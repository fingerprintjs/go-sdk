# RawDeviceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FontPreferences** | Pointer to [**FontPreferences**](FontPreferences.md) |  | [optional] 
**Emoji** | Pointer to [**Emoji**](Emoji.md) |  | [optional] 
**Fonts** | Pointer to **[]string** | List of fonts detected on the device. | [optional] 
**DeviceMemory** | Pointer to **int32** | Rounded amount of RAM (in gigabytes) reported by the browser. | [optional] 
**Timezone** | Pointer to **string** | Timezone identifier detected on the client. | [optional] 
**Canvas** | Pointer to [**Canvas**](Canvas.md) |  | [optional] 
**Languages** | Pointer to **[][]string** | Navigator languages reported by the agent including fallbacks. Each inner array represents ordered language preferences reported by different APIs.  | [optional] 
**WebglExtensions** | Pointer to [**WebGlExtensions**](WebGlExtensions.md) |  | [optional] 
**WebglBasics** | Pointer to [**WebGlBasics**](WebGlBasics.md) |  | [optional] 
**ScreenResolution** | Pointer to **[]int32** | Current screen resolution. | [optional] 
**TouchSupport** | Pointer to [**TouchSupport**](TouchSupport.md) |  | [optional] 
**Oscpu** | Pointer to **string** | Navigator `oscpu` string. | [optional] 
**Architecture** | Pointer to **int32** | Integer representing the CPU architecture exposed by the browser. | [optional] 
**CookiesEnabled** | Pointer to **bool** | Whether the cookies are enabled in the browser. | [optional] 
**HardwareConcurrency** | Pointer to **int32** | Number of logical CPU cores reported by the browser. | [optional] 
**DateTimeLocale** | Pointer to **string** | Locale derived from the Intl.DateTimeFormat API. Negative values indicate known error states. The negative statuses can be: - \"-1\": A permanent status for browsers that don't support Intl API. - \"-2\": A permanent status for browsers that don't supportDateTimeFormat constructor. - \"-3\": A permanent status for browsers in which DateTimeFormat locale is undefined or null.  | [optional] 
**Vendor** | Pointer to **string** | Navigator vendor string. | [optional] 
**ColorDepth** | Pointer to **int32** | Screen color depth in bits. | [optional] 
**Platform** | Pointer to **string** | Navigator platform string. | [optional] 
**SessionStorage** | Pointer to **bool** | Whether sessionStorage is available. | [optional] 
**LocalStorage** | Pointer to **bool** | Whether localStorage is available. | [optional] 
**Audio** | Pointer to **float64** | AudioContext fingerprint or negative status when unavailable. The negative statuses can be: - -1: A permanent status for those browsers which are known to always suspend audio context - -2: A permanent status for browsers that don't support the signal - -3: A temporary status that means that an unexpected timeout has happened  | [optional] 
**Plugins** | Pointer to [**[]PluginsInner**](PluginsInner.md) | Browser plugins reported by `navigator.plugins`. | [optional] 
**IndexedDb** | Pointer to **bool** | Whether IndexedDB is available. | [optional] 
**Math** | Pointer to **string** | Hash of Math APIs used for entropy collection. | [optional] 

## Methods

### NewRawDeviceAttributes

`func NewRawDeviceAttributes() *RawDeviceAttributes`

NewRawDeviceAttributes instantiates a new RawDeviceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawDeviceAttributesWithDefaults

`func NewRawDeviceAttributesWithDefaults() *RawDeviceAttributes`

NewRawDeviceAttributesWithDefaults instantiates a new RawDeviceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFontPreferences

`func (o *RawDeviceAttributes) GetFontPreferences() FontPreferences`

GetFontPreferences returns the FontPreferences field if non-nil, zero value otherwise.

### GetFontPreferencesOk

`func (o *RawDeviceAttributes) GetFontPreferencesOk() (*FontPreferences, bool)`

GetFontPreferencesOk returns a tuple with the FontPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontPreferences

`func (o *RawDeviceAttributes) SetFontPreferences(v FontPreferences)`

SetFontPreferences sets FontPreferences field to given value.

### HasFontPreferences

`func (o *RawDeviceAttributes) HasFontPreferences() bool`

