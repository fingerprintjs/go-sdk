# Integration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the specific integration, e.g. \"fingerprint-pro-react\". | [optional] 
**Version** | Pointer to **string** | The version of the specific integration, e.g. \"3.11.10\". | [optional] 
**Subintegration** | Pointer to [**IntegrationSubintegration**](IntegrationSubintegration.md) |  | [optional] 

## Methods

### NewIntegration

`func NewIntegration() *Integration`

NewIntegration instantiates a new Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationWithDefaults

`func NewIntegrationWithDefaults() *Integration`

NewIntegrationWithDefaults instantiates a new Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Integration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Integration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Integration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Integration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *Integration) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Integration) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Integration) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Integration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSubintegration

`func (o *Integration) GetSubintegration() IntegrationSubintegration`

GetSubintegration returns the Subintegration field if non-nil, zero value otherwise.

### GetSubintegrationOk

`func (o *Integration) GetSubintegrationOk() (*IntegrationSubintegration, bool)`

GetSubintegrationOk returns a tuple with the Subintegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubintegration

`func (o *Integration) SetSubintegration(v IntegrationSubintegration)`

SetSubintegration sets Subintegration field to given value.

### HasSubintegration

`func (o *Integration) HasSubintegration() bool`

HasSubintegration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

