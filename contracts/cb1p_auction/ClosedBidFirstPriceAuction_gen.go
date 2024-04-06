// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cb1p_auction

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

// Cb1pAuctionMetaData contains all meta data concerning the Cb1pAuction contract.
var Cb1pAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AwaitResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"auctionType\",\"type\":\"string\"}],\"name\":\"BidTooLow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"prcd\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"DecisionMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"auctionType\",\"type\":\"string\"}],\"name\":\"HighestBidIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"auctionType\",\"type\":\"string\"}],\"name\":\"NewBidHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"rating\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"review\",\"type\":\"string\"}],\"name\":\"RateAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"RevealAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawBid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"abort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"asset_id\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"asset_owner\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auction_type\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"bidHash\",\"type\":\"bytes32\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"checkAverageScore\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"not_winner_platform\",\"type\":\"bool\"}],\"name\":\"closeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonString\",\"type\":\"string\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auction_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_asset_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_asset_owner\",\"type\":\"string\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_score\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_feedback\",\"type\":\"string\"}],\"name\":\"provide_feedback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"revealAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200207e3803806200207e833981016040819052620000349162000085565b6001600160a01b038116608052600080546001600160a01b031916331790556040805180820190915260048152630636231760e41b60208201526001906200007d90826200015c565b505062000228565b6000602082840312156200009857600080fd5b81516001600160a01b0381168114620000b057600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620000e257607f821691505b6020821081036200010357634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200015757600081815260208120601f850160051c81016020861015620001325750805b601f850160051c820191505b8181101562000153578281556001016200013e565b5050505b505050565b81516001600160401b03811115620001785762000178620000b7565b6200019081620001898454620000cd565b8462000109565b602080601f831160018114620001c85760008415620001af5750858301515b600019600386901b1c1916600185901b17855562000153565b600085815260208120601f198616915b82811015620001f957888601518255948401946001909101908401620001d8565b5085821015620002185787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b608051611e2562000259600039600081816102e30152818161060101528181610d4d01526110220152611e256000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c80638da5cb5b116100ad578063d0a1414a11610071578063d0a1414a1461029d578063d960d573146102b0578063e1081bb3146102c3578063ea1591bb146102cb578063fc0c546a146102de57600080fd5b80638da5cb5b146102235780639348cef714610236578063b14c63c514610249578063c290d69114610277578063c84d2f6a1461028a57600080fd5b8063451df52e116100f4578063451df52e1461019657806355f78c8d146101d75780635a4463a8146101ea5780637efbf8ac146101fd5780638d10b0a61461021057600080fd5b8063176321e9146101265780632e1a7d4d1461013b57806342d21ef71461016357806344a770bf14610183575b600080fd5b6101396101343660046116c6565b610305565b005b61014e61014936600461170d565b610502565b60405190151581526020015b60405180910390f35b61017661017136600461170d565b6106e0565b60405161015a919061176c565b61017661019136600461170d565b61077a565b6101bf6101a436600461170d565b6004602052600090815260409020546001600160a01b031681565b6040516001600160a01b03909116815260200161015a565b6101396101e536600461170d565b610793565b6101396101f8366004611786565b6108ab565b61013961020b3660046117f3565b61099c565b61017661021e36600461170d565b610bab565b6000546101bf906001600160a01b031681565b610139610244366004611839565b610bc4565b61026961025736600461170d565b60056020526000908152604090205481565b60405190815260200161015a565b61013961028536600461170d565b610ecc565b61013961029836600461186c565b6110ec565b6101396102ab3660046116c6565b6112ba565b6102696102be36600461170d565b61147c565b610176611540565b6101396102d9366004611839565b61154d565b6101bf7f000000000000000000000000000000000000000000000000000000000000000081565b60405165656e64696e6760d01b602082015260260160408051601f19818403018152828252805160209182012060008681526006835292909220919261034c9291016118d6565b60405160208183030381529060405280519060200120146103b45760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e472073746174757300000060448201526064015b60405180910390fd5b6000828152600460205260409020546001600160a01b031633146103ea5760405162461bcd60e51b81526004016103ab9061194c565b60405180604001604052806007815260200166636c6f73696e6760c81b81525060066000848152602001908152602001600020908161042991906119cb565b506000828152600460209081526040808320546005835281842054600790935281842091517f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c79461048e9488946001600160a01b039094169390929091908890611b08565b60405180910390a160008281526005602090815260408083205460048352818420546001600160a01b03168452600290925282208054919290916104d3908490611b72565b909155505050600090815260046020908152604080832080546001600160a01b03191690556005909152812055565b6040516337b832b760e11b602082015260009060240160408051601f19818403018152828252805160209182012060008681526006835292909220919261054a9291016118d6565b60405160208183030381529060405280519060200120036105ad5760405162461bcd60e51b815260206004820152601760248201527f436f6e747261637420696e204f50454e2073746174757300000000000000000060448201526064016103ab565b33600090815260026020526040902054801561068b573360008181526002602052604080822091909155516323b872dd60e01b81523060048201526024810191909152604481018290526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303816000875af115801561064a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066e9190611b8b565b61068b573360009081526002602052604081209190915592915050565b6000838152600760205260409081902090517f9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c916106cf9186919033908690611ba8565b60405180910390a150600192915050565b600660205260009081526040902080546106f99061189c565b80601f01602080910402602001604051908101604052809291908181526020018280546107259061189c565b80156107725780601f1061074757610100808354040283529160200191610772565b820191906000526020600020905b81548152906001019060200180831161075557829003601f168201915b505050505081565b600860205260009081526040902080546106f99061189c565b6000546001600160a01b031633146107bd5760405162461bcd60e51b81526004016103ab90611bdd565b6040516337b832b760e11b602082015260240160408051601f1981840301815282825280516020918201206000858152600683529290922091926108029291016118d6565b60405160208183030381529060405280519060200120146108355760405162461bcd60e51b81526004016103ab90611c24565b604051806040016040528060068152602001651c995d99585b60d21b81525060066000838152602001908152602001600020908161087391906119cb565b506040518181527f3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba906020015b60405180910390a150565b6000546001600160a01b0316331461090f5760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79206f776e65722063616e20637265617465206e65772061756374696f6044820152603760f91b60648201526084016103ab565b600083815260046020818152604080842080546001600160a01b031916905560058252808420849055805180820182529283526337b832b760e11b8383015286845260069091529091209061096490826119cb565b50600083815260076020526040902061097d83826119cb565b50600083815260086020526040902061099682826119cb565b50505050565b60405166636c6f73696e6760c81b602082015260270160408051601f1981840301815282825280516020918201206000878152600683529290922091926109e49291016118d6565b6040516020818303038152906040528051906020012014610a475760405162461bcd60e51b815260206004820152601e60248201527f436f6e7472616374206e6f7420696e20434c4f53494e4720737461747573000060448201526064016103ab565b6000838152600460205260409020546001600160a01b03163314610a7d5760405162461bcd60e51b81526004016103ab9061194c565b600083815260086020526040908190209051600a91610a9b916118d6565b9081526040805191829003602090810183208054600181018255600091825282822001869055868152600890915220600991610ad791906118d6565b908152604051602091819003820190208054600181018255600091825291902001610b0282826119cb565b506040518060400160405280600681526020016518db1bdcd95960d21b815250600660008581526020019081526020016000209081610b4191906119cb565b50600083815260046020908152604080832080546001600160a01b031916905560079091529081902090517fcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa791610b9e9186919086908690611c5b565b60405180910390a1505050565b600760205260009081526040902080546106f99061189c565b604051651c995d99585b60d21b602082015260260160408051601f198184030181528282528051602091820120600086815260068352929092209192610c0b9291016118d6565b6040516020818303038152906040528051906020012014610c6e5760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e2052455645414c2073746174757300000060448201526064016103ab565b6000828152600560205260409020548111610ce35760008281526007602090815260408083206005909252918290205491517ff0d11e9dd98376d7e17bc7e9a556f6fba73d4ffc5e409408d71c1288788cf90392610cd792869290913391879190600190611c97565b60405180910390a15050565b600082815260036020908152604080832033845282529182902054825191820184905291016040516020818303038152906040528051906020012014610d2857600080fd5b6040516323b872dd60e01b8152336004820152306024820152604481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303816000875af1158015610d9e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dc29190611b8b565b905080610e0a5760405162461bcd60e51b81526020600482015260166024820152752a37b5b2b7103a3930b739b332b9103330b4b632b21760511b60448201526064016103ab565b60008381526005602052604090205415610e615760008381526005602090815260408083205460048352818420546001600160a01b0316845260029092528220805491929091610e5b908490611b72565b90915550505b600083815260046020908152604080832080546001600160a01b031916339081179091556005835281842086905560079092529182902091517fb95c1199e4385d3e33ad9cffdc96a2f61491b3798965fd93633c02ff7ade77f092610b9e9287928790600190611cdf565b60405166636c6f73696e6760c81b602082015260270160408051601f198184030181528282528051602091820120600085815260068352929092209192610f149291016118d6565b6040516020818303038152906040528051906020012014610f775760405162461bcd60e51b815260206004820152601e60248201527f436f6e7472616374206e6f7420696e20436c6f73696e6720737461747573000060448201526064016103ab565b6000546001600160a01b03163314610fdd5760405162461bcd60e51b8152602060048201526024808201527f4f6e6c79206f776e65722063616e206275726e206269646465722773207061796044820152631b595b9d60e21b60648201526084016103ab565b600081815260056020526040902054156110a8576000818152600560205260408082205490516323b872dd60e01b8152306004820152602481019290925260448201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303816000875af1158015611073573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110979190611b8b565b506000818152600560205260408120555b600081815260056020908152604091829020548251848152918201527f1369343ebae961e006afa581954579f5715145c2db730b9dd4dae66302ec017491016108a0565b6000546001600160a01b031633146111165760405162461bcd60e51b81526004016103ab90611bdd565b604051651c995d99585b60d21b602082015260260160408051601f19818403018152828252805160209182012060008681526006835292909220919261115d9291016118d6565b60405160208183030381529060405280519060200120146111905760405162461bcd60e51b81526004016103ab90611c24565b60405180604001604052806006815260200165656e64696e6760d01b8152506006600084815260200190815260200160002090816111ce91906119cb565b5080806111e75750600082815260056020526040902054155b15611268576040518060400160405280600681526020016518db1bdcd95960d21b81525060066000848152602001908152602001600020908161122a91906119cb565b5060008281526005602090815260408083205460048352818420546001600160a01b03168452600290925282208054919290916104d3908490611b72565b6000828152600460209081526040918290205482518581526001600160a01b03909116918101919091527fa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b48474685002739101610cd7565b60405165656e64696e6760d01b602082015260260160408051601f1981840301815282825280516020918201206000868152600683529290922091926113019291016118d6565b60405160208183030381529060405280519060200120146113645760405162461bcd60e51b815260206004820152601d60248201527f436f6e7472616374206e6f7420696e20454e44494e472073746174757300000060448201526064016103ab565b6000828152600460205260409020546001600160a01b0316331461139a5760405162461bcd60e51b81526004016103ab9061194c565b60405180604001604052806007815260200166636c6f73696e6760c81b8152506006600084815260200190815260200160002090816113d991906119cb565b5060008281526005602090815260408083205483546001600160a01b0316845260029092528220805491929091611411908490611b72565b9091555050600082815260046020908152604080832054600583528184205460079093529281902090517f70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c793610cd79387936001600160a01b03909216929091906001908890611b08565b600081815260086020526040808220905182918291600a9161149d916118d6565b90815260405190819003602001902054905060005b8181101561152157600085815260086020526040908190209051600a916114d8916118d6565b908152602001604051809103902081815481106114f7576114f7611d2c565b90600052602060002001548361150d9190611d42565b92508061151981611d6a565b9150506114b2565b508061152e836064611d83565b6115389190611db3565b949350505050565b600180546106f99061189c565b6040516337b832b760e11b602082015260240160408051601f1981840301815282825280516020918201206000868152600683529290922091926115929291016118d6565b60405160208183030381529060405280519060200120146115c55760405162461bcd60e51b81526004016103ab90611c24565b60008281526003602090815260408083203380855290835281842085905585845260079092529182902091517f313cc01ec1ef841e136089d8dc2dcaa79e223469f4c1634e325d5dfada14848492610cd79286928690600190611cdf565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261164a57600080fd5b813567ffffffffffffffff8082111561166557611665611623565b604051601f8301601f19908116603f0116810190828211818310171561168d5761168d611623565b816040528381528660208588010111156116a657600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080604083850312156116d957600080fd5b82359150602083013567ffffffffffffffff8111156116f757600080fd5b61170385828601611639565b9150509250929050565b60006020828403121561171f57600080fd5b5035919050565b6000815180845260005b8181101561174c57602081850181015186830182015201611730565b506000602082860101526020601f19601f83011685010191505092915050565b60208152600061177f6020830184611726565b9392505050565b60008060006060848603121561179b57600080fd5b83359250602084013567ffffffffffffffff808211156117ba57600080fd5b6117c687838801611639565b935060408601359150808211156117dc57600080fd5b506117e986828701611639565b9150509250925092565b60008060006060848603121561180857600080fd5b8335925060208401359150604084013567ffffffffffffffff81111561182d57600080fd5b6117e986828701611639565b6000806040838503121561184c57600080fd5b50508035926020909101359150565b801515811461186957600080fd5b50565b6000806040838503121561187f57600080fd5b8235915060208301356118918161185b565b809150509250929050565b600181811c908216806118b057607f821691505b6020821081036118d057634e487b7160e01b600052602260045260246000fd5b50919050565b60008083546118e48161189c565b600182811680156118fc576001811461191157611940565b60ff1984168752821515830287019450611940565b8760005260208060002060005b858110156119375781548a82015290840190820161191e565b50505082870194505b50929695505050505050565b6020808252601690820152754e6f7420617574686f72697a6564206163636573732160501b604082015260600190565b601f8211156119c657600081815260208120601f850160051c810160208610156119a35750805b601f850160051c820191505b818110156119c2578281556001016119af565b5050505b505050565b815167ffffffffffffffff8111156119e5576119e5611623565b6119f9816119f3845461189c565b8461197c565b602080601f831160018114611a2e5760008415611a165750858301515b600019600386901b1c1916600185901b1785556119c2565b600085815260208120601f198616915b82811015611a5d57888601518255948401946001909101908401611a3e565b5085821015611a7b5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60008154611a988161189c565b808552602060018381168015611ab55760018114611acf57611afd565b60ff1985168884015283151560051b880183019550611afd565b866000528260002060005b85811015611af55781548a8201860152908301908401611ada565b890184019650505b505050505092915050565b86815260018060a01b038616602082015284604082015260c060608201526000611b3560c0830186611a8b565b841515608084015282810360a0840152611b4f8185611726565b9998505050505050505050565b634e487b7160e01b600052601160045260246000fd5b80820180821115611b8557611b85611b5c565b92915050565b600060208284031215611b9d57600080fd5b815161177f8161185b565b848152608060208201526000611bc16080830186611a8b565b6001600160a01b03949094166040830152506060015292915050565b60208082526027908201527f4f6e6c79206f776e65722063616e206368616e676520636f6e747261637427736040820152662073746174757360c81b606082015260800190565b6020808252601b908201527f436f6e7472616374206e6f7420696e204f50454e207374617475730000000000604082015260600190565b848152608060208201526000611c746080830186611a8b565b8460408401528281036060840152611c8c8185611726565b979650505050505050565b86815260c060208201526000611cb060c0830188611a8b565b6001600160a01b0387166040840152606083018690526080830185905282810360a0840152611b4f8185611a8b565b85815260a060208201526000611cf860a0830187611a8b565b6001600160a01b0386166040840152606083018590528281036080840152611d208185611a8b565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b8082018281126000831280158216821582161715611d6257611d62611b5c565b505092915050565b600060018201611d7c57611d7c611b5c565b5060010190565b80820260008212600160ff1b84141615611d9f57611d9f611b5c565b8181058314821517611b8557611b85611b5c565b600082611dd057634e487b7160e01b600052601260045260246000fd5b600160ff1b821460001984141615611dea57611dea611b5c565b50059056fea26469706673582212206ef0f0fbf5404eb41803571492dcd60eb66554786533d8c507f9f66baf79a34c64736f6c63430008120033",
}

