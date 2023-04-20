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

// FileProveParams is an auto generated low-level Go binding around an user-defined struct.
type FileProveParams struct {
	FileHash    []byte
	ProveData   ProveData
	BlockHeight *big.Int
	NodeWallet  common.Address
	Profit      uint64
	SectorID    uint64
}

// MerkleNode is an auto generated low-level Go binding around an user-defined struct.
type MerkleNode struct {
	Layer uint64
	Index uint64
	Hash  []byte
}

// MerklePath is an auto generated low-level Go binding around an user-defined struct.
type MerklePath struct {
	PathLen uint64
	Path    []MerkleNode
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

// NodeInfo is an auto generated low-level Go binding around an user-defined struct.
type NodeInfo struct {
	Pledge      uint64
	Profit      uint64
	Volume      uint64
	RestVol     uint64
	ServiceTime uint64
	WalletAddr  common.Address
	NodeAddr    []byte
}

// ProveConfig is an auto generated low-level Go binding around an user-defined struct.
type ProveConfig struct {
	SECTORPROVEBLOCKNUM uint64
}

// ProveData is an auto generated low-level Go binding around an user-defined struct.
type ProveData struct {
	Proofs     []byte
	BlockNum   uint64
	Tags       [][]byte
	MerklePath []MerklePath
}

// ProveDetail is an auto generated low-level Go binding around an user-defined struct.
type ProveDetail struct {
	NodeAddr    []byte
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
	ProveData       ProveData
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"port\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"}],\"name\":\"DNSNodeRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DNSNodeUnReg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"DnsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"FsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structNameInfo\",\"name\":\"newer\",\"type\":\"tuple\"}],\"name\":\"NotifyNameInfoAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"PDPVerifyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nodeAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"CheckNodeSectorProvedInTime\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteProveDetails\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"BlockNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"MerklePath_\",\"type\":\"tuple[]\"}],\"internalType\":\"structProveData\",\"name\":\"ProveData_\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"NodeWallet\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"}],\"internalType\":\"structFileProveParams\",\"name\":\"fileProve\",\"type\":\"tuple\"}],\"name\":\"FileProve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"GetProveDetailList\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"NodeAddr\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ChallengeHeight\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"BlockNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"MerklePath_\",\"type\":\"tuple[]\"}],\"internalType\":\"structProveData\",\"name\":\"ProveData_\",\"type\":\"tuple\"}],\"internalType\":\"structSectorProveParams\",\"name\":\"sectorProve\",\"type\":\"tuple\"}],\"name\":\"SectorProve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"SetLastPunishmentHeightForNode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"NodeAddr\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"details\",\"type\":\"tuple[]\"}],\"name\":\"UpdateProveDetailList\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveDetailNum\",\"type\":\"uint64\"}],\"internalType\":\"structProveDetailMeta\",\"name\":\"details\",\"type\":\"tuple\"}],\"name\":\"UpdateProveDetailMeta\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0xd11e680a\",\"type\":\"bytes32\"}],\"name\":\"c_0xd11e680a\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractIFile\",\"name\":\"_fs\",\"type\":\"address\"},{\"internalType\":\"contractINode\",\"name\":\"_node\",\"type\":\"address\"},{\"internalType\":\"contractIPDP\",\"name\":\"_pdp\",\"type\":\"address\"},{\"internalType\":\"contractISector\",\"name\":\"_sector\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"SECTOR_PROVE_BLOCK_NUM\",\"type\":\"uint64\"}],\"internalType\":\"structProveConfig\",\"name\":\"proveConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractProveExtra\",\"name\":\"_proveExtra\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"NodeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"}],\"name\":\"profitSplitForSector\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"NodeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"times\",\"type\":\"uint64\"}],\"name\":\"punishForSector\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5061d52c80620000226000396000f3fe6080604052600436106100c75760003560e01c80639696161411610074578063bb4e4e421161004e578063bb4e4e421461021e578063be5ac5ea1461023a578063d2561e0a14610256576100c7565b80639696161414610188578063977fdfe2146101b1578063a0004412146101ee576100c7565b806352e06146116100a557806352e0614614610134578063648d225d146101505780638d4c13a91461016c576100c7565b80630e3459fd146100cc578063181197f7146100fc5780633ec0f5df14610118575b600080fd5b6100e660048036038101906100e1919061afaa565b61027f565b6040516100f3919061c155565b60405180910390f35b6101166004803603810190610111919061acde565b6112e2565b005b610132600480360381019061012d919061abbb565b6113f9565b005b61014e6004803603810190610149919061ac9d565b6114eb565b005b61016a6004803603810190610165919061b100565b6115ff565b005b6101866004803603810190610181919061aebe565b61243e565b005b34801561019457600080fd5b506101af60048036038101906101aa919061ac74565b616689565b005b3480156101bd57600080fd5b506101d860048036038101906101d3919061ac9d565b61668c565b6040516101e5919061c133565b60405180910390f35b6102086004803603810190610203919061b02a565b6167c9565b604051610215919061c400565b60405180910390f35b6102386004803603810190610233919061ad4a565b616e81565b005b610254600480360381019061024f919061b0bf565b616f98565b005b34801561026257600080fd5b5061027d6004803603810190610278919061ad9e565b61864f565b005b60006102ad7f3deac3417ddec3f47b742d262a30e32a4f426dfcd663b84fffaff1863318abd360001b616689565b6102d97fb24a3efcbd56de310049010d7866e481c664d28b2ffa3440f942e22c7c54459660001b616689565b6103057fbc7d1080be83595935022a839600d9239e76d68ec8bd533e17c41767445cee8360001b616689565b60005b8461016001515181101561127d576103427f8c490fabd1b1b238a23c11fe893266d2094712314af8d1f6652565a35229a91760001b616689565b61036e7f05f5f75fe106cbebe775c0b48909e7e30841a7cb6dff8f24a14010036876860760001b616689565b600085610160015182815181106103ae577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015190506103e47f8abb5b78acd72eaaa7087cae9f73ac605145632decf1c9ece5977529f210577160001b616689565b6104107f599f08cabc01a93e8ec4b1c5bddaca7042933ca8f3570ae8a4322c11ca84955d60001b616689565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663defce5d4836040518263ffffffff1660e01b815260040161046d919061c170565b60006040518083038186803b15801561048557600080fd5b505afa158015610499573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906104c2919061ae7d565b90506104f07f45acdf87ea9b227cd5f36a23361546df830375ed8701ecb4d6fb186775ffaa6960001b616689565b61051c7f8d0819833dfcd56e5a413654d010c009323d423c169f19a40874accde90e869b60001b616689565b600061052b826000015161668c565b90506105597ff92c71d9ea1794010d8c9076e71131b70550d031264c42479e6da62f4f1e9cb260001b616689565b6105857f2c447c33cad9228cee164a1c78aad431b1cbbbab549e7ba7f00fe79770ebbd6b60001b616689565b60006105b37f8a7d4c44ce1105275da25e94ec57319e75139ce398fb7e3d301e5270976f410660001b616689565b6105df7fac120f83581d3f571ee73661890add761fd3c2beb198f2a5f23ec0a99c1b39d160001b616689565b600083610100015190506106157f91cd826c4379367bd943208a01cb78e33fd4b626382908ae39c5c01973f8393f60001b616689565b6106417ff513cfafad589a3d1284730a81e616edf4449a89f9982c4cee1e1ac8e964658960001b616689565b600061066f7fa699f42204145ed94f1dd15373cd92cbb32fe2fa2b454dfa2e96a174b5b761b660001b616689565b61069b7fcb3b4e4e57b2b7b39bb0af199b1ed7c60b941e775af39bbb7890c6fb5cd7d6ca60001b616689565b6106a3619596565b6106cf7fdd750d587bb877ff6f0cf5c550bcafc6838d7b5950235068aa7d4803cfc3dfd960001b616689565b6106fb7fd5904041d4df235051fd87441fe6f5ca256ecff8b63dc0cd8a5d6fdb220b152860001b616689565b60005b8551811015610bcb576107337f4f40ffdfab6583f6648ce51d5fdd530fa811f79dfed6b067f95183cbb480209a60001b616689565b61075f7f2f18a8a73fd8d0bb225ff1f8b178a083c902482a230abc4ac583a921eb1cfb4f60001b616689565b8c6000015173ffffffffffffffffffffffffffffffffffffffff16868a815181106107b3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff161415610b8c576108077f576a4ce9c4e20c4979fd603c96d54188c49c9b5185f0bd3bb8f6b0d9bb7b240160001b616689565b6108337f27387611edafbd7c786e39076b2a9cf6d1fb582c805cdcd8cc2707abc5b34c3960001b616689565b61085f7f197a08499db9b2dc32a06f5bbbaed662c934a3cfb50210e3b883dc7ca6f0adbc60001b616689565b6001925061088f7f87d85117315d03e1e59002bcc1882801e7a413b37de2d6b9a34bacf140f98be960001b616689565b6108bb7f24302a883d9908178ec90bb71d2a1e7d223ba01a29d6ebb0803b1302eeb449dd60001b616689565b8589815181106108f4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151915061092a7fb02d233e7c581b8e201ba26a414b46b3bae08ba2bbfa41fa254c219582d9513260001b616689565b6109567f53b441dfaa935a05f7b5a9d7dea4dd50adddd657c3372f526df169964e6b632e60001b616689565b60008260400151905061098b7fe05ddd66bc59c8a672453a2b5880abe9f5c86c3ee4128711c7df18554e2dfb4160001b616689565b6109b77f5bc9af7d0d3d4a806d3e010778a95347a5d4a7f2c2d891539dc22df51650749f60001b616689565b8760e0015167ffffffffffffffff168167ffffffffffffffff1614806109dc57508443115b15610ad757610a0d7f10f7e9f9ee133406fe194eba56f71c703eca55bdbff0efb6c78d313e986c312e60001b616689565b610a397fbb3acf41521eb4c3aa16aa1daae17e0f7dc8847dac7bcbfde4a6277d6866f8ac60001b616689565b610a657fee45cb0609f16fe82df07b1eede2758d1cfa681c688f3fb3d5736c04f268335c60001b616689565b6001836080019015159081151581525050610aa27ff0cd79669726dbb1c8657cdb5685b94bd24533c627b1a9e4870045957c547b7360001b616689565b610ace7f3373e859977780bb4d5465dff26993141075ed585f6f1301a79887286e7bacae60001b616689565b60019550610b04565b610b037fea909091a3a07bf1b660662d6aa5d7e4c6d33ef2ed51642a680130e04933cf7c60001b616689565b5b610b307f9711bb4f03cda90e0024278e8416e83a20d636ee2ed230e020408fa96e5da5bc60001b616689565b826040018051809190610b429061cf45565b67ffffffffffffffff1667ffffffffffffffff1681525050610b867fdd975d28baf866003347bfaeaaf1069370e3b31fbd12c5325a20b63aaaa5a1e360001b616689565b50610bcb565b610bb87f22b4882f1d0f96a767b22c393d5446e5d82e1f1c91faf2dace01eb6b4a03c78c60001b616689565b8080610bc39061cefc565b9150506106fe565b50610bf87f8174475ad5054a84b7acfe20ae4cb0630d8fde8cf08495eeb381b474271d396b60001b616689565b610c247fbc15ca9f5863554f12e658b356518aca58729599bf8a881c91515314c70e0b1960001b616689565b81610cbe57610c557fb533844f559f73222cc37c8e2606fce50a2a3478271d23b7256cb7683157b91c60001b616689565b610c817f7c454e5af6135c5284f0d10bbc3aa8bcec3709f0fbfbd65d794eac0c2f04601360001b616689565b610cad7f95b87252f8eee41b97c15bbbcb814baf3faca79efecd3bfa572eb3d908f0b52260001b616689565b6000985050505050505050506112db565b610cea7f057e52dd26a59fd40b5b9e97a52280a8e48bde1c6af403d0f853334cedec8ddc60001b616689565b610d167ff1114f4434ab1c9da6a39dabf80e20e64ee66ac1866b06e684dacd156c31538a60001b616689565b610d427f3c6c09712e1a7fd3b1ebc89bcd32ffd4b159ce43e9dcdcf5620faaae5145326c60001b616689565b60008c9050610d737f0e77efe37497432ac3110cce10ad9e84d675afc8305f178360a776ec272e60cb60001b616689565b610d9f7f61060be35e695ff03e2fa501522405854ef0de32a6a456eb6fb9e61fa0d47fd560001b616689565b6000879050610dd07f0bf47693d93b706ca0040d225f445441ee899ae37b673b6ea4952758b368a7dd60001b616689565b610dfc7fde0ac2fb9abfebbbfed427e50ee0ee201a95c1fb4f417f12e82a7f83e99fa99060001b616689565b60008d9050610e2d7ffb5a9547d7fd2ba7a87f9065bb763aec6977d5390ec338e20e444d3cab5996d360001b616689565b610e597f07a2da1119b54bee964acfa651c6f327744b0c813d1aacae64d39cce010d333760001b616689565b60008d9050610e8a7f347e16cf260d04e4b3c1ca20339d2c3e37df01643c796c3a4528bbb99f28200960001b616689565b610eb67ff159827a714aa99e03d5c5200a6e991600d82da68d5188c4831e54d146a311df60001b616689565b610ec483600001518a6112e2565b610ef07f6d6cf2fb6fb8cde963e2ad0e4394b629007cb83843a861366d9c3a98694e52d560001b616689565b610f1c7fb0decd63afe0985544cbea42a5ad016c21bcae6f1ee2cc92a7bc534d5921e57a60001b616689565b871561123257610f4e7fd92b7f8dad6db03390b5afa82d0bd5b606075a88bdfe4eab3c289f3fd00328fe60001b616689565b610f7a7f6eb262c2e13f2b4de1a16d93eeb1e861923184c9b7761f2f62b34ed4c6c144b960001b616689565b610fa67f26f2acb11d4a85a0b9f3d86047123706369bcc118754ad2c56107ca53919e6c560001b616689565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d63877b9600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1686868a8f886040518863ffffffff1660e01b8152600401611051979695949392919061c229565b600060405180830381600087803b15801561106b57600080fd5b505af115801561107f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906110a8919061ae3c565b506110d57f87a0c90b3a7863ac42cda6fdbbc289506793958d8e4d463991a22dcf6d83c82f60001b616689565b6111017f3c6f218d76c4a30fc2743aa8e375dbc2c2569602f411e25e7ea2bbfa9a99ca7260001b616689565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166247c00385856040518363ffffffff1660e01b815260040161115d92919061c808565b600060405180830381600087803b15801561117757600080fd5b505af115801561118b573d6000803e3d6000fd5b505050506111bb7f8408b5e8d4ada4cf69594f0e3ee32c4a928a008e60db3b07a310966efb3eb93d60001b616689565b6111e77fab3b512d579a8b72c48470a5390ef8dfd7944e18b1d37d9951941c298d770c4960001b616689565b7fb1dc0b58437cd20174f0de80b765e50f8ca2ef9591496e576a65034e7c4d4f4560014385600001518560a00151604051611225949392919061c399565b60405180910390a161125f565b61125e7f0a79f63708b5baf897508d92510454f1c8e2ea96ebdc589adfb55a1aacc5e97760001b616689565b5b505050505050505050505080806112759061cefc565b915050610308565b506112aa7fe4cddb0c05d8a8762a2b2578bb7f9204e271466da33b1fc04bdfca62740e0a1f60001b616689565b6112d67f652a15132a17159259bdbe25dade82cbefbaa46f163540bda85c73c4ceb4031360001b616689565b600190505b9392505050565b61130e7f71df4ff92a47d21b07b77bdd0ba9c8c0a9f7756eb512cf07e803bb1c4a1d6e3760001b616689565b61133a7fcaf72246348cc6edc05d98eb2a3d86ce035cb741f03152dd856c53d52c5f24af60001b616689565b6113667fcf571d53f661f1c46d02f541660064e50256b8b17c1b6df2d0572ae31a741e3b60001b616689565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663181197f783836040518363ffffffff1660e01b81526004016113c392919061c192565b600060405180830381600087803b1580156113dd57600080fd5b505af11580156113f1573d6000803e3d6000fd5b505050505050565b6114257ffa6e2a9f4d246a4cc97ea86d76e44706e1228cf8588c3e22f5deed6536b291aa60001b616689565b6114517f7df98702883a46de76d7dbe3c0a0f2c19fc3dd84c15755505e9701fafcea869f60001b616689565b61147d7f11fd93252e4f829deb2cdb9244abc046a73965cf4dfe7141ce95223006faaf0a60001b616689565b80600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002081905550505050565b6115177f34ac7d6dcfaa3b4d11a1f7bd50df7643a26e312445de9e260cbed74723a25ea060001b616689565b6115437f95fad68d566d7da9391f6e69d8b4f8e8026bdf19f636b018091b8a5378a3631960001b616689565b61156f7f1bf8f24b3693669f581a40d9ba8405d8fd8169c6adead4b879483bc7fee1ff7260001b616689565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166352e06146826040518263ffffffff1660e01b81526004016115ca919061c170565b600060405180830381600087803b1580156115e457600080fd5b505af11580156115f8573d6000803e3d6000fd5b5050505050565b61162b7f2f6e9d3aefdbf502e66ccbc8d48720ef56fd7d92379b5a3f43ba48a6a7a4b00760001b616689565b6116577f0efcf1103787a1b63f0826fbe043957c5b11c299b0fecbdb9e2725775d14f21260001b616689565b6116837f31da2ea4d5c18a0d7c100604445bfa1d15cc2d2a3c1e0002232c9f3b876ce69460001b616689565b6000816000015190506116b87f1b227c221ef241c34b598235089c7d79d6b01679bc0b45fb97428b684e2dd1fa60001b616689565b6116e47f4747042fbb68806831cbbe3a2b96ace7b7bfeb5b6e9f708d35bd31c32409c07b60001b616689565b6000826020015190506117197f22eb937de68e0974957f3fcc156a5dc949d86a1c970908bd3aeecb4caebdfffe60001b616689565b6117457f238fbcc1b8f4369e764d126467e3af4199f13839e85d9bca7e1534e8dd82b29e60001b616689565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a65374a846040518263ffffffff1660e01b81526004016117a2919061c0ef565b60006040518083038186803b1580156117ba57600080fd5b505afa1580156117ce573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906117f7919061aeff565b90506118257fed99d247fc233c01e154a99842c326accbf6ade9cfeef7609c594e2d54b7d25460001b616689565b6118517fc635e5d0ed0fb1a75f712a07359aa34078a2ae06e8dfa3a812c66994747e596960001b616689565b42816080015167ffffffffffffffff161015611980576118937f09a2de91f7d60d409a19644af2f180afa575b51673509d44bccc5b8bf7434d9060001b616689565b6118bf7fd608d8e58be9f8a2159eb598992666c1ff097ddd36755c04ed52adfde492d54360001b616689565b6118eb7f28adc1a0c77a51086717203fe88eb86508179be4dd0bac1da9a1c17daba6b19060001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f56040516119189061c422565b60405180910390a161194c7f7a80b7c9a8c2a5ed88dc6727c0da75d4d6640347dd53a1a4f5d9bfb31bcabac160001b616689565b6119787ff2ee18171863fdd7b91c7b68f0638dcb70af236fabcf3f3027eefe92798d10c260001b616689565b50505061243b565b6119ac7f1b82e4c4bf31b6605db5648242275e15b575f0196987333dbf5231c610e9e32d60001b616689565b6119d87f28e0af72030ef0c98a6275c72ec60827b7f1bbe45e1ba59138a312f161e045a660001b616689565b611a047f73aa32948b39a7fb2f097f6eb960c832c134e552bc13c091e787b44acdf27ad460001b616689565b60008267ffffffffffffffff161415611b3057611a437f6381b8e0280557d645bba46c5dd923a11b4e7454bf4e7b3741e361fa9022038a60001b616689565b611a6f7f2a998b449b0c1428ba176a5dd2b727be34aebe478a1307c73a33e581a6f0a84360001b616689565b611a9b7f5b86b404db162a0e469fbeaa5e466c0872b1e67b8912ef45c72832b46ba2695b60001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5604051611ac89061c455565b60405180910390a1611afc7febf1495d43b6b6e1f6ddda261933451668563585d597a426e621f72a1f6b4e7560001b616689565b611b287ff28be862649908850f010c6ef1921108756efb970ce69704ec00c4ab5b46c99960001b616689565b50505061243b565b611b5c7f1b2447d498c95a2d8fbad04727c1110b14c62903216e3118d14e946e91b6d25d60001b616689565b611b887fedc29eb0997b46fb08214df78e368db0e408e9147e519353eb4e097793f199e860001b616689565b611bb47ff8576549f61c61e1060f7022a86b946ce6f0fc82d2f1c7882ac37c3debd9528260001b616689565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d7866040518263ffffffff1660e01b8152600401611c11919061c83f565b60006040518083038186803b158015611c2957600080fd5b505afa158015611c3d573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190611c66919061af69565b9050611c947fee6d94b53a7db12e2d16660e0eb204216619fcd454760d9173cdaf09e25190a860001b616689565b611cc07f7b37572bb9a9022af6f7e9abc841842e37a198de89e35aaa39739169e76b1b4b60001b616689565b600081610100015167ffffffffffffffff161415611df157611d037efeb91a1a072d5703d78dc22221216038743be99c05918381ac7a584ff1e23b60001b616689565b611d2f7f4c018bc4b7aeaae1390cbc52ee7fe07dcd9dedb55ce1ae9b56a9238e85023c8860001b616689565b611d5b7fad7bd0b61bb4d54d81c2fdf8aee6c3eb0a6e8ac7221b611e85db64f483f7ec6760001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5604051611d889061c455565b60405180910390a1611dbc7f134f121f6e16717e74cecb93d027ab2f6e687f1cfc818613a105424138d2ff2c60001b616689565b611de87f04dc74a85f3b1323828998cb7637520eb46fcf286a43c8d0d983b62fd874334b60001b616689565b5050505061243b565b611e1d7fcb41fb96f831ad7e93f7f85621ba7f1f0bfb20b302cafe4bfa83e5b7fc061d2760001b616689565b611e497f779731721316997c6ed77215b6e4b891b1384d088075ac7945831db995d33aee60001b616689565b611e757f12eff7823fcac789050bb9cbfa4fa0fdf5ac2a85abbb1b5787944676cb2a888c60001b616689565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c5c81b2083608001516040518263ffffffff1660e01b8152600401611ed5919061c3e5565b6101606040518083038186803b158015611eee57600080fd5b505afa158015611f02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f26919061b129565b9050611f547f4da8f75ca349f5f07eaa004461001e25cdcdec97025187a4833fd947148c25ef60001b616689565b611f807f61b728957a0a241378da094e6abb957d500448a0dc8341f8e6339b58055cc70260001b616689565b6000439050611fb17feb18741943d444bbd404eabfa566ec10cade8e54b50500e576df2caf173384fb60001b616689565b611fdd7f52e0b22b4d166141ba7ca7d1c94515fd1f101745b8b18d47dffe2c827647908860001b616689565b808260c0015167ffffffffffffffff168460c00151611ffc919061cb49565b101561211e5761202e7fcfab248e1c5d97cd18d990f803396ca7b0bb1a954be23885e8feca78d94e669660001b616689565b61205a7f3d23e6c553581b6f929ffe0e8a4382293abacec516b29f4ff7de85cf4ccce70860001b616689565b6120867f56dcbf72e3969f12c624e71c3e2aae5099c25b9eede90f782e3e5c3a048105c060001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f56040516120b39061c455565b60405180910390a16120e77fe9ce957c00085d35bbb46c3b07df47b03c98180544aea0cf65e26b6f10436e3760001b616689565b6121137f0a8d36d0c33ee07764393cb5e61a02fb40078ca0c65879848c698ce67551f1b660001b616689565b50505050505061243b565b61214a7f03bd7b4d3f57c0813a17c6f57bc4b8598bffe0fbd7eb2dfa50df6a24e93de86560001b616689565b6121767f50202b7905cfc4eaa56ffa107b4094f66a294fabe36b4ef47b80fea5a5bc36f560001b616689565b6121a27fdf01756b2cac2d707d81128a2018b2bec20373bfaa78c0053f7b2a49a657b79060001b616689565b60006121ae8787618b81565b90506121dc7f7723b0ebdef1c20c36caf00086f9918aa4a0bfd3f550fed4b23f87c6833623ba60001b616689565b6122087f83df60494323ab23bc21621ca3a2dbef6cdbe68c60f654d7df0726c84c1d370e60001b616689565b600061221685858486618c74565b90506122447f431c05bfc5c49ff2401eb82e5c72b96aae47fd7de2cb579c458ef3bfe9936f8560001b616689565b6122707ffb39729cfacb0b935ee908499d710ad58c16ce01779317d25e3ccc807139bb4d60001b616689565b60008167ffffffffffffffff1614156123a1576122af7f449f3ddfa16fea32298c61f8b6ef40248189ebe0575c17b0cd22dad19557d6af60001b616689565b6122db7f9703dbbd8b61df0a7cea0a5ff7e106cba1090460011f368d7b8ea26e311db01060001b616689565b6123077f75a810181c91a6fb6af8fb48f657cf9071c6cc8a3b0a4972c11254deff8d5a1660001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f56040516123349061c455565b60405180910390a16123687f4537d3c43fcc299c8b9c9e9d08d642f3cc8dbaa0d51dcb800c8abfe0581548fc60001b616689565b6123947f5eca77280042437194bf6c6ea1a95fa600a5208ed3ad07aa5e1f2e4ffa7d212260001b616689565b505050505050505061243b565b6123cd7fa3f4c3f175bf160241e854737cb33b09e582184970412a7a284194674c3b325160001b616689565b6123f97f2c6ea7fedb6e731d0a82f0ccf5e317d95e7e5f1f0c68f8a314771621d939007860001b616689565b6124257f462279288aa83df58d1ae81c1deb74da47c1c7c061ace9c26737823b2e130ca060001b616689565b612431858786846167c9565b5050505050505050505b50565b61246a7f5f85ef22dfac40a33b3f495f9eb1350432751b62de72255949d00ac157bb134960001b616689565b6124967f46f9676451b5ab7f4532b63ac115ae7ea3454f97f049e8416e1470ef4680765060001b616689565b6124c27fea553f637f5d8a914f91addf543fe683e54fa421d15de35d63e295b7ac01f93960001b616689565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663defce5d483600001516040518263ffffffff1660e01b8152600401612523919061c170565b60006040518083038186803b15801561253b57600080fd5b505afa15801561254f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190612578919061ae7d565b90506125a67f1c1a0a4c8f23f596302638f7372ad34f3798d3305b61bc8972194f398d9b18e560001b616689565b6125d27fb2e2aaad0ab37a5427fc5aca5bab3bd74903d31b94932a9483411ba18e5184a160001b616689565b806102a00151156127e4576126097f40cd5c41b91d315fc4e069750c1196397e615bab01329f709409c8834fee315460001b616689565b6126357f4de29afa45493ab84d48a6c0ac5af6bca2a07fcd79948d69e62cefcdd45070d960001b616689565b6126617fc4907b977582772db06f7d4127427c87367b4dcd8af0dc6eb40507a0b302744160001b616689565b806020015173ffffffffffffffffffffffffffffffffffffffff16826060015173ffffffffffffffffffffffffffffffffffffffff16146127b3576126c87f3196ed18ad61f5dfd0689ddd4c14a34fffdc1d39c3003c0e3499643c0187add360001b616689565b6126f47f154241118b30393f634f9d086edf879d8e068a08f682313514ba0a1f96aa4e1960001b616689565b6127207f61ec1e82402d9d4c9c3341dfcaba8151caec64724cca53a363652d98c2f0837e60001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f560405161274d9061c5de565b60405180910390a16127817f6647820278f1da0d7023386e2bb8cfb690cb3bf4ab7631e04284b242f09a45e460001b616689565b6127ad7f7a20cdc1d7e6c4ad11ef4632b804d31a5fafdced7a30a7276617cf6a2328e79d60001b616689565b50616686565b6127df7ff94d5edaeaad4db0896db7ca9e7c4d8d3338f4072e07917690ea47b916e1072560001b616689565b612fe9565b6128107ff4cba7ba8e2377dde5cb7820f681e0e80249ca9a583d80f70cb8ea5f31606c3f60001b616689565b61283c7f49fdfc1e902aa5743f01180d81eda9b72f499ff4065a1e41a3e673a89f16d0af60001b616689565b6128687f1380fbbe8dcf304633e4cdd310aae62be76fb8850fb92379178dd6308a44f65d60001b616689565b60006128967f10bf3d89eda5687063c445dbf9e0d22d2606000851a9f7d7b24c6e35c54fe68760001b616689565b6128c27f47fb92d16e2cb7d71b7dea5b0c50157a06f59fbf21a260ecfa7f1cc710ff3c2160001b616689565b600082610220015190506128f87fdfe8e14b88ba6ff1b6e0cb189586167c5ea8c97f2c6bc96a6139a3e69569feb560001b616689565b6129247f7e1a4e66e9cc3b580795b15cd4c8b6897aae650d33f27b933c04d588e6ea8cd960001b616689565b60005b8151811015612af85761295c7f057ddf46edca86eb7f00637222619b3261320de87407fa9f3e0a6ccef6f99c1e60001b616689565b6129887ff6849ccb9903a9fcd173131f432e5df2e174896a84f28d3dd47da01b9b51f4de60001b616689565b846060015173ffffffffffffffffffffffffffffffffffffffff168282815181106129dc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff161415612ab957612a2c7ff94f209a7fe3fd875458384e815dfb40ddadeb2aad4f608b56b48e6fd56ae13d60001b616689565b612a587f16e6386be2c7d36513454f4046ef1751c9e6f5d501b91a383cd963bebe551c5b60001b616689565b612a847fdef73e009295af699733477617781f7baf4d1a362e8efdec64428e0f3832fc7060001b616689565b60019250612ab47f60d06f8a22d9798644c83fcc4c0182356445bc00c2db232440b13d3612778ae960001b616689565b612af8565b612ae57fd6910c03f18d6a510ca81da88170b6e1068e7e34afc4ca329958277efccf938460001b616689565b8080612af09061cefc565b915050612927565b50612b257fd84b47b459b1ce3fb7984b4b1c71063940e3bdffd9f479228c153a0d81392f2560001b616689565b612b517f0c0dd0125f355b53632f6e131a32bfebf9426c246b90a70a685b2da8c58b850e60001b616689565b81612e1757612b827f0f7aa51b137c35a98cea1cd044389326c37203ebf6cf89876dec998a475b280760001b616689565b612bae7f3f74455a5e14bf19c1c1f8ddf10d89fc0ffb92fbb989d62a86b83925573ffc8460001b616689565b612bda7fedea1638f64ccd0fecf9cad20051e1f1af556db193b263b167a06047daf2777260001b616689565b60008361024001519050612c107f4b1f05ac12ea814592f462b69321996f3d4b944aebe83fd646b76fccab5e75ac60001b616689565b612c3c7f183896cf402e7d6065542be70e2a5f7862fc4faa183124116bb4f5392f169d7660001b616689565b60005b8151811015612e1057612c747f54683e71ac6a775dd8e5052ff868ac11cd54b290e528da1e7b6257cac057b09960001b616689565b612ca07fd325c02e3f6c874ad3cba34e9f934898a40f5304502607db1ff24705d13233c760001b616689565b856060015173ffffffffffffffffffffffffffffffffffffffff16828281518110612cf4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff161415612dd157612d447f7015100c93e015c251e59d9b1c17eb9ef9920913ba85c7ba591885dc1780966260001b616689565b612d707f27b723be59a808a17f001549c96a02ac20afc889fc9c46d2140d88fef459871c60001b616689565b612d9c7f33e381ae7eb5f7cfb3eb59c3fb3be1008cdc0f55ce52a177c80f1d9f589967f960001b616689565b60019350612dcc7f6469bd8aaf86321de85623ec47052e30b0dceff90ad50a8bb6e0a9992a3f8faa60001b616689565b612e10565b612dfd7feab3512b76dc9d0a3402c372bc8c68ae7e275986b60ee49434d3d1385eb3331c60001b616689565b8080612e089061cefc565b915050612c3f565b5050612e44565b612e437fd4220e08b37055facad043442eb76bdecad213da7be069bfa5255d6c6efc6d8660001b616689565b5b612e707f8cba125a8ea65c23eeaf188d8028fd93299f965c2485af89683a7acc0bfa3f3860001b616689565b612e9c7f694b7006cb1aff13a093f563fb2f02dbc7d72b701d78d018d01195a24a61839460001b616689565b81612fba57612ecd7fa125c54a64b4f68e3e53b6ea2c0799aba47738e70ba16a6887d3abf4a921eb1760001b616689565b612ef97f8347f961d7e88b48c4fe19517c2b45f96e481851af8c8a6adb4820b2e4cb40fe60001b616689565b612f257f173d888a066df7fdecf4a61797138e6bbce63dca71bc0886290488301ba247d260001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5604051612f529061c776565b60405180910390a1612f867f6f26249de4993a1beb9a5d8cbb6fb5e593cad721b21a0d6228e1840bb33b2d6060001b616689565b612fb27fa2406e418e49e3479cfcedb77a2dcdf7e46e9232d8ee63edf43423fa193b28e760001b616689565b505050616686565b612fe67fc03b78965696e00c2d5ee301fbf8c6b64fa0cb8f304fe87fea463899938e192f60001b616689565b50505b6130157f561f7d330adcbd6c94febcc524377ec0ebb1cc95dde6c36a1c450898268f72d360001b616689565b6130417f48ba6325894538064ff3a509057c1ff7ae85d4a0af662bd215beb52e0890886360001b616689565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a65374a84606001516040518263ffffffff1660e01b81526004016130a2919061c0ef565b60006040518083038186803b1580156130ba57600080fd5b505afa1580156130ce573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906130f7919061aeff565b90506131257fd6b50d8341fca8c01bd13a89f9315faf6275cc3398e19adacf73b808b002d1ea60001b616689565b6131517f1eaee8c64e7c83378c683ca4dd8a542926d0736723423c63e6aa6def2abef56660001b616689565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323fdc52c600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1685600001516040518363ffffffff1660e01b81526004016131d692919061c1f9565b60006040518083038186803b1580156131ee57600080fd5b505afa158015613202573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061322b919061ac0a565b90506132597fd2520db128be1426a346509af7574c40207daaf583bdbd9a47789333d7609e2a60001b616689565b6132857fb25d505d3e36cb41f38f082942c73ec92ecfbc6d6610edbd469d7f3dc17e273960001b616689565b60008460a0015167ffffffffffffffff16141580156132a8575082610100015143105b15613570576132d97fd34ff95d5ac1ed77816d994378bf7120581503a75dc7648a3b308648f3408aea60001b616689565b6133057ffc699b0daad118b69ad09f0dd53ee1c0223b27dd3cf7b1c767a5b5423d26750360001b616689565b6133317f76b455d0996693cde843bc7b6b5af18464397de1ceec4e0b95d247d01f96b64f60001b616689565b60005b815181101561356a576133697f868065019a2fdbd396cbe7c6df5dd9668f2066a79b47b6afc34d351fc810553360001b616689565b6133957fd4b8ce9cfdca31cb71f86272985887040c3a2b660bc888ebfb34d5c2491f30fd60001b616689565b846060015173ffffffffffffffffffffffffffffffffffffffff168282815181106133e9577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff16141561352b5761343d7fd6998f04ca9407b232f5d1bd07ea0ef383fb377d15508e45dd4f6cdb482b2d9260001b616689565b6134697ff8b60575471ec025c708429ef9cdd983f2a9c07f2b53ae9da32158578dfd64bf60001b616689565b6134957f5a93cb76020e177394c8ece3274cb340786f61dcb914eaf2f9d791498b5986af60001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f56040516134c29061c710565b60405180910390a16134f67f74e926096b21f10d76933d6f83418e0b438492832d4be918ef058a5c788d13c560001b616689565b6135227f0df7ce77f2a3d989322f063d1fe48cfded10465069f9c56ba2db2486e2d000fd60001b616689565b50505050616686565b6135577fc19ea389dd70191aa18be92ffd445a4746b34bbb74f0f88f69e8cf08b21a96e760001b616689565b80806135629061cefc565b915050613334565b5061359d565b61359c7f6323da6a603ca32026b826c8c99ccdf0573ffbb61e3964c1e0866aef9e39f6f960001b616689565b5b6135c97fbcfe502efe8186cc9fcecfed17f6f4e6ba8a987a4091ead5a79794771b3015dc60001b616689565b6135f57fe67a63b9092a5cae906549713b7088a6243a27959ddf7fe7e11ce5572d0191c060001b616689565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d19c1f32600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1687876040518463ffffffff1660e01b81526004016136789392919061c2b5565b60206040518083038186803b15801561369057600080fd5b505afa1580156136a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136c8919061ac4b565b90506136f67f322043e3ccc358278f0d0121cd8406bf70549e0614f515649df95c0342481da360001b616689565b6137227fb1876fa65f7acb39ef40af796e3970c04aa4906daa1264b0db825c79c8835d8260001b616689565b80613841576137537f19cbf92daeece1f87ba6804203e4dd75b92ef06cd6f88d0131b9caf6e5beae1060001b616689565b61377f7ff9dabf0367a1f85c49eeaf5319c70e2c369b7540415cf83bcf6343f097abba3760001b616689565b6137ab7fcf16e79c12e4f0b19c95278154217fb7fadfdcc52b87cc41afcf6becb8c4e94060001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f56040516137d89061c6aa565b60405180910390a161380c7ff289ad10c788f3eebbb709117fefbe7227f193d720294aeede6e9ebbc9c8107360001b616689565b6138387ffced4198b8d6931e07efff860e3360926fa863b1703bf9eb4ed07a842dc800ab60001b616689565b50505050616686565b61386d7f16941b71cc8d56e0e86f9e9441d2023d0ba9ac52e6f2d9855e4dba0ace4541f760001b616689565b6138997f2f83d452d89521a050a50eceb66aa1f779f50de54f6b026f0bc70e62c82e84e960001b616689565b6138c57f898a584e98e69a9ca1e53a69a245fbba1b13ebec5fe6b46f5e56911cbbd10b6b60001b616689565b60006138f37f2241d5e601e2936d121533f92a8c88daca1544568e462fc072f4c5adb782f3cb60001b616689565b61391f7f59ed5fb7c58dad201cf823b9a0636ca2a35f753dc64f2340a222c28b5be59c1e60001b616689565b600061394d7f56d00d28b1ec0b9ffab0386454b6375463d6adf524ff4e4e9e32e81e6983115860001b616689565b6139797fcd6e57abb6b7aed2b1d47110887606e08f25511590b765f83c507e9246b5352f60001b616689565b60006139a77f5ac519b8b5bed2c0bd5f7decdf9514cecc1b5fef6fb6930302ca4bf094c35ac860001b616689565b6139d37f559c4935e345459312fd09640760d41bbf80273c8f5bd28b8db66ddcf81affbe60001b616689565b6139db619596565b613a077fb99ed542fec444dbcdcce0828b1726f9a3e66a81c4276ceb772c94a4b682929260001b616689565b613a337f75be63674bd8ff43dc0128a612b29e3a54c5b2da1ea98eec6a8fb613c5a58f0660001b616689565b60008861010001519050613a697ffbbc01a7c393669835e7499b8afca3f817cb36fb4441d2820b80eb2dbe72373360001b616689565b613a957f9a6310cc410310da4ba1da298b8571b92c856c80624d96c480a38cd1f36a72e060001b616689565b60005b875181101561434357613acd7f3d8e4753b0d125c753b189867c4a55750fe994ca36929296dd93ccac8c865cd060001b616689565b613af97f9afc969c0a62734cf79447cf44637cb05fea47d0d7dd3603a29a576f0f55c25360001b616689565b8a6060015173ffffffffffffffffffffffffffffffffffffffff16888281518110613b4d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff16141561430457613ba17fe9821a20800af258e8131101115c38dd8b9bf12a4bb197688aba8dfe8dbf4aa760001b616689565b613bcd7f9936b2351f8ec787d9df1532670600d1d556880e3ef57e5a964be060e68db19460001b616689565b613bf97fc5d88f095c3ac7d3791604208b8b63d1cb972dca5eb99597a95c8b2ac38c2d9160001b616689565b82604001519350613c2c7fb9aac04104c795fe72cd451409727dce1e5f9630a362b9482c2ec272096fbc9260001b616689565b613c587fba76aee975e1515cc66e8e3ec790d0745733e86d794518157709579d99f016be60001b616689565b8960e0015167ffffffffffffffff168467ffffffffffffffff161480613c7d57508143115b15613d7857613cae7fbc0542a144690afd515f8500c40b3f598dca28b51dbaec4989bdee0df7f4a85960001b616689565b613cda7f646d42c9a0eed5462bb292596441fad603b1dc11a722f99135c2f10e20dac09f60001b616689565b613d067f5c50fd9ef9781e6a9d1ba2b8c7751c156d32c3d7d2287dba4c319a4dfa9768be60001b616689565b6001836080019015159081151581525050613d437f23130b45f094f9c5a3c67e629a8097a2570dceecacfc878b522b3a9aad16fb0660001b616689565b613d6f7f1271fc1afc69555fe1a10a49abf84b8eb2f33115a58a6ef9cdfc9babb1e2c90c60001b616689565b60019450613da5565b613da47fdef310875fedf1490341bb78c88e43e7bbc70ae1e2cfe314c8b4853b1576737460001b616689565b5b613dd17f9d510fcb4f78c4a050459971e2c38def9fbf49e0ec0df430f0879a22506b954260001b616689565b613dfd7ff5eb2534ea4faf7cf1595501fe25d2b40f00ed3ea819b049f1d1d636de38454160001b616689565b8960e0015167ffffffffffffffff168467ffffffffffffffff161115613f3d57613e497f5e59772b98f62bd4e0334b2289642ed000fac451b54e6331faf74101f241eed360001b616689565b613e757f7a4fb93602535bdc8c8c2fc6b055c119e84f8056741f0d68d21864081546872e60001b616689565b613ea17fbe8f78c210dbd9361bb0f6568557f80097fb1aea60cf8916c378d61ce4bc0ffa60001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5604051613ece9061c677565b60405180910390a1613f027fb7a7062a855379ed7fa39f97bc28fc02faacf6a9de3af5d4d632ec376af008a460001b616689565b613f2e7f9de6342480e80053f4f5da630bfd83cfb0feba051ce246ca2be0dd195c59957b60001b616689565b50505050505050505050616686565b613f697fdcd070660d30a6d7e633bad0b9387bc7012babd907cd32c5f363710a04bcbb9460001b616689565b613f957f6ceca726ffc97eba2dec0b65a6e802b6067a44ca98b04151dae93cf103448fd260001b616689565b613fc17ffdb5d743669d13dfd1c87ca677b9609bf378fb15d1ef6c42e96604d4ae4dd90f60001b616689565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663624e1dd08c61010001516040518263ffffffff1660e01b8152600401614023919061c85a565b60206040518083038186803b15801561403b57600080fd5b505afa15801561404f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614073919061ac4b565b90506140a17f3ec815210fa65e96dc3bdb02082ad8a52a6339aac822ad92732ffcee163d19ea60001b616689565b6140cd7fb258b278e6859a328e5d2e0b00a33489bb0abd68fc65ea176f188bf4896b6bf660001b616689565b80156141f4576140ff7f803fe12f258cc39780baf882231cf5f83e5846c33c9fee6ce20020a88ce6cdcd60001b616689565b61412b7ffde209dfba836f9e2a47cd953fad471fa6dfce0971afe1d186fe222c61038e4660001b616689565b6141577f9fa48c6dbd8807be31e99cbdb4277142b4f43177e75fd4b2f79fc417f4fdfb9960001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f56040516141849061c644565b60405180910390a16141b87f2b2a81c6690e3399518f49677919993acfa92e031f118b78a0a54fa29ed06dbc60001b616689565b6141e47f219c4c3a911d428f6f66f3e8f2d97ef5b7ba61bba7c0203ae05be3cca35ea98960001b616689565b5050505050505050505050616686565b6142207f8f55fc7d92d8848b83ab70619950603e79f567b1b25672b056793b042184d96960001b616689565b61424c7fd2d16f5904a93fa4ff6781d8d3031ef2b0dec0a609f6bb0753aa8a40b8a783fa60001b616689565b83604001805180919061425e9061cf45565b67ffffffffffffffff1667ffffffffffffffff16815250506142a27fcd78335f5b1297333eb40d248c7078d6afb08ca12ef90c70960a03b13d02b3a960001b616689565b6142ce7f4ef76c5373868ae1cb6ddf8c493025102222d4ad86f1b29db0ea0763051c4f1460001b616689565b600196506142fe7f2c730680c9630b0ad0e2ddbfc8386e99fed259b91c2fc235f734a86bb070304560001b616689565b50614343565b6143307ffef7f287c7b5f96c6d02804bf7a96d1f4a54510831078552da5ffd46738e023b60001b616689565b808061433b9061cefc565b915050613a98565b506143707fedf508a0defcb52cf71cabb807b0b358bf280aa1118182078474c368da5904f160001b616689565b61439c7f06e997bf50c04977f25663648c717062019300817c6b8fc3638f639240cf606960001b616689565b84614e69576143cd7f8d197a9a6c071a993a46bed0019b56adc025136d02589e5af364d632dacf249860001b616689565b6143f97f5337af193e804560be31facaa78ec45277d755b77bb3d76dc3c180d723a1ef8460001b616689565b6144257fd96a03050d7250da356bf1a0bcc414622a594693efdba2befd32849bf5465a7660001b616689565b6001896101200151614437919061cb9f565b67ffffffffffffffff1687511415614568576144757f579b8527f550c7d14c48e037e3c776dd74b8367641b0da19537b8d9f699b16c360001b616689565b6144a17f68178fa1c8a8fe03ea4bc7af703ddfd6b7fd83848a228968057439c564e28ce160001b616689565b6144cd7f61d85cf498c5ca6fb946ae62ac9f70811fdffe7e375244f085963217f069a20e60001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f56040516144fa9061c6dd565b60405180910390a161452e7f8a912884a3cf06da526109956092fad6413886c209d816a0467cc4952990b53460001b616689565b61455a7f2f250f843ab9a9c7db2b2576b80efe0434015e95a36eaeced7cb0694e0eb3d0a60001b616689565b505050505050505050616686565b6145947ff1d72059ec45d65fd08b792be39765f24d2f58fae4f27dbd350b2473f5da615b60001b616689565b6145c07ffde1746556cbc214f6243f719f24cbc5d1fdf463af521f9efec3ba4ae8bf65b360001b616689565b6145ec7f9487b7a3c05c78ceb690facec0bd2f67dd0f89537ef1f5b212ef132311b4fa1560001b616689565b8860a001518960800151614600919061cc0e565b67ffffffffffffffff16886060015167ffffffffffffffff16101561473e5761464b7f4181c3c6fb06f0db2f141293179da57f892b7e98fab8ab5c828f3ddc5ad9383360001b616689565b6146777fc8e04f49e84e3319476a22905cfa5634c120f2d09c42c08e773a0206584dca0c60001b616689565b6146a37f617d9655e16e78dcb1e3880be75f4e9e460a94878a35339257deeb0e93094b6460001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f56040516146d09061c743565b60405180910390a16147047f3fe5609e0b1fe4d7211c30960c93159345e48dc4bd3bea3752a61a0bac53037760001b616689565b6147307f5016af6b33df7d7b00af87fad3e21eee0989581ade8895d0a6ef7b4b718bfbab60001b616689565b505050505050505050616686565b61476a7f3dad443de2cf2679ccbc952c3b69462499a76dced36ac8515c054b9575a2f2b460001b616689565b6147967fcdd42c7c51b72dd6e9b3ece0e351cace3248f196fb69b34c383ddae972ab429d60001b616689565b6147c27f28633bf20b0e5b11d1e74e9b8ae8852c7857fa9ce37370a8fda08b3e0346136560001b616689565b8860a0015189608001516147d6919061cc0e565b886060018181516147e7919061cc84565b91509067ffffffffffffffff16908167ffffffffffffffff16815250506148307fba718e57122ea1b36b67a19cb8263d67c8b2fdf029a3893803caec7ec565b62960001b616689565b61485c7f408683bceb1ea9d17b5927d463d4922c530964fd04b9f1545e39ebfb1baa312c60001b616689565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fc215358896040518263ffffffff1660e01b81526004016148b7919061c7a9565b600060405180830381600087803b1580156148d157600080fd5b505af11580156148e5573d6000803e3d6000fd5b505050506149157f9e5e264582d2405548adc495e56c042d37df56d25a1b1ab272682751a47d5ed360001b616689565b6149417f6f0605a79d80b04b5ea8e904666c4c588acf468756b0570891e66917ace3799660001b616689565b8760c00151826000018190525061497a7fba7645343e2c7d8c5f841fd34c713378612182cdade43f94bde48b990c12662960001b616689565b6149a67f28f77dbc3ef008ed8813853f2ea214259441c3f45b35d88335425851d488d54a60001b616689565b8960600151826020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050614a0e7f865ea4a1e28763f0bd3a90582a0a27f9dae11663db98374cb8da97da41c3810960001b616689565b614a3a7fccad96e7b9b0ee2b25f38923a40eb5b982cb94ddd810ffe16d0f2c34de8bcbb960001b616689565b6001826040019067ffffffffffffffff16908167ffffffffffffffff1681525050614a877fb5b662721291077a195e7c37e481c31b8edce2d6f57972806d4dfb8742e0588a60001b616689565b614ab37fde6440a961f70dddbffb83ee8bac34c5098857fb95ea1eaba7878688d301187660001b616689565b43826060018181525050614ae97f5267065fb838760a51f71fddd38334f5ad7bdb7207db09039b3c2058113a566760001b616689565b614b157f8e2afd9d8e6c3d2c0c4a71d94c970e905e7e2d8501fd1faf6fe2226d188548f160001b616689565b6000826080019015159081151581525050614b527f9d4e98256e3adddd1463e93c6eae8def35509c1f9990356dc1f4db95030c312760001b616689565b614b7e7f7e9c6fc34227e27e28e44de92e924a2eb4722af7bb3e2ca669164c20b9a85b7860001b616689565b600060018851614b8e919061cb49565b67ffffffffffffffff811115614bcd577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015614c0657816020015b614bf3619596565b815260200190600190039081614beb5790505b509050614c357f258ec7f7824f7111596f105bfecc6aff85c79e27d207f29f462522f28fdad20660001b616689565b614c617fb59c82e54913b7fe53b433e46d8b724cbd0c87426c2baa4929991e43c15f699060001b616689565b60005b8851811015614d5d57614c997f7a475fec3912cf1630196f648fceb05a28f01ed6f88717b66be220bca96e9e3360001b616689565b614cc57f05de970411c64ab952e76fa57360f844c07e7242dfc9670aba8a8337a47e4b6b60001b616689565b888181518110614cfe577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151828281518110614d3f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052508080614d559061cefc565b915050614c64565b50614d8a7ffa8780f832554c1053f79a89f022c2fb8eb083516874dbf857eee3209e1c888560001b616689565b614db67fb7c5a46f81d0625596b5dc724da624cd3ad5825c072992740b5c0afa70aaa66d60001b616689565b828160018351614dc6919061cc50565b81518110614dfd577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181905250614e347f1e8df5be6d792b0651f8cae95fa5b23e7a16b669d470e35ad559a3a8f4a8b12a60001b616689565b614e607f2e4c8e77b44d2a587cd43b43549ab443b5eac69675bd641175a58f33b8b4f96160001b616689565b80975050614e96565b614e957f5f1549ed98a70c1b6ae1a727d646e5f83403563b8347f3bcdbb074920828ffe660001b616689565b5b614ec27f2262ac0c6ca060e89b4f939482f8f9ad4ccfc9923485e75155dc9babe10b93cb60001b616689565b614eee7f443c2f2a462b6db7bd6bf813182ea95f2d424c1154279c8ca1a31465253c5ea460001b616689565b614efc8960000151886112e2565b614f287fe0cc19f6af402c234201f0c2f8bfb586fdbaf27698451dc1b07b0c9a2abef99e60001b616689565b614f547f469dc7fe9a470c92aa4047665024678b3d54a7506c0c8bcd361833e49b0fba8d60001b616689565b60008a60a001519050614f897f19736d67a4e6b6cc672945619327f3d4cd5eb94cd2a9207a0b0a35823035bc1360001b616689565b614fb57fe94a7d3c8ed59a84a7b32849aa46bc93e0f57ffed80de32e0de58291962f7cad60001b616689565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506150087f175bc03b7ee975662fe1d371fc7322669d291e791a68e666713a46e49440856860001b616689565b6150347f293ad5e510c463fc6858d53d2338b567a865962ebdebf2d6921b4a43ee172d8460001b616689565b60008b90506150657f9a216f6e9577791a145a51fcd56e8d49047a9722499886e57a5a3f46f3eced5e60001b616689565b6150917f85979b963fe63e767959368e6586f12c04d1dd728295d8513140884f7741efcf60001b616689565b60008d6040015190506150c67f416dae7ffb89e0edda74033cfdc9c3422f382a7436c2f64d07b196c42c842cef60001b616689565b6150f27f04bb5a158682d12d67654e8eacfed588bc3092a85a30395c42436e6e59bfc9a060001b616689565b600060405180604001604052808e60a0015173ffffffffffffffffffffffffffffffffffffffff1681526020018667ffffffffffffffff16815250905061515b7f9153ef37b724b865934aea2a01384d0e4f8f80e3f28bd04651f4e23b248ddf8760001b616689565b6151877f85e54eef8644a218ff92c4971e1c80bb30b12e8e4d6a55ae9c75431ec8491b0460001b616689565b89615cbf576151b87f30ce213ca774f3aa223c947fe56ccc34f32caa8d4128a65db9f336fecf14808f60001b616689565b6151e47f95bc56e7bcd1e5dc8529d3167ffc1efac96845739de6bd93d905dfc7af07572360001b616689565b6152107fd4cc12bafba502ff004598008f15a89249f62a9ecd2184ed3a503e88005a2d7060001b616689565b60008473ffffffffffffffffffffffffffffffffffffffff16632ba010d7836040518263ffffffff1660e01b815260040161524b919061c83f565b60006040518083038186803b15801561526357600080fd5b505afa158015615277573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906152a0919061af69565b90506152ce7f20cbdda1f9950a0e3c62076957d1d02a0a877d8b224ca3f3fb53f46317e28fbd60001b616689565b6152fa7ff52ade7834dfa6d6019f1989237db34e62fb83105255bd1e30461f9033f9d62f60001b616689565b836102a0015115158161014001511515146154345761533b7fb1c819d1ec4f0fddf4d9e97cf43e6bee164bf8d77a91e812c2d5c78b4de0263a60001b616689565b6153677f8355ea64c774eba80ef21f78f35925a7960efc88b6956c566f79deedd8d6357660001b616689565b6153937f6062b167398af84e61604d24b65ca88ada854b3ace366cefbc59b24fab97516e60001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f56040516153c09061c611565b60405180910390a16153f47faafc2b701e8d76607c8b8088a7633650981e891584734cc6265749c9e33a8fce60001b616689565b6154207f3a9784abc054ada9937718ef69414c940035e4b6013d2e270be91a854ed75e3960001b616689565b505050505050505050505050505050616686565b6154607f0883f1e5e7b87e70cb7a13671067441ee94956ff174607964662c4cd5c6b697660001b616689565b61548c7f9f0ccce39d80b9f420ac45ed0d6aa8121f4926d82a88c875c4759c76f37999a860001b616689565b6154b87fa69f351b18f478d741f104f98d5cbb0ece520f7a3192e511280f5b1e513f1d6060001b616689565b8473ffffffffffffffffffffffffffffffffffffffff1663955f98b782866040518363ffffffff1660e01b81526004016154f392919061c808565b600060405180830381600087803b15801561550d57600080fd5b505af1158015615521573d6000803e3d6000fd5b505050506155517f27a8e7008ac86add9446fa442e126c39dd223b420ad151d14a71345692858e8a60001b616689565b61557d7fbfdbac2fd90c30097aed5d69122697db71e67015bf5e1c020f75c16438fa28c260001b616689565b8473ffffffffffffffffffffffffffffffffffffffff16632ba010d7836040518263ffffffff1660e01b81526004016155b6919061c83f565b60006040518083038186803b1580156155ce57600080fd5b505afa1580156155e2573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061560b919061af69565b90506156397fb17d4fd1ce5de5600c210881dcd46e5cd071eabe376761e18f1e79973b99d4df60001b616689565b6156657f44c97b5714dcddbb05d9e78533fcd72c1fdc7301d9a7adfcf55a21ce5d1722bc60001b616689565b836102a00151151581610140015115151461579f576156a67f4d2b2ac0d71f5eb9c9a660b02a0f2d3caff649cf4ea8612c41509d5247dc8cad60001b616689565b6156d27f892da246e3bad94a5abb4436facce7c6cf951d4f4e6e68930c4c6c35781e9f1660001b616689565b6156fe7f25385bd693afb486c5ddbc3f8939cc40263e7f70b3a2330e329bd8fa9d1db16c60001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f560405161572b9061c611565b60405180910390a161575f7fab156d9c91d056a1589d1534bbcf35b56ff7032fa6234fe2e1fd9283a0ff8a9960001b616689565b61578b7fb35f5bc71a2a12d4f097ade69498893f3409a473601e722c540bda23d32ff03660001b616689565b505050505050505050505050505050616686565b6157cb7f56019d7e6a6a3192028963f98c2b1b508d0dabca7a1846dd6d17cf96c063b94e60001b616689565b6157f77f5fc9e4a44a88fe1139c0705d84a9127f5d40c887877e570a1b4d9ea8899c845160001b616689565b6158237f1576a94f2669026fc2d5bf9338bd93b61f27295f6ef7e2a01cdd31a07ebe53c760001b616689565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663de4d268b87600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1688876040518563ffffffff1660e01b81526004016158a8949392919061c34d565b600060405180830381600087803b1580156158c257600080fd5b505af11580156158d6573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906158ff919061ae3c565b905061592d7f25d534cdd78b63be61720903a11136f3321e1e7ecdb609d4b4240a991868c24d60001b616689565b6159597fe17df7365a21bb076b71d887b5542e7acde09d8d006883cf72b614f3bea5c77e60001b616689565b600081511115615a8b5761598f7f0119369d2e2a2f378a97a0a9a1eca79fd5d2c103f47cc6be106cde3a7443c5ca60001b616689565b6159bb7f72e0a9b7da05c047eca70017e90e00009bbea4ef09a189ee8ea40b2e7526b72460001b616689565b6159e77fc64a17d8495c517c2e9defdc2442da4d45b9a20bdc2e1b73eb7b8367b6e9c51a60001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f581604051615a16919061c5a9565b60405180910390a1615a4a7fbd5780e5a291a71826beb49c1da2bae6f1d97b8089d1f8d3b7794e135e3232b560001b616689565b615a767f49ae4fa2f647c01cca367c36f470c019b3d433c4abdecab61b69ccc8905b51d460001b616689565b50505050505050505050505050505050616686565b615ab77fc246aa9932aab50855f77c3ba7a8212e6fd526f5acbc71a0a53fe22553f45b5d60001b616689565b615ae37fefeb3b253d3af9baf889ddd6e1105fc7023df455dc16794de2a2c77bbea1f5dd60001b616689565b615b0f7fbd45fd1848e5fe5fd8fa430a6f23309339c43d017aa710557da26ce1f79028c060001b616689565b60008260c001511415615bc857615b487f279151516655f10881d83cbc00ed5b381d1970b9bbad985b9407dcce48edc95560001b616689565b615b747f50f927a8f04c99b624e2dbbce0d4c665706e768e6f0d3fc1f2555b5e86221d3060001b616689565b615ba07f08a16795acbd745ef698867270d24c6411a8cd94fb23b6517b5fbffdaa5c57da60001b616689565b8460c0015167ffffffffffffffff1684615bba919061cb49565b8260c0018181525050615bf5565b615bf47f698a59800af6cbf0d694906d1af6f0d851a6db7ed581bbb72a61a7eb56f717f960001b616689565b5b615c217f387191c4f2379591359e8f2da502b961b8d6aafb57e48b2cc3a21303e81f5bf260001b616689565b615c4d7fa385030a5f207d16ba6ad63fa2f7e28585a32418cefa9812827e94a7fe9294fa60001b616689565b8573ffffffffffffffffffffffffffffffffffffffff16632384a6aa836040518263ffffffff1660e01b8152600401615c86919061c7e6565b600060405180830381600087803b158015615ca057600080fd5b505af1158015615cb4573d6000803e3d6000fd5b505050505050615cec565b615ceb7f58e9abc12036c89bce13cf8add00825668bfe41b24a4935bba7d6988fd42416060001b616689565b5b615d187f770f96e348c38770eeb22108ed9500b10eec9743eb918296bc43c400c69b558360001b616689565b615d447f0cb7cfd94fc4dfcd6b30c5ae3d5e375af3714bfbea704ca533629a4ed18f976260001b616689565b60008d9050615d757f7a06f877ef94cc23f5bb8b09b6e325a157b72d333a4a828252b1054c458e2f3f60001b616689565b615da17ff4ff3c206b808b47c7f80dacc4b23a7b6927f5a2f260a0c704774642a4304bbe60001b616689565b6000889050615dd27fe119c20d531fa01dac3af6fd7b2c47f2d5656890cf619ff5e7c7382d07d2d1c760001b616689565b615dfe7f51ac4be057ffe91154cb2c7e4c1a6f61bb07f43f9143eb4c1cb47c735d7ce56460001b616689565b60008e9050615e2f7f08480f7fc7b278e50f4736c6482ce9aceec6e6403b3500bb1c602b09f402c74760001b616689565b615e5b7f242adb204ed7cb8fa427401d2fb005238eb32ebb66db74900f2d0835d017d6e760001b616689565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b158015615ec557600080fd5b505afa158015615ed9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615efd919061b129565b9050615f2b7ffcfebbdb3ea8cc624158f9f7b915f41ed5ea1a8b1b823ed1469375a9453ddb8060001b616689565b615f577f17cf61d6db4f88d11c363fcfdc7dea6bf728bf8c06184559bcea1c634c3f0a0a60001b616689565b8c156165a857615f897f0affa4d0cffda842f46359b4161792a8f0f29cfddd99ce7fbda16c63982a473660001b616689565b615fb57f5ad6a4f724bcc8098591c64d46b3c6ca06035395732ed9fc2f6b9294fa75357060001b616689565b615fe17f044c8f126ad756a294d196bfbf8d509b34d7b7d869f37625249781d93b6b2e1d60001b616689565b60008967ffffffffffffffff16146162175761601f7f1faf0d98e36d9dd36c941754e2e3820f8b4d8f1d516a1923e6571dae4f358f1760001b616689565b61604b7f33931cd62fc4fd246b1a603a79a79a98a816ca200bfcd6c3ba37964e54b9915c60001b616689565b6160777f1ccc81516ede87e3865087732e32146ab962a2308ad567deb21f989ce0e5398b60001b616689565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d7876040518263ffffffff1660e01b81526004016160d4919061c83f565b60006040518083038186803b1580156160ec57600080fd5b505afa158015616100573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190616129919061af69565b90506161577ff6e9298b2cfe8e8a844aa3d65b5c427ed3da4dc350c6848166a519af787fecc960001b616689565b6161837f94ea6c00e40a04bb25130ecba0f746327168b2cae702b801bb3b63941c8c79c060001b616689565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166247c003828a6040518363ffffffff1660e01b81526004016161df92919061c808565b600060405180830381600087803b1580156161f957600080fd5b505af115801561620d573d6000803e3d6000fd5b5050505050616244565b6162437f496ca4691d2eafc9264ad37fd564d33135a52d576e008df930568f6e2673271160001b616689565b5b6162707f0268b8045dd458b5a1aa0bb8dd249e4a0b093d44c2102f1116e2560740e655b660001b616689565b61629c7fb90abbffaae7b9ea42e82f876c3f0bdae7d91fc4e3fc7430d1c626316f7ddb4260001b616689565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d63877b9600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168b898989896040518863ffffffff1660e01b8152600401616349979695949392919061c229565b600060405180830381600087803b15801561636357600080fd5b505af1158015616377573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906163a0919061ae3c565b90506163ce7feaa9644f51fc6c416c9a617e00adff8e30e7378860354d6a4a714c6c30252daa60001b616689565b6163fa7fd82b48cc490787d7ca3a0eca6cd51ba78f5af7a31323387197979aa25afd64bc60001b616689565b6040516020016164099061c0da565b6040516020818303038152906040528051906020012081604051602001616430919061c0c3565b6040516020818303038152906040528051906020012014616576576164777f0426add888eca79fccd7b09294b81477805725d13c8061729a7dc9d68b15afa360001b616689565b6164a37fc77d4861fbda2f3aecf0b7a284663b779555a4ad3b4bc45b9005e4d54d6d56b760001b616689565b6164cf7ffb2a06557cfd80b678869ab88144fa0abe2a8fd94ef437f9650e6c5f70295f5260001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5816040516164fe919061c5a9565b60405180910390a16165327fc9d3be598d3508fb27c0b4427a4fb014f9d9abc580c1235f97465c621b83a47d60001b616689565b61655e7f8749b1be0fb46e2fc2160deb91979a9f5f3b7511276fda171b1aa32d15fa2dd460001b616689565b50505050505050505050505050505050505050616686565b6165a27f40088d432a13a6253ce379b22b1b4146fce968ec7fdd5b14a5d82896435bda1560001b616689565b506165d5565b6165d47f51f7df9b20bd77df4a180340cdf6ea3cee3ea5c48955ef5ef7d8499f26db025360001b616689565b5b6166017f4df20ebfb7e30e08d85406b5e74cb13c8ee23ac0adca9c4ccea352cf7090685960001b616689565b61662d7f9bef352b600a4337db65d8f81397ace06b2c3d11c382c5254d5870c4d516f9c760001b616689565b7fd936063ae3cbbf2dbcd2adc837a997ab97adb09f24c566aa79d101e88928dea660074389600001518760a0015160405161666b949392919061c399565b60405180910390a15050505050505050505050505050505050505b50565b50565b60606166ba7f42eab3da22c71474f2e8e3d11f9ba81d7d3849e8a0657334ce7136d8dc4bc04260001b616689565b6166e67f65ad0079586d6e316f9e48f1356de288ab67194a95af67c25613fc590741e16760001b616689565b6167127f7f17d00ce2c3f70fe33b9d1c31098acfdf5309e28d38be28b4ea094d99b41f9360001b616689565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663977fdfe2836040518263ffffffff1660e01b815260040161676d919061c170565b60006040518083038186803b15801561678557600080fd5b505afa158015616799573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906167c2919061ac0a565b9050919050565b60606167f77f2dc3e74ed06327e750a04f7973d9cc4ba92baf1735ce8999760adda3b6f498f360001b616689565b6168237feb649525fd3a622e4f9e65c07eff8b78febfa42355f887b8e1b8ff8e6ae5fb2a60001b616689565b61684f7fcce9702fae1b162d4ef8eadd0d9d2bc3b19818372c33569e522f2ae97470e27660001b616689565b600061685b8487619362565b83616866919061cc0e565b90506168947f337e6c06b90f105f1ce6ef2274bc30f65a720afc67faa408cf1c1456e83e5a3260001b616689565b6168c07fa7bcf1c07513d340543f562a55703050018dd2950064574aa8d99b185ad6dd4f60001b616689565b8067ffffffffffffffff16856000015167ffffffffffffffff16106169975761690b7f8be2f03dbf2f9c7a8ae4a7cfff6052cba8c923b2c9aca7bf04e70a288d38fce860001b616689565b6169377f87e10a80b0d28de62def846345b5b0d3eda2c0dbde3da872dfc4c23ee8ff672460001b616689565b6169637fb61b422991e4fa7a47f113d799bac13f3ca0c8c649e66d5e2d85f543bef3813860001b616689565b8085600001818151616975919061cc84565b91509067ffffffffffffffff16908167ffffffffffffffff1681525050616a9c565b6169c37f78f1b655ab91ab6dc7765648a5d9958404963b67ffad387a22356ca56b430ffe60001b616689565b6169ef7f25d53b6c0ec9a6956707a546d3ce52e9e538674a4c12377ffaae3a3ff0670be360001b616689565b616a1b7f4e551dc288230aae9c54de3b939fb744a2c30108c9294c0ee4b09e7497265e9360001b616689565b6000856000019067ffffffffffffffff16908167ffffffffffffffff1681525050616a687f4e42490f9cdde24595522a5fbad07e46ace99021af895b5899dfbc2250ee8a5060001b616689565b616a947f457174c9c2fc44155cb25c3655f8a267e83adcaa4d8e97bd59b82f14cc931f7660001b616689565b846000015190505b616ac87ffcf91e8d1862e4a8424b03130cb818afc05988c9e39874d41c1ce3a5f1a746de60001b616689565b616af47f79aff1ed82d9c5d5bf71aef4ed57a1695b9ac3aa2f66516934ec9598a80d3b1060001b616689565b60008167ffffffffffffffff161115616d7557616b337f9ad1774e1080cbd124fadb150322373913292897a4d21185b12e66a56a3ebed260001b616689565b616b5f7f183376d1279869d2032f843b01e70cf6680fbeccba931458df878ae5b71fefbd60001b616689565b616b8b7f11b0743298566c0300bc734d8b551fc95a97761f822e76ae89ef25d17d9ce06560001b616689565b8067ffffffffffffffff16341015616c5f57616bc97fcc77bf3360566559f59a7543e0e0f83369b1056ae517233d5bd3162656826d1a60001b616689565b616bf57f7291841fbbf6972b890f3cf9d18f39fdf50cab360179aca816d91f514064e04f60001b616689565b616c217fb4bdc8cc0ebfa9708d05192f6b701e5cd89c9af36e251590dca8109bbe0591c260001b616689565b6040518060400160405280601681526020017f50756e697368466f72536563746f72206661696c656400000000000000000000815250915050616e79565b616c8b7f8570568c44d599a3cb6638278cef3c85086d81d56401a886537f985e038d226e60001b616689565b616cb77fd878f4ecededfa2f20fa34c02810601432d93ae89c4dbabf495b94c916ad501d60001b616689565b616ce37fc12c2be1004e803d44aeac5c4d3b1be8658f604a759322d9facf25da64ccae0560001b616689565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fc215358866040518263ffffffff1660e01b8152600401616d3e919061c7a9565b600060405180830381600087803b158015616d5857600080fd5b505af1158015616d6c573d6000803e3d6000fd5b50505050616da2565b616da17f5b1b0c2fce7a97be7061ede99e8d1987eeaacdd3791b374059e6ae4bc66db1b960001b616689565b5b616dce7fcd6a49fbab9ab8c281058aa4828824a743f83097947e0b968b7dd3cffda0f08060001b616689565b616dfa7fbf05df137c86520dbfb992acb1bb1ea1bcae91488803ad3041e87849318fe71560001b616689565b616e0d86600001518760200151436113f9565b616e397f7183cf7776337e6bac23b5ce3afbdfca34778a8bfff0e6329f567ba087f8ef1a60001b616689565b616e657f2964200475fbc9fc1be8f9b5caab497149d0e34d59582e270134468bab828db660001b616689565b604051806020016040528060008152509150505b949350505050565b616ead7f77e3a988cb79abd920d13dc4d1442c213c3973fca948689caa0ebc29a2ad026160001b616689565b616ed97f78bea9b74f45b74959056ec99488d4177f869e68532432521d9f8365f7e4042360001b616689565b616f057feb45fd09ab78ea9c6f0b66a279a3d930c213f244f3457d2e606ae03845cac3f560001b616689565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bb4e4e4283836040518363ffffffff1660e01b8152600401616f6292919061c1c9565b600060405180830381600087803b158015616f7c57600080fd5b505af1158015616f90573d6000803e3d6000fd5b505050505050565b616fc47f6d0f2364c8abd9bcfa223c14106c24cc5f85503cca50cfe9c6256d57470806b660001b616689565b616ff07f9d11f1b3bb12b1eba51efbcf06ed9f0a302572731bc6b2ac1226d4a1f1e830ca60001b616689565b61701c7f73112e6f01f06ee973717eafa4ebcc6df77cd1ba13367c8f895a8c3a33b63f0c60001b616689565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a65374a83600001516040518263ffffffff1660e01b815260040161707d919061c0ef565b60006040518083038186803b15801561709557600080fd5b505afa1580156170a9573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906170d2919061aeff565b90506171007fcd2b3575cc27716ea6d7c9e9420be810cd44ad288c0f25e25579a7dfced341fe60001b616689565b61712c7f6f1b81d806acad174a30c6945df5728ef42f4493af34c75c8389a0e7f144121460001b616689565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d76040518060400160405280866000015173ffffffffffffffffffffffffffffffffffffffff168152602001866020015167ffffffffffffffff168152506040518263ffffffff1660e01b81526004016171c5919061c83f565b60006040518083038186803b1580156171dd57600080fd5b505afa1580156171f1573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061721a919061af69565b90506172487fc6a0d67c1cc8889235b3a9111dd4313d18cf347f79f07b8d7d9dd97d91f5ea8b60001b616689565b6172747fe3745b2ffbc3704ed57526431998ca0deb22513c3580339e4b721d51ed9e211e60001b616689565b6000816020015167ffffffffffffffff1614156173a3576172b77f5321ca37086a202c4983b8e5f5d66b34f3efbc43e3e9731a7e20ff5932ff169c60001b616689565b6172e37f4a0279dbcd756e23b5923e794ad95fa4f2193b065e3b7f78e3c80b3380a88a0560001b616689565b61730f7fab4b973685e8d3f448cca591458956afe5e53c4422a6113fad3243cdfc7daee360001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f560405161733c9061c543565b60405180910390a16173707ffc8d7475fa00156f917f6830a0d9c01d28b866c5f54bb149ad8482f81705a0d260001b616689565b61739c7fe260844df8d7e3efe7790900be85a6be49abcd4f4fc748d5b7f5fbb23c92364560001b616689565b505061864c565b6173cf7f39bb3fce3360601641d52f4888acf79988b74a94f1aafd61791a97acb8ccf8a460001b616689565b6173fb7f2c01a6639d313235923591310d3169744f938b048e698c5cff4e5c4661ff382160001b616689565b6174277fbbb6adc178fb9fdd6d82a97bb540f8fb85ab2751959f9eb65fe34d5abc0d3f3260001b616689565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561749157600080fd5b505afa1580156174a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906174c9919061b129565b90506174f77f2746af48b2dd555320f5160ce11b3d133b7def05fdb0818cfcc50b116bd28a3160001b616689565b6175237fa2afb0a44ae6704c95c3113ccdf920692d11a37404db841a9765ff3455c6ff8e60001b616689565b8160c001514310156176485761755b7ffa14a14fc9bbf2e42113b57fc82c8d93d67e295f0f9c7318b1587bd2a5459bbb60001b616689565b6175877fc31722ca1681e6e8b510eef7046a5057f457587befd916b6ecd2dd09f22af37e60001b616689565b6175b37f0eb303f976c6cf70eb88e2b8bdaa620bb251ce8f4dcc753a87fee0b0402f183460001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f56040516175e09061c576565b60405180910390a16176147f148600aabc56b98d2cbdb1fcf0f81dae48b393fb14c3a493e2347711e69b8fa860001b616689565b6176407f6f127d275f179817211cb11bdd346a13704500ff8def732745f7dff3de433ed660001b616689565b50505061864c565b6176747ffb71fa9feadaea42976834c178731384b0f92b2af391a89089d9c693776cb46960001b616689565b6176a07f6992741b0179e037e80793c3e4a7bcb41fd25a0995479579a1bb5e4df0615ff960001b616689565b6176cc7ffcd42b6ccbb5761fce1921d9417cf33ce4cab4cfd6529dbc3b59f3a4632ba2d860001b616689565b8160c00151846040015167ffffffffffffffff16146177fe576177117fed426d12dd0600af0d2ff97dc6ad1b8f9d3753017a08f5ae366c3e95a01c4c7160001b616689565b61773d7fb173593dccd4e051acb063072a634f4f2fa95fa7cbb0cb0e31e1a0ebda277f9760001b616689565b6177697f83e3052e9fa86d98956c626e6e573df8cf31e566865827ca10bea5aa283b3b1360001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f56040516177969061c510565b60405180910390a16177ca7fa9d8f6e84dd4a7489360726497e3e59b0f576dc085a74d75fc0b31a97561015260001b616689565b6177f67f107e0bb873f002065230dace1ee6f387707c9b0b8b8d19330edf1201e26d0b1f60001b616689565b50505061864c565b61782a7f31494691372c000c48257f03ebd102cf0f4d9f4fcc95ca6e717da51fadb7620b60001b616689565b6178567f8458b44ae888999f6fefa74b10d65e5e95695fe453f490674b1758779ae6dfca60001b616689565b6178827f387e408c2459b38728934097fee3b0cb499dd0eab9a7d2ec665bff9d2c0c530260001b616689565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d047290e600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600560149054906101000a900467ffffffffffffffff1688876040518563ffffffff1660e01b815260040161791d949392919061c2fa565b600060405180830381600087803b15801561793757600080fd5b505af115801561794b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190617974919061ae3c565b90506179a27f655248a02b78ce56f68fc0f1778459664c6de6627e6c03f4f45e70ccfed5b08760001b616689565b6179ce7ff2347dbbd24e09ca2c4b381afda5f06fdedbcaaedc22cf4885f06ba03a20a08060001b616689565b600081511115617d0957617a047f342694e1fc4591037559a5a758485b1954ee1cd538ef2810bcf32032d2f8977160001b616689565b617a307fb5d3d946e1c0b9f49e06b4b618210b59195f450c35324b5b5e258487c56a68c460001b616689565b617a5c7f69c2ff71d909c10dfd0e71fc7ed95b6b9393c7503b492ec856c1ab2e9ee7178160001b616689565b6000617a6b84868560016167c9565b9050617a997f2af0db830e4cc2181345fcb59eb1f9ffed5a0ea22c228d526c9a691ea4d33bd160001b616689565b617ac57f480c52fb92764f3da3e1c0527be85af74f5fb602791c14b581876dd59f64fa0460001b616689565b600081511115617bec57617afb7ff110d9cba96f1642518c43ffb6e0684cce8b4658ca7d456dbd065bb3e12ae6d560001b616689565b617b277fcff58256e9c90a1912033f1e5ec63dc57c643ccc0b636a20d64d0c5aaf8e0ff560001b616689565b617b537f5896d3db5e43d31fa2814f43c15342e911a6e7bebf451c9c1393385a6888dd2060001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f581604051617b82919061c4a8565b60405180910390a1617bb67fe65e0ed3e69b3676412bd4cdb4160b5c8ef277f9fa1b3c24400e7b13e5796d0460001b616689565b617be27f5f448c8d26d7266831c9ff0e30e1bd496389fd9ddc62af098f963bb61897d75c60001b616689565b505050505061864c565b617c187f13cc12d0fc90c4cbeb40ca7fac392ab3438bc81f246a6ad3675faa5ef562cebe60001b616689565b617c447fb6eb29dace601d4a3672d6ab4598ef2c23eafeac1a536b82fc45e7cf1b7775f260001b616689565b617c707fe8a93941a4ebd190ab3e81c8fac1ef415b2758c89e2f1e1de079c6555560517760001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f582604051617c9f919061c4a8565b60405180910390a1617cd37f5019cb21e43b51aaaa8253fd522b3c1ac15c10de733448ab06e1c44a441047ac60001b616689565b617cff7f89a4c76cdf65ce2247a8564574d97abbe5539f798d27ec6de5e186e24fa143d860001b616689565b505050505061864c565b617d357fbbd5c988781505bc7ca4bafcf15f056ff712a59babcb26cfd1c8a5d3d40a35f960001b616689565b617d617fe79271fdfd1ff8caaf208e8102f97b11dd9c388d6fea58a4c900b67eb20a9bb560001b616689565b617d8d7f84dccf235ba9adf4eda61e1b9a70c507d5c69ecf3c7c67fd3ecb272f78d45dcf60001b616689565b6000617d9a84868561027f565b9050617dc87f632720a588847165941012b79b4e63a9d43aef5a44b2f0ef5834fa40048d8e4060001b616689565b617df47f74c109e973c133c9af63ba45965d8ac97eb8124ad749eb9dfdd15ac4676a9e8060001b616689565b80617f1457617e257ff6f2b441b0ed92aec11315cbaba58a7997c9426d4d57b34054b4dc45624adfd960001b616689565b617e517fe2d4d0fa3fb90c02bbebcbc5a92c4f4f575972983d4b02b667562646635d2e7760001b616689565b617e7d7f4e03257b707f735fadbaf2ffe82d63a5d99660354e14b6e5f828665a450521ff60001b616689565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5604051617eaa9061c4dd565b60405180910390a1617ede7fde29e3cf07e5a28ecc3ba09411b7872851be34131f3808f35a550804f64ff20a60001b616689565b617f0a7fbbaa500dd57fd741ecb72c5432bed8bc4fa3529568b4f639ea1d139e2d534a8560001b616689565b505050505061864c565b617f407fb9f9441c590574cd81f925313b0af1f1fad09317cf22d1a45e42a801187f0a0960001b616689565b617f6c7ff05ccd65f4a954b2eb6584e3fbd37855bb8d0bc74f2eb8a65379beb9c319f08360001b616689565b617f987f3857e7fed3b81b0325dbb070d38c00d4fe666c8fc3e5b83204cf762e46f6788060001b616689565b60008460a00151141561803857617fd17f6934e24617cbf885fb426686d5fb14d521eb985bed756f1e6bf61d09db6ff9e060001b616689565b617ffd7fb9cf7bea8e9074268596235591247f833200e97b443a32f4941d69b980aec52c60001b616689565b6180297f3cf31bf378be9f47157370ad07e3657d16b6558982434386cc6af97a3c46ea9960001b616689565b438460a0018181525050618065565b6180647f05f1431e53bb9c152feb1eb0ba895ec8bf68517a22d22fceb3b1555ee43147c860001b616689565b5b6180917f31134a187de735708f26121de58f5c40eb4d32e54741ccf65be2f9e14b4be31760001b616689565b6180bd7f13ee3d1a873ddb36a79c4ed8b23b2b61453c04297fcbb24672c87a14b4e70f6960001b616689565b8260c0015167ffffffffffffffff16436180d7919061cb49565b8460c001818152505061810c7f33db634171f55364c338032b7422f1a718b944ac2322d2132b58f2f1e4a6636a60001b616689565b6181387f82ac0407ad8285f2ec12dc52f5f46c089e524f5cad34527de195b8b916f3392860001b616689565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632384a6aa856040518263ffffffff1660e01b8152600401618193919061c7e6565b600060405180830381600087803b1580156181ad57600080fd5b505af11580156181c1573d6000803e3d6000fd5b505050506181f17f7684309b83fed64cec6eef0f617d4d26fb9a3e92bfb9747d04348675b5d6fdc160001b616689565b61821d7f8305bfa411b5fb2befe0efc3276cf93a4a3f0d020cea5f6e6de3cf275f6d3f9760001b616689565b8361014001516182b5576182537f5ff316874df1b26d3479f52a97f83198fb0f327b4af53cb586e6e604700c781a60001b616689565b61827f7f0435f83aa5886a1b346e6f74644c0c2a2bed72c8cbcb57bc3e0887be9ead395e60001b616689565b6182ab7ff86c2a00f2587d6ddfe569b6a418a3cbc606036e90bb497b44ce28635288757760001b616689565b505050505061864c565b6182e17fd00f86b9241a73c0fdde31cdb6d328b80d71083712750a0cfe943559dfb7e8c760001b616689565b61830d7f949f6408fd3bbfb306b1456c6b32eef9a540caf3c9d205a148dd1341081023ee60001b616689565b6183397f671ac03c34d77d6db8ce245c14212e4b39517cdab1dbc54e9af2e063bc02212c60001b616689565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639908f2bf8660000151436040518363ffffffff1660e01b815260040161839c92919061c10a565b60606040518083038186803b1580156183b457600080fd5b505afa1580156183c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906183ec919061af40565b905061841a7f0a2f9021ebc686f38bf8717e2be89c1d80f55cbb1c1153726ccb52da0bfd0d8660001b616689565b6184467f5831bc7661b4be44015cc7cb62d082c2feb24bd78d03149dcaa0df73efe12fcc60001b616689565b4381602001818152505061847c7fced5b36e223ceaf0a73374877b13ebd696beef26687a50d83234c11d23cd865560001b616689565b6184a87f5fc43201a99c18d0fdb9e9e1cc240031b545a84b3999f89fbe55d470b97321c660001b616689565b8460000151816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506185107f0b0bb5257f8575f8b69c88462152a793fc782c9ac28e5269948fd0cbbb9f39bc60001b616689565b61853c7f3b2bb7c19f1ba7de9f2597ed9712825b2670d7c11c38421af6160c99c9aa63fa60001b616689565b8460600151816040019067ffffffffffffffff16908167ffffffffffffffff168152505061858c7f58ef28a5bbfb73a09f728f66efa040d1821b40137cee127a21720482bc08e08060001b616689565b6185b87f1fbf759866b35dfcaf58535296a61ce24d6bea00a823b82dcc12ad327bd1869a60001b616689565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631581d545826040518263ffffffff1660e01b8152600401618613919061c7cb565b600060405180830381600087803b15801561862d57600080fd5b505af1158015618641573d6000803e3d6000fd5b505050505050505050505b50565b600060019054906101000a900460ff166186775760008054906101000a900460ff1615618680565b61867f6194c6565b5b6186bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016186b69061c488565b60405180910390fd5b60008060019054906101000a900460ff16159050801561870f576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b61873b7f9ea2bb4f8fe479a9013ae7d5974299e3d2b636387220c7e34c5bce7be3eb3cc160001b616689565b6187677f51f5b4d6b36a1cefa859dccf0f8fe440df7024cfe82d9e2d6738e537f03d6b7260001b616689565b6187937fecfe7674311fbbe10133bae32e6b6e8f42b6f97ff6678012ed855008a9cf13a960001b616689565b87600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506188007f3f31efca014eededb091b57c92bfc1bcd0fd437e424f70b194e333c1594a4d2d60001b616689565b61882c7f70687339367acbc036b13850f2a473936adb65758bf5af764c3b8d09d4db17f660001b616689565b86600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506188997fc22359bb0ba2861f3325598043813b69226bd056fc9f256a3c418bdade36dc8b60001b616689565b6188c57f09f226d3c3086332e1d4e84d275903838f33288cb813562ed4d7ee8e27a9c3bd60001b616689565b85600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506189327f0a20267dfbe10a78ef185554a811e7cc8f70eb2f55e3e04d40299e6161f7f51e60001b616689565b61895e7fd2c6a01b04debc30b8eebcf5c6c8034032423414cae8f1d6a46e008f7e08d87660001b616689565b84600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506189cb7f2f543be3dbf757aff14807484e271454367182c74055fba696722fdd4fdf420260001b616689565b6189f77fbb4bebd0515152fa6b028f53c08d9b082c3420cb0c563a8ee9a2f8571c44ab5860001b616689565b83600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550618a647f50703c9ed2b0d536ef07fb62bfa966bce93d89aff69cc9f511f1178a9de55b2e60001b616689565b618a907f9161ee60a9e75e0bc40ee494d71aaf56d48084b0d1b72e18a84d9b5e2b33960360001b616689565b8260000151600560146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550618ae97f5351380353d71f4d437fe7190e170e4449f8d787fc2bcef40dd38839035d156b60001b616689565b618b157f3729b74bb1e2654e092e8aff528b9785d947348d79d447ecc0a7cc0c79dcedbd60001b616689565b81600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015618b775760008060016101000a81548160ff0219169083151502179055505b5050505050505050565b6000618baf7f6721bc449323e5228e1b8041967abfad25c86ccad80e392bd409ec07db4e606f60001b616689565b618bdb7f8bf2e8c4f1f5a00b1b614313e7866c699da80c183fc7077e0769187b54565c8660001b616689565b618c077f14a7369437be1394211ef740d62ef652d8108e09e7cd077db88f7fa3cce6cd2f60001b616689565b600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002054905092915050565b6000618ca27fce7a05e9f552fc10dc5901fc63f56f1fdc6f4cb4307336d797ac432adfe6b5ad60001b616689565b618cce7fd4a8950fe6517d07b0ddbff88348707cfc0a0ab571b8a66b8a3eba242749bfe860001b616689565b618cfa7f7da27cad8c4705d325c028183005d6b88d2ae1a6f198ac41f54fb9fd8c1569f260001b616689565b60008460c001519050618d2f7f1075b93c60caeb351ceef0854d0bea84c6608510590730a6489a299e94568cee60001b616689565b618d5b7f7b4c522d6e207706e839eb6324d2c0bc919b0904f7d4e93680d197d3f88e9c6260001b616689565b60008660c001519050618d907f266cba52989e95322618fc17221e16bf5bfd537f0a6ca4601ae8410080901f5360001b616689565b618dbc7fcaa9a7995c6d4f165651863854434678d459ef7a53b2d559ac657a449222f16160001b616689565b838267ffffffffffffffff1682618dd3919061cb49565b10618e6757618e047fecbfb7f0c5f26fc221858457761d6720901ec88d61af0bd36cf645e0cf81e1ac60001b616689565b618e307f16a46fd6b1f2e5a60ce91d3ee1d9e54c652354329d93ead872c98e99fd25631760001b616689565b618e5c7f5f8c5b6718bd5acb690238e7463e1e741ebbcda2866828ef80feadb6f203211860001b616689565b60009250505061935a565b618e937fe95bba7cd5f2c5cb4a0f88206a918202b71dea924c234d64568bde2fcde8087a60001b616689565b618ebf7f50eec388e5ad4c5d3fe6ee577d2dc9f688e9532456a859c755fc41fe2874606460001b616689565b618eeb7facac5867d967b55683f650cefceffc1e9c048979b2dee60f4d45a428654d1c9b60001b616689565b6000828286618efa919061cc50565b618f04919061cbdd565b9050618f327f5bb265b82ab75cfe4b0f3431cb1981fa3e896e1c3445664a4ac6794785503bae60001b616689565b618f5e7f6e7a018472b25e5ba59831076ee2c7cbfbae9a3ca3da435efb8c40eae9878bcf60001b616689565b6000618f8c7f44f50733dac1727ce5abdfd1b5f029bd697b75c81c9b59a904f4cdeda81be6a160001b616689565b618fb87f5d607b53424da1060e66d9e1447d4972e78578cd568b40c8ad682455677d5d9e60001b616689565b6000871461919157618fec7f9b9c6708ab4d24057a5a6a804ec919a8c24296419bd921c5075e0ceb5152453560001b616689565b6190187f2ec1df7b5ca60fc51ec14bd5d4261b79f37708636ef29e9bd5d6cb87a45cbce260001b616689565b6190447f9a9d476ef473a571a1fffdaf4c54627af6c002d4c3169531a0dadf41d28adcdd60001b616689565b8367ffffffffffffffff168361905a919061cb49565b8711156191035761908d7f8d1d77d830524668d88060e5c1f004ed300e4c88db857e1797cfe62e4b0818d960001b616689565b6190b97fdbb6db3eee13604f0aa8029b03999ee96cbf89e9923cf8cbcb2557a4343c522b60001b616689565b6190e57f684d9aa69b6f139f99ea24d1829aa2b816b004358faf252e77c2d9179e0d6c9260001b616689565b8383886190f2919061cc50565b6190fc919061cbdd565b905061918c565b61912f7f82110a22bda580241f4819883fb296a90e4a8434fa329494876fb9031c93e4fb60001b616689565b61915b7ffd6952ec8180865dd7523e6d4e7dc15084b9c5b771c2eec06447eb83726f5a9760001b616689565b6191877ffcaf039d484c0e2b1930616a575753bf90ad821d942d3eff4411324fc059d1df60001b616689565b600090505b6191be565b6191bd7f2f9e1d3a76aeca934f0a067360b1607a9350ebd0b8af06a8bbe4be526900f6ed60001b616689565b5b6191ea7f7b6c53f09daa0084b6ef1430ccdd630318c182c6ae8df00b0b2395c9f860982260001b616689565b6192167fb9ed2b04a68e9b3e51e5d9022d5dfc51ec97b78dbc80bc263dc87cf206f05d6960001b616689565b8067ffffffffffffffff168267ffffffffffffffff1610156192c35761925e7fbdc7b727a983f967117cdbddb6da08de37bf6db1eecdd8d076b2a34315cfb5b360001b616689565b61928a7fff30f2171f1db255a036ed70c8d9ea21bb09a09447e47479777ebdd6b71b560b60001b616689565b6192b67f7aefe052566505ec6705769c090ae0e9c5df3e94f862bbc8963d176321cbad9060001b616689565b600094505050505061935a565b6192ef7fdc74aa70dc728613bd4fba5c7dbb01ad19b333de38fff1d6a30964f6a15dad2d60001b616689565b61931b7f8ee0b7bcb8b01ed68e9d28dfcb51b045693e5a295147feaf42234555969ce3c660001b616689565b6193477ff3bbf7f40551837b59192f9130c07b30b7b2f4f4d0e2fd20cd366e3321055a4860001b616689565b8082619353919061cc84565b9450505050505b949350505050565b60006193907f46a83deb927fcb9e72d8dbf5ddc100b8ea5217a9f17b5181fd0c18eef1ccce4360001b616689565b6193bc7f8bcae6b87cf9332345c5a8e1ee0d5b95b0753b8dc1005a9004953fbdb321c9e760001b616689565b6193e87f093995ad57bd3af2ce398cc4579fc35b3fde68ef3a62afba33cc983ff9c6237e60001b616689565b60006002905061941a7fca8e62d7611c977acedd19a28a01623c50a5262f6bd89458cb7960621deba26960001b616689565b6194467f53d301a8dacf364d4b1ea214e71933f2ed816f1465405dbea6863db44856748760001b616689565b60006194568585606001516194d7565b82619461919061cc0e565b905061948f7f5e4584e4a3f63999fbebd1baabd1254233352787679a3ad461188bd0a8c8d19260001b616689565b6194bb7f7388171399841473cf8c271e5542d6bcc3e8bd87c3889439b1e7dc26625cf0ea60001b616689565b809250505092915050565b60006194d130619583565b15905090565b60006195057fc971adab28968779c7ab0dc657047fff590f0101837511a61aaa0a7acabd495f60001b616689565b6195317fc6e43f99ef921b8319fcd707fb1328d54544b9ee4556cdab1f6d7790a499c81460001b616689565b61955d7f5705e9f765fa6001e5a3765a0e7c90bd98b40410f263afa238cb084c122f3af860001b616689565b620fa000828460600151619571919061cc0e565b61957b919061cbdd565b905092915050565b600080823b905060008111915050919050565b6040518060a0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff168152602001600081526020016000151581525090565b60006195fa6195f58461c89a565b61c875565b9050808382526020820190508285602086028201111561961957600080fd5b60005b85811015619649578161962f8882619a46565b84526020840193506020830192505060018101905061961c565b5050509392505050565b60006196666196618461c8c6565b61c875565b9050808382526020820190508285602086028201111561968557600080fd5b60005b858110156196cf57813567ffffffffffffffff8111156196a757600080fd5b8086016196b48982619bc0565b85526020850194506020840193505050600181019050619688565b5050509392505050565b60006196ec6196e78461c8c6565b61c875565b9050808382526020820190508285602086028201111561970b57600080fd5b60005b8581101561975557815167ffffffffffffffff81111561972d57600080fd5b80860161973a8982619bea565b8552602085019450602084019350505060018101905061970e565b5050509392505050565b600061977261976d8461c8f2565b61c875565b9050808382526020820190508285602086028201111561979157600080fd5b60005b858110156197db57813567ffffffffffffffff8111156197b357600080fd5b8086016197c0898261a067565b85526020850194506020840193505050600181019050619794565b5050509392505050565b60006197f86197f38461c91e565b61c875565b9050808382526020820190508285602086028201111561981757600080fd5b60005b8581101561986157813567ffffffffffffffff81111561983957600080fd5b808601619846898261a0df565b8552602085019450602084019350505060018101905061981a565b5050509392505050565b600061987e6198798461c94a565b61c875565b9050808382526020820190508285602086028201111561989d57600080fd5b60005b858110156198e757813567ffffffffffffffff8111156198bf57600080fd5b8086016198cc898261a4d3565b855260208501945060208401935050506001810190506198a0565b5050509392505050565b60006199046198ff8461c94a565b61c875565b9050808382526020820190508285602086028201111561992357600080fd5b60005b8581101561996d57815167ffffffffffffffff81111561994557600080fd5b808601619952898261a573565b85526020850194506020840193505050600181019050619926565b5050509392505050565b600061998a6199858461c976565b61c875565b9050828152602081018484840111156199a257600080fd5b6199ad84828561ce89565b509392505050565b60006199c86199c38461c976565b61c875565b9050828152602081018484840111156199e057600080fd5b6199eb84828561ce98565b509392505050565b6000619a06619a018461c9a7565b61c875565b905082815260208101848484011115619a1e57600080fd5b619a2984828561ce98565b509392505050565b600081359050619a408161d3d9565b92915050565b600081519050619a558161d3d9565b92915050565b600082601f830112619a6c57600080fd5b8151619a7c8482602086016195e7565b91505092915050565b600082601f830112619a9657600080fd5b8135619aa6848260208601619653565b91505092915050565b600082601f830112619ac057600080fd5b8151619ad08482602086016196d9565b91505092915050565b600082601f830112619aea57600080fd5b8135619afa84826020860161975f565b91505092915050565b600082601f830112619b1457600080fd5b8135619b248482602086016197e5565b91505092915050565b600082601f830112619b3e57600080fd5b8135619b4e84826020860161986b565b91505092915050565b600082601f830112619b6857600080fd5b8151619b788482602086016198f1565b91505092915050565b600081359050619b908161d3f0565b92915050565b600081519050619ba58161d3f0565b92915050565b600081359050619bba8161d407565b92915050565b600082601f830112619bd157600080fd5b8135619be1848260208601619977565b91505092915050565b600082601f830112619bfb57600080fd5b8151619c0b8482602086016199b5565b91505092915050565b600081359050619c238161d41e565b92915050565b600081359050619c388161d435565b92915050565b600081359050619c4d8161d44c565b92915050565b600081359050619c628161d463565b92915050565b600081359050619c778161d47a565b92915050565b600081359050619c8c8161d491565b92915050565b600081359050619ca18161d4a8565b92915050565b600081519050619cb68161d4a8565b92915050565b600081519050619ccb8161d4b8565b92915050565b600082601f830112619ce257600080fd5b8151619cf28482602086016199f3565b91505092915050565b60006103208284031215619d0e57600080fd5b619d196102e061c875565b9050600082015167ffffffffffffffff811115619d3557600080fd5b619d4184828501619bea565b6000830152506020619d5584828501619a46565b602083015250604082015167ffffffffffffffff811115619d7557600080fd5b619d8184828501619bea565b6040830152506060619d958482850161aba6565b6060830152506080619da98482850161aba6565b60808301525060a0619dbd8482850161aba6565b60a08301525060c0619dd18482850161aba6565b60c08301525060e0619de58482850161aba6565b60e083015250610100619dfa8482850161ab7c565b61010083015250610120619e108482850161aba6565b61012083015250610140619e268482850161aba6565b6101408301525061016082015167ffffffffffffffff811115619e4857600080fd5b619e5484828501619bea565b61016083015250610180619e6a8482850161aba6565b610180830152506101a0619e808482850161ab7c565b6101a0830152506101c0619e9684828501619b96565b6101c0830152506101e0619eac84828501619cbc565b6101e083015250610200619ec28482850161aba6565b6102008301525061022082015167ffffffffffffffff811115619ee457600080fd5b619ef084828501619a5b565b6102208301525061024082015167ffffffffffffffff811115619f1257600080fd5b619f1e84828501619a5b565b6102408301525061026082015167ffffffffffffffff811115619f4057600080fd5b619f4c84828501619bea565b61026083015250610280619f6284828501619ca7565b610280830152506102a0619f7884828501619b96565b6102a0830152506102c0619f8e8482850161a2d3565b6102c08301525092915050565b600060c08284031215619fad57600080fd5b619fb760c061c875565b9050600082013567ffffffffffffffff811115619fd357600080fd5b619fdf84828501619bc0565b600083015250602082013567ffffffffffffffff811115619fff57600080fd5b61a00b8482850161a3cb565b602083015250604061a01f8482850161ab67565b604083015250606061a03384828501619a31565b606083015250608061a0478482850161ab91565b60808301525060a061a05b8482850161ab91565b60a08301525092915050565b60006060828403121561a07957600080fd5b61a083606061c875565b9050600061a0938482850161ab91565b600083015250602061a0a78482850161ab91565b602083015250604082013567ffffffffffffffff81111561a0c757600080fd5b61a0d384828501619bc0565b60408301525092915050565b60006040828403121561a0f157600080fd5b61a0fb604061c875565b9050600061a10b8482850161ab91565b600083015250602082013567ffffffffffffffff81111561a12b57600080fd5b61a13784828501619ad9565b60208301525092915050565b600060e0828403121561a15557600080fd5b61a15f60e061c875565b9050600061a16f8482850161ab91565b600083015250602061a1838482850161ab91565b602083015250604061a1978482850161ab91565b604083015250606061a1ab8482850161ab91565b606083015250608061a1bf8482850161ab91565b60808301525060a061a1d384828501619a31565b60a08301525060c082013567ffffffffffffffff81111561a1f357600080fd5b61a1ff84828501619bc0565b60c08301525092915050565b600060e0828403121561a21d57600080fd5b61a22760e061c875565b9050600061a2378482850161aba6565b600083015250602061a24b8482850161aba6565b602083015250604061a25f8482850161aba6565b604083015250606061a2738482850161aba6565b606083015250608061a2878482850161aba6565b60808301525060a061a29b84828501619a46565b60a08301525060c082015167ffffffffffffffff81111561a2bb57600080fd5b61a2c784828501619bea565b60c08301525092915050565b60006060828403121561a2e557600080fd5b61a2ef606061c875565b9050600061a2ff8482850161aba6565b600083015250602061a3138482850161aba6565b602083015250604061a3278482850161aba6565b60408301525092915050565b60006060828403121561a34557600080fd5b61a34f606061c875565b9050600061a35f84828501619a46565b600083015250602061a3738482850161ab7c565b602083015250604061a3878482850161aba6565b60408301525092915050565b60006020828403121561a3a557600080fd5b61a3af602061c875565b9050600061a3bf8482850161ab91565b60008301525092915050565b60006080828403121561a3dd57600080fd5b61a3e7608061c875565b9050600082013567ffffffffffffffff81111561a40357600080fd5b61a40f84828501619bc0565b600083015250602061a4238482850161ab91565b602083015250604082013567ffffffffffffffff81111561a44357600080fd5b61a44f84828501619a85565b604083015250606082013567ffffffffffffffff81111561a46f57600080fd5b61a47b84828501619b03565b60608301525092915050565b60006040828403121561a49957600080fd5b61a4a3604061c875565b9050600061a4b38482850161ab91565b600083015250602061a4c78482850161ab91565b60208301525092915050565b600060a0828403121561a4e557600080fd5b61a4ef60a061c875565b9050600082013567ffffffffffffffff81111561a50b57600080fd5b61a51784828501619bc0565b600083015250602061a52b84828501619a31565b602083015250604061a53f8482850161ab91565b604083015250606061a5538482850161ab67565b606083015250608061a56784828501619b81565b60808301525092915050565b600060a0828403121561a58557600080fd5b61a58f60a061c875565b9050600082015167ffffffffffffffff81111561a5ab57600080fd5b61a5b784828501619bea565b600083015250602061a5cb84828501619a46565b602083015250604061a5df8482850161aba6565b604083015250606061a5f38482850161ab7c565b606083015250608061a60784828501619b96565b60808301525092915050565b6000610180828403121561a62657600080fd5b61a63161018061c875565b9050600061a64184828501619a31565b600083015250602061a6558482850161ab91565b602083015250604061a6698482850161ab91565b604083015250606061a67d8482850161ab91565b606083015250608061a69184828501619c92565b60808301525060a061a6a58482850161ab67565b60a08301525060c061a6b98482850161ab67565b60c08301525060e061a6cd8482850161ab91565b60e08301525061010061a6e28482850161ab91565b6101008301525061012061a6f88482850161ab91565b6101208301525061014061a70e84828501619b81565b6101408301525061016082013567ffffffffffffffff81111561a73057600080fd5b61a73c84828501619a85565b6101608301525092915050565b6000610180828403121561a75c57600080fd5b61a76761018061c875565b9050600061a77784828501619a46565b600083015250602061a78b8482850161aba6565b602083015250604061a79f8482850161aba6565b604083015250606061a7b38482850161aba6565b606083015250608061a7c784828501619ca7565b60808301525060a061a7db8482850161ab7c565b60a08301525060c061a7ef8482850161ab7c565b60c08301525060e061a8038482850161aba6565b60e08301525061010061a8188482850161aba6565b6101008301525061012061a82e8482850161aba6565b6101208301525061014061a84484828501619b96565b6101408301525061016082015167ffffffffffffffff81111561a86657600080fd5b61a87284828501619aaf565b6101608301525092915050565b60006080828403121561a89157600080fd5b61a89b608061c875565b9050600061a8ab84828501619a31565b600083015250602061a8bf8482850161ab91565b602083015250604061a8d38482850161ab91565b604083015250606082013567ffffffffffffffff81111561a8f357600080fd5b61a8ff8482850161a3cb565b60608301525092915050565b60006040828403121561a91d57600080fd5b61a927604061c875565b9050600061a93784828501619a31565b600083015250602061a94b8482850161ab91565b60208301525092915050565b6000610160828403121561a96a57600080fd5b61a97561016061c875565b9050600061a9858482850161ab91565b600083015250602061a9998482850161ab91565b602083015250604061a9ad8482850161ab91565b604083015250606061a9c18482850161ab91565b606083015250608061a9d58482850161ab91565b60808301525060a061a9e98482850161ab91565b60a08301525060c061a9fd8482850161ab91565b60c08301525060e061aa118482850161ab91565b60e08301525061010061aa268482850161ab91565b6101008301525061012061aa3c8482850161ab91565b6101208301525061014061aa528482850161ab91565b6101408301525092915050565b6000610160828403121561aa7257600080fd5b61aa7d61016061c875565b9050600061aa8d8482850161aba6565b600083015250602061aaa18482850161aba6565b602083015250604061aab58482850161aba6565b604083015250606061aac98482850161aba6565b606083015250608061aadd8482850161aba6565b60808301525060a061aaf18482850161aba6565b60a08301525060c061ab058482850161aba6565b60c08301525060e061ab198482850161aba6565b60e08301525061010061ab2e8482850161aba6565b6101008301525061012061ab448482850161aba6565b6101208301525061014061ab5a8482850161aba6565b6101408301525092915050565b60008135905061ab768161d4c8565b92915050565b60008151905061ab8b8161d4c8565b92915050565b60008135905061aba08161d4df565b92915050565b60008151905061abb58161d4df565b92915050565b60008060006060848603121561abd057600080fd5b600061abde86828701619a31565b935050602061abef8682870161ab91565b925050604061ac008682870161ab67565b9150509250925092565b60006020828403121561ac1c57600080fd5b600082015167ffffffffffffffff81111561ac3657600080fd5b61ac4284828501619b57565b91505092915050565b60006020828403121561ac5d57600080fd5b600061ac6b84828501619b96565b91505092915050565b60006020828403121561ac8657600080fd5b600061ac9484828501619bab565b91505092915050565b60006020828403121561acaf57600080fd5b600082013567ffffffffffffffff81111561acc957600080fd5b61acd584828501619bc0565b91505092915050565b6000806040838503121561acf157600080fd5b600083013567ffffffffffffffff81111561ad0b57600080fd5b61ad1785828601619bc0565b925050602083013567ffffffffffffffff81111561ad3457600080fd5b61ad4085828601619b2d565b9150509250929050565b6000806060838503121561ad5d57600080fd5b600083013567ffffffffffffffff81111561ad7757600080fd5b61ad8385828601619bc0565b925050602061ad948582860161a487565b9150509250929050565b600080600080600080600060e0888a03121561adb957600080fd5b600061adc78a828b01619c14565b975050602061add88a828b01619c29565b965050604061ade98a828b01619c3e565b955050606061adfa8a828b01619c53565b945050608061ae0b8a828b01619c68565b93505060a061ae1c8a828b0161a393565b92505060c061ae2d8a828b01619c7d565b91505092959891949750929550565b60006020828403121561ae4e57600080fd5b600082015167ffffffffffffffff81111561ae6857600080fd5b61ae7484828501619cd1565b91505092915050565b60006020828403121561ae8f57600080fd5b600082015167ffffffffffffffff81111561aea957600080fd5b61aeb584828501619cfb565b91505092915050565b60006020828403121561aed057600080fd5b600082013567ffffffffffffffff81111561aeea57600080fd5b61aef684828501619f9b565b91505092915050565b60006020828403121561af1157600080fd5b600082015167ffffffffffffffff81111561af2b57600080fd5b61af378482850161a20b565b91505092915050565b60006060828403121561af5257600080fd5b600061af608482850161a333565b91505092915050565b60006020828403121561af7b57600080fd5b600082015167ffffffffffffffff81111561af9557600080fd5b61afa18482850161a749565b91505092915050565b60008060006101a0848603121561afc057600080fd5b600084013567ffffffffffffffff81111561afda57600080fd5b61afe68682870161a613565b935050602084013567ffffffffffffffff81111561b00357600080fd5b61b00f8682870161a143565b925050604061b0208682870161a957565b9150509250925092565b6000806000806101c0858703121561b04157600080fd5b600085013567ffffffffffffffff81111561b05b57600080fd5b61b0678782880161a613565b945050602085013567ffffffffffffffff81111561b08457600080fd5b61b0908782880161a143565b935050604061b0a18782880161a957565b9250506101a061b0b38782880161ab91565b91505092959194509250565b60006020828403121561b0d157600080fd5b600082013567ffffffffffffffff81111561b0eb57600080fd5b61b0f78482850161a87f565b91505092915050565b60006040828403121561b11257600080fd5b600061b1208482850161a90b565b91505092915050565b6000610160828403121561b13c57600080fd5b600061b14a8482850161aa5f565b91505092915050565b600061b15f838361b1bb565b60208301905092915050565b600061b177838361b429565b905092915050565b600061b18b838361bad7565b905092915050565b600061b19f838361bb27565b905092915050565b600061b1b3838361bd24565b905092915050565b61b1c48161ccb8565b82525050565b61b1d38161ccb8565b82525050565b600061b1e48261ca28565b61b1ee818561cab6565b935061b1f98361c9d8565b8060005b8381101561b22a57815161b211888261b153565b975061b21c8361ca75565b92505060018101905061b1fd565b5085935050505092915050565b600061b2428261ca33565b61b24c818561cac7565b93508360208202850161b25e8561c9e8565b8060005b8581101561b29a578484038952815161b27b858261b16b565b945061b2868361ca82565b925060208a0199505060018101905061b262565b50829750879550505050505092915050565b600061b2b78261ca3e565b61b2c1818561cad8565b93508360208202850161b2d38561c9f8565b8060005b8581101561b30f578484038952815161b2f0858261b17f565b945061b2fb8361ca8f565b925060208a0199505060018101905061b2d7565b50829750879550505050505092915050565b600061b32c8261ca49565b61b336818561cae9565b93508360208202850161b3488561ca08565b8060005b8581101561b384578484038952815161b365858261b193565b945061b3708361ca9c565b925060208a0199505060018101905061b34c565b50829750879550505050505092915050565b600061b3a18261ca54565b61b3ab818561cafa565b93508360208202850161b3bd8561ca18565b8060005b8581101561b3f9578484038952815161b3da858261b1a7565b945061b3e58361caa9565b925060208a0199505060018101905061b3c1565b50829750879550505050505092915050565b61b4148161ccca565b82525050565b61b4238161ccca565b82525050565b600061b4348261ca5f565b61b43e818561cb0b565b935061b44e81856020860161ce98565b61b4578161d032565b840191505092915050565b600061b46d8261ca5f565b61b477818561cb1c565b935061b48781856020860161ce98565b61b4908161d032565b840191505092915050565b61b4a48161cdc3565b82525050565b61b4b38161cde7565b82525050565b61b4c28161ce0b565b82525050565b61b4d18161ce2f565b82525050565b61b4e08161ce53565b82525050565b61b4ef8161ce65565b82525050565b61b4fe8161ce65565b82525050565b61b50d8161ce77565b82525050565b600061b51e8261ca6a565b61b528818561cb2d565b935061b53881856020860161ce98565b61b5418161d032565b840191505092915050565b600061b5578261ca6a565b61b561818561cb3e565b935061b57181856020860161ce98565b80840191505092915050565b600061b58a60178361cb2d565b915061b5958261d043565b602082019050919050565b600061b5ad60158361cb2d565b915061b5b88261d06c565b602082019050919050565b600061b5d0601c8361cb2d565b915061b5db8261d095565b602082019050919050565b600061b5f360228361cb2d565b915061b5fe8261d0be565b604082019050919050565b600061b616601b8361cb2d565b915061b6218261d10d565b602082019050919050565b600061b63960108361cb2d565b915061b6448261d136565b602082019050919050565b600061b65c60198361cb2d565b915061b6678261d15f565b602082019050919050565b600061b67f601b8361cb2d565b915061b68a8261d188565b602082019050919050565b600061b6a2601b8361cb2d565b915061b6ad8261d1b1565b602082019050919050565b600061b6c5602e8361cb2d565b915061b6d08261d1da565b604082019050919050565b600061b6e860148361cb2d565b915061b6f38261d229565b602082019050919050565b600061b70b60148361cb2d565b915061b7168261d252565b602082019050919050565b600061b72e60168361cb2d565b915061b7398261d27b565b602082019050919050565b600061b751600b8361cb2d565b915061b75c8261d2a4565b602082019050919050565b600061b77460008361cb3e565b915061b77f8261d2cd565b600082019050919050565b600061b79760168361cb2d565b915061b7a28261d2d0565b602082019050919050565b600061b7ba60198361cb2d565b915061b7c58261d2f9565b602082019050919050565b600061b7dd60158361cb2d565b915061b7e88261d322565b602082019050919050565b600061b80060098361cb2d565b915061b80b8261d34b565b602082019050919050565b600061b82360108361cb2d565b915061b82e8261d374565b602082019050919050565b600061032083016000830151848203600086015261b857828261b429565b915050602083015161b86c602086018261b1bb565b506040830151848203604086015261b884828261b429565b915050606083015161b899606086018261c0a5565b50608083015161b8ac608086018261c0a5565b5060a083015161b8bf60a086018261c0a5565b5060c083015161b8d260c086018261c0a5565b5060e083015161b8e560e086018261c0a5565b5061010083015161b8fa61010086018261c087565b5061012083015161b90f61012086018261c0a5565b5061014083015161b92461014086018261c0a5565b5061016083015184820361016086015261b93e828261b429565b91505061018083015161b95561018086018261c0a5565b506101a083015161b96a6101a086018261c087565b506101c083015161b97f6101c086018261b40b565b506101e083015161b9946101e086018261b504565b5061020083015161b9a961020086018261c0a5565b5061022083015184820361022086015261b9c3828261b1d9565b91505061024083015184820361024086015261b9df828261b1d9565b91505061026083015184820361026086015261b9fb828261b429565b91505061028083015161ba1261028086018261b4e6565b506102a083015161ba276102a086018261b40b565b506102c083015161ba3c6102c086018261bc00565b508091505092915050565b600060c083016000830151848203600086015261ba64828261b429565b9150506020830151848203602086015261ba7e828261bc84565b915050604083015161ba93604086018261c087565b50606083015161baa6606086018261b1bb565b50608083015161bab9608086018261c0a5565b5060a083015161bacc60a086018261c0a5565b508091505092915050565b600060608301600083015161baef600086018261c0a5565b50602083015161bb02602086018261c0a5565b506040830151848203604086015261bb1a828261b429565b9150508091505092915050565b600060408301600083015161bb3f600086018261c0a5565b506020830151848203602086015261bb57828261b2ac565b9150508091505092915050565b600060e08301600083015161bb7c600086018261c0a5565b50602083015161bb8f602086018261c0a5565b50604083015161bba2604086018261c0a5565b50606083015161bbb5606086018261c0a5565b50608083015161bbc8608086018261c0a5565b5060a083015161bbdb60a086018261b1bb565b5060c083015184820360c086015261bbf3828261b429565b9150508091505092915050565b60608201600082015161bc16600085018261c0a5565b50602082015161bc29602085018261c0a5565b50604082015161bc3c604085018261c0a5565b50505050565b60608201600082015161bc58600085018261b1bb565b50602082015161bc6b602085018261c087565b50604082015161bc7e604085018261c0a5565b50505050565b6000608083016000830151848203600086015261bca1828261b429565b915050602083015161bcb6602086018261c0a5565b506040830151848203604086015261bcce828261b237565b9150506060830151848203606086015261bce8828261b321565b9150508091505092915050565b60408201600082015161bd0b600085018261c0a5565b50602082015161bd1e602085018261c0a5565b50505050565b600060a083016000830151848203600086015261bd41828261b429565b915050602083015161bd56602086018261b1bb565b50604083015161bd69604086018261c0a5565b50606083015161bd7c606086018261c087565b50608083015161bd8f608086018261b40b565b508091505092915050565b600060a083016000830151848203600086015261bdb7828261b429565b915050602083015161bdcc602086018261b1bb565b50604083015161bddf604086018261c0a5565b50606083015161bdf2606086018261c087565b50608083015161be05608086018261b40b565b508091505092915050565b60006101808301600083015161be29600086018261b1bb565b50602083015161be3c602086018261c0a5565b50604083015161be4f604086018261c0a5565b50606083015161be62606086018261c0a5565b50608083015161be75608086018261b4e6565b5060a083015161be8860a086018261c087565b5060c083015161be9b60c086018261c087565b5060e083015161beae60e086018261c0a5565b5061010083015161bec361010086018261c0a5565b5061012083015161bed861012086018261c0a5565b5061014083015161beed61014086018261b40b565b5061016083015184820361016086015261bf07828261b237565b9150508091505092915050565b600060808301600083015161bf2c600086018261b1bb565b50602083015161bf3f602086018261c0a5565b50604083015161bf52604086018261c0a5565b506060830151848203606086015261bf6a828261bc84565b9150508091505092915050565b60408201600082015161bf8d600085018261b1bb565b50602082015161bfa0602085018261c0a5565b50505050565b6101608201600082015161bfbd600085018261c0a5565b50602082015161bfd0602085018261c0a5565b50604082015161bfe3604085018261c0a5565b50606082015161bff6606085018261c0a5565b50608082015161c009608085018261c0a5565b5060a082015161c01c60a085018261c0a5565b5060c082015161c02f60c085018261c0a5565b5060e082015161c04260e085018261c0a5565b5061010082015161c05761010085018261c0a5565b5061012082015161c06c61012085018261c0a5565b5061014082015161c08161014085018261c0a5565b50505050565b61c0908161cda5565b82525050565b61c09f8161cda5565b82525050565b61c0ae8161cdaf565b82525050565b61c0bd8161cdaf565b82525050565b600061c0cf828461b54c565b915081905092915050565b600061c0e58261b767565b9150819050919050565b600060208201905061c104600083018461b1ca565b92915050565b600060408201905061c11f600083018561b1ca565b61c12c602083018461c096565b9392505050565b6000602082019050818103600083015261c14d818461b396565b905092915050565b600060208201905061c16a600083018461b41a565b92915050565b6000602082019050818103600083015261c18a818461b462565b905092915050565b6000604082019050818103600083015261c1ac818561b462565b9050818103602083015261c1c0818461b396565b90509392505050565b6000606082019050818103600083015261c1e3818561b462565b905061c1f2602083018461bcf5565b9392505050565b600060408201905061c20e600083018561b4aa565b818103602083015261c220818461b462565b90509392505050565b60006102208201905061c23f600083018a61b4aa565b61c24c602083018961b49b565b818103604083015261c25e818861b839565b9050818103606083015261c272818761bb64565b9050818103608083015261c286818661bd9a565b905081810360a083015261c29a818561b396565b905061c2a960c083018461bfa6565b98975050505050505050565b600060608201905061c2ca600083018661b4b9565b818103602083015261c2dc818561ba47565b9050818103604083015261c2f0818461b839565b9050949350505050565b600060808201905061c30f600083018761b4b9565b61c31c602083018661c0b4565b818103604083015261c32e818561bf14565b9050818103606083015261c342818461be10565b905095945050505050565b600060a08201905061c362600083018761b4c8565b61c36f602083018661b49b565b818103604083015261c381818561b839565b905061c390606083018461bf77565b95945050505050565b600060808201905061c3ae600083018761b4d7565b61c3bb602083018661c096565b818103604083015261c3cd818561b462565b905061c3dc606083018461b1ca565b95945050505050565b600060208201905061c3fa600083018461b4f5565b92915050565b6000602082019050818103600083015261c41a818461b513565b905092915050565b6000604082019050818103600083015261c43b8161b695565b9050818103602083015261c44e8161b57d565b9050919050565b6000604082019050818103600083015261c46e8161b695565b9050818103602083015261c4818161b672565b9050919050565b6000602082019050818103600083015261c4a18161b6b8565b9050919050565b6000604082019050818103600083015261c4c18161b744565b9050818103602083015261c4d5818461b513565b905092915050565b6000604082019050818103600083015261c4f68161b744565b9050818103602083015261c5098161b5c3565b9050919050565b6000604082019050818103600083015261c5298161b744565b9050818103602083015261c53c8161b5e6565b9050919050565b6000604082019050818103600083015261c55c8161b744565b9050818103602083015261c56f8161b64f565b9050919050565b6000604082019050818103600083015261c58f8161b744565b9050818103602083015261c5a28161b7d0565b9050919050565b6000604082019050818103600083015261c5c28161b7f3565b9050818103602083015261c5d6818461b513565b905092915050565b6000604082019050818103600083015261c5f78161b7f3565b9050818103602083015261c60a8161b5a0565b9050919050565b6000604082019050818103600083015261c62a8161b7f3565b9050818103602083015261c63d8161b609565b9050919050565b6000604082019050818103600083015261c65d8161b7f3565b9050818103602083015261c6708161b62c565b9050919050565b6000604082019050818103600083015261c6908161b7f3565b9050818103602083015261c6a38161b6db565b9050919050565b6000604082019050818103600083015261c6c38161b7f3565b9050818103602083015261c6d68161b6fe565b9050919050565b6000604082019050818103600083015261c6f68161b7f3565b9050818103602083015261c7098161b721565b9050919050565b6000604082019050818103600083015261c7298161b7f3565b9050818103602083015261c73c8161b78a565b9050919050565b6000604082019050818103600083015261c75c8161b7f3565b9050818103602083015261c76f8161b7ad565b9050919050565b6000604082019050818103600083015261c78f8161b7f3565b9050818103602083015261c7a28161b816565b9050919050565b6000602082019050818103600083015261c7c3818461bb64565b905092915050565b600060608201905061c7e0600083018461bc42565b92915050565b6000602082019050818103600083015261c800818461be10565b905092915050565b6000604082019050818103600083015261c822818561be10565b9050818103602083015261c836818461b839565b90509392505050565b600060408201905061c854600083018461bf77565b92915050565b600060208201905061c86f600083018461c096565b92915050565b600061c87f61c890565b905061c88b828261cecb565b919050565b6000604051905090565b600067ffffffffffffffff82111561c8b55761c8b461d003565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561c8e15761c8e061d003565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561c90d5761c90c61d003565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561c9395761c93861d003565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561c9655761c96461d003565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561c9915761c99061d003565b5b61c99a8261d032565b9050602081019050919050565b600067ffffffffffffffff82111561c9c25761c9c161d003565b5b61c9cb8261d032565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600061cb548261cda5565b915061cb5f8361cda5565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561cb945761cb9361cf76565b5b828201905092915050565b600061cbaa8261cdaf565b915061cbb58361cdaf565b92508267ffffffffffffffff0382111561cbd25761cbd161cf76565b5b828201905092915050565b600061cbe88261cdaf565b915061cbf38361cdaf565b92508261cc035761cc0261cfa5565b5b828204905092915050565b600061cc198261cdaf565b915061cc248361cdaf565b92508167ffffffffffffffff048311821515161561cc455761cc4461cf76565b5b828202905092915050565b600061cc5b8261cda5565b915061cc668361cda5565b92508282101561cc795761cc7861cf76565b5b828203905092915050565b600061cc8f8261cdaf565b915061cc9a8361cdaf565b92508282101561ccad5761ccac61cf76565b5b828203905092915050565b600061ccc38261cd85565b9050919050565b60008115159050919050565b6000819050919050565b600061cceb8261ccb8565b9050919050565b600061ccfd8261ccb8565b9050919050565b600061cd0f8261ccb8565b9050919050565b600061cd218261ccb8565b9050919050565b600061cd338261ccb8565b9050919050565b600061cd458261ccb8565b9050919050565b600081905061cd5a8261d39d565b919050565b600081905061cd6d8261d3b1565b919050565b600081905061cd808261d3c5565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600061cdce8261cdd5565b9050919050565b600061cde08261cd85565b9050919050565b600061cdf28261cdf9565b9050919050565b600061ce048261cd85565b9050919050565b600061ce168261ce1d565b9050919050565b600061ce288261cd85565b9050919050565b600061ce3a8261ce41565b9050919050565b600061ce4c8261cd85565b9050919050565b600061ce5e8261cd4c565b9050919050565b600061ce708261cd5f565b9050919050565b600061ce828261cd72565b9050919050565b82818337600083830152505050565b60005b8381101561ceb657808201518184015260208101905061ce9b565b8381111561cec5576000848401525b50505050565b61ced48261d032565b810181811067ffffffffffffffff8211171561cef35761cef261d003565b5b80604052505050565b600061cf078261cda5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561cf3a5761cf3961cf76565b5b600182019050919050565b600061cf508261cdaf565b915067ffffffffffffffff82141561cf6b5761cf6a61cf76565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f4e6f64655365727669636554696d654e6f744d61746368000000000000000000600082015250565b7f46696c6550726f76654e6f7446696c654f776e65720000000000000000000000600082015250565b7f536563746f7250726f766550726f66697453706c69744661696c656400000000600082015250565b7f536563746f7250726f76654368616c6c656e67654865696768744e6f744d617460008201527f6368000000000000000000000000000000000000000000000000000000000000602082015250565b7f46696c6550726f7665536563746f72547970654e6f744d617463680000000000600082015250565b7f46696c6550726f76654578706972656400000000000000000000000000000000600082015250565b7f536563746f7250726f7665536563746f724e6f74457869737400000000000000600082015250565b7f4e6f6465536563746f7250726f766564496e54696d654572726f720000000000600082015250565b7f436865636b4e6f6465536563746f7250726f766564496e54696d650000000000600082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f46696c6550726f766554696d6573457863656564000000000000000000000000600082015250565b7f46696c6550726f7665436865636b4661696c6564000000000000000000000000600082015250565b7f46696c6550726f7665436f70794e756d45786365656400000000000000000000600082015250565b7f536563746f7250726f7665000000000000000000000000000000000000000000600082015250565b50565b7f46696c6550726f7665416c726561647950726f76656400000000000000000000600082015250565b7f46696c6550726f76654e6f6465566f6c4e6f74456e6f75676800000000000000600082015250565b7f536563746f7250726f76654e6f74457870697265640000000000000000000000600082015250565b7f46696c6550726f76650000000000000000000000000000000000000000000000600082015250565b7f46696c6550726f76654e6f744e6f646500000000000000000000000000000000600082015250565b600b811061d3ae5761d3ad61cfd4565b5b50565b6003811061d3c25761d3c161cfd4565b5b50565b6002811061d3d65761d3d561cfd4565b5b50565b61d3e28161ccb8565b811461d3ed57600080fd5b50565b61d3f98161ccca565b811461d40457600080fd5b50565b61d4108161ccd6565b811461d41b57600080fd5b50565b61d4278161cce0565b811461d43257600080fd5b50565b61d43e8161ccf2565b811461d44957600080fd5b50565b61d4558161cd04565b811461d46057600080fd5b50565b61d46c8161cd16565b811461d47757600080fd5b50565b61d4838161cd28565b811461d48e57600080fd5b50565b61d49a8161cd3a565b811461d4a557600080fd5b50565b6003811061d4b557600080fd5b50565b6002811061d4c557600080fd5b50565b61d4d18161cda5565b811461d4dc57600080fd5b50565b61d4e88161cdaf565b811461d4f357600080fd5b5056fea2646970667358221220a7301743ba96535e8d754a2ed6a584ff35039971079d165e90d3386fc9f3c97a64736f6c63430008040033",
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

