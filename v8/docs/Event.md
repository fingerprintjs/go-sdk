# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | **string** | Unique identifier of the user's request. The first portion of the event_id is a unix epoch milliseconds timestamp For example: `1758130560902.8tRtrH`  | 
**Timestamp** | **int64** | Timestamp of the event with millisecond precision in Unix time. | 
**LinkedID** | Pointer to **string** | A customer-provided id that was sent with the request. | [optional] 
**EnvironmentID** | Pointer to **string** | Environment Id of the event. For example: `ae_47abaca3db2c7c43`  | [optional] 
**Suspect** | Pointer to **bool** | Field is `true` if you have previously set the `suspect` flag for this event using the [Server API Update event endpoint](https://dev.fingerprint.com/reference/updateevent). | [optional] 
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
**FactoryResetTimestamp** | Pointer to **int64** | The time of the most recent factory reset that happened on the **mobile device** is expressed as Unix epoch time. When a factory reset cannot be detected on the mobile device or when the request is initiated from a browser,  this field will correspond to the *epoch* time (i.e 1 Jan 1970 UTC) as a value of 0. See [Factory Reset Detection](https://dev.fingerprint.com/docs/smart-signals-overview#factory-reset-detection) to learn more about this Smart Signal.  | [optional] 
**Frida** | Pointer to **bool** | [Frida](https://frida.re/docs/) detection for Android and iOS devices. There are 2 values: * `true` - Frida detected * `false` - No signs of Frida or the client is not a mobile device.  | [optional] 
**IPBlockList** | Pointer to [**IPBlockList**](IPBlockList.md) |  | [optional] 
**IPInfo** | Pointer to [**IPInfo**](IPInfo.md) |  | [optional] 
**Proxy** | Pointer to **bool** | IP address was used by a public proxy provider or belonged to a known recent residential proxy  | [optional] 
**ProxyConfidence** | Pointer to [**ProxyConfidence**](ProxyConfidence.md) |  | [optional] 
**ProxyDetails** | Pointer to [**ProxyDetails**](ProxyDetails.md) |  | [optional] 
**Incognito** | Pointer to **bool** | `true` if we detected incognito mode used in the browser, `false` otherwise.  | [optional] 
**Jailbroken** | Pointer to **bool** | iOS specific jailbreak detection. There are 2 values:  * `true` - Jailbreak detected. * `false` - No signs of jailbreak or the client is not iOS.  | [optional] 
**LocationSpoofing** | Pointer to **bool** | Flag indicating whether the request came from a mobile device with location spoofing enabled. | [optional] 
**MITMAttack** | Pointer to **bool** | * `true` - When requests made from your users' mobile devices to Fingerprint servers have been intercepted and potentially modified.  * `false` - Otherwise or when the request originated from a browser. See [MitM Attack Detection](https://dev.fingerprint.com/docs/smart-signals-reference#mitm-attack-detection) to learn more about this Smart Signal.  | [optional] 
**PrivacySettings** | Pointer to **bool** | `true` if the request is from a privacy aware browser (e.g. Tor) or from a browser in which fingerprinting is blocked. Otherwise `false`.  | [optional] 
**RootApps** | Pointer to **bool** | Android specific root management apps detection. There are 2 values:  * `true` - Root Management Apps detected (e.g. Magisk). * `false` - No Root Management Apps detected or the client isn't Android.  | [optional] 
**RuleAction** | Pointer to [**EventRuleAction**](EventRuleAction.md) |  | [optional] 
**SuspectScore** | Pointer to **int32** | Suspect Score is an easy way to integrate Smart Signals into your fraud protection work flow.  It is a weighted representation of all Smart Signals present in the payload that helps identify suspicious activity. The value range is [0; S] where S is sum of all Smart Signals weights.  See more details here: https://dev.fingerprint.com/docs/suspect-score  | [optional] 
**Tampering** | Pointer to **bool** | Flag indicating browser tampering was detected. This happens when either:   * There are inconsistencies in the browser configuration that cross internal tampering thresholds (see `tampering_details.anomaly_score`).   * The browser signature resembles an \"anti-detect\" browser specifically designed to evade fingerprinting (see `tampering_details.anti_detect_browser`).  | [optional] 
**TamperingDetails** | Pointer to [**TamperingDetails**](TamperingDetails.md) |  | [optional] 
**Velocity** | Pointer to [**Velocity**](Velocity.md) |  | [optional] 
**VirtualMachine** | Pointer to **bool** | `true` if the request came from a browser running inside a virtual machine (e.g. VMWare), `false` otherwise.  | [optional] 
**VPN** | Pointer to **bool** | VPN or other anonymizing service has been used when sending the request.  | [optional] 
**VPNConfidence** | Pointer to [**VPNConfidence**](VPNConfidence.md) |  | [optional] 
**VPNOriginTimezone** | Pointer to **string** | Local timezone which is used in timezone_mismatch method.  | [optional] 
**VPNOriginCountry** | Pointer to **string** | Country of the request (only for Android SDK version >= 2.4.0, ISO 3166 format or unknown).  | [optional] 
**VPNMethods** | Pointer to [**VPNMethods**](VPNMethods.md) |  | [optional] 
**HighActivityDevice** | Pointer to **bool** | Flag indicating if the request came from a high-activity visitor. | [optional] 
**RawDeviceAttributes** | Pointer to [**RawDeviceAttributes**](RawDeviceAttributes.md) |  | [optional] 

## Methods

### NewEvent

`func NewEvent(eventID string, timestamp int64, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventID

`func (o *Event) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *Event) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *Event) SetEventID(v string)`

SetEventID sets EventID field to given value.


### GetTimestamp

`func (o *Event) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Event) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Event) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetLinkedID

