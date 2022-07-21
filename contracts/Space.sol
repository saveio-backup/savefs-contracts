//SPDX-License-Identifier: Unlicensed
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./type.sol";
import "./interface.sol";

contract Space is Initializable, ISpace {
    IConfig config;
    IFile fs;

    uint64 constant UserSpaceOps_None_None =
        (uint8(UserSpaceType.None) << 4) | uint8(UserSpaceType.None);
    uint64 constant UserspaceOps_None_Add =
        (uint8(UserSpaceType.None) << 4) | uint8(UserSpaceType.Add);
    uint64 constant UserspaceOps_None_Revoke =
        (uint8(UserSpaceType.None) << 4) | uint8(UserSpaceType.Revoke);
    uint64 constant UserspaceOps_Add_None =
        (uint8(UserSpaceType.Add) << 4) | uint8(UserSpaceType.None);
    uint64 constant UserspaceOps_Add_Add =
        (uint8(UserSpaceType.Add) << 4) | uint8(UserSpaceType.Add);
    uint64 constant UserspaceOps_Add_Revoke =
        (uint8(UserSpaceType.Add) << 4) | uint8(UserSpaceType.Revoke);
    uint64 constant UserspaceOps_Revoke_None =
        (uint8(UserSpaceType.Revoke) << 4) | uint8(UserSpaceType.None);
    uint64 constant UserspaceOps_Revoke_Add =
        (uint8(UserSpaceType.Revoke) << 4) | uint8(UserSpaceType.Add);
    uint64 constant UserspaceOps_Revoke_Revoke =
        (uint8(UserSpaceType.Revoke) << 4) | uint8(UserSpaceType.Revoke);

    mapping(address => UserSpace) userSpace; // walletAddr => UserSpace

    function initialize(IConfig _config, IFile _fs) public initializer {
        config = _config;
        fs = _fs;
    }

    function GetUserSpace(address walletAddr)
        public
        view
        virtual
        override
        returns (UserSpace memory)
    {
        return userSpace[walletAddr];
    }

    function GetUpdateCost(UserSpaceParams memory params)
        public
        view
        virtual
        override
        returns (TransferState memory)
    {
        ChangeReturn memory ret = getUserspaceChange(params);
        return ret.state;
    }

    function ManageUserSpace(UserSpaceParams memory params)
        public
        payable
        virtual
        override
    {
        ChangeReturn memory ret = getUserspaceChange(params);
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
        emit SetUserSpaceEvent(
            FsEvent.SET_USER_SPACE,
            block.number,
            params.WalletAddr,
            params.Size.Type,
            params.Size.Value,
            params.BlockCount.Type,
            params.BlockCount.Value
        );
    }

    function UpdateUserSpace(address walletAddr, UserSpace memory _userSpace)
        public
        payable
        virtual
        override
    {
        userSpace[walletAddr] = _userSpace;
    }

    function DeleteUserSpace(address walletAddr) public virtual override {
        UserSpace memory _userSpace = GetUserSpace(walletAddr);
        if (_userSpace.Used == 0 && _userSpace.Balance > 0) {
            payable(walletAddr).transfer(_userSpace.Balance);
        }
        if (
            _userSpace.ExpireHeight > 0 &&
            _userSpace.ExpireHeight <= block.number
        ) {
            deleteExpiredUserSpace(_userSpace, walletAddr);
        }
    }

    function isValidUserSpaceOperation(UserSpaceOperation memory op)
        private
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
        private
        pure
        returns (uint64)
    {
        uint64 n = (uint64(t1) << 4) | uint64(t2);
        return n;
    }

    function getUserSpaceOperationsFromParams(UserSpaceParams memory params)
        private
        pure
        returns (uint64)
    {
        bool r1 = isValidUserSpaceOperation(params.Size);
        if (!r1) {
            revert ParamsError();
        }
        bool r2 = isValidUserSpaceOperation(params.BlockCount);
        if (!r2) {
            revert ParamsError();
        }
        uint64 n = combineUserSpaceTypes(
            params.Size.Type,
            params.BlockCount.Type
        );
        return n;
    }

    function isRevokeUserSpace(UserSpaceParams memory params)
        private
        pure
        returns (bool)
    {
        return
            params.Size.Type == UserSpaceType.Revoke ||
            params.BlockCount.Type == UserSpaceType.Revoke;
    }

    function checkForUserSpaceRevoke(UserSpaceParams memory params)
        private
        view
        returns (bool)
    {
        bytes[] memory fileList = fs.GetFileList(params.Owner);
        if (fileList.length > 0) {
            return false;
        }
        if (params.WalletAddr != params.Owner) {
            return false;
        }
        return true;
    }

    function checkUserSpaceParams(UserSpaceParams memory params)
        private
        view
        returns (bool)
    {
        if (params.Size.Value <= 0) {
            return false;
        }
        if (params.BlockCount.Value <= 0) {
            return false;
        }
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

    function processExpiredUserSpace(
        UserSpace memory _userSpace,
        uint256 currentHeight
    ) private pure returns (UserSpace memory) {
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
    ) private pure returns (bool) {
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
    ) private view returns (StorageFee memory) {
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

    struct AddParams {
        UserSpace oldUserSpace;
        uint64 addSize;
        uint64 addBlockCount;
        uint256 currentHeight;
        Setting setting;
        bytes[] fileList;
    }

    struct AddReturn {
        UserSpace newUserSpace;
        FileInfo[] updatedFiles;
        bool success;
    }

    function AddUserSpace(AddParams memory params)
        private
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
                (ret.updatedFiles, ret.success) = fs.UpdateFilesForRenew(
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

    struct RevokeParams {
        UserSpace oldUserSpace;
        uint64 revokeSize;
        uint64 revokeBlockCount;
        uint256 currentHeight;
        Setting setting;
    }

    struct RevokeReturn {
        UserSpace newUserSpace;
        uint64 amount;
        bool success;
    }

    function RevokeUserSpace(RevokeParams memory params)
        private
        view
        returns (RevokeReturn memory)
    {
        RevokeReturn memory ret;
        if (params.oldUserSpace.Remain < params.revokeSize) {
            ret.success = false;
            return ret;
        }
        if (
            params.oldUserSpace.ExpireHeight - params.revokeBlockCount <
            params.currentHeight
        ) {
            ret.success = false;
            return ret;
        }
        UserSpace memory newUserSpace;
        newUserSpace.Used = params.oldUserSpace.Used;
        newUserSpace.Remain = params.oldUserSpace.Remain - params.revokeSize;
        newUserSpace.ExpireHeight =
            params.oldUserSpace.ExpireHeight -
            params.revokeBlockCount;
        newUserSpace.Balance = params.oldUserSpace.Balance;
        StorageFee memory fee1 = calcDepositFeeForUserSpace(
            params.oldUserSpace,
            params.setting,
            block.number
        );
        StorageFee memory fee2 = calcDepositFeeForUserSpace(
            newUserSpace,
            params.setting,
            block.number
        );
        uint64 fee1Sum = fee1.TxnFee + fee1.SpaceFee + fee1.ValidationFee;
        uint64 fee2Sum = fee2.TxnFee + fee2.SpaceFee + fee2.ValidationFee;
        if (fee1Sum <= fee2Sum) {
            ret.success = false;
            return ret;
        }
        uint64 amount = fee1Sum - fee2Sum;
        if (newUserSpace.Balance < amount) {
            ret.success = false;
            return ret;
        }
        newUserSpace.Balance -= amount;
        return ret;
    }

    struct ProcessRevokeParams {
        UserSpaceParams userSpaceParams;
        UserSpace oldUserspace;
        Setting setting;
        bytes[] fileList;
        uint64 ops;
    }

    struct ProcessRevokeReturn {
        UserSpace userSpace;
        uint64 addedAmount;
        uint64 revokedAmount;
        FileInfo[] update;
        bool success;
    }

    function processForUserSpaceOneAddOneRevoke(
        ProcessRevokeParams memory params
    ) private view returns (ProcessRevokeReturn memory) {
        ProcessRevokeReturn memory ret;
        uint64 addedSize;
        uint64 addedBlockCount;
        uint64 revokedSize;
        uint64 revokedBlockCount;
        if (params.ops == UserspaceOps_Add_Revoke) {
            addedSize = params.userSpaceParams.Size.Value;
            revokedBlockCount = params.userSpaceParams.BlockCount.Value;
        }
        if (params.ops == UserspaceOps_Revoke_Add) {
            revokedSize = params.userSpaceParams.Size.Value;
            addedBlockCount = params.userSpaceParams.BlockCount.Value;
        }
        AddParams memory addParams;
        addParams.oldUserSpace = params.oldUserspace;
        addParams.addSize = addedSize;
        addParams.addBlockCount = addedBlockCount;
        addParams.currentHeight = block.number;
        addParams.setting = params.setting;
        addParams.fileList = params.fileList;
        AddReturn memory addReturn = AddUserSpace(addParams);
        if (!addReturn.success) {
            ret.success = false;
            return ret;
        }
        RevokeParams memory revokeParams;
        revokeParams.oldUserSpace = addReturn.newUserSpace;
        revokeParams.revokeSize = revokedSize;
        revokeParams.revokeBlockCount = revokedBlockCount;
        revokeParams.currentHeight = block.number;
        revokeParams.setting = params.setting;
        RevokeReturn memory revokeReturn = RevokeUserSpace(revokeParams);
        if (!revokeReturn.success) {
            ret.success = false;
            return ret;
        }
        ret.userSpace = revokeReturn.newUserSpace;
        ret.addedAmount = addReturn.newUserSpace.Balance;
        ret.revokedAmount = revokeReturn.newUserSpace.Balance;
        ret.update = addReturn.updatedFiles;
        ret.success = true;
        return ret;
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
        private
        view
        returns (ProcessReturn memory)
    {
        ProcessReturn memory ret;

        uint64 transferIn;
        uint64 transferOut;

        bytes[] memory fileList = fs.GetFileList(params.userSpaceParams.Owner);
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

            AddReturn memory addRet = AddUserSpace(addParams);
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
            RevokeParams memory revokeParams;
            revokeParams.oldUserSpace = params.oldUserSpace;
            revokeParams.revokeSize = params.userSpaceParams.Size.Value;
            revokeParams.revokeBlockCount = params
                .userSpaceParams
                .BlockCount
                .Value;
            revokeParams.currentHeight = block.number;
            revokeParams.setting = params.setting;
            RevokeReturn memory revokeReturn = RevokeUserSpace(revokeParams);
            (ret.newUserSpace, transferOut, ret.success) = (
                revokeReturn.newUserSpace,
                revokeReturn.amount,
                revokeReturn.success
            );
            if (!ret.success) {
                return ret;
            }
        }
        if (ops == UserspaceOps_Add_Revoke || ops == UserspaceOps_Revoke_Add) {
            ProcessRevokeParams memory processRevokeParams;
            processRevokeParams.userSpaceParams = params.userSpaceParams;
            processRevokeParams.oldUserspace = params.oldUserSpace;
            processRevokeParams.setting = params.setting;
            processRevokeParams.fileList = fileList;
            processRevokeParams.ops = ops;
            ProcessRevokeReturn
                memory processRevokeReturn = processForUserSpaceOneAddOneRevoke(
                    processRevokeParams
                );
            (
                ret.newUserSpace,
                transferIn,
                transferOut,
                ret.updatedFiles,
                ret.success
            ) = (
                processRevokeReturn.userSpace,
                processRevokeReturn.addedAmount,
                processRevokeReturn.revokedAmount,
                processRevokeReturn.update,
                processRevokeReturn.success
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
    }

    function getUserspaceChange(UserSpaceParams memory params)
        private
        view
        returns (ChangeReturn memory)
    {
        ChangeReturn memory ret;
        Setting memory setting = config.GetSetting();
        bool checkRes = checkUserSpaceParams(params);
        if (!checkRes) {
            revert UserspaceChangeError(1);
        }
        UserSpace memory oldUserSpace = GetUserSpace(params.Owner);
        if (
            oldUserSpace.ExpireHeight > 0 &&
            oldUserSpace.ExpireHeight <= block.number
        ) {
            oldUserSpace = processExpiredUserSpace(oldUserSpace, block.number);
        }
        if (
            oldUserSpace.ExpireHeight == 0 ||
            oldUserSpace.ExpireHeight == block.number
        ) {
            bool b = checkForFirstUserSpaceOperation(setting, params);
            if (!b) {
                revert UserspaceChangeError(2);
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
        return ret;
    }

    function deleteExpiredUserSpace(
        UserSpace memory _userSpace,
        address walletAddr
    ) private {
        Setting memory setting = config.GetSetting();
        if (
            setting.DefaultProvePeriod + _userSpace.ExpireHeight > block.number
        ) {
            revert UserspaceDeleteError();
        }
        bytes[] memory deletedFiles;
        uint64 amount;
        bool success;
        bytes[] memory fileList = fs.GetFileList(walletAddr);
        StorageType[] memory sType = new StorageType[](1);
        sType[0] = StorageType.Normal;
        (deletedFiles, amount, success) = fs.DeleteExpiredFilesFromList(
            fileList,
            walletAddr,
            sType
        );
        bytes[] memory unsettledList = fs.GetUnSettledFileList(walletAddr);
        (deletedFiles, amount, success) = fs.DeleteExpiredFilesFromList(
            unsettledList,
            walletAddr,
            sType
        );
        for (uint256 i = 0; i < deletedFiles.length; i++) {
            for (uint256 j = 0; j < unsettledList.length; j++) {
                if (keccak256(deletedFiles[i]) == keccak256(unsettledList[j])) {
                    delete unsettledList[j];
                }
            }
        }
        fs.UpdateFileList(walletAddr, unsettledList);
    }
}