// GetProveDetailList is a free data retrieval call binding the contract method 0x977fdfe2.
//
// Solidity: function GetProveDetailList(bytes fileHash) view returns((bytes,address,uint64,uint256,bool)[])
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
// Solidity: function GetProveDetailList(bytes fileHash) view returns((bytes,address,uint64,uint256,bool)[])
func (_Store *StoreSession) GetProveDetailList(fileHash []byte) ([]ProveDetail, error) {
	return _Store.Contract.GetProveDetailList(&_Store.CallOpts, fileHash)
}

// GetProveDetailList is a free data retrieval call binding the contract method 0x977fdfe2.
//
// Solidity: function GetProveDetailList(bytes fileHash) view returns((bytes,address,uint64,uint256,bool)[])
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

// FileProve is a paid mutator transaction binding the contract method 0x8d4c13a9.
//
// Solidity: function FileProve((bytes,(bytes,uint64,bytes[],(uint64,(uint64,uint64,bytes)[])[]),uint256,address,uint64,uint64) fileProve) payable returns()
func (_Store *StoreTransactor) FileProve(opts *bind.TransactOpts, fileProve FileProveParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "FileProve", fileProve)
}

// FileProve is a paid mutator transaction binding the contract method 0x8d4c13a9.
//
// Solidity: function FileProve((bytes,(bytes,uint64,bytes[],(uint64,(uint64,uint64,bytes)[])[]),uint256,address,uint64,uint64) fileProve) payable returns()
func (_Store *StoreSession) FileProve(fileProve FileProveParams) (*types.Transaction, error) {
	return _Store.Contract.FileProve(&_Store.TransactOpts, fileProve)
}

