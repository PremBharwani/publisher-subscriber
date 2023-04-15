
// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/Sub.sol";

contract TestSubCurrency {

    Sub public SUB;

    // Run before every test function
    function beforeEach() public {
        SUB = new Sub();
    }

    //function to test create subscriber
    function test_create_subscriber() public {
        SUB.create_subscriber(address(int));
        uint[] memory access;
        access = SUB.get_subscriber(address(int));
    }

    //function to add a subscriber to an event stream
    function test_subscribe_to_event() public {
        uint stream_id = 7;
        SUB.create_subscriber(address(this));
        SUB.subscribe_to_event(stream_id, address(this));
        uint[] memory access = SUB.get_subscriber(address(this));
        Assert.equal(access[stream_id], true, "Adding subscriber to stream_id successful");
    }

    //function to remove a subscriber from an event stream
    function test_unsubscribe_to_event() public {
        uint stream_id = 2;
        SUB.create_subscriber(address(this));
        SUB.subscribe_to_event(stream_id, address(this));
        uint[] memory access;
        access = SUB.get_subscriber(address(this));
        Assert.equal(access[stream_id], true, "Adding subscriber to stream_id successful");
        SUB.unsubscribe_to_event(2, address(this));
        Assert.equal(access[stream_id], 0, "Removing subscriber unsuccesful");
    }

    //function to test delete subscriber
    function test_delete_subscriber() public {
        uint stream_id = 6;
        SUB.create_subscriber(address(this));
        SUB.subscribe_to_event(stream_id, address(this));
        uint[] memory access = SUB.get_subscriber(address(this));
        Assert.equal(access[stream_id], true, "Adding subscriber to stream_id successful");
        access = SUB.get_subscriber(address(this));
        Assert.equal(access.length, 0, "Subscriber Deletion unsuccessful");
    }
}