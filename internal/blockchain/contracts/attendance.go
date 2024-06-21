// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// AttendanceAttendanceRecord is an auto generated low-level Go binding around an user-defined struct.
type AttendanceAttendanceRecord struct {
	Id          *big.Int
	EmployerId  *big.Int
	CheckInTime *big.Int
	Notes       string
}

// AttendanceMetaData contains all meta data concerning the Attendance contract.
var AttendanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"checkInTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"notes\",\"type\":\"string\"}],\"name\":\"AttendanceRecorded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"checkInTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"notes\",\"type\":\"string\"}],\"name\":\"AttendanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"entity\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAuthorized\",\"type\":\"bool\"}],\"name\":\"AuthorizationUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"entity\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAuthorized\",\"type\":\"bool\"}],\"name\":\"authorizeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"employerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toDate\",\"type\":\"uint256\"}],\"name\":\"getAttendance\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"employerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkInTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"notes\",\"type\":\"string\"}],\"internalType\":\"structAttendance.AttendanceRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkInTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"notes\",\"type\":\"string\"}],\"name\":\"recordAttendance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"employerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkInTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"notes\",\"type\":\"string\"}],\"name\":\"updateAttendance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506113e4806100d76000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80638da5cb5b1461005c578063952e6ed71461007a578063b608d3a4146100aa578063cc59d295146100c6578063fa9f82cd146100e2575b600080fd5b6100646100fe565b60405161007191906108ed565b60405180910390f35b610094600480360381019061008f9190610952565b610122565b6040516100a19190610b69565b60405180910390f35b6100c460048036038101906100bf9190610cc0565b61042a565b005b6100e060048036038101906100db9190610cc0565b6105fe565b005b6100fc60048036038101906100f79190610da7565b61074d565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060600060028054905067ffffffffffffffff81111561014557610144610b95565b5b60405190808252806020026020018201604052801561017e57816020015b61016b610884565b8152602001906001900390816101635790505b5090506000805b60028054905081101561037257600060019050600088141580156101ce575087600283815481106101b9576101b8610de7565b5b90600052602060002090600402016001015414155b156101d857600090505b60008711801561020c575086600283815481106101f8576101f7610de7565b5b906000526020600020906004020160020154105b1561021657600090505b60008611801561024a5750856002838154811061023657610235610de7565b5b906000526020600020906004020160020154115b1561025457600090505b8015610364576002828154811061026e5761026d610de7565b5b90600052602060002090600402016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820180546102b590610e45565b80601f01602080910402602001604051908101604052809291908181526020018280546102e190610e45565b801561032e5780601f106103035761010080835404028352916020019161032e565b820191906000526020600020905b81548152906001019060200180831161031157829003601f168201915b50505050508152505084848151811061034a57610349610de7565b5b6020026020010181905250828061036090610ea5565b9350505b508080600101915050610185565b5060008167ffffffffffffffff81111561038f5761038e610b95565b5b6040519080825280602002602001820160405280156103c857816020015b6103b5610884565b8152602001906001900390816103ad5790505b50905060005b8281101561041c578381815181106103e9576103e8610de7565b5b602002602001015182828151811061040457610403610de7565b5b602002602001018190525080806001019150506103ce565b508093505050509392505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104af90610f70565b60405180910390fd5b60005b6002805490508110156105bc5784600282815481106104dd576104dc610de7565b5b906000526020600020906004020160000154036105af576040518060800160405280868152602001858152602001848152602001838152506002828154811061052957610528610de7565b5b9060005260206000209060040201600082015181600001556020820151816001015560408201518160020155606082015181600301908161056a919061113c565b5090505083857fa0d3071670661a556feb29023709b20ab2a62d5c56f63343d67c087d28be3df785856040516105a1929190611256565b60405180910390a3506105f8565b80806001019150506104bb565b506040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ef906112d2565b60405180910390fd5b50505050565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661068a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161068190611364565b60405180910390fd5b600260405180608001604052808681526020018581526020018481526020018381525090806001815401808255809150506001900390600052602060002090600402016000909190919091506000820151816000015560208201518160010155604082015181600201556060820151816003019081610709919061113c565b50505082847fbd00b0172a98c460bbec3c75c55ec0715c680024b0009f618e7599ffb7ef0182848460405161073f929190611256565b60405180910390a350505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107d290610f70565b60405180910390fd5b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167f2c59b8f31260880424e870332281eba93ddfef67e37d699cebe5a7ba89f7a8a6826040516108789190611393565b60405180910390a25050565b6040518060800160405280600081526020016000815260200160008152602001606081525090565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006108d7826108ac565b9050919050565b6108e7816108cc565b82525050565b600060208201905061090260008301846108de565b92915050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61092f8161091c565b811461093a57600080fd5b50565b60008135905061094c81610926565b92915050565b60008060006060848603121561096b5761096a610912565b5b60006109798682870161093d565b935050602061098a8682870161093d565b925050604061099b8682870161093d565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6109da8161091c565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610a1a5780820151818401526020810190506109ff565b60008484015250505050565b6000601f19601f8301169050919050565b6000610a42826109e0565b610a4c81856109eb565b9350610a5c8185602086016109fc565b610a6581610a26565b840191505092915050565b6000608083016000830151610a8860008601826109d1565b506020830151610a9b60208601826109d1565b506040830151610aae60408601826109d1565b5060608301518482036060860152610ac68282610a37565b9150508091505092915050565b6000610adf8383610a70565b905092915050565b6000602082019050919050565b6000610aff826109a5565b610b0981856109b0565b935083602082028501610b1b856109c1565b8060005b85811015610b575784840389528151610b388582610ad3565b9450610b4383610ae7565b925060208a01995050600181019050610b1f565b50829750879550505050505092915050565b60006020820190508181036000830152610b838184610af4565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610bcd82610a26565b810181811067ffffffffffffffff82111715610bec57610beb610b95565b5b80604052505050565b6000610bff610908565b9050610c0b8282610bc4565b919050565b600067ffffffffffffffff821115610c2b57610c2a610b95565b5b610c3482610a26565b9050602081019050919050565b82818337600083830152505050565b6000610c63610c5e84610c10565b610bf5565b905082815260208101848484011115610c7f57610c7e610b90565b5b610c8a848285610c41565b509392505050565b600082601f830112610ca757610ca6610b8b565b5b8135610cb7848260208601610c50565b91505092915050565b60008060008060808587031215610cda57610cd9610912565b5b6000610ce88782880161093d565b9450506020610cf98782880161093d565b9350506040610d0a8782880161093d565b925050606085013567ffffffffffffffff811115610d2b57610d2a610917565b5b610d3787828801610c92565b91505092959194509250565b610d4c816108cc565b8114610d5757600080fd5b50565b600081359050610d6981610d43565b92915050565b60008115159050919050565b610d8481610d6f565b8114610d8f57600080fd5b50565b600081359050610da181610d7b565b92915050565b60008060408385031215610dbe57610dbd610912565b5b6000610dcc85828601610d5a565b9250506020610ddd85828601610d92565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610e5d57607f821691505b602082108103610e7057610e6f610e16565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610eb08261091c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610ee257610ee1610e76565b5b600182019050919050565b600082825260208201905092915050565b7f6f6e6c79206f776e65722063616e20706572666f726d2074686973206163746960008201527f6f6e000000000000000000000000000000000000000000000000000000000000602082015250565b6000610f5a602283610eed565b9150610f6582610efe565b604082019050919050565b60006020820190508181036000830152610f8981610f4d565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610ff27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610fb5565b610ffc8683610fb5565b95508019841693508086168417925050509392505050565b6000819050919050565b600061103961103461102f8461091c565b611014565b61091c565b9050919050565b6000819050919050565b6110538361101e565b61106761105f82611040565b848454610fc2565b825550505050565b600090565b61107c61106f565b61108781848461104a565b505050565b5b818110156110ab576110a0600082611074565b60018101905061108d565b5050565b601f8211156110f0576110c181610f90565b6110ca84610fa5565b810160208510156110d9578190505b6110ed6110e585610fa5565b83018261108c565b50505b505050565b600082821c905092915050565b6000611113600019846008026110f5565b1980831691505092915050565b600061112c8383611102565b9150826002028217905092915050565b611145826109e0565b67ffffffffffffffff81111561115e5761115d610b95565b5b6111688254610e45565b6111738282856110af565b600060209050601f8311600181146111a65760008415611194578287015190505b61119e8582611120565b865550611206565b601f1984166111b486610f90565b60005b828110156111dc578489015182556001820191506020850194506020810190506111b7565b868310156111f957848901516111f5601f891682611102565b8355505b6001600288020188555050505b505050505050565b6112178161091c565b82525050565b6000611228826109e0565b6112328185610eed565b93506112428185602086016109fc565b61124b81610a26565b840191505092915050565b600060408201905061126b600083018561120e565b818103602083015261127d818461121d565b90509392505050565b7f7265636f7264206e6f7420666f756e6400000000000000000000000000000000600082015250565b60006112bc601083610eed565b91506112c782611286565b602082019050919050565b600060208201905081810360008301526112eb816112af565b9050919050565b7f6e6f7420617574686f72697a656420746f20706572666f726d2074686973206160008201527f6374696f6e000000000000000000000000000000000000000000000000000000602082015250565b600061134e602583610eed565b9150611359826112f2565b604082019050919050565b6000602082019050818103600083015261137d81611341565b9050919050565b61138d81610d6f565b82525050565b60006020820190506113a86000830184611384565b9291505056fea26469706673582212201d7af750aa5af8b4f79c693c6439810e4c13ca291febc41864d81048787ffcb064736f6c634300081a0033",
}

