# EventRuleActionAllow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**RuleActionType**](RuleActionType.md) |  | 
**RequestHeaderModifications** | Pointer to [**RequestHeaderModifications**](RequestHeaderModifications.md) |  | [optional] 

## Methods

### NewEventRuleActionAllow

`func NewEventRuleActionAllow(type_ RuleActionType, ) *EventRuleActionAllow`

NewEventRuleActionAllow instantiates a new EventRuleActionAllow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRuleActionAllowWithDefaults

`func NewEventRuleActionAllowWithDefaults() *EventRuleActionAllow`

NewEventRuleActionAllowWithDefaults instantiates a new EventRuleActionAllow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventRuleActionAllow) GetType() RuleActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventRuleActionAllow) GetTypeOk() (*RuleActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventRuleActionAllow) SetType(v RuleActionType)`

SetType sets Type field to given value.


### GetRequestHeaderModifications

`func (o *EventRuleActionAllow) GetRequestHeaderModifications() RequestHeaderModifications`

GetRequestHeaderModifications returns the RequestHeaderModifications field if non-nil, zero value otherwise.

### GetRequestHeaderModificationsOk

`func (o *EventRuleActionAllow) GetRequestHeaderModificationsOk() (*RequestHeaderModifications, bool)`

GetRequestHeaderModificationsOk returns a tuple with the RequestHeaderModifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaderModifications

`func (o *EventRuleActionAllow) SetRequestHeaderModifications(v RequestHeaderModifications)`

SetRequestHeaderModifications sets RequestHeaderModifications field to given value.

### HasRequestHeaderModifications

`func (o *EventRuleActionAllow) HasRequestHeaderModifications() bool`

HasRequestHeaderModifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


