// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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

// WhiteList is an auto generated low-level Go binding around an user-defined struct.
type WhiteList struct {
	Addr         common.Address
	BaseHeight   uint64
	ExpireHeight uint64
}

// WhiteListParams is an auto generated low-level Go binding around an user-defined struct.
type WhiteListParams struct {
	FileHash []byte
	Op       uint8
	List     []WhiteList
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileHash\",\"type\":\"bytes\"}],\"name\":\"GetWhiteList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"FileHash\",\"type\":\"bytes\"},{\"internalType\":\"enumWhiteListOpType\",\"name\":\"Op\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"BaseHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ExpireHeight\",\"type\":\"uint64\"}],\"internalType\":\"structWhiteList[]\",\"name\":\"List\",\"type\":\"tuple[]\"}],\"internalType\":\"structWhiteListParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"WhiteListOperate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"c__0xe6f45c7f\",\"type\":\"bytes32\"}],\"name\":\"c_0xe6f45c7f\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611fd5806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630c7a28c61461005157806335deb19c1461006d5780635baf5697146100895780638129fc1c146100b9575b600080fd5b61006b6004803603810190610066919061198e565b6100c3565b005b610087600480360381019061008291906119f8565b6100c6565b005b6100a3600480360381019061009e91906119b7565b611078565b6040516100b09190611bc0565b60405180910390f35b6100c161143b565b005b50565b6100f27f5a4f101a18d93929d3f3cbc1e08b77385002ea5985e1e2ac99e19aa26dc863ce60001b6100c3565b61011e7f392a0c8a5bd31778c0361416d1bc447d2c5e63ab213ba270119f34bc3104ac2360001b6100c3565b61014a7fae888716793846db60ecdf1a033210ba634985ef4a5ae07abb824c5f00796c0360001b6100c3565b60006004811115610184577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b816020015160048111156101c1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415610461576101f37f27e607cce62616d1a7f2075298d67bcd52b2311e9da97eaba90c06e0a561cf2960001b6100c3565b61021f7f170e9e7c26df8a258fd924d79c3757155b2d9b0dae9e224d59329ae88201ff4760001b6100c3565b61024b7f57a3fba6539af432095145374b544f4caf943aeaa393cc552992d8665e8898bf60001b6100c3565b6000600182600001516040516102619190611ba9565b9081526020016040518091039020905061029d7f06e56e833b12fab9bc852292bd7d5f9b15e21506dc1fe39f5d8ae245b377f8e960001b6100c3565b6102c97f023a3c1d17e20f0cce974a57597ad252973048a366fb23c32db5301622928bc160001b6100c3565b60005b82604001515181101561045a576103057f6d52a2c09de7456a73f9ab2d1b1d4e51df191942a22a52bbb994fdb94784508260001b6100c3565b6103317fa4210e1fbaf25e44d4bedc85631d3807e9f547df3acc5501e89ce15b73aeea2b60001b6100c3565b818360400151828151811061036f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050808061045290611df1565b9150506102cc565b505061048e565b61048d7f8c6f75d9fcb13ef51cd3b77ef2ddd9c73448c12c4757e279aa8cfc6ee68a2c2d60001b6100c3565b5b6104ba7fcbde6d8123f7173963ff672431be281f6db115a88f7e419d1ca91f5ca8f9e7a360001b6100c3565b6104e67fd967696e6f9218356eb3a4965d91260494962ca46f4485f65817105faab6bf5060001b6100c3565b60016004811115610520577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8160200151600481111561055d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156109a45761058f7f34070bf6665a4a99d2a7a347a561b2745910de8bc63d38cbac00dfba6d7f602160001b6100c3565b6105bb7f08f72f384f166c85e52233f4c7d4221ecac2e15d02b3d730763ffec16cf0e5e360001b6100c3565b6105e77f711a794e92bbeec68f4d452e73688aa712b4497f312cfb6cbc00cf7779ddc2df60001b6100c3565b6000600182600001516040516105fd9190611ba9565b908152602001604051809103902090506106397f0308bcb1fa822d496caabe69e127ae7ce2df40c0cc22ec31ed6a572d1c9562bf60001b6100c3565b6106657fc5252aa63f9af4f77c083874ca49c17033455415504480fe3414e13967d7a37660001b6100c3565b60005b82604001515181101561099d576106a17f99a9a12b8bd4b732deae45502fa80a2cf4af472df1b8fcbaaee015221559543c60001b6100c3565b6106cd7f98de212467c0b1ddddd11e61f7599b6f93fed670630ff579b5c8692fc91720f260001b6100c3565b60005b8280549050811015610989576107087f7d3df5c06142599bee73a4b124ad7bec3b6774d5863409292fbb40452232aa8060001b6100c3565b6107347f1cd8b0224bbbe9e9ba1fa2d2e885d4269be143cf4114f9221ba41644c5c5b73d60001b6100c3565b83604001518281518110610771577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff168382815481106107cc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906002020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561094a576108467ffc941cc3955f842f82cc913c82884ec1921a764c1be0c6b6905a2252b0aebdc460001b6100c3565b6108727f70db5dbe642bb8e2df526f22055cd9fd8f04d5e643b4106237e7ee5c80d6bad260001b6100c3565b8281815481106108ab577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000209060020201600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549067ffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff021916905550506109457f33248e4011054842d98be520ede92c763c2c79e6aa1da1c677ce79b79c82699760001b6100c3565b610989565b6109767ffd1966157fb8a6246457b9e3710a923dd5a6e65c283e9386154bacb7b24f058d60001b6100c3565b808061098190611df1565b9150506106d0565b50808061099590611df1565b915050610668565b50506109d1565b6109d07fde072bedf97c29075bbcf43258fb145d6aebd9a7074404579c258f2efddebf0c60001b6100c3565b5b6109fd7f7e3e0e89b95f89f56a5607a1bc3e215af294d5ba78e5f999f1d5097abcb61a9b60001b6100c3565b610a297fc40368409f7c2a979a16c958d848aacc59eb0e8df553c55644fd85eee98d7d4860001b6100c3565b60026004811115610a63577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81602001516004811115610aa0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415610ebb57610ad27fdfbb19046e70bb5f0bdcf3395389f66e9f8442933e5d4e9755c63023186c161e60001b6100c3565b610afe7f999a38972d43e9258f079a1db3d0e95e4ee8c403d0501fb6658a54274ea7df8e60001b6100c3565b610b2a7fc89f34f44d7fb8d2d29701b0564f59647cf412e89fa645a131227cbb1e973d1360001b6100c3565b600060018260000151604051610b409190611ba9565b90815260200160405180910390209050610b7c7f172c9f4bfa0aa26f0775a9f3afa9340d2a759d2d219bc8d69142b80acfcdd71460001b6100c3565b610ba87f17673e9ce8b2db59bd609fd9c6032ab4777337586b1700e4e9d6688a87816c6460001b6100c3565b60005b8180549050811015610e2a57610be37f4da3e5b5ce5fe9aa7a78f4b851e7370cfbd4cf2993b2aa2bc2748ec635400a8760001b6100c3565b610c0f7f706a38bed56db71f888bd5fc09d5f55748e468813ac63f175ae9bd3025284b3760001b6100c3565b818181548110610c48577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906002020160000160149054906101000a900467ffffffffffffffff1667ffffffffffffffff16828281548110610cb1577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906002020160010160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1611610dea57610d127f0eb371abaf305bd498cef55562a8017c3f801ab167868e68ea7daf387c0a3d3460001b6100c3565b610d3e7fcff7a659347bbfb54f47e7a8821906f8d679791d6d832c8746cd74c370c7817760001b6100c3565b818181548110610d77577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000209060020201600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549067ffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff02191690555050610e17565b610e167f0a202b397da5fe10e4b58cc065564d8f3c9094d1115a879d7f882e6d7336743160001b6100c3565b5b8080610e2290611df1565b915050610bab565b50610e577ff5bac4d0d45ae268d39705b521a7ad989e803c52453a2b54ec4df5865d7a0c9f60001b6100c3565b610e837fc583b14985527487fcdbd0c92563c254a2cea956bf05a152514b94543641982060001b6100c3565b8060018360000151604051610e989190611ba9565b9081526020016040518091039020908054610eb492919061156f565b5050610ee8565b610ee77f71bfd4c7ed9e991d0656750eeb4f37f744755058c6c5675b00a31f4f71d2a79a60001b6100c3565b5b610f147f364dfc70c0b2c36b8b59370ebd931e6142cb2be61a7449d25a85c39057b4246060001b6100c3565b610f407f951a8968e77fae7ba46b3bf4ba8f9e59ed0a532dcb684a1690803bf7009d12b560001b6100c3565b60036004811115610f7a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81602001516004811115610fb7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b141561104857610fe97f6dc0ed75a407d638194dac2d538a4688eedd860e51f064578bac52f08b76607860001b6100c3565b6110157f6f841bed5c6e96c4752cbcee6bf6536bd73de873463b23a74b6e72edb29b311560001b6100c3565b600181600001516040516110299190611ba9565b9081526020016040518091039020600061104391906116b4565b611075565b6110747fb1d2e4e934c776948ad9ce38e5193c8f10b0ac1a9bca3d51bb61259f7ae1d39d60001b6100c3565b5b50565b60606110a67fe104b7c07670bd8a325aea1cdcf7ec02732ce84b2a9b2cea6f919b6ff701adb960001b6100c3565b6110d27f9fece8810fc585959f89fcc6a5aa85b44244a2649aff668845b44e5b298f74f760001b6100c3565b6110fe7f2009f889782796c458f3c4c6288559d11774bc4052dae36fc3125098ef4a86f160001b6100c3565b61112a7f5644bc08837188350630e31428515c02929500dc6e2f3655dd2269ad8ceff1f660001b6100c3565b600082511161116e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161116590611c02565b60405180910390fd5b61119a7fab4cae8990c60cf7a34474535d9977a345644993a4ecd9a6777cf7b265a94c2a60001b6100c3565b6111c67f2431b02b896a91f13e692cc7732ed2c4e5b9176b7ecee54651ff1f80af35900060001b6100c3565b6111f27f12c4b79e4c1d30fd870d4ad3ddacac78389d4b34ad852e042ad4c1f3a60730a960001b6100c3565b61121e7fd6f6dee114bd1c494b662ea9a418d21b37c845be390db6a96c9145f2b712478460001b6100c3565b60006001836040516112309190611ba9565b90815260200160405180910390208054905011611282576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161127990611be2565b60405180910390fd5b6112ae7f685d0dc81a0a267848edbca8e224fcf418376188b06ec69a9bcf82c4a9e86c0160001b6100c3565b6112da7f68dbce1ed6ccec1321dcaafb966f47848f7a9dcc2a4e71a73d2be7f2a4f7bca060001b6100c3565b6113067f74f14d48f7fe8fdff40382f4613fa488e26df52ddb672706153276d7c78c7e1d60001b6100c3565b6001826040516113169190611ba9565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b8282101561143057838290600052602060002090600202016040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505081526020019060010190611344565b505050509050919050565b600060019054906101000a900460ff166114635760008054906101000a900460ff161561146c565b61146b61154b565b5b6114ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114a290611c22565b60405180910390fd5b60008060019054906101000a900460ff1615905080156114fb576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b6115277f112b3894bcd4cf4b0a3acabca5cd216bfe3037a9007922c01be65808db60b8aa60001b6100c3565b80156115485760008060016101000a81548160ff0219169083151502179055505b50565b60006115563061155c565b15905090565b600080823b905060008111915050919050565b8280548282559060005260206000209060020281019282156116a35760005260206000209160020282015b828111156116a25782826000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000820160149054906101000a900467ffffffffffffffff168160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160009054906101000a900467ffffffffffffffff168160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050509160020191906002019061159a565b5b5090506116b091906116d8565b5090565b50805460008255600202906000526020600020908101906116d591906116d8565b50565b5b8082111561174857600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549067ffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff0219169055506002016116d9565b5090565b600061175f61175a84611c67565b611c42565b9050808382526020820190508285606086028201111561177e57600080fd5b60005b858110156117ae57816117948882611919565b845260208401935060608301925050600181019050611781565b5050509392505050565b60006117cb6117c684611c93565b611c42565b9050828152602081018484840111156117e357600080fd5b6117ee848285611d7e565b509392505050565b60008135905061180581611f4a565b92915050565b600082601f83011261181c57600080fd5b813561182c84826020860161174c565b91505092915050565b60008135905061184481611f61565b92915050565b600082601f83011261185b57600080fd5b813561186b8482602086016117b8565b91505092915050565b60008135905061188381611f78565b92915050565b60006060828403121561189b57600080fd5b6118a56060611c42565b9050600082013567ffffffffffffffff8111156118c157600080fd5b6118cd8482850161184a565b60008301525060206118e184828501611874565b602083015250604082013567ffffffffffffffff81111561190157600080fd5b61190d8482850161180b565b60408301525092915050565b60006060828403121561192b57600080fd5b6119356060611c42565b90506000611945848285016117f6565b600083015250602061195984828501611979565b602083015250604061196d84828501611979565b60408301525092915050565b60008135905061198881611f88565b92915050565b6000602082840312156119a057600080fd5b60006119ae84828501611835565b91505092915050565b6000602082840312156119c957600080fd5b600082013567ffffffffffffffff8111156119e357600080fd5b6119ef8482850161184a565b91505092915050565b600060208284031215611a0a57600080fd5b600082013567ffffffffffffffff811115611a2457600080fd5b611a3084828501611889565b91505092915050565b6000611a458383611b58565b60608301905092915050565b611a5a81611d24565b82525050565b6000611a6b82611cd4565b611a758185611cf7565b9350611a8083611cc4565b8060005b83811015611ab1578151611a988882611a39565b9750611aa383611cea565b925050600181019050611a84565b5085935050505092915050565b6000611ac982611cdf565b611ad38185611d08565b9350611ae3818560208601611d8d565b80840191505092915050565b6000611afc601283611d13565b9150611b0782611ea9565b602082019050919050565b6000611b1f601683611d13565b9150611b2a82611ed2565b602082019050919050565b6000611b42602e83611d13565b9150611b4d82611efb565b604082019050919050565b606082016000820151611b6e6000850182611a51565b506020820151611b816020850182611b9a565b506040820151611b946040850182611b9a565b50505050565b611ba381611d6a565b82525050565b6000611bb58284611abe565b915081905092915050565b60006020820190508181036000830152611bda8184611a60565b905092915050565b60006020820190508181036000830152611bfb81611aef565b9050919050565b60006020820190508181036000830152611c1b81611b12565b9050919050565b60006020820190508181036000830152611c3b81611b35565b9050919050565b6000611c4c611c5d565b9050611c588282611dc0565b919050565b6000604051905090565b600067ffffffffffffffff821115611c8257611c81611e69565b5b602082029050602081019050919050565b600067ffffffffffffffff821115611cae57611cad611e69565b5b611cb782611e98565b9050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611d2f82611d40565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b82818337600083830152505050565b60005b83811015611dab578082015181840152602081019050611d90565b83811115611dba576000848401525b50505050565b611dc982611e98565b810181811067ffffffffffffffff82111715611de857611de7611e69565b5b80604052505050565b6000611dfc82611d60565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611e2f57611e2e611e3a565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f77686974654c69737420697320656d7074790000000000000000000000000000600082015250565b7f66696c6548617368206d75737420626520656d70747900000000000000000000600082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b611f5381611d24565b8114611f5e57600080fd5b50565b611f6a81611d36565b8114611f7557600080fd5b50565b60058110611f8557600080fd5b50565b611f9181611d6a565b8114611f9c57600080fd5b5056fea26469706673582212209590962d575b8b31717539d88a50040ccddfb987de8de5275c72851409ef256064736f6c63430008040033",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// GetWhiteList is a free data retrieval call binding the contract method 0x5baf5697.
//
// Solidity: function GetWhiteList(bytes fileHash) view returns((address,uint64,uint64)[])
func (_Store *StoreCaller) GetWhiteList(opts *bind.CallOpts, fileHash []byte) ([]WhiteList, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GetWhiteList", fileHash)

	if err != nil {
		return *new([]WhiteList), err
	}

	out0 := *abi.ConvertType(out[0], new([]WhiteList)).(*[]WhiteList)

	return out0, err

}

// GetWhiteList is a free data retrieval call binding the contract method 0x5baf5697.
//
// Solidity: function GetWhiteList(bytes fileHash) view returns((address,uint64,uint64)[])
func (_Store *StoreSession) GetWhiteList(fileHash []byte) ([]WhiteList, error) {
	return _Store.Contract.GetWhiteList(&_Store.CallOpts, fileHash)
}

// GetWhiteList is a free data retrieval call binding the contract method 0x5baf5697.
//
// Solidity: function GetWhiteList(bytes fileHash) view returns((address,uint64,uint64)[])
func (_Store *StoreCallerSession) GetWhiteList(fileHash []byte) ([]WhiteList, error) {
	return _Store.Contract.GetWhiteList(&_Store.CallOpts, fileHash)
}

// C0xe6f45c7f is a free data retrieval call binding the contract method 0x0c7a28c6.
//
// Solidity: function c_0xe6f45c7f(bytes32 c__0xe6f45c7f) pure returns()
func (_Store *StoreCaller) C0xe6f45c7f(opts *bind.CallOpts, c__0xe6f45c7f [32]byte) error {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "c_0xe6f45c7f", c__0xe6f45c7f)

	if err != nil {
		return err
	}

	return err

}

