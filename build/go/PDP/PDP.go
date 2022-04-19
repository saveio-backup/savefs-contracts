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

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"HashValue\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveNum\",\"type\":\"uint64\"}],\"internalType\":\"structGenChallengeParams\",\"name\":\"gParams\",\"type\":\"tuple\"}],\"name\":\"GenChallenge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"SectorInfo_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"Challenges\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"ProveFileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"BlockNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Tags\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"MerklePath_\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"PlotData\",\"type\":\"bytes\"}],\"internalType\":\"structSectorProveData\",\"name\":\"ProveData\",\"type\":\"tuple\"}],\"internalType\":\"structPrepareForPdpVerificationParams\",\"name\":\"pParams\",\"type\":\"tuple\"}],\"name\":\"PrepareForPdpVerification\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"FileIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"UpdatedChal\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"Path\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"FileInfo_\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"Success\",\"type\":\"bool\"}],\"internalType\":\"structPdpVerificationReturns\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"PlotData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"}],\"internalType\":\"structVerifyPlotDataParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"VerifyPlotData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"FileIds\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"Challenges\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"MerklePath_\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"}],\"internalType\":\"structVerifyProofWithMerklePathForFileParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"VerifyProofWithMerklePathForFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0x7b63f42e\",\"type\":\"bytes32\"}],\"name\":\"c_0x7b63f42e\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506120d5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c8063b6ddeca011610050578063b6ddeca0146100e1578063cca9fb6e14610111578063fd7c98081461012d57610072565b80632e19aeff146100775780638129fc1c146100a75780639f9ca644146100b1575b600080fd5b610091600480360381019061008c9190611306565b61015d565b60405161009e9190611b2e565b60405180910390f35b6100af6101ec565b005b6100cb60048036038101906100c691906112c5565b6102fc565b6040516100d89190611b69565b60405180910390f35b6100fb60048036038101906100f69190611347565b61045a565b6040516101089190611b2e565b60405180910390f35b61012b6004803603810190610126919061125b565b6104e8565b005b61014760048036038101906101429190611284565b6104eb565b6040516101549190611b0c565b60405180910390f35b600061018b7f5efc60bee29c0fab48f02d348a6fb1488c1e4a747ff4674e48884cd07d47d82b60001b6104e8565b6101b77f7b75ebb277e9f8a00296301e0334183f97c247ad425fd158c063e923c698c2d360001b6104e8565b6101e37f60e37000abcfd880e47df587a1a2cc4409b97fc516e395e6137913b9ece4d20e60001b6104e8565b60019050919050565b600060019054906101000a900460ff166102145760008054906101000a900460ff161561021d565b61021c6105d4565b5b61025c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025390611b49565b60405180910390fd5b60008060019054906101000a900460ff1615905080156102ac576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b6102d87f44db0aaf6904b3315a40e84f3442f84be6501dc9d7e936acea091c7eda167a7f60001b6104e8565b80156102f95760008060016101000a81548160ff0219169083151502179055505b50565b6103046105f8565b6103307fce680254565d432b46831a931b297d93b45a98bbe4593a2c47176a5d09f4d1e560001b6104e8565b61035c7fee9e120134b16802739bfd54717171eda884099791851f31e142c7cae1fabbcd60001b6104e8565b6103887f3689a3d102cb26631f84265f0e76baca5bb050d923401ec3d2534a76bdefadce60001b6104e8565b6103906105f8565b6103bc7f9abfbb5ca58067695023328cd39c06b369e6e051346a6a46002d12f913f5d1fd60001b6104e8565b6103e87ffac7c10c79bc52656f5a435013d57d2e8df6a210f5cbf50d59d8e02704674d5460001b6104e8565b60018160c00190151590811515815250506104257f723c39863d15fd2055cd47bfca0d40683f541d9e01ebd97e7b90ade8695f6f0060001b6104e8565b6104517f804d271d8c17ea46edba27abc7994828bb4ffffc50fb0538838aacc4fd07477e60001b6104e8565b80915050919050565b60006104887fe47f26a3071c363b5827c764e86b7fc35471727f1bfdc858ee28115d3277477860001b6104e8565b6104b37ec89a30f62f49d1ef51eac6a5216936c9f53411f7e15eab33134c8909e78a6260001b6104e8565b6104df7ff1d6e8c892a9832748f724d5568d0734b972497e0c8fa5d02ebccac4cf9b9b9260001b6104e8565b60019050919050565b50565b60606105197ff1bb8c03959f8a1ce1a18da3f3d470a8291f92fa9dc941e8544491cfdb39815f60001b6104e8565b6105457fd82d45ee036f12215a680de7226f350b0f83f34c22d0007b7859ef729ea8607060001b6104e8565b6105717fe1f1d7d1b05ed5cd0562db2ec49b3674f4caf6034fd0033e689334ab9b864d6860001b6104e8565b606061059f7fa29cb9714ea6158b19ac002f5f8246093198aae56ec2aa2166d3ce33f7b6c90c60001b6104e8565b6105cb7f5918c56eb3df23b1a0c7c04b5deee0b89afb33cc797aaf2a947b416678dc51b860001b6104e8565b80915050919050565b60006105df306105e5565b15905090565b600080823b905060008111915050919050565b6040518060e00160405280606081526020016060815260200160608152602001606081526020016060815260200161062e61063d565b81526020016000151581525090565b604051806102e0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff1681526020016000815260200160001515815260200160006001811115610754577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600067ffffffffffffffff168152602001606081526020016060815260200160608152602001600060028111156107b9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016000151581526020016107cf6107d5565b81525090565b6040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b600061082761082284611bb0565b611b8b565b9050808382526020820190508285602086028201111561084657600080fd5b60005b8581101561089057813567ffffffffffffffff81111561086857600080fd5b8086016108758982610b37565b85526020850194506020840193505050600181019050610849565b5050509392505050565b60006108ad6108a884611bdc565b611b8b565b905080838252602082019050828560408602820111156108cc57600080fd5b60005b858110156108fc57816108e28882610b76565b8452602084019350604083019250506001810190506108cf565b5050509392505050565b600061091961091484611c08565b611b8b565b9050808382526020820190508285602086028201111561093857600080fd5b60005b8581101561098257813567ffffffffffffffff81111561095a57600080fd5b8086016109678982610c4e565b8552602085019450602084019350505060018101905061093b565b5050509392505050565b600061099f61099a84611c34565b611b8b565b905080838252602082019050828560208602820111156109be57600080fd5b60005b85811015610a0857813567ffffffffffffffff8111156109e057600080fd5b8086016109ed8982610cc6565b855260208501945060208401935050506001810190506109c1565b5050509392505050565b6000610a25610a2084611c60565b611b8b565b905082815260208101848484011115610a3d57600080fd5b610a48848285611eac565b509392505050565b600081359050610a5f81612005565b92915050565b600082601f830112610a7657600080fd5b8135610a86848260208601610814565b91505092915050565b600082601f830112610aa057600080fd5b8135610ab084826020860161089a565b91505092915050565b600082601f830112610aca57600080fd5b8135610ada848260208601610906565b91505092915050565b600082601f830112610af457600080fd5b8135610b0484826020860161098c565b91505092915050565b600081359050610b1c8161201c565b92915050565b600081359050610b3181612033565b92915050565b600082601f830112610b4857600080fd5b8135610b58848260208601610a12565b91505092915050565b600081359050610b708161204a565b92915050565b600060408284031215610b8857600080fd5b610b926040611b8b565b90506000610ba284828501611231565b6000830152506020610bb684828501611231565b60208301525092915050565b600060808284031215610bd457600080fd5b610bde6080611b8b565b90506000610bee84828501610a50565b600083015250602082013567ffffffffffffffff811115610c0e57600080fd5b610c1a84828501610b37565b6020830152506040610c2e84828501611246565b6040830152506060610c4284828501611246565b60608301525092915050565b600060608284031215610c6057600080fd5b610c6a6060611b8b565b90506000610c7a84828501611246565b6000830152506020610c8e84828501611246565b602083015250604082013567ffffffffffffffff811115610cae57600080fd5b610cba84828501610b37565b60408301525092915050565b600060408284031215610cd857600080fd5b610ce26040611b8b565b90506000610cf284828501611246565b600083015250602082013567ffffffffffffffff811115610d1257600080fd5b610d1e84828501610ab9565b60208301525092915050565b600060608284031215610d3c57600080fd5b610d466060611b8b565b90506000610d5684828501611246565b6000830152506020610d6a84828501611246565b6020830152506040610d7e84828501611246565b60408301525092915050565b600060608284031215610d9c57600080fd5b610da66060611b8b565b9050600082013567ffffffffffffffff811115610dc257600080fd5b610dce84828501610e32565b600083015250602082013567ffffffffffffffff811115610dee57600080fd5b610dfa84828501610a8f565b602083015250604082013567ffffffffffffffff811115610e1a57600080fd5b610e2684828501610f68565b60408301525092915050565b60006101808284031215610e4557600080fd5b610e50610180611b8b565b90506000610e6084828501610a50565b6000830152506020610e7484828501611246565b6020830152506040610e8884828501611246565b6040830152506060610e9c84828501611246565b6060830152506080610eb084828501610b61565b60808301525060a0610ec48482850161121c565b60a08301525060c0610ed88482850161121c565b60c08301525060e0610eec84828501611246565b60e083015250610100610f0184828501611246565b61010083015250610120610f1784828501611246565b61012083015250610140610f2d84828501610b0d565b6101408301525061016082013567ffffffffffffffff811115610f4f57600080fd5b610f5b84828501610a65565b6101608301525092915050565b600060c08284031215610f7a57600080fd5b610f8460c0611b8b565b90506000610f9484828501611246565b6000830152506020610fa884828501611246565b602083015250604082013567ffffffffffffffff811115610fc857600080fd5b610fd484828501610b37565b604083015250606082013567ffffffffffffffff811115610ff457600080fd5b61100084828501610b37565b606083015250608082013567ffffffffffffffff81111561102057600080fd5b61102c84828501610ae3565b60808301525060a082013567ffffffffffffffff81111561104c57600080fd5b61105884828501610b37565b60a08301525092915050565b600060a0828403121561107657600080fd5b6110806060611b8b565b9050600061109084828501610d2a565b600083015250606082013567ffffffffffffffff8111156110b057600080fd5b6110bc84828501610b37565b60208301525060806110d084828501611246565b60408301525092915050565b600060e082840312156110ee57600080fd5b6110f860e0611b8b565b9050600061110884828501611246565b600083015250602082013567ffffffffffffffff81111561112857600080fd5b61113484828501610b37565b602083015250604082013567ffffffffffffffff81111561115457600080fd5b61116084828501610a65565b604083015250606082013567ffffffffffffffff81111561118057600080fd5b61118c84828501610a65565b606083015250608082013567ffffffffffffffff8111156111ac57600080fd5b6111b884828501610a8f565b60808301525060a082013567ffffffffffffffff8111156111d857600080fd5b6111e484828501610ae3565b60a08301525060c082013567ffffffffffffffff81111561120457600080fd5b61121084828501610b37565b60c08301525092915050565b60008135905061122b8161205a565b92915050565b60008135905061124081612071565b92915050565b60008135905061125581612088565b92915050565b60006020828403121561126d57600080fd5b600061127b84828501610b22565b91505092915050565b60006020828403121561129657600080fd5b600082013567ffffffffffffffff8111156112b057600080fd5b6112bc84828501610bc2565b91505092915050565b6000602082840312156112d757600080fd5b600082013567ffffffffffffffff8111156112f157600080fd5b6112fd84828501610d8a565b91505092915050565b60006020828403121561131857600080fd5b600082013567ffffffffffffffff81111561133257600080fd5b61133e84828501611064565b91505092915050565b60006020828403121561135957600080fd5b600082013567ffffffffffffffff81111561137357600080fd5b61137f848285016110dc565b91505092915050565b600061139483836113f4565b60208301905092915050565b60006113ac838361169a565b905092915050565b60006113c08383611714565b60408301905092915050565b60006113d88383611951565b905092915050565b60006113ec83836119a1565b905092915050565b6113fd81611dec565b82525050565b600061140e82611ce1565b6114188185611d64565b935061142383611c91565b8060005b8381101561145457815161143b8882611388565b975061144683611d23565b925050600181019050611427565b5085935050505092915050565b600061146c82611cec565b6114768185611d75565b93508360208202850161148885611ca1565b8060005b858110156114c457848403895281516114a585826113a0565b94506114b083611d30565b925060208a0199505060018101905061148c565b50829750879550505050505092915050565b60006114e182611cf7565b6114eb8185611d86565b93506114f683611cb1565b8060005b8381101561152757815161150e88826113b4565b975061151983611d3d565b9250506001810190506114fa565b5085935050505092915050565b600061153f82611cf7565b6115498185611d97565b935061155483611cb1565b8060005b8381101561158557815161156c88826113b4565b975061157783611d3d565b925050600181019050611558565b5085935050505092915050565b600061159d82611d02565b6115a78185611da8565b9350836020820285016115b985611cc1565b8060005b858110156115f557848403895281516115d685826113cc565b94506115e183611d4a565b925060208a019950506001810190506115bd565b50829750879550505050505092915050565b600061161282611d0d565b61161c8185611db9565b93508360208202850161162e85611cd1565b8060005b8581101561166a578484038952815161164b85826113e0565b945061165683611d57565b925060208a01995050600181019050611632565b50829750879550505050505092915050565b61168581611dfe565b82525050565b61169481611dfe565b82525050565b60006116a582611d18565b6116af8185611dca565b93506116bf818560208601611ebb565b6116c881611f7d565b840191505092915050565b6116dc81611e88565b82525050565b6116eb81611e9a565b82525050565b60006116fe602e83611ddb565b915061170982611f8e565b604082019050919050565b60408201600082015161172a6000850182611aee565b50602082015161173d6020850182611aee565b50505050565b6000610320830160008301518482036000860152611761828261169a565b915050602083015161177660208601826113f4565b506040830151848203604086015261178e828261169a565b91505060608301516117a36060860182611afd565b5060808301516117b66080860182611afd565b5060a08301516117c960a0860182611afd565b5060c08301516117dc60c0860182611afd565b5060e08301516117ef60e0860182611afd565b50610100830151611804610100860182611adf565b50610120830151611819610120860182611afd565b5061014083015161182e610140860182611afd565b50610160830151848203610160860152611848828261169a565b91505061018083015161185f610180860182611afd565b506101a08301516118746101a0860182611adf565b506101c08301516118896101c086018261167c565b506101e083015161189e6101e08601826116e2565b506102008301516118b3610200860182611afd565b506102208301518482036102208601526118cd8282611403565b9150506102408301518482036102408601526118e98282611403565b915050610260830151848203610260860152611905828261169a565b91505061028083015161191c6102808601826116d3565b506102a08301516119316102a086018261167c565b506102c08301516119466102c0860182611a9d565b508091505092915050565b60006060830160008301516119696000860182611afd565b50602083015161197c6020860182611afd565b5060408301518482036040860152611994828261169a565b9150508091505092915050565b60006040830160008301516119b96000860182611afd565b50602083015184820360208601526119d18282611592565b9150508091505092915050565b600060e08301600083015184820360008601526119fb8282611461565b91505060208301518482036020860152611a158282611461565b91505060408301518482036040860152611a2f82826114d6565b91505060608301518482036060860152611a498282611607565b91505060808301518482036080860152611a63828261169a565b91505060a083015184820360a0860152611a7d8282611743565b91505060c0830151611a9260c086018261167c565b508091505092915050565b606082016000820151611ab36000850182611afd565b506020820151611ac66020850182611afd565b506040820151611ad96040850182611afd565b50505050565b611ae881611e5a565b82525050565b611af781611e64565b82525050565b611b0681611e74565b82525050565b60006020820190508181036000830152611b268184611534565b905092915050565b6000602082019050611b43600083018461168b565b92915050565b60006020820190508181036000830152611b62816116f1565b9050919050565b60006020820190508181036000830152611b8381846119de565b905092915050565b6000611b95611ba6565b9050611ba18282611eee565b919050565b6000604051905090565b600067ffffffffffffffff821115611bcb57611bca611f4e565b5b602082029050602081019050919050565b600067ffffffffffffffff821115611bf757611bf6611f4e565b5b602082029050602081019050919050565b600067ffffffffffffffff821115611c2357611c22611f4e565b5b602082029050602081019050919050565b600067ffffffffffffffff821115611c4f57611c4e611f4e565b5b602082029050602081019050919050565b600067ffffffffffffffff821115611c7b57611c7a611f4e565b5b611c8482611f7d565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611df782611e3a565b9050919050565b60008115159050919050565b6000819050919050565b6000819050611e2282611fdd565b919050565b6000819050611e3582611ff1565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b600067ffffffffffffffff82169050919050565b6000611e9382611e14565b9050919050565b6000611ea582611e27565b9050919050565b82818337600083830152505050565b60005b83811015611ed9578082015181840152602081019050611ebe565b83811115611ee8576000848401525b50505050565b611ef782611f7d565b810181811067ffffffffffffffff82111715611f1657611f15611f4e565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b60038110611fee57611fed611f1f565b5b50565b6002811061200257612001611f1f565b5b50565b61200e81611dec565b811461201957600080fd5b50565b61202581611dfe565b811461203057600080fd5b50565b61203c81611e0a565b811461204757600080fd5b50565b6003811061205757600080fd5b50565b61206381611e5a565b811461206e57600080fd5b50565b61207a81611e64565b811461208557600080fd5b50565b61209181611e74565b811461209c57600080fd5b5056fea2646970667358221220ebb7a12dd6dead766ed8b90536ab23e9a95ea4cce2e33d1e5a762bd481844e3664736f6c63430008040033",
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

