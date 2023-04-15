// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

// // npm install @openzeppelin/contracts
// import "@openzeppelin/contracts/utils/Strings.sol";

// to do: payable functions

contract Events {
    // using Strings for address;
    uint public num_queues = 0;//counts the total number of queues

    event topic_added(uint stream_id);
    event topic_deleted(uint stream_id);
    mapping (uint => address) public owners; //owner of the topic
    
    function add_topic() public {
        num_queues++;
        owners[num_queues]=msg.sender;
        emit topic_added(num_queues);
    }

    function delete_topic(uint _topic_id) public {
        require(owners[_topic_id]==msg.sender, "You are not authorized to delete this queue!");
        delete owners[_topic_id];
        emit topic_deleted(_topic_id);
    }

}
