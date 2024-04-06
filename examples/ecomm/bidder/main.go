package main

import (
	"flag"
	"fmt"

	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	platform = "eth"
	auc_type = "english"

	usr_name = "Bidder 1"
	bid_key  string

	ethClient *ethclient.Client
	quoClient *ethclient.Client

	assetClient *ecomm.AssetClient
)

const (
	password = "password"

	contractInfoFile = "../tmp/contract_info.json"
	userInfoFile     = "../tmp/user_info.json"
	logInfoFile      = "../tmp/log.csv"
	timeInfoFile     = "../tmp/timer"
)

func main() {
	// load clients
	var err error
	ethClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8545", "localhost"))
	check(err)
	quoClient, err = ethclient.Dial(fmt.Sprintf("http://%s:8546", "localhost"))
	check(err)
	assetClient = ecomm.NewAssetClient()

	// // Handle bash command
	// command := flag.String("c", "", "command")
	// id_ := flag.String("id", "", "Auction ID")
	// amount_ := flag.String("amt", "", "Bid amount")

	// feedback := flag.String("fb", "", "Detail feedback with format 'score@comments'")

	flag.StringVar(&platform, "p", platform, "specify platform")
	flag.StringVar(&auc_type, "t", auc_type, "choose auction type")
	flag.StringVar(&usr_name, "usr", usr_name, "Load User/Bidder Information")
	flag.Parse()

	fmt.Println("Load User/Bidder: ", usr_name)
	// bid_key = ecomm.load_bidder_key(usr_name)

	// switch *command {
	// case "bid":
	// 	amount := new(big.Int)
	// 	amount.SetString(*amount_, 10)
	// 	id, _ := strconv.Atoi(*id_)
	// 	ecomm.bidAuction(id, amount)
	// case "bidH":
	// 	amount := new(big.Int)
	// 	amount.SetString(*amount_, 10)
	// 	id, _ := strconv.Atoi(*id_)
	// 	ecomm.bidAuctionH(id, amount)
	// case "check":
	// 	id, _ := strconv.Atoi(*id_)
	// 	ecomm.check_winner(id)
	// case "prcd":
	// 	id, _ := strconv.Atoi(*id_)
	// 	ecomm.sign_auction_result(id, true)
	// case "abt":
	// 	id, _ := strconv.Atoi(*id_)
	// 	ecomm.sign_auction_result(id, false)
	// case "with":
	// 	id, _ := strconv.Atoi(*id_)
	// 	ecomm.withdraw(id)
	// case "rate":
	// 	id, _ := strconv.Atoi(*id_)
	// 	ecomm.provide_feedback(id, *feedback)
	// }
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
