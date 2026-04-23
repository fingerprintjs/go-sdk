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
**Languages** | Pointer to **[][]string** | Navigator languages reported by the agent including fallbacks. Each inner array represents ordered language preferences reported by different APIs. Available for both browsers and iOS devices  | [optional] 
**WebglExtensions** | Pointer to [**WebGlExtensions**](WebGlExtensions.md) |  | [optional] 
**WebglBasics** | Pointer to [**WebGlBasics**](WebGlBasics.md) |  | [optional] 
**ScreenResolution** | Pointer to **[]int32** | Current screen resolution. Available for both browsers and iOS devices | [optional] 
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
**DeviceModel** | Pointer to **string** | Device model string. Available only for Android and iOS devices. | [optional] 
**DeviceManufacturer** | Pointer to **string** | Device manufacturer string. Available only for Android and iOS devices. | [optional] 
**FontHash** | Pointer to **string** | Unique identifier for the user’s installed fonts. | [optional] 
**TimezoneOffset** | Pointer to **string** | UTC offset in \"±HH:MM\" format derived from the detected IANA timezone. | [optional] 


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


