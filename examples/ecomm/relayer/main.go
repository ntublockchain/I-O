package main

import (
	"flag"
	"strings"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/cb1p_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/cb2p_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/dutch_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/english_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/stable_coin"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	zkNodes = "localhost:2181"

	assetClient *ecomm.AssetClient
	//fabric_network *gateway.Network
	ccsvc *cclib.CCService

	// auctionResults   map[int]*ecomm.FinalizeAuctionArgs
	// auctionResultsMu sync.Mutex

	ethClient *ethclient.Client
	quoClient *ethclient.Client

	rootT *bind.TransactOpts
	// eth_ERC20    stable_coin.StableCoin
	// quo_ERC20    stable_coin.StableCoin
	// fabric_ERC20 string

	// english_auction_info ecomm.AuctionInfo
	// cb1p_auction_info    ecomm.AuctionInfo

	contract_info ecomm.ContractInfo

	eth_english_auction_contract *english_auction.EnglishAuction
	eth_dutch_auction_contract   *dutch_auction.DutchAuction
	eth_cb1p_contract            *cb1p_auction.Cb1pAuction
	eth_cb2p_contract            *cb2p_auction.Cb2pAuction

	quo_english_auction_contract *english_auction.EnglishAuction
	quo_dutch_auction_contract   *dutch_auction.DutchAuction
	quo_cb1p_contract            *cb1p_auction.Cb1pAuction
	quo_cb2p_contract            *cb2p_auction.Cb2pAuction
)

const (
	platform = "eth"
	root_key = "../../keys/key0"
	password = "password"

	contractInfoFile = "../tmp/contract_info.json"
	auctionInfoFile  = "../tmp/auction_info.json"
	logInfoFile      = "../tmp/log.csv"
)

func main() {
	// flag.StringVar(&platform, "p", platform, "Monitors wich platform")

	flag.StringVar(&zkNodes, "zk", zkNodes, "comma separated zoolkeeper nodes")
	flag.Parse()

	// Initialize
	//auctionResults = make(map[int]*ecomm.FinalizeAuctionArgs)

	assetClient = ecomm.NewAssetClient()
	ethClient = ecomm.NewEthClient()
	quoClient = ecomm.NewQuorumClient()
	rootT, _ = cclib.NewTransactor(root_key, password)

	ccsvc, _ = cclib.NewEventService(strings.Split(zkNodes, ","), "relayer") //zookeeper node
	//check(err)

	ecomm.ReadJsonFile(contractInfoFile, &contract_info)

	ccsvc.Register(ecomm.AssetAddingEvent, chainCodeEvent)
	ccsvc.Register(ecomm.AuctionStartingEvent, chainCodeEvent)

	ccsvc.Register(ecomm.EndClosedBidEvent, chainCodeEvent)

	ccsvc.Register(ecomm.DetermineWinnerEvent, chainCodeEvent)
	ccsvc.Register(ecomm.AuctionClosingEvent, chainCodeEvent)

	//ccsvc.Register(ecomm.CancelAuctionEvent, chainCodeEvent)
	ccsvc.Register(ecomm.FinAuctionEvent, chainCodeEvent)

	ccsvc.Register(ecomm.BidEvent, smartContractEvent)
	ccsvc.Register(ecomm.BidHashEvent, smartContractEvent)
	ccsvc.Register(ecomm.WithdrawEvent, smartContractEvent)
	ccsvc.Register(ecomm.CommitAuctionResultEvent, smartContractEvent)
	ccsvc.Register(ecomm.ProvideFeedbackEvent, smartContractEvent)

	err := ccsvc.Start(true)
	check(err)

	startContractListener(contract_info)
	startFabricListener(assetClient)

	// go startAuctionListener("english_auction", english_auction.EthAddr.String(), "eth")
	// go startAuctionListener("english_auction", english_auction.QuoAddr.String(), "quo")
	// startListeningForAuctionEvents()
	// startListeningForAuctionEvents()
}

func load_ERC20() (*stable_coin.StableCoin, *stable_coin.StableCoin, *ecomm.Erc20Client) {
	var contract_info ecomm.ContractInfo
	ecomm.ReadJsonFile(contractInfoFile, &contract_info)

	eth_ERC20, err := stable_coin.NewStableCoin(contract_info.EthERC20, ethClient)
	check(err)

	quo_ERC20, err := stable_coin.NewStableCoin(contract_info.QuoERC20, quoClient)
	check(err)

	fabric_ERC20 := ecomm.NewErc20Client(contract_info.FabricTokenName)

	return eth_ERC20, quo_ERC20, fabric_ERC20
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
