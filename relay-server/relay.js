const express = require('express');
const Web3 = require('web3');
const fs = require('fs');
const { exec } = require('child_process');

const app = express();
app.use(express.json());
const web3 = new Web3('http://localhost:7545');

// the API endpoint
app.post('/send-events', (req, res) => {
  
    //   Request body:
    // {  
    //   "type": relay_events
    //   "data": ["event1", "event2"]
    // 
    // }

  // get data from req
  console.log(typeof req.body)
  console.log(req.body.type)
  const jsonString = JSON.stringify(req.body);
  
  // write data to json for the truffle js file to take as input
  console.log("before writing json")
  console.log(jsonString)
  
  fs.writeFile('relay_events.json', jsonString, (err) => {
    if (err) {
      console.error(err);
    } else {
      console.log('JSON made');
    }
  });

  // call the js file
  console.log("before calling relay-events")


  exec('truffle exec call-relay-events.js', (error, stdout, stderr) => {
    if (error) {
      console.error(`exec error: ${error}`);
      return;
    }
    console.log(`stdout: ${stdout}`);
    console.error(`stderr: ${stderr}`);
  });

  const ret = "OK"
  res.json({ ret });
});

// Start the server
const port = 3000;
app.listen(port, () => {
  console.log(`Server running on port ${port}`);
});
