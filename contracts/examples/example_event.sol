// SPDX-License-Identifier: GPL-3.0




//This is a example smart contract to use the events contract


pragma solidity >=0.7.0 <0.9.0;
import {Events} from "../Events.sol";     // through this way we can import the events contract




contract example_event{

    Events public eventQueue;    // This line declares the instance of events contract
    constructor(){
        eventQueue = Events(0x5B38Da6a701c568545dCfcB03FcB875f56beddC4); 
    }

    uint public stream_id = eventQueue.add_topic() ; // This stores the event stream id of the created queue

    function test()public {
        eventQueue.delete_topic(stream_id) ; // This deletes queue corresponding to event stream
    }

}