// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PriceSubmitterBinding

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

// PriceSubmitterBindingMetaData contains all meta data concerning the PriceSubmitterBinding contract.
var PriceSubmitterBindingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposedGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldGovernance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newGoveranance\",\"type\":\"address\"}],\"name\":\"GovernanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIFtsoGenesis[]\",\"name\":\"ftsos\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"hashes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PriceHashesSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIFtsoGenesis[]\",\"name\":\"ftsos\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"randoms\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PricesRevealed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"claimGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFtsoManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFtsoRegistry\",\"outputs\":[{\"internalType\":\"contractIFtsoRegistryGenesis\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTrustedAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVoterWhitelister\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialiseFixedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"proposeGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposedGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_ftsoIndices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_randoms\",\"type\":\"uint256[]\"}],\"name\":\"revealPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFtsoRegistryGenesis\",\"name\":\"_ftsoRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voterWhitelister\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ftsoManager\",\"type\":\"address\"}],\"name\":\"setContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_trustedAddresses\",\"type\":\"address[]\"}],\"name\":\"setTrustedAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_ftsoIndices\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_hashes\",\"type\":\"bytes32[]\"}],\"name\":\"submitPriceHashes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"transferGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"voterWhitelistBitmap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ftsoIndex\",\"type\":\"uint256\"}],\"name\":\"voterWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_removedVoters\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_ftsoIndex\",\"type\":\"uint256\"}],\"name\":\"votersRemovedFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PriceSubmitterBindingABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceSubmitterBindingMetaData.ABI instead.
var PriceSubmitterBindingABI = PriceSubmitterBindingMetaData.ABI

// PriceSubmitterBinding is an auto generated Go binding around an Ethereum contract.
type PriceSubmitterBinding struct {
	PriceSubmitterBindingCaller     // Read-only binding to the contract
	PriceSubmitterBindingTransactor // Write-only binding to the contract
	PriceSubmitterBindingFilterer   // Log filterer for contract events
}

// PriceSubmitterBindingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceSubmitterBindingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceSubmitterBindingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceSubmitterBindingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceSubmitterBindingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceSubmitterBindingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceSubmitterBindingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceSubmitterBindingSession struct {
	Contract     *PriceSubmitterBinding // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PriceSubmitterBindingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceSubmitterBindingCallerSession struct {
	Contract *PriceSubmitterBindingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// PriceSubmitterBindingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceSubmitterBindingTransactorSession struct {
	Contract     *PriceSubmitterBindingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// PriceSubmitterBindingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceSubmitterBindingRaw struct {
	Contract *PriceSubmitterBinding // Generic contract binding to access the raw methods on
}

// PriceSubmitterBindingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceSubmitterBindingCallerRaw struct {
	Contract *PriceSubmitterBindingCaller // Generic read-only contract binding to access the raw methods on
}

// PriceSubmitterBindingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceSubmitterBindingTransactorRaw struct {
	Contract *PriceSubmitterBindingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceSubmitterBinding creates a new instance of PriceSubmitterBinding, bound to a specific deployed contract.
func NewPriceSubmitterBinding(address common.Address, backend bind.ContractBackend) (*PriceSubmitterBinding, error) {
	contract, err := bindPriceSubmitterBinding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceSubmitterBinding{PriceSubmitterBindingCaller: PriceSubmitterBindingCaller{contract: contract}, PriceSubmitterBindingTransactor: PriceSubmitterBindingTransactor{contract: contract}, PriceSubmitterBindingFilterer: PriceSubmitterBindingFilterer{contract: contract}}, nil
}

// NewPriceSubmitterBindingCaller creates a new read-only instance of PriceSubmitterBinding, bound to a specific deployed contract.
func NewPriceSubmitterBindingCaller(address common.Address, caller bind.ContractCaller) (*PriceSubmitterBindingCaller, error) {
	contract, err := bindPriceSubmitterBinding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceSubmitterBindingCaller{contract: contract}, nil
}

// NewPriceSubmitterBindingTransactor creates a new write-only instance of PriceSubmitterBinding, bound to a specific deployed contract.
func NewPriceSubmitterBindingTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceSubmitterBindingTransactor, error) {
	contract, err := bindPriceSubmitterBinding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceSubmitterBindingTransactor{contract: contract}, nil
}

// NewPriceSubmitterBindingFilterer creates a new log filterer instance of PriceSubmitterBinding, bound to a specific deployed contract.
func NewPriceSubmitterBindingFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceSubmitterBindingFilterer, error) {
	contract, err := bindPriceSubmitterBinding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceSubmitterBindingFilterer{contract: contract}, nil
}

