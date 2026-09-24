# DeviceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceManufacturer** | Pointer to **string** | Raw device manufacturer string as reported by the device OS. Not normalized: casing is vendor-defined (samsung, Xiaomi, OPPO, HUAWEI). Always `Apple` on iOS. | [optional] 
**DeviceModel** | Pointer to **string** | Raw device model identifier, as reported by the mobile OS. | [optional] 
**OsVersion** | Pointer to **string** | Mobile operating system version. Component count is not fixed and must not be assumed by consumers: iOS always reports `major.minor.patch` (e.g. `17.4.1`), while Android's precision varies by OS era and which raw signal resolved it — `major` only (`9`, `13`) since Android 10 dropped point releases, `major.minor` (`16.1`) from Android 16 (API 36+) reintroducing a minor component, or a genuine `major.minor.patch` (`8.1.0`) on pre-Android 10 devices that shipped real point releases. Never a fabricated/zero-padded component. | [optional] 


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


