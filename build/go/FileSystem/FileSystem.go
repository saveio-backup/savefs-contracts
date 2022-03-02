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

// FileReNewInfo is an auto generated low-level Go binding around an user-defined struct.
type FileReNewInfo struct {
	FileHash   []byte
	FromAddr   common.Address
	ReNewTimes uint64
}

// FileSystemOwnerChange is an auto generated low-level Go binding around an user-defined struct.
type FileSystemOwnerChange struct {
	FileHash []byte
	CurOwner common.Address
	NewOwner common.Address
}

// FileSystemPriChange is an auto generated low-level Go binding around an user-defined struct.
type FileSystemPriChange struct {
	FileHash  []byte
	Privilege uint64
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

// StorageFee is an auto generated low-level Go binding around an user-defined struct.
type StorageFee struct {
	TxnFee        uint64
	SpaceFee      uint64
	ValidationFee uint64
}

// UploadOption is an auto generated low-level Go binding around an user-defined struct.
type UploadOption struct {
	FileDesc        []byte
	FileSize        uint64
	ProveInterval   uint64
	ProveLevel      uint64
	ExpiredHeight   *big.Int
	Privilege       uint64
	CopyNum         uint64
	Encrypt         bool
	EncryptPassword []byte
	RegisterDNS     bool
	BindDNS         bool
	DnsURL          []byte
	WhiteList       []WhiteList
	Share           bool
	StorageType     uint8
}

// WhiteList is an auto generated low-level Go binding around an user-defined struct.
type WhiteList struct {
	Addr         common.Address
	BaseHeight   uint64
	ExpireHeight uint64
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DifferenceFileOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"FileNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProfit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"OpError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceWrongExpiredHeight\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"AddFileToUnSettleList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"currentHeight\",\"type\":\"uint256\"}],\"name\":\"CalcDepositFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"proveTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"CalcFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"beginHeight\",\"type\":\"uint256\"}],\"name\":\"CalcProveTimesByUploadInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"CurOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NewOwner\",\"type\":\"address\"}],\"internalType\":\"structFileSystem.OwnerChange\",\"name\":\"ownerChange\",\"type\":\"tuple\"}],\"name\":\"ChangeFileOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"privilege\",\"type\":\"uint64\"}],\"internalType\":\"structFileSystem.PriChange\",\"name\":\"priChange\",\"type\":\"tuple\"}],\"name\":\"ChangeFilePrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"rmInfo\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rmList\",\"type\":\"bool\"}],\"name\":\"CleanupForDeleteFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DelFileFromCandidateList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DelFileFromList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DelFileFromPrimaryList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DelFileFromUnSettledList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"enumStorageType[]\",\"name\":\"storageType\",\"type\":\"uint8[]\"}],\"name\":\"DeleteExpiredFilesFromList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteFileInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"}],\"name\":\"DeleteFiles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteUnsettledFiles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FromAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ReNewTimes\",\"type\":\"uint64\"}],\"internalType\":\"structFileReNewInfo\",\"name\":\"fileReNewInfo\",\"type\":\"tuple\"}],\"name\":\"FileReNew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"GetFileInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"}],\"name\":\"GetFileInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetFileList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnProveCandidateFiles\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnProvePrimaryFiles\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnSettledFileList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"}],\"name\":\"GetUploadStorageFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"StoreFile\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"f\",\"type\":\"tuple\"}],\"name\":\"UpdateFileInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"newExpireHeight\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"UpdateFileInfoForRenew\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"list\",\"type\":\"bytes[]\"}],\"name\":\"UpdateFileList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"newExpireHeight\",\"type\":\"uint256\"}],\"name\":\"UpdateFilesForRenew\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"}],\"name\":\"calcSingleValidFeeForFile\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"calcStorageFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"calcStorageFeeForOneNode\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"currentHeight\",\"type\":\"uint256\"}],\"name\":\"calcUploadFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"proveTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"}],\"name\":\"calcValidFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"proveTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"}],\"name\":\"calcValidFeeForOneNode\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"files\",\"type\":\"tuple[]\"}],\"name\":\"deleteFilesInner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"deleteProveDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractNode\",\"name\":\"_node\",\"type\":\"address\"},{\"internalType\":\"contractSpace\",\"name\":\"_space\",\"type\":\"address\"},{\"internalType\":\"contractSector\",\"name\":\"_sector\",\"type\":\"address\"},{\"internalType\":\"contractProve\",\"name\":\"_prove\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260048054600160a01b600160e01b031916600560a01b17908190556200003f906001600160401b03600160a01b909104166201518062000082565b600580546001600160401b03929092166001600160801b0319909216919091176a0f424000000000000000001790553480156200007b57600080fd5b50620000bc565b6001600160401b039182169116600082620000a157620000a1620000a6565b500490565b634e487b7160e01b600052601260045260246000fd5b6182ac80620000cc6000396000f3fe6080604052600436106102855760003560e01c80638d41f9f811610153578063cc76e80d116100cb578063defce5d41161007f578063e4bca97311610064578063e4bca97314610753578063ebaf025814610773578063f54cd2951461079357600080fd5b8063defce5d414610706578063df49410a1461073357600080fd5b8063d49ce874116100b0578063d49ce874146106a6578063d70e6272146106c6578063dc1ec30b146106e657600080fd5b8063cc76e80d14610666578063ce9045541461068657600080fd5b8063b64ab1ef11610122578063c43007e511610107578063c43007e5146105f9578063c6e8392a14610619578063ca6142cb1461064657600080fd5b8063b64ab1ef146105b9578063b6af10fb146105d957600080fd5b80638d41f9f81461053957806395b0b54b146105595780639a2b5e63146105795780639cd103e91461059957600080fd5b80633d4dcb9811610201578063655c6694116101b557806370b1d8891161019a57806370b1d889146104cc578063740f6913146104ec578063793cab8e1461051957600080fd5b8063655c66941461048c578063681850d7146104ac57600080fd5b806341bc86cb116101e657806341bc86cb14610439578063432152ce1461045957806357d943991461047957600080fd5b80633d4dcb98146103eb5780633f2cc9a01461040b57600080fd5b80631bc558801161025857806328a8565c1161023d57806328a8565c1461036f57806334fddaac1461039c5780633ad525a9146103bc57600080fd5b80631bc558801461032f578063207e33be1461034f57600080fd5b80630be097021461028a5780631459457a146102ac578063170bbdf5146102cc578063178e4dc014610302575b600080fd5b34801561029657600080fd5b506102aa6102a5366004616bce565b6107a6565b005b3480156102b857600080fd5b506102aa6102c7366004616e29565b610ab6565b3480156102d857600080fd5b506102ec6102e736600461710d565b610bb5565b6040516102f99190617eec565b60405180910390f35b34801561030e57600080fd5b5061032261031d366004617213565b610be2565b6040516102f99190617ea8565b34801561033b57600080fd5b506102aa61034a366004616c3d565b610c38565b34801561035b57600080fd5b506102aa61036a366004616f7d565b610d31565b34801561037b57600080fd5b5061038f61038a366004616bce565b610d9b565b6040516102f99190617c22565b3480156103a857600080fd5b506102aa6103b7366004616c3d565b610e8a565b3480156103c857600080fd5b506103dc6103d7366004616cb8565b610ec6565b6040516102f993929190617c33565b3480156103f757600080fd5b506102aa610406366004616df5565b6111ca565b34801561041757600080fd5b5061042b610426366004616d32565b611367565b6040516102f9929190617c71565b34801561044557600080fd5b506103226104543660046171df565b61142d565b34801561046557600080fd5b506102aa610474366004616c84565b611506565b6102aa610487366004616f2b565b611633565b34801561049857600080fd5b506102aa6104a7366004616c3d565b612109565b3480156104b857600080fd5b506102aa6104c7366004616e9e565b6121e7565b3480156104d857600080fd5b506102ec6104e736600461710d565b612526565b3480156104f857600080fd5b5061050c610507366004617038565b61253e565b6040516102f99190617c91565b34801561052557600080fd5b506102ec6105343660046170c7565b6126b2565b34801561054557600080fd5b5061038f610554366004616bce565b6126c8565b34801561056557600080fd5b506102aa610574366004616df5565b6129d0565b34801561058557600080fd5b5061038f610594366004616bce565b612af3565b3480156105a557600080fd5b506102aa6105b4366004616ed2565b612bd7565b3480156105c557600080fd5b506102aa6105d4366004616c3d565b612ccc565b3480156105e557600080fd5b506102aa6105f4366004616fb1565b612daa565b34801561060557600080fd5b506102aa610614366004616d8d565b612dd7565b34801561062557600080fd5b50610639610634366004616c84565b61376b565b6040516102f99190617c60565b34801561065257600080fd5b506102ec610661366004617095565b614283565b34801561067257600080fd5b50610322610681366004617172565b6142ac565b34801561069257600080fd5b506102aa6106a1366004616df5565b614322565b3480156106b257600080fd5b506103226106c1366004617213565b6143d5565b3480156106d257600080fd5b506102aa6106e1366004616bec565b6144be565b3480156106f257600080fd5b5061038f610701366004616bce565b6144e7565b34801561071257600080fd5b50610726610721366004616df5565b6147e5565b6040516102f99190617e56565b34801561073f57600080fd5b506102ec61074e3660046170c7565b614d33565b34801561075f57600080fd5b506102ec61076e36600461724b565b614d5e565b34801561077f57600080fd5b506102aa61078e366004616c3d565b614d8a565b6102aa6107a1366004616e9e565b614e68565b60006107b182610d9b565b60408051600280825260608201835292935060009290916020830190803683370190505090506000816000815181106107fa57634e487b7160e01b600052603260045260246000fd5b6020026020010190600181111561082157634e487b7160e01b600052602160045260246000fd5b9081600181111561084257634e487b7160e01b600052602160045260246000fd5b8152505060018160018151811061086957634e487b7160e01b600052603260045260246000fd5b6020026020010190600181111561089057634e487b7160e01b600052602160045260246000fd5b908160018111156108b157634e487b7160e01b600052602160045260246000fd5b90525060606000806108c4858786610ec6565b9194509250905060005b8351811015610aad576001600160a01b0387166000908152600b6020908152604080832080548251818502810185019093528083529192909190849084015b828210156109b957838290600052602060002001805461092c906180f5565b80601f0160208091040260200160405190810160405280929190818152602001828054610958906180f5565b80156109a55780601f1061097a576101008083540402835291602001916109a5565b820191906000526020600020905b81548152906001019060200180831161098857829003601f168201915b50505050508152602001906001019061090d565b50505050905060005b8151811015610a6e578583815181106109eb57634e487b7160e01b600052603260045260246000fd5b602002602001015180519060200120828281518110610a1a57634e487b7160e01b600052603260045260246000fd5b6020026020010151805190602001201415610a5c57818181518110610a4f57634e487b7160e01b600052603260045260246000fd5b6020026020010160608152505b80610a6681618148565b9150506109c2565b506001600160a01b0388166000908152600b602090815260409091208251610a98928401906159af565b50508080610aa590618148565b9150506108ce565b50505050505050565b600054610100900460ff16610ad15760005460ff1615610ad5565b303b155b610afa5760405162461bcd60e51b8152600401610af190617df6565b60405180910390fd5b600054610100900460ff16158015610b1c576000805461ffff19166101011790555b600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff16620100006001600160a01b038981169190910291909117909155600180546001600160a01b031990811688841617909155600280548216878416179055600380548216868416179055600480549091169184169190911790558015610bad576000805461ff00191690555b505050505050565b6000610bc28585846126b2565b610bcd846001617f89565b610bd79190617fdd565b90505b949350505050565b6040805160608101825260008082526020820181905291810182905290610c098584614d5e565b90506000610c2c85838860c001518960200151888b608001516106819190618019565b925050505b9392505050565b60005b6001600160a01b0383166000908152600a6020526040902054811015610d1a578180519060200120600a6000856001600160a01b03166001600160a01b031681526020019081526020016000208281548110610ca757634e487b7160e01b600052603260045260246000fd5b90600052602060002001604051610cbe9190617bed565b60405180910390201415610d1f576001600160a01b0383166000908152600a60205260409020805482908110610d0457634e487b7160e01b600052603260045260246000fd5b906000526020600020016000610d1a9190615a0c565b505050565b80610d2981618148565b915050610c3b565b6000610d4082600001516147e5565b905081602001516001600160a01b031681602001516001600160a01b031614610d7b5760405162461bcd60e51b8152600401610af190617dd6565b60408201516001600160a01b03166020820152610d97816121e7565b5050565b6001600160a01b0381166000908152600b60209081526040808320805482518185028101850190935280835260609492939192909184015b82821015610e7f578382906000526020600020018054610df2906180f5565b80601f0160208091040260200160405190810160405280929190818152602001828054610e1e906180f5565b8015610e6b5780601f10610e4057610100808354040283529160200191610e6b565b820191906000526020600020905b815481529060010190602001808311610e4e57829003601f168201915b505050505081526020019060010190610dd3565b505050509050919050565b6001600160a01b0382166000908152600b6020908152604082208054600181018255908352918190208351610d1a939190910191840190615a49565b6060600080600086516001600160401b03811115610ef457634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610f2757816020015b6060815260200190600190039081610f125790505b5090506000808060008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b158015610f7e57600080fd5b505afa158015610f92573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fb69190617019565b905060005b8b51811015611168576000610ff68d8381518110610fe957634e487b7160e01b600052603260045260246000fd5b60200260200101516147e5565b90506000805b8c51811015611092578c818151811061102557634e487b7160e01b600052603260045260246000fd5b6020026020010151600181111561104c57634e487b7160e01b600052602160045260246000fd5b836101e00151600181111561107157634e487b7160e01b600052602160045260246000fd5b14156110805760019150611092565b8061108a81618148565b915050610ffc565b508061109f575050611156565b438460c001516001600160401b03168361010001516110be9190617f71565b11156110cb575050611156565b6101408201516110db9087617f89565b95506110e982600180612bd7565b8d838151811061110957634e487b7160e01b600052603260045260246000fd5b602002602001015188886001600160401b03168151811061113a57634e487b7160e01b600052603260045260246000fd5b6020026020010181905250868061115090618163565b97505050505b8061116081618148565b915050610fbb565b506001600160401b038316156111b8576040516001600160a01b038b16906001600160401b03851680156108fc02916000818181858888f193505050501580156111b6573d6000803e3d6000fd5b505b50929990985091965090945050505050565b6006816040516111da9190617be1565b90815260405190819003602001902060006111f58282615a0c565b6001820180546001600160a01b0319169055611215600283016000615a0c565b60006003830181905560048301805467ffffffffffffffff19169055600583018190556006830180546001600160801b0319169055611258906007840190615a0c565b60088201805467ffffffffffffffff19169055600060098301819055600a8301805469ffffffffffffffffffff1916905561129790600b840190615ac9565b6112a5600c83016000615ac9565b6112b3600d83016000615a0c565b50600e8101805461ffff19169055600f0180547fffffffffffffffff000000000000000000000000000000000000000000000000169055600480546040517f52e061460000000000000000000000000000000000000000000000000000000081526001600160a01b03909116916352e061469161133291859101617c9f565b600060405180830381600087803b15801561134c57600080fd5b505af1158015611360573d6000803e3d6000fd5b5050505050565b60606000606060005b865181101561141d57600061139e888381518110610fe957634e487b7160e01b600052603260045260246000fd5b90506000816101e0015160018111156113c757634e487b7160e01b600052602160045260246000fd5b146113d2575061140b565b80610100015186116113e4575061140b565b60006113f188888461253e565b905080611408578360009550955050505050611425565b50505b8061141581618148565b915050611370565b509150600190505b935093915050565b604080516060810182526000808252602080830182905292820152908201516001600160401b03166114715760405162461bcd60e51b8152600401610af190617e16565b60008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156114c157600080fd5b505afa1580156114d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114f99190617019565b9050610c318382436143d5565b60008082516001600160401b0381111561153057634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561156957816020015b611556615ae7565b81526020019060019003908161154e5790505b50905060005b83518110156115e657600061159d858381518110610fe957634e487b7160e01b600052603260045260246000fd5b905080602001519350808383815181106115c757634e487b7160e01b600052603260045260246000fd5b60200260200101819052505080806115de90618148565b91505061156f565b506115f081612dd7565b7fa5d0a6140e999641864ed8958af5eea1d36d821658cfc056552962fab40291ec60014385856040516116269493929190617cd0565b60405180910390a1505050565b6000600682600001516040516116499190617be1565b908152602001604051809103902060090154116116785760405162461bcd60e51b8152600401610af190617de6565b6001815160405160069161168b91617be1565b908152604051908190036020019020600a015460ff6101009091041660018111156116c657634e487b7160e01b600052602160045260246000fd5b146116e35760405162461bcd60e51b8152600401610af190617e06565b43600682600001516040516116f89190617be1565b908152602001604051809103902060050154116117275760405162461bcd60e51b8152600401610af190617dc6565b60008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561177757600080fd5b505afa15801561178b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117af9190617019565b90506000600683600001516040516117c79190617be1565b9081526020016040518091039020604051806102e00160405290816000820180546117f1906180f5565b80601f016020809104026020016040519081016040528092919081815260200182805461181d906180f5565b801561186a5780601f1061183f5761010080835404028352916020019161186a565b820191906000526020600020905b81548152906001019060200180831161184d57829003601f168201915b505050918352505060018201546001600160a01b0316602082015260028201805460409092019161189a906180f5565b80601f01602080910402602001604051908101604052809291908181526020018280546118c6906180f5565b80156119135780601f106118e857610100808354040283529160200191611913565b820191906000526020600020905b8154815290600101906020018083116118f657829003601f168201915b505050918352505060038201546001600160401b038082166020840152600160401b80830482166040850152600160801b830482166060850152600160c01b909204811660808401526004840154811660a0840152600584015460c0840152600684015480821660e08501529190910416610100820152600782018054610120909201916119a0906180f5565b80601f01602080910402602001604051908101604052809291908181526020018280546119cc906180f5565b8015611a195780601f106119ee57610100808354040283529160200191611a19565b820191906000526020600020905b8154815290600101906020018083116119fc57829003601f168201915b505050918352505060088201546001600160401b0316602082015260098201546040820152600a82015460ff80821615156060840152608090920191610100909104166001811115611a7b57634e487b7160e01b600052602160045260246000fd5b6001811115611a9a57634e487b7160e01b600052602160045260246000fd5b8152600a8201546201000090046001600160401b0316602080830191909152600b830180546040805182850281018501825282815294019392830182828015611b0c57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611aee575b50505050508152602001600c8201805480602002602001604051908101604052809291908181526020018280548015611b6e57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611b50575b50505050508152602001600d82018054611b87906180f5565b80601f0160208091040260200160405190810160405280929190818152602001828054611bb3906180f5565b8015611c005780601f10611bd557610100808354040283529160200191611c00565b820191906000526020600020905b815481529060010190602001808311611be357829003601f168201915b5050509183525050600e82015460209091019060ff166002811115611c3557634e487b7160e01b600052602160045260246000fd5b6002811115611c5457634e487b7160e01b600052602160045260246000fd5b8152600e820154610100900460ff16151560208083019190915260408051606081018252600f909401546001600160401b038082168652600160401b8204811693860193909352600160801b9004909116838201529081019190915284015161012082015160a08301516080840151939450600093611cf193879390929091611cdd9190617fdd565b8660c0015189604001516106819190617fdd565b9050600081602001518260400151611d099190617f89565b9050806001600160401b0316341015611d525734816040517fb0b78f49000000000000000000000000000000000000000000000000000000008152600401610af1929190617ed1565b84604001518360e001818151611d689190617f89565b6001600160401b031690525061014083018051829190611d89908390617f89565b6001600160401b0316905250604085015160c0840151611da99190617fdd565b6001600160401b03168361010001818151611dc49190617f71565b90525084516040518491600691611ddb9190617be1565b90815260200160405180910390206000820151816000019080519060200190611e05929190615a49565b506020828101516001830180546001600160a01b0319166001600160a01b0390921691909117905560408301518051611e449260028501920190615a49565b506060820151600382018054608085015160a086015160c08701516001600160401b039586166001600160801b031994851617600160401b9387168402176fffffffffffffffffffffffffffffffff16600160801b9287169290920277ffffffffffffffffffffffffffffffffffffffffffffffff1691909117600160c01b918616919091021790925560e085015160048501805467ffffffffffffffff1916918516919091179055610100850151600585015561012085015160068501805461014088015192861693169290921793169091029190911790556101608201518051611f3a916007840191602090910190615a49565b5061018082015160088201805467ffffffffffffffff19166001600160401b039092169190911790556101a082015160098201556101c0820151600a8201805460ff19811692151592831782556101e08501519261ffff1990911661ff001990911617610100836001811115611fc057634e487b7160e01b600052602160045260246000fd5b0217905550610200820151600a820180546001600160401b03909216620100000269ffffffffffffffff000019909216919091179055610220820151805161201291600b840191602090910190615bb9565b50610240820151805161202f91600c840191602090910190615bb9565b50610260820151805161204c91600d840191602090910190615a49565b50610280820151600e8201805460ff1916600183600281111561207f57634e487b7160e01b600052602160045260246000fd5b02179055506102a0820151600e820180549115156101000261ff00199092169190911790556102c0909101518051600f909201805460208301516040909301516001600160401b03908116600160801b0267ffffffffffffffff60801b19948216600160401b026001600160801b0319909316919095161717919091169190911790555050505050565b60005b6001600160a01b038316600090815260096020526040902054811015610d1a57818051906020012060096000856001600160a01b03166001600160a01b03168152602001908152602001600020828154811061217857634e487b7160e01b600052603260045260246000fd5b9060005260206000200160405161218f9190617bed565b604051809103902014156121d5576001600160a01b0383166000908152600960205260409020805482908110610d0457634e487b7160e01b600052603260045260246000fd5b806121df81618148565b91505061210c565b80600682600001516040516121fc9190617be1565b90815260200160405180910390206000820151816000019080519060200190612226929190615a49565b506020828101516001830180546001600160a01b0319166001600160a01b03909216919091179055604083015180516122659260028501920190615a49565b506060820151600382018054608085015160a086015160c08701516001600160401b039586166001600160801b031994851617600160401b9387168402176fffffffffffffffffffffffffffffffff16600160801b9287169290920277ffffffffffffffffffffffffffffffffffffffffffffffff1691909117600160c01b918616919091021790925560e085015160048501805467ffffffffffffffff191691851691909117905561010085015160058501556101208501516006850180546101408801519286169316929092179316909102919091179055610160820151805161235b916007840191602090910190615a49565b5061018082015160088201805467ffffffffffffffff19166001600160401b039092169190911790556101a082015160098201556101c0820151600a8201805460ff19811692151592831782556101e08501519261ffff1990911661ff0019909116176101008360018111156123e157634e487b7160e01b600052602160045260246000fd5b0217905550610200820151600a820180546001600160401b03909216620100000269ffffffffffffffff000019909216919091179055610220820151805161243391600b840191602090910190615bb9565b50610240820151805161245091600c840191602090910190615bb9565b50610260820151805161246d91600d840191602090910190615a49565b50610280820151600e8201805460ff191660018360028111156124a057634e487b7160e01b600052602160045260246000fd5b02179055506102a0820151600e820180549115156101000261ff00199092169190911790556102c0909101518051600f909201805460208301516040909301516001600160401b03908116600160801b0267ffffffffffffffff60801b19948216600160401b026001600160801b03199093169190951617179190911691909117905550565b6000612533858484614d33565b610bcd856001617f89565b610100810182905260006125c4604080516101e08101825260608082526000602083018190529282018390528082018390526080820183905260a0820183905260c0820183905260e0820183905261010082018190526101208201839052610140820183905261016082018190526101808201526101a08101829052906101c082015290565b61010083015160808083019190915260c0808501516001600160401b039081166040850152610120860151169083015283015160a08401516126069190617fdd565b6001600160401b031660208201526101a08301516000612627838884610be2565b905060008160400151826020015183600001516126449190617f89565b61264e9190617f89565b90508561014001516001600160401b0316816001600160401b03161161267b576000945050505050610c31565b6001600160401b0381166101408701526126958484614d5e565b6001600160401b031660e087015250600193505050509392505050565b60006126be8483614283565b610bda9084617fdd565b6001600160a01b0381166000908152600960209081526040808320805482518185028101850190935280835260609493849084015b828210156127a957838290600052602060002001805461271c906180f5565b80601f0160208091040260200160405190810160405280929190818152602001828054612748906180f5565b80156127955780601f1061276a57610100808354040283529160200191612795565b820191906000526020600020905b81548152906001019060200180831161277857829003601f168201915b5050505050815260200190600101906126fd565b505050509050600081516001600160401b038111156127d857634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561280b57816020015b60608152602001906001900390816127f65790505b5090506000805b83518110156129c65760045484516000916001600160a01b03169063977fdfe29087908590811061285357634e487b7160e01b600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b81526004016128779190617c9f565b60006040518083038186803b15801561288f57600080fd5b505afa1580156128a3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526128cb9190810190616dc1565b90506000805b825181101561293a57886001600160a01b031683828151811061290457634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516001600160a01b03161415612928576001915061293a565b8061293281618148565b9150506128d1565b50806129475750506129b4565b85838151811061296757634e487b7160e01b600052603260045260246000fd5b602002602001015185856001600160401b03168151811061299857634e487b7160e01b600052603260045260246000fd5b602002602001018190525083806129ae90618163565b94505050505b806129be81618148565b915050612812565b5090949350505050565b6006816040516129e09190617be1565b90815260405190819003602001902060006129fb8282615a0c565b6001820180546001600160a01b0319169055612a1b600283016000615a0c565b60006003830181905560048301805467ffffffffffffffff19169055600583018190556006830180546001600160801b0319169055612a5e906007840190615a0c565b60088201805467ffffffffffffffff19169055600060098301819055600a8301805469ffffffffffffffffffff19169055612a9d90600b840190615ac9565b612aab600c83016000615ac9565b612ab9600d83016000615a0c565b50600e8101805461ffff19169055600f0180547fffffffffffffffff00000000000000000000000000000000000000000000000016905550565b6001600160a01b0381166000908152600860209081526040808320805482518185028101850190935280835260609492939192909184015b82821015610e7f578382906000526020600020018054612b4a906180f5565b80601f0160208091040260200160405190810160405280929190818152602001828054612b76906180f5565b8015612bc35780601f10612b9857610100808354040283529160200191612bc3565b820191906000526020600020905b815481529060010190602001808311612ba657829003601f168201915b505050505081526020019060010190612b2b565b82518215612bff57612be8816129d0565b612bf1816111ca565b612bff846020015182612ccc565b8115612cc657612c13846020015182614d8a565b60005b84610220015151811015612c6c57612c5a8561022001518281518110612c4c57634e487b7160e01b600052603260045260246000fd5b602002602001015183612109565b80612c6481618148565b915050612c16565b5060005b8461024001515181101561136057612cb48561024001518281518110612ca657634e487b7160e01b600052603260045260246000fd5b602002602001015183610c38565b80612cbe81618148565b915050612c70565b50505050565b60005b6001600160a01b0383166000908152600b6020526040902054811015610d1a578180519060200120600b6000856001600160a01b03166001600160a01b031681526020019081526020016000208281548110612d3b57634e487b7160e01b600052603260045260246000fd5b90600052602060002001604051612d529190617bed565b60405180910390201415612d98576001600160a01b0383166000908152600b60205260409020805482908110610d0457634e487b7160e01b600052603260045260246000fd5b80612da281618148565b915050612ccf565b6000612db982600001516147e5565b60208301516001600160401b031660608201529050610d97816121e7565b8051612de05750565b60008082600081518110612e0457634e487b7160e01b600052603260045260246000fd5b602002602001015160200151905060008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b158015612e6257600080fd5b505afa158015612e76573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e9a9190617019565b905060005b8451811015612f36576000858281518110612eca57634e487b7160e01b600052603260045260246000fd5b60200260200101519050836001600160a01b031681602001516001600160a01b031614612f23576040517f2f7fecf500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5080612f2e81618148565b915050612e9f565b5060005b8451811015613718576000858281518110612f6557634e487b7160e01b600052603260045260246000fd5b60200260200101519050600060078260000151604051612f859190617be1565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b8282101561301357600084815260209081902060408051808201909152908401546001600160a01b03811682527401000000000000000000000000000000000000000090046001600160401b031681830152825260019092019101612fb3565b50505050905060005b81518110156131655760035482516000916001600160a01b031690632ba010d79085908590811061305d57634e487b7160e01b600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b81526004016130819190617e9a565b60006040518083038186803b15801561309957600080fd5b505afa1580156130ad573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526130d59190810190616fe5565b6003546040517e47c0030000000000000000000000000000000000000000000000000000000081529192506001600160a01b0316906247c0039061311f9084908890600401617e75565b600060405180830381600087803b15801561313957600080fd5b505af115801561314d573d6000803e3d6000fd5b5050505050808061315d90618148565b91505061301c565b506101408201516001600160401b031661318c5761318582600180612bd7565b5050613706565b6004805483516040517f977fdfe20000000000000000000000000000000000000000000000000000000081526000936001600160a01b039093169263977fdfe2926131d992909101617c9f565b60006040518083038186803b1580156131f157600080fd5b505afa158015613205573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261322d9190810190616dc1565b90506000836101400151905060008460a00151856080015161324f9190617fdd565b9050600061325d8883614283565b905060005b84518110156134965760015485516000916001600160a01b031690631a65374a908890859081106132a357634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516040518263ffffffff1660e01b81526004016132cb9190617bf9565b60e06040518083038186803b1580156132e357600080fd5b505afa1580156132f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061331b9190616f5f565b9050600083600188858151811061334257634e487b7160e01b600052603260045260246000fd5b6020026020010151604001516133589190618034565b6133629190617fdd565b9050600061337c8c878c6101a001514361074e9190618019565b9050600061338a8284617f89565b9050876001600160401b0316816001600160401b031611156133d8576040517f870303a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80846020018181516133ea9190617f89565b6001600160401b03169052506001546040517fc819d86c0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063c819d86c9061343f908790600401617e67565b600060405180830381600087803b15801561345957600080fd5b505af115801561346d573d6000803e3d6000fd5b50505050808861347d9190618034565b975050505050808061348e90618148565b915050613262565b506000866101e0015160018111156134be57634e487b7160e01b600052602160045260246000fd5b14156134d5576134ce838b617f89565b99506136f3565b6001866101e0015160018111156134fc57634e487b7160e01b600052602160045260246000fd5b14156136f3576002546020870151604051632a7069e160e11b81526000926001600160a01b0316916354e0d3c2916135379190600401617bf9565b60a06040518083038186803b15801561354f57600080fd5b505afa158015613563573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135879190617291565b90508660a00151876080015161359d9190617fdd565b6001600160401b031681600001516001600160401b03161061363c5783816040018181516135cb9190617f89565b6001600160401b031690525060a087015160808801516135eb9190617fdd565b816020018181516135fc9190617f89565b6001600160401b031690525060a0870151608088015161361c9190617fdd565b8151829061362b908390618034565b6001600160401b0316905250613671565b60016040517fc8c84b2f000000000000000000000000000000000000000000000000000000008152600401610af19190617da8565b60025460208801516040517fed108de90000000000000000000000000000000000000000000000000000000081526001600160a01b039092169163ed108de9916136bf918590600401617c07565b600060405180830381600087803b1580156136d957600080fd5b505af11580156136ed573d6000803e3d6000fd5b50505050505b6136ff86600180612bd7565b5050505050505b8061371081618148565b915050612f3a565b506001600160401b03831661372d5750505050565b6040516001600160a01b038316906001600160401b03851680156108fc02916000818181858888f19350505050158015611360573d6000803e3d6000fd5b6060600082511161378e5760405162461bcd60e51b8152600401610af190617db6565b600082516001600160401b038111156137b757634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156137f057816020015b6137dd615ae7565b8152602001906001900390816137d55790505b50905060005b835181101561427c57600084828151811061382157634e487b7160e01b600052603260045260246000fd5b60200260200101519050600060068260405161383d9190617be1565b9081526020016040518091039020604051806102e0016040529081600082018054613867906180f5565b80601f0160208091040260200160405190810160405280929190818152602001828054613893906180f5565b80156138e05780601f106138b5576101008083540402835291602001916138e0565b820191906000526020600020905b8154815290600101906020018083116138c357829003601f168201915b505050918352505060018201546001600160a01b03166020820152600282018054604090920191613910906180f5565b80601f016020809104026020016040519081016040528092919081815260200182805461393c906180f5565b80156139895780601f1061395e57610100808354040283529160200191613989565b820191906000526020600020905b81548152906001019060200180831161396c57829003601f168201915b505050918352505060038201546001600160401b038082166020840152600160401b80830482166040850152600160801b830482166060850152600160c01b909204811660808401526004840154811660a0840152600584015460c0840152600684015480821660e0850152919091041661010082015260078201805461012090920191613a16906180f5565b80601f0160208091040260200160405190810160405280929190818152602001828054613a42906180f5565b8015613a8f5780601f10613a6457610100808354040283529160200191613a8f565b820191906000526020600020905b815481529060010190602001808311613a7257829003601f168201915b505050918352505060088201546001600160401b0316602082015260098201546040820152600a82015460ff80821615156060840152608090920191610100909104166001811115613af157634e487b7160e01b600052602160045260246000fd5b6001811115613b1057634e487b7160e01b600052602160045260246000fd5b8152600a8201546201000090046001600160401b0316602080830191909152600b830180546040805182850281018501825282815294019392830182828015613b8257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613b64575b50505050508152602001600c8201805480602002602001604051908101604052809291908181526020018280548015613be457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613bc6575b50505050508152602001600d82018054613bfd906180f5565b80601f0160208091040260200160405190810160405280929190818152602001828054613c29906180f5565b8015613c765780601f10613c4b57610100808354040283529160200191613c76565b820191906000526020600020905b815481529060010190602001808311613c5957829003601f168201915b5050509183525050600e82015460209091019060ff166002811115613cab57634e487b7160e01b600052602160045260246000fd5b6002811115613cca57634e487b7160e01b600052602160045260246000fd5b8152600e820154610100900460ff16151560208083019190915260408051606081018252600f909401546001600160401b038082168652600160401b8204811693860193909352600160801b9004909116838201520152805151909150613d465781604051636c5249c160e01b8152600401610af19190617c9f565b600682604051613d569190617be1565b9081526020016040518091039020604051806102e0016040529081600082018054613d80906180f5565b80601f0160208091040260200160405190810160405280929190818152602001828054613dac906180f5565b8015613df95780601f10613dce57610100808354040283529160200191613df9565b820191906000526020600020905b815481529060010190602001808311613ddc57829003601f168201915b505050918352505060018201546001600160a01b03166020820152600282018054604090920191613e29906180f5565b80601f0160208091040260200160405190810160405280929190818152602001828054613e55906180f5565b8015613ea25780601f10613e7757610100808354040283529160200191613ea2565b820191906000526020600020905b815481529060010190602001808311613e8557829003601f168201915b505050918352505060038201546001600160401b038082166020840152600160401b80830482166040850152600160801b830482166060850152600160c01b909204811660808401526004840154811660a0840152600584015460c0840152600684015480821660e0850152919091041661010082015260078201805461012090920191613f2f906180f5565b80601f0160208091040260200160405190810160405280929190818152602001828054613f5b906180f5565b8015613fa85780601f10613f7d57610100808354040283529160200191613fa8565b820191906000526020600020905b815481529060010190602001808311613f8b57829003601f168201915b505050918352505060088201546001600160401b0316602082015260098201546040820152600a82015460ff8082161515606084015260809092019161010090910416600181111561400a57634e487b7160e01b600052602160045260246000fd5b600181111561402957634e487b7160e01b600052602160045260246000fd5b8152600a8201546201000090046001600160401b0316602080830191909152600b83018054604080518285028101850182528281529401939283018282801561409b57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161407d575b50505050508152602001600c82018054806020026020016040519081016040528092919081815260200182805480156140fd57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116140df575b50505050508152602001600d82018054614116906180f5565b80601f0160208091040260200160405190810160405280929190818152602001828054614142906180f5565b801561418f5780601f106141645761010080835404028352916020019161418f565b820191906000526020600020905b81548152906001019060200180831161417257829003601f168201915b5050509183525050600e82015460209091019060ff1660028111156141c457634e487b7160e01b600052602160045260246000fd5b60028111156141e357634e487b7160e01b600052602160045260246000fd5b8152600e820154610100900460ff16151560208083019190915260408051606081018252600f909401546001600160401b038082168652600160401b8204811693860193909352600160801b9004909116838201520152845185908590811061425c57634e487b7160e01b600052603260045260246000fd5b60200260200101819052505050808061427490618148565b9150506137f6565b5092915050565b6000620fa0008284606001516142999190617fdd565b6142a39190617fbc565b90505b92915050565b60408051606080820183526000808352602080840182905283850182905284519283018552818352820181905292810183905290916142ed88888888610bb5565b905060006142fd89888888612526565b6001600160401b03928316604085015290911660208301525090505b95945050505050565b600061432d826147e5565b60408051600180825281830190925291925060009190816020015b614350615ae7565b815260200190600190039081614348579050509050818160008151811061438757634e487b7160e01b600052603260045260246000fd5b602002602001018190525061439b81612dd7565b7fb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f456001438585602001516040516116269493929190617d0c565b60408051606081018252600080825260208201819052918101919091526101808401515160009062989680901561441857614411816004617fdd565b9150614426565b614423816003617fdd565b91505b604080516060810182526000602082018190529181018290526001600160401b03841681526101c0880151909190600181111561447357634e487b7160e01b600052602160045260246000fd5b1415614483579250610c31915050565b6000614490888888610be2565b6040818101516001600160401b03908116918501919091526020918201511690830152509695505050505050565b6001600160a01b03821660009081526008602090815260409091208251610d1a928401906159af565b6001600160a01b0381166000908152600a60209081526040808320805482518185028101850190935280835260609493849084015b828210156145c857838290600052602060002001805461453b906180f5565b80601f0160208091040260200160405190810160405280929190818152602001828054614567906180f5565b80156145b45780601f10614589576101008083540402835291602001916145b4565b820191906000526020600020905b81548152906001019060200180831161459757829003601f168201915b50505050508152602001906001019061451c565b505050509050600081516001600160401b038111156145f757634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561462a57816020015b60608152602001906001900390816146155790505b5090506000805b83518110156129c65760045484516000916001600160a01b03169063977fdfe29087908590811061467257634e487b7160e01b600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b81526004016146969190617c9f565b60006040518083038186803b1580156146ae57600080fd5b505afa1580156146c2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526146ea9190810190616dc1565b90506000805b825181101561475957886001600160a01b031683828151811061472357634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516001600160a01b031614156147475760019150614759565b8061475181618148565b9150506146f0565b50806147665750506147d3565b85838151811061478657634e487b7160e01b600052603260045260246000fd5b602002602001015185856001600160401b0316815181106147b757634e487b7160e01b600052603260045260246000fd5b602002602001018190525083806147cd90618163565b94505050505b806147dd81618148565b915050614631565b6147ed615ae7565b81600081511161480f5760405162461bcd60e51b8152600401610af190617e26565b60006006846040516148219190617be1565b9081526020016040518091039020604051806102e001604052908160008201805461484b906180f5565b80601f0160208091040260200160405190810160405280929190818152602001828054614877906180f5565b80156148c45780601f10614899576101008083540402835291602001916148c4565b820191906000526020600020905b8154815290600101906020018083116148a757829003601f168201915b505050918352505060018201546001600160a01b031660208201526002820180546040909201916148f4906180f5565b80601f0160208091040260200160405190810160405280929190818152602001828054614920906180f5565b801561496d5780601f106149425761010080835404028352916020019161496d565b820191906000526020600020905b81548152906001019060200180831161495057829003601f168201915b505050918352505060038201546001600160401b038082166020840152600160401b80830482166040850152600160801b830482166060850152600160c01b909204811660808401526004840154811660a0840152600584015460c0840152600684015480821660e08501529190910416610100820152600782018054610120909201916149fa906180f5565b80601f0160208091040260200160405190810160405280929190818152602001828054614a26906180f5565b8015614a735780601f10614a4857610100808354040283529160200191614a73565b820191906000526020600020905b815481529060010190602001808311614a5657829003601f168201915b505050918352505060088201546001600160401b0316602082015260098201546040820152600a82015460ff80821615156060840152608090920191610100909104166001811115614ad557634e487b7160e01b600052602160045260246000fd5b6001811115614af457634e487b7160e01b600052602160045260246000fd5b8152600a8201546201000090046001600160401b0316602080830191909152600b830180546040805182850281018501825282815294019392830182828015614b6657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311614b48575b50505050508152602001600c8201805480602002602001604051908101604052809291908181526020018280548015614bc857602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311614baa575b50505050508152602001600d82018054614be1906180f5565b80601f0160208091040260200160405190810160405280929190818152602001828054614c0d906180f5565b8015614c5a5780601f10614c2f57610100808354040283529160200191614c5a565b820191906000526020600020905b815481529060010190602001808311614c3d57829003601f168201915b5050509183525050600e82015460209091019060ff166002811115614c8f57634e487b7160e01b600052602160045260246000fd5b6002811115614cae57634e487b7160e01b600052602160045260246000fd5b8152600e820154610100900460ff16151560208083019190915260408051606081018252600f909401546001600160401b038082168652600160401b8204811693860193909352600160801b9004909116838201520152805151909150614d2a5783604051636c5249c160e01b8152600401610af19190617c9f565b91505b50919050565b6000620fa00082848660200151614d4a9190617fdd565b614d549190617fdd565b610bda9190617fbc565b60008260400151828460800151614d759190618019565b614d7f9190617fbc565b6142a3906001617f89565b60005b6001600160a01b038316600090815260086020526040902054811015610d1a57818051906020012060086000856001600160a01b03166001600160a01b031681526020019081526020016000208281548110614df957634e487b7160e01b600052603260045260246000fd5b90600052602060002001604051614e109190617bed565b60405180910390201415614e56576001600160a01b0383166000908152600860205260409020805482908110610d0457634e487b7160e01b600052603260045260246000fd5b80614e6081618148565b915050614d8d565b8051604051600691614e7991617be1565b908152602001604051809103902060090154600014614eaa5760405162461bcd60e51b8152600401610af190617e36565b4381610100015111614ece5760405162461bcd60e51b8152600401610af190617dc6565b60008060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b158015614f1e57600080fd5b505afa158015614f32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614f569190617019565b905060008261028001516002811115614f7f57634e487b7160e01b600052602160045260246000fd5b1415614f97576005546001600160401b031660c08301525b60018261028001516002811115614fbe57634e487b7160e01b600052602160045260246000fd5b1415614fd6576005546001600160401b031660c08301525b60028261028001516002811115614ffd57634e487b7160e01b600052602160045260246000fd5b1415615015576005546001600160401b031660c08301525b60016101c0830152615099604080516101e08101825260608082526000602083018190529282018390528082018390526080820183905260a0820183905260c0820183905260e0820183905261010082018190526101208201839052610140820183905261016082018190526101808201526101a08101829052906101c082015290565b61010083015160808083019190915260c0808501516001600160401b039081166040850152610120860151169083015283015160a08401516150db9190617fdd565b6001600160401b0316602082015260006150f6828443610be2565b90508060400151816020015182600001516151119190617f89565b61511b9190617f89565b6001600160401b03166101408501526151348243614d5e565b6001600160401b031660e08501526000846101e00151600181111561516957634e487b7160e01b600052602160045260246000fd5b14156153a1576002546020850151604051632a7069e160e11b81526000926001600160a01b0316916354e0d3c2916151a49190600401617bf9565b60a06040518083038186803b1580156151bc57600080fd5b505afa1580156151d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906151f49190617291565b90508461014001516001600160401b031681604001516001600160401b031610156152585780604001518561014001516040517f62fe63d9000000000000000000000000000000000000000000000000000000008152600401610af1929190617efa565b84608001518560a0015161526c9190617fdd565b6001600160401b031681602001516001600160401b031610156152c75780602001518560a001516040517fd39dfd03000000000000000000000000000000000000000000000000000000008152600401610af1929190617efa565b846101000151816060015110156153175780606001518561010001516040517f0c803412000000000000000000000000000000000000000000000000000000008152600401610af1929190617eb6565b8461014001518160400181815161532e9190618034565b6001600160401b031690525060a0850151608086015161534e9190617fdd565b8160200181815161535f9190618034565b6001600160401b031690525060a0850151608086015161537f9190617fdd565b8151829061538e908390617f89565b6001600160401b03169052506153d89050565b8361014001516001600160401b03163410156153cf5760405162461bcd60e51b8152600401610af190617e46565b60016101e08501525b60808301516001600160401b0316610180850152436101a0850152835160405185916006916154079190617be1565b90815260200160405180910390206000820151816000019080519060200190615431929190615a49565b506020828101516001830180546001600160a01b0319166001600160a01b03909216919091179055604083015180516154709260028501920190615a49565b506060820151600382018054608085015160a086015160c08701516001600160401b039586166001600160801b031994851617600160401b9387168402176fffffffffffffffffffffffffffffffff16600160801b9287169290920277ffffffffffffffffffffffffffffffffffffffffffffffff1691909117600160c01b918616919091021790925560e085015160048501805467ffffffffffffffff1916918516919091179055610100850151600585015561012085015160068501805461014088015192861693169290921793169091029190911790556101608201518051615566916007840191602090910190615a49565b5061018082015160088201805467ffffffffffffffff19166001600160401b039092169190911790556101a082015160098201556101c0820151600a8201805460ff19811692151592831782556101e08501519261ffff1990911661ff0019909116176101008360018111156155ec57634e487b7160e01b600052602160045260246000fd5b0217905550610200820151600a820180546001600160401b03909216620100000269ffffffffffffffff000019909216919091179055610220820151805161563e91600b840191602090910190615bb9565b50610240820151805161565b91600c840191602090910190615bb9565b50610260820151805161567891600d840191602090910190615a49565b50610280820151600e8201805460ff191660018360028111156156ab57634e487b7160e01b600052602160045260246000fd5b02179055506102a0820151600e820180549115156101000261ff00199092169190911790556102c0909101518051600f90920180546020808401516040948501516001600160401b03908116600160801b0267ffffffffffffffff60801b19928216600160401b026001600160801b031990951691909716179290921791909116939093179055858201516001600160a01b0316600090815260088352908120865181546001810183558284529284902081519294615771949190910192910190615a49565b5060005b8561022001515181101561580d5760006008600088610220015184815181106157ae57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b03168252818101929092526040016000908120895181546001810183558284529284902081519295506157f8949301920190615a49565b5050808061580590618148565b915050615775565b5060005b856102400151518110156158a957600060086000886102400151848151811061584a57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0316825281810192909252604001600090812089518154600181018355828452928490208151929550615894949301920190615a49565b505080806158a190618148565b915050615811565b50604080518082018252600080825260208201526101208701516001600160401b0316815260048054885193517fbb4e4e4200000000000000000000000000000000000000000000000000000000815292936001600160a01b039091169263bb4e4e429261591a9291869101617cb0565b600060405180830381600087803b15801561593457600080fd5b505af1158015615948573d6000803e3d6000fd5b505050507fa08fe92f4f83b40431d08f9f130fd22b658ad78cf83027611d629d83427f0d1860004388600001518961020001518a602001518b61014001518c6102a0015160405161599f9796959493929190617d39565b60405180910390a1505050505050565b8280548282559060005260206000209081019282156159fc579160200282015b828111156159fc57825180516159ec918491602090910190615a49565b50916020019190600101906159cf565b50615a08929150615c0e565b5090565b508054615a18906180f5565b6000825580601f10615a28575050565b601f016020900490600052602060002090810190615a469190615c2b565b50565b828054615a55906180f5565b90600052602060002090601f016020900481019282615a775760008555615abd565b82601f10615a9057805160ff1916838001178555615abd565b82800160010185558215615abd579182015b82811115615abd578251825591602001919060010190615aa2565b50615a08929150615c2b565b5080546000825590600052602060002090810190615a469190615c2b565b604080516102e08101825260608082526000602083018190529282018190528082018390526080820183905260a0820183905260c0820183905260e08201839052610100820183905261012082018390526101408201839052610160820181905261018082018390526101a082018390526101c082018390526101e0820183905261020082018390526102208201819052610240820181905261026082015290610280820190815260006020808301829052604080516060810182528381529182018390528181019290925291015290565b828054828255906000526020600020908101928215615abd579160200282015b82811115615abd57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190615bd9565b80821115615a08576000615c228282615a0c565b50600101615c0e565b5b80821115615a085760008155600101615c2c565b6000615c53615c4e84617f24565b617f08565b90508083825260208201905082856020860282011115615c7257600080fd5b60005b85811015615c9e5781615c888882615f71565b8452506020928301929190910190600101615c75565b5050509392505050565b6000615cb6615c4e84617f24565b90508083825260208201905082856020860282011115615cd557600080fd5b60005b85811015615c9e5781356001600160401b03811115615cf657600080fd5b808601615d038982616084565b855250506020928301929190910190600101615cd8565b6000615d28615c4e84617f24565b90508083825260208201905082856020860282011115615d4757600080fd5b60005b85811015615c9e5781516001600160401b03811115615d6857600080fd5b808601615d7589826160a5565b855250506020928301929190910190600101615d4a565b6000615d9a615c4e84617f24565b90508083825260208201905082856020860282011115615db957600080fd5b60005b85811015615c9e5781615dcf88826160e7565b8452506020928301929190910190600101615dbc565b6000615df3615c4e84617f24565b90508083825260208201905082856020860282011115615e1257600080fd5b60005b85811015615c9e5781356001600160401b03811115615e3357600080fd5b808601615e4089826160f2565b855250506020928301929190910190600101615e15565b6000615e65615c4e84617f24565b8381529050602081018260a085028101861015615e8157600080fd5b60005b85811015615c9e5781615e97888261659f565b84525060209092019160a09190910190600101615e84565b6000615ebd615c4e84617f24565b83815290506020810182606085028101861015615ed957600080fd5b60005b85811015615c9e5781615eef8882616b78565b84525060209092019160609190910190600101615edc565b6000615f15615c4e84617f47565b905082815260208101848484011115615f2d57600080fd5b615f388482856180bd565b509392505050565b6000615f4e615c4e84617f47565b905082815260208101848484011115615f6657600080fd5b615f388482856180c9565b80356142a681618222565b80516142a681618222565b600082601f830112615f9857600080fd5b8135610bda848260208601615c40565b600082601f830112615fb957600080fd5b8135610bda848260208601615ca8565b600082601f830112615fda57600080fd5b8151610bda848260208601615d1a565b600082601f830112615ffb57600080fd5b8135610bda848260208601615d8c565b600082601f83011261601c57600080fd5b8135610bda848260208601615de5565b600082601f83011261603d57600080fd5b8151610bda848260208601615e57565b600082601f83011261605e57600080fd5b8135610bda848260208601615eaf565b80356142a681618236565b80516142a681618236565b600082601f83011261609557600080fd5b8135610bda848260208601615f07565b600082601f8301126160b657600080fd5b8151610bda848260208601615f40565b80356142a68161823e565b80356142a681618247565b80516142a681618247565b80356142a681618254565b6000610320828403121561610557600080fd5b6161106102e0617f08565b905081356001600160401b0381111561612857600080fd5b61613484828501616084565b825250602061614584848301615f71565b60208301525060408201356001600160401b0381111561616457600080fd5b61617084828501616084565b604083015250606061618484828501616bb8565b606083015250608061619884828501616bb8565b60808301525060a06161ac84828501616bb8565b60a08301525060c06161c084828501616bb8565b60c08301525060e06161d484828501616bb8565b60e0830152506101006161e984828501616ba2565b610100830152506101206161ff84828501616bb8565b6101208301525061014061621584828501616bb8565b610140830152506101608201356001600160401b0381111561623657600080fd5b61624284828501616084565b6101608301525061018061625884828501616bb8565b610180830152506101a061626e84828501616ba2565b6101a0830152506101c06162848482850161606e565b6101c0830152506101e061629a848285016160e7565b6101e0830152506102006162b084828501616bb8565b610200830152506102208201356001600160401b038111156162d157600080fd5b6162dd84828501615f87565b610220830152506102408201356001600160401b038111156162fe57600080fd5b61630a84828501615f87565b610240830152506102608201356001600160401b0381111561632b57600080fd5b61633784828501616084565b6102608301525061028061634d848285016160d1565b610280830152506102a06163638482850161606e565b6102a0830152506102c061637984828501616507565b6102c08301525092915050565b60006060828403121561639857600080fd5b6163a26060617f08565b905081356001600160401b038111156163ba57600080fd5b6163c684828501616084565b82525060206163d784848301615f71565b60208301525060406163eb84828501616bb8565b60408301525092915050565b600060e0828403121561640957600080fd5b61641360e0617f08565b905060006164218484616bc3565b825250602061643284848301616bc3565b602083015250604061644684828501616bc3565b604083015250606061645a84828501616bc3565b606083015250608061646e84828501616bc3565b60808301525060a061648284828501615f7c565b60a08301525060c061649684828501615f7c565b60c08301525092915050565b6000606082840312156164b457600080fd5b6164be6060617f08565b905081356001600160401b038111156164d657600080fd5b6164e284828501616084565b82525060206164f384848301615f71565b60208301525060406163eb84828501615f71565b60006060828403121561651957600080fd5b6165236060617f08565b905060006165318484616bb8565b82525060206163d784848301616bb8565b60006040828403121561655457600080fd5b61655e6040617f08565b905081356001600160401b0381111561657657600080fd5b61658284828501616084565b825250602061659384848301616bb8565b60208301525092915050565b600060a082840312156165b157600080fd5b6165bb60a0617f08565b905060006165c98484615f7c565b82525060206165da84848301615f7c565b60208301525060406165ee84828501616bc3565b604083015250606061660284828501616bad565b606083015250608061661684828501616079565b60808301525092915050565b6000610180828403121561663557600080fd5b616640610180617f08565b9050600061664e8484615f7c565b825250602061665f84848301616bc3565b602083015250604061667384828501616bc3565b604083015250606061668784828501616bc3565b606083015250608061669b848285016160dc565b60808301525060a06166af84828501616bad565b60a08301525060c06166c384828501616bad565b60c08301525060e06166d784828501616bc3565b60e0830152506101006166ec84828501616bc3565b6101008301525061012061670284828501616bc3565b6101208301525061014061671884828501616079565b610140830152506101608201516001600160401b0381111561673957600080fd5b61674584828501615fc9565b6101608301525092915050565b6000610160828403121561676557600080fd5b616770610160617f08565b9050600061677e8484616bb8565b825250602061678f84848301616bb8565b60208301525060406167a384828501616bb8565b60408301525060606167b784828501616bb8565b60608301525060806167cb84828501616bb8565b60808301525060a06167df84828501616bb8565b60a08301525060c06167f384828501616bb8565b60c08301525060e061680784828501616bb8565b60e08301525061010061681c84828501616bb8565b6101008301525061012061683284828501616bb8565b6101208301525061014061684884828501616bb8565b6101408301525092915050565b6000610160828403121561686857600080fd5b616873610160617f08565b905060006168818484616bc3565b825250602061689284848301616bc3565b60208301525060406168a684828501616bc3565b60408301525060606168ba84828501616bc3565b60608301525060806168ce84828501616bc3565b60808301525060a06168e284828501616bc3565b60a08301525060c06168f684828501616bc3565b60c08301525060e061690a84828501616bc3565b60e08301525061010061691f84828501616bc3565b6101008301525061012061693584828501616bc3565b6101208301525061014061684884828501616bc3565b60006101e0828403121561695e57600080fd5b6169696101e0617f08565b905081356001600160401b0381111561698157600080fd5b61698d84828501616084565b825250602061699e84848301616bb8565b60208301525060406169b284828501616bb8565b60408301525060606169c684828501616bb8565b60608301525060806169da84828501616ba2565b60808301525060a06169ee84828501616bb8565b60a08301525060c0616a0284828501616bb8565b60c08301525060e0616a168482850161606e565b60e0830152506101008201356001600160401b03811115616a3657600080fd5b616a4284828501616084565b61010083015250610120616a588482850161606e565b61012083015250610140616a6e8482850161606e565b610140830152506101608201356001600160401b03811115616a8f57600080fd5b616a9b84828501616084565b610160830152506101808201356001600160401b03811115616abc57600080fd5b616ac88482850161604d565b610180830152506101a0616ade8482850161606e565b6101a0830152506101c0616af4848285016160e7565b6101c08301525092915050565b600060a08284031215616b1357600080fd5b616b1d60a0617f08565b90506000616b2b8484616bc3565b8252506020616b3c84848301616bc3565b6020830152506040616b5084828501616bc3565b6040830152506060616b6484828501616bad565b606083015250608061661684828501616bad565b600060608284031215616b8a57600080fd5b616b946060617f08565b905060006165318484615f71565b80356142a681618261565b80516142a681618261565b80356142a681618267565b80516142a681618267565b600060208284031215616be057600080fd5b6000610bda8484615f71565b60008060408385031215616bff57600080fd5b6000616c0b8585615f71565b92505060208301356001600160401b03811115616c2757600080fd5b616c3385828601615fa8565b9150509250929050565b60008060408385031215616c5057600080fd5b6000616c5c8585615f71565b92505060208301356001600160401b03811115616c7857600080fd5b616c3385828601616084565b600060208284031215616c9657600080fd5b81356001600160401b03811115616cac57600080fd5b610bda84828501615fa8565b600080600060608486031215616ccd57600080fd5b83356001600160401b03811115616ce357600080fd5b616cef86828701615fa8565b9350506020616d0086828701615f71565b92505060408401356001600160401b03811115616d1c57600080fd5b616d2886828701615fea565b9150509250925092565b60008060006101a08486031215616d4857600080fd5b83356001600160401b03811115616d5e57600080fd5b616d6a86828701615fa8565b9350506020616d7b86828701616752565b925050610180616d2886828701616ba2565b600060208284031215616d9f57600080fd5b81356001600160401b03811115616db557600080fd5b610bda8482850161600b565b600060208284031215616dd357600080fd5b81516001600160401b03811115616de957600080fd5b610bda8482850161602c565b600060208284031215616e0757600080fd5b81356001600160401b03811115616e1d57600080fd5b610bda84828501616084565b600080600080600060a08688031215616e4157600080fd5b6000616e4d88886160c6565b9550506020616e5e888289016160c6565b9450506040616e6f888289016160c6565b9350506060616e80888289016160c6565b9250506080616e91888289016160c6565b9150509295509295909350565b600060208284031215616eb057600080fd5b81356001600160401b03811115616ec657600080fd5b610bda848285016160f2565b600080600060608486031215616ee757600080fd5b83356001600160401b03811115616efd57600080fd5b616f09868287016160f2565b9350506020616f1a8682870161606e565b9250506040616d288682870161606e565b600060208284031215616f3d57600080fd5b81356001600160401b03811115616f5357600080fd5b610bda84828501616386565b600060e08284031215616f7157600080fd5b6000610bda84846163f7565b600060208284031215616f8f57600080fd5b81356001600160401b03811115616fa557600080fd5b610bda848285016164a2565b600060208284031215616fc357600080fd5b81356001600160401b03811115616fd957600080fd5b610bda84828501616542565b600060208284031215616ff757600080fd5b81516001600160401b0381111561700d57600080fd5b610bda84828501616622565b6000610160828403121561702c57600080fd5b6000610bda8484616855565b60008060006101a0848603121561704e57600080fd5b600061705a8686616752565b93505061016061706c86828701616ba2565b9250506101808401356001600160401b0381111561708957600080fd5b616d28868287016160f2565b60008061018083850312156170a957600080fd5b60006170b58585616752565b925050610160616c3385828601616bb8565b60008060006101a084860312156170dd57600080fd5b60006170e98686616752565b9350506101606170fb86828701616bb8565b925050610180616d2886828701616bb8565b6000806000806101c0858703121561712457600080fd5b60006171308787616752565b94505061016061714287828801616bb8565b93505061018061715487828801616bb8565b9250506101a061716687828801616bb8565b91505092959194509250565b60008060008060006101e0868803121561718b57600080fd5b60006171978888616752565b9550506101606171a988828901616bb8565b9450506101806171bb88828901616bb8565b9350506101a06171cd88828901616bb8565b9250506101c0616e9188828901616bb8565b6000602082840312156171f157600080fd5b81356001600160401b0381111561720757600080fd5b610bda8482850161694b565b60008060006101a0848603121561722957600080fd5b83356001600160401b0381111561723f57600080fd5b616d6a8682870161694b565b6000806040838503121561725e57600080fd5b82356001600160401b0381111561727457600080fd5b6172808582860161694b565b9250506020616c3385828601616ba2565b600060a082840312156172a357600080fd5b6000610bda8484616b01565b60006172bb83836172db565b505060200190565b60006142a38383617453565b60006142a38383617793565b6172e481618051565b82525050565b60006172f4825190565b80845260209384019383018060005b8381101561732857815161731788826172af565b975060208301925050600101617303565b509495945050505050565b600061733d825190565b808452602084019350836020820285016173578560200190565b8060005b8581101561738c578484038952815161737485826172c3565b94506020830160209a909a019992505060010161735b565b5091979650505050505050565b60006173a3825190565b808452602084019350836020820285016173bd8560200190565b8060005b8581101561738c57848403895281516173da85826172c3565b94506020830160209a909a01999250506001016173c1565b60006173fc825190565b808452602084019350836020820285016174168560200190565b8060005b8581101561738c578484038952815161743385826172cf565b94506020830160209a909a019992505060010161741a565b8015156172e4565b600061745d825190565b8084526020840193506174748185602086016180c9565b601f01601f19169290920192915050565b600061748f825190565b61749d8185602086016180c9565b9290920192915050565b600081546174b4816180f5565b6001821680156174cb57600181146174dc5761750c565b60ff1983168652818601935061750c565b60008581526020902060005b83811015617504578154888201526001909101906020016174e8565b838801955050505b50505092915050565b6172e48161808b565b6172e481618096565b6172e4816180a1565b6172e4816180ac565b601181526000602082017f66696c654c69737420697320656d707479000000000000000000000000000000815291505b5060200190565b600c81526000602082017f66696c652065787069726564000000000000000000000000000000000000000081529150617569565b602b81526000602082017f43757272656e74206f776e6572206d75737420626520746865206f776e65722081527f6f66207468652066696c65000000000000000000000000000000000000000000602082015291505b5060400190565b600e81526000602082017f66696c65206e6f7420657869737400000000000000000000000000000000000081529150617569565b602e81526000602082017f496e697469616c697a61626c653a20636f6e747261637420697320616c72656181527f647920696e697469616c697a6564000000000000000000000000000000000000602082015291506175fa565b600f81526000602082017f66696c652074797065206572726f72000000000000000000000000000000000081529150617569565b601f81526000602082017f66696c6553697a65206d7573742062652067726561746572207468616e20300081529150617569565b601a81526000602082017f66696c6548617368206d757374206265206e6f7420656d70747900000000000081529150617569565b601281526000602082017f66696c6520616c7265616479206578697374000000000000000000000000000081529150617569565b601481526000602082017f696e73756666696369656e74206465706f73697400000000000000000000000081529150617569565b8051610320808452600091908401906177ac8282617453565b91505060208301516177c160208601826172db565b50604083015184820360408601526177d98282617453565b91505060608301516177ee6060860182617bd2565b5060808301516178016080860182617bd2565b5060a083015161781460a0860182617bd2565b5060c083015161782760c0860182617bd2565b5060e083015161783a60e0860182617bd2565b5061010083015161784f610100860182617bcc565b50610120830151617864610120860182617bd2565b50610140830151617879610140860182617bd2565b506101608301518482036101608601526178938282617453565b9150506101808301516178aa610180860182617bd2565b506101a08301516178bf6101a0860182617bcc565b506101c08301516178d46101c086018261744b565b506101e08301516178e96101e0860182617527565b506102008301516178fe610200860182617bd2565b5061022083015184820361022086015261791882826172ea565b91505061024083015184820361024086015261793482826172ea565b9150506102608301518482036102608601526179508282617453565b91505061028083015161796761028086018261751e565b506102a083015161797c6102a086018261744b565b506102c0830151615f386102c0860182617a14565b805160e08301906179a28482617bd2565b5060208201516179b56020850182617bd2565b5060408201516179c86040850182617bd2565b5060608201516179db6060850182617bd2565b5060808201516179ee6080850182617bd2565b5060a0820151617a0160a08501826172db565b5060c0820151612cc660c08501826172db565b80516060830190617a258482617bd2565b506020820151617a386020850182617bd2565b506040820151612cc66040850182617bd2565b80516040830190617a5c8482617bd2565b506020820151612cc66020850182617bd2565b8051600090610180840190617a8485826172db565b506020830151617a976020860182617bd2565b506040830151617aaa6040860182617bd2565b506060830151617abd6060860182617bd2565b506080830151617ad0608086018261751e565b5060a0830151617ae360a0860182617bcc565b5060c0830151617af660c0860182617bcc565b5060e0830151617b0960e0860182617bd2565b50610100830151617b1e610100860182617bd2565b50610120830151617b33610120860182617bd2565b50610140830151617b4861014086018261744b565b506101608301518482036101608601526143198282617333565b80516040830190617a5c84826172db565b805160a0830190617b848482617bd2565b506020820151617b976020850182617bd2565b506040820151617baa6040850182617bd2565b506060820151617bbd6060850182617bcc565b506080820151612cc660808501825b806172e4565b6001600160401b0381166172e4565b6000610c318284617485565b6000610c3182846174a7565b602081016142a682846172db565b60c08101617c1582856172db565b610c316020830184617b73565b602080825281016142a38184617399565b60608082528101617c448186617399565b9050617c536020830185617bd2565b610bda604083018461744b565b602080825281016142a381846173f2565b60408082528101617c8281856173f2565b9050610c31602083018461744b565b602081016142a6828461744b565b602080825281016142a38184617453565b60608082528101617cc18185617453565b9050610c316020830184617a4b565b60808101617cde8287617515565b617ceb6020830186617bcc565b8181036040830152617cfd8185617399565b905061431960608301846172db565b60808101617d1a8287617515565b617d276020830186617bcc565b8181036040830152617cfd8185617453565b60e08101617d47828a617515565b617d546020830189617bcc565b8181036040830152617d668188617453565b9050617d756060830187617bd2565b617d8260808301866172db565b617d8f60a0830185617bd2565b617d9c60c083018461744b565b98975050505050505050565b602081016142a68284617530565b602080825281016142a681617539565b602080825281016142a681617570565b602080825281016142a6816175a4565b602080825281016142a681617601565b602080825281016142a681617635565b602080825281016142a68161768f565b602080825281016142a6816176c3565b602080825281016142a6816176f7565b602080825281016142a68161772b565b602080825281016142a68161775f565b602080825281016142a38184617793565b60e081016142a68284617991565b60408082528101617e868185617a6f565b90508181036020830152610bda8184617793565b604081016142a68284617b62565b606081016142a68284617a14565b60408101617ec48285617bcc565b610c316020830184617bcc565b60408101617edf8285617bcc565b610c316020830184617530565b602081016142a68284617bd2565b60408101617edf8285617530565b6000617f1360405190565b9050617f1f828261811c565b919050565b60006001600160401b03821115617f3d57617f3d6181dc565b5060209081020190565b60006001600160401b03821115617f6057617f606181dc565b601f19601f83011660200192915050565b60008219821115617f8457617f84618184565b500190565b60006001600160401b03821691506001600160401b0383169250826001600160401b0303821115617f8457617f84618184565b6001600160401b039182169116600082617fd857617fd861819a565b500490565b60006001600160401b03821691506001600160401b0383169250816001600160401b03048311821515161561801457618014618184565b500290565b6000825b92508282101561802f5761802f618184565b500390565b60006001600160401b03821691506001600160401b03831661801d565b60006001600160a01b0382166142a6565b60006142a682618051565b80617f1f816181f2565b80617f1f81618202565b80617f1f81618212565b60006142a68261806d565b60006142a682618077565b60006142a682618081565b60006001600160401b0382166142a6565b82818337506000910152565b60005b838110156180e45781810151838201526020016180cc565b83811115612cc65750506000910152565b60028104600182168061810957607f821691505b60208210811415614d2d57614d2d6181c6565b601f19601f83011681018181106001600160401b0382111715618141576181416181dc565b6040525050565b600060001982141561815c5761815c618184565b5060010190565b60006001600160401b03821691506001600160401b0382141561815c5761815c5b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600a8110615a4657615a466181b0565b60038110615a4657615a466181b0565b60028110615a4657615a466181b0565b61822b81618051565b8114615a4657600080fd5b80151561822b565b61822b81618062565b60038110615a4657600080fd5b60028110615a4657600080fd5b8061822b565b6001600160401b03811661822b56fea264697066735822122016df45ed6608f81a9e164b7a12b860ad5e0fc1f6da95620a6e868a39c3f7d91064736f6c63430008040033",
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
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// CalcDepositFee is a free data retrieval call binding the contract method 0x178e4dc0.
//
// Solidity: function CalcDepositFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) pure returns((uint64,uint64,uint64))
func (_Store *StoreCaller) CalcDepositFee(opts *bind.CallOpts, uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "CalcDepositFee", uploadOption, setting, currentHeight)

	if err != nil {
		return *new(StorageFee), err
	}

	out0 := *abi.ConvertType(out[0], new(StorageFee)).(*StorageFee)

	return out0, err

}

