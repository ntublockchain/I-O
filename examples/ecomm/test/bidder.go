package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/cb1p_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/cb2p_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/english_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/stable_coin"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func load_bidder_key(name string) string {
	users, err := ecomm.ReadUsersFromFile(userInfoFile)
	check(err)

	for _, user := range users {
		//fmt.Println("Find ", name, "in", user.UserID)
		if name == user.UserID {
			//fmt.Println("Find!", user.KeyFile)
			return user.KeyFile
		}
	}

	return "../../keys/key2"
}

func bidAuction(auction_id int, amount *big.Int, bid_key, platform string) {
	t := time.Now()

	var contract_info ecomm.ContractInfo
	ecomm.ReadJsonFile(contractInfoFile, &contract_info)
	erc20_address := contract_info.EthERC20
	client := ethClient

	a, err := assetClient.GetAuction(auction_id)
	check(err)

	if a.AucType != "english" {
		log.Fatalf("Incorrect auction type %s", a.AucType)
	}

	auction_addr := common.HexToAddress(a.EthAddr)
	// @todo: require platform either is 'quo' or 'eth'
	if platform != "eth" {
		auction_addr = common.HexToAddress(a.QuorumAddr)
		client = quoClient
		erc20_address = contract_info.QuoERC20
	}

	//var auction_contract english_auction.EnglishAuction
	auction_contract, err := english_auction.NewEnglishAuction(auction_addr, client)

	check(err)

	bidT, err := cclib.NewTransactor(bid_key, "password")
	check(err)

	// Bid more than highest
	if amount.Cmp(big.NewInt(0)) == 0 {
		highest, err := auction_contract.HighestBid(&bind.CallOpts{}, big.NewInt(int64(auction_id)))
		highest.Div(highest, ecomm.DecimalB)
		check(err)
		amount.Add(highest, big.NewInt(1))
		//log.Printf("highest: %s bidAmount: %s", highest, amount)
	}

	//eventID := fmt.Sprintf("%s_%s_%s_%s", a.AssetID, platform, bidT.From.String()[36:], amount)
	keyWords := fmt.Sprintf("%s_%s_%s", platform, bidT.From.String()[36:], amount)
	ecomm.LogEvent(logInfoFile, a.AssetID, ecomm.BidEvent, keyWords, t, "", 0)

	// @todo: Make approve and bid in a single transaction
	// Approve amount of bid through ERC20 contract
	MDAI, _ := stable_coin.NewStableCoin(erc20_address, client)
	// valueB, err := MDAI.BalanceOf(&bind.CallOpts{}, auction_addr)
	// log.Printf("Auction contract orig balance: %s", valueB)

	tx1, err := MDAI.Approve(bidT, auction_addr, big.NewInt(0).Mul(big.NewInt(amount.Int64()), ecomm.DecimalB))
	if err != nil {
		log.Fatalf("Failed to approve: %v", err)
	}
	receipt1 := ecomm.WaitTx(client, tx1, "Approve Auction Contract's allowance")
	// allB, err := MDAI.Allowance(&bind.CallOpts{}, bidT.From, auction_addr)
	// log.Printf("Auction contract orig allowance: %s", allB)

	tx2, err := auction_contract.Bid(bidT, big.NewInt(int64(auction_id)), big.NewInt(0).Mul(big.NewInt(amount.Int64()), ecomm.DecimalB))
	if err != nil {
		log.Fatalf("Failed to bid: %v", err)
	}
	receipt2 := ecomm.WaitTx(client, tx2, fmt.Sprintf("Bid on Auction ID: %d through contract: %s", a.AuctionID, auction_addr))

	total_cost := receipt1.GasUsed + receipt2.GasUsed
	//fmt.Print(total_cost)
	ecomm.UpdateLog(logInfoFile, a.AssetID, ecomm.BidEvent, keyWords, total_cost, "")
}

