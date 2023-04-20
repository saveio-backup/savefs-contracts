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

// SectorConfig is an auto generated low-level Go binding around an user-defined struct.
type SectorConfig struct {
	SECTORFILEINFOGROUPMAXLEN uint64
}

// SectorInfo is an auto generated low-level Go binding around an user-defined struct.
type SectorInfo struct {
	NodeAddr         common.Address
	SectorID         uint64
	Size             uint64
	Used             uint64
	ProveLevel       uint8
	FirstProveHeight *big.Int
	NextProveHeight  *big.Int
	TotalBlockNum    uint64
	FileNum          uint64
	GroupNum         uint64
	IsPlots          bool
	FileList         [][]byte
}

// SectorRef is an auto generated low-level Go binding around an user-defined struct.
type SectorRef struct {
	NodeAddr common.Address
	SectorId uint64
}

// TransferState is an auto generated low-level Go binding around an user-defined struct.
type TransferState struct {
	From  common.Address
	To    common.Address
	Value uint64
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"port\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"}],\"name\":\"DNSNodeRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DNSNodeUnReg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"DnsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"FsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structNameInfo\",\"name\":\"newer\",\"type\":\"tuple\"}],\"name\":\"NotifyNameInfoAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"PDPVerifyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nodeAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"AddFileToSector\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"}],\"name\":\"CreateSector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"DeleteFileFromSector\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"DeleteSector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"GetSectorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"GetSectorsForNode\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sectorID\",\"type\":\"uint64\"}],\"name\":\"IsSectorRefByFileInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sector\",\"type\":\"tuple\"}],\"name\":\"UpdateSectorInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0x29d6276d\",\"type\":\"bytes32\"}],\"name\":\"c_0x29d6276d\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINode\",\"name\":\"_node\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"SECTOR_FILE_INFO_GROUP_MAX_LEN\",\"type\":\"uint64\"}],\"internalType\":\"structSectorConfig\",\"name\":\"sectorConfig\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5061925980620000226000396000f3fe6080604052600436106100b05760003560e01c8063931bb19a116100695780639a7d0a281161004e5780639a7d0a28146101c1578063ba921004146101fe578063e3168f9e14610227576100b0565b8063931bb19a1461017c578063955f98b7146101a5576100b0565b8063112346c21161009a578063112346c2146100fa5780632384a6aa146101235780632ba010d71461013f576100b0565b806247c003146100b55780630daf9945146100d1575b600080fd5b6100cf60048036038101906100ca9190618081565b610264565b005b3480156100dd57600080fd5b506100f860048036038101906100f39190617f9a565b610729565b005b34801561010657600080fd5b50610121600480360381019061011c9190617fc3565b61072c565b005b61013d60048036038101906101389190618040565b61095c565b005b34801561014b57600080fd5b50610166600480360381019061016191906180ed565b610f7d565b60405161017391906189d3565b60405180910390f35b34801561018857600080fd5b506101a3600480360381019061019e91906180ed565b6116c7565b005b6101bf60048036038101906101ba9190618081565b611a14565b005b3480156101cd57600080fd5b506101e860048036038101906101e39190617f35565b611fa7565b6040516101f5919061877f565b60405180910390f35b34801561020a57600080fd5b5061022560048036038101906102209190618040565b612304565b005b34801561023357600080fd5b5061024e60048036038101906102499190617f0c565b6131a0565b60405161025b919061875d565b60405180910390f35b6102907feb952083feba7092c52ec86244555c277e74337471df7a58497396ac6140822560001b610729565b6102bc7f195784c0e90535dcb8dc88a0794ef7bb9abc49d02accd0897b9efd42d04e1a3260001b610729565b6102e87f559b8cf9edb2a0f65db21fc2c7378c06d1d699d0fb5e83df1457cff0d4e5846860001b610729565b600061030183600001518460200151846000015161417c565b905061032f7fbb61a01393788562202651365fc64f6c6eafa4ef7eed9a17b639db43ea68e42b60001b610729565b8261010001805180919061034290618d44565b67ffffffffffffffff1667ffffffffffffffff16815250506103867f910a74511ed667cb2e524533e389233a0daf739a5d1beb9e5d6245a23936f73d60001b610729565b6103b27f916c1e4ffad4d697867fece4567c854d326935c5085cf858af25a4d5fa0175f060001b610729565b81608001518360e0018181516103c89190618c0c565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506104117fcc3bb648b3077836cc0ed9c8d76e8783a7e010d8b4dc6ad8e5deb880076608b760001b610729565b61043d7fe688fbdd40994e5f8a50fed49729dd2d927dbb57bbe3655ffac2c99d987d956360001b610729565b8160a0015182608001516104519190618b96565b836060018181516104629190618c0c565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506104ab7ffac408a9f30e6d11f4c6cb8430a26863833fbad54bc182e7b558b55f14472f5260001b610729565b6104d77fecc9aa0421f4dc3119d0b8b81e93afb6140987c6e93d9fbe96f06487a474204160001b610729565b8015610565576105097f4dfb28646a37d6e6a0b5ec55089f102b217b4318df6a08e8563efb6ed7a42a6a60001b610729565b6105357fe0c82aa8b6de3d2f182d6e1ec09fd5bbdd67e659f1571f32e2bd714409369ad760001b610729565b8261012001805180919061054890618d44565b67ffffffffffffffff1667ffffffffffffffff1681525050610592565b6105917ffef0aaac5152be3ea501022cfd01c35e0ea3aebf41e26823bff9b2617fbe685c60001b610729565b5b6105be7fbbbc64262f33e7f72c267ce18adb0880ac5541498174bfcb2787dbb9fd56e5a460001b610729565b6105ea7f05003c0d861dcd2625df9bb55a5ffea39d660873e6952f3c7b827c78eced624060001b610729565b600083610100015167ffffffffffffffff1614156106965761062e7f96267e95db7d239d639ad8b27db8efc947918c9b230f8317b2ba3b0cacf8914e60001b610729565b61065a7f9aca5a6ab1b12736aff80257289c18563b5d0fdc9e1ba11c7246a9578ea0081660001b610729565b6106867f3701e837323716cdcf753bfd5ecb43bcd88c7ae91f9110e2387a3f2065fda0d260001b610729565b60008360c00181815250506106c3565b6106c27f3688637dca373c889ea5c092757c6436bdfbd618f5de590801405a7492c248fb60001b610729565b5b6106ef7f4710ac3fb072cd1148af2fb6d9cf1d01a2d647766a62041f16fff4ddfada6f3760001b610729565b61071b7fe08c2677b247446db64e59a0f0a6cf41cd658578da4e1002a0faee5f081324f360001b610729565b6107248361095c565b505050565b50565b600060019054906101000a900460ff166107545760008054906101000a900460ff161561075d565b61075c614915565b5b61079c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610793906189b3565b60405180910390fd5b60008060019054906101000a900460ff1615905080156107ec576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b6108187fe9d33170bd4eff341d80bb0574dbdb5f882fdfae39cc97587855171c7da54db360001b610729565b6108447f48db40ba5313edf6667342fdde0440fe7e0037fd7efbed7f054639e1bc676f1460001b610729565b6108707fdfc4e791214d51aa3fc0c7833cc99bad91f5a5cb75ca92246552d0896688a74860001b610729565b82600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506108dd7f4d5a5856f3df4084622094fd9ba690e5a80007ae585418c50e101b7a261ded4360001b610729565b6109097f3bde1d56688725600227b62125548bf62bee8bf22e7a161da4dfc10d4bcdd48b60001b610729565b8160000151600060166101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555080156109575760008060016101000a81548160ff0219169083151502179055505b505050565b6109887f75c626efc38084d79f8ee6df9dcb8c89a28fcada25d6e40448d7584a5a1a2ffa60001b610729565b6109b47fbf26b259cc6089b46dd0547389f9b82084e0d0a666cfb0c81ce70f342a4efacb60001b610729565b6109e07f89732853787499aed52b3d410dd0f584dd3448418280e1390424510931f3067b60001b610729565b600060016000836000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050610a537fe2f6192bed5bbcd2b37adda254fb83dbfd747e78a79e6cf13ff536d1eb61cb6c60001b610729565b610a7f7fba0adf568209eb1b0f7546d2b6020b6361feb9fd4ec892f51570f009796a0ebb60001b610729565b60005b81805490508167ffffffffffffffff161015610ecd57610ac47f539f6281ead98718fd8664a0309673dc2432996e5ddb9569a8a49922279b8fed60001b610729565b610af07fd69897fffe3296e501ccd99616ff7873fe4e24c0de134bd8767477f710054e0060001b610729565b826020015167ffffffffffffffff16828267ffffffffffffffff1681548110610b42577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160000160149054906101000a900467ffffffffffffffff1667ffffffffffffffff161415610e8e57610ba47faf850519318c82c1c3fa3a3e53cb01f8b19d5497eebb28bc858d3451616892c660001b610729565b610bd07f27240eb0b147be739c82319c591b7b9b940fd4a64d8241e9556ad3f8d5b1941260001b610729565b610bfc7f2664b230457a8bd2475494e808be7345737e1252a423a09b042a51062a8a584b60001b610729565b82828267ffffffffffffffff1681548110610c40577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160106101000a81548160ff02191690836002811115610d73577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060a0820151816002015560c0820151816003015560e08201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101008201518160040160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101208201518160040160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101408201518160040160186101000a81548160ff021916908315150217905550610160820151816005019080519060200190610e59929190616e5e565b50905050610e897f8f4354772d303d981acf9984e9cf6337ce21361a27c0ac2450d45c80c834458360001b610729565b610ecd565b610eba7f03989da341b6b0161c92870e82c26918f7116de0d992d5aae104f171b87efbb360001b610729565b8080610ec590618e1a565b915050610a82565b50610efa7fe63b7ad17da268a53d89360c6c49719c0a21276c4c2de96b6e7fb5932dac155a60001b610729565b610f267f9c59f08a3f2705bffb642a6b736a54362adfec291def39e601f34e7989ab864b60001b610729565b8060016000846000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020908054610f78929190616ebe565b505050565b610f856171ca565b610fb17f6cc8cac6c7d675e2879b10793621ecd6737e0bc4e25e65c0ffde20a369afb07560001b610729565b610fdd7f30a2d873ec9a0e05405a64acd62abf0a8e2a05dfe7324de2f0aa9911267cd99c60001b610729565b6110097f3269afd974400e9b2276e36c70d29835220fe23a27dcb61dffa7571b9dd1182e60001b610729565b600060016000846000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156113b25783829060005260206000209060060201604051806101800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900460ff1660028111156111c9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115611201577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200160028201548152602001600382015481526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160189054906101000a900460ff1615151515815260200160058201805480602002602001604051908101604052809291908181526020016000905b8282101561139b57838290600052602060002001805461130e90618d6e565b80601f016020809104026020016040519081016040528092919081815260200182805461133a90618d6e565b80156113875780601f1061135c57610100808354040283529160200191611387565b820191906000526020600020905b81548152906001019060200180831161136a57829003601f168201915b5050505050815260200190600101906112ef565b50505050815250508152602001906001019061106e565b5050505090506113e47f3e6de94e58de63d9f3f71a3ad0e704606558a7585b7975d73812d335249a988360001b610729565b6114107f60ba2b4f9b9679593e2289bd0d90ab55d03757e9572b525deececd9ea351167d60001b610729565b60005b81518167ffffffffffffffff161015611603576114527feb655aea1f02c983e1e6130df954872794c1db8850260f4663edca1b51e2b71160001b610729565b61147e7f536db5d5bf5214a700437f965d02c21e3535de60e7e8888c812d360385f47ac060001b610729565b836020015167ffffffffffffffff16828267ffffffffffffffff16815181106114d0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015167ffffffffffffffff1614156115c4576115187f639074483d1901c3018c2d625b59bf3868c6dad4f3ea6d1d18a51c34d9eb4b8e60001b610729565b6115447fccfcabe0834e39f5adbf1acf5e736de0876629cd8af9a36526f57c078ba0b1ef60001b610729565b6115707f6d93f10669a26f0cdb8c1f236631cb77b4718359a5ddb98de8356f07a02d995d60001b610729565b818167ffffffffffffffff16815181106115b3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151925050506116c2565b6115f07f3c18aa6ec1cc88706e1ce4163760623015c4794a0aa7b80275441b8d6f840be360001b610729565b80806115fb90618e1a565b915050611413565b506116307f3fe3dd61f08cc45d16d7dc36e5e73a8929756f0fa1f5d02aec259c47269042f060001b610729565b61165c7f74468df947eec6f029a9b2bdabb1ceee9cfca11891d5f71ddea9b5a03d7a14ea60001b610729565b6116646171ca565b6116907f10badb172aef9c46736c4fc3741dd5db3069f36669c5b35dd84419d6ecaa7dcd60001b610729565b6116bc7f29f2e258fa0f711f324b45a6e1a5755f906c071a6260665e84516a523dfa78e160001b610729565b80925050505b919050565b6116f37ff6519937eb9683a7cf49c9c29a289e04fbb915b5502c89f76a9b08a8fa77367360001b610729565b61171f7fc3cc7288d3c61427a25c67609ceedcfeadec42ea2dc2f565fd5be42b86e4e58160001b610729565b61174b7ff851c8ecdcbe34918871e4ad4c1100f471a4d35bc8d175118eace55ed70d2b8d60001b610729565b600061175682610f7d565b90506117847fb4c87b314edf9731430514561373c649022716e3856dc678b73662113e36c0f160001b610729565b6117b07f0a488e8187669845b26ed4dc8e482ecb3ba717817edc6ca4c1d97dc2a9008bd560001b610729565b600081610100015167ffffffffffffffff1611156118df576117f47f14144e52841eddb0dd0eebd9d1dac1cea49843cfbd214b7c1032d45bd941697560001b610729565b6118207f09a59e0e286b4fef5721f194aebb25de740784ea9ced76a95be2576b69d3300060001b610729565b61184c7f9dd22a9b6f3af6e07035e733f938e0e9b5475dbac5479613e69c903562e58a4460001b610729565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f560405161187990618980565b60405180910390a16118ad7fe768f81d8ebcd8ce084573fcd3e8234cfbfc295af0fb8f600cb4e3efe48387a960001b610729565b6118d97f3ef96621b4c5882000ba282badf407dd1f943f9fe4868e106dc97798a16a9f3060001b610729565b50611a11565b61190b7f4079f1f78db03ab4bea16a18d9184ee88d3856c67687387634ce2dbf85ebe8d860001b610729565b6119377f6e124a1e13bde74783ba6a9b1d25797e829212c8671f134ae84009aa8867de4660001b610729565b6119637f2509dab7c4fe4b3b62e2e6b7023dee855426b88c0e720af95d70bf89d7f85c7d60001b610729565b61197582600001518360200151614926565b6119a17f6367699f2d5fbc1730f96f1841b6586ff68fd72db87f2b873080488fecc38faf60001b610729565b6119cd7f867fcde07dfa61d1b8d53210288f905923ba80fb02376c92d9e3f30385d4e6de60001b610729565b7f4fca90fd1b1cd962cf67f21703f1380572181450811cdd09c4add7ed1cb0c12c600943338560200151604051611a07949392919061879a565b60405180910390a1505b50565b611a407fddb214da9e849fda34200c661e404152b3f01fa7837fff3e33e9ac117f6ebc7960001b610729565b611a6c7f49d3651e8914f35c7699420a3e4c9be84e747d0dff4b8cf8213a0839870fa61260001b610729565b611a987f40b0ad238fbfbf191c4b826e4c60a829f72a00014312c3440df6534a605b405d60001b610729565b816040015167ffffffffffffffff168160a001518260800151611abb9190618b96565b8360600151611aca9190618b58565b67ffffffffffffffff161115611bf057611b067fc0fc4ac426e3baca4cc161867a9ff529b2f75054ade7b6be74c41e9b768659e260001b610729565b611b327f78eefa34481002f2f560724c5926b071eaf6a422846edec35273e22b413502b260001b610729565b611b5e7f9b25c91c09561f9457a054d598917d7eb31de50c91391141038e11fede35687360001b610729565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5604051611b8b9061894d565b60405180910390a1611bbf7f978a397b385931ebb52617af701560995f342482bec77849355c4d51160ae68c60001b610729565b611beb7f4f6ad527ee81acc0b5f88bba4489fd286042edb12076286254596bbbe41dbd3660001b610729565b611fa3565b611c1c7f1a4bd2d48c34418bb6fd448e7fcb02b85875eacabb4e1da5ff4a752ddf3cddae60001b610729565b611c487fdc12ef9f644fc321ffaa3a7d38ac2ebf0db96eccf69f9a396e6aa5d160e4744f60001b610729565b611c747f807727d6f91c5cebb8553be2f5c14d2ea8e9708bae265e82373b5a4b7cdfe20360001b610729565b6000611caf83600001518460200151604051806040016040528086600001518152602001866080015167ffffffffffffffff16815250614a1a565b9050611cdd7f8235d736916096e59afb8ea0bc693af38c97c68e0b3c35180b7feff0d5d01dc560001b610729565b82610100018051809190611cf090618e1a565b67ffffffffffffffff1667ffffffffffffffff1681525050611d347f8bbf0cd23506fff915235ecb9bc4bc04bb6d505588804c283c7f14c4c51643d960001b610729565b611d607f7ebff368104a11152676102a6676ad54c6ab59a3b4db8b34d4b0dbc30c09dbf660001b610729565b8160a001518260800151611d749190618b96565b83606001818151611d859190618b58565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050611dce7f44c22ea3032cf7c5ac6a6423961b1537059142d6c250854d671b49d1330c278360001b610729565b611dfa7f6565241b2c877aa1a4aeb48a055977a33a45ec2e4abfc7dfe84256a731ab0dbf60001b610729565b81608001518360e001818151611e109190618b58565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050611e597f23bfbf198bd6d4caad84ec0a1a7b13ba8fb1031f2237c0b475c6eb91ca4a78b660001b610729565b611e857f419b0a745f3401cb7584e51da95a78aaca6ec87424a7835cc84fe8249fa0c67d60001b610729565b8015611f1357611eb77fbf450a765bf8a736a40b1c89218ae9c6980e345b6613b53a0a8d59c4261b07ca60001b610729565b611ee37fec0fa697135557320a726401463e45b07acd495c16b26255355f0c4e293dfc9d60001b610729565b82610120018051809190611ef690618e1a565b67ffffffffffffffff1667ffffffffffffffff1681525050611f40565b611f3f7fc0011a2ac3c6c3df310c5b5d119b0b5ec5d1bd467b13f66b0d557bf329c04f8b60001b610729565b5b611f6c7ffce46f3db6e113c42e44ba9162e5e38d62b21924b083c4131f691dbe9da731c960001b610729565b611f987f728d9043b7d96a6514a78c3ef4ed18799ca1bad85bb2e658235fb72ca7d6f07760001b610729565b611fa18361095c565b505b5050565b6000611fd57fa911d3e14cf6c46bf1166d52c2f796b8e706a2d617d91197da9225fdbd2cd32b60001b610729565b6120017f175ce577e779be7bc2456f730cb9500beab6b41a94ef453e832916fb580d639f60001b610729565b61202d7faab7bd44d28842261a16e83387b7dacc7511b3651714f9491ab6579366412aac60001b610729565b6000612038846131a0565b90506120667f936874d05462bbc4e6ffc6643bc8a421961ca09cbe3c5bb80b827f5532e02f8260001b610729565b6120927f7c03669cf846975971f0bf30af537f4b255e0a126a8fd0ead086794d95fe1d5260001b610729565b60005b815181101561229f576120ca7f52f742130a67a071fa8daedc79b756dfe4316fa820772a87f558ab18752a255060001b610729565b6120f67f608c959a4aa863a7ed85af62781fff85bf510af904e90c5d844c87079ff2e78160001b610729565b8473ffffffffffffffffffffffffffffffffffffffff16828281518110612146577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff161480156121cc57508367ffffffffffffffff168282815181106121b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015167ffffffffffffffff16145b15612260576121fd7f36807f9a61399157db56bc0e1e170467b89a6c0af1334be3575f2e96b279657d60001b610729565b6122297f3ee45caf904e724fd79bf03c4b02808cae8abd06711e53d05b3a050437737abb60001b610729565b6122557f7fb3c617f58689a23d1a790430efd4fa49e5b7c2406cf12dcdc995f035bc2adb60001b610729565b6001925050506122fe565b61228c7f2b97b10137a473de55a2b75be639f0f4cf3f5da355e4bdb985e28ddef31b51cd60001b610729565b808061229790618dd1565b915050612095565b506122cc7fdf4a84425c8cfc13ef3cda35c4a3d565aed5bc645ecfe662608fae25208b7ae660001b610729565b6122f87f4b7e051eefe535df6118b303df61c13456c63356668cd40078a1f27711e5db7c60001b610729565b60009150505b92915050565b6123307f39e62894578a7c92076036a9984514aa1d79f3d5b3204cc2c59203bc1af5f93160001b610729565b61235c7f1b0170b5479c63117a766b14aac6e507b1ec53a0d365b9d2a005a173222367f360001b610729565b6123887f634e04efcb22428d9db75270c51908916256bdf2d2dedc102b3f46137a0b8e2660001b610729565b6000816020015167ffffffffffffffff1614156124b5576123cb7f27b433d70f01d0655da81889ccd02f2c3c3a4f6b50ec19faea121bf5429a62b160001b610729565b6123f77f66c5a6586cccdc30bda03c5fb49fb02fd62403367003edfdf4ffa473cbecb5c960001b610729565b6124237f3e958c3c0afbe3a70e58fea1c2d2ba99fdff6e739cbfd2dd89458e1a44dcc64660001b610729565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f560405161245090618881565b60405180910390a16124847faaa45d2a0185dd4e12b64d44feff7cdb71ed48f6c1eabff96ec9ca257f72f0f760001b610729565b6124b07fa9bfcc12713553287d0449f950ff3a771272b85a94678e42377449b4a2f397ec60001b610729565b61319d565b6124e17f46b86f1727584480722c1fd1c07c89e3f3d93999c2e846d0f8823fbcb04462ad60001b610729565b61250d7f87491697d8a4debb5bced2cb6e6344b89b8976750b0e0339958af00464f08c3e60001b610729565b6125397fd9de970421b77b6433cc6e43d8c45a1226f2898a133535472b36700a89409fa360001b610729565b6000816040015167ffffffffffffffff1614156126665761257c7f2aaab7742ab5e38db74dde3dfbf39972851252b93750ecfb36d7c1fe9e4d450360001b610729565b6125a87fdef1b32d71d7a4340bdca1f00a1b0d89caf052b87e67e88171242d24d01553e560001b610729565b6125d47f521b7aa02e24b01674d4e02d0164ead315f22b5bd3ab7ee661a05f42436c9c3760001b610729565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f56040516126019061891a565b60405180910390a16126357fa3957cc699f514fd08ef32df91834694f4bcca4c79a3409ec0f31debab9059c860001b610729565b6126617fc730b90893a81805f5d6cbe1b6d3692fa88fe99f487f963dcbe5118e9e2c98bf60001b610729565b61319d565b6126927fd39dac6676cdecc1b08976d127ff1e2defdb2e73f6a1c5db33ce864d2f48f56f60001b610729565b6126be7f58826982c34acbdca9f628283e92e38092dcc31e0ac1ffd23ad44fe5cfc1e90560001b610729565b6126ea7fa3962079949f8f9b5a24e4886ade3343f4ce5e4952cecdfff165e2c37f548c5760001b610729565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636683899482600001516040518263ffffffff1660e01b81526004016127499190618742565b60206040518083038186803b15801561276157600080fd5b505afa158015612775573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127999190617f71565b6128b3576127c97ff69865c544b7c7302f40c06fd97f69fec8b2e61ca32e28a1c17fc64b6c43106c60001b610729565b6127f57fb98484254ade14368b4202e45cd3220454af3fde8491b258b94ba75f7373be9e60001b610729565b6128217f52dbe9a5403e90fa122a940f1a9f3cda93c9a401fd4e83286a5c02f1843d8bea60001b610729565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f560405161284e9061884e565b60405180910390a16128827f42cf6aaf4bbc1bc0a032ef61bbecda5fca711b28c87fd9334993614193ce111360001b610729565b6128ae7f823ece50824eea773ea841829378c72bbf26c84df2c2b20f45f912da47ae863960001b610729565b61319d565b6128df7f88267011accd2f661b7387bab6e271c561d7ff23d4acd81976c602f31bd511b460001b610729565b61290b7fa3bdf6526d9f90c425a407afad3869c1e658145bf14f99a5ffb1c1ce0548530160001b610729565b6129377f71b2aea860f1799db345dab18e797fca3d86949c5822dde6cf8e22d8a8229e7060001b610729565b600061297e6040518060400160405280846000015173ffffffffffffffffffffffffffffffffffffffff168152602001846020015167ffffffffffffffff16815250610f7d565b6040015167ffffffffffffffff1614612aa7576129bd7f89e4660d44bdb431a0d11d1cbe59883822e9e9cc0ed4f842fadc16d45abc752960001b610729565b6129e97f0b1562ee637ea86618cb529a53590eb27429f035fe98907eacd17e7680b355d560001b610729565b612a157f4ae870ed93e62ce5c8ad047ede430a7dbc177c5d653d1ed1ac09823a288c399060001b610729565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5604051612a42906188e7565b60405180910390a1612a767fa5b2ed468e511015701295f3cd191ad5c638ecb991367494b9990a3541b1c68c60001b610729565b612aa27f35ca5cf06b6a8b4459c98724b12916bb567fbd0226ad97de31752b9a401b48d760001b610729565b61319d565b612ad37f816099082480c536887ae0312402cf0e460aa640e6457cb81afb11862f319ac760001b610729565b612aff7f5606baee1a0e5a44fcfd8cca807482bf8adb5a67c8a49f865ba4ced6584119dc60001b610729565b612b2b7f74a98de4bf33751a682ea76669b5f026465b5f687bd9cad5bfcc0e725d77ad2b60001b610729565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a65374a83600001516040518263ffffffff1660e01b8152600401612b8b9190618742565b60006040518083038186803b158015612ba357600080fd5b505afa158015612bb7573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190612be09190617fff565b9050612c0e7f592419a190fc9e02b88e3fa8659098639a3672a1b12c61201966692a6249308960001b610729565b612c3a7f9b1cefa54805b5cc286f56133602056a7e431f0b3578411c23822e70720e30f060001b610729565b6000612c498360000151615375565b9050612c777f5712560e7c97660c8876b2d5681dbb21ba3af45e7e41a96074216717e491729b60001b610729565b612ca37fec521717318424adc12d694f54d6559290e9045cd5485f525fd7b6d4f93ceab560001b610729565b816040015167ffffffffffffffff16818460400151612cc29190618b58565b67ffffffffffffffff161115612dea57612cfe7fc10e97725937950a6262f0fdaff9b94079cc12f2aeba9d73905c04ce215fc68760001b610729565b612d2a7ff691088d1b30117cfef676c773c9d348c6e0606d1b1673fc5351e218bd8bfa3560001b610729565b612d567f48b9624effe7778534fff7919e64c1b6f78c4beafccc0021556c04fd308ee09d60001b610729565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5604051612d83906188b4565b60405180910390a1612db77fc9b59fe6a63fb8674c751261ddad76a9f5e7a547376f55bfa298adb337d229d460001b610729565b612de37fdacaf73088ea9a8d01f6c8b8c79d297de29e6ff57e203c9ea615cee0a8e4975f60001b610729565b505061319d565b612e167f6856c4f32e7a4cfa1f0199c4a2d126ac2769e9234d9865cddd4b6d02caa1db8660001b610729565b612e427f9601ffedcee56bf3059c3ef43d778f83ad16b04446e1df82d1cad2b7bf2b3dd360001b610729565b612e6e7f86e27efcdebfb4d5474f2eec7d2a0edf30b0dd8aea892b58e1813c138765099060001b610729565b60016000846000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002083908060018154018082558091505060019003906000526020600020906006020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160106101000a81548160ff02191690836002811115613000577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060a0820151816002015560c0820151816003015560e08201518160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101008201518160040160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101208201518160040160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101408201518160040160186101000a81548160ff0219169083151502179055506101608201518160050190805190602001906130e6929190616e5e565b5050506131157f9524614ba9167109b543aaff7439f5c275e683d4a4a2acf285436365c06b538760001b610729565b6131417f2142eb5a46e020608336893407d86e7475cb1a1e91ec8ab7c2c5e482ffdc624960001b610729565b7fc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb600843856000015186602001518760800151886040015189610140015160405161319297969594939291906187df565b60405180910390a150505b50565b60606131ce7fbcda5ba371a7f69c07e2405d070f272bcefe0f09faf32e75738a58f243298c2260001b610729565b6131fa7f5a06cd60fbec2b62896818ed8b539ec955911452b26b5852c205688307142b2560001b610729565b6132267fd2f981150a84495ce28b5d13ff45df53d5975a8c32c073bec1933b082c00697460001b610729565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156135cb5783829060005260206000209060060201604051806101800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900460ff1660028111156133e2577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600281111561341a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200160028201548152602001600382015481526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160189054906101000a900460ff1615151515815260200160058201805480602002602001604051908101604052809291908181526020016000905b828210156135b457838290600052602060002001805461352790618d6e565b80601f016020809104026020016040519081016040528092919081815260200182805461355390618d6e565b80156135a05780601f10613575576101008083540402835291602001916135a0565b820191906000526020600020905b81548152906001019060200180831161358357829003601f168201915b505050505081526020019060010190613508565b505050508152505081526020019060010190613287565b5050505090506135fd7f890a20cf6f3c9344a3ad8e6dd25a08e3ecefae672485425e86e6ad834ee2ffc360001b610729565b6136297f7727d2e92f696e5e53db8a810bf820b86717893a6b82d8c4d9a5e7e9257ae9e160001b610729565b60006136577feabf93b2441e6a7f74a5b8b67e3b78210042b7b7b67d6ccfe266a4e5616c8ca160001b610729565b6136837f275a42e0474ece248be89a88f6e8e2bceade28b6de37f7b3960267575444500960001b610729565b60005b82518167ffffffffffffffff1610156137fd576136c57f2e6264b91d592038f6ad3ab797ce3121f46b4652d246adf7b9f53f602e0701c560001b610729565b6136f17f42208ee0009085148bb58358f48a2d5603fe2102b7fa7d3f3bf762ecda3dd94760001b610729565b6000838267ffffffffffffffff1681518110613736577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040015167ffffffffffffffff1611156137bd5761377e7fd97be3eca8181264c836e317eb8d1b38e93ee191ab018d31eac433bde5d370fc60001b610729565b6137aa7f83f72c62164435d35168399c71b73ff022c2877a5ff78c0393144e71bd7a2daa60001b610729565b81806137b590618e1a565b9250506137ea565b6137e97f2539917ac74c673d6dff5ba6cf934198adb3bc7e2058e95a1025d897ed00faa060001b610729565b5b80806137f590618e1a565b915050613686565b5061382a7f6a15aa123246fe313e0e3330356d22a5e2569966e457bc7e5820db7d14f0260660001b610729565b6138567f4b1a25695416685362af1bae8b545ac65d29008bb29df66f903f6bc324a25d3460001b610729565b60008167ffffffffffffffff1667ffffffffffffffff8111156138a2577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156138db57816020015b6138c86171ca565b8152602001906001900390816138c05790505b50905061390a7f06afa8adf80c65625622b58ead406cbd2c05667d6e6d549c807cd2f2c26345e960001b610729565b6139367ffbdd17d45e5a6c1a8f1e52e1abc90afb16600c875858fe7e15122eba4b17c24460001b610729565b60006139647f8c0dc30327d43cdb5dc53a63d5931f520883b9ba25d5ea72adc2e355234a715060001b610729565b6139907ffdf6c0ff35317e305b5fd095eefc97c0e0c61d83e9cb10f654fd352974dca8dd60001b610729565b60005b84518167ffffffffffffffff161015614117576139d27ff6259bb1bda64ffcf546dcd2c4006690ca2e83de344dd408c0558570f33551e060001b610729565b6139fe7f8272ae08a4a198db3b3dd99b557dccd929f13bf561717614f3c639d3a757a4b360001b610729565b6000858267ffffffffffffffff1681518110613a43577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040015167ffffffffffffffff1611156140d757613a8b7f0ae256cc92f989f741d7a12d01686c96206f6e0aba044723965e0cebb819ebf160001b610729565b613ab77f37d4e9372d20cdf1d3fc620fc83dd393a50c02512dfa983d187afd825e71bd3a60001b610729565b613ae37f248176321fb360f7890b7c05b8a23d3aa47f17caa55aa7ea06c492e694038dcc60001b610729565b6000858267ffffffffffffffff1681518110613b28577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101519050613b5e7f9d3da89f4c831cdef94e9f7f8ffec67015837508289c1fc40a864e5f1ce756dd60001b610729565b613b8a7f56af549ca8efc942405965a9dbf5708f9e9d429166332b6d2b08ef68129c925460001b610729565b6000613ba0898360200151846101200151615984565b9050613bce7f859fb220f60898b3c60d007adecff053fc136836b96cf70f35213dc4dda2a40760001b610729565b613bfa7fa19efe1068aacba148c356cd8e06f8103bcce739b37f02d064c36db0bd8271ee60001b610729565b600060036000836000015167ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b82821015613d3d5783829060005260206000209060020201604051806040016040529081600082018054613c7a90618d6e565b80601f0160208091040260200160405190810160405280929190818152602001828054613ca690618d6e565b8015613cf35780601f10613cc857610100808354040283529160200191613cf3565b820191906000526020600020905b815481529060010190602001808311613cd657829003601f168201915b505050505081526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505081526020019060010190613c47565b505050509050613d6f7f3bc744ad2b08ab0b6ab78540a7c5b706a2e12174ef854dce1997d9a58b2fd6e860001b610729565b613d9b7f8bd0eabb317b749d9015748561e8ca340cfac2118430133465b61ca2f77490df60001b610729565b6000815167ffffffffffffffff811115613dde577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015613e1157816020015b6060815260200190600190039081613dfc5790505b509050613e407f46c6095ac2a21c019dd62ba1c618104e4bc94c03a6ed728fa11734384f67277d60001b610729565b613e6c7f87ebab337c8b6fa32ee175b6790b3d69dac8ab3a5277d2854e051d87690b9a0c60001b610729565b60005b82518167ffffffffffffffff161015613f8a57613eae7f339b274eb512f8bdf6045c2503d3b86c1d3f20362a5e086ee39556693c25ae1360001b610729565b613eda7f8b46055fdf26de0ac251a9c560042156540785a912e6fc62cedf33a5ffd2207260001b610729565b828167ffffffffffffffff1681518110613f1d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015160000151828267ffffffffffffffff1681518110613f6c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052508080613f8290618e1a565b915050613e6f565b50613fb77f4e95b7411e113a0b9c10176e2a0afab24377ec2728330900b6aae782ed60f13960001b610729565b613fe37f6779a55fedaae25675101536cbeca85715f19ab3ddeeecd479ded2ddc893482360001b610729565b808461016001819052506140197f818c773f3425a7348647b99a15eef74e8dfde17a23b1d99723beed6ba19318b960001b610729565b6140457f9b0ed96866191bbfe074e66b3475a217ce310aedd162bfee5f3ea5ac0c0e2ac160001b610729565b83878767ffffffffffffffff1681518110614089577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052506140c07f4b18d1f42f31d265732b3dbbeaf51046eca941be91644aa466958bab5260fbf260001b610729565b85806140cb90618e1a565b96505050505050614104565b6141037ff4a2ebfead1e50018a53831afe07d662d30c820498f94e567b66743c2ca83a9c60001b610729565b5b808061410f90618e1a565b915050613993565b506141447fb821db5a7723d7c71d9c5412c56a0b30f14fb229da71535af4cf9a1cd23374e160001b610729565b6141707feb3e6fa3984ad2d0a4f6e3097f9af9676e0bafe4a9c61d14617f31c57dbe37b360001b610729565b81945050505050919050565b60006141aa7f62bbff9b3b2a34aba10ed417d231e684d95e571603d7b2ed9c9450cccd31dbf060001b610729565b6141d67ff0c74ddf226926618fc64921427c651c0e82438b427f3a355e2507ffbbb44fc460001b610729565b6142027f14d5dbc7db0c9a9f32cafa409de22830176bb0b0fcbd80e12ca4cb1fb0d06f2560001b610729565b600061420e8585615c19565b905061423c7fe2dfc218977cea9ac17b547e67071c8cc36321643194c1c6423ba8996e14592660001b610729565b6142687f9d53afa3c3647802e11b2059b923c7e24ae0dcccfa6a71458cf2e281331af94860001b610729565b6000600190505b8167ffffffffffffffff168167ffffffffffffffff16116148af576142b67f2955fe9c793e5b157a8d2baf4e959a8f801d43bd292bd798904cf4ca839437ca60001b610729565b6142e27fa41a928131529231433bd1ea420f31e6dabdd98ad8300e595f8da6360f6b6e7760001b610729565b60006142ef878784615984565b905061431d7fec56e19d75a8998c1d8511e446feceadc059e595e2048b924af56aaae2dacad460001b610729565b6143497fb1393472fcab9fafa961892267f8f8d51ab5b1a7e0ab6c96e2c08ccb72f6017e60001b610729565b6000806143568388615d47565b915091506143867fc8aa6dbb2794e7ac7856362eb7e317fbeae93abd114f0f365feda73f0aafee0160001b610729565b6143b27f9628d0cd33c7d02cbcb5421bf2a2c3429a7de44ce36c5e10cc4e046162f601c460001b610729565b801561486d576143e47fe5af8c406e0992131ab58e8d7036df3bd0601d7e7ba7be1b22f3533acb9d033e60001b610729565b6144107f46e22851f39ca2a7eb703b697ca0f0db377f4c41575411253bab76f3fe692b9860001b610729565b61443c7f28e562253317db5fa93fd2bee9f7483234f05f6b8ec2b5750491d1df49a6be8c60001b610729565b600060036000856000015167ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002080549050905061449b7fa875e28d8f7d66293998449c11401252963e3935847b1930b74a0696cb6f6a3f60001b610729565b60036000856000015167ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000208367ffffffffffffffff1681548110614508577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002090600202016000808201600061452791906172b7565b6001820160006101000a81549067ffffffffffffffff021916905550506145707f132c945af6a17a2b1fcc7066a234dd9247c13cb4c91f4710c7e39aafc01eae0660001b610729565b61459c7f7924361b0f88199ff7f89b4ef70f177215b836bf72b2d28525317e1ec55b4e0b60001b610729565b60008367ffffffffffffffff161415614641576145db7f0a83a4d45c3fe6fffa74d000c2af53f94dbb42a8aac07c14b7fe67799cb47f3d60001b610729565b6146077f416269f96c478181050db99564237eb7f556a362ccf90debe6eb2c5add21829560001b610729565b6146337f211298ea21a56cd1eccba5f9e81a284dc4eccfcfbc1e48ce3e770eed83b10f0260001b610729565b87846020018190525061466e565b61466d7f8639293c1756487dd9e9b14919a8c09105dcb1426b39e9c639e892f73896de3660001b610729565b5b61469a7f1a13171ab6c0dc308efa4284d442f861f791f51a02573eb7348e17eeafb5a46a60001b610729565b6146c67fb64938aa6ace3098ea3f196a15ae8e32bae8e4c9ba4571131b4f38b4c86e8b9e60001b610729565b6001816146d39190618bd8565b8367ffffffffffffffff161415614776576147107fe787cbeccc972df7b6198f1e94cdecc640aa3c2ed4af97f5f6013c0f2807a1c260001b610729565b61473c7f6525de481b85d76d7f9386146094c42540fa79b9fb9d84f52bcb07bd347c3c1260001b610729565b6147687f4e466086cdabc3f77be063c9dced88f0aa63ee1c06e267b8b8a82d10056a8c6460001b610729565b8784604001819052506147a3565b6147a27f39ee405d4b60a9738c3a84f03e1153a121c2e8d980a65fb1524899b69851ffed60001b610729565b5b6147cf7f4c6674fa2a055cad830993b1c5a66dee9d8c12f9a5485cb522a25dddda9dda4360001b610729565b6147fb7fd56ad3d51bd56e05329d217f082efa63e15cc3d68c243816e7d0d65cc74cae1560001b610729565b6148068a8a8661629a565b6148327fa1d8091c4ba5bae5f6ff4622edf5f9520c7d985f5a6c297e06d1500ef87f515560001b610729565b61485e7f0f4398314955e2058e07c9475b52580c08e663c03f10da52991e65a16cc25a9060001b610729565b6001965050505050505061490e565b6148997fcb81b5582d32e057de0cd1d22d7016eb1af78d8fb757008499e1b538ff4d8c8f60001b610729565b50505080806148a790618e1a565b91505061426f565b506148dc7f0ef1637fec469998559e9f634e332630393bfa8713d7e349150ac1aaa088e07660001b610729565b6149087f81d45c19d54fdf67065eb6dffeec7517a621b04ea085329ad0551d6ca1ea8a5160001b610729565b60009150505b9392505050565b600061492030616433565b15905090565b6149527f50bcb00647133ad91079f0b6d4d3e3f099d5ce494625ee921bc4d8e3ddc5463060001b610729565b61497e7f8a1781bf86c073f4c0e1c69f6489e72f63601e04533642f3aa650d1d65b5cbbc60001b610729565b6149aa7ff312e84113018708a44a4de2ba25b79845e3181d8f782835b6760a4a5c40e11c60001b610729565b6149b48282616446565b6149e07f1e02fd8013672bdcd83ea52c7d3c8af8daa980b3d435affe2e753aa06af71b5a60001b610729565b614a0c7ffb6a8fb8ece1b5503a7cb9b5642a8de5936919236c3e78977f1f2772297091cb60001b610729565b614a1682826165cb565b5050565b6000614a487f457d6482736ed6851b82f9dd38ffc6077292ed011d12a8fb89a08cd9c69c7a2960001b610729565b614a747f7c46366fcc2b4d8fdc1ae589a88a4051c30f904985f69f4847cfdcca0a3fd9f160001b610729565b614aa07f7429f174463937961035cf1d40ef0f8a7a92aadb0b09a9dbe67105f7c4390a8760001b610729565b614aa86172f7565b614ad47f13120feb051b95d5e5c9dd7e40caac99f1b10e3bca054de71ea024746fa81f8c60001b610729565b614b007fad13d9b01132a71d811fce21e908a18ad384dbd668611f3fe5a29902534fa5c560001b610729565b6000614b2e7fa07fda3f9be7a553cf5b4669e4c0db35d29bd55792627df3bd9938168625f34160001b610729565b614b5a7fb40c1a81d204af649aa0c496b2b6580f0b8377e848440c6cbcd7bf7e76ca6c4060001b610729565b6000614b887f6ee5af59007cd5d3990495b24398bb0f36ee156edabf6b35ef4bb6f55180383960001b610729565b614bb47f464950c1ae9fd9aa75adb28a639656b07e66c61f6f150470f8997d9cb7b0e9f460001b610729565b6000614bf360405180604001604052808a73ffffffffffffffffffffffffffffffffffffffff1681526020018967ffffffffffffffff16815250610f7d565b9050614c217f6d3268463e7335e8c75e8fa943fdb8327188c645726bb11317cf1f7c5f536bce60001b610729565b614c4d7f6c839f8dc6c510ef295dc60ea69ef2ad4a022551303ede34fad46080a1c3bf1160001b610729565b6060614c7b7fb702be2a81441ead6a3093f45366bdc3cb24d3e0f7bd43a9a9742c2f62f2d9be60001b610729565b614ca77fb8f2c482015ed578a27b310292750a98e409565df741d4241d2bbc99c309261160001b610729565b8161012001519350614cdb7f10d2dd6e1a26f42a74af4fd64d020fd292b1d1ddbfb4eb348f3fe3a96363c89560001b610729565b614d077fbb53999dcb4db6dab6776e040c79980d29457cee1031e683866b557e14fefe4460001b610729565b60008467ffffffffffffffff161415614e8257614d467fccca4db22079fa069ca01888dbf38cc7fa7df84ebe29af4428aaf9cde713202f60001b610729565b614d727f35ddf38223ea09fc75c177d9e1da025d9244b0d83205de62a89c9d33759f5b2760001b610729565b614d9e7f55a77c6452f80b6764b7b5e4e751298762384c1451d03babb126ddcb79341c1760001b610729565b60019350614dce7fafdd98af31a117a31f840fddd69d4e6aca752ebc175bce844b620d8792c1b0d360001b610729565b614dfa7f5d622005848bcb207138753b81e104f976dd35fa25b2c8b222c34bf17a184f7360001b610729565b60019250614e2a7fac9b7193793e48398476cb01b24ccaeb10f65e8fde8ebf4331905f31da67076760001b610729565b614e567fb5b1aefca44ca5f6acfe2f01ddb368886bf9a479b709d99c7a971848cdc3649260001b610729565b60405180606001604052808567ffffffffffffffff1681526020018281526020018281525094506152aa565b614eae7fc3fc924a11509dbcd39917df4649733d605d21b82625d16405b4266ab8dc0e9160001b610729565b614eda7f09475e33f036c47faec29aee8c664ae0656e37ea6968bc943f0168092e47149860001b610729565b614f067fdcf8218e9ac6c0947efd7f8a3f3538e9895fecb008df20b146de5ecd6555b30160001b610729565b614f11898986615984565b9450614f3f7f3e1bba040a7db48a79f46aabfbe2fcf85b23cd1f8cbb66fb7227c14cc6c4f00260001b610729565b614f6b7f1e8a9d2b75b8882eddb8da9d4ac4fba5b736d9219892130988b4fc5354cc677f60001b610729565b600060036000876000015167ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156150ae5783829060005260206000209060020201604051806040016040529081600082018054614feb90618d6e565b80601f016020809104026020016040519081016040528092919081815260200182805461501790618d6e565b80156150645780601f1061503957610100808354040283529160200191615064565b820191906000526020600020905b81548152906001019060200180831161504757829003601f168201915b505050505081526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505081526020019060010190614fb8565b5050505090506150e07fc8377a1dcc8725c31198f223541431507db7e6abbd9ae548ec9a945f20c2067460001b610729565b61510c7fdab241cb54e5f3e2b75790a8427075e7aaf218f4ff9f3dfa0a4f916ddbf5fae260001b610729565b600060169054906101000a900467ffffffffffffffff1667ffffffffffffffff168151141561527b576151617f6669e193621796f6f062d1f552fc1107763536879da48ad5661bf8c688a6437460001b610729565b61518d7f5b73bacb9c7048ad4ab55e42ec78bdf7f8fa678d65ed506a02652e001da1305560001b610729565b848061519890618e1a565b9550506151c77fc429578b12d1adc129a36b810e1c460eda41ba9dc2c049d3612597322af9f5ef60001b610729565b6151f37feffb96e6752363e752b295827e822060aa2528977eca65e1dd86a895b3346af560001b610729565b600193506152237f19bf67571e113ed3026fb57da64f39b019f4459783d49d23f531917cc239105960001b610729565b61524f7f1cdd5420f7529eb80ce12d9cf6e09aa4f5c2bb90682d35e44d521f592e5c0e2a60001b610729565b60405180606001604052808667ffffffffffffffff1681526020018381526020018381525095506152a8565b6152a77fdba3aca52a0c298de70063e2bffb0f34bac5f031798a871851fbb123ab92adaa60001b610729565b5b505b6152d67fb6396f30d4e605c653fcd96fe2f75a8b762e264d191971238caa1d01d1e2f9da60001b610729565b6153027f5c7c527ecede565292ae235170e815641b7d0f1b194e8354be3a159596dc46ef60001b610729565b61530e8989878a616a30565b61533a7f23f538fedb78f37f74191f5292f618fd80136790a5a574e48beff0d3a4b1b8e760001b610729565b6153667fd915f1f65672be0747606fb22423626540a29123e0847809e04175538c80e5ba60001b610729565b82955050505050509392505050565b60006153a37fd9447aefcc5cfa296c1db4ca42e6efaeaf789a181cd3cc94dc6161be98d5eb0560001b610729565b6153cf7f68c16abbede51e1fc63470812ebb02f45e097dc029852df7d138f3c6cf7d405560001b610729565b6153fb7f964bfe3262d7d56e473f6c6d0f32789ccf19cd9ddbb47f20b122d6ba41dbc80760001b610729565b60006154297f5f78408f76d587a710e0fe84f07cb2ab27019799a56bcc2aeb233ef664ad71bd60001b610729565b6154557f7165b1ce2e3fe1d2c9572fc20a2ed212f5cf0c45db3c6fd70de44d3feac781df60001b610729565b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156157fa5783829060005260206000209060060201604051806101800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160109054906101000a900460ff166002811115615611577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115615649577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200160028201548152602001600382015481526020016004820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016004820160189054906101000a900460ff1615151515815260200160058201805480602002602001604051908101604052809291908181526020016000905b828210156157e357838290600052602060002001805461575690618d6e565b80601f016020809104026020016040519081016040528092919081815260200182805461578290618d6e565b80156157cf5780601f106157a4576101008083540402835291602001916157cf565b820191906000526020600020905b8154815290600101906020018083116157b257829003601f168201915b505050505081526020019060010190615737565b5050505081525050815260200190600101906154b6565b50505050905061582c7fcf7323efdd8553829776d98ebcd3cd4abc23f9eb63ba2ff8b6c0919175a0025a60001b610729565b6158587f8b88f20ba055da90bc8ad74164cbcd427dc317f1b382639c2007ebd9179326b460001b610729565b60005b8151811015615921576158907ff89062e5a73d8e8e24eeb8c05de30c4ffcf6f81c7dac14c9d16cb0b5248609d960001b610729565b6158bc7f11b21dbec2d05b8b72afc10dc0191a0616f0ae4ef5b9890f7c07837fb2862e0d60001b610729565b8181815181106158f5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151604001518361590c9190618b58565b9250808061591990618dd1565b91505061585b565b5061594e7f4333457829c5af5729c5e22e8eb1a4585ee4a250b2fd05b2cc0b3b32dc1af85c60001b610729565b61597a7fc072e8f009a9eff37c80d56b1d0cd46884bf1cd69c65201105be20a436396f7260001b610729565b8192505050919050565b61598c6172f7565b6159b87f28d69446bdfc85aaba8c73df75c45d1f59f9301b8a5a598b48476d045ff571a760001b610729565b6159e47f87c82cef7e55e73e488cfdf8cbba6c900138094ea50e6687ff6a1c87eef35b9160001b610729565b615a107f136518c0a28e070ddea0b1d80cd20f9b1ed92bb933a44264cccd6b4a80e66aab60001b610729565b6000848484604051602001615a27939291906186ee565b6040516020818303038152906040529050615a647fe9e5c26ec77ad82619040e2356609c7312e3b88c429807e36f42681ba5453b0d60001b610729565b615a907f43cf36297652f2c1bc45bda99e6dc10a138470905a4b308a42b9f6c35ae14c9860001b610729565b600281604051615aa0919061872b565b90815260200160405180910390206040518060600160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054615afb90618d6e565b80601f0160208091040260200160405190810160405280929190818152602001828054615b2790618d6e565b8015615b745780601f10615b4957610100808354040283529160200191615b74565b820191906000526020600020905b815481529060010190602001808311615b5757829003601f168201915b50505050508152602001600282018054615b8d90618d6e565b80601f0160208091040260200160405190810160405280929190818152602001828054615bb990618d6e565b8015615c065780601f10615bdb57610100808354040283529160200191615c06565b820191906000526020600020905b815481529060010190602001808311615be957829003601f168201915b5050505050815250509150509392505050565b6000615c477f2f90b92b80b3f22ef05dbaf3ec89141f1b7fbfc547ab97d4712fa7bff094a7c460001b610729565b615c737f7799a8def87e8f1536428191148b21d907d074b971d9f2af60df4e9174d52d4260001b610729565b615c9f7f9309f9b2c5b2611d5740a4f3e3a1808bfe8981a5e278fffa2739260c377b4cf160001b610729565b6000615cde60405180604001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018567ffffffffffffffff16815250610f7d565b9050615d0c7fda0332c4501ff34716b8ae74e1cb182fb7f6616fcc5b60e8eb3a434daf8c150c60001b610729565b615d387ff3a0f560cd9dff412c030a0a23b1cf32ac42ce3df44a6c79d9dd2f49e3e048bb60001b610729565b80610120015191505092915050565b600080615d767f38f0e2bcfac2aeb3348be2393ffb089cd8ebd21f46e6715cf314c409e93266c860001b610729565b615da27f4918d8d632291e6ac47d26c34a84884513a005411119db164fdf452f082b944960001b610729565b615dce7f48e297f96683fa9ad917277776d91c8985e2a92f351b2c87c19a9c0b1e53305160001b610729565b600060036000866000015167ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b82821015615f115783829060005260206000209060020201604051806040016040529081600082018054615e4e90618d6e565b80601f0160208091040260200160405190810160405280929190818152602001828054615e7a90618d6e565b8015615ec75780601f10615e9c57610100808354040283529160200191615ec7565b820191906000526020600020905b815481529060010190602001808311615eaa57829003601f168201915b505050505081526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505081526020019060010190615e1b565b505050509050615f437f2cb2693e4aebafd2b7dc12aa60303b52cd07546043000f74490592abdab2231e60001b610729565b615f6f7f0fc5a9e9d7092fa429a1506caffd1ea42790aee2efd7fa4b67c7ad17998d1f1a60001b610729565b60008151141561600a57615fa57f7fcc4f9cc58d5cd9a3bbbc7168ced1699a9eab33e7afd0217147ded9e793f45960001b610729565b615fd17f1862580d4d56e0c45d0b5aee7912916619459c512b03afeec5b4daff7ded546e60001b610729565b615ffd7f87eb45878493323932251eef6466b0bb4435370b50d3b1ed8ad0bced5a2b9b3a60001b610729565b6000809250925050616293565b6160367f32bd96b518d3469418db08fc0f7b94c129a24450fa857021054659dc0e74a90e60001b610729565b6160627f7c408a12ca6d3ff9197b5483324c5a2ade568da50489555267505edd600785d660001b610729565b61608e7f631d642abc00356c5d34923a07f566c95c141df6b7859da8f9c4c1b14e7df55360001b610729565b60005b81518167ffffffffffffffff161015616231576160d07f5ea0bb6edcb01f2d61457313edaf6335a087a18512b916e9748c046c6415916160001b610729565b6160fc7fe78977092a0703b876a9f92f0565d97523149c5a7dec86fc087f3eece964e40160001b610729565b8480519060200120828267ffffffffffffffff1681518110616147577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151600001518051906020012014156161f25761618c7f4374d2e61714182d53fa1c8a7d95db5f74ebe27dc3b856b5a965bac602727e0260001b610729565b6161b87f54508d9b029428a14b78723ee414a99d8db24103c081888a472efcad20d8ff2060001b610729565b6161e47fb1bdfba3cf2beb774b81fba74c66d4ba1c7fda2317f3357be2bdd4140e97e72d60001b610729565b806001935093505050616293565b61621e7fff8984b7603a4351a5ca4b7f007cbb4b127a35e4e9583813abd72f170e5af40560001b610729565b808061622990618e1a565b915050616091565b5061625e7fa9a145d1a719ecf4e7cab56893c66cfa0af2bf79c953d4bec01edb7f7aa8e0b360001b610729565b61628a7f759b808197974fea3daa4c9afe8b59ad718ee40a01bb4267756c96d0ab62174960001b610729565b60008092509250505b9250929050565b6162c67f3f019eba0a6a282b9b04a436d31feb8e3bf2a0923d1d2d608c234038ed870e7d60001b610729565b6162f27ffa7624b5757643bcbfd269b4e43dd4b98f070a2023f9ce517e60f46a2a225e6c60001b610729565b61631e7fd416837649abc04f9b45ea9d2b4a884907b827df05ec23dcb25c7c26f0641e3460001b610729565b600083838360000151604051602001616339939291906186ee565b60405160208183030381529060405290506163767ffaa060d3048f3d8b8ea2917d49bc09a4dc94b3a0ae5be9b0a84ce9c1c324e71f60001b610729565b6163a27f24bb84df2c2f8d2ba0e44a8dc4fba0a410e0a810d250de792b4ef0798b6840b660001b610729565b816002826040516163b3919061872b565b908152602001604051809103902060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550602082015181600101908051906020019061640c929190617322565b506040820151816002019080519060200190616429929190617322565b5090505050505050565b600080823b905060008111915050919050565b6164727f0611ef49ce715b7fae002feea5249a4ddaf9058e0ed0957d78b864138e2ed4b560001b610729565b61649e7f476221235fccddd81f8c1501361f0a1d9dd094314023bebbdfd0672e14bdf81460001b610729565b6164ca7f7769a79402d40a8e1d6f100e0c40eab27abf40e07f588b12f484d704972932aa60001b610729565b60006164d68383615c19565b90506165047f70cb0cb80b423a5140ac7e80d296d80367fc261b786a24d08165ae023322bbf760001b610729565b6165307f1f9dab3ecb36e321bbe3cc2e18de9a8fd507ab8566a802d8c41f7cbdac6a140b60001b610729565b60005b8167ffffffffffffffff168167ffffffffffffffff1610156165c55761657b7f4fcd43470fa02e8bc771d34d0cdaa02cbb6ab6f6738a3476fa1d67f898fac2f560001b610729565b6165a77fd2cf4bf803336f4a4e5d3bb21a2f8d92b0b49289f80ce80acba1be4125fa5f9160001b610729565b6165b2848483616cc5565b80806165bd90618e1a565b915050616533565b50505050565b6165f77faaee39ab9bfaa79f1d9af5a070fa3850f38de076e5b17447042331664bf9c20660001b610729565b6166237ffd33ab2e18f68164515b524444c029bcc2e67ee8dd1bbdbffb9a7e701a5d97fc60001b610729565b61664f7f3d0852de2ff0d2cb91481bc522dbc77ed7f2f12c2bd25c6d5ed933d8c743810460001b610729565b60005b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490508167ffffffffffffffff161015616a2b576166d27f5e9713a08249afa1c07ef58fe07bb63064363102cba89a9ef0c3de53c4c5930f60001b610729565b6166fe7f0c30df1c813fd84bba940fa6ba4ec4716b9f4fef1ba763b89b759c6dae0f603160001b610729565b8167ffffffffffffffff16600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208267ffffffffffffffff168154811061678a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906006020160000160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1614156169ec576167ec7f84b75cb0141704832f341df1f9cf9cc5662426815d152e3d2983ea7c7496614260001b610729565b6168187f3c35f6f221c8577a867b2c25e5ef966a326cc03dbd6a0689cf58302469101d2b60001b610729565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208167ffffffffffffffff1681548110616899577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000209060060201600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549067ffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff02191690556001820160086101000a81549067ffffffffffffffff02191690556001820160106101000a81549060ff0219169055600282016000905560038201600090556004820160006101000a81549067ffffffffffffffff02191690556004820160086101000a81549067ffffffffffffffff02191690556004820160106101000a81549067ffffffffffffffff02191690556004820160186101000a81549060ff02191690556005820160006169b991906173a8565b50506169e77fdc3cb7ae1670a712523e2fc4e21fe2c4752c96f1f2acbf4690e6fd08053278a360001b610729565b616a2b565b616a187f04e576e7202f8f48222ee79bde432448b547d07e89d32cd7d35846e5fc6add6060001b610729565b8080616a2390618e1a565b915050616652565b505050565b616a5c7f389a0f5328209e3295ff5c2c37dcd22b41171fe3ed9e2642bb971291e7462c8f60001b610729565b616a887f9928b1cd7a7947e2c71c25a93ad06d56b25be91447f3a29cb1b7365a06efc49d60001b610729565b616ab47ffbcb1d30ab84655d8eed8e2c3b19cb68fa05b98b8c0593359d05368b483577a160001b610729565b600084848460000151604051602001616acf939291906186ee565b6040516020818303038152906040529050616b0c7f071aa4f44e402d32ab01d68708f5a8a7b29048eaf325483b604a9f2190fa12ae60001b610729565b616b387f8c5d4d881d374418037664bfd7e34625d89b21f2e8b974b44ffb7e2274bd5dfb60001b610729565b82600282604051616b49919061872b565b908152602001604051809103902060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506020820151816001019080519060200190616ba2929190617322565b506040820151816002019080519060200190616bbf929190617322565b50905050616bef7fddfed950469df18443916696eed5585837dc4295da26a6628964a8dad07634e860001b610729565b616c1b7fb2e1677842f101318da10939f713c79e94826af73535477bdd41ab57f0c9ecf660001b610729565b60036000846000015167ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000208290806001815401808255809150506001900390600052602060002090600202016000909190919091506000820151816000019080519060200190616c8c929190617322565b5060208201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050505050505050565b616cf17f9668b27ffb0894d03d592ee85097a9d85999df04dae884c00c8695ff89ba635560001b610729565b616d1d7f0f2366730038f71599292bdfdc23a12b9e46e1d6659d91ad7cc71df96783e9d360001b610729565b616d497f9cd6cc6108380c61764b02bbbcc3b2814a02b95f2fb3b6214a658f6de9c07b8360001b610729565b6000838383604051602001616d60939291906186ee565b6040516020818303038152906040529050616d9d7f23e4f2944cb125a6ecc0e78725a8a46a5c4359012e08c18171b6307f08be1df060001b610729565b600281604051616dad919061872b565b9081526020016040518091039020600080820160006101000a81549067ffffffffffffffff0219169055600182016000616de791906172b7565b600282016000616df791906172b7565b5050616e257fbed2d35d37dd4f6490f1b939f8a5a51759af9d128bc917c68e744c01bb171bae60001b610729565b600360008367ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206000616e5891906173c9565b50505050565b828054828255906000526020600020908101928215616ead579160200282015b82811115616eac578251829080519060200190616e9c929190617322565b5091602001919060010190616e7e565b5b509050616eba91906173ed565b5090565b8280548282559060005260206000209060060281019282156171b95760005260206000209160060282015b828111156171b85782826000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000820160149054906101000a900467ffffffffffffffff168160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160009054906101000a900467ffffffffffffffff168160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160089054906101000a900467ffffffffffffffff168160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160109054906101000a900460ff168160010160106101000a81548160ff02191690836002811115617081577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060028201548160020155600382015481600301556004820160009054906101000a900467ffffffffffffffff168160040160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506004820160089054906101000a900467ffffffffffffffff168160040160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506004820160109054906101000a900467ffffffffffffffff168160040160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506004820160189054906101000a900460ff168160040160186101000a81548160ff02191690831515021790555060058201816005019080546171a6929190617411565b50505091600601919060060190616ee9565b5b5090506171c69190617479565b5090565b604051806101800160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160006002811115617260577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016000815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600015158152602001606081525090565b5080546172c390618d6e565b6000825580601f106172d557506172f4565b601f0160209004906000526020600020908101906172f391906175a1565b5b50565b6040518060600160405280600067ffffffffffffffff16815260200160608152602001606081525090565b82805461732e90618d6e565b90600052602060002090601f0160209004810192826173505760008555617397565b82601f1061736957805160ff1916838001178555617397565b82800160010185558215617397579182015b8281111561739657825182559160200191906001019061737b565b5b5090506173a491906175a1565b5090565b50805460008255906000526020600020908101906173c691906173ed565b50565b50805460008255600202906000526020600020908101906173ea91906175be565b50565b5b8082111561740d576000818161740491906172b7565b506001016173ee565b5090565b8280548282559060005260206000209081019282156174685760005260206000209182015b8281111561746757828290805461744c90618d6e565b617457929190617600565b5091600101919060010190617436565b5b50905061747591906173ed565b5090565b5b8082111561759d57600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549067ffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff02191690556001820160086101000a81549067ffffffffffffffff02191690556001820160106101000a81549060ff0219169055600282016000905560038201600090556004820160006101000a81549067ffffffffffffffff02191690556004820160086101000a81549067ffffffffffffffff02191690556004820160106101000a81549067ffffffffffffffff02191690556004820160186101000a81549060ff021916905560058201600061759491906173a8565b5060060161747a565b5090565b5b808211156175ba5760008160009055506001016175a2565b5090565b5b808211156175fc57600080820160006175d891906172b7565b6001820160006101000a81549067ffffffffffffffff0219169055506002016175bf565b5090565b82805461760c90618d6e565b90600052602060002090601f01602090048101928261762e576000855561767c565b82601f1061763f578054855561767c565b8280016001018555821561767c57600052602060002091601f016020900482015b8281111561767b578254825591600101919060010190617660565b5b50905061768991906175a1565b5090565b60006176a061769b84618a1a565b6189f5565b905080838252602082019050828560208602820111156176bf57600080fd5b60005b858110156176ef57816176d588826177fb565b8452602084019350602083019250506001810190506176c2565b5050509392505050565b600061770c61770784618a46565b6189f5565b9050808382526020820190508285602086028201111561772b57600080fd5b60005b8581101561777557813567ffffffffffffffff81111561774d57600080fd5b80860161775a89826178b8565b8552602085019450602084019350505060018101905061772e565b5050509392505050565b600061779261778d84618a72565b6189f5565b9050828152602081018484840111156177aa57600080fd5b6177b5848285618d02565b509392505050565b60006177d06177cb84618a72565b6189f5565b9050828152602081018484840111156177e857600080fd5b6177f3848285618d11565b509392505050565b60008135905061780a81619179565b92915050565b60008151905061781f81619179565b92915050565b600082601f83011261783657600080fd5b813561784684826020860161768d565b91505092915050565b600082601f83011261786057600080fd5b81356178708482602086016176f9565b91505092915050565b60008135905061788881619190565b92915050565b60008151905061789d81619190565b92915050565b6000813590506178b2816191a7565b92915050565b600082601f8301126178c957600080fd5b81356178d984826020860161777f565b91505092915050565b600082601f8301126178f357600080fd5b81516179038482602086016177bd565b91505092915050565b60008135905061791b816191be565b92915050565b600081359050617930816191d5565b92915050565b600081359050617945816191e5565b92915050565b6000610320828403121561795e57600080fd5b6179696102e06189f5565b9050600082013567ffffffffffffffff81111561798557600080fd5b617991848285016178b8565b60008301525060206179a5848285016177fb565b602083015250604082013567ffffffffffffffff8111156179c557600080fd5b6179d1848285016178b8565b60408301525060606179e584828501617ee2565b60608301525060806179f984828501617ee2565b60808301525060a0617a0d84828501617ee2565b60a08301525060c0617a2184828501617ee2565b60c08301525060e0617a3584828501617ee2565b60e083015250610100617a4a84828501617ecd565b61010083015250610120617a6084828501617ee2565b61012083015250610140617a7684828501617ee2565b6101408301525061016082013567ffffffffffffffff811115617a9857600080fd5b617aa4848285016178b8565b61016083015250610180617aba84828501617ee2565b610180830152506101a0617ad084828501617ecd565b6101a0830152506101c0617ae684828501617879565b6101c0830152506101e0617afc84828501617936565b6101e083015250610200617b1284828501617ee2565b6102008301525061022082013567ffffffffffffffff811115617b3457600080fd5b617b4084828501617825565b6102208301525061024082013567ffffffffffffffff811115617b6257600080fd5b617b6e84828501617825565b6102408301525061026082013567ffffffffffffffff811115617b9057600080fd5b617b9c848285016178b8565b61026083015250610280617bb284828501617921565b610280830152506102a0617bc884828501617879565b6102a0830152506102c0617bde84828501617cb3565b6102c08301525092915050565b600060e08284031215617bfd57600080fd5b617c0760e06189f5565b90506000617c1784828501617ef7565b6000830152506020617c2b84828501617ef7565b6020830152506040617c3f84828501617ef7565b6040830152506060617c5384828501617ef7565b6060830152506080617c6784828501617ef7565b60808301525060a0617c7b84828501617810565b60a08301525060c082015167ffffffffffffffff811115617c9b57600080fd5b617ca7848285016178e2565b60c08301525092915050565b600060608284031215617cc557600080fd5b617ccf60606189f5565b90506000617cdf84828501617ee2565b6000830152506020617cf384828501617ee2565b6020830152506040617d0784828501617ee2565b60408301525092915050565b600060208284031215617d2557600080fd5b617d2f60206189f5565b90506000617d3f84828501617ee2565b60008301525092915050565b60006101808284031215617d5e57600080fd5b617d696101806189f5565b90506000617d79848285016177fb565b6000830152506020617d8d84828501617ee2565b6020830152506040617da184828501617ee2565b6040830152506060617db584828501617ee2565b6060830152506080617dc984828501617921565b60808301525060a0617ddd84828501617ecd565b60a08301525060c0617df184828501617ecd565b60c08301525060e0617e0584828501617ee2565b60e083015250610100617e1a84828501617ee2565b61010083015250610120617e3084828501617ee2565b61012083015250610140617e4684828501617879565b6101408301525061016082013567ffffffffffffffff811115617e6857600080fd5b617e748482850161784f565b6101608301525092915050565b600060408284031215617e9357600080fd5b617e9d60406189f5565b90506000617ead848285016177fb565b6000830152506020617ec184828501617ee2565b60208301525092915050565b600081359050617edc816191f5565b92915050565b600081359050617ef18161920c565b92915050565b600081519050617f068161920c565b92915050565b600060208284031215617f1e57600080fd5b6000617f2c848285016177fb565b91505092915050565b60008060408385031215617f4857600080fd5b6000617f56858286016177fb565b9250506020617f6785828601617ee2565b9150509250929050565b600060208284031215617f8357600080fd5b6000617f918482850161788e565b91505092915050565b600060208284031215617fac57600080fd5b6000617fba848285016178a3565b91505092915050565b60008060408385031215617fd657600080fd5b6000617fe48582860161790c565b9250506020617ff585828601617d13565b9150509250929050565b60006020828403121561801157600080fd5b600082015167ffffffffffffffff81111561802b57600080fd5b61803784828501617beb565b91505092915050565b60006020828403121561805257600080fd5b600082013567ffffffffffffffff81111561806c57600080fd5b61807884828501617d4b565b91505092915050565b6000806040838503121561809457600080fd5b600083013567ffffffffffffffff8111156180ae57600080fd5b6180ba85828601617d4b565b925050602083013567ffffffffffffffff8111156180d757600080fd5b6180e38582860161794b565b9150509250929050565b6000604082840312156180ff57600080fd5b600061810d84828501617e81565b91505092915050565b6000618122838361827b565b905092915050565b60006181368383618493565b905092915050565b61814781618c40565b82525050565b61815681618c40565b82525050565b61816d61816882618c40565b618e4b565b82525050565b600061817e82618ac3565b6181888185618b09565b93508360208202850161819a85618aa3565b8060005b858110156181d657848403895281516181b78582618116565b94506181c283618aef565b925060208a0199505060018101905061819e565b50829750879550505050505092915050565b60006181f382618ace565b6181fd8185618b1a565b93508360208202850161820f85618ab3565b8060005b8581101561824b578484038952815161822c858261812a565b945061823783618afc565b925060208a01995050600181019050618213565b50829750879550505050505092915050565b61826681618c52565b82525050565b61827581618c52565b82525050565b600061828682618ad9565b6182908185618b2b565b93506182a0818560208601618d11565b6182a981618f3d565b840191505092915050565b6182bd81618cde565b82525050565b6182cc81618cf0565b82525050565b6182db81618cf0565b82525050565b60006182ec82618ae4565b6182f68185618b4d565b9350618306818560208601618d11565b80840191505092915050565b600061831f601383618b3c565b915061832a82618f68565b602082019050919050565b6000618342600c83618b3c565b915061834d82618f91565b602082019050919050565b6000618365600d83618b3c565b915061837082618fba565b602082019050919050565b6000618388600f83618b3c565b915061839382618fe3565b602082019050919050565b60006183ab600c83618b3c565b91506183b68261900c565b602082019050919050565b60006183ce602e83618b3c565b91506183d982619035565b604082019050919050565b60006183f1600e83618b3c565b91506183fc82619084565b602082019050919050565b6000618414601383618b3c565b915061841f826190ad565b602082019050919050565b6000618437600e83618b3c565b9150618442826190d6565b602082019050919050565b600061845a601583618b3c565b9150618465826190ff565b602082019050919050565b600061847d600983618b3c565b915061848882619128565b602082019050919050565b6000610180830160008301516184ac600086018261813e565b5060208301516184bf60208601826186b9565b5060408301516184d260408601826186b9565b5060608301516184e560608601826186b9565b5060808301516184f860808601826182c3565b5060a083015161850b60a086018261869b565b5060c083015161851e60c086018261869b565b5060e083015161853160e08601826186b9565b506101008301516185466101008601826186b9565b5061012083015161855b6101208601826186b9565b5061014083015161857061014086018261825d565b5061016083015184820361016086015261858a8282618173565b9150508091505092915050565b6000610180830160008301516185b0600086018261813e565b5060208301516185c360208601826186b9565b5060408301516185d660408601826186b9565b5060608301516185e960608601826186b9565b5060808301516185fc60808601826182c3565b5060a083015161860f60a086018261869b565b5060c083015161862260c086018261869b565b5060e083015161863560e08601826186b9565b5061010083015161864a6101008601826186b9565b5061012083015161865f6101208601826186b9565b5061014083015161867461014086018261825d565b5061016083015184820361016086015261868e8282618173565b9150508091505092915050565b6186a481618cc0565b82525050565b6186b381618cc0565b82525050565b6186c281618cca565b82525050565b6186d181618cca565b82525050565b6186e86186e382618cca565b618e6f565b82525050565b60006186fa828661815c565b60148201915061870a82856186d7565b60088201915061871a82846186d7565b600882019150819050949350505050565b600061873782846182e1565b915081905092915050565b6000602082019050618757600083018461814d565b92915050565b6000602082019050818103600083015261877781846181e8565b905092915050565b6000602082019050618794600083018461826c565b92915050565b60006080820190506187af60008301876182b4565b6187bc60208301866186aa565b6187c9604083018561814d565b6187d660608301846186c8565b95945050505050565b600060e0820190506187f4600083018a6182b4565b61880160208301896186aa565b61880e604083018861814d565b61881b60608301876186c8565b61882860808301866182d2565b61883560a08301856186c8565b61884260c083018461826c565b98975050505050505050565b6000604082019050818103600083015261886781618335565b9050818103602083015261887a81618312565b9050919050565b6000604082019050818103600083015261889a81618335565b905081810360208301526188ad81618358565b9050919050565b600060408201905081810360008301526188cd81618335565b905081810360208301526188e081618407565b9050919050565b6000604082019050818103600083015261890081618335565b905081810360208301526189138161844d565b9050919050565b6000604082019050818103600083015261893381618335565b9050818103602083015261894681618470565b9050919050565b600060408201905081810360008301526189668161837b565b90508181036020830152618979816183e4565b9050919050565b600060408201905081810360008301526189998161839e565b905081810360208301526189ac8161842a565b9050919050565b600060208201905081810360008301526189cc816183c1565b9050919050565b600060208201905081810360008301526189ed8184618597565b905092915050565b60006189ff618a10565b9050618a0b8282618da0565b919050565b6000604051905090565b600067ffffffffffffffff821115618a3557618a34618f0e565b5b602082029050602081019050919050565b600067ffffffffffffffff821115618a6157618a60618f0e565b5b602082029050602081019050919050565b600067ffffffffffffffff821115618a8d57618a8c618f0e565b5b618a9682618f3d565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000618b6382618cca565b9150618b6e83618cca565b92508267ffffffffffffffff03821115618b8b57618b8a618e81565b5b828201905092915050565b6000618ba182618cca565b9150618bac83618cca565b92508167ffffffffffffffff0483118215151615618bcd57618bcc618e81565b5b828202905092915050565b6000618be382618cc0565b9150618bee83618cc0565b925082821015618c0157618c00618e81565b5b828203905092915050565b6000618c1782618cca565b9150618c2283618cca565b925082821015618c3557618c34618e81565b5b828203905092915050565b6000618c4b82618ca0565b9050919050565b60008115159050919050565b6000819050919050565b6000618c7382618c40565b9050919050565b6000819050618c8882619151565b919050565b6000819050618c9b82619165565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b6000618ce982618c7a565b9050919050565b6000618cfb82618c8d565b9050919050565b82818337600083830152505050565b60005b83811015618d2f578082015181840152602081019050618d14565b83811115618d3e576000848401525b50505050565b6000618d4f82618cca565b91506000821415618d6357618d62618e81565b5b600182039050919050565b60006002820490506001821680618d8657607f821691505b60208210811415618d9a57618d99618edf565b5b50919050565b618da982618f3d565b810181811067ffffffffffffffff82111715618dc857618dc7618f0e565b5b80604052505050565b6000618ddc82618cc0565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415618e0f57618e0e618e81565b5b600182019050919050565b6000618e2582618cca565b915067ffffffffffffffff821415618e4057618e3f618e81565b5b600182019050919050565b6000618e5682618e5d565b9050919050565b6000618e6882618f5b565b9050919050565b6000618e7a82618f4e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160c01b9050919050565b60008160601b9050919050565b7f4e6f6465206e6f74207265676973746572656400000000000000000000000000600082015250565b7f437265617465536563746f720000000000000000000000000000000000000000600082015250565b7f536563746f724944206973203000000000000000000000000000000000000000600082015250565b7f41646446696c65546f536563746f720000000000000000000000000000000000600082015250565b7f44656c657465536563746f720000000000000000000000000000000000000000600082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f4e6f74456e6f7567685370616365000000000000000000000000000000000000600082015250565b7f496e73756666696369656e7420766f6c756d6500000000000000000000000000600082015250565b7f4e6f74456d707479536563746f72000000000000000000000000000000000000600082015250565b7f536563746f7220616c7265616479206578697374730000000000000000000000600082015250565b7f53697a6520697320300000000000000000000000000000000000000000000000600082015250565b600b811061916257619161618eb0565b5b50565b6003811061917657619175618eb0565b5b50565b61918281618c40565b811461918d57600080fd5b50565b61919981618c52565b81146191a457600080fd5b50565b6191b081618c5e565b81146191bb57600080fd5b50565b6191c781618c68565b81146191d257600080fd5b50565b600381106191e257600080fd5b50565b600281106191f257600080fd5b50565b6191fe81618cc0565b811461920957600080fd5b50565b61921581618cca565b811461922057600080fd5b5056fea26469706673582212202a4026450e3bb69af018f3537263975ff8e9526814e693c53a4f2a99637a594664736f6c63430008040033",
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