// CalcDepositFee is a free data retrieval call binding the contract method 0x178e4dc0.
//
// Solidity: function CalcDepositFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) pure returns((uint64,uint64,uint64))
func (_Store *StoreSession) CalcDepositFee(uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	return _Store.Contract.CalcDepositFee(&_Store.CallOpts, uploadOption, setting, currentHeight)
}

// CalcDepositFee is a free data retrieval call binding the contract method 0x178e4dc0.
//
// Solidity: function CalcDepositFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) pure returns((uint64,uint64,uint64))
func (_Store *StoreCallerSession) CalcDepositFee(uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	return _Store.Contract.CalcDepositFee(&_Store.CallOpts, uploadOption, setting, currentHeight)
}

// CalcFee is a free data retrieval call binding the contract method 0xcc76e80d.
//
// Solidity: function CalcFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 copyNum, uint64 fileSize, uint64 duration) pure returns((uint64,uint64,uint64))
func (_Store *StoreCaller) CalcFee(opts *bind.CallOpts, setting Setting, proveTime uint64, copyNum uint64, fileSize uint64, duration uint64) (StorageFee, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "CalcFee", setting, proveTime, copyNum, fileSize, duration)

	if err != nil {
		return *new(StorageFee), err
	}

	out0 := *abi.ConvertType(out[0], new(StorageFee)).(*StorageFee)

	return out0, err

}