func bidAuctionH(auction_id int, bidAmount *big.Int, bid_key, platform string) {
	t := time.Now()

	var contract_info ecomm.ContractInfo
	ecomm.ReadJsonFile(contractInfoFile, &contract_info)
	//erc20_address := contract_info.EthERC20
	client := ethClient

	auction, err := assetClient.GetAuction(auction_id)
	check(err)
	if auction.AucType == "english" {
		log.Fatalf("Incorrect auction id which only support %s", auction.AucType)
	}

	auction_addr := common.HexToAddress(auction.EthAddr)
	if platform != "eth" {
		auction_addr = common.HexToAddress(auction.QuorumAddr)
		client = quoClient
	}

	// auction_contract, err := cb1p_auction.NewCb1pAuction(auction_addr, client)
	// check(err)
	var auction_contract ecomm.AuctionContractCloseBid
	//var auction_contract_close_bid ecomm.AuctionContractCloseBid

	switch auction.AucType {
	case "cb1p":
		auction_contract, err = cb1p_auction.NewCb1pAuction(auction_addr, client)
	case "cb2p":
		auction_contract, err = cb2p_auction.NewCb2pAuction(auction_addr, client)
	// case "cb1p":
	// 	auction_contract_cb, err = cb1p_auction.NewCb1pAuction(auction_addr, client)
	default:
		log.Fatalf("Incorrect Auction Type!")
		return
	}
	check(err)

	bidT, err := cclib.NewTransactor(bid_key, "password")
	check(err)

	bidAmount.Mul(bidAmount, ecomm.DecimalB)
	bidHash := solsha3.SoliditySHA3(solsha3.Uint256(bidAmount))
	//log.Printf("Calculated hash of uint(%s): %s", bidAmount, hex.EncodeToString(bidHash[:]))

	// Convert bidHash to [32]byte to match the Go binding's expectation
	var bidHashArray [32]byte
	copy(bidHashArray[:], bidHash)

	keyWords := fmt.Sprintf("%s_%s_%s", platform, bidT.From.String()[36:], hex.EncodeToString(bidHashArray[:])[60:])
	ecomm.LogEvent(logInfoFile, auction.AssetID, ecomm.BidHashEvent, keyWords, t, "", 0)

	tx, _ := auction_contract.Bid(bidT, big.NewInt(int64(auction_id)), bidHashArray)
	receipt := ecomm.WaitTx(client, tx, fmt.Sprintf("Bid on Auction ID: %d through contract: %s", auction.AuctionID, auction_addr))
	ecomm.UpdateLog(logInfoFile, auction.AssetID, ecomm.BidHashEvent, keyWords, receipt.GasUsed, "")
}

func revealBid(auction_id int, amount *big.Int, bid_key, platform string) {
	t := time.Now()

	var contract_info ecomm.ContractInfo
	ecomm.ReadJsonFile(contractInfoFile, &contract_info)
	erc20_address := contract_info.EthERC20
	client := ethClient

	auction, err := assetClient.GetAuction(auction_id)
	check(err)
	if auction.AucType == "english" {
		log.Fatalf("Incorrect auction id which only support %s", auction.AucType)
	}

	auction_addr := common.HexToAddress(auction.EthAddr)
	// @todo: require platform either is 'quo' or 'eth'
	if platform != "eth" {
		auction_addr = common.HexToAddress(auction.QuorumAddr)
		client = quoClient
		erc20_address = contract_info.QuoERC20
	}
	//auction_contract, err := cb1p_auction.NewCb1pAuction(auction_addr, client)
	var auction_contract ecomm.AuctionContractCloseBid
	//var auction_contract_close_bid ecomm.AuctionContractCloseBid

	switch auction.AucType {
	case "cb1p":
		auction_contract, err = cb1p_auction.NewCb1pAuction(auction_addr, client)
	case "cb2p":
		auction_contract, err = cb2p_auction.NewCb2pAuction(auction_addr, client)
	default:
		log.Fatalf("Incorrect Auction Type!")
		return
	}
	check(err)

	bidT, err := cclib.NewTransactor(bid_key, "password")
	check(err)

	//eventID := fmt.Sprintf("%s_%s_%s_%s", a.AssetID, platform, bidT.From.String()[36:], amount)
	keyWords := fmt.Sprintf("%s_%s_%s", platform, bidT.From.String()[36:], amount)
	ecomm.LogEvent(logInfoFile, auction.AssetID, ecomm.BidEvent, keyWords, t, "", 0)
	amount.Mul(amount, ecomm.DecimalB)

	// @todo: Make approve and bid in a single transaction
	// Approve amount of bid through ERC20 contract
	MDAI, _ := stable_coin.NewStableCoin(erc20_address, client)
	// valueB, err := MDAI.BalanceOf(&bind.CallOpts{}, auction_addr)
	// log.Printf("Auction contract orig balance: %s", valueB)

	tx1, err := MDAI.Approve(bidT, auction_addr, amount)
	if err != nil {
		log.Fatalf("Failed to approve: %v", err)
	}
	receipt1 := ecomm.WaitTx(client, tx1, "Approve Auction Contract's allowance")
	// allB, err := MDAI.Allowance(&bind.CallOpts{}, bidT.From, auction_addr)
	// log.Printf("Auction contract orig allowance: %s", allB)

	tx2, err := auction_contract.Reveal(bidT, big.NewInt(int64(auction_id)), amount)
	if err != nil {
		log.Fatalf("Failed to bid: %v", err)
	}
	receipt2 := ecomm.WaitTx(client, tx2, fmt.Sprintf("Bid on Auction ID: %d through contract: %s", auction.AuctionID, auction_addr))

	total_cost := receipt1.GasUsed + receipt2.GasUsed
	//fmt.Print(total_cost)
	ecomm.UpdateLog(logInfoFile, auction.AssetID, ecomm.BidEvent, keyWords, total_cost, "")
}

