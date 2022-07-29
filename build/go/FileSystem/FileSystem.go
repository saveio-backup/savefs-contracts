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

// Challenge is an auto generated low-level Go binding around an user-defined struct.
type Challenge struct {
	Index uint32
	Rand  uint32
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

// FileProveParams is an auto generated low-level Go binding around an user-defined struct.
type FileProveParams struct {
	FileHash    []byte
	ProveData   []byte
	BlockHeight *big.Int
	NodeWallet  common.Address
	Profit      uint64
	SectorID    uint64
}

// FileReNewInfo is an auto generated low-level Go binding around an user-defined struct.
type FileReNewInfo struct {
	FileHash   []byte
	FromAddr   common.Address
	ReNewTimes uint64
}

// GenChallengeParams is an auto generated low-level Go binding around an user-defined struct.
type GenChallengeParams struct {
	WalletAddr   common.Address
	HashValue    []byte
	FileBlockNum uint64
	ProveNum     uint64
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

// OwnerChange is an auto generated low-level Go binding around an user-defined struct.
type OwnerChange struct {
	FileHash []byte
	CurOwner common.Address
	NewOwner common.Address
}

// PdpVerificationReturns is an auto generated low-level Go binding around an user-defined struct.
type PdpVerificationReturns struct {
	FileIDs     [][]byte
	Tags        [][]byte
	UpdatedChal []Challenge
	Path        []MerklePath
	RootHashes  []byte
	FileInfo    FileInfo
	Success     bool
}

// PlotInfo is an auto generated low-level Go binding around an user-defined struct.
type PlotInfo struct {
	NumberID   uint64
	StartNonce uint64
	Nonces     uint64
}

// PrepareForPdpVerificationParams is an auto generated low-level Go binding around an user-defined struct.
type PrepareForPdpVerificationParams struct {
	SectorInfo SectorInfo
	Challenges []Challenge
	ProveData  SectorProveData
}

// PriChange is an auto generated low-level Go binding around an user-defined struct.
type PriChange struct {
	FileHash  []byte
	Privilege uint64
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

// SectorProveData is an auto generated low-level Go binding around an user-defined struct.
type SectorProveData struct {
	ProveFileNum uint64
	BlockNum     uint64
	Proofs       []byte
	Tags         []byte
	MerklePath   []MerklePath
	PlotData     []byte
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

// UserSpace is an auto generated low-level Go binding around an user-defined struct.
type UserSpace struct {
	Used         uint64
	Remain       uint64
	Balance      uint64
	ExpireHeight *big.Int
	UpdateHeight *big.Int
}

// UserSpaceOperation is an auto generated low-level Go binding around an user-defined struct.
type UserSpaceOperation struct {
	Type  uint8
	Value uint64
}

// UserSpaceParams is an auto generated low-level Go binding around an user-defined struct.
type UserSpaceParams struct {
	WalletAddr common.Address
	Owner      common.Address
	Size       UserSpaceOperation
	BlockCount UserSpaceOperation
}

// VerifyPlotDataParams is an auto generated low-level Go binding around an user-defined struct.
type VerifyPlotDataParams struct {
	PlotInfo PlotInfo
	PlotData []byte
	Index    uint64
}

// VerifyProofWithMerklePathForFileParams is an auto generated low-level Go binding around an user-defined struct.
type VerifyProofWithMerklePathForFileParams struct {
	Version    uint64
	Proofs     []byte
	FileIds    [][]byte
	Tags       [][]byte
	Challenges []Challenge
	MerklePath []MerklePath
	RootHashes []byte
}

// WhiteList is an auto generated low-level Go binding around an user-defined struct.
type WhiteList struct {
	Addr         common.Address
	BaseHeight   uint64
	ExpireHeight uint64
}

// WhiteListParams is an auto generated low-level Go binding around an user-defined struct.
type WhiteListParams struct {
	FileHash []byte
	Op       uint8
	List     []WhiteList
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DifferenceFileOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"FileNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"FileProveFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileProveNotFileOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FirstUserSpaceOperationError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProfit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeSectorProvedInTimeError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"want\",\"type\":\"uint64\"}],\"name\":\"NotEmptySector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughPledge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"want\",\"type\":\"uint64\"}],\"name\":\"NotEnoughVolume\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"OpError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParamsError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SectorOpError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SectorProveFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"UserspaceChangeError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserspaceDeleteError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceWrongExpiredHeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroProfit\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"provLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"AddFileToSector\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"AddFileToUnSettleList\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"}],\"name\":\"AddSectorRefForFileInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"currentHeight\",\"type\":\"uint256\"}],\"name\":\"CalcDepositFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"proveTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"copyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"}],\"name\":\"CalcFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"CalculateNodePledge\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"Cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"CurOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NewOwner\",\"type\":\"address\"}],\"internalType\":\"structOwnerChange\",\"name\":\"ownerChange\",\"type\":\"tuple\"}],\"name\":\"ChangeFileOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"privilege\",\"type\":\"uint64\"}],\"internalType\":\"structPriChange\",\"name\":\"priChange\",\"type\":\"tuple\"}],\"name\":\"ChangeFilePrivilege\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"CheckNodeSectorProvedInTime\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"rmInfo\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rmList\",\"type\":\"bool\"}],\"name\":\"CleanupForDeleteFile\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"}],\"name\":\"CreateSector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"enumStorageType[]\",\"name\":\"storageType\",\"type\":\"uint8[]\"}],\"name\":\"DeleteExpiredFilesFromList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"sectorInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"DeleteFileFromSector\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"}],\"name\":\"DeleteFiles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"DeleteProveDetails\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"DeleteSector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteUnsettledFiles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteUserSpace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"NodeWallet\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"}],\"internalType\":\"structFileProveParams\",\"name\":\"fileProve\",\"type\":\"tuple\"}],\"name\":\"FileProve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FromAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ReNewTimes\",\"type\":\"uint64\"}],\"internalType\":\"structFileReNewInfo\",\"name\":\"fileReNewInfo\",\"type\":\"tuple\"}],\"name\":\"FileReNew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"HashValue\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveNum\",\"type\":\"uint64\"}],\"internalType\":\"structGenChallengeParams\",\"name\":\"gParams\",\"type\":\"tuple\"}],\"name\":\"GenChallenge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"GetFileInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"}],\"name\":\"GetFileInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetFileList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"GetNodeInfoByNodeAddr\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetNodeInfoByWalletAddr\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetNodeList\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"GetProveDetailList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"Finished\",\"type\":\"bool\"}],\"internalType\":\"structProveDetail[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorId\",\"type\":\"uint64\"}],\"internalType\":\"structSectorRef\",\"name\":\"sectorRef\",\"type\":\"tuple\"}],\"name\":\"GetSectorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"GetSectorsForNode\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetSetting\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"}],\"name\":\"GetSettingWithProveLevel\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnProveCandidateFiles\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnProvePrimaryFiles\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUnSettledFileList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"Size\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"BlockCount\",\"type\":\"tuple\"}],\"internalType\":\"structUserSpaceParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCost\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structTransferState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"Encrypt\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"EncryptPassword\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"RegisterDNS\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"BindDNS\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"DnsURL\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"WhiteList_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"Share\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"}],\"internalType\":\"structUploadOption\",\"name\":\"uploadOption\",\"type\":\"tuple\"}],\"name\":\"GetUploadStorageFee\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"TxnFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"SpaceFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ValidationFee\",\"type\":\"uint64\"}],\"internalType\":\"structStorageFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetUserSpace\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"GetWhiteList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"IsNodeRegisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"Size\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumUserSpaceType\",\"name\":\"Type\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"internalType\":\"structUserSpaceOperation\",\"name\":\"BlockCount\",\"type\":\"tuple\"}],\"internalType\":\"structUserSpaceParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"ManageUserSpace\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"NodeUpdate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"SectorInfo_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"Challenges\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"ProveFileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"BlockNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Tags\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"MerklePath_\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"PlotData\",\"type\":\"bytes\"}],\"internalType\":\"structSectorProveData\",\"name\":\"ProveData\",\"type\":\"tuple\"}],\"internalType\":\"structPrepareForPdpVerificationParams\",\"name\":\"pParams\",\"type\":\"tuple\"}],\"name\":\"PrepareForPdpVerification\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"FileIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"UpdatedChal\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"Path\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"FileInfo_\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"Success\",\"type\":\"bool\"}],\"internalType\":\"structPdpVerificationReturns\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"Register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ChallengeHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"ProveData\",\"type\":\"bytes\"}],\"internalType\":\"structSectorProveParams\",\"name\":\"sectorProve\",\"type\":\"tuple\"}],\"name\":\"SectorProve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"fileInfo\",\"type\":\"tuple\"}],\"name\":\"StoreFile\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"f\",\"type\":\"tuple\"}],\"name\":\"UpdateFileInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"_list\",\"type\":\"bytes[]\"}],\"name\":\"UpdateFileList\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_fileList\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"setting\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"newExpireHeight\",\"type\":\"uint256\"}],\"name\":\"UpdateFilesForRenew\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"UpdateNodeInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveDetailNum\",\"type\":\"uint64\"}],\"internalType\":\"structProveDetailMeta\",\"name\":\"details\",\"type\":\"tuple\"}],\"name\":\"UpdateProveDetailMeta\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"_sector\",\"type\":\"tuple\"}],\"name\":\"UpdateSectorInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Remain\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Balance\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"UpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structUserSpace\",\"name\":\"_userSpace\",\"type\":\"tuple\"}],\"name\":\"UpdateUserSpace\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"PlotData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"}],\"internalType\":\"structVerifyPlotDataParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"VerifyPlotData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"FileIds\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"Challenges\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"MerklePath_\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"}],\"internalType\":\"structVerifyProofWithMerklePathForFileParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"VerifyProofWithMerklePathForFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"enumWhiteListOpType\",\"name\":\"Op\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"List\",\"type\":\"tuple[]\"}],\"internalType\":\"structWhiteListParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"WhiteListOperate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"WithDrawProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0xf932268c\",\"type\":\"bytes32\"}],\"name\":\"c_0xf932268c\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractIFile\",\"name\":\"_file\",\"type\":\"address\"},{\"internalType\":\"contractIList\",\"name\":\"_list\",\"type\":\"address\"},{\"internalType\":\"contractINode\",\"name\":\"_node\",\"type\":\"address\"},{\"internalType\":\"contractIPDP\",\"name\":\"_pdp\",\"type\":\"address\"},{\"internalType\":\"contractISpace\",\"name\":\"_space\",\"type\":\"address\"},{\"internalType\":\"contractISector\",\"name\":\"_sector\",\"type\":\"address\"},{\"internalType\":\"contractIProve\",\"name\":\"_prove\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5061c14580620000226000396000f3fe6080604052600436106103a15760003560e01c80638a29e2de116101e7578063bb4e4e421161010d578063dcf74960116100a0578063ed108de91161006f578063ed108de914610dcb578063ed326a7114610de7578063f54cd29514610e03578063fd7c980814610e1f576103a1565b8063dcf7496014610d0c578063defce5d414610d28578063dfae2e4414610d65578063e3168f9e14610d8e576103a1565b8063cc76e80d116100dc578063cc76e80d14610c4d578063ce90455414610c8a578063d70e627214610cb3578063dc1ec30b14610ccf576103a1565b8063bb4e4e4214610b9b578063c5c81b2014610bb7578063c6e8392a14610bf4578063c819d86c14610c31576103a1565b80639a2b5e6311610185578063ad42030b11610154578063ad42030b14610acf578063b6af10fb14610b0c578063b6ddeca014610b35578063ba92100414610b72576103a1565b80639a2b5e6314610a0e5780639cd103e914610a4b5780639f9ca64414610a67578063aba7239614610aa4576103a1565b80639260665f116101c15780639260665f14610963578063931bb19a1461098c578063955f98b7146109b5578063977fdfe2146109d1576103a1565b80638a29e2de146108e15780638d41f9f81461090a5780638e27053014610947576103a1565b806335deb19c116102cc578063432152ce1161026a5780635baf5697116102395780635baf56971461082f578063648d225d1461086c5780636683899414610888578063681850d7146108c5576103a1565b8063432152ce1461079157806352e06146146107ba57806354e0d3c2146107d657806357d9439914610813576103a1565b80633984ff7d116102a65780633984ff7d146106b45780633ad525a9146106e45780633f2cc9a01461072357806341bc86cb14610754576103a1565b806335deb19c146106325780633778febe1461065b5780633899831a14610698576103a1565b80631a65374a1161034457806328a8565c1161031357806328a8565c1461055f5780632ba010d71461059c5780632e19aeff146105d957806334fddaac14610616576103a1565b80631a65374a146104b2578063207e33be146104ef57806322cf12cf146105185780632384a6aa14610543576103a1565b80630df7e99e116103805780630df7e99e146104075780630f9fa2eb14610430578063178e4dc014610459578063199499c014610496576103a1565b806247c003146103a657806309cbe391146103c25780630be09702146103de575b600080fd5b6103c060048036038101906103bb9190619263565b610e5c565b005b6103dc60048036038101906103d791906192cf565b610f72565b005b3480156103ea57600080fd5b5061040560048036038101906104009190618865565b611086565b005b34801561041357600080fd5b5061042e60048036038101906104299190618d46565b61119a565b005b34801561043c57600080fd5b5061045760048036038101906104529190618865565b61119d565b005b34801561046557600080fd5b50610480600480360381019061047b9190619472565b6112b1565b60405161048d919061b605565b60405180910390f35b6104b060048036038101906104ab919061908b565b6113f5565b005b3480156104be57600080fd5b506104d960048036038101906104d49190618865565b611509565b6040516104e6919061b458565b60405180910390f35b3480156104fb57600080fd5b50610516600480360381019061051191906190dd565b611647565b005b34801561052457600080fd5b5061052d61175b565b60405161053a919061b591565b60405180910390f35b61055d600480360381019061055891906191e1565b61188d565b005b34801561056b57600080fd5b5061058660048036038101906105819190618865565b6119a1565b604051610593919061b109565b60405180910390f35b3480156105a857600080fd5b506105c360048036038101906105be9190619310565b611ade565b6040516105d0919061b4fb565b60405180910390f35b3480156105e557600080fd5b5061060060048036038101906105fb919061952d565b611c21565b60405161060d919061b2ea565b60405180910390f35b610630600480360381019061062b91906188e2565b611d59565b005b34801561063e57600080fd5b50610659600480360381019061065491906195af565b611e70565b005b34801561066757600080fd5b50610682600480360381019061067d9190618865565b611f84565b60405161068f919061b458565b60405180910390f35b6106b260048036038101906106ad91906194db565b6120c2565b005b6106ce60048036038101906106c991906194db565b6121d6565b6040516106db919061b620565b60405180910390f35b3480156106f057600080fd5b5061070b600480360381019061070691906189f4565b612316565b60405161071a9392919061b1b0565b60405180910390f35b61073d60048036038101906107389190618a73565b612464565b60405161074b92919061b232565b60405180910390f35b34801561076057600080fd5b5061077b60048036038101906107769190619431565b6125ae565b604051610788919061b605565b60405180910390f35b34801561079d57600080fd5b506107b860048036038101906107b39190618972565b6126ec565b005b6107d460048036038101906107cf9190618d6f565b612800565b005b3480156107e257600080fd5b506107fd60048036038101906107f89190618865565b612914565b60405161080a919061b6b8565b60405180910390f35b61082d60048036038101906108289190619009565b612a52565b005b34801561083b57600080fd5b5061085660048036038101906108519190618d6f565b612b66565b604051610863919061b2c8565b60405180910390f35b61088660048036038101906108819190619310565b612ca3565b005b34801561089457600080fd5b506108af60048036038101906108aa9190618865565b612db7565b6040516108bc919061b2ea565b60405180910390f35b6108df60048036038101906108da9190618edf565b612eef565b005b3480156108ed57600080fd5b5061090860048036038101906109039190618e04565b613003565b005b34801561091657600080fd5b50610931600480360381019061092c9190618865565b6135e3565b60405161093e919061b109565b60405180910390f35b610961600480360381019061095c9190618fc8565b613720565b005b34801561096f57600080fd5b5061098a60048036038101906109859190618865565b613834565b005b34801561099857600080fd5b506109b360048036038101906109ae9190619310565b613948565b005b6109cf60048036038101906109ca9190619263565b613a5c565b005b3480156109dd57600080fd5b506109f860048036038101906109f39190618d6f565b613b73565b604051610a05919061b284565b60405180910390f35b348015610a1a57600080fd5b50610a356004803603810190610a309190618865565b613cb0565b604051610a42919061b109565b60405180910390f35b610a656004803603810190610a609190618f61565b613ded565b005b348015610a7357600080fd5b50610a8e6004803603810190610a89919061915f565b613f07565b604051610a9b919061b495565b60405180910390f35b348015610ab057600080fd5b50610ab961404a565b604051610ac6919061b262565b60405180910390f35b348015610adb57600080fd5b50610af66004803603810190610af1919061908b565b61417a565b604051610b03919061b739565b60405180910390f35b348015610b1857600080fd5b50610b336004803603810190610b2e91906191a0565b6142b2565b005b348015610b4157600080fd5b50610b5c6004803603810190610b57919061956e565b6143c6565b604051610b69919061b2ea565b60405180910390f35b348015610b7e57600080fd5b50610b996004803603810190610b9491906191e1565b6144fe565b005b610bb56004803603810190610bb09190618db0565b614612565b005b348015610bc357600080fd5b50610bde6004803603810190610bd99190618eb6565b614729565b604051610beb919061b591565b60405180910390f35b348015610c0057600080fd5b50610c1b6004803603810190610c169190618972565b614868565b604051610c28919061b210565b60405180910390f35b610c4b6004803603810190610c46919061908b565b6149a5565b005b348015610c5957600080fd5b50610c746004803603810190610c6f9190619363565b614ab9565b604051610c81919061b605565b60405180910390f35b348015610c9657600080fd5b50610cb16004803603810190610cac9190618d6f565b614c03565b005b610ccd6004803603810190610cc8919061888e565b614d17565b005b348015610cdb57600080fd5b50610cf66004803603810190610cf19190618865565b614e2e565b604051610d03919061b109565b60405180910390f35b610d266004803603810190610d2191906191e1565b614f6b565b005b348015610d3457600080fd5b50610d4f6004803603810190610d4a9190618d6f565b61507f565b604051610d5c919061b392565b60405180910390f35b348015610d7157600080fd5b50610d8c6004803603810190610d879190618865565b6151c2565b005b348015610d9a57600080fd5b50610db56004803603810190610db09190618865565b6152d6565b604051610dc2919061b2a6565b60405180910390f35b610de56004803603810190610de09190618936565b615413565b005b610e016004803603810190610dfc919061908b565b61552a565b005b610e1d6004803603810190610e189190618edf565b61563e565b005b348015610e2b57600080fd5b50610e466004803603810190610e41919061904a565b615752565b604051610e53919061b1ee565b60405180910390f35b610e887fd2e806315dc2a4ba763e0a293e4acd30f97718d9aede9ecc18af2b90eaa4808260001b61119a565b610eb47fec00034f4f2bc5a39e2910dc3a9c640418b5afced4e81c3f6e3163ab1cf9827060001b61119a565b610ee07fa49a149498b034dbe17e6165cac34616fe0115b0b45c26f8d392357d276c0f5560001b61119a565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166247c00383836040518363ffffffff1660e01b8152600401610f3c92919061b51d565b600060405180830381600087803b158015610f5657600080fd5b505af1158015610f6a573d6000803e3d6000fd5b505050505050565b610f9e7fce109943eeeb25c775ce072034abfb93211c3e24990964642c72bde6bef766f960001b61119a565b610fca7f30defc0f6682b76eb60b8a3ffd5b8488d39eff211d50c025be3ae3207fcbb3b860001b61119a565b610ff67f1cc351ba91eddc906805fe3619d1e40b8a39837704006f4f33ba18466959e44c60001b61119a565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166309cbe391826040518263ffffffff1660e01b8152600401611051919061b554565b600060405180830381600087803b15801561106b57600080fd5b505af115801561107f573d6000803e3d6000fd5b5050505050565b6110b27fb4dd96101b3788ab7617accc638684af496c52b3f1fa8b2bbc5a9716534ef1a960001b61119a565b6110de7f731e5548d061f6c1d77ff1d26c035306f79cc896aa67561f7668839f56664ed660001b61119a565b61110a7fa09601fe43324eeff4b68ec691b3ca813080687104f3fee22ff0d90c87012eee60001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630be09702826040518263ffffffff1660e01b8152600401611165919061b065565b600060405180830381600087803b15801561117f57600080fd5b505af1158015611193573d6000803e3d6000fd5b5050505050565b50565b6111c97f88d19bee915b54012625f8db26323685142c7d918a3033e90254d6b507afebca60001b61119a565b6111f57f28612fb58a9b3c0c1270b8b4203e0354223227c282cc6c9161de70324e22072460001b61119a565b6112217ffb004402f3c64e970ace50f71c8efa71adfec08af24a89f99b06c4f72cff8e9060001b61119a565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630f9fa2eb826040518263ffffffff1660e01b815260040161127c919061b065565b600060405180830381600087803b15801561129657600080fd5b505af11580156112aa573d6000803e3d6000fd5b5050505050565b6112b96158b3565b6112e57fc5f428682777b52a1195c58cf6dbf0782469202a04cb62c1a53a60398068cad960001b61119a565b6113117fc026b1b8bbda4e5cf40751b4d76eb9297fcf50319eda18a8bbaa81f7f22a486260001b61119a565b61133d7f4dec6b376a1dd0f297d069eb339ee64182c990588f487fc2d2b3cf609ecae4da60001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663178e4dc08585856040518463ffffffff1660e01b815260040161139c9392919061b65d565b60606040518083038186803b1580156113b457600080fd5b505afa1580156113c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113ec91906193df565b90509392505050565b6114217f70dbd4859f1d5367f26caa4813ad107a2230a4cefb1e490da72b59f99c045f6460001b61119a565b61144d7f562e901b7d6d070e2ba422030a6ef229582e0a0b09b879d65f449e2677dcbf5560001b61119a565b6114797f497e6f1e758585237edfe1e40e93bbb49b26b768cbd59d6b9993339ced40e24360001b61119a565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663199499c0826040518263ffffffff1660e01b81526004016114d4919061b458565b600060405180830381600087803b1580156114ee57600080fd5b505af1158015611502573d6000803e3d6000fd5b5050505050565b6115116158f2565b61153d7f5fd2e50c5014f5bf3b12129853651ca85879826c96589f6f72a86f1d8c52f23f60001b61119a565b6115697f389c4d379e21c44eb5b21bc9aa0a54b75861070a4807de54a52617fb02f1e13660001b61119a565b6115957f67bff11df8c1e703a4987ca28ea21990b6f8bd1ed8d6d60c31a631bfa81866a160001b61119a565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a65374a836040518263ffffffff1660e01b81526004016115f0919061b065565b60e06040518083038186803b15801561160857600080fd5b505afa15801561161c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061164091906190b4565b9050919050565b6116737fcfc6251ac2342f13b1c9ad88fd534e21ef2d0c12b2074c77f9b5c3170298010260001b61119a565b61169f7fdfa764d56aa1f203918be5ae5a2d7509f395a939034a5ffd66f50ef343c2fed560001b61119a565b6116cb7fb8ad0dc66fbd8405b0588240faab036fb57b1390c3f5d3862972ac82acd3aa1b60001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663207e33be826040518263ffffffff1660e01b8152600401611726919061b473565b600060405180830381600087803b15801561174057600080fd5b505af1158015611754573d6000803e3d6000fd5b5050505050565b61176361598d565b61178f7f89a80bd4b4c5cf8680bf2c5fcef57fced3d802d7baa4864809b30a81a89b1a7060001b61119a565b6117bb7f37fa811746aeac280a5303e1aaaa791093607c338b48952f18f43b922009893560001b61119a565b6117e77f8f67295b86948472f43eb8cc425e9609c8f44feef09c942b4346566df5daa69860001b61119a565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561185057600080fd5b505afa158015611864573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118889190619339565b905090565b6118b97f88f392db3b3a0b0d85f2d64b9e58217fe909e6f2f754afad870fd813dc93dc6760001b61119a565b6118e57f5c175619f78ba3b250956c3ad91042d79188892da263efcb0d901afa4eb8f6ba60001b61119a565b6119117ff2fb3b4359a607398f9ad49dc4e42da2f8c3c54d0737057caeb24ee2a3b61bd560001b61119a565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632384a6aa826040518263ffffffff1660e01b815260040161196c919061b4fb565b600060405180830381600087803b15801561198657600080fd5b505af115801561199a573d6000803e3d6000fd5b5050505050565b60606119cf7f7f7e0919e2acdf40e901c7e40fa7b1c2df938d51c404f6bb9a25867e2d0f290e60001b61119a565b6119fb7ffdb31422e226a16a7fe9d20866d029df63e6b45a6a3eed5a2f769fe915b6915160001b61119a565b611a277fcc333576ff241bbe437adb82cee23939bbb66cf73156e508ae5ae54d9591405260001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166328a8565c836040518263ffffffff1660e01b8152600401611a82919061b065565b60006040518083038186803b158015611a9a57600080fd5b505afa158015611aae573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190611ad791906189b3565b9050919050565b611ae6615a55565b611b127f2a16411995f53db9b54bdd1dee457cc3f67f5e95657fd762a1f4164650656da160001b61119a565b611b3e7f1626dd66f27a1342e16ea89599b148b883024ca20c65107a8fb2329b61dcb1be60001b61119a565b611b6a7f63c5c0d8d33eed59c9380a8ccc7cf0c6d9015d547679d823bf08fc4a5f0773fc60001b61119a565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632ba010d7836040518263ffffffff1660e01b8152600401611bc5919061b576565b60006040518083038186803b158015611bdd57600080fd5b505afa158015611bf1573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190611c1a9190619222565b9050919050565b6000611c4f7f96d70670c3e53b479c1cb30bea01276aef6b140f499ba2c4e1431ec5dd18105360001b61119a565b611c7b7ff41770d92745d29cd55b304e21b9e983e433df06365c0a8155ac5b2bcba1b30b60001b61119a565b611ca77f90146d4c8452007ebc44bf64be1c18158f72a2a93b5f64a76df35215584002f860001b61119a565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e19aeff836040518263ffffffff1660e01b8152600401611d02919061b6d3565b60206040518083038186803b158015611d1a57600080fd5b505afa158015611d2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d529190618d1d565b9050919050565b611d857f300cf6f92a03db6d19b104dafac6a36081e907dfb767914e0a153a933f07bc2d60001b61119a565b611db17f8cd666412b05667d7434525490a2f009ec52d8352e2f8232a0e38a8d5e94b19f60001b61119a565b611ddd7fb7dd9e77d1026f10b5e49109a30facd772bc6d963bbfe4e0d8722e23a41d452460001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166334fddaac83836040518363ffffffff1660e01b8152600401611e3a92919061b0b0565b600060405180830381600087803b158015611e5457600080fd5b505af1158015611e68573d6000803e3d6000fd5b505050505050565b611e9c7f7382f9127dab7dcc89fec5b93121b9ba8301f3e38fd9e5556553b1dde8c2d75a60001b61119a565b611ec87fbe9a207edcbd89a54a2167b2ae723cfa101c01fa0dccbc3829d00daec60c611660001b61119a565b611ef47f18ca08420cae939a66637b43539964e2651fe1862232ddd62e8088fb3d5a6f2c60001b61119a565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166335deb19c826040518263ffffffff1660e01b8152600401611f4f919061b717565b600060405180830381600087803b158015611f6957600080fd5b505af1158015611f7d573d6000803e3d6000fd5b5050505050565b611f8c6158f2565b611fb87ff0fe2d266fa3cb748b9541197b24c05d0205d02dd9a1f25dd7fb711bf535fd2d60001b61119a565b611fe47f903ec7a6b26b468ef5eb11e3ff11d43fbbd7070da65767a337779f67f9c6f11660001b61119a565b6120107f0303911048bd119f61b4ead5684861be3fc67ad06b545bd8f0057f931f1b915660001b61119a565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633778febe836040518263ffffffff1660e01b815260040161206b919061b065565b60e06040518083038186803b15801561208357600080fd5b505afa158015612097573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120bb91906190b4565b9050919050565b6120ee7f84b44cfd83199133343ae6fa5ed4c4fece1cff2471ebf60bee9816bb88c4770460001b61119a565b61211a7fd032bdaf3b97387e4900efbdfdb7a21b92140aad5bbf3b7e27d05563060dc6ff60001b61119a565b6121467fadb1eefbd4f6fc77efbd19b12b4f313661711c24405d7e2884be4ce799ab8e6c60001b61119a565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633899831a826040518263ffffffff1660e01b81526004016121a1919061b69d565b600060405180830381600087803b1580156121bb57600080fd5b505af11580156121cf573d6000803e3d6000fd5b5050505050565b6121de615b42565b61220a7f44f7fb21a6eb7dbc6b3f65616ec5ea9523dd14cc50e93228a009794709dd252660001b61119a565b6122367fd34361e409db58c523f0db00ebb7601b5170f557b3a3bfc3402488f7cbd89e5560001b61119a565b6122627f95db1da3dab86c86559f7527ce1da2399fc234130e44f8e8ef5b2079ca5e169b60001b61119a565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633984ff7d836040518263ffffffff1660e01b81526004016122bd919061b69d565b606060405180830381600087803b1580156122d757600080fd5b505af11580156122eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061230f9190619408565b9050919050565b60606000806123477ffdc7af7671ac2078f4d68766b0da9a0786aac4fd0d3357253470725301593ea960001b61119a565b6123737fcd0062c96438bace13e909e44c52c2c7068a1b8449fc0a861436aab90823bd9a60001b61119a565b61239f7f9fb2a76282d2761d7a9b52153fa18b529475749f50fa4e038d32212ae488666e60001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633ad525a98787876040518463ffffffff1660e01b81526004016123fe9392919061b12b565b600060405180830381600087803b15801561241857600080fd5b505af115801561242c573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906124559190618adc565b92509250925093509350939050565b606060006124947f52b9ba10e088c50467f1c91d0c7aaef698d62e4878664eafdf0393c9722467aa60001b61119a565b6124c07fd40b248500fc8a0ec15c27df1e0f99d6a5c527a5892572d07029545ed18db04c60001b61119a565b6124ec7f3a58ba27ba5c28442d70b3f15b1c7dcfe3d6095b412ce9feb7273149d6c5a71a60001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633f2cc9a08686866040518463ffffffff1660e01b815260040161254b9392919061b170565b600060405180830381600087803b15801561256557600080fd5b505af1158015612579573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906125a29190618bc5565b91509150935093915050565b6125b66158b3565b6125e27fe8a8d62fe144028d33c5163b41639568f608113c898fec552b10a660ec6c671860001b61119a565b61260e7f7c1759c1c4f76cbd14f15ebfe64a2d9d38dedcd9d22bbb4790b3f67730c98e6860001b61119a565b61263a7f3e6949047aacaabaab9f120ab869cbb19a36cc3b319f665caf58eb712495da4b60001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166341bc86cb836040518263ffffffff1660e01b8152600401612695919061b63b565b60606040518083038186803b1580156126ad57600080fd5b505afa1580156126c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126e591906193df565b9050919050565b6127187faae72bf78ec47a65179f30bdbc7ed2892e72342f07340a079cde0803af2296ab60001b61119a565b6127447f1d566b872d17321f91e00d78ba7afa2bd3fba992cd1cdc4956007d181a65d30060001b61119a565b6127707f8ff220307c7ca3d5807724a07a27092bce7219387323c1b25ee73bf6ff18798960001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663432152ce826040518263ffffffff1660e01b81526004016127cb919061b109565b600060405180830381600087803b1580156127e557600080fd5b505af11580156127f9573d6000803e3d6000fd5b5050505050565b61282c7ff2bff2b9c2fecd7c4817fb815185351ff1834baf5f0f2667e02d14ea078f629b60001b61119a565b6128587f59569fc53b43aa9fa65717405a03ec6443298cd87f204633040885c3be34180b60001b61119a565b6128847f0113a9c2a0730beec8f0d02ed0a8fba2aa2d92b66ffcc8ff152fdb64a90dd24560001b61119a565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166352e06146826040518263ffffffff1660e01b81526004016128df919061b305565b600060405180830381600087803b1580156128f957600080fd5b505af115801561290d573d6000803e3d6000fd5b5050505050565b61291c615b99565b6129487fe319b6b80c37ea2b46467088970cf531298446f33f35a41b6665782f60ed44e460001b61119a565b6129747fff385b82f75e56d4bb300e416a2122d989cd01e3df3d55a2786d496dd7ba213f60001b61119a565b6129a07ffd1df4cf08d8eadec5ef2de4b69d2303c17c113b5cd19eac6b5908491e5da3e660001b61119a565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166354e0d3c2836040518263ffffffff1660e01b81526004016129fb919061b065565b60a06040518083038186803b158015612a1357600080fd5b505afa158015612a27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a4b9190619504565b9050919050565b612a7e7fb2d1939039cbf1776a203345c32b55e3c54e76fedf3f5bd890f7089b1cf896a960001b61119a565b612aaa7f64adbc88baff2a26766ec8ad0b7d90e0d50d91f3874f399fd0c3297f5991780d60001b61119a565b612ad67fb01cc526242d271eab287d3a61d83e7a64d5321b469e5b03cb9687d1dea1b4c460001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166357d94399826040518263ffffffff1660e01b8152600401612b31919061b414565b600060405180830381600087803b158015612b4b57600080fd5b505af1158015612b5f573d6000803e3d6000fd5b5050505050565b6060612b947f595043c25c07c14d8043a8e370e6eb107a95fc5b54bf9283ae5f430b5c8e92a260001b61119a565b612bc07f3aeb6064fc0d8fbde5deb74e2d9ef7949813697be92fe9767a41ee614745ddb460001b61119a565b612bec7f88202f146db499514ec78aec6eeb0270e2da00a81093c21315e1dd4482c7213760001b61119a565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635baf5697836040518263ffffffff1660e01b8152600401612c47919061b305565b60006040518083038186803b158015612c5f57600080fd5b505afa158015612c73573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190612c9c9190618cdc565b9050919050565b612ccf7f45788de8ad540e56709161e75a03c1dd9fa6a7b772beced2707dbd921885085260001b61119a565b612cfb7fc6648b774b6f2c073ad256fead1d228b0d940510df666a2b2539fe6cfb08e88a60001b61119a565b612d277f93db5604dbd2fda01d7405a7718d1453c3baae798ffe12bbb610a09bbdc2f2e560001b61119a565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663648d225d826040518263ffffffff1660e01b8152600401612d82919061b576565b600060405180830381600087803b158015612d9c57600080fd5b505af1158015612db0573d6000803e3d6000fd5b5050505050565b6000612de57f2b1fe0625a7c2b626900ef7ee6230342cf4b53bbc7264569ef19fced6dcff0f260001b61119a565b612e117f8a75505bc2e1e10f496d0b43900b7db55175831922f9e67a145167d78e6ceed760001b61119a565b612e3d7f17b41c85aedac7fce4f2b9dd409cd10b2f5d6babb9de3286dbbb73b38a8441bd60001b61119a565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166366838994836040518263ffffffff1660e01b8152600401612e98919061b065565b60206040518083038186803b158015612eb057600080fd5b505afa158015612ec4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ee89190618d1d565b9050919050565b612f1b7f87435aa24b4551774fcf3840aaf49c5029e49a22118a88e1560d179b472be77060001b61119a565b612f477f6203cde49a21c49bba959a47e684cf2be509ffbfa679251c9f1e529f8b2ca49460001b61119a565b612f737f942392f696535504907e1484a2836b05eb372186fe0ab1cf5847574e93f5b89a60001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663681850d7826040518263ffffffff1660e01b8152600401612fce919061b392565b600060405180830381600087803b158015612fe857600080fd5b505af1158015612ffc573d6000803e3d6000fd5b5050505050565b600060019054906101000a900460ff1661302b5760008054906101000a900460ff1615613034565b61303361588f565b5b613073576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161306a9061b372565b60405180910390fd5b60008060019054906101000a900460ff1615905080156130c3576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b6130ef7fbeb0c18c3b5959792aeedecf4ec467fd8781cfa7a2febd9a8b63cb2c2eb1299760001b61119a565b61311b7f12a025aeb60a05873cc4bee98fb3c26db209ddd25a9a1bdd043b2514251552ea60001b61119a565b6131477fe88ae820bc720d91f887f351eff7e49ec40675afe5d7ec1c632b5763f853dcac60001b61119a565b88600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506131b47fec61c2aae3d281a19ab7fb83ffd1eca0e66270053e05320ee3a269b6311a9b7c60001b61119a565b6131e07f0683080fba0a003ac9dcddb1273028c9e6c73e96b6c6bfdbb57333fa08eedf8e60001b61119a565b87600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061324d7f7cc1d8fbdc60442063a6c36c884aa0a586de70ce98707e945c3c442e1e561d8d60001b61119a565b6132797f75656f1ca1af66237896799c35733aea6a5154284f9a0c53921f869c97e26cc360001b61119a565b86600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506132e67f3a3ee52369554bf836dadbb5d5b728a32827e1f779022874bc7dfd1437bb01cf60001b61119a565b6133127f2754c3984177016ada3b0064ca93424fe264d3771eb2509aa5fc591c3240ad4d60001b61119a565b85600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061337f7fdf289e3c27361ef15525d8c1df465d26704f140c9c6b660090fd6bbb6bc8980f60001b61119a565b6133ab7fca65b36673b358aade37874f9840204b3f8d5e10ba4111749bbc47e72c473a7860001b61119a565b84600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506134187fc5f1af80fa3a2ab79cc4e1c4038f889a311e0fdd58f25f73a5cefdd106e49cf260001b61119a565b6134447f45f0c163523d1fd527777f0f62bc33bc64ee9622ddd2574c9cf8d7237848fe6360001b61119a565b81600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506134b17f9058d257e028c461883ad25241ce606ede6875b3fdbbdd0753f967f8c634299660001b61119a565b6134dd7fd162bb9ea16c9720e74048ec6b97c0e0519bc867d9f672afb43850a93033258660001b61119a565b82600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061354a7fdae49b48de6ed1eade2ef7c90a5890b0087ca77915578964798bfd09567ef04860001b61119a565b6135767f0e281ecb84fa84b7149e25ce0ff93cc89732c33bcf59dac09de1a9e0a543239860001b61119a565b83600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080156135d85760008060016101000a81548160ff0219169083151502179055505b505050505050505050565b60606136117f7c74f65a495472bde2b27ef29600f7178c2972573b2876710dbbabc5e2df789760001b61119a565b61363d7fbc889f0cbce83e92b5498f793727c3108f4dac88226c5b14f95122931b68726d60001b61119a565b6136697f9db6289575a87ee850556a589b492261241a2f553835778f7df81e2effb5b7c660001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638d41f9f8836040518263ffffffff1660e01b81526004016136c4919061b065565b60006040518083038186803b1580156136dc57600080fd5b505afa1580156136f0573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061371991906189b3565b9050919050565b61374c7f0576cf18e5c048a7fb9048fc9a2dc28529fac75fd0778a89ce053af9618b79fc60001b61119a565b6137787fcdd059375549ca8389bfdb160e98685f45eff3a5f78a880170c738973e3c726e60001b61119a565b6137a47f99ab62b18d4cbbd0f164d9e977d71079bbac52ddc0bbfd957dd8c824d482e9ed60001b61119a565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638e270530826040518263ffffffff1660e01b81526004016137ff919061b3f2565b600060405180830381600087803b15801561381957600080fd5b505af115801561382d573d6000803e3d6000fd5b5050505050565b6138607f07d074414f408453296b4cc3d6c60a713427f59549bca720ee739da6769cf18d60001b61119a565b61388c7f20935a463091c7a1977a1856caa68dbf39e509cf2b020627e59c793fb49426e660001b61119a565b6138b87fad3e7e503b60c933d82b68a7d7dd0a219452b6cdc3ab2868556112dceb30f5cc60001b61119a565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639260665f826040518263ffffffff1660e01b8152600401613913919061b065565b600060405180830381600087803b15801561392d57600080fd5b505af1158015613941573d6000803e3d6000fd5b5050505050565b6139747f4d7ed268f1602a04e0e3efe5864577e4e588d89afd47494460c2b0a7dc74691360001b61119a565b6139a07f82847ccfdd980d9cbebf02d91dcaa7a597b25c323da48f117a31497d81b6a94860001b61119a565b6139cc7f875589ae2fde8e26afb82308418104158aba292870fead11b8278c247d386b0360001b61119a565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663931bb19a826040518263ffffffff1660e01b8152600401613a27919061b576565b600060405180830381600087803b158015613a4157600080fd5b505af1158015613a55573d6000803e3d6000fd5b5050505050565b613a887f2af98c140e1a1f851e30d2e8dbcc18bcbd23288a704e73f23b51fe0fffa8b07860001b61119a565b613ab47ff489a58b8a8c93700f2e0fa898de6ca0a0df59655f8aa621d43aee9bc7ccf05060001b61119a565b613ae07f0abe7eedc48660dc78b6e77360040b70504d77d6c07ea528aea63872b4d7d24060001b61119a565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663955f98b783836040518363ffffffff1660e01b8152600401613b3d92919061b51d565b600060405180830381600087803b158015613b5757600080fd5b505af1158015613b6b573d6000803e3d6000fd5b505050505050565b6060613ba17f3caae36004d312d3cdede134d6500c4a0d655a207b7d5452f7fd0a53598c78cc60001b61119a565b613bcd7f7e31c494c0dcbba853e2aa041fbdc2605fe7859a780606a22a60362913bf965d60001b61119a565b613bf97facfb304ce86f9a7cea2e243fc4361db429bbdb89e9d0e9daca2756b21f3d84ff60001b61119a565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663977fdfe2836040518263ffffffff1660e01b8152600401613c54919061b305565b60006040518083038186803b158015613c6c57600080fd5b505afa158015613c80573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190613ca99190618c5a565b9050919050565b6060613cde7fc0f5663eb0cc9b1bec9bd5d97779765960754440addc3803524bdbfcacd4b30660001b61119a565b613d0a7f4b6a7919bb1651cb6dbc50222550454704bf79b5a7136c59b2dccb1a9aa1a2a860001b61119a565b613d367f91cb907892dbb6b6f083a6a59792fa302d948027278add4d6c5c581c7367f93f60001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639a2b5e63836040518263ffffffff1660e01b8152600401613d91919061b065565b60006040518083038186803b158015613da957600080fd5b505afa158015613dbd573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190613de691906189b3565b9050919050565b613e197fee53a51a7eada6df07c197a5bfd0e8dbd9e9a71e227a80ed953551a8556bcd9960001b61119a565b613e457f37567a62930d65b742ffd43ee239965a58bd7d818050c1ecaacae662c756268a60001b61119a565b613e717f516c4d4efd3c3a6cc9404c26cade673111ce2d9e5491ab54f27c252cdf404bdf60001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639cd103e98484846040518463ffffffff1660e01b8152600401613ed09392919061b3b4565b600060405180830381600087803b158015613eea57600080fd5b505af1158015613efe573d6000803e3d6000fd5b50505050505050565b613f0f615be6565b613f3b7fd3cb3d2327a246c4cd4e1845adf1b749149ca4efd3a091e0da449ea879011aee60001b61119a565b613f677ffda0237e28b9e349802adfd2e66a8d2e748e17c45ecc69bc92c52852a142d7c160001b61119a565b613f937fb8521e69112295e146a2a56afdc68ac6d16e6007d06223fe9fb62c99903a4cb660001b61119a565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639f9ca644836040518263ffffffff1660e01b8152600401613fee919061b4b7565b60006040518083038186803b15801561400657600080fd5b505afa15801561401a573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190614043919061911e565b9050919050565b60606140787f376d91b4d239f358a56d47d02cb1ec2d63ec5562aa4656e2ef9ac68cb12849bc60001b61119a565b6140a47f63c569f52ea47a8a1633db0e5c5c6d44148f9df3c5c1a561acdb1935f44bf71860001b61119a565b6140d07f724e43763a20d5231c8048daf871f0e7161039769758c5dc7309042c8fcc55fa60001b61119a565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663aba723966040518163ffffffff1660e01b815260040160006040518083038186803b15801561413857600080fd5b505afa15801561414c573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906141759190618c19565b905090565b60006141a87f0e4fccf47d78015c942c74a3a01e252145f526251d757e4829e269a228cbfddf60001b61119a565b6141d47fa478436545e495f2ca37f993399b9509b02a1e7af966d5d028bf19e98539f9a460001b61119a565b6142007f17f80c9df4dc33a5bf0ff5ba938d63fded1a5620ccb1ebef7b97526c4c93416b60001b61119a565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ad42030b836040518263ffffffff1660e01b815260040161425b919061b458565b60206040518083038186803b15801561427357600080fd5b505afa158015614287573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906142ab91906195f0565b9050919050565b6142de7f88795b064d758197c6231bca6ff1abf4704df764ee17f1cd2209874d9948f30c60001b61119a565b61430a7f048b2d81908f98af644d837112a9c55a076d3f86ec900aa9754976784161eaf760001b61119a565b6143367fb10088e13a7e4f3a7664a25b84eaaced2bd0fe147db5c22dae8b444adfd3f2f560001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b6af10fb826040518263ffffffff1660e01b8152600401614391919061b4d9565b600060405180830381600087803b1580156143ab57600080fd5b505af11580156143bf573d6000803e3d6000fd5b5050505050565b60006143f47f241eb8be7fb8381f61be442864539ed0cc779c3f17fbc99dce5bd7a418fd1d1a60001b61119a565b6144207fb13e96b83c33671b92aba159200347ace08fdb73ebaefc4d028c70f4a2c15b1660001b61119a565b61444c7fc60545c3b5ca81d023be3c1123aca016133a652457f1e8af4bd1ee0eff70c44f60001b61119a565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b6ddeca0836040518263ffffffff1660e01b81526004016144a7919061b6f5565b60206040518083038186803b1580156144bf57600080fd5b505afa1580156144d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906144f79190618d1d565b9050919050565b61452a7f47fbe7cc5ce3a347da0fbe7267486d2f34639e3d7d0018e7f5d9dfb2464a1f1b60001b61119a565b6145567fbab1d88b497c517df2b2c7e6e27527bf7f0278ebed835f805e50e7c8b299c10360001b61119a565b6145827f762d554c419c6e9852cf33e1dd0b7d24c5aef4435f224a383e33e7446662259d60001b61119a565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ba921004826040518263ffffffff1660e01b81526004016145dd919061b4fb565b600060405180830381600087803b1580156145f757600080fd5b505af115801561460b573d6000803e3d6000fd5b5050505050565b61463e7f493a884e44960ffd2f8a386ce5ff0de0ffdb8472b1c0bf693e8ea8de8da50a1c60001b61119a565b61466a7f339dd3d8b2e9bc5b17d5215c3d4bebcca1947b116934921c679d3e0d44ac842260001b61119a565b6146967f5257540a222088e01d96e74f08ab02067c2307a57b416200cc62514eef633a7460001b61119a565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bb4e4e4283836040518363ffffffff1660e01b81526004016146f392919061b327565b600060405180830381600087803b15801561470d57600080fd5b505af1158015614721573d6000803e3d6000fd5b505050505050565b61473161598d565b61475d7f35bb61585f2d0f79542f2b6f61390ff5c46279d93dbc95f516abbb93e6f2f32960001b61119a565b6147897fa967d1ea180303b5c4a2501e665211b25582913c57eeb66201df90d3b525e64760001b61119a565b6147b57f71c403881640f72fdd65f3f1561c1f014b9b0ea9777f41e3c2ef942cd6b3c4b860001b61119a565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c5c81b20836040518263ffffffff1660e01b8152600401614810919061b357565b6101606040518083038186803b15801561482957600080fd5b505afa15801561483d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906148619190619339565b9050919050565b60606148967f985b7ab83a4abc3507c5fc5b855574ba8b2873bc28625ad35f0eca3b61e42eeb60001b61119a565b6148c27f9be56937361888d999c7d260408f0a2a1874ac7dca1da1736222210ba25e80aa60001b61119a565b6148ee7f5e226d4585f2333a705c93eeb8cab76625c8191a3f7c6f60a1d5adb582d6b40360001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c6e8392a836040518263ffffffff1660e01b8152600401614949919061b109565b60006040518083038186803b15801561496157600080fd5b505afa158015614975573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061499e9190618b84565b9050919050565b6149d17fca17739822d208daf63af2cb9900f1223893bdbc54f694db072998f6edc8e47660001b61119a565b6149fd7fdc9cb3087cab37f4766216a14fc57d5547f311a53bf0a1f7afa8ce5867b2be9f60001b61119a565b614a297f029d24da3a322f1d4e295e72f058c12c39933e5a64c23b48d26323472a13b3e460001b61119a565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c819d86c826040518263ffffffff1660e01b8152600401614a84919061b458565b600060405180830381600087803b158015614a9e57600080fd5b505af1158015614ab2573d6000803e3d6000fd5b5050505050565b614ac16158b3565b614aed7f0ab1a5e6531a68bba3f2e36713218d6173139eea1d32da1a287fbc529e52d7fc60001b61119a565b614b197ff232141d24b602ac66494de2e8cce2c1d2abe4cba6e819607f75b656d37d351160001b61119a565b614b457f89af18d43534692b9d86dba9e8cace2870f707f99df094fc4c2fbeab59af47b660001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cc76e80d87878787876040518663ffffffff1660e01b8152600401614ba895949392919061b5ad565b60606040518083038186803b158015614bc057600080fd5b505afa158015614bd4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614bf891906193df565b905095945050505050565b614c2f7f54500a6783eda9496140f4cc8d5fa663733ab50e82a44b542b745ab2ac291ced60001b61119a565b614c5b7f033b9cb46e7d5932f9528da94064f3ec5d513c32c1ed156cf8645d4139f8e55160001b61119a565b614c877fd5b18986e590a15797aeb8bd5d19512d9070a5a9d431a49ae42823923723d11c60001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ce904554826040518263ffffffff1660e01b8152600401614ce2919061b305565b600060405180830381600087803b158015614cfc57600080fd5b505af1158015614d10573d6000803e3d6000fd5b5050505050565b614d437fe4f8e4d0894f4c857b36c1be120923e25d0b755a93780959e3f0f32c22bf7f4a60001b61119a565b614d6f7fc595bc1c69fa7cb8083d1247c8e17fe04083a0cbe2b332b8fb9a7bc8d08d23cf60001b61119a565b614d9b7fdf73244492136209d7bdbf80cc692f0ce7f10192a74cdaf2c552bb6c75f3516b60001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d70e627283836040518363ffffffff1660e01b8152600401614df892919061b080565b600060405180830381600087803b158015614e1257600080fd5b505af1158015614e26573d6000803e3d6000fd5b505050505050565b6060614e5c7f093293d5d01cba92dc385aaf1a8318ccbfee8de8a1b1c3f38c736482fffef12860001b61119a565b614e887fcd8918773c91b7439c461530ed29760c63198d5eef7edcc837c109fcfa3a481e60001b61119a565b614eb47fc9e84d858cc695223c4d76f598d3f48e1fee1b5cba922b4c963406834063025c60001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dc1ec30b836040518263ffffffff1660e01b8152600401614f0f919061b065565b60006040518083038186803b158015614f2757600080fd5b505afa158015614f3b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190614f6491906189b3565b9050919050565b614f977f104da1373a82bf2e3552e3cc8235ec0336e593bb05c10d453ac3d554a046de5560001b61119a565b614fc37fd6ee737c7737e519ad64cb70b2da2dd620bd64a0de9b4edc83a7cb47d09a92c660001b61119a565b614fef7f91879346b82ce23d5a52cabc894257d53b5980ca66267c95b2921f89ba230a1c60001b61119a565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dcf74960826040518263ffffffff1660e01b815260040161504a919061b4fb565b600060405180830381600087803b15801561506457600080fd5b505af1158015615078573d6000803e3d6000fd5b5050505050565b615087615c2b565b6150b37f548a45acd9544d17a7505fad2819b81821505f5eceb74b090ac2416817ca229260001b61119a565b6150df7fa64d753b0b83c24285141f4f12574d4844f77443b7809e6c2a9205f83f153f4160001b61119a565b61510b7f68779d2501d40fec81c72fb898e9dc7ae032e410239c6f3fafbca32a62b83ec760001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663defce5d4836040518263ffffffff1660e01b8152600401615166919061b305565b60006040518083038186803b15801561517e57600080fd5b505afa158015615192573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906151bb9190618f20565b9050919050565b6151ee7f089e4e737fa972fd0e8382657d64f1e653490b4f4ef2a539617484cf2e304a8a60001b61119a565b61521a7f5c40fe374ac06827c3fbb02593bcab7e55c5fe653bffac7116c14ba07a24403e60001b61119a565b6152467f8b3a28b38bebbd96adc1762cb4714b025b77873623d0e14b06a8e520a985350060001b61119a565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dfae2e44826040518263ffffffff1660e01b81526004016152a1919061b065565b600060405180830381600087803b1580156152bb57600080fd5b505af11580156152cf573d6000803e3d6000fd5b5050505050565b60606153047f7549f3daf8c7ca08990589b4a193c1aa97a8e581d41183d1b9b463ddf1672faf60001b61119a565b6153307f0fe2cf1f8f7bf29a3c26dbf105ccdfaa189c594c61e8325504d3cbc1db3f694160001b61119a565b61535c7f3b77180895bd39fe0e1dd3e56ea55c9358c9c273501087926cf0fd60dc56fab460001b61119a565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e3168f9e836040518263ffffffff1660e01b81526004016153b7919061b065565b60006040518083038186803b1580156153cf57600080fd5b505afa1580156153e3573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061540c9190618c9b565b9050919050565b61543f7f595294265ee4fdb866bd6e87af413a28ddd177a57b7d1106a0ab9b6bf05f29ed60001b61119a565b61546b7f059c33059b94da8f2e9d7ced40771ffbfe3af7a35306d9b0b10baa1df2348de560001b61119a565b6154977f720342c5f674c4b504f9200abd1d59d2faa36335ae45c997cebab403c179f9eb60001b61119a565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ed108de983836040518363ffffffff1660e01b81526004016154f492919061b0e0565b600060405180830381600087803b15801561550e57600080fd5b505af1158015615522573d6000803e3d6000fd5b505050505050565b6155567fd9ce3a06fba3271baf279a04c04b103b1d83521bbeb384dad07416eb1cc8bcf460001b61119a565b6155827f4888cce7c26d831f6bf44a717312305f3187321333bad49b71a437e3e762c56060001b61119a565b6155ae7f016114b4c1dbf21179a386eae1ea0c9fbfc0d38e579a087d0a0b7fe415c15ba260001b61119a565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ed326a71826040518263ffffffff1660e01b8152600401615609919061b458565b600060405180830381600087803b15801561562357600080fd5b505af1158015615637573d6000803e3d6000fd5b5050505050565b61566a7f5baaa6dea589413405948c2157373e342b5fc2c2a2d9c4b20d089679bbfcdcaf60001b61119a565b6156967f9d814bf827cc756e0f7a99073651dcf70a3c161a51f1862fd5879c3078fb987160001b61119a565b6156c27fb7bedabaa61a0f79da344e9c2f3c2cbc9ecca13e932780b8eea539d8fff921b160001b61119a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f54cd295826040518263ffffffff1660e01b815260040161571d919061b392565b600060405180830381600087803b15801561573757600080fd5b505af115801561574b573d6000803e3d6000fd5b5050505050565b60606157807fcae8017273054186cd489974cac882b9ed4b5dbfca6887ecf68c4a37c27b136d60001b61119a565b6157ac7fe2ffed2f8b1da5908bdb441c1cedd638d765b2e74cc7779993adfb982044401360001b61119a565b6157d87f5f50fa7ad86ed4d835488b9e702d788ccd0d848f4cd217e20a1a88119c2cb5b060001b61119a565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fd7c9808836040518263ffffffff1660e01b8152600401615833919061b436565b60006040518083038186803b15801561584b57600080fd5b505afa15801561585f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906158889190618b43565b9050919050565b600061589a306158a0565b15905090565b600080823b905060008111915050919050565b6040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6040518060e00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681525090565b604051806101600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b604051806101800160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160006002811115615aeb577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016000815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600015158152602001606081525090565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6040518060a00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600081525090565b6040518060e001604052806060815260200160608152602001606081526020016060815260200160608152602001615c1c615c2b565b81526020016000151581525090565b604051806102e0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff1681526020016000815260200160001515815260200160006001811115615d42577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600067ffffffffffffffff16815260200160608152602001606081526020016060815260200160006002811115615da7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600015158152602001615dbd615dc3565b81525090565b6040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6000615e15615e108461b779565b61b754565b90508083825260208201905082856020860282011115615e3457600080fd5b60005b85811015615e645781615e4a888261667a565b845260208401935060208301925050600181019050615e37565b5050509392505050565b6000615e81615e7c8461b779565b61b754565b90508083825260208201905082856020860282011115615ea057600080fd5b60005b85811015615ed05781615eb6888261668f565b845260208401935060208301925050600181019050615ea3565b5050509392505050565b6000615eed615ee88461b7a5565b61b754565b90508083825260208201905082856020860282011115615f0c57600080fd5b60005b85811015615f5657813567ffffffffffffffff811115615f2e57600080fd5b808601615f3b89826169ad565b85526020850194506020840193505050600181019050615f0f565b5050509392505050565b6000615f73615f6e8461b7a5565b61b754565b90508083825260208201905082856020860282011115615f9257600080fd5b60005b85811015615fdc57815167ffffffffffffffff811115615fb457600080fd5b808601615fc189826169d7565b85526020850194506020840193505050600181019050615f95565b5050509392505050565b6000615ff9615ff48461b7d1565b61b754565b9050808382526020820190508285602086028201111561601857600080fd5b60005b85811015616048578161602e8882616ad3565b84526020840193506020830192505060018101905061601b565b5050509392505050565b60006160656160608461b7fd565b61b754565b9050808382526020820190508285604086028201111561608457600080fd5b60005b858110156160b4578161609a8882616b27565b845260208401935060408301925050600181019050616087565b5050509392505050565b60006160d16160cc8461b7fd565b61b754565b905080838252602082019050828560408602820111156160f057600080fd5b60005b8581101561612057816161068882616b73565b8452602084019350604083019250506001810190506160f3565b5050509392505050565b600061613d6161388461b829565b61b754565b9050808382526020820190508285602086028201111561615c57600080fd5b60005b858110156161a657815167ffffffffffffffff81111561617e57600080fd5b80860161618b8982616e5f565b8552602085019450602084019350505060018101905061615f565b5050509392505050565b60006161c36161be8461b855565b61b754565b905080838252602082019050828560208602820111156161e257600080fd5b60005b8581101561622c57813567ffffffffffffffff81111561620457600080fd5b80860161621189826172cf565b855260208501945060208401935050506001810190506161e5565b5050509392505050565b60006162496162448461b855565b61b754565b9050808382526020820190508285602086028201111561626857600080fd5b60005b858110156162b257815167ffffffffffffffff81111561628a57600080fd5b8086016162978982617347565b8552602085019450602084019350505060018101905061626b565b5050509392505050565b60006162cf6162ca8461b881565b61b754565b905080838252602082019050828560208602820111156162ee57600080fd5b60005b8581101561633857813567ffffffffffffffff81111561631057600080fd5b80860161631d89826173bf565b855260208501945060208401935050506001810190506162f1565b5050509392505050565b60006163556163508461b881565b61b754565b9050808382526020820190508285602086028201111561637457600080fd5b60005b858110156163be57815167ffffffffffffffff81111561639657600080fd5b8086016163a38982617423565b85526020850194506020840193505050600181019050616377565b5050509392505050565b60006163db6163d68461b8ad565b61b754565b905080838252602082019050828560e08602820111156163fa57600080fd5b60005b8581101561642a57816164108882617537565b845260208401935060e083019250506001810190506163fd565b5050509392505050565b60006164476164428461b8d9565b61b754565b905080838252602082019050828560a086028201111561646657600080fd5b60005b85811015616496578161647c88826179b7565b845260208401935060a08301925050600181019050616469565b5050509392505050565b60006164b36164ae8461b905565b61b754565b905080838252602082019050828560208602820111156164d257600080fd5b60005b8581101561651c57815167ffffffffffffffff8111156164f457600080fd5b8086016165018982617b75565b855260208501945060208401935050506001810190506164d5565b5050509392505050565b60006165396165348461b931565b61b754565b9050808382526020820190508285606086028201111561655857600080fd5b60005b85811015616588578161656e8882618727565b84526020840193506060830192505060018101905061655b565b5050509392505050565b60006165a56165a08461b931565b61b754565b905080838252602082019050828560608602820111156165c457600080fd5b60005b858110156165f457816165da8882618787565b8452602084019350606083019250506001810190506165c7565b5050509392505050565b600061661161660c8461b95d565b61b754565b90508281526020810184848401111561662957600080fd5b61663484828561be0c565b509392505050565b600061664f61664a8461b95d565b61b754565b90508281526020810184848401111561666757600080fd5b61667284828561be1b565b509392505050565b6000813590506166898161bf8d565b92915050565b60008151905061669e8161bf8d565b92915050565b600082601f8301126166b557600080fd5b81356166c5848260208601615e02565b91505092915050565b600082601f8301126166df57600080fd5b81516166ef848260208601615e6e565b91505092915050565b600082601f83011261670957600080fd5b8135616719848260208601615eda565b91505092915050565b600082601f83011261673357600080fd5b8151616743848260208601615f60565b91505092915050565b600082601f83011261675d57600080fd5b813561676d848260208601615fe6565b91505092915050565b600082601f83011261678757600080fd5b8135616797848260208601616052565b91505092915050565b600082601f8301126167b157600080fd5b81516167c18482602086016160be565b91505092915050565b600082601f8301126167db57600080fd5b81516167eb84826020860161612a565b91505092915050565b600082601f83011261680557600080fd5b81356168158482602086016161b0565b91505092915050565b600082601f83011261682f57600080fd5b815161683f848260208601616236565b91505092915050565b600082601f83011261685957600080fd5b81356168698482602086016162bc565b91505092915050565b600082601f83011261688357600080fd5b8151616893848260208601616342565b91505092915050565b600082601f8301126168ad57600080fd5b81516168bd8482602086016163c8565b91505092915050565b600082601f8301126168d757600080fd5b81516168e7848260208601616434565b91505092915050565b600082601f83011261690157600080fd5b81516169118482602086016164a0565b91505092915050565b600082601f83011261692b57600080fd5b813561693b848260208601616526565b91505092915050565b600082601f83011261695557600080fd5b8151616965848260208601616592565b91505092915050565b60008135905061697d8161bfa4565b92915050565b6000815190506169928161bfa4565b92915050565b6000813590506169a78161bfbb565b92915050565b600082601f8301126169be57600080fd5b81356169ce8482602086016165fe565b91505092915050565b600082601f8301126169e857600080fd5b81516169f884826020860161663c565b91505092915050565b600081359050616a108161bfd2565b92915050565b600081359050616a258161bfe9565b92915050565b600081359050616a3a8161c000565b92915050565b600081359050616a4f8161c017565b92915050565b600081359050616a648161c02e565b92915050565b600081359050616a798161c045565b92915050565b600081359050616a8e8161c05c565b92915050565b600081359050616aa38161c073565b92915050565b600081359050616ab88161c08a565b92915050565b600081519050616acd8161c08a565b92915050565b600081359050616ae28161c09a565b92915050565b600081519050616af78161c09a565b92915050565b600081359050616b0c8161c0aa565b92915050565b600081359050616b218161c0ba565b92915050565b600060408284031215616b3957600080fd5b616b43604061b754565b90506000616b5384828501618811565b6000830152506020616b6784828501618811565b60208301525092915050565b600060408284031215616b8557600080fd5b616b8f604061b754565b90506000616b9f84828501618826565b6000830152506020616bb384828501618826565b60208301525092915050565b60006103208284031215616bd257600080fd5b616bdd6102e061b754565b9050600082013567ffffffffffffffff811115616bf957600080fd5b616c05848285016169ad565b6000830152506020616c198482850161667a565b602083015250604082013567ffffffffffffffff811115616c3957600080fd5b616c45848285016169ad565b6040830152506060616c598482850161883b565b6060830152506080616c6d8482850161883b565b60808301525060a0616c818482850161883b565b60a08301525060c0616c958482850161883b565b60c08301525060e0616ca98482850161883b565b60e083015250610100616cbe848285016187e7565b61010083015250610120616cd48482850161883b565b61012083015250610140616cea8482850161883b565b6101408301525061016082013567ffffffffffffffff811115616d0c57600080fd5b616d18848285016169ad565b61016083015250610180616d2e8482850161883b565b610180830152506101a0616d44848285016187e7565b6101a0830152506101c0616d5a8482850161696e565b6101c0830152506101e0616d7084828501616ad3565b6101e083015250610200616d868482850161883b565b6102008301525061022082013567ffffffffffffffff811115616da857600080fd5b616db4848285016166a4565b6102208301525061024082013567ffffffffffffffff811115616dd657600080fd5b616de2848285016166a4565b6102408301525061026082013567ffffffffffffffff811115616e0457600080fd5b616e10848285016169ad565b61026083015250610280616e2684828501616aa9565b610280830152506102a0616e3c8482850161696e565b6102a0830152506102c0616e528482850161779f565b6102c08301525092915050565b60006103208284031215616e7257600080fd5b616e7d6102e061b754565b9050600082015167ffffffffffffffff811115616e9957600080fd5b616ea5848285016169d7565b6000830152506020616eb98482850161668f565b602083015250604082015167ffffffffffffffff811115616ed957600080fd5b616ee5848285016169d7565b6040830152506060616ef984828501618850565b6060830152506080616f0d84828501618850565b60808301525060a0616f2184828501618850565b60a08301525060c0616f3584828501618850565b60c08301525060e0616f4984828501618850565b60e083015250610100616f5e848285016187fc565b61010083015250610120616f7484828501618850565b61012083015250610140616f8a84828501618850565b6101408301525061016082015167ffffffffffffffff811115616fac57600080fd5b616fb8848285016169d7565b61016083015250610180616fce84828501618850565b610180830152506101a0616fe4848285016187fc565b6101a0830152506101c0616ffa84828501616983565b6101c0830152506101e061701084828501616ae8565b6101e08301525061020061702684828501618850565b6102008301525061022082015167ffffffffffffffff81111561704857600080fd5b617054848285016166ce565b6102208301525061024082015167ffffffffffffffff81111561707657600080fd5b617082848285016166ce565b6102408301525061026082015167ffffffffffffffff8111156170a457600080fd5b6170b0848285016169d7565b610260830152506102806170c684828501616abe565b610280830152506102a06170dc84828501616983565b6102a0830152506102c06170f2848285016177ff565b6102c08301525092915050565b600060c0828403121561711157600080fd5b61711b60c061b754565b9050600082013567ffffffffffffffff81111561713757600080fd5b617143848285016169ad565b600083015250602082013567ffffffffffffffff81111561716357600080fd5b61716f848285016169ad565b6020830152506040617183848285016187e7565b60408301525060606171978482850161667a565b60608301525060806171ab8482850161883b565b60808301525060a06171bf8482850161883b565b60a08301525092915050565b6000606082840312156171dd57600080fd5b6171e7606061b754565b9050600082013567ffffffffffffffff81111561720357600080fd5b61720f848285016169ad565b60008301525060206172238482850161667a565b60208301525060406172378482850161883b565b60408301525092915050565b60006080828403121561725557600080fd5b61725f608061b754565b9050600061726f8482850161667a565b600083015250602082013567ffffffffffffffff81111561728f57600080fd5b61729b848285016169ad565b60208301525060406172af8482850161883b565b60408301525060606172c38482850161883b565b60608301525092915050565b6000606082840312156172e157600080fd5b6172eb606061b754565b905060006172fb8482850161883b565b600083015250602061730f8482850161883b565b602083015250604082013567ffffffffffffffff81111561732f57600080fd5b61733b848285016169ad565b60408301525092915050565b60006060828403121561735957600080fd5b617363606061b754565b9050600061737384828501618850565b600083015250602061738784828501618850565b602083015250604082015167ffffffffffffffff8111156173a757600080fd5b6173b3848285016169d7565b60408301525092915050565b6000604082840312156173d157600080fd5b6173db604061b754565b905060006173eb8482850161883b565b600083015250602082013567ffffffffffffffff81111561740b57600080fd5b617417848285016167f4565b60208301525092915050565b60006040828403121561743557600080fd5b61743f604061b754565b9050600061744f84828501618850565b600083015250602082015167ffffffffffffffff81111561746f57600080fd5b61747b8482850161681e565b60208301525092915050565b600060e0828403121561749957600080fd5b6174a360e061b754565b905060006174b38482850161883b565b60008301525060206174c78482850161883b565b60208301525060406174db8482850161883b565b60408301525060606174ef8482850161883b565b60608301525060806175038482850161883b565b60808301525060a06175178482850161667a565b60a08301525060c061752b8482850161667a565b60c08301525092915050565b600060e0828403121561754957600080fd5b61755360e061b754565b9050600061756384828501618850565b600083015250602061757784828501618850565b602083015250604061758b84828501618850565b604083015250606061759f84828501618850565b60608301525060806175b384828501618850565b60808301525060a06175c78482850161668f565b60a08301525060c06175db8482850161668f565b60c08301525092915050565b6000606082840312156175f957600080fd5b617603606061b754565b9050600082013567ffffffffffffffff81111561761f57600080fd5b61762b848285016169ad565b600083015250602061763f8482850161667a565b60208301525060406176538482850161667a565b60408301525092915050565b600060e0828403121561767157600080fd5b61767b60e061b754565b9050600082015167ffffffffffffffff81111561769757600080fd5b6176a384828501616722565b600083015250602082015167ffffffffffffffff8111156176c357600080fd5b6176cf84828501616722565b602083015250604082015167ffffffffffffffff8111156176ef57600080fd5b6176fb848285016167a0565b604083015250606082015167ffffffffffffffff81111561771b57600080fd5b61772784828501616872565b606083015250608082015167ffffffffffffffff81111561774757600080fd5b617753848285016169d7565b60808301525060a082015167ffffffffffffffff81111561777357600080fd5b61777f84828501616e5f565b60a08301525060c061779384828501616983565b60c08301525092915050565b6000606082840312156177b157600080fd5b6177bb606061b754565b905060006177cb8482850161883b565b60008301525060206177df8482850161883b565b60208301525060406177f38482850161883b565b60408301525092915050565b60006060828403121561781157600080fd5b61781b606061b754565b9050600061782b84828501618850565b600083015250602061783f84828501618850565b602083015250604061785384828501618850565b60408301525092915050565b60006060828403121561787157600080fd5b61787b606061b754565b9050600082013567ffffffffffffffff81111561789757600080fd5b6178a384828501617a3f565b600083015250602082013567ffffffffffffffff8111156178c357600080fd5b6178cf84828501616776565b602083015250604082013567ffffffffffffffff8111156178ef57600080fd5b6178fb84828501617cab565b60408301525092915050565b60006040828403121561791957600080fd5b617923604061b754565b9050600082013567ffffffffffffffff81111561793f57600080fd5b61794b848285016169ad565b600083015250602061795f8482850161883b565b60208301525092915050565b60006040828403121561797d57600080fd5b617987604061b754565b905060006179978482850161883b565b60008301525060206179ab8482850161883b565b60208301525092915050565b600060a082840312156179c957600080fd5b6179d360a061b754565b905060006179e38482850161668f565b60008301525060206179f78482850161668f565b6020830152506040617a0b84828501618850565b6040830152506060617a1f848285016187fc565b6060830152506080617a3384828501616983565b60808301525092915050565b60006101808284031215617a5257600080fd5b617a5d61018061b754565b90506000617a6d8482850161667a565b6000830152506020617a818482850161883b565b6020830152506040617a958482850161883b565b6040830152506060617aa98482850161883b565b6060830152506080617abd84828501616aa9565b60808301525060a0617ad1848285016187e7565b60a08301525060c0617ae5848285016187e7565b60c08301525060e0617af98482850161883b565b60e083015250610100617b0e8482850161883b565b61010083015250610120617b248482850161883b565b61012083015250610140617b3a8482850161696e565b6101408301525061016082013567ffffffffffffffff811115617b5c57600080fd5b617b68848285016166f8565b6101608301525092915050565b60006101808284031215617b8857600080fd5b617b9361018061b754565b90506000617ba38482850161668f565b6000830152506020617bb784828501618850565b6020830152506040617bcb84828501618850565b6040830152506060617bdf84828501618850565b6060830152506080617bf384828501616abe565b60808301525060a0617c07848285016187fc565b60a08301525060c0617c1b848285016187fc565b60c08301525060e0617c2f84828501618850565b60e083015250610100617c4484828501618850565b61010083015250610120617c5a84828501618850565b61012083015250610140617c7084828501616983565b6101408301525061016082015167ffffffffffffffff811115617c9257600080fd5b617c9e84828501616722565b6101608301525092915050565b600060c08284031215617cbd57600080fd5b617cc760c061b754565b90506000617cd78482850161883b565b6000830152506020617ceb8482850161883b565b602083015250604082013567ffffffffffffffff811115617d0b57600080fd5b617d17848285016169ad565b604083015250606082013567ffffffffffffffff811115617d3757600080fd5b617d43848285016169ad565b606083015250608082013567ffffffffffffffff811115617d6357600080fd5b617d6f84828501616848565b60808301525060a082013567ffffffffffffffff811115617d8f57600080fd5b617d9b848285016169ad565b60a08301525092915050565b600060808284031215617db957600080fd5b617dc3608061b754565b90506000617dd38482850161667a565b6000830152506020617de78482850161883b565b6020830152506040617dfb8482850161883b565b604083015250606082013567ffffffffffffffff811115617e1b57600080fd5b617e27848285016169ad565b60608301525092915050565b600060408284031215617e4557600080fd5b617e4f604061b754565b90506000617e5f8482850161667a565b6000830152506020617e738482850161883b565b60208301525092915050565b60006101608284031215617e9257600080fd5b617e9d61016061b754565b90506000617ead8482850161883b565b6000830152506020617ec18482850161883b565b6020830152506040617ed58482850161883b565b6040830152506060617ee98482850161883b565b6060830152506080617efd8482850161883b565b60808301525060a0617f118482850161883b565b60a08301525060c0617f258482850161883b565b60c08301525060e0617f398482850161883b565b60e083015250610100617f4e8482850161883b565b61010083015250610120617f648482850161883b565b61012083015250610140617f7a8482850161883b565b6101408301525092915050565b60006101608284031215617f9a57600080fd5b617fa561016061b754565b90506000617fb584828501618850565b6000830152506020617fc984828501618850565b6020830152506040617fdd84828501618850565b6040830152506060617ff184828501618850565b606083015250608061800584828501618850565b60808301525060a061801984828501618850565b60a08301525060c061802d84828501618850565b60c08301525060e061804184828501618850565b60e08301525061010061805684828501618850565b6101008301525061012061806c84828501618850565b6101208301525061014061808284828501618850565b6101408301525092915050565b6000606082840312156180a157600080fd5b6180ab606061b754565b905060006180bb84828501618850565b60008301525060206180cf84828501618850565b60208301525060406180e384828501618850565b60408301525092915050565b60006060828403121561810157600080fd5b61810b606061b754565b9050600061811b8482850161668f565b600083015250602061812f8482850161668f565b602083015250604061814384828501618850565b60408301525092915050565b60006101e0828403121561816257600080fd5b61816d6101e061b754565b9050600082013567ffffffffffffffff81111561818957600080fd5b618195848285016169ad565b60008301525060206181a98482850161883b565b60208301525060406181bd8482850161883b565b60408301525060606181d18482850161883b565b60608301525060806181e5848285016187e7565b60808301525060a06181f98482850161883b565b60a08301525060c061820d8482850161883b565b60c08301525060e06182218482850161696e565b60e08301525061010082013567ffffffffffffffff81111561824257600080fd5b61824e848285016169ad565b610100830152506101206182648482850161696e565b6101208301525061014061827a8482850161696e565b6101408301525061016082013567ffffffffffffffff81111561829c57600080fd5b6182a8848285016169ad565b6101608301525061018082013567ffffffffffffffff8111156182ca57600080fd5b6182d68482850161691a565b610180830152506101a06182ec8482850161696e565b6101a0830152506101c061830284828501616ad3565b6101c08301525092915050565b60006040828403121561832157600080fd5b61832b604061b754565b9050600061833b84828501616afd565b600083015250602061834f8482850161883b565b60208301525092915050565b600060c0828403121561836d57600080fd5b618377608061b754565b905060006183878482850161667a565b600083015250602061839b8482850161667a565b60208301525060406183af8482850161830f565b60408301525060806183c38482850161830f565b60608301525092915050565b600060a082840312156183e157600080fd5b6183eb60a061b754565b905060006183fb8482850161883b565b600083015250602061840f8482850161883b565b60208301525060406184238482850161883b565b6040830152506060618437848285016187e7565b606083015250608061844b848285016187e7565b60808301525092915050565b600060a0828403121561846957600080fd5b61847360a061b754565b9050600061848384828501618850565b600083015250602061849784828501618850565b60208301525060406184ab84828501618850565b60408301525060606184bf848285016187fc565b60608301525060806184d3848285016187fc565b60808301525092915050565b600060a082840312156184f157600080fd5b6184fb606061b754565b9050600061850b8482850161779f565b600083015250606082013567ffffffffffffffff81111561852b57600080fd5b618537848285016169ad565b602083015250608061854b8482850161883b565b60408301525092915050565b600060e0828403121561856957600080fd5b61857360e061b754565b905060006185838482850161883b565b600083015250602082013567ffffffffffffffff8111156185a357600080fd5b6185af848285016169ad565b602083015250604082013567ffffffffffffffff8111156185cf57600080fd5b6185db848285016166f8565b604083015250606082013567ffffffffffffffff8111156185fb57600080fd5b618607848285016166f8565b606083015250608082013567ffffffffffffffff81111561862757600080fd5b61863384828501616776565b60808301525060a082013567ffffffffffffffff81111561865357600080fd5b61865f84828501616848565b60a08301525060c082013567ffffffffffffffff81111561867f57600080fd5b61868b848285016169ad565b60c08301525092915050565b6000606082840312156186a957600080fd5b6186b3606061b754565b9050600082013567ffffffffffffffff8111156186cf57600080fd5b6186db848285016169ad565b60008301525060206186ef84828501616b12565b602083015250604082013567ffffffffffffffff81111561870f57600080fd5b61871b8482850161691a565b60408301525092915050565b60006060828403121561873957600080fd5b618743606061b754565b905060006187538482850161667a565b60008301525060206187678482850161883b565b602083015250604061877b8482850161883b565b60408301525092915050565b60006060828403121561879957600080fd5b6187a3606061b754565b905060006187b38482850161668f565b60008301525060206187c784828501618850565b60208301525060406187db84828501618850565b60408301525092915050565b6000813590506187f68161c0ca565b92915050565b60008151905061880b8161c0ca565b92915050565b6000813590506188208161c0e1565b92915050565b6000815190506188358161c0e1565b92915050565b60008135905061884a8161c0f8565b92915050565b60008151905061885f8161c0f8565b92915050565b60006020828403121561887757600080fd5b60006188858482850161667a565b91505092915050565b600080604083850312156188a157600080fd5b60006188af8582860161667a565b925050602083013567ffffffffffffffff8111156188cc57600080fd5b6188d8858286016166f8565b9150509250929050565b600080604083850312156188f557600080fd5b60006189038582860161667a565b925050602083013567ffffffffffffffff81111561892057600080fd5b61892c858286016169ad565b9150509250929050565b60008060c0838503121561894957600080fd5b60006189578582860161667a565b9250506020618968858286016183cf565b9150509250929050565b60006020828403121561898457600080fd5b600082013567ffffffffffffffff81111561899e57600080fd5b6189aa848285016166f8565b91505092915050565b6000602082840312156189c557600080fd5b600082015167ffffffffffffffff8111156189df57600080fd5b6189eb84828501616722565b91505092915050565b600080600060608486031215618a0957600080fd5b600084013567ffffffffffffffff811115618a2357600080fd5b618a2f868287016166f8565b9350506020618a408682870161667a565b925050604084013567ffffffffffffffff811115618a5d57600080fd5b618a698682870161674c565b9150509250925092565b60008060006101a08486031215618a8957600080fd5b600084013567ffffffffffffffff811115618aa357600080fd5b618aaf868287016166f8565b9350506020618ac086828701617e7f565b925050610180618ad2868287016187e7565b9150509250925092565b600080600060608486031215618af157600080fd5b600084015167ffffffffffffffff811115618b0b57600080fd5b618b1786828701616722565b9350506020618b2886828701618850565b9250506040618b3986828701616983565b9150509250925092565b600060208284031215618b5557600080fd5b600082015167ffffffffffffffff811115618b6f57600080fd5b618b7b848285016167a0565b91505092915050565b600060208284031215618b9657600080fd5b600082015167ffffffffffffffff811115618bb057600080fd5b618bbc848285016167ca565b91505092915050565b60008060408385031215618bd857600080fd5b600083015167ffffffffffffffff811115618bf257600080fd5b618bfe858286016167ca565b9250506020618c0f85828601616983565b9150509250929050565b600060208284031215618c2b57600080fd5b600082015167ffffffffffffffff811115618c4557600080fd5b618c518482850161689c565b91505092915050565b600060208284031215618c6c57600080fd5b600082015167ffffffffffffffff811115618c8657600080fd5b618c92848285016168c6565b91505092915050565b600060208284031215618cad57600080fd5b600082015167ffffffffffffffff811115618cc757600080fd5b618cd3848285016168f0565b91505092915050565b600060208284031215618cee57600080fd5b600082015167ffffffffffffffff811115618d0857600080fd5b618d1484828501616944565b91505092915050565b600060208284031215618d2f57600080fd5b6000618d3d84828501616983565b91505092915050565b600060208284031215618d5857600080fd5b6000618d6684828501616998565b91505092915050565b600060208284031215618d8157600080fd5b600082013567ffffffffffffffff811115618d9b57600080fd5b618da7848285016169ad565b91505092915050565b60008060608385031215618dc357600080fd5b600083013567ffffffffffffffff811115618ddd57600080fd5b618de9858286016169ad565b9250506020618dfa8582860161796b565b9150509250929050565b600080600080600080600080610100898b031215618e2157600080fd5b6000618e2f8b828c01616a01565b9850506020618e408b828c01616a16565b9750506040618e518b828c01616a2b565b9650506060618e628b828c01616a40565b9550506080618e738b828c01616a55565b94505060a0618e848b828c01616a94565b93505060c0618e958b828c01616a7f565b92505060e0618ea68b828c01616a6a565b9150509295985092959890939650565b600060208284031215618ec857600080fd5b6000618ed684828501616aa9565b91505092915050565b600060208284031215618ef157600080fd5b600082013567ffffffffffffffff811115618f0b57600080fd5b618f1784828501616bbf565b91505092915050565b600060208284031215618f3257600080fd5b600082015167ffffffffffffffff811115618f4c57600080fd5b618f5884828501616e5f565b91505092915050565b600080600060608486031215618f7657600080fd5b600084013567ffffffffffffffff811115618f9057600080fd5b618f9c86828701616bbf565b9350506020618fad8682870161696e565b9250506040618fbe8682870161696e565b9150509250925092565b600060208284031215618fda57600080fd5b600082013567ffffffffffffffff811115618ff457600080fd5b619000848285016170ff565b91505092915050565b60006020828403121561901b57600080fd5b600082013567ffffffffffffffff81111561903557600080fd5b619041848285016171cb565b91505092915050565b60006020828403121561905c57600080fd5b600082013567ffffffffffffffff81111561907657600080fd5b61908284828501617243565b91505092915050565b600060e0828403121561909d57600080fd5b60006190ab84828501617487565b91505092915050565b600060e082840312156190c657600080fd5b60006190d484828501617537565b91505092915050565b6000602082840312156190ef57600080fd5b600082013567ffffffffffffffff81111561910957600080fd5b619115848285016175e7565b91505092915050565b60006020828403121561913057600080fd5b600082015167ffffffffffffffff81111561914a57600080fd5b6191568482850161765f565b91505092915050565b60006020828403121561917157600080fd5b600082013567ffffffffffffffff81111561918b57600080fd5b6191978482850161785f565b91505092915050565b6000602082840312156191b257600080fd5b600082013567ffffffffffffffff8111156191cc57600080fd5b6191d884828501617907565b91505092915050565b6000602082840312156191f357600080fd5b600082013567ffffffffffffffff81111561920d57600080fd5b61921984828501617a3f565b91505092915050565b60006020828403121561923457600080fd5b600082015167ffffffffffffffff81111561924e57600080fd5b61925a84828501617b75565b91505092915050565b6000806040838503121561927657600080fd5b600083013567ffffffffffffffff81111561929057600080fd5b61929c85828601617a3f565b925050602083013567ffffffffffffffff8111156192b957600080fd5b6192c585828601616bbf565b9150509250929050565b6000602082840312156192e157600080fd5b600082013567ffffffffffffffff8111156192fb57600080fd5b61930784828501617da7565b91505092915050565b60006040828403121561932257600080fd5b600061933084828501617e33565b91505092915050565b6000610160828403121561934c57600080fd5b600061935a84828501617f87565b91505092915050565b60008060008060006101e0868803121561937c57600080fd5b600061938a88828901617e7f565b95505061016061939c8882890161883b565b9450506101806193ae8882890161883b565b9350506101a06193c08882890161883b565b9250506101c06193d28882890161883b565b9150509295509295909350565b6000606082840312156193f157600080fd5b60006193ff8482850161808f565b91505092915050565b60006060828403121561941a57600080fd5b6000619428848285016180ef565b91505092915050565b60006020828403121561944357600080fd5b600082013567ffffffffffffffff81111561945d57600080fd5b6194698482850161814f565b91505092915050565b60008060006101a0848603121561948857600080fd5b600084013567ffffffffffffffff8111156194a257600080fd5b6194ae8682870161814f565b93505060206194bf86828701617e7f565b9250506101806194d1868287016187e7565b9150509250925092565b600060c082840312156194ed57600080fd5b60006194fb8482850161835b565b91505092915050565b600060a0828403121561951657600080fd5b600061952484828501618457565b91505092915050565b60006020828403121561953f57600080fd5b600082013567ffffffffffffffff81111561955957600080fd5b619565848285016184df565b91505092915050565b60006020828403121561958057600080fd5b600082013567ffffffffffffffff81111561959a57600080fd5b6195a684828501618557565b91505092915050565b6000602082840312156195c157600080fd5b600082013567ffffffffffffffff8111156195db57600080fd5b6195e784828501618697565b91505092915050565b60006020828403121561960257600080fd5b600061961084828501618850565b91505092915050565b6000619625838361970d565b60208301905092915050565b600061963d8383619cf7565b905092915050565b60006196518383619d87565b60208301905092915050565b60006196698383619dd7565b60408301905092915050565b60006196818383619e06565b905092915050565b6000619695838361a365565b905092915050565b60006196a9838361a3b5565b905092915050565b60006196bd838361a3f2565b60e08301905092915050565b60006196d5838361a729565b60a08301905092915050565b60006196ed838361a791565b905092915050565b6000619701838361afd8565b60608301905092915050565b6197168161bc72565b82525050565b6197258161bc72565b82525050565b60006197368261ba3e565b619740818561bb51565b935061974b8361b98e565b8060005b8381101561977c5781516197638882619619565b975061976e8361bac2565b92505060018101905061974f565b5085935050505092915050565b60006197948261ba49565b61979e818561bb62565b9350836020820285016197b08561b99e565b8060005b858110156197ec57848403895281516197cd8582619631565b94506197d88361bacf565b925060208a019950506001810190506197b4565b50829750879550505050505092915050565b60006198098261ba49565b619813818561bb73565b9350836020820285016198258561b99e565b8060005b8581101561986157848403895281516198428582619631565b945061984d8361bacf565b925060208a01995050600181019050619829565b50829750879550505050505092915050565b600061987e8261ba54565b619888818561bc2e565b93506198938361b9ae565b8060005b838110156198c45781516198ab8882619645565b97506198b68361badc565b925050600181019050619897565b5085935050505092915050565b60006198dc8261ba5f565b6198e6818561bb84565b93506198f18361b9be565b8060005b83811015619922578151619909888261965d565b97506199148361bae9565b9250506001810190506198f5565b5085935050505092915050565b600061993a8261ba5f565b619944818561bb95565b935061994f8361b9be565b8060005b83811015619980578151619967888261965d565b97506199728361bae9565b925050600181019050619953565b5085935050505092915050565b60006199988261ba6a565b6199a2818561bba6565b9350836020820285016199b48561b9ce565b8060005b858110156199f057848403895281516199d18582619675565b94506199dc8361baf6565b925060208a019950506001810190506199b8565b50829750879550505050505092915050565b6000619a0d8261ba75565b619a17818561bbb7565b935083602082028501619a298561b9de565b8060005b85811015619a655784840389528151619a468582619689565b9450619a518361bb03565b925060208a01995050600181019050619a2d565b50829750879550505050505092915050565b6000619a828261ba80565b619a8c818561bbc8565b935083602082028501619a9e8561b9ee565b8060005b85811015619ada5784840389528151619abb858261969d565b9450619ac68361bb10565b925060208a01995050600181019050619aa2565b50829750879550505050505092915050565b6000619af78261ba8b565b619b01818561bbd9565b9350619b0c8361b9fe565b8060005b83811015619b3d578151619b2488826196b1565b9750619b2f8361bb1d565b925050600181019050619b10565b5085935050505092915050565b6000619b558261ba96565b619b5f818561bbea565b9350619b6a8361ba0e565b8060005b83811015619b9b578151619b8288826196c9565b9750619b8d8361bb2a565b925050600181019050619b6e565b5085935050505092915050565b6000619bb38261baa1565b619bbd818561bbfb565b935083602082028501619bcf8561ba1e565b8060005b85811015619c0b5784840389528151619bec85826196e1565b9450619bf78361bb37565b925060208a01995050600181019050619bd3565b50829750879550505050505092915050565b6000619c288261baac565b619c32818561bc0c565b9350619c3d8361ba2e565b8060005b83811015619c6e578151619c5588826196f5565b9750619c608361bb44565b925050600181019050619c41565b5085935050505092915050565b6000619c868261baac565b619c90818561bc1d565b9350619c9b8361ba2e565b8060005b83811015619ccc578151619cb388826196f5565b9750619cbe8361bb44565b925050600181019050619c9f565b5085935050505092915050565b619ce28161bc84565b82525050565b619cf18161bc84565b82525050565b6000619d028261bab7565b619d0c818561bc3f565b9350619d1c81856020860161be1b565b619d258161bedd565b840191505092915050565b6000619d3b8261bab7565b619d45818561bc50565b9350619d5581856020860161be1b565b619d5e8161bedd565b840191505092915050565b619d728161bdc4565b82525050565b619d818161bdc4565b82525050565b619d908161bdd6565b82525050565b619d9f8161bde8565b82525050565b619dae8161bdfa565b82525050565b6000619dc1602e8361bc61565b9150619dcc8261beee565b604082019050919050565b604082016000820151619ded600085018261b038565b506020820151619e00602085018261b038565b50505050565b6000610320830160008301518482036000860152619e248282619cf7565b9150506020830151619e39602086018261970d565b5060408301518482036040860152619e518282619cf7565b9150506060830151619e66606086018261b047565b506080830151619e79608086018261b047565b5060a0830151619e8c60a086018261b047565b5060c0830151619e9f60c086018261b047565b5060e0830151619eb260e086018261b047565b50610100830151619ec761010086018261b01a565b50610120830151619edc61012086018261b047565b50610140830151619ef161014086018261b047565b50610160830151848203610160860152619f0b8282619cf7565b915050610180830151619f2261018086018261b047565b506101a0830151619f376101a086018261b01a565b506101c0830151619f4c6101c0860182619cd9565b506101e0830151619f616101e0860182619d87565b50610200830151619f7661020086018261b047565b50610220830151848203610220860152619f90828261972b565b915050610240830151848203610240860152619fac828261972b565b915050610260830151848203610260860152619fc88282619cf7565b915050610280830151619fdf610280860182619d69565b506102a0830151619ff46102a0860182619cd9565b506102c083015161a0096102c086018261a61d565b508091505092915050565b600061032083016000830151848203600086015261a0328282619cf7565b915050602083015161a047602086018261970d565b506040830151848203604086015261a05f8282619cf7565b915050606083015161a074606086018261b047565b50608083015161a087608086018261b047565b5060a083015161a09a60a086018261b047565b5060c083015161a0ad60c086018261b047565b5060e083015161a0c060e086018261b047565b5061010083015161a0d561010086018261b01a565b5061012083015161a0ea61012086018261b047565b5061014083015161a0ff61014086018261b047565b5061016083015184820361016086015261a1198282619cf7565b91505061018083015161a13061018086018261b047565b506101a083015161a1456101a086018261b01a565b506101c083015161a15a6101c0860182619cd9565b506101e083015161a16f6101e0860182619d87565b5061020083015161a18461020086018261b047565b5061022083015184820361022086015261a19e828261972b565b91505061024083015184820361024086015261a1ba828261972b565b91505061026083015184820361026086015261a1d68282619cf7565b91505061028083015161a1ed610280860182619d69565b506102a083015161a2026102a0860182619cd9565b506102c083015161a2176102c086018261a61d565b508091505092915050565b600060c083016000830151848203600086015261a23f8282619cf7565b9150506020830151848203602086015261a2598282619cf7565b915050604083015161a26e604086018261b01a565b50606083015161a281606086018261970d565b50608083015161a294608086018261b047565b5060a083015161a2a760a086018261b047565b508091505092915050565b6000606083016000830151848203600086015261a2cf8282619cf7565b915050602083015161a2e4602086018261970d565b50604083015161a2f7604086018261b047565b508091505092915050565b600060808301600083015161a31a600086018261970d565b506020830151848203602086015261a3328282619cf7565b915050604083015161a347604086018261b047565b50606083015161a35a606086018261b047565b508091505092915050565b600060608301600083015161a37d600086018261b047565b50602083015161a390602086018261b047565b506040830151848203604086015261a3a88282619cf7565b9150508091505092915050565b600060408301600083015161a3cd600086018261b047565b506020830151848203602086015261a3e58282619a02565b9150508091505092915050565b60e08201600082015161a408600085018261b047565b50602082015161a41b602085018261b047565b50604082015161a42e604085018261b047565b50606082015161a441606085018261b047565b50608082015161a454608085018261b047565b5060a082015161a46760a085018261970d565b5060c082015161a47a60c085018261970d565b50505050565b60e08201600082015161a496600085018261b047565b50602082015161a4a9602085018261b047565b50604082015161a4bc604085018261b047565b50606082015161a4cf606085018261b047565b50608082015161a4e2608085018261b047565b5060a082015161a4f560a085018261970d565b5060c082015161a50860c085018261970d565b50505050565b6000606083016000830151848203600086015261a52b8282619cf7565b915050602083015161a540602086018261970d565b50604083015161a553604086018261970d565b508091505092915050565b600060e083016000830151848203600086015261a57b8282619789565b9150506020830151848203602086015261a5958282619789565b9150506040830151848203604086015261a5af82826198d1565b9150506060830151848203606086015261a5c98282619a77565b9150506080830151848203608086015261a5e38282619cf7565b91505060a083015184820360a086015261a5fd8282619e06565b91505060c083015161a61260c0860182619cd9565b508091505092915050565b60608201600082015161a633600085018261b047565b50602082015161a646602085018261b047565b50604082015161a659604085018261b047565b50505050565b6000606083016000830151848203600086015261a67c828261a791565b9150506020830151848203602086015261a69682826198d1565b9150506040830151848203604086015261a6b0828261a999565b9150508091505092915050565b6000604083016000830151848203600086015261a6da8282619cf7565b915050602083015161a6ef602086018261b047565b508091505092915050565b60408201600082015161a710600085018261b047565b50602082015161a723602085018261b047565b50505050565b60a08201600082015161a73f600085018261970d565b50602082015161a752602085018261970d565b50604082015161a765604085018261b047565b50606082015161a778606085018261b01a565b50608082015161a78b6080850182619cd9565b50505050565b60006101808301600083015161a7aa600086018261970d565b50602083015161a7bd602086018261b047565b50604083015161a7d0604086018261b047565b50606083015161a7e3606086018261b047565b50608083015161a7f66080860182619d69565b5060a083015161a80960a086018261b01a565b5060c083015161a81c60c086018261b01a565b5060e083015161a82f60e086018261b047565b5061010083015161a84461010086018261b047565b5061012083015161a85961012086018261b047565b5061014083015161a86e610140860182619cd9565b5061016083015184820361016086015261a8888282619789565b9150508091505092915050565b60006101808301600083015161a8ae600086018261970d565b50602083015161a8c1602086018261b047565b50604083015161a8d4604086018261b047565b50606083015161a8e7606086018261b047565b50608083015161a8fa6080860182619d69565b5060a083015161a90d60a086018261b01a565b5060c083015161a92060c086018261b01a565b5060e083015161a93360e086018261b047565b5061010083015161a94861010086018261b047565b5061012083015161a95d61012086018261b047565b5061014083015161a972610140860182619cd9565b5061016083015184820361016086015261a98c8282619789565b9150508091505092915050565b600060c08301600083015161a9b1600086018261b047565b50602083015161a9c4602086018261b047565b506040830151848203604086015261a9dc8282619cf7565b9150506060830151848203606086015261a9f68282619cf7565b9150506080830151848203608086015261aa108282619a77565b91505060a083015184820360a086015261aa2a8282619cf7565b9150508091505092915050565b600060808301600083015161aa4f600086018261970d565b50602083015161aa62602086018261b047565b50604083015161aa75604086018261b047565b506060830151848203606086015261aa8d8282619cf7565b9150508091505092915050565b60408201600082015161aab0600085018261970d565b50602082015161aac3602085018261b047565b50505050565b6101608201600082015161aae0600085018261b047565b50602082015161aaf3602085018261b047565b50604082015161ab06604085018261b047565b50606082015161ab19606085018261b047565b50608082015161ab2c608085018261b047565b5060a082015161ab3f60a085018261b047565b5060c082015161ab5260c085018261b047565b5060e082015161ab6560e085018261b047565b5061010082015161ab7a61010085018261b047565b5061012082015161ab8f61012085018261b047565b5061014082015161aba461014085018261b047565b50505050565b60608201600082015161abc0600085018261b047565b50602082015161abd3602085018261b047565b50604082015161abe6604085018261b047565b50505050565b60608201600082015161ac02600085018261970d565b50602082015161ac15602085018261970d565b50604082015161ac28604085018261b047565b50505050565b60006101e083016000830151848203600086015261ac4c8282619cf7565b915050602083015161ac61602086018261b047565b50604083015161ac74604086018261b047565b50606083015161ac87606086018261b047565b50608083015161ac9a608086018261b01a565b5060a083015161acad60a086018261b047565b5060c083015161acc060c086018261b047565b5060e083015161acd360e0860182619cd9565b5061010083015184820361010086015261aced8282619cf7565b91505061012083015161ad04610120860182619cd9565b5061014083015161ad19610140860182619cd9565b5061016083015184820361016086015261ad338282619cf7565b91505061018083015184820361018086015261ad4f8282619c1d565b9150506101a083015161ad666101a0860182619cd9565b506101c083015161ad7b6101c0860182619d87565b508091505092915050565b60408201600082015161ad9c6000850182619d96565b50602082015161adaf602085018261b047565b50505050565b60c08201600082015161adcb600085018261970d565b50602082015161adde602085018261970d565b50604082015161adf1604085018261ad86565b50606082015161ae04608085018261ad86565b50505050565b60a08201600082015161ae20600085018261b047565b50602082015161ae33602085018261b047565b50604082015161ae46604085018261b047565b50606082015161ae59606085018261b01a565b50608082015161ae6c608085018261b01a565b50505050565b600060a08301600083015161ae8a600086018261a61d565b506020830151848203606086015261aea28282619cf7565b915050604083015161aeb7608086018261b047565b508091505092915050565b600060e08301600083015161aeda600086018261b047565b506020830151848203602086015261aef28282619cf7565b9150506040830151848203604086015261af0c8282619789565b9150506060830151848203606086015261af268282619789565b9150506080830151848203608086015261af4082826198d1565b91505060a083015184820360a086015261af5a8282619a77565b91505060c083015184820360c086015261af748282619cf7565b9150508091505092915050565b6000606083016000830151848203600086015261af9e8282619cf7565b915050602083015161afb36020860182619da5565b506040830151848203604086015261afcb8282619c1d565b9150508091505092915050565b60608201600082015161afee600085018261970d565b50602082015161b001602085018261b047565b50604082015161b014604085018261b047565b50505050565b61b0238161bd96565b82525050565b61b0328161bd96565b82525050565b61b0418161bda0565b82525050565b61b0508161bdb0565b82525050565b61b05f8161bdb0565b82525050565b600060208201905061b07a600083018461971c565b92915050565b600060408201905061b095600083018561971c565b818103602083015261b0a781846197fe565b90509392505050565b600060408201905061b0c5600083018561971c565b818103602083015261b0d78184619d30565b90509392505050565b600060c08201905061b0f5600083018561971c565b61b102602083018461ae0a565b9392505050565b6000602082019050818103600083015261b12381846197fe565b905092915050565b6000606082019050818103600083015261b14581866197fe565b905061b154602083018561971c565b818103604083015261b1668184619873565b9050949350505050565b60006101a082019050818103600083015261b18b81866197fe565b905061b19a602083018561aac9565b61b1a861018083018461b029565b949350505050565b6000606082019050818103600083015261b1ca81866197fe565b905061b1d9602083018561b056565b61b1e66040830184619ce8565b949350505050565b6000602082019050818103600083015261b208818461992f565b905092915050565b6000602082019050818103600083015261b22a818461998d565b905092915050565b6000604082019050818103600083015261b24c818561998d565b905061b25b6020830184619ce8565b9392505050565b6000602082019050818103600083015261b27c8184619aec565b905092915050565b6000602082019050818103600083015261b29e8184619b4a565b905092915050565b6000602082019050818103600083015261b2c08184619ba8565b905092915050565b6000602082019050818103600083015261b2e28184619c7b565b905092915050565b600060208201905061b2ff6000830184619ce8565b92915050565b6000602082019050818103600083015261b31f8184619d30565b905092915050565b6000606082019050818103600083015261b3418185619d30565b905061b350602083018461a6fa565b9392505050565b600060208201905061b36c6000830184619d78565b92915050565b6000602082019050818103600083015261b38b81619db4565b9050919050565b6000602082019050818103600083015261b3ac818461a014565b905092915050565b6000606082019050818103600083015261b3ce818661a014565b905061b3dd6020830185619ce8565b61b3ea6040830184619ce8565b949350505050565b6000602082019050818103600083015261b40c818461a222565b905092915050565b6000602082019050818103600083015261b42e818461a2b2565b905092915050565b6000602082019050818103600083015261b450818461a302565b905092915050565b600060e08201905061b46d600083018461a480565b92915050565b6000602082019050818103600083015261b48d818461a50e565b905092915050565b6000602082019050818103600083015261b4af818461a55e565b905092915050565b6000602082019050818103600083015261b4d1818461a65f565b905092915050565b6000602082019050818103600083015261b4f3818461a6bd565b905092915050565b6000602082019050818103600083015261b515818461a895565b905092915050565b6000604082019050818103600083015261b537818561a895565b9050818103602083015261b54b818461a014565b90509392505050565b6000602082019050818103600083015261b56e818461aa37565b905092915050565b600060408201905061b58b600083018461aa9a565b92915050565b60006101608201905061b5a7600083018461aac9565b92915050565b60006101e08201905061b5c3600083018861aac9565b61b5d161016083018761b056565b61b5df61018083018661b056565b61b5ed6101a083018561b056565b61b5fb6101c083018461b056565b9695505050505050565b600060608201905061b61a600083018461abaa565b92915050565b600060608201905061b635600083018461abec565b92915050565b6000602082019050818103600083015261b655818461ac2e565b905092915050565b60006101a082019050818103600083015261b678818661ac2e565b905061b687602083018561aac9565b61b69561018083018461b029565b949350505050565b600060c08201905061b6b2600083018461adb5565b92915050565b600060a08201905061b6cd600083018461ae0a565b92915050565b6000602082019050818103600083015261b6ed818461ae72565b905092915050565b6000602082019050818103600083015261b70f818461aec2565b905092915050565b6000602082019050818103600083015261b731818461af81565b905092915050565b600060208201905061b74e600083018461b056565b92915050565b600061b75e61b76f565b905061b76a828261be4e565b919050565b6000604051905090565b600067ffffffffffffffff82111561b7945761b79361beae565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561b7c05761b7bf61beae565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561b7ec5761b7eb61beae565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561b8185761b81761beae565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561b8445761b84361beae565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561b8705761b86f61beae565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561b89c5761b89b61beae565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561b8c85761b8c761beae565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561b8f45761b8f361beae565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561b9205761b91f61beae565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561b94c5761b94b61beae565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561b9785761b97761beae565b5b61b9818261bedd565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600061bc7d8261bd76565b9050919050565b60008115159050919050565b6000819050919050565b600061bca58261bc72565b9050919050565b600061bcb78261bc72565b9050919050565b600061bcc98261bc72565b9050919050565b600061bcdb8261bc72565b9050919050565b600061bced8261bc72565b9050919050565b600061bcff8261bc72565b9050919050565b600061bd118261bc72565b9050919050565b600061bd238261bc72565b9050919050565b600081905061bd388261bf3d565b919050565b600081905061bd4b8261bf51565b919050565b600081905061bd5e8261bf65565b919050565b600081905061bd718261bf79565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b600067ffffffffffffffff82169050919050565b600061bdcf8261bd2a565b9050919050565b600061bde18261bd3d565b9050919050565b600061bdf38261bd50565b9050919050565b600061be058261bd63565b9050919050565b82818337600083830152505050565b60005b8381101561be3957808201518184015260208101905061be1e565b8381111561be48576000848401525b50505050565b61be578261bedd565b810181811067ffffffffffffffff8211171561be765761be7561beae565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6003811061bf4e5761bf4d61be7f565b5b50565b6002811061bf625761bf6161be7f565b5b50565b6003811061bf765761bf7561be7f565b5b50565b6005811061bf8a5761bf8961be7f565b5b50565b61bf968161bc72565b811461bfa157600080fd5b50565b61bfad8161bc84565b811461bfb857600080fd5b50565b61bfc48161bc90565b811461bfcf57600080fd5b50565b61bfdb8161bc9a565b811461bfe657600080fd5b50565b61bff28161bcac565b811461bffd57600080fd5b50565b61c0098161bcbe565b811461c01457600080fd5b50565b61c0208161bcd0565b811461c02b57600080fd5b50565b61c0378161bce2565b811461c04257600080fd5b50565b61c04e8161bcf4565b811461c05957600080fd5b50565b61c0658161bd06565b811461c07057600080fd5b50565b61c07c8161bd18565b811461c08757600080fd5b50565b6003811061c09757600080fd5b50565b6002811061c0a757600080fd5b50565b6003811061c0b757600080fd5b50565b6005811061c0c757600080fd5b50565b61c0d38161bd96565b811461c0de57600080fd5b50565b61c0ea8161bda0565b811461c0f557600080fd5b50565b61c1018161bdb0565b811461c10c57600080fd5b5056fea264697066735822122065b1c4712abacdb7651d44cb1fad887b75ada11bf2eb5728419b3685f44349f564736f6c63430008040033",
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

