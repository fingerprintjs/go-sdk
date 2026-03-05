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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


