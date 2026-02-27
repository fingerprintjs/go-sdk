# IPBlockList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailSpam** | Pointer to **bool** | IP address was part of a known email spam attack (SMTP). | [optional] 
**AttackSource** | Pointer to **bool** | IP address was part of a known network attack (SSH/HTTPS). | [optional] 
**TorNode** | Pointer to **bool** | IP address was part of known TOR network activity. | [optional] 

## Methods

### NewIPBlockList

`func NewIPBlockList() *IPBlockList`

NewIPBlockList instantiates a new IPBlockList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPBlockListWithDefaults

`func NewIPBlockListWithDefaults() *IPBlockList`

NewIPBlockListWithDefaults instantiates a new IPBlockList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailSpam

`func (o *IPBlockList) GetEmailSpam() bool`

GetEmailSpam returns the EmailSpam field if non-nil, zero value otherwise.

### GetEmailSpamOk

`func (o *IPBlockList) GetEmailSpamOk() (*bool, bool)`

GetEmailSpamOk returns a tuple with the EmailSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSpam

`func (o *IPBlockList) SetEmailSpam(v bool)`

SetEmailSpam sets EmailSpam field to given value.

### HasEmailSpam

`func (o *IPBlockList) HasEmailSpam() bool`

HasEmailSpam returns a boolean if a field has been set.

### GetAttackSource

`func (o *IPBlockList) GetAttackSource() bool`

GetAttackSource returns the AttackSource field if non-nil, zero value otherwise.

### GetAttackSourceOk

`func (o *IPBlockList) GetAttackSourceOk() (*bool, bool)`

GetAttackSourceOk returns a tuple with the AttackSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackSource

`func (o *IPBlockList) SetAttackSource(v bool)`

SetAttackSource sets AttackSource field to given value.

### HasAttackSource

`func (o *IPBlockList) HasAttackSource() bool`

HasAttackSource returns a boolean if a field has been set.

### GetTorNode

`func (o *IPBlockList) GetTorNode() bool`

GetTorNode returns the TorNode field if non-nil, zero value otherwise.

### GetTorNodeOk

`func (o *IPBlockList) GetTorNodeOk() (*bool, bool)`

GetTorNodeOk returns a tuple with the TorNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorNode

`func (o *IPBlockList) SetTorNode(v bool)`

SetTorNode sets TorNode field to given value.

### HasTorNode

`func (o *IPBlockList) HasTorNode() bool`

HasTorNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

