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

// PocProve is an auto generated low-level Go binding around an user-defined struct.
type PocProve struct {
	Miner    common.Address
	Height   *big.Int
	PlotSize uint64
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

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"FileProveFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileProveNotFileOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeSectorProvedInTimeError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SectorProveFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"CheckNodeSectorProvedInTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteProveDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"NodeWallet\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"}],\"internalType\":\"structFileProveParams\",\"name\":\"fileProve\",\"type\":\"tuple\"}],\"name\":\"FileProve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetPocProveList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"Height\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"PlotSize\",\"type\":\"uint64\"}],\"internalType\":\"structPocProve[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"GetProveDetailList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ChallengeHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"}],\"internalType\":\"structSectorProveParams\",\"name\":\"sectorProve\",\"type\":\"tuple\"}],\"name\":\"SectorProve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"SetLastPunishmentHeightForNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail\",\"name\":\"detail\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"details\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"}],\"name\":\"SettleForFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"details\",\"type\":\"tuple[]\"}],\"name\":\"UpdateProveDetailList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveDetailNum\",\"type\":\"uint64\"}],\"internalType\":\"structProveDetailMeta\",\"name\":\"details\",\"type\":\"tuple\"}],\"name\":\"UpdateProveDetailMeta\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0xd11e680a\",\"type\":\"bytes32\"}],\"name\":\"c_0xd11e680a\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ChallengeHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"}],\"internalType\":\"structSectorProveParams\",\"name\":\"sectorProve\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"}],\"name\":\"checkSectorProve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"getPocProve\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"Height\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"PlotSize\",\"type\":\"uint64\"}],\"internalType\":\"structPocProve\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractFileSystem\",\"name\":\"_fs\",\"type\":\"address\"},{\"internalType\":\"contractNode\",\"name\":\"_node\",\"type\":\"address\"},{\"internalType\":\"contractPDP\",\"name\":\"_pdp\",\"type\":\"address\"},{\"internalType\":\"contractSector\",\"name\":\"_sector\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"SECTOR_PROVE_BLOCK_NUM\",\"type\":\"uint64\"}],\"internalType\":\"structProveConfig\",\"name\":\"proveConfig\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"times\",\"type\":\"uint64\"}],\"name\":\"punishForSector\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"Height\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"PlotSize\",\"type\":\"uint64\"}],\"internalType\":\"structPocProve\",\"name\":\"prove\",\"type\":\"tuple\"}],\"name\":\"putPocProve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506201165880620000236000396000f3fe6080604052600436106200010a5760003560e01c80637b85255a11620000975780639908f2bf11620000615780639908f2bf1462000365578063bb4e4e4214620003a9578063f76031f614620003d7578063fda593511462000405576200010a565b80637b85255a14620002975780638e27053014620002c55780639696161414620002f3578063977fdfe21462000321576200010a565b806327ab443411620000d957806327ab443414620001df5780633ec0f5df146200020d57806352e06146146200023b578063648d225d1462000269576200010a565b806309cbe391146200010f5780630edec0f2146200013d5780630fece869146200016d5780631581d54514620001b1575b600080fd5b3480156200011c57600080fd5b506200013b60048036038101906200013591906200f22c565b62000425565b005b3480156200014a57600080fd5b5062000155620014ad565b60405162000164919062010617565b60405180910390f35b3480156200017a57600080fd5b506200019960048036038101906200019391906200f271565b620014e0565b604051620001a891906201065f565b60405180910390f35b348015620001be57600080fd5b50620001dd6004803603810190620001d791906200f133565b62002a0a565b005b348015620001ec57600080fd5b506200020b60048036038101906200020591906200ee1d565b62002bc8565b005b3480156200021a57600080fd5b506200023960048036038101906200023391906200ece5565b62002e16565b005b3480156200024857600080fd5b506200026760048036038101906200026191906200edd8565b62002f0e565b005b3480156200027657600080fd5b506200029560048036038101906200028f91906200f2e4565b62003032565b005b348015620002a457600080fd5b50620002c36004803603810190620002bd91906200efc5565b62003c0e565b005b348015620002d257600080fd5b50620002f16004803603810190620002eb91906200f07d565b62004a51565b005b3480156200030057600080fd5b506200031f60048036038101906200031991906200edac565b62007dcb565b005b3480156200032e57600080fd5b506200034d60048036038101906200034791906200edd8565b62007dce565b6040516200035c91906201063b565b60405180910390f35b3480156200037257600080fd5b506200039160048036038101906200038b91906200eca4565b620082b3565b604051620003a0919062010994565b60405180910390f35b348015620003b657600080fd5b50620003d56004803603810190620003cf91906200ee90565b62008491565b005b348015620003e457600080fd5b50620004036004803603810190620003fd91906200eeea565b620085a1565b005b6200042360048036038101906200041d91906200f1a4565b62008a5c565b005b620004537f52a98d01f32c4d1dce46f9b818475cd3ab23a96c883555e656257c492038ad9960001b62007dcb565b620004817fd22b1ecbf0339db49703a8bb58cfff0175cd5cc42df043e8d0ffdaf89ecbcf8f60001b62007dcb565b620004af7f2373d625948d3e524d519c1c5900b4e36bd26cf8c09ec83c1fa6324754e3b3b360001b62007dcb565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633778febe83600001516040518263ffffffff1660e01b8152600401620005129190620105c6565b60e06040518083038186803b1580156200052b57600080fd5b505afa15801562000540573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200056691906200f0c2565b9050620005967f8d9c2a0120886a5aaa9f18eda54c9823dd299a10caa7075c36d9be66ad06c5c560001b62007dcb565b620005c47f18cf4d2d377b061f5aafb5d95206df70e96d7067365e0f26bb61d0c3831f3ed260001b62007dcb565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d76040518060400160405280866000015173ffffffffffffffffffffffffffffffffffffffff168152602001866020015167ffffffffffffffff168152506040518263ffffffff1660e01b81526004016200065f919062010a34565b60006040518083038186803b1580156200067857600080fd5b505afa1580156200068d573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190620006b891906200f15f565b9050620006e87f3ca6f5846a74b648927d994564d30683a2d3322d0015a07f85e341beff83ce0360001b62007dcb565b620007167fbed21c8f3ccc61094dd246c83b8a57c26155c838406e9cc88fa1905d7805cddd60001b62007dcb565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156200078157600080fd5b505afa15801562000796573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620007bc91906200f310565b9050620007ec7f92a263e0f7fce99ce10b411c0d7b506178e3523f76b37f933e152d85c9795ef860001b62007dcb565b6200081a7fc53207b556a82b61e8d55d45fafdff2d21c9ea09972cb69194220a29e8ca9dc060001b62007dcb565b8160c00151431015620008c357620008557fd02960f429ea5d0b6e78bc10e9014b4fc8c6dcb1b83126084f5fcc8bc23a443160001b62007dcb565b620008837fa2ce1678e85ebf4ac1ff55de11cb693db4749dd55a1498f2c7e8de6a44361d5d60001b62007dcb565b60016040517fcab88a3c000000000000000000000000000000000000000000000000000000008152600401620008ba91906201075e565b60405180910390fd5b620008f17f2e3dd35d1e93dd2d12c9b81cd837b47a72457f8114a3eb391951813ea92accc660001b62007dcb565b6200091f7f8a288a30e53d223ec12f697555b41ffd362de3bbd559268cd7a8a615b74776ac60001b62007dcb565b6200094d7f1c11be8c7833a9c29a0286dfeb1d79d08437862d3bf0f7ec99c3ae64ae98147f60001b62007dcb565b8160c00151846040015167ffffffffffffffff161462000ab357620009957f5aca5d1a0c017caedfc17679100858029ff0b9ca168d390aeb2293f512518c6b60001b62007dcb565b620009c37f5ff17198d2f86688937e2d66db6d6961886360239ca70dceccbcf7a31a55f38760001b62007dcb565b620009f17fdfe3d37a596e1c5737f7b28325f95f4d5d809b8c54e87960e1a3ff0d969b2dbb60001b62007dcb565b62000a456040518060400160405280601681526020017f4368616c6c656e6765486569676874206572726f723a00000000000000000000815250856040015167ffffffffffffffff168460c001516200908f565b62000a737f39f2f0c80339f8a20e89e87f7fbb315f2b58c353bcc2c19a89fe49659f291ee760001b62007dcb565b60026040517fcab88a3c00000000000000000000000000000000000000000000000000000000815260040162000aaa91906201077b565b60405180910390fd5b62000ae17f5165ee976b9836bb433900cc23c73645533d1b32298a133155b63d852b6adf5260001b62007dcb565b62000b0f7f75386f3147cbfddd7db9fb0aab2963b04d024f06705cbb808d610a48b5fd704560001b62007dcb565b62000b3d7f4b1e5305eef31cea3532f31ad5c5cd10b4d8a200aa0ee20349d6142e6c47ef5360001b62007dcb565b600062000b4b8584620014e0565b905062000b7b7f7544975bfe18be3fdb57bee43c9d6234c3daa5d6ac38ef8e45ac31b5198a5d9960001b62007dcb565b62000ba97fcd7f3eaa18eba7044f49e228f851d574c0d29205ad3b998a960153c00da7a9ad60001b62007dcb565b8062000cb65762000bdd7fb64a679f7f9fe8eab2a03c8dfd087abe150a5141bce2b4a58102856b669eaaf860001b62007dcb565b62000c0b7fcd57a7f1ffedcb2b9253253e6e9e7bb214c3dcdf1d92f558f6b50734fc6505fb60001b62007dcb565b62000c397fbe8ffdc8997715b0254950b4369f5e0e8bd5984036546376c08449e966e6157760001b62007dcb565b62000c48838584600162008a5c565b62000c767f645769f1b21d4064fc12919ec2fb385d13645ce2558dd90b6fe5a3dfae25d5b960001b62007dcb565b60036040517fcab88a3c00000000000000000000000000000000000000000000000000000000815260040162000cad919062010798565b60405180910390fd5b62000ce47fecaaa99bd0de2f39619fdee68ee6b4c349ad5ab026aff1441c6ea5e3198df36160001b62007dcb565b62000d127fd85ab96555044395a91bf12f8bf847ca16406672c7c481515563208519e1610b60001b62007dcb565b62000d407f6a938f362bd7d8257a50727c623165992f1b0524862b64ed2ee9f45384d35c4b60001b62007dcb565b600062000d4f84868562009132565b905062000d7f7f591f24fe2a938e4042c98fcfd965064c961d18786652495ae6ac333395c80e8b60001b62007dcb565b62000dad7fdacca2a8e2ef852f9e44a5f795207510b2b71301d7e5aaf300d92ca1cc9facb460001b62007dcb565b8062000e4f5762000de17fe0d5637626943fd49d96fd5bcb460320535009f789ae5577712a2aa96717909460001b62007dcb565b62000e0f7f1b996e2a735db9f17042d4545f82b10ea4e790b7b452dea00dc75788cf97d60560001b62007dcb565b60056040517fcab88a3c00000000000000000000000000000000000000000000000000000000815260040162000e469190620107d2565b60405180910390fd5b62000e7d7f97b48fc267e77fccb7dc43ecad070e29871320dc1ee055020e32b6c765e2b90460001b62007dcb565b62000eab7fcb9f9495d337eafb10bef4c77d80e567fea2e183ec666acde4c5d168d8d7d0f260001b62007dcb565b62000ed97ff0c32d95a6c03a442955875459a83e98a0350059d8a63ecf9801ba50405292c760001b62007dcb565b60008460a00151141562000f815762000f157f54e1364924f69afe3c273694aaac68a5e40081dfa2e509782de53e39ad0f098d60001b62007dcb565b62000f437fbacbf44b24f6cd41e00cc77f0314eca6ebce1d84ad336dc47539c9c4846f924860001b62007dcb565b62000f717f554c4a82174e6be72a24357c25c018c08e3d314583363cf44b5281f8d7f53ea160001b62007dcb565b438460a001818152505062000fb0565b62000faf7f9f2beb20d88aedd2d4097a2ca20d44a8dda8df6db34ba2acae9ed5263c30f35460001b62007dcb565b5b62000fde7f1d9e55adcbced2e9f5cac09b59669946deaca4f4a564e6f589e042ffabaed4ba60001b62007dcb565b6200100c7f42c5d7b397ea59074ef3017405102011900008d90d8cd9f4f2516df583a76ada60001b62007dcb565b8260c0015167ffffffffffffffff164362001028919062010e62565b8460c00181815250506200105f7f3d97537de75785d84f26099d104b642fdf91e296b9a0e768e2b70e9e197775ed60001b62007dcb565b6200108d7f16ccb16e4a41378b9ed43c80b0950bc4a8f402d2a8c271fe0023792182edea3260001b62007dcb565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632384a6aa856040518263ffffffff1660e01b8152600401620010ea9190620109d5565b600060405180830381600087803b1580156200110557600080fd5b505af11580156200111a573d6000803e3d6000fd5b505050506200114c7f1c466f470101914b7e7bc84279416ebfdb0ab50749d0e4fdb910a00d4a02938560001b62007dcb565b6200117a7f99d4b32a0fcf8eb9e87f21d232efc4066a946d986e48a5890d9d46aed34c6f2060001b62007dcb565b8361014001516200122157620011b37fea609f7d3df430484a145960f1e9b9da1b091f32b5409566f433cdcd51c401ca60001b62007dcb565b620011e17f5c90e46ad93e0975126af0d56eec156288b5d96d07c049d32bb1097475633ae960001b62007dcb565b60046040517fcab88a3c000000000000000000000000000000000000000000000000000000008152600401620012189190620107b5565b60405180910390fd5b6200124f7f1305f6ad2c8ed618af6dd9524578e0559d2b815ca252a89d3004dc3635bff02c60001b62007dcb565b6200127d7fb5e3f4b34a63fc3d95607c604b0dfa68a2486479be7e63fbaf7686eb15230a8460001b62007dcb565b620012ab7fb77c6d79f5d1109995f72a7e3c86b8951b9b4faa8270563931c106f6dce1797560001b62007dcb565b6000620012bd856000015143620082b3565b9050620012ed7f9177a2d82a45530837b2acf20453430f2e90a7e78d41f405a9bfaa84d7f6134260001b62007dcb565b6200131b7fe535de601ad48db66f0ef4ec8e5bba7f74bc5d082a4e6980feff9ff4416d6ba660001b62007dcb565b43816020018181525050620013537fc12372f42e1d5ff8768d8c4dc009632e8022f979ea30eb702760d054b58a135160001b62007dcb565b620013817f05f9c76f7747b746e60dae86b9db91d0e48f111cca695d3b1b0bb325ac71b1d960001b62007dcb565b8460000151816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050620013eb7f4fb8f8c4788dbcc5195d53a1cf7d202c648a7c664efe112c305bbf49a34d63db60001b62007dcb565b620014197fe693f4fb98ddf07777f2eb3ee2b0900295190f930c35d4afd2dcf2a7d853a6f860001b62007dcb565b8460600151816040019067ffffffffffffffff16908167ffffffffffffffff16815250506200146b7f4712799cc8cd2c0818ab0e574c6c30dfe0c62331140eb677074c0ab50abdf37f60001b62007dcb565b620014997f65273b12e8392db757daa3be4b80cf456a07f17b5852cd0c5c9caf54d24b016d60001b62007dcb565b620014a48162002a0a565b50505050505050565b6060620014dd7faf0cd1c18a1f9dda1ed0c761220fb00a15c1fb1903a46f67d2a1c23a1af5014560001b62007dcb565b90565b6000620015107fbe35e717fbf53305d666bac8f52cb130e353c413ccb2bc35f0f45fd51e29517e60001b62007dcb565b6200153e7f4dd39864b9b8080e1da2ae9b91eb7cc506b0e6bf978c51bdd2f65f3b37448d7460001b62007dcb565b6200156c7f8a1ceb46dd0440ac7bbb59994ab900420acce918a53607b148c3026d28c9a47860001b62007dcb565b60606200159c7feaea0371208a24c73a528452c54a53cf1e311506ad5d7fdef6386cb85ba37b4d60001b62007dcb565b620015ca7f13c1105e0f99568fae82f20d912f7c9e6769a4cc4bec876b96d962122434328a60001b62007dcb565b620015d46200cb8e565b620016027f3d525515d89955f64ede1165cc9b7ef3a85e313f9049eea99b689b72ac6ac57160001b62007dcb565b620016307fecbaef36c300f880b3caff55a3c2a7733aa2690071b4515c4de8e544678f9a9460001b62007dcb565b8460000151816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506200169a7f4ea2041baae17fecea2d40ed90784c7c98d1a836b4b30ab6673d83e31b5c58fb60001b62007dcb565b620016c87fd4d4c842e209b216e1bb3ccf9e34b9060f390c769760d60978c547f24ad8f85760001b62007dcb565b818160200181905250620016ff7fe1b000915fef0abe3427ecd2eaf95b141809eb1fe4ce13f2378d98cb275d802b60001b62007dcb565b6200172d7f289cba5f2dd1d2cc91fb01d2cdb5041976776ad2d9f483892aa5e569c1613eb360001b62007dcb565b8360e00151816040019067ffffffffffffffff16908167ffffffffffffffff16815250506200177f7fd8e419049fc94571543e014f6d982cf5bb0bdeea0bb04c79cf5962a9fbfbb7c260001b62007dcb565b620017ad7fbfd117ffa83c488bd1f481b34e5558914c3a913f66520713f86dd29a4d0d77a460001b62007dcb565b600460149054906101000a900467ffffffffffffffff16816060019067ffffffffffffffff16908167ffffffffffffffff1681525050620018117f4d4cde182772095554d79aebb92700ee08a346db41720244c205d4adcb3716f060001b62007dcb565b6200183f7f04521ddc36d77af77189d9ffe57d4e3521fe3dc0fcf8a1631a0661a437feb70e60001b62007dcb565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fd7c9808836040518263ffffffff1660e01b81526004016200189e919062010953565b60006040518083038186803b158015620018b757600080fd5b505afa158015620018cc573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190620018f791906200ed3b565b9050620019277fdacc2bd7ce38720b67d183371caaa0f24e60a62bc21edb631d516c4aa6d02cf560001b62007dcb565b620019557fd89f9d94a68689eae8e69ac3c4f15de1b8d9e8a56021d0e20065d5ec7c6f4ce660001b62007dcb565b6200195f6200cbe0565b6200198d7fa09e05b3cf073ed2effb9d30e2087b389c12ed23e4405723e38bea1bb5ef425a60001b62007dcb565b620019bb7f3af92384b91affa8ca1013ed8d6384d9a4976c568badca5950099c54a67ad3d360001b62007dcb565b620019c56200cc2a565b620019f37f12e7f9e03e2a39ee339815411dc12c5fca5ffa3fa0e043598555658db050000960001b62007dcb565b62001a217ffee56e7255445e9958b6ad3a045e09cb3fb51e59de0c24565fdb858e896017a160001b62007dcb565b86816000018190525062001a587f115ad34b2d05efbb8921a2f385d96f45bf0c89fb8605958e61c05fc7f30ccbdd60001b62007dcb565b62001a867f2f082e63060b7670254433dd60679cd1280710eb974f7d3c9532e0539ad5a08460001b62007dcb565b81816040018190525062001abd7f6aa3e9023062aa0c35eefed3c7af2ea727867554a36abbccdeda3d71f094bc8460001b62007dcb565b62001aeb7f8f02e60abfc8d672123b33dcd59e26cb5397af2ca7f90fe16b5fc7cfd48e628660001b62007dcb565b62001af56200cc5b565b62001b237f5d84eae77ba7510ea5f393a73b9a1b0b7cb06bc1f532675b3343d859a78657e760001b62007dcb565b62001b517f83899ccfcf976e2818ec29708a2771fa2e985a297bde71c7b81a534fce97fe8160001b62007dcb565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639f9ca644836040518263ffffffff1660e01b815260040162001bae9190620109b1565b60006040518083038186803b15801562001bc757600080fd5b505afa15801562001bdc573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019062001c0791906200f0ee565b905062001c377f9a8541901b452bc70728d98109b1e82273ad10f010480177cb67092ceaee87af60001b62007dcb565b62001c657f09250155911b4c74ff4431abe4bd8cd822933e21869b24fdc15577660f627eaa60001b62007dcb565b8060c0015162001d095762001c9d7fc08591f3432c3a62392ee2fd31ed91e89aab59c0bd7c39058b27b35a64b51eb060001b62007dcb565b62001ccb7f376874d940ba3b3327d1c15313f9a3a69540f30235a4c730951e9a8916c1064060001b62007dcb565b62001cf97fa3927e45c75163765bfa0dce1318a3859e2be6f2c4ad9f213efe7c9e1513625360001b62007dcb565b6000965050505050505062002a04565b62001d377f859a62f32495d1fd93d54d1598f5467ccf90a0d4745fdb590c9f61ec71043c7360001b62007dcb565b62001d657fdce8d1da1d99fd386ab167fa2b785a3e5206f2a98ebfe8057d868fe06b7da9d360001b62007dcb565b62001d937f4c9eee751618169de62f50e7100004ec3ca864fc2c07fa4569082fedc4b2fcef60001b62007dcb565b62001d9d6200cca2565b62001dcb7f11b9140c3e27ef65e382d373090cedaf5641d10ac6041a42fffc8f275de5f32a60001b62007dcb565b62001df97fb9749b1b62eb9df4cc4d5db5561bbed52b29267984ef91238da69940dae0a80a60001b62007dcb565b6000816000019067ffffffffffffffff16908167ffffffffffffffff168152505062001e487f5269af5f5b22ef5fccdd7a87dd08ee52ef1aaa2f4dd88343084786fb1d94315760001b62007dcb565b62001e767feccc78ac89a1f98f7e2d7258c43ee9a828bf86bdd305ee59dcb785fed0738b7060001b62007dcb565b8360400151816020018190525062001eb17f454e0214d681b4b63f724d326e0d0661391c13735c8c679578f4c07df072be0560001b62007dcb565b62001edf7f9f9b420cc110d1fd67b11a835b85e3b7cc1b6eb3aacec91a13c13cffbbec32bf60001b62007dcb565b8160000151816040018190525062001f1a7f723824323ebac42a4bcc38fd4bead264152f828c68c4a13463b8ac0c79e6d5c460001b62007dcb565b62001f487feedbd47c367f9517f3eb649e940603c7aaf532ed8668a0c65e087e949531d2dc60001b62007dcb565b8160200151816060018190525062001f837f2958981ffff30cd568bf95f7d1ca2c4b2c1b6fde5edbce4b332cf1d73e820e4b60001b62007dcb565b62001fb17f31cf12c99ed3d404486649dffddab0c5edbaf5d7b0cfcc5e3cbe65768dc120eb60001b62007dcb565b8160400151816080018190525062001fec7f8b61f31a21fa507d407d3d6ddf1a19d475d3ec4d06e0a0f189cac7116189759f60001b62007dcb565b6200201a7fd256f4c73beb4f11ac24f614165c4af0800927c2305646647ac691abd475130860001b62007dcb565b81606001518160a00181905250620020557f66fbd2a524fb997c3e82cdf94e21f92099225be9ec97b1acb720cb70fa73062260001b62007dcb565b620020837fb2c6872138be6dbf223b08fedfe35a0ee33e775c0723f4f5e54d5da01abc99a160001b62007dcb565b81608001518160c00181905250620020be7ff02e33d5bd52057e47723bbba4cebd153d29043cd2a806046a64a5277038b85c60001b62007dcb565b620020ec7fd49fcf361432ebf1464a00416ba681cc70ea122828b75352bd267fbcf752162860001b62007dcb565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b6ddeca0836040518263ffffffff1660e01b81526004016200214b919062010ad7565b60206040518083038186803b1580156200216457600080fd5b505afa15801562002179573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200219f91906200ed80565b9050620021cf7fc7a94e088072a7ea4f206a8b5d25a5fd8f54f3eb9c6384b22b87d4884860720b60001b62007dcb565b620021fd7f5bb8fb295deeb9747ba42e5b8f61cb12a270563f7533c4dadddbd091d6b5379260001b62007dcb565b806200229f57620022317fbe878fe0d7fbc9144ed479938f48fbc5e4816f2b3710c56f01f850ba3afe9d4d60001b62007dcb565b6200225f7fabbf519eeb13113fd93651e1a8e4ea6c317ada15cb7c019ee4ad2f8314476d7260001b62007dcb565b6200228d7fb84fffaacf916ea99b2243c7375633655735b3d17eff16b5fe2dac58536097c160001b62007dcb565b60009850505050505050505062002a04565b620022cd7ff03fa4ef51b4bae369da51a08838fc3f826c12d982adb6499887f8143d76657760001b62007dcb565b620022fb7f354a94e6eda2abc178686d2918ba9c31aa02b5d03b2a287a9a6dee230152878a60001b62007dcb565b620023297f5744ba7f8c7b4d876fb4689dabce586c6b6fef0f9d26a79d1a5d5b08adac96d560001b62007dcb565b896101400151156200296c57620023637fe130323275832d315616de68598c18ed6050c94e4f115bf999210f859b379ef660001b62007dcb565b620023917fe41f63410fc8503ed943f0185754934b8e63e8631eac1173bd4ecde15898875d60001b62007dcb565b620023bf7fe1ebd55178cddbfd96e1ef8557d5da40f14d46df7da01e5a30b3f6096efd1d8c60001b62007dcb565b8260a001516102a001511580620023ed575060008360a001516102c001516040015167ffffffffffffffff16145b156200242757620024217f0505be9033dce48c9d5e8375e60a73f9dd9e45437f681933002bd024137f9cdb60001b62007dcb565b62002456565b620024557f8c1f0d7c0b08480def4b8fdd0cef311dd7d1c845a2a14f05d673835cc1e2558c60001b62007dcb565b5b620024847fa703b1219893c972323e8d2bb20dda8925670fd58b18243ad8c778d147043d0a60001b62007dcb565b620024b27f38a6f9c257f256361a195953524f12720c70caff12dde71d7e20bfe6d84c1bef60001b62007dcb565b620024bc6200cce9565b620024ea7fede461ac5e07ef51deae8fefde97fab96c4518b69c8e0cb87df02ef4bfce36fc60001b62007dcb565b620025187feac4e28cbd209b0449c70e3654f4599cb45bb8054f16def61024a8a870c3961560001b62007dcb565b8360a001516102c001518160000181905250620025587ff9bbb499be44c6918fec029501af4ead958731569dce56b4f20cf142ec172c8f60001b62007dcb565b620025867f54e21210838fd1158e4003e08201cf9eebcf55ac628e863343adb2141ed5f5bb60001b62007dcb565b8560a001518160200181905250620025c17f498e06133caed9b6d7bcb977f3c54ca71e1877105cdf61f5b4db67b9e195f9e860001b62007dcb565b620025ef7fe2cefe2c1e76bbb20e9adf9f2b2e8bb20dc1d43b5df44b925eb29fa11823106e60001b62007dcb565b600087511115620026f657620026287f5ef42882416a805bc8f4c878f12c296c429ce6534ecdeaf39d358f9e9fc0b76260001b62007dcb565b620026567f1b5e1b499a99c47da6c323d400ddfc5a427cbf9fff81f40c9296a33898c42abf60001b62007dcb565b620026847f86057a41dab64bff135b7ffa61a466f7b16ab0394a8a16422bef6c707555b02b60001b62007dcb565b86600081518110620026bf577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516000015163ffffffff16816040019067ffffffffffffffff16908167ffffffffffffffff168152505062002725565b620027247fe123f61ff8622bfd89a988931bffb9f89e253f14353129aff641e050c5017a0460001b62007dcb565b5b620027537fd8781b1a6b66a879aa918331fb247b12cc9048f302d89869dbdb3094e6a1ff7d60001b62007dcb565b620027817fcce7522ca9c2b75bacded5bec499468412268ac78ef12901aadde8221e73162760001b62007dcb565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e19aeff836040518263ffffffff1660e01b8152600401620027e0919062010ab3565b60206040518083038186803b158015620027f957600080fd5b505afa1580156200280e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200283491906200ed80565b9050620028647f566dfa7350c0db209aa46f7d40032eabb8e380cc9979c1f5679eed25280c1f6560001b62007dcb565b620028927fc5f8cf93a32278525218c4e16b1ac6b546ae9f5ae963304ba2371d3287c49c0c60001b62007dcb565b806200293657620028c67ff9baa8263878341f897a245415407a43004546e8e4645784372ff855749ea8fb60001b62007dcb565b620028f47f30ce5e10943a966f585bf4df9b137b397a1a4b4b73f402cbf2f1db0cd4a98a3560001b62007dcb565b620029227fed76efb3fb31a856bd19d83eaa21a5f57408046ab5da1bd68f34e5957fb01b7d60001b62007dcb565b60009a505050505050505050505062002a04565b620029647f64cb5bef198db57354022901898704fd8bef66fcad9072d7523c535a2697d65260001b62007dcb565b50506200299b565b6200299a7fdd103d623be02d735e31347324c0c514bf10cf435c9cb66f3efd209f062b027060001b62007dcb565b5b620029c97feb71a0685f7b60e74fc77cb558af72e4858175b0fd4647fdf08aab08159a3e9660001b62007dcb565b620029f77f51b2b8fad82f352bc9dcfb6a39b5fac9f9689aff7a827646f51afb82de21eecf60001b62007dcb565b6001985050505050505050505b92915050565b62002a387fd667bd8aa6e78b4bcb94f58282d05ebc225155dc32258d7c025d7a2f74c13d2960001b62007dcb565b62002a667f6ec5c30b4fc512ffc63febfce60ce11cdb4593c8da489eb82f839ef5d39c9fcc60001b62007dcb565b62002a947f58bdeae4e12a1931b4c32b12443d5afa603c5c7ee7b9b3fbea255b22db864e8e60001b62007dcb565b60008160000151826020015160405160200162002ab392919062010564565b604051602081830303815290604052905062002af27f39fce1ca179ec1983a61703eec8edc41f159f563d7e052baf082c8c56704c70260001b62007dcb565b62002b207f0b8768cb94c7135128d0dbb31ba32317a8b1129cc242d3dbac6cd625bafa7e0660001b62007dcb565b8160088260405162002b339190620105ad565b908152602001604051809103902060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050505050565b62002bf67f278bdf5abe17a911ce5505ac97c0fbfc83916f3eeeed2bc10f648485a38bc86e60001b62007dcb565b62002c247fa8055c7b7986afe080c0e939e12a285369d5ef65c99ccf96b50ec8f6a3ec921a60001b62007dcb565b62002c527f144a7a75ded974443a7b4001cee160d4bd0f61babb30f7fbc8bb44ed0a62826d60001b62007dcb565b600060058360405162002c66919062010594565b9081526020016040518091039020905062002ca47f6560b0d97e7169eede6ae4b747136736170e31e0f5314264332309eaf5e4e9f360001b62007dcb565b62002cd27f0b32ed3c0e1f9da2f255bc41bdfea3a0049a2d5da59811d6a17f5502b1ef0ca260001b62007dcb565b60005b825181101562002e105762002d0d7fe4d30d60a4489deb682edfb5fb66faf4a976d7f7546cad97a12b7bd79861c98d60001b62007dcb565b62002d3b7f318899e60312fac39f1ea2bb39656137a3755e90caacb1e56f60128c07e6b68760001b62007dcb565b600083828151811062002d77577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151905062002daf7fe14529329d48899dd309a680836606b4f891d20feeb828e7a1565c3ba8b1a91060001b62007dcb565b62002ddd7fd1c782eeb1da7c6376a568b4d6174b9680865644c023d43eed15476c06227bb360001b62007dcb565b62002df88160000151828562009fd09092919063ffffffff16565b5050808062002e079062011295565b91505062002cd5565b50505050565b62002e447f208fe3968fbd6dcd6a0e456bbece4e03dcebc7ff7c0f1741dba23a776009f04760001b62007dcb565b62002e727f4f48310bf022a0521f33891749f17d12ed7ced384777aa01ca314f2e2ff9ed3460001b62007dcb565b62002ea07f03a9cb743a5e9bbf61cd83c00c8b90940bca85d47a02164e80480a107a1ebf7b60001b62007dcb565b80600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002081905550505050565b62002f3c7fdb10f62e8da07d23c956a3ce17fc89806d429d2b212b90fcbe9c5241fafdf67960001b62007dcb565b62002f6a7f702641b0ac5eb6cfc0b282212a9a5d5a0b69c15ff52e3ec912c8bc066aef658460001b62007dcb565b60058160405162002f7c919062010594565b9081526020016040518091039020600060018201600062002f9e91906200cd1c565b6002820160009055505062002fd67fe226fab1292adf954a7556353498f89f470b960a169bacc49334be0a5c33174660001b62007dcb565b60068160405162002fe8919062010594565b9081526020016040518091039020600080820160006101000a81549067ffffffffffffffff02191690556000820160086101000a81549067ffffffffffffffff0219169055505050565b620030607f309ef9531f36185dff0010ad240f0b4bc8feaf74663f6c9977bdacea648afdd060001b62007dcb565b6200308e7f2914caac75c34be6795e751888733db9698ab0fc0a37b20234bd7e01cd29749d60001b62007dcb565b620030bc7f3c22dc82e5ea8515b58890a6fe4ec7bbd13405ff97ac5ef8a51065d89ac0635160001b62007dcb565b600081600001519050620030f37f065bfff9efcb6a0a0d4283dedc1569c157384b51ba5a42275f087166d194adff60001b62007dcb565b620031217f07a0877b470595a1bed7f4069258b1aae748ee2be40e2b449368b886f16b083c60001b62007dcb565b600082602001519050620031587f88de31fd45be0814fc625788152c43510d26f29f7f6d23b6e8ac628eb62034ed60001b62007dcb565b620031867f95e170eb493c17563ecaeb89ad6b8f068747a80986d5af1c5a16634687e59da060001b62007dcb565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633778febe846040518263ffffffff1660e01b8152600401620031e59190620105c6565b60e06040518083038186803b158015620031fe57600080fd5b505afa15801562003213573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200323991906200f0c2565b9050620032697f22b7f4607c3c44ae18e9e36d7b17c52f1d087a4974c042a68381b6e9cccd881e60001b62007dcb565b620032977fa661b4f61ab7ec4753954ddf0338c8fa96c8b4028d3dab2de1ad9f23c65543d860001b62007dcb565b42816080015167ffffffffffffffff1610156200333c57620032dc7fa26fb8bcf4f1771b2d7e5e63707d42330bb9214ab6033606eeca3114d809238760001b62007dcb565b6200330a7fc7f0202eb3b34b0f924ae85277882ec7f79a2d7b8356eef88daa6adc6bd2515760001b62007dcb565b6040517f26841a7e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6200336a7fd3905f6d07088ed9979f11a10da74291ffce44b443893defdd1c665877d0b6b560001b62007dcb565b620033987fe13d43abe6c90cd08f5456e8b2ca5fff2777219ccf19f80b16e39e7f64d94b3460001b62007dcb565b620033c67f92abe34b44d997f7c8261084d1ee1462fad86ef77f80e2ac8e3148673f9fe0ad60001b62007dcb565b60008267ffffffffffffffff1614156200346857620034087f9ffc4a4bb782793ddc737bfabf881edea5554217ba73d2a8531ad959e01b6dbf60001b62007dcb565b620034367fde8b9b760af09b3aa0b7a17627806f5196254eb5936592657d802308275b7ba560001b62007dcb565b6040517f26841a7e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620034967f667c7270e3db2c393e325079f1c42c98f5f8dc240afc97c2e228dd59ca15d1ca60001b62007dcb565b620034c47f431991219f3c3d47a7c62a7c600a38ef35480a50ad3f20c6b8067ba0206bb37a60001b62007dcb565b620034f27fe61f871a49ed694e9fc45fe012bd30cb7ac37251e6e8e5be461bfba02dd99cd960001b62007dcb565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d7866040518263ffffffff1660e01b815260040162003551919062010a34565b60006040518083038186803b1580156200356a57600080fd5b505afa1580156200357f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190620035aa91906200f15f565b9050620035da7fe44b4e1a428ae440304511e02428486862160379d2ac030bed49b2286e8dd18c60001b62007dcb565b620036087f37c8a6e0bf563b1a46265d01226189c22f6231fafb88e276f0f3f7dd49fe88e960001b62007dcb565b600081610100015167ffffffffffffffff161415620036af576200364f7fc3ff640996cf3ec2132cbea9875372a82f61351ca7ccd27ab462f8f7904784f260001b62007dcb565b6200367d7f82cb862f7865a5afaaa4b8e630163c232fda31627b9424afa5282602a92e147960001b62007dcb565b6040517f26841a7e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620036dd7ff3f3f2f560c7dd95784aefa3dcb9a0684b90fdd333f45f302ca3e520e1394fb360001b62007dcb565b6200370b7f763f0f24c8d6d0427860811393bfa52de3dc5552dc32ed1b787da3358f9b985e60001b62007dcb565b620037397f2171b27be3e70c20e42724b33e45e8cd952c29dc33a8c49ed787345fae8a854460001b62007dcb565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c5c81b2083608001516040518263ffffffff1660e01b81526004016200379b919062010741565b6101606040518083038186803b158015620037b557600080fd5b505afa158015620037ca573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620037f091906200f310565b9050620038207f96bf9f6abb6718472a89a02014655cedd02d9eea7ab5f2ed67dd4a43b68ce2f260001b62007dcb565b6200384e7fb904005eb0cf39c25ecf40554858ed6fcfbea47011394a7492416d1e8f05088d60001b62007dcb565b6000439050620038817f4e716e60f83e39dfe0e2c81ca10e7bfa538424639198f32cefb8d7beef5a35e960001b62007dcb565b620038af7f9bfac304b9cab5359b1ebf9f64acfb68bf6772737b9c62f9e9844536c24380ee60001b62007dcb565b808260c0015167ffffffffffffffff168460c00151620038d0919062010e62565b10156200396557620039057fc875b6911556b7889f3b56de25e365269d4e4156e598c06134a5bc829356175a60001b62007dcb565b620039337f9f76c73a085d5dcba5b121f015dc64ce51246c7797a28f59a85761e3117c2df060001b62007dcb565b6040517f26841a7e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620039937f6fa469a54555b4473f41910a6cffc15ee95264b922678645137f27f9f2d8d5b060001b62007dcb565b620039c17fc9c54d1d82e0b475ccb8cf34a8bc11ed8f0e971d6fe65c055b6bfb034422d8cd60001b62007dcb565b620039ef7fdf586645913f5534a1aabe70f303ec1648758df327a49bf116fbeef2699dd28960001b62007dcb565b6000620039fd87876200a657565b905062003a2d7fdf424496ae761981dbded3ba8c231100ed17d8e8870365f83424e37eb7aec92660001b62007dcb565b62003a5b7fca070eb9f348ec1d0adf6ec015415922bc180299ad9c31fa4234d7f8e42c156b60001b62007dcb565b600062003a6b858584866200a750565b905062003a9b7f52b4058550c148c2f0b79344148b1327225559b65055fe253be8e0f6a234270c60001b62007dcb565b62003ac97fa61cbde5d9adb043f84800f44e858aa7224910b3982a172ac44eab089d61ffe360001b62007dcb565b60008167ffffffffffffffff16141562003b6b5762003b0b7fc1959350b64ab9f61df5d4df4ae4cb3a353f43fefdf48d2b401e26e1a107945660001b62007dcb565b62003b397fdd7481b9f33e37f240f007826d2846725113a8ed4adf66630aeeb30c5dcb477b60001b62007dcb565b6040517f26841a7e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62003b997f7a7a0c93aea5a57b1db2bdb253e6229690517677b9f29dc8ddf6b112313565b460001b62007dcb565b62003bc77f7e8894f83ff0cfaff8da165fb61ecb962d9ecd4d6e1f0ac63caccee517f9c35760001b62007dcb565b62003bf57fce56d9eea9cdbe46d55ad476369fef4022f210ab7add2397645ba74e3de9c1aa60001b62007dcb565b62003c038587868462008a5c565b505050505050505050565b62003c3c7f2e99d5282cbe9b41396ae6919676609262c8d8d0ca75d6bb408c72208b2240be60001b62007dcb565b62003c6a7fd989359f747afac85feeaa7593de470cc0b82d783d7dc01f21859671cc52fc1560001b62007dcb565b62003c987fef12e34febd5dca298c9256a6863428bd118cb7cd8c75c6244dc38d4b830dcd360001b62007dcb565b600062003ca78685846200ae9a565b905062003cd77f80af9ff788fc10746f90a7644a9ff73e43de402ca7b6e11b427b2e57de3cb34660001b62007dcb565b62003d057f9369b24158b5773b6eb142032c753c6de77cb435a4de3c0f5230c082baa65ab560001b62007dcb565b8067ffffffffffffffff1686610140015167ffffffffffffffff16101562003dc35762003d557f2ad417701ed8b01f165513726224b5208aa68af68a38dbaf230ba9733c9a72e860001b62007dcb565b62003d837fabc5f60ca7d1dfaa706e98a572065d20c946f34a3d7d23ec7987ad4d099738bb60001b62007dcb565b60096040517fcd442d6500000000000000000000000000000000000000000000000000000000815260040162003dba919062010846565b60405180910390fd5b62003df17f8e83d575f8f37dfea9f380f944bbf77494a65a0bd507c7d584f2d23cd214f16360001b62007dcb565b62003e1f7f291c615dfae5b522de5d1c1cb4e9eaf3d96fbca69d843788da59ae79b7a3bd0860001b62007dcb565b62003e4d7f2d944243b9bddf92d4a48845310078ed2cf4254eb2f2983f0897f026efee7f3d60001b62007dcb565b808560600181815162003e61919062010ebf565b91509067ffffffffffffffff16908167ffffffffffffffff168152505062003eac7f6a7ac2c5149de31d04e87289ab173ae1a47dee174bde2919e8e2eae942b2e4af60001b62007dcb565b62003eda7fe7c522d4634f7e51363a25ead2c246adef2f41b3bf391b640af8ec708a03546d60001b62007dcb565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c819d86c866040518263ffffffff1660e01b815260040162003f37919062010977565b600060405180830381600087803b15801562003f5257600080fd5b505af115801562003f67573d6000803e3d6000fd5b5050505062003f997f478bf5bd5e7dbaa523ea69eda77cb8495a1a1ff7239a0a6f189a0973d933a10e60001b62007dcb565b62003fc77fe29dda0f9d87cc11f63a2ec441e6d8d6c1e0b0f8b61db9e2430ab823a605836560001b62007dcb565b80866101400181815162003fdc919062010fc0565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050620040267efad54934db310a1db713b9c93e7269423a2b3536bc746c64638e96fc485e8160001b62007dcb565b620040547fb5d74f82321cfce1104b93ed3e67d01ebeb87c9b51b3fa8b6e59ab535290c67d60001b62007dcb565b6000866101c0019015159081151581525050620040947feec7f9682bff4dff108b9daf13e5eda10aacb22d0ff25cca5e52dea7ab0ff7b060001b62007dcb565b620040c27fc9d8d5b17f9e5ede8b4a493b7913daefba83d8e4bdcbabcbfb5668f81baced7460001b62007dcb565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663681850d7876040518263ffffffff1660e01b81526004016200411f9190620108eb565b600060405180830381600087803b1580156200413a57600080fd5b505af11580156200414f573d6000803e3d6000fd5b50505050620041817fde96ca5b50fc244f9de7b52f733d5ed627e648134a1dbb22d178a8065e15627f60001b62007dcb565b620041af7f850af0896df406f3d6347ddb631be8fafc943539234a94a6b611581a84939dae60001b62007dcb565b6000620041df7faedc44903731f7aaf0677b20130653760cbedb6787d6fe0faef7ec327aee701a60001b62007dcb565b6200420d7f706f21ce7f27f5fcdd51ed62f082eaa7b539fd849b1872175561412cd5f328d260001b62007dcb565b60005b84518110156200437957620042487f9822250f810fb5b5f0181f5698003e94f1e4261dbaa9fc2a0502a1f830f29fed60001b62007dcb565b620042767f95918a839dcca798859f7ac840c513022c7882bbbacc0b7fd7865c73705d1aee60001b62007dcb565b848181518110620042b0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015160800151156200433457620042f07f3599049c8a9afd14d3412e4e4ec718b86a8d625c07193e09a05441b905e5627660001b62007dcb565b6200431e7f369dca1a76d0fb0ce2103d50a14645fa678c89e7576bcea9cff8d9bb670c0e9a60001b62007dcb565b81806200432b90620112e3565b92505062004363565b620043627fb1c6669e50a07b8604c911f0feaf86de1e6fe2344873e8ee4ba0c0bfcba3a11b60001b62007dcb565b5b8080620043709062011295565b91505062004210565b50620043a87f78e02b59f2ea02684b65bc60a92bc56df855f05f17097eda82f6639f9451198160001b62007dcb565b620043d67fcd0712d85369c9305432746f10c2c03de35b8a5389e9819d9f9abe532e8c091860001b62007dcb565b60018167ffffffffffffffff1614156200451157620044187f85c6d6306b8108696e412fb0c3e3ed7de5eb00ccf861586f7dece7fb4d6f231760001b62007dcb565b620044467f7949420d4b9cd606ca80380d4b8836b6700ebb3a031138a5a06c3c46aacb517e60001b62007dcb565b620044747f823225b88e1d8be44aa625f4591b82efaa5557452235edb8a649b5bf0947050b60001b62007dcb565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639cd103e988600060016040518463ffffffff1660e01b8152600401620044d7939291906201090f565b600060405180830381600087803b158015620044f257600080fd5b505af115801562004507573d6000803e3d6000fd5b5050505062004540565b6200453f7f8563f60887e0eb19b2e80cd69cda512409bb735cae8ec95db3fe3020b30e7e2c60001b62007dcb565b5b6200456e7fe086c67b268fc789036c861973d6db250ffdc70367540f3c9f0122dd0bb369fe60001b62007dcb565b6200459c7f39b14fbd207100409c141b8e627278b2fcafb3af1836acfdb54243552ed9549f60001b62007dcb565b6001876101200151620045b0919062010ebf565b67ffffffffffffffff168167ffffffffffffffff1614156200488257620045fa7f535f0afff06828b9a3bad9e681ecfc4ab48e44c7f9f395e2a2927a4223cd4f3260001b62007dcb565b620046287f88786ca25b0a6d286633981808c0c81bc6b1afee0e5838ffa4238dc3d18ea54a60001b62007dcb565b620046567fb37ba7f94497d003f1e72744d3d47539199e0fb1ce79080ceda18d61dbd29a7b60001b62007dcb565b600087610140015167ffffffffffffffff1611156200475a576200469d7fd0a6368437b90ac510f422c884b3ba20e52583613c0fee50cae8c888704f521a60001b62007dcb565b620046cb7f9016cb6926b44ffbcc59260b6a43dfbf2dbb9d69f3a713491d4e3b91195218ae60001b62007dcb565b620046f97f42b294934692bf93f57fed9519ec7df6e6cf62ce4c9a2d831fca55eab5025e6160001b62007dcb565b866020015173ffffffffffffffffffffffffffffffffffffffff166108fc88610140015167ffffffffffffffff169081150290604051600060405180830381858888f1935050505015801562004753573d6000803e3d6000fd5b5062004789565b620047887f0177dec289b8f2526afd362eec9db354328c162e05fa621b34dee9c4d3af567f60001b62007dcb565b5b620047b77f4dbdf830786377b5a61c71fe8f6dd1a173ffc7bf39e77f2252ad671ac7f2a75a60001b62007dcb565b620047e57f046253b7fae8085fd95fad7084326abf59e3567d74b4954bfd785533b453e65860001b62007dcb565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639cd103e988600160006040518463ffffffff1660e01b815260040162004848939291906201090f565b600060405180830381600087803b1580156200486357600080fd5b505af115801562004878573d6000803e3d6000fd5b50505050620049a8565b620048b07fda4f3826a42bc301260d211b53b113c04a11dbf8e2d44f6fb9f8178af17dfdac60001b62007dcb565b620048de7f39c0002d2ba1a4cd8083ea52428316a0043c7b24a453dabd0f033b6838cec0b560001b62007dcb565b6200490c7fa1dd33eda6387973352feea6a1bcd4086509ba2380bd627442dc858c1c65822960001b62007dcb565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166334fddaac886020015189600001516040518363ffffffff1660e01b815260040162004973929190620105e3565b600060405180830381600087803b1580156200498e57600080fd5b505af1158015620049a3573d6000803e3d6000fd5b505050505b620049d67fd906e0e68b5ec7fb4600d02287d875391838731415febbd92221a9251fffb32960001b62007dcb565b62004a047f116f6de949e83736e71fa0cd1b626c5f5a0c492b17b4d2df41efce8490639ff860001b62007dcb565b7f123b9998b2f027b02db4cb2eadef421fe9f1c21627569438588262d3a6baac6c6006438860a001518560405162004a409493929190620106a0565b60405180910390a150505050505050565b62004a7f7f0f37245d1094ee6542ef5556083a0020a077f034d6d780dc2b9313ce16cd5dfb60001b62007dcb565b62004aad7f50b277d23a8e03fd5f4c12764bd44351959bddb7ded3df9538b503d2c727e55760001b62007dcb565b62004adb7f5adc8d4f4d7a477875b7dfd6b610cb542621449e2d2077135d771333ff5d42b460001b62007dcb565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801562004b4657600080fd5b505afa15801562004b5b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062004b8191906200f310565b905062004bb17ffd8702c811b350a86045a376eeb00ad572c69a992d090b420e5329a0b2d9b98c60001b62007dcb565b62004bdf7f2f41e2cd2413c64350bf786bb4f8d541edb1c957b6e1df4b16455a4fd6b6be4460001b62007dcb565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663defce5d484600001516040518263ffffffff1660e01b815260040162004c4291906201067c565b60006040518083038186803b15801562004c5b57600080fd5b505afa15801562004c70573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019062004c9b91906200ef80565b905062004ccb7f3bccc43751f4498a0355e02df6206998ef019dad44c4521cee56097b54c2905460001b62007dcb565b62004cf97f40165c08f35e772a70c71c562f9cac1b70d03fc7adf6229f46a0a85388a184f660001b62007dcb565b806102a001511562004e8d5762004d337fdbc6002eca082629f6f8adb05be36d425ba8bc4e18f3b8a209a27bf4fd9d939f60001b62007dcb565b62004d617f991b1dc2ab63d9e73dbf297609a8d6ef5e7abdf35c02125e8b0334e1eeaab59a60001b62007dcb565b62004d8f7f2d1f6f56c2bf4f80910c7f9b91cde3179b8d242c1d863e535928e8c46bd0ce5660001b62007dcb565b806020015173ffffffffffffffffffffffffffffffffffffffff16836060015173ffffffffffffffffffffffffffffffffffffffff161462004e595762004df97f32b7539161197c774a2b90300b803df603ecf18e4d7205a0a1ab386b3f18c2de60001b62007dcb565b62004e277fd2e9076d3b1f29391c78961dd81873801916dafd0211407129f4bb398890795c60001b62007dcb565b6040517f65fd380a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62004e877fd4f87265a8e455774008a82b6f6e2f228da126f38c9af3cd62e718ca6e7b995a60001b62007dcb565b62005665565b62004ebb7f49516454ad1909cf8e2b080967efcb5e82ce25ed64ebc256d218cbdf361c057260001b62007dcb565b62004ee97fad130773bae93052fc71e17926cdc3786ef2bf3b2c3eec744b3ddb0bf571edcd60001b62007dcb565b62004f177f4f01fc12b60872eaecc1f0b3232a5405e43682b8dd7e03838fb0231b985bd15160001b62007dcb565b600062004f477fc678e8547e4dd67c92a03295c1bd6bb2bf81af3894d447d396ca3882bcdfadb760001b62007dcb565b62004f757f2e4b897e127bf439763d91a388841e7b864f12f52db11e498004b3f4a1a5eb8e60001b62007dcb565b6000826102200151905062004fad7f487d6fc02beddb9c6e7375d4f9838f3c1ffd4e3e5292a294364c2fbd0598b08660001b62007dcb565b62004fdb7f5b0add7cad7c89ec11e13bbfe5042280862b6a3faf46ac21f1e8c30a33b18d1760001b62007dcb565b60005b8151811015620051c357620050167fce0c85e23f32a73d8a607d82b18d35a4b0ef6dd773c9ba4fe7a9bb8ed9005fb860001b62007dcb565b620050447fb7694aea5332daec8d04ad02d30445113ef1ca4e0794e3376e8b7d22e1046b0060001b62007dcb565b856060015173ffffffffffffffffffffffffffffffffffffffff1682828151811062005099577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1614156200517f57620050ec7f30076260222f766f0e692dfec16b9782badb1a945e187fa107030cfa918655a260001b62007dcb565b620051197e9d3295b06ae142f00cb08da4b87ed4931dde71b0bd55241bf8c6119345a00d60001b62007dcb565b620051477fb6a550dc706d56c15d71c667b2a82157499eac2027f3697180b152e53b85c3b960001b62007dcb565b60019250620051797ff789299c91951d87c97785833c8de3a5df9332047f3afb77e8de7f7a666a2df060001b62007dcb565b620051c3565b620051ad7ff73ec4b5c97f8f40dc352934e3e8f9e383c2ad6e3c10ceecb8e9717367acb57360001b62007dcb565b8080620051ba9062011295565b91505062004fde565b50620051f27fda0dff376f24bba595edb6350018ba8308ce77ef33987d67471c54dcdd5a2d5960001b62007dcb565b620052207feb240abeaf690f25ea4e088fbe6efced9ac0d045a722fb957ec4daa1691b8e7060001b62007dcb565b816200550757620052547f2218e7a4a77fc02cd161f2ba253fd7a358f593ac611b48d078aa60c60d14e86a60001b62007dcb565b620052827fe4170e320d89e4dbd42edff5bfc4a07ea30e8dbdd1c4367878a0f52dacfa5d8e60001b62007dcb565b620052b07f9c991a3db94afae4abc330aa4a7340a9c6401544895e38bab9a58acae938506b60001b62007dcb565b60008361024001519050620052e87f3fd7c9bdf394090454d4d35e65ab3ba6a0d3bc31cb58477e392b906de6cd37b460001b62007dcb565b620053167fcb5f52505ff28f5d770a7520a810127b4d11a3fa6c6cf5dbe7fa3228b5be229260001b62007dcb565b60005b8151811015620054ff57620053517fcf5ee8ac67b2714408d0f6eb7da82f61ef005ec3c189842600baa5fd33e48a2e60001b62007dcb565b6200537f7f5691a547b7f59abed046b7338dae5ef5acf68f9310462cba6a0b1e8fca9658e960001b62007dcb565b866060015173ffffffffffffffffffffffffffffffffffffffff16828281518110620053d4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff161415620054bb57620054277f5e4c8049834abd2083c7ff7cdf03f688973c319cb5959b3980136ee504e1c9a760001b62007dcb565b620054557f614db3316a5bfa2e12c2b0bf29b848778ef03751d56daedef5eb09c1c69dfc5e60001b62007dcb565b620054837ff3f58db16b1c3f95f9385912b232b52db5d1e113d10a5e03581a6d9f434d254360001b62007dcb565b60019350620054b57fbf03e16d82fc02815197a62c8423507bb4e960b9826f67589401505d1c02736060001b62007dcb565b620054ff565b620054e97f716a642f926ccd67946dc7e71f597a8985898a5e3f3876dba2377aeee7daf52660001b62007dcb565b8080620054f69062011295565b91505062005319565b505062005536565b620055357f2f5af5888f79870c350780c94cb8e400792ab4cc31ddf32b8ec083a853e945ef60001b62007dcb565b5b620055647fad69cbc3edf995939c8f266aca6bead961b3fef314476c85214740733ad058d060001b62007dcb565b620055927f3637d853d640eec20c8eb32ecc33d1716dfa152228c7f43f5ad3293754d6059260001b62007dcb565b816200563457620055c67f48049f60921fe4073ad1fb4026b35d5e4de39d93be8b10c545c90a0d892ccb7960001b62007dcb565b620055f47fa12c1c03c908c2d92863891d6bd5f9f1d39f461229e63ec4b00e5f6ec737c0c360001b62007dcb565b60016040517fcd442d650000000000000000000000000000000000000000000000000000000081526004016200562b91906201075e565b60405180910390fd5b620056627f7bed2600832e78b78af2b012119d0f37be5ebb3dc249f02a1265aee5a31a091060001b62007dcb565b50505b620056937f15d9ad84822fddcad19082529c5f79872ab53d83f62c47747736e7597076fa2260001b62007dcb565b620056c17fe041bfe03372568f851a67eabbd54dcae7a537a801518b51a9b45d0c6a4e4aed60001b62007dcb565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a65374a85606001516040518263ffffffff1660e01b8152600401620057249190620105c6565b60e06040518083038186803b1580156200573d57600080fd5b505afa15801562005752573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200577891906200f0c2565b9050620057a87fb7904ba779d7dd9e8f94caf3a8a39912eef06172e6376253adedfa33ba58f94c60001b62007dcb565b620057d67f5e36c656103aff715e4ecd8010e2a0e4f30e3cea239491d776226430978d5eef60001b62007dcb565b6000620057e783600001516200b0af565b9050620058177fd7926d66c6d9a69d70b5f94880e41402198d0e19e7c2380a72dbcfc44829bfb360001b62007dcb565b620058457f73eba1007c4c6817bd9bb2d5bc4ef045ebdcc590e2f8cf2ce7d6a853b1db5d7960001b62007dcb565b60008560a0015167ffffffffffffffff161415801562005869575082610100015143105b1562005ac7576200589d7f08a8d1fdf9c8598e46d904e8abdb35a9012acd5e39a18c37618889f97d2b84aa60001b62007dcb565b620058cb7f70711224d311837b6089458f3d735442962e4628c11c6934b51f81462f7c440a60001b62007dcb565b620058f97faa48d71a50b3a8cf5ff291deced677c1b3b9687dbbc914d35e00cbb86b9d49c660001b62007dcb565b60005b815181101562005ac057620059347fceaf56f33254053bf278b675d1cdb7e7b3d58269dcfc916e4d9a49e62ac2bd8c60001b62007dcb565b620059627f90f9050d14f8dcc11d4ee3cb5f447148f594f2411efbaf44857620115f298b0260001b62007dcb565b856060015173ffffffffffffffffffffffffffffffffffffffff16828281518110620059b7577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff16141562005a7c5762005a0e7f94cbbce47d25e2c7c87588dc7f201e3225745f49b257c3809ac301ae24340c8660001b62007dcb565b62005a3c7f44f05e1b273905dbca6a070144e96243b5fd998fda83c774b9552da4e555751760001b62007dcb565b60026040517fcd442d6500000000000000000000000000000000000000000000000000000000815260040162005a7391906201077b565b60405180910390fd5b62005aaa7fe6bb7edfe12aa96440400836183eb37c83ba965fb35d95f1ff7c365e0a0b760e60001b62007dcb565b808062005ab79062011295565b915050620058fc565b5062005af6565b62005af57f68fc95c78e0cd8a47c6ccbeaa0ea7a6fc1029ad7b9587e13682bf19960af81d060001b62007dcb565b5b62005b247fd79ce2cd19d14061c65ec4617ddc245d9cbaa04836dba2257f25acf33ec96b9060001b62007dcb565b62005b527ffa3976faa4f233cd76a51fe5deb5f27f877c98aba8ac9722b834eaf0e487ed1560001b62007dcb565b600062005b6086856200b45f565b905062005b907f4ce49860b5e0360854214064de601d9b0428e28618beb5a7f3b975e33144776a60001b62007dcb565b62005bbe7f9baedcc545482aba9a35de05fb02cf25157c53ace06a6fdbae9240f9bbd34cfe60001b62007dcb565b8062005c605762005bf27f66deb016a2c7340a2b93294377bb4375d8f055d9c204177cf91cee389b61f32960001b62007dcb565b62005c207f2a26c5a0bf9d2b45d0aa4f534be4c04c0f27377969913c6c1e027bc02f8fabaa60001b62007dcb565b60036040517fcd442d6500000000000000000000000000000000000000000000000000000000815260040162005c57919062010798565b60405180910390fd5b62005c8e7f17a349f6c508bd849237d6ca6c6406ade0e014bfb2e4bb7d7d5839b50c234b4860001b62007dcb565b62005cbc7fcdc731413ef0d99abf34ba6f20e0a1b8a910773206c7dc5557bd17b468716cd460001b62007dcb565b62005cea7fbf8e04e1325a21512d5f274c9677a7ac6457f1d15a44e355a7420ca13585707060001b62007dcb565b600062005d1a7f31102aa8ccd8667514b8fd08c7636c0a1e602296f20bd0460c7cafe949184ac060001b62007dcb565b62005d487f4aa71b8fadb34ae7c55156a583bae72d4c7c20aaa1fd643567645d2c3695bae560001b62007dcb565b600062005d787fdb154f979c4973495fe8d10f94dc3be44542e551b234457a7de73c79343fe1e660001b62007dcb565b62005da67f076b5c095a55d80c489ebef76b914ea971c66582178637fe4793adf40ef62c1360001b62007dcb565b600062005dd67fc61d5295071b35e0002f8bb10afbcdeda61acbdc2c3ba3544be9b9788b5e57e760001b62007dcb565b62005e047f01c59d689cd4736da2ce845eac2bec1e84ccf0a0b6b714636c00f8470894978760001b62007dcb565b62005e0e6200cd3f565b62005e3c7f2aac02d3a10e51b456eb080c2e57e6a569ab7e451f3eb30a64ff028c82e19b8160001b62007dcb565b62005e6a7f1e337c5ab7476da9e627b4c9a65a7a2d733166bf7bcd89094c82855395909fad60001b62007dcb565b6000886101000151905062005ea27fb60cfb854ade38bfa3c6d151864b7e8d4c4c4a920bb011816fa1f14a13172fa760001b62007dcb565b62005ed07fea95d86ac872c4f42c55813b9986b748f6cfe30b725af4bcc3d26676eab08ff560001b62007dcb565b60005b8751811015620066175762005f0b7faf68160529ea7005d5ab6186a73efbd7a876521b23c1e9852bde0728313dbfc860001b62007dcb565b62005f397f31056821ae0ab0bea78d33444ea6f3f9230acfabf905d5eb56b366f7c83a063160001b62007dcb565b8b6060015173ffffffffffffffffffffffffffffffffffffffff1688828151811062005f8e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff161415620065d35762005fe57f90d7a0451ee7b3c0db2cc10dcb87a8055ea8e5a60cbf6245a7e8fd7ab5c5064360001b62007dcb565b620060137fd26c94a085f9341efe607f7a861d71e012f78eb068ee1aa365edf196a8dedf0660001b62007dcb565b620060417f466de43131e5884c78e632af312cc7f6133820eac22bc31ba9989514f4e898d260001b62007dcb565b82604001519350620060767fab1e7ac727887619d8a3dddb863dc6da34ae49dde914d2de9acbe597862e834660001b62007dcb565b620060a47ff958b70cb5124380aa154e8cde8ecc6081c99c5de7c75d3334c078d3e2843af660001b62007dcb565b8960e0015167ffffffffffffffff168467ffffffffffffffff161480620060ca57508143115b15620061d157620060fe7f9dc98a1902d59d0e17be140d5cae13f2a378b9cce070a5597f2b26a43bfaa59960001b62007dcb565b6200612c7fd958756a3f9a192ac71527ed38462281cee0509ede955ebbb6b505c69b7f477360001b62007dcb565b6200615a7fbeaadde72cea2776801bbfa43017f34e94c5c91c64c5d2c27b3a7a9aea3f688f60001b62007dcb565b6001836080019015159081151581525050620061997f51bc3cd067ba5c32b7c0e16495ff35860b3da1e6fcad4dcb8a6d961fe7421f3460001b62007dcb565b620061c77f871ba2862f5f0f6b4f318110d643505cc09631946bd23703fda9e49396e7be0660001b62007dcb565b6001945062006200565b620061ff7ff5479c4c07d735ee38f08b86bf5600a4acc5c33f308e4f4df39434bce32ae38560001b62007dcb565b5b6200622e7f2f0765d1a8599ad4c6b1339a5deba4740dfaae208c0417da9877650c77f4de3b60001b62007dcb565b6200625c7f1097c5bdb8378cea33bce60eab4e20c80e72a30cf3b48ef3b287d13d17220a0060001b62007dcb565b8960e0015167ffffffffffffffff168467ffffffffffffffff1611156200631957620062ab7f97e5d56dfefdf5d478a924d5063e97860ee08277dfda1c819725575dcceeea5e60001b62007dcb565b620062d97fa5229b4b1951fd2f97b96136b8dc82c9a23d400ae22fd280572ee5b22826250460001b62007dcb565b60046040517fcd442d65000000000000000000000000000000000000000000000000000000008152600401620063109190620107b5565b60405180910390fd5b620063477f991cda3e4a8a87eaf36d615f25ca1fae390498836879bb35928ea3749721564160001b62007dcb565b620063757ff24062659cff4a7f2249a0774f377856aa16a6e6cf3affa6c0ae45f6e80200e860001b62007dcb565b620063a37f1dd66f97cc736107512f4dc7c857ea715189193e8773cd04845884b4b200045060001b62007dcb565b6000620063b58b61010001516200c106565b9050620063e57f4af059b9b5584c4b47afd33abbb21ac42897c5337c111095ff194bca8c8d2c6560001b62007dcb565b620064137f66e1ad802da42798aaab0c16856cfad847c2af03a0a042fb55843d533ed397d160001b62007dcb565b8015620064b657620064487f5c7ebd9f13ed692d8bfa45e5b5fe2bb8b96f5e7a2de0c9964221fa772e79c25460001b62007dcb565b620064767fa02ce339ec7824036cd4f92a803f864e8d00b6012efcde8530f9b4dade32859660001b62007dcb565b60056040517fcd442d65000000000000000000000000000000000000000000000000000000008152600401620064ad9190620107d2565b60405180910390fd5b620064e47f4e4e55d54ff34e97976379a3b63ea0a37c516cfb611d77a44a4302d4520661e060001b62007dcb565b620065127fc4606c0d1a325221799bd3893b6b65409de49f58959fc53f74138c6c107ba78a60001b62007dcb565b8360400180518091906200652690620112e3565b67ffffffffffffffff1667ffffffffffffffff16815250506200656c7f2d335617b01d33ca1c43261649dafea7bdf7d55a6e59c281c92c28b369fe2c1c60001b62007dcb565b6200659a7ffb85d4ed3488c5279346824de02dc6979f3f3ba5b4fb93ffb2ddb0e89a2efb7960001b62007dcb565b60019650620065cc7fbc27faa4ad6d6257a265a86e66ea9e498d06d40ddc64abdb4cef1701a82e650560001b62007dcb565b5062006617565b620066017fe6438c61e9fc16dee939bc612f7744982fbbec328c27e65f276997976a37f95a60001b62007dcb565b80806200660e9062011295565b91505062005ed3565b50620066467f654b76213247bef1dd5b384d007295b3faa419cf820b13e2924736ff9863cca560001b62007dcb565b620066747f31abf154acfa9c890b7b44f2730b36b87c8ff5a02f3888b08c0b3b922c4bab0960001b62007dcb565b84620070c857620066a87f57d9467eb25d53b550c70f02612c5f8f22babf3b9e58a5803dd49e46941f799d60001b62007dcb565b620066d67f36e7b707deabb58bebbdc28821f6d0f50b6d16e0bf895d7bc25f54c2e62ed39260001b62007dcb565b620067047f5ecf10f430a069cada01a8e4efafb2f19da916e44333eda7bb14d5e66882a16860001b62007dcb565b600189610120015162006718919062010ebf565b67ffffffffffffffff1687511415620067c757620067597f8f6e4b45dc1d470189e41089d3dac7061847ddc946b44bdee6f73b3bd8acf28060001b62007dcb565b620067877f908cd1a75f898875e07becb067dc3d87872daca307b743fa0f04657e72e1793d60001b62007dcb565b60066040517fcd442d65000000000000000000000000000000000000000000000000000000008152600401620067be9190620107ef565b60405180910390fd5b620067f57f4a6395058dfbe966abdbc562d905bb356e766f1fcc45699a6365d3a486757f9760001b62007dcb565b620068237fc449cc6fd29fa6b4ec8d7c52bc1e0385786bea67dd8f140ac2b7bf57dba78d4b60001b62007dcb565b620068517f344fcf6ec12f1a252194c40d6434cd9bff7e5b78a745e683f70f4aa871bbce5760001b62007dcb565b8860a00151896080015162006867919062010f3c565b67ffffffffffffffff16886060015167ffffffffffffffff1610156200692357620068b57f35b89a7804e6248f89e4c0e90a60fa428766ebb0191f169770da213949ab36b460001b62007dcb565b620068e37fc2c7ab60b347c201481cfdba2ed6f8346fa7ebda66dcfdfe79c90ee226180adb60001b62007dcb565b60076040517fcd442d650000000000000000000000000000000000000000000000000000000081526004016200691a91906201080c565b60405180910390fd5b620069517f0e1ed403a7d2df29aa94c56b17cf5574da135107b90a5e03d367ae712d1da91d60001b62007dcb565b6200697f7fc2e03d82f409d0de8d0cf12fc4730ae724369d8086da8b82a104268105be51b160001b62007dcb565b620069ad7f842686004b2a4efd2e924a9be272614647bb5ea7f90f530857899616c5aa60c260001b62007dcb565b8860a001518960800151620069c3919062010f3c565b88606001818151620069d6919062010fc0565b91509067ffffffffffffffff16908167ffffffffffffffff168152505062006a217fc769f587e08ef033b47dd14b102ea40178e2023b0d224f6bd410714f2f0e798b60001b62007dcb565b62006a4f7f7fed0300b22ee14f3a154f136e7a0967bb643f2a3fdb006664404bb7ee94352960001b62007dcb565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c819d86c896040518263ffffffff1660e01b815260040162006aac919062010977565b600060405180830381600087803b15801562006ac757600080fd5b505af115801562006adc573d6000803e3d6000fd5b5050505062006b0e7f580d44ae7f38d805036aebe0ae726da384034cf1f31e5984f98c68150393d9e060001b62007dcb565b62006b3c7f1641ef084252f5710d5c1c2486a78ba1eab7ff91523cb8a32a811c1d63f717ea60001b62007dcb565b8760c00151826000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505062006ba67fbd2530872a955b63ac3b5571480e163bcfbe93e7cd2953fbd947b2f4b205486c60001b62007dcb565b62006bd47fa9b61643be691616d262115cb43e28206088a87af72dfff0ccbd650f266d133960001b62007dcb565b8a60600151826020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505062006c3e7ff11470d72c67bfef7ee351fbc13dbb959bcd67cfb073e21c8913abeaf7ec7f8260001b62007dcb565b62006c6c7f493b484f6f70d2c6fffbf34a90e8045e5d7e0144950f28380bfb975c0029aeff60001b62007dcb565b6001826040019067ffffffffffffffff16908167ffffffffffffffff168152505062006cbb7fdae5e583cf336d5fa2237f5431bf8b8ef25c65005d5e41c7582f3e9ae30e923760001b62007dcb565b62006ce97f0833be629a04c0019c36ccdfe8d1ef8dc5858520f674082d214f9ed2fefd3e5f60001b62007dcb565b4382606001818152505062006d217fd73abc0253b8ef6f032994447af98bc4a5c86e4b588c2fea9a8ff0c7edf4f12b60001b62007dcb565b62006d4f7f5bec59ab1c32e516cc7e03e081b32b7ce77ccfe6a4e29e4a90c1e057d3ce046e60001b62007dcb565b600082608001901515908115158152505062006d8e7f887a6074bf649d35fc26a2d95772c9e5dc04d402bf7f5881739c5cd878ed192960001b62007dcb565b62006dbc7f7f83ee04645312cf87972f31aec4d93a3a66b2e28b60f5e7576ebcabc48ad5d660001b62007dcb565b60006001885162006dce919062010e62565b67ffffffffffffffff81111562006e0e577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801562006e4b57816020015b62006e376200cd3f565b81526020019060019003908162006e2d5790505b50905062006e7c7fcddd69650b0b0f84c96262441ce7e3062c7a59cd4ee4cad02f9b8d42e3e74a5d60001b62007dcb565b62006eaa7f2d313e4e7b939ddab0d139a4153669e99a187ea4a88b5a049faf69714bf2ea6b60001b62007dcb565b60005b885181101562006fb05762006ee57fe3f1b6837b84f034321e86b6a2561ddc332227a505de6d025b7d08562f6e93c960001b62007dcb565b62006f137f69c5886ae649ec666ae670ad3695eba88beeb2e7c8be7150b00318d5dedda96560001b62007dcb565b88818151811062006f4d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015182828151811062006f8f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181905250808062006fa79062011295565b91505062006ead565b5062006fdf7ffc69c08df168fe45cc8bfe55b036f685ac76948277b96d2ee634035a9a56976460001b62007dcb565b6200700d7f9bc5fc0d41cbbe637f5db991d2f9d463643190e7ea26e2e2057cb78b77e6576460001b62007dcb565b8281600183516200701f919062010f85565b8151811062007057577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181905250620070907f92607830c0537e5917bd71c4c9d319a6bcb974ac6f7d0f6e4e043a32b0b0a10860001b62007dcb565b620070be7f96d0410b004b6d287622cfc6498b48457e9abf320f2ec54cb945bbe9b3a5a38460001b62007dcb565b80975050620070f7565b620070f67f2206c0c67469616760a2c6b79525c7bc9f3eb7649404b6cbb711fa733dac8f1560001b62007dcb565b5b620071257f91fb238c6c32bcb0ca62b7038d4aa9506dfbff4e800a5a80a8a06be69247956560001b62007dcb565b620071537f685a3098d28d5e6aea1aad476c3b918395404a48fbe6a2cbfb70dbac4ad65d8560001b62007dcb565b6200716389600001518862002bc8565b620071917f22df9e600b94769026c7c26fa6efdfa71c08457e972e9dce8c9904c8d74d2f9f60001b62007dcb565b620071bf7ff4f388e2e5ca21fb41f1cf62966b93542fac989ead09b54f0fb75e0f00e6740560001b62007dcb565b84620078a357620071f37f2168f4ad4ec31a4a24c4ee020f43222cee5da4e81fe52efb8f5f505053e99bb460001b62007dcb565b620072217f8df3530a42ab5391f144ab40f00833dda660b17504ed139c575ce968619164e860001b62007dcb565b6200724f7f0b55ede9a25e214eff982b82950e8794da059d861f525b4146d36c0a025581b560001b62007dcb565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d760405180604001604052808c60c0015173ffffffffffffffffffffffffffffffffffffffff1681526020018f60a0015167ffffffffffffffff168152506040518263ffffffff1660e01b8152600401620072ea919062010a34565b60006040518083038186803b1580156200730357600080fd5b505afa15801562007318573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906200734391906200f15f565b9050620073737f05eff38d788bbfc4f744769421c28568bf9019ac5cd77b929a40767234ac0bad60001b62007dcb565b620073a17fc71aea22379865bc3e4d8f805cc058f9b5b2280081547230dbd57391f634bfe260001b62007dcb565b896102a0015115158161014001511515146200745357620073e57f3ddb6da24c09a6ec0cd5bea169f3020e7a024a1522607c2a595208a71d914e6d60001b62007dcb565b620074137f3185ab22d7a756c409c1fe36dba1302ae619a41889aa45fe1056aa4dd532dbf660001b62007dcb565b60086040517fcd442d650000000000000000000000000000000000000000000000000000000081526004016200744a919062010829565b60405180910390fd5b620074817ffc0e9e2cecba13d9501baa9e074f63710e36de5d315099a1326ff80b329e39ab60001b62007dcb565b620074af7f50b9fc14544e4153a7ac873fd958cf5d8add32ad7b334f2ec086e13da384e94660001b62007dcb565b620074dd7fbf200f643997f0c31a632b98b0ec1c9288755ecb59349ad74ac5588b850d8aa460001b62007dcb565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663955f98b7828c6040518363ffffffff1660e01b81526004016200753c929190620109f9565b600060405180830381600087803b1580156200755757600080fd5b505af11580156200756c573d6000803e3d6000fd5b505050506200759e7f40ff5548ab414950c9640130aa0e5b8d4ecc24b2d07e2c72ba25391d06096e0b60001b62007dcb565b620075cc7ff283d518ce9e8031d47e7655dbe4863c200f9538974a785d2a14b0cfa652462360001b62007dcb565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dcf74960826040518263ffffffff1660e01b8152600401620076299190620109d5565b600060405180830381600087803b1580156200764457600080fd5b505af115801562007659573d6000803e3d6000fd5b505050506200768b7fa1e1d65e6636c12311cc600108ac74693390ff9aac778798bfc354d4309102c860001b62007dcb565b620076b97ff3b77d4b43e3d2c3ad71273cd3c9a24a39b93b2897f298caf5e9592a4ff0aa2e60001b62007dcb565b60008160c0015114156200778057620076f57f944fffa1ebf48a4c5bb803bc658702a572d56e4dc6cfc39fa35affaba096a8ff60001b62007dcb565b620077237fb737fb5d11ff70fdf1ad2da8a362c1851f517e46ac0064e8153b2529fd8f5ca060001b62007dcb565b620077517f179670ddfbee6ba3325a8d6ac4d1aa6d4e3f56606172543bce964e4477333f4260001b62007dcb565b8960c0015167ffffffffffffffff168c6040015162007771919062010e62565b8160c0018181525050620077af565b620077ae7f44d259532cf45dc9ccbe0cceb2e2653206891b795e316f6db6615c125580612160001b62007dcb565b5b620077dd7f40464701613f5ca06f802831aeba536deaef91e53b98d135a26b20fc71e0a48960001b62007dcb565b6200780b7f42a8b0a4c0fbb45241871705740a2a702f033b564189e16915339e6a399f6b6060001b62007dcb565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632384a6aa826040518263ffffffff1660e01b8152600401620078689190620109d5565b600060405180830381600087803b1580156200788357600080fd5b505af115801562007898573d6000803e3d6000fd5b5050505050620078d2565b620078d17fd275554609bb8ce56f34a6d79e653e9dc6784fccf08ed5d7c77766f738dc83f460001b62007dcb565b5b620079007fc3601494002ee296670f7ae30d6e5338fd25ab349df274a63719b71a7918fa8860001b62007dcb565b6200792e7ff614dcb71a494be5360271ee9ebe88ba2e28dfbc5538781e07f0894bf79bf39c60001b62007dcb565b831562007ceb57620079637f1103d7a072057aa8f8004cc6739e544342e5170f8d175bd7a75f153840876c0260001b62007dcb565b620079917f8587c199af6ea0dcb78521edd933b0c373c4daf2ad8740a405b095c609c1a6b260001b62007dcb565b620079bf7fa269072e18852bc6aba264dacd7a0d8cd661949e390ccf77562ce2610948161f60001b62007dcb565b60008b60a0015167ffffffffffffffff161462007c4b5762007a047f5bab5d4d9591d77fe281fab59ff5465c36d8ba8359494ca185cad34fc778ba8060001b62007dcb565b62007a327f181536a9806f586125e649c2dce46bd99f2db65f32238582f48133520879bfe860001b62007dcb565b62007a607f57db23d9080db44dc07efe2216df060dec3879fead9f30918badcbf1be33066a60001b62007dcb565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d760405180604001604052808c60c0015173ffffffffffffffffffffffffffffffffffffffff1681526020018f60a0015167ffffffffffffffff168152506040518263ffffffff1660e01b815260040162007afb919062010a34565b60006040518083038186803b15801562007b1457600080fd5b505afa15801562007b29573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019062007b5491906200f15f565b905062007b847f589b8d6d13b6aaf5cd9cb2add9ba35149f2ce0addea2ba7cf6aaeed13e4f38e560001b62007dcb565b62007bb27f98b0b781bba2aa9d3b71f9788b667eb14e43d5c97ed7cfe459aa736c1f5d5e6760001b62007dcb565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166247c003828c6040518363ffffffff1660e01b815260040162007c10929190620109f9565b600060405180830381600087803b15801562007c2b57600080fd5b505af115801562007c40573d6000803e3d6000fd5b505050505062007c7a565b62007c797f7ab413b54fe1641cce7c27fcd40db3fdd5b77d0ad21460ffbb64037bb214094d60001b62007dcb565b5b62007ca87fd28446571037ce54ef250496f56e9e658b8c7c7a6a12976d8ec18a0f057477ff60001b62007dcb565b62007cd67f043babd9bb9e0e9c76fa9f0c68cd587caad0e589b31503cd963ad0f16d26900d60001b62007dcb565b62007ce58989848a8e62003c0e565b62007d1a565b62007d197fa2bce3ea1622cbebfa50eac412afc04f9407d906fc0fad49dc7b5f545461f98f60001b62007dcb565b5b62007d487f0881b14f4cbac6a640f56904577740546c91fdd1fe5af3ade0caf3dcf1694f5860001b62007dcb565b62007d767fc73ef999d2b7020dc4f80340c44c5cce0afcb65bdc6b735017a66ffedf1e995660001b62007dcb565b7fd936063ae3cbbf2dbcd2adc837a997ab97adb09f24c566aa79d101e88928dea66007438b600001518b60a0015160405162007db69493929190620106ed565b60405180910390a15050505050505050505050565b50565b606062007dfe7f1afa166f9206d7acadf80b68c76817c6dfecb49889b0acb41f3af875cb4b998860001b62007dcb565b62007e2c7fee0e95267b4bdc677e4d2c30e3e537c8e689ce71ed410b933ab45ae6b711494760001b62007dcb565b62007e5a7f8d10afe56b2353d3ac50c9d3c8b29015d39720599e20a3b9ce88f4b05211add860001b62007dcb565b600060058360405162007e6e919062010594565b9081526020016040518091039020905062007eac7f52cb81c196e79141f8a2b997ca09889be0f2dba0c56a64acdab301fe23f09e0360001b62007dcb565b62007eda7fa412963f5783fa46f92181bbabd06f13eb15132070bf72fce0bb67f1c6952c2560001b62007dcb565b6000816002015467ffffffffffffffff81111562007f21577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801562007f5e57816020015b62007f4a6200cd3f565b81526020019060019003908162007f405790505b50905062007f8f7fda5fbb2deab399ff1e7b5122b29bee87ab0bbb19cd577f2851b7c52f9e86f32660001b62007dcb565b62007fbd7faed76bd3e02d3e22125b5cee0cd2191d12a0e7fcbe6a77917dada412b5d9a8f260001b62007dcb565b600082600201541415620080605762007ff97f1b9502f8d88523498966d13f0259b917ceda06de15b97ab776dd38907c24f95e60001b62007dcb565b620080277fad945ca0cbd2170b6491a8366ccc72e40ca53836cd11a83a961529f0afdfdec360001b62007dcb565b620080557f920c38f36b3f5f66a72a427d85e95ad565a4aa37994caf321ca876830b8d0ec260001b62007dcb565b8092505050620082ae565b6200808e7f9f7fb924ec4e72fe7a63d885dba9bf26c0f5c4c0e3e9516a438a654e68a56e4f60001b62007dcb565b620080bc7f58d90a4a575b4f93c784394f5ce5e9e4c3a7f8af46ae87daa9ec7f93f284e23260001b62007dcb565b620080ea7f23fd89998caa391ffbd0d023973356e4f6441648835a2c52e11baf42fb0e646960001b62007dcb565b6000620080f7836200c2c3565b90505b6200810f81846200c3d390919063ffffffff16565b156200824b57620081437fe1818bb37b1b0aef35ce99c07992a868f30411ee6a552a27b2baf24097b996ff60001b62007dcb565b620081717fcf68ae11dd5355f2eb71939e3c9ea309194896ab10c93fcc31fe3a98c022433e60001b62007dcb565b60006200818882856200c47190919063ffffffff16565b915050620081b97f533833147fbceab0a04b24fcf808af5e521d9688e180cb7f8c19b3a5852bd4b560001b62007dcb565b620081e77ff28fd36563106efdc6d9fe2b1199afe1670dd325140464fa68fe4fe1788843df60001b62007dcb565b8083838151811062008222577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181905250506200824381846200c72b90919063ffffffff16565b9050620080fa565b506200827a7fb7c915a3ccfffb3e0f9913f0199c0f56e7b6ef28b0108788ad38f218e9f350c960001b62007dcb565b620082a87f79dab4d53fc5137479b96022fbc3553c4f217b1b54d4a03ec3e1a708ca43e5e660001b62007dcb565b80925050505b919050565b620082bd6200cda6565b620082eb7f9f10345e5f43458e57a947ca6b83559efa265af8f4fc590e330a25afdc57a26d60001b62007dcb565b620083197f48df4fc357928b767793e678821caa8d1bff6d60a6cfd8356e6f1c43eae73dfe60001b62007dcb565b620083477fd196f59e1f7debf759858b6133745fbe8b0ef083d3732ae70816c38d8fdb9a8560001b62007dcb565b600083836040516020016200835e92919062010564565b60405160208183030381529060405290506200839d7f57690fbf70fe6a4270c39626a9d12abba54408c16c5a28f5850c83e3f4c20c7b60001b62007dcb565b620083cb7fab2a24b3553aa2af51df49b80913d17a13de2367cf1292f6c69fd1a73e3d1c9960001b62007dcb565b600881604051620083dd9190620105ad565b90815260200160405180910390206040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505091505092915050565b620084bf7f84dd572c893bc4666be08db74e72b794cf1cf3766b3f8787e93c0c5f0879bc0b60001b62007dcb565b620084ed7f856149dbfbdc9ade8bfbbf465ca1de45f6a684bdab79b1ba141d1f15e915fa3b60001b62007dcb565b6200851b7ffba6b8807d935dbfd35ea6e619646dd01b9badd36d38a7ff1347cf28b1bf4c2d60001b62007dcb565b806006836040516200852e919062010594565b908152602001604051809103902060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050505050565b600060019054906101000a900460ff16620085cb5760008054906101000a900460ff1615620085d6565b620085d56200c910565b5b62008618576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200860f90620108c9565b60405180910390fd5b60008060019054906101000a900460ff16159050801562008669576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b620086977fbda71fc2da12636ef1b8e0c563f9c770f260c1061c6254f081da9d722603ce8e60001b62007dcb565b620086c57f0c1ab0ef414358ff4ed868e7e53bf6fda8b9f4e829ff4e5274b7e544c4a8f09d60001b62007dcb565b620086f37f118e289b04c14820f63a5046af7efbe20e52fdf6639f0a1555ed548e487983a660001b62007dcb565b86600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620087627f8b70a8d18fe5262854ecc8f67f9d9f6b4904fc4e788a9c0f1435c167a141402e60001b62007dcb565b620087907fc9176ba6248ab3320be3a221cb8ff1badae7755cc04c45e1fea5879387fa4b0c60001b62007dcb565b85600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620087ff7fa42db23c8e839c0f8bb3fe58e51e846ec182674fbadf7392aa79f0b9ffda288d60001b62007dcb565b6200882d7fcba2a0cc461453a5a7dcfee5c9dcd715d4f83d37ac41e41c4d62c8c7a00a9a1460001b62007dcb565b84600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506200889c7f14af8f1265bd3d38aadbe6e425ddd83e7b561818296a0cce7fc06c33c0ff7ccf60001b62007dcb565b620088ca7f792ce25d7bf81fff1ba38410a538981d6334604ba0e7cc3c202dde46c45f185b60001b62007dcb565b83600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620089397f541b62456456bfe57e9d2e62c294a4b1e8f71207f40471e92286403d39466a6760001b62007dcb565b620089677fe091ec7180f6261f0ad63793a76152fd3674588765497234488b7ebbe52ad43760001b62007dcb565b82600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620089d67f6a5e2e78cfc77af8e92a700bcab862c87a0179ac0fc1fd7979c551526c8376ee60001b62007dcb565b62008a047f6807fac6fd60913fee8164afedb4b765b64bd5e8e61dd658cb62ef2a74c49ab860001b62007dcb565b8160000151600460146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550801562008a535760008060016101000a81548160ff0219169083151502179055505b50505050505050565b62008a8a7f536794cbd9b840eb0a0422b8b9f269becb308dc1d96f806c95d1caf314dba8af60001b62007dcb565b62008ab87fb7e66e972fc516e98db148ac6b928dfa155a046c596910f4ccebd7b1e9034b6460001b62007dcb565b62008ae67f09d0637848afd1783ac77a56716cf0d8d54d334e6623790ddd398e17d212790360001b62007dcb565b600062008af483866200c923565b8262008b01919062010f3c565b905062008b317fbe44f4e8ee306bb4efc4b36e17b6ea80aa59bc72308089baf49a8d575007944b60001b62007dcb565b62008b5f7f464afa7cc4a013c18352eac9c6ee3776276ead8d95c7070cf47c91b0ff53636d60001b62007dcb565b8067ffffffffffffffff16846000015167ffffffffffffffff161062008c405762008bad7f505e1a8e4d383caa6cbe518fb40f3aaa80abf4a5e6fa027cdacdd0312e247dc960001b62007dcb565b62008bdb7f43d49191d7a47b166ac4dfcabb8850de88eebc0baa14ec066293818a985bf47760001b62007dcb565b62008c097f5d4026ae0d46f72e9b0d9ea63064be789031a8a4fb19a2e33a3961e39e10787a60001b62007dcb565b808460000181815162008c1d919062010fc0565b91509067ffffffffffffffff16908167ffffffffffffffff168152505062008d4f565b62008c6e7fbd4f89b98aea14b10997bfe2b2489e4a110dc2a339161222e454f9b60f9e042460001b62007dcb565b62008c9c7f95022b7fc316ef2c8520d78abbf1367f539b72bb19f83b60b711f79c3a96c7e960001b62007dcb565b62008cca7fada2b0058b797e953714e1501f65a7084987774e4c2c5ea20999a752017c3aa260001b62007dcb565b6000846000019067ffffffffffffffff16908167ffffffffffffffff168152505062008d197f2ce826fd2805531f69a6f088261036d17b686419a474b4474a9fb741d1b6f88d60001b62007dcb565b62008d477f3a2a0611813370ab03dd50c6d884d5c10c0d34fcb44132043eb846e20b88aa7160001b62007dcb565b836000015190505b62008d7d7fe1942f2043a44bee3a7b91e647909cf54f99a7303b380ff65250c2ec209c106e60001b62007dcb565b62008dab7f2da9621272571da9356bf8c908f919ee9ea5d8497e4dab18dbc910bb62e713a360001b62007dcb565b60008167ffffffffffffffff16111562008fe85762008ded7fa6520bf8905555f96f501407c95e497772cc9b950fdad2d1629804f22d26683b60001b62007dcb565b62008e1b7f7bf28374ed8a80f7bf6e499592c5096014e36c4a5b5039f3be1b6986164e42f460001b62007dcb565b62008e497f58ef7815c29ece677717aeae48e3d185f910ebf43b34076b54074b904a047ca460001b62007dcb565b62008e777f8ef6f8f1a61283a5dc772d4207fe2a7dca9244113b7793d4b8ee93978412a7b960001b62007dcb565b8067ffffffffffffffff1634101562008ec7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162008ebe90620108a7565b60405180910390fd5b62008ef57fa1d50950a2fc0b61db43fbbe13fd3d5bc24a5e7bdcf652b92fd7ac871c746cd860001b62007dcb565b62008f237fc986e9ceee21fcae08b852797474834920616b5b5a34d75daf4ff2d996ef199760001b62007dcb565b62008f517f3d648d98081a28c662806c1d60075156710ae9c776c78095692b00876546dbef60001b62007dcb565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c819d86c856040518263ffffffff1660e01b815260040162008fae919062010977565b600060405180830381600087803b15801562008fc957600080fd5b505af115801562008fde573d6000803e3d6000fd5b5050505062009017565b620090167ff3405b84f112f278f637612ef29278744afa095b1c442160da473e152a71547e60001b62007dcb565b5b620090457fd5053904a73dbeef0b4942f867723551bfb37e4d0c7df651bef2c4732653330c60001b62007dcb565b620090737fbdca501af8a0a6ad9d4028c7bca3d9eacc83fc2cffb2262df5b8d8f0ed25194b60001b62007dcb565b62009088856000015186602001514362002e16565b5050505050565b6200912d838383604051602401620090aa9392919062010863565b6040516020818303038152906040527f969cdd03000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506200ca99565b505050565b6000620091627f684ea5201810e9d25600327497310fb8f683536f06e88823e879d23cf06599fe60001b62007dcb565b620091907f18f66947524524b68607518bd00d1a8234cd736614d7568f21c80f2b0f7e1df560001b62007dcb565b620091be7f10da145549d112cd2e2043ea10552f530e6e3e9319350adf88dd17bb55e35e8060001b62007dcb565b60005b84610100015167ffffffffffffffff1681101562009f6757620092077faa7776014ad7ead320e53e7eb756c0466a3dfd96ec72c68fd35fec5669183f9e60001b62007dcb565b620092357f1611fa7507d3c7a7b97f8c114810f9d64e7b6e3d8b60f7c4292fbf658eb708e860001b62007dcb565b6000856101600151828151811062009276577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101519050620092ae7fddfec0236af114c110286b2ae9d7064158171bf2d36b258d274734cf8192076860001b62007dcb565b620092dc7f743ff5aec4f1024b2952d331bf3e4587b52da15f256d749289ea419bdea9f7df60001b62007dcb565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663defce5d4836040518263ffffffff1660e01b81526004016200933b91906201067c565b60006040518083038186803b1580156200935457600080fd5b505afa15801562009369573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906200939491906200ef80565b9050620093c47fe34a54510219ff588df2b407b38455d180f34be2fa4fc74da2633ddc38da774e60001b62007dcb565b620093f27f02fb1b83b5523d6f551ce66ad701a62d08c8eb2cfe4265c57167b45974a6ad5f60001b62007dcb565b600062009403826000015162007dce565b9050620094337feb50cea8771d265890b27830442fc47a426c3e6e9e63362065253c6dadede08860001b62007dcb565b620094617fceb09aba30a3199b72cfa41c182edbcb7ee31049ba446f863c2752f1e8835cda60001b62007dcb565b6000620094917f16facca15f2f6c51ac021a3361592fdf63ae781007a0fa0f951c12d544150f2860001b62007dcb565b620094bf7ffe1b9038f684ec62a3c18fb32270a38ce288de2e9ed5d0e32af0f39348ea1aef60001b62007dcb565b60008361010001519050620094f77fe31d54d43646917ba335c8eb03466ce6fba9e410085aacc0d99a4816814b6c2160001b62007dcb565b620095257fa3b899637dc2d6e76152961bbd741ead7586a1ee0e61368ac71146ad53ba271660001b62007dcb565b6000620095557f99b41c0504365ec000ec5664522d04aab39f93c70b97ac79df898c48b488d81560001b62007dcb565b620095837f889aad19aa8b2f12547bd53701e9300e71ee7bc4427fad8e29a6775a3366552b60001b62007dcb565b6200958d6200cd3f565b620095bb7f98acb880014096c45dd184cb1805ba6f0185415f45b8cd0fd6480c38ef09336f60001b62007dcb565b620095e97fc8aae2e013c780d288b51dd0b66c48f14940cb5640151bee18e5495beb8b662860001b62007dcb565b60005b855181101562009aee57620096247f9fdc6bd6749635a1bd7b994977e872a7a52b31e50e19468847eaad914b76040260001b62007dcb565b620096527fae3cec05bd1fab36d190e64821b4b4667235c73277434d68cdd5d59cebaec8e660001b62007dcb565b8c6000015173ffffffffffffffffffffffffffffffffffffffff16868a81518110620096a7577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff16141562009aaa57620096fe7f0ea378c22593b075a1553e77c583adf098113802e1e7c0eccaea6167d17d5d9160001b62007dcb565b6200972c7f70c4fdbb6e4d725c55cf48302213ea152a8249f9022285d2a460c6532dedcd0f60001b62007dcb565b6200975a7f417622f68f162bcaf709d5670d9bd17491ece2d5bd8116aa4992094e494212c160001b62007dcb565b600192506200978c7f02643bddce67eb1cc54d7af9cde5aba0363f56b003c9d642e2a4f9d1d0fa862460001b62007dcb565b620097ba7f984d28118f6c760be0f94ce2a7c45279e1790c546357978b904d56253386363460001b62007dcb565b858981518110620097f4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015191506200982c7ffc72fb01c32146750a854c3646e45eea74dd146a3bd7937a897deee2420dca8e60001b62007dcb565b6200985a7f27b6c5c5e6e1f53b1ebf68ff816a859d1830f9e54a90d3bf41a6462611a98ee860001b62007dcb565b600082604001519050620098917fb58600b74bdf43479be771e686231ea9a75e2337f75a1131f8a75e349415ec7860001b62007dcb565b620098bf7f8c9af00630ac40f2d84f7c5c732b9a73a6dece5b8b165a160594486978ea12c960001b62007dcb565b8760e0015167ffffffffffffffff168167ffffffffffffffff161480620098e557508443115b15620099ec57620099197fed80116c3927465935ae23138ecb80f4bc4e0bbbb4d2a94a99e262409a76e57960001b62007dcb565b620099477f4690e51229e5d2de6165d161580ffce0455403dcfb32984548987fc613cad28360001b62007dcb565b620099757f099de8c4bfce0e0d1bed7d4b04bbf611f0ef75587a2746cdbe90d5547f3090b360001b62007dcb565b6001836080019015159081151581525050620099b47f9d14d9f0a0e742071225cf533990a5717c3bc708c0097d3533606177d0d4bcb860001b62007dcb565b620099e27f51381c4c3a12663aeeacb609607249f71575ad991b5ea4d31f013dc2b859186760001b62007dcb565b6001955062009a1b565b62009a1a7fafd347aee829f2ba5a4eb2e589cdb253c1f7db48c5fad8d5aa5e58774c9dfbff60001b62007dcb565b5b62009a497f512457dc556dce832118a6ebfcac21e75a9b37101b8d9662eee35a7117cf4ec460001b62007dcb565b82604001805180919062009a5d90620112e3565b67ffffffffffffffff1667ffffffffffffffff168152505062009aa37f550ef2fb7c49fc6d77cbb8b123eb957697564572ab6590bb86bb1bf390b785b560001b62007dcb565b5062009aee565b62009ad87f3f83e125c5326eb100196d2919e8cf2fd9387f7a020320286028121c4d6c9a6660001b62007dcb565b808062009ae59062011295565b915050620095ec565b5062009b1d7f2a44b86db298a125303398a7c9f6247ec8d63b57eba0f09ad6e45f3c7d6985e560001b62007dcb565b62009b4b7f5d6e82547a1f04db5eee28d104471e90988c714f638279121241e3bc845074d460001b62007dcb565b8162009bed5762009b7f7f69950c9a4b665e9c725182ed38d7a21fd7de7539d4a42989d4da0c40a6d831c960001b62007dcb565b62009bad7f104fb3edcd520a754865432d4fc337660c1ba85ac865acd1819cb8950eefa39560001b62007dcb565b62009bdb7f5eaeca7f2a1742f3093024cb0e303f902f8df3f330bcae237998f355ef66706060001b62007dcb565b60009850505050505050505062009fc9565b62009c1b7ff91aebc4718dba927c7467b1d042f76381450d86f8933b8b8da03c4078a3eae160001b62007dcb565b62009c497f36250df4dd8b183a7fb1b421c5c3b663d2ea970d2a334ca1351cb1fa206a742660001b62007dcb565b62009c777faca70f6d17ebb05e61824afc24478df44e73a278c4c4b65b2e669c76ec88412d60001b62007dcb565b62009c8786600001518662002bc8565b62009cb57fb04d4fb548cc63eed460ea30fd245295d36db2096d4c01969a37b7efc8d6ea6660001b62007dcb565b62009ce37fde26ac31f03b73db5828c190c68bca0be18bb6e1f1136dfe37be7dd39e3c331160001b62007dcb565b831562009f1b5762009d187f7f5349473afd2977e2d77d957f818ecfb27c0e56efb85cbe37f948104d5db24b60001b62007dcb565b62009d467ff917801a0e6c0c648c006cc4b9f9157d81fc37e93563c376dcd92fdbd7cef9c960001b62007dcb565b62009d747fdb723c2bff791c15de5abf1f3740f39638634c2a492b23456ecacd87fa7728db60001b62007dcb565b62009d83868c83888e62003c0e565b62009db17f141104dea343b265cfc021c46cd8e3f763013c2649cbb1fe60df109c5617c31960001b62007dcb565b62009ddf7f32afae8a03e621772dd9d046dbdfb399ffe862f15b86a602ffb8a6c343c7e36860001b62007dcb565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166247c0038d886040518363ffffffff1660e01b815260040162009e3d929190620109f9565b600060405180830381600087803b15801562009e5857600080fd5b505af115801562009e6d573d6000803e3d6000fd5b5050505062009e9f7fa7a1b85f95d5a698d15ad8716b7bb26c27b15111c5167736b7e6e3b33cf7b64060001b62007dcb565b62009ecd7fcd9d34a640bfb0b7c6e4431dff94853584f69087123a16e689e8a1ee6b503e1460001b62007dcb565b7fb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f4560014388600001518e60a0015160405162009f0d9493929190620106ed565b60405180910390a162009f4a565b62009f497f64463b3609d561de03d551e0013ce32684a254644fb78cfc13b56508195f401a60001b62007dcb565b5b50505050505050808062009f5e9062011295565b915050620091c1565b5062009f967faa198af1e59bcf78448189ddcf19d88221575de4199d505c24b2bffb5eb3704760001b62007dcb565b62009fc47f4d62b62944ed4e74438e5a6d9a1578e21f76e2a1db10c06300db7ab3797a8e1760001b62007dcb565b600190505b9392505050565b60006200a0007fd141d1c821d4478f556803ba4ce1f7c5ced656b12254a78f03e5d9f33f6ecc3160001b6200cac2565b6200a02e7f4e9a6aaeee87e16b21bb377a83303e28b9f2b7dc2019eb6bc2944be13011910e60001b6200cac2565b6200a05c7fbfe18e57a3f2cabc5d872e43bb674ce1cf8916b83df9427e4e3098fcfdee136060001b6200cac2565b60008460000160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015490506200a0d37fcca473718337ff5da4b6fadff922afcfa023163de0585980a1b1e13adae34da960001b6200cac2565b6200a1017f43d33b4f220709ac8536a38ac78e1b0b15d197fba5c427abdcb2a925bf25271960001b6200cac2565b828560000160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506060820151816002015560808201518160030160006101000a81548160ff0219169083151502179055509050506200a25e7f3dbf6506da0f8e4a232a6cbd0d6fdbcc46a6159b1b57bf674f6978070f43e07c60001b6200cac2565b6200a28c7f4cf39f8763a1a3902f1e72d6b2f34b943edf45963c2ae5a2ae6c1bd00f24008d60001b6200cac2565b60008111156200a2fd576200a2c47f9faf1ee22fde19560011cedda1b06b6c953045ea7d03e4bdf847a9dc9ef5059f60001b6200cac2565b6200a2f27f7d63221be0372f37731dcc32697c8e5de73379fb45a6ac7a77060f363c35829f60001b6200cac2565b60019150506200a650565b6200a32b7fc7093f5578794c496533b8c5a7994544b3efa0689d1a8f9aa1ac506b966f45cc60001b6200cac2565b6200a3597fe9b379ca2bbc5e1f9409d7614d58ac42c54d3e3d03ff54a0f56026a9351fa5ee60001b6200cac2565b6200a3877fe387f0e2b02cdd2bad03d9905d392476fdc656054a9961a5b577a01f7cd5978460001b6200cac2565b846001018054905090506200a3bf7f8cf11e2e078ffa55c325627aa16e47f8f848d61673153358151244df4e12dcea60001b6200cac2565b6200a3ed7f993b5c039a2de062246b4ea17865891aa7773ea915b4f1d4ce04bb1e996fa87e60001b6200cac2565b846001016001816001815401808255809150500390600052602060002050506200a43a7f586393b1dcc16d25d5833f9e1a05de9a23f821e316f6a7011556d3b729be1abb60001b6200cac2565b6200a4687ff114c5545f9a950732b6919d235a334c6190e5f07b19823aacfa12c433f1c0a960001b6200cac2565b6001816200a477919062010e62565b8560000160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001819055506200a4ed7f683532729e017ac735f77036d1e9bf46cd05c1a09ee7b1e138d52bc0dbde0c2660001b6200cac2565b6200a51b7fda11dc8ef37fc8e4b64fe565597f064a54d7712b451b12a17c7c6819ce7c926e60001b6200cac2565b838560010182815481106200a559577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506200a5d27f6a3a45a757ef250d7b372805f67274e349618b179f559fb81b8489bf16370d4160001b6200cac2565b8460020160008154809291906200a5e99062011295565b91905055506200a61c7f50948d4a08e96cf4b026c117649246e8edb86fbbcc421528651502e5663df54260001b6200cac2565b6200a64a7ffacb90579ea92a10f0cc3136c8f1947142b7b7ff6e42102f5c8b430fd910740660001b6200cac2565b60009150505b9392505050565b60006200a6877f7c1ae9cb9a5ecb22440cb57ad0521926f6e09bfae9690cc35207185ee3f87fc360001b62007dcb565b6200a6b57f4fae10ebcf9c22362664d1dfc08028b17b455067af9b995c7ba89d4ad6fd3de960001b62007dcb565b6200a6e37f3df189fbd98bae753489cc676e6659e5e4bd24ef2694c3d043bd642d3f49dced60001b62007dcb565b600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002054905092915050565b60006200a7807f9d16bf398be3c245c0af14b0f8c3f0dfac777bd7b36f90e4c98af3bcacc0683d60001b62007dcb565b6200a7ae7f82a68b8708234b8f61118a41d1038e5cd53cfc5128105d12bd2bced2fb3bf18c60001b62007dcb565b6200a7dc7fc9d843b4d545a896b03b19af8e972cc658ace049573369bf4d62fe51cac8281a60001b62007dcb565b60008460c0015190506200a8137ffa1152683f5955bea192dcb4d69b10c67aa843dfa6238922c7367e131572efad60001b62007dcb565b6200a8417f24b0161cf09012b815d11e678961ceb9b420acf3b018748ed8c2a1bb42a5453560001b62007dcb565b60008660c0015190506200a8787fd8ce2e292fda381d0e0e39505ff235bbac68f353a5f259ca334047a453b2bb2e60001b62007dcb565b6200a8a67f01fdc3b429b0d1c9efcd1684b3b2f7b1a7446e63e5ade377fbbc9fb8338d2c9f60001b62007dcb565b838267ffffffffffffffff16826200a8bf919062010e62565b106200a95b576200a8f37fd7b8ca8b74eaa6f2356e3d495f3c4fa7b1d32be71cf6fc3089f375dd5e06ba7b60001b62007dcb565b6200a9217f71907692e79e3b7bd3b17b23da102c4ffa8d754acb6cee4620838d3ec2e56b7360001b62007dcb565b6200a94f7f15ebb0ccbb6b038eba119ac6fbd8c68f20a93d4b5d265dd9bc66adc9c4bdbf8260001b62007dcb565b6000925050506200ae92565b6200a9897f07ee16e4afee4f842416a060e70ed891a26243278cad8bcc06426ce663830f1960001b62007dcb565b6200a9b77fb33a70fb115b11df890b074f5b502b99105c9f3fe4fed35a1a462955e2b21ab660001b62007dcb565b6200a9e57f7b3fe115c3901837c158d7fbdeba0486e080062640db7e949964db7c8365850460001b62007dcb565b60008282866200a9f6919062010f85565b6200aa02919062010f04565b90506200aa327fe3b123bac2e24150c40dfe232a196aeff1794cde9ae4946b1be376c73c69587b60001b62007dcb565b6200aa607fa71188729b2cbb8ce9a6e1b3ec1fda884b14f32d3e96ec347b846db9cb80080260001b62007dcb565b60006200aa907f58e91497643c1580ba8fce5f7cacf6d51f8322a10bb38835a7e220464ca385d760001b62007dcb565b6200aabe7fed295cc3993b9662909eb75fcbd087d3cbc84a36f674ebb051a0345d90b436d060001b62007dcb565b600087146200acb3576200aaf57f68fb92549ba302f5f6d68f9db6db74dc94cb9969a0bd4104ae3e03af4e57d05a60001b62007dcb565b6200ab237fec2a4853956d1a4ceee7d97ae5465e3ad9ca72d081ccd813eddd65e3b5751db260001b62007dcb565b6200ab517f3ef1f21028d4433515328e55c6fa6a05e1183951db140f5dcd72436c80c32b1960001b62007dcb565b8367ffffffffffffffff16836200ab69919062010e62565b8711156200ac1e576200ab9f7f553018eedbae47833e54f385743afe14888cadbb2c126e65919fc54affaf3aff60001b62007dcb565b6200abcd7fa7575b482d181c385c2ba39b309484c6f7cb6d29b454ce02f0e316c06e4dbd4560001b62007dcb565b6200abfb7f4debc683fe29d3539159ef3f14990d7b13452b8fedd8380d24cb85ffaff1cea660001b62007dcb565b8383886200ac0a919062010f85565b6200ac16919062010f04565b90506200acad565b6200ac4c7f884af1070517ac2ed45f38baad5dda9644a296a30b27eb93f2524e08ed4ec1d560001b62007dcb565b6200ac7a7fbcfc721eecabde701227005b095dcd00eff448be4cdf428465f3dac8976a722060001b62007dcb565b6200aca87f115c79494652cfe53c817163dbc759d0600f09382ed80b9107170bc45dfa858060001b62007dcb565b600090505b6200ace2565b6200ace17f16efee67891375a520960d2e92e20fa0017d6dec9a763524657631589ce43bd960001b62007dcb565b5b6200ad107fdb2472675f915eb6e8b0a0de1cc84a3ffef954c1e44773805f219b0c3e9f4ac760001b62007dcb565b6200ad3e7f87160bfe64d1541eab8e5a68fb1642ad5c0171073914851607ece21840e1706160001b62007dcb565b8067ffffffffffffffff168267ffffffffffffffff1610156200adf3576200ad897f63922a29142a0358345154a478cd6a8faf4f6f8d0270895d787fc3b1a190119660001b62007dcb565b6200adb77fabd395443da03e3ad8ede4513f64207d653102e79d504e73a7353fea456ccb6860001b62007dcb565b6200ade57fe680dfd37d93155d4e2d4b4b6ce1e34e38c136db1327b57d7fdd61acfd394c3460001b62007dcb565b60009450505050506200ae92565b6200ae217f39ec3013285af66a7ade83e184f95b09974504dcf49e2d5f7533f949d773f35b60001b62007dcb565b6200ae4f7f59bae41e7fe4b67d67cedb9d7b8acda2ff75d9124e5bddfb821572e03e13234760001b62007dcb565b6200ae7d7ff70df2859d958d2321d2ac45f874391ef1d5354bc8fade31ddc10a257a7acf6960001b62007dcb565b80826200ae8b919062010fc0565b9450505050505b949350505050565b60006200aeca7fb3cfe50daf8ef3a3cdba7a7b0ef0b28688e263feddc6bf90df6ce97be64872c960001b62007dcb565b6200aef87f0301213e8aea82888992c1c4744f671e3d14dc13b49e6bdc4a1bfe0a803d1e2260001b62007dcb565b6200af267f97ea5e44a15d15a08569b8d4a2c7b71b7a22ed152c582b7f5ccd59944391e89660001b62007dcb565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cc76e80d84600187604001516200af7a919062010fc0565b60008960a001518a608001516200af92919062010f3c565b8a6101a001518b61010001516200afaa919062010f85565b6040518663ffffffff1660e01b81526004016200afcc95949392919062010a51565b60606040518083038186803b1580156200afe557600080fd5b505afa1580156200affa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200b02091906200f33d565b90506200b0507fc22c2cc4502ba745b46b8e37aa66e23529ccce8decbfe10f394065802fef609960001b62007dcb565b6200b07e7fa847c562115d4ab05e5fc04fc2578360c532990a09c849278caa38d61f574c5b60001b62007dcb565b8060400151816020015182600001516200b099919062010ebf565b6200b0a5919062010ebf565b9150509392505050565b60606200b0df7fa1c1531ab3acf95261d9b2969670ab5b4008c9228fc856516ea7ccae7bfe21ee60001b62007dcb565b6200b10d7f11430196cc337324af63b84a80f51a8234cfc39d6999fe5662de45e84c94071560001b62007dcb565b6200b13b7fad3069dda3b0afd2ebc3198991c3e049280ce0378a34b66e4c0f49b00ffd95c960001b62007dcb565b60006200b1488362007dce565b90506200b1787fe3772f0a256e29e0cc62671400c74ef97b75f9ed6be3ca06722bf440eb28ae0860001b62007dcb565b6200b1a67fc567b8a687cef7503f084b0ea0e2737de8c96d009afe1a129fd04974a7ed62cf60001b62007dcb565b60005b81518110156200b3f9576200b1e17fff34ce214d3c58b6e41fa49793c512eae8b807d323f6fbfbaa74b493d0c34e2660001b62007dcb565b6200b20f7feb9afd447d0119eec6c0bf812d6bcb349a1eb45e0435856a057de64e3f125ab060001b62007dcb565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a65374a8484815181106200b289577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151602001516040518263ffffffff1660e01b81526004016200b2b39190620105c6565b60e06040518083038186803b1580156200b2cc57600080fd5b505afa1580156200b2e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200b30791906200f0c2565b90506200b3377fe65b9897435d700c8304fa5ca5ff6ff9dd29bd0733e551d15b75fc36c5cc345660001b62007dcb565b6200b3657f5237757807b06ddfe1cc56cdda044d550cb19fab8dca0a6d0fae6478381d13a260001b62007dcb565b8060c001518383815181106200b3a4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505080806200b3f09062011295565b9150506200b1a9565b506200b4287f537c13e5110ceec19a55abd540176c6136e77cfdd8746061ae6984d5301fbeba60001b62007dcb565b6200b4567fc1c6b1599e638b5eba413331567c540a63623696fa47257665315d642ad0b0d060001b62007dcb565b80915050919050565b60006200b48f7fc2b3b51d26d7c6e045c54a94fc06eb36c13011911f2ca3bf9dce12714bd404a660001b62007dcb565b6200b4bd7fab1b6ae163aa2e0fec62d6dbb52f562b0c11107961013b05421348d52d3e73eb60001b62007dcb565b6200b4eb7f2278269b1cfbc50c666fc5331539577da103648da63bc3f5efae711797c38e4060001b62007dcb565b60004390506200b51e7fd48153f1bc8afeae9380a460749ee46c5190d9c419f6d37a905781664d52259360001b62007dcb565b6200b54c7f4ff685308f415e073f1b4afacccb2d5c984eafdf210e4dc8e5d5d539189eebfe60001b62007dcb565b8260c0015167ffffffffffffffff16816200b568919062010e62565b846040015111806200b5985750808360c0015167ffffffffffffffff1685604001516200b596919062010e62565b105b156200b633576200b5cc7fa6de4e2fbca443dd7a19e6a98ee53b5e9aa4eb794feb4cb1b2bd4d6af870bc5860001b62007dcb565b6200b5fa7fecceaebcda79444288bfb34d12365eb3523f0c40d48361ab5cf5221bb1e496ab60001b62007dcb565b6200b6287f82ff7e7079e511bff144d646f26b9ae8af1bda6d3b185962f6a7578473e6f89360001b62007dcb565b60009150506200c100565b6200b6617f32ca81f6229f93f82b8ffadd1fc84c5cbcf31dd5808ce03001bc019bb5d2143360001b62007dcb565b6200b68f7f9158d9aabe2e8c73bf234598aae1bce7419b34d1d83dff4f864baa7911387c2760001b62007dcb565b6200b6bd7f7b31e6448150e378fd424f78eeed9d0e8226b62145d2463f33d115e5b9664e1660001b62007dcb565b6200b6c76200cde7565b6200b6f57febf274387122f6e647dd3e8f12d50be88a139515909a6133eee62b19d322072f60001b62007dcb565b6200b7237f7b39fbb6c2bbe3c19b89e4e27fc187a46d865c2762683f1c348ffcbe68dbeeff60001b62007dcb565b6200b72d6200ce01565b6200b75b7ffd9e830571f8761de93e3cb19849f06163fb254427bc9eeeb432377ee64113d060001b62007dcb565b6200b7897f7ab79d94dabee964058a9156f66230c3ecc33e58410c344bc875694b72b9763160001b62007dcb565b60606200b7b97fe372bd25dc7f9e6fa454a37bc1a960ecd5969c34887feddc6e335c15a88aa66860001b62007dcb565b6200b7e77f15121c0b850669f94a78525cd1b5a71bc16d76eeffcf26c9776dfe842a75f55760001b62007dcb565b6200b7f16200cb8e565b6200b81f7f269d7965001c76d5597794a39134b6dfd6db0b6c520d21bdea5d596d077715f760001b62007dcb565b6200b84d7f65c939eb1e3eab9f9ef7dac2863f6e5ea4441a821470557e9211e2eaf960c20c60001b62007dcb565b8760600151816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506200b8b77fb9fe5d4f938de3d1e9c42c65b572ee1dce7f1331160ec9d9a2291e299465264e60001b62007dcb565b6200b8e57f9127c5d15a7e2978b988b5c76178781fa2462b214fffd991aa8846e1573bed1260001b62007dcb565b8181602001819052506200b91c7fbfdf8ac41f529f29541e8a14089d58e6c9b9518eeab096ce53b3e1ba9471033560001b62007dcb565b6200b94a7f15655164e6a675a4d407e61760e80d4d321d47852e73baca1fcd4b8ed45bca9a60001b62007dcb565b8660800151816040019067ffffffffffffffff16908167ffffffffffffffff16815250506200b99c7f6bf238f3442036d704b7977b2097f0f864574275e838b7f5710a7d9efbc9e28460001b62007dcb565b6200b9ca7fe69f0c008700781b84f176d21b4b5ca692cfeff2734b280212f3c640caaac9ca60001b62007dcb565b866101800151816060019067ffffffffffffffff16908167ffffffffffffffff16815250506200ba1d7f0ed1dac139e052371f8a59c72b6517c9d7c0ccab730f810767bcc1234eef84b660001b62007dcb565b6200ba4b7f168e26c5a2db2ebdd536541a6065fd1d5d232d76164b19c1f9de70812a61f2a860001b62007dcb565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fd7c9808836040518263ffffffff1660e01b81526004016200baaa919062010953565b60006040518083038186803b1580156200bac357600080fd5b505afa1580156200bad8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906200bb0391906200ed3b565b90506200bb337f70f8460e0df1e662bdd3bb161fa62c91eaa4f3869d0b6c3c18315ca4d551b6f960001b62007dcb565b6200bb617f428e1bcaf18ec42118c981ece3690cc01c78ce3b8cde95b33e4710cd3de11f2860001b62007dcb565b6200bb6b6200cca2565b6200bb997f447e2a3f9e6091f2e06b7a7a44bb6c67c1ea14ee04ac9fc6a4e6883e0ef1d98e60001b62007dcb565b6200bbc77f37d1caacf35ea27e0640500ee71b009bf02398ad7ed06660ca087c41f6950fbf60001b62007dcb565b6000816000019067ffffffffffffffff16908167ffffffffffffffff16815250506200bc167f6b9c3c081d57f84e9d92be63a74710806732ac28519f4a2ce6616ed4cbf7987560001b62007dcb565b6200bc447f3050161797bfe99a59b154f918a93d7afa3f9c0196bac0bbf238264a383ae61f60001b62007dcb565b846000015181602001819052506200bc7f7f4e67a1a494c30a7969749327b462d4aef97c7499e524d19e44ab2307be60bc2f60001b62007dcb565b6200bcad7ff06e594d70a43e13e0662eda226446f76525970e3e075f4d61aa0ca8825b1e0b60001b62007dcb565b856020015181604001819052506200bce87fb7a771bacad19bf13eec124b19f32ccc8d2a89bb868ec47132dee8963470617a60001b62007dcb565b6200bd167f0a61728aa13ff8c9161f312c8d0c285c4f9e24c01624a4b632e3713b6ba7b75160001b62007dcb565b846040015181606001819052506200bd517f6ad0207da14cab32652c3cc3615e044385d6262afc32f29ee463d4eb5d11effa60001b62007dcb565b6200bd7f7fcf50f72305df6a8cbd1f8b829fd8f89a06959167862fe93886bfafe8471672ea60001b62007dcb565b8181608001819052506200bdb67f2877180e16a1e3a2a31c9fe4361110200531f46bab258e0463d68a2da88cab8560001b62007dcb565b6200bde47fed92ef99784f600c0a28bc427136a55a1bb53aa178da38ae135312742fb2a43460001b62007dcb565b84606001518160a001819052506200be1f7f77a4b74e3dbc6038ee5eafad46df982db96edf66105f7facf14609869b42ef8360001b62007dcb565b6200be4d7f22a1c1fdc2d9d141f1ded49c3f41c0e59dd06731a130319938f252c1c96c32b960001b62007dcb565b85600001518160c001819052506200be887f5cebc0aeb10717c4e10ffa3996f9df7b073129404c38ee669681d8cf85fa8df460001b62007dcb565b6200beb67f4ab226050d857943558606a4011a47904bd7d94b3b51ff34c71d96b81c9b4ac960001b62007dcb565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b6ddeca0836040518263ffffffff1660e01b81526004016200bf15919062010ad7565b60206040518083038186803b1580156200bf2e57600080fd5b505afa1580156200bf43573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200bf6991906200ed80565b90506200bf997f990a4b325c824bda58530e775531f68e5e05708119980c1f6b9ea94afa6b174460001b62007dcb565b6200bfc77f6c32ff0ad4c8bb80357afd3065745afdb3cefa0538f813750697a94b12f62d6b60001b62007dcb565b806200c069576200bffb7f7db5dbee663494370987686f6ab84c5f3547de9ab1f5d7b492cc11147467e8eb60001b62007dcb565b6200c0297f49a7146638182a662761a7666f9a6c8ad0fef03e65220f4b6e843e52f50a8d8c60001b62007dcb565b6200c0577f923bb086fd2bfd4670edb92514242b35c09f60b616fa50aa498a1330ac02043260001b62007dcb565b6000985050505050505050506200c100565b6200c0977f22c09e3276608fc99f94aa4211b7fd55ea44e64a51616ec4bdae2432425274d560001b62007dcb565b6200c0c57fedc0aee1c377c31376f840125a9259f020683c77895dcaf8a2d0b4561593a59160001b62007dcb565b6200c0f37f57438687a312b8020421f070ee48b6e229b231f3b46cf9f5da32a0797da4c98460001b62007dcb565b6001985050505050505050505b92915050565b60006200c1367f63b82f01d2eadf86878c7430dc7186ceaca7ad439ca3f313f9969fbd5babcdca60001b62007dcb565b6200c1647fc1bf39448ec93245a27d8d0ae63c1b8b1362d90e6e3fc8788c66995f04a71a0160001b62007dcb565b6200c1927f4f0e318f55ec82b3a61015222cae868c2ac671e3a57bb96df97aeda8d79d14ba60001b62007dcb565b814311156200c22f576200c1c97fc4b1dbc24c2e4223eae0144af6b6c8031dc8346c721fdf2e701ab9096e0b5d8c60001b62007dcb565b6200c1f77fe4839430c604ef5ce3664ce4f680ff051af6f328bbf7c4d57d8c831ca534fd8760001b62007dcb565b6200c2257f2d92f7571e52cf83dcb54655202b391c9321c01149c1117bf4b2d8f4447f0e2260001b62007dcb565b600190506200c2be565b6200c25d7f86ebe836082c5f1cd2eaebd85b7d50416d5e7ae3233cae37027b1189a6c33cdb60001b62007dcb565b6200c28b7f1c84f7acbc5bbed589b86e08a7bf69257274820542d43e4e0553c88058666cf460001b62007dcb565b6200c2b97fdb6d2b7fa2a881c4a0987efc14ce99931634796541a4f6958e93c053c407019f60001b62007dcb565b600090505b919050565b60006200c2f37fef24a6af12e53070e9e4733e1f62ba265b957aeed1a255b1c10cfaf955b2eb2660001b6200cac2565b6200c3217f9f6729093e5604f53a3aa28319536264fd8961fb59557b43a3750b85068bbe8a60001b6200cac2565b6200c34f7f38e791e6a450741fa1c87a4fd34a69b3cd97b7b7f2024926687c86b54efd75c960001b6200cac2565b60006200c35e8360006200c72b565b90506200c38e7f433440103c2e88580e58067630275418c60ec540504ccb9372a29df875a516bf60001b6200cac2565b6200c3bc7f1924095f638c56734ab6c4365e452317b456447cc700f043628331c27ebd78af60001b6200cac2565b6001816200c3cb919062010f85565b915050919050565b60006200c4037f316e6c9528f6854f31ba3e89e3b3dc63bc0068a4af7fae07ea015bcb566ea05260001b6200cac2565b6200c4317f8f4685549187fd84f2f3158619e2e6fe3afa5c22bf50c87fe8767c0be5cb6f5860001b6200cac2565b6200c45f7fb28c0899bf0ffa7d903849fdab8a36fe8ebf2be4e803020554e3bd038a04103360001b6200cac2565b82600101805490508210905092915050565b60006200c47d6200cd3f565b6200c4ab7f2ee32202b14075dab4d543099a07964c02909d6a7d8d61c6c288e2208b0664eb60001b6200cac2565b6200c4d97f61e4694cbf00b9302c6a7a1efcd3e0074ee3dec089af2fc27ab219c54e8fe66a60001b6200cac2565b6200c5077f84a2caea5f9ff227fe6db47c92fc516d90b642255965900de72ae23a07d95ccb60001b6200cac2565b8360010183815481106200c544577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691506200c5a27f7fce783fd7ea0987a558d13ee100000bf54b2fc787abbd179f70e3a19e19668360001b6200cac2565b6200c5d07f7294e5e53c0e68349e7e238c9387fb16cfc97cd845edb1cf999d936eb5f98b1560001b6200cac2565b8360000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001016040518060a00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600282015481526020016003820160009054906101000a900460ff16151515158152505090509250929050565b60006200c75b7f01a23c7dbf36e917e35a5011b0a9b4af23836e2f32868b23a2214b738e2b57e060001b6200cac2565b6200c7897f6ae211447691763725c956ade578ce08d125ba2b364334314e273ed6a1b5d8eb60001b6200cac2565b81806200c7969062011295565b9250506200c7c77f2b9a5fdd2497d438882650436bde0800c78a48ebe0910b8b82947d0f1fb5e47360001b6200cac2565b6200c7f57f7c662f3ae00e13f864e3fbce308bf3a02ac24dba13ed2d8842636d2ac662180d60001b6200cac2565b5b8260010180549050821080156200c86157508260010182815481106200c845577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160000160149054906101000a900460ff165b156200c8ab576200c8957f72d48d21f2c306d91eb806375cc2254967a74d51eb005856674c4fea2fb1b8ad60001b6200cac2565b81806200c8a29062011295565b9250506200c7f6565b6200c8d97fdf1d1c4a9a6974451a9d0b9458a21371865bbe23d393b53ce018375a372bb59d60001b6200cac2565b6200c9077f99f0a7d60a13110313addb675a0b40c2dfdc51304926dfb4f2592cc5b8f15acd60001b6200cac2565b81905092915050565b60006200c91d306200cac5565b15905090565b60006200c9537f6c2a5f6c037d8863e6fe8813bd47a7c71a761a9fd5a5a2484bc57fb5df0822ba60001b62007dcb565b6200c9817f104ccfd88b5b213e62c5cc6a623f694dec46621d1d81a4bcc5b1bd7f3857c0c260001b62007dcb565b6200c9af7f7c75d777d50ba4832b156ed258d430ba6df5f57e456e8e87f121ada14f161c7660001b62007dcb565b6000600290506200c9e37fb4fe9c9fd00232c0d9dd18667b963f679cba54ea9f5a37313338b77187e195e060001b62007dcb565b6200ca117f7c3742e0ee3df7b4adad58176dab410c0fe0f229f5fc681d45d8fd2c536ea2cc60001b62007dcb565b60006200ca238585606001516200cad8565b826200ca30919062010f3c565b90506200ca607f8363aa485d172a8d9c85fdd107791eec0afcb902ef42830e4297c1d95f2c78d560001b62007dcb565b6200ca8e7f9c8e3b748cbc4eaf4d1173e324acb5077f6b1992befcdbc983ed877a1271917c60001b62007dcb565b809250505092915050565b60008151905060006a636f6e736f6c652e6c6f679050602083016000808483855afa5050505050565b50565b600080823b905060008111915050919050565b60006200cb087f0406d87923fe1fa1e5ee329503d20e485c45fe6032fc7b6a30290fdf2c8bf9fc60001b62007dcb565b6200cb367fb99541977541724c677117c477e674d068a04a3b693b722bff8e6075b971580660001b62007dcb565b6200cb647f09687d2867fd306615aa40a5060bb23993013c7c2b053bd5fc672787d9ea24de60001b62007dcb565b620fa0008284606001516200cb7a919062010f3c565b6200cb86919062010f04565b905092915050565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6040518060c00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081526020016060815260200160608152602001606081525090565b60405180606001604052806200cc3f6200ce33565b8152602001606081526020016200cc556200cbe0565b81525090565b6040518060e0016040528060608152602001606081526020016060815260200160608152602001606081526020016200cc936200cf21565b81526020016000151581525090565b6040518060e00160405280600067ffffffffffffffff1681526020016060815260200160608152602001606081526020016060815260200160608152602001606081525090565b60405180606001604052806200ccfe6200d0bd565b815260200160608152602001600067ffffffffffffffff1681525090565b50805460008255906000526020600020908101906200cd3c91906200d0fc565b50565b6040518060a00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff168152602001600081526020016000151581525090565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600067ffffffffffffffff1681525090565b604051806040016040528060608152602001606081525090565b604051806080016040528060608152602001600067ffffffffffffffff16815260200160608152602001606081525090565b604051806101800160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600060028111156200ceca577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016000815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600015158152602001606081525090565b604051806102e0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff16815260200160008152602001600015158152602001600060018111156200d039577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600067ffffffffffffffff168152602001606081526020016060815260200160608152602001600060028111156200d09f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016000151581526020016200d0b76200d0bd565b81525090565b6040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b5b808211156200d14c57600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549060ff0219169055506001016200d0fd565b5090565b60006200d1676200d1618462010b24565b62010afb565b905080838252602082019050828560208602820111156200d18757600080fd5b60005b858110156200d1bb57816200d1a088826200d5ee565b8452602084019350602083019250506001810190506200d18a565b5050509392505050565b60006200d1dc6200d1d68462010b24565b62010afb565b905080838252602082019050828560208602820111156200d1fc57600080fd5b60005b858110156200d23057816200d21588826200d605565b8452602084019350602083019250506001810190506200d1ff565b5050509392505050565b60006200d2516200d24b8462010b53565b62010afb565b905080838252602082019050828560208602820111156200d27157600080fd5b60005b858110156200d2c057813567ffffffffffffffff8111156200d29557600080fd5b8086016200d2a489826200d7c9565b855260208501945060208401935050506001810190506200d274565b5050509392505050565b60006200d2e16200d2db8462010b53565b62010afb565b905080838252602082019050828560208602820111156200d30157600080fd5b60005b858110156200d35057815167ffffffffffffffff8111156200d32557600080fd5b8086016200d33489826200d7f6565b855260208501945060208401935050506001810190506200d304565b5050509392505050565b60006200d3716200d36b8462010b82565b62010afb565b905080838252602082019050828560408602820111156200d39157600080fd5b60005b858110156200d3c557816200d3aa88826200d8f2565b8452602084019350604083019250506001810190506200d394565b5050509392505050565b60006200d3e66200d3e08462010bb1565b62010afb565b905080838252602082019050828560208602820111156200d40657600080fd5b60005b858110156200d45557815167ffffffffffffffff8111156200d42a57600080fd5b8086016200d43989826200dfd0565b855260208501945060208401935050506001810190506200d409565b5050509392505050565b60006200d4766200d4708462010be0565b62010afb565b905080838252602082019050828560208602820111156200d49657600080fd5b60005b858110156200d4e557815167ffffffffffffffff8111156200d4ba57600080fd5b8086016200d4c989826200e052565b855260208501945060208401935050506001810190506200d499565b5050509392505050565b60006200d5066200d5008462010c0f565b62010afb565b905080838252602082019050828560a08602820111156200d52657600080fd5b60005b858110156200d55a57816200d53f88826200e562565b845260208401935060a083019250506001810190506200d529565b5050509392505050565b60006200d57b6200d5758462010c3e565b62010afb565b9050828152602081018484840111156200d59457600080fd5b6200d5a18482856201121a565b509392505050565b60006200d5c06200d5ba8462010c3e565b62010afb565b9050828152602081018484840111156200d5d957600080fd5b6200d5e684828562011229565b509392505050565b6000813590506200d5ff81620114e2565b92915050565b6000815190506200d61681620114e2565b92915050565b600082601f8301126200d62e57600080fd5b81356200d6408482602086016200d150565b91505092915050565b600082601f8301126200d65b57600080fd5b81516200d66d8482602086016200d1c5565b91505092915050565b600082601f8301126200d68857600080fd5b81356200d69a8482602086016200d23a565b91505092915050565b600082601f8301126200d6b557600080fd5b81516200d6c78482602086016200d2ca565b91505092915050565b600082601f8301126200d6e257600080fd5b81516200d6f48482602086016200d35a565b91505092915050565b600082601f8301126200d70f57600080fd5b81516200d7218482602086016200d3cf565b91505092915050565b600082601f8301126200d73c57600080fd5b81516200d74e8482602086016200d45f565b91505092915050565b600082601f8301126200d76957600080fd5b81356200d77b8482602086016200d4ef565b91505092915050565b6000813590506200d79581620114fc565b92915050565b6000815190506200d7ac81620114fc565b92915050565b6000813590506200d7c38162011516565b92915050565b600082601f8301126200d7db57600080fd5b81356200d7ed8482602086016200d564565b91505092915050565b600082601f8301126200d80857600080fd5b81516200d81a8482602086016200d5a9565b91505092915050565b6000813590506200d8348162011530565b92915050565b6000813590506200d84b816201154a565b92915050565b6000813590506200d8628162011564565b92915050565b6000813590506200d879816201157e565b92915050565b6000813590506200d8908162011598565b92915050565b6000813590506200d8a781620115b2565b92915050565b6000815190506200d8be81620115b2565b92915050565b6000813590506200d8d581620115c3565b92915050565b6000815190506200d8ec81620115c3565b92915050565b6000604082840312156200d90557600080fd5b6200d911604062010afb565b905060006200d923848285016200ec5f565b60008301525060206200d939848285016200ec5f565b60208301525092915050565b600061032082840312156200d95957600080fd5b6200d9666102e062010afb565b9050600082013567ffffffffffffffff8111156200d98357600080fd5b6200d991848285016200d7c9565b60008301525060206200d9a7848285016200d5ee565b602083015250604082013567ffffffffffffffff8111156200d9c857600080fd5b6200d9d6848285016200d7c9565b60408301525060606200d9ec848285016200ec76565b60608301525060806200da02848285016200ec76565b60808301525060a06200da18848285016200ec76565b60a08301525060c06200da2e848285016200ec76565b60c08301525060e06200da44848285016200ec76565b60e0830152506101006200da5b848285016200ec31565b610100830152506101206200da73848285016200ec76565b610120830152506101406200da8b848285016200ec76565b6101408301525061016082013567ffffffffffffffff8111156200daae57600080fd5b6200dabc848285016200d7c9565b610160830152506101806200dad4848285016200ec76565b610180830152506101a06200daec848285016200ec31565b6101a0830152506101c06200db04848285016200d784565b6101c0830152506101e06200db1c848285016200d8c4565b6101e0830152506102006200db34848285016200ec76565b6102008301525061022082013567ffffffffffffffff8111156200db5757600080fd5b6200db65848285016200d61c565b6102208301525061024082013567ffffffffffffffff8111156200db8857600080fd5b6200db96848285016200d61c565b6102408301525061026082013567ffffffffffffffff8111156200dbb957600080fd5b6200dbc7848285016200d7c9565b610260830152506102806200dbdf848285016200d896565b610280830152506102a06200dbf7848285016200d784565b6102a0830152506102c06200dc0f848285016200e397565b6102c08301525092915050565b600061032082840312156200dc3057600080fd5b6200dc3d6102e062010afb565b9050600082015167ffffffffffffffff8111156200dc5a57600080fd5b6200dc68848285016200d7f6565b60008301525060206200dc7e848285016200d605565b602083015250604082015167ffffffffffffffff8111156200dc9f57600080fd5b6200dcad848285016200d7f6565b60408301525060606200dcc3848285016200ec8d565b60608301525060806200dcd9848285016200ec8d565b60808301525060a06200dcef848285016200ec8d565b60a08301525060c06200dd05848285016200ec8d565b60c08301525060e06200dd1b848285016200ec8d565b60e0830152506101006200dd32848285016200ec48565b610100830152506101206200dd4a848285016200ec8d565b610120830152506101406200dd62848285016200ec8d565b6101408301525061016082015167ffffffffffffffff8111156200dd8557600080fd5b6200dd93848285016200d7f6565b610160830152506101806200ddab848285016200ec8d565b610180830152506101a06200ddc3848285016200ec48565b6101a0830152506101c06200dddb848285016200d79b565b6101c0830152506101e06200ddf3848285016200d8db565b6101e0830152506102006200de0b848285016200ec8d565b6102008301525061022082015167ffffffffffffffff8111156200de2e57600080fd5b6200de3c848285016200d649565b6102208301525061024082015167ffffffffffffffff8111156200de5f57600080fd5b6200de6d848285016200d649565b6102408301525061026082015167ffffffffffffffff8111156200de9057600080fd5b6200de9e848285016200d7f6565b610260830152506102806200deb6848285016200d8ad565b610280830152506102a06200dece848285016200d79b565b6102a0830152506102c06200dee6848285016200e400565b6102c08301525092915050565b600060c082840312156200df0657600080fd5b6200df1260c062010afb565b9050600082013567ffffffffffffffff8111156200df2f57600080fd5b6200df3d848285016200d7c9565b600083015250602082013567ffffffffffffffff8111156200df5e57600080fd5b6200df6c848285016200d7c9565b60208301525060406200df82848285016200ec31565b60408301525060606200df98848285016200d5ee565b60608301525060806200dfae848285016200ec76565b60808301525060a06200dfc4848285016200ec76565b60a08301525092915050565b6000606082840312156200dfe357600080fd5b6200dfef606062010afb565b905060006200e001848285016200ec8d565b60008301525060206200e017848285016200ec8d565b602083015250604082015167ffffffffffffffff8111156200e03857600080fd5b6200e046848285016200d7f6565b60408301525092915050565b6000604082840312156200e06557600080fd5b6200e071604062010afb565b905060006200e083848285016200ec8d565b600083015250602082015167ffffffffffffffff8111156200e0a457600080fd5b6200e0b2848285016200d6fd565b60208301525092915050565b600060e082840312156200e0d157600080fd5b6200e0dd60e062010afb565b905060006200e0ef848285016200ec76565b60008301525060206200e105848285016200ec76565b60208301525060406200e11b848285016200ec76565b60408301525060606200e131848285016200ec76565b60608301525060806200e147848285016200ec76565b60808301525060a06200e15d848285016200d5ee565b60a08301525060c06200e173848285016200d5ee565b60c08301525092915050565b600060e082840312156200e19257600080fd5b6200e19e60e062010afb565b905060006200e1b0848285016200ec8d565b60008301525060206200e1c6848285016200ec8d565b60208301525060406200e1dc848285016200ec8d565b60408301525060606200e1f2848285016200ec8d565b60608301525060806200e208848285016200ec8d565b60808301525060a06200e21e848285016200d605565b60a08301525060c06200e234848285016200d605565b60c08301525092915050565b600060e082840312156200e25357600080fd5b6200e25f60e062010afb565b9050600082015167ffffffffffffffff8111156200e27c57600080fd5b6200e28a848285016200d6a3565b600083015250602082015167ffffffffffffffff8111156200e2ab57600080fd5b6200e2b9848285016200d6a3565b602083015250604082015167ffffffffffffffff8111156200e2da57600080fd5b6200e2e8848285016200d6d0565b604083015250606082015167ffffffffffffffff8111156200e30957600080fd5b6200e317848285016200d72a565b606083015250608082015167ffffffffffffffff8111156200e33857600080fd5b6200e346848285016200d7f6565b60808301525060a082015167ffffffffffffffff8111156200e36757600080fd5b6200e375848285016200dc1c565b60a08301525060c06200e38b848285016200d79b565b60c08301525092915050565b6000606082840312156200e3aa57600080fd5b6200e3b6606062010afb565b905060006200e3c8848285016200ec76565b60008301525060206200e3de848285016200ec76565b60208301525060406200e3f4848285016200ec76565b60408301525092915050565b6000606082840312156200e41357600080fd5b6200e41f606062010afb565b905060006200e431848285016200ec8d565b60008301525060206200e447848285016200ec8d565b60208301525060406200e45d848285016200ec8d565b60408301525092915050565b6000606082840312156200e47c57600080fd5b6200e488606062010afb565b905060006200e49a848285016200d5ee565b60008301525060206200e4b0848285016200ec31565b60208301525060406200e4c6848285016200ec76565b60408301525092915050565b6000602082840312156200e4e557600080fd5b6200e4f1602062010afb565b905060006200e503848285016200ec76565b60008301525092915050565b6000604082840312156200e52257600080fd5b6200e52e604062010afb565b905060006200e540848285016200ec76565b60008301525060206200e556848285016200ec76565b60208301525092915050565b600060a082840312156200e57557600080fd5b6200e58160a062010afb565b905060006200e593848285016200d5ee565b60008301525060206200e5a9848285016200d5ee565b60208301525060406200e5bf848285016200ec76565b60408301525060606200e5d5848285016200ec31565b60608301525060806200e5eb848285016200d784565b60808301525092915050565b600061018082840312156200e60b57600080fd5b6200e61861018062010afb565b905060006200e62a848285016200d5ee565b60008301525060206200e640848285016200ec76565b60208301525060406200e656848285016200ec76565b60408301525060606200e66c848285016200ec76565b60608301525060806200e682848285016200d896565b60808301525060a06200e698848285016200ec31565b60a08301525060c06200e6ae848285016200ec31565b60c08301525060e06200e6c4848285016200ec76565b60e0830152506101006200e6db848285016200ec76565b610100830152506101206200e6f3848285016200ec76565b610120830152506101406200e70b848285016200d784565b6101408301525061016082013567ffffffffffffffff8111156200e72e57600080fd5b6200e73c848285016200d676565b6101608301525092915050565b600061018082840312156200e75d57600080fd5b6200e76a61018062010afb565b905060006200e77c848285016200d605565b60008301525060206200e792848285016200ec8d565b60208301525060406200e7a8848285016200ec8d565b60408301525060606200e7be848285016200ec8d565b60608301525060806200e7d4848285016200d8ad565b60808301525060a06200e7ea848285016200ec48565b60a08301525060c06200e800848285016200ec48565b60c08301525060e06200e816848285016200ec8d565b60e0830152506101006200e82d848285016200ec8d565b610100830152506101206200e845848285016200ec8d565b610120830152506101406200e85d848285016200d79b565b6101408301525061016082015167ffffffffffffffff8111156200e88057600080fd5b6200e88e848285016200d6a3565b6101608301525092915050565b6000608082840312156200e8ae57600080fd5b6200e8ba608062010afb565b905060006200e8cc848285016200d5ee565b60008301525060206200e8e2848285016200ec76565b60208301525060406200e8f8848285016200ec76565b604083015250606082013567ffffffffffffffff8111156200e91957600080fd5b6200e927848285016200d7c9565b60608301525092915050565b6000604082840312156200e94657600080fd5b6200e952604062010afb565b905060006200e964848285016200d5ee565b60008301525060206200e97a848285016200ec76565b60208301525092915050565b600061016082840312156200e99a57600080fd5b6200e9a761016062010afb565b905060006200e9b9848285016200ec76565b60008301525060206200e9cf848285016200ec76565b60208301525060406200e9e5848285016200ec76565b60408301525060606200e9fb848285016200ec76565b60608301525060806200ea11848285016200ec76565b60808301525060a06200ea27848285016200ec76565b60a08301525060c06200ea3d848285016200ec76565b60c08301525060e06200ea53848285016200ec76565b60e0830152506101006200ea6a848285016200ec76565b610100830152506101206200ea82848285016200ec76565b610120830152506101406200ea9a848285016200ec76565b6101408301525092915050565b600061016082840312156200eabb57600080fd5b6200eac861016062010afb565b905060006200eada848285016200ec8d565b60008301525060206200eaf0848285016200ec8d565b60208301525060406200eb06848285016200ec8d565b60408301525060606200eb1c848285016200ec8d565b60608301525060806200eb32848285016200ec8d565b60808301525060a06200eb48848285016200ec8d565b60a08301525060c06200eb5e848285016200ec8d565b60c08301525060e06200eb74848285016200ec8d565b60e0830152506101006200eb8b848285016200ec8d565b610100830152506101206200eba3848285016200ec8d565b610120830152506101406200ebbb848285016200ec8d565b6101408301525092915050565b6000606082840312156200ebdb57600080fd5b6200ebe7606062010afb565b905060006200ebf9848285016200ec8d565b60008301525060206200ec0f848285016200ec8d565b60208301525060406200ec25848285016200ec8d565b60408301525092915050565b6000813590506200ec4281620115d4565b92915050565b6000815190506200ec5981620115d4565b92915050565b6000815190506200ec7081620115ee565b92915050565b6000813590506200ec878162011608565b92915050565b6000815190506200ec9e8162011608565b92915050565b600080604083850312156200ecb857600080fd5b60006200ecc8858286016200d5ee565b92505060206200ecdb858286016200ec31565b9150509250929050565b6000806000606084860312156200ecfb57600080fd5b60006200ed0b868287016200d5ee565b93505060206200ed1e868287016200ec76565b92505060406200ed31868287016200ec31565b9150509250925092565b6000602082840312156200ed4e57600080fd5b600082015167ffffffffffffffff8111156200ed6957600080fd5b6200ed77848285016200d6d0565b91505092915050565b6000602082840312156200ed9357600080fd5b60006200eda3848285016200d79b565b91505092915050565b6000602082840312156200edbf57600080fd5b60006200edcf848285016200d7b2565b91505092915050565b6000602082840312156200edeb57600080fd5b600082013567ffffffffffffffff8111156200ee0657600080fd5b6200ee14848285016200d7c9565b91505092915050565b600080604083850312156200ee3157600080fd5b600083013567ffffffffffffffff8111156200ee4c57600080fd5b6200ee5a858286016200d7c9565b925050602083013567ffffffffffffffff8111156200ee7857600080fd5b6200ee86858286016200d757565b9150509250929050565b600080606083850312156200eea457600080fd5b600083013567ffffffffffffffff8111156200eebf57600080fd5b6200eecd858286016200d7c9565b92505060206200eee0858286016200e50f565b9150509250929050565b60008060008060008060c087890312156200ef0457600080fd5b60006200ef1489828a016200d823565b96505060206200ef2789828a016200d83a565b95505060406200ef3a89828a016200d851565b94505060606200ef4d89828a016200d868565b93505060806200ef6089828a016200d87f565b92505060a06200ef7389828a016200e4d2565b9150509295509295509295565b6000602082840312156200ef9357600080fd5b600082015167ffffffffffffffff8111156200efae57600080fd5b6200efbc848285016200dc1c565b91505092915050565b600080600080600061032086880312156200efdf57600080fd5b600086013567ffffffffffffffff8111156200effa57600080fd5b6200f008888289016200d945565b95505060206200f01b888289016200e0be565b9450506101006200f02f888289016200e562565b9350506101a086013567ffffffffffffffff8111156200f04e57600080fd5b6200f05c888289016200d757565b9250506101c06200f070888289016200e986565b9150509295509295909350565b6000602082840312156200f09057600080fd5b600082013567ffffffffffffffff8111156200f0ab57600080fd5b6200f0b9848285016200def3565b91505092915050565b600060e082840312156200f0d557600080fd5b60006200f0e5848285016200e17f565b91505092915050565b6000602082840312156200f10157600080fd5b600082015167ffffffffffffffff8111156200f11c57600080fd5b6200f12a848285016200e240565b91505092915050565b6000606082840312156200f14657600080fd5b60006200f156848285016200e469565b91505092915050565b6000602082840312156200f17257600080fd5b600082015167ffffffffffffffff8111156200f18d57600080fd5b6200f19b848285016200e749565b91505092915050565b60008060008061028085870312156200f1bc57600080fd5b600085013567ffffffffffffffff8111156200f1d757600080fd5b6200f1e5878288016200e5f7565b94505060206200f1f8878288016200e0be565b9350506101006200f20c878288016200e986565b9250506102606200f220878288016200ec76565b91505092959194509250565b6000602082840312156200f23f57600080fd5b600082013567ffffffffffffffff8111156200f25a57600080fd5b6200f268848285016200e89b565b91505092915050565b600080604083850312156200f28557600080fd5b600083013567ffffffffffffffff8111156200f2a057600080fd5b6200f2ae858286016200e89b565b925050602083013567ffffffffffffffff8111156200f2cc57600080fd5b6200f2da858286016200e5f7565b9150509250929050565b6000604082840312156200f2f757600080fd5b60006200f307848285016200e933565b91505092915050565b600061016082840312156200f32457600080fd5b60006200f334848285016200eaa7565b91505092915050565b6000606082840312156200f35057600080fd5b60006200f360848285016200ebc8565b91505092915050565b60006200f37783836200f413565b60208301905092915050565b60006200f39183836200f79d565b905092915050565b60006200f3a783836200fa0a565b60408301905092915050565b60006200f3c183836200fce4565b905092915050565b60006200f3d783836200fd3a565b905092915050565b60006200f3ed83836200fe5f565b60608301905092915050565b60006200f40783836200ff53565b60a08301905092915050565b6200f41e8162010ffb565b82525050565b6200f42f8162010ffb565b82525050565b6200f44a6200f4448262010ffb565b62011319565b82525050565b60006200f45d8262010ce4565b6200f469818562010da2565b93506200f4768362010c74565b8060005b838110156200f4ad5781516200f49188826200f369565b97506200f49e8362010d47565b9250506001810190506200f47a565b5085935050505092915050565b60006200f4c78262010cef565b6200f4d3818562010db3565b9350836020820285016200f4e78562010c84565b8060005b858110156200f52957848403895281516200f50785826200f383565b94506200f5148362010d54565b925060208a019950506001810190506200f4eb565b50829750879550505050505092915050565b60006200f5488262010cfa565b6200f554818562010dc4565b93506200f5618362010c94565b8060005b838110156200f5985781516200f57c88826200f399565b97506200f5898362010d61565b9250506001810190506200f565565b5085935050505092915050565b60006200f5b28262010d05565b6200f5be818562010dd5565b9350836020820285016200f5d28562010ca4565b8060005b858110156200f61457848403895281516200f5f285826200f3b3565b94506200f5ff8362010d6e565b925060208a019950506001810190506200f5d6565b50829750879550505050505092915050565b60006200f6338262010d10565b6200f63f818562010de6565b9350836020820285016200f6538562010cb4565b8060005b858110156200f69557848403895281516200f67385826200f3c9565b94506200f6808362010d7b565b925060208a019950506001810190506200f657565b50829750879550505050505092915050565b60006200f6b48262010d1b565b6200f6c0818562010df7565b93506200f6cd8362010cc4565b8060005b838110156200f7045781516200f6e888826200f3df565b97506200f6f58362010d88565b9250506001810190506200f6d1565b5085935050505092915050565b60006200f71e8262010d26565b6200f72a818562010e08565b93506200f7378362010cd4565b8060005b838110156200f76e5781516200f75288826200f3f9565b97506200f75f8362010d95565b9250506001810190506200f73b565b5085935050505092915050565b6200f786816201100f565b82525050565b6200f797816201100f565b82525050565b60006200f7aa8262010d31565b6200f7b6818562010e19565b93506200f7c881856020860162011229565b6200f7d38162011407565b840191505092915050565b60006200f7eb8262010d31565b6200f7f7818562010e2a565b93506200f80981856020860162011229565b6200f8148162011407565b840191505092915050565b60006200f82c8262010d31565b6200f838818562010e3b565b93506200f84a81856020860162011229565b80840191505092915050565b6200f8618162011116565b82525050565b6200f872816201112a565b82525050565b6200f883816201112a565b82525050565b6200f894816201113e565b82525050565b6200f8a58162011152565b82525050565b6200f8b68162011166565b82525050565b6200f8c7816201117a565b82525050565b6200f8d8816201118e565b82525050565b6200f8e981620111a2565b82525050565b6200f8fa81620111b6565b82525050565b6200f90b81620111ca565b82525050565b6200f91c81620111de565b82525050565b6200f92d81620111f2565b82525050565b6200f93e8162011206565b82525050565b60006200f9518262010d3c565b6200f95d818562010e46565b93506200f96f81856020860162011229565b6200f97a8162011407565b840191505092915050565b60006200f9928262010d3c565b6200f99e818562010e57565b93506200f9b081856020860162011229565b80840191505092915050565b60006200f9cb60168362010e46565b91506200f9d88262011425565b602082019050919050565b60006200f9f2602e8362010e46565b91506200f9ff826201144e565b604082019050919050565b6040820160008201516200fa22600085018262010531565b5060208201516200fa37602085018262010531565b50505050565b60006103208301600083015184820360008601526200fa5d82826200f79d565b91505060208301516200fa7460208601826200f413565b50604083015184820360408601526200fa8e82826200f79d565b91505060608301516200faa5606086018262010542565b5060808301516200faba608086018262010542565b5060a08301516200facf60a086018262010542565b5060c08301516200fae460c086018262010542565b5060e08301516200faf960e086018262010542565b506101008301516200fb10610100860182620104f4565b506101208301516200fb2761012086018262010542565b506101408301516200fb3e61014086018262010542565b506101608301518482036101608601526200fb5a82826200f79d565b9150506101808301516200fb7361018086018262010542565b506101a08301516200fb8a6101a0860182620104f4565b506101c08301516200fba16101c08601826200f77b565b506101e08301516200fbb86101e08601826200f889565b506102008301516200fbcf61020086018262010542565b506102208301518482036102208601526200fbeb82826200f450565b9150506102408301518482036102408601526200fc0982826200f450565b9150506102608301518482036102608601526200fc2782826200f79d565b9150506102808301516200fc406102808601826200f867565b506102a08301516200fc576102a08601826200f77b565b506102c08301516200fc6e6102c08601826200fe17565b508091505092915050565b60006080830160008301516200fc9360008601826200f413565b50602083015184820360208601526200fcad82826200f79d565b91505060408301516200fcc4604086018262010542565b5060608301516200fcd9606086018262010542565b508091505092915050565b60006060830160008301516200fcfe600086018262010542565b5060208301516200fd13602086018262010542565b50604083015184820360408601526200fd2d82826200f79d565b9150508091505092915050565b60006040830160008301516200fd54600086018262010542565b50602083015184820360208601526200fd6e82826200f5a5565b9150508091505092915050565b60e0820160008201516200fd93600085018262010542565b5060208201516200fda8602085018262010542565b5060408201516200fdbd604085018262010542565b5060608201516200fdd2606085018262010542565b5060808201516200fde7608085018262010542565b5060a08201516200fdfc60a08501826200f413565b5060c08201516200fe1160c08501826200f413565b50505050565b6060820160008201516200fe2f600085018262010542565b5060208201516200fe44602085018262010542565b5060408201516200fe59604085018262010542565b50505050565b6060820160008201516200fe7760008501826200f413565b5060208201516200fe8c6020850182620104f4565b5060408201516200fea1604085018262010542565b50505050565b6060820160008201516200febf60008501826200f413565b5060208201516200fed46020850182620104f4565b5060408201516200fee9604085018262010542565b50505050565b600060608301600083015184820360008601526200ff0e82826200ffc5565b915050602083015184820360208601526200ff2a82826200f53b565b915050604083015184820360408601526200ff468282620101fd565b9150508091505092915050565b60a0820160008201516200ff6b60008501826200f413565b5060208201516200ff8060208501826200f413565b5060408201516200ff95604085018262010542565b5060608201516200ffaa6060850182620104f4565b5060808201516200ffbf60808501826200f77b565b50505050565b6000610180830160008301516200ffe060008601826200f413565b5060208301516200fff5602086018262010542565b5060408301516201000a604086018262010542565b5060608301516201001f606086018262010542565b5060808301516201003460808601826200f867565b5060a08301516201004960a0860182620104f4565b5060c08301516201005e60c0860182620104f4565b5060e08301516201007360e086018262010542565b506101008301516201008a61010086018262010542565b50610120830151620100a161012086018262010542565b50610140830151620100b86101408601826200f77b565b50610160830151848203610160860152620100d482826200f4ba565b9150508091505092915050565b600061018083016000830151620100fc60008601826200f413565b50602083015162010111602086018262010542565b50604083015162010126604086018262010542565b5060608301516201013b606086018262010542565b5060808301516201015060808601826200f867565b5060a08301516201016560a0860182620104f4565b5060c08301516201017a60c0860182620104f4565b5060e08301516201018f60e086018262010542565b50610100830151620101a661010086018262010542565b50610120830151620101bd61012086018262010542565b50610140830151620101d46101408601826200f77b565b50610160830151848203610160860152620101f082826200f4ba565b9150508091505092915050565b600060c08301600083015162010217600086018262010542565b5060208301516201022c602086018262010542565b50604083015184820360408601526201024682826200f79d565b915050606083015184820360608601526201026282826200f79d565b915050608083015184820360808601526201027e82826200f626565b91505060a083015184820360a08601526201029a82826200f79d565b9150508091505092915050565b604082016000820151620102bf60008501826200f413565b506020820151620102d4602085018262010542565b50505050565b61016082016000820151620102f3600085018262010542565b50602082015162010308602085018262010542565b5060408201516201031d604085018262010542565b50606082015162010332606085018262010542565b50608082015162010347608085018262010542565b5060a08201516201035c60a085018262010542565b5060c08201516201037160c085018262010542565b5060e08201516201038660e085018262010542565b506101008201516201039d61010085018262010542565b50610120820151620103b461012085018262010542565b50610140820151620103cb61014085018262010542565b50505050565b600060a083016000830151620103eb60008601826200fe17565b50602083015184820360608601526201040582826200f79d565b91505060408301516201041c608086018262010542565b508091505092915050565b600060e08301600083015162010441600086018262010542565b50602083015184820360208601526201045b82826200f79d565b915050604083015184820360408601526201047782826200f4ba565b915050606083015184820360608601526201049382826200f4ba565b91505060808301518482036080860152620104af82826200f53b565b91505060a083015184820360a0860152620104cb82826200f626565b91505060c083015184820360c0860152620104e782826200f79d565b9150508091505092915050565b620104ff81620110e8565b82525050565b6201051081620110e8565b82525050565b6201052b6201052582620110e8565b62011341565b82525050565b6201053c81620110f2565b82525050565b6201054d8162011102565b82525050565b6201055e8162011102565b82525050565b60006201057282856200f435565b60148201915062010584828462010516565b6020820191508190509392505050565b6000620105a282846200f81f565b915081905092915050565b6000620105bb82846200f985565b915081905092915050565b6000602082019050620105dd60008301846200f424565b92915050565b6000604082019050620105fa60008301856200f424565b81810360208301526201060e81846200f7de565b90509392505050565b600060208201905081810360008301526201063381846200f6a7565b905092915050565b600060208201905081810360008301526201065781846200f711565b905092915050565b60006020820190506201067660008301846200f78c565b92915050565b600060208201905081810360008301526201069881846200f7de565b905092915050565b6000608082019050620106b760008301876200f856565b620106c6602083018662010505565b620106d560408301856200f424565b620106e4606083018462010553565b95945050505050565b60006080820190506201070460008301876200f856565b62010713602083018662010505565b81810360408301526201072781856200f7de565b90506201073860608301846200f424565b95945050505050565b60006020820190506201075860008301846200f878565b92915050565b60006020820190506201077560008301846200f8ab565b92915050565b60006020820190506201079260008301846200f8bc565b92915050565b6000602082019050620107af60008301846200f8cd565b92915050565b6000602082019050620107cc60008301846200f8de565b92915050565b6000602082019050620107e960008301846200f8ef565b92915050565b60006020820190506201080660008301846200f900565b92915050565b60006020820190506201082360008301846200f911565b92915050565b60006020820190506201084060008301846200f922565b92915050565b60006020820190506201085d60008301846200f933565b92915050565b600060608201905081810360008301526201087f81866200f944565b905062010890602083018562010505565b6201089f604083018462010505565b949350505050565b60006020820190508181036000830152620108c2816200f9bc565b9050919050565b60006020820190508181036000830152620108e4816200f9e3565b9050919050565b600060208201905081810360008301526201090781846200fa3d565b905092915050565b600060608201905081810360008301526201092b81866200fa3d565b90506201093c60208301856200f78c565b6201094b60408301846200f78c565b949350505050565b600060208201905081810360008301526201096f81846200fc79565b905092915050565b600060e0820190506201098e60008301846200fd7b565b92915050565b6000606082019050620109ab60008301846200fea7565b92915050565b60006020820190508181036000830152620109cd81846200feef565b905092915050565b60006020820190508181036000830152620109f18184620100e1565b905092915050565b6000604082019050818103600083015262010a158185620100e1565b9050818103602083015262010a2b81846200fa3d565b90509392505050565b600060408201905062010a4b6000830184620102a7565b92915050565b60006101e08201905062010a696000830188620102da565b62010a7961016083018762010553565b62010a896101808301866200f89a565b62010a996101a083018562010553565b62010aa96101c083018462010553565b9695505050505050565b6000602082019050818103600083015262010acf8184620103d1565b905092915050565b6000602082019050818103600083015262010af3818462010427565b905092915050565b600062010b0762010b1a565b905062010b1582826201125f565b919050565b6000604051905090565b600067ffffffffffffffff82111562010b425762010b41620113d8565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562010b715762010b70620113d8565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562010ba05762010b9f620113d8565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562010bcf5762010bce620113d8565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562010bfe5762010bfd620113d8565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562010c2d5762010c2c620113d8565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562010c5c5762010c5b620113d8565b5b62010c678262011407565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600081905092915050565b600062010e6f82620110e8565b915062010e7c83620110e8565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111562010eb45762010eb36201134b565b5b828201905092915050565b600062010ecc8262011102565b915062010ed98362011102565b92508267ffffffffffffffff0382111562010ef95762010ef86201134b565b5b828201905092915050565b600062010f118262011102565b915062010f1e8362011102565b92508262010f315762010f306201137a565b5b828204905092915050565b600062010f498262011102565b915062010f568362011102565b92508167ffffffffffffffff048311821515161562010f7a5762010f796201134b565b5b828202905092915050565b600062010f9282620110e8565b915062010f9f83620110e8565b92508282101562010fb55762010fb46201134b565b5b828203905092915050565b600062010fcd8262011102565b915062010fda8362011102565b92508282101562010ff05762010fef6201134b565b5b828203905092915050565b60006201100882620110c8565b9050919050565b60008115159050919050565b6000819050919050565b6000620110328262010ffb565b9050919050565b6000620110468262010ffb565b9050919050565b60006201105a8262010ffb565b9050919050565b60006201106e8262010ffb565b9050919050565b6000620110828262010ffb565b9050919050565b600081905062011099826201149d565b919050565b6000819050620110ae82620114b4565b919050565b6000819050620110c382620114cb565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b600067ffffffffffffffff82169050919050565b6000620111238262011089565b9050919050565b600062011137826201109e565b9050919050565b60006201114b82620110b3565b9050919050565b60006201115f8262011102565b9050919050565b6000620111738262011102565b9050919050565b6000620111878262011102565b9050919050565b60006201119b8262011102565b9050919050565b6000620111af8262011102565b9050919050565b6000620111c38262011102565b9050919050565b6000620111d78262011102565b9050919050565b6000620111eb8262011102565b9050919050565b6000620111ff8262011102565b9050919050565b6000620112138262011102565b9050919050565b82818337600083830152505050565b60005b83811015620112495780820151818401526020810190506201122c565b8381111562011259576000848401525b50505050565b6201126a8262011407565b810181811067ffffffffffffffff821117156201128c576201128b620113d8565b5b80604052505050565b6000620112a282620110e8565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415620112d857620112d76201134b565b5b600182019050919050565b6000620112f08262011102565b915067ffffffffffffffff8214156201130e576201130d6201134b565b5b600182019050919050565b600062011326826201132d565b9050919050565b60006201133a8262011418565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f50756e697368466f72536563746f72206661696c656400000000000000000000600082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b600a8110620114b157620114b0620113a9565b5b50565b60038110620114c857620114c7620113a9565b5b50565b60028110620114df57620114de620113a9565b5b50565b620114ed8162010ffb565b8114620114f957600080fd5b50565b62011507816201100f565b81146201151357600080fd5b50565b62011521816201101b565b81146201152d57600080fd5b50565b6201153b8162011025565b81146201154757600080fd5b50565b620115558162011039565b81146201156157600080fd5b50565b6201156f816201104d565b81146201157b57600080fd5b50565b620115898162011061565b81146201159557600080fd5b50565b620115a38162011075565b8114620115af57600080fd5b50565b60038110620115c057600080fd5b50565b60028110620115d157600080fd5b50565b620115df81620110e8565b8114620115eb57600080fd5b50565b620115f981620110f2565b81146201160557600080fd5b50565b620116138162011102565b81146201161f57600080fd5b5056fea2646970667358221220de8f278001ef67ebb607491cad05c32b73a0aa0c0e2cbb0f9289396040aa93b364736f6c63430008040033",
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

