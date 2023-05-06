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

// GenChallengeParams is an auto generated low-level Go binding around an user-defined struct.
type GenChallengeParams struct {
	WalletAddr   common.Address
	HashValue    [32]byte
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

// PdpVerificationReturns is an auto generated low-level Go binding around an user-defined struct.
type PdpVerificationReturns struct {
	FileIDs     [][]byte
	Tags        [][]byte
	UpdatedChal []Challenge
	Path        []MerklePath
	RootHashes  [][]byte
	FileInfo    FileInfo
	Error       string
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

// ProofParams is an auto generated low-level Go binding around an user-defined struct.
type ProofParams struct {
	Version    uint64
	Proofs     []byte
	FileIds    [][]byte
	Tags       [][]byte
	RootHashes [][]byte
}

// ProofRecord is an auto generated low-level Go binding around an user-defined struct.
type ProofRecord struct {
	Proof            ProofParams
	State            bool
	LastUpdateHeight *big.Int
}

// ProofRecordWithParams is an auto generated low-level Go binding around an user-defined struct.
type ProofRecordWithParams struct {
	Proof      ProofParams
	Challenge  []Challenge
	MerklePath []MerklePath
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
	Tags         [][]byte
	MerklePath   []MerklePath
	PlotData     []byte
}

// TransferState is an auto generated low-level Go binding around an user-defined struct.
type TransferState struct {
	From  common.Address
	To    common.Address
	Value uint64
}

// VerifyPlotDataParams is an auto generated low-level Go binding around an user-defined struct.
type VerifyPlotDataParams struct {
	PlotInfo PlotInfo
	PlotData []byte
	Index    uint64
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"port\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"}],\"name\":\"DNSNodeRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DNSNodeUnReg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"DnsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"FsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structNameInfo\",\"name\":\"newer\",\"type\":\"tuple\"}],\"name\":\"NotifyNameInfoAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"PDPVerifyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nodeAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"VerifyProofWithMerklePathForFileEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"HashValue\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveNum\",\"type\":\"uint64\"}],\"internalType\":\"structGenChallengeParams\",\"name\":\"gParams\",\"type\":\"tuple\"}],\"name\":\"GenChallenge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge\",\"name\":\"chg\",\"type\":\"tuple\"}],\"name\":\"GetChallengeKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"GetChallengeList\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"FileIds\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"RootHashes\",\"type\":\"bytes[]\"}],\"internalType\":\"structProofParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"GetKeyByProofParams\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode\",\"name\":\"mn\",\"type\":\"tuple\"}],\"name\":\"GetMerkleNodeKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath\",\"name\":\"mp\",\"type\":\"tuple\"}],\"name\":\"GetMerklePathKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"GetMerklePathList\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetUnVerifyProofList\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"FileIds\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"RootHashes\",\"type\":\"bytes[]\"}],\"internalType\":\"structProofParams\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"challenge\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"merklePath\",\"type\":\"tuple[]\"}],\"internalType\":\"structProofRecordWithParams[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"SectorInfo_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"Challenges\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"ProveFileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"BlockNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"MerklePath_\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"PlotData\",\"type\":\"bytes\"}],\"internalType\":\"structSectorProveData\",\"name\":\"ProveData\",\"type\":\"tuple\"}],\"internalType\":\"structPrepareForPdpVerificationParams\",\"name\":\"pParams\",\"type\":\"tuple\"}],\"name\":\"PrepareForPdpVerification\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"FileIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"UpdatedChal\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"Path\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"RootHashes\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"RootHash\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"FileID\",\"type\":\"bytes\"}],\"internalType\":\"structFileProveParam\",\"name\":\"FileProveParam_\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"FileInfo_\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"Error\",\"type\":\"string\"}],\"internalType\":\"structPdpVerificationReturns\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"chgs\",\"type\":\"tuple[]\"}],\"name\":\"SaveChallenge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"mp\",\"type\":\"tuple[]\"}],\"name\":\"SaveMerklePath\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"FileIds\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"RootHashes\",\"type\":\"bytes[]\"}],\"internalType\":\"structProofParams\",\"name\":\"vParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"chgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"mp\",\"type\":\"tuple[]\"}],\"name\":\"SubmitVerifyProofRequest\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"PlotData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"}],\"internalType\":\"structVerifyPlotDataParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"VerifyPlotData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"FileIds\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"RootHashes\",\"type\":\"bytes[]\"}],\"internalType\":\"structProofParams\",\"name\":\"Proof\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"State\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"LastUpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structProofRecord\",\"name\":\"vParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"chgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"mp\",\"type\":\"tuple[]\"}],\"name\":\"VerifyProof\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"FileIds\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"RootHashes\",\"type\":\"bytes[]\"}],\"internalType\":\"structProofParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"VerifyProofWithMerklePathForFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"bytesToUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"HashValue\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveNum\",\"type\":\"uint64\"}],\"internalType\":\"structGenChallengeParams\",\"name\":\"gParams\",\"type\":\"tuple\"}],\"name\":\"genChallengeKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFile\",\"name\":\"_file\",\"type\":\"address\"},{\"internalType\":\"contractISector\",\"name\":\"_sector\",\"type\":\"address\"},{\"internalType\":\"contractPDPExtra\",\"name\":\"_pdpExtra\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615c8180620000216000396000f3fe6080604052600436106101295760003560e01c80638b90b4cd116100a5578063c0c53b8b11610074578063d48841c011610059578063d48841c014610333578063efea778314610353578063f5dbc78e1461037357600080fd5b8063c0c53b8b14610300578063c70a45db1461032057600080fd5b80638b90b4cd14610273578063904a8d4114610286578063b750ab8a146102b3578063b8527b31146102d357600080fd5b806344c2b91b116100fc5780635bf7f1d2116100e15780635bf7f1d214610220578063743b9eb61461024057806383807a431461025357600080fd5b806344c2b91b146101de57806357d60d6e146101f357600080fd5b806302d06d051461012e5780632e19aeff146101645780633c279f5d146101915780633d1731b8146101be575b600080fd5b34801561013a57600080fd5b5061014e6101493660046145a9565b610395565b60405161015b9190615679565b60405180910390f35b34801561017057600080fd5b5061018461017f366004614954565b610428565b60405161015b9190615552565b34801561019d57600080fd5b506101b16101ac366004614774565b610448565b60405161015b9190615560565b3480156101ca57600080fd5b506101b16101d93660046147c6565b61048a565b6101f16101ec3660046145dd565b610544565b005b3480156101ff57600080fd5b5061021361020e36600461482e565b610602565b60405161015b9190615643565b34801561022c57600080fd5b5061018461023b366004614862565b610c58565b6101f161024e366004614644565b610ffe565b34801561025f57600080fd5b506101b161026e366004614722565b611285565b6101f1610281366004614896565b6112a4565b34801561029257600080fd5b506102a66102a13660046145a9565b611364565b60405161015b919061551f565b3480156102bf57600080fd5b506101b16102ce366004614792565b6114dc565b3480156102df57600080fd5b506102f36102ee3660046145a9565b6114fb565b60405161015b9190615530565b34801561030c57600080fd5b506101f161031b3660046146a1565b61185c565b6101f161032e36600461491d565b6119c8565b34801561033f57600080fd5b506102a661034e366004614774565b611aa3565b34801561035f57600080fd5b506101b161036e366004614862565b611fea565b34801561037f57600080fd5b506103886121de565b60405161015b9190615541565b60008060005b8351811015610421576103af816001615738565b84516103bb919061595d565b6103c6906008615907565b6103d1906002615837565b8482815181106103f157634e487b7160e01b600052603260045260246000fd5b0160200151610403919060f81c615907565b61040d9083615738565b91508061041981615a80565b91505061039b565b5092915050565b600061044082604001516001600160401b031661240c565b506001919050565b606060008260000151836020015184604001518560600151604051602001610473949392919061540b565b60408051601f198184030181529190529392505050565b60608060005b8360200151518163ffffffff1610156104215783602001518163ffffffff16815181106104cd57634e487b7160e01b600052603260045260246000fd5b60200260200101516040015184602001518263ffffffff168151811061050357634e487b7160e01b600052603260045260246000fd5b60200260200101516020015160405160200161052092919061548c565b6040516020818303038152906040529150808061053c90615a9b565b915050610490565b60005b81518163ffffffff1610156105fd57600061058e838363ffffffff168151811061058157634e487b7160e01b600052603260045260246000fd5b6020026020010151611285565b90506105e881848463ffffffff16815181106105ba57634e487b7160e01b600052603260045260246000fd5b60200260200101516006876040516105d29190615468565b9081526040519081900360200190209190612480565b505080806105f590615a9b565b915050610547565b505050565b61060a61313a565b61061261313a565b600254835160408086015190517f19ac473f0000000000000000000000000000000000000000000000000000000081526000936001600160a01b0316926319ac473f9261066192600401615654565b60006040518083038186803b15801561067957600080fd5b505afa15801561068d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106b591908101906146ee565b8051909150156106ca5760c082015292915050565b8351610100015160006001600160401b03808316908111156106fc57634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561074257816020015b60408051808201909152606081526000602082015281526020019060019003908161071a5790505b50905060005b826001600160401b0316816001600160401b031610156108d857600087600001516101600151826001600160401b03168151811061079657634e487b7160e01b600052603260045260246000fd5b602002602001015190508083836001600160401b0316815181106107ca57634e487b7160e01b600052603260045260246000fd5b602090810291909101015152600080546040517fdefce5d4000000000000000000000000000000000000000000000000000000008152620100009091046001600160a01b03169063defce5d490610825908590600401615560565b60006040518083038186803b15801561083d57600080fd5b505afa158015610851573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526108799190810190614740565b9050806080015184846001600160401b0316815181106108a957634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160401b03909216910152508190506108d081615aba565b915050610748565b50816001600160401b03166001600160401b0381111561090857634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561093b57816020015b60608152602001906001900390816109265790505b5084526001600160401b038083169081111561096757634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561099a57816020015b60608152602001906001900390816109855790505b5060208501526001600160401b03808316908111156109c957634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610a0f57816020015b6040805180820190915260008152606060208201528152602001906001900390816109e75790505b5060608501526001600160401b0380831690811115610a3e57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610a7157816020015b6060815260200190600190039081610a5c5790505b5060808501526001600160401b0380831690811115610aa057634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610ae557816020015b6040805180820190915260008082526020820152815260200190600190039081610abe5790505b50604080860191909152600254600054602089015192517f83095ce20000000000000000000000000000000000000000000000000000000081526001600160a01b03928316936383095ce293610b4b936201000090041691879187918b90600401615571565b60006040518083038186803b158015610b6357600080fd5b505afa158015610b77573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610b9f91908101906147fa565b60025460208801516040808a015190517f799b76720000000000000000000000000000000000000000000000000000000081529397506001600160a01b039092169263799b767292610bfa9287928792908b90600401615687565b60006040518083038186803b158015610c1257600080fd5b505afa158015610c26573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c4e91908101906147fa565b9695505050505050565b600080610c6483611fea565b90506000600360000182604051610c7b9190615468565b9081526040805191829003602001822061010083019091526001810180546001600160401b031660608401908152600290920180549192849290918491608085019190610cc790615a2d565b80601f0160208091040260200160405190810160405280929190818152602001828054610cf390615a2d565b8015610d405780601f10610d1557610100808354040283529160200191610d40565b820191906000526020600020905b815481529060010190602001808311610d2357829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b82821015610e1a578382906000526020600020018054610d8d90615a2d565b80601f0160208091040260200160405190810160405280929190818152602001828054610db990615a2d565b8015610e065780601f10610ddb57610100808354040283529160200191610e06565b820191906000526020600020905b815481529060010190602001808311610de957829003601f168201915b505050505081526020019060010190610d6e565b50505050815260200160038201805480602002602001604051908101604052809291908181526020016000905b82821015610ef3578382906000526020600020018054610e6690615a2d565b80601f0160208091040260200160405190810160405280929190818152602001828054610e9290615a2d565b8015610edf5780601f10610eb457610100808354040283529160200191610edf565b820191906000526020600020905b815481529060010190602001808311610ec257829003601f168201915b505050505081526020019060010190610e47565b50505050815260200160048201805480602002602001604051908101604052809291908181526020016000905b82821015610fcc578382906000526020600020018054610f3f90615a2d565b80601f0160208091040260200160405190810160405280929190818152602001828054610f6b90615a2d565b8015610fb85780601f10610f8d57610100808354040283529160200191610fb8565b820191906000526020600020905b815481529060010190602001808311610f9b57829003601f168201915b505050505081526020019060010190610f20565b505050915250508152600582015460ff1615156020808301919091526006909201546040909101520151949350505050565b60005b81518163ffffffff1610156105fd576000611048838363ffffffff168151811061103b57634e487b7160e01b600052603260045260246000fd5b602002602001015161048a565b90506000838363ffffffff168151811061107257634e487b7160e01b600052603260045260246000fd5b602002602001015160200151516001600160401b038111156110a457634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156110d757816020015b60608152602001906001900390816110c25790505b50905060005b848463ffffffff168151811061110357634e487b7160e01b600052603260045260246000fd5b602002602001015160200151518163ffffffff161015611243576000611186868663ffffffff168151811061114857634e487b7160e01b600052603260045260246000fd5b6020026020010151602001518363ffffffff168151811061117957634e487b7160e01b600052603260045260246000fd5b60200260200101516114dc565b90506111fc81878763ffffffff16815181106111b257634e487b7160e01b600052603260045260246000fd5b6020026020010151602001518463ffffffff16815181106111e357634e487b7160e01b600052603260045260246000fd5b602002602001015160086125c09092919063ffffffff16565b5080838363ffffffff168151811061122457634e487b7160e01b600052603260045260246000fd5b602002602001018190525050808061123b90615a9b565b9150506110dd565b5061126f82826007886040516112599190615468565b9081526040519081900360200190209190612689565b505050808061127d90615a9b565b915050611001565b60606000826000015183602001516040516020016104739291906154ae565b60006112af84611fea565b90506112bb8184610544565b6112c58183610ffe565b61130e604080516101008101909152600060608083019182526080830181905260a0830181905260c0830181905260e08301528190815260006020820181905260409091015290565b84815260006020820152611324600383836126f4565b507f7f6b6a2262c2fc2cc1b1785be448eae36656117d8f558facfe3c199e26256ea1600a60405161135591906155c6565b60405180910390a15050505050565b606060006006836040516113789190615468565b908152604051908190036020019020600101546001600160401b038111156113b057634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156113f557816020015b60408051808201909152600080825260208201528152602001906001900390816113ce5790505b509050600061142060068560405161140d9190615468565b9081526020016040518091039020612808565b90505b61144c816006866040516114379190615468565b90815260405190819003602001902090612823565b1561042157600061147c826006876040516114679190615468565b90815260405190819003602001902090612831565b915050808383815181106114a057634e487b7160e01b600052603260045260246000fd5b6020026020010181905250506114d5816006866040516114c09190615468565b9081526040519081900360200190209061295a565b9050611423565b606060008260400151836020015160405160200161047392919061548c565b6060600060078360405161150f9190615468565b908152604051908190036020019020600101546001600160401b0381111561154757634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561158d57816020015b6040805180820190915260008152606060208201528152602001906001900390816115655790505b50905060006115b86007856040516115a59190615468565b90815260200160405180910390206129d0565b90505b6115cf816007866040516114379190615468565b156104215760006115ff826007876040516115ea9190615468565b908152604051908190036020019020906129de565b91505080516001600160401b0381111561162957634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561167657816020015b604080516060808201835260008083526020830152918101919091528152602001906001900390816116475790505b5083838151811061169757634e487b7160e01b600052603260045260246000fd5b60200260200101516020018190525060005b815181101561182a5760086000018282815181106116d757634e487b7160e01b600052603260045260246000fd5b60200260200101516040516116ec9190615468565b9081526040805191829003602090810183206060840183526001810180546001600160401b0380821687526801000000000000000090910416928501929092526002018054919284019161173f90615a2d565b80601f016020809104026020016040519081016040528092919081815260200182805461176b90615a2d565b80156117b85780601f1061178d576101008083540402835291602001916117b8565b820191906000526020600020905b81548152906001019060200180831161179b57829003601f168201915b5050505050815250508484815181106117e157634e487b7160e01b600052603260045260246000fd5b602002602001015160200151828151811061180c57634e487b7160e01b600052603260045260246000fd5b6020026020010181905250808061182290615a80565b9150506116a9565b5050611855816007866040516118409190615468565b90815260405190819003602001902090612b9c565b90506115bb565b600054610100900460ff161580801561187c5750600054600160ff909116105b806118965750303b158015611896575060005460ff166001145b6118d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118cc906155e2565b60405180910390fd5b6000805460ff1916600117905580156118f8576000805461ff0019166101001790555b600080546001600160a01b0380871662010000027fffffffffffffffffffff0000000000000000000000000000000000000000ffff90921691909117909155600180548583167fffffffffffffffffffffffff000000000000000000000000000000000000000091821617909155600280549285169290911691909117905580156119c2576000805461ff00191690556040517f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906119b9906001906155d4565b60405180910390a15b50505050565b611a11604080516101008101909152600060608083019182526080830181905260a0830181905260c0830181905260e08301528190815260006020820181905260409091015290565b83515181516001600160401b0390911690528351602090810151825182015284516040908101518351820152855160609081015184519091015285516080908101518451909101528186015115159183019190915243908201528351600090611a7990611fea565b9050611a858185610544565b611a8f8184610ffe565b611a9b600382846126f4565b505050505050565b6020808201518251604051606093600092611ac29290918591016153e5565b60405160208183030381529060405290506000600282604051611ae59190615468565b602060405180830381855afa158015611b02573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190611b25919061458b565b604080516024808252606082019092529192506000919060208201818036833701905050905060005b60208163ffffffff161015611bd657828163ffffffff1660208110611b8357634e487b7160e01b600052603260045260246000fd5b1a60f81b828263ffffffff1681518110611bad57634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a90535080611bce81615a9b565b915050611b4e565b5060005b60048163ffffffff161015611c6c57828163ffffffff1660208110611c0f57634e487b7160e01b600052603260045260246000fd5b1a60f81b82611c1f836020615750565b63ffffffff1681518110611c4357634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a90535080611c6481615a9b565b915050611bda565b506000806000600389604001516001600160401b031611611ca65750505060408601516001600160401b0316606087015260018080611d23565b600389604001516001600160401b0316118015611cdc575088606001516001600160401b031689604001516001600160401b0316105b15611ce957600360608a01525b88606001518960400151611cfd91906157ce565b925088606001518960400151611d139190615b30565b611d1d908461577a565b91508290505b600089606001516001600160401b03166001600160401b03811115611d5857634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015611d9d57816020015b6040805180820190915260008082526020820152815260200190600190039081611d765790505b50905085600060015b8c606001516001600160401b03168163ffffffff1611611fd9578c606001516001600160401b03168163ffffffff161415611ddf578594505b60408051600480825281830190925260009160208201818036833701905050905060005b60048163ffffffff161015611e9b5789611e1d8286615750565b63ffffffff1681518110611e4157634e487b7160e01b600052603260045260246000fd5b602001015160f81c60f81b828263ffffffff1681518110611e7257634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a90535080611e9381615a9b565b915050611e03565b506000611ea782610395565b905088611eb5600185615978565b63ffffffff16611ec59190615926565b87611ed1836001615750565b63ffffffff16611ee19190615b30565b611eeb919061577a565b86611ef7600186615978565b63ffffffff1681518110611f1b57634e487b7160e01b600052603260045260246000fd5b60209081029190910181015163ffffffff928316905286918616908110611f5257634e487b7160e01b600052603260045260246000fd5b611f5f91901a60016157ad565b60ff1686611f6e600186615978565b63ffffffff1681518110611f9257634e487b7160e01b600052603260045260246000fd5b60209081029190910181015163ffffffff90921691015283611fb381615a9b565b9450611fc29050602085615b09565b935050508080611fd190615a9b565b915050611da6565b50919b9a5050505050505050505050565b60608060005b8360400151518163ffffffff16101561206b578184604001518263ffffffff168151811061202e57634e487b7160e01b600052603260045260246000fd5b6020026020010151604051602001612047929190615474565b6040516020818303038152906040529150808061206390615a9b565b915050611ff0565b50606060005b8460600151518163ffffffff1610156120ec578185606001518263ffffffff16815181106120af57634e487b7160e01b600052603260045260246000fd5b60200260200101516040516020016120c8929190615474565b604051602081830303815290604052915080806120e490615a9b565b915050612071565b50606060005b8560800151518163ffffffff16101561216d578186608001518263ffffffff168151811061213057634e487b7160e01b600052603260045260246000fd5b6020026020010151604051602001612149929190615474565b6040516020818303038152906040529150808061216590615a9b565b9150506120f2565b506000856000015186602001518585856040516020016121919594939291906154d4565b60408051601f1981840301815290829052805160208083019190912091935083926000916121c191849101615453565b60408051601f198184030181529190529998505050505050505050565b60606000806121ed6003612c0c565b90505b6121fb600382612823565b1561223e57600061220d600383612c1a565b91505060008160400151111561222b578261222781615a80565b9350505b506122376003826130a9565b90506121f0565b506005546000906001600160401b0381111561226a57634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156122e457816020015b6122d1604080516101008101909152600060608083019182526080830181905260a0830181905260c0830181905260e08301528190815260200160608152602001606081525090565b8152602001906001900390816122885790505b5090506123226040518060a0016040528060006001600160401b03168152602001606081526020016060815260200160608152602001606081525090565b6060806000806123326003612c0c565b90505b612340600382612823565b15612400576000612352600383612c1a565b915050600081604001511161236757506123ee565b80519550600061237687611fea565b905061238181611364565b955061238c816114fb565b945060405180606001604052808881526020018781526020018681525088856001600160401b0316815181106123d257634e487b7160e01b600052603260045260246000fd5b602002602001018190525083806123e890615aba565b94505050505b6123f96003826130a9565b9050612335565b50939695505050505050565b61247d816040516024016124209190615679565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff82c50f100000000000000000000000000000000000000000000000000000000179052613119565b50565b60008084600001846040516124959190615468565b908152604051908190036020018120549150839086906124b6908790615468565b908152604051908190036020908101909120825160019091018054939092015163ffffffff9081166401000000000267ffffffffffffffff19909416911617919091179055801561250b5760019150506125b9565b506001808501805480830182556000919091529061252a908290615738565b604051869061253a908790615468565b908152604051908190036020019020556001850180548591908390811061257157634e487b7160e01b600052603260045260246000fd5b9060005260206000209060020201600001908051906020019061259592919061325a565b506002850180549060006125a883615a80565b919050555060009150506125b9565b505b9392505050565b60008084600001846040516125d59190615468565b908152604051908190036020018120549150839086906125f6908790615468565b908152604080519182900360209081019092208351600182018054868601516001600160401b0390811668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000000000000000000090921693169290921791909117815591840151805192936126759360029093019291019061325a565b50508115905061250b5760019150506125b9565b600080846000018460405161269e9190615468565b908152604051908190036020018120549150839086906126bf908790615468565b908152602001604051809103902060010190805190602001906126e39291906132de565b50801561250b5760019150506125b9565b60008084600001846040516127099190615468565b9081526040519081900360200181205491508390869061272a908790615468565b908152604051602091819003820190208251805160018301805467ffffffffffffffff19166001600160401b0390921691909117815581840151805191949293859361277d93600290920192019061325a565b50604082015180516127999160028401916020909101906132de565b50606082015180516127b59160038401916020909101906132de565b50608082015180516127d19160048401916020909101906132de565b505050602082015160058201805460ff1916911515919091179055604090910151600690910155801561250b5760019150506125b9565b60008061281683600061295a565b90506125b960018261595d565b600182015481105b92915050565b604080518082019091526000808252602082015260609083600101838154811061286b57634e487b7160e01b600052603260045260246000fd5b9060005260206000209060020201600001805461288790615a2d565b80601f01602080910402602001604051908101604052809291908181526020018280546128b390615a2d565b80156129005780601f106128d557610100808354040283529160200191612900565b820191906000526020600020905b8154815290600101906020018083116128e357829003601f168201915b5050505050915083600001826040516129199190615468565b90815260408051918290036020908101832083830190925260019091015463ffffffff80821684526401000000009091041690820152919491935090915050565b60008161296681615a80565b9250505b6001830154821080156129b3575082600101828154811061299b57634e487b7160e01b600052603260045260246000fd5b600091825260209091206001600290920201015460ff165b156129ca57816129c281615a80565b92505061296a565b50919050565b600080612816836000612b9c565b606080836001018381548110612a0457634e487b7160e01b600052603260045260246000fd5b90600052602060002090600202016000018054612a2090615a2d565b80601f0160208091040260200160405190810160405280929190818152602001828054612a4c90615a2d565b8015612a995780601f10612a6e57610100808354040283529160200191612a99565b820191906000526020600020905b815481529060010190602001808311612a7c57829003601f168201915b505050505091508360000182604051612ab29190615468565b9081526020016040518091039020600101805480602002602001604051908101604052809291908181526020016000905b82821015612b8f578382906000526020600020018054612b0290615a2d565b80601f0160208091040260200160405190810160405280929190818152602001828054612b2e90615a2d565b8015612b7b5780601f10612b5057610100808354040283529160200191612b7b565b820191906000526020600020905b815481529060010190602001808311612b5e57829003601f168201915b505050505081526020019060010190612ae3565b5050505090509250929050565b600081612ba881615a80565b9250505b600183015482108015612bf55750826001018281548110612bdd57634e487b7160e01b600052603260045260246000fd5b600091825260209091206001600290920201015460ff165b156129ca5781612c0481615a80565b925050612bac565b6000806128168360006130a9565b6040805161010081018252600060608281018281526080840182905260a0840182905260c0840182905260e0840182905283526020830182905292820152836001018381548110612c7b57634e487b7160e01b600052603260045260246000fd5b90600052602060002090600202016000018054612c9790615a2d565b80601f0160208091040260200160405190810160405280929190818152602001828054612cc390615a2d565b8015612d105780601f10612ce557610100808354040283529160200191612d10565b820191906000526020600020905b815481529060010190602001808311612cf357829003601f168201915b505050505091508360000182604051612d299190615468565b9081526040805191829003602001822061010083019091526001810180546001600160401b031660608401908152600290920180549192849290918491608085019190612d7590615a2d565b80601f0160208091040260200160405190810160405280929190818152602001828054612da190615a2d565b8015612dee5780601f10612dc357610100808354040283529160200191612dee565b820191906000526020600020905b815481529060010190602001808311612dd157829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b82821015612ec8578382906000526020600020018054612e3b90615a2d565b80601f0160208091040260200160405190810160405280929190818152602001828054612e6790615a2d565b8015612eb45780601f10612e8957610100808354040283529160200191612eb4565b820191906000526020600020905b815481529060010190602001808311612e9757829003601f168201915b505050505081526020019060010190612e1c565b50505050815260200160038201805480602002602001604051908101604052809291908181526020016000905b82821015612fa1578382906000526020600020018054612f1490615a2d565b80601f0160208091040260200160405190810160405280929190818152602001828054612f4090615a2d565b8015612f8d5780601f10612f6257610100808354040283529160200191612f8d565b820191906000526020600020905b815481529060010190602001808311612f7057829003601f168201915b505050505081526020019060010190612ef5565b50505050815260200160048201805480602002602001604051908101604052809291908181526020016000905b8282101561307a578382906000526020600020018054612fed90615a2d565b80601f016020809104026020016040519081016040528092919081815260200182805461301990615a2d565b80156130665780601f1061303b57610100808354040283529160200191613066565b820191906000526020600020905b81548152906001019060200180831161304957829003601f168201915b505050505081526020019060010190612fce565b505050915250508152600582015460ff1615156020820152600690910154604090910152919491935090915050565b6000816130b581615a80565b9250505b60018301548210801561310257508260010182815481106130ea57634e487b7160e01b600052603260045260246000fd5b600091825260209091206001600290920201015460ff165b156129ca578161311181615a80565b9250506130b9565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b6040518060e00160405280606081526020016060815260200160608152602001606081526020016060815260200161324d604080516102e0810182526060808252600060208084018290528385018390528284018290526080840182905260a0840182905260c0840182905260e084018290526101008401829052610120840182905261014084018290528451808601865283815280820184905261016085015261018084018290526101a084018290526101c084018290526101e08401829052610200840182905261022084018390526102408401839052610260840183905261028084018290526102a08401829052845192830185528183528201819052928101929092526102c081019190915290565b8152602001606081525090565b82805461326690615a2d565b90600052602060002090601f01602090048101928261328857600085556132ce565b82601f106132a157805160ff19168380011785556132ce565b828001600101855582156132ce579182015b828111156132ce5782518255916020019190600101906132b3565b506132da929150613337565b5090565b82805482825590600052602060002090810192821561332b579160200282015b8281111561332b578251805161331b91849160209091019061325a565b50916020019190600101906132fe565b506132da92915061334c565b5b808211156132da5760008155600101613338565b808211156132da5760006133608282613369565b5060010161334c565b50805461337590615a2d565b6000825580601f10613385575050565b601f01602090049060005260206000209081019061247d9190613337565b60006133b66133b1846156eb565b6156cf565b905080838252602082019050828560208602820111156133d557600080fd5b60005b8581101561340157816133eb88826137d4565b84525060209283019291909101906001016133d8565b5050509392505050565b60006134196133b1846156eb565b9050808382526020820190508285602086028201111561343857600080fd5b60005b858110156134015781356001600160401b0381111561345957600080fd5b808601613466898261393c565b85525050602092830192919091019060010161343b565b600061348b6133b1846156eb565b905080838252602082019050828560208602820111156134aa57600080fd5b60005b858110156134015781516001600160401b038111156134cb57600080fd5b8086016134d8898261395d565b8552505060209283019291909101906001016134ad565b60006134fd6133b1846156eb565b8381529050602081018260408502810186101561351957600080fd5b60005b85811015613401578161352f88826139aa565b8452506020909201916040919091019060010161351c565b60006135556133b1846156eb565b8381529050602081018260408502810186101561357157600080fd5b60005b85811015613401578161358788826139f1565b84525060209092019160409190910190600101613574565b60006135ad6133b1846156eb565b905080838252602082019050828560208602820111156135cc57600080fd5b60005b858110156134015781356001600160401b038111156135ed57600080fd5b8086016135fa8982613d99565b8552505060209283019291909101906001016135cf565b600061361f6133b1846156eb565b9050808382526020820190508285602086028201111561363e57600080fd5b60005b858110156134015781516001600160401b0381111561365f57600080fd5b80860161366c8982613e0b565b855250506020928301929190910190600101613641565b60006136916133b1846156eb565b905080838252602082019050828560208602820111156136b057600080fd5b60005b858110156134015781356001600160401b038111156136d157600080fd5b8086016136de8982613e71565b8552505060209283019291909101906001016136b3565b60006137036133b1846156eb565b9050808382526020820190508285602086028201111561372257600080fd5b60005b858110156134015781516001600160401b0381111561374357600080fd5b8086016137508982613ec3565b855250506020928301929190910190600101613725565b60006137756133b18461570e565b90508281526020810184848401111561378d57600080fd5b6125b78482856159f5565b60006137a66133b18461570e565b9050828152602081018484840111156137be57600080fd5b6125b7848285615a01565b803561282b81615beb565b805161282b81615beb565b600082601f8301126137f057600080fd5b81516138008482602086016133a3565b949350505050565b600082601f83011261381957600080fd5b813561380084826020860161340b565b600082601f83011261383a57600080fd5b815161380084826020860161347d565b600082601f83011261385b57600080fd5b81356138008482602086016134ef565b600082601f83011261387c57600080fd5b8151613800848260208601613547565b600082601f83011261389d57600080fd5b813561380084826020860161359f565b600082601f8301126138be57600080fd5b8151613800848260208601613611565b600082601f8301126138df57600080fd5b8135613800848260208601613683565b600082601f83011261390057600080fd5b81516138008482602086016136f5565b803561282b81615bff565b805161282b81615bff565b803561282b81615c07565b805161282b81615c07565b600082601f83011261394d57600080fd5b8135613800848260208601613767565b600082601f83011261396e57600080fd5b8151613800848260208601613798565b803561282b81615c0d565b803561282b81615c16565b805161282b81615c16565b805161282b81615c23565b6000604082840312156139bc57600080fd5b6139c660406156cf565b905060006139d4848461455f565b82525060206139e58484830161455f565b60208301525092915050565b600060408284031215613a0357600080fd5b613a0d60406156cf565b90506000613a1b848461456a565b82525060206139e58484830161456a565b60006103208284031215613a3f57600080fd5b613a4a6102e06156cf565b82519091506001600160401b03811115613a6357600080fd5b613a6f8482850161395d565b8252506020613a80848483016137d4565b60208301525060408201516001600160401b03811115613a9f57600080fd5b613aab8482850161395d565b6040830152506060613abf84828501614580565b6060830152506080613ad384828501614580565b60808301525060a0613ae784828501614580565b60a08301525060c0613afb84828501614580565b60c08301525060e0613b0f84828501614580565b60e083015250610100613b2484828501613931565b61010083015250610120613b3a84828501614580565b61012083015250610140613b5084828501614580565b610140830152506101608201516001600160401b03811115613b7157600080fd5b613b7d84828501613cc1565b61016083015250610180613b9384828501614580565b610180830152506101a0613ba984828501613931565b6101a0830152506101c0613bbf8482850161391b565b6101c0830152506101e0613bd58482850161399f565b6101e083015250610200613beb84828501614580565b610200830152506102208201516001600160401b03811115613c0c57600080fd5b613c18848285016137df565b610220830152506102408201516001600160401b03811115613c3957600080fd5b613c45848285016137df565b610240830152506102608201516001600160401b03811115613c6657600080fd5b613c728482850161395d565b61026083015250610280613c8884828501613994565b610280830152506102a0613c9e8482850161391b565b6102a0830152506102c0613cb4848285016140b0565b6102c08301525092915050565b600060408284031215613cd357600080fd5b613cdd60406156cf565b82519091506001600160401b03811115613cf657600080fd5b613d028482850161395d565b82525060208201516001600160401b03811115613d1e57600080fd5b6139e58482850161395d565b600060808284031215613d3c57600080fd5b613d4660806156cf565b90506000613d5484846137c9565b8252506020613d6584848301613926565b6020830152506040613d7984828501614575565b6040830152506060613d8d84828501614575565b60608301525092915050565b600060608284031215613dab57600080fd5b613db560606156cf565b90506000613dc38484614575565b8252506020613dd484848301614575565b60208301525060408201356001600160401b03811115613df357600080fd5b613dff8482850161393c565b60408301525092915050565b600060608284031215613e1d57600080fd5b613e2760606156cf565b90506000613e358484614580565b8252506020613e4684848301614580565b60208301525060408201516001600160401b03811115613e6557600080fd5b613dff8482850161395d565b600060408284031215613e8357600080fd5b613e8d60406156cf565b90506000613e9b8484614575565b82525060208201356001600160401b03811115613eb757600080fd5b6139e58482850161388c565b600060408284031215613ed557600080fd5b613edf60406156cf565b90506000613eed8484614580565b82525060208201516001600160401b03811115613f0957600080fd5b6139e5848285016138ad565b600060e08284031215613f2757600080fd5b613f3160e06156cf565b82519091506001600160401b03811115613f4a57600080fd5b613f5684828501613829565b82525060208201516001600160401b03811115613f7257600080fd5b613f7e84828501613829565b60208301525060408201516001600160401b03811115613f9d57600080fd5b613fa98482850161386b565b60408301525060608201516001600160401b03811115613fc857600080fd5b613fd4848285016138ef565b60608301525060808201516001600160401b03811115613ff357600080fd5b613fff84828501613829565b60808301525060a08201516001600160401b0381111561401e57600080fd5b61402a84828501613a2c565b60a08301525060c08201516001600160401b0381111561404957600080fd5b6140558482850161395d565b60c08301525092915050565b60006060828403121561407357600080fd5b61407d60606156cf565b9050600061408b8484614575565b825250602061409c84848301614575565b6020830152506040613dff84828501614575565b6000606082840312156140c257600080fd5b6140cc60606156cf565b905060006140da8484614580565b82525060206140eb84848301614580565b6020830152506040613dff84828501614580565b60006060828403121561411157600080fd5b61411b60606156cf565b905081356001600160401b0381111561413357600080fd5b61413f848285016142d6565b82525060208201356001600160401b0381111561415b57600080fd5b6141678482850161384a565b60208301525060408201356001600160401b0381111561418657600080fd5b613dff84828501614406565b600060a082840312156141a457600080fd5b6141ae60a06156cf565b905060006141bc8484614575565b82525060208201356001600160401b038111156141d857600080fd5b6141e48482850161393c565b60208301525060408201356001600160401b0381111561420357600080fd5b61420f84828501613808565b60408301525060608201356001600160401b0381111561422e57600080fd5b61423a84828501613808565b60608301525060808201356001600160401b0381111561425957600080fd5b61426584828501613808565b60808301525092915050565b60006060828403121561428357600080fd5b61428d60606156cf565b905081356001600160401b038111156142a557600080fd5b6142b184828501614192565b82525060206142c284848301613910565b6020830152506040613dff84828501613926565b600061018082840312156142e957600080fd5b6142f46101806156cf565b9050600061430284846137c9565b825250602061431384848301614575565b602083015250604061432784828501614575565b604083015250606061433b84828501614575565b606083015250608061434f84828501613989565b60808301525060a061436384828501613926565b60a08301525060c061437784828501613926565b60c08301525060e061438b84828501614575565b60e0830152506101006143a084828501614575565b610100830152506101206143b684828501614575565b610120830152506101406143cc84828501613910565b610140830152506101608201356001600160401b038111156143ed57600080fd5b6143f984828501613808565b6101608301525092915050565b600060c0828403121561441857600080fd5b61442260c06156cf565b905060006144308484614575565b825250602061444184848301614575565b60208301525060408201356001600160401b0381111561446057600080fd5b61446c8482850161393c565b60408301525060608201356001600160401b0381111561448b57600080fd5b61449784828501613808565b60608301525060808201356001600160401b038111156144b657600080fd5b6144c2848285016138ce565b60808301525060a08201356001600160401b038111156144e157600080fd5b6144ed8482850161393c565b60a08301525092915050565b600060a0828403121561450b57600080fd5b61451560606156cf565b905060006145238484614061565b82525060608201356001600160401b0381111561453f57600080fd5b61454b8482850161393c565b6020830152506080613dff84828501614575565b803561282b81615c30565b805161282b81615c30565b803561282b81615c3c565b805161282b81615c3c565b60006020828403121561459d57600080fd5b60006138008484613931565b6000602082840312156145bb57600080fd5b81356001600160401b038111156145d157600080fd5b6138008482850161393c565b600080604083850312156145f057600080fd5b82356001600160401b0381111561460657600080fd5b6146128582860161393c565b92505060208301356001600160401b0381111561462e57600080fd5b61463a8582860161384a565b9150509250929050565b6000806040838503121561465757600080fd5b82356001600160401b0381111561466d57600080fd5b6146798582860161393c565b92505060208301356001600160401b0381111561469557600080fd5b61463a858286016138ce565b6000806000606084860312156146b657600080fd5b60006146c2868661397e565b93505060206146d38682870161397e565b92505060406146e48682870161397e565b9150509250925092565b60006020828403121561470057600080fd5b81516001600160401b0381111561471657600080fd5b6138008482850161395d565b60006040828403121561473457600080fd5b600061380084846139aa565b60006020828403121561475257600080fd5b81516001600160401b0381111561476857600080fd5b61380084828501613a2c565b60006080828403121561478657600080fd5b60006138008484613d2a565b6000602082840312156147a457600080fd5b81356001600160401b038111156147ba57600080fd5b61380084828501613d99565b6000602082840312156147d857600080fd5b81356001600160401b038111156147ee57600080fd5b61380084828501613e71565b60006020828403121561480c57600080fd5b81516001600160401b0381111561482257600080fd5b61380084828501613f15565b60006020828403121561484057600080fd5b81356001600160401b0381111561485657600080fd5b613800848285016140ff565b60006020828403121561487457600080fd5b81356001600160401b0381111561488a57600080fd5b61380084828501614192565b6000806000606084860312156148ab57600080fd5b83356001600160401b038111156148c157600080fd5b6148cd86828701614192565b93505060208401356001600160401b038111156148e957600080fd5b6148f58682870161384a565b92505060408401356001600160401b0381111561491157600080fd5b6146e4868287016138ce565b60008060006060848603121561493257600080fd5b83356001600160401b0381111561494857600080fd5b6148cd86828701614271565b60006020828403121561496657600080fd5b81356001600160401b0381111561497c57600080fd5b613800848285016144f9565b600061499483836149ec565b505060200190565b60006125b98383614d02565b60006149b48383614d83565b505060400190565b60006125b98383614fe0565b60006125b9838361501f565b60006125b983836151b0565b60006125b983836151fc565b6149f58161598f565b82525050565b6149f5614a078261598f565b615adf565b6000614a16825190565b80845260209384019383018060005b83811015614a4a578151614a398882614988565b975060208301925050600101614a25565b509495945050505050565b6000614a5f825190565b80845260208401935083602082028501614a798560200190565b8060005b85811015614aae5784840389528151614a96858261499c565b94506020830160209a909a0199925050600101614a7d565b5091979650505050505050565b6000614ac5825190565b80845260209384019383018060005b83811015614a4a578151614ae888826149a8565b975060208301925050600101614ad4565b6000614b03825190565b80845260209384019383018060005b83811015614a4a578151614b2688826149a8565b975060208301925050600101614b12565b6000614b41825190565b80845260208401935083602082028501614b5b8560200190565b8060005b85811015614aae5784840389528151614b7885826149bc565b94506020830160209a909a0199925050600101614b5f565b6000614b9a825190565b80845260208401935083602082028501614bb48560200190565b8060005b85811015614aae5784840389528151614bd185826149c8565b94506020830160209a909a0199925050600101614bb8565b6000614bf3825190565b80845260208401935083602082028501614c0d8560200190565b8060005b85811015614aae5784840389528151614c2a85826149c8565b94506020830160209a909a0199925050600101614c11565b6000614c4c825190565b80845260208401935083602082028501614c668560200190565b8060005b85811015614aae5784840389528151614c8385826149d4565b94506020830160209a909a0199925050600101614c6a565b6000614ca5825190565b80845260208401935083602082028501614cbf8560200190565b8060005b85811015614aae5784840389528151614cdc85826149e0565b94506020830160209a909a0199925050600101614cc3565b8015156149f5565b806149f5565b6000614d0c825190565b808452602084019350614d23818560208601615a01565b601f01601f19169290920192915050565b6000614d3e825190565b614d4c818560208601615a01565b9290920192915050565b6149f5816159a0565b6149f5816159c9565b6149f5816159d4565b6149f5816159df565b6149f5816159ea565b80516040830190614d9484826153a9565b5060208201516119c260208501826153a9565b805161032080845260009190840190614dc08282614d02565b9150506020830151614dd560208601826149ec565b5060408301518482036040860152614ded8282614d02565b9150506060830151614e0260608601826153c4565b506080830151614e1560808601826153c4565b5060a0830151614e2860a08601826153c4565b5060c0830151614e3b60c08601826153c4565b5060e0830151614e4e60e08601826153c4565b50610100830151614e63610100860182614cfc565b50610120830151614e786101208601826153c4565b50610140830151614e8d6101408601826153c4565b50610160830151848203610160860152614ea78282614fa5565b915050610180830151614ebe6101808601826153c4565b506101a0830151614ed36101a0860182614cfc565b506101c0830151614ee86101c0860182614cf4565b506101e0830151614efd6101e0860182614d71565b50610200830151614f126102008601826153c4565b50610220830151848203610220860152614f2c8282614a0c565b915050610240830151848203610240860152614f488282614a0c565b915050610260830151848203610260860152614f648282614d02565b915050610280830151614f7b610280860182614d68565b506102a0830151614f906102a0860182614cf4565b506102c08301516125b76102c08601826150ff565b8051604080845260009190840190614fbd8282614d02565b91505060208301518482036020860152614fd78282614d02565b95945050505050565b80516000906060840190614ff485826153c4565b50602083015161500760208601826153c4565b5060408301518482036040860152614fd78282614d02565b8051600090604084019061503385826153c4565b5060208301518482036020860152614fd78282614b37565b805160e0808452600091908401906150638282614a55565b9150506020830151848203602086015261507d8282614a55565b915050604083015184820360408601526150978282614abb565b915050606083015184820360608601526150b18282614b90565b915050608083015184820360808601526150cb8282614a55565b91505060a083015184820360a08601526150e58282614da7565b91505060c083015184820360c0860152614fd78282614d02565b8051606083019061511084826153c4565b50602082015161512360208501826153c4565b5060408201516119c260408501826153c4565b805160009060a084019061514a85826153c4565b50602083015184820360208601526151628282614d02565b9150506040830151848203604086015261517c8282614a55565b915050606083015184820360608601526151968282614a55565b91505060808301518482036080860152614fd78282614a55565b80516060808452600091908401906151c88282615136565b915050602083015184820360208601526151e28282614abb565b91505060408301518482036040860152614fd78282614b90565b80516040808452600091908401906152148282614d02565b91505060208301516125b760208601826153c4565b805160009061018084019061523e85826149ec565b50602083015161525160208601826153c4565b50604083015161526460408601826153c4565b50606083015161527760608601826153c4565b50608083015161528a6080860182614d68565b5060a083015161529d60a0860182614cfc565b5060c08301516152b060c0860182614cfc565b5060e08301516152c360e08601826153c4565b506101008301516152d86101008601826153c4565b506101208301516152ed6101208601826153c4565b50610140830151615302610140860182614cf4565b50610160830151848203610160860152614fd78282614a55565b805160009060c084019061533085826153c4565b50602083015161534360208601826153c4565b506040830151848203604086015261535b8282614d02565b915050606083015184820360608601526153758282614a55565b9150506080830151848203608086015261538f8282614b90565b91505060a083015184820360a0860152614fd78282614d02565b63ffffffff81166149f5565b6149f563ffffffff8216615af1565b6001600160401b0381166149f5565b6149f56001600160401b038216615afd565b60006153f182856149fb565b6014820191506154018284614cfc565b5060200192915050565b600061541782876149fb565b6014820191506154278286614cfc565b60208201915061543782856153d3565b60088201915061544782846153d3565b50600801949350505050565b600061545f8284614cfc565b50602001919050565b60006125b98284614d34565b60006154808285614d34565b91506138008284614d34565b60006154988285614d34565b91506154a482846153d3565b5060080192915050565b60006154ba82856153b5565b6004820191506154ca82846153b5565b5060040192915050565b60006154e082886153d3565b6008820191506154f08287614d34565b91506154fc8286614d34565b91506155088285614d34565b91506155148284614d34565b979650505050505050565b602080825281016125b98184614af9565b602080825281016125b98184614be9565b602080825281016125b98184614c42565b6020810161282b8284614cf4565b602080825281016125b98184614d02565b60a0810161557f8288614d56565b61558c60208301876153c4565b818103604083015261559e8186614c9b565b905081810360608301526155b28185614af9565b90508181036080830152615514818461504b565b6020810161282b8284614d5f565b6020810161282b8284614d7a565b6020808252810161282b81602e81527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160208201527f647920696e697469616c697a6564000000000000000000000000000000000000604082015260600190565b602080825281016125b9818461504b565b604080825281016156658185615229565b90508181036020830152613800818461531c565b6020810161282b8284614cfc565b60a0810161569582886153c4565b81810360208301526156a78187614c9b565b905081810360408301526156bb8186614af9565b905081810360608301526155b2818561531c565b60006156da60405190565b90506156e68282615a54565b919050565b60006001600160401b0382111561570457615704615ba5565b5060209081020190565b60006001600160401b0382111561572757615727615ba5565b601f19601f83011660200192915050565b6000821982111561574b5761574b615b4d565b500190565b600063ffffffff8216915063ffffffff831692508263ffffffff0382111561574b5761574b615b4d565b60006001600160401b03821691506001600160401b0383169250826001600160401b030382111561574b5761574b615b4d565b600060ff8216915060ff831692508260ff0382111561574b5761574b615b4d565b6001600160401b0391821691166000826157ea576157ea615b63565b500490565b80825b600185111561582e5780860481111561580d5761580d615b4d565b600185161561581b57908102905b80026158278560011c90565b94506157f2565b94509492505050565b60006125b96000198484600082615850575060016125b9565b8161585d575060006125b9565b8160018114615873576002811461587d576158aa565b60019150506125b9565b60ff84111561588e5761588e615b4d565b8360020a9150848211156158a4576158a4615b4d565b506125b9565b5060208310610133831016604e8410600b84101617156158dd575081810a838111156158d8576158d8615b4d565b6125b9565b6158ea84848460016157ef565b9250905081840481111561590057615900615b4d565b0292915050565b600081600019048311821515161561592157615921615b4d565b500290565b60006001600160401b03821691506001600160401b0383169250816001600160401b03048311821515161561592157615921615b4d565b6000825b92508282101561597357615973615b4d565b500390565b600063ffffffff8216915063ffffffff8316615961565b60006001600160a01b03821661282b565b600061282b8261598f565b806156e681615bbb565b806156e681615bcb565b806156e681615bdb565b600061282b826159ab565b600061282b826159b5565b600061282b826159bf565b600060ff821661282b565b82818337506000910152565b60005b83811015615a1c578181015183820152602001615a04565b838111156119c25750506000910152565b600281046001821680615a4157607f821691505b602082108114156129ca576129ca615b8f565b601f19601f83011681018181106001600160401b0382111715615a7957615a79615ba5565b6040525050565b6000600019821415615a9457615a94615b4d565b5060010190565b600063ffffffff8216915063ffffffff821415615a9457615a94615b4d565b60006001600160401b03821691506001600160401b03821415615a9457615a94615b4d565b600061282b82600061282b8260601b90565b600061282b8260e01b90565b600061282b8260c01b90565b600063ffffffff8216915063ffffffff83165b925082615b2b57615b2b615b63565b500690565b60006001600160401b03821691506001600160401b038316615b1c565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600b811061247d5761247d615b79565b6003811061247d5761247d615b79565b6002811061247d5761247d615b79565b615bf48161598f565b811461247d57600080fd5b801515615bf4565b80615bf4565b615bf4816159a0565b6003811061247d57600080fd5b6002811061247d57600080fd5b63ffffffff8116615bf4565b6001600160401b038116615bf456fea2646970667358221220fed750d15b6ba52c3fe1292ebfa60f0bc0ec96822eb83a328457dc660bae09e764736f6c63430008040033",
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

// GenChallenge is a free data retrieval call binding the contract method 0xd48841c0.
//
// Solidity: function GenChallenge((address,bytes32,uint64,uint64) gParams) view returns((uint32,uint32)[])
func (_Store *StoreCaller) GenChallenge(opts *bind.CallOpts, gParams GenChallengeParams) ([]Challenge, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GenChallenge", gParams)

	if err != nil {
		return *new([]Challenge), err
	}

	out0 := *abi.ConvertType(out[0], new([]Challenge)).(*[]Challenge)

	return out0, err

}

// GenChallenge is a free data retrieval call binding the contract method 0xd48841c0.
//
// Solidity: function GenChallenge((address,bytes32,uint64,uint64) gParams) view returns((uint32,uint32)[])
func (_Store *StoreSession) GenChallenge(gParams GenChallengeParams) ([]Challenge, error) {
	return _Store.Contract.GenChallenge(&_Store.CallOpts, gParams)
}

// GenChallenge is a free data retrieval call binding the contract method 0xd48841c0.
//
// Solidity: function GenChallenge((address,bytes32,uint64,uint64) gParams) view returns((uint32,uint32)[])
func (_Store *StoreCallerSession) GenChallenge(gParams GenChallengeParams) ([]Challenge, error) {
	return _Store.Contract.GenChallenge(&_Store.CallOpts, gParams)
}

// GetChallengeKey is a free data retrieval call binding the contract method 0x83807a43.
//
// Solidity: function GetChallengeKey((uint32,uint32) chg) pure returns(bytes)
func (_Store *StoreCaller) GetChallengeKey(opts *bind.CallOpts, chg Challenge) ([]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetChallengeKey", chg)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetChallengeKey is a free data retrieval call binding the contract method 0x83807a43.
//
// Solidity: function GetChallengeKey((uint32,uint32) chg) pure returns(bytes)
func (_Store *StoreSession) GetChallengeKey(chg Challenge) ([]byte, error) {
	return _Store.Contract.GetChallengeKey(&_Store.CallOpts, chg)
}

// GetChallengeKey is a free data retrieval call binding the contract method 0x83807a43.
//
// Solidity: function GetChallengeKey((uint32,uint32) chg) pure returns(bytes)
func (_Store *StoreCallerSession) GetChallengeKey(chg Challenge) ([]byte, error) {
	return _Store.Contract.GetChallengeKey(&_Store.CallOpts, chg)
}

// GetChallengeList is a free data retrieval call binding the contract method 0x904a8d41.
//
// Solidity: function GetChallengeList(bytes key) view returns((uint32,uint32)[])
func (_Store *StoreCaller) GetChallengeList(opts *bind.CallOpts, key []byte) ([]Challenge, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetChallengeList", key)

	if err != nil {
		return *new([]Challenge), err
	}

	out0 := *abi.ConvertType(out[0], new([]Challenge)).(*[]Challenge)

	return out0, err

}

// GetChallengeList is a free data retrieval call binding the contract method 0x904a8d41.
//
// Solidity: function GetChallengeList(bytes key) view returns((uint32,uint32)[])
func (_Store *StoreSession) GetChallengeList(key []byte) ([]Challenge, error) {
	return _Store.Contract.GetChallengeList(&_Store.CallOpts, key)
}

// GetChallengeList is a free data retrieval call binding the contract method 0x904a8d41.
//
// Solidity: function GetChallengeList(bytes key) view returns((uint32,uint32)[])
func (_Store *StoreCallerSession) GetChallengeList(key []byte) ([]Challenge, error) {
	return _Store.Contract.GetChallengeList(&_Store.CallOpts, key)
}

// GetKeyByProofParams is a free data retrieval call binding the contract method 0xefea7783.
//
// Solidity: function GetKeyByProofParams((uint64,bytes,bytes[],bytes[],bytes[]) vParams) pure returns(bytes)
func (_Store *StoreCaller) GetKeyByProofParams(opts *bind.CallOpts, vParams ProofParams) ([]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetKeyByProofParams", vParams)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetKeyByProofParams is a free data retrieval call binding the contract method 0xefea7783.
//
// Solidity: function GetKeyByProofParams((uint64,bytes,bytes[],bytes[],bytes[]) vParams) pure returns(bytes)
func (_Store *StoreSession) GetKeyByProofParams(vParams ProofParams) ([]byte, error) {
	return _Store.Contract.GetKeyByProofParams(&_Store.CallOpts, vParams)
}

// GetKeyByProofParams is a free data retrieval call binding the contract method 0xefea7783.
//
// Solidity: function GetKeyByProofParams((uint64,bytes,bytes[],bytes[],bytes[]) vParams) pure returns(bytes)
func (_Store *StoreCallerSession) GetKeyByProofParams(vParams ProofParams) ([]byte, error) {
	return _Store.Contract.GetKeyByProofParams(&_Store.CallOpts, vParams)
}

// GetMerkleNodeKey is a free data retrieval call binding the contract method 0xb750ab8a.
//
// Solidity: function GetMerkleNodeKey((uint64,uint64,bytes) mn) pure returns(bytes)
func (_Store *StoreCaller) GetMerkleNodeKey(opts *bind.CallOpts, mn MerkleNode) ([]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetMerkleNodeKey", mn)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetMerkleNodeKey is a free data retrieval call binding the contract method 0xb750ab8a.
//
// Solidity: function GetMerkleNodeKey((uint64,uint64,bytes) mn) pure returns(bytes)
func (_Store *StoreSession) GetMerkleNodeKey(mn MerkleNode) ([]byte, error) {
	return _Store.Contract.GetMerkleNodeKey(&_Store.CallOpts, mn)
}

// GetMerkleNodeKey is a free data retrieval call binding the contract method 0xb750ab8a.
//
// Solidity: function GetMerkleNodeKey((uint64,uint64,bytes) mn) pure returns(bytes)
func (_Store *StoreCallerSession) GetMerkleNodeKey(mn MerkleNode) ([]byte, error) {
	return _Store.Contract.GetMerkleNodeKey(&_Store.CallOpts, mn)
}

// GetMerklePathKey is a free data retrieval call binding the contract method 0x3d1731b8.
//
// Solidity: function GetMerklePathKey((uint64,(uint64,uint64,bytes)[]) mp) pure returns(bytes)
func (_Store *StoreCaller) GetMerklePathKey(opts *bind.CallOpts, mp MerklePath) ([]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetMerklePathKey", mp)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetMerklePathKey is a free data retrieval call binding the contract method 0x3d1731b8.
//
// Solidity: function GetMerklePathKey((uint64,(uint64,uint64,bytes)[]) mp) pure returns(bytes)
func (_Store *StoreSession) GetMerklePathKey(mp MerklePath) ([]byte, error) {
	return _Store.Contract.GetMerklePathKey(&_Store.CallOpts, mp)
}

// GetMerklePathKey is a free data retrieval call binding the contract method 0x3d1731b8.
//
// Solidity: function GetMerklePathKey((uint64,(uint64,uint64,bytes)[]) mp) pure returns(bytes)
func (_Store *StoreCallerSession) GetMerklePathKey(mp MerklePath) ([]byte, error) {
	return _Store.Contract.GetMerklePathKey(&_Store.CallOpts, mp)
}

// GetMerklePathList is a free data retrieval call binding the contract method 0xb8527b31.
//
// Solidity: function GetMerklePathList(bytes key) view returns((uint64,(uint64,uint64,bytes)[])[])
func (_Store *StoreCaller) GetMerklePathList(opts *bind.CallOpts, key []byte) ([]MerklePath, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetMerklePathList", key)

	if err != nil {
		return *new([]MerklePath), err
	}

	out0 := *abi.ConvertType(out[0], new([]MerklePath)).(*[]MerklePath)

	return out0, err

}

// GetMerklePathList is a free data retrieval call binding the contract method 0xb8527b31.
//
// Solidity: function GetMerklePathList(bytes key) view returns((uint64,(uint64,uint64,bytes)[])[])
func (_Store *StoreSession) GetMerklePathList(key []byte) ([]MerklePath, error) {
	return _Store.Contract.GetMerklePathList(&_Store.CallOpts, key)
}

// GetMerklePathList is a free data retrieval call binding the contract method 0xb8527b31.
//
// Solidity: function GetMerklePathList(bytes key) view returns((uint64,(uint64,uint64,bytes)[])[])
func (_Store *StoreCallerSession) GetMerklePathList(key []byte) ([]MerklePath, error) {
	return _Store.Contract.GetMerklePathList(&_Store.CallOpts, key)
}

// GetUnVerifyProofList is a free data retrieval call binding the contract method 0xf5dbc78e.
//
// Solidity: function GetUnVerifyProofList() view returns(((uint64,bytes,bytes[],bytes[],bytes[]),(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[])[])
func (_Store *StoreCaller) GetUnVerifyProofList(opts *bind.CallOpts) ([]ProofRecordWithParams, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetUnVerifyProofList")

	if err != nil {
		return *new([]ProofRecordWithParams), err
	}

	out0 := *abi.ConvertType(out[0], new([]ProofRecordWithParams)).(*[]ProofRecordWithParams)

	return out0, err

}

// GetUnVerifyProofList is a free data retrieval call binding the contract method 0xf5dbc78e.
//
// Solidity: function GetUnVerifyProofList() view returns(((uint64,bytes,bytes[],bytes[],bytes[]),(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[])[])
func (_Store *StoreSession) GetUnVerifyProofList() ([]ProofRecordWithParams, error) {
	return _Store.Contract.GetUnVerifyProofList(&_Store.CallOpts)
}

// GetUnVerifyProofList is a free data retrieval call binding the contract method 0xf5dbc78e.
//
// Solidity: function GetUnVerifyProofList() view returns(((uint64,bytes,bytes[],bytes[],bytes[]),(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[])[])
func (_Store *StoreCallerSession) GetUnVerifyProofList() ([]ProofRecordWithParams, error) {
	return _Store.Contract.GetUnVerifyProofList(&_Store.CallOpts)
}

// PrepareForPdpVerification is a free data retrieval call binding the contract method 0x57d60d6e.
//
// Solidity: function PrepareForPdpVerification(((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]),(uint32,uint32)[],(uint64,uint64,bytes,bytes[],(uint64,(uint64,uint64,bytes)[])[],bytes)) pParams) view returns((bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes[],(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)),string))
func (_Store *StoreCaller) PrepareForPdpVerification(opts *bind.CallOpts, pParams PrepareForPdpVerificationParams) (PdpVerificationReturns, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "PrepareForPdpVerification", pParams)

	if err != nil {
		return *new(PdpVerificationReturns), err
	}

	out0 := *abi.ConvertType(out[0], new(PdpVerificationReturns)).(*PdpVerificationReturns)

	return out0, err

}

// PrepareForPdpVerification is a free data retrieval call binding the contract method 0x57d60d6e.
//
// Solidity: function PrepareForPdpVerification(((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]),(uint32,uint32)[],(uint64,uint64,bytes,bytes[],(uint64,(uint64,uint64,bytes)[])[],bytes)) pParams) view returns((bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes[],(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)),string))
func (_Store *StoreSession) PrepareForPdpVerification(pParams PrepareForPdpVerificationParams) (PdpVerificationReturns, error) {
	return _Store.Contract.PrepareForPdpVerification(&_Store.CallOpts, pParams)
}

// PrepareForPdpVerification is a free data retrieval call binding the contract method 0x57d60d6e.
//
// Solidity: function PrepareForPdpVerification(((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]),(uint32,uint32)[],(uint64,uint64,bytes,bytes[],(uint64,(uint64,uint64,bytes)[])[],bytes)) pParams) view returns((bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes[],(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,(bytes,bytes),uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)),string))
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

// VerifyProofWithMerklePathForFile is a free data retrieval call binding the contract method 0x5bf7f1d2.
//
// Solidity: function VerifyProofWithMerklePathForFile((uint64,bytes,bytes[],bytes[],bytes[]) vParams) view returns(bool)
func (_Store *StoreCaller) VerifyProofWithMerklePathForFile(opts *bind.CallOpts, vParams ProofParams) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "VerifyProofWithMerklePathForFile", vParams)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProofWithMerklePathForFile is a free data retrieval call binding the contract method 0x5bf7f1d2.
//
// Solidity: function VerifyProofWithMerklePathForFile((uint64,bytes,bytes[],bytes[],bytes[]) vParams) view returns(bool)
func (_Store *StoreSession) VerifyProofWithMerklePathForFile(vParams ProofParams) (bool, error) {
	return _Store.Contract.VerifyProofWithMerklePathForFile(&_Store.CallOpts, vParams)
}

// VerifyProofWithMerklePathForFile is a free data retrieval call binding the contract method 0x5bf7f1d2.
//
// Solidity: function VerifyProofWithMerklePathForFile((uint64,bytes,bytes[],bytes[],bytes[]) vParams) view returns(bool)
func (_Store *StoreCallerSession) VerifyProofWithMerklePathForFile(vParams ProofParams) (bool, error) {
	return _Store.Contract.VerifyProofWithMerklePathForFile(&_Store.CallOpts, vParams)
}

// BytesToUint is a free data retrieval call binding the contract method 0x02d06d05.
//
// Solidity: function bytesToUint(bytes b) pure returns(uint256)
func (_Store *StoreCaller) BytesToUint(opts *bind.CallOpts, b []byte) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "bytesToUint", b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BytesToUint is a free data retrieval call binding the contract method 0x02d06d05.
//
// Solidity: function bytesToUint(bytes b) pure returns(uint256)
func (_Store *StoreSession) BytesToUint(b []byte) (*big.Int, error) {
	return _Store.Contract.BytesToUint(&_Store.CallOpts, b)
}

// BytesToUint is a free data retrieval call binding the contract method 0x02d06d05.
//
// Solidity: function bytesToUint(bytes b) pure returns(uint256)
func (_Store *StoreCallerSession) BytesToUint(b []byte) (*big.Int, error) {
	return _Store.Contract.BytesToUint(&_Store.CallOpts, b)
}

// GenChallengeKey is a free data retrieval call binding the contract method 0x3c279f5d.
//
// Solidity: function genChallengeKey((address,bytes32,uint64,uint64) gParams) pure returns(string)
func (_Store *StoreCaller) GenChallengeKey(opts *bind.CallOpts, gParams GenChallengeParams) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "genChallengeKey", gParams)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GenChallengeKey is a free data retrieval call binding the contract method 0x3c279f5d.
//
// Solidity: function genChallengeKey((address,bytes32,uint64,uint64) gParams) pure returns(string)
func (_Store *StoreSession) GenChallengeKey(gParams GenChallengeParams) (string, error) {
	return _Store.Contract.GenChallengeKey(&_Store.CallOpts, gParams)
}

// GenChallengeKey is a free data retrieval call binding the contract method 0x3c279f5d.
//
// Solidity: function genChallengeKey((address,bytes32,uint64,uint64) gParams) pure returns(string)
func (_Store *StoreCallerSession) GenChallengeKey(gParams GenChallengeParams) (string, error) {
	return _Store.Contract.GenChallengeKey(&_Store.CallOpts, gParams)
}

// SaveChallenge is a paid mutator transaction binding the contract method 0x44c2b91b.
//
// Solidity: function SaveChallenge(bytes key, (uint32,uint32)[] chgs) payable returns()
func (_Store *StoreTransactor) SaveChallenge(opts *bind.TransactOpts, key []byte, chgs []Challenge) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SaveChallenge", key, chgs)
}

// SaveChallenge is a paid mutator transaction binding the contract method 0x44c2b91b.
//
// Solidity: function SaveChallenge(bytes key, (uint32,uint32)[] chgs) payable returns()
func (_Store *StoreSession) SaveChallenge(key []byte, chgs []Challenge) (*types.Transaction, error) {
	return _Store.Contract.SaveChallenge(&_Store.TransactOpts, key, chgs)
}

// SaveChallenge is a paid mutator transaction binding the contract method 0x44c2b91b.
//
// Solidity: function SaveChallenge(bytes key, (uint32,uint32)[] chgs) payable returns()
func (_Store *StoreTransactorSession) SaveChallenge(key []byte, chgs []Challenge) (*types.Transaction, error) {
	return _Store.Contract.SaveChallenge(&_Store.TransactOpts, key, chgs)
}

// SaveMerklePath is a paid mutator transaction binding the contract method 0x743b9eb6.
//
// Solidity: function SaveMerklePath(bytes key, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreTransactor) SaveMerklePath(opts *bind.TransactOpts, key []byte, mp []MerklePath) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SaveMerklePath", key, mp)
}

