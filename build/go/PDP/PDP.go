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
	FileIDs     []byte
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
	FileIds    []byte
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"port\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"}],\"name\":\"DNSNodeRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DNSNodeUnReg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"DnsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"FsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structNameInfo\",\"name\":\"newer\",\"type\":\"tuple\"}],\"name\":\"NotifyNameInfoAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"PDPVerifyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nodeAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"VerifyProofWithMerklePathForFileEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"HashValue\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveNum\",\"type\":\"uint64\"}],\"internalType\":\"structGenChallengeParams\",\"name\":\"gParams\",\"type\":\"tuple\"}],\"name\":\"GenChallenge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge\",\"name\":\"chg\",\"type\":\"tuple\"}],\"name\":\"GetChallengeKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"GetChallengeList\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"FileIds\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"}],\"internalType\":\"structProofParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"GetKeyByProofParams\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode\",\"name\":\"mn\",\"type\":\"tuple\"}],\"name\":\"GetMerkleNodeKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath\",\"name\":\"mp\",\"type\":\"tuple\"}],\"name\":\"GetMerklePathKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"GetMerklePathList\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetUnVerifyProofList\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"FileIds\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"}],\"internalType\":\"structProofParams\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"challenge\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"merklePath\",\"type\":\"tuple[]\"}],\"internalType\":\"structProofRecordWithParams[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"SectorInfo_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"Challenges\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"ProveFileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"BlockNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"MerklePath_\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"PlotData\",\"type\":\"bytes\"}],\"internalType\":\"structSectorProveData\",\"name\":\"ProveData\",\"type\":\"tuple\"}],\"internalType\":\"structPrepareForPdpVerificationParams\",\"name\":\"pParams\",\"type\":\"tuple\"}],\"name\":\"PrepareForPdpVerification\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileIDs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"UpdatedChal\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"Path\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"FileInfo_\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"Success\",\"type\":\"bool\"}],\"internalType\":\"structPdpVerificationReturns\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"chgs\",\"type\":\"tuple[]\"}],\"name\":\"SaveChallenge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"mp\",\"type\":\"tuple[]\"}],\"name\":\"SaveMerklePath\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"FileIds\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"}],\"internalType\":\"structProofParams\",\"name\":\"vParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"chgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"mp\",\"type\":\"tuple[]\"}],\"name\":\"SubmitVerifyProofRequest\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"PlotData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"}],\"internalType\":\"structVerifyPlotDataParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"VerifyPlotData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"FileIds\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"}],\"internalType\":\"structProofParams\",\"name\":\"Proof\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"State\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"LastUpdateHeight\",\"type\":\"uint256\"}],\"internalType\":\"structProofRecord\",\"name\":\"vParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"chgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"mp\",\"type\":\"tuple[]\"}],\"name\":\"VerifyProof\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"FileIds\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"}],\"internalType\":\"structProofParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"VerifyProofWithMerklePathForFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"bytesToUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0x7b63f42e\",\"type\":\"bytes32\"}],\"name\":\"c_0x7b63f42e\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"HashValue\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveNum\",\"type\":\"uint64\"}],\"internalType\":\"structGenChallengeParams\",\"name\":\"gParams\",\"type\":\"tuple\"}],\"name\":\"genChallengeKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5061ac8e80620000226000396000f3fe6080604052600436106101445760003560e01c80638129fc1c116100c0578063b8527b3111610074578063e0bdfb4011610059578063e0bdfb401461047c578063f5dbc78e14610498578063fd7c9808146104c357610144565b8063b8527b3114610416578063cca9fb6e1461045357610144565b8063904a8d41116100a5578063904a8d411461035f578063958e549c1461039c578063b750ab8a146103d957610144565b80638129fc1c1461030b57806383807a431461032257610144565b806344c2b91b1161011757806357d60d6e116100fc57806357d60d6e14610275578063743b9eb6146102b25780637cdb2459146102ce57610144565b806344c2b91b1461021c578063507af8051461023857610144565b806302d06d0514610149578063173dee68146101865780632e19aeff146101a25780633d1731b8146101df575b600080fd5b34801561015557600080fd5b50610170600480360381019061016b9190618f00565b610500565b60405161017d919061a074565b60405180910390f35b6101a0600480360381019061019b919061921e565b61074b565b005b3480156101ae57600080fd5b506101c960048036038101906101c491906192b5565b610c80565b6040516101d69190619fb8565b60405180910390f35b3480156101eb57600080fd5b50610206600480360381019061020191906190c4565b610d7e565b6040516102139190619fd3565b60405180910390f35b61023660048036038101906102319190618f41565b611001565b005b34801561024457600080fd5b5061025f600480360381019061025a9190619146565b61122b565b60405161026c9190619fb8565b60405180910390f35b34801561028157600080fd5b5061029c60048036038101906102979190619105565b6116a9565b6040516102a9919061a052565b60405180910390f35b6102cc60048036038101906102c79190618fad565b611870565b005b3480156102da57600080fd5b506102f560048036038101906102f09190619042565b611ed0565b604051610302919061a010565b60405180910390f35b34801561031757600080fd5b50610320611ff1565b005b34801561032e57600080fd5b5061034960048036038101906103449190619019565b612101565b6040516103569190619fd3565b60405180910390f35b34801561036b57600080fd5b5061038660048036038101906103819190618f00565b612216565b6040516103939190619f52565b60405180910390f35b3480156103a857600080fd5b506103c360048036038101906103be9190619146565b6125ba565b6040516103d09190619fd3565b60405180910390f35b3480156103e557600080fd5b5061040060048036038101906103fb9190619083565b612b62565b60405161040d9190619fd3565b60405180910390f35b34801561042257600080fd5b5061043d60048036038101906104389190618f00565b612c77565b60405161044a9190619f74565b60405180910390f35b34801561045f57600080fd5b5061047a60048036038101906104759190618eae565b613356565b005b61049660048036038101906104919190619187565b613359565b005b3480156104a457600080fd5b506104ad61367d565b6040516104ba9190619f96565b60405180910390f35b3480156104cf57600080fd5b506104ea60048036038101906104e59190619042565b613f96565b6040516104f79190619f52565b60405180910390f35b600061052e7f4798e23059a1935d93b02013f9b70d288619e89b8053545c12d30b98792e76b860001b613356565b61055a7f561ce1be6363840fc86d541dafb0b126a62dfaf632805675c424889a2bab6c5b60001b613356565b6105867f3580762c7cb245c677dc60faf9169fe9d75f41e4d2534af6b7ca4869ab9f5c4760001b613356565b60006105b47fd2bddbdd57f436c79427a2877438aa55c132a60777e411becf6b7984a7a9af3060001b613356565b6105e07fbe2ddbb3a8fafc3e1d52a75d9503601e64c9395ea5720530def6dfc5deaf4dd860001b613356565b60005b83518110156106e9576106187fbe50ccaf2547bc03ad1d7a02a05551cb908505903173ba791c8bc190fc00bab660001b613356565b6106447fc3de97f24f6d455a1e24872d57c8ce0e4b76b5f44175875a62b1e2661a93bed760001b613356565b600181610651919061a361565b845161065d919061a6a4565b6008610669919061a608565b6002610675919061a4ea565b8482815181106106ae577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b60f81c60ff166106c9919061a608565b826106d4919061a361565b915080806106e19061a8cf565b9150506105e3565b506107167f0246e7e1ddc834d4b85a0d24893b7c38226d33e3a24d2ac00a085397212e004460001b613356565b6107427f542ffdda3ca9b8a322515fb605caa0ce36f51827d6b175046230e0cdc24c2bf060001b613356565b80915050919050565b6107777fb47cfe15a4818e985c22ff97816cbcd517a0808eaaa0c73805be229b7e9ebd9a60001b613356565b6107a37f5d1461836ef28ddc91b89a398fef01d80f82207612805eb054884f554871cdbd60001b613356565b6107cf7fe1bb36cfdb944ec1a9bbb2735fab8b0f053da9914cb5cfb9f1e7e26446c4117560001b613356565b6107d7617fa7565b6108037f27a28d25a92d022cbbfcf430872055aac86143e6e2c2c769b782f314083d7c3260001b613356565b61082f7f9372b0bfda95dfbdecaafa5a9f1b6fa6565135990586caaffea62ec5a0727b4a60001b613356565b83600001516000015181600001516000019067ffffffffffffffff16908167ffffffffffffffff16815250506108877f53f44add46d313c43f4c7e483a639af69c9c0cf91528ac4f38d3abcae2a4348a60001b613356565b6108b37f3b0b9609d5b6602933412c347eb4a1fbba52b4d5a2e627c29417ddb16955591160001b613356565b8360000151602001518160000151602001819052506108f47fc1ac6fd6175f5b21ad78b9678c835536df82aa1f3ca5bff6ea0ac5a837cf90d460001b613356565b6109207f6875747c92d6f6bb89fe5191541b808e413f1b00f8209a94b0c3ab812772997260001b613356565b8360000151604001518160000151604001819052506109617f8a839cbc3e42263d68d0912409eae969e1f5316b481505d7cbecd153202fbc5c60001b613356565b61098d7fefe4b787044ab85df6f5ca579d2f2755a5d7a00b2f96ad92dc78368d6f1ef2d760001b613356565b8360000151606001518160000151606001819052506109ce7f17c04bdec2ed321a52db872406b709e047883de6aecfc88df58db6c7fd5f239860001b613356565b6109fa7ffcb3cbb9fcaa3980ca52956208ebf74e42ef93ea0de876bd1eb7fb90fb178c1860001b613356565b836000015160800151816000015160800181905250610a3b7fd97cee35214a85cf70eebaae73ac9fa907e9bf68540574e14e433ba88149942660001b613356565b610a677f9ecbc4049f504b851f717657570f4aea45a89e74693851997559fffb01716d7c60001b613356565b8360200151816020019015159081151581525050610aa77f81c42f9c5703d19757b6c9f7aef94620d755d1d1ef7f3a42c822f6d098620dbb60001b613356565b610ad37f7051d7a75558a841ee95462632197c32247b751c5be8ad820ce4b5d09490948160001b613356565b43816040018181525050610b097f85763f3642e4f8b8e8faa540b116671efec731acfe2ccfc3aa711fb654c78d0160001b613356565b610b357f170b6b90716d2b740d160904cd9719de0363ee7a7d9105ba7126b4041a925b7f60001b613356565b6000610b4485600001516125ba565b9050610b727f1736aa89ed002c01e5c6cdcd9ceaca9565c1be0f5772972c5ca468a3c70e276a60001b613356565b610b9e7f462ffd82ece63e025ddc5c6c19ebfd70952218c569713ffc42db14c8fab1001f60001b613356565b610ba88185611001565b610bd47f5e1030bd5c1eca173b804df2312b4ac70c117ea1d90c9b564601f91f5c339c2760001b613356565b610c007f4a4a7e989ce58282e5eccb80f62dcad3bbd4383b575719bfb2a472c9e5d5095360001b613356565b610c0a8184611870565b610c367f239d82dee4d1c1e06df20090da6d462fdbe225fb29a74ef1e008af302f988c1760001b613356565b610c627f152cc3322c6d0e42020843ae4b9ff9b729b1ebb1c516708a1787df8f7f01b31160001b613356565b610c78818360016154a49092919063ffffffff16565b505050505050565b6000610cae7f3975b31b83cb9aca0d9401d83ab84f65ae89845e838e85e0945961162f320ce260001b613356565b610cda7f48fd146058f3905ddc2ac44f22b99098adbef9a73f4193cc13c7d635dfdb2a6960001b613356565b610d067fc738e0ac2ada58d796c785bbd9472cfdfa52204fee6e2581bd1cb7372f7c84cd60001b613356565b610d1d826040015167ffffffffffffffff16615a63565b610d497f503b3270df9fe77f025e81d8b5ec2b6782d79b43ec5d8cf692c9e542f9bc1e5560001b613356565b610d757f5c6a018027cfa8e523934cb05cfa863bc4ed091426b7b0a491af6f9ebfd74eae60001b613356565b60019050919050565b6060610dac7f2ef3261d304397fade30a2863822cc38bb7edd7938db97bde8092c434bfef58160001b613356565b610dd87f3fad293f3937a5d556f11a1b4ecfd19d648462630eb4c49028ac1bf6dbe4184e60001b613356565b610e047fb15d31d2ac8e8c7b84c8948fc032bceb8356e99b9024809f31e4a107296d726960001b613356565b6060610e327fa9b72a6421f40200a2507f822753305a856bf5ecc348072acb09cb196a6990bf60001b613356565b610e5e7f3d4c3a00217e6b1a14b6c047400f71cf5976fa737fefa9d96ff948fb517e7e0660001b613356565b60005b8360200151518163ffffffff161015610f9f57610ea07f8647a079b8c15d1a97accf2bea0aa6d34322630afceff937cdb8c6c7967aa25b60001b613356565b610ecc7f6b025daff2e3afc18ae935d7c5fdacba325e54ec7f0c8f1183e17e69d5f84f8460001b613356565b83602001518163ffffffff1681518110610f0f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516040015184602001518263ffffffff1681518110610f5e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015160200151604051602001610f7b929190619e94565b60405160208183030381529060405291508080610f979061a918565b915050610e61565b50610fcc7fd677c24f38973f65f6a51a5db54e9a4d168d501e0c272b30959e0d670bf1126560001b613356565b610ff87f33f866deeb76bf492b4e8af29fdbec987349e61f284286bd2a33b9f1d0da590060001b613356565b80915050919050565b61102d7f0cf8882a6e670597ce701065042731166c76229cc2fb8b86b10076c9f7f0fb4d60001b613356565b6110597fafe5420563008f886007373c918df602196117023971b8109c031ad47fee263860001b613356565b6110857f07dac24474ed4982792b371fac314a132f467f5090fa2680ae9e8edb6741b50a60001b613356565b60005b81518163ffffffff161015611226576110c27ecf1ecd0f2e96fd578c09e06523afe69be8188560438a255338ccc173cf9b2e60001b613356565b6110ee7fd8b8adaff04f0f6b2c8a707aa0f015bc023fa019b3d87b98f8404d8bd5c9213460001b613356565b600061113f838363ffffffff1681518110611132577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151612101565b905061116d7f78c7d8365d3296bb9d8bb2d69b6e8162d14c14acb9bc29b8420ee8b849a759ed60001b613356565b6111997ff51399000028c4b80657a85421820ec97019b898acf9385d6dcbd8eb9de65d1f60001b613356565b61121181848463ffffffff16815181106111dc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516004876040516111f49190619e31565b9081526020016040518091039020615afc9092919063ffffffff16565b5050808061121e9061a918565b915050611088565b505050565b60006112597f35a7235e6422dc3e6faf0e37972937210da34c482a99f7c35a1c5412ded436fd60001b613356565b6112857fe5d0c138fc7e7729510df8c7ed6bfb09463d818bb0beffeca0d8b43d662024aa60001b613356565b6112b17ff7e39c8aa1eed7e8bdae88e4f8687683f85d757861d5b990ea56f436e122750460001b613356565b60006112bc836125ba565b90506112ea7f5666818069b6f45b0ede916cb6ae8f51402fdd4d18409fb4de3710a645198d1460001b613356565b6113167f4ff6ed2560a76cd3afcb61add3818e6cf9c640e5c2ada0d47a15a0fa9333e0b060001b613356565b600060016000018260405161132b9190619e31565b9081526020016040518091039020600101604051806060016040529081600082016040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820180546113999061a86c565b80601f01602080910402602001604051908101604052809291908181526020018280546113c59061a86c565b80156114125780601f106113e757610100808354040283529160200191611412565b820191906000526020600020905b8154815290600101906020018083116113f557829003601f168201915b5050505050815260200160028201805461142b9061a86c565b80601f01602080910402602001604051908101604052809291908181526020018280546114579061a86c565b80156114a45780601f10611479576101008083540402835291602001916114a4565b820191906000526020600020905b81548152906001019060200180831161148757829003601f168201915b5050505050815260200160038201805480602002602001604051908101604052809291908181526020016000905b8282101561157e5783829060005260206000200180546114f19061a86c565b80601f016020809104026020016040519081016040528092919081815260200182805461151d9061a86c565b801561156a5780601f1061153f5761010080835404028352916020019161156a565b820191906000526020600020905b81548152906001019060200180831161154d57829003601f168201915b5050505050815260200190600101906114d2565b5050505081526020016004820180546115969061a86c565b80601f01602080910402602001604051908101604052809291908181526020018280546115c29061a86c565b801561160f5780601f106115e45761010080835404028352916020019161160f565b820191906000526020600020905b8154815290600101906020018083116115f257829003601f168201915b50505050508152505081526020016005820160009054906101000a900460ff16151515158152602001600682015481525050905061166f7fbd4eba4b78d4f840697851f642707af0a53e7b6e252cee5e9bb55dfb26ea404360001b613356565b61169b7fa7f9c6f448f2561908f17b3426923c34ac720e483b314995d3baa1cb58fe954960001b613356565b806020015192505050919050565b6116b1617fd0565b6116dd7fc25f1242e1b2fcba055bb68ac76038a6b43663806ea61c7c4d01b85cece835b060001b613356565b6117097fde085b8c7bb830898531e8e4e7921997a812c030ab8dc92b503aeddb0e13d04260001b613356565b6117357fb93c6c65a12e78d0c8f5c945473c7f4faec4d642f5d8fc3d350a2de46b39e8bb60001b613356565b611746826000015160000151616031565b6117727fae5b48e5e72f0a100bb5acd1292478d0426f3010c793d1079ae1db0b5a5443e060001b613356565b61179e7f1e2fa8fd7b33af198c4d901c9bb3b28f037f6ef447cf100026aba95487ad669260001b613356565b6117a6617fd0565b6117d27fc64cbb3dc85a51c682ad83d8c7654e17062559398667eeb3b5987fbe97285ea060001b613356565b6117fe7fd3a9fb7fa1371ef00b85ab5c3ebeeea637b83eeed11b04ea323250459382a78160001b613356565b60018160c001901515908115158152505061183b7fd7276d074d3a249bd8b308574c95f7f8ddcdbceeed90478c224808460e808f5960001b613356565b6118677fabb9e156d12ac7a66b109c336845d45dd7e6cd39b45b1b4fc084c8eb8390273460001b613356565b80915050919050565b61189c7f8228780cce9f4c68bbff62bf03fe6d8fe82f51e072661fba6f3f89ce458d35ae60001b613356565b6118c87f7cd4a549dffbe58362c6db3cfbb498a50d92182ba8ba764fc666eb343887068f60001b613356565b6118f47f1ba027032f68bdf2d6f49245c37c77e0a4461df7667ea28de5eb633917623e7160001b613356565b60005b81518163ffffffff161015611ecb576119327fb80f57b712096411b68fb70959ed4efbd3240365e002ebe630f4bf01dcf15fec60001b613356565b61195e7f1a78233a48b4688205736b49cf4572854e2e7ad73448ced19ee668b5118ec9ef60001b613356565b60006119af838363ffffffff16815181106119a2577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151610d7e565b90506119dd7f3fdf31873dbbb325217710a4a5fb88b98fb22fe0a1404cefe1ec29e2fb1e245d60001b613356565b611a097f805ec2bb86e7f77d8b63aa8834fcaf8f9ce197cc2d551fc93371c55b2ae5226a60001b613356565b6000838363ffffffff1681518110611a4a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151602001515167ffffffffffffffff811115611a96577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015611ac957816020015b6060815260200190600190039081611ab45790505b509050611af87f3833f126db8a9e9d7d9298be746b2a9b7dd478ddc1ca9b03c4e9c99030b031b860001b613356565b611b247f7d639eaa37b112e859209cd5546989d57a68d1b592a9d091be0e95adf7301af060001b613356565b60005b848463ffffffff1681518110611b66577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015160200151518163ffffffff161015611e2b57611bac7f18d664471c5ef35a60a0e72d2dc900fda2c1acd880e8bf365bfcdb719dc1896660001b613356565b611bd87f7a96271b90e3af2419a2c988cc999ddeff18758c2b03b26fe01ed08458bcaf6e60001b613356565b6000611c73868663ffffffff1681518110611c1c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151602001518363ffffffff1681518110611c66577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151612b62565b9050611ca17f77b50e56891a303a548061917eba237456e4365c56b1783cb95d852abd0cdb4360001b613356565b611ccd7fc56f93839990d7314d23c5da6ba78fd38b873c1af14390d50bd4699e645624b660001b613356565b611d7381878763ffffffff1681518110611d10577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151602001518463ffffffff1681518110611d5a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015160066160ca9092919063ffffffff16565b50611da07f650f53ad7e4c53b06462c84cb9da6b9229f391c565c2fbd1554b22db8a65bf7960001b613356565b611dcc7f044ed8a090255a6a7ac240412ea6100e50f1c9f4d19a0079b2c0636b393e87aa60001b613356565b80838363ffffffff1681518110611e0c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181905250508080611e239061a918565b915050611b27565b50611e577eb50ee83ca9291f0d1e5f0d1b13dd5c0b3283cde0212801c7b304014414bb8060001b613356565b611e837f5a0d5e38e7fdfe8a3ce2665b06cdd7c2b1cc43f0b34500909fdadeec712cc00860001b613356565b611eb58282600588604051611e989190619e31565b908152602001604051809103902061662c9092919063ffffffff16565b5050508080611ec39061a918565b9150506118f7565b505050565b6060611efe7f3700605baaa596f56e3aa9f457e94a1cada9e96aeecb50881015d6562e07bcac60001b613356565b611f2a7f205356761809f828a03ce79a12a10cae306418afe294857b8ec2ebee66a4609260001b613356565b611f567ff1437bc0abdcf149f88b5f2111f1c181a3bd1fddbbb50ae9b4e8a8443938cc6560001b613356565b60008260000151836020015184604001518560600151604051602001611f7f9493929190619dcc565b6040516020818303038152906040529050611fbc7fc8d92ed5ed8e7306abc374362b3b4d717f26914defd7df5ed450f4e3a09b3d6f60001b613356565b611fe87f398864ec6006bf32f10661b7168d27f93c70eeb576d8b9a316e6c6d7da80ef4b60001b613356565b80915050919050565b600060019054906101000a900460ff166120195760008054906101000a900460ff1615612022565b612021616b24565b5b612061576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120589061a032565b60405180910390fd5b60008060019054906101000a900460ff1615905080156120b1576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b6120dd7fa229b811df4f66d6c58e173f2b8a07c9b0e203528a8034af2be8bae9497b57cf60001b613356565b80156120fe5760008060016101000a81548160ff0219169083151502179055505b50565b606061212f7f95f20bd4fdc12c942b47502eb34078e077a2f550655c5adfaf2684f4a311f3ee60001b613356565b61215b7f8229e54e47c788b6e2856308be5d2ca9e678e6c3a5db9bb8c1bf345fbe77082960001b613356565b6121877f44c36570451a6a2c754f72a00a4a5d70b664f54fac451916afa748e5a613c84e60001b613356565b6000826000015183602001516040516020016121a4929190619ebc565b60405160208183030381529060405290506121e17f152fbf788a711d25f94f5131d0424b938d0b9ec130ec78c0c42f5d54f1223c1360001b613356565b61220d7fb92a467eb89bb0e7b6fa8c1c9f0a375fc95dcfd0376acc155772dd81cb8df8b560001b613356565b80915050919050565b60606122447ff60c93b3ab36fc025ac656854ed959aa01dc0bce8987aa79a46c1d3f4a9d6efa60001b613356565b6122707f644065295cecf9e868907c4ce53b009298384839edd67adb8800405f4b93daa760001b613356565b61229c7f499b831e88c5d3779afd4c45c26402dfa1fbc615e9da10a42d56739f323139fe60001b613356565b60006004836040516122ae9190619e31565b90815260200160405180910390206001018054905067ffffffffffffffff811115612302577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561233b57816020015b612328618015565b8152602001906001900390816123205790505b50905061236a7ffacd226dec52ebd4083513362e3a95165771df74593c98d17b595d74bb00e3a060001b613356565b6123967f74341ea022293690010cc66f74a9efa278b0034890f6fa73befc18f5823a26f160001b613356565b60006123be6004856040516123ab9190619e31565b9081526020016040518091039020616b35565b90505b6123f1816004866040516123d59190619e31565b9081526020016040518091039020616c3790919063ffffffff16565b15612558576124227ffa3b94b3d3ffe02b89b515b6aeaabb6756e35097795129d26a31b4cf0f63cc8860001b613356565b61244e7faa7a44493f715a86d42315117040d2fd8c87be60a72a906dc8e503705020076d60001b613356565b6000612480826004876040516124649190619e31565b9081526020016040518091039020616ccf90919063ffffffff16565b9150506124af7fdc445fecefec329bdbe54b406eb327b96efdbd00faab0315a1a6b75fc694b2d660001b613356565b6124db7fe922cca990c9890a0be4e99b23f8d7ae8948bb2b6a76a2609e7f24c67f17348660001b613356565b80838381518110612515577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018190525050612551816004866040516125359190619e31565b9081526020016040518091039020616f1090919063ffffffff16565b90506123c1565b506125857fc2ec75f28a1458b53f13c48018d29b003a7421a35d64faa0226af5bf87b7e13560001b613356565b6125b17f3a96ac1e11cac5049b166960973e84388673bbc8c73f463e4099816bd54e736d60001b613356565b80915050919050565b60606125e87f08153d73ba4a053d4394cdfbe6fd66013dda662fc9522b68ff45b9fbad3f1dc360001b613356565b6126147f121d9fccf665f4b17a436ec4e91542e62e0c9cfa36b58f70a16ffd83af10253c60001b613356565b6126407fc71d8b06fd52744f8ca8290f02ed69e031c8ef42f0003139a31f91aa72d64bf060001b613356565b606061266e7fc4151949408fd0bee0f0a2164c5b6ea2be670a3fd5001d35e7633b98b0582dc760001b613356565b61269a7f45802ecfa93b576664690d580469380099e07d1893ad0e1a62e8becce30f3f6260001b613356565b60005b8360400151518163ffffffff16101561278c576126dc7fc2ba9d6b59afa180384132b5815ecaa3cf19d1667b45263f5ae230a71a7e334f60001b613356565b6127087f75dabb0fefd18c88c700adccff6fa75c5c29c64cbba5dd46b2b73288051f4f6060001b613356565b8184604001518263ffffffff168151811061274c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b604051602001612768929190619e48565b604051602081830303815290604052915080806127849061a918565b91505061269d565b506127b97f40f88f1e5a9c92c4da8af702422a0c2f0da7106c8a02370001fe110f57cd7e8a60001b613356565b6127e57f0190fc62b26344bc0b9a007a382615a99b1b525c0421a2304cd990a839dd299960001b613356565b60606128137f47aceda595d02109e2304e8518fca55aa5953676fb98ba73e718958b8f04a57260001b613356565b61283f7f03f58f7ba898d802637de9803f7582a6ab6e9805bd7fd3bad21db9fb1f3c39df60001b613356565b60005b8460600151518163ffffffff16101561292e576128817ff137b1d53ffa25df37b0f65cb865b1891373f9b151e8750b61e88eda55cabc5b60001b613356565b6128ad7f85218afa6fe9cd90353f58982b980b5d2648ffbf4a59fbb12bdfc5844162ccf960001b613356565b8185606001518263ffffffff16815181106128f1577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015160405160200161290a929190619e70565b604051602081830303815290604052915080806129269061a918565b915050612842565b5061295b7f980826820b4e0115ac83cf41fa3b0a4118862d361cad398cf66dda46461f104260001b613356565b6129877f601d2f7f7389883cd6e61292238d0f602415d9f9198349bb8ece94dce9c8e69960001b613356565b600084600001518560200151848488608001516040516020016129ae959493929190619ee8565b60405160208183030381529060405290506129eb7f250bab79cdda3dd8bac1a189748f2a4cc90b416d29f5f11af948ab4c20a31fd560001b613356565b612a177f952be3f95a465c341a4fb54f7307b6e26aeb15f7e8f4f1a4b8feb0aca454376960001b613356565b6000819050612a487f1b2af4e046d94a7007c0cdd6a68b42fb311a333685c2c968112fc80b32dc34d360001b613356565b612a747fb8258761d47da97aad07e11720425d6cbdf64c6956625160d119e3857bd7951160001b613356565b600081805190602001209050612aac7f0dbdbb6e79e08581cd88b6dc3fb4ed8ca455fa1672dc51073c1a50e4db1e7bd460001b613356565b612ad87f777a42740c016407f391d9fcdc81c515595fdc5e7a243473c6302cafef0c326f60001b613356565b600081604051602001612aeb9190619e16565b6040516020818303038152906040529050612b287f06d1a3b7e04efc8c2b2c8f61705657d03eb2761b96974054f8ebfdb7522e06d460001b613356565b612b547fd04d99867ba2c52308262ac6cb5784030e0a4f055c54e9638f3adbf062f3e0b860001b613356565b809650505050505050919050565b6060612b907f606c0a54fbe1838c989873776db0dbdcc9c7c860dab1d8316c43e548cd55212b60001b613356565b612bbc7ff3447b213efda5f804ac50ec13d4cc2defdb2899ffd686256a3da7a4bead9e4560001b613356565b612be87fccfa8bf01cdd4dfe33d6b5e15b274c0cc4b61b6e50f4de4b660ca037e586919560001b613356565b600082604001518360200151604051602001612c05929190619e94565b6040516020818303038152906040529050612c427ff09706b9ecd7eda7d0b08726f53826e20b97bb34533379a42f063dc3e199ab4c60001b613356565b612c6e7f630661fb54e7a62a56c0d51e2d66a0c8aa9e524b85978fde6348204613984aaf60001b613356565b80915050919050565b6060612ca57fa37cb037c2843c4b1c577022e6666ef82ead9ba7785d6c7a68f57d6c05be72d460001b613356565b612cd17f90ede87fc5a3d9a5af03555fad0903156e5f3ad743e2e71667d97dcf42b6838f60001b613356565b612cfd7fde192b4b63a9375da38adefe9aec5cda77155162a4bd04887c19e803ecd2814760001b613356565b6000600583604051612d0f9190619e31565b90815260200160405180910390206001018054905067ffffffffffffffff811115612d63577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015612d9c57816020015b612d8961803b565b815260200190600190039081612d815790505b509050612dcb7ffc1edd5bbb1acda7bf34f0f585bdd98d172f65d0dea30503cf5062a17a1eedbb60001b613356565b612df77fad7348770d28e369fc3d967a634aadec4562a9d3214d220e5df96ce377610e1f60001b613356565b6000612e1f600585604051612e0c9190619e31565b90815260200160405180910390206170e3565b90505b612e5281600586604051612e369190619e31565b90815260200160405180910390206171e590919063ffffffff16565b156132f457612e837ffebfd8305ddae9ca391ed0476d26c8dfcf7847f0c53fe56eb8d1e3877bc2741860001b613356565b612eaf7f16b9e675f11f4ada6a3e2092c96e4f64b64f06590dd29740bdb9d6b533884a3f60001b613356565b6000612ee182600587604051612ec59190619e31565b908152602001604051809103902061727d90919063ffffffff16565b915050612f107fcc2ffb622e2a721b17e94d079b822618e506c0a5e440d0fd626c357688860fe360001b613356565b612f3c7feabd8147fc8f81421bdf4d8b4270392cd18ed6be648eba4a172e0ab0eb927f1160001b613356565b805167ffffffffffffffff811115612f7d577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015612fb657816020015b612fa361805f565b815260200190600190039081612f9b5790505b50838381518110612ff0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020018190525061302b7f96ef178e4a5574ad4f213e75f2b5e568fb5675034cb24ca7e1ed08c3eb5f1b7160001b613356565b6130577fc3f368abfbe163764bf610aef5c7c952d31885b7cd4313acbe9f71ddd5a8ca0360001b613356565b60005b81518110156132bb5761308f7f1504792ed1147b567976a18b3910bbed17635006b85fb6af86e7a7d1725e7db360001b613356565b6130bb7fe8ae67aa2da7716f299179909db20452b329468a37ee2978157fcc0bbd9033cc60001b613356565b60066000018282815181106130f9577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015160405161310e9190619e31565b90815260200160405180910390206001016040518060600160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201805461319e9061a86c565b80601f01602080910402602001604051908101604052809291908181526020018280546131ca9061a86c565b80156132175780601f106131ec57610100808354040283529160200191613217565b820191906000526020600020905b8154815290600101906020018083116131fa57829003601f168201915b505050505081525050848481518110613259577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015160200151828151811061329d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018190525080806132b39061a8cf565b91505061305a565b50506132ed816005866040516132d19190619e31565b908152602001604051809103902061753090919063ffffffff16565b9050612e22565b506133217f2d528651833ec73c1197b943836d5fa59df3d2a8dc2faa2d062c08161458cf8b60001b613356565b61334d7f7ad64c90ba12d135128451ce69c5e571ca13b86af87172fd600976209212527360001b613356565b80915050919050565b50565b6133857f95bed97b996bd8a92dafd0a1af7981e0f45abfa9e09b83217eae5b05def1644260001b613356565b6133b17f305940d7b1065ac58cf237b33db79301a734ef34d61ce691c210340a513a059160001b613356565b6133dd7fd101c9e55bf9438eb693067f349af7ea81b0a15d236cbe983a33a618072c993460001b613356565b60006133e8846125ba565b90506134167f93f5fb964c4f56c4fd2d8c21edf43e9a5b43105161959905b2107733cbb8a74860001b613356565b6134427f3f527fdea596ea7962b12bd9bec9f70d54303342a9d59907a8699bad886c7dba60001b613356565b61344c8184611001565b6134787f32f09a15221a86543baac8d2f9f75c8d5623f2a09aa0130dda346d0ffffab17d60001b613356565b6134a47f598dd765ba7222f2b8d1dc264f0cd304bd7d7fc8379fc826e16d0020823d081d60001b613356565b6134ae8183611870565b6134da7f3e2dd7c13b965033695e62df30e0468ec135d06f238c19c212890b3a3c6876b260001b613356565b6135067fed04265c7c5cf7d0c9a06ed5bb7bdfd9d17c0bf7c6be0de81d27040222e545d360001b613356565b61350e617fa7565b61353a7f27a4cf058858cb9ba34ec2bd4333ad4df4506042fbb8a57910b46a469342707760001b613356565b6135667f1912323c8d862490705e6e9fee154cd01f8f866e827e00b9bf362e5bafe6370660001b613356565b60008160200190151590811515815250506135a37fe3a9fa61b982ae3f95684ec2d737731a31fa70c42270391489cb39956da81fc260001b613356565b6135cf7ff9cb822ecee4274d7be6b276508d206bcd62dacdcb2bf9acf4988bd992de8acd60001b613356565b6135e5828260016154a49092919063ffffffff16565b506136127fc84998d91b2710d739d8e5f78da2e7d6d1d4fb8d21869ffe7ee263d04a35100e60001b613356565b61363e7f8f960ebdb13eafca07b4930e1b473202a3cb008bc1e8d22bffdd13733030a25860001b613356565b7f7f6b6a2262c2fc2cc1b1785be448eae36656117d8f558facfe3c199e26256ea1600a60405161366e9190619ff5565b60405180910390a15050505050565b60606136ab7f90246bd187de13f550249ed4d92a6cefb1e0b56c5e23fa012f627ddb6829868460001b613356565b6136d77f19f67ee7fe904152bcf6fd97a37f427861dec33fd4a3913b1643b9cb667af68e60001b613356565b6137037fb3793d135d63823a045881b2b9c07b60cf748313d5b301eece0d53b2a9aaf12a60001b613356565b60006137317fad3f5cf51823b7889496904878daee9b01d9750a8db700d05208c281d94faa4460001b613356565b61375d7f5b317b2dd084f7d2f802cb0e647ff66b42441649136eac618c0164db9b969c6e60001b613356565b60006137696001617703565b90505b61378081600161780590919063ffffffff16565b1561390b576137b17ff793dc5b6a4fe6c3d73976270cb4a783d2d22f4af4bbf8c5179808da52e76e8a60001b613356565b6137dd7f8ce26e67a6c1228cdd2ff0dd737269f02689a891f00e7ce0dbc4568a9e5de90c60001b613356565b60006137f382600161789d90919063ffffffff16565b9150506138227f097a53f4f31a30e720b5d18d1b92c30d9c93ab56036d63e23b7caf2f19f7eb0860001b613356565b61384e7f8a982553283cc1eb78212ddd1b901ac249d1fc10924112f5861cecf5e369a00a60001b613356565b80602001516138c2576138837f64399d71e5b80f42f0f020d2397b4f09c6a92fbad71396c5417d49727364da5b60001b613356565b6138af7fe1cd1fbef05b4fbb2649dc4ab1c38aac5e56d3d02d570551549d695dcc0396d060001b613356565b82806138ba9061a8cf565b9350506138ef565b6138ee7fec2116f1fb280fcc19937f07e95cf151c3b8790342a381946a18fe07ca8ec90460001b613356565b5b50613904816001617d8c90919063ffffffff16565b905061376c565b506139387fbacbb99b5bfac579329a50bb311b0187abb141d57d2b156c6e26b529e2e0c15060001b613356565b6139647fff2e6e7bd6231c49979d8e102007d1851b0aabe83e1fe75fa87e7e15e43dbe7960001b613356565b60008167ffffffffffffffff8111156139a6577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156139df57816020015b6139cc618094565b8152602001906001900390816139c45790505b509050613a0e7f861f2f396f548139f635af0595e9f70218b9ef441fb8315214aed9344c53594660001b613356565b613a3a7f94a10a6eb241cda59cbb1bd83ef7a7ca7246a3a37f4ae755398293a064785bb860001b613356565b613a426180bb565b613a6e7f015ac08ece14ef760e9c1ff3fc6592b895ccd95dfe378fcf4945c3e82fbc497860001b613356565b613a9a7fe8df19290d49bd92e6db0c22189461ae138dd4a7023a716b0ee234edde2b14f760001b613356565b6060613ac87fb206e18fc9a583a368c528b5a3fde2487b09bb9633d8d075ee6b27f267b8e0c460001b613356565b613af47f057186ba9d7be81c57a287e4c6b3143f94f8060f753551dd6a648d84baf79f5360001b613356565b6060613b227f48b83cc4fcc123a3e2e357f8ab3d9fae50e9d9ee0f369bcdcb0093fcbb3fabdd60001b613356565b613b4e7f9d5d4b9e404935787287dea97bfee32f061720461eba57bbe20e2ca27b81d4cb60001b613356565b6000613b5a6001617703565b90505b613b7181600161780590919063ffffffff16565b15613f3257613ba27f1293f1c01d24eedb3a5a30f57952f2b3985ef85ff81da531d28f71dcc921c11c60001b613356565b613bce7f86529c8d299e9f6c44de9986adfb10531e7c4d5dfa73cbfe45839605c037af9160001b613356565b6000613be482600161789d90919063ffffffff16565b915050613c137f54261b2326b5818be3162991584764b0666146fbfcc2a345ab84bf8cfb32a97760001b613356565b613c3f7f6cd5e2ce50291b457bbbd5e78dc7a10082f134b07c47d6c9ad9043423556c85160001b613356565b806020015115613ca757613c757fe80a097df90e76205028fcd3c8a55760580e471f1912c693c0e51a54287932db60001b613356565b613ca17f69a517b545dc65113c96778be43ee20546edd00b340c9fbd6e62ddd5c72221c760001b613356565b50613f17565b613cd37f6837f03823eaa44442e45a88b9c18c36068b7db7e5764c8dd25b139d9089369260001b613356565b613cff7f8ebfbdf0cd13c61fd32cf5f94acf3f933a3a5bd410b66f1be8e0961add619b8e60001b613356565b613d2b7f68a75fcdb4b1642a462663ef29febe2909687d7e07d37daf9bd23a10863af81e60001b613356565b80600001519450613d5e7feb4e2957e60ddaaa1519187c71fa8d01e02b711a1712a9c06dcb360194b7f1e260001b613356565b613d8a7fe75a9662752de24219e7cb601686120f74abb78f0ac8c5a24411558a1168b17e60001b613356565b6000613d95866125ba565b9050613dc37fd73500ec049ce6b5f2500dfb690f6ec5bbd876a75525cbaab32c4e40eb84c68560001b613356565b613def7f40a57e82587c0debe6478e3e001c807c42167c2c80e829f0acf4d4dc29f1f0cd60001b613356565b613df881612216565b9450613e267fe42f0fe697251b7923948743324394c33212b9236d2ed6787e09d38adbeab36160001b613356565b613e527fdce7c8fe98651d65ad3cfeaf1bf117d5dcad1dd7133587f26ab0883998e36a8160001b613356565b613e5b81612c77565b9350613e897f2a61f8c1a9d47e30cd7c4c14f5b49c708d30dd04cd5ee3d76cc2f343fec10ef760001b613356565b613eb57f918a85d90d0d616c72c71c6aab98153138d0bb691c130aceb4df17af201930d160001b613356565b604051806060016040528087815260200186815260200185815250878481518110613f09577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018190525050505b613f2b816001617d8c90919063ffffffff16565b9050613b5d565b50613f5f7f8d5b5d7cb1862da20882a7cf420f6e9ae13a9ef72152a15543d054164e606ab060001b613356565b613f8b7f0cf686f501fd87e6a3dd12763a5fc56f6250890cb20016a3787a4766140a15b460001b613356565b839550505050505090565b6060613fc47fa486626a3500ae7e68689ee33e12436a801f3876fe204bf504ed90523ee2342a60001b613356565b613ff07faba923875ab1d03d37c339081cb85a78d5c83e1926328953511a7df4f1c938be60001b613356565b61401c7f68e2fed54d21b188a8be504c16cd69a43bd31058177570c82a83351f02c1adda60001b613356565b6000826020015190506140517f3505f2c01d6564160c25231fdb932202e59155a7c41cd1e9d3a3cdaa44f103bc60001b613356565b61407d7f87bb72e8ebd1dbc8746b793e62f49a73ac0eb861d4a5bc7dada92cbe61cc0dbd60001b613356565b6000836000015182604051602001614096929190619da4565b60405160208183030381529060405290506140d37fa6326939f0d38d9d4ab8ade05716897c0f202b4b6d59de333222301ef384521360001b613356565b6140ff7f8daa746192148b3bf87fa73d3ef30d13d222a7aebfa14570525468037323563760001b613356565b60006002826040516141119190619e31565b602060405180830381855afa15801561412e573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906141519190618ed7565b905061417f7fbf8fa58349832a8dc7d0975fcbee127fa25abf137521c8bbf0d584290c05a3ef60001b613356565b6141ab7f718f7aa9fd38a22f9bd24e5a12973b34ca36afdf872fb696a8c8cbbad6466b1660001b613356565b6000602467ffffffffffffffff8111156141ee577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f1916602001820160405280156142205781602001600182028036833780820191505090505b50905061424f7fcfbed26571f978ed1de26af02f06a8cd7defd9dcf94ae78bb25306008b03a2d060001b613356565b61427b7f3b8bd05e0230d114b33baf5cfd2124a8c52d5898aaf8e7fdf95b368aee3fe10b60001b613356565b60005b60208163ffffffff1610156143a9576142b97fcb3f9b32029e2e86ea5745418e50c135230baae028b9af289cccb2b32335837b60001b613356565b6142e57f6433135c887a3bcaa6e60aff058d210239ad89dfbf9dbad135585a50aa54989f60001b613356565b828163ffffffff1660208110614324577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b1a60f81b828263ffffffff1681518110614367577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806143a19061a918565b91505061427e565b506143d67fea6e2d08122abb3d8a7643c9db01310983c8ce73b409c458bc991430aa43e86860001b613356565b6144027f35e358c5513b5b4995183b2e5c6ca93ae131af0fbd2fa7bd233494884e17eaed60001b613356565b60005b60048163ffffffff16101561453c576144407f8b8952852bd002edfb0974eb22ce763f95812bbb2fa897d0910f3ec1a396a19f60001b613356565b61446c7f64c7b34f3e8eff86769ad2d13a4cd3255bbf7e511337b254abc712fca0d7665d60001b613356565b828163ffffffff16602081106144ab577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b1a60f81b826020836144bd919061a3b7565b63ffffffff16815181106144fa577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806145349061a918565b915050614405565b506145697f40311cca264e4630a1be08dc1f6f28ac3658ae6f3ee7d9112933448fc0ca834460001b613356565b6145957f9105564755410f7ce646845b40e7e56e87f269c4af90ef2888a2dba72d254c4760001b613356565b60006145c37f0b3e1807ee979df0d7122f810fdd2b64fdbcde8695dc9c644a68d534ecb1af3c60001b613356565b6145ef7f10063bd89b68b170f9523dbcb0890bdadb3069c976202e5a7e0571a8f7da6eec60001b613356565b600061461d7feae4d3e24bbe7732496288cbb3aa525fa908f107a71ffcaacf63483cc4e4a0f360001b613356565b6146497ff3b30201f5e2bc182531ecf6859fb7c905ad6160f68860aae0bfd4bfbca1069360001b613356565b60006146777fdf418b10a814e6f2422e5dfbf24afb96b0fc9c7d40a195ccbfab9418b8b72f5660001b613356565b6146a37fabd7b0929e32f1d8f4e34954f97ed7bbff47a83788b0183521ab7f8e52e1f68b60001b613356565b6003896040015167ffffffffffffffff161161487a576146e57fe975486f86488a1d52e796ba008c197bcf01706c444e766d53f6ada53438566a60001b613356565b6147117fc17cb80bc16a5c2e34e5576a72210151f348f213accbd127c669e85d6f27049d60001b613356565b61473d7f502ef7d1db4da6534fb4c3b727ee0ed19337bc52d893c29d235c898784a0b29960001b613356565b6001925061476d7f04e669c20d21a5fe6ead64b2fa697375795cc49837e56ef5002581a9aad0cdbf60001b613356565b6147997f5d6c937cc651fee6993a1edde60580361873b2370f2996b2c6d43876e274dd9860001b613356565b600191506147c97f638e0b69fd3994d0fa07d3d50fedb2bc85150cb6835390de37b5dfebf6f695fa60001b613356565b6147f57f4e723b30ef1fa233fd6cf50ab36d733f26d823569cd6bc439d2431c4cd425b3360001b613356565b600190506148257fc89902327ca252dae801271b8953703753d54e03dd9158267a00a845c9e063b860001b613356565b6148517fa53a479e0b71559d802d125665f0e62e8f2eab598ff3bddaa2c1248311e75e3c60001b613356565b8860400151896060019067ffffffffffffffff16908167ffffffffffffffff1681525050614b56565b6148a67f0c941daf34a735779a3a91fab28c8e1ffaa352c1f0ac88d6066238647de25c7760001b613356565b6148d27fb2b62b843684ceb2d9831eb530f632b7536f72307e1c0b76fd5634637b32cad260001b613356565b6148fe7f3eb6e351bf2b7b1b53608a9326f27d510af7821585e6829b416d327dc4e6d0d660001b613356565b6003896040015167ffffffffffffffff161180156149375750886060015167ffffffffffffffff16896040015167ffffffffffffffff16105b156149e6576149687f35085e10299efeee0ee1a72bc37be672bd61c0af49b2229e4df8034a3b2aa8f560001b613356565b6149947f45a8aa7219beb3088408e87f39697efe03fc553460e64896785ed11e0b450efd60001b613356565b6149c07f8d0a072d9748a7e83412b919e560562b14f08f5f7e28594930896d2b9845c8c660001b613356565b6003896060019067ffffffffffffffff16908167ffffffffffffffff1681525050614a13565b614a127fd488e42b98a33c529141a02d809f08467ff62f3875805a3e0cefa31b34f96c0f60001b613356565b5b614a3f7fc895c1b1c319c04c6ad0ae31c91a69cfb47d50b919effe4c715acc6396586c5b60001b613356565b614a6b7fb45e357356b3a53d67174a0e41fa0edc1e7a7b4f72b5c79aaf07b1a5b1186f0560001b613356565b88606001518960400151614a7f919061a466565b9250614aad7f5aaa303bd6704909330623a7d13d2c2b58857e882df585122c3217952a3c9e3960001b613356565b614ad97f910c48224b566050d6f05efba029fcb20424f547f93e494ea8d78fb8963a66f060001b613356565b88606001518960400151614aed919061a9d2565b83614af8919061a3f1565b9150614b267fda03c328f949a70e506afede871dbebee231ecb7ec37ebc81bfbc5ada803dc4660001b613356565b614b527f3671ee0b0b75b988c779cb8d97e26929aadf25b1b9c06bd6e7ec56711de44d1b60001b613356565b8290505b614b827fad593b758a22258132b918b3f0416e728be2d60ddc8cfc336ef83188d995382660001b613356565b614bae7fcbf1709381319ba3ee54ea2c0afe7e6135d58fd2951b70c84537de4a0f250fac60001b613356565b6000896060015167ffffffffffffffff1667ffffffffffffffff811115614bfe577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015614c3757816020015b614c24618015565b815260200190600190039081614c1c5790505b509050614c667f7346601aebd6d9c922f951e93aee56d9521856a979387991fdb6322c79d71bcc60001b613356565b614c927ff6ea27dc853e99769d39e128a032e533755e3c480331e4474184ac2108b104bd60001b613356565b6000869050614cc37f01fbb2d227304ea9c2a67ca5278e2008d313a611bdd06204e0cf1c0b220f387260001b613356565b614cef7fe2b9fa4b37fa8e0ec2f8f11e36a61dde561e59c36464570dcd8024b8449e5fa860001b613356565b6000614d1d7f957bbca5d2acd13007ee7425419f50c25f377b1f03b4852db5d7a4e50614f3da60001b613356565b614d497f53bba485906fd1f4aded578041a52d70067910c46ba16d90795da7ab254b080c60001b613356565b6000600190505b8c6060015167ffffffffffffffff168163ffffffff161161543957614d977f226d3444d122ffecd53372c411574efbee93933be0a46779846a5308055d809260001b613356565b614dc37f8750f8757a2e6b2e6bd0cbb6a3613a1604e2b6b85200936204638cda1d80ebae60001b613356565b8c6060015167ffffffffffffffff168163ffffffff161415614e6b57614e0b7fa13c409d162ac2f082f5bea49013bd2e894fb2fe326e78d650cdb403e823dee760001b613356565b614e377f32312ea2687113ab56c178cafba63b3195ac6333fd83d89eaef1c930d0f33c9260001b613356565b614e637fca45358cee04bd0042e0dfaba7cc0217a943fd28ad54464b5c2ccf1dc683644160001b613356565b859450614e98565b614e977fc63ec981427c4bd70f66bd01d9bfc5fd0f536fcff5a6d357abe9a8573807f52160001b613356565b5b614ec47f66949b8f0b33f97bf38fe2b63e6f8b4b3a746c7952a948c4097a9e79eda86ceb60001b613356565b614ef07f7aa237a45705ac37364a95b2ddc95662e4347b37688c240e2abf969b48bd250260001b613356565b6000600467ffffffffffffffff811115614f33577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f191660200182016040528015614f655781602001600182028036833780820191505090505b509050614f947f122db8a9118cf194c3a8d0e600bb9461ea6af80accf2fdf9554d2db29ccb1faa60001b613356565b614fc07f628df02cd5a2c24772beed64eb3d5b4cd11f9a4067453c27209906e1d523d29160001b613356565b60005b60048163ffffffff16101561510057614ffe7f6dde327f73f3be4a8e08bb358c7e00820a44ed1ec6b3f6c5fa5aeb4e8233b47f60001b613356565b61502a7f610298c5dea83f45bc2fb153e2a95c53a1830d40928c38c4fc4b602ec71bdf9a60001b613356565b898185615037919061a3b7565b63ffffffff1681518110615074577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602001015160f81c60f81b828263ffffffff16815181106150be577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806150f89061a918565b915050614fc3565b5061512d7f822db6a75d06bd5ef3e50a50772732e235246b2fef294cbe72a84c8553736bd860001b613356565b6151597f5f833927560259102fa390c44712d1bd534729c789a6320b4c96ebbc8a23e2d860001b613356565b600061516482610500565b90506151927ff898b4d52fca48b04a3b7cc3639196e6dbd03caefeb51054319f04d27adf63e560001b613356565b6151be7f7a523026173c56a626a993914c74ebc3fad4bb514004da55422e6a3858b880c560001b613356565b886001846151cc919061a6d8565b63ffffffff166151dc919061a662565b876001836151ea919061a3b7565b63ffffffff166151fa919061a9d2565b615204919061a3f1565b86600185615212919061a6d8565b63ffffffff168151811061524f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516000019063ffffffff16908163ffffffff16815250506152997fc7475e4547abd8535159092a87c875cb8d43a3ab3326e6d43bb9eb009f0fbd6260001b613356565b6152c57f6bdd2fa6d1f791ed793e457946567ae8fc7ca5811875ae21c44fdc90c1911f0c60001b613356565b6001858563ffffffff1660208110615306577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b1a60f81b60f81c615317919061a42f565b60ff1686600185615328919061a6d8565b63ffffffff1681518110615365577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020019063ffffffff16908163ffffffff16815250506153af7f79d21cd1a83de0a09d151e0ee9efc3c128858769b1ada6bc14327163a0f8a39f60001b613356565b83806153ba9061a918565b9450506153e97f27abf0a3d7cee3cd0f717280d27397ddada84fdcd0094939c2a6b8a79959360460001b613356565b6154157fcffde9eccacc2deee77a3a0ead4f5aa0ab031446bf3f777ec15604a0b94d08ba60001b613356565b602084615422919061a9a1565b9350505080806154319061a918565b915050614d50565b506154667fd812030ceee15aa8e8c6784308556c4e255504335e36cb0359543c848cf26ed360001b613356565b6154927f236413256bcbfdca3c471fab2368daeb7a5b09e0426f8f985e5765e79d634e3a60001b613356565b829a5050505050505050505050919050565b60006154d27fe3599aa9866cfd03a867a2e97b63ab0b9eec44deb8ee00022d6aa0366a56649f60001b617f5f565b6154fe7f27ccab27f3553fd26fa5890d1898bf38eef4f71e3204e610b8a97d7f84d7b42760001b617f5f565b61552a7f873910f18cdacd497f49322f2b2d620c780c518a35c07391e3d911284ba37e8f60001b617f5f565b6000846000018460405161553e9190619e31565b908152602001604051809103902060000154905061557e7f594350d6d06161fe464b9abeddcfa9352e5bca993d0a4e8e5f15f70a285b2f9b60001b617f5f565b6155aa7f442d4ff95aef9f651bf908e237f033c0c980294426189587e256d18da715485b60001b617f5f565b8285600001856040516155bd9190619e31565b908152602001604051809103902060010160008201518160000160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160010190805190602001906156229291906180f4565b50604082015181600201908051906020019061563f9291906180f4565b50606082015181600301908051906020019061565c92919061817a565b5060808201518160040190805190602001906156799291906180f4565b50505060208201518160050160006101000a81548160ff021916908315150217905550604082015181600601559050506156d57ff6e491192a94810eaf4ced274cda08b1585f6f23405a4c293e044d269b80e54460001b617f5f565b6157017f66fd6693224a99900a0e4c5def4f1e7cd1c8b30273a7f557a4a7af8324f7534360001b617f5f565b600081111561576c576157367fe4e08727881ee50fefc30e2bb5746740f9c77f2543b4134eb8f2e2485257049660001b617f5f565b6157627f3f507d363dc820267a8186002018055a3f7d94ea4b524df1d630bcf25f48d1a160001b617f5f565b6001915050615a5c565b6157987f42bfb13fb8cdbc67ebe0205cec8c1aaa05fa0e80808cb3e2cb6664b7269f1bb360001b617f5f565b6157c47f6dfea76d801421108a1a759338ba0262f0057a49e353ef92f9c16127d8677e7d60001b617f5f565b6157f07ff3a95e80748b7546658ec98e3c92d2dc79bcc39dee07a94dda156ab3f476839460001b617f5f565b846001018054905090506158267f73c462a8bdf2952e850d8d2cff05fbc10a5a2dab94dcfb8c1f5d98e1b68d291360001b617f5f565b6158527f12ce1b49014235dee106ff533c9fe94a82612e60dbf1c489dc6d588587d7b05d60001b617f5f565b846001016001816001815401808255809150500390600052602060002090505061589e7f3ececb0bf5a71c495b960c2ac2dd4a5ae2821e25ca105aec9d8f63f8988e810c60001b617f5f565b6158ca7fc959a3960414006c558e1e7fb94560439f94323719e7b2a5296d94ea3eb3be5060001b617f5f565b6001816158d7919061a361565b85600001856040516158e99190619e31565b90815260200160405180910390206000018190555061592a7f5bb4d5d28c3d6ce7e6720dde2e8362387539bc4d8f5fe6afe40cab9390d6355460001b617f5f565b6159567f05a36436f6c25a1b673b492caeca270e822c2bf50e5e5ee4d73e09c2a7587ec260001b617f5f565b83856001018281548110615993577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906002020160000190805190602001906159b79291906180f4565b506159e47f7f83147f36e4534253f70e375bd16a10a13491f12015cff20c7a09cebc4cd78060001b617f5f565b8460020160008154809291906159f99061a8cf565b9190505550615a2a7f3da7e1574f15a1691a51f44ca5b394673e73be7939d48bae73101f9ef51ec95060001b617f5f565b615a567f2c30fbefa2096b36355e0d1d386bee878474db38f09df58a4e4867897cda806560001b617f5f565b60009150505b9392505050565b615af981604051602401615a77919061a074565b6040516020818303038152906040527ff5b1bba9000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050617f62565b50565b6000615b2a7f6ef357f72959a55f84604fe70ffe73525d6ad10e918f31ab0ad0eefb2aa4d18060001b617f8b565b615b567fc7c2abb5171ef5f16531531d85ace34b596bb0e0554d44cdde90e214ac721ad460001b617f8b565b615b827f10b3ca98ffeb31a58ec949eafa1836b77fcf086a8177dc62f008f1abeaddc8fb60001b617f8b565b60008460000184604051615b969190619e31565b9081526020016040518091039020600001549050615bd67f804e1a1c44f0cecc7d0bb2b33338e9d9e3fe4e90509802164e42bfec45461ad160001b617f8b565b615c027f5aaf457aa3705b945a431f1aa1ab2f90912151004e1e9389a462f8641368ccaa60001b617f8b565b828560000185604051615c159190619e31565b908152602001604051809103902060010160008201518160000160006101000a81548163ffffffff021916908363ffffffff16021790555060208201518160000160046101000a81548163ffffffff021916908363ffffffff160217905550905050615ca37f087abe6c9036ffb33bc97265fcc119a6a4bf6edf436058b9a551da187db46f1860001b617f8b565b615ccf7f65f5d8ff0db0d762da8964aab3c3ddf662aaf7597738d919bf334bc3b58d45a960001b617f8b565b6000811115615d3a57615d047ffc77ba7df43f44eb1d9ae7b2fc947ca886692c525ab31dc3605dfa4be145fe7260001b617f8b565b615d307f08b7bf8cf9a11d9f2b0194fb04b6c2a4d1bc55a671bd2702411a7c9937fb83d660001b617f8b565b600191505061602a565b615d667fd1332c3553c4d4f3e28e14d1acc1db5936a2937c04ffb5bb5d5c3a48dadfb6ac60001b617f8b565b615d927fbf566f49f49c0af01635b258800a128c175ec266509f7a5c7548a64f509e3b0a60001b617f8b565b615dbe7f872562c06d04c8c83d2c8940a225bafc77a988cba2428c527dfe2d3574cb803b60001b617f8b565b84600101805490509050615df47f7a634fc97b24ef3f41b74bef6c921c16e46de5bb8416cae56d77e9e10198534a60001b617f8b565b615e207f60b06689f7a7fc0a179d66ec567ba9b9233c3a847c8b528519dd69e4632c82cb60001b617f8b565b8460010160018160018154018082558091505003906000526020600020905050615e6c7fcd51db06b5f1697854b22815ab3da0ac3925e934b24db4534741a156cb40392660001b617f8b565b615e987f81897fa430b25c069a5d98a18dc7c3be1ad7f694dba7bae8ecc8bb84921d921460001b617f8b565b600181615ea5919061a361565b8560000185604051615eb79190619e31565b908152602001604051809103902060000181905550615ef87f40139dd37e52903611a992cc4f6706b02ec59f83e572e6b6c35f0287e8ce4b3760001b617f8b565b615f247f0b33e57eb8e713e1a012b74e08a5faf3d00f3f7f36ff9583871ff9d78a3a9ad060001b617f8b565b83856001018281548110615f61577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002090600202016000019080519060200190615f859291906180f4565b50615fb27f3ea40f4330405680d8b7d36a3c85e059ba474865b9f808b8ce5db8c6c2e63a7860001b617f8b565b846002016000815480929190615fc79061a8cf565b9190505550615ff87fc559256d939ad4860f9aa1a9b9bde71f9a115c616b81ac4792fa89971245df4b60001b617f8b565b6160247f7c7528e6091c094ba275de2a328131ca76b462ab624c5fa447c077f73ecd7ea760001b617f8b565b60009150505b9392505050565b6160c7816040516024016160459190619f37565b6040516020818303038152906040527f2c2ecbc2000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050617f62565b50565b60006160f87fc0fd925fc82f36017ec6d307faa7fc2fd5246df9c1159809a00a73d3ea51f74160001b617f8e565b6161247fe90980aa2759b455eadc0f4b382aa14cabe81edfc9c2459ebde27a2424a8b3b460001b617f8e565b6161507fdbed26c2d5444a2b0da6c0a3dfe3bdb879d4a04fae53b52e5b1e2cd7a46fd00960001b617f8e565b600084600001846040516161649190619e31565b90815260200160405180910390206000015490506161a47feac0b0a1012b00cb2545db7fd1337ec0c0ac7673755356f532d66e5d6e40f78d60001b617f8e565b6161d07fad1ce68898f8323ba8d145f28e338f891623480a138f32e24d731092e920641d60001b617f8e565b8285600001856040516161e39190619e31565b908152602001604051809103902060010160008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550604082015181600101908051906020019061626e9291906180f4565b5090505061629e7f3b40c1265eda9c7c81356706a009493c0d07f3acc76c65ba3d7238dd443821c860001b617f8e565b6162ca7f09fc7e0b9cf5da0ffdded534e7c8ff1240b6e60bea44d4120027950ec801c86660001b617f8e565b6000811115616335576162ff7f5c9782a6b217d59527df4558e5fe8929e8cf24d4c00f589108b128511c20ab4660001b617f8e565b61632b7f5aeae2acc5dba95c9cf9022406fc416e2d75ed1bfe08688e11dd53906288473560001b617f8e565b6001915050616625565b6163617f0e69694489e2d0763238b6ece00ef6cb5b4419d1784ff355d1e15dffbc41150d60001b617f8e565b61638d7fa1d811679aec368407d6d7d6ad9d60171dfa894ee3838cd00c1929e8121f42a360001b617f8e565b6163b97f19299aec6e7041fea4186cb31bb0a0f9ea804b8316b66f706e48f4519ec0cdef60001b617f8e565b846001018054905090506163ef7f19a5d8760b619737388ba3e79b017c07d83b913d554227ab6c681c5b966b5a0260001b617f8e565b61641b7f58d0437dd4253c077b62fd65c3fdb9721972ad39f735938d03171c466236278360001b617f8e565b84600101600181600181540180825580915050039060005260206000209050506164677f31ac163ff66c1046de9ee7b3e682ee059f7825998122e4600386504f9049b74760001b617f8e565b6164937f6960cd16dd242cb6371a9d12ca19e419afa82eeabff79ebdb56d7ca4f52d5c3e60001b617f8e565b6001816164a0919061a361565b85600001856040516164b29190619e31565b9081526020016040518091039020600001819055506164f37facdbf58217aea94c90917baa34df080d682a25125810f011811674efd08baf1c60001b617f8e565b61651f7f67c96ca88d936b440822e9bcbc9db2dc12d558d45d1f4b4d50127b36d4f7fcd060001b617f8e565b8385600101828154811061655c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906002020160000190805190602001906165809291906180f4565b506165ad7fb60a049518e4b4623124016684ee24af5c7b8b9233607e32e5fde52ff5effbbe60001b617f8e565b8460020160008154809291906165c29061a8cf565b91905055506165f37fc23f57cfae0ccbdf6d8965e84cc8e69ee6df02f6304d00db74bee9f0c511cac360001b617f8e565b61661f7f32b79282ce3a8421e4c63ad410e823ec743a5d95e194a7320cdf967ec9bad7b160001b617f8e565b60009150505b9392505050565b600061665a7feca4f5f79de0d2c22d178eb2124d92d33b03c9005c55530ee4e6226f44063f3960001b617f91565b6166867f394635ae8d9187a3726ef6862c1f36d32d3225aaa61709f0c5eaf17dc4a9030360001b617f91565b6166b27f57ed5a03297a61f366edf4e3b6708ca1bb87f818c8587a8d34f1baa7984386a060001b617f91565b600084600001846040516166c69190619e31565b90815260200160405180910390206000015490506167067fa8d55023d88a815ee83ffcb7ebea811c389dbff5d5af0f78ef3f94db756fb68160001b617f91565b6167327f980f44685ed6c2f980c407e162aab029c62ed959c2a6cf3dc5a00c1aef46d3e360001b617f91565b8285600001856040516167459190619e31565b9081526020016040518091039020600101908051906020019061676992919061817a565b506167967f54c25574a793cd4a86452d23223fa7463dc99ca956200dde04e9c2c3ab05f66760001b617f91565b6167c27ffd38675d9a85e65af2fce2200155763c6e44688cbc5b078bab5bd0cd11a38f7060001b617f91565b600081111561682d576167f77f3954867c8654f129c1ac57c94c0db4d6e2cfb61e732b8ad04b8a45c6ec0bae7c60001b617f91565b6168237fe6a50bf4f53b0312bc415087c7ff7d209b13079ecf6a68cd3e3b15307df5b25b60001b617f91565b6001915050616b1d565b6168597f709be2b466ae3b6459e0ba80b2b1c4205ff4944da1e6f54791a8807e7244ef6d60001b617f91565b6168857f349f70d4cfa3c81178f92dc296754a99fc4947abbc6d135e6b9f4f27e46bfea660001b617f91565b6168b17f4a7b48c2da595fa01a2409f798bd2a7d7c1554c88714afa1ecc14119a0fdeeab60001b617f91565b846001018054905090506168e77f86c87f2721e2aa4eda18f7439567ef7d593d90559fe484f502fbe9517473cc0c60001b617f91565b6169137f939bf4630ca2499fc5869f0ef5dd2dc3c190643c688f201cdd83387dbcd7c38460001b617f91565b846001016001816001815401808255809150500390600052602060002090505061695f7f360ff9442d6e3faa2e412bb6bcf47be9be0226ad69d7bfd1c10c8bf638a6395a60001b617f91565b61698b7f1a8bf7dcdd526da69316eeb4cf773481138fa332815b828f2a5cd861d02d284960001b617f91565b600181616998919061a361565b85600001856040516169aa9190619e31565b9081526020016040518091039020600001819055506169eb7fdf355ae9d26eefbf11bc080737adea51ad45098b8c865cc6865b11c1ad4da3b560001b617f91565b616a177f75e449b107ecb9582c9a70d55323bc4e445e2c34710711a6c268cf09a6e472f960001b617f91565b83856001018281548110616a54577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002090600202016000019080519060200190616a789291906180f4565b50616aa57f45cfec0c088b9c43c82ef98db081c0e63d7fdee8ce2ee0541668079249a90dab60001b617f91565b846002016000815480929190616aba9061a8cf565b9190505550616aeb7f763e665d6e13d389d349d76ccf74602ef670597ddaa16541ed2f8df1ae73784060001b617f91565b616b177fbc14492c807df2b0b8b2255cb1257185806118878d691e9f37df17c0aba1590560001b617f91565b60009150505b9392505050565b6000616b2f30617f94565b15905090565b6000616b637f130ffc31a4104284bd94c19edfb657769b150ff7af0e4866c4ccbf7a120682d760001b617f8b565b616b8f7f91d843dfe844cc7e6e3e7529eae9d50feb63ab966a4f93e71263d01390f2b19860001b617f8b565b616bbb7f36e34fc65bcc7e1b6dcfb5fc9ec0c80e465afec13fb812510a0027b1bddc7f5260001b617f8b565b6000616bc8836000616f10565b9050616bf67f5ef0683b41aa81020f0cd79382dde3832c1b9be1cb7c720bb2b9f6898fd9714960001b617f8b565b616c227fd8a0050b38c41e5307f3929602252565913827b8ff37b4935efd97fabcfc307b60001b617f8b565b600181616c2f919061a6a4565b915050919050565b6000616c657f5e127334ed08e65d10eacdf29531a85b65e33dbaae2ff75ae51b9ed09a1b86f160001b617f8b565b616c917f901517fcebd04dd3124dba0d75834b843adcab4557c1cbb26071c9259e8dc54860001b617f8b565b616cbd7f0641386a5e2cc1910b28e27052259a3b01894d191de24dfb56e41bb8e789c9c760001b617f8b565b82600101805490508210905092915050565b6060616cd9618015565b616d057fd1777fa6262ea58903df49900ffcf1d1f0c202f19d8b7111627034ffafd1ac5f60001b617f8b565b616d317f039897155cfc6b50a470c7cd71725b506b0b92216f17bded5d252ee634e36b4d60001b617f8b565b616d5d7f2627cdf1ccd5af175ed6fe0ddbf236c9c1229634bcca6b4f03ffca0bad2cfb2d60001b617f8b565b836001018381548110616d99577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002090600202016000018054616db59061a86c565b80601f0160208091040260200160405190810160405280929190818152602001828054616de19061a86c565b8015616e2e5780601f10616e0357610100808354040283529160200191616e2e565b820191906000526020600020905b815481529060010190602001808311616e1157829003601f168201915b50505050509150616e617f0ef3278fb51e169301193c07b5fae232341bca2885661b7ac886741b14ea566460001b617f8b565b616e8d7fdc321dd175b9a9f414b68716bd9786b3a59de85aff5ad2765624371413f890ed60001b617f8b565b8360000182604051616e9f9190619e31565b90815260200160405180910390206001016040518060400160405290816000820160009054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016000820160049054906101000a900463ffffffff1663ffffffff1663ffffffff168152505090509250929050565b6000616f3e7fd9386dcb02406058385ad637acd6cccc487484466975953890e72c6e0cc25dc360001b617f8b565b616f6a7f977f4d8b7d744b9173e8ad45ec7eaf45f9a06611f970fc214f6087d977bf5dea60001b617f8b565b8180616f759061a8cf565b925050616fa47f805f286d2f14e1fc9e9f19f9e910717966dcac6758502457b342cc0fbb85bb0760001b617f8b565b616fd07f813ee58e8dc2616e8fbe19826f3942d4b2fa7d6d427b6800ccdaec71794f632a60001b617f8b565b5b82600101805490508210801561703e575082600101828154811061701e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906002020160010160009054906101000a900460ff165b156170825761706f7fba300d329c9b012a4d2482dc5a6b8c788e6e063838b464859b11c7716225838460001b617f8b565b818061707a9061a8cf565b925050616fd1565b6170ae7fafc41a651201ee79e0174d1c0b5e23caa3e5a1c5c201b6269e75b27684fa8c0760001b617f8b565b6170da7f17e59f04579f4a718fac0539632715365475c9bbec9bf2a73a2c020bdddfa03c60001b617f8b565b81905092915050565b60006171117f15c80879221e95d3b907f37c28604dcedc0343a242a565c4d3f8a6c0bf17252c60001b617f91565b61713d7f16232020e838f01034f251720322d256d094219318b2a05649df1aebf23c216160001b617f91565b6171697fe642033ca4eba80c1bc0106ab2b45d362e0ba573f919ee481217d03bf92aa74660001b617f91565b6000617176836000617530565b90506171a47fee1159216c898189dcfe21d5d3ab474a3a93ee75dcbcd0e6a80f1f6ba5eb155760001b617f91565b6171d07faddc10af81d371412c90196a2115862a686966f9b64b14e5fc069e3c2629d3f660001b617f91565b6001816171dd919061a6a4565b915050919050565b60006172137f9d18063e1a3deb17cb6536001c31b957c5b80f654e913e99c5783589ee12e26a60001b617f91565b61723f7fb1c2e084339635b43d98ee5cd2085d007c84df5d4f7b2c86cbf87c5feddd03c760001b617f91565b61726b7fd795eea709d34e25e6fdccdce68af067d17ee71b20f67216a9b245bfc92c1ccb60001b617f91565b82600101805490508210905092915050565b6060806172ac7f9c818ce95e7a63c05f310a4655052fd047762693a1c3b0798b08a8e586d61a8960001b617f91565b6172d87fe61daac294c54202209dc730515f0e5b0fe902049f7c5ff61594f39ac85b343760001b617f91565b6173047f441adbc5a32ea300121afb1c4c3cdd34b2a8b65eac43b4301047e875dee5271160001b617f91565b836001018381548110617340577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000209060020201600001805461735c9061a86c565b80601f01602080910402602001604051908101604052809291908181526020018280546173889061a86c565b80156173d55780601f106173aa576101008083540402835291602001916173d5565b820191906000526020600020905b8154815290600101906020018083116173b857829003601f168201915b505050505091506174087fa0f9a6b1a6fd6b1ec2edcb7481c90147fd0ed543458acc241d1e0f44bb4457c660001b617f91565b6174347f342b9747f9946503492f084bcb1c62df4711d16e27486b1ee3227e78f92ae13060001b617f91565b83600001826040516174469190619e31565b9081526020016040518091039020600101805480602002602001604051908101604052809291908181526020016000905b828210156175235783829060005260206000200180546174969061a86c565b80601f01602080910402602001604051908101604052809291908181526020018280546174c29061a86c565b801561750f5780601f106174e45761010080835404028352916020019161750f565b820191906000526020600020905b8154815290600101906020018083116174f257829003601f168201915b505050505081526020019060010190617477565b5050505090509250929050565b600061755e7f8030a7e3360b0bf8e5aeebc4a2b785617ff082e1351171920e4f57bc6d6ffa1d60001b617f91565b61758a7f4f1e58269379299398fe5a3f59ee3a9e2aec10d0deb7bcc48cda37bc01fcdea960001b617f91565b81806175959061a8cf565b9250506175c47fae2a0fb5cfb6678f2896e82d14331e111b0b4a893b2e6a9942c37f9461b542bb60001b617f91565b6175f07fb119eaeb6a64d6d1091cc1f39b3f18d5c03bc72f16abd0f2b10fa0f116f0e99060001b617f91565b5b82600101805490508210801561765e575082600101828154811061763e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906002020160010160009054906101000a900460ff165b156176a25761768f7f0ca6005de74b2bd76f75ec31df525f526a390c4985de9d8f6acc15793b3c7d4060001b617f91565b818061769a9061a8cf565b9250506175f1565b6176ce7f020dec87aa2c10fb42d7b67accba194806f84258eb9f8f0607d6a907a49189bc60001b617f91565b6176fa7fa77217cf99f0a358b3015950f4453c4aee24b87db18d3699be4b14b946a3aecc60001b617f91565b81905092915050565b60006177317f105dcc925039f971dff081840bb61a8f6f562872c3779752d0e3805624fdbed760001b617f5f565b61775d7f4ebe780020a8856250128e486383bdc8f0263c2b796266651538f1cd9dc81f6f60001b617f5f565b6177897fee531fb2dafb1a67cc890334c83bf7c8a2d913746b020314c6dc6b64256d129760001b617f5f565b6000617796836000617d8c565b90506177c47f9aff586a987c32c174a790d4fe5159a9e458f5f0b8ae91d48c0f5b0c44ea788a60001b617f5f565b6177f07f2e6aabb4354902f526d2224e582c34782fdb35031d57e80d33d1b2a592cecdb960001b617f5f565b6001816177fd919061a6a4565b915050919050565b60006178337f6264f85486d3a9a79dffc54f2b0fb9ccee55e725b23dbac4b8b04fbc017b395160001b617f5f565b61785f7f530217e7cfac26dacd62e72406ee1ea96f32ac6414aef86277349f17f48f21ab60001b617f5f565b61788b7ff2fb66e5c39d0342ff8d5d9f55d8910b1b59e03ab70035d9a48174e32187eaaa60001b617f5f565b82600101805490508210905092915050565b60606178a7617fa7565b6178d37f66076eb0a095b470716b3c2dc725bd6183799cbcb28250d9079601151463e77d60001b617f5f565b6178ff7f9f2b10463ceb327a81440a9974c1f0b8827b105b140b1d997f0e08a02f4471c460001b617f5f565b61792b7fc3bf4e2efffb3d38a7a5048a8a2a8ab435ce5afb1fe1b1b513714f0a8fd1f8b360001b617f5f565b836001018381548110617967577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906002020160000180546179839061a86c565b80601f01602080910402602001604051908101604052809291908181526020018280546179af9061a86c565b80156179fc5780601f106179d1576101008083540402835291602001916179fc565b820191906000526020600020905b8154815290600101906020018083116179df57829003601f168201915b50505050509150617a2f7faca6829264c00b50d4b9c59e9d602cac615b234202ed32d33c7aac9c1d1b511760001b617f5f565b617a5b7f76618b0bb6efb99f207ce739cbbc3891f64aa7d534fdff7e7fa3f0651189e9c160001b617f5f565b8360000182604051617a6d9190619e31565b9081526020016040518091039020600101604051806060016040529081600082016040518060a00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182018054617adb9061a86c565b80601f0160208091040260200160405190810160405280929190818152602001828054617b079061a86c565b8015617b545780601f10617b2957610100808354040283529160200191617b54565b820191906000526020600020905b815481529060010190602001808311617b3757829003601f168201915b50505050508152602001600282018054617b6d9061a86c565b80601f0160208091040260200160405190810160405280929190818152602001828054617b999061a86c565b8015617be65780601f10617bbb57610100808354040283529160200191617be6565b820191906000526020600020905b815481529060010190602001808311617bc957829003601f168201915b5050505050815260200160038201805480602002602001604051908101604052809291908181526020016000905b82821015617cc0578382906000526020600020018054617c339061a86c565b80601f0160208091040260200160405190810160405280929190818152602001828054617c5f9061a86c565b8015617cac5780601f10617c8157610100808354040283529160200191617cac565b820191906000526020600020905b815481529060010190602001808311617c8f57829003601f168201915b505050505081526020019060010190617c14565b505050508152602001600482018054617cd89061a86c565b80601f0160208091040260200160405190810160405280929190818152602001828054617d049061a86c565b8015617d515780601f10617d2657610100808354040283529160200191617d51565b820191906000526020600020905b815481529060010190602001808311617d3457829003601f168201915b50505050508152505081526020016005820160009054906101000a900460ff1615151515815260200160068201548152505090509250929050565b6000617dba7fe79bf4b510edd2604e95ff581da8acc7e3cce668860a1ec505dc80407c90919260001b617f5f565b617de67f86134b744847b7ea22025fedf702dba4c21473831ce6a8c6f077a48a8e6d7aa860001b617f5f565b8180617df19061a8cf565b925050617e207fe9464569b34f0d13f57d128bc72b6937bd102145ae61b5ae1ab948b872d1e8f860001b617f5f565b617e4c7f059f51b65c85bf170303fd45947488eb07bf50818d4c066aca88c87742b77c6e60001b617f5f565b5b826001018054905082108015617eba5750826001018281548110617e9a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906002020160010160009054906101000a900460ff165b15617efe57617eeb7f108a7e805c9fabd9c435445942989fe80f5c98c22c04929d95dde5625008ac1860001b617f5f565b8180617ef69061a8cf565b925050617e4d565b617f2a7fb2bdd0bdcdad21e31c7d63fe1ea2e2921dc22759043e9447629a6eacbe24a35860001b617f5f565b617f567f2702eea04cfcc6b80a61083489bd45cce4a987e383a037ba21b9e480ae80c46a60001b617f5f565b81905092915050565b50565b60008151905060006a636f6e736f6c652e6c6f679050602083016000808483855afa5050505050565b50565b50565b50565b600080823b905060008111915050919050565b6040518060600160405280617fba6180bb565b8152602001600015158152602001600081525090565b6040518060e0016040528060608152602001606081526020016060815260200160608152602001606081526020016180066181da565b81526020016000151581525090565b6040518060400160405280600063ffffffff168152602001600063ffffffff1681525090565b6040518060400160405280600067ffffffffffffffff168152602001606081525090565b6040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081525090565b60405180606001604052806180a76180bb565b815260200160608152602001606081525090565b6040518060a00160405280600067ffffffffffffffff168152602001606081526020016060815260200160608152602001606081525090565b8280546181009061a86c565b90600052602060002090601f0160209004810192826181225760008555618169565b82601f1061813b57805160ff1916838001178555618169565b82800160010185558215618169579182015b8281111561816857825182559160200191906001019061814d565b5b5090506181769190618372565b5090565b8280548282559060005260206000209081019282156181c9579160200282015b828111156181c85782518290805190602001906181b89291906180f4565b509160200191906001019061819a565b5b5090506181d6919061838f565b5090565b604051806102e0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff16815260200160008152602001600015158152602001600060018111156182f1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600067ffffffffffffffff16815260200160608152602001606081526020016060815260200160006002811115618356577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200160001515815260200161836c6183b3565b81525090565b5b8082111561838b576000816000905550600101618373565b5090565b5b808211156183af57600081816183a691906183f2565b50600101618390565b5090565b6040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b5080546183fe9061a86c565b6000825580601f10618410575061842f565b601f01602090049060005260206000209081019061842e9190618372565b5b50565b60006184456184408461a0b4565b61a08f565b9050808382526020820190508285602086028201111561846457600080fd5b60005b858110156184ae57813567ffffffffffffffff81111561848657600080fd5b808601618493898261876a565b85526020850194506020840193505050600181019050618467565b5050509392505050565b60006184cb6184c68461a0e0565b61a08f565b905080838252602082019050828560408602820111156184ea57600080fd5b60005b8581101561851a578161850088826187a9565b8452602084019350604083019250506001810190506184ed565b5050509392505050565b60006185376185328461a10c565b61a08f565b9050808382526020820190508285602086028201111561855657600080fd5b60005b858110156185a057813567ffffffffffffffff81111561857857600080fd5b8086016185858982618881565b85526020850194506020840193505050600181019050618559565b5050509392505050565b60006185bd6185b88461a138565b61a08f565b905080838252602082019050828560208602820111156185dc57600080fd5b60005b8581101561862657813567ffffffffffffffff8111156185fe57600080fd5b80860161860b89826188f9565b855260208501945060208401935050506001810190506185df565b5050509392505050565b600061864361863e8461a164565b61a08f565b90508281526020810184848401111561865b57600080fd5b61866684828561a82a565b509392505050565b60008135905061867d8161abbe565b92915050565b600082601f83011261869457600080fd5b81356186a4848260208601618432565b91505092915050565b600082601f8301126186be57600080fd5b81356186ce8482602086016184b8565b91505092915050565b600082601f8301126186e857600080fd5b81356186f8848260208601618524565b91505092915050565b600082601f83011261871257600080fd5b81356187228482602086016185aa565b91505092915050565b60008135905061873a8161abd5565b92915050565b60008135905061874f8161abec565b92915050565b6000815190506187648161abec565b92915050565b600082601f83011261877b57600080fd5b813561878b848260208601618630565b91505092915050565b6000813590506187a38161ac03565b92915050565b6000604082840312156187bb57600080fd5b6187c5604061a08f565b905060006187d584828501618e84565b60008301525060206187e984828501618e84565b60208301525092915050565b60006080828403121561880757600080fd5b618811608061a08f565b905060006188218482850161866e565b600083015250602082013567ffffffffffffffff81111561884157600080fd5b61884d8482850161876a565b602083015250604061886184828501618e99565b604083015250606061887584828501618e99565b60608301525092915050565b60006060828403121561889357600080fd5b61889d606061a08f565b905060006188ad84828501618e99565b60008301525060206188c184828501618e99565b602083015250604082013567ffffffffffffffff8111156188e157600080fd5b6188ed8482850161876a565b60408301525092915050565b60006040828403121561890b57600080fd5b618915604061a08f565b9050600061892584828501618e99565b600083015250602082013567ffffffffffffffff81111561894557600080fd5b618951848285016186d7565b60208301525092915050565b60006060828403121561896f57600080fd5b618979606061a08f565b9050600061898984828501618e99565b600083015250602061899d84828501618e99565b60208301525060406189b184828501618e99565b60408301525092915050565b6000606082840312156189cf57600080fd5b6189d9606061a08f565b9050600082013567ffffffffffffffff8111156189f557600080fd5b618a0184828501618bc5565b600083015250602082013567ffffffffffffffff811115618a2157600080fd5b618a2d848285016186ad565b602083015250604082013567ffffffffffffffff811115618a4d57600080fd5b618a5984828501618cfb565b60408301525092915050565b600060a08284031215618a7757600080fd5b618a8160a061a08f565b90506000618a9184828501618e99565b600083015250602082013567ffffffffffffffff811115618ab157600080fd5b618abd8482850161876a565b602083015250604082013567ffffffffffffffff811115618add57600080fd5b618ae98482850161876a565b604083015250606082013567ffffffffffffffff811115618b0957600080fd5b618b1584828501618683565b606083015250608082013567ffffffffffffffff811115618b3557600080fd5b618b418482850161876a565b60808301525092915050565b600060608284031215618b5f57600080fd5b618b69606061a08f565b9050600082013567ffffffffffffffff811115618b8557600080fd5b618b9184828501618a65565b6000830152506020618ba58482850161872b565b6020830152506040618bb984828501618e6f565b60408301525092915050565b60006101808284031215618bd857600080fd5b618be361018061a08f565b90506000618bf38482850161866e565b6000830152506020618c0784828501618e99565b6020830152506040618c1b84828501618e99565b6040830152506060618c2f84828501618e99565b6060830152506080618c4384828501618794565b60808301525060a0618c5784828501618e6f565b60a08301525060c0618c6b84828501618e6f565b60c08301525060e0618c7f84828501618e99565b60e083015250610100618c9484828501618e99565b61010083015250610120618caa84828501618e99565b61012083015250610140618cc08482850161872b565b6101408301525061016082013567ffffffffffffffff811115618ce257600080fd5b618cee84828501618683565b6101608301525092915050565b600060c08284031215618d0d57600080fd5b618d1760c061a08f565b90506000618d2784828501618e99565b6000830152506020618d3b84828501618e99565b602083015250604082013567ffffffffffffffff811115618d5b57600080fd5b618d678482850161876a565b604083015250606082013567ffffffffffffffff811115618d8757600080fd5b618d9384828501618683565b606083015250608082013567ffffffffffffffff811115618db357600080fd5b618dbf84828501618701565b60808301525060a082013567ffffffffffffffff811115618ddf57600080fd5b618deb8482850161876a565b60a08301525092915050565b600060a08284031215618e0957600080fd5b618e13606061a08f565b90506000618e238482850161895d565b600083015250606082013567ffffffffffffffff811115618e4357600080fd5b618e4f8482850161876a565b6020830152506080618e6384828501618e99565b60408301525092915050565b600081359050618e7e8161ac13565b92915050565b600081359050618e938161ac2a565b92915050565b600081359050618ea88161ac41565b92915050565b600060208284031215618ec057600080fd5b6000618ece84828501618740565b91505092915050565b600060208284031215618ee957600080fd5b6000618ef784828501618755565b91505092915050565b600060208284031215618f1257600080fd5b600082013567ffffffffffffffff811115618f2c57600080fd5b618f388482850161876a565b91505092915050565b60008060408385031215618f5457600080fd5b600083013567ffffffffffffffff811115618f6e57600080fd5b618f7a8582860161876a565b925050602083013567ffffffffffffffff811115618f9757600080fd5b618fa3858286016186ad565b9150509250929050565b60008060408385031215618fc057600080fd5b600083013567ffffffffffffffff811115618fda57600080fd5b618fe68582860161876a565b925050602083013567ffffffffffffffff81111561900357600080fd5b61900f85828601618701565b9150509250929050565b60006040828403121561902b57600080fd5b6000619039848285016187a9565b91505092915050565b60006020828403121561905457600080fd5b600082013567ffffffffffffffff81111561906e57600080fd5b61907a848285016187f5565b91505092915050565b60006020828403121561909557600080fd5b600082013567ffffffffffffffff8111156190af57600080fd5b6190bb84828501618881565b91505092915050565b6000602082840312156190d657600080fd5b600082013567ffffffffffffffff8111156190f057600080fd5b6190fc848285016188f9565b91505092915050565b60006020828403121561911757600080fd5b600082013567ffffffffffffffff81111561913157600080fd5b61913d848285016189bd565b91505092915050565b60006020828403121561915857600080fd5b600082013567ffffffffffffffff81111561917257600080fd5b61917e84828501618a65565b91505092915050565b60008060006060848603121561919c57600080fd5b600084013567ffffffffffffffff8111156191b657600080fd5b6191c286828701618a65565b935050602084013567ffffffffffffffff8111156191df57600080fd5b6191eb868287016186ad565b925050604084013567ffffffffffffffff81111561920857600080fd5b61921486828701618701565b9150509250925092565b60008060006060848603121561923357600080fd5b600084013567ffffffffffffffff81111561924d57600080fd5b61925986828701618b4d565b935050602084013567ffffffffffffffff81111561927657600080fd5b619282868287016186ad565b925050604084013567ffffffffffffffff81111561929f57600080fd5b6192ab86828701618701565b9150509250925092565b6000602082840312156192c757600080fd5b600082013567ffffffffffffffff8111156192e157600080fd5b6192ed84828501618df7565b91505092915050565b60006193028383619376565b60208301905092915050565b600061931a838361975a565b905092915050565b600061932e8383619886565b60408301905092915050565b60006193468383619ac3565b905092915050565b600061935a8383619b13565b905092915050565b600061936e8383619cdc565b905092915050565b61937f8161a70c565b82525050565b61938e8161a70c565b82525050565b6193a56193a08261a70c565b61a945565b82525050565b60006193b68261a1f5565b6193c0818561a29b565b93506193cb8361a195565b8060005b838110156193fc5781516193e388826192f6565b97506193ee8361a24d565b9250506001810190506193cf565b5085935050505092915050565b60006194148261a200565b61941e818561a2ac565b9350836020820285016194308561a1a5565b8060005b8581101561946c578484038952815161944d858261930e565b94506194588361a25a565b925060208a01995050600181019050619434565b50829750879550505050505092915050565b60006194898261a20b565b619493818561a2bd565b935061949e8361a1b5565b8060005b838110156194cf5781516194b68882619322565b97506194c18361a267565b9250506001810190506194a2565b5085935050505092915050565b60006194e78261a20b565b6194f1818561a2ce565b93506194fc8361a1b5565b8060005b8381101561952d5781516195148882619322565b975061951f8361a267565b925050600181019050619500565b5085935050505092915050565b60006195458261a216565b61954f818561a2df565b9350836020820285016195618561a1c5565b8060005b8581101561959d578484038952815161957e858261933a565b94506195898361a274565b925060208a01995050600181019050619565565b50829750879550505050505092915050565b60006195ba8261a221565b6195c4818561a2f0565b9350836020820285016195d68561a1d5565b8060005b8581101561961257848403895281516195f3858261934e565b94506195fe8361a281565b925060208a019950506001810190506195da565b50829750879550505050505092915050565b600061962f8261a221565b619639818561a301565b93508360208202850161964b8561a1d5565b8060005b858110156196875784840389528151619668858261934e565b94506196738361a281565b925060208a0199505060018101905061964f565b50829750879550505050505092915050565b60006196a48261a22c565b6196ae818561a312565b9350836020820285016196c08561a1e5565b8060005b858110156196fc57848403895281516196dd8582619362565b94506196e88361a28e565b925060208a019950506001810190506196c4565b50829750879550505050505092915050565b6197178161a71e565b82525050565b6197268161a71e565b82525050565b61973d6197388261a72a565b61a957565b82525050565b61975461974f8261a756565b61a961565b82525050565b60006197658261a237565b61976f818561a323565b935061977f81856020860161a839565b6197888161aaee565b840191505092915050565b600061979e8261a237565b6197a8818561a334565b93506197b881856020860161a839565b6197c18161aaee565b840191505092915050565b60006197d78261a237565b6197e1818561a345565b93506197f181856020860161a839565b80840191505092915050565b6198068161a7f4565b82525050565b6198158161a806565b82525050565b6198248161a818565b82525050565b60006198358261a242565b61983f818561a350565b935061984f81856020860161a839565b6198588161aaee565b840191505092915050565b6000619870602e8361a350565b915061987b8261ab33565b604082019050919050565b60408201600082015161989c6000850182619d58565b5060208201516198af6020850182619d58565b50505050565b60006103208301600083015184820360008601526198d3828261975a565b91505060208301516198e86020860182619376565b5060408301518482036040860152619900828261975a565b91505060608301516199156060860182619d7e565b5060808301516199286080860182619d7e565b5060a083015161993b60a0860182619d7e565b5060c083015161994e60c0860182619d7e565b5060e083015161996160e0860182619d7e565b50610100830151619976610100860182619d3a565b5061012083015161998b610120860182619d7e565b506101408301516199a0610140860182619d7e565b506101608301518482036101608601526199ba828261975a565b9150506101808301516199d1610180860182619d7e565b506101a08301516199e66101a0860182619d3a565b506101c08301516199fb6101c086018261970e565b506101e0830151619a106101e086018261981b565b50610200830151619a25610200860182619d7e565b50610220830151848203610220860152619a3f82826193ab565b915050610240830151848203610240860152619a5b82826193ab565b915050610260830151848203610260860152619a77828261975a565b915050610280830151619a8e61028086018261980c565b506102a0830151619aa36102a086018261970e565b506102c0830151619ab86102c0860182619c0f565b508091505092915050565b6000606083016000830151619adb6000860182619d7e565b506020830151619aee6020860182619d7e565b5060408301518482036040860152619b06828261975a565b9150508091505092915050565b6000604083016000830151619b2b6000860182619d7e565b5060208301518482036020860152619b43828261953a565b9150508091505092915050565b600060e0830160008301518482036000860152619b6d828261975a565b91505060208301518482036020860152619b878282619409565b91505060408301518482036040860152619ba1828261947e565b91505060608301518482036060860152619bbb82826195af565b91505060808301518482036080860152619bd5828261975a565b91505060a083015184820360a0860152619bef82826198b5565b91505060c0830151619c0460c086018261970e565b508091505092915050565b606082016000820151619c256000850182619d7e565b506020820151619c386020850182619d7e565b506040820151619c4b6040850182619d7e565b50505050565b600060a083016000830151619c696000860182619d7e565b5060208301518482036020860152619c81828261975a565b91505060408301518482036040860152619c9b828261975a565b91505060608301518482036060860152619cb58282619409565b91505060808301518482036080860152619ccf828261975a565b9150508091505092915050565b60006060830160008301518482036000860152619cf98282619c51565b91505060208301518482036020860152619d13828261947e565b91505060408301518482036040860152619d2d82826195af565b9150508091505092915050565b619d438161a7b9565b82525050565b619d528161a7b9565b82525050565b619d618161a7c3565b82525050565b619d78619d738261a7c3565b61a97d565b82525050565b619d878161a7d3565b82525050565b619d9e619d998261a7d3565b61a98f565b82525050565b6000619db08285619394565b601482019150619dc082846197cc565b91508190509392505050565b6000619dd88287619394565b601482019150619de882866197cc565b9150619df48285619d8d565b600882019150619e048284619d8d565b60088201915081905095945050505050565b6000619e228284619743565b60208201915081905092915050565b6000619e3d82846197cc565b915081905092915050565b6000619e5482856197cc565b9150619e60828461972c565b6001820191508190509392505050565b6000619e7c82856197cc565b9150619e8882846197cc565b91508190509392505050565b6000619ea082856197cc565b9150619eac8284619d8d565b6008820191508190509392505050565b6000619ec88285619d67565b600482019150619ed88284619d67565b6004820191508190509392505050565b6000619ef48288619d8d565b600882019150619f0482876197cc565b9150619f1082866197cc565b9150619f1c82856197cc565b9150619f2882846197cc565b91508190509695505050505050565b6000602082019050619f4c6000830184619385565b92915050565b60006020820190508181036000830152619f6c81846194dc565b905092915050565b60006020820190508181036000830152619f8e8184619624565b905092915050565b60006020820190508181036000830152619fb08184619699565b905092915050565b6000602082019050619fcd600083018461971d565b92915050565b60006020820190508181036000830152619fed8184619793565b905092915050565b600060208201905061a00a60008301846197fd565b92915050565b6000602082019050818103600083015261a02a818461982a565b905092915050565b6000602082019050818103600083015261a04b81619863565b9050919050565b6000602082019050818103600083015261a06c8184619b50565b905092915050565b600060208201905061a0896000830184619d49565b92915050565b600061a09961a0aa565b905061a0a5828261a89e565b919050565b6000604051905090565b600067ffffffffffffffff82111561a0cf5761a0ce61aabf565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561a0fb5761a0fa61aabf565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561a1275761a12661aabf565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561a1535761a15261aabf565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561a17f5761a17e61aabf565b5b61a1888261aaee565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600061a36c8261a7b9565b915061a3778361a7b9565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561a3ac5761a3ab61aa03565b5b828201905092915050565b600061a3c28261a7c3565b915061a3cd8361a7c3565b92508263ffffffff0382111561a3e65761a3e561aa03565b5b828201905092915050565b600061a3fc8261a7d3565b915061a4078361a7d3565b92508267ffffffffffffffff0382111561a4245761a42361aa03565b5b828201905092915050565b600061a43a8261a7e7565b915061a4458361a7e7565b92508260ff0382111561a45b5761a45a61aa03565b5b828201905092915050565b600061a4718261a7d3565b915061a47c8361a7d3565b92508261a48c5761a48b61aa32565b5b828204905092915050565b6000808291508390505b600185111561a4e15780860481111561a4bd5761a4bc61aa03565b5b600185161561a4cc5780820291505b808102905061a4da8561ab26565b945061a4a1565b94509492505050565b600061a4f58261a7b9565b915061a5008361a7b9565b925061a52d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848461a535565b905092915050565b60008261a545576001905061a601565b8161a553576000905061a601565b816001811461a569576002811461a5735761a5a2565b600191505061a601565b60ff84111561a5855761a58461aa03565b5b8360020a91508482111561a59c5761a59b61aa03565b5b5061a601565b5060208310610133831016604e8410600b841016171561a5d75782820a90508381111561a5d25761a5d161aa03565b5b61a601565b61a5e4848484600161a497565b9250905081840481111561a5fb5761a5fa61aa03565b5b81810290505b9392505050565b600061a6138261a7b9565b915061a61e8361a7b9565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561a6575761a65661aa03565b5b828202905092915050565b600061a66d8261a7d3565b915061a6788361a7d3565b92508167ffffffffffffffff048311821515161561a6995761a69861aa03565b5b828202905092915050565b600061a6af8261a7b9565b915061a6ba8361a7b9565b92508282101561a6cd5761a6cc61aa03565b5b828203905092915050565b600061a6e38261a7c3565b915061a6ee8361a7c3565b92508282101561a7015761a70061aa03565b5b828203905092915050565b600061a7178261a799565b9050919050565b60008115159050919050565b60007fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b6000819050919050565b600081905061a76e8261ab82565b919050565b600081905061a7818261ab96565b919050565b600081905061a7948261abaa565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b600061a7ff8261a760565b9050919050565b600061a8118261a773565b9050919050565b600061a8238261a786565b9050919050565b82818337600083830152505050565b60005b8381101561a85757808201518184015260208101905061a83c565b8381111561a866576000848401525b50505050565b6000600282049050600182168061a88457607f821691505b6020821081141561a8985761a89761aa90565b5b50919050565b61a8a78261aaee565b810181811067ffffffffffffffff8211171561a8c65761a8c561aabf565b5b80604052505050565b600061a8da8261a7b9565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561a90d5761a90c61aa03565b5b600182019050919050565b600061a9238261a7c3565b915063ffffffff82141561a93a5761a93961aa03565b5b600182019050919050565b600061a9508261a96b565b9050919050565b6000819050919050565b6000819050919050565b600061a9768261ab19565b9050919050565b600061a9888261ab0c565b9050919050565b600061a99a8261aaff565b9050919050565b600061a9ac8261a7c3565b915061a9b78361a7c3565b92508261a9c75761a9c661aa32565b5b828206905092915050565b600061a9dd8261a7d3565b915061a9e88361a7d3565b92508261a9f85761a9f761aa32565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160c01b9050919050565b60008160e01b9050919050565b60008160601b9050919050565b60008160011c9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b600b811061ab935761ab9261aa61565b5b50565b6003811061aba75761aba661aa61565b5b50565b6002811061abbb5761abba61aa61565b5b50565b61abc78161a70c565b811461abd257600080fd5b50565b61abde8161a71e565b811461abe957600080fd5b50565b61abf58161a756565b811461ac0057600080fd5b50565b6003811061ac1057600080fd5b50565b61ac1c8161a7b9565b811461ac2757600080fd5b50565b61ac338161a7c3565b811461ac3e57600080fd5b50565b61ac4a8161a7d3565b811461ac5557600080fd5b5056fea26469706673582212202bf75132de76d76c44c479575518b00a52f5bc1519c22db83f0806c776f1b80d64736f6c63430008040033",
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

