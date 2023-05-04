// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// FileInfo is an auto generated low-level Go binding around an user-defined struct.
type FileInfo struct {
	FileHash       []byte
	FileOwner      common.Address
	FileDesc       []byte
	Privilege      uint64
	FileBlockNum   uint64
	FileBlockSize  uint64
	ProveInterval  uint64
	ProveTimes     uint64
	ExpiredHeight  *big.Int
	CopyNum        uint64
	Deposit        uint64
	FileProveParam FileProveParam
	ProveBlockNum  uint64
	BlockHeight    *big.Int
	ValidFlag      bool
	StorageType    uint8
	RealFileSize   uint64
	PrimaryNodes   []common.Address
	CandidateNodes []common.Address
	BlocksRoot     []byte
	ProveLevel     uint8
	IsPlotFile     bool
	PlotInfo       PlotInfo
}

// FileProveParam is an auto generated low-level Go binding around an user-defined struct.
type FileProveParam struct {
	RootHash []byte
	FileID   []byte
}

// NameInfo is an auto generated low-level Go binding around an user-defined struct.
type NameInfo struct {
	Type        uint64
	Header      []byte
	URL         []byte
	Name        []byte
	NameOwner   common.Address
	Desc        []byte
	BlockHeight *big.Int
	TTL         uint64
}

// PlotInfo is an auto generated low-level Go binding around an user-defined struct.
type PlotInfo struct {
	NumberID   uint64
	StartNonce uint64
	Nonces     uint64
}

// Setting is an auto generated low-level Go binding around an user-defined struct.
type Setting struct {
	GasPrice             uint64
	GasPerGBPerBlock     uint64
	GasPerKBPerBlock     uint64
	GasForChallenge      uint64
	MaxProveBlockNum     uint64
	MinVolume            uint64
	DefaultProvePeriod   uint64
	DefaultProveLevel    uint64
	DefaultCopyNum       uint64
	DefaultBlockInterval uint64
	MinSectorSize        uint64
}

// SpaceAddParams is an auto generated low-level Go binding around an user-defined struct.
type SpaceAddParams struct {
	OldUserSpace  UserSpace
	AddSize       uint64
	AddBlockCount uint64
	CurrentHeight *big.Int
	Setting       Setting
	FileList      [][]byte
}

// SpaceAddReturn is an auto generated low-level Go binding around an user-defined struct.
type SpaceAddReturn struct {
	NewUserSpace UserSpace
	UpdatedFiles []FileInfo
	Success      bool
}

// SpaceChangeReturn is an auto generated low-level Go binding around an user-defined struct.
type SpaceChangeReturn struct {
	NewUserSpace UserSpace
	State        TransferState
	UpdatedFiles []FileInfo
}

// SpaceProcessParams is an auto generated low-level Go binding around an user-defined struct.
type SpaceProcessParams struct {
	UserSpaceParams UserSpaceParams
	OldUserSpace    UserSpace
	Setting         Setting
}

// SpaceProcessReturn is an auto generated low-level Go binding around an user-defined struct.
type SpaceProcessReturn struct {
	NewUserSpace UserSpace
	State        TransferState
	UpdatedFiles []FileInfo
	Success      bool
}

// SpaceProcessRevokeParams is an auto generated low-level Go binding around an user-defined struct.
type SpaceProcessRevokeParams struct {
	UserSpaceParams UserSpaceParams
	OldUserspace    UserSpace
	Setting         Setting
	FileList        [][]byte
	Ops             uint64
}

// SpaceProcessRevokeReturn is an auto generated low-level Go binding around an user-defined struct.
type SpaceProcessRevokeReturn struct {
	UserSpace     UserSpace
	AddedAmount   uint64
	RevokedAmount uint64
	Update        []FileInfo
	Success       bool
}

// TransferState is an auto generated low-level Go binding around an user-defined struct.
type TransferState struct {
	From  common.Address
	To    common.Address
	Value uint64
}

// UserSpace is an auto generated low-level Go binding around an user-defined struct.
type UserSpace struct {
	Used         uint64
	Remain       uint64
	Balance      uint64
	ExpireHeight *big.Int
	UpdateHeight *big.Int
}

// UserSpaceOperation is an auto generated low-level Go binding around an user-defined struct.
type UserSpaceOperation struct {
	Type  uint8
	Value uint64
}

