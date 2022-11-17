// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gateway_registry

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

// IGatewayRegistryGateway is an auto generated low-level Go binding around an user-defined struct.
type IGatewayRegistryGateway struct {
	PublicKey     [32]byte
	Version       uint8
	Owner         common.Address
	AntennaGain   uint8
	FrequencyPlan uint8
	Location      int64
	Altitude      uint8
}

// GatewayRegistryMetaData contains all meta data concerning the GatewayRegistry contract.
var GatewayRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"uint8[]\",\"name\":\"fps\",\"type\":\"uint8[]\"},{\"internalType\":\"string[]\",\"name\":\"fpNames\",\"type\":\"string[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"FrequencyPlanAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"FrequencyPlanRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"gatewayId\",\"type\":\"bytes32\"}],\"name\":\"GatewayOffboarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"gatewayId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"GatewayOnboarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"onboards\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"offboards\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"updates\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"transfers\",\"type\":\"bool\"}],\"name\":\"GatewayRegistryPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"gatewayId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"GatewayTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"gatewayId\",\"type\":\"bytes32\"}],\"name\":\"GatewayUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GATEWAY_OFFBOARD_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GATEWAY_ONBOARDER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GATEWAY_TRANSFER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GATEWAY_UPDATER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRUSTED_FORWARDER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"fp\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"addFrequencyPlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"frequencyPlans\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"gatewayCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"gatewayOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"gateways\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"publicKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"antennaGain\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"frequencyPlan\",\"type\":\"uint8\"},{\"internalType\":\"int64\",\"name\":\"location\",\"type\":\"int64\"},{\"internalType\":\"uint8\",\"name\":\"altitude\",\"type\":\"uint8\"}],\"internalType\":\"structIGatewayRegistry.Gateway\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"gatewaysPaged\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"publicKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"antennaGain\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"frequencyPlan\",\"type\":\"uint8\"},{\"internalType\":\"int64\",\"name\":\"location\",\"type\":\"int64\"},{\"internalType\":\"uint8\",\"name\":\"altitude\",\"type\":\"uint8\"}],\"internalType\":\"structIGatewayRegistry.Gateway[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"gatewayId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"offboard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"gatewayId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"onboard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registeredGateways\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"fp\",\"type\":\"uint8\"}],\"name\":\"removeFrequencyPlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"onboards\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"offboards\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"updates\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"transfers\",\"type\":\"bool\"}],\"name\":\"setPauseStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"gatewayId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"currentOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"gatewayId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"antennaGain\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"frequencyPlan\",\"type\":\"uint8\"},{\"internalType\":\"int64\",\"name\":\"location\",\"type\":\"int64\"},{\"internalType\":\"uint8\",\"name\":\"altitude\",\"type\":\"uint8\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GatewayRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayRegistryMetaData.ABI instead.
var GatewayRegistryABI = GatewayRegistryMetaData.ABI

// GatewayRegistry is an auto generated Go binding around an Ethereum contract.
type GatewayRegistry struct {
	GatewayRegistryCaller     // Read-only binding to the contract
	GatewayRegistryTransactor // Write-only binding to the contract
	GatewayRegistryFilterer   // Log filterer for contract events
}

// GatewayRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayRegistrySession struct {
	Contract     *GatewayRegistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayRegistryCallerSession struct {
	Contract *GatewayRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// GatewayRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayRegistryTransactorSession struct {
	Contract     *GatewayRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// GatewayRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayRegistryRaw struct {
	Contract *GatewayRegistry // Generic contract binding to access the raw methods on
}

// GatewayRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayRegistryCallerRaw struct {
	Contract *GatewayRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayRegistryTransactorRaw struct {
	Contract *GatewayRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayRegistry creates a new instance of GatewayRegistry, bound to a specific deployed contract.
func NewGatewayRegistry(address common.Address, backend bind.ContractBackend) (*GatewayRegistry, error) {
	contract, err := bindGatewayRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayRegistry{GatewayRegistryCaller: GatewayRegistryCaller{contract: contract}, GatewayRegistryTransactor: GatewayRegistryTransactor{contract: contract}, GatewayRegistryFilterer: GatewayRegistryFilterer{contract: contract}}, nil
}

// NewGatewayRegistryCaller creates a new read-only instance of GatewayRegistry, bound to a specific deployed contract.
func NewGatewayRegistryCaller(address common.Address, caller bind.ContractCaller) (*GatewayRegistryCaller, error) {
	contract, err := bindGatewayRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayRegistryCaller{contract: contract}, nil
}

// NewGatewayRegistryTransactor creates a new write-only instance of GatewayRegistry, bound to a specific deployed contract.
func NewGatewayRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayRegistryTransactor, error) {
	contract, err := bindGatewayRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayRegistryTransactor{contract: contract}, nil
}

// NewGatewayRegistryFilterer creates a new log filterer instance of GatewayRegistry, bound to a specific deployed contract.
func NewGatewayRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayRegistryFilterer, error) {
	contract, err := bindGatewayRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayRegistryFilterer{contract: contract}, nil
}

// bindGatewayRegistry binds a generic wrapper to an already deployed contract.
func bindGatewayRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayRegistry *GatewayRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayRegistry.Contract.GatewayRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayRegistry *GatewayRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.GatewayRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayRegistry *GatewayRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.GatewayRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayRegistry *GatewayRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayRegistry *GatewayRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayRegistry *GatewayRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayRegistry *GatewayRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayRegistry *GatewayRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayRegistry.Contract.DEFAULTADMINROLE(&_GatewayRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayRegistry *GatewayRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayRegistry.Contract.DEFAULTADMINROLE(&_GatewayRegistry.CallOpts)
}

