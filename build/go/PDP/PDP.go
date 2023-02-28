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

// ProofParams is an auto generated low-level Go binding around an user-defined struct.
type ProofParams struct {
	Version    uint64
	Proofs     []byte
	FileIds    [][]byte
	Tags       [][]byte
	RootHashes []byte
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
	Tags         []byte
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"port\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"}],\"name\":\"DNSNodeRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DNSNodeUnReg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"DnsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"FsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structNameInfo\",\"name\":\"newer\",\"type\":\"tuple\"}],\"name\":\"NotifyNameInfoAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"PDPVerifyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nodeAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"VerifyProofWithMerklePathForFileEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"HashValue\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveNum\",\"type\":\"uint64\"}],\"internalType\":\"structGenChallengeParams\",\"name\":\"gParams\",\"type\":\"tuple\"}],\"name\":\"GenChallenge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge\",\"name\":\"chg\",\"type\":\"tuple\"}],\"name\":\"GetChallengeKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"GetChallengeList\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"FileIds\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"}],\"internalType\":\"structProofParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"GetKeyByProofParams\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode\",\"name\":\"mn\",\"type\":\"tuple\"}],\"name\":\"GetMerkleNodeKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath\",\"name\":\"mp\",\"type\":\"tuple\"}],\"name\":\"GetMerklePathKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"GetMerklePathList\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetUnVerifyProofList\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"FileIds\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"}],\"internalType\":\"structProofParams\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"challenge\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"merklePath\",\"type\":\"tuple[]\"}],\"internalType\":\"structProofRecordWithParams[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"SectorInfo_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"Challenges\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"ProveFileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"BlockNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Tags\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"MerklePath_\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"PlotData\",\"type\":\"bytes\"}],\"internalType\":\"structSectorProveData\",\"name\":\"ProveData\",\"type\":\"tuple\"}],\"internalType\":\"structPrepareForPdpVerificationParams\",\"name\":\"pParams\",\"type\":\"tuple\"}],\"name\":\"PrepareForPdpVerification\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"FileIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"UpdatedChal\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"Path\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"FileInfo_\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"Success\",\"type\":\"bool\"}],\"internalType\":\"structPdpVerificationReturns\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"chgs\",\"type\":\"tuple[]\"}],\"name\":\"SaveChallenge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"mp\",\"type\":\"tuple[]\"}],\"name\":\"SaveMerklePath\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"FileIds\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"}],\"internalType\":\"structProofParams\",\"name\":\"Proof\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"State\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"LastUpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structProofRecord\",\"name\":\"vParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"chgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"mp\",\"type\":\"tuple[]\"}],\"name\":\"SubmitVerifyProofRequest\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"PlotData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"}],\"internalType\":\"structVerifyPlotDataParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"VerifyPlotData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"FileIds\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"}],\"internalType\":\"structProofParams\",\"name\":\"Proof\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"State\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"LastUpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structProofRecord\",\"name\":\"vParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"chgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"mp\",\"type\":\"tuple[]\"}],\"name\":\"VerifyProof\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"FileIds\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"}],\"internalType\":\"structProofParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"VerifyProofWithMerklePathForFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"bytesToUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"HashValue\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveNum\",\"type\":\"uint64\"}],\"internalType\":\"structGenChallengeParams\",\"name\":\"gParams\",\"type\":\"tuple\"}],\"name\":\"genChallengeKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50614731806100206000396000f3fe6080604052600436106101295760003560e01c8063904a8d41116100a5578063b750ab8a11610074578063ba9da17611610059578063ba9da17614610328578063f5dbc78e14610348578063fd7c98081461036a57600080fd5b8063b750ab8a146102db578063b8527b31146102fb57600080fd5b8063904a8d411461024e5780639f9ca6441461027b578063ab33ced4146102a8578063ae78fb6f146102bb57600080fd5b80635498d974116100fc5780637cdb2459116100e15780637cdb2459146101f95780638129fc1c1461021957806383807a431461022e57600080fd5b80635498d974146101d3578063743b9eb6146101e657600080fd5b806302d06d051461012e5780632e19aeff146101645780633d1731b81461019157806344c2b91b146101be575b600080fd5b34801561013a57600080fd5b5061014e6101493660046134c2565b61038a565b60405161015b91906141b5565b60405180910390f35b34801561017057600080fd5b5061018461017f36600461376d565b61041d565b60405161015b9190614116565b34801561019d57600080fd5b506101b16101ac366004613640565b61043d565b60405161015b9190614124565b6101d16101cc3660046134f6565b6104f7565b005b6101d16101e13660046136dc565b6105b5565b6101d16101f436600461355d565b610690565b34801561020557600080fd5b506101b16102143660046135d8565b610917565b34801561022557600080fd5b506101d1610959565b34801561023a57600080fd5b506101b16102493660046135ba565b6109ee565b34801561025a57600080fd5b5061026e6102693660046134c2565b610a0d565b60405161015b91906140e3565b34801561028757600080fd5b5061029b610296366004613674565b610b85565b60405161015b91906141a4565b6101d16102b63660046136dc565b610bae565b3480156102c757600080fd5b506101846102d63660046136a8565b610c25565b3480156102e757600080fd5b506101b16102f636600461360c565b610f86565b34801561030757600080fd5b5061031b6103163660046134c2565b610fa5565b60405161015b91906140f4565b34801561033457600080fd5b506101b16103433660046136a8565b611306565b34801561035457600080fd5b5061035d611479565b60405161015b9190614105565b34801561037657600080fd5b5061026e6103853660046135d8565b611685565b60008060005b8351811015610416576103a481600161422c565b84516103b09190614451565b6103bb9060086143fb565b6103c690600261432b565b8482815181106103e657634e487b7160e01b600052603260045260246000fd5b01602001516103f8919060f81c6143fb565b610402908361422c565b91508061040e8161456b565b915050610390565b5092915050565b600061043582604001516001600160401b0316611bcc565b506001919050565b60608060005b8360200151518163ffffffff1610156104165783602001518163ffffffff168151811061048057634e487b7160e01b600052603260045260246000fd5b60200260200101516040015184602001518263ffffffff16815181106104b657634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516040516020016104d3929190614042565b604051602081830303815290604052915080806104ef90614586565b915050610443565b60005b81518163ffffffff1610156105b0576000610541838363ffffffff168151811061053457634e487b7160e01b600052603260045260246000fd5b60200260200101516109ee565b905061059b81848463ffffffff168151811061056d57634e487b7160e01b600052603260045260246000fd5b6020026020010151600487604051610585919061401e565b9081526040519081900360200190209190611c3d565b505080806105a890614586565b9150506104fa565b505050565b6105fe604080516101008101909152600060608083019182526080830181905260a0830181905260c0830181905260e08301528190815260006020820181905260409091015290565b83515181516001600160401b039091169052835160209081015182518201528451604090810151835182015285516060908101518451909101528551608090810151845190910152818601511515918301919091524390820152835160009061066690611306565b905061067281856104f7565b61067c8184610690565b61068860018284611d7d565b505050505050565b60005b81518163ffffffff1610156105b05760006106da838363ffffffff16815181106106cd57634e487b7160e01b600052603260045260246000fd5b602002602001015161043d565b90506000838363ffffffff168151811061070457634e487b7160e01b600052603260045260246000fd5b602002602001015160200151516001600160401b0381111561073657634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561076957816020015b60608152602001906001900390816107545790505b50905060005b848463ffffffff168151811061079557634e487b7160e01b600052603260045260246000fd5b602002602001015160200151518163ffffffff1610156108d5576000610818868663ffffffff16815181106107da57634e487b7160e01b600052603260045260246000fd5b6020026020010151602001518363ffffffff168151811061080b57634e487b7160e01b600052603260045260246000fd5b6020026020010151610f86565b905061088e81878763ffffffff168151811061084457634e487b7160e01b600052603260045260246000fd5b6020026020010151602001518463ffffffff168151811061087557634e487b7160e01b600052603260045260246000fd5b60200260200101516006611e919092919063ffffffff16565b5080838363ffffffff16815181106108b657634e487b7160e01b600052603260045260246000fd5b60200260200101819052505080806108cd90614586565b91505061076f565b5061090182826005886040516108eb919061401e565b9081526040519081900360200190209190611f5a565b505050808061090f90614586565b915050610693565b6060600082600001518360200151846040015185606001516040516020016109429493929190613fc5565b60408051601f198184030181529190529392505050565b600054610100900460ff166109745760005460ff1615610978565b303b155b6109b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ae90614143565b60405180910390fd5b600054610100900460ff161580156109d9576000805461ffff19166101011790555b80156109eb576000805461ff00191690555b50565b6060600082600001518360200151604051602001610942929190614064565b60606000600483604051610a21919061401e565b908152604051908190036020019020600101546001600160401b03811115610a5957634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610a9e57816020015b6040805180820190915260008082526020820152815260200190600190039081610a775790505b5090506000610ac9600485604051610ab6919061401e565b9081526020016040518091039020611fc5565b90505b610af581600486604051610ae0919061401e565b90815260405190819003602001902090611fe0565b15610416576000610b2582600487604051610b10919061401e565b90815260405190819003602001902090611fee565b91505080838381518110610b4957634e487b7160e01b600052603260045260246000fd5b602002602001018190525050610b7e81600486604051610b69919061401e565b90815260405190819003602001902090612117565b9050610acc565b610b8d612923565b815151610b999061218d565b610ba1612923565b600160c082015292915050565b6000610bbd8460000151611306565b9050610bc981846104f7565b610bd38183610690565b60006020850152610be660018286611d7d565b507f7f6b6a2262c2fc2cc1b1785be448eae36656117d8f558facfe3c199e26256ea1600a604051610c179190614135565b60405180910390a150505050565b600080610c3183611306565b90506000600160000182604051610c48919061401e565b9081526040805191829003602001822061010083019091526001810180546001600160401b031660608401908152600290920180549192849290918491608085019190610c9490614518565b80601f0160208091040260200160405190810160405280929190818152602001828054610cc090614518565b8015610d0d5780601f10610ce257610100808354040283529160200191610d0d565b820191906000526020600020905b815481529060010190602001808311610cf057829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b82821015610de7578382906000526020600020018054610d5a90614518565b80601f0160208091040260200160405190810160405280929190818152602001828054610d8690614518565b8015610dd35780601f10610da857610100808354040283529160200191610dd3565b820191906000526020600020905b815481529060010190602001808311610db657829003601f168201915b505050505081526020019060010190610d3b565b50505050815260200160038201805480602002602001604051908101604052809291908181526020016000905b82821015610ec0578382906000526020600020018054610e3390614518565b80601f0160208091040260200160405190810160405280929190818152602001828054610e5f90614518565b8015610eac5780601f10610e8157610100808354040283529160200191610eac565b820191906000526020600020905b815481529060010190602001808311610e8f57829003601f168201915b505050505081526020019060010190610e14565b505050508152602001600482018054610ed890614518565b80601f0160208091040260200160405190810160405280929190818152602001828054610f0490614518565b8015610f515780601f10610f2657610100808354040283529160200191610f51565b820191906000526020600020905b815481529060010190602001808311610f3457829003601f168201915b505050919092525050508152600582015460ff1615156020808301919091526006909201546040909101520151949350505050565b6060600082604001518360200151604051602001610942929190614042565b60606000600583604051610fb9919061401e565b908152604051908190036020019020600101546001600160401b03811115610ff157634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561103757816020015b60408051808201909152600081526060602082015281526020019060019003908161100f5790505b509050600061106260058560405161104f919061401e565b90815260200160405180910390206121fe565b90505b61107981600586604051610ae0919061401e565b156104165760006110a982600587604051611094919061401e565b9081526040519081900360200190209061220c565b91505080516001600160401b038111156110d357634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561112057816020015b604080516060808201835260008083526020830152918101919091528152602001906001900390816110f15790505b5083838151811061114157634e487b7160e01b600052603260045260246000fd5b60200260200101516020018190525060005b81518110156112d457600660000182828151811061118157634e487b7160e01b600052603260045260246000fd5b6020026020010151604051611196919061401e565b9081526040805191829003602090810183206060840183526001810180546001600160401b038082168752680100000000000000009091041692850192909252600201805491928401916111e990614518565b80601f016020809104026020016040519081016040528092919081815260200182805461121590614518565b80156112625780601f1061123757610100808354040283529160200191611262565b820191906000526020600020905b81548152906001019060200180831161124557829003601f168201915b50505050508152505084848151811061128b57634e487b7160e01b600052603260045260246000fd5b60200260200101516020015182815181106112b657634e487b7160e01b600052603260045260246000fd5b602002602001018190525080806112cc9061456b565b915050611153565b50506112ff816005866040516112ea919061401e565b908152604051908190036020019020906123ca565b9050611065565b60608060005b8360400151518163ffffffff161015611387578184604001518263ffffffff168151811061134a57634e487b7160e01b600052603260045260246000fd5b602002602001015160405160200161136392919061402a565b6040516020818303038152906040529150808061137f90614586565b91505061130c565b50606060005b8460600151518163ffffffff161015611408578185606001518263ffffffff16815181106113cb57634e487b7160e01b600052603260045260246000fd5b60200260200101516040516020016113e492919061402a565b6040516020818303038152906040529150808061140090614586565b91505061138d565b508351602080860151608087015160405160009461142d94909392889288920161408a565b60408051601f19818403018152908290528051602080830191909120919350839260009161145d91849101614009565b60408051601f1981840301815291905298975050505050505050565b6060600080611488600161243a565b90505b611496600182611fe0565b156114d55760006114a8600183612448565b91505080602001516114c257826114be8161456b565b9350505b506114ce600182612892565b905061148b565b506000816001600160401b038111156114fe57634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561157857816020015b611565604080516101008101909152600060608083019182526080830181905260a0830181905260c0830181905260e08301528190815260200160608152602001606081525090565b81526020019060019003908161151c5790505b5090506115b66040518060a0016040528060006001600160401b03168152602001606081526020016060815260200160608152602001606081525090565b60608060006115c5600161243a565b90505b6115d3600182611fe0565b1561167a5760006115e5600183612448565b9150508060200151156115f85750611668565b80519450600061160786611306565b905061161281610a0d565b945061161d81610fa5565b935060405180606001604052808781526020018681526020018581525087848151811061165a57634e487b7160e01b600052603260045260246000fd5b602002602001018190525050505b611673600182612892565b90506115c8565b509295945050505050565b60208082015182516040516060936000926116a4929091859101613fa9565b604051602081830303815290604052905060006002826040516116c7919061401e565b602060405180830381855afa1580156116e4573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061170791906134a4565b604080516024808252606082019092529192506000919060208201818036833701905050905060005b60208163ffffffff1610156117b857828163ffffffff166020811061176557634e487b7160e01b600052603260045260246000fd5b1a60f81b828263ffffffff168151811061178f57634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a905350806117b081614586565b915050611730565b5060005b60048163ffffffff16101561184e57828163ffffffff16602081106117f157634e487b7160e01b600052603260045260246000fd5b1a60f81b82611801836020614244565b63ffffffff168151811061182557634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a9053508061184681614586565b9150506117bc565b506000806000600389604001516001600160401b0316116118885750505060408601516001600160401b0316606087015260018080611905565b600389604001516001600160401b03161180156118be575088606001516001600160401b031689604001516001600160401b0316105b156118cb57600360608a01525b886060015189604001516118df91906142c2565b9250886060015189604001516118f591906145f6565b6118ff908461426e565b91508290505b600089606001516001600160401b03166001600160401b0381111561193a57634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561197f57816020015b60408051808201909152600080825260208201528152602001906001900390816119585790505b50905085600060015b8c606001516001600160401b03168163ffffffff1611611bbb578c606001516001600160401b03168163ffffffff1614156119c1578594505b60408051600480825281830190925260009160208201818036833701905050905060005b60048163ffffffff161015611a7d57896119ff8286614244565b63ffffffff1681518110611a2357634e487b7160e01b600052603260045260246000fd5b602001015160f81c60f81b828263ffffffff1681518110611a5457634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a90535080611a7581614586565b9150506119e5565b506000611a898261038a565b905088611a9760018561446c565b63ffffffff16611aa7919061441a565b87611ab3836001614244565b63ffffffff16611ac391906145f6565b611acd919061426e565b86611ad960018661446c565b63ffffffff1681518110611afd57634e487b7160e01b600052603260045260246000fd5b60209081029190910181015163ffffffff928316905286918616908110611b3457634e487b7160e01b600052603260045260246000fd5b611b4191901a60016142a1565b60ff1686611b5060018661446c565b63ffffffff1681518110611b7457634e487b7160e01b600052603260045260246000fd5b60209081029190910181015163ffffffff90921691015283611b9581614586565b9450611ba490506020856145cf565b935050508080611bb390614586565b915050611988565b50919b9a5050505050505050505050565b6109eb81604051602401611be091906141b5565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff5b1bba900000000000000000000000000000000000000000000000000000000179052612902565b6000808460000184604051611c52919061401e565b90815260405190819003602001812054915083908690611c7390879061401e565b908152604051908190036020908101909120825160019091018054939092015163ffffffff9081166401000000000267ffffffffffffffff199094169116179190911790558015611cc8576001915050611d76565b5060018085018054808301825560009190915290611ce790829061422c565b6040518690611cf790879061401e565b9081526040519081900360200190205560018501805485919083908110611d2e57634e487b7160e01b600052603260045260246000fd5b90600052602060002090600202016000019080519060200190611d52929190612a33565b50600285018054906000611d658361456b565b91905055506000915050611d76565b505b9392505050565b6000808460000184604051611d92919061401e565b90815260405190819003602001812054915083908690611db390879061401e565b908152604051602091819003820190208251805160018301805467ffffffffffffffff19166001600160401b03909216919091178155818401518051919492938593611e06936002909201920190612a33565b5060408201518051611e22916002840191602090910190612ab7565b5060608201518051611e3e916003840191602090910190612ab7565b5060808201518051611e5a916004840191602090910190612a33565b505050602082015160058201805460ff19169115159190911790556040909101516006909101558015611cc8576001915050611d76565b6000808460000184604051611ea6919061401e565b90815260405190819003602001812054915083908690611ec790879061401e565b908152604080519182900360209081019092208351600182018054868601516001600160401b0390811668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009092169316929092179190911781559184015180519293611f4693600290930192910190612a33565b505081159050611cc8576001915050611d76565b6000808460000184604051611f6f919061401e565b90815260405190819003602001812054915083908690611f9090879061401e565b90815260200160405180910390206001019080519060200190611fb4929190612ab7565b508015611cc8576001915050611d76565b600080611fd3836000612117565b9050611d76600182614451565b600182015481105b92915050565b604080518082019091526000808252602082015260609083600101838154811061202857634e487b7160e01b600052603260045260246000fd5b9060005260206000209060020201600001805461204490614518565b80601f016020809104026020016040519081016040528092919081815260200182805461207090614518565b80156120bd5780601f10612092576101008083540402835291602001916120bd565b820191906000526020600020905b8154815290600101906020018083116120a057829003601f168201915b5050505050915083600001826040516120d6919061401e565b90815260408051918290036020908101832083830190925260019091015463ffffffff80821684526401000000009091041690820152919491935090915050565b6000816121238161456b565b9250505b600183015482108015612170575082600101828154811061215857634e487b7160e01b600052603260045260246000fd5b600091825260209091206001600290920201015460ff165b15612187578161217f8161456b565b925050612127565b50919050565b6109eb816040516024016121a191906140d5565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f2c2ecbc200000000000000000000000000000000000000000000000000000000179052612902565b600080611fd38360006123ca565b60608083600101838154811061223257634e487b7160e01b600052603260045260246000fd5b9060005260206000209060020201600001805461224e90614518565b80601f016020809104026020016040519081016040528092919081815260200182805461227a90614518565b80156122c75780601f1061229c576101008083540402835291602001916122c7565b820191906000526020600020905b8154815290600101906020018083116122aa57829003601f168201915b5050505050915083600001826040516122e0919061401e565b9081526020016040518091039020600101805480602002602001604051908101604052809291908181526020016000905b828210156123bd57838290600052602060002001805461233090614518565b80601f016020809104026020016040519081016040528092919081815260200182805461235c90614518565b80156123a95780601f1061237e576101008083540402835291602001916123a9565b820191906000526020600020905b81548152906001019060200180831161238c57829003601f168201915b505050505081526020019060010190612311565b5050505090509250929050565b6000816123d68161456b565b9250505b600183015482108015612423575082600101828154811061240b57634e487b7160e01b600052603260045260246000fd5b600091825260209091206001600290920201015460ff165b1561218757816124328161456b565b9250506123da565b600080611fd3836000612892565b6040805161010081018252600060608281018281526080840182905260a0840182905260c0840182905260e08401829052835260208301829052928201528360010183815481106124a957634e487b7160e01b600052603260045260246000fd5b906000526020600020906002020160000180546124c590614518565b80601f01602080910402602001604051908101604052809291908181526020018280546124f190614518565b801561253e5780601f106125135761010080835404028352916020019161253e565b820191906000526020600020905b81548152906001019060200180831161252157829003601f168201915b505050505091508360000182604051612557919061401e565b9081526040805191829003602001822061010083019091526001810180546001600160401b0316606084019081526002909201805491928492909184916080850191906125a390614518565b80601f01602080910402602001604051908101604052809291908181526020018280546125cf90614518565b801561261c5780601f106125f15761010080835404028352916020019161261c565b820191906000526020600020905b8154815290600101906020018083116125ff57829003601f168201915b5050505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b828210156126f657838290600052602060002001805461266990614518565b80601f016020809104026020016040519081016040528092919081815260200182805461269590614518565b80156126e25780601f106126b7576101008083540402835291602001916126e2565b820191906000526020600020905b8154815290600101906020018083116126c557829003601f168201915b50505050508152602001906001019061264a565b50505050815260200160038201805480602002602001604051908101604052809291908181526020016000905b828210156127cf57838290600052602060002001805461274290614518565b80601f016020809104026020016040519081016040528092919081815260200182805461276e90614518565b80156127bb5780601f10612790576101008083540402835291602001916127bb565b820191906000526020600020905b81548152906001019060200180831161279e57829003601f168201915b505050505081526020019060010190612723565b5050505081526020016004820180546127e790614518565b80601f016020809104026020016040519081016040528092919081815260200182805461281390614518565b80156128605780601f1061283557610100808354040283529160200191612860565b820191906000526020600020905b81548152906001019060200180831161284357829003601f168201915b505050919092525050508152600582015460ff1615156020820152600690910154604090910152919491935090915050565b60008161289e8161456b565b9250505b6001830154821080156128eb57508260010182815481106128d357634e487b7160e01b600052603260045260246000fd5b600091825260209091206001600290920201015460ff165b1561218757816128fa8161456b565b9250506128a2565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b6040518060e001604052806060815260200160608152602001606081526020016060815260200160608152602001612a26604080516102e08101825260608082526000602083018190529282018190528082018390526080820183905260a0820183905260c0820183905260e08201839052610100820183905261012082018390526101408201839052610160820181905261018082018390526101a082018390526101c082018390526101e0820183905261020082018390526102208201819052610240820181905261026082015290610280820190815260006020808301829052604080516060810182528381529182018390528181019290925291015290565b8152600060209091015290565b828054612a3f90614518565b90600052602060002090601f016020900481019282612a615760008555612aa7565b82601f10612a7a57805160ff1916838001178555612aa7565b82800160010185558215612aa7579182015b82811115612aa7578251825591602001919060010190612a8c565b50612ab3929150612b10565b5090565b828054828255906000526020600020908101928215612b04579160200282015b82811115612b045782518051612af4918491602090910190612a33565b5091602001919060010190612ad7565b50612ab3929150612b25565b5b80821115612ab35760008155600101612b11565b80821115612ab3576000612b398282612b42565b50600101612b25565b508054612b4e90614518565b6000825580601f10612b5e575050565b601f0160209004906000526020600020908101906109eb9190612b10565b6000612b8f612b8a846141df565b6141c3565b90508083825260208201905082856020860282011115612bae57600080fd5b60005b85811015612bf35781356001600160401b03811115612bcf57600080fd5b808601612bdc8982612e17565b855250506020928301929190910190600101612bb1565b5050509392505050565b6000612c0b612b8a846141df565b83815290506020810182604085028101861015612c2757600080fd5b60005b85811015612bf35781612c3d8882612e43565b84525060209092019160409190910190600101612c2a565b6000612c63612b8a846141df565b90508083825260208201905082856020860282011115612c8257600080fd5b60005b85811015612bf35781356001600160401b03811115612ca357600080fd5b808601612cb08982612f10565b855250506020928301929190910190600101612c85565b6000612cd5612b8a846141df565b90508083825260208201905082856020860282011115612cf457600080fd5b60005b85811015612bf35781356001600160401b03811115612d1557600080fd5b808601612d228982612f82565b855250506020928301929190910190600101612cf7565b6000612d47612b8a84614202565b905082815260208101848484011115612d5f57600080fd5b611d748482856144e0565b8035611fe8816146b1565b600082601f830112612d8657600080fd5b8135612d96848260208601612b7c565b949350505050565b600082601f830112612daf57600080fd5b8135612d96848260208601612bfd565b600082601f830112612dd057600080fd5b8135612d96848260208601612c55565b600082601f830112612df157600080fd5b8135612d96848260208601612cc7565b8035611fe8816146c5565b8051611fe8816146cd565b600082601f830112612e2857600080fd5b8135612d96848260208601612d39565b8035611fe8816146d3565b600060408284031215612e5557600080fd5b612e5f60406141c3565b90506000612e6d848461348e565b8252506020612e7e8484830161348e565b60208301525092915050565b600060808284031215612e9c57600080fd5b612ea660806141c3565b90506000612eb48484612d6a565b82525060208201356001600160401b03811115612ed057600080fd5b612edc84828501612e17565b6020830152506040612ef084828501613499565b6040830152506060612f0484828501613499565b60608301525092915050565b600060608284031215612f2257600080fd5b612f2c60606141c3565b90506000612f3a8484613499565b8252506020612f4b84848301613499565b60208301525060408201356001600160401b03811115612f6a57600080fd5b612f7684828501612e17565b60408301525092915050565b600060408284031215612f9457600080fd5b612f9e60406141c3565b90506000612fac8484613499565b82525060208201356001600160401b03811115612fc857600080fd5b612e7e84828501612dbf565b600060608284031215612fe657600080fd5b612ff060606141c3565b90506000612ffe8484613499565b825250602061300f84848301613499565b6020830152506040612f7684828501613499565b60006060828403121561303557600080fd5b61303f60606141c3565b905081356001600160401b0381111561305757600080fd5b613063848285016131fa565b82525060208201356001600160401b0381111561307f57600080fd5b61308b84828501612d9e565b60208301525060408201356001600160401b038111156130aa57600080fd5b612f768482850161332a565b600060a082840312156130c857600080fd5b6130d260a06141c3565b905060006130e08484613499565b82525060208201356001600160401b038111156130fc57600080fd5b61310884828501612e17565b60208301525060408201356001600160401b0381111561312757600080fd5b61313384828501612d75565b60408301525060608201356001600160401b0381111561315257600080fd5b61315e84828501612d75565b60608301525060808201356001600160401b0381111561317d57600080fd5b61318984828501612e17565b60808301525092915050565b6000606082840312156131a757600080fd5b6131b160606141c3565b905081356001600160401b038111156131c957600080fd5b6131d5848285016130b6565b82525060206131e684848301612e01565b6020830152506040612f7684828501613483565b6000610180828403121561320d57600080fd5b6132186101806141c3565b905060006132268484612d6a565b825250602061323784848301613499565b602083015250604061324b84828501613499565b604083015250606061325f84828501613499565b606083015250608061327384828501612e38565b60808301525060a061328784828501613483565b60a08301525060c061329b84828501613483565b60c08301525060e06132af84828501613499565b60e0830152506101006132c484828501613499565b610100830152506101206132da84828501613499565b610120830152506101406132f084828501612e01565b610140830152506101608201356001600160401b0381111561331157600080fd5b61331d84828501612d75565b6101608301525092915050565b600060c0828403121561333c57600080fd5b61334660c06141c3565b905060006133548484613499565b825250602061336584848301613499565b60208301525060408201356001600160401b0381111561338457600080fd5b61339084828501612e17565b60408301525060608201356001600160401b038111156133af57600080fd5b6133bb84828501612e17565b60608301525060808201356001600160401b038111156133da57600080fd5b6133e684828501612de0565b60808301525060a08201356001600160401b0381111561340557600080fd5b61341184828501612e17565b60a08301525092915050565b600060a0828403121561342f57600080fd5b61343960606141c3565b905060006134478484612fd4565b82525060608201356001600160401b0381111561346357600080fd5b61346f84828501612e17565b6020830152506080612f7684828501613499565b8035611fe8816146cd565b8035611fe8816146e0565b8035611fe8816146ec565b6000602082840312156134b657600080fd5b6000612d968484612e0c565b6000602082840312156134d457600080fd5b81356001600160401b038111156134ea57600080fd5b612d9684828501612e17565b6000806040838503121561350957600080fd5b82356001600160401b0381111561351f57600080fd5b61352b85828601612e17565b92505060208301356001600160401b0381111561354757600080fd5b61355385828601612d9e565b9150509250929050565b6000806040838503121561357057600080fd5b82356001600160401b0381111561358657600080fd5b61359285828601612e17565b92505060208301356001600160401b038111156135ae57600080fd5b61355385828601612de0565b6000604082840312156135cc57600080fd5b6000612d968484612e43565b6000602082840312156135ea57600080fd5b81356001600160401b0381111561360057600080fd5b612d9684828501612e8a565b60006020828403121561361e57600080fd5b81356001600160401b0381111561363457600080fd5b612d9684828501612f10565b60006020828403121561365257600080fd5b81356001600160401b0381111561366857600080fd5b612d9684828501612f82565b60006020828403121561368657600080fd5b81356001600160401b0381111561369c57600080fd5b612d9684828501613023565b6000602082840312156136ba57600080fd5b81356001600160401b038111156136d057600080fd5b612d96848285016130b6565b6000806000606084860312156136f157600080fd5b83356001600160401b0381111561370757600080fd5b61371386828701613195565b93505060208401356001600160401b0381111561372f57600080fd5b61373b86828701612d9e565b92505060408401356001600160401b0381111561375757600080fd5b61376386828701612de0565b9150509250925092565b60006020828403121561377f57600080fd5b81356001600160401b0381111561379557600080fd5b612d968482850161341d565b60006137ad83836137f9565b505060200190565b6000611d768383613ab6565b60006137cd8383613b25565b505060400190565b6000611d768383613d4d565b6000611d768383613d95565b6000611d768383613f21565b61380281614483565b82525050565b61380261381482614483565b6145a5565b6000613823825190565b80845260209384019383018060005b8381101561385757815161384688826137a1565b975060208301925050600101613832565b509495945050505050565b600061386c825190565b808452602084019350836020820285016138868560200190565b8060005b858110156138bb57848403895281516138a385826137b5565b94506020830160209a909a019992505060010161388a565b5091979650505050505050565b60006138d2825190565b80845260209384019383018060005b838110156138575781516138f588826137c1565b9750602083019250506001016138e1565b6000613910825190565b80845260209384019383018060005b8381101561385757815161393388826137c1565b97506020830192505060010161391f565b600061394e825190565b808452602084019350836020820285016139688560200190565b8060005b858110156138bb578484038952815161398585826137d5565b94506020830160209a909a019992505060010161396c565b60006139a7825190565b808452602084019350836020820285016139c18560200190565b8060005b858110156138bb57848403895281516139de85826137e1565b94506020830160209a909a01999250506001016139c5565b6000613a00825190565b80845260208401935083602082028501613a1a8560200190565b8060005b858110156138bb5784840389528151613a3785826137e1565b94506020830160209a909a0199925050600101613a1e565b6000613a59825190565b80845260208401935083602082028501613a738560200190565b8060005b858110156138bb5784840389528151613a9085826137ed565b94506020830160209a909a0199925050600101613a77565b801515613802565b80613802565b6000613ac0825190565b808452602084019350613ad78185602086016144ec565b601f01601f19169290920192915050565b6000613af2825190565b613b008185602086016144ec565b9290920192915050565b613802816144bf565b613802816144ca565b613802816144d5565b80516040830190613b368482613f6d565b506020820151613b496020850182613f6d565b50505050565b805161032080845260009190840190613b688282613ab6565b9150506020830151613b7d60208601826137f9565b5060408301518482036040860152613b958282613ab6565b9150506060830151613baa6060860182613f88565b506080830151613bbd6080860182613f88565b5060a0830151613bd060a0860182613f88565b5060c0830151613be360c0860182613f88565b5060e0830151613bf660e0860182613f88565b50610100830151613c0b610100860182613ab0565b50610120830151613c20610120860182613f88565b50610140830151613c35610140860182613f88565b50610160830151848203610160860152613c4f8282613ab6565b915050610180830151613c66610180860182613f88565b506101a0830151613c7b6101a0860182613ab0565b506101c0830151613c906101c0860182613aa8565b506101e0830151613ca56101e0860182613b1c565b50610200830151613cba610200860182613f88565b50610220830151848203610220860152613cd48282613819565b915050610240830151848203610240860152613cf08282613819565b915050610260830151848203610260860152613d0c8282613ab6565b915050610280830151613d23610280860182613b13565b506102a0830151613d386102a0860182613aa8565b506102c0830151611d746102c0860182613e70565b80516000906060840190613d618582613f88565b506020830151613d746020860182613f88565b5060408301518482036040860152613d8c8282613ab6565b95945050505050565b80516000906040840190613da98582613f88565b5060208301518482036020860152613d8c8282613944565b805160e080845260009190840190613dd98282613862565b91505060208301518482036020860152613df38282613862565b91505060408301518482036040860152613e0d82826138c8565b91505060608301518482036060860152613e27828261399d565b91505060808301518482036080860152613e418282613ab6565b91505060a083015184820360a0860152613e5b8282613b4f565b91505060c0830151611d7460c0860182613aa8565b80516060830190613e818482613f88565b506020820151613e946020850182613f88565b506040820151613b496040850182613f88565b805160009060a0840190613ebb8582613f88565b5060208301518482036020860152613ed38282613ab6565b91505060408301518482036040860152613eed8282613862565b91505060608301518482036060860152613f078282613862565b91505060808301518482036080860152613d8c8282613ab6565b8051606080845260009190840190613f398282613ea7565b91505060208301518482036020860152613f5382826138c8565b91505060408301518482036040860152613d8c828261399d565b63ffffffff8116613802565b61380263ffffffff82166145b7565b6001600160401b038116613802565b6138026001600160401b0382166145c3565b6000613fb58285613808565b601482019150612d968284613ae8565b6000613fd18287613808565b601482019150613fe18286613ae8565b9150613fed8285613f97565b600882019150613ffd8284613f97565b50600801949350505050565b60006140158284613ab0565b50602001919050565b6000611d768284613ae8565b60006140368285613ae8565b9150612d968284613ae8565b600061404e8285613ae8565b915061405a8284613f97565b5060080192915050565b60006140708285613f79565b6004820191506140808284613f79565b5060040192915050565b60006140968288613f97565b6008820191506140a68287613ae8565b91506140b28286613ae8565b91506140be8285613ae8565b91506140ca8284613ae8565b979650505050505050565b60208101611fe882846137f9565b60208082528101611d768184613906565b60208082528101611d7681846139f6565b60208082528101611d768184613a4f565b60208101611fe88284613aa8565b60208082528101611d768184613ab6565b60208101611fe88284613b0a565b60208082528101611fe881602e81527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160208201527f647920696e697469616c697a6564000000000000000000000000000000000000604082015260600190565b60208082528101611d768184613dc1565b60208101611fe88284613ab0565b60006141ce60405190565b90506141da828261453f565b919050565b60006001600160401b038211156141f8576141f861466b565b5060209081020190565b60006001600160401b0382111561421b5761421b61466b565b601f19601f83011660200192915050565b6000821982111561423f5761423f614613565b500190565b600063ffffffff8216915063ffffffff831692508263ffffffff0382111561423f5761423f614613565b60006001600160401b03821691506001600160401b0383169250826001600160401b030382111561423f5761423f614613565b600060ff8216915060ff831692508260ff0382111561423f5761423f614613565b6001600160401b0391821691166000826142de576142de614629565b500490565b80825b60018511156143225780860481111561430157614301614613565b600185161561430f57908102905b800261431b8560011c90565b94506142e6565b94509492505050565b6000611d76600019848460008261434457506001611d76565b8161435157506000611d76565b816001811461436757600281146143715761439e565b6001915050611d76565b60ff84111561438257614382614613565b8360020a91508482111561439857614398614613565b50611d76565b5060208310610133831016604e8410600b84101617156143d1575081810a838111156143cc576143cc614613565b611d76565b6143de84848460016142e3565b925090508184048111156143f4576143f4614613565b0292915050565b600081600019048311821515161561441557614415614613565b500290565b60006001600160401b03821691506001600160401b0383169250816001600160401b03048311821515161561441557614415614613565b6000825b92508282101561446757614467614613565b500390565b600063ffffffff8216915063ffffffff8316614455565b600073ffffffffffffffffffffffffffffffffffffffff8216611fe8565b806141da81614681565b806141da81614691565b806141da816146a1565b6000611fe8826144a1565b6000611fe8826144ab565b6000611fe8826144b5565b82818337506000910152565b60005b838110156145075781810151838201526020016144ef565b83811115613b495750506000910152565b60028104600182168061452c57607f821691505b6020821081141561218757612187614655565b601f19601f83011681018181106001600160401b03821117156145645761456461466b565b6040525050565b600060001982141561457f5761457f614613565b5060010190565b600063ffffffff8216915063ffffffff82141561457f5761457f614613565b6000611fe8826000611fe88260601b90565b6000611fe88260e01b90565b6000611fe88260c01b90565b600063ffffffff8216915063ffffffff83165b9250826145f1576145f1614629565b500690565b60006001600160401b03821691506001600160401b0383166145e2565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600b81106109eb576109eb61463f565b600381106109eb576109eb61463f565b600281106109eb576109eb61463f565b6146ba81614483565b81146109eb57600080fd5b8015156146ba565b806146ba565b600381106109eb57600080fd5b63ffffffff81166146ba565b6001600160401b0381166146ba56fea2646970667358221220620ee6866483e849e2c58e8a46b24691e359f46b91f3424f81a36d78daa1dc5a64736f6c63430008040033",
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

