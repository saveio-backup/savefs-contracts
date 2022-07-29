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

// TransferState is an auto generated low-level Go binding around an user-defined struct.
type TransferState struct {
	From  common.Address
	To    common.Address
	Value uint64
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DifferenceFileOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"FileNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"FileProveFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileProveNotFileOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FirstUserSpaceOperationError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProfit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeSectorProvedInTimeError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"want\",\"type\":\"uint64\"}],\"name\":\"NotEmptySector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughPledge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"NotEnoughTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"got\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"want\",\"type\":\"uint64\"}],\"name\":\"NotEnoughVolume\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"OpError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParamsError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SectorOpError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SectorProveFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"UserspaceChangeError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserspaceDeleteError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceInsufficientSpace\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"UserspaceWrongExpiredHeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroProfit\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumProveLevel\",\"name\":\"provLevel\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlots\",\"type\":\"bool\"}],\"name\":\"CreateSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"fileHashs\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"DeleteFilesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sectorId\",\"type\":\"uint64\"}],\"name\":\"DeleteSectorEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"FilePDPSuccessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"From\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"To\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"Value\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structTransferState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"GetUpdateCostEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"profit\",\"type\":\"uint64\"}],\"name\":\"ProveFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"volume\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"serviceTime\",\"type\":\"uint64\"}],\"name\":\"RegisterNodeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"sizeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumUserSpaceType\",\"name\":\"countType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"SetUserSpaceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fileSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cost\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPlotFile\",\"type\":\"bool\"}],\"name\":\"StoreFileEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumFsEvent\",\"name\":\"eventType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"UnRegisterNodeEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"CalculateNodePledge\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"Cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"GetNodeInfoByNodeAddr\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"GetNodeInfoByWalletAddr\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetNodeList\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"GetPledgeUpdate\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"IsNodeRegisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"NodeUpdate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"Register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"Pledge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Profit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"Volume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"RestVol\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ServiceTime\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"WalletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"NodeAddr\",\"type\":\"address\"}],\"internalType\":\"structNodeInfo\",\"name\":\"nodeInfo\",\"type\":\"tuple\"}],\"name\":\"UpdateNodeInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"}],\"name\":\"WithDrawProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0x34653303\",\"type\":\"bytes32\"}],\"name\":\"c_0x34653303\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"_config\",\"type\":\"address\"},{\"internalType\":\"contractISector\",\"name\":\"_sector\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506160ea80620000216000396000f3fe6080604052600436106100d25760003560e01c8063a0217c951161007f578063c819d86c11610059578063c819d86c146102a1578063d73ea9f0146102bd578063dfae2e44146102e6578063ed326a711461030f576100d2565b8063a0217c95146101fc578063aba7239614610239578063ad42030b14610264576100d2565b8063485cc955116100b0578063485cc9551461016d57806366838994146101965780639260665f146101d3576100d2565b8063199499c0146100d75780631a65374a146100f35780633778febe14610130575b600080fd5b6100f160048036038101906100ec91906154e0565b61032b565b005b3480156100ff57600080fd5b5061011a60048036038101906101159190615411565b610d97565b6040516101279190615a18565b60405180910390f35b34801561013c57600080fd5b5061015760048036038101906101529190615411565b61101a565b6040516101649190615a18565b60405180910390f35b34801561017957600080fd5b50610194600480360381019061018f91906154a4565b61129d565b005b3480156101a257600080fd5b506101bd60048036038101906101b89190615411565b6114e1565b6040516101ca919061588a565b60405180910390f35b3480156101df57600080fd5b506101fa60048036038101906101f59190615411565b6115d3565b005b34801561020857600080fd5b50610223600480360381019061021e91906154e0565b611e65565b604051610230919061593d565b60405180910390f35b34801561024557600080fd5b5061024e612217565b60405161025b9190615868565b60405180910390f35b34801561027057600080fd5b5061028b600480360381019061028691906154e0565b6126eb565b6040516102989190615a77565b60405180910390f35b6102bb60048036038101906102b691906154e0565b612898565b005b3480156102c957600080fd5b506102e460048036038101906102df919061547b565b612adf565b005b3480156102f257600080fd5b5061030d60048036038101906103089190615411565b612ae2565b005b610329600480360381019061032491906154e0565b6135bb565b005b80600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561039557600080fd5b505afa1580156103a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103cd9190615509565b6103f97fbf8a8c387f8176db333a624231e31562142d0be3628d1ef7bdbf899734dd26b960001b612adf565b6104257f0ad01aa2636e455e0a773960061ca8bbe3d3d532ac39777e046e2decf296a66f60001b612adf565b6104517fc12717b786816122e512ae97b09710b6a1c24af8ca4bae7d72770ca869221c4760001b612adf565b61047d7f941c45d4a978a9e76e54c3c0675aeb2cc06e9834a547825efe6f38c6821833a860001b612adf565b8060a0015167ffffffffffffffff16826040015167ffffffffffffffff1610156104dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d3906159d8565b60405180910390fd5b6105087f1e143db984e2e1e3d33f1aff8ac7696d707fd14ca6236a77323fc955139b41da60001b612adf565b6105347f282641e2f9cd8d00642f9609cf07ba5051fb19c520c3f86bae1dc0059af1fba160001b612adf565b8260a001516105657fb4193ca930d2f8872236e6f00f896a6b842e98b2273ba4e30c4ca3232274387b60001b612adf565b6105917f9d22638cb2dc3b80952544b4bc3386ab1bf623638a69013b16f241d830c8806e60001b612adf565b6105bd7fdea786468597a205c87507fba612951402470ce559f8ca2ba119a1383ff4181560001b612adf565b6105e97fa55438107e8b3a3718aaefca85ce65a58b32f7ca4d1513039b17e8ee36a3c7f460001b612adf565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff161461068c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161068390615998565b60405180910390fd5b6106b87f6021cd2638c35f5a5292d7d54ef0930bdf3788b878c725e9c3f06ad5492b9de160001b612adf565b6106e47fb94ce12b572b71737603f1cbdee0f1ad27d22a6a0f491a19097fe4e9076e431f60001b612adf565b6107107f24cb1ce8d5208e489f13b44f8848fcaa7fadc6025499f9ca627a3e40d4073e7b60001b612adf565b61073c7f6af01478d1b4ec6b150cff696424bd3c58402070a34ee0e32d5b37261e84855860001b612adf565b6107687f1107b63241d1a9cc1dceb78a9741faa83736922f7d639a7c5ff930b67e7244a660001b612adf565b6000610773856126eb565b90506107a17fdaf35e364780582c86fcd8a5d1bab89d359867464b0028052b8b7dfa92b1e4c060001b612adf565b6107cd7f4311452a4c4cf57481d3cd952ba1f88fa7fe2c6813d42b5a0a6b406c6d8552cd60001b612adf565b8067ffffffffffffffff163410156108765761080b7fd3010f812d5e75bf878f8e301fab7b41858f14bbbacf9ccfc9e01bbc1ee18c5860001b612adf565b6108377f67cd530a8f2d4521ac92ca761c8c1d64edd051ad216003f8e859475c234fa2c260001b612adf565b34816040517fb169552500000000000000000000000000000000000000000000000000000000815260040161086d929190615a4e565b60405180910390fd5b6108a27f436408e3689cce5048e9fcd2c968582046996fc40c1650c1c57541e519de235660001b612adf565b6108ce7f11500c6bf67ac595df8106c7d3f2d2d43e21919c879b3a2016dcfddc26bcc2d960001b612adf565b6108fa7febcf8f5aa8e43d52c06abfc16da555c5b4081a1fabac76a97a6150bc72df8d7360001b612adf565b80856000019067ffffffffffffffff16908167ffffffffffffffff16815250506109467f200623418baf773a381da7131a57368b86a6658e89802cfbb79693795cc4b43a60001b612adf565b6109727f3d43407b1741d0a1da4553956dbc57debe0fa820d8bc44a4f17534640ae294f660001b612adf565b6000856020019067ffffffffffffffff16908167ffffffffffffffff16815250506109bf7f4ac56d70fc398c221cdc556193bec62feff22c0c21944b6dd20874e65ed8a4a560001b612adf565b6109eb7f319e778da4e944c197c97845a13ba083944f806000852b66e36366e5acd768c560001b612adf565b8460400151856060019067ffffffffffffffff16908167ffffffffffffffff1681525050610a3b7fd76aa07dac3e7ad3bb66b7ce294a2ac1c77627f28ffaa60e827ec593b75f40ab60001b612adf565b610a677ffa0d9fae6719b6917c62471872c49b773d85488a318738627cc743108e10c8d260001b612adf565b84600260008760a0015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160010160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050610c537fa6070818c3c1fe0962ec77ab808dd9db3212b85dd193bed927ed2dbde032785a60001b612adf565b610c7f7f4bf4c9c931fcd650ee3d6a1ef58804358858cd280a6c0bd60427de27f11c070b60001b612adf565b60038560a001519080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610d127fd1c27a8919c4fb8b04b14d99f20c1d0c9881952eecd4f982ea046b28b34abbcb60001b612adf565b610d3e7f0c7baceb357dd15e3a878f49de0c63074c9d6fe6a5e90d1cb9619b159013525260001b612adf565b7f2668f2f26843f9357774392703ed879a5a5ca781dfb60354b056070fed6db8086004438760a001518860c0015189604001518a60800151604051610d88969594939291906158dc565b60405180910390a15050505050565b610d9f614d97565b610dcb7f23ee019e49504d440e7a225ac3c5f58cfe9fb83cd8462b66a4b4bbbccc5579f760001b612adf565b610df77f11f6bcd50f121d0cf55a82a2760447e1634431d344515b70d23e0f9f4bd621b860001b612adf565b610e237f1a82006246ac2dad0be4ba56452366234f322a2f04a2c76d45a03bb7b839c94b60001b612adf565b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050919050565b611022614d97565b61104e7f2cf9b2a095b1f5edc6d2a9169d75667a60090533fac0fd55fc8cc59a30dc4c8f60001b612adf565b61107a7f0d696201539e237fb3e94dc332b41748e9e4c8654bac37583a93417f53e757f060001b612adf565b6110a67f77a02684cd5dcfa701c023011f31769bf7a299bb07f90c0add8642986bf5a5c260001b612adf565b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050919050565b600060019054906101000a900460ff166112c55760008054906101000a900460ff16156112ce565b6112cd614a8e565b5b61130d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161130490615978565b60405180910390fd5b60008060019054906101000a900460ff16159050801561135d576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b6113897fe9ff7da3bb3c885d9b9b1888132e65511b4c394f7d9ed469c3c5edecb11a9b0160001b612adf565b6113b57fe5983b6e05b3850d428113e33407d5f2739d47fd9bd288b76cdcb95ff797949260001b612adf565b6113e17f4487c56fa051dd21a2746d9e4ee13f8f1091d5339e615d3c2d1a9e334aa22bf860001b612adf565b82600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061144e7f4974bd478cb2b02730a508e12f48f57fdb5cf8e069822ad5ad82229a0a144a6760001b612adf565b61147a7f7bea6d4bd86efac1223cfcb05872273d33186937d6b57b95b682ecb2384caceb60001b612adf565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080156114dc5760008060016101000a81548160ff0219169083151502179055505b505050565b600061150f7fd4da4880a0dacafad998c7bfacbedbc41968adf0489e16b403922d23350bb4ee60001b612adf565b61153b7f22b73fc3f0d90a282a039bbdf67defdfd68623c5efb98659b40f8a66e3e3095160001b612adf565b6115677fb3a33e6900fed3b38e44cf00e580114c73a0169de001a068083e472645a5138c60001b612adf565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1614159050919050565b806116007f64cc528f3b26c3ecdcf22ae75147b31ec378cc9e58c7b2af744a7a6bd800092f60001b612adf565b61162c7f0c23ab576e78509f2fff518b8b54c8576a75603ad454b6c2b8eedda882173eb260001b612adf565b6116587f8e92030e863f71aef0aacd7e07ca10840e0aa5da6128c6b272bbca6bf08648ea60001b612adf565b6116847f6cd292f9418bdc421db4f8b0e161ca4f17e7080706247ddcf83ccef436c8dde060001b612adf565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff161415611728576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161171f90615958565b60405180910390fd5b6117547fc6d8f6e34239c3686b98ae8f5ec421d51eb73ed78c648fc9745fb71598ed2abf60001b612adf565b6117807f6d15d8164848bc9e87aa6483d01a4dc44e29fe8d832228b211c4e4760b34f4c460001b612adf565b6117ac7f48ee2dd6676fa57b6d1f07fbf9c4c19bb850d5953e4dd84257a1c4e0be77763b60001b612adf565b6117d87f69ba76a5d088b02656ee9768d2b992f73a6b9c0dc13ee44ed5d5b1dbfe4fbc6760001b612adf565b6118047fe1ee543e1f43d22adad842937dd63069da64793dcf85be7d69ed4dd32453e85360001b612adf565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050611a247fe627b0df9a171094460c3a10ec8d4efea513500eb9bcade86b837eb21ad0a02960001b612adf565b611a507fae6fd83a758e9f742c97ece6962a8be74184f9f9f720b56fa3911d8ff4260aa060001b612adf565b6000816020015167ffffffffffffffff161115611bc257611a937febada0309f90961d87e1190056eaf5734465ad9f423dc27650422d00ae869e3f60001b612adf565b611abf7f9fcb7b7a95bf1cec84ce91e2d4a196a87ef512f47d70e4d998c0caee37e1ed0d60001b612adf565b611aeb7fd0968a7f29ef53160c482ff11179912a365a179e30af998dceb7468ef6f94c3360001b612adf565b8060a0015173ffffffffffffffffffffffffffffffffffffffff166108fc826020015167ffffffffffffffff169081150290604051600060405180830381858888f19350505050158015611b43573d6000803e3d6000fd5b50611b707f680532113b773a55228b5ed1da4a0b66048c15425347afc0661400b4bb79612360001b612adf565b611b9c7f28f089d137a86ed57adcbe5dd8dbf73c1aa7aa22ebc58e9c8f4b548392db4bb660001b612adf565b6000816020019067ffffffffffffffff16908167ffffffffffffffff1681525050611c4c565b611bee7fc4bca03e25da89c7e725606cc11a5a2ebced4be68b68db2891ebb491343fa56360001b612adf565b611c1a7f7dbf40fd4dcd34ebd06576512dd625fc48daa5e6da23c5daa85ea57aec97ca5160001b612adf565b6040517f154b67a500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611c787f0a4d32c141ffe7ebfffbe00612a805f142abafe48ddfea6cd67828b6f91ad7bb60001b612adf565b611ca47f658184fedc8e8c019682bb785641995402d198ea62a2c3675c4f322f8b697d2e60001b612adf565b80600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160010160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050505050565b6000611e937f1ae249284f98d58dad7e883121467a1035e6953cd42f3046e3fd29e15aa6fe1b60001b612adf565b611ebf7faca7a2b922199faec971e09195b877cb031e9ef43d311ed37786598da90f838860001b612adf565b611eeb7fbf3e49eba964647ec017dcec25bede02e15a357c33ac0e8885acce993242216060001b612adf565b6000600260008460a0015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050905061210f7f7607a0295d55a8eb7ffd039bb4bcce61c08fc52f846a2f3da502add7bb4cc8eb60001b612adf565b61213b7faf1258733bc40b4b42990f0600033f18d9de065032880849763f0e40c95020fd60001b612adf565b6000612146846126eb565b90506121747fa8d20b3d53b4eabcbd9f027d43a7b9b56cdbfd504a65c4482f937c8aab3524c960001b612adf565b6121a07fc890a44d3ef6edaa82d84d9abe345dfd50f2bf67569fdfc733908785cc0e999c60001b612adf565b6000826000015190506121d57f35da2acd194083119c69f764bf9199298f05634bb3c20611059db9d37b5a783660001b612adf565b6122017f6cdffce80685b9661fb5b1fcb7b327dd4b7666fa33d8f06fe840d1e5892d63a760001b612adf565b808261220d9190615c0a565b9350505050919050565b60606122457ff5fdacd4eb389fcd6920d798b3d3bb5bed4d22d5ae6a1875ce4a6abb1ae5baf760001b612adf565b6122717f64fd4cbc133bf8ebfb2efb56aafaf5607c6328f05648844a30f0b0169e976ec460001b612adf565b61229d7f61e4ffa1ce571191a385a0de95ef4ec8e176734b713a86ca4c8a9cc44fe7f97360001b612adf565b600060038054905067ffffffffffffffff8111156122e4577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561231d57816020015b61230a614d97565b8152602001906001900390816123025790505b50905061234c7f77151b9d1289e9fa3bcb7b62ffe5647f54cddebeb51bd2644cf378210c86f5cf60001b612adf565b6123787f4768e41194e7ebb9f2cbee20c1c5ca8aabed5eb3660744857f8f0ca5db634f5360001b612adf565b60005b60038054905081101561268b576123b47fc10ed0880dc5f7703e9d2fe92094fa07e18ccc6456b086f464f8d8497b6ed9bc60001b612adf565b6123e07fbab93353795279e40f4665ab5f46333b0b8c438a8b85f2743f0ba21d7e6f861760001b612adf565b600260006003838154811061241e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505082828151811061266d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181905250808061268390615dec565b91505061237b565b506126b87fe8237261b1fbee57652f16437a0555e6615ea2474d07692b3cb8f1d2120a9c8760001b612adf565b6126e47ffe6b8791df6ae743f76fce8a153579261af45d2ca56765638742a3a6188254e560001b612adf565b8091505090565b60006127197fd4bd31d49534c02c34ae5fdceb669bbb65c404ffe206f7135b5f26d7ae09a04960001b612adf565b6127457ff5e45e2678bfaa91c022107d18f4df7a30f8b792bd4ed62c9be1fa388ecbce7a60001b612adf565b6127717f3ab838b935ee438c818c47b6143dbf7f7c09097df1c6675da7338170046a7d4160001b612adf565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b1580156127db57600080fd5b505afa1580156127ef573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128139190615509565b90506128417ff642072f288b54ea03567f65f0c9797e49425fa94a00aac8ebec4593528c4d1560001b612adf565b61286d7fbfe98675a953cee5c5ea6165501f6e05d4a69791f44dcf329944c03216c92b5360001b612adf565b8260400151816020015182600001516128869190615bc8565b6128909190615bc8565b915050919050565b6128c47f67b02bfed4cd358f9425055f2e477bf888aa1fa6598c50bf529fee816023c8c560001b612adf565b6128f07f65b8b2a86911ca4c240bb94acb2aa59c71307b1d02b28bb3f33833b45737288060001b612adf565b61291c7f24ec539c3a1c343d49136a2a3a347cc516ad8f12ef417d41dae3a2ac4ee4b55e60001b612adf565b80600260008360a0015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160010160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555090505050565b50565b80612b0f7f64cc528f3b26c3ecdcf22ae75147b31ec378cc9e58c7b2af744a7a6bd800092f60001b612adf565b612b3b7f0c23ab576e78509f2fff518b8b54c8576a75603ad454b6c2b8eedda882173eb260001b612adf565b612b677f8e92030e863f71aef0aacd7e07ca10840e0aa5da6128c6b272bbca6bf08648ea60001b612adf565b612b937f6cd292f9418bdc421db4f8b0e161ca4f17e7080706247ddcf83ccef436c8dde060001b612adf565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff161415612c37576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c2e90615958565b60405180910390fd5b612c637fc6d8f6e34239c3686b98ae8f5ec421d51eb73ed78c648fc9745fb71598ed2abf60001b612adf565b612c8f7f6d15d8164848bc9e87aa6483d01a4dc44e29fe8d832228b211c4e4760b34f4c460001b612adf565b612cbb7ff2a73aceb245e9139cc62908f57adcb59a45767405015c6b3e13982ea98d2ab360001b612adf565b612ce77ff8d8e2512ffb6cd98ca9645ba11f720a4e927162572648cd94ceff6e290da52460001b612adf565b612d137faa245e6b393b99ade175e01ef3d589ca0c565a7acefce6742eb6d9461c4f11d660001b612adf565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050612f337f5dfaf366c07acb5333b752ac077e3a28c35cb067bc0bfdd818f5fe4d83458c2260001b612adf565b612f5f7f84c05b4014b4f50d2efc4796fc1a23ea4a664feed1aefd93c682b422058fecf660001b612adf565b6000816000015167ffffffffffffffff16111561306757612fa27f80cce1370effeeb2f64bb6f42ae0bc2beebde5f3e19f717838469899bdda15c860001b612adf565b612fce7f8606372e0432c7bbdf6413578552f2def6b7d16df9ae1da523005afae248857560001b612adf565b612ffa7f14b4ba308e39d329249ae9b72e0664176a1d09221864385c5e9210d7df9cb14f60001b612adf565b8060a0015173ffffffffffffffffffffffffffffffffffffffff166108fc8260200151836000015161302c9190615b8a565b67ffffffffffffffff169081150290604051600060405180830381858888f19350505050158015613061573d6000803e3d6000fd5b50613094565b6130937fd361b025c5d40f51c26ec0474713cdb7621c7733ff7c1761731c730f83e65fdb60001b612adf565b5b6130c07f4bf8b18ce3caf2d007d8d4cf4d8ff4ea11ac519d4f1d6d2ab54272125f49402e60001b612adf565b600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160006101000a81549067ffffffffffffffff02191690556000820160086101000a81549067ffffffffffffffff02191690556000820160106101000a81549067ffffffffffffffff02191690556000820160186101000a81549067ffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff02191690556001820160086101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550506132037f472f4aaca687356100fca86f5a4e7edc470995e30db83127f571d92ed67e389160001b612adf565b61322f7f2ac3da16a56d7863d7a1e6bed7716563c17105ed1d665e62046449dcc1b17a8260001b612adf565b61323883614a9f565b6132647fbccd274cb08120410760fe5a997640e488f88bb66d1d3c1ee5c8a1b5af85b97d60001b612adf565b6132907fbff08fdb9df0b289bc0d0940c53be8b4626a3fa9ef11cbc74b6cd7bec8b9098060001b612adf565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e3168f9e8360c001516040518263ffffffff1660e01b81526004016132f1919061584d565b60006040518083038186803b15801561330957600080fd5b505afa15801561331d573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190613346919061543a565b90506133747fc5ad841f7635095e7b2af6c81ec9e65109ba3c43a7ad370dd31b9dbb98717ba960001b612adf565b6133a07fc47a5abd8eddca33c9735a300d7e4d41f68d0128ba334718ebd8a3b72d1082d060001b612adf565b60005b8151811015613520576133d87f1915eecd08f0b8ee6e6bf2960f591c4caef9a7fb7099f34ce333b21add19f26560001b612adf565b6134047f3153cb85f984b2579b48787df0dc7428316cf9d321d2b35b989331d2a674470a60001b612adf565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663931bb19a60405180604001604052808660c0015173ffffffffffffffffffffffffffffffffffffffff1681526020018585815181106134a6577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516020015167ffffffffffffffff168152506040518263ffffffff1660e01b81526004016134db9190615a33565b600060405180830381600087803b1580156134f557600080fd5b505af1158015613509573d6000803e3d6000fd5b50505050808061351890615dec565b9150506133a3565b5061354d7f46660e5b06b4761991fd65da2308e177270df4f4ee8932bd2c8daaf3f86fa73060001b612adf565b6135797fa59377a7a7474a95865b9a34c15d7fda35f154ec8682db5ec2314c287aa10c0f60001b612adf565b7f39a789265f7fefbc3aebf3560cea4e1c1f57e6e31faa063f91797173a0daed37600543866040516135ad939291906158a5565b60405180910390a150505050565b80600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cf12cf6040518163ffffffff1660e01b81526004016101606040518083038186803b15801561362557600080fd5b505afa158015613639573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061365d9190615509565b6136897fbf8a8c387f8176db333a624231e31562142d0be3628d1ef7bdbf899734dd26b960001b612adf565b6136b57f0ad01aa2636e455e0a773960061ca8bbe3d3d532ac39777e046e2decf296a66f60001b612adf565b6136e17fc12717b786816122e512ae97b09710b6a1c24af8ca4bae7d72770ca869221c4760001b612adf565b61370d7f941c45d4a978a9e76e54c3c0675aeb2cc06e9834a547825efe6f38c6821833a860001b612adf565b8060a0015167ffffffffffffffff16826040015167ffffffffffffffff16101561376c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613763906159d8565b60405180910390fd5b6137987f1e143db984e2e1e3d33f1aff8ac7696d707fd14ca6236a77323fc955139b41da60001b612adf565b6137c47f282641e2f9cd8d00642f9609cf07ba5051fb19c520c3f86bae1dc0059af1fba160001b612adf565b8260a001516137f57f64cc528f3b26c3ecdcf22ae75147b31ec378cc9e58c7b2af744a7a6bd800092f60001b612adf565b6138217f0c23ab576e78509f2fff518b8b54c8576a75603ad454b6c2b8eedda882173eb260001b612adf565b61384d7f8e92030e863f71aef0aacd7e07ca10840e0aa5da6128c6b272bbca6bf08648ea60001b612adf565b6138797f6cd292f9418bdc421db4f8b0e161ca4f17e7080706247ddcf83ccef436c8dde060001b612adf565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff16141561391d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161391490615958565b60405180910390fd5b6139497fc6d8f6e34239c3686b98ae8f5ec421d51eb73ed78c648fc9745fb71598ed2abf60001b612adf565b6139757f6d15d8164848bc9e87aa6483d01a4dc44e29fe8d832228b211c4e4760b34f4c460001b612adf565b6139a17fd6a51745b9d9d8da195a5acce5bb9824e3d606449e6af3a5396f1b612f7d396260001b612adf565b6139cd7f8904a53175f8357cc1eea062d4e022d0f7c7b414dc9e5dd8ac9adcd99b218f2260001b612adf565b6139f97f80f698e14b3abb30f22f9a4aad543b308ac8dadf7cf0dfef3846205d5808b11e60001b612adf565b613a257fbf49df713d17ff4945f747909ba35ac52041f3f76d8e6442d24768ce10bc43b360001b612adf565b8360a0015173ffffffffffffffffffffffffffffffffffffffff16600260008660a0015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614613afd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613af4906159f8565b60405180910390fd5b613b297f26bde2d8f257677c1c495ef6b4b37fb95f09ff6abb9a5d895eb47bcd5aaa730360001b612adf565b613b557fde66975fc066fb4752d8b3a75abb3732ac3b74346c9125e74ea8967bc18ebbdf60001b612adf565b613b817fe8675cb5ba873f8f8c372b060bd0a46b089ab1f96dbc7116fbda46e8dd074ddd60001b612adf565b6000600260008660a0015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050613da57f6d0779394c99ccf49e7cd924e2b872e42ea29530d780a58c3fd764dc19aa578c60001b612adf565b613dd17f21c9da4711da6bc2565734458f9f793957b67028fd9946e681a09d4dab12682160001b612adf565b6000613ddc866126eb565b9050613e0a7f519aa3b5330bbd348623fac4d6b7e8dab5b0ff33ac4d3a4797cac85697b7b89560001b612adf565b613e367f6f8c7dbe6a015ef7edc96ddd8addbda7877f85ad31bfd6f10876f14b5e706e2560001b612adf565b600082600001519050613e6b7f9aa4bf6384111d0608cca02df3bef02bfe4c7271554b7cded6af8f8d9b4095c760001b612adf565b613e977f9f850e05514e6b4b4ef87a5965129689222e702d722320d470b7349d854c495d60001b612adf565b613e9f614e32565b613ecb7fcfc368e25f0ad2a5aa262b567a1b647ff38a9acac735710499d195e229a8ae9760001b612adf565b613ef77f54b40e5cc5ad7291f5c37830b62720565af48ea54f5d20045ccb47fc418c454260001b612adf565b8167ffffffffffffffff168367ffffffffffffffff1610156140ef57613f3f7fc84db016922d97455696982988e930e432a596f0574e3adc15900361ce7f469f60001b612adf565b613f6b7ff7ec8d0f68f3098405c6ceffa0c0529cc5c481138f23e48f389b6c6c7c22c35660001b612adf565b613f977fd8bfda616a43a93b2f0552d83dabf7b7aee9585b8002bec26cac3bd9cd3f47c560001b612adf565b30816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050613ffb7f3bddc323a6e75584cbd56518fdb194ba35a83ae52cf24fef31915a7ce6b68ceb60001b612adf565b6140277fa79fe0a7eed152da8b21dbff7a33f19472725f4d5bfb02452b7c337b3a0fc24a60001b612adf565b8760a00151816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505061408f7fb91072e901ae0246ef927a8778aceae4d4c16952f16592f7728b22d576cc117360001b612adf565b6140bb7f94147e01512c345b5eafde5f46c612706f724bab5b4bbd2e06e59579df3c19e960001b612adf565b8284600001516140cb9190615c86565b816040019067ffffffffffffffff16908167ffffffffffffffff168152505061436d565b61411b7f1b395e699bc6f0c3dd5ba2a3aa9a43a9f62b162d2f5b83f43276c79eabde38f760001b612adf565b6141477f7322a9a25aef793677c7fec72e62ec64ef8c823acf37cd0e07b75fd15aa9879460001b612adf565b8167ffffffffffffffff168367ffffffffffffffff16111561433f5761418f7f3338e77e825fb9bdd10435ff38bb69796590b38b19d8de747fbe39c2c1cca7b960001b612adf565b6141bb7fc60f80b0cdaa4c9fab3d393803dd1329e65302960f8cf1efff3c199a6bdac30660001b612adf565b6141e77fba3a2a0763984b7c2b7e41afa6e894a607221f691a01347d1cafa825c8d677a960001b612adf565b8760a00151816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505061424f7f07a991de73aa1f795531299477e73d06bf0e84a72e0f4cc9170f9090a30cab8060001b612adf565b61427b7f494ba8e4cc137775ba2ee1e3ccdbd40452832699672f8fcfa5f2d3887a9f34e260001b612adf565b30816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506142df7f49b5ab08b14325e45e08b6cb9a1e8b7cdbfba3f839a50e5453d8b1357e28ea5c60001b612adf565b61430b7f7afb7cf5372fede9faea20523878af8b8d3384804f2bf55bac9d43401442060260001b612adf565b83600001518361431b9190615c86565b816040019067ffffffffffffffff16908167ffffffffffffffff168152505061436c565b61436b7fe2ad526ef8a56c6139a0486df8e0852b4495c3d8a57d4669ae8183a5f5800d5460001b612adf565b5b5b6143997fe095c1a363521a26563d4fd2479d6c442ee089b6187c1f1b1e53dbe4e0137ac060001b612adf565b6143c57fbe3750e339430aa2abed2da0cb0d78eda8d65cb9b7b6a58d118cabf43c2922f360001b612adf565b8167ffffffffffffffff168367ffffffffffffffff16146146b15761440c7f7c9cee8112f0da947d4eff65637fd0860d6323dbdd691ad9f2ea48947e22b73560001b612adf565b6144387fcd92c7abb190aa9aa0a47a1c28f3dd1dcfad603ab1ab22a4946e6fc6041c004360001b612adf565b6144647f34cb71ba48180cdf4f55936019b90ad048b05b865d42d90b68cdeda5240c09e760001b612adf565b3073ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614156145ce576144c87fd18d927f24dbb238659fc8acf6635ca79b6b84ed2e5b0c61bf7253ff08ae9ff160001b612adf565b6144f47f5d5bfbb36ca0af6b0cbebd9f2db3b426f906e3d7a90ce98ba897009259b0c4b160001b612adf565b6145207f840e1bebd22d3afd14ee12fd91991910f267674c18c361416102f49b48a78a5360001b612adf565b61454c7ff19174a8b37589f7b27192bf99ef65f389b63d0d7185cb6e25e6542b1afc748b60001b612adf565b806040015167ffffffffffffffff1634101561459d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614594906159b8565b60405180910390fd5b6145c97f5636d991e8ecaacf9a6b1599a19935ef12211c97ed70596655326a71784ebf6260001b612adf565b6146ac565b6145fa7f8ca0d304290010c58189634708506a30e54d707e82adcdd65eb5fd78a76b509c60001b612adf565b6146267f8783dc37fb8374ea02268ce1b563a4cdf7a54d17a68df791941dfa8ac84f2df560001b612adf565b6146527fb48cc303a2e3e4810cf11051efb57eda4d894046e891e72163e43a42a6c492eb60001b612adf565b806000015173ffffffffffffffffffffffffffffffffffffffff166108fc826040015167ffffffffffffffff169081150290604051600060405180830381858888f193505050501580156146aa573d6000803e3d6000fd5b505b6146de565b6146dd7ff5c97d4f01edf8f80fa438cb21f219b50050c1f5a2da7f4a1fec7f9e29adb53560001b612adf565b5b61470a7f9a7c54e639a28ed73c9a8d33e095b4e0894b56a2c760f5e44f4f682df5bc095f60001b612adf565b6147367fe99049203297eabafc4ad0dff2737b85969d526855aa38b700484f91c66c1bac60001b612adf565b82886000019067ffffffffffffffff16908167ffffffffffffffff16815250506147827f63f17f447c89ac802de0f7cd3cf359f349fdce066bbb51866a33912f75a1788660001b612adf565b6147ae7f1c59f35f8a8892a64762b49585e2921f74c4b7c421875f86286d6fcc3418363360001b612adf565b8360200151886020019067ffffffffffffffff16908167ffffffffffffffff16815250506147fe7fd9562eb0e64d3eda6a6a4001fcff0cfccc287d0c2dffb5212f34510438d2432060001b612adf565b61482a7f721c0c119fcfd905a336d0cb5ecf89157ff11a966a25dee01d9dc79bbddd383660001b612adf565b8360400151886040015185606001516148439190615b8a565b61484d9190615c86565b886060019067ffffffffffffffff16908167ffffffffffffffff16815250506148987f30b22a855e730fe99d801efaee4366c5bc3dcfa4081f5b49678a0d09082b0f0c60001b612adf565b6148c47f51f51f3f872cfe49b295f312208feae6bdd051aec49c53afd386faa546dce87960001b612adf565b87600260008a60a0015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160010160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050505050505050505050565b6000614a9930614d84565b15905090565b614acb7f02655aa970313a85b4b835d25eab764df0d67b060903628b7fe1250d2f6881a660001b612adf565b614af77fc674ab4565bf7df037e3152a02a47e4bab1845f0b82782f890943bc5d4fbe4f860001b612adf565b614b237fb668bd8a1f6728e6196824512fddf7cb96db208deb76c9bab7423540b0871f0b60001b612adf565b60005b600380549050811015614d7f57614b5f7ff28dd60d060420e9c4a9866fa3932b6551edb164c30eeaa53c750b4b20c2ed1060001b612adf565b614b8b7f13e5be30fc83247e5447967c03cebe81c431ff3e9c4d1e717c01c733dc90089360001b612adf565b8173ffffffffffffffffffffffffffffffffffffffff1660038281548110614bdc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415614d4057614c4f7f11d1bdb4faed2e46c0ec6a6c45a07e634a751b0ff52307ab0e9ec2a19541e7b260001b612adf565b614c7b7fbdf0cbb4f5c69f6cb9c1e0bf2ba47d8c5a03c60b187cc0d26d01189d310fd40760001b612adf565b60038181548110614cb5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055614d0e7fe861081cd79b21edef5f62faa510e88cf62c42d68e8a53b505a51642ff0b3fb860001b612adf565b614d3a7f4adc6ad755503c0c979c7e90129d15c7ca5f77b51e9bb6c3ff0f3b6a2a304a9160001b612adf565b50614d81565b614d6c7f111089d887a2d72633f6935903e475fd45a36135379343e21271795e84ce575360001b612adf565b8080614d7790615dec565b915050614b26565b505b50565b600080823b905060008111915050919050565b6040518060e00160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681525090565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6000614e9c614e9784615ab7565b615a92565b90508083825260208201905082856020860282011115614ebb57600080fd5b60005b85811015614f0557815167ffffffffffffffff811115614edd57600080fd5b808601614eea898261507b565b85526020850194506020840193505050600181019050614ebe565b5050509392505050565b6000614f22614f1d84615ae3565b615a92565b90508083825260208201905082856020860282011115614f4157600080fd5b60005b85811015614f8b57815167ffffffffffffffff811115614f6357600080fd5b808601614f708982615194565b85526020850194506020840193505050600181019050614f44565b5050509392505050565b6000614fa8614fa384615b0f565b615a92565b905082815260208101848484011115614fc057600080fd5b614fcb848285615d88565b509392505050565b600081359050614fe281616003565b92915050565b600081519050614ff781616003565b92915050565b600082601f83011261500e57600080fd5b815161501e848260208601614e89565b91505092915050565b600082601f83011261503857600080fd5b8151615048848260208601614f0f565b91505092915050565b6000815190506150608161601a565b92915050565b60008135905061507581616031565b92915050565b600082601f83011261508c57600080fd5b815161509c848260208601614f95565b91505092915050565b6000813590506150b481616048565b92915050565b6000813590506150c98161605f565b92915050565b6000815190506150de81616076565b92915050565b600060e082840312156150f657600080fd5b61510060e0615a92565b90506000615110848285016153e7565b6000830152506020615124848285016153e7565b6020830152506040615138848285016153e7565b604083015250606061514c848285016153e7565b6060830152506080615160848285016153e7565b60808301525060a061517484828501614fd3565b60a08301525060c061518884828501614fd3565b60c08301525092915050565b600061018082840312156151a757600080fd5b6151b2610180615a92565b905060006151c284828501614fe8565b60008301525060206151d6848285016153fc565b60208301525060406151ea848285016153fc565b60408301525060606151fe848285016153fc565b6060830152506080615212848285016150cf565b60808301525060a0615226848285016153d2565b60a08301525060c061523a848285016153d2565b60c08301525060e061524e848285016153fc565b60e083015250610100615263848285016153fc565b61010083015250610120615279848285016153fc565b6101208301525061014061528f84828501615051565b6101408301525061016082015167ffffffffffffffff8111156152b157600080fd5b6152bd84828501614ffd565b6101608301525092915050565b600061016082840312156152dd57600080fd5b6152e8610160615a92565b905060006152f8848285016153fc565b600083015250602061530c848285016153fc565b6020830152506040615320848285016153fc565b6040830152506060615334848285016153fc565b6060830152506080615348848285016153fc565b60808301525060a061535c848285016153fc565b60a08301525060c0615370848285016153fc565b60c08301525060e0615384848285016153fc565b60e083015250610100615399848285016153fc565b610100830152506101206153af848285016153fc565b610120830152506101406153c5848285016153fc565b6101408301525092915050565b6000815190506153e181616086565b92915050565b6000813590506153f68161609d565b92915050565b60008151905061540b8161609d565b92915050565b60006020828403121561542357600080fd5b600061543184828501614fd3565b91505092915050565b60006020828403121561544c57600080fd5b600082015167ffffffffffffffff81111561546657600080fd5b61547284828501615027565b91505092915050565b60006020828403121561548d57600080fd5b600061549b84828501615066565b91505092915050565b600080604083850312156154b757600080fd5b60006154c5858286016150a5565b92505060206154d6858286016150ba565b9150509250929050565b600060e082840312156154f257600080fd5b6000615500848285016150e4565b91505092915050565b6000610160828403121561551c57600080fd5b600061552a848285016152ca565b91505092915050565b600061553f83836156c6565b60e08301905092915050565b61555481615cba565b82525050565b61556381615cba565b82525050565b600061557482615b50565b61557e8185615b68565b935061558983615b40565b8060005b838110156155ba5781516155a18882615533565b97506155ac83615b5b565b92505060018101905061558d565b5085935050505092915050565b6155d081615ccc565b82525050565b6155df81615d64565b82525050565b6155ee81615d19565b82525050565b6000615601601383615b79565b915061560c82615ed3565b602082019050919050565b6000615624602e83615b79565b915061562f82615efc565b604082019050919050565b6000615647601783615b79565b915061565282615f4b565b602082019050919050565b600061566a601183615b79565b915061567582615f74565b602082019050919050565b600061568d601383615b79565b915061569882615f9d565b602082019050919050565b60006156b0601783615b79565b91506156bb82615fc6565b602082019050919050565b60e0820160008201516156dc600085018261582f565b5060208201516156ef602085018261582f565b506040820151615702604085018261582f565b506060820151615715606085018261582f565b506080820151615728608085018261582f565b5060a082015161573b60a085018261554b565b5060c082015161574e60c085018261554b565b50505050565b60e08201600082015161576a600085018261582f565b50602082015161577d602085018261582f565b506040820151615790604085018261582f565b5060608201516157a3606085018261582f565b5060808201516157b6608085018261582f565b5060a08201516157c960a085018261554b565b5060c08201516157dc60c085018261554b565b50505050565b6040820160008201516157f8600085018261554b565b50602082015161580b602085018261582f565b50505050565b61581a81615d46565b82525050565b61582981615d76565b82525050565b61583881615d50565b82525050565b61584781615d50565b82525050565b6000602082019050615862600083018461555a565b92915050565b600060208201905081810360008301526158828184615569565b905092915050565b600060208201905061589f60008301846155c7565b92915050565b60006060820190506158ba60008301866155d6565b6158c76020830185615811565b6158d4604083018461555a565b949350505050565b600060c0820190506158f160008301896155d6565b6158fe6020830188615811565b61590b604083018761555a565b615918606083018661555a565b615925608083018561583e565b61593260a083018461583e565b979650505050505050565b600060208201905061595260008301846155e5565b92915050565b60006020820190508181036000830152615971816155f4565b9050919050565b6000602082019050818103600083015261599181615617565b9050919050565b600060208201905081810360008301526159b18161563a565b9050919050565b600060208201905081810360008301526159d18161565d565b9050919050565b600060208201905081810360008301526159f181615680565b9050919050565b60006020820190508181036000830152615a11816156a3565b9050919050565b600060e082019050615a2d6000830184615754565b92915050565b6000604082019050615a4860008301846157e2565b92915050565b6000604082019050615a636000830185615811565b615a706020830184615820565b9392505050565b6000602082019050615a8c600083018461583e565b92915050565b6000615a9c615aad565b9050615aa88282615dbb565b919050565b6000604051905090565b600067ffffffffffffffff821115615ad257615ad1615e93565b5b602082029050602081019050919050565b600067ffffffffffffffff821115615afe57615afd615e93565b5b602082029050602081019050919050565b600067ffffffffffffffff821115615b2a57615b29615e93565b5b615b3382615ec2565b9050602081019050919050565b6000819050602082019050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b6000615b9582615d50565b9150615ba083615d50565b92508267ffffffffffffffff03821115615bbd57615bbc615e35565b5b828201905092915050565b6000615bd382615d50565b9150615bde83615d50565b92508167ffffffffffffffff0483118215151615615bff57615bfe615e35565b5b828202905092915050565b6000615c1582615d19565b9150615c2083615d19565b9250827fffffffffffffffffffffffffffffffffffffffffffffffff800000000000000001821260008412151615615c5b57615c5a615e35565b5b82677fffffffffffffff018213600084121615615c7b57615c7a615e35565b5b828203905092915050565b6000615c9182615d50565b9150615c9c83615d50565b925082821015615caf57615cae615e35565b5b828203905092915050565b6000615cc582615d26565b9050919050565b60008115159050919050565b6000819050919050565b6000615ced82615cba565b9050919050565b6000615cff82615cba565b9050919050565b6000819050615d1482615fef565b919050565b60008160070b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b6000615d6f82615d06565b9050919050565b6000615d8182615d50565b9050919050565b60005b83811015615da6578082015181840152602081019050615d8b565b83811115615db5576000848401525b50505050565b615dc482615ec2565b810181811067ffffffffffffffff82111715615de357615de2615e93565b5b80604052505050565b6000615df782615d46565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415615e2a57615e29615e35565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f4e6f6465206e6f74207265676973746572656400000000000000000000000000600082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f4e6f646520616c72656164792072656769737465726564000000000000000000600082015250565b7f4e6f7420656e6f75676820706c65646765000000000000000000000000000000600082015250565b7f566f6c756d6520697320746f6f20736d616c6c00000000000000000000000000600082015250565b7f4e6f64652077616c6c657441646472206368616e676564000000000000000000600082015250565b600a811061600057615fff615e64565b5b50565b61600c81615cba565b811461601757600080fd5b50565b61602381615ccc565b811461602e57600080fd5b50565b61603a81615cd8565b811461604557600080fd5b50565b61605181615ce2565b811461605c57600080fd5b50565b61606881615cf4565b811461607357600080fd5b50565b6003811061608357600080fd5b50565b61608f81615d46565b811461609a57600080fd5b50565b6160a681615d50565b81146160b157600080fd5b5056fea264697066735822122043a2bcb294a99a3d79aea2df29cc7fbc6cfd7577a0b1dd8e33380108f52042e064736f6c63430008040033",
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

// GetPledgeUpdate is a free data retrieval call binding the contract method 0xa0217c95.
//
// Solidity: function GetPledgeUpdate((uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo) view returns(int64)
func (_Store *StoreCaller) GetPledgeUpdate(opts *bind.CallOpts, nodeInfo NodeInfo) (int64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetPledgeUpdate", nodeInfo)

	if err != nil {
		return *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(int64)).(*int64)

	return out0, err

}

// GetPledgeUpdate is a free data retrieval call binding the contract method 0xa0217c95.
//
// Solidity: function GetPledgeUpdate((uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo) view returns(int64)
func (_Store *StoreSession) GetPledgeUpdate(nodeInfo NodeInfo) (int64, error) {
	return _Store.Contract.GetPledgeUpdate(&_Store.CallOpts, nodeInfo)
}

// GetPledgeUpdate is a free data retrieval call binding the contract method 0xa0217c95.
//
// Solidity: function GetPledgeUpdate((uint64,uint64,uint64,uint64,uint64,address,address) nodeInfo) view returns(int64)
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
