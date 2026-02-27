# Velocity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinctIP** | Pointer to [**VelocityData**](VelocityData.md) |  | [optional] 
**DistinctLinkedID** | Pointer to [**VelocityData**](VelocityData.md) |  | [optional] 
**DistinctCountry** | Pointer to [**VelocityData**](VelocityData.md) |  | [optional] 
**Events** | Pointer to [**VelocityData**](VelocityData.md) |  | [optional] 
**IPEvents** | Pointer to [**VelocityData**](VelocityData.md) |  | [optional] 
**DistinctIPByLinkedID** | Pointer to [**VelocityData**](VelocityData.md) |  | [optional] 
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

### GetDistinctIP

`func (o *Velocity) GetDistinctIP() VelocityData`

GetDistinctIP returns the DistinctIP field if non-nil, zero value otherwise.

### GetDistinctIPOk

`func (o *Velocity) GetDistinctIPOk() (*VelocityData, bool)`

GetDistinctIPOk returns a tuple with the DistinctIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctIP

`func (o *Velocity) SetDistinctIP(v VelocityData)`

SetDistinctIP sets DistinctIP field to given value.

### HasDistinctIP

`func (o *Velocity) HasDistinctIP() bool`

HasDistinctIP returns a boolean if a field has been set.

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

### GetIPEvents

`func (o *Velocity) GetIPEvents() VelocityData`

GetIPEvents returns the IPEvents field if non-nil, zero value otherwise.

### GetIPEventsOk

`func (o *Velocity) GetIPEventsOk() (*VelocityData, bool)`

GetIPEventsOk returns a tuple with the IPEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPEvents

`func (o *Velocity) SetIPEvents(v VelocityData)`

SetIPEvents sets IPEvents field to given value.

### HasIPEvents

`func (o *Velocity) HasIPEvents() bool`

HasIPEvents returns a boolean if a field has been set.

### GetDistinctIPByLinkedID

`func (o *Velocity) GetDistinctIPByLinkedID() VelocityData`

GetDistinctIPByLinkedID returns the DistinctIPByLinkedID field if non-nil, zero value otherwise.

### GetDistinctIPByLinkedIDOk

`func (o *Velocity) GetDistinctIPByLinkedIDOk() (*VelocityData, bool)`

GetDistinctIPByLinkedIDOk returns a tuple with the DistinctIPByLinkedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctIPByLinkedID

`func (o *Velocity) SetDistinctIPByLinkedID(v VelocityData)`

SetDistinctIPByLinkedID sets DistinctIPByLinkedID field to given value.

### HasDistinctIPByLinkedID

`func (o *Velocity) HasDistinctIPByLinkedID() bool`

HasDistinctIPByLinkedID returns a boolean if a field has been set.

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