// UserSpaceParams is an auto generated low-level Go binding around an user-defined struct.
type UserSpaceParams struct {
	WalletAddr common.Address
	Owner      common.Address
	Size       UserSpaceOperation
	BlockCount UserSpaceOperation
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"port\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"}],\"name\":\"DNSNodeRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DNSNodeUnReg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"DnsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"FsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structNameInfo\",\"name\":\"newer\",\"type\":\"tuple\"}],\"name\":\"NotifyNameInfoAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"PDPVerifyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nodeAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"oldUserSpace\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"addSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"addBlockCount\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"currentHeight\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"fileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSpace.AddParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"AddUserSpace\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"newUserSpace\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"RootHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"FileID\",\"type\":\"bytes\"}],\"internalType\":\"structFileProveParam\",\"name\":\"FileProveParam_\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"updatedFiles\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structSpace.AddReturn\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteUserSpace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"Size\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"BlockCount\",\"type\":\"tuple\"}],\"internalType\":\"structUserSpaceParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCost\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structTransferState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUserSpace\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"Size\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"BlockCount\",\"type\":\"tuple\"}],\"internalType\":\"structUserSpaceParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"ManageUserSpace\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"_userSpace\",\"type\":\"tuple\"}],\"name\":\"UpdateUserSpace\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"}],\"name\":\"calcSingleValidFeeForFile\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"calcStorageFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"calcStorageFeeForOneNode\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"proveTime\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"}],\"name\":\"calcValidFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"proveTime\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"}],\"name\":\"calcValidFeeForOneNode\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"Size\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"BlockCount\",\"type\":\"tuple\"}],\"internalType\":\"structUserSpaceParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"getUserspaceChange\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"newUserSpace\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"RootHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"FileID\",\"type\":\"bytes\"}],\"internalType\":\"structFileProveParam\",\"name\":\"FileProveParam_\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"updatedFiles\",\"type\":\"tuple[]\"}],\"internalType\":\"structSpace.ChangeReturn\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractIFile\",\"name\":\"_fs\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"Size\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"BlockCount\",\"type\":\"tuple\"}],\"internalType\":\"structUserSpaceParams\",\"name\":\"userSpaceParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"oldUserspace\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"fileList\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"ops\",\"type\":\"uint64\"}],\"internalType\":\"structSpace.ProcessRevokeParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"processForUserSpaceOneAddOneRevoke\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"userSpace\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"addedAmount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revokedAmount\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"RootHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"FileID\",\"type\":\"bytes\"}],\"internalType\":\"structFileProveParam\",\"name\":\"FileProveParam_\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"update\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structSpace.ProcessRevokeReturn\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"Size\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"BlockCount\",\"type\":\"tuple\"}],\"internalType\":\"structUserSpaceParams\",\"name\":\"userSpaceParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"oldUserSpace\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"}],\"internalType\":\"structSpace.ProcessParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"processForUserSpaceOperations\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"newUserSpace\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"RootHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"FileID\",\"type\":\"bytes\"}],\"internalType\":\"structFileProveParam\",\"name\":\"FileProveParam_\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"updatedFiles\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structSpace.ProcessReturn\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506151d9806100206000396000f3fe6080604052600436106100e85760003560e01c806354e0d3c21161008a578063918a066511610059578063918a066514610307578063ca6142cb14610327578063ed108de914610347578063f1feacbf146103f857600080fd5b806354e0d3c2146101e55780636a4fe520146102a65780636c799c13146102c6578063722b25b9146102e657600080fd5b80633899831a116100c65780633899831a146101655780633984ff7d1461017857806344b2d82d146101a5578063485cc955146101c557600080fd5b80630f9fa2eb146100ed578063284c60031461010f57806337bafa5214610145575b600080fd5b3480156100f957600080fd5b5061010d610108366004613f0e565b610418565b005b34801561011b57600080fd5b5061012f61012a3660046141f6565b610583565b60405161013c9190614e5e565b60405180910390f35b34801561015157600080fd5b5061012f61016036600461415f565b6105b6565b61010d6101733660046142b3565b6105e1565b34801561018457600080fd5b506101986101933660046142b3565b61088f565b60405161013c9190614e30565b3480156101b157600080fd5b5061012f6101c0366004614119565b610a84565b3480156101d157600080fd5b5061010d6101e0366004614043565b610a9a565b3480156101f157600080fd5b50610299610200366004613f0e565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506001600160a01b0316600090815260026020818152604092839020835160a08101855281546001600160401b038082168352680100000000000000008204811694830194909452600160801b90049092169382019390935260018301546060820152910154608082015290565b60405161013c9190614e50565b6102b96102b4366004614073565b610b9f565b60405161013c9190614dc7565b6102d96102d43660046140a7565b610f68565b60405161013c9190614e0e565b6102f96102f43660046142b3565b611342565b60405161013c929190614dd8565b34801561031357600080fd5b5061012f61032236600461423c565b611637565b34801561033357600080fd5b5061012f6103423660046141c4565b61164f565b61010d610355366004613f2c565b6001600160a01b039091166000908152600260208181526040928390208451815492860151948601516001600160401b03908116600160801b027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff96821668010000000000000000026fffffffffffffffffffffffffffffffff199095169190921617929092179390931617825560608301516001830155608090920151910155565b61040b6104063660046140c6565b611678565b60405161013c9190614e1f565b6040805160a0808201835260008083526020808401829052838501829052606080850183905260809485018390526001600160a01b03871683526002808352928690208651948501875280546001600160401b03808216808852680100000000000000008304821695880195909552600160801b909104169685019690965260018601549084015293015491810191909152901580156104c55750600081604001516001600160401b0316115b1561050e5760408082015190516001600160a01b038416916001600160401b031680156108fc02916000818181858888f1935050505015801561050c573d6000803e3d6000fd5b505b60008160600151118015610526575043816060015111155b1561057f5760006105378284611899565b80519091501561057d577f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5816040516105709190614d56565b60405180910390a1505050565b505b5050565b6000620fa0008284866020015161059a9190614f55565b6105a49190614f55565b6105ae9190614f38565b949350505050565b60006105c3858584610a84565b6105ce846001614eed565b6105d89190614f55565b95945050505050565b6000806105ed83611342565b8051919350915015610628577f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5816040516105709190614d0f565b6020820151604001516001600160401b0316156106ee576020820151516001600160a01b03163014156106a8576020808301519081015160409182015191516001600160a01b03909116916001600160401b031680156108fc02916000818181858888f193505050501580156106a2573d6000803e3d6000fd5b506106ee565b8160200151604001516001600160401b03163410156106ee577f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f560405161057090614d33565b60005b8260400151518110156107a257600154604084015180516001600160a01b039092169163e0d208a591908490811061073957634e487b7160e01b600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b815260040161075d9190614dfd565b600060405180830381600087803b15801561077757600080fd5b505af115801561078b573d6000803e3d6000fd5b50505050808061079a90615088565b9150506106f1565b5081516020808501516001600160a01b0316600090815260028083526040918290208451815485870151858801516001600160401b03908116600160801b027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff92821668010000000000000000026fffffffffffffffffffffffffffffffff1990941691909416179190911716178155606080860151600183015560809095015191015585518187015180519084015194880151805194015192517f0c5dd947b790e17c40c62b29772f0e0ca54c86e473b5bf74cfac2a1f9156ee91956105709560039543959493614ca7565b6040805160608101825260008082526020820181905291810191909152604080516060810182526000808252602082018190529181019190915260008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561091957600080fd5b505afa15801561092d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095191906140fa565b9050600061095e85611dd3565b90508061096e5750909392505050565b6020808601516001600160a01b031660009081526002808352604091829020825160a08101845281546001600160401b038082168352680100000000000000008204811696830196909652600160801b9004909416928401929092526001820154606084018190529101546080830152158015906109f0575043816060015111155b15610a34576040805160a081018252600080825260208083018290529282018190526060808301829052608092830182905281855292840152439183018290528201525b6060810151158015610a495750438160600151145b15610a6e576000610a5a8488611e8e565b905080610a6c57509295945050505050565b505b610a79868285611f3d565b979650505050505050565b6000610a90848361164f565b6105ae9084614f55565b600054610100900460ff16610ab55760005460ff1615610ab9565b303b155b610af8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aef90614d66565b60405180910390fd5b600054610100900460ff16158015610b1a576000805461ffff19166101011790555b600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff16620100006001600160a01b038681169190910291909117909155600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016918416919091179055801561057d576000805461ff0019169055505050565b610ba7612f90565b610baf612f90565b825160600151610c595780516000905260208084015182516001600160401b0391821692019190915260408401516060850151610bf0929190911690614ed5565b81516060015280516000604090910181905281516080850151610c14919043612165565b9050806040015181602001518260000151610c2f9190614eed565b610c399190614eed565b82516001600160401b039091166040918201526001908301525092915050565b600083604001516001600160401b0316846000015160600151610c7c9190614ed5565b905060008460200151856000015160200151610c989190614eed565b85515184516001600160401b0391821690528451818316602090910152845160600184905286516040908101518651921691015285516080870151919250600091610ce4919043612165565b90506000610cfb8560000151886080015143612165565b90506000826040015183602001518460000151610d189190614eed565b610d229190614eed565b90506000826040015183602001518460000151610d3f9190614eed565b610d499190614eed565b9050816001600160401b0316816001600160401b031611610d7857505060006040860152509295945050505050565b6000610d848383614fac565b8a516000908190528b5160808d01519293509091610da3919043612165565b90506000816040015182602001518360000151610dc09190614eed565b610dca9190614eed565b9050806001600160401b03168c60000151604001516001600160401b031611610e0457505060006040890152509598975050505050505050565b8b5160400151600090610e18908390614fac565b9050836001600160401b0316816001600160401b03161115610e3d5760009350610e6c565b610e478185614fac565b9350838b60000151604001818151610e5f9190614eed565b6001600160401b03169052505b60408d01516001600160401b031615610f505760015460a08e015160808f01516040517f3f2cc9a00000000000000000000000000000000000000000000000000000000081526001600160a01b0390931692633f2cc9a092610ed49290918f90600401614c78565b600060405180830381600087803b158015610eee57600080fd5b505af1158015610f02573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610f2a9190810190613ffd565b151560408d0181905260208d0191909152610f505750989b9a5050505050505050505050565b5050600160408a015250969998505050505050505050565b610f70612fd9565b610f78612fd9565b600154835160200151604051639a2b5e6360e01b8152600092839283926001600160a01b0390921691639a2b5e6391610fb391600401614c18565b60006040518083038186803b158015610fcb57600080fd5b505afa158015610fdf573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526110079190810190613f66565b905060008061101988600001516122ee565b8051919350915015611038575050600060608501525091949350505050565b6001600160401b0382166011148061105957506001600160401b0382166010145b8061106d57506001600160401b0382166001145b1561110b5761107a613030565b6020808a0151825289516040908101518201516001600160401b03908116848401528b516060908101519093015116818401524391830191909152890151608082015260a0810184905260006110cf82610b9f565b8051604080820151602084015182850151151560608e01819052928d0152918b5290985090915061110857509598975050505050505050565b50505b6001600160401b0382166022148061112c57506001600160401b0382166002145b8061114057506001600160401b0382166020145b156111cd5761114d6130ed565b6020808a0151825289516040908101518201516001600160401b03908116848401528b5160609081015190930151168184015243918301919091528901516080820152600061119b826123d7565b805160208201516040830151151560608c01819052918b5297509091506111ca57509598975050505050505050565b50505b6001600160401b038216601214806111ee57506001600160401b0382166021145b1561127c576111fb6131a1565b885181526020808a0151908201526040808a015190820152606081018490526001600160401b0383166080820152600061123482611678565b8051602082015160408084015160608086015160808701511515918f01829052928e0192909252928c5290995090975090915061127957509598975050505050505050565b50505b8551606001516112925750939695505050505050565b8551436080909101526001600160401b03808516908616106112f157875151602080880180516001600160a01b039093169092529051309101526112d68486614fac565b60208701516001600160401b0390911660409091015261132f565b6020868101805130905289515190516001600160a01b039091169101526113188585614fac565b60208701516001600160401b039091166040909101525b5050600160608501525091949350505050565b61134a613243565b6060611354613243565b60008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156113a457600080fd5b505afa1580156113b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113dc91906140fa565b905060006113e986611dd3565b90508061143057505060408051808201909152601881527f696e76616c69642075736572737061636520706172616d73000000000000000060208201529094909350915050565b60006114d387602001516040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506001600160a01b0316600090815260026020818152604092839020835160a08101855281546001600160401b038082168352680100000000000000008204811694830194909452600160801b90049092169382019390935260018301546060820152910154608082015290565b80519091506000906001600160401b0316158015906114fe575060208201516001600160401b031615155b8015611516575060408201516001600160401b031615155b80156115255750608082015115155b80156115345750606082015115155b9050808015611547575043826060015111155b1561158b576040805160a081018252600080825260208083018290529282018190526060808301829052608092830182905281865292850152439184018290528301525b80158061159b5750438260600151145b156115de5760006115ac858a611e8e565b9050806115dc57856040518060600160405280602181526020016151836021913997509750505050505050915091565b505b6115e661329c565b8881526020810183905260408101859052600061160282610f68565b80518852602080820151818a0152604091820151828a0152815190810190915260008152969a96995095975050505050505050565b6000611644858484610583565b6105ce856001614eed565b6000620fa0008284606001516116659190614f55565b61166f9190614f38565b90505b92915050565b6040805161014081018252600060a0820181815260c0830182905260e083018290526101008301829052610120830182905282526020820181905291810182905260608082015260808101919091526040805161014081018252600060a0820181815260c0830182905260e0830182905261010083018290526101208301829052825260208201819052918101829052606080820152608081019190915260808301516000908190819081906001600160401b03166012141561174f57508551604081015160209081015160609092015101519093505b60808701516001600160401b03166021141561177f57865160408101516020908101516060909201510151935091505b611787613030565b60208089015182526001600160401b038087169183019190915284166040808301919091524360608084019190915290890151608083015288015160a082015260006117d282610b9f565b905080604001516117f157505060006080860152509295945050505050565b6117f96130ed565b815181526001600160401b03808616602083015284166040808301919091524360608301528a015160808201526000611831826123d7565b90508060400151611852575050600060808801525094979650505050505050565b8051895282516040908101516001600160401b039081166020808d0191909152925182015116908a0152909101516060880152505060016080860152509295945050505050565b606060008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156118eb57600080fd5b505afa1580156118ff573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061192391906140fa565b90504384606001518260c001516001600160401b03166119439190614ed5565b111561198457505060408051808201909152600b81527f6e6f7420657870697265640000000000000000000000000000000000000000006020820152611672565b600154604051639a2b5e6360e01b8152606091600091829182916001600160a01b0390911690639a2b5e63906119be908a90600401614c18565b60006040518083038186803b1580156119d657600080fd5b505afa1580156119ea573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611a129190810190613f66565b6040805160018082528183019092529192506000919060208083019080368337019050509050600081600081518110611a5b57634e487b7160e01b600052603260045260246000fd5b60200260200101906001811115611a8257634e487b7160e01b600052602160045260246000fd5b90816001811115611aa357634e487b7160e01b600052602160045260246000fd5b905250600154604051633ad525a960e01b81526001600160a01b0390911690633ad525a990611ada9085908c908690600401614c46565b600060405180830381600087803b158015611af457600080fd5b505af1158015611b08573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611b309190810190613f9a565b6001546040517f28a8565c00000000000000000000000000000000000000000000000000000000815293985091965094506000916001600160a01b03909116906328a8565c90611b84908c90600401614c18565b60006040518083038186803b158015611b9c57600080fd5b505afa158015611bb0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611bd89190810190613f66565b600154604051633ad525a960e01b81529192506001600160a01b031690633ad525a990611c0d9084908d908790600401614c46565b600060405180830381600087803b158015611c2757600080fd5b505af1158015611c3b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611c639190810190613f9a565b9197509550935060005b8651811015611d385760005b8251811015611d2557828181518110611ca257634e487b7160e01b600052603260045260246000fd5b602002602001015180519060200120888381518110611cd157634e487b7160e01b600052603260045260246000fd5b6020026020010151805190602001201415611d1357828181518110611d0657634e487b7160e01b600052603260045260246000fd5b6020026020010160608152505b80611d1d81615088565b915050611c79565b5080611d3081615088565b915050611c6d565b506001546040517fd70e62720000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063d70e627290611d84908c908590600401614c26565b600060405180830381600087803b158015611d9e57600080fd5b505af1158015611db2573d6000803e3d6000fd5b50506040805160208101909152600081529c9b505050505050505050505050565b6000808260400151602001516001600160401b031611611df557506000919050565b60008260600151602001516001600160401b031611611e1657506000919050565b600080611e22846122ee565b8051919350915015611e38575060009392505050565b6001600160401b038216611e50575060009392505050565b6000611e5b85612660565b90508015611e83576000611e6e866126be565b905080611e815750600095945050505050565b505b506001949350505050565b6000806000611e9c846122ee565b8051919350915015611eb357600092505050611672565b6001600160401b038216601114611ecf57600092505050611672565b6040840151602001516001600160401b03161580611efc57506060840151602001516001600160401b0316155b15611f0c57600092505050611672565b8460c001516001600160401b03168460600151602001516001600160401b03161015611e8357600092505050611672565b6040805160a080820183526000808352602080840182905283850182905260608085018390526080808601849052865180830188528481528084018590528088018590528751958601885284865285840185905285880185905285830185905290850184905286519182018752838252918101839052948501829052929390808080611fc88b6122ee565b8051919350915015611fe4575093955091935061215d92505050565b6001600160401b0382166011148061200557506001600160401b0382166010145b8061201957506001600160401b0382166001145b1561205a5761203b8a8c60400151602001518d6060015160200151438d612790565b805192985090955091501561205a575093955091935061215d92505050565b6001600160401b0382166022148061207b57506001600160401b0382166002145b8061208f57506001600160401b0382166020145b156120a4575093955091935061215d92505050565b6001600160401b038216601214806120c557506001600160401b0382166021145b156120da575093955091935061215d92505050565b826001600160401b0316846001600160401b03161115612125576120fe8385614fac565b6001600160401b031660408601528a516001600160a01b0316855260006020860152612152565b61212f8484614fac565b6001600160401b03166040860152600085528a516001600160a01b031660208601525b509395509193505050505b935093915050565b60408051606081018252600080825260208201819052918101919091526121fe604080516101e08101825260608082526000602083018190529282018390528082018390526080820183905260a0820183905260c0820183905260e0820183905261010082018190526101208201839052610140820183905261016082018190526101808201526101a08101829052906101c082015290565b6020850151855161220f9190614eed565b6001600160401b03908116602083015260c0808601518216604080850191909152606088015160808501526101008701519092169083015260015490517f178e4dc00000000000000000000000000000000000000000000000000000000081526000916001600160a01b03169063178e4dc09061229490859089908990600401614e3e565b60606040518083038186803b1580156122ac57600080fd5b505afa1580156122c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122e49190614295565b9695505050505050565b600060606000612301846040015161285d565b90508061234a5760006040518060400160405280601681526020017f496e76616c69642073697a65206f7065726174696f6e000000000000000000008152509250925050915091565b6000612359856060015161285d565b9050806123a35760006040518060400160405280601d81526020017f496e76616c696420626c6f636b20636f756e74206f7065726174696f6e000000815250935093505050915091565b6040850151516060860151516000916123bb91612939565b6040805160208101909152600081529097909650945050505050565b60408051610100810182526000606082018181526080830182905260a0830182905260c0830182905260e083018290528252602082018190529181019190915260408051610100810182526000606082018181526080830182905260a0830182905260c0830182905260e083018290528252602082018190529181019190915282602001516001600160401b03168360000151602001516001600160401b0316101561248a576000604082015292915050565b826060015183604001516001600160401b03168460000151606001516124b09190614f91565b10156124c3576000604082015292915050565b6040805160a081018252600080825260208083018290529282018190526060820181905260808201528451516001600160401b0316815284820151855190920151909161250f91614fac565b6001600160401b0390811660208301526040850151855160600151612538929190911690614f91565b606082015283516040908101516001600160401b031690820152835160808501516000916125669143612165565b9050600061257983876080015143612165565b905060008260400151836020015184600001516125969190614eed565b6125a09190614eed565b905060008260400151836020015184600001516125bd9190614eed565b6125c79190614eed565b9050806001600160401b0316826001600160401b0316116125f5575050600060408501525091949350505050565b60006126018284614fac565b9050806001600160401b031686604001516001600160401b0316101561263557505060006040860152509295945050505050565b80866040018181516126479190614fac565b6001600160401b03169052509598975050505050505050565b60006002604083015151600281111561268957634e487b7160e01b600052602160045260246000fd5b14806116725750600260608301515160028111156126b757634e487b7160e01b600052602160045260246000fd5b1492915050565b6001546020820151604051639a2b5e6360e01b815260009283926001600160a01b0390911691639a2b5e63916126f691600401614c18565b60006040518083038186803b15801561270e57600080fd5b505afa158015612722573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261274a9190810190613f66565b80519091501561275d5750600092915050565b82602001516001600160a01b031683600001516001600160a01b0316146127875750600092915050565b50600192915050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040805160a081018252600080825260208201819052918101829052606081810183905260808201839052906000896060015160001415612827576128038a8a8a898b612992565b50604080820151815160208101909252600082529196509094509250612852915050565b6128348a8a8a898b612992565b60408051602081019091526000815291965094509250612852915050565b955095509592505050565b600060028251600281111561288257634e487b7160e01b600052602160045260246000fd5b14156128aa5760208201516001600160401b03166128a257506000919050565b506001919050565b6001825160028111156128cd57634e487b7160e01b600052602160045260246000fd5b14156128ed5760208201516001600160401b03166128a257506000919050565b60008251600281111561291057634e487b7160e01b600052602160045260246000fd5b14156129315760208201516001600160401b0316156128a257506000919050565b506000919050565b60008082600281111561295c57634e487b7160e01b600052602160045260246000fd5b600485600281111561297e57634e487b7160e01b600052602160045260246000fd5b6001600160401b0316901b17949350505050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915260008060006129d760008a888961010001518c8c8b612a66565b90508060400151816020015182600001516129f29190614eed565b6129fc9190614eed565b91508789602001818151612a109190614eed565b6001600160401b0390811690915260608b018051918a16925090612a35908390614ed5565b905250604089018051839190612a4c908390614eed565b6001600160401b0316905250979890975095505050505050565b6040805160608082018352600080835260208084018290528385018290528451808401865282815280820183905280860183905285516101e0810187528481529182018390529481018290528083018290526080810182905260a0810182905260c0810182905260e0810182905261010081018390526101208101829052610140810182905261016081018390526101808101929092526101a082018190526101c08201529091889160208301518351612b209190614eed565b6001600160401b03908116602083015260c0808b015182166040840152606085015160808401526101008b0151909116908201526000808c6002811115612b7757634e487b7160e01b600052602160045260246000fd5b1415612db1576000876001600160401b031611612c0d576000886001600160401b031611612ba457600197505b612bbd82878660600151612bb89190614f91565b612e37565b9050612bcb8a828b8b6105b6565b6001600160401b031660408401526060840151612bf3908b908b908b90610322908b90614f91565b6001600160401b0316602084015250909250610a79915050565b6000886001600160401b031611612c98576000886001600160401b031611612c3457600197505b612c4782886001600160401b0316612e37565b9050612c638a828b876000015188602001516101609190614eed565b6001600160401b0316604084015283516020850151612bf3918c918c91612c8991614eed565b8a6001600160401b0316611637565b6000886001600160401b031611612cae57600197505b60008685606001511115612ccf57868560600151612ccc9190614f91565b90505b612cd98382612e37565b91506000612ce98c848d8d6105b6565b90506000612cf98d8d8d86611637565b9050612d0e858b6001600160401b0316612e37565b93506000612d378e868f8f8c602001518d60000151612d2d9190614eed565b6101609190614eed565b90506000612d6e8f8f8f8c602001518d60000151612d559190614eed565b612d5f9190614eed565b8f6001600160401b0316611637565b9050612d7a8285614eed565b6001600160401b03166040890152612d928184614eed565b6001600160401b0316602089015250959750610a799650505050505050565b60028c6002811115612dd357634e487b7160e01b600052602160045260246000fd5b1415612df357612df08285602001516001600160401b0316612e37565b90505b60018c6002811115612e1557634e487b7160e01b600052602160045260246000fd5b1415612e275782945050505050610a79565b50909a9950505050505050505050565b604082015160009081906001600160401b0316612eac5760608401516001600160401b0316612e6857600060608501525b612e9d84606001516001600160401b03166002811115612e9857634e487b7160e01b600052602160045260246000fd5b612ecf565b6001600160401b031660408501525b6040840151612ec4906001600160401b031684614f20565b6105ae906001614ed5565b60006143808082612ee1826002614f55565b90506000612ef0846008614f55565b90506002866002811115612f1457634e487b7160e01b600052602160045260246000fd5b1415612f235795945050505050565b6001866002811115612f4557634e487b7160e01b600052602160045260246000fd5b1415612f545750949350505050565b6000866002811115612f7657634e487b7160e01b600052602160045260246000fd5b1415612f86575090949350505050565b5090949350505050565b6040805161010081019091526000606082018181526080830182905260a0830182905260c0830182905260e083019190915281905b815260606020820152600060409091015290565b604080516101208101825260006080820181815260a0830182905260c0830182905260e083018290526101008301829052825282516060810184528181526020808201839052938101919091529091820190612fc5565b604080516101608101909152600060c0820181815260e0830182905261010083018290526101208301829052610140830191909152819081526000602082018190526040820181905260608201526080016130e06040805161016081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101829052610120810182905261014081019190915290565b8152602001606081525090565b604080516101408101909152600060a0820181815260c0830182905260e0830182905261010083018290526101208301919091528190815260006020820181905260408201819052606082015260800161319c6040805161016081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101829052610120810182905261014081019190915290565b905290565b6040518060a001604052806131b461333c565b81526040805160a081018252600080825260208281018290529282018190526060820181905260808201529101908152604080516101608101825260008082526020828101829052928201819052606082018190526080820181905260a0820181905260c0820181905260e0820181905261010082018190526101208201819052610140820152910190612fc5565b6040805161010081019091526000606082018181526080830182905260a0830182905260c0830182905260e0830191909152819081526040805160608101825260008082526020828101829052928201529101906130e0565b60405180606001604052806132af61333c565b81526040805160a081018252600080825260208281018290529282018190526060820181905260808201529101908152604080516101608101825260008082526020828101829052928201819052606082018190526080820181905260a0820181905260c0820181905260e082018190526101008201819052610120820181905261014082015291015290565b604080516080810182526000808252602082015290810161336d604080518082019091526000808252602082015290565b815260200161319c604080518082019091526000808252602082015290565b600061339f61339a84614e88565b614e6c565b905080838252602082019050828560208602820111156133be57600080fd5b60005b858110156133ea57816133d488826135bf565b84525060209283019291909101906001016133c1565b5050509392505050565b600061340261339a84614e88565b9050808382526020820190508285602086028201111561342157600080fd5b60005b858110156133ea5781356001600160401b0381111561344257600080fd5b80860161344f8982613659565b855250506020928301929190910190600101613424565b600061347461339a84614e88565b9050808382526020820190508285602086028201111561349357600080fd5b60005b858110156133ea5781516001600160401b038111156134b457600080fd5b8086016134c1898261367a565b855250506020928301929190910190600101613496565b60006134e661339a84614e88565b9050808382526020820190508285602086028201111561350557600080fd5b60005b858110156133ea5781516001600160401b0381111561352657600080fd5b8086016135338982613778565b855250506020928301929190910190600101613508565b600061355861339a84614eab565b90508281526020810184848401111561357057600080fd5b61357b848285615024565b509392505050565b600061359161339a84614eab565b9050828152602081018484840111156135a957600080fd5b61357b848285615030565b80356116728161512e565b80516116728161512e565b600082601f8301126135db57600080fd5b81516105ae84826020860161338c565b600082601f8301126135fc57600080fd5b81356105ae8482602086016133f4565b600082601f83011261361d57600080fd5b81516105ae848260208601613466565b600082601f83011261363e57600080fd5b81516105ae8482602086016134d8565b805161167281615142565b600082601f83011261366a57600080fd5b81356105ae84826020860161354a565b600082601f83011261368b57600080fd5b81516105ae848260208601613583565b80356116728161514a565b805161167281615153565b805161167281615160565b803561167281615153565b600061028082840312156136da57600080fd5b6136e460c0614e6c565b905060006136f28484613e6f565b82525060a061370384848301613ef8565b60208301525060c061371784828501613ef8565b60408301525060e061372b84828501613ee2565b60608301525061010061374084828501613bcc565b6080830152506102608201356001600160401b0381111561376057600080fd5b61376c848285016135eb565b60a08301525092915050565b6000610320828403121561378b57600080fd5b6137966102e0614e6c565b82519091506001600160401b038111156137af57600080fd5b6137bb8482850161367a565b82525060206137cc848483016135bf565b60208301525060408201516001600160401b038111156137eb57600080fd5b6137f78482850161367a565b604083015250606061380b84828501613f03565b606083015250608061381f84828501613f03565b60808301525060a061383384828501613f03565b60a08301525060c061384784828501613f03565b60c08301525060e061385b84828501613f03565b60e08301525061010061387084828501613eed565b6101008301525061012061388684828501613f03565b6101208301525061014061389c84828501613f03565b610140830152506101608201516001600160401b038111156138bd57600080fd5b6138c984828501613a0d565b610160830152506101806138df84828501613f03565b610180830152506101a06138f584828501613eed565b6101a0830152506101c061390b8482850161364e565b6101c0830152506101e0613921848285016136b1565b6101e08301525061020061393784828501613f03565b610200830152506102208201516001600160401b0381111561395857600080fd5b613964848285016135ca565b610220830152506102408201516001600160401b0381111561398557600080fd5b613991848285016135ca565b610240830152506102608201516001600160401b038111156139b257600080fd5b6139be8482850161367a565b610260830152506102806139d4848285016136a6565b610280830152506102a06139ea8482850161364e565b6102a0830152506102c0613a0084828501613a82565b6102c08301525092915050565b600060408284031215613a1f57600080fd5b613a296040614e6c565b82519091506001600160401b03811115613a4257600080fd5b613a4e8482850161367a565b82525060208201516001600160401b03811115613a6a57600080fd5b613a768482850161367a565b60208301525092915050565b600060608284031215613a9457600080fd5b613a9e6060614e6c565b90506000613aac8484613f03565b8252506020613abd84848301613f03565b6020830152506040613ad184828501613f03565b60408301525092915050565b60006102c08284031215613af057600080fd5b613afa6060614e6c565b90506000613b088484613e00565b82525060c0613b1984848301613e6f565b602083015250610160613ad184828501613bcc565b60006103008284031215613b4157600080fd5b613b4b60a0614e6c565b90506000613b598484613e00565b82525060c0613b6a84848301613e6f565b602083015250610160613b7f84828501613bcc565b6040830152506102c08201356001600160401b03811115613b9f57600080fd5b613bab848285016135eb565b6060830152506102e0613bc084828501613ef8565b60808301525092915050565b60006101608284031215613bdf57600080fd5b613bea610160614e6c565b90506000613bf88484613ef8565b8252506020613c0984848301613ef8565b6020830152506040613c1d84828501613ef8565b6040830152506060613c3184828501613ef8565b6060830152506080613c4584828501613ef8565b60808301525060a0613c5984828501613ef8565b60a08301525060c0613c6d84828501613ef8565b60c08301525060e0613c8184828501613ef8565b60e083015250610100613c9684828501613ef8565b61010083015250610120613cac84828501613ef8565b61012083015250610140613cc284828501613ef8565b6101408301525092915050565b60006101608284031215613ce257600080fd5b613ced610160614e6c565b90506000613cfb8484613f03565b8252506020613d0c84848301613f03565b6020830152506040613d2084828501613f03565b6040830152506060613d3484828501613f03565b6060830152506080613d4884828501613f03565b60808301525060a0613d5c84828501613f03565b60a08301525060c0613d7084828501613f03565b60c08301525060e0613d8484828501613f03565b60e083015250610100613d9984828501613f03565b61010083015250610120613daf84828501613f03565b61012083015250610140613cc284828501613f03565b600060408284031215613dd757600080fd5b613de16040614e6c565b90506000613def84846136bc565b8252506020613a7684848301613ef8565b600060c08284031215613e1257600080fd5b613e1c6080614e6c565b90506000613e2a84846135b4565b8252506020613e3b848483016135b4565b6020830152506040613e4f84828501613dc5565b6040830152506080613e6384828501613dc5565b60608301525092915050565b600060a08284031215613e8157600080fd5b613e8b60a0614e6c565b90506000613e998484613ef8565b8252506020613eaa84848301613ef8565b6020830152506040613ebe84828501613ef8565b6040830152506060613ed284828501613ee2565b6060830152506080613bc0848285015b80356116728161516d565b80516116728161516d565b803561167281615173565b805161167281615173565b600060208284031215613f2057600080fd5b60006105ae84846135b4565b60008060c08385031215613f3f57600080fd5b6000613f4b85856135b4565b9250506020613f5c85828601613e6f565b9150509250929050565b600060208284031215613f7857600080fd5b81516001600160401b03811115613f8e57600080fd5b6105ae8482850161360c565b600080600060608486031215613faf57600080fd5b83516001600160401b03811115613fc557600080fd5b613fd18682870161360c565b9350506020613fe286828701613f03565b9250506040613ff38682870161364e565b9150509250925092565b6000806040838503121561401057600080fd5b82516001600160401b0381111561402657600080fd5b6140328582860161362d565b9250506020613f5c8582860161364e565b6000806040838503121561405657600080fd5b6000614062858561369b565b9250506020613f5c8582860161369b565b60006020828403121561408557600080fd5b81356001600160401b0381111561409b57600080fd5b6105ae848285016136c7565b60006102c082840312156140ba57600080fd5b60006105ae8484613add565b6000602082840312156140d857600080fd5b81356001600160401b038111156140ee57600080fd5b6105ae84828501613b2e565b6000610160828403121561410d57600080fd5b60006105ae8484613ccf565b60008060006101a0848603121561412f57600080fd5b600061413b8686613bcc565b93505061016061414d86828701613ee2565b925050610180613ff386828701613ef8565b6000806000806101c0858703121561417657600080fd5b60006141828787613bcc565b94505061016061419487828801613ee2565b9350506101806141a687828801613ef8565b9250506101a06141b887828801613ef8565b91505092959194509250565b60008061018083850312156141d857600080fd5b60006141e48585613bcc565b925050610160613f5c85828601613ef8565b60008060006101a0848603121561420c57600080fd5b60006142188686613bcc565b93505061016061422a86828701613ef8565b925050610180613ff386828701613ee2565b6000806000806101c0858703121561425357600080fd5b600061425f8787613bcc565b94505061016061427187828801613ef8565b93505061018061428387828801613ef8565b9250506101a06141b887828801613ee2565b6000606082840312156142a757600080fd5b60006105ae8484613a82565b600060c082840312156142c557600080fd5b60006105ae8484613e00565b60006142dd838361431d565b505060200190565b600061166f83836144b8565b60006142dd83836144fc565b600061166f8383614626565b60006143158383614bf2565b505060600190565b61432681614fc9565b82525050565b6000614336825190565b80845260209384019383018060005b8381101561436a57815161435988826142d1565b975060208301925050600101614345565b509495945050505050565b600061437f825190565b808452602084019350836020820285016143998560200190565b8060005b858110156143ce57848403895281516143b685826142e5565b94506020830160209a909a019992505060010161439d565b5091979650505050505050565b60006143e5825190565b80845260209384019383018060005b8381101561436a57815161440888826142f1565b9750602083019250506001016143f4565b6000614423825190565b8084526020840193508360208202850161443d8560200190565b8060005b858110156143ce578484038952815161445a85826142fd565b94506020830160209a909a0199925050600101614441565b600061447c825190565b80845260209384019383018060005b8381101561436a57815161449f8882614309565b97506020830192505060010161448b565b801515614326565b60006144c2825190565b8084526020840193506144d9818560208601615030565b601f01601f19169290920192915050565b61432681615003565b6143268161500e565b61432681615019565b600f81526000602082017f4d616e6167655573657253706163650000000000000000000000000000000000815291505b5060200190565b601281526000602082017f496e73756666696369656e742066756e6473000000000000000000000000000081529150614535565b600f81526000602082017f44656c657465557365725370616365000000000000000000000000000000000081529150614535565b805160009060e08401906145b88582614b95565b50602083015184820360a08601526145d08282614419565b915050604083015161357b60c08601826144b0565b80516000906101208401906145fa8582614b95565b50602083015161460d60a0860182614a29565b5060408301518482036101008601526105d88282614419565b80516103208084526000919084019061463f82826144b8565b9150506020830151614654602086018261431d565b506040830151848203604086015261466c82826144b8565b91505060608301516146816060860182614c09565b5060808301516146946080860182614c09565b5060a08301516146a760a0860182614c09565b5060c08301516146ba60c0860182614c09565b5060e08301516146cd60e0860182614c09565b506101008301516146e2610100860182614c03565b506101208301516146f7610120860182614c09565b5061014083015161470c610140860182614c09565b506101608301518482036101608601526147268282614824565b91505061018083015161473d610180860182614c09565b506101a08301516147526101a0860182614c03565b506101c08301516147676101c08601826144b0565b506101e083015161477c6101e08601826144fc565b50610200830151614791610200860182614c09565b506102208301518482036102208601526147ab828261432c565b9150506102408301518482036102408601526147c7828261432c565b9150506102608301518482036102608601526147e382826144b8565b9150506102808301516147fa6102808601826144f3565b506102a083015161480f6102a08601826144b0565b506102c083015161357b6102c0860182614856565b805160408084526000919084019061483c82826144b8565b915050602083015184820360208601526105d882826144b8565b805160608301906148678482614c09565b50602082015161487a6020850182614c09565b50604082015161488d6040850182614c09565b50505050565b80516000906101408401906148a88582614b95565b5060208301516148bb60a0860182614a29565b5060408301518482036101008601526148d48282614419565b915050606083015161357b6101208601826144b0565b80516000906101208401906148ff8582614b95565b50602083015161491260a0860182614c09565b50604083015161492560c0860182614c09565b50606083015184820360e086015261493d8282614419565b915050608083015161357b6101008601826144b0565b80516101608301906149658482614c09565b5060208201516149786020850182614c09565b50604082015161498b6040850182614c09565b50606082015161499e6060850182614c09565b5060808201516149b16080850182614c09565b5060a08201516149c460a0850182614c09565b5060c08201516149d760c0850182614c09565b5060e08201516149ea60e0850182614c09565b506101008201516149ff610100850182614c09565b50610120820151614a14610120850182614c09565b5061014082015161488d610140850182614c09565b80516060830190614a3a848261431d565b50602082015161487a602085018261431d565b80516101e080845260009190840190614a6682826144b8565b9150506020830151614a7b6020860182614c09565b506040830151614a8e6040860182614c09565b506060830151614aa16060860182614c09565b506080830151614ab46080860182614c03565b5060a0830151614ac760a0860182614c09565b5060c0830151614ada60c0860182614c09565b5060e0830151614aed60e08601826144b0565b50610100830151848203610100860152614b0782826144b8565b915050610120830151614b1e6101208601826144b0565b50610140830151614b336101408601826144b0565b50610160830151848203610160860152614b4d82826144b8565b915050610180830151848203610180860152614b698282614472565b9150506101a0830151614b806101a08601826144b0565b506101c083015161357b6101c08601826144fc565b805160a0830190614ba68482614c09565b506020820151614bb96020850182614c09565b506040820151614bcc6040850182614c09565b506060820151614bdf6060850182614c03565b50608082015161488d6080850182614c03565b80516060830190614867848261431d565b80614326565b6001600160401b038116614326565b60208101611672828461431d565b60408101614c34828561431d565b81810360208301526105ae8184614375565b60608082528101614c578186614375565b9050614c66602083018561431d565b81810360408301526105d881846143db565b6101a08082528101614c8a8186614375565b9050614c996020830185614953565b6105ae610180830184614c03565b60e08101614cb5828a6144ea565b614cc26020830189614c03565b614ccf604083018861431d565b614cdc60608301876144f3565b614ce96080830186614c09565b614cf660a08301856144f3565b614d0360c0830184614c09565b98975050505050505050565b60408082528101614d1f81614505565b9050818103602083015261166f81846144b8565b60408082528101614d4381614505565b905081810360208301526116728161453c565b60408082528101614d1f81614570565b6020808252810161167281602e81527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160208201527f647920696e697469616c697a6564000000000000000000000000000000000000604082015260600190565b6020808252810161166f81846145a4565b60408082528101614de981856145e5565b905081810360208301526105ae81846144b8565b6020808252810161166f8184614626565b6020808252810161166f8184614893565b6020808252810161166f81846148ea565b606081016116728284614a29565b6101a08082528101614c8a8186614a4d565b60a081016116728284614b95565b602081016116728284614c09565b6000614e7760405190565b9050614e83828261505c565b919050565b60006001600160401b03821115614ea157614ea16150e5565b5060209081020190565b60006001600160401b03821115614ec457614ec46150e5565b601f19601f83011660200192915050565b60008219821115614ee857614ee86150a3565b500190565b60006001600160401b03821691506001600160401b0383169250826001600160401b0303821115614ee857614ee86150a3565b6000825b925082614f3357614f336150b9565b500490565b60006001600160401b03821691506001600160401b038316614f24565b60006001600160401b03821691506001600160401b0383169250816001600160401b030483118215151615614f8c57614f8c6150a3565b500290565b6000825b925082821015614fa757614fa76150a3565b500390565b60006001600160401b03821691506001600160401b038316614f95565b60006001600160a01b038216611672565b600061167282614fc9565b80614e83816150fb565b80614e838161510e565b80614e838161511e565b600061167282614fe5565b600061167282614fef565b600061167282614ff9565b82818337506000910152565b60005b8381101561504b578181015183820152602001615033565b8381111561488d5750506000910152565b601f19601f83011681018181106001600160401b0382111715615081576150816150e5565b6040525050565b600060001982141561509c5761509c6150a3565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600b811061510b5761510b6150cf565b50565b6003811061510b5761510b6150cf565b6002811061510b5761510b6150cf565b61513781614fc9565b811461510b57600080fd5b801515615137565b61513781614fda565b6003811061510b57600080fd5b6002811061510b57600080fd5b80615137565b6001600160401b03811661513756fe696e76616c696420666972737420757365727370616365206f7065726174696f6ea26469706673582212205ea007808b132f36aea13becfc2db4b25d1c04926776db23c68885627f6f91cb64736f6c63430008040033",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// GetUpdateCost is a free data retrieval call binding the contract method 0x3984ff7d.
//
// Solidity: function GetUpdateCost((address,address,(uint8,uint64),(uint8,uint64)) params) view returns((address,address,uint64))
func (_Store *StoreCaller) GetUpdateCost(opts *bind.CallOpts, params UserSpaceParams) (TransferState, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetUpdateCost", params)

	if err != nil {
		return *new(TransferState), err
	}

	out0 := *abi.ConvertType(out[0], new(TransferState)).(*TransferState)

	return out0, err

}

