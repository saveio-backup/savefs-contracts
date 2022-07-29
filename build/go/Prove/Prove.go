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

// FileProveParams is an auto generated low-level Go binding around an user-defined struct.
type FileProveParams struct {
	FileHash    []byte
	ProveData   []byte
	BlockHeight *big.Int
	NodeWallet  common.Address
	Profit      uint64
	SectorID    uint64
}

// NodeInfo is an auto generated low-level Go binding around an user-defined struct.
type NodeInfo struct {
	Pledge      uint64
	Profit      uint64
	Volume      uint64
	RestVol     uint64
	ServiceTime uint64
	WalletAddr  common.Address
	NodeAddr    common.Address
}

// PlotInfo is an auto generated low-level Go binding around an user-defined struct.
type PlotInfo struct {
	NumberID   uint64
	StartNonce uint64
	Nonces     uint64
}

// ProveConfig is an auto generated low-level Go binding around an user-defined struct.
type ProveConfig struct {
	SECTORPROVEBLOCKNUM uint64
}

// ProveDetail is an auto generated low-level Go binding around an user-defined struct.
type ProveDetail struct {
	NodeAddr    common.Address
	WalletAddr  common.Address
	ProveTimes  uint64
	BlockHeight *big.Int
	Finished    bool
}

// ProveDetailMeta is an auto generated low-level Go binding around an user-defined struct.
type ProveDetailMeta struct {
	CopyNum        uint64
	ProveDetailNum uint64
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

// SectorProveParams is an auto generated low-level Go binding around an user-defined struct.
type SectorProveParams struct {
	NodeAddr        common.Address
	SectorID        uint64
	ChallengeHeight uint64
	ProveData       []byte
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

// TransferState is an auto generated low-level Go binding around an user-defined struct.
type TransferState struct {
	From  common.Address
	To    common.Address
	Value uint64
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DifferenceFileOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"FileNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"FileProveFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileProveNotFileOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FirstUserSpaceOperationError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProfit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeSectorProvedInTimeError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"want\",\"type\":\"uint64\"}],\"name\":\"NotEmptySector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughPledge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"want\",\"type\":\"uint64\"}],\"name\":\"NotEnoughVolume\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"OpError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParamsError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SectorOpError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SectorProveFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"UserspaceChangeError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserspaceDeleteError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceWrongExpiredHeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroProfit\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"provLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"CheckNodeSectorProvedInTime\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteProveDetails\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"NodeWallet\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"}],\"internalType\":\"structFileProveParams\",\"name\":\"fileProve\",\"type\":\"tuple\"}],\"name\":\"FileProve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"GetProveDetailList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ChallengeHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"}],\"internalType\":\"structSectorProveParams\",\"name\":\"sectorProve\",\"type\":\"tuple\"}],\"name\":\"SectorProve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"SetLastPunishmentHeightForNode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail\",\"name\":\"detail\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"details\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"}],\"name\":\"SettleForFile\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"details\",\"type\":\"tuple[]\"}],\"name\":\"UpdateProveDetailList\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveDetailNum\",\"type\":\"uint64\"}],\"internalType\":\"structProveDetailMeta\",\"name\":\"details\",\"type\":\"tuple\"}],\"name\":\"UpdateProveDetailMeta\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0xd11e680a\",\"type\":\"bytes32\"}],\"name\":\"c_0xd11e680a\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ChallengeHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"}],\"internalType\":\"structSectorProveParams\",\"name\":\"sectorProve\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"}],\"name\":\"checkSectorProve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractIFile\",\"name\":\"_fs\",\"type\":\"address\"},{\"internalType\":\"contractINode\",\"name\":\"_node\",\"type\":\"address\"},{\"internalType\":\"contractIPDP\",\"name\":\"_pdp\",\"type\":\"address\"},{\"internalType\":\"contractISector\",\"name\":\"_sector\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"SECTOR_PROVE_BLOCK_NUM\",\"type\":\"uint64\"}],\"internalType\":\"structProveConfig\",\"name\":\"proveConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractProveExtra\",\"name\":\"_proveExtra\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"}],\"name\":\"profitSplitForSector\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"times\",\"type\":\"uint64\"}],\"name\":\"punishForSector\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5061f45080620000226000396000f3fe6080604052600436106100dd5760003560e01c80638e2705301161007f5780639a19c98a116100595780639a19c98a14610249578063bb4e4e4214610279578063d2561e0a14610295578063fda59351146102be576100dd565b80638e270530146101c757806396961614146101e3578063977fdfe21461020c576100dd565b80633ec0f5df116100bb5780633ec0f5df1461015757806352e0614614610173578063648d225d1461018f5780637b85255a146101ab576100dd565b806309cbe391146100e25780630fece869146100fe57806327ab44341461013b575b600080fd5b6100fc60048036038101906100f7919061d45d565b6102da565b005b34801561010a57600080fd5b506101256004803603810190610120919061d49e565b6113dd565b604051610132919061e55c565b60405180910390f35b6101556004803603810190610150919061d017565b612826565b005b610171600480360381019061016c919061ceb3565b61293d565b005b61018d6004803603810190610188919061cfd6565b612a2f565b005b6101a960048036038101906101a4919061d50a565b612b43565b005b6101c560048036038101906101c0919061d1b6565b6136a8565b005b6101e160048036038101906101dc919061d261565b614459565b005b3480156101ef57600080fd5b5061020a6004803603810190610205919061cfad565b6175b6565b005b34801561021857600080fd5b50610233600480360381019061022e919061cfd6565b6175b9565b604051610240919061e53a565b60405180910390f35b610263600480360381019061025e919061d376565b6176f6565b604051610270919061e55c565b60405180910390f35b610293600480360381019061028e919061d083565b6184f4565b005b3480156102a157600080fd5b506102bc60048036038101906102b7919061d0d7565b61860b565b005b6102d860048036038101906102d3919061d3df565b618b3d565b005b6103067fcba2a0cc461453a5a7dcfee5c9dcd715d4f83d37ac41e41c4d62c8c7a00a9a1460001b6175b6565b6103327f14af8f1265bd3d38aadbe6e425ddd83e7b561818296a0cce7fc06c33c0ff7ccf60001b6175b6565b61035e7f792ce25d7bf81fff1ba38410a538981d6334604ba0e7cc3c202dde46c45f185b60001b6175b6565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633778febe83600001516040518263ffffffff1660e01b81526004016103bf919061e4c6565b60e06040518083038186803b1580156103d757600080fd5b505afa1580156103eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061040f919061d2a2565b905061043d7f541b62456456bfe57e9d2e62c294a4b1e8f71207f40471e92286403d39466a6760001b6175b6565b6104697fe091ec7180f6261f0ad63793a76152fd3674588765497234488b7ebbe52ad43760001b6175b6565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d76040518060400160405280866000015173ffffffffffffffffffffffffffffffffffffffff168152602001866020015167ffffffffffffffff168152506040518263ffffffff1660e01b8152600401610502919061e950565b60006040518083038186803b15801561051a57600080fd5b505afa15801561052e573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610557919061d335565b90506105857f6a5e2e78cfc77af8e92a700bcab862c87a0179ac0fc1fd7979c551526c8376ee60001b6175b6565b6105b17f6807fac6fd60913fee8164afedb4b765b64bd5e8e61dd658cb62ef2a74c49ab860001b6175b6565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561061b57600080fd5b505afa15801561062f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610653919061d533565b90506106817f0f37245d1094ee6542ef5556083a0020a077f034d6d780dc2b9313ce16cd5dfb60001b6175b6565b6106ad7f50b277d23a8e03fd5f4c12764bd44351959bddb7ded3df9538b503d2c727e55760001b6175b6565b8160c0015143101561074f576106e57f5adc8d4f4d7a477875b7dfd6b610cb542621449e2d2077135d771333ff5d42b460001b6175b6565b6107117ffd8702c811b350a86045a376eeb00ad572c69a992d090b420e5329a0b2d9b98c60001b6175b6565b60016040517fcab88a3c000000000000000000000000000000000000000000000000000000008152600401610746919061e6ac565b60405180910390fd5b61077b7f2f41e2cd2413c64350bf786bb4f8d541edb1c957b6e1df4b16455a4fd6b6be4460001b6175b6565b6107a77f3bccc43751f4498a0355e02df6206998ef019dad44c4521cee56097b54c2905460001b6175b6565b6107d37f40165c08f35e772a70c71c562f9cac1b70d03fc7adf6229f46a0a85388a184f660001b6175b6565b8160c00151846040015167ffffffffffffffff161461092c576108187fdbc6002eca082629f6f8adb05be36d425ba8bc4e18f3b8a209a27bf4fd9d939f60001b6175b6565b6108447f991b1dc2ab63d9e73dbf297609a8d6ef5e7abdf35c02125e8b0334e1eeaab59a60001b6175b6565b6108707f2d1f6f56c2bf4f80910c7f9b91cde3179b8d242c1d863e535928e8c46bd0ce5660001b6175b6565b6108c26040518060400160405280601681526020017f4368616c6c656e6765486569676874206572726f723a00000000000000000000815250856040015167ffffffffffffffff168460c0015161912b565b6108ee7f32b7539161197c774a2b90300b803df603ecf18e4d7205a0a1ab386b3f18c2de60001b6175b6565b60026040517fcab88a3c000000000000000000000000000000000000000000000000000000008152600401610923919061e6c7565b60405180910390fd5b6109587fd2e9076d3b1f29391c78961dd81873801916dafd0211407129f4bb398890795c60001b6175b6565b6109847fd4f87265a8e455774008a82b6f6e2f228da126f38c9af3cd62e718ca6e7b995a60001b6175b6565b6109b07f49516454ad1909cf8e2b080967efcb5e82ce25ed64ebc256d218cbdf361c057260001b6175b6565b60006109bc85846113dd565b90506109ea7fad130773bae93052fc71e17926cdc3786ef2bf3b2c3eec744b3ddb0bf571edcd60001b6175b6565b610a167f4f01fc12b60872eaecc1f0b3232a5405e43682b8dd7e03838fb0231b985bd15160001b6175b6565b80610b1657610a477fc678e8547e4dd67c92a03295c1bd6bb2bf81af3894d447d396ca3882bcdfadb760001b6175b6565b610a737f2e4b897e127bf439763d91a388841e7b864f12f52db11e498004b3f4a1a5eb8e60001b6175b6565b610a9f7f487d6fc02beddb9c6e7375d4f9838f3c1ffd4e3e5292a294364c2fbd0598b08660001b6175b6565b610aac8385846001618b3d565b610ad87f5b0add7cad7c89ec11e13bbfe5042280862b6a3faf46ac21f1e8c30a33b18d1760001b6175b6565b60036040517fcab88a3c000000000000000000000000000000000000000000000000000000008152600401610b0d919061e6e2565b60405180910390fd5b610b427fce0c85e23f32a73d8a607d82b18d35a4b0ef6dd773c9ba4fe7a9bb8ed9005fb860001b6175b6565b610b6e7fb7694aea5332daec8d04ad02d30445113ef1ca4e0794e3376e8b7d22e1046b0060001b6175b6565b610b9a7f30076260222f766f0e692dfec16b9782badb1a945e187fa107030cfa918655a260001b6175b6565b6000610ba78486856176f6565b9050610bd47e9d3295b06ae142f00cb08da4b87ed4931dde71b0bd55241bf8c6119345a00d60001b6175b6565b610c007fb6a550dc706d56c15d71c667b2a82157499eac2027f3697180b152e53b85c3b960001b6175b6565b80610c9b57610c317ff789299c91951d87c97785833c8de3a5df9332047f3afb77e8de7f7a666a2df060001b6175b6565b610c5d7ff73ec4b5c97f8f40dc352934e3e8f9e383c2ad6e3c10ceecb8e9717367acb57360001b6175b6565b60056040517fcab88a3c000000000000000000000000000000000000000000000000000000008152600401610c92919061e718565b60405180910390fd5b610cc77fda0dff376f24bba595edb6350018ba8308ce77ef33987d67471c54dcdd5a2d5960001b6175b6565b610cf37feb240abeaf690f25ea4e088fbe6efced9ac0d045a722fb957ec4daa1691b8e7060001b6175b6565b610d1f7f2218e7a4a77fc02cd161f2ba253fd7a358f593ac611b48d078aa60c60d14e86a60001b6175b6565b60008460a001511415610dbf57610d587fe4170e320d89e4dbd42edff5bfc4a07ea30e8dbdd1c4367878a0f52dacfa5d8e60001b6175b6565b610d847f9c991a3db94afae4abc330aa4a7340a9c6401544895e38bab9a58acae938506b60001b6175b6565b610db07f3fd7c9bdf394090454d4d35e65ab3ba6a0d3bc31cb58477e392b906de6cd37b460001b6175b6565b438460a0018181525050610dec565b610deb7fcb5f52505ff28f5d770a7520a810127b4d11a3fa6c6cf5dbe7fa3228b5be229260001b6175b6565b5b610e187fcf5ee8ac67b2714408d0f6eb7da82f61ef005ec3c189842600baa5fd33e48a2e60001b6175b6565b610e447f5691a547b7f59abed046b7338dae5ef5acf68f9310462cba6a0b1e8fca9658e960001b6175b6565b8260c0015167ffffffffffffffff1643610e5e919061ed04565b8460c0018181525050610e937f5e4c8049834abd2083c7ff7cdf03f688973c319cb5959b3980136ee504e1c9a760001b6175b6565b610ebf7f614db3316a5bfa2e12c2b0bf29b848778ef03751d56daedef5eb09c1c69dfc5e60001b6175b6565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632384a6aa856040518263ffffffff1660e01b8152600401610f1a919061e8f7565b600060405180830381600087803b158015610f3457600080fd5b505af1158015610f48573d6000803e3d6000fd5b50505050610f787ff3f58db16b1c3f95f9385912b232b52db5d1e113d10a5e03581a6d9f434d254360001b6175b6565b610fa47fbf03e16d82fc02815197a62c8423507bb4e960b9826f67589401505d1c02736060001b6175b6565b83610140015161104457610fda7f716a642f926ccd67946dc7e71f597a8985898a5e3f3876dba2377aeee7daf52660001b6175b6565b6110067f2f5af5888f79870c350780c94cb8e400792ab4cc31ddf32b8ec083a853e945ef60001b6175b6565b60046040517fcab88a3c00000000000000000000000000000000000000000000000000000000815260040161103b919061e6fd565b60405180910390fd5b6110707fad69cbc3edf995939c8f266aca6bead961b3fef314476c85214740733ad058d060001b6175b6565b61109c7f3637d853d640eec20c8eb32ecc33d1716dfa152228c7f43f5ad3293754d6059260001b6175b6565b6110c87f48049f60921fe4073ad1fb4026b35d5e4de39d93be8b10c545c90a0d892ccb7960001b6175b6565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639908f2bf8660000151436040518363ffffffff1660e01b815260040161112b92919061e511565b60606040518083038186803b15801561114357600080fd5b505afa158015611157573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061117b919061d30c565b90506111a97fa12c1c03c908c2d92863891d6bd5f9f1d39f461229e63ec4b00e5f6ec737c0c360001b6175b6565b6111d57f7bed2600832e78b78af2b012119d0f37be5ebb3dc249f02a1265aee5a31a091060001b6175b6565b4381602001818152505061120b7f15d9ad84822fddcad19082529c5f79872ab53d83f62c47747736e7597076fa2260001b6175b6565b6112377fe041bfe03372568f851a67eabbd54dcae7a537a801518b51a9b45d0c6a4e4aed60001b6175b6565b8460000151816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505061129f7fb7904ba779d7dd9e8f94caf3a8a39912eef06172e6376253adedfa33ba58f94c60001b6175b6565b6112cb7f5e36c656103aff715e4ecd8010e2a0e4f30e3cea239491d776226430978d5eef60001b6175b6565b8460600151816040019067ffffffffffffffff16908167ffffffffffffffff168152505061131b7fd7926d66c6d9a69d70b5f94880e41402198d0e19e7c2380a72dbcfc44829bfb360001b6175b6565b6113477f73eba1007c4c6817bd9bb2d5bc4ef045ebdcc590e2f8cf2ce7d6a853b1db5d7960001b6175b6565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631581d545826040518263ffffffff1660e01b81526004016113a2919061e8ba565b600060405180830381600087803b1580156113bc57600080fd5b505af11580156113d0573d6000803e3d6000fd5b5050505050505050505050565b600061140b7fcf68ae11dd5355f2eb71939e3c9ea309194896ab10c93fcc31fe3a98c022433e60001b6175b6565b6114377f533833147fbceab0a04b24fcf808af5e521d9688e180cb7f8c19b3a5852bd4b560001b6175b6565b6114637ff28fd36563106efdc6d9fe2b1199afe1670dd325140464fa68fe4fe1788843df60001b6175b6565b60606114917fb7c915a3ccfffb3e0f9913f0199c0f56e7b6ef28b0108788ad38f218e9f350c960001b6175b6565b6114bd7f79dab4d53fc5137479b96022fbc3553c4f217b1b54d4a03ec3e1a708ca43e5e660001b6175b6565b6114c561af58565b6114f17f84dd572c893bc4666be08db74e72b794cf1cf3766b3f8787e93c0c5f0879bc0b60001b6175b6565b61151d7f856149dbfbdc9ade8bfbbf465ca1de45f6a684bdab79b1ba141d1f15e915fa3b60001b6175b6565b8460000151816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506115857ffba6b8807d935dbfd35ea6e619646dd01b9badd36d38a7ff1347cf28b1bf4c2d60001b6175b6565b6115b17f278bdf5abe17a911ce5505ac97c0fbfc83916f3eeeed2bc10f648485a38bc86e60001b6175b6565b8181602001819052506115e67fa8055c7b7986afe080c0e939e12a285369d5ef65c99ccf96b50ec8f6a3ec921a60001b6175b6565b6116127f144a7a75ded974443a7b4001cee160d4bd0f61babb30f7fbc8bb44ed0a62826d60001b6175b6565b8360e00151816040019067ffffffffffffffff16908167ffffffffffffffff16815250506116627f6560b0d97e7169eede6ae4b747136736170e31e0f5314264332309eaf5e4e9f360001b6175b6565b61168e7f0b32ed3c0e1f9da2f255bc41bdfea3a0049a2d5da59811d6a17f5502b1ef0ca260001b6175b6565b600560149054906101000a900467ffffffffffffffff16816060019067ffffffffffffffff16908167ffffffffffffffff16815250506116f07fe4d30d60a4489deb682edfb5fb66faf4a976d7f7546cad97a12b7bd79861c98d60001b6175b6565b61171c7f318899e60312fac39f1ea2bb39656137a3755e90caacb1e56f60128c07e6b68760001b6175b6565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fd7c9808836040518263ffffffff1660e01b8152600401611779919061e87d565b60006040518083038186803b15801561179157600080fd5b505afa1580156117a5573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906117ce919061cf02565b90506117fc7fe14529329d48899dd309a680836606b4f891d20feeb828e7a1565c3ba8b1a91060001b6175b6565b6118287fd1c782eeb1da7c6376a568b4d6174b9680865644c023d43eed15476c06227bb360001b6175b6565b61183061afaa565b61185c7fdb10f62e8da07d23c956a3ce17fc89806d429d2b212b90fcbe9c5241fafdf67960001b6175b6565b6118887f702641b0ac5eb6cfc0b282212a9a5d5a0b69c15ff52e3ec912c8bc066aef658460001b6175b6565b61189061aff4565b6118bc7fe226fab1292adf954a7556353498f89f470b960a169bacc49334be0a5c33174660001b6175b6565b6118e87f2e99d5282cbe9b41396ae6919676609262c8d8d0ca75d6bb408c72208b2240be60001b6175b6565b86816000018190525061191d7fd989359f747afac85feeaa7593de470cc0b82d783d7dc01f21859671cc52fc1560001b6175b6565b6119497fef12e34febd5dca298c9256a6863428bd118cb7cd8c75c6244dc38d4b830dcd360001b6175b6565b81816040018190525061197e7f80af9ff788fc10746f90a7644a9ff73e43de402ca7b6e11b427b2e57de3cb34660001b6175b6565b6119aa7f9369b24158b5773b6eb142032c753c6de77cb435a4de3c0f5230c082baa65ab560001b6175b6565b6119b261b021565b6119de7f2ad417701ed8b01f165513726224b5208aa68af68a38dbaf230ba9733c9a72e860001b6175b6565b611a0a7fabc5f60ca7d1dfaa706e98a572065d20c946f34a3d7d23ec7987ad4d099738bb60001b6175b6565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639f9ca644836040518263ffffffff1660e01b8152600401611a65919061e8d5565b60006040518083038186803b158015611a7d57600080fd5b505afa158015611a91573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190611aba919061d2cb565b9050611ae87f8e83d575f8f37dfea9f380f944bbf77494a65a0bd507c7d584f2d23cd214f16360001b6175b6565b611b147f291c615dfae5b522de5d1c1cb4e9eaf3d96fbca69d843788da59ae79b7a3bd0860001b6175b6565b8060c00151611bb057611b497f2d944243b9bddf92d4a48845310078ed2cf4254eb2f2983f0897f026efee7f3d60001b6175b6565b611b757f6a7ac2c5149de31d04e87289ab173ae1a47dee174bde2919e8e2eae942b2e4af60001b6175b6565b611ba17fe7c522d4634f7e51363a25ead2c246adef2f41b3bf391b640af8ec708a03546d60001b6175b6565b60009650505050505050612820565b611bdc7f478bf5bd5e7dbaa523ea69eda77cb8495a1a1ff7239a0a6f189a0973d933a10e60001b6175b6565b611c087fe29dda0f9d87cc11f63a2ec441e6d8d6c1e0b0f8b61db9e2430ab823a605836560001b6175b6565b611c337efad54934db310a1db713b9c93e7269423a2b3536bc746c64638e96fc485e8160001b6175b6565b611c3b61b066565b611c677fb5d74f82321cfce1104b93ed3e67d01ebeb87c9b51b3fa8b6e59ab535290c67d60001b6175b6565b611c937feec7f9682bff4dff108b9daf13e5eda10aacb22d0ff25cca5e52dea7ab0ff7b060001b6175b6565b6000816000019067ffffffffffffffff16908167ffffffffffffffff1681525050611ce07fc9d8d5b17f9e5ede8b4a493b7913daefba83d8e4bdcbabcbfb5668f81baced7460001b6175b6565b611d0c7fde96ca5b50fc244f9de7b52f733d5ed627e648134a1dbb22d178a8065e15627f60001b6175b6565b83604001518160200181905250611d457f850af0896df406f3d6347ddb631be8fafc943539234a94a6b611581a84939dae60001b6175b6565b611d717faedc44903731f7aaf0677b20130653760cbedb6787d6fe0faef7ec327aee701a60001b6175b6565b81600001518160400181905250611daa7f706f21ce7f27f5fcdd51ed62f082eaa7b539fd849b1872175561412cd5f328d260001b6175b6565b611dd67f9822250f810fb5b5f0181f5698003e94f1e4261dbaa9fc2a0502a1f830f29fed60001b6175b6565b81602001518160600181905250611e0f7f95918a839dcca798859f7ac840c513022c7882bbbacc0b7fd7865c73705d1aee60001b6175b6565b611e3b7f3599049c8a9afd14d3412e4e4ec718b86a8d625c07193e09a05441b905e5627660001b6175b6565b81604001518160800181905250611e747f369dca1a76d0fb0ce2103d50a14645fa678c89e7576bcea9cff8d9bb670c0e9a60001b6175b6565b611ea07fb1c6669e50a07b8604c911f0feaf86de1e6fe2344873e8ee4ba0c0bfcba3a11b60001b6175b6565b81606001518160a00181905250611ed97f78e02b59f2ea02684b65bc60a92bc56df855f05f17097eda82f6639f9451198160001b6175b6565b611f057fcd0712d85369c9305432746f10c2c03de35b8a5389e9819d9f9abe532e8c091860001b6175b6565b81608001518160c00181905250611f3e7f85c6d6306b8108696e412fb0c3e3ed7de5eb00ccf861586f7dece7fb4d6f231760001b6175b6565b611f6a7f7949420d4b9cd606ca80380d4b8836b6700ebb3a031138a5a06c3c46aacb517e60001b6175b6565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b6ddeca0836040518263ffffffff1660e01b8152600401611fc7919061e9e5565b60206040518083038186803b158015611fdf57600080fd5b505afa158015611ff3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612017919061cf84565b90506120457f823225b88e1d8be44aa625f4591b82efaa5557452235edb8a649b5bf0947050b60001b6175b6565b6120717f8563f60887e0eb19b2e80cd69cda512409bb735cae8ec95db3fe3020b30e7e2c60001b6175b6565b8061210b576120a27fe086c67b268fc789036c861973d6db250ffdc70367540f3c9f0122dd0bb369fe60001b6175b6565b6120ce7f39b14fbd207100409c141b8e627278b2fcafb3af1836acfdb54243552ed9549f60001b6175b6565b6120fa7f535f0afff06828b9a3bad9e681ecfc4ab48e44c7f9f395e2a2927a4223cd4f3260001b6175b6565b600098505050505050505050612820565b6121377f88786ca25b0a6d286633981808c0c81bc6b1afee0e5838ffa4238dc3d18ea54a60001b6175b6565b6121637fb37ba7f94497d003f1e72744d3d47539199e0fb1ce79080ceda18d61dbd29a7b60001b6175b6565b61218f7fd0a6368437b90ac510f422c884b3ba20e52583613c0fee50cae8c888704f521a60001b6175b6565b8961014001511561278e576121c67f9016cb6926b44ffbcc59260b6a43dfbf2dbb9d69f3a713491d4e3b91195218ae60001b6175b6565b6121f27f42b294934692bf93f57fed9519ec7df6e6cf62ce4c9a2d831fca55eab5025e6160001b6175b6565b61221e7f0177dec289b8f2526afd362eec9db354328c162e05fa621b34dee9c4d3af567f60001b6175b6565b8260a001516102a00151158061224b575060008360a001516102c001516040015167ffffffffffffffff16145b156122815761227c7f4dbdf830786377b5a61c71fe8f6dd1a173ffc7bf39e77f2252ad671ac7f2a75a60001b6175b6565b6122ae565b6122ad7f046253b7fae8085fd95fad7084326abf59e3567d74b4954bfd785533b453e65860001b6175b6565b5b6122da7fda4f3826a42bc301260d211b53b113c04a11dbf8e2d44f6fb9f8178af17dfdac60001b6175b6565b6123067f39c0002d2ba1a4cd8083ea52428316a0043c7b24a453dabd0f033b6838cec0b560001b6175b6565b61230e61b0ad565b61233a7fa1dd33eda6387973352feea6a1bcd4086509ba2380bd627442dc858c1c65822960001b6175b6565b6123667fd906e0e68b5ec7fb4600d02287d875391838731415febbd92221a9251fffb32960001b6175b6565b8360a001516102c0015181600001819052506123a47f116f6de949e83736e71fa0cd1b626c5f5a0c492b17b4d2df41efce8490639ff860001b6175b6565b6123d07f536794cbd9b840eb0a0422b8b9f269becb308dc1d96f806c95d1caf314dba8af60001b6175b6565b8560a0015181602001819052506124097fb7e66e972fc516e98db148ac6b928dfa155a046c596910f4ccebd7b1e9034b6460001b6175b6565b6124357f09d0637848afd1783ac77a56716cf0d8d54d334e6623790ddd398e17d212790360001b6175b6565b6000875111156125335761246b7fbe44f4e8ee306bb4efc4b36e17b6ea80aa59bc72308089baf49a8d575007944b60001b6175b6565b6124977f464afa7cc4a013c18352eac9c6ee3776276ead8d95c7070cf47c91b0ff53636d60001b6175b6565b6124c37f505e1a8e4d383caa6cbe518fb40f3aaa80abf4a5e6fa027cdacdd0312e247dc960001b6175b6565b866000815181106124fd577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516000015163ffffffff16816040019067ffffffffffffffff16908167ffffffffffffffff1681525050612560565b61255f7f43d49191d7a47b166ac4dfcabb8850de88eebc0baa14ec066293818a985bf47760001b6175b6565b5b61258c7f5d4026ae0d46f72e9b0d9ea63064be789031a8a4fb19a2e33a3961e39e10787a60001b6175b6565b6125b87fbd4f89b98aea14b10997bfe2b2489e4a110dc2a339161222e454f9b60f9e042460001b6175b6565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e19aeff836040518263ffffffff1660e01b8152600401612615919061e9c3565b60206040518083038186803b15801561262d57600080fd5b505afa158015612641573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612665919061cf84565b90506126937f95022b7fc316ef2c8520d78abbf1367f539b72bb19f83b60b711f79c3a96c7e960001b6175b6565b6126bf7fada2b0058b797e953714e1501f65a7084987774e4c2c5ea20999a752017c3aa260001b6175b6565b8061275b576126f07f2ce826fd2805531f69a6f088261036d17b686419a474b4474a9fb741d1b6f88d60001b6175b6565b61271c7f3a2a0611813370ab03dd50c6d884d5c10c0d34fcb44132043eb846e20b88aa7160001b6175b6565b6127487fe1942f2043a44bee3a7b91e647909cf54f99a7303b380ff65250c2ec209c106e60001b6175b6565b60009a5050505050505050505050612820565b6127877f2da9621272571da9356bf8c908f919ee9ea5d8497e4dab18dbc910bb62e713a360001b6175b6565b50506127bb565b6127ba7fa6520bf8905555f96f501407c95e497772cc9b950fdad2d1629804f22d26683b60001b6175b6565b5b6127e77f7bf28374ed8a80f7bf6e499592c5096014e36c4a5b5039f3be1b6986164e42f460001b6175b6565b6128137f58ef7815c29ece677717aeae48e3d185f910ebf43b34076b54074b904a047ca460001b6175b6565b6001985050505050505050505b92915050565b6128527fceaf56f33254053bf278b675d1cdb7e7b3d58269dcfc916e4d9a49e62ac2bd8c60001b6175b6565b61287e7f90f9050d14f8dcc11d4ee3cb5f447148f594f2411efbaf44857620115f298b0260001b6175b6565b6128aa7f94cbbce47d25e2c7c87588dc7f201e3225745f49b257c3809ac301ae24340c8660001b6175b6565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166327ab443483836040518363ffffffff1660e01b815260040161290792919061e599565b600060405180830381600087803b15801561292157600080fd5b505af1158015612935573d6000803e3d6000fd5b505050505050565b6129697f98b0b781bba2aa9d3b71f9788b667eb14e43d5c97ed7cfe459aa736c1f5d5e6760001b6175b6565b6129957f7ab413b54fe1641cce7c27fcd40db3fdd5b77d0ad21460ffbb64037bb214094d60001b6175b6565b6129c17fd28446571037ce54ef250496f56e9e658b8c7c7a6a12976d8ec18a0f057477ff60001b6175b6565b80600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002081905550505050565b612a5b7f44f05e1b273905dbca6a070144e96243b5fd998fda83c774b9552da4e555751760001b6175b6565b612a877fe6bb7edfe12aa96440400836183eb37c83ba965fb35d95f1ff7c365e0a0b760e60001b6175b6565b612ab37f68fc95c78e0cd8a47c6ccbeaa0ea7a6fc1029ad7b9587e13682bf19960af81d060001b6175b6565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166352e06146826040518263ffffffff1660e01b8152600401612b0e919061e577565b600060405180830381600087803b158015612b2857600080fd5b505af1158015612b3c573d6000803e3d6000fd5b5050505050565b612b6f7f7f83ee04645312cf87972f31aec4d93a3a66b2e28b60f5e7576ebcabc48ad5d660001b6175b6565b612b9b7fcddd69650b0b0f84c96262441ce7e3062c7a59cd4ee4cad02f9b8d42e3e74a5d60001b6175b6565b612bc77f2d313e4e7b939ddab0d139a4153669e99a187ea4a88b5a049faf69714bf2ea6b60001b6175b6565b600081600001519050612bfc7fe3f1b6837b84f034321e86b6a2561ddc332227a505de6d025b7d08562f6e93c960001b6175b6565b612c287f69c5886ae649ec666ae670ad3695eba88beeb2e7c8be7150b00318d5dedda96560001b6175b6565b600082602001519050612c5d7ffc69c08df168fe45cc8bfe55b036f685ac76948277b96d2ee634035a9a56976460001b6175b6565b612c897f9bc5fc0d41cbbe637f5db991d2f9d463643190e7ea26e2e2057cb78b77e6576460001b6175b6565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633778febe846040518263ffffffff1660e01b8152600401612ce6919061e4c6565b60e06040518083038186803b158015612cfe57600080fd5b505afa158015612d12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d36919061d2a2565b9050612d647f92607830c0537e5917bd71c4c9d319a6bcb974ac6f7d0f6e4e043a32b0b0a10860001b6175b6565b612d907f96d0410b004b6d287622cfc6498b48457e9abf320f2ec54cb945bbe9b3a5a38460001b6175b6565b42816080015167ffffffffffffffff161015612e3057612dd27f2206c0c67469616760a2c6b79525c7bc9f3eb7649404b6cbb711fa733dac8f1560001b6175b6565b612dfe7f91fb238c6c32bcb0ca62b7038d4aa9506dfbff4e800a5a80a8a06be69247956560001b6175b6565b6040517f26841a7e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612e5c7f685a3098d28d5e6aea1aad476c3b918395404a48fbe6a2cbfb70dbac4ad65d8560001b6175b6565b612e887f22df9e600b94769026c7c26fa6efdfa71c08457e972e9dce8c9904c8d74d2f9f60001b6175b6565b612eb47ff4f388e2e5ca21fb41f1cf62966b93542fac989ead09b54f0fb75e0f00e6740560001b6175b6565b60008267ffffffffffffffff161415612f5157612ef37f2168f4ad4ec31a4a24c4ee020f43222cee5da4e81fe52efb8f5f505053e99bb460001b6175b6565b612f1f7f8df3530a42ab5391f144ab40f00833dda660b17504ed139c575ce968619164e860001b6175b6565b6040517f26841a7e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612f7d7f0b55ede9a25e214eff982b82950e8794da059d861f525b4146d36c0a025581b560001b6175b6565b612fa97f05eff38d788bbfc4f744769421c28568bf9019ac5cd77b929a40767234ac0bad60001b6175b6565b612fd57fc71aea22379865bc3e4d8f805cc058f9b5b2280081547230dbd57391f634bfe260001b6175b6565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d7866040518263ffffffff1660e01b8152600401613032919061e950565b60006040518083038186803b15801561304a57600080fd5b505afa15801561305e573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190613087919061d335565b90506130b57f3ddb6da24c09a6ec0cd5bea169f3020e7a024a1522607c2a595208a71d914e6d60001b6175b6565b6130e17f3185ab22d7a756c409c1fe36dba1302ae619a41889aa45fe1056aa4dd532dbf660001b6175b6565b600081610100015167ffffffffffffffff161415613183576131257ffc0e9e2cecba13d9501baa9e074f63710e36de5d315099a1326ff80b329e39ab60001b6175b6565b6131517f50b9fc14544e4153a7ac873fd958cf5d8add32ad7b334f2ec086e13da384e94660001b6175b6565b6040517f26841a7e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6131af7fbf200f643997f0c31a632b98b0ec1c9288755ecb59349ad74ac5588b850d8aa460001b6175b6565b6131db7f40ff5548ab414950c9640130aa0e5b8d4ecc24b2d07e2c72ba25391d06096e0b60001b6175b6565b6132077ff283d518ce9e8031d47e7655dbe4863c200f9538974a785d2a14b0cfa652462360001b6175b6565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c5c81b2083608001516040518263ffffffff1660e01b8152600401613267919061e691565b6101606040518083038186803b15801561328057600080fd5b505afa158015613294573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132b8919061d533565b90506132e67fa1e1d65e6636c12311cc600108ac74693390ff9aac778798bfc354d4309102c860001b6175b6565b6133127ff3b77d4b43e3d2c3ad71273cd3c9a24a39b93b2897f298caf5e9592a4ff0aa2e60001b6175b6565b60004390506133437f944fffa1ebf48a4c5bb803bc658702a572d56e4dc6cfc39fa35affaba096a8ff60001b6175b6565b61336f7fb737fb5d11ff70fdf1ad2da8a362c1851f517e46ac0064e8153b2529fd8f5ca060001b6175b6565b808260c0015167ffffffffffffffff168460c0015161338e919061ed04565b101561341e576133c07f179670ddfbee6ba3325a8d6ac4d1aa6d4e3f56606172543bce964e4477333f4260001b6175b6565b6133ec7f44d259532cf45dc9ccbe0cceb2e2653206891b795e316f6db6615c125580612160001b6175b6565b6040517f26841a7e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61344a7f40464701613f5ca06f802831aeba536deaef91e53b98d135a26b20fc71e0a48960001b6175b6565b6134767f42a8b0a4c0fbb45241871705740a2a702f033b564189e16915339e6a399f6b6060001b6175b6565b6134a27fd275554609bb8ce56f34a6d79e653e9dc6784fccf08ed5d7c77766f738dc83f460001b6175b6565b60006134ae87876191ca565b90506134dc7fc3601494002ee296670f7ae30d6e5338fd25ab349df274a63719b71a7918fa8860001b6175b6565b6135087ff614dcb71a494be5360271ee9ebe88ba2e28dfbc5538781e07f0894bf79bf39c60001b6175b6565b6000613516858584866192bd565b90506135447f1103d7a072057aa8f8004cc6739e544342e5170f8d175bd7a75f153840876c0260001b6175b6565b6135707f8587c199af6ea0dcb78521edd933b0c373c4daf2ad8740a405b095c609c1a6b260001b6175b6565b60008167ffffffffffffffff16141561360d576135af7fa269072e18852bc6aba264dacd7a0d8cd661949e390ccf77562ce2610948161f60001b6175b6565b6135db7f5bab5d4d9591d77fe281fab59ff5465c36d8ba8359494ca185cad34fc778ba8060001b6175b6565b6040517f26841a7e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6136397f181536a9806f586125e649c2dce46bd99f2db65f32238582f48133520879bfe860001b6175b6565b6136657f57db23d9080db44dc07efe2216df060dec3879fead9f30918badcbf1be33066a60001b6175b6565b6136917f589b8d6d13b6aaf5cd9cb2add9ba35149f2ce0addea2ba7cf6aaeed13e4f38e560001b6175b6565b61369d85878684618b3d565b505050505050505050565b6136d47f9baedcc545482aba9a35de05fb02cf25157c53ace06a6fdbae9240f9bbd34cfe60001b6175b6565b6137007f66deb016a2c7340a2b93294377bb4375d8f055d9c204177cf91cee389b61f32960001b6175b6565b61372c7f2a26c5a0bf9d2b45d0aa4f534be4c04c0f27377969913c6c1e027bc02f8fabaa60001b6175b6565b60006137398685846199ab565b90506137677f17a349f6c508bd849237d6ca6c6406ade0e014bfb2e4bb7d7d5839b50c234b4860001b6175b6565b6137937fcdc731413ef0d99abf34ba6f20e0a1b8a910773206c7dc5557bd17b468716cd460001b6175b6565b8067ffffffffffffffff1686610140015167ffffffffffffffff16101561384a576137e07fbf8e04e1325a21512d5f274c9677a7ac6457f1d15a44e355a7420ca13585707060001b6175b6565b61380c7f31102aa8ccd8667514b8fd08c7636c0a1e602296f20bd0460c7cafe949184ac060001b6175b6565b60096040517fcd442d65000000000000000000000000000000000000000000000000000000008152600401613841919061e784565b60405180910390fd5b6138767f4aa71b8fadb34ae7c55156a583bae72d4c7c20aaa1fd643567645d2c3695bae560001b6175b6565b6138a27fdb154f979c4973495fe8d10f94dc3be44542e551b234457a7de73c79343fe1e660001b6175b6565b6138ce7f076b5c095a55d80c489ebef76b914ea971c66582178637fe4793adf40ef62c1360001b6175b6565b80856060018181516138e0919061ed5a565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506139297fc61d5295071b35e0002f8bb10afbcdeda61acbdc2c3ba3544be9b9788b5e57e760001b6175b6565b6139557f01c59d689cd4736da2ce845eac2bec1e84ccf0a0b6b714636c00f8470894978760001b6175b6565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c819d86c866040518263ffffffff1660e01b81526004016139b0919061e89f565b600060405180830381600087803b1580156139ca57600080fd5b505af11580156139de573d6000803e3d6000fd5b50505050613a0e7f2aac02d3a10e51b456eb080c2e57e6a569ab7e451f3eb30a64ff028c82e19b8160001b6175b6565b613a3a7f1e337c5ab7476da9e627b4c9a65a7a2d733166bf7bcd89094c82855395909fad60001b6175b6565b808661014001818151613a4d919061ee3f565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050613a967fb60cfb854ade38bfa3c6d151864b7e8d4c4c4a920bb011816fa1f14a13172fa760001b6175b6565b613ac27fea95d86ac872c4f42c55813b9986b748f6cfe30b725af4bcc3d26676eab08ff560001b6175b6565b6000866101c0019015159081151581525050613b007faf68160529ea7005d5ab6186a73efbd7a876521b23c1e9852bde0728313dbfc860001b6175b6565b613b2c7f31056821ae0ab0bea78d33444ea6f3f9230acfabf905d5eb56b366f7c83a063160001b6175b6565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663681850d7876040518263ffffffff1660e01b8152600401613b87919061e81d565b600060405180830381600087803b158015613ba157600080fd5b505af1158015613bb5573d6000803e3d6000fd5b50505050613be57f90d7a0451ee7b3c0db2cc10dcb87a8055ea8e5a60cbf6245a7e8fd7ab5c5064360001b6175b6565b613c117fd26c94a085f9341efe607f7a861d71e012f78eb068ee1aa365edf196a8dedf0660001b6175b6565b6000613c3f7f466de43131e5884c78e632af312cc7f6133820eac22bc31ba9989514f4e898d260001b6175b6565b613c6b7fab1e7ac727887619d8a3dddb863dc6da34ae49dde914d2de9acbe597862e834660001b6175b6565b60005b8451811015613dc457613ca37ff958b70cb5124380aa154e8cde8ecc6081c99c5de7c75d3334c078d3e2843af660001b6175b6565b613ccf7f9dc98a1902d59d0e17be140d5cae13f2a378b9cce070a5597f2b26a43bfaa59960001b6175b6565b848181518110613d08577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516080015115613d8457613d457fd958756a3f9a192ac71527ed38462281cee0509ede955ebbb6b505c69b7f477360001b6175b6565b613d717fbeaadde72cea2776801bbfa43017f34e94c5c91c64c5d2c27b3a7a9aea3f688f60001b6175b6565b8180613d7c9061f134565b925050613db1565b613db07f51bc3cd067ba5c32b7c0e16495ff35860b3da1e6fcad4dcb8a6d961fe7421f3460001b6175b6565b5b8080613dbc9061f0eb565b915050613c6e565b50613df17f871ba2862f5f0f6b4f318110d643505cc09631946bd23703fda9e49396e7be0660001b6175b6565b613e1d7ff5479c4c07d735ee38f08b86bf5600a4acc5c33f308e4f4df39434bce32ae38560001b6175b6565b60018167ffffffffffffffff161415613f4c57613e5c7f2f0765d1a8599ad4c6b1339a5deba4740dfaae208c0417da9877650c77f4de3b60001b6175b6565b613e887f1097c5bdb8378cea33bce60eab4e20c80e72a30cf3b48ef3b287d13d17220a0060001b6175b6565b613eb47f97e5d56dfefdf5d478a924d5063e97860ee08277dfda1c819725575dcceeea5e60001b6175b6565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639cd103e988600060016040518463ffffffff1660e01b8152600401613f159392919061e83f565b600060405180830381600087803b158015613f2f57600080fd5b505af1158015613f43573d6000803e3d6000fd5b50505050613f79565b613f787fa5229b4b1951fd2f97b96136b8dc82c9a23d400ae22fd280572ee5b22826250460001b6175b6565b5b613fa57f991cda3e4a8a87eaf36d615f25ca1fae390498836879bb35928ea3749721564160001b6175b6565b613fd17ff24062659cff4a7f2249a0774f377856aa16a6e6cf3affa6c0ae45f6e80200e860001b6175b6565b6001876101200151613fe3919061ed5a565b67ffffffffffffffff168167ffffffffffffffff16141561429a5761402a7f1dd66f97cc736107512f4dc7c857ea715189193e8773cd04845884b4b200045060001b6175b6565b6140567f4af059b9b5584c4b47afd33abbb21ac42897c5337c111095ff194bca8c8d2c6560001b6175b6565b6140827f66e1ad802da42798aaab0c16856cfad847c2af03a0a042fb55843d533ed397d160001b6175b6565b600087610140015167ffffffffffffffff16111561417d576140c67f5c7ebd9f13ed692d8bfa45e5b5fe2bb8b96f5e7a2de0c9964221fa772e79c25460001b6175b6565b6140f27fa02ce339ec7824036cd4f92a803f864e8d00b6012efcde8530f9b4dade32859660001b6175b6565b61411e7f4e4e55d54ff34e97976379a3b63ea0a37c516cfb611d77a44a4302d4520661e060001b6175b6565b866020015173ffffffffffffffffffffffffffffffffffffffff166108fc88610140015167ffffffffffffffff169081150290604051600060405180830381858888f19350505050158015614177573d6000803e3d6000fd5b506141aa565b6141a97fc4606c0d1a325221799bd3893b6b65409de49f58959fc53f74138c6c107ba78a60001b6175b6565b5b6141d67f2d335617b01d33ca1c43261649dafea7bdf7d55a6e59c281c92c28b369fe2c1c60001b6175b6565b6142027ffb85d4ed3488c5279346824de02dc6979f3f3ba5b4fb93ffb2ddb0e89a2efb7960001b6175b6565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639cd103e988600160006040518463ffffffff1660e01b81526004016142639392919061e83f565b600060405180830381600087803b15801561427d57600080fd5b505af1158015614291573d6000803e3d6000fd5b505050506143b6565b6142c67fbc27faa4ad6d6257a265a86e66ea9e498d06d40ddc64abdb4cef1701a82e650560001b6175b6565b6142f27fe6438c61e9fc16dee939bc612f7744982fbbec328c27e65f276997976a37f95a60001b6175b6565b61431e7f654b76213247bef1dd5b384d007295b3faa419cf820b13e2924736ff9863cca560001b6175b6565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166334fddaac886020015189600001516040518363ffffffff1660e01b815260040161438392919061e4e1565b600060405180830381600087803b15801561439d57600080fd5b505af11580156143b1573d6000803e3d6000fd5b505050505b6143e27f31abf154acfa9c890b7b44f2730b36b87c8ff5a02f3888b08c0b3b922c4bab0960001b6175b6565b61440e7f57d9467eb25d53b550c70f02612c5f8f22babf3b9e58a5803dd49e46941f799d60001b6175b6565b7f123b9998b2f027b02db4cb2eadef421fe9f1c21627569438588262d3a6baac6c6006438860a0015185604051614448949392919061e600565b60405180910390a150505050505050565b6144857f9f15293e9a6c35ea9f90871ef178fb9ca3d5b6931d7bd618b41f74e239ad082b60001b6175b6565b6144b17f589933865349507e7380f7e213f9f30f42484050b4d256381f5e5c839776557660001b6175b6565b6144dd7f5b3ca1a9c68f9de4acbae0648b48b9f5e1d6880b93e1ec2675b27a07ace6d08e60001b6175b6565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561454757600080fd5b505afa15801561455b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061457f919061d533565b90506145ad7fd0c835d251d0ee3f7ceb4da4ef869e6df9ad3abb5a09faafbe9c3296ca52e2b060001b6175b6565b6145d97f3e78f9f2a23678d63b6c64c9de871e4f66bc66dfbfe8c73d78203c7825e70ba260001b6175b6565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663defce5d484600001516040518263ffffffff1660e01b815260040161463a919061e577565b60006040518083038186803b15801561465257600080fd5b505afa158015614666573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061468f919061d175565b90506146bd7fc9d38e5985cb4b0fea15b9d62cf78ddee17888b0a6bda23d23fcb95c8a99118760001b6175b6565b6146e97f1c3bf8c3e67148b51ec013d7d81210ba48acccfb2b1c603037f8fdee673cd7a660001b6175b6565b806102a001511561486e576147207fcd9594440d4e163f22f13f3972ab9d6328eafa7fa23476ca3854d94b50ebdd0f60001b6175b6565b61474c7fc870ad9b75503fc72cbfec939a3b6ae68ab9c8c1c7598e448a31f27a8044182e60001b6175b6565b6147787f8a45b15d4d8babf9fd9e80c2cc8fb836eb245b344dca69a6826be4963be0d33c60001b6175b6565b806020015173ffffffffffffffffffffffffffffffffffffffff16836060015173ffffffffffffffffffffffffffffffffffffffff161461483d576147df7f683313b696049293fc6c6b1c4297f779e786609a79daf51767f20cb7922726f360001b6175b6565b61480b7fe1d904c805d4ff324f9d484038d676397ccaa45f5e5e0df399376e35683428c260001b6175b6565b6040517f65fd380a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6148697ff306c205058aa978db4d3966690550568bf446be55d798adcb65452a5ea507c260001b6175b6565b614ff0565b61489a7fcebb75de18aa010981b585f795b05852cb050137802cb69c20ec371b2351180b60001b6175b6565b6148c67f30f341aa46b3e12b029520e0db2e3a2231f31561e321e445d7d4e1d40951ba2460001b6175b6565b6148f27fcf2b1661f50de81f59c91894a295c6837a39d366e38e3e2189521ec2cd4a4aeb60001b6175b6565b60006149207fff4d661d5268cdb7260ffc0dbc966ffc06e1a10112204143030d29103522442660001b6175b6565b61494c7fe24ea467fb39d277185f0a424c3c71ba635051c648ec96637566c4529095352260001b6175b6565b600082610220015190506149827f44c979fb4e60cb0de9c1994a9e81606ef77a8dbee722dbeed8cccd7d6b8109a860001b6175b6565b6149ae7f9414cae6ceecc5fd78a89994212b4d6c8f34d4386d647786e46a3be3b500d1fa60001b6175b6565b60005b8151811015614b82576149e67f0c361c218158c8d2b73d3816da70d49f466294ba77419f99dccc5956dc5cbcde60001b6175b6565b614a127fc1232b7e96ca0534a507996ea8cb5048071f2e945ae2fb2d2c1f515772acb25f60001b6175b6565b856060015173ffffffffffffffffffffffffffffffffffffffff16828281518110614a66577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff161415614b4357614ab67f9d045721b08572e7d92d628a194e0e06c2358f401972582ba8025850907c412c60001b6175b6565b614ae27f0ce17d79b13b854c273ffef87ab9ad7f2aa2fcfcdf4d7b7fc553102516c25ed260001b6175b6565b614b0e7fba7697b770f65ad986472bf7072c622834bc06dc11d0feb75abaa1873af5d52a60001b6175b6565b60019250614b3e7f082905da3c37fe86c1eccfcb172bc1bf0b5e6ce44b52324aa40dbb2fb6d2d96b60001b6175b6565b614b82565b614b6f7ffe8efcbff7a0a06b810552ba674d0415e7c592fdc7be5fb4c6d010a4a6507e6860001b6175b6565b8080614b7a9061f0eb565b9150506149b1565b50614baf7fdccb0e68ace00cfdcb4c5a4dc4051d900323e692df92d7d5e58320bd579e264060001b6175b6565b614bdb7fc3bae6467366736a9a874902bb464ac59f0fa6ff8b34ddadfdc0f7d8729ce35460001b6175b6565b81614ea157614c0c7f1fac2d8b5ee5868dc923de1db474a25d94e018635a5ed2894938b03ff524bb3560001b6175b6565b614c387f62090d4552c638b5ffe521ae23d5130e16e82caa213abfa6a68aa1bb26d7674360001b6175b6565b614c647f3fab7c5ac8d46aabda47fafadcbcccd69e160f675b81db3db3c3974da2303a4460001b6175b6565b60008361024001519050614c9a7f41d97894412044404a20f39774f608a168adbb99100fd00d565e12b85fe3d67160001b6175b6565b614cc67f0ac9e0b9d0af7f7e46d7162f61a86c3cf7ad2ef691ad903b07a7a911da43bb0460001b6175b6565b60005b8151811015614e9a57614cfe7fb926ec24d49fb89053fa162dcc50d4fcba816a4b634cd5dee2305e268a1c3eff60001b6175b6565b614d2a7fefcde9498c8724dda2c5f556c052264bcba51ea17c56b555097993bb4775777360001b6175b6565b866060015173ffffffffffffffffffffffffffffffffffffffff16828281518110614d7e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff161415614e5b57614dce7fb2e290160be074801c19c1552bff088d9a0b5fab93871e58a9334c677f68129860001b6175b6565b614dfa7f886bc391879a533ee06736c769953449f3042a660846014cd3f48695135e126c60001b6175b6565b614e267f4ff1c17bb4b150dc5acf85550688e023bbdfdcb0e00d421104cae0e63f6863aa60001b6175b6565b60019350614e567f9227955e4b20c42d52c5dc855513238d49291b4f102f02756c8a5c15534b601b60001b6175b6565b614e9a565b614e877f7cb81dbe2d5948593051969d3e7a62ca61406f6bac11df2e09da406814d3894560001b6175b6565b8080614e929061f0eb565b915050614cc9565b5050614ece565b614ecd7f871cb5821c1dedad8f51a98296421479fd74b4243e5b003453591aad82160cba60001b6175b6565b5b614efa7f251a9e690415eaaaf8433725dc7fa83ed7cbe9717112c41953d4f96b162946a360001b6175b6565b614f267f44cf5e2e81f52ac91a00e95dcb11a6b4122d258a061723764cb6b0a694e93d3f60001b6175b6565b81614fc157614f577f262fe490489396783490abe4ca27dc44e9e49d816b29d532a83113190fb15ca560001b6175b6565b614f837fd5dc6d711d4c884ffa15ee617444de1216774dfc2a5c938ba842959b87d4951760001b6175b6565b60016040517fcd442d65000000000000000000000000000000000000000000000000000000008152600401614fb8919061e6ac565b60405180910390fd5b614fed7fef7bcaa5af352e4d6dc12ac744771d588cb10647212ecd02229cf173e055ddf960001b6175b6565b50505b61501c7ffeea2c465f317dc149e28a03df67f09a68214a6bdb9f943788051b8d659134d060001b6175b6565b6150487fc4523bb5f995be9ed6cfee66541efd8bcd3f59856f6261b2e55efbb2a636d22f60001b6175b6565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a65374a85606001516040518263ffffffff1660e01b81526004016150a9919061e4c6565b60e06040518083038186803b1580156150c157600080fd5b505afa1580156150d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906150f9919061d2a2565b90506151277f230d1f70430ca9609ddc05bdc42b9792ae093938b3739bf54422b10f0867166360001b6175b6565b6151537f74c3d770a3b86cde06533477017beeb435151cbe91ac9fd90c0887bf559f6fbf60001b6175b6565b60006151628360000151619ba6565b90506151907fd8aa09ac14dbc8811fdb4f71620ee7774126d15351f1d0bf352d8bb6a5a65f4e60001b6175b6565b6151bc7f926471e540c28afada3a1e1e2fe0273f3e77aed9185d7442427ad57432c072fe60001b6175b6565b60008560a0015167ffffffffffffffff16141580156151df575082610100015143105b15615423576152107f80ec2b35fdf062733d4f5d7079ab912640207ba30e7a515074bfdb609f00dcfd60001b6175b6565b61523c7fc0176102031847a4daf4a2c602db4477e7481f8f70b9cae6f6683c6bf4188fcd60001b6175b6565b6152687fa74ccab719c4e9cc1a0cfe63021f4d9ff6b2a5b7c1e0300a59792551dee72f8360001b6175b6565b60005b815181101561541d576152a07f20409ff02f453ad1a58b7f805730694296712e8f2827fc19d40e91613f7104a760001b6175b6565b6152cc7f37ea62e350d2fb3532ff2564210d529f90b4db8beb1ebf796e1eed4e9d7fa0f260001b6175b6565b856060015173ffffffffffffffffffffffffffffffffffffffff16828281518110615320577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff1614156153de576153747f176898c460b3653e43230efa8683b44115aac047c1c5aabf93f77de45029e48f60001b6175b6565b6153a07f324037464d18e274dff73c06e6bc72926d5fe4ca20d2925c3fd33e6850a78cc360001b6175b6565b60026040517fcd442d650000000000000000000000000000000000000000000000000000000081526004016153d5919061e6c7565b60405180910390fd5b61540a7f2dea8c9a60be242bee62df2d33b0dab3f89e5ad73c5d76d93873a2d83eac013960001b6175b6565b80806154159061f0eb565b91505061526b565b50615450565b61544f7f479df34880a1fd20ab1af66f1ff690d0ce48c709a1837e8515ed785ca985ecdb60001b6175b6565b5b61547c7f2e9a8ab34dee140963d9e5821c912490ee795be903345a8b310d4d90a95537ad60001b6175b6565b6154a87fa4bc3eb0e48088650730d7d502ca90150139a4c86f805151a97f920e3dcabd3360001b6175b6565b60006154b48685619f32565b90506154e27f19763910d7e9583697c3d0e8a46f726e9592044ab98ffff392c06593eb84906960001b6175b6565b61550e7f8e8d4d4b30912198c3781dd3ecaf37c7f52be6f2bbf446a2957fbb3368b5876960001b6175b6565b806155a95761553f7f0723cc4e39dc35297b3fc4e3775218d68db78558fea2d6b747fbdbd67bcbdaa860001b6175b6565b61556b7fe92d6d2e11b6d0568a6fc2248be232647afc74466f6b81c8093c0bb808c80d8760001b6175b6565b60036040517fcd442d650000000000000000000000000000000000000000000000000000000081526004016155a0919061e6e2565b60405180910390fd5b6155d57f1706435028b52c17b44ad9c4a6489856ffbacf649c8b5ceff9eb18ed1a87413a60001b6175b6565b6156017fb0f3d7bc581d1a2344029e9a8bd62a0ee8f290a15b0a09cea7c4c34dcd9cac6960001b6175b6565b61562d7f4780f44050cc954668bfc97e8e2ae4984703e1e2fd4702be3a6df1ff9705a8da60001b6175b6565b600061565b7f61f5af4c98450ecbf5ffc6347438b5d02b0b3db091ec7ccac0bd76e89b7be50060001b6175b6565b6156877f34d8c8c4c37aee73106d21c387272bf7929b0289a853e002c499b2bc3f65c1af60001b6175b6565b60006156b57f9932ebe51c548b548b943633c3b121ac3a856039de911142608c8af43461e08560001b6175b6565b6156e17f49665c6dfca803761a1357f8506a1d5bd684c4aa1b5286565216745992b04f4760001b6175b6565b600061570f7f06187d03ce702695b16aad7cff5d7f999433a65c657e56950d001e2daddcd60e60001b6175b6565b61573b7fc98f94b89d4cc4826f7684ed7598b49781b3c497b525686b044ff2e5b2f5a9e160001b6175b6565b61574361b0de565b61576f7f3cd67a62a79f05486462372a9239accbda82219c9fb2ad0912786108dab06cc060001b6175b6565b61579b7f8a0826b53edc1fdd2ce85ce2ed3d4cf483374cc61f25aae27c2cec181801d26560001b6175b6565b600088610100015190506157d17f6303de697d9156c419c1489b715b4e82b0e3b428da790a61e34f1e9f9a59baa460001b6175b6565b6157fd7f1e3198a74a3ce219499bc00caf440ec0970d7512ec5292b445f290fbd18ca2c760001b6175b6565b60005b8751811015615ef4576158357f134fe93911582a8bccb0bcf9d80ae9ff988a0a07806d3d9284ed8318e74ee91f60001b6175b6565b6158617fb59ad2332cd24d7562b004c015227e0d9adf2d1c6f6b135be1de28ad8fc59ec860001b6175b6565b8b6060015173ffffffffffffffffffffffffffffffffffffffff168882815181106158b5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff161415615eb5576159097ff586f0b7cc853713edf2af317ee6ff4466739d5729148d9c7a0e85a8324ed58760001b6175b6565b6159357f261fe10024cab039818aa23dff4d0b3dafa9eae6cdae26fbe173c99515f6778460001b6175b6565b6159617f256a690e3645f0deda8073c23eb27863f93abdc294d49e2e63bab9627d63a3f660001b6175b6565b826040015193506159947f67dcd611e82a480e6e75fd334044181dfb19ab234cfdb88e0c16152ce4905cf460001b6175b6565b6159c07fde67596ab778995cacb6630cdb733cbbbdb0c7c964b3fbe700fc0f7b54392d6760001b6175b6565b8960e0015167ffffffffffffffff168467ffffffffffffffff1614806159e557508143115b15615ae057615a167f9a4a028ca52c58a162d98ba4c5d2608e95caf097bc68480aa1f5da6320ae2e1660001b6175b6565b615a427f7a815d78cee0956de00045d54abef3cc27c8ed72bbc13fdf94227989e980531860001b6175b6565b615a6e7f3ecacd23fe54882d40fa5937ce24fd1d01ccd2bbf8fb9a7ac5ee3f374f80c78f60001b6175b6565b6001836080019015159081151581525050615aab7f545754fcaf2fc06d1b272d3562a22b5cff9ca9d0050b79dbd8f58e3004a21f7260001b6175b6565b615ad77f35b954cc46436bb747de5aea1cdd8e296a71e118fbfa02f8a1f8644b7dfd65b360001b6175b6565b60019450615b0d565b615b0c7f5109b6adb6b35a7e5890dc0f20a830b7ac928704b551d938de6aa9a38984295960001b6175b6565b5b615b397f086066136e2e21292269d7762d7dcc92987524a4d4cbf562bd4c1b85ae21331f60001b6175b6565b615b657f900f94bac40378819bbe558c7b3c094dcca699364f667ebccdb589ee2d440baa60001b6175b6565b8960e0015167ffffffffffffffff168467ffffffffffffffff161115615c1b57615bb17fe80aabf202c9c6a837f43581a2f184b86629f9a3ac317f4453d2680843bc530960001b6175b6565b615bdd7f9562a3e74de9238548aad023c4eb65aff30d6c2314b6fe2fd5270387462bf69560001b6175b6565b60046040517fcd442d65000000000000000000000000000000000000000000000000000000008152600401615c12919061e6fd565b60405180910390fd5b615c477f33bdd2b44eb750539360062a6533375beb0a72a68b79b2c8796669490ad6551760001b6175b6565b615c737f86aaa7b2817ca740c6acd11e94dcfc6358e77187bab6fbf914dbd918982f4f9b60001b6175b6565b615c9f7f800af83943b3b34fae845ba70c9cbc39dffaff57163921b32d9810b1de43fc9360001b6175b6565b6000615caf8b610100015161ab52565b9050615cdd7f825bff47ef7f2e534cde19945792a66eee9ce85027e2188c1bf04c9e05fbe6ed60001b6175b6565b615d097fb8b87c67d887ce6c05ad0d4dc05682b50234589ccf5689c26c94d626935f0d6c60001b6175b6565b8015615da557615d3b7f20f8cca40902901329c000b1522979d5ad01648863f1bf1f742c205997b2cec960001b6175b6565b615d677f3b1959f19399c4d10fc750cb40ba3d06e6372b383380df8504e405f77e5d625160001b6175b6565b60056040517fcd442d65000000000000000000000000000000000000000000000000000000008152600401615d9c919061e718565b60405180910390fd5b615dd17f39bac7b45c09dafcee7290534e7b5a91fc2bc2dc10410064e12d69f35362704e60001b6175b6565b615dfd7fd7437c3e322f58fd49b6694f273d92ddee87330b1acbc512b3146c4a4ebcccea60001b6175b6565b836040018051809190615e0f9061f134565b67ffffffffffffffff1667ffffffffffffffff1681525050615e537f1f2654316c0e72836cc67be4e4031fd36b85cd3fd414293a0f16dd2c64affb2760001b6175b6565b615e7f7fde50e8648b99fff0ec46215f84a6771c2f2f4d83a77ef95bb2b63d9eb45ea36160001b6175b6565b60019650615eaf7fe34a05149548ea98f8f4b69d599304660966e2ce3aa200f55c8ef5c88381943a60001b6175b6565b50615ef4565b615ee17f30784fee221b931f61ca6a7a18892ca0ce0b818b6437c250ecb6ddeb56c307e060001b6175b6565b8080615eec9061f0eb565b915050615800565b50615f217f36d420d843b50a3dbd6f1bb8da0f962b4912d3c45a4dbb3fcb6191cd5f0b7c2f60001b6175b6565b615f4d7fc31de5ba66255068d43213793fc816ba05e8ac4f6aa03b043de20f64fa9a1dba60001b6175b6565b8461693757615f7e7f85ce4019bdce638780080b717d66b121be20be71a2fc33399f3675e3cf356dc460001b6175b6565b615faa7fba8210e63551a5d97bfdeb12db24baf414555c9498bb29b1355c2835ca42168260001b6175b6565b615fd67f42afaee83070f6e6b1d14a5f558b38efb9c182d1207dd2d6ea76ec66bb76f06f60001b6175b6565b6001896101200151615fe8919061ed5a565b67ffffffffffffffff1687511415616090576160267f01870235c5db73c8f895e4cc93dbe889b76c8e8d7148f1a450f0471141870bac60001b6175b6565b6160527f5f06ef1537656376ec8c3dcb656a9fa0d11d659edd0b9e935e366a2f0b6fd16260001b6175b6565b60066040517fcd442d65000000000000000000000000000000000000000000000000000000008152600401616087919061e733565b60405180910390fd5b6160bc7f7328d26958f584ba4dec7e2c41afe66ab358009e2c1ff9b6c40419b4731c0eda60001b6175b6565b6160e87f45eb456219bde368a3e68f13167140634c051c536efb12d357b97d1ae5dc668060001b6175b6565b6161147fc8d3119369e6c0f80bc42884ec1b50164c4aeadca9dffa05bb5f79fa86d9bf8660001b6175b6565b8860a001518960800151616128919061edc9565b67ffffffffffffffff16886060015167ffffffffffffffff1610156161dd576161737f8981412390b0ff1b0c43c762d1f96c2359b48814fd34ba1d3adaa2071c0b6f9160001b6175b6565b61619f7f5b4dac620cd8a88ee1a7a1a0aa35e65ed1dac63416b137cc60c41b07bf5d177a60001b6175b6565b60076040517fcd442d650000000000000000000000000000000000000000000000000000000081526004016161d4919061e74e565b60405180910390fd5b6162097f06a15976853fa326eb7935ce909bff96648f71361ae82e98aa87363338bbb05a60001b6175b6565b6162357f8b7a98b2b7730b4392003d9225825a7c1c790333a68f6fb7dcf14d124d112be260001b6175b6565b6162617fcf368bdeb36b9f7abdac489588a550cf525db7b5c155aecd6e36cdc981c2de3d60001b6175b6565b8860a001518960800151616275919061edc9565b88606001818151616286919061ee3f565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506162cf7fccbb14dc1e7a4626141ab8ef1c3a1d1360fcd8b8cebaae8d9af71357370b7cdb60001b6175b6565b6162fb7f3b9be9b497d50978ab8698af0d172ef4224019582d9273f8fd658936ebfa335560001b6175b6565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c819d86c896040518263ffffffff1660e01b8152600401616356919061e89f565b600060405180830381600087803b15801561637057600080fd5b505af1158015616384573d6000803e3d6000fd5b505050506163b47fc0ab0ed02b402fe74bcfa4f1261c96746e8ee6221b0eeb2323cd5544d612f8cf60001b6175b6565b6163e07f0fa1563893cb29f57dd78875d08fe16025d12192d168f2e1effa6ec54fce972660001b6175b6565b8760c00151826000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506164487f687186e8987d4042c74c0c51284e3b3efa184c6ea290272af3ad4a00bf513ef760001b6175b6565b6164747f94522384eaeba1b20cf7f7432dc1b2783607fe2e53d356e90843f486ba7aef9060001b6175b6565b8a60600151826020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506164dc7f071810d7a09de2577b0449fc0e48eaa4d6b5a1666d0601aa63f0e644436456b360001b6175b6565b6165087fe6e7d43836ada4b099c7207b567d4e7281c5081511473cb07b01426d5004480460001b6175b6565b6001826040019067ffffffffffffffff16908167ffffffffffffffff16815250506165557fef9ed1af064754850afaba05aac67fb37359a5c7384e9ef7c82641c590248ff660001b6175b6565b6165817fbb59c892ce6240be1cd174def91de2c85a4c5ee5185cb89a43656385d469553d60001b6175b6565b438260600181815250506165b77f31774f40f9289fb9bf2f4be013a8f5473b3376a5d83efea8eb304c3ddd78422760001b6175b6565b6165e37f86d249fd2ec2fd8b6589d7540d10026ca9daa3141811cfdc34bbd6ccffea5a7860001b6175b6565b60008260800190151590811515815250506166207f693064a16f67fac8e362c0e37be42daef980b678c5690712e4d0924d05ac376060001b6175b6565b61664c7fbf3308cd3437f2d1ae760717aea198a80b5d1add63d4f66e270bf8a063b24b6860001b6175b6565b60006001885161665c919061ed04565b67ffffffffffffffff81111561669b577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156166d457816020015b6166c161b0de565b8152602001906001900390816166b95790505b5090506167037f7a6aa92a831aa382e26501da3967318a610f5116e77d33251d51a5d9c55040bb60001b6175b6565b61672f7f244fe727a8d00b822d2992356a30824e07d1530554d9fd6191b04dbb41c32c8660001b6175b6565b60005b885181101561682b576167677f21fb66466133ecfbd3c03b7bb904d6be9666e5fe1bb3c9b663fc83ada757076b60001b6175b6565b6167937f8b9252ad2b10856a728812feb4a7c8e4ece2cbb30d1826fd12cabdff4c9fbaf060001b6175b6565b8881815181106167cc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015182828151811061680d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018190525080806168239061f0eb565b915050616732565b506168587f15fda1e0d5f8b4f06ca5a4c1f6b0d2fe38e68f0fcb85c3badfdba2d1e4b28e5360001b6175b6565b6168847ffe10f30226cb266af72728c1a8721687a76f39ad904cfac3eb01bbf217c0bb6b60001b6175b6565b828160018351616894919061ee0b565b815181106168cb577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052506169027f694cf3db14d4f3b9635cf395b0aca7d4ddda24d4840fe73e9bea8014f2f98fa760001b6175b6565b61692e7f09735a879116d8310d79080c7f29f8dcb9bed2ecc923a55ca9c397d43dd293aa60001b6175b6565b80975050616964565b6169637fc52751fd82b08ba2f98e1d70d7fa473660c97006d83ebd9a0ec42f437019899260001b6175b6565b5b6169907fea7f07299e05e99d4f0c8ed5e6b04e3819f329323368029c3f7a7f3fb28b456d60001b6175b6565b6169bc7faa4f72d5efd49a726272e1ad18f39c140ef05138d1ce8d82c0ddc2731bbd29cc60001b6175b6565b6169ca896000015188612826565b6169f67ff990a32c3ab92db80d9efa4364b8489570b0a12b22e0b067bf3cf7cb570c680360001b6175b6565b616a227f08d9836fc245804ca925acc004290fa638ea77489d8a42a757ac76fd100f169860001b6175b6565b846170c257616a537ff03791d934798f7081f4c8abb8077558708f2bf63ff5c4d7bd7b705245bfe20c60001b6175b6565b616a7f7fde9b7ae06056dc0217bd8ffa7215c81370c85ea57aa40ab478dff090dde2a34c60001b6175b6565b616aab7f8fab81ca0d13b30ca85741827fb4b76f69b79078c5c558e49c55d429e14aaa9d60001b6175b6565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d760405180604001604052808c60c0015173ffffffffffffffffffffffffffffffffffffffff1681526020018f60a0015167ffffffffffffffff168152506040518263ffffffff1660e01b8152600401616b44919061e950565b60006040518083038186803b158015616b5c57600080fd5b505afa158015616b70573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190616b99919061d335565b9050616bc77fbfc582fad12ad310e367ebb39a9ac15d4212665cf85046911fff11c7297bbcd560001b6175b6565b616bf37f97614f19bc29697ea4cd72f185fe801d7cad85eac526a55114103c0da6229b5760001b6175b6565b896102a001511515816101400151151514616c9e57616c347f348a3056b1d3e31d9b53568e9ad9abbe52a742b003bc06172d2af2b5a94588f960001b6175b6565b616c607f9dd80a638fe9727466ca07ba928b122c8a7c1a23f9e8bc4f57887fade25d498960001b6175b6565b60086040517fcd442d65000000000000000000000000000000000000000000000000000000008152600401616c95919061e769565b60405180910390fd5b616cca7f671e7c25fa73c33ffc06d67327811591981d5acc809a4b51ed4c2a050f147a8b60001b6175b6565b616cf67ffbb3f14c62902127692bf9fdfca59c138166b398bf763df3ad8fffe4eefaa9a860001b6175b6565b616d227fafe592b5b0e1c3827c8fdea15b6e4173145e0ba600f47443792423ec89bfe9fe60001b6175b6565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663955f98b7828c6040518363ffffffff1660e01b8152600401616d7f92919061e919565b600060405180830381600087803b158015616d9957600080fd5b505af1158015616dad573d6000803e3d6000fd5b50505050616ddd7f155497e399ee1827264bfe8eaa4d31482bc15865a3a8d0023270d44dee01571c60001b6175b6565b616e097f8aa0a1a2cfd7c3cc648a891383d79a1eb7b99646e3b976e45dabb6a6f2f8815560001b6175b6565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dcf74960826040518263ffffffff1660e01b8152600401616e64919061e8f7565b600060405180830381600087803b158015616e7e57600080fd5b505af1158015616e92573d6000803e3d6000fd5b50505050616ec27f17679fd09d656f58e527d7a049fec504c2c8db17c10bb2af8c7a145b5a163ae960001b6175b6565b616eee7f84716c2da1b234839177ac8e15200cee549947b5ae60953fc482ec7e2301cf2560001b6175b6565b60008160c001511415616fab57616f277fae2d5888261c568c467455386b0890185eca70b3e5d32fe678fcf5c49f68ae7a60001b6175b6565b616f537f69fdc029659f3b6e29e8f6687e0988626e33904191b895ec5d1cd814c2a81ad860001b6175b6565b616f7f7f806fdf35726e64687ff5f9eef8cf84b6ff5cc3c9258aa2afa635f48783a7fdd060001b6175b6565b8960c0015167ffffffffffffffff168c60400151616f9d919061ed04565b8160c0018181525050616fd7565b616fd67e47b3efd127e62f79d3a9b9552015fec95e78b5d5853e71dd2de912ace1c79060001b6175b6565b5b6170037ffc57cdb196f09757664f80fef3576e51dfde3c5c7b75834dcd8acfc4e4581e2d60001b6175b6565b61702f7fb72f0547c1d1d0c85d86c070687dfbff9d4a2f713a1ca3288328f2236099ebac60001b6175b6565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632384a6aa826040518263ffffffff1660e01b815260040161708a919061e8f7565b600060405180830381600087803b1580156170a457600080fd5b505af11580156170b8573d6000803e3d6000fd5b50505050506170ef565b6170ee7fe878a439f521524744f66e6ff2780d1c81f60355ee2fb8a08f4af6aad57e4bb860001b6175b6565b5b61711b7f5501f654d2140e0918f3314f5d0438cd4e1bf11d09c6a755a7de0bad6a56bf9f60001b6175b6565b6171477f742ba4e799bd21a32698048e1ca812eb4469cda1a732f0d91466f3f9ad67e4b360001b6175b6565b83156174de576171797f7af605023729e2dc8b48b1527df22e8323c41f2476e8ae40fde7eebb8ada299660001b6175b6565b6171a57fbb2615bdef852e8e27cf745b935a78767244523d906c5815742ebfddaef530fb60001b6175b6565b6171d17f028be751e9dedb1aa89ff2effb9c70f4529a17654d5185458566c19168ee45b860001b6175b6565b60008b60a0015167ffffffffffffffff1614617447576172137f081a77aea6fd69ffd5bb3ada071823d644000773732181a4c07fd0472246444860001b6175b6565b61723f7f05f20afa332f37fe0106cf73362e52ea186065d5c46f2fc62e9234cb0b55e8ef60001b6175b6565b61726b7fe6a83b3108a059b1c9bf6205cc18ada9dc590be6eed499d0d5dedd0b3e8685d760001b6175b6565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d760405180604001604052808c60c0015173ffffffffffffffffffffffffffffffffffffffff1681526020018f60a0015167ffffffffffffffff168152506040518263ffffffff1660e01b8152600401617304919061e950565b60006040518083038186803b15801561731c57600080fd5b505afa158015617330573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190617359919061d335565b90506173877f97c46fa0aa10d619e0ecf1bea82415ac9333fdecc3e25e60a186740c2ded893760001b6175b6565b6173b37fd33642d30672fd233de6824391dc593d5109f27edb589bd7a04b940617177f9660001b6175b6565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166247c003828c6040518363ffffffff1660e01b815260040161740f92919061e919565b600060405180830381600087803b15801561742957600080fd5b505af115801561743d573d6000803e3d6000fd5b5050505050617474565b6174737fbda71fc2da12636ef1b8e0c563f9c770f260c1061c6254f081da9d722603ce8e60001b6175b6565b5b6174a07f0c1ab0ef414358ff4ed868e7e53bf6fda8b9f4e829ff4e5274b7e544c4a8f09d60001b6175b6565b6174cc7f118e289b04c14820f63a5046af7efbe20e52fdf6639f0a1555ed548e487983a660001b6175b6565b6174d98989848a8e6136a8565b61750b565b61750a7f8b70a8d18fe5262854ecc8f67f9d9f6b4904fc4e788a9c0f1435c167a141402e60001b6175b6565b5b6175377fc9176ba6248ab3320be3a221cb8ff1badae7755cc04c45e1fea5879387fa4b0c60001b6175b6565b6175637fa42db23c8e839c0f8bb3fe58e51e846ec182674fbadf7392aa79f0b9ffda288d60001b6175b6565b7fd936063ae3cbbf2dbcd2adc837a997ab97adb09f24c566aa79d101e88928dea66007438b600001518b60a001516040516175a1949392919061e645565b60405180910390a15050505050505050505050565b50565b60606175e77f08a8d1fdf9c8598e46d904e8abdb35a9012acd5e39a18c37618889f97d2b84aa60001b6175b6565b6176137f70711224d311837b6089458f3d735442962e4628c11c6934b51f81462f7c440a60001b6175b6565b61763f7faa48d71a50b3a8cf5ff291deced677c1b3b9687dbbc914d35e00cbb86b9d49c660001b6175b6565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663977fdfe2836040518263ffffffff1660e01b815260040161769a919061e577565b60006040518083038186803b1580156176b257600080fd5b505afa1580156176c6573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906176ef919061cf43565b9050919050565b60006177247f57690fbf70fe6a4270c39626a9d12abba54408c16c5a28f5850c83e3f4c20c7b60001b6175b6565b6177507fab2a24b3553aa2af51df49b80913d17a13de2367cf1292f6c69fd1a73e3d1c9960001b6175b6565b61777c7fd667bd8aa6e78b4bcb94f58282d05ebc225155dc32258d7c025d7a2f74c13d2960001b6175b6565b60005b84610100015167ffffffffffffffff1681101561848f576177c27f6ec5c30b4fc512ffc63febfce60ce11cdb4593c8da489eb82f839ef5d39c9fcc60001b6175b6565b6177ee7f58bdeae4e12a1931b4c32b12443d5afa603c5c7ee7b9b3fbea255b22db864e8e60001b6175b6565b6000856101600151828151811061782e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015190506178647f39fce1ca179ec1983a61703eec8edc41f159f563d7e052baf082c8c56704c70260001b6175b6565b6178907f0b8768cb94c7135128d0dbb31ba32317a8b1129cc242d3dbac6cd625bafa7e0660001b6175b6565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663defce5d4836040518263ffffffff1660e01b81526004016178ed919061e577565b60006040518083038186803b15801561790557600080fd5b505afa158015617919573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190617942919061d175565b90506179707faf0cd1c18a1f9dda1ed0c761220fb00a15c1fb1903a46f67d2a1c23a1af5014560001b6175b6565b61799c7f309ef9531f36185dff0010ad240f0b4bc8feaf74663f6c9977bdacea648afdd060001b6175b6565b60006179ab82600001516175b9565b90506179d97f2914caac75c34be6795e751888733db9698ab0fc0a37b20234bd7e01cd29749d60001b6175b6565b617a057f3c22dc82e5ea8515b58890a6fe4ec7bbd13405ff97ac5ef8a51065d89ac0635160001b6175b6565b6000617a337f065bfff9efcb6a0a0d4283dedc1569c157384b51ba5a42275f087166d194adff60001b6175b6565b617a5f7f07a0877b470595a1bed7f4069258b1aae748ee2be40e2b449368b886f16b083c60001b6175b6565b60008361010001519050617a957f88de31fd45be0814fc625788152c43510d26f29f7f6d23b6e8ac628eb62034ed60001b6175b6565b617ac17f95e170eb493c17563ecaeb89ad6b8f068747a80986d5af1c5a16634687e59da060001b6175b6565b6000617aef7f22b7f4607c3c44ae18e9e36d7b17c52f1d087a4974c042a68381b6e9cccd881e60001b6175b6565b617b1b7fa661b4f61ab7ec4753954ddf0338c8fa96c8b4028d3dab2de1ad9f23c65543d860001b6175b6565b617b2361b0de565b617b4f7fa26fb8bcf4f1771b2d7e5e63707d42330bb9214ab6033606eeca3114d809238760001b6175b6565b617b7b7fc7f0202eb3b34b0f924ae85277882ec7f79a2d7b8356eef88daa6adc6bd2515760001b6175b6565b60005b855181101561804b57617bb37fd3905f6d07088ed9979f11a10da74291ffce44b443893defdd1c665877d0b6b560001b6175b6565b617bdf7fe13d43abe6c90cd08f5456e8b2ca5fff2777219ccf19f80b16e39e7f64d94b3460001b6175b6565b8c6000015173ffffffffffffffffffffffffffffffffffffffff16868a81518110617c33577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff16141561800c57617c877f92abe34b44d997f7c8261084d1ee1462fad86ef77f80e2ac8e3148673f9fe0ad60001b6175b6565b617cb37f9ffc4a4bb782793ddc737bfabf881edea5554217ba73d2a8531ad959e01b6dbf60001b6175b6565b617cdf7fde8b9b760af09b3aa0b7a17627806f5196254eb5936592657d802308275b7ba560001b6175b6565b60019250617d0f7f667c7270e3db2c393e325079f1c42c98f5f8dc240afc97c2e228dd59ca15d1ca60001b6175b6565b617d3b7f431991219f3c3d47a7c62a7c600a38ef35480a50ad3f20c6b8067ba0206bb37a60001b6175b6565b858981518110617d74577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101519150617daa7fe61f871a49ed694e9fc45fe012bd30cb7ac37251e6e8e5be461bfba02dd99cd960001b6175b6565b617dd67fe44b4e1a428ae440304511e02428486862160379d2ac030bed49b2286e8dd18c60001b6175b6565b600082604001519050617e0b7f37c8a6e0bf563b1a46265d01226189c22f6231fafb88e276f0f3f7dd49fe88e960001b6175b6565b617e377fc3ff640996cf3ec2132cbea9875372a82f61351ca7ccd27ab462f8f7904784f260001b6175b6565b8760e0015167ffffffffffffffff168167ffffffffffffffff161480617e5c57508443115b15617f5757617e8d7f82cb862f7865a5afaaa4b8e630163c232fda31627b9424afa5282602a92e147960001b6175b6565b617eb97ff3f3f2f560c7dd95784aefa3dcb9a0684b90fdd333f45f302ca3e520e1394fb360001b6175b6565b617ee57f763f0f24c8d6d0427860811393bfa52de3dc5552dc32ed1b787da3358f9b985e60001b6175b6565b6001836080019015159081151581525050617f227f2171b27be3e70c20e42724b33e45e8cd952c29dc33a8c49ed787345fae8a854460001b6175b6565b617f4e7f96bf9f6abb6718472a89a02014655cedd02d9eea7ab5f2ed67dd4a43b68ce2f260001b6175b6565b60019550617f84565b617f837fb904005eb0cf39c25ecf40554858ed6fcfbea47011394a7492416d1e8f05088d60001b6175b6565b5b617fb07f4e716e60f83e39dfe0e2c81ca10e7bfa538424639198f32cefb8d7beef5a35e960001b6175b6565b826040018051809190617fc29061f134565b67ffffffffffffffff1667ffffffffffffffff16815250506180067f9bfac304b9cab5359b1ebf9f64acfb68bf6772737b9c62f9e9844536c24380ee60001b6175b6565b5061804b565b6180387fc875b6911556b7889f3b56de25e365269d4e4156e598c06134a5bc829356175a60001b6175b6565b80806180439061f0eb565b915050617b7e565b506180787f9f76c73a085d5dcba5b121f015dc64ce51246c7797a28f59a85761e3117c2df060001b6175b6565b6180a47f6fa469a54555b4473f41910a6cffc15ee95264b922678645137f27f9f2d8d5b060001b6175b6565b8161813e576180d57fc9c54d1d82e0b475ccb8cf34a8bc11ed8f0e971d6fe65c055b6bfb034422d8cd60001b6175b6565b6181017fdf586645913f5534a1aabe70f303ec1648758df327a49bf116fbeef2699dd28960001b6175b6565b61812d7fdf424496ae761981dbded3ba8c231100ed17d8e8870365f83424e37eb7aec92660001b6175b6565b6000985050505050505050506184ed565b61816a7fca070eb9f348ec1d0adf6ec015415922bc180299ad9c31fa4234d7f8e42c156b60001b6175b6565b6181967f52b4058550c148c2f0b79344148b1327225559b65055fe253be8e0f6a234270c60001b6175b6565b6181c27fa61cbde5d9adb043f84800f44e858aa7224910b3982a172ac44eab089d61ffe360001b6175b6565b6181d0866000015186612826565b6181fc7fc1959350b64ab9f61df5d4df4ae4cb3a353f43fefdf48d2b401e26e1a107945660001b6175b6565b6182287fdd7481b9f33e37f240f007826d2846725113a8ed4adf66630aeeb30c5dcb477b60001b6175b6565b83156184485761825a7f7a7a0c93aea5a57b1db2bdb253e6229690517677b9f29dc8ddf6b112313565b460001b6175b6565b6182867f7e8894f83ff0cfaff8da165fb61ecb962d9ecd4d6e1f0ac63caccee517f9c35760001b6175b6565b6182b27fce56d9eea9cdbe46d55ad476369fef4022f210ab7add2397645ba74e3de9c1aa60001b6175b6565b6182bf868c83888e6136a8565b6182eb7f208fe3968fbd6dcd6a0e456bbece4e03dcebc7ff7c0f1741dba23a776009f04760001b6175b6565b6183177f4f48310bf022a0521f33891749f17d12ed7ced384777aa01ca314f2e2ff9ed3460001b6175b6565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166247c0038d886040518363ffffffff1660e01b815260040161837392919061e919565b600060405180830381600087803b15801561838d57600080fd5b505af11580156183a1573d6000803e3d6000fd5b505050506183d17f03a9cb743a5e9bbf61cd83c00c8b90940bca85d47a02164e80480a107a1ebf7b60001b6175b6565b6183fd7fa1c1531ab3acf95261d9b2969670ab5b4008c9228fc856516ea7ccae7bfe21ee60001b6175b6565b7fb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f4560014388600001518e60a0015160405161843b949392919061e645565b60405180910390a1618475565b6184747f11430196cc337324af63b84a80f51a8234cfc39d6999fe5662de45e84c94071560001b6175b6565b5b5050505050505080806184879061f0eb565b91505061777f565b506184bc7fad3069dda3b0afd2ebc3198991c3e049280ce0378a34b66e4c0f49b00ffd95c960001b6175b6565b6184e87fe3772f0a256e29e0cc62671400c74ef97b75f9ed6be3ca06722bf440eb28ae0860001b6175b6565b600190505b9392505050565b6185207fd79ce2cd19d14061c65ec4617ddc245d9cbaa04836dba2257f25acf33ec96b9060001b6175b6565b61854c7ffa3976faa4f233cd76a51fe5deb5f27f877c98aba8ac9722b834eaf0e487ed1560001b6175b6565b6185787f4ce49860b5e0360854214064de601d9b0428e28618beb5a7f3b975e33144776a60001b6175b6565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bb4e4e4283836040518363ffffffff1660e01b81526004016185d592919061e5d0565b600060405180830381600087803b1580156185ef57600080fd5b505af1158015618603573d6000803e3d6000fd5b505050505050565b600060019054906101000a900460ff166186335760008054906101000a900460ff161561863c565b61863b61acfb565b5b61867b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016186729061e7fd565b60405180910390fd5b60008060019054906101000a900460ff1615905080156186cb576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b6186f77f6140b6f4e6fa3e9b599d34dd60104821e2e0550022b72a2e534320b030d9774460001b6175b6565b6187237fe55eb305c9d32f3970dca08f911ca89d7cd395e95d6b185bd3f1efc270509b4860001b6175b6565b61874f7f99c7183088b21dbcaf364c583f3d252794f3da376854392d6fdb9898ca0119ee60001b6175b6565b87600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506187bc7fea03cc40849685a27631df737f2705e57cee339bd864e72af29d0bf037bbf73e60001b6175b6565b6187e87f0476483c2880d795b0895d2aa9a9c3e444377acc4daaf2dbab3746dfa231474a60001b6175b6565b86600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506188557f641ccac28d3a35f5545c6c0740aa766a11ddd9dd4d673deb40d1745e7b0632d160001b6175b6565b6188817fa39c07e3dba029091967c98621522d8bbf07623f7634f177ec9e62fdfbde38ee60001b6175b6565b85600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506188ee7f02eafc399085de6a920c6c70b48df88f85d469934cf29780c43809597f5dee9a60001b6175b6565b61891a7f1e1f30a711d3c9af1c6bf2d8c003e4c9fd3e2f245ed3eab12dd4d4822897dccb60001b6175b6565b84600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506189877faf20f7a3201791d09b65ceb498758d74ec40f4d33f554a38d96b2bebece8398860001b6175b6565b6189b37f4aa2ffe7cc89408e6aed76822c34c6f8547bedc2b858f8c5343be3d6f331ff6b60001b6175b6565b83600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550618a207f3130f98f6192068f0a82127e3587e2600cc3bcc2802072336542f2c02aa15fe660001b6175b6565b618a4c7fd531193aea250c71ef80fb0234d9ad80cdb50e78c45f44fb48ca81724036c0e960001b6175b6565b8260000151600560146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550618aa57f1bc1ee7a83f8b832092fc284f738a7184e47adc4d07d047c8c9848d02662366960001b6175b6565b618ad17fb7be1d129a91606fc7e7f815735711d11984ba80bbded164447049bd799ecee960001b6175b6565b81600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015618b335760008060016101000a81548160ff0219169083151502179055505b5050505050505050565b618b697f36e7b707deabb58bebbdc28821f6d0f50b6d16e0bf895d7bc25f54c2e62ed39260001b6175b6565b618b957f5ecf10f430a069cada01a8e4efafb2f19da916e44333eda7bb14d5e66882a16860001b6175b6565b618bc17f8f6e4b45dc1d470189e41089d3dac7061847ddc946b44bdee6f73b3bd8acf28060001b6175b6565b6000618bcd838661ad0c565b82618bd8919061edc9565b9050618c067f908cd1a75f898875e07becb067dc3d87872daca307b743fa0f04657e72e1793d60001b6175b6565b618c327f4a6395058dfbe966abdbc562d905bb356e766f1fcc45699a6365d3a486757f9760001b6175b6565b8067ffffffffffffffff16846000015167ffffffffffffffff1610618d0957618c7d7fc449cc6fd29fa6b4ec8d7c52bc1e0385786bea67dd8f140ac2b7bf57dba78d4b60001b6175b6565b618ca97f344fcf6ec12f1a252194c40d6434cd9bff7e5b78a745e683f70f4aa871bbce5760001b6175b6565b618cd57f35b89a7804e6248f89e4c0e90a60fa428766ebb0191f169770da213949ab36b460001b6175b6565b8084600001818151618ce7919061ee3f565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050618e0e565b618d357fc2c7ab60b347c201481cfdba2ed6f8346fa7ebda66dcfdfe79c90ee226180adb60001b6175b6565b618d617f0e1ed403a7d2df29aa94c56b17cf5574da135107b90a5e03d367ae712d1da91d60001b6175b6565b618d8d7fc2e03d82f409d0de8d0cf12fc4730ae724369d8086da8b82a104268105be51b160001b6175b6565b6000846000019067ffffffffffffffff16908167ffffffffffffffff1681525050618dda7f842686004b2a4efd2e924a9be272614647bb5ea7f90f530857899616c5aa60c260001b6175b6565b618e067fc769f587e08ef033b47dd14b102ea40178e2023b0d224f6bd410714f2f0e798b60001b6175b6565b836000015190505b618e3a7f7fed0300b22ee14f3a154f136e7a0967bb643f2a3fdb006664404bb7ee94352960001b6175b6565b618e667f580d44ae7f38d805036aebe0ae726da384034cf1f31e5984f98c68150393d9e060001b6175b6565b60008167ffffffffffffffff16111561908c57618ea57f1641ef084252f5710d5c1c2486a78ba1eab7ff91523cb8a32a811c1d63f717ea60001b6175b6565b618ed17fbd2530872a955b63ac3b5571480e163bcfbe93e7cd2953fbd947b2f4b205486c60001b6175b6565b618efd7fa9b61643be691616d262115cb43e28206088a87af72dfff0ccbd650f266d133960001b6175b6565b618f297ff11470d72c67bfef7ee351fbc13dbb959bcd67cfb073e21c8913abeaf7ec7f8260001b6175b6565b8067ffffffffffffffff16341015618f76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401618f6d9061e7dd565b60405180910390fd5b618fa27f493b484f6f70d2c6fffbf34a90e8045e5d7e0144950f28380bfb975c0029aeff60001b6175b6565b618fce7fdae5e583cf336d5fa2237f5431bf8b8ef25c65005d5e41c7582f3e9ae30e923760001b6175b6565b618ffa7f0833be629a04c0019c36ccdfe8d1ef8dc5858520f674082d214f9ed2fefd3e5f60001b6175b6565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c819d86c856040518263ffffffff1660e01b8152600401619055919061e89f565b600060405180830381600087803b15801561906f57600080fd5b505af1158015619083573d6000803e3d6000fd5b505050506190b9565b6190b87fd73abc0253b8ef6f032994447af98bc4a5c86e4b588c2fea9a8ff0c7edf4f12b60001b6175b6565b5b6190e57f5bec59ab1c32e516cc7e03e081b32b7ce77ccfe6a4e29e4a90c1e057d3ce046e60001b6175b6565b6191117f887a6074bf649d35fc26a2d95772c9e5dc04d402bf7f5881739c5cd878ed192960001b6175b6565b619124856000015186602001514361293d565b5050505050565b6191c58383836040516024016191439392919061e79f565b6040516020818303038152906040527f969cdd03000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505061ae70565b505050565b60006191f87f428e1bcaf18ec42118c981ece3690cc01c78ce3b8cde95b33e4710cd3de11f2860001b6175b6565b6192247f447e2a3f9e6091f2e06b7a7a44bb6c67c1ea14ee04ac9fc6a4e6883e0ef1d98e60001b6175b6565b6192507f37d1caacf35ea27e0640500ee71b009bf02398ad7ed06660ca087c41f6950fbf60001b6175b6565b600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002054905092915050565b60006192eb7fc567b8a687cef7503f084b0ea0e2737de8c96d009afe1a129fd04974a7ed62cf60001b6175b6565b6193177fff34ce214d3c58b6e41fa49793c512eae8b807d323f6fbfbaa74b493d0c34e2660001b6175b6565b6193437feb9afd447d0119eec6c0bf812d6bcb349a1eb45e0435856a057de64e3f125ab060001b6175b6565b60008460c0015190506193787fe65b9897435d700c8304fa5ca5ff6ff9dd29bd0733e551d15b75fc36c5cc345660001b6175b6565b6193a47f5237757807b06ddfe1cc56cdda044d550cb19fab8dca0a6d0fae6478381d13a260001b6175b6565b60008660c0015190506193d97f537c13e5110ceec19a55abd540176c6136e77cfdd8746061ae6984d5301fbeba60001b6175b6565b6194057fc1c6b1599e638b5eba413331567c540a63623696fa47257665315d642ad0b0d060001b6175b6565b838267ffffffffffffffff168261941c919061ed04565b106194b05761944d7fc2b3b51d26d7c6e045c54a94fc06eb36c13011911f2ca3bf9dce12714bd404a660001b6175b6565b6194797fab1b6ae163aa2e0fec62d6dbb52f562b0c11107961013b05421348d52d3e73eb60001b6175b6565b6194a57f2278269b1cfbc50c666fc5331539577da103648da63bc3f5efae711797c38e4060001b6175b6565b6000925050506199a3565b6194dc7fd48153f1bc8afeae9380a460749ee46c5190d9c419f6d37a905781664d52259360001b6175b6565b6195087f4ff685308f415e073f1b4afacccb2d5c984eafdf210e4dc8e5d5d539189eebfe60001b6175b6565b6195347fa6de4e2fbca443dd7a19e6a98ee53b5e9aa4eb794feb4cb1b2bd4d6af870bc5860001b6175b6565b6000828286619543919061ee0b565b61954d919061ed98565b905061957b7fecceaebcda79444288bfb34d12365eb3523f0c40d48361ab5cf5221bb1e496ab60001b6175b6565b6195a77f82ff7e7079e511bff144d646f26b9ae8af1bda6d3b185962f6a7578473e6f89360001b6175b6565b60006195d57f32ca81f6229f93f82b8ffadd1fc84c5cbcf31dd5808ce03001bc019bb5d2143360001b6175b6565b6196017f9158d9aabe2e8c73bf234598aae1bce7419b34d1d83dff4f864baa7911387c2760001b6175b6565b600087146197da576196357f7b31e6448150e378fd424f78eeed9d0e8226b62145d2463f33d115e5b9664e1660001b6175b6565b6196617febf274387122f6e647dd3e8f12d50be88a139515909a6133eee62b19d322072f60001b6175b6565b61968d7f7b39fbb6c2bbe3c19b89e4e27fc187a46d865c2762683f1c348ffcbe68dbeeff60001b6175b6565b8367ffffffffffffffff16836196a3919061ed04565b87111561974c576196d67ffd9e830571f8761de93e3cb19849f06163fb254427bc9eeeb432377ee64113d060001b6175b6565b6197027f7ab79d94dabee964058a9156f66230c3ecc33e58410c344bc875694b72b9763160001b6175b6565b61972e7fe372bd25dc7f9e6fa454a37bc1a960ecd5969c34887feddc6e335c15a88aa66860001b6175b6565b83838861973b919061ee0b565b619745919061ed98565b90506197d5565b6197787f15121c0b850669f94a78525cd1b5a71bc16d76eeffcf26c9776dfe842a75f55760001b6175b6565b6197a47f269d7965001c76d5597794a39134b6dfd6db0b6c520d21bdea5d596d077715f760001b6175b6565b6197d07f65c939eb1e3eab9f9ef7dac2863f6e5ea4441a821470557e9211e2eaf960c20c60001b6175b6565b600090505b619807565b6198067fb9fe5d4f938de3d1e9c42c65b572ee1dce7f1331160ec9d9a2291e299465264e60001b6175b6565b5b6198337f9127c5d15a7e2978b988b5c76178781fa2462b214fffd991aa8846e1573bed1260001b6175b6565b61985f7fbfdf8ac41f529f29541e8a14089d58e6c9b9518eeab096ce53b3e1ba9471033560001b6175b6565b8067ffffffffffffffff168267ffffffffffffffff16101561990c576198a77f15655164e6a675a4d407e61760e80d4d321d47852e73baca1fcd4b8ed45bca9a60001b6175b6565b6198d37f6bf238f3442036d704b7977b2097f0f864574275e838b7f5710a7d9efbc9e28460001b6175b6565b6198ff7fe69f0c008700781b84f176d21b4b5ca692cfeff2734b280212f3c640caaac9ca60001b6175b6565b60009450505050506199a3565b6199387f0ed1dac139e052371f8a59c72b6517c9d7c0ccab730f810767bcc1234eef84b660001b6175b6565b6199647f168e26c5a2db2ebdd536541a6065fd1d5d232d76164b19c1f9de70812a61f2a860001b6175b6565b6199907f70f8460e0df1e662bdd3bb161fa62c91eaa4f3869d0b6c3c18315ca4d551b6f960001b6175b6565b808261999c919061ee3f565b9450505050505b949350505050565b60006199d97f920c38f36b3f5f66a72a427d85e95ad565a4aa37994caf321ca876830b8d0ec260001b6175b6565b619a057f9f7fb924ec4e72fe7a63d885dba9bf26c0f5c4c0e3e9516a438a654e68a56e4f60001b6175b6565b619a317f58d90a4a575b4f93c784394f5ce5e9e4c3a7f8af46ae87daa9ec7f93f284e23260001b6175b6565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cc76e80d8460018760400151619a83919061ee3f565b60008960a001518a60800151619a99919061edc9565b8a6101a001518b6101000151619aaf919061ee0b565b6040518663ffffffff1660e01b8152600401619acf95949392919061e96b565b60606040518083038186803b158015619ae757600080fd5b505afa158015619afb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190619b1f919061d55d565b9050619b4d7f23fd89998caa391ffbd0d023973356e4f6441648835a2c52e11baf42fb0e646960001b6175b6565b619b797fe1818bb37b1b0aef35ce99c07992a868f30411ee6a552a27b2baf24097b996ff60001b6175b6565b806040015181602001518260000151619b92919061ed5a565b619b9c919061ed5a565b9150509392505050565b6060619bd47f043babd9bb9e0e9c76fa9f0c68cd587caad0e589b31503cd963ad0f16d26900d60001b6175b6565b619c007fa2bce3ea1622cbebfa50eac412afc04f9407d906fc0fad49dc7b5f545461f98f60001b6175b6565b619c2c7f0881b14f4cbac6a640f56904577740546c91fdd1fe5af3ade0caf3dcf1694f5860001b6175b6565b6000619c37836175b9565b9050619c657fc73ef999d2b7020dc4f80340c44c5cce0afcb65bdc6b735017a66ffedf1e995660001b6175b6565b619c917f52a98d01f32c4d1dce46f9b818475cd3ab23a96c883555e656257c492038ad9960001b6175b6565b60005b8151811015619ed057619cc97fd22b1ecbf0339db49703a8bb58cfff0175cd5cc42df043e8d0ffdaf89ecbcf8f60001b6175b6565b619cf57f2373d625948d3e524d519c1c5900b4e36bd26cf8c09ec83c1fa6324754e3b3b360001b6175b6565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a65374a848481518110619d6e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151602001516040518263ffffffff1660e01b8152600401619d96919061e4c6565b60e06040518083038186803b158015619dae57600080fd5b505afa158015619dc2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190619de6919061d2a2565b9050619e147f8d9c2a0120886a5aaa9f18eda54c9823dd299a10caa7075c36d9be66ad06c5c560001b6175b6565b619e407f18cf4d2d377b061f5aafb5d95206df70e96d7067365e0f26bb61d0c3831f3ed260001b6175b6565b8060c00151838381518110619e7e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050508080619ec89061f0eb565b915050619c94565b50619efd7f3ca6f5846a74b648927d994564d30683a2d3322d0015a07f85e341beff83ce0360001b6175b6565b619f297fbed21c8f3ccc61094dd246c83b8a57c26155c838406e9cc88fa1905d7805cddd60001b6175b6565b80915050919050565b6000619f607f92a263e0f7fce99ce10b411c0d7b506178e3523f76b37f933e152d85c9795ef860001b6175b6565b619f8c7fc53207b556a82b61e8d55d45fafdff2d21c9ea09972cb69194220a29e8ca9dc060001b6175b6565b619fb87fd02960f429ea5d0b6e78bc10e9014b4fc8c6dcb1b83126084f5fcc8bc23a443160001b6175b6565b6000439050619fe97fa2ce1678e85ebf4ac1ff55de11cb693db4749dd55a1498f2c7e8de6a44361d5d60001b6175b6565b61a0157f2e3dd35d1e93dd2d12c9b81cd837b47a72457f8114a3eb391951813ea92accc660001b6175b6565b8260c0015167ffffffffffffffff168161a02f919061ed04565b8460400151118061a05c5750808360c0015167ffffffffffffffff16856040015161a05a919061ed04565b105b1561a0ef5761a08d7f8a288a30e53d223ec12f697555b41ffd362de3bbd559268cd7a8a615b74776ac60001b6175b6565b61a0b97f1c11be8c7833a9c29a0286dfeb1d79d08437862d3bf0f7ec99c3ae64ae98147f60001b6175b6565b61a0e57f5aca5d1a0c017caedfc17679100858029ff0b9ca168d390aeb2293f512518c6b60001b6175b6565b600091505061ab4c565b61a11b7f5ff17198d2f86688937e2d66db6d6961886360239ca70dceccbcf7a31a55f38760001b6175b6565b61a1477fdfe3d37a596e1c5737f7b28325f95f4d5d809b8c54e87960e1a3ff0d969b2dbb60001b6175b6565b61a1737f39f2f0c80339f8a20e89e87f7fbb315f2b58c353bcc2c19a89fe49659f291ee760001b6175b6565b61a17b61b145565b61a1a77f5165ee976b9836bb433900cc23c73645533d1b32298a133155b63d852b6adf5260001b6175b6565b61a1d37f75386f3147cbfddd7db9fb0aab2963b04d024f06705cbb808d610a48b5fd704560001b6175b6565b61a1db61b15f565b61a2077f4b1e5305eef31cea3532f31ad5c5cd10b4d8a200aa0ee20349d6142e6c47ef5360001b6175b6565b61a2337f7544975bfe18be3fdb57bee43c9d6234c3daa5d6ac38ef8e45ac31b5198a5d9960001b6175b6565b606061a2617fcd7f3eaa18eba7044f49e228f851d574c0d29205ad3b998a960153c00da7a9ad60001b6175b6565b61a28d7fb64a679f7f9fe8eab2a03c8dfd087abe150a5141bce2b4a58102856b669eaaf860001b6175b6565b61a29561af58565b61a2c17fcd57a7f1ffedcb2b9253253e6e9e7bb214c3dcdf1d92f558f6b50734fc6505fb60001b6175b6565b61a2ed7fbe8ffdc8997715b0254950b4369f5e0e8bd5984036546376c08449e966e6157760001b6175b6565b8760600151816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505061a3557f645769f1b21d4064fc12919ec2fb385d13645ce2558dd90b6fe5a3dfae25d5b960001b6175b6565b61a3817fecaaa99bd0de2f39619fdee68ee6b4c349ad5ab026aff1441c6ea5e3198df36160001b6175b6565b81816020018190525061a3b67fd85ab96555044395a91bf12f8bf847ca16406672c7c481515563208519e1610b60001b6175b6565b61a3e27f6a938f362bd7d8257a50727c623165992f1b0524862b64ed2ee9f45384d35c4b60001b6175b6565b8660800151816040019067ffffffffffffffff16908167ffffffffffffffff168152505061a4327f591f24fe2a938e4042c98fcfd965064c961d18786652495ae6ac333395c80e8b60001b6175b6565b61a45e7fdacca2a8e2ef852f9e44a5f795207510b2b71301d7e5aaf300d92ca1cc9facb460001b6175b6565b866101800151816060019067ffffffffffffffff16908167ffffffffffffffff168152505061a4af7fe0d5637626943fd49d96fd5bcb460320535009f789ae5577712a2aa96717909460001b6175b6565b61a4db7f1b996e2a735db9f17042d4545f82b10ea4e790b7b452dea00dc75788cf97d60560001b6175b6565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fd7c9808836040518263ffffffff1660e01b815260040161a538919061e87d565b60006040518083038186803b15801561a55057600080fd5b505afa15801561a564573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061a58d919061cf02565b905061a5bb7f97b48fc267e77fccb7dc43ecad070e29871320dc1ee055020e32b6c765e2b90460001b6175b6565b61a5e77fcb9f9495d337eafb10bef4c77d80e567fea2e183ec666acde4c5d168d8d7d0f260001b6175b6565b61a5ef61b066565b61a61b7ff0c32d95a6c03a442955875459a83e98a0350059d8a63ecf9801ba50405292c760001b6175b6565b61a6477f54e1364924f69afe3c273694aaac68a5e40081dfa2e509782de53e39ad0f098d60001b6175b6565b6000816000019067ffffffffffffffff16908167ffffffffffffffff168152505061a6947fbacbf44b24f6cd41e00cc77f0314eca6ebce1d84ad336dc47539c9c4846f924860001b6175b6565b61a6c07f554c4a82174e6be72a24357c25c018c08e3d314583363cf44b5281f8d7f53ea160001b6175b6565b8460000151816020018190525061a6f97f9f2beb20d88aedd2d4097a2ca20d44a8dda8df6db34ba2acae9ed5263c30f35460001b6175b6565b61a7257f1d9e55adcbced2e9f5cac09b59669946deaca4f4a564e6f589e042ffabaed4ba60001b6175b6565b8560200151816040018190525061a75e7f42c5d7b397ea59074ef3017405102011900008d90d8cd9f4f2516df583a76ada60001b6175b6565b61a78a7f3d97537de75785d84f26099d104b642fdf91e296b9a0e768e2b70e9e197775ed60001b6175b6565b8460400151816060018190525061a7c37f16ccb16e4a41378b9ed43c80b0950bc4a8f402d2a8c271fe0023792182edea3260001b6175b6565b61a7ef7f1c466f470101914b7e7bc84279416ebfdb0ab50749d0e4fdb910a00d4a02938560001b6175b6565b81816080018190525061a8247f99d4b32a0fcf8eb9e87f21d232efc4066a946d986e48a5890d9d46aed34c6f2060001b6175b6565b61a8507fea609f7d3df430484a145960f1e9b9da1b091f32b5409566f433cdcd51c401ca60001b6175b6565b84606001518160a0018190525061a8897f5c90e46ad93e0975126af0d56eec156288b5d96d07c049d32bb1097475633ae960001b6175b6565b61a8b57f1305f6ad2c8ed618af6dd9524578e0559d2b815ca252a89d3004dc3635bff02c60001b6175b6565b85600001518160c0018190525061a8ee7fb5e3f4b34a63fc3d95607c604b0dfa68a2486479be7e63fbaf7686eb15230a8460001b6175b6565b61a91a7fb77c6d79f5d1109995f72a7e3c86b8951b9b4faa8270563931c106f6dce1797560001b6175b6565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b6ddeca0836040518263ffffffff1660e01b815260040161a977919061e9e5565b60206040518083038186803b15801561a98f57600080fd5b505afa15801561a9a3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061a9c7919061cf84565b905061a9f57f9177a2d82a45530837b2acf20453430f2e90a7e78d41f405a9bfaa84d7f6134260001b6175b6565b61aa217fe535de601ad48db66f0ef4ec8e5bba7f74bc5d082a4e6980feff9ff4416d6ba660001b6175b6565b8061aabb5761aa527fc12372f42e1d5ff8768d8c4dc009632e8022f979ea30eb702760d054b58a135160001b6175b6565b61aa7e7f05f9c76f7747b746e60dae86b9db91d0e48f111cca695d3b1b0bb325ac71b1d960001b6175b6565b61aaaa7f4fb8f8c4788dbcc5195d53a1cf7d202c648a7c664efe112c305bbf49a34d63db60001b6175b6565b60009850505050505050505061ab4c565b61aae77fe693f4fb98ddf07777f2eb3ee2b0900295190f930c35d4afd2dcf2a7d853a6f860001b6175b6565b61ab137f4712799cc8cd2c0818ab0e574c6c30dfe0c62331140eb677074c0ab50abdf37f60001b6175b6565b61ab3f7f65273b12e8392db757daa3be4b80cf456a07f17b5852cd0c5c9caf54d24b016d60001b6175b6565b6001985050505050505050505b92915050565b600061ab807f1afa166f9206d7acadf80b68c76817c6dfecb49889b0acb41f3af875cb4b998860001b6175b6565b61abac7fee0e95267b4bdc677e4d2c30e3e537c8e689ce71ed410b933ab45ae6b711494760001b6175b6565b61abd87f8d10afe56b2353d3ac50c9d3c8b29015d39720599e20a3b9ce88f4b05211add860001b6175b6565b8143111561ac6d5761ac0c7f52cb81c196e79141f8a2b997ca09889be0f2dba0c56a64acdab301fe23f09e0360001b6175b6565b61ac387fa412963f5783fa46f92181bbabd06f13eb15132070bf72fce0bb67f1c6952c2560001b6175b6565b61ac647fda5fbb2deab399ff1e7b5122b29bee87ab0bbb19cd577f2851b7c52f9e86f32660001b6175b6565b6001905061acf6565b61ac997faed76bd3e02d3e22125b5cee0cd2191d12a0e7fcbe6a77917dada412b5d9a8f260001b6175b6565b61acc57f1b9502f8d88523498966d13f0259b917ceda06de15b97ab776dd38907c24f95e60001b6175b6565b61acf17fad945ca0cbd2170b6491a8366ccc72e40ca53836cd11a83a961529f0afdfdec360001b6175b6565b600090505b919050565b600061ad063061ae99565b15905090565b600061ad3a7f3d648d98081a28c662806c1d60075156710ae9c776c78095692b00876546dbef60001b6175b6565b61ad667ff3405b84f112f278f637612ef29278744afa095b1c442160da473e152a71547e60001b6175b6565b61ad927fd5053904a73dbeef0b4942f867723551bfb37e4d0c7df651bef2c4732653330c60001b6175b6565b60006002905061adc47fbdca501af8a0a6ad9d4028c7bca3d9eacc83fc2cffb2262df5b8d8f0ed25194b60001b6175b6565b61adf07f9f10345e5f43458e57a947ca6b83559efa265af8f4fc590e330a25afdc57a26d60001b6175b6565b600061ae0085856060015161aeac565b8261ae0b919061edc9565b905061ae397f48df4fc357928b767793e678821caa8d1bff6d60a6cfd8356e6f1c43eae73dfe60001b6175b6565b61ae657fd196f59e1f7debf759858b6133745fbe8b0ef083d3732ae70816c38d8fdb9a8560001b6175b6565b809250505092915050565b60008151905060006a636f6e736f6c652e6c6f679050602083016000808483855afa5050505050565b600080823b905060008111915050919050565b600061aeda7f8ef6f8f1a61283a5dc772d4207fe2a7dca9244113b7793d4b8ee93978412a7b960001b6175b6565b61af067fa1d50950a2fc0b61db43fbbe13fd3d5bc24a5e7bdcf652b92fd7ac871c746cd860001b6175b6565b61af327fc986e9ceee21fcae08b852797474834920616b5b5a34d75daf4ff2d996ef199760001b6175b6565b620fa00082846060015161af46919061edc9565b61af50919061ed98565b905092915050565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6040518060c00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081526020016060815260200160608152602001606081525090565b604051806060016040528061b00761b191565b81526020016060815260200161b01b61afaa565b81525090565b6040518060e00160405280606081526020016060815260200160608152602001606081526020016060815260200161b05761b27e565b81526020016000151581525090565b6040518060e00160405280600067ffffffffffffffff1681526020016060815260200160608152602001606081526020016060815260200160608152602001606081525090565b604051806060016040528061b0c061b416565b815260200160608152602001600067ffffffffffffffff1681525090565b6040518060a00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff168152602001600081526020016000151581525090565b604051806040016040528060608152602001606081525090565b604051806080016040528060608152602001600067ffffffffffffffff16815260200160608152602001606081525090565b604051806101800160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681526020016000600281111561b227577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016000815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600015158152602001606081525090565b604051806102e0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600081526020016000151581526020016000600181111561b395577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600067ffffffffffffffff1681526020016060815260200160608152602001606081526020016000600281111561b3fa577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200160001515815260200161b41061b416565b81525090565b6040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b600061b46861b4638461ea2c565b61ea07565b9050808382526020820190508285602086028201111561b48757600080fd5b60005b8581101561b4b7578161b49d888261b905565b84526020840193506020830192505060018101905061b48a565b5050509392505050565b600061b4d461b4cf8461ea2c565b61ea07565b9050808382526020820190508285602086028201111561b4f357600080fd5b60005b8581101561b523578161b509888261b91a565b84526020840193506020830192505060018101905061b4f6565b5050509392505050565b600061b54061b53b8461ea58565b61ea07565b9050808382526020820190508285602086028201111561b55f57600080fd5b60005b8581101561b5a957813567ffffffffffffffff81111561b58157600080fd5b80860161b58e898261bae8565b8552602085019450602084019350505060018101905061b562565b5050509392505050565b600061b5c661b5c18461ea58565b61ea07565b9050808382526020820190508285602086028201111561b5e557600080fd5b60005b8581101561b62f57815167ffffffffffffffff81111561b60757600080fd5b80860161b614898261bb12565b8552602085019450602084019350505060018101905061b5e8565b5050509392505050565b600061b64c61b6478461ea84565b61ea07565b9050808382526020820190508285604086028201111561b66b57600080fd5b60005b8581101561b69b578161b681888261bc0e565b84526020840193506040830192505060018101905061b66e565b5050509392505050565b600061b6b861b6b38461eab0565b61ea07565b9050808382526020820190508285602086028201111561b6d757600080fd5b60005b8581101561b72157815167ffffffffffffffff81111561b6f957600080fd5b80860161b706898261c266565b8552602085019450602084019350505060018101905061b6da565b5050509392505050565b600061b73e61b7398461eadc565b61ea07565b9050808382526020820190508285602086028201111561b75d57600080fd5b60005b8581101561b7a757815167ffffffffffffffff81111561b77f57600080fd5b80860161b78c898261c2de565b8552602085019450602084019350505060018101905061b760565b5050509392505050565b600061b7c461b7bf8461eb08565b61ea07565b905080838252602082019050828560a086028201111561b7e357600080fd5b60005b8581101561b813578161b7f9888261c786565b845260208401935060a0830192505060018101905061b7e6565b5050509392505050565b600061b83061b82b8461eb08565b61ea07565b905080838252602082019050828560a086028201111561b84f57600080fd5b60005b8581101561b87f578161b865888261c80e565b845260208401935060a0830192505060018101905061b852565b5050509392505050565b600061b89c61b8978461eb34565b61ea07565b90508281526020810184848401111561b8b457600080fd5b61b8bf84828561f078565b509392505050565b600061b8da61b8d58461eb34565b61ea07565b90508281526020810184848401111561b8f257600080fd5b61b8fd84828561f087565b509392505050565b60008135905061b9148161f2e6565b92915050565b60008151905061b9298161f2e6565b92915050565b600082601f83011261b94057600080fd5b813561b95084826020860161b455565b91505092915050565b600082601f83011261b96a57600080fd5b815161b97a84826020860161b4c1565b91505092915050565b600082601f83011261b99457600080fd5b813561b9a484826020860161b52d565b91505092915050565b600082601f83011261b9be57600080fd5b815161b9ce84826020860161b5b3565b91505092915050565b600082601f83011261b9e857600080fd5b815161b9f884826020860161b639565b91505092915050565b600082601f83011261ba1257600080fd5b815161ba2284826020860161b6a5565b91505092915050565b600082601f83011261ba3c57600080fd5b815161ba4c84826020860161b72b565b91505092915050565b600082601f83011261ba6657600080fd5b813561ba7684826020860161b7b1565b91505092915050565b600082601f83011261ba9057600080fd5b815161baa084826020860161b81d565b91505092915050565b60008135905061bab88161f2fd565b92915050565b60008151905061bacd8161f2fd565b92915050565b60008135905061bae28161f314565b92915050565b600082601f83011261baf957600080fd5b813561bb0984826020860161b889565b91505092915050565b600082601f83011261bb2357600080fd5b815161bb3384826020860161b8c7565b91505092915050565b60008135905061bb4b8161f32b565b92915050565b60008135905061bb608161f342565b92915050565b60008135905061bb758161f359565b92915050565b60008135905061bb8a8161f370565b92915050565b60008135905061bb9f8161f387565b92915050565b60008135905061bbb48161f39e565b92915050565b60008135905061bbc98161f3b5565b92915050565b60008151905061bbde8161f3b5565b92915050565b60008135905061bbf38161f3c5565b92915050565b60008151905061bc088161f3c5565b92915050565b60006040828403121561bc2057600080fd5b61bc2a604061ea07565b9050600061bc3a8482850161ce74565b600083015250602061bc4e8482850161ce74565b60208301525092915050565b6000610320828403121561bc6d57600080fd5b61bc786102e061ea07565b9050600082013567ffffffffffffffff81111561bc9457600080fd5b61bca08482850161bae8565b600083015250602061bcb48482850161b905565b602083015250604082013567ffffffffffffffff81111561bcd457600080fd5b61bce08482850161bae8565b604083015250606061bcf48482850161ce89565b606083015250608061bd088482850161ce89565b60808301525060a061bd1c8482850161ce89565b60a08301525060c061bd308482850161ce89565b60c08301525060e061bd448482850161ce89565b60e08301525061010061bd598482850161ce4a565b6101008301525061012061bd6f8482850161ce89565b6101208301525061014061bd858482850161ce89565b6101408301525061016082013567ffffffffffffffff81111561bda757600080fd5b61bdb38482850161bae8565b6101608301525061018061bdc98482850161ce89565b610180830152506101a061bddf8482850161ce4a565b6101a0830152506101c061bdf58482850161baa9565b6101c0830152506101e061be0b8482850161bbe4565b6101e08301525061020061be218482850161ce89565b6102008301525061022082013567ffffffffffffffff81111561be4357600080fd5b61be4f8482850161b92f565b6102208301525061024082013567ffffffffffffffff81111561be7157600080fd5b61be7d8482850161b92f565b6102408301525061026082013567ffffffffffffffff81111561be9f57600080fd5b61beab8482850161bae8565b6102608301525061028061bec18482850161bbba565b610280830152506102a061bed78482850161baa9565b6102a0830152506102c061beed8482850161c5e2565b6102c08301525092915050565b6000610320828403121561bf0d57600080fd5b61bf186102e061ea07565b9050600082015167ffffffffffffffff81111561bf3457600080fd5b61bf408482850161bb12565b600083015250602061bf548482850161b91a565b602083015250604082015167ffffffffffffffff81111561bf7457600080fd5b61bf808482850161bb12565b604083015250606061bf948482850161ce9e565b606083015250608061bfa88482850161ce9e565b60808301525060a061bfbc8482850161ce9e565b60a08301525060c061bfd08482850161ce9e565b60c08301525060e061bfe48482850161ce9e565b60e08301525061010061bff98482850161ce5f565b6101008301525061012061c00f8482850161ce9e565b6101208301525061014061c0258482850161ce9e565b6101408301525061016082015167ffffffffffffffff81111561c04757600080fd5b61c0538482850161bb12565b6101608301525061018061c0698482850161ce9e565b610180830152506101a061c07f8482850161ce5f565b6101a0830152506101c061c0958482850161babe565b6101c0830152506101e061c0ab8482850161bbf9565b6101e08301525061020061c0c18482850161ce9e565b6102008301525061022082015167ffffffffffffffff81111561c0e357600080fd5b61c0ef8482850161b959565b6102208301525061024082015167ffffffffffffffff81111561c11157600080fd5b61c11d8482850161b959565b6102408301525061026082015167ffffffffffffffff81111561c13f57600080fd5b61c14b8482850161bb12565b6102608301525061028061c1618482850161bbcf565b610280830152506102a061c1778482850161babe565b6102a0830152506102c061c18d8482850161c642565b6102c08301525092915050565b600060c0828403121561c1ac57600080fd5b61c1b660c061ea07565b9050600082013567ffffffffffffffff81111561c1d257600080fd5b61c1de8482850161bae8565b600083015250602082013567ffffffffffffffff81111561c1fe57600080fd5b61c20a8482850161bae8565b602083015250604061c21e8482850161ce4a565b604083015250606061c2328482850161b905565b606083015250608061c2468482850161ce89565b60808301525060a061c25a8482850161ce89565b60a08301525092915050565b60006060828403121561c27857600080fd5b61c282606061ea07565b9050600061c2928482850161ce9e565b600083015250602061c2a68482850161ce9e565b602083015250604082015167ffffffffffffffff81111561c2c657600080fd5b61c2d28482850161bb12565b60408301525092915050565b60006040828403121561c2f057600080fd5b61c2fa604061ea07565b9050600061c30a8482850161ce9e565b600083015250602082015167ffffffffffffffff81111561c32a57600080fd5b61c3368482850161ba01565b60208301525092915050565b600060e0828403121561c35457600080fd5b61c35e60e061ea07565b9050600061c36e8482850161ce89565b600083015250602061c3828482850161ce89565b602083015250604061c3968482850161ce89565b604083015250606061c3aa8482850161ce89565b606083015250608061c3be8482850161ce89565b60808301525060a061c3d28482850161b905565b60a08301525060c061c3e68482850161b905565b60c08301525092915050565b600060e0828403121561c40457600080fd5b61c40e60e061ea07565b9050600061c41e8482850161ce9e565b600083015250602061c4328482850161ce9e565b602083015250604061c4468482850161ce9e565b604083015250606061c45a8482850161ce9e565b606083015250608061c46e8482850161ce9e565b60808301525060a061c4828482850161b91a565b60a08301525060c061c4968482850161b91a565b60c08301525092915050565b600060e0828403121561c4b457600080fd5b61c4be60e061ea07565b9050600082015167ffffffffffffffff81111561c4da57600080fd5b61c4e68482850161b9ad565b600083015250602082015167ffffffffffffffff81111561c50657600080fd5b61c5128482850161b9ad565b602083015250604082015167ffffffffffffffff81111561c53257600080fd5b61c53e8482850161b9d7565b604083015250606082015167ffffffffffffffff81111561c55e57600080fd5b61c56a8482850161ba2b565b606083015250608082015167ffffffffffffffff81111561c58a57600080fd5b61c5968482850161bb12565b60808301525060a082015167ffffffffffffffff81111561c5b657600080fd5b61c5c28482850161befa565b60a08301525060c061c5d68482850161babe565b60c08301525092915050565b60006060828403121561c5f457600080fd5b61c5fe606061ea07565b9050600061c60e8482850161ce89565b600083015250602061c6228482850161ce89565b602083015250604061c6368482850161ce89565b60408301525092915050565b60006060828403121561c65457600080fd5b61c65e606061ea07565b9050600061c66e8482850161ce9e565b600083015250602061c6828482850161ce9e565b602083015250604061c6968482850161ce9e565b60408301525092915050565b60006060828403121561c6b457600080fd5b61c6be606061ea07565b9050600061c6ce8482850161b91a565b600083015250602061c6e28482850161ce5f565b602083015250604061c6f68482850161ce9e565b60408301525092915050565b60006020828403121561c71457600080fd5b61c71e602061ea07565b9050600061c72e8482850161ce89565b60008301525092915050565b60006040828403121561c74c57600080fd5b61c756604061ea07565b9050600061c7668482850161ce89565b600083015250602061c77a8482850161ce89565b60208301525092915050565b600060a0828403121561c79857600080fd5b61c7a260a061ea07565b9050600061c7b28482850161b905565b600083015250602061c7c68482850161b905565b602083015250604061c7da8482850161ce89565b604083015250606061c7ee8482850161ce4a565b606083015250608061c8028482850161baa9565b60808301525092915050565b600060a0828403121561c82057600080fd5b61c82a60a061ea07565b9050600061c83a8482850161b91a565b600083015250602061c84e8482850161b91a565b602083015250604061c8628482850161ce9e565b604083015250606061c8768482850161ce5f565b606083015250608061c88a8482850161babe565b60808301525092915050565b6000610180828403121561c8a957600080fd5b61c8b461018061ea07565b9050600061c8c48482850161b905565b600083015250602061c8d88482850161ce89565b602083015250604061c8ec8482850161ce89565b604083015250606061c9008482850161ce89565b606083015250608061c9148482850161bbba565b60808301525060a061c9288482850161ce4a565b60a08301525060c061c93c8482850161ce4a565b60c08301525060e061c9508482850161ce89565b60e08301525061010061c9658482850161ce89565b6101008301525061012061c97b8482850161ce89565b6101208301525061014061c9918482850161baa9565b6101408301525061016082013567ffffffffffffffff81111561c9b357600080fd5b61c9bf8482850161b983565b6101608301525092915050565b6000610180828403121561c9df57600080fd5b61c9ea61018061ea07565b9050600061c9fa8482850161b91a565b600083015250602061ca0e8482850161ce9e565b602083015250604061ca228482850161ce9e565b604083015250606061ca368482850161ce9e565b606083015250608061ca4a8482850161bbcf565b60808301525060a061ca5e8482850161ce5f565b60a08301525060c061ca728482850161ce5f565b60c08301525060e061ca868482850161ce9e565b60e08301525061010061ca9b8482850161ce9e565b6101008301525061012061cab18482850161ce9e565b6101208301525061014061cac78482850161babe565b6101408301525061016082015167ffffffffffffffff81111561cae957600080fd5b61caf58482850161b9ad565b6101608301525092915050565b60006080828403121561cb1457600080fd5b61cb1e608061ea07565b9050600061cb2e8482850161b905565b600083015250602061cb428482850161ce89565b602083015250604061cb568482850161ce89565b604083015250606082013567ffffffffffffffff81111561cb7657600080fd5b61cb828482850161bae8565b60608301525092915050565b60006040828403121561cba057600080fd5b61cbaa604061ea07565b9050600061cbba8482850161b905565b600083015250602061cbce8482850161ce89565b60208301525092915050565b6000610160828403121561cbed57600080fd5b61cbf861016061ea07565b9050600061cc088482850161ce89565b600083015250602061cc1c8482850161ce89565b602083015250604061cc308482850161ce89565b604083015250606061cc448482850161ce89565b606083015250608061cc588482850161ce89565b60808301525060a061cc6c8482850161ce89565b60a08301525060c061cc808482850161ce89565b60c08301525060e061cc948482850161ce89565b60e08301525061010061cca98482850161ce89565b6101008301525061012061ccbf8482850161ce89565b6101208301525061014061ccd58482850161ce89565b6101408301525092915050565b6000610160828403121561ccf557600080fd5b61cd0061016061ea07565b9050600061cd108482850161ce9e565b600083015250602061cd248482850161ce9e565b602083015250604061cd388482850161ce9e565b604083015250606061cd4c8482850161ce9e565b606083015250608061cd608482850161ce9e565b60808301525060a061cd748482850161ce9e565b60a08301525060c061cd888482850161ce9e565b60c08301525060e061cd9c8482850161ce9e565b60e08301525061010061cdb18482850161ce9e565b6101008301525061012061cdc78482850161ce9e565b6101208301525061014061cddd8482850161ce9e565b6101408301525092915050565b60006060828403121561cdfc57600080fd5b61ce06606061ea07565b9050600061ce168482850161ce9e565b600083015250602061ce2a8482850161ce9e565b602083015250604061ce3e8482850161ce9e565b60408301525092915050565b60008135905061ce598161f3d5565b92915050565b60008151905061ce6e8161f3d5565b92915050565b60008151905061ce838161f3ec565b92915050565b60008135905061ce988161f403565b92915050565b60008151905061cead8161f403565b92915050565b60008060006060848603121561cec857600080fd5b600061ced68682870161b905565b935050602061cee78682870161ce89565b925050604061cef88682870161ce4a565b9150509250925092565b60006020828403121561cf1457600080fd5b600082015167ffffffffffffffff81111561cf2e57600080fd5b61cf3a8482850161b9d7565b91505092915050565b60006020828403121561cf5557600080fd5b600082015167ffffffffffffffff81111561cf6f57600080fd5b61cf7b8482850161ba7f565b91505092915050565b60006020828403121561cf9657600080fd5b600061cfa48482850161babe565b91505092915050565b60006020828403121561cfbf57600080fd5b600061cfcd8482850161bad3565b91505092915050565b60006020828403121561cfe857600080fd5b600082013567ffffffffffffffff81111561d00257600080fd5b61d00e8482850161bae8565b91505092915050565b6000806040838503121561d02a57600080fd5b600083013567ffffffffffffffff81111561d04457600080fd5b61d0508582860161bae8565b925050602083013567ffffffffffffffff81111561d06d57600080fd5b61d0798582860161ba55565b9150509250929050565b6000806060838503121561d09657600080fd5b600083013567ffffffffffffffff81111561d0b057600080fd5b61d0bc8582860161bae8565b925050602061d0cd8582860161c73a565b9150509250929050565b600080600080600080600060e0888a03121561d0f257600080fd5b600061d1008a828b0161bb3c565b975050602061d1118a828b0161bb51565b965050604061d1228a828b0161bb66565b955050606061d1338a828b0161bb7b565b945050608061d1448a828b0161bb90565b93505060a061d1558a828b0161c702565b92505060c061d1668a828b0161bba5565b91505092959891949750929550565b60006020828403121561d18757600080fd5b600082015167ffffffffffffffff81111561d1a157600080fd5b61d1ad8482850161befa565b91505092915050565b6000806000806000610320868803121561d1cf57600080fd5b600086013567ffffffffffffffff81111561d1e957600080fd5b61d1f58882890161bc5a565b955050602061d2068882890161c342565b94505061010061d2188882890161c786565b9350506101a086013567ffffffffffffffff81111561d23657600080fd5b61d2428882890161ba55565b9250506101c061d2548882890161cbda565b9150509295509295909350565b60006020828403121561d27357600080fd5b600082013567ffffffffffffffff81111561d28d57600080fd5b61d2998482850161c19a565b91505092915050565b600060e0828403121561d2b457600080fd5b600061d2c28482850161c3f2565b91505092915050565b60006020828403121561d2dd57600080fd5b600082015167ffffffffffffffff81111561d2f757600080fd5b61d3038482850161c4a2565b91505092915050565b60006060828403121561d31e57600080fd5b600061d32c8482850161c6a2565b91505092915050565b60006020828403121561d34757600080fd5b600082015167ffffffffffffffff81111561d36157600080fd5b61d36d8482850161c9cc565b91505092915050565b6000806000610260848603121561d38c57600080fd5b600084013567ffffffffffffffff81111561d3a657600080fd5b61d3b28682870161c896565b935050602061d3c38682870161c342565b92505061010061d3d58682870161cbda565b9150509250925092565b600080600080610280858703121561d3f657600080fd5b600085013567ffffffffffffffff81111561d41057600080fd5b61d41c8782880161c896565b945050602061d42d8782880161c342565b93505061010061d43f8782880161cbda565b92505061026061d4518782880161ce89565b91505092959194509250565b60006020828403121561d46f57600080fd5b600082013567ffffffffffffffff81111561d48957600080fd5b61d4958482850161cb02565b91505092915050565b6000806040838503121561d4b157600080fd5b600083013567ffffffffffffffff81111561d4cb57600080fd5b61d4d78582860161cb02565b925050602083013567ffffffffffffffff81111561d4f457600080fd5b61d5008582860161c896565b9150509250929050565b60006040828403121561d51c57600080fd5b600061d52a8482850161cb8e565b91505092915050565b6000610160828403121561d54657600080fd5b600061d5548482850161cce2565b91505092915050565b60006060828403121561d56f57600080fd5b600061d57d8482850161cdea565b91505092915050565b600061d592838361d60a565b60208301905092915050565b600061d5aa838361d8bf565b905092915050565b600061d5be838361da82565b60408301905092915050565b600061d5d6838361dd22565b905092915050565b600061d5ea838361dd72565b905092915050565b600061d5fe838361df4e565b60a08301905092915050565b61d6138161ee73565b82525050565b61d6228161ee73565b82525050565b600061d6338261ebc5565b61d63d818561ec6b565b935061d6488361eb65565b8060005b8381101561d67957815161d660888261d586565b975061d66b8361ec1d565b92505060018101905061d64c565b5085935050505092915050565b600061d6918261ebd0565b61d69b818561ec7c565b93508360208202850161d6ad8561eb75565b8060005b8581101561d6e9578484038952815161d6ca858261d59e565b945061d6d58361ec2a565b925060208a0199505060018101905061d6b1565b50829750879550505050505092915050565b600061d7068261ebdb565b61d710818561ec8d565b935061d71b8361eb85565b8060005b8381101561d74c57815161d733888261d5b2565b975061d73e8361ec37565b92505060018101905061d71f565b5085935050505092915050565b600061d7648261ebe6565b61d76e818561ec9e565b93508360208202850161d7808561eb95565b8060005b8581101561d7bc578484038952815161d79d858261d5ca565b945061d7a88361ec44565b925060208a0199505060018101905061d784565b50829750879550505050505092915050565b600061d7d98261ebf1565b61d7e3818561ecaf565b93508360208202850161d7f58561eba5565b8060005b8581101561d831578484038952815161d812858261d5de565b945061d81d8361ec51565b925060208a0199505060018101905061d7f9565b50829750879550505050505092915050565b600061d84e8261ebfc565b61d858818561ecc0565b935061d8638361ebb5565b8060005b8381101561d89457815161d87b888261d5f2565b975061d8868361ec5e565b92505060018101905061d867565b5085935050505092915050565b61d8aa8161ee85565b82525050565b61d8b98161ee85565b82525050565b600061d8ca8261ec07565b61d8d4818561ecd1565b935061d8e481856020860161f087565b61d8ed8161f221565b840191505092915050565b600061d9038261ec07565b61d90d818561ece2565b935061d91d81856020860161f087565b61d9268161f221565b840191505092915050565b61d93a8161ef8e565b82525050565b61d9498161efa0565b82525050565b61d9588161efa0565b82525050565b61d9678161efb2565b82525050565b61d9768161efc4565b82525050565b61d9858161efd6565b82525050565b61d9948161efe8565b82525050565b61d9a38161effa565b82525050565b61d9b28161f00c565b82525050565b61d9c18161f01e565b82525050565b61d9d08161f030565b82525050565b61d9df8161f042565b82525050565b61d9ee8161f054565b82525050565b61d9fd8161f066565b82525050565b600061da0e8261ec12565b61da18818561ecf3565b935061da2881856020860161f087565b61da318161f221565b840191505092915050565b600061da4960168361ecf3565b915061da548261f232565b602082019050919050565b600061da6c602e8361ecf3565b915061da778261f25b565b604082019050919050565b60408201600082015161da98600085018261e499565b50602082015161daab602085018261e499565b50505050565b600061032083016000830151848203600086015261dacf828261d8bf565b915050602083015161dae4602086018261d60a565b506040830151848203604086015261dafc828261d8bf565b915050606083015161db11606086018261e4a8565b50608083015161db24608086018261e4a8565b5060a083015161db3760a086018261e4a8565b5060c083015161db4a60c086018261e4a8565b5060e083015161db5d60e086018261e4a8565b5061010083015161db7261010086018261e47b565b5061012083015161db8761012086018261e4a8565b5061014083015161db9c61014086018261e4a8565b5061016083015184820361016086015261dbb6828261d8bf565b91505061018083015161dbcd61018086018261e4a8565b506101a083015161dbe26101a086018261e47b565b506101c083015161dbf76101c086018261d8a1565b506101e083015161dc0c6101e086018261d95e565b5061020083015161dc2161020086018261e4a8565b5061022083015184820361022086015261dc3b828261d628565b91505061024083015184820361024086015261dc57828261d628565b91505061026083015184820361026086015261dc73828261d8bf565b91505061028083015161dc8a61028086018261d940565b506102a083015161dc9f6102a086018261d8a1565b506102c083015161dcb46102c086018261de3d565b508091505092915050565b600060808301600083015161dcd7600086018261d60a565b506020830151848203602086015261dcef828261d8bf565b915050604083015161dd04604086018261e4a8565b50606083015161dd17606086018261e4a8565b508091505092915050565b600060608301600083015161dd3a600086018261e4a8565b50602083015161dd4d602086018261e4a8565b506040830151848203604086015261dd65828261d8bf565b9150508091505092915050565b600060408301600083015161dd8a600086018261e4a8565b506020830151848203602086015261dda2828261d759565b9150508091505092915050565b60e08201600082015161ddc5600085018261e4a8565b50602082015161ddd8602085018261e4a8565b50604082015161ddeb604085018261e4a8565b50606082015161ddfe606085018261e4a8565b50608082015161de11608085018261e4a8565b5060a082015161de2460a085018261d60a565b5060c082015161de3760c085018261d60a565b50505050565b60608201600082015161de53600085018261e4a8565b50602082015161de66602085018261e4a8565b50604082015161de79604085018261e4a8565b50505050565b60608201600082015161de95600085018261d60a565b50602082015161dea8602085018261e47b565b50604082015161debb604085018261e4a8565b50505050565b6000606083016000830151848203600086015261dede828261dfb6565b9150506020830151848203602086015261def8828261d6fb565b9150506040830151848203604086015261df12828261e1be565b9150508091505092915050565b60408201600082015161df35600085018261e4a8565b50602082015161df48602085018261e4a8565b50505050565b60a08201600082015161df64600085018261d60a565b50602082015161df77602085018261d60a565b50604082015161df8a604085018261e4a8565b50606082015161df9d606085018261e47b565b50608082015161dfb0608085018261d8a1565b50505050565b60006101808301600083015161dfcf600086018261d60a565b50602083015161dfe2602086018261e4a8565b50604083015161dff5604086018261e4a8565b50606083015161e008606086018261e4a8565b50608083015161e01b608086018261d940565b5060a083015161e02e60a086018261e47b565b5060c083015161e04160c086018261e47b565b5060e083015161e05460e086018261e4a8565b5061010083015161e06961010086018261e4a8565b5061012083015161e07e61012086018261e4a8565b5061014083015161e09361014086018261d8a1565b5061016083015184820361016086015261e0ad828261d686565b9150508091505092915050565b60006101808301600083015161e0d3600086018261d60a565b50602083015161e0e6602086018261e4a8565b50604083015161e0f9604086018261e4a8565b50606083015161e10c606086018261e4a8565b50608083015161e11f608086018261d940565b5060a083015161e13260a086018261e47b565b5060c083015161e14560c086018261e47b565b5060e083015161e15860e086018261e4a8565b5061010083015161e16d61010086018261e4a8565b5061012083015161e18261012086018261e4a8565b5061014083015161e19761014086018261d8a1565b5061016083015184820361016086015261e1b1828261d686565b9150508091505092915050565b600060c08301600083015161e1d6600086018261e4a8565b50602083015161e1e9602086018261e4a8565b506040830151848203604086015261e201828261d8bf565b9150506060830151848203606086015261e21b828261d8bf565b9150506080830151848203608086015261e235828261d7ce565b91505060a083015184820360a086015261e24f828261d8bf565b9150508091505092915050565b60408201600082015161e272600085018261d60a565b50602082015161e285602085018261e4a8565b50505050565b6101608201600082015161e2a2600085018261e4a8565b50602082015161e2b5602085018261e4a8565b50604082015161e2c8604085018261e4a8565b50606082015161e2db606085018261e4a8565b50608082015161e2ee608085018261e4a8565b5060a082015161e30160a085018261e4a8565b5060c082015161e31460c085018261e4a8565b5060e082015161e32760e085018261e4a8565b5061010082015161e33c61010085018261e4a8565b5061012082015161e35161012085018261e4a8565b5061014082015161e36661014085018261e4a8565b50505050565b600060a08301600083015161e384600086018261de3d565b506020830151848203606086015261e39c828261d8bf565b915050604083015161e3b1608086018261e4a8565b508091505092915050565b600060e08301600083015161e3d4600086018261e4a8565b506020830151848203602086015261e3ec828261d8bf565b9150506040830151848203604086015261e406828261d686565b9150506060830151848203606086015261e420828261d686565b9150506080830151848203608086015261e43a828261d6fb565b91505060a083015184820360a086015261e454828261d7ce565b91505060c083015184820360c086015261e46e828261d8bf565b9150508091505092915050565b61e4848161ef60565b82525050565b61e4938161ef60565b82525050565b61e4a28161ef6a565b82525050565b61e4b18161ef7a565b82525050565b61e4c08161ef7a565b82525050565b600060208201905061e4db600083018461d619565b92915050565b600060408201905061e4f6600083018561d619565b818103602083015261e508818461d8f8565b90509392505050565b600060408201905061e526600083018561d619565b61e533602083018461e48a565b9392505050565b6000602082019050818103600083015261e554818461d843565b905092915050565b600060208201905061e571600083018461d8b0565b92915050565b6000602082019050818103600083015261e591818461d8f8565b905092915050565b6000604082019050818103600083015261e5b3818561d8f8565b9050818103602083015261e5c7818461d843565b90509392505050565b6000606082019050818103600083015261e5ea818561d8f8565b905061e5f9602083018461df1f565b9392505050565b600060808201905061e615600083018761d931565b61e622602083018661e48a565b61e62f604083018561d619565b61e63c606083018461e4b7565b95945050505050565b600060808201905061e65a600083018761d931565b61e667602083018661e48a565b818103604083015261e679818561d8f8565b905061e688606083018461d619565b95945050505050565b600060208201905061e6a6600083018461d94f565b92915050565b600060208201905061e6c1600083018461d97c565b92915050565b600060208201905061e6dc600083018461d98b565b92915050565b600060208201905061e6f7600083018461d99a565b92915050565b600060208201905061e712600083018461d9a9565b92915050565b600060208201905061e72d600083018461d9b8565b92915050565b600060208201905061e748600083018461d9c7565b92915050565b600060208201905061e763600083018461d9d6565b92915050565b600060208201905061e77e600083018461d9e5565b92915050565b600060208201905061e799600083018461d9f4565b92915050565b6000606082019050818103600083015261e7b9818661da03565b905061e7c8602083018561e48a565b61e7d5604083018461e48a565b949350505050565b6000602082019050818103600083015261e7f68161da3c565b9050919050565b6000602082019050818103600083015261e8168161da5f565b9050919050565b6000602082019050818103600083015261e837818461dab1565b905092915050565b6000606082019050818103600083015261e859818661dab1565b905061e868602083018561d8b0565b61e875604083018461d8b0565b949350505050565b6000602082019050818103600083015261e897818461dcbf565b905092915050565b600060e08201905061e8b4600083018461ddaf565b92915050565b600060608201905061e8cf600083018461de7f565b92915050565b6000602082019050818103600083015261e8ef818461dec1565b905092915050565b6000602082019050818103600083015261e911818461e0ba565b905092915050565b6000604082019050818103600083015261e933818561e0ba565b9050818103602083015261e947818461dab1565b90509392505050565b600060408201905061e965600083018461e25c565b92915050565b60006101e08201905061e981600083018861e28b565b61e98f61016083018761e4b7565b61e99d61018083018661d96d565b61e9ab6101a083018561e4b7565b61e9b96101c083018461e4b7565b9695505050505050565b6000602082019050818103600083015261e9dd818461e36c565b905092915050565b6000602082019050818103600083015261e9ff818461e3bc565b905092915050565b600061ea1161ea22565b905061ea1d828261f0ba565b919050565b6000604051905090565b600067ffffffffffffffff82111561ea475761ea4661f1f2565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561ea735761ea7261f1f2565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561ea9f5761ea9e61f1f2565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561eacb5761eaca61f1f2565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561eaf75761eaf661f1f2565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561eb235761eb2261f1f2565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561eb4f5761eb4e61f1f2565b5b61eb588261f221565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600061ed0f8261ef60565b915061ed1a8361ef60565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561ed4f5761ed4e61f165565b5b828201905092915050565b600061ed658261ef7a565b915061ed708361ef7a565b92508267ffffffffffffffff0382111561ed8d5761ed8c61f165565b5b828201905092915050565b600061eda38261ef7a565b915061edae8361ef7a565b92508261edbe5761edbd61f194565b5b828204905092915050565b600061edd48261ef7a565b915061eddf8361ef7a565b92508167ffffffffffffffff048311821515161561ee005761edff61f165565b5b828202905092915050565b600061ee168261ef60565b915061ee218361ef60565b92508282101561ee345761ee3361f165565b5b828203905092915050565b600061ee4a8261ef7a565b915061ee558361ef7a565b92508282101561ee685761ee6761f165565b5b828203905092915050565b600061ee7e8261ef40565b9050919050565b60008115159050919050565b6000819050919050565b600061eea68261ee73565b9050919050565b600061eeb88261ee73565b9050919050565b600061eeca8261ee73565b9050919050565b600061eedc8261ee73565b9050919050565b600061eeee8261ee73565b9050919050565b600061ef008261ee73565b9050919050565b600081905061ef158261f2aa565b919050565b600081905061ef288261f2be565b919050565b600081905061ef3b8261f2d2565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b600067ffffffffffffffff82169050919050565b600061ef998261ef07565b9050919050565b600061efab8261ef1a565b9050919050565b600061efbd8261ef2d565b9050919050565b600061efcf8261ef7a565b9050919050565b600061efe18261ef7a565b9050919050565b600061eff38261ef7a565b9050919050565b600061f0058261ef7a565b9050919050565b600061f0178261ef7a565b9050919050565b600061f0298261ef7a565b9050919050565b600061f03b8261ef7a565b9050919050565b600061f04d8261ef7a565b9050919050565b600061f05f8261ef7a565b9050919050565b600061f0718261ef7a565b9050919050565b82818337600083830152505050565b60005b8381101561f0a557808201518184015260208101905061f08a565b8381111561f0b4576000848401525b50505050565b61f0c38261f221565b810181811067ffffffffffffffff8211171561f0e25761f0e161f1f2565b5b80604052505050565b600061f0f68261ef60565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561f1295761f12861f165565b5b600182019050919050565b600061f13f8261ef7a565b915067ffffffffffffffff82141561f15a5761f15961f165565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f50756e697368466f72536563746f72206661696c656400000000000000000000600082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b600a811061f2bb5761f2ba61f1c3565b5b50565b6003811061f2cf5761f2ce61f1c3565b5b50565b6002811061f2e35761f2e261f1c3565b5b50565b61f2ef8161ee73565b811461f2fa57600080fd5b50565b61f3068161ee85565b811461f31157600080fd5b50565b61f31d8161ee91565b811461f32857600080fd5b50565b61f3348161ee9b565b811461f33f57600080fd5b50565b61f34b8161eead565b811461f35657600080fd5b50565b61f3628161eebf565b811461f36d57600080fd5b50565b61f3798161eed1565b811461f38457600080fd5b50565b61f3908161eee3565b811461f39b57600080fd5b50565b61f3a78161eef5565b811461f3b257600080fd5b50565b6003811061f3c257600080fd5b50565b6002811061f3d257600080fd5b50565b61f3de8161ef60565b811461f3e957600080fd5b50565b61f3f58161ef6a565b811461f40057600080fd5b50565b61f40c8161ef7a565b811461f41757600080fd5b5056fea2646970667358221220da3e5e375bc96d57999beafb1184857c0289f11e2aff2688e97a772f5c65d1fe64736f6c63430008040033",
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

