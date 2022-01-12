//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./Type.sol";
import "./Config.sol";
import "./FileSystem.sol";

contract Space is Initializable {
    Config config;
    FileSystem fs;

    uint64 UserSpaceOps_None_None =
        (uint8(UserSpaceType.None) << 4) | uint8(UserSpaceType.None);
    uint64 UserspaceOps_None_Add =
        (uint8(UserSpaceType.None) << 4) | uint8(UserSpaceType.Add);
    uint64 UserspaceOps_None_Revoke =
        (uint8(UserSpaceType.None) << 4) | uint8(UserSpaceType.Revoke);
    uint64 UserspaceOps_Add_None =
        (uint8(UserSpaceType.Add) << 4) | uint8(UserSpaceType.None);
    uint64 UserspaceOps_Add_Add =
        (uint8(UserSpaceType.Add) << 4) | uint8(UserSpaceType.Add);
    uint64 UserspaceOps_Add_Revoke =
        (uint8(UserSpaceType.Add) << 4) | uint8(UserSpaceType.Revoke);
    uint64 UserspaceOps_Revoke_None =
        (uint8(UserSpaceType.Revoke) << 4) | uint8(UserSpaceType.None);
    uint64 UserspaceOps_Revoke_Add =
        (uint8(UserSpaceType.Revoke) << 4) | uint8(UserSpaceType.Add);
    uint64 UserspaceOps_Revoke_Revoke =
        (uint8(UserSpaceType.Revoke) << 4) | uint8(UserSpaceType.Revoke);

    mapping(address => UserSpace) userSpace; // walletAddr => UserSpace

    error ParamsError();
    error FirstUserSpaceOperation();

    function initialize(Config _config, FileSystem _fs) public initializer {
        config = _config;
        fs = _fs;
    }

    function GetUserSpace(address walletAddr)
        public
        view
        returns (UserSpace memory)
    {
        require(userSpace[walletAddr].UpdateHeight > 0, "userSpace is empty");
        return userSpace[walletAddr];
    }

    function isValidUserSpaceOperation(UserSpaceOperation memory op)
        public
        pure
        returns (bool)
    {
        if (op.Type == UserSpaceType.Revoke) {
            if (op.Value == 0) {
                return false;
            }
            return true;
        }
        if (op.Type == UserSpaceType.Add) {
            if (op.Value == 0) {
                return false;
            }
            return true;
        }
        if (op.Type == UserSpaceType.None) {
            if (op.Value != 0) {
                return false;
            }
            return true;
        }
        return false;
    }

    function combineUserSpaceTypes(UserSpaceType t1, UserSpaceType t2)
        public
        pure
        returns (uint64)
    {
        uint64 n = (uint8(t1) << 4) | uint8(t2);
        return n;
    }

    function getUserSpaceOperationsFromParams(UserSpaceParams memory params)
        public
        pure
        returns (uint64)
    {
        bool r1 = isValidUserSpaceOperation(params.Size);
        if (r1) {
            revert ParamsError();
        }
        bool r2 = isValidUserSpaceOperation(params.BlockCount);
        if (r2) {
            revert ParamsError();
        }
        uint64 n = combineUserSpaceTypes(
            params.Size.Type,
            params.BlockCount.Type
        );
        return n;
    }

    function isRevokeUserSpace(UserSpaceParams memory params)
        public
        pure
        returns (bool)
    {
        return
            params.Size.Type == UserSpaceType.Revoke ||
            params.BlockCount.Type == UserSpaceType.Revoke;
    }

    function checkForUserSpaceRevoke(UserSpaceParams memory params)
        public
        view
        returns (bool)
    {
        FileList memory fileList = fs.GetFileList(params.Owner);
        if (fileList.FileNum > 0) {
            return false;
        }
        if (params.WalletAddr != params.Owner) {
            return false;
        }
        return true;
    }

    function checkUserSpaceParams(UserSpaceParams memory params)
        public
        view
        returns (bool)
    {
        require(
            params.Size.Value > 0,
            "params.Size.value must be greater than 0"
        );
        require(
            params.BlockCount.Value > 0,
            "params.BlockCount.value must be greater than 0"
        );
        uint64 n = getUserSpaceOperationsFromParams(params);
        if (n == uint64(UserSpaceType.None)) {
            return false;
        }
        bool b1 = isRevokeUserSpace(params);
        if (b1) {
            bool b2 = checkForUserSpaceRevoke(params);
            if (!b2) {
                return false;
            }
        }
        return true;
    }

    function getOldUserSpace(address owner)
        public
        view
        returns (UserSpace memory)
    {
        return userSpace[owner];
    }

    function processExpiredUserSpace(
        UserSpace memory _userSpace,
        uint256 currentHeight
    ) public pure returns (UserSpace memory) {
        _userSpace.Used = 0;
        _userSpace.Remain = 0;
        //_userSpace.Balance = 0
        _userSpace.ExpireHeight = currentHeight;
        _userSpace.UpdateHeight = currentHeight;
        return _userSpace;
    }

    function checkForFirstUserSpaceOperation(
        Setting memory setting,
        UserSpaceParams memory params
    ) public view returns (bool) {
        uint64 ops = getUserSpaceOperationsFromParams(params);
        if (ops != UserspaceOps_Add_Add) {
            return false;
        }
        if (params.Size.Value == 0 || params.BlockCount.Value == 0) {
            return false;
        }
        if (params.BlockCount.Value < setting.DefaultProvePeriod) {
            return false;
        }
        return true;
    }

    function getUserspaceChange(UserSpaceParams memory params) public view {
        Setting memory setting = config.GetSetting();
        bool checkRes = checkUserSpaceParams(params);
        if (!checkRes) {
            revert ParamsError();
        }
        UserSpace memory oldUserSpace = getOldUserSpace(params.Owner);
        if (
            oldUserSpace.ExpireHeight > 0 &&
            oldUserSpace.ExpireHeight <= block.number
        ) {
            oldUserSpace = processExpiredUserSpace(oldUserSpace, block.number);
        }
        if (
            oldUserSpace.ExpireHeight > 0 &&
            oldUserSpace.ExpireHeight == block.number
        ) {
            bool b = checkForFirstUserSpaceOperation(setting, params);
            if (!b) {
                revert FirstUserSpaceOperation();
            }
        }
        // TODO
        // processForUserSpaceOperations();
    }

    function ManageUserSpace(UserSpaceParams memory params) public payable {
        // TODO
    }
}
