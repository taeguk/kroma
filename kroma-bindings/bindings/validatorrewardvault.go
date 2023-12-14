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

// ValidatorRewardVaultMetaData contains all meta data concerning the ValidatorRewardVault contract.
var ValidatorRewardVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardDivider\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Rewarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_WITHDRAWAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECIPIENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_DIVIDER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_POOL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalProcessed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalReserved\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawToL2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x61016060405234801561001157600080fd5b50604051610eb9380380610eb983398101604081905261003091610063565b6000608081905260a0819052600160c05260e0819052610100526001600160a01b0391909116610120526101405261009d565b6000806040838503121561007657600080fd5b82516001600160a01b038116811461008d57600080fd5b6020939093015192949293505050565b60805160a05160c05160e051610100516101205161014051610dae61010b6000396000818161019501526107f901526000818161024501526102c8015260006106ef015260006106c60152600061069d0152600060de01526000818161028f01526108470152610dae6000f3fe6080604052600436106100c05760003560e01c80636ed39f6211610074578063b98debbf1161004e578063b98debbf14610233578063c71b0e1c14610267578063d3e5792b1461027d57600080fd5b80636ed39f62146101c557806370a08231146101da57806384411d651461021d57600080fd5b80633ccfd60b116100a55780633ccfd60b1461014c57806354fd4d501461016157806362aba76b1461018357600080fd5b80630d9019e1146100cc57806321670f221461012a57600080fd5b366100c757005b600080fd5b3480156100d857600080fd5b506101007f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561013657600080fd5b5061014a610145366004610b66565b6102b1565b005b34801561015857600080fd5b5061014a6105f0565b34801561016d57600080fd5b50610176610696565b6040516101219190610c0a565b34801561018f57600080fd5b506101b77f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610121565b3480156101d157600080fd5b5061014a610739565b3480156101e657600080fd5b506101b76101f5366004610c24565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205490565b34801561022957600080fd5b506101b760005481565b34801561023f57600080fd5b506101007f000000000000000000000000000000000000000000000000000000000000000081565b34801561027357600080fd5b506101b760035481565b34801561028957600080fd5b506101b77f000000000000000000000000000000000000000000000000000000000000000081565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000167fffffffffffffffffffffffffeeeeffffffffffffffffffffffffffffffffeeef330173ffffffffffffffffffffffffffffffffffffffff16146103d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604860248201527f56616c696461746f725265776172645661756c743a2066756e6374696f6e206360448201527f616e206f6e6c792062652063616c6c65642066726f6d207468652056616c696460648201527f61746f72506f6f6c000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821661047c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603360248201527f56616c696461746f725265776172645661756c743a2076616c696461746f722060448201527f616464726573732063616e6e6f7420626520300000000000000000000000000060648201526084016103d0565b60008181526002602052604090205460ff1615610541576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604e60248201527f56616c696461746f725265776172645661756c743a207468652072657761726460448201527f2068617320616c7265616479206265656e207061696420666f7220746865204c60648201527f3220626c6f636b206e756d626572000000000000000000000000000000000000608482015260a4016103d0565b600061054b6107f5565b600380548201905573ffffffffffffffffffffffffffffffffffffffff84166000818152600160208181526040808420805487019055878452600282529283902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016909217909155905183815292935084927f291e8ba3c0f4b0bd86e6e2346fcee1e7ca0975b1cc1886bfbc722d34f3f24d91910160405180910390a3505050565b60006105fa610835565b604080516020810182526000815290517fe11013dd0000000000000000000000000000000000000000000000000000000081529192507342000000000000000000000000000000000000099163e11013dd9184916106619133916188b89190600401610c3f565b6000604051808303818588803b15801561067a57600080fd5b505af115801561068e573d6000803e3d6000fd5b505050505050565b60606106c17f0000000000000000000000000000000000000000000000000000000000000000610982565b6106ea7f0000000000000000000000000000000000000000000000000000000000000000610982565b6107137f0000000000000000000000000000000000000000000000000000000000000000610982565b60405160200161072593929190610c83565b604051602081830303815290604052905090565b6000610743610835565b90506000610762335a8460405180602001604052806000815250610a40565b9050806107f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f56616c696461746f725265776172645661756c743a20455448207472616e736660448201527f6572206661696c6564000000000000000000000000000000000000000000000060648201526084016103d0565b5050565b60007f0000000000000000000000000000000000000000000000000000000000000000600354476108269190610cf9565b6108309190610d37565b905090565b336000908152600160205260408120547f000000000000000000000000000000000000000000000000000000000000000081101561091b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605660248201527f56616c696461746f725265776172645661756c743a207769746864726177616c60448201527f20616d6f756e74206d7573742062652067726561746572207468616e206d696e60648201527f696d756d207769746864726177616c20616d6f756e7400000000000000000000608482015260a4016103d0565b33600081815260016020908152604080832083905560038054869003905582548501909255815184815290810183905280820192909252517fc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba9181900360600190a1919050565b6060600061098f83610a5a565b600101905060008167ffffffffffffffff8111156109af576109af610d72565b6040519080825280601f01601f1916602001820160405280156109d9576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85049450846109e357509392505050565b600080600080845160208601878a8af19695505050505050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310610aa3577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310610acf576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310610aed57662386f26fc10000830492506010015b6305f5e1008310610b05576305f5e100830492506008015b6127108310610b1957612710830492506004015b60648310610b2b576064830492506002015b600a8310610b37576001015b92915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610b6157600080fd5b919050565b60008060408385031215610b7957600080fd5b610b8283610b3d565b946020939093013593505050565b60005b83811015610bab578181015183820152602001610b93565b83811115610bba576000848401525b50505050565b60008151808452610bd8816020860160208601610b90565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610c1d6020830184610bc0565b9392505050565b600060208284031215610c3657600080fd5b610c1d82610b3d565b73ffffffffffffffffffffffffffffffffffffffff8416815263ffffffff83166020820152606060408201526000610c7a6060830184610bc0565b95945050505050565b60008451610c95818460208901610b90565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610cd1816001850160208a01610b90565b60019201918201528351610cec816002840160208801610b90565b0160020195945050505050565b600082821015610d32577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b500390565b600082610d6d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c634300080f000a",
}

// ValidatorRewardVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorRewardVaultMetaData.ABI instead.
var ValidatorRewardVaultABI = ValidatorRewardVaultMetaData.ABI

// ValidatorRewardVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorRewardVaultMetaData.Bin instead.
var ValidatorRewardVaultBin = ValidatorRewardVaultMetaData.Bin

// DeployValidatorRewardVault deploys a new Ethereum contract, binding an instance of ValidatorRewardVault to it.
func DeployValidatorRewardVault(auth *bind.TransactOpts, backend bind.ContractBackend, _validatorPool common.Address, _rewardDivider *big.Int) (common.Address, *types.Transaction, *ValidatorRewardVault, error) {
	parsed, err := ValidatorRewardVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorRewardVaultBin), backend, _validatorPool, _rewardDivider)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorRewardVault{ValidatorRewardVaultCaller: ValidatorRewardVaultCaller{contract: contract}, ValidatorRewardVaultTransactor: ValidatorRewardVaultTransactor{contract: contract}, ValidatorRewardVaultFilterer: ValidatorRewardVaultFilterer{contract: contract}}, nil
}

// ValidatorRewardVault is an auto generated Go binding around an Ethereum contract.
type ValidatorRewardVault struct {
	ValidatorRewardVaultCaller     // Read-only binding to the contract
	ValidatorRewardVaultTransactor // Write-only binding to the contract
	ValidatorRewardVaultFilterer   // Log filterer for contract events
}

// ValidatorRewardVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorRewardVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorRewardVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorRewardVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorRewardVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorRewardVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorRewardVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorRewardVaultSession struct {
	Contract     *ValidatorRewardVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ValidatorRewardVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorRewardVaultCallerSession struct {
	Contract *ValidatorRewardVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ValidatorRewardVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorRewardVaultTransactorSession struct {
	Contract     *ValidatorRewardVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ValidatorRewardVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorRewardVaultRaw struct {
	Contract *ValidatorRewardVault // Generic contract binding to access the raw methods on
}

// ValidatorRewardVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorRewardVaultCallerRaw struct {
	Contract *ValidatorRewardVaultCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorRewardVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorRewardVaultTransactorRaw struct {
	Contract *ValidatorRewardVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorRewardVault creates a new instance of ValidatorRewardVault, bound to a specific deployed contract.
func NewValidatorRewardVault(address common.Address, backend bind.ContractBackend) (*ValidatorRewardVault, error) {
	contract, err := bindValidatorRewardVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardVault{ValidatorRewardVaultCaller: ValidatorRewardVaultCaller{contract: contract}, ValidatorRewardVaultTransactor: ValidatorRewardVaultTransactor{contract: contract}, ValidatorRewardVaultFilterer: ValidatorRewardVaultFilterer{contract: contract}}, nil
}

// NewValidatorRewardVaultCaller creates a new read-only instance of ValidatorRewardVault, bound to a specific deployed contract.
func NewValidatorRewardVaultCaller(address common.Address, caller bind.ContractCaller) (*ValidatorRewardVaultCaller, error) {
	contract, err := bindValidatorRewardVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardVaultCaller{contract: contract}, nil
}

// NewValidatorRewardVaultTransactor creates a new write-only instance of ValidatorRewardVault, bound to a specific deployed contract.
func NewValidatorRewardVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorRewardVaultTransactor, error) {
	contract, err := bindValidatorRewardVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardVaultTransactor{contract: contract}, nil
}

// NewValidatorRewardVaultFilterer creates a new log filterer instance of ValidatorRewardVault, bound to a specific deployed contract.
func NewValidatorRewardVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorRewardVaultFilterer, error) {
	contract, err := bindValidatorRewardVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardVaultFilterer{contract: contract}, nil
}

// bindValidatorRewardVault binds a generic wrapper to an already deployed contract.
func bindValidatorRewardVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorRewardVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorRewardVault *ValidatorRewardVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorRewardVault.Contract.ValidatorRewardVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorRewardVault *ValidatorRewardVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.ValidatorRewardVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorRewardVault *ValidatorRewardVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.ValidatorRewardVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorRewardVault *ValidatorRewardVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorRewardVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorRewardVault *ValidatorRewardVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorRewardVault *ValidatorRewardVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.contract.Transact(opts, method, params...)
}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCaller) MINWITHDRAWALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorRewardVault.contract.Call(opts, &out, "MIN_WITHDRAWAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultSession) MINWITHDRAWALAMOUNT() (*big.Int, error) {
	return _ValidatorRewardVault.Contract.MINWITHDRAWALAMOUNT(&_ValidatorRewardVault.CallOpts)
}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCallerSession) MINWITHDRAWALAMOUNT() (*big.Int, error) {
	return _ValidatorRewardVault.Contract.MINWITHDRAWALAMOUNT(&_ValidatorRewardVault.CallOpts)
}