// CalculateNodePledge is a free data retrieval call binding the contract method 0xad42030b.
//
// Solidity: function CalculateNodePledge((uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo) view returns(uint64)
func (_Store *StoreCaller) CalculateNodePledge(opts *bind.CallOpts, nodeInfo NodeInfo) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "CalculateNodePledge", nodeInfo)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalculateNodePledge is a free data retrieval call binding the contract method 0xad42030b.
//
// Solidity: function CalculateNodePledge((uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo) view returns(uint64)
func (_Store *StoreSession) CalculateNodePledge(nodeInfo NodeInfo) (uint64, error) {
	return _Store.Contract.CalculateNodePledge(&_Store.CallOpts, nodeInfo)
}

// CalculateNodePledge is a free data retrieval call binding the contract method 0xad42030b.
//
// Solidity: function CalculateNodePledge((uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo) view returns(uint64)
func (_Store *StoreCallerSession) CalculateNodePledge(nodeInfo NodeInfo) (uint64, error) {
	return _Store.Contract.CalculateNodePledge(&_Store.CallOpts, nodeInfo)
}

// GenChallenge is a free data retrieval call binding the contract method 0xfd7c9808.
//
// Solidity: function GenChallenge((address,bytes,uint64,uint64) gParams) view returns((uint32,uint32)[])
func (_Store *StoreCaller) GenChallenge(opts *bind.CallOpts, gParams GenChallengeParams) ([]Challenge, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GenChallenge", gParams)

	if err != nil {
		return *new([]Challenge), err
	}

	out0 := *abi.ConvertType(out[0], new([]Challenge)).(*[]Challenge)

	return out0, err

}

