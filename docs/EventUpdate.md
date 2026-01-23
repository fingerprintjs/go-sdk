# EventUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkedId** | Pointer to **string** | Linked Id value to assign to the existing event | [optional] 
**Tags** | Pointer to **map[string]interface{}** | A customer-provided value or an object that was sent with the identification request or updated later. | [optional] 
**Suspect** | Pointer to **bool** | Suspect flag indicating observed suspicious or fraudulent event | [optional] 

## Methods

### NewEventUpdate

`func NewEventUpdate() *EventUpdate`

NewEventUpdate instantiates a new EventUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventUpdateWithDefaults

`func NewEventUpdateWithDefaults() *EventUpdate`

NewEventUpdateWithDefaults instantiates a new EventUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkedId

`func (o *EventUpdate) GetLinkedId() string`

GetLinkedId returns the LinkedId field if non-nil, zero value otherwise.

### GetLinkedIdOk

`func (o *EventUpdate) GetLinkedIdOk() (*string, bool)`

GetLinkedIdOk returns a tuple with the LinkedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedId

`func (o *EventUpdate) SetLinkedId(v string)`

SetLinkedId sets LinkedId field to given value.

### HasLinkedId

`func (o *EventUpdate) HasLinkedId() bool`

HasLinkedId returns a boolean if a field has been set.

### GetTags

`func (o *EventUpdate) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EventUpdate) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EventUpdate) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *EventUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSuspect

`func (o *EventUpdate) GetSuspect() bool`

GetSuspect returns the Suspect field if non-nil, zero value otherwise.

### GetSuspectOk

`func (o *EventUpdate) GetSuspectOk() (*bool, bool)`

GetSuspectOk returns a tuple with the Suspect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspect

`func (o *EventUpdate) SetSuspect(v bool)`

SetSuspect sets Suspect field to given value.

### HasSuspect

`func (o *EventUpdate) HasSuspect() bool`

HasSuspect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


