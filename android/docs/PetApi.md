# PetApi

All URIs are relative to *https://dev-virtserver.swaggerhub.com/acme/another-pet/base*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addDog**](PetApi.md#addDog) | **POST** /dogs | Add a new dog to the store


<a name="addDog"></a>
# **addDog**
> addDog()

Add a new dog to the store



### Example
```java
// Import classes:
//import io.swagger.client.api.PetApi;

PetApi apiInstance = new PetApi();
try {
    apiInstance.addDog();
} catch (ApiException e) {
    System.err.println("Exception when calling PetApi#addDog");
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

