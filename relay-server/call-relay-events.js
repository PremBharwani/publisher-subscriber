const Sub = artifacts.require('Sub');
const axios = require('axios');
const fs = require('fs');

const jsonData = JSON.parse(fs.readFileSync('relay_events.json'));

const args = jsonData.data;

// const args = ["event1", "event2"]

module.exports = async function(callback) {
    try {

        // Fetch the deployed exchange
        const sub = await Sub.deployed()
        console.log('Sub contract fetched', Sub.address)

        await sub.relay_events(args)
        console.log(`relayed events called`)
        const flag = await sub.relay_eventsCalled()
        console.log(flag)
        
    }
    catch(error) {
      console.log(error)
    }
    callback()
  }