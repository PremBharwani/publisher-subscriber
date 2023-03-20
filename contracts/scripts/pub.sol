// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;
import "./event.sol" as c2;


library publisher_contract{

    address owner;
    constructor() {
        owner = msg.sender;
    }
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function.");
        _;
    }
    struct publisher1{
        string name;
        address address_publisher;
        bool exist;
        uint[] access;
    }
    mapping (address => publisher1) publisher;

    event published(uint indexed stream_id, uint indexed pub_id);
    event publisher_added(uint indexed stream_id, uint indexed pub_id);
    event publisher_created(string indexed name, address indexed address_publisher);
    event publisher_removed(uint indexed stream_id, uint indexed pub_id);

    function create_publisher(string memory _name, address _address_publisher,string stream_id) public view returns(uint){
        require(msg.sender == _address_publisher, "you are not allowed to create publisher with this address");
        publisher[_address_publisher] = publisher1(_name, _address_publisher, true, new bool[](0));
        emit publisher_created(_name, _address_publisher);
        access.push(stream_id);
        return publisher[_address_publisher].access.length;
    }

    function pub_to_event(uint data, uint stream_id, uint pub_id) public view {
        require(publisher[pub_id].exist == true, "publisher does not exist");
        require(publisher[pub_id].address_publisher == msg.sender, "you are not allowed to publish to this event");
        bool check= false;
        for(uint i = 0; i < publisher[pub_id].access.length; i++){
            if(publisher[pub_id].access[i] == stream_id){
                check = true;
            }
        }
        require(check == true, "publisher does not have access to this event");
        c2.publish_to_event_stream(stream_id, data);
        emit published(stream_id,pub_id);
        
    }

    function add_publisher(uint stream_id, uint pub_id) public view {
        require(publisher[pub_id].exist == true, "publisher does not exist");
        require(publisher[pub_id].address_publisher == msg.sender, "you are not allowed to add this publisher");
        bool check= false;
        for(uint i = 0; i < publisher[pub_id].access.length; i++){
            if(publisher[pub_id].access[i] == stream_id){
                check = true;
            }
        }
        require(check == false, "publisher already has access to this event");
        publisher[pub_id].access.push(stream_id);
        emit publisher_added(stream_id);
    }

    function remove_publisher(uint stream_id) public view{
        require(publisher[pub_id].exist == true, "publisher does not exist");
        require(publisher[pub_id].address_publisher == msg.sender, "you are not allowed to remove this publisher");
        bool check= false;
        for(uint i = 0; i < publisher[pub_id].access.length; i++){
            if(publisher[pub_id].access[i] == stream_id){
                check = true;
            }
        }
        require(check == true, "publisher does not have access to this event");
        for(uint i = 0; i < publisher[pub_id].access.length; i++){
            if(publisher[pub_id].access[i] == stream_id){
                publisher[pub_id].access[i] = 0;
            }
        }
        emit publisher_removed(stream_id,pub_id);
    
    }
    function delete_publisher(uint _id) public onlyOwner {
        require(publisher[_id].exist == true, "publisher does not exist");
        require(publisher[_id].address_publisher == msg.sender, "you are not allowed to delete this publisher");
        publisher[_id].exist = false;
        publisher[_id].address_publisher = address(0);
        publisher[_id].name = "";
        publisher[_id].access = new bool[](0);
    }

}