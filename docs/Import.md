# Import

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SegmentKey** | **string** |  | 
**CreationTime** | **int64** |  | 
**Mode** | **string** |  | 
**Status** | **string** |  | 
**Files** | Pointer to [**[]FileRep**](FileRep.md) |  | [optional] 
**Links** | [**map[string]Link**](Link.md) |  | 

## Methods

### NewImport

`func NewImport(id string, segmentKey string, creationTime int64, mode string, status string, links map[string]Link, ) *Import`

NewImport instantiates a new Import object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportWithDefaults

`func NewImportWithDefaults() *Import`

NewImportWithDefaults instantiates a new Import object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Import) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Import) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Import) SetId(v string)`

SetId sets Id field to given value.


### GetSegmentKey

`func (o *Import) GetSegmentKey() string`

GetSegmentKey returns the SegmentKey field if non-nil, zero value otherwise.

### GetSegmentKeyOk

`func (o *Import) GetSegmentKeyOk() (*string, bool)`

GetSegmentKeyOk returns a tuple with the SegmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentKey

`func (o *Import) SetSegmentKey(v string)`

SetSegmentKey sets SegmentKey field to given value.


### GetCreationTime

`func (o *Import) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Import) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Import) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.


### GetMode

`func (o *Import) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Import) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Import) SetMode(v string)`

SetMode sets Mode field to given value.


### GetStatus

`func (o *Import) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Import) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Import) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFiles

`func (o *Import) GetFiles() []FileRep`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Import) GetFilesOk() (*[]FileRep, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Import) SetFiles(v []FileRep)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *Import) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetLinks

`func (o *Import) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Import) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Import) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


