# EdgeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | [**[]EdgeRequestHeadersInner**](EdgeRequestHeadersInner.md) | Ordered header entries from the request made to your server. Each entry represents one header line. If one header name appears as multiple lines, send each as a separate item in the array.  Headers that contain authentication or session data must still be included, but with with their value set to an empty string. This includes headers like `Authorization` and `Cookie`, but may contain more depending on your specific project, for instance `Proxy-Authenticate` or `X-Api-Key`. Omitting the headers entirely changes the shape of the request and can affect detection. Never forward the real secret values.  Whenever possible, we recommend preserving header order and capitalization to provide the best accuracy, however it’s not a strict requirement if your runtime does not maintain http header order or canonicalizes header names.  | 
**Method** | **string** | The original HTTP method of the request. If supported in your runtime, preserve the original casing. | 
**URL** | **string** | Absolute URL of the request, without a \\#fragment suffix. Only HTTP and HTTPS schemes are supported. | 
**Ipv4Address** | Pointer to **string** | Client IPv4 address observed by your server. | [optional] 
**Ipv6Address** | Pointer to **string** | Client IPv6 address observed by your server. | [optional] 
**LinkedID** | Pointer to **string** | A customer-provided id that was sent with the request. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | A customer-provided value or an object that was sent with the identification request or updated later. | [optional] 


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


