// this file if for passing the queue events from the API, back to the contract by calling a function in sub.sol

const subContract = artifacts.require('Sub')
const axios = require('axios');

var events;
// read the logs and wait till a subscriber needs data

// events=get_events_from_the_queue_somehow()-->

axios.get('http://localhost:8080/get-event-data')
  .then(response => {
    console.log(response.data); 
    events=response.data;
  })
  .catch(error => {
    console.error(error);
  });

module.exports = async function(callback) {
    try {

        // Fetch the deployed contract
        const subFunc = await subContract.deployed()
        console.log('subContract fetched', subContract.address)

        await subFunc.relay_events(events)
        console.log(`Passed events to the contract`)

        
    }
    catch(error) {
      console.log(error)
    }
  
    callback()
}
