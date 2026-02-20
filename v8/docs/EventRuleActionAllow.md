# EventRuleActionAllow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RulesetId** | **string** | The ID of the evaluated ruleset. | 
**RuleId** | Pointer to **string** | The ID of the rule that matched the identification event. | [optional] 
**RuleExpression** | Pointer to **string** | The expression of the rule that matched the identification event. | [optional] 
**Type** | [**RuleActionType**](RuleActionType.md) |  | 
**RequestHeaderModifications** | Pointer to [**RequestHeaderModifications**](RequestHeaderModifications.md) |  | [optional] 

## Methods

### NewEventRuleActionAllow

`func NewEventRuleActionAllow(rulesetId string, type_ RuleActionType, ) *EventRuleActionAllow`

NewEventRuleActionAllow instantiates a new EventRuleActionAllow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRuleActionAllowWithDefaults

`func NewEventRuleActionAllowWithDefaults() *EventRuleActionAllow`

NewEventRuleActionAllowWithDefaults instantiates a new EventRuleActionAllow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRulesetId

`func (o *EventRuleActionAllow) GetRulesetId() string`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *EventRuleActionAllow) GetRulesetIdOk() (*string, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetId

`func (o *EventRuleActionAllow) SetRulesetId(v string)`

SetRulesetId sets RulesetId field to given value.


### GetRuleId

`func (o *EventRuleActionAllow) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *EventRuleActionAllow) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *EventRuleActionAllow) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *EventRuleActionAllow) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

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


