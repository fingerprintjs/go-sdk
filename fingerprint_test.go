package fingerprint_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	fingerprint "github.com/fingerprintjs/go-sdk/v8"
	"github.com/fingerprintjs/go-sdk/v8/internal/openapi"
)

func ExampleNew() {
	_ = fingerprint.New(fingerprint.WithAPIKey("SECRET_API_KEY"))
}

func ExampleWithAPIKey() {
	_ = fingerprint.New(fingerprint.WithAPIKey("SECRET_API_KEY"))
}

func ExampleWithHTTPClient() {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	_ = fingerprint.New(fingerprint.WithHTTPClient(httpClient))
}

func ExampleWithRegion() {
	// For EU
	_ = fingerprint.New(fingerprint.WithRegion(fingerprint.RegionEU))

	// For Asia
	_ = fingerprint.New(fingerprint.WithRegion(fingerprint.RegionAsia))
}

func ExampleWithDebug() {
	_ = fingerprint.New(fingerprint.WithDebug(true))
}

func ExampleWithBaseURL() {
	_ = fingerprint.New(fingerprint.WithBaseURL("https://example.com"))
}

type FingerprintApiImpl struct{}

func (f FingerprintApiImpl) DeleteVisitorData(visitorID string) openapi.ApiDeleteVisitorDataRequest {
	//TODO implement me
	panic("implement me")
}

func (f FingerprintApiImpl) DeleteVisitorDataExecute(ctx context.Context, r openapi.ApiDeleteVisitorDataRequest) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (f FingerprintApiImpl) GetEvent(eventID string) openapi.ApiGetEventRequest {
	//TODO implement me
	panic("implement me")
}

