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

    struct AddParams {
        UserSpace oldUserSpace;
        uint64 addSize;
        uint64 addBlockCount;
        uint256 currentHeight;
        Setting setting;
        FileList fileList;
    }

    struct AddReturn {
        UserSpace newUserSpace;
        FileInfo[] updatedFiles;
        bool success;
    }

    function fsAddUserSpace(AddParams memory params)
        public
        view
        returns (AddReturn memory)
    {
        AddReturn memory ret;

        if (params.oldUserSpace.ExpireHeight == 0) {
            ret.newUserSpace.Used = 0;
            ret.newUserSpace.Remain = params.addSize;
            ret.newUserSpace.ExpireHeight =
                params.currentHeight +
                params.addBlockCount;
            ret.newUserSpace.Balance = 0;
            StorageFee memory fee = calcDepositFeeForUserSpace(
                ret.newUserSpace,
                params.setting,
                block.number
            );
            ret.newUserSpace.Balance =
                fee.TxnFee +
                fee.SpaceFee +
                fee.ValidationFee;
            ret.success = true;
            return ret;
        } else {
            uint256 newExpiredHeight = params.oldUserSpace.ExpireHeight +
                params.addBlockCount;
            uint64 newRemain = params.oldUserSpace.Remain + params.addSize;
            ret.newUserSpace.Used = params.oldUserSpace.Used;
            ret.newUserSpace.Remain = newRemain;
            ret.newUserSpace.ExpireHeight = newExpiredHeight;
            ret.newUserSpace.Balance = params.oldUserSpace.Balance;
            StorageFee memory fee1 = calcDepositFeeForUserSpace(
                params.oldUserSpace,
                params.setting,
                block.number
            );
            StorageFee memory fee2 = calcDepositFeeForUserSpace(
                ret.newUserSpace,
                params.setting,
                block.number
            );
            uint64 fee1Sum = fee1.TxnFee + fee1.SpaceFee + fee1.ValidationFee;
            uint64 fee2Sum = fee2.TxnFee + fee2.SpaceFee + fee2.ValidationFee;
            if (fee2Sum <= fee1Sum) {
                ret.success = false;
                return ret;
            }
            uint64 deposit = fee2Sum - fee1Sum;
            params.oldUserSpace.Used = 0;
            StorageFee memory feeForRemaining = calcDepositFeeForUserSpace(
                params.oldUserSpace,
                params.setting,
                block.number
            );
            uint64 feeForRemainingSum = feeForRemaining.TxnFee +
                feeForRemaining.SpaceFee +
                feeForRemaining.ValidationFee;
            if (params.oldUserSpace.Balance <= feeForRemainingSum) {
                ret.success = false;
                return ret;
            }
            uint64 availableInBalance = params.oldUserSpace.Balance -
                feeForRemainingSum;
            if (availableInBalance > deposit) {
                deposit = 0;
            } else {
                deposit = deposit - availableInBalance;
                ret.newUserSpace.Balance += deposit;
            }
            if (params.addBlockCount != 0) {
                (ret.updatedFiles, ret.success) = updateFilesForRenew(
                    params.fileList,
                    params.setting,
                    newExpiredHeight
                );
                if (!ret.success) {
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

    struct ProcessParams {
        UserSpaceParams userSpaceParams;
        UserSpace oldUserSpace;
        Setting setting;
    }

    struct ProcessReturn {
        UserSpace newUserSpace;
        TransferState state;
        FileInfo[] updatedFiles;
        bool success;
    }

    function processForUserSpaceOperations(ProcessParams memory params)
        public
        view
        returns (ProcessReturn memory)
    {
        ProcessReturn memory ret;

        uint64 transferIn;
        uint64 transferOut;

        FileList memory fileList = fs.GetFileList(params.userSpaceParams.Owner);
        uint64 ops = getUserSpaceOperationsFromParams(params.userSpaceParams);
        if (
            ops == UserspaceOps_Add_Add ||
            ops == UserspaceOps_Add_None ||
            ops == UserspaceOps_None_Add
        ) {
            AddParams memory addParams;
            addParams.oldUserSpace = params.oldUserSpace;
            addParams.addSize = params.userSpaceParams.Size.Value;
            addParams.addBlockCount = params.userSpaceParams.BlockCount.Value;
            addParams.currentHeight = block.number;
            addParams.setting = params.setting;
            addParams.fileList = fileList;

            AddReturn memory addRet = fsAddUserSpace(addParams);
            (ret.newUserSpace, transferIn, ret.updatedFiles, ret.success) = (
                addRet.newUserSpace,
                addRet.newUserSpace.Balance,
                addRet.updatedFiles,
                addRet.success
            );
            if (!ret.success) {
                return ret;
            }
        }
        if (
            ops == UserspaceOps_Revoke_Revoke ||
            ops == UserspaceOps_None_Revoke ||
            ops == UserspaceOps_Revoke_None
        ) {
            (ret.newUserSpace, transferOut, ret.success) = fsRevokeUserspace(
                params.oldUserSpace,
                params.userSpaceParams.Size.Value,
                params.userSpaceParams.BlockCount.Value,
                block.number,
                params.setting
            );
            if (!ret.success) {
                return ret;
            }
        }
        if (ops == UserspaceOps_Add_Revoke || ops == UserspaceOps_Revoke_Add) {
            (
                ret.newUserSpace,
                transferIn,
                transferOut,
                ret.updatedFiles,
                ret.success
            ) = processForUserSpaceOneAddOneRevoke(
                params.userSpaceParams,
                params.oldUserSpace,
                params.setting,
                fileList,
                ops
            );
            if (!ret.success) {
                return ret;
            }
        }
        if (ret.newUserSpace.ExpireHeight == 0) {
            return ret;
        }
        ret.newUserSpace.UpdateHeight = block.number;
        if (transferIn >= transferOut) {
            ret.state.From = params.userSpaceParams.WalletAddr;
            ret.state.To = address(this);
            ret.state.Value = transferIn - transferOut;
        } else {
            ret.state.From = address(this);
            ret.state.To = params.userSpaceParams.WalletAddr;
            ret.state.Value = transferOut - transferIn;
        }
        ret.success = true;
        return ret;
    }

    struct ChangeReturn {
        UserSpace newUserSpace;
        TransferState state;
        FileInfo[] updatedFiles;
        bool success;
    }

    function getUserspaceChange(UserSpaceParams memory params)
        public
        view
        returns (ChangeReturn memory)
    {
        ChangeReturn memory ret;
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
        ProcessParams memory processParams;
        processParams.userSpaceParams = params;
        processParams.oldUserSpace = oldUserSpace;
        processParams.setting = setting;
        ProcessReturn memory processRet = processForUserSpaceOperations(
            processParams
        );
        ret.newUserSpace = processRet.newUserSpace;
        ret.state = processRet.state;
        ret.updatedFiles = processRet.updatedFiles;
        ret.success = processRet.success;
        return ret;
    }

    function ManageUserSpace(UserSpaceParams memory params) public payable {
        ChangeReturn memory ret = getUserspaceChange(params);
        if (!ret.success) {
            revert UserspaceChangeError();
        }
        if (ret.state.Value > 0) {
            if (ret.state.From == address(this)) {
                payable(ret.state.To).transfer(ret.state.Value);
            } else {
                if (msg.value < ret.state.Value) {
                    revert InsufficientFunds();
                }
            }
        }
        for (uint256 i = 0; i < ret.updatedFiles.length; i++) {
            fs.UpdateFileInfo(ret.updatedFiles[i]);
        }
        userSpace[params.Owner] = ret.newUserSpace;
    }
}