// RECIPIENT is a free data retrieval call binding the contract method 0x0d9019e1.
//
// Solidity: function RECIPIENT() view returns(address)
func (_ValidatorRewardVault *ValidatorRewardVaultCaller) RECIPIENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorRewardVault.contract.Call(opts, &out, "RECIPIENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RECIPIENT is a free data retrieval call binding the contract method 0x0d9019e1.
//
// Solidity: function RECIPIENT() view returns(address)
func (_ValidatorRewardVault *ValidatorRewardVaultSession) RECIPIENT() (common.Address, error) {
	return _ValidatorRewardVault.Contract.RECIPIENT(&_ValidatorRewardVault.CallOpts)
}

// RECIPIENT is a free data retrieval call binding the contract method 0x0d9019e1.
//
// Solidity: function RECIPIENT() view returns(address)
func (_ValidatorRewardVault *ValidatorRewardVaultCallerSession) RECIPIENT() (common.Address, error) {
	return _ValidatorRewardVault.Contract.RECIPIENT(&_ValidatorRewardVault.CallOpts)
}

// REWARDDIVIDER is a free data retrieval call binding the contract method 0x62aba76b.
//
// Solidity: function REWARD_DIVIDER() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCaller) REWARDDIVIDER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorRewardVault.contract.Call(opts, &out, "REWARD_DIVIDER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REWARDDIVIDER is a free data retrieval call binding the contract method 0x62aba76b.
//
// Solidity: function REWARD_DIVIDER() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultSession) REWARDDIVIDER() (*big.Int, error) {
	return _ValidatorRewardVault.Contract.REWARDDIVIDER(&_ValidatorRewardVault.CallOpts)
}

// REWARDDIVIDER is a free data retrieval call binding the contract method 0x62aba76b.
//
// Solidity: function REWARD_DIVIDER() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCallerSession) REWARDDIVIDER() (*big.Int, error) {
	return _ValidatorRewardVault.Contract.REWARDDIVIDER(&_ValidatorRewardVault.CallOpts)
}