// GetKeyByProofParams is a free data retrieval call binding the contract method 0xba9da176.
//
// Solidity: function GetKeyByProofParams((uint64,bytes,bytes[],bytes[],bytes) vParams) pure returns(bytes)
func (_Store *StoreCaller) GetKeyByProofParams(opts *bind.CallOpts, vParams ProofParams) ([]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetKeyByProofParams", vParams)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetKeyByProofParams is a free data retrieval call binding the contract method 0xba9da176.
//
// Solidity: function GetKeyByProofParams((uint64,bytes,bytes[],bytes[],bytes) vParams) pure returns(bytes)
func (_Store *StoreSession) GetKeyByProofParams(vParams ProofParams) ([]byte, error) {
	return _Store.Contract.GetKeyByProofParams(&_Store.CallOpts, vParams)
}

// GetKeyByProofParams is a free data retrieval call binding the contract method 0xba9da176.
//
// Solidity: function GetKeyByProofParams((uint64,bytes,bytes[],bytes[],bytes) vParams) pure returns(bytes)
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
// Solidity: function GetUnVerifyProofList() view returns(((uint64,bytes,bytes[],bytes[],bytes),(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[])[])
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
// Solidity: function GetUnVerifyProofList() view returns(((uint64,bytes,bytes[],bytes[],bytes),(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[])[])
func (_Store *StoreSession) GetUnVerifyProofList() ([]ProofRecordWithParams, error) {
	return _Store.Contract.GetUnVerifyProofList(&_Store.CallOpts)
}

// GetUnVerifyProofList is a free data retrieval call binding the contract method 0xf5dbc78e.
//
// Solidity: function GetUnVerifyProofList() view returns(((uint64,bytes,bytes[],bytes[],bytes),(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[])[])
func (_Store *StoreCallerSession) GetUnVerifyProofList() ([]ProofRecordWithParams, error) {
	return _Store.Contract.GetUnVerifyProofList(&_Store.CallOpts)
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

// VerifyProofWithMerklePathForFile is a free data retrieval call binding the contract method 0xae78fb6f.
//
// Solidity: function VerifyProofWithMerklePathForFile((uint64,bytes,bytes[],bytes[],bytes) vParams) view returns(bool)
func (_Store *StoreCaller) VerifyProofWithMerklePathForFile(opts *bind.CallOpts, vParams ProofParams) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "VerifyProofWithMerklePathForFile", vParams)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProofWithMerklePathForFile is a free data retrieval call binding the contract method 0xae78fb6f.
//
// Solidity: function VerifyProofWithMerklePathForFile((uint64,bytes,bytes[],bytes[],bytes) vParams) view returns(bool)
func (_Store *StoreSession) VerifyProofWithMerklePathForFile(vParams ProofParams) (bool, error) {
	return _Store.Contract.VerifyProofWithMerklePathForFile(&_Store.CallOpts, vParams)
}

// VerifyProofWithMerklePathForFile is a free data retrieval call binding the contract method 0xae78fb6f.
//
// Solidity: function VerifyProofWithMerklePathForFile((uint64,bytes,bytes[],bytes[],bytes) vParams) view returns(bool)
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

// GenChallengeKey is a free data retrieval call binding the contract method 0x7cdb2459.
//
// Solidity: function genChallengeKey((address,bytes,uint64,uint64) gParams) pure returns(string)
func (_Store *StoreCaller) GenChallengeKey(opts *bind.CallOpts, gParams GenChallengeParams) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "genChallengeKey", gParams)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GenChallengeKey is a free data retrieval call binding the contract method 0x7cdb2459.
//
// Solidity: function genChallengeKey((address,bytes,uint64,uint64) gParams) pure returns(string)
func (_Store *StoreSession) GenChallengeKey(gParams GenChallengeParams) (string, error) {
	return _Store.Contract.GenChallengeKey(&_Store.CallOpts, gParams)
}