// GetUpdateCost is a free data retrieval call binding the contract method 0x3984ff7d.
//
// Solidity: function GetUpdateCost((address,address,(uint8,uint64),(uint8,uint64)) params) view returns((address,address,uint64))
func (_Store *StoreSession) GetUpdateCost(params UserSpaceParams) (TransferState, error) {
	return _Store.Contract.GetUpdateCost(&_Store.CallOpts, params)
}

// GetUpdateCost is a free data retrieval call binding the contract method 0x3984ff7d.
//
// Solidity: function GetUpdateCost((address,address,(uint8,uint64),(uint8,uint64)) params) view returns((address,address,uint64))
func (_Store *StoreCallerSession) GetUpdateCost(params UserSpaceParams) (TransferState, error) {
	return _Store.Contract.GetUpdateCost(&_Store.CallOpts, params)
}

// GetUserSpace is a free data retrieval call binding the contract method 0x54e0d3c2.
//
// Solidity: function GetUserSpace(address walletAddr) view returns((uint64,uint64,uint64,uint256,uint256))
func (_Store *StoreCaller) GetUserSpace(opts *bind.CallOpts, walletAddr common.Address) (UserSpace, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetUserSpace", walletAddr)

	if err != nil {
		return *new(UserSpace), err
	}

	out0 := *abi.ConvertType(out[0], new(UserSpace)).(*UserSpace)

	return out0, err

}

