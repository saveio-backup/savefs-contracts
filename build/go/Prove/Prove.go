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

// ProveFileProveParams is an auto generated low-level Go binding around an user-defined struct.
type ProveFileProveParams struct {
	FileHash    []byte
	ProveData   []byte
	BlockHeight *big.Int
	NodeWallet  common.Address
	Profit      uint64
	SectorID    uint64
}

// ProveSectorProveParams is an auto generated low-level Go binding around an user-defined struct.
type ProveSectorProveParams struct {
	NodeAddr        common.Address
	SectorID        uint64
	ChallengeHeight uint64
	ProveData       []byte
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"FileProveFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileProveNotFileOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeSectorProvedInTimeError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SectorProveFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"CheckNodeSectorProvedInTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteProveDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"NodeWallet\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"}],\"internalType\":\"structProve.FileProveParams\",\"name\":\"fileProve\",\"type\":\"tuple\"}],\"name\":\"FileProve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetPocProveList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"Height\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"PlotSize\",\"type\":\"uint64\"}],\"internalType\":\"structPocProve[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"GetProveDetailList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ChallengeHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"}],\"internalType\":\"structProve.SectorProveParams\",\"name\":\"sectorProve\",\"type\":\"tuple\"}],\"name\":\"SectorProve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"SetLastPunishmentHeightForNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail\",\"name\":\"detail\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"details\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"}],\"name\":\"SettleForFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"details\",\"type\":\"tuple[]\"}],\"name\":\"UpdateProveDetailList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveDetailNum\",\"type\":\"uint64\"}],\"internalType\":\"structProveDetailMeta\",\"name\":\"details\",\"type\":\"tuple\"}],\"name\":\"UpdateProveDetailMeta\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0xd11e680a\",\"type\":\"bytes32\"}],\"name\":\"c_0xd11e680a\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ChallengeHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"}],\"internalType\":\"structProve.SectorProveParams\",\"name\":\"sectorProve\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"}],\"name\":\"checkSectorProve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"getPocProve\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"Height\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"PlotSize\",\"type\":\"uint64\"}],\"internalType\":\"structPocProve\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractFileSystem\",\"name\":\"_fs\",\"type\":\"address\"},{\"internalType\":\"contractNode\",\"name\":\"_node\",\"type\":\"address\"},{\"internalType\":\"contractPDP\",\"name\":\"_pdp\",\"type\":\"address\"},{\"internalType\":\"contractSector\",\"name\":\"_sector\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"times\",\"type\":\"uint64\"}],\"name\":\"punishForSector\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"Height\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"PlotSize\",\"type\":\"uint64\"}],\"internalType\":\"structPocProve\",\"name\":\"prove\",\"type\":\"tuple\"}],\"name\":\"putPocProve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526020600060026101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055503480156200003b57600080fd5b50620114f8806200004d6000396000f3fe6080604052600436106200010a5760003560e01c8063648d225d11620000975780639908f2bf11620000615780639908f2bf1462000365578063bb4e4e4214620003a9578063ef5d51b414620003d7578063fda593511462000405576200010a565b8063648d225d14620002975780638e27053014620002c55780639696161414620002f3578063977fdfe21462000321576200010a565b80631581d54511620000d95780631581d54514620001df57806327ab4434146200020d5780633ec0f5df146200023b57806352e061461462000269576200010a565b806309cbe391146200010f5780630edec0f2146200013d5780630fece869146200016d5780631459457a14620001b1575b600080fd5b3480156200011c57600080fd5b506200013b60048036038101906200013591906200f0ea565b62000425565b005b3480156200014a57600080fd5b5062000155620014ad565b604051620001649190620104b7565b60405180910390f35b3480156200017a57600080fd5b506200019960048036038101906200019391906200f12f565b620014e0565b604051620001a89190620104ff565b60405180910390f35b348015620001be57600080fd5b50620001dd6004803603810190620001d791906200edbc565b62002a0a565b005b348015620001ec57600080fd5b506200020b60048036038101906200020591906200eff1565b62002e3b565b005b3480156200021a57600080fd5b506200023960048036038101906200023391906200ecef565b62002ff9565b005b3480156200024857600080fd5b506200026760048036038101906200026191906200ebb7565b62003247565b005b3480156200027657600080fd5b506200029560048036038101906200028f91906200ecaa565b6200333f565b005b348015620002a457600080fd5b50620002c36004803603810190620002bd91906200f1a2565b62003462565b005b348015620002d257600080fd5b50620002f16004803603810190620002eb91906200ef3b565b6200403e565b005b3480156200030057600080fd5b506200031f60048036038101906200031991906200ec7e565b620073b9565b005b3480156200032e57600080fd5b506200034d60048036038101906200034791906200ecaa565b620073bc565b6040516200035c9190620104db565b60405180910390f35b3480156200037257600080fd5b506200039160048036038101906200038b91906200eb76565b620078a1565b604051620003a0919062010834565b60405180910390f35b348015620003b657600080fd5b50620003d56004803603810190620003cf91906200ed62565b62007a7f565b005b348015620003e457600080fd5b50620004036004803603810190620003fd91906200ee83565b62007b8f565b005b6200042360048036038101906200041d91906200f062565b620089d3565b005b620004537f5ff17198d2f86688937e2d66db6d6961886360239ca70dceccbcf7a31a55f38760001b620073b9565b620004817fdfe3d37a596e1c5737f7b28325f95f4d5d809b8c54e87960e1a3ff0d969b2dbb60001b620073b9565b620004af7f39f2f0c80339f8a20e89e87f7fbb315f2b58c353bcc2c19a89fe49659f291ee760001b620073b9565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633778febe83600001516040518263ffffffff1660e01b815260040162000512919062010466565b60e06040518083038186803b1580156200052b57600080fd5b505afa15801562000540573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200056691906200ef80565b9050620005967f5165ee976b9836bb433900cc23c73645533d1b32298a133155b63d852b6adf5260001b620073b9565b620005c47f75386f3147cbfddd7db9fb0aab2963b04d024f06705cbb808d610a48b5fd704560001b620073b9565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d76040518060400160405280866000015173ffffffffffffffffffffffffffffffffffffffff168152602001866020015167ffffffffffffffff168152506040518263ffffffff1660e01b81526004016200065f9190620108d4565b60006040518083038186803b1580156200067857600080fd5b505afa1580156200068d573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190620006b891906200f01d565b9050620006e87f4b1e5305eef31cea3532f31ad5c5cd10b4d8a200aa0ee20349d6142e6c47ef5360001b620073b9565b620007167f7544975bfe18be3fdb57bee43c9d6234c3daa5d6ac38ef8e45ac31b5198a5d9960001b620073b9565b600080600a9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156200078157600080fd5b505afa15801562000796573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620007bc91906200f1ce565b9050620007ec7fcd7f3eaa18eba7044f49e228f851d574c0d29205ad3b998a960153c00da7a9ad60001b620073b9565b6200081a7fb64a679f7f9fe8eab2a03c8dfd087abe150a5141bce2b4a58102856b669eaaf860001b620073b9565b8160c00151431015620008c357620008557fcd57a7f1ffedcb2b9253253e6e9e7bb214c3dcdf1d92f558f6b50734fc6505fb60001b620073b9565b620008837fbe8ffdc8997715b0254950b4369f5e0e8bd5984036546376c08449e966e6157760001b620073b9565b60016040517fcab88a3c000000000000000000000000000000000000000000000000000000008152600401620008ba9190620105fe565b60405180910390fd5b620008f17f645769f1b21d4064fc12919ec2fb385d13645ce2558dd90b6fe5a3dfae25d5b960001b620073b9565b6200091f7fecaaa99bd0de2f39619fdee68ee6b4c349ad5ab026aff1441c6ea5e3198df36160001b620073b9565b6200094d7fd85ab96555044395a91bf12f8bf847ca16406672c7c481515563208519e1610b60001b620073b9565b8160c00151846040015167ffffffffffffffff161462000ab357620009957f6a938f362bd7d8257a50727c623165992f1b0524862b64ed2ee9f45384d35c4b60001b620073b9565b620009c37f591f24fe2a938e4042c98fcfd965064c961d18786652495ae6ac333395c80e8b60001b620073b9565b620009f17fdacca2a8e2ef852f9e44a5f795207510b2b71301d7e5aaf300d92ca1cc9facb460001b620073b9565b62000a456040518060400160405280601681526020017f4368616c6c656e6765486569676874206572726f723a00000000000000000000815250856040015167ffffffffffffffff168460c0015162009006565b62000a737fe0d5637626943fd49d96fd5bcb460320535009f789ae5577712a2aa96717909460001b620073b9565b60026040517fcab88a3c00000000000000000000000000000000000000000000000000000000815260040162000aaa91906201061b565b60405180910390fd5b62000ae17f1b996e2a735db9f17042d4545f82b10ea4e790b7b452dea00dc75788cf97d60560001b620073b9565b62000b0f7f97b48fc267e77fccb7dc43ecad070e29871320dc1ee055020e32b6c765e2b90460001b620073b9565b62000b3d7fcb9f9495d337eafb10bef4c77d80e567fea2e183ec666acde4c5d168d8d7d0f260001b620073b9565b600062000b4b8584620014e0565b905062000b7b7ff0c32d95a6c03a442955875459a83e98a0350059d8a63ecf9801ba50405292c760001b620073b9565b62000ba97f54e1364924f69afe3c273694aaac68a5e40081dfa2e509782de53e39ad0f098d60001b620073b9565b8062000cb65762000bdd7fbacbf44b24f6cd41e00cc77f0314eca6ebce1d84ad336dc47539c9c4846f924860001b620073b9565b62000c0b7f554c4a82174e6be72a24357c25c018c08e3d314583363cf44b5281f8d7f53ea160001b620073b9565b62000c397f9f2beb20d88aedd2d4097a2ca20d44a8dda8df6db34ba2acae9ed5263c30f35460001b620073b9565b62000c488385846001620089d3565b62000c767f1d9e55adcbced2e9f5cac09b59669946deaca4f4a564e6f589e042ffabaed4ba60001b620073b9565b60036040517fcab88a3c00000000000000000000000000000000000000000000000000000000815260040162000cad919062010638565b60405180910390fd5b62000ce47f42c5d7b397ea59074ef3017405102011900008d90d8cd9f4f2516df583a76ada60001b620073b9565b62000d127f3d97537de75785d84f26099d104b642fdf91e296b9a0e768e2b70e9e197775ed60001b620073b9565b62000d407f16ccb16e4a41378b9ed43c80b0950bc4a8f402d2a8c271fe0023792182edea3260001b620073b9565b600062000d4f848685620090a9565b905062000d7f7f1c466f470101914b7e7bc84279416ebfdb0ab50749d0e4fdb910a00d4a02938560001b620073b9565b62000dad7f99d4b32a0fcf8eb9e87f21d232efc4066a946d986e48a5890d9d46aed34c6f2060001b620073b9565b8062000e4f5762000de17fea609f7d3df430484a145960f1e9b9da1b091f32b5409566f433cdcd51c401ca60001b620073b9565b62000e0f7f5c90e46ad93e0975126af0d56eec156288b5d96d07c049d32bb1097475633ae960001b620073b9565b60056040517fcab88a3c00000000000000000000000000000000000000000000000000000000815260040162000e46919062010672565b60405180910390fd5b62000e7d7f1305f6ad2c8ed618af6dd9524578e0559d2b815ca252a89d3004dc3635bff02c60001b620073b9565b62000eab7fb5e3f4b34a63fc3d95607c604b0dfa68a2486479be7e63fbaf7686eb15230a8460001b620073b9565b62000ed97fb77c6d79f5d1109995f72a7e3c86b8951b9b4faa8270563931c106f6dce1797560001b620073b9565b60008460a00151141562000f815762000f157f9177a2d82a45530837b2acf20453430f2e90a7e78d41f405a9bfaa84d7f6134260001b620073b9565b62000f437fe535de601ad48db66f0ef4ec8e5bba7f74bc5d082a4e6980feff9ff4416d6ba660001b620073b9565b62000f717fc12372f42e1d5ff8768d8c4dc009632e8022f979ea30eb702760d054b58a135160001b620073b9565b438460a001818152505062000fb0565b62000faf7f05f9c76f7747b746e60dae86b9db91d0e48f111cca695d3b1b0bb325ac71b1d960001b620073b9565b5b62000fde7f4fb8f8c4788dbcc5195d53a1cf7d202c648a7c664efe112c305bbf49a34d63db60001b620073b9565b6200100c7fe693f4fb98ddf07777f2eb3ee2b0900295190f930c35d4afd2dcf2a7d853a6f860001b620073b9565b8260c0015167ffffffffffffffff164362001028919062010d02565b8460c00181815250506200105f7f4712799cc8cd2c0818ab0e574c6c30dfe0c62331140eb677074c0ab50abdf37f60001b620073b9565b6200108d7f65273b12e8392db757daa3be4b80cf456a07f17b5852cd0c5c9caf54d24b016d60001b620073b9565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632384a6aa856040518263ffffffff1660e01b8152600401620010ea919062010875565b600060405180830381600087803b1580156200110557600080fd5b505af11580156200111a573d6000803e3d6000fd5b505050506200114c7f1afa166f9206d7acadf80b68c76817c6dfecb49889b0acb41f3af875cb4b998860001b620073b9565b6200117a7fee0e95267b4bdc677e4d2c30e3e537c8e689ce71ed410b933ab45ae6b711494760001b620073b9565b8361014001516200122157620011b37f8d10afe56b2353d3ac50c9d3c8b29015d39720599e20a3b9ce88f4b05211add860001b620073b9565b620011e17f52cb81c196e79141f8a2b997ca09889be0f2dba0c56a64acdab301fe23f09e0360001b620073b9565b60046040517fcab88a3c00000000000000000000000000000000000000000000000000000000815260040162001218919062010655565b60405180910390fd5b6200124f7fa412963f5783fa46f92181bbabd06f13eb15132070bf72fce0bb67f1c6952c2560001b620073b9565b6200127d7fda5fbb2deab399ff1e7b5122b29bee87ab0bbb19cd577f2851b7c52f9e86f32660001b620073b9565b620012ab7faed76bd3e02d3e22125b5cee0cd2191d12a0e7fcbe6a77917dada412b5d9a8f260001b620073b9565b6000620012bd856000015143620078a1565b9050620012ed7f1b9502f8d88523498966d13f0259b917ceda06de15b97ab776dd38907c24f95e60001b620073b9565b6200131b7fad945ca0cbd2170b6491a8366ccc72e40ca53836cd11a83a961529f0afdfdec360001b620073b9565b43816020018181525050620013537f920c38f36b3f5f66a72a427d85e95ad565a4aa37994caf321ca876830b8d0ec260001b620073b9565b620013817f9f7fb924ec4e72fe7a63d885dba9bf26c0f5c4c0e3e9516a438a654e68a56e4f60001b620073b9565b8460000151816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050620013eb7f58d90a4a575b4f93c784394f5ce5e9e4c3a7f8af46ae87daa9ec7f93f284e23260001b620073b9565b620014197f23fd89998caa391ffbd0d023973356e4f6441648835a2c52e11baf42fb0e646960001b620073b9565b8460600151816040019067ffffffffffffffff16908167ffffffffffffffff16815250506200146b7fe1818bb37b1b0aef35ce99c07992a868f30411ee6a552a27b2baf24097b996ff60001b620073b9565b620014997fcf68ae11dd5355f2eb71939e3c9ea309194896ab10c93fcc31fe3a98c022433e60001b620073b9565b620014a48162002e3b565b50505050505050565b6060620014dd7f9ffc4a4bb782793ddc737bfabf881edea5554217ba73d2a8531ad959e01b6dbf60001b620073b9565b90565b6000620015107fdacc2bd7ce38720b67d183371caaa0f24e60a62bc21edb631d516c4aa6d02cf560001b620073b9565b6200153e7fd89f9d94a68689eae8e69ac3c4f15de1b8d9e8a56021d0e20065d5ec7c6f4ce660001b620073b9565b6200156c7fa09e05b3cf073ed2effb9d30e2087b389c12ed23e4405723e38bea1bb5ef425a60001b620073b9565b60606200159c7f3af92384b91affa8ca1013ed8d6384d9a4976c568badca5950099c54a67ad3d360001b620073b9565b620015ca7f12e7f9e03e2a39ee339815411dc12c5fca5ffa3fa0e043598555658db050000960001b620073b9565b620015d46200cb06565b620016027ffee56e7255445e9958b6ad3a045e09cb3fb51e59de0c24565fdb858e896017a160001b620073b9565b620016307f115ad34b2d05efbb8921a2f385d96f45bf0c89fb8605958e61c05fc7f30ccbdd60001b620073b9565b8460000151816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506200169a7f2f082e63060b7670254433dd60679cd1280710eb974f7d3c9532e0539ad5a08460001b620073b9565b620016c87f6aa3e9023062aa0c35eefed3c7af2ea727867554a36abbccdeda3d71f094bc8460001b620073b9565b818160200181905250620016ff7f8f02e60abfc8d672123b33dcd59e26cb5397af2ca7f90fe16b5fc7cfd48e628660001b620073b9565b6200172d7f5d84eae77ba7510ea5f393a73b9a1b0b7cb06bc1f532675b3343d859a78657e760001b620073b9565b8360e00151816040019067ffffffffffffffff16908167ffffffffffffffff16815250506200177f7f83899ccfcf976e2818ec29708a2771fa2e985a297bde71c7b81a534fce97fe8160001b620073b9565b620017ad7f9a8541901b452bc70728d98109b1e82273ad10f010480177cb67092ceaee87af60001b620073b9565b600060029054906101000a900467ffffffffffffffff16816060019067ffffffffffffffff16908167ffffffffffffffff1681525050620018117f09250155911b4c74ff4431abe4bd8cd822933e21869b24fdc15577660f627eaa60001b620073b9565b6200183f7fc08591f3432c3a62392ee2fd31ed91e89aab59c0bd7c39058b27b35a64b51eb060001b620073b9565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fd7c9808836040518263ffffffff1660e01b81526004016200189e9190620107f3565b60006040518083038186803b158015620018b757600080fd5b505afa158015620018cc573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190620018f791906200ec0d565b9050620019277f376874d940ba3b3327d1c15313f9a3a69540f30235a4c730951e9a8916c1064060001b620073b9565b620019557fa3927e45c75163765bfa0dce1318a3859e2be6f2c4ad9f213efe7c9e1513625360001b620073b9565b6200195f6200cb58565b6200198d7f859a62f32495d1fd93d54d1598f5467ccf90a0d4745fdb590c9f61ec71043c7360001b620073b9565b620019bb7fdce8d1da1d99fd386ab167fa2b785a3e5206f2a98ebfe8057d868fe06b7da9d360001b620073b9565b620019c56200cba2565b620019f37f4c9eee751618169de62f50e7100004ec3ca864fc2c07fa4569082fedc4b2fcef60001b620073b9565b62001a217f11b9140c3e27ef65e382d373090cedaf5641d10ac6041a42fffc8f275de5f32a60001b620073b9565b86816000018190525062001a587fb9749b1b62eb9df4cc4d5db5561bbed52b29267984ef91238da69940dae0a80a60001b620073b9565b62001a867f5269af5f5b22ef5fccdd7a87dd08ee52ef1aaa2f4dd88343084786fb1d94315760001b620073b9565b81816040018190525062001abd7feccc78ac89a1f98f7e2d7258c43ee9a828bf86bdd305ee59dcb785fed0738b7060001b620073b9565b62001aeb7f454e0214d681b4b63f724d326e0d0661391c13735c8c679578f4c07df072be0560001b620073b9565b62001af56200cbd3565b62001b237f9f9b420cc110d1fd67b11a835b85e3b7cc1b6eb3aacec91a13c13cffbbec32bf60001b620073b9565b62001b517f723824323ebac42a4bcc38fd4bead264152f828c68c4a13463b8ac0c79e6d5c460001b620073b9565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639f9ca644836040518263ffffffff1660e01b815260040162001bae919062010851565b60006040518083038186803b15801562001bc757600080fd5b505afa15801562001bdc573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019062001c0791906200efac565b905062001c377feedbd47c367f9517f3eb649e940603c7aaf532ed8668a0c65e087e949531d2dc60001b620073b9565b62001c657f2958981ffff30cd568bf95f7d1ca2c4b2c1b6fde5edbce4b332cf1d73e820e4b60001b620073b9565b8060c0015162001d095762001c9d7f31cf12c99ed3d404486649dffddab0c5edbaf5d7b0cfcc5e3cbe65768dc120eb60001b620073b9565b62001ccb7f8b61f31a21fa507d407d3d6ddf1a19d475d3ec4d06e0a0f189cac7116189759f60001b620073b9565b62001cf97fd256f4c73beb4f11ac24f614165c4af0800927c2305646647ac691abd475130860001b620073b9565b6000965050505050505062002a04565b62001d377f66fbd2a524fb997c3e82cdf94e21f92099225be9ec97b1acb720cb70fa73062260001b620073b9565b62001d657fb2c6872138be6dbf223b08fedfe35a0ee33e775c0723f4f5e54d5da01abc99a160001b620073b9565b62001d937ff02e33d5bd52057e47723bbba4cebd153d29043cd2a806046a64a5277038b85c60001b620073b9565b62001d9d6200cc1a565b62001dcb7fd49fcf361432ebf1464a00416ba681cc70ea122828b75352bd267fbcf752162860001b620073b9565b62001df97fc7a94e088072a7ea4f206a8b5d25a5fd8f54f3eb9c6384b22b87d4884860720b60001b620073b9565b6000816000019067ffffffffffffffff16908167ffffffffffffffff168152505062001e487f5bb8fb295deeb9747ba42e5b8f61cb12a270563f7533c4dadddbd091d6b5379260001b620073b9565b62001e767fbe878fe0d7fbc9144ed479938f48fbc5e4816f2b3710c56f01f850ba3afe9d4d60001b620073b9565b8360400151816020018190525062001eb17fabbf519eeb13113fd93651e1a8e4ea6c317ada15cb7c019ee4ad2f8314476d7260001b620073b9565b62001edf7fb84fffaacf916ea99b2243c7375633655735b3d17eff16b5fe2dac58536097c160001b620073b9565b8160000151816040018190525062001f1a7ff03fa4ef51b4bae369da51a08838fc3f826c12d982adb6499887f8143d76657760001b620073b9565b62001f487f354a94e6eda2abc178686d2918ba9c31aa02b5d03b2a287a9a6dee230152878a60001b620073b9565b8160200151816060018190525062001f837f5744ba7f8c7b4d876fb4689dabce586c6b6fef0f9d26a79d1a5d5b08adac96d560001b620073b9565b62001fb17fe130323275832d315616de68598c18ed6050c94e4f115bf999210f859b379ef660001b620073b9565b8160400151816080018190525062001fec7fe41f63410fc8503ed943f0185754934b8e63e8631eac1173bd4ecde15898875d60001b620073b9565b6200201a7fe1ebd55178cddbfd96e1ef8557d5da40f14d46df7da01e5a30b3f6096efd1d8c60001b620073b9565b81606001518160a00181905250620020557f0505be9033dce48c9d5e8375e60a73f9dd9e45437f681933002bd024137f9cdb60001b620073b9565b620020837f8c1f0d7c0b08480def4b8fdd0cef311dd7d1c845a2a14f05d673835cc1e2558c60001b620073b9565b81608001518160c00181905250620020be7fa703b1219893c972323e8d2bb20dda8925670fd58b18243ad8c778d147043d0a60001b620073b9565b620020ec7f38a6f9c257f256361a195953524f12720c70caff12dde71d7e20bfe6d84c1bef60001b620073b9565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b6ddeca0836040518263ffffffff1660e01b81526004016200214b919062010977565b60206040518083038186803b1580156200216457600080fd5b505afa15801562002179573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200219f91906200ec52565b9050620021cf7fede461ac5e07ef51deae8fefde97fab96c4518b69c8e0cb87df02ef4bfce36fc60001b620073b9565b620021fd7feac4e28cbd209b0449c70e3654f4599cb45bb8054f16def61024a8a870c3961560001b620073b9565b806200229f57620022317ff9bbb499be44c6918fec029501af4ead958731569dce56b4f20cf142ec172c8f60001b620073b9565b6200225f7f54e21210838fd1158e4003e08201cf9eebcf55ac628e863343adb2141ed5f5bb60001b620073b9565b6200228d7f498e06133caed9b6d7bcb977f3c54ca71e1877105cdf61f5b4db67b9e195f9e860001b620073b9565b60009850505050505050505062002a04565b620022cd7fe2cefe2c1e76bbb20e9adf9f2b2e8bb20dc1d43b5df44b925eb29fa11823106e60001b620073b9565b620022fb7f5ef42882416a805bc8f4c878f12c296c429ce6534ecdeaf39d358f9e9fc0b76260001b620073b9565b620023297f1b5e1b499a99c47da6c323d400ddfc5a427cbf9fff81f40c9296a33898c42abf60001b620073b9565b896101400151156200296c57620023637f86057a41dab64bff135b7ffa61a466f7b16ab0394a8a16422bef6c707555b02b60001b620073b9565b620023917fe123f61ff8622bfd89a988931bffb9f89e253f14353129aff641e050c5017a0460001b620073b9565b620023bf7fd8781b1a6b66a879aa918331fb247b12cc9048f302d89869dbdb3094e6a1ff7d60001b620073b9565b8260a0015161028001511580620023ed575060008360a001516102a001516040015167ffffffffffffffff16145b156200242757620024217fcce7522ca9c2b75bacded5bec499468412268ac78ef12901aadde8221e73162760001b620073b9565b62002456565b620024557f566dfa7350c0db209aa46f7d40032eabb8e380cc9979c1f5679eed25280c1f6560001b620073b9565b5b620024847fc5f8cf93a32278525218c4e16b1ac6b546ae9f5ae963304ba2371d3287c49c0c60001b620073b9565b620024b27ff9baa8263878341f897a245415407a43004546e8e4645784372ff855749ea8fb60001b620073b9565b620024bc6200cc61565b620024ea7f30ce5e10943a966f585bf4df9b137b397a1a4b4b73f402cbf2f1db0cd4a98a3560001b620073b9565b620025187fed76efb3fb31a856bd19d83eaa21a5f57408046ab5da1bd68f34e5957fb01b7d60001b620073b9565b8360a001516102a001518160000181905250620025587f64cb5bef198db57354022901898704fd8bef66fcad9072d7523c535a2697d65260001b620073b9565b620025867fdd103d623be02d735e31347324c0c514bf10cf435c9cb66f3efd209f062b027060001b620073b9565b8560a001518160200181905250620025c17feb71a0685f7b60e74fc77cb558af72e4858175b0fd4647fdf08aab08159a3e9660001b620073b9565b620025ef7f51b2b8fad82f352bc9dcfb6a39b5fac9f9689aff7a827646f51afb82de21eecf60001b620073b9565b600087511115620026f657620026287f0406d87923fe1fa1e5ee329503d20e485c45fe6032fc7b6a30290fdf2c8bf9fc60001b620073b9565b620026567fb99541977541724c677117c477e674d068a04a3b693b722bff8e6075b971580660001b620073b9565b620026847f09687d2867fd306615aa40a5060bb23993013c7c2b053bd5fc672787d9ea24de60001b620073b9565b86600081518110620026bf577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516000015163ffffffff16816040019067ffffffffffffffff16908167ffffffffffffffff168152505062002725565b620027247f6c2a5f6c037d8863e6fe8813bd47a7c71a761a9fd5a5a2484bc57fb5df0822ba60001b620073b9565b5b620027537f104ccfd88b5b213e62c5cc6a623f694dec46621d1d81a4bcc5b1bd7f3857c0c260001b620073b9565b620027817f7c75d777d50ba4832b156ed258d430ba6df5f57e456e8e87f121ada14f161c7660001b620073b9565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e19aeff836040518263ffffffff1660e01b8152600401620027e0919062010953565b60206040518083038186803b158015620027f957600080fd5b505afa1580156200280e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200283491906200ec52565b9050620028647fb4fe9c9fd00232c0d9dd18667b963f679cba54ea9f5a37313338b77187e195e060001b620073b9565b620028927f7c3742e0ee3df7b4adad58176dab410c0fe0f229f5fc681d45d8fd2c536ea2cc60001b620073b9565b806200293657620028c67f8363aa485d172a8d9c85fdd107791eec0afcb902ef42830e4297c1d95f2c78d560001b620073b9565b620028f47f9c8e3b748cbc4eaf4d1173e324acb5077f6b1992befcdbc983ed877a1271917c60001b620073b9565b620029227f684ea5201810e9d25600327497310fb8f683536f06e88823e879d23cf06599fe60001b620073b9565b60009a505050505050505050505062002a04565b620029647f18f66947524524b68607518bd00d1a8234cd736614d7568f21c80f2b0f7e1df560001b620073b9565b50506200299b565b6200299a7f10da145549d112cd2e2043ea10552f530e6e3e9319350adf88dd17bb55e35e8060001b620073b9565b5b620029c97faa7776014ad7ead320e53e7eb756c0466a3dfd96ec72c68fd35fec5669183f9e60001b620073b9565b620029f77f1611fa7507d3c7a7b97f8c114810f9d64e7b6e3d8b60f7c4292fbf658eb708e860001b620073b9565b6001985050505050505050505b92915050565b600060019054906101000a900460ff1662002a345760008054906101000a900460ff161562002a3f565b62002a3e62009f48565b5b62002a81576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162002a789062010769565b60405180910390fd5b60008060019054906101000a900460ff16159050801562002ad2576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b62002b007f2f41e2cd2413c64350bf786bb4f8d541edb1c957b6e1df4b16455a4fd6b6be4460001b620073b9565b62002b2e7f3bccc43751f4498a0355e02df6206998ef019dad44c4521cee56097b54c2905460001b620073b9565b62002b5c7f40165c08f35e772a70c71c562f9cac1b70d03fc7adf6229f46a0a85388a184f660001b620073b9565b856000600a6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062002bcb7fdbc6002eca082629f6f8adb05be36d425ba8bc4e18f3b8a209a27bf4fd9d939f60001b620073b9565b62002bf97f991b1dc2ab63d9e73dbf297609a8d6ef5e7abdf35c02125e8b0334e1eeaab59a60001b620073b9565b84600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062002c687f2d1f6f56c2bf4f80910c7f9b91cde3179b8d242c1d863e535928e8c46bd0ce5660001b620073b9565b62002c967f32b7539161197c774a2b90300b803df603ecf18e4d7205a0a1ab386b3f18c2de60001b620073b9565b83600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062002d057fd2e9076d3b1f29391c78961dd81873801916dafd0211407129f4bb398890795c60001b620073b9565b62002d337fd4f87265a8e455774008a82b6f6e2f228da126f38c9af3cd62e718ca6e7b995a60001b620073b9565b82600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062002da27f49516454ad1909cf8e2b080967efcb5e82ce25ed64ebc256d218cbdf361c057260001b620073b9565b62002dd07fad130773bae93052fc71e17926cdc3786ef2bf3b2c3eec744b3ddb0bf571edcd60001b620073b9565b81600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550801562002e335760008060016101000a81548160ff0219169083151502179055505b505050505050565b62002e697fa26fb8bcf4f1771b2d7e5e63707d42330bb9214ab6033606eeca3114d809238760001b620073b9565b62002e977fc7f0202eb3b34b0f924ae85277882ec7f79a2d7b8356eef88daa6adc6bd2515760001b620073b9565b62002ec57fd3905f6d07088ed9979f11a10da74291ffce44b443893defdd1c665877d0b6b560001b620073b9565b60008160000151826020015160405160200162002ee492919062010404565b604051602081830303815290604052905062002f237fe13d43abe6c90cd08f5456e8b2ca5fff2777219ccf19f80b16e39e7f64d94b3460001b620073b9565b62002f517f92abe34b44d997f7c8261084d1ee1462fad86ef77f80e2ac8e3148673f9fe0ad60001b620073b9565b8160088260405162002f6491906201044d565b908152602001604051809103902060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050505050565b620030277f80af9ff788fc10746f90a7644a9ff73e43de402ca7b6e11b427b2e57de3cb34660001b620073b9565b620030557f9369b24158b5773b6eb142032c753c6de77cb435a4de3c0f5230c082baa65ab560001b620073b9565b620030837f2ad417701ed8b01f165513726224b5208aa68af68a38dbaf230ba9733c9a72e860001b620073b9565b600060058360405162003097919062010434565b90815260200160405180910390209050620030d57fabc5f60ca7d1dfaa706e98a572065d20c946f34a3d7d23ec7987ad4d099738bb60001b620073b9565b620031037f8e83d575f8f37dfea9f380f944bbf77494a65a0bd507c7d584f2d23cd214f16360001b620073b9565b60005b825181101562003241576200313e7f291c615dfae5b522de5d1c1cb4e9eaf3d96fbca69d843788da59ae79b7a3bd0860001b620073b9565b6200316c7f2d944243b9bddf92d4a48845310078ed2cf4254eb2f2983f0897f026efee7f3d60001b620073b9565b6000838281518110620031a8577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101519050620031e07f6a7ac2c5149de31d04e87289ab173ae1a47dee174bde2919e8e2eae942b2e4af60001b620073b9565b6200320e7fe7c522d4634f7e51363a25ead2c246adef2f41b3bf391b640af8ec708a03546d60001b620073b9565b620032298160000151828562009f5b9092919063ffffffff16565b50508080620032389062011135565b91505062003106565b50505050565b620032757fab1b6ae163aa2e0fec62d6dbb52f562b0c11107961013b05421348d52d3e73eb60001b620073b9565b620032a37f2278269b1cfbc50c666fc5331539577da103648da63bc3f5efae711797c38e4060001b620073b9565b620032d17fd48153f1bc8afeae9380a460749ee46c5190d9c419f6d37a905781664d52259360001b620073b9565b80600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002081905550505050565b6200336d7f478bf5bd5e7dbaa523ea69eda77cb8495a1a1ff7239a0a6f189a0973d933a10e60001b620073b9565b6200339b7fe29dda0f9d87cc11f63a2ec441e6d8d6c1e0b0f8b61db9e2430ab823a605836560001b620073b9565b600581604051620033ad919062010434565b90815260200160405180910390206000600182016000620033cf91906200cc94565b60028201600090555050620034067efad54934db310a1db713b9c93e7269423a2b3536bc746c64638e96fc485e8160001b620073b9565b60068160405162003418919062010434565b9081526020016040518091039020600080820160006101000a81549067ffffffffffffffff02191690556000820160086101000a81549067ffffffffffffffff0219169055505050565b620034907fde8b9b760af09b3aa0b7a17627806f5196254eb5936592657d802308275b7ba560001b620073b9565b620034be7f667c7270e3db2c393e325079f1c42c98f5f8dc240afc97c2e228dd59ca15d1ca60001b620073b9565b620034ec7f431991219f3c3d47a7c62a7c600a38ef35480a50ad3f20c6b8067ba0206bb37a60001b620073b9565b600081600001519050620035237fe61f871a49ed694e9fc45fe012bd30cb7ac37251e6e8e5be461bfba02dd99cd960001b620073b9565b620035517fe44b4e1a428ae440304511e02428486862160379d2ac030bed49b2286e8dd18c60001b620073b9565b600082602001519050620035887f37c8a6e0bf563b1a46265d01226189c22f6231fafb88e276f0f3f7dd49fe88e960001b620073b9565b620035b67fc3ff640996cf3ec2132cbea9875372a82f61351ca7ccd27ab462f8f7904784f260001b620073b9565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633778febe846040518263ffffffff1660e01b815260040162003615919062010466565b60e06040518083038186803b1580156200362e57600080fd5b505afa15801562003643573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200366991906200ef80565b9050620036997f82cb862f7865a5afaaa4b8e630163c232fda31627b9424afa5282602a92e147960001b620073b9565b620036c77ff3f3f2f560c7dd95784aefa3dcb9a0684b90fdd333f45f302ca3e520e1394fb360001b620073b9565b42816080015167ffffffffffffffff1610156200376c576200370c7f763f0f24c8d6d0427860811393bfa52de3dc5552dc32ed1b787da3358f9b985e60001b620073b9565b6200373a7f2171b27be3e70c20e42724b33e45e8cd952c29dc33a8c49ed787345fae8a854460001b620073b9565b6040517f26841a7e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6200379a7f96bf9f6abb6718472a89a02014655cedd02d9eea7ab5f2ed67dd4a43b68ce2f260001b620073b9565b620037c87fb904005eb0cf39c25ecf40554858ed6fcfbea47011394a7492416d1e8f05088d60001b620073b9565b620037f67f4e716e60f83e39dfe0e2c81ca10e7bfa538424639198f32cefb8d7beef5a35e960001b620073b9565b60008267ffffffffffffffff1614156200389857620038387f9bfac304b9cab5359b1ebf9f64acfb68bf6772737b9c62f9e9844536c24380ee60001b620073b9565b620038667fc875b6911556b7889f3b56de25e365269d4e4156e598c06134a5bc829356175a60001b620073b9565b6040517f26841a7e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620038c67f9f76c73a085d5dcba5b121f015dc64ce51246c7797a28f59a85761e3117c2df060001b620073b9565b620038f47f6fa469a54555b4473f41910a6cffc15ee95264b922678645137f27f9f2d8d5b060001b620073b9565b620039227fc9c54d1d82e0b475ccb8cf34a8bc11ed8f0e971d6fe65c055b6bfb034422d8cd60001b620073b9565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d7866040518263ffffffff1660e01b8152600401620039819190620108d4565b60006040518083038186803b1580156200399a57600080fd5b505afa158015620039af573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190620039da91906200f01d565b905062003a0a7fdf586645913f5534a1aabe70f303ec1648758df327a49bf116fbeef2699dd28960001b620073b9565b62003a387fdf424496ae761981dbded3ba8c231100ed17d8e8870365f83424e37eb7aec92660001b620073b9565b600081610100015167ffffffffffffffff16141562003adf5762003a7f7fca070eb9f348ec1d0adf6ec015415922bc180299ad9c31fa4234d7f8e42c156b60001b620073b9565b62003aad7f52b4058550c148c2f0b79344148b1327225559b65055fe253be8e0f6a234270c60001b620073b9565b6040517f26841a7e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62003b0d7fa61cbde5d9adb043f84800f44e858aa7224910b3982a172ac44eab089d61ffe360001b620073b9565b62003b3b7fc1959350b64ab9f61df5d4df4ae4cb3a353f43fefdf48d2b401e26e1a107945660001b620073b9565b62003b697fdd7481b9f33e37f240f007826d2846725113a8ed4adf66630aeeb30c5dcb477b60001b620073b9565b600080600a9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c5c81b2083608001516040518263ffffffff1660e01b815260040162003bcb9190620105e1565b6101606040518083038186803b15801562003be557600080fd5b505afa15801562003bfa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062003c2091906200f1ce565b905062003c507f7a7a0c93aea5a57b1db2bdb253e6229690517677b9f29dc8ddf6b112313565b460001b620073b9565b62003c7e7f7e8894f83ff0cfaff8da165fb61ecb962d9ecd4d6e1f0ac63caccee517f9c35760001b620073b9565b600043905062003cb17fce56d9eea9cdbe46d55ad476369fef4022f210ab7add2397645ba74e3de9c1aa60001b620073b9565b62003cdf7f208fe3968fbd6dcd6a0e456bbece4e03dcebc7ff7c0f1741dba23a776009f04760001b620073b9565b808260c0015167ffffffffffffffff168460c0015162003d00919062010d02565b101562003d955762003d357f4f48310bf022a0521f33891749f17d12ed7ced384777aa01ca314f2e2ff9ed3460001b620073b9565b62003d637f03a9cb743a5e9bbf61cd83c00c8b90940bca85d47a02164e80480a107a1ebf7b60001b620073b9565b6040517f26841a7e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62003dc37fa1c1531ab3acf95261d9b2969670ab5b4008c9228fc856516ea7ccae7bfe21ee60001b620073b9565b62003df17f11430196cc337324af63b84a80f51a8234cfc39d6999fe5662de45e84c94071560001b620073b9565b62003e1f7fad3069dda3b0afd2ebc3198991c3e049280ce0378a34b66e4c0f49b00ffd95c960001b620073b9565b600062003e2d87876200a5e2565b905062003e5d7fe3772f0a256e29e0cc62671400c74ef97b75f9ed6be3ca06722bf440eb28ae0860001b620073b9565b62003e8b7fc567b8a687cef7503f084b0ea0e2737de8c96d009afe1a129fd04974a7ed62cf60001b620073b9565b600062003e9b858584866200a6db565b905062003ecb7fff34ce214d3c58b6e41fa49793c512eae8b807d323f6fbfbaa74b493d0c34e2660001b620073b9565b62003ef97feb9afd447d0119eec6c0bf812d6bcb349a1eb45e0435856a057de64e3f125ab060001b620073b9565b60008167ffffffffffffffff16141562003f9b5762003f3b7fe65b9897435d700c8304fa5ca5ff6ff9dd29bd0733e551d15b75fc36c5cc345660001b620073b9565b62003f697f5237757807b06ddfe1cc56cdda044d550cb19fab8dca0a6d0fae6478381d13a260001b620073b9565b6040517f26841a7e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62003fc97f537c13e5110ceec19a55abd540176c6136e77cfdd8746061ae6984d5301fbeba60001b620073b9565b62003ff77fc1c6b1599e638b5eba413331567c540a63623696fa47257665315d642ad0b0d060001b620073b9565b620040257fc2b3b51d26d7c6e045c54a94fc06eb36c13011911f2ca3bf9dce12714bd404a660001b620073b9565b6200403385878684620089d3565b505050505050505050565b6200406c7f4f01fc12b60872eaecc1f0b3232a5405e43682b8dd7e03838fb0231b985bd15160001b620073b9565b6200409a7fc678e8547e4dd67c92a03295c1bd6bb2bf81af3894d447d396ca3882bcdfadb760001b620073b9565b620040c87f2e4b897e127bf439763d91a388841e7b864f12f52db11e498004b3f4a1a5eb8e60001b620073b9565b600080600a9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156200413357600080fd5b505afa15801562004148573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200416e91906200f1ce565b90506200419e7f487d6fc02beddb9c6e7375d4f9838f3c1ffd4e3e5292a294364c2fbd0598b08660001b620073b9565b620041cc7f5b0add7cad7c89ec11e13bbfe5042280862b6a3faf46ac21f1e8c30a33b18d1760001b620073b9565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663defce5d484600001516040518263ffffffff1660e01b81526004016200422f91906201051c565b60006040518083038186803b1580156200424857600080fd5b505afa1580156200425d573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906200428891906200ee3e565b9050620042b87fce0c85e23f32a73d8a607d82b18d35a4b0ef6dd773c9ba4fe7a9bb8ed9005fb860001b620073b9565b620042e67fb7694aea5332daec8d04ad02d30445113ef1ca4e0794e3376e8b7d22e1046b0060001b620073b9565b806102800151156200447957620043207f30076260222f766f0e692dfec16b9782badb1a945e187fa107030cfa918655a260001b620073b9565b6200434d7e9d3295b06ae142f00cb08da4b87ed4931dde71b0bd55241bf8c6119345a00d60001b620073b9565b6200437b7fb6a550dc706d56c15d71c667b2a82157499eac2027f3697180b152e53b85c3b960001b620073b9565b806020015173ffffffffffffffffffffffffffffffffffffffff16836060015173ffffffffffffffffffffffffffffffffffffffff16146200444557620043e57ff789299c91951d87c97785833c8de3a5df9332047f3afb77e8de7f7a666a2df060001b620073b9565b620044137ff73ec4b5c97f8f40dc352934e3e8f9e383c2ad6e3c10ceecb8e9717367acb57360001b620073b9565b6040517f65fd380a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620044737fda0dff376f24bba595edb6350018ba8308ce77ef33987d67471c54dcdd5a2d5960001b620073b9565b62004c52565b620044a77feb240abeaf690f25ea4e088fbe6efced9ac0d045a722fb957ec4daa1691b8e7060001b620073b9565b620044d57f2218e7a4a77fc02cd161f2ba253fd7a358f593ac611b48d078aa60c60d14e86a60001b620073b9565b620045037fe4170e320d89e4dbd42edff5bfc4a07ea30e8dbdd1c4367878a0f52dacfa5d8e60001b620073b9565b6000620045337f9c991a3db94afae4abc330aa4a7340a9c6401544895e38bab9a58acae938506b60001b620073b9565b620045617f3fd7c9bdf394090454d4d35e65ab3ba6a0d3bc31cb58477e392b906de6cd37b460001b620073b9565b60008261022001519050620045997fcb5f52505ff28f5d770a7520a810127b4d11a3fa6c6cf5dbe7fa3228b5be229260001b620073b9565b620045c77fcf5ee8ac67b2714408d0f6eb7da82f61ef005ec3c189842600baa5fd33e48a2e60001b620073b9565b60005b8151811015620047b057620046027f5691a547b7f59abed046b7338dae5ef5acf68f9310462cba6a0b1e8fca9658e960001b620073b9565b620046307f5e4c8049834abd2083c7ff7cdf03f688973c319cb5959b3980136ee504e1c9a760001b620073b9565b856060015173ffffffffffffffffffffffffffffffffffffffff1682828151811062004685577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1614156200476c57620046d87f614db3316a5bfa2e12c2b0bf29b848778ef03751d56daedef5eb09c1c69dfc5e60001b620073b9565b620047067ff3f58db16b1c3f95f9385912b232b52db5d1e113d10a5e03581a6d9f434d254360001b620073b9565b620047347fbf03e16d82fc02815197a62c8423507bb4e960b9826f67589401505d1c02736060001b620073b9565b60019250620047667f716a642f926ccd67946dc7e71f597a8985898a5e3f3876dba2377aeee7daf52660001b620073b9565b620047b0565b6200479a7f2f5af5888f79870c350780c94cb8e400792ab4cc31ddf32b8ec083a853e945ef60001b620073b9565b8080620047a79062011135565b915050620045ca565b50620047df7fad69cbc3edf995939c8f266aca6bead961b3fef314476c85214740733ad058d060001b620073b9565b6200480d7f3637d853d640eec20c8eb32ecc33d1716dfa152228c7f43f5ad3293754d6059260001b620073b9565b8162004af457620048417f48049f60921fe4073ad1fb4026b35d5e4de39d93be8b10c545c90a0d892ccb7960001b620073b9565b6200486f7fa12c1c03c908c2d92863891d6bd5f9f1d39f461229e63ec4b00e5f6ec737c0c360001b620073b9565b6200489d7f7bed2600832e78b78af2b012119d0f37be5ebb3dc249f02a1265aee5a31a091060001b620073b9565b60008361024001519050620048d57f15d9ad84822fddcad19082529c5f79872ab53d83f62c47747736e7597076fa2260001b620073b9565b620049037fe041bfe03372568f851a67eabbd54dcae7a537a801518b51a9b45d0c6a4e4aed60001b620073b9565b60005b815181101562004aec576200493e7fb7904ba779d7dd9e8f94caf3a8a39912eef06172e6376253adedfa33ba58f94c60001b620073b9565b6200496c7f5e36c656103aff715e4ecd8010e2a0e4f30e3cea239491d776226430978d5eef60001b620073b9565b866060015173ffffffffffffffffffffffffffffffffffffffff16828281518110620049c1577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff16141562004aa85762004a147fd7926d66c6d9a69d70b5f94880e41402198d0e19e7c2380a72dbcfc44829bfb360001b620073b9565b62004a427f73eba1007c4c6817bd9bb2d5bc4ef045ebdcc590e2f8cf2ce7d6a853b1db5d7960001b620073b9565b62004a707f08a8d1fdf9c8598e46d904e8abdb35a9012acd5e39a18c37618889f97d2b84aa60001b620073b9565b6001935062004aa27f70711224d311837b6089458f3d735442962e4628c11c6934b51f81462f7c440a60001b620073b9565b62004aec565b62004ad67faa48d71a50b3a8cf5ff291deced677c1b3b9687dbbc914d35e00cbb86b9d49c660001b620073b9565b808062004ae39062011135565b91505062004906565b505062004b23565b62004b227fceaf56f33254053bf278b675d1cdb7e7b3d58269dcfc916e4d9a49e62ac2bd8c60001b620073b9565b5b62004b517f90f9050d14f8dcc11d4ee3cb5f447148f594f2411efbaf44857620115f298b0260001b620073b9565b62004b7f7f94cbbce47d25e2c7c87588dc7f201e3225745f49b257c3809ac301ae24340c8660001b620073b9565b8162004c215762004bb37f44f05e1b273905dbca6a070144e96243b5fd998fda83c774b9552da4e555751760001b620073b9565b62004be17fe6bb7edfe12aa96440400836183eb37c83ba965fb35d95f1ff7c365e0a0b760e60001b620073b9565b60016040517fcd442d6500000000000000000000000000000000000000000000000000000000815260040162004c189190620105fe565b60405180910390fd5b62004c4f7f68fc95c78e0cd8a47c6ccbeaa0ea7a6fc1029ad7b9587e13682bf19960af81d060001b620073b9565b50505b62004c807fd79ce2cd19d14061c65ec4617ddc245d9cbaa04836dba2257f25acf33ec96b9060001b620073b9565b62004cae7ffa3976faa4f233cd76a51fe5deb5f27f877c98aba8ac9722b834eaf0e487ed1560001b620073b9565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a65374a85606001516040518263ffffffff1660e01b815260040162004d11919062010466565b60e06040518083038186803b15801562004d2a57600080fd5b505afa15801562004d3f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062004d6591906200ef80565b905062004d957f4ce49860b5e0360854214064de601d9b0428e28618beb5a7f3b975e33144776a60001b620073b9565b62004dc37f9baedcc545482aba9a35de05fb02cf25157c53ace06a6fdbae9240f9bbd34cfe60001b620073b9565b600062004dd483600001516200ae25565b905062004e047f66deb016a2c7340a2b93294377bb4375d8f055d9c204177cf91cee389b61f32960001b620073b9565b62004e327f2a26c5a0bf9d2b45d0aa4f534be4c04c0f27377969913c6c1e027bc02f8fabaa60001b620073b9565b60008560a0015167ffffffffffffffff161415801562004e56575082610100015143105b15620050b45762004e8a7f17a349f6c508bd849237d6ca6c6406ade0e014bfb2e4bb7d7d5839b50c234b4860001b620073b9565b62004eb87fcdc731413ef0d99abf34ba6f20e0a1b8a910773206c7dc5557bd17b468716cd460001b620073b9565b62004ee67fbf8e04e1325a21512d5f274c9677a7ac6457f1d15a44e355a7420ca13585707060001b620073b9565b60005b8151811015620050ad5762004f217f31102aa8ccd8667514b8fd08c7636c0a1e602296f20bd0460c7cafe949184ac060001b620073b9565b62004f4f7f4aa71b8fadb34ae7c55156a583bae72d4c7c20aaa1fd643567645d2c3695bae560001b620073b9565b856060015173ffffffffffffffffffffffffffffffffffffffff1682828151811062004fa4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff161415620050695762004ffb7fdb154f979c4973495fe8d10f94dc3be44542e551b234457a7de73c79343fe1e660001b620073b9565b620050297f076b5c095a55d80c489ebef76b914ea971c66582178637fe4793adf40ef62c1360001b620073b9565b60026040517fcd442d650000000000000000000000000000000000000000000000000000000081526004016200506091906201061b565b60405180910390fd5b620050977fc61d5295071b35e0002f8bb10afbcdeda61acbdc2c3ba3544be9b9788b5e57e760001b620073b9565b8080620050a49062011135565b91505062004ee9565b50620050e3565b620050e27f01c59d689cd4736da2ce845eac2bec1e84ccf0a0b6b714636c00f8470894978760001b620073b9565b5b620051117f2aac02d3a10e51b456eb080c2e57e6a569ab7e451f3eb30a64ff028c82e19b8160001b620073b9565b6200513f7f1e337c5ab7476da9e627b4c9a65a7a2d733166bf7bcd89094c82855395909fad60001b620073b9565b60006200514d86856200b1d5565b90506200517d7fb60cfb854ade38bfa3c6d151864b7e8d4c4c4a920bb011816fa1f14a13172fa760001b620073b9565b620051ab7fea95d86ac872c4f42c55813b9986b748f6cfe30b725af4bcc3d26676eab08ff560001b620073b9565b806200524d57620051df7faf68160529ea7005d5ab6186a73efbd7a876521b23c1e9852bde0728313dbfc860001b620073b9565b6200520d7f31056821ae0ab0bea78d33444ea6f3f9230acfabf905d5eb56b366f7c83a063160001b620073b9565b60036040517fcd442d6500000000000000000000000000000000000000000000000000000000815260040162005244919062010638565b60405180910390fd5b6200527b7f90d7a0451ee7b3c0db2cc10dcb87a8055ea8e5a60cbf6245a7e8fd7ab5c5064360001b620073b9565b620052a97fd26c94a085f9341efe607f7a861d71e012f78eb068ee1aa365edf196a8dedf0660001b620073b9565b620052d77f466de43131e5884c78e632af312cc7f6133820eac22bc31ba9989514f4e898d260001b620073b9565b6000620053077fab1e7ac727887619d8a3dddb863dc6da34ae49dde914d2de9acbe597862e834660001b620073b9565b620053357ff958b70cb5124380aa154e8cde8ecc6081c99c5de7c75d3334c078d3e2843af660001b620073b9565b6000620053657f9dc98a1902d59d0e17be140d5cae13f2a378b9cce070a5597f2b26a43bfaa59960001b620073b9565b620053937fd958756a3f9a192ac71527ed38462281cee0509ede955ebbb6b505c69b7f477360001b620073b9565b6000620053c37fbeaadde72cea2776801bbfa43017f34e94c5c91c64c5d2c27b3a7a9aea3f688f60001b620073b9565b620053f17f51bc3cd067ba5c32b7c0e16495ff35860b3da1e6fcad4dcb8a6d961fe7421f3460001b620073b9565b620053fb6200ccb7565b620054297f871ba2862f5f0f6b4f318110d643505cc09631946bd23703fda9e49396e7be0660001b620073b9565b620054577ff5479c4c07d735ee38f08b86bf5600a4acc5c33f308e4f4df39434bce32ae38560001b620073b9565b600088610100015190506200548f7f2f0765d1a8599ad4c6b1339a5deba4740dfaae208c0417da9877650c77f4de3b60001b620073b9565b620054bd7f1097c5bdb8378cea33bce60eab4e20c80e72a30cf3b48ef3b287d13d17220a0060001b620073b9565b60005b875181101562005c0457620054f87f97e5d56dfefdf5d478a924d5063e97860ee08277dfda1c819725575dcceeea5e60001b620073b9565b620055267fa5229b4b1951fd2f97b96136b8dc82c9a23d400ae22fd280572ee5b22826250460001b620073b9565b8b6060015173ffffffffffffffffffffffffffffffffffffffff168882815181106200557b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff16141562005bc057620055d27f991cda3e4a8a87eaf36d615f25ca1fae390498836879bb35928ea3749721564160001b620073b9565b620056007ff24062659cff4a7f2249a0774f377856aa16a6e6cf3affa6c0ae45f6e80200e860001b620073b9565b6200562e7f1dd66f97cc736107512f4dc7c857ea715189193e8773cd04845884b4b200045060001b620073b9565b82604001519350620056637f4af059b9b5584c4b47afd33abbb21ac42897c5337c111095ff194bca8c8d2c6560001b620073b9565b620056917f66e1ad802da42798aaab0c16856cfad847c2af03a0a042fb55843d533ed397d160001b620073b9565b8960e0015167ffffffffffffffff168467ffffffffffffffff161480620056b757508143115b15620057be57620056eb7f5c7ebd9f13ed692d8bfa45e5b5fe2bb8b96f5e7a2de0c9964221fa772e79c25460001b620073b9565b620057197fa02ce339ec7824036cd4f92a803f864e8d00b6012efcde8530f9b4dade32859660001b620073b9565b620057477f4e4e55d54ff34e97976379a3b63ea0a37c516cfb611d77a44a4302d4520661e060001b620073b9565b6001836080019015159081151581525050620057867fc4606c0d1a325221799bd3893b6b65409de49f58959fc53f74138c6c107ba78a60001b620073b9565b620057b47f2d335617b01d33ca1c43261649dafea7bdf7d55a6e59c281c92c28b369fe2c1c60001b620073b9565b60019450620057ed565b620057ec7ffb85d4ed3488c5279346824de02dc6979f3f3ba5b4fb93ffb2ddb0e89a2efb7960001b620073b9565b5b6200581b7fbc27faa4ad6d6257a265a86e66ea9e498d06d40ddc64abdb4cef1701a82e650560001b620073b9565b620058497fe6438c61e9fc16dee939bc612f7744982fbbec328c27e65f276997976a37f95a60001b620073b9565b8960e0015167ffffffffffffffff168467ffffffffffffffff1611156200590657620058987f654b76213247bef1dd5b384d007295b3faa419cf820b13e2924736ff9863cca560001b620073b9565b620058c67f31abf154acfa9c890b7b44f2730b36b87c8ff5a02f3888b08c0b3b922c4bab0960001b620073b9565b60046040517fcd442d65000000000000000000000000000000000000000000000000000000008152600401620058fd919062010655565b60405180910390fd5b620059347f57d9467eb25d53b550c70f02612c5f8f22babf3b9e58a5803dd49e46941f799d60001b620073b9565b620059627f36e7b707deabb58bebbdc28821f6d0f50b6d16e0bf895d7bc25f54c2e62ed39260001b620073b9565b620059907f5ecf10f430a069cada01a8e4efafb2f19da916e44333eda7bb14d5e66882a16860001b620073b9565b6000620059a28b61010001516200be7c565b9050620059d27f8f6e4b45dc1d470189e41089d3dac7061847ddc946b44bdee6f73b3bd8acf28060001b620073b9565b62005a007f908cd1a75f898875e07becb067dc3d87872daca307b743fa0f04657e72e1793d60001b620073b9565b801562005aa35762005a357f4a6395058dfbe966abdbc562d905bb356e766f1fcc45699a6365d3a486757f9760001b620073b9565b62005a637fc449cc6fd29fa6b4ec8d7c52bc1e0385786bea67dd8f140ac2b7bf57dba78d4b60001b620073b9565b60056040517fcd442d6500000000000000000000000000000000000000000000000000000000815260040162005a9a919062010672565b60405180910390fd5b62005ad17f344fcf6ec12f1a252194c40d6434cd9bff7e5b78a745e683f70f4aa871bbce5760001b620073b9565b62005aff7f35b89a7804e6248f89e4c0e90a60fa428766ebb0191f169770da213949ab36b460001b620073b9565b83604001805180919062005b139062011183565b67ffffffffffffffff1667ffffffffffffffff168152505062005b597fc2c7ab60b347c201481cfdba2ed6f8346fa7ebda66dcfdfe79c90ee226180adb60001b620073b9565b62005b877f0e1ed403a7d2df29aa94c56b17cf5574da135107b90a5e03d367ae712d1da91d60001b620073b9565b6001965062005bb97fc2e03d82f409d0de8d0cf12fc4730ae724369d8086da8b82a104268105be51b160001b620073b9565b5062005c04565b62005bee7f842686004b2a4efd2e924a9be272614647bb5ea7f90f530857899616c5aa60c260001b620073b9565b808062005bfb9062011135565b915050620054c0565b5062005c337fc769f587e08ef033b47dd14b102ea40178e2023b0d224f6bd410714f2f0e798b60001b620073b9565b62005c617f7fed0300b22ee14f3a154f136e7a0967bb643f2a3fdb006664404bb7ee94352960001b620073b9565b84620066b55762005c957f580d44ae7f38d805036aebe0ae726da384034cf1f31e5984f98c68150393d9e060001b620073b9565b62005cc37f1641ef084252f5710d5c1c2486a78ba1eab7ff91523cb8a32a811c1d63f717ea60001b620073b9565b62005cf17fbd2530872a955b63ac3b5571480e163bcfbe93e7cd2953fbd947b2f4b205486c60001b620073b9565b600189610120015162005d05919062010d5f565b67ffffffffffffffff168751141562005db45762005d467fa9b61643be691616d262115cb43e28206088a87af72dfff0ccbd650f266d133960001b620073b9565b62005d747ff11470d72c67bfef7ee351fbc13dbb959bcd67cfb073e21c8913abeaf7ec7f8260001b620073b9565b60066040517fcd442d6500000000000000000000000000000000000000000000000000000000815260040162005dab91906201068f565b60405180910390fd5b62005de27f493b484f6f70d2c6fffbf34a90e8045e5d7e0144950f28380bfb975c0029aeff60001b620073b9565b62005e107fdae5e583cf336d5fa2237f5431bf8b8ef25c65005d5e41c7582f3e9ae30e923760001b620073b9565b62005e3e7f0833be629a04c0019c36ccdfe8d1ef8dc5858520f674082d214f9ed2fefd3e5f60001b620073b9565b8860a00151896080015162005e54919062010ddc565b67ffffffffffffffff16886060015167ffffffffffffffff16101562005f105762005ea27fd73abc0253b8ef6f032994447af98bc4a5c86e4b588c2fea9a8ff0c7edf4f12b60001b620073b9565b62005ed07f5bec59ab1c32e516cc7e03e081b32b7ce77ccfe6a4e29e4a90c1e057d3ce046e60001b620073b9565b60076040517fcd442d6500000000000000000000000000000000000000000000000000000000815260040162005f079190620106ac565b60405180910390fd5b62005f3e7f887a6074bf649d35fc26a2d95772c9e5dc04d402bf7f5881739c5cd878ed192960001b620073b9565b62005f6c7f7f83ee04645312cf87972f31aec4d93a3a66b2e28b60f5e7576ebcabc48ad5d660001b620073b9565b62005f9a7fcddd69650b0b0f84c96262441ce7e3062c7a59cd4ee4cad02f9b8d42e3e74a5d60001b620073b9565b8860a00151896080015162005fb0919062010ddc565b8860600181815162005fc3919062010e60565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506200600e7f2d313e4e7b939ddab0d139a4153669e99a187ea4a88b5a049faf69714bf2ea6b60001b620073b9565b6200603c7fe3f1b6837b84f034321e86b6a2561ddc332227a505de6d025b7d08562f6e93c960001b620073b9565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c819d86c896040518263ffffffff1660e01b815260040162006099919062010817565b600060405180830381600087803b158015620060b457600080fd5b505af1158015620060c9573d6000803e3d6000fd5b50505050620060fb7f69c5886ae649ec666ae670ad3695eba88beeb2e7c8be7150b00318d5dedda96560001b620073b9565b620061297ffc69c08df168fe45cc8bfe55b036f685ac76948277b96d2ee634035a9a56976460001b620073b9565b8760c00151826000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050620061937f9bc5fc0d41cbbe637f5db991d2f9d463643190e7ea26e2e2057cb78b77e6576460001b620073b9565b620061c17f92607830c0537e5917bd71c4c9d319a6bcb974ac6f7d0f6e4e043a32b0b0a10860001b620073b9565b8a60600151826020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506200622b7f96d0410b004b6d287622cfc6498b48457e9abf320f2ec54cb945bbe9b3a5a38460001b620073b9565b620062597f2206c0c67469616760a2c6b79525c7bc9f3eb7649404b6cbb711fa733dac8f1560001b620073b9565b6001826040019067ffffffffffffffff16908167ffffffffffffffff1681525050620062a87f91fb238c6c32bcb0ca62b7038d4aa9506dfbff4e800a5a80a8a06be69247956560001b620073b9565b620062d67f685a3098d28d5e6aea1aad476c3b918395404a48fbe6a2cbfb70dbac4ad65d8560001b620073b9565b438260600181815250506200630e7f22df9e600b94769026c7c26fa6efdfa71c08457e972e9dce8c9904c8d74d2f9f60001b620073b9565b6200633c7ff4f388e2e5ca21fb41f1cf62966b93542fac989ead09b54f0fb75e0f00e6740560001b620073b9565b60008260800190151590811515815250506200637b7f2168f4ad4ec31a4a24c4ee020f43222cee5da4e81fe52efb8f5f505053e99bb460001b620073b9565b620063a97f8df3530a42ab5391f144ab40f00833dda660b17504ed139c575ce968619164e860001b620073b9565b600060018851620063bb919062010d02565b67ffffffffffffffff811115620063fb577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156200643857816020015b620064246200ccb7565b8152602001906001900390816200641a5790505b509050620064697f0b55ede9a25e214eff982b82950e8794da059d861f525b4146d36c0a025581b560001b620073b9565b620064977f05eff38d788bbfc4f744769421c28568bf9019ac5cd77b929a40767234ac0bad60001b620073b9565b60005b88518110156200659d57620064d27fc71aea22379865bc3e4d8f805cc058f9b5b2280081547230dbd57391f634bfe260001b620073b9565b620065007f3ddb6da24c09a6ec0cd5bea169f3020e7a024a1522607c2a595208a71d914e6d60001b620073b9565b8881815181106200653a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518282815181106200657c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052508080620065949062011135565b9150506200649a565b50620065cc7f3185ab22d7a756c409c1fe36dba1302ae619a41889aa45fe1056aa4dd532dbf660001b620073b9565b620065fa7ffc0e9e2cecba13d9501baa9e074f63710e36de5d315099a1326ff80b329e39ab60001b620073b9565b8281600183516200660c919062010e25565b8151811062006644577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052506200667d7f50b9fc14544e4153a7ac873fd958cf5d8add32ad7b334f2ec086e13da384e94660001b620073b9565b620066ab7fbf200f643997f0c31a632b98b0ec1c9288755ecb59349ad74ac5588b850d8aa460001b620073b9565b80975050620066e4565b620066e37f40ff5548ab414950c9640130aa0e5b8d4ecc24b2d07e2c72ba25391d06096e0b60001b620073b9565b5b620067127ff283d518ce9e8031d47e7655dbe4863c200f9538974a785d2a14b0cfa652462360001b620073b9565b620067407fa1e1d65e6636c12311cc600108ac74693390ff9aac778798bfc354d4309102c860001b620073b9565b6200675089600001518862002ff9565b6200677e7ff3b77d4b43e3d2c3ad71273cd3c9a24a39b93b2897f298caf5e9592a4ff0aa2e60001b620073b9565b620067ac7f944fffa1ebf48a4c5bb803bc658702a572d56e4dc6cfc39fa35affaba096a8ff60001b620073b9565b8462006e9057620067e07fb737fb5d11ff70fdf1ad2da8a362c1851f517e46ac0064e8153b2529fd8f5ca060001b620073b9565b6200680e7f179670ddfbee6ba3325a8d6ac4d1aa6d4e3f56606172543bce964e4477333f4260001b620073b9565b6200683c7f44d259532cf45dc9ccbe0cceb2e2653206891b795e316f6db6615c125580612160001b620073b9565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d760405180604001604052808c60c0015173ffffffffffffffffffffffffffffffffffffffff1681526020018f60a0015167ffffffffffffffff168152506040518263ffffffff1660e01b8152600401620068d79190620108d4565b60006040518083038186803b158015620068f057600080fd5b505afa15801562006905573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906200693091906200f01d565b9050620069607f40464701613f5ca06f802831aeba536deaef91e53b98d135a26b20fc71e0a48960001b620073b9565b6200698e7f42a8b0a4c0fbb45241871705740a2a702f033b564189e16915339e6a399f6b6060001b620073b9565b896102800151151581610140015115151462006a4057620069d27fd275554609bb8ce56f34a6d79e653e9dc6784fccf08ed5d7c77766f738dc83f460001b620073b9565b62006a007fc3601494002ee296670f7ae30d6e5338fd25ab349df274a63719b71a7918fa8860001b620073b9565b60086040517fcd442d6500000000000000000000000000000000000000000000000000000000815260040162006a379190620106c9565b60405180910390fd5b62006a6e7ff614dcb71a494be5360271ee9ebe88ba2e28dfbc5538781e07f0894bf79bf39c60001b620073b9565b62006a9c7f1103d7a072057aa8f8004cc6739e544342e5170f8d175bd7a75f153840876c0260001b620073b9565b62006aca7f8587c199af6ea0dcb78521edd933b0c373c4daf2ad8740a405b095c609c1a6b260001b620073b9565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663848fd7c8828c6040518363ffffffff1660e01b815260040162006b2992919062010899565b600060405180830381600087803b15801562006b4457600080fd5b505af115801562006b59573d6000803e3d6000fd5b5050505062006b8b7fa269072e18852bc6aba264dacd7a0d8cd661949e390ccf77562ce2610948161f60001b620073b9565b62006bb97f5bab5d4d9591d77fe281fab59ff5465c36d8ba8359494ca185cad34fc778ba8060001b620073b9565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dcf74960826040518263ffffffff1660e01b815260040162006c16919062010875565b600060405180830381600087803b15801562006c3157600080fd5b505af115801562006c46573d6000803e3d6000fd5b5050505062006c787f181536a9806f586125e649c2dce46bd99f2db65f32238582f48133520879bfe860001b620073b9565b62006ca67f57db23d9080db44dc07efe2216df060dec3879fead9f30918badcbf1be33066a60001b620073b9565b60008160c00151141562006d6d5762006ce27f589b8d6d13b6aaf5cd9cb2add9ba35149f2ce0addea2ba7cf6aaeed13e4f38e560001b620073b9565b62006d107f98b0b781bba2aa9d3b71f9788b667eb14e43d5c97ed7cfe459aa736c1f5d5e6760001b620073b9565b62006d3e7f7ab413b54fe1641cce7c27fcd40db3fdd5b77d0ad21460ffbb64037bb214094d60001b620073b9565b8960c0015167ffffffffffffffff168c6040015162006d5e919062010d02565b8160c001818152505062006d9c565b62006d9b7fd28446571037ce54ef250496f56e9e658b8c7c7a6a12976d8ec18a0f057477ff60001b620073b9565b5b62006dca7f043babd9bb9e0e9c76fa9f0c68cd587caad0e589b31503cd963ad0f16d26900d60001b620073b9565b62006df87fa2bce3ea1622cbebfa50eac412afc04f9407d906fc0fad49dc7b5f545461f98f60001b620073b9565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632384a6aa826040518263ffffffff1660e01b815260040162006e55919062010875565b600060405180830381600087803b15801562006e7057600080fd5b505af115801562006e85573d6000803e3d6000fd5b505050505062006ebf565b62006ebe7f0881b14f4cbac6a640f56904577740546c91fdd1fe5af3ade0caf3dcf1694f5860001b620073b9565b5b62006eed7fc73ef999d2b7020dc4f80340c44c5cce0afcb65bdc6b735017a66ffedf1e995660001b620073b9565b62006f1b7f52a98d01f32c4d1dce46f9b818475cd3ab23a96c883555e656257c492038ad9960001b620073b9565b8315620072d95762006f507fd22b1ecbf0339db49703a8bb58cfff0175cd5cc42df043e8d0ffdaf89ecbcf8f60001b620073b9565b62006f7e7f2373d625948d3e524d519c1c5900b4e36bd26cf8c09ec83c1fa6324754e3b3b360001b620073b9565b62006fac7f8d9c2a0120886a5aaa9f18eda54c9823dd299a10caa7075c36d9be66ad06c5c560001b620073b9565b60008b60a0015167ffffffffffffffff1614620072395762006ff17f18cf4d2d377b061f5aafb5d95206df70e96d7067365e0f26bb61d0c3831f3ed260001b620073b9565b6200701f7f3ca6f5846a74b648927d994564d30683a2d3322d0015a07f85e341beff83ce0360001b620073b9565b6200704d7fbed21c8f3ccc61094dd246c83b8a57c26155c838406e9cc88fa1905d7805cddd60001b620073b9565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d760405180604001604052808c60c0015173ffffffffffffffffffffffffffffffffffffffff1681526020018f60a0015167ffffffffffffffff168152506040518263ffffffff1660e01b8152600401620070e89190620108d4565b60006040518083038186803b1580156200710157600080fd5b505afa15801562007116573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906200714191906200f01d565b9050620071717f92a263e0f7fce99ce10b411c0d7b506178e3523f76b37f933e152d85c9795ef860001b620073b9565b6200719f7fc53207b556a82b61e8d55d45fafdff2d21c9ea09972cb69194220a29e8ca9dc060001b620073b9565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a3f02918828c6040518363ffffffff1660e01b8152600401620071fe92919062010899565b600060405180830381600087803b1580156200721957600080fd5b505af11580156200722e573d6000803e3d6000fd5b505050505062007268565b620072677fd02960f429ea5d0b6e78bc10e9014b4fc8c6dcb1b83126084f5fcc8bc23a443160001b620073b9565b5b620072967fa2ce1678e85ebf4ac1ff55de11cb693db4749dd55a1498f2c7e8de6a44361d5d60001b620073b9565b620072c47f2e3dd35d1e93dd2d12c9b81cd837b47a72457f8114a3eb391951813ea92accc660001b620073b9565b620072d38989848a8e62007b8f565b62007308565b620073077f8a288a30e53d223ec12f697555b41ffd362de3bbd559268cd7a8a615b74776ac60001b620073b9565b5b620073367f1c11be8c7833a9c29a0286dfeb1d79d08437862d3bf0f7ec99c3ae64ae98147f60001b620073b9565b620073647f5aca5d1a0c017caedfc17679100858029ff0b9ca168d390aeb2293f512518c6b60001b620073b9565b7fd936063ae3cbbf2dbcd2adc837a997ab97adb09f24c566aa79d101e88928dea66007438b600001518b60a00151604051620073a494939291906201058d565b60405180910390a15050505050505050505050565b50565b6060620073ec7f533833147fbceab0a04b24fcf808af5e521d9688e180cb7f8c19b3a5852bd4b560001b620073b9565b6200741a7ff28fd36563106efdc6d9fe2b1199afe1670dd325140464fa68fe4fe1788843df60001b620073b9565b620074487fb7c915a3ccfffb3e0f9913f0199c0f56e7b6ef28b0108788ad38f218e9f350c960001b620073b9565b60006005836040516200745c919062010434565b908152602001604051809103902090506200749a7f79dab4d53fc5137479b96022fbc3553c4f217b1b54d4a03ec3e1a708ca43e5e660001b620073b9565b620074c87f84dd572c893bc4666be08db74e72b794cf1cf3766b3f8787e93c0c5f0879bc0b60001b620073b9565b6000816002015467ffffffffffffffff8111156200750f577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156200754c57816020015b620075386200ccb7565b8152602001906001900390816200752e5790505b5090506200757d7f856149dbfbdc9ade8bfbbf465ca1de45f6a684bdab79b1ba141d1f15e915fa3b60001b620073b9565b620075ab7ffba6b8807d935dbfd35ea6e619646dd01b9badd36d38a7ff1347cf28b1bf4c2d60001b620073b9565b6000826002015414156200764e57620075e77f278bdf5abe17a911ce5505ac97c0fbfc83916f3eeeed2bc10f648485a38bc86e60001b620073b9565b620076157fa8055c7b7986afe080c0e939e12a285369d5ef65c99ccf96b50ec8f6a3ec921a60001b620073b9565b620076437f144a7a75ded974443a7b4001cee160d4bd0f61babb30f7fbc8bb44ed0a62826d60001b620073b9565b80925050506200789c565b6200767c7f6560b0d97e7169eede6ae4b747136736170e31e0f5314264332309eaf5e4e9f360001b620073b9565b620076aa7f0b32ed3c0e1f9da2f255bc41bdfea3a0049a2d5da59811d6a17f5502b1ef0ca260001b620073b9565b620076d87fe4d30d60a4489deb682edfb5fb66faf4a976d7f7546cad97a12b7bd79861c98d60001b620073b9565b6000620076e5836200c039565b90505b620076fd81846200c14990919063ffffffff16565b156200783957620077317f318899e60312fac39f1ea2bb39656137a3755e90caacb1e56f60128c07e6b68760001b620073b9565b6200775f7fe14529329d48899dd309a680836606b4f891d20feeb828e7a1565c3ba8b1a91060001b620073b9565b60006200777682856200c1e790919063ffffffff16565b915050620077a77fd1c782eeb1da7c6376a568b4d6174b9680865644c023d43eed15476c06227bb360001b620073b9565b620077d57fdb10f62e8da07d23c956a3ce17fc89806d429d2b212b90fcbe9c5241fafdf67960001b620073b9565b8083838151811062007810577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181905250506200783181846200c4a190919063ffffffff16565b9050620076e8565b50620078687f702641b0ac5eb6cfc0b282212a9a5d5a0b69c15ff52e3ec912c8bc066aef658460001b620073b9565b620078967fe226fab1292adf954a7556353498f89f470b960a169bacc49334be0a5c33174660001b620073b9565b80925050505b919050565b620078ab6200cd1e565b620078d97f07a0877b470595a1bed7f4069258b1aae748ee2be40e2b449368b886f16b083c60001b620073b9565b620079077f88de31fd45be0814fc625788152c43510d26f29f7f6d23b6e8ac628eb62034ed60001b620073b9565b620079357f95e170eb493c17563ecaeb89ad6b8f068747a80986d5af1c5a16634687e59da060001b620073b9565b600083836040516020016200794c92919062010404565b60405160208183030381529060405290506200798b7f22b7f4607c3c44ae18e9e36d7b17c52f1d087a4974c042a68381b6e9cccd881e60001b620073b9565b620079b97fa661b4f61ab7ec4753954ddf0338c8fa96c8b4028d3dab2de1ad9f23c65543d860001b620073b9565b600881604051620079cb91906201044d565b90815260200160405180910390206040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505091505092915050565b62007aad7f2e99d5282cbe9b41396ae6919676609262c8d8d0ca75d6bb408c72208b2240be60001b620073b9565b62007adb7fd989359f747afac85feeaa7593de470cc0b82d783d7dc01f21859671cc52fc1560001b620073b9565b62007b097fef12e34febd5dca298c9256a6863428bd118cb7cd8c75c6244dc38d4b830dcd360001b620073b9565b8060068360405162007b1c919062010434565b908152602001604051809103902060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050505050565b62007bbd7fb5d74f82321cfce1104b93ed3e67d01ebeb87c9b51b3fa8b6e59ab535290c67d60001b620073b9565b62007beb7feec7f9682bff4dff108b9daf13e5eda10aacb22d0ff25cca5e52dea7ab0ff7b060001b620073b9565b62007c197fc9d8d5b17f9e5ede8b4a493b7913daefba83d8e4bdcbabcbfb5668f81baced7460001b620073b9565b600062007c288685846200c686565b905062007c587fde96ca5b50fc244f9de7b52f733d5ed627e648134a1dbb22d178a8065e15627f60001b620073b9565b62007c867f850af0896df406f3d6347ddb631be8fafc943539234a94a6b611581a84939dae60001b620073b9565b8067ffffffffffffffff1686610140015167ffffffffffffffff16101562007d445762007cd67faedc44903731f7aaf0677b20130653760cbedb6787d6fe0faef7ec327aee701a60001b620073b9565b62007d047f706f21ce7f27f5fcdd51ed62f082eaa7b539fd849b1872175561412cd5f328d260001b620073b9565b60096040517fcd442d6500000000000000000000000000000000000000000000000000000000815260040162007d3b9190620106e6565b60405180910390fd5b62007d727f9822250f810fb5b5f0181f5698003e94f1e4261dbaa9fc2a0502a1f830f29fed60001b620073b9565b62007da07f95918a839dcca798859f7ac840c513022c7882bbbacc0b7fd7865c73705d1aee60001b620073b9565b62007dce7f3599049c8a9afd14d3412e4e4ec718b86a8d625c07193e09a05441b905e5627660001b620073b9565b808560600181815162007de2919062010d5f565b91509067ffffffffffffffff16908167ffffffffffffffff168152505062007e2d7f369dca1a76d0fb0ce2103d50a14645fa678c89e7576bcea9cff8d9bb670c0e9a60001b620073b9565b62007e5b7fb1c6669e50a07b8604c911f0feaf86de1e6fe2344873e8ee4ba0c0bfcba3a11b60001b620073b9565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c819d86c866040518263ffffffff1660e01b815260040162007eb8919062010817565b600060405180830381600087803b15801562007ed357600080fd5b505af115801562007ee8573d6000803e3d6000fd5b5050505062007f1a7f78e02b59f2ea02684b65bc60a92bc56df855f05f17097eda82f6639f9451198160001b620073b9565b62007f487fcd0712d85369c9305432746f10c2c03de35b8a5389e9819d9f9abe532e8c091860001b620073b9565b80866101400181815162007f5d919062010e60565b91509067ffffffffffffffff16908167ffffffffffffffff168152505062007fa87f85c6d6306b8108696e412fb0c3e3ed7de5eb00ccf861586f7dece7fb4d6f231760001b620073b9565b62007fd67f7949420d4b9cd606ca80380d4b8836b6700ebb3a031138a5a06c3c46aacb517e60001b620073b9565b6000866101c0019015159081151581525050620080167f823225b88e1d8be44aa625f4591b82efaa5557452235edb8a649b5bf0947050b60001b620073b9565b620080447f8563f60887e0eb19b2e80cd69cda512409bb735cae8ec95db3fe3020b30e7e2c60001b620073b9565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631aa05c5c876040518263ffffffff1660e01b8152600401620080a191906201078b565b600060405180830381600087803b158015620080bc57600080fd5b505af1158015620080d1573d6000803e3d6000fd5b50505050620081037fe086c67b268fc789036c861973d6db250ffdc70367540f3c9f0122dd0bb369fe60001b620073b9565b620081317f39b14fbd207100409c141b8e627278b2fcafb3af1836acfdb54243552ed9549f60001b620073b9565b6000620081617f535f0afff06828b9a3bad9e681ecfc4ab48e44c7f9f395e2a2927a4223cd4f3260001b620073b9565b6200818f7f88786ca25b0a6d286633981808c0c81bc6b1afee0e5838ffa4238dc3d18ea54a60001b620073b9565b60005b8451811015620082fb57620081ca7fb37ba7f94497d003f1e72744d3d47539199e0fb1ce79080ceda18d61dbd29a7b60001b620073b9565b620081f87fd0a6368437b90ac510f422c884b3ba20e52583613c0fee50cae8c888704f521a60001b620073b9565b84818151811062008232577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516080015115620082b657620082727f9016cb6926b44ffbcc59260b6a43dfbf2dbb9d69f3a713491d4e3b91195218ae60001b620073b9565b620082a07f42b294934692bf93f57fed9519ec7df6e6cf62ce4c9a2d831fca55eab5025e6160001b620073b9565b8180620082ad9062011183565b925050620082e5565b620082e47f0177dec289b8f2526afd362eec9db354328c162e05fa621b34dee9c4d3af567f60001b620073b9565b5b8080620082f29062011135565b91505062008192565b506200832a7f4dbdf830786377b5a61c71fe8f6dd1a173ffc7bf39e77f2252ad671ac7f2a75a60001b620073b9565b620083587f046253b7fae8085fd95fad7084326abf59e3567d74b4954bfd785533b453e65860001b620073b9565b60018167ffffffffffffffff16141562008493576200839a7fda4f3826a42bc301260d211b53b113c04a11dbf8e2d44f6fb9f8178af17dfdac60001b620073b9565b620083c87f39c0002d2ba1a4cd8083ea52428316a0043c7b24a453dabd0f033b6838cec0b560001b620073b9565b620083f67fa1dd33eda6387973352feea6a1bcd4086509ba2380bd627442dc858c1c65822960001b620073b9565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166334afd8c188600060016040518463ffffffff1660e01b81526004016200845993929190620107af565b600060405180830381600087803b1580156200847457600080fd5b505af115801562008489573d6000803e3d6000fd5b50505050620084c2565b620084c17fd906e0e68b5ec7fb4600d02287d875391838731415febbd92221a9251fffb32960001b620073b9565b5b620084f07f116f6de949e83736e71fa0cd1b626c5f5a0c492b17b4d2df41efce8490639ff860001b620073b9565b6200851e7f536794cbd9b840eb0a0422b8b9f269becb308dc1d96f806c95d1caf314dba8af60001b620073b9565b600187610120015162008532919062010d5f565b67ffffffffffffffff168167ffffffffffffffff16141562008804576200857c7fb7e66e972fc516e98db148ac6b928dfa155a046c596910f4ccebd7b1e9034b6460001b620073b9565b620085aa7f09d0637848afd1783ac77a56716cf0d8d54d334e6623790ddd398e17d212790360001b620073b9565b620085d87fbe44f4e8ee306bb4efc4b36e17b6ea80aa59bc72308089baf49a8d575007944b60001b620073b9565b600087610140015167ffffffffffffffff161115620086dc576200861f7f464afa7cc4a013c18352eac9c6ee3776276ead8d95c7070cf47c91b0ff53636d60001b620073b9565b6200864d7f505e1a8e4d383caa6cbe518fb40f3aaa80abf4a5e6fa027cdacdd0312e247dc960001b620073b9565b6200867b7f43d49191d7a47b166ac4dfcabb8850de88eebc0baa14ec066293818a985bf47760001b620073b9565b866020015173ffffffffffffffffffffffffffffffffffffffff166108fc88610140015167ffffffffffffffff169081150290604051600060405180830381858888f19350505050158015620086d5573d6000803e3d6000fd5b506200870b565b6200870a7f5d4026ae0d46f72e9b0d9ea63064be789031a8a4fb19a2e33a3961e39e10787a60001b620073b9565b5b620087397fbd4f89b98aea14b10997bfe2b2489e4a110dc2a339161222e454f9b60f9e042460001b620073b9565b620087677f95022b7fc316ef2c8520d78abbf1367f539b72bb19f83b60b711f79c3a96c7e960001b620073b9565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166334afd8c188600160006040518463ffffffff1660e01b8152600401620087ca93929190620107af565b600060405180830381600087803b158015620087e557600080fd5b505af1158015620087fa573d6000803e3d6000fd5b505050506200892a565b620088327fada2b0058b797e953714e1501f65a7084987774e4c2c5ea20999a752017c3aa260001b620073b9565b620088607f2ce826fd2805531f69a6f088261036d17b686419a474b4474a9fb741d1b6f88d60001b620073b9565b6200888e7f3a2a0611813370ab03dd50c6d884d5c10c0d34fcb44132043eb846e20b88aa7160001b620073b9565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166334fddaac886020015189600001516040518363ffffffff1660e01b8152600401620088f592919062010483565b600060405180830381600087803b1580156200891057600080fd5b505af115801562008925573d6000803e3d6000fd5b505050505b620089587fe1942f2043a44bee3a7b91e647909cf54f99a7303b380ff65250c2ec209c106e60001b620073b9565b620089867f2da9621272571da9356bf8c908f919ee9ea5d8497e4dab18dbc910bb62e713a360001b620073b9565b7f123b9998b2f027b02db4cb2eadef421fe9f1c21627569438588262d3a6baac6c6006438860a0015185604051620089c2949392919062010540565b60405180910390a150505050505050565b62008a017fa6520bf8905555f96f501407c95e497772cc9b950fdad2d1629804f22d26683b60001b620073b9565b62008a2f7f7bf28374ed8a80f7bf6e499592c5096014e36c4a5b5039f3be1b6986164e42f460001b620073b9565b62008a5d7f58ef7815c29ece677717aeae48e3d185f910ebf43b34076b54074b904a047ca460001b620073b9565b600062008a6b83866200c89b565b8262008a78919062010ddc565b905062008aa87f8ef6f8f1a61283a5dc772d4207fe2a7dca9244113b7793d4b8ee93978412a7b960001b620073b9565b62008ad67fa1d50950a2fc0b61db43fbbe13fd3d5bc24a5e7bdcf652b92fd7ac871c746cd860001b620073b9565b8067ffffffffffffffff16846000015167ffffffffffffffff161062008bb75762008b247fc986e9ceee21fcae08b852797474834920616b5b5a34d75daf4ff2d996ef199760001b620073b9565b62008b527f3d648d98081a28c662806c1d60075156710ae9c776c78095692b00876546dbef60001b620073b9565b62008b807ff3405b84f112f278f637612ef29278744afa095b1c442160da473e152a71547e60001b620073b9565b808460000181815162008b94919062010e60565b91509067ffffffffffffffff16908167ffffffffffffffff168152505062008cc6565b62008be57fd5053904a73dbeef0b4942f867723551bfb37e4d0c7df651bef2c4732653330c60001b620073b9565b62008c137fbdca501af8a0a6ad9d4028c7bca3d9eacc83fc2cffb2262df5b8d8f0ed25194b60001b620073b9565b62008c417f9f10345e5f43458e57a947ca6b83559efa265af8f4fc590e330a25afdc57a26d60001b620073b9565b6000846000019067ffffffffffffffff16908167ffffffffffffffff168152505062008c907f48df4fc357928b767793e678821caa8d1bff6d60a6cfd8356e6f1c43eae73dfe60001b620073b9565b62008cbe7fd196f59e1f7debf759858b6133745fbe8b0ef083d3732ae70816c38d8fdb9a8560001b620073b9565b836000015190505b62008cf47f57690fbf70fe6a4270c39626a9d12abba54408c16c5a28f5850c83e3f4c20c7b60001b620073b9565b62008d227fab2a24b3553aa2af51df49b80913d17a13de2367cf1292f6c69fd1a73e3d1c9960001b620073b9565b60008167ffffffffffffffff16111562008f5f5762008d647fd667bd8aa6e78b4bcb94f58282d05ebc225155dc32258d7c025d7a2f74c13d2960001b620073b9565b62008d927f6ec5c30b4fc512ffc63febfce60ce11cdb4593c8da489eb82f839ef5d39c9fcc60001b620073b9565b62008dc07f58bdeae4e12a1931b4c32b12443d5afa603c5c7ee7b9b3fbea255b22db864e8e60001b620073b9565b62008dee7f39fce1ca179ec1983a61703eec8edc41f159f563d7e052baf082c8c56704c70260001b620073b9565b8067ffffffffffffffff1634101562008e3e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162008e359062010747565b60405180910390fd5b62008e6c7f0b8768cb94c7135128d0dbb31ba32317a8b1129cc242d3dbac6cd625bafa7e0660001b620073b9565b62008e9a7faf0cd1c18a1f9dda1ed0c761220fb00a15c1fb1903a46f67d2a1c23a1af5014560001b620073b9565b62008ec87f309ef9531f36185dff0010ad240f0b4bc8feaf74663f6c9977bdacea648afdd060001b620073b9565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c819d86c856040518263ffffffff1660e01b815260040162008f25919062010817565b600060405180830381600087803b15801562008f4057600080fd5b505af115801562008f55573d6000803e3d6000fd5b5050505062008f8e565b62008f8d7f2914caac75c34be6795e751888733db9698ab0fc0a37b20234bd7e01cd29749d60001b620073b9565b5b62008fbc7f3c22dc82e5ea8515b58890a6fe4ec7bbd13405ff97ac5ef8a51065d89ac0635160001b620073b9565b62008fea7f065bfff9efcb6a0a0d4283dedc1569c157384b51ba5a42275f087166d194adff60001b620073b9565b62008fff856000015186602001514362003247565b5050505050565b620090a4838383604051602401620090219392919062010703565b6040516020818303038152906040527f969cdd03000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506200ca11565b505050565b6000620090d97f99b41c0504365ec000ec5664522d04aab39f93c70b97ac79df898c48b488d81560001b620073b9565b620091077f889aad19aa8b2f12547bd53701e9300e71ee7bc4427fad8e29a6775a3366552b60001b620073b9565b620091357f98acb880014096c45dd184cb1805ba6f0185415f45b8cd0fd6480c38ef09336f60001b620073b9565b60005b84610100015167ffffffffffffffff1681101562009edf576200917e7fc8aae2e013c780d288b51dd0b66c48f14940cb5640151bee18e5495beb8b662860001b620073b9565b620091ac7f9fdc6bd6749635a1bd7b994977e872a7a52b31e50e19468847eaad914b76040260001b620073b9565b60008561016001518281518110620091ed577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101519050620092257fae3cec05bd1fab36d190e64821b4b4667235c73277434d68cdd5d59cebaec8e660001b620073b9565b620092537f0ea378c22593b075a1553e77c583adf098113802e1e7c0eccaea6167d17d5d9160001b620073b9565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663defce5d4836040518263ffffffff1660e01b8152600401620092b291906201051c565b60006040518083038186803b158015620092cb57600080fd5b505afa158015620092e0573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906200930b91906200ee3e565b90506200933b7f70c4fdbb6e4d725c55cf48302213ea152a8249f9022285d2a460c6532dedcd0f60001b620073b9565b620093697f417622f68f162bcaf709d5670d9bd17491ece2d5bd8116aa4992094e494212c160001b620073b9565b60006200937a8260000151620073bc565b9050620093aa7f02643bddce67eb1cc54d7af9cde5aba0363f56b003c9d642e2a4f9d1d0fa862460001b620073b9565b620093d87f984d28118f6c760be0f94ce2a7c45279e1790c546357978b904d56253386363460001b620073b9565b6000620094087ffc72fb01c32146750a854c3646e45eea74dd146a3bd7937a897deee2420dca8e60001b620073b9565b620094367f27b6c5c5e6e1f53b1ebf68ff816a859d1830f9e54a90d3bf41a6462611a98ee860001b620073b9565b600083610100015190506200946e7fb58600b74bdf43479be771e686231ea9a75e2337f75a1131f8a75e349415ec7860001b620073b9565b6200949c7f8c9af00630ac40f2d84f7c5c732b9a73a6dece5b8b165a160594486978ea12c960001b620073b9565b6000620094cc7fed80116c3927465935ae23138ecb80f4bc4e0bbbb4d2a94a99e262409a76e57960001b620073b9565b620094fa7f4690e51229e5d2de6165d161580ffce0455403dcfb32984548987fc613cad28360001b620073b9565b620095046200ccb7565b620095327f099de8c4bfce0e0d1bed7d4b04bbf611f0ef75587a2746cdbe90d5547f3090b360001b620073b9565b620095607f9d14d9f0a0e742071225cf533990a5717c3bc708c0097d3533606177d0d4bcb860001b620073b9565b60005b855181101562009a65576200959b7f51381c4c3a12663aeeacb609607249f71575ad991b5ea4d31f013dc2b859186760001b620073b9565b620095c97fafd347aee829f2ba5a4eb2e589cdb253c1f7db48c5fad8d5aa5e58774c9dfbff60001b620073b9565b8c6000015173ffffffffffffffffffffffffffffffffffffffff16868a815181106200961e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff16141562009a2157620096757f512457dc556dce832118a6ebfcac21e75a9b37101b8d9662eee35a7117cf4ec460001b620073b9565b620096a37f550ef2fb7c49fc6d77cbb8b123eb957697564572ab6590bb86bb1bf390b785b560001b620073b9565b620096d17f3f83e125c5326eb100196d2919e8cf2fd9387f7a020320286028121c4d6c9a6660001b620073b9565b60019250620097037f2a44b86db298a125303398a7c9f6247ec8d63b57eba0f09ad6e45f3c7d6985e560001b620073b9565b620097317f5d6e82547a1f04db5eee28d104471e90988c714f638279121241e3bc845074d460001b620073b9565b8589815181106200976b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101519150620097a37f69950c9a4b665e9c725182ed38d7a21fd7de7539d4a42989d4da0c40a6d831c960001b620073b9565b620097d17f104fb3edcd520a754865432d4fc337660c1ba85ac865acd1819cb8950eefa39560001b620073b9565b600082604001519050620098087f5eaeca7f2a1742f3093024cb0e303f902f8df3f330bcae237998f355ef66706060001b620073b9565b620098367ff91aebc4718dba927c7467b1d042f76381450d86f8933b8b8da03c4078a3eae160001b620073b9565b8760e0015167ffffffffffffffff168167ffffffffffffffff1614806200985c57508443115b156200996357620098907f36250df4dd8b183a7fb1b421c5c3b663d2ea970d2a334ca1351cb1fa206a742660001b620073b9565b620098be7faca70f6d17ebb05e61824afc24478df44e73a278c4c4b65b2e669c76ec88412d60001b620073b9565b620098ec7fb04d4fb548cc63eed460ea30fd245295d36db2096d4c01969a37b7efc8d6ea6660001b620073b9565b60018360800190151590811515815250506200992b7fde26ac31f03b73db5828c190c68bca0be18bb6e1f1136dfe37be7dd39e3c331160001b620073b9565b620099597f7f5349473afd2977e2d77d957f818ecfb27c0e56efb85cbe37f948104d5db24b60001b620073b9565b6001955062009992565b620099917ff917801a0e6c0c648c006cc4b9f9157d81fc37e93563c376dcd92fdbd7cef9c960001b620073b9565b5b620099c07fdb723c2bff791c15de5abf1f3740f39638634c2a492b23456ecacd87fa7728db60001b620073b9565b826040018051809190620099d49062011183565b67ffffffffffffffff1667ffffffffffffffff168152505062009a1a7f141104dea343b265cfc021c46cd8e3f763013c2649cbb1fe60df109c5617c31960001b620073b9565b5062009a65565b62009a4f7f32afae8a03e621772dd9d046dbdfb399ffe862f15b86a602ffb8a6c343c7e36860001b620073b9565b808062009a5c9062011135565b91505062009563565b5062009a947fa7a1b85f95d5a698d15ad8716b7bb26c27b15111c5167736b7e6e3b33cf7b64060001b620073b9565b62009ac27fcd9d34a640bfb0b7c6e4431dff94853584f69087123a16e689e8a1ee6b503e1460001b620073b9565b8162009b645762009af67f64463b3609d561de03d551e0013ce32684a254644fb78cfc13b56508195f401a60001b620073b9565b62009b247faa198af1e59bcf78448189ddcf19d88221575de4199d505c24b2bffb5eb3704760001b620073b9565b62009b527f4d62b62944ed4e74438e5a6d9a1578e21f76e2a1db10c06300db7ab3797a8e1760001b620073b9565b60009850505050505050505062009f41565b62009b927f9d16bf398be3c245c0af14b0f8c3f0dfac777bd7b36f90e4c98af3bcacc0683d60001b620073b9565b62009bc07f82a68b8708234b8f61118a41d1038e5cd53cfc5128105d12bd2bced2fb3bf18c60001b620073b9565b62009bee7fc9d843b4d545a896b03b19af8e972cc658ace049573369bf4d62fe51cac8281a60001b620073b9565b62009bfe86600001518662002ff9565b62009c2c7ffa1152683f5955bea192dcb4d69b10c67aa843dfa6238922c7367e131572efad60001b620073b9565b62009c5a7f24b0161cf09012b815d11e678961ceb9b420acf3b018748ed8c2a1bb42a5453560001b620073b9565b831562009e935762009c8f7fd8ce2e292fda381d0e0e39505ff235bbac68f353a5f259ca334047a453b2bb2e60001b620073b9565b62009cbd7f01fdc3b429b0d1c9efcd1684b3b2f7b1a7446e63e5ade377fbbc9fb8338d2c9f60001b620073b9565b62009ceb7fd7b8ca8b74eaa6f2356e3d495f3c4fa7b1d32be71cf6fc3089f375dd5e06ba7b60001b620073b9565b62009cfa868c83888e62007b8f565b62009d287f71907692e79e3b7bd3b17b23da102c4ffa8d754acb6cee4620838d3ec2e56b7360001b620073b9565b62009d567f15ebb0ccbb6b038eba119ac6fbd8c68f20a93d4b5d265dd9bc66adc9c4bdbf8260001b620073b9565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a3f029188d886040518363ffffffff1660e01b815260040162009db592919062010899565b600060405180830381600087803b15801562009dd057600080fd5b505af115801562009de5573d6000803e3d6000fd5b5050505062009e177f07ee16e4afee4f842416a060e70ed891a26243278cad8bcc06426ce663830f1960001b620073b9565b62009e457fb33a70fb115b11df890b074f5b502b99105c9f3fe4fed35a1a462955e2b21ab660001b620073b9565b7fb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f4560014388600001518e60a0015160405162009e8594939291906201058d565b60405180910390a162009ec2565b62009ec17f7b3fe115c3901837c158d7fbdeba0486e080062640db7e949964db7c8365850460001b620073b9565b5b50505050505050808062009ed69062011135565b91505062009138565b5062009f0e7fe3b123bac2e24150c40dfe232a196aeff1794cde9ae4946b1be376c73c69587b60001b620073b9565b62009f3c7fa71188729b2cbb8ce9a6e1b3ec1fda884b14f32d3e96ec347b846db9cb80080260001b620073b9565b600190505b9392505050565b600062009f55306200ca3a565b15905090565b600062009f8b7f8f3afc474225eabb087904841767e9529d25f82ddc7692a65d57d76c22bfab9d60001b6200ca4d565b62009fb97f339d7439d04ad42028b51fce3ca43aa4bb8ccd46f70ad48425d5317271036fb760001b6200ca4d565b62009fe77f9681274581b7f3a8dead554f2827be1f746bab016e4d4fdf188502db1a79c42b60001b6200ca4d565b60008460000160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015490506200a05e7f10728c320fe944762b3bb1224ce8aa85b6d64fc58171b963a0321306f3e585f660001b6200ca4d565b6200a08c7f5f369b3449a06d504798e3eca430e10c4bef40b1ebc5ea7b7b8309c85ae7d2c960001b6200ca4d565b828560000160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506060820151816002015560808201518160030160006101000a81548160ff0219169083151502179055509050506200a1e97fa06a3df516226113f9b1252d6e8c465ecce0ae398fb66ba703cb0fde35b168a960001b6200ca4d565b6200a2177fdefe07d2c841614cb4cbce63eef482214ff3c78b1578c4db49990e709dc2d7c360001b6200ca4d565b60008111156200a288576200a24f7fba2f7cd8e41ebf3f9cfd7822ccb14c0eadd1c9852e2860535bdd759ddab1fe8e60001b6200ca4d565b6200a27d7fd00fe8735280698c4c1fa4b28a191f63c9e57d44d44475a9e2c416ddd5324f8a60001b6200ca4d565b60019150506200a5db565b6200a2b67fef269b8596545a59484fd3ada4416e3ba760bd28aad340e2a933fa20a8e7691e60001b6200ca4d565b6200a2e47fa7d755d8621c0bce51cff342ab984211e52a2c183be44d7379b1cf441c417b6460001b6200ca4d565b6200a3127f31c43eae29d5eebb3f32b7d9eb5ac3a8734f36a68e4f60995793ccc4c562accc60001b6200ca4d565b846001018054905090506200a34a7fb19c275b45bfd54b2579af124c7e78edfe6603aa97f503db90366ed9435f744060001b6200ca4d565b6200a3787fe302f187a39a5f1895bfaa159149dd7f4073c85c92ab601488b230409179525b60001b6200ca4d565b846001016001816001815401808255809150500390600052602060002050506200a3c57f5b4400f4a9aeba42aa289d593e52c043f1cce28ecde51fd68e49fcfec94e261f60001b6200ca4d565b6200a3f37f8aa0ba5239a2d1533d8c90cfde5f62ac5a000d08a24ced99dd73f589aeb02e4e60001b6200ca4d565b6001816200a402919062010d02565b8560000160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001819055506200a4787f75c091f3d90fb99926e297742cddb8e1b433d7a4e3a4e2db06c66e1b3748b43f60001b6200ca4d565b6200a4a67faac2e5448d88c03f88d402c10a53a4d3bf214673871561ef04ba7e5c1231c4af60001b6200ca4d565b838560010182815481106200a4e4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506200a55d7fd31d5c2af99c0a6b0fcd184aca90824ba588b7a6014d2572303776a7cce8aad560001b6200ca4d565b8460020160008154809291906200a5749062011135565b91905055506200a5a77f7c6f6b476928224a628fd4474e5f4129b76364caccee7ab3441d930fec45db7460001b6200ca4d565b6200a5d57ffd08eb3980dfc76114a88e6e71daebcbe38e249a00605e46fe4351fbc677a2fc60001b6200ca4d565b60009150505b9392505050565b60006200a6127fa4ead4f076428a6805299c709c19335fcd6572944b121309962e07c6a17ad42d60001b620073b9565b6200a6407fa23155b61a1badc1d3f258acf9d4d5295ae44690e66b6ac3bdbee588953723dd60001b620073b9565b6200a66e7f9f4ca2b6087b74347d17ed24971e6c29a0bfe1557017f7396d029786f37fa86060001b620073b9565b600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002054905092915050565b60006200a70b7f58e91497643c1580ba8fce5f7cacf6d51f8322a10bb38835a7e220464ca385d760001b620073b9565b6200a7397fed295cc3993b9662909eb75fcbd087d3cbc84a36f674ebb051a0345d90b436d060001b620073b9565b6200a7677f68fb92549ba302f5f6d68f9db6db74dc94cb9969a0bd4104ae3e03af4e57d05a60001b620073b9565b60008460c0015190506200a79e7fec2a4853956d1a4ceee7d97ae5465e3ad9ca72d081ccd813eddd65e3b5751db260001b620073b9565b6200a7cc7f3ef1f21028d4433515328e55c6fa6a05e1183951db140f5dcd72436c80c32b1960001b620073b9565b60008660c0015190506200a8037f553018eedbae47833e54f385743afe14888cadbb2c126e65919fc54affaf3aff60001b620073b9565b6200a8317fa7575b482d181c385c2ba39b309484c6f7cb6d29b454ce02f0e316c06e4dbd4560001b620073b9565b838267ffffffffffffffff16826200a84a919062010d02565b106200a8e6576200a87e7f4debc683fe29d3539159ef3f14990d7b13452b8fedd8380d24cb85ffaff1cea660001b620073b9565b6200a8ac7f884af1070517ac2ed45f38baad5dda9644a296a30b27eb93f2524e08ed4ec1d560001b620073b9565b6200a8da7fbcfc721eecabde701227005b095dcd00eff448be4cdf428465f3dac8976a722060001b620073b9565b6000925050506200ae1d565b6200a9147f115c79494652cfe53c817163dbc759d0600f09382ed80b9107170bc45dfa858060001b620073b9565b6200a9427f16efee67891375a520960d2e92e20fa0017d6dec9a763524657631589ce43bd960001b620073b9565b6200a9707fdb2472675f915eb6e8b0a0de1cc84a3ffef954c1e44773805f219b0c3e9f4ac760001b620073b9565b60008282866200a981919062010e25565b6200a98d919062010da4565b90506200a9bd7f87160bfe64d1541eab8e5a68fb1642ad5c0171073914851607ece21840e1706160001b620073b9565b6200a9eb7f63922a29142a0358345154a478cd6a8faf4f6f8d0270895d787fc3b1a190119660001b620073b9565b60006200aa1b7fabd395443da03e3ad8ede4513f64207d653102e79d504e73a7353fea456ccb6860001b620073b9565b6200aa497fe680dfd37d93155d4e2d4b4b6ce1e34e38c136db1327b57d7fdd61acfd394c3460001b620073b9565b600087146200ac3e576200aa807f39ec3013285af66a7ade83e184f95b09974504dcf49e2d5f7533f949d773f35b60001b620073b9565b6200aaae7f59bae41e7fe4b67d67cedb9d7b8acda2ff75d9124e5bddfb821572e03e13234760001b620073b9565b6200aadc7ff70df2859d958d2321d2ac45f874391ef1d5354bc8fade31ddc10a257a7acf6960001b620073b9565b8367ffffffffffffffff16836200aaf4919062010d02565b8711156200aba9576200ab2a7f7c1ae9cb9a5ecb22440cb57ad0521926f6e09bfae9690cc35207185ee3f87fc360001b620073b9565b6200ab587f4fae10ebcf9c22362664d1dfc08028b17b455067af9b995c7ba89d4ad6fd3de960001b620073b9565b6200ab867f3df189fbd98bae753489cc676e6659e5e4bd24ef2694c3d043bd642d3f49dced60001b620073b9565b8383886200ab95919062010e25565b6200aba1919062010da4565b90506200ac38565b6200abd77f8e246c1a755bd75fa2389c13d689e4f333a653339650497fd1ae124516c836ee60001b620073b9565b6200ac057fe8079187a836be3b4d673661217f7b40d0d271047b80b255cacc779057839ac660001b620073b9565b6200ac337fa727f5f7e433573bb315c0a95e45a5b04deea9529f7e4a52dd783dd226ffd28c60001b620073b9565b600090505b6200ac6d565b6200ac6c7fb9315226d90b978ea073c4fb13ab2408fd8985c41a539a47f889e195146166de60001b620073b9565b5b6200ac9b7fafc3c1a429be5d65f92af8fbfbd9456d3a50b185789ea899c5c5caa133ea3efa60001b620073b9565b6200acc97f6f8e87283b68f7c4aceb61f61d230e58f3dc98388f9bdbcad567440fff30b66860001b620073b9565b8067ffffffffffffffff168267ffffffffffffffff1610156200ad7e576200ad147fb3c57dd538fc693723a0af43aae7ca1407d3911019480fd68f4dedd3b5ecbea060001b620073b9565b6200ad427f1313b83a0a078ec69b8d409f500166414232bb9463fb593d506f24a28661eac860001b620073b9565b6200ad707f777a002212a0e0c28b98e17730cd2fc8eb83ab2192ee5510df5a68c9af4bb38360001b620073b9565b60009450505050506200ae1d565b6200adac7f416dd1b5579a46cf1512a1c88f464c91a1b1edc81ae6af3304301c4e0a15093160001b620073b9565b6200adda7ff596eefafbc73f6bd273a3984e825a5a17bb25be097eb0844ca8b86c0aa1a05f60001b620073b9565b6200ae087fb9e1d7efa4b9aad009188c4672fbd62cef391172c4b1cd535cf72706255659f160001b620073b9565b80826200ae16919062010e60565b9450505050505b949350505050565b60606200ae557f4ff685308f415e073f1b4afacccb2d5c984eafdf210e4dc8e5d5d539189eebfe60001b620073b9565b6200ae837fa6de4e2fbca443dd7a19e6a98ee53b5e9aa4eb794feb4cb1b2bd4d6af870bc5860001b620073b9565b6200aeb17fecceaebcda79444288bfb34d12365eb3523f0c40d48361ab5cf5221bb1e496ab60001b620073b9565b60006200aebe83620073bc565b90506200aeee7f82ff7e7079e511bff144d646f26b9ae8af1bda6d3b185962f6a7578473e6f89360001b620073b9565b6200af1c7f32ca81f6229f93f82b8ffadd1fc84c5cbcf31dd5808ce03001bc019bb5d2143360001b620073b9565b60005b81518110156200b16f576200af577f9158d9aabe2e8c73bf234598aae1bce7419b34d1d83dff4f864baa7911387c2760001b620073b9565b6200af857f7b31e6448150e378fd424f78eeed9d0e8226b62145d2463f33d115e5b9664e1660001b620073b9565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a65374a8484815181106200afff577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151602001516040518263ffffffff1660e01b81526004016200b029919062010466565b60e06040518083038186803b1580156200b04257600080fd5b505afa1580156200b057573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200b07d91906200ef80565b90506200b0ad7febf274387122f6e647dd3e8f12d50be88a139515909a6133eee62b19d322072f60001b620073b9565b6200b0db7f7b39fbb6c2bbe3c19b89e4e27fc187a46d865c2762683f1c348ffcbe68dbeeff60001b620073b9565b8060c001518383815181106200b11a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505080806200b1669062011135565b9150506200af1f565b506200b19e7ffd9e830571f8761de93e3cb19849f06163fb254427bc9eeeb432377ee64113d060001b620073b9565b6200b1cc7f7ab79d94dabee964058a9156f66230c3ecc33e58410c344bc875694b72b9763160001b620073b9565b80915050919050565b60006200b2057fe372bd25dc7f9e6fa454a37bc1a960ecd5969c34887feddc6e335c15a88aa66860001b620073b9565b6200b2337f15121c0b850669f94a78525cd1b5a71bc16d76eeffcf26c9776dfe842a75f55760001b620073b9565b6200b2617f269d7965001c76d5597794a39134b6dfd6db0b6c520d21bdea5d596d077715f760001b620073b9565b60004390506200b2947f65c939eb1e3eab9f9ef7dac2863f6e5ea4441a821470557e9211e2eaf960c20c60001b620073b9565b6200b2c27fb9fe5d4f938de3d1e9c42c65b572ee1dce7f1331160ec9d9a2291e299465264e60001b620073b9565b8260c0015167ffffffffffffffff16816200b2de919062010d02565b846040015111806200b30e5750808360c0015167ffffffffffffffff1685604001516200b30c919062010d02565b105b156200b3a9576200b3427f9127c5d15a7e2978b988b5c76178781fa2462b214fffd991aa8846e1573bed1260001b620073b9565b6200b3707fbfdf8ac41f529f29541e8a14089d58e6c9b9518eeab096ce53b3e1ba9471033560001b620073b9565b6200b39e7f15655164e6a675a4d407e61760e80d4d321d47852e73baca1fcd4b8ed45bca9a60001b620073b9565b60009150506200be76565b6200b3d77f6bf238f3442036d704b7977b2097f0f864574275e838b7f5710a7d9efbc9e28460001b620073b9565b6200b4057fe69f0c008700781b84f176d21b4b5ca692cfeff2734b280212f3c640caaac9ca60001b620073b9565b6200b4337f0ed1dac139e052371f8a59c72b6517c9d7c0ccab730f810767bcc1234eef84b660001b620073b9565b6200b43d6200cd5f565b6200b46b7f168e26c5a2db2ebdd536541a6065fd1d5d232d76164b19c1f9de70812a61f2a860001b620073b9565b6200b4997f70f8460e0df1e662bdd3bb161fa62c91eaa4f3869d0b6c3c18315ca4d551b6f960001b620073b9565b6200b4a36200cd79565b6200b4d17f428e1bcaf18ec42118c981ece3690cc01c78ce3b8cde95b33e4710cd3de11f2860001b620073b9565b6200b4ff7f447e2a3f9e6091f2e06b7a7a44bb6c67c1ea14ee04ac9fc6a4e6883e0ef1d98e60001b620073b9565b60606200b52f7f37d1caacf35ea27e0640500ee71b009bf02398ad7ed06660ca087c41f6950fbf60001b620073b9565b6200b55d7f6b9c3c081d57f84e9d92be63a74710806732ac28519f4a2ce6616ed4cbf7987560001b620073b9565b6200b5676200cb06565b6200b5957f3050161797bfe99a59b154f918a93d7afa3f9c0196bac0bbf238264a383ae61f60001b620073b9565b6200b5c37f4e67a1a494c30a7969749327b462d4aef97c7499e524d19e44ab2307be60bc2f60001b620073b9565b8760600151816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506200b62d7ff06e594d70a43e13e0662eda226446f76525970e3e075f4d61aa0ca8825b1e0b60001b620073b9565b6200b65b7fb7a771bacad19bf13eec124b19f32ccc8d2a89bb868ec47132dee8963470617a60001b620073b9565b8181602001819052506200b6927f0a61728aa13ff8c9161f312c8d0c285c4f9e24c01624a4b632e3713b6ba7b75160001b620073b9565b6200b6c07f6ad0207da14cab32652c3cc3615e044385d6262afc32f29ee463d4eb5d11effa60001b620073b9565b8660800151816040019067ffffffffffffffff16908167ffffffffffffffff16815250506200b7127fcf50f72305df6a8cbd1f8b829fd8f89a06959167862fe93886bfafe8471672ea60001b620073b9565b6200b7407f2877180e16a1e3a2a31c9fe4361110200531f46bab258e0463d68a2da88cab8560001b620073b9565b866101800151816060019067ffffffffffffffff16908167ffffffffffffffff16815250506200b7937fed92ef99784f600c0a28bc427136a55a1bb53aa178da38ae135312742fb2a43460001b620073b9565b6200b7c17f77a4b74e3dbc6038ee5eafad46df982db96edf66105f7facf14609869b42ef8360001b620073b9565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fd7c9808836040518263ffffffff1660e01b81526004016200b8209190620107f3565b60006040518083038186803b1580156200b83957600080fd5b505afa1580156200b84e573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906200b87991906200ec0d565b90506200b8a97f22a1c1fdc2d9d141f1ded49c3f41c0e59dd06731a130319938f252c1c96c32b960001b620073b9565b6200b8d77f5cebc0aeb10717c4e10ffa3996f9df7b073129404c38ee669681d8cf85fa8df460001b620073b9565b6200b8e16200cc1a565b6200b90f7f4ab226050d857943558606a4011a47904bd7d94b3b51ff34c71d96b81c9b4ac960001b620073b9565b6200b93d7f990a4b325c824bda58530e775531f68e5e05708119980c1f6b9ea94afa6b174460001b620073b9565b6000816000019067ffffffffffffffff16908167ffffffffffffffff16815250506200b98c7f6c32ff0ad4c8bb80357afd3065745afdb3cefa0538f813750697a94b12f62d6b60001b620073b9565b6200b9ba7f7db5dbee663494370987686f6ab84c5f3547de9ab1f5d7b492cc11147467e8eb60001b620073b9565b846000015181602001819052506200b9f57f49a7146638182a662761a7666f9a6c8ad0fef03e65220f4b6e843e52f50a8d8c60001b620073b9565b6200ba237f923bb086fd2bfd4670edb92514242b35c09f60b616fa50aa498a1330ac02043260001b620073b9565b856020015181604001819052506200ba5e7f22c09e3276608fc99f94aa4211b7fd55ea44e64a51616ec4bdae2432425274d560001b620073b9565b6200ba8c7fedc0aee1c377c31376f840125a9259f020683c77895dcaf8a2d0b4561593a59160001b620073b9565b846040015181606001819052506200bac77f57438687a312b8020421f070ee48b6e229b231f3b46cf9f5da32a0797da4c98460001b620073b9565b6200baf57f63b82f01d2eadf86878c7430dc7186ceaca7ad439ca3f313f9969fbd5babcdca60001b620073b9565b8181608001819052506200bb2c7fc1bf39448ec93245a27d8d0ae63c1b8b1362d90e6e3fc8788c66995f04a71a0160001b620073b9565b6200bb5a7f4f0e318f55ec82b3a61015222cae868c2ac671e3a57bb96df97aeda8d79d14ba60001b620073b9565b84606001518160a001819052506200bb957fc4b1dbc24c2e4223eae0144af6b6c8031dc8346c721fdf2e701ab9096e0b5d8c60001b620073b9565b6200bbc37fe4839430c604ef5ce3664ce4f680ff051af6f328bbf7c4d57d8c831ca534fd8760001b620073b9565b85600001518160c001819052506200bbfe7f2d92f7571e52cf83dcb54655202b391c9321c01149c1117bf4b2d8f4447f0e2260001b620073b9565b6200bc2c7f86ebe836082c5f1cd2eaebd85b7d50416d5e7ae3233cae37027b1189a6c33cdb60001b620073b9565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b6ddeca0836040518263ffffffff1660e01b81526004016200bc8b919062010977565b60206040518083038186803b1580156200bca457600080fd5b505afa1580156200bcb9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200bcdf91906200ec52565b90506200bd0f7f1c84f7acbc5bbed589b86e08a7bf69257274820542d43e4e0553c88058666cf460001b620073b9565b6200bd3d7fdb6d2b7fa2a881c4a0987efc14ce99931634796541a4f6958e93c053c407019f60001b620073b9565b806200bddf576200bd717fb3cfe50daf8ef3a3cdba7a7b0ef0b28688e263feddc6bf90df6ce97be64872c960001b620073b9565b6200bd9f7f0301213e8aea82888992c1c4744f671e3d14dc13b49e6bdc4a1bfe0a803d1e2260001b620073b9565b6200bdcd7f97ea5e44a15d15a08569b8d4a2c7b71b7a22ed152c582b7f5ccd59944391e89660001b620073b9565b6000985050505050505050506200be76565b6200be0d7fc22c2cc4502ba745b46b8e37aa66e23529ccce8decbfe10f394065802fef609960001b620073b9565b6200be3b7fa847c562115d4ab05e5fc04fc2578360c532990a09c849278caa38d61f574c5b60001b620073b9565b6200be697fbe35e717fbf53305d666bac8f52cb130e353c413ccb2bc35f0f45fd51e29517e60001b620073b9565b6001985050505050505050505b92915050565b60006200beac7f4dd39864b9b8080e1da2ae9b91eb7cc506b0e6bf978c51bdd2f65f3b37448d7460001b620073b9565b6200beda7f8a1ceb46dd0440ac7bbb59994ab900420acce918a53607b148c3026d28c9a47860001b620073b9565b6200bf087feaea0371208a24c73a528452c54a53cf1e311506ad5d7fdef6386cb85ba37b4d60001b620073b9565b814311156200bfa5576200bf3f7f13c1105e0f99568fae82f20d912f7c9e6769a4cc4bec876b96d962122434328a60001b620073b9565b6200bf6d7f3d525515d89955f64ede1165cc9b7ef3a85e313f9049eea99b689b72ac6ac57160001b620073b9565b6200bf9b7fecbaef36c300f880b3caff55a3c2a7733aa2690071b4515c4de8e544678f9a9460001b620073b9565b600190506200c034565b6200bfd37f4ea2041baae17fecea2d40ed90784c7c98d1a836b4b30ab6673d83e31b5c58fb60001b620073b9565b6200c0017fd4d4c842e209b216e1bb3ccf9e34b9060f390c769760d60978c547f24ad8f85760001b620073b9565b6200c02f7fe1b000915fef0abe3427ecd2eaf95b141809eb1fe4ce13f2378d98cb275d802b60001b620073b9565b600090505b919050565b60006200c0697f1c7a237730a13a69ae762c1d97d27b28cc8a74265a93eb0e2acc1df3f891436760001b6200ca4d565b6200c0977f714986d000f06d5358b12b4cfa6095f9ea704ed6daa0e8e3b3e01ab6080d338060001b6200ca4d565b6200c0c57f5aa2883e05feb664a3d1489dbb9fd2a7a37439d0e0e948d9c1e6c5f67ea91be360001b6200ca4d565b60006200c0d48360006200c4a1565b90506200c1047ffde14ee918369832f361533e938548900512a04f6355ef3ae2c728a92f104d3e60001b6200ca4d565b6200c1327f75726af36d58dbbb21f6c181c84c54217073c3c1e3c41c4a95b9b2c969c448fd60001b6200ca4d565b6001816200c141919062010e25565b915050919050565b60006200c1797f637b25c6f1c111685215fbe84d2c62a0ba32a52bb577d097fc51ed5c9e6d229260001b6200ca4d565b6200c1a77fa0425d0733f8d94a38d14d7a4b9088ac50676259a0a8d57a0cbde0eb0a414d4360001b6200ca4d565b6200c1d57fcfb54b1b0e71f6389d122e92c0c97e06ecef2ead46fc77752194ecc239414b2c60001b6200ca4d565b82600101805490508210905092915050565b60006200c1f36200ccb7565b6200c2217f3e1bc6fdde5c31ccf6ff47cf0f4a78bcafab84ab792cc3fc0ae1c01b92a00cb860001b6200ca4d565b6200c24f7f8669ee4ced1fda9e7c547c26b3cd8bb7b35299021bfe6da147a99327581a566960001b6200ca4d565b6200c27d7fa631463aeda9e8ac74d6350bf2d7e3ae49c58c9aedec0e4255049b5748f911e360001b6200ca4d565b8360010183815481106200c2ba577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691506200c3187fd6b5403f2b2962ee58464100a27cbe9e219b6dd73c34bb713319d8a928a35fc460001b6200ca4d565b6200c3467f1df447b03bcedb44291862f04095d480490338a351a3bed484fc57b5c311884060001b6200ca4d565b8360000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001016040518060a00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600282015481526020016003820160009054906101000a900460ff16151515158152505090509250929050565b60006200c4d17f560d1c37e0385cf78dc51d2d22f0590515f444b3c06eb97a97d755c3a61576e260001b6200ca4d565b6200c4ff7f64faf7bcf57b0e70b5ffcd5dcf187cd246bff041e9aefd7304369ceeab5a1f4160001b6200ca4d565b81806200c50c9062011135565b9250506200c53d7f32d6760f3aa00d45a62d732b0363ac9ec344e5718c637bdf63bae74a6f6d1a7c60001b6200ca4d565b6200c56b7fbff44b25d644d63592232190b7142dce2e8552e4c0b842b526873d7ae13d0e9f60001b6200ca4d565b5b8260010180549050821080156200c5d757508260010182815481106200c5bb577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160000160149054906101000a900460ff165b156200c621576200c60b7f695cebe7514e4f7f049cb23dc3b1b72f890f30a1be36331885c90c0640d1919860001b6200ca4d565b81806200c6189062011135565b9250506200c56c565b6200c64f7f85e9943001605b67418ede1f4ca0d233dff10b61caf21ebeccfe62dff1abeeb760001b6200ca4d565b6200c67d7fea5a3a6e11ff893fe1734dde53a483720658db1f7e225bc1c20866097995a00660001b6200ca4d565b81905092915050565b60006200c6b67f289cba5f2dd1d2cc91fb01d2cdb5041976776ad2d9f483892aa5e569c1613eb360001b620073b9565b6200c6e47fd8e419049fc94571543e014f6d982cf5bb0bdeea0bb04c79cf5962a9fbfbb7c260001b620073b9565b6200c7127fbfd117ffa83c488bd1f481b34e5558914c3a913f66520713f86dd29a4d0d77a460001b620073b9565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cc76e80d84600187604001516200c766919062010e60565b60008960a001518a608001516200c77e919062010ddc565b8a6101a001518b61010001516200c796919062010e25565b6040518663ffffffff1660e01b81526004016200c7b8959493929190620108f1565b60606040518083038186803b1580156200c7d157600080fd5b505afa1580156200c7e6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200c80c91906200f1fb565b90506200c83c7f4d4cde182772095554d79aebb92700ee08a346db41720244c205d4adcb3716f060001b620073b9565b6200c86a7f04521ddc36d77af77189d9ffe57d4e3521fe3dc0fcf8a1631a0661a437feb70e60001b620073b9565b8060400151816020015182600001516200c885919062010d5f565b6200c891919062010d5f565b9150509392505050565b60006200c8cb7f02fb1b83b5523d6f551ce66ad701a62d08c8eb2cfe4265c57167b45974a6ad5f60001b620073b9565b6200c8f97feb50cea8771d265890b27830442fc47a426c3e6e9e63362065253c6dadede08860001b620073b9565b6200c9277fceb09aba30a3199b72cfa41c182edbcb7ee31049ba446f863c2752f1e8835cda60001b620073b9565b6000600290506200c95b7f16facca15f2f6c51ac021a3361592fdf63ae781007a0fa0f951c12d544150f2860001b620073b9565b6200c9897ffe1b9038f684ec62a3c18fb32270a38ce288de2e9ed5d0e32af0f39348ea1aef60001b620073b9565b60006200c99b8585606001516200ca50565b826200c9a8919062010ddc565b90506200c9d87fe31d54d43646917ba335c8eb03466ce6fba9e410085aacc0d99a4816814b6c2160001b620073b9565b6200ca067fa3b899637dc2d6e76152961bbd741ead7586a1ee0e61368ac71146ad53ba271660001b620073b9565b809250505092915050565b60008151905060006a636f6e736f6c652e6c6f679050602083016000808483855afa5050505050565b600080823b905060008111915050919050565b50565b60006200ca807fddfec0236af114c110286b2ae9d7064158171bf2d36b258d274734cf8192076860001b620073b9565b6200caae7f743ff5aec4f1024b2952d331bf3e4587b52da15f256d749289ea419bdea9f7df60001b620073b9565b6200cadc7fe34a54510219ff588df2b407b38455d180f34be2fa4fc74da2633ddc38da774e60001b620073b9565b620fa0008284606001516200caf2919062010ddc565b6200cafe919062010da4565b905092915050565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6040518060c00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081526020016060815260200160608152602001606081525090565b60405180606001604052806200cbb76200cdab565b8152602001606081526020016200cbcd6200cb58565b81525090565b6040518060e0016040528060608152602001606081526020016060815260200160608152602001606081526020016200cc0b6200ce99565b81526020016000151581525090565b6040518060e00160405280600067ffffffffffffffff1681526020016060815260200160608152602001606081526020016060815260200160608152602001606081525090565b60405180606001604052806200cc766200d02e565b815260200160608152602001600067ffffffffffffffff1681525090565b50805460008255906000526020600020908101906200ccb491906200d06d565b50565b6040518060a00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff168152602001600081526020016000151581525090565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600067ffffffffffffffff1681525090565b604051806040016040528060608152602001606081525090565b604051806080016040528060608152602001600067ffffffffffffffff16815260200160608152602001606081525090565b604051806101800160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600060028111156200ce42577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016000815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600015158152602001606081525090565b604051806102c0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff16815260200160008152602001600015158152602001600060018111156200cfb1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600067ffffffffffffffff1681526020016060815260200160608152602001600060028111156200d010577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016000151581526020016200d0286200d02e565b81525090565b6040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b5b808211156200d0bd57600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549060ff0219169055506001016200d06e565b5090565b60006200d0d86200d0d284620109c4565b6201099b565b905080838252602082019050828560208602820111156200d0f857600080fd5b60005b858110156200d12c57816200d11188826200d55f565b8452602084019350602083019250506001810190506200d0fb565b5050509392505050565b60006200d14d6200d14784620109c4565b6201099b565b905080838252602082019050828560208602820111156200d16d57600080fd5b60005b858110156200d1a157816200d18688826200d576565b8452602084019350602083019250506001810190506200d170565b5050509392505050565b60006200d1c26200d1bc84620109f3565b6201099b565b905080838252602082019050828560208602820111156200d1e257600080fd5b60005b858110156200d23157813567ffffffffffffffff8111156200d20657600080fd5b8086016200d21589826200d73a565b855260208501945060208401935050506001810190506200d1e5565b5050509392505050565b60006200d2526200d24c84620109f3565b6201099b565b905080838252602082019050828560208602820111156200d27257600080fd5b60005b858110156200d2c157815167ffffffffffffffff8111156200d29657600080fd5b8086016200d2a589826200d767565b855260208501945060208401935050506001810190506200d275565b5050509392505050565b60006200d2e26200d2dc8462010a22565b6201099b565b905080838252602082019050828560408602820111156200d30257600080fd5b60005b858110156200d33657816200d31b88826200d863565b8452602084019350604083019250506001810190506200d305565b5050509392505050565b60006200d3576200d3518462010a51565b6201099b565b905080838252602082019050828560208602820111156200d37757600080fd5b60005b858110156200d3c657815167ffffffffffffffff8111156200d39b57600080fd5b8086016200d3aa89826200dedf565b855260208501945060208401935050506001810190506200d37a565b5050509392505050565b60006200d3e76200d3e18462010a80565b6201099b565b905080838252602082019050828560208602820111156200d40757600080fd5b60005b858110156200d45657815167ffffffffffffffff8111156200d42b57600080fd5b8086016200d43a89826200df61565b855260208501945060208401935050506001810190506200d40a565b5050509392505050565b60006200d4776200d4718462010aaf565b6201099b565b905080838252602082019050828560a08602820111156200d49757600080fd5b60005b858110156200d4cb57816200d4b088826200e434565b845260208401935060a083019250506001810190506200d49a565b5050509392505050565b60006200d4ec6200d4e68462010ade565b6201099b565b9050828152602081018484840111156200d50557600080fd5b6200d512848285620110ba565b509392505050565b60006200d5316200d52b8462010ade565b6201099b565b9050828152602081018484840111156200d54a57600080fd5b6200d557848285620110c9565b509392505050565b6000813590506200d5708162011382565b92915050565b6000815190506200d5878162011382565b92915050565b600082601f8301126200d59f57600080fd5b81356200d5b18482602086016200d0c1565b91505092915050565b600082601f8301126200d5cc57600080fd5b81516200d5de8482602086016200d136565b91505092915050565b600082601f8301126200d5f957600080fd5b81356200d60b8482602086016200d1ab565b91505092915050565b600082601f8301126200d62657600080fd5b81516200d6388482602086016200d23b565b91505092915050565b600082601f8301126200d65357600080fd5b81516200d6658482602086016200d2cb565b91505092915050565b600082601f8301126200d68057600080fd5b81516200d6928482602086016200d340565b91505092915050565b600082601f8301126200d6ad57600080fd5b81516200d6bf8482602086016200d3d0565b91505092915050565b600082601f8301126200d6da57600080fd5b81356200d6ec8482602086016200d460565b91505092915050565b6000813590506200d706816201139c565b92915050565b6000815190506200d71d816201139c565b92915050565b6000813590506200d73481620113b6565b92915050565b600082601f8301126200d74c57600080fd5b81356200d75e8482602086016200d4d5565b91505092915050565b600082601f8301126200d77957600080fd5b81516200d78b8482602086016200d51a565b91505092915050565b6000813590506200d7a581620113d0565b92915050565b6000813590506200d7bc81620113ea565b92915050565b6000813590506200d7d38162011404565b92915050565b6000813590506200d7ea816201141e565b92915050565b6000813590506200d8018162011438565b92915050565b6000813590506200d8188162011452565b92915050565b6000815190506200d82f8162011452565b92915050565b6000813590506200d8468162011463565b92915050565b6000815190506200d85d8162011463565b92915050565b6000604082840312156200d87657600080fd5b6200d88260406201099b565b905060006200d894848285016200eb31565b60008301525060206200d8aa848285016200eb31565b60208301525092915050565b600061030082840312156200d8ca57600080fd5b6200d8d76102c06201099b565b9050600082013567ffffffffffffffff8111156200d8f457600080fd5b6200d902848285016200d73a565b60008301525060206200d918848285016200d55f565b602083015250604082013567ffffffffffffffff8111156200d93957600080fd5b6200d947848285016200d73a565b60408301525060606200d95d848285016200eb48565b60608301525060806200d973848285016200eb48565b60808301525060a06200d989848285016200eb48565b60a08301525060c06200d99f848285016200eb48565b60c08301525060e06200d9b5848285016200eb48565b60e0830152506101006200d9cc848285016200eb03565b610100830152506101206200d9e4848285016200eb48565b610120830152506101406200d9fc848285016200eb48565b6101408301525061016082013567ffffffffffffffff8111156200da1f57600080fd5b6200da2d848285016200d73a565b610160830152506101806200da45848285016200eb48565b610180830152506101a06200da5d848285016200eb03565b6101a0830152506101c06200da75848285016200d6f5565b6101c0830152506101e06200da8d848285016200d835565b6101e0830152506102006200daa5848285016200eb48565b6102008301525061022082013567ffffffffffffffff8111156200dac857600080fd5b6200dad6848285016200d58d565b6102208301525061024082013567ffffffffffffffff8111156200daf957600080fd5b6200db07848285016200d58d565b610240830152506102606200db1f848285016200d807565b610260830152506102806200db37848285016200d6f5565b610280830152506102a06200db4f848285016200e2a6565b6102a08301525092915050565b600061030082840312156200db7057600080fd5b6200db7d6102c06201099b565b9050600082015167ffffffffffffffff8111156200db9a57600080fd5b6200dba8848285016200d767565b60008301525060206200dbbe848285016200d576565b602083015250604082015167ffffffffffffffff8111156200dbdf57600080fd5b6200dbed848285016200d767565b60408301525060606200dc03848285016200eb5f565b60608301525060806200dc19848285016200eb5f565b60808301525060a06200dc2f848285016200eb5f565b60a08301525060c06200dc45848285016200eb5f565b60c08301525060e06200dc5b848285016200eb5f565b60e0830152506101006200dc72848285016200eb1a565b610100830152506101206200dc8a848285016200eb5f565b610120830152506101406200dca2848285016200eb5f565b6101408301525061016082015167ffffffffffffffff8111156200dcc557600080fd5b6200dcd3848285016200d767565b610160830152506101806200dceb848285016200eb5f565b610180830152506101a06200dd03848285016200eb1a565b6101a0830152506101c06200dd1b848285016200d70c565b6101c0830152506101e06200dd33848285016200d84c565b6101e0830152506102006200dd4b848285016200eb5f565b6102008301525061022082015167ffffffffffffffff8111156200dd6e57600080fd5b6200dd7c848285016200d5ba565b6102208301525061024082015167ffffffffffffffff8111156200dd9f57600080fd5b6200ddad848285016200d5ba565b610240830152506102606200ddc5848285016200d81e565b610260830152506102806200dddd848285016200d70c565b610280830152506102a06200ddf5848285016200e30f565b6102a08301525092915050565b600060c082840312156200de1557600080fd5b6200de2160c06201099b565b9050600082013567ffffffffffffffff8111156200de3e57600080fd5b6200de4c848285016200d73a565b600083015250602082013567ffffffffffffffff8111156200de6d57600080fd5b6200de7b848285016200d73a565b60208301525060406200de91848285016200eb03565b60408301525060606200dea7848285016200d55f565b60608301525060806200debd848285016200eb48565b60808301525060a06200ded3848285016200eb48565b60a08301525092915050565b6000606082840312156200def257600080fd5b6200defe60606201099b565b905060006200df10848285016200eb5f565b60008301525060206200df26848285016200eb5f565b602083015250604082015167ffffffffffffffff8111156200df4757600080fd5b6200df55848285016200d767565b60408301525092915050565b6000604082840312156200df7457600080fd5b6200df8060406201099b565b905060006200df92848285016200eb5f565b600083015250602082015167ffffffffffffffff8111156200dfb357600080fd5b6200dfc1848285016200d66e565b60208301525092915050565b600060e082840312156200dfe057600080fd5b6200dfec60e06201099b565b905060006200dffe848285016200eb48565b60008301525060206200e014848285016200eb48565b60208301525060406200e02a848285016200eb48565b60408301525060606200e040848285016200eb48565b60608301525060806200e056848285016200eb48565b60808301525060a06200e06c848285016200d55f565b60a08301525060c06200e082848285016200d55f565b60c08301525092915050565b600060e082840312156200e0a157600080fd5b6200e0ad60e06201099b565b905060006200e0bf848285016200eb5f565b60008301525060206200e0d5848285016200eb5f565b60208301525060406200e0eb848285016200eb5f565b60408301525060606200e101848285016200eb5f565b60608301525060806200e117848285016200eb5f565b60808301525060a06200e12d848285016200d576565b60a08301525060c06200e143848285016200d576565b60c08301525092915050565b600060e082840312156200e16257600080fd5b6200e16e60e06201099b565b9050600082015167ffffffffffffffff8111156200e18b57600080fd5b6200e199848285016200d614565b600083015250602082015167ffffffffffffffff8111156200e1ba57600080fd5b6200e1c8848285016200d614565b602083015250604082015167ffffffffffffffff8111156200e1e957600080fd5b6200e1f7848285016200d641565b604083015250606082015167ffffffffffffffff8111156200e21857600080fd5b6200e226848285016200d69b565b606083015250608082015167ffffffffffffffff8111156200e24757600080fd5b6200e255848285016200d767565b60808301525060a082015167ffffffffffffffff8111156200e27657600080fd5b6200e284848285016200db5c565b60a08301525060c06200e29a848285016200d70c565b60c08301525092915050565b6000606082840312156200e2b957600080fd5b6200e2c560606201099b565b905060006200e2d7848285016200eb48565b60008301525060206200e2ed848285016200eb48565b60208301525060406200e303848285016200eb48565b60408301525092915050565b6000606082840312156200e32257600080fd5b6200e32e60606201099b565b905060006200e340848285016200eb5f565b60008301525060206200e356848285016200eb5f565b60208301525060406200e36c848285016200eb5f565b60408301525092915050565b6000606082840312156200e38b57600080fd5b6200e39760606201099b565b905060006200e3a9848285016200d55f565b60008301525060206200e3bf848285016200eb03565b60208301525060406200e3d5848285016200eb48565b60408301525092915050565b6000604082840312156200e3f457600080fd5b6200e40060406201099b565b905060006200e412848285016200eb48565b60008301525060206200e428848285016200eb48565b60208301525092915050565b600060a082840312156200e44757600080fd5b6200e45360a06201099b565b905060006200e465848285016200d55f565b60008301525060206200e47b848285016200d55f565b60208301525060406200e491848285016200eb48565b60408301525060606200e4a7848285016200eb03565b60608301525060806200e4bd848285016200d6f5565b60808301525092915050565b600061018082840312156200e4dd57600080fd5b6200e4ea6101806201099b565b905060006200e4fc848285016200d55f565b60008301525060206200e512848285016200eb48565b60208301525060406200e528848285016200eb48565b60408301525060606200e53e848285016200eb48565b60608301525060806200e554848285016200d807565b60808301525060a06200e56a848285016200eb03565b60a08301525060c06200e580848285016200eb03565b60c08301525060e06200e596848285016200eb48565b60e0830152506101006200e5ad848285016200eb48565b610100830152506101206200e5c5848285016200eb48565b610120830152506101406200e5dd848285016200d6f5565b6101408301525061016082013567ffffffffffffffff8111156200e60057600080fd5b6200e60e848285016200d5e7565b6101608301525092915050565b600061018082840312156200e62f57600080fd5b6200e63c6101806201099b565b905060006200e64e848285016200d576565b60008301525060206200e664848285016200eb5f565b60208301525060406200e67a848285016200eb5f565b60408301525060606200e690848285016200eb5f565b60608301525060806200e6a6848285016200d81e565b60808301525060a06200e6bc848285016200eb1a565b60a08301525060c06200e6d2848285016200eb1a565b60c08301525060e06200e6e8848285016200eb5f565b60e0830152506101006200e6ff848285016200eb5f565b610100830152506101206200e717848285016200eb5f565b610120830152506101406200e72f848285016200d70c565b6101408301525061016082015167ffffffffffffffff8111156200e75257600080fd5b6200e760848285016200d614565b6101608301525092915050565b6000608082840312156200e78057600080fd5b6200e78c60806201099b565b905060006200e79e848285016200d55f565b60008301525060206200e7b4848285016200eb48565b60208301525060406200e7ca848285016200eb48565b604083015250606082013567ffffffffffffffff8111156200e7eb57600080fd5b6200e7f9848285016200d73a565b60608301525092915050565b6000604082840312156200e81857600080fd5b6200e82460406201099b565b905060006200e836848285016200d55f565b60008301525060206200e84c848285016200eb48565b60208301525092915050565b600061016082840312156200e86c57600080fd5b6200e8796101606201099b565b905060006200e88b848285016200eb48565b60008301525060206200e8a1848285016200eb48565b60208301525060406200e8b7848285016200eb48565b60408301525060606200e8cd848285016200eb48565b60608301525060806200e8e3848285016200eb48565b60808301525060a06200e8f9848285016200eb48565b60a08301525060c06200e90f848285016200eb48565b60c08301525060e06200e925848285016200eb48565b60e0830152506101006200e93c848285016200eb48565b610100830152506101206200e954848285016200eb48565b610120830152506101406200e96c848285016200eb48565b6101408301525092915050565b600061016082840312156200e98d57600080fd5b6200e99a6101606201099b565b905060006200e9ac848285016200eb5f565b60008301525060206200e9c2848285016200eb5f565b60208301525060406200e9d8848285016200eb5f565b60408301525060606200e9ee848285016200eb5f565b60608301525060806200ea04848285016200eb5f565b60808301525060a06200ea1a848285016200eb5f565b60a08301525060c06200ea30848285016200eb5f565b60c08301525060e06200ea46848285016200eb5f565b60e0830152506101006200ea5d848285016200eb5f565b610100830152506101206200ea75848285016200eb5f565b610120830152506101406200ea8d848285016200eb5f565b6101408301525092915050565b6000606082840312156200eaad57600080fd5b6200eab960606201099b565b905060006200eacb848285016200eb5f565b60008301525060206200eae1848285016200eb5f565b60208301525060406200eaf7848285016200eb5f565b60408301525092915050565b6000813590506200eb148162011474565b92915050565b6000815190506200eb2b8162011474565b92915050565b6000815190506200eb42816201148e565b92915050565b6000813590506200eb5981620114a8565b92915050565b6000815190506200eb7081620114a8565b92915050565b600080604083850312156200eb8a57600080fd5b60006200eb9a858286016200d55f565b92505060206200ebad858286016200eb03565b9150509250929050565b6000806000606084860312156200ebcd57600080fd5b60006200ebdd868287016200d55f565b93505060206200ebf0868287016200eb48565b92505060406200ec03868287016200eb03565b9150509250925092565b6000602082840312156200ec2057600080fd5b600082015167ffffffffffffffff8111156200ec3b57600080fd5b6200ec49848285016200d641565b91505092915050565b6000602082840312156200ec6557600080fd5b60006200ec75848285016200d70c565b91505092915050565b6000602082840312156200ec9157600080fd5b60006200eca1848285016200d723565b91505092915050565b6000602082840312156200ecbd57600080fd5b600082013567ffffffffffffffff8111156200ecd857600080fd5b6200ece6848285016200d73a565b91505092915050565b600080604083850312156200ed0357600080fd5b600083013567ffffffffffffffff8111156200ed1e57600080fd5b6200ed2c858286016200d73a565b925050602083013567ffffffffffffffff8111156200ed4a57600080fd5b6200ed58858286016200d6c8565b9150509250929050565b600080606083850312156200ed7657600080fd5b600083013567ffffffffffffffff8111156200ed9157600080fd5b6200ed9f858286016200d73a565b92505060206200edb2858286016200e3e1565b9150509250929050565b600080600080600060a086880312156200edd557600080fd5b60006200ede5888289016200d794565b95505060206200edf8888289016200d7ab565b94505060406200ee0b888289016200d7c2565b93505060606200ee1e888289016200d7d9565b92505060806200ee31888289016200d7f0565b9150509295509295909350565b6000602082840312156200ee5157600080fd5b600082015167ffffffffffffffff8111156200ee6c57600080fd5b6200ee7a848285016200db5c565b91505092915050565b600080600080600061032086880312156200ee9d57600080fd5b600086013567ffffffffffffffff8111156200eeb857600080fd5b6200eec6888289016200d8b6565b95505060206200eed9888289016200dfcd565b9450506101006200eeed888289016200e434565b9350506101a086013567ffffffffffffffff8111156200ef0c57600080fd5b6200ef1a888289016200d6c8565b9250506101c06200ef2e888289016200e858565b9150509295509295909350565b6000602082840312156200ef4e57600080fd5b600082013567ffffffffffffffff8111156200ef6957600080fd5b6200ef77848285016200de02565b91505092915050565b600060e082840312156200ef9357600080fd5b60006200efa3848285016200e08e565b91505092915050565b6000602082840312156200efbf57600080fd5b600082015167ffffffffffffffff8111156200efda57600080fd5b6200efe8848285016200e14f565b91505092915050565b6000606082840312156200f00457600080fd5b60006200f014848285016200e378565b91505092915050565b6000602082840312156200f03057600080fd5b600082015167ffffffffffffffff8111156200f04b57600080fd5b6200f059848285016200e61b565b91505092915050565b60008060008061028085870312156200f07a57600080fd5b600085013567ffffffffffffffff8111156200f09557600080fd5b6200f0a3878288016200e4c9565b94505060206200f0b6878288016200dfcd565b9350506101006200f0ca878288016200e858565b9250506102606200f0de878288016200eb48565b91505092959194509250565b6000602082840312156200f0fd57600080fd5b600082013567ffffffffffffffff8111156200f11857600080fd5b6200f126848285016200e76d565b91505092915050565b600080604083850312156200f14357600080fd5b600083013567ffffffffffffffff8111156200f15e57600080fd5b6200f16c858286016200e76d565b925050602083013567ffffffffffffffff8111156200f18a57600080fd5b6200f198858286016200e4c9565b9150509250929050565b6000604082840312156200f1b557600080fd5b60006200f1c5848285016200e805565b91505092915050565b600061016082840312156200f1e257600080fd5b60006200f1f2848285016200e979565b91505092915050565b6000606082840312156200f20e57600080fd5b60006200f21e848285016200ea9a565b91505092915050565b60006200f23583836200f2d1565b60208301905092915050565b60006200f24f83836200f65b565b905092915050565b60006200f26583836200f8c8565b60408301905092915050565b60006200f27f83836200fb84565b905092915050565b60006200f29583836200fbda565b905092915050565b60006200f2ab83836200fcff565b60608301905092915050565b60006200f2c583836200fdf3565b60a08301905092915050565b6200f2dc8162010e9b565b82525050565b6200f2ed8162010e9b565b82525050565b6200f3086200f3028262010e9b565b620111b9565b82525050565b60006200f31b8262010b84565b6200f327818562010c42565b93506200f3348362010b14565b8060005b838110156200f36b5781516200f34f88826200f227565b97506200f35c8362010be7565b9250506001810190506200f338565b5085935050505092915050565b60006200f3858262010b8f565b6200f391818562010c53565b9350836020820285016200f3a58562010b24565b8060005b858110156200f3e757848403895281516200f3c585826200f241565b94506200f3d28362010bf4565b925060208a019950506001810190506200f3a9565b50829750879550505050505092915050565b60006200f4068262010b9a565b6200f412818562010c64565b93506200f41f8362010b34565b8060005b838110156200f4565781516200f43a88826200f257565b97506200f4478362010c01565b9250506001810190506200f423565b5085935050505092915050565b60006200f4708262010ba5565b6200f47c818562010c75565b9350836020820285016200f4908562010b44565b8060005b858110156200f4d257848403895281516200f4b085826200f271565b94506200f4bd8362010c0e565b925060208a019950506001810190506200f494565b50829750879550505050505092915050565b60006200f4f18262010bb0565b6200f4fd818562010c86565b9350836020820285016200f5118562010b54565b8060005b858110156200f55357848403895281516200f53185826200f287565b94506200f53e8362010c1b565b925060208a019950506001810190506200f515565b50829750879550505050505092915050565b60006200f5728262010bbb565b6200f57e818562010c97565b93506200f58b8362010b64565b8060005b838110156200f5c25781516200f5a688826200f29d565b97506200f5b38362010c28565b9250506001810190506200f58f565b5085935050505092915050565b60006200f5dc8262010bc6565b6200f5e8818562010ca8565b93506200f5f58362010b74565b8060005b838110156200f62c5781516200f61088826200f2b7565b97506200f61d8362010c35565b9250506001810190506200f5f9565b5085935050505092915050565b6200f6448162010eaf565b82525050565b6200f6558162010eaf565b82525050565b60006200f6688262010bd1565b6200f674818562010cb9565b93506200f686818560208601620110c9565b6200f69181620112a7565b840191505092915050565b60006200f6a98262010bd1565b6200f6b5818562010cca565b93506200f6c7818560208601620110c9565b6200f6d281620112a7565b840191505092915050565b60006200f6ea8262010bd1565b6200f6f6818562010cdb565b93506200f708818560208601620110c9565b80840191505092915050565b6200f71f8162010fb6565b82525050565b6200f7308162010fca565b82525050565b6200f7418162010fca565b82525050565b6200f7528162010fde565b82525050565b6200f7638162010ff2565b82525050565b6200f7748162011006565b82525050565b6200f785816201101a565b82525050565b6200f796816201102e565b82525050565b6200f7a78162011042565b82525050565b6200f7b88162011056565b82525050565b6200f7c9816201106a565b82525050565b6200f7da816201107e565b82525050565b6200f7eb8162011092565b82525050565b6200f7fc81620110a6565b82525050565b60006200f80f8262010bdc565b6200f81b818562010ce6565b93506200f82d818560208601620110c9565b6200f83881620112a7565b840191505092915050565b60006200f8508262010bdc565b6200f85c818562010cf7565b93506200f86e818560208601620110c9565b80840191505092915050565b60006200f88960168362010ce6565b91506200f89682620112c5565b602082019050919050565b60006200f8b0602e8362010ce6565b91506200f8bd82620112ee565b604082019050919050565b6040820160008201516200f8e06000850182620103d1565b5060208201516200f8f56020850182620103d1565b50505050565b60006103008301600083015184820360008601526200f91b82826200f65b565b91505060208301516200f93260208601826200f2d1565b50604083015184820360408601526200f94c82826200f65b565b91505060608301516200f9636060860182620103e2565b5060808301516200f9786080860182620103e2565b5060a08301516200f98d60a0860182620103e2565b5060c08301516200f9a260c0860182620103e2565b5060e08301516200f9b760e0860182620103e2565b506101008301516200f9ce61010086018262010394565b506101208301516200f9e5610120860182620103e2565b506101408301516200f9fc610140860182620103e2565b506101608301518482036101608601526200fa1882826200f65b565b9150506101808301516200fa31610180860182620103e2565b506101a08301516200fa486101a086018262010394565b506101c08301516200fa5f6101c08601826200f639565b506101e08301516200fa766101e08601826200f747565b506102008301516200fa8d610200860182620103e2565b506102208301518482036102208601526200faa982826200f30e565b9150506102408301518482036102408601526200fac782826200f30e565b9150506102608301516200fae06102608601826200f725565b506102808301516200faf76102808601826200f639565b506102a08301516200fb0e6102a08601826200fcb7565b508091505092915050565b60006080830160008301516200fb3360008601826200f2d1565b50602083015184820360208601526200fb4d82826200f65b565b91505060408301516200fb646040860182620103e2565b5060608301516200fb796060860182620103e2565b508091505092915050565b60006060830160008301516200fb9e6000860182620103e2565b5060208301516200fbb36020860182620103e2565b50604083015184820360408601526200fbcd82826200f65b565b9150508091505092915050565b60006040830160008301516200fbf46000860182620103e2565b50602083015184820360208601526200fc0e82826200f463565b9150508091505092915050565b60e0820160008201516200fc336000850182620103e2565b5060208201516200fc486020850182620103e2565b5060408201516200fc5d6040850182620103e2565b5060608201516200fc726060850182620103e2565b5060808201516200fc876080850182620103e2565b5060a08201516200fc9c60a08501826200f2d1565b5060c08201516200fcb160c08501826200f2d1565b50505050565b6060820160008201516200fccf6000850182620103e2565b5060208201516200fce46020850182620103e2565b5060408201516200fcf96040850182620103e2565b50505050565b6060820160008201516200fd1760008501826200f2d1565b5060208201516200fd2c602085018262010394565b5060408201516200fd416040850182620103e2565b50505050565b6060820160008201516200fd5f60008501826200f2d1565b5060208201516200fd74602085018262010394565b5060408201516200fd896040850182620103e2565b50505050565b600060608301600083015184820360008601526200fdae82826200fe65565b915050602083015184820360208601526200fdca82826200f3f9565b915050604083015184820360408601526200fde682826201009d565b9150508091505092915050565b60a0820160008201516200fe0b60008501826200f2d1565b5060208201516200fe2060208501826200f2d1565b5060408201516200fe356040850182620103e2565b5060608201516200fe4a606085018262010394565b5060808201516200fe5f60808501826200f639565b50505050565b6000610180830160008301516200fe8060008601826200f2d1565b5060208301516200fe956020860182620103e2565b5060408301516200feaa6040860182620103e2565b5060608301516200febf6060860182620103e2565b5060808301516200fed460808601826200f725565b5060a08301516200fee960a086018262010394565b5060c08301516200fefe60c086018262010394565b5060e08301516200ff1360e0860182620103e2565b506101008301516200ff2a610100860182620103e2565b506101208301516200ff41610120860182620103e2565b506101408301516200ff586101408601826200f639565b506101608301518482036101608601526200ff7482826200f378565b9150508091505092915050565b6000610180830160008301516200ff9c60008601826200f2d1565b5060208301516200ffb16020860182620103e2565b5060408301516200ffc66040860182620103e2565b5060608301516200ffdb6060860182620103e2565b5060808301516200fff060808601826200f725565b5060a08301516201000560a086018262010394565b5060c08301516201001a60c086018262010394565b5060e08301516201002f60e0860182620103e2565b5061010083015162010046610100860182620103e2565b506101208301516201005d610120860182620103e2565b50610140830151620100746101408601826200f639565b506101608301518482036101608601526201009082826200f378565b9150508091505092915050565b600060c083016000830151620100b76000860182620103e2565b506020830151620100cc6020860182620103e2565b5060408301518482036040860152620100e682826200f65b565b915050606083015184820360608601526201010282826200f65b565b915050608083015184820360808601526201011e82826200f4e4565b91505060a083015184820360a08601526201013a82826200f65b565b9150508091505092915050565b6040820160008201516201015f60008501826200f2d1565b506020820151620101746020850182620103e2565b50505050565b61016082016000820151620101936000850182620103e2565b506020820151620101a86020850182620103e2565b506040820151620101bd6040850182620103e2565b506060820151620101d26060850182620103e2565b506080820151620101e76080850182620103e2565b5060a0820151620101fc60a0850182620103e2565b5060c08201516201021160c0850182620103e2565b5060e08201516201022660e0850182620103e2565b506101008201516201023d610100850182620103e2565b5061012082015162010254610120850182620103e2565b506101408201516201026b610140850182620103e2565b50505050565b600060a0830160008301516201028b60008601826200fcb7565b5060208301518482036060860152620102a582826200f65b565b9150506040830151620102bc6080860182620103e2565b508091505092915050565b600060e083016000830151620102e16000860182620103e2565b5060208301518482036020860152620102fb82826200f65b565b915050604083015184820360408601526201031782826200f378565b915050606083015184820360608601526201033382826200f378565b915050608083015184820360808601526201034f82826200f3f9565b91505060a083015184820360a08601526201036b82826200f4e4565b91505060c083015184820360c08601526201038782826200f65b565b9150508091505092915050565b6201039f8162010f88565b82525050565b620103b08162010f88565b82525050565b620103cb620103c58262010f88565b620111e1565b82525050565b620103dc8162010f92565b82525050565b620103ed8162010fa2565b82525050565b620103fe8162010fa2565b82525050565b60006201041282856200f2f3565b601482019150620104248284620103b6565b6020820191508190509392505050565b60006201044282846200f6dd565b915081905092915050565b60006201045b82846200f843565b915081905092915050565b60006020820190506201047d60008301846200f2e2565b92915050565b60006040820190506201049a60008301856200f2e2565b8181036020830152620104ae81846200f69c565b90509392505050565b60006020820190508181036000830152620104d381846200f565565b905092915050565b60006020820190508181036000830152620104f781846200f5cf565b905092915050565b60006020820190506201051660008301846200f64a565b92915050565b600060208201905081810360008301526201053881846200f69c565b905092915050565b60006080820190506201055760008301876200f714565b620105666020830186620103a5565b6201057560408301856200f2e2565b620105846060830184620103f3565b95945050505050565b6000608082019050620105a460008301876200f714565b620105b36020830186620103a5565b8181036040830152620105c781856200f69c565b9050620105d860608301846200f2e2565b95945050505050565b6000602082019050620105f860008301846200f736565b92915050565b60006020820190506201061560008301846200f769565b92915050565b60006020820190506201063260008301846200f77a565b92915050565b60006020820190506201064f60008301846200f78b565b92915050565b60006020820190506201066c60008301846200f79c565b92915050565b60006020820190506201068960008301846200f7ad565b92915050565b6000602082019050620106a660008301846200f7be565b92915050565b6000602082019050620106c360008301846200f7cf565b92915050565b6000602082019050620106e060008301846200f7e0565b92915050565b6000602082019050620106fd60008301846200f7f1565b92915050565b600060608201905081810360008301526201071f81866200f802565b9050620107306020830185620103a5565b6201073f6040830184620103a5565b949350505050565b6000602082019050818103600083015262010762816200f87a565b9050919050565b6000602082019050818103600083015262010784816200f8a1565b9050919050565b60006020820190508181036000830152620107a781846200f8fb565b905092915050565b60006060820190508181036000830152620107cb81866200f8fb565b9050620107dc60208301856200f64a565b620107eb60408301846200f64a565b949350505050565b600060208201905081810360008301526201080f81846200fb19565b905092915050565b600060e0820190506201082e60008301846200fc1b565b92915050565b60006060820190506201084b60008301846200fd47565b92915050565b600060208201905081810360008301526201086d81846200fd8f565b905092915050565b600060208201905081810360008301526201089181846200ff81565b905092915050565b60006040820190508181036000830152620108b581856200ff81565b90508181036020830152620108cb81846200f8fb565b90509392505050565b6000604082019050620108eb600083018462010147565b92915050565b60006101e0820190506201090960008301886201017a565b62010919610160830187620103f3565b620109296101808301866200f758565b620109396101a0830185620103f3565b620109496101c0830184620103f3565b9695505050505050565b600060208201905081810360008301526201096f818462010271565b905092915050565b60006020820190508181036000830152620109938184620102c7565b905092915050565b6000620109a7620109ba565b9050620109b58282620110ff565b919050565b6000604051905090565b600067ffffffffffffffff821115620109e257620109e162011278565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562010a115762010a1062011278565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562010a405762010a3f62011278565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562010a6f5762010a6e62011278565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562010a9e5762010a9d62011278565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562010acd5762010acc62011278565b5b602082029050602081019050919050565b600067ffffffffffffffff82111562010afc5762010afb62011278565b5b62010b0782620112a7565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600081905092915050565b600062010d0f8262010f88565b915062010d1c8362010f88565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111562010d545762010d53620111eb565b5b828201905092915050565b600062010d6c8262010fa2565b915062010d798362010fa2565b92508267ffffffffffffffff0382111562010d995762010d98620111eb565b5b828201905092915050565b600062010db18262010fa2565b915062010dbe8362010fa2565b92508262010dd15762010dd06201121a565b5b828204905092915050565b600062010de98262010fa2565b915062010df68362010fa2565b92508167ffffffffffffffff048311821515161562010e1a5762010e19620111eb565b5b828202905092915050565b600062010e328262010f88565b915062010e3f8362010f88565b92508282101562010e555762010e54620111eb565b5b828203905092915050565b600062010e6d8262010fa2565b915062010e7a8362010fa2565b92508282101562010e905762010e8f620111eb565b5b828203905092915050565b600062010ea88262010f68565b9050919050565b60008115159050919050565b6000819050919050565b600062010ed28262010e9b565b9050919050565b600062010ee68262010e9b565b9050919050565b600062010efa8262010e9b565b9050919050565b600062010f0e8262010e9b565b9050919050565b600062010f228262010e9b565b9050919050565b600081905062010f39826201133d565b919050565b600081905062010f4e8262011354565b919050565b600081905062010f63826201136b565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b600067ffffffffffffffff82169050919050565b600062010fc38262010f29565b9050919050565b600062010fd78262010f3e565b9050919050565b600062010feb8262010f53565b9050919050565b600062010fff8262010fa2565b9050919050565b6000620110138262010fa2565b9050919050565b6000620110278262010fa2565b9050919050565b60006201103b8262010fa2565b9050919050565b60006201104f8262010fa2565b9050919050565b6000620110638262010fa2565b9050919050565b6000620110778262010fa2565b9050919050565b60006201108b8262010fa2565b9050919050565b60006201109f8262010fa2565b9050919050565b6000620110b38262010fa2565b9050919050565b82818337600083830152505050565b60005b83811015620110e9578082015181840152602081019050620110cc565b83811115620110f9576000848401525b50505050565b6201110a82620112a7565b810181811067ffffffffffffffff821117156201112c576201112b62011278565b5b80604052505050565b6000620111428262010f88565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415620111785762011177620111eb565b5b600182019050919050565b6000620111908262010fa2565b915067ffffffffffffffff821415620111ae57620111ad620111eb565b5b600182019050919050565b6000620111c682620111cd565b9050919050565b6000620111da82620112b8565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f50756e697368466f72536563746f72206661696c656400000000000000000000600082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b600a811062011351576201135062011249565b5b50565b6003811062011368576201136762011249565b5b50565b600281106201137f576201137e62011249565b5b50565b6201138d8162010e9b565b81146201139957600080fd5b50565b620113a78162010eaf565b8114620113b357600080fd5b50565b620113c18162010ebb565b8114620113cd57600080fd5b50565b620113db8162010ec5565b8114620113e757600080fd5b50565b620113f58162010ed9565b81146201140157600080fd5b50565b6201140f8162010eed565b81146201141b57600080fd5b50565b620114298162010f01565b81146201143557600080fd5b50565b620114438162010f15565b81146201144f57600080fd5b50565b600381106201146057600080fd5b50565b600281106201147157600080fd5b50565b6201147f8162010f88565b81146201148b57600080fd5b50565b620114998162010f92565b8114620114a557600080fd5b50565b620114b38162010fa2565b8114620114bf57600080fd5b5056fea2646970667358221220e1f169778dab9c812af8298446d07021688c1969c67b266a39934e745a84e6d164736f6c63430008040033",
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
func (_Store *StoreCaller) CheckSectorProve(opts *bind.CallOpts, sectorProve ProveSectorProveParams, sectorInfo SectorInfo) (bool, error) {
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
func (_Store *StoreSession) CheckSectorProve(sectorProve ProveSectorProveParams, sectorInfo SectorInfo) (bool, error) {
	return _Store.Contract.CheckSectorProve(&_Store.CallOpts, sectorProve, sectorInfo)
}

// CheckSectorProve is a free data retrieval call binding the contract method 0x0fece869.
//
// Solidity: function checkSectorProve((address,uint64,uint64,bytes) sectorProve, (address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) view returns(bool)
func (_Store *StoreCallerSession) CheckSectorProve(sectorProve ProveSectorProveParams, sectorInfo SectorInfo) (bool, error) {
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
func (_Store *StoreTransactor) FileProve(opts *bind.TransactOpts, fileProve ProveFileProveParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "FileProve", fileProve)
}

// FileProve is a paid mutator transaction binding the contract method 0x8e270530.
//
// Solidity: function FileProve((bytes,bytes,uint256,address,uint64,uint64) fileProve) returns()
func (_Store *StoreSession) FileProve(fileProve ProveFileProveParams) (*types.Transaction, error) {
	return _Store.Contract.FileProve(&_Store.TransactOpts, fileProve)
}

// FileProve is a paid mutator transaction binding the contract method 0x8e270530.
//
// Solidity: function FileProve((bytes,bytes,uint256,address,uint64,uint64) fileProve) returns()
func (_Store *StoreTransactorSession) FileProve(fileProve ProveFileProveParams) (*types.Transaction, error) {
	return _Store.Contract.FileProve(&_Store.TransactOpts, fileProve)
}

// SectorProve is a paid mutator transaction binding the contract method 0x09cbe391.
//
// Solidity: function SectorProve((address,uint64,uint64,bytes) sectorProve) returns()
func (_Store *StoreTransactor) SectorProve(opts *bind.TransactOpts, sectorProve ProveSectorProveParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SectorProve", sectorProve)
}

// SectorProve is a paid mutator transaction binding the contract method 0x09cbe391.
//
// Solidity: function SectorProve((address,uint64,uint64,bytes) sectorProve) returns()
func (_Store *StoreSession) SectorProve(sectorProve ProveSectorProveParams) (*types.Transaction, error) {
	return _Store.Contract.SectorProve(&_Store.TransactOpts, sectorProve)
}

// SectorProve is a paid mutator transaction binding the contract method 0x09cbe391.
//
// Solidity: function SectorProve((address,uint64,uint64,bytes) sectorProve) returns()
func (_Store *StoreTransactorSession) SectorProve(sectorProve ProveSectorProveParams) (*types.Transaction, error) {
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

// SettleForFile is a paid mutator transaction binding the contract method 0xef5d51b4.
//
// Solidity: function SettleForFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo, (uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo, (address,address,uint64,uint256,bool) detail, (address,address,uint64,uint256,bool)[] details, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting) returns()
func (_Store *StoreTransactor) SettleForFile(opts *bind.TransactOpts, fileInfo FileInfo, nodeInfo NodeInfo, detail ProveDetail, details []ProveDetail, setting Setting) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SettleForFile", fileInfo, nodeInfo, detail, details, setting)
}

// SettleForFile is a paid mutator transaction binding the contract method 0xef5d51b4.
//
// Solidity: function SettleForFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo, (uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo, (address,address,uint64,uint256,bool) detail, (address,address,uint64,uint256,bool)[] details, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting) returns()
func (_Store *StoreSession) SettleForFile(fileInfo FileInfo, nodeInfo NodeInfo, detail ProveDetail, details []ProveDetail, setting Setting) (*types.Transaction, error) {
	return _Store.Contract.SettleForFile(&_Store.TransactOpts, fileInfo, nodeInfo, detail, details, setting)
}

// SettleForFile is a paid mutator transaction binding the contract method 0xef5d51b4.
//
// Solidity: function SettleForFile((bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)) fileInfo, (uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo, (address,address,uint64,uint256,bool) detail, (address,address,uint64,uint256,bool)[] details, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting) returns()
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

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _config, address _fs, address _node, address _pdp, address _sector) returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts, _config common.Address, _fs common.Address, _node common.Address, _pdp common.Address, _sector common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize", _config, _fs, _node, _pdp, _sector)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _config, address _fs, address _node, address _pdp, address _sector) returns()
func (_Store *StoreSession) Initialize(_config common.Address, _fs common.Address, _node common.Address, _pdp common.Address, _sector common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _fs, _node, _pdp, _sector)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _config, address _fs, address _node, address _pdp, address _sector) returns()
func (_Store *StoreTransactorSession) Initialize(_config common.Address, _fs common.Address, _node common.Address, _pdp common.Address, _sector common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _fs, _node, _pdp, _sector)
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