// GetKeyByProofParams is a free data retrieval call binding the contract method 0x958e549c.
//
// Solidity: function GetKeyByProofParams((uint64,bytes,bytes,bytes[],bytes) vParams) pure returns(bytes)
func (_Store *StoreCaller) GetKeyByProofParams(opts *bind.CallOpts, vParams ProofParams) ([]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetKeyByProofParams", vParams)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetKeyByProofParams is a free data retrieval call binding the contract method 0x958e549c.
//
// Solidity: function GetKeyByProofParams((uint64,bytes,bytes,bytes[],bytes) vParams) pure returns(bytes)
func (_Store *StoreSession) GetKeyByProofParams(vParams ProofParams) ([]byte, error) {
	return _Store.Contract.GetKeyByProofParams(&_Store.CallOpts, vParams)
}

// GetKeyByProofParams is a free data retrieval call binding the contract method 0x958e549c.
//
// Solidity: function GetKeyByProofParams((uint64,bytes,bytes,bytes[],bytes) vParams) pure returns(bytes)
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
// Solidity: function GetUnVerifyProofList() view returns(((uint64,bytes,bytes,bytes[],bytes),(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[])[])
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
// Solidity: function GetUnVerifyProofList() view returns(((uint64,bytes,bytes,bytes[],bytes),(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[])[])
func (_Store *StoreSession) GetUnVerifyProofList() ([]ProofRecordWithParams, error) {
	return _Store.Contract.GetUnVerifyProofList(&_Store.CallOpts)
}

// GetUnVerifyProofList is a free data retrieval call binding the contract method 0xf5dbc78e.
//
// Solidity: function GetUnVerifyProofList() view returns(((uint64,bytes,bytes,bytes[],bytes),(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[])[])
func (_Store *StoreCallerSession) GetUnVerifyProofList() ([]ProofRecordWithParams, error) {
	return _Store.Contract.GetUnVerifyProofList(&_Store.CallOpts)
}

// PrepareForPdpVerification is a free data retrieval call binding the contract method 0x57d60d6e.
//
// Solidity: function PrepareForPdpVerification(((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]),(uint32,uint32)[],(uint64,uint64,bytes,bytes[],(uint64,(uint64,uint64,bytes)[])[],bytes)) pParams) view returns((bytes,bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)),bool))
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
// Solidity: function PrepareForPdpVerification(((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]),(uint32,uint32)[],(uint64,uint64,bytes,bytes[],(uint64,(uint64,uint64,bytes)[])[],bytes)) pParams) view returns((bytes,bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)),bool))
func (_Store *StoreSession) PrepareForPdpVerification(pParams PrepareForPdpVerificationParams) (PdpVerificationReturns, error) {
	return _Store.Contract.PrepareForPdpVerification(&_Store.CallOpts, pParams)
}

// PrepareForPdpVerification is a free data retrieval call binding the contract method 0x57d60d6e.
//
// Solidity: function PrepareForPdpVerification(((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]),(uint32,uint32)[],(uint64,uint64,bytes,bytes[],(uint64,(uint64,uint64,bytes)[])[],bytes)) pParams) view returns((bytes,bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)),bool))
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

// VerifyProofWithMerklePathForFile is a free data retrieval call binding the contract method 0x507af805.
//
// Solidity: function VerifyProofWithMerklePathForFile((uint64,bytes,bytes,bytes[],bytes) vParams) view returns(bool)
func (_Store *StoreCaller) VerifyProofWithMerklePathForFile(opts *bind.CallOpts, vParams ProofParams) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "VerifyProofWithMerklePathForFile", vParams)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProofWithMerklePathForFile is a free data retrieval call binding the contract method 0x507af805.
//
// Solidity: function VerifyProofWithMerklePathForFile((uint64,bytes,bytes,bytes[],bytes) vParams) view returns(bool)
func (_Store *StoreSession) VerifyProofWithMerklePathForFile(vParams ProofParams) (bool, error) {
	return _Store.Contract.VerifyProofWithMerklePathForFile(&_Store.CallOpts, vParams)
}