// C0xe6f45c7f is a free data retrieval call binding the contract method 0x0c7a28c6.
//
// Solidity: function c_0xe6f45c7f(bytes32 c__0xe6f45c7f) pure returns()
func (_Store *StoreSession) C0xe6f45c7f(c__0xe6f45c7f [32]byte) error {
	return _Store.Contract.C0xe6f45c7f(&_Store.CallOpts, c__0xe6f45c7f)
}

// C0xe6f45c7f is a free data retrieval call binding the contract method 0x0c7a28c6.
//
// Solidity: function c_0xe6f45c7f(bytes32 c__0xe6f45c7f) pure returns()
func (_Store *StoreCallerSession) C0xe6f45c7f(c__0xe6f45c7f [32]byte) error {
	return _Store.Contract.C0xe6f45c7f(&_Store.CallOpts, c__0xe6f45c7f)
}

// WhiteListOperate is a paid mutator transaction binding the contract method 0x35deb19c.
//
// Solidity: function WhiteListOperate((bytes,uint8,(address,uint64,uint64)[]) params) returns()
func (_Store *StoreTransactor) WhiteListOperate(opts *bind.TransactOpts, params WhiteListParams) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "WhiteListOperate", params)
}

// WhiteListOperate is a paid mutator transaction binding the contract method 0x35deb19c.
//
// Solidity: function WhiteListOperate((bytes,uint8,(address,uint64,uint64)[]) params) returns()
func (_Store *StoreSession) WhiteListOperate(params WhiteListParams) (*types.Transaction, error) {
	return _Store.Contract.WhiteListOperate(&_Store.TransactOpts, params)
}

// WhiteListOperate is a paid mutator transaction binding the contract method 0x35deb19c.
//
// Solidity: function WhiteListOperate((bytes,uint8,(address,uint64,uint64)[]) params) returns()
func (_Store *StoreTransactorSession) WhiteListOperate(params WhiteListParams) (*types.Transaction, error) {
	return _Store.Contract.WhiteListOperate(&_Store.TransactOpts, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Store *StoreSession) Initialize() (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Store *StoreTransactorSession) Initialize() (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts)
}
