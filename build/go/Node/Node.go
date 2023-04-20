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

// TransferState is an auto generated low-level Go binding around an user-defined struct.
type TransferState struct {
	From  common.Address
	To    common.Address
	Value uint64
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"ip\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"port\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"deposit\",\"type\":\"uint64\"}],\"name\":\"DNSNodeRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DNSNodeUnReg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"DnsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"FsError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"NotifyHeaderTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Type\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"Header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"URL\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"NameOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"Desc\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"BlockHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"TTL\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structNameInfo\",\"name\":\"newer\",\"type\":\"tuple\"}],\"name\":\"NotifyNameInfoAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"url\",\"type\":\"bytes\"}],\"name\":\"NotifyNameInfoTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"PDPVerifyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nodeAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"NodeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"CalculateNodePledge\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"Cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeAddr\",\"type\":\"bytes\"}],\"name\":\"GetNodeInfoByNodeAddr\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"NodeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structNodeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetNodeInfoByWalletAddr\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"NodeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structNodeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetNodeList\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"NodeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structNodeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"NodeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"GetPledgeUpdate\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"IsNodeRegisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"NodeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"NodeUpdate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"NodeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"Register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"NodeAddr\",\"type\":\"bytes\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"UpdateNodeInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"WithDrawProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0x34653303\",\"type\":\"bytes32\"}],\"name\":\"c_0x34653303\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractISector\",\"name\":\"_sector\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50616e0980620000216000396000f3fe6080604052600436106100d25760003560e01c8063a315b8741161007f578063dc88870511610059578063dc8887051461028d578063dfae2e44146102a9578063fc16cb88146102d2578063fc2153581461030f576100d2565b8063a315b8741461021d578063aba7239614610239578063d73ea9f014610264576100d2565b8063485cc955116100b0578063485cc9551461018e57806366838994146101b75780639260665f146101f4576100d2565b80631a65374a146100d75780633d031b84146101145780634189abe014610151575b600080fd5b3480156100e357600080fd5b506100fe60048036038101906100f99190615cd5565b61032b565b60405161010b91906165d8565b60405180910390f35b34801561012057600080fd5b5061013b60048036038101906101369190615d68565b6105ea565b60405161014891906165d8565b60405180910390f35b34801561015d57600080fd5b5061017860048036038101906101739190615de5565b610b72565b60405161018591906163c5565b60405180910390f35b34801561019a57600080fd5b506101b560048036038101906101b09190615da9565b610f60565b005b3480156101c357600080fd5b506101de60048036038101906101d99190615cd5565b6111a4565b6040516101eb919061630b565b60405180910390f35b34801561020057600080fd5b5061021b60048036038101906102169190615cd5565b611296565b005b61023760048036038101906102329190615de5565b611c1d565b005b34801561024557600080fd5b5061024e6127e8565b60405161025b91906162e9565b60405180910390f35b34801561027057600080fd5b5061028b60048036038101906102869190615d3f565b612cf8565b005b6102a760048036038101906102a29190615de5565b612cfb565b005b3480156102b557600080fd5b506102d060048036038101906102cb9190615cd5565b6142e8565b005b3480156102de57600080fd5b506102f960048036038101906102f49190615de5565b614e3b565b6040516103069190616615565b60405180910390f35b61032960048036038101906103249190615de5565b614fe8565b005b61033361550e565b61035f7fef3f22caa4c87b9087b325d7344cdd84320a6b63ca3e88c32de754231c08c16660001b612cf8565b61038b7f76db31b7c9bdfcbe738bd1a8012de3a12e89b2399b7a0ec8700ee44933cd59c360001b612cf8565b6103b77f8b6f2e71b5bd2f096c8d852ae13de815fd8c7647b2ab042c54073ba591f8dfd560001b612cf8565b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201805461056190616983565b80601f016020809104026020016040519081016040528092919081815260200182805461058d90616983565b80156105da5780601f106105af576101008083540402835291602001916105da565b820191906000526020600020905b8154815290600101906020018083116105bd57829003601f168201915b5050505050815250509050919050565b6105f261550e565b61061e7f1637f4d28dd5ab84961db3b91c0a80091417611f151c24d324e1ed68c4091c6b60001b612cf8565b61064a7f10e67ac57afcd2768791a2c038a283a742ac8d53fa3ba6294c9a8b7c9a5b293660001b612cf8565b6106767f561c1a9db7ade08f07911a977ac029946539ecc7de0ee445f54b23d51cfa3cd260001b612cf8565b61067e61550e565b6106aa7fb025d79863e48f94745d8c392d8d45795690b736f68a3aa34e27a26766083a5c60001b612cf8565b6106d67f56d3787e3409ef6f298878022278449459b637a1fe6fdd2eac5892c2408c75ca60001b612cf8565b60005b600380549050811015610b0f576107127fbf88e7889db94e4349c991d5c4087de44a62578d78e45d34fe60c89071a46ab660001b612cf8565b61073e7f420dddabf713a6088aad803e92613337e85aaeef8c81aaae71639c108e63b0c260001b612cf8565b600260006003838154811061077c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201805461094c90616983565b80601f016020809104026020016040519081016040528092919081815260200182805461097890616983565b80156109c55780601f1061099a576101008083540402835291602001916109c5565b820191906000526020600020905b8154815290600101906020018083116109a857829003601f168201915b50505050508152505091506109fc7fa1945aa1b2aaa03f611a986acbfe89ed15aa6b5012efd1b739cb1989b7735f1260001b612cf8565b610a287ff68041f540386277d400129c9673b392613bf5f7d56f32ce540b57a22fadb58360001b612cf8565b83805190602001208260c00151805190602001201415610ad057610a6e7f90271c3b1504f12c8cf9ce541d5b48dfee74dd69444e0cf09096878b74cb510060001b612cf8565b610a9a7fa096f5ce7eb0630f7708fc212f2cb5b232d86f6210b708d62a58bc978e6369b460001b612cf8565b610ac67f13e697312dfc8016f04745c5f46b4307a6b03b2b73e604781ba8a43edeb0679360001b612cf8565b8192505050610b6d565b610afc7f5f4319c98a71acf6be0f15419d0b4dc7c34460393b05721bc01e73268078547e60001b612cf8565b8080610b07906169e6565b9150506106d9565b50610b3c7f42ce493a3665fec75d365ed99de1c46b8460c6fc38884d24ffd115e42650b95160001b612cf8565b610b687fd8370223d9d963afa57a35c2343548e5859242f5db110396737be1456e9b8bf160001b612cf8565b809150505b919050565b6000610ba07f53c4a51de6c16253feb94dae132102db37e73f8139a5bb5537c0ab3bd7f774a060001b612cf8565b610bcc7f8da1b3ab0f8dba2ca7ca706e9cb1f4e44486919287d97dc848c3e4f9841ce9e660001b612cf8565b610bf87f9d0f56a503ace172acfe5c6331d710e7dbb185c4e43df76b523a55f94c14ab7060001b612cf8565b6000600260008460a0015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282018054610da890616983565b80601f0160208091040260200160405190810160405280929190818152602001828054610dd490616983565b8015610e215780601f10610df657610100808354040283529160200191610e21565b820191906000526020600020905b815481529060010190602001808311610e0457829003601f168201915b5050505050815250509050610e587f3d2ba02e9d9c96f9a45da09969d3b6d453078794c1e957b7ee08e492962a77e360001b612cf8565b610e847fe962f24aaf1050ffe020f645fc8c664d75e3f6ad9cd09e57245410cc20e6e38b60001b612cf8565b6000610e8f84614e3b565b9050610ebd7f581a8e012ed1b45979e557632251327617c0fbef5d6379e44c78757129ee698760001b612cf8565b610ee97f7afecd4e41659ed400beabd87bec8ef6da154d6c23ce643ed0ae7056f2edc34d60001b612cf8565b600082600001519050610f1e7f5822928daf1ae5cbcdeae78f1f23dffed102c45278af4e8a3363cd148074571e60001b612cf8565b610f4a7f8ddbcf5615d733204c7d42f150c9d418b4bd455bfdea20bd522247852eb556f260001b612cf8565b8082610f5691906167d5565b9350505050919050565b600060019054906101000a900460ff16610f885760008054906101000a900460ff1615610f91565b610f90615205565b5b610fd0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fc790616479565b60405180910390fd5b60008060019054906101000a900460ff161590508015611020576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b61104c7fe58e32a5bfba885754cf1d870599a01c6638e7c66e8619b3009f580927ede5df60001b612cf8565b6110787f41d5f52f4276c04f737067c1f1ba1e1100fb6c78af0c20bb8299934d2c03d1f360001b612cf8565b6110a47f9cd91cc71f2e334e89b921ba89b3f5b77d3c3093bcd3b729ac186e54f43fc19d60001b612cf8565b82600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506111117f83a37eaf2d62dc375b5127792c1fe758fdf5720dffee8221660b206f9fe9a41a60001b612cf8565b61113d7fc96dd02d36543d9cae426551196fb2dc96a9ae4ccbccb04e07ce8062522b77a260001b612cf8565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550801561119f5760008060016101000a81548160ff0219169083151502179055505b505050565b60006111d27f7b3961b9c9f390d880faf3589e167bf1c86ed5b9168d8d5170cb6a73503f59b160001b612cf8565b6111fe7f9f3177a5eeabd0849e9d0096b18df7fb00dc22b5551124a7b7d7d73707ac2cf460001b612cf8565b61122a7f48237c4bc9596dbd25607c40ead5efe4812609abba901c9a43ce5da64f29f37560001b612cf8565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1614159050919050565b6112c27f9253689584c352178145e1ab506747543d222f8cf87ab7cf5fbe11225e0225c460001b612cf8565b6112ee7f3fe3df996f3aa8d6dbc3cb30e97714020c7a01cfd22bf3663fbf96b6eb612f2360001b612cf8565b61131a7f2cbc282cd6d7ce1a79c54c95ca793159b9d46deba469f497f733cc5d0b021f7e60001b612cf8565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff161415611499576113af7fc50143bb531600a5f17400efe21be74ec571938d524574fa5f2705f0256e420760001b612cf8565b6113db7f72dd5a25715b0450959c831e287f8875bd2236eec4d9c647a4e5879f48b57ccc60001b612cf8565b6114077f25887d502421f20b51a7f2257b5e67f8dce71b293e523365eedb48f5f812632560001b612cf8565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5604051611434906164cc565b60405180910390a16114687fcd183911fe5cf213b3477286b052204b02d5aa2394e09b3e9f4eb21685489ffe60001b612cf8565b6114947fbde125ac38bf3626eaa7432c47164c37b793c2f836330e95c5775113cd62ca8460001b612cf8565b611c1a565b6114c57fe269833ce6ecff9aea80f4ae5268abb8eb780b99e7c358f37d4a3c6aafc60eb560001b612cf8565b6114f17f05f08d6a0646b44d7e56d24ce5353f8f1a9f187e41e9fb5f649b94a4bbc3adf860001b612cf8565b61151d7f106389421c7b5ae7d1310310fbb7c2ebe6c200224c38df27ccbe8d0ca4aa288460001b612cf8565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820180546116c990616983565b80601f01602080910402602001604051908101604052809291908181526020018280546116f590616983565b80156117425780601f1061171757610100808354040283529160200191611742565b820191906000526020600020905b81548152906001019060200180831161172557829003601f168201915b50505050508152505090506117797f7aed07885a54d70e65e85a1dc871e4ffc817f46971469a005020da749fbd4b1860001b612cf8565b6117a57f67d5854014e53199153b33ff311fafe334b662c61dd337a2b4a4d7a77368430760001b612cf8565b6000816020015167ffffffffffffffff161115611917576117e87f49fa910bef426dee0b346ec1b3ad811eb5fb48039058c927506f984d7cea38be60001b612cf8565b6118147f80f5e9d518178d4907459915d127de62a3a770763286fac623190d4dd8e136fa60001b612cf8565b6118407f740f9732a7dd60ab2b9da4304e3147d3eed94cdbc068b8b6b66f1f37485652c060001b612cf8565b8060a0015173ffffffffffffffffffffffffffffffffffffffff166108fc826020015167ffffffffffffffff169081150290604051600060405180830381858888f19350505050158015611898573d6000803e3d6000fd5b506118c57f1e8e38da4306c99de1df18dff177e9dcfbf8dc70d7e41029b7681b742f04d45960001b612cf8565b6118f17f35f0647a72aae3d1c567ec9f8faa5dc7e5fc18ff3150c848e5630c4d2e3bc37660001b612cf8565b6000816020019067ffffffffffffffff16908167ffffffffffffffff1681525050611a2e565b6119437f207b1d035febaa8f406797fe315e4a5186b04b9ee3aaa3e645997c51009e34a760001b612cf8565b61196f7f948ddf9303c492191a0fe7d1c4505156cb49873a2aa507800dd4d1521e020c6760001b612cf8565b61199b7fa08f60a0eb396ede7852520f99f813261ea026d75250c92d93dbd7b9486ed25760001b612cf8565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f56040516119c8906164ff565b60405180910390a16119fc7f5ff306945058b452f44334298568b7a9e7840ffcc66a35c4b102d095033f4d9960001b612cf8565b611a287f1bd42386ae2bd789f67a5b5d42224559ebf5add57a0ccd0626459ccb7a6532bc60001b612cf8565b50611c1a565b611a5a7f6f02fc31ef401fd5a877268f655db9b68c0fab78fa69176b7cbecb6589bcf52260001b612cf8565b611a867f0d11f02e441c530714d925088256de78675cbfdd689910cd2a1a2379d876f1ac60001b612cf8565b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160010160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c0820151816002019080519060200190611c14929190615593565b50905050505b50565b611c497fad62471026fa59a2d429f6c91a0cd42686094db98cea2ad0d2e9fc9a39f7111b60001b612cf8565b611c757fb6d2b8c90742a0dcbce9467ebcbaa6a5d296cc2dd771d8e25f8b5d0652f1b85960001b612cf8565b611ca17f78fc087275e9a459389f1d787b1798bfe96150d936bd450a437e06778f9a257760001b612cf8565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b158015611d0b57600080fd5b505afa158015611d1f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d439190615e26565b9050611d717f8e077b5857041924bb52145cfd74356a4f10a659b8d650fa73c76474431d058a60001b612cf8565b611d9d7f7f1767e2ed4db02dcd5f664a55c9cfdc7db7274baedf75ccabf19c6c04fffa2d60001b612cf8565b8060a0015167ffffffffffffffff16826040015167ffffffffffffffff161015611ed857611ded7f03a1a60fe5639c04c3d9ac317e9d7a1f48f8547b4730ef2ee6ccd46d19753c2a60001b612cf8565b611e197f4157d077a83a926b76fe3c2f4700e0d39af5f78a234251b75eb1f6c97155b20f60001b612cf8565b611e457f68955abf1ba1e13f40626f2eca5b834b6b32940ba9332995ca5c788c9e1326b860001b612cf8565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f5604051611e7290616446565b60405180910390a1611ea67f702bc1f677cf21925768b1f46234d1c9e2459e5dc4bfb7391dd5582b7149c61a60001b612cf8565b611ed27f65e8b5ec918cba75163efaf20c9cdf46f8bf1fcb4c3bf4ba2482b36694a295b160001b612cf8565b506127e5565b611f047f6aa01e6c2433c8de17c8972963bc3a1978b99e0354f93f765bbb78d2b789c18860001b612cf8565b611f307fe67b698c1f7d32166daed36a088c729594816eb9060fd6b6575c3a0bbc7b990b60001b612cf8565b611f5c7fcebbd7992b022ad4b6957fe2aabaad0e3aac60a304a8d7d2beee2f7504142ce860001b612cf8565b6000600260008460a0015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff16146120df57611ff47f8f30a3ec9bbbf646d2d3eefe34ff344ef302c42d3472e2d6975c61b457062e9660001b612cf8565b6120207f1c4d6a97ef8b9eed7a6cae6e576edf0785b2e840a8b0dfba33a7995e0ae6da7e60001b612cf8565b61204c7f9e89d109255966a3b5d2761d948f88798261c5aff870d6c1b7f15e7386c4062360001b612cf8565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f560405161207990616413565b60405180910390a16120ad7f449c55fbeaa762bdcdfefdfab4756affc71ea7c1b17eb22f5c46f60f8e151b0960001b612cf8565b6120d97fde65f00bea73f487b51dd4898d2629380873c1e50700b4f94cc70005d4c3c93660001b612cf8565b506127e5565b61210b7f5b1875b7d6bec5bd4df819948ada8438ce2a88161460af669ece3ca192d4b49060001b612cf8565b6121377f2507ccec57d4a448fd0195ef8c50116129b4440eccdc547b0938d2692246b9a660001b612cf8565b6121637f3870f5ae8f505abcf52afc451d893bec607aa1ff2a40f4ff9c5b982875d075ae60001b612cf8565b600061216e83614e3b565b905061219c7fcb60604e025b89d2b1881086c829aa6691fd86af9946c17a7081b9fc75b3f17460001b612cf8565b6121c87fd628fd0333de939f51f2dc5029b4920340d01fc8ee3c6b18b5842ee8a4e4d72460001b612cf8565b8067ffffffffffffffff163410156122f2576122067f2237bb2b232b69b964023b03bfb047169aff71f1f01fd97cf3d133966eccdc6c60001b612cf8565b6122327f9ba7cc8add68a601c31a0e7f3378f66655eb7eeee6da1927ca58b4707f4cfaad60001b612cf8565b61225e7f7a72208f43b730aa7beb227cdd2fe40e1eac43ff487dbb39628dd41f9826a02e60001b612cf8565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f560405161228b906163e0565b60405180910390a16122bf7f48084705b766d2f13dd7d18a8144f7256846a48984ce3b73df5c44e4d7481f8b60001b612cf8565b6122eb7fa3834badf57ee86fef052238b967d241331d634b1dbdcacd8f0d70bceaeea86460001b612cf8565b50506127e5565b61231e7f3e9f9ffd8d757b24910bd106d7ad859c7fe89d19d276f9d8b6d7e99a1d43fdab60001b612cf8565b61234a7ff6efb65aa9720bb7b2d5e91b5e65a22fa2dddf262ffd2f7b1c371fc06caf8aaa60001b612cf8565b6123767fb919127c0959b0c6f8954174c5d763823e466adfa815c35597fe21d7ad02d64260001b612cf8565b80836000019067ffffffffffffffff16908167ffffffffffffffff16815250506123c27fb1a0e4cd2da5df80364fcbe0d16776ef21500631884562787a826607180c3c2360001b612cf8565b6123ee7f33032b0f9de0b5ceb2ebf2fbf2699e58b0eedd63c94dfe6b0ceb6f819b4f589860001b612cf8565b6000836020019067ffffffffffffffff16908167ffffffffffffffff168152505061243b7f856bb5a6aaa3513afec64f740da3cf0158602847d76bee521a7d228c92b0a3de60001b612cf8565b6124677f48bc67afc5329bc0fad4260d5afddfa7002f5d59b532209a75aa355c200966b660001b612cf8565b8260400151836060019067ffffffffffffffff16908167ffffffffffffffff16815250506124b77ffd6b83bc2b5a83edfb17a13c7a84dc6cb4dfba1e2047876a2c10cec9add2f87560001b612cf8565b6124e37f441ce91ed98d467f9fb2d547a965ee4319e7d5bc194098c1b6dd940c9219c34660001b612cf8565b82600260008560a0015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160010160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c0820151816002019080519060200190612675929190615593565b509050506126a57f84d99f1ed9a5e6de92c6bcfa7d9ec455b172c22d44d6bf2f96299da113e14a7760001b612cf8565b6126d17fa7446fb2e32a67718efcdf8f095c12187511f5a16d16d5002271a47c6bcc673c60001b612cf8565b60038360a001519080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506127647f71a5bbfd9a7440b19cff8f1dc48fe2e8095736c7596cfe896282e97d0de2744260001b612cf8565b6127907f703533d128ac7b5ec7036b3e415a7b9634beed32ff4b3f9222eb5482bc7bfe0860001b612cf8565b7f79778bfbb76fd8f676064400dca61e3f85ef0138b7c2944bdbd1c9cb131a05516004438560a001518660c00151876040015188608001516040516127da9695949392919061635d565b60405180910390a150505b50565b60606128167f33ea01eb18e0e61f24ebeab18a2acc5071977a8e0272a72e03ac8d967a7ecb3360001b612cf8565b6128427fe4caf8bd3e865de4c721dd0368945ac8d8e639da667c536bb213732d893390d160001b612cf8565b61286e7fd92fa4d4368f5ba4eb2dcb341efe74bedad3004a03bc4b004a7b57994d3f070460001b612cf8565b600060038054905067ffffffffffffffff8111156128b5577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156128ee57816020015b6128db61550e565b8152602001906001900390816128d35790505b50905061291d7f4bf4b1a4babd8297bdb502ee422cc5a38bca3fdab3a9e76e83a9d21d90ffc25760001b612cf8565b6129497f8a32d71d486d24b1a40f073ef773d89288943ae18092cc523d547dfc264fd2a360001b612cf8565b60005b600380549050811015612c98576129857ff95d5342e0b854d2d11a72ad71f523950ee74b2647239908988c4fbc8597b66560001b612cf8565b6129b17fdaa09a5f5eed60758e1282a0d99dfdd054c50331c8edb7443091531abb4f1d6560001b612cf8565b60026000600383815481106129ef577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282018054612bbf90616983565b80601f0160208091040260200160405190810160405280929190818152602001828054612beb90616983565b8015612c385780601f10612c0d57610100808354040283529160200191612c38565b820191906000526020600020905b815481529060010190602001808311612c1b57829003601f168201915b505050505081525050828281518110612c7a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052508080612c90906169e6565b91505061294c565b50612cc57f389a8e8a8430df2d7b18a459d77cbd485eadca27e9c39654080b612ff6f5037e60001b612cf8565b612cf17faf003cfb7763171586c5b7b48961dd79d1cacdbe7600f8ba45abe557f49afaef60001b612cf8565b8091505090565b50565b612d277fe251c6f740c2fd27a9bb1135c61cdf861f6d68305e147df781d37e5fbd10fbaa60001b612cf8565b612d537f682f779d5730ad76bccbfe72ce41c2d647fda54ce680c3a123e8e4a800508b7b60001b612cf8565b612d7f7f4537d271c5935a47e090f3b8838212c8ceccbef377c62ac4dc69669070cfe1ed60001b612cf8565b612dab7fb9257f7bbc5449e1ec4e8e6316d2e1ffe548e9319807d8f39aa8805ebe76598060001b612cf8565b8060a0015173ffffffffffffffffffffffffffffffffffffffff16600260008360a0015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614612e83576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e7a906165b8565b60405180910390fd5b612eaf7f515a3f6acb375a17feab2617b4b3f82d22410af606bd939b16ac1c171d62c53e60001b612cf8565b612edb7f88564a125dd1d75df269f8fd50c8af84f4645acbf466754c388c707fe9761b2760001b612cf8565b612f077fc04219793d9dede159d278f1f26989c19624b3903b8d354cd8b619a03931bea860001b612cf8565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b158015612f7157600080fd5b505afa158015612f85573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fa99190615e26565b9050612fd77f83a63f5339b46cfc4f1c42574601f415fe08cfda48eedf6349c0d9c94862967260001b612cf8565b6130037f747a43b8fa031c5f51498290ceac77ad15c9b430057d98e58f43bab1029029f060001b612cf8565b8060a0015167ffffffffffffffff16826040015167ffffffffffffffff16101561313e576130537f389d93eaf13274df7a8dd0a373a32a93bfb3c7dc3793bcf8ceb4ab9f0461b1ad60001b612cf8565b61307f7f21ebb0094cc542ddd636391fa29c979df6e0825111a8c6c4379e8491e950a1a060001b612cf8565b6130ab7fcb2a48c3c7e93298951adab09d7587ddcfbc3536a719d0d880db1c63300642f160001b612cf8565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f56040516130d890616585565b60405180910390a161310c7f35c32e8b9ff60127e4cf8023fe38d9385964f2db84a6ae93594da79c621e98d860001b612cf8565b6131387fc20683891576616724398eb8c860bfd3b21d272d0421def38de214ac99f7cc8160001b612cf8565b506142e5565b61316a7faf2a99b6cbbe74089e40958554d3fbd7babcf9a55d4ff73e3fb70d42d138f62160001b612cf8565b6131967f668a11e287c6a3687735c33280de8f0ab351ec1eb3295834fe763c69def2cef760001b612cf8565b6131c27fa77aca42ccae4fd7d4380dd18cb7399a6b80e23fc6fc940cb606a6df8771d16a60001b612cf8565b6000600260008460a0015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1614156133465761325b7f60365d26b6c9e5c734274ed9aee58506eafff61828978e66dbd609e62be7984d60001b612cf8565b6132877f200131c54d6303a1cc52c1fd58f3cb54db7b259b0dc4e365b63b3701c31bdbe160001b612cf8565b6132b37f283a173606f5756d7965d22b02aa21527ad94dbbd018118d980e8f5feb54ff4b60001b612cf8565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f56040516132e090616552565b60405180910390a16133147f4ba893f8d5a156e9327e3ce5aef27b79f9bfc7a50cff7cb0e8cdebeef36d7f5c60001b612cf8565b6133407f55917d7977811e17dbe85528815730b00f30ad63e039203dd8fbdc80ddecc21160001b612cf8565b506142e5565b6133727fc3c2fd1b205a0029bd9c83ca832b1f248bdb6f850cfc63dc6b0235938735aecc60001b612cf8565b61339e7fd82805bd496fb2ddc49aeff2dff8156e133e39c94874933139658c62acfe3b9f60001b612cf8565b6133ca7f872db38be83a668d87cc0c38a94e965a67bbf553a66d6cdfcd13ef13d23f912660001b612cf8565b6000600260008460a0015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201805461357a90616983565b80601f01602080910402602001604051908101604052809291908181526020018280546135a690616983565b80156135f35780601f106135c8576101008083540402835291602001916135f3565b820191906000526020600020905b8154815290600101906020018083116135d657829003601f168201915b505050505081525050905061362a7f7138e7144348e036aa2f70e55e92f707a50808bd727f9114e9cfaf48f4a8cc7260001b612cf8565b6136567fbcd31ea1bf51a4a60a0288e512f3a70176e298e7e359f42dd189105a54ae905c60001b612cf8565b600061366184614e3b565b905061368f7f658f9e6b44f41704ba0639545f503caa34041df4a53d1e3abb9e883710c2896a60001b612cf8565b6136bb7fe8d3eca35324085adc4c942b30cb0dd7db9a8d993af5838660b0f2b3dcadabc460001b612cf8565b6000826000015190506136f07f8aa134eb20c7a8bea51aed561e58780b4bd273c2e841e61bcebd8608559705fd60001b612cf8565b61371c7f11b3f642806dd0f47d63e2600c078cf9536d93a91b940e3f47a69385511a718f60001b612cf8565b613724615619565b6137507f42d4f125dcaecb4f2d7ceb3cbf1ed1c7ecf8962442bdd343cee4a489d01180e460001b612cf8565b61377c7f83711410fc17b8989e0a06f3b05dcb5ae276e9d457e5ae5f2cf8b1924a5eb29060001b612cf8565b8167ffffffffffffffff168367ffffffffffffffff161015613974576137c47f449ee5ec33befd5369c6f4f0cd50fe5747b91bf91b7a92b586d8db6d92e8536160001b612cf8565b6137f07f644452b427cf7d9a8232ffa588677b36e1fe7b877dcde32718962244900ca72e60001b612cf8565b61381c7fbcc3a5a25b4344241f14a577763d2aab6810c8f745f7712432942721472accc660001b612cf8565b30816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506138807f948e4d371ca1c9217749baa84491a8e93241fb039264e21d2e90ad5fbe42fe9b60001b612cf8565b6138ac7f1e4024e7a1d112a8e4e54e8670287e020ffac88d9c3523fa3a770a0443f4864860001b612cf8565b8560a00151816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506139147f768d56b98a86f010dc3f4d7c969f1a4ea8881abc35d826486675b338f37789c660001b612cf8565b6139407f4d61d26df6feb2482e1a6c2052732bc099b7e4adda7b8ef45725240bed968efd60001b612cf8565b8284600001516139509190616851565b816040019067ffffffffffffffff16908167ffffffffffffffff1681525050613bf2565b6139a07f314f7e1dcf0da5bb6301d235644f13cef807aae4b2ced6102372b89f17df63bc60001b612cf8565b6139cc7f7670cd82a5671da6dab74a683a9f9d89d3cc9e18d5eba52e2c51ccb96d67e9f860001b612cf8565b8167ffffffffffffffff168367ffffffffffffffff161115613bc457613a147f463dc71af073c9eb3736757cd63473cc7218ae9e9cbfc4ea488fef60157692a960001b612cf8565b613a407f74cd8482803dd534f16f631399eff07ed68dfb8055cce94fba7cfee782b1efee60001b612cf8565b613a6c7f76971a0a1fc7d799940e22f0a12e24cf10a55939443d25b272a9682b4fc2f1d160001b612cf8565b8560a00151816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050613ad47f4439a590f9a51c93d0fe7ebed89aa7c51383f70f63fd34262e101eae22965a0660001b612cf8565b613b007fb3961941569f3882bed7043c4aed63c2540cf75be61962d41c45c9d76066dd9f60001b612cf8565b30816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050613b647fedc3e5e563615f11c0a4390c3083ac4f98328be4a32df5fedba6568a8c4f863c60001b612cf8565b613b907fedda3a2c4ecd0b0857a2a375add033c7d5b0aa8fc5570a9e8b7df1058f17a7db60001b612cf8565b836000015183613ba09190616851565b816040019067ffffffffffffffff16908167ffffffffffffffff1681525050613bf1565b613bf07ffb84bc8b61c7c2af449319d80acf09e3ef6a0b182d322905a2939e3d65433cbc60001b612cf8565b5b5b613c1e7f94a1ffdc4138175dcb3c79bffeaf3d626f56028fdd4d2d0f8ef14ab1fcd3de1360001b612cf8565b613c4a7f3ed2cfd4e315befdd04b651700c2fb8510fffa3901599645a461634ce0c73b0260001b612cf8565b8167ffffffffffffffff168367ffffffffffffffff1614613f3657613c917fbb5c488b4f775ed788d0c617afa312a24a2fbb9ebda29fea1eaf1ff6ef55d6e060001b612cf8565b613cbd7f470214bcf1cf605e73b0114f0e317f0fc6ebf5b8ef38bad99e49c6a551129d6560001b612cf8565b613ce97f57b6eb91dcf8d1d2ca49e81584b905d1c75c07a68d5acc0478950e920f35d4fb60001b612cf8565b3073ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff161415613e5357613d4d7fad0d6f3d6f75ca01764351420bee26d8b7ceceb1d4986764b7f12436380c975260001b612cf8565b613d797f6340bab76b2b57f7e924568a7c43c243d4a77681c2869666290e708800f7bf7a60001b612cf8565b613da57f226ccfca2ed156a4d3c8a69f7ac9e0efc7bd2f1b5efd7ef90f4a9a6ac9543e6f60001b612cf8565b613dd17fcb662d7717d0df63cead4dcaea800b184a104efd8425e54aff5849fa0cc4388f60001b612cf8565b806040015167ffffffffffffffff16341015613e22576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613e1990616532565b60405180910390fd5b613e4e7ff0ff903e09163aeb8dc89e71818985283a92cee0df873f88696c9a73daef78ea60001b612cf8565b613f31565b613e7f7fbd39273f47debdedd848a8be06a438bd5833076b0cffb19cd47ed3782612284660001b612cf8565b613eab7fd1542b83c69ab4cd957e055e07f75eacc993f5312a2cc343612231a2a2b85c4460001b612cf8565b613ed77f1febcf3d4b34e81a6fc42021b124512cbc3d2901dc2a2d9aa3204ca79fa4205860001b612cf8565b806000015173ffffffffffffffffffffffffffffffffffffffff166108fc826040015167ffffffffffffffff169081150290604051600060405180830381858888f19350505050158015613f2f573d6000803e3d6000fd5b505b613f63565b613f627ff2a04f4782900fdd84486a314032eefcac98457f546845e54963a8f1a89fe88460001b612cf8565b5b613f8f7f5d3044561ea14bc10ffd5782690def71dee0d37a30c1ed38df16bc1aa582562960001b612cf8565b613fbb7f94a2da0497381477e42b918de4e1371786117cb61e4f6f6925caa580140cec2f60001b612cf8565b82866000019067ffffffffffffffff16908167ffffffffffffffff16815250506140077f8176860b7d24d973040ef82c3b46fa0baec3676466b889d4f9bacb20a9e7386f60001b612cf8565b6140337f0b5080df49b2b00e7ce9c59cdf26215142116249ed2357dd1983f5146ec58aec60001b612cf8565b8360200151866020019067ffffffffffffffff16908167ffffffffffffffff16815250506140837f1e8a6af3c7cf7418646d08418574f560fdfb1489c5c522155f496e3c76df63fe60001b612cf8565b6140af7f5cde4068e8bac87da65ab47e5c6e656bcd28b74fb30d37f26fee3545e8401f8560001b612cf8565b8360400151866040015185606001516140c89190616755565b6140d29190616851565b866060019067ffffffffffffffff16908167ffffffffffffffff168152505061411d7fe04c3d3d6434176de7cb17c7cefc342f57cbaffab28b99f84a9a3f7f43ecee6460001b612cf8565b6141497fef2bbbd797194f400816ed833d9111ebb457f3d7a822cf1a89cf8327b51aa7d860001b612cf8565b85600260008860a0015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160010160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160020190805190602001906142db929190615593565b5090505050505050505b50565b6143147f681ef90df3d31f0df5da2d954ad4b895ebfb034f530c2b54389d39c389a815c960001b612cf8565b6143407f8d1fe62c90daab025fa6a3cd27c1427921fd4b2067b5ba21d3ca50d53ef64a9c60001b612cf8565b61436c7f3c07124f0c97f09ea130e08f28a4be10d814fa4f5ea215a77f16e1e831bbc9d560001b612cf8565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1614156144eb576144017feb1be83bfcb8b153ecff5d50af6fc24f138a1bc4f7fd5e615dcc9889225538e360001b612cf8565b61442d7fa5b3fcfe79887dc22081fc80b2aef865aa04503cd03b02f511d3e21920b778aa60001b612cf8565b6144597fd8fe2e424a017faf669a78292b102e53846c39820df73c4ed7c466c86576748060001b612cf8565b7f61cdd2c66b5fad56c57c04cfe850e75dc43bd34ab1d159139ddffedc4f29b5f560405161448690616499565b60405180910390a16144ba7f3126dfcdfa525c775f90a043ad40ebd0776e031fc9fa72cbab4f0c5013ea1f2d60001b612cf8565b6144e67fd1944d33aae65a348cb3953c1b40c269f09bb482ca23c4a05e43a69009606e5960001b612cf8565b614e38565b6145177ff7934eb6f5b7c5932c4cd5ece259dda13ac2453aac7ac865a9d9e691dd8e67a460001b612cf8565b6145437f6519038a4be245e74a615a27652ba15965956d9a84c12377519848825a49946a60001b612cf8565b61456f7f0b5ea19a24d00e1866031fe906f35ec6d5839de681eeeb061e97238b46fe97bf60001b612cf8565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201805461471b90616983565b80601f016020809104026020016040519081016040528092919081815260200182805461474790616983565b80156147945780601f1061476957610100808354040283529160200191614794565b820191906000526020600020905b81548152906001019060200180831161477757829003601f168201915b50505050508152505090506147cb7f878885270d30e62a53a488853c7b89f6998d01c17275891d5cec63c587a7479b60001b612cf8565b6147f77ffc0ea194031ae6e3b44ca69aa5afdf2e2fdaa37cf9eb4e19cf8066dc07103d1860001b612cf8565b6000816000015167ffffffffffffffff1611156148ff5761483a7f1cb08df11a1d1b8a8958d50bc3a27c2450fc341575a1c61615852a13558c776c60001b612cf8565b6148667f3e61f9f1816723e48404337a15f510787861b54c5c414a2550e9a3972069c11360001b612cf8565b6148927f03de7f1b2abf917e87ddab2f8e2b18c561fd23bf26bd757e76aca7749f35d8db60001b612cf8565b8060a0015173ffffffffffffffffffffffffffffffffffffffff166108fc826020015183600001516148c49190616755565b67ffffffffffffffff169081150290604051600060405180830381858888f193505050501580156148f9573d6000803e3d6000fd5b5061492c565b61492b7fad080984b45775d99142442a3fe81896ef6f06c60e73a89afe22abe68ceef4dd60001b612cf8565b5b6149587f032b5872f0b0f7cca1eca8edf6a745d051401b8af1dc3de6ba15fc4f47f4803160001b612cf8565b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160006101000a81549067ffffffffffffffff02191690556000820160086101000a81549067ffffffffffffffff02191690556000820160106101000a81549067ffffffffffffffff02191690556000820160186101000a81549067ffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff02191690556001820160086101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600282016000614a569190615670565b5050614a847fbee9f32fc67a80fb1623e83551513960d03c7868fdddeb0529889baab81e034860001b612cf8565b614ab07f98ddb43b3ca642a17bb2a7dac22f2e33f0d5d43a67c57ebf4c5ac464176a9cd160001b612cf8565b614ab982615216565b614ae57fbbedd5c49748d4f0118ff7c0b09d94aa511181d1450b4c65368323dcc7c7dbac60001b612cf8565b614b117f44d9548ed81f3eaaee73d027f5282917ad4b596e47f73d886f4ebcc557e85c0b60001b612cf8565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e3168f9e8360a001516040518263ffffffff1660e01b8152600401614b7291906162ce565b60006040518083038186803b158015614b8a57600080fd5b505afa158015614b9e573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190614bc79190615cfe565b9050614bf57f14bf3459f582883946f2a2093b8442487ff5c34f4d9c8a55abcd347a2fd8c14660001b612cf8565b614c217f65d948e680a3aed4b8e52a22bcd1a4e2ff2790ab79cb9af2f0431367d4c571df60001b612cf8565b60005b8151811015614da157614c597ff80c07868bab5ffc11e6c5b5c61ad537946ca8cd1cd4d1ebdf92030f13084b3960001b612cf8565b614c857f04463249d4ea727a91c14216062029c9e3c21c40101cd7311aa2bf4414defabe60001b612cf8565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663931bb19a60405180604001604052808660a0015173ffffffffffffffffffffffffffffffffffffffff168152602001858581518110614d27577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015167ffffffffffffffff168152506040518263ffffffff1660e01b8152600401614d5c91906165fa565b600060405180830381600087803b158015614d7657600080fd5b505af1158015614d8a573d6000803e3d6000fd5b505050508080614d99906169e6565b915050614c24565b50614dcd7e6417137c5730a052a9786492ffa94109282d8c26490d74efdc988889e7a0ec60001b612cf8565b614df97fe25e7927bed54ac9965aad7d535811eb84e94e60eae353d5a4253fdfc27d453660001b612cf8565b7f39a789265f7fefbc3aebf3560cea4e1c1f57e6e31faa063f91797173a0daed3760054385604051614e2d93929190616326565b60405180910390a150505b50565b6000614e697f83bf6ae836e2e0b9be0be88c257450758813ff02d7330b85a1ba807143378cee60001b612cf8565b614e957f7165cc5af104a240dda6637d2d99b10da6e55c62f6ba1f7824a39f1dda91069760001b612cf8565b614ec17fbfb5fff3d9750759d3fe130ff3a3b883b95019062a8e41707992a4157107df4760001b612cf8565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b158015614f2b57600080fd5b505afa158015614f3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614f639190615e26565b9050614f917fc7059ac500dd46e07f71f1bb7d784323ef446736ae9c91ba3bf0f2c78d07045060001b612cf8565b614fbd7f056b133d7ef96bfb40101283791e8fb9549089386e07dbbd676f70cf3825daa360001b612cf8565b826040015181602001518260000151614fd69190616793565b614fe09190616793565b915050919050565b6150147fd4552ee1ae0144603460299fbccf2ebfcb4038c1ba5be255d19572d0148a69d560001b612cf8565b6150407f31baa4ec59752275c68fa343874995bdc91bfd0dc3a75e8307c43854f69f9f8b60001b612cf8565b61506c7f480e8327c7b35b07138bb0fc230472a8752ebd8657299466b8ed52033e216d6f60001b612cf8565b80600260008360a0015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160010160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160020190805190602001906151fe929190615593565b5090505050565b6000615210306154fb565b15905090565b6152427f1b107251a590273332773fbb8631c45fa3aa4aa211012570db4dd5a12bb12a5660001b612cf8565b61526e7f4bd65cb3fe5b37d6a475245e1b47294217f59f1680283fc2275b66f109ba4def60001b612cf8565b61529a7fbf952990b71fe255e826d757cfb68f1c3a09c7d7181c64e7cfe3e1c45679953d60001b612cf8565b60005b6003805490508110156154f6576152d67fe12119f1677e86ddbf5936c15be4b207971edc0a6f6756dec34dffff210417eb60001b612cf8565b6153027fb2f3df2936d8804de99e0132fc984d3ebd2b5de47b4ef7cf78cd29d7ce14a64e60001b612cf8565b8173ffffffffffffffffffffffffffffffffffffffff1660038281548110615353577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156154b7576153c67f41ee6ff70163046871559ec91d25b4dfdcb43327983e06b71a9785286106c34b60001b612cf8565b6153f27fbd2eac8de887072eda2c22e5c06ceadbacf2cffa7c50658bb16688fe0b85f17060001b612cf8565b6003818154811061542c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556154857f8711dafddec33f7dc590e7745234a55f38323f53da3ec79449fb18ce26b05f4060001b612cf8565b6154b17fb4046ebb0276ec5df6f25a34c44a987b8b1a9722b262a5375af5b3aacf63faf460001b612cf8565b506154f8565b6154e37fd65a70f13e43fbec4ca6a35beefac8752373566e99c59d55e9f093353450a74f60001b612cf8565b80806154ee906169e6565b91505061529d565b505b50565b600080823b905060008111915050919050565b6040518060e00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b82805461559f90616983565b90600052602060002090601f0160209004810192826155c15760008555615608565b82601f106155da57805160ff1916838001178555615608565b82800160010185558215615608579182015b828111156156075782518255916020019190600101906155ec565b5b50905061561591906156b0565b5090565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff1681525090565b50805461567c90616983565b6000825580601f1061568e57506156ad565b601f0160209004906000526020600020908101906156ac91906156b0565b5b50565b5b808211156156c95760008160009055506001016156b1565b5090565b60006156e06156db84616655565b616630565b905080838252602082019050828560208602820111156156ff57600080fd5b60005b8581101561574957815167ffffffffffffffff81111561572157600080fd5b80860161572e8982615927565b85526020850194506020840193505050600181019050615702565b5050509392505050565b600061576661576184616681565b616630565b9050808382526020820190508285602086028201111561578557600080fd5b60005b858110156157cf57815167ffffffffffffffff8111156157a757600080fd5b8086016157b48982615a58565b85526020850194506020840193505050600181019050615788565b5050509392505050565b60006157ec6157e7846166ad565b616630565b90508281526020810184848401111561580457600080fd5b61580f848285616941565b509392505050565b600061582a615825846166ad565b616630565b90508281526020810184848401111561584257600080fd5b61584d848285616950565b509392505050565b60008135905061586481616d22565b92915050565b60008151905061587981616d22565b92915050565b600082601f83011261589057600080fd5b81516158a08482602086016156cd565b91505092915050565b600082601f8301126158ba57600080fd5b81516158ca848260208601615753565b91505092915050565b6000815190506158e281616d39565b92915050565b6000813590506158f781616d50565b92915050565b600082601f83011261590e57600080fd5b813561591e8482602086016157d9565b91505092915050565b600082601f83011261593857600080fd5b8151615948848260208601615817565b91505092915050565b60008135905061596081616d67565b92915050565b60008135905061597581616d7e565b92915050565b60008151905061598a81616d95565b92915050565b600060e082840312156159a257600080fd5b6159ac60e0616630565b905060006159bc84828501615cab565b60008301525060206159d084828501615cab565b60208301525060406159e484828501615cab565b60408301525060606159f884828501615cab565b6060830152506080615a0c84828501615cab565b60808301525060a0615a2084828501615855565b60a08301525060c082013567ffffffffffffffff811115615a4057600080fd5b615a4c848285016158fd565b60c08301525092915050565b60006101808284031215615a6b57600080fd5b615a76610180616630565b90506000615a868482850161586a565b6000830152506020615a9a84828501615cc0565b6020830152506040615aae84828501615cc0565b6040830152506060615ac284828501615cc0565b6060830152506080615ad68482850161597b565b60808301525060a0615aea84828501615c96565b60a08301525060c0615afe84828501615c96565b60c08301525060e0615b1284828501615cc0565b60e083015250610100615b2784828501615cc0565b61010083015250610120615b3d84828501615cc0565b61012083015250610140615b53848285016158d3565b6101408301525061016082015167ffffffffffffffff811115615b7557600080fd5b615b818482850161587f565b6101608301525092915050565b60006101608284031215615ba157600080fd5b615bac610160616630565b90506000615bbc84828501615cc0565b6000830152506020615bd084828501615cc0565b6020830152506040615be484828501615cc0565b6040830152506060615bf884828501615cc0565b6060830152506080615c0c84828501615cc0565b60808301525060a0615c2084828501615cc0565b60a08301525060c0615c3484828501615cc0565b60c08301525060e0615c4884828501615cc0565b60e083015250610100615c5d84828501615cc0565b61010083015250610120615c7384828501615cc0565b61012083015250610140615c8984828501615cc0565b6101408301525092915050565b600081519050615ca581616da5565b92915050565b600081359050615cba81616dbc565b92915050565b600081519050615ccf81616dbc565b92915050565b600060208284031215615ce757600080fd5b6000615cf584828501615855565b91505092915050565b600060208284031215615d1057600080fd5b600082015167ffffffffffffffff811115615d2a57600080fd5b615d36848285016158a9565b91505092915050565b600060208284031215615d5157600080fd5b6000615d5f848285016158e8565b91505092915050565b600060208284031215615d7a57600080fd5b600082013567ffffffffffffffff811115615d9457600080fd5b615da0848285016158fd565b91505092915050565b60008060408385031215615dbc57600080fd5b6000615dca85828601615951565b9250506020615ddb85828601615966565b9150509250929050565b600060208284031215615df757600080fd5b600082013567ffffffffffffffff811115615e1157600080fd5b615e1d84828501615990565b91505092915050565b60006101608284031215615e3957600080fd5b6000615e4784828501615b8e565b91505092915050565b6000615e5c838361613a565b905092915050565b615e6d81616885565b82525050565b615e7c81616885565b82525050565b6000615e8d826166ee565b615e978185616711565b935083602082028501615ea9856166de565b8060005b85811015615ee55784840389528151615ec68582615e50565b9450615ed183616704565b925060208a01995050600181019050615ead565b50829750879550505050505092915050565b615f0081616897565b82525050565b6000615f11826166f9565b615f1b8185616722565b9350615f2b818560208601616950565b615f3481616aeb565b840191505092915050565b6000615f4a826166f9565b615f548185616733565b9350615f64818560208601616950565b615f6d81616aeb565b840191505092915050565b615f818161692f565b82525050565b615f90816168e4565b82525050565b6000615fa3601383616744565b9150615fae82616afc565b602082019050919050565b6000615fc6601383616744565b9150615fd182616b25565b602082019050919050565b6000615fe9600883616744565b9150615ff482616b4e565b602082019050919050565b600061600c602e83616744565b915061601782616b77565b604082019050919050565b600061602f600683616744565b915061603a82616bc6565b602082019050919050565b6000616052601783616744565b915061605d82616bef565b602082019050919050565b6000616075600e83616744565b915061608082616c18565b602082019050919050565b6000616098601183616744565b91506160a382616c41565b602082019050919050565b60006160bb600a83616744565b91506160c682616c6a565b602082019050919050565b60006160de601383616744565b91506160e982616c93565b602082019050919050565b6000616101601783616744565b915061610c82616cbc565b602082019050919050565b6000616124600b83616744565b915061612f82616ce5565b602082019050919050565b600060e08301600083015161615260008601826162b0565b50602083015161616560208601826162b0565b50604083015161617860408601826162b0565b50606083015161618b60608601826162b0565b50608083015161619e60808601826162b0565b5060a08301516161b160a0860182615e64565b5060c083015184820360c08601526161c98282615f06565b9150508091505092915050565b600060e0830160008301516161ee60008601826162b0565b50602083015161620160208601826162b0565b50604083015161621460408601826162b0565b50606083015161622760608601826162b0565b50608083015161623a60808601826162b0565b5060a083015161624d60a0860182615e64565b5060c083015184820360c08601526162658282615f06565b9150508091505092915050565b6040820160008201516162886000850182615e64565b50602082015161629b60208501826162b0565b50505050565b6162aa81616911565b82525050565b6162b98161691b565b82525050565b6162c88161691b565b82525050565b60006020820190506162e36000830184615e73565b92915050565b600060208201905081810360008301526163038184615e82565b905092915050565b60006020820190506163206000830184615ef7565b92915050565b600060608201905061633b6000830186615f78565b61634860208301856162a1565b6163556040830184615e73565b949350505050565b600060c0820190506163726000830189615f78565b61637f60208301886162a1565b61638c6040830187615e73565b818103606083015261639e8186615f3f565b90506163ad60808301856162bf565b6163ba60a08301846162bf565b979650505050505050565b60006020820190506163da6000830184615f87565b92915050565b600060408201905081810360008301526163f981615fdc565b9050818103602083015261640c81615fb9565b9050919050565b6000604082019050818103600083015261642c81615fdc565b9050818103602083015261643f81616045565b9050919050565b6000604082019050818103600083015261645f81615fdc565b90508181036020830152616472816160d1565b9050919050565b6000602082019050818103600083015261649281615fff565b9050919050565b600060408201905081810360008301526164b281616022565b905081810360208301526164c581615f96565b9050919050565b600060408201905081810360008301526164e581616068565b905081810360208301526164f881615f96565b9050919050565b6000604082019050818103600083015261651881616068565b9050818103602083015261652b81616117565b9050919050565b6000602082019050818103600083015261654b8161608b565b9050919050565b6000604082019050818103600083015261656b816160ae565b9050818103602083015261657e81615f96565b9050919050565b6000604082019050818103600083015261659e816160ae565b905081810360208301526165b1816160d1565b9050919050565b600060208201905081810360008301526165d1816160f4565b9050919050565b600060208201905081810360008301526165f281846161d6565b905092915050565b600060408201905061660f6000830184616272565b92915050565b600060208201905061662a60008301846162bf565b92915050565b600061663a61664b565b905061664682826169b5565b919050565b6000604051905090565b600067ffffffffffffffff8211156166705761666f616abc565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561669c5761669b616abc565b5b602082029050602081019050919050565b600067ffffffffffffffff8211156166c8576166c7616abc565b5b6166d182616aeb565b9050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b60006167608261691b565b915061676b8361691b565b92508267ffffffffffffffff0382111561678857616787616a2f565b5b828201905092915050565b600061679e8261691b565b91506167a98361691b565b92508167ffffffffffffffff04831182151516156167ca576167c9616a2f565b5b828202905092915050565b60006167e0826168e4565b91506167eb836168e4565b9250827fffffffffffffffffffffffffffffffffffffffffffffffff80000000000000000182126000841215161561682657616825616a2f565b5b82677fffffffffffffff01821360008412161561684657616845616a2f565b5b828203905092915050565b600061685c8261691b565b91506168678361691b565b92508282101561687a57616879616a2f565b5b828203905092915050565b6000616890826168f1565b9050919050565b60008115159050919050565b6000819050919050565b60006168b882616885565b9050919050565b60006168ca82616885565b9050919050565b60008190506168df82616d0e565b919050565b60008160070b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600061693a826168d1565b9050919050565b82818337600083830152505050565b60005b8381101561696e578082015181840152602081019050616953565b8381111561697d576000848401525b50505050565b6000600282049050600182168061699b57607f821691505b602082108114156169af576169ae616a8d565b5b50919050565b6169be82616aeb565b810181811067ffffffffffffffff821117156169dd576169dc616abc565b5b80604052505050565b60006169f182616911565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415616a2457616a23616a2f565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f4e6f6465206e6f74207265676973746572656400000000000000000000000000600082015250565b7f496e73756666696369656e7420706c6564676500000000000000000000000000600082015250565b7f5265676973746572000000000000000000000000000000000000000000000000600082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f43616e63656c0000000000000000000000000000000000000000000000000000600082015250565b7f4e6f646520616c72656164792072656769737465726564000000000000000000600082015250565b7f576974684472617750726f666974000000000000000000000000000000000000600082015250565b7f4e6f7420656e6f75676820706c65646765000000000000000000000000000000600082015250565b7f4e6f646555706461746500000000000000000000000000000000000000000000600082015250565b7f566f6c756d6520697320746f6f20736d616c6c00000000000000000000000000600082015250565b7f4e6f64652077616c6c657441646472206368616e676564000000000000000000600082015250565b7f5a65726f2070726f666974000000000000000000000000000000000000000000600082015250565b600b8110616d1f57616d1e616a5e565b5b50565b616d2b81616885565b8114616d3657600080fd5b50565b616d4281616897565b8114616d4d57600080fd5b50565b616d59816168a3565b8114616d6457600080fd5b50565b616d70816168ad565b8114616d7b57600080fd5b50565b616d87816168bf565b8114616d9257600080fd5b50565b60038110616da257600080fd5b50565b616dae81616911565b8114616db957600080fd5b50565b616dc58161691b565b8114616dd057600080fd5b5056fea2646970667358221220cf5e75eb63008a17cc46c785742ae0a999fad428d71be898e051ad1ecdb1e8ef64736f6c63430008040033",
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

// CalculateNodePledge is a free data retrieval call binding the contract method 0xfc16cb88.
//
// Solidity: function CalculateNodePledge((uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo) view returns(uint64)
func (_Store *StoreCaller) CalculateNodePledge(opts *bind.CallOpts, nodeInfo NodeInfo) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "CalculateNodePledge", nodeInfo)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CalculateNodePledge is a free data retrieval call binding the contract method 0xfc16cb88.
//
// Solidity: function CalculateNodePledge((uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo) view returns(uint64)
func (_Store *StoreSession) CalculateNodePledge(nodeInfo NodeInfo) (uint64, error) {
	return _Store.Contract.CalculateNodePledge(&_Store.CallOpts, nodeInfo)
}

