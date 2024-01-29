// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/event"
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

// WrappedEtherMetaData contains all meta data concerning the WrappedEther contract.
var WrappedEtherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101606040523480156200001257600080fd5b506040518060400160405280600d81526020016c2bb930b83832b21022ba3432b960991b81525080604051806040016040528060018152602001603160f81b8152506040518060400160405280600d81526020016c2bb930b83832b21022ba3432b960991b815250604051806040016040528060048152602001630ae8aa8960e31b8152508160039081620000a89190620002d4565b506004620000b78282620002d4565b505050620000d56005836200019060201b620009591790919060201c565b61012052620000f281600662000190602090811b6200095917901c565b61014052815160208084019190912060e052815190820120610100524660a0526200018060e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b60805250503060c0525062000415565b6000602083511015620001b057620001a883620001e0565b9050620001da565b82620001c7836200022c60201b6200098a1760201c565b90620001d49082620002d4565b5060ff90505b92915050565b600080829050601f8151111562000217578260405163305a27a960e01b81526004016200020e9190620003a0565b60405180910390fd5b80516200022482620003f0565b179392505050565b90565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200025a57607f821691505b6020821081036200027b57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620002cf57600081815260208120601f850160051c81016020861015620002aa5750805b601f850160051c820191505b81811015620002cb57828155600101620002b6565b5050505b505050565b81516001600160401b03811115620002f057620002f06200022f565b620003088162000301845462000245565b8462000281565b602080601f831160018114620003405760008415620003275750858301515b600019600386901b1c1916600185901b178555620002cb565b600085815260208120601f198616915b82811015620003715788860151825594840194600190910190840162000350565b5085821015620003905787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600060208083528351808285015260005b81811015620003cf57858101830151858201604001528201620003b1565b506000604082860101526040601f19601f8301168501019250505092915050565b805160208083015191908110156200027b5760001960209190910360031b1b16919050565b60805160a05160c05160e051610100516101205161014051611b6b62000470600039600061068e0152600061066301526000611138015260006111100152600061106b01526000611095015260006110bf0152611b6b6000f3fe6080604052600436106101125760003560e01c806370a08231116100a5578063a457c2d711610074578063d0e30db011610059578063d0e30db014610311578063d505accf14610319578063dd62ed3e1461033957600080fd5b8063a457c2d7146102d1578063a9059cbb146102f157600080fd5b806370a08231146102315780637ecebe001461027457806384b0196e1461029457806395d89b41146102bc57600080fd5b80632e1a7d4d116100e15780632e1a7d4d146101c0578063313ce567146101e05780633644e515146101fc578063395093511461021157600080fd5b806306fdde0314610126578063095ea7b31461015157806318160ddd1461018157806323b872dd146101a057600080fd5b366101215761011f61038c565b005b600080fd5b34801561013257600080fd5b5061013b6103ea565b6040516101489190611664565b60405180910390f35b34801561015d57600080fd5b5061017161016c3660046116a7565b61047c565b6040519015158152602001610148565b34801561018d57600080fd5b506002545b604051908152602001610148565b3480156101ac57600080fd5b506101716101bb3660046116d1565b610496565b3480156101cc57600080fd5b5061011f6101db36600461170d565b6104ba565b3480156101ec57600080fd5b5060405160128152602001610148565b34801561020857600080fd5b506101926105cf565b34801561021d57600080fd5b5061017161022c3660046116a7565b6105de565b34801561023d57600080fd5b5061019261024c366004611726565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b34801561028057600080fd5b5061019261028f366004611726565b61062a565b3480156102a057600080fd5b506102a9610655565b6040516101489796959493929190611741565b3480156102c857600080fd5b5061013b6106fa565b3480156102dd57600080fd5b506101716102ec3660046116a7565b610709565b3480156102fd57600080fd5b5061017161030c3660046116a7565b6107c0565b61011f61038c565b34801561032557600080fd5b5061011f610334366004611800565b6107ce565b34801561034557600080fd5b50610192610354366004611873565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b33610397813461098d565b8073ffffffffffffffffffffffffffffffffffffffff167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c346040516103df91815260200190565b60405180910390a250565b6060600380546103f9906118a6565b80601f0160208091040260200160405190810160405280929190818152602001828054610425906118a6565b80156104725780601f1061044757610100808354040283529160200191610472565b820191906000526020600020905b81548152906001019060200180831161045557829003601f168201915b5050505050905090565b60003361048a818585610a66565b60019150505b92915050565b6000336104a4858285610be6565b6104af858585610ca3565b506001949350505050565b336104c58183610ec4565b60008173ffffffffffffffffffffffffffffffffffffffff168360405160006040518083038185875af1925050503d806000811461051f576040519150601f19603f3d011682016040523d82523d6000602084013e610524565b606091505b505090508061057a5760405162461bcd60e51b815260206004820152601360248201527f776974686472617720455448206661696c65640000000000000000000000000060448201526064015b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff167f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65846040516105c291815260200190565b60405180910390a2505050565b60006105d9611051565b905090565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061048a90829086906106259087906118f3565b610a66565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260076020526040812054610490565b6000606080828080836106897f00000000000000000000000000000000000000000000000000000000000000006005611189565b6106b47f00000000000000000000000000000000000000000000000000000000000000006006611189565b604080516000808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b6060600480546103f9906118a6565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152812054909190838110156107b35760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610571565b6104af8286868403610a66565b60003361048a818585610ca3565b8342111561081e5760405162461bcd60e51b815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e650000006044820152606401610571565b60007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c988888861084d8c61122d565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e00160405160208183030381529060405280519060200120905060006108b582611262565b905060006108c5828787876112aa565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146109425760405162461bcd60e51b815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e617475726500006044820152606401610571565b61094d8a8a8a610a66565b50505050505050505050565b60006020835110156109755761096e836112d2565b9050610490565b8161098084826119aa565b5060ff9050610490565b90565b73ffffffffffffffffffffffffffffffffffffffff82166109f05760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610571565b8060026000828254610a0291906118f3565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff8316610aee5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610571565b73ffffffffffffffffffffffffffffffffffffffff8216610b775760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610571565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610c9d5781811015610c905760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610571565b610c9d8484848403610a66565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610d2c5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610571565b73ffffffffffffffffffffffffffffffffffffffff8216610db55760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610571565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610e515760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610571565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610c9d565b73ffffffffffffffffffffffffffffffffffffffff8216610f4d5760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610571565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090205481811015610fe95760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610571565b73ffffffffffffffffffffffffffffffffffffffff83166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610bd9565b505050565b60003073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156110b757507f000000000000000000000000000000000000000000000000000000000000000046145b156110e157507f000000000000000000000000000000000000000000000000000000000000000090565b6105d9604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b606060ff831461119c5761096e83611329565b8180546111a8906118a6565b80601f01602080910402602001604051908101604052809291908181526020018280546111d4906118a6565b80156112215780601f106111f657610100808354040283529160200191611221565b820191906000526020600020905b81548152906001019060200180831161120457829003601f168201915b50505050509050610490565b73ffffffffffffffffffffffffffffffffffffffff811660009081526007602052604090208054600181018255905b50919050565b600061049061126f611051565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b60008060006112bb87878787611368565b915091506112c881611457565b5095945050505050565b600080829050601f8151111561131657826040517f305a27a90000000000000000000000000000000000000000000000000000000081526004016105719190611664565b805161132182611ac4565b179392505050565b60606000611336836115bf565b604080516020808252818301909252919250600091906020820181803683375050509182525060208101929092525090565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561139f575060009050600361144e565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156113f3573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166114475760006001925092505061144e565b9150600090505b94509492505050565b600081600481111561146b5761146b611b06565b036114735750565b600181600481111561148757611487611b06565b036114d45760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610571565b60028160048111156114e8576114e8611b06565b036115355760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610571565b600381600481111561154957611549611b06565b036115bc5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610571565b50565b600060ff8216601f811115610490576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000815180845260005b818110156116265760208185018101518683018201520161160a565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6020815260006116776020830184611600565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146116a257600080fd5b919050565b600080604083850312156116ba57600080fd5b6116c38361167e565b946020939093013593505050565b6000806000606084860312156116e657600080fd5b6116ef8461167e565b92506116fd6020850161167e565b9150604084013590509250925092565b60006020828403121561171f57600080fd5b5035919050565b60006020828403121561173857600080fd5b6116778261167e565b7fff00000000000000000000000000000000000000000000000000000000000000881681526000602060e08184015261177d60e084018a611600565b838103604085015261178f818a611600565b6060850189905273ffffffffffffffffffffffffffffffffffffffff8816608086015260a0850187905284810360c0860152855180825283870192509083019060005b818110156117ee578351835292840192918401916001016117d2565b50909c9b505050505050505050505050565b600080600080600080600060e0888a03121561181b57600080fd5b6118248861167e565b96506118326020890161167e565b95506040880135945060608801359350608088013560ff8116811461185657600080fd5b9699959850939692959460a0840135945060c09093013592915050565b6000806040838503121561188657600080fd5b61188f8361167e565b915061189d6020840161167e565b90509250929050565b600181811c908216806118ba57607f821691505b60208210810361125c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b80820180821115610490577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f82111561104c57600081815260208120601f850160051c810160208610156119835750805b601f850160051c820191505b818110156119a25782815560010161198f565b505050505050565b815167ffffffffffffffff8111156119c4576119c461192d565b6119d8816119d284546118a6565b8461195c565b602080601f831160018114611a2b57600084156119f55750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556119a2565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015611a7857888601518255948401946001909101908401611a59565b5085821015611ab457878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b8051602080830151919081101561125c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60209190910360031b1b16919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfea2646970667358221220174e320af26fcf09b5b8cf899f0380bce059f15b43fed7d79f6eb2cf7832ff2b64736f6c63430008100033",
}

