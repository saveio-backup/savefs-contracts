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

// FSConfig is an auto generated low-level Go binding around an user-defined struct.
type FSConfig struct {
	DEFAULTBLOCKINTERVAL uint64
	DEFAULTPROVEPERIOD   uint64
	INSECTORSIZE         uint64
}

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

// OwnerChange is an auto generated low-level Go binding around an user-defined struct.
type OwnerChange struct {
	FileHash []byte
	CurOwner common.Address
	NewOwner common.Address
}

// PlotInfo is an auto generated low-level Go binding around an user-defined struct.
type PlotInfo struct {
	NumberID   uint64
	StartNonce uint64
	Nonces     uint64
}

// PriChange is an auto generated low-level Go binding around an user-defined struct.
type PriChange struct {
	FileHash  []byte
	Privilege uint64
}

// SectorRef is an auto generated low-level Go binding around an user-defined struct.
type SectorRef struct {
	NodeAddr common.Address
	SectorId uint64
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

// TransferState is an auto generated low-level Go binding around an user-defined struct.
type TransferState struct {
	From  common.Address
	To    common.Address
	Value uint64
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"port\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"}],\"name\":\"DNSNodeRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DNSNodeUnReg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"DnsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"FsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structNameInfo\",\"name\":\"newer\",\"type\":\"tuple\"}],\"name\":\"NotifyNameInfoAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"PDPVerifyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nodeAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"ref\",\"type\":\"tuple\"}],\"name\":\"AddFileSectorRef\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"AddFileToUnSettleList\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"currentHeight\",\"type\":\"uint256\"}],\"name\":\"CalcDepositFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"proveTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"CalcFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"CurOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NewOwner\",\"type\":\"address\"}],\"internalType\":\"structOwnerChange\",\"name\":\"ownerChange\",\"type\":\"tuple\"}],\"name\":\"ChangeFileOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"privilege\",\"type\":\"uint64\"}],\"internalType\":\"structPriChange\",\"name\":\"priChange\",\"type\":\"tuple\"}],\"name\":\"ChangeFilePrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"rmInfo\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rmList\",\"type\":\"bool\"}],\"name\":\"CleanupForDeleteFile\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"enumStorageType[]\",\"name\":\"storageType\",\"type\":\"uint8[]\"}],\"name\":\"DeleteExpiredFilesFromList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteFile\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteFileInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"}],\"name\":\"DeleteFiles\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteProveDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteUnsettledFiles\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FromAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ReNewTimes\",\"type\":\"uint64\"}],\"internalType\":\"structFileReNewInfo\",\"name\":\"fileReNewInfo\",\"type\":\"tuple\"}],\"name\":\"FileReNew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"GetFileInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"}],\"name\":\"GetFileInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetFileList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnProveCandidateFiles\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnProvePrimaryFiles\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnSettledFileList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"}],\"name\":\"GetUploadStorageFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"StoreFile\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"f\",\"type\":\"tuple\"}],\"name\":\"UpdateFileInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"list\",\"type\":\"bytes[]\"}],\"name\":\"UpdateFileList\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"newExpireHeight\",\"type\":\"uint256\"}],\"name\":\"UpdateFilesForRenew\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0x74cde3b0\",\"type\":\"bytes32\"}],\"name\":\"c_0x74cde3b0\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"currentHeight\",\"type\":\"uint256\"}],\"name\":\"calcUploadFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"files\",\"type\":\"tuple[]\"}],\"name\":\"deleteFilesInner\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"error\",\"type\":\"string\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractINode\",\"name\":\"_node\",\"type\":\"address\"},{\"internalType\":\"contractISpace\",\"name\":\"_space\",\"type\":\"address\"},{\"internalType\":\"contractISector\",\"name\":\"_sector\",\"type\":\"address\"},{\"internalType\":\"contractIProve\",\"name\":\"_prove\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"DEFAULT_BLOCK_INTERVAL\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DEFAULT_PROVE_PERIOD\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"IN_SECTOR_SIZE\",\"type\":\"uint64\"}],\"internalType\":\"structFSConfig\",\"name\":\"fsConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractFileExtra\",\"name\":\"_fileExtra\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5061c76d80620000226000396000f3fe6080604052600436106101c25760003560e01c80638d41f9f8116100f7578063c6e8392a11610095578063d70e627211610064578063d70e627214610625578063dc1ec30b14610641578063defce5d41461067e578063f54cd295146106bb576101c2565b8063c6e8392a14610552578063cc76e80d1461058f578063ce904554146105cc578063d49ce874146105e8576101c2565b80639cd103e9116100d15780639cd103e9146104b4578063b6af10fb146104d0578063bc1c783e146104f9578063c43007e514610522576101c2565b80638d41f9f81461041157806395b0b54b1461044e5780639a2b5e6314610477576101c2565b80633f2cc9a01161016457806352e061461161013e57806352e061461461039457806357d94399146103bd5780635a0e7482146103d9578063681850d7146103f5576101c2565b80633f2cc9a01461030a57806341bc86cb1461033b578063432152ce14610378576101c2565b8063207e33be116101a0578063207e33be1461024957806328a8565c1461027257806334fddaac146102af5780633ad525a9146102cb576101c2565b80630be09702146101c7578063178e4dc0146101e35780631ca4a47914610220575b600080fd5b6101e160048036038101906101dc9190619d09565b6106d7565b005b3480156101ef57600080fd5b5061020a6004803603810190610205919061a579565b610fd7565b604051610217919061bc20565b60405180910390f35b34801561022c57600080fd5b506102476004803603810190610242919061a09c565b61111b565b005b34801561025557600080fd5b50610270600480360381019061026b919061a3a6565b61111e565b005b34801561027e57600080fd5b5061029960048036038101906102949190619d09565b611232565b6040516102a6919061b60f565b60405180910390f35b6102c960048036038101906102c49190619d86565b61136f565b005b3480156102d757600080fd5b506102f260048036038101906102ed9190619e5c565b611486565b6040516103019392919061b671565b60405180910390f35b610324600480360381019061031f9190619edb565b612140565b60405161033292919061b6d1565b60405180910390f35b34801561034757600080fd5b50610362600480360381019061035d919061a538565b61228a565b60405161036f919061bc20565b60405180910390f35b610392600480360381019061038d9190619dda565b612526565b005b3480156103a057600080fd5b506103bb60048036038101906103b6919061a0c5565b612b47565b005b6103d760048036038101906103d2919061a324565b612d40565b005b6103f360048036038101906103ee919061a106565b613127565b005b61040f600480360381019061040a919061a23b565b61323e565b005b34801561041d57600080fd5b5061043860048036038101906104339190619d09565b613352565b604051610445919061b60f565b60405180910390f35b34801561045a57600080fd5b506104756004803603810190610470919061a0c5565b6134f6565b005b34801561048357600080fd5b5061049e60048036038101906104999190619d09565b61360a565b6040516104ab919061b60f565b60405180910390f35b6104ce60048036038101906104c9919061a2bd565b613747565b005b3480156104dc57600080fd5b506104f760048036038101906104f2919061a3e7565b613b8e565b005b34801561050557600080fd5b50610520600480360381019061051b919061a15a565b613ca2565b005b61053c60048036038101906105379190619f44565b6142ff565b604051610549919061b861565b60405180910390f35b34801561055e57600080fd5b5061057960048036038101906105749190619dda565b6164df565b604051610586919061b6af565b60405180910390f35b34801561059b57600080fd5b506105b660048036038101906105b1919061a493565b61661c565b6040516105c3919061bc20565b60405180910390f35b6105e660048036038101906105e1919061a0c5565b616766565b005b3480156105f457600080fd5b5061060f600480360381019061060a919061a579565b616c1b565b60405161061c919061bc20565b60405180910390f35b61063f600480360381019061063a9190619d32565b616d5f565b005b34801561064d57600080fd5b5061066860048036038101906106639190619d09565b616e76565b604051610675919061b60f565b60405180910390f35b34801561068a57600080fd5b506106a560048036038101906106a0919061a0c5565b61701a565b6040516106b2919061b9fd565b60405180910390f35b6106d560048036038101906106d0919061a23b565b6171bb565b005b6107037f26356d783de7752a7d46894d2231a693e850a456a8a1b35bad3b86e288fa1c8660001b61111b565b61072f7f4f704a41b0765fc86119168d3cd70dc331d5885df28210e68cd3422e2dba657060001b61111b565b61075b7f8008cf59186da37a702172d2a68da550665ab7079a2688a7e09c6414b41dc9f860001b61111b565b600061076682611232565b90506107947f2f29791bf11d16c4cf4f6d903c79c100853ada232d85bb22b25208be3d11201060001b61111b565b6107c07f7080d7d9f1e73ab4ac4f4c1a5c05f0b91239c4f850f5ab42afe443f8d8ffb1e460001b61111b565b6000600267ffffffffffffffff811115610803577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156108315781602001602082028036833780820191505090505b5090506108607fc3d6005748efa9859c707d5f2130f56e1b875d874e5884a9bee40dbb76d286a760001b61111b565b61088c7f770b5bdf045ba8beceabb20836938f3d35a32684cff78974f5f6bd843a186a9460001b61111b565b6000816000815181106108c8577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101906001811115610908577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b90816001811115610942577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815250506109727fc1b69b8f1f55f73b39229c38f116aee7ab0f04a8e658048b1e01082613ced6c660001b61111b565b61099e7f8433e8e7971718971902a67f108a62ab5d4f388d5d7db2aec9b7f0b6c16358de60001b61111b565b6001816001815181106109da577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101906001811115610a1a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b90816001811115610a54577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81525050610a847fdf62a77585d60d6f37a341028231b43ea52d2a6cc794e7b028e84c0042fe673660001b61111b565b610ab07f16cc541b2d15b31da8c5b2b97b6de1863383ed0783ed97ba4145fe56034dfb3660001b61111b565b6060610ade7f9b3438273d6f9d626078fe4b69d27637c302c86e3e548e6ed8d5b635e87a09c960001b61111b565b610b0a7f99f5099fa0b3c273dd0620f2f191f0774b1fb5c7f995ace78db2fda5c599d28160001b61111b565b6000610b387ff09b53a0a2a1c2486bc23fb0c18608e6dd37e18669ffcd956cc3c5585f564f6560001b61111b565b610b647f4d55d4c3a1605a468d4cbfaaf72cb0893f6bd07e21942cccdef6ae18b7e3558260001b61111b565b6000610b927f9091c291a8e08e6d719e8f247574d23ffe079242fdf0e8aaf23f25ca59e1163e60001b61111b565b610bbe7f60d3fcffda0280be21926d2818c0b0ffee6442fd15171c0d7de50d5d596579bb60001b61111b565b610bc9858786611486565b809350819450829550505050610c017f136790c5fa4ba4912ffd4f636547f05e893d8a1c3345a06cd04b36e5400bb6d060001b61111b565b610c2d7fa3e80a2f49bb0bc416d971a6f1c659bff6a4882c7b60b28e7d234b3bdfaffdd560001b61111b565b60005b8351811015610fce57610c657f7b27dccaec691ce2febdd9ab5f017c587d4f2a20c49e0e627cda9ed433cfad5d60001b61111b565b610c917f77fa937ba401ecace4458765dc245854c2b754182e154457cbf54d702afebbd260001b61111b565b6000610c9c88611232565b9050610cca7fc6fa121bb4e5d99a66cb968b29185399e610d8bb8d9f54945475bc73197f457960001b61111b565b610cf67fafd5e5b7f953877b459b6cafde79e61559650554415f436ba19d95ac49f9259660001b61111b565b60005b8151811015610ed257610d2e7fd61abaa70680e1cc693c748cb38bd64f7da28cf29fb369ec8b8b4d06b41d1f1060001b61111b565b610d5a7f9b857e82b9aa3ebeb394bb6ca24411322f067f7a880c25756fd89723b44b671f60001b61111b565b858381518110610d93577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015180519060200120828281518110610ddb577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151805190602001201415610e9257610e1c7ffdba164df587e544ccb544ab009714ffbf89d26d7ff2e5cceeec4513f53ffc1b60001b61111b565b610e487f3c4e18d46fdd03d312c191de97ca5df87e6c44f7077c420f09887a2850bdde2460001b61111b565b818181518110610e81577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001016060815250610ebf565b610ebe7ffaeb77e2a4e0b0be7df33a5842c07d9b0d0e5ddac00d4f775ccecc8baa3bc9dc60001b61111b565b5b8080610eca9061c303565b915050610cf9565b50610eff7fc98f4ba2dace4e23d07a428c2bbe1687a33547ac03ed5a895eea90aaadd2800060001b61111b565b610f2b7f1c88b253dd40748e8591b571c4f3c8417d9ff8a0349f990bdf88e3175d7d5b2d60001b61111b565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a0ad3d0889836040518363ffffffff1660e01b8152600401610f8892919061b586565b600060405180830381600087803b158015610fa257600080fd5b505af1158015610fb6573d6000803e3d6000fd5b50505050508080610fc69061c303565b915050610c30565b50505050505050565b610fdf6181cc565b61100b7fa97fbb64ed3e4ddc01bf67bb8409fbe9f1f12f761f697363daaa205c42ee843060001b61111b565b6110377f536573c05ebc542079903ac3b9c1c43c228f7d76dcdf70842b202b4a9a26487c60001b61111b565b6110637f7999798da1e1b9f6aa9626e4759a53a988806d4b7e1fbd9112a0f13f6f88eaed60001b61111b565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663178e4dc08585856040518463ffffffff1660e01b81526004016110c29392919061bc3b565b60606040518083038186803b1580156110da57600080fd5b505afa1580156110ee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611112919061a50f565b90509392505050565b50565b61114a7ffd4cd084a07f6c1c212a222143462c0f7b514ff09ededdc04864725b23601e3b60001b61111b565b6111767fff7a76eaa71099f1a727f622c860f12d7ddd1387705ef60b25a368e7c6119fd460001b61111b565b6111a27f8a2362af76dd53c33fe8921de4bfc7f151ee55d26d9d6f18b4ebda0092ac72f360001b61111b565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663207e33be826040518263ffffffff1660e01b81526004016111fd919061ba9b565b600060405180830381600087803b15801561121757600080fd5b505af115801561122b573d6000803e3d6000fd5b5050505050565b60606112607fbd7f21ad5fb3f1bd6281afd3a9e81916b9167e717c7bc0599f2bac337d55101860001b61111b565b61128c7f9d516223f7242ae0a07fd2ddacaf49110b3436c6b6901d12a473404fd17c138360001b61111b565b6112b87f9419561f7125cd17c05d20fb6c9f7de53d7a9fed9727937522cf5a69444e8fc860001b61111b565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166328a8565c836040518263ffffffff1660e01b8152600401611313919061b56b565b60006040518083038186803b15801561132b57600080fd5b505afa15801561133f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906113689190619e1b565b9050919050565b61139b7f4fdc36a78ab025dff86e83bce3ae52158ca976b996c041dcec4484f9ae5fa59f60001b61111b565b6113c77fa7959443072becbea7b1790c63363a55a051081c57afb4928a87b910f92adec860001b61111b565b6113f37f21b900bdb6dc906738d90466d1632c993a2ea17ee9fc432061b536e388ee6b7060001b61111b565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166334fddaac83836040518363ffffffff1660e01b815260040161145092919061b5b6565b600060405180830381600087803b15801561146a57600080fd5b505af115801561147e573d6000803e3d6000fd5b505050505050565b60606000806114b77f8069413918a946633a8b4deaff1120e2b798eb0338c16f7f04cab7d3a9c2464c60001b61111b565b6114e37f10a741fda6896f17e52346c695440a0cb22f1fe97b3da85519eac6f3b2547fa560001b61111b565b61150f7ff56feb6e9e2d24f131ae29b928c97532e2565afcfa0a3c1636a7d1936bec8cde60001b61111b565b6000865167ffffffffffffffff811115611552577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561158557816020015b60608152602001906001900390816115705790505b5090506115b47f82620880b9e5f2948bf667629bbf321fb3e185453e20a0fcad7a3eb1dabb846660001b61111b565b6115e07f699139cc03d052e9165f1b1c088e9ab3d2eec9bea0d7da1d210fa23697b6599d60001b61111b565b600061160e7f38d2a52cfd3d10b4aa352e2a91765650c583997a4b2a9db2c5339461317081bf60001b61111b565b61163a7f8f878437cd813a53772286d326a7647e85be3c242f59fc2a2562f59c180f602160001b61111b565b60006116687ffb501f2bccbe4d4a52a86af6830f03882cb763aa153c1810f2f11d40fcf76c1e60001b61111b565b6116947f1d617871948890fa732f8b0a1f7267f3f1a9a7d5d9e6971d7be3aec1815d120160001b61111b565b60006116c27f50f7d00b1c941ecd7c4b2125dca8b3e30af01ab1fca3f7ac9434afe5d4593d0760001b61111b565b6116ee7f96d43dd0ab49be079785640fa8b188b8c882c9c081503946fc0a77706e74acb260001b61111b565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561175857600080fd5b505afa15801561176c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611790919061a469565b90506117be7f796b8e29860dd9ba340f140d66b4b1d4080f6ffbab1f8373ae7cfdbb43ba6e6160001b61111b565b6117ea7fbc4892c915ddc29af56ebe72ca6f1627d0b065fdadc244d5c2cca80f499bf44160001b61111b565b60005b8b51811015611f5e576118227fd177410671a17daa5c5a100af5cc96d5fe8e01f57a56a8addddc125ef7bb631660001b61111b565b61184e7f23589d27ea8dfadc31fb99fdc33273d8b64af2cbbc468ece50935f987f45ebb660001b61111b565b60006118998d838151811061188c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015161701a565b90506118c77f736e7e6317d3db0000c9e9b2cf173d4d699f93ce3db40b1279166b1af11ef7b360001b61111b565b6118f37fdffbc3d16cfcdc27cd15132d167c0295fcd0fc3b71ae3f0dce3c5e06f12c78bb60001b61111b565b60006119217f013c016af8037e8a61fec716811ed8ca3d908a2dc62c75746027b4ed11b64e0860001b61111b565b61194d7ff963fb6a25f3425e5a18179005cb66eb47fa5805573e58901101ff9da3ba3e8660001b61111b565b60005b8c51811015611b66576119857ff7b6f58cb5c6911377cd6c83d3908d22bd16136068d1f3b52cc39a767205f66160001b61111b565b6119b17f0497a1fdecb568782aaf3a903ad8e53da23de09daae2ae4d743b12d63f0a2f5c60001b61111b565b8c81815181106119ea577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516001811115611a2a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b836101e001516001811115611a68577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415611b2757611a9a7fb23521024a916a67741ec66c0ab586e7cc053f1f1b2c5cb4f8dc534d414f972760001b61111b565b611ac67f6b32999acebc05994b86674373cddd87f4a90365a56df064f93a188524ccfaeb60001b61111b565b611af27f74c4ceb4f779e495148b329f67fd5d7b8bd9c741c7a5027aa7801adefc8df32660001b61111b565b60019150611b227f5ea177551736f8935a63635b6446940b2c97e987c9f8357a84f4aa64359d547460001b61111b565b611b66565b611b537f40e5d63282d49249d90651686e768402b090a4c1c274ce411ff7c9154558c3dd60001b61111b565b8080611b5e9061c303565b915050611950565b50611b937f6f44decabee134d35621bd0361f848540c8e00213a6c5ee9db06a4f1201d03bf60001b61111b565b611bbf7f2b698865645f98d8eca6ae41a7830344c374e45ffa0b8e709fc8e96a32d9c56560001b61111b565b80611c2357611bf07ff3781870cabe050db413ac503bb3f98fef91be88f0c971d95b664b1b23ec899b60001b61111b565b611c1c7f75185595b9e862ef90d343b5713155b13c7d1e909d8463fd8ca85b9cc27b587860001b61111b565b5050611f4b565b611c4f7f5a3acb4b904086ebd4a00ade2e8c404c90aa4ecb69c4917164ab8ae3c5f8a84660001b61111b565b611c7b7f5391c7c33545c8be0fdfabe26a8aee09a3bae324ae3c41731958d540a2d84f9f60001b61111b565b611ca77f7d4cc6829167e437517fbe735e531ff8cb5cd62d4e34635a2b84e8e7de834d3e60001b61111b565b438460c0015167ffffffffffffffff16836101000151611cc7919061bf74565b1115611d2c57611cf97fe09da45012ea0b4b0b4539966fdd0115241949f964e2d23fc73fa08a393a962c60001b61111b565b611d257f863959c3727e0c82bf9a4d078b44f82fddba4976ed0ebc3c2fc1eba6831cd42f60001b61111b565b5050611f4b565b611d587f54dcff8d6eaa569c9d0c047ffb5d7fec5e6be28acf99408e7ad181e0f1ca74c660001b61111b565b611d847fbc4e18fac22be2a119cd465c245c20c259463ea5faf978e2b5c902b18113260760001b61111b565b611db07f5cfebeefc37cf01ec75ba2acc712bb299737e37b812901961d985fb2dab11c0c60001b61111b565b81610140015186611dc1919061bfca565b9550611def7f9ade799a8daaddabaa3733eb3ff21cea56af7c76e195fef5d4676516cc2a439860001b61111b565b611e1b7f5289d3b55c2b6dabf2baa9e42cf7627056dc3909a0367f69d23cf459adf11a0e60001b61111b565b611e2782600180613747565b611e537fc459527952f80dbfec428050c93e2038f9eabfbfb2d3a1ff02bf521b7016c61d60001b61111b565b611e7f7f911b00e92ebe68906b641c943693eb72b443816fa888a71364e2539db34f8ada60001b61111b565b8d8381518110611eb8577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151888867ffffffffffffffff1681518110611f03577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181905250611f3a7f11b076df21be5d6015a4ede58644db5f1fe7816be62a35ce37e6df1a7b0f073860001b61111b565b8680611f459061c34c565b97505050505b8080611f569061c303565b9150506117ed565b50611f8b7f0a0e1a953412786a860101505f181b9852a83a7bc5205d8ea591ce40f9a27d9060001b61111b565b611fb77f949fc0e9cdfe968be7cdccd740c0b06a299082422a8ef519ebf6599798dde02160001b61111b565b60008367ffffffffffffffff1611156120a457611ff67fb9a48efac06c7b393f3f14c8a2c10d2cbcfb56045c554b9dba4bc9083abe444760001b61111b565b6120227ffbe160f6ed79ecf07b1cae6f1d5b750c05d0579de87e6da5b29e72bba293df4e60001b61111b565b61204e7fd8b78b1220a681b62e9bc67ac8a62928e80daa0e061cbab261da112f2a71c0bd60001b61111b565b8973ffffffffffffffffffffffffffffffffffffffff166108fc8467ffffffffffffffff169081150290604051600060405180830381858888f1935050505015801561209e573d6000803e3d6000fd5b506120d1565b6120d07f5a6ca6ed7b8464f67b583d6049ec6df204c56952cc70e0092a4e6502b3e21d6e60001b61111b565b5b6120fd7f4fb6ee5f6768fcc36096fedfc7f2720221b085ed3cddc06418e854073aa255d760001b61111b565b6121297f634e45f7f73fefb36bac3079f314da56619453a351fccf04f2126ad68060aea560001b61111b565b848383975097509750505050505093509350939050565b606060006121707fa6c53a4b48094df66d51e82b545158527d68f97869d09b3b3eb98130389fb4e760001b61111b565b61219c7f9153531b7826af85a59b761bf6a15ceb3df117757c17aa306e1e339400afe99560001b61111b565b6121c87fc3d1e78c40085da77334f6038206bec265e8596ae1189b786ff4586fe9aad74460001b61111b565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633f2cc9a08686866040518463ffffffff1660e01b81526004016122279392919061b631565b600060405180830381600087803b15801561224157600080fd5b505af1158015612255573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061227e9190619fc6565b91509150935093915050565b6122926181cc565b6122be7f4e89669487f343bdb522bca44c96b8dd2b73874293145f62598f16ee1ddc7cc360001b61111b565b6122ea7fe12956a7ea69f095c0ae99d61ec936e806c473810f85d2b83406a1a3534f079160001b61111b565b6123167f188ccca9118e6da1a83be9bb4764144567b25c6201e053c3064fbb323dfdf4ca60001b61111b565b6123427fce377d6b9eaff600a7d94f52a9996d49b269553314bcf07574b4918613b8766f60001b61111b565b6000826020015167ffffffffffffffff1611612393576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161238a9061b9a8565b60405180910390fd5b6123bf7f68652f6b7cbfd4a46d37a0c875e7ee56998a7e3aa9be5318cbffc633d4a989bb60001b61111b565b6123eb7fc36490bf7e9d045ab5f8ae28d5e0d42cdeaec296a44b49e90be4696084630bc260001b61111b565b6124177f72cd0fb4d532b6acc4af3ccc503bd0eb4c8711723ac3bb43c257190b20816d4c60001b61111b565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561248157600080fd5b505afa158015612495573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124b9919061a469565b90506124e77f8631c745c8a9c9560bbc31d74e8e722e12cd26930d5c339c4e907010c8dc89cb60001b61111b565b6125137f58c3f7b744edc57f1facd828c3bdfaeca5656e8a32416ad4fd14fab4423ff34960001b61111b565b61251e838243616c1b565b915050919050565b6125527f784e842fda899a59fa9042c10a34dacb86d43b9ad361c42c0cd26e4dbbe9490560001b61111b565b61257e7fb15341980292161c3d7e42c78e77e83af9d6b9fae7fb2d67906bc48f93757e7f60001b61111b565b6125aa7f211c60f2a19a0e5509626d882be4211011fa68965114a66c714f222c80f77fb360001b61111b565b60006125d87f9fc7b9410d544159d50e223e1cbcfcd618420b25675294df8a31c3f60cc1220c60001b61111b565b6126047f7ed59e16900736260816dd7b49d31679276e7048ef7ca2f9e290f7500cee15e260001b61111b565b6000825167ffffffffffffffff811115612647577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561268057816020015b61266d61820b565b8152602001906001900390816126655790505b5090506126af7f30d5210fb64b85d7e5c57fe677f0964688b0896b009e08ed56ed1de2ca16f67360001b61111b565b6126db7fc79e2bb05bab63d66cf2f62e2b48d23599202a5eb521ab74bcffe3f68203ca1f60001b61111b565b60005b835181101561289c576127137f758f46183f51d52833b5974d264560cc6b2f760b1a502c5eecfd3eea8f4b6e4360001b61111b565b61273f7fe9a32bfb4e22fe769681bf0361700b30d11929ccfb7f9612f6e02d1d5f353ba960001b61111b565b600061278a85838151811061277d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015161701a565b90506127b87fe5b7ac972b584db469bdeb8dd3eabec7d9013825b1e0ce635b50c05a3656fb3860001b61111b565b6127e47fe76434c95393533222e3fc444e22ce9e92a3d454ddf6967ec7b27fee18f50fcc60001b61111b565b806020015193506128177f9165392f205fb3f3754456b49ee6c83abe365c20635c4898ceaf949efa41a5bc60001b61111b565b6128437f10aa042bb000489fba945db2f119ac75f427012458707dfce07fc814c0b3728060001b61111b565b8083838151811061287d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052505080806128949061c303565b9150506126de565b506128c97f5f6a3473742f06584e71c21840ef116dcff7b5d93c506dadf6126b9b0f3c0e3e60001b61111b565b6128f57fde5f95d6b099054d29dc702b405a20ad0edb309cd3ea3c4ee09e5e8cc2699f9b60001b61111b565b6000612900826142ff565b905061292e7fe197a2fd4143194ffd107a03ddc41f028e2dda62471245a1b0836504226151f460001b61111b565b61295a7f24c8bd42fa84caf499efd2f74f062846a02422ea8c6d94e54009bf8e23658d2760001b61111b565b600081511115612a7e5761298f7ef293bbc27b52f6f82c0fc7618adfc592daf50cb5c79ee86ce0f09f807dd8f460001b61111b565b6129bb7fdeaf9c600e8f2b38d27d992e5833432c8752c4dc0bfdebb47efe66258c8e3cf160001b61111b565b6129e77fc2b9c24d841168b2a61c6b1042925c8e191d1c3af5c3e2b9df8f0fed5b1e792b60001b61111b565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f581604051612a16919061b953565b60405180910390a1612a4a7f3197d40d2e7afdfcc3a7467b5cbcbacf27faad1274672e73c6e6b46b7ed4d33360001b61111b565b612a767f9573a774254f2fd3a73f1782242a15e38cdee763afbf2c150594120a02e3eae360001b61111b565b505050612b44565b612aaa7f3935ab56406c645e890289ea822c98c38d69d3e3f00896dacb7dc999564f38f160001b61111b565b612ad67f2b5e1795e2701b3b1ad6392dbe3acc705f33025d27d4ce41fe22ee2300faed0560001b61111b565b612b027f8ea639e76c0cc11434026d984d76e441100544d62b3aefa425d40d76bfe70e8e60001b61111b565b7fa5d0a6140e999641864ed8958af5eea1d36d821658cfc056552962fab40291ec6001438686604051612b38949392919061b753565b60405180910390a15050505b50565b612b737f193b1c43c6e6aec1e8e916177c0b49f1c889b65ad242af56fd8c2cf1d28640ab60001b61111b565b612b9f7f2bb36c1267d4a41a8e90a15d516f2f93d483c42795f44ccc9dc5244043fda21260001b61111b565b612bcb7fe296bffab4b2e2425344b805a49d4bc0093e081e1263fb84fb66271cd12845eb60001b61111b565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166395b0b54b826040518263ffffffff1660e01b8152600401612c26919061b701565b600060405180830381600087803b158015612c4057600080fd5b505af1158015612c54573d6000803e3d6000fd5b50505050612c847f7d421dc844471ece90b9a43f016c1e23eab0ac1ad959d17f3b8fe0768b89791160001b61111b565b612cb07f42de7111d7bca2bc3f2e6ab12e378f19c525ed6e5a8579ff4b9d603e9daab53c60001b61111b565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166352e06146826040518263ffffffff1660e01b8152600401612d0b919061b701565b600060405180830381600087803b158015612d2557600080fd5b505af1158015612d39573d6000803e3d6000fd5b5050505050565b612d6c7f16a707eb5ef8e43eb64b6cebddba55779b015ea04ca224a2affcb1b90f5eb93260001b61111b565b612d987f32a7cb81cf638f667044637f588e83a3a3b1c5c31144859711599936d9cb925e60001b61111b565b612dc47f0580942fa82566e67293a2703ecef065ad4596cff716b64760c96c2c0097389060001b61111b565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b158015612e2e57600080fd5b505afa158015612e42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e66919061a469565b9050612e947f4c17339a3e7c6dafb4320ef8ceebf44dfee5e829a2fade932afcfae0ce7fb08860001b61111b565b612ec07fd31a87f290a190f5617c58e9c83a8a3352ce94007e3cc68bd4df31d65e00b71560001b61111b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e37891a3484866040518463ffffffff1660e01b8152600401612f2092919061bb31565b6000604051808303818588803b158015612f3957600080fd5b505af1158015612f4d573d6000803e3d6000fd5b50505050506040513d6000823e3d601f19601f82011682018060405250810190612f77919061a1fa565b9050612fa57f5b36d0e324582cf41859fa06082ffa9fc0205d232fbec8f4ff3d44b908c4bf1f60001b61111b565b612fd17f225899894222b0c2f7e7c6414276fb176d230202a53da041325b79c6ba09d47360001b61111b565b6000815111156130f5576130077f106064cf049beb22e8878ca945539639ff2be7f4180c76953de82c9f7e14efaf60001b61111b565b6130337fcf3074887d0e83d1d07b8ee9f68f74ffd08109575ad99c87157f0ad71b5deaf960001b61111b565b61305f7fd3074b316838bb728182c6994162f1a00121b85252a664ce072e74ee17018fe560001b61111b565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f58160405161308e919061b9c8565b60405180910390a16130c27f688e23e6a63a76362900ce48e4349b84dbf1fd6dd45c9f54f75b44d93b073c7f60001b61111b565b6130ee7f6e06f9ad50dfd88b88db2cfe599ce025901dc3b904e6363e577f865546041e1860001b61111b565b5050613124565b6131217fc382e0520b0bdeb1a0ff8c20bb2f2744f9b9f37b5e2421ca1ec9a464e082ba9860001b61111b565b50505b50565b6131537f77cc859894ffc6e09ab4c99f0611172a6b47611ef443c2c11db24684195304e460001b61111b565b61317f7f0e82a1fca189aebbbf56a1226f6e4be354b884b652287316e903afcaf73a568560001b61111b565b6131ab7fdcf023669476bf92ef432590b89ab4981741601995d610dfa47404d69a1af52e60001b61111b565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635a0e748283836040518363ffffffff1660e01b815260040161320892919061b723565b600060405180830381600087803b15801561322257600080fd5b505af1158015613236573d6000803e3d6000fd5b505050505050565b61326a7f6a49e54e278efbfec2ce9ba30bc856ca78dd6dddfd1306dd4b09306d9bbfeefc60001b61111b565b6132967f1a49c3fdbedfa25f0e7ebc0b97dbc49da2ab80963cec23d7d15a17f0c755297460001b61111b565b6132c27f636a212783f91e89984acd4d40cc98b07be9ec0d2046d7ef626046683a7ac4ee60001b61111b565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663681850d7826040518263ffffffff1660e01b815260040161331d919061b9fd565b600060405180830381600087803b15801561333757600080fd5b505af115801561334b573d6000803e3d6000fd5b5050505050565b60606133807f57bbb7e2d4930138d63b21fbc0aa1ae60ae5bad4702a57f9c69d1a0833e71e5660001b61111b565b6133ac7f39cecb4bcf390d57e16a108c1ad1c18f41756ac60400ab4e27562631d1cb6f2060001b61111b565b6133d87f3dd4b96a0dbd56dfc8518e5e667ef395671c1268b08475739f82528f42ec3d9e60001b61111b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f161a50e846040518263ffffffff1660e01b8152600401613435919061b56b565b60006040518083038186803b15801561344d57600080fd5b505afa158015613461573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061348a9190619e1b565b90506134b87f2ba6bd68eb68b9607a7c9e45456634f850ef02ea451c0c564847938e9619120a60001b61111b565b6134e47fe4ef5c2e61aa738c844edacd9c81e650986983ff798816aaf12519f5bd0a6f6360001b61111b565b6134ee8184617995565b915050919050565b6135227ff9a5f618bd0933b6dd5fec709b4fe427a743e14447453fd747f9fc02ac0fe08a60001b61111b565b61354e7fd8ca8379967b7099c916a5468cef939e506d3855a7f061c9b8061dd00e2fd75260001b61111b565b61357a7f594c8e1e6734497eb2834d8726391ea58da6e07645cc9c0ae358df52a8f4641560001b61111b565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166395b0b54b826040518263ffffffff1660e01b81526004016135d5919061b701565b600060405180830381600087803b1580156135ef57600080fd5b505af1158015613603573d6000803e3d6000fd5b5050505050565b60606136387fe5df031cb514f8b192cb54e2587adb22896b80ac4f2fafb0b4324a8c6ed8cf9160001b61111b565b6136647f75b49e645798e93f79faaa68f7d3680978f9575d5e0d2170cb445641788dda2360001b61111b565b6136907fdfddb20503d5804f43904fa72bcb4e32554c3efba79dda3a12608d74aeb447d160001b61111b565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639a2b5e63836040518263ffffffff1660e01b81526004016136eb919061b56b565b60006040518083038186803b15801561370357600080fd5b505afa158015613717573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906137409190619e1b565b9050919050565b6137737f772e6831b2e0d83cd33471ed5576f04a4cdc3f071af5855ef2d6c9f964ef031b60001b61111b565b61379f7f993f030b7589cb30843f0a05b00362416728eb2972f945e0c98933197cad5e5160001b61111b565b6137cb7f258f93f57c2fa778bb7e357c63c7db4d7e94c5f1feb9615110d21df39f52e38260001b61111b565b81156139bb576137fd7f7e1f5d5602914d2b9c39bda0bd3a44d336b8f6b4b84cb27d0018843e879ff13f60001b61111b565b6138297faa0135159e5233ad17645b6a53f1bf76252013366b361a09b740c4b5f585d88a60001b61111b565b6138557fd66304f6dd37ebe9046ba9a88a4e22b49acf1262c2923c8af9bd31b6e0c3061160001b61111b565b61386283600001516134f6565b61388e7f5d65814baba119f2ef19e96f7f71d9d7b4ccfa5dd743b8a835a4e076c4b237f860001b61111b565b6138ba7fe1aa69e0b7d6a496bde94b8585ad3bdb6cddf7eab666f0f84ce1f56acb07d3a960001b61111b565b6138c78360000151612b47565b6138f37f887f6a9663cb6b5a067531954fc6e3a7a55a72d77603a416ba4d79a674ed753660001b61111b565b61391f7f443d9d9b6b937f84700101d06baab7ba427e96ab7d627c6f7843b5ac4edc325e60001b61111b565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b64ab1ef846020015185600001516040518363ffffffff1660e01b815260040161398492919061b5b6565b600060405180830381600087803b15801561399e57600080fd5b505af11580156139b2573d6000803e3d6000fd5b505050506139e8565b6139e77f038bf703b5cf8dd35b85200cdd9e94d55e5c58946b7307d2f46fe6b18cc16a0760001b61111b565b5b613a147f59c101b0e823ad7f8839b84bdb8282e9d6595351822a435bc42365d2f9ac5fab60001b61111b565b613a407fbf2f78ac5e882b8aef216b160f288679aa2cabaace93d53e8d23353b9ee1a37460001b61111b565b8015613b5c57613a727f3de2739ca3f7ca00dd215c9223b9641da83fd1b8062d35262e52a95c04c25ad560001b61111b565b613a9e7ff7da24ef499be2c5e7b2b33556e318c802359fa97cbed8f321fe649fc49760c260001b61111b565b613aca7ff638f53111b02f81514c9bcf601b78be924090d53c69212003e67e75b084d8c660001b61111b565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e457c9d9846040518263ffffffff1660e01b8152600401613b25919061b9fd565b600060405180830381600087803b158015613b3f57600080fd5b505af1158015613b53573d6000803e3d6000fd5b50505050613b89565b613b887f922e7c0263cced1e25c004f92f6a379f960d58745e89689a265fac9223b3ec7660001b61111b565b5b505050565b613bba7f5a4abcc7d063adc28719d2fc6eb9a81693da6125e5124aebed461613af0ecc0f60001b61111b565b613be67fd60e34cb7fcccc4ee77471fc8134273fb72a32a1504c686f3c01fc5ff780b78260001b61111b565b613c127f1d4ac84624f7e0b5080aae2302090d06318199f2ed098202131210df5e47e0a060001b61111b565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b6af10fb826040518263ffffffff1660e01b8152600401613c6d919061babd565b600060405180830381600087803b158015613c8757600080fd5b505af1158015613c9b573d6000803e3d6000fd5b5050505050565b600060019054906101000a900460ff16613cca5760008054906101000a900460ff1615613cd3565b613cd26181a8565b5b613d12576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613d099061b988565b60405180910390fd5b60008060019054906101000a900460ff161590508015613d62576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b613d8e7f01b3736f6eff7ddc8f6afb13cd8e8fcb117db2fd846ea51e0490529f7b15d42260001b61111b565b613dba7f5e10819a111e32f7c6a7a10c3ec67f71f90caeefd43cba9f5cd23a52e8ebcf8960001b61111b565b613de67f32256303d8f980ef24d10cd93ffea002a8f2d71541b53553a0d635237d9eea9360001b61111b565b87600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550613e537f4b74fc60260d6ce80bf02547814dadee7f0369d3feef880848a7233f41e2de8d60001b61111b565b613e7f7f1d053d424193b161d01535461153af5157626738f02a945c3a3657b1b592b6d960001b61111b565b86600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550613eec7f330b35d89173bba704284c38b411bdf0293c2e9dd656c1771dd202fa869a619360001b61111b565b613f187f9e0c7a561877f9020d3212ab515baaa6d38b6da890097490709ed0792774868d60001b61111b565b85600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550613f857f817deb35fdb1ec9a451624ab4c27edd95c252847266b91f7137ec98ade40194b60001b61111b565b613fb17fdf49098e65948843b39a64aa0401e3e18a3d8cf3cdf1c3d3107b9f78253d73cf60001b61111b565b84600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061401e7fec2341b2521469c34b6bba6500c9a97f9bb342ae51072488a458be332ee194b760001b61111b565b61404a7ffc30cdc392986ed4bddcf7b50495aba427116f565c458e0fa7ed118962ad9d9060001b61111b565b83600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506140b77f1e28ec252f42fb3be0e33efc41e3c7bfc1a2edbf03ad9a156802026705699b3560001b61111b565b6140e37f4597b4cd4587c452a3bf67842096c1a35e3970de842709497fdd1fd364a3274c60001b61111b565b8260000151600560146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061413c7fb759217f8ed5689e96c04f6c9923d48487c439672f070fd0a6087e10f5756a1860001b61111b565b6141687f12cb69f2d55820a03d1c89c149391a9c64e5f3beb40cd73bc46548cbfcf9589960001b61111b565b600560149054906101000a900467ffffffffffffffff16836020015161418e919061c008565b600660006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506141e27fb2b1003e6801fa7923697df5b6a6a4db4adc7c60c81840ed6a34c2543c91cea160001b61111b565b61420e7f64cd6c86d2fd248b1134f3f13d6f38bfdaee1970f35ec32e596fa0cc0e46e66560001b61111b565b8260400151600660086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506142677f8bb509bab61e59b8256dfc6d2d2d78ab5326fa224db1b712f20c58989218ade360001b61111b565b6142937f6e4d2e5436388070783987d0288565eedf9e561c166929f90517f8c8ca9dbc5460001b61111b565b81600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080156142f55760008060016101000a81548160ff0219169083151502179055505b5050505050505050565b606061432d7fa0fc9958fcf2ea0bafd5b153f3d711dea349a208f2239207465da361695fef7a60001b61111b565b6143597f6b35c7284a8f40fd69ca0c00ec80f477d5f9b614345fc31db026a721ef55c59860001b61111b565b6143857f06dfd88a1b0ad8c927a122a5f4647b7ee46561ed3492dc39080fbb4d56b51cb060001b61111b565b600082511415614450576143bb7fd7352f1cb1b8dc84251f7eaa86f75434a3616fa63f7322c25dc24799c86efcf960001b61111b565b6143e77f9ad6611d88687b7f0a3a4b58ad2e3491b13822810c221e7304f936fb8c11ea1d60001b61111b565b6144137fc12857feb9abc8fc2aab5dec394e331601de004049761a3aebb49253dddda5f760001b61111b565b6040518060400160405280600881526020017f6e6f2066696c657300000000000000000000000000000000000000000000000081525090506164da565b61447c7f23e53fa17505563a154f65bd8f23427f42df39984b15fca5f2c8fbd9066c3d4b60001b61111b565b6144a87fb776e427309acd77583e84f0561c4c176d5961bf4dba6af3f7c4b4d7b20a4a6760001b61111b565b6144d47fc028de4b2bf55925a85fc29efaaa930338a20702b2d7c70a40d9a541ba6221ef60001b61111b565b60006145027f9924d44aca9d23b180a07cb950d6c9e9572de81694ae92b51dd98bf23c98eeca60001b61111b565b61452e7fa5fb9536ab322519356ebd2025d1293c4eb3bfd0cafb387ffba0f6590a7f975860001b61111b565b60008360008151811061456a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015190506145a47fe1a49a31b5652c9e5a51147b263d4f9fdef7c594a001132e190d9e455e4e5a9260001b61111b565b6145d07f53fb7283bbdd0f07365b1dbdb390c61d5b7c7829b524f8c2d0731b3afa2da20760001b61111b565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561463a57600080fd5b505afa15801561464e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614672919061a469565b90506146a07f5e3b7578d113eae42a300c360a007a45aea8f77df05d2a34d9e962a00be9037360001b61111b565b6146cc7fa9b09ac123429e28343893c20ed6aaf1949366bce30847681e82772d127f704a60001b61111b565b60005b85518110156161b8576147047f3adef0c9bc0fd75636ba8c2eb3511c893f1caba5c68b65a26268b58f4eb6325360001b61111b565b6147307f321e2112d8eb816b4a1dd73d01bb7f9a9bc712b747a295c244ea105402c8bb4660001b61111b565b600086828151811061476b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015190506147a17f28de12618d4fe162e1d5cbf445e08b3847760213f9653d7a97a491acd4cb844660001b61111b565b6147cd7f54937c087ef0219be3d141819f5874762f19088f081a6bba6108a4ef1bfffdab60001b61111b565b60008160000151511415614839576148077f9962cf27adb32f95121d30efb1cd5072ea466cbfa85ca1ccdc333ab550be06e460001b61111b565b6148337fad28abdfb1ed096d3c4aa19e56ddf433e83decf26d89941e96e20d78cdf2870c60001b61111b565b506161a5565b6148657f58ae49a730331db5999c5289912411466143a0c4bfca8117197d07c25cc580e860001b61111b565b6148917f7dc9bd9855b25855f8406c5a670fef50a0f296a98fed98705d9cd6d47bb58a0c60001b61111b565b6148bd7f822f44a3145f2faf3cf2f3c5e17c8b2a5bccfcc92f1ce4b248f68b0d2bdb52fb60001b61111b565b8373ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff16146149ba576149207ff07877f616bf59b1b8c2ff5fd7087f55779a8d61debf2b182d0766d150ad337e60001b61111b565b61494c7f5b559063798a48dd67b31baa27e2bed38cdab62f2a9b17c1a23c18673dbab7f660001b61111b565b6149787fc05f8bc4537cb0bd95648b4d5e59ff7205620607ecb16c3790a803ad000ca78860001b61111b565b6040518060400160405280601781526020017f66696c65206f776e657220697320646966666572656e74000000000000000000815250955050505050506164da565b6149e67f7b7d3376e453b6d9d8b199f77c9d16cb3a89d9da2208d4c317085388c2eb46e560001b61111b565b614a127fb1c291d6832998efc7c4e3bc118438f032cda6248d8b1ffa5db172abadcf44a260001b61111b565b614a3e7f6b7cebc6e742c128b11940f7a4c107a7b2efbb28c5b2733dbfe8f8410b91a9a460001b61111b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634968e2d683600001516040518263ffffffff1660e01b8152600401614a9f919061b701565b60006040518083038186803b158015614ab757600080fd5b505afa158015614acb573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190614af4919061a05b565b9050614b227f67f0650fbea054a1d38531b367af96413485d449077a186a19b9071d1649f9f260001b61111b565b614b4e7fa4269a0c6ccdababbcf4add1b59804a5b7147c87972c8142bf8aaebda4a235f960001b61111b565b60005b8151811015614da057614b867fa1c44da533e0d49a06976da41b0b9e56d5c68b99ced8b56f54b0b940cbb83d0460001b61111b565b614bb27f6dd3e75f224fc9d396c24aa8ddb294c5f1952ab80257e3f5ca5f6fdb956423c460001b61111b565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d7848481518110614c2b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b8152600401614c4f919061bb16565b60006040518083038186803b158015614c6757600080fd5b505afa158015614c7b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190614ca4919061a428565b9050614cd27f9b1fa0c0cc73c74760d6b4aaf582140ac17b9a8837c7230eb824c29de48d50de60001b61111b565b614cfe7f8f1578b5f720282de86b829a38371614cab2125b71d901e246b0c29b4bf9a51060001b61111b565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166247c00382866040518363ffffffff1660e01b8152600401614d5a92919061badf565b600060405180830381600087803b158015614d7457600080fd5b505af1158015614d88573d6000803e3d6000fd5b50505050508080614d989061c303565b915050614b51565b50614dcd7fe67714f4c4ae172d105a27d4c91563cf56547ab3c319822b338ac576ee2f0b6a60001b61111b565b614df97f036aff75003a99e416617c33665857a13b1930ebdef15d6720980c59ef161b0460001b61111b565b600082610140015167ffffffffffffffff161415614ed457614e3d7f8e10b7321a7f7d84781e91c4b9c6361835780e22e7e9c5f0505da3f3c5a720bd60001b61111b565b614e697f4d7501e025ba8ad77712f5db2b94ec878c37c9f0292e1d0af62ae343f6917aec60001b61111b565b614e957f0b66cdca9b63603415122aa2cba6c919613e0eb2616e030be67d4d3ad92fbc7160001b61111b565b614ea182600180613747565b614ecd7fae3bdb1e7ebe70f8e107f27a1b5b7562894192663f6f6584ebebebca279da84e60001b61111b565b50506161a5565b614f007f337d2ed0b74c8647ff9988e5727914d436af27016763126fd3ccb8cc9e5c907e60001b61111b565b614f2c7f6fea5a316ee47f37a31e665057ed53bd1754de17fabd9dbf1365d4de9b893f9c60001b61111b565b614f587fc2347d274aa5ad98ec82ac8e29bd0aaa1fa573a99107ede80de193e7f5178f7360001b61111b565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663977fdfe284600001516040518263ffffffff1660e01b8152600401614fb9919061b701565b60006040518083038186803b158015614fd157600080fd5b505afa158015614fe5573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061500e919061a01a565b905061503c7fcf814f086163941db00bc5e76f8b5ce36e3867f4d1b60ad9f3792e1500fa264860001b61111b565b6150687f56beb726d63435555573c3285dd9d782e4fbdc622daae12ea1a706c22228544760001b61111b565b6000836101400151905061509e7f04170e1c81f456e45a278a3c89eae2563cc94f18fdc7831514af17d7310c050260001b61111b565b6150ca7f7059907310d13527ae2e2aab409719e3111beeb5bf7678d19974aa2aa406d2bb60001b61111b565b60008460a0015185608001516150e0919061c039565b905061510e7f42710b87aa1b9fcef871c180eef79368d471d4a0c125d32c9852b6dfeb560bce60001b61111b565b61513a7f997039a1f5b3eec727f41461c4f08be7cedd15315236b779e58c24a57f7949d560001b61111b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ca6142cb89846040518363ffffffff1660e01b815260040161519992919061bb63565b60206040518083038186803b1580156151b157600080fd5b505afa1580156151c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906151e9919061a60b565b90506152177fda285427af4b0cb8040a63ab18c34ad7887374bea47afa9220c7a7ef4b88679260001b61111b565b6152437fdd6a881aee0687fee55b5b6e536cf1105b2ee635c56114fae921e9c16e9838ca60001b61111b565b60005b84518110156159335761527b7f38ae8fc3a1dff4b82092afe5da249e8fdd737799ce363eaeeb1a6a8370bc919360001b61111b565b6152a77f2b37b211acd131de6ab3ad4cf1da081e0f5881bea13650ac41b9e4d91691403260001b61111b565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a65374a878481518110615320577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151602001516040518263ffffffff1660e01b8152600401615348919061b56b565b60006040518083038186803b15801561536057600080fd5b505afa158015615374573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061539d919061a365565b90506153cb7f663396fbffe6654d30db6476345bdafbec82267a26d09bda0c593e595378b30c60001b61111b565b6153f77f9f12c4e4da357f41bd044f15dc8c7ce4af272325f60f0943fcd988027b2aff3d60001b61111b565b6000836001888581518110615435577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040015161544b919061c0af565b615455919061c039565b90506154837f39ec0fa728a018b44e8830c59111801b66c03743362ca4d8438c1dc782207e4060001b61111b565b6154af7f4c09f860f2dbd2cc0809471da21353ecc638d0d61dd018e33681eb28bb7355af60001b61111b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663df49410a8d888d6101a0015143615502919061c07b565b6040518463ffffffff1660e01b81526004016155209392919061bb8e565b60206040518083038186803b15801561553857600080fd5b505afa15801561554c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615570919061a60b565b905061559e7fe5dc8152d1bec4005c2b4f0e016d5bc4198537486ac70ba1754934a58656eee160001b61111b565b6155ca7f2028d31c4555b22dcfc6f1ac8b86a9c8750469df2c1a831820fdb8577f74150c60001b61111b565b600081836155d8919061bfca565b90506156067facbf5741b46c0ea17872ab51408d3de09d930254d521094e91c1269c6f0998d560001b61111b565b6156327ff6ac1ae29ee8a6a40ae9dfce7f70e3878d4b5778edfc6847b3297169ce033e2860001b61111b565b8767ffffffffffffffff168167ffffffffffffffff16111561571e5761567a7ff67814b8a3067e9ab7cfa259f55b098593a28e586e89b4eceb9a6dc1023a693560001b61111b565b6156a67f748f8f46ebdc5aee887cec3e581ab9d16bc35bd80fe26aa1c2527a2ae682422e60001b61111b565b6156d27f6ae951a8b7739c322a3f271fdd6479dd818cc23f63d440e9a000cefcceb4183760001b61111b565b6040518060400160405280601181526020017f70726f66697420697320696e76616c69640000000000000000000000000000008152509f505050505050505050505050505050506164da565b61574a7f15a1793a0e5bf3c678fd8ca75d16179919ecf8b91710ebfab3a4679b7bb1247960001b61111b565b6157767f5ef5e9925e82e918f10b5a0f43889b15d6df8e27474fc747e5644dab2354fd3a60001b61111b565b6157a27feee226e72b8ff3856bb535d7a0b129b431930f7d36616e11ec2675e9cc3b3c5660001b61111b565b80846020018181516157b4919061bfca565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506157fd7fd599af408f861aa802ab7dedff0ec57575cc19020c581601cce19844552b465060001b61111b565b6158297fff44e2b936db3fb821ab3b669c34c01fdbbbf675b2837cf111192b21e0ee78b760001b61111b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fc215358856040518263ffffffff1660e01b8152600401615884919061ba79565b600060405180830381600087803b15801561589e57600080fd5b505af11580156158b2573d6000803e3d6000fd5b505050506158e27f0df477b9373332ed4c126adac860524163ce443e4d39012277ac18c2675a1ec460001b61111b565b61590e7f4098e863a84cadaa02f3afb6797323002190e7faa560875dd12de5d52a6a73b460001b61111b565b808861591a919061c0af565b975050505050808061592b9061c303565b915050615246565b506159607f457ebc0616b32d5f2bc80bd4b1bdceb11554be83c7f41ad1bcfef5698cfb9f6760001b61111b565b61598c7f1c92d66d1e5fb59e973ba6ed08dfcc4f41825591853dd9ef5dbd87a6d67773e660001b61111b565b600060018111156159c6577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b866101e001516001811115615a04577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415615f7457615a367fbed5132d708bcb0b85b457937b95c3ff22379ea0413eefd3477787b863a7096e60001b61111b565b615a627f7ead1144687fb7745cb02b96667600e168bea3306d750c7ab23ad474a2944a7860001b61111b565b615a8e7fc4ca6fa365a9bc56f25e49edcfc530b7043326886c1ec9a698ef80c56d78800c60001b61111b565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166354e0d3c288602001516040518263ffffffff1660e01b8152600401615aef919061b56b565b60a06040518083038186803b158015615b0757600080fd5b505afa158015615b1b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615b3f919061a5e2565b9050615b6d7f6694b0b46ac55f0a5d15a22edece2ec6c8b8c4e9571ad39c533a0501ce4afb7960001b61111b565b615b997fd6b504cf9222b30917f291e40ff6928775dfbe9b8dfaa80ad8cfb22274aa0b9260001b61111b565b8660a001518760800151615bad919061c039565b67ffffffffffffffff16816000015167ffffffffffffffff1610615db757615bf77f1ac6e850550673eca718ef6df26e236decdb785115e2785674994b9c4bf4d9af60001b61111b565b615c237f97aaefd4418fa2e18bab941e2da1cd4ef5054b41de95b84dea0ace5b130b2c7e60001b61111b565b615c4f7fb01e39298031752523b06afb2bb7b7bee0e61ae9fe4106730058be13882a2cbb60001b61111b565b8381604001818151615c61919061bfca565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050615caa7f7f1002e2eb6cbf696b656918a26eeff91bf739207d58f6413f1955227e9094f360001b61111b565b615cd67f91337974e83ebf1bbaa27ea17a9e03176ff5bd32e9f70fd2decef3ee99f60cae60001b61111b565b8660a001518760800151615cea919061c039565b81602001818151615cfb919061bfca565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050615d447f21a59de1fc2b30c0248a37143c4791067e359d35a26304355009f115d96f76e860001b61111b565b615d707f29c0b13b1f4d9e2132d23eb44d2d496bc79acbc1f3aa2af2b1eb6664ad82db7160001b61111b565b8660a001518760800151615d84919061c039565b81600001818151615d95919061c0af565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050615e83565b615de37ff69fc921af84f96419038936dc95d5ad67bad79724b46a0c20d344a50d965e0560001b61111b565b615e0f7fb613733ffd58221dca93b634dd2ed2d48a7a4d7c6f1c84a9ca4549b9e61ccf1d60001b61111b565b615e3b7fdf036df2a135f0be9d8e9ff5934a3201a75f555f7e1d97663fc9fefb03f8805360001b61111b565b6040518060400160405280601581526020017f7573657220737061636520697320696e76616c696400000000000000000000008152509b5050505050505050505050506164da565b615eaf7faa3e2cf6743af2b92e9609fa5736b5deba59889b15e31c84650f2c35003b030360001b61111b565b615edb7fab92a9e23e6b9d12cecf79253cb23e60913de5754f5f50924684a0364d0e33fd60001b61111b565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ed108de98860200151836040518363ffffffff1660e01b8152600401615f3c92919061b5e6565b600060405180830381600087803b158015615f5657600080fd5b505af1158015615f6a573d6000803e3d6000fd5b5050505050615fa1565b615fa07f79aa943ebe74edbd09b5c89a58f93ee2c123d7ff41594116603158af7207b6b360001b61111b565b5b615fcd7fc3393ecfb68a0df2aee8db666047f4c1ac387f20dfc44e760ef9fe3ed459d19760001b61111b565b615ff97f7847e0dbaa9c513805f3e0248994695f00a0d0001d40837d10ff8368f06f692f60001b61111b565b600180811115616032577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b866101e001516001811115616070577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b141561610d576160a27fdeb71e6653bbde65727b7010aacfc9dba425d7ba2ab991514348a52ff1a8d6c660001b61111b565b6160ce7f8f4e42704c8d0c7fba9fc6f842d63382c8b177bba395339662635d852d3d71ad60001b61111b565b6160fa7f74e6f0ce4a5a6325003da6e8b56a9f3334697fd80d73f789e459449b2976e2bb60001b61111b565b828a616106919061bfca565b995061613a565b6161397f6b3e8087d561a744c82ec7d9f46f034dbf6bd71cf1a5b021383ecedb3f28849c60001b61111b565b5b6161667f90d9290543e00cd851afcb8443fc3a81bf8a2b24bfda2bc56fe96800cb6cb9ec60001b61111b565b6161927f4f80e65f3a411092973a5a633793802a683ad7ced62ecf662f34e96014b84ad660001b61111b565b61619e86600180613747565b5050505050505b80806161b09061c303565b9150506146cf565b506161e57fcd722c7a5770d99774d9fdb392083a3a9e55faa06b85cbcf4d3cbcf96931321560001b61111b565b6162117f7e81b2235948fd0ca38d0b7dd980201974f6cb28319941020e1e02901731ae7c60001b61111b565b60008367ffffffffffffffff16111561643f576162507f85db5023508c3ab0badd7c0447e79723e7216a7603656945a572ed7259c0ebfd60001b61111b565b61627c7f21289aecf15b25dbc8919944b3ccd6cdcc813adcc9b04151d301d924f9d8cb0f60001b61111b565b6162a87f8ca69b8a1cc321de42a1f248eedd23aa6bda6687ab76beea2d89f02327e7a1ff60001b61111b565b60008273ffffffffffffffffffffffffffffffffffffffff166108fc8567ffffffffffffffff169081150290604051600060405180830381858888f1935050505090506163177ff88d20408387ca0a98f0b478dfad80d7ed2596d6ca1f46bbb9e956304734673c60001b61111b565b6163437ff90a52eee0cd483fbaa315d346b42f1089bf1751f567c19a28e84e7377e1534860001b61111b565b8061640d576163747fe0b8ea0daebc00a17ef4dc33f64ae130e719542e175818244756afc009706a1860001b61111b565b6163a07fa0e58f77ea6ccd03f5dbf336d246fae9c605941e525630bf02f699ad48e2349160001b61111b565b6163cc7f2f0562be154e86af85100fd6917c7f1c51adb9415227564511868b8d78b5d23860001b61111b565b6040518060400160405280600d81526020017f726566756e64206661696c6564000000000000000000000000000000000000008152509450505050506164da565b6164397faf32846c8117b21c027b2dd23417bff8981ccfe8b072e47200db9b208c74cd1460001b61111b565b5061646c565b61646b7f686a1d66d1df37bc345d3101ff20dbcc3b627e8419e12b99b8a9814baec1516d60001b61111b565b5b6164987f4bb89a1b75f115450cac8bdaaa92b749acd486f9fce7cc65eb025f3b5ba46f8260001b61111b565b6164c47f5e3b2591a670c2a17b00c01277bb41a3c77da233e70a3d4a291b13c1b3486d9960001b61111b565b6040518060200160405280600081525093505050505b919050565b606061650d7f9556a4b1f9f1470c16e97062b1378b7ba86e8eee0e2032af9a89abd569a25c4760001b61111b565b6165397f8194132932aa4428e777a88674380942eb0efb8cc0655f5ba5761a13e1a30f7b60001b61111b565b6165657f72e2d2a71702576faf3d3cc94c64720fcfc87e10872adf901ea0daa4eb2e01bb60001b61111b565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c6e8392a836040518263ffffffff1660e01b81526004016165c0919061b60f565b60006040518083038186803b1580156165d857600080fd5b505afa1580156165ec573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906166159190619f85565b9050919050565b6166246181cc565b6166507f9d31c2c007ac09bb6be771a8f751f265228599bbd6e52ec65866da87deaf682f60001b61111b565b61667c7f0c13aaceba1a35fd07a4343b502c4ed172098daa9f41ed406d94fe856f06f1fd60001b61111b565b6166a87fa68a50d139cbd4c8363eb188df1edb253170e925fdb9ba01567318f997eed24e60001b61111b565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cc76e80d87878787876040518663ffffffff1660e01b815260040161670b95949392919061bbc8565b60606040518083038186803b15801561672357600080fd5b505afa158015616737573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061675b919061a50f565b905095945050505050565b6167927f648801a1b4f990e963961bdef6d614f776c714f2b65913acd980901ff9837e7060001b61111b565b6167be7f2817071a33a945e7e04355ed52153fef3766faf6154cfe3e9db7fd7501ca945560001b61111b565b6167ea7faae3224047ee3181e7865918259d93e6f3827fc896f3a69e9aeb78e3e6fe8fe560001b61111b565b60006167f58261701a565b90506168237f09634cb394ecb80d663a9cf06a9ceee22c1fb98d9c0991e00a5543ee3aa9b81260001b61111b565b61684f7fb57cb61c9551bd38d510a3a8e3cbfeec946d64b26924a111c00bdff1a9bbd91560001b61111b565b6000600167ffffffffffffffff811115616892577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156168cb57816020015b6168b861820b565b8152602001906001900390816168b05790505b5090506168fa7f13eb7926457ffd2b311b94cf0cb3b9092bd9a2d0a44f6ccde417167e5e0ccd4860001b61111b565b6169267fba489c821a41bdf815582d893ee7cb3c87d863c989b10345943f0001b2ab646460001b61111b565b8181600081518110616961577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052506169987f6df796efd035d63cfa8dae7bc125ae668f582c5e9337d8f0a30e37da3eaf4f4d60001b61111b565b6169c47f20764f59df3e073c1594e62fb444767f74f76f3b9d782354a87c82c50b893b6960001b61111b565b60006169cf826142ff565b90506169fd7f598fada4affeb8d80107622246c2f4abe3eecc19ec6563f4f018b06ec331cb7160001b61111b565b616a297fbb65845065d087a58e68a52eec17067509f7b2ea7c2dccd2b983599583851cc160001b61111b565b600081511115616b4e57616a5f7f91a9cde8380b701262370e2a80978d399e8c58b728aeaeb809c52c3197964fdb60001b61111b565b616a8b7fc4b53b20363c2ba004eac6cce26cbe368ee79596b0ae5403c64fdd5c08dbfc9e60001b61111b565b616ab77f65c48709489a780bc84379c1b01485b5b807ef1ea76e5095f148f3d3a6f3436660001b61111b565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f581604051616ae6919061b883565b60405180910390a1616b1a7f3266956389b4d7929e5767dd7d9177594b4e4755df1322154203cd6d314194ca60001b61111b565b616b467f2cc0b0027deeabff21b746d5ec44bad349d659178454a6de94eaa3fdcb40663960001b61111b565b505050616c18565b616b7a7fbd05c97dceff5a0d46746b5622314e05d05747233299db139a93367f23f254b060001b61111b565b616ba67f192d63b13641c4563e1f817de5580a7095924f8d521166c2548e8101f458651660001b61111b565b616bd27f751f889209730d631f0955d3cadf0811056ba8ae980e328f1c8006fa20a0039060001b61111b565b7fb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f45600143868660200151604051616c0c949392919061b79f565b60405180910390a15050505b50565b616c236181cc565b616c4f7f9d3c353e3e67e92266134e0278afdc617ae2eff66ebd9d3b0916fe493855f1c760001b61111b565b616c7b7ff3457cdd220d4fb20eea3b4106e03101b1c1b4f52c590fb34626b5c59cde859560001b61111b565b616ca77f55530debb388c499f88c88783c0c77f91342213794a2fafaeef8b078f374e2d660001b61111b565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d49ce8748585856040518463ffffffff1660e01b8152600401616d069392919061bc3b565b60606040518083038186803b158015616d1e57600080fd5b505afa158015616d32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616d56919061a50f565b90509392505050565b616d8b7f7e82b8dcbc8981f44236e152a0ff3bb497613baf948fb50641245af58628608060001b61111b565b616db77f0eef6aaff3b8b21754a8cf9dc8cfbf5971771c06f387f65dce7d4ddadf9b952860001b61111b565b616de37f796ef51cfc0e53b7f93d6ef1fb7b8b3f17cf0687ea2182bcaec029825eb367ea60001b61111b565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d70e627283836040518363ffffffff1660e01b8152600401616e4092919061b586565b600060405180830381600087803b158015616e5a57600080fd5b505af1158015616e6e573d6000803e3d6000fd5b505050505050565b6060616ea47f99caf209c2a02718805907bc51a7598fb8cf9990efb99790633ff6d79aeb046560001b61111b565b616ed07fc9d2fb222a43c11e6b444794288e523d52faa1758e137f0a047e3df20e9e7aa260001b61111b565b616efc7fcd7fc9e6ca48940bacf79e21db92ffcd7966117af550b35bca3e78344076ad0a60001b61111b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630c18d9cc846040518263ffffffff1660e01b8152600401616f59919061b56b565b60006040518083038186803b158015616f7157600080fd5b505afa158015616f85573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190616fae9190619e1b565b9050616fdc7f97c179f3c19aaaa1a5dbd2a0522a5579034f89cbdbaface5b0f631d0da943c2d60001b61111b565b6170087fa6f7263a4d6a50cbaa8512a7ae675d531bb8e684aef796ffcf1bd79ff39fcdd660001b61111b565b6170128184617995565b915050919050565b61702261820b565b61704e7f48f58a35ef1956e876e689c6f6954290547fda63f9ff5f6dbe3a1e796e11912e60001b61111b565b61707a7f5b1aa24d93a53080fc0939561ca6d1a193daf26304f0e9c6d236b71cbda2cf6c60001b61111b565b6170a67f9f14d8fc0ae9c25ec0b95f7a16531fb45ccf49161a287c47f367d626936705d560001b61111b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663defce5d4846040518263ffffffff1660e01b8152600401617103919061b701565b60006040518083038186803b15801561711b57600080fd5b505afa15801561712f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190617158919061a27c565b90506171867f6b4112f3855c37b88278de50da68f524b862e801d91a5fbde27e5e725d4e184360001b61111b565b6171b27f9b840a9b603a21011a5f774f51792d9fdbe687fb7504f9d42df2d08df32bf16a60001b61111b565b80915050919050565b6171e77fb936844d254aa5dc16d14d652fb69c16e52abfd9bd6989a002faaf1ed5b60e3660001b61111b565b6172137f9d72b5f492367f56809ea030d9adbad196e149f702fcb0d905ac0e9c5b06835e60001b61111b565b61723f7f3343b1203c7bc72f8315ea20611461483c4c32851297f950f9c94fc529714d3b60001b61111b565b600061724e826000015161701a565b905061727c7f620015d9dd6813fb0304c21b76064a7118c0a018f4cf2bd54c97664c856f156560001b61111b565b6172a87f260c0a02711454783ead9899e43134267a5171f73d2116d725df7848b18bee6d60001b61111b565b6000816101a0015111156173cd576172e27feb2a43b8843782daef737a57cdbf992b04f7b206310423ebe0c027c537f7fc2d60001b61111b565b61730e7f99b12319bdbf442e8ba467f4c3cfd67034826778853e906e2d10949940095c6860001b61111b565b61733a7fb645f827a48e86d8ce58ee34a6f499bc191718c1d66db7be2c2bdc6ac26c428d60001b61111b565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f56040516173679061b920565b60405180910390a161739b7fa2c5e2884a67514bb1ed5534efa2e979412cd315f140d39439cbf9e0dd8e179d60001b61111b565b6173c77f2c343682864ad39a079e93a4a70046362b1d33a7beb26c7cfe906bc4b9ea7db260001b61111b565b50617992565b6173f97fb1469e734e3c6b75479c0f1fec17e4e08b0237fb0d4d403793a40fad96811a1a60001b61111b565b6174257fbfb7cce59265944dcc6d9799fc091542b1710328e4528fa50d6dcfa444c5321560001b61111b565b6174517f106ebf6c6d2d24b8f4abf4cfb2ba5017873956407483c0f9dd950e827f20d7fc60001b61111b565b4382610100015110156175755761748a7f17941a9b01305f578cd3225cb32a307f0c156e1ba0701fc42b272a564c661d2360001b61111b565b6174b67f5f71e1af2b56b8e04eb06de9fb659b47d6556f9e01c3397bfecca8a5ffd8a29060001b61111b565b6174e27f82941522694af52ebcb35e85c5c63fd5bcbb235aa97d362bf49170e1ce3bbf4260001b61111b565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f560405161750f9061b8ed565b60405180910390a16175437f7870969cc652385fd9d0c43b43a4bd462acfd4fb00af0beb4327c77abd59b92f60001b61111b565b61756f7fe78b80af86de95ea086138790fe619d5fc804bc54bb1927fb74fe534c7f32f4360001b61111b565b50617992565b6175a17f08268161b2ee89bc4c19ec04403925739112e49ccac00cc195e75024af3c49ce60001b61111b565b6175cd7f705c0072b4c3eb47c25851f3a5e136df3508f00569913e586a8a2b99634eb30d60001b61111b565b6175f97ff9d4be24bd69ee9ad1ac47c53bb9e9b1f73f5a351e66c12bbcd05693dd87fdf360001b61111b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166396f8b9323485600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900467ffffffffffffffff166040518763ffffffff1660e01b81526004016176db95949392919061ba1f565b6000604051808303818588803b1580156176f457600080fd5b505af1158015617708573d6000803e3d6000fd5b50505050506040513d6000823e3d601f19601f82011682018060405250810190617732919061a1fa565b90506177607fd528286c67ad18d7379bdfce8950519f7f0aaa23550891e112e4fdeb711d9f5b60001b61111b565b61778c7f79408af21bd3389774b4d38742dbbf69b9a6ac453e48b1d7735661b930f7474160001b61111b565b6000815111156178b0576177c27f8d21a9f66b8f6d04c1aa12c58b72469bb6cd08dc6e6274535e37793f67e52a2260001b61111b565b6177ee7fad6d20d32c2a148cc29c58f8ad383a76a1a64ef8d9a9d0a8b29b1bf62df13bfe60001b61111b565b61781a7fae256053a90f163f71db2f52f6e2e07c5663578af5d4644419cd28eef5c4abe860001b61111b565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f581604051617849919061b8b8565b60405180910390a161787d7fa0aada19aa37980aa02fddae17af0668f4ab33169def611b43829d0b0d55f33560001b61111b565b6178a97f77114236bb7d5aa9d804de678f7d4edd6721401d8810e4452cf328a52a982ecc60001b61111b565b5050617992565b6178dc7f3fbcff6776c43dd29034506bf4be40817a942d008a7b0915df43cbed7ee52c1c60001b61111b565b6179087f5dcfcdfafe47398ca1f91797a118673ce54ac97e9db62bcb6aeb2817605f375560001b61111b565b6179347fa2e356a24180bd84250e5b332783d5fb45e3c0e944cacb9a5915accbd303d18a60001b61111b565b7fa08fe92f4f83b40431d08f9f130fd22b658ad78cf83027611d629d83427f0d1860004385600001518661020001518760200151886101400151896102a00151604051617987979695949392919061b7eb565b60405180910390a150505b50565b60606179c37f4ef5f0786e6fedd47a9abc4aee786743b6f683e8aee9ea10addf1521023f94f560001b61111b565b6179ef7f763f69cc0a37478d05a5f2fcdfeb9c2d3f596065f6514f502cade4e0719435ec60001b61111b565b617a1b7f4a612f078ac03f70c8f33bc7c1654fc9129204910a61be66c0bc01dac343d27660001b61111b565b6000835167ffffffffffffffff811115617a5e577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015617a9157816020015b6060815260200190600190039081617a7c5790505b509050617ac07f69b895d2f3d9ccb094f2af9343015fa00171d55b893d1737a0775a919c17edb160001b61111b565b617aec7fd288a71c73c0d60618ced37672f45371d286d8ad552a990c0e3d8f6964c7deac60001b61111b565b6000617b1a7fa021e4ed30afb55f27b0f56d85b6da0928817de009ba5c5ef388fc8f18a9f8c460001b61111b565b617b467f0203e434e2cf5e711ad31bff08c64863a222a17e8c79fac1496f1b0754a4b00960001b61111b565b60005b855181101561814457617b7e7f989e0319dd913c90000c8fac9997f4d2e99f7b5a9903499f9364aaa5b64fd43d60001b61111b565b617baa7fee1735b2c6286d8974659928b2b7100eebf911bf43dd5ed75c0f5efa465585b860001b61111b565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663977fdfe2888481518110617c23577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040518263ffffffff1660e01b8152600401617c47919061b701565b60006040518083038186803b158015617c5f57600080fd5b505afa158015617c73573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190617c9c919061a01a565b9050617cca7f617c79df8de33fc585e1bd523825d172f1acf0450d2661cd491d3ec4191f42b160001b61111b565b617cf67f0d42ec05be1e33e66912fccb017e71c781b78ccc8c41695f819f5dc9722aec3760001b61111b565b6000617d247f73329e4968ef0b07d97166dbc7f045c8acdc34bdf42e78c42f060f81c180036c60001b61111b565b617d507fc7339e5a170ea18dcbae190d4444df4bc68fad5378f35ecbc4cb411eedf6845960001b61111b565b60005b8251811015617f2457617d887fc353404ca4c19ba27032e793451678a7dadb3b5dd47105896c242be27a9689d360001b61111b565b617db47f5e807eb001ab7c2410a44d8041ccf6450d7e15e14c455b787048df8e052288c060001b61111b565b8773ffffffffffffffffffffffffffffffffffffffff16838281518110617e04577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff161415617ee557617e587f0bb2a98990eb4cc16f1c653705be552fae84d599b3cca06586f75aad781596fe60001b61111b565b617e847f08d44427e8a7364952dcf9679b13cc64cb12daf4fa4af00fa21d5b4af63bf21660001b61111b565b617eb07f07ee6e177e13c9916b0ff2500ce76fb5735592ae0d91251546be2063a162f78660001b61111b565b60019150617ee07fa5ab6cbf5e0b463885343c87750ff909dd5439144067a81602d475f235b7d75460001b61111b565b617f24565b617f117fe73471a222768c59aae4a209809650a6f666042dfd908b369fe19d060681670d60001b61111b565b8080617f1c9061c303565b915050617d53565b50617f517ff30e8722efeacccf0c9e0e7a76b75ab46f07b4c9ea3cb0e4657f3dceb82adcbe60001b61111b565b617f7d7fef7730f5644b4b2059078830ebe0d9dab1ae0f55f062519aa15327031da6503560001b61111b565b80617fe157617fae7f86baeabba06f3ea6dbf55b19c9a75f8e211575275d5d358a1960819887bc04b060001b61111b565b617fda7f1ed635e7f01f7ec0a7c03811b349c099f0e810e284e32b5484b84230372f25d060001b61111b565b5050618131565b61800d7f295d14135f9c21510f12767373258055fd0eb7cb26642414e500b20c8d24096660001b61111b565b6180397f704070d36d5fa0b1be06a3c021baa841e1ff59f121ede1a92c6cb3c8e4b92f6560001b61111b565b6180657fb4b4070104497464f85767642a8e35e10f0c438432af83ff73cbe376f1c88aff60001b61111b565b87838151811061809e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151858567ffffffffffffffff16815181106180e9577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052506181207fcd8af20bc23a4957de37bc438bd1aa72ab55852ed55554f06249c92b59b6f53460001b61111b565b838061812b9061c34c565b94505050505b808061813c9061c303565b915050617b49565b506181717f36ee888a4ebba1ad329110f820cd7dda15b9d8c0b704acf7534510f806088ecf60001b61111b565b61819d7fa60755a7321739a9669c263ef20c0eeda93e810feb1c852d491199e6b06452d760001b61111b565b819250505092915050565b60006181b3306181b9565b15905090565b600080823b905060008111915050919050565b6040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b604051806102e0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff1681526020016000815260200160001515815260200160006001811115618322577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600067ffffffffffffffff16815260200160608152602001606081526020016060815260200160006002811115618387577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200160001515815260200161839d6183a3565b81525090565b6040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b60006183f56183f08461bca0565b61bc7b565b9050808382526020820190508285602086028201111561841457600080fd5b60005b85811015618444578161842a8882618956565b845260208401935060208301925050600181019050618417565b5050509392505050565b600061846161845c8461bca0565b61bc7b565b9050808382526020820190508285602086028201111561848057600080fd5b60005b858110156184b05781618496888261896b565b845260208401935060208301925050600181019050618483565b5050509392505050565b60006184cd6184c88461bccc565b61bc7b565b905080838252602082019050828560208602820111156184ec57600080fd5b60005b8581101561853657813567ffffffffffffffff81111561850e57600080fd5b80860161851b8982618b63565b855260208501945060208401935050506001810190506184ef565b5050509392505050565b600061855361854e8461bccc565b61bc7b565b9050808382526020820190508285602086028201111561857257600080fd5b60005b858110156185bc57815167ffffffffffffffff81111561859457600080fd5b8086016185a18982618b8d565b85526020850194506020840193505050600181019050618575565b5050509392505050565b60006185d96185d48461bcf8565b61bc7b565b905080838252602082019050828560208602820111156185f857600080fd5b60005b85811015618628578161860e8882618c5f565b8452602084019350602083019250506001810190506185fb565b5050509392505050565b60006186456186408461bd24565b61bc7b565b9050808382526020820190508285602086028201111561866457600080fd5b60005b858110156186ae57813567ffffffffffffffff81111561868657600080fd5b8086016186938982618d13565b85526020850194506020840193505050600181019050618667565b5050509392505050565b60006186cb6186c68461bd24565b61bc7b565b905080838252602082019050828560208602820111156186ea57600080fd5b60005b8581101561873457815167ffffffffffffffff81111561870c57600080fd5b8086016187198982618fb3565b855260208501945060208401935050506001810190506186ed565b5050509392505050565b600061875161874c8461bd50565b61bc7b565b9050808382526020820190508285602086028201111561877057600080fd5b60005b858110156187ba57815167ffffffffffffffff81111561879257600080fd5b80860161879f898261952f565b85526020850194506020840193505050600181019050618773565b5050509392505050565b60006187d76187d28461bd7c565b61bc7b565b905080838252602082019050828560408602820111156187f657600080fd5b60005b85811015618826578161880c8882619751565b8452602084019350604083019250506001810190506187f9565b5050509392505050565b600061884361883e8461bda8565b61bc7b565b9050808382526020820190508285606086028201111561886257600080fd5b60005b8581101561889257816188788882619c55565b845260208401935060608301925050600181019050618865565b5050509392505050565b60006188af6188aa8461bdd4565b61bc7b565b9050828152602081018484840111156188c757600080fd5b6188d284828561c290565b509392505050565b60006188ed6188e88461bdd4565b61bc7b565b90508281526020810184848401111561890557600080fd5b61891084828561c29f565b509392505050565b600061892b6189268461be05565b61bc7b565b90508281526020810184848401111561894357600080fd5b61894e84828561c29f565b509392505050565b6000813590506189658161c61a565b92915050565b60008151905061897a8161c61a565b92915050565b600082601f83011261899157600080fd5b81356189a18482602086016183e2565b91505092915050565b600082601f8301126189bb57600080fd5b81516189cb84826020860161844e565b91505092915050565b600082601f8301126189e557600080fd5b81356189f58482602086016184ba565b91505092915050565b600082601f830112618a0f57600080fd5b8151618a1f848260208601618540565b91505092915050565b600082601f830112618a3957600080fd5b8135618a498482602086016185c6565b91505092915050565b600082601f830112618a6357600080fd5b8135618a73848260208601618632565b91505092915050565b600082601f830112618a8d57600080fd5b8151618a9d8482602086016186b8565b91505092915050565b600082601f830112618ab757600080fd5b8151618ac784826020860161873e565b91505092915050565b600082601f830112618ae157600080fd5b8151618af18482602086016187c4565b91505092915050565b600082601f830112618b0b57600080fd5b8135618b1b848260208601618830565b91505092915050565b600081359050618b338161c631565b92915050565b600081519050618b488161c631565b92915050565b600081359050618b5d8161c648565b92915050565b600082601f830112618b7457600080fd5b8135618b8484826020860161889c565b91505092915050565b600082601f830112618b9e57600080fd5b8151618bae8482602086016188da565b91505092915050565b600081359050618bc68161c65f565b92915050565b600081359050618bdb8161c676565b92915050565b600081359050618bf08161c68d565b92915050565b600081359050618c058161c6a4565b92915050565b600081359050618c1a8161c6bb565b92915050565b600081359050618c2f8161c6d2565b92915050565b600081359050618c448161c6e9565b92915050565b600081519050618c598161c6e9565b92915050565b600081359050618c6e8161c6f9565b92915050565b600081519050618c838161c6f9565b92915050565b600082601f830112618c9a57600080fd5b8151618caa848260208601618918565b91505092915050565b600060608284031215618cc557600080fd5b618ccf606061bc7b565b90506000618cdf84828501619cdf565b6000830152506020618cf384828501619cdf565b6020830152506040618d0784828501619cdf565b60408301525092915050565b60006103208284031215618d2657600080fd5b618d316102e061bc7b565b9050600082013567ffffffffffffffff811115618d4d57600080fd5b618d5984828501618b63565b6000830152506020618d6d84828501618956565b602083015250604082013567ffffffffffffffff811115618d8d57600080fd5b618d9984828501618b63565b6040830152506060618dad84828501619cdf565b6060830152506080618dc184828501619cdf565b60808301525060a0618dd584828501619cdf565b60a08301525060c0618de984828501619cdf565b60c08301525060e0618dfd84828501619cdf565b60e083015250610100618e1284828501619cb5565b61010083015250610120618e2884828501619cdf565b61012083015250610140618e3e84828501619cdf565b6101408301525061016082013567ffffffffffffffff811115618e6057600080fd5b618e6c84828501618b63565b61016083015250610180618e8284828501619cdf565b610180830152506101a0618e9884828501619cb5565b6101a0830152506101c0618eae84828501618b24565b6101c0830152506101e0618ec484828501618c5f565b6101e083015250610200618eda84828501619cdf565b6102008301525061022082013567ffffffffffffffff811115618efc57600080fd5b618f0884828501618980565b6102208301525061024082013567ffffffffffffffff811115618f2a57600080fd5b618f3684828501618980565b6102408301525061026082013567ffffffffffffffff811115618f5857600080fd5b618f6484828501618b63565b61026083015250610280618f7a84828501618c35565b610280830152506102a0618f9084828501618b24565b6102a0830152506102c0618fa68482850161940b565b6102c08301525092915050565b60006103208284031215618fc657600080fd5b618fd16102e061bc7b565b9050600082015167ffffffffffffffff811115618fed57600080fd5b618ff984828501618b8d565b600083015250602061900d8482850161896b565b602083015250604082015167ffffffffffffffff81111561902d57600080fd5b61903984828501618b8d565b604083015250606061904d84828501619cf4565b606083015250608061906184828501619cf4565b60808301525060a061907584828501619cf4565b60a08301525060c061908984828501619cf4565b60c08301525060e061909d84828501619cf4565b60e0830152506101006190b284828501619cca565b610100830152506101206190c884828501619cf4565b610120830152506101406190de84828501619cf4565b6101408301525061016082015167ffffffffffffffff81111561910057600080fd5b61910c84828501618b8d565b6101608301525061018061912284828501619cf4565b610180830152506101a061913884828501619cca565b6101a0830152506101c061914e84828501618b39565b6101c0830152506101e061916484828501618c74565b6101e08301525061020061917a84828501619cf4565b6102008301525061022082015167ffffffffffffffff81111561919c57600080fd5b6191a8848285016189aa565b6102208301525061024082015167ffffffffffffffff8111156191ca57600080fd5b6191d6848285016189aa565b6102408301525061026082015167ffffffffffffffff8111156191f857600080fd5b61920484828501618b8d565b6102608301525061028061921a84828501618c4a565b610280830152506102a061923084828501618b39565b6102a0830152506102c06192468482850161946b565b6102c08301525092915050565b60006060828403121561926557600080fd5b61926f606061bc7b565b9050600082013567ffffffffffffffff81111561928b57600080fd5b61929784828501618b63565b60008301525060206192ab84828501618956565b60208301525060406192bf84828501619cdf565b60408301525092915050565b600060e082840312156192dd57600080fd5b6192e760e061bc7b565b905060006192f784828501619cf4565b600083015250602061930b84828501619cf4565b602083015250604061931f84828501619cf4565b604083015250606061933384828501619cf4565b606083015250608061934784828501619cf4565b60808301525060a061935b8482850161896b565b60a08301525060c082015167ffffffffffffffff81111561937b57600080fd5b61938784828501618b8d565b60c08301525092915050565b6000606082840312156193a557600080fd5b6193af606061bc7b565b9050600082013567ffffffffffffffff8111156193cb57600080fd5b6193d784828501618b63565b60008301525060206193eb84828501618956565b60208301525060406193ff84828501618956565b60408301525092915050565b60006060828403121561941d57600080fd5b619427606061bc7b565b9050600061943784828501619cdf565b600083015250602061944b84828501619cdf565b602083015250604061945f84828501619cdf565b60408301525092915050565b60006060828403121561947d57600080fd5b619487606061bc7b565b9050600061949784828501619cf4565b60008301525060206194ab84828501619cf4565b60208301525060406194bf84828501619cf4565b60408301525092915050565b6000604082840312156194dd57600080fd5b6194e7604061bc7b565b9050600082013567ffffffffffffffff81111561950357600080fd5b61950f84828501618b63565b600083015250602061952384828501619cdf565b60208301525092915050565b600060a0828403121561954157600080fd5b61954b60a061bc7b565b9050600082015167ffffffffffffffff81111561956757600080fd5b61957384828501618b8d565b60008301525060206195878482850161896b565b602083015250604061959b84828501619cf4565b60408301525060606195af84828501619cca565b60608301525060806195c384828501618b39565b60808301525092915050565b600061018082840312156195e257600080fd5b6195ed61018061bc7b565b905060006195fd8482850161896b565b600083015250602061961184828501619cf4565b602083015250604061962584828501619cf4565b604083015250606061963984828501619cf4565b606083015250608061964d84828501618c4a565b60808301525060a061966184828501619cca565b60a08301525060c061967584828501619cca565b60c08301525060e061968984828501619cf4565b60e08301525061010061969e84828501619cf4565b610100830152506101206196b484828501619cf4565b610120830152506101406196ca84828501618b39565b6101408301525061016082015167ffffffffffffffff8111156196ec57600080fd5b6196f8848285016189fe565b6101608301525092915050565b60006040828403121561971757600080fd5b619721604061bc7b565b9050600061973184828501618956565b600083015250602061974584828501619cdf565b60208301525092915050565b60006040828403121561976357600080fd5b61976d604061bc7b565b9050600061977d8482850161896b565b600083015250602061979184828501619cf4565b60208301525092915050565b600061016082840312156197b057600080fd5b6197bb61016061bc7b565b905060006197cb84828501619cdf565b60008301525060206197df84828501619cdf565b60208301525060406197f384828501619cdf565b604083015250606061980784828501619cdf565b606083015250608061981b84828501619cdf565b60808301525060a061982f84828501619cdf565b60a08301525060c061984384828501619cdf565b60c08301525060e061985784828501619cdf565b60e08301525061010061986c84828501619cdf565b6101008301525061012061988284828501619cdf565b6101208301525061014061989884828501619cdf565b6101408301525092915050565b600061016082840312156198b857600080fd5b6198c361016061bc7b565b905060006198d384828501619cf4565b60008301525060206198e784828501619cf4565b60208301525060406198fb84828501619cf4565b604083015250606061990f84828501619cf4565b606083015250608061992384828501619cf4565b60808301525060a061993784828501619cf4565b60a08301525060c061994b84828501619cf4565b60c08301525060e061995f84828501619cf4565b60e08301525061010061997484828501619cf4565b6101008301525061012061998a84828501619cf4565b610120830152506101406199a084828501619cf4565b6101408301525092915050565b6000606082840312156199bf57600080fd5b6199c9606061bc7b565b905060006199d984828501619cf4565b60008301525060206199ed84828501619cf4565b6020830152506040619a0184828501619cf4565b60408301525092915050565b60006101e08284031215619a2057600080fd5b619a2b6101e061bc7b565b9050600082013567ffffffffffffffff811115619a4757600080fd5b619a5384828501618b63565b6000830152506020619a6784828501619cdf565b6020830152506040619a7b84828501619cdf565b6040830152506060619a8f84828501619cdf565b6060830152506080619aa384828501619cb5565b60808301525060a0619ab784828501619cdf565b60a08301525060c0619acb84828501619cdf565b60c08301525060e0619adf84828501618b24565b60e08301525061010082013567ffffffffffffffff811115619b0057600080fd5b619b0c84828501618b63565b61010083015250610120619b2284828501618b24565b61012083015250610140619b3884828501618b24565b6101408301525061016082013567ffffffffffffffff811115619b5a57600080fd5b619b6684828501618b63565b6101608301525061018082013567ffffffffffffffff811115619b8857600080fd5b619b9484828501618afa565b610180830152506101a0619baa84828501618b24565b6101a0830152506101c0619bc084828501618c5f565b6101c08301525092915050565b600060a08284031215619bdf57600080fd5b619be960a061bc7b565b90506000619bf984828501619cf4565b6000830152506020619c0d84828501619cf4565b6020830152506040619c2184828501619cf4565b6040830152506060619c3584828501619cca565b6060830152506080619c4984828501619cca565b60808301525092915050565b600060608284031215619c6757600080fd5b619c71606061bc7b565b90506000619c8184828501618956565b6000830152506020619c9584828501619cdf565b6020830152506040619ca984828501619cdf565b60408301525092915050565b600081359050619cc48161c709565b92915050565b600081519050619cd98161c709565b92915050565b600081359050619cee8161c720565b92915050565b600081519050619d038161c720565b92915050565b600060208284031215619d1b57600080fd5b6000619d2984828501618956565b91505092915050565b60008060408385031215619d4557600080fd5b6000619d5385828601618956565b925050602083013567ffffffffffffffff811115619d7057600080fd5b619d7c858286016189d4565b9150509250929050565b60008060408385031215619d9957600080fd5b6000619da785828601618956565b925050602083013567ffffffffffffffff811115619dc457600080fd5b619dd085828601618b63565b9150509250929050565b600060208284031215619dec57600080fd5b600082013567ffffffffffffffff811115619e0657600080fd5b619e12848285016189d4565b91505092915050565b600060208284031215619e2d57600080fd5b600082015167ffffffffffffffff811115619e4757600080fd5b619e53848285016189fe565b91505092915050565b600080600060608486031215619e7157600080fd5b600084013567ffffffffffffffff811115619e8b57600080fd5b619e97868287016189d4565b9350506020619ea886828701618956565b925050604084013567ffffffffffffffff811115619ec557600080fd5b619ed186828701618a28565b9150509250925092565b60008060006101a08486031215619ef157600080fd5b600084013567ffffffffffffffff811115619f0b57600080fd5b619f17868287016189d4565b9350506020619f288682870161979d565b925050610180619f3a86828701619cb5565b9150509250925092565b600060208284031215619f5657600080fd5b600082013567ffffffffffffffff811115619f7057600080fd5b619f7c84828501618a52565b91505092915050565b600060208284031215619f9757600080fd5b600082015167ffffffffffffffff811115619fb157600080fd5b619fbd84828501618a7c565b91505092915050565b60008060408385031215619fd957600080fd5b600083015167ffffffffffffffff811115619ff357600080fd5b619fff85828601618a7c565b925050602061a01085828601618b39565b9150509250929050565b60006020828403121561a02c57600080fd5b600082015167ffffffffffffffff81111561a04657600080fd5b61a05284828501618aa6565b91505092915050565b60006020828403121561a06d57600080fd5b600082015167ffffffffffffffff81111561a08757600080fd5b61a09384828501618ad0565b91505092915050565b60006020828403121561a0ae57600080fd5b600061a0bc84828501618b4e565b91505092915050565b60006020828403121561a0d757600080fd5b600082013567ffffffffffffffff81111561a0f157600080fd5b61a0fd84828501618b63565b91505092915050565b6000806060838503121561a11957600080fd5b600083013567ffffffffffffffff81111561a13357600080fd5b61a13f85828601618b63565b925050602061a15085828601619705565b9150509250929050565b6000806000806000806000610120888a03121561a17657600080fd5b600061a1848a828b01618bcc565b975050602061a1958a828b01618be1565b965050604061a1a68a828b01618c20565b955050606061a1b78a828b01618c0b565b945050608061a1c88a828b01618bf6565b93505060a061a1d98a828b01618cb3565b92505061010061a1eb8a828b01618bb7565b91505092959891949750929550565b60006020828403121561a20c57600080fd5b600082015167ffffffffffffffff81111561a22657600080fd5b61a23284828501618c89565b91505092915050565b60006020828403121561a24d57600080fd5b600082013567ffffffffffffffff81111561a26757600080fd5b61a27384828501618d13565b91505092915050565b60006020828403121561a28e57600080fd5b600082015167ffffffffffffffff81111561a2a857600080fd5b61a2b484828501618fb3565b91505092915050565b60008060006060848603121561a2d257600080fd5b600084013567ffffffffffffffff81111561a2ec57600080fd5b61a2f886828701618d13565b935050602061a30986828701618b24565b925050604061a31a86828701618b24565b9150509250925092565b60006020828403121561a33657600080fd5b600082013567ffffffffffffffff81111561a35057600080fd5b61a35c84828501619253565b91505092915050565b60006020828403121561a37757600080fd5b600082015167ffffffffffffffff81111561a39157600080fd5b61a39d848285016192cb565b91505092915050565b60006020828403121561a3b857600080fd5b600082013567ffffffffffffffff81111561a3d257600080fd5b61a3de84828501619393565b91505092915050565b60006020828403121561a3f957600080fd5b600082013567ffffffffffffffff81111561a41357600080fd5b61a41f848285016194cb565b91505092915050565b60006020828403121561a43a57600080fd5b600082015167ffffffffffffffff81111561a45457600080fd5b61a460848285016195cf565b91505092915050565b6000610160828403121561a47c57600080fd5b600061a48a848285016198a5565b91505092915050565b60008060008060006101e0868803121561a4ac57600080fd5b600061a4ba8882890161979d565b95505061016061a4cc88828901619cdf565b94505061018061a4de88828901619cdf565b9350506101a061a4f088828901619cdf565b9250506101c061a50288828901619cdf565b9150509295509295909350565b60006060828403121561a52157600080fd5b600061a52f848285016199ad565b91505092915050565b60006020828403121561a54a57600080fd5b600082013567ffffffffffffffff81111561a56457600080fd5b61a57084828501619a0d565b91505092915050565b60008060006101a0848603121561a58f57600080fd5b600084013567ffffffffffffffff81111561a5a957600080fd5b61a5b586828701619a0d565b935050602061a5c68682870161979d565b92505061018061a5d886828701619cb5565b9150509250925092565b600060a0828403121561a5f457600080fd5b600061a60284828501619bcd565b91505092915050565b60006020828403121561a61d57600080fd5b600061a62b84828501619cf4565b91505092915050565b600061a640838361a68c565b60208301905092915050565b600061a658838361a8e3565b905092915050565b600061a66c838361ab00565b905092915050565b600061a680838361b4ed565b60608301905092915050565b61a6958161c0e3565b82525050565b61a6a48161c0e3565b82525050565b600061a6b58261be76565b61a6bf818561beec565b935061a6ca8361be36565b8060005b8381101561a6fb57815161a6e2888261a634565b975061a6ed8361beb8565b92505060018101905061a6ce565b5085935050505092915050565b600061a7138261be81565b61a71d818561befd565b93508360208202850161a72f8561be46565b8060005b8581101561a76b578484038952815161a74c858261a64c565b945061a7578361bec5565b925060208a0199505060018101905061a733565b50829750879550505050505092915050565b600061a7888261be81565b61a792818561bf0e565b93508360208202850161a7a48561be46565b8060005b8581101561a7e0578484038952815161a7c1858261a64c565b945061a7cc8361bec5565b925060208a0199505060018101905061a7a8565b50829750879550505050505092915050565b600061a7fd8261be8c565b61a807818561bf1f565b93508360208202850161a8198561be56565b8060005b8581101561a855578484038952815161a836858261a660565b945061a8418361bed2565b925060208a0199505060018101905061a81d565b50829750879550505050505092915050565b600061a8728261be97565b61a87c818561bf30565b935061a8878361be66565b8060005b8381101561a8b857815161a89f888261a674565b975061a8aa8361bedf565b92505060018101905061a88b565b5085935050505092915050565b61a8ce8161c0f5565b82525050565b61a8dd8161c0f5565b82525050565b600061a8ee8261bea2565b61a8f8818561bf41565b935061a90881856020860161c29f565b61a9118161c439565b840191505092915050565b600061a9278261bea2565b61a931818561bf52565b935061a94181856020860161c29f565b61a94a8161c439565b840191505092915050565b61a95e8161c1ee565b82525050565b61a96d8161c212565b82525050565b61a97c8161c236565b82525050565b61a98b8161c25a565b82525050565b61a99a8161c26c565b82525050565b61a9a98161c27e565b82525050565b600061a9ba8261bead565b61a9c4818561bf63565b935061a9d481856020860161c29f565b61a9dd8161c439565b840191505092915050565b600061a9f560318361bf63565b915061aa008261c44a565b604082019050919050565b600061aa18600a8361bf63565b915061aa238261c499565b602082019050919050565b600061aa3b60098361bf63565b915061aa468261c4c2565b602082019050919050565b600061aa5e600b8361bf63565b915061aa698261c4eb565b602082019050919050565b600061aa81602e8361bf63565b915061aa8c8261c514565b604082019050919050565b600061aaa4601f8361bf63565b915061aaaf8261c563565b602082019050919050565b600061aac760098361bf63565b915061aad28261c58c565b602082019050919050565b600061aaea60128361bf63565b915061aaf58261c5b5565b602082019050919050565b600061032083016000830151848203600086015261ab1e828261a8e3565b915050602083015161ab33602086018261a68c565b506040830151848203604086015261ab4b828261a8e3565b915050606083015161ab60606086018261b54d565b50608083015161ab73608086018261b54d565b5060a083015161ab8660a086018261b54d565b5060c083015161ab9960c086018261b54d565b5060e083015161abac60e086018261b54d565b5061010083015161abc161010086018261b52f565b5061012083015161abd661012086018261b54d565b5061014083015161abeb61014086018261b54d565b5061016083015184820361016086015261ac05828261a8e3565b91505061018083015161ac1c61018086018261b54d565b506101a083015161ac316101a086018261b52f565b506101c083015161ac466101c086018261a8c5565b506101e083015161ac5b6101e086018261a9a0565b5061020083015161ac7061020086018261b54d565b5061022083015184820361022086015261ac8a828261a6aa565b91505061024083015184820361024086015261aca6828261a6aa565b91505061026083015184820361026086015261acc2828261a8e3565b91505061028083015161acd961028086018261a991565b506102a083015161acee6102a086018261a8c5565b506102c083015161ad036102c086018261b058565b508091505092915050565b600061032083016000830151848203600086015261ad2c828261a8e3565b915050602083015161ad41602086018261a68c565b506040830151848203604086015261ad59828261a8e3565b915050606083015161ad6e606086018261b54d565b50608083015161ad81608086018261b54d565b5060a083015161ad9460a086018261b54d565b5060c083015161ada760c086018261b54d565b5060e083015161adba60e086018261b54d565b5061010083015161adcf61010086018261b52f565b5061012083015161ade461012086018261b54d565b5061014083015161adf961014086018261b54d565b5061016083015184820361016086015261ae13828261a8e3565b91505061018083015161ae2a61018086018261b54d565b506101a083015161ae3f6101a086018261b52f565b506101c083015161ae546101c086018261a8c5565b506101e083015161ae696101e086018261a9a0565b5061020083015161ae7e61020086018261b54d565b5061022083015184820361022086015261ae98828261a6aa565b91505061024083015184820361024086015261aeb4828261a6aa565b91505061026083015184820361026086015261aed0828261a8e3565b91505061028083015161aee761028086018261a991565b506102a083015161aefc6102a086018261a8c5565b506102c083015161af116102c086018261b058565b508091505092915050565b6000606083016000830151848203600086015261af39828261a8e3565b915050602083015161af4e602086018261a68c565b50604083015161af61604086018261b54d565b508091505092915050565b600060e08301600083015161af84600086018261b54d565b50602083015161af97602086018261b54d565b50604083015161afaa604086018261b54d565b50606083015161afbd606086018261b54d565b50608083015161afd0608086018261b54d565b5060a083015161afe360a086018261a68c565b5060c083015184820360c086015261affb828261a8e3565b9150508091505092915050565b6000606083016000830151848203600086015261b025828261a8e3565b915050602083015161b03a602086018261a68c565b50604083015161b04d604086018261a68c565b508091505092915050565b60608201600082015161b06e600085018261b54d565b50602082015161b081602085018261b54d565b50604082015161b094604085018261b54d565b50505050565b6000604083016000830151848203600086015261b0b7828261a8e3565b915050602083015161b0cc602086018261b54d565b508091505092915050565b60006101808301600083015161b0f0600086018261a68c565b50602083015161b103602086018261b54d565b50604083015161b116604086018261b54d565b50606083015161b129606086018261b54d565b50608083015161b13c608086018261a991565b5060a083015161b14f60a086018261b52f565b5060c083015161b16260c086018261b52f565b5060e083015161b17560e086018261b54d565b5061010083015161b18a61010086018261b54d565b5061012083015161b19f61012086018261b54d565b5061014083015161b1b461014086018261a8c5565b5061016083015184820361016086015261b1ce828261a708565b9150508091505092915050565b60408201600082015161b1f1600085018261a68c565b50602082015161b204602085018261b54d565b50505050565b6101608201600082015161b221600085018261b54d565b50602082015161b234602085018261b54d565b50604082015161b247604085018261b54d565b50606082015161b25a606085018261b54d565b50608082015161b26d608085018261b54d565b5060a082015161b28060a085018261b54d565b5060c082015161b29360c085018261b54d565b5060e082015161b2a660e085018261b54d565b5061010082015161b2bb61010085018261b54d565b5061012082015161b2d061012085018261b54d565b5061014082015161b2e561014085018261b54d565b50505050565b60608201600082015161b301600085018261b54d565b50602082015161b314602085018261b54d565b50604082015161b327604085018261b54d565b50505050565b60006101e083016000830151848203600086015261b34b828261a8e3565b915050602083015161b360602086018261b54d565b50604083015161b373604086018261b54d565b50606083015161b386606086018261b54d565b50608083015161b399608086018261b52f565b5060a083015161b3ac60a086018261b54d565b5060c083015161b3bf60c086018261b54d565b5060e083015161b3d260e086018261a8c5565b5061010083015184820361010086015261b3ec828261a8e3565b91505061012083015161b40361012086018261a8c5565b5061014083015161b41861014086018261a8c5565b5061016083015184820361016086015261b432828261a8e3565b91505061018083015184820361018086015261b44e828261a867565b9150506101a083015161b4656101a086018261a8c5565b506101c083015161b47a6101c086018261a9a0565b508091505092915050565b60a08201600082015161b49b600085018261b54d565b50602082015161b4ae602085018261b54d565b50604082015161b4c1604085018261b54d565b50606082015161b4d4606085018261b52f565b50608082015161b4e7608085018261b52f565b50505050565b60608201600082015161b503600085018261a68c565b50602082015161b516602085018261b54d565b50604082015161b529604085018261b54d565b50505050565b61b5388161c1d0565b82525050565b61b5478161c1d0565b82525050565b61b5568161c1da565b82525050565b61b5658161c1da565b82525050565b600060208201905061b580600083018461a69b565b92915050565b600060408201905061b59b600083018561a69b565b818103602083015261b5ad818461a77d565b90509392505050565b600060408201905061b5cb600083018561a69b565b818103602083015261b5dd818461a91c565b90509392505050565b600060c08201905061b5fb600083018561a69b565b61b608602083018461b485565b9392505050565b6000602082019050818103600083015261b629818461a77d565b905092915050565b60006101a082019050818103600083015261b64c818661a77d565b905061b65b602083018561b20a565b61b66961018083018461b53e565b949350505050565b6000606082019050818103600083015261b68b818661a77d565b905061b69a602083018561b55c565b61b6a7604083018461a8d4565b949350505050565b6000602082019050818103600083015261b6c9818461a7f2565b905092915050565b6000604082019050818103600083015261b6eb818561a7f2565b905061b6fa602083018461a8d4565b9392505050565b6000602082019050818103600083015261b71b818461a91c565b905092915050565b6000606082019050818103600083015261b73d818561a91c565b905061b74c602083018461b1db565b9392505050565b600060808201905061b768600083018761a982565b61b775602083018661b53e565b818103604083015261b787818561a77d565b905061b796606083018461a69b565b95945050505050565b600060808201905061b7b4600083018761a982565b61b7c1602083018661b53e565b818103604083015261b7d3818561a91c565b905061b7e2606083018461a69b565b95945050505050565b600060e08201905061b800600083018a61a982565b61b80d602083018961b53e565b818103604083015261b81f818861a91c565b905061b82e606083018761b55c565b61b83b608083018661a69b565b61b84860a083018561b55c565b61b85560c083018461a8d4565b98975050505050505050565b6000602082019050818103600083015261b87b818461a9af565b905092915050565b6000604082019050818103600083015261b89c8161aa0b565b9050818103602083015261b8b0818461a9af565b905092915050565b6000604082019050818103600083015261b8d18161aa2e565b9050818103602083015261b8e5818461a9af565b905092915050565b6000604082019050818103600083015261b9068161aa2e565b9050818103602083015261b9198161a9e8565b9050919050565b6000604082019050818103600083015261b9398161aa2e565b9050818103602083015261b94c8161aadd565b9050919050565b6000604082019050818103600083015261b96c8161aa51565b9050818103602083015261b980818461a9af565b905092915050565b6000602082019050818103600083015261b9a18161aa74565b9050919050565b6000602082019050818103600083015261b9c18161aa97565b9050919050565b6000604082019050818103600083015261b9e18161aaba565b9050818103602083015261b9f5818461a9af565b905092915050565b6000602082019050818103600083015261ba17818461ad0e565b905092915050565b600060a082019050818103600083015261ba39818861ad0e565b905061ba48602083018761a955565b61ba55604083018661a973565b61ba62606083018561a964565b61ba6f608083018461b55c565b9695505050505050565b6000602082019050818103600083015261ba93818461af6c565b905092915050565b6000602082019050818103600083015261bab5818461b008565b905092915050565b6000602082019050818103600083015261bad7818461b09a565b905092915050565b6000604082019050818103600083015261baf9818561b0d7565b9050818103602083015261bb0d818461ad0e565b90509392505050565b600060408201905061bb2b600083018461b1db565b92915050565b60006101808201905061bb47600083018561b20a565b81810361016083015261bb5a818461af1c565b90509392505050565b60006101808201905061bb79600083018561b20a565b61bb8761016083018461b55c565b9392505050565b60006101a08201905061bba4600083018661b20a565b61bbb261016083018561b55c565b61bbc061018083018461b55c565b949350505050565b60006101e08201905061bbde600083018861b20a565b61bbec61016083018761b55c565b61bbfa61018083018661b55c565b61bc086101a083018561b55c565b61bc166101c083018461b55c565b9695505050505050565b600060608201905061bc35600083018461b2eb565b92915050565b60006101a082019050818103600083015261bc56818661b32d565b905061bc65602083018561b20a565b61bc7361018083018461b53e565b949350505050565b600061bc8561bc96565b905061bc91828261c2d2565b919050565b6000604051905090565b600067ffffffffffffffff82111561bcbb5761bcba61c40a565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561bce75761bce661c40a565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561bd135761bd1261c40a565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561bd3f5761bd3e61c40a565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561bd6b5761bd6a61c40a565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561bd975761bd9661c40a565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561bdc35761bdc261c40a565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561bdef5761bdee61c40a565b5b61bdf88261c439565b9050602081019050919050565b600067ffffffffffffffff82111561be205761be1f61c40a565b5b61be298261c439565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600061bf7f8261c1d0565b915061bf8a8361c1d0565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561bfbf5761bfbe61c37d565b5b828201905092915050565b600061bfd58261c1da565b915061bfe08361c1da565b92508267ffffffffffffffff0382111561bffd5761bffc61c37d565b5b828201905092915050565b600061c0138261c1da565b915061c01e8361c1da565b92508261c02e5761c02d61c3ac565b5b828204905092915050565b600061c0448261c1da565b915061c04f8361c1da565b92508167ffffffffffffffff048311821515161561c0705761c06f61c37d565b5b828202905092915050565b600061c0868261c1d0565b915061c0918361c1d0565b92508282101561c0a45761c0a361c37d565b5b828203905092915050565b600061c0ba8261c1da565b915061c0c58361c1da565b92508282101561c0d85761c0d761c37d565b5b828203905092915050565b600061c0ee8261c1b0565b9050919050565b60008115159050919050565b6000819050919050565b600061c1168261c0e3565b9050919050565b600061c1288261c0e3565b9050919050565b600061c13a8261c0e3565b9050919050565b600061c14c8261c0e3565b9050919050565b600061c15e8261c0e3565b9050919050565b600061c1708261c0e3565b9050919050565b600081905061c1858261c5de565b919050565b600081905061c1988261c5f2565b919050565b600081905061c1ab8261c606565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600061c1f98261c200565b9050919050565b600061c20b8261c1b0565b9050919050565b600061c21d8261c224565b9050919050565b600061c22f8261c1b0565b9050919050565b600061c2418261c248565b9050919050565b600061c2538261c1b0565b9050919050565b600061c2658261c177565b9050919050565b600061c2778261c18a565b9050919050565b600061c2898261c19d565b9050919050565b82818337600083830152505050565b60005b8381101561c2bd57808201518184015260208101905061c2a2565b8381111561c2cc576000848401525b50505050565b61c2db8261c439565b810181811067ffffffffffffffff8211171561c2fa5761c2f961c40a565b5b80604052505050565b600061c30e8261c1d0565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561c3415761c34061c37d565b5b600182019050919050565b600061c3578261c1da565b915067ffffffffffffffff82141561c3725761c37161c37d565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f65787069726564486569676874206d757374206265206772656174657220746860008201527f616e2063757272656e7420686569676874000000000000000000000000000000602082015250565b7f44656c65746546696c6500000000000000000000000000000000000000000000600082015250565b7f53746f726546696c650000000000000000000000000000000000000000000000600082015250565b7f44656c65746546696c6573000000000000000000000000000000000000000000600082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f66696c6553697a65206d7573742062652067726561746572207468616e203000600082015250565b7f46696c6552654e65770000000000000000000000000000000000000000000000600082015250565b7f66696c6520616c72656164792065786973740000000000000000000000000000600082015250565b600b811061c5ef5761c5ee61c3db565b5b50565b6003811061c6035761c60261c3db565b5b50565b6002811061c6175761c61661c3db565b5b50565b61c6238161c0e3565b811461c62e57600080fd5b50565b61c63a8161c0f5565b811461c64557600080fd5b50565b61c6518161c101565b811461c65c57600080fd5b50565b61c6688161c10b565b811461c67357600080fd5b50565b61c67f8161c11d565b811461c68a57600080fd5b50565b61c6968161c12f565b811461c6a157600080fd5b50565b61c6ad8161c141565b811461c6b857600080fd5b50565b61c6c48161c153565b811461c6cf57600080fd5b50565b61c6db8161c165565b811461c6e657600080fd5b50565b6003811061c6f657600080fd5b50565b6002811061c70657600080fd5b50565b61c7128161c1d0565b811461c71d57600080fd5b50565b61c7298161c1da565b811461c73457600080fd5b5056fea2646970667358221220efe7bb5a2163c0d19cdf9c0a4285f8eee0ff4d97e96e6a71230497052d382d8464736f6c63430008040033",
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

