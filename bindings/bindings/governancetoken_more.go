// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const GovernanceTokenStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/governance/GovernanceToken.sol:GovernanceToken\",\"label\":\"_balances\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1001,\"contract\":\"contracts/governance/GovernanceToken.sol:GovernanceToken\",\"label\":\"_allowances\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1002,\"contract\":\"contracts/governance/GovernanceToken.sol:GovernanceToken\",\"label\":\"_totalSupply\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"contracts/governance/GovernanceToken.sol:GovernanceToken\",\"label\":\"_name\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_string_storage\"},{\"astId\":1004,\"contract\":\"contracts/governance/GovernanceToken.sol:GovernanceToken\",\"label\":\"_symbol\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_string_storage\"},{\"astId\":1005,\"contract\":\"contracts/governance/GovernanceToken.sol:GovernanceToken\",\"label\":\"_nameFallback\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_string_storage\"},{\"astId\":1006,\"contract\":\"contracts/governance/GovernanceToken.sol:GovernanceToken\",\"label\":\"_versionFallback\",\"offset\":0,\"slot\":\"6\",\"type\":\"t_string_storage\"},{\"astId\":1007,\"contract\":\"contracts/governance/GovernanceToken.sol:GovernanceToken\",\"label\":\"_nonces\",\"offset\":0,\"slot\":\"7\",\"type\":\"t_mapping(t_address,t_struct(Counter)1014_storage)\"},{\"astId\":1008,\"contract\":\"contracts/governance/GovernanceToken.sol:GovernanceToken\",\"label\":\"_PERMIT_TYPEHASH_DEPRECATED_SLOT\",\"offset\":0,\"slot\":\"8\",\"type\":\"t_bytes32\"},{\"astId\":1009,\"contract\":\"contracts/governance/GovernanceToken.sol:GovernanceToken\",\"label\":\"_delegates\",\"offset\":0,\"slot\":\"9\",\"type\":\"t_mapping(t_address,t_address)\"},{\"astId\":1010,\"contract\":\"contracts/governance/GovernanceToken.sol:GovernanceToken\",\"label\":\"_checkpoints\",\"offset\":0,\"slot\":\"10\",\"type\":\"t_mapping(t_address,t_array(t_struct(Checkpoint)1013_storage)dyn_storage)\"},{\"astId\":1011,\"contract\":\"contracts/governance/GovernanceToken.sol:GovernanceToken\",\"label\":\"_totalSupplyCheckpoints\",\"offset\":0,\"slot\":\"11\",\"type\":\"t_array(t_struct(Checkpoint)1013_storage)dyn_storage\"},{\"astId\":1012,\"contract\":\"contracts/governance/GovernanceToken.sol:GovernanceToken\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"12\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_struct(Checkpoint)1013_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct ERC20Votes.Checkpoint[]\",\"numberOfBytes\":\"32\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_address,t_address)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e address)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_address\"},\"t_mapping(t_address,t_array(t_struct(Checkpoint)1013_storage)dyn_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e struct ERC20Votes.Checkpoint[])\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_array(t_struct(Checkpoint)1013_storage)dyn_storage\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_struct(Counter)1014_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e struct Counters.Counter)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_struct(Counter)1014_storage\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_struct(Checkpoint)1013_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ERC20Votes.Checkpoint\",\"numberOfBytes\":\"32\"},\"t_struct(Counter)1014_storage\":{\"encoding\":\"inplace\",\"label\":\"struct Counters.Counter\",\"numberOfBytes\":\"32\"},\"t_uint224\":{\"encoding\":\"inplace\",\"label\":\"uint224\",\"numberOfBytes\":\"28\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"}}}"

var GovernanceTokenStorageLayout = new(solc.StorageLayout)

