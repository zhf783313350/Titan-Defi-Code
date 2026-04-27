// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
// 引入 UUPS 升级相关的具名库
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
// 引入 NFT 标准接口
import {IERC721} from "@openzeppelin/contracts/token/ERC721/IERC721.sol";
// 引入 Chainlink 价格喂送接口
import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

contract AuctionMarket is
    Initializable,
    UUPSUpgradeable,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable
{
    struct Auction {
        address seller; // 卖家地址
        address nftAddr; // NFT 合约地址
        uint256 tokenId; // NFT ID
        uint256 minPriceInUsd; // 设定的最低美元起拍价 (18位精度)
        address highestBidder; // 最高出价者地址
        uint256 highestBidAmount; // 最高出价金额 (单位: Wei)
        uint256 endTime; // 结束时间戳
        bool isActive; // 拍卖是否处于激活状态
    }
event AuctionCreated(uint256 indexed auctionId, address indexed seller, uint256 tokenId, uint256 minPriceInUsd);
    mapping(uint256 => Auction) public auctions;
    uint256 public totalAuctions;
    AggregatorV3Interface internal ethUsdPriceFeed;
event BidPlaced(uint256 indexed auctionId, address indexed bidder, uint256 amount);
    constructor() {
        _disableInitializers();
    }

    function initialize(address _priceFeed) public initializer {
        __Ownable_init(msg.sender);
        __ReentrancyGuard_init();
        ethUsdPriceFeed = AggregatorV3Interface(_priceFeed);
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyOwner {}

   function getLatestPrice() public view returns (uint256) {
     (, int256 price, , , ) = ethUsdPriceFeed.latestRoundData();
        require(price > 0, "Invalid price");
          // 添加下面这行指令，告诉编译器我已经检查过这个转换是安全的
    // forge-lint: disable-next-line(unsafe-typecast)
        return uint256(price);
    }

    function createAuction( address _nftAddr, uint256 _tokenId, uint256 _minPriceUsd, uint256 _duration
    ) external {
        // 步骤1: 托管 NFT (用户必须先执行 Approve)
        IERC721(_nftAddr).transferFrom(msg.sender, address(this), _tokenId);

        totalAuctions++;
        // 步骤2: 存入拍卖记录
        auctions[totalAuctions] = Auction({
            seller: msg.sender,
            nftAddr: _nftAddr,
            tokenId: _tokenId,
            minPriceInUsd: _minPriceUsd,
            highestBidder: address(0),
            highestBidAmount: 0,
            endTime: block.timestamp + _duration,
            isActive: true
        });
        emit AuctionCreated(totalAuctions, msg.sender, _tokenId, _minPriceUsd);
    }

    function bid(uint256 _auctionId) external payable nonReentrant {
        Auction storage auction = auctions[_auctionId];
        require(auction.isActive, "Auction not active");
        require(block.timestamp < auction.endTime, "Auction expired");
        require(
            msg.value > auction.highestBidAmount,
            "Bid lower than current highest"
        );

        uint256  ethPrice = getLatestPrice();
        require(ethPrice > 0, "Invalid oracle price");
        uint256 bidValueInUsd = (uint256(ethPrice) * msg.value) / 1e8;
        require(
            bidValueInUsd >= auction.minPriceInUsd,
            "Bid value below USD floor"
        );

        if (auction.highestBidder != address(0)) {
            (bool success, ) = payable(auction.highestBidder).call{
                value: auction.highestBidAmount
            }("");
            require(success, "Refund failed");
        }
        auction.highestBidder = msg.sender;
        auction.highestBidAmount = msg.value;
         emit BidPlaced(_auctionId, msg.sender, msg.value);
    }

    function endAuction(uint256 _auctionId) external nonReentrant {
        Auction storage auction = auctions[_auctionId];
        require(auction.isActive, "Already closed");
        require(block.timestamp >= auction.endTime, "Not finished");
        auction.isActive = false;
        if (auction.highestBidder != address(0)) {
            IERC721(auction.nftAddr).transferFrom(
                address(this),
                auction.highestBidder,
                auction.tokenId
            );
            (bool success, ) = payable(auction.seller).call{value: auction.highestBidAmount}("");
            require(success, "Transfer to seller failed");
        } else {
            IERC721(auction.nftAddr).transferFrom(
                address(this),
                auction.seller,
                auction.tokenId
            );
        }
    }
}
