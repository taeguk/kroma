// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// KromaTokenMinterMetaData contains all meta data concerning the KromaTokenMinter contract.
var KromaTokenMinterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_shares\",\"type\":\"uint64[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEPOSITOR_ACCOUNT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_DECREASE_CAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_INCREASE_CAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_MAX_INCREASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_MIN_DECREASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHARE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_shares\",\"type\":\"uint64[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kromaToken\",\"outputs\":[{\"internalType\":\"contractKromaToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recipients\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"shareOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620011d7380380620011d78339810160408190526200003491620004a9565b73420000000000000000000000000000000000001060805262000058828262000060565b5050620005ea565b600054610100900460ff1615808015620000815750600054600160ff909116105b80620000b157506200009e30620003a160201b620008331760201c565b158015620000b1575060005460ff166001145b6200011a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff1916600117905580156200013e576000805461ff0019166101001790555b8151835114620001915760405162461bcd60e51b815260206004820152601760248201527f696e76616c6964206c656e677468206f66206172726179000000000000000000604482015260640162000111565b6000805b845181101562000310576000858281518110620001b657620001b662000587565b6020026020010151905060006001600160a01b0316816001600160a01b031603620002245760405162461bcd60e51b815260206004820152601d60248201527f726563697069656e7420616464726573732063616e6e6f742062652030000000604482015260640162000111565b60008583815181106200023b576200023b62000587565b60200260200101516001600160401b0316905080600003620002945760405162461bcd60e51b8152602060048201526011602482015270073686172652063616e6e6f74206265203607c1b604482015260640162000111565b620002a08185620005b3565b600180548082019091557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b039094166001600160a01b031990941684179055600092835260026020526040909220559150806200030781620005ce565b91505062000195565b5060648114620003545760405162461bcd60e51b815260206004820152600e60248201526d696e76616c69642073686172657360901b604482015260640162000111565b5080156200039c576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6001600160a01b03163b151590565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620003f157620003f1620003b0565b604052919050565b60006001600160401b03821115620004155762000415620003b0565b5060051b60200190565b600082601f8301126200043157600080fd5b815160206200044a6200044483620003f9565b620003c6565b82815260059290921b840181019181810190868411156200046a57600080fd5b8286015b848110156200049e5780516001600160401b0381168114620004905760008081fd5b83529183019183016200046e565b509695505050505050565b60008060408385031215620004bd57600080fd5b82516001600160401b0380821115620004d557600080fd5b818501915085601f830112620004ea57600080fd5b81516020620004fd6200044483620003f9565b82815260059290921b840181019181810190898411156200051d57600080fd5b948201945b83861015620005545785516001600160a01b0381168114620005445760008081fd5b8252948201949082019062000522565b918801519196509093505050808211156200056e57600080fd5b506200057d858286016200041f565b9150509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60008219821115620005c957620005c96200059d565b500190565b600060018201620005e357620005e36200059d565b5060010190565b608051610bca6200060d6000396000818161013e01526103320152610bca6000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c8063c27ca14911610081578063e06939a31161005b578063e06939a314610198578063e591b282146101a0578063fdb6583e146101bb57600080fd5b8063c27ca149146100d9578063c9d2b49614610139578063d1bc76a11461018557600080fd5b80636745d032116100b25780636745d0321461011c5780637eb1184514610114578063a0712d681461012457600080fd5b80630ccfab45146100d957806321e5e2c4146100f45780632c7dc24214610114575b600080fd5b6100e1600a81565b6040519081526020015b60405180910390f35b6100e1610102366004610878565b60026020526000908152604090205481565b6100e1606481565b6100e1600781565b61013761013236600461089a565b6101ce565b005b6101607f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100eb565b61016061019336600461089a565b6103aa565b6100e1600381565b61016073deaddeaddeaddeaddeaddeaddeaddeaddead000281565b6101376101c93660046109d7565b6103e1565b3373deaddeaddeaddeaddeaddeaddeaddeaddead000214610276576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f6f6e6c79206465706f7369746f722063616e2063616c6c20746869732066756e60448201527f6374696f6e00000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b60005b6001548110156103a65760006001828154811061029857610298610a97565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16808352600290915260408220549092509060646102d98387610af5565b6102e39190610b32565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018390529192507f0000000000000000000000000000000000000000000000000000000000000000909116906340c10f1990604401600060405180830381600087803b15801561037857600080fd5b505af115801561038c573d6000803e3d6000fd5b50505050505050808061039e90610b6d565b915050610279565b5050565b600181815481106103ba57600080fd5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b600054610100900460ff16158080156104015750600054600160ff909116105b8061041b5750303b15801561041b575060005460ff166001145b6104a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161026d565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561050557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b8151835114610570576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f696e76616c6964206c656e677468206f66206172726179000000000000000000604482015260640161026d565b6000805b845181101561075f57600085828151811061059157610591610a97565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610631576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f726563697069656e7420616464726573732063616e6e6f742062652030000000604482015260640161026d565b600085838151811061064557610645610a97565b602002602001015167ffffffffffffffff169050806000036106c3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f73686172652063616e6e6f742062652030000000000000000000000000000000604482015260640161026d565b6106cd8185610ba5565b600180548082019091557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601805473ffffffffffffffffffffffffffffffffffffffff9094167fffffffffffffffffffffffff0000000000000000000000000000000000000000909416841790556000928352600260205260409092205591508061075781610b6d565b915050610574565b50606481146107ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f696e76616c696420736861726573000000000000000000000000000000000000604482015260640161026d565b50801561082e57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b803573ffffffffffffffffffffffffffffffffffffffff8116811461087357600080fd5b919050565b60006020828403121561088a57600080fd5b6108938261084f565b9392505050565b6000602082840312156108ac57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610929576109296108b3565b604052919050565b600067ffffffffffffffff82111561094b5761094b6108b3565b5060051b60200190565b600082601f83011261096657600080fd5b8135602061097b61097683610931565b6108e2565b82815260059290921b8401810191818101908684111561099a57600080fd5b8286015b848110156109cc57803567ffffffffffffffff811681146109bf5760008081fd5b835291830191830161099e565b509695505050505050565b600080604083850312156109ea57600080fd5b823567ffffffffffffffff80821115610a0257600080fd5b818501915085601f830112610a1657600080fd5b81356020610a2661097683610931565b82815260059290921b84018101918181019089841115610a4557600080fd5b948201945b83861015610a6a57610a5b8661084f565b82529482019490820190610a4a565b96505086013592505080821115610a8057600080fd5b50610a8d85828601610955565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610b2d57610b2d610ac6565b500290565b600082610b68577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610b9e57610b9e610ac6565b5060010190565b60008219821115610bb857610bb8610ac6565b50019056fea164736f6c634300080f000a",
}

