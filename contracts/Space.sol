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
    error FirstUserSpaceOperationError();
    error UserspaceChangeError();
    error InsufficientFunds();

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

    function calcDepositFeeForUserSpace(
        UserSpace memory _userSpace,
        Setting memory setting,
        uint256 currentHeight
    ) public view returns (StorageFee memory) {
        UploadOption memory uploadOpt;
        uploadOpt.FileSize = _userSpace.Used + _userSpace.Remain;
        uploadOpt.ProveInterval = setting.DefaultProvePeriod;
        uploadOpt.ExpiredHeight = _userSpace.ExpireHeight;
        uploadOpt.CopyNum = setting.DefaultCopyNum;
        StorageFee memory fee = fs.CalcDepositFee(
            uploadOpt,
            setting,
            currentHeight
        );
        return fee;
    }

    function updateFilesForRenew(
        FileList memory fileList,
        Setting memory setting,
        uint256 newExpireHeight
    ) public view returns (FileInfo[] memory, bool) {
        FileInfo[] memory fileInfo;
        for (uint256 i = 0; i < fileList.List.length; i++) {
            FileInfo memory _fileInfo = fs.GetFileInfo(fileList.List[i]);
            if (_fileInfo.StorageType_ != StorageType.Normal) {
                continue;
            }
            if (newExpireHeight <= _fileInfo.ExpiredHeight) {
                continue;
            }
            bool res = fs.UpdateFileInfoForRenew(
                setting,
                newExpireHeight,
                _fileInfo
            );
            if (!res) {
                return (fileInfo, false);
            }
        }
        return (fileInfo, true);
    }

    struct AddReturn {
        UserSpace newUserSpace;
        FileInfo[] updatedFiles;
        bool success;
    }

    function fsAddUserSpace(
        UserSpace memory oldUserSpace,
        uint64 addSize,
        uint64 addBlockCount,
        uint256 currentHeight,
        Setting memory setting,
        FileList memory fileList
    ) public view returns (AddReturn memory) {
        AddReturn memory ret;

        // avoid stack too deep error
        UserSpace memory _oldUserSpace;
        _oldUserSpace.Used = oldUserSpace.Used;
        _oldUserSpace.Remain = oldUserSpace.Remain;
        _oldUserSpace.Balance = oldUserSpace.Balance;
        _oldUserSpace.ExpireHeight = oldUserSpace.ExpireHeight;
        _oldUserSpace.UpdateHeight = oldUserSpace.UpdateHeight;

        // avoid stack too deep error
        Setting memory _setting;
        _setting.GasPrice = setting.GasPrice;
        _setting.GasPerGBPerBlock = setting.GasPerGBPerBlock;
        _setting.GasPerKBPerBlock = setting.GasPerKBPerBlock;
        _setting.GasForChallenge = setting.GasForChallenge;
        _setting.MaxProveBlockNum = setting.MaxProveBlockNum;
        _setting.MinVolume = setting.MinVolume;
        _setting.DefaultProvePeriod = setting.DefaultProvePeriod;
        _setting.DefaultProveLevel = setting.DefaultProveLevel;
        _setting.DefaultCopyNum = setting.DefaultCopyNum;

        if (_oldUserSpace.ExpireHeight == 0) {
            ret.newUserSpace.Used = 0;
            ret.newUserSpace.Remain = addSize;
            ret.newUserSpace.ExpireHeight = currentHeight + addBlockCount;
            ret.newUserSpace.Balance = 0;
            StorageFee memory fee = calcDepositFeeForUserSpace(
                ret.newUserSpace,
                _setting,
                block.number
            );
            ret.newUserSpace.Balance =
                fee.TxnFee +
                fee.SpaceFee +
                fee.ValidationFee;
            ret.success = true;
            return ret;
        } else {
            uint256 newExpiredHeight = _oldUserSpace.ExpireHeight +
                addBlockCount;
            uint64 newRemain = _oldUserSpace.Remain + addSize;
            ret.newUserSpace.Used = _oldUserSpace.Used;
            ret.newUserSpace.Remain = newRemain;
            ret.newUserSpace.ExpireHeight = newExpiredHeight;
            ret.newUserSpace.Balance = _oldUserSpace.Balance;
            StorageFee memory fee1 = calcDepositFeeForUserSpace(
                _oldUserSpace,
                _setting,
                block.number
            );
            StorageFee memory fee2 = calcDepositFeeForUserSpace(
                ret.newUserSpace,
                _setting,
                block.number
            );
            uint64 fee1Sum = fee1.TxnFee + fee1.SpaceFee + fee1.ValidationFee;
            uint64 fee2Sum = fee2.TxnFee + fee2.SpaceFee + fee2.ValidationFee;
            if (fee2Sum <= fee1Sum) {
                ret.success = false;
                return ret;
            }
            uint64 deposit = fee2Sum - fee1Sum;
            _oldUserSpace.Used = 0;
            StorageFee memory feeForRemaining = calcDepositFeeForUserSpace(
                _oldUserSpace,
                _setting,
                block.number
            );
            uint64 feeForRemainingSum = feeForRemaining.TxnFee +
                feeForRemaining.SpaceFee +
                feeForRemaining.ValidationFee;
            if (_oldUserSpace.Balance <= feeForRemainingSum) {
                ret.success = false;
                return ret;
            }
            uint64 availableInBalance = _oldUserSpace.Balance -
                feeForRemainingSum;
            if (availableInBalance > deposit) {
                deposit = 0;
            } else {
                deposit = deposit - availableInBalance;
                ret.newUserSpace.Balance += deposit;
            }
            if (addBlockCount != 0) {
                bool res;
                (ret.updatedFiles, res) = updateFilesForRenew(
                    fileList,
                    _setting,
                    newExpiredHeight
                );
                if (!res) {
                    ret.success = false;
                    return ret;
                }
            }
            ret.success = true;
            return ret;
        }
    }

    function fsRevokeUserspace(
        UserSpace memory oldUserspace,
        uint64 revokeSize,
        uint64 revokeBlockCount,
        uint256 currentHeight,
        Setting memory setting
    )
        public
        view
        returns (
            UserSpace memory,
            uint64,
            bool
        )
    {
        // TODO
    }

    function processForUserSpaceOneAddOneRevoke(
        UserSpaceParams memory params,
        UserSpace memory oldUserspace,
        Setting memory setting,
        FileList memory fileList,
        uint64 ops
    )
        public
        view
        returns (
            UserSpace memory,
            uint64,
            uint64,
            FileInfo[] memory,
            bool
        )
    {
        // TODO
    }

    function processForUserSpaceOperations(
        UserSpaceParams memory params,
        UserSpace memory oldUserSpace,
        Setting memory setting
    )
        public
        view
        returns (
            UserSpace memory,
            TransferState memory,
            FileInfo[] memory,
            bool
        )
    {
        UserSpace memory newUserSpace;
        FileInfo[] memory updatedFiles;
        uint64 transferIn;
        uint64 transferOut;
        TransferState memory state;
        bool res;

        // avoid stack too deep error
        UserSpaceParams memory _params;
        _params.WalletAddr = params.WalletAddr;
        _params.Owner = params.Owner;
        _params.Size = params.Size;
        _params.BlockCount = params.BlockCount;
        // avoid stack too deep error
        Setting memory _setting;
        _setting.GasPrice = setting.GasPrice;
        _setting.GasPerGBPerBlock = setting.GasPerGBPerBlock;
        _setting.GasPerKBPerBlock = setting.GasPerKBPerBlock;
        _setting.GasForChallenge = setting.GasForChallenge;
        _setting.MaxProveBlockNum = setting.MaxProveBlockNum;
        _setting.MinVolume = setting.MinVolume;
        _setting.DefaultProvePeriod = setting.DefaultProvePeriod;
        _setting.DefaultProveLevel = setting.DefaultProveLevel;
        _setting.DefaultCopyNum = setting.DefaultCopyNum;
        // avoid stack too deep error
        UserSpace memory _oldUserSpace;
        _oldUserSpace.Used = oldUserSpace.Used;
        _oldUserSpace.Remain = oldUserSpace.Remain;
        _oldUserSpace.Balance = oldUserSpace.Balance;
        _oldUserSpace.ExpireHeight = oldUserSpace.ExpireHeight;
        _oldUserSpace.UpdateHeight = oldUserSpace.UpdateHeight;

        FileList memory fileList = fs.GetFileList(_params.Owner);
        uint64 ops = getUserSpaceOperationsFromParams(_params);
        if (
            ops == UserspaceOps_Add_Add ||
            ops == UserspaceOps_Add_None ||
            ops == UserspaceOps_None_Add
        ) {
            AddReturn memory ret = fsAddUserSpace(
                _oldUserSpace,
                _params.Size.Value,
                _params.BlockCount.Value,
                block.number,
                _setting,
                fileList
            );
            (newUserSpace, transferIn, updatedFiles, res) = (
                ret.newUserSpace,
                ret.newUserSpace.Balance,
                ret.updatedFiles,
                ret.success
            );
            if (!res) {
                return (newUserSpace, state, updatedFiles, false);
            }
        }
        if (
            ops == UserspaceOps_Revoke_Revoke ||
            ops == UserspaceOps_None_Revoke ||
            ops == UserspaceOps_Revoke_None
        ) {
            (newUserSpace, transferOut, res) = fsRevokeUserspace(
                _oldUserSpace,
                _params.Size.Value,
                _params.BlockCount.Value,
                block.number,
                _setting
            );
            if (!res) {
                return (newUserSpace, state, updatedFiles, false);
            }
        }
        if (ops == UserspaceOps_Add_Revoke || ops == UserspaceOps_Revoke_Add) {
            (
                newUserSpace,
                transferIn,
                transferOut,
                updatedFiles,
                res
            ) = processForUserSpaceOneAddOneRevoke(
                _params,
                _oldUserSpace,
                _setting,
                fileList,
                ops
            );
            if (!res) {
                return (newUserSpace, state, updatedFiles, false);
            }
        }
        if (newUserSpace.ExpireHeight == 0) {
            return (newUserSpace, state, updatedFiles, false);
        }
        newUserSpace.UpdateHeight = block.number;
        if (transferIn >= transferOut) {
            state.From = _params.WalletAddr;
            state.To = address(this);
            state.Value = transferIn - transferOut;
        } else {
            state.From = address(this);
            state.To = _params.WalletAddr;
            state.Value = transferOut - transferIn;
        }
        return (newUserSpace, state, updatedFiles, true);
    }

    function getUserspaceChange(UserSpaceParams memory params)
        public
        view
        returns (
            UserSpace memory,
            TransferState memory,
            FileInfo[] memory,
            bool
        )
    {
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
                revert FirstUserSpaceOperationError();
            }
        }
        UserSpace memory newUserSpace;
        TransferState memory state;
        FileInfo[] memory updatedFiles;
        bool res;
        (
            newUserSpace,
            state,
            updatedFiles,
            res
        ) = processForUserSpaceOperations(params, oldUserSpace, setting);
        return (newUserSpace, state, updatedFiles, res);
    }

    function ManageUserSpace(UserSpaceParams memory params) public payable {
        UserSpace memory newUserSpace;
        TransferState memory state;
        FileInfo[] memory updatedFiles;
        bool res;
        (newUserSpace, state, updatedFiles, res) = getUserspaceChange(params);
        if (!res) {
            revert UserspaceChangeError();
        }
        if (state.Value > 0) {
            if (state.From == address(this)) {
                payable(state.To).transfer(state.Value);
            } else {
                if (msg.value < state.Value) {
                    revert InsufficientFunds();
                }
            }
        }
        for (uint256 i = 0; i < updatedFiles.length; i++) {
            fs.UpdateFileInfo(updatedFiles[i]);
        }
        userSpace[params.Owner] = newUserSpace;
    }
}