// bindPriceSubmitterBinding binds a generic wrapper to an already deployed contract.
func bindPriceSubmitterBinding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceSubmitterBindingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceSubmitterBinding *PriceSubmitterBindingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceSubmitterBinding.Contract.PriceSubmitterBindingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceSubmitterBinding *PriceSubmitterBindingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.PriceSubmitterBindingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceSubmitterBinding *PriceSubmitterBindingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.PriceSubmitterBindingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceSubmitterBinding *PriceSubmitterBindingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceSubmitterBinding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.contract.Transact(opts, method, params...)
}

// GetFtsoManager is a free data retrieval call binding the contract method 0xb39c6858.
//
// Solidity: function getFtsoManager() view returns(address)
func (_PriceSubmitterBinding *PriceSubmitterBindingCaller) GetFtsoManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceSubmitterBinding.contract.Call(opts, &out, "getFtsoManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFtsoManager is a free data retrieval call binding the contract method 0xb39c6858.
//
// Solidity: function getFtsoManager() view returns(address)
func (_PriceSubmitterBinding *PriceSubmitterBindingSession) GetFtsoManager() (common.Address, error) {
	return _PriceSubmitterBinding.Contract.GetFtsoManager(&_PriceSubmitterBinding.CallOpts)
}

// GetFtsoManager is a free data retrieval call binding the contract method 0xb39c6858.
//
// Solidity: function getFtsoManager() view returns(address)
func (_PriceSubmitterBinding *PriceSubmitterBindingCallerSession) GetFtsoManager() (common.Address, error) {
	return _PriceSubmitterBinding.Contract.GetFtsoManager(&_PriceSubmitterBinding.CallOpts)
}

// GetFtsoRegistry is a free data retrieval call binding the contract method 0x8c9d28b6.
//
// Solidity: function getFtsoRegistry() view returns(address)
func (_PriceSubmitterBinding *PriceSubmitterBindingCaller) GetFtsoRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceSubmitterBinding.contract.Call(opts, &out, "getFtsoRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFtsoRegistry is a free data retrieval call binding the contract method 0x8c9d28b6.
//
// Solidity: function getFtsoRegistry() view returns(address)
func (_PriceSubmitterBinding *PriceSubmitterBindingSession) GetFtsoRegistry() (common.Address, error) {
	return _PriceSubmitterBinding.Contract.GetFtsoRegistry(&_PriceSubmitterBinding.CallOpts)
}

// GetFtsoRegistry is a free data retrieval call binding the contract method 0x8c9d28b6.
//
// Solidity: function getFtsoRegistry() view returns(address)
func (_PriceSubmitterBinding *PriceSubmitterBindingCallerSession) GetFtsoRegistry() (common.Address, error) {
	return _PriceSubmitterBinding.Contract.GetFtsoRegistry(&_PriceSubmitterBinding.CallOpts)
}

// GetTrustedAddresses is a free data retrieval call binding the contract method 0xffacb84e.
//
// Solidity: function getTrustedAddresses() view returns(address[])
func (_PriceSubmitterBinding *PriceSubmitterBindingCaller) GetTrustedAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PriceSubmitterBinding.contract.Call(opts, &out, "getTrustedAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTrustedAddresses is a free data retrieval call binding the contract method 0xffacb84e.
//
// Solidity: function getTrustedAddresses() view returns(address[])
func (_PriceSubmitterBinding *PriceSubmitterBindingSession) GetTrustedAddresses() ([]common.Address, error) {
	return _PriceSubmitterBinding.Contract.GetTrustedAddresses(&_PriceSubmitterBinding.CallOpts)
}

// GetTrustedAddresses is a free data retrieval call binding the contract method 0xffacb84e.
//
// Solidity: function getTrustedAddresses() view returns(address[])
func (_PriceSubmitterBinding *PriceSubmitterBindingCallerSession) GetTrustedAddresses() ([]common.Address, error) {
	return _PriceSubmitterBinding.Contract.GetTrustedAddresses(&_PriceSubmitterBinding.CallOpts)
}

// GetVoterWhitelister is a free data retrieval call binding the contract method 0x71e1fad9.
//
// Solidity: function getVoterWhitelister() view returns(address)
func (_PriceSubmitterBinding *PriceSubmitterBindingCaller) GetVoterWhitelister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceSubmitterBinding.contract.Call(opts, &out, "getVoterWhitelister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVoterWhitelister is a free data retrieval call binding the contract method 0x71e1fad9.
//
// Solidity: function getVoterWhitelister() view returns(address)
func (_PriceSubmitterBinding *PriceSubmitterBindingSession) GetVoterWhitelister() (common.Address, error) {
	return _PriceSubmitterBinding.Contract.GetVoterWhitelister(&_PriceSubmitterBinding.CallOpts)
}

// GetVoterWhitelister is a free data retrieval call binding the contract method 0x71e1fad9.
//
// Solidity: function getVoterWhitelister() view returns(address)
func (_PriceSubmitterBinding *PriceSubmitterBindingCallerSession) GetVoterWhitelister() (common.Address, error) {
	return _PriceSubmitterBinding.Contract.GetVoterWhitelister(&_PriceSubmitterBinding.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PriceSubmitterBinding *PriceSubmitterBindingCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceSubmitterBinding.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PriceSubmitterBinding *PriceSubmitterBindingSession) Governance() (common.Address, error) {
	return _PriceSubmitterBinding.Contract.Governance(&_PriceSubmitterBinding.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PriceSubmitterBinding *PriceSubmitterBindingCallerSession) Governance() (common.Address, error) {
	return _PriceSubmitterBinding.Contract.Governance(&_PriceSubmitterBinding.CallOpts)
}

// Initialise is a free data retrieval call binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _governance) pure returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingCaller) Initialise(opts *bind.CallOpts, _governance common.Address) error {
	var out []interface{}
	err := _PriceSubmitterBinding.contract.Call(opts, &out, "initialise", _governance)

	if err != nil {
		return err
	}

	return err

}

// Initialise is a free data retrieval call binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _governance) pure returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingSession) Initialise(_governance common.Address) error {
	return _PriceSubmitterBinding.Contract.Initialise(&_PriceSubmitterBinding.CallOpts, _governance)
}

// Initialise is a free data retrieval call binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _governance) pure returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingCallerSession) Initialise(_governance common.Address) error {
	return _PriceSubmitterBinding.Contract.Initialise(&_PriceSubmitterBinding.CallOpts, _governance)
}

// ProposedGovernance is a free data retrieval call binding the contract method 0x60f7ac97.
//
// Solidity: function proposedGovernance() view returns(address)
func (_PriceSubmitterBinding *PriceSubmitterBindingCaller) ProposedGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceSubmitterBinding.contract.Call(opts, &out, "proposedGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposedGovernance is a free data retrieval call binding the contract method 0x60f7ac97.
//
// Solidity: function proposedGovernance() view returns(address)
func (_PriceSubmitterBinding *PriceSubmitterBindingSession) ProposedGovernance() (common.Address, error) {
	return _PriceSubmitterBinding.Contract.ProposedGovernance(&_PriceSubmitterBinding.CallOpts)
}

// ProposedGovernance is a free data retrieval call binding the contract method 0x60f7ac97.
//
// Solidity: function proposedGovernance() view returns(address)
func (_PriceSubmitterBinding *PriceSubmitterBindingCallerSession) ProposedGovernance() (common.Address, error) {
	return _PriceSubmitterBinding.Contract.ProposedGovernance(&_PriceSubmitterBinding.CallOpts)
}

// VoterWhitelistBitmap is a free data retrieval call binding the contract method 0x7ac420ad.
//
// Solidity: function voterWhitelistBitmap(address _voter) view returns(uint256)
func (_PriceSubmitterBinding *PriceSubmitterBindingCaller) VoterWhitelistBitmap(opts *bind.CallOpts, _voter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceSubmitterBinding.contract.Call(opts, &out, "voterWhitelistBitmap", _voter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoterWhitelistBitmap is a free data retrieval call binding the contract method 0x7ac420ad.
//
// Solidity: function voterWhitelistBitmap(address _voter) view returns(uint256)
func (_PriceSubmitterBinding *PriceSubmitterBindingSession) VoterWhitelistBitmap(_voter common.Address) (*big.Int, error) {
	return _PriceSubmitterBinding.Contract.VoterWhitelistBitmap(&_PriceSubmitterBinding.CallOpts, _voter)
}

// VoterWhitelistBitmap is a free data retrieval call binding the contract method 0x7ac420ad.
//
// Solidity: function voterWhitelistBitmap(address _voter) view returns(uint256)
func (_PriceSubmitterBinding *PriceSubmitterBindingCallerSession) VoterWhitelistBitmap(_voter common.Address) (*big.Int, error) {
	return _PriceSubmitterBinding.Contract.VoterWhitelistBitmap(&_PriceSubmitterBinding.CallOpts, _voter)
}

// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
//
// Solidity: function claimGovernance() returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactor) ClaimGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceSubmitterBinding.contract.Transact(opts, "claimGovernance")
}

// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
//
// Solidity: function claimGovernance() returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingSession) ClaimGovernance() (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.ClaimGovernance(&_PriceSubmitterBinding.TransactOpts)
}

// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
//
// Solidity: function claimGovernance() returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactorSession) ClaimGovernance() (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.ClaimGovernance(&_PriceSubmitterBinding.TransactOpts)
}

// InitialiseFixedAddress is a paid mutator transaction binding the contract method 0xc9f960eb.
//
// Solidity: function initialiseFixedAddress() returns(address)
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactor) InitialiseFixedAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceSubmitterBinding.contract.Transact(opts, "initialiseFixedAddress")
}

// InitialiseFixedAddress is a paid mutator transaction binding the contract method 0xc9f960eb.
//
// Solidity: function initialiseFixedAddress() returns(address)
func (_PriceSubmitterBinding *PriceSubmitterBindingSession) InitialiseFixedAddress() (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.InitialiseFixedAddress(&_PriceSubmitterBinding.TransactOpts)
}

// InitialiseFixedAddress is a paid mutator transaction binding the contract method 0xc9f960eb.
//
// Solidity: function initialiseFixedAddress() returns(address)
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactorSession) InitialiseFixedAddress() (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.InitialiseFixedAddress(&_PriceSubmitterBinding.TransactOpts)
}

