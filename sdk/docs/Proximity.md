# Proximity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A stable privacy-preserving identifier for a given proximity zone.  | 
**PrecisionRadius** | **int32** | The radius of the proximity zone’s precision level, in meters.  | 
**Confidence** | **float32** | A value between &#x60;0&#x60; and &#x60;1&#x60; representing the likelihood that the true device location lies within the mapped proximity zone.   * Scores closer to &#x60;1&#x60; indicate high confidence that the location is inside the mapped proximity zone.   * Scores closer to &#x60;0&#x60; indicate lower confidence, suggesting the true location may fall in an adjacent zone.  | 

## Methods

### NewProximity

`func NewProximity(id string, precisionRadius int32, confidence float32, ) *Proximity`

NewProximity instantiates a new Proximity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProximityWithDefaults

`func NewProximityWithDefaults() *Proximity`

NewProximityWithDefaults instantiates a new Proximity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Proximity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Proximity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Proximity) SetId(v string)`

SetId sets Id field to given value.


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


