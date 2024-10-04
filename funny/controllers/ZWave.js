'use strict';

var utils = require('../utils/writer.js');
var ZWave = require('../service/ZWaveService');

module.exports.getLightingSummary = function getLightingSummary (req, res, next) {
  ZWave.getLightingSummary()
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.getSwitchState = function getSwitchState (req, res, next) {
  var deviceId = req.swagger.params['deviceId'].value;
  ZWave.getSwitchState(deviceId)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.setDimmer = function setDimmer (req, res, next) {
  var deviceId = req.swagger.params['deviceId'].value;
  var value = req.swagger.params['value'].value;
  ZWave.setDimmer(deviceId,value)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.setDimmerTimer = function setDimmerTimer (req, res, next) {
  var deviceId = req.swagger.params['deviceId'].value;
  var value = req.swagger.params['value'].value;
  var timeunit = req.swagger.params['timeunit'].value;
  var units = req.swagger.params['units'].value;
  ZWave.setDimmerTimer(deviceId,value,timeunit,units)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.setSwitch = function setSwitch (req, res, next) {
  var deviceId = req.swagger.params['deviceId'].value;
  var value = req.swagger.params['value'].value;
  ZWave.setSwitch(deviceId,value)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.setSwitchTimer = function setSwitchTimer (req, res, next) {
  var deviceId = req.swagger.params['deviceId'].value;
  var value = req.swagger.params['value'].value;
  var minutes = req.swagger.params['minutes'].value;
  ZWave.setSwitchTimer(deviceId,value,minutes)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