// KromaTokenMinterABI is the input ABI used to generate the binding from.
// Deprecated: Use KromaTokenMinterMetaData.ABI instead.
var KromaTokenMinterABI = KromaTokenMinterMetaData.ABI

// KromaTokenMinterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KromaTokenMinterMetaData.Bin instead.
var KromaTokenMinterBin = KromaTokenMinterMetaData.Bin

// DeployKromaTokenMinter deploys a new Ethereum contract, binding an instance of KromaTokenMinter to it.
func DeployKromaTokenMinter(auth *bind.TransactOpts, backend bind.ContractBackend, _recipients []common.Address, _shares []uint64) (common.Address, *types.Transaction, *KromaTokenMinter, error) {
	parsed, err := KromaTokenMinterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KromaTokenMinterBin), backend, _recipients, _shares)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KromaTokenMinter{KromaTokenMinterCaller: KromaTokenMinterCaller{contract: contract}, KromaTokenMinterTransactor: KromaTokenMinterTransactor{contract: contract}, KromaTokenMinterFilterer: KromaTokenMinterFilterer{contract: contract}}, nil
}

// KromaTokenMinter is an auto generated Go binding around an Ethereum contract.
type KromaTokenMinter struct {
	KromaTokenMinterCaller     // Read-only binding to the contract
	KromaTokenMinterTransactor // Write-only binding to the contract
	KromaTokenMinterFilterer   // Log filterer for contract events
}

// KromaTokenMinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type KromaTokenMinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KromaTokenMinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KromaTokenMinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KromaTokenMinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KromaTokenMinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KromaTokenMinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KromaTokenMinterSession struct {
	Contract     *KromaTokenMinter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KromaTokenMinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KromaTokenMinterCallerSession struct {
	Contract *KromaTokenMinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// KromaTokenMinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KromaTokenMinterTransactorSession struct {
	Contract     *KromaTokenMinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// KromaTokenMinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type KromaTokenMinterRaw struct {
	Contract *KromaTokenMinter // Generic contract binding to access the raw methods on
}

// KromaTokenMinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KromaTokenMinterCallerRaw struct {
	Contract *KromaTokenMinterCaller // Generic read-only contract binding to access the raw methods on
}

// KromaTokenMinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KromaTokenMinterTransactorRaw struct {
	Contract *KromaTokenMinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKromaTokenMinter creates a new instance of KromaTokenMinter, bound to a specific deployed contract.
func NewKromaTokenMinter(address common.Address, backend bind.ContractBackend) (*KromaTokenMinter, error) {
	contract, err := bindKromaTokenMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KromaTokenMinter{KromaTokenMinterCaller: KromaTokenMinterCaller{contract: contract}, KromaTokenMinterTransactor: KromaTokenMinterTransactor{contract: contract}, KromaTokenMinterFilterer: KromaTokenMinterFilterer{contract: contract}}, nil
}

// NewKromaTokenMinterCaller creates a new read-only instance of KromaTokenMinter, bound to a specific deployed contract.
func NewKromaTokenMinterCaller(address common.Address, caller bind.ContractCaller) (*KromaTokenMinterCaller, error) {
	contract, err := bindKromaTokenMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KromaTokenMinterCaller{contract: contract}, nil
}

// NewKromaTokenMinterTransactor creates a new write-only instance of KromaTokenMinter, bound to a specific deployed contract.
func NewKromaTokenMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*KromaTokenMinterTransactor, error) {
	contract, err := bindKromaTokenMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KromaTokenMinterTransactor{contract: contract}, nil
}

// NewKromaTokenMinterFilterer creates a new log filterer instance of KromaTokenMinter, bound to a specific deployed contract.
func NewKromaTokenMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*KromaTokenMinterFilterer, error) {
	contract, err := bindKromaTokenMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KromaTokenMinterFilterer{contract: contract}, nil
}

// bindKromaTokenMinter binds a generic wrapper to an already deployed contract.
func bindKromaTokenMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KromaTokenMinterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KromaTokenMinter *KromaTokenMinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KromaTokenMinter.Contract.KromaTokenMinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KromaTokenMinter *KromaTokenMinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KromaTokenMinter.Contract.KromaTokenMinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KromaTokenMinter *KromaTokenMinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KromaTokenMinter.Contract.KromaTokenMinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KromaTokenMinter *KromaTokenMinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KromaTokenMinter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KromaTokenMinter *KromaTokenMinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KromaTokenMinter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KromaTokenMinter *KromaTokenMinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KromaTokenMinter.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() view returns(address)
func (_KromaTokenMinter *KromaTokenMinterCaller) DEPOSITORACCOUNT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KromaTokenMinter.contract.Call(opts, &out, "DEPOSITOR_ACCOUNT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() view returns(address)
func (_KromaTokenMinter *KromaTokenMinterSession) DEPOSITORACCOUNT() (common.Address, error) {
	return _KromaTokenMinter.Contract.DEPOSITORACCOUNT(&_KromaTokenMinter.CallOpts)
}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() view returns(address)
func (_KromaTokenMinter *KromaTokenMinterCallerSession) DEPOSITORACCOUNT() (common.Address, error) {
	return _KromaTokenMinter.Contract.DEPOSITORACCOUNT(&_KromaTokenMinter.CallOpts)
}

// MINTDECREASECAP is a free data retrieval call binding the contract method 0xc27ca149.
//
// Solidity: function MINT_DECREASE_CAP() view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterCaller) MINTDECREASECAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KromaTokenMinter.contract.Call(opts, &out, "MINT_DECREASE_CAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTDECREASECAP is a free data retrieval call binding the contract method 0xc27ca149.
//
// Solidity: function MINT_DECREASE_CAP() view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterSession) MINTDECREASECAP() (*big.Int, error) {
	return _KromaTokenMinter.Contract.MINTDECREASECAP(&_KromaTokenMinter.CallOpts)
}

