package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
)

// Use key 1 as default auctioneer
func create(asset_name string, auc_type string, usr_name string) {
	t := time.Now()
	//fmt.Println("Auc type:", auc_type)

	if !stringInSlice(support_auc_types, auc_type) {
		log.Println("[fabric] not support auction type")
		return
	}
	log.Println("[fabric] Adding asset")
	auc_key := load_auctioneer(usr_name)
	aucT, _ := cclib.NewTransactor(auc_key, password)
	_, err := assetClient.AddAsset(asset_name, aucT.From.Hex(), auc_type)
	check(err)

	ecomm.LogEvent(logInfoFile, asset_name, ecomm.AssetAddingEvent, auc_type, t, "", 0)
}

func reveal(auctionID int) {
	t := time.Now()

	a, err := assetClient.GetAuction(auctionID)
	check(err)
	ecomm.LogEvent(logInfoFile, a.AssetID, ecomm.EndClosedBidEvent, a.AucType, t, "", 0)

	if a.Status != "open" {
		err = fmt.Errorf("auction status error")
		check(err)
	}

	log.Println("[fabric] End ClosedBid")
	_, err = assetClient.EndClosedBid(auctionID)
	check(err)
}

func cancel(auctionID int) {
	t := time.Now()

	a, err := assetClient.GetAuction(auctionID)
	check(err)

	if a.Status != "open" {
		err = fmt.Errorf("auction status error")
		check(err)
	}

	log.Println("[fabric] Cancel auction")
	_, err = assetClient.CancelAuction(auctionID)
	check(err)

	ecomm.LogEvent(logInfoFile, a.AssetID, ecomm.CancelAuctionEvent, "", t, "", 0)
}

func close(auctionID int) {
	t := time.Now()

	a, err := assetClient.GetAuction(auctionID)
	check(err)

	if (a.Status != "open") && (a.Status != "reveal") {
		err = fmt.Errorf("auction status error")
		check(err)
	}

	log.Println("[fabric] Determine Auction Winner")
	_, err = assetClient.DetermineWinner(auctionID)
	check(err)

	//payload, _ := json.Marshal(a)
	ecomm.LogEvent(logInfoFile, a.AssetID, ecomm.DetermineWinnerEvent, "", t, "", 0)
	//cclib.LogEventToFile(logInfoFile, ecomm.AuctionClosingEvent, payload, t, timeInfoFile)
}
