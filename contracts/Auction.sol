pragma solidity >=0.5.0 <0.6.0;

contract Aunction {
    address cryptopearBalance;
    address buyerBalance;
    address sellerBalance;
    uint itemprice;
    string itemname;
    uint256[] public inventory = new uint256[](100);
    
    //events
    event newitem(string itemname);
    event newprice(uint itemprice);
    //event NewItem(uint itemId, string itemname, uint price);
    
    // struct CryptoPear{
    //     string itemname;
    //     string price;
    // }
    
    function setitem(string memory item) public {
        itemname = item;
        emit newitem(item);
    }
    
    function getitem() public view returns (string memory){
        return itemname;
    }
    
    function setprice(uint item) public {
        itemprice = item;
        emit newprice(item);
    }
    
    function getprice() public view returns (uint){
        return itemprice;
    }