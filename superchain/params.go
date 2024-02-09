package superchain

import (
	"math/big"
)

var uint128Max, ok = big.NewInt(0).SetString("ffffffffffffffffffffffffffffffff", 16)

func init() {
	if !ok {
		panic("cannot construct uint128Max")
	}
}

type ResourceConfig struct {
	MaxResourceLimit            uint32
	ElasticityMultiplier        uint8
	BaseFeeMaxChangeDenominator uint8
	MinimumBaseFee              uint32
	SystemTxMaxGas              uint32
	MaximumBaseFee              *big.Int
}

// OPMainnetResourceConfig describes the resource metering configuration from OP Mainnet
var OPMainnetResourceConfig = ResourceConfig{
	MaxResourceLimit:            20000000,
	ElasticityMultiplier:        10,
	BaseFeeMaxChangeDenominator: 8,
	MinimumBaseFee:              1000000000,
	SystemTxMaxGas:              1000000,
	MaximumBaseFee:              uint128Max,
}

type L2OOParams struct {
	SubmissionInterval        *big.Int // Interval in blocks at which checkpoints must be submitted.
	L2BlockTime               *big.Int // The time per L2 block, in seconds.
	FinalizationPeriodSeconds *big.Int // The minimum time (in seconds) that must elapse before a withdrawal can be finalized.
}

// OPMainnetL2OOParams describes the L2OutputOracle parameters from OP Mainnet
var OPMainnetL2OOParams = L2OOParams{
	SubmissionInterval:        big.NewInt(120),
	L2BlockTime:               big.NewInt(2),
	FinalizationPeriodSeconds: big.NewInt(12),
}

// OPGoerliL2OOParams describes the L2OutputOracle parameters from OP Goerli
var OPGoerliL2OOParams = L2OOParams{
	SubmissionInterval:        big.NewInt(120),
	L2BlockTime:               big.NewInt(2),
	FinalizationPeriodSeconds: big.NewInt(12),
}

// OPGoerliDev0L2OOParams describes the L2OutputOracle parameters from OP Goerli
var OPGoerliDev0L2OOParams = L2OOParams{
	SubmissionInterval:        big.NewInt(120),
	L2BlockTime:               big.NewInt(2),
	FinalizationPeriodSeconds: big.NewInt(12),
}

// OPSepoliaL2OOParams describes the L2OutputOracle parameters from OP Goerli
var OPSepoliaL2OOParams = L2OOParams{
	SubmissionInterval:        big.NewInt(120),
	L2BlockTime:               big.NewInt(2),
	FinalizationPeriodSeconds: big.NewInt(12),
}

// OPSepoliaDev0L2OOParams describes the L2OutputOracle parameters from OP Goerli
var OPSepoliaDev0L2OOParams = L2OOParams{
	SubmissionInterval:        big.NewInt(120),
	L2BlockTime:               big.NewInt(2),
	FinalizationPeriodSeconds: big.NewInt(12),
}

