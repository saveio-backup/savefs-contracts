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

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"GetSetting\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"}],\"name\":\"GetSettingWithProveLevel\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0xea7eb5a7\",\"type\":\"bytes32\"}],\"name\":\"c_0xea7eb5a7\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061114e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806322cf12cf146100465780636c5ab29f14610064578063c5c81b2014610080575b600080fd5b61004e6100b0565b60405161005b9190611046565b60405180910390f35b61007e60048036038101906100799190610f04565b6106e0565b005b61009a60048036038101906100959190610f2d565b6106e3565b6040516100a79190611046565b60405180910390f35b6100b8610e12565b6100e47f25d761c9f975c5dcebb1527bd6c91c83fc5e16d5a604c09f21248986d83e33b160001b6106e0565b6101107fa80353c277b53781550e8004b1183c216bd93b917d47010877fff9bb52de772e60001b6106e0565b61013c7ff05d7d6f1ba86c204ce69188f35d2c944728d8559f9f05bc3b0b77400ebe5a4d60001b6106e0565b610144610e12565b6101707f37e84a7390b556cf030ee8c20b97ed7d71e8f6135be5db59c68f3562f30873ef60001b6106e0565b61019c7fa8c6e9662e351b8c75cce8fb32853a5499a472a41d2029fa271be02ff1bd8ee560001b6106e0565b6001816000019067ffffffffffffffff16908167ffffffffffffffff16815250506101e97fef666733d531e627a1919556160cc7c161761406a77b5c3de0bcb1da574aa76a60001b6106e0565b6102157f12c87d395ab915dd9042a1819ceb3ea5f0ffcf4236420920b69055c3b235c83a60001b6106e0565b6001816020019067ffffffffffffffff16908167ffffffffffffffff16815250506102627f2ccf26ed9c534b634feab9391b63c8cbd26b236f99115300e0e2dfd19bebd9e960001b6106e0565b61028e7f3be63c117ebfa011c25170601a6c3bd43cc6a108fee6227571c64dcf891e95b760001b6106e0565b6001816040019067ffffffffffffffff16908167ffffffffffffffff16815250506102db7fcd3537bff4e09eb61f4824541bf653e4852ebedd49a7be17210ccf11a038443360001b6106e0565b6103077f53f421ebc62b0de495bb3db322ac1e998c5683a67d233f3e0858ff65191f6be360001b6106e0565b62030d40816060019067ffffffffffffffff16908167ffffffffffffffff16815250506103567f53573f5987f90cbd0cd6921943b89dd0c40ce907c97195c2ba01846f0529868c60001b6106e0565b6103827f53a2464869bc6d494e715d29a41c0f7fe3c19af2addff420ca32b4b878c7a2f660001b6106e0565b6020816080019067ffffffffffffffff16908167ffffffffffffffff16815250506103cf7f6a872a919f2f332a8c486e49d9f9fb203c5a0e38369e74901b7198e75bb8c76060001b6106e0565b6103fb7f5e00c08b59e97462d174caf1e2587d125dd222a4dd00ce8ee1cd68b466af6de360001b6106e0565b620f42408160a0019067ffffffffffffffff16908167ffffffffffffffff168152505061044a7f796c17eacfab6188ca7ed4319a6c13e2c30a7021f434110cf62565c3fb5c4ab860001b6106e0565b6104767f74b5e848b29d1b6c3abc26050b14af5bf073ea45e7327245d59670213551b6df60001b6106e0565b6143808160c0019067ffffffffffffffff16908167ffffffffffffffff16815250506104c47fc9aa0981ac513ffeb618117c2e000e79b436285fd3a3079dfd6fe6354c3d8cd460001b6106e0565b6104f07f9f97aa5cb3acc197a8171a9d7bd2edce17b19830724d2d416c6d5f98f0ee0a0360001b6106e0565b60018160e0019067ffffffffffffffff16908167ffffffffffffffff168152505061053d7f2847aa11c20d9c577abf2394c8ddbb95012a12166c8ef3ccc9b3eb1a43e45ff460001b6106e0565b6105697f5e4bb4bd07dad6c14da1c76335806a1432024f74477b8530e1b32eb3a1b87be460001b6106e0565b600281610100019067ffffffffffffffff16908167ffffffffffffffff16815250506105b77f4c3c41793b9fa035a29dfc6e70035d9c21f267ca4e820589cf4da3f606b5696960001b6106e0565b6105e37fa39e8a7fdf63f8371cc2bfaa7d3d1b0e5cefd9af8a31902b5e8bbacae06e42de60001b6106e0565b600581610120019067ffffffffffffffff16908167ffffffffffffffff16815250506106317f3850faac504ab2a8611b052165c03b054a970d9a60bce8d540ad78850286b79460001b6106e0565b61065d7ff0f748406f19a1a4d98e2e36d9d2a905b4016f6935faa847ceb6a9c66d6b2bce60001b6106e0565b620f424081610140019067ffffffffffffffff16908167ffffffffffffffff16815250506106ad7fe26160b8f051f556343f31e1d250b10b0219c9124c75a16460e67b03508eb68360001b6106e0565b6106d97fc678ba4c35663e79443fb1e214039601fc842fe5b35dba01f567fca6ccd48de060001b6106e0565b8091505090565b50565b6106eb610e12565b6107177f2f051ee7d764748f04c000da93c0175a11df665135278c88d5efe5c82da83b8560001b6106e0565b6107437fa6797a0969ee35120d3ede94ab84b3b912d59725401c954655cf6d7dd3355b0f60001b6106e0565b61076f7fb503a6c37b579b8e8005cc7e699fcb9c4efc9ade251c9f481a1d7afd329776d560001b6106e0565b60006107796100b0565b90506107a77f456352522c85cc1113bb1c513c70af80c0bc8807d1608f523b34a60817a7b40360001b6106e0565b6107d37ff7168de228a32c4aa5217eb424bf3150d427c6fc6fff194ba2a163154511133f60001b6106e0565b6107dc8361085c565b8160c0019067ffffffffffffffff16908167ffffffffffffffff16815250506108277f328190f7d2335fb4817fe4f0a4385abb59145a2cd556d97c613d1bd50afd64f960001b6106e0565b6108537f221dcf64f94dfa75f5e3387343b72e3b183c42c3f409826888d316b7e487369560001b6106e0565b80915050919050565b600061088a7f614df1bd4327fb0c869ecf3e236a4ee2649289e5c02064dd7411186a1125d24960001b6106e0565b6108b67f26e89168660b6ccdf25691fcc110d08a39bdf4eceedc969d0c6c5f45594340c760001b6106e0565b6108e27f64e445cc373687869264b5e180eb3a6ec2e4f2a7060469ed5e000c019da4772660001b6106e0565b60006108ec6100b0565b905061091a7f9b03e1c9338d4c5e189fb17c56fc1072f0d300adbdba403f8166a52c55e05f5e60001b6106e0565b6109467f178b4b8cbbfcac236b7eee87a2db07e39a756a91e7877670da062377915855bf60001b6106e0565b60006002811115610980577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8360028111156109b9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415610a50576109eb7f168f6a8da2e4e7364cd8d88a7a58fde70ca9a0d58bf8e45baa6e1ea0d4c3d5c260001b6106e0565b610a177f5cda96c5cd21dceadeaa8c83bbf06872914dfd160e61114af4e70090f950eb4560001b6106e0565b610a437f9a39b5ee9852edb6c54d46a8c0df4a9629702b9e9296d483ea1a433dc0ddab1960001b6106e0565b8060c00151915050610e0d565b610a7c7fa5f7c1d5d50cb03d7acfb5f1d3aeaab17764001ed69bcbec24ea66205691ecea60001b6106e0565b610aa87f4bad6ed4794dbcfe94cea0a6ae415e2f5b0957b7c07a16d8d96ebf035137153f60001b6106e0565b610ad47f96a16c4b58bd18911a06835b7d96e5a2b134a9205492a09b2e84888405376c2e60001b6106e0565b60016002811115610b0e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b836002811115610b47577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415610bea57610b797f5628b65c4a34916ce8c367a693953bdaf86fd43928fe227089b849bdbc5847f560001b6106e0565b610ba57f31b9fecfe2d8f6c4e4b5217112397513d85a47e21f120271de45b01b9c26006460001b6106e0565b610bd17f40884f6f1ebb0d843b2ea7abcac5247e050bdf64783ff787aea3f8f8d428ebcf60001b6106e0565b60028160c00151610be29190611062565b915050610e0d565b610c167fd55b0d6139f352e34894955aba3c7926281def55b9550ffa93c042321832a22a60001b6106e0565b610c427f4e4ee34ed87474c85792ae3604b75094c383eec58cfa1150bf6a803e1663049b60001b6106e0565b610c6e7ff4b9c6d6804f4e52e7dae60a45ecd834c597a063203dc734f6930ec7758a094c60001b6106e0565b600280811115610ca7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b836002811115610ce0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415610d8357610d127f6d7c9451c27f2ebbdcddf13a1e0dbe7a84ce0ee333768f0b90d39acf4cdb864e60001b6106e0565b610d3e7f9d6340a64a50ca1ea8e056197b630013f07ef3d5a169e853f0b687d728ba18d160001b6106e0565b610d6a7fe666bfb1f53aa51da17df7e6ed561101132f79f5e588ffc7c49339ba62d3292960001b6106e0565b60088160c00151610d7b9190611062565b915050610e0d565b610daf7f25b8ff10f37ff975574e3f4de5220dda500ecce328f728d26ffc429af6187cf860001b6106e0565b610ddb7fc4f443a7e180c0c4f030453d2101c13f39034589679f5e7f09707b75029c322560001b6106e0565b610e077f031e36f0a055e7aba370ef121b5eba7a68ed4336247eb347cfbcc87c3f67b72a60001b6106e0565b60009150505b919050565b604051806101600160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b600081359050610ee9816110f1565b92915050565b600081359050610efe81611108565b92915050565b600060208284031215610f1657600080fd5b6000610f2484828501610eda565b91505092915050565b600060208284031215610f3f57600080fd5b6000610f4d84828501610eef565b91505092915050565b61016082016000820151610f6d6000850182611037565b506020820151610f806020850182611037565b506040820151610f936040850182611037565b506060820151610fa66060850182611037565b506080820151610fb96080850182611037565b5060a0820151610fcc60a0850182611037565b5060c0820151610fdf60c0850182611037565b5060e0820151610ff260e0850182611037565b50610100820151611007610100850182611037565b5061012082015161101c610120850182611037565b50610140820151611031610140850182611037565b50505050565b611040816110ae565b82525050565b60006101608201905061105c6000830184610f56565b92915050565b600061106d826110ae565b9150611078836110ae565b92508167ffffffffffffffff0483118215151615611099576110986110c2565b5b828202905092915050565b6000819050919050565b600067ffffffffffffffff82169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6110fa816110a4565b811461110557600080fd5b50565b6003811061111557600080fd5b5056fea2646970667358221220a39746d1f61573884bf00800cd02c0aac940ab4d44bbd2163a85805bcfbe232364736f6c63430008040033",
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

