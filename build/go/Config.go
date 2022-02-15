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
	ABI: "[{\"inputs\":[{\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"}],\"name\":\"GetProveIntervalByProveLevel\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetSetting\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumProveLevel\",\"name\":\"proveLevel\",\"type\":\"uint8\"}],\"name\":\"GetSettingWithProveLevel\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"GasPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerGBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasPerKBPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GasForChallenge\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MaxProveBlockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinVolume\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProvePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultProveLevel\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultCopyNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DefaultBlockInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MinSectorSize\",\"type\":\"uint64\"}],\"internalType\":\"structSetting\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610619806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806322cf12cf14610046578063a86593981461011a578063c5c81b201461013a575b600080fd5b604080516101608082018352600080835260208084018290528385018290526060808501839052608080860184905260a080870185905260c080880186905260e0808901879052610100808a01889052610120808b018990526101409a8b01989098528a519889018b526001808a528988018190529a89018b905262030d409589019590955292870194909452620f424090860181905261438093860193909352840195909552600294830194909452600590820152908101919091525b6040516101119190610559565b60405180910390f35b61012d6101283660046103fa565b61014d565b6040516101119190610568565b6101046101483660046103fa565b6102c8565b600080610213604080516101608082018352600080835260208084018290528385018290526060808501839052608080860184905260a080870185905260c080880186905260e0808901879052610100808a01889052610120808b018990526101409a8b01989098528a519889018b526001808a528988018190529a89018b905262030d409589019590955292870194909452620f4240908601819052614380938601939093528401959095526002948301949094526005908201529081019190915290565b9050600083600281111561023757634e487b7160e01b600052602160045260246000fd5b14156102475760c0015192915050565b600183600281111561026957634e487b7160e01b600052602160045260246000fd5b14156102865760c081015161027f90600261057e565b9392505050565b60028360028111156102a857634e487b7160e01b600052602160045260246000fd5b14156102be5760c081015161027f90600861057e565b60c0015192915050565b604080516101608082018352600080835260208084018290528385018290526060808501839052608080860184905260a080870185905260c080880186905260e0808901879052610100808a01889052610120808b01899052610140808c018a90528c51808c018e528a8152808a018b9052808e018b90528089018b90528088018b90528087018b90528086018b90528085018b90528084018b90528083018b90528101999099528b51998a018c526001808b528a89018190529b8a018c905262030d40968a019690965293880195909552620f42409187018290526143809087015292850196909652600295840195909552600594830194909452810192909252906103d48361014d565b67ffffffffffffffff1660c082015292915050565b80356103f4816105d3565b92915050565b60006020828403121561040c57600080fd5b600061041884846103e9565b949350505050565b805167ffffffffffffffff81168352610160830190506020820151610451602085018267ffffffffffffffff169052565b50604082015161046d604085018267ffffffffffffffff169052565b506060820151610489606085018267ffffffffffffffff169052565b5060808201516104a5608085018267ffffffffffffffff169052565b5060a08201516104c160a085018267ffffffffffffffff169052565b5060c08201516104dd60c085018267ffffffffffffffff169052565b5060e08201516104f960e085018267ffffffffffffffff169052565b5061010082015161051761010085018267ffffffffffffffff169052565b5061012082015161053561012085018267ffffffffffffffff169052565b5061014082015161055361014085018267ffffffffffffffff169052565b50505050565b61016081016103f48284610420565b67ffffffffffffffff82168152602081016103f4565b600067ffffffffffffffff8216915067ffffffffffffffff831692508167ffffffffffffffff04831182151516156105b8576105b86105bd565b500290565b634e487b7160e01b600052601160045260246000fd5b600381106105e057600080fd5b5056fea2646970667358221220c50584ec88f0f2faeb77ceebe840f34f86f5539a37d8ef739822da78866e21ba64736f6c63430008040033",
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

// GetProveIntervalByProveLevel is a free data retrieval call binding the contract method 0xa8659398.
//
// Solidity: function GetProveIntervalByProveLevel(uint8 proveLevel) pure returns(uint64)
func (_Store *StoreCaller) GetProveIntervalByProveLevel(opts *bind.CallOpts, proveLevel uint8) (uint64, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetProveIntervalByProveLevel", proveLevel)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetProveIntervalByProveLevel is a free data retrieval call binding the contract method 0xa8659398.
//
// Solidity: function GetProveIntervalByProveLevel(uint8 proveLevel) pure returns(uint64)
func (_Store *StoreSession) GetProveIntervalByProveLevel(proveLevel uint8) (uint64, error) {
	return _Store.Contract.GetProveIntervalByProveLevel(&_Store.CallOpts, proveLevel)
}

// GetProveIntervalByProveLevel is a free data retrieval call binding the contract method 0xa8659398.
//
// Solidity: function GetProveIntervalByProveLevel(uint8 proveLevel) pure returns(uint64)
func (_Store *StoreCallerSession) GetProveIntervalByProveLevel(proveLevel uint8) (uint64, error) {
	return _Store.Contract.GetProveIntervalByProveLevel(&_Store.CallOpts, proveLevel)
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