// GetProveDetailList is a free data retrieval call binding the contract method 0x977fdfe2.
//
// Solidity: function GetProveDetailList(bytes fileHash) view returns((address,address,uint64,uint256,bool)[])
func (_Store *StoreCaller) GetProveDetailList(opts *bind.CallOpts, fileHash []byte) ([]ProveDetail, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetProveDetailList", fileHash)

	if err != nil {
		return *new([]ProveDetail), err
	}

	out0 := *abi.ConvertType(out[0], new([]ProveDetail)).(*[]ProveDetail)

	return out0, err

}

// GetProveDetailList is a free data retrieval call binding the contract method 0x977fdfe2.
//
// Solidity: function GetProveDetailList(bytes fileHash) view returns((address,address,uint64,uint256,bool)[])
func (_Store *StoreSession) GetProveDetailList(fileHash []byte) ([]ProveDetail, error) {
	return _Store.Contract.GetProveDetailList(&_Store.CallOpts, fileHash)
}

// GetProveDetailList is a free data retrieval call binding the contract method 0x977fdfe2.
//
// Solidity: function GetProveDetailList(bytes fileHash) view returns((address,address,uint64,uint256,bool)[])
func (_Store *StoreCallerSession) GetProveDetailList(fileHash []byte) ([]ProveDetail, error) {
	return _Store.Contract.GetProveDetailList(&_Store.CallOpts, fileHash)
}

