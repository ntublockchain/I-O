// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cb2p_auction

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

// Cb2pAuctionMetaData contains all meta data concerning the Cb2pAuction contract.
var Cb2pAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AwaitResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"prcd\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"DecisionMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HighestBidIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"NewBidHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"rating\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"review\",\"type\":\"string\"}],\"name\":\"RateAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"}],\"name\":\"RevealAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawBid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"asset_id\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"asset_owner\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"checkAverageScore\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"not_winner_platform\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"secondHighestPrice\",\"type\":\"uint256\"}],\"name\":\"closeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auction_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_asset_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_asset_owner\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_score\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_feedback\",\"type\":\"string\"}],\"name\":\"provide_feedback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"revealAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"secondHighestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051611cac380380611cac83398101604081905261002f91610052565b6001600160a01b0316608052600080546001600160a01b03191633179055610082565b60006020828403121561006457600080fd5b81516001600160a01b038116811461007b57600080fd5b9392505050565b608051611c016100ab600039600081816102dd015281816105fb0152610f0e0152611c016000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c80638d10b0a6116100a2578063b14c63c511610071578063b14c63c51461027f578063d0a1414a1461029f578063d960d573146102b2578063ea1591bb146102c5578063fc0c546a146102d857600080fd5b80638d10b0a6146102185780638da5cb5b1461022b5780639348cef71461023e578063966ffcff1461025157600080fd5b8063451df52e116100e9578063451df52e1461018b57806355f78c8d146101cc5780635a4463a8146101df5780637efbf8ac146101f257806388d3c98e1461020557600080fd5b8063176321e91461011b5780632e1a7d4d1461013057806342d21ef71461015857806344a770bf14610178575b600080fd5b61012e61012936600461151c565b6102ff565b005b61014361013e366004611563565b6104fc565b60405190151581526020015b60405180910390f35b61016b610166366004611563565b6106da565b60405161014f91906115c2565b61016b610186366004611563565b610774565b6101b4610199366004611563565b6003602052600090815260409020546001600160a01b031681565b6040516001600160a01b03909116815260200161014f565b61012e6101da366004611563565b61078d565b61012e6101ed3660046115dc565b6108a4565b61012e610200366004611649565b610996565b61012e6102133660046116a0565b610b87565b61016b610226366004611563565b610db4565b6000546101b4906001600160a01b031681565b61012e61024c3660046116d8565b610dcd565b61027161025f366004611563565b60046020526000908152604090205481565b60405190815260200161014f565b61027161028d366004611563565b60056020526000908152604090205481565b61012e6102ad36600461151c565b611094565b6102716102c0366004611563565b6112e2565b61012e6102d33660046116d8565b6113a6565b6101b47f000000000000000000000000000000000000000000000000000000000000000081565b60405165656e64696e6760d01b602082015260260160408051601f198184030181528282528051602091820120600086815260068352929092209192610346929101611734565b60405160208183030381529060405280519060200120146103ae5760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e472073746174757300000060448201526064015b60405180910390fd5b6000828152600360205260409020546001600160a01b031633146103e45760405162461bcd60e51b81526004016103a5906117aa565b60405180604001604052806007815260200166636c6f73696e6760c81b8152506006600084815260200190815260200160002090816104239190611829565b506000828152600360209081526040808320546005835281842054600790935281842091517f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7946104889488946001600160a01b039094169390929091908890611966565b60405180910390a160008281526005602090815260408083205460038352818420546001600160a01b03168452600190925282208054919290916104cd9084906119d0565b909155505050600090815260036020908152604080832080546001600160a01b03191690556005909152812055565b6040516337b832b760e11b602082015260009060240160408051601f198184030181528282528051602091820120600086815260068352929092209192610544929101611734565b60405160208183030381529060405280519060200120036105a75760405162461bcd60e51b815260206004820152601760248201527f436f6e747261637420696e204f50454e2073746174757300000000000000000060448201526064016103a5565b336000908152600160205260409020548015610685573360008181526001602052604080822091909155516323b872dd60e01b81523060048201526024810191909152604481018290526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303816000875af1158015610644573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066891906119e9565b610685573360009081526001602052604081209190915592915050565b6000838152600760205260409081902090517f9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c916106c99186919033908690611a06565b60405180910390a150600192915050565b600660205260009081526040902080546106f3906116fa565b80601f016020809104026020016040519081016040528092919081815260200182805461071f906116fa565b801561076c5780601f106107415761010080835404028352916020019161076c565b820191906000526020600020905b81548152906001019060200180831161074f57829003601f168201915b505050505081565b600860205260009081526040902080546106f3906116fa565b6000546001600160a01b031633146107b75760405162461bcd60e51b81526004016103a590611a3b565b6040516337b832b760e11b602082015260240160408051601f1981840301815282825280516020918201206000858152600683529290922091926107fc929101611734565b604051602081830303815290604052805190602001201461082f5760405162461bcd60e51b81526004016103a590611a82565b604051806040016040528060068152602001651c995d99585b60d21b81525060066000838152602001908152602001600020908161086d9190611829565b506040518181527f3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba9060200160405180910390a150565b6000546001600160a01b031633146109085760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79206f776e65722063616e20637265617465206e65772061756374696f6044820152603760f91b60648201526084016103a5565b600083815260036020908152604080832080546001600160a01b03191690556005825280832083905580518082018252600481526337b832b760e11b8184015286845260069092529091209061095e9082611829565b5060008381526007602052604090206109778382611829565b5060008381526008602052604090206109908282611829565b50505050565b60405166636c6f73696e6760c81b602082015260270160408051601f1981840301815282825280516020918201206000878152600683529290922091926109de929101611734565b6040516020818303038152906040528051906020012014610a415760405162461bcd60e51b815260206004820152601e60248201527f436f6e7472616374206e6f7420696e20434c4f53494e4720737461747573000060448201526064016103a5565b6000838152600360205260409020546001600160a01b03163314610a775760405162461bcd60e51b81526004016103a5906117aa565b600083815260086020526040908190209051600a91610a9591611734565b9081526040805191829003602090810183208054600181018255600091825282822001869055868152600890915220600991610ad19190611734565b908152604051602091819003820190208054600181018255600091825291902001610afc8282611829565b506000838152600760205260409081902090517fcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa791610b419186919086908690611ab9565b60405180910390a16040518060400160405280600681526020016518db1bdcd95960d21b8152506006600085815260200190815260200160002090816109909190611829565b6000546001600160a01b03163314610bb15760405162461bcd60e51b81526004016103a590611a3b565b6040516337b832b760e11b602082015260240160408051601f198184030181528282528051602091820120600087815260068352929092209192610bf6929101611734565b6040516020818303038152906040528051906020012014610c295760405162461bcd60e51b81526004016103a590611a82565b60405180604001604052806006815260200165656e64696e6760d01b815250600660008581526020019081526020016000209081610c679190611829565b508180610c805750600083815260056020526040902054155b15610d31576040518060400160405280600681526020016518db1bdcd95960d21b815250600660008581526020019081526020016000209081610cc39190611829565b5060008381526005602090815260408083205460038352818420546001600160a01b0316845260019092528220805491929091610d019084906119d0565b90915550505060009182525060036020908152604080832080546001600160a01b03191690556005909152812055565b600083815260046020526040902054811115610d595760008381526004602052604090208190555b6000838152600360209081526040918290205482518681526001600160a01b03909116918101919091527fa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b484746850027391015b60405180910390a1505050565b600760205260009081526040902080546106f3906116fa565b604051651c995d99585b60d21b602082015260260160408051601f198184030181528282528051602091820120600086815260068352929092209192610e14929101611734565b6040516020818303038152906040528051906020012014610e475760405162461bcd60e51b81526004016103a590611a82565b6000828152600560205260409020548111610ea45760405162461bcd60e51b815260206004820152601e60248201527f546865726520616c7265616479206973206120686967686572206269642e000060448201526064016103a5565b600082815260026020908152604080832033845282529182902054825191820184905291016040516020818303038152906040528051906020012014610ee957600080fd5b6040516323b872dd60e01b8152336004820152306024820152604481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303816000875af1158015610f5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f8391906119e9565b905080610fcb5760405162461bcd60e51b81526020600482015260166024820152752a37b5b2b7103a3930b739b332b9103330b4b632b21760511b60448201526064016103a5565b600083815260056020526040902054156110225760008381526005602090815260408083205460038352818420546001600160a01b031684526001909252822080549192909161101c9084906119d0565b90915550505b600083815260036020908152604080832080546001600160a01b03191633908117909155600583528184208054600485528386205586905560079092529182902091517f463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde0892610da79287928790611a06565b60405165656e64696e6760d01b602082015260260160408051601f1981840301815282825280516020918201206000868152600683529290922091926110db929101611734565b604051602081830303815290604052805190602001201461113e5760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e472073746174757300000060448201526064016103a5565b6000828152600360205260409020546001600160a01b031633146111745760405162461bcd60e51b81526004016103a5906117aa565b60405180604001604052806007815260200166636c6f73696e6760c81b8152506006600084815260200190815260200160002090816111b39190611829565b506000828152600460209081526040808320546005909252909120541115611234576000828152600460209081526040808320546005909252909120546111fa9190611af5565b6000838152600360209081526040808320546001600160a01b0316835260019091528120805490919061122e9084906119d0565b90915550505b60008281526004602090815260408083205483546001600160a01b031684526001909252822080549192909161126b9084906119d0565b9091555050600082815260036020908152604080832054600583528184205460079093529281902090517f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7936112d69387936001600160a01b03909216929091906001908890611966565b60405180910390a15050565b600081815260086020526040808220905182918291600a9161130391611734565b90815260405190819003602001902054905060005b8181101561138757600085815260086020526040908190209051600a9161133e91611734565b9081526020016040518091039020818154811061135d5761135d611b08565b9060005260206000200154836113739190611b1e565b92508061137f81611b46565b915050611318565b5080611394836064611b5f565b61139e9190611b8f565b949350505050565b6040516337b832b760e11b602082015260240160408051601f1981840301815282825280516020918201206000868152600683529290922091926113eb929101611734565b604051602081830303815290604052805190602001201461141e5760405162461bcd60e51b81526004016103a590611a82565b60008281526002602090815260408083203380855290835281842085905585845260079092529182902091517f6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f926112d69286928690611a06565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126114a057600080fd5b813567ffffffffffffffff808211156114bb576114bb611479565b604051601f8301601f19908116603f011681019082821181831017156114e3576114e3611479565b816040528381528660208588010111156114fc57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561152f57600080fd5b82359150602083013567ffffffffffffffff81111561154d57600080fd5b6115598582860161148f565b9150509250929050565b60006020828403121561157557600080fd5b5035919050565b6000815180845260005b818110156115a257602081850181015186830182015201611586565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006115d5602083018461157c565b9392505050565b6000806000606084860312156115f157600080fd5b83359250602084013567ffffffffffffffff8082111561161057600080fd5b61161c8783880161148f565b9350604086013591508082111561163257600080fd5b5061163f8682870161148f565b9150509250925092565b60008060006060848603121561165e57600080fd5b8335925060208401359150604084013567ffffffffffffffff81111561168357600080fd5b61163f8682870161148f565b801515811461169d57600080fd5b50565b6000806000606084860312156116b557600080fd5b8335925060208401356116c78161168f565b929592945050506040919091013590565b600080604083850312156116eb57600080fd5b50508035926020909101359150565b600181811c9082168061170e57607f821691505b60208210810361172e57634e487b7160e01b600052602260045260246000fd5b50919050565b6000808354611742816116fa565b6001828116801561175a576001811461176f5761179e565b60ff198416875282151583028701945061179e565b8760005260208060002060005b858110156117955781548a82015290840190820161177c565b50505082870194505b50929695505050505050565b6020808252601690820152754e6f7420617574686f72697a6564206163636573732160501b604082015260600190565b601f82111561182457600081815260208120601f850160051c810160208610156118015750805b601f850160051c820191505b818110156118205782815560010161180d565b5050505b505050565b815167ffffffffffffffff81111561184357611843611479565b6118578161185184546116fa565b846117da565b602080601f83116001811461188c57600084156118745750858301515b600019600386901b1c1916600185901b178555611820565b600085815260208120601f198616915b828110156118bb5788860151825594840194600190910190840161189c565b50858210156118d95787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600081546118f6816116fa565b808552602060018381168015611913576001811461192d5761195b565b60ff1985168884015283151560051b88018301955061195b565b866000528260002060005b858110156119535781548a8201860152908301908401611938565b890184019650505b505050505092915050565b86815260018060a01b038616602082015284604082015260c06060820152600061199360c08301866118e9565b841515608084015282810360a08401526119ad818561157c565b9998505050505050505050565b634e487b7160e01b600052601160045260246000fd5b808201808211156119e3576119e36119ba565b92915050565b6000602082840312156119fb57600080fd5b81516115d58161168f565b848152608060208201526000611a1f60808301866118e9565b6001600160a01b03949094166040830152506060015292915050565b60208082526027908201527f4f6e6c79206f776e65722063616e206368616e676520636f6e747261637427736040820152662073746174757360c81b606082015260800190565b6020808252601b908201527f436f6e7472616374206e6f7420696e204f50454e207374617475730000000000604082015260600190565b848152608060208201526000611ad260808301866118e9565b8460408401528281036060840152611aea818561157c565b979650505050505050565b818103818111156119e3576119e36119ba565b634e487b7160e01b600052603260045260246000fd5b8082018281126000831280158216821582161715611b3e57611b3e6119ba565b505092915050565b600060018201611b5857611b586119ba565b5060010190565b80820260008212600160ff1b84141615611b7b57611b7b6119ba565b81810583148215176119e3576119e36119ba565b600082611bac57634e487b7160e01b600052601260045260246000fd5b600160ff1b821460001984141615611bc657611bc66119ba565b50059056fea2646970667358221220ecf49492ec562621b9e0014bb3f8c6320b93a13de76d6eb990e2a85d3c7e904964736f6c63430008120033",
}

