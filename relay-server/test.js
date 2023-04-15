const axios = require('axios');

const apiUrl = 'http://localhost:3000/send-events';

const jsonData = {
  type: 'relay_events',
  data: ['event1', 'event2']
};

axios.post(apiUrl, jsonData)
  .then(response => {
    console.log('Response:', response.data);
  })
  .catch(error => {
    console.error('Error:', error);
  });
