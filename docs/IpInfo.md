# IPInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**V4** | Pointer to [**IPInfoV4**](IPInfoV4.md) |  | [optional] 
**V6** | Pointer to [**IPInfoV6**](IPInfoV6.md) |  | [optional] 

## Methods

### NewIPInfo

`func NewIPInfo() *IPInfo`

NewIPInfo instantiates a new IPInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPInfoWithDefaults

`func NewIPInfoWithDefaults() *IPInfo`

NewIPInfoWithDefaults instantiates a new IPInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetV4

`func (o *IPInfo) GetV4() IPInfoV4`

GetV4 returns the V4 field if non-nil, zero value otherwise.

### GetV4Ok

`func (o *IPInfo) GetV4Ok() (*IPInfoV4, bool)`

GetV4Ok returns a tuple with the V4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4

`func (o *IPInfo) SetV4(v IPInfoV4)`

SetV4 sets V4 field to given value.

### HasV4

`func (o *IPInfo) HasV4() bool`

HasV4 returns a boolean if a field has been set.

### GetV6

`func (o *IPInfo) GetV6() IPInfoV6`

GetV6 returns the V6 field if non-nil, zero value otherwise.

### GetV6Ok

`func (o *IPInfo) GetV6Ok() (*IPInfoV6, bool)`

GetV6Ok returns a tuple with the V6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6

`func (o *IPInfo) SetV6(v IPInfoV6)`

SetV6 sets V6 field to given value.

### HasV6

`func (o *IPInfo) HasV6() bool`

HasV6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


