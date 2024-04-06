// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.18;

interface IERC20 {
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
}

contract DutchAuction {
    // Parameters of the auction
    uint public immutable startingPrice;
    uint public immutable startAt;
    uint public immutable discountRate;

    // Outcome of the auction.
    address public winningBidder;
    uint public winningBid;

    // Set to true at the end, disallows any change.
    // By default initialized to `false`.

    string public status;
    address public owner;

    string public asset_id;

    int32[] private score;

    // Events that will be emitted on changes.
    event WinningBid(string id, address bidder, uint amount);
    event DecisionMade(address winner, uint amount, string id, bool prcd, string jsonString);
    event WaitResponse(address winner);
    event RateAuction(string id, int rating);

    IERC20 public token;

    constructor(IERC20 _token, string memory _asset_id, uint _startingPrice, uint _discountRate) {
        //beneficiary = payable(msg.sender);
        token = _token;
        status = "open";
        owner = msg.sender;
        asset_id = _asset_id;
        startingPrice = _startingPrice;
        startAt = block.timestamp;
        discountRate = _discountRate;
    }

    function getPrice() public view returns (uint) {
        uint timeElapsed = block.timestamp - startAt;
        uint discount = discountRate * timeElapsed;
        return startingPrice - discount;
    }

    function bid(uint bidAmount) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("open")), "Contract not in OPEN status");

        uint price = getPrice();
        require(bidAmount >= price, "Bid is too low");

        // Attempt to transfer the tokens from the bidder to the contract
        bool transferSuccessful = token.transferFrom(msg.sender, address(this), bidAmount);
        
        // Check that the token transfer was successful
        require(transferSuccessful, "Token transfer failed.");
        
        // Update the highest bid and highest bidder
        winningBidder = msg.sender;
        winningBid = bidAmount;

        status = "closed";
        
        emit WinningBid(asset_id, msg.sender, bidAmount);
    }

    function abort(string memory jsonString) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("ending")), "Contract not in ENDING status");

        require(msg.sender == winningBidder, "Not authorized access!");

        status = "closing";
        emit DecisionMade(winningBidder, winningBid, asset_id, false, jsonString);

        if(winningBid > 0) {
            bool refundSuccessful = token.transferFrom(address(this), winningBidder, winningBid);
            require(refundSuccessful, "Refund failed.");
        }

        winningBidder = address(0);
        winningBid = 0;
    }

    function commit(string memory jsonString) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("ending")), "Contract not in ENDING status");
        // For testing only
        require(msg.sender == winningBidder, "Not authorized access!");
        
        status = "closing";

        token.transferFrom(address(this), owner, winningBid);
        emit DecisionMade(winningBidder, winningBid, asset_id, true, jsonString);
    }

    function provide_feedback(int32 _score) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("closing")), "Contract not in CLOSING status");
        // For testing only
        require(msg.sender == winningBidder, "Not authorized access!");

        score.push(_score);

        emit RateAuction(asset_id, _score);
        status = "closed";
    }

    function checkAverageFeedback() public view returns (int) {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("closed")), "Contract not in CLOSED status");

        int total = 0;
        for(uint i=0;i<score.length;i++) {
            total += score[i];
        }

        // solidity does not support floats, so we multiply the rating by 100 to achieve accuracy up to two decimals (the user's client will have to divide the result by 100)
        return (100*total/int(score.length));
    }
}
