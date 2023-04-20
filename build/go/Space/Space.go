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
	FileProveParam []byte
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"port\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"}],\"name\":\"DNSNodeRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DNSNodeUnReg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"DnsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"FsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structNameInfo\",\"name\":\"newer\",\"type\":\"tuple\"}],\"name\":\"NotifyNameInfoAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"PDPVerifyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nodeAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"oldUserSpace\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"addSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"addBlockCount\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"currentHeight\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"fileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSpace.AddParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"AddUserSpace\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"newUserSpace\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"updatedFiles\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structSpace.AddReturn\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteUserSpace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"Size\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"BlockCount\",\"type\":\"tuple\"}],\"internalType\":\"structUserSpaceParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCost\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structTransferState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUserSpace\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"Size\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"BlockCount\",\"type\":\"tuple\"}],\"internalType\":\"structUserSpaceParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"ManageUserSpace\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"_userSpace\",\"type\":\"tuple\"}],\"name\":\"UpdateUserSpace\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0x936f0bd2\",\"type\":\"bytes32\"}],\"name\":\"c_0x936f0bd2\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"}],\"name\":\"calcSingleValidFeeForFile\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"calcStorageFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"calcStorageFeeForOneNode\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"proveTime\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"}],\"name\":\"calcValidFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"proveTime\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"}],\"name\":\"calcValidFeeForOneNode\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"Size\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"BlockCount\",\"type\":\"tuple\"}],\"internalType\":\"structUserSpaceParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"getUserspaceChange\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"newUserSpace\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"updatedFiles\",\"type\":\"tuple[]\"}],\"internalType\":\"structSpace.ChangeReturn\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractIFile\",\"name\":\"_fs\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"Size\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"BlockCount\",\"type\":\"tuple\"}],\"internalType\":\"structUserSpaceParams\",\"name\":\"userSpaceParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"oldUserspace\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"fileList\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"ops\",\"type\":\"uint64\"}],\"internalType\":\"structSpace.ProcessRevokeParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"processForUserSpaceOneAddOneRevoke\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"userSpace\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"addedAmount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revokedAmount\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"update\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structSpace.ProcessRevokeReturn\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"Size\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"BlockCount\",\"type\":\"tuple\"}],\"internalType\":\"structUserSpaceParams\",\"name\":\"userSpaceParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"oldUserSpace\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"}],\"internalType\":\"structSpace.ProcessParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"processForUserSpaceOperations\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"newUserSpace\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"updatedFiles\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"internalType\":\"structSpace.ProcessReturn\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506201359480620000236000396000f3fe6080604052600436106200010a5760003560e01c80636a4fe5201162000097578063b5bc8ee21162000061578063b5bc8ee214620003c6578063ca6142cb14620003f4578063ed108de91462000438578063f1feacbf1462000458576200010a565b80636a4fe52014620002df5780636c799c131462000315578063722b25b9146200034b578063918a06651462000382576200010a565b80633984ff7d11620000d95780633984ff7d14620001e557806344b2d82d1462000229578063485cc955146200026d57806354e0d3c2146200029b576200010a565b80630f9fa2eb146200010f578063284c6003146200013d57806337bafa5214620001815780633899831a14620001c5575b600080fd5b3480156200011c57600080fd5b506200013b6004803603810190620001359190620114ad565b6200048e565b005b3480156200014a57600080fd5b5062000169600480360381019062000163919062011885565b620009d7565b60405162000178919062012c35565b60405180910390f35b3480156200018e57600080fd5b50620001ad6004803603810190620001a79190620117d2565b62000a9b565b604051620001bc919062012c35565b60405180910390f35b620001e36004803603810190620001dd91906201197a565b62000b59565b005b348015620001f257600080fd5b506200021160048036038101906200020b91906201197a565b620015da565b60405162000220919062012bb5565b60405180910390f35b3480156200023657600080fd5b506200025560048036038101906200024f919062011779565b62001ef7565b60405162000264919062012c35565b60405180910390f35b3480156200027a57600080fd5b5062000299600480360381019062000293919062011654565b62001fa5565b005b348015620002a857600080fd5b50620002c76004803603810190620002c19190620114ad565b620021fc565b604051620002d6919062012c18565b60405180910390f35b620002fd6004803603810190620002f7919062011695565b6200238b565b6040516200030c919062012aea565b60405180910390f35b6200033360048036038101906200032d9190620116da565b620039f9565b60405162000342919062012b6d565b60405180910390f35b6200036960048036038101906200036391906201197a565b62005b4d565b6040516200037992919062012b0e565b60405180910390f35b3480156200038f57600080fd5b50620003ae6004803603810190620003a89190620118de565b62006779565b604051620003bd919062012c35565b60405180910390f35b348015620003d357600080fd5b50620003f26004803603810190620003ec919062011628565b62006837565b005b3480156200040157600080fd5b506200042060048036038101906200041a919062011842565b6200683a565b6040516200042f919062012c35565b60405180910390f35b620004566004803603810190620004509190620114d9565b620068f0565b005b62000476600480360381019062000470919062011707565b62006a61565b60405162000485919062012b91565b60405180910390f35b620004bc7ffc2ea88c75e42b3e159a43394ed7388eb352b308242c7b611c1a40e795a0152060001b62006837565b620004ea7f010e6cb480eb77c5cff88964dd4b27f0112f3f3c13113c5a7a054362ebcb759560001b62006837565b620005187ffd75f1f4cde08d4be5b67961067d8241a45c59562e3dc02b1ac90b2172eb057660001b62006837565b60006200052582620021fc565b9050620005557f270d6d1f8deb3c09c15819f0f82510cef6e2f32df8f625f9a3477c60ec9f21f360001b62006837565b620005837f85dd633638d373d42c5ff6c92744f3d79f11a802c6ebba05912a79fbc104899760001b62006837565b6000816000015167ffffffffffffffff16148015620005b057506000816040015167ffffffffffffffff16115b156200069c57620005e47f883c455fe5f8ad0cbefd197a6fa7c2914bc30987b62ed0c63f7bb6598f8a64cc60001b62006837565b620006127ff4c1a1527e6664c9363144bfc3e08c17300517a315fbc5a39620517d4644850b60001b62006837565b620006407fcae84dda4f1280296a4e18e4373a5449a85efef9282c4b80e687e63955bf3adc60001b62006837565b8173ffffffffffffffffffffffffffffffffffffffff166108fc826040015167ffffffffffffffff169081150290604051600060405180830381858888f1935050505015801562000695573d6000803e3d6000fd5b50620006cb565b620006ca7fe02fcf8e7240f24fcc2358cf7a6deefbe8294f571b6291d7108b4e6758cba6c260001b62006837565b5b620006f97f4303a009938d68287e638cc7d457eabce96a5c764b3f151cdd88f6ab898a3f3660001b62006837565b620007277f77865a25d11566dcafa98efb63d96d332829067c85019082c25cd1f4ade8f5a160001b62006837565b6000816060015111801562000740575043816060015111155b15620009a357620007747f25bbf0ba25a84ca675b9659c76c2bd50a0375fd3331623f460724bbaa0f61bc260001b62006837565b620007a27fe784265d18787a3046e9e6705aca39ca5b534b2612990c22354e9b071af3760760001b62006837565b620007d07f85e719f821aca70644d631d0cebe3fce67efeafe4a5fccdb040003db650a317e60001b62006837565b6000620007de828462007d39565b90506200080e7f214c8d56753b813a2f4bf0270132c5f3998e1f6a00f09f04eaead7312308dc3860001b62006837565b6200083c7fd70ea466184dd0da52a997fadc7881b16cc43f58c329ea5a72a1e7562925f03360001b62006837565b6000815111156200096e57620008757fe0e63329c082b79ab19b859531f97f68a0c886fa4a3e54be804bc6aa16d1720960001b62006837565b620008a37f2bcdf3af517acf778e2791492495f3d01f3f5cf474b75b6e1b7d3dca7c00825160001b62006837565b620008d17fbd228856c955ffd3fe8055ba002f64857950baa5e3af563a68dcb7ea61e7f9b660001b62006837565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f58160405162000902919062012a8f565b60405180910390a1620009387f1a46f5f635bb3b30782afb62d7657c72a5062fde4390919632161bc8a178afd860001b62006837565b620009667fd2f51d240ede680149c856c1e342b6e7037b5c6df338885536e0355c663a781b60001b62006837565b5050620009d4565b6200099c7fb5ba9a520cf4d93b25670f6be8cc1990763a7439939b37c970b4a9007c69bf0a60001b62006837565b50620009d2565b620009d17f981530207dad6471bbc4264d49d44b6400b2cc87236ac54d676c45de0cac66e160001b62006837565b5b505b50565b600062000a077fb291ab7cb0c00d5f7f77c6db230efb928c028f9f9e01b0c83cc9e12a907f595460001b62006837565b62000a357fcc81d82f8d06225d5c8cfa63d35860b48dedc278a3db7a9cee24ce0b947408e360001b62006837565b62000a637fa9807f6f1639b48c26ec842f26c4daa913b3453c97f10cec4b016adfc07c5da760001b62006837565b620fa0008284866020015162000a7a919062012fa5565b62000a86919062012fa5565b62000a92919062012f6d565b90509392505050565b600062000acb7f2fc0db0eeed7cfee12c29bd3c963282756a8afcac5e7c486de01915d6f95554b60001b62006837565b62000af97f77b8f942a2f22687a4584ff5e67b7a30ab9376531a4945b03e18a963834e6d9060001b62006837565b62000b277fec3933607744c9ee36dee30c124a64c276b4a0e89376c10d2649cd79329fdfff60001b62006837565b62000b3485858462001ef7565b60018462000b43919062012ef0565b62000b4f919062012fa5565b9050949350505050565b62000b877fb1e12ef8c0f581db6370ca08f72d08ac30b29a649f8cfdcc996a29b5c34ae2fa60001b62006837565b62000bb57fe98e29747b136de7a12729030d8d9cfbbe6c4df8be94f1d399896727c2c5db4960001b62006837565b62000be37f7dbd1960dbb8ea81fc5722e83822b0b0f5afd695f5d2c904a904fb9d9f4399ae60001b62006837565b60008062000bf18362005b4d565b9150915062000c237fab9b9062b5d158fb0addef2496708a7c3cfb39cdea56d5a2e6e16db5cbf62b8260001b62006837565b62000c517fc340d85d53284671b152356a5e51094bd672b5193d11ad4f8d8cd7f9fac0743760001b62006837565b60008151111562000d835762000c8a7fc97a3ecb728d10b56800b67cde08cd7974eaeeb8d442682c2f99be45f6ea5a6060001b62006837565b62000cb87fd97bfb765407515a62c33fb836cb02f4143dbb521eaf2532d080ba4b9ef96c9160001b62006837565b62000ce67ff856598a7a28f5291f02babdf2c990706941d6ba20e1f6c6945821670b2804eb60001b62006837565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f58160405162000d17919062012a1f565b60405180910390a162000d4d7f061e16043e04cbfba48a30b99193a4bbf0b1479c7c4b6391b0383f9cc32880b760001b62006837565b62000d7b7fd3a34e398025350166a972b80c9eb81d862eb56cebcf3ba72ce6e5b21cab38b160001b62006837565b5050620015d7565b62000db17f530dce0b99c80021d96d6ee5d0704474c3e0dff4a3f4299f9567fbe62e1e4b1d60001b62006837565b62000ddf7f84b31f08790356fa8f8850671f5d3bd7e137968ea358dfb43ea981db25a90f5860001b62006837565b62000e0d7fb92076a51a548be837397ad2266b0afffb320f2269b7a720de9f910634a3474960001b62006837565b600082602001516040015167ffffffffffffffff161115620011e15762000e577f57fd0cc122eab129693c4ed38b130baf1578bfcbf0704fc1d88a503ed97459e460001b62006837565b62000e857f7f80f1f1c48d9b7801597461ba50d67f75498a9726967f80be217fb775ebb59b60001b62006837565b62000eb37f88fb5a6b8509f5e567ca0b01afe770eba7452b32b6f1e344482cb6e5afe92ee960001b62006837565b3073ffffffffffffffffffffffffffffffffffffffff1682602001516000015173ffffffffffffffffffffffffffffffffffffffff16141562000fe25762000f1e7f3dadcbb265ae9c7ce3ca53203ef5e60b9606f2fc3c6c6436c6d9a0e20f5908df60001b62006837565b62000f4c7f1049c3f48d17ffc10adb2c2fa09765043c1aab37af3bcfea0f0b539fb466f1c560001b62006837565b62000f7a7fb54613a2f7684227ee2e6f5d64371fe5d9ce259ddc0b7a0d0398cf08f668f42f60001b62006837565b81602001516020015173ffffffffffffffffffffffffffffffffffffffff166108fc83602001516040015167ffffffffffffffff169081150290604051600060405180830381858888f1935050505015801562000fdb573d6000803e3d6000fd5b50620011db565b620010107fccdcfdfff2eb415727b937d1ff73c8acc649cd1355277c7efbb5e16e27c4f71f60001b62006837565b6200103e7fa2bcd7081c7065083c4261539bcfa9104abdfc5d5fd86fbcc9e98507307cef8960001b62006837565b6200106c7ffb98751c57ef1b121d0b9849c9fbec146c46ab8eefc619bfc5ccaa25a990a0d660001b62006837565b81602001516040015167ffffffffffffffff16341015620011ac57620010b57f3b4a17ea9a5d2493676108907e4361938cc731861c0d3e95f9f581801a64efdd60001b62006837565b620010e37f0b3c5cede03e94bc8afff6b694a402374c23736ac7b818b888a658dddee240c060001b62006837565b620011117ffcc947b0957d7eb4a55667a174d1cf9d3e313eec5107a5033f41a1cabaa6175760001b62006837565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5604051620011409062012a58565b60405180910390a1620011767fccd5f448ff346299f1ebdf41218c3ee697413532a5d3b43008caa2db7602714a60001b62006837565b620011a47f1a64abde407e3304bf5ae864d88f87a9c207239c33d2a5cf8f9aec76b9054e3a60001b62006837565b5050620015d7565b620011da7ff0aa01c0feee1a04ab6c92a365b8ff1c4dc46a2bf0402bd6220d8f8265219ea860001b62006837565b5b62001210565b6200120f7fd06e884eb9ad7633e358f0365964945de367d6448854a1bc50bfc343bb519e9460001b62006837565b5b6200123e7f9e966354315f0d5b05e201c440859ecb79a9a4f38c132957ee0c178ba382a3a260001b62006837565b6200126c7f1097348ee20259988d51189de989da9f9d4b3d46ebe7e22e5d1e5762a67fc63660001b62006837565b60005b826040015151811015620013c557620012ab7f07047f11f2466737e4f81da0f81e017b17294ed82b65679cd019fa99c3d0f1be60001b62006837565b620012d97fe31dac448a36c89ef45348817a23cd0aa7de19aa7fa1c3822cbe8358003239ec60001b62006837565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663681850d78460400151838151811062001355577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b81526004016200137b919062012b49565b600060405180830381600087803b1580156200139657600080fd5b505af1158015620013ab573d6000803e3d6000fd5b505050508080620013bc9062013213565b9150506200126f565b50620013f47f4e72b97173ff444de03a7b3c7316064b20979e7386c7c8e4525160bbc661709660001b62006837565b620014227f61137bc99cbdf7e2931cd9405299bcaae3d54c4784df3ea5ed88227c8d63b5d960001b62006837565b816000015160026000856020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010155608082015181600201559050506200153c7f9101951560db0ad94b121c6d4e0f4cbcae6ea3ed4993de2d17a727d7067894a560001b62006837565b6200156a7fb82545513b4d55e182527cf3d48da8619a01cd2178e0b24144bf95b492ba0e2960001b62006837565b7f0c5dd947b790e17c40c62b29772f0e0ca54c86e473b5bf74cfac2a1f9156ee916003438560000151866040015160000151876040015160200151886060015160000151896060015160200151604051620015cc9796959493929190620129a2565b60405180910390a150505b50565b620015e4620100eb565b620016127ff36542fededf1c69619821af2873c95c892c125404ddfc6649c5c3c3df32831160001b62006837565b620016407f40b982557fdcb7283268200cbbffa18d97fd2a8a15608bfa5c168c5b748742b160001b62006837565b6200166e7f601c72c1e17992e78fd5f6396c794c031726c0a84f0a1860699d5070bc96d2c960001b62006837565b62001678620100eb565b620016a67fedd7a248d112394ecb2a7cd12432e7b24ffd1108111a7351d0659476230f770660001b62006837565b620016d47fe96b2e90f4e616a524afc7edcd6fe5c2a0b6a17b48be91553e4527cc840a5e3260001b62006837565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156200173f57600080fd5b505afa15801562001754573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200177a91906201174c565b9050620017aa7f6df4dd4bf9bf930bd772485e7aeed7cfaa9417140965dcaadc3389ff4429c10160001b62006837565b620017d87f613078797348076d263420ef808e31c4f782768e5363eb944b9bfcc33730316260001b62006837565b6000620017e58562008b98565b9050620018157fda3dfc04b30f370b7556820d80442ad64836aa6ee991eba95d3d2f90f3c25e0e60001b62006837565b620018437fef43d695fb35018e6046b7047a89b854a2ddb56c70f83be15a42ab21e524dceb60001b62006837565b80620018df57620018777f6323a4964b769c93b48a8b5c3cbf05b45fdd75af84aa11d20726f07edcf66fc060001b62006837565b620018a57f1e55f09a08ed9f7e222a7f0f82e4053ca560895cdeeb974cadf2e7b21441180e60001b62006837565b620018d37f9dc78d816619b5d21c311d39bb084600d0e5763707a6ce19ca7889cdfd6f3b8f60001b62006837565b82935050505062001ef2565b6200190d7ff81c435f23db222575bc55ad76fedb7c48e98f19185d565f4d4afadb52d2921660001b62006837565b6200193b7f96ec3810af5ee8fce7fcd87b8999993ae961bbb7600dc8b4d2883a79aa9e6d4060001b62006837565b620019697fec313e4dda6fee4c24462214ed0c12da9fe831165dec6cac8851ee72149b754f60001b62006837565b600060026000876020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201548152602001600282015481525050905062001a937f16222d3636f6113e2c4237cfc80e9596a79de67758f64b97f7cdf4862e18c51060001b62006837565b62001ac17f051ad4fd9e8a1ce8f87cc7def2368b1e4b73587325a24473f5d46eefeea2f66360001b62006837565b600081606001511415801562001adb575043816060015111155b1562001b7e5762001b0f7f0dab5ecbdab6d50b4a728795aedae7744a495c1709aa05bc361ebe37f9d069a560001b62006837565b62001b3d7fcbd84d5b43ce278a3bb9793c0bfb0b8db8111ccba69ab7e6c0accf110e97200560001b62006837565b62001b6b7f2550850958a9eed905534e582cbd46aed30f9a23bb50dcfeb6ee38e6ac059d2360001b62006837565b62001b77814362009478565b5062001bad565b62001bac7f6f0308e87117b21e52f5ed2100d3540651db65ec9be374d7ad15164d16e7002160001b62006837565b5b62001bdb7fc8c45c6f7a96ad57833d50e31e201af73a26846bd88091669797278ed6a929f760001b62006837565b62001c097f79df323ac2ebb17f23b7723fb705e5e656e26de38d34899bbd8e02ab7db3e24d60001b62006837565b6000816060015114801562001c215750438160600151145b1562001df05762001c557fb9eb1ec70caa0d31efe4721534446d953849772e63574f3139d7b48a1a35b4dc60001b62006837565b62001c837f9eea07d6e587fee6ba17f9e3f3a61ad091c51096f1e3b2f6bc90d66e6926000f60001b62006837565b62001cb17fb306c791964bb9209437a8b3fb55f5330189d5254ff1bc494b7eb430a5662e8c60001b62006837565b600062001cbf8488620096db565b905062001cef7f49e115935277d843bd205ae87dc4da67a9a5d9df68bc9dfcaf0cebeded5762d860001b62006837565b62001d1d7fc20c33cb25eba21bc66eb981c1e623ca344971980599bf8a4f42e9f420b2866260001b62006837565b8062001dbb5762001d517f670251cfc16b9ea58fa866c4be650f2b5be7e5edb93a07c6f129fbd51a9c863560001b62006837565b62001d7f7fb7f858523fc7487c9b63fad1f0b9a190e8e03614aaa784d12ca8c6cc74ff0a7560001b62006837565b62001dad7f8ff585a50783447eb2dcd33d41343127ab675b9b275562a8ecb6ef5863e672f460001b62006837565b849550505050505062001ef2565b62001de97fb143d80bb595060aa8de05fefd1b19aab89f5f9af9e72b5794b0f8032af4d80460001b62006837565b5062001e1f565b62001e1e7f5f3a1847c75a920591712ae3064fcefb1d1795df12615bce42de40a2e335bf9160001b62006837565b5b62001e4d7fe39e8c7f4a7f064be0cb07a2fcb988f6f2e2f7d0af593527544ab6ff4dd3c79b60001b62006837565b62001e7b7ff864df909dc71b75efa13cf12583622fc0d58f27dff6d58cd73801ab33d1df5460001b62006837565b62001e8886828562009d6b565b90508094505062001ebc7f0649b54e86a7e22648ae4f6042a704ce36ef04a1fef72a3834a03390d3ba610760001b62006837565b62001eea7f66d0519de3ab51fd91093900b66a592a2999f70d0b2a741c14c60e27d349034e60001b62006837565b839450505050505b919050565b600062001f277f66bbcd9dc94e7bf2be4b32b903b69cf0085cdcbdc75e0f7799f57ded872a88e560001b62006837565b62001f557f5d24838b34df215a4cf6465a91cfcea8ad872d54646ef552ef905184eea1008760001b62006837565b62001f837f8532a213f030fc77b6ee81c5887110cb32fd629d20b396a741168a63d6d5fdb960001b62006837565b62001f8f84836200683a565b8362001f9c919062012fa5565b90509392505050565b600060019054906101000a900460ff1662001fcf5760008054906101000a900460ff161562001fda565b62001fd96200af3c565b5b6200201c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620020139062012ac8565b60405180910390fd5b60008060019054906101000a900460ff1615905080156200206d576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b6200209b7f806debfe1ae8c78ea8dfc5ec63ea1cab17dabf25fa5bbb3c40af78d7cd99004d60001b62006837565b620020c97fb69c6427e4810e3fb7c167dffacd6a592a4c44cb2aec585ddab91b09c132607560001b62006837565b620020f77f53e6e83f8ef7c7ef97029afa089c5fb1926af25d63e3014243eb616e62510fac60001b62006837565b82600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620021667f5862be41fd6cd3f6d365059b0f80d3a11a89fc3a3882b6e4f5e740c1d21e9ba360001b62006837565b620021947f398985efeee4838b25a3221ad93e136d5d50f6c05c090cfd19293fbabb8f72cd60001b62006837565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015620021f75760008060016101000a81548160ff0219169083151502179055505b505050565b6200220662010142565b620022347f9d88c25a4613a24d1839172ea061d062e25e6725fe5de171b0d29a61ecc6a44f60001b62006837565b620022627ffa7ed8fba48174d0161ee9f1fce6f590311f78083454a4e1f90fe03e2f40115560001b62006837565b620022907f50ece3b800a0513e860cf67e7b068733276516f30f6f4a4fae5ad49c540cf11860001b62006837565b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182015481526020016002820154815250509050919050565b620023956201018f565b620023c37f94901f06318d36a4d34c2398b11ecd225576d8c2d8c5f6053442d4360874952060001b62006837565b620023f17f2c30de56d6688de66e02d99c8bdbd7b560587ccabd041b6d614431b62c0a26e060001b62006837565b6200241f7fce70d77a694b0145e4d43f9be9b08c3cd61c944e4d6bb9bd2198d0eccf60a83660001b62006837565b620024296201018f565b620024577f08a26d85264d178718fc6400620620580cef4a2ecadf642427aa07d19e4ec4b960001b62006837565b620024857fdeb89aaeea19839b12c696c3c2776b1f50e185cf5f2510e5024f9a9111f1aa1e60001b62006837565b60008360000151606001511415620028c357620024c57f5d642fe317d611248ba3f84fa394e71556414ce436feffd005e95702bee3b22a60001b62006837565b620024f37f057df35bbcda83c375fa2bac2fd04b77fd6426b7b0868d274f9ee2908315e39960001b62006837565b620025217f8eaed6ea1f7c6b08db0e728f79508767b24f28cb580061c67b2581360912f3df60001b62006837565b600081600001516000019067ffffffffffffffff16908167ffffffffffffffff1681525050620025747fdcd81a72fe24bd53aaeba817699bee7efd6cec90d4cdcf0e127e0e4310c7563760001b62006837565b620025a27f3d9d319ac0eed2402b2f85f8db996a7bf61c8df7ff98d498d90e2cf6f6c0b94760001b62006837565b826020015181600001516020019067ffffffffffffffff16908167ffffffffffffffff1681525050620025f87f3207ab229d75fa3bc2805f2c570135703d29d7026efb0944cbc068a4b54a28a660001b62006837565b620026267fa55f2cfdf8cbda40d7fff1816342577e46760e1f13f4d31ad33c94a961da5fe660001b62006837565b826040015167ffffffffffffffff16836060015162002646919062012e93565b81600001516060018181525050620026817f443a8a1c805ee44c21a72cccf59e20f4e66c7160e393fff7d07b60f681baa37c60001b62006837565b620026af7fa937a29c7dcdaa9817f58c543a1581968ad1bb02f7855fb13f2c018cd577637860001b62006837565b600081600001516040019067ffffffffffffffff16908167ffffffffffffffff1681525050620027027f20bcedae92fb60e7bbcf22a961fc31f14804420cb0881f0357addb652c7061bb60001b62006837565b620027307fc90d32c8c73fb9c13d6e069f687921911be42c4c89b6b0d55c499cb963210a3960001b62006837565b60006200274782600001518560800151436200af4f565b9050620027777fbb5818a8885cbfa163674dc42767102a8ab123a015646fca6c72eb288db0ca2e60001b62006837565b620027a57ff4dd40fc786bf2e9172376176f8934527dd7e11306a35994c8ae03cbcabd3f4e60001b62006837565b806040015181602001518260000151620027c0919062012ef0565b620027cc919062012ef0565b82600001516040019067ffffffffffffffff16908167ffffffffffffffff16815250506200281d7f386428c0bb057094b76f46505aa26b6095885c7227dd98f4ac6413fcd98aac5160001b62006837565b6200284b7fb0aac12b1cc1cb0936f17116232d96d1a90cd780641b9f6449af0de261bf9ceb60001b62006837565b60018260400190151590811515815250506200288a7f45da9c4a065d6b4c5db1746a15a62fb3730c1cdb675b44683657d19dbc96182d60001b62006837565b620028b87f49f1ffbea42c6505fc2289a6a6dcd8cf5475ba173fbab48651b7b83413864b2760001b62006837565b8192505050620039f4565b620028f17f8d27476cf43866733f0e716be20f845e2877d0d2d0d55b61d0b5a28d01a0ac0a60001b62006837565b6200291f7f8fdeaf89c9c5e2442dc0f865c113c0e92bb16c4dea718c948df4f9efc6a3d8c160001b62006837565b6200294d7f9c69e9fc32599af0d0918b66a7fc206e9713644144bc39f71e14ae9a1223f5fa60001b62006837565b6000836040015167ffffffffffffffff1684600001516060015162002973919062012e93565b9050620029a37f3a9fb206e7134fa3d2877e3a496eef5ea8044734159b8e8c2688ad06a1b5909560001b62006837565b620029d17f1c23e92d885edc68bdbfda1860420b111274eede8659c9ea5701b6b5a6a72e3460001b62006837565b60008460200151856000015160200151620029ed919062012ef0565b905062002a1d7f52450bea36d7b565cf5bdb8114dc7541fac5a741474a37e677b150e33a61b73560001b62006837565b62002a4b7fd8781d1e251b23cc0269304fc86cce2d6d6055bbb440f8822cdcac2eab36613260001b62006837565b84600001516000015183600001516000019067ffffffffffffffff16908167ffffffffffffffff168152505062002aa57f2c941fea80867fdcbd6b0c658b7af10a10c5b893e503d427144147e37a8e5cab60001b62006837565b62002ad37f6d3a60c3340afd12a5a3d42df7b66ae7df3f3cfeb9d3ae82be4e99e16ee8e05260001b62006837565b8083600001516020019067ffffffffffffffff16908167ffffffffffffffff168152505062002b257f0abce7d3586db8b2dfd62d517f52f7e7c01be12c9591260ba0cb0c3067953ee760001b62006837565b62002b537fe7f8663b49966dc9e0f85ccbc77a836b392bfde13d773f14d063d1aea472cfff60001b62006837565b818360000151606001818152505062002b8f7fa1b6beebe6bc9c9a4d170004c33777b2da5a5bd7332a3c66feb6bd275be8218f60001b62006837565b62002bbd7f79698355fd93d8f1a05688fd1f4f2cd0ad9a83d2e6f1f949b3173208a76961dc60001b62006837565b84600001516040015183600001516040019067ffffffffffffffff16908167ffffffffffffffff168152505062002c177fad16b5cee6b4bdc5e9bcd31ea2990257ef33f7c8f05647f35496cd38d50b7b5660001b62006837565b62002c457f5aa1d0b7f3b5e464184678283ff7ad60984fccbf1061e108c32bd9f12f235bdd60001b62006837565b600062002c5c86600001518760800151436200af4f565b905062002c8c7fb7fef740c334ddc9598bd517b3817e8082b1710ab37886e59697f9d58ae5c73660001b62006837565b62002cba7fbb5034361a5a91c23a7e2c990612fc57e991123cc363470810b6dce946c95fb360001b62006837565b600062002cd185600001518860800151436200af4f565b905062002d017f2dc879ab1cdb203cd8a80d4b01f0f256fe1eaf15dbfd5a6ea98038b235673e5e60001b62006837565b62002d2f7f7d8fe82e44b6c0fcafc1159004e8cb6150f39d11b403e4e3ea243529daec514f60001b62006837565b600082604001518360200151846000015162002d4c919062012ef0565b62002d58919062012ef0565b905062002d887fe3819ce057c17c0a2333c8d56cb62d4d63a3bca3ce112065a751e1d76daefb4560001b62006837565b62002db67f7c6f71e016a26fbff14b2135c87d06eaef55fd800252b9756738ec505982b6da60001b62006837565b600082604001518360200151846000015162002dd3919062012ef0565b62002ddf919062012ef0565b905062002e0f7f4829dc1588f886c1c13f7ee1ddcb4bbf2dd1886385711ee15074f627575ecf6860001b62006837565b62002e3d7f0bc77aa1fe5cdc15a704c4adf4543d82429773a29bccff3fa19728171594977760001b62006837565b8167ffffffffffffffff168167ffffffffffffffff161162002f605762002e877f325dd7621bf2c8e5385ef66b149e30aeabad8d20ad920b9540e12a7fff623b6560001b62006837565b62002eb57f5b78c7a742769f033e23e8ce7370115b3e2ce1d6f4a717ff7d0ef2cb2463be9c60001b62006837565b62002ee37f36b2747b5e41397242f6e020d536cfab526fa243ce4d074c300a28327e9e4b5260001b62006837565b600087604001901515908115158152505062002f227fc5ae3c93a711136d11e358fd96b513259651b08b7d90d8073caddab1d946907160001b62006837565b62002f507f81fdc1c84d59b59925d6e64e5504ce926c41ce32d30e7176e935bb53233398f060001b62006837565b86975050505050505050620039f4565b62002f8e7f86322c5d34f2d03b4c572a6d73b2ed877170df2c0f456f999852521e5ca0c3bb60001b62006837565b62002fbc7fb8dc64039f0e5c9f98be58b63e77cedd5f11bdd553e26f4bdf00615d48df4a0060001b62006837565b62002fea7f8ecc7df171077fb642f29971f43ced1aea9ff59177a41a610a10401375bc86ad60001b62006837565b6000828262002ffa919062013029565b90506200302a7fc4448c46973e01ffd33969fe8b968418d75f4378d6e338368909da3874e6b52f60001b62006837565b620030587fe578963647b681c1f03809ac13d7b7ad077445b839db70a7de488180d0a8deaf60001b62006837565b60008a600001516000019067ffffffffffffffff16908167ffffffffffffffff1681525050620030ab7fcfce77dd49bdc00a25e5f1943925ab1904783e6550b90bfe0b4e244bb051925f60001b62006837565b620030d97f5d9f49ebd68644aef24a8a3cdefba1af756e7df8336157a3778512725beab62d60001b62006837565b6000620030f08b600001518c60800151436200af4f565b9050620031207fa7445bfe7e88aeea7702ac82ffd1fdd26d70a810517b495d2198b93af277122d60001b62006837565b6200314e7f68f7234e7f5af6674a9b377cff0a13afc21cacbf2354e75288b668c51206382460001b62006837565b60008160400151826020015183600001516200316b919062012ef0565b62003177919062012ef0565b9050620031a77f5af2cec3d94dfebf03ebd70207d55d3fafa57bb25932e2cc7ba43f300d49a9d560001b62006837565b620031d57f67952e9d3131762bb952598b1f62e867c7b0409db0e458304a3fed15d9d217b860001b62006837565b8067ffffffffffffffff168c600001516040015167ffffffffffffffff16116200330357620032277feced80709ac27ab3f6a3abc9762e8d904eb8a616a3c826c9d988d5c10d3b349560001b62006837565b620032557f26d7d88f6ead4b6d68736530551c5dafa267ff2b2809cacac56f941e41e0da9b60001b62006837565b620032837f2cca043c0d710b8f6b80ee84e81cbe3aa8c991864febba3ec1c1b240fb3a0b1b60001b62006837565b60008a6040019015159081151581525050620032c27fa11c77ea9836a53301f35bb9edfaecbe044bf5a7f9e1e50cc39b10c93af5cfce60001b62006837565b620032f07f61c92cd783d025966af5027e4dee13e34f9bff4437cd0cf3f31b71ca7b741f1060001b62006837565b899a5050505050505050505050620039f4565b620033317f77a98082621d9d004296b2a25336ce691595cba880f86dc86a92d1a3360fefcf60001b62006837565b6200335f7f356b54a38199f6a98551d93224487466c0e86b24fe1e83c98da123abcb9d8c0d60001b62006837565b6200338d7f97140c0f765bb3b3372e3906b8a2f510ca29a1b338834b62f113b5d3ff12ccc760001b62006837565b6000818d6000015160400151620033a5919062013029565b9050620033d57f384e71c1e0f08e9b1743836a1f5046758dfbe98750f28477bad4dea85f5e349060001b62006837565b620034037f5d73b339d95cff5bcbab710c932ec1e7b740b47ce3b8b65eb867df8af2421d7d60001b62006837565b8367ffffffffffffffff168167ffffffffffffffff161115620034b3576200344d7e290292dd07a4803f57580c00064e55d047766179d4be9cde3aa18211af287660001b62006837565b6200347b7f6f45e478c1c1134e64248d61fc6c4f8cd6ac24f0fc37be0e8b63ca86333b2d7f60001b62006837565b620034a97f20a0f41b69b4c35f7eec99d72bc822e9a4cab296c46b19622396bd277182b1a760001b62006837565b60009350620035df565b620034e17f8abe6e2ae6be176d77ddf20962e09ea8a5838b6dfb88dbc22afe9c2f7c169aea60001b62006837565b6200350f7ff137e1c00a5b91f2e1f337a6442df292932df1ae2242d4cdd43ab1cc4c53c0ec60001b62006837565b6200353d7f474872b098f61482c3a7036844bf6d87efae4db9cf1f60ad93889f390f5a66a060001b62006837565b80846200354b919062013029565b93506200357b7ffcc4175632518bdf10e41b5a1c60e39056d3edc357fe7cb6b94a98216a8d16ef60001b62006837565b620035a97f137c6d4db9a0c0bcb41a2b463bdf3e1c13bb0b1aa828e08976e538d485931b4860001b62006837565b838b60000151604001818151620035c1919062012ef0565b91509067ffffffffffffffff16908167ffffffffffffffff16815250505b6200360d7f0dfd7c55a05299765ce4648b5609088526ee0fc6e11e8bcc7f56f3115edbb0ca60001b62006837565b6200363b7f648c01e73586286c69dc0efaa94d450dd93990f3c06ae05d9830c76b3c03d3c860001b62006837565b60008d6040015167ffffffffffffffff1614620038ed57620036807fc0b21eaf7fbfade476a2265e988524535d6f230361090ab8a67c29a4246e8c5860001b62006837565b620036ae7f49502ca55d2e3122d98f9f7e5db816804f4c139cf7ccc5d707b02b738416078760001b62006837565b620036dc7f840b841584ac2026200ee71b968b1b140a0793823a3f9d9ac8850dd453ce92b960001b62006837565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633f2cc9a08e60a001518f608001518d6040518463ffffffff1660e01b815260040162003745939291906201295c565b600060405180830381600087803b1580156200376057600080fd5b505af115801562003775573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190620037a09190620115ce565b8c6020018d60400182151515158152508290525050620037e37f7308286d66cd99b1752bb5829b1dec83d40d2cfe1fc482530856649d1e085a5360001b62006837565b620038117f94e3cfcae53b2338cdc53333195a662688bdf7f0aa1784f2400362a146a67c4360001b62006837565b8a60400151620038b957620038497fc7fcba0c01cc6c4b02f824bf216bfd2a0c01f82b6e7b39cd08eea752fc6b30b360001b62006837565b620038777f7ede18f855920155642a6f9f694c7e3a35820cbe1e3b36c963b687f85ae3ddfd60001b62006837565b620038a57f704b4c18cc54d1eec165a50be627767569148358ebd4166ee7ea1cca0d8aa29660001b62006837565b8a9b505050505050505050505050620039f4565b620038e77f962090a735484b655361d07e23d86c218e5c6e239640ad7c3db1b98543c5fd3e60001b62006837565b6200391c565b6200391b7f6bd044f65017c910bd021b5b6b9d43e6691b6def55c0850594929297230562cb60001b62006837565b5b6200394a7f82e3ad0f531332a041cb321e5dd15c9d4e9f20e68f28fc203ddf6c087f0fc03f60001b62006837565b620039787fffa37c1a1b319ac1fd88aeef7caea7d47ccb770bc0ff7d8fd409535871d87aad60001b62006837565b60018b6040019015159081151581525050620039b77f931c83021e4d0e3402a7fe328f4d259698c11f04a1d0c926e384053ec06c5f8460001b62006837565b620039e57f2c65a0543c813d857947c040e60ea0be186bdeba45440e3a1d861e3ad16cac9f60001b62006837565b8a9b5050505050505050505050505b919050565b62003a03620101ba565b62003a317fa26a7ed4d4d260af308700c8bc5530898c57a5b95b2422721c3ac094e7eb46d560001b62006837565b62003a5f7f27eb79f06e870ef1107e5bceefe5e35bc5e94282dc1e4a8de5f60bc2c8fbfa0960001b62006837565b62003a8d7f5fb80fea2a30bc577368429d38b60ff47b8c76bceadd690da40edf52a9478eb960001b62006837565b62003a97620101ba565b62003ac57fc0a16611b403f0e311ce1ee49ce8a6eae3ebf9d2d583f5cab38d62ed4efff71b60001b62006837565b62003af37f17973e0c5b96d9221216aaf772ecb23bb127db749615121a2afe0206de2f1f3160001b62006837565b600062003b237f03efb4be2780d78f901f936483b77acb3819a1411898982e88276d3816e942fa60001b62006837565b62003b517f107562484378d922357b7ee75f0c4a500059e460fe32587b4d88d019150ed33060001b62006837565b600062003b817fab2718a1b994386d393ab4915042da76e9a64224ef47e0ab521fdac0e0c2a5b460001b62006837565b62003baf7fa87c193cac5b676d9a0c3da7c5d665585ecbcf52398966841953fd77f15581cb60001b62006837565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639a2b5e638760000151602001516040518263ffffffff1660e01b815260040162003c169190620128c0565b60006040518083038186803b15801562003c2f57600080fd5b505afa15801562003c44573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019062003c6f91906201151a565b905062003c9f7f8247404e6abfd47b3b22beebb45446bdf392ac632795cee83287a61b8468e01b60001b62006837565b62003ccd7f2d06be3336712f5caeda3742677a87241c9ad41333c4d381e0c86de33e5d671060001b62006837565b60008062003cdf88600001516200b366565b9150915062003d117f8e7236a9cfc96717dfadc32650bbf4f17a0135cafc5e1bec6b45d734648241f360001b62006837565b62003d3f7f3f104f3d9242bac34182507423cf7723c7630e6c6cb316d28ceab2e828a55bfe60001b62006837565b60008151111562003e505762003d787f37e08c54517514b8827bdf414b8d84b7124304518c87fca14a080db4ab24992060001b62006837565b62003da67fbdf0b1cf918665f28cefcb8b8dce886e677ce8b69c18752e7339ed0e3741a5d060001b62006837565b62003dd47f9d9c8735243ad2296006930bbfac730af58f6d7e538fff5a83d10f304c165eb760001b62006837565b600086606001901515908115158152505062003e137f0b3deedcdfe894abbf6813b75712f876a49c9f70bc019c535ddc8cee9b477b1160001b62006837565b62003e417f19ed4b4509176e8d7c850495a8f7b301f94bdb055b089a0040d3145203c1392f60001b62006837565b85965050505050505062005b48565b62003e7e7fe45b13998f85f2f32d2168dab0c872421e7aa84d3f709fc7de83b960a57f43f460001b62006837565b62003eac7f506dd673bf63588d26f6748d2f2db63bb55a6214d6f90daea5fe3b00fcf0d46060001b62006837565b62003eda7fdc3676bd4a24ff81cab1e5af1c5b7621bd908d359802bc8006cf8db5cb959c2c60001b62006837565b6001600281111562003f15577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60046001600281111562003f52577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168267ffffffffffffffff1614806200401057506000600281111562003fb3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60046001600281111562003ff0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168267ffffffffffffffff16145b80620040af57506001600281111562004052577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004600060028111156200408f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168267ffffffffffffffff16145b156200462d57620040e37f136c37321bc25936bfe7200d9c42697dba8f8b3196347d0e68f0b0f5d35913ad60001b62006837565b620041117f157802733df6e5255549b67d846cab1c01e41c711147a63a47551714fca6c58b60001b62006837565b6200413f7f785027ffde68e94581430309174cee6841807b690389d74daf8727b2fef8528760001b62006837565b62004149620101f4565b620041777fcfb9306404c499a42baeba28529f9f0f2c867be1fbe38b3e0404939f4fbde5a560001b62006837565b620041a57f3184074958ea1445e597a1ec00cfb7ef69595eb6a0e56cda051dbe3504978b5060001b62006837565b88602001518160000181905250620041e07f4df77aa9115d8916d731bcd3d775f09fe883e7107e5cb52f841b98a8d11226e860001b62006837565b6200420e7fb7e5d25ee874958085722a663f9235dbd554c74a59040371ba4108ef9ca5035560001b62006837565b88600001516040015160200151816020019067ffffffffffffffff16908167ffffffffffffffff1681525050620042687f7ee98045cba9a758f390c16a7718898c0543a9301a167e275638f877b2ac0ba560001b62006837565b620042967fdee58df9c0081bed72975f2f43465e32e052b2080e4a27337a36a81a8bda719260001b62006837565b88600001516060015160200151816040019067ffffffffffffffff16908167ffffffffffffffff1681525050620042f07f664a607fc3a59825c72f26f7808443737f005811b02b89fe61db4c9374d8102360001b62006837565b6200431e7f53a01afdbbecc4fb802a6ca2a0b566fae226d9c2c74ac48fc57c241105df29c260001b62006837565b43816060018181525050620043567f3334114fe7504e814c3c074c9536eb44d93644950287b41c83ecb50c7b74e5f560001b62006837565b620043847f2a3973a6cbf8bc4314195e950a44d9a09d60754f748f09cbc433dc6e0c244bca60001b62006837565b88604001518160800181905250620043bf7f3e111004bdbb073c355b20dea5d5c648c0c8a33088eb95f343f6f8ae089f98bc60001b62006837565b620043ed7f7d324280054f05e79e21f3237533347683572bf3c6c246252f0eeed9d2a8977060001b62006837565b838160a00181905250620044247fd966894ea605f18664e6027a30ffa64e3d5d0d42571999c81b228ddd62b5b40460001b62006837565b620044527fca09e42ad29193860c49c3e03a2147eb89c8f71e0c0e84c0e739656d191c278d60001b62006837565b60006200445f826200238b565b90506200448f7feafc5f48278602e1662f25a13be1144d9e2f2be95b2236dc50308c2f6eaf710f60001b62006837565b620044bd7f69d64561e6d585bb9e236954c49886c9a6fdaf7dd477089968d48116714963f560001b62006837565b8060000151816000015160400151826020015183604001518b6000018c6040018d6060018315151515815250839052839b5084905250505050620045247f4f0e6a8bdd3d5aa555b17738a7c97748eebd0c5547918fb644bf0e0ddc1510fa60001b62006837565b620045527fa21281e193fd7cfa7eb74c5fa6166a64f04fcce15c34f2cd4d3f0c05535ee12960001b62006837565b8760600151620045f7576200458a7f64c84063ea062f2af5d7cda189b5f733ea5186f1919dd2a751642a856c2b6c1f60001b62006837565b620045b87f34ce5992cf2030b250131654fe0512ba09696806cf7ea66820961c09abdd7fec60001b62006837565b620045e67f55e8028af7707c714d82e75172dab8bf7bbc4e6e785cdb0c65075e9e7f737bca60001b62006837565b879850505050505050505062005b48565b620046257fd6940e6758de479346c9aa48911e4a8937c25220ceed190096abf1077a7d9a5060001b62006837565b50506200465c565b6200465b7f7436afcd264a437dcd1ed82ec30e94dd3408de89b12f4bf44a112ac6be75fbbd60001b62006837565b5b6200468a7fc0b7774d3487d5a14ef985da514a36f751ea47af6454865fed85971915dab0f960001b62006837565b620046b87f1250c9da3140e378e792b1d7afe30955deb460fb6300d140f875378e29da6c9b60001b62006837565b600280811115620046f2577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60046002808111156200472e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168267ffffffffffffffff161480620047eb57506002808111156200478e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600460006002811115620047cb577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168267ffffffffffffffff16145b80620048895750600060028111156200482d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600460028081111562004869577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168267ffffffffffffffff16145b1562004d9157620048bd7f4ccee1a0dac1eb0002e228893eea31b2c8fa4645a7cc5e2e1f03d8dbc95fd82760001b62006837565b620048eb7fbcbf8d1ac584fb869d315cb207656552e117171f0af432e11cce9bbeeb17324a60001b62006837565b620049197f638849b80fac8d75a1681902927fa969071f472b8f82261f6aea8d0aabe8dc8e60001b62006837565b620049236201024e565b620049517fccc5dade75599eac82e6f1608fe69f76e2230bffc9123505b4597c0ac710170160001b62006837565b6200497f7f2228f191882cf2adedcca1826249c9d529ec605315b9f062111eeef4c2b7641160001b62006837565b88602001518160000181905250620049ba7f543c1d0591d5e47b81f1b7894f375ecb91142ef79cd436861164e1daf3a811ef60001b62006837565b620049e87ffbf296496f6e341ad44fc298407f63321e2b8575e7154de4860aac62d823ff1760001b62006837565b88600001516040015160200151816020019067ffffffffffffffff16908167ffffffffffffffff168152505062004a427fa2f58d075414024c2331923da0e9d72899ac60dbff17487f7b754204d19f4a0b60001b62006837565b62004a707fc8e4db9c1d14fe061f7c3f8c8a6cc3084bdff8515e874749d4aafeddf571adf860001b62006837565b88600001516060015160200151816040019067ffffffffffffffff16908167ffffffffffffffff168152505062004aca7f29273161b70229feeec9981bfee5339b535590eee8afc0c86ca518506df63c8460001b62006837565b62004af87fbce01222fdbdf92011e1671fa05f0ed9c7ec67c4f86e6dc7e62442438a42630060001b62006837565b4381606001818152505062004b307f3f4afcb0794a23579b5c9780477acbfcfc5c2cebe4d44158eb4f4143732b405e60001b62006837565b62004b5e7f02985321103df1a52655ddc024ef2c389818485a6066c51773654d49695f486b60001b62006837565b8860400151816080018190525062004b997f0266f0a24649546890acb24fc0d1ceddd31946ea9537f2d0f8a75e0936b2ee1960001b62006837565b62004bc77ffaad186d1776330372feedebfcce16f23ec99732f4ade4c4942f0062faef527560001b62006837565b600062004bd4826200b827565b905062004c047f5ea09fe25ec72d6431e511c9b3c5d530a36b1f35b9368152b8782d9922c67ddc60001b62006837565b62004c327fce7ef4a5e243add278b5738374779c67a31e7b916eca128a0aa106ac41639b4b60001b62006837565b8060000151816020015182604001518a6000018b606001821515151581525082995083905250505062004c887fea4d0d423c62b94e8703c250338c827718afaa03af8d424a3cef22afe920dca460001b62006837565b62004cb67f577fec203cbcbe89b6fd246441eda598f1c5e3de17c7e2e83f3cb8a3e6e1308d60001b62006837565b876060015162004d5b5762004cee7f63e8a8143bfcd02d7e4d29e70e7e127f4e9a80791679de9c2c4ffa7c6008cc4560001b62006837565b62004d1c7f5e20ebc77cb04f6348bfb6baafe8a5fb28b9ad7809e5b0f481a87cbb208aec0660001b62006837565b62004d4a7f605a1acb7788d8e3cbe90cd05a1a52d9f64d23aa28e529e72057c6463e62eb9a60001b62006837565b879850505050505050505062005b48565b62004d897f3e84dc2e2db074317cec14ac70e72409e910952bfad0f3d4d6a7365e492b794d60001b62006837565b505062004dc0565b62004dbf7ffde7b2ebc1391534aa9017e62fc56b7c29c7c2b4a0b28f82aead1c181f46b94f60001b62006837565b5b62004dee7fe79d00cb537e8eb959a39bd95ca33434761bd0f26659a94a3441c7bc8a8d76f960001b62006837565b62004e1c7f80ab98ca08c0e18f72d82ab1ce5e6fb8e8441db2ff635dfabc4d9b4221c6eadb60001b62006837565b60028081111562004e56577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60046001600281111562004e93577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168267ffffffffffffffff16148062004f5057506001600281111562004ef4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600460028081111562004f30577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168267ffffffffffffffff16145b15620054425762004f847f5ce359594e5f3f4555cd42089b2d3dd6f1a1181894fe06171ba56ff9200dcc6960001b62006837565b62004fb27f78a742f0d4b3edcaffa55d4cec761ff54ff2f26167cdb572e09cbf336c7b62da60001b62006837565b62004fe07f51b31934af427b6b8327cd5ed184b5af109403214936d66e277c51f900a2a09560001b62006837565b62004fea620102a1565b620050187f0a8f2192e3aaf65995732301a4e33441ddf8575b54206c303e4374e16cd3b86260001b62006837565b620050467fb6c992f3c8a84ca894f9b405d2cffda24506e39829e9e5ac9b263d2879adb54e60001b62006837565b88600001518160000181905250620050817fc10ec6d5b928508c5af777739fe4e3d4618a2a1dcfb7838bddd24a83ff67ddc660001b62006837565b620050af7f8263d898fca684c9b616589ce2c7e42f7849e310a625e25d47759c7d641b2ba460001b62006837565b88602001518160200181905250620050ea7f402de156d276fd7b3a1ed3560b44aef6c392f961ce96cfd05d69455d02ab46e160001b62006837565b620051187fd8f6adbb80b4ecdabfce678f2535a6190ba7ba0cc8cd40b28b606616a0d10ee360001b62006837565b88604001518160400181905250620051537f8eb6f86ef2214a2cfd81fe8827bb6ba2f0797ba36f761a16294018681c92aecb60001b62006837565b620051817f02cf8b8d15c91c0ac04ee773a5d5801e7640fdcbb70e0cc1de3b94ca6ed7074360001b62006837565b838160600181905250620051b87fb78a38ba4947f4ae88c8756759bc30af843bbd48d77ec35676027fe1aa15fa2460001b62006837565b620051e67f44cf614f1b5c9e167d3b50a017f1141609edc3b97cbd36c337083bcfbcaf888e60001b62006837565b82816080019067ffffffffffffffff16908167ffffffffffffffff1681525050620052347fa79ab0b8e9ba6345f47e7a2419f2b3b230958d0fa5693fdf207781a428024b6f60001b62006837565b620052627fb93e4e3d9ce25718ca1fcd44d80f831003dc0f515b4d35876e85a7576af93aa960001b62006837565b60006200526f8262006a61565b90506200529f7f3c829438cbc1201a901e8e9f8936f4d2510606bf4379c11010ef627c4ff6fba260001b62006837565b620052cd7f24d269c20ea5adf30d647e78951ad9b2e3c59a5084ca16f79fa649f72a8488aa60001b62006837565b806000015181602001518260400151836060015184608001518c6000018d6040018e6060018315151515815250839052839b50849c508590525050505050620053397fffad259bcb3780004fa8ddede0a0f986b84d704c9c086e7bca5de5b4170f836360001b62006837565b620053677f9cae3326e58979168c48f8ea55193ec2812f7bf814fb9736f6fd16ed0456e06360001b62006837565b87606001516200540c576200539f7fe3024a4413b1b02ccbb167459b505b74a64018db2404601f7133912bc8a8d81060001b62006837565b620053cd7f470e47ce1a0a814dab0ec294b1d2ba8857e1f12e41deef93318613e00e32ac7760001b62006837565b620053fb7fd01ace81817e69818e5d6e0517dc23f1c261f02603d7401e1e10ef3ee5e5d4cd60001b62006837565b879850505050505050505062005b48565b6200543a7ff33b89b3d974dd2b7aeece3121190c95530006f2966807778a895e61adb5cb6160001b62006837565b505062005471565b620054707f2cf47e665ac2a70d736551d0feee0be5f17b27c0bee53038fb571103f68a70be60001b62006837565b5b6200549f7f6e3d6f086bc6ce67c1417692b165e4143d78e1eb8af2b2a4adf87d66444a431760001b62006837565b620054cd7fc685b4d7235665615a4102e4c911bc8ff3991a8dd6b2d063884a667e021112ee60001b62006837565b6000866000015160600151141562005578576200550d7f9e66729d73f2a8a346ae85a0f796e751cd8ea006492f0a68e509d421ea138fa660001b62006837565b6200553b7ff325fb963463c783be47fe4d501018bb90456e805747ec930a0aac1a9b9b3d9060001b62006837565b620055697f490b52fef4dccedb6b017ed95f3470d4ac21d70ce65fb5978e0ff420c3e04dde60001b62006837565b85965050505050505062005b48565b620055a67f1de9dbd88ad8894ec21e4374aa799f9487302a1b43b382bfaa80866af8c4c91660001b62006837565b620055d47f151ed89958251132c5e435ca5cfc8b8acdf3ab3dd1bf8343befa5d165f4d7f4460001b62006837565b620056027f5f34b828a2541ab14b75e998ba8a519ba41146c7f4d778be1922384642a952e160001b62006837565b43866000015160800181815250506200563e7f309ef2be2d380785f767c478bafac5bc5b9443e6e44ba9d7d5718e34b363512b60001b62006837565b6200566c7fe0413f6681757f2bf6ba2032ff3bc76588409cc7e8bf7590846a6dbb391ea3b060001b62006837565b8367ffffffffffffffff168567ffffffffffffffff16106200588157620056b67fa75151f5635597b95b079332f68290b496a31b46396e504488dab11e38b9f48f60001b62006837565b620056e47faaebae7edcbbea5f67961636a89ff78b05631aa2cdad1b40a2efc648c40a8dfa60001b62006837565b620057127f64447ec8ac1e71dee55395c37941bbabe0044e2d96f5b52ab0d9fcaab3e6a94060001b62006837565b87600001516000015186602001516000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050620057847fbfbce938fdaa14ea561b14dc7af06d1d053d0e8d384bb8e49697e93ba1a782dc60001b62006837565b620057b27f702ed32e5a53d5d9a9fb0e71d04c16e7fac9be3947686b477cadb1187c98411860001b62006837565b3086602001516020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506200581c7fb3eebd33bc01349be0692b04c411ca0964ae6b9f2b6f2ffe77861c3a5f3e463b60001b62006837565b6200584a7fe31c69c17f2433ff116e2d4ddebd9e96f87b530818971b81cfae7685b22e07ea60001b62006837565b838562005858919062013029565b86602001516040019067ffffffffffffffff16908167ffffffffffffffff168152505062005a75565b620058af7fecfcc67d5e4d14e319f5ffb5492cba33fe3e69672a06f7971711e2180f177d9960001b62006837565b620058dd7f012033bd668b27d06c0d18cbbdb0b3bc7d23837bfee029acfc0a440e450b825960001b62006837565b6200590b7f90efc1eb69da66ecad09431243996618c99f253a0416e684efe74d3024b26b3160001b62006837565b3086602001516000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050620059757fc09e3a90f91d840d4c94dde5c0ab49b8169df767eddf94f475585772270315ea60001b62006837565b620059a37f57aac37e7682c8e5b7fab4283d350fc69cb6abde0546f7bcb315a66a45cc7aad60001b62006837565b87600001516000015186602001516020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505062005a157f761e83d23eb1827cb3450364b816ea5ca76b5b5e941960df1bb186d84abb471360001b62006837565b62005a437f52853c41b16b27320136d07f6f526163331c19a679b7b8a823ee8600bba9c61f60001b62006837565b848462005a51919062013029565b86602001516040019067ffffffffffffffff16908167ffffffffffffffff16815250505b62005aa37fe3fb9794bc1b571280ce1bd9b86caff25cc9187aba5e63fd8ee9945d04043ae360001b62006837565b62005ad17f371f172a4b273f8b283651194d3e7ec4bfde14144aa5b07017b19569075422d260001b62006837565b600186606001901515908115158152505062005b107fe9aac58e75b80d4ed9d1c485abe6df2c7b82f0dc05d81113f293a635a7667a7f60001b62006837565b62005b3e7fd5209e0e48f26526105519c053e8c523533e0774d9603375617384bd4c230c6360001b62006837565b8596505050505050505b919050565b62005b57620102f2565b606062005b877fd3f9a36da398bda76da452735ff780ec2a9246617cc3e415b56ed3306b088dde60001b62006837565b62005bb57fb6467d58a46a50ec03fe8171bc55c3f8afb63f2277d87e56420249daff79fc9160001b62006837565b62005be37ffe1da04508d4244e8065067ba6e199726f17b36ecedfe181ba2d5da5bebbe47660001b62006837565b62005bed620102f2565b62005c1b7feb49a4b738315f586c7910042b42c7b4a44a23eaee74c07468787fecad4a3b2560001b62006837565b62005c497fa2c616e24a7b9a19ef6cc903ca7944cf4ee6ad66ca014a2cb0df3240225ba1b360001b62006837565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801562005cb457600080fd5b505afa15801562005cc9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062005cef91906201174c565b905062005d1f7fb39068ac76e304c0a1868c8b0b03193d67111c72b48d62ef99784e92217bd9f160001b62006837565b62005d4d7fadb1699a2820a312c27359dfb75c2b831cc098ac89e8c923ede518cbf4a105d860001b62006837565b600062005d5a8662008b98565b905062005d8a7fc2086fb29f91389e3b7aececc43b9dac3eec1e72cf6bb532bcb77f834bf80ce360001b62006837565b62005db87f19b9aa2c4ff24bc0eb1e8e3b9a810c15b52aed6209fb938ad989048aea09a0ec60001b62006837565b8062005e8c5762005dec7fce962d9571a5a543dc97cfac8e02a63b4d2e70c313b2f90f735da9db14cfdbc260001b62006837565b62005e1a7fe08fe1f88eae01b2674f2e279eef37867c03098721eb16f9289fff416229962660001b62006837565b62005e487feecfdf52001996de8654766276b5defa8b9986fce4c3adcf07ce13f7d53147be60001b62006837565b826040518060400160405280601881526020017f696e76616c69642075736572737061636520706172616d7300000000000000008152509450945050505062006774565b62005eba7fb9cbf6d9e1ed8920e4db3672bee18e280a2c847f29f6442f6d168d02aaf439c560001b62006837565b62005ee87fadea1f9e06e6457a0e7b3f932c41aad81eeaa92611441e38693357befbf43bbc60001b62006837565b62005f167f7f53ef95f66be033aa38bf3307ccf521e10f0df4f3f780c3cdfcbf12a4c8b7b360001b62006837565b600062005f278760200151620021fc565b905062005f577fbc63b03fc80f2963419962bc2f84bcd83d56fdfd8ab18c3d3598af50c6124b9160001b62006837565b62005f857f9eefdc6273d5ab80595b6a25b9f4f8e498fad0e67494fcf9a3f9714e7bce28e760001b62006837565b600080826000015167ffffffffffffffff161415801562005fb557506000826020015167ffffffffffffffff1614155b801562005fd157506000826040015167ffffffffffffffff1614155b801562005fe357506000826080015114155b801562005ff557506000826060015114155b9050620060257f856c3978303e038721919817b21c4b90e16eb31f61d0ca4039f25245a9c98fb360001b62006837565b620060537f0eb43414b195fd6453c0ce53b0c7e2b8748c25dcb5f585afc682d708235b75de60001b62006837565b80801562006065575043826060015111155b156200610957620060997f7faaf6a2ba6a77c56b0adafc1a31a2d92bdd8eb01699acb7ea7e793b995e8d3060001b62006837565b620060c77fe4684d85e2758b524bbf0252e3afa24a112f7fec72334aa2c03b04bfad1d213f60001b62006837565b620060f57f011266dc5118119c9e0d7558cc2c6f05e7ccd90012b83e8bdbeebc059e1137a260001b62006837565b62006101824362009478565b915062006138565b620061377f031862b4373f77c8a17e015ba5f6dae88a2e0dac83f22d7c09b2dab8ace2d78660001b62006837565b5b620061667ff2fd90056b7457cc3beff272ad44149b83903455a9418b0060dfcc2e10791d9c60001b62006837565b620061947fb547568ee3c57979455557383dca42a0ca3e7353c3beed8471342164d098721f60001b62006837565b801580620061a55750438260600151145b156200639157620061d97fc5f9a901de2f418c4190d6d6f0ef2001bca93d9c31fbc0e46dc1ef6956f9f08860001b62006837565b620062077f311034ce722f876da16a948aec88267d5bf9873d7917489d85a988d10144e93b60001b62006837565b620062357f5648c113d28ebddd1e095f4e17a4140eb7e53a48183f8ee8da577a9081033fa460001b62006837565b600062006243858a620096db565b9050620062737f848e5f09fd91b36cf3121a6cbcf39c825afa0ea3b52ae8f1bf6d81af731d15b260001b62006837565b620062a17f61b3c2c88c06d4d62858459317a69a4f601827e49e743f1dd72015c064fb966660001b62006837565b806200635c57620062d57f946c7b596ccdc5d785a1c6bd310b62ec647f990e2fad62440af341012db2c49160001b62006837565b620063037f6420c04d37dafad502709da28ab943c4b68f4699fbfcb509038295391ee4cac760001b62006837565b620063317fe72594022edd7b6eaef553c687a5256b8c269d566a5ea2487352b3f78de1041760001b62006837565b856040518060600160405280602181526020016201353e602191399750975050505050505062006774565b6200638a7fde781902ec8dcf9981b9f9081aec8230a084ae12c17eca77ebee42068951a94b60001b62006837565b50620063c0565b620063bf7f5981bc0b8835fe4a598379b79ac16c741b8fc9c3863b71a6c8d7f8974bf54b7360001b62006837565b5b620063ee7f2dd53b0586caf62672767c4835a083a0ebc5bd9e37abf1e022db4c977f6b09f360001b62006837565b6200641c7f4791af200f5a0871ff321cf113e7053ef2a7b1b8d54d5520f70e3e2525e42ea760001b62006837565b6200642662010323565b620064547f3cc9d2d4a9bbb0f6d61bfa24b79d2350cbba216f0b236d377a7d65e3c1d4c5f860001b62006837565b620064827f034ed07d31a4a7b20d07399007e2dfc0e2c7b92886ddc613a62c167bd3cd6d8a60001b62006837565b888160000181905250620064b97ff134b935fda2cfab877136dcf33835221411016612c02e229d228e5f9e8b6b6b60001b62006837565b620064e77f0720cd603bfb0bee6b91d8b2a8b199e1aeca0f7dc525f6c7961a81ef924304bf60001b62006837565b8281602001819052506200651e7f979b611d4b806734761847c6c05c806028f0656acfd8471d80a387d28a99724960001b62006837565b6200654c7fc1e820cf5e9a8a720435a43d9ae3b551a44c1d2a899e0386b181dbba46a861a560001b62006837565b848160400181905250620065837f68cf0eb477bbda4ca58635aff51fd56004549b21c28ea76018afdac0c448605460001b62006837565b620065b17ff7940041329f7712ec04ba624c1ffc7f08d8dbd9ac4d48379b014604aa08ebb260001b62006837565b6000620065be82620039f9565b9050620065ee7f1503b18df03e00d50cc4397580459fa49c330f44ed02494d283f7329f7116d3160001b62006837565b6200661c7f4280ce5c08c4714d7975bb4c133e44561186bfc8d7f9b8eaef011001bcdff65d60001b62006837565b80600001518760000181905250620066577feaa362df0a04fa3f04ca6aa2a569fc5995a1f2faebe0c01e91ed684d5df5dd0c60001b62006837565b620066857fc9c35e84176745fc96d8533d98927b3c0bce31da4189fe0b597a6f9f85566c5860001b62006837565b80602001518760200181905250620066c07f1bd52817ab90401d08075211fcb2932d310d64c79acf2b6b3b663cf1150bf04d60001b62006837565b620066ee7f48f0c22c0109c43365cbd8edf612d9c58c85e2ab5cb1fae479146f14355534d160001b62006837565b80604001518760400181905250620067297f60423381da00755a74ee4f101e6f71c52ed4e497711d9750698971c42e4a613060001b62006837565b620067577ff8b9f02e5eb56d65f4e37f781ea65921a3f1c66068d41c41f60562b0ca7a06b460001b62006837565b866040518060200160405280600081525098509850505050505050505b915091565b6000620067a97fb8c81879d8bd9db6c49e26ff5937a2581d63036f9eddc5efc1ec16afebd5430360001b62006837565b620067d77f5e7c27562f2ac9bab7cb83021c0fee5fa24437f679f1f3073d16673a13517aa260001b62006837565b620068057f43de10916900b2f711ed510eed8e49c1a8032ca508a5423142c6f9defa505d8560001b62006837565b62006812858484620009d7565b60018562006821919062012ef0565b6200682d919062012fa5565b9050949350505050565b50565b60006200686a7f5117d8a363eb6ce7671ecb07a9e609be447b7a84831c57ef55bc273fe93f360160001b62006837565b620068987f7930ab16975c5fd1e801c4394ba81262e14772c0a41a7b71462959782b259a8660001b62006837565b620068c67fece9e35444179b516d24c64016081754c3d60291014499a1fcaab8c668204f0260001b62006837565b620fa000828460600151620068dc919062012fa5565b620068e8919062012f6d565b905092915050565b6200691e7f110672d320423ebcdb85ef558f4913879989432f08c01658625b37bfa777539260001b62006837565b6200694c7f478fb1eef468f06847b9dc78dfe24eda0217eb82641ca9d25e0e87dcc146531760001b62006837565b620069797e242a0972d4157c18b1d9750cd4d922309c4e8f1232c66a64409a760378915860001b62006837565b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010155608082015181600201559050505050565b62006a6b6201035c565b62006a997f3b3a701b8a92922e34ad1190b0e1b800983ef5a7171d0f0579bd3e37c6d8fd0d60001b62006837565b62006ac77f407d10c0592da2978ec07222adafc313e098ad9db98b235c4bd778abd0ce9cd560001b62006837565b62006af57fe433bb8efe0ebf25ba26517b83bb594fdbb3619b5cd7226b7f825fcf46e3b0ef60001b62006837565b62006aff6201035c565b62006b2d7f3a7eaec189f8747f6f696c8b635c1ba4b5980385a1cdf83eea9289c383ce2d7460001b62006837565b62006b5b7f4035eaf34c9535857ce53abd8af18325f171cd6175e02a3ab6f18b0d26181ec660001b62006837565b600062006b8b7f5fe35ac6e7827196acb7370d9b6977951e5c2f3bb03a999e97936b5685c061aa60001b62006837565b62006bb97ff098ad1330a457bb1e0cb42452716211259f867744aec1f57bcc3b933b59134460001b62006837565b600062006be97f69a633ac7f56ad7112abec74c25106c3cc4b5aec0ca2b051bc2b39f8bbd0c75860001b62006837565b62006c177f738446e6dc5a2ba1447a8fb9e4cb28fe61b7efaf274e5530780358cc9a09089160001b62006837565b600062006c477fa226d785705349e168eeca59fb85ccd0306f6ae688ed9e99305d366bd24892aa60001b62006837565b62006c757fc28446fd8c2e305de18670f44244867f1ce28699ae07f8631502e85fc545fe7460001b62006837565b600062006ca57f44b66f8df50aa7001a8794589b9c9f75c29b83597ad09d5ade068c4bc4b8b7ac60001b62006837565b62006cd37f981a67224cd292279316b402ec24d7fcf89f247f84007a5188298cfb5776c68e60001b62006837565b60028081111562006d0d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60046001600281111562006d4a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff16876080015167ffffffffffffffff16141562006e7d5762006da17f5fb3be2a6b7ae02b62e7bcc837e35b7f08f1449f14703d56289dbc12218ca71860001b62006837565b62006dcf7f8222e6688e638eed226e4d33bad4c30aad194d8e0e1150d2c839af84871bc75f60001b62006837565b62006dfd7f85475ebd8117f96d1c7cfe25ae58e88941e1b14c36a6375795e9abbd76a6d3d560001b62006837565b86600001516040015160200151935062006e3a7f1c2deecdfe4adbe909bbf3b402c0131cf28dfa924dc4b814de47d5a00c5f465660001b62006837565b62006e687f48a8bdc620b575af180888bb896eec6484dee20eac81fde10c6d80eac3d4bca660001b62006837565b86600001516060015160200151905062006eac565b62006eab7f14bbae46dcb3af48862685d432c33e318c3786657fc24e5ba8568bb762b161e860001b62006837565b5b62006eda7fa8d585082f0df7d48ca7af95572d4c3cd5c38343bdfb08743fa31ab629bf234b60001b62006837565b62006f087fc620f9cff3e0ef0951ac58a6d9063c95cd0fad96b83333180efb559eeeb3f69460001b62006837565b6001600281111562006f43577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600460028081111562006f7f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff16876080015167ffffffffffffffff161415620070b25762006fd67f85c4647c931a86099f12abf233aaf857bc06a3de2a760ff95c1bda63e8bdfd3c60001b62006837565b620070047f4854578d9232697732f53913dbb0dc012a41ab9c8fc0d6291909b08e1c80381060001b62006837565b620070327fbab5305e24ec910b954199e42be9ff6212f2794b92c708d9121231e7031f5da360001b62006837565b8660000151604001516020015191506200706f7f4141d3f00dca8bb5cdbf8d04b043a8648d9261c2a369794fd2d21dbac64d9d7160001b62006837565b6200709d7f5f099997e5c97d27abde5bc9f937d170a09dde252a0c314c78da16539b14a47960001b62006837565b866000015160600151602001519250620070e1565b620070e07f3d4541aa773f38d1a762dddfd5132978b4b533473b783aa8a7f44e6e429a2ced60001b62006837565b5b6200710f7f8025ec3ae1af069948f9e6dc85ae27ad0eebcaf543c241501124155b5afff01760001b62006837565b6200713d7ff71b5ded23589e0fc0e46cad3ac2f0fe21423ab6ebf9ee91011ed83796b020f360001b62006837565b62007147620101f4565b620071757fe64487c93ce31c27733a06e62fe7937c773e9cb2a596a222afd12b95bd5243b460001b62006837565b620071a37f12333483880e2d8e635e9febfb72af42839bd25fc5848ae7b95b72190dd5829460001b62006837565b87602001518160000181905250620071de7f1b7537191df050afb6b9f89c041a77db4700fc4e6a967c6c2a5bed3c7833fbd560001b62006837565b6200720c7fc22195c94f01da46b739ee932336e23ad3e5732207de9262102b5fd17d89290260001b62006837565b84816020019067ffffffffffffffff16908167ffffffffffffffff16815250506200725a7f1e6c114257ef9a9c4ba3ae07f9f99284e6e0ae696f95993cea4c94ec4ce7879160001b62006837565b620072887f09a4177bc3989a626e66916103f5e7234786c3fbb1e52088f643d971d13c8b1060001b62006837565b83816040019067ffffffffffffffff16908167ffffffffffffffff1681525050620072d67f409c1cb035ce733479a21acc65cd37ab68df51ded01afaa78625c9225569477c60001b62006837565b620073047f78dec0032bba88fddd3432fcd4ec0fe828aeb47db7be615664c918f3e8fe64ff60001b62006837565b438160600181815250506200733c7fc72c932d060a64c37f22fbc425f40752c68e70002c4fb1cc9cc4b32a588650d860001b62006837565b6200736a7fd8dbf105f0421b0bed6c5f6c9b343526ce3a78cc7abeb60c8a70e45ae2bee79660001b62006837565b87604001518160800181905250620073a57f1cef09511bb341cb0c76adff82d622d0708b75fb271df4c0ce55bf46f86ed9ad60001b62006837565b620073d37f3d3efb968db42ce62caf3d41a91fcc194d2646cd05e00fceec82ec7535e9804560001b62006837565b87606001518160a001819052506200740e7f612620f2f7418cbc0a6ac18c065808c1f74a461ce364edd3bdf0f22abca4f8cd60001b62006837565b6200743c7ff71123ee9ec5260cceae9cdef46e6dd69992709cca84ec8fbc93f7099dee13fb60001b62006837565b600062007449826200238b565b9050620074797f7bc4f5c96b8d05986d067767e5abbf1ab0af160f6cf7d052351c21428ad6678c60001b62006837565b620074a77fac05931a5c0c0b46b847a39c4469912997c17f02d8e16a13b41589a5eb3e289860001b62006837565b8060400151620075b857620074df7fd752cfd5ca4a27155d2fea0fc1d339c41928bd10ca11a11a48ce719c669ad73060001b62006837565b6200750d7f53865404c5b7adcd9a6cf660eaf4341dc8eb04d12e131a1c9e9fe0412e5afbec60001b62006837565b6200753b7fa3b74c2464ced51b4d338256c51563550ec308eda9e81ec5e2c4cfc688231e4760001b62006837565b60008760800190151590811515815250506200757a7f98bf3b32dadec4771335de2f878a1ea487b3b56b48c818eccb91c41574fb1ed860001b62006837565b620075a87f3b5a053178d642f235b791ea83232dd5cec34e97bece5de755eed435cb16cf4d60001b62006837565b8697505050505050505062007d34565b620075e67fc3af9c739948dfd27a1d434d6d9f5cc6b724c3e118bf8fa3fa8eca3f33f6d1a360001b62006837565b620076147f436a4d0420c18866b130cc33987892c33659ac659ddc71d0a25c5cfc5097a7b660001b62006837565b620076427fa3652798618ec1b30c24524ef58394f5e960ba578f82df6c37b7cd3e1cda37fc60001b62006837565b6200764c6201024e565b6200767a7f95a486fb317136fbd391a1593c8525ab33f1eaa7871b48d6da4464a606f26f6460001b62006837565b620076a87f4de02cc3a4eec35db7f6fe44d11ea423076706180798ee327bcf7e28c5d1f07f60001b62006837565b81600001518160000181905250620076e37fe3e9afa6b0b9da48fea63805c5d1e432217037b1b1e6c8ef4f93740de796a79e60001b62006837565b620077117f4a744dd49600a2518af0d15209a4fe6b0fa81971cdf1ff45e22d58c0b5ff79b460001b62006837565b84816020019067ffffffffffffffff16908167ffffffffffffffff16815250506200775f7fa3d2c66c293635992668987f1599908a0c35a0310de6dc014d1b0f75fa4fd97e60001b62006837565b6200778d7fb46b984509d265d5db669fb42d6c81a533eb26c66f9d5585ecbd8a8ff2373e7f60001b62006837565b83816040019067ffffffffffffffff16908167ffffffffffffffff1681525050620077db7f6f9a3f9544bbd1761182eec3630c4a701915415ff1056b9cb563c49b655c354760001b62006837565b620078097fa074ec1bb488468ffc8a9e9c943f1d040458496a892ee367fa91b70ba0b743d960001b62006837565b43816060018181525050620078417fce00dbf2184e08a8c8db77e37fb1a8040356f76452edb9c42c7cc8c33e0a0feb60001b62006837565b6200786f7f4cb7b3f418655d77348e25795dffc6362a2f01ea398a84d8a692d2b471a8494660001b62006837565b89604001518160800181905250620078aa7f846b88f0ca3337ab42ba3dc16ceb04be587b0ceb7f2bf1badb072aa4632fea9260001b62006837565b620078d87f378efdc189fd15843876f441a80a13d00ce05474e6317d57d580c97898f97c9860001b62006837565b6000620078e5826200b827565b9050620079157f1583ec6b9da8a8dab29383e7e7fcc3d2c791ccb0a13cd93973adf95ca6499e8c60001b62006837565b620079437f0653ac17fd3609096335425b23778783e0eeb896eb33c21955f99554f1a4366060001b62006837565b806040015162007a56576200797b7f1f9fcf76902d1818cfa4bd479dd7e83158cc51b272e871adfba83439874e048e60001b62006837565b620079a97fe8dfe6119266d42f39e22a87b1091035a5e4f90f944ddc0daf7421a3551ede4b60001b62006837565b620079d77fa3231c5fc846f5d4179a7f1d651b2bc3229a287112d5e526ea33efc5f0b74f1c60001b62006837565b600089608001901515908115158152505062007a167fb8cae1e9e9ed4987bd0f772e12adb8301503586eda776d802b752f26c9c4ee2f60001b62006837565b62007a447fa52ec1ad8e9653b5b0c1db7b36ef562380829c267fa3fc3c8e464973bd3445e060001b62006837565b88995050505050505050505062007d34565b62007a847f59d4690561fe9ec5bafb324e8889504ec4fe5278e96f59d0d518448e9722a69c60001b62006837565b62007ab27f152f2f8745a15bb8ce296efd72f2112e39dbeea196e8976ad1558a710e694f4060001b62006837565b62007ae07f8e33119928d174fdf0ede40101053b81196bffe76ecce7b7723e017de0b20afc60001b62006837565b8060000151896000018190525062007b1b7f6564e86b2e0e1a58a1aec74658e544a706a8af75715d4ceb20e4f2c8788df39860001b62006837565b62007b497f4981023d369c3146a9b0bc9b8b3f92fca813e3b3e713d0d0b799983ec12d9df660001b62006837565b826000015160400151896020019067ffffffffffffffff16908167ffffffffffffffff168152505062007b9f7f37e0147de5408a16f2d1ebd0440718b378f392adef6291fe379f8c31347c34f860001b62006837565b62007bcd7f683468f72a07324acebaf79848df2b6ecc7e4044edc6848358a25fc6c0f1233560001b62006837565b806000015160400151896040019067ffffffffffffffff16908167ffffffffffffffff168152505062007c237f7b3425b2df94518dfa2fc5ecb107863dbf65193296274c3824a19c263c33ffc260001b62006837565b62007c517f9d5362cc013d45fc44f468c8a4b087a80e8dbb1c971cbffc16f0d2bdf397682f60001b62006837565b8260200151896060018190525062007c8c7f82a111aa1bbe88ffdf54c08f56d85dbc4002148316a9e5a5c3bb2107cac94a2160001b62006837565b62007cba7fc353fd51673ed2e5cc8281576e03a457ef7732d6b108bb4b2f300813e04ced3b60001b62006837565b600189608001901515908115158152505062007cf97fd2bbc1a121dc379714e0198dcd16807d4a53573aa13c338a6533c48a8ae301e560001b62006837565b62007d277f36a130573f8c3a957ce8191708e1084a5a2d228d07ae0b09637f545c0c91805e60001b62006837565b8899505050505050505050505b919050565b606062007d697f1b87cec5659b70b5e7019994a8252085c98c12da9208ef0190b2348ec946a37960001b62006837565b62007d977fdaada2dd415496f10624e6705d16c760383150ddbf32d984fda48e7ad04da49f60001b62006837565b62007dc57fabd7a0d482777d6ae676bac15e65746a019f7a966be7a873aad8397e0d0af08960001b62006837565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801562007e3057600080fd5b505afa15801562007e45573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062007e6b91906201174c565b905062007e9b7fb279465411538f41202780f1293c4d9fac5bb06f72b1d3b48558d712dd9ad64d60001b62006837565b62007ec97fe4823e0929e0bb4e369edb4e7b1c83eced16ad2d4cae7f8f4c7501f14820ee7f60001b62006837565b4384606001518260c0015167ffffffffffffffff1662007eea919062012e93565b111562007fba5762007f1f7f182fd0275b1ea88abc3192249c966e65ae0760f8daa75a6114ed9a36fba2883560001b62006837565b62007f4d7f319f09f492cd0382a5c3651cc5dd2d185f36c0c8ac67ff4662bb4c95b7e622ac60001b62006837565b62007f7b7f315f60420cdc1b97f82ecc4dee57e996b960d2c00e889f030a2f76a7ac64345b60001b62006837565b6040518060400160405280600b81526020017f6e6f74206578706972656400000000000000000000000000000000000000000081525091505062008b92565b62007fe87fe3006d5eabf4a8166eb70deaacb1ab7081c212ce30a094e246ca05bca753159060001b62006837565b620080167f1e9b8825b6fa9258aae508fe205450360b50bf58c6b500e8f38cf19752dee36860001b62006837565b620080447fc5e2ed9a1dad3a0b53cbaf5dd55bf7356ba8842d45706fdde9dae2f922a5832560001b62006837565b6060620080747f7df9ed7f5c162a1bfb7c71fd8ae5e6d394c4e12685542371839805d24fa1739d60001b62006837565b620080a27f7b9f55888fd89b575bcb916b8ab531326faf4264c1bd36352f97ce932fbc48db60001b62006837565b6000620080d27ff141bbbeb68751ee5d08c892160e834265c6ef0cf3adbb7e3a1a13caa52c268f60001b62006837565b620081007fbc8693ed91fd21a7cc35e8765fb96ad45dbc086ac30a478efdf2dd1afde40a1360001b62006837565b6000620081307f6833c2a8875183a126ed02627806a2f475a195a2dd647643e37d7d4c51f8e42e60001b62006837565b6200815e7fd83fa78a858743cd8031eb12e82cd4e7665af4f50e6558c120143995d6dcf36b60001b62006837565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639a2b5e63886040518263ffffffff1660e01b8152600401620081bd9190620128c0565b60006040518083038186803b158015620081d657600080fd5b505afa158015620081eb573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906200821691906201151a565b9050620082467f7cd48684f70951f3d115615078c1143b6e292b050307d5492f425b6c197a65b860001b62006837565b620082747f139d85af3e6a2f8cd6a03ff9a7a70f7427081ca943a0bd48b9d58c295a63155060001b62006837565b6000600167ffffffffffffffff811115620082b8577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015620082e75781602001602082028036833780820191505090505b509050620083187fcd88f77fd48cfe40571a27a880a6ca7723a8a0ab6ddb802a68cb7f31a044135060001b62006837565b620083467f9489c0a8038475bcf343364692b3ad49f13388135a9818ee72098cb745f0f15160001b62006837565b60008160008151811062008383577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101906001811115620083c4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b90816001811115620083ff577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81525050620084317f55917990efe100c6d6cf4ec77fb0bbc3ec27a0244004237767a0406e0d5dcbbe60001b62006837565b6200845f7f6764dccde9e06a66dca7792693503d4b9564f423efee59e0860a26b5e79253e460001b62006837565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633ad525a9838a846040518463ffffffff1660e01b8152600401620084c09392919062012911565b600060405180830381600087803b158015620084db57600080fd5b505af1158015620084f0573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906200851b91906201155f565b809550819650829750505050620085557f17659c25b682183f188abd241d48da4ab93d19b9d7cd6a3f3d77e9517441a56c60001b62006837565b620085837fe5d144c40009228b9943b594611a0561f4ce40a60c88b6b96bcfe446c62ddc8360001b62006837565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166328a8565c8a6040518263ffffffff1660e01b8152600401620085e29190620128c0565b60006040518083038186803b158015620085fb57600080fd5b505afa15801562008610573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906200863b91906201151a565b90506200866b7f4649766d59e06dedceb92ed300069c02368e2549933f1a5bc0a58fadc54dcc6160001b62006837565b620086997fc486072e29f4dcbf1f3a85caa783b77ed95d08c921a789710d21e808d996cc2260001b62006837565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633ad525a9828b856040518463ffffffff1660e01b8152600401620086fa9392919062012911565b600060405180830381600087803b1580156200871557600080fd5b505af11580156200872a573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906200875591906201155f565b8096508197508298505050506200878f7f1deb8d8ca8c3510d724b3a038036794ff7d6a73fa5fbabc3cf17a0fa13c0ab3560001b62006837565b620087bd7fe88854c3d84f356e896dad8ddb10575506a76d4c2ff9b72709fd1218f327e24760001b62006837565b60005b865181101562008a2c57620087f87f7cd4d3428e353278c802269edcaea79d58fb804c5249a504dc854012eecf9c1660001b62006837565b620088267f2855879c0981519f351cb8d7238dc9cce4b79182376bae9029d3fe2b97e76e6960001b62006837565b60005b825181101562008a1557620088617fd09b876996961d595f11d44b1c93ed847bd1c0930a99cfbc59ca543e8415e9cc60001b62006837565b6200888f7f75f94dca39dc8164004872eeb3e8cf88011b73616c2b111708bdf17b6c63d1d660001b62006837565b828181518110620088c9577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518051906020012088838151811062008912577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151805190602001201415620089d057620089567fc396039f21a1cbfbe6b059c35bc4a28a5d006c94d53131f854d4c9179927654d60001b62006837565b620089847f53012faa7a297288d8bfb5d9980ffef807643497f04f5860cd215f3953eede9560001b62006837565b828181518110620089be577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001016060815250620089ff565b620089fe7fb602ee32d9b00ce20212cbc2c44f7a40c6e30fa5e44eedbf6f0050198c65acbf60001b62006837565b5b808062008a0c9062013213565b91505062008829565b50808062008a239062013213565b915050620087c0565b5062008a5b7fa82fe406bff1a0edd992cd530da8137a8fd657f2ac9238cbc48227887653901860001b62006837565b62008a897fb9c2fe7b35699a23d200f8b5118570a61a6ccf1269d4b28ca8c412735996dd7160001b62006837565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d70e62728a836040518363ffffffff1660e01b815260040162008ae8929190620128dd565b600060405180830381600087803b15801562008b0357600080fd5b505af115801562008b18573d6000803e3d6000fd5b5050505062008b4a7fc16228e63356a7ee419859d283a85ad7114fc95f62d22569c5aca5052c5713c360001b62006837565b62008b787fa835233da9c317ceeb829c68ddad4180422499e25c5f843f110f89009f29868860001b62006837565b604051806020016040528060008152509750505050505050505b92915050565b600062008bc87f4efcef3f7869db45c6c0dd01ee181f7ca63c5adb1c98941d4fd455d3ab0a4fb760001b62006837565b62008bf67f06edfa051733582e920c32798b6070f80151ac0b967c5ea19d162b9f57f9368d60001b62006837565b62008c247f9e1b708214a3314bff58c0cf54dcfd0a7844b3454be0dc81950fbce96ac2587960001b62006837565b600082604001516020015167ffffffffffffffff161162008cd35762008c6d7f882c693d0255e85412588588a23aefb41c97ad4549da2a03d4ed48936519b54060001b62006837565b62008c9b7fb84fc04413be701b927cf68890b3a6f0de1da8e13ec6f5e22a0ad545b609a8c460001b62006837565b62008cc97f90c3256075a41c40adac2a82be4da845c0fbcbe24670a69078d0dd3f974b393360001b62006837565b6000905062009473565b62008d017f0f4962b0eba61ab0c217af03b7dab05ee1cc2c8f306e207dd9d056e17bf01eaa60001b62006837565b62008d2f7f16ea86296b455ba58f2e9238f9b88d4acfa178d50a1cfe3240a4ddc2bd9c6bff60001b62006837565b62008d5d7fc33f8951ed9b29b0c5eeac7bc089dc2b59678d21162bdfffea3bdcb9aa8dcb8260001b62006837565b600082606001516020015167ffffffffffffffff161162008e0c5762008da67f5e84b73719a591a3c385aeadd49ea2199969a4c1ce69f57531878de25862c66860001b62006837565b62008dd47fd39bf90cf7cba5d5dd36e9bd4c6e52be991e65e6dee179b1ad32b4ae2c349f2560001b62006837565b62008e027f15b6020a3572fe821e8b7897559c979e199dae3b9d8b4a5c68e56e0d4386153c60001b62006837565b6000905062009473565b62008e3a7fa81a85bb9f6652b8dae772b60672d2628f9753e0c56805ad6bad914548b0977660001b62006837565b62008e687f6a81c97293d9c641b79a54567d6993e8547fd5e1b2f588d6cce03335f681917a60001b62006837565b62008e967f53d2c583b49b84852d8bb48cf60bc8c1bfe577b5cd496779edc7655dffa57ca660001b62006837565b60008062008ea4846200b366565b9150915062008ed67f18c04fef350b3438822b5676c9c24cc91940586eee520c9a0b7a92e88d59a0d760001b62006837565b62008f047fbb917b88efe37303ba75add38840f4ce06e58000bdb9194a7556ab4b7747bf5b60001b62006837565b60008151111562008fa55762008f3d7f22ed5229acd6d2474de20ed2aab0c04da3eb9187d7a9034325870b7b31c0bc0560001b62006837565b62008f6b7fe2a40648e289833e7e8a190939c5bff008735bae88bfcee64454c5cd98a0102460001b62006837565b62008f997f0be0f55d45e79d3e8ac3e845485191c57beacc763970cb8e13e5f8c3fb58bed260001b62006837565b60009250505062009473565b62008fd37f0a67c2ccb0ddb2efba4ebad388c30278ec28a6cb7366f6f45c1ba139d5f8ccd260001b62006837565b620090017f1afeeded232a26fb44ff12cd4812d8507971ee3dc48da0dac2f1f324c324894460001b62006837565b6200902f7fdee443b3039b85ec403b8595176cc226dca366b440b0d21270d475d1dcc6dfa760001b62006837565b600060028111156200906a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b67ffffffffffffffff168267ffffffffffffffff1614156200911c57620090b47f956506f0f725f393b1d4976a18f76c1df4ff929a8073639d041f1a141e9ed1f260001b62006837565b620090e27f4e46d52851a6684e3716dc289f6d6cbbd7cd9beda27593413cc7313778d9b91a60001b62006837565b620091107f3b5746e4fc129aa34daedd9712ffd242ff9032578e7224671188888ae7b5cbb560001b62006837565b60009250505062009473565b6200914a7f4d42bbeb863ab7ee609fbcdb61d154343e6e49a37db1239a3911e3a449db1eba60001b62006837565b620091787f919d7bd7b21cbebec18362351cfd6d1795b7ca6a41fd21add087654351d09c9660001b62006837565b620091a67fd090990afcd8744ea751cd864fc980b203dfdceeadd805e8109161164e88b37b60001b62006837565b6000620091b3856200c579565b9050620091e37f14e3eb0755b12c46766eb38f36bc9294b98ea2215a9cad01180f6075380fbdf760001b62006837565b620092117fce67cfd862ac1cc556b0fb63cd3ab953cae0641291d98aba0cddb2b143578e9460001b62006837565b8015620093e057620092467ff782d66862988964e0172ea54471137eb781b340b794ef3a4c4cbfdb05c09a8d60001b62006837565b620092747fa111fe4823a722b0195756db8c38d0a14f19ca98b9011ea59c474c4dc44e74f860001b62006837565b620092a27f7b3b74402fbccf6767d596dadf127aa03368e7d892b3788cb63822de6b36b08760001b62006837565b6000620092af866200c70e565b9050620092df7f2a0308beffd6f34d2bf1a956476ce435d4dbfc29edda301ddfc3bb2caa178d1d60001b62006837565b6200930d7ff59995debc7c6ca32ce7868aa36bb241542d7dc5cb21470eef2edab670d1af9a60001b62006837565b80620093ab57620093417f0f38aaa45b11c044100a188be45c34c51de46dedd4cfad80a1d2e8863af3493a60001b62006837565b6200936f7fc1493c5e0c947c324df6775ae395c933d979cd6e476f9ae61f813f8c99e3d46660001b62006837565b6200939d7f89e5bbefc751a2732f28fa7f909c6f5b55442077cae36c8efbd3011c51e0af1d60001b62006837565b600094505050505062009473565b620093d97fdb3204371f194de640067ea195a4eff6fce97afa5a9ca9b2667fca88f314727060001b62006837565b506200940f565b6200940e7f4758424f5d31437f578e0a956b151115612d4ff17951dfab01633d1b254ebedb60001b62006837565b5b6200943d7f9790e13a5c4258562f0a16ff2da5dfb5f55cc2ea95feb7a014229456efca6ef560001b62006837565b6200946b7f4634b3a81f1d910f95ea3ace2b44a48eeb363eda54377b8f178cd0108c77749460001b62006837565b600193505050505b919050565b6200948262010142565b620094b07fe184d69ec043b223c71685f20a46f97649ce6c733a5753331420c50a55ee397560001b62006837565b620094de7f0e20d51cecbf45d438fab9a122f2d81a8f15995a40c166d8555af5d03e44e21060001b62006837565b6200950c7f431027dba1c94b6b87df504d919b1eefd8b6182ac8e01f8c1531e76e976a02e160001b62006837565b6000836000019067ffffffffffffffff16908167ffffffffffffffff16815250506200955b7f8e6c7bbdb30c997e57ce8a3023b80160316efea3adef68c98c78e56ea947568d60001b62006837565b620095897f583ef8bc45d39c1ac9b3581d921207d88aae9346473e636f4279f90309dd356b60001b62006837565b6000836020019067ffffffffffffffff16908167ffffffffffffffff1681525050620095d87f13294eab4f5e84eccd287eaf02253ce573d5009155e8e15596881cad297fd2a660001b62006837565b620096067fddf4890078988ee50efde6deb29dd6b07337412098833d03fc6dce2b4fc209c660001b62006837565b818360600181815250506200963e7f6b205f9ef41897056e4e1e5a0a13d7a2235a5afa6215f352b0198e09ab0f6fcf60001b62006837565b6200966c7ff7d750adebeae16e62b35937aa905c799960db07d9226316a737892f324f7fd860001b62006837565b81836080018181525050620096a47f3a49e6bcb0879f7695e1235d9e1d5b05fa00076f1a49fbf0d43c9b47d619a60960001b62006837565b620096d27f1af7b88c8cb97c75204b5b6a1de1536e017bac3192bbbb7e1699a32294ac6f2e60001b62006837565b82905092915050565b60006200970b7f5133d9df7fe5921b447ad2ee1cb2380b7ea5eecf9a7a775223558e363cee321c60001b62006837565b620097397fa7d3313436ed19dd6240ba59393d3f6fb44b4d961a1e8861b77292735dbed60f60001b62006837565b620097677fae6a9c9318bfa6b4778c8bc265d01c76f7e2b7b7e62f99cfaf1360626da645a360001b62006837565b60008062009775846200b366565b91509150620097a77f41a6dca72f32ce8ea7f68bc07c307837f39f0d2c5ac9aa78c8b9ef976f409f3e60001b62006837565b620097d57f55623608c3b0590f3ddad955510f2746c6d0a9654b2d0195276b794f7bbb1bde60001b62006837565b60008151111562009876576200980e7f26c64baf4467b372751c84f4fb2db28063b54b9698f154e776ddf481d21f9d9960001b62006837565b6200983c7fa5be71ad07e62f49b9e3de168ffbe05736639e963ec9b98a0cc034a76266114960001b62006837565b6200986a7f276769d2c23495205bcdbe4cb944ea69afabf66e9e0a98dab563accda792469660001b62006837565b60009250505062009d65565b620098a47f0f0ee4c24b7dff7ee25cc4611859fb8f283580b1b7e5ccc8e1c1e7247669bd9060001b62006837565b620098d27f88b77207e1b7c72520461db4568f56fd4756ddf663507fcd71f2c8e179bf56a260001b62006837565b620099007ffd52bdf833c6cf68a20873a14783b9a22ef7b80d6849ab055eaa4c2669c8081d60001b62006837565b600160028111156200993b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60046001600281111562009978577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168267ffffffffffffffff161462009a3157620099ca7f456e8bb579c515bbcc3840cc633260efd1845619cb74d1dc81b5773cf6ddab0e60001b62006837565b620099f77e5cbba131b9d390c6d22c5c133b57edc26d8aefc737994d4df39603783bd12060001b62006837565b62009a257fddddbf2a651826ad19bb851dac0ea70c0996a0a1e0b0b6cb4bffd511866b6fd160001b62006837565b60009250505062009d65565b62009a5f7f4f66c7d806d99b9e8e71ec46e1e3e592f94efda3aae317f4f6e28df3d6da995560001b62006837565b62009a8d7f0e28c4300dbb63bf59291333dba0756423ec4df1b039ce756dab6b3151df2a6660001b62006837565b62009abb7f98734215a90a60fd7be1a6ba862b6d28b8b449907c297aa73a567a56d5271b6f60001b62006837565b600084604001516020015167ffffffffffffffff16148062009aef5750600084606001516020015167ffffffffffffffff16145b1562009b8b5762009b237f875695b3cf8cb39b0c02d0ff1087641f8f9e9c7e5e5dac41547bf55a7a28dddf60001b62006837565b62009b517ff301ce8d2318cc21e7b5d35598ac0b2a24d1bc653413f1736d18e45802a594b860001b62006837565b62009b7f7f63a23504aa555094cc43aae8e4e15897e2d18054b13985be2efad1395bf0769460001b62006837565b60009250505062009d65565b62009bb97f8a04043012fa6ff9c5f3c39c3efee1764d8d11ab7a22f7545d48b1356bd6a6dc60001b62006837565b62009be77fa529cfb0fec637d4450991c896a2157ef297fa750f8dd62164473f1513ee4bf960001b62006837565b62009c157f76f585de3d4fd4c1fe352d43692389b2703ec25714695e21762e24bacc574de360001b62006837565b8460c0015167ffffffffffffffff1684606001516020015167ffffffffffffffff16101562009cd45762009c6c7fbe51cfb277858a3ab7f9f191545985ac2cff7f2f9eba36adfc07c48fa65799f960001b62006837565b62009c9a7fd035e98f86eacb4c9fc44697399f418da5d10726bf17fd222b245c194804800e60001b62006837565b62009cc87f88fe1755f7c72338d0c2fdcbac83441f085e9fe3b3104d7e3a9655002c4cea3660001b62006837565b60009250505062009d65565b62009d027fbb7f56bf5db9a04a63b2438ad7195fdd1eef41a182494f461647e596e1a1012260001b62006837565b62009d307f5f4ad1343d05a2cbeef9b4ad2b87c6836feb9b12b278f56467ec1839bab7bdda60001b62006837565b62009d5e7ff18bc3b71c609fa46a345e4873c6e31e0d4bad410528c5deee8c4cf5c397ea1a60001b62006837565b6001925050505b92915050565b62009d7562010142565b62009d7f620100eb565b62009dad7f1ae9b671fd6794058d8b36bcb116535b233ccc0ea2e32f5ea27bd5764666198e60001b62006837565b62009ddb7fe2ac908f2c8b11e37b900d52695254e4ab5225402b5b33553874128726d56e7360001b62006837565b62009e097f6706ca443363f4b4150b2314a72adcdb6b88ee13a3289e1f56d2070d080399ad60001b62006837565b62009e1362010142565b62009e417ff7f32db8e3c1b6a593745ee84f0743060a18629f9d431d29d12200e33032a26f60001b62006837565b62009e6f7fab3627bd342dc2bdc975c2dd877754258aa1c41949a7a6b2e0ca9ba6c015ae5760001b62006837565b62009e79620100eb565b62009ea77fa49b4c2e0a6ea16c5cc239499d19d1237c1ee6a120f65de3d78574ccef4648ac60001b62006837565b62009ed57f1fb8d83a325937c852fb92164405e8f1beed1d8199cd8c93a0ebd3edd6b6637460001b62006837565b600062009f057f2c9250973bda70191fffcaee17942607da85f51b22d51b3da45b9dd0039c9ecf60001b62006837565b62009f337f3d5e8a118a8e06255c429528b51c7da54e318d97d3de0f951d6de8a8989af96960001b62006837565b600062009f637fee7fddf174990138eff78dbe0389f3c18bf0e8ba56c94a6db82f47db225c599b60001b62006837565b62009f917fa83dac38365f6dedf41ba2ee2a1651f9279500ba049718719d04f51499a658b160001b62006837565b60008062009f9f8b6200b366565b9150915062009fd17f5df3a06e5f19704423331af1d729209266bf1de47fb83cb1b16e689fc698f71e60001b62006837565b62009fff7ff30a53b79e768f97ebe73001a53feb97cb55598f97757122cef98760e915a7c660001b62006837565b6000815111156200a0a6576200a0387f0d1b192fa08ba58630d7d18782a90f2f90fdb86ec6a5a41346ac127fa8cd4b9a60001b62006837565b6200a0667f3a0c03c5b557cabfddc8921d0b037d4a7252149a9c6427d1d419debdc946aa2960001b62006837565b6200a0947fd75af609cc1b63c812d8d0611f0a054a76782911ef710704902ba802a051830160001b62006837565b8585975097505050505050506200af34565b6200a0d47f769fa09b4aa3cdd390805cec9de4bd3ba03934c78c766662cc9e34b4cef77a8260001b62006837565b6200a1027f6b71b9f9f10c315d8319ef7fcfc392a03797004883eb43d3861623a9772f357c60001b62006837565b6200a1307ffc41056e083e5d894571696546368de8eefc4a77a49e37361688c9b94034706c60001b62006837565b600160028111156200a16b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004600160028111156200a1a8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168267ffffffffffffffff1614806200a2665750600060028111156200a209577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004600160028111156200a246577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168267ffffffffffffffff16145b806200a3055750600160028111156200a2a8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004600060028111156200a2e5577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168267ffffffffffffffff16145b156200a4f7576200a3397f3a29f44def9df951f06ededc88845c8460b6c0252ab632aefdb70dd4545052e260001b62006837565b6200a3677fa2ee495d0035780830f75d06d99fb5c959a688ad6e1a499595814b391f3f820d60001b62006837565b6200a3957ffe943f881535565e1a2e0f2167df41bee7adaf254c6b4170e45ced2cce2089b460001b62006837565b6200a3b48a8c60400151602001518d6060015160200151438d6200cb44565b8093508196508298505050506200a3ee7f82d5368276cd6d68518def9465d8c669114aff057f92a3247521c3c671206ed360001b62006837565b6200a41c7f5b68063ac573eb53ad15da598c552d74dba8d94f68b7b5b34d4c1dd436d580b060001b62006837565b6000815111156200a4c3576200a4557fca064f5d491dbbb0d9e44706cd5e2e3988e9940086689e35dac08d18cb8c559260001b62006837565b6200a4837f3a6d854b87da176864fd3e21a50e0e13678468e2855d43aab453eafbf655bc8a60001b62006837565b6200a4b17f3acd40084913aa3c7b44398a763a4156b3f7be45e6bf82b7e588accd799cfea260001b62006837565b8585975097505050505050506200af34565b6200a4f17f9d3e7d3cbc4a79d630bd1c210b9bd8bcdd691660e3dfa3952222fffdb865f99060001b62006837565b6200a526565b6200a5257f9cdaa07826234a54afc0d07cb6a527aba28099fc061e49a818b87c0953191cbc60001b62006837565b5b6200a5547f509990849667e573d537713c118472d9429c7dd1d8cdc56ec5b34141088db79560001b62006837565b6200a5827fd53fc32b7149bfe0978cd0702a81884ce59fc3b789c8b34377e380fb0ae1272c60001b62006837565b6002808111156200a5bc577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60046002808111156200a5f8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168267ffffffffffffffff1614806200a6b557506002808111156200a658577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004600060028111156200a695577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168267ffffffffffffffff16145b806200a7535750600060028111156200a6f7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60046002808111156200a733577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168267ffffffffffffffff16145b156200a7f5576200a7877fe53e363b7e6b874c75228eb38d266b236d0e7547f0162fb9ccb266deb0a2986160001b62006837565b6200a7b57f9e38c85b8ef4ababb0977c0d0348e8bf1a1527619a426308a79b8abfb1cb114160001b62006837565b6200a7e37fa3f1f8e9c04b18a7ad9299c9cf32c51dca72d4865649046d292417c3b1ad28e960001b62006837565b8585975097505050505050506200af34565b6200a8237f4ff5665205ed52711fc7fa58e8aadddd7865dbd9f7a904207c7cb4a8a743a54460001b62006837565b6200a8517f7f8f8b5b671430ba2374243f158c5050d1fd50167ee6aa4c7d960a1877450e4660001b62006837565b6200a87f7fa2526b085e0e23a8a353d7f4430c246ba0fa81dc483aefdf70257579bd9481e460001b62006837565b6002808111156200a8b9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004600160028111156200a8f6577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168267ffffffffffffffff1614806200a9b35750600160028111156200a957577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60046002808111156200a993577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60ff16901b1760ff1667ffffffffffffffff168267ffffffffffffffff16145b156200aa55576200a9e77f68b3c772b5eea5d53f620641ced887f2a832b7bbe9e63dac5b023a602a42b38860001b62006837565b6200aa157fd9ac34e31dbe16716c8abf244dc35da62c0601f178c7f13d273b7c9a1272768d60001b62006837565b6200aa437f3428b8d3957b5ed8a1bf12e04cc790580632019b48cbadd88704733c8e229f8e60001b62006837565b8585975097505050505050506200af34565b6200aa837f9992dd28147edc79891f4e65e904ad4a03a16b28312ebe4bf1ed3836861c3b8560001b62006837565b6200aab17fe9ee47de24a300f369359a9dd7b4b1296567495fcb650c507695777bc615b76160001b62006837565b6200aadf7f5050b364d0795e764ffa12cce85574da1c3d2c4350b996360e7753e5184e066060001b62006837565b8267ffffffffffffffff168467ffffffffffffffff1611156200ace6576200ab2a7fc0a21ef6cfe08485ac3695cc673fcf60f81e74962b9fcbb885557cb6da38a72760001b62006837565b6200ab587fa8ebd72dc466cba8753ca6e57a41919649d5ab7084634782b2ddc88d9d99381c60001b62006837565b6200ab867fc5c2c1cb1b5dda34573bbba8478c506648f9fd9b838ad07cba5bc5188a13c3fc60001b62006837565b82846200ab94919062013029565b856040019067ffffffffffffffff16908167ffffffffffffffff16815250506200abe17f343956781e1a6bcf0ed1125af79c636d1f9404cc5de1781f229d4db4b20e090560001b62006837565b6200ac0f7fb4b1a79b35890f2de75300c6a0a757280be4ff8363417c545f13401e737f1bbe60001b62006837565b8a60000151856000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506200ac797fe78fb55c2e79419959f4221acd46ab9dcb1ababcf3ae3034472548c4c3f6a05360001b62006837565b6200aca77ff62f89b0c2120642cda774ff222372582d8ec58cc24e29fedcb04735a62d474860001b62006837565b6000856020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506200aecb565b6200ad147fcfdb25dcac78d3a197b468bf6f99fb2e792976ffd81b1ca2200b2f5710d4077560001b62006837565b6200ad427f54aae939d5f6dbc001f1c97f30e58869a8fe8f654303cfc1c34778e06f148db460001b62006837565b6200ad707f811c3c1919194f011d586915bca9f184464b172cebc8f042e9c354794f5937b160001b62006837565b83836200ad7e919062013029565b856040019067ffffffffffffffff16908167ffffffffffffffff16815250506200adcb7f48a5b3470e99e067391e5c58bf6121c9e617c86216545d4a3c3b6d00aa7bf63760001b62006837565b6200adf97f315e56b5ca1564514f3d29b3832e2ad3805b17c39d59d57dd92f56b5dae2c86060001b62006837565b6000856000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506200ae607f339a764e78200d8f5ab274dc6420594e5127b161e130d9b15d7bbeddf52ea4fa60001b62006837565b6200ae8e7f3e03fd1fd30de7c669290b9f3e3b7e4ba68fc784e22adf663eb369556a610a4860001b62006837565b8a60000151856020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505b6200aef97f2d29e41131b4cec5c42f5899e315a74b27a699f0905cac12bd6d67d7fc8e80ff60001b62006837565b6200af277f0ad42e3667fe67ae8f546c2db6d7b57e098166e9dc2bf2daeaade55501327d8f60001b62006837565b8585975097505050505050505b935093915050565b60006200af49306200ceef565b15905090565b6200af59620103a9565b6200af877f8643243e0e5e1a4fbbeade6bed94bdeada0e6a18255d7ba0aef31d50532a348160001b62006837565b6200afb57ff03a9e1e16b21fcf20d0248517dd63c28e5ba63a4c88f6fa8e2468a56c338e7f60001b62006837565b6200afe37f8486db0ab03d68d313776b068f24ca9d47fd5a906308bd95117b62f48bfccfcb60001b62006837565b6200afed620103e8565b6200b01b7f464ae6012f7f0430bde46b1d6bc4e66f366deb61c56f6664926c2ca31d464c7f60001b62006837565b6200b0497f668fd7bdb0ba661b94256bcad9a47a9a96da8138489aef569e16d316ff9f38ae60001b62006837565b846020015185600001516200b05f919062012ef0565b816020019067ffffffffffffffff16908167ffffffffffffffff16815250506200b0ac7f8c5c716530264676481104f6891d0b0616154df628dbc1a511db1c80a1de91d460001b62006837565b6200b0da7fb1d8a35ae3999abf50a497c1c1055a3e03301e7e68937784be83dd7b648b426460001b62006837565b8360c00151816040019067ffffffffffffffff16908167ffffffffffffffff16815250506200b12c7f475d4f0c0b1bb6d7d3f50f5ba41b629e5762e7abde87aa2ea5b170d5446e41b960001b62006837565b6200b15a7f22f3299bc02507b8457f6d15d8cb10a4f842ccb259101ec24511bb423948df5a60001b62006837565b84606001518160800181815250506200b1967ffbd7de86b449d97b201467b2d2344e82add5cc4bdb284ca769921a79cccc7a1c60001b62006837565b6200b1c47f9d6820c4277b914be8ff084684ea4c5618d922d4becc71a4a74b672201e06feb60001b62006837565b8361010001518160c0019067ffffffffffffffff16908167ffffffffffffffff16815250506200b2177fa890d880893a53683003e221b80ed2fa2b26c9597f44fa0f498741b05b169adc60001b62006837565b6200b2457f87f8436ef4cb0bfb9269694b560a4774f26a4c265590340d26b342102fe5318c60001b62006837565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663178e4dc08387876040518463ffffffff1660e01b81526004016200b2a89392919062012bd2565b60606040518083038186803b1580156200b2c157600080fd5b505afa1580156200b2d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200b2fc91906201194e565b90506200b32c7fff29b2536a9185535f4afd231534d46344c5b49fe4fd75411ff09443c98028d760001b62006837565b6200b35a7f2a5e32a832132589e9fb4961f0201d3fa0da02029b260c0a35a878acf0b0676860001b62006837565b80925050509392505050565b600060606200b3987f4174d79904d22fa365048852b3d0807a4c586f396e1a7f132d4da9da3df7becf60001b62006837565b6200b3c67fd7665f8990049c559717a1ee7f319e4ac49a9643afcc2eed9e5570f55a5865db60001b62006837565b6200b3f47f1e29820e09e4c4ed5f86c566bb2015114a8cf03ce601a8c9468a0054e08d5dda60001b62006837565b60006200b40584604001516200cf02565b90506200b4357f9b9213cbfe980e640f680bbfbbf7b72648d93661c1baf6eec0ce21c625ff207c60001b62006837565b6200b4637fa69e445b1c68f576469427b727d07c4457eddbdbc7a0524ef77c90357d8a358560001b62006837565b806200b536576200b4977fe014fde45f2b62805636a92a9468b1cef5664ce8807a47f853bb22c7e4539ace60001b62006837565b6200b4c57f45d664473fba00c2edc0219ff47ab81f3759b976146b7e11257718787f50905460001b62006837565b6200b4f37feedb8ee38fb6ba3949987a4d7f27cbb2d77f42a59fb95258c478e89f3ccf00f160001b62006837565b60006040518060400160405280601681526020017f496e76616c69642073697a65206f7065726174696f6e0000000000000000000081525092509250506200b822565b6200b5647f4195ec545ed188c66172b27615df8ab7191669ef902880a93c8c5066b86d008e60001b62006837565b6200b5927fccf9221c3e7ac58923b3dfcbb624af6973f07388000ffad3f51fc8de34acbe8460001b62006837565b6200b5c07f58746f963345f10f65639c17eee1980c3ac22d9af63115d5f8b47561943394e660001b62006837565b60006200b5d185606001516200cf02565b90506200b6017fd4c39e45946500a3cb9ed726a37af9273e87a2e268b3f59e5f6a883395b67bac60001b62006837565b6200b62f7fb2254f6bc8c4f52e503461c5379b7edfa896373f188c6e44eb2fa7dd9216abd360001b62006837565b806200b703576200b6637f938010d7cc9835581c068457bb1b10aaf70d7b6689d9c1893f533c9014d3c52b60001b62006837565b6200b6917f01b2edeae9306577a4747ac0218531eeb09b600e3ca78f665776d0cf9f04aa7560001b62006837565b6200b6bf7f2a793e31fde667325342357186181c880fa99e3806ef49606ed54f6fba27ceb660001b62006837565b60006040518060400160405280601d81526020017f496e76616c696420626c6f636b20636f756e74206f7065726174696f6e0000008152509350935050506200b822565b6200b7317f4aa9ed963a1f6157f28eaa4af0aab7e37ec135ef19a76bf40e7ce5b4a192c13060001b62006837565b6200b75f7f18ce8695f5082908b90189db7e25998bdd681bb8b4e21156a993077f28942ca960001b62006837565b6200b78d7f75b0e62bc88987f4607ddb7c1879e8c765f670a62595943c20d7cea4caa5f34a60001b62006837565b60006200b7ab8660400151600001518760600151600001516200d812565b90506200b7db7fa5822098ecce29c38648641cda57f4019c0bca812f277e802152a2476d7e411f60001b62006837565b6200b8097ff0e5035ad17789c3ebfdb7c4360c484a5ffbc564e092384382c8a1332b0b4f5460001b62006837565b8060405180602001604052806000815250945094505050505b915091565b6200b831620104d1565b6200b85f7f8b2b12076dfafccf25219aebe159fad8cf7448a314808d9da587fd868194713660001b62006837565b6200b88d7fd16f298662048da7e390724882a45be874340f93146f96aaa3c5b10345210d3c60001b62006837565b6200b8bb7f4648cadb6ba7067202eab180803b1ef1cc42920361f87c59f9fae7a52774ef8460001b62006837565b6200b8c5620104d1565b6200b8f37f20d0a752d915f3da4c300e4127ca7fc421334a00972ebedacedb8c6a91b1379d60001b62006837565b6200b9217fcda3e34d7df329c7981f3ef288297d7f2d3fadd6dfdc587f8b2c1efc5223a03d60001b62006837565b826020015167ffffffffffffffff1683600001516020015167ffffffffffffffff1610156200ba4b576200b9787f8b77fdec422e2264a8bc91f6b48113974ee59a2504b4e3b48f22da5e48cd2f4b60001b62006837565b6200b9a67fced9fba449a533679fb95de68c7440da088a021521e8bc1b28d273bd0c3c0a6460001b62006837565b6200b9d47f68aa59553e50e72b9e96d95ca7c73a57f81c597521e3c8e857e8bfbaebe2ebe260001b62006837565b60008160400190151590811515815250506200ba137f2939fdb093768f55967742cab6c65a2af6f4775ff44e2f90e765fbed77381a7260001b62006837565b6200ba417fd484dd463a009fe89bd9e5240ba7d0527e5166b7d39ea5ab6de488846e874d4460001b62006837565b809150506200c574565b6200ba797f6be2c2bd53da551e14961a334f8a646a3d89caa64ae915666f3e916ce10996b660001b62006837565b6200baa77f73522520dcdafe486506d7bf885985c6b4d644f5e900f25700c2c71b79b5a6a160001b62006837565b6200bad57f6e302e2a639ba0f7a42421d43391eb8a7f6233677b78a4e5b934674478c3f34560001b62006837565b8260600151836040015167ffffffffffffffff168460000151606001516200bafe919062012fee565b10156200bc06576200bb337f64209bdf8a1471b0dc777cef2cf7cf24198c60a3891cccbe151d01f9f024b2e860001b62006837565b6200bb617f4ba7758d4909ba708adf60ecca00e7aade54173ee1ca35f942cb26a6cce5a63360001b62006837565b6200bb8f7f24d3a4e4030f3002e49b8bbdf34d7f81f03480cfb6e817718900bc855381d34660001b62006837565b60008160400190151590811515815250506200bbce7f33908663a105cabf5eda93dfbfa42ca7e75a9969e60bf7c796878cd89bf59d7060001b62006837565b6200bbfc7fa0f799f77e62cd2a6cda5709c0ab507769acb5b8311e94daba8ee052fb4a879c60001b62006837565b809150506200c574565b6200bc347fc8933766aec9ffbc03225a431bb714b768fc5f076f344adf1787e1adb16ea5b160001b62006837565b6200bc627f9168189dc8f478f774ee79146f3d3a28840a56ca303c946161e66285118214e360001b62006837565b6200bc907fba57baca21843eb574c786950b54e96dca59cc9b9e41e6e425688b3d2ccceb7060001b62006837565b6200bc9a62010142565b6200bcc87f2d680bc2d6a7b625276c9db7f583481e0e3624aebb75b952e8649633136833d860001b62006837565b6200bcf67f49d7b8048bf01141265f2779e57bbb0bc01cc061162da68e8eba326e1364ca8360001b62006837565b836000015160000151816000019067ffffffffffffffff16908167ffffffffffffffff16815250506200bd4c7fab0391e0805624158731bf5de077b0d451e4c156f054cfdac7e0f5728d9a48df60001b62006837565b6200bd7a7fbce8cfb9ea3c78a92f8b7eb966fa12ea88bcd1685c0a4debafba6a82065654f260001b62006837565b83602001518460000151602001516200bd94919062013029565b816020019067ffffffffffffffff16908167ffffffffffffffff16815250506200bde17fb5d0cc327cb5e00ec8399acdf96c4867f58a3840f28c048338dfd78b9171e35560001b62006837565b6200be0f7fc4314f2694020e134c42821203fe9dadca0d6ed2c3187074625f312af4cecf2560001b62006837565b836040015167ffffffffffffffff168460000151606001516200be33919062012fee565b8160600181815250506200be6a7f43efb3e1fa87362eedf962faea88a65960bd9182afe9fb089d5bffbe5d1f64c160001b62006837565b6200be987fde1f294e8cb2ff0ea240ffd8dba0eaef9bc69746d6c92ab174dff8e0fcdb187960001b62006837565b836000015160400151816040019067ffffffffffffffff16908167ffffffffffffffff16815250506200beee7f7ff1fabc188639117c9545385481ca8f55de3c1910f57bff3602d2efdea75fed60001b62006837565b6200bf1c7f85e7d826828808c5e76cca884deaf2521467d5481fd1d73a40caee2431c7521360001b62006837565b60006200bf3385600001518660800151436200af4f565b90506200bf637fc57096ef8c899b7c9332b23e1ee44aaa031cd7431164bd6a8430216f573ba86b60001b62006837565b6200bf917f27f9cb6942082bd64cb506f15d913a7b997bbfd3b55d7a121f95850ce7a9125f60001b62006837565b60006200bfa4838760800151436200af4f565b90506200bfd47f98c2718c5c6143cb1e560f564a757f2708c09a6438d44cecb69efdd433ad540760001b62006837565b6200c0027ff755c900503ef3ac6ce24ef346fb94e2c16a747100e594be36c8883ae869867060001b62006837565b60008260400151836020015184600001516200c01f919062012ef0565b6200c02b919062012ef0565b90506200c05b7f32d83b38de38172425156c33b81e2b51899b042872bc5b9b3dd1ce500d8bc29b60001b62006837565b6200c0897f99e88618ec3588e3171038085245a9a86f6aa0071058bff12ab3f303c13de8b460001b62006837565b60008260400151836020015184600001516200c0a6919062012ef0565b6200c0b2919062012ef0565b90506200c0e27f2816fd0c2c05d7bbf4dfd894a42f465cb7ecc7472d26faecd6ddea6ef3f8cb6460001b62006837565b6200c1107ff38e34bb4d4f27f90218aa0392d6af836633880bfd33d2d759d9587c7ebb89e360001b62006837565b8067ffffffffffffffff168267ffffffffffffffff16116200c232576200c15a7f158d7d5f1eef353637a523e49b4cffd7a69cd46dd26dc658d6903ac56c058d4c60001b62006837565b6200c1887f7b9ccfa3347fd3039fa5afa0ec0e0ac8aac081eb2ad646318808e59d72802a2960001b62006837565b6200c1b67fd0df4f488e546616d783422d2cdda0860ae683b926747124b641f254aeea47a060001b62006837565b60008660400190151590811515815250506200c1f57fa6fde92b811f57d5ba61c95f7131cf0749b8e9dd9297fee2e66965541a5a9edf60001b62006837565b6200c2237fa6110a4a941684df55f67379eaa3a041d1c60866813cde0ff6acb76404a6250860001b62006837565b8596505050505050506200c574565b6200c2607fe4e6ff1f01e3e04eafe5c38739289a5714c5a8b89a47f80b400c0e56c6e4e38060001b62006837565b6200c28e7fe99120fa4cafd278b397045eddec25f380e707a83a359bd4bc9848da006b27aa60001b62006837565b6200c2bc7fa82fecd8944944644e36087c533dfd616ce410703b24242e9323527c414555c160001b62006837565b600081836200c2cc919062013029565b90506200c2fc7f3eb40f403d67813f20df67f8a186dc22c2563ce1f7326b26b29a8733a0ff44ba60001b62006837565b6200c32a7f38c13dfce0cd0436edca0eaf019f2650ea35d8bc8e0b9b6384a2d66bde6e310660001b62006837565b8067ffffffffffffffff16866040015167ffffffffffffffff1610156200c452576200c3797f9fec218436df19d80b39bed96969358c9a93f88f9f4c5cd7104b6d153fc3268960001b62006837565b6200c3a77f660093c33d589264889c74ae892387357d07fd3c9469105ecf17d7c3dfa829c260001b62006837565b6200c3d57ffd53809a8345ae467307f02ee38f72c92bad2181f96752253bed85a2372de4e560001b62006837565b60008760400190151590811515815250506200c4147f2f5aaf191fafd704488cd882c4ca8cddf4acd5b224d709d6be185590a92267fd60001b62006837565b6200c4427f36a47784bff315170bc1402194a9e487f247f5b01c265cf352adbd5759014b8c60001b62006837565b869750505050505050506200c574565b6200c4807ff6390e3d957cb4e226fc5dcdf60558631d7ca20f8f1ff11031eebc23c1d3064660001b62006837565b6200c4ae7fdafc448ad932266e3253d7b2bc14814106818a7eb676e176f641572efed2a1c560001b62006837565b6200c4dc7f4517625fbb26d3cddc3fc52744b731b5090c24e38e38738e5709de56a0a6c0f760001b62006837565b80866040018181516200c4f0919062013029565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506200c53b7f7bdc50ad2e3d51ae8f256348236132f419332b67f73859473ecf6a15b4830d1560001b62006837565b6200c5697fbce34b81f3f4903f1573fee671074da3398f2cbd1cb9af573c9547f90ab6942e60001b62006837565b869750505050505050505b919050565b60006200c5a97f640a31f11678a877d7b612e302c97ddb149108d7a3d1da51d54d1f5ba6d5865860001b62006837565b6200c5d77f707f24ee5b68d52a226babe27a7526b1d02140fa08270c453f3831888a1671c860001b62006837565b6200c6057fbc510d1d6354c5048d59c8174a5a476ffccf64f1ea56ac5a67f3f7461f91e26960001b62006837565b6002808111156200c63f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b82604001516000015160028111156200c681577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14806200c70757506002808111156200c6c3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b82606001516000015160028111156200c705577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b145b9050919050565b60006200c73e7f49fca7e4201c2acf0f4547cf5a5523f34714d602f42dbf4083bf647e29e72e9960001b62006837565b6200c76c7fb99a4ce04bce09c2814518ffbceafb443fa086df78c298ce6ef600ed085f11ef60001b62006837565b6200c79a7f951fc1f60e9f6164b75360507481abbe680800406498e617a19e79b72e518e3460001b62006837565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639a2b5e6384602001516040518263ffffffff1660e01b81526004016200c7fd9190620128c0565b60006040518083038186803b1580156200c81657600080fd5b505afa1580156200c82b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906200c85691906201151a565b90506200c8867f5827b329543fd699b236fe7f6bac67f0405f1ee1047d14b8a269564ed768240d60001b62006837565b6200c8b47faf9da07126d42c683aca58f20b00d65e540ae3e1ed4dec5504a95f566d4b3aea60001b62006837565b6000815111156200c954576200c8ed7ffeff89800707e50411b7b8f8c3764649837a4df9c940551d57c97aab9188608460001b62006837565b6200c91b7fef0a2d4dfe033dd37145bbe9a5f7e621d84405e783042b994450d5eede8d249d60001b62006837565b6200c9497f73fc0ae715f07c5b8607c9871b9000a56a3b7e9a9ccafc1b57960da1704d2b9b60001b62006837565b60009150506200cb3f565b6200c9827f4af8027b749d1ac446952460ce887a9aeec767e709109ca18e2fefee7645b88560001b62006837565b6200c9b07f7d53ec7f7a94cd79fba6ea47a20acf44e42265424a519a2833eba7efc1a1341d60001b62006837565b6200c9de7f364daff91cd32b70b89c76f94cfb8fd197abb9da13ae2ce5edb941aeaf31958a60001b62006837565b826020015173ffffffffffffffffffffffffffffffffffffffff16836000015173ffffffffffffffffffffffffffffffffffffffff16146200caaf576200ca487f6cd4bf3c3894300446c530eed6454ee53011b32813418b70f2bd68e830218f9360001b62006837565b6200ca767f756a414140eb368d4ef10347dc76f7135b3764c7ef87ce267d940df55d4b661c60001b62006837565b6200caa47f0b63e749caa275926d143054b4f06b9679f36a081542ec98ba8b6edb125c02ca60001b62006837565b60009150506200cb3f565b6200cadd7f12cce1e2e9887d201e03e4e55d6a66a2d84ac905e39b75fa520862d5d3f963cc60001b62006837565b6200cb0b7ff7b659137f9a82a9dd82ae43082d53b66035d0d68ce97adddb81c08bb66e2b5660001b62006837565b6200cb397ff4029a7fd47bc8b0d8ec6543f652b827c3bb40ff40cf9082014efdc3a922b44660001b62006837565b60019150505b919050565b6200cb4e62010142565b600060606200cb807fd5858872bfdf77bf5a2dcbbc6ff9293716c1bb95ccc47a4c780dacd258e4ac3560001b62006837565b6200cbae7faaaf7f20ad2ac01e391baa537ad9e636e1566a34b7bacedde30dcd844c5baa1760001b62006837565b6200cbdc7fede0f385cb4d59c1512bc7140c5b329119da6edc03ce85b47c290ddf21c2dc8560001b62006837565b6200cbe662010142565b6200cc147f420c81cce81072cedbb82b38a07b74a649cf6fffaab96dc911bcb0a2fa969a7160001b62006837565b6200cc427f78a877a08e459f4c3f2f3199cb47e0db14ababd256f12f706f1fb43b3772b2f260001b62006837565b60006200cc727fe0e7005b9ccd29ca3397cc9e47dd960659e4ad17e56c47b541c5e0946c6758ff60001b62006837565b6200cca07faa11feb05717cf333478a53f56d6d799ffb7d44f0ec92be82cab2dfa8cfa11f160001b62006837565b60008a6060015114156200cdcc576200ccdc7f67760af9d2697efbf13d52765b79026bc1294fd678154fe7ba8975569103f7da60001b62006837565b6200cd0a7fbe5b8e328fce12022e1d9853e4d8643d1ee01349905771d0accabe30dbc9f3a160001b62006837565b6200cd387faf6baf7c96cde38c27ff6a2ae5f3b13b90dd8d749539682a14d4f827ec0c60d760001b62006837565b6200cd478a8a8a898b6200d98b565b50809250506200cd7a7fdfade31931d960103fd47b7a34f464d9d028dec712bea6cf26cc2123aeab499660001b62006837565b6200cda87f43cd6147df2e50a64bab1ed62c487719312852d27894b8c311af5609f94276d760001b62006837565b8182604001516040518060200160405280600081525094509450945050506200cee4565b6200cdfa7f03e0e6232f88f883fc7f61ff96f58bfcebf45259ea5b4b072ac8c969fb2b3b9660001b62006837565b6200ce287faa780bc307c4736dcb62edb53fbe0fc295af4480706ae35b76e1e0dfb46c1e9160001b62006837565b6200ce567ff8d2369c3b15e526215ad9842fc22f8dae20806f2263fb8a99fb309b50802b0760001b62006837565b6200ce658a8a8a898b6200d98b565b80925081935050506200ce9b7fe8b749e35e079116d26d837aaae7d8b842a88c402c5f8579b7b6cc9c57e8537b60001b62006837565b6200cec97fb3e485ab242cb01e5ecba9976121eeb5a564bd62a00ea7fc1e95f0f80fc4782260001b62006837565b81816040518060200160405280600081525094509450945050505b955095509592505050565b600080823b905060008111915050919050565b60006200cf327f57d68698fcad0a27f869b8fb3464cfcff294ac6b3bd4ed8cabdf7cd02d6709d260001b62006837565b6200cf607fc2b5ec6150d1ffa3beb4ae7ece237e1049a03284da76ea5afe129e78433830ad60001b62006837565b6200cf8e7f034a81ba4815816821bfb724110fbe7a6febb44685758fd922eb5ba51d2534d460001b62006837565b6002808111156200cfc8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b826000015160028111156200d006577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156200d1d7576200d03b7fdc1a06131706e174fbaca9381287839a29f1f73bb3e3c450ea3d3f9582e40d6560001b62006837565b6200d0697f810b90f23aac409a77f9d8726e0f2e2233622ac8cdad0f6e08e96b1b58616b1b60001b62006837565b6200d0977feb4eaa9a044cc74eb659526f07219b69833b82a30615f8644a083f0a6776680f60001b62006837565b6000826020015167ffffffffffffffff1614156200d143576200d0dd7feeb4d8066bf0fa50158a6aadde4a12235b7c6f90ecaf17a590a74f7d0cf057fb60001b62006837565b6200d10b7f166d8ecb8e9fa5c93f513c2e1efc382b3211501bc5fd734eb2d91ed7f15c26d560001b62006837565b6200d1397f232b9581b39086a40304336caf57e7925f9118899e4a36d34bc4caf2d37878d560001b62006837565b600090506200d80d565b6200d1717fb7f2b58908e5785f602a4f33cc1158979cefbea3c9280a4a56e1e01c503ae45d60001b62006837565b6200d19f7f38008d5deb9b7f76f0e651344e0e83bb4794e5ff1e8d7e03f47adca10a1791c460001b62006837565b6200d1cd7f5af21774edd604b0c82eba7e8179d89bdcb80ad4c5c707160cebab99bf7db61460001b62006837565b600190506200d80d565b6200d2057f1c1226258493dba5d48d274415a96999b3819064c9ea4960115484282d1817f760001b62006837565b6200d2337ff4755b0816216d7f5dbc8644ea8e33eef7db9ad19efc17d759961c9abb5c44f460001b62006837565b6200d2617f716dea0b379a4f15ffd28984b4ec3b9cbe0e670321d4687801d88e6a08b8e1ba60001b62006837565b600160028111156200d29c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b826000015160028111156200d2da577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156200d4ab576200d30f7f8e4eff695649a1393e45f4bc17f233c62f6f63a29db3e846b096c353d05283d360001b62006837565b6200d33d7fbd540fc8adeea36f5e908d68f6e8acf4c0adb5e5aa1266aa06ccab74707b5d0260001b62006837565b6200d36b7f493b970154e7d90535d75b3361c311f2ed587a4c3604d6ab6ebc5f31e7d5553060001b62006837565b6000826020015167ffffffffffffffff1614156200d417576200d3b17fe46c4d78cea7545277a2dbe6416151af4afa4d11823adf2d723dc62fd0648b2f60001b62006837565b6200d3df7f4aa529311c10cc0188c3ed59bdf25467a7ee5507eccffacadff49e6cd4ae715960001b62006837565b6200d40d7fa5c0bd5f6650991d736e8667e553751a1e7637bdf2098e731e29ca0bea77391360001b62006837565b600090506200d80d565b6200d4457fec206abbbd63efc34cf1918dff70fcb1d83fb75d836e253e72abd73b117d155860001b62006837565b6200d4737fa1fb4590440ea1e025b532b13cfb718a0de3fa57df3df38d8b467568f022948760001b62006837565b6200d4a17f27c3d302441e5195ca19a02d82fa59824fe093556687f8ae55a57648510fc4ec60001b62006837565b600190506200d80d565b6200d4d97f1cffdb44ce6304e6deb0c70f9147770b123ffdd5bf78043dfb6f674edf7e815060001b62006837565b6200d5077febdfe219448333b94cb6e712f4bb8a7aae7cdea951ad3476900e5cab0e3ae1da60001b62006837565b6200d5357fed2c9e07f5bff3d50def6afb225ffbe1dc191ada55186f0d601bb3d1ad372c8160001b62006837565b600060028111156200d570577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b826000015160028111156200d5ae577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156200d77e576200d5e37fd60f4d1134d922705a2d79c49cbf037bd3782dd508e5d6663bf073e081b1394660001b62006837565b6200d6117f58ef133e8dd56dcb66049be3cff8ef58156feb8b1419165edde9ec742c71dddd60001b62006837565b6200d63f7f1594eb15ffb20ca37341d0c203bd6c2ce68e21f52826b7abc2c6c663b76a1c0460001b62006837565b6000826020015167ffffffffffffffff16146200d6ea576200d6847f741f85b7a5e3d69a55c4c0e2f554fbf7a4a3582b55a0851d0d5bc10530ed062860001b62006837565b6200d6b27f7222f9569d4ba14f91de52c257f3e7b1766bb82805067bbe7ae76d9d2687f50660001b62006837565b6200d6e07f0f6c5070ee2891fece2dca076f1e02708d6b23edd6431a456d2899916435a48c60001b62006837565b600090506200d80d565b6200d7187f87e7b9c237a36f3ac93d30faf96a5a739dce3e086a3d4644012599e218034ff960001b62006837565b6200d7467f51a6b0dc1aef9d0e76e1ed2c8d0599e5fe05c5a50e57e58e2dd67f6eaf5cd03560001b62006837565b6200d7747f57bb7d994f502b3459b7c53e209f60ed5c05a3160373ec45e0957f6b2855abe060001b62006837565b600190506200d80d565b6200d7ac7f4141456260ba6a89adeb5cb1a861def18cf2c64e36dc054bd82c1c0e873b16e160001b62006837565b6200d7da7fc31c134150501117f6eef6f4963c77bc35f554d1687a2158546991d5ef98d93460001b62006837565b6200d8087f8f04737cc2d9a35de350767ab950796e42dc8fa372c42b3049e966a1fd839cac60001b62006837565b600090505b919050565b60006200d8427f2baad1f3d237d57ea626a703235c5d488cb03e67f7ee11d62a78e5d06542e8ec60001b62006837565b6200d8707fc4af6ec68a6aa8a20d0d8b96fa38bf45cfab2990f9e0e4632e334814ffc8113660001b62006837565b6200d89e7fd4bab063701bbe39db65e03ef6203996c59df8ab33ec37d8cd44024f9cb1d9e860001b62006837565b60008260028111156200d8da577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60048560028111156200d916577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b67ffffffffffffffff16901b1790506200d9537fdde3a8c9d8ea6122931e1cb13fdb1e6411d5c2dfec12dd7381f6cee821cc918a60001b62006837565b6200d9817fe0163332134a590fc79fd43044cff7482e187c04560ff79dcf68ba059856bf6460001b62006837565b8091505092915050565b6200d99562010142565b60006200d9c57fe74b31fe6f0b6ef8c204e04ea976e28c541fca4d9eefe65399782affe310d0cd60001b62006837565b6200d9f37f3dd743f9ab3766bd9c1870bf3456428e0c76be198dd52b1a960e4578a2109d5d60001b62006837565b6200da217f7862bf68022d85c3cac8cf7d063d0fb914a81b0aa23a7739306a2e3ca9089c0860001b62006837565b60006200da517fac006b871bc90ad19b70aa10df10d1e74cc2f9654216c035a4bf418e25c58ed660001b62006837565b6200da7f7fd00e7c07d2fabebba07df6f60148ddffed633a5cd8e9db9f53fc9b3c5b049ad560001b62006837565b60006200da9860008a888961010001518c8c8b6200dd28565b90506200dac87f1651fe3913197659f038e3ce83ecd700b87e93033c769c0ce26391dbdc66650160001b62006837565b6200daf67f51b296c4cc554a307e28dfbe3034087608d1b7609f6490995173470735322b7860001b62006837565b8060400151816020015182600001516200db11919062012ef0565b6200db1d919062012ef0565b91506200db4d7fa706df7a05fdee9b4217cea2d6f8e725a619394040b8c993a50a271f894073f760001b62006837565b6200db7b7f69a90a14c78fe15e74ddefb613a283950484e8a0cfdb06169223832a4f548a7260001b62006837565b87896020018181516200db8f919062012ef0565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506200dbda7f24013219029456f7b5b52e7604132264d0944895641aeecfcb43a8123b4783c060001b62006837565b6200dc087f52d4cd8c80501e6bda303b259cd23ce4e725a8706126c70ac4ff33aea62263bf60001b62006837565b8667ffffffffffffffff16896060018181516200dc26919062012e93565b915081815250506200dc5b7fcdaa9186759eefc1019ee3095bb9a897a832f22ae8f40b43b5046207e95714cd60001b62006837565b6200dc897fc455bf3c089ae55097dfafad03fcae6443dc65d0ebfd1aa371a80840d942015360001b62006837565b81896040018181516200dc9d919062012ef0565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506200dce87f4c2dd5aabba6082f8d520e15439afbd6bb5a4808e5a24b1e6683988270289f8460001b62006837565b6200dd167f5989e99e3bec73c976b68550c8ba4c2420e9750a030a3882baae464c1aef89d160001b62006837565b88829350935050509550959350505050565b6200dd32620103a9565b6200dd607fb819e956f7d3f8b3ac86dd6d2e0a576215dafa32a26c3687cc0e98c1b2baba5860001b62006837565b6200dd8e7f35b3bcf316357f18aded9443bb0caab910c86a32f5a58b32203c6f49e302bab360001b62006837565b6200ddbc7fb9296c28d199f92a3470781b397bddcdae00ef971ed51cee01b2f99634b37a0460001b62006837565b60008790506200ddef7f48e24400d5ae468542840342c16bc996d63460a48fec973ea2fc4e0416d2b3cf60001b62006837565b6200de1d7f5113f64345784a463b4365e23122f93e0c7b33761a70efdcca4d1eea5be09bfb60001b62006837565b6200de27620103a9565b6200de557f36caa07a678a0fe3687dd8c4bbe2e1d8e31fe6b1cc63a0de37d8c6fba7a7883960001b62006837565b6200de837f7db483faf0a41999392d1de156ab97dee3e2420f033dcf45ba597e3dc12eb2ba60001b62006837565b6200de8d620103e8565b6200debb7f2de6fc4703488fe7b23e94c660e4335e189e8ce7b7650882a8daeb5f03de044560001b62006837565b6200dee97fd3124bf7494c52ea314eace23625f694d0e46308e600873f9675d2ad814a38fc60001b62006837565b826020015183600001516200deff919062012ef0565b816020019067ffffffffffffffff16908167ffffffffffffffff16815250506200df4c7f7a2790db87aec046d985b76e179d0afcab47ecf72e6716e7285c957e4e2a8f4960001b62006837565b6200df7a7f1c28df21f355e932600341ae1ab0e979595166cc4759e2bee1c84935dff418d360001b62006837565b8860c00151816040019067ffffffffffffffff16908167ffffffffffffffff16815250506200dfcc7f382d9cea3667c2ce356f17f5e66afdc355ec23f90afd09d0ab1737181589dd8f60001b62006837565b6200dffa7fc963c01a754e1fa09a6aaf0e1f6b3d8a457d763e42281beb36017b6dc9a2ff7760001b62006837565b82606001518160800181815250506200e0367f88476a615f63aaf7f79dedd08c2c9ba9805dacedfb8915b11d41f2db982351ab60001b62006837565b6200e0647fafc6dd75b1efcd3a4738a2936d29b80e63b94b3210a6eff73c7827e8ab0356b260001b62006837565b8861010001518160c0019067ffffffffffffffff16908167ffffffffffffffff16815250506200e0b77fc816be7d302ba8e42d290c2b2672e79656dca132a6ca750bf61cea495983877f60001b62006837565b6200e0e57f1a5e80b6a8882e13e0ec851a402df7329531658be152c590c107f19250e60d6c60001b62006837565b60006200e1157f9559f55a39ec2e701c8695e91d80ebea2ebf88a019eac650d8a7adbaad6e0e0360001b62006837565b6200e1437f63b2178f1fa8594cca89d72c8717b7b489305603005ef59f6aa59884e06f421860001b62006837565b600060028111156200e17e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8c60028111156200e1b8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156200f141576200e1ed7f2e4b7b6550f5af95441404fcc33f95be0f23a4dcf6be1f6bb8c99964b3ebc58d60001b62006837565b6200e21b7f902a88340dfa8d0dd87a99c3193eae807077ec3c231be54cdb4615fae2d816f460001b62006837565b6200e2497ff5fe8deb6a868b267f7e39561dc194a25e072f50641fcfbe0325707b0fa3747860001b62006837565b60008767ffffffffffffffff16116200e5c3576200e28a7fb6d3f960de9ae60667732361c5c3a6bc53d1df6d9dd497f2f8ac03f765956bd560001b62006837565b6200e2b87fba1688266914405c38f10eecbf64aec163967d560c58e05369b2185c615e738a60001b62006837565b6200e2e67f95512e103faf53d377c85f67cc5c0eb3b049c6a93292be64c722b82fdd08f2e460001b62006837565b60008867ffffffffffffffff16116200e38d576200e3277fd1c581cba581dc2d6961e54c8aba24189d3aee0aa4a3d04702a34e6b654480e360001b62006837565b6200e3557faa48902c9363282a8b9bfd742299c280f5c0dfc976646d404d7a284017faef5d60001b62006837565b6200e3837f131acfcf00a231b5ffd2d5f5e8d00fed6d294d30deb4360a8648ff465931748360001b62006837565b600197506200e3bc565b6200e3bb7f46f0d80f81e3c20aaf670f6dbcb3f8ad2096c136cb13dd606ca07ff15c10136060001b62006837565b5b6200e3ea7ffb42e5b72b8cd5555da7973c66b857e58d0eb101d7dcf5db23592b1c2b7e247a60001b62006837565b6200e4187f36772fb5fd862cb2f54ed6a6df3f3a5a1dd0c1dd095e10ebd188611936647b3a60001b62006837565b6200e435828786606001516200e42f919062012fee565b6200f52d565b90506200e4657f67aba781cd41ed1ed4ad82d21405aa7a87f3a0ff7f0c55722aa3fbbce04dc21960001b62006837565b6200e4937f1a4eeb626a01519a1e285eca09087e4159557f17badfb3a2a2773317b4be45f160001b62006837565b6200e4a18a828b8b62000a9b565b836040019067ffffffffffffffff16908167ffffffffffffffff16815250506200e4ee7fca9ea8d27ec3a79a0322e81079843f15736c9a4f13c2097e2c6eefdfb6b16f0660001b62006837565b6200e51c7f9fec916f8ae8ba0d6f060e33b3af3b9b59ae522179d6f4f89ac33ae5208a311060001b62006837565b6200e53b8a8a8a8988606001516200e535919062012fee565b62006779565b836020019067ffffffffffffffff16908167ffffffffffffffff16815250506200e5887fdb2662122ee61227566e5290f9697456faab6eadf8ace8c16de0e465de7b4b0260001b62006837565b6200e5b67f191c3d9872bc64dfd76c010f4d7bf4db906fd56cb3f3cd9abf49c3864bdbe25360001b62006837565b829450505050506200f522565b6200e5f17f8fc3dd39657e36441c7db66c80f56380f4d22bd4b64fc70c2dfc6c233fc6ae1b60001b62006837565b6200e61f7fcc28f3857146290a561c4e52f100387348c8a3588591d39c1edc96bbb8a7717d60001b62006837565b6200e64d7f32ebc03a52a30b583cc980953e1514fb3be65d80cac25368f44a7a1e5329b12960001b62006837565b60008867ffffffffffffffff16116200e9e3576200e68e7fb2f7d50f298a9bb090f7756bff4eee4af9683f77c0199274bf5da3fa1685cc6a60001b62006837565b6200e6bc7f37dcc96b7c3c6b67c0634996a531c3b2868b138f7a025df39c92529e1c22700360001b62006837565b6200e6ea7fe12d1bc5380226c5e8249cc277d71879579ec4794a4d19f5e7ce335b71608a0860001b62006837565b60008867ffffffffffffffff16116200e791576200e72b7f727057728f8ccf9d389bc2ce6012a61dbdcb394f4d61d865bfdc8b5f4a1f1dbb60001b62006837565b6200e7597f58ab2cd5cbb322e63dfa6424d629de810d8d057ea525f9fc71443c16637a7dc260001b62006837565b6200e7877f91ea1a062fb8065861d33a83450f34a3c39c8e2bdb763f8fe10fca856aa1254960001b62006837565b600197506200e7c0565b6200e7bf7f5066aacfc1c8c6c62147c27baa0838bf46bdbe917d17cd615116c26a6bc73b7760001b62006837565b5b6200e7ee7fb2efa8b11bbdfd646009f347ff20926a3e4b6ca06fd4c0a5f6dff529403466d360001b62006837565b6200e81c7fc247fa23a9a276f5f07304df467cf0f7457e9092a80d6c4a09c090bef539438c60001b62006837565b6200e832828867ffffffffffffffff166200f52d565b90506200e8627f7dcc825a59c78e884cbfaac4723e7eaf865b83d53e070e15650dd0e3f1de2c1660001b62006837565b6200e8907f9b85cc3588d53ea8292a3b46b954510d8d419792e64c8e583587b18188f0ff8260001b62006837565b6200e8b38a828b876000015188602001516200e8ad919062012ef0565b62000a9b565b836040019067ffffffffffffffff16908167ffffffffffffffff16815250506200e9007fe27650d6f241462d92da391fc51c991bcb32f364efb55edd2291602093ad89a260001b62006837565b6200e92e7fce084fd33706007b98365f34b503e22d758067d09fde7df6f8e28a37149140ea60001b62006837565b6200e95b8a8a866000015187602001516200e94a919062012ef0565b8a67ffffffffffffffff1662006779565b836020019067ffffffffffffffff16908167ffffffffffffffff16815250506200e9a87f7443233a11cf2bffb07d4f3f891f4080c2029cd0e2b1f41006514e0a976b660b60001b62006837565b6200e9d67f673a58f26b9013a78ded574cb1920690ba9ad5c1b589b35e647e1ff1018a314160001b62006837565b829450505050506200f522565b6200ea117fd4b4ce6401e27787868a18e8a588d927b3464a499a6f47904dcbf7f4f382298a60001b62006837565b6200ea3f7fbda7c22c2fe26d0b2dbdc4ad93508433316dee12b2a3dd464f1d02e03a5517a460001b62006837565b6200ea6d7fbc648f75a10d63e202aeb1a9d696a7364c52c91241086ea0f67f80e76c65015660001b62006837565b60008867ffffffffffffffff16116200eb14576200eaae7f55b4668de408e340ab4e66174ba14cbfe21b0a102a2815517b509d414791883f60001b62006837565b6200eadc7f252ff72dc6535211c282e449899ec8f6276eeb1bfe6bf2f4aa09bd983ccdae6c60001b62006837565b6200eb0a7ffa4a31eac1f41bc92dfaf8ef25b47b3297beceeab969ae3b42d7de1937318df760001b62006837565b600197506200eb43565b6200eb427fb46bc752216abf88209d01c3bdb729196e4c9bcbfd912e66f8f064056a51cdd260001b62006837565b5b6200eb717f70be17a53dbb492da6286bee022a7b4a84ae64324dcb4a7f80d43f796171fe1860001b62006837565b6200eb9f7f960dfce3a10484faa93678da396c583e4c112f82c4f12d7f974caf964471132260001b62006837565b60006200ebcf7fb6ea35027ca84578c0729e215f68756dc07ce76144e5a3373c569fa4d93a33c660001b62006837565b6200ebfd7f6dbedc7e6cf42c1abeaea2ddfef7fc78faff0f9a8d8eecc4de14dd2d78a5eab460001b62006837565b86856060015111156200ecae576200ec387f9f158184bc5505e83237c174a8ee2fe2102177c68baf83a972794b714e2d174460001b62006837565b6200ec667ffaeb1005c22841162982a84cdffe29987a8f2dbfb43142a3e4b934d0ae44293060001b62006837565b6200ec947fa3f46f866e32e3148901d84aa7e1f7982ae7861313a4393ebbb5463522f29cf760001b62006837565b8685606001516200eca6919062012fee565b90506200ecdd565b6200ecdc7f398bc6024872b3a247370a39568e1754f237946f7e9879768d9896fa8f64700260001b62006837565b5b6200ed0b7f1349561c04e72b6539cec4df3396e956aa47b4518746296ad2c8ff967ce83dc960001b62006837565b6200ed397f8bbb1a37307352cf1bb1ea9f5949727f2b5aa09bc234f125635564e4d54cd3c860001b62006837565b6200ed4583826200f52d565b91506200ed757f8c9c7dd860e98f2674bd43813cf97c18441ede734957926ef853e48bde6b46a960001b62006837565b6200eda37f6f3e68ea405a5e69a7547b07560e2a1156a7230a89c28ea264d91ab125793bb460001b62006837565b60006200edb38c848d8d62000a9b565b90506200ede37f19046806a87ba0016468db6bb110317dd0db157fc2ccf60807d685ae507173fb60001b62006837565b6200ee117ff214691087e7720aced487ea2577ea48805c49dce21cbd8d3883ff377529bcd360001b62006837565b60006200ee218d8d8d8662006779565b90506200ee517f24859ef8a51fbb449d4aa102cc073058909dbb964b1758d3e0857e1663f8203e60001b62006837565b6200ee7f7f576b658e13bbf63b0606ba5ab17ffc58a52548413b4ebd148d1203faa4b4761d60001b62006837565b6200ee95858b67ffffffffffffffff166200f52d565b93506200eec57f9632956edfd7fd4a880979efc8845dce85f334eabc523ff7676b9c593bedcc4660001b62006837565b6200eef37fa66e54b58969f6332e7cb085840322655d1c5d6ba7f268aa6b6e0a1845b4a30160001b62006837565b60006200ef258e868f8f8c602001518d600001516200ef13919062012ef0565b6200ef1f919062012ef0565b62000a9b565b90506200ef557fbc3ea8fccc06471c8e9b53de1b381d6534158e1105949422f3f992aad5ccffb760001b62006837565b6200ef837f3bda1521a1fd01a0bd5049dc7668a03b066151ed138029780d592748843dbcc660001b62006837565b60006200efbf8f8f8f8c602001518d600001516200efa2919062012ef0565b6200efae919062012ef0565b8f67ffffffffffffffff1662006779565b90506200efef7f7b5d61741e77b270e019cd16acbd995b00e1e945054241e4d4560165d951b0a060001b62006837565b6200f01d7fcf92e54ee2fa6d69ca8296ad242d8cf48e48713b838e1485e4b3f997c77a23e160001b62006837565b81846200f02b919062012ef0565b886040019067ffffffffffffffff16908167ffffffffffffffff16815250506200f0787fc45e8bb92c909cb7d3ef019a7b0d16e4f7d5621edee1528ca587cb49d4ea397a60001b62006837565b6200f0a67f6a0d7c87794da5b9906684e1f8a9b24ced6a49f5362963ef9b804ca81ad25a6160001b62006837565b80836200f0b4919062012ef0565b886020019067ffffffffffffffff16908167ffffffffffffffff16815250506200f1017f3a268f12c3a070be3509dd0e53bad6f505aec9d8c0a955068320a302570b4d2160001b62006837565b6200f12f7f26e7886a450b4c85975948b6c24c997d31745508d27eb55c5af4de184ca084e960001b62006837565b8799505050505050505050506200f522565b6200f16f7f6df2a58d494163bc125bac4912771d55e8f75075a0555dd95150bd1f9b841b5960001b62006837565b6200f19d7ff86f551c46cf1b8f4a665ec62fe49f7b8a82224e6dbab6eae7775e4d2007428860001b62006837565b6200f1cb7f21390cce13e97bf0c1baec95e79e7175cf9bfc2cc8fa4dd0b1b1e9bd3cf5059560001b62006837565b6002808111156200f205577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8c60028111156200f23f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156200f2f2576200f2747fa786573be0d56672602f56e66ceee2706344302bec6a0264f19b42482cdc450260001b62006837565b6200f2a27feeb17a4243243bb8d2cf636bdcf6e0cddbadb8e2662443a53590c3631f83639e60001b62006837565b6200f2d07f732e968a0bb3a1e07d9466d724ea9133cd0fc3b7211cadf2724a03e8b62f34ac60001b62006837565b6200f2ea82856020015167ffffffffffffffff166200f52d565b90506200f321565b6200f3207f7b0052324d73242f2ef88b824cd9d3ed91ef744549d99683c2270f38e7f747d860001b62006837565b5b6200f34f7f4cfbd4ba155c369dc24d6638851a8c695447ba911e119713760f79e1ac12c6e660001b62006837565b6200f37d7f11eb2b76f340d956c2a39387b2e4fb0deb0b9f78eada10eadb89a638f356111460001b62006837565b600160028111156200f3b8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8c60028111156200f3f2577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156200f490576200f4277fedff899a6a456ea79c9d0c6f47dd9bc84f34547ef92a7cd164686ad3217358a560001b62006837565b6200f4557fa4525607317a1e505d81c2dc056c55b5f060424c33d5cbcb86f66b91ca24cff360001b62006837565b6200f4837f75e1f233beddfb8ee49b594f021450bad065eea472f39391922137ae41e2e9c560001b62006837565b829450505050506200f522565b6200f4be7f5914ec8eac445bbae4121d8a6355c3101d462f30970778fb1b25cba9afb980e660001b62006837565b6200f4ec7fcbcf7e62b1d0332b065126bde2b7c0c39e7392ab4dc7dee2f21bb4c4e64ff42260001b62006837565b6200f51a7f5645672132b7e855909beac423f0c2a36e65c556a87ae3bd142ea9c56d4315bb60001b62006837565b829450505050505b979650505050505050565b60006200f55d7fede4d70b49c35d05d89cab1eb201cd8079e5baf126a3f6cf32c044900107770460001b62006837565b6200f58b7f7028ee2546f34ab5f776f3505e2d15d6ad79076aa98074fd284f96b1e2f709f360001b62006837565b6200f5b97f86fcab32bee16b145469e10b12b0274ee27ad5615eabe3b6d3b03d2f2decf45160001b62006837565b60006200f5e97f7e46b2e199e6e3d0a5d3a0f0c0112b0af4b15e356a4c27adbb57827948aa5da060001b62006837565b6200f6177f6b9c2a753a6ff5bb46817432d1576a230a2e912ca433e664cc4585dfc2eb289160001b62006837565b6000846040015167ffffffffffffffff1614156200f8bd576200f65d7f91eff5e5a1763c4bdd6ef151c07b8eefda08bc212bdc2af974f81b525661101160001b62006837565b6200f68b7f323630fe7dbd5ed0cedb83d1ccbf0a24ebf62f20a28fb3a1cf4d191bcd099f9d60001b62006837565b6200f6b97f21c258797ce6b8c097c17716d6dcf1cff114e169f17f61a262b1af15900a15b060001b62006837565b6000846060015167ffffffffffffffff1614156200f7bb576200f6ff7f777aa5e0fcea5c1496e739964a4b56eaea34ab049606f00d7ceb1c5f08db7a7960001b62006837565b6200f72d7f366234c829f5eaa6c295df21474be129e5dea607f098287a4283797bd686e17f60001b62006837565b6200f75b7f883fa25ddf56f7523e01ad5852e30a85e7eb6cb4239470982e3fcff23034865860001b62006837565b600060028111156200f796577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b846060019067ffffffffffffffff16908167ffffffffffffffff16815250506200f7ea565b6200f7e97f3ca6bf4c71a2614f6daa721a311f4332c17248d841df86a736d391f81969214f60001b62006837565b5b6200f8187f2ae3f0d5d508fd244a02a6dac8fbc91391b51086f527c74b5d796c1a9c6be8fd60001b62006837565b6200f8467f34f864ee05fb3c98dd3600ed432d2379a21599d091e0a541d03c95d63e59701760001b62006837565b6200f898846060015167ffffffffffffffff1660028111156200f892577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6200f9da565b846040019067ffffffffffffffff16908167ffffffffffffffff16815250506200f8ec565b6200f8eb7f70ff957e63aceb4b83d9c177c92f2bfe0e3537f30a507546538cfaeb772a82c860001b62006837565b5b6200f91a7f6b4bb3098e8489bcc5e9dfb3ab17c9b188d76ee2a192ae69e8ddc12ea02640a660001b62006837565b6200f9487f5143afdbd0a781f387092f9572b0c799cdb881ae6382a76d6f1bfa252086666860001b62006837565b6001846040015167ffffffffffffffff16846200f966919062012f35565b6200f972919062012e93565b90506200f9a27f692e2de839c7faa607dc150a1f2dd32fc53bf5fb4fc62037090666b30355bb7960001b62006837565b6200f9d07fd1674eb07f1dd5e5c65d871713939925b9e2517b36cea3331bb35e290cc6fc7c60001b62006837565b8091505092915050565b60006200fa0a7fb830271382133fd364bc74de07c5a279aa57a9cb0d70664c549558d1af47a61a60001b62006837565b6200fa387f6c104c467601d334f8ea1b7a7ee6822db68de4ed65f06a06b512d4c9b5e79fea60001b62006837565b6200fa667f242fd04f09370b3b3b67f0aac7d0d025883a7cb87d265371b0f129bbe0e3b78e60001b62006837565b600061438090506200fa9b7ff22f8fecb2cce7648b97a2b4db3bc200a948551057cacde47d769dc171f4738160001b62006837565b6200fac97ff2ea72b46cc0da9492f388d3579cd61f458811bdc7cf00afa9b49f9fd30d85af60001b62006837565b60008190506200fafc7f57448dbab44bd3c6f1ce9da10d02dc4cee9110f177f93de46cf7c7cd70fb734960001b62006837565b6200fb2a7ff42a476f78e9c54e706804e233a9f8dee7f6f6600561fa85349a8977cc9f760a60001b62006837565b60008260026200fb3b919062012fa5565b90506200fb6b7f1bb577b7372832788c8fdf0d0bc3076fe10184b7b0b902c1c40c526b887f89bf60001b62006837565b6200fb997ff82d75009c318a2e658a83e3dcb880fb4121cdd8221e773a5d1a7bf9e1365d0c60001b62006837565b60008360086200fbaa919062012fa5565b90506200fbda7fff40c3f2e702226cfdb192602c176a30c8df389d5016273db259b4de747452e360001b62006837565b6200fc087f33a533787b6b557a4dca2d20abc9d497b720bf49d3cc19dbeeba6371f3c1790460001b62006837565b6002808111156200fc42577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8660028111156200fc7c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156200fd1a576200fcb17f29538afb836b0fed861ff48782eeecde6e5982f060adef66febc0897237edba960001b62006837565b6200fcdf7f57e032a7227d3572c97e2123deec0a95a88b76fa8510579c9e2974aa9d3a791960001b62006837565b6200fd0d7f98f28d82d46b67a8884b894545d573029f229c4067dd683a7c6864b6172a88b060001b62006837565b80945050505050620100e6565b6200fd487fbe3bfdff3147df8bdaa628a724f74430ed2b014b06770ed27d8ed4d68ac57adb60001b62006837565b6200fd767f2bcd36c14906abf9e71f74580d29dbdea740d62aa77de189d8fe40b35c8af27860001b62006837565b6200fda47f8c85832451dd1fcef18a543e737bb38642e6c9412918fd0a9798745381c9b5de60001b62006837565b600160028111156200fddf577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8660028111156200fe19577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156200feb7576200fe4e7f39ff583e475c15bce0d7500444da225a42e28c181faeb16769fe57c0964e44ce60001b62006837565b6200fe7c7f475c62ac82b15ea5d3cdf3224bbfbe9c610f2edc6ff71d47e421093c7f45da3b60001b62006837565b6200feaa7f19f2ca86df216db8169bd0b6a6adb088cdad2557245f8a859b6f64878833492b60001b62006837565b81945050505050620100e6565b6200fee57ff3ea6d1fa4482e8d8b8b6279c1e33c350f044cea58afd75b6bfd308de4e22a4760001b62006837565b6200ff137f05724858d834000e6ca9a8fd9f7d47a82053ad20edc86e04511fba572e93479960001b62006837565b6200ff417f3819050f781a2769f76835ead78e6e5b281b2eb71aa0c012f67b4b5c2e76073d60001b62006837565b600060028111156200ff7c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8660028111156200ffb6577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b141562010054576200ffeb7f30d121430986bddfd9201a582308a8e3ff1f30da3df2f7ce26a1855582dbe6b860001b62006837565b620100197f5f13d2e6fb418cdd2476d04844bb1dcd70809dcccd7cf50765c8ae80635a39d060001b62006837565b620100477f253d0237604ffe85cd4de0adde965168ad2d985132eea99d953bddcdff09d64460001b62006837565b82945050505050620100e6565b620100827f1b992a7f69d59f7147f093099f9972040fc4ecc225937dd44bd32ff61e2964be60001b62006837565b620100b07f90d14d2386b35af2f4eb2c0bcecb8649f59c33deb0b7119a0077efb1a7fb714660001b62006837565b620100de7f7764b8b3398c7de7f2d0e7412fa1d9a40e26f4bc21ff49414bcbdb709b059e1460001b62006837565b829450505050505b919050565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6040518060a00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600081525090565b6040518060600160405280620101a462010142565b8152602001606081526020016000151581525090565b6040518060800160405280620101cf62010142565b8152602001620101de620100eb565b8152602001606081526020016000151581525090565b6040518060c001604052806201020962010142565b8152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600081526020016201024162010506565b8152602001606081525090565b6040518060a001604052806201026362010142565b8152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600081526020016201029b62010506565b81525090565b6040518060a00160405280620102b6620105ce565b8152602001620102c562010142565b8152602001620102d462010506565b815260200160608152602001600067ffffffffffffffff1681525090565b60405180606001604052806201030762010142565b815260200162010316620100eb565b8152602001606081525090565b604051806060016040528062010338620105ce565b81526020016201034762010142565b81526020016201035662010506565b81525090565b6040518060a001604052806201037162010142565b8152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081526020016000151581525090565b6040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b604051806101e0016040528060608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160001515815260200160608152602001600015158152602001600015158152602001606081526020016060815260200160001515815260200160006001811115620104cb577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81525090565b6040518060600160405280620104e662010142565b8152602001600067ffffffffffffffff1681526020016000151581525090565b604051806101600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016201061d62010632565b81526020016201062c62010632565b81525090565b60405180604001604052806000600281111562010678577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600067ffffffffffffffff1681525090565b6000620106a6620106a08462012c7b565b62012c52565b90508083825260208201905082856020860282011115620106c657600080fd5b60005b85811015620106fa5781620106df888262010955565b845260208401935060208301925050600181019050620106c9565b5050509392505050565b60006201071b620107158462012caa565b62012c52565b905080838252602082019050828560208602820111156201073b57600080fd5b60005b858110156201078a57813567ffffffffffffffff8111156201075f57600080fd5b8086016201076e898262010a4e565b855260208501945060208401935050506001810190506201073e565b5050509392505050565b6000620107ab620107a58462012caa565b62012c52565b90508083825260208201905082856020860282011115620107cb57600080fd5b60005b858110156201081a57815167ffffffffffffffff811115620107ef57600080fd5b808601620107fe898262010a7b565b85526020850194506020840193505050600181019050620107ce565b5050509392505050565b60006201083b620108358462012cd9565b62012c52565b905080838252602082019050828560208602820111156201085b57600080fd5b60005b85811015620108aa57815167ffffffffffffffff8111156201087f57600080fd5b8086016201088e898262010be2565b855260208501945060208401935050506001810190506201085e565b5050509392505050565b6000620108cb620108c58462012d08565b62012c52565b905082815260208101848484011115620108e457600080fd5b620108f184828562013198565b509392505050565b6000620109106201090a8462012d08565b62012c52565b9050828152602081018484840111156201092957600080fd5b62010936848285620131a7565b509392505050565b6000813590506201094f8162013454565b92915050565b600081519050620109668162013454565b92915050565b600082601f8301126201097e57600080fd5b8151620109908482602086016201068f565b91505092915050565b600082601f830112620109ab57600080fd5b8135620109bd84826020860162010704565b91505092915050565b600082601f830112620109d857600080fd5b8151620109ea84826020860162010794565b91505092915050565b600082601f83011262010a0557600080fd5b815162010a1784826020860162010824565b91505092915050565b60008151905062010a31816201346e565b92915050565b60008135905062010a488162013488565b92915050565b600082601f83011262010a6057600080fd5b813562010a72848260208601620108b4565b91505092915050565b600082601f83011262010a8d57600080fd5b815162010a9f848260208601620108f9565b91505092915050565b60008135905062010ab981620134a2565b92915050565b60008135905062010ad081620134bc565b92915050565b60008151905062010ae781620134d6565b92915050565b60008151905062010afe81620134e7565b92915050565b60008135905062010b1581620134f8565b92915050565b6000610280828403121562010b2f57600080fd5b62010b3b60c062012c52565b9050600062010b4d84828501620113bc565b60008301525060a062010b63848285016201147f565b60208301525060c062010b79848285016201147f565b60408301525060e062010b8f8482850162011451565b60608301525061010062010ba6848285016201103f565b60808301525061026082013567ffffffffffffffff81111562010bc857600080fd5b62010bd68482850162010999565b60a08301525092915050565b6000610320828403121562010bf657600080fd5b62010c036102e062012c52565b9050600082015167ffffffffffffffff81111562010c2057600080fd5b62010c2e8482850162010a7b565b600083015250602062010c448482850162010955565b602083015250604082015167ffffffffffffffff81111562010c6557600080fd5b62010c738482850162010a7b565b604083015250606062010c898482850162011496565b606083015250608062010c9f8482850162011496565b60808301525060a062010cb58482850162011496565b60a08301525060c062010ccb8482850162011496565b60c08301525060e062010ce18482850162011496565b60e08301525061010062010cf88482850162011468565b6101008301525061012062010d108482850162011496565b6101208301525061014062010d288482850162011496565b6101408301525061016082015167ffffffffffffffff81111562010d4b57600080fd5b62010d598482850162010a7b565b6101608301525061018062010d718482850162011496565b610180830152506101a062010d898482850162011468565b6101a0830152506101c062010da18482850162010a20565b6101c0830152506101e062010db98482850162010aed565b6101e08301525061020062010dd18482850162011496565b6102008301525061022082015167ffffffffffffffff81111562010df457600080fd5b62010e02848285016201096c565b6102208301525061024082015167ffffffffffffffff81111562010e2557600080fd5b62010e33848285016201096c565b6102408301525061026082015167ffffffffffffffff81111562010e5657600080fd5b62010e648482850162010a7b565b6102608301525061028062010e7c8482850162010ad6565b610280830152506102a062010e948482850162010a20565b6102a0830152506102c062010eac8482850162010eb9565b6102c08301525092915050565b60006060828403121562010ecc57600080fd5b62010ed8606062012c52565b9050600062010eea8482850162011496565b600083015250602062010f008482850162011496565b602083015250604062010f168482850162011496565b60408301525092915050565b60006102c0828403121562010f3657600080fd5b62010f42606062012c52565b9050600062010f54848285016201133d565b60008301525060c062010f6a84828501620113bc565b60208301525061016062010f81848285016201103f565b60408301525092915050565b6000610300828403121562010fa157600080fd5b62010fad60a062012c52565b9050600062010fbf848285016201133d565b60008301525060c062010fd584828501620113bc565b60208301525061016062010fec848285016201103f565b6040830152506102c082013567ffffffffffffffff8111156201100e57600080fd5b6201101c8482850162010999565b6060830152506102e062011033848285016201147f565b60808301525092915050565b600061016082840312156201105357600080fd5b6201106061016062012c52565b9050600062011072848285016201147f565b600083015250602062011088848285016201147f565b60208301525060406201109e848285016201147f565b6040830152506060620110b4848285016201147f565b6060830152506080620110ca848285016201147f565b60808301525060a0620110e0848285016201147f565b60a08301525060c0620110f6848285016201147f565b60c08301525060e06201110c848285016201147f565b60e08301525061010062011123848285016201147f565b610100830152506101206201113b848285016201147f565b6101208301525061014062011153848285016201147f565b6101408301525092915050565b600061016082840312156201117457600080fd5b6201118161016062012c52565b90506000620111938482850162011496565b6000830152506020620111a98482850162011496565b6020830152506040620111bf8482850162011496565b6040830152506060620111d58482850162011496565b6060830152506080620111eb8482850162011496565b60808301525060a0620112018482850162011496565b60a08301525060c0620112178482850162011496565b60c08301525060e06201122d8482850162011496565b60e083015250610100620112448482850162011496565b610100830152506101206201125c8482850162011496565b61012083015250610140620112748482850162011496565b6101408301525092915050565b6000606082840312156201129457600080fd5b620112a0606062012c52565b90506000620112b28482850162011496565b6000830152506020620112c88482850162011496565b6020830152506040620112de8482850162011496565b60408301525092915050565b600060408284031215620112fd57600080fd5b62011309604062012c52565b905060006201131b8482850162010b04565b600083015250602062011331848285016201147f565b60208301525092915050565b600060c082840312156201135057600080fd5b6201135c608062012c52565b905060006201136e848285016201093e565b600083015250602062011384848285016201093e565b60208301525060406201139a84828501620112ea565b6040830152506080620113b084828501620112ea565b60608301525092915050565b600060a08284031215620113cf57600080fd5b620113db60a062012c52565b90506000620113ed848285016201147f565b600083015250602062011403848285016201147f565b602083015250604062011419848285016201147f565b60408301525060606201142f8482850162011451565b6060830152506080620114458482850162011451565b60808301525092915050565b600081359050620114628162013509565b92915050565b600081519050620114798162013509565b92915050565b600081359050620114908162013523565b92915050565b600081519050620114a78162013523565b92915050565b600060208284031215620114c057600080fd5b6000620114d0848285016201093e565b91505092915050565b60008060c08385031215620114ed57600080fd5b6000620114fd858286016201093e565b92505060206201151085828601620113bc565b9150509250929050565b6000602082840312156201152d57600080fd5b600082015167ffffffffffffffff8111156201154857600080fd5b6201155684828501620109c6565b91505092915050565b6000806000606084860312156201157557600080fd5b600084015167ffffffffffffffff8111156201159057600080fd5b6201159e86828701620109c6565b9350506020620115b18682870162011496565b9250506040620115c48682870162010a20565b9150509250925092565b60008060408385031215620115e257600080fd5b600083015167ffffffffffffffff811115620115fd57600080fd5b6201160b85828601620109f3565b92505060206201161e8582860162010a20565b9150509250929050565b6000602082840312156201163b57600080fd5b60006201164b8482850162010a37565b91505092915050565b600080604083850312156201166857600080fd5b6000620116788582860162010aa8565b92505060206201168b8582860162010abf565b9150509250929050565b600060208284031215620116a857600080fd5b600082013567ffffffffffffffff811115620116c357600080fd5b620116d18482850162010b1b565b91505092915050565b60006102c08284031215620116ee57600080fd5b6000620116fe8482850162010f22565b91505092915050565b6000602082840312156201171a57600080fd5b600082013567ffffffffffffffff8111156201173557600080fd5b620117438482850162010f8d565b91505092915050565b600061016082840312156201176057600080fd5b6000620117708482850162011160565b91505092915050565b60008060006101a084860312156201179057600080fd5b6000620117a0868287016201103f565b935050610160620117b48682870162011451565b925050610180620117c8868287016201147f565b9150509250925092565b6000806000806101c08587031215620117ea57600080fd5b6000620117fa878288016201103f565b9450506101606201180e8782880162011451565b93505061018062011822878288016201147f565b9250506101a062011836878288016201147f565b91505092959194509250565b60008061018083850312156201185757600080fd5b600062011867858286016201103f565b9250506101606201187b858286016201147f565b9150509250929050565b60008060006101a084860312156201189c57600080fd5b6000620118ac868287016201103f565b935050610160620118c0868287016201147f565b925050610180620118d48682870162011451565b9150509250925092565b6000806000806101c08587031215620118f657600080fd5b600062011906878288016201103f565b9450506101606201191a878288016201147f565b9350506101806201192e878288016201147f565b9250506101a0620119428782880162011451565b91505092959194509250565b6000606082840312156201196157600080fd5b6000620119718482850162011281565b91505092915050565b600060c082840312156201198d57600080fd5b60006201199d848285016201133d565b91505092915050565b6000620119b4838362011a20565b60208301905092915050565b6000620119ce838362011c93565b905092915050565b6000620119e4838362011cf6565b60208301905092915050565b6000620119fe838362011ea3565b905092915050565b600062011a14838362012834565b60608301905092915050565b62011a2b8162013064565b82525050565b62011a3c8162013064565b82525050565b600062011a4f8262012d8e565b62011a5b818562012e1c565b935062011a688362012d3e565b8060005b8381101562011a9f57815162011a838882620119a6565b975062011a908362012ddb565b92505060018101905062011a6c565b5085935050505092915050565b600062011ab98262012d99565b62011ac5818562012e2d565b93508360208202850162011ad98562012d4e565b8060005b8581101562011b1b578484038952815162011af98582620119c0565b945062011b068362012de8565b925060208a0199505060018101905062011add565b50829750879550505050505092915050565b600062011b3a8262012da4565b62011b46818562012e60565b935062011b538362012d5e565b8060005b8381101562011b8a57815162011b6e8882620119d6565b975062011b7b8362012df5565b92505060018101905062011b57565b5085935050505092915050565b600062011ba48262012daf565b62011bb0818562012e3e565b93508360208202850162011bc48562012d6e565b8060005b8581101562011c06578484038952815162011be48582620119f0565b945062011bf18362012e02565b925060208a0199505060018101905062011bc8565b50829750879550505050505092915050565b600062011c258262012dba565b62011c31818562012e4f565b935062011c3e8362012d7e565b8060005b8381101562011c7557815162011c59888262011a06565b975062011c668362012e0f565b92505060018101905062011c42565b5085935050505092915050565b62011c8d8162013078565b82525050565b600062011ca08262012dc5565b62011cac818562012e71565b935062011cbe818560208601620131a7565b62011cc9816201331d565b840191505092915050565b62011cdf8162013148565b82525050565b62011cf0816201315c565b82525050565b62011d018162013170565b82525050565b62011d128162013184565b82525050565b600062011d258262012dd0565b62011d31818562012e82565b935062011d43818560208601620131a7565b62011d4e816201331d565b840191505092915050565b600062011d68600f8362012e82565b915062011d75826201332e565b602082019050919050565b600062011d8f60128362012e82565b915062011d9c8262013357565b602082019050919050565b600062011db6600f8362012e82565b915062011dc38262013380565b602082019050919050565b600062011ddd602e8362012e82565b915062011dea82620133a9565b604082019050919050565b600060e08301600083015162011e0f600086018262012750565b50602083015184820360a086015262011e29828262011b97565b915050604083015162011e4060c086018262011c82565b508091505092915050565b60006101208301600083015162011e66600086018262012750565b50602083015162011e7b60a08601826201254a565b50604083015184820361010086015262011e96828262011b97565b9150508091505092915050565b600061032083016000830151848203600086015262011ec3828262011c93565b915050602083015162011eda602086018262011a20565b506040830151848203604086015262011ef4828262011c93565b915050606083015162011f0b60608601826201289e565b50608083015162011f2060808601826201289e565b5060a083015162011f3560a08601826201289e565b5060c083015162011f4a60c08601826201289e565b5060e083015162011f5f60e08601826201289e565b5061010083015162011f766101008601826201287c565b5061012083015162011f8d6101208601826201289e565b5061014083015162011fa46101408601826201289e565b5061016083015184820361016086015262011fc0828262011c93565b91505061018083015162011fd96101808601826201289e565b506101a083015162011ff06101a08601826201287c565b506101c0830151620120076101c086018262011c82565b506101e08301516201201e6101e086018262011cf6565b50610200830151620120356102008601826201289e565b5061022083015184820361022086015262012051828262011a42565b9150506102408301518482036102408601526201206f828262011a42565b9150506102608301518482036102608601526201208d828262011c93565b915050610280830151620120a661028086018262011ce5565b506102a0830151620120bd6102a086018262011c82565b506102c0830151620120d46102c08601826201231b565b508091505092915050565b6000610320830160008301518482036000860152620120ff828262011c93565b915050602083015162012116602086018262011a20565b506040830151848203604086015262012130828262011c93565b91505060608301516201214760608601826201289e565b5060808301516201215c60808601826201289e565b5060a08301516201217160a08601826201289e565b5060c08301516201218660c08601826201289e565b5060e08301516201219b60e08601826201289e565b50610100830151620121b26101008601826201287c565b50610120830151620121c96101208601826201289e565b50610140830151620121e06101408601826201289e565b50610160830151848203610160860152620121fc828262011c93565b915050610180830151620122156101808601826201289e565b506101a08301516201222c6101a08601826201287c565b506101c0830151620122436101c086018262011c82565b506101e08301516201225a6101e086018262011cf6565b50610200830151620122716102008601826201289e565b506102208301518482036102208601526201228d828262011a42565b915050610240830151848203610240860152620122ab828262011a42565b915050610260830151848203610260860152620122c9828262011c93565b915050610280830151620122e261028086018262011ce5565b506102a0830151620122f96102a086018262011c82565b506102c0830151620123106102c08601826201231b565b508091505092915050565b6060820160008201516201233360008501826201289e565b5060208201516201234860208501826201289e565b5060408201516201235d60408501826201289e565b50505050565b6000610140830160008301516201237e600086018262012750565b5060208301516201239360a08601826201254a565b506040830151848203610100860152620123ae828262011b97565b9150506060830151620123c661012086018262011c82565b508091505092915050565b600061012083016000830151620123ec600086018262012750565b5060208301516201240160a08601826201289e565b5060408301516201241660c08601826201289e565b50606083015184820360e086015262012430828262011b97565b91505060808301516201244861010086018262011c82565b508091505092915050565b610160820160008201516201246c60008501826201289e565b5060208201516201248160208501826201289e565b5060408201516201249660408501826201289e565b506060820151620124ab60608501826201289e565b506080820151620124c060808501826201289e565b5060a0820151620124d560a08501826201289e565b5060c0820151620124ea60c08501826201289e565b5060e0820151620124ff60e08501826201289e565b50610100820151620125166101008501826201289e565b506101208201516201252d6101208501826201289e565b50610140820151620125446101408501826201289e565b50505050565b60608201600082015162012562600085018262011a20565b50602082015162012577602085018262011a20565b5060408201516201258c60408501826201289e565b50505050565b606082016000820151620125aa600085018262011a20565b506020820151620125bf602085018262011a20565b506040820151620125d460408501826201289e565b50505050565b60006101e0830160008301518482036000860152620125fa828262011c93565b91505060208301516201261160208601826201289e565b5060408301516201262660408601826201289e565b5060608301516201263b60608601826201289e565b5060808301516201265060808601826201287c565b5060a08301516201266560a08601826201289e565b5060c08301516201267a60c08601826201289e565b5060e08301516201268f60e086018262011c82565b50610100830151848203610100860152620126ab828262011c93565b915050610120830151620126c461012086018262011c82565b50610140830151620126db61014086018262011c82565b50610160830151848203610160860152620126f7828262011c93565b91505061018083015184820361018086015262012715828262011c18565b9150506101a08301516201272e6101a086018262011c82565b506101c0830151620127456101c086018262011cf6565b508091505092915050565b60a0820160008201516201276860008501826201289e565b5060208201516201277d60208501826201289e565b5060408201516201279260408501826201289e565b506060820151620127a760608501826201287c565b506080820151620127bc60808501826201287c565b50505050565b60a082016000820151620127da60008501826201289e565b506020820151620127ef60208501826201289e565b5060408201516201280460408501826201289e565b5060608201516201281960608501826201287c565b5060808201516201282e60808501826201287c565b50505050565b6060820160008201516201284c600085018262011a20565b5060208201516201286160208501826201289e565b5060408201516201287660408501826201289e565b50505050565b62012887816201312a565b82525050565b62012898816201312a565b82525050565b620128a98162013134565b82525050565b620128ba8162013134565b82525050565b6000602082019050620128d7600083018462011a31565b92915050565b6000604082019050620128f4600083018562011a31565b818103602083015262012908818462011aac565b90509392505050565b600060608201905081810360008301526201292d818662011aac565b90506201293e602083018562011a31565b818103604083015262012952818462011b2d565b9050949350505050565b60006101a082019050818103600083015262012979818662011aac565b90506201298a602083018562012453565b6201299a6101808301846201288d565b949350505050565b600060e082019050620129b9600083018a62011cd4565b620129c860208301896201288d565b620129d7604083018862011a31565b620129e6606083018762011d07565b620129f56080830186620128af565b62012a0460a083018562011d07565b62012a1360c0830184620128af565b98975050505050505050565b6000604082019050818103600083015262012a3a8162011d59565b9050818103602083015262012a50818462011d18565b905092915050565b6000604082019050818103600083015262012a738162011d59565b9050818103602083015262012a888162011d80565b9050919050565b6000604082019050818103600083015262012aaa8162011da7565b9050818103602083015262012ac0818462011d18565b905092915050565b6000602082019050818103600083015262012ae38162011dce565b9050919050565b6000602082019050818103600083015262012b06818462011df5565b905092915050565b6000604082019050818103600083015262012b2a818562011e4b565b9050818103602083015262012b40818462011d18565b90509392505050565b6000602082019050818103600083015262012b658184620120df565b905092915050565b6000602082019050818103600083015262012b89818462012363565b905092915050565b6000602082019050818103600083015262012bad8184620123d1565b905092915050565b600060608201905062012bcc600083018462012592565b92915050565b60006101a082019050818103600083015262012bef8186620125da565b905062012c00602083018562012453565b62012c106101808301846201288d565b949350505050565b600060a08201905062012c2f6000830184620127c2565b92915050565b600060208201905062012c4c6000830184620128af565b92915050565b600062012c5e62012c71565b905062012c6c8282620131dd565b919050565b6000604051905090565b600067ffffffffffffffff82111562012c995762012c98620132ee565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562012cc85762012cc7620132ee565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562012cf75762012cf6620132ee565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562012d265762012d25620132ee565b5b62012d31826201331d565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600062012ea0826201312a565b915062012ead836201312a565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111562012ee55762012ee462013261565b5b828201905092915050565b600062012efd8262013134565b915062012f0a8362013134565b92508267ffffffffffffffff0382111562012f2a5762012f2962013261565b5b828201905092915050565b600062012f42826201312a565b915062012f4f836201312a565b92508262012f625762012f6162013290565b5b828204905092915050565b600062012f7a8262013134565b915062012f878362013134565b92508262012f9a5762012f9962013290565b5b828204905092915050565b600062012fb28262013134565b915062012fbf8362013134565b92508167ffffffffffffffff048311821515161562012fe35762012fe262013261565b5b828202905092915050565b600062012ffb826201312a565b915062013008836201312a565b9250828210156201301e576201301d62013261565b5b828203905092915050565b6000620130368262013134565b9150620130438362013134565b92508282101562013059576201305862013261565b5b828203905092915050565b600062013071826201310a565b9050919050565b60008115159050919050565b6000819050919050565b60006201309b8262013064565b9050919050565b6000620130af8262013064565b9050919050565b6000819050620130c682620133f8565b919050565b6000819050620130db826201340f565b919050565b6000819050620130f08262013426565b919050565b600081905062013105826201343d565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b60006201315582620130b6565b9050919050565b60006201316982620130cb565b9050919050565b60006201317d82620130e0565b9050919050565b60006201319182620130f5565b9050919050565b82818337600083830152505050565b60005b83811015620131c7578082015181840152602081019050620131aa565b83811115620131d7576000848401525b50505050565b620131e8826201331d565b810181811067ffffffffffffffff821117156201320a5762013209620132ee565b5b80604052505050565b600062013220826201312a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141562013256576201325562013261565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f4d616e6167655573657253706163650000000000000000000000000000000000600082015250565b7f496e73756666696369656e742066756e64730000000000000000000000000000600082015250565b7f44656c6574655573657253706163650000000000000000000000000000000000600082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b600b81106201340c576201340b620132bf565b5b50565b60038110620134235762013422620132bf565b5b50565b600281106201343a5762013439620132bf565b5b50565b60038110620134515762013450620132bf565b5b50565b6201345f8162013064565b81146201346b57600080fd5b50565b620134798162013078565b81146201348557600080fd5b50565b620134938162013084565b81146201349f57600080fd5b50565b620134ad816201308e565b8114620134b957600080fd5b50565b620134c781620130a2565b8114620134d357600080fd5b50565b60038110620134e457600080fd5b50565b60028110620134f557600080fd5b50565b600381106201350657600080fd5b50565b62013514816201312a565b81146201352057600080fd5b50565b6201352e8162013134565b81146201353a57600080fd5b5056fe696e76616c696420666972737420757365727370616365206f7065726174696f6ea264697066735822122045562c06141177521ad6bacc81015d18744de8fd0f2597069c129ccbe51c64ac64736f6c63430008040033",
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