`func (o *Event) GetLinkedID() string`

GetLinkedID returns the LinkedID field if non-nil, zero value otherwise.

### GetLinkedIDOk

`func (o *Event) GetLinkedIDOk() (*string, bool)`

GetLinkedIDOk returns a tuple with the LinkedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedID

`func (o *Event) SetLinkedID(v string)`

SetLinkedID sets LinkedID field to given value.

### HasLinkedID

`func (o *Event) HasLinkedID() bool`

HasLinkedID returns a boolean if a field has been set.

### GetEnvironmentID

`func (o *Event) GetEnvironmentID() string`

GetEnvironmentID returns the EnvironmentID field if non-nil, zero value otherwise.

### GetEnvironmentIDOk

`func (o *Event) GetEnvironmentIDOk() (*string, bool)`

GetEnvironmentIDOk returns a tuple with the EnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentID

`func (o *Event) SetEnvironmentID(v string)`

SetEnvironmentID sets EnvironmentID field to given value.

### HasEnvironmentID

`func (o *Event) HasEnvironmentID() bool`

HasEnvironmentID returns a boolean if a field has been set.

### GetSuspect

`func (o *Event) GetSuspect() bool`

GetSuspect returns the Suspect field if non-nil, zero value otherwise.

### GetSuspectOk

`func (o *Event) GetSuspectOk() (*bool, bool)`

GetSuspectOk returns a tuple with the Suspect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspect

`func (o *Event) SetSuspect(v bool)`

SetSuspect sets Suspect field to given value.

### HasSuspect

`func (o *Event) HasSuspect() bool`

HasSuspect returns a boolean if a field has been set.

### GetSDK

`func (o *Event) GetSDK() SDK`

GetSDK returns the SDK field if non-nil, zero value otherwise.

### GetSDKOk

`func (o *Event) GetSDKOk() (*SDK, bool)`

GetSDKOk returns a tuple with the SDK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDK

`func (o *Event) SetSDK(v SDK)`

SetSDK sets SDK field to given value.

### HasSDK

`func (o *Event) HasSDK() bool`

