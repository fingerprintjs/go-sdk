# Identification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VisitorID** | **string** | String of 20 characters that uniquely identifies the visitor's browser or mobile device. | 
**Confidence** | Pointer to [**IdentificationConfidence**](IdentificationConfidence.md) |  | [optional] 
**VisitorFound** | **bool** | Attribute represents if a visitor had been identified before. | 
**FirstSeenAt** | Pointer to **int64** | Unix epoch time milliseconds timestamp indicating the time at which this visitor ID was first seen. example: `1758069706642` - Corresponding to Wed Sep 17 2025 00:41:46 GMT+0000  | [optional] 
**LastSeenAt** | Pointer to **int64** | Unix epoch time milliseconds timestamp indicating the time at which this visitor ID was last seen. example: `1758069706642` - Corresponding to Wed Sep 17 2025 00:41:46 GMT+0000  | [optional] 


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


