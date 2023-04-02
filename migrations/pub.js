var pub = artifacts.require("./scripts/pub.sol");

module.exports = function(deployer) {
  deployer.deploy(pub);
};
