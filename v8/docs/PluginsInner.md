# PluginsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**MimeTypes** | Pointer to [**[]PluginsInnerMimeTypesInner**](PluginsInnerMimeTypesInner.md) |  | [optional] 

## Methods

### NewPluginsInner

`func NewPluginsInner(name string, ) *PluginsInner`

NewPluginsInner instantiates a new PluginsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginsInnerWithDefaults

`func NewPluginsInnerWithDefaults() *PluginsInner`

NewPluginsInnerWithDefaults instantiates a new PluginsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PluginsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PluginsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PluginsInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PluginsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PluginsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PluginsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PluginsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMimeTypes

`func (o *PluginsInner) GetMimeTypes() []PluginsInnerMimeTypesInner`

GetMimeTypes returns the MimeTypes field if non-nil, zero value otherwise.

### GetMimeTypesOk

`func (o *PluginsInner) GetMimeTypesOk() (*[]PluginsInnerMimeTypesInner, bool)`

GetMimeTypesOk returns a tuple with the MimeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeTypes

`func (o *PluginsInner) SetMimeTypes(v []PluginsInnerMimeTypesInner)`

SetMimeTypes sets MimeTypes field to given value.

### HasMimeTypes

`func (o *PluginsInner) HasMimeTypes() bool`

HasMimeTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


