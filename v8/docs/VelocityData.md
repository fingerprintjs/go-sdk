# VelocityData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var5Minutes** | **int32** | Count for the last 5 minutes of velocity data, from the time of the event.  | 
**Var1Hour** | **int32** | Count for the last 1 hour of velocity data, from the time of the event.  | 
**Var24Hours** | Pointer to **int32** | The &#x60;24_hours&#x60; interval of &#x60;distinct_ip&#x60;, &#x60;distinct_linked_id&#x60;, &#x60;distinct_country&#x60;, &#x60;distinct_ip_by_linked_id&#x60; and &#x60;distinct_visitor_id_by_linked_id&#x60; will be omitted if the number of &#x60;events&#x60; for the visitor ID in the last 24 hours (&#x60;events.[&#39;24_hours&#39;]&#x60;) is higher than 20.000.  | [optional] 

## Methods

### NewVelocityData

`func NewVelocityData(var5Minutes int32, var1Hour int32, ) *VelocityData`

NewVelocityData instantiates a new VelocityData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVelocityDataWithDefaults

`func NewVelocityDataWithDefaults() *VelocityData`

NewVelocityDataWithDefaults instantiates a new VelocityData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar5Minutes

`func (o *VelocityData) GetVar5Minutes() int32`

GetVar5Minutes returns the Var5Minutes field if non-nil, zero value otherwise.

### GetVar5MinutesOk

`func (o *VelocityData) GetVar5MinutesOk() (*int32, bool)`

GetVar5MinutesOk returns a tuple with the Var5Minutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5Minutes

`func (o *VelocityData) SetVar5Minutes(v int32)`

SetVar5Minutes sets Var5Minutes field to given value.


### GetVar1Hour

`func (o *VelocityData) GetVar1Hour() int32`

GetVar1Hour returns the Var1Hour field if non-nil, zero value otherwise.

### GetVar1HourOk

`func (o *VelocityData) GetVar1HourOk() (*int32, bool)`

GetVar1HourOk returns a tuple with the Var1Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1Hour

`func (o *VelocityData) SetVar1Hour(v int32)`

SetVar1Hour sets Var1Hour field to given value.


### GetVar24Hours

`func (o *VelocityData) GetVar24Hours() int32`

GetVar24Hours returns the Var24Hours field if non-nil, zero value otherwise.

### GetVar24HoursOk

`func (o *VelocityData) GetVar24HoursOk() (*int32, bool)`

GetVar24HoursOk returns a tuple with the Var24Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar24Hours

`func (o *VelocityData) SetVar24Hours(v int32)`

SetVar24Hours sets Var24Hours field to given value.

### HasVar24Hours

`func (o *VelocityData) HasVar24Hours() bool`

HasVar24Hours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


