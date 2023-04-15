var pub = artifacts.require("./scripts/Pub.sol");

module.exports = function(deployer) {
  deployer.deploy(pub);
};
