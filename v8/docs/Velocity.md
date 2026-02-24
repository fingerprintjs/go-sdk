# Velocity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinctIp** | Pointer to [**VelocityData**](VelocityData.md) |  | [optional] 
**DistinctLinkedID** | Pointer to [**VelocityData**](VelocityData.md) |  | [optional] 
**DistinctCountry** | Pointer to [**VelocityData**](VelocityData.md) |  | [optional] 
**Events** | Pointer to [**VelocityData**](VelocityData.md) |  | [optional] 
**IpEvents** | Pointer to [**VelocityData**](VelocityData.md) |  | [optional] 
**DistinctIpByLinkedID** | Pointer to [**VelocityData**](VelocityData.md) |  | [optional] 
**DistinctVisitorIDByLinkedID** | Pointer to [**VelocityData**](VelocityData.md) |  | [optional] 

## Methods

### NewVelocity

`func NewVelocity() *Velocity`

NewVelocity instantiates a new Velocity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVelocityWithDefaults

`func NewVelocityWithDefaults() *Velocity`

NewVelocityWithDefaults instantiates a new Velocity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinctIp

`func (o *Velocity) GetDistinctIp() VelocityData`

GetDistinctIp returns the DistinctIp field if non-nil, zero value otherwise.

### GetDistinctIpOk

`func (o *Velocity) GetDistinctIpOk() (*VelocityData, bool)`

GetDistinctIpOk returns a tuple with the DistinctIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctIp

`func (o *Velocity) SetDistinctIp(v VelocityData)`

SetDistinctIp sets DistinctIp field to given value.

### HasDistinctIp

`func (o *Velocity) HasDistinctIp() bool`

HasDistinctIp returns a boolean if a field has been set.

### GetDistinctLinkedID

`func (o *Velocity) GetDistinctLinkedID() VelocityData`

GetDistinctLinkedID returns the DistinctLinkedID field if non-nil, zero value otherwise.

### GetDistinctLinkedIDOk

`func (o *Velocity) GetDistinctLinkedIDOk() (*VelocityData, bool)`

GetDistinctLinkedIDOk returns a tuple with the DistinctLinkedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctLinkedID

`func (o *Velocity) SetDistinctLinkedID(v VelocityData)`

SetDistinctLinkedID sets DistinctLinkedID field to given value.

### HasDistinctLinkedID

`func (o *Velocity) HasDistinctLinkedID() bool`

HasDistinctLinkedID returns a boolean if a field has been set.

### GetDistinctCountry

`func (o *Velocity) GetDistinctCountry() VelocityData`

GetDistinctCountry returns the DistinctCountry field if non-nil, zero value otherwise.

### GetDistinctCountryOk

`func (o *Velocity) GetDistinctCountryOk() (*VelocityData, bool)`

GetDistinctCountryOk returns a tuple with the DistinctCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctCountry

`func (o *Velocity) SetDistinctCountry(v VelocityData)`

SetDistinctCountry sets DistinctCountry field to given value.

### HasDistinctCountry

`func (o *Velocity) HasDistinctCountry() bool`

HasDistinctCountry returns a boolean if a field has been set.

### GetEvents

`func (o *Velocity) GetEvents() VelocityData`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Velocity) GetEventsOk() (*VelocityData, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Velocity) SetEvents(v VelocityData)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Velocity) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetIpEvents

`func (o *Velocity) GetIpEvents() VelocityData`

GetIpEvents returns the IpEvents field if non-nil, zero value otherwise.

### GetIpEventsOk

`func (o *Velocity) GetIpEventsOk() (*VelocityData, bool)`

GetIpEventsOk returns a tuple with the IpEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpEvents

`func (o *Velocity) SetIpEvents(v VelocityData)`

SetIpEvents sets IpEvents field to given value.

### HasIpEvents

`func (o *Velocity) HasIpEvents() bool`

HasIpEvents returns a boolean if a field has been set.

### GetDistinctIpByLinkedID

`func (o *Velocity) GetDistinctIpByLinkedID() VelocityData`

GetDistinctIpByLinkedID returns the DistinctIpByLinkedID field if non-nil, zero value otherwise.

### GetDistinctIpByLinkedIDOk

`func (o *Velocity) GetDistinctIpByLinkedIDOk() (*VelocityData, bool)`

GetDistinctIpByLinkedIDOk returns a tuple with the DistinctIpByLinkedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctIpByLinkedID

`func (o *Velocity) SetDistinctIpByLinkedID(v VelocityData)`

SetDistinctIpByLinkedID sets DistinctIpByLinkedID field to given value.

### HasDistinctIpByLinkedID

`func (o *Velocity) HasDistinctIpByLinkedID() bool`

HasDistinctIpByLinkedID returns a boolean if a field has been set.

### GetDistinctVisitorIDByLinkedID

`func (o *Velocity) GetDistinctVisitorIDByLinkedID() VelocityData`

GetDistinctVisitorIDByLinkedID returns the DistinctVisitorIDByLinkedID field if non-nil, zero value otherwise.

### GetDistinctVisitorIDByLinkedIDOk

`func (o *Velocity) GetDistinctVisitorIDByLinkedIDOk() (*VelocityData, bool)`

GetDistinctVisitorIDByLinkedIDOk returns a tuple with the DistinctVisitorIDByLinkedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctVisitorIDByLinkedID

`func (o *Velocity) SetDistinctVisitorIDByLinkedID(v VelocityData)`

SetDistinctVisitorIDByLinkedID sets DistinctVisitorIDByLinkedID field to given value.

### HasDistinctVisitorIDByLinkedID

`func (o *Velocity) HasDistinctVisitorIDByLinkedID() bool`

HasDistinctVisitorIDByLinkedID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