// GetUserSpace is a free data retrieval call binding the contract method 0x54e0d3c2.
//
// Solidity: function GetUserSpace(address walletAddr) view returns((uint64,uint64,uint64,uint256,uint256))
func (_Store *StoreSession) GetUserSpace(walletAddr common.Address) (UserSpace, error) {
	return _Store.Contract.GetUserSpace(&_Store.CallOpts, walletAddr)
}

// GetUserSpace is a free data retrieval call binding the contract method 0x54e0d3c2.
//
// Solidity: function GetUserSpace(address walletAddr) view returns((uint64,uint64,uint64,uint256,uint256))
func (_Store *StoreCallerSession) GetUserSpace(walletAddr common.Address) (UserSpace, error) {
	return _Store.Contract.GetUserSpace(&_Store.CallOpts, walletAddr)
}

// CalcSingleValidFeeForFile is a free data retrieval call binding the contract method 0xca6142cb.
//
// Solidity: function calcSingleValidFeeForFile((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 fileSize) pure returns(uint64)
func (_Store *StoreCaller) CalcSingleValidFeeForFile(opts *bind.CallOpts, setting Setting, fileSize uint64) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "calcSingleValidFeeForFile", setting, fileSize)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalcSingleValidFeeForFile is a free data retrieval call binding the contract method 0xca6142cb.
//
// Solidity: function calcSingleValidFeeForFile((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 fileSize) pure returns(uint64)
func (_Store *StoreSession) CalcSingleValidFeeForFile(setting Setting, fileSize uint64) (uint64, error) {
	return _Store.Contract.CalcSingleValidFeeForFile(&_Store.CallOpts, setting, fileSize)
}

// CalcSingleValidFeeForFile is a free data retrieval call binding the contract method 0xca6142cb.
//
// Solidity: function calcSingleValidFeeForFile((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 fileSize) pure returns(uint64)
func (_Store *StoreCallerSession) CalcSingleValidFeeForFile(setting Setting, fileSize uint64) (uint64, error) {
	return _Store.Contract.CalcSingleValidFeeForFile(&_Store.CallOpts, setting, fileSize)
}

// CalcStorageFee is a free data retrieval call binding the contract method 0x918a0665.
//
// Solidity: function calcStorageFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 copyNum, uint64 fileSize, uint256 duration) pure returns(uint64)
func (_Store *StoreCaller) CalcStorageFee(opts *bind.CallOpts, setting Setting, copyNum uint64, fileSize uint64, duration *big.Int) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "calcStorageFee", setting, copyNum, fileSize, duration)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalcStorageFee is a free data retrieval call binding the contract method 0x918a0665.
//
// Solidity: function calcStorageFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 copyNum, uint64 fileSize, uint256 duration) pure returns(uint64)
func (_Store *StoreSession) CalcStorageFee(setting Setting, copyNum uint64, fileSize uint64, duration *big.Int) (uint64, error) {
	return _Store.Contract.CalcStorageFee(&_Store.CallOpts, setting, copyNum, fileSize, duration)
}

// CalcStorageFee is a free data retrieval call binding the contract method 0x918a0665.
//
// Solidity: function calcStorageFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 copyNum, uint64 fileSize, uint256 duration) pure returns(uint64)
func (_Store *StoreCallerSession) CalcStorageFee(setting Setting, copyNum uint64, fileSize uint64, duration *big.Int) (uint64, error) {
	return _Store.Contract.CalcStorageFee(&_Store.CallOpts, setting, copyNum, fileSize, duration)
}

// CalcStorageFeeForOneNode is a free data retrieval call binding the contract method 0x284c6003.
//
// Solidity: function calcStorageFeeForOneNode((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 fileSize, uint256 duration) pure returns(uint64)
func (_Store *StoreCaller) CalcStorageFeeForOneNode(opts *bind.CallOpts, setting Setting, fileSize uint64, duration *big.Int) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "calcStorageFeeForOneNode", setting, fileSize, duration)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalcStorageFeeForOneNode is a free data retrieval call binding the contract method 0x284c6003.
//
// Solidity: function calcStorageFeeForOneNode((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 fileSize, uint256 duration) pure returns(uint64)
func (_Store *StoreSession) CalcStorageFeeForOneNode(setting Setting, fileSize uint64, duration *big.Int) (uint64, error) {
	return _Store.Contract.CalcStorageFeeForOneNode(&_Store.CallOpts, setting, fileSize, duration)
}

// CalcStorageFeeForOneNode is a free data retrieval call binding the contract method 0x284c6003.
//
// Solidity: function calcStorageFeeForOneNode((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 fileSize, uint256 duration) pure returns(uint64)
func (_Store *StoreCallerSession) CalcStorageFeeForOneNode(setting Setting, fileSize uint64, duration *big.Int) (uint64, error) {
	return _Store.Contract.CalcStorageFeeForOneNode(&_Store.CallOpts, setting, fileSize, duration)
}

// CalcValidFee is a free data retrieval call binding the contract method 0x37bafa52.
//
// Solidity: function calcValidFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 proveTime, uint64 copyNum, uint64 fileSize) pure returns(uint64)
func (_Store *StoreCaller) CalcValidFee(opts *bind.CallOpts, setting Setting, proveTime *big.Int, copyNum uint64, fileSize uint64) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "calcValidFee", setting, proveTime, copyNum, fileSize)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalcValidFee is a free data retrieval call binding the contract method 0x37bafa52.
//
// Solidity: function calcValidFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 proveTime, uint64 copyNum, uint64 fileSize) pure returns(uint64)
func (_Store *StoreSession) CalcValidFee(setting Setting, proveTime *big.Int, copyNum uint64, fileSize uint64) (uint64, error) {
	return _Store.Contract.CalcValidFee(&_Store.CallOpts, setting, proveTime, copyNum, fileSize)
}

// CalcValidFee is a free data retrieval call binding the contract method 0x37bafa52.
//
// Solidity: function calcValidFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 proveTime, uint64 copyNum, uint64 fileSize) pure returns(uint64)
func (_Store *StoreCallerSession) CalcValidFee(setting Setting, proveTime *big.Int, copyNum uint64, fileSize uint64) (uint64, error) {
	return _Store.Contract.CalcValidFee(&_Store.CallOpts, setting, proveTime, copyNum, fileSize)
}

// CalcValidFeeForOneNode is a free data retrieval call binding the contract method 0x44b2d82d.
//
// Solidity: function calcValidFeeForOneNode((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 proveTime, uint64 fileSize) pure returns(uint64)
func (_Store *StoreCaller) CalcValidFeeForOneNode(opts *bind.CallOpts, setting Setting, proveTime *big.Int, fileSize uint64) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "calcValidFeeForOneNode", setting, proveTime, fileSize)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalcValidFeeForOneNode is a free data retrieval call binding the contract method 0x44b2d82d.
//
// Solidity: function calcValidFeeForOneNode((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 proveTime, uint64 fileSize) pure returns(uint64)
func (_Store *StoreSession) CalcValidFeeForOneNode(setting Setting, proveTime *big.Int, fileSize uint64) (uint64, error) {
	return _Store.Contract.CalcValidFeeForOneNode(&_Store.CallOpts, setting, proveTime, fileSize)
}

