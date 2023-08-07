// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package basic

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

// BasicMetaData contains all meta data concerning the Basic contract.
var BasicMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b506040518060400160405280600b81526020017f426f74746c6544726f70730000000000000000000000000000000000000000008152506040518060400160405280600481526020017f44524f500000000000000000000000000000000000000000000000000000000081525081600390816200008e9190620003f7565b508060049081620000a09190620003f7565b505050620000c3620000b7620000c960201b60201c565b620000d060201b60201c565b620004db565b5f33905090565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200020f57607f821691505b602082108103620002255762000224620001ca565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620002897fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200024c565b6200029586836200024c565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620002df620002d9620002d384620002ad565b620002b6565b620002ad565b9050919050565b5f819050919050565b620002fa83620002bf565b620003126200030982620002e6565b84845462000258565b825550505050565b5f90565b620003286200031a565b62000335818484620002ef565b505050565b5b818110156200035c57620003505f826200031e565b6001810190506200033b565b5050565b601f821115620003ab5762000375816200022b565b62000380846200023d565b8101602085101562000390578190505b620003a86200039f856200023d565b8301826200033a565b50505b505050565b5f82821c905092915050565b5f620003cd5f1984600802620003b0565b1980831691505092915050565b5f620003e78383620003bc565b9150826002028217905092915050565b620004028262000193565b67ffffffffffffffff8111156200041e576200041d6200019d565b5b6200042a8254620001f7565b6200043782828562000360565b5f60209050601f8311600181146200046d575f841562000458578287015190505b620004648582620003da565b865550620004d3565b601f1984166200047d866200022b565b5f5b82811015620004a6578489015182556001820191506020850194506020810190506200047f565b86831015620004c65784890151620004c2601f891682620003bc565b8355505b6001600288020188555050505b505050505050565b6117ab80620004e95f395ff3fe608060405234801561000f575f80fd5b50600436106100f3575f3560e01c806370a0823111610095578063a457c2d711610064578063a457c2d714610273578063a9059cbb146102a3578063dd62ed3e146102d3578063f2fde38b14610303576100f3565b806370a08231146101fd578063715018a61461022d5780638da5cb5b1461023757806395d89b4114610255576100f3565b806323b872dd116100d157806323b872dd14610163578063313ce5671461019357806339509351146101b157806340c10f19146101e1576100f3565b806306fdde03146100f7578063095ea7b31461011557806318160ddd14610145575b5f80fd5b6100ff61031f565b60405161010c9190610efe565b60405180910390f35b61012f600480360381019061012a9190610faf565b6103af565b60405161013c9190611007565b60405180910390f35b61014d6103d1565b60405161015a919061102f565b60405180910390f35b61017d60048036038101906101789190611048565b6103da565b60405161018a9190611007565b60405180910390f35b61019b610408565b6040516101a891906110b3565b60405180910390f35b6101cb60048036038101906101c69190610faf565b610410565b6040516101d89190611007565b60405180910390f35b6101fb60048036038101906101f69190610faf565b610446565b005b610217600480360381019061021291906110cc565b61046f565b604051610224919061102f565b60405180910390f35b6102356104b4565b005b61023f6104c7565b60405161024c9190611106565b60405180910390f35b61025d6104ef565b60405161026a9190610efe565b60405180910390f35b61028d60048036038101906102889190610faf565b61057f565b60405161029a9190611007565b60405180910390f35b6102bd60048036038101906102b89190610faf565b6105f4565b6040516102ca9190611007565b60405180910390f35b6102ed60048036038101906102e8919061111f565b610616565b6040516102fa919061102f565b60405180910390f35b61031d600480360381019061031891906110cc565b610698565b005b60606003805461032e9061118a565b80601f016020809104026020016040519081016040528092919081815260200182805461035a9061118a565b80156103a55780601f1061037c576101008083540402835291602001916103a5565b820191905f5260205f20905b81548152906001019060200180831161038857829003601f168201915b5050505050905090565b5f806103b961071a565b90506103c6818585610721565b600191505092915050565b5f600254905090565b5f806103e461071a565b90506103f18582856108e4565b6103fc85858561096f565b60019150509392505050565b5f6012905090565b5f8061041a61071a565b905061043b81858561042c8589610616565b61043691906111e7565b610721565b600191505092915050565b61044e610bdb565b61046b82670de0b6b3a764000083610466919061121a565b610c59565b5050565b5f805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6104bc610bdb565b6104c55f610da7565b565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600480546104fe9061118a565b80601f016020809104026020016040519081016040528092919081815260200182805461052a9061118a565b80156105755780601f1061054c57610100808354040283529160200191610575565b820191905f5260205f20905b81548152906001019060200180831161055857829003601f168201915b5050505050905090565b5f8061058961071a565b90505f6105968286610616565b9050838110156105db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d2906112cb565b60405180910390fd5b6105e88286868403610721565b60019250505092915050565b5f806105fe61071a565b905061060b81858561096f565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b6106a0610bdb565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361070e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070590611359565b60405180910390fd5b61071781610da7565b50565b5f33905090565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361078f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610786906113e7565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036107fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f490611475565b60405180910390fd5b8060015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516108d7919061102f565b60405180910390a3505050565b5f6108ef8484610616565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610969578181101561095b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610952906114dd565b60405180910390fd5b6109688484848403610721565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036109dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d49061156b565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a4b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a42906115f9565b60405180910390fd5b610a56838383610e6a565b5f805f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610ad9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ad090611687565b60405180910390fd5b8181035f808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550815f808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610bc2919061102f565b60405180910390a3610bd5848484610e6f565b50505050565b610be361071a565b73ffffffffffffffffffffffffffffffffffffffff16610c016104c7565b73ffffffffffffffffffffffffffffffffffffffff1614610c57576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c4e906116ef565b60405180910390fd5b565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610cc7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cbe90611757565b60405180910390fd5b610cd25f8383610e6a565b8060025f828254610ce391906111e7565b92505081905550805f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055508173ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610d90919061102f565b60405180910390a3610da35f8383610e6f565b5050565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b505050565b505050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610eab578082015181840152602081019050610e90565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610ed082610e74565b610eda8185610e7e565b9350610eea818560208601610e8e565b610ef381610eb6565b840191505092915050565b5f6020820190508181035f830152610f168184610ec6565b905092915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610f4b82610f22565b9050919050565b610f5b81610f41565b8114610f65575f80fd5b50565b5f81359050610f7681610f52565b92915050565b5f819050919050565b610f8e81610f7c565b8114610f98575f80fd5b50565b5f81359050610fa981610f85565b92915050565b5f8060408385031215610fc557610fc4610f1e565b5b5f610fd285828601610f68565b9250506020610fe385828601610f9b565b9150509250929050565b5f8115159050919050565b61100181610fed565b82525050565b5f60208201905061101a5f830184610ff8565b92915050565b61102981610f7c565b82525050565b5f6020820190506110425f830184611020565b92915050565b5f805f6060848603121561105f5761105e610f1e565b5b5f61106c86828701610f68565b935050602061107d86828701610f68565b925050604061108e86828701610f9b565b9150509250925092565b5f60ff82169050919050565b6110ad81611098565b82525050565b5f6020820190506110c65f8301846110a4565b92915050565b5f602082840312156110e1576110e0610f1e565b5b5f6110ee84828501610f68565b91505092915050565b61110081610f41565b82525050565b5f6020820190506111195f8301846110f7565b92915050565b5f806040838503121561113557611134610f1e565b5b5f61114285828601610f68565b925050602061115385828601610f68565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806111a157607f821691505b6020821081036111b4576111b361115d565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6111f182610f7c565b91506111fc83610f7c565b9250828201905080821115611214576112136111ba565b5b92915050565b5f61122482610f7c565b915061122f83610f7c565b925082820261123d81610f7c565b91508282048414831517611254576112536111ba565b5b5092915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f775f8201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b5f6112b5602583610e7e565b91506112c08261125b565b604082019050919050565b5f6020820190508181035f8301526112e2816112a9565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f611343602683610e7e565b915061134e826112e9565b604082019050919050565b5f6020820190508181035f83015261137081611337565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f206164645f8201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b5f6113d1602483610e7e565b91506113dc82611377565b604082019050919050565b5f6020820190508181035f8301526113fe816113c5565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f2061646472655f8201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b5f61145f602283610e7e565b915061146a82611405565b604082019050919050565b5f6020820190508181035f83015261148c81611453565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000005f82015250565b5f6114c7601d83610e7e565b91506114d282611493565b602082019050919050565b5f6020820190508181035f8301526114f4816114bb565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f2061645f8201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b5f611555602583610e7e565b9150611560826114fb565b604082019050919050565b5f6020820190508181035f83015261158281611549565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f20616464725f8201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b5f6115e3602383610e7e565b91506115ee82611589565b604082019050919050565b5f6020820190508181035f830152611610816115d7565b9050919050565b7f45524332303a207472616e7366657220616d6f756e74206578636565647320625f8201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b5f611671602683610e7e565b915061167c82611617565b604082019050919050565b5f6020820190508181035f83015261169e81611665565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f6116d9602083610e7e565b91506116e4826116a5565b602082019050919050565b5f6020820190508181035f830152611706816116cd565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f2061646472657373005f82015250565b5f611741601f83610e7e565b915061174c8261170d565b602082019050919050565b5f6020820190508181035f83015261176e81611735565b905091905056fea2646970667358221220cb837989e04be3235383b4e8b3e49973861a0afab83abb4d4f238f3cee57f02764736f6c63430008140033",
}