// GetPocProveList is a free data retrieval call binding the contract method 0x0edec0f2.
//
// Solidity: function GetPocProveList() view returns((address,uint256,uint64)[])
func (_Store *StoreCaller) GetPocProveList(opts *bind.CallOpts) ([]PocProve, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetPocProveList")

	if err != nil {
		return *new([]PocProve), err
	}

	out0 := *abi.ConvertType(out[0], new([]PocProve)).(*[]PocProve)

	return out0, err

}

// GetPocProveList is a free data retrieval call binding the contract method 0x0edec0f2.
//
// Solidity: function GetPocProveList() view returns((address,uint256,uint64)[])
func (_Store *StoreSession) GetPocProveList() ([]PocProve, error) {
	return _Store.Contract.GetPocProveList(&_Store.CallOpts)
}

// GetPocProveList is a free data retrieval call binding the contract method 0x0edec0f2.
//
// Solidity: function GetPocProveList() view returns((address,uint256,uint64)[])
func (_Store *StoreCallerSession) GetPocProveList() ([]PocProve, error) {
	return _Store.Contract.GetPocProveList(&_Store.CallOpts)
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

// GetPocProve is a free data retrieval call binding the contract method 0x9908f2bf.
//
// Solidity: function getPocProve(address nodeAddr, uint256 height) view returns((address,uint256,uint64))
func (_Store *StoreCaller) GetPocProve(opts *bind.CallOpts, nodeAddr common.Address, height *big.Int) (PocProve, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getPocProve", nodeAddr, height)

	if err != nil {
		return *new(PocProve), err
	}

	out0 := *abi.ConvertType(out[0], new(PocProve)).(*PocProve)

	return out0, err

}

// GetPocProve is a free data retrieval call binding the contract method 0x9908f2bf.
//
// Solidity: function getPocProve(address nodeAddr, uint256 height) view returns((address,uint256,uint64))
func (_Store *StoreSession) GetPocProve(nodeAddr common.Address, height *big.Int) (PocProve, error) {
	return _Store.Contract.GetPocProve(&_Store.CallOpts, nodeAddr, height)
}

// GetPocProve is a free data retrieval call binding the contract method 0x9908f2bf.
//
// Solidity: function getPocProve(address nodeAddr, uint256 height) view returns((address,uint256,uint64))
func (_Store *StoreCallerSession) GetPocProve(nodeAddr common.Address, height *big.Int) (PocProve, error) {
	return _Store.Contract.GetPocProve(&_Store.CallOpts, nodeAddr, height)
}

// CheckNodeSectorProvedInTime is a paid mutator transaction binding the contract method 0x648d225d.
//
// Solidity: function CheckNodeSectorProvedInTime((address,uint64) sectorRef) returns()
func (_Store *StoreTransactor) CheckNodeSectorProvedInTime(opts *bind.TransactOpts, sectorRef SectorRef) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "CheckNodeSectorProvedInTime", sectorRef)
}