func withdraw(auction_id int, bid_key, platform string) {
	t := time.Now()
	client := ethClient

	auction, err := assetClient.GetAuction(auction_id)
	check(err)

	auction_addr := common.HexToAddress(auction.EthAddr)
	if platform != "eth" {
		auction_addr = common.HexToAddress(auction.QuorumAddr)
		client = quoClient
	}

	bidT, err := cclib.NewTransactor(bid_key, "password")
	check(err)
	//eventID := auction.AssetID + "_" + platform + "_" + bidT.From.String()[36:]
	//@todo withdraw bidAmount
	keyWords := fmt.Sprintf("%s_%s_%s", auction.AucType, platform, bidT.From.String()[36:])

	ecomm.LogEvent(logInfoFile, "NA", ecomm.WithdrawEvent, keyWords, t, "", 0)

	// same interface for all 4 kinds auction contracts
	// @todo support all 4 auctions
	auction_contract, err := english_auction.NewEnglishAuction(auction_addr, client)
	check(err)

	tx, err := auction_contract.Withdraw(bidT, big.NewInt(int64(auction_id)))
	check(err)
	receipt := ecomm.WaitTx(client, tx, fmt.Sprintf("Withdraw bid on Auction ID: %d through contract: %s", auction.AuctionID, auction_addr))
	//log.Printf("Gas: %d", receipt.GasUsed)
	ecomm.UpdateLog(logInfoFile, "NA", ecomm.WithdrawEvent, keyWords, receipt.GasUsed, "")
	//debugTransaction(tx)
	// log
	/////////////
}

// func check_winner(auction_id int, bid_key, platform string) {
// 	client := ethClient
// 	bidT, err := cclib.NewTransactor(bid_key, password)
// 	check(err)

// 	// Get Auction Contract deployed on Eth/Quo
// 	//assetClient := ecomm.NewAssetClient() // return Fabric asset contract
// 	a, err := assetClient.GetAuction(auction_id)
// 	check(err)

// 	auction_addr := common.HexToAddress(a.EthAddr)
// 	if platform != "eth" {
// 		auction_addr = common.HexToAddress(a.QuorumAddr)
// 		client = quoClient
// 	}

// 	auction_contract, err := english_auction.NewEnglishAuction(auction_addr, client)
// 	check(err)

// 	// Check winner
// 	highestBidder, err := auction_contract.HighestBidder(&bind.CallOpts{}, big.NewInt(int64(auction_id)))
// 	check(err)

// 	if bidT.From == highestBidder {
// 		fmt.Println("Waiting your response (abt/prcd) for Auction", auction_id)
// 	} else {
// 		fmt.Println("highest bidder:", highestBidder.Hex())
// 	}

// 	highestBid, err := auction_contract.HighestBid(&bind.CallOpts{}, big.NewInt(int64(auction_id)))
// 	check(err)
// 	fmt.Println("highest bid:", highestBid)
// }