// CalcDepositFee is a free data retrieval call binding the contract method 0x178e4dc0.
//
// Solidity: function CalcDepositFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) view returns((uint64,uint64,uint64))
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
// Solidity: function CalcDepositFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) view returns((uint64,uint64,uint64))
func (_Store *StoreSession) CalcDepositFee(uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	return _Store.Contract.CalcDepositFee(&_Store.CallOpts, uploadOption, setting, currentHeight)
}

// CalcDepositFee is a free data retrieval call binding the contract method 0x178e4dc0.
//
// Solidity: function CalcDepositFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) view returns((uint64,uint64,uint64))
func (_Store *StoreCallerSession) CalcDepositFee(uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	return _Store.Contract.CalcDepositFee(&_Store.CallOpts, uploadOption, setting, currentHeight)
}

// CalcFee is a free data retrieval call binding the contract method 0xcc76e80d.
//
// Solidity: function CalcFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 copyNum, uint64 fileSize, uint64 duration) view returns((uint64,uint64,uint64))
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
// Solidity: function CalcFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 copyNum, uint64 fileSize, uint64 duration) view returns((uint64,uint64,uint64))
func (_Store *StoreSession) CalcFee(setting Setting, proveTime uint64, copyNum uint64, fileSize uint64, duration uint64) (StorageFee, error) {
	return _Store.Contract.CalcFee(&_Store.CallOpts, setting, proveTime, copyNum, fileSize, duration)
}

