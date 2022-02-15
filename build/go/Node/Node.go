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

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughPledge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroProfit\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"CalculateNodePledge\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"Cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"GetNodeInfoByNodeAddr\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetNodeInfoByWalletAddr\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetNodeList\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"IsNodeRegisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"NodeUpdate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"Register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"UpdateNodeInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"WithDrawProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractSector\",\"name\":\"_sector\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612190806100206000396000f3fe6080604052600436106100bc5760003560e01c80639260665f11610074578063c819d86c1161004e578063c819d86c146102ae578063dfae2e44146103ba578063ed326a71146103da57600080fd5b80639260665f1461023f578063aba723961461025f578063ad42030b1461028157600080fd5b80633778febe116100a55780633778febe146100d6578063485cc955146101ca57806366838994146101ea57600080fd5b8063199499c0146100c15780631a65374a146100d6575b600080fd5b6100d46100cf366004611ad7565b6103ed565b005b3480156100e257600080fd5b506101b46100f1366004611a4b565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152506001600160a01b03908116600090815260026020818152604092839020835160e08101855281546001600160401b038082168352600160401b808304821695840195909552600160801b8204811696830196909652600160c01b9004851660608201526001820154948516608082015291909304841660a082015291015490911660c082015290565b6040516101c19190611ec3565b60405180910390f35b3480156101d657600080fd5b506100d46101e5366004611a9d565b6106f0565b3480156101f657600080fd5b50610232610205366004611a4b565b6001600160a01b0316600090815260026020526040902054600160801b90046001600160401b0316151590565b6040516101c19190611d82565b34801561024b57600080fd5b506100d461025a366004611a4b565b6107bb565b34801561026b57600080fd5b50610274610a0f565b6040516101c19190611d71565b34801561028d57600080fd5b506102a161029c366004611ad7565b610bb5565b6040516101c19190611efa565b6100d46102bc366004611ad7565b60a0810180516001600160a01b0390811660009081526002602081815260409283902086518154928801519488015160608901516001600160401b039283166fffffffffffffffffffffffffffffffff1990951694909417600160401b9683168702176fffffffffffffffffffffffffffffffff16600160801b9183169190910277ffffffffffffffffffffffffffffffffffffffffffffffff1617600160c01b9382169390930292909217815560808701516001820180549751919093166001600160e01b0319909716969096179585169093029490941790935560c0909301519290910180546001600160a01b03191692909116919091179055565b3480156103c657600080fd5b506100d46103d5366004611a4b565b610c6b565b6100d46103e8366004611ad7565b610fb0565b80600060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561043d57600080fd5b505afa158015610451573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104759190611af5565b8060a001516001600160401b031682604001516001600160401b031610156104b85760405162461bcd60e51b81526004016104af90611ea3565b60405180910390fd5b60a08301516001600160a01b038116600090815260026020526040902054600160801b90046001600160401b0316156105035760405162461bcd60e51b81526004016104af90611e83565b600061050e85610bb5565b9050806001600160401b03163410156105575734816040517fb16955250000000000000000000000000000000000000000000000000000000081526004016104af929190611edf565b6001600160401b038082168652600060208088018281526040808a018051861660608c0190815260a08c0180516001600160a01b0390811688526002968790528488208e5181549751865195518c16600160c01b0277ffffffffffffffffffffffffffffffffffffffffffffffff968d16600160801b02969096166fffffffffffffffffffffffffffffffff918d16600160401b9081026fffffffffffffffffffffffffffffffff19909b16938e169390931799909917169790971793909317835560808e01805160018086018054865186169a8b026001600160e01b031990911693909d16929092179b909b17905560c08f018051949098018054949092166001600160a01b03199485161790915560038054998a0181559097527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9097018054909116909317909255935191519051925193517f2668f2f26843f9357774392703ed879a5a5ca781dfb60354b056070fed6db808946106e19460049443949093909290611db8565b60405180910390a15050505050565b600054610100900460ff1661070b5760005460ff161561070f565b303b155b61072b5760405162461bcd60e51b81526004016104af90611e22565b600054610100900460ff1615801561074d576000805461ffff19166101011790555b600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff16620100006001600160a01b038681169190910291909117909155600180546001600160a01b03191691841691909117905580156107b6576000805461ff00191690555b505050565b6001600160a01b0381166000908152600260205260409020548190600160801b90046001600160401b03166108025760405162461bcd60e51b81526004016104af90611e12565b6001600160a01b03808316600090815260026020818152604092839020835160e08101855281546001600160401b038082168352600160401b8083048216958401869052600160801b8304821697840197909752600160c01b909104811660608301526001830154908116608083015294909404851660a08501529091015490921660c082015290156108e6578060a001516001600160a01b03166108fc82602001516001600160401b03169081150290604051600060405180830381858888f193505050501580156108d9573d6000803e3d6000fd5b5060006020820152610918565b6040517f154b67a500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0392831660009081526002602081815260409283902084518154928601519486015160608701516001600160401b039283166fffffffffffffffffffffffffffffffff1990951694909417600160401b9683168702176fffffffffffffffffffffffffffffffff16600160801b9183169190910277ffffffffffffffffffffffffffffffffffffffffffffffff1617600160c01b93821693909302929092178155608085015160018201805460a0880151929094166001600160e01b03199094169390931790881690940293909317905560c090920151910180546001600160a01b031916919093161790915550565b6003546060906000906001600160401b03811115610a3d57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610aa457816020015b6040805160e08101825260008082526020808301829052928201819052606082018190526080820181905260a0820181905260c08201528252600019909201910181610a5b5790505b50905060005b600354811015610baf576002600060038381548110610ad957634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101546001600160a01b039081168452838201949094526040928301909120825160e08101845281546001600160401b038082168352600160401b808304821695840195909552600160801b8204811695830195909552600160c01b9004841660608201526001820154938416608082015291909204831660a082015260029091015490911660c08201528251839083908110610b9157634e487b7160e01b600052603260045260246000fd5b60200260200101819052508080610ba7906120a3565b915050610aaa565b50919050565b600080600060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b158015610c0757600080fd5b505afa158015610c1b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3f9190611af5565b9050826040015181602001518260000151610c5a9190611fa9565b610c649190611fa9565b9392505050565b6001600160a01b0381166000908152600260205260409020548190600160801b90046001600160401b0316610cb25760405162461bcd60e51b81526004016104af90611e12565b6001600160a01b03808316600090815260026020818152604092839020835160e08101855281546001600160401b03808216808452600160401b808404831696850196909652600160801b8304821697840197909752600160c01b909104811660608301526001830154908116608083015292909204851660a08301529091015490921660c083015215610d9b578060a001516001600160a01b03166108fc82602001518360000151610d659190611f71565b6001600160401b03169081150290604051600060405180830381858888f19350505050158015610d99573d6000803e3d6000fd5b505b6001600160a01b038316600090815260026020819052604082209182556001820180546001600160e01b03191690550180546001600160a01b0319169055610de2836114d1565b60015460c08201516040517fe3168f9e0000000000000000000000000000000000000000000000000000000081526000926001600160a01b03169163e3168f9e91610e309190600401611d63565b60006040518083038186803b158015610e4857600080fd5b505afa158015610e5c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610e849190810190611a69565b905060005b8151811015610f6d57600160009054906101000a90046001600160a01b03166001600160a01b03166323de5b9860405180604001604052808660c001516001600160a01b03168152602001858581518110610ef457634e487b7160e01b600052603260045260246000fd5b6020026020010151602001516001600160401b03168152506040518263ffffffff1660e01b8152600401610f289190611ed1565b600060405180830381600087803b158015610f4257600080fd5b505af1158015610f56573d6000803e3d6000fd5b505050508080610f65906120a3565b915050610e89565b507f39a789265f7fefbc3aebf3560cea4e1c1f57e6e31faa063f91797173a0daed3760054386604051610fa293929190611d90565b60405180910390a150505050565b80600060029054906101000a90046001600160a01b03166001600160a01b03166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561100057600080fd5b505afa158015611014573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110389190611af5565b8060a001516001600160401b031682604001516001600160401b031610156110725760405162461bcd60e51b81526004016104af90611ea3565b60a08301516001600160a01b038116600090815260026020526040902054600160801b90046001600160401b03166110bc5760405162461bcd60e51b81526004016104af90611e12565b60a08401516001600160a01b03908116600081815260026020526040902060010154600160401b9004909116146111055760405162461bcd60e51b81526004016104af90611eb3565b60a0808501516001600160a01b039081166000908152600260208181526040808420815160e08101835281546001600160401b038082168352600160401b808304821696840196909652600160801b8204811694830194909452600160c01b900483166060820152600182015492831660808201529290910485169582019590955293015490911660c083015261119b86610bb5565b8251604080516060810182526000808252602082018190529181019190915291925090816001600160401b0316836001600160401b0316101561120e5730815260a08801516001600160a01b0316602082015283516111fb908490611fe5565b6001600160401b03166040820152611259565b816001600160401b0316836001600160401b031611156112595760a08801516001600160a01b03168152306020820152835161124a9084611fe5565b6001600160401b031660408201525b816001600160401b0316836001600160401b0316146113055760208101516001600160a01b03163014156112b95780604001516001600160401b03163410156112b45760405162461bcd60e51b81526004016104af90611e93565b611305565b80600001516001600160a01b03166108fc82604001516001600160401b03169081150290604051600060405180830381858888f19350505050158015611303573d6000803e3d6000fd5b505b6001600160401b038084168952602080860151909116908901526040808501519089015160608601516113389190611f71565b6113429190611fe5565b88606001906001600160401b031690816001600160401b03168152505087600260008a60a001516001600160a01b03166001600160a01b0316815260200190815260200160002060008201518160000160006101000a8154816001600160401b0302191690836001600160401b0316021790555060208201518160000160086101000a8154816001600160401b0302191690836001600160401b0316021790555060408201518160000160106101000a8154816001600160401b0302191690836001600160401b0316021790555060608201518160000160186101000a8154816001600160401b0302191690836001600160401b0316021790555060808201518160010160006101000a8154816001600160401b0302191690836001600160401b0316021790555060a08201518160010160086101000a8154816001600160a01b0302191690836001600160a01b0316021790555060c08201518160020160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050505050505050505050565b60005b60035481101561157457816001600160a01b03166003828154811061150957634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b03161415611562576003818154811061154557634e487b7160e01b600052603260045260246000fd5b600091825260209091200180546001600160a01b03191690555050565b8061156c816120a3565b9150506114d4565b5050565b600061158b61158684611f24565b611f08565b905080838252602082019050828560208602820111156115aa57600080fd5b60005b858110156115ef5781516001600160401b038111156115cb57600080fd5b8086016115d88982611715565b8552505060209283019291909101906001016115ad565b5050509392505050565b600061160761158684611f24565b9050808382526020820190508285602086028201111561162657600080fd5b60005b858110156115ef5781516001600160401b0381111561164757600080fd5b80860161165489826117f7565b855250506020928301929190910190600101611629565b600061167961158684611f47565b90508281526020810184848401111561169157600080fd5b61169c84828561204b565b509392505050565b80356116af81612113565b92915050565b80516116af81612113565b600082601f8301126116d157600080fd5b81516116e1848260208601611578565b949350505050565b600082601f8301126116fa57600080fd5b81516116e18482602086016115f9565b80516116af81612127565b600082601f83011261172657600080fd5b81516116e184826020860161166b565b80356116af8161212f565b80516116af81612138565b600060e0828403121561175e57600080fd5b61176860e0611f08565b905060006117768484611a35565b825250602061178784848301611a35565b602083015250604061179b84828501611a35565b60408301525060606117af84828501611a35565b60608301525060806117c384828501611a35565b60808301525060a06117d7848285016116a4565b60a08301525060c06117eb848285016116a4565b60c08301525092915050565b6000610180828403121561180a57600080fd5b611815610180611f08565b9050600061182384846116b5565b825250602061183484848301611a40565b602083015250604061184884828501611a40565b604083015250606061185c84828501611a40565b606083015250608061187084828501611741565b60808301525060a061188484828501611a2a565b60a08301525060c061189884828501611a2a565b60c08301525060e06118ac84828501611a40565b60e0830152506101006118c184828501611a40565b610100830152506101206118d784828501611a40565b610120830152506101406118ed8482850161170a565b610140830152506101608201516001600160401b0381111561190e57600080fd5b61191a848285016116c0565b6101608301525092915050565b6000610160828403121561193a57600080fd5b611945610160611f08565b905060006119538484611a40565b825250602061196484848301611a40565b602083015250604061197884828501611a40565b604083015250606061198c84828501611a40565b60608301525060806119a084828501611a40565b60808301525060a06119b484828501611a40565b60a08301525060c06119c884828501611a40565b60c08301525060e06119dc84828501611a40565b60e0830152506101006119f184828501611a40565b61010083015250610120611a0784828501611a40565b61012083015250610140611a1d84828501611a40565b6101408301525092915050565b80516116af81612145565b80356116af8161214b565b80516116af8161214b565b600060208284031215611a5d57600080fd5b60006116e184846116a4565b600060208284031215611a7b57600080fd5b81516001600160401b03811115611a9157600080fd5b6116e1848285016116e9565b60008060408385031215611ab057600080fd5b6000611abc8585611736565b9250506020611acd85828601611736565b9150509250929050565b600060e08284031215611ae957600080fd5b60006116e1848461174c565b60006101608284031215611b0857600080fd5b60006116e18484611927565b6000611b208383611c98565b505060e00190565b611b3181612009565b82525050565b6000611b41825190565b80845260209384019383018060005b83811015611b75578151611b648882611b14565b975060208301925050600101611b50565b509495945050505050565b801515611b31565b611b318161202f565b601381526000602082017f4e6f6465206e6f74207265676973746572656400000000000000000000000000815291505b5060200190565b601781526000602082017f4e6f646520616c7265616479207265676973746572656400000000000000000081529150611bc1565b601181526000602082017f4e6f7420656e6f75676820706c6564676500000000000000000000000000000081529150611bc1565b601381526000602082017f566f6c756d6520697320746f6f20736d616c6c0000000000000000000000000081529150611bc1565b601781526000602082017f4e6f64652077616c6c657441646472206368616e67656400000000000000000081529150611bc1565b805160e0830190611ca98482611d54565b506020820151611cbc6020850182611d54565b506040820151611ccf6040850182611d54565b506060820151611ce26060850182611d54565b506080820151611cf56080850182611d54565b5060a0820151611d0860a0850182611b28565b5060c0820151611d1b60c0850182611b28565b50505050565b80516040830190611d328482611b28565b506020820151611d1b6020850182611d54565b80611b31565b611b318161203a565b6001600160401b038116611b31565b602081016116af8284611b28565b60208082528101610c648184611b37565b602081016116af8284611b80565b60608101611d9e8286611b88565b611dab6020830185611d45565b6116e16040830184611b28565b60c08101611dc68289611b88565b611dd36020830188611d45565b611de06040830187611b28565b611ded6060830186611b28565b611dfa6080830185611d54565b611e0760a0830184611d54565b979650505050505050565b602080825281016116af81611b91565b602080825281016116af81602e81527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160208201527f647920696e697469616c697a6564000000000000000000000000000000000000604082015260600190565b602080825281016116af81611bc8565b602080825281016116af81611bfc565b602080825281016116af81611c30565b602080825281016116af81611c64565b60e081016116af8284611c98565b604081016116af8284611d21565b60408101611eed8285611d45565b610c646020830184611d4b565b602081016116af8284611d54565b6000611f1360405190565b9050611f1f8282612077565b919050565b60006001600160401b03821115611f3d57611f3d6120ea565b5060209081020190565b60006001600160401b03821115611f6057611f606120ea565b601f19601f83011660200192915050565b60006001600160401b03821691506001600160401b0383169250826001600160401b0303821115611fa457611fa46120be565b500190565b60006001600160401b03821691506001600160401b0383169250816001600160401b030483118215151615611fe057611fe06120be565b500290565b6001600160401b039182169116600082821015612004576120046120be565b500390565b60006001600160a01b0382166116af565b60006116af82612009565b80611f1f81612100565b60006116af82612025565b60006001600160401b0382166116af565b60005b8381101561206657818101518382015260200161204e565b83811115611d1b5750506000910152565b601f19601f83011681018181106001600160401b038211171561209c5761209c6120ea565b6040525050565b60006000198214156120b7576120b76120be565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600a8110612110576121106120d4565b50565b61211c81612009565b811461211057600080fd5b80151561211c565b61211c8161201a565b6003811061211057600080fd5b8061211c565b6001600160401b03811661211c56fea2646970667358221220fccbf0c8730813a7f057d3717e82befdb05d3baf7b22a90e3f19b7b1533f17e664736f6c63430008040033",
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

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _config, address _sector) returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts, _config common.Address, _sector common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize", _config, _sector)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _config, address _sector) returns()
func (_Store *StoreSession) Initialize(_config common.Address, _sector common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _sector)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _config, address _sector) returns()
func (_Store *StoreTransactorSession) Initialize(_config common.Address, _sector common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, _config, _sector)
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