// SaveMerklePath is a paid mutator transaction binding the contract method 0x743b9eb6.
//
// Solidity: function SaveMerklePath(bytes key, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreSession) SaveMerklePath(key []byte, mp []MerklePath) (*types.Transaction, error) {
	return _Store.Contract.SaveMerklePath(&_Store.TransactOpts, key, mp)
}

// SaveMerklePath is a paid mutator transaction binding the contract method 0x743b9eb6.
//
// Solidity: function SaveMerklePath(bytes key, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreTransactorSession) SaveMerklePath(key []byte, mp []MerklePath) (*types.Transaction, error) {
	return _Store.Contract.SaveMerklePath(&_Store.TransactOpts, key, mp)
}

// SubmitVerifyProofRequest is a paid mutator transaction binding the contract method 0x8b90b4cd.
//
// Solidity: function SubmitVerifyProofRequest((uint64,bytes,bytes[],bytes[],bytes[]) vParams, (uint32,uint32)[] chgs, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreTransactor) SubmitVerifyProofRequest(opts *bind.TransactOpts, vParams ProofParams, chgs []Challenge, mp []MerklePath) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SubmitVerifyProofRequest", vParams, chgs, mp)
}

// SubmitVerifyProofRequest is a paid mutator transaction binding the contract method 0x8b90b4cd.
//
// Solidity: function SubmitVerifyProofRequest((uint64,bytes,bytes[],bytes[],bytes[]) vParams, (uint32,uint32)[] chgs, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreSession) SubmitVerifyProofRequest(vParams ProofParams, chgs []Challenge, mp []MerklePath) (*types.Transaction, error) {
	return _Store.Contract.SubmitVerifyProofRequest(&_Store.TransactOpts, vParams, chgs, mp)
}