// GATEWAYOFFBOARDROLE is a free data retrieval call binding the contract method 0xd9044941.
//
// Solidity: function GATEWAY_OFFBOARD_ROLE() view returns(bytes32)
func (_GatewayRegistry *GatewayRegistryCaller) GATEWAYOFFBOARDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayRegistry.contract.Call(opts, &out, "GATEWAY_OFFBOARD_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GATEWAYOFFBOARDROLE is a free data retrieval call binding the contract method 0xd9044941.
//
// Solidity: function GATEWAY_OFFBOARD_ROLE() view returns(bytes32)
func (_GatewayRegistry *GatewayRegistrySession) GATEWAYOFFBOARDROLE() ([32]byte, error) {
	return _GatewayRegistry.Contract.GATEWAYOFFBOARDROLE(&_GatewayRegistry.CallOpts)
}

// GATEWAYOFFBOARDROLE is a free data retrieval call binding the contract method 0xd9044941.
//
// Solidity: function GATEWAY_OFFBOARD_ROLE() view returns(bytes32)
func (_GatewayRegistry *GatewayRegistryCallerSession) GATEWAYOFFBOARDROLE() ([32]byte, error) {
	return _GatewayRegistry.Contract.GATEWAYOFFBOARDROLE(&_GatewayRegistry.CallOpts)
}

// GATEWAYONBOARDERROLE is a free data retrieval call binding the contract method 0x473682c3.
//
// Solidity: function GATEWAY_ONBOARDER_ROLE() view returns(bytes32)
func (_GatewayRegistry *GatewayRegistryCaller) GATEWAYONBOARDERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayRegistry.contract.Call(opts, &out, "GATEWAY_ONBOARDER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GATEWAYONBOARDERROLE is a free data retrieval call binding the contract method 0x473682c3.
//
// Solidity: function GATEWAY_ONBOARDER_ROLE() view returns(bytes32)
func (_GatewayRegistry *GatewayRegistrySession) GATEWAYONBOARDERROLE() ([32]byte, error) {
	return _GatewayRegistry.Contract.GATEWAYONBOARDERROLE(&_GatewayRegistry.CallOpts)
}

// GATEWAYONBOARDERROLE is a free data retrieval call binding the contract method 0x473682c3.
//
// Solidity: function GATEWAY_ONBOARDER_ROLE() view returns(bytes32)
func (_GatewayRegistry *GatewayRegistryCallerSession) GATEWAYONBOARDERROLE() ([32]byte, error) {
	return _GatewayRegistry.Contract.GATEWAYONBOARDERROLE(&_GatewayRegistry.CallOpts)
}

// GATEWAYTRANSFERROLE is a free data retrieval call binding the contract method 0x569b9c5f.
//
// Solidity: function GATEWAY_TRANSFER_ROLE() view returns(bytes32)
func (_GatewayRegistry *GatewayRegistryCaller) GATEWAYTRANSFERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayRegistry.contract.Call(opts, &out, "GATEWAY_TRANSFER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GATEWAYTRANSFERROLE is a free data retrieval call binding the contract method 0x569b9c5f.
//
// Solidity: function GATEWAY_TRANSFER_ROLE() view returns(bytes32)
func (_GatewayRegistry *GatewayRegistrySession) GATEWAYTRANSFERROLE() ([32]byte, error) {
	return _GatewayRegistry.Contract.GATEWAYTRANSFERROLE(&_GatewayRegistry.CallOpts)
}

// GATEWAYTRANSFERROLE is a free data retrieval call binding the contract method 0x569b9c5f.
//
// Solidity: function GATEWAY_TRANSFER_ROLE() view returns(bytes32)
func (_GatewayRegistry *GatewayRegistryCallerSession) GATEWAYTRANSFERROLE() ([32]byte, error) {
	return _GatewayRegistry.Contract.GATEWAYTRANSFERROLE(&_GatewayRegistry.CallOpts)
}

// GATEWAYUPDATERROLE is a free data retrieval call binding the contract method 0xff1d8d5f.
//
// Solidity: function GATEWAY_UPDATER_ROLE() view returns(bytes32)
func (_GatewayRegistry *GatewayRegistryCaller) GATEWAYUPDATERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayRegistry.contract.Call(opts, &out, "GATEWAY_UPDATER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GATEWAYUPDATERROLE is a free data retrieval call binding the contract method 0xff1d8d5f.
//
// Solidity: function GATEWAY_UPDATER_ROLE() view returns(bytes32)
func (_GatewayRegistry *GatewayRegistrySession) GATEWAYUPDATERROLE() ([32]byte, error) {
	return _GatewayRegistry.Contract.GATEWAYUPDATERROLE(&_GatewayRegistry.CallOpts)
}

// GATEWAYUPDATERROLE is a free data retrieval call binding the contract method 0xff1d8d5f.
//
// Solidity: function GATEWAY_UPDATER_ROLE() view returns(bytes32)
func (_GatewayRegistry *GatewayRegistryCallerSession) GATEWAYUPDATERROLE() ([32]byte, error) {
	return _GatewayRegistry.Contract.GATEWAYUPDATERROLE(&_GatewayRegistry.CallOpts)
}

// TRUSTEDFORWARDERROLE is a free data retrieval call binding the contract method 0x00cba943.
//
// Solidity: function TRUSTED_FORWARDER_ROLE() view returns(bytes32)
func (_GatewayRegistry *GatewayRegistryCaller) TRUSTEDFORWARDERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayRegistry.contract.Call(opts, &out, "TRUSTED_FORWARDER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRUSTEDFORWARDERROLE is a free data retrieval call binding the contract method 0x00cba943.
//
// Solidity: function TRUSTED_FORWARDER_ROLE() view returns(bytes32)
func (_GatewayRegistry *GatewayRegistrySession) TRUSTEDFORWARDERROLE() ([32]byte, error) {
	return _GatewayRegistry.Contract.TRUSTEDFORWARDERROLE(&_GatewayRegistry.CallOpts)
}

// TRUSTEDFORWARDERROLE is a free data retrieval call binding the contract method 0x00cba943.
//
// Solidity: function TRUSTED_FORWARDER_ROLE() view returns(bytes32)
func (_GatewayRegistry *GatewayRegistryCallerSession) TRUSTEDFORWARDERROLE() ([32]byte, error) {
	return _GatewayRegistry.Contract.TRUSTEDFORWARDERROLE(&_GatewayRegistry.CallOpts)
}

// FrequencyPlans is a free data retrieval call binding the contract method 0x809ae202.
//
// Solidity: function frequencyPlans(uint8 ) view returns(string)
func (_GatewayRegistry *GatewayRegistryCaller) FrequencyPlans(opts *bind.CallOpts, arg0 uint8) (string, error) {
	var out []interface{}
	err := _GatewayRegistry.contract.Call(opts, &out, "frequencyPlans", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// FrequencyPlans is a free data retrieval call binding the contract method 0x809ae202.
//
// Solidity: function frequencyPlans(uint8 ) view returns(string)
func (_GatewayRegistry *GatewayRegistrySession) FrequencyPlans(arg0 uint8) (string, error) {
	return _GatewayRegistry.Contract.FrequencyPlans(&_GatewayRegistry.CallOpts, arg0)
}

// FrequencyPlans is a free data retrieval call binding the contract method 0x809ae202.
//
// Solidity: function frequencyPlans(uint8 ) view returns(string)
func (_GatewayRegistry *GatewayRegistryCallerSession) FrequencyPlans(arg0 uint8) (string, error) {
	return _GatewayRegistry.Contract.FrequencyPlans(&_GatewayRegistry.CallOpts, arg0)
}

// GatewayCount is a free data retrieval call binding the contract method 0xbe1e5aa5.
//
// Solidity: function gatewayCount(address owner) view returns(uint256)
func (_GatewayRegistry *GatewayRegistryCaller) GatewayCount(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GatewayRegistry.contract.Call(opts, &out, "gatewayCount", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GatewayCount is a free data retrieval call binding the contract method 0xbe1e5aa5.
//
// Solidity: function gatewayCount(address owner) view returns(uint256)
func (_GatewayRegistry *GatewayRegistrySession) GatewayCount(owner common.Address) (*big.Int, error) {
	return _GatewayRegistry.Contract.GatewayCount(&_GatewayRegistry.CallOpts, owner)
}

// GatewayCount is a free data retrieval call binding the contract method 0xbe1e5aa5.
//
// Solidity: function gatewayCount(address owner) view returns(uint256)
func (_GatewayRegistry *GatewayRegistryCallerSession) GatewayCount(owner common.Address) (*big.Int, error) {
	return _GatewayRegistry.Contract.GatewayCount(&_GatewayRegistry.CallOpts, owner)
}

// GatewayOwner is a free data retrieval call binding the contract method 0x78c8ae5f.
//
// Solidity: function gatewayOwner(bytes32 id) view returns(address)
func (_GatewayRegistry *GatewayRegistryCaller) GatewayOwner(opts *bind.CallOpts, id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _GatewayRegistry.contract.Call(opts, &out, "gatewayOwner", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GatewayOwner is a free data retrieval call binding the contract method 0x78c8ae5f.
//
// Solidity: function gatewayOwner(bytes32 id) view returns(address)
func (_GatewayRegistry *GatewayRegistrySession) GatewayOwner(id [32]byte) (common.Address, error) {
	return _GatewayRegistry.Contract.GatewayOwner(&_GatewayRegistry.CallOpts, id)
}

// GatewayOwner is a free data retrieval call binding the contract method 0x78c8ae5f.
//
// Solidity: function gatewayOwner(bytes32 id) view returns(address)
func (_GatewayRegistry *GatewayRegistryCallerSession) GatewayOwner(id [32]byte) (common.Address, error) {
	return _GatewayRegistry.Contract.GatewayOwner(&_GatewayRegistry.CallOpts, id)
}

// Gateways is a free data retrieval call binding the contract method 0xfbe336ff.
//
// Solidity: function gateways(bytes32 id) view returns((bytes32,uint8,address,uint8,uint8,int64,uint8))
func (_GatewayRegistry *GatewayRegistryCaller) Gateways(opts *bind.CallOpts, id [32]byte) (IGatewayRegistryGateway, error) {
	var out []interface{}
	err := _GatewayRegistry.contract.Call(opts, &out, "gateways", id)

	if err != nil {
		return *new(IGatewayRegistryGateway), err
	}

	out0 := *abi.ConvertType(out[0], new(IGatewayRegistryGateway)).(*IGatewayRegistryGateway)

	return out0, err

}

// Gateways is a free data retrieval call binding the contract method 0xfbe336ff.
//
// Solidity: function gateways(bytes32 id) view returns((bytes32,uint8,address,uint8,uint8,int64,uint8))
func (_GatewayRegistry *GatewayRegistrySession) Gateways(id [32]byte) (IGatewayRegistryGateway, error) {
	return _GatewayRegistry.Contract.Gateways(&_GatewayRegistry.CallOpts, id)
}

// Gateways is a free data retrieval call binding the contract method 0xfbe336ff.
//
// Solidity: function gateways(bytes32 id) view returns((bytes32,uint8,address,uint8,uint8,int64,uint8))
func (_GatewayRegistry *GatewayRegistryCallerSession) Gateways(id [32]byte) (IGatewayRegistryGateway, error) {
	return _GatewayRegistry.Contract.Gateways(&_GatewayRegistry.CallOpts, id)
}

// GatewaysPaged is a free data retrieval call binding the contract method 0xd3fa7495.
//
// Solidity: function gatewaysPaged(address owner, uint256 start, uint256 end) view returns((bytes32,uint8,address,uint8,uint8,int64,uint8)[])
func (_GatewayRegistry *GatewayRegistryCaller) GatewaysPaged(opts *bind.CallOpts, owner common.Address, start *big.Int, end *big.Int) ([]IGatewayRegistryGateway, error) {
	var out []interface{}
	err := _GatewayRegistry.contract.Call(opts, &out, "gatewaysPaged", owner, start, end)

	if err != nil {
		return *new([]IGatewayRegistryGateway), err
	}

	out0 := *abi.ConvertType(out[0], new([]IGatewayRegistryGateway)).(*[]IGatewayRegistryGateway)

	return out0, err

}

// GatewaysPaged is a free data retrieval call binding the contract method 0xd3fa7495.
//
// Solidity: function gatewaysPaged(address owner, uint256 start, uint256 end) view returns((bytes32,uint8,address,uint8,uint8,int64,uint8)[])
func (_GatewayRegistry *GatewayRegistrySession) GatewaysPaged(owner common.Address, start *big.Int, end *big.Int) ([]IGatewayRegistryGateway, error) {
	return _GatewayRegistry.Contract.GatewaysPaged(&_GatewayRegistry.CallOpts, owner, start, end)
}

// GatewaysPaged is a free data retrieval call binding the contract method 0xd3fa7495.
//
// Solidity: function gatewaysPaged(address owner, uint256 start, uint256 end) view returns((bytes32,uint8,address,uint8,uint8,int64,uint8)[])
func (_GatewayRegistry *GatewayRegistryCallerSession) GatewaysPaged(owner common.Address, start *big.Int, end *big.Int) ([]IGatewayRegistryGateway, error) {
	return _GatewayRegistry.Contract.GatewaysPaged(&_GatewayRegistry.CallOpts, owner, start, end)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayRegistry *GatewayRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GatewayRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayRegistry *GatewayRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GatewayRegistry.Contract.GetRoleAdmin(&_GatewayRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayRegistry *GatewayRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GatewayRegistry.Contract.GetRoleAdmin(&_GatewayRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayRegistry *GatewayRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _GatewayRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayRegistry *GatewayRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GatewayRegistry.Contract.HasRole(&_GatewayRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayRegistry *GatewayRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GatewayRegistry.Contract.HasRole(&_GatewayRegistry.CallOpts, role, account)
}

// RegisteredGateways is a free data retrieval call binding the contract method 0x06543c1a.
//
// Solidity: function registeredGateways() view returns(uint256)
func (_GatewayRegistry *GatewayRegistryCaller) RegisteredGateways(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayRegistry.contract.Call(opts, &out, "registeredGateways")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegisteredGateways is a free data retrieval call binding the contract method 0x06543c1a.
//
// Solidity: function registeredGateways() view returns(uint256)
func (_GatewayRegistry *GatewayRegistrySession) RegisteredGateways() (*big.Int, error) {
	return _GatewayRegistry.Contract.RegisteredGateways(&_GatewayRegistry.CallOpts)
}

// RegisteredGateways is a free data retrieval call binding the contract method 0x06543c1a.
//
// Solidity: function registeredGateways() view returns(uint256)
func (_GatewayRegistry *GatewayRegistryCallerSession) RegisteredGateways() (*big.Int, error) {
	return _GatewayRegistry.Contract.RegisteredGateways(&_GatewayRegistry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayRegistry *GatewayRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GatewayRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayRegistry *GatewayRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GatewayRegistry.Contract.SupportsInterface(&_GatewayRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayRegistry *GatewayRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GatewayRegistry.Contract.SupportsInterface(&_GatewayRegistry.CallOpts, interfaceId)
}

// AddFrequencyPlan is a paid mutator transaction binding the contract method 0x640335c7.
//
// Solidity: function addFrequencyPlan(uint8 fp, string name) returns()
func (_GatewayRegistry *GatewayRegistryTransactor) AddFrequencyPlan(opts *bind.TransactOpts, fp uint8, name string) (*types.Transaction, error) {
	return _GatewayRegistry.contract.Transact(opts, "addFrequencyPlan", fp, name)
}

// AddFrequencyPlan is a paid mutator transaction binding the contract method 0x640335c7.
//
// Solidity: function addFrequencyPlan(uint8 fp, string name) returns()
func (_GatewayRegistry *GatewayRegistrySession) AddFrequencyPlan(fp uint8, name string) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.AddFrequencyPlan(&_GatewayRegistry.TransactOpts, fp, name)
}

// AddFrequencyPlan is a paid mutator transaction binding the contract method 0x640335c7.
//
// Solidity: function addFrequencyPlan(uint8 fp, string name) returns()
func (_GatewayRegistry *GatewayRegistryTransactorSession) AddFrequencyPlan(fp uint8, name string) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.AddFrequencyPlan(&_GatewayRegistry.TransactOpts, fp, name)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address beneficiary) returns()
func (_GatewayRegistry *GatewayRegistryTransactor) Destroy(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.contract.Transact(opts, "destroy", beneficiary)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address beneficiary) returns()
func (_GatewayRegistry *GatewayRegistrySession) Destroy(beneficiary common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.Destroy(&_GatewayRegistry.TransactOpts, beneficiary)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address beneficiary) returns()
func (_GatewayRegistry *GatewayRegistryTransactorSession) Destroy(beneficiary common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.Destroy(&_GatewayRegistry.TransactOpts, beneficiary)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayRegistry *GatewayRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayRegistry *GatewayRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.GrantRole(&_GatewayRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayRegistry *GatewayRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.GrantRole(&_GatewayRegistry.TransactOpts, role, account)
}

// Offboard is a paid mutator transaction binding the contract method 0x55bb316c.
//
// Solidity: function offboard(bytes32 gatewayId, address owner) returns()
func (_GatewayRegistry *GatewayRegistryTransactor) Offboard(opts *bind.TransactOpts, gatewayId [32]byte, owner common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.contract.Transact(opts, "offboard", gatewayId, owner)
}

// Offboard is a paid mutator transaction binding the contract method 0x55bb316c.
//
// Solidity: function offboard(bytes32 gatewayId, address owner) returns()
func (_GatewayRegistry *GatewayRegistrySession) Offboard(gatewayId [32]byte, owner common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.Offboard(&_GatewayRegistry.TransactOpts, gatewayId, owner)
}

// Offboard is a paid mutator transaction binding the contract method 0x55bb316c.
//
// Solidity: function offboard(bytes32 gatewayId, address owner) returns()
func (_GatewayRegistry *GatewayRegistryTransactorSession) Offboard(gatewayId [32]byte, owner common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.Offboard(&_GatewayRegistry.TransactOpts, gatewayId, owner)
}

// Onboard is a paid mutator transaction binding the contract method 0x5a5bc82b.
//
// Solidity: function onboard(bytes32 gatewayId, uint8 version, address owner) returns()
func (_GatewayRegistry *GatewayRegistryTransactor) Onboard(opts *bind.TransactOpts, gatewayId [32]byte, version uint8, owner common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.contract.Transact(opts, "onboard", gatewayId, version, owner)
}

// Onboard is a paid mutator transaction binding the contract method 0x5a5bc82b.
//
// Solidity: function onboard(bytes32 gatewayId, uint8 version, address owner) returns()
func (_GatewayRegistry *GatewayRegistrySession) Onboard(gatewayId [32]byte, version uint8, owner common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.Onboard(&_GatewayRegistry.TransactOpts, gatewayId, version, owner)
}

// Onboard is a paid mutator transaction binding the contract method 0x5a5bc82b.
//
// Solidity: function onboard(bytes32 gatewayId, uint8 version, address owner) returns()
func (_GatewayRegistry *GatewayRegistryTransactorSession) Onboard(gatewayId [32]byte, version uint8, owner common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.Onboard(&_GatewayRegistry.TransactOpts, gatewayId, version, owner)
}

// RemoveFrequencyPlan is a paid mutator transaction binding the contract method 0x893729b3.
//
// Solidity: function removeFrequencyPlan(uint8 fp) returns()
func (_GatewayRegistry *GatewayRegistryTransactor) RemoveFrequencyPlan(opts *bind.TransactOpts, fp uint8) (*types.Transaction, error) {
	return _GatewayRegistry.contract.Transact(opts, "removeFrequencyPlan", fp)
}

// RemoveFrequencyPlan is a paid mutator transaction binding the contract method 0x893729b3.
//
// Solidity: function removeFrequencyPlan(uint8 fp) returns()
func (_GatewayRegistry *GatewayRegistrySession) RemoveFrequencyPlan(fp uint8) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.RemoveFrequencyPlan(&_GatewayRegistry.TransactOpts, fp)
}

// RemoveFrequencyPlan is a paid mutator transaction binding the contract method 0x893729b3.
//
// Solidity: function removeFrequencyPlan(uint8 fp) returns()
func (_GatewayRegistry *GatewayRegistryTransactorSession) RemoveFrequencyPlan(fp uint8) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.RemoveFrequencyPlan(&_GatewayRegistry.TransactOpts, fp)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GatewayRegistry *GatewayRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GatewayRegistry *GatewayRegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.RenounceRole(&_GatewayRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GatewayRegistry *GatewayRegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.RenounceRole(&_GatewayRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayRegistry *GatewayRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayRegistry *GatewayRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.RevokeRole(&_GatewayRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayRegistry *GatewayRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.RevokeRole(&_GatewayRegistry.TransactOpts, role, account)
}

// SetPauseStatus is a paid mutator transaction binding the contract method 0x1c6fd6ab.
//
// Solidity: function setPauseStatus(bool onboards, bool offboards, bool updates, bool transfers) returns()
func (_GatewayRegistry *GatewayRegistryTransactor) SetPauseStatus(opts *bind.TransactOpts, onboards bool, offboards bool, updates bool, transfers bool) (*types.Transaction, error) {
	return _GatewayRegistry.contract.Transact(opts, "setPauseStatus", onboards, offboards, updates, transfers)
}

// SetPauseStatus is a paid mutator transaction binding the contract method 0x1c6fd6ab.
//
// Solidity: function setPauseStatus(bool onboards, bool offboards, bool updates, bool transfers) returns()
func (_GatewayRegistry *GatewayRegistrySession) SetPauseStatus(onboards bool, offboards bool, updates bool, transfers bool) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.SetPauseStatus(&_GatewayRegistry.TransactOpts, onboards, offboards, updates, transfers)
}

// SetPauseStatus is a paid mutator transaction binding the contract method 0x1c6fd6ab.
//
// Solidity: function setPauseStatus(bool onboards, bool offboards, bool updates, bool transfers) returns()
func (_GatewayRegistry *GatewayRegistryTransactorSession) SetPauseStatus(onboards bool, offboards bool, updates bool, transfers bool) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.SetPauseStatus(&_GatewayRegistry.TransactOpts, onboards, offboards, updates, transfers)
}

// Transfer is a paid mutator transaction binding the contract method 0x3d6028ca.
//
// Solidity: function transfer(bytes32 gatewayId, address currentOwner, address newOwner) returns()
func (_GatewayRegistry *GatewayRegistryTransactor) Transfer(opts *bind.TransactOpts, gatewayId [32]byte, currentOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.contract.Transact(opts, "transfer", gatewayId, currentOwner, newOwner)
}

// Transfer is a paid mutator transaction binding the contract method 0x3d6028ca.
//
// Solidity: function transfer(bytes32 gatewayId, address currentOwner, address newOwner) returns()
func (_GatewayRegistry *GatewayRegistrySession) Transfer(gatewayId [32]byte, currentOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.Transfer(&_GatewayRegistry.TransactOpts, gatewayId, currentOwner, newOwner)
}

// Transfer is a paid mutator transaction binding the contract method 0x3d6028ca.
//
// Solidity: function transfer(bytes32 gatewayId, address currentOwner, address newOwner) returns()
func (_GatewayRegistry *GatewayRegistryTransactorSession) Transfer(gatewayId [32]byte, currentOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.Transfer(&_GatewayRegistry.TransactOpts, gatewayId, currentOwner, newOwner)
}

// Update is a paid mutator transaction binding the contract method 0x5f7bca1c.
//
// Solidity: function update(bytes32 gatewayId, uint8 version, address owner, uint8 antennaGain, uint8 frequencyPlan, int64 location, uint8 altitude) returns()
func (_GatewayRegistry *GatewayRegistryTransactor) Update(opts *bind.TransactOpts, gatewayId [32]byte, version uint8, owner common.Address, antennaGain uint8, frequencyPlan uint8, location int64, altitude uint8) (*types.Transaction, error) {
	return _GatewayRegistry.contract.Transact(opts, "update", gatewayId, version, owner, antennaGain, frequencyPlan, location, altitude)
}

// Update is a paid mutator transaction binding the contract method 0x5f7bca1c.
//
// Solidity: function update(bytes32 gatewayId, uint8 version, address owner, uint8 antennaGain, uint8 frequencyPlan, int64 location, uint8 altitude) returns()
func (_GatewayRegistry *GatewayRegistrySession) Update(gatewayId [32]byte, version uint8, owner common.Address, antennaGain uint8, frequencyPlan uint8, location int64, altitude uint8) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.Update(&_GatewayRegistry.TransactOpts, gatewayId, version, owner, antennaGain, frequencyPlan, location, altitude)
}

// Update is a paid mutator transaction binding the contract method 0x5f7bca1c.
//
// Solidity: function update(bytes32 gatewayId, uint8 version, address owner, uint8 antennaGain, uint8 frequencyPlan, int64 location, uint8 altitude) returns()
func (_GatewayRegistry *GatewayRegistryTransactorSession) Update(gatewayId [32]byte, version uint8, owner common.Address, antennaGain uint8, frequencyPlan uint8, location int64, altitude uint8) (*types.Transaction, error) {
	return _GatewayRegistry.Contract.Update(&_GatewayRegistry.TransactOpts, gatewayId, version, owner, antennaGain, frequencyPlan, location, altitude)
}

// GatewayRegistryFrequencyPlanAddedIterator is returned from FilterFrequencyPlanAdded and is used to iterate over the raw logs and unpacked data for FrequencyPlanAdded events raised by the GatewayRegistry contract.
type GatewayRegistryFrequencyPlanAddedIterator struct {
	Event *GatewayRegistryFrequencyPlanAdded // Event containing the contract specifics and raw log

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
func (it *GatewayRegistryFrequencyPlanAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayRegistryFrequencyPlanAdded)
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
		it.Event = new(GatewayRegistryFrequencyPlanAdded)
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
func (it *GatewayRegistryFrequencyPlanAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayRegistryFrequencyPlanAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayRegistryFrequencyPlanAdded represents a FrequencyPlanAdded event raised by the GatewayRegistry contract.
type GatewayRegistryFrequencyPlanAdded struct {
	Arg0 uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFrequencyPlanAdded is a free log retrieval operation binding the contract event 0x5eb1afacf56fce1c91289feeed93783eefc45e3111990d01a42bde4078ec6198.
//
// Solidity: event FrequencyPlanAdded(uint8 indexed arg0)
func (_GatewayRegistry *GatewayRegistryFilterer) FilterFrequencyPlanAdded(opts *bind.FilterOpts, arg0 []uint8) (*GatewayRegistryFrequencyPlanAddedIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _GatewayRegistry.contract.FilterLogs(opts, "FrequencyPlanAdded", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &GatewayRegistryFrequencyPlanAddedIterator{contract: _GatewayRegistry.contract, event: "FrequencyPlanAdded", logs: logs, sub: sub}, nil
}

// WatchFrequencyPlanAdded is a free log subscription operation binding the contract event 0x5eb1afacf56fce1c91289feeed93783eefc45e3111990d01a42bde4078ec6198.
//
// Solidity: event FrequencyPlanAdded(uint8 indexed arg0)
func (_GatewayRegistry *GatewayRegistryFilterer) WatchFrequencyPlanAdded(opts *bind.WatchOpts, sink chan<- *GatewayRegistryFrequencyPlanAdded, arg0 []uint8) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _GatewayRegistry.contract.WatchLogs(opts, "FrequencyPlanAdded", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayRegistryFrequencyPlanAdded)
				if err := _GatewayRegistry.contract.UnpackLog(event, "FrequencyPlanAdded", log); err != nil {
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

// ParseFrequencyPlanAdded is a log parse operation binding the contract event 0x5eb1afacf56fce1c91289feeed93783eefc45e3111990d01a42bde4078ec6198.
//
// Solidity: event FrequencyPlanAdded(uint8 indexed arg0)
func (_GatewayRegistry *GatewayRegistryFilterer) ParseFrequencyPlanAdded(log types.Log) (*GatewayRegistryFrequencyPlanAdded, error) {
	event := new(GatewayRegistryFrequencyPlanAdded)
	if err := _GatewayRegistry.contract.UnpackLog(event, "FrequencyPlanAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayRegistryFrequencyPlanRemovedIterator is returned from FilterFrequencyPlanRemoved and is used to iterate over the raw logs and unpacked data for FrequencyPlanRemoved events raised by the GatewayRegistry contract.
type GatewayRegistryFrequencyPlanRemovedIterator struct {
	Event *GatewayRegistryFrequencyPlanRemoved // Event containing the contract specifics and raw log

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
func (it *GatewayRegistryFrequencyPlanRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayRegistryFrequencyPlanRemoved)
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
		it.Event = new(GatewayRegistryFrequencyPlanRemoved)
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
func (it *GatewayRegistryFrequencyPlanRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayRegistryFrequencyPlanRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayRegistryFrequencyPlanRemoved represents a FrequencyPlanRemoved event raised by the GatewayRegistry contract.
type GatewayRegistryFrequencyPlanRemoved struct {
	Arg0 uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFrequencyPlanRemoved is a free log retrieval operation binding the contract event 0x898a81f51a8979704f8c50be34010e7d8aac481090dbff9badfd49079b58f06f.
//
// Solidity: event FrequencyPlanRemoved(uint8 indexed arg0)
func (_GatewayRegistry *GatewayRegistryFilterer) FilterFrequencyPlanRemoved(opts *bind.FilterOpts, arg0 []uint8) (*GatewayRegistryFrequencyPlanRemovedIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _GatewayRegistry.contract.FilterLogs(opts, "FrequencyPlanRemoved", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &GatewayRegistryFrequencyPlanRemovedIterator{contract: _GatewayRegistry.contract, event: "FrequencyPlanRemoved", logs: logs, sub: sub}, nil
}

// WatchFrequencyPlanRemoved is a free log subscription operation binding the contract event 0x898a81f51a8979704f8c50be34010e7d8aac481090dbff9badfd49079b58f06f.
//
// Solidity: event FrequencyPlanRemoved(uint8 indexed arg0)
func (_GatewayRegistry *GatewayRegistryFilterer) WatchFrequencyPlanRemoved(opts *bind.WatchOpts, sink chan<- *GatewayRegistryFrequencyPlanRemoved, arg0 []uint8) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _GatewayRegistry.contract.WatchLogs(opts, "FrequencyPlanRemoved", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayRegistryFrequencyPlanRemoved)
				if err := _GatewayRegistry.contract.UnpackLog(event, "FrequencyPlanRemoved", log); err != nil {
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

// ParseFrequencyPlanRemoved is a log parse operation binding the contract event 0x898a81f51a8979704f8c50be34010e7d8aac481090dbff9badfd49079b58f06f.
//
// Solidity: event FrequencyPlanRemoved(uint8 indexed arg0)
func (_GatewayRegistry *GatewayRegistryFilterer) ParseFrequencyPlanRemoved(log types.Log) (*GatewayRegistryFrequencyPlanRemoved, error) {
	event := new(GatewayRegistryFrequencyPlanRemoved)
	if err := _GatewayRegistry.contract.UnpackLog(event, "FrequencyPlanRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayRegistryGatewayOffboardedIterator is returned from FilterGatewayOffboarded and is used to iterate over the raw logs and unpacked data for GatewayOffboarded events raised by the GatewayRegistry contract.
type GatewayRegistryGatewayOffboardedIterator struct {
	Event *GatewayRegistryGatewayOffboarded // Event containing the contract specifics and raw log

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
func (it *GatewayRegistryGatewayOffboardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayRegistryGatewayOffboarded)
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
		it.Event = new(GatewayRegistryGatewayOffboarded)
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
func (it *GatewayRegistryGatewayOffboardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayRegistryGatewayOffboardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayRegistryGatewayOffboarded represents a GatewayOffboarded event raised by the GatewayRegistry contract.
type GatewayRegistryGatewayOffboarded struct {
	GatewayId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGatewayOffboarded is a free log retrieval operation binding the contract event 0x3fdaf0b55cfef0439169d60405b136ee3fd10646dd858e08a8334f9000c3eba7.
//
// Solidity: event GatewayOffboarded(bytes32 indexed gatewayId)
func (_GatewayRegistry *GatewayRegistryFilterer) FilterGatewayOffboarded(opts *bind.FilterOpts, gatewayId [][32]byte) (*GatewayRegistryGatewayOffboardedIterator, error) {

	var gatewayIdRule []interface{}
	for _, gatewayIdItem := range gatewayId {
		gatewayIdRule = append(gatewayIdRule, gatewayIdItem)
	}

	logs, sub, err := _GatewayRegistry.contract.FilterLogs(opts, "GatewayOffboarded", gatewayIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayRegistryGatewayOffboardedIterator{contract: _GatewayRegistry.contract, event: "GatewayOffboarded", logs: logs, sub: sub}, nil
}

// WatchGatewayOffboarded is a free log subscription operation binding the contract event 0x3fdaf0b55cfef0439169d60405b136ee3fd10646dd858e08a8334f9000c3eba7.
//
// Solidity: event GatewayOffboarded(bytes32 indexed gatewayId)
func (_GatewayRegistry *GatewayRegistryFilterer) WatchGatewayOffboarded(opts *bind.WatchOpts, sink chan<- *GatewayRegistryGatewayOffboarded, gatewayId [][32]byte) (event.Subscription, error) {

	var gatewayIdRule []interface{}
	for _, gatewayIdItem := range gatewayId {
		gatewayIdRule = append(gatewayIdRule, gatewayIdItem)
	}

	logs, sub, err := _GatewayRegistry.contract.WatchLogs(opts, "GatewayOffboarded", gatewayIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayRegistryGatewayOffboarded)
				if err := _GatewayRegistry.contract.UnpackLog(event, "GatewayOffboarded", log); err != nil {
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

// ParseGatewayOffboarded is a log parse operation binding the contract event 0x3fdaf0b55cfef0439169d60405b136ee3fd10646dd858e08a8334f9000c3eba7.
//
// Solidity: event GatewayOffboarded(bytes32 indexed gatewayId)
func (_GatewayRegistry *GatewayRegistryFilterer) ParseGatewayOffboarded(log types.Log) (*GatewayRegistryGatewayOffboarded, error) {
	event := new(GatewayRegistryGatewayOffboarded)
	if err := _GatewayRegistry.contract.UnpackLog(event, "GatewayOffboarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayRegistryGatewayOnboardedIterator is returned from FilterGatewayOnboarded and is used to iterate over the raw logs and unpacked data for GatewayOnboarded events raised by the GatewayRegistry contract.
type GatewayRegistryGatewayOnboardedIterator struct {
	Event *GatewayRegistryGatewayOnboarded // Event containing the contract specifics and raw log

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
func (it *GatewayRegistryGatewayOnboardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayRegistryGatewayOnboarded)
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
		it.Event = new(GatewayRegistryGatewayOnboarded)
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
func (it *GatewayRegistryGatewayOnboardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayRegistryGatewayOnboardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayRegistryGatewayOnboarded represents a GatewayOnboarded event raised by the GatewayRegistry contract.
type GatewayRegistryGatewayOnboarded struct {
	GatewayId [32]byte
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGatewayOnboarded is a free log retrieval operation binding the contract event 0xa60d6768fbd8c6717e7c8635f5bdf299dca6ff32aac817950e620eec5da47fdb.
//
// Solidity: event GatewayOnboarded(bytes32 indexed gatewayId, address indexed owner)
func (_GatewayRegistry *GatewayRegistryFilterer) FilterGatewayOnboarded(opts *bind.FilterOpts, gatewayId [][32]byte, owner []common.Address) (*GatewayRegistryGatewayOnboardedIterator, error) {

	var gatewayIdRule []interface{}
	for _, gatewayIdItem := range gatewayId {
		gatewayIdRule = append(gatewayIdRule, gatewayIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _GatewayRegistry.contract.FilterLogs(opts, "GatewayOnboarded", gatewayIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayRegistryGatewayOnboardedIterator{contract: _GatewayRegistry.contract, event: "GatewayOnboarded", logs: logs, sub: sub}, nil
}

// WatchGatewayOnboarded is a free log subscription operation binding the contract event 0xa60d6768fbd8c6717e7c8635f5bdf299dca6ff32aac817950e620eec5da47fdb.
//
// Solidity: event GatewayOnboarded(bytes32 indexed gatewayId, address indexed owner)
func (_GatewayRegistry *GatewayRegistryFilterer) WatchGatewayOnboarded(opts *bind.WatchOpts, sink chan<- *GatewayRegistryGatewayOnboarded, gatewayId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var gatewayIdRule []interface{}
	for _, gatewayIdItem := range gatewayId {
		gatewayIdRule = append(gatewayIdRule, gatewayIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _GatewayRegistry.contract.WatchLogs(opts, "GatewayOnboarded", gatewayIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayRegistryGatewayOnboarded)
				if err := _GatewayRegistry.contract.UnpackLog(event, "GatewayOnboarded", log); err != nil {
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

// ParseGatewayOnboarded is a log parse operation binding the contract event 0xa60d6768fbd8c6717e7c8635f5bdf299dca6ff32aac817950e620eec5da47fdb.
//
// Solidity: event GatewayOnboarded(bytes32 indexed gatewayId, address indexed owner)
func (_GatewayRegistry *GatewayRegistryFilterer) ParseGatewayOnboarded(log types.Log) (*GatewayRegistryGatewayOnboarded, error) {
	event := new(GatewayRegistryGatewayOnboarded)
	if err := _GatewayRegistry.contract.UnpackLog(event, "GatewayOnboarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayRegistryGatewayRegistryPausedIterator is returned from FilterGatewayRegistryPaused and is used to iterate over the raw logs and unpacked data for GatewayRegistryPaused events raised by the GatewayRegistry contract.
type GatewayRegistryGatewayRegistryPausedIterator struct {
	Event *GatewayRegistryGatewayRegistryPaused // Event containing the contract specifics and raw log

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
func (it *GatewayRegistryGatewayRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayRegistryGatewayRegistryPaused)
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
		it.Event = new(GatewayRegistryGatewayRegistryPaused)
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
func (it *GatewayRegistryGatewayRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayRegistryGatewayRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayRegistryGatewayRegistryPaused represents a GatewayRegistryPaused event raised by the GatewayRegistry contract.
type GatewayRegistryGatewayRegistryPaused struct {
	Onboards  bool
	Offboards bool
	Updates   bool
	Transfers bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGatewayRegistryPaused is a free log retrieval operation binding the contract event 0x0d30b4d10c0e924cbf1332f59500d2f843b7f4910bebc4e6347a492cf2dfdee5.
//
// Solidity: event GatewayRegistryPaused(bool onboards, bool offboards, bool updates, bool transfers)
func (_GatewayRegistry *GatewayRegistryFilterer) FilterGatewayRegistryPaused(opts *bind.FilterOpts) (*GatewayRegistryGatewayRegistryPausedIterator, error) {

	logs, sub, err := _GatewayRegistry.contract.FilterLogs(opts, "GatewayRegistryPaused")
	if err != nil {
		return nil, err
	}
	return &GatewayRegistryGatewayRegistryPausedIterator{contract: _GatewayRegistry.contract, event: "GatewayRegistryPaused", logs: logs, sub: sub}, nil
}

// WatchGatewayRegistryPaused is a free log subscription operation binding the contract event 0x0d30b4d10c0e924cbf1332f59500d2f843b7f4910bebc4e6347a492cf2dfdee5.
//
// Solidity: event GatewayRegistryPaused(bool onboards, bool offboards, bool updates, bool transfers)
func (_GatewayRegistry *GatewayRegistryFilterer) WatchGatewayRegistryPaused(opts *bind.WatchOpts, sink chan<- *GatewayRegistryGatewayRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _GatewayRegistry.contract.WatchLogs(opts, "GatewayRegistryPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayRegistryGatewayRegistryPaused)
				if err := _GatewayRegistry.contract.UnpackLog(event, "GatewayRegistryPaused", log); err != nil {
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

// ParseGatewayRegistryPaused is a log parse operation binding the contract event 0x0d30b4d10c0e924cbf1332f59500d2f843b7f4910bebc4e6347a492cf2dfdee5.
//
// Solidity: event GatewayRegistryPaused(bool onboards, bool offboards, bool updates, bool transfers)
func (_GatewayRegistry *GatewayRegistryFilterer) ParseGatewayRegistryPaused(log types.Log) (*GatewayRegistryGatewayRegistryPaused, error) {
	event := new(GatewayRegistryGatewayRegistryPaused)
	if err := _GatewayRegistry.contract.UnpackLog(event, "GatewayRegistryPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayRegistryGatewayTransferredIterator is returned from FilterGatewayTransferred and is used to iterate over the raw logs and unpacked data for GatewayTransferred events raised by the GatewayRegistry contract.
type GatewayRegistryGatewayTransferredIterator struct {
	Event *GatewayRegistryGatewayTransferred // Event containing the contract specifics and raw log

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
func (it *GatewayRegistryGatewayTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayRegistryGatewayTransferred)
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
		it.Event = new(GatewayRegistryGatewayTransferred)
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
func (it *GatewayRegistryGatewayTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayRegistryGatewayTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayRegistryGatewayTransferred represents a GatewayTransferred event raised by the GatewayRegistry contract.
type GatewayRegistryGatewayTransferred struct {
	GatewayId [32]byte
	OldOwner  common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGatewayTransferred is a free log retrieval operation binding the contract event 0xd83561db967da60f545e3a0538a24a25874f4aa9b3b14ac2e63415c8dea21c0e.
//
// Solidity: event GatewayTransferred(bytes32 indexed gatewayId, address indexed oldOwner, address indexed newOwner)
func (_GatewayRegistry *GatewayRegistryFilterer) FilterGatewayTransferred(opts *bind.FilterOpts, gatewayId [][32]byte, oldOwner []common.Address, newOwner []common.Address) (*GatewayRegistryGatewayTransferredIterator, error) {

	var gatewayIdRule []interface{}
	for _, gatewayIdItem := range gatewayId {
		gatewayIdRule = append(gatewayIdRule, gatewayIdItem)
	}
	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayRegistry.contract.FilterLogs(opts, "GatewayTransferred", gatewayIdRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayRegistryGatewayTransferredIterator{contract: _GatewayRegistry.contract, event: "GatewayTransferred", logs: logs, sub: sub}, nil
}

// WatchGatewayTransferred is a free log subscription operation binding the contract event 0xd83561db967da60f545e3a0538a24a25874f4aa9b3b14ac2e63415c8dea21c0e.
//
// Solidity: event GatewayTransferred(bytes32 indexed gatewayId, address indexed oldOwner, address indexed newOwner)
func (_GatewayRegistry *GatewayRegistryFilterer) WatchGatewayTransferred(opts *bind.WatchOpts, sink chan<- *GatewayRegistryGatewayTransferred, gatewayId [][32]byte, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var gatewayIdRule []interface{}
	for _, gatewayIdItem := range gatewayId {
		gatewayIdRule = append(gatewayIdRule, gatewayIdItem)
	}
	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayRegistry.contract.WatchLogs(opts, "GatewayTransferred", gatewayIdRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayRegistryGatewayTransferred)
				if err := _GatewayRegistry.contract.UnpackLog(event, "GatewayTransferred", log); err != nil {
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

// ParseGatewayTransferred is a log parse operation binding the contract event 0xd83561db967da60f545e3a0538a24a25874f4aa9b3b14ac2e63415c8dea21c0e.
//
// Solidity: event GatewayTransferred(bytes32 indexed gatewayId, address indexed oldOwner, address indexed newOwner)
func (_GatewayRegistry *GatewayRegistryFilterer) ParseGatewayTransferred(log types.Log) (*GatewayRegistryGatewayTransferred, error) {
	event := new(GatewayRegistryGatewayTransferred)
	if err := _GatewayRegistry.contract.UnpackLog(event, "GatewayTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayRegistryGatewayUpdatedIterator is returned from FilterGatewayUpdated and is used to iterate over the raw logs and unpacked data for GatewayUpdated events raised by the GatewayRegistry contract.
type GatewayRegistryGatewayUpdatedIterator struct {
	Event *GatewayRegistryGatewayUpdated // Event containing the contract specifics and raw log

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
func (it *GatewayRegistryGatewayUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayRegistryGatewayUpdated)
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
		it.Event = new(GatewayRegistryGatewayUpdated)
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
func (it *GatewayRegistryGatewayUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayRegistryGatewayUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayRegistryGatewayUpdated represents a GatewayUpdated event raised by the GatewayRegistry contract.
type GatewayRegistryGatewayUpdated struct {
	GatewayId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGatewayUpdated is a free log retrieval operation binding the contract event 0x6d07f738d8b7adf4f097fc73b96bb18ceeebfc0d2274b8b8e138973b5dc84016.
//
// Solidity: event GatewayUpdated(bytes32 indexed gatewayId)
func (_GatewayRegistry *GatewayRegistryFilterer) FilterGatewayUpdated(opts *bind.FilterOpts, gatewayId [][32]byte) (*GatewayRegistryGatewayUpdatedIterator, error) {

	var gatewayIdRule []interface{}
	for _, gatewayIdItem := range gatewayId {
		gatewayIdRule = append(gatewayIdRule, gatewayIdItem)
	}

	logs, sub, err := _GatewayRegistry.contract.FilterLogs(opts, "GatewayUpdated", gatewayIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayRegistryGatewayUpdatedIterator{contract: _GatewayRegistry.contract, event: "GatewayUpdated", logs: logs, sub: sub}, nil
}

// WatchGatewayUpdated is a free log subscription operation binding the contract event 0x6d07f738d8b7adf4f097fc73b96bb18ceeebfc0d2274b8b8e138973b5dc84016.
//
// Solidity: event GatewayUpdated(bytes32 indexed gatewayId)
func (_GatewayRegistry *GatewayRegistryFilterer) WatchGatewayUpdated(opts *bind.WatchOpts, sink chan<- *GatewayRegistryGatewayUpdated, gatewayId [][32]byte) (event.Subscription, error) {

	var gatewayIdRule []interface{}
	for _, gatewayIdItem := range gatewayId {
		gatewayIdRule = append(gatewayIdRule, gatewayIdItem)
	}

	logs, sub, err := _GatewayRegistry.contract.WatchLogs(opts, "GatewayUpdated", gatewayIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayRegistryGatewayUpdated)
				if err := _GatewayRegistry.contract.UnpackLog(event, "GatewayUpdated", log); err != nil {
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

// ParseGatewayUpdated is a log parse operation binding the contract event 0x6d07f738d8b7adf4f097fc73b96bb18ceeebfc0d2274b8b8e138973b5dc84016.
//
// Solidity: event GatewayUpdated(bytes32 indexed gatewayId)
func (_GatewayRegistry *GatewayRegistryFilterer) ParseGatewayUpdated(log types.Log) (*GatewayRegistryGatewayUpdated, error) {
	event := new(GatewayRegistryGatewayUpdated)
	if err := _GatewayRegistry.contract.UnpackLog(event, "GatewayUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the GatewayRegistry contract.
type GatewayRegistryRoleAdminChangedIterator struct {
	Event *GatewayRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *GatewayRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayRegistryRoleAdminChanged)
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
		it.Event = new(GatewayRegistryRoleAdminChanged)
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
func (it *GatewayRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the GatewayRegistry contract.
type GatewayRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayRegistry *GatewayRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GatewayRegistryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _GatewayRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GatewayRegistryRoleAdminChangedIterator{contract: _GatewayRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayRegistry *GatewayRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GatewayRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _GatewayRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayRegistryRoleAdminChanged)
				if err := _GatewayRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayRegistry *GatewayRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*GatewayRegistryRoleAdminChanged, error) {
	event := new(GatewayRegistryRoleAdminChanged)
	if err := _GatewayRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the GatewayRegistry contract.
type GatewayRegistryRoleGrantedIterator struct {
	Event *GatewayRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *GatewayRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayRegistryRoleGranted)
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
		it.Event = new(GatewayRegistryRoleGranted)
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
func (it *GatewayRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayRegistryRoleGranted represents a RoleGranted event raised by the GatewayRegistry contract.
type GatewayRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayRegistry *GatewayRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayRegistryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GatewayRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayRegistryRoleGrantedIterator{contract: _GatewayRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayRegistry *GatewayRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GatewayRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GatewayRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayRegistryRoleGranted)
				if err := _GatewayRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayRegistry *GatewayRegistryFilterer) ParseRoleGranted(log types.Log) (*GatewayRegistryRoleGranted, error) {
	event := new(GatewayRegistryRoleGranted)
	if err := _GatewayRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the GatewayRegistry contract.
type GatewayRegistryRoleRevokedIterator struct {
	Event *GatewayRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *GatewayRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayRegistryRoleRevoked)
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
		it.Event = new(GatewayRegistryRoleRevoked)
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
func (it *GatewayRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayRegistryRoleRevoked represents a RoleRevoked event raised by the GatewayRegistry contract.
type GatewayRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayRegistry *GatewayRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayRegistryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GatewayRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayRegistryRoleRevokedIterator{contract: _GatewayRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayRegistry *GatewayRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GatewayRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GatewayRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayRegistryRoleRevoked)
				if err := _GatewayRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayRegistry *GatewayRegistryFilterer) ParseRoleRevoked(log types.Log) (*GatewayRegistryRoleRevoked, error) {
	event := new(GatewayRegistryRoleRevoked)
	if err := _GatewayRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
