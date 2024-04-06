package ecomm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	AssetAddingEvent      = "auctioneer.add_asset"
	AuctionStartingEvent  = "relayer.start_auction"
	AuctionClosingEvent   = "relayer.close_auction"
	AuctionRevealingEvent = "relayer.reveal_auction"

	DetermineWinnerEvent = "auctioneer.determine_winner"
	CancelAuctionEvent   = "auctioneer.cancel_auction"
	EndClosedBidEvent    = "auctioneer.end_closed_bid"
	FinAuctionEvent      = "auctioneer.fin_auction"

	AuctionStateUpdatingEvent = "relayer.update_auction_state"

	BidEvent                 = "bidder.bid"
	BidHashEvent             = "bidder.conceal_bid"
	RevealBidEvent           = "bidder.reveal_bid"
	WithdrawEvent            = "bidder.withdraw"
	CommitAuctionResultEvent = "bidder.cmt_result"
	AbortAuctionResultEvent  = "bidder.abt_result"

	TransactionMinedEvent = "eth_quo.tx_mined"
	//SignAuctionResultEvent = "eth_quo.signed_result"

	ProvideFeedbackEvent = "eth_quo.provide_feedback"

	KafkaReceivedEvent   = "kafka.received"
	RelayerDetectedEvent = "relayer.detected"
	// end with BiddingAuctionEvent
)

type HighestBidIncreased struct {
	AuctionId   *big.Int
	Id          string
	Bidder      common.Address
	BidAmount   *big.Int
	AuctionType string
	Raw         types.Log // Blockchain specific contextual infos
}

type NewBidHash struct {
	AuctionId   *big.Int
	Id          string
	Bidder      common.Address
	BidHash     [32]byte
	AuctionType string
	Raw         types.Log // Blockchain specific contextual infos
}

type RevealAuction struct {
	AuctionId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

type BidTooLow struct {
	AuctionId   *big.Int
	Id          string
	Bidder      common.Address
	BidAmount   *big.Int
	HighestBid  *big.Int
	AuctionType string
	Raw         types.Log // Blockchain specific contextual infos
}

type WithdrawBid struct {
	AuctionId *big.Int
	Id        string
	Bidder    common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

type DecisionMade struct {
	AuctionId  *big.Int
	Winner     common.Address
	Amount     *big.Int
	Id         string
	Prcd       bool
	JsonString string
	Raw        types.Log // Blockchain specific contextual infos
}

type AwaitResponse struct {
	AuctionId *big.Int
	Winner    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

type RateAuction struct {
	AuctionId *big.Int
	Id        string
	Rating    *big.Int
	Review    string
	Raw       types.Log // Blockchain specific contextual infos
}