// Cb1pAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use Cb1pAuctionMetaData.ABI instead.
var Cb1pAuctionABI = Cb1pAuctionMetaData.ABI

// Cb1pAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Cb1pAuctionMetaData.Bin instead.
var Cb1pAuctionBin = Cb1pAuctionMetaData.Bin

// DeployCb1pAuction deploys a new Ethereum contract, binding an instance of Cb1pAuction to it.
func DeployCb1pAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *Cb1pAuction, error) {
	parsed, err := Cb1pAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Cb1pAuctionBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cb1pAuction{Cb1pAuctionCaller: Cb1pAuctionCaller{contract: contract}, Cb1pAuctionTransactor: Cb1pAuctionTransactor{contract: contract}, Cb1pAuctionFilterer: Cb1pAuctionFilterer{contract: contract}}, nil
}

// Cb1pAuction is an auto generated Go binding around an Ethereum contract.
type Cb1pAuction struct {
	Cb1pAuctionCaller     // Read-only binding to the contract
	Cb1pAuctionTransactor // Write-only binding to the contract
	Cb1pAuctionFilterer   // Log filterer for contract events
}

// Cb1pAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type Cb1pAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cb1pAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Cb1pAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cb1pAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Cb1pAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cb1pAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Cb1pAuctionSession struct {
	Contract     *Cb1pAuction      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Cb1pAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Cb1pAuctionCallerSession struct {
	Contract *Cb1pAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// Cb1pAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Cb1pAuctionTransactorSession struct {
	Contract     *Cb1pAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Cb1pAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type Cb1pAuctionRaw struct {
	Contract *Cb1pAuction // Generic contract binding to access the raw methods on
}

// Cb1pAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Cb1pAuctionCallerRaw struct {
	Contract *Cb1pAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// Cb1pAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Cb1pAuctionTransactorRaw struct {
	Contract *Cb1pAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCb1pAuction creates a new instance of Cb1pAuction, bound to a specific deployed contract.
func NewCb1pAuction(address common.Address, backend bind.ContractBackend) (*Cb1pAuction, error) {
	contract, err := bindCb1pAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cb1pAuction{Cb1pAuctionCaller: Cb1pAuctionCaller{contract: contract}, Cb1pAuctionTransactor: Cb1pAuctionTransactor{contract: contract}, Cb1pAuctionFilterer: Cb1pAuctionFilterer{contract: contract}}, nil
}

// NewCb1pAuctionCaller creates a new read-only instance of Cb1pAuction, bound to a specific deployed contract.
func NewCb1pAuctionCaller(address common.Address, caller bind.ContractCaller) (*Cb1pAuctionCaller, error) {
	contract, err := bindCb1pAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionCaller{contract: contract}, nil
}

// NewCb1pAuctionTransactor creates a new write-only instance of Cb1pAuction, bound to a specific deployed contract.
func NewCb1pAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*Cb1pAuctionTransactor, error) {
	contract, err := bindCb1pAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionTransactor{contract: contract}, nil
}

