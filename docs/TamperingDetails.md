# TamperingDetails

    ## Properties

    Name | Type | Description | Notes
    ------------ | ------------- | ------------- | -------------
        **AnomalyScore** | Pointer to **float64** | Confidence score (&#x60;0.0 - 1.0&#x60;) for tampering detection:   * Values above &#x60;0.5&#x60; indicate tampering.   * Values below &#x60;0.5&#x60; indicate genuine browsers.  | [optional] 
**AntiDetectBrowser** | Pointer to **bool** | True if the identified browser resembles an \&quot;anti-detect\&quot; browser, such as Incognition, which attempts to evade identification by manipulating its fingerprint.   | [optional] 

    ## Methods

        ### NewTamperingDetails

        `func NewTamperingDetails() *TamperingDetails`

        NewTamperingDetails instantiates a new TamperingDetails object
        This constructor will assign default values to properties that have it defined,
        and makes sure properties required by API are set, but the set of arguments
        will change when the set of required properties is changed

        ### NewTamperingDetailsWithDefaults

        `func NewTamperingDetailsWithDefaults() *TamperingDetails`

        NewTamperingDetailsWithDefaults instantiates a new TamperingDetails object
        This constructor will only assign default values to properties that have it defined,
        but it doesn't guarantee that properties required by API are set

            ### GetAnomalyScore

            `func (o *TamperingDetails) GetAnomalyScore() float64`

            GetAnomalyScore returns the AnomalyScore field if non-nil, zero value otherwise.

            ### GetAnomalyScoreOk

            `func (o *TamperingDetails) GetAnomalyScoreOk() (*float64, bool)`

            GetAnomalyScoreOk returns a tuple with the AnomalyScore field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetAnomalyScore

            `func (o *TamperingDetails) SetAnomalyScore(v float64)`

            SetAnomalyScore sets AnomalyScore field to given value.

                ### HasAnomalyScore

                `func (o *TamperingDetails) HasAnomalyScore() bool`

                HasAnomalyScore returns a boolean if a field has been set.

            ### GetAntiDetectBrowser

            `func (o *TamperingDetails) GetAntiDetectBrowser() bool`

            GetAntiDetectBrowser returns the AntiDetectBrowser field if non-nil, zero value otherwise.

            ### GetAntiDetectBrowserOk

            `func (o *TamperingDetails) GetAntiDetectBrowserOk() (*bool, bool)`

            GetAntiDetectBrowserOk returns a tuple with the AntiDetectBrowser field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetAntiDetectBrowser

            `func (o *TamperingDetails) SetAntiDetectBrowser(v bool)`

            SetAntiDetectBrowser sets AntiDetectBrowser field to given value.

                ### HasAntiDetectBrowser

                `func (o *TamperingDetails) HasAntiDetectBrowser() bool`

                HasAntiDetectBrowser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

