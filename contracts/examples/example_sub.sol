// SPDX-License-Identifier: GPL-3.0




//This is a example smart contract to use the subscriber contract


pragma solidity >=0.7.0 <0.9.0;
import {Sub} from "../Sub.sol";     // through this way we can import the subscriber contrac`t



contract example_sub{

    Sub public mySub;    // This line declares the instance of subscriber contract
    constructor(){
        mySub = Sub(0x97ecEd422C160411512f23836364bD1a66c01f2A);     //create instance of deployed subscriber contract at a specific address
    }

    function test_create_subscriber() public{
        mySub.create_subscriber(0xa8A2e80a87e71412635bBaD9D413acccDB7aFFBE);  //create subscriber with a address
    }

    function test_subscribe_to_event() public{
        mySub.subscribe_to_event(5, 0xa8A2e80a87e71412635bBaD9D413acccDB7aFFBE); //add event access to subscriber
    }

    function test_call_for_events() public{
        mySub.call_for_events(5, 0xa8A2e80a87e71412635bBaD9D413acccDB7aFFBE);   // through this  fucntion get events
    }

    function test_unsubscribe_to_event() public{
        mySub.unsubscribe_to_event(5, 0xa8A2e80a87e71412635bBaD9D413acccDB7aFFBE);  // remove event access from subscriber
    }

    function test_delete_subscriber() public{
        mySub.delete_subscriber(0xa8A2e80a87e71412635bBaD9D413acccDB7aFFBE);   // through this delete subscriber
    }
}