// CalcValidFeeForOneNode is a free data retrieval call binding the contract method 0x44b2d82d.
//
// Solidity: function calcValidFeeForOneNode((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 proveTime, uint64 fileSize) pure returns(uint64)
func (_Store *StoreCallerSession) CalcValidFeeForOneNode(setting Setting, proveTime *big.Int, fileSize uint64) (uint64, error) {
	return _Store.Contract.CalcValidFeeForOneNode(&_Store.CallOpts, setting, proveTime, fileSize)
}

// AddUserSpace is a paid mutator transaction binding the contract method 0x6a4fe520.
//
// Solidity: function AddUserSpace(((uint64,uint64,uint64,uint256,uint256),uint64,uint64,uint256,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64),bytes[]) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreTransactor) AddUserSpace(opts *bind.TransactOpts, params SpaceAddParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "AddUserSpace", params)
}

// AddUserSpace is a paid mutator transaction binding the contract method 0x6a4fe520.
//
// Solidity: function AddUserSpace(((uint64,uint64,uint64,uint256,uint256),uint64,uint64,uint256,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64),bytes[]) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreSession) AddUserSpace(params SpaceAddParams) (*types.Transaction, error) {
	return _Store.Contract.AddUserSpace(&_Store.TransactOpts, params)
}

// AddUserSpace is a paid mutator transaction binding the contract method 0x6a4fe520.
//
// Solidity: function AddUserSpace(((uint64,uint64,uint64,uint256,uint256),uint64,uint64,uint256,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64),bytes[]) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreTransactorSession) AddUserSpace(params SpaceAddParams) (*types.Transaction, error) {
	return _Store.Contract.AddUserSpace(&_Store.TransactOpts, params)
}

// DeleteUserSpace is a paid mutator transaction binding the contract method 0x0f9fa2eb.
//
// Solidity: function DeleteUserSpace(address walletAddr) returns()
func (_Store *StoreTransactor) DeleteUserSpace(opts *bind.TransactOpts, walletAddr common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteUserSpace", walletAddr)
}

// DeleteUserSpace is a paid mutator transaction binding the contract method 0x0f9fa2eb.
//
// Solidity: function DeleteUserSpace(address walletAddr) returns()
func (_Store *StoreSession) DeleteUserSpace(walletAddr common.Address) (*types.Transaction, error) {
	return _Store.Contract.DeleteUserSpace(&_Store.TransactOpts, walletAddr)
}

// DeleteUserSpace is a paid mutator transaction binding the contract method 0x0f9fa2eb.
//
// Solidity: function DeleteUserSpace(address walletAddr) returns()
func (_Store *StoreTransactorSession) DeleteUserSpace(walletAddr common.Address) (*types.Transaction, error) {
	return _Store.Contract.DeleteUserSpace(&_Store.TransactOpts, walletAddr)
}

// ManageUserSpace is a paid mutator transaction binding the contract method 0x3899831a.
//
// Solidity: function ManageUserSpace((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns()
func (_Store *StoreTransactor) ManageUserSpace(opts *bind.TransactOpts, params UserSpaceParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ManageUserSpace", params)
}

// ManageUserSpace is a paid mutator transaction binding the contract method 0x3899831a.
//
// Solidity: function ManageUserSpace((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns()
func (_Store *StoreSession) ManageUserSpace(params UserSpaceParams) (*types.Transaction, error) {
	return _Store.Contract.ManageUserSpace(&_Store.TransactOpts, params)
}

// ManageUserSpace is a paid mutator transaction binding the contract method 0x3899831a.
//
// Solidity: function ManageUserSpace((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns()
func (_Store *StoreTransactorSession) ManageUserSpace(params UserSpaceParams) (*types.Transaction, error) {
	return _Store.Contract.ManageUserSpace(&_Store.TransactOpts, params)
}

// UpdateUserSpace is a paid mutator transaction binding the contract method 0xed108de9.
//
// Solidity: function UpdateUserSpace(address walletAddr, (uint64,uint64,uint64,uint256,uint256) _userSpace) payable returns()
func (_Store *StoreTransactor) UpdateUserSpace(opts *bind.TransactOpts, walletAddr common.Address, _userSpace UserSpace) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateUserSpace", walletAddr, _userSpace)
}

// UpdateUserSpace is a paid mutator transaction binding the contract method 0xed108de9.
//
// Solidity: function UpdateUserSpace(address walletAddr, (uint64,uint64,uint64,uint256,uint256) _userSpace) payable returns()
func (_Store *StoreSession) UpdateUserSpace(walletAddr common.Address, _userSpace UserSpace) (*types.Transaction, error) {
	return _Store.Contract.UpdateUserSpace(&_Store.TransactOpts, walletAddr, _userSpace)
}

// UpdateUserSpace is a paid mutator transaction binding the contract method 0xed108de9.
//
// Solidity: function UpdateUserSpace(address walletAddr, (uint64,uint64,uint64,uint256,uint256) _userSpace) payable returns()
func (_Store *StoreTransactorSession) UpdateUserSpace(walletAddr common.Address, _userSpace UserSpace) (*types.Transaction, error) {
	return _Store.Contract.UpdateUserSpace(&_Store.TransactOpts, walletAddr, _userSpace)
}

// GetUserspaceChange is a paid mutator transaction binding the contract method 0x722b25b9.
//
// Solidity: function getUserspaceChange((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(address,address,uint64),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[]), string)
func (_Store *StoreTransactor) GetUserspaceChange(opts *bind.TransactOpts, params UserSpaceParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "getUserspaceChange", params)
}

// GetUserspaceChange is a paid mutator transaction binding the contract method 0x722b25b9.
//
// Solidity: function getUserspaceChange((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(address,address,uint64),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[]), string)
func (_Store *StoreSession) GetUserspaceChange(params UserSpaceParams) (*types.Transaction, error) {
	return _Store.Contract.GetUserspaceChange(&_Store.TransactOpts, params)
}

// GetUserspaceChange is a paid mutator transaction binding the contract method 0x722b25b9.
//
// Solidity: function getUserspaceChange((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(address,address,uint64),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[]), string)
func (_Store *StoreTransactorSession) GetUserspaceChange(params UserSpaceParams) (*types.Transaction, error) {
	return _Store.Contract.GetUserspaceChange(&_Store.TransactOpts, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _config, address _fs) returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts, _config common.Address, _fs common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize", _config, _fs)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _config, address _fs) returns()
func (_Store *StoreSession) Initialize(_config common.Address, _fs common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _fs)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _config, address _fs) returns()
func (_Store *StoreTransactorSession) Initialize(_config common.Address, _fs common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _fs)
}

// ProcessForUserSpaceOneAddOneRevoke is a paid mutator transaction binding the contract method 0xf1feacbf.
//
// Solidity: function processForUserSpaceOneAddOneRevoke(((address,address,(uint8,uint64),(uint8,uint64)),(uint64,uint64,uint64,uint256,uint256),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64),bytes[],uint64) params) payable returns(((uint64,uint64,uint64,uint256,uint256),uint64,uint64,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreTransactor) ProcessForUserSpaceOneAddOneRevoke(opts *bind.TransactOpts, params SpaceProcessRevokeParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "processForUserSpaceOneAddOneRevoke", params)
}

// ProcessForUserSpaceOneAddOneRevoke is a paid mutator transaction binding the contract method 0xf1feacbf.
//
// Solidity: function processForUserSpaceOneAddOneRevoke(((address,address,(uint8,uint64),(uint8,uint64)),(uint64,uint64,uint64,uint256,uint256),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64),bytes[],uint64) params) payable returns(((uint64,uint64,uint64,uint256,uint256),uint64,uint64,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreSession) ProcessForUserSpaceOneAddOneRevoke(params SpaceProcessRevokeParams) (*types.Transaction, error) {
	return _Store.Contract.ProcessForUserSpaceOneAddOneRevoke(&_Store.TransactOpts, params)
}

// ProcessForUserSpaceOneAddOneRevoke is a paid mutator transaction binding the contract method 0xf1feacbf.
//
// Solidity: function processForUserSpaceOneAddOneRevoke(((address,address,(uint8,uint64),(uint8,uint64)),(uint64,uint64,uint64,uint256,uint256),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64),bytes[],uint64) params) payable returns(((uint64,uint64,uint64,uint256,uint256),uint64,uint64,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreTransactorSession) ProcessForUserSpaceOneAddOneRevoke(params SpaceProcessRevokeParams) (*types.Transaction, error) {
	return _Store.Contract.ProcessForUserSpaceOneAddOneRevoke(&_Store.TransactOpts, params)
}

// ProcessForUserSpaceOperations is a paid mutator transaction binding the contract method 0x6c799c13.
//
// Solidity: function processForUserSpaceOperations(((address,address,(uint8,uint64),(uint8,uint64)),(uint64,uint64,uint64,uint256,uint256),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64)) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(address,address,uint64),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreTransactor) ProcessForUserSpaceOperations(opts *bind.TransactOpts, params SpaceProcessParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "processForUserSpaceOperations", params)
}

// ProcessForUserSpaceOperations is a paid mutator transaction binding the contract method 0x6c799c13.
//
// Solidity: function processForUserSpaceOperations(((address,address,(uint8,uint64),(uint8,uint64)),(uint64,uint64,uint64,uint256,uint256),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64)) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(address,address,uint64),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreSession) ProcessForUserSpaceOperations(params SpaceProcessParams) (*types.Transaction, error) {
	return _Store.Contract.ProcessForUserSpaceOperations(&_Store.TransactOpts, params)
}

// ProcessForUserSpaceOperations is a paid mutator transaction binding the contract method 0x6c799c13.
//
// Solidity: function processForUserSpaceOperations(((address,address,(uint8,uint64),(uint8,uint64)),(uint64,uint64,uint64,uint256,uint256),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64)) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(address,address,uint64),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreTransactorSession) ProcessForUserSpaceOperations(params SpaceProcessParams) (*types.Transaction, error) {
	return _Store.Contract.ProcessForUserSpaceOperations(&_Store.TransactOpts, params)
}

// StoreCreateSectorEventIterator is returned from FilterCreateSectorEvent and is used to iterate over the raw logs and unpacked data for CreateSectorEvent events raised by the Store contract.
type StoreCreateSectorEventIterator struct {
	Event *StoreCreateSectorEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreCreateSectorEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreCreateSectorEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreCreateSectorEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreCreateSectorEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreCreateSectorEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreCreateSectorEvent represents a CreateSectorEvent event raised by the Store contract.
type StoreCreateSectorEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	WalletAddr  common.Address
	SectorId    uint64
	ProveLevel  uint8
	Size        uint64
	IsPlots     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreateSectorEvent is a free log retrieval operation binding the contract event 0xc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb.
//
// Solidity: event CreateSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId, uint8 proveLevel, uint64 size, bool isPlots)
func (_Store *StoreFilterer) FilterCreateSectorEvent(opts *bind.FilterOpts) (*StoreCreateSectorEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "CreateSectorEvent")
	if err != nil {
		return nil, err
	}
	return &StoreCreateSectorEventIterator{contract: _Store.contract, event: "CreateSectorEvent", logs: logs, sub: sub}, nil
}

