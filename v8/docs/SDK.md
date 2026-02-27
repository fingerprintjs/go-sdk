# SDK

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** | Platform of the SDK used for the identification request. | 
**Version** | **string** | Version string of the SDK used for the identification request. For example: `\"3.12.1\"`  | 
**Integrations** | Pointer to [**[]Integration**](Integration.md) |  | [optional] 

## Methods

### NewSDK

`func NewSDK(platform string, version string, ) *SDK`

NewSDK instantiates a new SDK object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSDKWithDefaults

`func NewSDKWithDefaults() *SDK`

NewSDKWithDefaults instantiates a new SDK object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *SDK) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *SDK) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *SDK) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetVersion

`func (o *SDK) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SDK) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SDK) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetIntegrations

`func (o *SDK) GetIntegrations() []Integration`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *SDK) GetIntegrationsOk() (*[]Integration, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *SDK) SetIntegrations(v []Integration)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *SDK) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

