# \AccountsApi

All URIs are relative to *https://api.leap-play.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePassword**](AccountsApi.md#ChangePassword) | **Post** /api/v1/accounts/password/change | Changes the Password and provides a new Refresh Token (Auth)
[**ConfirmEmail**](AccountsApi.md#ConfirmEmail) | **Post** /api/v1/accounts/email/confirmation | Confirms ownership of an E-Mail Address by an E-Mail Confirmation Token
[**Register**](AccountsApi.md#Register) | **Post** /api/v1/accounts/register | Registers a User Account
[**RequestPassword**](AccountsApi.md#RequestPassword) | **Post** /api/v1/accounts/password/forgot | Receives a Reset Token that is required to reset the Password
[**ResendEMailConfirmation**](AccountsApi.md#ResendEMailConfirmation) | **Post** /api/v1/accounts/email/confirmation/send | Re-sends a message with a EMail Confirmation Token.
[**ResetPassword**](AccountsApi.md#ResetPassword) | **Post** /api/v1/accounts/password/reset | Resets the Password with an Reset Token


# **ChangePassword**
> ChangedPasswordUser ChangePassword(ctx, changePasswordRequest)
Changes the Password and provides a new Refresh Token (Auth)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **changePasswordRequest** | [**RequestChangePassword**](RequestChangePassword.md)| The change password request. | 

### Return type

[**ChangedPasswordUser**](ChangedPasswordUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfirmEmail**
> ConfirmEmail(ctx, emailConfirmationRequest)
Confirms ownership of an E-Mail Address by an E-Mail Confirmation Token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailConfirmationRequest** | [**RequestEmailConfirmation**](RequestEmailConfirmation.md)| The request email confirmation. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Register**
> Register(ctx, registerRequest)
Registers a User Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registerRequest** | [**RequestRegisterUser**](RequestRegisterUser.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestPassword**
> RequestPassword(ctx, forgotPasswordRequest)
Receives a Reset Token that is required to reset the Password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **forgotPasswordRequest** | [**RequestForgotPassword**](RequestForgotPassword.md)| Request object holding the required parameter | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResendEMailConfirmation**
> ResendEMailConfirmation(ctx, resendConfirmationEmailRequest)
Re-sends a message with a EMail Confirmation Token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resendConfirmationEmailRequest** | [**RequestResendConfirmationEmail**](RequestResendConfirmationEmail.md)| The resend email confirmation request. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetPassword**
> ResetPassword(ctx, resetPasswordRequest)
Resets the Password with an Reset Token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resetPasswordRequest** | [**RequestResetPassword**](RequestResetPassword.md)| The request reset password. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

