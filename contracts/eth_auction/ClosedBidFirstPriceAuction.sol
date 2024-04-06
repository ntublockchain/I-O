// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.18;

interface IERC20 {
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
}

contract ClosedBidFirstPriceAuction {
    address public owner;

    // Allowed withdrawals of previous bids
    mapping(address => uint) pendingReturns;

    // Current state of the auction.
    mapping (uint => mapping(address => bytes32)) bidHashes;
    address[] public highestBidder;
    uint[] public highestBid;
    string[] public status;
    string[] public asset_id;

    // feedback
    bytes32[] private feedback;
    int[] private score;

    // Events that will be emitted on changes.
    event NewBidHash(uint auction, string id, address bidder, bytes32 bidHash);
    event HighestBidIncreased(uint auction, string id, address bidder, uint amount);
    event RevealAuction(uint auction);
    event WithdrawBid(uint auction, string id, address bidder, uint amount);
    event DecisionMade(uint auction, address winner, uint amount, string id, bool prcd, string jsonString);
    event AwaitResponse(uint auction, address winner);
    event RateAuction(uint auction, string id, int rating, bytes32 review);

    IERC20 public immutable token;

    constructor(address _token) {
        token = IERC20(_token);
        owner = msg.sender;
    }

    function create(string memory _asset_id) public {
        require(msg.sender == owner, "Only owner can create new auction");

        highestBidder.push(address(0));
        highestBid.push(0);
        status.push("open");
        asset_id.push(_asset_id);
        feedback.push(bytes32(0));
        score.push(0);
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

    function closeAuction(uint auctionId, bool not_winner_platform) public {
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

        pendingReturns[owner] += highestBid[auctionId];
        emit DecisionMade(auctionId, highestBidder[auctionId], highestBid[auctionId], asset_id[auctionId], true, jsonString);
    }

    function provide_feedback(uint auctionId, int _score, bytes32 _feedback) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status[auctionId])) == keccak256(abi.encodePacked("closing")), "Contract not in CLOSING status");
        // For testing only
        require(msg.sender == highestBidder[auctionId], "Not authorized access!");

        score[auctionId] = _score;
        feedback[auctionId] = _feedback;

        emit RateAuction(auctionId, asset_id[auctionId], _score, _feedback);
        status[auctionId] = "closed";
    }

    function checkAverageScore() public view returns (int) {
        int total = 0;
        for(uint i=0;i<score.length;i++) {
            // Use hash to check status
            if(keccak256(abi.encodePacked(status[i])) == keccak256(abi.encodePacked("closed"))) {
                total += score[i];
            }
        }

        // solidity does not support floats, so we multiply the rating by 100 to achieve accuracy up to two decimals (the user's client will have to divide the result by 100)
        return (100*total/int(score.length));
    }
}