// NewCb1pAuctionFilterer creates a new log filterer instance of Cb1pAuction, bound to a specific deployed contract.
func NewCb1pAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*Cb1pAuctionFilterer, error) {
	contract, err := bindCb1pAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionFilterer{contract: contract}, nil
}

// bindCb1pAuction binds a generic wrapper to an already deployed contract.
func bindCb1pAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Cb1pAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cb1pAuction *Cb1pAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cb1pAuction.Contract.Cb1pAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cb1pAuction *Cb1pAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Cb1pAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cb1pAuction *Cb1pAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Cb1pAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cb1pAuction *Cb1pAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cb1pAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cb1pAuction *Cb1pAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cb1pAuction *Cb1pAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.contract.Transact(opts, method, params...)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_Cb1pAuction *Cb1pAuctionCaller) AssetId(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Cb1pAuction.contract.Call(opts, &out, "asset_id", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_Cb1pAuction *Cb1pAuctionSession) AssetId(arg0 *big.Int) (string, error) {
	return _Cb1pAuction.Contract.AssetId(&_Cb1pAuction.CallOpts, arg0)
}

// AssetId is a free data retrieval call binding the contract method 0x8d10b0a6.
//
// Solidity: function asset_id(uint256 ) view returns(string)
func (_Cb1pAuction *Cb1pAuctionCallerSession) AssetId(arg0 *big.Int) (string, error) {
	return _Cb1pAuction.Contract.AssetId(&_Cb1pAuction.CallOpts, arg0)
}

// AssetOwner is a free data retrieval call binding the contract method 0x44a770bf.
//
// Solidity: function asset_owner(uint256 ) view returns(string)
func (_Cb1pAuction *Cb1pAuctionCaller) AssetOwner(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Cb1pAuction.contract.Call(opts, &out, "asset_owner", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AssetOwner is a free data retrieval call binding the contract method 0x44a770bf.
//
// Solidity: function asset_owner(uint256 ) view returns(string)
func (_Cb1pAuction *Cb1pAuctionSession) AssetOwner(arg0 *big.Int) (string, error) {
	return _Cb1pAuction.Contract.AssetOwner(&_Cb1pAuction.CallOpts, arg0)
}

// AssetOwner is a free data retrieval call binding the contract method 0x44a770bf.
//
// Solidity: function asset_owner(uint256 ) view returns(string)
func (_Cb1pAuction *Cb1pAuctionCallerSession) AssetOwner(arg0 *big.Int) (string, error) {
	return _Cb1pAuction.Contract.AssetOwner(&_Cb1pAuction.CallOpts, arg0)
}

// AuctionType is a free data retrieval call binding the contract method 0xe1081bb3.
//
// Solidity: function auction_type() view returns(string)
func (_Cb1pAuction *Cb1pAuctionCaller) AuctionType(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cb1pAuction.contract.Call(opts, &out, "auction_type")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AuctionType is a free data retrieval call binding the contract method 0xe1081bb3.
//
// Solidity: function auction_type() view returns(string)
func (_Cb1pAuction *Cb1pAuctionSession) AuctionType() (string, error) {
	return _Cb1pAuction.Contract.AuctionType(&_Cb1pAuction.CallOpts)
}

// AuctionType is a free data retrieval call binding the contract method 0xe1081bb3.
//
// Solidity: function auction_type() view returns(string)
func (_Cb1pAuction *Cb1pAuctionCallerSession) AuctionType() (string, error) {
	return _Cb1pAuction.Contract.AuctionType(&_Cb1pAuction.CallOpts)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xd960d573.
//
// Solidity: function checkAverageScore(uint256 auctionId) view returns(int256)
func (_Cb1pAuction *Cb1pAuctionCaller) CheckAverageScore(opts *bind.CallOpts, auctionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cb1pAuction.contract.Call(opts, &out, "checkAverageScore", auctionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckAverageScore is a free data retrieval call binding the contract method 0xd960d573.
//
// Solidity: function checkAverageScore(uint256 auctionId) view returns(int256)
func (_Cb1pAuction *Cb1pAuctionSession) CheckAverageScore(auctionId *big.Int) (*big.Int, error) {
	return _Cb1pAuction.Contract.CheckAverageScore(&_Cb1pAuction.CallOpts, auctionId)
}

// CheckAverageScore is a free data retrieval call binding the contract method 0xd960d573.
//
// Solidity: function checkAverageScore(uint256 auctionId) view returns(int256)
func (_Cb1pAuction *Cb1pAuctionCallerSession) CheckAverageScore(auctionId *big.Int) (*big.Int, error) {
	return _Cb1pAuction.Contract.CheckAverageScore(&_Cb1pAuction.CallOpts, auctionId)
}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_Cb1pAuction *Cb1pAuctionCaller) HighestBid(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cb1pAuction.contract.Call(opts, &out, "highestBid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_Cb1pAuction *Cb1pAuctionSession) HighestBid(arg0 *big.Int) (*big.Int, error) {
	return _Cb1pAuction.Contract.HighestBid(&_Cb1pAuction.CallOpts, arg0)
}

// HighestBid is a free data retrieval call binding the contract method 0xb14c63c5.
//
// Solidity: function highestBid(uint256 ) view returns(uint256)
func (_Cb1pAuction *Cb1pAuctionCallerSession) HighestBid(arg0 *big.Int) (*big.Int, error) {
	return _Cb1pAuction.Contract.HighestBid(&_Cb1pAuction.CallOpts, arg0)
}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_Cb1pAuction *Cb1pAuctionCaller) HighestBidder(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Cb1pAuction.contract.Call(opts, &out, "highestBidder", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_Cb1pAuction *Cb1pAuctionSession) HighestBidder(arg0 *big.Int) (common.Address, error) {
	return _Cb1pAuction.Contract.HighestBidder(&_Cb1pAuction.CallOpts, arg0)
}

// HighestBidder is a free data retrieval call binding the contract method 0x451df52e.
//
// Solidity: function highestBidder(uint256 ) view returns(address)
func (_Cb1pAuction *Cb1pAuctionCallerSession) HighestBidder(arg0 *big.Int) (common.Address, error) {
	return _Cb1pAuction.Contract.HighestBidder(&_Cb1pAuction.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cb1pAuction *Cb1pAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cb1pAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cb1pAuction *Cb1pAuctionSession) Owner() (common.Address, error) {
	return _Cb1pAuction.Contract.Owner(&_Cb1pAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cb1pAuction *Cb1pAuctionCallerSession) Owner() (common.Address, error) {
	return _Cb1pAuction.Contract.Owner(&_Cb1pAuction.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_Cb1pAuction *Cb1pAuctionCaller) Status(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Cb1pAuction.contract.Call(opts, &out, "status", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_Cb1pAuction *Cb1pAuctionSession) Status(arg0 *big.Int) (string, error) {
	return _Cb1pAuction.Contract.Status(&_Cb1pAuction.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(string)
func (_Cb1pAuction *Cb1pAuctionCallerSession) Status(arg0 *big.Int) (string, error) {
	return _Cb1pAuction.Contract.Status(&_Cb1pAuction.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Cb1pAuction *Cb1pAuctionCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cb1pAuction.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Cb1pAuction *Cb1pAuctionSession) Token() (common.Address, error) {
	return _Cb1pAuction.Contract.Token(&_Cb1pAuction.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Cb1pAuction *Cb1pAuctionCallerSession) Token() (common.Address, error) {
	return _Cb1pAuction.Contract.Token(&_Cb1pAuction.CallOpts)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_Cb1pAuction *Cb1pAuctionTransactor) Abort(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "abort", auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_Cb1pAuction *Cb1pAuctionSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Abort(&_Cb1pAuction.TransactOpts, auctionId, jsonString)
}

// Abort is a paid mutator transaction binding the contract method 0x176321e9.
//
// Solidity: function abort(uint256 auctionId, string jsonString) returns()
func (_Cb1pAuction *Cb1pAuctionTransactorSession) Abort(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Abort(&_Cb1pAuction.TransactOpts, auctionId, jsonString)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_Cb1pAuction *Cb1pAuctionTransactor) Bid(opts *bind.TransactOpts, auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "bid", auctionId, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_Cb1pAuction *Cb1pAuctionSession) Bid(auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Bid(&_Cb1pAuction.TransactOpts, auctionId, bidHash)
}

// Bid is a paid mutator transaction binding the contract method 0xea1591bb.
//
// Solidity: function bid(uint256 auctionId, bytes32 bidHash) returns()
func (_Cb1pAuction *Cb1pAuctionTransactorSession) Bid(auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Bid(&_Cb1pAuction.TransactOpts, auctionId, bidHash)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_Cb1pAuction *Cb1pAuctionTransactor) CloseAuction(opts *bind.TransactOpts, auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "closeAuction", auctionId, not_winner_platform)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_Cb1pAuction *Cb1pAuctionSession) CloseAuction(auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.CloseAuction(&_Cb1pAuction.TransactOpts, auctionId, not_winner_platform)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xc84d2f6a.
//
// Solidity: function closeAuction(uint256 auctionId, bool not_winner_platform) returns()
func (_Cb1pAuction *Cb1pAuctionTransactorSession) CloseAuction(auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.CloseAuction(&_Cb1pAuction.TransactOpts, auctionId, not_winner_platform)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_Cb1pAuction *Cb1pAuctionTransactor) Commit(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "commit", auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_Cb1pAuction *Cb1pAuctionSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Commit(&_Cb1pAuction.TransactOpts, auctionId, jsonString)
}

// Commit is a paid mutator transaction binding the contract method 0xd0a1414a.
//
// Solidity: function commit(uint256 auctionId, string jsonString) returns()
func (_Cb1pAuction *Cb1pAuctionTransactorSession) Commit(auctionId *big.Int, jsonString string) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Commit(&_Cb1pAuction.TransactOpts, auctionId, jsonString)
}

// Create is a paid mutator transaction binding the contract method 0x5a4463a8.
//
// Solidity: function create(uint256 _auction_id, string _asset_id, string _asset_owner) returns()
func (_Cb1pAuction *Cb1pAuctionTransactor) Create(opts *bind.TransactOpts, _auction_id *big.Int, _asset_id string, _asset_owner string) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "create", _auction_id, _asset_id, _asset_owner)
}

// Create is a paid mutator transaction binding the contract method 0x5a4463a8.
//
// Solidity: function create(uint256 _auction_id, string _asset_id, string _asset_owner) returns()
func (_Cb1pAuction *Cb1pAuctionSession) Create(_auction_id *big.Int, _asset_id string, _asset_owner string) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Create(&_Cb1pAuction.TransactOpts, _auction_id, _asset_id, _asset_owner)
}

// Create is a paid mutator transaction binding the contract method 0x5a4463a8.
//
// Solidity: function create(uint256 _auction_id, string _asset_id, string _asset_owner) returns()
func (_Cb1pAuction *Cb1pAuctionTransactorSession) Create(_auction_id *big.Int, _asset_id string, _asset_owner string) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Create(&_Cb1pAuction.TransactOpts, _auction_id, _asset_id, _asset_owner)
}

// Pay is a paid mutator transaction binding the contract method 0xc290d691.
//
// Solidity: function pay(uint256 auctionId) returns()
func (_Cb1pAuction *Cb1pAuctionTransactor) Pay(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "pay", auctionId)
}

// Pay is a paid mutator transaction binding the contract method 0xc290d691.
//
// Solidity: function pay(uint256 auctionId) returns()
func (_Cb1pAuction *Cb1pAuctionSession) Pay(auctionId *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Pay(&_Cb1pAuction.TransactOpts, auctionId)
}

// Pay is a paid mutator transaction binding the contract method 0xc290d691.
//
// Solidity: function pay(uint256 auctionId) returns()
func (_Cb1pAuction *Cb1pAuctionTransactorSession) Pay(auctionId *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Pay(&_Cb1pAuction.TransactOpts, auctionId)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x7efbf8ac.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, string _feedback) returns()
func (_Cb1pAuction *Cb1pAuctionTransactor) ProvideFeedback(opts *bind.TransactOpts, auctionId *big.Int, _score *big.Int, _feedback string) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "provide_feedback", auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x7efbf8ac.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, string _feedback) returns()
func (_Cb1pAuction *Cb1pAuctionSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback string) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.ProvideFeedback(&_Cb1pAuction.TransactOpts, auctionId, _score, _feedback)
}

// ProvideFeedback is a paid mutator transaction binding the contract method 0x7efbf8ac.
//
// Solidity: function provide_feedback(uint256 auctionId, int256 _score, string _feedback) returns()
func (_Cb1pAuction *Cb1pAuctionTransactorSession) ProvideFeedback(auctionId *big.Int, _score *big.Int, _feedback string) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.ProvideFeedback(&_Cb1pAuction.TransactOpts, auctionId, _score, _feedback)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_Cb1pAuction *Cb1pAuctionTransactor) Reveal(opts *bind.TransactOpts, auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "reveal", auctionId, bidAmount)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_Cb1pAuction *Cb1pAuctionSession) Reveal(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Reveal(&_Cb1pAuction.TransactOpts, auctionId, bidAmount)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 auctionId, uint256 bidAmount) returns()
func (_Cb1pAuction *Cb1pAuctionTransactorSession) Reveal(auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Reveal(&_Cb1pAuction.TransactOpts, auctionId, bidAmount)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_Cb1pAuction *Cb1pAuctionTransactor) RevealAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "revealAuction", auctionId)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_Cb1pAuction *Cb1pAuctionSession) RevealAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.RevealAuction(&_Cb1pAuction.TransactOpts, auctionId)
}

// RevealAuction is a paid mutator transaction binding the contract method 0x55f78c8d.
//
// Solidity: function revealAuction(uint256 auctionId) returns()
func (_Cb1pAuction *Cb1pAuctionTransactorSession) RevealAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.RevealAuction(&_Cb1pAuction.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_Cb1pAuction *Cb1pAuctionTransactor) Withdraw(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.contract.Transact(opts, "withdraw", auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_Cb1pAuction *Cb1pAuctionSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Withdraw(&_Cb1pAuction.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 auctionId) returns(bool)
func (_Cb1pAuction *Cb1pAuctionTransactorSession) Withdraw(auctionId *big.Int) (*types.Transaction, error) {
	return _Cb1pAuction.Contract.Withdraw(&_Cb1pAuction.TransactOpts, auctionId)
}

// Cb1pAuctionAwaitResponseIterator is returned from FilterAwaitResponse and is used to iterate over the raw logs and unpacked data for AwaitResponse events raised by the Cb1pAuction contract.
type Cb1pAuctionAwaitResponseIterator struct {
	Event *Cb1pAuctionAwaitResponse // Event containing the contract specifics and raw log

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
func (it *Cb1pAuctionAwaitResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb1pAuctionAwaitResponse)
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
		it.Event = new(Cb1pAuctionAwaitResponse)
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
func (it *Cb1pAuctionAwaitResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb1pAuctionAwaitResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb1pAuctionAwaitResponse represents a AwaitResponse event raised by the Cb1pAuction contract.
type Cb1pAuctionAwaitResponse struct {
	AuctionId *big.Int
	Winner    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAwaitResponse is a free log retrieval operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auctionId, address winner)
func (_Cb1pAuction *Cb1pAuctionFilterer) FilterAwaitResponse(opts *bind.FilterOpts) (*Cb1pAuctionAwaitResponseIterator, error) {

	logs, sub, err := _Cb1pAuction.contract.FilterLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionAwaitResponseIterator{contract: _Cb1pAuction.contract, event: "AwaitResponse", logs: logs, sub: sub}, nil
}

// WatchAwaitResponse is a free log subscription operation binding the contract event 0xa4b690e89a49a57d32303c0c79679bacafe2c7f7b95dba3338b4847468500273.
//
// Solidity: event AwaitResponse(uint256 auctionId, address winner)
func (_Cb1pAuction *Cb1pAuctionFilterer) WatchAwaitResponse(opts *bind.WatchOpts, sink chan<- *Cb1pAuctionAwaitResponse) (event.Subscription, error) {

	logs, sub, err := _Cb1pAuction.contract.WatchLogs(opts, "AwaitResponse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb1pAuctionAwaitResponse)
				if err := _Cb1pAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
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
// Solidity: event AwaitResponse(uint256 auctionId, address winner)
func (_Cb1pAuction *Cb1pAuctionFilterer) ParseAwaitResponse(log types.Log) (*Cb1pAuctionAwaitResponse, error) {
	event := new(Cb1pAuctionAwaitResponse)
	if err := _Cb1pAuction.contract.UnpackLog(event, "AwaitResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb1pAuctionBidTooLowIterator is returned from FilterBidTooLow and is used to iterate over the raw logs and unpacked data for BidTooLow events raised by the Cb1pAuction contract.
type Cb1pAuctionBidTooLowIterator struct {
	Event *Cb1pAuctionBidTooLow // Event containing the contract specifics and raw log

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
func (it *Cb1pAuctionBidTooLowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb1pAuctionBidTooLow)
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
		it.Event = new(Cb1pAuctionBidTooLow)
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
func (it *Cb1pAuctionBidTooLowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb1pAuctionBidTooLowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb1pAuctionBidTooLow represents a BidTooLow event raised by the Cb1pAuction contract.
type Cb1pAuctionBidTooLow struct {
	AuctionId   *big.Int
	Id          string
	Bidder      common.Address
	BidAmount   *big.Int
	HighestBid  *big.Int
	AuctionType string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBidTooLow is a free log retrieval operation binding the contract event 0xf0d11e9dd98376d7e17bc7e9a556f6fba73d4ffc5e409408d71c1288788cf903.
//
// Solidity: event BidTooLow(uint256 auctionId, string id, address bidder, uint256 bidAmount, uint256 highestBid, string auctionType)
func (_Cb1pAuction *Cb1pAuctionFilterer) FilterBidTooLow(opts *bind.FilterOpts) (*Cb1pAuctionBidTooLowIterator, error) {

	logs, sub, err := _Cb1pAuction.contract.FilterLogs(opts, "BidTooLow")
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionBidTooLowIterator{contract: _Cb1pAuction.contract, event: "BidTooLow", logs: logs, sub: sub}, nil
}

// WatchBidTooLow is a free log subscription operation binding the contract event 0xf0d11e9dd98376d7e17bc7e9a556f6fba73d4ffc5e409408d71c1288788cf903.
//
// Solidity: event BidTooLow(uint256 auctionId, string id, address bidder, uint256 bidAmount, uint256 highestBid, string auctionType)
func (_Cb1pAuction *Cb1pAuctionFilterer) WatchBidTooLow(opts *bind.WatchOpts, sink chan<- *Cb1pAuctionBidTooLow) (event.Subscription, error) {

	logs, sub, err := _Cb1pAuction.contract.WatchLogs(opts, "BidTooLow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb1pAuctionBidTooLow)
				if err := _Cb1pAuction.contract.UnpackLog(event, "BidTooLow", log); err != nil {
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

// ParseBidTooLow is a log parse operation binding the contract event 0xf0d11e9dd98376d7e17bc7e9a556f6fba73d4ffc5e409408d71c1288788cf903.
//
// Solidity: event BidTooLow(uint256 auctionId, string id, address bidder, uint256 bidAmount, uint256 highestBid, string auctionType)
func (_Cb1pAuction *Cb1pAuctionFilterer) ParseBidTooLow(log types.Log) (*Cb1pAuctionBidTooLow, error) {
	event := new(Cb1pAuctionBidTooLow)
	if err := _Cb1pAuction.contract.UnpackLog(event, "BidTooLow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb1pAuctionDecisionMadeIterator is returned from FilterDecisionMade and is used to iterate over the raw logs and unpacked data for DecisionMade events raised by the Cb1pAuction contract.
type Cb1pAuctionDecisionMadeIterator struct {
	Event *Cb1pAuctionDecisionMade // Event containing the contract specifics and raw log

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
func (it *Cb1pAuctionDecisionMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb1pAuctionDecisionMade)
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
		it.Event = new(Cb1pAuctionDecisionMade)
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
func (it *Cb1pAuctionDecisionMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb1pAuctionDecisionMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb1pAuctionDecisionMade represents a DecisionMade event raised by the Cb1pAuction contract.
type Cb1pAuctionDecisionMade struct {
	AuctionId  *big.Int
	Winner     common.Address
	Amount     *big.Int
	Id         string
	Prcd       bool
	JsonString string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDecisionMade is a free log retrieval operation binding the contract event 0x70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7.
//
// Solidity: event DecisionMade(uint256 auctionId, address winner, uint256 amount, string id, bool prcd, string jsonString)
func (_Cb1pAuction *Cb1pAuctionFilterer) FilterDecisionMade(opts *bind.FilterOpts) (*Cb1pAuctionDecisionMadeIterator, error) {

	logs, sub, err := _Cb1pAuction.contract.FilterLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionDecisionMadeIterator{contract: _Cb1pAuction.contract, event: "DecisionMade", logs: logs, sub: sub}, nil
}

// WatchDecisionMade is a free log subscription operation binding the contract event 0x70ceb0775fa7938ed493a232f0be6acf5986ac714e9ced9ce323de68972df9c7.
//
// Solidity: event DecisionMade(uint256 auctionId, address winner, uint256 amount, string id, bool prcd, string jsonString)
func (_Cb1pAuction *Cb1pAuctionFilterer) WatchDecisionMade(opts *bind.WatchOpts, sink chan<- *Cb1pAuctionDecisionMade) (event.Subscription, error) {

	logs, sub, err := _Cb1pAuction.contract.WatchLogs(opts, "DecisionMade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb1pAuctionDecisionMade)
				if err := _Cb1pAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
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
// Solidity: event DecisionMade(uint256 auctionId, address winner, uint256 amount, string id, bool prcd, string jsonString)
func (_Cb1pAuction *Cb1pAuctionFilterer) ParseDecisionMade(log types.Log) (*Cb1pAuctionDecisionMade, error) {
	event := new(Cb1pAuctionDecisionMade)
	if err := _Cb1pAuction.contract.UnpackLog(event, "DecisionMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb1pAuctionHighestBidIncreasedIterator is returned from FilterHighestBidIncreased and is used to iterate over the raw logs and unpacked data for HighestBidIncreased events raised by the Cb1pAuction contract.
type Cb1pAuctionHighestBidIncreasedIterator struct {
	Event *Cb1pAuctionHighestBidIncreased // Event containing the contract specifics and raw log

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
func (it *Cb1pAuctionHighestBidIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb1pAuctionHighestBidIncreased)
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
		it.Event = new(Cb1pAuctionHighestBidIncreased)
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
func (it *Cb1pAuctionHighestBidIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb1pAuctionHighestBidIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb1pAuctionHighestBidIncreased represents a HighestBidIncreased event raised by the Cb1pAuction contract.
type Cb1pAuctionHighestBidIncreased struct {
	AuctionId   *big.Int
	Id          string
	Bidder      common.Address
	BidAmount   *big.Int
	AuctionType string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterHighestBidIncreased is a free log retrieval operation binding the contract event 0xb95c1199e4385d3e33ad9cffdc96a2f61491b3798965fd93633c02ff7ade77f0.
//
// Solidity: event HighestBidIncreased(uint256 auctionId, string id, address bidder, uint256 bidAmount, string auctionType)
func (_Cb1pAuction *Cb1pAuctionFilterer) FilterHighestBidIncreased(opts *bind.FilterOpts) (*Cb1pAuctionHighestBidIncreasedIterator, error) {

	logs, sub, err := _Cb1pAuction.contract.FilterLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionHighestBidIncreasedIterator{contract: _Cb1pAuction.contract, event: "HighestBidIncreased", logs: logs, sub: sub}, nil
}

// WatchHighestBidIncreased is a free log subscription operation binding the contract event 0xb95c1199e4385d3e33ad9cffdc96a2f61491b3798965fd93633c02ff7ade77f0.
//
// Solidity: event HighestBidIncreased(uint256 auctionId, string id, address bidder, uint256 bidAmount, string auctionType)
func (_Cb1pAuction *Cb1pAuctionFilterer) WatchHighestBidIncreased(opts *bind.WatchOpts, sink chan<- *Cb1pAuctionHighestBidIncreased) (event.Subscription, error) {

	logs, sub, err := _Cb1pAuction.contract.WatchLogs(opts, "HighestBidIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb1pAuctionHighestBidIncreased)
				if err := _Cb1pAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
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

// ParseHighestBidIncreased is a log parse operation binding the contract event 0xb95c1199e4385d3e33ad9cffdc96a2f61491b3798965fd93633c02ff7ade77f0.
//
// Solidity: event HighestBidIncreased(uint256 auctionId, string id, address bidder, uint256 bidAmount, string auctionType)
func (_Cb1pAuction *Cb1pAuctionFilterer) ParseHighestBidIncreased(log types.Log) (*Cb1pAuctionHighestBidIncreased, error) {
	event := new(Cb1pAuctionHighestBidIncreased)
	if err := _Cb1pAuction.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb1pAuctionNewBidHashIterator is returned from FilterNewBidHash and is used to iterate over the raw logs and unpacked data for NewBidHash events raised by the Cb1pAuction contract.
type Cb1pAuctionNewBidHashIterator struct {
	Event *Cb1pAuctionNewBidHash // Event containing the contract specifics and raw log

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
func (it *Cb1pAuctionNewBidHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb1pAuctionNewBidHash)
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
		it.Event = new(Cb1pAuctionNewBidHash)
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
func (it *Cb1pAuctionNewBidHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb1pAuctionNewBidHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb1pAuctionNewBidHash represents a NewBidHash event raised by the Cb1pAuction contract.
type Cb1pAuctionNewBidHash struct {
	AuctionId   *big.Int
	Id          string
	Bidder      common.Address
	BidHash     [32]byte
	AuctionType string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewBidHash is a free log retrieval operation binding the contract event 0x313cc01ec1ef841e136089d8dc2dcaa79e223469f4c1634e325d5dfada148484.
//
// Solidity: event NewBidHash(uint256 auctionId, string id, address bidder, bytes32 bidHash, string auctionType)
func (_Cb1pAuction *Cb1pAuctionFilterer) FilterNewBidHash(opts *bind.FilterOpts) (*Cb1pAuctionNewBidHashIterator, error) {

	logs, sub, err := _Cb1pAuction.contract.FilterLogs(opts, "NewBidHash")
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionNewBidHashIterator{contract: _Cb1pAuction.contract, event: "NewBidHash", logs: logs, sub: sub}, nil
}

// WatchNewBidHash is a free log subscription operation binding the contract event 0x313cc01ec1ef841e136089d8dc2dcaa79e223469f4c1634e325d5dfada148484.
//
// Solidity: event NewBidHash(uint256 auctionId, string id, address bidder, bytes32 bidHash, string auctionType)
func (_Cb1pAuction *Cb1pAuctionFilterer) WatchNewBidHash(opts *bind.WatchOpts, sink chan<- *Cb1pAuctionNewBidHash) (event.Subscription, error) {

	logs, sub, err := _Cb1pAuction.contract.WatchLogs(opts, "NewBidHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb1pAuctionNewBidHash)
				if err := _Cb1pAuction.contract.UnpackLog(event, "NewBidHash", log); err != nil {
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

// ParseNewBidHash is a log parse operation binding the contract event 0x313cc01ec1ef841e136089d8dc2dcaa79e223469f4c1634e325d5dfada148484.
//
// Solidity: event NewBidHash(uint256 auctionId, string id, address bidder, bytes32 bidHash, string auctionType)
func (_Cb1pAuction *Cb1pAuctionFilterer) ParseNewBidHash(log types.Log) (*Cb1pAuctionNewBidHash, error) {
	event := new(Cb1pAuctionNewBidHash)
	if err := _Cb1pAuction.contract.UnpackLog(event, "NewBidHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb1pAuctionPayIterator is returned from FilterPay and is used to iterate over the raw logs and unpacked data for Pay events raised by the Cb1pAuction contract.
type Cb1pAuctionPayIterator struct {
	Event *Cb1pAuctionPay // Event containing the contract specifics and raw log

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
func (it *Cb1pAuctionPayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb1pAuctionPay)
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
		it.Event = new(Cb1pAuctionPay)
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
func (it *Cb1pAuctionPayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb1pAuctionPayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb1pAuctionPay represents a Pay event raised by the Cb1pAuction contract.
type Cb1pAuctionPay struct {
	AuctionId *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPay is a free log retrieval operation binding the contract event 0x1369343ebae961e006afa581954579f5715145c2db730b9dd4dae66302ec0174.
//
// Solidity: event Pay(uint256 auctionId, uint256 amount)
func (_Cb1pAuction *Cb1pAuctionFilterer) FilterPay(opts *bind.FilterOpts) (*Cb1pAuctionPayIterator, error) {

	logs, sub, err := _Cb1pAuction.contract.FilterLogs(opts, "Pay")
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionPayIterator{contract: _Cb1pAuction.contract, event: "Pay", logs: logs, sub: sub}, nil
}

// WatchPay is a free log subscription operation binding the contract event 0x1369343ebae961e006afa581954579f5715145c2db730b9dd4dae66302ec0174.
//
// Solidity: event Pay(uint256 auctionId, uint256 amount)
func (_Cb1pAuction *Cb1pAuctionFilterer) WatchPay(opts *bind.WatchOpts, sink chan<- *Cb1pAuctionPay) (event.Subscription, error) {

	logs, sub, err := _Cb1pAuction.contract.WatchLogs(opts, "Pay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb1pAuctionPay)
				if err := _Cb1pAuction.contract.UnpackLog(event, "Pay", log); err != nil {
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

// ParsePay is a log parse operation binding the contract event 0x1369343ebae961e006afa581954579f5715145c2db730b9dd4dae66302ec0174.
//
// Solidity: event Pay(uint256 auctionId, uint256 amount)
func (_Cb1pAuction *Cb1pAuctionFilterer) ParsePay(log types.Log) (*Cb1pAuctionPay, error) {
	event := new(Cb1pAuctionPay)
	if err := _Cb1pAuction.contract.UnpackLog(event, "Pay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb1pAuctionRateAuctionIterator is returned from FilterRateAuction and is used to iterate over the raw logs and unpacked data for RateAuction events raised by the Cb1pAuction contract.
type Cb1pAuctionRateAuctionIterator struct {
	Event *Cb1pAuctionRateAuction // Event containing the contract specifics and raw log

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
func (it *Cb1pAuctionRateAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb1pAuctionRateAuction)
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
		it.Event = new(Cb1pAuctionRateAuction)
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
func (it *Cb1pAuctionRateAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb1pAuctionRateAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb1pAuctionRateAuction represents a RateAuction event raised by the Cb1pAuction contract.
type Cb1pAuctionRateAuction struct {
	AuctionId *big.Int
	Id        string
	Rating    *big.Int
	Review    string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRateAuction is a free log retrieval operation binding the contract event 0xcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa7.
//
// Solidity: event RateAuction(uint256 auctionId, string id, int256 rating, string review)
func (_Cb1pAuction *Cb1pAuctionFilterer) FilterRateAuction(opts *bind.FilterOpts) (*Cb1pAuctionRateAuctionIterator, error) {

	logs, sub, err := _Cb1pAuction.contract.FilterLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionRateAuctionIterator{contract: _Cb1pAuction.contract, event: "RateAuction", logs: logs, sub: sub}, nil
}

// WatchRateAuction is a free log subscription operation binding the contract event 0xcbb0dc5e5b19c111126e2c5f6b96c4cbc2b1fc3ef08c41178bffc7d7136acfa7.
//
// Solidity: event RateAuction(uint256 auctionId, string id, int256 rating, string review)
func (_Cb1pAuction *Cb1pAuctionFilterer) WatchRateAuction(opts *bind.WatchOpts, sink chan<- *Cb1pAuctionRateAuction) (event.Subscription, error) {

	logs, sub, err := _Cb1pAuction.contract.WatchLogs(opts, "RateAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb1pAuctionRateAuction)
				if err := _Cb1pAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
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
func (_Cb1pAuction *Cb1pAuctionFilterer) ParseRateAuction(log types.Log) (*Cb1pAuctionRateAuction, error) {
	event := new(Cb1pAuctionRateAuction)
	if err := _Cb1pAuction.contract.UnpackLog(event, "RateAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb1pAuctionRevealAuctionIterator is returned from FilterRevealAuction and is used to iterate over the raw logs and unpacked data for RevealAuction events raised by the Cb1pAuction contract.
type Cb1pAuctionRevealAuctionIterator struct {
	Event *Cb1pAuctionRevealAuction // Event containing the contract specifics and raw log

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
func (it *Cb1pAuctionRevealAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb1pAuctionRevealAuction)
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
		it.Event = new(Cb1pAuctionRevealAuction)
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
func (it *Cb1pAuctionRevealAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb1pAuctionRevealAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb1pAuctionRevealAuction represents a RevealAuction event raised by the Cb1pAuction contract.
type Cb1pAuctionRevealAuction struct {
	AuctionId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRevealAuction is a free log retrieval operation binding the contract event 0x3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba.
//
// Solidity: event RevealAuction(uint256 auctionId)
func (_Cb1pAuction *Cb1pAuctionFilterer) FilterRevealAuction(opts *bind.FilterOpts) (*Cb1pAuctionRevealAuctionIterator, error) {

	logs, sub, err := _Cb1pAuction.contract.FilterLogs(opts, "RevealAuction")
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionRevealAuctionIterator{contract: _Cb1pAuction.contract, event: "RevealAuction", logs: logs, sub: sub}, nil
}

// WatchRevealAuction is a free log subscription operation binding the contract event 0x3141f569af0f2a74e294bddba6f8391cf40e85348ee9fcf4e79ad4b75a4794ba.
//
// Solidity: event RevealAuction(uint256 auctionId)
func (_Cb1pAuction *Cb1pAuctionFilterer) WatchRevealAuction(opts *bind.WatchOpts, sink chan<- *Cb1pAuctionRevealAuction) (event.Subscription, error) {

	logs, sub, err := _Cb1pAuction.contract.WatchLogs(opts, "RevealAuction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb1pAuctionRevealAuction)
				if err := _Cb1pAuction.contract.UnpackLog(event, "RevealAuction", log); err != nil {
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
// Solidity: event RevealAuction(uint256 auctionId)
func (_Cb1pAuction *Cb1pAuctionFilterer) ParseRevealAuction(log types.Log) (*Cb1pAuctionRevealAuction, error) {
	event := new(Cb1pAuctionRevealAuction)
	if err := _Cb1pAuction.contract.UnpackLog(event, "RevealAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Cb1pAuctionWithdrawBidIterator is returned from FilterWithdrawBid and is used to iterate over the raw logs and unpacked data for WithdrawBid events raised by the Cb1pAuction contract.
type Cb1pAuctionWithdrawBidIterator struct {
	Event *Cb1pAuctionWithdrawBid // Event containing the contract specifics and raw log

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
func (it *Cb1pAuctionWithdrawBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cb1pAuctionWithdrawBid)
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
		it.Event = new(Cb1pAuctionWithdrawBid)
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
func (it *Cb1pAuctionWithdrawBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cb1pAuctionWithdrawBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cb1pAuctionWithdrawBid represents a WithdrawBid event raised by the Cb1pAuction contract.
type Cb1pAuctionWithdrawBid struct {
	AuctionId *big.Int
	Id        string
	Bidder    common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawBid is a free log retrieval operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auctionId, string id, address bidder, uint256 amount)
func (_Cb1pAuction *Cb1pAuctionFilterer) FilterWithdrawBid(opts *bind.FilterOpts) (*Cb1pAuctionWithdrawBidIterator, error) {

	logs, sub, err := _Cb1pAuction.contract.FilterLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return &Cb1pAuctionWithdrawBidIterator{contract: _Cb1pAuction.contract, event: "WithdrawBid", logs: logs, sub: sub}, nil
}

// WatchWithdrawBid is a free log subscription operation binding the contract event 0x9691d185d35558d296453bdc7848da198cbe1d417744d50c3aa8536d2f662e7c.
//
// Solidity: event WithdrawBid(uint256 auctionId, string id, address bidder, uint256 amount)
func (_Cb1pAuction *Cb1pAuctionFilterer) WatchWithdrawBid(opts *bind.WatchOpts, sink chan<- *Cb1pAuctionWithdrawBid) (event.Subscription, error) {

	logs, sub, err := _Cb1pAuction.contract.WatchLogs(opts, "WithdrawBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cb1pAuctionWithdrawBid)
				if err := _Cb1pAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
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
// Solidity: event WithdrawBid(uint256 auctionId, string id, address bidder, uint256 amount)
func (_Cb1pAuction *Cb1pAuctionFilterer) ParseWithdrawBid(log types.Log) (*Cb1pAuctionWithdrawBid, error) {
	event := new(Cb1pAuctionWithdrawBid)
	if err := _Cb1pAuction.contract.UnpackLog(event, "WithdrawBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