// CheckNodeSectorProvedInTime is a paid mutator transaction binding the contract method 0x648d225d.
//
// Solidity: function CheckNodeSectorProvedInTime((address,uint64) sectorRef) returns()
func (_Store *StoreSession) CheckNodeSectorProvedInTime(sectorRef SectorRef) (*types.Transaction, error) {
	return _Store.Contract.CheckNodeSectorProvedInTime(&_Store.TransactOpts, sectorRef)
}

// CheckNodeSectorProvedInTime is a paid mutator transaction binding the contract method 0x648d225d.
//
// Solidity: function CheckNodeSectorProvedInTime((address,uint64) sectorRef) returns()
func (_Store *StoreTransactorSession) CheckNodeSectorProvedInTime(sectorRef SectorRef) (*types.Transaction, error) {
	return _Store.Contract.CheckNodeSectorProvedInTime(&_Store.TransactOpts, sectorRef)
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

// FileProve is a paid mutator transaction binding the contract method 0x8e270530.
//
// Solidity: function FileProve((bytes,bytes,uint256,address,uint64,uint64) fileProve) returns()
func (_Store *StoreTransactor) FileProve(opts *bind.TransactOpts, fileProve FileProveParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "FileProve", fileProve)
}

// FileProve is a paid mutator transaction binding the contract method 0x8e270530.
//
// Solidity: function FileProve((bytes,bytes,uint256,address,uint64,uint64) fileProve) returns()
func (_Store *StoreSession) FileProve(fileProve FileProveParams) (*types.Transaction, error) {
	return _Store.Contract.FileProve(&_Store.TransactOpts, fileProve)
}

