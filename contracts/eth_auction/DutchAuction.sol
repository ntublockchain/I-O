// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.18;

interface IERC20 {
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
}

contract DutchAuction {
    address public owner;

    // Parameters of the auction
    uint[] public startingPrice;
    uint[] public startAt;
    uint[] public discountRate;

    // Allowed withdrawals of previous bids
    mapping(address => uint) pendingReturns;

    // Current state of the auction.
    address[] public winningBidder;
    uint[] public winningBid;
    string[] public status;
    string[] public asset_id;

    // feedback
    bytes32[] private feedback;
    int[] private score;

    // Events that will be emitted on changes.
    event BidReceived(uint auction, string id, address bidder, uint amount);
    event WithdrawBid(uint auction, string id, address bidder, uint amount);
    event DecisionMade(uint auction, address winner, uint amount, string id, bool prcd, string jsonString);
    event AwaitResponse(uint auction, address winner);
    event RateAuction(uint auction, string id, int rating, bytes32 review);

    IERC20 public immutable token;

    constructor(address _token) {
        token = IERC20(_token);
        owner = msg.sender;
    }

    function create(string memory _asset_id, uint _startingPrice, uint _discountRate) public {
        require(msg.sender == owner, "Only owner can create new auction");

        winningBidder.push(address(0));
        winningBid.push(0);
        status.push("open");
        asset_id.push(_asset_id);
        startingPrice.push(_startingPrice);
        startAt.push(block.timestamp);
        discountRate.push(_discountRate);
        feedback.push(bytes32(0));
        score.push(0);
    }

    function getPrice(uint auctionId) public view returns (uint) {
        uint timeElapsed = block.timestamp - startAt[auctionId];
        uint discount = discountRate[auctionId] * timeElapsed;
        return startingPrice[auctionId] - discount;
    }

    function bid(uint auctionId, uint bidAmount) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status[auctionId])) == keccak256(abi.encodePacked("open")), "Contract not in OPEN status");

        // only the first bid is relevant
        require(winningBid[auctionId] == 0, "Not the first bid");

        uint price = getPrice(auctionId);
        require(bidAmount >= price, "Bid is too low");

        // Attempt to transfer the tokens from the bidder to the contract
        bool transferSuccessful = token.transferFrom(msg.sender, address(this), bidAmount);
        
        // Check that the token transfer was successful
        require(transferSuccessful, "Token transfer failed.");
        
        // Update the winning bidder and close the auction
        winningBidder[auctionId] = msg.sender;
        winningBid[auctionId] = bidAmount;

        emit BidReceived(auctionId, asset_id[auctionId], msg.sender, bidAmount);
    }


    // Withdraw the winning bid
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
        if (not_winner_platform || winningBid[auctionId] == 0){
            status[auctionId] = "closed";

            pendingReturns[winningBidder[auctionId]] += winningBid[auctionId];
            winningBidder[auctionId] = address(0);
            winningBid[auctionId] = 0;

            return;
        }

        emit AwaitResponse(auctionId, winningBidder[auctionId]);
    }

    function abort(uint auctionId, string memory jsonString) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status[auctionId])) == keccak256(abi.encodePacked("ending")), "Contract not in ENDING status");

        require(msg.sender == winningBidder[auctionId], "Not authorized access!");

        status[auctionId] = "closing";
        emit DecisionMade(auctionId, winningBidder[auctionId], winningBid[auctionId], asset_id[auctionId], false, jsonString);

        pendingReturns[winningBidder[auctionId]] += winningBid[auctionId];
        winningBidder[auctionId] = address(0);
        winningBid[auctionId] = 0;
    }

    function commit(uint auctionId, string memory jsonString) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status[auctionId])) == keccak256(abi.encodePacked("ending")), "Contract not in ENDING status");
        // For testing only
        require(msg.sender == winningBidder[auctionId], "Not authorized access!");
        
        status[auctionId] = "closing";

        pendingReturns[owner] += winningBid[auctionId];
        emit DecisionMade(auctionId, winningBidder[auctionId], winningBid[auctionId], asset_id[auctionId], true, jsonString);
    }

    function provide_feedback(uint auctionId, int _score, bytes32 _feedback) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status[auctionId])) == keccak256(abi.encodePacked("closing")), "Contract not in CLOSING status");
        // For testing only
        require(msg.sender == winningBidder[auctionId], "Not authorized access!");

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
