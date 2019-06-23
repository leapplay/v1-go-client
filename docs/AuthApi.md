# \AuthApi

All URIs are relative to *https://api.leap-play.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Login**](AuthApi.md#Login) | **Post** /api/v1/auth/login | Receive an Access and Refresh Token
[**Logout**](AuthApi.md#Logout) | **Post** /api/v1/auth/logout | Invalidate the Refresh token (Auth)
[**RefreshToken**](AuthApi.md#RefreshToken) | **Post** /api/v1/auth/refreshtoken | Receive a new Access token


# **Login**
> LoginResponse Login(ctx, optional)
Receive an Access and Refresh Token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LoginOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoginOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginRequest** | [**optional.Interface of RequestLoginModel**](RequestLoginModel.md)| Login Request Dto | 

### Return type

[**LoginResponse**](LoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Logout**
> Logout(ctx, )
Invalidate the Refresh token (Auth)

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshToken**
> AccessToken RefreshToken(ctx, optional)
Receive a new Access token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RefreshTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RefreshTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestTokenRefresh** | [**optional.Interface of RequestTokenRefresh**](RequestTokenRefresh.md)| Refresh Token Dto | 
 **authorization** | **optional.String**| Any previous Access Token. | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

