# EventRuleAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RulesetId** | **string** | The ID of the evaluated ruleset. | 
**RuleId** | Pointer to **string** | The ID of the rule that matched the identification event. | [optional] 
**RuleExpression** | Pointer to **string** | The expression of the rule that matched the identification event. | [optional] 

## Methods

### NewEventRuleAction

`func NewEventRuleAction(rulesetId string, ) *EventRuleAction`

NewEventRuleAction instantiates a new EventRuleAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRuleActionWithDefaults

`func NewEventRuleActionWithDefaults() *EventRuleAction`

NewEventRuleActionWithDefaults instantiates a new EventRuleAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRulesetId

`func (o *EventRuleAction) GetRulesetId() string`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *EventRuleAction) GetRulesetIdOk() (*string, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetId

`func (o *EventRuleAction) SetRulesetId(v string)`

SetRulesetId sets RulesetId field to given value.


### GetRuleId

`func (o *EventRuleAction) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *EventRuleAction) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *EventRuleAction) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *EventRuleAction) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleExpression

`func (o *EventRuleAction) GetRuleExpression() string`

GetRuleExpression returns the RuleExpression field if non-nil, zero value otherwise.

### GetRuleExpressionOk

`func (o *EventRuleAction) GetRuleExpressionOk() (*string, bool)`

GetRuleExpressionOk returns a tuple with the RuleExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleExpression

`func (o *EventRuleAction) SetRuleExpression(v string)`

SetRuleExpression sets RuleExpression field to given value.

### HasRuleExpression

`func (o *EventRuleAction) HasRuleExpression() bool`

HasRuleExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