// FileProve is a paid mutator transaction binding the contract method 0x8d4c13a9.
//
// Solidity: function FileProve((bytes,(bytes,uint64,bytes[],(uint64,(uint64,uint64,bytes)[])[]),uint256,address,uint64,uint64) fileProve) payable returns()
func (_Store *StoreTransactorSession) FileProve(fileProve FileProveParams) (*types.Transaction, error) {
	return _Store.Contract.FileProve(&_Store.TransactOpts, fileProve)
}

// SectorProve is a paid mutator transaction binding the contract method 0xbe5ac5ea.
//
// Solidity: function SectorProve((address,uint64,uint64,(bytes,uint64,bytes[],(uint64,(uint64,uint64,bytes)[])[])) sectorProve) payable returns()
func (_Store *StoreTransactor) SectorProve(opts *bind.TransactOpts, sectorProve SectorProveParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SectorProve", sectorProve)
}

// SectorProve is a paid mutator transaction binding the contract method 0xbe5ac5ea.
//
// Solidity: function SectorProve((address,uint64,uint64,(bytes,uint64,bytes[],(uint64,(uint64,uint64,bytes)[])[])) sectorProve) payable returns()
func (_Store *StoreSession) SectorProve(sectorProve SectorProveParams) (*types.Transaction, error) {
	return _Store.Contract.SectorProve(&_Store.TransactOpts, sectorProve)
}

