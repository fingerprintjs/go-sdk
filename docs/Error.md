# Error

    ## Properties

    Name | Type | Description | Notes
    ------------ | ------------- | ------------- | -------------
        **Code** | [**ErrorCode**](ErrorCode.md) |  | 
**Message** | **string** |  | 

    ## Methods

        ### NewError

        `func NewError(code ErrorCode, message string, ) *Error`

        NewError instantiates a new Error object
        This constructor will assign default values to properties that have it defined,
        and makes sure properties required by API are set, but the set of arguments
        will change when the set of required properties is changed

        ### NewErrorWithDefaults

        `func NewErrorWithDefaults() *Error`

        NewErrorWithDefaults instantiates a new Error object
        This constructor will only assign default values to properties that have it defined,
        but it doesn't guarantee that properties required by API are set

            ### GetCode

            `func (o *Error) GetCode() ErrorCode`

            GetCode returns the Code field if non-nil, zero value otherwise.

            ### GetCodeOk

            `func (o *Error) GetCodeOk() (*ErrorCode, bool)`

            GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetCode

            `func (o *Error) SetCode(v ErrorCode)`

            SetCode sets Code field to given value.


            ### GetMessage

            `func (o *Error) GetMessage() string`

            GetMessage returns the Message field if non-nil, zero value otherwise.

            ### GetMessageOk

            `func (o *Error) GetMessageOk() (*string, bool)`

            GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
            and a boolean to check if the value has been set.

            ### SetMessage

            `func (o *Error) SetMessage(v string)`

            SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