// C0xd11e680a is a free data retrieval call binding the contract method 0x96961614.
//
// Solidity: function c_0xd11e680a(bytes32 c__0xd11e680a) pure returns()
func (_Store *StoreCaller) C0xd11e680a(opts *bind.CallOpts, c__0xd11e680a [32]byte) error {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "c_0xd11e680a", c__0xd11e680a)

	if err != nil {
		return err
	}

	return err

}

// C0xd11e680a is a free data retrieval call binding the contract method 0x96961614.
//
// Solidity: function c_0xd11e680a(bytes32 c__0xd11e680a) pure returns()
func (_Store *StoreSession) C0xd11e680a(c__0xd11e680a [32]byte) error {
	return _Store.Contract.C0xd11e680a(&_Store.CallOpts, c__0xd11e680a)
}

// C0xd11e680a is a free data retrieval call binding the contract method 0x96961614.
//
// Solidity: function c_0xd11e680a(bytes32 c__0xd11e680a) pure returns()
func (_Store *StoreCallerSession) C0xd11e680a(c__0xd11e680a [32]byte) error {
	return _Store.Contract.C0xd11e680a(&_Store.CallOpts, c__0xd11e680a)
}

// CheckSectorProve is a free data retrieval call binding the contract method 0x0fece869.
//
// Solidity: function checkSectorProve((address,uint64,uint64,bytes) sectorProve, (address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) view returns(bool)
func (_Store *StoreCaller) CheckSectorProve(opts *bind.CallOpts, sectorProve SectorProveParams, sectorInfo SectorInfo) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "checkSectorProve", sectorProve, sectorInfo)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckSectorProve is a free data retrieval call binding the contract method 0x0fece869.
//
// Solidity: function checkSectorProve((address,uint64,uint64,bytes) sectorProve, (address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) view returns(bool)
func (_Store *StoreSession) CheckSectorProve(sectorProve SectorProveParams, sectorInfo SectorInfo) (bool, error) {
	return _Store.Contract.CheckSectorProve(&_Store.CallOpts, sectorProve, sectorInfo)
}