// FileProve is a paid mutator transaction binding the contract method 0x8e270530.
//
// Solidity: function FileProve((bytes,bytes,uint256,address,uint64,uint64) fileProve) returns()
func (_Store *StoreTransactorSession) FileProve(fileProve FileProveParams) (*types.Transaction, error) {
	return _Store.Contract.FileProve(&_Store.TransactOpts, fileProve)
}

// SectorProve is a paid mutator transaction binding the contract method 0x09cbe391.
//
// Solidity: function SectorProve((address,uint64,uint64,bytes) sectorProve) returns()
func (_Store *StoreTransactor) SectorProve(opts *bind.TransactOpts, sectorProve SectorProveParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SectorProve", sectorProve)
}

// SectorProve is a paid mutator transaction binding the contract method 0x09cbe391.
//
// Solidity: function SectorProve((address,uint64,uint64,bytes) sectorProve) returns()
func (_Store *StoreSession) SectorProve(sectorProve SectorProveParams) (*types.Transaction, error) {
	return _Store.Contract.SectorProve(&_Store.TransactOpts, sectorProve)
}

// SectorProve is a paid mutator transaction binding the contract method 0x09cbe391.
//
// Solidity: function SectorProve((address,uint64,uint64,bytes) sectorProve) returns()
func (_Store *StoreTransactorSession) SectorProve(sectorProve SectorProveParams) (*types.Transaction, error) {
	return _Store.Contract.SectorProve(&_Store.TransactOpts, sectorProve)
}

