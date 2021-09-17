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
)

// ActivatableMetaData contains all meta data concerning the Activatable contract.
var ActivatableMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"active\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deactivate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"Deactivate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"Activate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
	Sigs: map[string]string{
		"0f15f4c0": "activate()",
		"02fb0c5e": "active()",
		"51b42b00": "deactivate()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060405260008054600160a860020a0319163317905561039e806100256000396000f3006080604052600436106100775763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302fb0c5e811461007c5780630f15f4c0146100a557806351b42b00146100bc578063715018a6146100d15780638da5cb5b146100e6578063f2fde38b14610117575b600080fd5b34801561008857600080fd5b50610091610138565b604080519115158252519081900360200190f35b3480156100b157600080fd5b506100ba610159565b005b3480156100c857600080fd5b506100ba6101f7565b3480156100dd57600080fd5b506100ba610257565b3480156100f257600080fd5b506100fb6102c3565b60408051600160a060020a039092168252519081900360200190f35b34801561012357600080fd5b506100ba600160a060020a03600435166102d2565b60005474010000000000000000000000000000000000000000900460ff1681565b60005474010000000000000000000000000000000000000000900460ff161561018157600080fd5b600054600160a060020a0316331461019857600080fd5b6000805474ff000000000000000000000000000000000000000019167401000000000000000000000000000000000000000017815560405133917ff7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba1591a2565b60005474010000000000000000000000000000000000000000900460ff16151561022057600080fd5b600054600160a060020a0316331461023757600080fd5b6000805474ff000000000000000000000000000000000000000019169055565b600054600160a060020a0316331461026e57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600054600160a060020a031633146102e957600080fd5b6102f2816102f5565b50565b600160a060020a038116151561030a57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820ef71c82f84515187cc3c80fbd40f5fc65eac29967394d8f5e63ff8515330370f0029",
}

// ActivatableABI is the input ABI used to generate the binding from.
// Deprecated: Use ActivatableMetaData.ABI instead.
var ActivatableABI = ActivatableMetaData.ABI

// Deprecated: Use ActivatableMetaData.Sigs instead.
// ActivatableFuncSigs maps the 4-byte function signature to its string representation.
var ActivatableFuncSigs = ActivatableMetaData.Sigs

// ActivatableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ActivatableMetaData.Bin instead.
var ActivatableBin = ActivatableMetaData.Bin

// DeployActivatable deploys a new Ethereum contract, binding an instance of Activatable to it.
func DeployActivatable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Activatable, error) {
	parsed, err := ActivatableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ActivatableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Activatable{ActivatableCaller: ActivatableCaller{contract: contract}, ActivatableTransactor: ActivatableTransactor{contract: contract}, ActivatableFilterer: ActivatableFilterer{contract: contract}}, nil
}

// Activatable is an auto generated Go binding around an Ethereum contract.
type Activatable struct {
	ActivatableCaller     // Read-only binding to the contract
	ActivatableTransactor // Write-only binding to the contract
	ActivatableFilterer   // Log filterer for contract events
}

// ActivatableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivatableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivatableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivatableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivatableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivatableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivatableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivatableSession struct {
	Contract     *Activatable      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ActivatableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivatableCallerSession struct {
	Contract *ActivatableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ActivatableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivatableTransactorSession struct {
	Contract     *ActivatableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ActivatableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivatableRaw struct {
	Contract *Activatable // Generic contract binding to access the raw methods on
}

// ActivatableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivatableCallerRaw struct {
	Contract *ActivatableCaller // Generic read-only contract binding to access the raw methods on
}

// ActivatableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivatableTransactorRaw struct {
	Contract *ActivatableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivatable creates a new instance of Activatable, bound to a specific deployed contract.
func NewActivatable(address common.Address, backend bind.ContractBackend) (*Activatable, error) {
	contract, err := bindActivatable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Activatable{ActivatableCaller: ActivatableCaller{contract: contract}, ActivatableTransactor: ActivatableTransactor{contract: contract}, ActivatableFilterer: ActivatableFilterer{contract: contract}}, nil
}

// NewActivatableCaller creates a new read-only instance of Activatable, bound to a specific deployed contract.
func NewActivatableCaller(address common.Address, caller bind.ContractCaller) (*ActivatableCaller, error) {
	contract, err := bindActivatable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivatableCaller{contract: contract}, nil
}

// NewActivatableTransactor creates a new write-only instance of Activatable, bound to a specific deployed contract.
func NewActivatableTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivatableTransactor, error) {
	contract, err := bindActivatable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivatableTransactor{contract: contract}, nil
}

// NewActivatableFilterer creates a new log filterer instance of Activatable, bound to a specific deployed contract.
func NewActivatableFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivatableFilterer, error) {
	contract, err := bindActivatable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivatableFilterer{contract: contract}, nil
}

