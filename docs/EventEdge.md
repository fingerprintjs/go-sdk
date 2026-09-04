# EventEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | **string** | Unique identifier of the user's request. The first portion of the event_id is a unix epoch milliseconds timestamp.  | 
**Timestamp** | **int64** | Timestamp of the event with millisecond precision in Unix time. | 
**LinkedID** | Pointer to **string** | A customer-provided id that was sent with the request. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | A customer-provided value or an object that was sent with the identification request or updated later. | [optional] 
**URL** | Pointer to **string** | Page URL from which the request was sent. | [optional] 
**BotInfo** | Pointer to [**BotInfo**](BotInfo.md) |  | [optional] 
**IPInfo** | [**IPInfo**](IPInfo.md) |  | 
**Proxy** | Pointer to **bool** | IP address was used by a public proxy provider or belonged to a known recent residential proxy  | [optional] 
**ProxyConfidence** | Pointer to [**ProxyConfidence**](ProxyConfidence.md) |  | [optional] 
**ProxyDetails** | Pointer to [**ProxyDetails**](ProxyDetails.md) |  | [optional] 
**VPN** | Pointer to **bool** | VPN or other anonymizing service has been used when sending the request.  | [optional] 
**VPNConfidence** | Pointer to [**VPNConfidence**](VPNConfidence.md) |  | [optional] 
**VPNMethods** | Pointer to [**VPNMethods**](VPNMethods.md) |  | [optional] 
**Source** | [**EventSource**](EventSource.md) |  | 


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