HasSDK returns a boolean if a field has been set.

### GetReplayed

`func (o *Event) GetReplayed() bool`

GetReplayed returns the Replayed field if non-nil, zero value otherwise.

### GetReplayedOk

`func (o *Event) GetReplayedOk() (*bool, bool)`

GetReplayedOk returns a tuple with the Replayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayed

`func (o *Event) SetReplayed(v bool)`

SetReplayed sets Replayed field to given value.

### HasReplayed

`func (o *Event) HasReplayed() bool`

HasReplayed returns a boolean if a field has been set.

### GetIdentification

`func (o *Event) GetIdentification() Identification`

GetIdentification returns the Identification field if non-nil, zero value otherwise.

### GetIdentificationOk

`func (o *Event) GetIdentificationOk() (*Identification, bool)`

GetIdentificationOk returns a tuple with the Identification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentification

`func (o *Event) SetIdentification(v Identification)`

SetIdentification sets Identification field to given value.

### HasIdentification

`func (o *Event) HasIdentification() bool`

HasIdentification returns a boolean if a field has been set.

### GetSupplementaryIDHighRecall

`func (o *Event) GetSupplementaryIDHighRecall() SupplementaryIDHighRecall`

GetSupplementaryIDHighRecall returns the SupplementaryIDHighRecall field if non-nil, zero value otherwise.

### GetSupplementaryIDHighRecallOk

`func (o *Event) GetSupplementaryIDHighRecallOk() (*SupplementaryIDHighRecall, bool)`

GetSupplementaryIDHighRecallOk returns a tuple with the SupplementaryIDHighRecall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementaryIDHighRecall

`func (o *Event) SetSupplementaryIDHighRecall(v SupplementaryIDHighRecall)`

SetSupplementaryIDHighRecall sets SupplementaryIDHighRecall field to given value.

### HasSupplementaryIDHighRecall

`func (o *Event) HasSupplementaryIDHighRecall() bool`

HasSupplementaryIDHighRecall returns a boolean if a field has been set.

### GetTags

`func (o *Event) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Event) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Event) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Event) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetURL

`func (o *Event) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *Event) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *Event) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *Event) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetBundleID

`func (o *Event) GetBundleID() string`

GetBundleID returns the BundleID field if non-nil, zero value otherwise.

### GetBundleIDOk

`func (o *Event) GetBundleIDOk() (*string, bool)`

GetBundleIDOk returns a tuple with the BundleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleID

`func (o *Event) SetBundleID(v string)`

SetBundleID sets BundleID field to given value.

### HasBundleID

`func (o *Event) HasBundleID() bool`

HasBundleID returns a boolean if a field has been set.

### GetPackageName

`func (o *Event) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *Event) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *Event) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *Event) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetIPAddress

`func (o *Event) GetIPAddress() string`

GetIPAddress returns the IPAddress field if non-nil, zero value otherwise.

### GetIPAddressOk

`func (o *Event) GetIPAddressOk() (*string, bool)`

GetIPAddressOk returns a tuple with the IPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPAddress

`func (o *Event) SetIPAddress(v string)`

SetIPAddress sets IPAddress field to given value.

### HasIPAddress

`func (o *Event) HasIPAddress() bool`

HasIPAddress returns a boolean if a field has been set.

### GetUserAgent

`func (o *Event) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *Event) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *Event) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *Event) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetClientReferrer

`func (o *Event) GetClientReferrer() string`

GetClientReferrer returns the ClientReferrer field if non-nil, zero value otherwise.

### GetClientReferrerOk

`func (o *Event) GetClientReferrerOk() (*string, bool)`

GetClientReferrerOk returns a tuple with the ClientReferrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReferrer

`func (o *Event) SetClientReferrer(v string)`

SetClientReferrer sets ClientReferrer field to given value.

### HasClientReferrer

`func (o *Event) HasClientReferrer() bool`