// C0x936f0bd2 is a free data retrieval call binding the contract method 0xb5bc8ee2.
//
// Solidity: function c_0x936f0bd2(bytes32 c__0x936f0bd2) pure returns()
func (_Store *StoreCaller) C0x936f0bd2(opts *bind.CallOpts, c__0x936f0bd2 [32]byte) error {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "c_0x936f0bd2", c__0x936f0bd2)

	if err != nil {
		return err
	}

	return err

}

// C0x936f0bd2 is a free data retrieval call binding the contract method 0xb5bc8ee2.
//
// Solidity: function c_0x936f0bd2(bytes32 c__0x936f0bd2) pure returns()
func (_Store *StoreSession) C0x936f0bd2(c__0x936f0bd2 [32]byte) error {
	return _Store.Contract.C0x936f0bd2(&_Store.CallOpts, c__0x936f0bd2)
}

// C0x936f0bd2 is a free data retrieval call binding the contract method 0xb5bc8ee2.
//
// Solidity: function c_0x936f0bd2(bytes32 c__0x936f0bd2) pure returns()
func (_Store *StoreCallerSession) C0x936f0bd2(c__0x936f0bd2 [32]byte) error {
	return _Store.Contract.C0x936f0bd2(&_Store.CallOpts, c__0x936f0bd2)
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
// Solidity: function AddUserSpace(((uint64,uint64,uint64,uint256,uint256),uint64,uint64,uint256,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64),bytes[]) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreTransactor) AddUserSpace(opts *bind.TransactOpts, params SpaceAddParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "AddUserSpace", params)
}

