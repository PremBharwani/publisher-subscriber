// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/scripts/sub.sol";

contract TestSubCurrency {

    sub public SUB;

    // Run before every test function
    function beforeEach() public {
        SUB = new sub();
    }


    function testcreate_subscriber() public {
        address subs_id = SUB.create_subscriber(this, new address[](0));
        Assert.equal(sub_id, address(this), "It should store the correct value");


        address subscriber_id ;
        address[] event_streams_subscribed ;
        (subscriber_id, event_streams_subscribed) = SUB.get_subscriber(address(this));
        Assert.equal(subscriber_id, address(this) ,"It should store the correct value");

    }

  
}