HasFontPreferences returns a boolean if a field has been set.

### GetEmoji

`func (o *RawDeviceAttributes) GetEmoji() Emoji`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *RawDeviceAttributes) GetEmojiOk() (*Emoji, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *RawDeviceAttributes) SetEmoji(v Emoji)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *RawDeviceAttributes) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetFonts

`func (o *RawDeviceAttributes) GetFonts() []string`

GetFonts returns the Fonts field if non-nil, zero value otherwise.

### GetFontsOk

`func (o *RawDeviceAttributes) GetFontsOk() (*[]string, bool)`

GetFontsOk returns a tuple with the Fonts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFonts

`func (o *RawDeviceAttributes) SetFonts(v []string)`

SetFonts sets Fonts field to given value.

### HasFonts

`func (o *RawDeviceAttributes) HasFonts() bool`

HasFonts returns a boolean if a field has been set.

### GetDeviceMemory

`func (o *RawDeviceAttributes) GetDeviceMemory() int32`

GetDeviceMemory returns the DeviceMemory field if non-nil, zero value otherwise.

### GetDeviceMemoryOk

`func (o *RawDeviceAttributes) GetDeviceMemoryOk() (*int32, bool)`

GetDeviceMemoryOk returns a tuple with the DeviceMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMemory

`func (o *RawDeviceAttributes) SetDeviceMemory(v int32)`

SetDeviceMemory sets DeviceMemory field to given value.

### HasDeviceMemory

`func (o *RawDeviceAttributes) HasDeviceMemory() bool`

HasDeviceMemory returns a boolean if a field has been set.

### GetTimezone

`func (o *RawDeviceAttributes) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *RawDeviceAttributes) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *RawDeviceAttributes) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *RawDeviceAttributes) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetCanvas

`func (o *RawDeviceAttributes) GetCanvas() Canvas`

GetCanvas returns the Canvas field if non-nil, zero value otherwise.

### GetCanvasOk

`func (o *RawDeviceAttributes) GetCanvasOk() (*Canvas, bool)`

GetCanvasOk returns a tuple with the Canvas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanvas

`func (o *RawDeviceAttributes) SetCanvas(v Canvas)`

SetCanvas sets Canvas field to given value.

### HasCanvas

`func (o *RawDeviceAttributes) HasCanvas() bool`

HasCanvas returns a boolean if a field has been set.

### GetLanguages

`func (o *RawDeviceAttributes) GetLanguages() [][]string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *RawDeviceAttributes) GetLanguagesOk() (*[][]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *RawDeviceAttributes) SetLanguages(v [][]string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *RawDeviceAttributes) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetWebglExtensions

`func (o *RawDeviceAttributes) GetWebglExtensions() WebGlExtensions`

GetWebglExtensions returns the WebglExtensions field if non-nil, zero value otherwise.

### GetWebglExtensionsOk

`func (o *RawDeviceAttributes) GetWebglExtensionsOk() (*WebGlExtensions, bool)`

GetWebglExtensionsOk returns a tuple with the WebglExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebglExtensions

`func (o *RawDeviceAttributes) SetWebglExtensions(v WebGlExtensions)`

SetWebglExtensions sets WebglExtensions field to given value.

### HasWebglExtensions

`func (o *RawDeviceAttributes) HasWebglExtensions() bool`

HasWebglExtensions returns a boolean if a field has been set.

### GetWebglBasics

`func (o *RawDeviceAttributes) GetWebglBasics() WebGlBasics`

GetWebglBasics returns the WebglBasics field if non-nil, zero value otherwise.

### GetWebglBasicsOk

`func (o *RawDeviceAttributes) GetWebglBasicsOk() (*WebGlBasics, bool)`

GetWebglBasicsOk returns a tuple with the WebglBasics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebglBasics

`func (o *RawDeviceAttributes) SetWebglBasics(v WebGlBasics)`

SetWebglBasics sets WebglBasics field to given value.

### HasWebglBasics

`func (o *RawDeviceAttributes) HasWebglBasics() bool`

HasWebglBasics returns a boolean if a field has been set.

### GetScreenResolution

`func (o *RawDeviceAttributes) GetScreenResolution() []int32`