// CheckSectorProve is a free data retrieval call binding the contract method 0x0fece869.
//
// Solidity: function checkSectorProve((address,uint64,uint64,bytes) sectorProve, (address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) view returns(bool)
func (_Store *StoreCallerSession) CheckSectorProve(sectorProve SectorProveParams, sectorInfo SectorInfo) (bool, error) {
	return _Store.Contract.CheckSectorProve(&_Store.CallOpts, sectorProve, sectorInfo)
}

// CheckNodeSectorProvedInTime is a paid mutator transaction binding the contract method 0x648d225d.
//
// Solidity: function CheckNodeSectorProvedInTime((address,uint64) sectorRef) payable returns()
func (_Store *StoreTransactor) CheckNodeSectorProvedInTime(opts *bind.TransactOpts, sectorRef SectorRef) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "CheckNodeSectorProvedInTime", sectorRef)
}

// CheckNodeSectorProvedInTime is a paid mutator transaction binding the contract method 0x648d225d.
//
// Solidity: function CheckNodeSectorProvedInTime((address,uint64) sectorRef) payable returns()
func (_Store *StoreSession) CheckNodeSectorProvedInTime(sectorRef SectorRef) (*types.Transaction, error) {
	return _Store.Contract.CheckNodeSectorProvedInTime(&_Store.TransactOpts, sectorRef)
}

