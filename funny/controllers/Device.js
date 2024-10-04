'use strict';

var utils = require('../utils/writer.js');
var Device = require('../service/DeviceService');

module.exports.getDevices = function getDevices (req, res, next) {
  var skip = req.swagger.params['skip'].value;
  var limit = req.swagger.params['limit'].value;
  Device.getDevices(skip,limit)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.register = function register (req, res, next) {
  var device = req.swagger.params['device'].value;
  Device.register(device)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
