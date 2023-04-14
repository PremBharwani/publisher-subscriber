// this file if for passing the queue events from the API, back to the contract by calling a function in sub.sol

const subContract = artifacts.require('Sub')

var events;

// events=get_events_from_the_queue_somehow()

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