// GenChallenge is a free data retrieval call binding the contract method 0xfd7c9808.
//
// Solidity: function GenChallenge((address,bytes,uint64,uint64) gParams) pure returns((uint32,uint32)[])
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
// Solidity: function GenChallenge((address,bytes,uint64,uint64) gParams) pure returns((uint32,uint32)[])
func (_Store *StoreSession) GenChallenge(gParams GenChallengeParams) ([]Challenge, error) {
	return _Store.Contract.GenChallenge(&_Store.CallOpts, gParams)
}

// GenChallenge is a free data retrieval call binding the contract method 0xfd7c9808.
//
// Solidity: function GenChallenge((address,bytes,uint64,uint64) gParams) pure returns((uint32,uint32)[])
func (_Store *StoreCallerSession) GenChallenge(gParams GenChallengeParams) ([]Challenge, error) {
	return _Store.Contract.GenChallenge(&_Store.CallOpts, gParams)
}

// PrepareForPdpVerification is a free data retrieval call binding the contract method 0x9f9ca644.
//
// Solidity: function PrepareForPdpVerification(((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]),(uint32,uint32)[],(uint64,uint64,bytes,bytes,(uint64,(uint64,uint64,bytes)[])[],bytes)) pParams) pure returns((bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)),bool))
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
// Solidity: function PrepareForPdpVerification(((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]),(uint32,uint32)[],(uint64,uint64,bytes,bytes,(uint64,(uint64,uint64,bytes)[])[],bytes)) pParams) pure returns((bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)),bool))
func (_Store *StoreSession) PrepareForPdpVerification(pParams PrepareForPdpVerificationParams) (PdpVerificationReturns, error) {
	return _Store.Contract.PrepareForPdpVerification(&_Store.CallOpts, pParams)
}