// CheckNodeSectorProvedInTime is a paid mutator transaction binding the contract method 0x648d225d.
//
// Solidity: function CheckNodeSectorProvedInTime((address,uint64) sectorRef) payable returns()
func (_Store *StoreTransactorSession) CheckNodeSectorProvedInTime(sectorRef SectorRef) (*types.Transaction, error) {
	return _Store.Contract.CheckNodeSectorProvedInTime(&_Store.TransactOpts, sectorRef)
}

// DeleteProveDetails is a paid mutator transaction binding the contract method 0x52e06146.
//
// Solidity: function DeleteProveDetails(bytes fileHash) payable returns()
func (_Store *StoreTransactor) DeleteProveDetails(opts *bind.TransactOpts, fileHash []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteProveDetails", fileHash)
}

// DeleteProveDetails is a paid mutator transaction binding the contract method 0x52e06146.
//
// Solidity: function DeleteProveDetails(bytes fileHash) payable returns()
func (_Store *StoreSession) DeleteProveDetails(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteProveDetails(&_Store.TransactOpts, fileHash)
}

// DeleteProveDetails is a paid mutator transaction binding the contract method 0x52e06146.
//
// Solidity: function DeleteProveDetails(bytes fileHash) payable returns()
func (_Store *StoreTransactorSession) DeleteProveDetails(fileHash []byte) (*types.Transaction, error) {
	return _Store.Contract.DeleteProveDetails(&_Store.TransactOpts, fileHash)
}