// SetLastPunishmentHeightForNode is a paid mutator transaction binding the contract method 0x3ec0f5df.
//
// Solidity: function SetLastPunishmentHeightForNode(address nodeAddr, uint64 sectorId, uint256 height) returns()
func (_Store *StoreTransactor) SetLastPunishmentHeightForNode(opts *bind.TransactOpts, nodeAddr common.Address, sectorId uint64, height *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SetLastPunishmentHeightForNode", nodeAddr, sectorId, height)
}

// SetLastPunishmentHeightForNode is a paid mutator transaction binding the contract method 0x3ec0f5df.
//
// Solidity: function SetLastPunishmentHeightForNode(address nodeAddr, uint64 sectorId, uint256 height) returns()
func (_Store *StoreSession) SetLastPunishmentHeightForNode(nodeAddr common.Address, sectorId uint64, height *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetLastPunishmentHeightForNode(&_Store.TransactOpts, nodeAddr, sectorId, height)
}

// SetLastPunishmentHeightForNode is a paid mutator transaction binding the contract method 0x3ec0f5df.
//
// Solidity: function SetLastPunishmentHeightForNode(address nodeAddr, uint64 sectorId, uint256 height) returns()
func (_Store *StoreTransactorSession) SetLastPunishmentHeightForNode(nodeAddr common.Address, sectorId uint64, height *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetLastPunishmentHeightForNode(&_Store.TransactOpts, nodeAddr, sectorId, height)
}