// Cb2pAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use Cb2pAuctionMetaData.ABI instead.
var Cb2pAuctionABI = Cb2pAuctionMetaData.ABI

// Cb2pAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Cb2pAuctionMetaData.Bin instead.
var Cb2pAuctionBin = Cb2pAuctionMetaData.Bin

// DeployCb2pAuction deploys a new Ethereum contract, binding an instance of Cb2pAuction to it.
func DeployCb2pAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *Cb2pAuction, error) {
	parsed, err := Cb2pAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Cb2pAuctionBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cb2pAuction{Cb2pAuctionCaller: Cb2pAuctionCaller{contract: contract}, Cb2pAuctionTransactor: Cb2pAuctionTransactor{contract: contract}, Cb2pAuctionFilterer: Cb2pAuctionFilterer{contract: contract}}, nil
}

// Cb2pAuction is an auto generated Go binding around an Ethereum contract.
type Cb2pAuction struct {
	Cb2pAuctionCaller     // Read-only binding to the contract
	Cb2pAuctionTransactor // Write-only binding to the contract
	Cb2pAuctionFilterer   // Log filterer for contract events
}

// Cb2pAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type Cb2pAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cb2pAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Cb2pAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cb2pAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Cb2pAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cb2pAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Cb2pAuctionSession struct {
	Contract     *Cb2pAuction      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Cb2pAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Cb2pAuctionCallerSession struct {
	Contract *Cb2pAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Cb2pAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Cb2pAuctionTransactorSession struct {
	Contract     *Cb2pAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Cb2pAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type Cb2pAuctionRaw struct {
	Contract *Cb2pAuction // Generic contract binding to access the raw methods on
}

// Cb2pAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Cb2pAuctionCallerRaw struct {
	Contract *Cb2pAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// Cb2pAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Cb2pAuctionTransactorRaw struct {
	Contract *Cb2pAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCb2pAuction creates a new instance of Cb2pAuction, bound to a specific deployed contract.
func NewCb2pAuction(address common.Address, backend bind.ContractBackend) (*Cb2pAuction, error) {
	contract, err := bindCb2pAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cb2pAuction{Cb2pAuctionCaller: Cb2pAuctionCaller{contract: contract}, Cb2pAuctionTransactor: Cb2pAuctionTransactor{contract: contract}, Cb2pAuctionFilterer: Cb2pAuctionFilterer{contract: contract}}, nil
}

// NewCb2pAuctionCaller creates a new read-only instance of Cb2pAuction, bound to a specific deployed contract.
func NewCb2pAuctionCaller(address common.Address, caller bind.ContractCaller) (*Cb2pAuctionCaller, error) {
	contract, err := bindCb2pAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionCaller{contract: contract}, nil
}

// NewCb2pAuctionTransactor creates a new write-only instance of Cb2pAuction, bound to a specific deployed contract.
func NewCb2pAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*Cb2pAuctionTransactor, error) {
	contract, err := bindCb2pAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionTransactor{contract: contract}, nil
}

// NewCb2pAuctionFilterer creates a new log filterer instance of Cb2pAuction, bound to a specific deployed contract.
func NewCb2pAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*Cb2pAuctionFilterer, error) {
	contract, err := bindCb2pAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionFilterer{contract: contract}, nil
}

