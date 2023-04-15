// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/Pub.sol";

contract TestPubCurrency {

    Pub public PUB;

    // Run before every test function
    function beforeEach() public {
        PUB = new Pub();
    }

    // Test that create_publisher works
    function test_create_publisher() public {
        // require(publisher_list[address(this)].exist == false, "publisher already exist");
        PUB.create_publisher(address(this));
        uint[] memory access;
        access = PUB.get_publisher(address(this));

    }

    // Test to give publisher access to stream_id
    function test_add_publisher() public{
        uint stream_id = 7;
        PUB.create_publisher(address(this));
        PUB.add_publisher(stream_id, address(this));
        uint[] memory access = PUB.get_publisher(address(this));
        bool check= false;
        for(uint i = 0; i < access.length; i++){
            if(access[i] == stream_id){
                check = true;
            }
        }
        Assert.equal(check, true, "Adding publisher to stream_id successful");
    }

    //Test whether publishing to stream was successful
    function test_publish_to_eventstream() public {
        uint stream_id = 7;
        PUB.create_publisher(address(this));
        PUB.add_publisher(stream_id, address(this));
        uint[] memory access = PUB.get_publisher(address(this));
         bool check= false;
        for(uint i = 0; i < access.length; i++){
            if(access[i] == stream_id){
                check = true;
            }
        }
        Assert.equal(check, true, "Adding publisher to stream_id successful");
        PUB.publish_to_eventstream("example", stream_id, address(this));
    }

    //Test to check remove_publisher
    function test_remove_publisher() public {
        uint stream_id = 2;
        PUB.create_publisher(address(this));
        PUB.add_publisher(stream_id, address(this));
        uint[] memory access;
         bool check= false;
        for(uint i = 0; i < access.length; i++){
            if(access[i] == stream_id){
                check = true;
            }
        }
        Assert.equal(check, true, "Adding publisher to stream_id successful");
        PUB.remove_publisher(2, address(this));
        Assert.equal(access[stream_id], 0, "Removing publisher unsuccesful");
    }


    //Test to check delete_publisher
    function test_delete_publihser() public {
        uint stream_id = 6;
        PUB.create_publisher(address(this));
        PUB.add_publisher(stream_id, address(this));
        uint[] memory access = PUB.get_publisher(address(this));
         bool check= false;
        for(uint i = 0; i < access.length; i++){
            if(access[i] == stream_id){
                check = true;
            }
        }
        Assert.equal(check, true, "Adding publisher to stream_id successful");
        PUB.delete_publisher(address(this));
        access = PUB.get_publisher(address(this));
        Assert.equal(access.length, 0, "Publisher Deletion unsuccessful");
    }

}
