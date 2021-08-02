// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.6.0;

contract ISwapper {
    function swap(
        address fromAssetHash, uint64 toPoolId, uint64 toChainId, 
        bytes memory toAssetHash, bytes memory toAddress, 
        uint amount, uint minOutAmount, uint fee, uint id
        ) external payable returns (bool){}
    
    function add_liquidity(
        address fromAssetHash, uint64 toPoolId, uint64 toChainId, 
        bytes memory toAddress, uint amount, uint minOutAmount, 
        uint fee, uint id
        ) external payable returns (bool){}
    
    function remove_liquidity(
        address fromAssetHash, uint64 toPoolId, uint64 toChainId, 
        bytes memory toAssetHash, bytes memory toAddress, uint amount, 
        uint minOutAmount, uint fee, uint id
        ) external payable returns (bool){}
}