// CalcFee is a free data retrieval call binding the contract method 0xcc76e80d.
//
// Solidity: function CalcFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 copyNum, uint64 fileSize, uint64 duration) pure returns((uint64,uint64,uint64))
func (_Store *StoreSession) CalcFee(setting Setting, proveTime uint64, copyNum uint64, fileSize uint64, duration uint64) (StorageFee, error) {
	return _Store.Contract.CalcFee(&_Store.CallOpts, setting, proveTime, copyNum, fileSize, duration)
}

// CalcFee is a free data retrieval call binding the contract method 0xcc76e80d.
//
// Solidity: function CalcFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 copyNum, uint64 fileSize, uint64 duration) pure returns((uint64,uint64,uint64))
func (_Store *StoreCallerSession) CalcFee(setting Setting, proveTime uint64, copyNum uint64, fileSize uint64, duration uint64) (StorageFee, error) {
	return _Store.Contract.CalcFee(&_Store.CallOpts, setting, proveTime, copyNum, fileSize, duration)
}

// CalcProveTimesByUploadInfo is a free data retrieval call binding the contract method 0xe4bca973.
//
// Solidity: function CalcProveTimesByUploadInfo((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, uint256 beginHeight) pure returns(uint64)
func (_Store *StoreCaller) CalcProveTimesByUploadInfo(opts *bind.CallOpts, uploadOption UploadOption, beginHeight *big.Int) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "CalcProveTimesByUploadInfo", uploadOption, beginHeight)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalcProveTimesByUploadInfo is a free data retrieval call binding the contract method 0xe4bca973.
//
// Solidity: function CalcProveTimesByUploadInfo((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, uint256 beginHeight) pure returns(uint64)
func (_Store *StoreSession) CalcProveTimesByUploadInfo(uploadOption UploadOption, beginHeight *big.Int) (uint64, error) {
	return _Store.Contract.CalcProveTimesByUploadInfo(&_Store.CallOpts, uploadOption, beginHeight)
}