// CalcFee is a free data retrieval call binding the contract method 0xcc76e80d.
//
// Solidity: function CalcFee((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 proveTime, uint64 copyNum, uint64 fileSize, uint64 duration) view returns((uint64,uint64,uint64))
func (_Store *StoreCallerSession) CalcFee(setting Setting, proveTime uint64, copyNum uint64, fileSize uint64, duration uint64) (StorageFee, error) {
	return _Store.Contract.CalcFee(&_Store.CallOpts, setting, proveTime, copyNum, fileSize, duration)
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

// C0x74cde3b0 is a free data retrieval call binding the contract method 0x1ca4a479.
//
// Solidity: function c_0x74cde3b0(bytes32 c__0x74cde3b0) pure returns()
func (_Store *StoreCaller) C0x74cde3b0(opts *bind.CallOpts, c__0x74cde3b0 [32]byte) error {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "c_0x74cde3b0", c__0x74cde3b0)

	if err != nil {
		return err
	}

	return err

}

// C0x74cde3b0 is a free data retrieval call binding the contract method 0x1ca4a479.
//
// Solidity: function c_0x74cde3b0(bytes32 c__0x74cde3b0) pure returns()
func (_Store *StoreSession) C0x74cde3b0(c__0x74cde3b0 [32]byte) error {
	return _Store.Contract.C0x74cde3b0(&_Store.CallOpts, c__0x74cde3b0)
}