// GenChallenge is a free data retrieval call binding the contract method 0xfd7c9808.
//
// Solidity: function GenChallenge((address,bytes,uint64,uint64) gParams) view returns((uint32,uint32)[])
func (_Store *StoreSession) GenChallenge(gParams GenChallengeParams) ([]Challenge, error) {
	return _Store.Contract.GenChallenge(&_Store.CallOpts, gParams)
}

// GenChallenge is a free data retrieval call binding the contract method 0xfd7c9808.
//
// Solidity: function GenChallenge((address,bytes,uint64,uint64) gParams) view returns((uint32,uint32)[])
func (_Store *StoreCallerSession) GenChallenge(gParams GenChallengeParams) ([]Challenge, error) {
	return _Store.Contract.GenChallenge(&_Store.CallOpts, gParams)
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

// GetNodeInfoByNodeAddr is a free data retrieval call binding the contract method 0x3778febe.
//
// Solidity: function GetNodeInfoByNodeAddr(address nodeAddr) view returns((uint64,uint64,uint64,uint64,uint64,address,address))
func (_Store *StoreCaller) GetNodeInfoByNodeAddr(opts *bind.CallOpts, nodeAddr common.Address) (NodeInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetNodeInfoByNodeAddr", nodeAddr)

	if err != nil {
		return *new(NodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(NodeInfo)).(*NodeInfo)

	return out0, err

}

// GetNodeInfoByNodeAddr is a free data retrieval call binding the contract method 0x3778febe.
//
// Solidity: function GetNodeInfoByNodeAddr(address nodeAddr) view returns((uint64,uint64,uint64,uint64,uint64,address,address))
func (_Store *StoreSession) GetNodeInfoByNodeAddr(nodeAddr common.Address) (NodeInfo, error) {
	return _Store.Contract.GetNodeInfoByNodeAddr(&_Store.CallOpts, nodeAddr)
}

// GetNodeInfoByNodeAddr is a free data retrieval call binding the contract method 0x3778febe.
//
// Solidity: function GetNodeInfoByNodeAddr(address nodeAddr) view returns((uint64,uint64,uint64,uint64,uint64,address,address))
func (_Store *StoreCallerSession) GetNodeInfoByNodeAddr(nodeAddr common.Address) (NodeInfo, error) {
	return _Store.Contract.GetNodeInfoByNodeAddr(&_Store.CallOpts, nodeAddr)
}

// GetNodeInfoByWalletAddr is a free data retrieval call binding the contract method 0x1a65374a.
//
// Solidity: function GetNodeInfoByWalletAddr(address walletAddr) view returns((uint64,uint64,uint64,uint64,uint64,address,address))
func (_Store *StoreCaller) GetNodeInfoByWalletAddr(opts *bind.CallOpts, walletAddr common.Address) (NodeInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetNodeInfoByWalletAddr", walletAddr)

	if err != nil {
		return *new(NodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(NodeInfo)).(*NodeInfo)

	return out0, err

}

// GetNodeInfoByWalletAddr is a free data retrieval call binding the contract method 0x1a65374a.
//
// Solidity: function GetNodeInfoByWalletAddr(address walletAddr) view returns((uint64,uint64,uint64,uint64,uint64,address,address))
func (_Store *StoreSession) GetNodeInfoByWalletAddr(walletAddr common.Address) (NodeInfo, error) {
	return _Store.Contract.GetNodeInfoByWalletAddr(&_Store.CallOpts, walletAddr)
}

// GetNodeInfoByWalletAddr is a free data retrieval call binding the contract method 0x1a65374a.
//
// Solidity: function GetNodeInfoByWalletAddr(address walletAddr) view returns((uint64,uint64,uint64,uint64,uint64,address,address))
func (_Store *StoreCallerSession) GetNodeInfoByWalletAddr(walletAddr common.Address) (NodeInfo, error) {
	return _Store.Contract.GetNodeInfoByWalletAddr(&_Store.CallOpts, walletAddr)
}

// GetNodeList is a free data retrieval call binding the contract method 0xaba72396.
//
// Solidity: function GetNodeList() view returns((uint64,uint64,uint64,uint64,uint64,address,address)[])
func (_Store *StoreCaller) GetNodeList(opts *bind.CallOpts) ([]NodeInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetNodeList")

	if err != nil {
		return *new([]NodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]NodeInfo)).(*[]NodeInfo)

	return out0, err

}

// GetNodeList is a free data retrieval call binding the contract method 0xaba72396.
//
// Solidity: function GetNodeList() view returns((uint64,uint64,uint64,uint64,uint64,address,address)[])
func (_Store *StoreSession) GetNodeList() ([]NodeInfo, error) {
	return _Store.Contract.GetNodeList(&_Store.CallOpts)
}

// GetNodeList is a free data retrieval call binding the contract method 0xaba72396.
//
// Solidity: function GetNodeList() view returns((uint64,uint64,uint64,uint64,uint64,address,address)[])
func (_Store *StoreCallerSession) GetNodeList() ([]NodeInfo, error) {
	return _Store.Contract.GetNodeList(&_Store.CallOpts)
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

// GetSetting is a free data retrieval call binding the contract method 0x22cf12cf.
//
// Solidity: function GetSetting() view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_Store *StoreCaller) GetSetting(opts *bind.CallOpts) (Setting, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetSetting")

	if err != nil {
		return *new(Setting), err
	}

	out0 := *abi.ConvertType(out[0], new(Setting)).(*Setting)

	return out0, err

}

// GetSetting is a free data retrieval call binding the contract method 0x22cf12cf.
//
// Solidity: function GetSetting() view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_Store *StoreSession) GetSetting() (Setting, error) {
	return _Store.Contract.GetSetting(&_Store.CallOpts)
}

// GetSetting is a free data retrieval call binding the contract method 0x22cf12cf.
//
// Solidity: function GetSetting() view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_Store *StoreCallerSession) GetSetting() (Setting, error) {
	return _Store.Contract.GetSetting(&_Store.CallOpts)
}