// CalcProveTimesByUploadInfo is a free data retrieval call binding the contract method 0xe4bca973.
//
// Solidity: function CalcProveTimesByUploadInfo((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, uint256 beginHeight) pure returns(uint64)
func (_Store *StoreCallerSession) CalcProveTimesByUploadInfo(uploadOption UploadOption, beginHeight *big.Int) (uint64, error) {
	return _Store.Contract.CalcProveTimesByUploadInfo(&_Store.CallOpts, uploadOption, beginHeight)
}

// GetFileInfo is a free data retrieval call binding the contract method 0xdefce5d4.
//
// Solidity: function GetFileInfo(bytes fileHash) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)))
func (_Store *StoreCaller) GetFileInfo(opts *bind.CallOpts, fileHash []byte) (FileInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetFileInfo", fileHash)

	if err != nil {
		return *new(FileInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(FileInfo)).(*FileInfo)

	return out0, err

}

// GetFileInfo is a free data retrieval call binding the contract method 0xdefce5d4.
//
// Solidity: function GetFileInfo(bytes fileHash) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)))
func (_Store *StoreSession) GetFileInfo(fileHash []byte) (FileInfo, error) {
	return _Store.Contract.GetFileInfo(&_Store.CallOpts, fileHash)
}

// GetFileInfo is a free data retrieval call binding the contract method 0xdefce5d4.
//
// Solidity: function GetFileInfo(bytes fileHash) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)))
func (_Store *StoreCallerSession) GetFileInfo(fileHash []byte) (FileInfo, error) {
	return _Store.Contract.GetFileInfo(&_Store.CallOpts, fileHash)
}

// GetFileInfos is a free data retrieval call binding the contract method 0xc6e8392a.
//
// Solidity: function GetFileInfos(bytes[] _fileList) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[])
func (_Store *StoreCaller) GetFileInfos(opts *bind.CallOpts, _fileList [][]byte) ([]FileInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetFileInfos", _fileList)

	if err != nil {
		return *new([]FileInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]FileInfo)).(*[]FileInfo)

	return out0, err

}