// C0x74cde3b0 is a free data retrieval call binding the contract method 0x1ca4a479.
//
// Solidity: function c_0x74cde3b0(bytes32 c__0x74cde3b0) pure returns()
func (_Store *StoreCallerSession) C0x74cde3b0(c__0x74cde3b0 [32]byte) error {
	return _Store.Contract.C0x74cde3b0(&_Store.CallOpts, c__0x74cde3b0)
}

// CalcUploadFee is a free data retrieval call binding the contract method 0xd49ce874.
//
// Solidity: function calcUploadFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) view returns((uint64,uint64,uint64))
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
// Solidity: function calcUploadFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) view returns((uint64,uint64,uint64))
func (_Store *StoreSession) CalcUploadFee(uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	return _Store.Contract.CalcUploadFee(&_Store.CallOpts, uploadOption, setting, currentHeight)
}

// CalcUploadFee is a free data retrieval call binding the contract method 0xd49ce874.
//
// Solidity: function calcUploadFee((bytes,uint64,uint64,uint64,uint256,uint64,uint64,bool,bytes,bool,bool,bytes,(address,uint64,uint64)[],bool,uint8) uploadOption, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 currentHeight) view returns((uint64,uint64,uint64))
func (_Store *StoreCallerSession) CalcUploadFee(uploadOption UploadOption, setting Setting, currentHeight *big.Int) (StorageFee, error) {
	return _Store.Contract.CalcUploadFee(&_Store.CallOpts, uploadOption, setting, currentHeight)
}

