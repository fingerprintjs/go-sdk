# EventRuleActionAllow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RulesetID** | **string** | The ID of the evaluated ruleset. | 
**RuleID** | Pointer to **string** | The ID of the rule that matched the identification event. | [optional] 
**RuleExpression** | Pointer to **string** | The expression of the rule that matched the identification event. | [optional] 
**Type** | [**RuleActionType**](RuleActionType.md) |  | 
**RequestHeaderModifications** | Pointer to [**RequestHeaderModifications**](RequestHeaderModifications.md) |  | [optional] 

## Methods

### NewEventRuleActionAllow

`func NewEventRuleActionAllow(rulesetID string, type_ RuleActionType, ) *EventRuleActionAllow`

NewEventRuleActionAllow instantiates a new EventRuleActionAllow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRuleActionAllowWithDefaults

`func NewEventRuleActionAllowWithDefaults() *EventRuleActionAllow`

NewEventRuleActionAllowWithDefaults instantiates a new EventRuleActionAllow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRulesetID

`func (o *EventRuleActionAllow) GetRulesetID() string`

GetRulesetID returns the RulesetID field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *EventRuleActionAllow) GetRulesetIdOk() (*string, bool)`

GetRulesetIdOk returns a tuple with the RulesetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetID

`func (o *EventRuleActionAllow) SetRulesetID(v string)`

SetRulesetID sets RulesetID field to given value.


### GetRuleID

`func (o *EventRuleActionAllow) GetRuleID() string`

GetRuleID returns the RuleID field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *EventRuleActionAllow) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleID

`func (o *EventRuleActionAllow) SetRuleID(v string)`

SetRuleID sets RuleID field to given value.

### HasRuleID

`func (o *EventRuleActionAllow) HasRuleID() bool`

HasRuleID returns a boolean if a field has been set.

### GetRuleExpression

`func (o *EventRuleActionAllow) GetRuleExpression() string`

GetRuleExpression returns the RuleExpression field if non-nil, zero value otherwise.

### GetRuleExpressionOk

`func (o *EventRuleActionAllow) GetRuleExpressionOk() (*string, bool)`

GetRuleExpressionOk returns a tuple with the RuleExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleExpression

`func (o *EventRuleActionAllow) SetRuleExpression(v string)`

SetRuleExpression sets RuleExpression field to given value.

### HasRuleExpression

`func (o *EventRuleActionAllow) HasRuleExpression() bool`

HasRuleExpression returns a boolean if a field has been set.

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


