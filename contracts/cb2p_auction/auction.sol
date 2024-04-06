// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.18;

interface IERC20 {
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
}

contract ClosedBidSecondPriceAuction {
    address public owner;

    // Allowed withdrawals of previous bids
    mapping(address => uint) pendingReturns;

    // Current state of the auction.
    mapping (uint => mapping(address => bytes32)) bidHashes;
    // uint: AuctionID
    mapping(uint => address) public highestBidder;
    mapping(uint => uint) public secondHighestBid;
    mapping(uint => uint) public highestBid;
    mapping(uint => string) public status;
    mapping(uint => string) public asset_id;
    mapping(uint => string) public asset_owner;

    // Feedback
    mapping(string => string[]) private feedback;
    mapping(string => int[]) private score;

    // Events that will be emitted on changes.
    event NewBidHash(uint auction, string id, address bidder, bytes32 bidHash);
    event HighestBidIncreased(uint auction, string id, address bidder, uint amount);
    event RevealAuction(uint auction);
    event WithdrawBid(uint auction, string id, address bidder, uint amount);
    event DecisionMade(uint auction, address winner, uint amount, string id, bool prcd, string jsonString);
    event AwaitResponse(uint auction, address winner);
    event RateAuction(uint auctionId, string id, int rating, string review);

    IERC20 public immutable token;

    constructor(address _token) {
        token = IERC20(_token);
        owner = msg.sender;
    }

    function create(uint _auction_id, string memory _asset_id, string memory _asset_owner) public {
        require(msg.sender == owner, "Only owner can create new auction");

        // Initialize the auction with default values
        highestBidder[_auction_id] = address(0);
        highestBid[_auction_id] = 0;
        status[_auction_id] = "open";
        asset_id[_auction_id] = _asset_id;
        asset_owner[_auction_id] = _asset_owner;

        // feedback and score are related to users, not auctions, so might not be set here
    }

    function bid(uint auctionId, bytes32 bidHash) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status[auctionId])) == keccak256(abi.encodePacked("open")), "Contract not in OPEN status");

        // Update the highest bid and highest bidder
        bidHashes[auctionId][msg.sender] = bidHash;
        emit NewBidHash(auctionId,asset_id[auctionId], msg.sender, bidHash);
    }

    function revealAuction(uint auctionId) public {
        require(msg.sender == owner, "Only owner can change contract's status");
        // Use hash to check status
        require(keccak256(abi.encodePacked(status[auctionId])) == keccak256(abi.encodePacked("open")), "Contract not in OPEN status");

        status[auctionId] = "reveal";

        emit RevealAuction(auctionId);
    }

    function reveal(uint auctionId, uint bidAmount) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status[auctionId])) == keccak256(abi.encodePacked("reveal")), "Contract not in OPEN status");

        // Check that the bid is higher than the current highest bid
        require(bidAmount > highestBid[auctionId], "There already is a higher bid.");

        // Check that the bid amount corresponds to the hash 
        require(keccak256(abi.encodePacked(bidAmount)) == bidHashes[auctionId][msg.sender]);

        // Attempt to transfer the tokens from the bidder to the contract
        bool transferSuccessful = token.transferFrom(msg.sender, address(this), bidAmount);
        
        // Check that the token transfer was successful
        require(transferSuccessful, "Token transfer failed.");

        // If there was a previous bid, allow the previous highest bidder to withdraw their bid
         if (highestBid[auctionId] != 0) {
            pendingReturns[highestBidder[auctionId]] += highestBid[auctionId];
        }
        
        // Update the highest bid and highest bidder
        highestBidder[auctionId] = msg.sender;
        secondHighestBid[auctionId] = highestBid[auctionId];
        highestBid[auctionId] = bidAmount;
        emit HighestBidIncreased(auctionId, asset_id[auctionId], msg.sender, bidAmount);
    }


    // Withdraw a bid that was overbid.
    function withdraw(uint auctionId) public returns (bool) {
        require(keccak256(abi.encodePacked(status[auctionId])) != keccak256(abi.encodePacked("open")), "Contract in OPEN status");

        uint amount = pendingReturns[msg.sender];
        if (amount > 0) {
            pendingReturns[msg.sender] = 0;

            if (!token.transferFrom(address(this), msg.sender, amount)) {
                pendingReturns[msg.sender] = amount;
                return false;
            }
        }
        emit WithdrawBid(auctionId, asset_id[auctionId], msg.sender, amount);
        return true;
    }

    function closeAuction(uint auctionId, bool not_winner_platform, uint secondHighestPrice) public {
        require(msg.sender == owner, "Only owner can change contract's status");
        // Use hash to check status
        require(keccak256(abi.encodePacked(status[auctionId])) == keccak256(abi.encodePacked("open")), "Contract not in OPEN status");

        status[auctionId] = "ending";
        if (not_winner_platform || highestBid[auctionId] == 0){
            status[auctionId] = "closed";

            pendingReturns[highestBidder[auctionId]] += highestBid[auctionId];
            highestBidder[auctionId] = address(0);
            highestBid[auctionId] = 0;

            return;
        }

        if(secondHighestPrice > secondHighestBid[auctionId]) {
            secondHighestBid[auctionId] = secondHighestPrice;
        }
        emit AwaitResponse(auctionId, highestBidder[auctionId]);
    }

    function abort(uint auctionId, string memory jsonString) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status[auctionId])) == keccak256(abi.encodePacked("ending")), "Contract not in ENDING status");

        require(msg.sender == highestBidder[auctionId], "Not authorized access!");

        status[auctionId] = "closing";
        emit DecisionMade(auctionId, highestBidder[auctionId], highestBid[auctionId], asset_id[auctionId], false, jsonString);

        pendingReturns[highestBidder[auctionId]] += highestBid[auctionId];
        highestBidder[auctionId] = address(0);
        highestBid[auctionId] = 0;
    }

    function commit(uint auctionId, string memory jsonString) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status[auctionId])) == keccak256(abi.encodePacked("ending")), "Contract not in ENDING status");
        // For testing only
        require(msg.sender == highestBidder[auctionId], "Not authorized access!");
        
        status[auctionId] = "closing";
        if(highestBid[auctionId] > secondHighestBid[auctionId]) {
            pendingReturns[highestBidder[auctionId]] += highestBid[auctionId] - secondHighestBid[auctionId];
        }
        pendingReturns[owner] += secondHighestBid[auctionId];
        emit DecisionMade(auctionId, highestBidder[auctionId], highestBid[auctionId], asset_id[auctionId], true, jsonString);
    }

    function provide_feedback(uint auctionId, int _score, string memory _feedback) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status[auctionId])) == keccak256(abi.encodePacked("closing")), "Contract not in CLOSING status");
        // For testing only
        require(msg.sender == highestBidder[auctionId], "Not authorized access!");

        score[asset_owner[auctionId]].push(_score);
        feedback[asset_owner[auctionId]].push(_feedback);

        emit RateAuction(auctionId, asset_id[auctionId], _score, _feedback);
        status[auctionId] = "closed";
    }

    function checkAverageScore(uint auctionId) public view returns (int) {
        int total = 0;
        uint l = score[asset_owner[auctionId]].length;
        for(uint i=0; i < l;i++) {
            total += score[asset_owner[auctionId]][i];
        }

        // solidity does not support floats, so we multiply the rating by 100 to achieve accuracy up to two decimals (the user's client will have to divide the result by 100)
        return (100*total/int(l));
    }


}
