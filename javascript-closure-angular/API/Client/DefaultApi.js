/**
 * @fileoverview AUTOMATICALLY GENERATED service for API.Client.DefaultApi.
 * Do not edit this file by hand or your changes will be lost next time it is
 * generated.
 *
 * This is the base document ( currently being edited )
 * Version: base
 * Generated at: 2016-11-25T06:45:55.633Z
 * Generated by: class io.swagger.codegen.languages.JavascriptClosureAngularClientCodegen
 */
goog.provide('API.Client.DefaultApi');


/**
 * @constructor
 * @param {!angular.$http} $http
 * @param {!Object} $httpParamSerializer
 * @param {!angular.$injector} $injector
 * @struct
 */
API.Client.DefaultApi = function($http, $httpParamSerializer, $injector) {
  /** @private {!string} */
  this.basePath_ = $injector.has('DefaultApiBasePath') ?
                   /** @type {!string} */ ($injector.get('DefaultApiBasePath')) :
                   'https://dev-virtserver.swaggerhub.com/acme/another-pet/base';

  /** @private {!Object<string, string>} */
  this.defaultHeaders_ = $injector.has('DefaultApiDefaultHeaders') ?
                   /** @type {!Object<string, string>} */ (
                       $injector.get('DefaultApiDefaultHeaders')) :
                   {};

  /** @private {!angular.$http} */
  this.http_ = $http;

  /** @package {!Object} */
  this.httpParamSerializer = $injector.get('$httpParamSerializer');
}
API.Client.DefaultApi.$inject = ['$http', '$httpParamSerializer', '$injector'];

/**
 * A cow
 * A cow
 * @param {!angular.$http.Config=} opt_extraHttpRequestParams Extra HTTP parameters to send.
 * @return {!angular.$q.Promise}
 */
API.Client.DefaultApi.prototype.addCow = function(opt_extraHttpRequestParams) {
  /** @const {string} */
  var path = this.basePath_ + '/cows';

  /** @type {!Object} */
  var queryParameters = {};

  /** @type {!Object} */
  var headerParams = angular.extend({}, this.defaultHeaders_);
  /** @type {!Object} */
  var httpRequestParams = {
    method: 'POST',
    url: path,
    json: true,
            params: queryParameters,
    headers: headerParams
  };

  if (opt_extraHttpRequestParams) {
    httpRequestParams = angular.extend(httpRequestParams, opt_extraHttpRequestParams);
  }

  return (/** @type {?} */ (this.http_))(httpRequestParams);
}

/**
 * Just a cat
 * Just a cat
 * @param {!angular.$http.Config=} opt_extraHttpRequestParams Extra HTTP parameters to send.
 * @return {!angular.$q.Promise}
 */
API.Client.DefaultApi.prototype.catsGet = function(opt_extraHttpRequestParams) {
  /** @const {string} */
  var path = this.basePath_ + '/cats';

  /** @type {!Object} */
  var queryParameters = {};

  /** @type {!Object} */
  var headerParams = angular.extend({}, this.defaultHeaders_);
  /** @type {!Object} */
  var httpRequestParams = {
    method: 'GET',
    url: path,
    json: true,
            params: queryParameters,
    headers: headerParams
  };

  if (opt_extraHttpRequestParams) {
    httpRequestParams = angular.extend(httpRequestParams, opt_extraHttpRequestParams);
  }

  return (/** @type {?} */ (this.http_))(httpRequestParams);
}
