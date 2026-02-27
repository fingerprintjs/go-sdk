# IPInfoV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Geolocation** | Pointer to [**Geolocation**](Geolocation.md) |  | [optional] 
**Asn** | Pointer to **string** |  | [optional] 
**AsnName** | Pointer to **string** |  | [optional] 
**AsnNetwork** | Pointer to **string** |  | [optional] 
**AsnType** | Pointer to **string** |  | [optional] 
**DatacenterResult** | Pointer to **bool** |  | [optional] 
**DatacenterName** | Pointer to **string** |  | [optional] 

## Methods

### NewIPInfoV4

`func NewIPInfoV4(address string, ) *IPInfoV4`

NewIPInfoV4 instantiates a new IPInfoV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPInfoV4WithDefaults

`func NewIPInfoV4WithDefaults() *IPInfoV4`

NewIPInfoV4WithDefaults instantiates a new IPInfoV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IPInfoV4) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IPInfoV4) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IPInfoV4) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetGeolocation

`func (o *IPInfoV4) GetGeolocation() Geolocation`

GetGeolocation returns the Geolocation field if non-nil, zero value otherwise.

### GetGeolocationOk

`func (o *IPInfoV4) GetGeolocationOk() (*Geolocation, bool)`

GetGeolocationOk returns a tuple with the Geolocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeolocation

`func (o *IPInfoV4) SetGeolocation(v Geolocation)`

SetGeolocation sets Geolocation field to given value.

### HasGeolocation

`func (o *IPInfoV4) HasGeolocation() bool`

HasGeolocation returns a boolean if a field has been set.

### GetAsn

`func (o *IPInfoV4) GetAsn() string`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *IPInfoV4) GetAsnOk() (*string, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *IPInfoV4) SetAsn(v string)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *IPInfoV4) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetAsnName

`func (o *IPInfoV4) GetAsnName() string`

GetAsnName returns the AsnName field if non-nil, zero value otherwise.

### GetAsnNameOk

`func (o *IPInfoV4) GetAsnNameOk() (*string, bool)`

GetAsnNameOk returns a tuple with the AsnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnName

`func (o *IPInfoV4) SetAsnName(v string)`

SetAsnName sets AsnName field to given value.

### HasAsnName

`func (o *IPInfoV4) HasAsnName() bool`

HasAsnName returns a boolean if a field has been set.

### GetAsnNetwork

`func (o *IPInfoV4) GetAsnNetwork() string`

GetAsnNetwork returns the AsnNetwork field if non-nil, zero value otherwise.

### GetAsnNetworkOk

`func (o *IPInfoV4) GetAsnNetworkOk() (*string, bool)`

GetAsnNetworkOk returns a tuple with the AsnNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnNetwork

`func (o *IPInfoV4) SetAsnNetwork(v string)`

SetAsnNetwork sets AsnNetwork field to given value.

### HasAsnNetwork

`func (o *IPInfoV4) HasAsnNetwork() bool`

HasAsnNetwork returns a boolean if a field has been set.

### GetAsnType

`func (o *IPInfoV4) GetAsnType() string`

GetAsnType returns the AsnType field if non-nil, zero value otherwise.

### GetAsnTypeOk

`func (o *IPInfoV4) GetAsnTypeOk() (*string, bool)`

GetAsnTypeOk returns a tuple with the AsnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsnType

`func (o *IPInfoV4) SetAsnType(v string)`

SetAsnType sets AsnType field to given value.

### HasAsnType

`func (o *IPInfoV4) HasAsnType() bool`

HasAsnType returns a boolean if a field has been set.

### GetDatacenterResult

`func (o *IPInfoV4) GetDatacenterResult() bool`

GetDatacenterResult returns the DatacenterResult field if non-nil, zero value otherwise.

### GetDatacenterResultOk

`func (o *IPInfoV4) GetDatacenterResultOk() (*bool, bool)`

GetDatacenterResultOk returns a tuple with the DatacenterResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterResult

`func (o *IPInfoV4) SetDatacenterResult(v bool)`

SetDatacenterResult sets DatacenterResult field to given value.

### HasDatacenterResult

`func (o *IPInfoV4) HasDatacenterResult() bool`

HasDatacenterResult returns a boolean if a field has been set.

### GetDatacenterName

`func (o *IPInfoV4) GetDatacenterName() string`

GetDatacenterName returns the DatacenterName field if non-nil, zero value otherwise.

### GetDatacenterNameOk

`func (o *IPInfoV4) GetDatacenterNameOk() (*string, bool)`

GetDatacenterNameOk returns a tuple with the DatacenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterName

`func (o *IPInfoV4) SetDatacenterName(v string)`

SetDatacenterName sets DatacenterName field to given value.

### HasDatacenterName

`func (o *IPInfoV4) HasDatacenterName() bool`

HasDatacenterName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