// SectorProve is a paid mutator transaction binding the contract method 0xbe5ac5ea.
//
// Solidity: function SectorProve((address,uint64,uint64,(bytes,uint64,bytes[],(uint64,(uint64,uint64,bytes)[])[])) sectorProve) payable returns()
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

// UpdateProveDetailList is a paid mutator transaction binding the contract method 0x181197f7.
//
// Solidity: function UpdateProveDetailList(bytes fileHash, (bytes,address,uint64,uint256,bool)[] details) payable returns()
func (_Store *StoreTransactor) UpdateProveDetailList(opts *bind.TransactOpts, fileHash []byte, details []ProveDetail) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateProveDetailList", fileHash, details)
}

// UpdateProveDetailList is a paid mutator transaction binding the contract method 0x181197f7.
//
// Solidity: function UpdateProveDetailList(bytes fileHash, (bytes,address,uint64,uint256,bool)[] details) payable returns()
func (_Store *StoreSession) UpdateProveDetailList(fileHash []byte, details []ProveDetail) (*types.Transaction, error) {
	return _Store.Contract.UpdateProveDetailList(&_Store.TransactOpts, fileHash, details)
}

// UpdateProveDetailList is a paid mutator transaction binding the contract method 0x181197f7.
//
// Solidity: function UpdateProveDetailList(bytes fileHash, (bytes,address,uint64,uint256,bool)[] details) payable returns()
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

