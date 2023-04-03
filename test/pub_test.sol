// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/scripts/pub.sol";

contract TestSubCurrency {

    pub public PUB;

    // Run before every test function
    function beforeEach() public {
        PUB = new pub();
    }

    // Test that create_publisher works
    function testcreate_publisher() public {
        uint stream_id = 5;
        uint access_length = PUB.create_publisher("test", address(this), stream_id);
        Assert.equal(access_length, 1, "It should store the correct value");
        string memory name;
        address address_publisher;
        uint[] memory access;
        (name, address_publisher, access) = PUB.get_publisher(address(this));
        Assert.equal(name, "test", "It should store the correct value");
    }

    // Test that it adds publisher to access
    function testadd_publisher() public {
        uint stream_id = 5;
        PUB.create_publisher("test", address(this), stream_id);
        PUB.add_publisher(6, address(this));
        // uint access_length = PUB.get_publisher(address(this)).access.length;
        string memory name;
        address address_publisher;
        uint[] memory access;
        (name, address_publisher, access) = PUB.get_publisher(address(this));
        Assert.equal(name, "test", "It should store the correct value");
        Assert.equal(access.length, 2, "It should store the correct value");
    }

    // // Test that remove_publisher works
    function testremove_publisher() public {
        uint stream_id = 5;
        PUB.create_publisher("test", address(this), stream_id);
        PUB.add_publisher(6, address(this));
        // uint access_length = PUB.get_publisher(address(this)).access.length;
        string memory name;
        address address_publisher;
        uint[] memory access;
        (name, address_publisher, access) = PUB.get_publisher(address(this));
        Assert.equal(name, "test", "It should store the correct value");
        Assert.equal(access[1], 6, "It should store the correct value");
        PUB.remove_publisher(6, address(this));
        (name, address_publisher, access) = PUB.get_publisher(address(this));
        Assert.equal(access[1], 0, "It should store the correct value");
    }
    // Test that delete_publisher works
    function testdelete_publisher() public {
        uint stream_id = 5;
        PUB.create_publisher("test", address(this), stream_id);
        PUB.add_publisher(6, address(this));
        PUB.delete_publisher(address(this));
        string memory name;
        address address_publisher;
        uint[] memory access;
        (name, address_publisher, access) = PUB.get_publisher(address(this));
        Assert.equal(name, "", "It should store the correct value");
    }
}