// AddFileSectorRef is a paid mutator transaction binding the contract method 0x5a0e7482.
//
// Solidity: function AddFileSectorRef(bytes fileHash, (address,uint64) ref) payable returns()
func (_Store *StoreTransactor) AddFileSectorRef(opts *bind.TransactOpts, fileHash []byte, ref SectorRef) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "AddFileSectorRef", fileHash, ref)
}

// AddFileSectorRef is a paid mutator transaction binding the contract method 0x5a0e7482.
//
// Solidity: function AddFileSectorRef(bytes fileHash, (address,uint64) ref) payable returns()
func (_Store *StoreSession) AddFileSectorRef(fileHash []byte, ref SectorRef) (*types.Transaction, error) {
	return _Store.Contract.AddFileSectorRef(&_Store.TransactOpts, fileHash, ref)
}

// AddFileSectorRef is a paid mutator transaction binding the contract method 0x5a0e7482.
//
// Solidity: function AddFileSectorRef(bytes fileHash, (address,uint64) ref) payable returns()
func (_Store *StoreTransactorSession) AddFileSectorRef(fileHash []byte, ref SectorRef) (*types.Transaction, error) {
	return _Store.Contract.AddFileSectorRef(&_Store.TransactOpts, fileHash, ref)
}

// AddFileToUnSettleList is a paid mutator transaction binding the contract method 0x34fddaac.
//
// Solidity: function AddFileToUnSettleList(address fileOwner, bytes fileHash) payable returns()
func (_Store *StoreTransactor) AddFileToUnSettleList(opts *bind.TransactOpts, fileOwner common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "AddFileToUnSettleList", fileOwner, fileHash)
}