// AttendanceABI is the input ABI used to generate the binding from.
// Deprecated: Use AttendanceMetaData.ABI instead.
var AttendanceABI = AttendanceMetaData.ABI

// AttendanceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AttendanceMetaData.Bin instead.
var AttendanceBin = AttendanceMetaData.Bin

// DeployAttendance deploys a new Ethereum contract, binding an instance of Attendance to it.
func DeployAttendance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Attendance, error) {
	parsed, err := AttendanceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AttendanceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Attendance{AttendanceCaller: AttendanceCaller{contract: contract}, AttendanceTransactor: AttendanceTransactor{contract: contract}, AttendanceFilterer: AttendanceFilterer{contract: contract}}, nil
}

// Attendance is an auto generated Go binding around an Ethereum contract.
type Attendance struct {
	AttendanceCaller     // Read-only binding to the contract
	AttendanceTransactor // Write-only binding to the contract
	AttendanceFilterer   // Log filterer for contract events
}

// AttendanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttendanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttendanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttendanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttendanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttendanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttendanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttendanceSession struct {
	Contract     *Attendance       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AttendanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttendanceCallerSession struct {
	Contract *AttendanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AttendanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttendanceTransactorSession struct {
	Contract     *AttendanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AttendanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttendanceRaw struct {
	Contract *Attendance // Generic contract binding to access the raw methods on
}

// AttendanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttendanceCallerRaw struct {
	Contract *AttendanceCaller // Generic read-only contract binding to access the raw methods on
}

// AttendanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttendanceTransactorRaw struct {
	Contract *AttendanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttendance creates a new instance of Attendance, bound to a specific deployed contract.
func NewAttendance(address common.Address, backend bind.ContractBackend) (*Attendance, error) {
	contract, err := bindAttendance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Attendance{AttendanceCaller: AttendanceCaller{contract: contract}, AttendanceTransactor: AttendanceTransactor{contract: contract}, AttendanceFilterer: AttendanceFilterer{contract: contract}}, nil
}

// NewAttendanceCaller creates a new read-only instance of Attendance, bound to a specific deployed contract.
func NewAttendanceCaller(address common.Address, caller bind.ContractCaller) (*AttendanceCaller, error) {
	contract, err := bindAttendance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttendanceCaller{contract: contract}, nil
}

// NewAttendanceTransactor creates a new write-only instance of Attendance, bound to a specific deployed contract.
func NewAttendanceTransactor(address common.Address, transactor bind.ContractTransactor) (*AttendanceTransactor, error) {
	contract, err := bindAttendance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttendanceTransactor{contract: contract}, nil
}

// NewAttendanceFilterer creates a new log filterer instance of Attendance, bound to a specific deployed contract.
func NewAttendanceFilterer(address common.Address, filterer bind.ContractFilterer) (*AttendanceFilterer, error) {
	contract, err := bindAttendance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttendanceFilterer{contract: contract}, nil
}

// bindAttendance binds a generic wrapper to an already deployed contract.
func bindAttendance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AttendanceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attendance *AttendanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Attendance.Contract.AttendanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attendance *AttendanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attendance.Contract.AttendanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attendance *AttendanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attendance.Contract.AttendanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attendance *AttendanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Attendance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attendance *AttendanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attendance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attendance *AttendanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attendance.Contract.contract.Transact(opts, method, params...)
}