// GetSectorInfo is a free data retrieval call binding the contract method 0x2ba010d7.
//
// Solidity: function GetSectorInfo((address,uint64) sectorRef) view returns((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]))
func (_Store *StoreCaller) GetSectorInfo(opts *bind.CallOpts, sectorRef SectorRef) (SectorInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetSectorInfo", sectorRef)

	if err != nil {
		return *new(SectorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(SectorInfo)).(*SectorInfo)

	return out0, err

}

// GetSectorInfo is a free data retrieval call binding the contract method 0x2ba010d7.
//
// Solidity: function GetSectorInfo((address,uint64) sectorRef) view returns((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]))
func (_Store *StoreSession) GetSectorInfo(sectorRef SectorRef) (SectorInfo, error) {
	return _Store.Contract.GetSectorInfo(&_Store.CallOpts, sectorRef)
}

// GetSectorInfo is a free data retrieval call binding the contract method 0x2ba010d7.
//
// Solidity: function GetSectorInfo((address,uint64) sectorRef) view returns((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]))
func (_Store *StoreCallerSession) GetSectorInfo(sectorRef SectorRef) (SectorInfo, error) {
	return _Store.Contract.GetSectorInfo(&_Store.CallOpts, sectorRef)
}

// GetSectorsForNode is a free data retrieval call binding the contract method 0xe3168f9e.
//
// Solidity: function GetSectorsForNode(address nodeAddr) view returns((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[])[])
func (_Store *StoreCaller) GetSectorsForNode(opts *bind.CallOpts, nodeAddr common.Address) ([]SectorInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetSectorsForNode", nodeAddr)

	if err != nil {
		return *new([]SectorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]SectorInfo)).(*[]SectorInfo)

	return out0, err

}

