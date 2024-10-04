'use strict';


/**
 *
 * returns List
 **/
exports.getZones = function() {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 *
 * zoneId String 
 * no response value expected for this operation
 **/
exports.quietZone = function(zoneId) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}