HasClientReferrer returns a boolean if a field has been set.

### GetBrowserDetails

`func (o *Event) GetBrowserDetails() BrowserDetails`

GetBrowserDetails returns the BrowserDetails field if non-nil, zero value otherwise.

### GetBrowserDetailsOk

`func (o *Event) GetBrowserDetailsOk() (*BrowserDetails, bool)`

GetBrowserDetailsOk returns a tuple with the BrowserDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserDetails

`func (o *Event) SetBrowserDetails(v BrowserDetails)`

SetBrowserDetails sets BrowserDetails field to given value.

### HasBrowserDetails

`func (o *Event) HasBrowserDetails() bool`

HasBrowserDetails returns a boolean if a field has been set.

### GetProximity

`func (o *Event) GetProximity() Proximity`

GetProximity returns the Proximity field if non-nil, zero value otherwise.

### GetProximityOk

`func (o *Event) GetProximityOk() (*Proximity, bool)`

GetProximityOk returns a tuple with the Proximity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProximity

`func (o *Event) SetProximity(v Proximity)`

SetProximity sets Proximity field to given value.

### HasProximity

`func (o *Event) HasProximity() bool`

HasProximity returns a boolean if a field has been set.

### GetBot

`func (o *Event) GetBot() BotResult`

GetBot returns the Bot field if non-nil, zero value otherwise.

### GetBotOk

`func (o *Event) GetBotOk() (*BotResult, bool)`

GetBotOk returns a tuple with the Bot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBot

`func (o *Event) SetBot(v BotResult)`

SetBot sets Bot field to given value.

### HasBot

`func (o *Event) HasBot() bool`

HasBot returns a boolean if a field has been set.

### GetBotType

`func (o *Event) GetBotType() string`

GetBotType returns the BotType field if non-nil, zero value otherwise.

### GetBotTypeOk

`func (o *Event) GetBotTypeOk() (*string, bool)`

GetBotTypeOk returns a tuple with the BotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotType

`func (o *Event) SetBotType(v string)`

SetBotType sets BotType field to given value.

### HasBotType

`func (o *Event) HasBotType() bool`

HasBotType returns a boolean if a field has been set.

### GetBotInfo

`func (o *Event) GetBotInfo() BotInfo`

GetBotInfo returns the BotInfo field if non-nil, zero value otherwise.

### GetBotInfoOk

`func (o *Event) GetBotInfoOk() (*BotInfo, bool)`

GetBotInfoOk returns a tuple with the BotInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotInfo

`func (o *Event) SetBotInfo(v BotInfo)`

SetBotInfo sets BotInfo field to given value.

### HasBotInfo

`func (o *Event) HasBotInfo() bool`

HasBotInfo returns a boolean if a field has been set.

### GetClonedApp

`func (o *Event) GetClonedApp() bool`

GetClonedApp returns the ClonedApp field if non-nil, zero value otherwise.

### GetClonedAppOk

`func (o *Event) GetClonedAppOk() (*bool, bool)`

GetClonedAppOk returns a tuple with the ClonedApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedApp

`func (o *Event) SetClonedApp(v bool)`

SetClonedApp sets ClonedApp field to given value.

### HasClonedApp

`func (o *Event) HasClonedApp() bool`

HasClonedApp returns a boolean if a field has been set.

### GetDeveloperTools

`func (o *Event) GetDeveloperTools() bool`

GetDeveloperTools returns the DeveloperTools field if non-nil, zero value otherwise.

### GetDeveloperToolsOk

`func (o *Event) GetDeveloperToolsOk() (*bool, bool)`

GetDeveloperToolsOk returns a tuple with the DeveloperTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperTools

`func (o *Event) SetDeveloperTools(v bool)`

SetDeveloperTools sets DeveloperTools field to given value.

### HasDeveloperTools

`func (o *Event) HasDeveloperTools() bool`

HasDeveloperTools returns a boolean if a field has been set.