// GetSectorsForNode is a free data retrieval call binding the contract method 0xe3168f9e.
//
// Solidity: function GetSectorsForNode(address nodeAddr) view returns((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[])[])
func (_Store *StoreSession) GetSectorsForNode(nodeAddr common.Address) ([]SectorInfo, error) {
	return _Store.Contract.GetSectorsForNode(&_Store.CallOpts, nodeAddr)
}

// GetSectorsForNode is a free data retrieval call binding the contract method 0xe3168f9e.
//
// Solidity: function GetSectorsForNode(address nodeAddr) view returns((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[])[])
func (_Store *StoreCallerSession) GetSectorsForNode(nodeAddr common.Address) ([]SectorInfo, error) {
	return _Store.Contract.GetSectorsForNode(&_Store.CallOpts, nodeAddr)
}

// IsSectorRefByFileInfo is a free data retrieval call binding the contract method 0x9a7d0a28.
//
// Solidity: function IsSectorRefByFileInfo(address nodeAddr, uint64 sectorID) view returns(bool)
func (_Store *StoreCaller) IsSectorRefByFileInfo(opts *bind.CallOpts, nodeAddr common.Address, sectorID uint64) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "IsSectorRefByFileInfo", nodeAddr, sectorID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSectorRefByFileInfo is a free data retrieval call binding the contract method 0x9a7d0a28.
//
// Solidity: function IsSectorRefByFileInfo(address nodeAddr, uint64 sectorID) view returns(bool)
func (_Store *StoreSession) IsSectorRefByFileInfo(nodeAddr common.Address, sectorID uint64) (bool, error) {
	return _Store.Contract.IsSectorRefByFileInfo(&_Store.CallOpts, nodeAddr, sectorID)
}