// AddUserSpace is a paid mutator transaction binding the contract method 0x6a4fe520.
//
// Solidity: function AddUserSpace(((uint64,uint64,uint64,uint256,uint256),uint64,uint64,uint256,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64),bytes[]) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreSession) AddUserSpace(params SpaceAddParams) (*types.Transaction, error) {
	return _Store.Contract.AddUserSpace(&_Store.TransactOpts, params)
}

// AddUserSpace is a paid mutator transaction binding the contract method 0x6a4fe520.
//
// Solidity: function AddUserSpace(((uint64,uint64,uint64,uint256,uint256),uint64,uint64,uint256,(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64),bytes[]) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
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
// Solidity: function getUserspaceChange((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(address,address,uint64),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[]), string)
func (_Store *StoreTransactor) GetUserspaceChange(opts *bind.TransactOpts, params UserSpaceParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "getUserspaceChange", params)
}

// GetUserspaceChange is a paid mutator transaction binding the contract method 0x722b25b9.
//
// Solidity: function getUserspaceChange((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(address,address,uint64),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[]), string)
func (_Store *StoreSession) GetUserspaceChange(params UserSpaceParams) (*types.Transaction, error) {
	return _Store.Contract.GetUserspaceChange(&_Store.TransactOpts, params)
}

// GetUserspaceChange is a paid mutator transaction binding the contract method 0x722b25b9.
//
// Solidity: function getUserspaceChange((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(address,address,uint64),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[]), string)
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
// Solidity: function processForUserSpaceOneAddOneRevoke(((address,address,(uint8,uint64),(uint8,uint64)),(uint64,uint64,uint64,uint256,uint256),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64),bytes[],uint64) params) payable returns(((uint64,uint64,uint64,uint256,uint256),uint64,uint64,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreTransactor) ProcessForUserSpaceOneAddOneRevoke(opts *bind.TransactOpts, params SpaceProcessRevokeParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "processForUserSpaceOneAddOneRevoke", params)
}

