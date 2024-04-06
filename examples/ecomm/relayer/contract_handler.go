package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
)

// Smart Contract handler
func handleHighestBidIncreasedEvent(eventPayload ecomm.HighestBidIncreased, bid ecomm.Bid, t time.Time) error {
	log.Printf("[%s] HighestBidIncreased Event", strings.ToUpper(bid.Platform))

	amount := new(big.Int).Div(eventPayload.BidAmount, ecomm.DecimalB).String()
	asset, _ := assetClient.GetAsset(eventPayload.Id)
	keyWords := fmt.Sprintf("%s_%s_%s", bid.Platform, eventPayload.Bidder.String()[36:], amount)
	//eventID := eventPayload.Id + "_" + bid.Platform + "_" + eventPayload.Bidder.String()[36:]
	ecomm.LogEvent(logInfoFile, asset.ID, ecomm.BidEvent, keyWords, t, "Highest bid increased to "+amount, 0)

	bid.BidAmount = eventPayload.BidAmount.String()
	bid.Bidder = eventPayload.Bidder
	bid.AssetID = asset.ID

	payloadJSON, _ := json.Marshal(bid)
	wrapper := ecomm.EventWrapper{Type: "Bid", Result: payloadJSON}
	payload, _ := json.Marshal(wrapper)

	ccsvc.Publish(ecomm.BidEvent, payload)
	return nil
}

// Smart Contract handler
func handleBidTooLowEvent(eventPayload ecomm.BidTooLow, bid ecomm.Bid, t time.Time) error {
	log.Printf("[%s] BidTooLow Event", strings.ToUpper(bid.Platform))

	amount := new(big.Int).Div(eventPayload.BidAmount, ecomm.DecimalB).String()
	asset, _ := assetClient.GetAsset(eventPayload.Id)
	keyWords := fmt.Sprintf("%s_%s_%s", bid.Platform, eventPayload.Bidder.String()[36:], amount)
	ecomm.LogEvent(logInfoFile, asset.ID, ecomm.BidEvent, keyWords, t, "Bid too low with amount "+amount, 0)

	bid.BidAmount = eventPayload.BidAmount.String()
	bid.Bidder = eventPayload.Bidder
	bid.AssetID = asset.ID

	payloadJSON, _ := json.Marshal(bid)
	wrapper := ecomm.EventWrapper{Type: "Bid", Result: payloadJSON}
	payload, _ := json.Marshal(wrapper)

	// asset, _ := assetClient.GetAsset(eventPayload.Id)
	// auction, _ := assetClient.GetAuction(asset.PendingAuctionID)
	// fmt.Println("find auction in new bid: ", auction.ID, "status: ", auction.Status)

	ccsvc.Publish(ecomm.BidEvent, payload)
	return nil
}

// Smart Contract handler
func handleNewBidHashEvent(eventPayload ecomm.NewBidHash, bidHash ecomm.BidHash, t time.Time) error {
	log.Printf("[%s] NewBidHash Event", strings.ToUpper(bidHash.Platform))
	//log.Print("[Relayer] bidHash", bidHash.BidHash)

	asset, _ := assetClient.GetAsset(eventPayload.Id)

	bidHash.BidHash = eventPayload.BidHash
	bidHash.Bidder = eventPayload.Bidder
	bidHash.AssetID = eventPayload.Id

	//log.Printf("Hash of uint(4): %s", hex.EncodeToString(bidHash.BidHash[:]))

	keyWords := fmt.Sprintf("%s_%s_%s", bidHash.Platform, eventPayload.Bidder.String()[36:], hex.EncodeToString(bidHash.BidHash[:])[60:])
	ecomm.LogEvent(logInfoFile, asset.ID, ecomm.BidHashEvent, keyWords, t, "", 0)

	payloadJSON, _ := json.Marshal(bidHash)
	wrapper := ecomm.EventWrapper{Type: "BidHash", Result: payloadJSON}
	payload, _ := json.Marshal(wrapper)

	// asset, _ := assetClient.GetAsset(eventPayload.Id)
	// auction, _ := assetClient.GetAuction(asset.PendingAuctionID)
	// fmt.Println("find auction in new bid: ", auction.ID, "status: ", auction.Status)

	ccsvc.Publish(ecomm.BidEvent, payload)
	return nil
}

func handleRevealAuctionEvent(eventPayload ecomm.RevealAuction, t time.Time) error {
	log.Printf("[QUO/ETH] RevealAuction Event")

	auction, _ := assetClient.GetAuction(int(eventPayload.AuctionId.Int64()))
	ecomm.LogEvent(logInfoFile, auction.AssetID, ecomm.AuctionRevealingEvent, auction.AucType, t, "", 0)

	payloadJSON, _ := json.Marshal(auction)
	wrapper := ecomm.EventWrapper{Type: "Reveal", Result: payloadJSON}
	payload, _ := json.Marshal(wrapper)

	ccsvc.Publish(ecomm.AuctionRevealingEvent, payload)

	return nil
}

func handleWithdrawBidEvent(eventPayload ecomm.WithdrawBid, bid ecomm.Bid, t time.Time) error {
	log.Printf("[%s] WithdrawBid Event", strings.ToUpper(bid.Platform))

	amount := new(big.Int).Div(eventPayload.Amount, ecomm.DecimalB).String()
	auction, _ := assetClient.GetAuction(int(eventPayload.AuctionId.Int64()))
	keyWords := fmt.Sprintf("%s_%s_%s", auction.AucType, bid.Platform, eventPayload.Bidder.String()[36:])
	//eventID := eventPayload.Id + "_" + bid.Platform + "_" + eventPayload.Bidder.String()[36:]
	ecomm.LogEvent(logInfoFile, "NA", ecomm.WithdrawEvent, keyWords, t, "Withdraw: "+amount, 0)

	bid.BidAmount = amount
	bid.Bidder = eventPayload.Bidder
	bid.AssetID = eventPayload.Id
	bid.AuctionID = auction.AuctionID

	//ecomm.UpdateLog(logInfoFile, ecomm.WithdrawEvent, eventID, "", 0, "Withdraw: MDAI "+amount)

	payloadJSON, _ := json.Marshal(bid)
	wrapper := ecomm.EventWrapper{Type: "Withdraw", Result: payloadJSON}
	payload, _ := json.Marshal(wrapper)

	ccsvc.Publish(ecomm.WithdrawEvent, payload)
	return nil
}

