# ProxyDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProxyType** | **string** | Residential proxies use real user IP addresses to appear as legitimate traffic,  while data center proxies are public proxies hosted in data centers  | 
**LastSeenAt** | Pointer to **int64** | Unix millisecond timestamp with hourly resolution of when this IP was last seen as a proxy  | [optional] 
**Provider** | Pointer to **string** | String representing the last proxy service provider detected when this IP was synced. An IP can be shared by multiple service providers.  | [optional] 

## Methods

### NewProxyDetails

`func NewProxyDetails(proxyType string, ) *ProxyDetails`

NewProxyDetails instantiates a new ProxyDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyDetailsWithDefaults

`func NewProxyDetailsWithDefaults() *ProxyDetails`

NewProxyDetailsWithDefaults instantiates a new ProxyDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProxyType

`func (o *ProxyDetails) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *ProxyDetails) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *ProxyDetails) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.


### GetLastSeenAt

`func (o *ProxyDetails) GetLastSeenAt() int64`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *ProxyDetails) GetLastSeenAtOk() (*int64, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *ProxyDetails) SetLastSeenAt(v int64)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *ProxyDetails) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### GetProvider

`func (o *ProxyDetails) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ProxyDetails) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ProxyDetails) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ProxyDetails) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

