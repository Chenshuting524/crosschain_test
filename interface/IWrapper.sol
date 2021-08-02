// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.6.0;

contract IWrapper {
    function lock(
        address fromAsset, uint64 toChainId, bytes memory toAddress, 
        uint amount, uint fee, uint id
        ) external{}
}