// GetFileInfos is a free data retrieval call binding the contract method 0xc6e8392a.
//
// Solidity: function GetFileInfos(bytes[] _fileList) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[])
func (_Store *StoreSession) GetFileInfos(_fileList [][]byte) ([]FileInfo, error) {
	return _Store.Contract.GetFileInfos(&_Store.CallOpts, _fileList)
}

// GetFileInfos is a free data retrieval call binding the contract method 0xc6e8392a.
//
// Solidity: function GetFileInfos(bytes[] _fileList) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[])
func (_Store *StoreCallerSession) GetFileInfos(_fileList [][]byte) ([]FileInfo, error) {
	return _Store.Contract.GetFileInfos(&_Store.CallOpts, _fileList)
}

// GetFileList is a free data retrieval call binding the contract method 0x9a2b5e63.
//
// Solidity: function GetFileList(address walletAddr) view returns(bytes[])
func (_Store *StoreCaller) GetFileList(opts *bind.CallOpts, walletAddr common.Address) ([][]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetFileList", walletAddr)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetFileList is a free data retrieval call binding the contract method 0x9a2b5e63.
//
// Solidity: function GetFileList(address walletAddr) view returns(bytes[])
func (_Store *StoreSession) GetFileList(walletAddr common.Address) ([][]byte, error) {
	return _Store.Contract.GetFileList(&_Store.CallOpts, walletAddr)
}

// GetFileList is a free data retrieval call binding the contract method 0x9a2b5e63.
//
// Solidity: function GetFileList(address walletAddr) view returns(bytes[])
func (_Store *StoreCallerSession) GetFileList(walletAddr common.Address) ([][]byte, error) {
	return _Store.Contract.GetFileList(&_Store.CallOpts, walletAddr)
}

// GetUnProveCandidateFiles is a free data retrieval call binding the contract method 0xdc1ec30b.
//
// Solidity: function GetUnProveCandidateFiles(address walletAddr) view returns(bytes[])
func (_Store *StoreCaller) GetUnProveCandidateFiles(opts *bind.CallOpts, walletAddr common.Address) ([][]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetUnProveCandidateFiles", walletAddr)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetUnProveCandidateFiles is a free data retrieval call binding the contract method 0xdc1ec30b.
//
// Solidity: function GetUnProveCandidateFiles(address walletAddr) view returns(bytes[])
func (_Store *StoreSession) GetUnProveCandidateFiles(walletAddr common.Address) ([][]byte, error) {
	return _Store.Contract.GetUnProveCandidateFiles(&_Store.CallOpts, walletAddr)
}

// GetUnProveCandidateFiles is a free data retrieval call binding the contract method 0xdc1ec30b.
//
// Solidity: function GetUnProveCandidateFiles(address walletAddr) view returns(bytes[])
func (_Store *StoreCallerSession) GetUnProveCandidateFiles(walletAddr common.Address) ([][]byte, error) {
	return _Store.Contract.GetUnProveCandidateFiles(&_Store.CallOpts, walletAddr)
}

// GetUnProvePrimaryFiles is a free data retrieval call binding the contract method 0x8d41f9f8.
//
// Solidity: function GetUnProvePrimaryFiles(address walletAddr) view returns(bytes[])
func (_Store *StoreCaller) GetUnProvePrimaryFiles(opts *bind.CallOpts, walletAddr common.Address) ([][]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetUnProvePrimaryFiles", walletAddr)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetUnProvePrimaryFiles is a free data retrieval call binding the contract method 0x8d41f9f8.
//
// Solidity: function GetUnProvePrimaryFiles(address walletAddr) view returns(bytes[])
func (_Store *StoreSession) GetUnProvePrimaryFiles(walletAddr common.Address) ([][]byte, error) {
	return _Store.Contract.GetUnProvePrimaryFiles(&_Store.CallOpts, walletAddr)
}

// GetUnProvePrimaryFiles is a free data retrieval call binding the contract method 0x8d41f9f8.
//
// Solidity: function GetUnProvePrimaryFiles(address walletAddr) view returns(bytes[])
func (_Store *StoreCallerSession) GetUnProvePrimaryFiles(walletAddr common.Address) ([][]byte, error) {
	return _Store.Contract.GetUnProvePrimaryFiles(&_Store.CallOpts, walletAddr)
}

// GetUnSettledFileList is a free data retrieval call binding the contract method 0x28a8565c.
//
// Solidity: function GetUnSettledFileList(address walletAddr) view returns(bytes[])
func (_Store *StoreCaller) GetUnSettledFileList(opts *bind.CallOpts, walletAddr common.Address) ([][]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetUnSettledFileList", walletAddr)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetUnSettledFileList is a free data retrieval call binding the contract method 0x28a8565c.
//
// Solidity: function GetUnSettledFileList(address walletAddr) view returns(bytes[])
func (_Store *StoreSession) GetUnSettledFileList(walletAddr common.Address) ([][]byte, error) {
	return _Store.Contract.GetUnSettledFileList(&_Store.CallOpts, walletAddr)
}

// GetUnSettledFileList is a free data retrieval call binding the contract method 0x28a8565c.
//
// Solidity: function GetUnSettledFileList(address walletAddr) view returns(bytes[])
func (_Store *StoreCallerSession) GetUnSettledFileList(walletAddr common.Address) ([][]byte, error) {
	return _Store.Contract.GetUnSettledFileList(&_Store.CallOpts, walletAddr)
}

// GetUploadStorageFee is a free data retrieval call binding the contract method 0x41bc86cb.
//
// Solidity: function GetUploadStorageFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption) view returns((uint64,uint64,uint64))
func (_Store *StoreCaller) GetUploadStorageFee(opts *bind.CallOpts, uploadOption UploadOption) (StorageFee, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetUploadStorageFee", uploadOption)

	if err != nil {
		return *new(StorageFee), err
	}

	out0 := *abi.ConvertType(out[0], new(StorageFee)).(*StorageFee)

	return out0, err

}

// GetUploadStorageFee is a free data retrieval call binding the contract method 0x41bc86cb.
//
// Solidity: function GetUploadStorageFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption) view returns((uint64,uint64,uint64))
func (_Store *StoreSession) GetUploadStorageFee(uploadOption UploadOption) (StorageFee, error) {
	return _Store.Contract.GetUploadStorageFee(&_Store.CallOpts, uploadOption)
}

// GetUploadStorageFee is a free data retrieval call binding the contract method 0x41bc86cb.
//
// Solidity: function GetUploadStorageFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption) view returns((uint64,uint64,uint64))
func (_Store *StoreCallerSession) GetUploadStorageFee(uploadOption UploadOption) (StorageFee, error) {
	return _Store.Contract.GetUploadStorageFee(&_Store.CallOpts, uploadOption)
}

// UpdateFileInfoForRenew is a free data retrieval call binding the contract method 0x740f6913.
//
// Solidity: function UpdateFileInfoForRenew((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) pure returns(bool)
func (_Store *StoreCaller) UpdateFileInfoForRenew(opts *bind.CallOpts, setting Setting, newExpireHeight *big.Int, fileInfo FileInfo) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "UpdateFileInfoForRenew", setting, newExpireHeight, fileInfo)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UpdateFileInfoForRenew is a free data retrieval call binding the contract method 0x740f6913.
//
// Solidity: function UpdateFileInfoForRenew((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) pure returns(bool)
func (_Store *StoreSession) UpdateFileInfoForRenew(setting Setting, newExpireHeight *big.Int, fileInfo FileInfo) (bool, error) {
	return _Store.Contract.UpdateFileInfoForRenew(&_Store.CallOpts, setting, newExpireHeight, fileInfo)
}

// UpdateFileInfoForRenew is a free data retrieval call binding the contract method 0x740f6913.
//
// Solidity: function UpdateFileInfoForRenew((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) pure returns(bool)
func (_Store *StoreCallerSession) UpdateFileInfoForRenew(setting Setting, newExpireHeight *big.Int, fileInfo FileInfo) (bool, error) {
	return _Store.Contract.UpdateFileInfoForRenew(&_Store.CallOpts, setting, newExpireHeight, fileInfo)
}

// UpdateFilesForRenew is a free data retrieval call binding the contract method 0x3f2cc9a0.
//
// Solidity: function UpdateFilesForRenew(bytes[] _fileList, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[], bool)
func (_Store *StoreCaller) UpdateFilesForRenew(opts *bind.CallOpts, _fileList [][]byte, setting Setting, newExpireHeight *big.Int) ([]FileInfo, bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "UpdateFilesForRenew", _fileList, setting, newExpireHeight)

	if err != nil {
		return *new([]FileInfo), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]FileInfo)).(*[]FileInfo)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// UpdateFilesForRenew is a free data retrieval call binding the contract method 0x3f2cc9a0.
//
// Solidity: function UpdateFilesForRenew(bytes[] _fileList, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[], bool)
func (_Store *StoreSession) UpdateFilesForRenew(_fileList [][]byte, setting Setting, newExpireHeight *big.Int) ([]FileInfo, bool, error) {
	return _Store.Contract.UpdateFilesForRenew(&_Store.CallOpts, _fileList, setting, newExpireHeight)
}

// UpdateFilesForRenew is a free data retrieval call binding the contract method 0x3f2cc9a0.
//
// Solidity: function UpdateFilesForRenew(bytes[] _fileList, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[], bool)
func (_Store *StoreCallerSession) UpdateFilesForRenew(_fileList [][]byte, setting Setting, newExpireHeight *big.Int) ([]FileInfo, bool, error) {
	return _Store.Contract.UpdateFilesForRenew(&_Store.CallOpts, _fileList, setting, newExpireHeight)
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

// CalcStorageFee is a free data retrieval call binding the contract method 0x70b1d889.
//
// Solidity: function calcStorageFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 copyNum, uint64 fileSize, uint64 duration) pure returns(uint64)
func (_Store *StoreCaller) CalcStorageFee(opts *bind.CallOpts, setting Setting, copyNum uint64, fileSize uint64, duration uint64) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "calcStorageFee", setting, copyNum, fileSize, duration)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalcStorageFee is a free data retrieval call binding the contract method 0x70b1d889.
//
// Solidity: function calcStorageFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 copyNum, uint64 fileSize, uint64 duration) pure returns(uint64)
func (_Store *StoreSession) CalcStorageFee(setting Setting, copyNum uint64, fileSize uint64, duration uint64) (uint64, error) {
	return _Store.Contract.CalcStorageFee(&_Store.CallOpts, setting, copyNum, fileSize, duration)
}

// CalcStorageFee is a free data retrieval call binding the contract method 0x70b1d889.
//
// Solidity: function calcStorageFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 copyNum, uint64 fileSize, uint64 duration) pure returns(uint64)
func (_Store *StoreCallerSession) CalcStorageFee(setting Setting, copyNum uint64, fileSize uint64, duration uint64) (uint64, error) {
	return _Store.Contract.CalcStorageFee(&_Store.CallOpts, setting, copyNum, fileSize, duration)
}

// CalcStorageFeeForOneNode is a free data retrieval call binding the contract method 0xdf49410a.
//
// Solidity: function calcStorageFeeForOneNode((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 fileSize, uint64 duration) pure returns(uint64)
func (_Store *StoreCaller) CalcStorageFeeForOneNode(opts *bind.CallOpts, setting Setting, fileSize uint64, duration uint64) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "calcStorageFeeForOneNode", setting, fileSize, duration)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalcStorageFeeForOneNode is a free data retrieval call binding the contract method 0xdf49410a.
//
// Solidity: function calcStorageFeeForOneNode((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 fileSize, uint64 duration) pure returns(uint64)
func (_Store *StoreSession) CalcStorageFeeForOneNode(setting Setting, fileSize uint64, duration uint64) (uint64, error) {
	return _Store.Contract.CalcStorageFeeForOneNode(&_Store.CallOpts, setting, fileSize, duration)
}