// GenChallengeKey is a free data retrieval call binding the contract method 0x7cdb2459.
//
// Solidity: function genChallengeKey((address,bytes,uint64,uint64) gParams) pure returns(string)
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

// SubmitVerifyProofRequest is a paid mutator transaction binding the contract method 0xab33ced4.
//
// Solidity: function SubmitVerifyProofRequest(((uint64,bytes,bytes[],bytes[],bytes),bool,uint256) vParams, (uint32,uint32)[] chgs, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreTransactor) SubmitVerifyProofRequest(opts *bind.TransactOpts, vParams ProofRecord, chgs []Challenge, mp []MerklePath) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SubmitVerifyProofRequest", vParams, chgs, mp)
}

// SubmitVerifyProofRequest is a paid mutator transaction binding the contract method 0xab33ced4.
//
// Solidity: function SubmitVerifyProofRequest(((uint64,bytes,bytes[],bytes[],bytes),bool,uint256) vParams, (uint32,uint32)[] chgs, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreSession) SubmitVerifyProofRequest(vParams ProofRecord, chgs []Challenge, mp []MerklePath) (*types.Transaction, error) {
	return _Store.Contract.SubmitVerifyProofRequest(&_Store.TransactOpts, vParams, chgs, mp)
}

// SubmitVerifyProofRequest is a paid mutator transaction binding the contract method 0xab33ced4.
//
// Solidity: function SubmitVerifyProofRequest(((uint64,bytes,bytes[],bytes[],bytes),bool,uint256) vParams, (uint32,uint32)[] chgs, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreTransactorSession) SubmitVerifyProofRequest(vParams ProofRecord, chgs []Challenge, mp []MerklePath) (*types.Transaction, error) {
	return _Store.Contract.SubmitVerifyProofRequest(&_Store.TransactOpts, vParams, chgs, mp)
}