GetScreenResolution returns the ScreenResolution field if non-nil, zero value otherwise.

### GetScreenResolutionOk

`func (o *RawDeviceAttributes) GetScreenResolutionOk() (*[]int32, bool)`

GetScreenResolutionOk returns a tuple with the ScreenResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenResolution

`func (o *RawDeviceAttributes) SetScreenResolution(v []int32)`

SetScreenResolution sets ScreenResolution field to given value.

### HasScreenResolution

`func (o *RawDeviceAttributes) HasScreenResolution() bool`

HasScreenResolution returns a boolean if a field has been set.

### GetTouchSupport

`func (o *RawDeviceAttributes) GetTouchSupport() TouchSupport`

GetTouchSupport returns the TouchSupport field if non-nil, zero value otherwise.

### GetTouchSupportOk

`func (o *RawDeviceAttributes) GetTouchSupportOk() (*TouchSupport, bool)`

GetTouchSupportOk returns a tuple with the TouchSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTouchSupport

`func (o *RawDeviceAttributes) SetTouchSupport(v TouchSupport)`

SetTouchSupport sets TouchSupport field to given value.

### HasTouchSupport

`func (o *RawDeviceAttributes) HasTouchSupport() bool`

HasTouchSupport returns a boolean if a field has been set.

### GetOscpu

`func (o *RawDeviceAttributes) GetOscpu() string`

GetOscpu returns the Oscpu field if non-nil, zero value otherwise.

### GetOscpuOk

`func (o *RawDeviceAttributes) GetOscpuOk() (*string, bool)`

GetOscpuOk returns a tuple with the Oscpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOscpu

`func (o *RawDeviceAttributes) SetOscpu(v string)`

SetOscpu sets Oscpu field to given value.

### HasOscpu

`func (o *RawDeviceAttributes) HasOscpu() bool`

HasOscpu returns a boolean if a field has been set.

### GetArchitecture

`func (o *RawDeviceAttributes) GetArchitecture() int32`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *RawDeviceAttributes) GetArchitectureOk() (*int32, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *RawDeviceAttributes) SetArchitecture(v int32)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *RawDeviceAttributes) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetCookiesEnabled

`func (o *RawDeviceAttributes) GetCookiesEnabled() bool`

GetCookiesEnabled returns the CookiesEnabled field if non-nil, zero value otherwise.

### GetCookiesEnabledOk

`func (o *RawDeviceAttributes) GetCookiesEnabledOk() (*bool, bool)`

GetCookiesEnabledOk returns a tuple with the CookiesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookiesEnabled

`func (o *RawDeviceAttributes) SetCookiesEnabled(v bool)`

SetCookiesEnabled sets CookiesEnabled field to given value.

### HasCookiesEnabled

`func (o *RawDeviceAttributes) HasCookiesEnabled() bool`

HasCookiesEnabled returns a boolean if a field has been set.

### GetHardwareConcurrency

`func (o *RawDeviceAttributes) GetHardwareConcurrency() int32`

GetHardwareConcurrency returns the HardwareConcurrency field if non-nil, zero value otherwise.

### GetHardwareConcurrencyOk

`func (o *RawDeviceAttributes) GetHardwareConcurrencyOk() (*int32, bool)`

GetHardwareConcurrencyOk returns a tuple with the HardwareConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareConcurrency

`func (o *RawDeviceAttributes) SetHardwareConcurrency(v int32)`

SetHardwareConcurrency sets HardwareConcurrency field to given value.

### HasHardwareConcurrency

`func (o *RawDeviceAttributes) HasHardwareConcurrency() bool`

HasHardwareConcurrency returns a boolean if a field has been set.

### GetDateTimeLocale

`func (o *RawDeviceAttributes) GetDateTimeLocale() string`

GetDateTimeLocale returns the DateTimeLocale field if non-nil, zero value otherwise.

### GetDateTimeLocaleOk

`func (o *RawDeviceAttributes) GetDateTimeLocaleOk() (*string, bool)`

GetDateTimeLocaleOk returns a tuple with the DateTimeLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeLocale

`func (o *RawDeviceAttributes) SetDateTimeLocale(v string)`

