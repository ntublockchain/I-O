// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.18;

interface IERC20 {
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
    function burn(address usr, uint wad) external;
}

contract ClosedBidSecondPriceAuction {

    // hashes of the submitters
    mapping (address => bytes32) bidHashes;

    // Current state of the auction.
    address public highestBidder;
    uint public highestBid;
    uint public secondHighestBid;

    // Set to true at the end, disallows any change.
    // By default initialized to `false`.

    string public status;
    address public owner;

    string public asset_id;

    int32[] private score;

    // Events that will be emitted on changes.
    event NewBidHash(string id, address bidder, bytes32 bidHash);
    event HighestBidIncreased(string id, address bidder, uint amount);
    event DecisionMade(address winner, uint amount, uint amount2, string id, bool prcd, string jsonString);
    event WaitResponse(address winner);
    event RateAuction(string id, int rating);

    IERC20 public token;

    constructor(IERC20 _token, string memory _asset_id) {
        //beneficiary = payable(msg.sender);
        token = _token;
        status = "open";
        owner = msg.sender;
        asset_id = _asset_id;
    }

    function bid(bytes32 bidHash) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("open")), "Contract not in OPEN status");
        
        // Update the highest bid and highest bidder
        bidHashes[msg.sender] = bidHash;
        emit NewBidHash(asset_id, msg.sender, bidHash);
    }

    function revealAuction() public {
        require(msg.sender == owner, "Only owner can change contract's status");
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("open")), "Contract not in OPEN status");

        status = "reveal";

        emit WaitResponse(highestBidder);
    }

    function reveal(uint bidAmount) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("reveal")), "Contract not in REVEAL status");

        // Check that the bid is higher than the current highest bid
        require(bidAmount > highestBid, "There already is a higher bid.");

        // Check that the bid amount corresponds to the hash 
        require(keccak256(abi.encodePacked(bidAmount)) == bidHashes[msg.sender]);

        // Attempt to transfer the tokens from the bidder to the contract
        bool transferSuccessful = token.transferFrom(msg.sender, address(this), bidAmount);
        
        // Check that the token transfer was successful
        require(transferSuccessful, "Token transfer failed.");

        // If there was a previous bid, allow the previous highest bidder to withdraw their bid
        if(highestBid > 0) {
            bool refundSuccessful = token.transferFrom(address(this), highestBidder, highestBid);
            require(refundSuccessful, "Refund failed.");
        }
        
        // Update the highest bid and highest bidder
        highestBidder = msg.sender;
        secondHighestBid = highestBid;
        highestBid = bidAmount;
        emit HighestBidIncreased(asset_id, msg.sender, bidAmount);
    }

    function closeAuction(bool not_winner_platform) public {
        require(msg.sender == owner, "Only owner can change contract's status");
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("reveal")), "Contract not in REVEAL status");

        status = "ending";
        if (not_winner_platform || highestBid == 0){
            status = "closed";

            highestBidder = address(0);
            highestBid = 0;

            return;
        }

        emit WaitResponse(highestBidder);
    }

    function abort(string memory jsonString) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("ending")), "Contract not in ENDING status");

        require(msg.sender == highestBidder, "Not authorized access!");

        status = "closing";
        emit DecisionMade(highestBidder, highestBid, secondHighestBid, asset_id, false, jsonString);

        if(highestBid > 0) {
            bool refundSuccessful = token.transferFrom(address(this), highestBidder, highestBid);
            require(refundSuccessful, "Refund failed.");
        }

        highestBidder = address(0);
        highestBid = 0;
    }

    function commit(string memory jsonString) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("ending")), "Contract not in ENDING status");
        // For testing only
        require(msg.sender == highestBidder, "Not authorized access!");
        
        status = "closing";

        token.burn(address(this), highestBid);
        emit DecisionMade(highestBidder, highestBid, secondHighestBid, asset_id, true, jsonString);
    }

    function provide_feedback(int32 _score) public {
        // Use hash to check status
        require(keccak256(abi.encodePacked(status)) == keccak256(abi.encodePacked("closing")), "Contract not in CLOSING status");
        // For testing only
        require(msg.sender == highestBidder, "Not authorized access!");

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
