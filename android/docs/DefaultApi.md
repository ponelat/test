# DefaultApi

All URIs are relative to *https://dev-virtserver.swaggerhub.com/acme/another-pet/base*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addCow**](DefaultApi.md#addCow) | **POST** /cows | Add a cow
[**catsGet**](DefaultApi.md#catsGet) | **GET** /cats | Just a cat


<a name="addCow"></a>
# **addCow**
> addCow()

Add a cow

Add a cow to our collection of animals

### Example
```java
// Import classes:
//import io.swagger.client.api.DefaultApi;

DefaultApi apiInstance = new DefaultApi();
try {
    apiInstance.addCow();
} catch (ApiException e) {
    System.err.println("Exception when calling DefaultApi#addCow");
    e.printStackTrace();
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="catsGet"></a>
# **catsGet**
> catsGet()

Just a cat

Just a cat

### Example
```java
// Import classes:
//import io.swagger.client.api.DefaultApi;

DefaultApi apiInstance = new DefaultApi();
try {
    apiInstance.catsGet();
} catch (ApiException e) {
    System.err.println("Exception when calling DefaultApi#catsGet");
    e.printStackTrace();
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

