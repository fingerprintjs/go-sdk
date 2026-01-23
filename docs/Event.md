# Event

    ## Properties

    Name | Type | Description | Notes
    ------------ | ------------- | ------------- | -------------
        **EventId** | **string** | Unique identifier of the user&#39;s request. The first portion of the event_id is a unix epoch milliseconds timestamp For example: &#x60;1758130560902.8tRtrH&#x60;  | 
**Timestamp** | **int64** | Timestamp of the event with millisecond precision in Unix time. | 
**LinkedId** | Pointer to **string** | A customer-provided id that was sent with the request. | [optional] 
**EnvironmentId** | Pointer to **string** | Environment Id of the event. For example: &#x60;ae_47abaca3db2c7c43&#x60;  | [optional] 
**Suspect** | Pointer to **bool** | Field is &#x60;true&#x60; if you have previously set the &#x60;suspect&#x60; flag for this event using the [Server API Update event endpoint](https://dev.fingerprint.com/reference/updateevent). | [optional] 
**Sdk** | Pointer to [**SDK**](SDK.md) |  | [optional] 
**Replayed** | Pointer to **bool** | &#x60;true&#x60; if we determined that this payload was replayed, &#x60;false&#x60; otherwise.  | [optional] 
**Identification** | Pointer to [**Identification**](Identification.md) |  | [optional] 
**SupplementaryIdHighRecall** | Pointer to [**SupplementaryIDHighRecall**](SupplementaryIDHighRecall.md) |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** | A customer-provided value or an object that was sent with the identification request or updated later. | [optional] 
**Url** | Pointer to **string** | Page URL from which the request was sent. For example &#x60;https://example.com/&#x60;  | [optional] 
**BundleId** | Pointer to **string** | Bundle Id of the iOS application integrated with the Fingerprint SDK for the event. For example: &#x60;com.foo.app&#x60;  | [optional] 
**PackageName** | Pointer to **string** | Package name of the Android application integrated with the Fingerprint SDK for the event. For example: &#x60;com.foo.app&#x60;  | [optional] 
**IpAddress** | Pointer to **string** | IP address of the requesting browser or bot. | [optional] 
**UserAgent** | Pointer to **string** | User Agent of the client, for example: &#x60;Mozilla/5.0 (Windows NT 6.1; Win64; x64) ....&#x60;  | [optional] 
**ClientReferrer** | Pointer to **string** | Client Referrer field corresponds to the &#x60;document.referrer&#x60; field gathered during an identification request. The value is an empty string if the user navigated to the page directly (not through a link, but, for example, by using a bookmark) For example: &#x60;https://example.com/blog/my-article&#x60;  | [optional] 
**BrowserDetails** | Pointer to [**BrowserDetails**](BrowserDetails.md) |  | [optional] 
**Proximity** | Pointer to [**Proximity**](Proximity.md) |  | [optional] 
**Bot** | Pointer to [**BotResult**](BotResult.md) |  | [optional] 
**BotType** | Pointer to **string** | Additional classification of the bot type if detected.  | [optional] 
**ClonedApp** | Pointer to **bool** | Android specific cloned application detection. There are 2 values:  * &#x60;true&#x60; - Presence of app cloners work detected (e.g. fully cloned application found or launch of it inside of a not main working profile detected). * &#x60;false&#x60; - No signs of cloned application detected or the client is not Android.  | [optional] 
**DeveloperTools** | Pointer to **bool** | &#x60;true&#x60; if the browser is Chrome with DevTools open or Firefox with Developer Tools open, &#x60;false&#x60; otherwise.  | [optional] 
**Emulator** | Pointer to **bool** | Android specific emulator detection. There are 2 values:  * &#x60;true&#x60; - Emulated environment detected (e.g. launch inside of AVD).  * &#x60;false&#x60; - No signs of emulated environment detected or the client is not Android.  | [optional] 
**FactoryResetTimestamp** | Pointer to **int64** | The time of the most recent factory reset that happened on the **mobile device** is expressed as Unix epoch time. When a factory reset cannot be detected on the mobile device or when the request is initiated from a browser,  this field will correspond to the *epoch* time (i.e 1 Jan 1970 UTC) as a value of 0. See [Factory Reset Detection](https://dev.fingerprint.com/docs/smart-signals-overview#factory-reset-detection) to learn more about this Smart Signal.  | [optional] 
**Frida** | Pointer to **bool** | [Frida](https://frida.re/docs/) detection for Android and iOS devices. There are 2 values: * &#x60;true&#x60; - Frida detected * &#x60;false&#x60; - No signs of Frida or the client is not a mobile device.  | [optional] 
**IpBlocklist** | Pointer to [**IPBlockList**](IPBlockList.md) |  | [optional] 
**IpInfo** | Pointer to [**IPInfo**](IPInfo.md) |  | [optional] 
**Proxy** | Pointer to **bool** | IP address was used by a public proxy provider or belonged to a known recent residential proxy  | [optional] 
**ProxyConfidence** | Pointer to [**ProxyConfidence**](ProxyConfidence.md) |  | [optional] 
**ProxyDetails** | Pointer to [**ProxyDetails**](ProxyDetails.md) |  | [optional] 
**Incognito** | Pointer to **bool** | &#x60;true&#x60; if we detected incognito mode used in the browser, &#x60;false&#x60; otherwise.  | [optional] 
**Jailbroken** | Pointer to **bool** | iOS specific jailbreak detection. There are 2 values:  * &#x60;true&#x60; - Jailbreak detected. * &#x60;false&#x60; - No signs of jailbreak or the client is not iOS.  | [optional] 
**LocationSpoofing** | Pointer to **bool** | Flag indicating whether the request came from a mobile device with location spoofing enabled. | [optional] 
**MitmAttack** | Pointer to **bool** | * &#x60;true&#x60; - When requests made from your users&#39; mobile devices to Fingerprint servers have been intercepted and potentially modified.  * &#x60;false&#x60; - Otherwise or when the request originated from a browser. See [MitM Attack Detection](https://dev.fingerprint.com/docs/smart-signals-reference#mitm-attack-detection) to learn more about this Smart Signal.  | [optional] 
**PrivacySettings** | Pointer to **bool** | &#x60;true&#x60; if the request is from a privacy aware browser (e.g. Tor) or from a browser in which fingerprinting is blocked. Otherwise &#x60;false&#x60;.  | [optional] 
**RootApps** | Pointer to **bool** | Android specific root management apps detection. There are 2 values:  * &#x60;true&#x60; - Root Management Apps detected (e.g. Magisk). * &#x60;false&#x60; - No Root Management Apps detected or the client isn&#39;t Android.  | [optional] 
**SuspectScore** | Pointer to **int32** | Suspect Score is an easy way to integrate Smart Signals into your fraud protection work flow.  It is a weighted representation of all Smart Signals present in the payload that helps identify suspicious activity. The value range is [0; S] where S is sum of all Smart Signals weights.  See more details here: https://dev.fingerprint.com/docs/suspect-score  | [optional] 
**Tampering** | Pointer to **bool** | Flag indicating browser tampering was detected. This happens when either:   * There are inconsistencies in the browser configuration that cross internal tampering thresholds (see &#x60;tampering_details.anomaly_score&#x60;).   * The browser signature resembles an \&quot;anti-detect\&quot; browser specifically designed to evade fingerprinting (see &#x60;tampering_details.anti_detect_browser&#x60;).  | [optional] 
**TamperingDetails** | Pointer to [**TamperingDetails**](TamperingDetails.md) |  | [optional] 
**Velocity** | Pointer to [**Velocity**](Velocity.md) |  | [optional] 
**VirtualMachine** | Pointer to **bool** | &#x60;true&#x60; if the request came from a browser running inside a virtual machine (e.g. VMWare), &#x60;false&#x60; otherwise.  | [optional] 
**Vpn** | Pointer to **bool** | VPN or other anonymizing service has been used when sending the request.  | [optional] 
**VpnConfidence** | Pointer to [**VpnConfidence**](VpnConfidence.md) |  | [optional] 
**VpnOriginTimezone** | Pointer to **string** | Local timezone which is used in timezone_mismatch method.  | [optional] 
**VpnOriginCountry** | Pointer to **string** | Country of the request (only for Android SDK version &gt;&#x3D; 2.4.0, ISO 3166 format or unknown).  | [optional] 
**VpnMethods** | Pointer to [**VpnMethods**](VpnMethods.md) |  | [optional] 
**HighActivityDevice** | Pointer to **bool** | Flag indicating if the request came from a high-activity visitor. | [optional] 

    ## Methods

        ### NewEvent

        `func NewEvent(eventId string, timestamp int64, ) *Event`

        NewEvent instantiates a new Event object
        This constructor will assign default values to properties that have it defined,
        and makes sure properties required by API are set, but the set of arguments
        will change when the set of required properties is changed

        ### NewEventWithDefaults

        `func NewEventWithDefaults() *Event`

        NewEventWithDefaults instantiates a new Event object
        This constructor will only assign default values to properties that have it defined,
        but it doesn't guarantee that properties required by API are set

            ### GetEventId

            `func (o *Event) GetEventId() string`

            GetEventId returns the EventId field if non-nil, zero value otherwise.

            ### GetEventIdOk

            `func (o *Event) GetEventIdOk() (*string, bool)`

            GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetEventId

            `func (o *Event) SetEventId(v string)`

            SetEventId sets EventId field to given value.


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


            ### GetLinkedId

            `func (o *Event) GetLinkedId() string`

            GetLinkedId returns the LinkedId field if non-nil, zero value otherwise.

            ### GetLinkedIdOk

            `func (o *Event) GetLinkedIdOk() (*string, bool)`

            GetLinkedIdOk returns a tuple with the LinkedId field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetLinkedId

            `func (o *Event) SetLinkedId(v string)`

            SetLinkedId sets LinkedId field to given value.

                ### HasLinkedId

                `func (o *Event) HasLinkedId() bool`

                HasLinkedId returns a boolean if a field has been set.

            ### GetEnvironmentId

            `func (o *Event) GetEnvironmentId() string`

            GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

            ### GetEnvironmentIdOk

            `func (o *Event) GetEnvironmentIdOk() (*string, bool)`

            GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetEnvironmentId

            `func (o *Event) SetEnvironmentId(v string)`

            SetEnvironmentId sets EnvironmentId field to given value.

                ### HasEnvironmentId

                `func (o *Event) HasEnvironmentId() bool`

                HasEnvironmentId returns a boolean if a field has been set.

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

            ### GetSdk

            `func (o *Event) GetSdk() SDK`

            GetSdk returns the Sdk field if non-nil, zero value otherwise.

            ### GetSdkOk

            `func (o *Event) GetSdkOk() (*SDK, bool)`

            GetSdkOk returns a tuple with the Sdk field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetSdk

            `func (o *Event) SetSdk(v SDK)`

            SetSdk sets Sdk field to given value.

                ### HasSdk

                `func (o *Event) HasSdk() bool`

                HasSdk returns a boolean if a field has been set.

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

            ### GetSupplementaryIdHighRecall

            `func (o *Event) GetSupplementaryIdHighRecall() SupplementaryIDHighRecall`

            GetSupplementaryIdHighRecall returns the SupplementaryIdHighRecall field if non-nil, zero value otherwise.

            ### GetSupplementaryIdHighRecallOk

            `func (o *Event) GetSupplementaryIdHighRecallOk() (*SupplementaryIDHighRecall, bool)`

            GetSupplementaryIdHighRecallOk returns a tuple with the SupplementaryIdHighRecall field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetSupplementaryIdHighRecall

            `func (o *Event) SetSupplementaryIdHighRecall(v SupplementaryIDHighRecall)`

            SetSupplementaryIdHighRecall sets SupplementaryIdHighRecall field to given value.

                ### HasSupplementaryIdHighRecall

                `func (o *Event) HasSupplementaryIdHighRecall() bool`

                HasSupplementaryIdHighRecall returns a boolean if a field has been set.

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

            ### GetUrl

            `func (o *Event) GetUrl() string`

            GetUrl returns the Url field if non-nil, zero value otherwise.

            ### GetUrlOk

            `func (o *Event) GetUrlOk() (*string, bool)`

            GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetUrl

            `func (o *Event) SetUrl(v string)`

            SetUrl sets Url field to given value.

                ### HasUrl

                `func (o *Event) HasUrl() bool`

                HasUrl returns a boolean if a field has been set.

            ### GetBundleId

            `func (o *Event) GetBundleId() string`

            GetBundleId returns the BundleId field if non-nil, zero value otherwise.

            ### GetBundleIdOk

            `func (o *Event) GetBundleIdOk() (*string, bool)`

            GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetBundleId

            `func (o *Event) SetBundleId(v string)`

            SetBundleId sets BundleId field to given value.

                ### HasBundleId

                `func (o *Event) HasBundleId() bool`

                HasBundleId returns a boolean if a field has been set.

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

            ### GetIpAddress

            `func (o *Event) GetIpAddress() string`

            GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

            ### GetIpAddressOk

            `func (o *Event) GetIpAddressOk() (*string, bool)`

            GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetIpAddress

            `func (o *Event) SetIpAddress(v string)`

            SetIpAddress sets IpAddress field to given value.

                ### HasIpAddress

                `func (o *Event) HasIpAddress() bool`

                HasIpAddress returns a boolean if a field has been set.

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

            ### GetIpBlocklist

            `func (o *Event) GetIpBlocklist() IPBlockList`

            GetIpBlocklist returns the IpBlocklist field if non-nil, zero value otherwise.

            ### GetIpBlocklistOk

            `func (o *Event) GetIpBlocklistOk() (*IPBlockList, bool)`

            GetIpBlocklistOk returns a tuple with the IpBlocklist field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetIpBlocklist

            `func (o *Event) SetIpBlocklist(v IPBlockList)`

            SetIpBlocklist sets IpBlocklist field to given value.

                ### HasIpBlocklist

                `func (o *Event) HasIpBlocklist() bool`

                HasIpBlocklist returns a boolean if a field has been set.

            ### GetIpInfo

            `func (o *Event) GetIpInfo() IPInfo`

            GetIpInfo returns the IpInfo field if non-nil, zero value otherwise.

            ### GetIpInfoOk

            `func (o *Event) GetIpInfoOk() (*IPInfo, bool)`

            GetIpInfoOk returns a tuple with the IpInfo field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetIpInfo

            `func (o *Event) SetIpInfo(v IPInfo)`

            SetIpInfo sets IpInfo field to given value.

                ### HasIpInfo

                `func (o *Event) HasIpInfo() bool`

                HasIpInfo returns a boolean if a field has been set.

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

            ### GetMitmAttack

            `func (o *Event) GetMitmAttack() bool`

            GetMitmAttack returns the MitmAttack field if non-nil, zero value otherwise.

            ### GetMitmAttackOk

            `func (o *Event) GetMitmAttackOk() (*bool, bool)`

            GetMitmAttackOk returns a tuple with the MitmAttack field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetMitmAttack

            `func (o *Event) SetMitmAttack(v bool)`

            SetMitmAttack sets MitmAttack field to given value.

                ### HasMitmAttack

                `func (o *Event) HasMitmAttack() bool`

                HasMitmAttack returns a boolean if a field has been set.

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

            ### GetVpn

            `func (o *Event) GetVpn() bool`

            GetVpn returns the Vpn field if non-nil, zero value otherwise.

            ### GetVpnOk

            `func (o *Event) GetVpnOk() (*bool, bool)`

            GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetVpn

            `func (o *Event) SetVpn(v bool)`

            SetVpn sets Vpn field to given value.

                ### HasVpn

                `func (o *Event) HasVpn() bool`

                HasVpn returns a boolean if a field has been set.

            ### GetVpnConfidence

            `func (o *Event) GetVpnConfidence() VpnConfidence`

            GetVpnConfidence returns the VpnConfidence field if non-nil, zero value otherwise.

            ### GetVpnConfidenceOk

            `func (o *Event) GetVpnConfidenceOk() (*VpnConfidence, bool)`

            GetVpnConfidenceOk returns a tuple with the VpnConfidence field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetVpnConfidence

            `func (o *Event) SetVpnConfidence(v VpnConfidence)`

            SetVpnConfidence sets VpnConfidence field to given value.

                ### HasVpnConfidence

                `func (o *Event) HasVpnConfidence() bool`

                HasVpnConfidence returns a boolean if a field has been set.

            ### GetVpnOriginTimezone

            `func (o *Event) GetVpnOriginTimezone() string`

            GetVpnOriginTimezone returns the VpnOriginTimezone field if non-nil, zero value otherwise.

            ### GetVpnOriginTimezoneOk

            `func (o *Event) GetVpnOriginTimezoneOk() (*string, bool)`

            GetVpnOriginTimezoneOk returns a tuple with the VpnOriginTimezone field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetVpnOriginTimezone

            `func (o *Event) SetVpnOriginTimezone(v string)`

            SetVpnOriginTimezone sets VpnOriginTimezone field to given value.

                ### HasVpnOriginTimezone

                `func (o *Event) HasVpnOriginTimezone() bool`

                HasVpnOriginTimezone returns a boolean if a field has been set.

            ### GetVpnOriginCountry

            `func (o *Event) GetVpnOriginCountry() string`

            GetVpnOriginCountry returns the VpnOriginCountry field if non-nil, zero value otherwise.

            ### GetVpnOriginCountryOk

            `func (o *Event) GetVpnOriginCountryOk() (*string, bool)`

            GetVpnOriginCountryOk returns a tuple with the VpnOriginCountry field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetVpnOriginCountry

            `func (o *Event) SetVpnOriginCountry(v string)`

            SetVpnOriginCountry sets VpnOriginCountry field to given value.

                ### HasVpnOriginCountry

                `func (o *Event) HasVpnOriginCountry() bool`

                HasVpnOriginCountry returns a boolean if a field has been set.

            ### GetVpnMethods

            `func (o *Event) GetVpnMethods() VpnMethods`

            GetVpnMethods returns the VpnMethods field if non-nil, zero value otherwise.

            ### GetVpnMethodsOk

            `func (o *Event) GetVpnMethodsOk() (*VpnMethods, bool)`

            GetVpnMethodsOk returns a tuple with the VpnMethods field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetVpnMethods

            `func (o *Event) SetVpnMethods(v VpnMethods)`

            SetVpnMethods sets VpnMethods field to given value.

                ### HasVpnMethods

                `func (o *Event) HasVpnMethods() bool`

                HasVpnMethods returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

