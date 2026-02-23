# Identification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VisitorID** | **string** | String of 20 characters that uniquely identifies the visitor&#39;s browser or mobile device. | 
**Confidence** | Pointer to [**IdentificationConfidence**](IdentificationConfidence.md) |  | [optional] 
**VisitorFound** | **bool** | Attribute represents if a visitor had been identified before. | 
**FirstSeenAt** | Pointer to **int64** | Unix epoch time milliseconds timestamp indicating the time at which this visitor ID was first seen. example: &#x60;1758069706642&#x60; - Corresponding to Wed Sep 17 2025 00:41:46 GMT+0000  | [optional] 
**LastSeenAt** | Pointer to **int64** | Unix epoch time milliseconds timestamp indicating the time at which this visitor ID was last seen. example: &#x60;1758069706642&#x60; - Corresponding to Wed Sep 17 2025 00:41:46 GMT+0000  | [optional] 

## Methods

### NewIdentification

`func NewIdentification(visitorID string, visitorFound bool, ) *Identification`

NewIdentification instantiates a new Identification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentificationWithDefaults

`func NewIdentificationWithDefaults() *Identification`

NewIdentificationWithDefaults instantiates a new Identification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisitorID

`func (o *Identification) GetVisitorID() string`

GetVisitorID returns the VisitorID field if non-nil, zero value otherwise.

### GetVisitorIdOk

`func (o *Identification) GetVisitorIdOk() (*string, bool)`

GetVisitorIdOk returns a tuple with the VisitorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitorID

`func (o *Identification) SetVisitorID(v string)`

SetVisitorID sets VisitorID field to given value.


### GetConfidence

`func (o *Identification) GetConfidence() IdentificationConfidence`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *Identification) GetConfidenceOk() (*IdentificationConfidence, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *Identification) SetConfidence(v IdentificationConfidence)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *Identification) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetVisitorFound

`func (o *Identification) GetVisitorFound() bool`

GetVisitorFound returns the VisitorFound field if non-nil, zero value otherwise.

### GetVisitorFoundOk

`func (o *Identification) GetVisitorFoundOk() (*bool, bool)`

GetVisitorFoundOk returns a tuple with the VisitorFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitorFound

`func (o *Identification) SetVisitorFound(v bool)`

SetVisitorFound sets VisitorFound field to given value.


### GetFirstSeenAt

`func (o *Identification) GetFirstSeenAt() int64`

GetFirstSeenAt returns the FirstSeenAt field if non-nil, zero value otherwise.

### GetFirstSeenAtOk

`func (o *Identification) GetFirstSeenAtOk() (*int64, bool)`

GetFirstSeenAtOk returns a tuple with the FirstSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenAt

`func (o *Identification) SetFirstSeenAt(v int64)`

SetFirstSeenAt sets FirstSeenAt field to given value.

### HasFirstSeenAt

`func (o *Identification) HasFirstSeenAt() bool`

HasFirstSeenAt returns a boolean if a field has been set.

### GetLastSeenAt

`func (o *Identification) GetLastSeenAt() int64`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *Identification) GetLastSeenAtOk() (*int64, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *Identification) SetLastSeenAt(v int64)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *Identification) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