### GetEmulator

`func (o *Event) GetEmulator() bool`

GetEmulator returns the Emulator field if non-nil, zero value otherwise.

### GetEmulatorOk

`func (o *Event) GetEmulatorOk() (*bool, bool)`

GetEmulatorOk returns a tuple with the Emulator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmulator

`func (o *Event) SetEmulator(v bool)`

SetEmulator sets Emulator field to given value.

### HasEmulator

`func (o *Event) HasEmulator() bool`

HasEmulator returns a boolean if a field has been set.

### GetFactoryResetTimestamp

`func (o *Event) GetFactoryResetTimestamp() int64`

GetFactoryResetTimestamp returns the FactoryResetTimestamp field if non-nil, zero value otherwise.

### GetFactoryResetTimestampOk

`func (o *Event) GetFactoryResetTimestampOk() (*int64, bool)`

GetFactoryResetTimestampOk returns a tuple with the FactoryResetTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactoryResetTimestamp

`func (o *Event) SetFactoryResetTimestamp(v int64)`

SetFactoryResetTimestamp sets FactoryResetTimestamp field to given value.

### HasFactoryResetTimestamp

`func (o *Event) HasFactoryResetTimestamp() bool`

HasFactoryResetTimestamp returns a boolean if a field has been set.

### GetFrida

`func (o *Event) GetFrida() bool`

GetFrida returns the Frida field if non-nil, zero value otherwise.

### GetFridaOk

`func (o *Event) GetFridaOk() (*bool, bool)`

GetFridaOk returns a tuple with the Frida field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrida

`func (o *Event) SetFrida(v bool)`

SetFrida sets Frida field to given value.

### HasFrida

`func (o *Event) HasFrida() bool`

HasFrida returns a boolean if a field has been set.

### GetIPBlockList

`func (o *Event) GetIPBlockList() IPBlockList`

GetIPBlockList returns the IPBlockList field if non-nil, zero value otherwise.

### GetIPBlockListOk

`func (o *Event) GetIPBlockListOk() (*IPBlockList, bool)`

GetIPBlockListOk returns a tuple with the IPBlockList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPBlockList

`func (o *Event) SetIPBlockList(v IPBlockList)`

SetIPBlockList sets IPBlockList field to given value.

### HasIPBlockList

`func (o *Event) HasIPBlockList() bool`

HasIPBlockList returns a boolean if a field has been set.

### GetIPInfo

`func (o *Event) GetIPInfo() IPInfo`

GetIPInfo returns the IPInfo field if non-nil, zero value otherwise.

### GetIPInfoOk

`func (o *Event) GetIPInfoOk() (*IPInfo, bool)`

GetIPInfoOk returns a tuple with the IPInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPInfo

`func (o *Event) SetIPInfo(v IPInfo)`

SetIPInfo sets IPInfo field to given value.

### HasIPInfo

`func (o *Event) HasIPInfo() bool`

HasIPInfo returns a boolean if a field has been set.

### GetProxy

`func (o *Event) GetProxy() bool`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *Event) GetProxyOk() (*bool, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *Event) SetProxy(v bool)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *Event) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetProxyConfidence

`func (o *Event) GetProxyConfidence() ProxyConfidence`

GetProxyConfidence returns the ProxyConfidence field if non-nil, zero value otherwise.

### GetProxyConfidenceOk

`func (o *Event) GetProxyConfidenceOk() (*ProxyConfidence, bool)`

GetProxyConfidenceOk returns a tuple with the ProxyConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyConfidence

`func (o *Event) SetProxyConfidence(v ProxyConfidence)`

SetProxyConfidence sets ProxyConfidence field to given value.

### HasProxyConfidence

`func (o *Event) HasProxyConfidence() bool`

HasProxyConfidence returns a boolean if a field has been set.

### GetProxyDetails

`func (o *Event) GetProxyDetails() ProxyDetails`

