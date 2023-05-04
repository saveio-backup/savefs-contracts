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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"port\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"}],\"name\":\"DNSNodeRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DNSNodeUnReg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"DnsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"FsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structNameInfo\",\"name\":\"newer\",\"type\":\"tuple\"}],\"name\":\"NotifyNameInfoAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"PDPVerifyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nodeAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"RootHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"FileID\",\"type\":\"bytes\"}],\"internalType\":\"structFileProveParam\",\"name\":\"FileProveParam_\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"AddFileToSector\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"}],\"name\":\"CreateSector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"RootHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"FileID\",\"type\":\"bytes\"}],\"internalType\":\"structFileProveParam\",\"name\":\"FileProveParam_\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"DeleteFileFromSector\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"DeleteSector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"GetSectorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"GetSectorsForNode\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sectorID\",\"type\":\"uint64\"}],\"name\":\"IsSectorRefByFileInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sector\",\"type\":\"tuple\"}],\"name\":\"UpdateSectorInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINode\",\"name\":\"_node\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"SECTOR_FILE_INFO_GROUP_MAX_LEN\",\"type\":\"uint64\"}],\"internalType\":\"structSectorConfig\",\"name\":\"sectorConfig\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50613f0c806100206000396000f3fe6080604052600436106100965760003560e01c80639a7d0a2811610069578063ba9210041161004e578063ba92100414610166578063c047d41014610186578063e3168f9e1461019957600080fd5b80639a7d0a2814610126578063b40014641461015357600080fd5b8063112346c21461009b5780632384a6aa146100bd5780632ba010d7146100d0578063931bb19a14610106575b600080fd5b3480156100a757600080fd5b506100bb6100b6366004613386565b6101c6565b005b6100bb6100cb3660046133ea565b6102e1565b3480156100dc57600080fd5b506100f06100eb36600461347b565b610547565b6040516100fd9190613b8f565b60405180910390f35b34801561011257600080fd5b506100bb61012136600461347b565b6108fe565b34801561013257600080fd5b5061014661014136600461332e565b6109a5565b6040516100fd919061398e565b6100bb61016136600461341e565b610a76565b34801561017257600080fd5b506100bb6101813660046133ea565b610b49565b6100bb61019436600461341e565b61100f565b3480156101a557600080fd5b506101b96101b4366004613310565b611154565b6040516100fd919061397d565b600054610100900460ff166101e15760005460ff16156101e5565b303b155b610224576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021b90613b2e565b60405180910390fd5b600054610100900460ff16158015610246576000805461ffff19166101011790555b6000805483516001600160401b0316760100000000000000000000000000000000000000000000027fffff0000000000000000ffffffffffffffffffffffffffffffffffffffffffff6001600160a01b0387166201000002167fffff00000000000000000000000000000000000000000000000000000000ffff9092169190911717905580156102dc576000805461ff00191690555b505050565b80516001600160a01b03166000908152600160205260408120905b81546001600160401b038216101561051f5782602001516001600160401b031682826001600160401b03168154811061034557634e487b7160e01b600052603260045260246000fd5b6000918252602090912060069091020154600160a01b90046001600160401b0316141561050d578282826001600160401b03168154811061039657634e487b7160e01b600052603260045260246000fd5b60009182526020918290208351600692909202018054928401516001600160a01b039092166001600160e01b031990931692909217600160a01b6001600160401b0392831602178255604083015160018301805460608601519284166fffffffffffffffffffffffffffffffff1990911617600160401b9290931691909102919091178082556080840151919060ff60801b1916600160801b83600281111561044f57634e487b7160e01b600052602160045260246000fd5b021790555060a0820151600282015560c0820151600382015560e08201516004820180546101008501516101208601516101408701511515600160c01b0260ff60c01b196001600160401b03928316600160801b021668ffffffffffffffffff60801b19938316600160401b026fffffffffffffffffffffffffffffffff19909516929096169190911792909217169290921791909117905561016082015180516105049160058401916020909101906125e0565b5090505061051f565b8061051781613dc7565b9150506102fc565b5081516001600160a01b0316600090815260016020526040902081546102dc9190839061263d565b60408051610180810182526000808252602080830182905282840182905260608084018390526080840183905260a0840183905260c0840183905260e0840183905261010084018390526101208401839052610140840183905261016084015284516001600160a01b03168252600181528382208054855181840281018401909652808652939492939091849084015b828210156107d957600084815260209081902060408051610180810182526006860290920180546001600160a01b0381168452600160a01b90046001600160401b0390811694840194909452600181015480851692840192909252600160401b820490931660608301529091906080830190600160801b900460ff16600281111561067257634e487b7160e01b600052602160045260246000fd5b600281111561069157634e487b7160e01b600052602160045260246000fd5b81526002820154602080830191909152600383015460408084019190915260048401546001600160401b038082166060860152600160401b820481166080860152600160801b82041660a0850152600160c01b900460ff16151560c084015260058401805482518185028101850190935280835260e090940193919290919060009084015b828210156107c257838290600052602060002001805461073590613d53565b80601f016020809104026020016040519081016040528092919081815260200182805461076190613d53565b80156107ae5780601f10610783576101008083540402835291602001916107ae565b820191906000526020600020905b81548152906001019060200180831161079157829003601f168201915b505050505081526020019060010190610716565b5050505081525050815260200190600101906105d7565b50505050905060005b8151816001600160401b031610156108925783602001516001600160401b031682826001600160401b03168151811061082b57634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516001600160401b031614156108805781816001600160401b03168151811061086f57634e487b7160e01b600052603260045260246000fd5b602002602001015192505050919050565b8061088a81613dc7565b9150506107e2565b50604080516101808101825260008082526020820181905291810182905260608082018390526080820183905260a0820183905260c0820183905260e08201839052610100820183905261012082018390526101408201929092526101608101919091525b9392505050565b600061090982610547565b6101008101519091506001600160401b031615610959577f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f560405161094d90613b0b565b60405180910390a15050565b61096b826000015183602001516117f1565b7f4fca90fd1b1cd962cf67f21703f1380572181450811cdd09c4add7ed1cb0c12c60094333856020015160405161094d949392919061399c565b6000806109b184611154565b905060005b8151811015610a6957846001600160a01b03168282815181106109e957634e487b7160e01b600052603260045260246000fd5b6020026020010151600001516001600160a01b0316148015610a475750836001600160401b0316828281518110610a3057634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516001600160401b0316145b15610a5757600192505050610a70565b80610a6181613dac565b9150506109b6565b5060009150505b92915050565b6000610a8f836000015184602001518460000151611809565b61010084018051919250610aa282613d33565b6001600160401b0316905250608082015160e084018051610ac4908390613c98565b6001600160401b031690525060a08201516080830151610ae49190613c41565b83606001818151610af59190613c98565b6001600160401b03169052508015610b25576101208301805190610b1882613d33565b6001600160401b03169052505b6101008301516001600160401b0316610b4057600060c08401525b6102dc836102e1565b60208101516001600160401b0316610b93577f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5604051610b8890613a5c565b60405180910390a150565b60408101516001600160401b0316610bd2577f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5604051610b8890613ac5565b60005481516040517f66838994000000000000000000000000000000000000000000000000000000008152620100009092046001600160a01b031691636683899491610c209160040161396f565b60206040518083038186803b158015610c3857600080fd5b505afa158015610c4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c709190613368565b610ca1577f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5604051610b8890613a39565b610cd8604051806040016040528083600001516001600160a01b0316815260200183602001516001600160401b0316815250610547565b604001516001600160401b031615610d17577f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5604051610b8890613aa2565b6000805482516040517f1a65374a000000000000000000000000000000000000000000000000000000008152620100009092046001600160a01b031691631a65374a91610d669160040161396f565b60006040518083038186803b158015610d7e57600080fd5b505afa158015610d92573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610dba91908101906133b6565b90506000610dcb8360000151611947565b905081604001516001600160401b0316818460400151610deb9190613c09565b6001600160401b03161115610e34577f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5604051610e2790613a7f565b60405180910390a1505050565b82516001600160a01b03908116600090815260016020818152604080842080548085018255908552938290208851600690950201805492890151949095166001600160e01b031990921691909117600160a01b6001600160401b0394851602178455860151908301805460608801519284166fffffffffffffffffffffffffffffffff1990911617600160401b9290931691909102919091178082556080860151869392909160ff60801b1916600160801b836002811115610f0657634e487b7160e01b600052602160045260246000fd5b021790555060a0820151600282015560c0820151600382015560e08201516004820180546101008501516101208601516101408701511515600160c01b0260ff60c01b196001600160401b03928316600160801b021668ffffffffffffffffff60801b19938316600160401b026fffffffffffffffffffffffffffffffff1990951692909616919091179290921716929092179190911790556101608201518051610fbb9160058401916020909101906125e0565b5050507fc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb6008438560000151866020015187608001518860400151896101400151604051610e2797969594939291906139d1565b81604001516001600160401b03168160a0015182608001516110319190613c41565b83606001516110409190613c09565b6001600160401b0316111561107c577f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f560405161094d90613ae8565b60006110b68360000151846020015160405180604001604052808660000151815260200186608001516001600160401b0316815250611be7565b610100840180519192506110c982613dc7565b6001600160401b031690525060a082015160808301516110e99190613c41565b836060018181516110fa9190613c09565b6001600160401b0316905250608082015160e08401805161111c908390613c09565b6001600160401b03169052508015610b4057610120830180519061113f82613dc7565b6001600160401b03169052506102dc836102e1565b6001600160a01b0381166000908152600160209081526040808320805482518185028101850190935280835260609493849084015b8282101561138b57600084815260209081902060408051610180810182526006860290920180546001600160a01b0381168452600160a01b90046001600160401b0390811694840194909452600181015480851692840192909252600160401b820490931660608301529091906080830190600160801b900460ff16600281111561122457634e487b7160e01b600052602160045260246000fd5b600281111561124357634e487b7160e01b600052602160045260246000fd5b81526002820154602080830191909152600383015460408084019190915260048401546001600160401b038082166060860152600160401b820481166080860152600160801b82041660a0850152600160c01b900460ff16151560c084015260058401805482518185028101850190935280835260e090940193919290919060009084015b828210156113745783829060005260206000200180546112e790613d53565b80601f016020809104026020016040519081016040528092919081815260200182805461131390613d53565b80156113605780601f1061133557610100808354040283529160200191611360565b820191906000526020600020905b81548152906001019060200180831161134357829003601f168201915b5050505050815260200190600101906112c8565b505050508152505081526020019060010190611189565b5050505090506000805b8251816001600160401b0316101561140d57600083826001600160401b0316815181106113d257634e487b7160e01b600052603260045260246000fd5b6020026020010151604001516001600160401b031611156113fb57816113f781613dc7565b9250505b8061140581613dc7565b915050611395565b506000816001600160401b03166001600160401b0381111561143f57634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156114d057816020015b60408051610180810182526000808252602080830182905292820181905260608083018290526080830182905260a0830182905260c0830182905260e0830182905261010083018290526101208301829052610140830191909152610160820152825260001990920191018161145d5790505b5090506000805b8451816001600160401b031610156117e657600085826001600160401b03168151811061151457634e487b7160e01b600052603260045260246000fd5b6020026020010151604001516001600160401b031611156117d457600085826001600160401b03168151811061155a57634e487b7160e01b600052603260045260246000fd5b60200260200101519050600061157a898360200151846101200151611e14565b80516001600160401b0316600090815260036020908152604080832080548251818502810185019093528083529495509293909291849084015b8282101561168b57838290600052602060002090600202016040518060400160405290816000820180546115e790613d53565b80601f016020809104026020016040519081016040528092919081815260200182805461161390613d53565b80156116605780601f1061163557610100808354040283529160200191611660565b820191906000526020600020905b81548152906001019060200180831161164357829003601f168201915b50505091835250506001918201546001600160401b03166020918201529183529290920191016115b4565b505050509050600081516001600160401b038111156116ba57634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156116ed57816020015b60608152602001906001900390816116d85790505b50905060005b8251816001600160401b031610156117815782816001600160401b03168151811061172e57634e487b7160e01b600052603260045260246000fd5b60200260200101516000015182826001600160401b03168151811061176357634e487b7160e01b600052603260045260246000fd5b6020026020010181905250808061177990613dc7565b9150506116f3565b50808461016001819052508387876001600160401b0316815181106117b657634e487b7160e01b600052603260045260246000fd5b602002602001018190525085806117cc90613dc7565b965050505050505b806117de81613dc7565b9150506114d7565b509095945050505050565b6117fb8282611fd0565b611805828261201e565b5050565b6000806118168585612191565b905060015b816001600160401b0316816001600160401b03161161193b576000611841878784611e14565b905060008061185083886121d0565b9150915080156119255782516001600160401b039081166000908152600360205260408082205486518416835291208054919290919085169081106118a557634e487b7160e01b600052603260045260246000fd5b600091825260208220600290910201906118bf828261285e565b50600101805467ffffffffffffffff191690556001600160401b0383166118e857602084018890525b6118f3600182613c7d565b836001600160401b0316141561190b57604084018890525b6119168a8a8661238c565b600196505050505050506108f7565b505050808061193390613dc7565b91505061181b565b50600095945050505050565b6001600160a01b0381166000908152600160209081526040808320805482518185028101850190935280835284938493929190849084015b82821015611b8157600084815260209081902060408051610180810182526006860290920180546001600160a01b0381168452600160a01b90046001600160401b0390811694840194909452600181015480851692840192909252600160401b820490931660608301529091906080830190600160801b900460ff166002811115611a1a57634e487b7160e01b600052602160045260246000fd5b6002811115611a3957634e487b7160e01b600052602160045260246000fd5b81526002820154602080830191909152600383015460408084019190915260048401546001600160401b038082166060860152600160401b820481166080860152600160801b82041660a0850152600160c01b900460ff16151560c084015260058401805482518185028101850190935280835260e090940193919290919060009084015b82821015611b6a578382906000526020600020018054611add90613d53565b80601f0160208091040260200160405190810160405280929190818152602001828054611b0990613d53565b8015611b565780601f10611b2b57610100808354040283529160200191611b56565b820191906000526020600020905b815481529060010190602001808311611b3957829003601f168201915b505050505081526020019060010190611abe565b50505050815250508152602001906001019061197f565b50505050905060005b8151811015611bde57818181518110611bb357634e487b7160e01b600052603260045260246000fd5b60200260200101516040015183611bca9190613c09565b925080611bd681613dac565b915050611b8a565b50909392505050565b6040805160608082018352600080835260208301829052928201526000806000611c3660405180604001604052808a6001600160a01b03168152602001896001600160401b0316815250610547565b6101208101519350905060606001600160401b038416611c78576040805160608101825260018082526020820184905291810183905295509350839250611dfb565b611c83898986611e14565b80516001600160401b0316600090815260036020908152604080832080548251818502810185019093528083529499509293909291849084015b82821015611d945783829060005260206000209060020201604051806040016040529081600082018054611cf090613d53565b80601f0160208091040260200160405190810160405280929190818152602001828054611d1c90613d53565b8015611d695780601f10611d3e57610100808354040283529160200191611d69565b820191906000526020600020905b815481529060010190602001808311611d4c57829003601f168201915b50505091835250506001918201546001600160401b0316602091820152918352929092019101611cbd565b505050509050600060169054906101000a90046001600160401b03166001600160401b031681511415611df95784611dcb81613dc7565b955050600193506040518060600160405280866001600160401b031681526020018381526020018381525095505b505b611e078989878a612432565b5090979650505050505050565b611e41604051806060016040528060006001600160401b0316815260200160608152602001606081525090565b6000848484604051602001611e589392919061392c565b6040516020818303038152906040529050600281604051611e799190613963565b9081526040805191829003602090810183206060840190925281546001600160401b0316835260018201805491840191611eb290613d53565b80601f0160208091040260200160405190810160405280929190818152602001828054611ede90613d53565b8015611f2b5780601f10611f0057610100808354040283529160200191611f2b565b820191906000526020600020905b815481529060010190602001808311611f0e57829003601f168201915b50505050508152602001600282018054611f4490613d53565b80601f0160208091040260200160405190810160405280929190818152602001828054611f7090613d53565b8015611fbd5780601f10611f9257610100808354040283529160200191611fbd565b820191906000526020600020905b815481529060010190602001808311611fa057829003601f168201915b5050505050815250509150509392505050565b6000611fdc8383612191565b905060005b816001600160401b0316816001600160401b031610156120185761200684848361254a565b8061201081613dc7565b915050611fe1565b50505050565b60005b6001600160a01b0383166000908152600160205260409020546001600160401b03821610156102dc576001600160a01b038316600090815260016020526040902080546001600160401b03808516929190841690811061209157634e487b7160e01b600052603260045260246000fd5b6000918252602090912060069091020154600160a01b90046001600160401b0316141561217f576001600160a01b038316600090815260016020526040902080546001600160401b0383169081106120f957634e487b7160e01b600052603260045260246000fd5b60009182526020822060069091020180546001600160e01b031916815560018101805470ffffffffffffffffffffffffffffffffff1916905560028101829055600381018290556004810180547fffffffffffffff0000000000000000000000000000000000000000000000000016905590612178600583018261289b565b5050505050565b8061218981613dc7565b915050612021565b6000806121c36040518060400160405280866001600160a01b03168152602001856001600160401b0316815250610547565b6101200151949350505050565b81516001600160401b03166000908152600360209081526040808320805482518185028101850190935280835284938493929190849084015b828210156122e0578382906000526020600020906002020160405180604001604052908160008201805461223c90613d53565b80601f016020809104026020016040519081016040528092919081815260200182805461226890613d53565b80156122b55780601f1061228a576101008083540402835291602001916122b5565b820191906000526020600020905b81548152906001019060200180831161229857829003601f168201915b50505091835250506001918201546001600160401b0316602091820152918352929092019101612209565b5050505090508051600014156122fd576000809250925050612385565b60005b8151816001600160401b0316101561237b57848051906020012082826001600160401b03168151811061234357634e487b7160e01b600052603260045260246000fd5b602002602001015160000151805190602001201415612369579250600191506123859050565b8061237381613dc7565b915050612300565b5060008092509250505b9250929050565b80516040516000916123a4918691869160200161392c565b6040516020818303038152906040529050816002826040516123c69190613963565b908152604051602091819003820190208251815467ffffffffffffffff19166001600160401b03909116178155828201518051919261240d926001850192909101906128b9565b50604082015180516124299160028401916020909101906128b9565b50505050505050565b815160405160009161244a918791879160200161392c565b60405160208183030381529060405290508260028260405161246c9190613963565b908152604051602091819003820190208251815467ffffffffffffffff19166001600160401b0390911617815582820151805191926124b3926001850192909101906128b9565b50604082015180516124cf9160028401916020909101906128b9565b505083516001600160401b0316600090815260036020908152604082208054600181018255908352918190208551805187955060029094029091019261251892849201906128b9565b50602091909101516001909101805467ffffffffffffffff19166001600160401b039092169190911790555050505050565b60008383836040516020016125619392919061392c565b60405160208183030381529060405290506002816040516125829190613963565b908152604051908190036020019020805467ffffffffffffffff1916815560006125af600183018261285e565b6125bd60028301600061285e565b50506001600160401b038216600090815260036020526040812061201891612939565b82805482825590600052602060002090810192821561262d579160200282015b8281111561262d578251805161261d9184916020909101906128b9565b5091602001919060010190612600565b5061263992915061295a565b5090565b8280548282559060005260206000209060060281019282156128525760005260206000209160060282015b8281111561285257825482546001600160a01b039091167fffffffffffffffffffffffff000000000000000000000000000000000000000082168117845584546001600160e01b031990921617600160a01b918290046001600160401b0390811690920217835560018085018054918501805492841667ffffffffffffffff1984168117825582546fffffffffffffffffffffffffffffffff1990941617600160401b9384900490941690920292909217808255915485928592600160801b9283900460ff1692909160ff60801b19169083600281111561275957634e487b7160e01b600052602160045260246000fd5b0217905550600282810154908201556003808301549082015560048083018054918301805467ffffffffffffffff1981166001600160401b0394851690811783558354600160401b908190048616026fffffffffffffffffffffffffffffffff1990921617178082558254600160801b908190049094169093027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff841681178255915460ff600160c01b918290041615150260ff60c01b1990921668ffffffffffffffffff60801b1990931692909217179055600580830180546128409284019190612977565b50505091600601919060060190612668565b506126399291506129cd565b50805461286a90613d53565b6000825580601f1061287a575050565b601f0160209004906000526020600020908101906128989190612a4f565b50565b5080546000825590600052602060002090810190612898919061295a565b8280546128c590613d53565b90600052602060002090601f0160209004810192826128e7576000855561292d565b82601f1061290057805160ff191683800117855561292d565b8280016001018555821561292d579182015b8281111561292d578251825591602001919060010190612912565b50612639929150612a4f565b50805460008255600202906000526020600020908101906128989190612a64565b8082111561263957600061296e828261285e565b5060010161295a565b82805482825590600052602060002090810192821561262d5760005260206000209182015b8281111561262d5782829080546129b290613d53565b6129bd929190612a94565b509160010191906001019061299c565b808211156126395780546001600160e01b031916815560018101805470ffffffffffffffffffffffffffffffffff19169055600060028201819055600382018190556004820180547fffffffffffffff00000000000000000000000000000000000000000000000000169055612a46600583018261289b565b506006016129cd565b5b808211156126395760008155600101612a50565b80821115612639576000612a78828261285e565b5060018101805467ffffffffffffffff19169055600201612a64565b828054612aa090613d53565b90600052602060002090601f016020900481019282612ac2576000855561292d565b82601f10612ad3578054855561292d565b8280016001018555821561292d57600052602060002091601f016020900482015b8281111561292d578254825591600101919060010190612af4565b6000612b22612b1d84613bbc565b613ba0565b90508083825260208201905082856020860282011115612b4157600080fd5b60005b85811015612b6d5781612b578882612c53565b8452506020928301929190910190600101612b44565b5050509392505050565b6000612b85612b1d84613bbc565b90508083825260208201905082856020860282011115612ba457600080fd5b60005b85811015612b6d5781356001600160401b03811115612bc557600080fd5b808601612bd28982612cc9565b855250506020928301929190910190600101612ba7565b6000612bf7612b1d84613bdf565b905082815260208101848484011115612c0f57600080fd5b612c1a848285613cfb565b509392505050565b6000612c30612b1d84613bdf565b905082815260208101848484011115612c4857600080fd5b612c1a848285613d07565b8035610a7081613e82565b8051610a7081613e82565b600082601f830112612c7a57600080fd5b8135612c8a848260208601612b0f565b949350505050565b600082601f830112612ca357600080fd5b8135612c8a848260208601612b77565b8035610a7081613e96565b8051610a7081613e96565b600082601f830112612cda57600080fd5b8135612c8a848260208601612be9565b600082601f830112612cfb57600080fd5b8151612c8a848260208601612c22565b8035610a7081613e9e565b8035610a7081613ea7565b8035610a7081613eb4565b60006103208284031215612d3f57600080fd5b612d4a6102e0613ba0565b905081356001600160401b03811115612d6257600080fd5b612d6e84828501612cc9565b8252506020612d7f84848301612c53565b60208301525060408201356001600160401b03811115612d9e57600080fd5b612daa84828501612cc9565b6040830152506060612dbe848285016132fa565b6060830152506080612dd2848285016132fa565b60808301525060a0612de6848285016132fa565b60a08301525060c0612dfa848285016132fa565b60c08301525060e0612e0e848285016132fa565b60e083015250610100612e23848285016132ef565b61010083015250610120612e39848285016132fa565b61012083015250610140612e4f848285016132fa565b610140830152506101608201356001600160401b03811115612e7057600080fd5b612e7c84828501612fc0565b61016083015250610180612e92848285016132fa565b610180830152506101a0612ea8848285016132ef565b6101a0830152506101c0612ebe84828501612cb3565b6101c0830152506101e0612ed484828501612d21565b6101e083015250610200612eea848285016132fa565b610200830152506102208201356001600160401b03811115612f0b57600080fd5b612f1784828501612c69565b610220830152506102408201356001600160401b03811115612f3857600080fd5b612f4484828501612c69565b610240830152506102608201356001600160401b03811115612f6557600080fd5b612f7184828501612cc9565b61026083015250610280612f8784828501612d16565b610280830152506102a0612f9d84828501612cb3565b6102a0830152506102c0612fb3848285016130f6565b6102c08301525092915050565b600060408284031215612fd257600080fd5b612fdc6040613ba0565b905081356001600160401b03811115612ff457600080fd5b61300084828501612cc9565b82525060208201356001600160401b0381111561301c57600080fd5b61302884828501612cc9565b60208301525092915050565b600060e0828403121561304657600080fd5b61305060e0613ba0565b9050600061305e8484613305565b825250602061306f84848301613305565b602083015250604061308384828501613305565b604083015250606061309784828501613305565b60608301525060806130ab84828501613305565b60808301525060a06130bf84828501612c5e565b60a08301525060c08201516001600160401b038111156130de57600080fd5b6130ea84828501612cea565b60c08301525092915050565b60006060828403121561310857600080fd5b6131126060613ba0565b9050600061312084846132fa565b8252506020613131848483016132fa565b6020830152506040613145848285016132fa565b60408301525092915050565b60006020828403121561316357600080fd5b61316d6020613ba0565b9050600061317b84846132fa565b82525092915050565b6000610180828403121561319757600080fd5b6131a2610180613ba0565b905060006131b08484612c53565b82525060206131c1848483016132fa565b60208301525060406131d5848285016132fa565b60408301525060606131e9848285016132fa565b60608301525060806131fd84828501612d16565b60808301525060a0613211848285016132ef565b60a08301525060c0613225848285016132ef565b60c08301525060e0613239848285016132fa565b60e08301525061010061324e848285016132fa565b61010083015250610120613264848285016132fa565b6101208301525061014061327a84828501612cb3565b610140830152506101608201356001600160401b0381111561329b57600080fd5b6132a784828501612c92565b6101608301525092915050565b6000604082840312156132c657600080fd5b6132d06040613ba0565b905060006132de8484612c53565b8252506020613028848483016132fa565b8035610a7081613ec1565b8035610a7081613ec7565b8051610a7081613ec7565b60006020828403121561332257600080fd5b6000612c8a8484612c53565b6000806040838503121561334157600080fd5b600061334d8585612c53565b925050602061335e858286016132fa565b9150509250929050565b60006020828403121561337a57600080fd5b6000612c8a8484612cbe565b6000806040838503121561339957600080fd5b60006133a58585612d0b565b925050602061335e85828601613151565b6000602082840312156133c857600080fd5b81516001600160401b038111156133de57600080fd5b612c8a84828501613034565b6000602082840312156133fc57600080fd5b81356001600160401b0381111561341257600080fd5b612c8a84828501613184565b6000806040838503121561343157600080fd5b82356001600160401b0381111561344757600080fd5b61345385828601613184565b92505060208301356001600160401b0381111561346f57600080fd5b61335e85828601612d2c565b60006040828403121561348d57600080fd5b6000612c8a84846132b4565b60006108f78383613598565b60006108f78383613809565b6134ba81613cb5565b82525050565b6134ba6134cc82613cb5565b613dec565b60006134db825190565b808452602084019350836020820285016134f58560200190565b8060005b8581101561352a57848403895281516135128582613499565b94506020830160209a909a01999250506001016134f9565b5091979650505050505050565b6000613541825190565b8084526020840193508360208202850161355b8560200190565b8060005b8581101561352a578484038952815161357885826134a5565b94506020830160209a909a019992505060010161355f565b8015156134ba565b60006135a2825190565b8084526020840193506135b9818560208601613d07565b601f01601f19169290920192915050565b6134ba81613ce5565b6134ba81613cf0565b60006135e6825190565b6135f4818560208601613d07565b9290920192915050565b601381526000602082017f4e6f6465206e6f74207265676973746572656400000000000000000000000000815291505b5060200190565b600c81526000602082017f437265617465536563746f7200000000000000000000000000000000000000008152915061362e565b600d81526000602082017f536563746f7249442069732030000000000000000000000000000000000000008152915061362e565b600f81526000602082017f41646446696c65546f536563746f7200000000000000000000000000000000008152915061362e565b600c81526000602082017f44656c657465536563746f7200000000000000000000000000000000000000008152915061362e565b600e81526000602082017f4e6f74456e6f75676853706163650000000000000000000000000000000000008152915061362e565b601381526000602082017f496e73756666696369656e7420766f6c756d65000000000000000000000000008152915061362e565b600e81526000602082017f4e6f74456d707479536563746f720000000000000000000000000000000000008152915061362e565b601581526000602082017f536563746f7220616c72656164792065786973747300000000000000000000008152915061362e565b600981526000602082017f53697a65206973203000000000000000000000000000000000000000000000008152915061362e565b805160009061018084019061381e85826134b1565b506020830151613831602086018261390b565b506040830151613844604086018261390b565b506060830151613857606086018261390b565b50608083015161386a60808601826135d3565b5060a083015161387d60a0860182613905565b5060c083015161389060c0860182613905565b5060e08301516138a360e086018261390b565b506101008301516138b861010086018261390b565b506101208301516138cd61012086018261390b565b506101408301516138e2610140860182613590565b506101608301518482036101608601526138fc82826134d1565b95945050505050565b806134ba565b6001600160401b0381166134ba565b6134ba6001600160401b038216613dfe565b600061393882866134c0565b601482019150613948828561391a565b600882019150613958828461391a565b506008019392505050565b60006108f782846135dc565b60208101610a7082846134b1565b602080825281016108f78184613537565b60208101610a708284613590565b608081016139aa82876135ca565b6139b76020830186613905565b6139c460408301856134b1565b6138fc606083018461390b565b60e081016139df828a6135ca565b6139ec6020830189613905565b6139f960408301886134b1565b613a06606083018761390b565b613a1360808301866135d3565b613a2060a083018561390b565b613a2d60c0830184613590565b98975050505050505050565b60408082528101613a4981613635565b90508181036020830152610a70816135fe565b60408082528101613a6c81613635565b90508181036020830152610a7081613669565b60408082528101613a8f81613635565b90508181036020830152610a7081613739565b60408082528101613ab281613635565b90508181036020830152610a70816137a1565b60408082528101613ad581613635565b90508181036020830152610a70816137d5565b60408082528101613af88161369d565b90508181036020830152610a7081613705565b60408082528101613b1b816136d1565b90508181036020830152610a708161376d565b60208082528101610a7081602e81527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160208201527f647920696e697469616c697a6564000000000000000000000000000000000000604082015260600190565b602080825281016108f78184613809565b6000613bab60405190565b9050613bb78282613d80565b919050565b60006001600160401b03821115613bd557613bd5613e4c565b5060209081020190565b60006001600160401b03821115613bf857613bf8613e4c565b601f19601f83011660200192915050565b60006001600160401b03821691506001600160401b0383169250826001600160401b0303821115613c3c57613c3c613e0a565b500190565b60006001600160401b03821691506001600160401b0383169250816001600160401b030483118215151615613c7857613c78613e0a565b500290565b6000825b925082821015613c9357613c93613e0a565b500390565b60006001600160401b03821691506001600160401b038316613c81565b60006001600160a01b038216610a70565b6000610a7082613cb5565b80613bb781613e62565b80613bb781613e72565b6000610a7082613cd1565b6000610a7082613cdb565b82818337506000910152565b60005b83811015613d22578181015183820152602001613d0a565b838111156120185750506000910152565b6001600160401b0316600081613d4b57613d4b613e0a565b506000190190565b600281046001821680613d6757607f821691505b60208210811415613d7a57613d7a613e36565b50919050565b601f19601f83011681018181106001600160401b0382111715613da557613da5613e4c565b6040525050565b6000600019821415613dc057613dc0613e0a565b5060010190565b60006001600160401b03821691506001600160401b03821415613dc057613dc0613e0a565b6000610a70826000610a708260601b90565b6000610a708260c01b90565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600b811061289857612898613e20565b6003811061289857612898613e20565b613e8b81613cb5565b811461289857600080fd5b801515613e8b565b613e8b81613cc6565b6003811061289857600080fd5b6002811061289857600080fd5b80613e8b565b6001600160401b038116613e8b56fea2646970667358221220c26df68dc8e32e07fa43780d0ca2e90f9f5d28d029baccdcfe43316dd5a63a0e64736f6c63430008040033",
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

