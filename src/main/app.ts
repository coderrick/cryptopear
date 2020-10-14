import express = require('express');

// Create a new express application instance
const app: express.Application = express();

app.get('/', function (req, res) {
  res.send('Route test, hello you made it!');
});

app.listen(3000, function () {
  console.log('listening on port 3000!');
});