func sign_auction_result(auction_id int) {
	t := time.Now()
	auction, err := assetClient.GetAuction(auction_id)
	check(err)

	ecomm.LogEvent(logInfoFile, auction.AssetID, ecomm.CommitAuctionResultEvent, auction.HighestBidder[36:], t, "", 0)

	var bidKey string
	inval_addr := true
	accounts, _ := ecomm.ReadUsersFromFile(userInfoFile)

	for i := 1; i < 9; i++ {
		if auction.HighestBidder == accounts[i].Address.String() {
			bidKey = load_bidder_key(accounts[i].UserID)
			inval_addr = false
			break
		}
	}

	if inval_addr {
		log.Fatalf("Invalid winner address%s", auction.HighestBidder)
	}

	var auction_addr common.Address
	var client *ethclient.Client
	if auction.HighestBidPlatform == "eth" {
		auction_addr = common.HexToAddress(auction.EthAddr)
		client = ethClient
	} else {
		auction_addr = common.HexToAddress(auction.QuorumAddr)
		client = quoClient
	}

	var auction_contract ecomm.AuctionContract
	// load auction contract
	switch auction.AucType {
	case "english":
		auction_contract, err = english_auction.NewEnglishAuction(auction_addr, client)
		check(err)
	case "dutch":
	case "cb1p":
		auction_contract, err = cb1p_auction.NewCb1pAuction(auction_addr, client)
		check(err)
	case "cb2p":
	default:
		log.Fatalf("Auction type error")
	}

	bidT, err := cclib.NewTransactor(bidKey, "password")
	check(err)
	auction_result := &ecomm.AuctionResult{
		Platform:    auction.HighestBidPlatform,
		AuctionID:   auction.AuctionID,
		AuctionAddr: auction_addr.Hex(),

		HighestBid:    auction.HighestBid,
		HighestBidder: bidT.From.Hex(),
	}

	signer, _ := cclib.NewSigner(bidKey, password)
	sig, err := signer.Sign(auction_result.Hash())
	check(err)

	auction_result.Signature = sig
	jsonString, err := json.Marshal(auction_result)
	check(err)

	tx, err := auction_contract.Commit(bidT, big.NewInt(int64(auction.AuctionID)), string(jsonString))
	check(err)
	receipt := ecomm.WaitTx(client, tx, fmt.Sprintf("Sign Auction Result on Auction ID: %d through contract: %s", auction.AuctionID, auction_addr))

	ecomm.UpdateLog(logInfoFile, auction.AssetID, ecomm.CommitAuctionResultEvent, bidT.From.String()[36:], receipt.GasUsed, auction.HighestBidPlatform)

}

func provide_feedback(auction_id, score int, feedback string) {
	// @reset timer
	t := time.Now()
	auction, err := assetClient.GetAuction(auction_id)
	check(err)

	ecomm.LogEvent(logInfoFile, auction.AssetID, ecomm.ProvideFeedbackEvent, auction.HighestBidder[36:], t, "", 0)

	var bidKey string
	inval_addr := true
	accounts, _ := ecomm.ReadUsersFromFile(userInfoFile)

	for i := 1; i < 9; i++ {
		if auction.HighestBidder == accounts[i].Address.String() {
			bidKey = load_bidder_key(accounts[i].UserID)
			inval_addr = false
			break
		}
	}

	if inval_addr {
		log.Fatalf("Invalid winner address%s", auction.HighestBidder)
	}

	var auction_addr common.Address
	var client *ethclient.Client
	if auction.HighestBidPlatform == "eth" {
		auction_addr = common.HexToAddress(auction.EthAddr)
		client = ethClient
	} else {
		auction_addr = common.HexToAddress(auction.QuorumAddr)
		client = quoClient
	}

	var auction_contract ecomm.AuctionContract
	// load auction contract
	switch auction.AucType {
	case "english":
		auction_contract, err = english_auction.NewEnglishAuction(auction_addr, client)
		check(err)
	case "dutch":
	case "cb1p":
		auction_contract, err = cb1p_auction.NewCb1pAuction(auction_addr, client)
		check(err)
	case "cb2p":
	default:
		log.Fatalf("Auction type error")
	}

	bidT, err := cclib.NewTransactor(bidKey, "password")
	check(err)

	tx, _ := auction_contract.ProvideFeedback(bidT, big.NewInt(int64(auction_id)), big.NewInt(int64(score)), feedback)
	receipt := ecomm.WaitTx(client, tx, fmt.Sprintf("Provide feedback on Auction ID: %d", auction.AuctionID))
	//t = time.Now()

	// payload, _ = json.Marshal(&ecomm.Tx{
	// 	Platform: platform,
	// 	Type:     "Feedback",
	// 	Hash:     receipt.TxHash,
	// })
	ecomm.UpdateLog(logInfoFile, auction.AssetID, ecomm.ProvideFeedbackEvent, auction.HighestBidder[36:], receipt.GasUsed, "")
	// cclib.LogEventToFile(logInfoFile, ecomm.TransactionMinedEvent, payload, t, timeInfoFile)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
