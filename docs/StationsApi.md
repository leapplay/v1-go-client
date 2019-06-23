# \StationsApi

All URIs are relative to *https://api.leap-play.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStation**](StationsApi.md#CreateStation) | **Put** /api/v1/stations | Creates a Station (Auth)
[**GetSessionLogsAll**](StationsApi.md#GetSessionLogsAll) | **Get** /api/v1/stations/sessions | Gets closed Sessions from all stations (Auth)
[**GetSessionLogsByStationId**](StationsApi.md#GetSessionLogsByStationId) | **Get** /api/v1/stations/{stationId}/sessions | Gets closed Sessions (Auth)
[**GetSettings**](StationsApi.md#GetSettings) | **Get** /api/v1/stations/settings | Gets the Settings of all Stations (Auth)
[**GetSettingsByStationId**](StationsApi.md#GetSettingsByStationId) | **Get** /api/v1/stations/{stationId}/settings | Get the Settings of an Station (Auth)
[**GetState**](StationsApi.md#GetState) | **Get** /api/v1/stations/{stationId} | Get the Station&#39;s State (Auth)
[**GetStates**](StationsApi.md#GetStates) | **Get** /api/v1/stations | Gets Collection of Station States (Auth)
[**SessionCreate**](StationsApi.md#SessionCreate) | **Put** /api/v1/stations/{stationId}/sessions | Creates a Session (Auth)
[**SessionStop**](StationsApi.md#SessionStop) | **Post** /api/v1/stations/{stationId}/sessions/current/stop | Stops the running Session (Auth)
[**SessionUpdate**](StationsApi.md#SessionUpdate) | **Post** /api/v1/stations/{stationId}/sessions/current/update | Updates the running Session. (Auth)
[**SetSettings**](StationsApi.md#SetSettings) | **Post** /api/v1/stations/{stationId}/settings | Sets the Settings for an Station (Auth)
[**SetStationMode**](StationsApi.md#SetStationMode) | **Post** /api/v1/stations/{stationId}/settings/mode | Sets the Operation Mode (Auth)
[**SetStationPassword**](StationsApi.md#SetStationPassword) | **Post** /api/v1/stations/{stationId}/settings/password | Changes the Password (Auth)
[**SetStationQrCode**](StationsApi.md#SetStationQrCode) | **Post** /api/v1/stations/{stationId}/settings/qrcode | Sets the QrCode (Auth)


# **CreateStation**
> StationSettings CreateStation(ctx, optional)
Creates a Station (Auth)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateStationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateStationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStation** | [**optional.Interface of RequestCreateStation**](RequestCreateStation.md)| Create Station Dto | 

### Return type

[**StationSettings**](StationSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSessionLogsAll**
> []SessionLog GetSessionLogsAll(ctx, optional)
Gets closed Sessions from all stations (Auth)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSessionLogsAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSessionLogsAllOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **take** | **optional.Int32**| Entries to return | [default to 50]
 **skip** | **optional.Int32**| Entries to skip | [default to 0]

### Return type

[**[]SessionLog**](SessionLog.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSessionLogsByStationId**
> []SessionLog GetSessionLogsByStationId(ctx, stationId, optional)
Gets closed Sessions (Auth)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stationId** | [**string**](.md)| Station Id | 
 **optional** | ***GetSessionLogsByStationIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSessionLogsByStationIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **take** | **optional.Int32**| Entries to return | [default to 50]
 **skip** | **optional.Int32**| Entries to skip | [default to 0]

### Return type

[**[]SessionLog**](SessionLog.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettings**
> []StationSettings GetSettings(ctx, )
Gets the Settings of all Stations (Auth)

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]StationSettings**](StationSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettingsByStationId**
> StationSettings GetSettingsByStationId(ctx, stationId)
Get the Settings of an Station (Auth)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stationId** | [**string**](.md)| Station Id | 

### Return type

[**StationSettings**](StationSettings.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetState**
> StationCurrentState GetState(ctx, stationId)
Get the Station's State (Auth)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stationId** | [**string**](.md)| Station Id | 

### Return type

[**StationCurrentState**](StationCurrentState.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStates**
> []StationCurrentState GetStates(ctx, optional)
Gets Collection of Station States (Auth)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetStatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkStateFilter** | **optional.String**| Optional Network State Filter | 

### Return type

[**[]StationCurrentState**](StationCurrentState.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionCreate**
> CreatedSession SessionCreate(ctx, stationId, optional)
Creates a Session (Auth)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stationId** | [**string**](.md)| Station Id | 
 **optional** | ***SessionCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SessionCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSessionRequest** | [**optional.Interface of RequestNewStationSession**](RequestNewStationSession.md)| New Session Request Dto | 

### Return type

[**CreatedSession**](CreatedSession.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionStop**
> StoppedSession SessionStop(ctx, stationId)
Stops the running Session (Auth)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stationId** | [**string**](.md)| Station Id | 

### Return type

[**StoppedSession**](StoppedSession.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionUpdate**
> UpdatedSession SessionUpdate(ctx, stationId, optional)
Updates the running Session. (Auth)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stationId** | [**string**](.md)| Station Id | 
 **optional** | ***SessionUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SessionUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSessionRequest** | [**optional.Interface of RequestSessionUpdate**](RequestSessionUpdate.md)| The Update Request Dto | 

### Return type

[**UpdatedSession**](UpdatedSession.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSettings**
> SetSettings(ctx, stationId, optional)
Sets the Settings for an Station (Auth)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stationId** | [**string**](.md)| Station Id | 
 **optional** | ***SetSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetSettingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setSettingsRequest** | [**optional.Interface of RequestStationSettings**](RequestStationSettings.md)| Settings Dto | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetStationMode**
> SetStationMode(ctx, stationId, optional)
Sets the Operation Mode (Auth)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stationId** | [**string**](.md)| Station Id | 
 **optional** | ***SetStationModeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetStationModeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setModeRequest** | [**optional.Interface of RequestStationMode**](RequestStationMode.md)| The Operation Mode | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetStationPassword**
> SetStationPassword(ctx, stationId, optional)
Changes the Password (Auth)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stationId** | [**string**](.md)| Station Id | 
 **optional** | ***SetStationPasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetStationPasswordOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setPasswordRequest** | [**optional.Interface of RequestSetStationPassword**](RequestSetStationPassword.md)| New Password | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetStationQrCode**
> SetStationQrCode(ctx, stationId, optional)
Sets the QrCode (Auth)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stationId** | [**string**](.md)| Station Id | 
 **optional** | ***SetStationQrCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetStationQrCodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setQrCodeRequest** | [**optional.Interface of RequestStationQrCode**](RequestStationQrCode.md)| QrCode | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

