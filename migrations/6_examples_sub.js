var examples_pub = artifacts.require("./examples/example_sub.sol");

module.exports = function(deployer) {
  deployer.deploy(examples_pub);
};