// ProposeGovernance is a paid mutator transaction binding the contract method 0xc373a08e.
//
// Solidity: function proposeGovernance(address _governance) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactor) ProposeGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _PriceSubmitterBinding.contract.Transact(opts, "proposeGovernance", _governance)
}

// ProposeGovernance is a paid mutator transaction binding the contract method 0xc373a08e.
//
// Solidity: function proposeGovernance(address _governance) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingSession) ProposeGovernance(_governance common.Address) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.ProposeGovernance(&_PriceSubmitterBinding.TransactOpts, _governance)
}

// ProposeGovernance is a paid mutator transaction binding the contract method 0xc373a08e.
//
// Solidity: function proposeGovernance(address _governance) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactorSession) ProposeGovernance(_governance common.Address) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.ProposeGovernance(&_PriceSubmitterBinding.TransactOpts, _governance)
}

// RevealPrices is a paid mutator transaction binding the contract method 0x60848b44.
//
// Solidity: function revealPrices(uint256 _epochId, uint256[] _ftsoIndices, uint256[] _prices, uint256[] _randoms) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactor) RevealPrices(opts *bind.TransactOpts, _epochId *big.Int, _ftsoIndices []*big.Int, _prices []*big.Int, _randoms []*big.Int) (*types.Transaction, error) {
	return _PriceSubmitterBinding.contract.Transact(opts, "revealPrices", _epochId, _ftsoIndices, _prices, _randoms)
}

