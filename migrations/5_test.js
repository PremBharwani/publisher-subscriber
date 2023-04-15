var test = artifacts.require("./examples/test.sol");

module.exports = function(deployer) {
  deployer.deploy(test);
};