// FileProve is a paid mutator transaction binding the contract method 0x8e270530.
//
// Solidity: function FileProve((bytes,bytes,uint256,address,uint64,uint64) fileProve) payable returns()
func (_Store *StoreTransactor) FileProve(opts *bind.TransactOpts, fileProve FileProveParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "FileProve", fileProve)
}

// FileProve is a paid mutator transaction binding the contract method 0x8e270530.
//
// Solidity: function FileProve((bytes,bytes,uint256,address,uint64,uint64) fileProve) payable returns()
func (_Store *StoreSession) FileProve(fileProve FileProveParams) (*types.Transaction, error) {
	return _Store.Contract.FileProve(&_Store.TransactOpts, fileProve)
}

// FileProve is a paid mutator transaction binding the contract method 0x8e270530.
//
// Solidity: function FileProve((bytes,bytes,uint256,address,uint64,uint64) fileProve) payable returns()
func (_Store *StoreTransactorSession) FileProve(fileProve FileProveParams) (*types.Transaction, error) {
	return _Store.Contract.FileProve(&_Store.TransactOpts, fileProve)
}

// SectorProve is a paid mutator transaction binding the contract method 0x09cbe391.
//
// Solidity: function SectorProve((address,uint64,uint64,bytes) sectorProve) payable returns()
func (_Store *StoreTransactor) SectorProve(opts *bind.TransactOpts, sectorProve SectorProveParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SectorProve", sectorProve)
}