// CalcStorageFeeForOneNode is a free data retrieval call binding the contract method 0xdf49410a.
//
// Solidity: function calcStorageFeeForOneNode((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 fileSize, uint64 duration) pure returns(uint64)
func (_Store *StoreCallerSession) CalcStorageFeeForOneNode(setting Setting, fileSize uint64, duration uint64) (uint64, error) {
	return _Store.Contract.CalcStorageFeeForOneNode(&_Store.CallOpts, setting, fileSize, duration)
}

// CalcUploadFee is a free data retrieval call binding the contract method 0xd49ce874.
//
// Solidity: function calcUploadFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) pure returns((uint64,uint64,uint64))
func (_Store *StoreCaller) CalcUploadFee(opts *bind.CallOpts, uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "calcUploadFee", uploadOption, setting, currentHeight)

	if err != nil {
		return *new(StorageFee), err
	}

	out0 := *abi.ConvertType(out[0], new(StorageFee)).(*StorageFee)

	return out0, err

}

// CalcUploadFee is a free data retrieval call binding the contract method 0xd49ce874.
//
// Solidity: function calcUploadFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) pure returns((uint64,uint64,uint64))
func (_Store *StoreSession) CalcUploadFee(uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	return _Store.Contract.CalcUploadFee(&_Store.CallOpts, uploadOption, setting, currentHeight)
}

