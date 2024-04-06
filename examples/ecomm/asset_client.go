package ecomm

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type AssetClient struct {
	contract *gateway.Contract
}

// load identity
func NewAssetClient() *AssetClient {
	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environment variable: %v", err)
	}

	walletPath := "wallet"
	// remove any existing wallet from prior runs
	os.RemoveAll(walletPath)
	wallet, err := gateway.NewFileSystemWallet(walletPath)
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}

	ccpPath := filepath.Join(
		"../../../",
		"fabric-samples",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)

	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	return &AssetClient{contract: network.GetContract("asset")}
}

func (cc *AssetClient) GetCCID() string {
	return "asset"
}

func (cc *AssetClient) AddAsset(id, owner, auc_type string) ([]byte, error) {
	return cc.contract.SubmitTransaction("AddAsset", id, owner, auc_type)
}

func (cc *AssetClient) StartAuction(args StartAuctionArgs) ([]byte, error) {
	b, _ := json.Marshal(args)
	return cc.contract.SubmitTransaction("StartAuction", string(b))
}

func (cc *AssetClient) DetermineWinner(ID int) ([]byte, error) {
	IDStr := strconv.FormatInt(int64(ID), 10)
	return cc.contract.SubmitTransaction("DetermineWinner", IDStr)
}

func (cc *AssetClient) CloseAuction(args CloseAuctionArgs) ([]byte, error) {
	b, _ := json.Marshal(args)
	return cc.contract.SubmitTransaction("CloseAuction", string(b))
}

func (cc *AssetClient) EndClosedBid(ID int) ([]byte, error) {
	IDStr := strconv.FormatInt(int64(ID), 10)
	return cc.contract.SubmitTransaction("EndClosedBid", IDStr)
}

func (cc *AssetClient) CancelAuction(ID int) ([]byte, error) {
	IDStr := strconv.FormatInt(int64(ID), 10)
	return cc.contract.SubmitTransaction("CancelAuction", IDStr)
}

func (cc *AssetClient) FinAuction(args AuctionResult, prcd bool) ([]byte, error) {
	b, _ := json.Marshal(args)
	prcdStr := strconv.FormatBool(prcd)
	return cc.contract.SubmitTransaction("FinAuction", string(b), prcdStr)
}

func (cc *AssetClient) GetAsset(assetID string) (*Asset, error) {
	var asset Asset
	res, err := cc.contract.EvaluateTransaction("GetAsset", assetID)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(res, &asset)
	return &asset, err
}

func (cc *AssetClient) GetAuction(auctionID int) (*Auction, error) {
	var auction Auction
	res, err := cc.contract.EvaluateTransaction("GetAuction", strconv.Itoa(auctionID))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(res, &auction)
	return &auction, err
}

func (cc *AssetClient) GetLastAuctionID() (int, error) {
	res, err := cc.contract.EvaluateTransaction("GetLastAuctionID")
	if err != nil {
		return 0, err
	}
	var id int
	err = json.Unmarshal(res, &id)
	return id, err
}

func (cc *AssetClient) Register(eventID string) (fab.Registration, <-chan *fab.CCEvent, error) {
	return cc.contract.RegisterEvent(eventID)
}

func (cc *AssetClient) Unregister(reg fab.Registration) {
	cc.contract.Unregister(reg)
}
