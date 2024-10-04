'use strict';


/**
 *
 * returns LightingSummary
 **/
exports.getLightingSummary = function() {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {"empty": false};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 *
 * deviceId String 
 * returns DeviceState
 **/
exports.getSwitchState = function(deviceId) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {"empty": false};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 *
 * deviceId String 
 * value Integer 
 * returns ApiResponse
 **/
exports.setDimmer = function(deviceId,value) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {"empty": false};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * sets a dimmer to a specific value on a timer
 *
 * deviceId String 
 * value Integer 
 * timeunit Integer 
 * units String  (optional)
 * returns ApiResponse
 **/
exports.setDimmerTimer = function(deviceId,value,timeunit,units) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {"empty": false};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 *
 * deviceId String 
 * value String 
 * returns ApiResponse
 **/
exports.setSwitch = function(deviceId,value) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {"empty": false};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * sets a switch to a specific value on a timer
 *
 * deviceId String 
 * value String 
 * minutes Integer 
 * returns ApiResponse
 **/
exports.setSwitchTimer = function(deviceId,value,minutes) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {"empty": false};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}