// CalculateNodePledge is a free data retrieval call binding the contract method 0xfc16cb88.
//
// Solidity: function CalculateNodePledge((uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo) view returns(uint64)
func (_Store *StoreCallerSession) CalculateNodePledge(nodeInfo NodeInfo) (uint64, error) {
	return _Store.Contract.CalculateNodePledge(&_Store.CallOpts, nodeInfo)
}

// GetNodeInfoByNodeAddr is a free data retrieval call binding the contract method 0x3d031b84.
//
// Solidity: function GetNodeInfoByNodeAddr(bytes nodeAddr) view returns((uint64,uint64,uint64,uint64,uint64,address,bytes))
func (_Store *StoreCaller) GetNodeInfoByNodeAddr(opts *bind.CallOpts, nodeAddr []byte) (NodeInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetNodeInfoByNodeAddr", nodeAddr)

	if err != nil {
		return *new(NodeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(NodeInfo)).(*NodeInfo)

	return out0, err

}

// GetNodeInfoByNodeAddr is a free data retrieval call binding the contract method 0x3d031b84.
//
// Solidity: function GetNodeInfoByNodeAddr(bytes nodeAddr) view returns((uint64,uint64,uint64,uint64,uint64,address,bytes))
func (_Store *StoreSession) GetNodeInfoByNodeAddr(nodeAddr []byte) (NodeInfo, error) {
	return _Store.Contract.GetNodeInfoByNodeAddr(&_Store.CallOpts, nodeAddr)
}

// GetNodeInfoByNodeAddr is a free data retrieval call binding the contract method 0x3d031b84.
//
// Solidity: function GetNodeInfoByNodeAddr(bytes nodeAddr) view returns((uint64,uint64,uint64,uint64,uint64,address,bytes))
func (_Store *StoreCallerSession) GetNodeInfoByNodeAddr(nodeAddr []byte) (NodeInfo, error) {
	return _Store.Contract.GetNodeInfoByNodeAddr(&_Store.CallOpts, nodeAddr)
}

// GetNodeInfoByWalletAddr is a free data retrieval call binding the contract method 0x1a65374a.
//
// Solidity: function GetNodeInfoByWalletAddr(address walletAddr) view returns((uint64,uint64,uint64,uint64,uint64,address,bytes))
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
// Solidity: function GetNodeInfoByWalletAddr(address walletAddr) view returns((uint64,uint64,uint64,uint64,uint64,address,bytes))
func (_Store *StoreSession) GetNodeInfoByWalletAddr(walletAddr common.Address) (NodeInfo, error) {
	return _Store.Contract.GetNodeInfoByWalletAddr(&_Store.CallOpts, walletAddr)
}

// GetNodeInfoByWalletAddr is a free data retrieval call binding the contract method 0x1a65374a.
//
// Solidity: function GetNodeInfoByWalletAddr(address walletAddr) view returns((uint64,uint64,uint64,uint64,uint64,address,bytes))
func (_Store *StoreCallerSession) GetNodeInfoByWalletAddr(walletAddr common.Address) (NodeInfo, error) {
	return _Store.Contract.GetNodeInfoByWalletAddr(&_Store.CallOpts, walletAddr)
}

// GetNodeList is a free data retrieval call binding the contract method 0xaba72396.
//
// Solidity: function GetNodeList() view returns((uint64,uint64,uint64,uint64,uint64,address,bytes)[])
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
// Solidity: function GetNodeList() view returns((uint64,uint64,uint64,uint64,uint64,address,bytes)[])
func (_Store *StoreSession) GetNodeList() ([]NodeInfo, error) {
	return _Store.Contract.GetNodeList(&_Store.CallOpts)
}

// GetNodeList is a free data retrieval call binding the contract method 0xaba72396.
//
// Solidity: function GetNodeList() view returns((uint64,uint64,uint64,uint64,uint64,address,bytes)[])
func (_Store *StoreCallerSession) GetNodeList() ([]NodeInfo, error) {
	return _Store.Contract.GetNodeList(&_Store.CallOpts)
}

// GetPledgeUpdate is a free data retrieval call binding the contract method 0x4189abe0.
//
// Solidity: function GetPledgeUpdate((uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo) view returns(int64)
func (_Store *StoreCaller) GetPledgeUpdate(opts *bind.CallOpts, nodeInfo NodeInfo) (int64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetPledgeUpdate", nodeInfo)

	if err != nil {
		return *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(int64)).(*int64)

	return out0, err

}

// GetPledgeUpdate is a free data retrieval call binding the contract method 0x4189abe0.
//
// Solidity: function GetPledgeUpdate((uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo) view returns(int64)
func (_Store *StoreSession) GetPledgeUpdate(nodeInfo NodeInfo) (int64, error) {
	return _Store.Contract.GetPledgeUpdate(&_Store.CallOpts, nodeInfo)
}

// GetPledgeUpdate is a free data retrieval call binding the contract method 0x4189abe0.
//
// Solidity: function GetPledgeUpdate((uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo) view returns(int64)
func (_Store *StoreCallerSession) GetPledgeUpdate(nodeInfo NodeInfo) (int64, error) {
	return _Store.Contract.GetPledgeUpdate(&_Store.CallOpts, nodeInfo)
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

// C0x34653303 is a free data retrieval call binding the contract method 0xd73ea9f0.
//
// Solidity: function c_0x34653303(bytes32 c__0x34653303) pure returns()
func (_Store *StoreCaller) C0x34653303(opts *bind.CallOpts, c__0x34653303 [32]byte) error {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "c_0x34653303", c__0x34653303)

	if err != nil {
		return err
	}

	return err

}

// C0x34653303 is a free data retrieval call binding the contract method 0xd73ea9f0.
//
// Solidity: function c_0x34653303(bytes32 c__0x34653303) pure returns()
func (_Store *StoreSession) C0x34653303(c__0x34653303 [32]byte) error {
	return _Store.Contract.C0x34653303(&_Store.CallOpts, c__0x34653303)
}

// C0x34653303 is a free data retrieval call binding the contract method 0xd73ea9f0.
//
// Solidity: function c_0x34653303(bytes32 c__0x34653303) pure returns()
func (_Store *StoreCallerSession) C0x34653303(c__0x34653303 [32]byte) error {
	return _Store.Contract.C0x34653303(&_Store.CallOpts, c__0x34653303)
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

// NodeUpdate is a paid mutator transaction binding the contract method 0xdc888705.
//
// Solidity: function NodeUpdate((uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo) payable returns()
func (_Store *StoreTransactor) NodeUpdate(opts *bind.TransactOpts, nodeInfo NodeInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "NodeUpdate", nodeInfo)
}

// NodeUpdate is a paid mutator transaction binding the contract method 0xdc888705.
//
// Solidity: function NodeUpdate((uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo) payable returns()
func (_Store *StoreSession) NodeUpdate(nodeInfo NodeInfo) (*types.Transaction, error) {
	return _Store.Contract.NodeUpdate(&_Store.TransactOpts, nodeInfo)
}

// NodeUpdate is a paid mutator transaction binding the contract method 0xdc888705.
//
// Solidity: function NodeUpdate((uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo) payable returns()
func (_Store *StoreTransactorSession) NodeUpdate(nodeInfo NodeInfo) (*types.Transaction, error) {
	return _Store.Contract.NodeUpdate(&_Store.TransactOpts, nodeInfo)
}

// Register is a paid mutator transaction binding the contract method 0xa315b874.
//
// Solidity: function Register((uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo) payable returns()
func (_Store *StoreTransactor) Register(opts *bind.TransactOpts, nodeInfo NodeInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "Register", nodeInfo)
}

