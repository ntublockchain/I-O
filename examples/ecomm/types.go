package ecomm

import (
	"encoding/json"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

type UserInfo struct {
	UserID  string
	Address common.Address
	KeyFile string
}

// may not need
type AuctionInfo struct {
	AuctionID int
	AucType   string

	AssetID string
	Owner   common.Address

	EthAddr common.Address
	QuoAddr common.Address
}

type AucConractInfo struct {
	AucType string
	Owner   common.Address
	EthAddr common.Address
	QuoAddr common.Address
}

type ContractInfo struct {
	FabricTokenName string
	EthERC20        common.Address
	QuoERC20        common.Address

	EnglishAuction AucConractInfo
	//DutchAuction   AucConractInfo
	Cb1pAuction AucConractInfo
	//Cb2pAuction    AucConractInfo
}

// type EnglishAuctionInfo struct {
// 	Owner   common.Address
// 	EthAddr common.Address
// 	QuoAddr common.Address
// }

// type ClosedBidAuctionInfo struct {
// 	Owner   common.Address
// 	EthAddr common.Address
// 	QuoAddr common.Address
// }

type Asset struct {
	ID               string `json:"id"`
	Owner            string `json:"owner"`
	PendingAuctionID int    `json:"pendingAuctionId"`
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

type Bid struct {
	Bidder      common.Address `json:"bidder"`
	BidAmount   string         `json:"bidAmount"`
	AuctionAddr common.Address `json:"auctionAddr"`

	Platform string `json:"platform"`

	AuctionID int    `json:"auctionID"`
	AssetID   string `json:"assetID"`
}

type BidHash struct {
	Bidder      common.Address `json:"bidder"`
	BidHash     [32]byte       `json:"bidHash"`
	AuctionAddr common.Address `json:"auctionAddr"`

	Platform string `json:"platform"`

	AuctionID int    `json:"auctionID"`
	AssetID   string `json:"assetID"`
}

type Tx struct {
	Platform string
	Type     string
	Hash     common.Hash
}

func (args *StartAuctionArgs) Hash() []byte {
	h := sha3.New256()

	h.Write([]byte(args.AssetID))
	h.Write([]byte(args.EthAddr))
	h.Write([]byte(args.QuorumAddr))

	return h.Sum(nil)
}

type AuctionResult struct {
	Platform    string
	AuctionID   int
	AuctionAddr string

	HighestBid    string
	HighestBidder string

	Signature []byte
}

type SignedAuctionResult struct {
	AuctionResult
	Signature []byte
}

type EventWrapper struct {
	Type   string          `json:"type"`
	Result json.RawMessage `json:"result"`
}

// type ContractChecker interface {
// 	HighestBid(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error)
// 	HighestBidder(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error)
// }

type AuctionContract interface {
	// Common across most auction types
	//CreateAuction(assetId string) error
	//Bid(opts *bind.TransactOpts, auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error) // Flexible input
	Withdraw(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error)
	CloseAuction(opts *bind.TransactOpts, auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) // Might be type-specific
	Abort(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error)
	Commit(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error)
	ProvideFeedback(opts *bind.TransactOpts, auctionId *big.Int, _score *big.Int, _feedback string) (*types.Transaction, error)

	HighestBid(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error)
	HighestBidder(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error)
}

type AuctionContractCloseBid interface {
	// Common across most auction types
	//CreateAuction(assetId string) error
	Bid(opts *bind.TransactOpts, auctionId *big.Int, bidHash [32]byte) (*types.Transaction, error)
	RevealAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error)
	Reveal(opts *bind.TransactOpts, auctionId *big.Int, bidAmount *big.Int) (*types.Transaction, error)
	Withdraw(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error)
	//	CloseAuction(opts *bind.TransactOpts, auctionId *big.Int, not_winner_platform bool) (*types.Transaction, error) // Might be type-specific
	Abort(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error)
	Commit(opts *bind.TransactOpts, auctionId *big.Int, jsonString string) (*types.Transaction, error)
	//ProvideFeedback(opts *bind.TransactOpts, auctionId *big.Int, _score *big.Int, _feedback string) (*types.Transaction, error)
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

type CloseAuctionArgs struct {
	AuctionID int `json:"auctionId"`

	HighestBid         string `json:"highestBid"`
	HighestBidder      string `json:"highestBidder"`
	HighestBidPlatform string `json:"highestBidPlatform"`
}

type AssetAddingEventPayload struct {
	AssetID string `json:"assetId"`
	AucType string `json:"aucType"`
}

type StartAuctionEventPayload struct {
	ID      int    `json:"id"`
	AucType string `json:"aucType"`
	Owner   string `json:"owner"`
}

// Struct used as input for creating new Auction
type StartAuctionArgs struct {
	AssetID    string
	AucType    string
	EthAddr    string
	QuorumAddr string

	Signature []byte
}

func VerifySignature(hash, signature []byte, addr string) bool {
	pubkey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return false
	}

	if addr == crypto.PubkeyToAddress(*pubkey).Hex() {
		return true
	}
	return false
}
