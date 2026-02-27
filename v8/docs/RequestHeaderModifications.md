# RequestHeaderModifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Remove** | Pointer to **[]string** | The list of headers to remove. | [optional] 
**Set** | Pointer to [**[]RuleActionHeaderField**](RuleActionHeaderField.md) | The list of headers to set, overwriting any existing headers with the same name. | [optional] 
**Append** | Pointer to [**[]RuleActionHeaderField**](RuleActionHeaderField.md) | The list of headers to append. | [optional] 

## Methods

### NewRequestHeaderModifications

`func NewRequestHeaderModifications() *RequestHeaderModifications`

NewRequestHeaderModifications instantiates a new RequestHeaderModifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestHeaderModificationsWithDefaults

`func NewRequestHeaderModificationsWithDefaults() *RequestHeaderModifications`

NewRequestHeaderModificationsWithDefaults instantiates a new RequestHeaderModifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemove

`func (o *RequestHeaderModifications) GetRemove() []string`

GetRemove returns the Remove field if non-nil, zero value otherwise.

### GetRemoveOk

`func (o *RequestHeaderModifications) GetRemoveOk() (*[]string, bool)`

GetRemoveOk returns a tuple with the Remove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemove

`func (o *RequestHeaderModifications) SetRemove(v []string)`

SetRemove sets Remove field to given value.

### HasRemove

`func (o *RequestHeaderModifications) HasRemove() bool`

HasRemove returns a boolean if a field has been set.

### GetSet

`func (o *RequestHeaderModifications) GetSet() []RuleActionHeaderField`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *RequestHeaderModifications) GetSetOk() (*[]RuleActionHeaderField, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *RequestHeaderModifications) SetSet(v []RuleActionHeaderField)`

SetSet sets Set field to given value.

### HasSet

`func (o *RequestHeaderModifications) HasSet() bool`

HasSet returns a boolean if a field has been set.

### GetAppend

`func (o *RequestHeaderModifications) GetAppend() []RuleActionHeaderField`

GetAppend returns the Append field if non-nil, zero value otherwise.

### GetAppendOk

`func (o *RequestHeaderModifications) GetAppendOk() (*[]RuleActionHeaderField, bool)`

GetAppendOk returns a tuple with the Append field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppend

`func (o *RequestHeaderModifications) SetAppend(v []RuleActionHeaderField)`

SetAppend sets Append field to given value.

### HasAppend

`func (o *RequestHeaderModifications) HasAppend() bool`

HasAppend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