// GetSettingWithProveLevel is a free data retrieval call binding the contract method 0xc5c81b20.
//
// Solidity: function GetSettingWithProveLevel(uint8 proveLevel) view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_Store *StoreCaller) GetSettingWithProveLevel(opts *bind.CallOpts, proveLevel uint8) (Setting, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetSettingWithProveLevel", proveLevel)

	if err != nil {
		return *new(Setting), err
	}

	out0 := *abi.ConvertType(out[0], new(Setting)).(*Setting)

	return out0, err

}

// GetSettingWithProveLevel is a free data retrieval call binding the contract method 0xc5c81b20.
//
// Solidity: function GetSettingWithProveLevel(uint8 proveLevel) view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_Store *StoreSession) GetSettingWithProveLevel(proveLevel uint8) (Setting, error) {
	return _Store.Contract.GetSettingWithProveLevel(&_Store.CallOpts, proveLevel)
}

// GetSettingWithProveLevel is a free data retrieval call binding the contract method 0xc5c81b20.
//
// Solidity: function GetSettingWithProveLevel(uint8 proveLevel) view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_Store *StoreCallerSession) GetSettingWithProveLevel(proveLevel uint8) (Setting, error) {
	return _Store.Contract.GetSettingWithProveLevel(&_Store.CallOpts, proveLevel)
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