// AddFileToSector is a paid mutator transaction binding the contract method 0xc047d410.
//
// Solidity: function AddFileToSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreTransactor) AddFileToSector(opts *bind.TransactOpts, sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "AddFileToSector", sectorInfo, fileInfo)
}

// AddFileToSector is a paid mutator transaction binding the contract method 0xc047d410.
//
// Solidity: function AddFileToSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreSession) AddFileToSector(sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.AddFileToSector(&_Store.TransactOpts, sectorInfo, fileInfo)
}

// AddFileToSector is a paid mutator transaction binding the contract method 0xc047d410.
//
// Solidity: function AddFileToSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
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

// DeleteFileFromSector is a paid mutator transaction binding the contract method 0xb4001464.
//
// Solidity: function DeleteFileFromSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreTransactor) DeleteFileFromSector(opts *bind.TransactOpts, sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteFileFromSector", sectorInfo, fileInfo)
}

// DeleteFileFromSector is a paid mutator transaction binding the contract method 0xb4001464.
//
// Solidity: function DeleteFileFromSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
func (_Store *StoreSession) DeleteFileFromSector(sectorInfo SectorInfo, fileInfo FileInfo) (*types.Transaction, error) {
	return _Store.Contract.DeleteFileFromSector(&_Store.TransactOpts, sectorInfo, fileInfo)
}

// DeleteFileFromSector is a paid mutator transaction binding the contract method 0xb4001464.
//
// Solidity: function DeleteFileFromSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo) payable returns()
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