// PrepareForPdpVerification is a free data retrieval call binding the contract method 0x9f9ca644.
//
// Solidity: function PrepareForPdpVerification(((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]),(uint32,uint32)[],(uint64,uint64,bytes,bytes,(uint64,(uint64,uint64,bytes)[])[],bytes)) pParams) pure returns((bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],bytes,uint8,bool,(uint64,uint64,uint64)),bool))
func (_Store *StoreCallerSession) PrepareForPdpVerification(pParams PrepareForPdpVerificationParams) (PdpVerificationReturns, error) {
	return _Store.Contract.PrepareForPdpVerification(&_Store.CallOpts, pParams)
}

// VerifyPlotData is a free data retrieval call binding the contract method 0x2e19aeff.
//
// Solidity: function VerifyPlotData(((uint64,uint64,uint64),bytes,uint64) vParams) pure returns(bool)
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
// Solidity: function VerifyPlotData(((uint64,uint64,uint64),bytes,uint64) vParams) pure returns(bool)
func (_Store *StoreSession) VerifyPlotData(vParams VerifyPlotDataParams) (bool, error) {
	return _Store.Contract.VerifyPlotData(&_Store.CallOpts, vParams)
}

// VerifyPlotData is a free data retrieval call binding the contract method 0x2e19aeff.
//
// Solidity: function VerifyPlotData(((uint64,uint64,uint64),bytes,uint64) vParams) pure returns(bool)
func (_Store *StoreCallerSession) VerifyPlotData(vParams VerifyPlotDataParams) (bool, error) {
	return _Store.Contract.VerifyPlotData(&_Store.CallOpts, vParams)
}

// VerifyProofWithMerklePathForFile is a free data retrieval call binding the contract method 0xb6ddeca0.
//
// Solidity: function VerifyProofWithMerklePathForFile((uint64,bytes,bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes) vParams) pure returns(bool)
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
// Solidity: function VerifyProofWithMerklePathForFile((uint64,bytes,bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes) vParams) pure returns(bool)
func (_Store *StoreSession) VerifyProofWithMerklePathForFile(vParams VerifyProofWithMerklePathForFileParams) (bool, error) {
	return _Store.Contract.VerifyProofWithMerklePathForFile(&_Store.CallOpts, vParams)
}

// VerifyProofWithMerklePathForFile is a free data retrieval call binding the contract method 0xb6ddeca0.
//
// Solidity: function VerifyProofWithMerklePathForFile((uint64,bytes,bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes) vParams) pure returns(bool)
func (_Store *StoreCallerSession) VerifyProofWithMerklePathForFile(vParams VerifyProofWithMerklePathForFileParams) (bool, error) {
	return _Store.Contract.VerifyProofWithMerklePathForFile(&_Store.CallOpts, vParams)
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