func handleDecisionMadeEvent(eventPayload ecomm.DecisionMade, t time.Time) error {
	payload := eventPayload.JsonString
	//cclib.LogEventToFile(logInfoFile, ecomm.RelayerDetectedEvent, []byte(eventPayload.JsonString), t, timeInfoFile)

	var result ecomm.AuctionResult

	err := json.Unmarshal([]byte(payload), &result)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	auction, err := assetClient.GetAuction(result.AuctionID)
	check(err)

	log.Printf("[%s] Decesion Made Event", strings.ToUpper(result.Platform))
	//bidT.From.String()[36:]
	proceed := eventPayload.Prcd
	if proceed {
		ecomm.LogEvent(logInfoFile, auction.AssetID, ecomm.CommitAuctionResultEvent, auction.HighestBidder[36:], t, "", 0)
		payloadJSON, _ := json.Marshal(result)
		wrapper := ecomm.EventWrapper{Type: "Commit", Result: payloadJSON}
		payload_, _ := json.Marshal(wrapper)

		ccsvc.Publish(ecomm.CommitAuctionResultEvent, payload_)
	} else {
		ecomm.LogEvent(logInfoFile, auction.AssetID, ecomm.AbortAuctionResultEvent, "", t, "", 0)
		ccsvc.Publish(ecomm.AbortAuctionResultEvent, []byte(payload))
	}

	t = time.Now()
	ecomm.LogEvent(logInfoFile, auction.AssetID, ecomm.FinAuctionEvent, "", t, "", 0)

	_, err = assetClient.FinAuction(result, proceed)
	check(err)

	return nil
}

func handleRateAuctionEvent(eventPayload ecomm.RateAuction, t time.Time) error {
	auction, err := assetClient.GetAuction(int(eventPayload.AuctionId.Int64()))
	check(err)

	ecomm.LogEvent(logInfoFile, auction.AssetID, ecomm.ProvideFeedbackEvent, auction.HighestBidder[36:], t, "", 0)

	payloadJSON, _ := json.Marshal(auction)
	wrapper := ecomm.EventWrapper{Type: "Feedback", Result: payloadJSON}
	payload, _ := json.Marshal(wrapper)

	ccsvc.Publish(ecomm.ProvideFeedbackEvent, payload)
	return nil
}

func smartContractEvent(eventPayload []byte) {
	t := time.Now()
	var wrapper ecomm.EventWrapper
	var event, keyWords, assetId string
	err := json.Unmarshal([]byte(eventPayload), &wrapper)
	check(err)

	switch wrapper.Type {
	case "Bid":
		var bid ecomm.Bid
		err = json.Unmarshal(wrapper.Result, &bid)
		check(err)
		//ecomm.Decimal
		event = ecomm.BidEvent
		amount := new(big.Int)
		amount.SetString(bid.BidAmount, 10)
		amount.Div(amount, ecomm.DecimalB)
		assetId = bid.AssetID
		//eventID = bid.AssetID + "_" + bid.Platform + "_" + bid.Bidder.String()[36:]
		keyWords = fmt.Sprintf("%s_%s_%s", bid.Platform, bid.Bidder.String()[36:], amount)

		//fmt.Printf("Received Asset: %+v\n", asset)
	case "Withdraw":
		var bid ecomm.Bid
		err = json.Unmarshal(wrapper.Result, &bid)
		check(err)

		assetId = "NA"
		event = ecomm.WithdrawEvent
		auction, _ := assetClient.GetAuction(bid.AuctionID)
		keyWords = fmt.Sprintf("%s_%s_%s", auction.AucType, bid.Platform, bid.Bidder.String()[36:])

	case "BidHash":
		var bidHash ecomm.BidHash
		err = json.Unmarshal(wrapper.Result, &bidHash)
		check(err)
		//log.Println("[kafka] bidhash: ", bidHash.BidHash)
		assetId = bidHash.AssetID
		event = ecomm.BidHashEvent
		keyWords = fmt.Sprintf("%s_%s_%s", bidHash.Platform, bidHash.Bidder.String()[36:], hex.EncodeToString(bidHash.BidHash[:])[60:])

	case "Reveal":
		var auction ecomm.Auction
		err = json.Unmarshal(wrapper.Result, &auction)
		check(err)

		event = ecomm.AuctionRevealingEvent
		keyWords = auction.AucType
	case "Commit":
		var result ecomm.AuctionResult
		err = json.Unmarshal(wrapper.Result, &result)
		check(err)

		event = ecomm.CommitAuctionResultEvent
		auction, _ := assetClient.GetAuction(result.AuctionID)
		assetId = auction.AssetID
		keyWords = auction.HighestBidder[36:]

	case "Feedback":
		var auction ecomm.Auction
		err = json.Unmarshal(wrapper.Result, &auction)
		check(err)

		assetId = auction.AssetID
		event = ecomm.ProvideFeedbackEvent
		keyWords = auction.HighestBidder[36:]

	default:
		fmt.Printf("Unknown type: %s\n", wrapper.Type)
	}

	ecomm.LogEvent(logInfoFile, assetId, event, keyWords, t, "", 0)
}