// RevealPrices is a paid mutator transaction binding the contract method 0x60848b44.
//
// Solidity: function revealPrices(uint256 _epochId, uint256[] _ftsoIndices, uint256[] _prices, uint256[] _randoms) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingSession) RevealPrices(_epochId *big.Int, _ftsoIndices []*big.Int, _prices []*big.Int, _randoms []*big.Int) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.RevealPrices(&_PriceSubmitterBinding.TransactOpts, _epochId, _ftsoIndices, _prices, _randoms)
}

// RevealPrices is a paid mutator transaction binding the contract method 0x60848b44.
//
// Solidity: function revealPrices(uint256 _epochId, uint256[] _ftsoIndices, uint256[] _prices, uint256[] _randoms) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactorSession) RevealPrices(_epochId *big.Int, _ftsoIndices []*big.Int, _prices []*big.Int, _randoms []*big.Int) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.RevealPrices(&_PriceSubmitterBinding.TransactOpts, _epochId, _ftsoIndices, _prices, _randoms)
}

// SetContractAddresses is a paid mutator transaction binding the contract method 0x8ab63380.
//
// Solidity: function setContractAddresses(address _ftsoRegistry, address _voterWhitelister, address _ftsoManager) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactor) SetContractAddresses(opts *bind.TransactOpts, _ftsoRegistry common.Address, _voterWhitelister common.Address, _ftsoManager common.Address) (*types.Transaction, error) {
	return _PriceSubmitterBinding.contract.Transact(opts, "setContractAddresses", _ftsoRegistry, _voterWhitelister, _ftsoManager)
}

// SetContractAddresses is a paid mutator transaction binding the contract method 0x8ab63380.
//
// Solidity: function setContractAddresses(address _ftsoRegistry, address _voterWhitelister, address _ftsoManager) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingSession) SetContractAddresses(_ftsoRegistry common.Address, _voterWhitelister common.Address, _ftsoManager common.Address) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.SetContractAddresses(&_PriceSubmitterBinding.TransactOpts, _ftsoRegistry, _voterWhitelister, _ftsoManager)
}

// SetContractAddresses is a paid mutator transaction binding the contract method 0x8ab63380.
//
// Solidity: function setContractAddresses(address _ftsoRegistry, address _voterWhitelister, address _ftsoManager) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactorSession) SetContractAddresses(_ftsoRegistry common.Address, _voterWhitelister common.Address, _ftsoManager common.Address) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.SetContractAddresses(&_PriceSubmitterBinding.TransactOpts, _ftsoRegistry, _voterWhitelister, _ftsoManager)
}