GetProxyDetails returns the ProxyDetails field if non-nil, zero value otherwise.

### GetProxyDetailsOk

`func (o *Event) GetProxyDetailsOk() (*ProxyDetails, bool)`

GetProxyDetailsOk returns a tuple with the ProxyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyDetails

`func (o *Event) SetProxyDetails(v ProxyDetails)`

SetProxyDetails sets ProxyDetails field to given value.

### HasProxyDetails

`func (o *Event) HasProxyDetails() bool`

HasProxyDetails returns a boolean if a field has been set.

### GetIncognito

`func (o *Event) GetIncognito() bool`

GetIncognito returns the Incognito field if non-nil, zero value otherwise.

### GetIncognitoOk

`func (o *Event) GetIncognitoOk() (*bool, bool)`

GetIncognitoOk returns a tuple with the Incognito field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncognito

`func (o *Event) SetIncognito(v bool)`

SetIncognito sets Incognito field to given value.

### HasIncognito

`func (o *Event) HasIncognito() bool`

HasIncognito returns a boolean if a field has been set.

### GetJailbroken

`func (o *Event) GetJailbroken() bool`

GetJailbroken returns the Jailbroken field if non-nil, zero value otherwise.

### GetJailbrokenOk

`func (o *Event) GetJailbrokenOk() (*bool, bool)`

GetJailbrokenOk returns a tuple with the Jailbroken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJailbroken

`func (o *Event) SetJailbroken(v bool)`

SetJailbroken sets Jailbroken field to given value.

### HasJailbroken

`func (o *Event) HasJailbroken() bool`

HasJailbroken returns a boolean if a field has been set.

### GetLocationSpoofing

`func (o *Event) GetLocationSpoofing() bool`

GetLocationSpoofing returns the LocationSpoofing field if non-nil, zero value otherwise.

### GetLocationSpoofingOk

`func (o *Event) GetLocationSpoofingOk() (*bool, bool)`

GetLocationSpoofingOk returns a tuple with the LocationSpoofing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationSpoofing

`func (o *Event) SetLocationSpoofing(v bool)`

SetLocationSpoofing sets LocationSpoofing field to given value.

### HasLocationSpoofing

`func (o *Event) HasLocationSpoofing() bool`

HasLocationSpoofing returns a boolean if a field has been set.

### GetMITMAttack

`func (o *Event) GetMITMAttack() bool`

GetMITMAttack returns the MITMAttack field if non-nil, zero value otherwise.

### GetMITMAttackOk

`func (o *Event) GetMITMAttackOk() (*bool, bool)`

GetMITMAttackOk returns a tuple with the MITMAttack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMITMAttack

`func (o *Event) SetMITMAttack(v bool)`

SetMITMAttack sets MITMAttack field to given value.

### HasMITMAttack

`func (o *Event) HasMITMAttack() bool`

HasMITMAttack returns a boolean if a field has been set.

### GetPrivacySettings

`func (o *Event) GetPrivacySettings() bool`

GetPrivacySettings returns the PrivacySettings field if non-nil, zero value otherwise.

### GetPrivacySettingsOk

`func (o *Event) GetPrivacySettingsOk() (*bool, bool)`

GetPrivacySettingsOk returns a tuple with the PrivacySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacySettings

`func (o *Event) SetPrivacySettings(v bool)`

SetPrivacySettings sets PrivacySettings field to given value.

### HasPrivacySettings

`func (o *Event) HasPrivacySettings() bool`

HasPrivacySettings returns a boolean if a field has been set.

### GetRootApps

`func (o *Event) GetRootApps() bool`

GetRootApps returns the RootApps field if non-nil, zero value otherwise.

### GetRootAppsOk

`func (o *Event) GetRootAppsOk() (*bool, bool)`

GetRootAppsOk returns a tuple with the RootApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootApps

`func (o *Event) SetRootApps(v bool)`

SetRootApps sets RootApps field to given value.