// AddFileToUnSettleList is a paid mutator transaction binding the contract method 0x34fddaac.
//
// Solidity: function AddFileToUnSettleList(address fileOwner, bytes fileHash) payable returns()
func (_Store *StoreSession) AddFileToUnSettleList(fileOwner common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.AddFileToUnSettleList(&_Store.TransactOpts, fileOwner, fileHash)
}

// AddFileToUnSettleList is a paid mutator transaction binding the contract method 0x34fddaac.
//
// Solidity: function AddFileToUnSettleList(address fileOwner, bytes fileHash) payable returns()
func (_Store *StoreTransactorSession) AddFileToUnSettleList(fileOwner common.Address, fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.AddFileToUnSettleList(&_Store.TransactOpts, fileOwner, fileHash)
}

// ChangeFileOwner is a paid mutator transaction binding the contract method 0x207e33be.
//
// Solidity: function ChangeFileOwner((bytes,address,address) ownerChange) returns()
func (_Store *StoreTransactor) ChangeFileOwner(opts *bind.TransactOpts, ownerChange OwnerChange) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ChangeFileOwner", ownerChange)
}

// ChangeFileOwner is a paid mutator transaction binding the contract method 0x207e33be.
//
// Solidity: function ChangeFileOwner((bytes,address,address) ownerChange) returns()
func (_Store *StoreSession) ChangeFileOwner(ownerChange OwnerChange) (*types.Transaction, error) {
	return _Store.Contract.ChangeFileOwner(&_Store.TransactOpts, ownerChange)
}

