// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// 引入 Foundry 官方提供的脚本工具
import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";

// 引入我们自己的合约代码
// 注意：这里使用了你在 remappings 里配置过的 src/ 路径
import {AuctionMarket} from "src/AuctionMarket.sol";
import {TitanNFT} from "src/TitanNFT.sol";
import {MockPriceFeed} from "src/mocks/MockPriceFeed.sol";

// 引入 OpenZeppelin 的代理合约，用于实现 UUPS 升级模式
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
contract DeployScript is Script{
     function run() external {
      uint256 deployerPrivateKey =vm.envUint("PRIVATE_KEY");
      vm.startBroadcast(deployerPrivateKey);
      MockPriceFeed priceFeed =new MockPriceFeed(3500 * 1e8);
      console.log("MockPriceFeed deployed at:", address(priceFeed));
      TitanNFT nft=new TitanNFT();
       console.log("TitanNFT deployed at:", address(nft));
       AuctionMarket marketImpl =new AuctionMarket();
       console.log("Market Implementation deployed at:", address(marketImpl));
     bytes memory initData=abi.encodeWithSelector(
        AuctionMarket.initialize.selector,
        address(priceFeed)
     );
     ERC1967Proxy proxy=new ERC1967Proxy(address(marketImpl),initData);
      console.log("----------------------------------------------");
      console.log(unicode"Market Proxy (最终使用这个地址!):", address(proxy));
        console.log("----------------------------------------------");
        vm.stopBroadcast();
     }
}