const Sub = artifacts.require('Sub');
const axios = require('axios');
const fs = require('fs');

// const args = ["event1", "event2"]

const all_args = process.argv 
console.log(all_args)
var args = []

// fill events from the passed args
var addr = all_args[4];
for (let i=5;i<all_args.length;i++){
    args.push(all_args[i])
}

module.exports = async function(callback) {
    try {

        // Fetch the deployed exchange
        const sub = await Sub.deployed()
        console.log('Sub contract fetched', Sub.address)

        await sub.relay_events(args,addr)
        console.log(`relayed events called`)
        // sanity check (make the relay_events_called public to view it here )
        const flag = await sub.relay_events_called()
        console.log(flag)

        const ev = await sub.checkEvent()
        console.log(ev)
        
    }
    catch(error) {
      console.log(error)
    }
    callback()
  }