### HasRootApps

`func (o *Event) HasRootApps() bool`

HasRootApps returns a boolean if a field has been set.

### GetRuleAction

`func (o *Event) GetRuleAction() EventRuleAction`

GetRuleAction returns the RuleAction field if non-nil, zero value otherwise.

### GetRuleActionOk

`func (o *Event) GetRuleActionOk() (*EventRuleAction, bool)`

GetRuleActionOk returns a tuple with the RuleAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleAction

`func (o *Event) SetRuleAction(v EventRuleAction)`

SetRuleAction sets RuleAction field to given value.

### HasRuleAction

`func (o *Event) HasRuleAction() bool`

HasRuleAction returns a boolean if a field has been set.

### GetSuspectScore

`func (o *Event) GetSuspectScore() int32`

GetSuspectScore returns the SuspectScore field if non-nil, zero value otherwise.

### GetSuspectScoreOk

`func (o *Event) GetSuspectScoreOk() (*int32, bool)`

GetSuspectScoreOk returns a tuple with the SuspectScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspectScore

`func (o *Event) SetSuspectScore(v int32)`

SetSuspectScore sets SuspectScore field to given value.

### HasSuspectScore

`func (o *Event) HasSuspectScore() bool`

HasSuspectScore returns a boolean if a field has been set.

### GetTampering

`func (o *Event) GetTampering() bool`

GetTampering returns the Tampering field if non-nil, zero value otherwise.

### GetTamperingOk

`func (o *Event) GetTamperingOk() (*bool, bool)`

GetTamperingOk returns a tuple with the Tampering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTampering

`func (o *Event) SetTampering(v bool)`

SetTampering sets Tampering field to given value.

### HasTampering

`func (o *Event) HasTampering() bool`

HasTampering returns a boolean if a field has been set.

### GetTamperingDetails

`func (o *Event) GetTamperingDetails() TamperingDetails`

GetTamperingDetails returns the TamperingDetails field if non-nil, zero value otherwise.

### GetTamperingDetailsOk

`func (o *Event) GetTamperingDetailsOk() (*TamperingDetails, bool)`

GetTamperingDetailsOk returns a tuple with the TamperingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTamperingDetails

`func (o *Event) SetTamperingDetails(v TamperingDetails)`

SetTamperingDetails sets TamperingDetails field to given value.

### HasTamperingDetails

`func (o *Event) HasTamperingDetails() bool`

HasTamperingDetails returns a boolean if a field has been set.

### GetVelocity

`func (o *Event) GetVelocity() Velocity`

GetVelocity returns the Velocity field if non-nil, zero value otherwise.

### GetVelocityOk

`func (o *Event) GetVelocityOk() (*Velocity, bool)`

GetVelocityOk returns a tuple with the Velocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocity

`func (o *Event) SetVelocity(v Velocity)`

SetVelocity sets Velocity field to given value.

### HasVelocity

`func (o *Event) HasVelocity() bool`

HasVelocity returns a boolean if a field has been set.

### GetVirtualMachine

`func (o *Event) GetVirtualMachine() bool`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *Event) GetVirtualMachineOk() (*bool, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *Event) SetVirtualMachine(v bool)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *Event) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.

### GetVPN

`func (o *Event) GetVPN() bool`

GetVPN returns the VPN field if non-nil, zero value otherwise.

### GetVPNOk

`func (o *Event) GetVPNOk() (*bool, bool)`

GetVPNOk returns a tuple with the VPN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVPN

`func (o *Event) SetVPN(v bool)`

SetVPN sets VPN field to given value.

### HasVPN

`func (o *Event) HasVPN() bool`

HasVPN returns a boolean if a field has been set.

### GetVPNConfidence

`func (o *Event) GetVPNConfidence() VPNConfidence`

GetVPNConfidence returns the VPNConfidence field if non-nil, zero value otherwise.

### GetVPNConfidenceOk