// ProcessForUserSpaceOneAddOneRevoke is a paid mutator transaction binding the contract method 0xf1feacbf.
//
// Solidity: function processForUserSpaceOneAddOneRevoke(((address,address,(uint8,uint64),(uint8,uint64)),(uint64,uint64,uint64,uint256,uint256),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64),bytes[],uint64) params) payable returns(((uint64,uint64,uint64,uint256,uint256),uint64,uint64,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreSession) ProcessForUserSpaceOneAddOneRevoke(params SpaceProcessRevokeParams) (*types.Transaction, error) {
	return _Store.Contract.ProcessForUserSpaceOneAddOneRevoke(&_Store.TransactOpts, params)
}

// ProcessForUserSpaceOneAddOneRevoke is a paid mutator transaction binding the contract method 0xf1feacbf.
//
// Solidity: function processForUserSpaceOneAddOneRevoke(((address,address,(uint8,uint64),(uint8,uint64)),(uint64,uint64,uint64,uint256,uint256),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64),bytes[],uint64) params) payable returns(((uint64,uint64,uint64,uint256,uint256),uint64,uint64,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreTransactorSession) ProcessForUserSpaceOneAddOneRevoke(params SpaceProcessRevokeParams) (*types.Transaction, error) {
	return _Store.Contract.ProcessForUserSpaceOneAddOneRevoke(&_Store.TransactOpts, params)
}