// SubmitVerifyProofRequest is a paid mutator transaction binding the contract method 0x8b90b4cd.
//
// Solidity: function SubmitVerifyProofRequest((uint64,bytes,bytes[],bytes[],bytes[]) vParams, (uint32,uint32)[] chgs, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreTransactorSession) SubmitVerifyProofRequest(vParams ProofParams, chgs []Challenge, mp []MerklePath) (*types.Transaction, error) {
	return _Store.Contract.SubmitVerifyProofRequest(&_Store.TransactOpts, vParams, chgs, mp)
}

// VerifyProof is a paid mutator transaction binding the contract method 0xc70a45db.
//
// Solidity: function VerifyProof(((uint64,bytes,bytes[],bytes[],bytes[]),bool,uint256) vParams, (uint32,uint32)[] chgs, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreTransactor) VerifyProof(opts *bind.TransactOpts, vParams ProofRecord, chgs []Challenge, mp []MerklePath) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "VerifyProof", vParams, chgs, mp)
}

// VerifyProof is a paid mutator transaction binding the contract method 0xc70a45db.
//
// Solidity: function VerifyProof(((uint64,bytes,bytes[],bytes[],bytes[]),bool,uint256) vParams, (uint32,uint32)[] chgs, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreSession) VerifyProof(vParams ProofRecord, chgs []Challenge, mp []MerklePath) (*types.Transaction, error) {
	return _Store.Contract.VerifyProof(&_Store.TransactOpts, vParams, chgs, mp)
}

