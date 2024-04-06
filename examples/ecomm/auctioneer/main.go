package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/cclib"
	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

const (
	password = "password"

	userInfoFile     = "../tmp/user_info.json"
	contractInfoFile = "../tmp/contract_info.json"
	logInfoFile      = "../tmp/log.csv"
	auctionInfoFile  = "../tmp/auction_info.json"
	timeInfoFile     = "../tmp/timer"
)

var (
	assetClient *ecomm.AssetClient

	aucT              *bind.TransactOpts
	usr_name          = "Auctioneer 1"
	support_auc_types = []string{"english", "dutch", "cb1p", "cb2p"}
	//auc_key  = "../../keys/key1"
)

func main() {
	assetClient = ecomm.NewAssetClient()

	command := flag.String("c", "", "command")
	asset := flag.String("ast", "", "Asset name")
	id := flag.String("id", "", "Auction ID")
	auc_type := flag.String("t", "", "Auction type")

	flag.StringVar(&usr_name, "usr", usr_name, "Load User/auctioneer Information")
	flag.Parse()

	fmt.Println("Load auctioneer: ", usr_name)
	auc_key := load_auctioneer(usr_name)
	aucT, _ = cclib.NewTransactor(auc_key, password)

	switch *command {
	case "create":
		create(*asset, *auc_type)
	case "reveal":
		id_, _ := strconv.Atoi(*id)
		reveal(id_)
	case "close":
		id_, _ := strconv.Atoi(*id)
		close(id_)
	case "cancel":
		id_, _ := strconv.Atoi(*id)
		cancel(id_)
	case "check":
		id_, _ := strconv.Atoi(*id)
		check_status(id_)
	default:
		fmt.Println("command not found")
	}
}

// Use key 1 as default auctioneer
func create(asset_name string, auc_type string) {
	t := time.Now()
	//fmt.Println("Auc type:", auc_type)

	if !stringInSlice(support_auc_types, auc_type) {
		log.Println("[fabric] not support auction type")
		return
	}

	log.Println("[fabric] Adding asset")
	_, err := assetClient.AddAsset(asset_name, aucT.From.Hex(), auc_type)
	check(err)

	ecomm.LogEvent(logInfoFile, asset_name, ecomm.AssetAddingEvent, auc_type, t, "", 0)
}

func reveal(auctionID int) {
	t := time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	a, err := assetClient.GetAuction(auctionID)
	check(err)

	if a.Status != "open" {
		err = fmt.Errorf("auction status error")
		check(err)
	}

	log.Println("[fabric] Reveal auction")
	//_, err = assetClient.CloseAuction(auctionID)
	check(err)

	payload, _ := json.Marshal(a)
	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.AuctionClosingEvent, payload, t, timeInfoFile)

	//@reset
	cclib.LastEventTimestamp.Set(t, timeInfoFile)
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

	ecomm.LogEvent(logInfoFile, ecomm.CancelAuctionEvent, a.AssetID, a.AucType, t, "", 0)
}

func close(auctionID int) {
	t := time.Now()
	cclib.LastEventTimestamp.Set(t, timeInfoFile)

	a, err := assetClient.GetAuction(auctionID)
	check(err)

	if a.Status != "open" {
		err = fmt.Errorf("auction status error")
		check(err)
	}

	log.Println("[fabric] Conclude auction")
	//_, err = assetClient.CloseAuction(auctionID)
	check(err)

	payload, _ := json.Marshal(a)
	t = time.Now()
	cclib.LogEventToFile(logInfoFile, ecomm.AuctionClosingEvent, payload, t, timeInfoFile)

	//@reset
	cclib.LastEventTimestamp.Set(t, timeInfoFile)
}

func check_status(auctionID int) {
	a, err := assetClient.GetAuction(auctionID)
	check(err)

	fmt.Println("auction ID:", auctionID, "AssetID:", a.AssetID, "Status: ", a.Status)

}

func load_auctioneer(name string) string {
	users, err := ecomm.ReadUsersFromFile(userInfoFile)
	check(err)

	for _, user := range users {
		if name == user.UserID {
			return user.KeyFile
		}
	}

	return "../../keys/key1"
}

// stringInSlice checks if a string exists in a slice of strings.
func stringInSlice(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true // Found the target string in the slice
		}
	}
	return false // Target string not found in the slice
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
