
// SPDX-License-Identifier: MIT
// OpenZeppelin Contracts (last updated v4.9.0) (utils/cryptography/ECDSA.sol)

//pragma solidity >=0.7.0 <0.9.0;
pragma solidity ^0.8.0;

    contract RegistrarHolderWithDid {

    address public reg_user;
    string sigat;
    bytes32 sig;
    bytes sig1;
    string did_submit;
    string did_sub;
    string holder_add;
    string didholder1;
    string tran_hash;
    string transaction_hash;
    address holder_addres;
    

    mapping (address=>holder)holders;

    struct holder {

        address didholder;
        string did; 
        string name; 
        bool isregistered;   

    }

    constructor() {
        reg_user = msg.sender;
        //systemAddress = _systemAddress;
        sig1 = "0x15e614e7dd17599a2c0f824a7247dda0675c13787e2333fc6dc0b671625db7b69b44fac2a8af73202bae54293e0e46ad";
        sigat = string(abi.encodePacked(sig1));
        did_submit = "DShjcTEbHi1QmrjXeaF4H7";
        did_sub = string(abi.encodePacked(did_submit));
        holder_add = "0x6457b27f7a486931d5d6d894d63fd3f8285d8069";
        holder_addres = stringToAddress(holder_add);
        //didholder1 = string(abi.encodePacked(holder_add));
        tran_hash = "0xf66cc514b24516dfecd9be35e3b3a2427ad7f5302ed9c0515fb5ed4b2c4e1363";
        transaction_hash = string(abi.encodePacked(tran_hash));
    }

    // function or(string memory x, string memory y) external pure returns (string) {
    //     return (x | y);
    // }

    modifier onlyRegistered {

        require(holders[msg.sender].isregistered, "You are not registered");
        _;

    }

    function stringToAddress(string memory _address) public pure returns (address) {
    string memory cleanAddress = remove0xPrefix(_address);
    bytes20 _addressBytes = parseHexStringToBytes20(cleanAddress);
    return address(_addressBytes);
}

    function remove0xPrefix(string memory _hexString) internal pure returns (string memory) {
    if (bytes(_hexString).length >= 2 && bytes(_hexString)[0] == '0' && (bytes(_hexString)[1] == 'x' || bytes(_hexString)[1] == 'X')) {
        return substring(_hexString, 2, bytes(_hexString).length);
    }
    return _hexString;
    }

    function substring(string memory _str, uint256 _start, uint256 _end) internal pure returns (string memory) {
    bytes memory _strBytes = bytes(_str);
    bytes memory _result = new bytes(_end - _start);
    for (uint256 i = _start; i < _end; i++) {
        _result[i - _start] = _strBytes[i];
    }
    return string(_result);
    }

    function parseHexStringToBytes20(string memory _hexString) internal pure returns (bytes20) {
    bytes memory _bytesString = bytes(_hexString);
    uint160 _parsedBytes = 0;
    for (uint256 i = 0; i < _bytesString.length; i += 2) {
        _parsedBytes *= 256;
        uint8 _byteValue = parseByteToUint8(_bytesString[i]);
        _byteValue *= 16;
        _byteValue += parseByteToUint8(_bytesString[i + 1]);
        _parsedBytes += _byteValue;
    }
    return bytes20(_parsedBytes);
    }

    function parseByteToUint8(bytes1 _byte) internal pure returns (uint8) {
    if (uint8(_byte) >= 48 && uint8(_byte) <= 57) {
        return uint8(_byte) - 48;
    } else if (uint8(_byte) >= 65 && uint8(_byte) <= 70) {
        return uint8(_byte) - 55;
    } else if (uint8(_byte) >= 97 && uint8(_byte) <= 102) {
        return uint8(_byte) - 87;
    } else {
        revert(string(abi.encodePacked("Invalid byte value: ", _byte)));
    }
    }

    function registerHolders(address didholder,string memory did, string memory name) public returns(string memory) {

        // if (holders[didholder]){
        //     revert;
        // }
        require(holders[didholder].isregistered == false,"Multiple DID registration cannot be possible for a single Holder");
        bool validdid = iscompredid((abi.encodePacked(did_sub)),(abi.encodePacked(did)));
        bool validholder = iscompreholderadd((abi.encodePacked(holder_addres)),(abi.encodePacked(didholder)));

        if(validdid && validholder)
        {

        holders[didholder] = holder(didholder,did,name,true);
        string memory validinvalid = "Valid DID of Holder 1 and Holder1 genered address provided";
        string memory registration = " and Holder1 Registration has been done successfully";
        string memory validinvalid_registration = append (validinvalid,registration);
        return validinvalid_registration;
        }
        else
        {
        holders[didholder] = holder(didholder,did,name,false);
        string memory validinvalid  = "Eithar Invalid  DID of Holder1 or Holder1 genered address provided ";
        return validinvalid ;
        }

    }

    function append(string memory a, string memory b) internal pure returns (string memory) {

    return string(abi.encodePacked(a, b));

    }

    function getHolderDetails(address didholder) public view returns (address,string memory,string memory){
        
        return(holders[didholder].didholder,holders[didholder].did,holders[didholder].name);
        //console.log("Registered for the user DID:" + holders[didholder].did);
    }

    function iscompredid(bytes memory did1, bytes memory did2) public pure returns (bool validDid){
     
        if ((keccak256(abi.encodePacked(did1))) == (keccak256(abi.encodePacked(did2)))){
            return true;
        }
    }

    function iscompresign(bytes memory sign1, bytes memory sign2) public pure returns (bool validSignature){
     
        if ((keccak256(abi.encodePacked(sign1))) == (keccak256(abi.encodePacked(sign2)))){
           return true;
        }
    }

    function iscomprehash(bytes memory hash1, bytes memory hash2) public pure returns (bool validhash){
     
        if ((keccak256(abi.encodePacked(hash1))) == (keccak256(abi.encodePacked(hash2)))){
            return true;
        }
    }

    function iscompreholderadd(bytes memory holderadd1, bytes memory holderadd2) public pure returns (bool validAddr){
     
        if ((keccak256(abi.encodePacked(holderadd1))) == (keccak256(abi.encodePacked(holderadd2)))){
            return true;
        }
    }

    function isValid_Signature_Hash(string memory hash, string memory signature) public view returns (string memory, bool) {

        bytes memory sig2;
        sig2 = bytes(signature);

        bool validsign = iscompresign(((abi.encodePacked(sigat))),(abi.encodePacked(signature)));
        bool validhash = iscomprehash(((abi.encodePacked(transaction_hash))),(abi.encodePacked(hash)));
     
        if(validsign && validhash)
        {

        string memory validinvalid = "Valid AnonCred generated Verifier Signature and Hash value provided";
        bool val = true;
        return (validinvalid,val) ;
        }
        else
        {
        string memory validinvalid = "Invalid AnonCred generated Verifier Signature or Hash value provided";
        bool val = false;
        return (validinvalid, val) ;
        }

    }


    }