// SetTrustedAddresses is a paid mutator transaction binding the contract method 0x9ec2b581.
//
// Solidity: function setTrustedAddresses(address[] _trustedAddresses) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactor) SetTrustedAddresses(opts *bind.TransactOpts, _trustedAddresses []common.Address) (*types.Transaction, error) {
	return _PriceSubmitterBinding.contract.Transact(opts, "setTrustedAddresses", _trustedAddresses)
}

// SetTrustedAddresses is a paid mutator transaction binding the contract method 0x9ec2b581.
//
// Solidity: function setTrustedAddresses(address[] _trustedAddresses) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingSession) SetTrustedAddresses(_trustedAddresses []common.Address) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.SetTrustedAddresses(&_PriceSubmitterBinding.TransactOpts, _trustedAddresses)
}

// SetTrustedAddresses is a paid mutator transaction binding the contract method 0x9ec2b581.
//
// Solidity: function setTrustedAddresses(address[] _trustedAddresses) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactorSession) SetTrustedAddresses(_trustedAddresses []common.Address) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.SetTrustedAddresses(&_PriceSubmitterBinding.TransactOpts, _trustedAddresses)
}

// SubmitPriceHashes is a paid mutator transaction binding the contract method 0xc5adc539.
//
// Solidity: function submitPriceHashes(uint256 _epochId, uint256[] _ftsoIndices, bytes32[] _hashes) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactor) SubmitPriceHashes(opts *bind.TransactOpts, _epochId *big.Int, _ftsoIndices []*big.Int, _hashes [][32]byte) (*types.Transaction, error) {
	return _PriceSubmitterBinding.contract.Transact(opts, "submitPriceHashes", _epochId, _ftsoIndices, _hashes)
}

// SubmitPriceHashes is a paid mutator transaction binding the contract method 0xc5adc539.
//
// Solidity: function submitPriceHashes(uint256 _epochId, uint256[] _ftsoIndices, bytes32[] _hashes) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingSession) SubmitPriceHashes(_epochId *big.Int, _ftsoIndices []*big.Int, _hashes [][32]byte) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.SubmitPriceHashes(&_PriceSubmitterBinding.TransactOpts, _epochId, _ftsoIndices, _hashes)
}

// SubmitPriceHashes is a paid mutator transaction binding the contract method 0xc5adc539.
//
// Solidity: function submitPriceHashes(uint256 _epochId, uint256[] _ftsoIndices, bytes32[] _hashes) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactorSession) SubmitPriceHashes(_epochId *big.Int, _ftsoIndices []*big.Int, _hashes [][32]byte) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.SubmitPriceHashes(&_PriceSubmitterBinding.TransactOpts, _epochId, _ftsoIndices, _hashes)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _governance) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactor) TransferGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _PriceSubmitterBinding.contract.Transact(opts, "transferGovernance", _governance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _governance) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingSession) TransferGovernance(_governance common.Address) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.TransferGovernance(&_PriceSubmitterBinding.TransactOpts, _governance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _governance) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactorSession) TransferGovernance(_governance common.Address) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.TransferGovernance(&_PriceSubmitterBinding.TransactOpts, _governance)
}

// VoterWhitelisted is a paid mutator transaction binding the contract method 0x9d986f91.
//
// Solidity: function voterWhitelisted(address _voter, uint256 _ftsoIndex) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactor) VoterWhitelisted(opts *bind.TransactOpts, _voter common.Address, _ftsoIndex *big.Int) (*types.Transaction, error) {
	return _PriceSubmitterBinding.contract.Transact(opts, "voterWhitelisted", _voter, _ftsoIndex)
}

// VoterWhitelisted is a paid mutator transaction binding the contract method 0x9d986f91.
//
// Solidity: function voterWhitelisted(address _voter, uint256 _ftsoIndex) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingSession) VoterWhitelisted(_voter common.Address, _ftsoIndex *big.Int) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.VoterWhitelisted(&_PriceSubmitterBinding.TransactOpts, _voter, _ftsoIndex)
}

// VoterWhitelisted is a paid mutator transaction binding the contract method 0x9d986f91.
//
// Solidity: function voterWhitelisted(address _voter, uint256 _ftsoIndex) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactorSession) VoterWhitelisted(_voter common.Address, _ftsoIndex *big.Int) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.VoterWhitelisted(&_PriceSubmitterBinding.TransactOpts, _voter, _ftsoIndex)
}

// VotersRemovedFromWhitelist is a paid mutator transaction binding the contract method 0x76794efb.
//
// Solidity: function votersRemovedFromWhitelist(address[] _removedVoters, uint256 _ftsoIndex) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactor) VotersRemovedFromWhitelist(opts *bind.TransactOpts, _removedVoters []common.Address, _ftsoIndex *big.Int) (*types.Transaction, error) {
	return _PriceSubmitterBinding.contract.Transact(opts, "votersRemovedFromWhitelist", _removedVoters, _ftsoIndex)
}