// GetUserSpace is a free data retrieval call binding the contract method 0x54e0d3c2.
//
// Solidity: function GetUserSpace(address walletAddr) view returns((uint64,uint64,uint64,uint256,uint256))
func (_Store *StoreCaller) GetUserSpace(opts *bind.CallOpts, walletAddr common.Address) (UserSpace, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetUserSpace", walletAddr)

	if err != nil {
		return *new(UserSpace), err
	}

	out0 := *abi.ConvertType(out[0], new(UserSpace)).(*UserSpace)

	return out0, err

}

// GetUserSpace is a free data retrieval call binding the contract method 0x54e0d3c2.
//
// Solidity: function GetUserSpace(address walletAddr) view returns((uint64,uint64,uint64,uint256,uint256))
func (_Store *StoreSession) GetUserSpace(walletAddr common.Address) (UserSpace, error) {
	return _Store.Contract.GetUserSpace(&_Store.CallOpts, walletAddr)
}

// GetUserSpace is a free data retrieval call binding the contract method 0x54e0d3c2.
//
// Solidity: function GetUserSpace(address walletAddr) view returns((uint64,uint64,uint64,uint256,uint256))
func (_Store *StoreCallerSession) GetUserSpace(walletAddr common.Address) (UserSpace, error) {
	return _Store.Contract.GetUserSpace(&_Store.CallOpts, walletAddr)
}