// SettleForFile is a paid mutator transaction binding the contract method 0x7b85255a.
//
// Solidity: function SettleForFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo, (uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo, (address,address,uint64,uint256,bool) detail, (address,address,uint64,uint256,bool)[] details, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting) returns()
func (_Store *StoreTransactor) SettleForFile(opts *bind.TransactOpts, fileInfo FileInfo, nodeInfo NodeInfo, detail ProveDetail, details []ProveDetail, setting Setting) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SettleForFile", fileInfo, nodeInfo, detail, details, setting)
}

// SettleForFile is a paid mutator transaction binding the contract method 0x7b85255a.
//
// Solidity: function SettleForFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo, (uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo, (address,address,uint64,uint256,bool) detail, (address,address,uint64,uint256,bool)[] details, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting) returns()
func (_Store *StoreSession) SettleForFile(fileInfo FileInfo, nodeInfo NodeInfo, detail ProveDetail, details []ProveDetail, setting Setting) (*types.Transaction, error) {
	return _Store.Contract.SettleForFile(&_Store.TransactOpts, fileInfo, nodeInfo, detail, details, setting)
}

// SettleForFile is a paid mutator transaction binding the contract method 0x7b85255a.
//
// Solidity: function SettleForFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)) fileInfo, (uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo, (address,address,uint64,uint256,bool) detail, (address,address,uint64,uint256,bool)[] details, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting) returns()
func (_Store *StoreTransactorSession) SettleForFile(fileInfo FileInfo, nodeInfo NodeInfo, detail ProveDetail, details []ProveDetail, setting Setting) (*types.Transaction, error) {
	return _Store.Contract.SettleForFile(&_Store.TransactOpts, fileInfo, nodeInfo, detail, details, setting)
}

