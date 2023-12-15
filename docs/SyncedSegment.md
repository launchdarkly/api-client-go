# SyncedSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonResponseBody** | Pointer to **string** |  | [optional] 
**RequestParser** | Pointer to [**RequestParser**](RequestParser.md) |  | [optional] 

## Methods

### NewSyncedSegment

`func NewSyncedSegment() *SyncedSegment`

NewSyncedSegment instantiates a new SyncedSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncedSegmentWithDefaults

`func NewSyncedSegmentWithDefaults() *SyncedSegment`

NewSyncedSegmentWithDefaults instantiates a new SyncedSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonResponseBody

`func (o *SyncedSegment) GetJsonResponseBody() string`

GetJsonResponseBody returns the JsonResponseBody field if non-nil, zero value otherwise.

### GetJsonResponseBodyOk

`func (o *SyncedSegment) GetJsonResponseBodyOk() (*string, bool)`

GetJsonResponseBodyOk returns a tuple with the JsonResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonResponseBody

`func (o *SyncedSegment) SetJsonResponseBody(v string)`

SetJsonResponseBody sets JsonResponseBody field to given value.

### HasJsonResponseBody

`func (o *SyncedSegment) HasJsonResponseBody() bool`

HasJsonResponseBody returns a boolean if a field has been set.

### GetRequestParser

`func (o *SyncedSegment) GetRequestParser() RequestParser`

GetRequestParser returns the RequestParser field if non-nil, zero value otherwise.

### GetRequestParserOk

`func (o *SyncedSegment) GetRequestParserOk() (*RequestParser, bool)`

GetRequestParserOk returns a tuple with the RequestParser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParser

`func (o *SyncedSegment) SetRequestParser(v RequestParser)`

SetRequestParser sets RequestParser field to given value.

### HasRequestParser

`func (o *SyncedSegment) HasRequestParser() bool`

HasRequestParser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


