# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | **string** | Unique identifier of the user's request. The first portion of the event_id is a unix epoch milliseconds timestamp For example: `1758130560902.8tRtrH`  | 
**Timestamp** | **int64** | Timestamp of the event with millisecond precision in Unix time. | 
**IncrementalIdentificationStatus** | Pointer to [**IncrementalIdentificationStatus**](IncrementalIdentificationStatus.md) |  | [optional] 
**LinkedID** | Pointer to **string** | A customer-provided id that was sent with the request. | [optional] 
**EnvironmentID** | Pointer to **string** | Environment Id of the event. For example: `ae_47abaca3db2c7c43`  | [optional] 
**Suspect** | Pointer to **bool** | Field is `true` if you have previously set the `suspect` flag for this event using the [Server API Update event endpoint](https://docs.fingerprint.com/reference/server-api-v4-update-event). | [optional] 
**SDK** | Pointer to [**SDK**](SDK.md) |  | [optional] 
**Replayed** | Pointer to **bool** | `true` if we determined that this payload was replayed, `false` otherwise.  | [optional] 
**Identification** | Pointer to [**Identification**](Identification.md) |  | [optional] 
**SupplementaryIDHighRecall** | Pointer to [**SupplementaryIDHighRecall**](SupplementaryIDHighRecall.md) |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** | A customer-provided value or an object that was sent with the identification request or updated later. | [optional] 
**URL** | Pointer to **string** | Page URL from which the request was sent. For example `https://example.com/`  | [optional] 
**BundleID** | Pointer to **string** | Bundle Id of the iOS application integrated with the Fingerprint SDK for the event. For example: `com.foo.app`  | [optional] 
**PackageName** | Pointer to **string** | Package name of the Android application integrated with the Fingerprint SDK for the event. For example: `com.foo.app`  | [optional] 
**IPAddress** | Pointer to **string** | IP address of the requesting browser or bot. | [optional] 
**UserAgent** | Pointer to **string** | User Agent of the client, for example: `Mozilla/5.0 (Windows NT 6.1; Win64; x64) ....`  | [optional] 
**ClientReferrer** | Pointer to **string** | Client Referrer field corresponds to the `document.referrer` field gathered during an identification request. The value is an empty string if the user navigated to the page directly (not through a link, but, for example, by using a bookmark) For example: `https://example.com/blog/my-article`  | [optional] 
**BrowserDetails** | Pointer to [**BrowserDetails**](BrowserDetails.md) |  | [optional] 
**Proximity** | Pointer to [**Proximity**](Proximity.md) |  | [optional] 
**Bot** | Pointer to [**BotResult**](BotResult.md) |  | [optional] 
**BotType** | Pointer to **string** | Additional classification of the bot type if detected.  | [optional] 
**BotInfo** | Pointer to [**BotInfo**](BotInfo.md) |  | [optional] 
**ClonedApp** | Pointer to **bool** | Android specific cloned application detection. There are 2 values:  * `true` - Presence of app cloners work detected (e.g. fully cloned application found or launch of it inside of a not main working profile detected). * `false` - No signs of cloned application detected or the client is not Android.  | [optional] 
**DeveloperTools** | Pointer to **bool** | `true` if the browser is Chrome with DevTools open or Firefox with Developer Tools open, `false` otherwise.  | [optional] 
**Emulator** | Pointer to **bool** | Android specific emulator detection. There are 2 values:  * `true` - Emulated environment detected (e.g. launch inside of AVD).  * `false` - No signs of emulated environment detected or the client is not Android.  | [optional] 
**FactoryResetTimestamp** | Pointer to **int64** | The time of the most recent factory reset that happened on the **mobile device** is expressed as Unix epoch time. When a factory reset cannot be detected on the mobile device or when the request is initiated from a browser,  this field will correspond to the *epoch* time (i.e 1 Jan 1970 UTC) as a value of 0. See [Factory Reset Detection](https://docs.fingerprint.com/docs/smart-signals-reference#factory-reset-detection) to learn more about this Smart Signal.  | [optional] 
**Frida** | Pointer to **bool** | [Frida](https://frida.re/docs/) detection for Android and iOS devices. There are 2 values: * `true` - Frida detected * `false` - No signs of Frida or the client is not a mobile device.  | [optional] 
**IPBlockList** | Pointer to [**IPBlockList**](IPBlockList.md) |  | [optional] 
**IPInfo** | Pointer to [**IPInfo**](IPInfo.md) |  | [optional] 
**Proxy** | Pointer to **bool** | IP address was used by a public proxy provider or belonged to a known recent residential proxy  | [optional] 
**ProxyConfidence** | Pointer to [**ProxyConfidence**](ProxyConfidence.md) |  | [optional] 
**ProxyDetails** | Pointer to [**ProxyDetails**](ProxyDetails.md) |  | [optional] 
**ProxyMlScore** | Pointer to **float64** | Machine learning–based proxy score, represented as a floating-point value between 0 and 1 (inclusive), with up to three decimal places of precision. A higher score means a higher confidence in the positive `proxy` detection result  | [optional] 
**Incognito** | Pointer to **bool** | `true` if we detected incognito mode used in the browser, `false` otherwise.  | [optional] 
**Jailbroken** | Pointer to **bool** | iOS specific jailbreak detection. There are 2 values:  * `true` - Jailbreak detected. * `false` - No signs of jailbreak or the client is not iOS.  | [optional] 
**LocationSpoofing** | Pointer to **bool** | Flag indicating whether the request came from a mobile device with location spoofing enabled. | [optional] 
**MITMAttack** | Pointer to **bool** | * `true` - When requests made from your users' mobile devices to Fingerprint servers have been intercepted and potentially modified.  * `false` - Otherwise or when the request originated from a browser. See [MitM Attack Detection](https://docs.fingerprint.com/docs/smart-signals-reference#mitm-attack-detection) to learn more about this Smart Signal.  | [optional] 
**PrivacySettings** | Pointer to **bool** | `true` if the request is from a privacy aware browser (e.g. Tor) or from a browser in which fingerprinting is blocked. Otherwise `false`.  | [optional] 
**RootApps** | Pointer to **bool** | Android specific root management apps detection. There are 2 values:  * `true` - Root Management Apps detected (e.g. Magisk). * `false` - No Root Management Apps detected or the client isn't Android.  | [optional] 
**RuleAction** | Pointer to [**EventRuleAction**](EventRuleAction.md) |  | [optional] 
**Simulator** | Pointer to **bool** | iOS specific simulator detection. There are 2 values: * `true` - Simulator environment detected. * `false` - No signs of simulator or the client is not iOS.  | [optional] 
**SuspectScore** | Pointer to **int32** | Suspect Score is an easy way to integrate Smart Signals into your fraud protection work flow.  It is a weighted representation of all Smart Signals present in the payload that helps identify suspicious activity. The value range is [0; S] where S is sum of all Smart Signals weights.  See more details here: https://docs.fingerprint.com/docs/suspect-score  | [optional] 
**Tampering** | Pointer to **bool** | The field can be used as a standalone flag for tampering detection. Alternatively, the more granular fields documented below can be used for workflows that require more context. * `true` if tampering is detected through an anomalous browser signature, anti-detect browser detection, or other tampering-related methods * `false` if none of the tampering checks return a positive result  | [optional] 
**TamperingConfidence** | Pointer to [**TamperingConfidence**](TamperingConfidence.md) |  | [optional] 
**TamperingMlScore** | Pointer to **float64** | The output of this model is captured as tampering_ml_score, a number indicating how likely an event is coming from an anti detect browser. Values close to 1 signify higher confidence and we consider anything above the threshold of 0.8 to be actionable (the result and anti_detect_browser fields conveniently captures that fact)  | [optional] 
**TamperingDetails** | Pointer to [**TamperingDetails**](TamperingDetails.md) |  | [optional] 
**Velocity** | Pointer to [**Velocity**](Velocity.md) |  | [optional] 
**VirtualMachine** | Pointer to **bool** | `true` if the request came from a browser running inside a virtual machine (e.g. VMWare), `false` otherwise.  | [optional] 
**VirtualMachineMlScore** | Pointer to **float64** | Machine learning–based virtual machine score,  represented as a floating-point value between 0 and 1 (inclusive), with up to three decimal places of precision. A higher score means a higher confidence in the positive `virtual_machine` detection result  | [optional] 
**VPN** | Pointer to **bool** | VPN or other anonymizing service has been used when sending the request.  | [optional] 
**VPNConfidence** | Pointer to [**VPNConfidence**](VPNConfidence.md) |  | [optional] 
**VPNOriginTimezone** | Pointer to **string** | Local timezone which is used in timezone_mismatch method.  | [optional] 
**VPNOriginCountry** | Pointer to **string** | Country of the request (only for Android SDK version >= 2.4.0, ISO 3166 format or unknown).  | [optional] 
**VPNMethods** | Pointer to [**VPNMethods**](VPNMethods.md) |  | [optional] 
**HighActivityDevice** | Pointer to **bool** | Flag indicating if the request came from a high-activity visitor. | [optional] 
**RareDevice** | Pointer to **bool** | `true` if the device is considered rare based on its combination of hardware and software attributes.  A device is classified as rare if it falls within the top 99.9 percentile (lowest-frequency segment) of observed traffic,  or if its configuration has not been previously seen (`not_seen`). > This Smart Signal is currently in beta and only available to select customers. If you are interested, please [contact our support team](https://fingerprint.com/support/).  | [optional] 
**RareDevicePercentileBucket** | Pointer to [**RareDevicePercentileBucket**](RareDevicePercentileBucket.md) |  | [optional] 
**RawDeviceAttributes** | Pointer to [**RawDeviceAttributes**](RawDeviceAttributes.md) |  | [optional] 


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