`func (o *Event) GetVPNConfidenceOk() (*VPNConfidence, bool)`

GetVPNConfidenceOk returns a tuple with the VPNConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVPNConfidence

`func (o *Event) SetVPNConfidence(v VPNConfidence)`

SetVPNConfidence sets VPNConfidence field to given value.

### HasVPNConfidence

`func (o *Event) HasVPNConfidence() bool`

HasVPNConfidence returns a boolean if a field has been set.

### GetVPNOriginTimezone

`func (o *Event) GetVPNOriginTimezone() string`

GetVPNOriginTimezone returns the VPNOriginTimezone field if non-nil, zero value otherwise.

### GetVPNOriginTimezoneOk

`func (o *Event) GetVPNOriginTimezoneOk() (*string, bool)`

GetVPNOriginTimezoneOk returns a tuple with the VPNOriginTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVPNOriginTimezone

`func (o *Event) SetVPNOriginTimezone(v string)`

SetVPNOriginTimezone sets VPNOriginTimezone field to given value.

### HasVPNOriginTimezone

`func (o *Event) HasVPNOriginTimezone() bool`

HasVPNOriginTimezone returns a boolean if a field has been set.

### GetVPNOriginCountry

`func (o *Event) GetVPNOriginCountry() string`

GetVPNOriginCountry returns the VPNOriginCountry field if non-nil, zero value otherwise.

### GetVPNOriginCountryOk

`func (o *Event) GetVPNOriginCountryOk() (*string, bool)`

GetVPNOriginCountryOk returns a tuple with the VPNOriginCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVPNOriginCountry

`func (o *Event) SetVPNOriginCountry(v string)`

SetVPNOriginCountry sets VPNOriginCountry field to given value.

### HasVPNOriginCountry

`func (o *Event) HasVPNOriginCountry() bool`

HasVPNOriginCountry returns a boolean if a field has been set.

### GetVPNMethods

`func (o *Event) GetVPNMethods() VPNMethods`

GetVPNMethods returns the VPNMethods field if non-nil, zero value otherwise.

### GetVPNMethodsOk

`func (o *Event) GetVPNMethodsOk() (*VPNMethods, bool)`

GetVPNMethodsOk returns a tuple with the VPNMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVPNMethods

`func (o *Event) SetVPNMethods(v VPNMethods)`

SetVPNMethods sets VPNMethods field to given value.

### HasVPNMethods

`func (o *Event) HasVPNMethods() bool`

HasVPNMethods returns a boolean if a field has been set.

### GetHighActivityDevice

`func (o *Event) GetHighActivityDevice() bool`

GetHighActivityDevice returns the HighActivityDevice field if non-nil, zero value otherwise.

### GetHighActivityDeviceOk

`func (o *Event) GetHighActivityDeviceOk() (*bool, bool)`

GetHighActivityDeviceOk returns a tuple with the HighActivityDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighActivityDevice

`func (o *Event) SetHighActivityDevice(v bool)`

SetHighActivityDevice sets HighActivityDevice field to given value.

### HasHighActivityDevice

`func (o *Event) HasHighActivityDevice() bool`

HasHighActivityDevice returns a boolean if a field has been set.

### GetRawDeviceAttributes

`func (o *Event) GetRawDeviceAttributes() RawDeviceAttributes`

GetRawDeviceAttributes returns the RawDeviceAttributes field if non-nil, zero value otherwise.

### GetRawDeviceAttributesOk

`func (o *Event) GetRawDeviceAttributesOk() (*RawDeviceAttributes, bool)`

GetRawDeviceAttributesOk returns a tuple with the RawDeviceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawDeviceAttributes

`func (o *Event) SetRawDeviceAttributes(v RawDeviceAttributes)`

SetRawDeviceAttributes sets RawDeviceAttributes field to given value.

### HasRawDeviceAttributes

`func (o *Event) HasRawDeviceAttributes() bool`

HasRawDeviceAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