// VotersRemovedFromWhitelist is a paid mutator transaction binding the contract method 0x76794efb.
//
// Solidity: function votersRemovedFromWhitelist(address[] _removedVoters, uint256 _ftsoIndex) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingSession) VotersRemovedFromWhitelist(_removedVoters []common.Address, _ftsoIndex *big.Int) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.VotersRemovedFromWhitelist(&_PriceSubmitterBinding.TransactOpts, _removedVoters, _ftsoIndex)
}

// VotersRemovedFromWhitelist is a paid mutator transaction binding the contract method 0x76794efb.
//
// Solidity: function votersRemovedFromWhitelist(address[] _removedVoters, uint256 _ftsoIndex) returns()
func (_PriceSubmitterBinding *PriceSubmitterBindingTransactorSession) VotersRemovedFromWhitelist(_removedVoters []common.Address, _ftsoIndex *big.Int) (*types.Transaction, error) {
	return _PriceSubmitterBinding.Contract.VotersRemovedFromWhitelist(&_PriceSubmitterBinding.TransactOpts, _removedVoters, _ftsoIndex)
}

// PriceSubmitterBindingGovernanceProposedIterator is returned from FilterGovernanceProposed and is used to iterate over the raw logs and unpacked data for GovernanceProposed events raised by the PriceSubmitterBinding contract.
type PriceSubmitterBindingGovernanceProposedIterator struct {
	Event *PriceSubmitterBindingGovernanceProposed // Event containing the contract specifics and raw log

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
func (it *PriceSubmitterBindingGovernanceProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceSubmitterBindingGovernanceProposed)
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
		it.Event = new(PriceSubmitterBindingGovernanceProposed)
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
func (it *PriceSubmitterBindingGovernanceProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceSubmitterBindingGovernanceProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceSubmitterBindingGovernanceProposed represents a GovernanceProposed event raised by the PriceSubmitterBinding contract.
type PriceSubmitterBindingGovernanceProposed struct {
	ProposedGovernance common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceProposed is a free log retrieval operation binding the contract event 0x1f95fb40be3a947982072902a887b521248d1d8931a39eb38f84f4d6fd758b69.
//
// Solidity: event GovernanceProposed(address proposedGovernance)
func (_PriceSubmitterBinding *PriceSubmitterBindingFilterer) FilterGovernanceProposed(opts *bind.FilterOpts) (*PriceSubmitterBindingGovernanceProposedIterator, error) {

	logs, sub, err := _PriceSubmitterBinding.contract.FilterLogs(opts, "GovernanceProposed")
	if err != nil {
		return nil, err
	}
	return &PriceSubmitterBindingGovernanceProposedIterator{contract: _PriceSubmitterBinding.contract, event: "GovernanceProposed", logs: logs, sub: sub}, nil
}

// WatchGovernanceProposed is a free log subscription operation binding the contract event 0x1f95fb40be3a947982072902a887b521248d1d8931a39eb38f84f4d6fd758b69.
//
// Solidity: event GovernanceProposed(address proposedGovernance)
func (_PriceSubmitterBinding *PriceSubmitterBindingFilterer) WatchGovernanceProposed(opts *bind.WatchOpts, sink chan<- *PriceSubmitterBindingGovernanceProposed) (event.Subscription, error) {

	logs, sub, err := _PriceSubmitterBinding.contract.WatchLogs(opts, "GovernanceProposed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceSubmitterBindingGovernanceProposed)
				if err := _PriceSubmitterBinding.contract.UnpackLog(event, "GovernanceProposed", log); err != nil {
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

// ParseGovernanceProposed is a log parse operation binding the contract event 0x1f95fb40be3a947982072902a887b521248d1d8931a39eb38f84f4d6fd758b69.
//
// Solidity: event GovernanceProposed(address proposedGovernance)
func (_PriceSubmitterBinding *PriceSubmitterBindingFilterer) ParseGovernanceProposed(log types.Log) (*PriceSubmitterBindingGovernanceProposed, error) {
	event := new(PriceSubmitterBindingGovernanceProposed)
	if err := _PriceSubmitterBinding.contract.UnpackLog(event, "GovernanceProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceSubmitterBindingGovernanceUpdatedIterator is returned from FilterGovernanceUpdated and is used to iterate over the raw logs and unpacked data for GovernanceUpdated events raised by the PriceSubmitterBinding contract.
type PriceSubmitterBindingGovernanceUpdatedIterator struct {
	Event *PriceSubmitterBindingGovernanceUpdated // Event containing the contract specifics and raw log

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
func (it *PriceSubmitterBindingGovernanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceSubmitterBindingGovernanceUpdated)
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
		it.Event = new(PriceSubmitterBindingGovernanceUpdated)
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
func (it *PriceSubmitterBindingGovernanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceSubmitterBindingGovernanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceSubmitterBindingGovernanceUpdated represents a GovernanceUpdated event raised by the PriceSubmitterBinding contract.
type PriceSubmitterBindingGovernanceUpdated struct {
	OldGovernance  common.Address
	NewGoveranance common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovernanceUpdated is a free log retrieval operation binding the contract event 0x434a2db650703b36c824e745330d6397cdaa9ee2cc891a4938ae853e1c50b68d.
//
// Solidity: event GovernanceUpdated(address oldGovernance, address newGoveranance)
func (_PriceSubmitterBinding *PriceSubmitterBindingFilterer) FilterGovernanceUpdated(opts *bind.FilterOpts) (*PriceSubmitterBindingGovernanceUpdatedIterator, error) {

	logs, sub, err := _PriceSubmitterBinding.contract.FilterLogs(opts, "GovernanceUpdated")
	if err != nil {
		return nil, err
	}
	return &PriceSubmitterBindingGovernanceUpdatedIterator{contract: _PriceSubmitterBinding.contract, event: "GovernanceUpdated", logs: logs, sub: sub}, nil
}

// WatchGovernanceUpdated is a free log subscription operation binding the contract event 0x434a2db650703b36c824e745330d6397cdaa9ee2cc891a4938ae853e1c50b68d.
//
// Solidity: event GovernanceUpdated(address oldGovernance, address newGoveranance)
func (_PriceSubmitterBinding *PriceSubmitterBindingFilterer) WatchGovernanceUpdated(opts *bind.WatchOpts, sink chan<- *PriceSubmitterBindingGovernanceUpdated) (event.Subscription, error) {

	logs, sub, err := _PriceSubmitterBinding.contract.WatchLogs(opts, "GovernanceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceSubmitterBindingGovernanceUpdated)
				if err := _PriceSubmitterBinding.contract.UnpackLog(event, "GovernanceUpdated", log); err != nil {
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

// ParseGovernanceUpdated is a log parse operation binding the contract event 0x434a2db650703b36c824e745330d6397cdaa9ee2cc891a4938ae853e1c50b68d.
//
// Solidity: event GovernanceUpdated(address oldGovernance, address newGoveranance)
func (_PriceSubmitterBinding *PriceSubmitterBindingFilterer) ParseGovernanceUpdated(log types.Log) (*PriceSubmitterBindingGovernanceUpdated, error) {
	event := new(PriceSubmitterBindingGovernanceUpdated)
	if err := _PriceSubmitterBinding.contract.UnpackLog(event, "GovernanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceSubmitterBindingPriceHashesSubmittedIterator is returned from FilterPriceHashesSubmitted and is used to iterate over the raw logs and unpacked data for PriceHashesSubmitted events raised by the PriceSubmitterBinding contract.
type PriceSubmitterBindingPriceHashesSubmittedIterator struct {
	Event *PriceSubmitterBindingPriceHashesSubmitted // Event containing the contract specifics and raw log

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
func (it *PriceSubmitterBindingPriceHashesSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceSubmitterBindingPriceHashesSubmitted)
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
		it.Event = new(PriceSubmitterBindingPriceHashesSubmitted)
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
func (it *PriceSubmitterBindingPriceHashesSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceSubmitterBindingPriceHashesSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceSubmitterBindingPriceHashesSubmitted represents a PriceHashesSubmitted event raised by the PriceSubmitterBinding contract.
type PriceSubmitterBindingPriceHashesSubmitted struct {
	Submitter common.Address
	EpochId   *big.Int
	Ftsos     []common.Address
	Hashes    [][32]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceHashesSubmitted is a free log retrieval operation binding the contract event 0x90c022ade239639b1f8c4ebb8a76df5e03a7129df46cf9bcdae3c1450ea35434.
//
// Solidity: event PriceHashesSubmitted(address indexed submitter, uint256 indexed epochId, address[] ftsos, bytes32[] hashes, uint256 timestamp)
func (_PriceSubmitterBinding *PriceSubmitterBindingFilterer) FilterPriceHashesSubmitted(opts *bind.FilterOpts, submitter []common.Address, epochId []*big.Int) (*PriceSubmitterBindingPriceHashesSubmittedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _PriceSubmitterBinding.contract.FilterLogs(opts, "PriceHashesSubmitted", submitterRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &PriceSubmitterBindingPriceHashesSubmittedIterator{contract: _PriceSubmitterBinding.contract, event: "PriceHashesSubmitted", logs: logs, sub: sub}, nil
}

// WatchPriceHashesSubmitted is a free log subscription operation binding the contract event 0x90c022ade239639b1f8c4ebb8a76df5e03a7129df46cf9bcdae3c1450ea35434.
//
// Solidity: event PriceHashesSubmitted(address indexed submitter, uint256 indexed epochId, address[] ftsos, bytes32[] hashes, uint256 timestamp)
func (_PriceSubmitterBinding *PriceSubmitterBindingFilterer) WatchPriceHashesSubmitted(opts *bind.WatchOpts, sink chan<- *PriceSubmitterBindingPriceHashesSubmitted, submitter []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _PriceSubmitterBinding.contract.WatchLogs(opts, "PriceHashesSubmitted", submitterRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceSubmitterBindingPriceHashesSubmitted)
				if err := _PriceSubmitterBinding.contract.UnpackLog(event, "PriceHashesSubmitted", log); err != nil {
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

// ParsePriceHashesSubmitted is a log parse operation binding the contract event 0x90c022ade239639b1f8c4ebb8a76df5e03a7129df46cf9bcdae3c1450ea35434.
//
// Solidity: event PriceHashesSubmitted(address indexed submitter, uint256 indexed epochId, address[] ftsos, bytes32[] hashes, uint256 timestamp)
func (_PriceSubmitterBinding *PriceSubmitterBindingFilterer) ParsePriceHashesSubmitted(log types.Log) (*PriceSubmitterBindingPriceHashesSubmitted, error) {
	event := new(PriceSubmitterBindingPriceHashesSubmitted)
	if err := _PriceSubmitterBinding.contract.UnpackLog(event, "PriceHashesSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceSubmitterBindingPricesRevealedIterator is returned from FilterPricesRevealed and is used to iterate over the raw logs and unpacked data for PricesRevealed events raised by the PriceSubmitterBinding contract.
type PriceSubmitterBindingPricesRevealedIterator struct {
	Event *PriceSubmitterBindingPricesRevealed // Event containing the contract specifics and raw log

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
func (it *PriceSubmitterBindingPricesRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceSubmitterBindingPricesRevealed)
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
		it.Event = new(PriceSubmitterBindingPricesRevealed)
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
func (it *PriceSubmitterBindingPricesRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceSubmitterBindingPricesRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceSubmitterBindingPricesRevealed represents a PricesRevealed event raised by the PriceSubmitterBinding contract.
type PriceSubmitterBindingPricesRevealed struct {
	Voter     common.Address
	EpochId   *big.Int
	Ftsos     []common.Address
	Prices    []*big.Int
	Randoms   []*big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPricesRevealed is a free log retrieval operation binding the contract event 0xa32444a31df2f9a116229eec3193d223a6bad89f7670ff17b8e5c7014a377da1.
//
// Solidity: event PricesRevealed(address indexed voter, uint256 indexed epochId, address[] ftsos, uint256[] prices, uint256[] randoms, uint256 timestamp)
func (_PriceSubmitterBinding *PriceSubmitterBindingFilterer) FilterPricesRevealed(opts *bind.FilterOpts, voter []common.Address, epochId []*big.Int) (*PriceSubmitterBindingPricesRevealedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _PriceSubmitterBinding.contract.FilterLogs(opts, "PricesRevealed", voterRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &PriceSubmitterBindingPricesRevealedIterator{contract: _PriceSubmitterBinding.contract, event: "PricesRevealed", logs: logs, sub: sub}, nil
}

// WatchPricesRevealed is a free log subscription operation binding the contract event 0xa32444a31df2f9a116229eec3193d223a6bad89f7670ff17b8e5c7014a377da1.
//
// Solidity: event PricesRevealed(address indexed voter, uint256 indexed epochId, address[] ftsos, uint256[] prices, uint256[] randoms, uint256 timestamp)
func (_PriceSubmitterBinding *PriceSubmitterBindingFilterer) WatchPricesRevealed(opts *bind.WatchOpts, sink chan<- *PriceSubmitterBindingPricesRevealed, voter []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _PriceSubmitterBinding.contract.WatchLogs(opts, "PricesRevealed", voterRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceSubmitterBindingPricesRevealed)
				if err := _PriceSubmitterBinding.contract.UnpackLog(event, "PricesRevealed", log); err != nil {
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

// ParsePricesRevealed is a log parse operation binding the contract event 0xa32444a31df2f9a116229eec3193d223a6bad89f7670ff17b8e5c7014a377da1.
//
// Solidity: event PricesRevealed(address indexed voter, uint256 indexed epochId, address[] ftsos, uint256[] prices, uint256[] randoms, uint256 timestamp)
func (_PriceSubmitterBinding *PriceSubmitterBindingFilterer) ParsePricesRevealed(log types.Log) (*PriceSubmitterBindingPricesRevealed, error) {
	event := new(PriceSubmitterBindingPricesRevealed)
	if err := _PriceSubmitterBinding.contract.UnpackLog(event, "PricesRevealed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