// Register is a paid mutator transaction binding the contract method 0xa315b874.
//
// Solidity: function Register((uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo) payable returns()
func (_Store *StoreSession) Register(nodeInfo NodeInfo) (*types.Transaction, error) {
	return _Store.Contract.Register(&_Store.TransactOpts, nodeInfo)
}

// Register is a paid mutator transaction binding the contract method 0xa315b874.
//
// Solidity: function Register((uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo) payable returns()
func (_Store *StoreTransactorSession) Register(nodeInfo NodeInfo) (*types.Transaction, error) {
	return _Store.Contract.Register(&_Store.TransactOpts, nodeInfo)
}

// UpdateNodeInfo is a paid mutator transaction binding the contract method 0xfc215358.
//
// Solidity: function UpdateNodeInfo((uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo) payable returns()
func (_Store *StoreTransactor) UpdateNodeInfo(opts *bind.TransactOpts, nodeInfo NodeInfo) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "UpdateNodeInfo", nodeInfo)
}

// UpdateNodeInfo is a paid mutator transaction binding the contract method 0xfc215358.
//
// Solidity: function UpdateNodeInfo((uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo) payable returns()
func (_Store *StoreSession) UpdateNodeInfo(nodeInfo NodeInfo) (*types.Transaction, error) {
	return _Store.Contract.UpdateNodeInfo(&_Store.TransactOpts, nodeInfo)
}

// UpdateNodeInfo is a paid mutator transaction binding the contract method 0xfc215358.
//
// Solidity: function UpdateNodeInfo((uint64,uint64,uint64,uint64,uint64,address,bytes) nodeInfo) payable returns()
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
