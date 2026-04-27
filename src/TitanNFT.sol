// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// 具名导入 OpenZeppelin 库，防止 Linter 报错
import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract  TitanNFT is ERC721,Ownable {
     // 状态变量：记录下一个生成的代币 ID
   uint256 private _nextTokenId;
   // 构造函数：初始化代币名称 "Titan NFT" 和符号 "TNFT"
    // 同时显式传递部署者为初始所有者
   constructor()ERC721("Titan NFT", "TNFT")Ownable(msg.sender){}
  // 修改后的 TitanNFT.sol
function safeMint(address to, uint256 tokenId) public onlyOwner {
    // 删掉内部计数逻辑，直接铸造传入的 ID
    _safeMint(to, tokenId);
}
}