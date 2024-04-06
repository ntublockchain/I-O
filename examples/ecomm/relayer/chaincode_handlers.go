package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/contracts/cb1p_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/cb2p_auction"
	"github.com/Guy1m0/Blockchain-I-O/contracts/english_auction"

	//"github.com/Guy1m0/Blockchain-I-O/contracts/eth_auction"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// fabric relayer
func handleAddAssetEvent(eventPayload []byte) error {
	log.Println("[fabric] Get Asset")
	t := time.Now()

	var result ecomm.AssetAddingEventPayload
	err := json.Unmarshal(eventPayload, &result)
	check(err)

	assetID := result.AssetID
	asset, err := assetClient.GetAsset(assetID)
	check(err)

	auc_type := result.AucType
	//fmt.Println("Auc Type:", auc_type)
	ecomm.LogEvent(logInfoFile, assetID, ecomm.AssetAddingEvent, auc_type, t, "", 0)

	//payloadJSON, _ := json.Marshal(asset)
	wrapper := ecomm.EventWrapper{Type: "Asset", Result: eventPayload}
	payload, _ := json.Marshal(wrapper)

	// @todo: change this payload to be AssetAddingEventPayload?
	ccsvc.Publish(ecomm.AssetAddingEvent, payload)

	log.Println("[fabric] Start auction")
	t = time.Now()

	var ethAddr, quoAddr common.Address
	switch auc_type {
	case "english":
		ethAddr = contract_info.EnglishAuction.EthAddr
		quoAddr = contract_info.EnglishAuction.QuoAddr
	// case "dutch":
	// 	ethAddr = contract_info.DutchAuction.EthAddr
	// 	quoAddr = contract_info.DutchAuction.QuoAddr
	case "cb1p":
		ethAddr = contract_info.Cb1pAuction.EthAddr
		quoAddr = contract_info.Cb1pAuction.QuoAddr
	// case "cb2p":
	// 	ethAddr = contract_info.Cb2pAuction.EthAddr
	// 	quoAddr = contract_info.Cb2pAuction.QuoAddr
	default:
		fmt.Println("Auction type error")
	}

	args := ecomm.StartAuctionArgs{
		AssetID:    asset.ID,
		AucType:    auc_type,
		EthAddr:    ethAddr.String(),
		QuorumAddr: quoAddr.String(),
	}

	_, err = assetClient.StartAuction(args)

	check(err)
	ecomm.LogEvent(logInfoFile, assetID, ecomm.AuctionStartingEvent, auc_type, t, "", 0)
	return nil
}

// fabric relayer
func handleStartAuctionEvent(eventPayload []byte) error {
	var result ecomm.StartAuctionEventPayload
	var tx *types.Transaction
	var receipt1, receipt2 *types.Receipt

	authT, _ := cclib.NewTransactor(root_key, password)

	err := json.Unmarshal(eventPayload, &result)
	check(err)

	auction, err := assetClient.GetAuction(result.ID)
	check(err)
	// ethAddr := auction.EthAddr
	// quoAddr := auction.QuorumAddr

	switch auction.AucType {
	case "english":
		tx, err = eth_english_auction_contract.Create(authT, big.NewInt(int64(auction.AuctionID)), auction.AssetID, result.Owner)
		check(err)
		receipt1 = ecomm.WaitTx(ethClient, tx, fmt.Sprintf("Create new auction with type: %s and ID: %d", auction.AucType, auction.AuctionID))

		tx, err = quo_english_auction_contract.Create(authT, big.NewInt(int64(auction.AuctionID)), auction.AssetID, result.Owner)
		check(err)
		receipt2 = ecomm.WaitTx(quoClient, tx, fmt.Sprintf("Create new auction with type: %s and ID: %d", auction.AucType, auction.AuctionID))
	case "dutch":

	case "cb1p":
		tx, err = eth_cb1p_contract.Create(authT, big.NewInt(int64(auction.AuctionID)), auction.AssetID, result.Owner)
		check(err)
		receipt1 = ecomm.WaitTx(ethClient, tx, fmt.Sprintf("Create new auction with type: %s and ID: %d", auction.AucType, auction.AuctionID))

		tx, err = quo_cb1p_contract.Create(authT, big.NewInt(int64(auction.AuctionID)), auction.AssetID, result.Owner)
		check(err)
		receipt2 = ecomm.WaitTx(quoClient, tx, fmt.Sprintf("Create new auction with type: %s and ID: %d", auction.AucType, auction.AuctionID))

	case "cb2p":

	default:
		log.Fatalf("Auction type error")

	}

	cost := receipt1.GasUsed
	note := "ETH:" + strconv.FormatUint(cost, 10)
	cost += receipt2.GasUsed
	note += " QUO:" + strconv.FormatUint(receipt2.GasUsed, 10)

	t := time.Now()
	ecomm.LogEvent(logInfoFile, auction.AssetID, ecomm.AuctionStartingEvent, auction.AucType, t, note, cost)

	log.Println("[fabirc] Start Auction with ID: ", result.ID)
	//log.Println("AuctionID", auction.AuctionID)

	payloadJSON, _ := json.Marshal(auction)
	wrapper := ecomm.EventWrapper{Type: "Start Auction", Result: payloadJSON}
	payload, _ := json.Marshal(wrapper)

	err = ccsvc.Publish(ecomm.AuctionStartingEvent, payload)
	ecomm.AddAuctionToFile(auctionInfoFile, ecomm.AuctionInfo{
		AuctionID: auction.AuctionID,
		AucType:   auction.AucType,
		Owner:     common.HexToAddress(result.Owner),
		AssetID:   auction.AssetID,
		EthAddr:   common.HexToAddress(auction.EthAddr),
		QuoAddr:   common.HexToAddress(auction.QuorumAddr),
	})

	return err
}

