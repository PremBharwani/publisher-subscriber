// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0 ;

// to do: need to handle cases when queue is full
// to do: need to design a mechanism to clear the queue as the size is fixed
// to do: payable functions

contract EventQueue {

    // queue is actually a fixed size array
    // The idea is that each subscriber will have a variable lastSeenIndex for every topic it is subscribed to
    // we only store the hashes of data in the queue (which is stored on the blockchain)

    mapping (uint => bytes32[200]) public queues ; // map from queue ID to the queue associated with it
    // The size has to be declared before because memory arrays don't support the push operation (https://docs.soliditylang.org/en/v0.8.12/types.html#allocating-memory-arrays:~:text=Memory%20arrays%20with,copy%20every%20element.)
    // bytes32 is just enough for SHA-256 hashes

    mapping (uint => uint) public queue_next_index ; // in a static array, we'll need to maintain the next empty index 

    uint public num_queues=0 ; // need to keep as you can't get the number of mapped keys 
    
    mapping (uint => address) public owners ; // each queue has an owner (creator of the queue), & has the power to delete it

    // generates a new unique ID for a queue and return it
    // for now, the queue IDs are integers (uint) assigned in an incremental fashion (1,2, ...)
    function get_new_id() private returns (uint) {
        
        num_queues++ ;
        return num_queues ;
    }

    // creates an new event queue (here, a fixed size array) and returns the id of the queue 
    function create_event_stream() public returns (uint) {

        bytes32[200] memory queue ;
        uint queue_id=get_new_id() ;
        queues[queue_id]=queue ;
        owners[queue_id]=msg.sender ;
        queue_next_index[queue_id]=0 ;

        return queue_id ;
    }

    modifier only_if_owner(uint _event_id, address _user_address) {
        
        require(owners[_event_id]==_user_address, "You are not authorized to delete this queue!");
        _;
    }
    
    // deletes an existing event queue and returns the id of the queue
    // only the creator of the queue can delete it
    function delete_event_stream(uint _event_id) public returns (uint) {
        
        delete queues[_event_id] ;
        delete owners[_event_id] ;
        delete queue_next_index[_event_id] ;

        return _event_id ;
    }

    // pub.sol will call this function. It is assumed that all the checks will be performed in that contract 
    function publish_to_event_stream(uint _event_id, bytes32 _hash) public only_if_owner(_event_id, msg.sender) returns (uint)  {
        
        // queues[_event_id].push(_hash) ;
        queues[_event_id][queue_next_index[_event_id]]=_hash ;
        queue_next_index[_event_id]++ ;
        return _event_id ;
    }

    // sub.sol will call this function. It is assumed that all the checks will be performed in that contract 
    // returns the messages not seen by the subscriber from the queue
    function get_data_from_event_stream(uint _event_id, uint _last_seen_index) public view returns (bytes32[200] memory) {
        
        bytes32[200] memory latest_msg_hashes ;
        uint j=0 ;
        for(uint i=_last_seen_index+1; i<queues[_event_id].length; i++){
            latest_msg_hashes[j]=(queues[_event_id][i]);
            j++ ;
        }

        return latest_msg_hashes ;
    }


}