// WatchCreateSectorEvent is a free log subscription operation binding the contract event 0xc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb.
//
// Solidity: event CreateSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId, uint8 proveLevel, uint64 size, bool isPlots)
func (_Store *StoreFilterer) WatchCreateSectorEvent(opts *bind.WatchOpts, sink chan<- *StoreCreateSectorEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "CreateSectorEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreCreateSectorEvent)
				if err := _Store.contract.UnpackLog(event, "CreateSectorEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreateSectorEvent is a log parse operation binding the contract event 0xc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb.
//
// Solidity: event CreateSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId, uint8 proveLevel, uint64 size, bool isPlots)
func (_Store *StoreFilterer) ParseCreateSectorEvent(log types.Log) (*StoreCreateSectorEvent, error) {
	event := new(StoreCreateSectorEvent)
	if err := _Store.contract.UnpackLog(event, "CreateSectorEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDNSNodeRegisterIterator is returned from FilterDNSNodeRegister and is used to iterate over the raw logs and unpacked data for DNSNodeRegister events raised by the Store contract.
type StoreDNSNodeRegisterIterator struct {
	Event *StoreDNSNodeRegister // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreDNSNodeRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDNSNodeRegister)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreDNSNodeRegister)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreDNSNodeRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDNSNodeRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDNSNodeRegister represents a DNSNodeRegister event raised by the Store contract.
type StoreDNSNodeRegister struct {
	Ip         []byte
	Port       []byte
	WalletAddr common.Address
	Deposit    uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDNSNodeRegister is a free log retrieval operation binding the contract event 0x602afcbd95c0dda31970fec49aaef3df0519ed4e405e43e1b0b4c7050d6fa2bf.
//
// Solidity: event DNSNodeRegister(bytes ip, bytes port, address walletAddr, uint64 deposit)
func (_Store *StoreFilterer) FilterDNSNodeRegister(opts *bind.FilterOpts) (*StoreDNSNodeRegisterIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DNSNodeRegister")
	if err != nil {
		return nil, err
	}
	return &StoreDNSNodeRegisterIterator{contract: _Store.contract, event: "DNSNodeRegister", logs: logs, sub: sub}, nil
}

// WatchDNSNodeRegister is a free log subscription operation binding the contract event 0x602afcbd95c0dda31970fec49aaef3df0519ed4e405e43e1b0b4c7050d6fa2bf.
//
// Solidity: event DNSNodeRegister(bytes ip, bytes port, address walletAddr, uint64 deposit)
func (_Store *StoreFilterer) WatchDNSNodeRegister(opts *bind.WatchOpts, sink chan<- *StoreDNSNodeRegister) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DNSNodeRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDNSNodeRegister)
				if err := _Store.contract.UnpackLog(event, "DNSNodeRegister", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDNSNodeRegister is a log parse operation binding the contract event 0x602afcbd95c0dda31970fec49aaef3df0519ed4e405e43e1b0b4c7050d6fa2bf.
//
// Solidity: event DNSNodeRegister(bytes ip, bytes port, address walletAddr, uint64 deposit)
func (_Store *StoreFilterer) ParseDNSNodeRegister(log types.Log) (*StoreDNSNodeRegister, error) {
	event := new(StoreDNSNodeRegister)
	if err := _Store.contract.UnpackLog(event, "DNSNodeRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDNSNodeUnRegIterator is returned from FilterDNSNodeUnReg and is used to iterate over the raw logs and unpacked data for DNSNodeUnReg events raised by the Store contract.
type StoreDNSNodeUnRegIterator struct {
	Event *StoreDNSNodeUnReg // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreDNSNodeUnRegIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDNSNodeUnReg)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreDNSNodeUnReg)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreDNSNodeUnRegIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDNSNodeUnRegIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDNSNodeUnReg represents a DNSNodeUnReg event raised by the Store contract.
type StoreDNSNodeUnReg struct {
	WalletAddr common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDNSNodeUnReg is a free log retrieval operation binding the contract event 0x7e854b98aa965bf68504b0e009e217e7dae7ae9b1cbde5d5968ecbd85fc5e11d.
//
// Solidity: event DNSNodeUnReg(address walletAddr)
func (_Store *StoreFilterer) FilterDNSNodeUnReg(opts *bind.FilterOpts) (*StoreDNSNodeUnRegIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DNSNodeUnReg")
	if err != nil {
		return nil, err
	}
	return &StoreDNSNodeUnRegIterator{contract: _Store.contract, event: "DNSNodeUnReg", logs: logs, sub: sub}, nil
}

// WatchDNSNodeUnReg is a free log subscription operation binding the contract event 0x7e854b98aa965bf68504b0e009e217e7dae7ae9b1cbde5d5968ecbd85fc5e11d.
//
// Solidity: event DNSNodeUnReg(address walletAddr)
func (_Store *StoreFilterer) WatchDNSNodeUnReg(opts *bind.WatchOpts, sink chan<- *StoreDNSNodeUnReg) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DNSNodeUnReg")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDNSNodeUnReg)
				if err := _Store.contract.UnpackLog(event, "DNSNodeUnReg", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDNSNodeUnReg is a log parse operation binding the contract event 0x7e854b98aa965bf68504b0e009e217e7dae7ae9b1cbde5d5968ecbd85fc5e11d.
//
// Solidity: event DNSNodeUnReg(address walletAddr)
func (_Store *StoreFilterer) ParseDNSNodeUnReg(log types.Log) (*StoreDNSNodeUnReg, error) {
	event := new(StoreDNSNodeUnReg)
	if err := _Store.contract.UnpackLog(event, "DNSNodeUnReg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDeleteFileEventIterator is returned from FilterDeleteFileEvent and is used to iterate over the raw logs and unpacked data for DeleteFileEvent events raised by the Store contract.
type StoreDeleteFileEventIterator struct {
	Event *StoreDeleteFileEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreDeleteFileEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDeleteFileEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreDeleteFileEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreDeleteFileEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDeleteFileEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDeleteFileEvent represents a DeleteFileEvent event raised by the Store contract.
type StoreDeleteFileEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	FileHash    []byte
	WalletAddr  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeleteFileEvent is a free log retrieval operation binding the contract event 0xb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f45.
//
// Solidity: event DeleteFileEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, address walletAddr)
func (_Store *StoreFilterer) FilterDeleteFileEvent(opts *bind.FilterOpts) (*StoreDeleteFileEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DeleteFileEvent")
	if err != nil {
		return nil, err
	}
	return &StoreDeleteFileEventIterator{contract: _Store.contract, event: "DeleteFileEvent", logs: logs, sub: sub}, nil
}

// WatchDeleteFileEvent is a free log subscription operation binding the contract event 0xb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f45.
//
// Solidity: event DeleteFileEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, address walletAddr)
func (_Store *StoreFilterer) WatchDeleteFileEvent(opts *bind.WatchOpts, sink chan<- *StoreDeleteFileEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DeleteFileEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDeleteFileEvent)
				if err := _Store.contract.UnpackLog(event, "DeleteFileEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeleteFileEvent is a log parse operation binding the contract event 0xb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f45.
//
// Solidity: event DeleteFileEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, address walletAddr)
func (_Store *StoreFilterer) ParseDeleteFileEvent(log types.Log) (*StoreDeleteFileEvent, error) {
	event := new(StoreDeleteFileEvent)
	if err := _Store.contract.UnpackLog(event, "DeleteFileEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDeleteFilesEventIterator is returned from FilterDeleteFilesEvent and is used to iterate over the raw logs and unpacked data for DeleteFilesEvent events raised by the Store contract.
type StoreDeleteFilesEventIterator struct {
	Event *StoreDeleteFilesEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreDeleteFilesEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDeleteFilesEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreDeleteFilesEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreDeleteFilesEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDeleteFilesEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDeleteFilesEvent represents a DeleteFilesEvent event raised by the Store contract.
type StoreDeleteFilesEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	FileHashs   [][]byte
	WalletAddr  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeleteFilesEvent is a free log retrieval operation binding the contract event 0xa5d0a6140e999641864ed8958af5eea1d36d821658cfc056552962fab40291ec.
//
// Solidity: event DeleteFilesEvent(uint8 eventType, uint256 blockHeight, bytes[] fileHashs, address walletAddr)
func (_Store *StoreFilterer) FilterDeleteFilesEvent(opts *bind.FilterOpts) (*StoreDeleteFilesEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DeleteFilesEvent")
	if err != nil {
		return nil, err
	}
	return &StoreDeleteFilesEventIterator{contract: _Store.contract, event: "DeleteFilesEvent", logs: logs, sub: sub}, nil
}

// WatchDeleteFilesEvent is a free log subscription operation binding the contract event 0xa5d0a6140e999641864ed8958af5eea1d36d821658cfc056552962fab40291ec.
//
// Solidity: event DeleteFilesEvent(uint8 eventType, uint256 blockHeight, bytes[] fileHashs, address walletAddr)
func (_Store *StoreFilterer) WatchDeleteFilesEvent(opts *bind.WatchOpts, sink chan<- *StoreDeleteFilesEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DeleteFilesEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDeleteFilesEvent)
				if err := _Store.contract.UnpackLog(event, "DeleteFilesEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeleteFilesEvent is a log parse operation binding the contract event 0xa5d0a6140e999641864ed8958af5eea1d36d821658cfc056552962fab40291ec.
//
// Solidity: event DeleteFilesEvent(uint8 eventType, uint256 blockHeight, bytes[] fileHashs, address walletAddr)
func (_Store *StoreFilterer) ParseDeleteFilesEvent(log types.Log) (*StoreDeleteFilesEvent, error) {
	event := new(StoreDeleteFilesEvent)
	if err := _Store.contract.UnpackLog(event, "DeleteFilesEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDeleteSectorEventIterator is returned from FilterDeleteSectorEvent and is used to iterate over the raw logs and unpacked data for DeleteSectorEvent events raised by the Store contract.
type StoreDeleteSectorEventIterator struct {
	Event *StoreDeleteSectorEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreDeleteSectorEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDeleteSectorEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreDeleteSectorEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreDeleteSectorEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDeleteSectorEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDeleteSectorEvent represents a DeleteSectorEvent event raised by the Store contract.
type StoreDeleteSectorEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	WalletAddr  common.Address
	SectorId    uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeleteSectorEvent is a free log retrieval operation binding the contract event 0x4fca90fd1b1cd962cf67f21703f1380572181450811cdd09c4add7ed1cb0c12c.
//
// Solidity: event DeleteSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId)
func (_Store *StoreFilterer) FilterDeleteSectorEvent(opts *bind.FilterOpts) (*StoreDeleteSectorEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DeleteSectorEvent")
	if err != nil {
		return nil, err
	}
	return &StoreDeleteSectorEventIterator{contract: _Store.contract, event: "DeleteSectorEvent", logs: logs, sub: sub}, nil
}

// WatchDeleteSectorEvent is a free log subscription operation binding the contract event 0x4fca90fd1b1cd962cf67f21703f1380572181450811cdd09c4add7ed1cb0c12c.
//
// Solidity: event DeleteSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId)
func (_Store *StoreFilterer) WatchDeleteSectorEvent(opts *bind.WatchOpts, sink chan<- *StoreDeleteSectorEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DeleteSectorEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDeleteSectorEvent)
				if err := _Store.contract.UnpackLog(event, "DeleteSectorEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeleteSectorEvent is a log parse operation binding the contract event 0x4fca90fd1b1cd962cf67f21703f1380572181450811cdd09c4add7ed1cb0c12c.
//
// Solidity: event DeleteSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId)
func (_Store *StoreFilterer) ParseDeleteSectorEvent(log types.Log) (*StoreDeleteSectorEvent, error) {
	event := new(StoreDeleteSectorEvent)
	if err := _Store.contract.UnpackLog(event, "DeleteSectorEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDnsErrorIterator is returned from FilterDnsError and is used to iterate over the raw logs and unpacked data for DnsError events raised by the Store contract.
type StoreDnsErrorIterator struct {
	Event *StoreDnsError // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreDnsErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDnsError)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreDnsError)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreDnsErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDnsErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDnsError represents a DnsError event raised by the Store contract.
type StoreDnsError struct {
	Method string
	Msg    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDnsError is a free log retrieval operation binding the contract event 0xd6cf900f8b0f19180820a0d00eb7439f8143ae0a9c9c1b724fb637a51ae41e35.
//
// Solidity: event DnsError(string method, string msg)
func (_Store *StoreFilterer) FilterDnsError(opts *bind.FilterOpts) (*StoreDnsErrorIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DnsError")
	if err != nil {
		return nil, err
	}
	return &StoreDnsErrorIterator{contract: _Store.contract, event: "DnsError", logs: logs, sub: sub}, nil
}

// WatchDnsError is a free log subscription operation binding the contract event 0xd6cf900f8b0f19180820a0d00eb7439f8143ae0a9c9c1b724fb637a51ae41e35.
//
// Solidity: event DnsError(string method, string msg)
func (_Store *StoreFilterer) WatchDnsError(opts *bind.WatchOpts, sink chan<- *StoreDnsError) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DnsError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDnsError)
				if err := _Store.contract.UnpackLog(event, "DnsError", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDnsError is a log parse operation binding the contract event 0xd6cf900f8b0f19180820a0d00eb7439f8143ae0a9c9c1b724fb637a51ae41e35.
//
// Solidity: event DnsError(string method, string msg)
func (_Store *StoreFilterer) ParseDnsError(log types.Log) (*StoreDnsError, error) {
	event := new(StoreDnsError)
	if err := _Store.contract.UnpackLog(event, "DnsError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreFilePDPSuccessEventIterator is returned from FilterFilePDPSuccessEvent and is used to iterate over the raw logs and unpacked data for FilePDPSuccessEvent events raised by the Store contract.
type StoreFilePDPSuccessEventIterator struct {
	Event *StoreFilePDPSuccessEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreFilePDPSuccessEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreFilePDPSuccessEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreFilePDPSuccessEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreFilePDPSuccessEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreFilePDPSuccessEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreFilePDPSuccessEvent represents a FilePDPSuccessEvent event raised by the Store contract.
type StoreFilePDPSuccessEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	FileHash    []byte
	WalletAddr  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFilePDPSuccessEvent is a free log retrieval operation binding the contract event 0xd936063ae3cbbf2dbcd2adc837a997ab97adb09f24c566aa79d101e88928dea6.
//
// Solidity: event FilePDPSuccessEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, address walletAddr)
func (_Store *StoreFilterer) FilterFilePDPSuccessEvent(opts *bind.FilterOpts) (*StoreFilePDPSuccessEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "FilePDPSuccessEvent")
	if err != nil {
		return nil, err
	}
	return &StoreFilePDPSuccessEventIterator{contract: _Store.contract, event: "FilePDPSuccessEvent", logs: logs, sub: sub}, nil
}

// WatchFilePDPSuccessEvent is a free log subscription operation binding the contract event 0xd936063ae3cbbf2dbcd2adc837a997ab97adb09f24c566aa79d101e88928dea6.
//
// Solidity: event FilePDPSuccessEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, address walletAddr)
func (_Store *StoreFilterer) WatchFilePDPSuccessEvent(opts *bind.WatchOpts, sink chan<- *StoreFilePDPSuccessEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "FilePDPSuccessEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreFilePDPSuccessEvent)
				if err := _Store.contract.UnpackLog(event, "FilePDPSuccessEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFilePDPSuccessEvent is a log parse operation binding the contract event 0xd936063ae3cbbf2dbcd2adc837a997ab97adb09f24c566aa79d101e88928dea6.
//
// Solidity: event FilePDPSuccessEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, address walletAddr)
func (_Store *StoreFilterer) ParseFilePDPSuccessEvent(log types.Log) (*StoreFilePDPSuccessEvent, error) {
	event := new(StoreFilePDPSuccessEvent)
	if err := _Store.contract.UnpackLog(event, "FilePDPSuccessEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreFsErrorIterator is returned from FilterFsError and is used to iterate over the raw logs and unpacked data for FsError events raised by the Store contract.
type StoreFsErrorIterator struct {
	Event *StoreFsError // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreFsErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreFsError)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreFsError)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreFsErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreFsErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreFsError represents a FsError event raised by the Store contract.
type StoreFsError struct {
	Method string
	Msg    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFsError is a free log retrieval operation binding the contract event 0x61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5.
//
// Solidity: event FsError(string method, string msg)
func (_Store *StoreFilterer) FilterFsError(opts *bind.FilterOpts) (*StoreFsErrorIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "FsError")
	if err != nil {
		return nil, err
	}
	return &StoreFsErrorIterator{contract: _Store.contract, event: "FsError", logs: logs, sub: sub}, nil
}

// WatchFsError is a free log subscription operation binding the contract event 0x61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5.
//
// Solidity: event FsError(string method, string msg)
func (_Store *StoreFilterer) WatchFsError(opts *bind.WatchOpts, sink chan<- *StoreFsError) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "FsError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreFsError)
				if err := _Store.contract.UnpackLog(event, "FsError", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFsError is a log parse operation binding the contract event 0x61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5.
//
// Solidity: event FsError(string method, string msg)
func (_Store *StoreFilterer) ParseFsError(log types.Log) (*StoreFsError, error) {
	event := new(StoreFsError)
	if err := _Store.contract.UnpackLog(event, "FsError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreGetUpdateCostEventIterator is returned from FilterGetUpdateCostEvent and is used to iterate over the raw logs and unpacked data for GetUpdateCostEvent events raised by the Store contract.
type StoreGetUpdateCostEventIterator struct {
	Event *StoreGetUpdateCostEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreGetUpdateCostEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreGetUpdateCostEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreGetUpdateCostEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreGetUpdateCostEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreGetUpdateCostEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreGetUpdateCostEvent represents a GetUpdateCostEvent event raised by the Store contract.
type StoreGetUpdateCostEvent struct {
	State TransferState
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGetUpdateCostEvent is a free log retrieval operation binding the contract event 0x7658e81ba2dec69f3dcdee55eae0141d069972172244313e3e13419927008767.
//
// Solidity: event GetUpdateCostEvent((address,address,uint64) state)
func (_Store *StoreFilterer) FilterGetUpdateCostEvent(opts *bind.FilterOpts) (*StoreGetUpdateCostEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "GetUpdateCostEvent")
	if err != nil {
		return nil, err
	}
	return &StoreGetUpdateCostEventIterator{contract: _Store.contract, event: "GetUpdateCostEvent", logs: logs, sub: sub}, nil
}

// WatchGetUpdateCostEvent is a free log subscription operation binding the contract event 0x7658e81ba2dec69f3dcdee55eae0141d069972172244313e3e13419927008767.
//
// Solidity: event GetUpdateCostEvent((address,address,uint64) state)
func (_Store *StoreFilterer) WatchGetUpdateCostEvent(opts *bind.WatchOpts, sink chan<- *StoreGetUpdateCostEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "GetUpdateCostEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreGetUpdateCostEvent)
				if err := _Store.contract.UnpackLog(event, "GetUpdateCostEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGetUpdateCostEvent is a log parse operation binding the contract event 0x7658e81ba2dec69f3dcdee55eae0141d069972172244313e3e13419927008767.
//
// Solidity: event GetUpdateCostEvent((address,address,uint64) state)
func (_Store *StoreFilterer) ParseGetUpdateCostEvent(log types.Log) (*StoreGetUpdateCostEvent, error) {
	event := new(StoreGetUpdateCostEvent)
	if err := _Store.contract.UnpackLog(event, "GetUpdateCostEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreNotifyHeaderAddIterator is returned from FilterNotifyHeaderAdd and is used to iterate over the raw logs and unpacked data for NotifyHeaderAdd events raised by the Store contract.
type StoreNotifyHeaderAddIterator struct {
	Event *StoreNotifyHeaderAdd // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreNotifyHeaderAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNotifyHeaderAdd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreNotifyHeaderAdd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreNotifyHeaderAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNotifyHeaderAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNotifyHeaderAdd represents a NotifyHeaderAdd event raised by the Store contract.
type StoreNotifyHeaderAdd struct {
	Owner  common.Address
	Header []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotifyHeaderAdd is a free log retrieval operation binding the contract event 0x356512f13710983e552444fcd4bd47c18383e96dfe51938f6d64117da87cc869.
//
// Solidity: event NotifyHeaderAdd(address owner, bytes header)
func (_Store *StoreFilterer) FilterNotifyHeaderAdd(opts *bind.FilterOpts) (*StoreNotifyHeaderAddIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NotifyHeaderAdd")
	if err != nil {
		return nil, err
	}
	return &StoreNotifyHeaderAddIterator{contract: _Store.contract, event: "NotifyHeaderAdd", logs: logs, sub: sub}, nil
}

// WatchNotifyHeaderAdd is a free log subscription operation binding the contract event 0x356512f13710983e552444fcd4bd47c18383e96dfe51938f6d64117da87cc869.
//
// Solidity: event NotifyHeaderAdd(address owner, bytes header)
func (_Store *StoreFilterer) WatchNotifyHeaderAdd(opts *bind.WatchOpts, sink chan<- *StoreNotifyHeaderAdd) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NotifyHeaderAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNotifyHeaderAdd)
				if err := _Store.contract.UnpackLog(event, "NotifyHeaderAdd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotifyHeaderAdd is a log parse operation binding the contract event 0x356512f13710983e552444fcd4bd47c18383e96dfe51938f6d64117da87cc869.
//
// Solidity: event NotifyHeaderAdd(address owner, bytes header)
func (_Store *StoreFilterer) ParseNotifyHeaderAdd(log types.Log) (*StoreNotifyHeaderAdd, error) {
	event := new(StoreNotifyHeaderAdd)
	if err := _Store.contract.UnpackLog(event, "NotifyHeaderAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreNotifyHeaderTransferIterator is returned from FilterNotifyHeaderTransfer and is used to iterate over the raw logs and unpacked data for NotifyHeaderTransfer events raised by the Store contract.
type StoreNotifyHeaderTransferIterator struct {
	Event *StoreNotifyHeaderTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreNotifyHeaderTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNotifyHeaderTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreNotifyHeaderTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreNotifyHeaderTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNotifyHeaderTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNotifyHeaderTransfer represents a NotifyHeaderTransfer event raised by the Store contract.
type StoreNotifyHeaderTransfer struct {
	From   common.Address
	To     common.Address
	Header []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotifyHeaderTransfer is a free log retrieval operation binding the contract event 0x178b65a7736b9ddda8192e46b4aa3b7916e057db13cde938ab6818a4450253ca.
//
// Solidity: event NotifyHeaderTransfer(address from, address to, bytes header)
func (_Store *StoreFilterer) FilterNotifyHeaderTransfer(opts *bind.FilterOpts) (*StoreNotifyHeaderTransferIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NotifyHeaderTransfer")
	if err != nil {
		return nil, err
	}
	return &StoreNotifyHeaderTransferIterator{contract: _Store.contract, event: "NotifyHeaderTransfer", logs: logs, sub: sub}, nil
}

// WatchNotifyHeaderTransfer is a free log subscription operation binding the contract event 0x178b65a7736b9ddda8192e46b4aa3b7916e057db13cde938ab6818a4450253ca.
//
// Solidity: event NotifyHeaderTransfer(address from, address to, bytes header)
func (_Store *StoreFilterer) WatchNotifyHeaderTransfer(opts *bind.WatchOpts, sink chan<- *StoreNotifyHeaderTransfer) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NotifyHeaderTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNotifyHeaderTransfer)
				if err := _Store.contract.UnpackLog(event, "NotifyHeaderTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotifyHeaderTransfer is a log parse operation binding the contract event 0x178b65a7736b9ddda8192e46b4aa3b7916e057db13cde938ab6818a4450253ca.
//
// Solidity: event NotifyHeaderTransfer(address from, address to, bytes header)
func (_Store *StoreFilterer) ParseNotifyHeaderTransfer(log types.Log) (*StoreNotifyHeaderTransfer, error) {
	event := new(StoreNotifyHeaderTransfer)
	if err := _Store.contract.UnpackLog(event, "NotifyHeaderTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreNotifyNameInfoAddIterator is returned from FilterNotifyNameInfoAdd and is used to iterate over the raw logs and unpacked data for NotifyNameInfoAdd events raised by the Store contract.
type StoreNotifyNameInfoAddIterator struct {
	Event *StoreNotifyNameInfoAdd // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreNotifyNameInfoAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNotifyNameInfoAdd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreNotifyNameInfoAdd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreNotifyNameInfoAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNotifyNameInfoAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNotifyNameInfoAdd represents a NotifyNameInfoAdd event raised by the Store contract.
type StoreNotifyNameInfoAdd struct {
	Owner common.Address
	Url   []byte
	Newer NameInfo
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNotifyNameInfoAdd is a free log retrieval operation binding the contract event 0x517828a430f1ccf98c9b3af7c0aba5d12949858f987f9ffb773eab305e801e41.
//
// Solidity: event NotifyNameInfoAdd(address owner, bytes url, (uint64,bytes,bytes,bytes,address,bytes,uint256,uint64) newer)
func (_Store *StoreFilterer) FilterNotifyNameInfoAdd(opts *bind.FilterOpts) (*StoreNotifyNameInfoAddIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NotifyNameInfoAdd")
	if err != nil {
		return nil, err
	}
	return &StoreNotifyNameInfoAddIterator{contract: _Store.contract, event: "NotifyNameInfoAdd", logs: logs, sub: sub}, nil
}

// WatchNotifyNameInfoAdd is a free log subscription operation binding the contract event 0x517828a430f1ccf98c9b3af7c0aba5d12949858f987f9ffb773eab305e801e41.
//
// Solidity: event NotifyNameInfoAdd(address owner, bytes url, (uint64,bytes,bytes,bytes,address,bytes,uint256,uint64) newer)
func (_Store *StoreFilterer) WatchNotifyNameInfoAdd(opts *bind.WatchOpts, sink chan<- *StoreNotifyNameInfoAdd) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NotifyNameInfoAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNotifyNameInfoAdd)
				if err := _Store.contract.UnpackLog(event, "NotifyNameInfoAdd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotifyNameInfoAdd is a log parse operation binding the contract event 0x517828a430f1ccf98c9b3af7c0aba5d12949858f987f9ffb773eab305e801e41.
//
// Solidity: event NotifyNameInfoAdd(address owner, bytes url, (uint64,bytes,bytes,bytes,address,bytes,uint256,uint64) newer)
func (_Store *StoreFilterer) ParseNotifyNameInfoAdd(log types.Log) (*StoreNotifyNameInfoAdd, error) {
	event := new(StoreNotifyNameInfoAdd)
	if err := _Store.contract.UnpackLog(event, "NotifyNameInfoAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreNotifyNameInfoChangeIterator is returned from FilterNotifyNameInfoChange and is used to iterate over the raw logs and unpacked data for NotifyNameInfoChange events raised by the Store contract.
type StoreNotifyNameInfoChangeIterator struct {
	Event *StoreNotifyNameInfoChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreNotifyNameInfoChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNotifyNameInfoChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreNotifyNameInfoChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreNotifyNameInfoChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNotifyNameInfoChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNotifyNameInfoChange represents a NotifyNameInfoChange event raised by the Store contract.
type StoreNotifyNameInfoChange struct {
	Owner common.Address
	Url   []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNotifyNameInfoChange is a free log retrieval operation binding the contract event 0x0fe7706eab23c889b6b5ba01289446319a34004a5a0fa59306dbd02f9fe08150.
//
// Solidity: event NotifyNameInfoChange(address owner, bytes url)
func (_Store *StoreFilterer) FilterNotifyNameInfoChange(opts *bind.FilterOpts) (*StoreNotifyNameInfoChangeIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NotifyNameInfoChange")
	if err != nil {
		return nil, err
	}
	return &StoreNotifyNameInfoChangeIterator{contract: _Store.contract, event: "NotifyNameInfoChange", logs: logs, sub: sub}, nil
}

// WatchNotifyNameInfoChange is a free log subscription operation binding the contract event 0x0fe7706eab23c889b6b5ba01289446319a34004a5a0fa59306dbd02f9fe08150.
//
// Solidity: event NotifyNameInfoChange(address owner, bytes url)
func (_Store *StoreFilterer) WatchNotifyNameInfoChange(opts *bind.WatchOpts, sink chan<- *StoreNotifyNameInfoChange) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NotifyNameInfoChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNotifyNameInfoChange)
				if err := _Store.contract.UnpackLog(event, "NotifyNameInfoChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotifyNameInfoChange is a log parse operation binding the contract event 0x0fe7706eab23c889b6b5ba01289446319a34004a5a0fa59306dbd02f9fe08150.
//
// Solidity: event NotifyNameInfoChange(address owner, bytes url)
func (_Store *StoreFilterer) ParseNotifyNameInfoChange(log types.Log) (*StoreNotifyNameInfoChange, error) {
	event := new(StoreNotifyNameInfoChange)
	if err := _Store.contract.UnpackLog(event, "NotifyNameInfoChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreNotifyNameInfoTransferIterator is returned from FilterNotifyNameInfoTransfer and is used to iterate over the raw logs and unpacked data for NotifyNameInfoTransfer events raised by the Store contract.
type StoreNotifyNameInfoTransferIterator struct {
	Event *StoreNotifyNameInfoTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreNotifyNameInfoTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNotifyNameInfoTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreNotifyNameInfoTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreNotifyNameInfoTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNotifyNameInfoTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNotifyNameInfoTransfer represents a NotifyNameInfoTransfer event raised by the Store contract.
type StoreNotifyNameInfoTransfer struct {
	From common.Address
	To   common.Address
	Url  []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNotifyNameInfoTransfer is a free log retrieval operation binding the contract event 0xa65269eec0d2bac560b4cb6d761c6d074a751cc06a834fe32a025e5e21720e2f.
//
// Solidity: event NotifyNameInfoTransfer(address from, address to, bytes url)
func (_Store *StoreFilterer) FilterNotifyNameInfoTransfer(opts *bind.FilterOpts) (*StoreNotifyNameInfoTransferIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NotifyNameInfoTransfer")
	if err != nil {
		return nil, err
	}
	return &StoreNotifyNameInfoTransferIterator{contract: _Store.contract, event: "NotifyNameInfoTransfer", logs: logs, sub: sub}, nil
}

// WatchNotifyNameInfoTransfer is a free log subscription operation binding the contract event 0xa65269eec0d2bac560b4cb6d761c6d074a751cc06a834fe32a025e5e21720e2f.
//
// Solidity: event NotifyNameInfoTransfer(address from, address to, bytes url)
func (_Store *StoreFilterer) WatchNotifyNameInfoTransfer(opts *bind.WatchOpts, sink chan<- *StoreNotifyNameInfoTransfer) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NotifyNameInfoTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNotifyNameInfoTransfer)
				if err := _Store.contract.UnpackLog(event, "NotifyNameInfoTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotifyNameInfoTransfer is a log parse operation binding the contract event 0xa65269eec0d2bac560b4cb6d761c6d074a751cc06a834fe32a025e5e21720e2f.
//
// Solidity: event NotifyNameInfoTransfer(address from, address to, bytes url)
func (_Store *StoreFilterer) ParseNotifyNameInfoTransfer(log types.Log) (*StoreNotifyNameInfoTransfer, error) {
	event := new(StoreNotifyNameInfoTransfer)
	if err := _Store.contract.UnpackLog(event, "NotifyNameInfoTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorePDPVerifyEventIterator is returned from FilterPDPVerifyEvent and is used to iterate over the raw logs and unpacked data for PDPVerifyEvent events raised by the Store contract.
type StorePDPVerifyEventIterator struct {
	Event *StorePDPVerifyEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorePDPVerifyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorePDPVerifyEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorePDPVerifyEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorePDPVerifyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorePDPVerifyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorePDPVerifyEvent represents a PDPVerifyEvent event raised by the Store contract.
type StorePDPVerifyEvent struct {
	EventType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPDPVerifyEvent is a free log retrieval operation binding the contract event 0x7f6b6a2262c2fc2cc1b1785be448eae36656117d8f558facfe3c199e26256ea1.
//
// Solidity: event PDPVerifyEvent(uint8 eventType)
func (_Store *StoreFilterer) FilterPDPVerifyEvent(opts *bind.FilterOpts) (*StorePDPVerifyEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "PDPVerifyEvent")
	if err != nil {
		return nil, err
	}
	return &StorePDPVerifyEventIterator{contract: _Store.contract, event: "PDPVerifyEvent", logs: logs, sub: sub}, nil
}

// WatchPDPVerifyEvent is a free log subscription operation binding the contract event 0x7f6b6a2262c2fc2cc1b1785be448eae36656117d8f558facfe3c199e26256ea1.
//
// Solidity: event PDPVerifyEvent(uint8 eventType)
func (_Store *StoreFilterer) WatchPDPVerifyEvent(opts *bind.WatchOpts, sink chan<- *StorePDPVerifyEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "PDPVerifyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorePDPVerifyEvent)
				if err := _Store.contract.UnpackLog(event, "PDPVerifyEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePDPVerifyEvent is a log parse operation binding the contract event 0x7f6b6a2262c2fc2cc1b1785be448eae36656117d8f558facfe3c199e26256ea1.
//
// Solidity: event PDPVerifyEvent(uint8 eventType)
func (_Store *StoreFilterer) ParsePDPVerifyEvent(log types.Log) (*StorePDPVerifyEvent, error) {
	event := new(StorePDPVerifyEvent)
	if err := _Store.contract.UnpackLog(event, "PDPVerifyEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreProveFileEventIterator is returned from FilterProveFileEvent and is used to iterate over the raw logs and unpacked data for ProveFileEvent events raised by the Store contract.
type StoreProveFileEventIterator struct {
	Event *StoreProveFileEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreProveFileEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreProveFileEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreProveFileEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreProveFileEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreProveFileEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreProveFileEvent represents a ProveFileEvent event raised by the Store contract.
type StoreProveFileEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	WalletAddr  common.Address
	Profit      uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProveFileEvent is a free log retrieval operation binding the contract event 0x123b9998b2f027b02db4cb2eadef421fe9f1c21627569438588262d3a6baac6c.
//
// Solidity: event ProveFileEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 profit)
func (_Store *StoreFilterer) FilterProveFileEvent(opts *bind.FilterOpts) (*StoreProveFileEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "ProveFileEvent")
	if err != nil {
		return nil, err
	}
	return &StoreProveFileEventIterator{contract: _Store.contract, event: "ProveFileEvent", logs: logs, sub: sub}, nil
}

// WatchProveFileEvent is a free log subscription operation binding the contract event 0x123b9998b2f027b02db4cb2eadef421fe9f1c21627569438588262d3a6baac6c.
//
// Solidity: event ProveFileEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 profit)
func (_Store *StoreFilterer) WatchProveFileEvent(opts *bind.WatchOpts, sink chan<- *StoreProveFileEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "ProveFileEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreProveFileEvent)
				if err := _Store.contract.UnpackLog(event, "ProveFileEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProveFileEvent is a log parse operation binding the contract event 0x123b9998b2f027b02db4cb2eadef421fe9f1c21627569438588262d3a6baac6c.
//
// Solidity: event ProveFileEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 profit)
func (_Store *StoreFilterer) ParseProveFileEvent(log types.Log) (*StoreProveFileEvent, error) {
	event := new(StoreProveFileEvent)
	if err := _Store.contract.UnpackLog(event, "ProveFileEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreRegisterNodeEventIterator is returned from FilterRegisterNodeEvent and is used to iterate over the raw logs and unpacked data for RegisterNodeEvent events raised by the Store contract.
type StoreRegisterNodeEventIterator struct {
	Event *StoreRegisterNodeEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreRegisterNodeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreRegisterNodeEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreRegisterNodeEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreRegisterNodeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreRegisterNodeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreRegisterNodeEvent represents a RegisterNodeEvent event raised by the Store contract.
type StoreRegisterNodeEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	WalletAddr  common.Address
	NodeAddr    []byte
	Volume      uint64
	ServiceTime uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegisterNodeEvent is a free log retrieval operation binding the contract event 0x79778bfbb76fd8f676064400dca61e3f85ef0138b7c2944bdbd1c9cb131a0551.
//
// Solidity: event RegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr, bytes nodeAddr, uint64 volume, uint64 serviceTime)
func (_Store *StoreFilterer) FilterRegisterNodeEvent(opts *bind.FilterOpts) (*StoreRegisterNodeEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "RegisterNodeEvent")
	if err != nil {
		return nil, err
	}
	return &StoreRegisterNodeEventIterator{contract: _Store.contract, event: "RegisterNodeEvent", logs: logs, sub: sub}, nil
}

// WatchRegisterNodeEvent is a free log subscription operation binding the contract event 0x79778bfbb76fd8f676064400dca61e3f85ef0138b7c2944bdbd1c9cb131a0551.
//
// Solidity: event RegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr, bytes nodeAddr, uint64 volume, uint64 serviceTime)
func (_Store *StoreFilterer) WatchRegisterNodeEvent(opts *bind.WatchOpts, sink chan<- *StoreRegisterNodeEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "RegisterNodeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreRegisterNodeEvent)
				if err := _Store.contract.UnpackLog(event, "RegisterNodeEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegisterNodeEvent is a log parse operation binding the contract event 0x79778bfbb76fd8f676064400dca61e3f85ef0138b7c2944bdbd1c9cb131a0551.
//
// Solidity: event RegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr, bytes nodeAddr, uint64 volume, uint64 serviceTime)
func (_Store *StoreFilterer) ParseRegisterNodeEvent(log types.Log) (*StoreRegisterNodeEvent, error) {
	event := new(StoreRegisterNodeEvent)
	if err := _Store.contract.UnpackLog(event, "RegisterNodeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreSetUserSpaceEventIterator is returned from FilterSetUserSpaceEvent and is used to iterate over the raw logs and unpacked data for SetUserSpaceEvent events raised by the Store contract.
type StoreSetUserSpaceEventIterator struct {
	Event *StoreSetUserSpaceEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreSetUserSpaceEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreSetUserSpaceEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreSetUserSpaceEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreSetUserSpaceEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreSetUserSpaceEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreSetUserSpaceEvent represents a SetUserSpaceEvent event raised by the Store contract.
type StoreSetUserSpaceEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	WalletAddr  common.Address
	SizeType    uint8
	Size        uint64
	CountType   uint8
	Count       uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetUserSpaceEvent is a free log retrieval operation binding the contract event 0x0c5dd947b790e17c40c62b29772f0e0ca54c86e473b5bf74cfac2a1f9156ee91.
//
// Solidity: event SetUserSpaceEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint8 sizeType, uint64 size, uint8 countType, uint64 count)
func (_Store *StoreFilterer) FilterSetUserSpaceEvent(opts *bind.FilterOpts) (*StoreSetUserSpaceEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "SetUserSpaceEvent")
	if err != nil {
		return nil, err
	}
	return &StoreSetUserSpaceEventIterator{contract: _Store.contract, event: "SetUserSpaceEvent", logs: logs, sub: sub}, nil
}

// WatchSetUserSpaceEvent is a free log subscription operation binding the contract event 0x0c5dd947b790e17c40c62b29772f0e0ca54c86e473b5bf74cfac2a1f9156ee91.
//
// Solidity: event SetUserSpaceEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint8 sizeType, uint64 size, uint8 countType, uint64 count)
func (_Store *StoreFilterer) WatchSetUserSpaceEvent(opts *bind.WatchOpts, sink chan<- *StoreSetUserSpaceEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "SetUserSpaceEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreSetUserSpaceEvent)
				if err := _Store.contract.UnpackLog(event, "SetUserSpaceEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetUserSpaceEvent is a log parse operation binding the contract event 0x0c5dd947b790e17c40c62b29772f0e0ca54c86e473b5bf74cfac2a1f9156ee91.
//
// Solidity: event SetUserSpaceEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint8 sizeType, uint64 size, uint8 countType, uint64 count)
func (_Store *StoreFilterer) ParseSetUserSpaceEvent(log types.Log) (*StoreSetUserSpaceEvent, error) {
	event := new(StoreSetUserSpaceEvent)
	if err := _Store.contract.UnpackLog(event, "SetUserSpaceEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreStoreFileEventIterator is returned from FilterStoreFileEvent and is used to iterate over the raw logs and unpacked data for StoreFileEvent events raised by the Store contract.
type StoreStoreFileEventIterator struct {
	Event *StoreStoreFileEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreStoreFileEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreStoreFileEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreStoreFileEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreStoreFileEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreStoreFileEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreStoreFileEvent represents a StoreFileEvent event raised by the Store contract.
type StoreStoreFileEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	FileHash    []byte
	FileSize    uint64
	WalletAddr  common.Address
	Cost        uint64
	IsPlotFile  bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStoreFileEvent is a free log retrieval operation binding the contract event 0xa08fe92f4f83b40431d08f9f130fd22b658ad78cf83027611d629d83427f0d18.
//
// Solidity: event StoreFileEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, uint64 fileSize, address walletAddr, uint64 cost, bool isPlotFile)
func (_Store *StoreFilterer) FilterStoreFileEvent(opts *bind.FilterOpts) (*StoreStoreFileEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "StoreFileEvent")
	if err != nil {
		return nil, err
	}
	return &StoreStoreFileEventIterator{contract: _Store.contract, event: "StoreFileEvent", logs: logs, sub: sub}, nil
}

// WatchStoreFileEvent is a free log subscription operation binding the contract event 0xa08fe92f4f83b40431d08f9f130fd22b658ad78cf83027611d629d83427f0d18.
//
// Solidity: event StoreFileEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, uint64 fileSize, address walletAddr, uint64 cost, bool isPlotFile)
func (_Store *StoreFilterer) WatchStoreFileEvent(opts *bind.WatchOpts, sink chan<- *StoreStoreFileEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "StoreFileEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreStoreFileEvent)
				if err := _Store.contract.UnpackLog(event, "StoreFileEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStoreFileEvent is a log parse operation binding the contract event 0xa08fe92f4f83b40431d08f9f130fd22b658ad78cf83027611d629d83427f0d18.
//
// Solidity: event StoreFileEvent(uint8 eventType, uint256 blockHeight, bytes fileHash, uint64 fileSize, address walletAddr, uint64 cost, bool isPlotFile)
func (_Store *StoreFilterer) ParseStoreFileEvent(log types.Log) (*StoreStoreFileEvent, error) {
	event := new(StoreStoreFileEvent)
	if err := _Store.contract.UnpackLog(event, "StoreFileEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreUnRegisterNodeEventIterator is returned from FilterUnRegisterNodeEvent and is used to iterate over the raw logs and unpacked data for UnRegisterNodeEvent events raised by the Store contract.
type StoreUnRegisterNodeEventIterator struct {
	Event *StoreUnRegisterNodeEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StoreUnRegisterNodeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreUnRegisterNodeEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StoreUnRegisterNodeEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StoreUnRegisterNodeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreUnRegisterNodeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreUnRegisterNodeEvent represents a UnRegisterNodeEvent event raised by the Store contract.
type StoreUnRegisterNodeEvent struct {
	EventType   uint8
	BlockHeight *big.Int
	WalletAddr  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnRegisterNodeEvent is a free log retrieval operation binding the contract event 0x39a789265f7fefbc3aebf3560cea4e1c1f57e6e31faa063f91797173a0daed37.
//
// Solidity: event UnRegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr)
func (_Store *StoreFilterer) FilterUnRegisterNodeEvent(opts *bind.FilterOpts) (*StoreUnRegisterNodeEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "UnRegisterNodeEvent")
	if err != nil {
		return nil, err
	}
	return &StoreUnRegisterNodeEventIterator{contract: _Store.contract, event: "UnRegisterNodeEvent", logs: logs, sub: sub}, nil
}

// WatchUnRegisterNodeEvent is a free log subscription operation binding the contract event 0x39a789265f7fefbc3aebf3560cea4e1c1f57e6e31faa063f91797173a0daed37.
//
// Solidity: event UnRegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr)
func (_Store *StoreFilterer) WatchUnRegisterNodeEvent(opts *bind.WatchOpts, sink chan<- *StoreUnRegisterNodeEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "UnRegisterNodeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreUnRegisterNodeEvent)
				if err := _Store.contract.UnpackLog(event, "UnRegisterNodeEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnRegisterNodeEvent is a log parse operation binding the contract event 0x39a789265f7fefbc3aebf3560cea4e1c1f57e6e31faa063f91797173a0daed37.
//
// Solidity: event UnRegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr)
func (_Store *StoreFilterer) ParseUnRegisterNodeEvent(log types.Log) (*StoreUnRegisterNodeEvent, error) {
	event := new(StoreUnRegisterNodeEvent)
	if err := _Store.contract.UnpackLog(event, "UnRegisterNodeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