// GetSetting is a free data retrieval call binding the contract method 0x22cf12cf.
//
// Solidity: function GetSetting() pure returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
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
// Solidity: function GetSetting() pure returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_Store *StoreSession) GetSetting() (Setting, error) {
	return _Store.Contract.GetSetting(&_Store.CallOpts)
}

// GetSetting is a free data retrieval call binding the contract method 0x22cf12cf.
//
// Solidity: function GetSetting() pure returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_Store *StoreCallerSession) GetSetting() (Setting, error) {
	return _Store.Contract.GetSetting(&_Store.CallOpts)
}

// GetSettingWithProveLevel is a free data retrieval call binding the contract method 0xc5c81b20.
//
// Solidity: function GetSettingWithProveLevel(uint8 proveLevel) pure returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
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
// Solidity: function GetSettingWithProveLevel(uint8 proveLevel) pure returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_Store *StoreSession) GetSettingWithProveLevel(proveLevel uint8) (Setting, error) {
	return _Store.Contract.GetSettingWithProveLevel(&_Store.CallOpts, proveLevel)
}

// GetSettingWithProveLevel is a free data retrieval call binding the contract method 0xc5c81b20.
//
// Solidity: function GetSettingWithProveLevel(uint8 proveLevel) pure returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_Store *StoreCallerSession) GetSettingWithProveLevel(proveLevel uint8) (Setting, error) {
	return _Store.Contract.GetSettingWithProveLevel(&_Store.CallOpts, proveLevel)
}