SetDateTimeLocale sets DateTimeLocale field to given value.

### HasDateTimeLocale

`func (o *RawDeviceAttributes) HasDateTimeLocale() bool`

HasDateTimeLocale returns a boolean if a field has been set.

### GetVendor

`func (o *RawDeviceAttributes) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *RawDeviceAttributes) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *RawDeviceAttributes) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *RawDeviceAttributes) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetColorDepth

`func (o *RawDeviceAttributes) GetColorDepth() int32`

GetColorDepth returns the ColorDepth field if non-nil, zero value otherwise.

### GetColorDepthOk

`func (o *RawDeviceAttributes) GetColorDepthOk() (*int32, bool)`

GetColorDepthOk returns a tuple with the ColorDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorDepth

`func (o *RawDeviceAttributes) SetColorDepth(v int32)`

SetColorDepth sets ColorDepth field to given value.

### HasColorDepth

`func (o *RawDeviceAttributes) HasColorDepth() bool`

HasColorDepth returns a boolean if a field has been set.

### GetPlatform

`func (o *RawDeviceAttributes) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *RawDeviceAttributes) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *RawDeviceAttributes) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *RawDeviceAttributes) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSessionStorage

`func (o *RawDeviceAttributes) GetSessionStorage() bool`

GetSessionStorage returns the SessionStorage field if non-nil, zero value otherwise.

### GetSessionStorageOk

`func (o *RawDeviceAttributes) GetSessionStorageOk() (*bool, bool)`

GetSessionStorageOk returns a tuple with the SessionStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStorage

`func (o *RawDeviceAttributes) SetSessionStorage(v bool)`

SetSessionStorage sets SessionStorage field to given value.

### HasSessionStorage

`func (o *RawDeviceAttributes) HasSessionStorage() bool`

HasSessionStorage returns a boolean if a field has been set.

### GetLocalStorage

`func (o *RawDeviceAttributes) GetLocalStorage() bool`

GetLocalStorage returns the LocalStorage field if non-nil, zero value otherwise.

### GetLocalStorageOk

`func (o *RawDeviceAttributes) GetLocalStorageOk() (*bool, bool)`

GetLocalStorageOk returns a tuple with the LocalStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStorage

`func (o *RawDeviceAttributes) SetLocalStorage(v bool)`

SetLocalStorage sets LocalStorage field to given value.

### HasLocalStorage

`func (o *RawDeviceAttributes) HasLocalStorage() bool`

HasLocalStorage returns a boolean if a field has been set.

### GetAudio

`func (o *RawDeviceAttributes) GetAudio() float64`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *RawDeviceAttributes) GetAudioOk() (*float64, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *RawDeviceAttributes) SetAudio(v float64)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *RawDeviceAttributes) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### GetPlugins

`func (o *RawDeviceAttributes) GetPlugins() []PluginsInner`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *RawDeviceAttributes) GetPluginsOk() (*[]PluginsInner, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *RawDeviceAttributes) SetPlugins(v []PluginsInner)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *RawDeviceAttributes) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetIndexedDb

`func (o *RawDeviceAttributes) GetIndexedDb() bool`

GetIndexedDb returns the IndexedDb field if non-nil, zero value otherwise.

### GetIndexedDbOk

`func (o *RawDeviceAttributes) GetIndexedDbOk() (*bool, bool)`

GetIndexedDbOk returns a tuple with the IndexedDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedDb

`func (o *RawDeviceAttributes) SetIndexedDb(v bool)`

SetIndexedDb sets IndexedDb field to given value.

### HasIndexedDb

`func (o *RawDeviceAttributes) HasIndexedDb() bool`

HasIndexedDb returns a boolean if a field has been set.

### GetMath

`func (o *RawDeviceAttributes) GetMath() string`

GetMath returns the Math field if non-nil, zero value otherwise.

### GetMathOk

`func (o *RawDeviceAttributes) GetMathOk() (*string, bool)`

GetMathOk returns a tuple with the Math field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMath

`func (o *RawDeviceAttributes) SetMath(v string)`

SetMath sets Math field to given value.

### HasMath

`func (o *RawDeviceAttributes) HasMath() bool`

HasMath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

