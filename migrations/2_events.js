var EventQueue = artifacts.require("./scripts/Events.sol");

module.exports = function(deployer) {
  deployer.deploy(EventQueue);
};