func handleEndClosedBidEvent(eventPayload []byte) error {
	log.Printf("[fabric] End ClosedBid")
	t := time.Now()

	var result ecomm.Auction
	err := json.Unmarshal(eventPayload, &result)
	check(err)

	log.Println(result.AssetID, ecomm.EndClosedBidEvent, result.AucType)
	ecomm.LogEvent(logInfoFile, result.AssetID, ecomm.EndClosedBidEvent, result.AucType, t, "", 0)

	payloadJSON, _ := json.Marshal(result)
	wrapper := ecomm.EventWrapper{Type: "End ClosedBid", Result: payloadJSON}
	payload, _ := json.Marshal(wrapper)

	ccsvc.Publish(ecomm.EndClosedBidEvent, payload)

	t = time.Now()
	ecomm.LogEvent(logInfoFile, result.AssetID, ecomm.AuctionRevealingEvent, result.AucType, t, "", 0)

	log.Println("[ETH/QUO] Change contract state")
	authT, err := cclib.NewTransactor(root_key, password)
	check(err)

	var eth_auction_contract, quo_auction_contract ecomm.AuctionContractCloseBid
	eth_auction_addr := common.HexToAddress(result.EthAddr)
	quo_auction_addr := common.HexToAddress(result.QuorumAddr)

	// load auction contract
	switch result.AucType {
	case "cb1p":
		eth_auction_contract, err = cb1p_auction.NewCb1pAuction(eth_auction_addr, ethClient)
		check(err)

		quo_auction_contract, err = cb1p_auction.NewCb1pAuction(quo_auction_addr, quoClient)
		check(err)
	case "cb2p":
		eth_auction_contract, err = cb2p_auction.NewCb2pAuction(eth_auction_addr, ethClient)
		check(err)

		quo_auction_contract, err = cb2p_auction.NewCb2pAuction(quo_auction_addr, quoClient)
		check(err)
	default:
		log.Fatalf("Auction type error")
	}

	// Change Auction Contract on Eth
	tx, _ := eth_auction_contract.RevealAuction(authT, big.NewInt(int64(result.AuctionID)))
	receipt1 := ecomm.WaitTx(ethClient, tx, fmt.Sprintf("Change Auction %d status to 'Reveal'", result.AuctionID))

	// Change Auction Contract on Quo
	tx, _ = quo_auction_contract.RevealAuction(authT, big.NewInt(int64(result.AuctionID)))
	receipt2 := ecomm.WaitTx(quoClient, tx, fmt.Sprintf("Change Auction %d status to 'Reveal'", result.AuctionID))

	cost := receipt1.GasUsed
	//note := "ETH:" + strconv.FormatUint(cost, 10)
	cost += receipt2.GasUsed
	//note += " QUO:" + strconv.FormatUint(receipt2.GasUsed, 10)

	//t := time.Now()
	ecomm.UpdateLog(logInfoFile, result.AssetID, ecomm.AuctionRevealingEvent, result.AucType, cost, "")

	err = ccsvc.Publish(ecomm.AuctionRevealingEvent, payload)

	return err
}

