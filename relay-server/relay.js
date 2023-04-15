const express = require('express');
const Web3 = require('web3');
const fs = require('fs');
const { exec } = require('child_process');

const app = express();
app.use(express.json());
// const web3 = new Web3('http://localhost:7545');

// the API endpoint
app.post('/send-events', (req, res) => {
  
    //   Request body format:
    // {
    //   "type": "relay_events",
    //   "sub_id": "0xc0ffee254729296a45a3885639AC7E10F9d54979",
    //   "data": ["event1", "event2"]
    // }

  // get data from req
  // console.log(typeof req.body)
  // console.log(req.body.type)

  let cmd_args = ""
  cmd_args+=req.body.sub_id
  cmd_args+=" "
  for (let i=0;i<req.body.data.length;i++){
    cmd_args+=req.body.data[i];
    cmd_args+=" "
  }

  // call the js file
  // console.log("before calling relay-events")

  var cmd ="truffle exec call-relay-events.js "+cmd_args 

  exec(cmd, (error, stdout, stderr) => {
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