var GovernanceTokenDeployedBin = "0x608060405234801561001057600080fd5b50600436106101e55760003560e01c8063715018a61161010f5780639ab24eb0116100a2578063d505accf11610071578063d505accf14610479578063dd62ed3e1461048c578063f1127ed8146104d2578063f2fde38b1461052457600080fd5b80639ab24eb01461042d578063a457c2d714610440578063a9059cbb14610453578063c3cda5201461046657600080fd5b80638da5cb5b116100de5780638da5cb5b146103d55780638e539e8c146103f357806391ddadf41461040657806395d89b411461042557600080fd5b8063715018a61461038c57806379cc6790146103945780637ecebe00146103a757806384b0196e146103ba57600080fd5b80633a46b1a811610187578063587cde1e11610156578063587cde1e146102bd5780635c19a95c1461031b5780636fcfff451461032e57806370a082311461035657600080fd5b80633a46b1a81461027a57806340c10f191461028d57806342966c68146102a25780634bf5d7e9146102b557600080fd5b806323b872dd116101c357806323b872dd1461023d578063313ce567146102505780633644e5151461025f578063395093511461026757600080fd5b806306fdde03146101ea578063095ea7b31461020857806318160ddd1461022b575b600080fd5b6101f2610537565b6040516101ff91906129be565b60405180910390f35b61021b6102163660046129fa565b6105c9565b60405190151581526020016101ff565b6002545b6040519081526020016101ff565b61021b61024b366004612a24565b6105e3565b604051601281526020016101ff565b61022f610607565b61021b6102753660046129fa565b610616565b61022f6102883660046129fa565b610662565b6102a061029b3660046129fa565b610717565b005b6102a06102b0366004612a60565b61072d565b6101f261073a565b6102f66102cb366004612a79565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600960205260409020541690565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101ff565b6102a0610329366004612a79565b6107ec565b61034161033c366004612a79565b6107f6565b60405163ffffffff90911681526020016101ff565b61022f610364366004612a79565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6102a0610825565b6102a06103a23660046129fa565b610839565b61022f6103b5366004612a79565b61084e565b6103c2610879565b6040516101ff9796959493929190612a94565b600c5473ffffffffffffffffffffffffffffffffffffffff166102f6565b61022f610401366004612a60565b61091e565b61040e6109a3565b60405165ffffffffffff90911681526020016101ff565b6101f26109ae565b61022f61043b366004612a79565b6109bd565b61021b61044e3660046129fa565b610aa1565b61021b6104613660046129fa565b610b72565b6102a0610474366004612b64565b610b80565b6102a0610487366004612bbc565b610cf7565b61022f61049a366004612c26565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6104e56104e0366004612c59565b610eb6565b60408051825163ffffffff1681526020928301517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1692810192909252016101ff565b6102a0610532366004612a79565b610f5c565b60606003805461054690612c99565b80601f016020809104026020016040519081016040528092919081815260200182805461057290612c99565b80156105bf5780601f10610594576101008083540402835291602001916105bf565b820191906000526020600020905b8154815290600101906020018083116105a257829003601f168201915b5050505050905090565b6000336105d7818585611044565b60019150505b92915050565b6000336105f18582856111f7565b6105fc8585856112ce565b506001949350505050565b6000610611611543565b905090565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906105d7908290869061065d908790612d15565b611044565b600061066c6109a3565b65ffffffffffff1682106106e1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4552433230566f7465733a20667574757265206c6f6f6b75700000000000000060448201526064015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83166000908152600a60205260409020610710908361167b565b9392505050565b61071f6117ac565b610729828261182d565b5050565b6107373382611837565b50565b6060436107456109a3565b65ffffffffffff16146107b4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4552433230566f7465733a2062726f6b656e20636c6f636b206d6f646500000060448201526064016106d8565b5060408051808201909152601d81527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c74000000602082015290565b6107373382611841565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600a60205260408120546105dd906118df565b61082d6117ac565b6108376000611979565b565b6108448233836111f7565b6107298282611837565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600760205260408120546105dd565b6000606080828080836108ad7f000000000000000000000000000000000000000000000000000000000000000060056119f0565b6108d87f000000000000000000000000000000000000000000000000000000000000000060066119f0565b604080516000808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b60006109286109a3565b65ffffffffffff168210610998576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4552433230566f7465733a20667574757265206c6f6f6b75700000000000000060448201526064016106d8565b6105dd600b8361167b565b600061061143611a94565b60606004805461054690612c99565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600a60205260408120548015610a795773ffffffffffffffffffffffffffffffffffffffff83166000908152600a6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8301908110610a4257610a42612d57565b60009182526020909120015464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16610a7c565b60005b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169392505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610b65576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084016106d8565b6105fc8286868403611044565b6000336105d78185856112ce565b83421115610bea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4552433230566f7465733a207369676e6174757265206578706972656400000060448201526064016106d8565b604080517fe48329057bfd03d55e49b547132e39cffd9c1820ad7b9d4c5307691425d15adf602082015273ffffffffffffffffffffffffffffffffffffffff8816918101919091526060810186905260808101859052600090610c7190610c699060a00160405160208183030381529060405280519060200120611b2c565b858585611b74565b9050610c7c81611b9c565b8614610ce4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4552433230566f7465733a20696e76616c6964206e6f6e63650000000000000060448201526064016106d8565b610cee8188611841565b50505050505050565b83421115610d61576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e6500000060448201526064016106d8565b60007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610d908c611b9c565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090506000610df882611b2c565b90506000610e0882878787611b74565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610e9f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e6174757265000060448201526064016106d8565b610eaa8a8a8a611044565b50505050505050505050565b604080518082019091526000808252602082015273ffffffffffffffffffffffffffffffffffffffff83166000908152600a60205260409020805463ffffffff8416908110610f0757610f07612d57565b60009182526020918290206040805180820190915291015463ffffffff8116825264010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16918101919091529392505050565b610f646117ac565b73ffffffffffffffffffffffffffffffffffffffff8116611007576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016106d8565b61073781611979565b600060208351101561102c5761102583611bd1565b90506105dd565b816110378482612dd4565b5060ff90506105dd565b90565b73ffffffffffffffffffffffffffffffffffffffff83166110e6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f726573730000000000000000000000000000000000000000000000000000000060648201526084016106d8565b73ffffffffffffffffffffffffffffffffffffffff8216611189576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f737300000000000000000000000000000000000000000000000000000000000060648201526084016106d8565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146112c857818110156112bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016106d8565b6112c88484848403611044565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316611371576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f647265737300000000000000000000000000000000000000000000000000000060648201526084016106d8565b73ffffffffffffffffffffffffffffffffffffffff8216611414576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f657373000000000000000000000000000000000000000000000000000000000060648201526084016106d8565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156114ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e6365000000000000000000000000000000000000000000000000000060648201526084016106d8565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a36112c8848484611c2d565b60003073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156115a957507f000000000000000000000000000000000000000000000000000000000000000046145b156115d357507f000000000000000000000000000000000000000000000000000000000000000090565b610611604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b8154600090818160058111156116d557600061169684611c38565b6116a09085612eee565b600088815260209020909150869082015463ffffffff1611156116c5578091506116d3565b6116d0816001612d15565b92505b505b808210156117225760006116e98383611d20565b600088815260209020909150869082015463ffffffff16111561170e5780915061171c565b611719816001612d15565b92505b506116d5565b80156117815760008681526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff015464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16611784565b60005b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169695505050505050565b600c5473ffffffffffffffffffffffffffffffffffffffff163314610837576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016106d8565b6107298282611d3b565b6107298282611e01565b73ffffffffffffffffffffffffffffffffffffffff8281166000818152600960208181526040808420805485845282862054949093528787167fffffffffffffffffffffffff00000000000000000000000000000000000000008416811790915590519190951694919391928592917f3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f9190a46112c8828483611e19565b600063ffffffff821115611975576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203360448201527f322062697473000000000000000000000000000000000000000000000000000060648201526084016106d8565b5090565b600c805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b606060ff8314611a035761102583611fbe565b818054611a0f90612c99565b80601f0160208091040260200160405190810160405280929190818152602001828054611a3b90612c99565b8015611a885780601f10611a5d57610100808354040283529160200191611a88565b820191906000526020600020905b815481529060010190602001808311611a6b57829003601f168201915b505050505090506105dd565b600065ffffffffffff821115611975576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203460448201527f382062697473000000000000000000000000000000000000000000000000000060648201526084016106d8565b60006105dd611b39611543565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b6000806000611b8587878787611ffd565b91509150611b92816120ec565b5095945050505050565b73ffffffffffffffffffffffffffffffffffffffff811660009081526007602052604090208054600181018255905b50919050565b600080829050601f81511115611c1557826040517f305a27a90000000000000000000000000000000000000000000000000000000081526004016106d891906129be565b8051611c2082612f01565b179392505050565b505050565b611c2883838361229f565b600081600003611c4a57506000919050565b60006001611c57846122de565b901c6001901b90506001818481611c7057611c70612f43565b048201901c90506001818481611c8857611c88612f43565b048201901c90506001818481611ca057611ca0612f43565b048201901c90506001818481611cb857611cb8612f43565b048201901c90506001818481611cd057611cd0612f43565b048201901c90506001818481611ce857611ce8612f43565b048201901c90506001818481611d0057611d00612f43565b048201901c905061071081828581611d1a57611d1a612f43565b04612372565b6000611d2f6002848418612f72565b61071090848416612d15565b611d458282612388565b6002547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1015611df3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4552433230566f7465733a20746f74616c20737570706c79207269736b73206f60448201527f766572666c6f77696e6720766f7465730000000000000000000000000000000060648201526084016106d8565b6112c8600b6124838361248f565b611e0b8282612694565b6112c8600b61285f8361248f565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614158015611e555750600081115b15611c285773ffffffffffffffffffffffffffffffffffffffff831615611f0a5773ffffffffffffffffffffffffffffffffffffffff83166000908152600a602052604081208190611eaa9061285f8561248f565b915091508473ffffffffffffffffffffffffffffffffffffffff167fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a7248383604051611eff929190918252602082015260400190565b60405180910390a250505b73ffffffffffffffffffffffffffffffffffffffff821615611c285773ffffffffffffffffffffffffffffffffffffffff82166000908152600a602052604081208190611f5a906124838561248f565b915091508373ffffffffffffffffffffffffffffffffffffffff167fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a7248383604051611faf929190918252602082015260400190565b60405180910390a25050505050565b60606000611fcb8361286b565b604080516020808252818301909252919250600091906020820181803683375050509182525060208101929092525090565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561203457506000905060036120e3565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015612088573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166120dc576000600192509250506120e3565b9150600090505b94509492505050565b600081600481111561210057612100612fad565b036121085750565b600181600481111561211c5761211c612fad565b03612183576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016106d8565b600281600481111561219757612197612fad565b036121fe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016106d8565b600381600481111561221257612212612fad565b03610737576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016106d8565b73ffffffffffffffffffffffffffffffffffffffff838116600090815260096020526040808220548584168352912054611c2892918216911683611e19565b600080608083901c156122f357608092831c92015b604083901c1561230557604092831c92015b602083901c1561231757602092831c92015b601083901c1561232957601092831c92015b600883901c1561233b57600892831c92015b600483901c1561234d57600492831c92015b600283901c1561235f57600292831c92015b600183901c156105dd5760010192915050565b60008183106123815781610710565b5090919050565b73ffffffffffffffffffffffffffffffffffffffff8216612405576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016106d8565b80600260008282546124179190612d15565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a361072960008383611c2d565b60006107108284612d15565b8254600090819081811561250f5760008781526020902082017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0160408051808201909152905463ffffffff8116825264010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166020820152612524565b60408051808201909152600080825260208201525b905080602001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16935061255984868863ffffffff16565b9250600082118015612583575061256e6109a3565b65ffffffffffff16816000015163ffffffff16145b156125fb57612591836128ac565b60008881526020902083017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0180547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092166401000000000263ffffffff90921691909117905561268a565b86604051806040016040528061261f6126126109a3565b65ffffffffffff166118df565b63ffffffff168152602001612633866128ac565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90811690915282546001810184556000938452602093849020835194909301519091166401000000000263ffffffff909316929092179101555b5050935093915050565b73ffffffffffffffffffffffffffffffffffffffff8216612737576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f730000000000000000000000000000000000000000000000000000000000000060648201526084016106d8565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260208190526040902054818110156127ed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f636500000000000000000000000000000000000000000000000000000000000060648201526084016106d8565b73ffffffffffffffffffffffffffffffffffffffff83166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3611c2883600084611c2d565b60006107108284612eee565b600060ff8216601f8111156105dd576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff821115611975576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203260448201527f323420626974730000000000000000000000000000000000000000000000000060648201526084016106d8565b6000815180845260005b8181101561298057602081850181015186830182015201612964565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000610710602083018461295a565b803573ffffffffffffffffffffffffffffffffffffffff811681146129f557600080fd5b919050565b60008060408385031215612a0d57600080fd5b612a16836129d1565b946020939093013593505050565b600080600060608486031215612a3957600080fd5b612a42846129d1565b9250612a50602085016129d1565b9150604084013590509250925092565b600060208284031215612a7257600080fd5b5035919050565b600060208284031215612a8b57600080fd5b610710826129d1565b7fff00000000000000000000000000000000000000000000000000000000000000881681526000602060e081840152612ad060e084018a61295a565b8381036040850152612ae2818a61295a565b6060850189905273ffffffffffffffffffffffffffffffffffffffff8816608086015260a0850187905284810360c0860152855180825283870192509083019060005b81811015612b4157835183529284019291840191600101612b25565b50909c9b505050505050505050505050565b803560ff811681146129f557600080fd5b60008060008060008060c08789031215612b7d57600080fd5b612b86876129d1565b95506020870135945060408701359350612ba260608801612b53565b92506080870135915060a087013590509295509295509295565b600080600080600080600060e0888a031215612bd757600080fd5b612be0886129d1565b9650612bee602089016129d1565b95506040880135945060608801359350612c0a60808901612b53565b925060a0880135915060c0880135905092959891949750929550565b60008060408385031215612c3957600080fd5b612c42836129d1565b9150612c50602084016129d1565b90509250929050565b60008060408385031215612c6c57600080fd5b612c75836129d1565b9150602083013563ffffffff81168114612c8e57600080fd5b809150509250929050565b600181811c90821680612cad57607f821691505b602082108103611bcb577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156105dd576105dd612ce6565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b601f821115611c2857600081815260208120601f850160051c81016020861015612dad5750805b601f850160051c820191505b81811015612dcc57828155600101612db9565b505050505050565b815167ffffffffffffffff811115612dee57612dee612d28565b612e0281612dfc8454612c99565b84612d86565b602080601f831160018114612e555760008415612e1f5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555612dcc565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015612ea257888601518255948401946001909101908401612e83565b5085821015612ede57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b818103818111156105dd576105dd612ce6565b80516020808301519190811015611bcb577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60209190910360031b1b16919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082612fa8577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfea164736f6c6343000810000a"

func init() {
	if err := json.Unmarshal([]byte(GovernanceTokenStorageLayoutJSON), GovernanceTokenStorageLayout); err != nil {
		panic(err)
	}

	layouts["GovernanceToken"] = GovernanceTokenStorageLayout
	deployedBytecodes["GovernanceToken"] = GovernanceTokenDeployedBin
}