// func handleCancelAuctionEvent(eventPayload []byte) error {
// 	t := time.Now()

// 	var result ecomm.Auction
// 	err := json.Unmarshal(eventPayload, &result)
// 	check(err)

// 	auction, err := assetClient.GetAuction(result.AuctionID)
// 	check(err)
// 	log.Println("[fabric] Cancel Auction with ID:", result.AuctionID)

// 	ecomm.LogEvent(logInfoFile, auction.AssetID, ecomm.CancelAuctionEvent, "", t, "", 0)

// 	log.Println("[ETH/QUO] Close auctions on both platforms")
// 	// load auction contract
// 	eth_auction_addr := common.HexToAddress(auction.EthAddr)
// 	eth_auction_contract, err := english_auction.NewEnglishAuction(eth_auction_addr, ethClient)
// 	check(err)

// 	quo_auction_addr := common.HexToAddress(auction.QuorumAddr)
// 	quo_auction_contract, err := english_auction.NewEnglishAuction(quo_auction_addr, quoClient)
// 	check(err)

// 	log.Println("[ETH/QUO] Change contract state")
// 	authT, err := cclib.NewTransactor(root_key, password)
// 	check(err)

// 	// Change Auction Contract on Eth
// 	tx, _ := eth_auction_contract.CloseAuction(authT, true)
// 	receipt := ecomm.WaitTx(ethClient, tx, fmt.Sprintf("Change Auction %s status to 'ENDED'", eth_auction_addr))

// 	cost := receipt.GasUsed
// 	note := "ETH:" + strconv.FormatUint(cost, 10)

// 	// Change Auction Contract on Quo
// 	tx, _ = quo_auction_contract.CloseAuction(authT, true)
// 	receipt = ecomm.WaitTx(quoClient, tx, fmt.Sprintf("Change Auction %s status to 'ENDED'", quo_auction_addr))

// 	cost += receipt.GasUsed
// 	note += " QUO:" + strconv.FormatUint(receipt.GasUsed, 10)

// 	ecomm.UpdateLog(logInfoFile, auction.AssetID, ecomm.CancelAuctionEvent, "", cost, note)
// 	//ccsvc.Publish(ecomm.AuctionClosingEvent, payload)

// 	payloadJSON, _ := json.Marshal(auction)
// 	wrapper := ecomm.EventWrapper{Type: "Cancel Auction", Result: payloadJSON}
// 	payload, _ := json.Marshal(wrapper)

// 	return ccsvc.Publish(ecomm.CancelAuctionEvent, payload)
// }

