// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
import {Test, console} from "forge-std/Test.sol";
import {AuctionMarket} from "src/AuctionMarket.sol";
import {TitanNFT} from "src/TitanNFT.sol";
import {MockPriceFeed} from "src/mocks/MockPriceFeed.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract AuctionMarketTest  is Test {
    AuctionMarket public market;
    TitanNFT public nft;
    MockPriceFeed public priceFeed;
    address public admin = address(1);
    address public seller = address(2);
 address public bidder = address(3);
    function setUp() public {
        vm.startPrank(admin);
        priceFeed = new MockPriceFeed(3000 * 1e8);
        AuctionMarket implementation = new AuctionMarket();
        bytes memory initData=abi.encodeWithSelector(
            AuctionMarket.initialize.selector,
            address(priceFeed)
        );
        ERC1967Proxy proxy=new ERC1967Proxy(address(implementation),initData);
        market=AuctionMarket(address(proxy));
        nft=new TitanNFT();
        vm.stopPrank();
    }
    function test_CreateAuction()public{
        vm.prank(admin);
        nft.safeMint(seller, 0);
        vm.startPrank(seller);
        nft.approve(address(market),0);
        market.createAuction(address(nft),0,1000*1e18,1 days);

          (address auctionSeller,,,,,,,bool active) = market.auctions(1);
          assertEq(auctionSeller,seller);
          assertTrue(active);
    }
    function  test_BidSuccess()public{
        test_CreateAuction();
        vm.deal(bidder,10 ether);
        vm.startPrank(bidder);
        market.bid{value:0.5 ether}(1);
        vm.stopPrank();
      (,,,,address highestBidder, uint256 amount,,) =  market.auctions(1);
      assertEq(highestBidder,bidder);
      assertEq(amount,0.5 ether);
    }
}
