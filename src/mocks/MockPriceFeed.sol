// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
contract MockPriceFeed {
    int256 private _price;
    constructor(int256 initialPrice){
        _price=initialPrice;
    }
    function updatePrice(int256 newPrice)public{
        _price=newPrice;
    }
    function latestRoundData()external view returns   (uint80, int256, uint256, uint256, uint80){
         return (0, _price, 0, block.timestamp, 0);
    }
}