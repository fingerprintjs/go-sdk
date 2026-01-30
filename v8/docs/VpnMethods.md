# VpnMethods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimezoneMismatch** | Pointer to **bool** | The browser timezone doesn&#39;t match the timezone inferred from the request IP address. | [optional] 
**PublicVpn** | Pointer to **bool** | Request IP address is owned and used by a public VPN service provider. | [optional] 
**AuxiliaryMobile** | Pointer to **bool** | This method applies to mobile devices only. Indicates the result of additional methods used to detect a VPN in mobile devices. | [optional] 
**OsMismatch** | Pointer to **bool** | The browser runs on a different operating system than the operating system inferred from the request network signature. | [optional] 
**Relay** | Pointer to **bool** | Request IP address belongs to a relay service provider, indicating the use of relay services like [Apple Private relay](https://support.apple.com/en-us/102602) or [Cloudflare Warp](https://developers.cloudflare.com/warp-client/).  * Like VPNs, relay services anonymize the visitor&#39;s true IP address. * Unlike traditional VPNs, relay services don&#39;t let visitors spoof their location by choosing an exit node in a different country.  This field allows you to differentiate VPN users and relay service users in your fraud prevention logic.  | [optional] 

## Methods

### NewVpnMethods

`func NewVpnMethods() *VpnMethods`

NewVpnMethods instantiates a new VpnMethods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnMethodsWithDefaults

`func NewVpnMethodsWithDefaults() *VpnMethods`

NewVpnMethodsWithDefaults instantiates a new VpnMethods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezoneMismatch

`func (o *VpnMethods) GetTimezoneMismatch() bool`

GetTimezoneMismatch returns the TimezoneMismatch field if non-nil, zero value otherwise.

### GetTimezoneMismatchOk

`func (o *VpnMethods) GetTimezoneMismatchOk() (*bool, bool)`

GetTimezoneMismatchOk returns a tuple with the TimezoneMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneMismatch

`func (o *VpnMethods) SetTimezoneMismatch(v bool)`

SetTimezoneMismatch sets TimezoneMismatch field to given value.

### HasTimezoneMismatch

`func (o *VpnMethods) HasTimezoneMismatch() bool`

HasTimezoneMismatch returns a boolean if a field has been set.

### GetPublicVpn

`func (o *VpnMethods) GetPublicVpn() bool`

GetPublicVpn returns the PublicVpn field if non-nil, zero value otherwise.

### GetPublicVpnOk

`func (o *VpnMethods) GetPublicVpnOk() (*bool, bool)`

GetPublicVpnOk returns a tuple with the PublicVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicVpn

`func (o *VpnMethods) SetPublicVpn(v bool)`

SetPublicVpn sets PublicVpn field to given value.

### HasPublicVpn

`func (o *VpnMethods) HasPublicVpn() bool`

HasPublicVpn returns a boolean if a field has been set.

### GetAuxiliaryMobile

`func (o *VpnMethods) GetAuxiliaryMobile() bool`

GetAuxiliaryMobile returns the AuxiliaryMobile field if non-nil, zero value otherwise.

### GetAuxiliaryMobileOk

`func (o *VpnMethods) GetAuxiliaryMobileOk() (*bool, bool)`

GetAuxiliaryMobileOk returns a tuple with the AuxiliaryMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxiliaryMobile

`func (o *VpnMethods) SetAuxiliaryMobile(v bool)`

SetAuxiliaryMobile sets AuxiliaryMobile field to given value.

### HasAuxiliaryMobile

`func (o *VpnMethods) HasAuxiliaryMobile() bool`

HasAuxiliaryMobile returns a boolean if a field has been set.

### GetOsMismatch

`func (o *VpnMethods) GetOsMismatch() bool`

GetOsMismatch returns the OsMismatch field if non-nil, zero value otherwise.

### GetOsMismatchOk

`func (o *VpnMethods) GetOsMismatchOk() (*bool, bool)`

GetOsMismatchOk returns a tuple with the OsMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsMismatch

`func (o *VpnMethods) SetOsMismatch(v bool)`

SetOsMismatch sets OsMismatch field to given value.

### HasOsMismatch

`func (o *VpnMethods) HasOsMismatch() bool`

HasOsMismatch returns a boolean if a field has been set.

### GetRelay

`func (o *VpnMethods) GetRelay() bool`

GetRelay returns the Relay field if non-nil, zero value otherwise.

### GetRelayOk

`func (o *VpnMethods) GetRelayOk() (*bool, bool)`

GetRelayOk returns a tuple with the Relay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelay

`func (o *VpnMethods) SetRelay(v bool)`

SetRelay sets Relay field to given value.

### HasRelay

`func (o *VpnMethods) HasRelay() bool`

HasRelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


