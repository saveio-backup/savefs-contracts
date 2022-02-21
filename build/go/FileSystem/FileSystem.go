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
	ABI: "[{\"inputs\":[],\"name\":\"DifferenceFileOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"FileNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProfit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"OpError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceWrongExpiredHeight\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"AddFileToUnSettleList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"currentHeight\",\"type\":\"uint256\"}],\"name\":\"CalcDepositFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"proveTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"CalcFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"beginHeight\",\"type\":\"uint256\"}],\"name\":\"CalcProveTimesByUploadInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"CurOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NewOwner\",\"type\":\"address\"}],\"internalType\":\"structFileSystem.OwnerChange\",\"name\":\"ownerChange\",\"type\":\"tuple\"}],\"name\":\"ChangeFileOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"privilege\",\"type\":\"uint64\"}],\"internalType\":\"structFileSystem.PriChange\",\"name\":\"priChange\",\"type\":\"tuple\"}],\"name\":\"ChangeFilePrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"rmInfo\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rmList\",\"type\":\"bool\"}],\"name\":\"CleanupForDeleteFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DelFileFromCandidateList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DelFileFromList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DelFileFromPrimaryList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DelFileFromUnSettledList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"enumStorageType[]\",\"name\":\"storageType\",\"type\":\"uint8[]\"}],\"name\":\"DeleteExpiredFilesFromList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteFileInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"}],\"name\":\"DeleteFiles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteUnsettledFiles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FromAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ReNewTimes\",\"type\":\"uint64\"}],\"internalType\":\"structFileReNewInfo\",\"name\":\"fileReNewInfo\",\"type\":\"tuple\"}],\"name\":\"FileReNew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"GetFileInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"}],\"name\":\"GetFileInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetFileList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnProveCandidateFiles\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnProvePrimaryFiles\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnSettledFileList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"}],\"name\":\"GetUploadStorageFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"StoreFile\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"f\",\"type\":\"tuple\"}],\"name\":\"UpdateFileInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"newExpireHeight\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"UpdateFileInfoForRenew\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"list\",\"type\":\"bytes[]\"}],\"name\":\"UpdateFileList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"newExpireHeight\",\"type\":\"uint256\"}],\"name\":\"UpdateFilesForRenew\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0xf932268c\",\"type\":\"bytes32\"}],\"name\":\"c_0xf932268c\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"}],\"name\":\"calcSingleValidFeeForFile\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"calcStorageFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"calcStorageFeeForOneNode\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"currentHeight\",\"type\":\"uint256\"}],\"name\":\"calcUploadFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"proveTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"}],\"name\":\"calcValidFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"proveTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"}],\"name\":\"calcValidFeeForOneNode\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"files\",\"type\":\"tuple[]\"}],\"name\":\"deleteFilesInner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"deleteProveDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractNode\",\"name\":\"_node\",\"type\":\"address\"},{\"internalType\":\"contractSpace\",\"name\":\"_space\",\"type\":\"address\"},{\"internalType\":\"contractSector\",\"name\":\"_sector\",\"type\":\"address\"},{\"internalType\":\"contractProve\",\"name\":\"_prove\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526005600460146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550600460149054906101000a900467ffffffffffffffff1662015180620000559190620000bd565b600560006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550620f4240600560086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550348015620000b657600080fd5b5062000138565b6000620000ca82620000f5565b9150620000d783620000f5565b925082620000ea57620000e962000109565b5b828204905092915050565b600067ffffffffffffffff82169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6201445880620001496000396000f3fe608060405260043610620002e45760003560e01c806370b1d8891162000193578063cc76e80d11620000df578063defce5d4116200008b578063e4bca973116200006d578063e4bca9731462000abe578063eb7336a01462000b02578063ebaf02581462000b4657620002e4565b8063defce5d41462000a36578063df49410a1462000a7a57620002e4565b8063d49ce87411620000c1578063d49ce8741462000980578063d70e627214620009c4578063dc1ec30b14620009f257620002e4565b8063cc76e80d146200090e578063ce904554146200095257620002e4565b8063af6a8370116200013f578063b6af10fb1162000121578063b6af10fb1462000858578063c6e8392a1462000886578063ca6142cb14620008ca57620002e4565b8063af6a8370146200080a578063b64ab1ef146200082a57620002e4565b80638d41f9f811620001755780638d41f9f8146200075457806395b0b54b14620007985780639a2b5e6314620007c657620002e4565b806370b1d88914620006cc578063793cab8e146200071057620002e4565b806328a8565c11620002535780633f2cc9a011620001ff578063432152ce11620001e1578063432152ce146200065057806357d94399146200067e578063655c6694146200069e57620002e4565b80633f2cc9a014620005c757806341bc86cb146200060c57620002e4565b806334fddaac116200023557806334fddaac14620005255780633ad525a914620005535780633d4dcb98146200059957620002e4565b806328a8565c14620004b357806334afd8c114620004f757620002e4565b8063178e4dc011620002b35780631bc5588011620002955780631bc5588014620004295780631fe07caa1462000457578063207e33be146200048557620002e4565b8063178e4dc014620003b75780631aa05c5c14620003fb57620002e4565b80630be0970214620002e95780630df7e99e14620003175780631459457a1462000345578063170bbdf51462000373575b600080fd5b348015620002f657600080fd5b506200031560048036038101906200030f919062011d2a565b62000b74565b005b3480156200032457600080fd5b506200034360048036038101906200033d919062011fd2565b620015a0565b005b3480156200035257600080fd5b506200037160048036038101906200036b919062012043565b620015a3565b005b3480156200038057600080fd5b506200039f6004803603810190620003999190620123f4565b620019d4565b604051620003ae9190620138bb565b60405180910390f35b348015620003c457600080fd5b50620003e36004803603810190620003dd919062012530565b62001a92565b604051620003f2919062013844565b60405180910390f35b3480156200040857600080fd5b50620004276004803603810190620004219190620120c5565b62001c26565b005b3480156200043657600080fd5b506200045560048036038101906200044f919062011db0565b6200210c565b005b3480156200046457600080fd5b506200048360048036038101906200047d919062011f48565b62002451565b005b3480156200049257600080fd5b50620004b16004803603810190620004ab9190620121ea565b6200442e565b005b348015620004c057600080fd5b50620004df6004803603810190620004d9919062011d2a565b620046ff565b604051620004ee9190620133d9565b60405180910390f35b3480156200050457600080fd5b506200052360048036038101906200051d91906201210a565b620048ab565b005b3480156200053257600080fd5b506200055160048036038101906200054b919062011db0565b62004edb565b005b3480156200056057600080fd5b506200057f600480360381019062000579919062011e4f565b62004fe4565b6040516200059093929190620133fd565b60405180910390f35b348015620005a657600080fd5b50620005c56004803603810190620005bf919062011ffe565b62005d2c565b005b348015620005d457600080fd5b50620005f36004803603810190620005ed919062011ed7565b620060c7565b6040516200060392919062013465565b60405180910390f35b3480156200061957600080fd5b50620006386004803603810190620006329190620124eb565b620066f2565b60405162000647919062013844565b60405180910390f35b3480156200065d57600080fd5b506200067c600480360381019062000676919062011e0a565b620069ab565b005b6200069c600480360381019062000696919062012179565b62006e51565b005b348015620006ab57600080fd5b50620006ca6004803603810190620006c4919062011db0565b620083cf565b005b348015620006d957600080fd5b50620006f86004803603810190620006f29190620123f4565b62008714565b604051620007079190620138bb565b60405180910390f35b3480156200071d57600080fd5b506200073c60048036038101906200073691906201239b565b620087d2565b6040516200074b9190620138bb565b60405180910390f35b3480156200076157600080fd5b506200078060048036038101906200077a919062011d2a565b62008880565b6040516200078f9190620133d9565b60405180910390f35b348015620007a557600080fd5b50620007c46004803603810190620007be919062011ffe565b62009263565b005b348015620007d357600080fd5b50620007f26004803603810190620007ec919062011d2a565b62009511565b604051620008019190620133d9565b60405180910390f35b620008286004803603810190620008229190620120c5565b620096bd565b005b3480156200083757600080fd5b5062000856600480360381019062000850919062011db0565b6200bc09565b005b3480156200086557600080fd5b506200088460048036038101906200087e91906201222f565b6200bf4d565b005b3480156200089357600080fd5b50620008b26004803603810190620008ac919062011e0a565b6200c0d5565b604051620008c1919062013441565b60405180910390f35b348015620008d757600080fd5b50620008f66004803603810190620008f0919062012358565b6200d512565b604051620009059190620138bb565b60405180910390f35b3480156200091b57600080fd5b506200093a600480360381019062000934919062012464565b6200d5c8565b60405162000949919062013844565b60405180910390f35b3480156200095f57600080fd5b506200097e600480360381019062000978919062011ffe565b6200d8a5565b005b3480156200098d57600080fd5b50620009ac6004803603810190620009a6919062012530565b6200dbcd565b604051620009bb919062013844565b60405180910390f35b348015620009d157600080fd5b50620009f06004803603810190620009ea919062011d56565b6200e2cb565b005b348015620009ff57600080fd5b5062000a1e600480360381019062000a18919062011d2a565b6200e3af565b60405162000a2d9190620133d9565b60405180910390f35b34801562000a4357600080fd5b5062000a62600480360381019062000a5c919062011ffe565b6200ed92565b60405162000a719190620137ab565b60405180910390f35b34801562000a8757600080fd5b5062000aa6600480360381019062000aa091906201239b565b6200f858565b60405162000ab59190620138bb565b60405180910390f35b34801562000acb57600080fd5b5062000aea600480360381019062000ae49190620125a1565b6200f91c565b60405162000af99190620138bb565b60405180910390f35b34801562000b0f57600080fd5b5062000b2e600480360381019062000b289190620122e6565b6200f9e1565b60405162000b3d919062013499565b60405180910390f35b34801562000b5357600080fd5b5062000b72600480360381019062000b6c919062011db0565b620100e8565b005b62000ba27fbfe1a6994139c5c3920edf29161b82f022ff62ad1ab4a250834c06835a8c344660001b620015a0565b62000bd07f7bcc5fdd9ec0580c1722f61e75e2a62b10d0d71770cc7324436aca6efc22fcf960001b620015a0565b62000bfe7f9105a3a8553720748ef3a0cfbdebd7109dcce51d78bf89ea9c327684fc2ed7de60001b620015a0565b600062000c0b82620046ff565b905062000c3b7ff646fd7c0d38b20681bf0586134cb86672a277fc328eeefc556e6efbe95f3d4660001b620015a0565b62000c697f0cc778ea07ee202fd18b2705852842da3dcb7cbf439e4e0a821f9611801c962760001b620015a0565b6000600267ffffffffffffffff81111562000cad577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801562000cdc5781602001602082028036833780820191505090505b50905062000d0d7fb64fcc500acc72e76b1848a0ad077381b819865ffa83369ffd161f552e2fd9ed60001b620015a0565b62000d3b7fe4a62461237b44e16869ffc69797af14a272a13baf3eb2b412dc2fbfbd66000160001b620015a0565b60008160008151811062000d78577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010190600181111562000db9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9081600181111562000df4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152505062000e267fca4edae81da8eeb08c4484c02b1e6151deda998a354857fe9f91aaf8c5e6af4760001b620015a0565b62000e547f870c243e274a3db17179c525c3b7ab115dc1bfc3fdb42fccb21becc1a1a0b24260001b620015a0565b60018160018151811062000e91577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010190600181111562000ed2577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9081600181111562000f0d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152505062000f3f7fe6a3c0fbfe2be967f37f499c0f2a74de4f231e11420b3f5f89edf9c4a509ed4160001b620015a0565b62000f6d7faf522b3197196cec4d0743bc6045edcd673f31563b41276ae1621a9fd4709c3660001b620015a0565b606062000f9d7fc166f1f38d83c8c12c65ee27a04d3fc0eac5da2dbc69a041994f27300ac4982160001b620015a0565b62000fcb7f7cf81d927fcb489089907f5bdeac3fd9368422e60270e67b1d2264dfb82b036960001b620015a0565b600062000ffb7fab403076d8474c4a1221c1f940f84b0242d8b058bc2377442ae9685e1fd15e7660001b620015a0565b620010297f113b3da8e79fe1b8a48119013ef02b1a60b6d2fac8b2e72110150f91f01d5a6a60001b620015a0565b6000620010597f6eecab00b91fc42f7567b5e0307da3bcdc5050b8c965d23e0835a7468a82049860001b620015a0565b620010877fa031341cdb61fe16eda0b3bebe42c7df9608ce0dea049b26776b05ef262131fd60001b620015a0565b6200109485878662004fe4565b809350819450829550505050620010ce7f704981f986bd5173be4afa36ce2a703df7c8a2d4604ec9189ca46c87ad89f8f360001b620015a0565b620010fc7f3c339fa14726e804890f4753bfa51d6a5270383cab3195999a7e6734696d792a60001b620015a0565b60005b83518110156200159757620011377fb65e33cda92155c40a62d362b87bb240c3c13ebdb57c436fae96ee5c9f05e58860001b620015a0565b620011657f538bd8a7ec958bc898413e0a64adcf2e4f29ca801d67134bf01edfe2bc1b21b860001b620015a0565b6000600b60008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156200127c578382906000526020600020018054620011e89062013ee5565b80601f0160208091040260200160405190810160405280929190818152602001828054620012169062013ee5565b8015620012675780601f106200123b5761010080835404028352916020019162001267565b820191906000526020600020905b8154815290600101906020018083116200124957829003601f168201915b505050505081526020019060010190620011c6565b505050509050620012b07fa74cd022319f44cef0d19be54826a544d4e719ac292f6dbf6d094c893a8ae8a160001b620015a0565b620012de7f406b247f893c76c384514430b2940b5d2419f279946863b5e9a608aae14ded0460001b620015a0565b60005b8151811015620014cd57620013197f290c43f8188491d9b7e9453926d8bd30aeefab5fc523e2ab141fda0b84535cf960001b620015a0565b620013477fafeb82dcf7ab32416f024cc84323c42a611caa1dba81732c52233ef9d9f53e4260001b620015a0565b85838151811062001381577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015180519060200120828281518110620013ca577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015180519060200120141562001488576200140e7f8cc5a2bac7d4765f8118076b9184c3b7dbc258e8641ec0cc8da292147306276960001b620015a0565b6200143c7f53e7b591012f2b40706d30807bf14750a5c09698ba4a9eee97f34cc88c6342c660001b620015a0565b81818151811062001476577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001016060815250620014b7565b620014b67f7f36442dc390b58f719b50e43f6e33bb9b69535074d08077f75e7b1db8ba4eb460001b620015a0565b5b8080620014c49062013f51565b915050620012e1565b50620014fc7f77b74c4251ee86c46c6e36a82c45754d953e260a7106636b6f284faf49cd0aa260001b620015a0565b6200152a7fadd58cd4ed0661c490feb138b9ae2842f2a03017e9b4535bf309ab7876aaf50e60001b620015a0565b80600b60008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090805190602001906200157f92919062010453565b505080806200158e9062013f51565b915050620010ff565b50505050505050565b50565b600060019054906101000a900460ff16620015cd5760008054906101000a900460ff1615620015d8565b620015d76201042d565b5b6200161a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200161190620136df565b60405180910390fd5b60008060019054906101000a900460ff1615905080156200166b576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b620016997f24c038856a123a351ec5c98bb117cb09149daa55e640031b1d8b1358e9f4886a60001b620015a0565b620016c77fa97789f41db57f357f3488ebfc68d20349da283e89335fc1dae994fcecaaabbd60001b620015a0565b620016f57f167acaa778b52b0c023c82d706a1695195011dc22ddefda802d6bf2cb33212b660001b620015a0565b85600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620017647f5bc409bf112e50c0b8106dcfb85c8768859033ffca97c18d78ca6d0966273a8b60001b620015a0565b620017927fcd544d73c3ff60a21e9c799bf6a160189eab5aadd68094969cf2f9e33dbe729460001b620015a0565b84600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620018017fd921e2e99ea30ead8738c6fd252327de7856ba808e8226704f367a512d439b2560001b620015a0565b6200182f7fc52f52cda96c6738011ea1e110849683642e743dfc9c061054569bbd3326459e60001b620015a0565b83600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506200189e7fc1d97d74fe6595ec7ba319d5afcdeb58beefd65fc1591ddc83ef18e054516a8960001b620015a0565b620018cc7f22cd66f2d2611ad157d487a561d1e87ab1d61f7f6b81d8cedd896c27ddd4f9cd60001b620015a0565b82600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506200193b7f14c0a98e3670b109cff339d59fb3b486fced2394053e64ae3b25fe9c247ba72f60001b620015a0565b620019697f890a7f6f3711fdc59755177d27724d66706a2c2412a841d3d4a1d1a1a8313ded60001b620015a0565b81600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015620019cc5760008060016101000a81548160ff0219169083151502179055505b505050505050565b600062001a047f880a161c3ba002f98117f1ef5f63d73872c6609c1f2cfd9c883e74b35826496c60001b620015a0565b62001a327fa255701f88816329d852a4ea9749c203c3db798d1ce2879aa8e8b22406a819d260001b620015a0565b62001a607fb313296457da02f8adaa7bfc56b2a35fef7516fc5616fa54e4bf991244c6336560001b620015a0565b62001a6d858584620087d2565b60018462001a7c919062013bf5565b62001a88919062013c72565b9050949350505050565b62001a9c620104ba565b62001aca7f4e91a6681efa3db9ad639967a2c82936ee06769fd7e9ab2b90cc2ae33c1f8a2960001b620015a0565b62001af87f9f208eafe9c9a21956a1236aff1641893c24f7725cd4984fd51ef87c96c30e6f60001b620015a0565b62001b267f37b8ec15a0ad47b0ecc19089c7489787c113552898ad84af2e7444579f78986160001b620015a0565b600062001b3485846200f91c565b905062001b647fbbb16efbcee538bab3f0d880f6510d16da0d619eb024dcf2a775e64b115038e060001b620015a0565b62001b927f017530f9276699c0934f17fe346e81e82c694264ccaaedf81802cd3842ed5d0560001b620015a0565b600062001bbc85838860c001518960200151888b6080015162001bb6919062013cbb565b6200d5c8565b905062001bec7fa7b766ec203c491458a4d371f22ded7db42f460eccc6b776230b5688be25cf9e60001b620015a0565b62001c1a7f280190ba76b3f309f97c7c32c6dabd1cd81be7db942c38d0020fa4f92e533c7860001b620015a0565b80925050509392505050565b62001c547f8c60b50f8cb941e11ef012eac6e238c8449dcdee55ee332b37549c53db1c2f5f60001b620015a0565b62001c827fda3bc62f49fc366348d1100ba32e8b2e88684b40624f7332de0565f35eb614ca60001b620015a0565b62001cb07f3fd428942c87989e6a3472b7764e3ddc97e2afa1d79c9a744f36a16245c3de1b60001b620015a0565b806006826000015160405162001cc791906201335d565b9081526020016040518091039020600082015181600001908051906020019062001cf3929190620104f9565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201908051906020019062001d59929190620104f9565b5060608201518160030160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160030160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160030160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160030160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060e08201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061010082015181600501556101208201518160060160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101408201518160060160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061016082015181600701908051906020019062001ecf929190620104f9565b506101808201518160080160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101a082015181600901556101c082015181600a0160006101000a81548160ff0219169083151502179055506101e082015181600a0160016101000a81548160ff0219169083600181111562001f7f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555061020082015181600a0160026101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061022082015181600b01908051906020019062001fd39291906201058a565b5061024082015181600c01908051906020019062001ff39291906201058a565b5061026082015181600d0160006101000a81548160ff0219169083600281111562002047577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555061028082015181600d0160016101000a81548160ff0219169083151502179055506102a082015181600e0160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505090505050565b6200213a7f969a90bf79c018a79a82dbc0ac5a7f992106179602efb699e16f46e212fd437360001b620015a0565b620021687f264b9c954684728d32674a7eb6775add718ae7e5e4a4d20b89606f63e298bf9e60001b620015a0565b620021967f2c5c929dfdac5df60c00ca60e13514c29bfa78beb04120499c8bddba12124d7760001b620015a0565b60005b600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490508110156200244c57620022127f68de6f91f5ee1a8079c782d29c33d90ccd9523231b758f1bc1ecb8e80480ed4f60001b620015a0565b620022407f9e56d2be8696791257ffed1c089b4a8a0e127e009b61868467ef3fdf1f64319260001b620015a0565b8180519060200120600a60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208281548110620022c0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002001604051620022d9919062013376565b604051809103902014156200240857620023167fa1781f850664cac5b11a6b9e481dfb02edcae89096e6e0508a9208112410c33b60001b620015a0565b620023447fe1dfee9a86c62f3285b21d6c5ad6f99cc02bc819ddfcc5cd9c21a3dfb2bce81160001b620015a0565b600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208181548110620023bc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020016000620023d4919062010619565b620024027f5b3457f4e6bf0c5d630cce97a025036cdc2c0d4ce500b54d52859f1d6dc997a560001b620015a0565b6200244c565b620024367f50a091d1db902d839c2b19bf3770bda484c7b93f0b12c36a3285a8fff4437ac760001b620015a0565b8080620024439062013f51565b91505062002199565b505050565b6200247f7ffec1e0c92f64247167617933cf4a94ca73cbf5b53834d04a45702275e5b0871360001b620015a0565b620024ad7f5c6f83f77a69331b91f32995d0bf3d5251f62018371d1ba82f515d8d675939a460001b620015a0565b620024db7f7edafe165a55d8bf1b374f81e56317170cedfd8a504706dbad1c13306b7ef8e360001b620015a0565b6000815114156200257657620025147f5b742bf846744189dcb20662dbe034cf3a8d728e1bebc1aa4635b3818d86cacd60001b620015a0565b620025427fb2dc7b267e1f7052516ddd142ad721c4243b7d5f1de11b4630048c403649b20860001b620015a0565b620025707f0d5b81b5145e3b5cf0435caa8a7baca595268cf341f6e948ce5d4f9d206b662e60001b620015a0565b6200442b565b620025a47fa07972223aac44602f2ede0469abb68f6ead68e45749438ca80ab54634e2ab1660001b620015a0565b620025d27f1536f27a76046b679e62b09682fd42873a57b3ec725392b790d0aad981b9a6fe60001b620015a0565b620026007f8abf294a065c517380985b36111f22fe635ace0e880f4f661c3e212861c4fae160001b620015a0565b6000620026307f9f2ef5f0877c3583bcac870c298bd2f0e145eea4cb257b67f75bc6da6678006b60001b620015a0565b6200265e7f77f7c945626170bcad0de9afe11f86c83f8f3c6a2fee463d001f7fcab134fda360001b620015a0565b6000826000815181106200269b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151602001519050620026d77fcaf1dec1d4ab98ea29ce77c4ae043c39e777fda52a8c3bf0d1a57e827788766260001b620015a0565b620027057f02500afe9f3ad8d5c3c4b6209a4d49c4100aa07cdccac6fbfe5d1a2e6596a2e860001b620015a0565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156200277057600080fd5b505afa15801562002785573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620027ab9190620122b9565b9050620027db7fbe491b7a7084baf0e8f3fd530436ec5f73c6aa86e0d755e67ee6d9f44dd3dfd360001b620015a0565b620028097feb0fdb6bfacd32e721fe398f7b50939d1bc9d7bc0b39509df95a44edf061d64160001b620015a0565b60005b845181101562002a1f57620028447f833e2ba4c35cc0d040958d790b057665c8ecfd7531557b1e0ff14c4c107a177c60001b620015a0565b620028727fc8b18c248d4b62c6dd2bfb55daafb0f55d643a91181c89abfc4aa054ddccd26060001b620015a0565b6000858281518110620028ae577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101519050620028e67f596fe214bc38b0aaa0c3efd2f9597f9d4c029b7e56047111d9d9ce4b0274983b60001b620015a0565b620029147fd798db5a967634c6fc6ab9ee01f8352a103d773df45b89a93bc2858ba70af5b560001b620015a0565b8373ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614620029da576200297a7f722dfd020b96eeda4d628c96c4b94069620005784e5417b5851ef0eb53e00d7460001b620015a0565b620029a87f3bd61223627eb34fae9620dfdd0947aeb721f184dd7347b4a4b9b7605291e07a60001b620015a0565b6040517f2f7fecf500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62002a087fe4323601c35961f2934c5742906856954f7f1b292fad95fe5e7b35ec06c2539760001b620015a0565b50808062002a169062013f51565b9150506200280c565b5062002a4e7f5368e6403ccbf59553954750121332161cd900abc20ebc68bf7f282756701d8d60001b620015a0565b62002a7c7f8d1fcacfdfee8d3dad1f0988fb749df24eeeae47125dc2450a00258268da98f960001b620015a0565b60005b8451811015620042475762002ab77f7582c1f7d7cb8ef97095497fa459be80917aa3027590111019662f86a16d387660001b620015a0565b62002ae57f2a9e8e7dc8d63d7c23a030724c07bbdf2e1e85c49b070d15dccd15266b5b3ac460001b620015a0565b600085828151811062002b21577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151905062002b597f42743ef157b72ec28f1c725160ce66abf49d43b22267404dd90227a1d7c76fb460001b620015a0565b62002b877f26d8c049d8338f867bf194bd1530f18e59807f6a63444dc832b0130a6ebf4b0760001b620015a0565b60006007826000015160405162002b9f91906201335d565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b8282101562002c85578382906000526020600020016040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250508152602001906001019062002bcd565b50505050905062002cb97ff494f922483f335cfaecd0760b6928a97501002ac99c80f756704e0e7f3095a460001b620015a0565b62002ce77f50496bc33f510b3fb9d2750123a6fcf7550fb0cf37bd2c044a139f55f8778cfe60001b620015a0565b60005b815181101562002f515762002d227f63c930a7cf17a38b073c96d64be7e3a7d49514103c9506208f2bb5137d49afc760001b620015a0565b62002d507f6e38fe8aca7e8bcb1d5536358ed018563172e75ecb575372a4fd0a5cd4954c2e60001b620015a0565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d784848151811062002dca577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b815260040162002df0919062013827565b60006040518083038186803b15801562002e0957600080fd5b505afa15801562002e1e573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019062002e49919062012274565b905062002e797fc9ff4220484d3f9458a0d2e3b462a02f90036b0fa557b8e0bd54f58cbcc925b560001b620015a0565b62002ea77fa568c33573a18ff02b069e40d62f69c830bdb31120b07120b751ef32826c3ebb60001b620015a0565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a3f0291882866040518363ffffffff1660e01b815260040162002f06929190620137ec565b600060405180830381600087803b15801562002f2157600080fd5b505af115801562002f36573d6000803e3d6000fd5b5050505050808062002f489062013f51565b91505062002cea565b5062002f807f46a5877324f3c7c7bb5f268a888d9421406e799fd8e8c9c94ff6da06e53245dc60001b620015a0565b62002fae7f2051d4eb957f2a75ad63acba42ae5da2926d1c05e7895ab94db6358bbf65883460001b620015a0565b600082610140015167ffffffffffffffff161415620030955762002ff57f9161c85b0f093c4bd1c157308e36dc2d38b5484542091b21fc63c27fa9295f7b60001b620015a0565b620030237f3e3a6be7ba1cb650e05c00bd4d4b10d76de373c5131abe671f56e6bbe2fa5cc060001b620015a0565b620030517fbadb00f25ee0bb872631c63a38cfb86c8b1aa13ce8104dc1a45c90dcac98718660001b620015a0565b6200305f82600180620048ab565b6200308d7f5a1122b1559a58cc549474032921bd01c7f8fb6555d11e12a6ef5b78d90b958560001b620015a0565b505062004231565b620030c37fc10acb0d43a791642286c564333f99bea4a316c6be59468fad2b97f1858add7f60001b620015a0565b620030f17f5b91100d35730d278baea1728df98ad9e365875f6da418c8f7ab092a6334a96760001b620015a0565b6200311f7f4cbf12a30e116d25102eac7fddbb8653c968ba354501a37147c593fbdc7246d760001b620015a0565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663977fdfe284600001516040518263ffffffff1660e01b8152600401620031829190620134b6565b60006040518083038186803b1580156200319b57600080fd5b505afa158015620031b0573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190620031db919062011f8d565b90506200320b7ffff9218c5caa959e92bf145df8d38d706bd35cd3fb9d2585604445e60164923260001b620015a0565b620032397f33d375f9d42d7c701b3bf432bd06f131c29e994a68b158f2834d631a0ac458dd60001b620015a0565b60008361014001519050620032717f040afd990879fcdc42ee17775af451ddfdce4af252de00bba93191b9be85db9960001b620015a0565b6200329f7fe6b1a7f1e1c3c164dab0fafec13071d7e6645a8704751cff4a9e392cef4e483860001b620015a0565b60008460a001518560800151620032b7919062013c72565b9050620032e77ffe31150f992673f9d3d8379febc9cda01f9d5419f1d5bd911a2aae867d6dd5af60001b620015a0565b620033157f8d6e19b5c22b49b24cdc6f2dbffaa98566f654c03a7227e523d3e6ca4ba3257d60001b620015a0565b60006200332388836200d512565b9050620033537f3fda9983dcc516dc5352523d698059269843831a8d3f69867085ea802b850e4860001b620015a0565b620033817f4e971878e934e38074b23682d9a1de4183b86e8eee0ab9295a67f66349ac1dc760001b620015a0565b60005b8451811015620039c757620033bc7f7c1959527d9fc9c7e0255520d47c16d9585faa3b91d61deed26eeb00090a46d460001b620015a0565b620033ea7f6cf24c1ef949a958d9c8493bd221d5233402e6b30be41f6619f84941892b8efb60001b620015a0565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a65374a87848151811062003464577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151602001516040518263ffffffff1660e01b81526004016200348e91906201338f565b60e06040518083038186803b158015620034a757600080fd5b505afa158015620034bc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620034e29190620121be565b9050620035127ff2d4470e3730069059ff1fbe4f49d8f5edd0a81d25911434817b67ea873ec9fa60001b620015a0565b620035407fb6430f90b728ec97641f542a39251f29e419ea8fa8c823bdddeceac1b416caa960001b620015a0565b60008360018885815181106200357f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040015162003597919062013cf6565b620035a3919062013c72565b9050620035d37f3ea4dfe0e0ee29c668b945e46d6ee0420bb6e25d82001cbe93ca5d20876d0e3d60001b620015a0565b620036017fbb67ed0bf8a63594c5d0ee3e5a5ac7745a4f693faae1e266bd66048bc03aa57860001b620015a0565b6000620036228c878c6101a00151436200361c919062013cbb565b6200f858565b9050620036527fdf28b9eb13a43e9f17ea81aa346906479ea69d844f9c870d7176c963fcc98a2960001b620015a0565b620036807f97cf79d2a257e2bc91a3900cc9b35124ab0d9384a1131ff3b8ba68afaed93e8460001b620015a0565b6000818362003690919062013bf5565b9050620036c07faeb2e81a509bbc1eaeaf899d7a1bad6e389fbb49da7745493ef49e4c1c32219d60001b620015a0565b620036ee7fae97696aa73793f3c8d819b6ebf5109eea208256dd7fa3ddb496bcab661c19d460001b620015a0565b8767ffffffffffffffff168167ffffffffffffffff1611156200379957620037397f0bb1ded3dc58a436ec95237250838c3d7a70bda698712f2275c642a3c65e51ab60001b620015a0565b620037677fb1b55a6c00587969576bc73cc96df668bd3505e9ad6075eed6d7b3fb4208e6b260001b620015a0565b6040517f870303a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620037c77ff4ed5bf7270322275c120c565e8560a4df1b77963b12d5e892767a9410aaf0bd60001b620015a0565b620037f57f337f5aa028d412380996bae7990084095129a3908b924255d6472470ee04299d60001b620015a0565b620038237f395c7dd84494560cf7074d5b39fcb620177637ac6d1c4bff9503d4209ccd97fd60001b620015a0565b808460200181815162003837919062013bf5565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050620038827f2a1f3ac5aeedbae936eeecb997b3135e9f7f66e04ba0e700a62d8872f44f913660001b620015a0565b620038b07f1aeeb9a1b11480db21e3a54789d9a1c8164f6449bc197e5f3bf6bccdd439328c60001b620015a0565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c819d86c856040518263ffffffff1660e01b81526004016200390d9190620137cf565b600060405180830381600087803b1580156200392857600080fd5b505af11580156200393d573d6000803e3d6000fd5b505050506200396f7fbb08638ef09e7e6b5afe90409c5d1042ba32ed2447eab02e76b73b0382e914a660001b620015a0565b6200399d7f766475e6d5edfa4ad466cd1f2ea945d4175e3cfc3683acae4a2f654b8cf000c060001b620015a0565b8088620039ab919062013cf6565b9750505050508080620039be9062013f51565b91505062003384565b50620039f67fe7e679d8de5580b377c52cbb98f9b612be10b6e3e49d1f7d7ccff30a53ce323060001b620015a0565b62003a247fdc77753cb50f0510cd4df0ca43b4c3667e7c3aaefaeff0bf42decf3468eec63a60001b620015a0565b6000600181111562003a5f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b866101e00151600181111562003a9e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b141562003b455762003ad37fa710508cc8fd98abba4e1e3a3d43eae2fb4b463eb4960af11e4abff42be86b7460001b620015a0565b62003b017f7498183a7af821a54bfdb8a0047c3d10e1fc80d68a3e8d4242e35192dfe387b160001b620015a0565b62003b2f7ffc3936561907adabb910231832e0d57f298ad160444e092724e614f4e698210e60001b620015a0565b828a62003b3d919062013bf5565b9950620041c0565b62003b737f593c183963387acfdd09f0afbf837bb83e20b4d00fd0d072d3c7d12552952ecc60001b620015a0565b62003ba17f227032a4d1df4bbd0868f6ab7e5fd77efc850eb45f43c8f1fc4d67ec84341d6860001b620015a0565b60018081111562003bdb577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b866101e00151600181111562003c1a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415620041905762003c4f7f1448b95e8691d4d6076281a0c8b43d3b63cb6e278f91d9d19a99a7185b14485560001b620015a0565b62003c7d7f97a7415bde6afff03e4dce77fbbd01ba90c27116097d00fe081a8b58e1652c3560001b620015a0565b62003cab7fd2d92f5c985868675c9c052e8c1cb29692eb01a00922196338a937592bad03b360001b620015a0565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166354e0d3c288602001516040518263ffffffff1660e01b815260040162003d0e91906201338f565b60a06040518083038186803b15801562003d2757600080fd5b505afa15801562003d3c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062003d629190620125fb565b905062003d927f6635b6c18b61ad05ff6085e05bdeb72c85f5617bfe0b5e39ecbd1f12d18d25f560001b620015a0565b62003dc07ff6fb0adeb720b9b7adb754dc2bd3cf5e394ea270ce25102c2c29628c6acf167b60001b620015a0565b8660a00151876080015162003dd6919062013c72565b67ffffffffffffffff16816000015167ffffffffffffffff161062003ffa5762003e237ffeadf1187eba19b4dedd01e784e492c0c20da5f2bb3a3692a67d41f5df24ee9c60001b620015a0565b62003e517fc71ccbf38068984707bb8d283618a998098b7cf6dfb2b27a736049139e917dff60001b620015a0565b62003e7f7f9e65e97b6e5a85e9acee41cf5afa3bd955dbdd51ff742835e46a24359c8e333f60001b620015a0565b838160400181815162003e93919062013bf5565b91509067ffffffffffffffff16908167ffffffffffffffff168152505062003ede7f558c6dddd8699dbbed23751de97d8359c6314f354aace96272f842c6bc694a1b60001b620015a0565b62003f0c7f87bcd5b2e5230e7e1ae3cd6021063ac3274ffd36a7d5cc10c14d589af1b4044560001b620015a0565b8660a00151876080015162003f22919062013c72565b8160200181815162003f35919062013bf5565b91509067ffffffffffffffff16908167ffffffffffffffff168152505062003f807f921cb1d70bcaf7d903fb866472d9952263c59de050bda85b0b36bafd2b382be560001b620015a0565b62003fae7f2cf463ddbcb1745e05483e5aff5c2ddf39662b2529cdd54554002fbab84015e060001b620015a0565b8660a00151876080015162003fc4919062013c72565b8160000181815162003fd7919062013cf6565b91509067ffffffffffffffff16908167ffffffffffffffff168152505062004096565b620040287fedc538cf02256fed4ce0eb662276466efedf12421ca36c9f8c7c7469b476bfbe60001b620015a0565b620040567f595a05b8db2d056bb1902a6f6cbd5e7c83ac179740bc04445a06d2718d265b1d60001b620015a0565b60016040517fc8c84b2f0000000000000000000000000000000000000000000000000000000081526004016200408d91906201363a565b60405180910390fd5b620040c47fc48475f37b0d45795ffaf6a5db31b820886bd10e00b04ea6c9970b50ad901e2e60001b620015a0565b620040f27f474d5811acd5ea3ee50823d08cd42550faf6b0ac439acca4873d19bb57c8764060001b620015a0565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ed108de98860200151836040518363ffffffff1660e01b815260040162004155929190620133ac565b600060405180830381600087803b1580156200417057600080fd5b505af115801562004185573d6000803e3d6000fd5b5050505050620041bf565b620041be7fa5529085b72aa4c0aa6fa65ba3b0d5fca2a4b26673f18abc2030ad75ee4602e960001b620015a0565b5b5b620041ee7f676c6b045ee52448a2839cfdeafadd1a579b1e6fe4214b6c3452d557c7f2c70860001b620015a0565b6200421c7fdf2f288b96b7c5b1712b47265cde57abbbc3fa5611d8d827c9db7b60724e3a7460001b620015a0565b6200422a86600180620048ab565b5050505050505b80806200423e9062013f51565b91505062002a7f565b50620042767f77457e8a0f34ae5998ddda0f18b716941b9a0a82b490404441b3c6e2be5a3dec60001b620015a0565b620042a47fb5c9253e2b19e7d9cad31d6a08a8063fe6f4e767ba3a6c95635b7f4d23494c5c60001b620015a0565b60008367ffffffffffffffff1614156200434b57620042e67feb0df40120ea2ee7285d6d8f9f912539375b874b094134f0b49b86c55cb22c2f60001b620015a0565b620043147ff9d75f8d49ef5a80ced7f72a7e94d1e45eb315f01eabfe1708b940e356c7056760001b620015a0565b620043427f54d9a71d4c2ccce5312953d7e1ee31270836f4d0abe9b7490f9f61863b890a5860001b620015a0565b5050506200442b565b620043797f2e0aa2bfdce8de0f057a2f81c2c3682b85ebc4fee4dd05a80535ba2c877b03cc60001b620015a0565b620043a77fce3e204081da479054d8dca9f9ac6d9dd7440620797cb88da4141b046d44461460001b620015a0565b620043d57f18b204f7426e4f90fae55fd63d354d515a315218c697fb49115b638c3824eb3260001b620015a0565b8173ffffffffffffffffffffffffffffffffffffffff166108fc8467ffffffffffffffff169081150290604051600060405180830381858888f1935050505015801562004426573d6000803e3d6000fd5b505050505b50565b6200445c7f0586e30b57be65e3ec9f86dc4cdf927f9886d8f5e73c596f6218c020ad717ebf60001b620015a0565b6200448a7f8d77dc23123b50fe58d4d181a26c2d360f4d8b44ecee2ebb4e73c8ae3a18c1e660001b620015a0565b620044b87f9a5653a2eaa6d6cdf3fb7187f885c0dcf3b3ffd70bcc64cc3014fb8febae87fd60001b620015a0565b6000620044c982600001516200ed92565b9050620044f97f8893787abbeb55c1e961ba93eb8d9c85f8ce04bd170e6a38dc3f1485e1b9b42060001b620015a0565b620045277f3f13b7f759a50557a7c26e9b9f6ee9990ce85966d48780e06b1248d8b59a403f60001b620015a0565b620045557fd2b19eb65fae328aa751afeed0f67d241dc4fee6d2fe97fc65d931a0d0c7967060001b620015a0565b816020015173ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614620045ce576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620045c5906201369b565b60405180910390fd5b620045fc7fd8c707514eac9dc587d13bf51e560c29a5dbef3c0b4263d785ab1be578794a4260001b620015a0565b6200462a7fe654e10850e2fcb6da9922b85f3f108bd37e12fefceff0b8738b6fd0f993037160001b620015a0565b620046587f478c582e21d9516412c65a8574c762cce218b8e79100fbc3bc22047484a8d79d60001b620015a0565b8160400151816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050620046c27fa0c83b9f239e712c72e7d53608b37be13d49a33da87536cd27bfed1fbc91c2d260001b620015a0565b620046f07fd24646d62f6bc931cb4e0f2dd0588bd65201134c2a2cfc248d42e1da68c43b4e60001b620015a0565b620046fb8162001c26565b5050565b60606200472f7f1429b295cf906a2d14788dd8ea97a649275d99383540b7bf96d73106cd23076560001b620015a0565b6200475d7f23c5a9f785df02655797d51d91f4f35352a1e3cb255214770895860bd24bdd5a60001b620015a0565b6200478b7fd43e529f56e455ac2f4fb45259e3b83f3e2aa2f872983dc4c36529a53d10149960001b620015a0565b600b60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b82821015620048a05783829060005260206000200180546200480c9062013ee5565b80601f01602080910402602001604051908101604052809291908181526020018280546200483a9062013ee5565b80156200488b5780601f106200485f576101008083540402835291602001916200488b565b820191906000526020600020905b8154815290600101906020018083116200486d57829003601f168201915b505050505081526020019060010190620047ea565b505050509050919050565b620048d97f958eed56fdafd1dee1feb6a2adce329ed83a818ca82d4f3e9951dfdb19a26a6560001b620015a0565b620049077fdb124247ff4c649c0c8e2b6341cf83d9c2d5addd8d00486e04962b0020e3bdb660001b620015a0565b620049357fceb0d9c67ce42d504cfe5bf5217887a65b6b815d8a0acac7b22ff3626d74343260001b620015a0565b6000836000015190506200496c7f3b10d4392a724f64d01ab62981bdbfebdaa6a138cabfb454527414f120c2a95360001b620015a0565b6200499a7f17df30f62f2c7c62e98a332cba855584657331a50921289a498bfe0bf42543e460001b620015a0565b821562004b0e57620049cf7f9d71ca31f8c70a1030a4a2562436b86b36a622113b961f401bd71adbafe8120e60001b620015a0565b620049fd7faa9c2eb2b7201248cc0c653e0a511cda314395fd9d5e868599b9f5505f24463560001b620015a0565b62004a2a7e8265f59f47a982857c668d7b4a65efdd282296c2bd2367f9740500f71087d060001b620015a0565b62004a358162009263565b62004a637f5d8428453f6e92ce2fad737e28c81e6d08e388ec07ace3fae69694d869221d4360001b620015a0565b62004a917f5eb30bafd704c7b6a22c401822cddc09f8209b2306c1d2104694f81fc6a07fec60001b620015a0565b62004a9c8162005d2c565b62004aca7fb8df03942981319cbb18aed20e2981ff1fa0ca501144cd96e461b892ffa5844260001b620015a0565b62004af87fce544cae4e90aaef5b2d9488288fc12c5e5982f87db30c702e6620702291e04660001b620015a0565b62004b088460200151826200bc09565b62004b3d565b62004b3c7fec78cf7e8e91585096452345627bb71b1045af1833a57f2dea84eb6d2023051c60001b620015a0565b5b62004b6b7fe2a0196d1f505023ee1ef74394e4f21479eaa79624eaf3db739dcf816d672ef360001b620015a0565b62004b997f9d8b1d6cd7b1a382665c171f128ddb755628372a260a3d6b5ac018eab5c1d4a860001b620015a0565b811562004ea65762004bce7fe8a61fe8d78f357872bfb46d8eaf2e075f221c61c8e53f73ee930fe745defea060001b620015a0565b62004bfc7f28e115199828f8b23b7093ef0f6a139054118a674b70085f61833f1fb777242060001b620015a0565b62004c2a7fb6fb671f0fd736857a0ea8f531bd27f125f7fcfadda6675becfa9878a202779160001b620015a0565b62004c3a846020015182620100e8565b62004c687f62a554a5e8c99d561be3d36a05e1dd5451322f31fbfb1ba62304cf6f3962777d60001b620015a0565b62004c967f76f219ddc649b089e4ea142cbddff8e18a04c5710858b9cfa435cfe818426e7e60001b620015a0565b60005b8461022001515181101562004d6c5762004cd67f2d4a186ef5dcd1948396f16a804c8d6057042635b1e076bec95154062962b48360001b620015a0565b62004d047f46abce34920238dc89a48186b9f06e80272b0a6b9347b2edcf09be302aeed18560001b620015a0565b62004d56856102200151828151811062004d47577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015183620083cf565b808062004d639062013f51565b91505062004c99565b5062004d9b7f1eb67b138d643afe5d08525cfd5ac3f8844274ddc973c2b9547bbfa7657acf3660001b620015a0565b62004dc97f816b5edcb116f5cf4b7f89b2a5d7a082708ad20ec82eedfdd49f81380e1eeb8760001b620015a0565b60005b8461024001515181101562004e9f5762004e097f33b9fa0f7a541df7837fad43c95bb632fe7df21fe07b9ffd62f9a039cdbdc58d60001b620015a0565b62004e377f41f5c625df63eb1fd41cee4e28869378bfb881add1bd86e3aced521570514f0160001b620015a0565b62004e89856102400151828151811062004e7a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151836200210c565b808062004e969062013f51565b91505062004dcc565b5062004ed5565b62004ed47fec59598a07ea9f6504111408edacad04477a537e83fce736aad95f2d8e8dd17660001b620015a0565b5b50505050565b62004f097fe648fc8c65a2ebedeafae6ec7519fd4397a1e2b4011c85238dfcdd5d3cef469d60001b620015a0565b62004f377fadcffe0225e69819f6e7e26c9dd6d9530ae911b021cd434fe70c7a42312ca0ea60001b620015a0565b62004f657fe5e5a57148cd6e37d886f38d7ef20525f6290d446b85e6f781747a32880319ef60001b620015a0565b600b60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190806001815401808255809150506001900390600052602060002001600090919091909150908051906020019062004fdf929190620104f9565b505050565b6060600080620050177f1b1b0212a93dfb3b904a5fdd751d7b502f6dc88a19fc0bbb752f7fe49f9af24460001b620015a0565b620050457f58494430edebd0bfc25dacfe8e9c116a568776ac7286d2a926b527b9d7c0069460001b620015a0565b620050737ff6828bf3114da7f08750dae437fcf16c48972b15cd9f330c6cbfcb0e73a5d7ea60001b620015a0565b6000865167ffffffffffffffff811115620050b7577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015620050ec57816020015b6060815260200190600190039081620050d65790505b5090506200511d7f2080fe531c704284ca9356f9302cd072d7a39d48ab8341227501d6e4ceef99e160001b620015a0565b6200514b7fea67f9b66be064445b1748b5f096828c679f2de93343a02076cff35fde7ed55960001b620015a0565b60006200517b7f1add6da8ee3f6ba578d4e447e109d644317c4022ba910bd61c45227f0393268060001b620015a0565b620051a97f120d3e6be5131e1413abf41f1fccb91d6cc4d5e5975e0ed3035fd8b7a5080c8960001b620015a0565b6000620051d97f8f85ea4682355a67f38023a766f7501af204ae69530ea809a22ee3eb3507363660001b620015a0565b620052077fafd419c1af0ff347baef3621ba4e45236cf645d10468edbce4533409c1429f0f60001b620015a0565b6000620052377f74b08d2e3558ef981c442a3b6f217925bc7caae9d7723a26aa7d28ccd8b3dec960001b620015a0565b620052657f86d041f35e25c5f6f2e7791f0113620939cc879649e92cd77e0e1dc8271c9c4360001b620015a0565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b158015620052d057600080fd5b505afa158015620052e5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200530b9190620122b9565b90506200533b7f405dd08ecaffd0f74468cc8551d16c96c7c6bc3b0e382a82e6128c25100a4a9e60001b620015a0565b620053697fa805888c8acbb6ca934192a87ac98cd34f9779df5551be637ff49ecd8fde23f460001b620015a0565b60005b8b5181101562005b3757620053a47f174b30a2fc5e050f8268f1355b0ed19e5dd85ec739f3e1b3624331c3fde9cb1d60001b620015a0565b620053d27f1cd874e1ee896136f87dba1cb043958a016cbd92452d2ccee2a0073ad7398bf460001b620015a0565b6000620054208d838151811062005412577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516200ed92565b9050620054507f506db94eb2080c8150dc396fd554c0c911906f22aad4267835d347c11a2cb03060001b620015a0565b6200547e7fd4bb94c1916f2b5d8096b0d607692e10247e648222e19c045b813504bf824a2860001b620015a0565b6000620054ae7f8220d22fec50a13c1077b57020d05c2d17d619cbfd6ee62d121cfb1ef489c63660001b620015a0565b620054dc7f555811514dee7d461e40903b051c65e8b6060e252142c0cdf0c4e9c292c6b87d60001b620015a0565b60005b8c518110156200570c57620055177f6a416f336b7f35827e0d7e174577f20406569be6daa4e7b3773a0a5d8b75c6dc60001b620015a0565b620055457f3855e23598b247d76ca0fa210cbb9aa86d0c263a1710c64194ca8ea08b08020760001b620015a0565b8c81815181106200557f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516001811115620055c0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b836101e001516001811115620055ff577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415620056c857620056347fa943e4a93fee2ab99bf1ca9ed36f4d052df1f86be70336acc5fd89e40394422060001b620015a0565b620056627ff02c8956f11b0335c7e2eae97034a6e64091dd767989a9719b02c5a78f284a8360001b620015a0565b620056907f1cb610e47574e5666dd3bd891e2f445445ebd07a3c18844d3b41fb6b8d2a37fc60001b620015a0565b60019150620056c27f9f10e5b9647235fa3306ee2b62dd9bb3c335e3b077713037789703104944f2fb60001b620015a0565b6200570c565b620056f67f1d1c75644c2971330c0b84fe3c988168a543487eab9a89612feec8bf649444ba60001b620015a0565b8080620057039062013f51565b915050620054df565b506200573b7f5f75ab263a06300b4372c2bb83f5dda60988655bec508216def027de7827711060001b620015a0565b620057697f66589873f30c5f2b42addfdad2ec2d4fb2c08e7add63cbeac0c9cd67a11a0ce060001b620015a0565b80620057d3576200579d7f4190543ce769b9973d293c75d872d10ac222300cfc26ddba3e019d097d70b59260001b620015a0565b620057cb7f67abbd33dfc55f75ecc348ca69152ba7285c5392748b74df2ce7f0c7fdf6321760001b620015a0565b505062005b21565b620058017fb4b6a28a49a7092351d025e37a9e0c18e8e2e0b818d2ab7caf0e8f67ead364d360001b620015a0565b6200582f7f5944910d27bd38b2ac9cb780fe9e7937407a6cb75683a1b0cf8796a1044fa61060001b620015a0565b6200585d7f165944f80b87f76e325962a6b249087d508c16881b11e470403855580a006f3a60001b620015a0565b438460c0015167ffffffffffffffff168361010001516200587f919062013b98565b1115620058ea57620058b47fe13da2b1b58484afe5e3677da2e17300d01024d06769b69b5b75f24e094129f660001b620015a0565b620058e27f2fa34e256e0cda332c9fd48aa33b0c84ed6453df96ace899eac89f497082f32660001b620015a0565b505062005b21565b620059187fb9de3f8d94fc35de895453f5230018b4c6f43c569d54ba27aab1a693258e43e460001b620015a0565b620059467f66e0871a83953cd29d41d2e4f8a9e4c6cb9ec0a727b2020cc8cb0c7748fa37f760001b620015a0565b620059747fc735706bd6112444995bce76f422654ff0d79ff7994747ad50d87aabb850676e60001b620015a0565b8161014001518662005987919062013bf5565b9550620059b77f26311b37e4bfd4755ffafe382e1d71071623e5e2a07cddd67daceb520de2c64060001b620015a0565b620059e57f9e8683e1b0b5aa14a71406c61d4184d91ececebe77af04edeb6a20ee474aa36260001b620015a0565b620059f382600180620048ab565b62005a217f1e74b958b6c76e3d0d4e7f3f937b38d9a8854a3632b64f0674c51ce2f1e23e1f60001b620015a0565b62005a4f7fea4bdad1d85a9ab02103c89a505c7a96f6539c654d59d16a4af43288ae5b2a9b60001b620015a0565b8d838151811062005a89577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151888867ffffffffffffffff168151811062005ad5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018190525062005b0e7f439242bc37ca4a4c87f069e153edb6fd6a8c67dd1b3480d6ca7c0a06f8106d3f60001b620015a0565b868062005b1b9062013f9f565b97505050505b808062005b2e9062013f51565b9150506200536c565b5062005b667f3b1331e7f62e037876d65d6c88c662a459747d88d88dd89e2761174a8a1ae5c660001b620015a0565b62005b947f0c22055de75bd34d8b9103d179cffaea3e66e550b482d864e78158ab1a4554ad60001b620015a0565b60008367ffffffffffffffff16111562005c8a5762005bd67f354630a43ff387e74a2fd2d8f3083c2e62a1e2d3c4af9d6dde6100160d9696ee60001b620015a0565b62005c047fc8f418fa287040deb2d2097e74f3b9bc872088318b47b468b5cff84470270f1760001b620015a0565b62005c327ff14f7e455a52bc8a661d6f61fc30432c3e29157dd5ca7002e229620558851ce460001b620015a0565b8973ffffffffffffffffffffffffffffffffffffffff166108fc8467ffffffffffffffff169081150290604051600060405180830381858888f1935050505015801562005c83573d6000803e3d6000fd5b5062005cb9565b62005cb87fb14e8ae0fb2f14b08c41fdededc353b50721ef93564eabeb5fb668765a96041860001b620015a0565b5b62005ce77f1454479f0c8650b4809ffba427d909ae52c97123c24f000e444041f78da0175660001b620015a0565b62005d157f6806dbdaa4e5b4063411a681076c3b69d21ba4660663c76b667409d899ab1f2360001b620015a0565b848383975097509750505050505093509350939050565b62005d5a7f639c1b25dc58bacd876d45ec5d597116c0c2cef6667a1ddc86875217813c15a360001b620015a0565b62005d887f92c1de1403fbda8fae9718b088715cfb8f3787c4da1dbf44164e8b9e4a6bcb6360001b620015a0565b60068160405162005d9a91906201335d565b90815260200160405180910390206000808201600062005dbb919062010619565b6001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560028201600062005df4919062010619565b6003820160006101000a81549067ffffffffffffffff02191690556003820160086101000a81549067ffffffffffffffff02191690556003820160106101000a81549067ffffffffffffffff02191690556003820160186101000a81549067ffffffffffffffff02191690556004820160006101000a81549067ffffffffffffffff021916905560058201600090556006820160006101000a81549067ffffffffffffffff02191690556006820160086101000a81549067ffffffffffffffff021916905560078201600062005ecb919062010619565b6008820160006101000a81549067ffffffffffffffff02191690556009820160009055600a820160006101000a81549060ff0219169055600a820160016101000a81549060ff0219169055600a820160026101000a81549067ffffffffffffffff0219169055600b8201600062005f4391906201065f565b600c8201600062005f5591906201065f565b600d820160006101000a81549060ff0219169055600d820160016101000a81549060ff0219169055600e8201600080820160006101000a81549067ffffffffffffffff02191690556000820160086101000a81549067ffffffffffffffff02191690556000820160106101000a81549067ffffffffffffffff021916905550505050620060057f46df501b828f66bb3584107d7d6bba29236465293566e299cfecaf4ade57723e60001b620015a0565b620060337fdb703d80a5d70bb51644e2a3efe9d313984bc908f8a3478e3a76eccc08f45c9460001b620015a0565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166352e06146826040518263ffffffff1660e01b8152600401620060909190620134b6565b600060405180830381600087803b158015620060ab57600080fd5b505af1158015620060c0573d6000803e3d6000fd5b5050505050565b60606000620060f97fcd5cfe51a4c6469b3d1e24f084efd9b4b658068c35f2641bf704b18c5fa5f08660001b620015a0565b620061277ffd83ca1eeb55d9e798ef6a1e5dda8484f9766f4036a06ac12761b480d31f87c460001b620015a0565b620061557f0239c4dd6dc1065463e0a22351d7a7b5508ee169c2c793e0a063e54a7ce564f560001b620015a0565b6060620061857ff66051e17f76627f2d53ed53656cf358851d8db646b8897a2e222771544aa56760001b620015a0565b620061b37f26d2ac6fd0a897c30bb5bceb8173852a7d5caf31fe24647708eaa6fa126f64eb60001b620015a0565b60005b86518110156200668457620061ee7f60065f963c9d71e293f25e58671f200bf9c513b521353ff94b85d9c64363bf7c60001b620015a0565b6200621c7fb75ad8e05510fe1ed6e7e20dc0a386002b1cd28947be51e1fc38667ee9d949a160001b620015a0565b60006200626a8883815181106200625c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516200ed92565b90506200629a7fcd2002e3dec244b5e4322e10bcb6257fde29577ea868e4dd6245449f9cdf27d360001b620015a0565b620062c87f84efe295e61d25dc20dabe94318bff15bda9b00690a4ebff604745cd06db81a560001b620015a0565b6000600181111562006303577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b816101e00151600181111562006342577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14620063ab57620063767fbca3d29fa5b2281a996c894fe154e36f6ecba9491dd5bb46e9e915872d29696160001b620015a0565b620063a47fa784c98d9bb91970f40fd7e3e55ec1da4a73af358c43af7465bfaefe077379ce60001b620015a0565b506200666e565b620063d97f5f049a6d20ee8954720ed052d48d10fdaeebb9be4ce662522fe2e273c68625da60001b620015a0565b620064077fc328f9f4aadb4dfcfa0c933dcde1900deb4951fa9c98c89ce86eab49f1dca15860001b620015a0565b620064357f0693eeca19bea5f4cf8939e06670b9790c00644db5c2ac89a875bdd1f6dc3dd760001b620015a0565b8061010001518611620064a557620064707f56c2a32830c0bec428b1f3e93116bd115b6a10a18b37381f30865c12969dca4260001b620015a0565b6200649e7f721421755122ba597fe3e094786facd8dafa166f425d5711822029108862247860001b620015a0565b506200666e565b620064d37f7cb5379989cf75f03acf505a4e88c9ec793e698c9486805acbc3f3452c4e4ac960001b620015a0565b620065017f7b26c618f5595a902926dce4eba153c186e7299d75e3a66158b669ff0ace33c960001b620015a0565b6200652f7f217e0bbf81979bc8174959e73b12e34858f32bd7130f7c2c9c3f1f3d61081cc260001b620015a0565b60006200653e8888846200f9e1565b90506200656e7f169ad7ee9ba9c6d102fc4f9e85adac96f05101adc28ff38a5198834fdf51610d60001b620015a0565b6200659c7f084f4205d7c8296cb9d766978521980e1bd49132a0a76755f25a49f7b8318fea60001b620015a0565b806200663d57620065d07f0adfa40c30c7870865b0657f8dd46c9a734012328c182263552d2e1be584a55960001b620015a0565b620065fe7f052e70b433acdb52f3e35dc9babc521180c17cde8039644c72266f70b8512b6360001b620015a0565b6200662c7f1f0f5b6c23a32892d8d203b8dac8cb283bfe2410378642cdcf962100f351de3a60001b620015a0565b8360009550955050505050620066ea565b6200666b7fd04a05a7cef13603c2c665d1dc0bc684a446826863efc95b6b46acb8efafa76e60001b620015a0565b50505b80806200667b9062013f51565b915050620061b6565b50620066b37f2a7f931bccf0c4219b8fb3268df5c8ea5f8cfcabe1725b7fb82fc67bf4f4463e60001b620015a0565b620066e17f7e12c3dae6e9b1186c1d66c29c9ccb7a8dbfe77407e49d87a4084c6ba7607f0860001b620015a0565b80600192509250505b935093915050565b620066fc620104ba565b6200672a7f447e082935b60124cf7d67a9e39b717c25b8941c25b81201e109122845824c0260001b620015a0565b620067587fbf63ee34c621f3a9272e67cd11cd0b6768c19765d9548255c8a9ccdd78440b6460001b620015a0565b620067867f63c9253e2607d9058069a3c09f2e1091e81594f4ae5eacb6de828e744da0425860001b620015a0565b620067b47f867a02572bd89c1dfd96049f4a189f54d37f73a254a81f9f491b3b8465e3be8b60001b620015a0565b6000826020015167ffffffffffffffff161162006808576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620067ff9062013723565b60405180910390fd5b620068367fbaa62176caf78a3cba6c909e27dda9360cb993374322f1412476dd001f5ad70d60001b620015a0565b620068647f57a8ff50a61d952950fad07c2466d071a00684c355d3ec578c8f75414d46ca3b60001b620015a0565b620068927fd0b1a533f9b5ab69bc5e3f2d8b097a71e01cf8a819ae45ab5bb0d166b5f852e960001b620015a0565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b158015620068fd57600080fd5b505afa15801562006912573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620069389190620122b9565b9050620069687f283c882f63869cff2b4c5701680a09d34b9776baaf3eea492db71b41c3d27be560001b620015a0565b620069967feae3ca41b54342974c3cf68f821e34dc39f4c6d3358928868e6fc64cbd987bea60001b620015a0565b620069a38382436200dbcd565b915050919050565b620069d97f0c6ae0a4e89937a95680b764edbdd4b873d70af1f8507d1e6f2b992f1296c2e360001b620015a0565b62006a077f28c5aced2023321b4ccfdee6c001705366f67b98ae0f90f96266205d80b4793d60001b620015a0565b62006a357f5d02a89c7e25afc1a50c75bae83ced379c1c52bd1c9040129708c29a01227c8e60001b620015a0565b600062006a657f95be1fe609c3d716f162de967a2470a9d39fcaab833d5b8f03a78d9fcf853a7e60001b620015a0565b62006a937f7c94a48499c9fe2708fbcbf5259341769177eb5ec12d75324b532b0fee70043360001b620015a0565b6000825167ffffffffffffffff81111562006ad7577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801562006b1457816020015b62006b0062010682565b81526020019060019003908162006af65790505b50905062006b457fdd0206a010ef963d409282a886808b7b9381e7454be786adfb12fcd58a27f2b360001b620015a0565b62006b737f12c5ce26c452a7b83292a1cf9e94f651047444af3a2c471652bcdc370ea426a360001b620015a0565b60005b835181101562006d485762006bae7ff8e0c00ea63238e947730512e66a045c88980cfc2a5697b52f76caccfd20201a60001b620015a0565b62006bdc7f6b3db2a2a89b712e39dc8cd29bdd6018a1c67f75275d4331ee63c21a8da75c9060001b620015a0565b600062006c2a85838151811062006c1c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516200ed92565b905062006c5a7f4ec6c0b28c10db44e5c2a084bf45824f83c55e87350afd7033d16140038ece6960001b620015a0565b62006c887fcf0e8737dc93a82d5257532020ee7a703e1e3ad256ad7968c440a2fc02aaaa7e60001b620015a0565b8060200151935062006cbd7f3a15c5c0c506433295c2dafc65e14f508ab36d895cb94f48a4694d093273863d60001b620015a0565b62006ceb7ff85fb9f7dd8a2bebdbe2e8c4cb5cd62aa685ab1164d7f6a3711fe088579d275c60001b620015a0565b8083838151811062006d26577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018190525050808062006d3f9062013f51565b91505062006b76565b5062006d777f057e2a5b0b9fd7dd1d9bac5efd407dabcb9fc0a7d5575ba9019419808e6e81ca60001b620015a0565b62006da57f9f1bc3830680a2f300bd0a956aa5afccc3c783b9519f5574d28be70fa29a7e1d60001b620015a0565b62006db08162002451565b62006dde7f107b92c1cd97edad0fd90de45b6a8ff55ff2dcdad83c5527b493d30ffa2a9db660001b620015a0565b62006e0c7f4304d515de4d365cec1c9add6533ad68533aaaffc20c543e0568226c7ab1f1ed60001b620015a0565b7fa5d0a6140e999641864ed8958af5eea1d36d821658cfc056552962fab40291ec600143858560405162006e4494939291906201350e565b60405180910390a1505050565b62006e7f7f3e9991811ccf4429588f79283ed27119012f8983ad2ea50e40c9815621d7b1c360001b620015a0565b62006ead7fae21f9f6f684141ff5d0e00271fdcf59775d4e662a63122b170169f8ea82304660001b620015a0565b62006edb7ff6d73a5c93d5376b2a0669dbc907714570bc06af6d6e2d7a659b73833e0e81bc60001b620015a0565b62006f097f5b1e2d0fd003c0a913297c15e867afea02b03711635ec908206edd08decae66760001b620015a0565b60006006826000015160405162006f2191906201335d565b9081526020016040518091039020600901541162006f76576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162006f6d90620136bd565b60405180910390fd5b62006fa47fc232047193d89dd107fb5857b0f6ae860b7c043fa13331f950e16d8413ca01cc60001b620015a0565b62006fd27f31b02e90378b75fe911716448a115407f943ca56d98b29d63f4ca1a53be7c7fb60001b620015a0565b620070007f1a27cff00c88e431b78b0d26982a3c41ddd60fd7eea6fb0acac22033e827c91c60001b620015a0565b6200702e7f981b9648ec972d2e47ea401e022795346b2128779d7384752120816e21e7665160001b620015a0565b60018081111562007068577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600682600001516040516200707e91906201335d565b9081526020016040518091039020600a0160019054906101000a900460ff166001811115620070d6577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1462007119576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620071109062013701565b60405180910390fd5b620071477fe499b64fe35656dc0a4c4e77a8983871300fbfee55c4945b3dfc583c06ced0fa60001b620015a0565b620071757f81d6e1ec87d8279ad412de3dccaa256078557bafb8a11a458f572bd480e4ac6360001b620015a0565b620071a37f9e329ed053749ac56adb3c618eeed7dd79fb666e3fca83643d237a4ce3b33b2560001b620015a0565b620071d17ff2bc6fc60153386fd8ce0fc6133387c918ec7d796b57257f82e8903235746c9f60001b620015a0565b4360068260000151604051620071e891906201335d565b908152602001604051809103902060050154116200723d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620072349062013679565b60405180910390fd5b6200726b7ffbbf85826141628e49de79f7f404deaee0dc7a767b333eae663b55a71c19ea3f60001b620015a0565b620072997f961b7992dcc5345cc732c81d7e18b2f0ddcba1af02b03585d06d2f19d75c3ed060001b620015a0565b620072c77fbb385af377f8540ddb2be261478f1f59fef205dc5b8e6f332e1f8e088954922260001b620015a0565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156200733257600080fd5b505afa15801562007347573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200736d9190620122b9565b90506200739d7ffdb5bb71d145d5b3fc9d0e09ef10b30290a45e5c19b4e3f492b94f35d08b4c1860001b620015a0565b620073cb7f921d46e74b139a78908d3ba9228a22dc320bd3db2846540ac802e86033f1ebc460001b620015a0565b600060068360000151604051620073e391906201335d565b9081526020016040518091039020604051806102c00160405290816000820180546200740f9062013ee5565b80601f01602080910402602001604051908101604052809291908181526020018280546200743d9062013ee5565b80156200748e5780601f1062007462576101008083540402835291602001916200748e565b820191906000526020600020905b8154815290600101906020018083116200747057829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282018054620074ff9062013ee5565b80601f01602080910402602001604051908101604052809291908181526020018280546200752d9062013ee5565b80156200757e5780601f1062007552576101008083540402835291602001916200757e565b820191906000526020600020905b8154815290600101906020018083116200756057829003601f168201915b505050505081526020016003820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600582015481526020016006820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016006820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600782018054620077019062013ee5565b80601f01602080910402602001604051908101604052809291908181526020018280546200772f9062013ee5565b8015620077805780601f10620077545761010080835404028352916020019162007780565b820191906000526020600020905b8154815290600101906020018083116200776257829003601f168201915b505050505081526020016008820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160098201548152602001600a820160009054906101000a900460ff16151515158152602001600a820160019054906101000a900460ff1660018111156200782c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600181111562007865577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600a820160029054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600b82018054806020026020016040519081016040528092919081815260200182805480156200792257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311620078d7575b50505050508152602001600c8201805480602002602001604051908101604052809291908181526020018280548015620079b257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831162007967575b50505050508152602001600d820160009054906101000a900460ff16600281111562007a07577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600281111562007a40577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600d820160019054906101000a900460ff16151515158152602001600e82016040518060600160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505081525050905062007b397ff95b5cc806e1fe96d109a790d392fe21d65b6471ea42fa98d2ec90996aaa702160001b620015a0565b62007b677f61c3ec6b41aa5d892cc158332e078622b7675d7b28e5ec52d38631d3b490c8a560001b620015a0565b600062007bab8385604001518461012001518560a00151866080015162007b8f919062013c72565b8660c00151896040015162007ba5919062013c72565b6200d5c8565b905062007bdb7f1b43559dd947cb5e13cc583585127d502f7dbeef1083fe15c04a902c60ac7f3660001b620015a0565b62007c097f44c50de07efcdc0fffbc404d9eb5aeb2ec073688d7d98dd78767305c1f4676d160001b620015a0565b60008160200151826040015162007c21919062013bf5565b905062007c517f8d64a5094b58eb77946f12f86d32a680fc0026fac51a7ec4412539f426861c6f60001b620015a0565b62007c7f7f9bcadc2a0e6ceb48581e1001a638e8cc24c6274c639c91bc341f173338cf071760001b620015a0565b8067ffffffffffffffff1634101562007d2f5762007cc07f40ce558a527bb9faf205d2546431885ce64f6e7917e55575a3c61a38071abc5c60001b620015a0565b62007cee7fb6be8962502d55a9414f78d2d0b138d4d5d13bd187a08304b744e7c5049510a960001b620015a0565b34816040517fb0b78f4900000000000000000000000000000000000000000000000000000000815260040162007d269291906201388e565b60405180910390fd5b62007d5d7fcb2004a05b6ae580e073350a258ab5e7655298e0517afc044fd98a72dda2712360001b620015a0565b62007d8b7f44d5e14b343b985dcaa77ac740c1d4e494392de3d8e179b73fcbff884e0d89d160001b620015a0565b62007db97f5eb9fe692af152eeaeae6cbcffe67da91d08ae8572acdf60b88869dcf5a19a1c60001b620015a0565b84604001518360e00181815162007dd1919062013bf5565b91509067ffffffffffffffff16908167ffffffffffffffff168152505062007e1c7fe66b23ecc532fe479a1fa2d9d96b13f9ba1d0374b9a8b19e1d4f73369a895a8e60001b620015a0565b62007e4a7f7c795242005044e557eceb72bbe6d71a4861e15b7ae779e347f62e7cd31109bd60001b620015a0565b80836101400181815162007e5f919062013bf5565b91509067ffffffffffffffff16908167ffffffffffffffff168152505062007eaa7f3f1d625862e8616925d4bc4a13bcf40dd24712f7b4fc4d3b3d9c315152cd5dce60001b620015a0565b62007ed87f42e0b5c64f3aa706d678eb71d6e13564fab06e1bec72830f91450c133eedaa5660001b620015a0565b84604001518360c0015162007eee919062013c72565b67ffffffffffffffff16836101000181815162007f0c919062013b98565b9150818152505062007f417fa2999d9ddd7087cb98b36a781b16a450a2ba2d008477701bd368fe8d1fe547fe60001b620015a0565b62007f6f7f5ad31e9a4cd2ffb5f7e0db63b08f5ced4005b01145383a781bae521849f402f060001b620015a0565b826006866000015160405162007f8691906201335d565b9081526020016040518091039020600082015181600001908051906020019062007fb2929190620104f9565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201908051906020019062008018929190620104f9565b5060608201518160030160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160030160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160030160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160030160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060e08201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061010082015181600501556101208201518160060160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101408201518160060160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101608201518160070190805190602001906200818e929190620104f9565b506101808201518160080160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101a082015181600901556101c082015181600a0160006101000a81548160ff0219169083151502179055506101e082015181600a0160016101000a81548160ff021916908360018111156200823e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555061020082015181600a0160026101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061022082015181600b019080519060200190620082929291906201058a565b5061024082015181600c019080519060200190620082b29291906201058a565b5061026082015181600d0160006101000a81548160ff0219169083600281111562008306577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555061028082015181600d0160016101000a81548160ff0219169083151502179055506102a082015181600e0160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050509050505050505050565b620083fd7f1f94138e9a33b18d4f1458abef85f3863f456e5b1aa7e26077ca85085bbc06a160001b620015a0565b6200842b7f772e38d3099bf19a8dfb845f612c923be2eabb9a7f2dad03ba2d1e4be13dcfd760001b620015a0565b620084597fb860fb88daf96cf03b146758004ee725ec265e0dd7001c60be8b99d7f4c5f18a60001b620015a0565b60005b600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490508110156200870f57620084d57f4de2d54671678c15be153bb8d329c0fd442b46834bfe03fe56cc62b0d034aeef60001b620015a0565b620085037ff4002df7aae4444c68150d646ab3ff0f9a2e3b3d804623f8cff010e6538e62ca60001b620015a0565b8180519060200120600960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020828154811062008583577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020016040516200859c919062013376565b60405180910390201415620086cb57620085d97f80e0526bfd7cfda7d7022153ac33fdb4dff324cefc7644e68e175be57becf66a60001b620015a0565b620086077f41225f43f2382c9f83c85a90f86600e9857936fbf9727ac048c70bf399dad05460001b620015a0565b600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081815481106200867f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002001600062008697919062010619565b620086c57f619b9e0390f34a875722793ad5ed0917e4db3a5ff2b1f7e5ded979393feb53ba60001b620015a0565b6200870f565b620086f97f1dbab3312bb21e744593709286852d193ffebea4ad8265edfa2852903cafcb2360001b620015a0565b8080620087069062013f51565b9150506200845c565b505050565b6000620087447fb195c33d923490965474ed61da2086b01cea0bff8ec52609939c3283785965e560001b620015a0565b620087727f3d36a9f405bc0ce9ea63a7a293f5191fb1fd905b40caa99ca7c960d6cec54c4360001b620015a0565b620087a07f4c93587fe3303c8f4b9ce7cb659791580e0bb9c4087f3120911a1965381ba62d60001b620015a0565b620087ad8584846200f858565b600185620087bc919062013bf5565b620087c8919062013c72565b9050949350505050565b6000620088027fc3f82c4d6c9ab658616bf090fa5f016979e0dc644a6c91614be4a5f4fdfb021460001b620015a0565b620088307f11f58a6b53e0398335c024a1e1b668ef5f10bc496bbbebd5cb5cdd5af3234b3560001b620015a0565b6200885e7fe74735d3ccb835b9096ea2f5e4307ae72fd5de874483b9a1b1b306713ae7563760001b620015a0565b6200886a84836200d512565b8362008877919062013c72565b90509392505050565b6060620088b07f388baf0212ca8889cebfc57d28ea7c190d1b5b0b04ac62b3adb10417f30d2dc660001b620015a0565b620088de7f53f0fb8a16094b39ea00fd8ace536e378e685671fbd9dfd5b0b8eb5687912a6960001b620015a0565b6200890c7f7746d3145f2d20f9e2c98fdb18b0827008f5771f206d363ea6dc0bee54142af560001b620015a0565b6000600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101562008a235783829060005260206000200180546200898f9062013ee5565b80601f0160208091040260200160405190810160405280929190818152602001828054620089bd9062013ee5565b801562008a0e5780601f10620089e25761010080835404028352916020019162008a0e565b820191906000526020600020905b815481529060010190602001808311620089f057829003601f168201915b5050505050815260200190600101906200896d565b50505050905062008a577f79b8ae175be0a6c5474e1da01b6c33146f8316ad20e6d4629cc21e65eb4dfc5b60001b620015a0565b62008a857fd5ea27528f4bcab085eb2c96b80c4e9a57be31cbaa88865713be71bd9da6b02660001b620015a0565b6000815167ffffffffffffffff81111562008ac9577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801562008afe57816020015b606081526020019060019003908162008ae85790505b50905062008b2f7f5661b06e75adb6526474e1b3ff2b51492410037aaa8e5be9d62592e2c1e70e8060001b620015a0565b62008b5d7f811ca069aee492ceebd10e4bcddf4cb095aa3f7b05827a5ad920702b1e249ab760001b620015a0565b600062008b8d7f7b23d086bdb40b89524c7ebbf6efdcbeb1ad51138ef78ff843bf3cebe044b4f260001b620015a0565b62008bbb7f282a2e203bf20e4b12d64adac9b63ec9eef99d7e7013eb7a117eb7384fbfac2360001b620015a0565b60005b8351811015620091fb5762008bf67f7ded9006d32f67ec194c72a67bdc3091afefe3352479e2332220f68d140bddf860001b620015a0565b62008c247fc0ad66a5be81cb1f141ab5aeb52b146f7163c726c4929b8a766e14335e4c9a3f60001b620015a0565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663977fdfe286848151811062008c9e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b815260040162008cc49190620134b6565b60006040518083038186803b15801562008cdd57600080fd5b505afa15801562008cf2573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019062008d1d919062011f8d565b905062008d4d7feffd430b4ee804f8cd2c2340b0c4e9ab454ad200d9eb1cc1fc941dc200ccd21e60001b620015a0565b62008d7b7f507cc4bb977412c31099b28cad2199c8537b4ba2041befb2caad5bf1cf92f0b360001b620015a0565b600062008dab7f4ed64843a3ca1c5d48a23d6a2cc8437a307ab09e430ec152b7eb806ca26c34b960001b620015a0565b62008dd97fe259954f4165993c53f66e49a4c5658bb56b75858681bd750b6e8fbfabf007e360001b620015a0565b60005b825181101562008fc25762008e147f504625989b9ec2c61960b1b1f38cec27b159d6813d95835e2cdc773142a5692d60001b620015a0565b62008e427feaee30452ea629889d22a584c8c806c8ea340ce9c9182975e48c1e2dd98d29f660001b620015a0565b8873ffffffffffffffffffffffffffffffffffffffff1683828151811062008e93577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff16141562008f7e5762008eea7fbdfcf26de3561f1e8b68867cf2760b2760aaff774478a7d39796d163480a8e6460001b620015a0565b62008f187f9f972284d511d7c3ee430b88dc4d5e4a45185740880b7a0504940527fd8b407a60001b620015a0565b62008f467fe2c77ed87cedf8dddffd03196a9f7a22f4ee58367060e9c2158b55a794fe0c8a60001b620015a0565b6001915062008f787f460a25f2acad0469abaad3d90531af060c3730ede7fd8581d157ef4e4c7e394060001b620015a0565b62008fc2565b62008fac7f10189c223c7807a0b5ca34fd36833fd7f5a799a00731b4a1157e966326a81e9560001b620015a0565b808062008fb99062013f51565b91505062008ddc565b5062008ff17f8c378dd10f86bbd1ccc62258965490dba24c8f9120ae381082c54bfc025e7a5960001b620015a0565b6200901f7f4a9751caf88e39c9ad5ed65d96db81398384f9a37bc6024301698dde0ff6a4f660001b620015a0565b806200908957620090537f42ef728a33b1349b163f6a139ee1a3b6e550d25b64da49e05f25f4fb36e6b2cb60001b620015a0565b620090817fab338a37a9a4ed52f8a714dd7943b321c5bd474d1a46624cd99f8cd00d0648cc60001b620015a0565b5050620091e5565b620090b77fb9da73a04beb3ff88caebedd89002391d5935a0ac75b16810bebc819957e5b1b60001b620015a0565b620090e57f240e9da687309d02d789b76bab78e360733d6de668348f937744fffcc3393c1d60001b620015a0565b620091137f702bb253c15ca8c6e28e245e9635668b33ff78d35698bd73c3c34197f46a480b60001b620015a0565b8583815181106200914d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151858567ffffffffffffffff168151811062009199577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181905250620091d27ff531e3c8787290c8678a3e2f0f8bbd86077fa896fe4bb5bd750189c3ec4f527e60001b620015a0565b8380620091df9062013f9f565b94505050505b8080620091f29062013f51565b91505062008bbe565b506200922a7f056756053e5182fd9bfaa62c53a52f52fbec57565d531b067d34de72a6be260f60001b620015a0565b620092587fdb01789ab99ab24e5787b38f325d35c5e01c49fd8c2ebf7a7cb4fb4d535bb3f060001b620015a0565b819350505050919050565b620092917f87b4148840511a558d9520aa269ad194c3befdd158a8313e31f245d5d654391d60001b620015a0565b620092bf7f2c0dc343a7de0fd692ee77940655c221a4f426e744f119b36e5ee7300ebb5db160001b620015a0565b600681604051620092d191906201335d565b908152602001604051809103902060008082016000620092f2919062010619565b6001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160006200932b919062010619565b6003820160006101000a81549067ffffffffffffffff02191690556003820160086101000a81549067ffffffffffffffff02191690556003820160106101000a81549067ffffffffffffffff02191690556003820160186101000a81549067ffffffffffffffff02191690556004820160006101000a81549067ffffffffffffffff021916905560058201600090556006820160006101000a81549067ffffffffffffffff02191690556006820160086101000a81549067ffffffffffffffff021916905560078201600062009402919062010619565b6008820160006101000a81549067ffffffffffffffff02191690556009820160009055600a820160006101000a81549060ff0219169055600a820160016101000a81549060ff0219169055600a820160026101000a81549067ffffffffffffffff0219169055600b820160006200947a91906201065f565b600c820160006200948c91906201065f565b600d820160006101000a81549060ff0219169055600d820160016101000a81549060ff0219169055600e8201600080820160006101000a81549067ffffffffffffffff02191690556000820160086101000a81549067ffffffffffffffff02191690556000820160106101000a81549067ffffffffffffffff02191690555050505050565b6060620095417f62a02aba0a01aff589b3934d3ff3b584bf6585f0eae0e29ce27520a14d44df2a60001b620015a0565b6200956f7f390e2f304bae752d8cc90c59ecea01be9e709d948682ffea2a7a20fd069cb8e960001b620015a0565b6200959d7fa561962b7a6dabdccf63b8a3ce8f86b536e8dd05d5ac95390281f77a2974d40460001b620015a0565b600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b82821015620096b25783829060005260206000200180546200961e9062013ee5565b80601f01602080910402602001604051908101604052809291908181526020018280546200964c9062013ee5565b80156200969d5780601f1062009671576101008083540402835291602001916200969d565b820191906000526020600020905b8154815290600101906020018083116200967f57829003601f168201915b505050505081526020019060010190620095fc565b505050509050919050565b620096eb7f67019cf9fccb4f80a4caee87009ece260606c01f458095d5ff841b3b4ac1c80560001b620015a0565b620097197f129579356cca7a0a41602dce565cddc4d9aefee3d2da3ba6654fa3863ab8651860001b620015a0565b620097477f2fc6a7e4a7e99902408144a526d72240ed118c5532b26e3de1bcd113cb9993d160001b620015a0565b620097757faa1a1b7cf3cb3eadbb2ec979cee1b256ce63ba462af24fe34132f5fbfbf6258060001b620015a0565b6000600682600001516040516200978d91906201335d565b90815260200160405180910390206009015414620097e2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620097d99062013767565b60405180910390fd5b620098107fe8140b4abe7593d0578e30c4e3e3969f8718eafb61539843ed8a2ddc5e835f2860001b620015a0565b6200983e7f34b11acbd5fa3a61151082edb2b8b7d99af998f5ed49753604c1b1653aa9c5cc60001b620015a0565b6200986c7f61f794a4e9b4d35baf4a0cfc5eb8b347bd3eceebfa177a817c6dcc21f5b5cccb60001b620015a0565b6200989a7f0769876d7522279fb30115f9fb9ce7133c4328dc1eda15cf1dddc2493e56c12660001b620015a0565b4381610100015111620098e4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620098db9062013679565b60405180910390fd5b620099127fc1b826d6bc33134e676fd18c11ebf1e38979453839b38b6a2e5461dc13f47cc860001b620015a0565b620099407f15410b331db777f9e3124cab1c29961c540e596de47829aaf50bdf1526ab1fa460001b620015a0565b6200996e7f4d423ee8d384a2e9421a6d127a1f5989671402b3841959e160d50a1639ac070d60001b620015a0565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b158015620099d957600080fd5b505afa158015620099ee573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062009a149190620122b9565b905062009a447f03bfda97ce2f32529924ce17aa6b6679bb2504bd6b32441b1f10e4993c8760ba60001b620015a0565b62009a727f2f443119e878cd8c96fff0abed6b0e95bb4a64de85744a05785604855841d3aa60001b620015a0565b6000600281111562009aad577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b826102600151600281111562009aec577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b141562009bb95762009b217f1221c6390881ff415a4bb3ebe9fe07e6e4ba5dc8751179787a9bb9ced6733d6f60001b620015a0565b62009b4f7f11b63e4db0ff33f4ec928284eca1145cbf6a664bbbd4e2fec712e46f7ced3f5060001b620015a0565b62009b7d7fb80fd6cb192a74545fb0006d28744b6a09cb317cf265bf450e816d77fbc352af60001b620015a0565b600560009054906101000a900467ffffffffffffffff168260c0019067ffffffffffffffff16908167ffffffffffffffff168152505062009be8565b62009be77fb80b728a5f386cca4a81807588afdf58b2d5c880ef266d9ca9a1625e36c2f7cd60001b620015a0565b5b62009c167f884cdeeac7e0b4b64a24cb2e435ce790c31886fa8bb34da0218248c579169de560001b620015a0565b62009c447f7676a18e0ed939a80bf384e96137add3fa77742cbf9d5ad50c11e65a37ddba1a60001b620015a0565b6001600281111562009c7f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b826102600151600281111562009cbe577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b141562009d8b5762009cf37f73e6f8b6eac0720a019ba5e7651649459f3bcfc805b0bf270a919201dca5ce1160001b620015a0565b62009d217fce9fc85a7213753f5203eb0573b3cf0756877b019544478864efe34c1b3b127e60001b620015a0565b62009d4f7f3a6ddb196f47d6038cd44f9f2aa91e652bc56c4adf81919244e1160cf314112860001b620015a0565b600560009054906101000a900467ffffffffffffffff168260c0019067ffffffffffffffff16908167ffffffffffffffff168152505062009dba565b62009db97fc1938dafece2f497da4878e6f86f390519a14f71badd7018d05bbd771b3a51a060001b620015a0565b5b62009de87fd74a703b65743152804fde57b9477936709bac93f200266a0c874f00f68f5b8860001b620015a0565b62009e167f4ade62425b730cc3d9934307e6dd2d4bdd0566016ad0b3bddb74ecdd7c61e41b60001b620015a0565b60028081111562009e50577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b826102600151600281111562009e8f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b141562009f5c5762009ec47fbf31ddd5b543927bd0ef9b67abc63c493abcf9ba604a3b6d0811659329b2452f60001b620015a0565b62009ef27f2c3fbb47390a63e8480dc10a1f812e30dc9ff35e6da1f5f938a8e5f99a0f2a3360001b620015a0565b62009f207fd437b53f04f2eae768159aa9ef9297b8aa9da0bfda3f64794cf2955933f21bb460001b620015a0565b600560009054906101000a900467ffffffffffffffff168260c0019067ffffffffffffffff16908167ffffffffffffffff168152505062009f8b565b62009f8a7f970a09f32507fc20ce31ba01dd673be592ce026f938c277e8d00fe1a3717ab4d60001b620015a0565b5b62009fb97f4b3bb5ce59c0fabf6c04027f862efb42b9829d804997b3e02248bdca7f95077d60001b620015a0565b62009fe77fd761a8ae2fa58b5a2f18c798212fc679a730428feead73bd5d0ce68b7571bb7060001b620015a0565b6001826101c00190151590811515815250506200a0277f5dd4a42b222c75d1e428eb6c4b35f04f2bd113ecd8411f5819e7fc766a70485c60001b620015a0565b6200a0557f3d9a19adfed7b393178e734f9811f5d81ebb7bcb69fe6494a824c67d08a1ffff60001b620015a0565b6200a05f62010817565b6200a08d7f9ea16bcb96e0a9de9b060c175dc7c2fe8b5f5e61525f365b9c6473f28534cb9f60001b620015a0565b6200a0bb7fd38e261fd40727a7c42fe1bf95b3833d66ca47a1ad4dc4cabe2bae116d60f63760001b620015a0565b8261010001518160800181815250506200a0f87ffa297a33092e9ed4e9f9108420625c4df6aaf792ccb430ed6bfae9ed56df465160001b620015a0565b6200a1267f20758964a68f56709ca6dae4735572a8b0409348a438e71217269be5aff7ff2b60001b620015a0565b8260c00151816040019067ffffffffffffffff16908167ffffffffffffffff16815250506200a1787f781c85f49e865980d5ccde5443f8d67c88b173fccd41fcfd88327c8deeebd34e60001b620015a0565b6200a1a67f13e970289b93971335134c53f6d1d2070d03bf8d758f0d2f47106aa7f66cb94660001b620015a0565b8261012001518160c0019067ffffffffffffffff16908167ffffffffffffffff16815250506200a1f97f7531ad35f7b2d4c0b5258794b46ce84a80bbff6b6e218cf53c2c9f8fda484f5f60001b620015a0565b6200a2277f59975163f464eeeb625479e1ac69fefdbe60dc069fdf3f3668a7fdbcf224af6360001b620015a0565b82608001518360a001516200a23d919062013c72565b816020019067ffffffffffffffff16908167ffffffffffffffff16815250506200a28a7ff79dc69b33ef26c71a0f0e4c22e3b32b0d48fec555ef342ce164a30156fb696460001b620015a0565b6200a2b87fc4f37235380951ca9ca054719429dc8637b9ef6839bacc749f6012b53cb075f460001b620015a0565b60006200a2c782844362001a92565b90506200a2f77f3ba3bf987e52f312e64c7d8a38330e66a7452fe9df5f2cd5c68453fc0267fbe760001b620015a0565b6200a3257f04abcfc39202ee7f8e925a1decf58889b009d87d313ebc037ccada9eaa0b1aac60001b620015a0565b8060400151816020015182600001516200a340919062013bf5565b6200a34c919062013bf5565b84610140019067ffffffffffffffff16908167ffffffffffffffff16815250506200a39a7f5f50d342d134a0372e97a3d506f49381b0fc757ba586dbdf0a34c9a517e2984c60001b620015a0565b6200a3c87f0b30cae3e0aba407ab3885836c0a30271cac40019b4477beb744255af12d7bc060001b620015a0565b6200a3d482436200f91c565b8460e0019067ffffffffffffffff16908167ffffffffffffffff16815250506200a4217fbd42828f125948f244949752bb7bd81f13c342004ae6d0395abbc686fd8f2c9a60001b620015a0565b6200a44f7f409796b52537f7e5955a4243eca2e28f04a6b9b173d1f8dc0cff3e6fa5ee59cd60001b620015a0565b600060018111156200a48a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b846101e0015160018111156200a4c9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156200abec576200a4fe7f1479adb8afe4794d69ad8c4aba29f82c9434d5cb49325c446cdb8914e534a2e060001b620015a0565b6200a52c7f5e785cb86e6cf030931e0abae28ac5a9197d1d74a1d471389c6132c8a1426b3f60001b620015a0565b6200a55a7f9cb9f632848548094a484a02abacccc92ae4b7137d873b10abc1dec928dd777460001b620015a0565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166354e0d3c286602001516040518263ffffffff1660e01b81526004016200a5bd91906201338f565b60a06040518083038186803b1580156200a5d657600080fd5b505afa1580156200a5eb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200a6119190620125fb565b90506200a6417f4bc14e30120e64e30cba7bca50d3e896091707ea147a7b5661992cd8bb2a8e3d60001b620015a0565b6200a66f7f8c90efca346a7e7e6c062e8b8b0ce22989b5fafa6f4b2931b8ba1de9acad914260001b620015a0565b84610140015167ffffffffffffffff16816040015167ffffffffffffffff1610156200a73a576200a6c27e952324f4814e140cabcf35efa7990248b38511894b9744165b960e5fbfc35360001b620015a0565b6200a6f07ff08c3fe625d5030617c1099a1f7340f398fa7724aa671d26693ebf3214b2670b60001b620015a0565b80604001518561014001516040517f62fe63d90000000000000000000000000000000000000000000000000000000081526004016200a731929190620138d8565b60405180910390fd5b6200a7687fdf8ed198e73d2731e4e4cba189fd6e739f96689a5bd7a83aaab4d97f41bcbb8560001b620015a0565b6200a7967f1d78b2f786405b24946f50eec23fdb02f744334747fc70aa1c3df77a733e30f460001b620015a0565b6200a7c47fd7ff92196792e1d01f7b5abc5d90ab6eb312ad1557d7e9ed898cdada23b3443a60001b620015a0565b84608001518560a001516200a7da919062013c72565b67ffffffffffffffff16816020015167ffffffffffffffff1610156200a89f576200a8287f9b44bd83e8eca3ad20a65eda17edc84d6d78fa7197d71e9191ae9a715fecf1e260001b620015a0565b6200a8567f4924eb45e77693ede686422ac1276125c280861908344a59187e5133c210bcc760001b620015a0565b80602001518560a001516040517fd39dfd030000000000000000000000000000000000000000000000000000000081526004016200a896929190620138d8565b60405180910390fd5b6200a8cd7feff9fb634b58e60e575d86e6fcde63b4eebc1e40021987af76cd8215458c05f260001b620015a0565b6200a8fb7f540d560d4f33c9ebc3a1413252c62c200392d5ca524d5cc0b90ac62088c48dde60001b620015a0565b6200a9297ff130a7a8f5f2ac96c46df65c77cd7320c6730a94417018deae3e00c2ac89f46060001b620015a0565b846101000151816060015110156200a9e1576200a9697ff2cf3a1e53fbc5f992de757ac4db78204164ec34e63cd02b83163da8ad58495060001b620015a0565b6200a9977fffee23fd612c326f9e7ffb1eba576dbb63442700c0bfbe7305d49f5567ff901160001b620015a0565b80606001518561010001516040517f0c8034120000000000000000000000000000000000000000000000000000000081526004016200a9d892919062013861565b60405180910390fd5b6200aa0f7fc305c6179e209930eeca61dc9244f9f8a8a146d49c42cf024cbce3ad46e770cd60001b620015a0565b6200aa3d7f2b32de74fe4815d0fbc38e765c24b7e4dde562265af48d7d4dc64f8e1f94655760001b620015a0565b6200aa6b7fb93a08e573fe5aa186015a0662db7f07457c013e15be8b95b2fb0180ee84b33960001b620015a0565b846101400151816040018181516200aa84919062013cf6565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506200aacf7fba009a7cfabb9654f6a672058ee982660885421484fa3af692100e78e11c58ae60001b620015a0565b6200aafd7f4c158c26e138b54b4cca747b05493d75cb465280964877455f7377751e2e121760001b620015a0565b8460a0015185608001516200ab13919062013c72565b816020018181516200ab26919062013cf6565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506200ab717f8a0dbd432c7164af56d3d52de3c6442ab9b7f19ab0425bf8752048414532323f60001b620015a0565b6200ab9f7f62cc8b6ed4686abb7418ed14c24ee360a06c7480f9987312914b82c76fd404e260001b620015a0565b8460a0015185608001516200abb5919062013c72565b816000018181516200abc8919062013bf5565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050506200ae04565b6200ac1a7f200272346cd8bf268e5ac1ade077dee3a786c8cdc1649fd770a1f4162a4b212660001b620015a0565b6200ac487ff53027ba5569a61a7b0dd189760f02eb5375a43e5c54dd3875002106443fdb7760001b620015a0565b6200ac767f1aa350aeb8445d62dc5c88058de45a7b7342894e086d6c2065aef5fa5781a06760001b620015a0565b6200aca47f2506346691b230a587fcd4c86a7084f357a741648936eafe972d58ab317d28a860001b620015a0565b83610140015167ffffffffffffffff163410156200acf9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200acf09062013789565b60405180910390fd5b6200ad277f70aa5b90757d4acc23e59215cfe919ce03f38707614047805b6c6d0329c05a3a60001b620015a0565b6200ad557f7ec3882214732c029850267d55f6e9888ce66ddc329c5426504a8546904d3b6760001b620015a0565b6200ad837f5073e3fc43b87f60e13e9ad0141b08f32c8bdf1024c04044b30b17430ed3f05e60001b620015a0565b6001846101e0019060018111156200adc4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b908160018111156200adff577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815250505b6200ae327f47dc3c728d9aad722386d6ec2973fb608d952a9be7d9e080e6e8b0a80338bc3660001b620015a0565b6200ae607fc3df9c86daab4b2736d603ea8137f4d4610c210355438f1f2baea58acb1db35260001b620015a0565b826080015184610180019067ffffffffffffffff16908167ffffffffffffffff16815250506200aeb37f3c6d1240d08716de2cd588ced6a63df66c8e327f1a6f3b7a9f08951e82a3dd4560001b620015a0565b6200aee17f87a3cb63d3e6cf2e60f4e1c8950c05c246319424f8dfc460d004fbafa74caa9c60001b620015a0565b43846101a00181815250506200af1a7fcac91423f201d19cb2bb455fa3c41359a907f94ea1e7add26bc2177ad9e86aaa60001b620015a0565b6200af487f3c6e2e4161b61d98b40030ce95b6fe5758fe384e85d378a59ebb767df9f157ee60001b620015a0565b83600685600001516040516200af5f91906201335d565b908152602001604051809103902060008201518160000190805190602001906200af8b929190620104f9565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020190805190602001906200aff1929190620104f9565b5060608201518160030160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160030160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160030160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160030160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060e08201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061010082015181600501556101208201518160060160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101408201518160060160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101608201518160070190805190602001906200b167929190620104f9565b506101808201518160080160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101a082015181600901556101c082015181600a0160006101000a81548160ff0219169083151502179055506101e082015181600a0160016101000a81548160ff021916908360018111156200b217577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555061020082015181600a0160026101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061022082015181600b0190805190602001906200b26b9291906201058a565b5061024082015181600c0190805190602001906200b28b9291906201058a565b5061026082015181600d0160006101000a81548160ff021916908360028111156200b2df577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555061028082015181600d0160016101000a81548160ff0219169083151502179055506102a082015181600e0160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050509050506200b3cf7fb83d52fa82fd630d69314173c7a974e5dac556af46cfd8c85f26d3e07e98362a60001b620015a0565b6200b3fd7f57455c5fb2b987f5651ad56e4cbb938381b29f19c89939aff5e7be3115c6b62c60001b620015a0565b600060086000866020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506200b4727fc43aa74a0d97c9e21159797ae853b80e81addf72f69c572ff7d793d05721cb1060001b620015a0565b6200b4a07f6136151bc8bf0b4879cf3c3cd98f8d64d00fb2c65c28972d86a511d58f538c6960001b620015a0565b8085600001519080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906200b4e0929190620104f9565b506200b50f7f017e4a80834f76ec8277f87e0d93447b45fd5c8f80684f3e2c2af9c32375befc60001b620015a0565b6200b53d7fbd54cfd3a9146eee0afb53d2a3f93d70566c869b4b05917d7c64b52af7863ac160001b620015a0565b60005b856102200151518110156200b6e8576200b57d7f35e7bdafa21c98c3c025a9afb78d6402f75ceeed14ff3bde2746704013105dd060001b620015a0565b6200b5ab7f4d492e388c813287c07aa5b10b2794f6f2d839339ba881d9945163324ed5261560001b620015a0565b60006008600088610220015184815181106200b5f0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506200b6627f3aefa0f9b9aea7a8d0b1db248d032b2c3e80d088583db7730fe7714ea77cbf0260001b620015a0565b6200b6907f9a16d308a298ce4317b1d35918ba2677ac974857f1f739996ec6c3327de01c2e60001b620015a0565b8087600001519080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906200b6d0929190620104f9565b505080806200b6df9062013f51565b9150506200b540565b506200b7177f106ea2301b67d3da118c5314e4651798ac4273dbd3ec00bd792618b97653be8e60001b620015a0565b6200b7457ffd68b719fce5c674a6005a2f4291e039f2bb8fcfe46314aac96420d9a1c8d23160001b620015a0565b60005b856102400151518110156200b8f0576200b7857fc63b5d6e797490c7b610cf05c6c5329548b54f954483839814174f5531ee32a460001b620015a0565b6200b7b37f6a12cb375ef74befdac946b195a1907cfeacd392aca6ed522d26b3c636fb7cae60001b620015a0565b60006008600088610240015184815181106200b7f8577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506200b86a7f38b264ea93f0b51eec711e14ed38d493733c7771527277519a76d0137f95da5660001b620015a0565b6200b8987fab578f961e96f7e0d9e980b428e1caeb76293eb2be36024bebda0027ecc3857260001b620015a0565b8087600001519080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906200b8d8929190620104f9565b505080806200b8e79062013f51565b9150506200b748565b506200b91f7f7bb3abe63085651b900427d704548b9f6f92dd598300edfabe476ad6da0a682c60001b620015a0565b6200b94d7fb9a2bcd169436f695dc40128957614ef38c1cb883d0bd14fed8cd4244aed30b260001b620015a0565b6200b95762010900565b6200b9857fd93403149fe66c5fa56f3896de7e3c89509ab34a5b86528d0f104eec1fad540e60001b620015a0565b6200b9b37f452678ef4f605b5bab621803e787064ce3668351cc6a6b7f25ed6d28664d060660001b620015a0565b856101200151816000019067ffffffffffffffff16908167ffffffffffffffff16815250506200ba067fbcdff6b64b4dc066d651baa27bdd3dbcd28f299895efbb44c022a28ed28a801360001b620015a0565b6200ba347f450ce935f17fe5cb13f764fdba8098a418b4dc096fe52b15171ea5e5130604cf60001b620015a0565b6000816020019067ffffffffffffffff16908167ffffffffffffffff16815250506200ba837f843b7d9b779e805acbf276a29187d459bc1320f792004023e95bc74f0773216460001b620015a0565b6200bab17f68dd9b25246f7382bf0c9c678f0404b0f1f04695784718dfae534f4be41629ef60001b620015a0565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bb4e4e428760000151836040518363ffffffff1660e01b81526004016200bb14929190620134da565b600060405180830381600087803b1580156200bb2f57600080fd5b505af11580156200bb44573d6000803e3d6000fd5b505050506200bb767f1c017d94f61fad97b863f4014cf61e27747882b50232355319af1e507309e03e60001b620015a0565b6200bba47fa8b25ac94132cb5de833164098b3f8ab60595b902d380942ca5cf92c35ec00fb60001b620015a0565b7fa08fe92f4f83b40431d08f9f130fd22b658ad78cf83027611d629d83427f0d1860004388600001518961020001518a602001518b61014001518c61028001516040516200bbf99796959493929190620135b6565b60405180910390a1505050505050565b6200bc377f60f4ad7a32a98583407569415f6bb676c8e5600b3263cec10bdecf3b66394a8560001b620015a0565b6200bc657fc8319b01485dac020e8b5b51bb2d5db722b7cd3b93ddebb01d1ca09f852af6e260001b620015a0565b6200bc937fb7a256da74bfb28af42e2f39baff2c27316e7d840a84f93d210ca225980d829660001b620015a0565b60005b600b60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490508110156200bf48576200bd0e7e6e027161f2f321aafb51a5025906fec65d3f71c010dcc7a7fec027a0ec5ce360001b620015a0565b6200bd3c7f85f4a5bd4455ed052654f8e0324e0ec109507a45d44dff13dce52183f2a34bfd60001b620015a0565b8180519060200120600b60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002082815481106200bdbc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020016040516200bdd5919062013376565b604051809103902014156200bf04576200be127fb3da975fd0934fc57a83f19c98f177026fd935cad5aeb530a5be48fafcbf50f360001b620015a0565b6200be407fb2835011e520bff0614d05b41f22229df6c8c61385e5d36f53d1ccdd6e7367f560001b620015a0565b600b60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081815481106200beb8577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160006200bed0919062010619565b6200befe7fd27e87e69fd5b31e928274138ea3dea87dac3c5e1b0069a957b45dd9148d7f3560001b620015a0565b6200bf48565b6200bf327fef72770d7d84b23281a09f97cb36d2ddf24ed749080be126f8fbe4c28ffafbf860001b620015a0565b80806200bf3f9062013f51565b9150506200bc96565b505050565b6200bf7b7f99aeff814146c54ea04bbe440a73f62f7199f4c3cf83831eb3a649e223b280a760001b620015a0565b6200bfa97fdbe9861e8e5762841e618df500e933e87ee1212151579015c41674ec784d5ec260001b620015a0565b6200bfd77f122c6ff6b646163db8db3492449180ae21dea711655f77250de30bed8d03817260001b620015a0565b60006200bfe882600001516200ed92565b90506200c0187f4298f47ad90cf278fc0374a16b4330b66bc9212a79aff3673af3946f3b116d4a60001b620015a0565b6200c0467f7079fc557e49f9573c55f0c78a14cb3a3a05e63227cc2dfad9e267c01a570c1c60001b620015a0565b8160200151816060019067ffffffffffffffff16908167ffffffffffffffff16815250506200c0987fb51d7e6f39d45df623a1f53c6b914bdee6665563ec43712d7e49e4ddd6990b1e60001b620015a0565b6200c0c67faedba580f299a7c1244118fed3a9c7387a4ccd12a95d3d5dedc3a24ac462048260001b620015a0565b6200c0d18162001c26565b5050565b60606200c1057fdb6458cc05e08b336e0790c1451ea98da8c55f399b2a316045d0549c396e9e3360001b620015a0565b6200c1337fc65e1f8e6f042b7c7aa93b55d0fc464984a780f68ff21cd64a56b899352f7cea60001b620015a0565b6200c1617f9a304036fb6719d5f8d0fffd6370a620c2d130c7c28d326e6f5f84460b25efb960001b620015a0565b6200c18f7f641159e61bdf2c900ee6a738405853a032863463826fd94d402662f2d8240f5060001b620015a0565b60008251116200c1d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200c1cd9062013657565b60405180910390fd5b6200c2047f6277ba5a34555d5469f97c1f931d34a369e81cb48854a651be60df154855812860001b620015a0565b6200c2327ffa16a54cf63b8cd89c7675386925c636f2c47c5c746221bc7c152e8354fe069160001b620015a0565b6200c2607f5fb8e95eb88d48f32ef01617185defc615c22c6c828c0419a074d8582cd0620a60001b620015a0565b6000825167ffffffffffffffff8111156200c2a4577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156200c2e157816020015b6200c2cd62010682565b8152602001906001900390816200c2c35790505b5090506200c3127f51ef10fbad4408679bbbf1d905136fdb6618b47ecff0aca9649216229ca580b460001b620015a0565b6200c3407f5ab2dc6e80adc23371dac0b7a9dad1ef7d36e49d92922bc361e29328f47e641060001b620015a0565b60005b83518110156200d4ac576200c37b7f8dd1976522cf158a4277cb8ae8a6399490fac084fffc5ba86a51f33069f1853960001b620015a0565b6200c3a97f0a4499cf04fb742c3d3042d95bae68c89ebb31b5a0a4745b16bc7500a49c820c60001b620015a0565b60008482815181106200c3e5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015190506200c41d7f9d0b34ee6cfe5f201aa5540590e09a118d2fc7b5a32143010983702ad120ef2360001b620015a0565b6200c44b7f4f699d550ff91e67e4483bb58a04563e94e1c0d6500ae2fa3f41aa785117b93960001b620015a0565b60006006826040516200c45f91906201335d565b9081526020016040518091039020604051806102c00160405290816000820180546200c48b9062013ee5565b80601f01602080910402602001604051908101604052809291908181526020018280546200c4b99062013ee5565b80156200c50a5780601f106200c4de576101008083540402835291602001916200c50a565b820191906000526020600020905b8154815290600101906020018083116200c4ec57829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820180546200c57b9062013ee5565b80601f01602080910402602001604051908101604052809291908181526020018280546200c5a99062013ee5565b80156200c5fa5780601f106200c5ce576101008083540402835291602001916200c5fa565b820191906000526020600020905b8154815290600101906020018083116200c5dc57829003601f168201915b505050505081526020016003820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600582015481526020016006820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016006820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016007820180546200c77d9062013ee5565b80601f01602080910402602001604051908101604052809291908181526020018280546200c7ab9062013ee5565b80156200c7fc5780601f106200c7d0576101008083540402835291602001916200c7fc565b820191906000526020600020905b8154815290600101906020018083116200c7de57829003601f168201915b505050505081526020016008820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160098201548152602001600a820160009054906101000a900460ff16151515158152602001600a820160019054906101000a900460ff1660018111156200c8a8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60018111156200c8e1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600a820160029054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600b82018054806020026020016040519081016040528092919081815260200182805480156200c99e57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116200c953575b50505050508152602001600c82018054806020026020016040519081016040528092919081815260200182805480156200ca2e57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116200c9e3575b50505050508152602001600d820160009054906101000a900460ff1660028111156200ca83577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60028111156200cabc577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600d820160019054906101000a900460ff16151515158152602001600e82016040518060600160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250508152505090506200cbb57f4bb39638ccc1a9f38d690908c14e1296d5a7c53e49e1afc4c6e7b5f8de71ed5360001b620015a0565b6200cbe37f21b0024845eabfff4136cae851052a4550ee32d964bf9521863a54a959d189eb60001b620015a0565b600081600001515114156200cc8d576200cc207f8bfb80018024877eecf7365b3f2bf626355fafd59f0514835f7d44600a0f377860001b620015a0565b6200cc4e7feb48c84303c06de35c0a22c1fb5eeefa35531efd3897a76346caeed5a171e17460001b620015a0565b816040517f6c5249c10000000000000000000000000000000000000000000000000000000081526004016200cc849190620134b6565b60405180910390fd5b6200ccbb7f2f15e999569c7e08b0701a041b321e49d8ba0b7e2c2bf79056227601e913a7d260001b620015a0565b6200cce97fba98946e89854548220573adb3d7e1d3847c705c16e671c28499dd8b6780ef1060001b620015a0565b6200cd177f2da02bc1ef8fc815c8e1b07922ea78f8f53c46e7e2a4c0041888a666e38521e760001b620015a0565b6006826040516200cd2991906201335d565b9081526020016040518091039020604051806102c00160405290816000820180546200cd559062013ee5565b80601f01602080910402602001604051908101604052809291908181526020018280546200cd839062013ee5565b80156200cdd45780601f106200cda8576101008083540402835291602001916200cdd4565b820191906000526020600020905b8154815290600101906020018083116200cdb657829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820180546200ce459062013ee5565b80601f01602080910402602001604051908101604052809291908181526020018280546200ce739062013ee5565b80156200cec45780601f106200ce98576101008083540402835291602001916200cec4565b820191906000526020600020905b8154815290600101906020018083116200cea657829003601f168201915b505050505081526020016003820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600582015481526020016006820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016006820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016007820180546200d0479062013ee5565b80601f01602080910402602001604051908101604052809291908181526020018280546200d0759062013ee5565b80156200d0c65780601f106200d09a576101008083540402835291602001916200d0c6565b820191906000526020600020905b8154815290600101906020018083116200d0a857829003601f168201915b505050505081526020016008820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160098201548152602001600a820160009054906101000a900460ff16151515158152602001600a820160019054906101000a900460ff1660018111156200d172577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60018111156200d1ab577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600a820160029054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600b82018054806020026020016040519081016040528092919081815260200182805480156200d26857602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116200d21d575b50505050508152602001600c82018054806020026020016040519081016040528092919081815260200182805480156200d2f857602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116200d2ad575b50505050508152602001600d820160009054906101000a900460ff1660028111156200d34d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60028111156200d386577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600d820160019054906101000a900460ff16151515158152602001600e82016040518060600160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050815250508484815181106200d489577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181905250505080806200d4a39062013f51565b9150506200c343565b506200d4db7f436d834e84d366d7ed3fdab4b2731aea1fcb52cb1a25922f0107770427dff21a60001b620015a0565b6200d5097fd38dc8ee092fe161bb932e0e1b507bae35797ae4513841ac7488c29e170e1a6d60001b620015a0565b80915050919050565b60006200d5427f438b7c63a5809859924121b74371bffdb86715a85d11f4142fb04730466f503460001b620015a0565b6200d5707fa257f053309ee13448581c267e478ec419ed24ebb37c69dc325c0ee4d9f23a4e60001b620015a0565b6200d59e7f44a6cdc69e1afba364d68d3c73a49210c19debc0dc4ae6ecac73a44eff72e0eb60001b620015a0565b620fa0008284606001516200d5b4919062013c72565b6200d5c0919062013c3a565b905092915050565b6200d5d2620104ba565b6200d6007fe53554de5b2c6dcbb32fb5ddc14a029494fa514fdf498d2d40865dcda916b45f60001b620015a0565b6200d62e7f9c13f4ff734a6f7eee3bd47e6f495b0056c856ddbe4a0cbe4577a966fa7b346860001b620015a0565b6200d65c7f99daf98c674c78c79c2c0af2516ffd710cc5f198e098930c864aafd9adb920a360001b620015a0565b6200d666620104ba565b6200d6947fa062ea5685fee4aab19e0e7ad8be4acfde38de148ccf5adca7b9be7a93c56e0b60001b620015a0565b6200d6c27f9efaaf89811b62ce32fa7a3db9890652fca00f1827de21afe479768a74b7703760001b620015a0565b60006200d6d288888888620019d4565b90506200d7027f217985ad4a7266d851008a41b4e34e40d43570ab7c5f8caac3721c7731a09c4060001b620015a0565b6200d7307f7061159c8da20a70dd4a7e7602c2486f7f9284afaa07fd8c66c70da0aea843cb60001b620015a0565b60006200d7408988888862008714565b90506200d7707ffc760423793427029ac6e6caa5c91221b38bfe2ddc75098cc682e5e7a4e59e8560001b620015a0565b6200d79e7f4700e78bbb56c6c3cb3949c24c4a48600d4f0c5f990b0769b30e5a0bfee4d40860001b620015a0565b81836040019067ffffffffffffffff16908167ffffffffffffffff16815250506200d7ec7f79cb8ddeb2a6319b9a71e06be79c27d2ce5f5c4a256aa0e6e051ebafbe71bc1460001b620015a0565b6200d81a7ffd56fcbfeae8b7f81d992296d272259147969496a822f0c75311a11c9508a19f60001b620015a0565b80836020019067ffffffffffffffff16908167ffffffffffffffff16815250506200d8687faeeab0f86657b027adf6135d4b646cd8d1997aea8b834ad8f1e9820161e978a260001b620015a0565b6200d8967fe66dd2e4b236471e68dfa46856642dc70c0591ebe4d54a3bf6b66d1957f6638860001b620015a0565b82935050505095945050505050565b6200d8d37f944bf0004173e26ccacff736a32e72e30731cc1357af87c6a64b068ba0aac50260001b620015a0565b6200d9017fe20adb1a3dc2e3ad849b255a8e79fb7e371bf8c976a3f2c1c3e156f617f9e96060001b620015a0565b6200d92f7f931025eaa38d4100f9536d325d95da9bf06f411f760ceea4409a79f889a2ea0060001b620015a0565b60006200d93c826200ed92565b90506200d96c7f40e7ba2dcc65eba1ed28c934e69e2787a99b2e2b6866ba0787afe8c6fd0a144f60001b620015a0565b6200d99a7fe04029617dec975d0c5fa51940a0213540119e2d4d26a6a1f75e459ea5339f8260001b620015a0565b6000600167ffffffffffffffff8111156200d9de577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156200da1b57816020015b6200da0762010682565b8152602001906001900390816200d9fd5790505b5090506200da4c7fff7ef79382f5411b6b61e1e0f19eb62a8f4238076ef00b9b5bceb015b4e8db2760001b620015a0565b6200da7a7f8d385dc8fc6b11adc66c09365ddeaa5a54152b265b49cc2418b8d79e967f443560001b620015a0565b81816000815181106200dab6577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052506200daef7fd24c944d706b4662785cc657c7efcce90012d00964f546832cc5c30973e261c560001b620015a0565b6200db1d7f48950c13e89965881ae3d403cc658eab4310d04b567409fc9f7dab142bc56cb860001b620015a0565b6200db288162002451565b6200db567ff4f33d55631a58ada81a0821224271f35eecf482a489cf712d448c7dcb4b341760001b620015a0565b6200db847fd4684b5b53cc0618e1c7bd5b83df1555772d0700830a4cd421438a7bc284512d60001b620015a0565b7fb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f456001438585602001516040516200dbc0949392919062013562565b60405180910390a1505050565b6200dbd7620104ba565b6200dc057fb2465e4b5c6ee98bb70addaa69d1bbf2b2cf1f471a2cb19d079fd9e4162574aa60001b620015a0565b6200dc337fddc24047eb9ac7f2c488986ef60d8f5c8110a4bb81102134217573da52dbb84960001b620015a0565b6200dc617f16528a43ed8102cacbf9d96a60d43372852b65e42499dc2ea560606046ca737560001b620015a0565b60006200dc917fe582a221657b948230177c6476bc47a665d7d805eb4f2b9ac77278b7d201922e60001b620015a0565b6200dcbf7fc7ffd4f66bae2be4ac58ad6a49398278b0005c85dee7162a0c19f645d2ae2be060001b620015a0565b60006298968090506200dcf57f0d60a796501a28fc32c78678cf4ffa493da5c07cd50f630b1b64f2e49f51118b60001b620015a0565b6200dd237ff7a9c29e816451d3dd3f705ac9c22f9a282f17dbc7703ba04f3a89db001a11f260001b620015a0565b60008661018001515111156200ddd4576200dd617f9b5f4500b980ab024a912e1192232e1d34875790084998618cf410eae20b2f5160001b620015a0565b6200dd8f7fb06787a1d86be30928c28df7c7bc75d6f3d8238424ae538d0278bd41abc9063460001b620015a0565b6200ddbd7f863a0eabc2f3a141356a9f08e21aecbd4eba8d0ae5d7f5c3414fa687e451bb5260001b620015a0565b6004816200ddcc919062013c72565b91506200de70565b6200de027fc6f44ff32a0bcf7ba8a4716a3e0abd272b463afd2c9d02ac9fab5d71758d855d60001b620015a0565b6200de307f9fd52ccc2f41147a23410e80ebe5727d77a796ee95e390d52f61b13c530f758a60001b620015a0565b6200de5e7f1f640d5d0ba9262d20dd82045622636cf1debcc40e73c68b5b561235f034c38d60001b620015a0565b6003816200de6d919062013c72565b91505b6200de9e7fce8c73828107763f824cc7a2fdbcecd6653e252e8dfbc6d1750e3a6f0b11632a60001b620015a0565b6200decc7fbc43b330e0c207bd19776067478d6493561f06b7f915a15d7dc56deb818aac5360001b620015a0565b6200ded6620104ba565b6200df047fbe192d68aab4df86ab4b98a7a9335eccc0bf203d6f2d5aa8a24d2c3b654d023260001b620015a0565b6200df327fc030bd9adf8a5bef1541ef5faaeff6ee3f884feb14a8a1aa840cac16b7ea3c5f60001b620015a0565b82816000019067ffffffffffffffff16908167ffffffffffffffff16815250506200df807fe5d0de1b3c8958b76ad4ec361afcf3515314d533038b7c2e21ebd0a161a94cde60001b620015a0565b6200dfae7f1952e3ea2ae72b328ac1d20e4581218c2541a1fa56e39850bd51ec7597daa21160001b620015a0565b600060018111156200dfe9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b876101c0015160018111156200e028577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156200e0c5576200e05d7f0ebb73d09e14ff731e20626c777eab77e923a0c40c848fb2fae3bf5ef692ead760001b620015a0565b6200e08b7f0b2f01f68a1e0fbe450eae0aa512e394a8204bfd805f4fa6ec4d13863b518c7860001b620015a0565b6200e0b97f74e6ecd310e70b198143f17816b79dbe720368983000ecfd4f2d56f88ec2b40060001b620015a0565b8093505050506200e2c4565b6200e0f37fbd0472ec4cc9554e7069b20ea2f0873d9a60da0621ce9a6beb0cc37b82522e3760001b620015a0565b6200e1217fed9399dd85d4438fdeb2365229bdd5d7a427a130cc1492e9ce92a6487648c34960001b620015a0565b6200e14f7f8d6cda09b6ac598965523e960cf0efdde48898dad6bda53b86a403ee0ea0ff8360001b620015a0565b60006200e15e88888862001a92565b90506200e18e7fd54804b2ee6d59a2aa938e306af5a218ae12e0ef286fd1e4f7a45ad87c09ce6160001b620015a0565b6200e1bc7f11a1fa1ab968cb207dac44ba78e6a32101fcc6f7e556578e73614df995b36ff160001b620015a0565b8060400151826040019067ffffffffffffffff16908167ffffffffffffffff16815250506200e20e7fd62d8bc4a654258f1479e5b6bfbfe30d65376a5b7382613032f4c306d57be34f60001b620015a0565b6200e23c7f1769590e259297417b0378a986ef120be903c8c356df373b2c857a491bf8597460001b620015a0565b8060200151826020019067ffffffffffffffff16908167ffffffffffffffff16815250506200e28e7fb219c0b9d4050add9615ce47cd042f593150936f4b0a36de6db6b4487ff22b1160001b620015a0565b6200e2bc7fb122c90c16be5246fa626aac8e1e18ec51dae63b99f4735df8540e400957216060001b620015a0565b819450505050505b9392505050565b6200e2f97f5898ff48267bc362603a052d3f8cbcbc76704e71b12ad683cab294a0aab838d260001b620015a0565b6200e3277f90bab23fc9a94161cbb3ad896fa9587e398713cb77d4a16c3c03710bb7d8e8e960001b620015a0565b6200e3557fac31b44f3ca53a31ce9b1469dd4c82a6e995303d8021db80c646b6d38b821e3060001b620015a0565b80600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090805190602001906200e3aa92919062010453565b505050565b60606200e3df7f387fc187174dc5bf3f189e4b690132eff002c22d10b933ab9f8db4142efabd8e60001b620015a0565b6200e40d7f9ae13d2bb882495c21f1b892685a4983901d54a519db1054bfc7434dee2e0c6660001b620015a0565b6200e43b7f90f6755c20e8557b6498a716e740a280e3060e0378ae5a008c12d4c2882a5bb660001b620015a0565b6000600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156200e5525783829060005260206000200180546200e4be9062013ee5565b80601f01602080910402602001604051908101604052809291908181526020018280546200e4ec9062013ee5565b80156200e53d5780601f106200e511576101008083540402835291602001916200e53d565b820191906000526020600020905b8154815290600101906020018083116200e51f57829003601f168201915b5050505050815260200190600101906200e49c565b5050505090506200e5867f4b3f6eafc480a7379146cca4fa23335282a7215c27f64ead52865de3b9e1d13060001b620015a0565b6200e5b47f66d109be6d09d9150ece19d0312ae36e25e7f8bed259900be86106b87048644860001b620015a0565b6000815167ffffffffffffffff8111156200e5f8577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156200e62d57816020015b60608152602001906001900390816200e6175790505b5090506200e65e7fd89e51d83ac654caeab43c3e7fb336cbd2347c365cfe6f2a7814dd5e528ed60a60001b620015a0565b6200e68c7f068088e1e0a3164cbec40d00bb33366f9028fe61a5e8c43c4d0d1e71fca5232160001b620015a0565b60006200e6bc7f57f546cfeb2ecdb41382d9a94561376beacbf13961fae380ee3fdfe517e8462660001b620015a0565b6200e6ea7f202378abc95088401786fd8d5f9f94b6b3eabd3abca05b67e24acc89e12704a060001b620015a0565b60005b83518110156200ed2a576200e7257fedceecd1c56994214bb7345439cbdb617d7bbea4a3a68c505d550254038406af60001b620015a0565b6200e7537fa390cb4e8bb8bdd8717b87ba4a4044cb89611814d4b246654618625afddd2dbd60001b620015a0565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663977fdfe28684815181106200e7cd577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b81526004016200e7f39190620134b6565b60006040518083038186803b1580156200e80c57600080fd5b505afa1580156200e821573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906200e84c919062011f8d565b90506200e87c7f0e9a73e50c0a93382778abeb582b3cca3089c4cd1ff82036f4f16db53a87325960001b620015a0565b6200e8aa7f679f353d2641982e84b49a838e61c4b9a83eb82320777752dd0583f6e8e7308460001b620015a0565b60006200e8da7feb6291909339b89d5da54b588b5ff89d166078a702e0e4adf72532971fa7c4a860001b620015a0565b6200e9087f939f69ec31cd9e529374f946d88b4041c2228e2338fcd88c25e4fcf8d4fa343260001b620015a0565b60005b82518110156200eaf1576200e9437f91005e23bb9c69a1b427e5670bd2721e990fb07187aa595ff7f14ec603b362c160001b620015a0565b6200e9717f2c84a07270952b2ca79bdd33d571914fd8fa6b4434acd41bc200bad73b32a1da60001b620015a0565b8873ffffffffffffffffffffffffffffffffffffffff168382815181106200e9c2577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff1614156200eaad576200ea197f0c01464459099109af6ddb252e27b60276a1f77b4ba0ea40a511d08950bc266260001b620015a0565b6200ea477f4a16a3888408d130267800974cb3f5e6a5e9b6d79655204ed3b1cad92ebe378a60001b620015a0565b6200ea757f7912575e61593fe96f1d13272db936b7df89b855d08af2740fb192877759ab4a60001b620015a0565b600191506200eaa77f61c2f39b6dfbebf6aa4b08795f43da369833f08168f16d2bbea18e98037a48d360001b620015a0565b6200eaf1565b6200eadb7f1ae1b38b7b7495121937c55fe8d91aa8e3d6383075bc0ae87d1ba60c7e7482f860001b620015a0565b80806200eae89062013f51565b9150506200e90b565b506200eb207fb53f9b6348cabac9b3a8b7be41201dba84c1c50549ac3c9e361701bb37539dd560001b620015a0565b6200eb4e7fecbc1df80962fef1814fcc10acd9307f01ca060a681b78ce0be586e07f71410a60001b620015a0565b806200ebb8576200eb827fbe3619b2c3d340abd8740482d1d5bf3ecc7a3217e1cf7bbf928f12aa0707e7f660001b620015a0565b6200ebb07f77b0c12f7bfdd306b6f1bc623eb1784c042d33b3caa93bea46fc291388d2763960001b620015a0565b50506200ed14565b6200ebe67ffa626994d9dad6c2993be845812fff6b38adc278e284f8cfcdda4bf1d28bab1f60001b620015a0565b6200ec147fb021bc26b7f58e22c635847dbb9505287fd72a27a30358a09a415e6d46a4ff3460001b620015a0565b6200ec427f46f9a942d369ce077ce569a431899e1e51eb97ecf28fadda98bd81f4c91ce3de60001b620015a0565b8583815181106200ec7c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151858567ffffffffffffffff16815181106200ecc8577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052506200ed017f7917f4dc92fe5bc718f1a0c0b3e3e64658e4d5c1ffaff204d6b9941947def0c860001b620015a0565b83806200ed0e9062013f9f565b94505050505b80806200ed219062013f51565b9150506200e6ed565b506200ed597fc5484101124372f011979c6677f1a3f14cd94c01457fb7b80441a411101b851a60001b620015a0565b6200ed877f4725ddfbc6d7d212581bb45f78d106d3dabb63258454133b4376ccdaecb2984160001b620015a0565b819350505050919050565b6200ed9c62010682565b816200edcb7f57879634bd2dfa3ddbf06b3b735c29ad8e51016bb0c4f5d407ba7a411c0ef36660001b620015a0565b6200edf97faef3e50a0b73da2aafc5f30a7b84e5a51881148024382c3586930fa1611d05dd60001b620015a0565b6200ee277fed3311f93163521dac2206de7278dfb3d29d54811c042eacbe492691aa58815760001b620015a0565b6200ee557f8cda38560b56ef3acd37f866324212bd50f30886763063555be10aa0659f8cac60001b620015a0565b60008151116200ee9c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200ee939062013745565b60405180910390fd5b6200eeca7fcaf13f8d7d3adcfa35498a36d395538f7986f622c21f900dc9f6f4e3cd2da8a060001b620015a0565b6200eef87f9f485013c6bd038643674c2e83960d78f4624f205ad8fcca46880105f03e8fc360001b620015a0565b6200ef267fc044375fadb95d358cee357feb34383d9cd5405865bfc6ab5f36bef1efc31b3460001b620015a0565b6200ef547fcf4a4be628adacfb14428def7410ad07ff6fc3ad8578eecbc4b7f3597fac52f160001b620015a0565b6200ef827fb10642d10e46a8f2b6cdde9473b70fc6091b1997274ce55d46926208737522ac60001b620015a0565b60006006846040516200ef9691906201335d565b9081526020016040518091039020604051806102c00160405290816000820180546200efc29062013ee5565b80601f01602080910402602001604051908101604052809291908181526020018280546200eff09062013ee5565b80156200f0415780601f106200f015576101008083540402835291602001916200f041565b820191906000526020600020905b8154815290600101906020018083116200f02357829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820180546200f0b29062013ee5565b80601f01602080910402602001604051908101604052809291908181526020018280546200f0e09062013ee5565b80156200f1315780601f106200f105576101008083540402835291602001916200f131565b820191906000526020600020905b8154815290600101906020018083116200f11357829003601f168201915b505050505081526020016003820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016003820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600582015481526020016006820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016006820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016007820180546200f2b49062013ee5565b80601f01602080910402602001604051908101604052809291908181526020018280546200f2e29062013ee5565b80156200f3335780601f106200f307576101008083540402835291602001916200f333565b820191906000526020600020905b8154815290600101906020018083116200f31557829003601f168201915b505050505081526020016008820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160098201548152602001600a820160009054906101000a900460ff16151515158152602001600a820160019054906101000a900460ff1660018111156200f3df577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60018111156200f418577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600a820160029054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600b82018054806020026020016040519081016040528092919081815260200182805480156200f4d557602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116200f48a575b50505050508152602001600c82018054806020026020016040519081016040528092919081815260200182805480156200f56557602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116200f51a575b50505050508152602001600d820160009054906101000a900460ff1660028111156200f5ba577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60028111156200f5f3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600d820160019054906101000a900460ff16151515158152602001600e82016040518060600160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250508152505090506200f6ec7f6a91e296095e9591810a3946baffc11ba323b2de00c1b0db24e4bb1bd927c57160001b620015a0565b6200f71a7f925a21be4c480dcb9911569499d901357d0f1d6e69e5712a161402e33dea108460001b620015a0565b600081600001515114156200f7c4576200f7577f75259cb50e4defc4f3ab59cceda4487f4f84be9296792b7cf5ccd438dd74a2a860001b620015a0565b6200f7857fd9a98aa7f5271ed6c5339c9f92e613610cc8a42f138b163e90950e425cf3303660001b620015a0565b836040517f6c5249c10000000000000000000000000000000000000000000000000000000081526004016200f7bb9190620134b6565b60405180910390fd5b6200f7f27f52133f711aa408c13cc50008f42e1bfc4185b0e8917aad8d80b2a97607390c9d60001b620015a0565b6200f8207f5f105ba679e5fc64a3480e95b5b824e38d53d3212f88691274338a38c56ad18b60001b620015a0565b6200f84e7f5ae2833e3301c0662eb171913d4fa8296ef4a90b05ba2ee257cc34876a5e128960001b620015a0565b8092505050919050565b60006200f8887f1fe9fc1472a69e0466fd47bfa10832b522e056154c6e360255eae6696802263560001b620015a0565b6200f8b67fe0925de2bee54b459f5437ea8911810de76fd7826191d4d4e30e358406a5d08960001b620015a0565b6200f8e47f1f139fb914682594f7d735af39cffe774340f9bcf8e119b1fbf6ffa116380d5660001b620015a0565b620fa000828486602001516200f8fb919062013c72565b6200f907919062013c72565b6200f913919062013c3a565b90509392505050565b60006200f94c7f8c6d1ddd2276911fff5706bdd5e97f5a0a3be042836862980245cccff080280860001b620015a0565b6200f97a7f7393aa05998eb8d0019286df986114f0c27a37c68238cf531fcb29d468e9c0b160001b620015a0565b6200f9a87f9e7875b337bc99fa44106d38412015164105b797494c6bb2ab24b3b6a9b831ff60001b620015a0565b600183604001518385608001516200f9c1919062013cbb565b6200f9cd919062013c3a565b6200f9d9919062013bf5565b905092915050565b60006200fa117ff6b7b5fae9dedda0f487a6d9b50a912b42669f8ebde0237969a23c7dc148a40d60001b620015a0565b6200fa3f7fcd0f23b0fdd0c6c6f86891931674f7e5253a4cc6f22d237e348124e16b82963060001b620015a0565b6200fa6d7f73a73f34f582d6c37316641eb60b79129fa0301a8de5329d20bcd22c48e57e5860001b620015a0565b82826101000181815250506200faa67f1502a0f78e02d97146ad524a6ad70324b9c1cf46c49132b1ccff31981ee7322b60001b620015a0565b6200fad47f27998aa87b098b76afe9d84f1f1f84f6ec28f705683a78064aa9d97467fd6bd560001b620015a0565b6200fade62010817565b6200fb0c7f4391215a194a1e339ed3d899f0237e4e25e5567cd7267d60bcaeb9480c44d94460001b620015a0565b6200fb3a7f902ed3e3ffe158a6577b5a678ce1f61d7cc7aee32a7a04f190f5e1926f1de72b60001b620015a0565b8261010001518160800181815250506200fb777f9402a71f36be7735c925ab19c4b222d6cf4c5362136f215dbe81585132b4a10260001b620015a0565b6200fba57f6cca6dbb597602c989f9385928265941dd0fdeb19901a94984091d1027d132ee60001b620015a0565b8260c00151816040019067ffffffffffffffff16908167ffffffffffffffff16815250506200fbf77f9d61adcb7d8de38326080451d6b6062c3e9883bd703eef940123b6328178508960001b620015a0565b6200fc257fb85fa0c2ba1f3dfc45f9ac728996b9b5115769cc119e59d06cbb7448bde4705560001b620015a0565b8261012001518160c0019067ffffffffffffffff16908167ffffffffffffffff16815250506200fc787fa8f23140eeff226f2fa651cecb4bf71f4bd4eeb04377777ee4f07b23206efcec60001b620015a0565b6200fca67f17d4ce84ab38db9751999ef5804c7d9880a7b4f8f6844a0678dac15ec8ad5d8660001b620015a0565b82608001518360a001516200fcbc919062013c72565b816020019067ffffffffffffffff16908167ffffffffffffffff16815250506200fd097f2fa0fa6bdd49e88c619efa12bac96ffea381dc91c102493b972439716953cbb760001b620015a0565b6200fd377fd56a2cd8bcddae1bf3be685feca29b6994c2f94751d90c5a45fa58a1258c906860001b620015a0565b6000836101a0015190506200fd6f7f019c0f8331752afe2515b71d3e6cd862f51d400c32358d8729baa76c9c858c7a60001b620015a0565b6200fd9d7f4def059c2fd8442a85b5a8206516cb80bf8c283fa15f36166ccaaa1431c312e660001b620015a0565b60006200fdac83888462001a92565b90506200fddc7f5026f785df6e533659c03b246091f1feb2ebd386d2415d220829115bc16b1d1660001b620015a0565b6200fe0a7f0c1d72ae9f3a74452e2c8d3f4f17594377e401afb2fe7ba1f6a5c738bce8ab8160001b620015a0565b60008160400151826020015183600001516200fe27919062013bf5565b6200fe33919062013bf5565b90506200fe637f76b71838af0234fcda6c36fb8da83e4b05c436bcc25e4141ae23db9e5e49f70f60001b620015a0565b6200fe917fc12e84a892cbb1acff726b76b238bc62b3a1f80a1f84289b44854662b835b3f660001b620015a0565b85610140015167ffffffffffffffff168167ffffffffffffffff16116200ff4a576200fee07ff7f52db577dfcc483aa2ca8ce9f28ebec6bb73fbef4a477c06460d3a3265dd8760001b620015a0565b6200ff0e7f9a88d05f70bf6088172eabb953cfb548175ec95a760357fae58ee72638477ad360001b620015a0565b6200ff3c7f64268433e8261a1e559e661e622d7f9e985588f27c64be78866cf8571392a56560001b620015a0565b6000945050505050620100e1565b6200ff787f2f15d3deb42c3213a46690100d5f689ea3c412ed686c960e24d714f9e3e1f50960001b620015a0565b6200ffa67fc44bb44de2474117bb8f10ef00a8254d19c5a4b8e5bd96ade6054c543b7820d460001b620015a0565b6200ffd47f0ed23c6776c4fc8914ad94a8384c85fdb666f2e39e7f2b7f874e1b7b9e59ffd660001b620015a0565b8086610140019067ffffffffffffffff16908167ffffffffffffffff1681525050620100237f09b906236dba6b1858c675081a6b2ed2c160b4d711fc31930382ea9dbe9004e860001b620015a0565b620100517f5fa72ba24b0f066106b91e31054526a2e6a2fa8f8814932f497573d90de6f16060001b620015a0565b6201005d84846200f91c565b8660e0019067ffffffffffffffff16908167ffffffffffffffff1681525050620100aa7f6cd79e0e068846b37280ec647c6654c5c31c05dfe22638c1886841e3f9433da260001b620015a0565b620100d87f8927dbfa57261b5e38bcc7244c769af0e5a5a1d7827cfcdf3cbb857a5275a28a60001b620015a0565b60019450505050505b9392505050565b620101167f203dc81212c80726daf4194adc73ded580b9b650705f22d5fb48495128b5e40660001b620015a0565b620101447f3c9f51c76a2c5c0d3fb2964cdac85e0a239461872035fc298fa5ac048805be6860001b620015a0565b620101727f8d5061506bb3c941d1368ab7b0d7c96bc3d122c40e37124dbde7cad089b3b62360001b620015a0565b60005b600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490508110156201042857620101ee7fcaba45db971d21f9d4a03d9e43d29098d41f747fd71bc38b21828e6e70ce484460001b620015a0565b6201021c7f9dc08cd7c408ed1de7bb5137fab198aa099646990a36b887ab6878e76708903c60001b620015a0565b8180519060200120600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002082815481106201029c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002001604051620102b5919062013376565b60405180910390201415620103e457620102f27f6290722ae13fd47962163472d09b14446e7895d79982f46f4ef19798b7b1753260001b620015a0565b620103207f195845903c31ef7d6e353dfee81d50838b743bde79239c60634880d279a7a89360001b620015a0565b600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020818154811062010398577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020016000620103b0919062010619565b620103de7f5c30b6d9dcdcbbd07ff74f956c74d210f08139274540453dcdf7489c60cf331760001b620015a0565b62010428565b620104127f28b476c4a37d663b13f6004a1ecd9ea2f97ca3bc40690c820921a70347ee85f460001b620015a0565b80806201041f9062013f51565b91505062010175565b505050565b60006201043a3062010440565b15905090565b600080823b905060008111915050919050565b828054828255906000526020600020908101928215620104a7579160200282015b82811115620104a657825182908051906020019062010495929190620104f9565b509160200191906001019062010474565b5b509050620104b691906201092e565b5090565b6040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b828054620105079062013ee5565b90600052602060002090601f0160209004810192826201052b576000855562010577565b82601f106201054657805160ff191683800117855562010577565b8280016001018555821562010577579182015b828111156201057657825182559160200191906001019062010559565b5b50905062010586919062010956565b5090565b82805482825590600052602060002090810192821562010606579160200282015b82811115620106055782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190620105ab565b5b50905062010615919062010956565b5090565b508054620106279062013ee5565b6000825580601f106201063b57506201065c565b601f0160209004906000526020600020908101906201065b919062010956565b5b50565b50805460008255906000526020600020908101906201067f919062010956565b50565b604051806102c0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff16815260200160008152602001600015158152602001600060018111156201079a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600067ffffffffffffffff168152602001606081526020016060815260200160006002811115620107f9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016000151581526020016201081162010975565b81525090565b604051806101e0016040528060608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160001515815260200160608152602001600015158152602001600015158152602001606081526020016060815260200160001515815260200160006001811115620108fa577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81525090565b6040518060400160405280600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b5b8082111562010952576000818162010948919062010619565b506001016201092f565b5090565b5b808211156201097157600081600090555060010162010957565b5090565b6040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6000620109cb620109c5846201392e565b62013905565b90508083825260208201905082856020860282011115620109eb57600080fd5b60005b8581101562010a1f578162010a04888262010dc2565b845260208401935060208301925050600181019050620109ee565b5050509392505050565b600062010a4062010a3a846201395d565b62013905565b9050808382526020820190508285602086028201111562010a6057600080fd5b60005b8581101562010aaf57813567ffffffffffffffff81111562010a8457600080fd5b80860162010a93898262010f70565b8552602085019450602084019350505060018101905062010a63565b5050509392505050565b600062010ad062010aca846201395d565b62013905565b9050808382526020820190508285602086028201111562010af057600080fd5b60005b8581101562010b3f57815167ffffffffffffffff81111562010b1457600080fd5b80860162010b23898262010f9d565b8552602085019450602084019350505060018101905062010af3565b5050509392505050565b600062010b6062010b5a846201398c565b62013905565b9050808382526020820190508285602086028201111562010b8057600080fd5b60005b8581101562010bb4578162010b9988826201106b565b84526020840193506020830192505060018101905062010b83565b5050509392505050565b600062010bd562010bcf84620139bb565b62013905565b9050808382526020820190508285602086028201111562010bf557600080fd5b60005b8581101562010c4457813567ffffffffffffffff81111562010c1957600080fd5b80860162010c28898262011082565b8552602085019450602084019350505060018101905062010bf8565b5050509392505050565b600062010c6562010c5f84620139ea565b62013905565b905080838252602082019050828560a086028201111562010c8557600080fd5b60005b8581101562010cb9578162010c9e8882620115c2565b845260208401935060a0830192505060018101905062010c88565b5050509392505050565b600062010cda62010cd48462013a19565b62013905565b9050808382526020820190508285606086028201111562010cfa57600080fd5b60005b8581101562010d2e578162010d13888262011c65565b84526020840193506060830192505060018101905062010cfd565b5050509392505050565b600062010d4f62010d498462013a48565b62013905565b90508281526020810184848401111562010d6857600080fd5b62010d7584828562013ea0565b509392505050565b600062010d9462010d8e8462013a48565b62013905565b90508281526020810184848401111562010dad57600080fd5b62010dba84828562013eaf565b509392505050565b60008135905062010dd381620142fc565b92915050565b60008151905062010dea81620142fc565b92915050565b600082601f83011262010e0257600080fd5b813562010e14848260208601620109b4565b91505092915050565b600082601f83011262010e2f57600080fd5b813562010e4184826020860162010a29565b91505092915050565b600082601f83011262010e5c57600080fd5b815162010e6e84826020860162010ab9565b91505092915050565b600082601f83011262010e8957600080fd5b813562010e9b84826020860162010b49565b91505092915050565b600082601f83011262010eb657600080fd5b813562010ec884826020860162010bbe565b91505092915050565b600082601f83011262010ee357600080fd5b815162010ef584826020860162010c4e565b91505092915050565b600082601f83011262010f1057600080fd5b813562010f2284826020860162010cc3565b91505092915050565b60008135905062010f3c8162014316565b92915050565b60008151905062010f538162014316565b92915050565b60008135905062010f6a8162014330565b92915050565b600082601f83011262010f8257600080fd5b813562010f9484826020860162010d38565b91505092915050565b600082601f83011262010faf57600080fd5b815162010fc184826020860162010d7d565b91505092915050565b60008135905062010fdb816201434a565b92915050565b60008135905062010ff28162014364565b92915050565b60008135905062011009816201437e565b92915050565b600081359050620110208162014398565b92915050565b6000813590506201103781620143b2565b92915050565b6000813590506201104e81620143cc565b92915050565b6000815190506201106581620143cc565b92915050565b6000813590506201107c81620143dd565b92915050565b600061030082840312156201109657600080fd5b620110a36102c062013905565b9050600082013567ffffffffffffffff811115620110c057600080fd5b620110ce8482850162010f70565b6000830152506020620110e48482850162010dc2565b602083015250604082013567ffffffffffffffff8111156201110557600080fd5b620111138482850162010f70565b6040830152506060620111298482850162011cfc565b60608301525060806201113f8482850162011cfc565b60808301525060a0620111558482850162011cfc565b60a08301525060c06201116b8482850162011cfc565b60c08301525060e0620111818482850162011cfc565b60e083015250610100620111988482850162011cce565b61010083015250610120620111b08482850162011cfc565b61012083015250610140620111c88482850162011cfc565b6101408301525061016082013567ffffffffffffffff811115620111eb57600080fd5b620111f98482850162010f70565b61016083015250610180620112118482850162011cfc565b610180830152506101a0620112298482850162011cce565b6101a0830152506101c0620112418482850162010f2b565b6101c0830152506101e062011259848285016201106b565b6101e083015250610200620112718482850162011cfc565b6102008301525061022082013567ffffffffffffffff8111156201129457600080fd5b620112a28482850162010df0565b6102208301525061024082013567ffffffffffffffff811115620112c557600080fd5b620112d38482850162010df0565b61024083015250610260620112eb848285016201103d565b61026083015250610280620113038482850162010f2b565b610280830152506102a06201131b84828501620114ed565b6102a08301525092915050565b6000606082840312156201133b57600080fd5b62011347606062013905565b9050600082013567ffffffffffffffff8111156201136457600080fd5b620113728482850162010f70565b6000830152506020620113888482850162010dc2565b60208301525060406201139e8482850162011cfc565b60408301525092915050565b600060e08284031215620113bd57600080fd5b620113c960e062013905565b90506000620113db8482850162011d13565b6000830152506020620113f18482850162011d13565b6020830152506040620114078482850162011d13565b60408301525060606201141d8482850162011d13565b6060830152506080620114338482850162011d13565b60808301525060a0620114498482850162010dd9565b60a08301525060c06201145f8482850162010dd9565b60c08301525092915050565b6000606082840312156201147e57600080fd5b6201148a606062013905565b9050600082013567ffffffffffffffff811115620114a757600080fd5b620114b58482850162010f70565b6000830152506020620114cb8482850162010dc2565b6020830152506040620114e18482850162010dc2565b60408301525092915050565b6000606082840312156201150057600080fd5b6201150c606062013905565b905060006201151e8482850162011cfc565b6000830152506020620115348482850162011cfc565b60208301525060406201154a8482850162011cfc565b60408301525092915050565b6000604082840312156201156957600080fd5b62011575604062013905565b9050600082013567ffffffffffffffff8111156201159257600080fd5b620115a08482850162010f70565b6000830152506020620115b68482850162011cfc565b60208301525092915050565b600060a08284031215620115d557600080fd5b620115e160a062013905565b90506000620115f38482850162010dd9565b6000830152506020620116098482850162010dd9565b60208301525060406201161f8482850162011d13565b6040830152506060620116358482850162011ce5565b60608301525060806201164b8482850162010f42565b60808301525092915050565b600061018082840312156201166b57600080fd5b6201167861018062013905565b905060006201168a8482850162010dd9565b6000830152506020620116a08482850162011d13565b6020830152506040620116b68482850162011d13565b6040830152506060620116cc8482850162011d13565b6060830152506080620116e28482850162011054565b60808301525060a0620116f88482850162011ce5565b60a08301525060c06201170e8482850162011ce5565b60c08301525060e0620117248482850162011d13565b60e0830152506101006201173b8482850162011d13565b61010083015250610120620117538482850162011d13565b610120830152506101406201176b8482850162010f42565b6101408301525061016082015167ffffffffffffffff8111156201178e57600080fd5b6201179c8482850162010e4a565b6101608301525092915050565b60006101608284031215620117bd57600080fd5b620117ca61016062013905565b90506000620117dc8482850162011cfc565b6000830152506020620117f28482850162011cfc565b6020830152506040620118088482850162011cfc565b60408301525060606201181e8482850162011cfc565b6060830152506080620118348482850162011cfc565b60808301525060a06201184a8482850162011cfc565b60a08301525060c0620118608482850162011cfc565b60c08301525060e0620118768482850162011cfc565b60e0830152506101006201188d8482850162011cfc565b61010083015250610120620118a58482850162011cfc565b61012083015250610140620118bd8482850162011cfc565b6101408301525092915050565b60006101608284031215620118de57600080fd5b620118eb61016062013905565b90506000620118fd8482850162011d13565b6000830152506020620119138482850162011d13565b6020830152506040620119298482850162011d13565b60408301525060606201193f8482850162011d13565b6060830152506080620119558482850162011d13565b60808301525060a06201196b8482850162011d13565b60a08301525060c0620119818482850162011d13565b60c08301525060e0620119978482850162011d13565b60e083015250610100620119ae8482850162011d13565b61010083015250610120620119c68482850162011d13565b61012083015250610140620119de8482850162011d13565b6101408301525092915050565b60006101e08284031215620119ff57600080fd5b62011a0c6101e062013905565b9050600082013567ffffffffffffffff81111562011a2957600080fd5b62011a378482850162010f70565b600083015250602062011a4d8482850162011cfc565b602083015250604062011a638482850162011cfc565b604083015250606062011a798482850162011cfc565b606083015250608062011a8f8482850162011cce565b60808301525060a062011aa58482850162011cfc565b60a08301525060c062011abb8482850162011cfc565b60c08301525060e062011ad18482850162010f2b565b60e08301525061010082013567ffffffffffffffff81111562011af357600080fd5b62011b018482850162010f70565b6101008301525061012062011b198482850162010f2b565b6101208301525061014062011b318482850162010f2b565b6101408301525061016082013567ffffffffffffffff81111562011b5457600080fd5b62011b628482850162010f70565b6101608301525061018082013567ffffffffffffffff81111562011b8557600080fd5b62011b938482850162010efe565b610180830152506101a062011bab8482850162010f2b565b6101a0830152506101c062011bc3848285016201106b565b6101c08301525092915050565b600060a0828403121562011be357600080fd5b62011bef60a062013905565b9050600062011c018482850162011d13565b600083015250602062011c178482850162011d13565b602083015250604062011c2d8482850162011d13565b604083015250606062011c438482850162011ce5565b606083015250608062011c598482850162011ce5565b60808301525092915050565b60006060828403121562011c7857600080fd5b62011c84606062013905565b9050600062011c968482850162010dc2565b600083015250602062011cac8482850162011cfc565b602083015250604062011cc28482850162011cfc565b60408301525092915050565b60008135905062011cdf81620143ee565b92915050565b60008151905062011cf681620143ee565b92915050565b60008135905062011d0d8162014408565b92915050565b60008151905062011d248162014408565b92915050565b60006020828403121562011d3d57600080fd5b600062011d4d8482850162010dc2565b91505092915050565b6000806040838503121562011d6a57600080fd5b600062011d7a8582860162010dc2565b925050602083013567ffffffffffffffff81111562011d9857600080fd5b62011da68582860162010e1d565b9150509250929050565b6000806040838503121562011dc457600080fd5b600062011dd48582860162010dc2565b925050602083013567ffffffffffffffff81111562011df257600080fd5b62011e008582860162010f70565b9150509250929050565b60006020828403121562011e1d57600080fd5b600082013567ffffffffffffffff81111562011e3857600080fd5b62011e468482850162010e1d565b91505092915050565b60008060006060848603121562011e6557600080fd5b600084013567ffffffffffffffff81111562011e8057600080fd5b62011e8e8682870162010e1d565b935050602062011ea18682870162010dc2565b925050604084013567ffffffffffffffff81111562011ebf57600080fd5b62011ecd8682870162010e77565b9150509250925092565b60008060006101a0848603121562011eee57600080fd5b600084013567ffffffffffffffff81111562011f0957600080fd5b62011f178682870162010e1d565b935050602062011f2a86828701620117a9565b92505061018062011f3e8682870162011cce565b9150509250925092565b60006020828403121562011f5b57600080fd5b600082013567ffffffffffffffff81111562011f7657600080fd5b62011f848482850162010ea4565b91505092915050565b60006020828403121562011fa057600080fd5b600082015167ffffffffffffffff81111562011fbb57600080fd5b62011fc98482850162010ed1565b91505092915050565b60006020828403121562011fe557600080fd5b600062011ff58482850162010f59565b91505092915050565b6000602082840312156201201157600080fd5b600082013567ffffffffffffffff8111156201202c57600080fd5b6201203a8482850162010f70565b91505092915050565b600080600080600060a086880312156201205c57600080fd5b60006201206c8882890162010fca565b95505060206201207f8882890162010fe1565b9450506040620120928882890162011026565b9350506060620120a5888289016201100f565b9250506080620120b88882890162010ff8565b9150509295509295909350565b600060208284031215620120d857600080fd5b600082013567ffffffffffffffff811115620120f357600080fd5b620121018482850162011082565b91505092915050565b6000806000606084860312156201212057600080fd5b600084013567ffffffffffffffff8111156201213b57600080fd5b620121498682870162011082565b93505060206201215c8682870162010f2b565b92505060406201216f8682870162010f2b565b9150509250925092565b6000602082840312156201218c57600080fd5b600082013567ffffffffffffffff811115620121a757600080fd5b620121b58482850162011328565b91505092915050565b600060e08284031215620121d157600080fd5b6000620121e184828501620113aa565b91505092915050565b600060208284031215620121fd57600080fd5b600082013567ffffffffffffffff8111156201221857600080fd5b62012226848285016201146b565b91505092915050565b6000602082840312156201224257600080fd5b600082013567ffffffffffffffff8111156201225d57600080fd5b6201226b8482850162011556565b91505092915050565b6000602082840312156201228757600080fd5b600082015167ffffffffffffffff811115620122a257600080fd5b620122b08482850162011657565b91505092915050565b60006101608284031215620122cd57600080fd5b6000620122dd84828501620118ca565b91505092915050565b60008060006101a08486031215620122fd57600080fd5b60006201230d86828701620117a9565b935050610160620123218682870162011cce565b92505061018084013567ffffffffffffffff8111156201234057600080fd5b6201234e8682870162011082565b9150509250925092565b60008061018083850312156201236d57600080fd5b60006201237d85828601620117a9565b925050610160620123918582860162011cfc565b9150509250929050565b60008060006101a08486031215620123b257600080fd5b6000620123c286828701620117a9565b935050610160620123d68682870162011cfc565b925050610180620123ea8682870162011cfc565b9150509250925092565b6000806000806101c085870312156201240c57600080fd5b60006201241c87828801620117a9565b945050610160620124308782880162011cfc565b935050610180620124448782880162011cfc565b9250506101a0620124588782880162011cfc565b91505092959194509250565b60008060008060006101e086880312156201247e57600080fd5b60006201248e88828901620117a9565b955050610160620124a28882890162011cfc565b945050610180620124b68882890162011cfc565b9350506101a0620124ca8882890162011cfc565b9250506101c0620124de8882890162011cfc565b9150509295509295909350565b600060208284031215620124fe57600080fd5b600082013567ffffffffffffffff8111156201251957600080fd5b6201252784828501620119eb565b91505092915050565b60008060006101a084860312156201254757600080fd5b600084013567ffffffffffffffff8111156201256257600080fd5b6201257086828701620119eb565b93505060206201258386828701620117a9565b925050610180620125978682870162011cce565b9150509250925092565b60008060408385031215620125b557600080fd5b600083013567ffffffffffffffff811115620125d057600080fd5b620125de85828601620119eb565b9250506020620125f18582860162011cce565b9150509250929050565b600060a082840312156201260e57600080fd5b60006201261e8482850162011bd0565b91505092915050565b60006201263583836201266d565b60208301905092915050565b60006201264f83836201289e565b905092915050565b600062012665838362012bac565b905092915050565b620126788162013d31565b82525050565b620126898162013d31565b82525050565b60006201269c8262013ac3565b620126a8818562013b16565b9350620126b58362013a7e565b8060005b83811015620126ec578151620126d0888262012627565b9750620126dd8362013aef565b925050600181019050620126b9565b5085935050505092915050565b6000620127068262013ace565b62012712818562013b27565b935083602082028501620127268562013a8e565b8060005b8581101562012768578484038952815162012746858262012641565b9450620127538362013afc565b925060208a019950506001810190506201272a565b50829750879550505050505092915050565b6000620127878262013ace565b62012793818562013b38565b935083602082028501620127a78562013a8e565b8060005b85811015620127e95784840389528151620127c7858262012641565b9450620127d48362013afc565b925060208a01995050600181019050620127ab565b50829750879550505050505092915050565b6000620128088262013ad9565b62012814818562013b49565b935083602082028501620128288562013a9e565b8060005b858110156201286a578484038952815162012848858262012657565b9450620128558362013b09565b925060208a019950506001810190506201282c565b50829750879550505050505092915050565b620128878162013d45565b82525050565b620128988162013d45565b82525050565b6000620128ab8262013ae4565b620128b7818562013b5a565b9350620128c981856020860162013eaf565b620128d481620140c0565b840191505092915050565b6000620128ec8262013ae4565b620128f8818562013b6b565b93506201290a81856020860162013eaf565b6201291581620140c0565b840191505092915050565b60006201292d8262013ae4565b62012939818562013b7c565b93506201294b81856020860162013eaf565b80840191505092915050565b60008154620129668162013ee5565b62012972818662013b7c565b94506001821660008114620129905760018114620129a257620129d9565b60ff19831686528186019350620129d9565b620129ad8562013aae565b60005b83811015620129d157815481890152600182019150602081019050620129b0565b838801955050505b50505092915050565b620129ed8162013e3c565b82525050565b620129fe8162013e50565b82525050565b62012a0f8162013e64565b82525050565b62012a208162013e78565b82525050565b600062012a3560118362013b87565b915062012a4282620140d1565b602082019050919050565b600062012a5c600c8362013b87565b915062012a6982620140fa565b602082019050919050565b600062012a83602b8362013b87565b915062012a908262014123565b604082019050919050565b600062012aaa600e8362013b87565b915062012ab78262014172565b602082019050919050565b600062012ad1602e8362013b87565b915062012ade826201419b565b604082019050919050565b600062012af8600f8362013b87565b915062012b0582620141ea565b602082019050919050565b600062012b1f601f8362013b87565b915062012b2c8262014213565b602082019050919050565b600062012b46601a8362013b87565b915062012b53826201423c565b602082019050919050565b600062012b6d60128362013b87565b915062012b7a8262014265565b602082019050919050565b600062012b9460148362013b87565b915062012ba1826201428e565b602082019050919050565b600061030083016000830151848203600086015262012bcc82826201289e565b915050602083015162012be360208601826201266d565b506040830151848203604086015262012bfd82826201289e565b915050606083015162012c1460608601826201333b565b50608083015162012c2960808601826201333b565b5060a083015162012c3e60a08601826201333b565b5060c083015162012c5360c08601826201333b565b5060e083015162012c6860e08601826201333b565b5061010083015162012c7f61010086018262013308565b5061012083015162012c966101208601826201333b565b5061014083015162012cad6101408601826201333b565b5061016083015184820361016086015262012cc982826201289e565b91505061018083015162012ce26101808601826201333b565b506101a083015162012cf96101a086018262013308565b506101c083015162012d106101c08601826201287c565b506101e083015162012d276101e086018262012a04565b5061020083015162012d3e6102008601826201333b565b5061022083015184820361022086015262012d5a82826201268f565b91505061024083015184820361024086015262012d7882826201268f565b91505061026083015162012d91610260860182620129f3565b5061028083015162012da86102808601826201287c565b506102a083015162012dbf6102a086018262013084565b508091505092915050565b600061030083016000830151848203600086015262012dea82826201289e565b915050602083015162012e0160208601826201266d565b506040830151848203604086015262012e1b82826201289e565b915050606083015162012e3260608601826201333b565b50608083015162012e4760808601826201333b565b5060a083015162012e5c60a08601826201333b565b5060c083015162012e7160c08601826201333b565b5060e083015162012e8660e08601826201333b565b5061010083015162012e9d61010086018262013308565b5061012083015162012eb46101208601826201333b565b5061014083015162012ecb6101408601826201333b565b5061016083015184820361016086015262012ee782826201289e565b91505061018083015162012f006101808601826201333b565b506101a083015162012f176101a086018262013308565b506101c083015162012f2e6101c08601826201287c565b506101e083015162012f456101e086018262012a04565b5061020083015162012f5c6102008601826201333b565b5061022083015184820361022086015262012f7882826201268f565b91505061024083015184820361024086015262012f9682826201268f565b91505061026083015162012faf610260860182620129f3565b5061028083015162012fc66102808601826201287c565b506102a083015162012fdd6102a086018262013084565b508091505092915050565b60e0820160008201516201300060008501826201333b565b5060208201516201301560208501826201333b565b5060408201516201302a60408501826201333b565b5060608201516201303f60608501826201333b565b5060808201516201305460808501826201333b565b5060a08201516201306960a08501826201266d565b5060c08201516201307e60c08501826201266d565b50505050565b6060820160008201516201309c60008501826201333b565b506020820151620130b160208501826201333b565b506040820151620130c660408501826201333b565b50505050565b604082016000820151620130e460008501826201333b565b506020820151620130f960208501826201333b565b50505050565b6000610180830160008301516201311a60008601826201266d565b5060208301516201312f60208601826201333b565b5060408301516201314460408601826201333b565b5060608301516201315960608601826201333b565b5060808301516201316e6080860182620129f3565b5060a08301516201318360a086018262013308565b5060c08301516201319860c086018262013308565b5060e0830151620131ad60e08601826201333b565b50610100830151620131c46101008601826201333b565b50610120830151620131db6101208601826201333b565b50610140830151620131f26101408601826201287c565b506101608301518482036101608601526201320e8282620126f9565b9150508091505092915050565b6040820160008201516201323360008501826201266d565b5060208201516201324860208501826201333b565b50505050565b6060820160008201516201326660008501826201333b565b5060208201516201327b60208501826201333b565b5060408201516201329060408501826201333b565b50505050565b60a082016000820151620132ae60008501826201333b565b506020820151620132c360208501826201333b565b506040820151620132d860408501826201333b565b506060820151620132ed606085018262013308565b50608082015162013302608085018262013308565b50505050565b620133138162013e1e565b82525050565b620133248162013e1e565b82525050565b620133358162013e8c565b82525050565b620133468162013e28565b82525050565b620133578162013e28565b82525050565b60006201336b828462012920565b915081905092915050565b600062013384828462012957565b915081905092915050565b6000602082019050620133a660008301846201267e565b92915050565b600060c082019050620133c360008301856201267e565b620133d2602083018462013296565b9392505050565b60006020820190508181036000830152620133f581846201277a565b905092915050565b600060608201905081810360008301526201341981866201277a565b90506201342a60208301856201334c565b6201343960408301846201288d565b949350505050565b600060208201905081810360008301526201345d8184620127fb565b905092915050565b60006040820190508181036000830152620134818185620127fb565b90506201349260208301846201288d565b9392505050565b6000602082019050620134b060008301846201288d565b92915050565b60006020820190508181036000830152620134d28184620128df565b905092915050565b60006060820190508181036000830152620134f68185620128df565b9050620135076020830184620130cc565b9392505050565b6000608082019050620135256000830187620129e2565b62013534602083018662013319565b81810360408301526201354881856201277a565b90506201355960608301846201267e565b95945050505050565b6000608082019050620135796000830187620129e2565b62013588602083018662013319565b81810360408301526201359c8185620128df565b9050620135ad60608301846201267e565b95945050505050565b600060e082019050620135cd600083018a620129e2565b620135dc602083018962013319565b8181036040830152620135f08188620128df565b90506201360160608301876201334c565b6201361060808301866201267e565b6201361f60a08301856201334c565b6201362e60c08301846201288d565b98975050505050505050565b600060208201905062013651600083018462012a15565b92915050565b60006020820190508181036000830152620136728162012a26565b9050919050565b60006020820190508181036000830152620136948162012a4d565b9050919050565b60006020820190508181036000830152620136b68162012a74565b9050919050565b60006020820190508181036000830152620136d88162012a9b565b9050919050565b60006020820190508181036000830152620136fa8162012ac2565b9050919050565b600060208201905081810360008301526201371c8162012ae9565b9050919050565b600060208201905081810360008301526201373e8162012b10565b9050919050565b60006020820190508181036000830152620137608162012b37565b9050919050565b60006020820190508181036000830152620137828162012b5e565b9050919050565b60006020820190508181036000830152620137a48162012b85565b9050919050565b60006020820190508181036000830152620137c7818462012dca565b905092915050565b600060e082019050620137e6600083018462012fe8565b92915050565b60006040820190508181036000830152620138088185620130ff565b905081810360208301526201381e818462012dca565b90509392505050565b60006040820190506201383e60008301846201321b565b92915050565b60006060820190506201385b60008301846201324e565b92915050565b600060408201905062013878600083018562013319565b62013887602083018462013319565b9392505050565b6000604082019050620138a5600083018562013319565b620138b460208301846201332a565b9392505050565b6000602082019050620138d260008301846201334c565b92915050565b6000604082019050620138ef60008301856201332a565b620138fe60208301846201332a565b9392505050565b60006201391162013924565b90506201391f828262013f1b565b919050565b6000604051905090565b600067ffffffffffffffff8211156201394c576201394b62014091565b5b602082029050602081019050919050565b600067ffffffffffffffff8211156201397b576201397a62014091565b5b602082029050602081019050919050565b600067ffffffffffffffff821115620139aa57620139a962014091565b5b602082029050602081019050919050565b600067ffffffffffffffff821115620139d957620139d862014091565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562013a085762013a0762014091565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562013a375762013a3662014091565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562013a665762013a6562014091565b5b62013a7182620140c0565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b60008190508160005260206000209050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600062013ba58262013e1e565b915062013bb28362013e1e565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111562013bea5762013be962013fd5565b5b828201905092915050565b600062013c028262013e28565b915062013c0f8362013e28565b92508267ffffffffffffffff0382111562013c2f5762013c2e62013fd5565b5b828201905092915050565b600062013c478262013e28565b915062013c548362013e28565b92508262013c675762013c6662014004565b5b828204905092915050565b600062013c7f8262013e28565b915062013c8c8362013e28565b92508167ffffffffffffffff048311821515161562013cb05762013caf62013fd5565b5b828202905092915050565b600062013cc88262013e1e565b915062013cd58362013e1e565b92508282101562013ceb5762013cea62013fd5565b5b828203905092915050565b600062013d038262013e28565b915062013d108362013e28565b92508282101562013d265762013d2562013fd5565b5b828203905092915050565b600062013d3e8262013dfe565b9050919050565b60008115159050919050565b6000819050919050565b600062013d688262013d31565b9050919050565b600062013d7c8262013d31565b9050919050565b600062013d908262013d31565b9050919050565b600062013da48262013d31565b9050919050565b600062013db88262013d31565b9050919050565b600081905062013dcf82620142b7565b919050565b600081905062013de482620142ce565b919050565b600081905062013df982620142e5565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600062013e498262013dbf565b9050919050565b600062013e5d8262013dd4565b9050919050565b600062013e718262013de9565b9050919050565b600062013e858262013e28565b9050919050565b600062013e998262013e28565b9050919050565b82818337600083830152505050565b60005b8381101562013ecf57808201518184015260208101905062013eb2565b8381111562013edf576000848401525b50505050565b6000600282049050600182168062013efe57607f821691505b6020821081141562013f155762013f1462014062565b5b50919050565b62013f2682620140c0565b810181811067ffffffffffffffff8211171562013f485762013f4762014091565b5b80604052505050565b600062013f5e8262013e1e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141562013f945762013f9362013fd5565b5b600182019050919050565b600062013fac8262013e28565b915067ffffffffffffffff82141562013fca5762013fc962013fd5565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f66696c654c69737420697320656d707479000000000000000000000000000000600082015250565b7f66696c6520657870697265640000000000000000000000000000000000000000600082015250565b7f43757272656e74206f776e6572206d75737420626520746865206f776e65722060008201527f6f66207468652066696c65000000000000000000000000000000000000000000602082015250565b7f66696c65206e6f74206578697374000000000000000000000000000000000000600082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f66696c652074797065206572726f720000000000000000000000000000000000600082015250565b7f66696c6553697a65206d7573742062652067726561746572207468616e203000600082015250565b7f66696c6548617368206d757374206265206e6f7420656d707479000000000000600082015250565b7f66696c6520616c72656164792065786973740000000000000000000000000000600082015250565b7f696e73756666696369656e74206465706f736974000000000000000000000000600082015250565b600a8110620142cb57620142ca62014033565b5b50565b60038110620142e257620142e162014033565b5b50565b60028110620142f957620142f862014033565b5b50565b620143078162013d31565b81146201431357600080fd5b50565b620143218162013d45565b81146201432d57600080fd5b50565b6201433b8162013d51565b81146201434757600080fd5b50565b620143558162013d5b565b81146201436157600080fd5b50565b6201436f8162013d6f565b81146201437b57600080fd5b50565b620143898162013d83565b81146201439557600080fd5b50565b620143a38162013d97565b8114620143af57600080fd5b50565b620143bd8162013dab565b8114620143c957600080fd5b50565b60038110620143da57600080fd5b50565b60028110620143eb57600080fd5b50565b620143f98162013e1e565b81146201440557600080fd5b50565b620144138162013e28565b81146201441f57600080fd5b5056fea2646970667358221220997623a5ca5486ac0c17dd5251f9fd72ae01d782759839ea67877dd3206ce35164736f6c63430008040033",
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
// Solidity: function GetFileInfo(bytes fileHash) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)))
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
// Solidity: function GetFileInfo(bytes fileHash) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)))
func (_Store *StoreSession) GetFileInfo(fileHash []byte) (FileInfo, error) {
	return _Store.Contract.GetFileInfo(&_Store.CallOpts, fileHash)
}

