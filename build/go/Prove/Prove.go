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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"FileProveFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileProveNotFileOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeSectorProvedInTimeError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SectorProveFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"CheckNodeSectorProvedInTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteProveDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"NodeWallet\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"}],\"internalType\":\"structProve.FileProveParams\",\"name\":\"fileProve\",\"type\":\"tuple\"}],\"name\":\"FileProve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetPocProveList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"Height\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"PlotSize\",\"type\":\"uint64\"}],\"internalType\":\"structPocProve[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"GetProveDetailList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ChallengeHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"}],\"internalType\":\"structProve.SectorProveParams\",\"name\":\"sectorProve\",\"type\":\"tuple\"}],\"name\":\"SectorProve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"SetLastPunishmentHeightForNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail\",\"name\":\"detail\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"details\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"}],\"name\":\"SettleForFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"details\",\"type\":\"tuple[]\"}],\"name\":\"UpdateProveDetailList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveDetailNum\",\"type\":\"uint64\"}],\"internalType\":\"structProveDetailMeta\",\"name\":\"details\",\"type\":\"tuple\"}],\"name\":\"UpdateProveDetailMeta\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ChallengeHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"}],\"internalType\":\"structProve.SectorProveParams\",\"name\":\"sectorProve\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"}],\"name\":\"checkSectorProve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"getPocProve\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"Height\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"PlotSize\",\"type\":\"uint64\"}],\"internalType\":\"structPocProve\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractFileSystem\",\"name\":\"_fs\",\"type\":\"address\"},{\"internalType\":\"contractNode\",\"name\":\"_node\",\"type\":\"address\"},{\"internalType\":\"contractPDP\",\"name\":\"_pdp\",\"type\":\"address\"},{\"internalType\":\"contractSector\",\"name\":\"_sector\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"times\",\"type\":\"uint64\"}],\"name\":\"punishForSector\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"Height\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"PlotSize\",\"type\":\"uint64\"}],\"internalType\":\"structPocProve\",\"name\":\"prove\",\"type\":\"tuple\"}],\"name\":\"putPocProve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000805462010000600160501b0319166220000017905534801561002757600080fd5b50615fbe80620000386000396000f3fe6080604052600436106100e85760003560e01c806352e061461161008a5780639908f2bf116100595780639908f2bf14610299578063bb4e4e42146102c6578063ef5d51b4146102e6578063fda593511461030657600080fd5b806352e061461461020c578063648d225d1461022c5780638e2705301461024c578063977fdfe21461026c57600080fd5b80631459457a116100c65780631459457a146101615780631581d5451461018157806327ab4434146101a15780633ec0f5df146101c157600080fd5b806309cbe391146100ed5780630edec0f21461010f5780630fece86914610134575b600080fd5b3480156100f957600080fd5b5061010d610108366004614d1a565b610319565b005b34801561011b57600080fd5b50606060405161012b919061596b565b60405180910390f35b34801561014057600080fd5b5061015461014f366004614d4e565b6106f3565b60405161012b919061598d565b34801561016d57600080fd5b5061010d61017c366004614a86565b610b8f565b34801561018d57600080fd5b5061010d61019c366004614c4e565b610c9a565b3480156101ad57600080fd5b5061010d6101bc3660046149e3565b610d48565b3480156101cd57600080fd5b5061010d6101dc366004614910565b6001600160a01b0390921660009081526007602090815260408083206001600160401b0390941683529290522055565b34801561021857600080fd5b5061010d6102273660046149af565b610dd6565b34801561023857600080fd5b5061010d610247366004614dab565b610e47565b34801561025857600080fd5b5061010d610267366004614bc8565b611147565b34801561027857600080fd5b5061028c6102873660046149af565b611db0565b60405161012b919061597c565b3480156102a557600080fd5b506102b96102b43660046148d6565b611edf565b60405161012b9190615b5f565b3480156102d257600080fd5b5061010d6102e1366004614a40565b611f7d565b3480156102f257600080fd5b5061010d610301366004614b2f565b611fe4565b61010d610314366004614ca0565b6123ed565b6002548151604051631bbc7f5f60e11b81526000926001600160a01b031691633778febe9161034b919060040161593d565b60e06040518083038186803b15801561036357600080fd5b505afa158015610377573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061039b9190614bfc565b6004805460408051808201825286516001600160a01b0390811682526020808901516001600160401b0316908301529151632ba010d760e01b81529495506000949190921692632ba010d7926103f392909101615bb4565b60006040518083038186803b15801561040b57600080fd5b505afa15801561041f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526104479190810190614c6c565b9050600080600a9054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561049957600080fd5b505afa1580156104ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d19190614dc9565b90508160c001514310156105045760016040516332ae228f60e21b81526004016104fb9190615a2b565b60405180910390fd5b8160c0015184604001516001600160401b0316146105895761056d6040518060400160405280601681526020017f4368616c6c656e6765486569676874206572726f723a0000000000000000000081525085604001516001600160401b03168460c00151612524565b60026040516332ae228f60e21b81526004016104fb9190615a2b565b600061059585846106f3565b9050806105c5576105a983858460016123ed565b60036040516332ae228f60e21b81526004016104fb9190615a2b565b60006105d284868561259e565b9050806105f55760056040516332ae228f60e21b81526004016104fb9190615a2b565b60a0840151610605574360a08501525b60c083015161061d906001600160401b031643615c9e565b60c0850152600480546040516311c2535560e11b81526001600160a01b0390911691632384a6aa9161065191889101615b7e565b600060405180830381600087803b15801561066b57600080fd5b505af115801561067f573d6000803e3d6000fd5b505050508361014001516106a95760046040516332ae228f60e21b81526004016104fb9190615a2b565b60006106b9856000015143611edf565b43602082015285516001600160a01b0316815260608601516001600160401b0316604082015290506106ea81610c9a565b50505050505050565b60408051608081018252600080825260606020830181905282840182815281840183815287516001600160a01b03908116865260e08801516001600160401b03908116909352845462010000900490921690526003549451631faf930160e31b8152929491939285929091169063fd7c980890610774908590600401615b40565b60006040518083038186803b15801561078c57600080fd5b505afa1580156107a0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526107c8919081019061495d565b90506108156040518060c0016040528060006001600160401b0316815260200160006001600160401b03168152602001606081526020016060815260200160608152602001606081525090565b61081d613197565b8681526040810182905261082f61325e565b6003546040517f9f9ca6440000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690639f9ca64490610878908590600401615b6d565b60006040518083038186803b15801561089057600080fd5b505afa1580156108a4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526108cc9190810190614c1a565b90508060c001516108e65760009650505050505050610b89565b61092f6040518060e0016040528060006001600160401b031681526020016060815260200160608152602001606081526020016060815260200160608152602001606081525090565b6000808252604080860151602080850191909152845182850152840151606080850191909152848201516080808601919091529085015160a085015284015160c084015260035490516305b6ef6560e51b81526001600160a01b039091169063b6ddeca0906109a2908590600401615c24565b60206040518083038186803b1580156109ba57600080fd5b505afa1580156109ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109f29190614991565b905080610a0a57600098505050505050505050610b89565b89610140015115610b7c578260a0015161028001511580610a3f575060a08301516102a00151604001516001600160401b0316155b506040805160c081018252600060608083018281526080840183905260a08401839052835260208301529181019190915260a0808501516102a0015182528601516020820152865115610ac55786600081518110610aad57634e487b7160e01b600052603260045260246000fd5b60209081029190910101515163ffffffff1660408201525b6003546040517f2e19aeff0000000000000000000000000000000000000000000000000000000081526000916001600160a01b031690632e19aeff90610b0f908590600401615c13565b60206040518083038186803b158015610b2757600080fd5b505afa158015610b3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b5f9190614991565b905080610b795760009a5050505050505050505050610b89565b50505b6001985050505050505050505b92915050565b600054610100900460ff16610baa5760005460ff1615610bae565b303b155b610bca5760405162461bcd60e51b81526004016104fb90615aa1565b600054610100900460ff16158015610bec576000805461ffff19166101011790555b600080547fffff0000000000000000000000000000000000000000ffffffffffffffffffff166a01000000000000000000006001600160a01b0389811691909102919091179091556001805473ffffffffffffffffffffffffffffffffffffffff1990811688841617909155600280548216878416179055600380548216868416179055600480549091169184169190911790558015610c92576000805461ff00191690555b505050505050565b600081600001518260200151604051602001610cb792919061590b565b604051602081830303815290604052905081600882604051610cd99190615931565b90815260408051602092819003830190208351815473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0390911617815591830151600183015591909101516002909101805467ffffffffffffffff19166001600160401b039092169190911790555050565b6000600583604051610d5a9190615931565b9081526020016040518091039020905060005b8251811015610dd0576000838281518110610d9857634e487b7160e01b600052603260045260246000fd5b60200260200101519050610dbb816000015182856128c79092919063ffffffff16565b50508080610dc890615e4e565b915050610d6d565b50505050565b600581604051610de69190615931565b9081526040519081900360200190206000610e046001830182613364565b60028201600090555050600681604051610e1e9190615931565b90815260405190819003602001902080546fffffffffffffffffffffffffffffffff1916905550565b80516020820151600254604051631bbc7f5f60e11b81526000916001600160a01b031690633778febe90610e7f90869060040161593d565b60e06040518083038186803b158015610e9757600080fd5b505afa158015610eab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ecf9190614bfc565b90504281608001516001600160401b03161015610eff576040516313420d3f60e11b815260040160405180910390fd5b6001600160401b038216610f26576040516313420d3f60e11b815260040160405180910390fd5b60048054604051632ba010d760e01b81526000926001600160a01b0390921691632ba010d791610f5891899101615bb4565b60006040518083038186803b158015610f7057600080fd5b505afa158015610f84573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610fac9190810190614c6c565b90508061010001516001600160401b031660001415610fde576040516313420d3f60e11b815260040160405180910390fd5b6000805460808301516040517fc5c81b200000000000000000000000000000000000000000000000000000000081526a01000000000000000000009092046001600160a01b03169163c5c81b209161103891600401615a1d565b6101606040518083038186803b15801561105157600080fd5b505afa158015611065573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110899190614dc9565b90506000439050808260c001516001600160401b03168460c001516110ae9190615c9e565b10156110cd576040516313420d3f60e11b815260040160405180910390fd5b6001600160a01b03861660009081526007602090815260408083206001600160401b03891684529091528120549061110785858486612a45565b90506001600160401b038116611130576040516313420d3f60e11b815260040160405180910390fd5b61113c858786846123ed565b505050505050505050565b600080600a9054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561119757600080fd5b505afa1580156111ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111cf9190614dc9565b60015483516040516337bf397560e21b81529293506000926001600160a01b039092169163defce5d4916112059160040161599b565b60006040518083038186803b15801561121d57600080fd5b505afa158015611231573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526112599190810190614afb565b9050806102800151156112be5780602001516001600160a01b031683606001516001600160a01b0316146112b9576040517f65fd380a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113d1565b610220810151600090815b81518110156113325785606001516001600160a01b031682828151811061130057634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b031614156113205760019250611332565b8061132a81615e4e565b9150506112c9565b50816113ad5761024083015160005b81518110156113aa5786606001516001600160a01b031682828151811061137857634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b0316141561139857600193506113aa565b806113a281615e4e565b915050611341565b50505b816113ce57600160405163cd442d6560e01b81526004016104fb9190615a2b565b50505b60025460608401516040517f1a65374a0000000000000000000000000000000000000000000000000000000081526000926001600160a01b031691631a65374a9161141f919060040161593d565b60e06040518083038186803b15801561143757600080fd5b505afa15801561144b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061146f9190614bfc565b905060006114808360000151612b12565b90508460a001516001600160401b03166000141580156114a4575082610100015143105b1561152e5760005b815181101561152c5785606001516001600160a01b03168282815181106114e357634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516001600160a01b0316141561151a57600260405163cd442d6560e01b81526004016104fb9190615a2b565b8061152481615e4e565b9150506114ac565b505b600061153a8685612c36565b90508061155d57600360405163cd442d6560e01b81526004016104fb9190615a2b565b6040805160a08101825260008082526020820181905291810182905260608101829052608081018290528190819061010088015160005b87518110156116c8578b606001516001600160a01b03168882815181106115cb57634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516001600160a01b031614156116b657826040015193508960e001516001600160401b0316846001600160401b0316148061161057508143115b156116215760016080840181905294505b8960e001516001600160401b0316846001600160401b0316111561165b57600460405163cd442d6560e01b81526004016104fb9190615a2b565b600061166b8b6101000151612eca565b9050801561168f57600560405163cd442d6560e01b81526004016104fb9190615a2b565b6040840180519061169f82615e69565b6001600160401b0316905250600196506116c89050565b806116c081615e4e565b915050611594565b5084611963576101208901516116df906001615cb6565b6001600160401b03168751141561170c57600660405163cd442d6560e01b81526004016104fb9190615a2b565b8860a0015189608001516117209190615d0a565b6001600160401b031688606001516001600160401b0316101561175957600760405163cd442d6560e01b81526004016104fb9190615a2b565b8860a00151896080015161176d9190615d0a565b8860600181815161177e9190615d61565b6001600160401b0316905250600254604051633206761b60e21b81526001600160a01b039091169063c819d86c906117ba908b90600401615b51565b600060405180830381600087803b1580156117d457600080fd5b505af11580156117e8573d6000803e3d6000fd5b50505060c08901516001600160a01b0390811684526060808e015190911660208501526001604085018190524391850191909152600060808501819052895190925061183391615c9e565b6001600160401b0381111561185857634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156118b157816020015b6040805160a0810182526000808252602080830182905292820181905260608201819052608082015282526000199092019101816118765790505b50905060005b8851811015611926578881815181106118e057634e487b7160e01b600052603260045260246000fd5b602002602001015182828151811061190857634e487b7160e01b600052603260045260246000fd5b6020026020010181905250808061191e90615e4e565b9150506118b7565b508281600183516119379190615d46565b8151811061195557634e487b7160e01b600052603260045260246000fd5b602090810291909101015296505b885161196f9088610d48565b84611bf5576000600460009054906101000a90046001600160a01b03166001600160a01b0316632ba010d760405180604001604052808c60c001516001600160a01b031681526020018f60a001516001600160401b03168152506040518263ffffffff1660e01b81526004016119e59190615bb4565b60006040518083038186803b1580156119fd57600080fd5b505afa158015611a11573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611a399190810190614c6c565b90508961028001511515816101400151151514611a6c57600860405163cd442d6560e01b81526004016104fb9190615a2b565b600480546040517f848fd7c80000000000000000000000000000000000000000000000000000000081526001600160a01b039091169163848fd7c891611ab69185918f9101615b8f565b600060405180830381600087803b158015611ad057600080fd5b505af1158015611ae4573d6000803e3d6000fd5b5050600480546040517fdcf749600000000000000000000000000000000000000000000000000000000081526001600160a01b03909116935063dcf749609250611b3091859101615b7e565b600060405180830381600087803b158015611b4a57600080fd5b505af1158015611b5e573d6000803e3d6000fd5b505050508060c0015160001415611b92578960c001516001600160401b03168c60400151611b8c9190615c9e565b60c08201525b600480546040516311c2535560e11b81526001600160a01b0390911691632384a6aa91611bc191859101615b7e565b600060405180830381600087803b158015611bdb57600080fd5b505af1158015611bef573d6000803e3d6000fd5b50505050505b8315611d5f5760a08b01516001600160401b031615611d52576000600460009054906101000a90046001600160a01b03166001600160a01b0316632ba010d760405180604001604052808c60c001516001600160a01b031681526020018f60a001516001600160401b03168152506040518263ffffffff1660e01b8152600401611c7f9190615bb4565b60006040518083038186803b158015611c9757600080fd5b505afa158015611cab573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611cd39190810190614c6c565b600480546040517fa3f029180000000000000000000000000000000000000000000000000000000081529293506001600160a01b03169163a3f0291891611d1e9185918f9101615b8f565b600060405180830381600087803b158015611d3857600080fd5b505af1158015611d4c573d6000803e3d6000fd5b50505050505b611d5f8989848a8e611fe4565b885160a08901516040517fd936063ae3cbbf2dbcd2adc837a997ab97adb09f24c566aa79d101e88928dea692611d9b92600792439291906159e1565b60405180910390a15050505050505050505050565b60606000600583604051611dc49190615931565b90815260200160405180910390209050600081600201546001600160401b03811115611e0057634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015611e5957816020015b6040805160a081018252600080825260208083018290529282018190526060820181905260808201528252600019909201910181611e1e5790505b509050816002015460001415611e70579392505050565b6000611e7b83612ee4565b90505b6001830154811015611ed7576000611e968483612eff565b91505080838381518110611eba57634e487b7160e01b600052603260045260246000fd5b602090810291909101015250611ed08382612fcc565b9050611e7e565b509392505050565b604080516060810182526000808252602082018190529181019190915260008383604051602001611f1192919061590b565b6040516020818303038152906040529050600881604051611f329190615931565b90815260408051918290036020908101832060608401835280546001600160a01b03168452600181015491840191909152600201546001600160401b03169082015291505092915050565b80600683604051611f8e9190615931565b9081526040519081900360209081019091208251815493909201516001600160401b0390811668010000000000000000026fffffffffffffffffffffffffffffffff199094169216919091179190911790555050565b6000611ff1868584613041565b9050806001600160401b03168661014001516001600160401b0316101561202e57600960405163cd442d6560e01b81526004016104fb9190615a2b565b80856060018181516120409190615cb6565b6001600160401b0316905250600254604051633206761b60e21b81526001600160a01b039091169063c819d86c9061207c908890600401615b51565b600060405180830381600087803b15801561209657600080fd5b505af11580156120aa573d6000803e3d6000fd5b505050508086610140018181516120c19190615d61565b6001600160401b031690525060006101c08701526001546040517f1aa05c5c0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690631aa05c5c9061211e908990600401615b02565b600060405180830381600087803b15801561213857600080fd5b505af115801561214c573d6000803e3d6000fd5b505050506000805b84518110156121ae5784818151811061217d57634e487b7160e01b600052603260045260246000fd5b6020026020010151608001511561219c578161219881615e69565b9250505b806121a681615e4e565b915050612154565b50806001600160401b03166001141561222957600180546040516334afd8c160e01b81526001600160a01b03909116916334afd8c1916121f6918b9160009190600401615b13565b600060405180830381600087803b15801561221057600080fd5b505af1158015612224573d6000803e3d6000fd5b505050505b61012087015161223a906001615cb6565b6001600160401b0316816001600160401b0316141561231f576101408701516001600160401b0316156122b45786602001516001600160a01b03166108fc8861014001516001600160401b03169081150290604051600060405180830381858888f193505050501580156122b2573d6000803e3d6000fd5b505b600180546040516334afd8c160e01b81526001600160a01b03909116916334afd8c1916122e8918b91600090600401615b13565b600060405180830381600087803b15801561230257600080fd5b505af1158015612316573d6000803e3d6000fd5b505050506123a2565b600154602088015188516040517f34fddaac0000000000000000000000000000000000000000000000000000000081526001600160a01b03909316926334fddaac9261236f92909160040161594b565b600060405180830381600087803b15801561238957600080fd5b505af115801561239d573d6000803e3d6000fd5b505050505b7f123b9998b2f027b02db4cb2eadef421fe9f1c21627569438588262d3a6baac6c6006438860a00151856040516123dc94939291906159ac565b60405180910390a150505050505050565b60006123f98386613135565b6124039083615d0a565b9050806001600160401b031684600001516001600160401b0316106124455780846000018181516124349190615d61565b6001600160401b031690525061244c565b5060008084525b6001600160401b038116156124e757806001600160401b03163410156124845760405162461bcd60e51b81526004016104fb90615a66565b600254604051633206761b60e21b81526001600160a01b039091169063c819d86c906124b4908790600401615b51565b600060405180830381600087803b1580156124ce57600080fd5b505af11580156124e2573d6000803e3d6000fd5b505050505b84516020808701516001600160a01b0390921660009081526007825260408082206001600160401b03909416825292909152204390555050505050565b61259983838360405160240161253c93929190615a39565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f969cdd0300000000000000000000000000000000000000000000000000000000179052613156565b505050565b6000805b8461010001516001600160401b03168110156128ba57600085610160015182815181106125df57634e487b7160e01b600052603260045260246000fd5b60209081029190910101516001546040516337bf397560e21b81529192506000916001600160a01b039091169063defce5d49061262090859060040161599b565b60006040518083038186803b15801561263857600080fd5b505afa15801561264c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526126749190810190614afb565b905060006126858260000151611db0565b6101008301516040805160a081018252600080825260208201819052918101829052606081018290526080810182905292935091829060005b85518110156127b0578c600001516001600160a01b0316868a815181106126f557634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516001600160a01b0316141561279e576001925085898151811061273457634e487b7160e01b600052603260045260246000fd5b602002602001015191506000826040015190508760e001516001600160401b0316816001600160401b0316148061276a57508443115b1561277b5760016080840181905295505b6040830180519061278b82615e69565b6001600160401b03169052506127b09050565b806127a881615e4e565b9150506126be565b50816127c7576000985050505050505050506128c0565b85516127d39086610d48565b83156128a0576127e6868c83888e611fe4565b600460009054906101000a90046001600160a01b03166001600160a01b031663a3f029188d886040518363ffffffff1660e01b8152600401612829929190615b8f565b600060405180830381600087803b15801561284357600080fd5b505af1158015612857573d6000803e3d6000fd5b5050875160a08e01516040517fb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f4594506128979350600192439290916159e1565b60405180910390a15b5050505050505080806128b290615e4e565b9150506125a2565b50600190505b9392505050565b6001600160a01b038083166000908152602085815260408083208054865160018301805491881673ffffffffffffffffffffffffffffffffffffffff1990921691909117905592860151600282018054938801516001600160401b0316600160a01b027fffffffff000000000000000000000000000000000000000000000000000000009094169190961617919091179093556060840151600384015560808401516004909301805493151560ff19909416939093179092559080156129915760019150506128c0565b50600180850180548083018255600091909152906129b0908290615c9e565b6001600160a01b038516600090815260208790526040902055600185018054859190839081106129f057634e487b7160e01b600052603260045260246000fd5b60009182526020822001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03939093169290921790915560028601805491612a3683615e4e565b919050555060009150506128c0565b60c080840151908501516000919083612a676001600160401b03841683615c9e565b10612a7757600092505050612b0a565b600082612a848387615d46565b612a8e9190615ce9565b905060008615612ad257612aab6001600160401b03851684615c9e565b871115612ace5783612abd8489615d46565b612ac79190615ce9565b9050612ad2565b5060005b806001600160401b0316826001600160401b03161015612af9576000945050505050612b0a565b612b038183615d61565b9450505050505b949350505050565b60606000612b1f83611db0565b905060005b8151811015612c2f5760025482516000916001600160a01b031690631a65374a90859085908110612b6557634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516040518263ffffffff1660e01b8152600401612b8d919061593d565b60e06040518083038186803b158015612ba557600080fd5b505afa158015612bb9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bdd9190614bfc565b90508060c00151838381518110612c0457634e487b7160e01b600052603260045260246000fd5b60209081029190910101516001600160a01b0390911690525080612c2781615e4e565b915050612b24565b5092915050565b60c08101516000904390612c53906001600160401b031682615c9e565b84604001511180612c7f5750808360c001516001600160401b03168560400151612c7d9190615c9e565b105b15612c8e576000915050610b89565b6040805180820190915260608082526020820152612cd660405180608001604052806060815260200160006001600160401b0316815260200160608152602001606081525090565b604080516080808201835260008083526060602084018190528385018281528185018381528c8301516001600160a01b039081168752948c01516001600160401b039081169092526101808c015190911690526003549451631faf930160e31b815290949192919091169063fd7c980890612d55908590600401615b40565b60006040518083038186803b158015612d6d57600080fd5b505afa158015612d81573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612da9919081019061495d565b9050612df46040518060e0016040528060006001600160401b031681526020016060815260200160608152602001606081526020016060815260200160608152602001606081525090565b60008082528551602080840191909152870151604080840191909152808701516060808501919091526080840185905287015160a0840152875160c084015260035490516305b6ef6560e51b81526001600160a01b039091169063b6ddeca090612e62908590600401615c24565b60206040518083038186803b158015612e7a57600080fd5b505afa158015612e8e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612eb29190614991565b905080610b7c57600098505050505050505050610b89565b600081431115612edc57506001919050565b506000919050565b600080612ef2836000612fcc565b90506128c0600182615d46565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101829052836001018381548110612f4c57634e487b7160e01b600052603260045260246000fd5b6000918252602080832091909101546001600160a01b03908116808452968252604092839020835160a081018552600182015483168152600282015492831693810193909352600160a01b9091046001600160401b0316928201929092526003820154606082015260049091015460ff1615156080820152939492505050565b600081612fd881615e4e565b9250505b600183015482108015613024575082600101828154811061300d57634e487b7160e01b600052603260045260246000fd5b600091825260209091200154600160a01b900460ff165b1561303b578161303381615e4e565b925050612fdc565b50919050565b60018054604084015160009283926001600160a01b03169163cc76e80d91869161306b9190615d61565b60008960a001518a608001516130819190615d0a565b8a6101a001518b61010001516130979190615d46565b6040518663ffffffff1660e01b81526004016130b7959493929190615bc2565b60606040518083038186803b1580156130cf57600080fd5b505afa1580156130e3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131079190614de8565b90508060400151816020015182600001516131229190615cb6565b61312c9190615cb6565b95945050505050565b60008060029050600061314c858560600151613177565b61312c9083615d0a565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b6000620fa00082846060015161318d9190615d0a565b6128c09190615ce9565b604080516101e08101909152600060608083018281526080840183905260a0840183905260c0840183905260e08401839052610100840183905261012084018390526101408401839052610160840183905261018084018390526101a08401929092526101c083015281908152602001606081526020016132596040518060c0016040528060006001600160401b0316815260200160006001600160401b03168152602001606081526020016060815260200160608152602001606081525090565b905290565b6040518060e001604052806060815260200160608152602001606081526020016060815260200160608152602001613357604080516102c08101825260608082526000602083018190529282018190528082018390526080820183905260a0820183905260c0820183905260e0820183905261010082018390526101208201839052610140820183905261016082015261018081018290526101a081018290526101c08101829052906101e0820190815260006020820181905260606040830181905280830152608090910190815260006020808301829052604080516060810182528381529182018390528181019290925291015290565b8152600060209091015290565b50805460008255906000526020600020908101906133829190613385565b50565b5b808211156133bc5780547fffffffffffffffffffffff000000000000000000000000000000000000000000168155600101613386565b5090565b60006133d36133ce84615c51565b615c35565b905080838252602082019050828560208602820111156133f257600080fd5b60005b8581101561341e5781613408888261375b565b84525060209283019291909101906001016133f5565b5050509392505050565b60006134366133ce84615c51565b9050808382526020820190508285602086028201111561345557600080fd5b60005b8581101561341e578161346b8882613766565b8452506020928301929190910190600101613458565b600061348f6133ce84615c51565b905080838252602082019050828560208602820111156134ae57600080fd5b60005b8581101561341e5781356001600160401b038111156134cf57600080fd5b8086016134dc898261388f565b8552505060209283019291909101906001016134b1565b60006135016133ce84615c51565b9050808382526020820190508285602086028201111561352057600080fd5b60005b8581101561341e5781516001600160401b0381111561354157600080fd5b80860161354e89826138b0565b855250506020928301929190910190600101613523565b60006135736133ce84615c51565b8381529050602081018260408502810186101561358f57600080fd5b60005b8581101561341e57816135a58882613908565b84525060209092019160409190910190600101613592565b60006135cb6133ce84615c51565b905080838252602082019050828560208602820111156135ea57600080fd5b60005b8581101561341e5781516001600160401b0381111561360b57600080fd5b8086016136188982613ed5565b8552505060209283019291909101906001016135ed565b600061363d6133ce84615c51565b9050808382526020820190508285602086028201111561365c57600080fd5b60005b8581101561341e5781516001600160401b0381111561367d57600080fd5b80860161368a8982613f47565b85525050602092830192919091019060010161365f565b60006136af6133ce84615c51565b8381529050602081018260a0850281018610156136cb57600080fd5b60005b8581101561341e57816136e18882614320565b84525060209092019160a091909101906001016136ce565b60006137076133ce84615c74565b90508281526020810184848401111561371f57600080fd5b611ed7848285615dea565b60006137386133ce84615c74565b90508281526020810184848401111561375057600080fd5b611ed7848285615df6565b8035610b8981615f28565b8051610b8981615f28565b600082601f83011261378257600080fd5b8135612b0a8482602086016133c0565b600082601f8301126137a357600080fd5b8151612b0a848260208601613428565b600082601f8301126137c457600080fd5b8135612b0a848260208601613481565b600082601f8301126137e557600080fd5b8151612b0a8482602086016134f3565b600082601f83011261380657600080fd5b8151612b0a848260208601613565565b600082601f83011261382757600080fd5b8151612b0a8482602086016135bd565b600082601f83011261384857600080fd5b8151612b0a84826020860161362f565b600082601f83011261386957600080fd5b8135612b0a8482602086016136a1565b8035610b8981615f3c565b8051610b8981615f3c565b600082601f8301126138a057600080fd5b8135612b0a8482602086016136f9565b600082601f8301126138c157600080fd5b8151612b0a84826020860161372a565b8035610b8981615f44565b8035610b8981615f4d565b8051610b8981615f4d565b8035610b8981615f5a565b8051610b8981615f5a565b60006040828403121561391a57600080fd5b6139246040615c35565b9050600061393284846148b5565b8252506020613943848483016148b5565b60208301525092915050565b6000610300828403121561396257600080fd5b61396d6102c0615c35565b905081356001600160401b0381111561398557600080fd5b6139918482850161388f565b82525060206139a28484830161375b565b60208301525060408201356001600160401b038111156139c157600080fd5b6139cd8482850161388f565b60408301525060606139e1848285016148c0565b60608301525060806139f5848285016148c0565b60808301525060a0613a09848285016148c0565b60a08301525060c0613a1d848285016148c0565b60c08301525060e0613a31848285016148c0565b60e083015250610100613a468482850161489f565b61010083015250610120613a5c848285016148c0565b61012083015250610140613a72848285016148c0565b610140830152506101608201356001600160401b03811115613a9357600080fd5b613a9f8482850161388f565b61016083015250610180613ab5848285016148c0565b610180830152506101a0613acb8482850161489f565b6101a0830152506101c0613ae184828501613879565b6101c0830152506101e0613af7848285016138f2565b6101e083015250610200613b0d848285016148c0565b610200830152506102208201356001600160401b03811115613b2e57600080fd5b613b3a84828501613771565b610220830152506102408201356001600160401b03811115613b5b57600080fd5b613b6784828501613771565b61024083015250610260613b7d848285016138dc565b61026083015250610280613b9384828501613879565b610280830152506102a0613ba98482850161420c565b6102a08301525092915050565b60006103008284031215613bc957600080fd5b613bd46102c0615c35565b82519091506001600160401b03811115613bed57600080fd5b613bf9848285016138b0565b8252506020613c0a84848301613766565b60208301525060408201516001600160401b03811115613c2957600080fd5b613c35848285016138b0565b6040830152506060613c49848285016148cb565b6060830152506080613c5d848285016148cb565b60808301525060a0613c71848285016148cb565b60a08301525060c0613c85848285016148cb565b60c08301525060e0613c99848285016148cb565b60e083015250610100613cae848285016148aa565b61010083015250610120613cc4848285016148cb565b61012083015250610140613cda848285016148cb565b610140830152506101608201516001600160401b03811115613cfb57600080fd5b613d07848285016138b0565b61016083015250610180613d1d848285016148cb565b610180830152506101a0613d33848285016148aa565b6101a0830152506101c0613d4984828501613884565b6101c0830152506101e0613d5f848285016138fd565b6101e083015250610200613d75848285016148cb565b610200830152506102208201516001600160401b03811115613d9657600080fd5b613da284828501613792565b610220830152506102408201516001600160401b03811115613dc357600080fd5b613dcf84828501613792565b61024083015250610260613de5848285016138e7565b61026083015250610280613dfb84828501613884565b610280830152506102a0613ba98482850161425b565b600060c08284031215613e2357600080fd5b613e2d60c0615c35565b905081356001600160401b03811115613e4557600080fd5b613e518482850161388f565b82525060208201356001600160401b03811115613e6d57600080fd5b613e798482850161388f565b6020830152506040613e8d8482850161489f565b6040830152506060613ea18482850161375b565b6060830152506080613eb5848285016148c0565b60808301525060a0613ec9848285016148c0565b60a08301525092915050565b600060608284031215613ee757600080fd5b613ef16060615c35565b90506000613eff84846148cb565b8252506020613f10848483016148cb565b60208301525060408201516001600160401b03811115613f2f57600080fd5b613f3b848285016138b0565b60408301525092915050565b600060408284031215613f5957600080fd5b613f636040615c35565b90506000613f7184846148cb565b82525060208201516001600160401b03811115613f8d57600080fd5b61394384828501613816565b600060e08284031215613fab57600080fd5b613fb560e0615c35565b90506000613fc384846148c0565b8252506020613fd4848483016148c0565b6020830152506040613fe8848285016148c0565b6040830152506060613ffc848285016148c0565b6060830152506080614010848285016148c0565b60808301525060a06140248482850161375b565b60a08301525060c06140388482850161375b565b60c08301525092915050565b600060e0828403121561405657600080fd5b61406060e0615c35565b9050600061406e84846148cb565b825250602061407f848483016148cb565b6020830152506040614093848285016148cb565b60408301525060606140a7848285016148cb565b60608301525060806140bb848285016148cb565b60808301525060a06140cf84828501613766565b60a08301525060c061403884828501613766565b600060e082840312156140f557600080fd5b6140ff60e0615c35565b82519091506001600160401b0381111561411857600080fd5b614124848285016137d4565b82525060208201516001600160401b0381111561414057600080fd5b61414c848285016137d4565b60208301525060408201516001600160401b0381111561416b57600080fd5b614177848285016137f5565b60408301525060608201516001600160401b0381111561419657600080fd5b6141a284828501613837565b60608301525060808201516001600160401b038111156141c157600080fd5b6141cd848285016138b0565b60808301525060a08201516001600160401b038111156141ec57600080fd5b6141f884828501613bb6565b60a08301525060c061403884828501613884565b60006060828403121561421e57600080fd5b6142286060615c35565b9050600061423684846148c0565b8252506020614247848483016148c0565b6020830152506040613f3b848285016148c0565b60006060828403121561426d57600080fd5b6142776060615c35565b9050600061428584846148cb565b8252506020614296848483016148cb565b6020830152506040613f3b848285016148cb565b6000606082840312156142bc57600080fd5b6142c66060615c35565b905060006142d4848461375b565b82525060206142478484830161489f565b6000604082840312156142f757600080fd5b6143016040615c35565b9050600061430f84846148c0565b8252506020613943848483016148c0565b600060a0828403121561433257600080fd5b61433c60a0615c35565b9050600061434a848461375b565b825250602061435b8484830161375b565b602083015250604061436f848285016148c0565b60408301525060606143838482850161489f565b606083015250608061439784828501613879565b60808301525092915050565b600061018082840312156143b657600080fd5b6143c1610180615c35565b905060006143cf848461375b565b82525060206143e0848483016148c0565b60208301525060406143f4848285016148c0565b6040830152506060614408848285016148c0565b606083015250608061441c848285016138dc565b60808301525060a06144308482850161489f565b60a08301525060c06144448482850161489f565b60c08301525060e0614458848285016148c0565b60e08301525061010061446d848285016148c0565b61010083015250610120614483848285016148c0565b6101208301525061014061449984828501613879565b610140830152506101608201356001600160401b038111156144ba57600080fd5b6144c6848285016137b3565b6101608301525092915050565b600061018082840312156144e657600080fd5b6144f1610180615c35565b905060006144ff8484613766565b8252506020614510848483016148cb565b6020830152506040614524848285016148cb565b6040830152506060614538848285016148cb565b606083015250608061454c848285016138e7565b60808301525060a0614560848285016148aa565b60a08301525060c0614574848285016148aa565b60c08301525060e0614588848285016148cb565b60e08301525061010061459d848285016148cb565b610100830152506101206145b3848285016148cb565b610120830152506101406145c984828501613884565b610140830152506101608201516001600160401b038111156145ea57600080fd5b6144c6848285016137d4565b60006080828403121561460857600080fd5b6146126080615c35565b90506000614620848461375b565b8252506020614631848483016148c0565b6020830152506040614645848285016148c0565b60408301525060608201356001600160401b0381111561466457600080fd5b6146708482850161388f565b60608301525092915050565b60006040828403121561468e57600080fd5b6146986040615c35565b9050600061430f848461375b565b600061016082840312156146b957600080fd5b6146c4610160615c35565b905060006146d284846148c0565b82525060206146e3848483016148c0565b60208301525060406146f7848285016148c0565b604083015250606061470b848285016148c0565b606083015250608061471f848285016148c0565b60808301525060a0614733848285016148c0565b60a08301525060c0614747848285016148c0565b60c08301525060e061475b848285016148c0565b60e083015250610100614770848285016148c0565b61010083015250610120614786848285016148c0565b6101208301525061014061479c848285016148c0565b6101408301525092915050565b600061016082840312156147bc57600080fd5b6147c7610160615c35565b905060006147d584846148cb565b82525060206147e6848483016148cb565b60208301525060406147fa848285016148cb565b604083015250606061480e848285016148cb565b6060830152506080614822848285016148cb565b60808301525060a0614836848285016148cb565b60a08301525060c061484a848285016148cb565b60c08301525060e061485e848285016148cb565b60e083015250610100614873848285016148cb565b61010083015250610120614889848285016148cb565b6101208301525061014061479c848285016148cb565b8035610b8981615f67565b8051610b8981615f67565b8051610b8981615f6d565b8035610b8981615f79565b8051610b8981615f79565b600080604083850312156148e957600080fd5b60006148f5858561375b565b92505060206149068582860161489f565b9150509250929050565b60008060006060848603121561492557600080fd5b6000614931868661375b565b9350506020614942868287016148c0565b92505060406149538682870161489f565b9150509250925092565b60006020828403121561496f57600080fd5b81516001600160401b0381111561498557600080fd5b612b0a848285016137f5565b6000602082840312156149a357600080fd5b6000612b0a8484613884565b6000602082840312156149c157600080fd5b81356001600160401b038111156149d757600080fd5b612b0a8482850161388f565b600080604083850312156149f657600080fd5b82356001600160401b03811115614a0c57600080fd5b614a188582860161388f565b92505060208301356001600160401b03811115614a3457600080fd5b61490685828601613858565b60008060608385031215614a5357600080fd5b82356001600160401b03811115614a6957600080fd5b614a758582860161388f565b9250506020614906858286016142e5565b600080600080600060a08688031215614a9e57600080fd5b6000614aaa88886138d1565b9550506020614abb888289016138d1565b9450506040614acc888289016138d1565b9350506060614add888289016138d1565b9250506080614aee888289016138d1565b9150509295509295909350565b600060208284031215614b0d57600080fd5b81516001600160401b03811115614b2357600080fd5b612b0a84828501613bb6565b60008060008060006103208688031215614b4857600080fd5b85356001600160401b03811115614b5e57600080fd5b614b6a8882890161394f565b9550506020614b7b88828901613f99565b945050610100614b8d88828901614320565b9350506101a08601356001600160401b03811115614baa57600080fd5b614bb688828901613858565b9250506101c0614aee888289016146a6565b600060208284031215614bda57600080fd5b81356001600160401b03811115614bf057600080fd5b612b0a84828501613e11565b600060e08284031215614c0e57600080fd5b6000612b0a8484614044565b600060208284031215614c2c57600080fd5b81516001600160401b03811115614c4257600080fd5b612b0a848285016140e3565b600060608284031215614c6057600080fd5b6000612b0a84846142aa565b600060208284031215614c7e57600080fd5b81516001600160401b03811115614c9457600080fd5b612b0a848285016144d3565b6000806000806102808587031215614cb757600080fd5b84356001600160401b03811115614ccd57600080fd5b614cd9878288016143a3565b9450506020614cea87828801613f99565b935050610100614cfc878288016146a6565b925050610260614d0e878288016148c0565b91505092959194509250565b600060208284031215614d2c57600080fd5b81356001600160401b03811115614d4257600080fd5b612b0a848285016145f6565b60008060408385031215614d6157600080fd5b82356001600160401b03811115614d7757600080fd5b614d83858286016145f6565b92505060208301356001600160401b03811115614d9f57600080fd5b614906858286016143a3565b600060408284031215614dbd57600080fd5b6000612b0a848461467c565b60006101608284031215614ddc57600080fd5b6000612b0a84846147a9565b600060608284031215614dfa57600080fd5b6000612b0a848461425b565b6000614e128383614e7a565b505060200190565b60006128c083836150bd565b6000614e328383615135565b505060400190565b60006128c0838361538f565b60006128c083836153ce565b6000614e5e83836154b4565b505060600190565b6000614e728383615524565b505060a00190565b614e8381615d7e565b82525050565b614e83614e9582615d7e565b615e8e565b6000614ea4825190565b80845260209384019383018060005b83811015614ed8578151614ec78882614e06565b975060208301925050600101614eb3565b509495945050505050565b6000614eed825190565b80845260208401935083602082028501614f078560200190565b8060005b85811015614f3c5784840389528151614f248582614e1a565b94506020830160209a909a0199925050600101614f0b565b5091979650505050505050565b6000614f53825190565b80845260209384019383018060005b83811015614ed8578151614f768882614e26565b975060208301925050600101614f62565b6000614f91825190565b80845260208401935083602082028501614fab8560200190565b8060005b85811015614f3c5784840389528151614fc88582614e3a565b94506020830160209a909a0199925050600101614faf565b6000614fea825190565b808452602084019350836020820285016150048560200190565b8060005b85811015614f3c57848403895281516150218582614e46565b94506020830160209a909a0199925050600101615008565b6000615043825190565b80845260209384019383018060005b83811015614ed85781516150668882614e52565b975060208301925050600101615052565b6000615081825190565b80845260209384019383018060005b83811015614ed85781516150a48882614e66565b975060208301925050600101615090565b801515614e83565b60006150c7825190565b8084526020840193506150de818560208601615df6565b601f01601f19169290920192915050565b60006150f9825190565b615107818560208601615df6565b9290920192915050565b614e8381615db8565b614e8381615dc3565b614e8381615dce565b614e8381615dd9565b8051604083019061514684826158f0565b506020820151610dd060208501826158f0565b80516103008084526000919084019061517282826150bd565b91505060208301516151876020860182614e7a565b506040830151848203604086015261519f82826150bd565b91505060608301516151b460608601826158fc565b5060808301516151c760808601826158fc565b5060a08301516151da60a08601826158fc565b5060c08301516151ed60c08601826158fc565b5060e083015161520060e08601826158fc565b506101008301516152156101008601826158ea565b5061012083015161522a6101208601826158fc565b5061014083015161523f6101408601826158fc565b5061016083015184820361016086015261525982826150bd565b9150506101808301516152706101808601826158fc565b506101a08301516152856101a08601826158ea565b506101c083015161529a6101c08601826150b5565b506101e08301516152af6101e0860182615123565b506102008301516152c46102008601826158fc565b506102208301518482036102208601526152de8282614e9a565b9150506102408301518482036102408601526152fa8282614e9a565b91505061026083015161531161026086018261511a565b506102808301516153266102808601826150b5565b506102a0830151611ed76102a086018261547d565b8051600090608084019061534f8582614e7a565b506020830151848203602086015261536782826150bd565b915050604083015161537c60408601826158fc565b506060830151611ed760608601826158fc565b805160009060608401906153a385826158fc565b5060208301516153b660208601826158fc565b506040830151848203604086015261312c82826150bd565b805160009060408401906153e285826158fc565b506020830151848203602086015261312c8282614f87565b805160e083019061540b84826158fc565b50602082015161541e60208501826158fc565b50604082015161543160408501826158fc565b50606082015161544460608501826158fc565b50608082015161545760808501826158fc565b5060a082015161546a60a0850182614e7a565b5060c0820151610dd060c0850182614e7a565b8051606083019061548e84826158fc565b5060208201516154a160208501826158fc565b506040820151610dd060408501826158fc565b805160608301906154c58482614e7a565b5060208201516154a160208501826158ea565b80516060808452600091908401906154f08282615581565b9150506020830151848203602086015261550a8282614f49565b9150506040830151848203604086015261312c8282615674565b805160a08301906155358482614e7a565b5060208201516155486020850182614e7a565b50604082015161555b60408501826158fc565b50606082015161556e60608501826158ea565b506080820151610dd060808501826150b5565b80516000906101808401906155968582614e7a565b5060208301516155a960208601826158fc565b5060408301516155bc60408601826158fc565b5060608301516155cf60608601826158fc565b5060808301516155e2608086018261511a565b5060a08301516155f560a08601826158ea565b5060c083015161560860c08601826158ea565b5060e083015161561b60e08601826158fc565b506101008301516156306101008601826158fc565b506101208301516156456101208601826158fc565b5061014083015161565a6101408601826150b5565b5061016083015184820361016086015261312c8282614ee3565b805160009060c084019061568885826158fc565b50602083015161569b60208601826158fc565b50604083015184820360408601526156b382826150bd565b915050606083015184820360608601526156cd82826150bd565b915050608083015184820360808601526156e78282614fe0565b91505060a083015184820360a086015261312c82826150bd565b805160408301906157128482614e7a565b506020820151610dd060208501826158fc565b805161016083019061573784826158fc565b50602082015161574a60208501826158fc565b50604082015161575d60408501826158fc565b50606082015161577060608501826158fc565b50608082015161578360808501826158fc565b5060a082015161579660a08501826158fc565b5060c08201516157a960c08501826158fc565b5060e08201516157bc60e08501826158fc565b506101008201516157d16101008501826158fc565b506101208201516157e66101208501826158fc565b50610140820151610dd06101408501826158fc565b805160009060a084019061580f858261547d565b506020830151848203606086015261582782826150bd565b9150506040830151611ed760808601826158fc565b805160009060e084019061585085826158fc565b506020830151848203602086015261586882826150bd565b915050604083015184820360408601526158828282614ee3565b9150506060830151848203606086015261589c8282614ee3565b915050608083015184820360808601526158b68282614f49565b91505060a083015184820360a08601526158d08282614fe0565b91505060c083015184820360c086015261312c82826150bd565b80614e83565b63ffffffff8116614e83565b6001600160401b038116614e83565b60006159178285614e89565b60148201915061592782846158ea565b5060200192915050565b60006128c082846150ef565b60208101610b898284614e7a565b604081016159598285614e7a565b8181036020830152612b0a81846150bd565b602080825281016128c08184615039565b602080825281016128c08184615077565b60208101610b8982846150b5565b602080825281016128c081846150bd565b608081016159ba8287615111565b6159c760208301866158ea565b6159d46040830185614e7a565b61312c60608301846158fc565b608081016159ef8287615111565b6159fc60208301866158ea565b8181036040830152615a0e81856150bd565b905061312c6060830184614e7a565b60208101610b89828461511a565b60208101610b89828461512c565b60608082528101615a4a81866150bd565b9050615a5960208301856158ea565b612b0a60408301846158ea565b60208082528101610b8981601681527f50756e697368466f72536563746f72206661696c656400000000000000000000602082015260400190565b60208082528101610b8981602e81527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160208201527f647920696e697469616c697a6564000000000000000000000000000000000000604082015260600190565b602080825281016128c08184615159565b60608082528101615b248186615159565b9050615b3360208301856150b5565b612b0a60408301846150b5565b602080825281016128c0818461533b565b60e08101610b8982846153fa565b60608101610b8982846154b4565b602080825281016128c081846154d8565b602080825281016128c08184615581565b60408082528101615ba08185615581565b90508181036020830152612b0a8184615159565b60408101610b898284615701565b6101e08101615bd18288615725565b615bdf6101608301876158fc565b615bed61018083018661512c565b615bfb6101a08301856158fc565b615c096101c08301846158fc565b9695505050505050565b602080825281016128c081846157fb565b602080825281016128c0818461583c565b6000615c4060405190565b9050615c4c8282615e22565b919050565b60006001600160401b03821115615c6a57615c6a615ee2565b5060209081020190565b60006001600160401b03821115615c8d57615c8d615ee2565b601f19601f83011660200192915050565b60008219821115615cb157615cb1615ea0565b500190565b60006001600160401b03821691506001600160401b0383169250826001600160401b0303821115615cb157615cb1615ea0565b6001600160401b039182169116600082615d0557615d05615eb6565b500490565b60006001600160401b03821691506001600160401b0383169250816001600160401b030483118215151615615d4157615d41615ea0565b500290565b6000825b925082821015615d5c57615d5c615ea0565b500390565b60006001600160401b03821691506001600160401b038316615d4a565b60006001600160a01b038216610b89565b6000610b8982615d7e565b80615c4c81615ef8565b80615c4c81615f08565b80615c4c81615f18565b6000610b8982615d9a565b6000610b8982615da4565b6000610b8982615dae565b60006001600160401b038216610b89565b82818337506000910152565b60005b83811015615e11578181015183820152602001615df9565b83811115610dd05750506000910152565b601f19601f83011681018181106001600160401b0382111715615e4757615e47615ee2565b6040525050565b6000600019821415615e6257615e62615ea0565b5060010190565b60006001600160401b03821691506001600160401b03821415615e6257615e62615ea0565b6000610b89826000610b898260601b90565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600a811061338257613382615ecc565b6003811061338257613382615ecc565b6002811061338257613382615ecc565b615f3181615d7e565b811461338257600080fd5b801515615f31565b615f3181615d8f565b6003811061338257600080fd5b6002811061338257600080fd5b80615f31565b63ffffffff8116615f31565b6001600160401b038116615f3156fea264697066735822122002075e72bd2b4ba6f3ede3bb7203aaba878285d42c864936707a61277508a96164736f6c63430008040033",
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