// ProcessForUserSpaceOperations is a paid mutator transaction binding the contract method 0x6c799c13.
//
// Solidity: function processForUserSpaceOperations(((address,address,(uint8,uint64),(uint8,uint64)),(uint64,uint64,uint64,uint256,uint256),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64)) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(address,address,uint64),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreTransactor) ProcessForUserSpaceOperations(opts *bind.TransactOpts, params SpaceProcessParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "processForUserSpaceOperations", params)
}

// ProcessForUserSpaceOperations is a paid mutator transaction binding the contract method 0x6c799c13.
//
// Solidity: function processForUserSpaceOperations(((address,address,(uint8,uint64),(uint8,uint64)),(uint64,uint64,uint64,uint256,uint256),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64)) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(address,address,uint64),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
func (_Store *StoreSession) ProcessForUserSpaceOperations(params SpaceProcessParams) (*types.Transaction, error) {
	return _Store.Contract.ProcessForUserSpaceOperations(&_Store.TransactOpts, params)
}

// ProcessForUserSpaceOperations is a paid mutator transaction binding the contract method 0x6c799c13.
//
// Solidity: function processForUserSpaceOperations(((address,address,(uint8,uint64),(uint8,uint64)),(uint64,uint64,uint64,uint256,uint256),(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64)) params) payable returns(((uint64,uint64,uint64,uint256,uint256),(address,address,uint64),(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[],bool))
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