// MINTDECREASECAP is a free data retrieval call binding the contract method 0xc27ca149.
//
// Solidity: function MINT_DECREASE_CAP() view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterCallerSession) MINTDECREASECAP() (*big.Int, error) {
	return _KromaTokenMinter.Contract.MINTDECREASECAP(&_KromaTokenMinter.CallOpts)
}

// MINTDENOMINATOR is a free data retrieval call binding the contract method 0x2c7dc242.
//
// Solidity: function MINT_DENOMINATOR() view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterCaller) MINTDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KromaTokenMinter.contract.Call(opts, &out, "MINT_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTDENOMINATOR is a free data retrieval call binding the contract method 0x2c7dc242.
//
// Solidity: function MINT_DENOMINATOR() view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterSession) MINTDENOMINATOR() (*big.Int, error) {
	return _KromaTokenMinter.Contract.MINTDENOMINATOR(&_KromaTokenMinter.CallOpts)
}

// MINTDENOMINATOR is a free data retrieval call binding the contract method 0x2c7dc242.
//
// Solidity: function MINT_DENOMINATOR() view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterCallerSession) MINTDENOMINATOR() (*big.Int, error) {
	return _KromaTokenMinter.Contract.MINTDENOMINATOR(&_KromaTokenMinter.CallOpts)
}

// MINTINCREASECAP is a free data retrieval call binding the contract method 0x0ccfab45.
//
// Solidity: function MINT_INCREASE_CAP() view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterCaller) MINTINCREASECAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KromaTokenMinter.contract.Call(opts, &out, "MINT_INCREASE_CAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTINCREASECAP is a free data retrieval call binding the contract method 0x0ccfab45.
//
// Solidity: function MINT_INCREASE_CAP() view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterSession) MINTINCREASECAP() (*big.Int, error) {
	return _KromaTokenMinter.Contract.MINTINCREASECAP(&_KromaTokenMinter.CallOpts)
}

// MINTINCREASECAP is a free data retrieval call binding the contract method 0x0ccfab45.
//
// Solidity: function MINT_INCREASE_CAP() view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterCallerSession) MINTINCREASECAP() (*big.Int, error) {
	return _KromaTokenMinter.Contract.MINTINCREASECAP(&_KromaTokenMinter.CallOpts)
}

// MINTMAXINCREASE is a free data retrieval call binding the contract method 0xe06939a3.
//
// Solidity: function MINT_MAX_INCREASE() view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterCaller) MINTMAXINCREASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KromaTokenMinter.contract.Call(opts, &out, "MINT_MAX_INCREASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTMAXINCREASE is a free data retrieval call binding the contract method 0xe06939a3.
//
// Solidity: function MINT_MAX_INCREASE() view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterSession) MINTMAXINCREASE() (*big.Int, error) {
	return _KromaTokenMinter.Contract.MINTMAXINCREASE(&_KromaTokenMinter.CallOpts)
}

// MINTMAXINCREASE is a free data retrieval call binding the contract method 0xe06939a3.
//
// Solidity: function MINT_MAX_INCREASE() view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterCallerSession) MINTMAXINCREASE() (*big.Int, error) {
	return _KromaTokenMinter.Contract.MINTMAXINCREASE(&_KromaTokenMinter.CallOpts)
}

