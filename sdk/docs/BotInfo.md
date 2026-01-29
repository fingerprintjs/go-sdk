# BotInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | The type and purpose of the bot. | 
**Provider** | **string** | The organization or company operating the bot. | 
**ProviderUrl** | Pointer to **string** | The URL of the bot provider&#39;s website. | [optional] 
**Name** | **string** | The specific name or identifier of the bot. | 
**Identity** | **string** | The verification status of the bot&#39;s identity:  * &#x60;verified&#x60; - well-known bot with publicly verifiable identity, directed by the bot provider.  * &#x60;signed&#x60; - bot that signs its platform via Web Bot Auth, directed by the bot provider’s customers.  * &#x60;spoofed&#x60; - bot that claims a public identity but fails verification.  * &#x60;unknown&#x60; - bot that does not publish a verifiable identity.  | 
**Confidence** | **string** | Confidence level of the bot identification. | 

## Methods

### NewBotInfo

`func NewBotInfo(category string, provider string, name string, identity string, confidence string, ) *BotInfo`

NewBotInfo instantiates a new BotInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotInfoWithDefaults

`func NewBotInfoWithDefaults() *BotInfo`

NewBotInfoWithDefaults instantiates a new BotInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *BotInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BotInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BotInfo) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetProvider

`func (o *BotInfo) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BotInfo) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BotInfo) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetProviderUrl

`func (o *BotInfo) GetProviderUrl() string`

GetProviderUrl returns the ProviderUrl field if non-nil, zero value otherwise.

### GetProviderUrlOk

`func (o *BotInfo) GetProviderUrlOk() (*string, bool)`

GetProviderUrlOk returns a tuple with the ProviderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUrl

`func (o *BotInfo) SetProviderUrl(v string)`

SetProviderUrl sets ProviderUrl field to given value.

### HasProviderUrl

`func (o *BotInfo) HasProviderUrl() bool`

HasProviderUrl returns a boolean if a field has been set.

### GetName

`func (o *BotInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BotInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BotInfo) SetName(v string)`

SetName sets Name field to given value.


### GetIdentity

`func (o *BotInfo) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *BotInfo) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *BotInfo) SetIdentity(v string)`

SetIdentity sets Identity field to given value.


### GetConfidence

`func (o *BotInfo) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *BotInfo) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *BotInfo) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