// ChangeFileOwner is a paid mutator transaction binding the contract method 0x207e33be.
//
// Solidity: function ChangeFileOwner((bytes,address,address) ownerChange) returns()
func (_Store *StoreTransactorSession) ChangeFileOwner(ownerChange OwnerChange) (*types.Transaction, error) {
	return _Store.Contract.ChangeFileOwner(&_Store.TransactOpts, ownerChange)
}

// ChangeFilePrivilege is a paid mutator transaction binding the contract method 0xb6af10fb.
//
// Solidity: function ChangeFilePrivilege((bytes,uint64) priChange) returns()
func (_Store *StoreTransactor) ChangeFilePrivilege(opts *bind.TransactOpts, priChange PriChange) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ChangeFilePrivilege", priChange)
}

// ChangeFilePrivilege is a paid mutator transaction binding the contract method 0xb6af10fb.
//
// Solidity: function ChangeFilePrivilege((bytes,uint64) priChange) returns()
func (_Store *StoreSession) ChangeFilePrivilege(priChange PriChange) (*types.Transaction, error) {
	return _Store.Contract.ChangeFilePrivilege(&_Store.TransactOpts, priChange)
}

// ChangeFilePrivilege is a paid mutator transaction binding the contract method 0xb6af10fb.
//
// Solidity: function ChangeFilePrivilege((bytes,uint64) priChange) returns()
func (_Store *StoreTransactorSession) ChangeFilePrivilege(priChange PriChange) (*types.Transaction, error) {
	return _Store.Contract.ChangeFilePrivilege(&_Store.TransactOpts, priChange)
}