func (f FingerprintApiImpl) GetEventExecute(ctx context.Context, r openapi.ApiGetEventRequest) (*openapi.Event, *http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (f FingerprintApiImpl) SearchEvents() openapi.ApiSearchEventsRequest {
	//TODO implement me
	panic("implement me")
}

func (f FingerprintApiImpl) SearchEventsExecute(ctx context.Context, r openapi.ApiSearchEventsRequest) (*openapi.EventSearch, *http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (f FingerprintApiImpl) UpdateEvent(eventID string) openapi.ApiUpdateEventRequest {
	//TODO implement me
	panic("implement me")
}

func (f FingerprintApiImpl) UpdateEventExecute(ctx context.Context, r openapi.ApiUpdateEventRequest) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func ExampleWithFingerprintAPI() {
	_ = fingerprint.New(fingerprint.WithFingerprintAPI(FingerprintApiImpl{}))
}

func ExampleClient_GetEvent() {
	eventID := "eventID_example"
	client := fingerprint.New(fingerprint.WithAPIKey("SECRET_API_KEY"))

	response, httpRes, err := client.GetEvent(context.TODO(), eventID)

	fmt.Printf("HTTP Response from `GetEvent`: %v\n", httpRes)

	if err != nil {
		fmt.Printf("Error when calling `GetEvent`: %v\n", err)
		fmt.Printf("Full HTTP response: %v\n", httpRes)
	}

	if response.Bot != nil {
		fmt.Printf("Got response with Botd: %v \n", response.Bot)
	}

	if response.Identification != nil {
		stringResponse, _ := json.Marshal(response.Identification)
		fmt.Printf("Got response with Identification: %s \n", string(stringResponse))
	}
}

func ExampleClient_GetEvent_With_Ruleset() {
	client := fingerprint.New(fingerprint.WithAPIKey("SECRET_API_KEY"))

	eventID := "eventID_example"
	rulesetID := "rulesetID_example"

	response, httpRes, err := client.GetEvent(context.TODO(), eventID, fingerprint.WithRulesetID(rulesetID))

	fmt.Printf("HTTP Response from `GetEvent`: %v\n", httpRes)

	if errResp, ok := fingerprint.AsErrorResponse(err); ok {
		switch errResp.Error.Code {
		case fingerprint.ErrorCodeEvent_not_found:
			fmt.Println("event not found")
		case fingerprint.ErrorCodeSecret_api_key_not_found:
			fmt.Println("secret api key not found")
		case fingerprint.ErrorCodeRuleset_not_found:
			fmt.Println("ruleset not found")
		default:
			fmt.Printf("unexpected error: %s\n", errResp.Error.Message)
		}
	} else {
		fmt.Printf("other error occurred: %v\n", err.Error())
	}

	if response.RuleAction != nil {
		if response.RuleAction.EventRuleActionAllow != nil {
			fmt.Printf("action is allow, rule id: %s, rule expression: %s\n", *response.RuleAction.EventRuleActionAllow.RuleID, *response.RuleAction.EventRuleActionAllow.RuleExpression)
			if response.RuleAction.EventRuleActionAllow.RequestHeaderModifications != nil {
				fmt.Printf("request header modifications to set %v, to append %v, to remove %v\n",
					response.RuleAction.EventRuleActionAllow.RequestHeaderModifications.Set,
					response.RuleAction.EventRuleActionAllow.RequestHeaderModifications.Append,
					response.RuleAction.EventRuleActionAllow.RequestHeaderModifications.Remove)
			}
		} else if response.RuleAction.EventRuleActionBlock != nil {
			fmt.Printf("action is block. rule id: %s, rule expression: %s, body %s statusCode: %d headers: %v\n",
				*response.RuleAction.EventRuleActionBlock.RuleID,
				*response.RuleAction.EventRuleActionBlock.RuleExpression,
				*response.RuleAction.EventRuleActionBlock.Body,
				*response.RuleAction.EventRuleActionBlock.StatusCode,
				response.RuleAction.EventRuleActionBlock.Headers)
		} else {
			fmt.Println("action is unexpected (please make sure library is at the latest version)")
		}
	}
}

func ExampleClient_SearchEvents() {
	limit := int32(10)                                         // int32 | Limit the number of events returned.  (optional) (default to 10)
	paginationKey := "paginationKey_example"                   // string | Use `pagination_key` to get the next page of results.  When more results are available (e.g., you requested up to 100 results for your query using `limit`, but there are more than 100 events total matching your request), the `pagination_key` field is added to the response. The pagination key is an arbitrary string that should not be interpreted in any way and should be passed as-is. In the following request, use that value in the `pagination_key` parameter to get the next page of results:  1. First request, returning most recent 200 events: `GET api-base-url/events?limit=100` 2. Use `response.pagination_key` to get the next page of results: `GET api-base-url/events?limit=100&pagination_key=1740815825085`  (optional)
	visitorID := "visitorID_example"                           // string | Unique [visitor identifier](https://dev.fingerprint.com/reference/get-function#visitorid) issued by Fingerprint Identification and all active Smart Signals. Filter for events matching this `visitor_id`.  (optional)
	bot := fingerprint.SearchEventsBotAll                      // SearchEventsBot | Filter events by the Bot Detection result, specifically:   `all` - events where any kind of bot was detected.   `good` - events where a good bot was detected.   `bad` - events where a bad bot was detected.   `none` - events where no bot was detected. > Note: When using this parameter, only events with the `bot` property set to a valid value are returned. Events without a `bot` Smart Signal result are left out of the response.  (optional)
	iPAddress := "iPAddress_example"                           // string | Filter events by IP address or IP range (if CIDR notation is used). If CIDR notation is not used, a /32 for IPv4 or /128 for IPv6 is assumed. Examples of range based queries: 10.0.0.0/24, 192.168.0.1/32  (optional)
	asn := "asn_example"                                       // string | Filter events by the ASN associated with the event's IP address. This corresponds to the `ip_info.(v4|v6).asn` property in the response.  (optional)
	linkedID := "linkedID_example"                             // string | Filter events by your custom identifier.  You can use [linked Ids](https://dev.fingerprint.com/reference/get-function#linkedid) to associate identification requests with your own identifier, for example, session Id, purchase Id, or transaction Id. You can then use this `linked_id` parameter to retrieve all events associated with your custom identifier.  (optional)
	uRL := "uRL_example"                                       // string | Filter events by the URL (`url` property) associated with the event.  (optional)
	bundleID := "bundleID_example"                             // string | Filter events by the Bundle ID (iOS) associated with the event.  (optional)
	packageName := "packageName_example"                       // string | Filter events by the Package Name (Android) associated with the event.  (optional)
	origin := "origin_example"                                 // string | Filter events by the origin field of the event. This is applicable to web events only (e.g., https://example.com)  (optional)
	start := int64(789)                                        // int64 | Filter events with a timestamp greater than the start time, in Unix time (milliseconds).  (optional)
	end := int64(789)                                          // int64 | Filter events with a timestamp smaller than the end time, in Unix time (milliseconds).  (optional)
	reverse := true                                            // bool | Sort events in reverse timestamp order.  (optional)
	suspect := true                                            // bool | Filter events previously tagged as suspicious via the [Update API](https://dev.fingerprint.com/reference/updateevent). > Note: When using this parameter, only events with the `suspect` property explicitly set to `true` or `false` are returned. Events with undefined `suspect` property are left out of the response.  (optional)
	vPN := true                                                // bool | Filter events by VPN Detection result. > Note: When using this parameter, only events with the `vpn` property set to `true` or `false` are returned. Events without a `vpn` Smart Signal result are left out of the response.  (optional)
	virtualMachine := true                                     // bool | Filter events by Virtual Machine Detection result. > Note: When using this parameter, only events with the `virtual_machine` property set to `true` or `false` are returned. Events without a `virtual_machine` Smart Signal result are left out of the response.  (optional)
	tampering := true                                          // bool | Filter events by Browser Tampering Detection result. > Note: When using this parameter, only events with the `tampering.result` property set to `true` or `false` are returned. Events without a `tampering` Smart Signal result are left out of the response.  (optional)
	antiDetectBrowser := true                                  // bool | Filter events by Anti-detect Browser Detection result. > Note: When using this parameter, only events with the `tampering.anti_detect_browser` property set to `true` or `false` are returned. Events without a `tampering` Smart Signal result are left out of the response.  (optional)
	incognito := true                                          // bool | Filter events by Browser Incognito Detection result. > Note: When using this parameter, only events with the `incognito` property set to `true` or `false` are returned. Events without an `incognito` Smart Signal result are left out of the response.  (optional)
	privacySettings := true                                    // bool | Filter events by Privacy Settings Detection result. > Note: When using this parameter, only events with the `privacy_settings` property set to `true` or `false` are returned. Events without a `privacy_settings` Smart Signal result are left out of the response.  (optional)
	jailbroken := true                                         // bool | Filter events by Jailbroken Device Detection result. > Note: When using this parameter, only events with the `jailbroken` property set to `true` or `false` are returned. Events without a `jailbroken` Smart Signal result are left out of the response.  (optional)
	frida := true                                              // bool | Filter events by Frida Detection result. > Note: When using this parameter, only events with the `frida` property set to `true` or `false` are returned. Events without a `frida` Smart Signal result are left out of the response.  (optional)
	factoryReset := true                                       // bool | Filter events by Factory Reset Detection result. > Note: When using this parameter, only events with a `factory_reset` time. Events without a `factory_reset` Smart Signal result are left out of the response.  (optional)
	clonedApp := true                                          // bool | Filter events by Cloned App Detection result. > Note: When using this parameter, only events with the `cloned_app` property set to `true` or `false` are returned. Events without a `cloned_app` Smart Signal result are left out of the response.  (optional)
	emulator := true                                           // bool | Filter events by Android Emulator Detection result. > Note: When using this parameter, only events with the `emulator` property set to `true` or `false` are returned. Events without an `emulator` Smart Signal result are left out of the response.  (optional)
	rootApps := true                                           // bool | Filter events by Rooted Device Detection result. > Note: When using this parameter, only events with the `root_apps` property set to `true` or `false` are returned. Events without a `root_apps` Smart Signal result are left out of the response.  (optional)
	vPNConfidence := fingerprint.SearchEventsVPNConfidenceHigh // SearchEventsVPNConfidence | Filter events by VPN Detection result confidence level. `high` - events with high VPN Detection confidence. `medium` - events with medium VPN Detection confidence. `low` - events with low VPN Detection confidence. > Note: When using this parameter, only events with the `vpn.confidence` property set to a valid value are returned. Events without a `vpn` Smart Signal result are left out of the response.  (optional)
	minSuspectScore := float32(3.4)                            // float32 | Filter events with Suspect Score result above a provided minimum threshold. > Note: When using this parameter, only events where the `suspect_score` property set to a value exceeding your threshold are returned. Events without a `suspect_score` Smart Signal result are left out of the response.  (optional)
	developerTools := true                                     // bool | Filter events by Developer Tools detection result. > Note: When using this parameter, only events with the `developer_tools` property set to `true` or `false` are returned. Events without a `developer_tools` Smart Signal result are left out of the response.  (optional)
	locationSpoofing := true                                   // bool | Filter events by Location Spoofing detection result. > Note: When using this parameter, only events with the `location_spoofing` property set to `true` or `false` are returned. Events without a `location_spoofing` Smart Signal result are left out of the response.  (optional)
	mITMAttack := true                                         // bool | Filter events by MITM (Man-in-the-Middle) Attack detection result. > Note: When using this parameter, only events with the `mitm_attack` property set to `true` or `false` are returned. Events without a `mitm_attack` Smart Signal result are left out of the response.  (optional)
	proxy := true                                              // bool | Filter events by Proxy detection result. > Note: When using this parameter, only events with the `proxy` property set to `true` or `false` are returned. Events without a `proxy` Smart Signal result are left out of the response.  (optional)
	sDKVersion := "sDKVersion_example"                         // string | Filter events by a specific SDK version associated with the identification event (`sdk.version` property). Example: `3.11.14`  (optional)
	sDKPlatform := fingerprint.SearchEventsSDKPlatformJs       // SearchEventsSDKPlatform | Filter events by the SDK Platform associated with the identification event (`sdk.platform` property) . `js` - Javascript agent (Web). `ios` - Apple iOS based devices. `android` - Android based devices.  (optional)
	environment := []string{"Inner_example"}                   // []string | Filter for events by providing one or more environment IDs (`environment_id` property).  (optional)
	proximityID := "proximityID_example"                       // string | Filter events by the most precise Proximity ID provided by default. > Note: When using this parameter, only events with the `proximity.id` property matching the provided ID are returned. Events without a `proximity` result are left out of the response.  (optional)
	totalHits := int64(789)                                    // int64 | When set, the response will include a `total_hits` property with a count of total query matches across all pages, up to the specified limit.  (optional)
	torNode := true                                            // bool | Filter events by Tor Node detection result. > Note: When using this parameter, only events with the `tor_node` property set to `true` or `false` are returned. Events without a `tor_node` detection result are left out of the response.  (optional)

	client := fingerprint.New(fingerprint.WithAPIKey("SECRET_API_KEY"))
	searchReq := fingerprint.NewSearchEventsRequest().
		Limit(limit).
		PaginationKey(paginationKey).
		VisitorID(visitorID).
		Bot(bot).
		IPAddress(iPAddress).
		Asn(asn).
		LinkedID(linkedID).
		URL(uRL).
		BundleID(bundleID).
		PackageName(packageName).
		Origin(origin).
		Start(start).
		End(end).
		Reverse(reverse).
		Suspect(suspect).
		VPN(vPN).
		VirtualMachine(virtualMachine).
		Tampering(tampering).
		AntiDetectBrowser(antiDetectBrowser).
		Incognito(incognito).
		PrivacySettings(privacySettings).
		Jailbroken(jailbroken).
		Frida(frida).
		FactoryReset(factoryReset).
		ClonedApp(clonedApp).
		Emulator(emulator).
		RootApps(rootApps).
		VPNConfidence(vPNConfidence).
		MinSuspectScore(minSuspectScore).
		DeveloperTools(developerTools).
		LocationSpoofing(locationSpoofing).
		MITMAttack(mITMAttack).
		Proxy(proxy).
		SDKVersion(sDKVersion).
		SDKPlatform(sDKPlatform).
		Environment(environment).
		ProximityID(proximityID).
		TotalHits(totalHits).
		TorNode(torNode)
	resp, httpRes, err := client.SearchEvents(context.TODO(), searchReq)

	fmt.Printf("Response from `SearchEvents`: %v\n", resp)
	fmt.Printf("HTTP Response from `SearchEvents`: %v\n", httpRes)

	if errResp, ok := fingerprint.AsErrorResponse(err); ok {
		switch errResp.Error.Code {
		case fingerprint.ErrorCodeEvent_not_found:
			fmt.Println("event not found")
		case fingerprint.ErrorCodeSecret_api_key_not_found:
			fmt.Println("secret api key not found")
		case fingerprint.ErrorCodeState_not_ready:
			fmt.Println("event is not ready yet for update, please try again later")
		default:
			fmt.Printf("unexpected error: %s\n", errResp.Error.Message)
		}
	} else {
		fmt.Printf("other error occurred: %v\n", err.Error())
	}
}

func ExampleClient_UpdateEvent() {
	tags := map[string]interface{}{
		"key": "value",
	}
	suspect := false
	linkedID := "new_linked_id"
	eventUpdateReq := fingerprint.EventUpdate{
		Suspect:  &suspect,
		LinkedID: &linkedID,
		Tags:     tags,
	}

	eventID := "eventID_example" // string | The unique event [identifier](https://dev.fingerprint.com/reference/get-function#event_id).

	client := fingerprint.New(fingerprint.WithAPIKey("SECRET_API_KEY"))
	httpRes, err := client.UpdateEvent(context.TODO(), eventID, eventUpdateReq)

	fmt.Printf("HTTP Response from `UpdateEvent`: %v\n", httpRes)

	if errResp, ok := fingerprint.AsErrorResponse(err); ok {
		switch errResp.Error.Code {
		case fingerprint.ErrorCodeEvent_not_found:
			fmt.Println("event not found")
		case fingerprint.ErrorCodeSecret_api_key_not_found:
			fmt.Println("secret api key not found")
		case fingerprint.ErrorCodeState_not_ready:
			fmt.Println("event is not ready yet for update, please try again later")
		default:
			fmt.Printf("unexpected error: %s\n", errResp.Error.Message)
		}
	} else {
		fmt.Printf("other error occurred: %v\n", err.Error())
	}
}

func ExampleClient_DeleteVisitorData() {
	visitorID := "visitorID_example"

	client := fingerprint.New(fingerprint.WithAPIKey("SECRET_API_KEY"))

	// Delete visitor data. If you are interested in using this API, please contact our support team (https://fingerprint.com/support/) to activate it for you
	httpRes, err := client.DeleteVisitorData(context.TODO(), visitorID)

	fmt.Printf("HTTP Response from `DeleteVisitorData`: %v\n", httpRes)

	if errResp, ok := fingerprint.AsErrorResponse(err); ok {
		switch errResp.Error.Code {
		case fingerprint.ErrorCodeEvent_not_found:
			fmt.Println("event not found")
		case fingerprint.ErrorCodeSecret_api_key_not_found:
			fmt.Println("secret api key not found")
		case fingerprint.ErrorCodeVisitor_not_found:
			fmt.Println("visitor not found")
		default:
			fmt.Printf("unexpected error: %s\n", errResp.Error.Message)
		}
	} else {
		fmt.Printf("other error occurred: %v\n", err.Error())
	}
}