// VerifyProof is a paid mutator transaction binding the contract method 0x5498d974.
//
// Solidity: function VerifyProof(((uint64,bytes,bytes[],bytes[],bytes),bool,uint256) vParams, (uint32,uint32)[] chgs, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreTransactor) VerifyProof(opts *bind.TransactOpts, vParams ProofRecord, chgs []Challenge, mp []MerklePath) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "VerifyProof", vParams, chgs, mp)
}

// VerifyProof is a paid mutator transaction binding the contract method 0x5498d974.
//
// Solidity: function VerifyProof(((uint64,bytes,bytes[],bytes[],bytes),bool,uint256) vParams, (uint32,uint32)[] chgs, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreSession) VerifyProof(vParams ProofRecord, chgs []Challenge, mp []MerklePath) (*types.Transaction, error) {
	return _Store.Contract.VerifyProof(&_Store.TransactOpts, vParams, chgs, mp)
}

// VerifyProof is a paid mutator transaction binding the contract method 0x5498d974.
//
// Solidity: function VerifyProof(((uint64,bytes,bytes[],bytes[],bytes),bool,uint256) vParams, (uint32,uint32)[] chgs, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreTransactorSession) VerifyProof(vParams ProofRecord, chgs []Challenge, mp []MerklePath) (*types.Transaction, error) {
	return _Store.Contract.VerifyProof(&_Store.TransactOpts, vParams, chgs, mp)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Store *StoreSession) Initialize() (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Store *StoreTransactorSession) Initialize() (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts)
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
