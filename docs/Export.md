# Export

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SegmentKey** | **string** |  | 
**CreationTime** | **int64** |  | 
**Status** | **string** |  | 
**SizeBytes** | **int64** |  | 
**Size** | **string** |  | 
**Initiator** | [**InitiatorRep**](InitiatorRep.md) |  | 
**Links** | [**map[string]Link**](Link.md) |  | 

## Methods

### NewExport

`func NewExport(id string, segmentKey string, creationTime int64, status string, sizeBytes int64, size string, initiator InitiatorRep, links map[string]Link, ) *Export`

NewExport instantiates a new Export object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportWithDefaults

`func NewExportWithDefaults() *Export`

NewExportWithDefaults instantiates a new Export object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Export) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Export) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Export) SetId(v string)`

SetId sets Id field to given value.


### GetSegmentKey

`func (o *Export) GetSegmentKey() string`

GetSegmentKey returns the SegmentKey field if non-nil, zero value otherwise.

### GetSegmentKeyOk

`func (o *Export) GetSegmentKeyOk() (*string, bool)`

GetSegmentKeyOk returns a tuple with the SegmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentKey

`func (o *Export) SetSegmentKey(v string)`

SetSegmentKey sets SegmentKey field to given value.


### GetCreationTime

`func (o *Export) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Export) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Export) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.


### GetStatus

`func (o *Export) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Export) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Export) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSizeBytes

`func (o *Export) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *Export) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *Export) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.


### GetSize

`func (o *Export) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Export) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Export) SetSize(v string)`

SetSize sets Size field to given value.


### GetInitiator

`func (o *Export) GetInitiator() InitiatorRep`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *Export) GetInitiatorOk() (*InitiatorRep, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *Export) SetInitiator(v InitiatorRep)`

SetInitiator sets Initiator field to given value.


### GetLinks

`func (o *Export) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Export) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Export) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


