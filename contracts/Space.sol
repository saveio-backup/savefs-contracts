//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./Type.sol";

contract Space is Initializable {
    mapping(address => UserSpace) userSpace; // walletAddr => UserSpace

    function initialize() public initializer {}

    function GetUserSpace(address walletAddr)
        public
        view
        returns (UserSpace memory)
    {
        require(userSpace[walletAddr].UpdateHeight > 0, "userSpace is empty");
        return userSpace[walletAddr];
    }

    function ManageUserSpace(UserSpaceParams memory params) public payable {
        
    }
}