// VerifyProof is a paid mutator transaction binding the contract method 0xc70a45db.
//
// Solidity: function VerifyProof(((uint64,bytes,bytes[],bytes[],bytes[]),bool,uint256) vParams, (uint32,uint32)[] chgs, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreTransactorSession) VerifyProof(vParams ProofRecord, chgs []Challenge, mp []MerklePath) (*types.Transaction, error) {
	return _Store.Contract.VerifyProof(&_Store.TransactOpts, vParams, chgs, mp)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _file, address _sector, address _pdpExtra) returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts, _file common.Address, _sector common.Address, _pdpExtra common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize", _file, _sector, _pdpExtra)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _file, address _sector, address _pdpExtra) returns()
func (_Store *StoreSession) Initialize(_file common.Address, _sector common.Address, _pdpExtra common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _file, _sector, _pdpExtra)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _file, address _sector, address _pdpExtra) returns()
func (_Store *StoreTransactorSession) Initialize(_file common.Address, _sector common.Address, _pdpExtra common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _file, _sector, _pdpExtra)
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

// StoreInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Store contract.
type StoreInitializedIterator struct {
	Event *StoreInitialized // Event containing the contract specifics and raw log

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
func (it *StoreInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreInitialized)
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
		it.Event = new(StoreInitialized)
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
func (it *StoreInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreInitialized represents a Initialized event raised by the Store contract.
type StoreInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Store *StoreFilterer) FilterInitialized(opts *bind.FilterOpts) (*StoreInitializedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StoreInitializedIterator{contract: _Store.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Store *StoreFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StoreInitialized) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreInitialized)
				if err := _Store.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Store *StoreFilterer) ParseInitialized(log types.Log) (*StoreInitialized, error) {
	event := new(StoreInitialized)
	if err := _Store.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// StoreVerifyProofWithMerklePathForFileEventIterator is returned from FilterVerifyProofWithMerklePathForFileEvent and is used to iterate over the raw logs and unpacked data for VerifyProofWithMerklePathForFileEvent events raised by the Store contract.
type StoreVerifyProofWithMerklePathForFileEventIterator struct {
	Event *StoreVerifyProofWithMerklePathForFileEvent // Event containing the contract specifics and raw log

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
func (it *StoreVerifyProofWithMerklePathForFileEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreVerifyProofWithMerklePathForFileEvent)
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
		it.Event = new(StoreVerifyProofWithMerklePathForFileEvent)
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
func (it *StoreVerifyProofWithMerklePathForFileEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreVerifyProofWithMerklePathForFileEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreVerifyProofWithMerklePathForFileEvent represents a VerifyProofWithMerklePathForFileEvent event raised by the Store contract.
type StoreVerifyProofWithMerklePathForFileEvent struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVerifyProofWithMerklePathForFileEvent is a free log retrieval operation binding the contract event 0x8d5d73f2c0d3568f31aa45a7dc46558789a3defcd708f945f4cfe63776c314c2.
//
// Solidity: event VerifyProofWithMerklePathForFileEvent()
func (_Store *StoreFilterer) FilterVerifyProofWithMerklePathForFileEvent(opts *bind.FilterOpts) (*StoreVerifyProofWithMerklePathForFileEventIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "VerifyProofWithMerklePathForFileEvent")
	if err != nil {
		return nil, err
	}
	return &StoreVerifyProofWithMerklePathForFileEventIterator{contract: _Store.contract, event: "VerifyProofWithMerklePathForFileEvent", logs: logs, sub: sub}, nil
}

// WatchVerifyProofWithMerklePathForFileEvent is a free log subscription operation binding the contract event 0x8d5d73f2c0d3568f31aa45a7dc46558789a3defcd708f945f4cfe63776c314c2.
//
// Solidity: event VerifyProofWithMerklePathForFileEvent()
func (_Store *StoreFilterer) WatchVerifyProofWithMerklePathForFileEvent(opts *bind.WatchOpts, sink chan<- *StoreVerifyProofWithMerklePathForFileEvent) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "VerifyProofWithMerklePathForFileEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreVerifyProofWithMerklePathForFileEvent)
				if err := _Store.contract.UnpackLog(event, "VerifyProofWithMerklePathForFileEvent", log); err != nil {
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

// ParseVerifyProofWithMerklePathForFileEvent is a log parse operation binding the contract event 0x8d5d73f2c0d3568f31aa45a7dc46558789a3defcd708f945f4cfe63776c314c2.
//
// Solidity: event VerifyProofWithMerklePathForFileEvent()
func (_Store *StoreFilterer) ParseVerifyProofWithMerklePathForFileEvent(log types.Log) (*StoreVerifyProofWithMerklePathForFileEvent, error) {
	event := new(StoreVerifyProofWithMerklePathForFileEvent)
	if err := _Store.contract.UnpackLog(event, "VerifyProofWithMerklePathForFileEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