// VALIDATORPOOL is a free data retrieval call binding the contract method 0xb98debbf.
//
// Solidity: function VALIDATOR_POOL() view returns(address)
func (_ValidatorRewardVault *ValidatorRewardVaultCaller) VALIDATORPOOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorRewardVault.contract.Call(opts, &out, "VALIDATOR_POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORPOOL is a free data retrieval call binding the contract method 0xb98debbf.
//
// Solidity: function VALIDATOR_POOL() view returns(address)
func (_ValidatorRewardVault *ValidatorRewardVaultSession) VALIDATORPOOL() (common.Address, error) {
	return _ValidatorRewardVault.Contract.VALIDATORPOOL(&_ValidatorRewardVault.CallOpts)
}

// VALIDATORPOOL is a free data retrieval call binding the contract method 0xb98debbf.
//
// Solidity: function VALIDATOR_POOL() view returns(address)
func (_ValidatorRewardVault *ValidatorRewardVaultCallerSession) VALIDATORPOOL() (common.Address, error) {
	return _ValidatorRewardVault.Contract.VALIDATORPOOL(&_ValidatorRewardVault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCaller) BalanceOf(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorRewardVault.contract.Call(opts, &out, "balanceOf", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultSession) BalanceOf(_addr common.Address) (*big.Int, error) {
	return _ValidatorRewardVault.Contract.BalanceOf(&_ValidatorRewardVault.CallOpts, _addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCallerSession) BalanceOf(_addr common.Address) (*big.Int, error) {
	return _ValidatorRewardVault.Contract.BalanceOf(&_ValidatorRewardVault.CallOpts, _addr)
}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCaller) TotalProcessed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorRewardVault.contract.Call(opts, &out, "totalProcessed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultSession) TotalProcessed() (*big.Int, error) {
	return _ValidatorRewardVault.Contract.TotalProcessed(&_ValidatorRewardVault.CallOpts)
}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCallerSession) TotalProcessed() (*big.Int, error) {
	return _ValidatorRewardVault.Contract.TotalProcessed(&_ValidatorRewardVault.CallOpts)
}

// TotalReserved is a free data retrieval call binding the contract method 0xc71b0e1c.
//
// Solidity: function totalReserved() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCaller) TotalReserved(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorRewardVault.contract.Call(opts, &out, "totalReserved")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserved is a free data retrieval call binding the contract method 0xc71b0e1c.
//
// Solidity: function totalReserved() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultSession) TotalReserved() (*big.Int, error) {
	return _ValidatorRewardVault.Contract.TotalReserved(&_ValidatorRewardVault.CallOpts)
}

// TotalReserved is a free data retrieval call binding the contract method 0xc71b0e1c.
//
// Solidity: function totalReserved() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCallerSession) TotalReserved() (*big.Int, error) {
	return _ValidatorRewardVault.Contract.TotalReserved(&_ValidatorRewardVault.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ValidatorRewardVault *ValidatorRewardVaultCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ValidatorRewardVault.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ValidatorRewardVault *ValidatorRewardVaultSession) Version() (string, error) {
	return _ValidatorRewardVault.Contract.Version(&_ValidatorRewardVault.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ValidatorRewardVault *ValidatorRewardVaultCallerSession) Version() (string, error) {
	return _ValidatorRewardVault.Contract.Version(&_ValidatorRewardVault.CallOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x21670f22.
//
// Solidity: function reward(address _validator, uint256 _l2BlockNumber) returns()
func (_ValidatorRewardVault *ValidatorRewardVaultTransactor) Reward(opts *bind.TransactOpts, _validator common.Address, _l2BlockNumber *big.Int) (*types.Transaction, error) {
	return _ValidatorRewardVault.contract.Transact(opts, "reward", _validator, _l2BlockNumber)
}

// Reward is a paid mutator transaction binding the contract method 0x21670f22.
//
// Solidity: function reward(address _validator, uint256 _l2BlockNumber) returns()
func (_ValidatorRewardVault *ValidatorRewardVaultSession) Reward(_validator common.Address, _l2BlockNumber *big.Int) (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.Reward(&_ValidatorRewardVault.TransactOpts, _validator, _l2BlockNumber)
}

// Reward is a paid mutator transaction binding the contract method 0x21670f22.
//
// Solidity: function reward(address _validator, uint256 _l2BlockNumber) returns()
func (_ValidatorRewardVault *ValidatorRewardVaultTransactorSession) Reward(_validator common.Address, _l2BlockNumber *big.Int) (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.Reward(&_ValidatorRewardVault.TransactOpts, _validator, _l2BlockNumber)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ValidatorRewardVault *ValidatorRewardVaultTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorRewardVault.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ValidatorRewardVault *ValidatorRewardVaultSession) Withdraw() (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.Withdraw(&_ValidatorRewardVault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ValidatorRewardVault *ValidatorRewardVaultTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.Withdraw(&_ValidatorRewardVault.TransactOpts)
}

// WithdrawToL2 is a paid mutator transaction binding the contract method 0x6ed39f62.
//
// Solidity: function withdrawToL2() returns()
func (_ValidatorRewardVault *ValidatorRewardVaultTransactor) WithdrawToL2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorRewardVault.contract.Transact(opts, "withdrawToL2")
}

// WithdrawToL2 is a paid mutator transaction binding the contract method 0x6ed39f62.
//
// Solidity: function withdrawToL2() returns()
func (_ValidatorRewardVault *ValidatorRewardVaultSession) WithdrawToL2() (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.WithdrawToL2(&_ValidatorRewardVault.TransactOpts)
}

// WithdrawToL2 is a paid mutator transaction binding the contract method 0x6ed39f62.
//
// Solidity: function withdrawToL2() returns()
func (_ValidatorRewardVault *ValidatorRewardVaultTransactorSession) WithdrawToL2() (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.WithdrawToL2(&_ValidatorRewardVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidatorRewardVault *ValidatorRewardVaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorRewardVault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidatorRewardVault *ValidatorRewardVaultSession) Receive() (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.Receive(&_ValidatorRewardVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidatorRewardVault *ValidatorRewardVaultTransactorSession) Receive() (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.Receive(&_ValidatorRewardVault.TransactOpts)
}

// ValidatorRewardVaultRewardedIterator is returned from FilterRewarded and is used to iterate over the raw logs and unpacked data for Rewarded events raised by the ValidatorRewardVault contract.
type ValidatorRewardVaultRewardedIterator struct {
	Event *ValidatorRewardVaultRewarded // Event containing the contract specifics and raw log

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
func (it *ValidatorRewardVaultRewardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRewardVaultRewarded)
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
		it.Event = new(ValidatorRewardVaultRewarded)
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
func (it *ValidatorRewardVaultRewardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRewardVaultRewardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRewardVaultRewarded represents a Rewarded event raised by the ValidatorRewardVault contract.
type ValidatorRewardVaultRewarded struct {
	Validator     common.Address
	L2BlockNumber *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRewarded is a free log retrieval operation binding the contract event 0x291e8ba3c0f4b0bd86e6e2346fcee1e7ca0975b1cc1886bfbc722d34f3f24d91.
//
// Solidity: event Rewarded(address indexed validator, uint256 indexed l2BlockNumber, uint256 amount)
func (_ValidatorRewardVault *ValidatorRewardVaultFilterer) FilterRewarded(opts *bind.FilterOpts, validator []common.Address, l2BlockNumber []*big.Int) (*ValidatorRewardVaultRewardedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _ValidatorRewardVault.contract.FilterLogs(opts, "Rewarded", validatorRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardVaultRewardedIterator{contract: _ValidatorRewardVault.contract, event: "Rewarded", logs: logs, sub: sub}, nil
}

// WatchRewarded is a free log subscription operation binding the contract event 0x291e8ba3c0f4b0bd86e6e2346fcee1e7ca0975b1cc1886bfbc722d34f3f24d91.
//
// Solidity: event Rewarded(address indexed validator, uint256 indexed l2BlockNumber, uint256 amount)
func (_ValidatorRewardVault *ValidatorRewardVaultFilterer) WatchRewarded(opts *bind.WatchOpts, sink chan<- *ValidatorRewardVaultRewarded, validator []common.Address, l2BlockNumber []*big.Int) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _ValidatorRewardVault.contract.WatchLogs(opts, "Rewarded", validatorRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRewardVaultRewarded)
				if err := _ValidatorRewardVault.contract.UnpackLog(event, "Rewarded", log); err != nil {
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

// ParseRewarded is a log parse operation binding the contract event 0x291e8ba3c0f4b0bd86e6e2346fcee1e7ca0975b1cc1886bfbc722d34f3f24d91.
//
// Solidity: event Rewarded(address indexed validator, uint256 indexed l2BlockNumber, uint256 amount)
func (_ValidatorRewardVault *ValidatorRewardVaultFilterer) ParseRewarded(log types.Log) (*ValidatorRewardVaultRewarded, error) {
	event := new(ValidatorRewardVaultRewarded)
	if err := _ValidatorRewardVault.contract.UnpackLog(event, "Rewarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorRewardVaultWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the ValidatorRewardVault contract.
type ValidatorRewardVaultWithdrawalIterator struct {
	Event *ValidatorRewardVaultWithdrawal // Event containing the contract specifics and raw log

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
func (it *ValidatorRewardVaultWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRewardVaultWithdrawal)
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
		it.Event = new(ValidatorRewardVaultWithdrawal)
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
func (it *ValidatorRewardVaultWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRewardVaultWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRewardVaultWithdrawal represents a Withdrawal event raised by the ValidatorRewardVault contract.
type ValidatorRewardVaultWithdrawal struct {
	Value *big.Int
	To    common.Address
	From  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_ValidatorRewardVault *ValidatorRewardVaultFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*ValidatorRewardVaultWithdrawalIterator, error) {

	logs, sub, err := _ValidatorRewardVault.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardVaultWithdrawalIterator{contract: _ValidatorRewardVault.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_ValidatorRewardVault *ValidatorRewardVaultFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ValidatorRewardVaultWithdrawal) (event.Subscription, error) {

	logs, sub, err := _ValidatorRewardVault.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRewardVaultWithdrawal)
				if err := _ValidatorRewardVault.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_ValidatorRewardVault *ValidatorRewardVaultFilterer) ParseWithdrawal(log types.Log) (*ValidatorRewardVaultWithdrawal, error) {
	event := new(ValidatorRewardVaultWithdrawal)
	if err := _ValidatorRewardVault.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