// GetAttendance is a free data retrieval call binding the contract method 0x952e6ed7.
//
// Solidity: function getAttendance(uint256 employerId, uint256 fromDate, uint256 toDate) view returns((uint256,uint256,uint256,string)[])
func (_Attendance *AttendanceCaller) GetAttendance(opts *bind.CallOpts, employerId *big.Int, fromDate *big.Int, toDate *big.Int) ([]AttendanceAttendanceRecord, error) {
	var out []interface{}
	err := _Attendance.contract.Call(opts, &out, "getAttendance", employerId, fromDate, toDate)

	if err != nil {
		return *new([]AttendanceAttendanceRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]AttendanceAttendanceRecord)).(*[]AttendanceAttendanceRecord)

	return out0, err

}

// GetAttendance is a free data retrieval call binding the contract method 0x952e6ed7.
//
// Solidity: function getAttendance(uint256 employerId, uint256 fromDate, uint256 toDate) view returns((uint256,uint256,uint256,string)[])
func (_Attendance *AttendanceSession) GetAttendance(employerId *big.Int, fromDate *big.Int, toDate *big.Int) ([]AttendanceAttendanceRecord, error) {
	return _Attendance.Contract.GetAttendance(&_Attendance.CallOpts, employerId, fromDate, toDate)
}