// ProfitSplitForSector is a paid mutator transaction binding the contract method 0x0e3459fd.
//
// Solidity: function profitSplitForSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting) payable returns(bool)
func (_Store *StoreTransactor) ProfitSplitForSector(opts *bind.TransactOpts, sectorInfo SectorInfo, nodeInfo NodeInfo, setting Setting) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "profitSplitForSector", sectorInfo, nodeInfo, setting)
}

// ProfitSplitForSector is a paid mutator transaction binding the contract method 0x0e3459fd.
//
// Solidity: function profitSplitForSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting) payable returns(bool)
func (_Store *StoreSession) ProfitSplitForSector(sectorInfo SectorInfo, nodeInfo NodeInfo, setting Setting) (*types.Transaction, error) {
	return _Store.Contract.ProfitSplitForSector(&_Store.TransactOpts, sectorInfo, nodeInfo, setting)
}

// ProfitSplitForSector is a paid mutator transaction binding the contract method 0x0e3459fd.
//
// Solidity: function profitSplitForSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting) payable returns(bool)
func (_Store *StoreTransactorSession) ProfitSplitForSector(sectorInfo SectorInfo, nodeInfo NodeInfo, setting Setting) (*types.Transaction, error) {
	return _Store.Contract.ProfitSplitForSector(&_Store.TransactOpts, sectorInfo, nodeInfo, setting)
}

// PunishForSector is a paid mutator transaction binding the contract method 0xa0004412.
//
// Solidity: function punishForSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 times) payable returns(string)
func (_Store *StoreTransactor) PunishForSector(opts *bind.TransactOpts, sectorInfo SectorInfo, nodeInfo NodeInfo, setting Setting, times uint64) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "punishForSector", sectorInfo, nodeInfo, setting, times)
}

// PunishForSector is a paid mutator transaction binding the contract method 0xa0004412.
//
// Solidity: function punishForSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 times) payable returns(string)
func (_Store *StoreSession) PunishForSector(sectorInfo SectorInfo, nodeInfo NodeInfo, setting Setting, times uint64) (*types.Transaction, error) {
	return _Store.Contract.PunishForSector(&_Store.TransactOpts, sectorInfo, nodeInfo, setting, times)
}

// PunishForSector is a paid mutator transaction binding the contract method 0xa0004412.
//
// Solidity: function punishForSector((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo, (uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo, (uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64) setting, uint64 times) payable returns(string)
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
