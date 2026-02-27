# EventSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]Event**](Event.md) |  | 
**PaginationKey** | Pointer to **string** | Use this value in the `pagination_key` parameter to request the next page of search results. | [optional] 
**TotalHits** | Pointer to **int64** | This value represents the total number of events matching the search query, up to the limit provided in the `total_hits` query parameter. Only present if the `total_hits` query parameter was provided. | [optional] 

## Methods

### NewEventSearch

`func NewEventSearch(events []Event, ) *EventSearch`

NewEventSearch instantiates a new EventSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSearchWithDefaults

`func NewEventSearchWithDefaults() *EventSearch`

NewEventSearchWithDefaults instantiates a new EventSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *EventSearch) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventSearch) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventSearch) SetEvents(v []Event)`

SetEvents sets Events field to given value.


### GetPaginationKey

`func (o *EventSearch) GetPaginationKey() string`

GetPaginationKey returns the PaginationKey field if non-nil, zero value otherwise.

### GetPaginationKeyOk

`func (o *EventSearch) GetPaginationKeyOk() (*string, bool)`

GetPaginationKeyOk returns a tuple with the PaginationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationKey

`func (o *EventSearch) SetPaginationKey(v string)`

SetPaginationKey sets PaginationKey field to given value.

### HasPaginationKey

`func (o *EventSearch) HasPaginationKey() bool`

HasPaginationKey returns a boolean if a field has been set.

### GetTotalHits

`func (o *EventSearch) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *EventSearch) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *EventSearch) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *EventSearch) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