// CalcUploadFee is a free data retrieval call binding the contract method 0xd49ce874.
//
// Solidity: function calcUploadFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) pure returns((uint64,uint64,uint64))
func (_Store *StoreCallerSession) CalcUploadFee(uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	return _Store.Contract.CalcUploadFee(&_Store.CallOpts, uploadOption, setting, currentHeight)
}

// CalcValidFee is a free data retrieval call binding the contract method 0x170bbdf5.
//
// Solidity: function calcValidFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 copyNum, uint64 fileSize) pure returns(uint64)
func (_Store *StoreCaller) CalcValidFee(opts *bind.CallOpts, setting Setting, proveTime uint64, copyNum uint64, fileSize uint64) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "calcValidFee", setting, proveTime, copyNum, fileSize)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalcValidFee is a free data retrieval call binding the contract method 0x170bbdf5.
//
// Solidity: function calcValidFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 copyNum, uint64 fileSize) pure returns(uint64)
func (_Store *StoreSession) CalcValidFee(setting Setting, proveTime uint64, copyNum uint64, fileSize uint64) (uint64, error) {
	return _Store.Contract.CalcValidFee(&_Store.CallOpts, setting, proveTime, copyNum, fileSize)
}

// CalcValidFee is a free data retrieval call binding the contract method 0x170bbdf5.
//
// Solidity: function calcValidFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 copyNum, uint64 fileSize) pure returns(uint64)
func (_Store *StoreCallerSession) CalcValidFee(setting Setting, proveTime uint64, copyNum uint64, fileSize uint64) (uint64, error) {
	return _Store.Contract.CalcValidFee(&_Store.CallOpts, setting, proveTime, copyNum, fileSize)
}

// CalcValidFeeForOneNode is a free data retrieval call binding the contract method 0x793cab8e.
//
// Solidity: function calcValidFeeForOneNode((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 fileSize) pure returns(uint64)
func (_Store *StoreCaller) CalcValidFeeForOneNode(opts *bind.CallOpts, setting Setting, proveTime uint64, fileSize uint64) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "calcValidFeeForOneNode", setting, proveTime, fileSize)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalcValidFeeForOneNode is a free data retrieval call binding the contract method 0x793cab8e.
//
// Solidity: function calcValidFeeForOneNode((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 fileSize) pure returns(uint64)
func (_Store *StoreSession) CalcValidFeeForOneNode(setting Setting, proveTime uint64, fileSize uint64) (uint64, error) {
	return _Store.Contract.CalcValidFeeForOneNode(&_Store.CallOpts, setting, proveTime, fileSize)
}

// CalcValidFeeForOneNode is a free data retrieval call binding the contract method 0x793cab8e.
//
// Solidity: function calcValidFeeForOneNode((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 fileSize) pure returns(uint64)
func (_Store *StoreCallerSession) CalcValidFeeForOneNode(setting Setting, proveTime uint64, fileSize uint64) (uint64, error) {
	return _Store.Contract.CalcValidFeeForOneNode(&_Store.CallOpts, setting, proveTime, fileSize)
}

// AddFileToUnSettleList is a paid mutator transaction binding the contract method 0x34fddaac.
//
// Solidity: function AddFileToUnSettleList(address fileOwner, bytes fileHash) returns()
func (_Store *StoreTransactor) AddFileToUnSettleList(opts *bind.TransactOpts, fileOwner common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "AddFileToUnSettleList", fileOwner, fileHash)
}

// AddFileToUnSettleList is a paid mutator transaction binding the contract method 0x34fddaac.
//
// Solidity: function AddFileToUnSettleList(address fileOwner, bytes fileHash) returns()
func (_Store *StoreSession) AddFileToUnSettleList(fileOwner common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.AddFileToUnSettleList(&_Store.TransactOpts, fileOwner, fileHash)
}

// AddFileToUnSettleList is a paid mutator transaction binding the contract method 0x34fddaac.
//
// Solidity: function AddFileToUnSettleList(address fileOwner, bytes fileHash) returns()
func (_Store *StoreTransactorSession) AddFileToUnSettleList(fileOwner common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.AddFileToUnSettleList(&_Store.TransactOpts, fileOwner, fileHash)
}

// ChangeFileOwner is a paid mutator transaction binding the contract method 0x207e33be.
//
// Solidity: function ChangeFileOwner((bytes,address,address) ownerChange) returns()
func (_Store *StoreTransactor) ChangeFileOwner(opts *bind.TransactOpts, ownerChange FileSystemOwnerChange) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ChangeFileOwner", ownerChange)
}

// ChangeFileOwner is a paid mutator transaction binding the contract method 0x207e33be.
//
// Solidity: function ChangeFileOwner((bytes,address,address) ownerChange) returns()
func (_Store *StoreSession) ChangeFileOwner(ownerChange FileSystemOwnerChange) (*types.Transaction, error) {
	return _Store.Contract.ChangeFileOwner(&_Store.TransactOpts, ownerChange)
}

// ChangeFileOwner is a paid mutator transaction binding the contract method 0x207e33be.
//
// Solidity: function ChangeFileOwner((bytes,address,address) ownerChange) returns()
func (_Store *StoreTransactorSession) ChangeFileOwner(ownerChange FileSystemOwnerChange) (*types.Transaction, error) {
	return _Store.Contract.ChangeFileOwner(&_Store.TransactOpts, ownerChange)
}

// ChangeFilePrivilege is a paid mutator transaction binding the contract method 0xb6af10fb.
//
// Solidity: function ChangeFilePrivilege((bytes,uint64) priChange) returns()
func (_Store *StoreTransactor) ChangeFilePrivilege(opts *bind.TransactOpts, priChange FileSystemPriChange) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ChangeFilePrivilege", priChange)
}

// ChangeFilePrivilege is a paid mutator transaction binding the contract method 0xb6af10fb.
//
// Solidity: function ChangeFilePrivilege((bytes,uint64) priChange) returns()
func (_Store *StoreSession) ChangeFilePrivilege(priChange FileSystemPriChange) (*types.Transaction, error) {
	return _Store.Contract.ChangeFilePrivilege(&_Store.TransactOpts, priChange)
}

// ChangeFilePrivilege is a paid mutator transaction binding the contract method 0xb6af10fb.
//
// Solidity: function ChangeFilePrivilege((bytes,uint64) priChange) returns()
func (_Store *StoreTransactorSession) ChangeFilePrivilege(priChange FileSystemPriChange) (*types.Transaction, error) {
	return _Store.Contract.ChangeFilePrivilege(&_Store.TransactOpts, priChange)
}

// CleanupForDeleteFile is a paid mutator transaction binding the contract method 0x9cd103e9.
//
// Solidity: function CleanupForDeleteFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo, bool rmInfo, bool rmList) returns()
func (_Store *StoreTransactor) CleanupForDeleteFile(opts *bind.TransactOpts, fileInfo FileInfo, rmInfo bool, rmList bool) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "CleanupForDeleteFile", fileInfo, rmInfo, rmList)
}

// CleanupForDeleteFile is a paid mutator transaction binding the contract method 0x9cd103e9.
//
// Solidity: function CleanupForDeleteFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo, bool rmInfo, bool rmList) returns()
func (_Store *StoreSession) CleanupForDeleteFile(fileInfo FileInfo, rmInfo bool, rmList bool) (*types.Transaction, error) {
	return _Store.Contract.CleanupForDeleteFile(&_Store.TransactOpts, fileInfo, rmInfo, rmList)
}

// CleanupForDeleteFile is a paid mutator transaction binding the contract method 0x9cd103e9.
//
// Solidity: function CleanupForDeleteFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo, bool rmInfo, bool rmList) returns()
func (_Store *StoreTransactorSession) CleanupForDeleteFile(fileInfo FileInfo, rmInfo bool, rmList bool) (*types.Transaction, error) {
	return _Store.Contract.CleanupForDeleteFile(&_Store.TransactOpts, fileInfo, rmInfo, rmList)
}

// DelFileFromCandidateList is a paid mutator transaction binding the contract method 0x1bc55880.
//
// Solidity: function DelFileFromCandidateList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreTransactor) DelFileFromCandidateList(opts *bind.TransactOpts, walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DelFileFromCandidateList", walletAddr, fileHash)
}

// DelFileFromCandidateList is a paid mutator transaction binding the contract method 0x1bc55880.
//
// Solidity: function DelFileFromCandidateList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreSession) DelFileFromCandidateList(walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DelFileFromCandidateList(&_Store.TransactOpts, walletAddr, fileHash)
}

// DelFileFromCandidateList is a paid mutator transaction binding the contract method 0x1bc55880.
//
// Solidity: function DelFileFromCandidateList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreTransactorSession) DelFileFromCandidateList(walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DelFileFromCandidateList(&_Store.TransactOpts, walletAddr, fileHash)
}

// DelFileFromList is a paid mutator transaction binding the contract method 0xebaf0258.
//
// Solidity: function DelFileFromList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreTransactor) DelFileFromList(opts *bind.TransactOpts, walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DelFileFromList", walletAddr, fileHash)
}