// MINTMINDECREASE is a free data retrieval call binding the contract method 0x6745d032.
//
// Solidity: function MINT_MIN_DECREASE() view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterCaller) MINTMINDECREASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KromaTokenMinter.contract.Call(opts, &out, "MINT_MIN_DECREASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTMINDECREASE is a free data retrieval call binding the contract method 0x6745d032.
//
// Solidity: function MINT_MIN_DECREASE() view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterSession) MINTMINDECREASE() (*big.Int, error) {
	return _KromaTokenMinter.Contract.MINTMINDECREASE(&_KromaTokenMinter.CallOpts)
}

// MINTMINDECREASE is a free data retrieval call binding the contract method 0x6745d032.
//
// Solidity: function MINT_MIN_DECREASE() view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterCallerSession) MINTMINDECREASE() (*big.Int, error) {
	return _KromaTokenMinter.Contract.MINTMINDECREASE(&_KromaTokenMinter.CallOpts)
}

// SHAREDENOMINATOR is a free data retrieval call binding the contract method 0x7eb11845.
//
// Solidity: function SHARE_DENOMINATOR() view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterCaller) SHAREDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KromaTokenMinter.contract.Call(opts, &out, "SHARE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHAREDENOMINATOR is a free data retrieval call binding the contract method 0x7eb11845.
//
// Solidity: function SHARE_DENOMINATOR() view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterSession) SHAREDENOMINATOR() (*big.Int, error) {
	return _KromaTokenMinter.Contract.SHAREDENOMINATOR(&_KromaTokenMinter.CallOpts)
}

// SHAREDENOMINATOR is a free data retrieval call binding the contract method 0x7eb11845.
//
// Solidity: function SHARE_DENOMINATOR() view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterCallerSession) SHAREDENOMINATOR() (*big.Int, error) {
	return _KromaTokenMinter.Contract.SHAREDENOMINATOR(&_KromaTokenMinter.CallOpts)
}

