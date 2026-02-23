# \FingerprintAPI

All URIs are relative to *https://api.fpjs.io/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVisitorData**](FingerprintAPI.md#DeleteVisitorData) | **Delete** /visitors/{visitor_id} | Delete data by visitor ID
[**GetEvent**](FingerprintAPI.md#GetEvent) | **Get** /events/{event_id} | Get an event by event ID
[**SearchEvents**](FingerprintAPI.md#SearchEvents) | **Get** /events | Search events
[**UpdateEvent**](FingerprintAPI.md#UpdateEvent) | **Patch** /events/{event_id} | Update an event



## DeleteVisitorData

> DeleteVisitorData(ctx, visitorId).Execute()

Delete data by visitor ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fingerprintjs/go-sdk/v8"
)

func main() {
	visitorId := "visitorId_example" // string | The [visitor ID](https://dev.fingerprint.com/reference/get-function#visitorid) you want to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FingerprintAPI.DeleteVisitorData(context.Background(), visitorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FingerprintAPI.DeleteVisitorData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**visitorId** | **string** | The [visitor ID](https://dev.fingerprint.com/reference/get-function#visitorid) you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVisitorDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvent

> Event GetEvent(ctx, eventId).RulesetId(rulesetId).Execute()

Get an event by event ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fingerprintjs/go-sdk/v8"
)

func main() {
	eventId := "eventId_example" // string | The unique [identifier](https://dev.fingerprint.com/reference/get-function#requestid) of each identification request (`requestId` can be used in its place).
	rulesetId := "rulesetId_example" // string | The ID of the ruleset to evaluate against the event, producing the action to take for this event. The resulting action is returned in the `rule_action` attribute of the response.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FingerprintAPI.GetEvent(context.Background(), eventId).RulesetId(rulesetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FingerprintAPI.GetEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvent`: Event
	fmt.Fprintf(os.Stdout, "Response from `FingerprintAPI.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | The unique [identifier](https://dev.fingerprint.com/reference/get-function#requestid) of each identification request (&#x60;requestId&#x60; can be used in its place). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rulesetId** | **string** | The ID of the ruleset to evaluate against the event, producing the action to take for this event. The resulting action is returned in the &#x60;rule_action&#x60; attribute of the response.  | 

### Return type

[**Event**](Event.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchEvents

> EventSearch SearchEvents(ctx).Limit(limit).PaginationKey(paginationKey).VisitorId(visitorId).Bot(bot).IpAddress(ipAddress).Asn(asn).LinkedId(linkedId).Url(url).BundleId(bundleId).PackageName(packageName).Origin(origin).Start(start).End(end).Reverse(reverse).Suspect(suspect).Vpn(vpn).VirtualMachine(virtualMachine).Tampering(tampering).AntiDetectBrowser(antiDetectBrowser).Incognito(incognito).PrivacySettings(privacySettings).Jailbroken(jailbroken).Frida(frida).FactoryReset(factoryReset).ClonedApp(clonedApp).Emulator(emulator).RootApps(rootApps).VpnConfidence(vpnConfidence).MinSuspectScore(minSuspectScore).DeveloperTools(developerTools).LocationSpoofing(locationSpoofing).MitmAttack(mitmAttack).Proxy(proxy).SdkVersion(sdkVersion).SdkPlatform(sdkPlatform).Environment(environment).ProximityId(proximityId).TotalHits(totalHits).TorNode(torNode).Execute()

Search events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fingerprintjs/go-sdk/v8"
)

func main() {
	limit := int32(10) // int32 | Limit the number of events returned.  (optional) (default to 10)
	paginationKey := "paginationKey_example" // string | Use `pagination_key` to get the next page of results.  When more results are available (e.g., you requested up to 100 results for your query using `limit`, but there are more than 100 events total matching your request), the `pagination_key` field is added to the response. The pagination key is an arbitrary string that should not be interpreted in any way and should be passed as-is. In the following request, use that value in the `pagination_key` parameter to get the next page of results:  1. First request, returning most recent 200 events: `GET api-base-url/events?limit=100` 2. Use `response.pagination_key` to get the next page of results: `GET api-base-url/events?limit=100&pagination_key=1740815825085`  (optional)
	visitorId := "visitorId_example" // string | Unique [visitor identifier](https://dev.fingerprint.com/reference/get-function#visitorid) issued by Fingerprint Identification and all active Smart Signals. Filter for events matching this `visitor_id`.  (optional)
	bot := openapiclient.SearchEventsBot("all") // SearchEventsBot | Filter events by the Bot Detection result, specifically:   `all` - events where any kind of bot was detected.   `good` - events where a good bot was detected.   `bad` - events where a bad bot was detected.   `none` - events where no bot was detected. > Note: When using this parameter, only events with the `bot` property set to a valid value are returned. Events without a `bot` Smart Signal result are left out of the response.  (optional)
	ipAddress := "ipAddress_example" // string | Filter events by IP address or IP range (if CIDR notation is used). If CIDR notation is not used, a /32 for IPv4 or /128 for IPv6 is assumed. Examples of range based queries: 10.0.0.0/24, 192.168.0.1/32  (optional)
	asn := "asn_example" // string | Filter events by the ASN associated with the event's IP address. This corresponds to the `ip_info.(v4|v6).asn` property in the response.  (optional)
	linkedId := "linkedId_example" // string | Filter events by your custom identifier.  You can use [linked Ids](https://dev.fingerprint.com/reference/get-function#linkedid) to associate identification requests with your own identifier, for example, session Id, purchase Id, or transaction Id. You can then use this `linked_id` parameter to retrieve all events associated with your custom identifier.  (optional)
	url := "url_example" // string | Filter events by the URL (`url` property) associated with the event.  (optional)
	bundleId := "bundleId_example" // string | Filter events by the Bundle ID (iOS) associated with the event.  (optional)
	packageName := "packageName_example" // string | Filter events by the Package Name (Android) associated with the event.  (optional)
	origin := "origin_example" // string | Filter events by the origin field of the event. This is applicable to web events only (e.g., https://example.com)  (optional)
	start := int64(789) // int64 | Filter events with a timestamp greater than the start time, in Unix time (milliseconds).  (optional)
	end := int64(789) // int64 | Filter events with a timestamp smaller than the end time, in Unix time (milliseconds).  (optional)
	reverse := true // bool | Sort events in reverse timestamp order.  (optional)
	suspect := true // bool | Filter events previously tagged as suspicious via the [Update API](https://dev.fingerprint.com/reference/updateevent). > Note: When using this parameter, only events with the `suspect` property explicitly set to `true` or `false` are returned. Events with undefined `suspect` property are left out of the response.  (optional)
	vpn := true // bool | Filter events by VPN Detection result. > Note: When using this parameter, only events with the `vpn` property set to `true` or `false` are returned. Events without a `vpn` Smart Signal result are left out of the response.  (optional)
	virtualMachine := true // bool | Filter events by Virtual Machine Detection result. > Note: When using this parameter, only events with the `virtual_machine` property set to `true` or `false` are returned. Events without a `virtual_machine` Smart Signal result are left out of the response.  (optional)
	tampering := true // bool | Filter events by Browser Tampering Detection result. > Note: When using this parameter, only events with the `tampering.result` property set to `true` or `false` are returned. Events without a `tampering` Smart Signal result are left out of the response.  (optional)
	antiDetectBrowser := true // bool | Filter events by Anti-detect Browser Detection result. > Note: When using this parameter, only events with the `tampering.anti_detect_browser` property set to `true` or `false` are returned. Events without a `tampering` Smart Signal result are left out of the response.  (optional)
	incognito := true // bool | Filter events by Browser Incognito Detection result. > Note: When using this parameter, only events with the `incognito` property set to `true` or `false` are returned. Events without an `incognito` Smart Signal result are left out of the response.  (optional)
	privacySettings := true // bool | Filter events by Privacy Settings Detection result. > Note: When using this parameter, only events with the `privacy_settings` property set to `true` or `false` are returned. Events without a `privacy_settings` Smart Signal result are left out of the response.  (optional)
	jailbroken := true // bool | Filter events by Jailbroken Device Detection result. > Note: When using this parameter, only events with the `jailbroken` property set to `true` or `false` are returned. Events without a `jailbroken` Smart Signal result are left out of the response.  (optional)
	frida := true // bool | Filter events by Frida Detection result. > Note: When using this parameter, only events with the `frida` property set to `true` or `false` are returned. Events without a `frida` Smart Signal result are left out of the response.  (optional)
	factoryReset := true // bool | Filter events by Factory Reset Detection result. > Note: When using this parameter, only events with a `factory_reset` time. Events without a `factory_reset` Smart Signal result are left out of the response.  (optional)
	clonedApp := true // bool | Filter events by Cloned App Detection result. > Note: When using this parameter, only events with the `cloned_app` property set to `true` or `false` are returned. Events without a `cloned_app` Smart Signal result are left out of the response.  (optional)
	emulator := true // bool | Filter events by Android Emulator Detection result. > Note: When using this parameter, only events with the `emulator` property set to `true` or `false` are returned. Events without an `emulator` Smart Signal result are left out of the response.  (optional)
	rootApps := true // bool | Filter events by Rooted Device Detection result. > Note: When using this parameter, only events with the `root_apps` property set to `true` or `false` are returned. Events without a `root_apps` Smart Signal result are left out of the response.  (optional)
	vpnConfidence := openapiclient.SearchEventsVpnConfidence("high") // SearchEventsVpnConfidence | Filter events by VPN Detection result confidence level. `high` - events with high VPN Detection confidence. `medium` - events with medium VPN Detection confidence. `low` - events with low VPN Detection confidence. > Note: When using this parameter, only events with the `vpn.confidence` property set to a valid value are returned. Events without a `vpn` Smart Signal result are left out of the response.  (optional)
	minSuspectScore := float32(3.4) // float32 | Filter events with Suspect Score result above a provided minimum threshold. > Note: When using this parameter, only events where the `suspect_score` property set to a value exceeding your threshold are returned. Events without a `suspect_score` Smart Signal result are left out of the response.  (optional)
	developerTools := true // bool | Filter events by Developer Tools detection result. > Note: When using this parameter, only events with the `developer_tools` property set to `true` or `false` are returned. Events without a `developer_tools` Smart Signal result are left out of the response.  (optional)
	locationSpoofing := true // bool | Filter events by Location Spoofing detection result. > Note: When using this parameter, only events with the `location_spoofing` property set to `true` or `false` are returned. Events without a `location_spoofing` Smart Signal result are left out of the response.  (optional)
	mitmAttack := true // bool | Filter events by MITM (Man-in-the-Middle) Attack detection result. > Note: When using this parameter, only events with the `mitm_attack` property set to `true` or `false` are returned. Events without a `mitm_attack` Smart Signal result are left out of the response.  (optional)
	proxy := true // bool | Filter events by Proxy detection result. > Note: When using this parameter, only events with the `proxy` property set to `true` or `false` are returned. Events without a `proxy` Smart Signal result are left out of the response.  (optional)
	sdkVersion := "sdkVersion_example" // string | Filter events by a specific SDK version associated with the identification event (`sdk.version` property). Example: `3.11.14`  (optional)
	sdkPlatform := openapiclient.SearchEventsSdkPlatform("js") // SearchEventsSdkPlatform | Filter events by the SDK Platform associated with the identification event (`sdk.platform` property) . `js` - Javascript agent (Web). `ios` - Apple iOS based devices. `android` - Android based devices.  (optional)
	environment := []string{"Inner_example"} // []string | Filter for events by providing one or more environment IDs (`environment_id` property).  (optional)
	proximityId := "proximityId_example" // string | Filter events by the most precise Proximity ID provided by default. > Note: When using this parameter, only events with the `proximity.id` property matching the provided ID are returned. Events without a `proximity` result are left out of the response.  (optional)
	totalHits := int64(789) // int64 | When set, the response will include a `total_hits` property with a count of total query matches across all pages, up to the specified limit.  (optional)
	torNode := true // bool | Filter events by Tor Node detection result. > Note: When using this parameter, only events with the `tor_node` property set to `true` or `false` are returned. Events without a `tor_node` detection result are left out of the response.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FingerprintAPI.SearchEvents(context.Background()).Limit(limit).PaginationKey(paginationKey).VisitorId(visitorId).Bot(bot).IpAddress(ipAddress).Asn(asn).LinkedId(linkedId).Url(url).BundleId(bundleId).PackageName(packageName).Origin(origin).Start(start).End(end).Reverse(reverse).Suspect(suspect).Vpn(vpn).VirtualMachine(virtualMachine).Tampering(tampering).AntiDetectBrowser(antiDetectBrowser).Incognito(incognito).PrivacySettings(privacySettings).Jailbroken(jailbroken).Frida(frida).FactoryReset(factoryReset).ClonedApp(clonedApp).Emulator(emulator).RootApps(rootApps).VpnConfidence(vpnConfidence).MinSuspectScore(minSuspectScore).DeveloperTools(developerTools).LocationSpoofing(locationSpoofing).MitmAttack(mitmAttack).Proxy(proxy).SdkVersion(sdkVersion).SdkPlatform(sdkPlatform).Environment(environment).ProximityId(proximityId).TotalHits(totalHits).TorNode(torNode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FingerprintAPI.SearchEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchEvents`: EventSearch
	fmt.Fprintf(os.Stdout, "Response from `FingerprintAPI.SearchEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of events returned.  | [default to 10]
 **paginationKey** | **string** | Use &#x60;pagination_key&#x60; to get the next page of results.  When more results are available (e.g., you requested up to 100 results for your query using &#x60;limit&#x60;, but there are more than 100 events total matching your request), the &#x60;pagination_key&#x60; field is added to the response. The pagination key is an arbitrary string that should not be interpreted in any way and should be passed as-is. In the following request, use that value in the &#x60;pagination_key&#x60; parameter to get the next page of results:  1. First request, returning most recent 200 events: &#x60;GET api-base-url/events?limit&#x3D;100&#x60; 2. Use &#x60;response.pagination_key&#x60; to get the next page of results: &#x60;GET api-base-url/events?limit&#x3D;100&amp;pagination_key&#x3D;1740815825085&#x60;  | 
 **visitorId** | **string** | Unique [visitor identifier](https://dev.fingerprint.com/reference/get-function#visitorid) issued by Fingerprint Identification and all active Smart Signals. Filter for events matching this &#x60;visitor_id&#x60;.  | 
 **bot** | [**SearchEventsBot**](SearchEventsBot.md) | Filter events by the Bot Detection result, specifically:   &#x60;all&#x60; - events where any kind of bot was detected.   &#x60;good&#x60; - events where a good bot was detected.   &#x60;bad&#x60; - events where a bad bot was detected.   &#x60;none&#x60; - events where no bot was detected. &gt; Note: When using this parameter, only events with the &#x60;bot&#x60; property set to a valid value are returned. Events without a &#x60;bot&#x60; Smart Signal result are left out of the response.  | 
 **ipAddress** | **string** | Filter events by IP address or IP range (if CIDR notation is used). If CIDR notation is not used, a /32 for IPv4 or /128 for IPv6 is assumed. Examples of range based queries: 10.0.0.0/24, 192.168.0.1/32  | 
 **asn** | **string** | Filter events by the ASN associated with the event&#39;s IP address. This corresponds to the &#x60;ip_info.(v4|v6).asn&#x60; property in the response.  | 
 **linkedId** | **string** | Filter events by your custom identifier.  You can use [linked Ids](https://dev.fingerprint.com/reference/get-function#linkedid) to associate identification requests with your own identifier, for example, session Id, purchase Id, or transaction Id. You can then use this &#x60;linked_id&#x60; parameter to retrieve all events associated with your custom identifier.  | 
 **url** | **string** | Filter events by the URL (&#x60;url&#x60; property) associated with the event.  | 
 **bundleId** | **string** | Filter events by the Bundle ID (iOS) associated with the event.  | 
 **packageName** | **string** | Filter events by the Package Name (Android) associated with the event.  | 
 **origin** | **string** | Filter events by the origin field of the event. This is applicable to web events only (e.g., https://example.com)  | 
 **start** | **int64** | Filter events with a timestamp greater than the start time, in Unix time (milliseconds).  | 
 **end** | **int64** | Filter events with a timestamp smaller than the end time, in Unix time (milliseconds).  | 
 **reverse** | **bool** | Sort events in reverse timestamp order.  | 
 **suspect** | **bool** | Filter events previously tagged as suspicious via the [Update API](https://dev.fingerprint.com/reference/updateevent). &gt; Note: When using this parameter, only events with the &#x60;suspect&#x60; property explicitly set to &#x60;true&#x60; or &#x60;false&#x60; are returned. Events with undefined &#x60;suspect&#x60; property are left out of the response.  | 
 **vpn** | **bool** | Filter events by VPN Detection result. &gt; Note: When using this parameter, only events with the &#x60;vpn&#x60; property set to &#x60;true&#x60; or &#x60;false&#x60; are returned. Events without a &#x60;vpn&#x60; Smart Signal result are left out of the response.  | 
 **virtualMachine** | **bool** | Filter events by Virtual Machine Detection result. &gt; Note: When using this parameter, only events with the &#x60;virtual_machine&#x60; property set to &#x60;true&#x60; or &#x60;false&#x60; are returned. Events without a &#x60;virtual_machine&#x60; Smart Signal result are left out of the response.  | 
 **tampering** | **bool** | Filter events by Browser Tampering Detection result. &gt; Note: When using this parameter, only events with the &#x60;tampering.result&#x60; property set to &#x60;true&#x60; or &#x60;false&#x60; are returned. Events without a &#x60;tampering&#x60; Smart Signal result are left out of the response.  | 
 **antiDetectBrowser** | **bool** | Filter events by Anti-detect Browser Detection result. &gt; Note: When using this parameter, only events with the &#x60;tampering.anti_detect_browser&#x60; property set to &#x60;true&#x60; or &#x60;false&#x60; are returned. Events without a &#x60;tampering&#x60; Smart Signal result are left out of the response.  | 
 **incognito** | **bool** | Filter events by Browser Incognito Detection result. &gt; Note: When using this parameter, only events with the &#x60;incognito&#x60; property set to &#x60;true&#x60; or &#x60;false&#x60; are returned. Events without an &#x60;incognito&#x60; Smart Signal result are left out of the response.  | 
 **privacySettings** | **bool** | Filter events by Privacy Settings Detection result. &gt; Note: When using this parameter, only events with the &#x60;privacy_settings&#x60; property set to &#x60;true&#x60; or &#x60;false&#x60; are returned. Events without a &#x60;privacy_settings&#x60; Smart Signal result are left out of the response.  | 
 **jailbroken** | **bool** | Filter events by Jailbroken Device Detection result. &gt; Note: When using this parameter, only events with the &#x60;jailbroken&#x60; property set to &#x60;true&#x60; or &#x60;false&#x60; are returned. Events without a &#x60;jailbroken&#x60; Smart Signal result are left out of the response.  | 
 **frida** | **bool** | Filter events by Frida Detection result. &gt; Note: When using this parameter, only events with the &#x60;frida&#x60; property set to &#x60;true&#x60; or &#x60;false&#x60; are returned. Events without a &#x60;frida&#x60; Smart Signal result are left out of the response.  | 
 **factoryReset** | **bool** | Filter events by Factory Reset Detection result. &gt; Note: When using this parameter, only events with a &#x60;factory_reset&#x60; time. Events without a &#x60;factory_reset&#x60; Smart Signal result are left out of the response.  | 
 **clonedApp** | **bool** | Filter events by Cloned App Detection result. &gt; Note: When using this parameter, only events with the &#x60;cloned_app&#x60; property set to &#x60;true&#x60; or &#x60;false&#x60; are returned. Events without a &#x60;cloned_app&#x60; Smart Signal result are left out of the response.  | 
 **emulator** | **bool** | Filter events by Android Emulator Detection result. &gt; Note: When using this parameter, only events with the &#x60;emulator&#x60; property set to &#x60;true&#x60; or &#x60;false&#x60; are returned. Events without an &#x60;emulator&#x60; Smart Signal result are left out of the response.  | 
 **rootApps** | **bool** | Filter events by Rooted Device Detection result. &gt; Note: When using this parameter, only events with the &#x60;root_apps&#x60; property set to &#x60;true&#x60; or &#x60;false&#x60; are returned. Events without a &#x60;root_apps&#x60; Smart Signal result are left out of the response.  | 
 **vpnConfidence** | [**SearchEventsVpnConfidence**](SearchEventsVpnConfidence.md) | Filter events by VPN Detection result confidence level. &#x60;high&#x60; - events with high VPN Detection confidence. &#x60;medium&#x60; - events with medium VPN Detection confidence. &#x60;low&#x60; - events with low VPN Detection confidence. &gt; Note: When using this parameter, only events with the &#x60;vpn.confidence&#x60; property set to a valid value are returned. Events without a &#x60;vpn&#x60; Smart Signal result are left out of the response.  | 
 **minSuspectScore** | **float32** | Filter events with Suspect Score result above a provided minimum threshold. &gt; Note: When using this parameter, only events where the &#x60;suspect_score&#x60; property set to a value exceeding your threshold are returned. Events without a &#x60;suspect_score&#x60; Smart Signal result are left out of the response.  | 
 **developerTools** | **bool** | Filter events by Developer Tools detection result. &gt; Note: When using this parameter, only events with the &#x60;developer_tools&#x60; property set to &#x60;true&#x60; or &#x60;false&#x60; are returned. Events without a &#x60;developer_tools&#x60; Smart Signal result are left out of the response.  | 
 **locationSpoofing** | **bool** | Filter events by Location Spoofing detection result. &gt; Note: When using this parameter, only events with the &#x60;location_spoofing&#x60; property set to &#x60;true&#x60; or &#x60;false&#x60; are returned. Events without a &#x60;location_spoofing&#x60; Smart Signal result are left out of the response.  | 
 **mitmAttack** | **bool** | Filter events by MITM (Man-in-the-Middle) Attack detection result. &gt; Note: When using this parameter, only events with the &#x60;mitm_attack&#x60; property set to &#x60;true&#x60; or &#x60;false&#x60; are returned. Events without a &#x60;mitm_attack&#x60; Smart Signal result are left out of the response.  | 
 **proxy** | **bool** | Filter events by Proxy detection result. &gt; Note: When using this parameter, only events with the &#x60;proxy&#x60; property set to &#x60;true&#x60; or &#x60;false&#x60; are returned. Events without a &#x60;proxy&#x60; Smart Signal result are left out of the response.  | 
 **sdkVersion** | **string** | Filter events by a specific SDK version associated with the identification event (&#x60;sdk.version&#x60; property). Example: &#x60;3.11.14&#x60;  | 
 **sdkPlatform** | [**SearchEventsSdkPlatform**](SearchEventsSdkPlatform.md) | Filter events by the SDK Platform associated with the identification event (&#x60;sdk.platform&#x60; property) . &#x60;js&#x60; - Javascript agent (Web). &#x60;ios&#x60; - Apple iOS based devices. &#x60;android&#x60; - Android based devices.  | 
 **environment** | **[]string** | Filter for events by providing one or more environment IDs (&#x60;environment_id&#x60; property).  | 
 **proximityId** | **string** | Filter events by the most precise Proximity ID provided by default. &gt; Note: When using this parameter, only events with the &#x60;proximity.id&#x60; property matching the provided ID are returned. Events without a &#x60;proximity&#x60; result are left out of the response.  | 
 **totalHits** | **int64** | When set, the response will include a &#x60;total_hits&#x60; property with a count of total query matches across all pages, up to the specified limit.  | 
 **torNode** | **bool** | Filter events by Tor Node detection result. &gt; Note: When using this parameter, only events with the &#x60;tor_node&#x60; property set to &#x60;true&#x60; or &#x60;false&#x60; are returned. Events without a &#x60;tor_node&#x60; detection result are left out of the response.  | 

### Return type

[**EventSearch**](EventSearch.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEvent

> UpdateEvent(ctx, eventId).EventUpdate(eventUpdate).Execute()

Update an event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fingerprintjs/go-sdk/v8"
)

func main() {
	eventId := "eventId_example" // string | The unique event [identifier](https://dev.fingerprint.com/reference/get-function#event_id).
	eventUpdate := *openapiclient.NewEventUpdate() // EventUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FingerprintAPI.UpdateEvent(context.Background(), eventId).EventUpdate(eventUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FingerprintAPI.UpdateEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | The unique event [identifier](https://dev.fingerprint.com/reference/get-function#event_id). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventUpdate** | [**EventUpdate**](EventUpdate.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

