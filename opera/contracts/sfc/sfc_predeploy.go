package sfc

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GetContractBin is SFC contract genesis implementation bin code
// Has to be compiled with flag bin-runtime
// Built from opera-sfc e86e5db95f98965f4489ad962565a3850126023a, solc 0.8.27, optimize-runs 200, cancun evmVersion
func GetContractBin() []byte {
	return hexutil.MustDecode("0x6080604052600436106103b6575f3560e01c80638cddb015116101e9578063c7be95de11610108578063df00c9221161009d578063e9a505a71161006d578063e9a505a714610dc6578063ebdf104c14610de5578063f2fde38b14610e04578063fb36025f14610e23575f5ffd5b8063df00c92214610d14578063e08d7e6614610d4e578063e261641a14610d6d578063e880a15914610da7575f5ffd5b8063d46fa518116100d8578063d46fa51814610c70578063d725e91f14610c8d578063db5ca3e514610cac578063dc31e1af14610cda575f5ffd5b8063c7be95de14610be7578063cc17278414610bfc578063cc8343aa14610c1b578063cfd4766314610c3a575f5ffd5b8063aa5d82921161017e578063b88a37e21161014e578063b88a37e214610b31578063c3de580e14610b5d578063c5f956af14610b9d578063c65ee0e114610bbc575f5ffd5b8063aa5d8292146109e4578063ad3cb1cc14610a3f578063b0ef386c14610a6f578063b5d8962714610a8e575f5ffd5b8063a198d229116101b9578063a198d22914610942578063a5a470ad1461097c578063a86a056f1461098f578063a8ab09ba146109c5575f5ffd5b80638cddb015146108b55780638da5cb5b146108d457806390a6c475146109105780639fa6dd351461092f575f5ffd5b806352d1902d116102d5578063736de9ae1161026a578063841e45611161023a578063841e456114610836578063854873e114610855578063860c2750146108815780638b0e9f3f146108a0575f5ffd5b8063736de9ae146107ba57806376671808146107ee57806376fed43a146108025780637cacb1d614610821575f5ffd5b80636099ecb2116102a55780636099ecb21461070c57806361e53fcc1461072b5780636f49866314610765578063715018a6146107a6575f5ffd5b806352d1902d1461069c57806354fd4d50146106b05780635601fe01146106ce57806358f95b80146106ed575f5ffd5b806339b80c001161034b57806346f1ca351161031b57806346f1ca351461062c5780634f1ef2861461064b5780634f7c4efb1461065e5780634f864df41461067d575f5ffd5b806339b80c001461053c5780633fbfd4df146105ba578063441a3e70146105d9578063468f35ee146105f8575f5ffd5b806318160ddd1161038657806318160ddd1461048d5780631e702f83146104a25780631f270152146104c157806328f7314814610527575f5ffd5b80630135b1db146103d857806308c3687414610416578063093b41d0146104375780630962ef791461046e575f5ffd5b366103d45760405163ab064ad360e01b815260040160405180910390fd5b5f5ffd5b3480156103e3575f5ffd5b506104036103f2366004614299565b60036020525f908152604090205481565b6040519081526020015b60405180910390f35b348015610421575f5ffd5b506104356104303660046142b2565b610e4e565b005b348015610442575f5ffd5b50601154610456906001600160a01b031681565b6040516001600160a01b03909116815260200161040d565b348015610479575f5ffd5b506104356104883660046142b2565b610eb0565b348015610498575f5ffd5b50610403600c5481565b3480156104ad575f5ffd5b506104356104bc3660046142c9565b610f80565b3480156104cc575f5ffd5b5061050c6104db3660046142e9565b600a60209081525f938452604080852082529284528284209052825290208054600182015460029092015490919083565b6040805193845260208401929092529082015260600161040d565b348015610532575f5ffd5b5061040360075481565b348015610547575f5ffd5b5061058d6105563660046142b2565b600d60208190525f918252604090912060088101546009820154600a830154600b840154600c850154949095015492949193909286565b604080519687526020870195909552938501929092526060840152608083015260a082015260c00161040d565b3480156105c5575f5ffd5b506104356105d4366004614319565b611009565b3480156105e4575f5ffd5b506104356105f33660046142c9565b611172565b348015610603575f5ffd5b50610456610612366004614299565b60146020525f90815260409020546001600160a01b031681565b348015610637575f5ffd5b50610435610646366004614299565b61118a565b610435610659366004614380565b6111c2565b348015610669575f5ffd5b506104356106783660046142c9565b6111dd565b348015610688575f5ffd5b50610435610697366004614441565b61128c565b3480156106a7575f5ffd5b506104036113d9565b3480156106bb575f5ffd5b50604051600160fa1b815260200161040d565b3480156106d9575f5ffd5b506104036106e83660046142b2565b6113f4565b3480156106f8575f5ffd5b506104036107073660046142c9565b611426565b348015610717575f5ffd5b5061040361072636600461446a565b611446565b348015610736575f5ffd5b506104036107453660046142c9565b5f918252600d602090815260408084209284526001909201905290205490565b348015610770575f5ffd5b5061040361077f36600461446a565b6001600160a01b03919091165f908152600860209081526040808320938352929052205490565b3480156107b1575f5ffd5b5061043561148b565b3480156107c5575f5ffd5b506104566107d4366004614299565b60156020525f90815260409020546001600160a01b031681565b3480156107f9575f5ffd5b5061040361149e565b34801561080d575f5ffd5b5061043561081c3660046144d6565b6114b3565b34801561082c575f5ffd5b5061040360015481565b348015610841575f5ffd5b50610435610850366004614299565b611505565b348015610860575f5ffd5b5061087461086f3660046142b2565b61152f565b60405161040d9190614560565b34801561088c575f5ffd5b5061043561089b366004614299565b6115c6565b3480156108ab575f5ffd5b5061040360065481565b3480156108c0575f5ffd5b506104356108cf36600461446a565b6115f0565b3480156108df575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610456565b34801561091b575f5ffd5b5061043561092a3660046142b2565b611617565b61043561093d3660046142b2565b61162b565b34801561094d575f5ffd5b5061040361095c3660046142c9565b5f918252600d602090815260408084209284526006909201905290205490565b61043561098a366004614572565b611636565b34801561099a575f5ffd5b506104036109a936600461446a565b600960209081525f928352604080842090915290825290205481565b3480156109d0575f5ffd5b506104356109df3660046142e9565b61177d565b3480156109ef575f5ffd5b50610a276109fe3660046142c9565b5f918252600d60209081526040808420928452600390920190529020546001600160401b031690565b6040516001600160401b03909116815260200161040d565b348015610a4a575f5ffd5b50610874604051806040016040528060058152602001640352e302e360dc1b81525081565b348015610a7a575f5ffd5b50610435610a89366004614299565b6117bc565b348015610a99575f5ffd5b50610aee610aa83660046142b2565b600260208190525f918252604090912080546001820154928201546003830154600484015460058501546006909501549395946001600160a01b03909316939192909187565b6040805197885260208801969096526001600160a01b03909416948601949094526060850191909152608084015260a083019190915260c082015260e00161040d565b348015610b3c575f5ffd5b50610b50610b4b3660046142b2565b611814565b60405161040d91906145b0565b348015610b68575f5ffd5b50610b8d610b773660046142b2565b5f90815260026020526040902054608016151590565b604051901515815260200161040d565b348015610ba8575f5ffd5b50600f54610456906001600160a01b031681565b348015610bc7575f5ffd5b50610403610bd63660046142b2565b600e6020525f908152604090205481565b348015610bf2575f5ffd5b5061040360055481565b348015610c07575f5ffd5b50610435610c163660046145f2565b611876565b348015610c26575f5ffd5b50610435610c35366004614623565b61193c565b348015610c45575f5ffd5b50610403610c5436600461446a565b600b60209081525f928352604080842090915290825290205481565b348015610c7b575f5ffd5b506010546001600160a01b0316610456565b348015610c98575f5ffd5b50610435610ca7366004614299565b611a66565b348015610cb7575f5ffd5b50610403610cc63660046142b2565b5f908152600d602052604090206009015490565b348015610ce5575f5ffd5b50610403610cf43660046142c9565b5f918252600d602090815260408084209284526004909201905290205490565b348015610d1f575f5ffd5b50610403610d2e3660046142c9565b5f918252600d602090815260408084209284526002909201905290205490565b348015610d59575f5ffd5b50610435610d68366004614695565b611b0c565b348015610d78575f5ffd5b50610403610d873660046142c9565b5f918252600d602090815260408084209284526005909201905290205490565b348015610db2575f5ffd5b50610435610dc1366004614299565b611bcd565b348015610dd1575f5ffd5b50601354610456906001600160a01b031681565b348015610df0575f5ffd5b50610435610dff3660046146c7565b611bf7565b348015610e0f575f5ffd5b50610435610e1e366004614299565b611e9d565b348015610e2e575f5ffd5b50610403610e3d366004614299565b60126020525f908152604090205481565b335f610e5a8284611edc565b9050610e67828483611f5f565b82826001600160a01b03167f663e0f63f4fc6b01be195c4b56111fd6f14b947d6264497119b08daf77e26da583604051610ea391815260200190565b60405180910390a3505050565b335f610ebc8284611edc565b90505f610ec883611fec565b6001600160a01b0316826040515f6040518083038185875af1925050503d805f8114610f0f576040519150601f19603f3d011682016040523d82523d5f602084013e610f14565b606091505b5050905080610f36576040516312171d8360e31b815260040160405180910390fd5b83836001600160a01b03167f70de20a533702af05c8faf1637846c4586a021bbc71b6928b089b6d555e4fbc284604051610f7291815260200190565b60405180910390a350505050565b5f546001600160a01b03163314610faa5760405163c78d372960e01b815260040160405180910390fd5b80610fc85760405163396bd83560e21b815260040160405180910390fd5b610fd28282612014565b610fdc825f61193c565b5f828152600260208190526040822001546001600160a01b0316906110049082908190612116565b505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f8115801561104d5750825b90505f826001600160401b031660011480156110685750303b155b905081158015611076575080155b156110945760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156110be57845460ff60401b1916600160401b1785555b6110c7866121f2565b6110cf612203565b60018a90555f80546001600160a01b03808b166001600160a01b03199283161790925560108054928a1692909116919091179055600c89905561110f4290565b5f8b8152600d6020526040902060080155831561116657845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050565b61118633838361118133611fec565b61220b565b5050565b6040516001600160a01b0382169033907f857125196131cfcd709c738c6d1fd2701ce70f2a03785aeadae6f4b47fe73c1d905f90a350565b6111ca61259c565b6111d382612640565b6111868282612648565b6111e5612704565b5f82815260026020526040902054608016611213576040516321b6a8f960e11b815260040160405180910390fd5b670de0b6b3a764000081111561123c576040516357c70d6360e01b815260040160405180910390fd5b5f828152600e6020526040908190208290555182907f047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917906112809084815260200190565b60405180910390a25050565b33611297818561275f565b50815f036112b857604051631f2a200560e01b815260040160405180910390fd5b6001600160a01b0381165f908152600a602090815260408083208784528252808320868452909152902060020154156113045760405163756f5c2d60e11b815260040160405180910390fd5b61131481858460015f60016127cd565b6001600160a01b0381165f908152600a602090815260408083208784528252808320868452909152902060020182905561134c61149e565b6001600160a01b0382165f908152600a60209081526040808320888452825280832087845290915281209182554260019092019190915561138e90859061193c565b8284826001600160a01b03167fd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57856040516113cb91815260200190565b60405180910390a450505050565b5f6113e261299e565b505f516020614b825f395f51905f5290565b5f818152600260208181526040808420909201546001600160a01b03168352600b815281832093835292909252205490565b5f828152600d602090815260408083208484529091529020545b92915050565b5f5f61145284846129e7565b6001600160a01b0385165f9081526008602090815260408083208784529091529020549091506114839082906147a5565b949350505050565b611493612704565b61149c5f612a51565b565b5f60015460016114ae91906147a5565b905090565b5f546001600160a01b031633146114dd5760405163c78d372960e01b815260040160405180910390fd5b6114ee858585855f5f875f5f612ac1565b6005548411156114fe5760058490555b5050505050565b61150d612704565b600f80546001600160a01b0319166001600160a01b0392909216919091179055565b60046020525f908152604090208054611547906147b8565b80601f0160208091040260200160405190810160405280929190818152602001828054611573906147b8565b80156115be5780601f10611595576101008083540402835291602001916115be565b820191905f5260205f20905b8154815290600101906020018083116115a157829003601f168201915b505050505081565b6115ce612704565b601080546001600160a01b0319166001600160a01b0392909216919091179055565b6115fa828261275f565b6111865760405163208e0a4160e11b815260040160405180910390fd5b61161f612704565b61162881612c6c565b50565b611628338234611f5f565b60105f9054906101000a90046001600160a01b03166001600160a01b031663c5f530af6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611686573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116aa91906147f0565b3410156116ca5760405163047447a360e11b815260040160405180910390fd5b604281141580611703575081815f8181106116e7576116e7614807565b9050013560f81c60f81b6001600160f81b03191660c060f81b14155b15611721576040516338497f4960e11b815260040160405180910390fd5b60125f61172e8484612cd3565b6001600160a01b0316815260208101919091526040015f2054156117655760405163028aeb6760e21b815260040160405180910390fd5b611770338383612cff565b6111863360055434611f5f565b5f546001600160a01b031633146117a75760405163c78d372960e01b815260040160405180910390fd5b6117b38383835f612d2d565b61100481612e7b565b6117c4612704565b6013546001600160a01b038083169116036117f257604051639b92bed360e01b815260040160405180910390fd5b601380546001600160a01b0319166001600160a01b0392909216919091179055565b5f818152600d602090815260409182902060070180548351818402810184019094528084526060939283018282801561186a57602002820191905f5260205f20905b815481526020019060010190808311611856575b50505050509050919050565b6013546001600160a01b031633146118a15760405163ea8e4eb560e01b815260040160405180910390fd5b6001600160a01b038281165f908152601560205260409020548183169116036118dd5760405163eb81e1a360e01b815260040160405180910390fd5b806001600160a01b0316826001600160a01b03160361190f5760405163367558c360e01b815260040160405180910390fd5b6001600160a01b039182165f90815260146020526040902080546001600160a01b03191691909216179055565b5f8281526002602052604090206004015461196a57604051635926e0c360e01b815260040160405180910390fd5b5f828152600260205260409020600181015490541561198657505f5b5f5460405163520337df60e11b815260048101859052602481018390526001600160a01b039091169063a4066fbe906044015f604051808303815f87803b1580156119cf575f5ffd5b505af11580156119e1573d5f5f3e3d5ffd5b505050508180156119f157508015155b15611004575f805484825260046020819052604092839020925163242a6e3f60e01b81526001600160a01b039092169263242a6e3f92611a34928892910161481b565b5f604051808303815f87803b158015611a4b575f5ffd5b505af1158015611a5d573d5f5f3e3d5ffd5b50505050505050565b336001600160a01b038216611a8e5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038181165f90815260146020526040902054811690831614611aca57604051630fe3b3c160e31b815260040160405180910390fd5b6001600160a01b039081165f9081526015602090815260408083208054949095166001600160a01b031994851617909455601490529190912080549091169055565b5f546001600160a01b03163314611b365760405163c78d372960e01b815260040160405180910390fd5b5f600d5f611b4261149e565b81526020019081526020015f2090505f5f90505b82811015611bb8575f848483818110611b7157611b71614807565b602090810292909201355f81815260028452604080822060010154948890529020839055600c860154909350611ba9915082906147a5565b600c8501555050600101611b56565b50611bc7600782018484614221565b50505050565b611bd5612704565b601180546001600160a01b0319166001600160a01b0392909216919091179055565b5f546001600160a01b03163314611c215760405163c78d372960e01b815260040160405180910390fd5b5f600d5f611c2d61149e565b81526020019081526020015f2090505f81600701805480602002602001604051908101604052809291908181526020018280548015611c8957602002820191905f5260205f20905b815481526020019060010190808311611c75575b50505050509050611d0e82828c8c808060200260200160405190810160405280939291908181526020018383602002808284375f81840152601f19601f820116905080830192505050505050508b8b808060200260200160405190810160405280939291908181526020018383602002808284375f92019190915250612eed92505050565b600180545f908152600d602052604090206008810154909190421115611d40576008820154611d3d90426148a9565b90505b611dc0818584868c8c808060200260200160405190810160405280939291908181526020018383602002808284375f81840152601f19601f820116905080830192505050505050508b8b808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525061310792505050565b611dff818584868c8c808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525061382592505050565b5050611e0961149e565b6001554260088301554360098301556010546040805163d9a7c1f960e01b815290516001600160a01b039092169163d9a7c1f9916004808201926020929091908290030181865afa158015611e60573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e8491906147f0565b600b83015550600c54600d909101555050505050505050565b611ea5612704565b6001600160a01b038116611ed357604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61162881612a51565b5f611ee7838361275f565b506001600160a01b0383165f90815260086020908152604080832085845290915281205490819003611f2c5760405163899aaa9d60e01b815260040160405180910390fd5b6001600160a01b0384165f908152600860209081526040808320868452909152812055611f5881612e7b565b9392505050565b5f82815260026020526040902060040154611f8d57604051635926e0c360e01b815260040160405180910390fd5b5f8281526002602052604090205415611fb9576040516353670afb60e11b815260040160405180910390fd5b611fc68383836001612d2d565b611fcf82613aa3565b6110045760405163c2eb4ead60e01b815260040160405180910390fd5b6001600160a01b038082165f9081526015602052604081205490911680611440575090919050565b5f8281526002602052604090205415801561202e57508015155b15612055575f8281526002602052604090206001015460075461205191906148a9565b6007555b5f82815260026020526040902054811115611186575f8281526002602052604081208281556006015490036120e45761208c61149e565b5f8381526002602090815260409182902060068101849055426005909101819055825193845290830152805184927fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e4792908290030190a25b817fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e8260405161128091815260200190565b6011546001600160a01b031615611004576011546040516001600160a01b03858116602483015284811660448301525f921690627a12009060640160408051601f198184030181529181526020820180516001600160e01b0316631fbcb08360e11b1790525161218691906148bc565b5f604051808303815f8787f1925050503d805f81146121c0576040519150601f19603f3d011682016040523d82523d5f602084013e6121c5565b606091505b50509050801580156121d45750815b15611bc7576040516347b4be6960e11b815260040160405180910390fd5b6121fa613b5a565b61162881613ba3565b61149c613b5a565b6001600160a01b0384165f908152600a60209081526040808320868452825280832085845282528083208151606081018352815480825260018301549482019490945260029091015491810191909152910361227a57604051630fe3b3c160e31b815260040160405180910390fd5b60208082015182515f8781526002909352604090922060050154909190158015906122b457505f8681526002602052604090206005015482115b156122d45750505f84815260026020526040902060058101546006909101545b60105f9054906101000a90046001600160a01b03166001600160a01b031663b82b84276040518163ffffffff1660e01b8152600401602060405180830381865afa158015612324573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061234891906147f0565b61235290836147a5565b42101561237257604051635ada9a9960e01b815260040160405180910390fd5b60105f9054906101000a90046001600160a01b03166001600160a01b031663650acd666040518163ffffffff1660e01b8152600401602060405180830381865afa1580156123c2573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123e691906147f0565b6123f090826147a5565b6123f861149e565b1015612417576040516323ea994d60e01b815260040160405180910390fd5b6001600160a01b0387165f908152600a60209081526040808320898452825280832088845282528083206002908101548a855290835281842054600e9093529083205490926080909216151591906124729084908490613bab565b6001600160a01b038b165f908152600a602090815260408083208d845282528083208c84529091528120818155600181018290556002015590508083116124cc576040516318f967fb60e01b815260040160405180910390fd5b5f6001600160a01b0388166124e183866148a9565b6040515f81818185875af1925050503d805f811461251a576040519150601f19603f3d011682016040523d82523d5f602084013e61251f565b606091505b5050905080612541576040516312171d8360e31b815260040160405180910390fd5b61254a82612c6c565b888a8c6001600160a01b03167f75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a218760405161258791815260200190565b60405180910390a45050505050505050505050565b306001600160a01b037f0000000000000000000000002685751d3c7a49ebf485e823079ac65e2a35a3dd16148061262257507f0000000000000000000000002685751d3c7a49ebf485e823079ac65e2a35a3dd6001600160a01b03166126165f516020614b825f395f51905f52546001600160a01b031690565b6001600160a01b031614155b1561149c5760405163703e46dd60e11b815260040160405180910390fd5b611628612704565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156126a2575060408051601f3d908101601f1916820190925261269f918101906147f0565b60015b6126ca57604051634c9c8ce360e01b81526001600160a01b0383166004820152602401611eca565b5f516020614b825f395f51905f5281146126fa57604051632a87526960e21b815260048101829052602401611eca565b6110048383613c11565b336127367f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b03161461149c5760405163118cdaa760e01b8152336004820152602401611eca565b5f5f61276b84846129e7565b905061277683613c66565b6001600160a01b0385165f8181526009602090815260408083208884528252808320949094559181526008825282812086825290915290812080548392906127bf9084906147a5565b909155505015159392505050565b6001600160a01b0386165f908152600b60209081526040808320888452909152812080548692906127ff9084906148a9565b90915550505f858152600260205260409020600101546128209085906148a9565b5f8681526002602052604090206001015560065461283f9085906148a9565b6006555f85815260026020526040902054612866578360075461286291906148a9565b6007555b5f612870866113f4565b9050801580159061288c57505f86815260026020526040902054155b1561296c5760105f9054906101000a90046001600160a01b03166001600160a01b031663c5f530af6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156128e1573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061290591906147f0565b81101561293657821561292b5760405163047447a360e11b815260040160405180910390fd5b612936866001612014565b818015612949575061294786613aa3565b155b156129675760405163c2eb4ead60e01b815260040160405180910390fd5b612977565b612977866001612014565b5f8681526002602081905260409091200154611a5d9088906001600160a01b031686612116565b306001600160a01b037f0000000000000000000000002685751d3c7a49ebf485e823079ac65e2a35a3dd161461149c5760405163703e46dd60e11b815260040160405180910390fd5b6001600160a01b0382165f90815260096020908152604080832084845290915281205481612a1484613c66565b6001600160a01b0386165f908152600b60209081526040808320888452909152812054919250612a4682878686613cbb565b979650505050505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b6001600160a01b0389165f9081526003602052604090205415612af757604051633f4dc7d360e11b815260040160405180910390fd5b6001600160a01b0389165f8181526003602081815260408084208d90558c845260028083528185208b81559384018a905560048085018a90556005850188905560068501899055930180546001600160a01b0319169095179094555220612b5f878983614916565b508760125f612b6e8a8a612cd3565b6001600160a01b03166001600160a01b031681526020019081526020015f2081905550886001600160a01b0316887f49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf8686604051612bd6929190918252602082015260400190565b60405180910390a38115612c2057604080518381526020810183905289917fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e47910160405180910390a25b8415612c6157877fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e86604051612c5891815260200190565b60405180910390a25b505050505050505050565b8015611628576040515f9082156108fc0290839083818181858288f19350505050158015612c9c573d5f5f3e3d5ffd5b506040518181527f8918bd6046d08b314e457977f29562c5d76a7030d79b1edba66e8a5da0b77ae89060200160405180910390a150565b5f612ce182600281866149cf565b604051612cef9291906149f6565b6040519081900390209392505050565b5f60055f8154612d0e90614a05565b91829055509050611bc7848285855f612d2561149e565b425f5f612ac1565b815f03612d4d57604051631f2a200560e01b815260040160405180910390fd5b612d57848461275f565b506001600160a01b0384165f908152600b60209081526040808320868452909152902054612d869083906147a5565b6001600160a01b0385165f908152600b60209081526040808320878452825280832093909355600290522060010154612dbf83826147a5565b5f85815260026020526040902060010155600654612dde9084906147a5565b6006555f84815260026020526040902054612e055782600754612e0191906147a5565b6007555b612e1084821561193c565b83856001600160a01b03167f9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb85604051612e4c91815260200190565b60405180910390a35f84815260026020819052604090912001546114fe9086906001600160a01b031684612116565b5f546040516366e7ea0f60e01b8152306004820152602481018390526001600160a01b03909116906366e7ea0f906044015f604051808303815f87803b158015612ec3575f5ffd5b505af1158015612ed5573d5f5f3e3d5ffd5b5050505080600c54612ee791906147a5565b600c5550565b5f5b83518110156114fe5760105f9054906101000a90046001600160a01b03166001600160a01b0316635a68f01a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612f48573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f6c91906147f0565b828281518110612f7e57612f7e614807565b602002602001015111801561301e575060105f9054906101000a90046001600160a01b03166001600160a01b031662cc7f836040518163ffffffff1660e01b8152600401602060405180830381865afa158015612fdd573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061300191906147f0565b83828151811061301357613013614807565b602002602001015110155b1561306a5761304784828151811061303857613038614807565b60200260200101516008612014565b61306a84828151811061305c5761305c614807565b60200260200101515f61193c565b82818151811061307c5761307c614807565b6020026020010151856005015f86848151811061309b5761309b614807565b602002602001015181526020019081526020015f20819055508181815181106130c6576130c6614807565b6020026020010151856006015f8684815181106130e5576130e5614807565b60209081029190910181015182528101919091526040015f2055600101612eef565b5f6040518060a0016040528085516001600160401b0381111561312c5761312c61436c565b604051908082528060200260200182016040528015613155578160200160208202803683370190505b5081526020015f815260200185516001600160401b0381111561317a5761317a61436c565b6040519080825280602002602001820160405280156131a3578160200160208202803683370190505b5081526020015f81526020015f81525090505f5f90505b84518110156132e3575f866004015f8784815181106131db576131db614807565b602002602001015181526020019081526020015f205490505f5f90508185848151811061320a5761320a614807565b60200260200101511115613240578185848151811061322b5761322b614807565b602002602001015161323d91906148a9565b90505b8986848151811061325357613253614807565b6020026020010151826132669190614a1d565b6132709190614a48565b8460400151848151811061328657613286614807565b602002602001018181525050836040015183815181106132a8576132a8614807565b602002602001015184606001516132bf91906147a5565b606085015260808401516132d49082906147a5565b608085015250506001016131ba565b505f5b84518110156133d0578784828151811061330257613302614807565b60200260200101518986848151811061331d5761331d614807565b60200260200101518a5f015f8a878151811061333b5761333b614807565b602002602001015181526020019081526020015f205461335b9190614a1d565b6133659190614a48565b61336f9190614a1d565b6133799190614a48565b825180518390811061338d5761338d614807565b602090810291909101015281518051829081106133ac576133ac614807565b602002602001015182602001516133c391906147a5565b60208301526001016132e6565b505f5b84518110156136d4575f61347b8960105f9054906101000a90046001600160a01b03166001600160a01b031663d9a7c1f96040518163ffffffff1660e01b8152600401602060405180830381865afa158015613431573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061345591906147f0565b855180518690811061346957613469614807565b60200260200101518660200151613d24565b90506134ad83608001518460400151848151811061349b5761349b614807565b60200260200101518560600151613d5f565b6134b790826147a5565b90505f8683815181106134cc576134cc614807565b6020908102919091018101515f8181526002808452604080832090910154601054825163a778651560e01b815292519496506001600160a01b03918216959394613564948994929093169263a778651592600480820193918290030181865afa15801561353b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061355f91906147f0565b613eb0565b6001600160a01b0383165f908152600b6020908152604080832087845290915290205490915080156135c8576001600160a01b0383165f908152600860209081526040808320878452909152812080548492906135c29084906147a5565b90915550505b5f6135d383876148a9565b5f86815260026020526040812060010154919250811561360d5781613600670de0b6b3a764000085614a1d565b61360a9190614a48565b90505b5f87815260018f0160205260409020546136289082906147a5565b8f6001015f8981526020019081526020015f20819055508a898151811061365157613651614807565b60200260200101518f6004015f8981526020019081526020015f20819055508b898151811061368257613682614807565b60200260200101518e6002015f8981526020019081526020015f20546136a891906147a5565b8f6002015f8981526020019081526020015f2081905550505050505050505080806001019150506133d3565b506080810151600a8701819055600c54111561370a5785600a0154600c5f8282546136ff91906148a9565b9091555061370f9050565b5f600c555b600f546001600160a01b031615611a5d575f670de0b6b3a764000060105f9054906101000a90046001600160a01b03166001600160a01b03166394c3e9146040518163ffffffff1660e01b8152600401602060405180830381865afa15801561377a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061379e91906147f0565b83608001516137ad9190614a1d565b6137b79190614a48565b90506137c281612e7b565b600f546040515f916001600160a01b031690620f424090849084818181858888f193505050503d805f8114613812576040519150601f19603f3d011682016040523d82523d5f602084013e613817565b606091505b505050505050505050505050565b5f5b8251811015613a9b575f83828151811061384357613843614807565b602002602001015190505f8761385e670de0b6b3a764000090565b85858151811061387057613870614807565b60200260200101516138829190614a1d565b61388c9190614a48565b9050670de0b6b3a76400008111156138a95750670de0b6b3a76400005b5f8281526003870160209081526040808320815160608101835290546001600160401b038116825263ffffffff600160401b8204811694830194909452600160601b900490921690820152906138ff8383613ece565b5f85815260038b0160209081526040918290208351815485840151868601516001600160401b039093166bffffffffffffffffffffffff1990921691909117600160401b63ffffffff928316021763ffffffff60601b1916600160601b91909216021790556010548251631c25433760e01b815292519394506001600160a01b031692631c2543379260048082019392918290030181865afa1580156139a7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906139cb9190614a5b565b6001600160401b0316815f01516001600160401b0316108015613a71575060105f9054906101000a90046001600160a01b03166001600160a01b0316633fa225486040518163ffffffff1660e01b8152600401602060405180830381865afa158015613a39573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613a5d9190614a81565b63ffffffff16816040015163ffffffff1610155b15613a8b57613a81846010612014565b613a8b845f61193c565b5050600190920191506138279050565b505050505050565b5f670de0b6b3a764000060105f9054906101000a90046001600160a01b03166001600160a01b0316632265f2846040518163ffffffff1660e01b8152600401602060405180830381865afa158015613afd573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b2191906147f0565b613b2a846113f4565b613b349190614a1d565b613b3e9190614a48565b5f92835260026020526040909220600101549190911115919050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661149c57604051631afcd79f60e31b815260040160405180910390fd5b611ea5613b5a565b5f821580613bc15750670de0b6b3a76400008210155b15613bcd57505f611f58565b670de0b6b3a7640000613be083826148a9565b613bea9086614a1d565b613bf49190614a48565b613bff9060016147a5565b905083811115611f5857509192915050565b613c1a826140b1565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115613c5e576110048282614114565b61118661417d565b5f8181526002602052604081206006015415613cb3575f828152600260205260409020600601546001541015613c9e57505060015490565b505f9081526002602052604090206006015490565b505060015490565b5f818310613cca57505f611483565b5f838152600d6020818152604080842088855260019081018352818520548786529383528185208986520190915290912054670de0b6b3a764000087613d1084846148a9565b613d1a9190614a1d565b612a469190614a48565b5f825f03613d3357505f611483565b5f613d3e8587614a1d565b905082613d4b8583614a1d565b613d559190614a48565b9695505050505050565b5f825f03613d6e57505f611f58565b5f82613d7a8587614a1d565b613d849190614a48565b9050670de0b6b3a764000060105f9054906101000a90046001600160a01b03166001600160a01b03166394c3e9146040518163ffffffff1660e01b8152600401602060405180830381865afa158015613ddf573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613e0391906147f0565b60105f9054906101000a90046001600160a01b03166001600160a01b031663c74dd6216040518163ffffffff1660e01b8152600401602060405180830381865afa158015613e53573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613e7791906147f0565b613e8990670de0b6b3a76400006148a9565b613e9391906148a9565b613e9d9083614a1d565b613ea79190614a48565b95945050505050565b5f670de0b6b3a7640000613ec48385614a1d565b611f589190614a48565b604080516060810182525f8082526020820181905291810191909152604080516060810182525f8082526020820181905291810191909152826040015163ffffffff165f03613f31576001600160401b0384168152600160408201529050611440565b5f83604001516001613f439190614aa4565b63ffffffff1690505f846020015163ffffffff16866001600160401b0316865f01516001600160401b0316600185613f7b9190614ac0565b613f859190614adf565b613f8f9190614b08565b613f999190614b08565b9050613fa58282614b27565b6001600160401b03168352613fba8282614b54565b63ffffffff166020840152670de0b6b3a764000083516001600160401b03161115613feb57670de0b6b3a764000083525b60105f9054906101000a90046001600160a01b03166001600160a01b0316633fa225486040518163ffffffff1660e01b8152600401602060405180830381865afa15801561403b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061405f9190614a81565b63ffffffff16856040015163ffffffff161015614096576040850151614086906001614aa4565b63ffffffff1660408401526140a7565b60408086015163ffffffff16908401525b5090949350505050565b806001600160a01b03163b5f036140e657604051634c9c8ce360e01b81526001600160a01b0382166004820152602401611eca565b5f516020614b825f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b03168460405161413091906148bc565b5f60405180830381855af49150503d805f8114614168576040519150601f19603f3d011682016040523d82523d5f602084013e61416d565b606091505b5091509150613ea785838361419c565b341561149c5760405163b398979f60e01b815260040160405180910390fd5b6060826141b1576141ac826141f8565b611f58565b81511580156141c857506001600160a01b0384163b155b156141f157604051639996b31560e01b81526001600160a01b0385166004820152602401611eca565b5080611f58565b8051156142085780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b828054828255905f5260205f2090810192821561425a579160200282015b8281111561425a57823582559160200191906001019061423f565b5061426692915061426a565b5090565b5b80821115614266575f815560010161426b565b80356001600160a01b0381168114614294575f5ffd5b919050565b5f602082840312156142a9575f5ffd5b611f588261427e565b5f602082840312156142c2575f5ffd5b5035919050565b5f5f604083850312156142da575f5ffd5b50508035926020909101359150565b5f5f5f606084860312156142fb575f5ffd5b6143048461427e565b95602085013595506040909401359392505050565b5f5f5f5f5f60a0868803121561432d575f5ffd5b85359450602086013593506143446040870161427e565b92506143526060870161427e565b91506143606080870161427e565b90509295509295909350565b634e487b7160e01b5f52604160045260245ffd5b5f5f60408385031215614391575f5ffd5b61439a8361427e565b915060208301356001600160401b038111156143b4575f5ffd5b8301601f810185136143c4575f5ffd5b80356001600160401b038111156143dd576143dd61436c565b604051601f8201601f19908116603f011681016001600160401b038111828210171561440b5761440b61436c565b604052818152828201602001871015614422575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f5f5f60608486031215614453575f5ffd5b505081359360208301359350604090920135919050565b5f5f6040838503121561447b575f5ffd5b6144848361427e565b946020939093013593505050565b5f5f83601f8401126144a2575f5ffd5b5081356001600160401b038111156144b8575f5ffd5b6020830191508360208285010111156144cf575f5ffd5b9250929050565b5f5f5f5f5f608086880312156144ea575f5ffd5b6144f38661427e565b94506020860135935060408601356001600160401b03811115614514575f5ffd5b61452088828901614492565b96999598509660600135949350505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f611f586020830184614532565b5f5f60208385031215614583575f5ffd5b82356001600160401b03811115614598575f5ffd5b6145a485828601614492565b90969095509350505050565b602080825282518282018190525f918401906040840190835b818110156145e75783518352602093840193909201916001016145c9565b509095945050505050565b5f5f60408385031215614603575f5ffd5b61460c8361427e565b915061461a6020840161427e565b90509250929050565b5f5f60408385031215614634575f5ffd5b823591506020830135801515811461464a575f5ffd5b809150509250929050565b5f5f83601f840112614665575f5ffd5b5081356001600160401b0381111561467b575f5ffd5b6020830191508360208260051b85010111156144cf575f5ffd5b5f5f602083850312156146a6575f5ffd5b82356001600160401b038111156146bb575f5ffd5b6145a485828601614655565b5f5f5f5f5f5f5f5f6080898b0312156146de575f5ffd5b88356001600160401b038111156146f3575f5ffd5b6146ff8b828c01614655565b90995097505060208901356001600160401b0381111561471d575f5ffd5b6147298b828c01614655565b90975095505060408901356001600160401b03811115614747575f5ffd5b6147538b828c01614655565b90955093505060608901356001600160401b03811115614771575f5ffd5b61477d8b828c01614655565b999c989b5096995094979396929594505050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561144057611440614791565b600181811c908216806147cc57607f821691505b6020821081036147ea57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f60208284031215614800575f5ffd5b5051919050565b634e487b7160e01b5f52603260045260245ffd5b828152604060208201525f5f8354614832816147b8565b806040860152600182165f8114614850576001811461486c5761489d565b60ff1983166060870152606082151560051b870101935061489d565b865f5260205f205f5b8381101561489457815488820160600152600190910190602001614875565b87016060019450505b50919695505050505050565b8181038181111561144057611440614791565b5f82518060208501845e5f920191825250919050565b601f82111561100457805f5260205f20601f840160051c810160208510156148f75750805b601f840160051c820191505b818110156114fe575f8155600101614903565b6001600160401b0383111561492d5761492d61436c565b6149418361493b83546147b8565b836148d2565b5f601f841160018114614972575f851561495b5750838201355b5f19600387901b1c1916600186901b1783556114fe565b5f83815260208120601f198716915b828110156149a15786850135825560209485019460019092019101614981565b50868210156149bd575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b5f5f858511156149dd575f5ffd5b838611156149e9575f5ffd5b5050820193919092039150565b818382375f9101908152919050565b5f60018201614a1657614a16614791565b5060010190565b808202811582820484141761144057611440614791565b634e487b7160e01b5f52601260045260245ffd5b5f82614a5657614a56614a34565b500490565b5f60208284031215614a6b575f5ffd5b81516001600160401b0381168114611f58575f5ffd5b5f60208284031215614a91575f5ffd5b815163ffffffff81168114611f58575f5ffd5b63ffffffff818116838216019081111561144057611440614791565b6001600160801b03828116828216039081111561144057611440614791565b6001600160801b038181168382160290811690818114614b0157614b01614791565b5092915050565b6001600160801b03818116838216019081111561144057611440614791565b5f6001600160801b03831680614b3f57614b3f614a34565b806001600160801b0384160491505092915050565b5f6001600160801b03831680614b6c57614b6c614a34565b806001600160801b038416069150509291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca2646970667358221220b85ce566241cb74ec11834f70798de9be5c9b0494cacd0dd034911314e5aff8364736f6c634300081b0033")
}

// ContractAddress is the SFC contract address
var ContractAddress = common.HexToAddress("0xfc00face00000000000000000000000000000000")
