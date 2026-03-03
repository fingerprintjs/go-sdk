# EventRuleAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RulesetID** | **string** | The ID of the evaluated ruleset. | 
**RuleID** | Pointer to **string** | The ID of the rule that matched the identification event. | [optional] 
**RuleExpression** | Pointer to **string** | The expression of the rule that matched the identification event. | [optional] 
**Type** | [**RuleActionType**](RuleActionType.md) |  | 
**RequestHeaderModifications** | Pointer to [**RequestHeaderModifications**](RequestHeaderModifications.md) |  | [optional] 
**StatusCode** | Pointer to **int32** | A valid HTTP status code. | [optional] 
**Headers** | Pointer to [**[]RuleActionHeaderField**](RuleActionHeaderField.md) | A list of headers to send. | [optional] 
**Body** | Pointer to **string** | The response body to send to the client. | [optional] 

## Methods

### NewEventRuleAction

`func NewEventRuleAction(rulesetID string, type_ RuleActionType, ) *EventRuleAction`

NewEventRuleAction instantiates a new EventRuleAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRuleActionWithDefaults

`func NewEventRuleActionWithDefaults() *EventRuleAction`

NewEventRuleActionWithDefaults instantiates a new EventRuleAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRulesetID

`func (o *EventRuleAction) GetRulesetID() string`

GetRulesetID returns the RulesetID field if non-nil, zero value otherwise.

### GetRulesetIDOk

`func (o *EventRuleAction) GetRulesetIDOk() (*string, bool)`

GetRulesetIDOk returns a tuple with the RulesetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetID

`func (o *EventRuleAction) SetRulesetID(v string)`

SetRulesetID sets RulesetID field to given value.


### GetRuleID

`func (o *EventRuleAction) GetRuleID() string`

GetRuleID returns the RuleID field if non-nil, zero value otherwise.

### GetRuleIDOk

`func (o *EventRuleAction) GetRuleIDOk() (*string, bool)`

GetRuleIDOk returns a tuple with the RuleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleID

`func (o *EventRuleAction) SetRuleID(v string)`

SetRuleID sets RuleID field to given value.

### HasRuleID

`func (o *EventRuleAction) HasRuleID() bool`

HasRuleID returns a boolean if a field has been set.

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

### GetType

`func (o *EventRuleAction) GetType() RuleActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventRuleAction) GetTypeOk() (*RuleActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventRuleAction) SetType(v RuleActionType)`

SetType sets Type field to given value.


### GetRequestHeaderModifications

`func (o *EventRuleAction) GetRequestHeaderModifications() RequestHeaderModifications`

GetRequestHeaderModifications returns the RequestHeaderModifications field if non-nil, zero value otherwise.

### GetRequestHeaderModificationsOk

`func (o *EventRuleAction) GetRequestHeaderModificationsOk() (*RequestHeaderModifications, bool)`

GetRequestHeaderModificationsOk returns a tuple with the RequestHeaderModifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaderModifications

`func (o *EventRuleAction) SetRequestHeaderModifications(v RequestHeaderModifications)`

SetRequestHeaderModifications sets RequestHeaderModifications field to given value.

### HasRequestHeaderModifications

`func (o *EventRuleAction) HasRequestHeaderModifications() bool`

HasRequestHeaderModifications returns a boolean if a field has been set.

### GetStatusCode

`func (o *EventRuleAction) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *EventRuleAction) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *EventRuleAction) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *EventRuleAction) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetHeaders

`func (o *EventRuleAction) GetHeaders() []RuleActionHeaderField`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *EventRuleAction) GetHeadersOk() (*[]RuleActionHeaderField, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *EventRuleAction) SetHeaders(v []RuleActionHeaderField)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *EventRuleAction) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBody

`func (o *EventRuleAction) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EventRuleAction) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EventRuleAction) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *EventRuleAction) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


