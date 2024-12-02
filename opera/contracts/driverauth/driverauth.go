package driverauth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GetContractBin is NodeDriverAuth contract genesis implementation bin code
// Has to be compiled with flag bin-runtime
// Built from opera-sfc e86e5db95f98965f4489ad962565a3850126023a, solc 0.8.27, optimize-runs 200, cancun evmVersion
func GetContractBin() []byte {
	return hexutil.MustDecode("0x60806040526004361061013c575f3560e01c806379bead38116100b3578063c0c53b8b1161006d578063c0c53b8b1461036d578063d6a0c7af1461038c578063e08d7e66146103ab578063ebdf104c146103ca578063f2fde38b146103e9578063fd1b6ec114610408575f5ffd5b806379bead38146102885780638da5cb5b146102a7578063a4066fbe146102d3578063a8ab09ba146102f2578063ad3cb1cc14610311578063b9cc6b1c1461034e575f5ffd5b80634b64e492116101045780634b64e492146101dd5780634f1ef286146101fc57806352d1902d1461020f57806366e7ea0f14610236578063715018a61461025557806376fed43a14610269575f5ffd5b80630aeeca00146101405780631cef4fab146101615780631e702f8314610180578063242a6e3f1461019f578063267ab446146101be575b5f5ffd5b34801561014b575f5ffd5b5061015f61015a366004611163565b610427565b005b34801561016c575f5ffd5b5061015f61017b36600461118e565b61048b565b34801561018b575f5ffd5b5061015f61019a3660046111d1565b6104a5565b3480156101aa575f5ffd5b5061015f6101b9366004611235565b610534565b3480156101c9575f5ffd5b5061015f6101d8366004611163565b6105c4565b3480156101e8575f5ffd5b5061015f6101f736600461127c565b6105fd565b61015f61020a3660046112ab565b610628565b34801561021a575f5ffd5b50610223610647565b6040519081526020015b60405180910390f35b348015610241575f5ffd5b5061015f61025036600461136e565b610662565b348015610260575f5ffd5b5061015f61070f565b348015610274575f5ffd5b5061015f610283366004611398565b610722565b348015610293575f5ffd5b5061015f6102a236600461136e565b6107b8565b3480156102b2575f5ffd5b506102bb6107f9565b6040516001600160a01b03909116815260200161022d565b3480156102de575f5ffd5b5061015f6102ed3660046111d1565b610827565b3480156102fd575f5ffd5b5061015f61030c3660046113f6565b610889565b34801561031c575f5ffd5b50610341604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161022d9190611428565b348015610359575f5ffd5b5061015f61036836600461145d565b6108f3565b348015610378575f5ffd5b5061015f61038736600461149b565b61092d565b348015610397575f5ffd5b5061015f6103a63660046114e3565b610a72565b3480156103b6575f5ffd5b5061015f6103c536600461155a565b610ab4565b3480156103d5575f5ffd5b5061015f6103e436600461158c565b610b10565b3480156103f4575f5ffd5b5061015f61040336600461127c565b610baf565b348015610413575f5ffd5b5061015f6104223660046114e3565b610bee565b61042f610c21565b6001546040516205776560e91b8152600481018390526001600160a01b0390911690630aeeca00906024015b5f604051808303815f87803b158015610472575f5ffd5b505af1158015610484573d5f5f3e3d5ffd5b5050505050565b610493610c21565b61049f84848484610c53565b50505050565b6001546001600160a01b031633146104d057604051630607323760e11b815260040160405180910390fd5b5f54604051631e702f8360e01b815260048101849052602481018390526001600160a01b0390911690631e702f83906044015b5f604051808303815f87803b15801561051a575f5ffd5b505af115801561052c573d5f5f3e3d5ffd5b505050505050565b5f546001600160a01b0316331461055e5760405163d42fccad60e01b815260040160405180910390fd5b60015460405163242a6e3f60e01b81526001600160a01b039091169063242a6e3f906105929086908690869060040161167e565b5f604051808303815f87803b1580156105a9575f5ffd5b505af11580156105bb573d5f5f3e3d5ffd5b50505050505050565b6105cc610c21565b60015460405163133d5a2360e11b8152600481018390526001600160a01b039091169063267ab4469060240161045b565b610605610c21565b610625816106116107f9565b303f6001546001600160a01b03163f610c53565b50565b610630610d00565b61063982610da4565b6106438282610dac565b5050565b5f610650610e6d565b505f5160206117e35f395f51905f5290565b5f546001600160a01b0316331461068c5760405163d42fccad60e01b815260040160405180910390fd5b5f546001600160a01b038381169116146106b957604051630ea42ef960e31b815260040160405180910390fd5b6001546001600160a01b039081169063e30443bc9084906106de9085908316316116a0565b6040516001600160e01b031960e085901b1681526001600160a01b0390921660048301526024820152604401610503565b610717610c21565b6107205f610eb6565b565b6001546001600160a01b0316331461074d57604051630607323760e11b815260040160405180910390fd5b5f54604051633b7f6a1d60e11b81526001600160a01b03909116906376fed43a9061078490889088908890889088906004016116bf565b5f604051808303815f87803b15801561079b575f5ffd5b505af11580156107ad573d5f5f3e3d5ffd5b505050505050505050565b6107c0610c21565b600154604051630f37d5a760e31b81526001600160a01b03848116600483015260248201849052909116906379bead3890604401610503565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b5f546001600160a01b031633146108515760405163d42fccad60e01b815260040160405180910390fd5b60015460405163520337df60e11b815260048101849052602481018390526001600160a01b039091169063a4066fbe90604401610503565b6001546001600160a01b031633146108b457604051630607323760e11b815260040160405180910390fd5b5f5460405163545584dd60e11b81526001600160a01b03858116600483015260248201859052604482018490529091169063a8ab09ba90606401610592565b6108fb610c21565b600154604051632e731ac760e21b81526001600160a01b039091169063b9cc6b1c9061050390859085906004016116f8565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156109715750825b90505f826001600160401b0316600114801561098c5750303b155b90508115801561099a575080155b156109b85760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156109e257845460ff60401b1916600160401b1785555b6109eb86610f26565b6109f3610f37565b600180546001600160a01b03808a166001600160a01b0319928316179092555f8054928b16929091169190911790558315610a6857845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b610a7a610c21565b60015460405163d6a0c7af60e01b81526001600160a01b03848116600483015283811660248301529091169063d6a0c7af90604401610503565b6001546001600160a01b03163314610adf57604051630607323760e11b815260040160405180910390fd5b5f54604051637046bf3360e11b81526001600160a01b039091169063e08d7e66906105039085908590600401611743565b6001546001600160a01b03163314610b3b57604051630607323760e11b815260040160405180910390fd5b5f54604051633af7c41360e21b81526001600160a01b039091169063ebdf104c90610b78908b908b908b908b908b908b908b908b90600401611756565b5f604051808303815f87803b158015610b8f575f5ffd5b505af1158015610ba1573d5f5f3e3d5ffd5b505050505050505050505050565b610bb7610c21565b6001600160a01b038116610be557604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61062581610eb6565b610bf6610c21565b813b1580610c035750803b155b15610a7a57604051636f7c43f160e01b815260040160405180910390fd5b33610c2a6107f9565b6001600160a01b0316146107205760405163118cdaa760e01b8152336004820152602401610bdc565b610c5c84610eb6565b836001600160a01b031663614619546040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610c94575f5ffd5b505af1158015610ca6573d5f5f3e3d5ffd5b50505050610cb383610eb6565b81303f14610cd4576040516311387fef60e21b815260040160405180910390fd5b6001546001600160a01b03163f811461049f5760405163f0c300ef60e01b815260040160405180910390fd5b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610d8657507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610d7a5f5160206117e35f395f51905f52546001600160a01b031690565b6001600160a01b031614155b156107205760405163703e46dd60e11b815260040160405180910390fd5b610625610c21565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610e06575060408051601f3d908101601f19168201909252610e03918101906117b5565b60015b610e2e57604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610bdc565b5f5160206117e35f395f51905f528114610e5e57604051632a87526960e21b815260048101829052602401610bdc565b610e688383610f3f565b505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107205760405163703e46dd60e11b815260040160405180910390fd5b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b610f2e610f94565b61062581610fdd565b610720610f94565b610f4882610fe5565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115610f8c57610e688282611048565b6106436110bc565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661072057604051631afcd79f60e31b815260040160405180910390fd5b610bb7610f94565b806001600160a01b03163b5f0361101a57604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610bdc565b5f5160206117e35f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b03168460405161106491906117cc565b5f60405180830381855af49150503d805f811461109c576040519150601f19603f3d011682016040523d82523d5f602084013e6110a1565b606091505b50915091506110b18583836110db565b925050505b92915050565b34156107205760405163b398979f60e01b815260040160405180910390fd5b6060826110f0576110eb8261113a565b611133565b815115801561110757506001600160a01b0384163b155b1561113057604051639996b31560e01b81526001600160a01b0385166004820152602401610bdc565b50805b9392505050565b80511561114a5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f60208284031215611173575f5ffd5b5035919050565b6001600160a01b0381168114610625575f5ffd5b5f5f5f5f608085870312156111a1575f5ffd5b84356111ac8161117a565b935060208501356111bc8161117a565b93969395505050506040820135916060013590565b5f5f604083850312156111e2575f5ffd5b50508035926020909101359150565b5f5f83601f840112611201575f5ffd5b5081356001600160401b03811115611217575f5ffd5b60208301915083602082850101111561122e575f5ffd5b9250929050565b5f5f5f60408486031215611247575f5ffd5b8335925060208401356001600160401b03811115611263575f5ffd5b61126f868287016111f1565b9497909650939450505050565b5f6020828403121561128c575f5ffd5b81356111338161117a565b634e487b7160e01b5f52604160045260245ffd5b5f5f604083850312156112bc575f5ffd5b82356112c78161117a565b915060208301356001600160401b038111156112e1575f5ffd5b8301601f810185136112f1575f5ffd5b80356001600160401b0381111561130a5761130a611297565b604051601f8201601f19908116603f011681016001600160401b038111828210171561133857611338611297565b60405281815282820160200187101561134f575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f5f6040838503121561137f575f5ffd5b823561138a8161117a565b946020939093013593505050565b5f5f5f5f5f608086880312156113ac575f5ffd5b85356113b78161117a565b94506020860135935060408601356001600160401b038111156113d8575f5ffd5b6113e4888289016111f1565b96999598509660600135949350505050565b5f5f5f60608486031215611408575f5ffd5b83356114138161117a565b95602085013595506040909401359392505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f5f6020838503121561146e575f5ffd5b82356001600160401b03811115611483575f5ffd5b61148f858286016111f1565b90969095509350505050565b5f5f5f606084860312156114ad575f5ffd5b83356114b88161117a565b925060208401356114c88161117a565b915060408401356114d88161117a565b809150509250925092565b5f5f604083850312156114f4575f5ffd5b82356114ff8161117a565b9150602083013561150f8161117a565b809150509250929050565b5f5f83601f84011261152a575f5ffd5b5081356001600160401b03811115611540575f5ffd5b6020830191508360208260051b850101111561122e575f5ffd5b5f5f6020838503121561156b575f5ffd5b82356001600160401b03811115611580575f5ffd5b61148f8582860161151a565b5f5f5f5f5f5f5f5f6080898b0312156115a3575f5ffd5b88356001600160401b038111156115b8575f5ffd5b6115c48b828c0161151a565b90995097505060208901356001600160401b038111156115e2575f5ffd5b6115ee8b828c0161151a565b90975095505060408901356001600160401b0381111561160c575f5ffd5b6116188b828c0161151a565b90955093505060608901356001600160401b03811115611636575f5ffd5b6116428b828c0161151a565b999c989b5096995094979396929594505050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b838152604060208201525f611697604083018486611656565b95945050505050565b808201808211156110b657634e487b7160e01b5f52601160045260245ffd5b60018060a01b0386168152846020820152608060408201525f6116e6608083018587611656565b90508260608301529695505050505050565b602081525f61170b602083018486611656565b949350505050565b8183525f6001600160fb1b0383111561172a575f5ffd5b8260051b80836020870137939093016020019392505050565b602081525f61170b602083018486611713565b608081525f611769608083018a8c611713565b828103602084015261177c81898b611713565b90508281036040840152611791818789611713565b905082810360608401526117a6818587611713565b9b9a5050505050505050505050565b5f602082840312156117c5575f5ffd5b5051919050565b5f82518060208501845e5f92019182525091905056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca264697066735822122068794a549ea260f5e8f76d22f28e70332219328d651409bff011f20bc678ea9864736f6c634300081b0033")
}

// ContractAddress is the NodeDriverAuth contract address
var ContractAddress = common.HexToAddress("0xd100ae0000000000000000000000000000000000")