// GetWhiteList is a free data retrieval call binding the contract method 0x5baf5697.
//
// Solidity: function GetWhiteList(bytes fileHash) view returns((address,uint64,uint64)[])
func (_Store *StoreCaller) GetWhiteList(opts *bind.CallOpts, fileHash []byte) ([]WhiteList, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetWhiteList", fileHash)

	if err != nil {
		return *new([]WhiteList), err
	}

	out0 := *abi.ConvertType(out[0], new([]WhiteList)).(*[]WhiteList)

	return out0, err

}

// GetWhiteList is a free data retrieval call binding the contract method 0x5baf5697.
//
// Solidity: function GetWhiteList(bytes fileHash) view returns((address,uint64,uint64)[])
func (_Store *StoreSession) GetWhiteList(fileHash []byte) ([]WhiteList, error) {
	return _Store.Contract.GetWhiteList(&_Store.CallOpts, fileHash)
}

// GetWhiteList is a free data retrieval call binding the contract method 0x5baf5697.
//
// Solidity: function GetWhiteList(bytes fileHash) view returns((address,uint64,uint64)[])
func (_Store *StoreCallerSession) GetWhiteList(fileHash []byte) ([]WhiteList, error) {
	return _Store.Contract.GetWhiteList(&_Store.CallOpts, fileHash)
}