// bindActivatable binds a generic wrapper to an already deployed contract.
func bindActivatable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivatableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Activatable *ActivatableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Activatable.Contract.ActivatableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Activatable *ActivatableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activatable.Contract.ActivatableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Activatable *ActivatableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Activatable.Contract.ActivatableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Activatable *ActivatableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Activatable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Activatable *ActivatableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activatable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Activatable *ActivatableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Activatable.Contract.contract.Transact(opts, method, params...)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Activatable *ActivatableCaller) Active(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Activatable.contract.Call(opts, &out, "active")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Activatable *ActivatableSession) Active() (bool, error) {
	return _Activatable.Contract.Active(&_Activatable.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Activatable *ActivatableCallerSession) Active() (bool, error) {
	return _Activatable.Contract.Active(&_Activatable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Activatable *ActivatableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Activatable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Activatable *ActivatableSession) Owner() (common.Address, error) {
	return _Activatable.Contract.Owner(&_Activatable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Activatable *ActivatableCallerSession) Owner() (common.Address, error) {
	return _Activatable.Contract.Owner(&_Activatable.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Activatable *ActivatableTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activatable.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Activatable *ActivatableSession) Activate() (*types.Transaction, error) {
	return _Activatable.Contract.Activate(&_Activatable.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Activatable *ActivatableTransactorSession) Activate() (*types.Transaction, error) {
	return _Activatable.Contract.Activate(&_Activatable.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Activatable *ActivatableTransactor) Deactivate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activatable.contract.Transact(opts, "deactivate")
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Activatable *ActivatableSession) Deactivate() (*types.Transaction, error) {
	return _Activatable.Contract.Deactivate(&_Activatable.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Activatable *ActivatableTransactorSession) Deactivate() (*types.Transaction, error) {
	return _Activatable.Contract.Deactivate(&_Activatable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Activatable *ActivatableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activatable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Activatable *ActivatableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Activatable.Contract.RenounceOwnership(&_Activatable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Activatable *ActivatableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Activatable.Contract.RenounceOwnership(&_Activatable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Activatable *ActivatableTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Activatable.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Activatable *ActivatableSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Activatable.Contract.TransferOwnership(&_Activatable.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Activatable *ActivatableTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Activatable.Contract.TransferOwnership(&_Activatable.TransactOpts, _newOwner)
}

// ActivatableActivateIterator is returned from FilterActivate and is used to iterate over the raw logs and unpacked data for Activate events raised by the Activatable contract.
type ActivatableActivateIterator struct {
	Event *ActivatableActivate // Event containing the contract specifics and raw log

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
func (it *ActivatableActivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivatableActivate)
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
		it.Event = new(ActivatableActivate)
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
func (it *ActivatableActivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivatableActivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivatableActivate represents a Activate event raised by the Activatable contract.
type ActivatableActivate struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterActivate is a free log retrieval operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: event Activate(address indexed _sender)
func (_Activatable *ActivatableFilterer) FilterActivate(opts *bind.FilterOpts, _sender []common.Address) (*ActivatableActivateIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Activatable.contract.FilterLogs(opts, "Activate", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ActivatableActivateIterator{contract: _Activatable.contract, event: "Activate", logs: logs, sub: sub}, nil
}

// WatchActivate is a free log subscription operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: event Activate(address indexed _sender)
func (_Activatable *ActivatableFilterer) WatchActivate(opts *bind.WatchOpts, sink chan<- *ActivatableActivate, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Activatable.contract.WatchLogs(opts, "Activate", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivatableActivate)
				if err := _Activatable.contract.UnpackLog(event, "Activate", log); err != nil {
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

// ParseActivate is a log parse operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: event Activate(address indexed _sender)
func (_Activatable *ActivatableFilterer) ParseActivate(log types.Log) (*ActivatableActivate, error) {
	event := new(ActivatableActivate)
	if err := _Activatable.contract.UnpackLog(event, "Activate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActivatableDeactivateIterator is returned from FilterDeactivate and is used to iterate over the raw logs and unpacked data for Deactivate events raised by the Activatable contract.
type ActivatableDeactivateIterator struct {
	Event *ActivatableDeactivate // Event containing the contract specifics and raw log

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
func (it *ActivatableDeactivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivatableDeactivate)
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
		it.Event = new(ActivatableDeactivate)
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
func (it *ActivatableDeactivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivatableDeactivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivatableDeactivate represents a Deactivate event raised by the Activatable contract.
type ActivatableDeactivate struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeactivate is a free log retrieval operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: event Deactivate(address indexed _sender)
func (_Activatable *ActivatableFilterer) FilterDeactivate(opts *bind.FilterOpts, _sender []common.Address) (*ActivatableDeactivateIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Activatable.contract.FilterLogs(opts, "Deactivate", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ActivatableDeactivateIterator{contract: _Activatable.contract, event: "Deactivate", logs: logs, sub: sub}, nil
}

// WatchDeactivate is a free log subscription operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: event Deactivate(address indexed _sender)
func (_Activatable *ActivatableFilterer) WatchDeactivate(opts *bind.WatchOpts, sink chan<- *ActivatableDeactivate, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Activatable.contract.WatchLogs(opts, "Deactivate", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivatableDeactivate)
				if err := _Activatable.contract.UnpackLog(event, "Deactivate", log); err != nil {
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

// ParseDeactivate is a log parse operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: event Deactivate(address indexed _sender)
func (_Activatable *ActivatableFilterer) ParseDeactivate(log types.Log) (*ActivatableDeactivate, error) {
	event := new(ActivatableDeactivate)
	if err := _Activatable.contract.UnpackLog(event, "Deactivate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActivatableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Activatable contract.
type ActivatableOwnershipRenouncedIterator struct {
	Event *ActivatableOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *ActivatableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivatableOwnershipRenounced)
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
		it.Event = new(ActivatableOwnershipRenounced)
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
func (it *ActivatableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivatableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivatableOwnershipRenounced represents a OwnershipRenounced event raised by the Activatable contract.
type ActivatableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Activatable *ActivatableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*ActivatableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Activatable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ActivatableOwnershipRenouncedIterator{contract: _Activatable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Activatable *ActivatableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *ActivatableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Activatable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivatableOwnershipRenounced)
				if err := _Activatable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Activatable *ActivatableFilterer) ParseOwnershipRenounced(log types.Log) (*ActivatableOwnershipRenounced, error) {
	event := new(ActivatableOwnershipRenounced)
	if err := _Activatable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActivatableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Activatable contract.
type ActivatableOwnershipTransferredIterator struct {
	Event *ActivatableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ActivatableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivatableOwnershipTransferred)
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
		it.Event = new(ActivatableOwnershipTransferred)
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
func (it *ActivatableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivatableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivatableOwnershipTransferred represents a OwnershipTransferred event raised by the Activatable contract.
type ActivatableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Activatable *ActivatableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ActivatableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Activatable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ActivatableOwnershipTransferredIterator{contract: _Activatable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Activatable *ActivatableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ActivatableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Activatable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivatableOwnershipTransferred)
				if err := _Activatable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Activatable *ActivatableFilterer) ParseOwnershipTransferred(log types.Log) (*ActivatableOwnershipTransferred, error) {
	event := new(ActivatableOwnershipTransferred)
	if err := _Activatable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestructibleMetaData contains all meta data concerning the Destructible contract.
var DestructibleMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"destroyAndSend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
	Sigs: map[string]string{
		"83197ef0": "destroy()",
		"f5074f41": "destroyAndSend(address)",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060405260008054600160a060020a0319163317905561029f806100256000396000f30060806040526004361061006c5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663715018a6811461007157806383197ef0146100885780638da5cb5b1461009d578063f2fde38b146100ce578063f5074f41146100ef575b600080fd5b34801561007d57600080fd5b50610086610110565b005b34801561009457600080fd5b5061008661017c565b3480156100a957600080fd5b506100b26101a1565b60408051600160a060020a039092168252519081900360200190f35b3480156100da57600080fd5b50610086600160a060020a03600435166101b0565b3480156100fb57600080fd5b50610086600160a060020a03600435166101d3565b600054600160a060020a0316331461012757600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a0316331461019357600080fd5b600054600160a060020a0316ff5b600054600160a060020a031681565b600054600160a060020a031633146101c757600080fd5b6101d0816101f6565b50565b600054600160a060020a031633146101ea57600080fd5b80600160a060020a0316ff5b600160a060020a038116151561020b57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820586cda9b042f4db12deedd2a2efc1a6571ce55522261bd25aa53a00dc5b8df320029",
}

// DestructibleABI is the input ABI used to generate the binding from.
// Deprecated: Use DestructibleMetaData.ABI instead.
var DestructibleABI = DestructibleMetaData.ABI

// Deprecated: Use DestructibleMetaData.Sigs instead.
// DestructibleFuncSigs maps the 4-byte function signature to its string representation.
var DestructibleFuncSigs = DestructibleMetaData.Sigs

// DestructibleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DestructibleMetaData.Bin instead.
var DestructibleBin = DestructibleMetaData.Bin

// DeployDestructible deploys a new Ethereum contract, binding an instance of Destructible to it.
func DeployDestructible(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Destructible, error) {
	parsed, err := DestructibleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DestructibleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Destructible{DestructibleCaller: DestructibleCaller{contract: contract}, DestructibleTransactor: DestructibleTransactor{contract: contract}, DestructibleFilterer: DestructibleFilterer{contract: contract}}, nil
}

// Destructible is an auto generated Go binding around an Ethereum contract.
type Destructible struct {
	DestructibleCaller     // Read-only binding to the contract
	DestructibleTransactor // Write-only binding to the contract
	DestructibleFilterer   // Log filterer for contract events
}

// DestructibleCaller is an auto generated read-only Go binding around an Ethereum contract.
type DestructibleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestructibleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DestructibleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestructibleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DestructibleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestructibleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DestructibleSession struct {
	Contract     *Destructible     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DestructibleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DestructibleCallerSession struct {
	Contract *DestructibleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DestructibleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DestructibleTransactorSession struct {
	Contract     *DestructibleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DestructibleRaw is an auto generated low-level Go binding around an Ethereum contract.
type DestructibleRaw struct {
	Contract *Destructible // Generic contract binding to access the raw methods on
}

// DestructibleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DestructibleCallerRaw struct {
	Contract *DestructibleCaller // Generic read-only contract binding to access the raw methods on
}

// DestructibleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DestructibleTransactorRaw struct {
	Contract *DestructibleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDestructible creates a new instance of Destructible, bound to a specific deployed contract.
func NewDestructible(address common.Address, backend bind.ContractBackend) (*Destructible, error) {
	contract, err := bindDestructible(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Destructible{DestructibleCaller: DestructibleCaller{contract: contract}, DestructibleTransactor: DestructibleTransactor{contract: contract}, DestructibleFilterer: DestructibleFilterer{contract: contract}}, nil
}

// NewDestructibleCaller creates a new read-only instance of Destructible, bound to a specific deployed contract.
func NewDestructibleCaller(address common.Address, caller bind.ContractCaller) (*DestructibleCaller, error) {
	contract, err := bindDestructible(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DestructibleCaller{contract: contract}, nil
}

// NewDestructibleTransactor creates a new write-only instance of Destructible, bound to a specific deployed contract.
func NewDestructibleTransactor(address common.Address, transactor bind.ContractTransactor) (*DestructibleTransactor, error) {
	contract, err := bindDestructible(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DestructibleTransactor{contract: contract}, nil
}

// NewDestructibleFilterer creates a new log filterer instance of Destructible, bound to a specific deployed contract.
func NewDestructibleFilterer(address common.Address, filterer bind.ContractFilterer) (*DestructibleFilterer, error) {
	contract, err := bindDestructible(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DestructibleFilterer{contract: contract}, nil
}

// bindDestructible binds a generic wrapper to an already deployed contract.
func bindDestructible(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DestructibleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destructible *DestructibleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Destructible.Contract.DestructibleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destructible *DestructibleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destructible.Contract.DestructibleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destructible *DestructibleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destructible.Contract.DestructibleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destructible *DestructibleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Destructible.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destructible *DestructibleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destructible.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destructible *DestructibleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destructible.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Destructible *DestructibleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Destructible.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Destructible *DestructibleSession) Owner() (common.Address, error) {
	return _Destructible.Contract.Owner(&_Destructible.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Destructible *DestructibleCallerSession) Owner() (common.Address, error) {
	return _Destructible.Contract.Owner(&_Destructible.CallOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Destructible *DestructibleTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destructible.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Destructible *DestructibleSession) Destroy() (*types.Transaction, error) {
	return _Destructible.Contract.Destroy(&_Destructible.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Destructible *DestructibleTransactorSession) Destroy() (*types.Transaction, error) {
	return _Destructible.Contract.Destroy(&_Destructible.TransactOpts)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_Destructible *DestructibleTransactor) DestroyAndSend(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _Destructible.contract.Transact(opts, "destroyAndSend", _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_Destructible *DestructibleSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _Destructible.Contract.DestroyAndSend(&_Destructible.TransactOpts, _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_Destructible *DestructibleTransactorSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _Destructible.Contract.DestroyAndSend(&_Destructible.TransactOpts, _recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Destructible *DestructibleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destructible.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Destructible *DestructibleSession) RenounceOwnership() (*types.Transaction, error) {
	return _Destructible.Contract.RenounceOwnership(&_Destructible.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Destructible *DestructibleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Destructible.Contract.RenounceOwnership(&_Destructible.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Destructible *DestructibleTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Destructible.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Destructible *DestructibleSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Destructible.Contract.TransferOwnership(&_Destructible.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Destructible *DestructibleTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Destructible.Contract.TransferOwnership(&_Destructible.TransactOpts, _newOwner)
}

// DestructibleOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Destructible contract.
type DestructibleOwnershipRenouncedIterator struct {
	Event *DestructibleOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *DestructibleOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestructibleOwnershipRenounced)
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
		it.Event = new(DestructibleOwnershipRenounced)
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
func (it *DestructibleOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestructibleOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestructibleOwnershipRenounced represents a OwnershipRenounced event raised by the Destructible contract.
type DestructibleOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Destructible *DestructibleFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*DestructibleOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Destructible.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DestructibleOwnershipRenouncedIterator{contract: _Destructible.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Destructible *DestructibleFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *DestructibleOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Destructible.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestructibleOwnershipRenounced)
				if err := _Destructible.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Destructible *DestructibleFilterer) ParseOwnershipRenounced(log types.Log) (*DestructibleOwnershipRenounced, error) {
	event := new(DestructibleOwnershipRenounced)
	if err := _Destructible.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestructibleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Destructible contract.
type DestructibleOwnershipTransferredIterator struct {
	Event *DestructibleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DestructibleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestructibleOwnershipTransferred)
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
		it.Event = new(DestructibleOwnershipTransferred)
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
func (it *DestructibleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestructibleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestructibleOwnershipTransferred represents a OwnershipTransferred event raised by the Destructible contract.
type DestructibleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Destructible *DestructibleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DestructibleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Destructible.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DestructibleOwnershipTransferredIterator{contract: _Destructible.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Destructible *DestructibleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DestructibleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Destructible.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestructibleOwnershipTransferred)
				if err := _Destructible.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Destructible *DestructibleFilterer) ParseOwnershipTransferred(log types.Log) (*DestructibleOwnershipTransferred, error) {
	event := new(DestructibleOwnershipTransferred)
	if err := _Destructible.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b5060008054600160a060020a0319163317905561020b806100326000396000f3006080604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663715018a6811461005b5780638da5cb5b14610072578063f2fde38b146100a3575b600080fd5b34801561006757600080fd5b506100706100c4565b005b34801561007e57600080fd5b50610087610130565b60408051600160a060020a039092168252519081900360200190f35b3480156100af57600080fd5b50610070600160a060020a036004351661013f565b600054600160a060020a031633146100db57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600054600160a060020a0316331461015657600080fd5b61015f81610162565b50565b600160a060020a038116151561017757600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a723058205520e4822844807b1250ae009f887e48f32a20b843a745a2eddd3d0e605ad7260029",
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Deprecated: Use OwnableMetaData.Sigs instead.
// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = OwnableMetaData.Sigs

// OwnableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OwnableMetaData.Bin instead.
var OwnableBin = OwnableMetaData.Bin

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := OwnableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// OwnableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Ownable contract.
type OwnableOwnershipRenouncedIterator struct {
	Event *OwnableOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipRenounced)
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
		it.Event = new(OwnableOwnershipRenounced)
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
func (it *OwnableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipRenounced represents a OwnershipRenounced event raised by the Ownable contract.
type OwnableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*OwnableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipRenouncedIterator{contract: _Ownable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipRenounced)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipRenounced(log types.Log) (*OwnableOwnershipRenounced, error) {
	event := new(OwnableOwnershipRenounced)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableMetaData contains all meta data concerning the Pausable contract.
var PausableMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
		"8456cb59": "pause()",
		"5c975abb": "paused()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
		"3f4ba83a": "unpause()",
	},
	Bin: "0x608060405260008054600160a860020a031916331790556103c4806100256000396000f3006080604052600436106100775763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633f4ba83a811461007c5780635c975abb14610093578063715018a6146100bc5780638456cb59146100d15780638da5cb5b146100e6578063f2fde38b14610117575b600080fd5b34801561008857600080fd5b50610091610138565b005b34801561009f57600080fd5b506100a86101bf565b604080519115158252519081900360200190f35b3480156100c857600080fd5b506100916101e0565b3480156100dd57600080fd5b5061009161024c565b3480156100f257600080fd5b506100fb6102e9565b60408051600160a060020a039092168252519081900360200190f35b34801561012357600080fd5b50610091600160a060020a03600435166102f8565b600054600160a060020a0316331461014f57600080fd5b60005474010000000000000000000000000000000000000000900460ff16151561017857600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a1565b60005474010000000000000000000000000000000000000000900460ff1681565b600054600160a060020a031633146101f757600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a0316331461026357600080fd5b60005474010000000000000000000000000000000000000000900460ff161561028b57600080fd5b6000805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a1565b600054600160a060020a031681565b600054600160a060020a0316331461030f57600080fd5b6103188161031b565b50565b600160a060020a038116151561033057600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a723058205f7697823a77447eb932c449992c009c856d23fb2a5ed859e810b08ed0ac7dfe0029",
}

// PausableABI is the input ABI used to generate the binding from.
// Deprecated: Use PausableMetaData.ABI instead.
var PausableABI = PausableMetaData.ABI

// Deprecated: Use PausableMetaData.Sigs instead.
// PausableFuncSigs maps the 4-byte function signature to its string representation.
var PausableFuncSigs = PausableMetaData.Sigs

// PausableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PausableMetaData.Bin instead.
var PausableBin = PausableMetaData.Bin

// DeployPausable deploys a new Ethereum contract, binding an instance of Pausable to it.
func DeployPausable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pausable, error) {
	parsed, err := PausableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PausableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pausable *PausableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pausable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pausable *PausableSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pausable *PausableCallerSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pausable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactorSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pausable *PausableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pausable *PausableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pausable.Contract.RenounceOwnership(&_Pausable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pausable *PausableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pausable.Contract.RenounceOwnership(&_Pausable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Pausable *PausableTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Pausable *PausableSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Pausable *PausableTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// PausableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Pausable contract.
type PausableOwnershipRenouncedIterator struct {
	Event *PausableOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PausableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableOwnershipRenounced)
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
		it.Event = new(PausableOwnershipRenounced)
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
func (it *PausableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableOwnershipRenounced represents a OwnershipRenounced event raised by the Pausable contract.
type PausableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Pausable *PausableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PausableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableOwnershipRenouncedIterator{contract: _Pausable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Pausable *PausableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PausableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableOwnershipRenounced)
				if err := _Pausable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Pausable *PausableFilterer) ParseOwnershipRenounced(log types.Log) (*PausableOwnershipRenounced, error) {
	event := new(PausableOwnershipRenounced)
	if err := _Pausable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pausable contract.
type PausableOwnershipTransferredIterator struct {
	Event *PausableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PausableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableOwnershipTransferred)
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
		it.Event = new(PausableOwnershipTransferred)
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
func (it *PausableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableOwnershipTransferred represents a OwnershipTransferred event raised by the Pausable contract.
type PausableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pausable *PausableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PausableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableOwnershipTransferredIterator{contract: _Pausable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pausable *PausableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PausableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableOwnershipTransferred)
				if err := _Pausable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pausable *PausableFilterer) ParseOwnershipTransferred(log types.Log) (*PausableOwnershipTransferred, error) {
	event := new(PausableOwnershipTransferred)
	if err := _Pausable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausablePauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Pausable contract.
type PausablePauseIterator struct {
	Event *PausablePause // Event containing the contract specifics and raw log

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
func (it *PausablePauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePause)
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
		it.Event = new(PausablePause)
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
func (it *PausablePauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePause represents a Pause event raised by the Pausable contract.
type PausablePause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Pausable *PausableFilterer) FilterPause(opts *bind.FilterOpts) (*PausablePauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PausablePauseIterator{contract: _Pausable.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Pausable *PausableFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PausablePause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePause)
				if err := _Pausable.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Pausable *PausableFilterer) ParsePause(log types.Log) (*PausablePause, error) {
	event := new(PausablePause)
	if err := _Pausable.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Pausable contract.
type PausableUnpauseIterator struct {
	Event *PausableUnpause // Event containing the contract specifics and raw log

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
func (it *PausableUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpause)
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
		it.Event = new(PausableUnpause)
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
func (it *PausableUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpause represents a Unpause event raised by the Pausable contract.
type PausableUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Pausable *PausableFilterer) FilterUnpause(opts *bind.FilterOpts) (*PausableUnpauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PausableUnpauseIterator{contract: _Pausable.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Pausable *PausableFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PausableUnpause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpause)
				if err := _Pausable.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Pausable *PausableFilterer) ParseUnpause(log types.Log) (*PausableUnpause, error) {
	event := new(PausableUnpause)
	if err := _Pausable.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomMetaData contains all meta data concerning the Room contract.
var RoomMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"active\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reward\",\"type\":\"uint256\"},{\"name\":\"_dest\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"sendReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refundToOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deactivate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardSent\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"destroyAndSend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_creator\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_depositor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_depositedValue\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_dest\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_reward\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"RewardSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_dest\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_refundedBalance\",\"type\":\"uint256\"}],\"name\":\"RefundedToOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"Deactivate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"Activate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
	Sigs: map[string]string{
		"0f15f4c0": "activate()",
		"02fb0c5e": "active()",
		"51b42b00": "deactivate()",
		"d0e30db0": "deposit()",
		"83197ef0": "destroy()",
		"f5074f41": "destroyAndSend(address)",
		"8da5cb5b": "owner()",
		"8456cb59": "pause()",
		"5c975abb": "paused()",
		"3e4e10d6": "refundToOwner()",
		"715018a6": "renounceOwnership()",
		"8f134025": "rewardSent(uint256)",
		"2673bd74": "sendReward(uint256,address,uint256)",
		"f2fde38b": "transferOwnership(address)",
		"3f4ba83a": "unpause()",
	},
	Bin: "0x608060408190526000805460a060020a61ffff021916905560208061086e833981016040525160008054600160a060020a03909216600160a060020a031992831633179092169190911790556108148061005a6000396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302fb0c5e81146100df5780630f15f4c0146101085780632673bd741461011f5780633e4e10d6146101465780633f4ba83a1461015b57806351b42b00146101705780635c975abb14610185578063715018a61461019a57806383197ef0146101af5780638456cb59146101c45780638da5cb5b146101d95780638f1340251461020a578063d0e30db014610222578063f2fde38b1461022a578063f5074f411461024b575b600080fd5b3480156100eb57600080fd5b506100f461026c565b604080519115158252519081900360200190f35b34801561011457600080fd5b5061011d61027c565b005b34801561012b57600080fd5b5061011d600435600160a060020a03602435166044356102f9565b34801561015257600080fd5b5061011d61040f565b34801561016757600080fd5b5061011d6104c3565b34801561017c57600080fd5b5061011d610539565b34801561019157600080fd5b506100f4610589565b3480156101a657600080fd5b5061011d610599565b3480156101bb57600080fd5b5061011d610605565b3480156101d057600080fd5b5061011d61062a565b3480156101e557600080fd5b506101ee6106a5565b60408051600160a060020a039092168252519081900360200190f35b34801561021657600080fd5b506100f46004356106b4565b61011d6106c9565b34801561023657600080fd5b5061011d600160a060020a0360043516610725565b34801561025757600080fd5b5061011d600160a060020a0360043516610748565b60005460a860020a900460ff1681565b60005460a860020a900460ff161561029357600080fd5b600054600160a060020a031633146102aa57600080fd5b6000805475ff000000000000000000000000000000000000000000191660a860020a17815560405133917ff7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba1591a2565b600054600160a060020a0316331461031057600080fd5b60008181526001602052604090205460ff161561032c57600080fd5b6000831161033957600080fd5b303183111561034757600080fd5b600160a060020a038216151561035c57600080fd5b600054600160a060020a038381169116141561037757600080fd5b6000818152600160208190526040808320805460ff191690921790915551600160a060020a0384169185156108fc02918691818181858888f193505050501580156103c6573d6000803e3d6000fd5b5060408051848152602081018390528151600160a060020a038516927f6379339f0ae63e95e65fad18ca2a7ec4e7e3f05f3cc5f7079f4d8da8cec34faa928290030190a2505050565b6000805460a860020a900460ff161561042757600080fd5b600054600160a060020a0316331461043e57600080fd5b600030311161044c57600080fd5b5060008054604051303192600160a060020a03909216916108fc841502918491818181858888f19350505050158015610489573d6000803e3d6000fd5b5060408051828152905133917fa9304f0ac97f820b11d5a66f29245fb1e76b1f1f93f46a84a71682e69d095e80919081900360200190a250565b600054600160a060020a031633146104da57600080fd5b60005460a060020a900460ff1615156104f257600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a1565b60005460a860020a900460ff16151561055157600080fd5b600054600160a060020a0316331461056857600080fd5b6000805475ff00000000000000000000000000000000000000000019169055565b60005460a060020a900460ff1681565b600054600160a060020a031633146105b057600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a0316331461061c57600080fd5b600054600160a060020a0316ff5b600054600160a060020a0316331461064157600080fd5b60005460a060020a900460ff161561065857600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a1565b600054600160a060020a031681565b60016020526000908152604090205460ff1681565b60005460a060020a900460ff16156106e057600080fd5b600034116106ed57600080fd5b60408051348152905133917f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4919081900360200190a2565b600054600160a060020a0316331461073c57600080fd5b6107458161076b565b50565b600054600160a060020a0316331461075f57600080fd5b80600160a060020a0316ff5b600160a060020a038116151561078057600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582022f6f45ceff2496c45ad71fcdcb399958432aeee6ab12f4b302eff083c294bd70029",
}

// RoomABI is the input ABI used to generate the binding from.
// Deprecated: Use RoomMetaData.ABI instead.
var RoomABI = RoomMetaData.ABI

// Deprecated: Use RoomMetaData.Sigs instead.
// RoomFuncSigs maps the 4-byte function signature to its string representation.
var RoomFuncSigs = RoomMetaData.Sigs

// RoomBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RoomMetaData.Bin instead.
var RoomBin = RoomMetaData.Bin

// DeployRoom deploys a new Ethereum contract, binding an instance of Room to it.
func DeployRoom(auth *bind.TransactOpts, backend bind.ContractBackend, _creator common.Address) (common.Address, *types.Transaction, *Room, error) {
	parsed, err := RoomMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RoomBin), backend, _creator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Room{RoomCaller: RoomCaller{contract: contract}, RoomTransactor: RoomTransactor{contract: contract}, RoomFilterer: RoomFilterer{contract: contract}}, nil
}

// Room is an auto generated Go binding around an Ethereum contract.
type Room struct {
	RoomCaller     // Read-only binding to the contract
	RoomTransactor // Write-only binding to the contract
	RoomFilterer   // Log filterer for contract events
}

// RoomCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoomCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoomTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoomFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoomSession struct {
	Contract     *Room             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoomCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoomCallerSession struct {
	Contract *RoomCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RoomTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoomTransactorSession struct {
	Contract     *RoomTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoomRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoomRaw struct {
	Contract *Room // Generic contract binding to access the raw methods on
}

// RoomCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoomCallerRaw struct {
	Contract *RoomCaller // Generic read-only contract binding to access the raw methods on
}

// RoomTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoomTransactorRaw struct {
	Contract *RoomTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoom creates a new instance of Room, bound to a specific deployed contract.
func NewRoom(address common.Address, backend bind.ContractBackend) (*Room, error) {
	contract, err := bindRoom(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Room{RoomCaller: RoomCaller{contract: contract}, RoomTransactor: RoomTransactor{contract: contract}, RoomFilterer: RoomFilterer{contract: contract}}, nil
}

// NewRoomCaller creates a new read-only instance of Room, bound to a specific deployed contract.
func NewRoomCaller(address common.Address, caller bind.ContractCaller) (*RoomCaller, error) {
	contract, err := bindRoom(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoomCaller{contract: contract}, nil
}

// NewRoomTransactor creates a new write-only instance of Room, bound to a specific deployed contract.
func NewRoomTransactor(address common.Address, transactor bind.ContractTransactor) (*RoomTransactor, error) {
	contract, err := bindRoom(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoomTransactor{contract: contract}, nil
}

// NewRoomFilterer creates a new log filterer instance of Room, bound to a specific deployed contract.
func NewRoomFilterer(address common.Address, filterer bind.ContractFilterer) (*RoomFilterer, error) {
	contract, err := bindRoom(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoomFilterer{contract: contract}, nil
}

// bindRoom binds a generic wrapper to an already deployed contract.
func bindRoom(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoomABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Room *RoomRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Room.Contract.RoomCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Room *RoomRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.Contract.RoomTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Room *RoomRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Room.Contract.RoomTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Room *RoomCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Room.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Room *RoomTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Room *RoomTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Room.Contract.contract.Transact(opts, method, params...)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Room *RoomCaller) Active(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Room.contract.Call(opts, &out, "active")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Room *RoomSession) Active() (bool, error) {
	return _Room.Contract.Active(&_Room.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() view returns(bool)
func (_Room *RoomCallerSession) Active() (bool, error) {
	return _Room.Contract.Active(&_Room.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Room *RoomCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Room.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Room *RoomSession) Owner() (common.Address, error) {
	return _Room.Contract.Owner(&_Room.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Room *RoomCallerSession) Owner() (common.Address, error) {
	return _Room.Contract.Owner(&_Room.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Room *RoomCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Room.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Room *RoomSession) Paused() (bool, error) {
	return _Room.Contract.Paused(&_Room.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Room *RoomCallerSession) Paused() (bool, error) {
	return _Room.Contract.Paused(&_Room.CallOpts)
}

// RewardSent is a free data retrieval call binding the contract method 0x8f134025.
//
// Solidity: function rewardSent(uint256 ) view returns(bool)
func (_Room *RoomCaller) RewardSent(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Room.contract.Call(opts, &out, "rewardSent", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardSent is a free data retrieval call binding the contract method 0x8f134025.
//
// Solidity: function rewardSent(uint256 ) view returns(bool)
func (_Room *RoomSession) RewardSent(arg0 *big.Int) (bool, error) {
	return _Room.Contract.RewardSent(&_Room.CallOpts, arg0)
}

// RewardSent is a free data retrieval call binding the contract method 0x8f134025.
//
// Solidity: function rewardSent(uint256 ) view returns(bool)
func (_Room *RoomCallerSession) RewardSent(arg0 *big.Int) (bool, error) {
	return _Room.Contract.RewardSent(&_Room.CallOpts, arg0)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Room *RoomTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Room *RoomSession) Activate() (*types.Transaction, error) {
	return _Room.Contract.Activate(&_Room.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Room *RoomTransactorSession) Activate() (*types.Transaction, error) {
	return _Room.Contract.Activate(&_Room.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Room *RoomTransactor) Deactivate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "deactivate")
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Room *RoomSession) Deactivate() (*types.Transaction, error) {
	return _Room.Contract.Deactivate(&_Room.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Room *RoomTransactorSession) Deactivate() (*types.Transaction, error) {
	return _Room.Contract.Deactivate(&_Room.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Room *RoomTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Room *RoomSession) Deposit() (*types.Transaction, error) {
	return _Room.Contract.Deposit(&_Room.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Room *RoomTransactorSession) Deposit() (*types.Transaction, error) {
	return _Room.Contract.Deposit(&_Room.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Room *RoomTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Room *RoomSession) Destroy() (*types.Transaction, error) {
	return _Room.Contract.Destroy(&_Room.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Room *RoomTransactorSession) Destroy() (*types.Transaction, error) {
	return _Room.Contract.Destroy(&_Room.TransactOpts)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_Room *RoomTransactor) DestroyAndSend(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "destroyAndSend", _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_Room *RoomSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _Room.Contract.DestroyAndSend(&_Room.TransactOpts, _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_Room *RoomTransactorSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _Room.Contract.DestroyAndSend(&_Room.TransactOpts, _recipient)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Room *RoomTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Room *RoomSession) Pause() (*types.Transaction, error) {
	return _Room.Contract.Pause(&_Room.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Room *RoomTransactorSession) Pause() (*types.Transaction, error) {
	return _Room.Contract.Pause(&_Room.TransactOpts)
}

// RefundToOwner is a paid mutator transaction binding the contract method 0x3e4e10d6.
//
// Solidity: function refundToOwner() returns()
func (_Room *RoomTransactor) RefundToOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "refundToOwner")
}

// RefundToOwner is a paid mutator transaction binding the contract method 0x3e4e10d6.
//
// Solidity: function refundToOwner() returns()
func (_Room *RoomSession) RefundToOwner() (*types.Transaction, error) {
	return _Room.Contract.RefundToOwner(&_Room.TransactOpts)
}

// RefundToOwner is a paid mutator transaction binding the contract method 0x3e4e10d6.
//
// Solidity: function refundToOwner() returns()
func (_Room *RoomTransactorSession) RefundToOwner() (*types.Transaction, error) {
	return _Room.Contract.RefundToOwner(&_Room.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Room *RoomTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Room *RoomSession) RenounceOwnership() (*types.Transaction, error) {
	return _Room.Contract.RenounceOwnership(&_Room.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Room *RoomTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Room.Contract.RenounceOwnership(&_Room.TransactOpts)
}

// SendReward is a paid mutator transaction binding the contract method 0x2673bd74.
//
// Solidity: function sendReward(uint256 _reward, address _dest, uint256 _id) returns()
func (_Room *RoomTransactor) SendReward(opts *bind.TransactOpts, _reward *big.Int, _dest common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "sendReward", _reward, _dest, _id)
}

// SendReward is a paid mutator transaction binding the contract method 0x2673bd74.
//
// Solidity: function sendReward(uint256 _reward, address _dest, uint256 _id) returns()
func (_Room *RoomSession) SendReward(_reward *big.Int, _dest common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Room.Contract.SendReward(&_Room.TransactOpts, _reward, _dest, _id)
}

// SendReward is a paid mutator transaction binding the contract method 0x2673bd74.
//
// Solidity: function sendReward(uint256 _reward, address _dest, uint256 _id) returns()
func (_Room *RoomTransactorSession) SendReward(_reward *big.Int, _dest common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Room.Contract.SendReward(&_Room.TransactOpts, _reward, _dest, _id)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Room *RoomTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Room *RoomSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Room.Contract.TransferOwnership(&_Room.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Room *RoomTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Room.Contract.TransferOwnership(&_Room.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Room *RoomTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Room *RoomSession) Unpause() (*types.Transaction, error) {
	return _Room.Contract.Unpause(&_Room.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Room *RoomTransactorSession) Unpause() (*types.Transaction, error) {
	return _Room.Contract.Unpause(&_Room.TransactOpts)
}

// RoomActivateIterator is returned from FilterActivate and is used to iterate over the raw logs and unpacked data for Activate events raised by the Room contract.
type RoomActivateIterator struct {
	Event *RoomActivate // Event containing the contract specifics and raw log

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
func (it *RoomActivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomActivate)
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
		it.Event = new(RoomActivate)
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
func (it *RoomActivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomActivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomActivate represents a Activate event raised by the Room contract.
type RoomActivate struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterActivate is a free log retrieval operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: event Activate(address indexed _sender)
func (_Room *RoomFilterer) FilterActivate(opts *bind.FilterOpts, _sender []common.Address) (*RoomActivateIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "Activate", _senderRule)
	if err != nil {
		return nil, err
	}
	return &RoomActivateIterator{contract: _Room.contract, event: "Activate", logs: logs, sub: sub}, nil
}

// WatchActivate is a free log subscription operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: event Activate(address indexed _sender)
func (_Room *RoomFilterer) WatchActivate(opts *bind.WatchOpts, sink chan<- *RoomActivate, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "Activate", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomActivate)
				if err := _Room.contract.UnpackLog(event, "Activate", log); err != nil {
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

// ParseActivate is a log parse operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: event Activate(address indexed _sender)
func (_Room *RoomFilterer) ParseActivate(log types.Log) (*RoomActivate, error) {
	event := new(RoomActivate)
	if err := _Room.contract.UnpackLog(event, "Activate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomDeactivateIterator is returned from FilterDeactivate and is used to iterate over the raw logs and unpacked data for Deactivate events raised by the Room contract.
type RoomDeactivateIterator struct {
	Event *RoomDeactivate // Event containing the contract specifics and raw log

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
func (it *RoomDeactivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomDeactivate)
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
		it.Event = new(RoomDeactivate)
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
func (it *RoomDeactivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomDeactivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomDeactivate represents a Deactivate event raised by the Room contract.
type RoomDeactivate struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeactivate is a free log retrieval operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: event Deactivate(address indexed _sender)
func (_Room *RoomFilterer) FilterDeactivate(opts *bind.FilterOpts, _sender []common.Address) (*RoomDeactivateIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "Deactivate", _senderRule)
	if err != nil {
		return nil, err
	}
	return &RoomDeactivateIterator{contract: _Room.contract, event: "Deactivate", logs: logs, sub: sub}, nil
}

// WatchDeactivate is a free log subscription operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: event Deactivate(address indexed _sender)
func (_Room *RoomFilterer) WatchDeactivate(opts *bind.WatchOpts, sink chan<- *RoomDeactivate, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "Deactivate", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomDeactivate)
				if err := _Room.contract.UnpackLog(event, "Deactivate", log); err != nil {
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

// ParseDeactivate is a log parse operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: event Deactivate(address indexed _sender)
func (_Room *RoomFilterer) ParseDeactivate(log types.Log) (*RoomDeactivate, error) {
	event := new(RoomDeactivate)
	if err := _Room.contract.UnpackLog(event, "Deactivate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Room contract.
type RoomDepositedIterator struct {
	Event *RoomDeposited // Event containing the contract specifics and raw log

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
func (it *RoomDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomDeposited)
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
		it.Event = new(RoomDeposited)
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
func (it *RoomDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomDeposited represents a Deposited event raised by the Room contract.
type RoomDeposited struct {
	Depositor      common.Address
	DepositedValue *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed _depositor, uint256 _depositedValue)
func (_Room *RoomFilterer) FilterDeposited(opts *bind.FilterOpts, _depositor []common.Address) (*RoomDepositedIterator, error) {

	var _depositorRule []interface{}
	for _, _depositorItem := range _depositor {
		_depositorRule = append(_depositorRule, _depositorItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "Deposited", _depositorRule)
	if err != nil {
		return nil, err
	}
	return &RoomDepositedIterator{contract: _Room.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed _depositor, uint256 _depositedValue)
func (_Room *RoomFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *RoomDeposited, _depositor []common.Address) (event.Subscription, error) {

	var _depositorRule []interface{}
	for _, _depositorItem := range _depositor {
		_depositorRule = append(_depositorRule, _depositorItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "Deposited", _depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomDeposited)
				if err := _Room.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed _depositor, uint256 _depositedValue)
func (_Room *RoomFilterer) ParseDeposited(log types.Log) (*RoomDeposited, error) {
	event := new(RoomDeposited)
	if err := _Room.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Room contract.
type RoomOwnershipRenouncedIterator struct {
	Event *RoomOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RoomOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomOwnershipRenounced)
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
		it.Event = new(RoomOwnershipRenounced)
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
func (it *RoomOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomOwnershipRenounced represents a OwnershipRenounced event raised by the Room contract.
type RoomOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Room *RoomFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RoomOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RoomOwnershipRenouncedIterator{contract: _Room.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Room *RoomFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RoomOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomOwnershipRenounced)
				if err := _Room.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Room *RoomFilterer) ParseOwnershipRenounced(log types.Log) (*RoomOwnershipRenounced, error) {
	event := new(RoomOwnershipRenounced)
	if err := _Room.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Room contract.
type RoomOwnershipTransferredIterator struct {
	Event *RoomOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RoomOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomOwnershipTransferred)
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
		it.Event = new(RoomOwnershipTransferred)
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
func (it *RoomOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomOwnershipTransferred represents a OwnershipTransferred event raised by the Room contract.
type RoomOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Room *RoomFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RoomOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RoomOwnershipTransferredIterator{contract: _Room.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Room *RoomFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RoomOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomOwnershipTransferred)
				if err := _Room.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Room *RoomFilterer) ParseOwnershipTransferred(log types.Log) (*RoomOwnershipTransferred, error) {
	event := new(RoomOwnershipTransferred)
	if err := _Room.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Room contract.
type RoomPauseIterator struct {
	Event *RoomPause // Event containing the contract specifics and raw log

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
func (it *RoomPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomPause)
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
		it.Event = new(RoomPause)
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
func (it *RoomPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomPause represents a Pause event raised by the Room contract.
type RoomPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Room *RoomFilterer) FilterPause(opts *bind.FilterOpts) (*RoomPauseIterator, error) {

	logs, sub, err := _Room.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &RoomPauseIterator{contract: _Room.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Room *RoomFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *RoomPause) (event.Subscription, error) {

	logs, sub, err := _Room.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomPause)
				if err := _Room.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Room *RoomFilterer) ParsePause(log types.Log) (*RoomPause, error) {
	event := new(RoomPause)
	if err := _Room.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomRefundedToOwnerIterator is returned from FilterRefundedToOwner and is used to iterate over the raw logs and unpacked data for RefundedToOwner events raised by the Room contract.
type RoomRefundedToOwnerIterator struct {
	Event *RoomRefundedToOwner // Event containing the contract specifics and raw log

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
func (it *RoomRefundedToOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomRefundedToOwner)
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
		it.Event = new(RoomRefundedToOwner)
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
func (it *RoomRefundedToOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomRefundedToOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomRefundedToOwner represents a RefundedToOwner event raised by the Room contract.
type RoomRefundedToOwner struct {
	Dest            common.Address
	RefundedBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRefundedToOwner is a free log retrieval operation binding the contract event 0xa9304f0ac97f820b11d5a66f29245fb1e76b1f1f93f46a84a71682e69d095e80.
//
// Solidity: event RefundedToOwner(address indexed _dest, uint256 _refundedBalance)
func (_Room *RoomFilterer) FilterRefundedToOwner(opts *bind.FilterOpts, _dest []common.Address) (*RoomRefundedToOwnerIterator, error) {

	var _destRule []interface{}
	for _, _destItem := range _dest {
		_destRule = append(_destRule, _destItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "RefundedToOwner", _destRule)
	if err != nil {
		return nil, err
	}
	return &RoomRefundedToOwnerIterator{contract: _Room.contract, event: "RefundedToOwner", logs: logs, sub: sub}, nil
}

// WatchRefundedToOwner is a free log subscription operation binding the contract event 0xa9304f0ac97f820b11d5a66f29245fb1e76b1f1f93f46a84a71682e69d095e80.
//
// Solidity: event RefundedToOwner(address indexed _dest, uint256 _refundedBalance)
func (_Room *RoomFilterer) WatchRefundedToOwner(opts *bind.WatchOpts, sink chan<- *RoomRefundedToOwner, _dest []common.Address) (event.Subscription, error) {

	var _destRule []interface{}
	for _, _destItem := range _dest {
		_destRule = append(_destRule, _destItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "RefundedToOwner", _destRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomRefundedToOwner)
				if err := _Room.contract.UnpackLog(event, "RefundedToOwner", log); err != nil {
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

// ParseRefundedToOwner is a log parse operation binding the contract event 0xa9304f0ac97f820b11d5a66f29245fb1e76b1f1f93f46a84a71682e69d095e80.
//
// Solidity: event RefundedToOwner(address indexed _dest, uint256 _refundedBalance)
func (_Room *RoomFilterer) ParseRefundedToOwner(log types.Log) (*RoomRefundedToOwner, error) {
	event := new(RoomRefundedToOwner)
	if err := _Room.contract.UnpackLog(event, "RefundedToOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomRewardSentIterator is returned from FilterRewardSent and is used to iterate over the raw logs and unpacked data for RewardSent events raised by the Room contract.
type RoomRewardSentIterator struct {
	Event *RoomRewardSent // Event containing the contract specifics and raw log

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
func (it *RoomRewardSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomRewardSent)
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
		it.Event = new(RoomRewardSent)
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
func (it *RoomRewardSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomRewardSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomRewardSent represents a RewardSent event raised by the Room contract.
type RoomRewardSent struct {
	Dest   common.Address
	Reward *big.Int
	Id     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardSent is a free log retrieval operation binding the contract event 0x6379339f0ae63e95e65fad18ca2a7ec4e7e3f05f3cc5f7079f4d8da8cec34faa.
//
// Solidity: event RewardSent(address indexed _dest, uint256 _reward, uint256 _id)
func (_Room *RoomFilterer) FilterRewardSent(opts *bind.FilterOpts, _dest []common.Address) (*RoomRewardSentIterator, error) {

	var _destRule []interface{}
	for _, _destItem := range _dest {
		_destRule = append(_destRule, _destItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "RewardSent", _destRule)
	if err != nil {
		return nil, err
	}
	return &RoomRewardSentIterator{contract: _Room.contract, event: "RewardSent", logs: logs, sub: sub}, nil
}

// WatchRewardSent is a free log subscription operation binding the contract event 0x6379339f0ae63e95e65fad18ca2a7ec4e7e3f05f3cc5f7079f4d8da8cec34faa.
//
// Solidity: event RewardSent(address indexed _dest, uint256 _reward, uint256 _id)
func (_Room *RoomFilterer) WatchRewardSent(opts *bind.WatchOpts, sink chan<- *RoomRewardSent, _dest []common.Address) (event.Subscription, error) {

	var _destRule []interface{}
	for _, _destItem := range _dest {
		_destRule = append(_destRule, _destItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "RewardSent", _destRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomRewardSent)
				if err := _Room.contract.UnpackLog(event, "RewardSent", log); err != nil {
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

// ParseRewardSent is a log parse operation binding the contract event 0x6379339f0ae63e95e65fad18ca2a7ec4e7e3f05f3cc5f7079f4d8da8cec34faa.
//
// Solidity: event RewardSent(address indexed _dest, uint256 _reward, uint256 _id)
func (_Room *RoomFilterer) ParseRewardSent(log types.Log) (*RoomRewardSent, error) {
	event := new(RoomRewardSent)
	if err := _Room.contract.UnpackLog(event, "RewardSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Room contract.
type RoomUnpauseIterator struct {
	Event *RoomUnpause // Event containing the contract specifics and raw log

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
func (it *RoomUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomUnpause)
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
		it.Event = new(RoomUnpause)
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
func (it *RoomUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomUnpause represents a Unpause event raised by the Room contract.
type RoomUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Room *RoomFilterer) FilterUnpause(opts *bind.FilterOpts) (*RoomUnpauseIterator, error) {

	logs, sub, err := _Room.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &RoomUnpauseIterator{contract: _Room.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Room *RoomFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *RoomUnpause) (event.Subscription, error) {

	logs, sub, err := _Room.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomUnpause)
				if err := _Room.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Room *RoomFilterer) ParseUnpause(log types.Log) (*RoomUnpause, error) {
	event := new(RoomUnpause)
	if err := _Room.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomFactoryMetaData contains all meta data concerning the RoomFactory contract.
var RoomFactoryMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"createRoom\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"destroyAndSend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_room\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_depositedValue\",\"type\":\"uint256\"}],\"name\":\"RoomCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
	Sigs: map[string]string{
		"3be272aa": "createRoom()",
		"83197ef0": "destroy()",
		"f5074f41": "destroyAndSend(address)",
		"8da5cb5b": "owner()",
		"8456cb59": "pause()",
		"5c975abb": "paused()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
		"3f4ba83a": "unpause()",
	},
	Bin: "0x608060405260008054600160a860020a03191633179055610d40806100256000396000f3006080604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633be272aa811461009d5780633f4ba83a146100a75780635c975abb146100bc578063715018a6146100e557806383197ef0146100fa5780638456cb591461010f5780638da5cb5b14610124578063f2fde38b14610155578063f5074f4114610176575b600080fd5b6100a5610197565b005b3480156100b357600080fd5b506100a5610232565b3480156100c857600080fd5b506100d16102a8565b604080519115158252519081900360200190f35b3480156100f157600080fd5b506100a56102b8565b34801561010657600080fd5b506100a5610324565b34801561011b57600080fd5b506100a5610349565b34801561013057600080fd5b506101396103c4565b60408051600160a060020a039092168252519081900360200190f35b34801561016157600080fd5b506100a5600160a060020a03600435166103d3565b34801561018257600080fd5b506100a5600160a060020a03600435166103f6565b6000805460a060020a900460ff16156101af57600080fd5b34336101b9610496565b600160a060020a039091168152604051908190036020019082f0801580156101e5573d6000803e3d6000fd5b5060408051600160a060020a038316815234602082015281519294503393507f6849f7a409ad97d39c5ffa074bf765330bf1a574da99d4c4831196ecd77ea8da929081900390910190a250565b600054600160a060020a0316331461024957600080fd5b60005460a060020a900460ff16151561026157600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a1565b60005460a060020a900460ff1681565b600054600160a060020a031633146102cf57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a0316331461033b57600080fd5b600054600160a060020a0316ff5b600054600160a060020a0316331461036057600080fd5b60005460a060020a900460ff161561037757600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a1565b600054600160a060020a031681565b600054600160a060020a031633146103ea57600080fd5b6103f381610419565b50565b600054600160a060020a0316331461040d57600080fd5b80600160a060020a0316ff5b600160a060020a038116151561042e57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60405161086e806104a7833901905600608060408190526000805460a060020a61ffff021916905560208061086e833981016040525160008054600160a060020a03909216600160a060020a031992831633179092169190911790556108148061005a6000396000f3006080604052600436106100da5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302fb0c5e81146100df5780630f15f4c0146101085780632673bd741461011f5780633e4e10d6146101465780633f4ba83a1461015b57806351b42b00146101705780635c975abb14610185578063715018a61461019a57806383197ef0146101af5780638456cb59146101c45780638da5cb5b146101d95780638f1340251461020a578063d0e30db014610222578063f2fde38b1461022a578063f5074f411461024b575b600080fd5b3480156100eb57600080fd5b506100f461026c565b604080519115158252519081900360200190f35b34801561011457600080fd5b5061011d61027c565b005b34801561012b57600080fd5b5061011d600435600160a060020a03602435166044356102f9565b34801561015257600080fd5b5061011d61040f565b34801561016757600080fd5b5061011d6104c3565b34801561017c57600080fd5b5061011d610539565b34801561019157600080fd5b506100f4610589565b3480156101a657600080fd5b5061011d610599565b3480156101bb57600080fd5b5061011d610605565b3480156101d057600080fd5b5061011d61062a565b3480156101e557600080fd5b506101ee6106a5565b60408051600160a060020a039092168252519081900360200190f35b34801561021657600080fd5b506100f46004356106b4565b61011d6106c9565b34801561023657600080fd5b5061011d600160a060020a0360043516610725565b34801561025757600080fd5b5061011d600160a060020a0360043516610748565b60005460a860020a900460ff1681565b60005460a860020a900460ff161561029357600080fd5b600054600160a060020a031633146102aa57600080fd5b6000805475ff000000000000000000000000000000000000000000191660a860020a17815560405133917ff7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba1591a2565b600054600160a060020a0316331461031057600080fd5b60008181526001602052604090205460ff161561032c57600080fd5b6000831161033957600080fd5b303183111561034757600080fd5b600160a060020a038216151561035c57600080fd5b600054600160a060020a038381169116141561037757600080fd5b6000818152600160208190526040808320805460ff191690921790915551600160a060020a0384169185156108fc02918691818181858888f193505050501580156103c6573d6000803e3d6000fd5b5060408051848152602081018390528151600160a060020a038516927f6379339f0ae63e95e65fad18ca2a7ec4e7e3f05f3cc5f7079f4d8da8cec34faa928290030190a2505050565b6000805460a860020a900460ff161561042757600080fd5b600054600160a060020a0316331461043e57600080fd5b600030311161044c57600080fd5b5060008054604051303192600160a060020a03909216916108fc841502918491818181858888f19350505050158015610489573d6000803e3d6000fd5b5060408051828152905133917fa9304f0ac97f820b11d5a66f29245fb1e76b1f1f93f46a84a71682e69d095e80919081900360200190a250565b600054600160a060020a031633146104da57600080fd5b60005460a060020a900460ff1615156104f257600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a1565b60005460a860020a900460ff16151561055157600080fd5b600054600160a060020a0316331461056857600080fd5b6000805475ff00000000000000000000000000000000000000000019169055565b60005460a060020a900460ff1681565b600054600160a060020a031633146105b057600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a0316331461061c57600080fd5b600054600160a060020a0316ff5b600054600160a060020a0316331461064157600080fd5b60005460a060020a900460ff161561065857600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a1565b600054600160a060020a031681565b60016020526000908152604090205460ff1681565b60005460a060020a900460ff16156106e057600080fd5b600034116106ed57600080fd5b60408051348152905133917f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4919081900360200190a2565b600054600160a060020a0316331461073c57600080fd5b6107458161076b565b50565b600054600160a060020a0316331461075f57600080fd5b80600160a060020a0316ff5b600160a060020a038116151561078057600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582022f6f45ceff2496c45ad71fcdcb399958432aeee6ab12f4b302eff083c294bd70029a165627a7a7230582065ba6fa71c51704a36a8ccbe4b0fcd705de4afc22f49cc01de70cc666777cdc70029",
}

// RoomFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use RoomFactoryMetaData.ABI instead.
var RoomFactoryABI = RoomFactoryMetaData.ABI

// Deprecated: Use RoomFactoryMetaData.Sigs instead.
// RoomFactoryFuncSigs maps the 4-byte function signature to its string representation.
var RoomFactoryFuncSigs = RoomFactoryMetaData.Sigs

// RoomFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RoomFactoryMetaData.Bin instead.
var RoomFactoryBin = RoomFactoryMetaData.Bin

// DeployRoomFactory deploys a new Ethereum contract, binding an instance of RoomFactory to it.
func DeployRoomFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RoomFactory, error) {
	parsed, err := RoomFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RoomFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RoomFactory{RoomFactoryCaller: RoomFactoryCaller{contract: contract}, RoomFactoryTransactor: RoomFactoryTransactor{contract: contract}, RoomFactoryFilterer: RoomFactoryFilterer{contract: contract}}, nil
}

// RoomFactory is an auto generated Go binding around an Ethereum contract.
type RoomFactory struct {
	RoomFactoryCaller     // Read-only binding to the contract
	RoomFactoryTransactor // Write-only binding to the contract
	RoomFactoryFilterer   // Log filterer for contract events
}

// RoomFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoomFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoomFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoomFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoomFactorySession struct {
	Contract     *RoomFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoomFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoomFactoryCallerSession struct {
	Contract *RoomFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RoomFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoomFactoryTransactorSession struct {
	Contract     *RoomFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RoomFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoomFactoryRaw struct {
	Contract *RoomFactory // Generic contract binding to access the raw methods on
}

// RoomFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoomFactoryCallerRaw struct {
	Contract *RoomFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// RoomFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoomFactoryTransactorRaw struct {
	Contract *RoomFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoomFactory creates a new instance of RoomFactory, bound to a specific deployed contract.
func NewRoomFactory(address common.Address, backend bind.ContractBackend) (*RoomFactory, error) {
	contract, err := bindRoomFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoomFactory{RoomFactoryCaller: RoomFactoryCaller{contract: contract}, RoomFactoryTransactor: RoomFactoryTransactor{contract: contract}, RoomFactoryFilterer: RoomFactoryFilterer{contract: contract}}, nil
}

// NewRoomFactoryCaller creates a new read-only instance of RoomFactory, bound to a specific deployed contract.
func NewRoomFactoryCaller(address common.Address, caller bind.ContractCaller) (*RoomFactoryCaller, error) {
	contract, err := bindRoomFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoomFactoryCaller{contract: contract}, nil
}

// NewRoomFactoryTransactor creates a new write-only instance of RoomFactory, bound to a specific deployed contract.
func NewRoomFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*RoomFactoryTransactor, error) {
	contract, err := bindRoomFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoomFactoryTransactor{contract: contract}, nil
}

// NewRoomFactoryFilterer creates a new log filterer instance of RoomFactory, bound to a specific deployed contract.
func NewRoomFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*RoomFactoryFilterer, error) {
	contract, err := bindRoomFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoomFactoryFilterer{contract: contract}, nil
}

// bindRoomFactory binds a generic wrapper to an already deployed contract.
func bindRoomFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoomFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoomFactory *RoomFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoomFactory.Contract.RoomFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoomFactory *RoomFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomFactory.Contract.RoomFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoomFactory *RoomFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoomFactory.Contract.RoomFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoomFactory *RoomFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoomFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoomFactory *RoomFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoomFactory *RoomFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoomFactory.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoomFactory *RoomFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoomFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoomFactory *RoomFactorySession) Owner() (common.Address, error) {
	return _RoomFactory.Contract.Owner(&_RoomFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoomFactory *RoomFactoryCallerSession) Owner() (common.Address, error) {
	return _RoomFactory.Contract.Owner(&_RoomFactory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RoomFactory *RoomFactoryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RoomFactory.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RoomFactory *RoomFactorySession) Paused() (bool, error) {
	return _RoomFactory.Contract.Paused(&_RoomFactory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RoomFactory *RoomFactoryCallerSession) Paused() (bool, error) {
	return _RoomFactory.Contract.Paused(&_RoomFactory.CallOpts)
}

// CreateRoom is a paid mutator transaction binding the contract method 0x3be272aa.
//
// Solidity: function createRoom() payable returns()
func (_RoomFactory *RoomFactoryTransactor) CreateRoom(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomFactory.contract.Transact(opts, "createRoom")
}

// CreateRoom is a paid mutator transaction binding the contract method 0x3be272aa.
//
// Solidity: function createRoom() payable returns()
func (_RoomFactory *RoomFactorySession) CreateRoom() (*types.Transaction, error) {
	return _RoomFactory.Contract.CreateRoom(&_RoomFactory.TransactOpts)
}

// CreateRoom is a paid mutator transaction binding the contract method 0x3be272aa.
//
// Solidity: function createRoom() payable returns()
func (_RoomFactory *RoomFactoryTransactorSession) CreateRoom() (*types.Transaction, error) {
	return _RoomFactory.Contract.CreateRoom(&_RoomFactory.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_RoomFactory *RoomFactoryTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomFactory.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_RoomFactory *RoomFactorySession) Destroy() (*types.Transaction, error) {
	return _RoomFactory.Contract.Destroy(&_RoomFactory.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_RoomFactory *RoomFactoryTransactorSession) Destroy() (*types.Transaction, error) {
	return _RoomFactory.Contract.Destroy(&_RoomFactory.TransactOpts)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_RoomFactory *RoomFactoryTransactor) DestroyAndSend(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _RoomFactory.contract.Transact(opts, "destroyAndSend", _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_RoomFactory *RoomFactorySession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _RoomFactory.Contract.DestroyAndSend(&_RoomFactory.TransactOpts, _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(address _recipient) returns()
func (_RoomFactory *RoomFactoryTransactorSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _RoomFactory.Contract.DestroyAndSend(&_RoomFactory.TransactOpts, _recipient)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RoomFactory *RoomFactoryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomFactory.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RoomFactory *RoomFactorySession) Pause() (*types.Transaction, error) {
	return _RoomFactory.Contract.Pause(&_RoomFactory.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RoomFactory *RoomFactoryTransactorSession) Pause() (*types.Transaction, error) {
	return _RoomFactory.Contract.Pause(&_RoomFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RoomFactory *RoomFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RoomFactory *RoomFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _RoomFactory.Contract.RenounceOwnership(&_RoomFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RoomFactory *RoomFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RoomFactory.Contract.RenounceOwnership(&_RoomFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_RoomFactory *RoomFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RoomFactory.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_RoomFactory *RoomFactorySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RoomFactory.Contract.TransferOwnership(&_RoomFactory.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_RoomFactory *RoomFactoryTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RoomFactory.Contract.TransferOwnership(&_RoomFactory.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RoomFactory *RoomFactoryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomFactory.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RoomFactory *RoomFactorySession) Unpause() (*types.Transaction, error) {
	return _RoomFactory.Contract.Unpause(&_RoomFactory.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RoomFactory *RoomFactoryTransactorSession) Unpause() (*types.Transaction, error) {
	return _RoomFactory.Contract.Unpause(&_RoomFactory.TransactOpts)
}

// RoomFactoryOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the RoomFactory contract.
type RoomFactoryOwnershipRenouncedIterator struct {
	Event *RoomFactoryOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RoomFactoryOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomFactoryOwnershipRenounced)
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
		it.Event = new(RoomFactoryOwnershipRenounced)
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
func (it *RoomFactoryOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomFactoryOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomFactoryOwnershipRenounced represents a OwnershipRenounced event raised by the RoomFactory contract.
type RoomFactoryOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_RoomFactory *RoomFactoryFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RoomFactoryOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RoomFactory.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RoomFactoryOwnershipRenouncedIterator{contract: _RoomFactory.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_RoomFactory *RoomFactoryFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RoomFactoryOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RoomFactory.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomFactoryOwnershipRenounced)
				if err := _RoomFactory.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_RoomFactory *RoomFactoryFilterer) ParseOwnershipRenounced(log types.Log) (*RoomFactoryOwnershipRenounced, error) {
	event := new(RoomFactoryOwnershipRenounced)
	if err := _RoomFactory.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RoomFactory contract.
type RoomFactoryOwnershipTransferredIterator struct {
	Event *RoomFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RoomFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomFactoryOwnershipTransferred)
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
		it.Event = new(RoomFactoryOwnershipTransferred)
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
func (it *RoomFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the RoomFactory contract.
type RoomFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RoomFactory *RoomFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RoomFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RoomFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RoomFactoryOwnershipTransferredIterator{contract: _RoomFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RoomFactory *RoomFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RoomFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RoomFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomFactoryOwnershipTransferred)
				if err := _RoomFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RoomFactory *RoomFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*RoomFactoryOwnershipTransferred, error) {
	event := new(RoomFactoryOwnershipTransferred)
	if err := _RoomFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomFactoryPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the RoomFactory contract.
type RoomFactoryPauseIterator struct {
	Event *RoomFactoryPause // Event containing the contract specifics and raw log

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
func (it *RoomFactoryPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomFactoryPause)
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
		it.Event = new(RoomFactoryPause)
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
func (it *RoomFactoryPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomFactoryPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomFactoryPause represents a Pause event raised by the RoomFactory contract.
type RoomFactoryPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_RoomFactory *RoomFactoryFilterer) FilterPause(opts *bind.FilterOpts) (*RoomFactoryPauseIterator, error) {

	logs, sub, err := _RoomFactory.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &RoomFactoryPauseIterator{contract: _RoomFactory.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_RoomFactory *RoomFactoryFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *RoomFactoryPause) (event.Subscription, error) {

	logs, sub, err := _RoomFactory.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomFactoryPause)
				if err := _RoomFactory.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_RoomFactory *RoomFactoryFilterer) ParsePause(log types.Log) (*RoomFactoryPause, error) {
	event := new(RoomFactoryPause)
	if err := _RoomFactory.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomFactoryRoomCreatedIterator is returned from FilterRoomCreated and is used to iterate over the raw logs and unpacked data for RoomCreated events raised by the RoomFactory contract.
type RoomFactoryRoomCreatedIterator struct {
	Event *RoomFactoryRoomCreated // Event containing the contract specifics and raw log

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
func (it *RoomFactoryRoomCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomFactoryRoomCreated)
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
		it.Event = new(RoomFactoryRoomCreated)
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
func (it *RoomFactoryRoomCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomFactoryRoomCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomFactoryRoomCreated represents a RoomCreated event raised by the RoomFactory contract.
type RoomFactoryRoomCreated struct {
	Creator        common.Address
	Room           common.Address
	DepositedValue *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRoomCreated is a free log retrieval operation binding the contract event 0x6849f7a409ad97d39c5ffa074bf765330bf1a574da99d4c4831196ecd77ea8da.
//
// Solidity: event RoomCreated(address indexed _creator, address _room, uint256 _depositedValue)
func (_RoomFactory *RoomFactoryFilterer) FilterRoomCreated(opts *bind.FilterOpts, _creator []common.Address) (*RoomFactoryRoomCreatedIterator, error) {

	var _creatorRule []interface{}
	for _, _creatorItem := range _creator {
		_creatorRule = append(_creatorRule, _creatorItem)
	}

	logs, sub, err := _RoomFactory.contract.FilterLogs(opts, "RoomCreated", _creatorRule)
	if err != nil {
		return nil, err
	}
	return &RoomFactoryRoomCreatedIterator{contract: _RoomFactory.contract, event: "RoomCreated", logs: logs, sub: sub}, nil
}

// WatchRoomCreated is a free log subscription operation binding the contract event 0x6849f7a409ad97d39c5ffa074bf765330bf1a574da99d4c4831196ecd77ea8da.
//
// Solidity: event RoomCreated(address indexed _creator, address _room, uint256 _depositedValue)
func (_RoomFactory *RoomFactoryFilterer) WatchRoomCreated(opts *bind.WatchOpts, sink chan<- *RoomFactoryRoomCreated, _creator []common.Address) (event.Subscription, error) {

	var _creatorRule []interface{}
	for _, _creatorItem := range _creator {
		_creatorRule = append(_creatorRule, _creatorItem)
	}

	logs, sub, err := _RoomFactory.contract.WatchLogs(opts, "RoomCreated", _creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomFactoryRoomCreated)
				if err := _RoomFactory.contract.UnpackLog(event, "RoomCreated", log); err != nil {
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

// ParseRoomCreated is a log parse operation binding the contract event 0x6849f7a409ad97d39c5ffa074bf765330bf1a574da99d4c4831196ecd77ea8da.
//
// Solidity: event RoomCreated(address indexed _creator, address _room, uint256 _depositedValue)
func (_RoomFactory *RoomFactoryFilterer) ParseRoomCreated(log types.Log) (*RoomFactoryRoomCreated, error) {
	event := new(RoomFactoryRoomCreated)
	if err := _RoomFactory.contract.UnpackLog(event, "RoomCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoomFactoryUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the RoomFactory contract.
type RoomFactoryUnpauseIterator struct {
	Event *RoomFactoryUnpause // Event containing the contract specifics and raw log

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
func (it *RoomFactoryUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomFactoryUnpause)
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
		it.Event = new(RoomFactoryUnpause)
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
func (it *RoomFactoryUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomFactoryUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomFactoryUnpause represents a Unpause event raised by the RoomFactory contract.
type RoomFactoryUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_RoomFactory *RoomFactoryFilterer) FilterUnpause(opts *bind.FilterOpts) (*RoomFactoryUnpauseIterator, error) {

	logs, sub, err := _RoomFactory.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &RoomFactoryUnpauseIterator{contract: _RoomFactory.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_RoomFactory *RoomFactoryFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *RoomFactoryUnpause) (event.Subscription, error) {

	logs, sub, err := _RoomFactory.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomFactoryUnpause)
				if err := _RoomFactory.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_RoomFactory *RoomFactoryFilterer) ParseUnpause(log types.Log) (*RoomFactoryUnpause, error) {
	event := new(RoomFactoryUnpause)
	if err := _RoomFactory.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
