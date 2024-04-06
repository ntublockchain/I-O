package fabric_asset

import (
	"strconv"

	"golang.org/x/crypto/sha3"
)

type AssetAddingEventPayload struct {
	AssetID string `json:"assetId"`
	AucType string `json:"aucType"`
}

// Struct used as input for creating new Auction
type StartAuctionArgs struct {
	AssetID    string
	AucType    string
	EthAddr    string
	QuorumAddr string

	Signature []byte
}

type CloseAuctionArgs struct {
	AuctionID int `json:"auctionId"`

	HighestBid         string `json:"highestBid"`
	HighestBidder      string `json:"highestBidder"`
	HighestBidPlatform string `json:"highestBidPlatform"`
}

type Auction struct {
	AuctionID  int    `json:"auctionId"`
	AssetID    string `json:"assetId"`
	AssetOwner string `json:"assetOwner"`

	AucType    string `json:"aucType"`
	EthAddr    string `json:"ethAddr"`
	QuorumAddr string `json:"quorumAddr"`

	Status string `json:"status"`

	HighestBid         string `json:"highestBid"`
	HighestBidder      string `json:"highestBidder"`
	HighestBidPlatform string `json:"highestBidPlatform"`
}

type StartAuctionEventPayload struct {
	ID      int    `json:"id"`
	AucType string `json:"aucType"`
	Owner   string `json:"owner"`
}

type Asset struct {
	ID               string
	Owner            string
	PendingAuctionID int
}

func (sa *StartAuctionArgs) Hash() []byte {
	h := sha3.New256()

	h.Write([]byte(sa.AssetID))
	h.Write([]byte(sa.EthAddr))
	h.Write([]byte(sa.QuorumAddr))

	return h.Sum(nil)
}

type AuctionResult struct {
	Platform    string
	AuctionID   int
	AuctionAddr string

	HighestBid    string
	HighestBidder string

	Signature []byte // acknowledged by bidder?
}

func (ar *AuctionResult) Hash() []byte {
	h := sha3.New256()

	h.Write([]byte(ar.Platform))
	h.Write([]byte(strconv.Itoa(ar.AuctionID)))
	h.Write([]byte(ar.AuctionAddr))

	h.Write([]byte(ar.HighestBid))
	h.Write([]byte(ar.HighestBidder))

	h.Write([]byte(""))

	return h.Sum(nil)
}

type CrossChainAuctionResult struct {
	AuctionResult
	Signatures [][]byte
}

type FinalizeAuctionArgs struct {
	AuctionID    int
	EthResult    CrossChainAuctionResult
	QuorumResult CrossChainAuctionResult
}