// VerifyProofWithMerklePathForFile is a free data retrieval call binding the contract method 0x507af805.
//
// Solidity: function VerifyProofWithMerklePathForFile((uint64,bytes,bytes,bytes[],bytes) vParams) view returns(bool)
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

// C0x7b63f42e is a free data retrieval call binding the contract method 0xcca9fb6e.
//
// Solidity: function c_0x7b63f42e(bytes32 c__0x7b63f42e) pure returns()
func (_Store *StoreCaller) C0x7b63f42e(opts *bind.CallOpts, c__0x7b63f42e [32]byte) error {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "c_0x7b63f42e", c__0x7b63f42e)

	if err != nil {
		return err
	}

	return err

}

// C0x7b63f42e is a free data retrieval call binding the contract method 0xcca9fb6e.
//
// Solidity: function c_0x7b63f42e(bytes32 c__0x7b63f42e) pure returns()
func (_Store *StoreSession) C0x7b63f42e(c__0x7b63f42e [32]byte) error {
	return _Store.Contract.C0x7b63f42e(&_Store.CallOpts, c__0x7b63f42e)
}

// C0x7b63f42e is a free data retrieval call binding the contract method 0xcca9fb6e.
//
// Solidity: function c_0x7b63f42e(bytes32 c__0x7b63f42e) pure returns()
func (_Store *StoreCallerSession) C0x7b63f42e(c__0x7b63f42e [32]byte) error {
	return _Store.Contract.C0x7b63f42e(&_Store.CallOpts, c__0x7b63f42e)
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

// SubmitVerifyProofRequest is a paid mutator transaction binding the contract method 0xe0bdfb40.
//
// Solidity: function SubmitVerifyProofRequest((uint64,bytes,bytes,bytes[],bytes) vParams, (uint32,uint32)[] chgs, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreTransactor) SubmitVerifyProofRequest(opts *bind.TransactOpts, vParams ProofParams, chgs []Challenge, mp []MerklePath) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "SubmitVerifyProofRequest", vParams, chgs, mp)
}

