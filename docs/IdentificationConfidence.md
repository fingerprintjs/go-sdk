# IdentificationConfidence

    ## Properties

    Name | Type | Description | Notes
    ------------ | ------------- | ------------- | -------------
        **Score** | **float64** | The confidence score is a floating-point number between 0 and 1 that represents the probability of accurate identification. | 
**Version** | Pointer to **string** | The version name of the method used to calculate the Confidence score. This field is only present for customers who opted in to an alternative calculation method. | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

    ## Methods

        ### NewIdentificationConfidence

        `func NewIdentificationConfidence(score float64, ) *IdentificationConfidence`

        NewIdentificationConfidence instantiates a new IdentificationConfidence object
        This constructor will assign default values to properties that have it defined,
        and makes sure properties required by API are set, but the set of arguments
        will change when the set of required properties is changed

        ### NewIdentificationConfidenceWithDefaults

        `func NewIdentificationConfidenceWithDefaults() *IdentificationConfidence`

        NewIdentificationConfidenceWithDefaults instantiates a new IdentificationConfidence object
        This constructor will only assign default values to properties that have it defined,
        but it doesn't guarantee that properties required by API are set

            ### GetScore

            `func (o *IdentificationConfidence) GetScore() float64`

            GetScore returns the Score field if non-nil, zero value otherwise.

            ### GetScoreOk

            `func (o *IdentificationConfidence) GetScoreOk() (*float64, bool)`

            GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetScore

            `func (o *IdentificationConfidence) SetScore(v float64)`

            SetScore sets Score field to given value.


            ### GetVersion

            `func (o *IdentificationConfidence) GetVersion() string`

            GetVersion returns the Version field if non-nil, zero value otherwise.

            ### GetVersionOk

            `func (o *IdentificationConfidence) GetVersionOk() (*string, bool)`

            GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetVersion

            `func (o *IdentificationConfidence) SetVersion(v string)`

            SetVersion sets Version field to given value.

                ### HasVersion

                `func (o *IdentificationConfidence) HasVersion() bool`

                HasVersion returns a boolean if a field has been set.

            ### GetComment

            `func (o *IdentificationConfidence) GetComment() string`

            GetComment returns the Comment field if non-nil, zero value otherwise.

            ### GetCommentOk

            `func (o *IdentificationConfidence) GetCommentOk() (*string, bool)`

            GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetComment

            `func (o *IdentificationConfidence) SetComment(v string)`

            SetComment sets Comment field to given value.

                ### HasComment

                `func (o *IdentificationConfidence) HasComment() bool`

                HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

