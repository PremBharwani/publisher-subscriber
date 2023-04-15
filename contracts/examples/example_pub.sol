// SPDX-License-Identifier: GPL-3.0




//This is a example smart contract to use the publisher contract
// We can make multiple publishers in this smart contract using their address
// After deploying this smart contract we can call different functions of this smart contract to execute the functions of publisher contract
// There can be several other ways also to use the publisher contract according to the user's need

pragma solidity >=0.7.0 <0.9.0;
import {Pub} from "../Pub.sol";     // through this way we can import the publisher contract



contract example_pub{

    Pub public myPub;    // This line declares the instance of publisher contract
    constructor(){
        myPub = Pub(0x0B0F0d4A09117feFfE604c87Db345E400DAeBE46);     //create instance of deployed publisher contract at a specific address
    }

    function test_create_publisher() public{
        myPub.create_publisher(0xc1AaD10265fbA919d119097A65fEF71201A2b86D);  //create publisher with a address
    }

    function test_add_publisher() public{
        myPub.add_publisher(5, 0xc1AaD10265fbA919d119097A65fEF71201A2b86D); //add event access to publisher
    }

    function test_publish_to_eventstream() public{
        myPub.publish_to_eventstream("data which needs to be transferred",5, 0xc1AaD10265fbA919d119097A65fEF71201A2b86D);  // this funtction will publish data to event stream
    }

    function test_remove_publisher() public{
        myPub.remove_publisher(5, 0xc1AaD10265fbA919d119097A65fEF71201A2b86D);  // remove event access from publisher
    }

    function test_delete_publisher() public{
        myPub.delete_publisher(0xc1AaD10265fbA919d119097A65fEF71201A2b86D);   // through this delete publisher
    }

}