# SupplementaryIDHighRecall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VisitorID** | **string** | String of 20 characters that uniquely identifies the visitor's browser or mobile device. | 
**VisitorFound** | **bool** | Attribute represents if a visitor had been identified before. | 
**Confidence** | Pointer to [**IdentificationConfidence**](IdentificationConfidence.md) |  | [optional] 
**FirstSeenAt** | Pointer to **int64** | Unix epoch time milliseconds timestamp indicating the time at which this ID was first seen. example: `1758069706642` - Corresponding to Wed Sep 17 2025 00:41:46 GMT+0000  | [optional] 
**LastSeenAt** | Pointer to **int64** | Unix epoch time milliseconds timestamp indicating the time at which this ID was last seen. example: `1758069706642` - Corresponding to Wed Sep 17 2025 00:41:46 GMT+0000  | [optional] 

## Methods

### NewSupplementaryIDHighRecall

`func NewSupplementaryIDHighRecall(visitorID string, visitorFound bool, ) *SupplementaryIDHighRecall`

NewSupplementaryIDHighRecall instantiates a new SupplementaryIDHighRecall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplementaryIDHighRecallWithDefaults

`func NewSupplementaryIDHighRecallWithDefaults() *SupplementaryIDHighRecall`

NewSupplementaryIDHighRecallWithDefaults instantiates a new SupplementaryIDHighRecall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisitorID

`func (o *SupplementaryIDHighRecall) GetVisitorID() string`

GetVisitorID returns the VisitorID field if non-nil, zero value otherwise.

### GetVisitorIDOk

`func (o *SupplementaryIDHighRecall) GetVisitorIDOk() (*string, bool)`

GetVisitorIDOk returns a tuple with the VisitorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitorID

`func (o *SupplementaryIDHighRecall) SetVisitorID(v string)`

SetVisitorID sets VisitorID field to given value.


### GetVisitorFound

`func (o *SupplementaryIDHighRecall) GetVisitorFound() bool`

GetVisitorFound returns the VisitorFound field if non-nil, zero value otherwise.

### GetVisitorFoundOk

`func (o *SupplementaryIDHighRecall) GetVisitorFoundOk() (*bool, bool)`

GetVisitorFoundOk returns a tuple with the VisitorFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitorFound

`func (o *SupplementaryIDHighRecall) SetVisitorFound(v bool)`

SetVisitorFound sets VisitorFound field to given value.


### GetConfidence

`func (o *SupplementaryIDHighRecall) GetConfidence() IdentificationConfidence`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *SupplementaryIDHighRecall) GetConfidenceOk() (*IdentificationConfidence, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *SupplementaryIDHighRecall) SetConfidence(v IdentificationConfidence)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *SupplementaryIDHighRecall) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetFirstSeenAt

`func (o *SupplementaryIDHighRecall) GetFirstSeenAt() int64`

GetFirstSeenAt returns the FirstSeenAt field if non-nil, zero value otherwise.

### GetFirstSeenAtOk

`func (o *SupplementaryIDHighRecall) GetFirstSeenAtOk() (*int64, bool)`

GetFirstSeenAtOk returns a tuple with the FirstSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenAt

`func (o *SupplementaryIDHighRecall) SetFirstSeenAt(v int64)`

SetFirstSeenAt sets FirstSeenAt field to given value.

### HasFirstSeenAt

`func (o *SupplementaryIDHighRecall) HasFirstSeenAt() bool`

HasFirstSeenAt returns a boolean if a field has been set.

### GetLastSeenAt

`func (o *SupplementaryIDHighRecall) GetLastSeenAt() int64`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *SupplementaryIDHighRecall) GetLastSeenAtOk() (*int64, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *SupplementaryIDHighRecall) SetLastSeenAt(v int64)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *SupplementaryIDHighRecall) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