// UpdateProveDetailList is a paid mutator transaction binding the contract method 0x27ab4434.
//
// Solidity: function UpdateProveDetailList(bytes fileHash, (address,address,uint64,uint256,bool)[] details) returns()
func (_Store *StoreTransactor) UpdateProveDetailList(opts *bind.TransactOpts, fileHash []byte, details []ProveDetail) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateProveDetailList", fileHash, details)
}

// UpdateProveDetailList is a paid mutator transaction binding the contract method 0x27ab4434.
//
// Solidity: function UpdateProveDetailList(bytes fileHash, (address,address,uint64,uint256,bool)[] details) returns()
func (_Store *StoreSession) UpdateProveDetailList(fileHash []byte, details []ProveDetail) (*types.Transaction, error) {
	return _Store.Contract.UpdateProveDetailList(&_Store.TransactOpts, fileHash, details)
}

// UpdateProveDetailList is a paid mutator transaction binding the contract method 0x27ab4434.
//
// Solidity: function UpdateProveDetailList(bytes fileHash, (address,address,uint64,uint256,bool)[] details) returns()
func (_Store *StoreTransactorSession) UpdateProveDetailList(fileHash []byte, details []ProveDetail) (*types.Transaction, error) {
	return _Store.Contract.UpdateProveDetailList(&_Store.TransactOpts, fileHash, details)
}