// GetAttendance is a free data retrieval call binding the contract method 0x952e6ed7.
//
// Solidity: function getAttendance(uint256 employerId, uint256 fromDate, uint256 toDate) view returns((uint256,uint256,uint256,string)[])
func (_Attendance *AttendanceCallerSession) GetAttendance(employerId *big.Int, fromDate *big.Int, toDate *big.Int) ([]AttendanceAttendanceRecord, error) {
	return _Attendance.Contract.GetAttendance(&_Attendance.CallOpts, employerId, fromDate, toDate)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Attendance *AttendanceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Attendance.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Attendance *AttendanceSession) Owner() (common.Address, error) {
	return _Attendance.Contract.Owner(&_Attendance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Attendance *AttendanceCallerSession) Owner() (common.Address, error) {
	return _Attendance.Contract.Owner(&_Attendance.CallOpts)
}

// AuthorizeAccount is a paid mutator transaction binding the contract method 0xfa9f82cd.
//
// Solidity: function authorizeAccount(address entity, bool isAuthorized) returns()
func (_Attendance *AttendanceTransactor) AuthorizeAccount(opts *bind.TransactOpts, entity common.Address, isAuthorized bool) (*types.Transaction, error) {
	return _Attendance.contract.Transact(opts, "authorizeAccount", entity, isAuthorized)
}

// AuthorizeAccount is a paid mutator transaction binding the contract method 0xfa9f82cd.
//
// Solidity: function authorizeAccount(address entity, bool isAuthorized) returns()
func (_Attendance *AttendanceSession) AuthorizeAccount(entity common.Address, isAuthorized bool) (*types.Transaction, error) {
	return _Attendance.Contract.AuthorizeAccount(&_Attendance.TransactOpts, entity, isAuthorized)
}

// AuthorizeAccount is a paid mutator transaction binding the contract method 0xfa9f82cd.
//
// Solidity: function authorizeAccount(address entity, bool isAuthorized) returns()
func (_Attendance *AttendanceTransactorSession) AuthorizeAccount(entity common.Address, isAuthorized bool) (*types.Transaction, error) {
	return _Attendance.Contract.AuthorizeAccount(&_Attendance.TransactOpts, entity, isAuthorized)
}

// RecordAttendance is a paid mutator transaction binding the contract method 0xcc59d295.
//
// Solidity: function recordAttendance(uint256 id, uint256 employeeId, uint256 checkInTime, string notes) returns()
func (_Attendance *AttendanceTransactor) RecordAttendance(opts *bind.TransactOpts, id *big.Int, employeeId *big.Int, checkInTime *big.Int, notes string) (*types.Transaction, error) {
	return _Attendance.contract.Transact(opts, "recordAttendance", id, employeeId, checkInTime, notes)
}

// RecordAttendance is a paid mutator transaction binding the contract method 0xcc59d295.
//
// Solidity: function recordAttendance(uint256 id, uint256 employeeId, uint256 checkInTime, string notes) returns()
func (_Attendance *AttendanceSession) RecordAttendance(id *big.Int, employeeId *big.Int, checkInTime *big.Int, notes string) (*types.Transaction, error) {
	return _Attendance.Contract.RecordAttendance(&_Attendance.TransactOpts, id, employeeId, checkInTime, notes)
}

// RecordAttendance is a paid mutator transaction binding the contract method 0xcc59d295.
//
// Solidity: function recordAttendance(uint256 id, uint256 employeeId, uint256 checkInTime, string notes) returns()
func (_Attendance *AttendanceTransactorSession) RecordAttendance(id *big.Int, employeeId *big.Int, checkInTime *big.Int, notes string) (*types.Transaction, error) {
	return _Attendance.Contract.RecordAttendance(&_Attendance.TransactOpts, id, employeeId, checkInTime, notes)
}

// UpdateAttendance is a paid mutator transaction binding the contract method 0xb608d3a4.
//
// Solidity: function updateAttendance(uint256 id, uint256 employerId, uint256 checkInTime, string notes) returns()
func (_Attendance *AttendanceTransactor) UpdateAttendance(opts *bind.TransactOpts, id *big.Int, employerId *big.Int, checkInTime *big.Int, notes string) (*types.Transaction, error) {
	return _Attendance.contract.Transact(opts, "updateAttendance", id, employerId, checkInTime, notes)
}

// UpdateAttendance is a paid mutator transaction binding the contract method 0xb608d3a4.
//
// Solidity: function updateAttendance(uint256 id, uint256 employerId, uint256 checkInTime, string notes) returns()
func (_Attendance *AttendanceSession) UpdateAttendance(id *big.Int, employerId *big.Int, checkInTime *big.Int, notes string) (*types.Transaction, error) {
	return _Attendance.Contract.UpdateAttendance(&_Attendance.TransactOpts, id, employerId, checkInTime, notes)
}

// UpdateAttendance is a paid mutator transaction binding the contract method 0xb608d3a4.
//
// Solidity: function updateAttendance(uint256 id, uint256 employerId, uint256 checkInTime, string notes) returns()
func (_Attendance *AttendanceTransactorSession) UpdateAttendance(id *big.Int, employerId *big.Int, checkInTime *big.Int, notes string) (*types.Transaction, error) {
	return _Attendance.Contract.UpdateAttendance(&_Attendance.TransactOpts, id, employerId, checkInTime, notes)
}

// AttendanceAttendanceRecordedIterator is returned from FilterAttendanceRecorded and is used to iterate over the raw logs and unpacked data for AttendanceRecorded events raised by the Attendance contract.
type AttendanceAttendanceRecordedIterator struct {
	Event *AttendanceAttendanceRecorded // Event containing the contract specifics and raw log

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
func (it *AttendanceAttendanceRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttendanceAttendanceRecorded)
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
		it.Event = new(AttendanceAttendanceRecorded)
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
func (it *AttendanceAttendanceRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttendanceAttendanceRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttendanceAttendanceRecorded represents a AttendanceRecorded event raised by the Attendance contract.
type AttendanceAttendanceRecorded struct {
	Id          *big.Int
	EmployeeId  *big.Int
	CheckInTime *big.Int
	Notes       string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAttendanceRecorded is a free log retrieval operation binding the contract event 0xbd00b0172a98c460bbec3c75c55ec0715c680024b0009f618e7599ffb7ef0182.
//
// Solidity: event AttendanceRecorded(uint256 indexed id, uint256 indexed employeeId, uint256 checkInTime, string notes)
func (_Attendance *AttendanceFilterer) FilterAttendanceRecorded(opts *bind.FilterOpts, id []*big.Int, employeeId []*big.Int) (*AttendanceAttendanceRecordedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var employeeIdRule []interface{}
	for _, employeeIdItem := range employeeId {
		employeeIdRule = append(employeeIdRule, employeeIdItem)
	}

	logs, sub, err := _Attendance.contract.FilterLogs(opts, "AttendanceRecorded", idRule, employeeIdRule)
	if err != nil {
		return nil, err
	}
	return &AttendanceAttendanceRecordedIterator{contract: _Attendance.contract, event: "AttendanceRecorded", logs: logs, sub: sub}, nil
}

// WatchAttendanceRecorded is a free log subscription operation binding the contract event 0xbd00b0172a98c460bbec3c75c55ec0715c680024b0009f618e7599ffb7ef0182.
//
// Solidity: event AttendanceRecorded(uint256 indexed id, uint256 indexed employeeId, uint256 checkInTime, string notes)
func (_Attendance *AttendanceFilterer) WatchAttendanceRecorded(opts *bind.WatchOpts, sink chan<- *AttendanceAttendanceRecorded, id []*big.Int, employeeId []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var employeeIdRule []interface{}
	for _, employeeIdItem := range employeeId {
		employeeIdRule = append(employeeIdRule, employeeIdItem)
	}

	logs, sub, err := _Attendance.contract.WatchLogs(opts, "AttendanceRecorded", idRule, employeeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttendanceAttendanceRecorded)
				if err := _Attendance.contract.UnpackLog(event, "AttendanceRecorded", log); err != nil {
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

// ParseAttendanceRecorded is a log parse operation binding the contract event 0xbd00b0172a98c460bbec3c75c55ec0715c680024b0009f618e7599ffb7ef0182.
//
// Solidity: event AttendanceRecorded(uint256 indexed id, uint256 indexed employeeId, uint256 checkInTime, string notes)
func (_Attendance *AttendanceFilterer) ParseAttendanceRecorded(log types.Log) (*AttendanceAttendanceRecorded, error) {
	event := new(AttendanceAttendanceRecorded)
	if err := _Attendance.contract.UnpackLog(event, "AttendanceRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttendanceAttendanceUpdatedIterator is returned from FilterAttendanceUpdated and is used to iterate over the raw logs and unpacked data for AttendanceUpdated events raised by the Attendance contract.
type AttendanceAttendanceUpdatedIterator struct {
	Event *AttendanceAttendanceUpdated // Event containing the contract specifics and raw log

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
func (it *AttendanceAttendanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttendanceAttendanceUpdated)
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
		it.Event = new(AttendanceAttendanceUpdated)
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
func (it *AttendanceAttendanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttendanceAttendanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttendanceAttendanceUpdated represents a AttendanceUpdated event raised by the Attendance contract.
type AttendanceAttendanceUpdated struct {
	Id          *big.Int
	EmployeeId  *big.Int
	CheckInTime *big.Int
	Notes       string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAttendanceUpdated is a free log retrieval operation binding the contract event 0xa0d3071670661a556feb29023709b20ab2a62d5c56f63343d67c087d28be3df7.
//
// Solidity: event AttendanceUpdated(uint256 indexed id, uint256 indexed employeeId, uint256 checkInTime, string notes)
func (_Attendance *AttendanceFilterer) FilterAttendanceUpdated(opts *bind.FilterOpts, id []*big.Int, employeeId []*big.Int) (*AttendanceAttendanceUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var employeeIdRule []interface{}
	for _, employeeIdItem := range employeeId {
		employeeIdRule = append(employeeIdRule, employeeIdItem)
	}

	logs, sub, err := _Attendance.contract.FilterLogs(opts, "AttendanceUpdated", idRule, employeeIdRule)
	if err != nil {
		return nil, err
	}
	return &AttendanceAttendanceUpdatedIterator{contract: _Attendance.contract, event: "AttendanceUpdated", logs: logs, sub: sub}, nil
}

// WatchAttendanceUpdated is a free log subscription operation binding the contract event 0xa0d3071670661a556feb29023709b20ab2a62d5c56f63343d67c087d28be3df7.
//
// Solidity: event AttendanceUpdated(uint256 indexed id, uint256 indexed employeeId, uint256 checkInTime, string notes)
func (_Attendance *AttendanceFilterer) WatchAttendanceUpdated(opts *bind.WatchOpts, sink chan<- *AttendanceAttendanceUpdated, id []*big.Int, employeeId []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var employeeIdRule []interface{}
	for _, employeeIdItem := range employeeId {
		employeeIdRule = append(employeeIdRule, employeeIdItem)
	}

	logs, sub, err := _Attendance.contract.WatchLogs(opts, "AttendanceUpdated", idRule, employeeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttendanceAttendanceUpdated)
				if err := _Attendance.contract.UnpackLog(event, "AttendanceUpdated", log); err != nil {
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

// ParseAttendanceUpdated is a log parse operation binding the contract event 0xa0d3071670661a556feb29023709b20ab2a62d5c56f63343d67c087d28be3df7.
//
// Solidity: event AttendanceUpdated(uint256 indexed id, uint256 indexed employeeId, uint256 checkInTime, string notes)
func (_Attendance *AttendanceFilterer) ParseAttendanceUpdated(log types.Log) (*AttendanceAttendanceUpdated, error) {
	event := new(AttendanceAttendanceUpdated)
	if err := _Attendance.contract.UnpackLog(event, "AttendanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttendanceAuthorizationUpdatedIterator is returned from FilterAuthorizationUpdated and is used to iterate over the raw logs and unpacked data for AuthorizationUpdated events raised by the Attendance contract.
type AttendanceAuthorizationUpdatedIterator struct {
	Event *AttendanceAuthorizationUpdated // Event containing the contract specifics and raw log

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
func (it *AttendanceAuthorizationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttendanceAuthorizationUpdated)
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
		it.Event = new(AttendanceAuthorizationUpdated)
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
func (it *AttendanceAuthorizationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttendanceAuthorizationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttendanceAuthorizationUpdated represents a AuthorizationUpdated event raised by the Attendance contract.
type AttendanceAuthorizationUpdated struct {
	Entity       common.Address
	IsAuthorized bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuthorizationUpdated is a free log retrieval operation binding the contract event 0x2c59b8f31260880424e870332281eba93ddfef67e37d699cebe5a7ba89f7a8a6.
//
// Solidity: event AuthorizationUpdated(address indexed entity, bool isAuthorized)
func (_Attendance *AttendanceFilterer) FilterAuthorizationUpdated(opts *bind.FilterOpts, entity []common.Address) (*AttendanceAuthorizationUpdatedIterator, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _Attendance.contract.FilterLogs(opts, "AuthorizationUpdated", entityRule)
	if err != nil {
		return nil, err
	}
	return &AttendanceAuthorizationUpdatedIterator{contract: _Attendance.contract, event: "AuthorizationUpdated", logs: logs, sub: sub}, nil
}

// WatchAuthorizationUpdated is a free log subscription operation binding the contract event 0x2c59b8f31260880424e870332281eba93ddfef67e37d699cebe5a7ba89f7a8a6.
//
// Solidity: event AuthorizationUpdated(address indexed entity, bool isAuthorized)
func (_Attendance *AttendanceFilterer) WatchAuthorizationUpdated(opts *bind.WatchOpts, sink chan<- *AttendanceAuthorizationUpdated, entity []common.Address) (event.Subscription, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}

	logs, sub, err := _Attendance.contract.WatchLogs(opts, "AuthorizationUpdated", entityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttendanceAuthorizationUpdated)
				if err := _Attendance.contract.UnpackLog(event, "AuthorizationUpdated", log); err != nil {
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

// ParseAuthorizationUpdated is a log parse operation binding the contract event 0x2c59b8f31260880424e870332281eba93ddfef67e37d699cebe5a7ba89f7a8a6.
//
// Solidity: event AuthorizationUpdated(address indexed entity, bool isAuthorized)
func (_Attendance *AttendanceFilterer) ParseAuthorizationUpdated(log types.Log) (*AttendanceAuthorizationUpdated, error) {
	event := new(AttendanceAuthorizationUpdated)
	if err := _Attendance.contract.UnpackLog(event, "AuthorizationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
