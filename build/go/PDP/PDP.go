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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"HashValue\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveNum\",\"type\":\"uint64\"}],\"internalType\":\"structGenChallengeParams\",\"name\":\"gParams\",\"type\":\"tuple\"}],\"name\":\"GenChallenge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"SectorInfo_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"Challenges\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"ProveFileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"BlockNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Tags\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"MerklePath_\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"PlotData\",\"type\":\"bytes\"}],\"internalType\":\"structSectorProveData\",\"name\":\"ProveData\",\"type\":\"tuple\"}],\"internalType\":\"structPrepareForPdpVerificationParams\",\"name\":\"pParams\",\"type\":\"tuple\"}],\"name\":\"PrepareForPdpVerification\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"FileIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"UpdatedChal\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"Path\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"FileInfo_\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"Success\",\"type\":\"bool\"}],\"internalType\":\"structPdpVerificationReturns\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"PlotData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"}],\"internalType\":\"structVerifyPlotDataParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"VerifyPlotData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"FileIds\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"Challenges\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"MerklePath_\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"}],\"internalType\":\"structVerifyProofWithMerklePathForFileParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"VerifyProofWithMerklePathForFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0x7b63f42e\",\"type\":\"bytes32\"}],\"name\":\"c_0x7b63f42e\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506120b2806100206000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c8063b6ddeca011610050578063b6ddeca0146100e1578063cca9fb6e14610111578063fd7c98081461012d57610072565b80632e19aeff146100775780638129fc1c146100a75780639f9ca644146100b1575b600080fd5b610091600480360381019061008c91906112ff565b61015d565b60405161009e9190611b0b565b60405180910390f35b6100af6101ec565b005b6100cb60048036038101906100c691906112be565b6102fc565b6040516100d89190611b46565b60405180910390f35b6100fb60048036038101906100f69190611340565b61045a565b6040516101089190611b0b565b60405180910390f35b61012b60048036038101906101269190611254565b6104e9565b005b6101476004803603810190610142919061127d565b6104ec565b6040516101549190611ae9565b60405180910390f35b600061018b7f500ccd357730fd6f355508855feb27a3de0231528b52f7a32a5816aad8aa3ff760001b6104e9565b6101b77f9460b79991bd48c7faed2af6ba324453b2c53bcdcdf2015c76137702c947f50f60001b6104e9565b6101e37f26fb3be9749c602152e3bd920e446d0adf8d2e732bcb3aaa2cc3cb0f7f79d18b60001b6104e9565b60019050919050565b600060019054906101000a900460ff166102145760008054906101000a900460ff161561021d565b61021c6105d4565b5b61025c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025390611b26565b60405180910390fd5b60008060019054906101000a900460ff1615905080156102ac576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b6102d87f9abfbb5ca58067695023328cd39c06b369e6e051346a6a46002d12f913f5d1fd60001b6104e9565b80156102f95760008060016101000a81548160ff0219169083151502179055505b50565b6103046105f8565b6103307ff1d6e8c892a9832748f724d5568d0734b972497e0c8fa5d02ebccac4cf9b9b9260001b6104e9565b61035c7f5efc60bee29c0fab48f02d348a6fb1488c1e4a747ff4674e48884cd07d47d82b60001b6104e9565b6103887f7b75ebb277e9f8a00296301e0334183f97c247ad425fd158c063e923c698c2d360001b6104e9565b6103906105f8565b6103bc7f60e37000abcfd880e47df587a1a2cc4409b97fc516e395e6137913b9ece4d20e60001b6104e9565b6103e87f0bff62c94de60acf4ea545e59bfb2c4219682415f225f09a8b4f31a2feaab20a60001b6104e9565b60018160c00190151590811515815250506104257fbd2a02a0b7f27c7998103ca9d86b6b052974ff079768d2a1a5208afe89709c9f60001b6104e9565b6104517fb8d9996cb9121b45fba1e1e7f36ab481c85fd42627df6f44da0808aa8054661d60001b6104e9565b80915050919050565b60006104887fad239245eae9f9ef967b882c3a2e70abc9c7856dc1475512a284c6ddf2b705c160001b6104e9565b6104b47fc106554398a2df8847db86711e4500039b17dc4f7e5d17c9ef89e686011030c060001b6104e9565b6104e07fbe0ce315415c1093249e555683512ba453f0589baee52f7b03435ae409e7c71560001b6104e9565b60019050919050565b50565b606061051a7ffac7c10c79bc52656f5a435013d57d2e8df6a210f5cbf50d59d8e02704674d5460001b6104e9565b6105467f723c39863d15fd2055cd47bfca0d40683f541d9e01ebd97e7b90ade8695f6f0060001b6104e9565b6105727f804d271d8c17ea46edba27abc7994828bb4ffffc50fb0538838aacc4fd07477e60001b6104e9565b60606105a07fe47f26a3071c363b5827c764e86b7fc35471727f1bfdc858ee28115d3277477860001b6104e9565b6105cb7ec89a30f62f49d1ef51eac6a5216936c9f53411f7e15eab33134c8909e78a6260001b6104e9565b80915050919050565b60006105df306105e5565b15905090565b600080823b905060008111915050919050565b6040518060e00160405280606081526020016060815260200160608152602001606081526020016060815260200161062e61063d565b81526020016000151581525090565b604051806102c0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff1681526020016000815260200160001515815260200160006001811115610754577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001600067ffffffffffffffff1681526020016060815260200160608152602001600060028111156107b2577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016000151581526020016107c86107ce565b81525090565b6040518060600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b600061082061081b84611b8d565b611b68565b9050808382526020820190508285602086028201111561083f57600080fd5b60005b8581101561088957813567ffffffffffffffff81111561086157600080fd5b80860161086e8982610b30565b85526020850194506020840193505050600181019050610842565b5050509392505050565b60006108a66108a184611bb9565b611b68565b905080838252602082019050828560408602820111156108c557600080fd5b60005b858110156108f557816108db8882610b6f565b8452602084019350604083019250506001810190506108c8565b5050509392505050565b600061091261090d84611be5565b611b68565b9050808382526020820190508285602086028201111561093157600080fd5b60005b8581101561097b57813567ffffffffffffffff81111561095357600080fd5b8086016109608982610c47565b85526020850194506020840193505050600181019050610934565b5050509392505050565b600061099861099384611c11565b611b68565b905080838252602082019050828560208602820111156109b757600080fd5b60005b85811015610a0157813567ffffffffffffffff8111156109d957600080fd5b8086016109e68982610cbf565b855260208501945060208401935050506001810190506109ba565b5050509392505050565b6000610a1e610a1984611c3d565b611b68565b905082815260208101848484011115610a3657600080fd5b610a41848285611e89565b509392505050565b600081359050610a5881611fe2565b92915050565b600082601f830112610a6f57600080fd5b8135610a7f84826020860161080d565b91505092915050565b600082601f830112610a9957600080fd5b8135610aa9848260208601610893565b91505092915050565b600082601f830112610ac357600080fd5b8135610ad38482602086016108ff565b91505092915050565b600082601f830112610aed57600080fd5b8135610afd848260208601610985565b91505092915050565b600081359050610b1581611ff9565b92915050565b600081359050610b2a81612010565b92915050565b600082601f830112610b4157600080fd5b8135610b51848260208601610a0b565b91505092915050565b600081359050610b6981612027565b92915050565b600060408284031215610b8157600080fd5b610b8b6040611b68565b90506000610b9b8482850161122a565b6000830152506020610baf8482850161122a565b60208301525092915050565b600060808284031215610bcd57600080fd5b610bd76080611b68565b90506000610be784828501610a49565b600083015250602082013567ffffffffffffffff811115610c0757600080fd5b610c1384828501610b30565b6020830152506040610c278482850161123f565b6040830152506060610c3b8482850161123f565b60608301525092915050565b600060608284031215610c5957600080fd5b610c636060611b68565b90506000610c738482850161123f565b6000830152506020610c878482850161123f565b602083015250604082013567ffffffffffffffff811115610ca757600080fd5b610cb384828501610b30565b60408301525092915050565b600060408284031215610cd157600080fd5b610cdb6040611b68565b90506000610ceb8482850161123f565b600083015250602082013567ffffffffffffffff811115610d0b57600080fd5b610d1784828501610ab2565b60208301525092915050565b600060608284031215610d3557600080fd5b610d3f6060611b68565b90506000610d4f8482850161123f565b6000830152506020610d638482850161123f565b6020830152506040610d778482850161123f565b60408301525092915050565b600060608284031215610d9557600080fd5b610d9f6060611b68565b9050600082013567ffffffffffffffff811115610dbb57600080fd5b610dc784828501610e2b565b600083015250602082013567ffffffffffffffff811115610de757600080fd5b610df384828501610a88565b602083015250604082013567ffffffffffffffff811115610e1357600080fd5b610e1f84828501610f61565b60408301525092915050565b60006101808284031215610e3e57600080fd5b610e49610180611b68565b90506000610e5984828501610a49565b6000830152506020610e6d8482850161123f565b6020830152506040610e818482850161123f565b6040830152506060610e958482850161123f565b6060830152506080610ea984828501610b5a565b60808301525060a0610ebd84828501611215565b60a08301525060c0610ed184828501611215565b60c08301525060e0610ee58482850161123f565b60e083015250610100610efa8482850161123f565b61010083015250610120610f108482850161123f565b61012083015250610140610f2684828501610b06565b6101408301525061016082013567ffffffffffffffff811115610f4857600080fd5b610f5484828501610a5e565b6101608301525092915050565b600060c08284031215610f7357600080fd5b610f7d60c0611b68565b90506000610f8d8482850161123f565b6000830152506020610fa18482850161123f565b602083015250604082013567ffffffffffffffff811115610fc157600080fd5b610fcd84828501610b30565b604083015250606082013567ffffffffffffffff811115610fed57600080fd5b610ff984828501610b30565b606083015250608082013567ffffffffffffffff81111561101957600080fd5b61102584828501610adc565b60808301525060a082013567ffffffffffffffff81111561104557600080fd5b61105184828501610b30565b60a08301525092915050565b600060a0828403121561106f57600080fd5b6110796060611b68565b9050600061108984828501610d23565b600083015250606082013567ffffffffffffffff8111156110a957600080fd5b6110b584828501610b30565b60208301525060806110c98482850161123f565b60408301525092915050565b600060e082840312156110e757600080fd5b6110f160e0611b68565b905060006111018482850161123f565b600083015250602082013567ffffffffffffffff81111561112157600080fd5b61112d84828501610b30565b602083015250604082013567ffffffffffffffff81111561114d57600080fd5b61115984828501610a5e565b604083015250606082013567ffffffffffffffff81111561117957600080fd5b61118584828501610a5e565b606083015250608082013567ffffffffffffffff8111156111a557600080fd5b6111b184828501610a88565b60808301525060a082013567ffffffffffffffff8111156111d157600080fd5b6111dd84828501610adc565b60a08301525060c082013567ffffffffffffffff8111156111fd57600080fd5b61120984828501610b30565b60c08301525092915050565b60008135905061122481612037565b92915050565b6000813590506112398161204e565b92915050565b60008135905061124e81612065565b92915050565b60006020828403121561126657600080fd5b600061127484828501610b1b565b91505092915050565b60006020828403121561128f57600080fd5b600082013567ffffffffffffffff8111156112a957600080fd5b6112b584828501610bbb565b91505092915050565b6000602082840312156112d057600080fd5b600082013567ffffffffffffffff8111156112ea57600080fd5b6112f684828501610d83565b91505092915050565b60006020828403121561131157600080fd5b600082013567ffffffffffffffff81111561132b57600080fd5b6113378482850161105d565b91505092915050565b60006020828403121561135257600080fd5b600082013567ffffffffffffffff81111561136c57600080fd5b611378848285016110d5565b91505092915050565b600061138d83836113ed565b60208301905092915050565b60006113a58383611693565b905092915050565b60006113b9838361170d565b60408301905092915050565b60006113d1838361192e565b905092915050565b60006113e5838361197e565b905092915050565b6113f681611dc9565b82525050565b600061140782611cbe565b6114118185611d41565b935061141c83611c6e565b8060005b8381101561144d5781516114348882611381565b975061143f83611d00565b925050600181019050611420565b5085935050505092915050565b600061146582611cc9565b61146f8185611d52565b93508360208202850161148185611c7e565b8060005b858110156114bd578484038952815161149e8582611399565b94506114a983611d0d565b925060208a01995050600181019050611485565b50829750879550505050505092915050565b60006114da82611cd4565b6114e48185611d63565b93506114ef83611c8e565b8060005b8381101561152057815161150788826113ad565b975061151283611d1a565b9250506001810190506114f3565b5085935050505092915050565b600061153882611cd4565b6115428185611d74565b935061154d83611c8e565b8060005b8381101561157e57815161156588826113ad565b975061157083611d1a565b925050600181019050611551565b5085935050505092915050565b600061159682611cdf565b6115a08185611d85565b9350836020820285016115b285611c9e565b8060005b858110156115ee57848403895281516115cf85826113c5565b94506115da83611d27565b925060208a019950506001810190506115b6565b50829750879550505050505092915050565b600061160b82611cea565b6116158185611d96565b93508360208202850161162785611cae565b8060005b85811015611663578484038952815161164485826113d9565b945061164f83611d34565b925060208a0199505060018101905061162b565b50829750879550505050505092915050565b61167e81611ddb565b82525050565b61168d81611ddb565b82525050565b600061169e82611cf5565b6116a88185611da7565b93506116b8818560208601611e98565b6116c181611f5a565b840191505092915050565b6116d581611e65565b82525050565b6116e481611e77565b82525050565b60006116f7602e83611db8565b915061170282611f6b565b604082019050919050565b6040820160008201516117236000850182611acb565b5060208201516117366020850182611acb565b50505050565b600061030083016000830151848203600086015261175a8282611693565b915050602083015161176f60208601826113ed565b50604083015184820360408601526117878282611693565b915050606083015161179c6060860182611ada565b5060808301516117af6080860182611ada565b5060a08301516117c260a0860182611ada565b5060c08301516117d560c0860182611ada565b5060e08301516117e860e0860182611ada565b506101008301516117fd610100860182611abc565b50610120830151611812610120860182611ada565b50610140830151611827610140860182611ada565b506101608301518482036101608601526118418282611693565b915050610180830151611858610180860182611ada565b506101a083015161186d6101a0860182611abc565b506101c08301516118826101c0860182611675565b506101e08301516118976101e08601826116db565b506102008301516118ac610200860182611ada565b506102208301518482036102208601526118c682826113fc565b9150506102408301518482036102408601526118e282826113fc565b9150506102608301516118f96102608601826116cc565b5061028083015161190e610280860182611675565b506102a08301516119236102a0860182611a7a565b508091505092915050565b60006060830160008301516119466000860182611ada565b5060208301516119596020860182611ada565b50604083015184820360408601526119718282611693565b9150508091505092915050565b60006040830160008301516119966000860182611ada565b50602083015184820360208601526119ae828261158b565b9150508091505092915050565b600060e08301600083015184820360008601526119d8828261145a565b915050602083015184820360208601526119f2828261145a565b91505060408301518482036040860152611a0c82826114cf565b91505060608301518482036060860152611a268282611600565b91505060808301518482036080860152611a408282611693565b91505060a083015184820360a0860152611a5a828261173c565b91505060c0830151611a6f60c0860182611675565b508091505092915050565b606082016000820151611a906000850182611ada565b506020820151611aa36020850182611ada565b506040820151611ab66040850182611ada565b50505050565b611ac581611e37565b82525050565b611ad481611e41565b82525050565b611ae381611e51565b82525050565b60006020820190508181036000830152611b03818461152d565b905092915050565b6000602082019050611b206000830184611684565b92915050565b60006020820190508181036000830152611b3f816116ea565b9050919050565b60006020820190508181036000830152611b6081846119bb565b905092915050565b6000611b72611b83565b9050611b7e8282611ecb565b919050565b6000604051905090565b600067ffffffffffffffff821115611ba857611ba7611f2b565b5b602082029050602081019050919050565b600067ffffffffffffffff821115611bd457611bd3611f2b565b5b602082029050602081019050919050565b600067ffffffffffffffff821115611c0057611bff611f2b565b5b602082029050602081019050919050565b600067ffffffffffffffff821115611c2c57611c2b611f2b565b5b602082029050602081019050919050565b600067ffffffffffffffff821115611c5857611c57611f2b565b5b611c6182611f5a565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611dd482611e17565b9050919050565b60008115159050919050565b6000819050919050565b6000819050611dff82611fba565b919050565b6000819050611e1282611fce565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b600067ffffffffffffffff82169050919050565b6000611e7082611df1565b9050919050565b6000611e8282611e04565b9050919050565b82818337600083830152505050565b60005b83811015611eb6578082015181840152602081019050611e9b565b83811115611ec5576000848401525b50505050565b611ed482611f5a565b810181811067ffffffffffffffff82111715611ef357611ef2611f2b565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b60038110611fcb57611fca611efc565b5b50565b60028110611fdf57611fde611efc565b5b50565b611feb81611dc9565b8114611ff657600080fd5b50565b61200281611ddb565b811461200d57600080fd5b50565b61201981611de7565b811461202457600080fd5b50565b6003811061203457600080fd5b50565b61204081611e37565b811461204b57600080fd5b50565b61205781611e41565b811461206257600080fd5b50565b61206e81611e51565b811461207957600080fd5b5056fea2646970667358221220a757ee802c4279ab1119be25e2f3fdd98d7d7c8e3b7f52b7f8724f586bb83df364736f6c63430008040033",
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
// Solidity: function PrepareForPdpVerification(((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]),(uint32,uint32)[],(uint64,uint64,bytes,bytes,(uint64,(uint64,uint64,bytes)[])[],bytes)) pParams) pure returns((bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)),bool))
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
// Solidity: function PrepareForPdpVerification(((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]),(uint32,uint32)[],(uint64,uint64,bytes,bytes,(uint64,(uint64,uint64,bytes)[])[],bytes)) pParams) pure returns((bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)),bool))
func (_Store *StoreSession) PrepareForPdpVerification(pParams PrepareForPdpVerificationParams) (PdpVerificationReturns, error) {
	return _Store.Contract.PrepareForPdpVerification(&_Store.CallOpts, pParams)
}

// PrepareForPdpVerification is a free data retrieval call binding the contract method 0x9f9ca644.
//
// Solidity: function PrepareForPdpVerification(((address,uint64,uint64,uint64,uint8,uint256,uint256,uint64,uint64,uint64,bool,bytes[]),(uint32,uint32)[],(uint64,uint64,bytes,bytes,(uint64,(uint64,uint64,bytes)[])[],bytes)) pParams) pure returns((bytes[],bytes[],(uint32,uint32)[],(uint64,(uint64,uint64,bytes)[])[],bytes,(bytes,address,bytes,uint64,uint64,uint64,uint64,uint64,uint256,uint64,uint64,bytes,uint64,uint256,bool,uint8,uint64,address[],address[],uint8,bool,(uint64,uint64,uint64)),bool))
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