// C0xea7eb5a7 is a free data retrieval call binding the contract method 0x6c5ab29f.
//
// Solidity: function c_0xea7eb5a7(bytes32 c__0xea7eb5a7) pure returns()
func (_Store *StoreCaller) C0xea7eb5a7(opts *bind.CallOpts, c__0xea7eb5a7 [32]byte) error {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "c_0xea7eb5a7", c__0xea7eb5a7)

	if err != nil {
		return err
	}

	return err

}

// C0xea7eb5a7 is a free data retrieval call binding the contract method 0x6c5ab29f.
//
// Solidity: function c_0xea7eb5a7(bytes32 c__0xea7eb5a7) pure returns()
func (_Store *StoreSession) C0xea7eb5a7(c__0xea7eb5a7 [32]byte) error {
	return _Store.Contract.C0xea7eb5a7(&_Store.CallOpts, c__0xea7eb5a7)
}

// C0xea7eb5a7 is a free data retrieval call binding the contract method 0x6c5ab29f.
//
// Solidity: function c_0xea7eb5a7(bytes32 c__0xea7eb5a7) pure returns()
func (_Store *StoreCallerSession) C0xea7eb5a7(c__0xea7eb5a7 [32]byte) error {
	return _Store.Contract.C0xea7eb5a7(&_Store.CallOpts, c__0xea7eb5a7)
}