// DelFileFromList is a paid mutator transaction binding the contract method 0xebaf0258.
//
// Solidity: function DelFileFromList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreSession) DelFileFromList(walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DelFileFromList(&_Store.TransactOpts, walletAddr, fileHash)
}

// DelFileFromList is a paid mutator transaction binding the contract method 0xebaf0258.
//
// Solidity: function DelFileFromList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreTransactorSession) DelFileFromList(walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DelFileFromList(&_Store.TransactOpts, walletAddr, fileHash)
}

// DelFileFromPrimaryList is a paid mutator transaction binding the contract method 0x655c6694.
//
// Solidity: function DelFileFromPrimaryList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreTransactor) DelFileFromPrimaryList(opts *bind.TransactOpts, walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DelFileFromPrimaryList", walletAddr, fileHash)
}

// DelFileFromPrimaryList is a paid mutator transaction binding the contract method 0x655c6694.
//
// Solidity: function DelFileFromPrimaryList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreSession) DelFileFromPrimaryList(walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DelFileFromPrimaryList(&_Store.TransactOpts, walletAddr, fileHash)
}

// DelFileFromPrimaryList is a paid mutator transaction binding the contract method 0x655c6694.
//
// Solidity: function DelFileFromPrimaryList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreTransactorSession) DelFileFromPrimaryList(walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DelFileFromPrimaryList(&_Store.TransactOpts, walletAddr, fileHash)
}

// DelFileFromUnSettledList is a paid mutator transaction binding the contract method 0xb64ab1ef.
//
// Solidity: function DelFileFromUnSettledList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreTransactor) DelFileFromUnSettledList(opts *bind.TransactOpts, walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DelFileFromUnSettledList", walletAddr, fileHash)
}

// DelFileFromUnSettledList is a paid mutator transaction binding the contract method 0xb64ab1ef.
//
// Solidity: function DelFileFromUnSettledList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreSession) DelFileFromUnSettledList(walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DelFileFromUnSettledList(&_Store.TransactOpts, walletAddr, fileHash)
}

// DelFileFromUnSettledList is a paid mutator transaction binding the contract method 0xb64ab1ef.
//
// Solidity: function DelFileFromUnSettledList(address walletAddr, bytes fileHash) returns()
func (_Store *StoreTransactorSession) DelFileFromUnSettledList(walletAddr common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DelFileFromUnSettledList(&_Store.TransactOpts, walletAddr, fileHash)
}

// DeleteExpiredFilesFromList is a paid mutator transaction binding the contract method 0x3ad525a9.
//
// Solidity: function DeleteExpiredFilesFromList(bytes[] _fileList, address walletAddr, uint8[] storageType) returns(bytes[], uint64, bool)
func (_Store *StoreTransactor) DeleteExpiredFilesFromList(opts *bind.TransactOpts, _fileList [][]byte, walletAddr common.Address, storageType []uint8) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteExpiredFilesFromList", _fileList, walletAddr, storageType)
}

// DeleteExpiredFilesFromList is a paid mutator transaction binding the contract method 0x3ad525a9.
//
// Solidity: function DeleteExpiredFilesFromList(bytes[] _fileList, address walletAddr, uint8[] storageType) returns(bytes[], uint64, bool)
func (_Store *StoreSession) DeleteExpiredFilesFromList(_fileList [][]byte, walletAddr common.Address, storageType []uint8) (*types.Transaction, error) {
	return _Store.Contract.DeleteExpiredFilesFromList(&_Store.TransactOpts, _fileList, walletAddr, storageType)
}

// DeleteExpiredFilesFromList is a paid mutator transaction binding the contract method 0x3ad525a9.
//
// Solidity: function DeleteExpiredFilesFromList(bytes[] _fileList, address walletAddr, uint8[] storageType) returns(bytes[], uint64, bool)
func (_Store *StoreTransactorSession) DeleteExpiredFilesFromList(_fileList [][]byte, walletAddr common.Address, storageType []uint8) (*types.Transaction, error) {
	return _Store.Contract.DeleteExpiredFilesFromList(&_Store.TransactOpts, _fileList, walletAddr, storageType)
}

// DeleteFile is a paid mutator transaction binding the contract method 0xce904554.
//
// Solidity: function DeleteFile(bytes fileHash) returns()
func (_Store *StoreTransactor) DeleteFile(opts *bind.TransactOpts, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteFile", fileHash)
}

// DeleteFile is a paid mutator transaction binding the contract method 0xce904554.
//
// Solidity: function DeleteFile(bytes fileHash) returns()
func (_Store *StoreSession) DeleteFile(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteFile(&_Store.TransactOpts, fileHash)
}

// DeleteFile is a paid mutator transaction binding the contract method 0xce904554.
//
// Solidity: function DeleteFile(bytes fileHash) returns()
func (_Store *StoreTransactorSession) DeleteFile(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteFile(&_Store.TransactOpts, fileHash)
}

// DeleteFileInfo is a paid mutator transaction binding the contract method 0x95b0b54b.
//
// Solidity: function DeleteFileInfo(bytes fileHash) returns()
func (_Store *StoreTransactor) DeleteFileInfo(opts *bind.TransactOpts, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteFileInfo", fileHash)
}

// DeleteFileInfo is a paid mutator transaction binding the contract method 0x95b0b54b.
//
// Solidity: function DeleteFileInfo(bytes fileHash) returns()
func (_Store *StoreSession) DeleteFileInfo(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteFileInfo(&_Store.TransactOpts, fileHash)
}

// DeleteFileInfo is a paid mutator transaction binding the contract method 0x95b0b54b.
//
// Solidity: function DeleteFileInfo(bytes fileHash) returns()
func (_Store *StoreTransactorSession) DeleteFileInfo(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteFileInfo(&_Store.TransactOpts, fileHash)
}

// DeleteFiles is a paid mutator transaction binding the contract method 0x432152ce.
//
// Solidity: function DeleteFiles(bytes[] fileHashs) returns()
func (_Store *StoreTransactor) DeleteFiles(opts *bind.TransactOpts, fileHashs [][]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteFiles", fileHashs)
}

// DeleteFiles is a paid mutator transaction binding the contract method 0x432152ce.
//
// Solidity: function DeleteFiles(bytes[] fileHashs) returns()
func (_Store *StoreSession) DeleteFiles(fileHashs [][]byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteFiles(&_Store.TransactOpts, fileHashs)
}

// DeleteFiles is a paid mutator transaction binding the contract method 0x432152ce.
//
// Solidity: function DeleteFiles(bytes[] fileHashs) returns()
func (_Store *StoreTransactorSession) DeleteFiles(fileHashs [][]byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteFiles(&_Store.TransactOpts, fileHashs)
}

// DeleteUnsettledFiles is a paid mutator transaction binding the contract method 0x0be09702.
//
// Solidity: function DeleteUnsettledFiles(address walletAddr) returns()
func (_Store *StoreTransactor) DeleteUnsettledFiles(opts *bind.TransactOpts, walletAddr common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteUnsettledFiles", walletAddr)
}

// DeleteUnsettledFiles is a paid mutator transaction binding the contract method 0x0be09702.
//
// Solidity: function DeleteUnsettledFiles(address walletAddr) returns()
func (_Store *StoreSession) DeleteUnsettledFiles(walletAddr common.Address) (*types.Transaction, error) {
	return _Store.Contract.DeleteUnsettledFiles(&_Store.TransactOpts, walletAddr)
}

// DeleteUnsettledFiles is a paid mutator transaction binding the contract method 0x0be09702.
//
// Solidity: function DeleteUnsettledFiles(address walletAddr) returns()
func (_Store *StoreTransactorSession) DeleteUnsettledFiles(walletAddr common.Address) (*types.Transaction, error) {
	return _Store.Contract.DeleteUnsettledFiles(&_Store.TransactOpts, walletAddr)
}

// FileReNew is a paid mutator transaction binding the contract method 0x57d94399.
//
// Solidity: function FileReNew((bytes,address,uint64) fileReNewInfo) payable returns()
func (_Store *StoreTransactor) FileReNew(opts *bind.TransactOpts, fileReNewInfo FileReNewInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "FileReNew", fileReNewInfo)
}

// FileReNew is a paid mutator transaction binding the contract method 0x57d94399.
//
// Solidity: function FileReNew((bytes,address,uint64) fileReNewInfo) payable returns()
func (_Store *StoreSession) FileReNew(fileReNewInfo FileReNewInfo) (*types.Transaction, error) {
	return _Store.Contract.FileReNew(&_Store.TransactOpts, fileReNewInfo)
}

// FileReNew is a paid mutator transaction binding the contract method 0x57d94399.
//
// Solidity: function FileReNew((bytes,address,uint64) fileReNewInfo) payable returns()
func (_Store *StoreTransactorSession) FileReNew(fileReNewInfo FileReNewInfo) (*types.Transaction, error) {
	return _Store.Contract.FileReNew(&_Store.TransactOpts, fileReNewInfo)
}

// StoreFile is a paid mutator transaction binding the contract method 0xf54cd295.
//
// Solidity: function StoreFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreTransactor) StoreFile(opts *bind.TransactOpts, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "StoreFile", fileInfo)
}

// StoreFile is a paid mutator transaction binding the contract method 0xf54cd295.
//
// Solidity: function StoreFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreSession) StoreFile(fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.StoreFile(&_Store.TransactOpts, fileInfo)
}

// StoreFile is a paid mutator transaction binding the contract method 0xf54cd295.
//
// Solidity: function StoreFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreTransactorSession) StoreFile(fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.StoreFile(&_Store.TransactOpts, fileInfo)
}

// UpdateFileInfo is a paid mutator transaction binding the contract method 0x681850d7.
//
// Solidity: function UpdateFileInfo((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) f) returns()
func (_Store *StoreTransactor) UpdateFileInfo(opts *bind.TransactOpts, f FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateFileInfo", f)
}

// UpdateFileInfo is a paid mutator transaction binding the contract method 0x681850d7.
//
// Solidity: function UpdateFileInfo((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) f) returns()
func (_Store *StoreSession) UpdateFileInfo(f FileInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileInfo(&_Store.TransactOpts, f)
}

// UpdateFileInfo is a paid mutator transaction binding the contract method 0x681850d7.
//
// Solidity: function UpdateFileInfo((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) f) returns()
func (_Store *StoreTransactorSession) UpdateFileInfo(f FileInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileInfo(&_Store.TransactOpts, f)
}

// UpdateFileList is a paid mutator transaction binding the contract method 0xd70e6272.
//
// Solidity: function UpdateFileList(address walletAddr, bytes[] list) returns()
func (_Store *StoreTransactor) UpdateFileList(opts *bind.TransactOpts, walletAddr common.Address, list [][]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateFileList", walletAddr, list)
}

// UpdateFileList is a paid mutator transaction binding the contract method 0xd70e6272.
//
// Solidity: function UpdateFileList(address walletAddr, bytes[] list) returns()
func (_Store *StoreSession) UpdateFileList(walletAddr common.Address, list [][]byte) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileList(&_Store.TransactOpts, walletAddr, list)
}

// UpdateFileList is a paid mutator transaction binding the contract method 0xd70e6272.
//
// Solidity: function UpdateFileList(address walletAddr, bytes[] list) returns()
func (_Store *StoreTransactorSession) UpdateFileList(walletAddr common.Address, list [][]byte) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileList(&_Store.TransactOpts, walletAddr, list)
}

// DeleteFilesInner is a paid mutator transaction binding the contract method 0xc43007e5.
//
// Solidity: function deleteFilesInner((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[] files) returns()
func (_Store *StoreTransactor) DeleteFilesInner(opts *bind.TransactOpts, files []FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "deleteFilesInner", files)
}

// DeleteFilesInner is a paid mutator transaction binding the contract method 0xc43007e5.
//
// Solidity: function deleteFilesInner((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[] files) returns()
func (_Store *StoreSession) DeleteFilesInner(files []FileInfo) (*types.Transaction, error) {
	return _Store.Contract.DeleteFilesInner(&_Store.TransactOpts, files)
}

// DeleteFilesInner is a paid mutator transaction binding the contract method 0xc43007e5.
//
// Solidity: function deleteFilesInner((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[] files) returns()
func (_Store *StoreTransactorSession) DeleteFilesInner(files []FileInfo) (*types.Transaction, error) {
	return _Store.Contract.DeleteFilesInner(&_Store.TransactOpts, files)
}

// DeleteProveDetails is a paid mutator transaction binding the contract method 0x3d4dcb98.
//
// Solidity: function deleteProveDetails(bytes fileHash) returns()
func (_Store *StoreTransactor) DeleteProveDetails(opts *bind.TransactOpts, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "deleteProveDetails", fileHash)
}

// DeleteProveDetails is a paid mutator transaction binding the contract method 0x3d4dcb98.
//
// Solidity: function deleteProveDetails(bytes fileHash) returns()
func (_Store *StoreSession) DeleteProveDetails(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteProveDetails(&_Store.TransactOpts, fileHash)
}

// DeleteProveDetails is a paid mutator transaction binding the contract method 0x3d4dcb98.
//
// Solidity: function deleteProveDetails(bytes fileHash) returns()
func (_Store *StoreTransactorSession) DeleteProveDetails(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteProveDetails(&_Store.TransactOpts, fileHash)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _config, address _node, address _space, address _sector, address _prove) returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts, _config common.Address, _node common.Address, _space common.Address, _sector common.Address, _prove common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize", _config, _node, _space, _sector, _prove)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _config, address _node, address _space, address _sector, address _prove) returns()
func (_Store *StoreSession) Initialize(_config common.Address, _node common.Address, _space common.Address, _sector common.Address, _prove common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _node, _space, _sector, _prove)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _config, address _node, address _space, address _sector, address _prove) returns()
func (_Store *StoreTransactorSession) Initialize(_config common.Address, _node common.Address, _space common.Address, _sector common.Address, _prove common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _node, _space, _sector, _prove)
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