// SectorProve is a paid mutator transaction binding the contract method 0x09cbe391.
//
// Solidity: function SectorProve((address,uint64,uint64,bytes) sectorProve) payable returns()
func (_Store *StoreSession) SectorProve(sectorProve SectorProveParams) (*types.Transaction, error) {
	return _Store.Contract.SectorProve(&_Store.TransactOpts, sectorProve)
}

// SectorProve is a paid mutator transaction binding the contract method 0x09cbe391.
//
// Solidity: function SectorProve((address,uint64,uint64,bytes) sectorProve) payable returns()
func (_Store *StoreTransactorSession) SectorProve(sectorProve SectorProveParams) (*types.Transaction, error) {
	return _Store.Contract.SectorProve(&_Store.TransactOpts, sectorProve)
}

// SetLastPunishmentHeightForNode is a paid mutator transaction binding the contract method 0x3ec0f5df.
//
// Solidity: function SetLastPunishmentHeightForNode(address nodeAddr, uint64 sectorId, uint256 height) payable returns()
func (_Store *StoreTransactor) SetLastPunishmentHeightForNode(opts *bind.TransactOpts, nodeAddr common.Address, sectorId uint64, height *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SetLastPunishmentHeightForNode", nodeAddr, sectorId, height)
}

// SetLastPunishmentHeightForNode is a paid mutator transaction binding the contract method 0x3ec0f5df.
//
// Solidity: function SetLastPunishmentHeightForNode(address nodeAddr, uint64 sectorId, uint256 height) payable returns()
func (_Store *StoreSession) SetLastPunishmentHeightForNode(nodeAddr common.Address, sectorId uint64, height *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetLastPunishmentHeightForNode(&_Store.TransactOpts, nodeAddr, sectorId, height)
}

