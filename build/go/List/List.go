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
	Bin: "0x608060405234801561001057600080fd5b50611fd4806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630c7a28c61461005157806335deb19c1461006d5780635baf5697146100895780638129fc1c146100b9575b600080fd5b61006b6004803603810190610066919061198d565b6100c3565b005b610087600480360381019061008291906119f7565b6100c6565b005b6100a3600480360381019061009e91906119b6565b611077565b6040516100b09190611bbf565b60405180910390f35b6100c161143a565b005b50565b6100f27f34b303d87768a9387d3532ff1783244e8b9f6af12112d5f4376de1d6af87d33560001b6100c3565b61011e7ff837e141650416f4668e84249e5099f9298c1a5f32ff9ca2a86d87cbee46684560001b6100c3565b61014a7f266e67e862b5888e203965e3fe73c18338bf946e247cf45dafc6386b70b8403f60001b6100c3565b60006004811115610184577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b816020015160048111156101c1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415610460576101f37f108a0d6496761ebc122a3ebc610dc59577a9031555b7ee32481684eb90150bce60001b6100c3565b61021f7ff5ccc275a8a6dfb2de099f11ebac5cb9200f3e32f93dc61ae7aab02adb0e486060001b6100c3565b61024a7e852b27ced6b0497c1aff8a57005242c3d4928f7a49e579df7af6cf0bd3603360001b6100c3565b6000600182600001516040516102609190611ba8565b9081526020016040518091039020905061029c7fe71c4c0d1cd3f251d9d36481efe9822b54df350f9b2061af273fabe136285db060001b6100c3565b6102c87f17b0583a98d5e2bbdc0d8f907eab3d5e582b2deac8e391aff5aa9f93565d86e460001b6100c3565b60005b826040015151811015610459576103047f8e0df832d7e343b164b18aa2b85dd40ffe7d4c5a4f8967329abe1b8b08b23d9860001b6100c3565b6103307f134eef0c39b40a7073eba665d4156061ea510c2ccbd67dfeda6ae564d0f6f35560001b6100c3565b818360400151828151811061036e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060408201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050808061045190611df0565b9150506102cb565b505061048d565b61048c7f639d759a8d22a0f1bd1ea45df2c743eaec86675ccfd25ad7b98b9c41cd770d2e60001b6100c3565b5b6104b97fb2b5f5d90d6848f44bb796783fb174fdbfd5bbfe9e677224f21ef1d8f2ea439e60001b6100c3565b6104e57f5b2e72df73d1e619a71e73879df74bd4f250f239ced316e25c1429f17d2e72ae60001b6100c3565b6001600481111561051f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8160200151600481111561055c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14156109a35761058e7f5949f321574d65f0d75c1a17f5005e94793de91ef2a64fd3cc4a240089f5680960001b6100c3565b6105ba7fcafb0dbfd7442060f48fd462fd17f663ec4f23790365ac21649ca75ab436bdeb60001b6100c3565b6105e67fc3d3ea2f054cf1e290050cc031a60b65f99ab8138b737227894ecfed9dd6399960001b6100c3565b6000600182600001516040516105fc9190611ba8565b908152602001604051809103902090506106387f3de513e6896472921a56b53fcf4baeff47e818e92828e5ee4da56b4e29b0373360001b6100c3565b6106647f72d22ff8b3cc49bcb04b09d0c0eb2d67ce45a4edb1daf853270f23d4f1c9b81d60001b6100c3565b60005b82604001515181101561099c576106a07f75c1ac7208e305bcf4c55be82339bd34ffede7f977219d600a916bfd8e98734060001b6100c3565b6106cc7fc0b0ddc3369cd8ecead1f858fe5e29bb2793f3eab3d8d4da956f9dee4f9eefec60001b6100c3565b60005b8280549050811015610988576107077f3d2fbcade80887b41a78959a6a5bab38b7614d25138bf9c17124bf17a5d0d47360001b6100c3565b6107337f8da4d7e4a7907d8a26fd856e6fd9c403193b0fc1f4f9010ac36c26288574077560001b6100c3565b83604001518281518110610770577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff168382815481106107cb577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906002020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610949576108457fd599c7d8fa95c936b322cfcb81ded05e3a4fe0a7b4a425f688c7760bb479b9be60001b6100c3565b6108717f5c8c358a6412eb27b0675783bc37eb912a3b248ab1c9b74a67550b5d45f48ca460001b6100c3565b8281815481106108aa577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000209060020201600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549067ffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff021916905550506109447fff3747144e3230ad4fe886c9c82b0863cc0ec92c361eef9189085b94984f4a4860001b6100c3565b610988565b6109757ff1a436d2d6bbbac4829f4da766631956a4feed61baa7d86cbdafd036ea38a5ba60001b6100c3565b808061098090611df0565b9150506106cf565b50808061099490611df0565b915050610667565b50506109d0565b6109cf7f40ea141d9fa382653f1410976fe3f623f7c9070c7ed5cdbebcab136f9e787a9260001b6100c3565b5b6109fc7f8494b76139d6fcd2f046b94bb1e60b15bc6ee4aaaba1ab1a3d98eea910a705a160001b6100c3565b610a287fb38d47f0db1ad3350be2306429636221ef9279169d874263deb7896543ac395b60001b6100c3565b60026004811115610a62577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81602001516004811115610a9f577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b1415610eba57610ad17f69597b5e015d60327c0f9037c4be51dec568ee12caa724888052abd51abba9a060001b6100c3565b610afd7fe6c02eb78dd8dda5cbb2edc651f51f2c3244481b43494ec066668dbfc7b74cf260001b6100c3565b610b297f181f0ddee4f86d9ce891ddf32a8ca783a89d3650352c52cfbb1a49fb574ff95a60001b6100c3565b600060018260000151604051610b3f9190611ba8565b90815260200160405180910390209050610b7b7f1d1faadb042b6422399221ce652d00b09f5c35141f001e4c458c4865595454af60001b6100c3565b610ba77f5766847d60913d4e4ba872da64485103c5424621be531114133e85acf581b58560001b6100c3565b60005b8180549050811015610e2957610be27ff9722036d0e64e56db790a7652a5b65c29dedd3bf0dde4a18bfa03664d71e2f660001b6100c3565b610c0e7fde5628854e2bac807fa53a982882fdb59377eda8fc4e469e3f3322850ec1c03d60001b6100c3565b818181548110610c47577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906002020160000160149054906101000a900467ffffffffffffffff1667ffffffffffffffff16828281548110610cb0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906002020160010160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1611610de957610d117f2e75c63a9a8481e805a91d5d67a387a049aae54eee14beb7b933acf8389a199060001b6100c3565b610d3d7f1ee25cf4836ff3fdf98b125eb008be3c41de96a052b1e0816eae79161516362f60001b6100c3565b818181548110610d76577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000209060020201600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549067ffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff02191690555050610e16565b610e157fcd8ea7d006eb48f824bc9784e157b227d9bd933360b825ceb7721ebb5794e05660001b6100c3565b5b8080610e2190611df0565b915050610baa565b50610e567fcae6d919719b4ec26529a9b961208d144966eeecdc958c3eab7839b4c04cdfa060001b6100c3565b610e827f7940cc244e73dd507c3ce6a0fc77a703788bd655c2cffad3e53b5a2b29bae52a60001b6100c3565b8060018360000151604051610e979190611ba8565b9081526020016040518091039020908054610eb392919061156e565b5050610ee7565b610ee67f85a60413248d87d0874cf1187516087c73fb99b56cbc9205d04bb75e440b0a3c60001b6100c3565b5b610f137fd4dfd7f39a8d512e08898f0a336b50d762d162ceef583cb249c8e33f783b4a9560001b6100c3565b610f3f7f7931a010de8ecafdc064a3f9f022e1b05f21d6cc3ad843617c1081cf0d4c121e60001b6100c3565b60036004811115610f79577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81602001516004811115610fb6577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b141561104757610fe87fbe5d6cca5b31a6074f3fdbb35083f0373c971005dff78f45c28687cc45d5e62860001b6100c3565b6110147fe295fd91917e233180d234a9c8c296fdf8c56e559b1727a0d8451ace9956abd860001b6100c3565b600181600001516040516110289190611ba8565b9081526020016040518091039020600061104291906116b3565b611074565b6110737f0dc24cc61875df5f0268eda82e81a05d3aea8a8cc4f93535296a805a49ffd47560001b6100c3565b5b50565b60606110a57f8ed166baba1dec77eb5839717de677b0790b1a1ec8b8680167613dc54fb7271a60001b6100c3565b6110d17f1d236a7cd410f005c221817bf768447755f1cd40b75a53ba29a5b7697d8278d260001b6100c3565b6110fd7f88b6bde6746628b0b64eb2d4fd3bce326459dcd16065aa60af7b99ec044fd5f660001b6100c3565b6111297f511eee4a7ce09400744fcb39a7007b662a6ab284e94154ea8fb298a36946491960001b6100c3565b600082511161116d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161116490611c01565b60405180910390fd5b6111997f2ec8d12430cb33724d18085c377a53287aa0119bea74c48caef4680bf11ea7b960001b6100c3565b6111c57fdd3ec1d2e55e9e335aad19854221077f9b31788bd66f72fbea5abd00f462fdd460001b6100c3565b6111f17f9b5d255c2976c51d116ad3dba008119581582e9300a5e69c7321419d4ad959e860001b6100c3565b61121d7fb6118e87520a9d151b894b0769f24a1779847c863e333818e1290246bb7e4f3960001b6100c3565b600060018360405161122f9190611ba8565b90815260200160405180910390208054905011611281576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161127890611be1565b60405180910390fd5b6112ad7f7d7aa1a05395cd31289a108d32b69da53ebce6ccdaf9cf8eb0010e1f3a3344a160001b6100c3565b6112d97fab1f40f64a9bbfcbb5839274b51f9b629fb8df18fa98a0f73394ce56ad1ff37960001b6100c3565b6113057f7ba88d8c8964f39eb8cb69c6a8aaed7de8e348525d3da52e40fea44b97ac94d560001b6100c3565b6001826040516113159190611ba8565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b8282101561142f57838290600052602060002090600202016040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505081526020019060010190611343565b505050509050919050565b600060019054906101000a900460ff166114625760008054906101000a900460ff161561146b565b61146a61154a565b5b6114aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114a190611c21565b60405180910390fd5b60008060019054906101000a900460ff1615905080156114fa576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b6115267fe6148bda7227aed3455cab6e82f8395274c1ee2c9af3b14c3fd64496b99ca89560001b6100c3565b80156115475760008060016101000a81548160ff0219169083151502179055505b50565b60006115553061155b565b15905090565b600080823b905060008111915050919050565b8280548282559060005260206000209060020281019282156116a25760005260206000209160020282015b828111156116a15782826000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000820160149054906101000a900467ffffffffffffffff168160000160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001820160009054906101000a900467ffffffffffffffff168160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505091600201919060020190611599565b5b5090506116af91906116d7565b5090565b50805460008255600202906000526020600020908101906116d491906116d7565b50565b5b8082111561174757600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549067ffffffffffffffff02191690556001820160006101000a81549067ffffffffffffffff0219169055506002016116d8565b5090565b600061175e61175984611c66565b611c41565b9050808382526020820190508285606086028201111561177d57600080fd5b60005b858110156117ad57816117938882611918565b845260208401935060608301925050600181019050611780565b5050509392505050565b60006117ca6117c584611c92565b611c41565b9050828152602081018484840111156117e257600080fd5b6117ed848285611d7d565b509392505050565b60008135905061180481611f49565b92915050565b600082601f83011261181b57600080fd5b813561182b84826020860161174b565b91505092915050565b60008135905061184381611f60565b92915050565b600082601f83011261185a57600080fd5b813561186a8482602086016117b7565b91505092915050565b60008135905061188281611f77565b92915050565b60006060828403121561189a57600080fd5b6118a46060611c41565b9050600082013567ffffffffffffffff8111156118c057600080fd5b6118cc84828501611849565b60008301525060206118e084828501611873565b602083015250604082013567ffffffffffffffff81111561190057600080fd5b61190c8482850161180a565b60408301525092915050565b60006060828403121561192a57600080fd5b6119346060611c41565b90506000611944848285016117f5565b600083015250602061195884828501611978565b602083015250604061196c84828501611978565b60408301525092915050565b60008135905061198781611f87565b92915050565b60006020828403121561199f57600080fd5b60006119ad84828501611834565b91505092915050565b6000602082840312156119c857600080fd5b600082013567ffffffffffffffff8111156119e257600080fd5b6119ee84828501611849565b91505092915050565b600060208284031215611a0957600080fd5b600082013567ffffffffffffffff811115611a2357600080fd5b611a2f84828501611888565b91505092915050565b6000611a448383611b57565b60608301905092915050565b611a5981611d23565b82525050565b6000611a6a82611cd3565b611a748185611cf6565b9350611a7f83611cc3565b8060005b83811015611ab0578151611a978882611a38565b9750611aa283611ce9565b925050600181019050611a83565b5085935050505092915050565b6000611ac882611cde565b611ad28185611d07565b9350611ae2818560208601611d8c565b80840191505092915050565b6000611afb601283611d12565b9150611b0682611ea8565b602082019050919050565b6000611b1e601683611d12565b9150611b2982611ed1565b602082019050919050565b6000611b41602e83611d12565b9150611b4c82611efa565b604082019050919050565b606082016000820151611b6d6000850182611a50565b506020820151611b806020850182611b99565b506040820151611b936040850182611b99565b50505050565b611ba281611d69565b82525050565b6000611bb48284611abd565b915081905092915050565b60006020820190508181036000830152611bd98184611a5f565b905092915050565b60006020820190508181036000830152611bfa81611aee565b9050919050565b60006020820190508181036000830152611c1a81611b11565b9050919050565b60006020820190508181036000830152611c3a81611b34565b9050919050565b6000611c4b611c5c565b9050611c578282611dbf565b919050565b6000604051905090565b600067ffffffffffffffff821115611c8157611c80611e68565b5b602082029050602081019050919050565b600067ffffffffffffffff821115611cad57611cac611e68565b5b611cb682611e97565b9050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611d2e82611d3f565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b82818337600083830152505050565b60005b83811015611daa578082015181840152602081019050611d8f565b83811115611db9576000848401525b50505050565b611dc882611e97565b810181811067ffffffffffffffff82111715611de757611de6611e68565b5b80604052505050565b6000611dfb82611d5f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611e2e57611e2d611e39565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f77686974654c69737420697320656d7074790000000000000000000000000000600082015250565b7f66696c6548617368206d75737420626520656d70747900000000000000000000600082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b611f5281611d23565b8114611f5d57600080fd5b50565b611f6981611d35565b8114611f7457600080fd5b50565b60058110611f8457600080fd5b50565b611f9081611d69565b8114611f9b57600080fd5b5056fea2646970667358221220b8e07b296da91871d5eae23255496fc30656c61d8eddeb67fc65cf8c21b513a764736f6c63430008040033",
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
