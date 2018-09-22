var express = require('express');
require('dotenv').config();
const PORT = process.env.PORT || 5000
var flash = require('connect-flash');
var passport = require('passport');
var request = require('request');
var session = require('express-session');
var app = express();
var bodyParser = require('body-parser');
var path = require('path');

app.use(require('cookie-parser')());
app.use(require('body-parser').urlencoded({ extended: true }));
const expressSession = require('express-session');
app.use(expressSession({secret: 'mySecretKey'}));
app.use(passport.initialize());
app.use(passport.session());
app.use('/public', express.static(__dirname + '/public'));
app.use(flash());
app.use(session({secret: 'keyboard cat'}))
app.use(bodyParser());
app.set('view engine', 'pug');
app.set('view options', { layout: false });
require('./routes/routes.js')(app);

var info = {
    "0": {
        "volumeTarget": 0.0,
        "purge_rate": 9000,
        "pump_id": 0,
        "rateW": 100.0,
        "volume": 10,
        "status": false,
        "name": "Outer",
        "direction": true,
        "syringe": 4.61,
        "used": true,
        "volumeTargetW": 0.0,
        "volumeW": 0.0,
        "rate": 100.0,
        "stalled": false,
        "force": 0
    },
    "1": {
        "volumeTarget": 0.0,
        "purge_rate": 4000,
        "pump_id": 1,
        "rateW": 100.0,
        "volume": 0.0,
        "status": false,
        "name": "Oil",
        "direction": true,
        "syringe": 4.61,
        "used": true,
        "volumeTargetW": 0.0,
        "volumeW": 0.0,
        "rate": 100.0,
        "stalled": false,
        "force": 0
    },
    "2": {
        "volumeTarget": 0.0,
        "purge_rate": 3000,
        "pump_id": 2,
        "rateW": 200.0,
        "volume": 0.0,
        "status": false,
        "name": "Inner",
        "direction": true,
        "syringe": 4.61,
        "used": true,
        "volumeTargetW": 0.0,
        "volumeW": 0.0,
        "rate": 200.0,
        "stalled": false,
        "force": 0
    },
    "3": {
        "volumeTarget": 0.0,
        "purge_rate": 1000,
        "pump_id": 3,
        "rateW": 250.0,
        "volume": 0.0,
        "status": false,
        "name": "1mL",
        "direction": true,
        "syringe": 4.61,
        "used": true,
        "volumeTargetW": 0.0,
        "volumeW": 0.0,
        "rate": 250.0,
        "stalled": false,
        "force": 0
    }
}

String.prototype.replaceAll = function(search, replacement) {
    var target = this;
    return target.replace(new RegExp(search, 'g'), replacement);
};


app.get("/refresh", function(req,res) {
    res.send()
})

app.post("/refresh", function(req,res) {
    var sample = {
        "data_pack": JSON.stringify(info).replaceAll("\"","\\\""),
        "success": 1
    }
    
    res.send(sample)
})

app.post("/update", function(req, res) {
    
    if (req.body.par == "rate") {
        info[req.body.pump.toString()].volume = req.body.value
        console.log(info[req.body.pump.toString()])
    }
    var sample1 = {
        "success": 1
    }
    res.send(sample1)
})

app.get("/", function(req,res) {
    res.send("ok")
})
app.listen(PORT);

console.log('Node listening on port %s', PORT);