// IsSectorRefByFileInfo is a free data retrieval call binding the contract method 0x9a7d0a28.
//
// Solidity: function IsSectorRefByFileInfo(address nodeAddr, uint64 sectorID) view returns(bool)
func (_Store *StoreCallerSession) IsSectorRefByFileInfo(nodeAddr common.Address, sectorID uint64) (bool, error) {
	return _Store.Contract.IsSectorRefByFileInfo(&_Store.CallOpts, nodeAddr, sectorID)
}

// C0x29d6276d is a free data retrieval call binding the contract method 0x0daf9945.
//
// Solidity: function c_0x29d6276d(bytes32 c__0x29d6276d) pure returns()
func (_Store *StoreCaller) C0x29d6276d(opts *bind.CallOpts, c__0x29d6276d [32]byte) error {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "c_0x29d6276d", c__0x29d6276d)

	if err != nil {
		return err
	}

	return err

}

// C0x29d6276d is a free data retrieval call binding the contract method 0x0daf9945.
//
// Solidity: function c_0x29d6276d(bytes32 c__0x29d6276d) pure returns()
func (_Store *StoreSession) C0x29d6276d(c__0x29d6276d [32]byte) error {
	return _Store.Contract.C0x29d6276d(&_Store.CallOpts, c__0x29d6276d)
}