func handleDetWinnerEvent(eventPayload []byte) error {
	//t := time.Now()
	var result ecomm.Auction
	err := json.Unmarshal(eventPayload, &result)
	check(err)

	log.Println("[fabric] DetermineWinner with ID:", result.AuctionID)
	//ecomm.LogEvent(logInfoFile, result.AssetID, ecomm.DetermineWinnerEvent, "", t, "", 0)

	var eth_auction_contract, quo_auction_contract ecomm.AuctionContract
	eth_auction_addr := common.HexToAddress(result.EthAddr)
	quo_auction_addr := common.HexToAddress(result.QuorumAddr)

	// load auction contract
	switch result.AucType {
	case "english":
		eth_auction_contract, err = english_auction.NewEnglishAuction(eth_auction_addr, ethClient)
		check(err)

		quo_auction_contract, err = english_auction.NewEnglishAuction(quo_auction_addr, quoClient)
		check(err)
	case "dutch":
	case "cb1p":
		eth_auction_contract, err = cb1p_auction.NewCb1pAuction(eth_auction_addr, ethClient)
		check(err)

		quo_auction_contract, err = cb1p_auction.NewCb1pAuction(quo_auction_addr, quoClient)
		check(err)
	case "cb2p":

	default:
		log.Fatalf("Auction type error")
	}

	log.Println("[ETH/QUO] Determin winner")
	// Check highest bid
	eth_highestBid, _ := eth_auction_contract.HighestBid(&bind.CallOpts{}, big.NewInt(int64(result.AuctionID)))
	quo_highestBid, _ := quo_auction_contract.HighestBid(&bind.CallOpts{}, big.NewInt(int64(result.AuctionID)))

	highestBidPlatform := "eth"
	highestBid := eth_highestBid.String()
	highestBidder, _ := eth_auction_contract.HighestBidder(&bind.CallOpts{}, big.NewInt(int64(result.AuctionID)))

	if eth_highestBid.Cmp(quo_highestBid) < 0 {
		highestBidPlatform = "quo"
		highestBid = quo_highestBid.String()
		highestBidder, _ = quo_auction_contract.HighestBidder(&bind.CallOpts{}, big.NewInt(int64(result.AuctionID)))
	}
	t := time.Now()
	ecomm.LogEvent(logInfoFile, result.AssetID, ecomm.DetermineWinnerEvent, "", t, fmt.Sprintf("Highest Bidder: %s on Platform: %s", highestBidder, strings.ToUpper(highestBidPlatform)), 0)

	args := ecomm.CloseAuctionArgs{
		AuctionID: result.AuctionID,

		HighestBid:         highestBid,
		HighestBidder:      highestBidder.String(),
		HighestBidPlatform: highestBidPlatform,
	}

	t = time.Now()

	payloadJSON, _ := json.Marshal(args)
	wrapper := ecomm.EventWrapper{Type: "Determine Winner", Result: payloadJSON}
	payload, _ := json.Marshal(wrapper)

	ccsvc.Publish(ecomm.DetermineWinnerEvent, payload)
	ecomm.LogEvent(logInfoFile, result.AssetID, ecomm.AuctionClosingEvent, "", t, "", 0)

	_, err = assetClient.CloseAuction(args)
	check(err)

	return nil

}
func handleCloseAuctionEvent(eventPayload []byte) error {
	var result ecomm.Auction
	err := json.Unmarshal(eventPayload, &result)
	check(err)

	log.Println("[ETH/QUO] Change contract state")
	authT, err := cclib.NewTransactor(root_key, password)
	check(err)

	var eth_auction_contract, quo_auction_contract ecomm.AuctionContract
	eth_auction_addr := common.HexToAddress(result.EthAddr)
	quo_auction_addr := common.HexToAddress(result.QuorumAddr)

	// load auction contract
	switch result.AucType {
	case "english":
		eth_auction_contract, err = english_auction.NewEnglishAuction(eth_auction_addr, ethClient)
		check(err)

		quo_auction_contract, err = english_auction.NewEnglishAuction(quo_auction_addr, quoClient)
		check(err)
	case "dutch":
	case "cb1p":
		eth_auction_contract, err = cb1p_auction.NewCb1pAuction(eth_auction_addr, ethClient)
		check(err)

		quo_auction_contract, err = cb1p_auction.NewCb1pAuction(quo_auction_addr, quoClient)
		check(err)
	case "cb2p":

	default:
		log.Fatalf("Auction type error")
	}

	var eth_bool, quo_bool bool
	if result.HighestBidPlatform == "eth" {
		eth_bool = false
		quo_bool = true
	} else {
		eth_bool = true
		quo_bool = false
	}

	// Change Auction Contract on Eth
	tx, _ := eth_auction_contract.CloseAuction(authT, big.NewInt(int64(result.AuctionID)), eth_bool)
	receipt1 := ecomm.WaitTx(ethClient, tx, fmt.Sprintf("Change Auction %d status to 'ENDING'", result.AuctionID))

	// Change Auction Contract on Quo
	tx, _ = quo_auction_contract.CloseAuction(authT, big.NewInt(int64(result.AuctionID)), quo_bool)
	receipt2 := ecomm.WaitTx(quoClient, tx, fmt.Sprintf("Change Auction %d status to 'ENDING'", result.AuctionID))

	cost := receipt1.GasUsed
	//note := "ETH:" + strconv.FormatUint(cost, 10)
	cost += receipt2.GasUsed
	//note += " QUO:" + strconv.FormatUint(receipt2.GasUsed, 10)

	t := time.Now()
	ecomm.LogEvent(logInfoFile, result.AssetID, ecomm.AuctionClosingEvent, "", t, "", cost)

	payloadJSON, _ := json.Marshal(result)
	wrapper := ecomm.EventWrapper{Type: "Close Auction", Result: payloadJSON}
	payload, _ := json.Marshal(wrapper)

	err = ccsvc.Publish(ecomm.AuctionClosingEvent, payload)

	return err
}