var OPGoerliBytecode = map[string]string{
	"L1CrossDomainMessenger": "0x60806040526004361061015f5760003560e01c80636e296e45116100c0578063b1b1b20911610074578063c4d66de811610059578063c4d66de8146103ed578063d764ad0b1461040d578063ecc704281461042057600080fd5b8063b1b1b2091461039d578063b28ade25146103cd57600080fd5b80638cbeeef2116100a55780638cbeeef2146102505780639fce812c14610329578063a4e7f8bd1461035d57600080fd5b80636e296e45146102fd57806383a740741461031257600080fd5b80633f827a5a1161011757806354fd4d50116100fc57806354fd4d50146102665780635644cfdf146102bc5780636425666b146102d257600080fd5b80633f827a5a146102285780634c1d6a691461025057600080fd5b80630ff754ea116101485780630ff754ea146101ac5780632828d7e8146101fe5780633dbb202b1461021357600080fd5b8063028f85f7146101645780630c56849814610197575b600080fd5b34801561017057600080fd5b50610179601081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b3480156101a357600080fd5b50610179603f81565b3480156101b857600080fd5b5060f9546101d99073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161018e565b34801561020a57600080fd5b50610179604081565b6102266102213660046117ba565b610485565b005b34801561023457600080fd5b5061023d600181565b60405161ffff909116815260200161018e565b34801561025c57600080fd5b50610179619c4081565b34801561027257600080fd5b506102af6040518060400160405280600581526020017f312e372e3000000000000000000000000000000000000000000000000000000081525081565b60405161018e919061188c565b3480156102c857600080fd5b5061017961138881565b3480156102de57600080fd5b5060f95473ffffffffffffffffffffffffffffffffffffffff166101d9565b34801561030957600080fd5b506101d96106e9565b34801561031e57600080fd5b5061017962030d4081565b34801561033557600080fd5b506101d97f000000000000000000000000420000000000000000000000000000000000000781565b34801561036957600080fd5b5061038d6103783660046118a6565b60ce6020526000908152604090205460ff1681565b604051901515815260200161018e565b3480156103a957600080fd5b5061038d6103b83660046118a6565b60cb6020526000908152604090205460ff1681565b3480156103d957600080fd5b506101796103e83660046118bf565b6107d5565b3480156103f957600080fd5b50610226610408366004611913565b610843565b61022661041b366004611930565b610a47565b34801561042c57600080fd5b5061047760cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b60405190815260200161018e565b6105be7f00000000000000000000000042000000000000000000000000000000000000076104b48585856107d5565b347fd764ad0b0000000000000000000000000000000000000000000000000000000061052060cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c60405160240161053c97969594939291906119ff565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526112d3565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a33858561064360cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b86604051610655959493929190611a5e565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2153016107b8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f7420736574000000000000000000000060648201526084015b60405180910390fd5b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b6000611388619c4080603f6107f1604063ffffffff8816611adb565b6107fb9190611b0b565b610806601088611adb565b6108139062030d40611b59565b61081d9190611b59565b6108279190611b59565b6108319190611b59565b61083b9190611b59565b949350505050565b6000546003907501000000000000000000000000000000000000000000900460ff16158015610891575060005460ff8083167401000000000000000000000000000000000000000090920416105b61091d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016107af565b6000805475010000000000000000000000000000000000000000007fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff9091167401000000000000000000000000000000000000000060ff8516027fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff161717905560f980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790556109e561136c565b600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b60f087901c60028110610b02576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206f722031206d657373616765732061726520737570706f7274656460648201527f20617420746869732074696d6500000000000000000000000000000000000000608482015260a4016107af565b8061ffff16600003610bf7576000610b53878986868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508f9250611445915050565b600081815260cb602052604090205490915060ff1615610bf5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a206c65676163792077697460448201527f6864726177616c20616c72656164792072656c6179656400000000000000000060648201526084016107af565b505b6000610c3d898989898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061146492505050565b9050610c47611487565b15610c7f57853414610c5b57610c5b611b85565b600081815260ce602052604090205460ff1615610c7a57610c7a611b85565b610dd1565b3415610d33576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a4016107af565b600081815260ce602052604090205460ff16610dd1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c617965640000000000000000000000000000000060648201526084016107af565b610dda8761157d565b15610e8d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a4016107af565b600081815260cb602052604090205460ff1615610f2c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c617965640000000000000000000060648201526084016107af565b610f4d85610f3e611388619c40611b59565b67ffffffffffffffff166115c3565b1580610f73575060cc5473ffffffffffffffffffffffffffffffffffffffff1661dead14155b1561108c57600081815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555182917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff3201611085576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d6573736167650000000000000000000000000000000000000060648201526084016107af565b50506112c5565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a16179055600061111d88619c405a6110e09190611bb4565b8988888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506115e192505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055905080156111b457600082815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a26112c1565b600082815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff32016112c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d6573736167650000000000000000000000000000000000000060648201526084016107af565b5050505b50505050505050565b905090565b60f9546040517fe9e05c4200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063e9e05c42908490611334908890839089906000908990600401611bcb565b6000604051808303818588803b15801561134d57600080fd5b505af1158015611361573d6000803e3d6000fd5b505050505050505050565b6000547501000000000000000000000000000000000000000000900460ff16611417576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016107af565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055565b6000611453858585856115fb565b805190602001209050949350505050565b6000611474878787878787611694565b8051906020012090509695505050505050565b60f95460009073ffffffffffffffffffffffffffffffffffffffff16331480156112ce575060f954604080517f9bf62d82000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000042000000000000000000000000000000000000078116931691639bf62d829160048083019260209291908290030181865afa15801561153d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115619190611c23565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b600073ffffffffffffffffffffffffffffffffffffffff82163014806115bd575060f95473ffffffffffffffffffffffffffffffffffffffff8381169116145b92915050565b600080603f83619c4001026040850201603f5a021015949350505050565b600080600080845160208601878a8af19695505050505050565b6060848484846040516024016116149493929190611c40565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcbd4ece9000000000000000000000000000000000000000000000000000000001790529050949350505050565b60608686868686866040516024016116b196959493929190611c8a565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461175557600080fd5b50565b60008083601f84011261176a57600080fd5b50813567ffffffffffffffff81111561178257600080fd5b60208301915083602082850101111561179a57600080fd5b9250929050565b803563ffffffff811681146117b557600080fd5b919050565b600080600080606085870312156117d057600080fd5b84356117db81611733565b9350602085013567ffffffffffffffff8111156117f757600080fd5b61180387828801611758565b90945092506118169050604086016117a1565b905092959194509250565b6000815180845260005b818110156118475760208185018101518683018201520161182b565b81811115611859576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061189f6020830184611821565b9392505050565b6000602082840312156118b857600080fd5b5035919050565b6000806000604084860312156118d457600080fd5b833567ffffffffffffffff8111156118eb57600080fd5b6118f786828701611758565b909450925061190a9050602085016117a1565b90509250925092565b60006020828403121561192557600080fd5b813561189f81611733565b600080600080600080600060c0888a03121561194b57600080fd5b87359650602088013561195d81611733565b9550604088013561196d81611733565b9450606088013593506080880135925060a088013567ffffffffffffffff81111561199757600080fd5b6119a38a828b01611758565b989b979a50959850939692959293505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a0830152611a5160c0830184866119b6565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff86168152608060208201526000611a8e6080830186886119b6565b905083604083015263ffffffff831660608301529695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681851681830481118215151615611b0257611b02611aac565b02949350505050565b600067ffffffffffffffff80841680611b4d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600067ffffffffffffffff808316818516808303821115611b7c57611b7c611aac565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600082821015611bc657611bc6611aac565b500390565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015267ffffffffffffffff84166040820152821515606082015260a060808201526000611c1860a0830184611821565b979650505050505050565b600060208284031215611c3557600080fd5b815161189f81611733565b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525060806040830152611c796080830185611821565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a0830152611cd560c0830184611821565b9897505050505050505056fea164736f6c634300080f000a",
}