// C0x29d6276d is a free data retrieval call binding the contract method 0x0daf9945.
//
// Solidity: function c_0x29d6276d(bytes32 c__0x29d6276d) pure returns()
func (_Store *StoreCallerSession) C0x29d6276d(c__0x29d6276d [32]byte) error {
	return _Store.Contract.C0x29d6276d(&_Store.CallOpts, c__0x29d6276d)
}

// AddFileToSector is a paid mutator transaction binding the contract method 0x955f98b7.
//
// Solidity: function AddFileToSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreTransactor) AddFileToSector(opts *bind.TransactOpts, sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "AddFileToSector", sectorInfo, fileInfo)
}

// AddFileToSector is a paid mutator transaction binding the contract method 0x955f98b7.
//
// Solidity: function AddFileToSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreSession) AddFileToSector(sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.AddFileToSector(&_Store.TransactOpts, sectorInfo, fileInfo)
}

// AddFileToSector is a paid mutator transaction binding the contract method 0x955f98b7.
//
// Solidity: function AddFileToSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreTransactorSession) AddFileToSector(sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.AddFileToSector(&_Store.TransactOpts, sectorInfo, fileInfo)
}

// CreateSector is a paid mutator transaction binding the contract method 0xba921004.
//
// Solidity: function CreateSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) returns()
func (_Store *StoreTransactor) CreateSector(opts *bind.TransactOpts, sectorInfo SectorInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "CreateSector", sectorInfo)
}