// GetFileInfo is a free data retrieval call binding the contract method 0xdefce5d4.
//
// Solidity: function GetFileInfo(bytes fileHash) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)))
func (_Store *StoreCallerSession) GetFileInfo(fileHash []byte) (FileInfo, error) {
	return _Store.Contract.GetFileInfo(&_Store.CallOpts, fileHash)
}

// GetFileInfos is a free data retrieval call binding the contract method 0xc6e8392a.
//
// Solidity: function GetFileInfos(bytes[] _fileList) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64))[])
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
// Solidity: function GetFileInfos(bytes[] _fileList) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64))[])
func (_Store *StoreSession) GetFileInfos(_fileList [][]byte) ([]FileInfo, error) {
	return _Store.Contract.GetFileInfos(&_Store.CallOpts, _fileList)
}

// GetFileInfos is a free data retrieval call binding the contract method 0xc6e8392a.
//
// Solidity: function GetFileInfos(bytes[] _fileList) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64))[])
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

// UpdateFileInfoForRenew is a free data retrieval call binding the contract method 0xeb7336a0.
//
// Solidity: function UpdateFileInfoForRenew((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo) pure returns(bool)
func (_Store *StoreCaller) UpdateFileInfoForRenew(opts *bind.CallOpts, setting Setting, newExpireHeight *big.Int, fileInfo FileInfo) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "UpdateFileInfoForRenew", setting, newExpireHeight, fileInfo)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UpdateFileInfoForRenew is a free data retrieval call binding the contract method 0xeb7336a0.
//
// Solidity: function UpdateFileInfoForRenew((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo) pure returns(bool)
func (_Store *StoreSession) UpdateFileInfoForRenew(setting Setting, newExpireHeight *big.Int, fileInfo FileInfo) (bool, error) {
	return _Store.Contract.UpdateFileInfoForRenew(&_Store.CallOpts, setting, newExpireHeight, fileInfo)
}