// BasicABI is the input ABI used to generate the binding from.
// Deprecated: Use BasicMetaData.ABI instead.
var BasicABI = BasicMetaData.ABI

// BasicBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BasicMetaData.Bin instead.
var BasicBin = BasicMetaData.Bin

// DeployBasic deploys a new Ethereum contract, binding an instance of Basic to it.
func DeployBasic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Basic, error) {
	parsed, err := BasicMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Basic{BasicCaller: BasicCaller{contract: contract}, BasicTransactor: BasicTransactor{contract: contract}, BasicFilterer: BasicFilterer{contract: contract}}, nil
}

// Basic is an auto generated Go binding around an Ethereum contract.
type Basic struct {
	BasicCaller     // Read-only binding to the contract
	BasicTransactor // Write-only binding to the contract
	BasicFilterer   // Log filterer for contract events
}

// BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicSession struct {
	Contract     *Basic            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicCallerSession struct {
	Contract *BasicCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTransactorSession struct {
	Contract     *BasicTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicRaw struct {
	Contract *Basic // Generic contract binding to access the raw methods on
}

// BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicCallerRaw struct {
	Contract *BasicCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTransactorRaw struct {
	Contract *BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasic creates a new instance of Basic, bound to a specific deployed contract.
func NewBasic(address common.Address, backend bind.ContractBackend) (*Basic, error) {
	contract, err := bindBasic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Basic{BasicCaller: BasicCaller{contract: contract}, BasicTransactor: BasicTransactor{contract: contract}, BasicFilterer: BasicFilterer{contract: contract}}, nil
}

// NewBasicCaller creates a new read-only instance of Basic, bound to a specific deployed contract.
func NewBasicCaller(address common.Address, caller bind.ContractCaller) (*BasicCaller, error) {
	contract, err := bindBasic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicCaller{contract: contract}, nil
}

// NewBasicTransactor creates a new write-only instance of Basic, bound to a specific deployed contract.
func NewBasicTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTransactor, error) {
	contract, err := bindBasic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTransactor{contract: contract}, nil
}

// NewBasicFilterer creates a new log filterer instance of Basic, bound to a specific deployed contract.
func NewBasicFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicFilterer, error) {
	contract, err := bindBasic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicFilterer{contract: contract}, nil
}

// bindBasic binds a generic wrapper to an already deployed contract.
func bindBasic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BasicMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Basic *BasicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Basic.Contract.BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Basic *BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Basic.Contract.BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Basic *BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Basic.Contract.BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Basic *BasicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Basic *BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Basic *BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Basic.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Basic *BasicCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Basic.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Basic *BasicSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Basic.Contract.Allowance(&_Basic.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Basic *BasicCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Basic.Contract.Allowance(&_Basic.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Basic *BasicCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Basic.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Basic *BasicSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Basic.Contract.BalanceOf(&_Basic.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Basic *BasicCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Basic.Contract.BalanceOf(&_Basic.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Basic *BasicCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Basic.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Basic *BasicSession) Decimals() (uint8, error) {
	return _Basic.Contract.Decimals(&_Basic.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Basic *BasicCallerSession) Decimals() (uint8, error) {
	return _Basic.Contract.Decimals(&_Basic.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Basic *BasicCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Basic.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Basic *BasicSession) Name() (string, error) {
	return _Basic.Contract.Name(&_Basic.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Basic *BasicCallerSession) Name() (string, error) {
	return _Basic.Contract.Name(&_Basic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Basic *BasicCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Basic.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Basic *BasicSession) Owner() (common.Address, error) {
	return _Basic.Contract.Owner(&_Basic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Basic *BasicCallerSession) Owner() (common.Address, error) {
	return _Basic.Contract.Owner(&_Basic.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Basic *BasicCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Basic.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Basic *BasicSession) Symbol() (string, error) {
	return _Basic.Contract.Symbol(&_Basic.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Basic *BasicCallerSession) Symbol() (string, error) {
	return _Basic.Contract.Symbol(&_Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Basic *BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Basic.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Basic *BasicSession) TotalSupply() (*big.Int, error) {
	return _Basic.Contract.TotalSupply(&_Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Basic *BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _Basic.Contract.TotalSupply(&_Basic.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Basic *BasicTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Basic.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Basic *BasicSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Basic.Contract.Approve(&_Basic.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Basic *BasicTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Basic.Contract.Approve(&_Basic.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Basic *BasicTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Basic.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Basic *BasicSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Basic.Contract.DecreaseAllowance(&_Basic.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Basic *BasicTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Basic.Contract.DecreaseAllowance(&_Basic.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Basic *BasicTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Basic.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Basic *BasicSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Basic.Contract.IncreaseAllowance(&_Basic.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Basic *BasicTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Basic.Contract.IncreaseAllowance(&_Basic.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address receiver, uint256 amount) returns()
func (_Basic *BasicTransactor) Mint(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Basic.contract.Transact(opts, "mint", receiver, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address receiver, uint256 amount) returns()
func (_Basic *BasicSession) Mint(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Basic.Contract.Mint(&_Basic.TransactOpts, receiver, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address receiver, uint256 amount) returns()
func (_Basic *BasicTransactorSession) Mint(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Basic.Contract.Mint(&_Basic.TransactOpts, receiver, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Basic *BasicTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Basic.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Basic *BasicSession) RenounceOwnership() (*types.Transaction, error) {
	return _Basic.Contract.RenounceOwnership(&_Basic.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Basic *BasicTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Basic.Contract.RenounceOwnership(&_Basic.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Basic *BasicTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Basic.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Basic *BasicSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Basic.Contract.Transfer(&_Basic.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Basic *BasicTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Basic.Contract.Transfer(&_Basic.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Basic *BasicTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Basic.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Basic *BasicSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Basic.Contract.TransferFrom(&_Basic.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Basic *BasicTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Basic.Contract.TransferFrom(&_Basic.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Basic *BasicTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Basic.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Basic *BasicSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Basic.Contract.TransferOwnership(&_Basic.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Basic *BasicTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Basic.Contract.TransferOwnership(&_Basic.TransactOpts, newOwner)
}

// BasicApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Basic contract.
type BasicApprovalIterator struct {
	Event *BasicApproval // Event containing the contract specifics and raw log

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
func (it *BasicApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicApproval)
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
		it.Event = new(BasicApproval)
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
func (it *BasicApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicApproval represents a Approval event raised by the Basic contract.
type BasicApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Basic *BasicFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BasicApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Basic.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BasicApprovalIterator{contract: _Basic.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Basic *BasicFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BasicApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Basic.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicApproval)
				if err := _Basic.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Basic *BasicFilterer) ParseApproval(log types.Log) (*BasicApproval, error) {
	event := new(BasicApproval)
	if err := _Basic.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasicOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Basic contract.
type BasicOwnershipTransferredIterator struct {
	Event *BasicOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BasicOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicOwnershipTransferred)
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
		it.Event = new(BasicOwnershipTransferred)
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
func (it *BasicOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicOwnershipTransferred represents a OwnershipTransferred event raised by the Basic contract.
type BasicOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Basic *BasicFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BasicOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Basic.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BasicOwnershipTransferredIterator{contract: _Basic.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Basic *BasicFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BasicOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Basic.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicOwnershipTransferred)
				if err := _Basic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Basic *BasicFilterer) ParseOwnershipTransferred(log types.Log) (*BasicOwnershipTransferred, error) {
	event := new(BasicOwnershipTransferred)
	if err := _Basic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Basic contract.
type BasicTransferIterator struct {
	Event *BasicTransfer // Event containing the contract specifics and raw log

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
func (it *BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTransfer)
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
		it.Event = new(BasicTransfer)
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
func (it *BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTransfer represents a Transfer event raised by the Basic contract.
type BasicTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Basic *BasicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BasicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Basic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BasicTransferIterator{contract: _Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Basic *BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Basic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTransfer)
				if err := _Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Basic *BasicFilterer) ParseTransfer(log types.Log) (*BasicTransfer, error) {
	event := new(BasicTransfer)
	if err := _Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
