// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const L1CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1002,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1003,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_1_0_1600\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1020_storage\"},{\"astId\":1004,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_51_0_20\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_52_0_1568\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1019_storage\"},{\"astId\":1006,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_101_0_1\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1007,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_102_0_1568\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)1019_storage\"},{\"astId\":1008,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_151_0_32\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_152_0_1568\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)1019_storage\"},{\"astId\":1010,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1011,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1012,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1013,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"receiveNonce\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_uint256\"},{\"astId\":1014,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_address\"},{\"astId\":1015,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_uint240\"},{\"astId\":1016,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1017,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"208\",\"type\":\"t_array(t_uint256)1018_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1018_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[42]\",\"numberOfBytes\":\"1344\"},\"t_array(t_uint256)1019_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1020_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L1CrossDomainMessengerDeployedBin = "0x60806040526004361061015f5760003560e01c80638129fc1c116100c0578063b1b1b20911610074578063d764ad0b11610059578063d764ad0b146103aa578063ecc70428146103bd578063f9a94d7f1461042257600080fd5b8063b1b1b2091461035a578063b28ade251461038a57600080fd5b80638cbeeef2116100a55780638cbeeef2146102575780639fce812c146102e6578063a4e7f8bd1461031a57600080fd5b80638129fc1c146102ba57806383a74074146102cf57600080fd5b80633f827a5a1161011757806354fd4d50116100fc57806354fd4d501461026d5780635644cfdf1461028f5780636e296e45146102a557600080fd5b80633f827a5a1461022f5780634c1d6a691461025757600080fd5b80630ff754ea116101485780630ff754ea146101ac5780632828d7e8146102055780633dbb202b1461021a57600080fd5b8063028f85f7146101645780630c56849814610197575b600080fd5b34801561017057600080fd5b50610179601081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b3480156101a357600080fd5b50610179603f81565b3480156101b857600080fd5b506101e07f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161018e565b34801561021157600080fd5b50610179604081565b61022d610228366004611b54565b610438565b005b34801561023b57600080fd5b50610244600181565b60405161ffff909116815260200161018e565b34801561026357600080fd5b50610179619c4081565b34801561027957600080fd5b5061028261069c565b60405161018e9190611c29565b34801561029b57600080fd5b5061017961138881565b3480156102b157600080fd5b506101e061073f565b3480156102c657600080fd5b5061022d61082b565b3480156102db57600080fd5b5061017962030d4081565b3480156102f257600080fd5b506101e07f000000000000000000000000000000000000000000000000000000000000000081565b34801561032657600080fd5b5061034a610335366004611c43565b60cf6020526000908152604090205460ff1681565b604051901515815260200161018e565b34801561036657600080fd5b5061034a610375366004611c43565b60cb6020526000908152604090205460ff1681565b34801561039657600080fd5b506101796103a5366004611c5c565b610a28565b61022d6103b8366004611cb0565b610a96565b3480156103c957600080fd5b5061041460ce547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b60405190815260200161018e565b34801561042e57600080fd5b5061041460cc5481565b6105717f0000000000000000000000000000000000000000000000000000000000000000610467858585610a28565b347fd764ad0b000000000000000000000000000000000000000000000000000000006104d360ce547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c6040516024016104ef9796959493929190611d7f565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611343565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a3385856105f660ce547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b86604051610608959493929190611dde565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060ce80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b60606106c77f000000000000000000000000000000000000000000000000000000000000000061150a565b6106f07f000000000000000000000000000000000000000000000000000000000000000061150a565b6107197f000000000000000000000000000000000000000000000000000000000000000061150a565b60405160200161072b93929190611e2c565b604051602081830303815290604052905090565b60cd5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff21530161080e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f7420736574000000000000000000000060648201526084015b60405180910390fd5b5060cd5473ffffffffffffffffffffffffffffffffffffffff1690565b6000547501000000000000000000000000000000000000000000900460ff1615808015610876575060005460017401000000000000000000000000000000000000000090910460ff16105b806108a85750303b1580156108a8575060005474010000000000000000000000000000000000000000900460ff166001145b610934576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610805565b600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000017905580156109ba57600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1675010000000000000000000000000000000000000000001790555b6109c26115c8565b8015610a2557600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6000611388619c4080603f610a44604063ffffffff8816611ed1565b610a4e9190611f01565b610a59601088611ed1565b610a669062030d40611f4f565b610a709190611f4f565b610a7a9190611f4f565b610a849190611f4f565b610a8e9190611f4f565b949350505050565b60f087901c60028110610b51576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206f722031206d657373616765732061726520737570706f7274656460648201527f20617420746869732074696d6500000000000000000000000000000000000000608482015260a401610805565b8061ffff16600003610c46576000610ba2878986868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508f92506116a1915050565b600081815260cb602052604090205490915060ff1615610c44576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a206c65676163792077697460448201527f6864726177616c20616c72656164792072656c617965640000000000000000006064820152608401610805565b505b6000610c8c898989898989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506116c092505050565b9050610c966116e3565b15610cd357853414610caa57610caa611f77565b600081815260cf602052604090205460ff1615610cc957610cc9611f77565b60cc899055610e25565b3415610d87576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a401610805565b600081815260cf602052604090205460ff16610e25576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c61796564000000000000000000000000000000006064820152608401610805565b610e2e87611807565b15610ee1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a401610805565b600081815260cb602052604090205460ff1615610f80576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c61796564000000000000000000006064820152608401610805565b610fa185610f92611388619c40611f4f565b67ffffffffffffffff1661187e565b1580610fc7575060cd5473ffffffffffffffffffffffffffffffffffffffff1661dead14155b156110e057600081815260cf602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555182917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff32016110d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d657373616765000000000000000000000000000000000000006064820152608401610805565b5050611319565b60cd80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a16179055600061117188619c405a6111349190611fa6565b8988888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061189c92505050565b60cd80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead1790559050801561120857600082815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a2611315565b600082815260cf602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff3201611315576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d657373616765000000000000000000000000000000000000006064820152608401610805565b5050505b50505050505050565b905090565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b7fecb8447569fbcb1b469a9cc521287ab0796bdd174c6ffd3b5e7c1d90a87f817e6113ae60ce547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663fd0ad31e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611419573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061143d9190611fb9565b6040805192835260208301919091520160405180910390a16040517fe9e05c4200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063e9e05c429084906114d2908890839089906000908990600401611fd2565b6000604051808303818588803b1580156114eb57600080fd5b505af11580156114ff573d6000803e3d6000fd5b505050505050505050565b60606000611517836118b6565b600101905060008167ffffffffffffffff8111156115375761153761202a565b6040519080825280601f01601f191660200182016040528015611561576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461156b57509392505050565b6000547501000000000000000000000000000000000000000000900460ff16611673576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610805565b60cd80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055565b60006116af85858585611998565b805190602001209050949350505050565b60006116d0878787878787611a31565b8051906020012090509695505050505050565b60003373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561132257507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639bf62d826040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117eb9190612059565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b600073ffffffffffffffffffffffffffffffffffffffff821630148061187857507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b92915050565b600080603f83619c4001026040850201603f5a021015949350505050565b600080600080845160208601878a8af19695505050505050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106118ff577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef8100000000831061192b576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061194957662386f26fc10000830492506010015b6305f5e1008310611961576305f5e100830492506008015b612710831061197557612710830492506004015b60648310611987576064830492506002015b600a83106118785760010192915050565b6060848484846040516024016119b19493929190612076565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcbd4ece9000000000000000000000000000000000000000000000000000000001790529050949350505050565b6060868686868686604051602401611a4e969594939291906120c0565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610a2557600080fd5b60008083601f840112611b0457600080fd5b50813567ffffffffffffffff811115611b1c57600080fd5b602083019150836020828501011115611b3457600080fd5b9250929050565b803563ffffffff81168114611b4f57600080fd5b919050565b60008060008060608587031215611b6a57600080fd5b8435611b7581611ad0565b9350602085013567ffffffffffffffff811115611b9157600080fd5b611b9d87828801611af2565b9094509250611bb0905060408601611b3b565b905092959194509250565b60005b83811015611bd6578181015183820152602001611bbe565b50506000910152565b60008151808452611bf7816020860160208601611bbb565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611c3c6020830184611bdf565b9392505050565b600060208284031215611c5557600080fd5b5035919050565b600080600060408486031215611c7157600080fd5b833567ffffffffffffffff811115611c8857600080fd5b611c9486828701611af2565b9094509250611ca7905060208501611b3b565b90509250925092565b600080600080600080600060c0888a031215611ccb57600080fd5b873596506020880135611cdd81611ad0565b95506040880135611ced81611ad0565b9450606088013593506080880135925060a088013567ffffffffffffffff811115611d1757600080fd5b611d238a828b01611af2565b989b979a50959850939692959293505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a0830152611dd160c083018486611d36565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff86168152608060208201526000611e0e608083018688611d36565b905083604083015263ffffffff831660608301529695505050505050565b60008451611e3e818460208901611bbb565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611e7a816001850160208a01611bbb565b60019201918201528351611e95816002840160208801611bbb565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681851681830481118215151615611ef857611ef8611ea2565b02949350505050565b600067ffffffffffffffff80841680611f43577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b67ffffffffffffffff818116838216019080821115611f7057611f70611ea2565b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b8181038181111561187857611878611ea2565b600060208284031215611fcb57600080fd5b5051919050565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015267ffffffffffffffff84166040820152821515606082015260a06080820152600061201f60a0830184611bdf565b979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006020828403121561206b57600080fd5b8151611c3c81611ad0565b600073ffffffffffffffffffffffffffffffffffffffff8087168352808616602084015250608060408301526120af6080830185611bdf565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a083015261210b60c0830184611bdf565b9897505050505050505056fea164736f6c6343000810000a"

func init() {
	if err := json.Unmarshal([]byte(L1CrossDomainMessengerStorageLayoutJSON), L1CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1CrossDomainMessenger"] = L1CrossDomainMessengerStorageLayout
	deployedBytecodes["L1CrossDomainMessenger"] = L1CrossDomainMessengerDeployedBin
}