// UpdateFileInfoForRenew is a free data retrieval call binding the contract method 0xeb7336a0.
//
// Solidity: function UpdateFileInfoForRenew((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo) pure returns(bool)
func (_Store *StoreCallerSession) UpdateFileInfoForRenew(setting Setting, newExpireHeight *big.Int, fileInfo FileInfo) (bool, error) {
	return _Store.Contract.UpdateFileInfoForRenew(&_Store.CallOpts, setting, newExpireHeight, fileInfo)
}

// UpdateFilesForRenew is a free data retrieval call binding the contract method 0x3f2cc9a0.
//
// Solidity: function UpdateFilesForRenew(bytes[] _fileList, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64))[], bool)
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
// Solidity: function UpdateFilesForRenew(bytes[] _fileList, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64))[], bool)
func (_Store *StoreSession) UpdateFilesForRenew(_fileList [][]byte, setting Setting, newExpireHeight *big.Int) ([]FileInfo, bool, error) {
	return _Store.Contract.UpdateFilesForRenew(&_Store.CallOpts, _fileList, setting, newExpireHeight)
}

// UpdateFilesForRenew is a free data retrieval call binding the contract method 0x3f2cc9a0.
//
// Solidity: function UpdateFilesForRenew(bytes[] _fileList, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight) view returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64))[], bool)
func (_Store *StoreCallerSession) UpdateFilesForRenew(_fileList [][]byte, setting Setting, newExpireHeight *big.Int) ([]FileInfo, bool, error) {
	return _Store.Contract.UpdateFilesForRenew(&_Store.CallOpts, _fileList, setting, newExpireHeight)
}