func handleAuctionClosedEvent(eventPayload []byte) error {
	var result ecomm.Auction
	err := json.Unmarshal(eventPayload, &result)
	check(err)

	auction, err := assetClient.GetAuction(result.AuctionID)
	check(err)

	log.Println("[fabric] Close Auction with ID:", auction.AuctionID)

	if auction.Status != "closed" {
		log.Fatalf("Failed to closed")
	}

	// atomic swap
	// if proceed {
	_, _, fabric_ERC20 := load_ERC20()
	amt, _ := new(big.Int).SetString(auction.HighestBid, 10)
	quotient := new(big.Int).Div(amt, ecomm.DecimalB)

	//asset, _ := assetClient.GetAsset(auction.AssetID)

	_, err = fabric_ERC20.Transfer(auction.AssetOwner, quotient.String())
	check(err)

	var from common.Address
	var client *ethclient.Client

	if auction.HighestBidPlatform == "eth" {
		from = common.HexToAddress(auction.EthAddr)
		client = ethClient
	} else {
		from = common.HexToAddress(auction.QuorumAddr)
		client = quoClient
	}

	auction_contract, err := english_auction.NewEnglishAuction(from, client)
	check(err)

	tx, err := auction_contract.Pay(rootT, big.NewInt(int64(auction.AuctionID)))
	check(err)
	receipt := ecomm.WaitTx(client, tx, fmt.Sprintf("Burn the bid %s placed by the winner %s", auction.HighestBid, auction.HighestBidder))

	t := time.Now()
	ecomm.LogEvent(logInfoFile, auction.AssetID, ecomm.FinAuctionEvent, "", t, auction.HighestBidPlatform, receipt.GasUsed)

	payloadJSON, _ := json.Marshal(auction)
	wrapper := ecomm.EventWrapper{Type: "Fin Auction", Result: payloadJSON}
	payload, _ := json.Marshal(wrapper)

	ccsvc.Publish(ecomm.FinAuctionEvent, payload)
	return nil
}

func chainCodeEvent(eventPayload []byte) {
	t := time.Now()
	var wrapper ecomm.EventWrapper
	var event, assetId string
	var keyWords = ""

	err := json.Unmarshal([]byte(eventPayload), &wrapper)
	check(err)
	log.Printf("[Kafka] Received eventPayload with type %s", wrapper.Type)

	switch wrapper.Type {
	case "Asset":
		var asset ecomm.AssetAddingEventPayload
		err = json.Unmarshal(wrapper.Result, &asset)
		check(err)

		event = ecomm.AssetAddingEvent
		assetId = asset.AssetID
		keyWords = asset.AucType
	case "Start Auction":
		var auction ecomm.Auction
		err = json.Unmarshal(wrapper.Result, &auction)
		check(err)

		event = ecomm.AuctionStartingEvent
		assetId = auction.AssetID
		keyWords = auction.AucType
	case "End ClosedBid":
		var auction ecomm.Auction
		err = json.Unmarshal(wrapper.Result, &auction)
		check(err)

		event = ecomm.EndClosedBidEvent
		assetId = auction.AssetID
		keyWords = auction.AucType
	case "Cancel Auction":
		var auction ecomm.Auction
		err = json.Unmarshal(wrapper.Result, &auction)
		check(err)

		event = ecomm.CancelAuctionEvent
		assetId = auction.AssetID
	case "Determine Winner":
		var args ecomm.CloseAuctionArgs
		err = json.Unmarshal(wrapper.Result, &args)
		check(err)
		auction, err := assetClient.GetAuction(args.AuctionID)
		check(err)

		event = ecomm.DetermineWinnerEvent
		assetId = auction.AssetID
		//log.Printf("assetId: %s, keywords: %s", assetId, keyWords)
	case "Close Auction":
		var auction ecomm.Auction
		err = json.Unmarshal(wrapper.Result, &auction)
		check(err)

		event = ecomm.AuctionClosingEvent
		assetId = auction.AssetID
	case "Fin Auction":
		var auction ecomm.Auction
		err = json.Unmarshal(wrapper.Result, &auction)
		check(err)

		event = ecomm.FinAuctionEvent
		assetId = auction.AssetID
	default:
		log.Fatalf("Unknown type: %s\n", wrapper.Type)
	}
	//time.Sleep(5 * time.Second)
	//log.Println("Kafka received event:", event, "with ID:", eventID)
	//cclib.LogEventToFile(logInfoFile, ecomm.KafkaReceivedEvent, payload, t, timeInfoFile)
	ecomm.LogEvent(logInfoFile, assetId, event, keyWords, t, "", 0)
}