// SubmitVerifyProofRequest is a paid mutator transaction binding the contract method 0xe0bdfb40.
//
// Solidity: function SubmitVerifyProofRequest((uint64,bytes,bytes,bytes[],bytes) vParams, (uint32,uint32)[] chgs, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreSession) SubmitVerifyProofRequest(vParams ProofParams, chgs []Challenge, mp []MerklePath) (*types.Transaction, error) {
	return _Store.Contract.SubmitVerifyProofRequest(&_Store.TransactOpts, vParams, chgs, mp)
}

// SubmitVerifyProofRequest is a paid mutator transaction binding the contract method 0xe0bdfb40.
//
// Solidity: function SubmitVerifyProofRequest((uint64,bytes,bytes,bytes[],bytes) vParams, (uint32,uint32)[] chgs, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreTransactorSession) SubmitVerifyProofRequest(vParams ProofParams, chgs []Challenge, mp []MerklePath) (*types.Transaction, error) {
	return _Store.Contract.SubmitVerifyProofRequest(&_Store.TransactOpts, vParams, chgs, mp)
}

// VerifyProof is a paid mutator transaction binding the contract method 0x173dee68.
//
// Solidity: function VerifyProof(((uint64,bytes,bytes,bytes[],bytes),bool,uint256) vParams, (uint32,uint32)[] chgs, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreTransactor) VerifyProof(opts *bind.TransactOpts, vParams ProofRecord, chgs []Challenge, mp []MerklePath) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "VerifyProof", vParams, chgs, mp)
}

// VerifyProof is a paid mutator transaction binding the contract method 0x173dee68.
//
// Solidity: function VerifyProof(((uint64,bytes,bytes,bytes[],bytes),bool,uint256) vParams, (uint32,uint32)[] chgs, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
func (_Store *StoreSession) VerifyProof(vParams ProofRecord, chgs []Challenge, mp []MerklePath) (*types.Transaction, error) {
	return _Store.Contract.VerifyProof(&_Store.TransactOpts, vParams, chgs, mp)
}

// VerifyProof is a paid mutator transaction binding the contract method 0x173dee68.
//
// Solidity: function VerifyProof(((uint64,bytes,bytes,bytes[],bytes),bool,uint256) vParams, (uint32,uint32)[] chgs, (uint64,(uint64,uint64,bytes)[])[] mp) payable returns()
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
