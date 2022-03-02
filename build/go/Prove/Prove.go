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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"FileProveFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileProveNotFileOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeSectorProvedInTimeError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SectorProveFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"CheckNodeSectorProvedInTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteProveDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"NodeWallet\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"}],\"internalType\":\"structProve.FileProveParams\",\"name\":\"fileProve\",\"type\":\"tuple\"}],\"name\":\"FileProve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetPocProveList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"Height\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"PlotSize\",\"type\":\"uint64\"}],\"internalType\":\"structPocProve[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"GetProveDetailList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ChallengeHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"}],\"internalType\":\"structProve.SectorProveParams\",\"name\":\"sectorProve\",\"type\":\"tuple\"}],\"name\":\"SectorProve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"SetLastPunishmentHeightForNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail\",\"name\":\"detail\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"details\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"}],\"name\":\"SettleForFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"details\",\"type\":\"tuple[]\"}],\"name\":\"UpdateProveDetailList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveDetailNum\",\"type\":\"uint64\"}],\"internalType\":\"structProveDetailMeta\",\"name\":\"details\",\"type\":\"tuple\"}],\"name\":\"UpdateProveDetailMeta\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ChallengeHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"}],\"internalType\":\"structProve.SectorProveParams\",\"name\":\"sectorProve\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"}],\"name\":\"checkSectorProve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"getPocProve\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"Height\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"PlotSize\",\"type\":\"uint64\"}],\"internalType\":\"structPocProve\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractFileSystem\",\"name\":\"_fs\",\"type\":\"address\"},{\"internalType\":\"contractNode\",\"name\":\"_node\",\"type\":\"address\"},{\"internalType\":\"contractPDP\",\"name\":\"_pdp\",\"type\":\"address\"},{\"internalType\":\"contractSector\",\"name\":\"_sector\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"times\",\"type\":\"uint64\"}],\"name\":\"punishForSector\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"Height\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"PlotSize\",\"type\":\"uint64\"}],\"internalType\":\"structPocProve\",\"name\":\"prove\",\"type\":\"tuple\"}],\"name\":\"putPocProve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000805462010000600160501b0319166220000017905534801561002757600080fd5b5061603b80620000386000396000f3fe6080604052600436106100e85760003560e01c806352e061461161008a578063977fdfe211610059578063977fdfe21461028c5780639908f2bf146102b9578063bb4e4e42146102e6578063fda593511461030657600080fd5b806352e061461461020c578063648d225d1461022c5780637b85255a1461024c5780638e2705301461026c57600080fd5b80631459457a116100c65780631459457a146101615780631581d5451461018157806327ab4434146101a15780633ec0f5df146101c157600080fd5b806309cbe391146100ed5780630edec0f21461010f5780630fece86914610134575b600080fd5b3480156100f957600080fd5b5061010d610108366004614d7b565b610319565b005b34801561011b57600080fd5b50606060405161012b91906159e8565b60405180910390f35b34801561014057600080fd5b5061015461014f366004614daf565b6106f3565b60405161012b9190615a0a565b34801561016d57600080fd5b5061010d61017c366004614ae7565b610b8f565b34801561018d57600080fd5b5061010d61019c366004614caf565b610c9a565b3480156101ad57600080fd5b5061010d6101bc366004614a44565b610d48565b3480156101cd57600080fd5b5061010d6101dc366004614971565b6001600160a01b0390921660009081526007602090815260408083206001600160401b0390941683529290522055565b34801561021857600080fd5b5061010d610227366004614a10565b610dd6565b34801561023857600080fd5b5061010d610247366004614e0c565b610e47565b34801561025857600080fd5b5061010d610267366004614b90565b611147565b34801561027857600080fd5b5061010d610287366004614c29565b611550565b34801561029857600080fd5b506102ac6102a7366004614a10565b6121b7565b60405161012b91906159f9565b3480156102c557600080fd5b506102d96102d4366004614937565b6122e6565b60405161012b9190615bdc565b3480156102f257600080fd5b5061010d610301366004614aa1565b612384565b61010d610314366004614d01565b6123eb565b6002548151604051631bbc7f5f60e11b81526000926001600160a01b031691633778febe9161034b91906004016159ba565b60e06040518083038186803b15801561036357600080fd5b505afa158015610377573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061039b9190614c5d565b6004805460408051808201825286516001600160a01b0390811682526020808901516001600160401b0316908301529151632ba010d760e01b81529495506000949190921692632ba010d7926103f392909101615c31565b60006040518083038186803b15801561040b57600080fd5b505afa15801561041f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526104479190810190614ccd565b9050600080600a9054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561049957600080fd5b505afa1580156104ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d19190614e2a565b90508160c001514310156105045760016040516332ae228f60e21b81526004016104fb9190615aa8565b60405180910390fd5b8160c0015184604001516001600160401b0316146105895761056d6040518060400160405280601681526020017f4368616c6c656e6765486569676874206572726f723a0000000000000000000081525085604001516001600160401b03168460c00151612522565b60026040516332ae228f60e21b81526004016104fb9190615aa8565b600061059585846106f3565b9050806105c5576105a983858460016123eb565b60036040516332ae228f60e21b81526004016104fb9190615aa8565b60006105d284868561259c565b9050806105f55760056040516332ae228f60e21b81526004016104fb9190615aa8565b60a0840151610605574360a08501525b60c083015161061d906001600160401b031643615d1b565b60c0850152600480546040516311c2535560e11b81526001600160a01b0390911691632384a6aa9161065191889101615bfb565b600060405180830381600087803b15801561066b57600080fd5b505af115801561067f573d6000803e3d6000fd5b505050508361014001516106a95760046040516332ae228f60e21b81526004016104fb9190615aa8565b60006106b98560000151436122e6565b43602082015285516001600160a01b0316815260608601516001600160401b0316604082015290506106ea81610c9a565b50505050505050565b60408051608081018252600080825260606020830181905282840182815281840183815287516001600160a01b03908116865260e08801516001600160401b03908116909352845462010000900490921690526003549451631faf930160e31b8152929491939285929091169063fd7c980890610774908590600401615bbd565b60006040518083038186803b15801561078c57600080fd5b505afa1580156107a0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526107c891908101906149be565b90506108156040518060c0016040528060006001600160401b0316815260200160006001600160401b03168152602001606081526020016060815260200160608152602001606081525090565b61081d613194565b8681526040810182905261082f61325b565b6003546040517f9f9ca6440000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690639f9ca64490610878908590600401615bea565b60006040518083038186803b15801561089057600080fd5b505afa1580156108a4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526108cc9190810190614c7b565b90508060c001516108e65760009650505050505050610b89565b61092f6040518060e0016040528060006001600160401b031681526020016060815260200160608152602001606081526020016060815260200160608152602001606081525090565b6000808252604080860151602080850191909152845182850152840151606080850191909152848201516080808601919091529085015160a085015284015160c084015260035490516305b6ef6560e51b81526001600160a01b039091169063b6ddeca0906109a2908590600401615ca1565b60206040518083038186803b1580156109ba57600080fd5b505afa1580156109ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109f291906149f2565b905080610a0a57600098505050505050505050610b89565b89610140015115610b7c578260a001516102a001511580610a3f575060a08301516102c00151604001516001600160401b0316155b506040805160c081018252600060608083018281526080840183905260a08401839052835260208301529181019190915260a0808501516102c0015182528601516020820152865115610ac55786600081518110610aad57634e487b7160e01b600052603260045260246000fd5b60209081029190910101515163ffffffff1660408201525b6003546040517f2e19aeff0000000000000000000000000000000000000000000000000000000081526000916001600160a01b031690632e19aeff90610b0f908590600401615c90565b60206040518083038186803b158015610b2757600080fd5b505afa158015610b3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b5f91906149f2565b905080610b795760009a5050505050505050505050610b89565b50505b6001985050505050505050505b92915050565b600054610100900460ff16610baa5760005460ff1615610bae565b303b155b610bca5760405162461bcd60e51b81526004016104fb90615b1e565b600054610100900460ff16158015610bec576000805461ffff19166101011790555b600080547fffff0000000000000000000000000000000000000000ffffffffffffffffffff166a01000000000000000000006001600160a01b0389811691909102919091179091556001805473ffffffffffffffffffffffffffffffffffffffff1990811688841617909155600280548216878416179055600380548216868416179055600480549091169184169190911790558015610c92576000805461ff00191690555b505050505050565b600081600001518260200151604051602001610cb7929190615988565b604051602081830303815290604052905081600882604051610cd991906159ae565b90815260408051602092819003830190208351815473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0390911617815591830151600183015591909101516002909101805467ffffffffffffffff19166001600160401b039092169190911790555050565b6000600583604051610d5a91906159ae565b9081526020016040518091039020905060005b8251811015610dd0576000838281518110610d9857634e487b7160e01b600052603260045260246000fd5b60200260200101519050610dbb816000015182856128c49092919063ffffffff16565b50508080610dc890615ecb565b915050610d6d565b50505050565b600581604051610de691906159ae565b9081526040519081900360200190206000610e04600183018261336b565b60028201600090555050600681604051610e1e91906159ae565b90815260405190819003602001902080546fffffffffffffffffffffffffffffffff1916905550565b80516020820151600254604051631bbc7f5f60e11b81526000916001600160a01b031690633778febe90610e7f9086906004016159ba565b60e06040518083038186803b158015610e9757600080fd5b505afa158015610eab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ecf9190614c5d565b90504281608001516001600160401b03161015610eff576040516313420d3f60e11b815260040160405180910390fd5b6001600160401b038216610f26576040516313420d3f60e11b815260040160405180910390fd5b60048054604051632ba010d760e01b81526000926001600160a01b0390921691632ba010d791610f5891899101615c31565b60006040518083038186803b158015610f7057600080fd5b505afa158015610f84573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610fac9190810190614ccd565b90508061010001516001600160401b031660001415610fde576040516313420d3f60e11b815260040160405180910390fd5b6000805460808301516040517fc5c81b200000000000000000000000000000000000000000000000000000000081526a01000000000000000000009092046001600160a01b03169163c5c81b209161103891600401615a9a565b6101606040518083038186803b15801561105157600080fd5b505afa158015611065573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110899190614e2a565b90506000439050808260c001516001600160401b03168460c001516110ae9190615d1b565b10156110cd576040516313420d3f60e11b815260040160405180910390fd5b6001600160a01b03861660009081526007602090815260408083206001600160401b03891684529091528120549061110785858486612a42565b90506001600160401b038116611130576040516313420d3f60e11b815260040160405180910390fd5b61113c858786846123eb565b505050505050505050565b6000611154868584612b0f565b9050806001600160401b03168661014001516001600160401b0316101561119157600960405163cd442d6560e01b81526004016104fb9190615aa8565b80856060018181516111a39190615d33565b6001600160401b0316905250600254604051633206761b60e21b81526001600160a01b039091169063c819d86c906111df908890600401615bce565b600060405180830381600087803b1580156111f957600080fd5b505af115801561120d573d6000803e3d6000fd5b505050508086610140018181516112249190615dde565b6001600160401b031690525060006101c08701526001546040517f681850d70000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063681850d790611281908990600401615b7f565b600060405180830381600087803b15801561129b57600080fd5b505af11580156112af573d6000803e3d6000fd5b505050506000805b8451811015611311578481815181106112e057634e487b7160e01b600052603260045260246000fd5b602002602001015160800151156112ff57816112fb81615ee6565b9250505b8061130981615ecb565b9150506112b7565b50806001600160401b03166001141561138c5760018054604051639cd103e960e01b81526001600160a01b0390911691639cd103e991611359918b9160009190600401615b90565b600060405180830381600087803b15801561137357600080fd5b505af1158015611387573d6000803e3d6000fd5b505050505b61012087015161139d906001615d33565b6001600160401b0316816001600160401b03161415611482576101408701516001600160401b0316156114175786602001516001600160a01b03166108fc8861014001516001600160401b03169081150290604051600060405180830381858888f19350505050158015611415573d6000803e3d6000fd5b505b60018054604051639cd103e960e01b81526001600160a01b0390911691639cd103e99161144b918b91600090600401615b90565b600060405180830381600087803b15801561146557600080fd5b505af1158015611479573d6000803e3d6000fd5b50505050611505565b600154602088015188516040517f34fddaac0000000000000000000000000000000000000000000000000000000081526001600160a01b03909316926334fddaac926114d29290916004016159c8565b600060405180830381600087803b1580156114ec57600080fd5b505af1158015611500573d6000803e3d6000fd5b505050505b7f123b9998b2f027b02db4cb2eadef421fe9f1c21627569438588262d3a6baac6c6006438860a001518560405161153f9493929190615a29565b60405180910390a150505050505050565b600080600a9054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156115a057600080fd5b505afa1580156115b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115d89190614e2a565b60015483516040516337bf397560e21b81529293506000926001600160a01b039092169163defce5d49161160e91600401615a18565b60006040518083038186803b15801561162657600080fd5b505afa15801561163a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526116629190810190614b5c565b9050806102a00151156116c75780602001516001600160a01b031683606001516001600160a01b0316146116c2576040517f65fd380a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6117da565b610220810151600090815b815181101561173b5785606001516001600160a01b031682828151811061170957634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03161415611729576001925061173b565b8061173381615ecb565b9150506116d2565b50816117b65761024083015160005b81518110156117b35786606001516001600160a01b031682828151811061178157634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b031614156117a157600193506117b3565b806117ab81615ecb565b91505061174a565b50505b816117d757600160405163cd442d6560e01b81526004016104fb9190615aa8565b50505b60025460608401516040517f1a65374a0000000000000000000000000000000000000000000000000000000081526000926001600160a01b031691631a65374a9161182891906004016159ba565b60e06040518083038186803b15801561184057600080fd5b505afa158015611854573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118789190614c5d565b905060006118898360000151612c03565b90508460a001516001600160401b03166000141580156118ad575082610100015143105b156119375760005b81518110156119355785606001516001600160a01b03168282815181106118ec57634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516001600160a01b0316141561192357600260405163cd442d6560e01b81526004016104fb9190615aa8565b8061192d81615ecb565b9150506118b5565b505b60006119438685612d27565b90508061196657600360405163cd442d6560e01b81526004016104fb9190615aa8565b6040805160a08101825260008082526020820181905291810182905260608101829052608081018290528190819061010088015160005b8751811015611ad1578b606001516001600160a01b03168882815181106119d457634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516001600160a01b03161415611abf57826040015193508960e001516001600160401b0316846001600160401b03161480611a1957508143115b15611a2a5760016080840181905294505b8960e001516001600160401b0316846001600160401b03161115611a6457600460405163cd442d6560e01b81526004016104fb9190615aa8565b6000611a748b6101000151612fbb565b90508015611a9857600560405163cd442d6560e01b81526004016104fb9190615aa8565b60408401805190611aa882615ee6565b6001600160401b031690525060019650611ad19050565b80611ac981615ecb565b91505061199d565b5084611d6c57610120890151611ae8906001615d33565b6001600160401b031687511415611b1557600660405163cd442d6560e01b81526004016104fb9190615aa8565b8860a001518960800151611b299190615d87565b6001600160401b031688606001516001600160401b03161015611b6257600760405163cd442d6560e01b81526004016104fb9190615aa8565b8860a001518960800151611b769190615d87565b88606001818151611b879190615dde565b6001600160401b0316905250600254604051633206761b60e21b81526001600160a01b039091169063c819d86c90611bc3908b90600401615bce565b600060405180830381600087803b158015611bdd57600080fd5b505af1158015611bf1573d6000803e3d6000fd5b50505060c08901516001600160a01b0390811684526060808e0151909116602085015260016040850181905243918501919091526000608085018190528951909250611c3c91615d1b565b6001600160401b03811115611c6157634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015611cba57816020015b6040805160a081018252600080825260208083018290529282018190526060820181905260808201528252600019909201910181611c7f5790505b50905060005b8851811015611d2f57888181518110611ce957634e487b7160e01b600052603260045260246000fd5b6020026020010151828281518110611d1157634e487b7160e01b600052603260045260246000fd5b60200260200101819052508080611d2790615ecb565b915050611cc0565b50828160018351611d409190615dc3565b81518110611d5e57634e487b7160e01b600052603260045260246000fd5b602090810291909101015296505b8851611d789088610d48565b84611ffe576000600460009054906101000a90046001600160a01b03166001600160a01b0316632ba010d760405180604001604052808c60c001516001600160a01b031681526020018f60a001516001600160401b03168152506040518263ffffffff1660e01b8152600401611dee9190615c31565b60006040518083038186803b158015611e0657600080fd5b505afa158015611e1a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611e429190810190614ccd565b9050896102a001511515816101400151151514611e7557600860405163cd442d6560e01b81526004016104fb9190615aa8565b600480546040517f955f98b70000000000000000000000000000000000000000000000000000000081526001600160a01b039091169163955f98b791611ebf9185918f9101615c0c565b600060405180830381600087803b158015611ed957600080fd5b505af1158015611eed573d6000803e3d6000fd5b5050600480546040517fdcf749600000000000000000000000000000000000000000000000000000000081526001600160a01b03909116935063dcf749609250611f3991859101615bfb565b600060405180830381600087803b158015611f5357600080fd5b505af1158015611f67573d6000803e3d6000fd5b505050508060c0015160001415611f9b578960c001516001600160401b03168c60400151611f959190615d1b565b60c08201525b600480546040516311c2535560e11b81526001600160a01b0390911691632384a6aa91611fca91859101615bfb565b600060405180830381600087803b158015611fe457600080fd5b505af1158015611ff8573d6000803e3d6000fd5b50505050505b83156121665760a08b01516001600160401b031615612159576000600460009054906101000a90046001600160a01b03166001600160a01b0316632ba010d760405180604001604052808c60c001516001600160a01b031681526020018f60a001516001600160401b03168152506040518263ffffffff1660e01b81526004016120889190615c31565b60006040518083038186803b1580156120a057600080fd5b505afa1580156120b4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526120dc9190810190614ccd565b600480546040517e47c0030000000000000000000000000000000000000000000000000000000081529293506001600160a01b0316916247c003916121259185918f9101615c0c565b600060405180830381600087803b15801561213f57600080fd5b505af1158015612153573d6000803e3d6000fd5b50505050505b6121668989848a8e611147565b885160a08901516040517fd936063ae3cbbf2dbcd2adc837a997ab97adb09f24c566aa79d101e88928dea6926121a29260079243929190615a5e565b60405180910390a15050505050505050505050565b606060006005836040516121cb91906159ae565b90815260200160405180910390209050600081600201546001600160401b0381111561220757634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561226057816020015b6040805160a0810182526000808252602080830182905292820181905260608201819052608082015282526000199092019101816122255790505b509050816002015460001415612277579392505050565b600061228283612fd5565b90505b60018301548110156122de57600061229d8483612ff0565b915050808383815181106122c157634e487b7160e01b600052603260045260246000fd5b6020908102919091010152506122d783826130bd565b9050612285565b509392505050565b604080516060810182526000808252602082018190529181019190915260008383604051602001612318929190615988565b604051602081830303815290604052905060088160405161233991906159ae565b90815260408051918290036020908101832060608401835280546001600160a01b03168452600181015491840191909152600201546001600160401b03169082015291505092915050565b8060068360405161239591906159ae565b9081526040519081900360209081019091208251815493909201516001600160401b0390811668010000000000000000026fffffffffffffffffffffffffffffffff199094169216919091179190911790555050565b60006123f78386613132565b6124019083615d87565b9050806001600160401b031684600001516001600160401b0316106124435780846000018181516124329190615dde565b6001600160401b031690525061244a565b5060008084525b6001600160401b038116156124e557806001600160401b03163410156124825760405162461bcd60e51b81526004016104fb90615ae3565b600254604051633206761b60e21b81526001600160a01b039091169063c819d86c906124b2908790600401615bce565b600060405180830381600087803b1580156124cc57600080fd5b505af11580156124e0573d6000803e3d6000fd5b505050505b84516020808701516001600160a01b0390921660009081526007825260408082206001600160401b03909416825292909152204390555050505050565b61259783838360405160240161253a93929190615ab6565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f969cdd0300000000000000000000000000000000000000000000000000000000179052613153565b505050565b6000805b8461010001516001600160401b03168110156128b757600085610160015182815181106125dd57634e487b7160e01b600052603260045260246000fd5b60209081029190910101516001546040516337bf397560e21b81529192506000916001600160a01b039091169063defce5d49061261e908590600401615a18565b60006040518083038186803b15801561263657600080fd5b505afa15801561264a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526126729190810190614b5c565b9050600061268382600001516121b7565b6101008301516040805160a081018252600080825260208201819052918101829052606081018290526080810182905292935091829060005b85518110156127ae578c600001516001600160a01b0316868a815181106126f357634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516001600160a01b0316141561279c576001925085898151811061273257634e487b7160e01b600052603260045260246000fd5b602002602001015191506000826040015190508760e001516001600160401b0316816001600160401b0316148061276857508443115b156127795760016080840181905295505b6040830180519061278982615ee6565b6001600160401b03169052506127ae9050565b806127a681615ecb565b9150506126bc565b50816127c5576000985050505050505050506128bd565b85516127d19086610d48565b831561289d576127e4868c83888e611147565b600460009054906101000a90046001600160a01b03166001600160a01b03166247c0038d886040518363ffffffff1660e01b8152600401612826929190615c0c565b600060405180830381600087803b15801561284057600080fd5b505af1158015612854573d6000803e3d6000fd5b5050875160a08e01516040517fb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f459450612894935060019243929091615a5e565b60405180910390a15b5050505050505080806128af90615ecb565b9150506125a0565b50600190505b9392505050565b6001600160a01b038083166000908152602085815260408083208054865160018301805491881673ffffffffffffffffffffffffffffffffffffffff1990921691909117905592860151600282018054938801516001600160401b0316600160a01b027fffffffff000000000000000000000000000000000000000000000000000000009094169190961617919091179093556060840151600384015560808401516004909301805493151560ff199094169390931790925590801561298e5760019150506128bd565b50600180850180548083018255600091909152906129ad908290615d1b565b6001600160a01b038516600090815260208790526040902055600185018054859190839081106129ed57634e487b7160e01b600052603260045260246000fd5b60009182526020822001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03939093169290921790915560028601805491612a3383615ecb565b919050555060009150506128bd565b60c080840151908501516000919083612a646001600160401b03841683615d1b565b10612a7457600092505050612b07565b600082612a818387615dc3565b612a8b9190615d66565b905060008615612acf57612aa86001600160401b03851684615d1b565b871115612acb5783612aba8489615dc3565b612ac49190615d66565b9050612acf565b5060005b806001600160401b0316826001600160401b03161015612af6576000945050505050612b07565b612b008183615dde565b9450505050505b949350505050565b60018054604084015160009283926001600160a01b03169163cc76e80d918691612b399190615dde565b60008960a001518a60800151612b4f9190615d87565b8a6101a001518b6101000151612b659190615dc3565b6040518663ffffffff1660e01b8152600401612b85959493929190615c3f565b60606040518083038186803b158015612b9d57600080fd5b505afa158015612bb1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bd59190614e49565b9050806040015181602001518260000151612bf09190615d33565b612bfa9190615d33565b95945050505050565b60606000612c10836121b7565b905060005b8151811015612d205760025482516000916001600160a01b031690631a65374a90859085908110612c5657634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516040518263ffffffff1660e01b8152600401612c7e91906159ba565b60e06040518083038186803b158015612c9657600080fd5b505afa158015612caa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cce9190614c5d565b90508060c00151838381518110612cf557634e487b7160e01b600052603260045260246000fd5b60209081029190910101516001600160a01b0390911690525080612d1881615ecb565b915050612c15565b5092915050565b60c08101516000904390612d44906001600160401b031682615d1b565b84604001511180612d705750808360c001516001600160401b03168560400151612d6e9190615d1b565b105b15612d7f576000915050610b89565b6040805180820190915260608082526020820152612dc760405180608001604052806060815260200160006001600160401b0316815260200160608152602001606081525090565b604080516080808201835260008083526060602084018190528385018281528185018381528c8301516001600160a01b039081168752948c01516001600160401b039081169092526101808c015190911690526003549451631faf930160e31b815290949192919091169063fd7c980890612e46908590600401615bbd565b60006040518083038186803b158015612e5e57600080fd5b505afa158015612e72573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612e9a91908101906149be565b9050612ee56040518060e0016040528060006001600160401b031681526020016060815260200160608152602001606081526020016060815260200160608152602001606081525090565b60008082528551602080840191909152870151604080840191909152808701516060808501919091526080840185905287015160a0840152875160c084015260035490516305b6ef6560e51b81526001600160a01b039091169063b6ddeca090612f53908590600401615ca1565b60206040518083038186803b158015612f6b57600080fd5b505afa158015612f7f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fa391906149f2565b905080610b7c57600098505050505050505050610b89565b600081431115612fcd57506001919050565b506000919050565b600080612fe38360006130bd565b90506128bd600182615dc3565b6040805160a081018252600080825260208201819052918101829052606081018290526080810182905283600101838154811061303d57634e487b7160e01b600052603260045260246000fd5b6000918252602080832091909101546001600160a01b03908116808452968252604092839020835160a081018552600182015483168152600282015492831693810193909352600160a01b9091046001600160401b0316928201929092526003820154606082015260049091015460ff1615156080820152939492505050565b6000816130c981615ecb565b9250505b60018301548210801561311557508260010182815481106130fe57634e487b7160e01b600052603260045260246000fd5b600091825260209091200154600160a01b900460ff165b1561312c578161312481615ecb565b9250506130cd565b50919050565b600080600290506000613149858560600151613174565b612bfa9083615d87565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b6000620fa00082846060015161318a9190615d87565b6128bd9190615d66565b604080516101e08101909152600060608083018281526080840183905260a0840183905260c0840183905260e08401839052610100840183905261012084018390526101408401839052610160840183905261018084018390526101a08401929092526101c083015281908152602001606081526020016132566040518060c0016040528060006001600160401b0316815260200160006001600160401b03168152602001606081526020016060815260200160608152602001606081525090565b905290565b6040518060e00160405280606081526020016060815260200160608152602001606081526020016060815260200161335e604080516102e08101825260608082526000602083018190529282018190528082018390526080820183905260a0820183905260c0820183905260e08201839052610100820183905261012082018390526101408201839052610160820181905261018082018390526101a082018390526101c082018390526101e0820183905261020082018390526102208201819052610240820181905261026082015290610280820190815260006020808301829052604080516060810182528381529182018390528181019290925291015290565b8152600060209091015290565b5080546000825590600052602060002090810190613389919061338c565b50565b5b808211156133c35780547fffffffffffffffffffffff00000000000000000000000000000000000000000016815560010161338d565b5090565b60006133da6133d584615cce565b615cb2565b905080838252602082019050828560208602820111156133f957600080fd5b60005b85811015613425578161340f8882613762565b84525060209283019291909101906001016133fc565b5050509392505050565b600061343d6133d584615cce565b9050808382526020820190508285602086028201111561345c57600080fd5b60005b858110156134255781613472888261376d565b845250602092830192919091019060010161345f565b60006134966133d584615cce565b905080838252602082019050828560208602820111156134b557600080fd5b60005b858110156134255781356001600160401b038111156134d657600080fd5b8086016134e38982613896565b8552505060209283019291909101906001016134b8565b60006135086133d584615cce565b9050808382526020820190508285602086028201111561352757600080fd5b60005b858110156134255781516001600160401b0381111561354857600080fd5b80860161355589826138b7565b85525050602092830192919091019060010161352a565b600061357a6133d584615cce565b8381529050602081018260408502810186101561359657600080fd5b60005b8581101561342557816135ac888261390f565b84525060209092019160409190910190600101613599565b60006135d26133d584615cce565b905080838252602082019050828560208602820111156135f157600080fd5b60005b858110156134255781516001600160401b0381111561361257600080fd5b80860161361f8982613f36565b8552505060209283019291909101906001016135f4565b60006136446133d584615cce565b9050808382526020820190508285602086028201111561366357600080fd5b60005b858110156134255781516001600160401b0381111561368457600080fd5b8086016136918982613fa8565b855250506020928301929190910190600101613666565b60006136b66133d584615cce565b8381529050602081018260a0850281018610156136d257600080fd5b60005b8581101561342557816136e88882614381565b84525060209092019160a091909101906001016136d5565b600061370e6133d584615cf1565b90508281526020810184848401111561372657600080fd5b6122de848285615e67565b600061373f6133d584615cf1565b90508281526020810184848401111561375757600080fd5b6122de848285615e73565b8035610b8981615fa5565b8051610b8981615fa5565b600082601f83011261378957600080fd5b8135612b078482602086016133c7565b600082601f8301126137aa57600080fd5b8151612b0784826020860161342f565b600082601f8301126137cb57600080fd5b8135612b07848260208601613488565b600082601f8301126137ec57600080fd5b8151612b078482602086016134fa565b600082601f83011261380d57600080fd5b8151612b0784826020860161356c565b600082601f83011261382e57600080fd5b8151612b078482602086016135c4565b600082601f83011261384f57600080fd5b8151612b07848260208601613636565b600082601f83011261387057600080fd5b8135612b078482602086016136a8565b8035610b8981615fb9565b8051610b8981615fb9565b600082601f8301126138a757600080fd5b8135612b07848260208601613700565b600082601f8301126138c857600080fd5b8151612b07848260208601613731565b8035610b8981615fc1565b8035610b8981615fca565b8051610b8981615fca565b8035610b8981615fd7565b8051610b8981615fd7565b60006040828403121561392157600080fd5b61392b6040615cb2565b905060006139398484614916565b825250602061394a84848301614916565b60208301525092915050565b6000610320828403121561396957600080fd5b6139746102e0615cb2565b905081356001600160401b0381111561398c57600080fd5b61399884828501613896565b82525060206139a984848301613762565b60208301525060408201356001600160401b038111156139c857600080fd5b6139d484828501613896565b60408301525060606139e884828501614921565b60608301525060806139fc84828501614921565b60808301525060a0613a1084828501614921565b60a08301525060c0613a2484828501614921565b60c08301525060e0613a3884828501614921565b60e083015250610100613a4d84828501614900565b61010083015250610120613a6384828501614921565b61012083015250610140613a7984828501614921565b610140830152506101608201356001600160401b03811115613a9a57600080fd5b613aa684828501613896565b61016083015250610180613abc84828501614921565b610180830152506101a0613ad284828501614900565b6101a0830152506101c0613ae884828501613880565b6101c0830152506101e0613afe848285016138f9565b6101e083015250610200613b1484828501614921565b610200830152506102208201356001600160401b03811115613b3557600080fd5b613b4184828501613778565b610220830152506102408201356001600160401b03811115613b6257600080fd5b613b6e84828501613778565b610240830152506102608201356001600160401b03811115613b8f57600080fd5b613b9b84828501613896565b61026083015250610280613bb1848285016138e3565b610280830152506102a0613bc784828501613880565b6102a0830152506102c0613bdd8482850161426d565b6102c08301525092915050565b60006103208284031215613bfd57600080fd5b613c086102e0615cb2565b82519091506001600160401b03811115613c2157600080fd5b613c2d848285016138b7565b8252506020613c3e8484830161376d565b60208301525060408201516001600160401b03811115613c5d57600080fd5b613c69848285016138b7565b6040830152506060613c7d8482850161492c565b6060830152506080613c918482850161492c565b60808301525060a0613ca58482850161492c565b60a08301525060c0613cb98482850161492c565b60c08301525060e0613ccd8482850161492c565b60e083015250610100613ce28482850161490b565b61010083015250610120613cf88482850161492c565b61012083015250610140613d0e8482850161492c565b610140830152506101608201516001600160401b03811115613d2f57600080fd5b613d3b848285016138b7565b61016083015250610180613d518482850161492c565b610180830152506101a0613d678482850161490b565b6101a0830152506101c0613d7d8482850161388b565b6101c0830152506101e0613d9384828501613904565b6101e083015250610200613da98482850161492c565b610200830152506102208201516001600160401b03811115613dca57600080fd5b613dd684828501613799565b610220830152506102408201516001600160401b03811115613df757600080fd5b613e0384828501613799565b610240830152506102608201516001600160401b03811115613e2457600080fd5b613e30848285016138b7565b61026083015250610280613e46848285016138ee565b610280830152506102a0613e5c8482850161388b565b6102a0830152506102c0613bdd848285016142bc565b600060c08284031215613e8457600080fd5b613e8e60c0615cb2565b905081356001600160401b03811115613ea657600080fd5b613eb284828501613896565b82525060208201356001600160401b03811115613ece57600080fd5b613eda84828501613896565b6020830152506040613eee84828501614900565b6040830152506060613f0284828501613762565b6060830152506080613f1684828501614921565b60808301525060a0613f2a84828501614921565b60a08301525092915050565b600060608284031215613f4857600080fd5b613f526060615cb2565b90506000613f60848461492c565b8252506020613f718484830161492c565b60208301525060408201516001600160401b03811115613f9057600080fd5b613f9c848285016138b7565b60408301525092915050565b600060408284031215613fba57600080fd5b613fc46040615cb2565b90506000613fd2848461492c565b82525060208201516001600160401b03811115613fee57600080fd5b61394a8482850161381d565b600060e0828403121561400c57600080fd5b61401660e0615cb2565b905060006140248484614921565b825250602061403584848301614921565b602083015250604061404984828501614921565b604083015250606061405d84828501614921565b606083015250608061407184828501614921565b60808301525060a061408584828501613762565b60a08301525060c061409984828501613762565b60c08301525092915050565b600060e082840312156140b757600080fd5b6140c160e0615cb2565b905060006140cf848461492c565b82525060206140e08484830161492c565b60208301525060406140f48482850161492c565b60408301525060606141088482850161492c565b606083015250608061411c8482850161492c565b60808301525060a06141308482850161376d565b60a08301525060c06140998482850161376d565b600060e0828403121561415657600080fd5b61416060e0615cb2565b82519091506001600160401b0381111561417957600080fd5b614185848285016137db565b82525060208201516001600160401b038111156141a157600080fd5b6141ad848285016137db565b60208301525060408201516001600160401b038111156141cc57600080fd5b6141d8848285016137fc565b60408301525060608201516001600160401b038111156141f757600080fd5b6142038482850161383e565b60608301525060808201516001600160401b0381111561422257600080fd5b61422e848285016138b7565b60808301525060a08201516001600160401b0381111561424d57600080fd5b61425984828501613bea565b60a08301525060c06140998482850161388b565b60006060828403121561427f57600080fd5b6142896060615cb2565b905060006142978484614921565b82525060206142a884848301614921565b6020830152506040613f9c84828501614921565b6000606082840312156142ce57600080fd5b6142d86060615cb2565b905060006142e6848461492c565b82525060206142f78484830161492c565b6020830152506040613f9c8482850161492c565b60006060828403121561431d57600080fd5b6143276060615cb2565b905060006143358484613762565b82525060206142a884848301614900565b60006040828403121561435857600080fd5b6143626040615cb2565b905060006143708484614921565b825250602061394a84848301614921565b600060a0828403121561439357600080fd5b61439d60a0615cb2565b905060006143ab8484613762565b82525060206143bc84848301613762565b60208301525060406143d084828501614921565b60408301525060606143e484828501614900565b60608301525060806143f884828501613880565b60808301525092915050565b6000610180828403121561441757600080fd5b614422610180615cb2565b905060006144308484613762565b825250602061444184848301614921565b602083015250604061445584828501614921565b604083015250606061446984828501614921565b606083015250608061447d848285016138e3565b60808301525060a061449184828501614900565b60a08301525060c06144a584828501614900565b60c08301525060e06144b984828501614921565b60e0830152506101006144ce84828501614921565b610100830152506101206144e484828501614921565b610120830152506101406144fa84828501613880565b610140830152506101608201356001600160401b0381111561451b57600080fd5b614527848285016137ba565b6101608301525092915050565b6000610180828403121561454757600080fd5b614552610180615cb2565b90506000614560848461376d565b82525060206145718484830161492c565b60208301525060406145858482850161492c565b60408301525060606145998482850161492c565b60608301525060806145ad848285016138ee565b60808301525060a06145c18482850161490b565b60a08301525060c06145d58482850161490b565b60c08301525060e06145e98482850161492c565b60e0830152506101006145fe8482850161492c565b610100830152506101206146148482850161492c565b6101208301525061014061462a8482850161388b565b610140830152506101608201516001600160401b0381111561464b57600080fd5b614527848285016137db565b60006080828403121561466957600080fd5b6146736080615cb2565b905060006146818484613762565b825250602061469284848301614921565b60208301525060406146a684828501614921565b60408301525060608201356001600160401b038111156146c557600080fd5b6146d184828501613896565b60608301525092915050565b6000604082840312156146ef57600080fd5b6146f96040615cb2565b905060006143708484613762565b6000610160828403121561471a57600080fd5b614725610160615cb2565b905060006147338484614921565b825250602061474484848301614921565b602083015250604061475884828501614921565b604083015250606061476c84828501614921565b606083015250608061478084828501614921565b60808301525060a061479484828501614921565b60a08301525060c06147a884828501614921565b60c08301525060e06147bc84828501614921565b60e0830152506101006147d184828501614921565b610100830152506101206147e784828501614921565b610120830152506101406147fd84828501614921565b6101408301525092915050565b6000610160828403121561481d57600080fd5b614828610160615cb2565b90506000614836848461492c565b82525060206148478484830161492c565b602083015250604061485b8482850161492c565b604083015250606061486f8482850161492c565b60608301525060806148838482850161492c565b60808301525060a06148978482850161492c565b60a08301525060c06148ab8482850161492c565b60c08301525060e06148bf8482850161492c565b60e0830152506101006148d48482850161492c565b610100830152506101206148ea8482850161492c565b610120830152506101406147fd8482850161492c565b8035610b8981615fe4565b8051610b8981615fe4565b8051610b8981615fea565b8035610b8981615ff6565b8051610b8981615ff6565b6000806040838503121561494a57600080fd5b60006149568585613762565b925050602061496785828601614900565b9150509250929050565b60008060006060848603121561498657600080fd5b60006149928686613762565b93505060206149a386828701614921565b92505060406149b486828701614900565b9150509250925092565b6000602082840312156149d057600080fd5b81516001600160401b038111156149e657600080fd5b612b07848285016137fc565b600060208284031215614a0457600080fd5b6000612b07848461388b565b600060208284031215614a2257600080fd5b81356001600160401b03811115614a3857600080fd5b612b0784828501613896565b60008060408385031215614a5757600080fd5b82356001600160401b03811115614a6d57600080fd5b614a7985828601613896565b92505060208301356001600160401b03811115614a9557600080fd5b6149678582860161385f565b60008060608385031215614ab457600080fd5b82356001600160401b03811115614aca57600080fd5b614ad685828601613896565b925050602061496785828601614346565b600080600080600060a08688031215614aff57600080fd5b6000614b0b88886138d8565b9550506020614b1c888289016138d8565b9450506040614b2d888289016138d8565b9350506060614b3e888289016138d8565b9250506080614b4f888289016138d8565b9150509295509295909350565b600060208284031215614b6e57600080fd5b81516001600160401b03811115614b8457600080fd5b612b0784828501613bea565b60008060008060006103208688031215614ba957600080fd5b85356001600160401b03811115614bbf57600080fd5b614bcb88828901613956565b9550506020614bdc88828901613ffa565b945050610100614bee88828901614381565b9350506101a08601356001600160401b03811115614c0b57600080fd5b614c178882890161385f565b9250506101c0614b4f88828901614707565b600060208284031215614c3b57600080fd5b81356001600160401b03811115614c5157600080fd5b612b0784828501613e72565b600060e08284031215614c6f57600080fd5b6000612b0784846140a5565b600060208284031215614c8d57600080fd5b81516001600160401b03811115614ca357600080fd5b612b0784828501614144565b600060608284031215614cc157600080fd5b6000612b07848461430b565b600060208284031215614cdf57600080fd5b81516001600160401b03811115614cf557600080fd5b612b0784828501614534565b6000806000806102808587031215614d1857600080fd5b84356001600160401b03811115614d2e57600080fd5b614d3a87828801614404565b9450506020614d4b87828801613ffa565b935050610100614d5d87828801614707565b925050610260614d6f87828801614921565b91505092959194509250565b600060208284031215614d8d57600080fd5b81356001600160401b03811115614da357600080fd5b612b0784828501614657565b60008060408385031215614dc257600080fd5b82356001600160401b03811115614dd857600080fd5b614de485828601614657565b92505060208301356001600160401b03811115614e0057600080fd5b61496785828601614404565b600060408284031215614e1e57600080fd5b6000612b0784846146dd565b60006101608284031215614e3d57600080fd5b6000612b07848461480a565b600060608284031215614e5b57600080fd5b6000612b0784846142bc565b6000614e738383614edb565b505060200190565b60006128bd838361511e565b6000614e938383615196565b505060400190565b60006128bd838361540c565b60006128bd838361544b565b6000614ebf8383615531565b505060600190565b6000614ed383836155a1565b505060a00190565b614ee481615dfb565b82525050565b614ee4614ef682615dfb565b615f0b565b6000614f05825190565b80845260209384019383018060005b83811015614f39578151614f288882614e67565b975060208301925050600101614f14565b509495945050505050565b6000614f4e825190565b80845260208401935083602082028501614f688560200190565b8060005b85811015614f9d5784840389528151614f858582614e7b565b94506020830160209a909a0199925050600101614f6c565b5091979650505050505050565b6000614fb4825190565b80845260209384019383018060005b83811015614f39578151614fd78882614e87565b975060208301925050600101614fc3565b6000614ff2825190565b8084526020840193508360208202850161500c8560200190565b8060005b85811015614f9d57848403895281516150298582614e9b565b94506020830160209a909a0199925050600101615010565b600061504b825190565b808452602084019350836020820285016150658560200190565b8060005b85811015614f9d57848403895281516150828582614ea7565b94506020830160209a909a0199925050600101615069565b60006150a4825190565b80845260209384019383018060005b83811015614f395781516150c78882614eb3565b9750602083019250506001016150b3565b60006150e2825190565b80845260209384019383018060005b83811015614f395781516151058882614ec7565b9750602083019250506001016150f1565b801515614ee4565b6000615128825190565b80845260208401935061513f818560208601615e73565b601f01601f19169290920192915050565b600061515a825190565b615168818560208601615e73565b9290920192915050565b614ee481615e35565b614ee481615e40565b614ee481615e4b565b614ee481615e56565b805160408301906151a7848261596d565b506020820151610dd0602085018261596d565b8051610320808452600091908401906151d3828261511e565b91505060208301516151e86020860182614edb565b5060408301518482036040860152615200828261511e565b91505060608301516152156060860182615979565b5060808301516152286080860182615979565b5060a083015161523b60a0860182615979565b5060c083015161524e60c0860182615979565b5060e083015161526160e0860182615979565b50610100830151615276610100860182615967565b5061012083015161528b610120860182615979565b506101408301516152a0610140860182615979565b506101608301518482036101608601526152ba828261511e565b9150506101808301516152d1610180860182615979565b506101a08301516152e66101a0860182615967565b506101c08301516152fb6101c0860182615116565b506101e08301516153106101e0860182615184565b50610200830151615325610200860182615979565b5061022083015184820361022086015261533f8282614efb565b91505061024083015184820361024086015261535b8282614efb565b915050610260830151848203610260860152615377828261511e565b91505061028083015161538e61028086018261517b565b506102a08301516153a36102a0860182615116565b506102c08301516122de6102c08601826154fa565b805160009060808401906153cc8582614edb565b50602083015184820360208601526153e4828261511e565b91505060408301516153f96040860182615979565b5060608301516122de6060860182615979565b805160009060608401906154208582615979565b5060208301516154336020860182615979565b5060408301518482036040860152612bfa828261511e565b8051600090604084019061545f8582615979565b5060208301518482036020860152612bfa8282614fe8565b805160e08301906154888482615979565b50602082015161549b6020850182615979565b5060408201516154ae6040850182615979565b5060608201516154c16060850182615979565b5060808201516154d46080850182615979565b5060a08201516154e760a0850182614edb565b5060c0820151610dd060c0850182614edb565b8051606083019061550b8482615979565b50602082015161551e6020850182615979565b506040820151610dd06040850182615979565b805160608301906155428482614edb565b50602082015161551e6020850182615967565b805160608084526000919084019061556d82826155fe565b915050602083015184820360208601526155878282614faa565b91505060408301518482036040860152612bfa82826156f1565b805160a08301906155b28482614edb565b5060208201516155c56020850182614edb565b5060408201516155d86040850182615979565b5060608201516155eb6060850182615967565b506080820151610dd06080850182615116565b80516000906101808401906156138582614edb565b5060208301516156266020860182615979565b5060408301516156396040860182615979565b50606083015161564c6060860182615979565b50608083015161565f608086018261517b565b5060a083015161567260a0860182615967565b5060c083015161568560c0860182615967565b5060e083015161569860e0860182615979565b506101008301516156ad610100860182615979565b506101208301516156c2610120860182615979565b506101408301516156d7610140860182615116565b50610160830151848203610160860152612bfa8282614f44565b805160009060c08401906157058582615979565b5060208301516157186020860182615979565b5060408301518482036040860152615730828261511e565b9150506060830151848203606086015261574a828261511e565b915050608083015184820360808601526157648282615041565b91505060a083015184820360a0860152612bfa828261511e565b8051604083019061578f8482614edb565b506020820151610dd06020850182615979565b80516101608301906157b48482615979565b5060208201516157c76020850182615979565b5060408201516157da6040850182615979565b5060608201516157ed6060850182615979565b5060808201516158006080850182615979565b5060a082015161581360a0850182615979565b5060c082015161582660c0850182615979565b5060e082015161583960e0850182615979565b5061010082015161584e610100850182615979565b50610120820151615863610120850182615979565b50610140820151610dd0610140850182615979565b805160009060a084019061588c85826154fa565b50602083015184820360608601526158a4828261511e565b91505060408301516122de6080860182615979565b805160009060e08401906158cd8582615979565b50602083015184820360208601526158e5828261511e565b915050604083015184820360408601526158ff8282614f44565b915050606083015184820360608601526159198282614f44565b915050608083015184820360808601526159338282614faa565b91505060a083015184820360a086015261594d8282615041565b91505060c083015184820360c0860152612bfa828261511e565b80614ee4565b63ffffffff8116614ee4565b6001600160401b038116614ee4565b60006159948285614eea565b6014820191506159a48284615967565b5060200192915050565b60006128bd8284615150565b60208101610b898284614edb565b604081016159d68285614edb565b8181036020830152612b07818461511e565b602080825281016128bd818461509a565b602080825281016128bd81846150d8565b60208101610b898284615116565b602080825281016128bd818461511e565b60808101615a378287615172565b615a446020830186615967565b615a516040830185614edb565b612bfa6060830184615979565b60808101615a6c8287615172565b615a796020830186615967565b8181036040830152615a8b818561511e565b9050612bfa6060830184614edb565b60208101610b89828461517b565b60208101610b89828461518d565b60608082528101615ac7818661511e565b9050615ad66020830185615967565b612b076040830184615967565b60208082528101610b8981601681527f50756e697368466f72536563746f72206661696c656400000000000000000000602082015260400190565b60208082528101610b8981602e81527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160208201527f647920696e697469616c697a6564000000000000000000000000000000000000604082015260600190565b602080825281016128bd81846151ba565b60608082528101615ba181866151ba565b9050615bb06020830185615116565b612b076040830184615116565b602080825281016128bd81846153b8565b60e08101610b898284615477565b60608101610b898284615531565b602080825281016128bd8184615555565b602080825281016128bd81846155fe565b60408082528101615c1d81856155fe565b90508181036020830152612b0781846151ba565b60408101610b89828461577e565b6101e08101615c4e82886157a2565b615c5c610160830187615979565b615c6a61018083018661518d565b615c786101a0830185615979565b615c866101c0830184615979565b9695505050505050565b602080825281016128bd8184615878565b602080825281016128bd81846158b9565b6000615cbd60405190565b9050615cc98282615e9f565b919050565b60006001600160401b03821115615ce757615ce7615f5f565b5060209081020190565b60006001600160401b03821115615d0a57615d0a615f5f565b601f19601f83011660200192915050565b60008219821115615d2e57615d2e615f1d565b500190565b60006001600160401b03821691506001600160401b0383169250826001600160401b0303821115615d2e57615d2e615f1d565b6001600160401b039182169116600082615d8257615d82615f33565b500490565b60006001600160401b03821691506001600160401b0383169250816001600160401b030483118215151615615dbe57615dbe615f1d565b500290565b6000825b925082821015615dd957615dd9615f1d565b500390565b60006001600160401b03821691506001600160401b038316615dc7565b60006001600160a01b038216610b89565b6000610b8982615dfb565b80615cc981615f75565b80615cc981615f85565b80615cc981615f95565b6000610b8982615e17565b6000610b8982615e21565b6000610b8982615e2b565b60006001600160401b038216610b89565b82818337506000910152565b60005b83811015615e8e578181015183820152602001615e76565b83811115610dd05750506000910152565b601f19601f83011681018181106001600160401b0382111715615ec457615ec4615f5f565b6040525050565b6000600019821415615edf57615edf615f1d565b5060010190565b60006001600160401b03821691506001600160401b03821415615edf57615edf615f1d565b6000610b89826000610b898260601b90565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600a811061338957613389615f49565b6003811061338957613389615f49565b6002811061338957613389615f49565b615fae81615dfb565b811461338957600080fd5b801515615fae565b615fae81615e0c565b6003811061338957600080fd5b6002811061338957600080fd5b80615fae565b63ffffffff8116615fae565b6001600160401b038116615fae56fea2646970667358221220ee926211e424659fed73379027e0bdadb204bdd13454367ca67ef7207c43601164736f6c63430008040033",
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
