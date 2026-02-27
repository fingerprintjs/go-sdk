# Canvas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Winding** | Pointer to **bool** |  | [optional] 
**Geometry** | Pointer to **string** | Hash of geometry rendering output or &#x60;unsupported&#x60; markers. | [optional] 
**Text** | Pointer to **string** | Hash of text rendering output or &#x60;unsupported&#x60; markers. | [optional] 

## Methods

### NewCanvas

`func NewCanvas() *Canvas`

NewCanvas instantiates a new Canvas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCanvasWithDefaults

`func NewCanvasWithDefaults() *Canvas`

NewCanvasWithDefaults instantiates a new Canvas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWinding

`func (o *Canvas) GetWinding() bool`

GetWinding returns the Winding field if non-nil, zero value otherwise.

### GetWindingOk

`func (o *Canvas) GetWindingOk() (*bool, bool)`

GetWindingOk returns a tuple with the Winding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinding

`func (o *Canvas) SetWinding(v bool)`

SetWinding sets Winding field to given value.

### HasWinding

`func (o *Canvas) HasWinding() bool`

HasWinding returns a boolean if a field has been set.

### GetGeometry

`func (o *Canvas) GetGeometry() string`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *Canvas) GetGeometryOk() (*string, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *Canvas) SetGeometry(v string)`

SetGeometry sets Geometry field to given value.

### HasGeometry

`func (o *Canvas) HasGeometry() bool`

HasGeometry returns a boolean if a field has been set.

### GetText

`func (o *Canvas) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Canvas) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Canvas) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Canvas) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