// C0xf932268c is a free data retrieval call binding the contract method 0x0df7e99e.
//
// Solidity: function c_0xf932268c(bytes32 c__0xf932268c) pure returns()
func (_Store *StoreCaller) C0xf932268c(opts *bind.CallOpts, c__0xf932268c [32]byte) error {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "c_0xf932268c", c__0xf932268c)

	if err != nil {
		return err
	}

	return err

}

// C0xf932268c is a free data retrieval call binding the contract method 0x0df7e99e.
//
// Solidity: function c_0xf932268c(bytes32 c__0xf932268c) pure returns()
func (_Store *StoreSession) C0xf932268c(c__0xf932268c [32]byte) error {
	return _Store.Contract.C0xf932268c(&_Store.CallOpts, c__0xf932268c)
}

// C0xf932268c is a free data retrieval call binding the contract method 0x0df7e99e.
//
// Solidity: function c_0xf932268c(bytes32 c__0xf932268c) pure returns()
func (_Store *StoreCallerSession) C0xf932268c(c__0xf932268c [32]byte) error {
	return _Store.Contract.C0xf932268c(&_Store.CallOpts, c__0xf932268c)
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

// CleanupForDeleteFile is a paid mutator transaction binding the contract method 0x34afd8c1.
//
// Solidity: function CleanupForDeleteFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo, bool rmInfo, bool rmList) returns()
func (_Store *StoreTransactor) CleanupForDeleteFile(opts *bind.TransactOpts, fileInfo FileInfo, rmInfo bool, rmList bool) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "CleanupForDeleteFile", fileInfo, rmInfo, rmList)
}

