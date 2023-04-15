var EventQueue = artifacts.require("./Events.sol");

module.exports = function(deployer) {
  deployer.deploy(EventQueue);
};
