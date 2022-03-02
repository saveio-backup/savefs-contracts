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

// ListWhiteListParams is an auto generated low-level Go binding around an user-defined struct.
type ListWhiteListParams struct {
	FileHash []byte
	Op       uint8
	List     []WhiteList
}

// WhiteList is an auto generated low-level Go binding around an user-defined struct.
type WhiteList struct {
	Addr         common.Address
	BaseHeight   uint64
	ExpireHeight uint64
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"GetWhiteList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"enumList.WhiteListOpType\",\"name\":\"Op\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"List\",\"type\":\"tuple[]\"}],\"internalType\":\"structList.WhiteListParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"WhiteListOperate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610d5b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806335deb19c146100465780635baf56971461005b5780638129fc1c14610084575b600080fd5b6100596100543660046109ac565b61008c565b005b61006e610069366004610977565b6104e8565b60405161007b9190610b3a565b60405180910390f35b610059610600565b6000816020015160048111156100b257634e487b7160e01b600052602160045260246000fd5b14156101a2576000600182600001516040516100ce9190610b27565b9081526020016040518091039020905060005b82604001515181101561019f57818360400151828151811061011357634e487b7160e01b600052603260045260246000fd5b60209081029190910181015182546001808201855560009485529383902082516002909202018054938301516001600160a01b039092166001600160e01b031990941693909317600160a01b67ffffffffffffffff9283160217835560409091015191909201805467ffffffffffffffff1916919092161790558061019781610cad565b9150506100e1565b50505b6001816020015160048111156101c857634e487b7160e01b600052602160045260246000fd5b1415610306576000600182600001516040516101e49190610b27565b9081526020016040518091039020905060005b8260400151518110156103035760005b82548110156102f0578360400151828151811061023457634e487b7160e01b600052603260045260246000fd5b6020026020010151600001516001600160a01b031683828154811061026957634e487b7160e01b600052603260045260246000fd5b60009182526020909120600290910201546001600160a01b031614156102de578281815481106102a957634e487b7160e01b600052603260045260246000fd5b6000918252602090912060029091020180546001600160e01b0319168155600101805467ffffffffffffffff191690556102f0565b806102e881610cad565b915050610207565b50806102fb81610cad565b9150506101f7565b50505b60028160200151600481111561032c57634e487b7160e01b600052602160045260246000fd5b141561048e576000600182600001516040516103489190610b27565b9081526020016040518091039020905060005b815481101561045a5781818154811061038457634e487b7160e01b600052603260045260246000fd5b906000526020600020906002020160000160149054906101000a900467ffffffffffffffff1667ffffffffffffffff168282815481106103d457634e487b7160e01b600052603260045260246000fd5b600091825260209091206001600290920201015467ffffffffffffffff16116104485781818154811061041757634e487b7160e01b600052603260045260246000fd5b6000918252602090912060029091020180546001600160e01b0319168155600101805467ffffffffffffffff191690555b8061045281610cad565b91505061035b565b5080600183600001516040516104709190610b27565b908152604051908190036020019020815461048b9290610671565b50505b6003816020015160048111156104b457634e487b7160e01b600052602160045260246000fd5b14156104e55780516040516001916104cb91610b27565b908152602001604051809103902060006104e59190610742565b50565b606060008251116105145760405162461bcd60e51b815260040161050b90610b5b565b60405180910390fd5b60006001836040516105269190610b27565b90815260405190819003602001902054116105535760405162461bcd60e51b815260040161050b90610b4b565b6001826040516105639190610b27565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156105f5576000848152602090819020604080516060810182526002860290920180546001600160a01b038116845267ffffffffffffffff600160a01b90910481168486015260019182015416918301919091529083529092019101610591565b505050509050919050565b600054610100900460ff1661061b5760005460ff161561061f565b303b155b61063b5760405162461bcd60e51b815260040161050b90610b6b565b600054610100900460ff1615801561065d576000805461ffff19166101011790555b80156104e5576000805461ff001916905550565b8280548282559060005260206000209060020281019282156107325760005260206000209160020282015b8281111561073257825482547fffffffffffffffffffffffff000000000000000000000000000000000000000081166001600160a01b039092169182178455845467ffffffffffffffff600160a01b9182900481169091026001600160e01b0319909216909217178355600180850154908401805467ffffffffffffffff1916919092161790556002928301929091019061069c565b5061073e92915061075f565b5090565b50805460008255600202906000526020600020908101906104e591905b5b8082111561073e5780546001600160e01b031916815560018101805467ffffffffffffffff19169055600201610760565b60006107a461079f84610be8565b610bcc565b838152905060208101826060850281018610156107c057600080fd5b60005b858110156107ee57816107d68882610921565b845250602090920191606091909101906001016107c3565b5050509392505050565b600061080661079f84610c0c565b90508281526020810184848401111561081e57600080fd5b610829848285610c48565b509392505050565b803561083c81610cf4565b92915050565b600082601f83011261085357600080fd5b8135610863848260208601610791565b949350505050565b600082601f83011261087c57600080fd5b81356108638482602086016107f8565b803561083c81610d08565b6000606082840312156108a957600080fd5b6108b36060610bcc565b9050813567ffffffffffffffff8111156108cc57600080fd5b6108d88482850161086b565b82525060206108e98484830161088c565b602083015250604082013567ffffffffffffffff81111561090957600080fd5b61091584828501610842565b60408301525092915050565b60006060828403121561093357600080fd5b61093d6060610bcc565b9050600061094b8484610831565b825250602061095c8484830161096c565b6020830152506040610915848285015b803561083c81610d15565b60006020828403121561098957600080fd5b813567ffffffffffffffff8111156109a057600080fd5b6108638482850161086b565b6000602082840312156109be57600080fd5b813567ffffffffffffffff8111156109d557600080fd5b61086384828501610897565b60006109ed8383610ada565b505060600190565b6109fe81610c37565b82525050565b6000610a0e825190565b80845260209384019383018060005b83811015610a42578151610a3188826109e1565b975060208301925050600101610a1d565b509495945050505050565b6000610a57825190565b610a65818560208601610c54565b9290920192915050565b601281526000602082017f77686974654c69737420697320656d7074790000000000000000000000000000815291505b5060200190565b601681526000602082017f66696c6548617368206d75737420626520656d7074790000000000000000000081529150610a9f565b80516060830190610aeb84826109f5565b506020820151610afe6020850182610b17565b506040820151610b116040850182610b17565b50505050565b67ffffffffffffffff81166109fe565b6000610b338284610a4d565b9392505050565b60208082528101610b338184610a04565b6020808252810161083c81610a6f565b6020808252810161083c81610aa6565b6020808252810161083c81602e81527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160208201527f647920696e697469616c697a6564000000000000000000000000000000000000604082015260600190565b6000610bd760405190565b9050610be38282610c80565b919050565b600067ffffffffffffffff821115610c0257610c02610cde565b5060209081020190565b600067ffffffffffffffff821115610c2657610c26610cde565b601f19601f83011660200192915050565b60006001600160a01b03821661083c565b82818337506000910152565b60005b83811015610c6f578181015183820152602001610c57565b83811115610b115750506000910152565b601f19601f830116810181811067ffffffffffffffff82111715610ca657610ca6610cde565b6040525050565b6000600019821415610cc157610cc1610cc8565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b610cfd81610c37565b81146104e557600080fd5b600581106104e557600080fd5b67ffffffffffffffff8116610cfd56fea264697066735822122078d2b04809c90c60e34ece6ce1ea7e7aa4f4ddf872f9bcd01507419f81d233bc64736f6c63430008040033",
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

// GetWhiteList is a free data retrieval call binding the contract method 0x5baf5697.
//
// Solidity: function GetWhiteList(bytes fileHash) view returns((address,uint64,uint64)[])
func (_Store *StoreCaller) GetWhiteList(opts *bind.CallOpts, fileHash []byte) ([]WhiteList, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetWhiteList", fileHash)

	if err != nil {
		return *new([]WhiteList), err
	}

	out0 := *abi.ConvertType(out[0], new([]WhiteList)).(*[]WhiteList)

	return out0, err

}

// GetWhiteList is a free data retrieval call binding the contract method 0x5baf5697.
//
// Solidity: function GetWhiteList(bytes fileHash) view returns((address,uint64,uint64)[])
func (_Store *StoreSession) GetWhiteList(fileHash []byte) ([]WhiteList, error) {
	return _Store.Contract.GetWhiteList(&_Store.CallOpts, fileHash)
}

// GetWhiteList is a free data retrieval call binding the contract method 0x5baf5697.
//
// Solidity: function GetWhiteList(bytes fileHash) view returns((address,uint64,uint64)[])
func (_Store *StoreCallerSession) GetWhiteList(fileHash []byte) ([]WhiteList, error) {
	return _Store.Contract.GetWhiteList(&_Store.CallOpts, fileHash)
}

// WhiteListOperate is a paid mutator transaction binding the contract method 0x35deb19c.
//
// Solidity: function WhiteListOperate((bytes,uint8,(address,uint64,uint64)[]) params) returns()
func (_Store *StoreTransactor) WhiteListOperate(opts *bind.TransactOpts, params ListWhiteListParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "WhiteListOperate", params)
}

// WhiteListOperate is a paid mutator transaction binding the contract method 0x35deb19c.
//
// Solidity: function WhiteListOperate((bytes,uint8,(address,uint64,uint64)[]) params) returns()
func (_Store *StoreSession) WhiteListOperate(params ListWhiteListParams) (*types.Transaction, error) {
	return _Store.Contract.WhiteListOperate(&_Store.TransactOpts, params)
}

// WhiteListOperate is a paid mutator transaction binding the contract method 0x35deb19c.
//
// Solidity: function WhiteListOperate((bytes,uint8,(address,uint64,uint64)[]) params) returns()
func (_Store *StoreTransactorSession) WhiteListOperate(params ListWhiteListParams) (*types.Transaction, error) {
	return _Store.Contract.WhiteListOperate(&_Store.TransactOpts, params)
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