// CleanupForDeleteFile is a paid mutator transaction binding the contract method 0x34afd8c1.
//
// Solidity: function CleanupForDeleteFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo, bool rmInfo, bool rmList) returns()
func (_Store *StoreSession) CleanupForDeleteFile(fileInfo FileInfo, rmInfo bool, rmList bool) (*types.Transaction, error) {
	return _Store.Contract.CleanupForDeleteFile(&_Store.TransactOpts, fileInfo, rmInfo, rmList)
}

// CleanupForDeleteFile is a paid mutator transaction binding the contract method 0x34afd8c1.
//
// Solidity: function CleanupForDeleteFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo, bool rmInfo, bool rmList) returns()
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

// StoreFile is a paid mutator transaction binding the contract method 0xaf6a8370.
//
// Solidity: function StoreFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreTransactor) StoreFile(opts *bind.TransactOpts, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "StoreFile", fileInfo)
}

// StoreFile is a paid mutator transaction binding the contract method 0xaf6a8370.
//
// Solidity: function StoreFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreSession) StoreFile(fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.StoreFile(&_Store.TransactOpts, fileInfo)
}

// StoreFile is a paid mutator transaction binding the contract method 0xaf6a8370.
//
// Solidity: function StoreFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreTransactorSession) StoreFile(fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.StoreFile(&_Store.TransactOpts, fileInfo)
}

