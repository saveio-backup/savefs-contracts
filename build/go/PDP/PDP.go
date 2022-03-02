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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"HashValue\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveNum\",\"type\":\"uint64\"}],\"internalType\":\"structGenChallengeParams\",\"name\":\"gParams\",\"type\":\"tuple\"}],\"name\":\"GenChallenge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"SectorID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Used\",\"type\":\"uint64\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"FirstProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"NextProveHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TotalBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GroupNum\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"IsPlots\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"FileList\",\"type\":\"bytes[]\"}],\"internalType\":\"structSectorInfo\",\"name\":\"SectorInfo_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"Challenges\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"ProveFileNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"BlockNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Tags\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"MerklePath_\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"PlotData\",\"type\":\"bytes\"}],\"internalType\":\"structSectorProveData\",\"name\":\"ProveData\",\"type\":\"tuple\"}],\"internalType\":\"structPrepareForPdpVerificationParams\",\"name\":\"pParams\",\"type\":\"tuple\"}],\"name\":\"PrepareForPdpVerification\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"FileIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"UpdatedChal\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"Path\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"FileOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"FileDesc\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Privilege\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"FileBlockSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ProveTimes\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"ExpiredHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"CopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Deposit\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"FileProveParam\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"ProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ValidFlag\",\"type\":\"bool\"},{\"internalType\":\"enumStorageType\",\"name\":\"StorageType_\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"RealFileSize\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"PrimaryNodes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"CandidateNodes\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"BlocksRoot\",\"type\":\"bytes\"},{\"internalType\":\"enumProveLevel\",\"name\":\"ProveLevel_\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"IsPlotFile\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structFileInfo\",\"name\":\"FileInfo_\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"Success\",\"type\":\"bool\"}],\"internalType\":\"structPdpVerificationReturns\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"NumberID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"StartNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Nonces\",\"type\":\"uint64\"}],\"internalType\":\"structPlotInfo\",\"name\":\"PlotInfo_\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"PlotData\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"}],\"internalType\":\"structVerifyPlotDataParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"VerifyPlotData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Version\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Proofs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"FileIds\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"Tags\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"Index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"Rand\",\"type\":\"uint32\"}],\"internalType\":\"structChallenge[]\",\"name\":\"Challenges\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"PathLen\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Layer\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Hash\",\"type\":\"bytes\"}],\"internalType\":\"structMerkleNode[]\",\"name\":\"Path\",\"type\":\"tuple[]\"}],\"internalType\":\"structMerklePath[]\",\"name\":\"MerklePath_\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"RootHashes\",\"type\":\"bytes\"}],\"internalType\":\"structVerifyProofWithMerklePathForFileParams\",\"name\":\"vParams\",\"type\":\"tuple\"}],\"name\":\"VerifyProofWithMerklePathForFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061155b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c80639f9ca644116100505780639f9ca644146100a0578063b6ddeca0146100c0578063fd7c9808146100ce57600080fd5b80632e19aeff1461006c5780638129fc1c14610096575b600080fd5b61008061007a366004610c4c565b50600190565b60405161008d91906112f6565b60405180910390f35b61009e6100ef565b005b6100b36100ae366004610c17565b610184565b60405161008d9190611365565b61008061007a366004610c81565b6100e26100dc366004610be2565b50606090565b60405161008d91906112e5565b600054610100900460ff1661010a5760005460ff161561010e565b303b155b61014d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161014490611304565b60405180910390fd5b600054610100900460ff1615801561016f576000805461ffff19166101011790555b8015610181576000805461ff00191690555b50565b61018c6101a1565b6101946101a1565b600160c082015292915050565b6040518060e0016040528060608152602001606081526020016060815260200160608152602001606081526020016102a4604080516102e08101825260608082526000602083018190529282018190528082018390526080820183905260a0820183905260c0820183905260e08201839052610100820183905261012082018390526101408201839052610160820181905261018082018390526101a082018390526101c082018390526101e0820183905261020082018390526102208201819052610240820181905261026082015290610280820190815260006020808301829052604080516060810182528381529182018390528181019290925291015290565b8152600060209091015290565b60006102c46102bf84611392565b611376565b905080838252602082019050828560208602820111156102e357600080fd5b60005b8581101561032957813567ffffffffffffffff81111561030557600080fd5b8086016103128982610552565b8552505060209283019291909101906001016102e6565b5050509392505050565b60006103416102bf84611392565b8381529050602081018260408502810186101561035d57600080fd5b60005b858110156103295781610373888261057e565b84525060209092019160409190910190600101610360565b60006103996102bf84611392565b905080838252602082019050828560208602820111156103b857600080fd5b60005b8581101561032957813567ffffffffffffffff8111156103da57600080fd5b8086016103e7898261064c565b8552505060209283019291909101906001016103bb565b600061040c6102bf84611392565b9050808382526020820190508285602086028201111561042b57600080fd5b60005b8581101561032957813567ffffffffffffffff81111561044d57600080fd5b80860161045a89826106bf565b85525050602092830192919091019060010161042e565b600061047f6102bf846113b6565b90508281526020810184848401111561049757600080fd5b6104a2848285611429565b509392505050565b80356104b5816114da565b92915050565b600082601f8301126104cc57600080fd5b81356104dc8482602086016102b1565b949350505050565b600082601f8301126104f557600080fd5b81356104dc848260208601610333565b600082601f83011261051657600080fd5b81356104dc84826020860161038b565b600082601f83011261053757600080fd5b81356104dc8482602086016103fe565b80356104b5816114ee565b600082601f83011261056357600080fd5b81356104dc848260208601610471565b80356104b5816114f6565b60006040828403121561059057600080fd5b61059a6040611376565b905060006105a88484610bcc565b82525060206105b984848301610bcc565b60208301525092915050565b6000608082840312156105d757600080fd5b6105e16080611376565b905060006105ef84846104aa565b825250602082013567ffffffffffffffff81111561060c57600080fd5b61061884828501610552565b602083015250604061062c84828501610bd7565b604083015250606061064084828501610bd7565b60608301525092915050565b60006060828403121561065e57600080fd5b6106686060611376565b905060006106768484610bd7565b825250602061068784848301610bd7565b602083015250604082013567ffffffffffffffff8111156106a757600080fd5b6106b384828501610552565b60408301525092915050565b6000604082840312156106d157600080fd5b6106db6040611376565b905060006106e98484610bd7565b825250602082013567ffffffffffffffff81111561070657600080fd5b6105b984828501610505565b60006060828403121561072457600080fd5b61072e6060611376565b9050600061073c8484610bd7565b825250602061074d84848301610bd7565b60208301525060406106b384828501610bd7565b60006060828403121561077357600080fd5b61077d6060611376565b9050813567ffffffffffffffff81111561079657600080fd5b6107a2848285016107f7565b825250602082013567ffffffffffffffff8111156107bf57600080fd5b6107cb848285016104e4565b602083015250604082013567ffffffffffffffff8111156107eb57600080fd5b6106b384828501610928565b6000610180828403121561080a57600080fd5b610815610180611376565b9050600061082384846104aa565b825250602061083484848301610bd7565b602083015250604061084884828501610bd7565b604083015250606061085c84828501610bd7565b606083015250608061087084828501610573565b60808301525060a061088484828501610bc1565b60a08301525060c061089884828501610bc1565b60c08301525060e06108ac84828501610bd7565b60e0830152506101006108c184828501610bd7565b610100830152506101206108d784828501610bd7565b610120830152506101406108ed84828501610547565b6101408301525061016082013567ffffffffffffffff81111561090f57600080fd5b61091b848285016104bb565b6101608301525092915050565b600060c0828403121561093a57600080fd5b61094460c0611376565b905060006109528484610bd7565b825250602061096384848301610bd7565b602083015250604082013567ffffffffffffffff81111561098357600080fd5b61098f84828501610552565b604083015250606082013567ffffffffffffffff8111156109af57600080fd5b6109bb84828501610552565b606083015250608082013567ffffffffffffffff8111156109db57600080fd5b6109e784828501610526565b60808301525060a082013567ffffffffffffffff811115610a0757600080fd5b610a1384828501610552565b60a08301525092915050565b600060a08284031215610a3157600080fd5b610a3b6060611376565b90506000610a498484610712565b825250606082013567ffffffffffffffff811115610a6657600080fd5b610a7284828501610552565b60208301525060806106b384828501610bd7565b600060e08284031215610a9857600080fd5b610aa260e0611376565b90506000610ab08484610bd7565b825250602082013567ffffffffffffffff811115610acd57600080fd5b610ad984828501610552565b602083015250604082013567ffffffffffffffff811115610af957600080fd5b610b05848285016104bb565b604083015250606082013567ffffffffffffffff811115610b2557600080fd5b610b31848285016104bb565b606083015250608082013567ffffffffffffffff811115610b5157600080fd5b610b5d848285016104e4565b60808301525060a082013567ffffffffffffffff811115610b7d57600080fd5b610b8984828501610526565b60a08301525060c082013567ffffffffffffffff811115610ba957600080fd5b610bb584828501610552565b60c08301525092915050565b80356104b581611503565b80356104b581611509565b80356104b581611515565b600060208284031215610bf457600080fd5b813567ffffffffffffffff811115610c0b57600080fd5b6104dc848285016105c5565b600060208284031215610c2957600080fd5b813567ffffffffffffffff811115610c4057600080fd5b6104dc84828501610761565b600060208284031215610c5e57600080fd5b813567ffffffffffffffff811115610c7557600080fd5b6104dc84828501610a1f565b600060208284031215610c9357600080fd5b813567ffffffffffffffff811115610caa57600080fd5b6104dc84828501610a86565b6000610cc28383610d09565b505060200190565b6000610cd68383610efd565b9392505050565b6000610ce98383610f41565b505060400190565b6000610cd68383611169565b6000610cd683836111b1565b610d12816113e1565b82525050565b6000610d22825190565b80845260209384019383018060005b83811015610d56578151610d458882610cb6565b975060208301925050600101610d31565b509495945050505050565b6000610d6b825190565b80845260208401935083602082028501610d858560200190565b8060005b85811015610dba5784840389528151610da28582610cca565b94506020830160209a909a0199925050600101610d89565b5091979650505050505050565b6000610dd1825190565b80845260209384019383018060005b83811015610d56578151610df48882610cdd565b975060208301925050600101610de0565b6000610e0f825190565b80845260209384019383018060005b83811015610d56578151610e328882610cdd565b975060208301925050600101610e1e565b6000610e4d825190565b80845260208401935083602082028501610e678560200190565b8060005b85811015610dba5784840389528151610e848582610cf1565b94506020830160209a909a0199925050600101610e6b565b6000610ea6825190565b80845260208401935083602082028501610ec08560200190565b8060005b85811015610dba5784840389528151610edd8582610cfd565b94506020830160209a909a0199925050600101610ec4565b801515610d12565b6000610f07825190565b808452602084019350610f1e818560208601611435565b601f01601f19169290920192915050565b610d1281611413565b610d128161141e565b80516040830190610f5284826112c9565b506020820151610f6560208501826112c9565b50505050565b805161032080845260009190840190610f848282610efd565b9150506020830151610f996020860182610d09565b5060408301518482036040860152610fb18282610efd565b9150506060830151610fc660608601826112d5565b506080830151610fd960808601826112d5565b5060a0830151610fec60a08601826112d5565b5060c0830151610fff60c08601826112d5565b5060e083015161101260e08601826112d5565b506101008301516110276101008601826112c3565b5061012083015161103c6101208601826112d5565b506101408301516110516101408601826112d5565b5061016083015184820361016086015261106b8282610efd565b9150506101808301516110826101808601826112d5565b506101a08301516110976101a08601826112c3565b506101c08301516110ac6101c0860182610ef5565b506101e08301516110c16101e0860182610f38565b506102008301516110d66102008601826112d5565b506102208301518482036102208601526110f08282610d18565b91505061024083015184820361024086015261110c8282610d18565b9150506102608301518482036102608601526111288282610efd565b91505061028083015161113f610280860182610f2f565b506102a08301516111546102a0860182610ef5565b506102c08301516104a26102c086018261128c565b8051600090606084019061117d85826112d5565b50602083015161119060208601826112d5565b50604083015184820360408601526111a88282610efd565b95945050505050565b805160009060408401906111c585826112d5565b50602083015184820360208601526111a88282610e43565b805160e0808452600091908401906111f58282610d61565b9150506020830151848203602086015261120f8282610d61565b915050604083015184820360408601526112298282610dc7565b915050606083015184820360608601526112438282610e9c565b9150506080830151848203608086015261125d8282610efd565b91505060a083015184820360a08601526112778282610f6b565b91505060c08301516104a260c0860182610ef5565b8051606083019061129d84826112d5565b5060208201516112b060208501826112d5565b506040820151610f6560408501826112d5565b80610d12565b63ffffffff8116610d12565b67ffffffffffffffff8116610d12565b60208082528101610cd68184610e05565b602081016104b58284610ef5565b602080825281016104b581602e81527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160208201527f647920696e697469616c697a6564000000000000000000000000000000000000604082015260600190565b60208082528101610cd681846111dd565b600061138160405190565b905061138d8282611461565b919050565b600067ffffffffffffffff8211156113ac576113ac6114a4565b5060209081020190565b600067ffffffffffffffff8211156113d0576113d06114a4565b601f19601f83011660200192915050565b600073ffffffffffffffffffffffffffffffffffffffff82166104b5565b8061138d816114ba565b8061138d816114ca565b60006104b5826113ff565b60006104b582611409565b82818337506000910152565b60005b83811015611450578181015183820152602001611438565b83811115610f655750506000910152565b601f19601f830116810181811067ffffffffffffffff82111715611487576114876114a4565b6040525050565b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600381106101815761018161148e565b600281106101815761018161148e565b6114e3816113e1565b811461018157600080fd5b8015156114e3565b6003811061018157600080fd5b806114e3565b63ffffffff81166114e3565b67ffffffffffffffff81166114e356fea2646970667358221220172302f5d9725aa6a4da70f40130a41b87e9ba5e9aca2b6aeec2fae5e0d3814c64736f6c63430008040033",
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
