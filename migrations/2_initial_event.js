var EventQueue = artifacts.require("./scripts/events.sol/EventQueue");

module.exports = function(deployer) {
  deployer.deploy(EventQueue);
};
