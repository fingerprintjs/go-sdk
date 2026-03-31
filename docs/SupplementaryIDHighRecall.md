# SupplementaryIDHighRecall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VisitorID** | **string** | The High Recall identifier for the visitor's browser. It is an alphanumeric string with a maximum length of 25 characters. | 
**VisitorFound** | **bool** | True if this is a returning browser and has been previously identified. Otherwise, false. | 
**Confidence** | Pointer to [**IdentificationConfidence**](IdentificationConfidence.md) |  | [optional] 
**FirstSeenAt** | Pointer to **int64** | Unix epoch timestamp (in milliseconds) indicating when the browser was first identified. example: `1758069706642` - Corresponding to Wed Sep 17 2025 00:41:46 GMT+0000  | [optional] 
**LastSeenAt** | Pointer to **int64** | Unix epoch timestamp (in milliseconds) corresponding to the most recent visit by this browser. example: `1758069706642` - Corresponding to Wed Sep 17 2025 00:41:46 GMT+0000  | [optional] 


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