// UpdateProveDetailMeta is a paid mutator transaction binding the contract method 0xbb4e4e42.
//
// Solidity: function UpdateProveDetailMeta(bytes fileHash, (uint64,uint64) details) returns()
func (_Store *StoreTransactor) UpdateProveDetailMeta(opts *bind.TransactOpts, fileHash []byte, details ProveDetailMeta) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateProveDetailMeta", fileHash, details)
}

// UpdateProveDetailMeta is a paid mutator transaction binding the contract method 0xbb4e4e42.
//
// Solidity: function UpdateProveDetailMeta(bytes fileHash, (uint64,uint64) details) returns()
func (_Store *StoreSession) UpdateProveDetailMeta(fileHash []byte, details ProveDetailMeta) (*types.Transaction, error) {
	return _Store.Contract.UpdateProveDetailMeta(&_Store.TransactOpts, fileHash, details)
}

// UpdateProveDetailMeta is a paid mutator transaction binding the contract method 0xbb4e4e42.
//
// Solidity: function UpdateProveDetailMeta(bytes fileHash, (uint64,uint64) details) returns()
func (_Store *StoreTransactorSession) UpdateProveDetailMeta(fileHash []byte, details ProveDetailMeta) (*types.Transaction, error) {
	return _Store.Contract.UpdateProveDetailMeta(&_Store.TransactOpts, fileHash, details)
}

// Initialize is a paid mutator transaction binding the contract method 0xf76031f6.
//
// Solidity: function initialize(address _config, address _fs, address _node, address _pdp, address _sector, (uint64) proveConfig) returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts, _config common.Address, _fs common.Address, _node common.Address, _pdp common.Address, _sector common.Address, proveConfig ProveConfig) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize", _config, _fs, _node, _pdp, _sector, proveConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0xf76031f6.
//
// Solidity: function initialize(address _config, address _fs, address _node, address _pdp, address _sector, (uint64) proveConfig) returns()
func (_Store *StoreSession) Initialize(_config common.Address, _fs common.Address, _node common.Address, _pdp common.Address, _sector common.Address, proveConfig ProveConfig) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _fs, _node, _pdp, _sector, proveConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0xf76031f6.
//
// Solidity: function initialize(address _config, address _fs, address _node, address _pdp, address _sector, (uint64) proveConfig) returns()
func (_Store *StoreTransactorSession) Initialize(_config common.Address, _fs common.Address, _node common.Address, _pdp common.Address, _sector common.Address, proveConfig ProveConfig) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _fs, _node, _pdp, _sector, proveConfig)
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

// PutPocProve is a paid mutator transaction binding the contract method 0x1581d545.
//
// Solidity: function putPocProve((address,uint256,uint64) prove) returns()
func (_Store *StoreTransactor) PutPocProve(opts *bind.TransactOpts, prove PocProve) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "putPocProve", prove)
}

// PutPocProve is a paid mutator transaction binding the contract method 0x1581d545.
//
// Solidity: function putPocProve((address,uint256,uint64) prove) returns()
func (_Store *StoreSession) PutPocProve(prove PocProve) (*types.Transaction, error) {
	return _Store.Contract.PutPocProve(&_Store.TransactOpts, prove)
}

// PutPocProve is a paid mutator transaction binding the contract method 0x1581d545.
//
// Solidity: function putPocProve((address,uint256,uint64) prove) returns()
func (_Store *StoreTransactorSession) PutPocProve(prove PocProve) (*types.Transaction, error) {
	return _Store.Contract.PutPocProve(&_Store.TransactOpts, prove)
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
