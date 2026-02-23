# EventRuleActionBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RulesetId** | **string** | The ID of the evaluated ruleset. | 
**RuleId** | Pointer to **string** | The ID of the rule that matched the identification event. | [optional] 
**RuleExpression** | Pointer to **string** | The expression of the rule that matched the identification event. | [optional] 
**Type** | [**RuleActionType**](RuleActionType.md) |  | 
**StatusCode** | Pointer to **int32** | A valid HTTP status code. | [optional] 
**Headers** | Pointer to [**[]RuleActionHeaderField**](RuleActionHeaderField.md) | A list of headers to send. | [optional] 
**Body** | Pointer to **string** | The response body to send to the client. | [optional] 

## Methods

### NewEventRuleActionBlock

`func NewEventRuleActionBlock(rulesetId string, type_ RuleActionType, ) *EventRuleActionBlock`

NewEventRuleActionBlock instantiates a new EventRuleActionBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRuleActionBlockWithDefaults

`func NewEventRuleActionBlockWithDefaults() *EventRuleActionBlock`

NewEventRuleActionBlockWithDefaults instantiates a new EventRuleActionBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRulesetId

`func (o *EventRuleActionBlock) GetRulesetId() string`

GetRulesetId returns the RulesetId field if non-nil, zero value otherwise.

### GetRulesetIdOk

`func (o *EventRuleActionBlock) GetRulesetIdOk() (*string, bool)`

GetRulesetIdOk returns a tuple with the RulesetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetId

`func (o *EventRuleActionBlock) SetRulesetId(v string)`

SetRulesetId sets RulesetId field to given value.


### GetRuleId

`func (o *EventRuleActionBlock) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *EventRuleActionBlock) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *EventRuleActionBlock) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *EventRuleActionBlock) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleExpression

`func (o *EventRuleActionBlock) GetRuleExpression() string`

GetRuleExpression returns the RuleExpression field if non-nil, zero value otherwise.

### GetRuleExpressionOk

`func (o *EventRuleActionBlock) GetRuleExpressionOk() (*string, bool)`

GetRuleExpressionOk returns a tuple with the RuleExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleExpression

`func (o *EventRuleActionBlock) SetRuleExpression(v string)`

SetRuleExpression sets RuleExpression field to given value.

### HasRuleExpression

`func (o *EventRuleActionBlock) HasRuleExpression() bool`

HasRuleExpression returns a boolean if a field has been set.

### GetType

`func (o *EventRuleActionBlock) GetType() RuleActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventRuleActionBlock) GetTypeOk() (*RuleActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventRuleActionBlock) SetType(v RuleActionType)`

SetType sets Type field to given value.


### GetStatusCode

`func (o *EventRuleActionBlock) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *EventRuleActionBlock) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *EventRuleActionBlock) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *EventRuleActionBlock) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetHeaders

`func (o *EventRuleActionBlock) GetHeaders() []RuleActionHeaderField`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *EventRuleActionBlock) GetHeadersOk() (*[]RuleActionHeaderField, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *EventRuleActionBlock) SetHeaders(v []RuleActionHeaderField)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *EventRuleActionBlock) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBody

`func (o *EventRuleActionBlock) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EventRuleActionBlock) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EventRuleActionBlock) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *EventRuleActionBlock) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