// CreateSector is a paid mutator transaction binding the contract method 0xba921004.
//
// Solidity: function CreateSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) returns()
func (_Store *StoreSession) CreateSector(sectorInfo SectorInfo) (*types.Transaction, error) {
	return _Store.Contract.CreateSector(&_Store.TransactOpts, sectorInfo)
}

// CreateSector is a paid mutator transaction binding the contract method 0xba921004.
//
// Solidity: function CreateSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) returns()
func (_Store *StoreTransactorSession) CreateSector(sectorInfo SectorInfo) (*types.Transaction, error) {
	return _Store.Contract.CreateSector(&_Store.TransactOpts, sectorInfo)
}

// DeleteFileFromSector is a paid mutator transaction binding the contract method 0x0047c003.
//
// Solidity: function DeleteFileFromSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreTransactor) DeleteFileFromSector(opts *bind.TransactOpts, sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteFileFromSector", sectorInfo, fileInfo)
}

// DeleteFileFromSector is a paid mutator transaction binding the contract method 0x0047c003.
//
// Solidity: function DeleteFileFromSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreSession) DeleteFileFromSector(sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.DeleteFileFromSector(&_Store.TransactOpts, sectorInfo, fileInfo)
}

// DeleteFileFromSector is a paid mutator transaction binding the contract method 0x0047c003.
//
// Solidity: function DeleteFileFromSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreTransactorSession) DeleteFileFromSector(sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.DeleteFileFromSector(&_Store.TransactOpts, sectorInfo, fileInfo)
}