// CleanupForDeleteFile is a paid mutator transaction binding the contract method 0x9cd103e9.
//
// Solidity: function CleanupForDeleteFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo, bool rmInfo, bool rmList) payable returns()
func (_Store *StoreTransactor) CleanupForDeleteFile(opts *bind.TransactOpts, fileInfo FileInfo, rmInfo bool, rmList bool) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "CleanupForDeleteFile", fileInfo, rmInfo, rmList)
}

// CleanupForDeleteFile is a paid mutator transaction binding the contract method 0x9cd103e9.
//
// Solidity: function CleanupForDeleteFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo, bool rmInfo, bool rmList) payable returns()
func (_Store *StoreSession) CleanupForDeleteFile(fileInfo FileInfo, rmInfo bool, rmList bool) (*types.Transaction, error) {
	return _Store.Contract.CleanupForDeleteFile(&_Store.TransactOpts, fileInfo, rmInfo, rmList)
}

// CleanupForDeleteFile is a paid mutator transaction binding the contract method 0x9cd103e9.
//
// Solidity: function CleanupForDeleteFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo, bool rmInfo, bool rmList) payable returns()
func (_Store *StoreTransactorSession) CleanupForDeleteFile(fileInfo FileInfo, rmInfo bool, rmList bool) (*types.Transaction, error) {
	return _Store.Contract.CleanupForDeleteFile(&_Store.TransactOpts, fileInfo, rmInfo, rmList)
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
// Solidity: function DeleteFile(bytes fileHash) payable returns()
func (_Store *StoreTransactor) DeleteFile(opts *bind.TransactOpts, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteFile", fileHash)
}

// DeleteFile is a paid mutator transaction binding the contract method 0xce904554.
//
// Solidity: function DeleteFile(bytes fileHash) payable returns()
func (_Store *StoreSession) DeleteFile(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteFile(&_Store.TransactOpts, fileHash)
}

// DeleteFile is a paid mutator transaction binding the contract method 0xce904554.
//
// Solidity: function DeleteFile(bytes fileHash) payable returns()
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
// Solidity: function DeleteFiles(bytes[] fileHashs) payable returns()
func (_Store *StoreTransactor) DeleteFiles(opts *bind.TransactOpts, fileHashs [][]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteFiles", fileHashs)
}

// DeleteFiles is a paid mutator transaction binding the contract method 0x432152ce.
//
// Solidity: function DeleteFiles(bytes[] fileHashs) payable returns()
func (_Store *StoreSession) DeleteFiles(fileHashs [][]byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteFiles(&_Store.TransactOpts, fileHashs)
}

// DeleteFiles is a paid mutator transaction binding the contract method 0x432152ce.
//
// Solidity: function DeleteFiles(bytes[] fileHashs) payable returns()
func (_Store *StoreTransactorSession) DeleteFiles(fileHashs [][]byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteFiles(&_Store.TransactOpts, fileHashs)
}

// DeleteProveDetails is a paid mutator transaction binding the contract method 0x52e06146.
//
// Solidity: function DeleteProveDetails(bytes fileHash) returns()
func (_Store *StoreTransactor) DeleteProveDetails(opts *bind.TransactOpts, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteProveDetails", fileHash)
}

// DeleteProveDetails is a paid mutator transaction binding the contract method 0x52e06146.
//
// Solidity: function DeleteProveDetails(bytes fileHash) returns()
func (_Store *StoreSession) DeleteProveDetails(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteProveDetails(&_Store.TransactOpts, fileHash)
}

// DeleteProveDetails is a paid mutator transaction binding the contract method 0x52e06146.
//
// Solidity: function DeleteProveDetails(bytes fileHash) returns()
func (_Store *StoreTransactorSession) DeleteProveDetails(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteProveDetails(&_Store.TransactOpts, fileHash)
}

// DeleteUnsettledFiles is a paid mutator transaction binding the contract method 0x0be09702.
//
// Solidity: function DeleteUnsettledFiles(address walletAddr) payable returns()
func (_Store *StoreTransactor) DeleteUnsettledFiles(opts *bind.TransactOpts, walletAddr common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteUnsettledFiles", walletAddr)
}

// DeleteUnsettledFiles is a paid mutator transaction binding the contract method 0x0be09702.
//
// Solidity: function DeleteUnsettledFiles(address walletAddr) payable returns()
func (_Store *StoreSession) DeleteUnsettledFiles(walletAddr common.Address) (*types.Transaction, error) {
	return _Store.Contract.DeleteUnsettledFiles(&_Store.TransactOpts, walletAddr)
}

// DeleteUnsettledFiles is a paid mutator transaction binding the contract method 0x0be09702.
//
// Solidity: function DeleteUnsettledFiles(address walletAddr) payable returns()
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
// Solidity: function UpdateFileInfo((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) f) payable returns()
func (_Store *StoreTransactor) UpdateFileInfo(opts *bind.TransactOpts, f FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateFileInfo", f)
}

// UpdateFileInfo is a paid mutator transaction binding the contract method 0x681850d7.
//
// Solidity: function UpdateFileInfo((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) f) payable returns()
func (_Store *StoreSession) UpdateFileInfo(f FileInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileInfo(&_Store.TransactOpts, f)
}

// UpdateFileInfo is a paid mutator transaction binding the contract method 0x681850d7.
//
// Solidity: function UpdateFileInfo((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) f) payable returns()
func (_Store *StoreTransactorSession) UpdateFileInfo(f FileInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileInfo(&_Store.TransactOpts, f)
}

// UpdateFileList is a paid mutator transaction binding the contract method 0xd70e6272.
//
// Solidity: function UpdateFileList(address walletAddr, bytes[] list) payable returns()
func (_Store *StoreTransactor) UpdateFileList(opts *bind.TransactOpts, walletAddr common.Address, list [][]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateFileList", walletAddr, list)
}

// UpdateFileList is a paid mutator transaction binding the contract method 0xd70e6272.
//
// Solidity: function UpdateFileList(address walletAddr, bytes[] list) payable returns()
func (_Store *StoreSession) UpdateFileList(walletAddr common.Address, list [][]byte) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileList(&_Store.TransactOpts, walletAddr, list)
}

// UpdateFileList is a paid mutator transaction binding the contract method 0xd70e6272.
//
// Solidity: function UpdateFileList(address walletAddr, bytes[] list) payable returns()
func (_Store *StoreTransactorSession) UpdateFileList(walletAddr common.Address, list [][]byte) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileList(&_Store.TransactOpts, walletAddr, list)
}

// UpdateFilesForRenew is a paid mutator transaction binding the contract method 0x3f2cc9a0.
//
// Solidity: function UpdateFilesForRenew(bytes[] _fileList, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight) payable returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[], bool)
func (_Store *StoreTransactor) UpdateFilesForRenew(opts *bind.TransactOpts, _fileList [][]byte, setting Setting, newExpireHeight *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateFilesForRenew", _fileList, setting, newExpireHeight)
}

// UpdateFilesForRenew is a paid mutator transaction binding the contract method 0x3f2cc9a0.
//
// Solidity: function UpdateFilesForRenew(bytes[] _fileList, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight) payable returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[], bool)
func (_Store *StoreSession) UpdateFilesForRenew(_fileList [][]byte, setting Setting, newExpireHeight *big.Int) (*types.Transaction, error) {
	return _Store.Contract.UpdateFilesForRenew(&_Store.TransactOpts, _fileList, setting, newExpireHeight)
}

// UpdateFilesForRenew is a paid mutator transaction binding the contract method 0x3f2cc9a0.
//
// Solidity: function UpdateFilesForRenew(bytes[] _fileList, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint256 newExpireHeight) payable returns((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[], bool)
func (_Store *StoreTransactorSession) UpdateFilesForRenew(_fileList [][]byte, setting Setting, newExpireHeight *big.Int) (*types.Transaction, error) {
	return _Store.Contract.UpdateFilesForRenew(&_Store.TransactOpts, _fileList, setting, newExpireHeight)
}

// DeleteFilesInner is a paid mutator transaction binding the contract method 0xc43007e5.
//
// Solidity: function deleteFilesInner((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[] files) payable returns(string error)
func (_Store *StoreTransactor) DeleteFilesInner(opts *bind.TransactOpts, files []FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "deleteFilesInner", files)
}

// DeleteFilesInner is a paid mutator transaction binding the contract method 0xc43007e5.
//
// Solidity: function deleteFilesInner((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[] files) payable returns(string error)
func (_Store *StoreSession) DeleteFilesInner(files []FileInfo) (*types.Transaction, error) {
	return _Store.Contract.DeleteFilesInner(&_Store.TransactOpts, files)
}

// DeleteFilesInner is a paid mutator transaction binding the contract method 0xc43007e5.
//
// Solidity: function deleteFilesInner((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64))[] files) payable returns(string error)
func (_Store *StoreTransactorSession) DeleteFilesInner(files []FileInfo) (*types.Transaction, error) {
	return _Store.Contract.DeleteFilesInner(&_Store.TransactOpts, files)
}

// Initialize is a paid mutator transaction binding the contract method 0xbc1c783e.
//
// Solidity: function initialize(address _config, address _node, address _space, address _sector, address _prove, (uint64,uint64,uint64) fsConfig, address _fileExtra) returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts, _config common.Address, _node common.Address, _space common.Address, _sector common.Address, _prove common.Address, fsConfig FSConfig, _fileExtra common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize", _config, _node, _space, _sector, _prove, fsConfig, _fileExtra)
}

// Initialize is a paid mutator transaction binding the contract method 0xbc1c783e.
//
// Solidity: function initialize(address _config, address _node, address _space, address _sector, address _prove, (uint64,uint64,uint64) fsConfig, address _fileExtra) returns()
func (_Store *StoreSession) Initialize(_config common.Address, _node common.Address, _space common.Address, _sector common.Address, _prove common.Address, fsConfig FSConfig, _fileExtra common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _node, _space, _sector, _prove, fsConfig, _fileExtra)
}

// Initialize is a paid mutator transaction binding the contract method 0xbc1c783e.
//
// Solidity: function initialize(address _config, address _node, address _space, address _sector, address _prove, (uint64,uint64,uint64) fsConfig, address _fileExtra) returns()
func (_Store *StoreTransactorSession) Initialize(_config common.Address, _node common.Address, _space common.Address, _sector common.Address, _prove common.Address, fsConfig FSConfig, _fileExtra common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _node, _space, _sector, _prove, fsConfig, _fileExtra)
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