// KromaToken is a free data retrieval call binding the contract method 0xc9d2b496.
//
// Solidity: function kromaToken() view returns(address)
func (_KromaTokenMinter *KromaTokenMinterCaller) KromaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KromaTokenMinter.contract.Call(opts, &out, "kromaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KromaToken is a free data retrieval call binding the contract method 0xc9d2b496.
//
// Solidity: function kromaToken() view returns(address)
func (_KromaTokenMinter *KromaTokenMinterSession) KromaToken() (common.Address, error) {
	return _KromaTokenMinter.Contract.KromaToken(&_KromaTokenMinter.CallOpts)
}

// KromaToken is a free data retrieval call binding the contract method 0xc9d2b496.
//
// Solidity: function kromaToken() view returns(address)
func (_KromaTokenMinter *KromaTokenMinterCallerSession) KromaToken() (common.Address, error) {
	return _KromaTokenMinter.Contract.KromaToken(&_KromaTokenMinter.CallOpts)
}

// Recipients is a free data retrieval call binding the contract method 0xd1bc76a1.
//
// Solidity: function recipients(uint256 ) view returns(address)
func (_KromaTokenMinter *KromaTokenMinterCaller) Recipients(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KromaTokenMinter.contract.Call(opts, &out, "recipients", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recipients is a free data retrieval call binding the contract method 0xd1bc76a1.
//
// Solidity: function recipients(uint256 ) view returns(address)
func (_KromaTokenMinter *KromaTokenMinterSession) Recipients(arg0 *big.Int) (common.Address, error) {
	return _KromaTokenMinter.Contract.Recipients(&_KromaTokenMinter.CallOpts, arg0)
}

// Recipients is a free data retrieval call binding the contract method 0xd1bc76a1.
//
// Solidity: function recipients(uint256 ) view returns(address)
func (_KromaTokenMinter *KromaTokenMinterCallerSession) Recipients(arg0 *big.Int) (common.Address, error) {
	return _KromaTokenMinter.Contract.Recipients(&_KromaTokenMinter.CallOpts, arg0)
}

// ShareOf is a free data retrieval call binding the contract method 0x21e5e2c4.
//
// Solidity: function shareOf(address ) view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterCaller) ShareOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KromaTokenMinter.contract.Call(opts, &out, "shareOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ShareOf is a free data retrieval call binding the contract method 0x21e5e2c4.
//
// Solidity: function shareOf(address ) view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterSession) ShareOf(arg0 common.Address) (*big.Int, error) {
	return _KromaTokenMinter.Contract.ShareOf(&_KromaTokenMinter.CallOpts, arg0)
}

// ShareOf is a free data retrieval call binding the contract method 0x21e5e2c4.
//
// Solidity: function shareOf(address ) view returns(uint256)
func (_KromaTokenMinter *KromaTokenMinterCallerSession) ShareOf(arg0 common.Address) (*big.Int, error) {
	return _KromaTokenMinter.Contract.ShareOf(&_KromaTokenMinter.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xfdb6583e.
//
// Solidity: function initialize(address[] _recipients, uint64[] _shares) returns()
func (_KromaTokenMinter *KromaTokenMinterTransactor) Initialize(opts *bind.TransactOpts, _recipients []common.Address, _shares []uint64) (*types.Transaction, error) {
	return _KromaTokenMinter.contract.Transact(opts, "initialize", _recipients, _shares)
}

// Initialize is a paid mutator transaction binding the contract method 0xfdb6583e.
//
// Solidity: function initialize(address[] _recipients, uint64[] _shares) returns()
func (_KromaTokenMinter *KromaTokenMinterSession) Initialize(_recipients []common.Address, _shares []uint64) (*types.Transaction, error) {
	return _KromaTokenMinter.Contract.Initialize(&_KromaTokenMinter.TransactOpts, _recipients, _shares)
}

// Initialize is a paid mutator transaction binding the contract method 0xfdb6583e.
//
// Solidity: function initialize(address[] _recipients, uint64[] _shares) returns()
func (_KromaTokenMinter *KromaTokenMinterTransactorSession) Initialize(_recipients []common.Address, _shares []uint64) (*types.Transaction, error) {
	return _KromaTokenMinter.Contract.Initialize(&_KromaTokenMinter.TransactOpts, _recipients, _shares)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 totalAmount) returns()
func (_KromaTokenMinter *KromaTokenMinterTransactor) Mint(opts *bind.TransactOpts, totalAmount *big.Int) (*types.Transaction, error) {
	return _KromaTokenMinter.contract.Transact(opts, "mint", totalAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 totalAmount) returns()
func (_KromaTokenMinter *KromaTokenMinterSession) Mint(totalAmount *big.Int) (*types.Transaction, error) {
	return _KromaTokenMinter.Contract.Mint(&_KromaTokenMinter.TransactOpts, totalAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 totalAmount) returns()
func (_KromaTokenMinter *KromaTokenMinterTransactorSession) Mint(totalAmount *big.Int) (*types.Transaction, error) {
	return _KromaTokenMinter.Contract.Mint(&_KromaTokenMinter.TransactOpts, totalAmount)
}

// KromaTokenMinterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the KromaTokenMinter contract.
type KromaTokenMinterInitializedIterator struct {
	Event *KromaTokenMinterInitialized // Event containing the contract specifics and raw log

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
func (it *KromaTokenMinterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaTokenMinterInitialized)
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
		it.Event = new(KromaTokenMinterInitialized)
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
func (it *KromaTokenMinterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaTokenMinterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaTokenMinterInitialized represents a Initialized event raised by the KromaTokenMinter contract.
type KromaTokenMinterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_KromaTokenMinter *KromaTokenMinterFilterer) FilterInitialized(opts *bind.FilterOpts) (*KromaTokenMinterInitializedIterator, error) {

	logs, sub, err := _KromaTokenMinter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &KromaTokenMinterInitializedIterator{contract: _KromaTokenMinter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_KromaTokenMinter *KromaTokenMinterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *KromaTokenMinterInitialized) (event.Subscription, error) {

	logs, sub, err := _KromaTokenMinter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaTokenMinterInitialized)
				if err := _KromaTokenMinter.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_KromaTokenMinter *KromaTokenMinterFilterer) ParseInitialized(log types.Log) (*KromaTokenMinterInitialized, error) {
	event := new(KromaTokenMinterInitialized)
	if err := _KromaTokenMinter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