// DeleteSector is a paid mutator transaction binding the contract method 0x931bb19a.
//
// Solidity: function DeleteSector((address,uint64) sectorRef) returns()
func (_Store *StoreTransactor) DeleteSector(opts *bind.TransactOpts, sectorRef SectorRef) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteSector", sectorRef)
}

// DeleteSector is a paid mutator transaction binding the contract method 0x931bb19a.
//
// Solidity: function DeleteSector((address,uint64) sectorRef) returns()
func (_Store *StoreSession) DeleteSector(sectorRef SectorRef) (*types.Transaction, error) {
	return _Store.Contract.DeleteSector(&_Store.TransactOpts, sectorRef)
}

// DeleteSector is a paid mutator transaction binding the contract method 0x931bb19a.
//
// Solidity: function DeleteSector((address,uint64) sectorRef) returns()
func (_Store *StoreTransactorSession) DeleteSector(sectorRef SectorRef) (*types.Transaction, error) {
	return _Store.Contract.DeleteSector(&_Store.TransactOpts, sectorRef)
}

// UpdateSectorInfo is a paid mutator transaction binding the contract method 0x2384a6aa.
//
// Solidity: function UpdateSectorInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sector) payable returns()
func (_Store *StoreTransactor) UpdateSectorInfo(opts *bind.TransactOpts, sector SectorInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateSectorInfo", sector)
}