// SetLastPunishmentHeightForNode is a paid mutator transaction binding the contract method 0x3ec0f5df.
//
// Solidity: function SetLastPunishmentHeightForNode(address nodeAddr, uint64 sectorId, uint256 height) payable returns()
func (_Store *StoreTransactorSession) SetLastPunishmentHeightForNode(nodeAddr common.Address, sectorId uint64, height *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetLastPunishmentHeightForNode(&_Store.TransactOpts, nodeAddr, sectorId, height)
}

// SettleForFile is a paid mutator transaction binding the contract method 0x7b85255a.
//
// Solidity: function SettleForFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo, (uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo, (address,address,uint64,uint256,bool) detail, (address,address,uint64,uint256,bool)[] details, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting) payable returns()
func (_Store *StoreTransactor) SettleForFile(opts *bind.TransactOpts, fileInfo FileInfo, nodeInfo NodeInfo, detail ProveDetail, details []ProveDetail, setting Setting) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SettleForFile", fileInfo, nodeInfo, detail, details, setting)
}

// SettleForFile is a paid mutator transaction binding the contract method 0x7b85255a.
//
// Solidity: function SettleForFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo, (uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo, (address,address,uint64,uint256,bool) detail, (address,address,uint64,uint256,bool)[] details, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting) payable returns()
func (_Store *StoreSession) SettleForFile(fileInfo FileInfo, nodeInfo NodeInfo, detail ProveDetail, details []ProveDetail, setting Setting) (*types.Transaction, error) {
	return _Store.Contract.SettleForFile(&_Store.TransactOpts, fileInfo, nodeInfo, detail, details, setting)
}

// SettleForFile is a paid mutator transaction binding the contract method 0x7b85255a.
//
// Solidity: function SettleForFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo, (uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo, (address,address,uint64,uint256,bool) detail, (address,address,uint64,uint256,bool)[] details, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting) payable returns()
func (_Store *StoreTransactorSession) SettleForFile(fileInfo FileInfo, nodeInfo NodeInfo, detail ProveDetail, details []ProveDetail, setting Setting) (*types.Transaction, error) {
	return _Store.Contract.SettleForFile(&_Store.TransactOpts, fileInfo, nodeInfo, detail, details, setting)
}

// UpdateProveDetailList is a paid mutator transaction binding the contract method 0x27ab4434.
//
// Solidity: function UpdateProveDetailList(bytes fileHash, (address,address,uint64,uint256,bool)[] details) payable returns()
func (_Store *StoreTransactor) UpdateProveDetailList(opts *bind.TransactOpts, fileHash []byte, details []ProveDetail) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateProveDetailList", fileHash, details)
}

// UpdateProveDetailList is a paid mutator transaction binding the contract method 0x27ab4434.
//
// Solidity: function UpdateProveDetailList(bytes fileHash, (address,address,uint64,uint256,bool)[] details) payable returns()
func (_Store *StoreSession) UpdateProveDetailList(fileHash []byte, details []ProveDetail) (*types.Transaction, error) {
	return _Store.Contract.UpdateProveDetailList(&_Store.TransactOpts, fileHash, details)
}

// UpdateProveDetailList is a paid mutator transaction binding the contract method 0x27ab4434.
//
// Solidity: function UpdateProveDetailList(bytes fileHash, (address,address,uint64,uint256,bool)[] details) payable returns()
func (_Store *StoreTransactorSession) UpdateProveDetailList(fileHash []byte, details []ProveDetail) (*types.Transaction, error) {
	return _Store.Contract.UpdateProveDetailList(&_Store.TransactOpts, fileHash, details)
}

// UpdateProveDetailMeta is a paid mutator transaction binding the contract method 0xbb4e4e42.
//
// Solidity: function UpdateProveDetailMeta(bytes fileHash, (uint64,uint64) details) payable returns()
func (_Store *StoreTransactor) UpdateProveDetailMeta(opts *bind.TransactOpts, fileHash []byte, details ProveDetailMeta) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateProveDetailMeta", fileHash, details)
}

// UpdateProveDetailMeta is a paid mutator transaction binding the contract method 0xbb4e4e42.
//
// Solidity: function UpdateProveDetailMeta(bytes fileHash, (uint64,uint64) details) payable returns()
func (_Store *StoreSession) UpdateProveDetailMeta(fileHash []byte, details ProveDetailMeta) (*types.Transaction, error) {
	return _Store.Contract.UpdateProveDetailMeta(&_Store.TransactOpts, fileHash, details)
}

// UpdateProveDetailMeta is a paid mutator transaction binding the contract method 0xbb4e4e42.
//
// Solidity: function UpdateProveDetailMeta(bytes fileHash, (uint64,uint64) details) payable returns()
func (_Store *StoreTransactorSession) UpdateProveDetailMeta(fileHash []byte, details ProveDetailMeta) (*types.Transaction, error) {
	return _Store.Contract.UpdateProveDetailMeta(&_Store.TransactOpts, fileHash, details)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2561e0a.
//
// Solidity: function initialize(address _config, address _fs, address _node, address _pdp, address _sector, (uint64) proveConfig, address _proveExtra) returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts, _config common.Address, _fs common.Address, _node common.Address, _pdp common.Address, _sector common.Address, proveConfig ProveConfig, _proveExtra common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize", _config, _fs, _node, _pdp, _sector, proveConfig, _proveExtra)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2561e0a.
//
// Solidity: function initialize(address _config, address _fs, address _node, address _pdp, address _sector, (uint64) proveConfig, address _proveExtra) returns()
func (_Store *StoreSession) Initialize(_config common.Address, _fs common.Address, _node common.Address, _pdp common.Address, _sector common.Address, proveConfig ProveConfig, _proveExtra common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _fs, _node, _pdp, _sector, proveConfig, _proveExtra)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2561e0a.
//
// Solidity: function initialize(address _config, address _fs, address _node, address _pdp, address _sector, (uint64) proveConfig, address _proveExtra) returns()
func (_Store *StoreTransactorSession) Initialize(_config common.Address, _fs common.Address, _node common.Address, _pdp common.Address, _sector common.Address, proveConfig ProveConfig, _proveExtra common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _fs, _node, _pdp, _sector, proveConfig, _proveExtra)
}

// ProfitSplitForSector is a paid mutator transaction binding the contract method 0x9a19c98a.
//
// Solidity: function profitSplitForSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting) payable returns(bool)
func (_Store *StoreTransactor) ProfitSplitForSector(opts *bind.TransactOpts, sectorInfo SectorInfo, nodeInfo NodeInfo, setting Setting) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "profitSplitForSector", sectorInfo, nodeInfo, setting)
}

// ProfitSplitForSector is a paid mutator transaction binding the contract method 0x9a19c98a.
//
// Solidity: function profitSplitForSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting) payable returns(bool)
func (_Store *StoreSession) ProfitSplitForSector(sectorInfo SectorInfo, nodeInfo NodeInfo, setting Setting) (*types.Transaction, error) {
	return _Store.Contract.ProfitSplitForSector(&_Store.TransactOpts, sectorInfo, nodeInfo, setting)
}

// ProfitSplitForSector is a paid mutator transaction binding the contract method 0x9a19c98a.
//
// Solidity: function profitSplitForSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting) payable returns(bool)
func (_Store *StoreTransactorSession) ProfitSplitForSector(sectorInfo SectorInfo, nodeInfo NodeInfo, setting Setting) (*types.Transaction, error) {
	return _Store.Contract.ProfitSplitForSector(&_Store.TransactOpts, sectorInfo, nodeInfo, setting)
}

// PunishForSector is a paid mutator transaction binding the contract method 0xfda59351.
//
// Solidity: function punishForSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 times) payable returns()
func (_Store *StoreTransactor) PunishForSector(opts *bind.TransactOpts, sectorInfo SectorInfo, nodeInfo NodeInfo, setting Setting, times uint64) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "punishForSector", sectorInfo, nodeInfo, setting, times)
}

// PunishForSector is a paid mutator transaction binding the contract method 0xfda59351.
//
// Solidity: function punishForSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 times) payable returns()
func (_Store *StoreSession) PunishForSector(sectorInfo SectorInfo, nodeInfo NodeInfo, setting Setting, times uint64) (*types.Transaction, error) {
	return _Store.Contract.PunishForSector(&_Store.TransactOpts, sectorInfo, nodeInfo, setting, times)
}

// PunishForSector is a paid mutator transaction binding the contract method 0xfda59351.
//
// Solidity: function punishForSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 times) payable returns()
func (_Store *StoreTransactorSession) PunishForSector(sectorInfo SectorInfo, nodeInfo NodeInfo, setting Setting, times uint64) (*types.Transaction, error) {
	return _Store.Contract.PunishForSector(&_Store.TransactOpts, sectorInfo, nodeInfo, setting, times)
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
	ProvLevel   uint8
	Size        uint64
	IsPlots     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreateSectorEvent is a free log retrieval operation binding the contract event 0xc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb.
//
// Solidity: event CreateSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId, uint8 provLevel, uint64 size, bool isPlots)
func (_Store *StoreFilterer) FilterCreateSectorEvent(opts *bind.FilterOpts) (*StoreCreateSectorEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "CreateSectorEvent")
	if err != nil {
		return nil, err
	}
	return &StoreCreateSectorEventIterator{contract: _Store.contract, event: "CreateSectorEvent", logs: logs, sub: sub}, nil
}

// WatchCreateSectorEvent is a free log subscription operation binding the contract event 0xc3aecd94f85ba0fad5908b524eaaa61cce4b92e04937d32d159f1d5bd8d029fb.
//
// Solidity: event CreateSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId, uint8 provLevel, uint64 size, bool isPlots)
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
// Solidity: event CreateSectorEvent(uint8 eventType, uint256 blockHeight, address walletAddr, uint64 sectorId, uint8 provLevel, uint64 size, bool isPlots)
func (_Store *StoreFilterer) ParseCreateSectorEvent(log types.Log) (*StoreCreateSectorEvent, error) {
	event := new(StoreCreateSectorEvent)
	if err := _Store.contract.UnpackLog(event, "CreateSectorEvent", log); err != nil {
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
	NodeAddr    common.Address
	Volume      uint64
	ServiceTime uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegisterNodeEvent is a free log retrieval operation binding the contract event 0x2668f2f26843f9357774392703ed879a5a5ca781dfb60354b056070fed6db808.
//
// Solidity: event RegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr, address nodeAddr, uint64 volume, uint64 serviceTime)
func (_Store *StoreFilterer) FilterRegisterNodeEvent(opts *bind.FilterOpts) (*StoreRegisterNodeEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "RegisterNodeEvent")
	if err != nil {
		return nil, err
	}
	return &StoreRegisterNodeEventIterator{contract: _Store.contract, event: "RegisterNodeEvent", logs: logs, sub: sub}, nil
}

// WatchRegisterNodeEvent is a free log subscription operation binding the contract event 0x2668f2f26843f9357774392703ed879a5a5ca781dfb60354b056070fed6db808.
//
// Solidity: event RegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr, address nodeAddr, uint64 volume, uint64 serviceTime)
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

// ParseRegisterNodeEvent is a log parse operation binding the contract event 0x2668f2f26843f9357774392703ed879a5a5ca781dfb60354b056070fed6db808.
//
// Solidity: event RegisterNodeEvent(uint8 eventType, uint256 blockHeight, address walletAddr, address nodeAddr, uint64 volume, uint64 serviceTime)
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
