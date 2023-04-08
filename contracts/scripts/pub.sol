// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;
// import {EventQueu} from "./events.sol";


contract pub{

    // this contract is stores information about all the publishers
    // it can create a publisher.
    // it also has the function to publish to an event stream
    // it also has the function to add or remove a publisher from an event stream


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
    mapping (address => publisher1) public publisher;

    event published(uint indexed stream_id, address indexed pub_id);
    event publisher_added(uint indexed stream_id, address indexed pub_id);
    event publisher_created(string indexed name, address indexed address_publisher);
    event publisher_removed(uint indexed stream_id, address indexed pub_id);

    //create publisher allows users to create a publisher and it emits a publisher 
    function create_publisher(string memory _name, address _address_publisher,uint stream_id) public returns(uint){
        require(msg.sender == _address_publisher, "you are not allowed to create publisher with this address");
        publisher[_address_publisher] = publisher1(_name, _address_publisher, true, new uint[](0));
        emit publisher_created(_name, _address_publisher);
        publisher[_address_publisher].access.push(stream_id);
        return publisher[_address_publisher].access.length;
    }

    function pub_to_event(bytes32 data, uint stream_id, address pub_id) public{
        require(publisher[pub_id].exist == true, "publisher does not exist");
        require(publisher[pub_id].address_publisher == msg.sender, "you are not allowed to publish to this event");
        bool check= false;
        for(uint i = 0; i < publisher[pub_id].access.length; i++){
            if(publisher[pub_id].access[i] == stream_id){
                check = true;
            }
        }
        require(check == true, "publisher does not have access to this event");
        // publish_to_event_stream(stream_id, data);
        emit published(stream_id,pub_id);

        //Issue - Implementation -> Data to be published to event stream 
        
    }

    function add_publisher(uint stream_id, address pub_id) public{
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
        emit publisher_added(stream_id,pub_id);
    }

    function remove_publisher(uint stream_id,address pub_id) public {
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
    function delete_publisher(address _id) public onlyOwner {
        require(publisher[_id].exist == true, "publisher does not exist");
        require(publisher[_id].address_publisher == msg.sender, "you are not allowed to delete this publisher");
        publisher[_id].exist = false;
        publisher[_id].address_publisher = address(0);
        publisher[_id].name = "";
        publisher[_id].access = new uint[](0);
    }

    function get_publisher(address _id) public view returns(string memory, address, uint[] memory){
        require(publisher[_id].exist == true, "publisher does not exist");
        return (publisher[_id].name, publisher[_id].address_publisher, publisher[_id].access);
    }

}