// UpdateSectorInfo is a paid mutator transaction binding the contract method 0x2384a6aa.
//
// Solidity: function UpdateSectorInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sector) payable returns()
func (_Store *StoreSession) UpdateSectorInfo(sector SectorInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateSectorInfo(&_Store.TransactOpts, sector)
}

// UpdateSectorInfo is a paid mutator transaction binding the contract method 0x2384a6aa.
//
// Solidity: function UpdateSectorInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sector) payable returns()
func (_Store *StoreTransactorSession) UpdateSectorInfo(sector SectorInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateSectorInfo(&_Store.TransactOpts, sector)
}

// Initialize is a paid mutator transaction binding the contract method 0x112346c2.
//
// Solidity: function initialize(address _node, (uint64) sectorConfig) returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts, _node common.Address, sectorConfig SectorConfig) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize", _node, sectorConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x112346c2.
//
// Solidity: function initialize(address _node, (uint64) sectorConfig) returns()
func (_Store *StoreSession) Initialize(_node common.Address, sectorConfig SectorConfig) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _node, sectorConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x112346c2.
//
// Solidity: function initialize(address _node, (uint64) sectorConfig) returns()
func (_Store *StoreTransactorSession) Initialize(_node common.Address, sectorConfig SectorConfig) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _node, sectorConfig)
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
