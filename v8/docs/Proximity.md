# Proximity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | A stable privacy-preserving identifier for a given proximity zone.  | 
**PrecisionRadius** | **int32** | The radius of the proximity zone’s precision level, in meters.  | 
**Confidence** | **float32** | A value between `0` and `1` representing the likelihood that the true device location lies within the mapped proximity zone.   * Scores closer to `1` indicate high confidence that the location is inside the mapped proximity zone.   * Scores closer to `0` indicate lower confidence, suggesting the true location may fall in an adjacent zone.  | 

## Methods

### NewProximity

`func NewProximity(iD string, precisionRadius int32, confidence float32, ) *Proximity`

NewProximity instantiates a new Proximity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProximityWithDefaults

`func NewProximityWithDefaults() *Proximity`

NewProximityWithDefaults instantiates a new Proximity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *Proximity) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *Proximity) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *Proximity) SetID(v string)`

SetID sets ID field to given value.


### GetPrecisionRadius

`func (o *Proximity) GetPrecisionRadius() int32`

GetPrecisionRadius returns the PrecisionRadius field if non-nil, zero value otherwise.

### GetPrecisionRadiusOk

`func (o *Proximity) GetPrecisionRadiusOk() (*int32, bool)`

GetPrecisionRadiusOk returns a tuple with the PrecisionRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecisionRadius

`func (o *Proximity) SetPrecisionRadius(v int32)`

SetPrecisionRadius sets PrecisionRadius field to given value.


### GetConfidence

`func (o *Proximity) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *Proximity) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *Proximity) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