// WrappedEtherABI is the input ABI used to generate the binding from.
// Deprecated: Use WrappedEtherMetaData.ABI instead.
var WrappedEtherABI = WrappedEtherMetaData.ABI

// WrappedEtherBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WrappedEtherMetaData.Bin instead.
var WrappedEtherBin = WrappedEtherMetaData.Bin

// DeployWrappedEther deploys a new Ethereum contract, binding an instance of WrappedEther to it.
func DeployWrappedEther(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WrappedEther, error) {
	parsed, err := WrappedEtherMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WrappedEtherBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WrappedEther{WrappedEtherCaller: WrappedEtherCaller{contract: contract}, WrappedEtherTransactor: WrappedEtherTransactor{contract: contract}, WrappedEtherFilterer: WrappedEtherFilterer{contract: contract}}, nil
}

// WrappedEther is an auto generated Go binding around an Ethereum contract.
type WrappedEther struct {
	WrappedEtherCaller     // Read-only binding to the contract
	WrappedEtherTransactor // Write-only binding to the contract
	WrappedEtherFilterer   // Log filterer for contract events
}

// WrappedEtherCaller is an auto generated read-only Go binding around an Ethereum contract.
type WrappedEtherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedEtherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WrappedEtherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedEtherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WrappedEtherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedEtherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WrappedEtherSession struct {
	Contract     *WrappedEther     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WrappedEtherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WrappedEtherCallerSession struct {
	Contract *WrappedEtherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WrappedEtherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WrappedEtherTransactorSession struct {
	Contract     *WrappedEtherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WrappedEtherRaw is an auto generated low-level Go binding around an Ethereum contract.
type WrappedEtherRaw struct {
	Contract *WrappedEther // Generic contract binding to access the raw methods on
}

// WrappedEtherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WrappedEtherCallerRaw struct {
	Contract *WrappedEtherCaller // Generic read-only contract binding to access the raw methods on
}

// WrappedEtherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WrappedEtherTransactorRaw struct {
	Contract *WrappedEtherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrappedEther creates a new instance of WrappedEther, bound to a specific deployed contract.
func NewWrappedEther(address common.Address, backend bind.ContractBackend) (*WrappedEther, error) {
	contract, err := bindWrappedEther(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WrappedEther{WrappedEtherCaller: WrappedEtherCaller{contract: contract}, WrappedEtherTransactor: WrappedEtherTransactor{contract: contract}, WrappedEtherFilterer: WrappedEtherFilterer{contract: contract}}, nil
}

// NewWrappedEtherCaller creates a new read-only instance of WrappedEther, bound to a specific deployed contract.
func NewWrappedEtherCaller(address common.Address, caller bind.ContractCaller) (*WrappedEtherCaller, error) {
	contract, err := bindWrappedEther(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedEtherCaller{contract: contract}, nil
}

// NewWrappedEtherTransactor creates a new write-only instance of WrappedEther, bound to a specific deployed contract.
func NewWrappedEtherTransactor(address common.Address, transactor bind.ContractTransactor) (*WrappedEtherTransactor, error) {
	contract, err := bindWrappedEther(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedEtherTransactor{contract: contract}, nil
}

// NewWrappedEtherFilterer creates a new log filterer instance of WrappedEther, bound to a specific deployed contract.
func NewWrappedEtherFilterer(address common.Address, filterer bind.ContractFilterer) (*WrappedEtherFilterer, error) {
	contract, err := bindWrappedEther(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrappedEtherFilterer{contract: contract}, nil
}

// bindWrappedEther binds a generic wrapper to an already deployed contract.
func bindWrappedEther(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WrappedEtherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedEther *WrappedEtherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedEther.Contract.WrappedEtherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedEther *WrappedEtherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedEther.Contract.WrappedEtherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedEther *WrappedEtherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedEther.Contract.WrappedEtherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedEther *WrappedEtherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedEther.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedEther *WrappedEtherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedEther.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedEther *WrappedEtherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedEther.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WrappedEther *WrappedEtherCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WrappedEther.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WrappedEther *WrappedEtherSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _WrappedEther.Contract.DOMAINSEPARATOR(&_WrappedEther.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WrappedEther *WrappedEtherCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _WrappedEther.Contract.DOMAINSEPARATOR(&_WrappedEther.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedEther *WrappedEtherCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedEther.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedEther *WrappedEtherSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrappedEther.Contract.Allowance(&_WrappedEther.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedEther *WrappedEtherCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrappedEther.Contract.Allowance(&_WrappedEther.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedEther *WrappedEtherCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedEther.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedEther *WrappedEtherSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrappedEther.Contract.BalanceOf(&_WrappedEther.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedEther *WrappedEtherCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrappedEther.Contract.BalanceOf(&_WrappedEther.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedEther *WrappedEtherCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WrappedEther.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedEther *WrappedEtherSession) Decimals() (uint8, error) {
	return _WrappedEther.Contract.Decimals(&_WrappedEther.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedEther *WrappedEtherCallerSession) Decimals() (uint8, error) {
	return _WrappedEther.Contract.Decimals(&_WrappedEther.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_WrappedEther *WrappedEtherCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _WrappedEther.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_WrappedEther *WrappedEtherSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _WrappedEther.Contract.Eip712Domain(&_WrappedEther.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_WrappedEther *WrappedEtherCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _WrappedEther.Contract.Eip712Domain(&_WrappedEther.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedEther *WrappedEtherCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrappedEther.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedEther *WrappedEtherSession) Name() (string, error) {
	return _WrappedEther.Contract.Name(&_WrappedEther.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedEther *WrappedEtherCallerSession) Name() (string, error) {
	return _WrappedEther.Contract.Name(&_WrappedEther.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WrappedEther *WrappedEtherCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedEther.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WrappedEther *WrappedEtherSession) Nonces(owner common.Address) (*big.Int, error) {
	return _WrappedEther.Contract.Nonces(&_WrappedEther.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WrappedEther *WrappedEtherCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _WrappedEther.Contract.Nonces(&_WrappedEther.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedEther *WrappedEtherCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrappedEther.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedEther *WrappedEtherSession) Symbol() (string, error) {
	return _WrappedEther.Contract.Symbol(&_WrappedEther.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedEther *WrappedEtherCallerSession) Symbol() (string, error) {
	return _WrappedEther.Contract.Symbol(&_WrappedEther.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedEther *WrappedEtherCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WrappedEther.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedEther *WrappedEtherSession) TotalSupply() (*big.Int, error) {
	return _WrappedEther.Contract.TotalSupply(&_WrappedEther.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedEther *WrappedEtherCallerSession) TotalSupply() (*big.Int, error) {
	return _WrappedEther.Contract.TotalSupply(&_WrappedEther.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrappedEther *WrappedEtherTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedEther.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrappedEther *WrappedEtherSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.Approve(&_WrappedEther.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrappedEther *WrappedEtherTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.Approve(&_WrappedEther.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WrappedEther *WrappedEtherTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedEther.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WrappedEther *WrappedEtherSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.DecreaseAllowance(&_WrappedEther.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WrappedEther *WrappedEtherTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.DecreaseAllowance(&_WrappedEther.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WrappedEther *WrappedEtherTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedEther.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WrappedEther *WrappedEtherSession) Deposit() (*types.Transaction, error) {
	return _WrappedEther.Contract.Deposit(&_WrappedEther.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WrappedEther *WrappedEtherTransactorSession) Deposit() (*types.Transaction, error) {
	return _WrappedEther.Contract.Deposit(&_WrappedEther.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WrappedEther *WrappedEtherTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedEther.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WrappedEther *WrappedEtherSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.IncreaseAllowance(&_WrappedEther.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WrappedEther *WrappedEtherTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.IncreaseAllowance(&_WrappedEther.TransactOpts, spender, addedValue)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WrappedEther *WrappedEtherTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WrappedEther.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WrappedEther *WrappedEtherSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WrappedEther.Contract.Permit(&_WrappedEther.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WrappedEther *WrappedEtherTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WrappedEther.Contract.Permit(&_WrappedEther.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WrappedEther *WrappedEtherTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedEther.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WrappedEther *WrappedEtherSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.Transfer(&_WrappedEther.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WrappedEther *WrappedEtherTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.Transfer(&_WrappedEther.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WrappedEther *WrappedEtherTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedEther.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WrappedEther *WrappedEtherSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.TransferFrom(&_WrappedEther.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WrappedEther *WrappedEtherTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.TransferFrom(&_WrappedEther.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WrappedEther *WrappedEtherTransactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _WrappedEther.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WrappedEther *WrappedEtherSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.Withdraw(&_WrappedEther.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WrappedEther *WrappedEtherTransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WrappedEther.Contract.Withdraw(&_WrappedEther.TransactOpts, wad)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WrappedEther *WrappedEtherTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedEther.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WrappedEther *WrappedEtherSession) Receive() (*types.Transaction, error) {
	return _WrappedEther.Contract.Receive(&_WrappedEther.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WrappedEther *WrappedEtherTransactorSession) Receive() (*types.Transaction, error) {
	return _WrappedEther.Contract.Receive(&_WrappedEther.TransactOpts)
}

// WrappedEtherApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WrappedEther contract.
type WrappedEtherApprovalIterator struct {
	Event *WrappedEtherApproval // Event containing the contract specifics and raw log

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
func (it *WrappedEtherApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedEtherApproval)
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
		it.Event = new(WrappedEtherApproval)
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
func (it *WrappedEtherApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedEtherApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedEtherApproval represents a Approval event raised by the WrappedEther contract.
type WrappedEtherApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedEther *WrappedEtherFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WrappedEtherApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrappedEther.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WrappedEtherApprovalIterator{contract: _WrappedEther.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedEther *WrappedEtherFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WrappedEtherApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrappedEther.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedEtherApproval)
				if err := _WrappedEther.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_WrappedEther *WrappedEtherFilterer) ParseApproval(log types.Log) (*WrappedEtherApproval, error) {
	event := new(WrappedEtherApproval)
	if err := _WrappedEther.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedEtherDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WrappedEther contract.
type WrappedEtherDepositIterator struct {
	Event *WrappedEtherDeposit // Event containing the contract specifics and raw log

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
func (it *WrappedEtherDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedEtherDeposit)
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
		it.Event = new(WrappedEtherDeposit)
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
func (it *WrappedEtherDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedEtherDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedEtherDeposit represents a Deposit event raised by the WrappedEther contract.
type WrappedEtherDeposit struct {
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WrappedEther *WrappedEtherFilterer) FilterDeposit(opts *bind.FilterOpts, dst []common.Address) (*WrappedEtherDepositIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WrappedEther.contract.FilterLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return &WrappedEtherDepositIterator{contract: _WrappedEther.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WrappedEther *WrappedEtherFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WrappedEtherDeposit, dst []common.Address) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WrappedEther.contract.WatchLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedEtherDeposit)
				if err := _WrappedEther.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WrappedEther *WrappedEtherFilterer) ParseDeposit(log types.Log) (*WrappedEtherDeposit, error) {
	event := new(WrappedEtherDeposit)
	if err := _WrappedEther.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedEtherEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the WrappedEther contract.
type WrappedEtherEIP712DomainChangedIterator struct {
	Event *WrappedEtherEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *WrappedEtherEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedEtherEIP712DomainChanged)
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
		it.Event = new(WrappedEtherEIP712DomainChanged)
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
func (it *WrappedEtherEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedEtherEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedEtherEIP712DomainChanged represents a EIP712DomainChanged event raised by the WrappedEther contract.
type WrappedEtherEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_WrappedEther *WrappedEtherFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*WrappedEtherEIP712DomainChangedIterator, error) {

	logs, sub, err := _WrappedEther.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &WrappedEtherEIP712DomainChangedIterator{contract: _WrappedEther.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_WrappedEther *WrappedEtherFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *WrappedEtherEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _WrappedEther.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedEtherEIP712DomainChanged)
				if err := _WrappedEther.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_WrappedEther *WrappedEtherFilterer) ParseEIP712DomainChanged(log types.Log) (*WrappedEtherEIP712DomainChanged, error) {
	event := new(WrappedEtherEIP712DomainChanged)
	if err := _WrappedEther.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedEtherTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WrappedEther contract.
type WrappedEtherTransferIterator struct {
	Event *WrappedEtherTransfer // Event containing the contract specifics and raw log

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
func (it *WrappedEtherTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedEtherTransfer)
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
		it.Event = new(WrappedEtherTransfer)
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
func (it *WrappedEtherTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedEtherTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedEtherTransfer represents a Transfer event raised by the WrappedEther contract.
type WrappedEtherTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedEther *WrappedEtherFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WrappedEtherTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedEther.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WrappedEtherTransferIterator{contract: _WrappedEther.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedEther *WrappedEtherFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WrappedEtherTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedEther.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedEtherTransfer)
				if err := _WrappedEther.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_WrappedEther *WrappedEtherFilterer) ParseTransfer(log types.Log) (*WrappedEtherTransfer, error) {
	event := new(WrappedEtherTransfer)
	if err := _WrappedEther.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WrappedEtherWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the WrappedEther contract.
type WrappedEtherWithdrawalIterator struct {
	Event *WrappedEtherWithdrawal // Event containing the contract specifics and raw log

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
func (it *WrappedEtherWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedEtherWithdrawal)
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
		it.Event = new(WrappedEtherWithdrawal)
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
func (it *WrappedEtherWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedEtherWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedEtherWithdrawal represents a Withdrawal event raised by the WrappedEther contract.
type WrappedEtherWithdrawal struct {
	Src common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WrappedEther *WrappedEtherFilterer) FilterWithdrawal(opts *bind.FilterOpts, src []common.Address) (*WrappedEtherWithdrawalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _WrappedEther.contract.FilterLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return &WrappedEtherWithdrawalIterator{contract: _WrappedEther.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WrappedEther *WrappedEtherFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WrappedEtherWithdrawal, src []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _WrappedEther.contract.WatchLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedEtherWithdrawal)
				if err := _WrappedEther.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WrappedEther *WrappedEtherFilterer) ParseWithdrawal(log types.Log) (*WrappedEtherWithdrawal, error) {
	event := new(WrappedEtherWithdrawal)
	if err := _WrappedEther.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}