// bindCb2pAuction binds a generic wrapper to an already deployed contract.
func bindCb2pAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Cb2pAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cb2pAuction *Cb2pAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cb2pAuction.Contract.Cb2pAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cb2pAuction *Cb2pAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Cb2pAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cb2pAuction *Cb2pAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Cb2pAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cb2pAuction *Cb2pAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cb2pAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cb2pAuction *Cb2pAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cb2pAuction *Cb2pAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.contract.Transact(opts, method, params...)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_Cb2pAuction *Cb2pAuctionCaller) AssetId(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Cb2pAuction.contract.Call(opts, &out, "asset_id", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_Cb2pAuction *Cb2pAuctionSession) AssetId(arg0 *big.Int) (string, error) {
	return _Cb2pAuction.Contract.AssetId(&_Cb2pAuction.CallOpts, arg0)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_Cb2pAuction *Cb2pAuctionCallerSession) AssetId(arg0 *big.Int) (string, error) {
	return _Cb2pAuction.Contract.AssetId(&_Cb2pAuction.CallOpts, arg0)
}

// AssetOwner is a free data retrieval call binding the contract method 0x44a770bf.
//
// Solidity: function asset_owner(uint256 ) view returns(string)
func (_Cb2pAuction *Cb2pAuctionCaller) AssetOwner(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Cb2pAuction.contract.Call(opts, &out, "asset_owner", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AssetOwner is a free data retrieval call binding the contract method 0x44a770bf.
//
// Solidity: function asset_owner(uint256 ) view returns(string)
func (_Cb2pAuction *Cb2pAuctionSession) AssetOwner(arg0 *big.Int) (string, error) {
	return _Cb2pAuction.Contract.AssetOwner(&_Cb2pAuction.CallOpts, arg0)
}

// AssetOwner is a free data retrieval call binding the contract method 0x44a770bf.
//
// Solidity: function asset_owner(uint256 ) view returns(string)
func (_Cb2pAuction *Cb2pAuctionCallerSession) AssetOwner(arg0 *big.Int) (string, error) {
	return _Cb2pAuction.Contract.AssetOwner(&_Cb2pAuction.CallOpts, arg0)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xd960d573.
//
// Solidity: function checkAverageScore(uint256 auctionId) view returns(int256)
func (_Cb2pAuction *Cb2pAuctionCaller) CheckAverageScore(opts *bind.CallOpts, auctionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cb2pAuction.contract.Call(opts, &out, "checkAverageScore", auctionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckAverageScore is a free data retrieval call binding the contract method 0xd960d573.
//
// Solidity: function checkAverageScore(uint256 auctionId) view returns(int256)
func (_Cb2pAuction *Cb2pAuctionSession) CheckAverageScore(auctionId *big.Int) (*big.Int, error) {
	return _Cb2pAuction.Contract.CheckAverageScore(&_Cb2pAuction.CallOpts, auctionId)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xd960d573.
//
// Solidity: function checkAverageScore(uint256 auctionId) view returns(int256)
func (_Cb2pAuction *Cb2pAuctionCallerSession) CheckAverageScore(auctionId *big.Int) (*big.Int, error) {
	return _Cb2pAuction.Contract.CheckAverageScore(&_Cb2pAuction.CallOpts, auctionId)
}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_Cb2pAuction *Cb2pAuctionCaller) HighestBid(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cb2pAuction.contract.Call(opts, &out, "highestBid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_Cb2pAuction *Cb2pAuctionSession) HighestBid(arg0 *big.Int) (*big.Int, error) {
	return _Cb2pAuction.Contract.HighestBid(&_Cb2pAuction.CallOpts, arg0)
}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_Cb2pAuction *Cb2pAuctionCallerSession) HighestBid(arg0 *big.Int) (*big.Int, error) {
	return _Cb2pAuction.Contract.HighestBid(&_Cb2pAuction.CallOpts, arg0)
}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_Cb2pAuction *Cb2pAuctionCaller) HighestBidder(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Cb2pAuction.contract.Call(opts, &out, "highestBidder", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_Cb2pAuction *Cb2pAuctionSession) HighestBidder(arg0 *big.Int) (common.Address, error) {
	return _Cb2pAuction.Contract.HighestBidder(&_Cb2pAuction.CallOpts, arg0)
}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_Cb2pAuction *Cb2pAuctionCallerSession) HighestBidder(arg0 *big.Int) (common.Address, error) {
	return _Cb2pAuction.Contract.HighestBidder(&_Cb2pAuction.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cb2pAuction *Cb2pAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cb2pAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cb2pAuction *Cb2pAuctionSession) Owner() (common.Address, error) {
	return _Cb2pAuction.Contract.Owner(&_Cb2pAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cb2pAuction *Cb2pAuctionCallerSession) Owner() (common.Address, error) {
	return _Cb2pAuction.Contract.Owner(&_Cb2pAuction.CallOpts)
}

// SecondHighestBid is a free data retrieval call binding the contract method 0x966ffcff.
//
// Solidity: function secondHighestBid(uint256 ) view returns(uint256)
func (_Cb2pAuction *Cb2pAuctionCaller) SecondHighestBid(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cb2pAuction.contract.Call(opts, &out, "secondHighestBid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondHighestBid is a free data retrieval call binding the contract method 0x966ffcff.
//
// Solidity: function secondHighestBid(uint256 ) view returns(uint256)
func (_Cb2pAuction *Cb2pAuctionSession) SecondHighestBid(arg0 *big.Int) (*big.Int, error) {
	return _Cb2pAuction.Contract.SecondHighestBid(&_Cb2pAuction.CallOpts, arg0)
}

// SecondHighestBid is a free data retrieval call binding the contract method 0x966ffcff.
//
// Solidity: function secondHighestBid(uint256 ) view returns(uint256)
func (_Cb2pAuction *Cb2pAuctionCallerSession) SecondHighestBid(arg0 *big.Int) (*big.Int, error) {
	return _Cb2pAuction.Contract.SecondHighestBid(&_Cb2pAuction.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_Cb2pAuction *Cb2pAuctionCaller) Status(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Cb2pAuction.contract.Call(opts, &out, "status", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_Cb2pAuction *Cb2pAuctionSession) Status(arg0 *big.Int) (string, error) {
	return _Cb2pAuction.Contract.Status(&_Cb2pAuction.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_Cb2pAuction *Cb2pAuctionCallerSession) Status(arg0 *big.Int) (string, error) {
	return _Cb2pAuction.Contract.Status(&_Cb2pAuction.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Cb2pAuction *Cb2pAuctionCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cb2pAuction.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Cb2pAuction *Cb2pAuctionSession) Token() (common.Address, error) {
	return _Cb2pAuction.Contract.Token(&_Cb2pAuction.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Cb2pAuction *Cb2pAuctionCallerSession) Token() (common.Address, error) {
	return _Cb2pAuction.Contract.Token(&_Cb2pAuction.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_Cb2pAuction *Cb2pAuctionTransactor) Abort(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb2pAuction.contract.Transact(opts, "abort", auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_Cb2pAuction *Cb2pAuctionSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Abort(&_Cb2pAuction.TransactOpts, auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_Cb2pAuction *Cb2pAuctionTransactorSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Abort(&_Cb2pAuction.TransactOpts, auctionId, jsonString)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_Cb2pAuction *Cb2pAuctionTransactor) Bid(opts *bind.TransactOpts, auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _Cb2pAuction.contract.Transact(opts, "bid", auctionId, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_Cb2pAuction *Cb2pAuctionSession) Bid(auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Bid(&_Cb2pAuction.TransactOpts, auctionId, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_Cb2pAuction *Cb2pAuctionTransactorSession) Bid(auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Bid(&_Cb2pAuction.TransactOpts, auctionId, bidHash)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x88d3c98e.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform, uint256 secondHighestPrice) returns()
func (_Cb2pAuction *Cb2pAuctionTransactor) CloseAuction(opts *bind.TransactOpts, auctionId *big.Int, not_winner_platform bool, secondHighestPrice *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.contract.Transact(opts, "closeAuction", auctionId, not_winner_platform, secondHighestPrice)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x88d3c98e.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform, uint256 secondHighestPrice) returns()
func (_Cb2pAuction *Cb2pAuctionSession) CloseAuction(auctionId *big.Int, not_winner_platform bool, secondHighestPrice *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.CloseAuction(&_Cb2pAuction.TransactOpts, auctionId, not_winner_platform, secondHighestPrice)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x88d3c98e.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform, uint256 secondHighestPrice) returns()
func (_Cb2pAuction *Cb2pAuctionTransactorSession) CloseAuction(auctionId *big.Int, not_winner_platform bool, secondHighestPrice *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.CloseAuction(&_Cb2pAuction.TransactOpts, auctionId, not_winner_platform, secondHighestPrice)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_Cb2pAuction *Cb2pAuctionTransactor) Commit(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb2pAuction.contract.Transact(opts, "commit", auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_Cb2pAuction *Cb2pAuctionSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Commit(&_Cb2pAuction.TransactOpts, auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_Cb2pAuction *Cb2pAuctionTransactorSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Commit(&_Cb2pAuction.TransactOpts, auctionId, jsonString)
}

// Create is a paid mutator transaction binding the contract method 0x5a4463a8.
//
// Solidity: function create(uint256 _auction_id, string _asset_id, string _asset_owner) returns()
func (_Cb2pAuction *Cb2pAuctionTransactor) Create(opts *bind.TransactOpts, _auction_id *big.Int, _asset_id string, _asset_owner string) (*types.Transaction, error) {
	return _Cb2pAuction.contract.Transact(opts, "create", _auction_id, _asset_id, _asset_owner)
}

// Create is a paid mutator transaction binding the contract method 0x5a4463a8.
//
// Solidity: function create(uint256 _auction_id, string _asset_id, string _asset_owner) returns()
func (_Cb2pAuction *Cb2pAuctionSession) Create(_auction_id *big.Int, _asset_id string, _asset_owner string) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Create(&_Cb2pAuction.TransactOpts, _auction_id, _asset_id, _asset_owner)
}

// Create is a paid mutator transaction binding the contract method 0x5a4463a8.
//
// Solidity: function create(uint256 _auction_id, string _asset_id, string _asset_owner) returns()
func (_Cb2pAuction *Cb2pAuctionTransactorSession) Create(_auction_id *big.Int, _asset_id string, _asset_owner string) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Create(&_Cb2pAuction.TransactOpts, _auction_id, _asset_id, _asset_owner)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x7efbf8ac.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, string _feedback) returns()
func (_Cb2pAuction *Cb2pAuctionTransactor) ProvideFeedback(opts *bind.TransactOpts, auctionId *big.Int, _score *big.Int, _feedback string) (*types.Transaction, error) {
	return _Cb2pAuction.contract.Transact(opts, "provide_feedback", auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x7efbf8ac.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, string _feedback) returns()
func (_Cb2pAuction *Cb2pAuctionSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback string) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.ProvideFeedback(&_Cb2pAuction.TransactOpts, auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x7efbf8ac.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, string _feedback) returns()
func (_Cb2pAuction *Cb2pAuctionTransactorSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback string) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.ProvideFeedback(&_Cb2pAuction.TransactOpts, auctionId, _score, _feedback)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_Cb2pAuction *Cb2pAuctionTransactor) Reveal(opts *bind.TransactOpts, auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.contract.Transact(opts, "reveal", auctionId, bidAmount)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_Cb2pAuction *Cb2pAuctionSession) Reveal(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Reveal(&_Cb2pAuction.TransactOpts, auctionId, bidAmount)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_Cb2pAuction *Cb2pAuctionTransactorSession) Reveal(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Reveal(&_Cb2pAuction.TransactOpts, auctionId, bidAmount)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_Cb2pAuction *Cb2pAuctionTransactor) RevealAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.contract.Transact(opts, "revealAuction", auctionId)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_Cb2pAuction *Cb2pAuctionSession) RevealAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.RevealAuction(&_Cb2pAuction.TransactOpts, auctionId)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_Cb2pAuction *Cb2pAuctionTransactorSession) RevealAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.RevealAuction(&_Cb2pAuction.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_Cb2pAuction *Cb2pAuctionTransactor) Withdraw(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.contract.Transact(opts, "withdraw", auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_Cb2pAuction *Cb2pAuctionSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Withdraw(&_Cb2pAuction.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_Cb2pAuction *Cb2pAuctionTransactorSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _Cb2pAuction.Contract.Withdraw(&_Cb2pAuction.TransactOpts, auctionId)
}

// Cb2pAuctionAwaitResponseIterator is returned from FilterAwaitResponse and is used to iterate over the raw logs and unpacked data for AwaitResponse events raised by the Cb2pAuction contract.
type Cb2pAuctionAwaitResponseIterator struct {
	Event *Cb2pAuctionAwaitResponse // Event containing the contract specifics and raw log

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
func (it *Cb2pAuctionAwaitResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb2pAuctionAwaitResponse)
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
		it.Event = new(Cb2pAuctionAwaitResponse)
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
func (it *Cb2pAuctionAwaitResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb2pAuctionAwaitResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb2pAuctionAwaitResponse represents a AwaitResponse event raised by the Cb2pAuction contract.
type Cb2pAuctionAwaitResponse struct {
	Auction *big.Int
	Winner  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAwaitResponse is a free log retrieval operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_Cb2pAuction *Cb2pAuctionFilterer) FilterAwaitResponse(opts *bind.FilterOpts) (*Cb2pAuctionAwaitResponseIterator, error) {

	logs, sub, err := _Cb2pAuction.contract.FilterLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionAwaitResponseIterator{contract: _Cb2pAuction.contract, event: "AwaitResponse", logs: logs, sub: sub}, nil
}

// WatchAwaitResponse is a free log subscription operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_Cb2pAuction *Cb2pAuctionFilterer) WatchAwaitResponse(opts *bind.WatchOpts, sink chan<- *Cb2pAuctionAwaitResponse) (event.Subscription, error) {

	logs, sub, err := _Cb2pAuction.contract.WatchLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb2pAuctionAwaitResponse)
				if err := _Cb2pAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
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

// ParseAwaitResponse is a log parse operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auction, address winner)
func (_Cb2pAuction *Cb2pAuctionFilterer) ParseAwaitResponse(log types.Log) (*Cb2pAuctionAwaitResponse, error) {
	event := new(Cb2pAuctionAwaitResponse)
	if err := _Cb2pAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb2pAuctionDecisionMadeIterator is returned from FilterDecisionMade and is used to iterate over the raw logs and unpacked data for DecisionMade events raised by the Cb2pAuction contract.
type Cb2pAuctionDecisionMadeIterator struct {
	Event *Cb2pAuctionDecisionMade // Event containing the contract specifics and raw log

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
func (it *Cb2pAuctionDecisionMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb2pAuctionDecisionMade)
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
		it.Event = new(Cb2pAuctionDecisionMade)
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
func (it *Cb2pAuctionDecisionMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb2pAuctionDecisionMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb2pAuctionDecisionMade represents a DecisionMade event raised by the Cb2pAuction contract.
type Cb2pAuctionDecisionMade struct {
	Auction    *big.Int
	Winner     common.Address
	Amount     *big.Int
	Id         string
	Prcd       bool
	JsonString string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDecisionMade is a free log retrieval operation binding the contract event 0x70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7.
//
// Solidity: event DecisionMade(uint256 auction, address winner, uint256 amount, string id, bool prcd, string jsonString)
func (_Cb2pAuction *Cb2pAuctionFilterer) FilterDecisionMade(opts *bind.FilterOpts) (*Cb2pAuctionDecisionMadeIterator, error) {

	logs, sub, err := _Cb2pAuction.contract.FilterLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionDecisionMadeIterator{contract: _Cb2pAuction.contract, event: "DecisionMade", logs: logs, sub: sub}, nil
}

// WatchDecisionMade is a free log subscription operation binding the contract event 0x70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7.
//
// Solidity: event DecisionMade(uint256 auction, address winner, uint256 amount, string id, bool prcd, string jsonString)
func (_Cb2pAuction *Cb2pAuctionFilterer) WatchDecisionMade(opts *bind.WatchOpts, sink chan<- *Cb2pAuctionDecisionMade) (event.Subscription, error) {

	logs, sub, err := _Cb2pAuction.contract.WatchLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb2pAuctionDecisionMade)
				if err := _Cb2pAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
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

// ParseDecisionMade is a log parse operation binding the contract event 0x70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7.
//
// Solidity: event DecisionMade(uint256 auction, address winner, uint256 amount, string id, bool prcd, string jsonString)
func (_Cb2pAuction *Cb2pAuctionFilterer) ParseDecisionMade(log types.Log) (*Cb2pAuctionDecisionMade, error) {
	event := new(Cb2pAuctionDecisionMade)
	if err := _Cb2pAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb2pAuctionHighestBidIncreasedIterator is returned from FilterHighestBidIncreased and is used to iterate over the raw logs and unpacked data for HighestBidIncreased events raised by the Cb2pAuction contract.
type Cb2pAuctionHighestBidIncreasedIterator struct {
	Event *Cb2pAuctionHighestBidIncreased // Event containing the contract specifics and raw log

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
func (it *Cb2pAuctionHighestBidIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb2pAuctionHighestBidIncreased)
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
		it.Event = new(Cb2pAuctionHighestBidIncreased)
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
func (it *Cb2pAuctionHighestBidIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb2pAuctionHighestBidIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb2pAuctionHighestBidIncreased represents a HighestBidIncreased event raised by the Cb2pAuction contract.
type Cb2pAuctionHighestBidIncreased struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHighestBidIncreased is a free log retrieval operation binding the contract event 0x463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08.
//
// Solidity: event HighestBidIncreased(uint256 auction, string id, address bidder, uint256 amount)
func (_Cb2pAuction *Cb2pAuctionFilterer) FilterHighestBidIncreased(opts *bind.FilterOpts) (*Cb2pAuctionHighestBidIncreasedIterator, error) {

	logs, sub, err := _Cb2pAuction.contract.FilterLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionHighestBidIncreasedIterator{contract: _Cb2pAuction.contract, event: "HighestBidIncreased", logs: logs, sub: sub}, nil
}

// WatchHighestBidIncreased is a free log subscription operation binding the contract event 0x463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08.
//
// Solidity: event HighestBidIncreased(uint256 auction, string id, address bidder, uint256 amount)
func (_Cb2pAuction *Cb2pAuctionFilterer) WatchHighestBidIncreased(opts *bind.WatchOpts, sink chan<- *Cb2pAuctionHighestBidIncreased) (event.Subscription, error) {

	logs, sub, err := _Cb2pAuction.contract.WatchLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb2pAuctionHighestBidIncreased)
				if err := _Cb2pAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
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

// ParseHighestBidIncreased is a log parse operation binding the contract event 0x463d814067347ed56fee7be37a4ebb3b44d11995db0cf30046c893c14bfdde08.
//
// Solidity: event HighestBidIncreased(uint256 auction, string id, address bidder, uint256 amount)
func (_Cb2pAuction *Cb2pAuctionFilterer) ParseHighestBidIncreased(log types.Log) (*Cb2pAuctionHighestBidIncreased, error) {
	event := new(Cb2pAuctionHighestBidIncreased)
	if err := _Cb2pAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb2pAuctionNewBidHashIterator is returned from FilterNewBidHash and is used to iterate over the raw logs and unpacked data for NewBidHash events raised by the Cb2pAuction contract.
type Cb2pAuctionNewBidHashIterator struct {
	Event *Cb2pAuctionNewBidHash // Event containing the contract specifics and raw log

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
func (it *Cb2pAuctionNewBidHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb2pAuctionNewBidHash)
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
		it.Event = new(Cb2pAuctionNewBidHash)
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
func (it *Cb2pAuctionNewBidHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb2pAuctionNewBidHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb2pAuctionNewBidHash represents a NewBidHash event raised by the Cb2pAuction contract.
type Cb2pAuctionNewBidHash struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	BidHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewBidHash is a free log retrieval operation binding the contract event 0x6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f.
//
// Solidity: event NewBidHash(uint256 auction, string id, address bidder, bytes32 bidHash)
func (_Cb2pAuction *Cb2pAuctionFilterer) FilterNewBidHash(opts *bind.FilterOpts) (*Cb2pAuctionNewBidHashIterator, error) {

	logs, sub, err := _Cb2pAuction.contract.FilterLogs(opts, "NewBidHash")
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionNewBidHashIterator{contract: _Cb2pAuction.contract, event: "NewBidHash", logs: logs, sub: sub}, nil
}

// WatchNewBidHash is a free log subscription operation binding the contract event 0x6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f.
//
// Solidity: event NewBidHash(uint256 auction, string id, address bidder, bytes32 bidHash)
func (_Cb2pAuction *Cb2pAuctionFilterer) WatchNewBidHash(opts *bind.WatchOpts, sink chan<- *Cb2pAuctionNewBidHash) (event.Subscription, error) {

	logs, sub, err := _Cb2pAuction.contract.WatchLogs(opts, "NewBidHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb2pAuctionNewBidHash)
				if err := _Cb2pAuction.contract.UnpackLog(event, "NewBidHash", log); err != nil {
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

// ParseNewBidHash is a log parse operation binding the contract event 0x6738406e3bb2425ad24e77066f32691ef2126fc5e51449aac89557df63d3e04f.
//
// Solidity: event NewBidHash(uint256 auction, string id, address bidder, bytes32 bidHash)
func (_Cb2pAuction *Cb2pAuctionFilterer) ParseNewBidHash(log types.Log) (*Cb2pAuctionNewBidHash, error) {
	event := new(Cb2pAuctionNewBidHash)
	if err := _Cb2pAuction.contract.UnpackLog(event, "NewBidHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb2pAuctionRateAuctionIterator is returned from FilterRateAuction and is used to iterate over the raw logs and unpacked data for RateAuction events raised by the Cb2pAuction contract.
type Cb2pAuctionRateAuctionIterator struct {
	Event *Cb2pAuctionRateAuction // Event containing the contract specifics and raw log

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
func (it *Cb2pAuctionRateAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb2pAuctionRateAuction)
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
		it.Event = new(Cb2pAuctionRateAuction)
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
func (it *Cb2pAuctionRateAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb2pAuctionRateAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb2pAuctionRateAuction represents a RateAuction event raised by the Cb2pAuction contract.
type Cb2pAuctionRateAuction struct {
	AuctionId *big.Int
	Id        string
	Rating    *big.Int
	Review    string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRateAuction is a free log retrieval operation binding the contract event 0xcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa7.
//
// Solidity: event RateAuction(uint256 auctionId, string id, int256 rating, string review)
func (_Cb2pAuction *Cb2pAuctionFilterer) FilterRateAuction(opts *bind.FilterOpts) (*Cb2pAuctionRateAuctionIterator, error) {

	logs, sub, err := _Cb2pAuction.contract.FilterLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionRateAuctionIterator{contract: _Cb2pAuction.contract, event: "RateAuction", logs: logs, sub: sub}, nil
}

// WatchRateAuction is a free log subscription operation binding the contract event 0xcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa7.
//
// Solidity: event RateAuction(uint256 auctionId, string id, int256 rating, string review)
func (_Cb2pAuction *Cb2pAuctionFilterer) WatchRateAuction(opts *bind.WatchOpts, sink chan<- *Cb2pAuctionRateAuction) (event.Subscription, error) {

	logs, sub, err := _Cb2pAuction.contract.WatchLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb2pAuctionRateAuction)
				if err := _Cb2pAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
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

// ParseRateAuction is a log parse operation binding the contract event 0xcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa7.
//
// Solidity: event RateAuction(uint256 auctionId, string id, int256 rating, string review)
func (_Cb2pAuction *Cb2pAuctionFilterer) ParseRateAuction(log types.Log) (*Cb2pAuctionRateAuction, error) {
	event := new(Cb2pAuctionRateAuction)
	if err := _Cb2pAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb2pAuctionRevealAuctionIterator is returned from FilterRevealAuction and is used to iterate over the raw logs and unpacked data for RevealAuction events raised by the Cb2pAuction contract.
type Cb2pAuctionRevealAuctionIterator struct {
	Event *Cb2pAuctionRevealAuction // Event containing the contract specifics and raw log

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
func (it *Cb2pAuctionRevealAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb2pAuctionRevealAuction)
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
		it.Event = new(Cb2pAuctionRevealAuction)
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
func (it *Cb2pAuctionRevealAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb2pAuctionRevealAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb2pAuctionRevealAuction represents a RevealAuction event raised by the Cb2pAuction contract.
type Cb2pAuctionRevealAuction struct {
	Auction *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRevealAuction is a free log retrieval operation binding the contract event 0x3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba.
//
// Solidity: event RevealAuction(uint256 auction)
func (_Cb2pAuction *Cb2pAuctionFilterer) FilterRevealAuction(opts *bind.FilterOpts) (*Cb2pAuctionRevealAuctionIterator, error) {

	logs, sub, err := _Cb2pAuction.contract.FilterLogs(opts, "RevealAuction")
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionRevealAuctionIterator{contract: _Cb2pAuction.contract, event: "RevealAuction", logs: logs, sub: sub}, nil
}

// WatchRevealAuction is a free log subscription operation binding the contract event 0x3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba.
//
// Solidity: event RevealAuction(uint256 auction)
func (_Cb2pAuction *Cb2pAuctionFilterer) WatchRevealAuction(opts *bind.WatchOpts, sink chan<- *Cb2pAuctionRevealAuction) (event.Subscription, error) {

	logs, sub, err := _Cb2pAuction.contract.WatchLogs(opts, "RevealAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb2pAuctionRevealAuction)
				if err := _Cb2pAuction.contract.UnpackLog(event, "RevealAuction", log); err != nil {
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

// ParseRevealAuction is a log parse operation binding the contract event 0x3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba.
//
// Solidity: event RevealAuction(uint256 auction)
func (_Cb2pAuction *Cb2pAuctionFilterer) ParseRevealAuction(log types.Log) (*Cb2pAuctionRevealAuction, error) {
	event := new(Cb2pAuctionRevealAuction)
	if err := _Cb2pAuction.contract.UnpackLog(event, "RevealAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb2pAuctionWithdrawBidIterator is returned from FilterWithdrawBid and is used to iterate over the raw logs and unpacked data for WithdrawBid events raised by the Cb2pAuction contract.
type Cb2pAuctionWithdrawBidIterator struct {
	Event *Cb2pAuctionWithdrawBid // Event containing the contract specifics and raw log

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
func (it *Cb2pAuctionWithdrawBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb2pAuctionWithdrawBid)
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
		it.Event = new(Cb2pAuctionWithdrawBid)
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
func (it *Cb2pAuctionWithdrawBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb2pAuctionWithdrawBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb2pAuctionWithdrawBid represents a WithdrawBid event raised by the Cb2pAuction contract.
type Cb2pAuctionWithdrawBid struct {
	Auction *big.Int
	Id      string
	Bidder  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawBid is a free log retrieval operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_Cb2pAuction *Cb2pAuctionFilterer) FilterWithdrawBid(opts *bind.FilterOpts) (*Cb2pAuctionWithdrawBidIterator, error) {

	logs, sub, err := _Cb2pAuction.contract.FilterLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return &Cb2pAuctionWithdrawBidIterator{contract: _Cb2pAuction.contract, event: "WithdrawBid", logs: logs, sub: sub}, nil
}

// WatchWithdrawBid is a free log subscription operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_Cb2pAuction *Cb2pAuctionFilterer) WatchWithdrawBid(opts *bind.WatchOpts, sink chan<- *Cb2pAuctionWithdrawBid) (event.Subscription, error) {

	logs, sub, err := _Cb2pAuction.contract.WatchLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb2pAuctionWithdrawBid)
				if err := _Cb2pAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
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

// ParseWithdrawBid is a log parse operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auction, string id, address bidder, uint256 amount)
func (_Cb2pAuction *Cb2pAuctionFilterer) ParseWithdrawBid(log types.Log) (*Cb2pAuctionWithdrawBid, error) {
	event := new(Cb2pAuctionWithdrawBid)
	if err := _Cb2pAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