// UpdateFileInfo is a paid mutator transaction binding the contract method 0x1aa05c5c.
//
// Solidity: function UpdateFileInfo((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) f) returns()
func (_Store *StoreTransactor) UpdateFileInfo(opts *bind.TransactOpts, f FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateFileInfo", f)
}

// UpdateFileInfo is a paid mutator transaction binding the contract method 0x1aa05c5c.
//
// Solidity: function UpdateFileInfo((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) f) returns()
func (_Store *StoreSession) UpdateFileInfo(f FileInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileInfo(&_Store.TransactOpts, f)
}

// UpdateFileInfo is a paid mutator transaction binding the contract method 0x1aa05c5c.
//
// Solidity: function UpdateFileInfo((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) f) returns()
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

// DeleteFilesInner is a paid mutator transaction binding the contract method 0x1fe07caa.
//
// Solidity: function deleteFilesInner((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64))[] files) returns()
func (_Store *StoreTransactor) DeleteFilesInner(opts *bind.TransactOpts, files []FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "deleteFilesInner", files)
}

// DeleteFilesInner is a paid mutator transaction binding the contract method 0x1fe07caa.
//
// Solidity: function deleteFilesInner((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64))[] files) returns()
func (_Store *StoreSession) DeleteFilesInner(files []FileInfo) (*types.Transaction, error) {
	return _Store.Contract.DeleteFilesInner(&_Store.TransactOpts, files)
}

// DeleteFilesInner is a paid mutator transaction binding the contract method 0x1fe07caa.
//
// Solidity: function deleteFilesInner((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64))[] files) returns()
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