// IsNodeRegisted is a free data retrieval call binding the contract method 0x66838994.
//
// Solidity: function IsNodeRegisted(address walletAddr) view returns(bool)
func (_Store *StoreCaller) IsNodeRegisted(opts *bind.CallOpts, walletAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "IsNodeRegisted", walletAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNodeRegisted is a free data retrieval call binding the contract method 0x66838994.
//
// Solidity: function IsNodeRegisted(address walletAddr) view returns(bool)
func (_Store *StoreSession) IsNodeRegisted(walletAddr common.Address) (bool, error) {
	return _Store.Contract.IsNodeRegisted(&_Store.CallOpts, walletAddr)
}

// IsNodeRegisted is a free data retrieval call binding the contract method 0x66838994.
//
// Solidity: function IsNodeRegisted(address walletAddr) view returns(bool)
func (_Store *StoreCallerSession) IsNodeRegisted(walletAddr common.Address) (bool, error) {
	return _Store.Contract.IsNodeRegisted(&_Store.CallOpts, walletAddr)
}

// PrepareForPdpVerification is a free data retrieval call binding the contract method 0x9f9ca644.
//
// Solidity: function PrepareForPdpVerification(((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]),(uint32,uint32)[],(uint64,uint64,bytes,bytes,(uint64,(uint64,uint64,bytes)[])[],bytes)) pParams) view returns((bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)),bool))
func (_Store *StoreCaller) PrepareForPdpVerification(opts *bind.CallOpts, pParams PrepareForPdpVerificationParams) (PdpVerificationReturns, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "PrepareForPdpVerification", pParams)

	if err != nil {
		return *new(PdpVerificationReturns), err
	}

	out0 := *abi.ConvertType(out[0], new(PdpVerificationReturns)).(*PdpVerificationReturns)

	return out0, err

}

// PrepareForPdpVerification is a free data retrieval call binding the contract method 0x9f9ca644.
//
// Solidity: function PrepareForPdpVerification(((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]),(uint32,uint32)[],(uint64,uint64,bytes,bytes,(uint64,(uint64,uint64,bytes)[])[],bytes)) pParams) view returns((bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)),bool))
func (_Store *StoreSession) PrepareForPdpVerification(pParams PrepareForPdpVerificationParams) (PdpVerificationReturns, error) {
	return _Store.Contract.PrepareForPdpVerification(&_Store.CallOpts, pParams)
}

// PrepareForPdpVerification is a free data retrieval call binding the contract method 0x9f9ca644.
//
// Solidity: function PrepareForPdpVerification(((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]),(uint32,uint32)[],(uint64,uint64,bytes,bytes,(uint64,(uint64,uint64,bytes)[])[],bytes)) pParams) view returns((bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)),bool))
func (_Store *StoreCallerSession) PrepareForPdpVerification(pParams PrepareForPdpVerificationParams) (PdpVerificationReturns, error) {
	return _Store.Contract.PrepareForPdpVerification(&_Store.CallOpts, pParams)
}

