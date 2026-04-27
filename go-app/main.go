package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"titan-defi-code/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ 读取 .env 失败")
	}

	rpcURL := os.Getenv("SEPOLIA_RPC_URL")
	wssURL := os.Getenv("SEPOLIA_WSS_URL")

	client, _ := ethclient.Dial(rpcURL)
	wssClient, _ := ethclient.Dial(wssURL)

// 1. 拍卖行代理地址 (Proxy) - 这是最重要的入口
marketAddr := common.HexToAddress("0x9EE55f0B08828542127046Df2e11f9F4e4AEa9bC")

// 2. TitanNFT 合约地址
nftAddr := common.HexToAddress("0x27E70283dcEb16AAa553335641C30a2CC34C2632")

	marketWatcher, _ := bindings.NewAuctionMarket(marketAddr, wssClient)

	// 【核心监听协程】严格按照要求的格式打印所有交易信息
	go func() {
		fmt.Println("🎧 [监听服务] 已启动，正在实时捕捉 Sepolia 链上动态...")
		createChan := make(chan *bindings.AuctionMarketAuctionCreated)
		bidChan := make(chan *bindings.AuctionMarketBidPlaced)

		marketWatcher.WatchAuctionCreated(nil, createChan, nil, nil)
		marketWatcher.WatchBidPlaced(nil, bidChan, nil, nil)

		for {
			select {
			case event := <-createChan:
				fmt.Println("\n==================== 🔔 监听到【新拍卖创建】事件 ====================")
				fmt.Printf("【交易元数据】\n")
				fmt.Printf("   交易哈希: %s\n", event.Raw.TxHash.Hex())
				fmt.Printf("   区块高度: %d\n", event.Raw.BlockNumber)
				fmt.Printf("   日志索引: %d\n", event.Raw.Index)
				fmt.Printf("\n【业务明细】\n")
				fmt.Printf("   拍卖 ID: %s\n", event.AuctionId.String())
				fmt.Printf("   卖家地址: %s\n", event.Seller.Hex())
				fmt.Printf("   NFT 编号: %s\n", event.TokenId.String())
				fmt.Printf("   起拍价格: %s (Wei)\n", event.MinPriceInUsd.String())
				fmt.Println("==============================================================")

			case event := <-bidChan:
				fmt.Println("\n==================== 💰 监听到【有人出价】事件 ====================")
				fmt.Printf("【交易元数据】\n")
				fmt.Printf("   交易哈希: %s\n", event.Raw.TxHash.Hex())
				fmt.Printf("   区块高度: %d\n", event.Raw.BlockNumber)
				fmt.Printf("   日志索引: %d\n", event.Raw.Index)
				fmt.Printf("\n【业务明细】\n")
				fmt.Printf("   拍卖 ID: %s\n", event.AuctionId.String())
				fmt.Printf("   出价人地址: %s\n", event.Bidder.Hex())
				fmt.Printf("   出价金额: %s (Wei)\n", event.Amount.String())
				fmt.Println("==============================================================")
			}
		}
	}()

	// 执行一次自动化挂单流程
	privateKeyHex := strings.TrimSpace(strings.TrimPrefix(os.Getenv("PRIVATE_KEY"), "0x"))
	privateKey, _ := crypto.HexToECDSA(privateKeyHex)
	chainID, _ := client.NetworkID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	market, _ := bindings.NewAuctionMarket(marketAddr, client)
	nft, _ := bindings.NewTitanNFT(nftAddr, client)

	tokenId := big.NewInt(time.Now().UnixNano() / 1e6)
	fmt.Printf("🚀 [业务启动] 操作账户: %s\n", auth.From.Hex())

	// 1. 铸造
	mintTx, _ := nft.SafeMint(auth, auth.From, tokenId)
	bind.WaitMined(context.Background(), client, mintTx)
	fmt.Printf("   ✅ [1/3] 铸造完成 #%s\n", tokenId.String())

	// 2. 授权
	auth.Nonce = nil
	approveTx, _ := nft.Approve(auth, marketAddr, tokenId)
	bind.WaitMined(context.Background(), client, approveTx)
	fmt.Println("   ✅ [2/3] 市场授权成功")

	// 3. 挂单
	auth.Nonce = nil
	minPrice := new(big.Int).Mul(big.NewInt(100), big.NewInt(1e18))
	auctionTx, _ := market.CreateAuction(auth, nftAddr, tokenId, minPrice, big.NewInt(86400))
	fmt.Printf("   [3/3] 挂单已发出，等待链上确认... Hash: %s\n", auctionTx.Hash().Hex())
	bind.WaitMined(context.Background(), client, auctionTx)

	fmt.Println("\n✅ 自动化任务完成。程序进入工业级常驻监听模式，按 Ctrl+C 退出。")
	select {}
}