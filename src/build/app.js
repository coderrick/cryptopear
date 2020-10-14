"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var express = require("express");
// Create a new express application instance
var app = express();
app.get('/', function (req, res) {
    res.send('Route test, hello you made it!');
});
app.listen(3000, function () {
    console.log('listening on port 3000!');
});