// VerifyPlotData is a free data retrieval call binding the contract method 0x2e19aeff.
//
// Solidity: function VerifyPlotData(((uint64,uint64,uint64),bytes,uint64) vParams) view returns(bool)
func (_Store *StoreCaller) VerifyPlotData(opts *bind.CallOpts, vParams VerifyPlotDataParams) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "VerifyPlotData", vParams)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyPlotData is a free data retrieval call binding the contract method 0x2e19aeff.
//
// Solidity: function VerifyPlotData(((uint64,uint64,uint64),bytes,uint64) vParams) view returns(bool)
func (_Store *StoreSession) VerifyPlotData(vParams VerifyPlotDataParams) (bool, error) {
	return _Store.Contract.VerifyPlotData(&_Store.CallOpts, vParams)
}

// VerifyPlotData is a free data retrieval call binding the contract method 0x2e19aeff.
//
// Solidity: function VerifyPlotData(((uint64,uint64,uint64),bytes,uint64) vParams) view returns(bool)
func (_Store *StoreCallerSession) VerifyPlotData(vParams VerifyPlotDataParams) (bool, error) {
	return _Store.Contract.VerifyPlotData(&_Store.CallOpts, vParams)
}

// VerifyProofWithMerklePathForFile is a free data retrieval call binding the contract method 0xb6ddeca0.
//
// Solidity: function VerifyProofWithMerklePathForFile((uint64,bytes,bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes) vParams) view returns(bool)
func (_Store *StoreCaller) VerifyProofWithMerklePathForFile(opts *bind.CallOpts, vParams VerifyProofWithMerklePathForFileParams) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "VerifyProofWithMerklePathForFile", vParams)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProofWithMerklePathForFile is a free data retrieval call binding the contract method 0xb6ddeca0.
//
// Solidity: function VerifyProofWithMerklePathForFile((uint64,bytes,bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes) vParams) view returns(bool)
func (_Store *StoreSession) VerifyProofWithMerklePathForFile(vParams VerifyProofWithMerklePathForFileParams) (bool, error) {
	return _Store.Contract.VerifyProofWithMerklePathForFile(&_Store.CallOpts, vParams)
}

// VerifyProofWithMerklePathForFile is a free data retrieval call binding the contract method 0xb6ddeca0.
//
// Solidity: function VerifyProofWithMerklePathForFile((uint64,bytes,bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes) vParams) view returns(bool)
func (_Store *StoreCallerSession) VerifyProofWithMerklePathForFile(vParams VerifyProofWithMerklePathForFileParams) (bool, error) {
	return _Store.Contract.VerifyProofWithMerklePathForFile(&_Store.CallOpts, vParams)
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

// AddSectorRefForFileInfo is a paid mutator transaction binding the contract method 0xdcf74960.
//
// Solidity: function AddSectorRefForFileInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) payable returns()
func (_Store *StoreTransactor) AddSectorRefForFileInfo(opts *bind.TransactOpts, sectorInfo SectorInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "AddSectorRefForFileInfo", sectorInfo)
}

// AddSectorRefForFileInfo is a paid mutator transaction binding the contract method 0xdcf74960.
//
// Solidity: function AddSectorRefForFileInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) payable returns()
func (_Store *StoreSession) AddSectorRefForFileInfo(sectorInfo SectorInfo) (*types.Transaction, error) {
	return _Store.Contract.AddSectorRefForFileInfo(&_Store.TransactOpts, sectorInfo)
}

// AddSectorRefForFileInfo is a paid mutator transaction binding the contract method 0xdcf74960.
//
// Solidity: function AddSectorRefForFileInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) sectorInfo) payable returns()
func (_Store *StoreTransactorSession) AddSectorRefForFileInfo(sectorInfo SectorInfo) (*types.Transaction, error) {
	return _Store.Contract.AddSectorRefForFileInfo(&_Store.TransactOpts, sectorInfo)
}

// Cancel is a paid mutator transaction binding the contract method 0xdfae2e44.
//
// Solidity: function Cancel(address walletAddr) returns()
func (_Store *StoreTransactor) Cancel(opts *bind.TransactOpts, walletAddr common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "Cancel", walletAddr)
}

// Cancel is a paid mutator transaction binding the contract method 0xdfae2e44.
//
// Solidity: function Cancel(address walletAddr) returns()
func (_Store *StoreSession) Cancel(walletAddr common.Address) (*types.Transaction, error) {
	return _Store.Contract.Cancel(&_Store.TransactOpts, walletAddr)
}

// Cancel is a paid mutator transaction binding the contract method 0xdfae2e44.
//
// Solidity: function Cancel(address walletAddr) returns()
func (_Store *StoreTransactorSession) Cancel(walletAddr common.Address) (*types.Transaction, error) {
	return _Store.Contract.Cancel(&_Store.TransactOpts, walletAddr)
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

// DeleteUserSpace is a paid mutator transaction binding the contract method 0x0f9fa2eb.
//
// Solidity: function DeleteUserSpace(address walletAddr) returns()
func (_Store *StoreTransactor) DeleteUserSpace(opts *bind.TransactOpts, walletAddr common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DeleteUserSpace", walletAddr)
}

// DeleteUserSpace is a paid mutator transaction binding the contract method 0x0f9fa2eb.
//
// Solidity: function DeleteUserSpace(address walletAddr) returns()
func (_Store *StoreSession) DeleteUserSpace(walletAddr common.Address) (*types.Transaction, error) {
	return _Store.Contract.DeleteUserSpace(&_Store.TransactOpts, walletAddr)
}

// DeleteUserSpace is a paid mutator transaction binding the contract method 0x0f9fa2eb.
//
// Solidity: function DeleteUserSpace(address walletAddr) returns()
func (_Store *StoreTransactorSession) DeleteUserSpace(walletAddr common.Address) (*types.Transaction, error) {
	return _Store.Contract.DeleteUserSpace(&_Store.TransactOpts, walletAddr)
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

// GetUpdateCost is a paid mutator transaction binding the contract method 0x3984ff7d.
//
// Solidity: function GetUpdateCost((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns((address,address,uint64))
func (_Store *StoreTransactor) GetUpdateCost(opts *bind.TransactOpts, params UserSpaceParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "GetUpdateCost", params)
}

// GetUpdateCost is a paid mutator transaction binding the contract method 0x3984ff7d.
//
// Solidity: function GetUpdateCost((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns((address,address,uint64))
func (_Store *StoreSession) GetUpdateCost(params UserSpaceParams) (*types.Transaction, error) {
	return _Store.Contract.GetUpdateCost(&_Store.TransactOpts, params)
}

// GetUpdateCost is a paid mutator transaction binding the contract method 0x3984ff7d.
//
// Solidity: function GetUpdateCost((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns((address,address,uint64))
func (_Store *StoreTransactorSession) GetUpdateCost(params UserSpaceParams) (*types.Transaction, error) {
	return _Store.Contract.GetUpdateCost(&_Store.TransactOpts, params)
}

// ManageUserSpace is a paid mutator transaction binding the contract method 0x3899831a.
//
// Solidity: function ManageUserSpace((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns()
func (_Store *StoreTransactor) ManageUserSpace(opts *bind.TransactOpts, params UserSpaceParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ManageUserSpace", params)
}

// ManageUserSpace is a paid mutator transaction binding the contract method 0x3899831a.
//
// Solidity: function ManageUserSpace((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns()
func (_Store *StoreSession) ManageUserSpace(params UserSpaceParams) (*types.Transaction, error) {
	return _Store.Contract.ManageUserSpace(&_Store.TransactOpts, params)
}

// ManageUserSpace is a paid mutator transaction binding the contract method 0x3899831a.
//
// Solidity: function ManageUserSpace((address,address,(uint8,uint64),(uint8,uint64)) params) payable returns()
func (_Store *StoreTransactorSession) ManageUserSpace(params UserSpaceParams) (*types.Transaction, error) {
	return _Store.Contract.ManageUserSpace(&_Store.TransactOpts, params)
}

// NodeUpdate is a paid mutator transaction binding the contract method 0xed326a71.
//
// Solidity: function NodeUpdate((uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo) payable returns()
func (_Store *StoreTransactor) NodeUpdate(opts *bind.TransactOpts, nodeInfo NodeInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "NodeUpdate", nodeInfo)
}

// NodeUpdate is a paid mutator transaction binding the contract method 0xed326a71.
//
// Solidity: function NodeUpdate((uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo) payable returns()
func (_Store *StoreSession) NodeUpdate(nodeInfo NodeInfo) (*types.Transaction, error) {
	return _Store.Contract.NodeUpdate(&_Store.TransactOpts, nodeInfo)
}

// NodeUpdate is a paid mutator transaction binding the contract method 0xed326a71.
//
// Solidity: function NodeUpdate((uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo) payable returns()
func (_Store *StoreTransactorSession) NodeUpdate(nodeInfo NodeInfo) (*types.Transaction, error) {
	return _Store.Contract.NodeUpdate(&_Store.TransactOpts, nodeInfo)
}

// Register is a paid mutator transaction binding the contract method 0x199499c0.
//
// Solidity: function Register((uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo) payable returns()
func (_Store *StoreTransactor) Register(opts *bind.TransactOpts, nodeInfo NodeInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "Register", nodeInfo)
}

// Register is a paid mutator transaction binding the contract method 0x199499c0.
//
// Solidity: function Register((uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo) payable returns()
func (_Store *StoreSession) Register(nodeInfo NodeInfo) (*types.Transaction, error) {
	return _Store.Contract.Register(&_Store.TransactOpts, nodeInfo)
}

// Register is a paid mutator transaction binding the contract method 0x199499c0.
//
// Solidity: function Register((uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo) payable returns()
func (_Store *StoreTransactorSession) Register(nodeInfo NodeInfo) (*types.Transaction, error) {
	return _Store.Contract.Register(&_Store.TransactOpts, nodeInfo)
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
// Solidity: function UpdateFileList(address walletAddr, bytes[] _list) payable returns()
func (_Store *StoreTransactor) UpdateFileList(opts *bind.TransactOpts, walletAddr common.Address, _list [][]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateFileList", walletAddr, _list)
}

// UpdateFileList is a paid mutator transaction binding the contract method 0xd70e6272.
//
// Solidity: function UpdateFileList(address walletAddr, bytes[] _list) payable returns()
func (_Store *StoreSession) UpdateFileList(walletAddr common.Address, _list [][]byte) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileList(&_Store.TransactOpts, walletAddr, _list)
}

// UpdateFileList is a paid mutator transaction binding the contract method 0xd70e6272.
//
// Solidity: function UpdateFileList(address walletAddr, bytes[] _list) payable returns()
func (_Store *StoreTransactorSession) UpdateFileList(walletAddr common.Address, _list [][]byte) (*types.Transaction, error) {
	return _Store.Contract.UpdateFileList(&_Store.TransactOpts, walletAddr, _list)
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

// UpdateNodeInfo is a paid mutator transaction binding the contract method 0xc819d86c.
//
// Solidity: function UpdateNodeInfo((uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo) payable returns()
func (_Store *StoreTransactor) UpdateNodeInfo(opts *bind.TransactOpts, nodeInfo NodeInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateNodeInfo", nodeInfo)
}

// UpdateNodeInfo is a paid mutator transaction binding the contract method 0xc819d86c.
//
// Solidity: function UpdateNodeInfo((uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo) payable returns()
func (_Store *StoreSession) UpdateNodeInfo(nodeInfo NodeInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateNodeInfo(&_Store.TransactOpts, nodeInfo)
}

// UpdateNodeInfo is a paid mutator transaction binding the contract method 0xc819d86c.
//
// Solidity: function UpdateNodeInfo((uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo) payable returns()
func (_Store *StoreTransactorSession) UpdateNodeInfo(nodeInfo NodeInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateNodeInfo(&_Store.TransactOpts, nodeInfo)
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

// UpdateSectorInfo is a paid mutator transaction binding the contract method 0x2384a6aa.
//
// Solidity: function UpdateSectorInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) _sector) payable returns()
func (_Store *StoreTransactor) UpdateSectorInfo(opts *bind.TransactOpts, _sector SectorInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateSectorInfo", _sector)
}

// UpdateSectorInfo is a paid mutator transaction binding the contract method 0x2384a6aa.
//
// Solidity: function UpdateSectorInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) _sector) payable returns()
func (_Store *StoreSession) UpdateSectorInfo(_sector SectorInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateSectorInfo(&_Store.TransactOpts, _sector)
}

// UpdateSectorInfo is a paid mutator transaction binding the contract method 0x2384a6aa.
//
// Solidity: function UpdateSectorInfo((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]) _sector) payable returns()
func (_Store *StoreTransactorSession) UpdateSectorInfo(_sector SectorInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateSectorInfo(&_Store.TransactOpts, _sector)
}

// UpdateUserSpace is a paid mutator transaction binding the contract method 0xed108de9.
//
// Solidity: function UpdateUserSpace(address walletAddr, (uint64,uint64,uint64,uint256,uint256) _userSpace) payable returns()
func (_Store *StoreTransactor) UpdateUserSpace(opts *bind.TransactOpts, walletAddr common.Address, _userSpace UserSpace) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateUserSpace", walletAddr, _userSpace)
}

// UpdateUserSpace is a paid mutator transaction binding the contract method 0xed108de9.
//
// Solidity: function UpdateUserSpace(address walletAddr, (uint64,uint64,uint64,uint256,uint256) _userSpace) payable returns()
func (_Store *StoreSession) UpdateUserSpace(walletAddr common.Address, _userSpace UserSpace) (*types.Transaction, error) {
	return _Store.Contract.UpdateUserSpace(&_Store.TransactOpts, walletAddr, _userSpace)
}

// UpdateUserSpace is a paid mutator transaction binding the contract method 0xed108de9.
//
// Solidity: function UpdateUserSpace(address walletAddr, (uint64,uint64,uint64,uint256,uint256) _userSpace) payable returns()
func (_Store *StoreTransactorSession) UpdateUserSpace(walletAddr common.Address, _userSpace UserSpace) (*types.Transaction, error) {
	return _Store.Contract.UpdateUserSpace(&_Store.TransactOpts, walletAddr, _userSpace)
}

// WhiteListOperate is a paid mutator transaction binding the contract method 0x35deb19c.
//
// Solidity: function WhiteListOperate((bytes,uint8,(address,uint64,uint64)[]) params) returns()
func (_Store *StoreTransactor) WhiteListOperate(opts *bind.TransactOpts, params WhiteListParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "WhiteListOperate", params)
}

// WhiteListOperate is a paid mutator transaction binding the contract method 0x35deb19c.
//
// Solidity: function WhiteListOperate((bytes,uint8,(address,uint64,uint64)[]) params) returns()
func (_Store *StoreSession) WhiteListOperate(params WhiteListParams) (*types.Transaction, error) {
	return _Store.Contract.WhiteListOperate(&_Store.TransactOpts, params)
}

// WhiteListOperate is a paid mutator transaction binding the contract method 0x35deb19c.
//
// Solidity: function WhiteListOperate((bytes,uint8,(address,uint64,uint64)[]) params) returns()
func (_Store *StoreTransactorSession) WhiteListOperate(params WhiteListParams) (*types.Transaction, error) {
	return _Store.Contract.WhiteListOperate(&_Store.TransactOpts, params)
}

// WithDrawProfit is a paid mutator transaction binding the contract method 0x9260665f.
//
// Solidity: function WithDrawProfit(address walletAddr) returns()
func (_Store *StoreTransactor) WithDrawProfit(opts *bind.TransactOpts, walletAddr common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "WithDrawProfit", walletAddr)
}

// WithDrawProfit is a paid mutator transaction binding the contract method 0x9260665f.
//
// Solidity: function WithDrawProfit(address walletAddr) returns()
func (_Store *StoreSession) WithDrawProfit(walletAddr common.Address) (*types.Transaction, error) {
	return _Store.Contract.WithDrawProfit(&_Store.TransactOpts, walletAddr)
}

// WithDrawProfit is a paid mutator transaction binding the contract method 0x9260665f.
//
// Solidity: function WithDrawProfit(address walletAddr) returns()
func (_Store *StoreTransactorSession) WithDrawProfit(walletAddr common.Address) (*types.Transaction, error) {
	return _Store.Contract.WithDrawProfit(&_Store.TransactOpts, walletAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address _config, address _file, address _list, address _node, address _pdp, address _space, address _sector, address _prove) returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts, _config common.Address, _file common.Address, _list common.Address, _node common.Address, _pdp common.Address, _space common.Address, _sector common.Address, _prove common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize", _config, _file, _list, _node, _pdp, _space, _sector, _prove)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address _config, address _file, address _list, address _node, address _pdp, address _space, address _sector, address _prove) returns()
func (_Store *StoreSession) Initialize(_config common.Address, _file common.Address, _list common.Address, _node common.Address, _pdp common.Address, _space common.Address, _sector common.Address, _prove common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _file, _list, _node, _pdp, _space, _sector, _prove)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address _config, address _file, address _list, address _node, address _pdp, address _space, address _sector, address _prove) returns()
func (_Store *StoreTransactorSession) Initialize(_config common.Address, _file common.Address, _list common.Address, _node common.Address, _pdp common.Address, _space common.Address, _sector common.Address, _prove common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _file, _list, _node, _pdp, _space, _sector, _prove)
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
