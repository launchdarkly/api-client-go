# ExpandedSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | A unique key used to reference the segment | 
**Name** | **string** | A human-friendly name for the segment | 
**EnvironmentId** | Pointer to **string** | Environment ID of the segment | [optional] 
**EnvironmentKey** | Pointer to **string** | Environment key of the segment | [optional] 
**Description** | Pointer to **string** | Description of the segment | [optional] 
**CreationDate** | Pointer to **int64** | Creation date in milliseconds | [optional] 
**LastModifiedDate** | Pointer to **int64** | Last modification date in milliseconds | [optional] 
**Deleted** | Pointer to **bool** | Whether the segment is deleted | [optional] 
**Tags** | Pointer to **[]string** | Tags for the segment | [optional] 
**Unbounded** | Pointer to **bool** | Whether the segment is unbounded | [optional] 
**Version** | Pointer to **int32** | Version of the segment | [optional] 
**Generation** | Pointer to **int32** | Generation of the segment | [optional] 
**Links** | Pointer to [**ParentAndSelfLinks**](ParentAndSelfLinks.md) |  | [optional] 

## Methods

### NewExpandedSegment

`func NewExpandedSegment(key string, name string, ) *ExpandedSegment`

NewExpandedSegment instantiates a new ExpandedSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandedSegmentWithDefaults

`func NewExpandedSegmentWithDefaults() *ExpandedSegment`

NewExpandedSegmentWithDefaults instantiates a new ExpandedSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ExpandedSegment) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ExpandedSegment) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ExpandedSegment) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ExpandedSegment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpandedSegment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpandedSegment) SetName(v string)`

SetName sets Name field to given value.


### GetEnvironmentId

`func (o *ExpandedSegment) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ExpandedSegment) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ExpandedSegment) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ExpandedSegment) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetEnvironmentKey

`func (o *ExpandedSegment) GetEnvironmentKey() string`

GetEnvironmentKey returns the EnvironmentKey field if non-nil, zero value otherwise.

### GetEnvironmentKeyOk

`func (o *ExpandedSegment) GetEnvironmentKeyOk() (*string, bool)`

GetEnvironmentKeyOk returns a tuple with the EnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentKey

`func (o *ExpandedSegment) SetEnvironmentKey(v string)`

SetEnvironmentKey sets EnvironmentKey field to given value.

### HasEnvironmentKey

`func (o *ExpandedSegment) HasEnvironmentKey() bool`

HasEnvironmentKey returns a boolean if a field has been set.

### GetDescription

`func (o *ExpandedSegment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpandedSegment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpandedSegment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpandedSegment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreationDate

`func (o *ExpandedSegment) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ExpandedSegment) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ExpandedSegment) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ExpandedSegment) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *ExpandedSegment) GetLastModifiedDate() int64`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *ExpandedSegment) GetLastModifiedDateOk() (*int64, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *ExpandedSegment) SetLastModifiedDate(v int64)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *ExpandedSegment) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetDeleted

`func (o *ExpandedSegment) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ExpandedSegment) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ExpandedSegment) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ExpandedSegment) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetTags

`func (o *ExpandedSegment) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExpandedSegment) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExpandedSegment) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExpandedSegment) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUnbounded

`func (o *ExpandedSegment) GetUnbounded() bool`

GetUnbounded returns the Unbounded field if non-nil, zero value otherwise.

### GetUnboundedOk

`func (o *ExpandedSegment) GetUnboundedOk() (*bool, bool)`

GetUnboundedOk returns a tuple with the Unbounded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbounded

`func (o *ExpandedSegment) SetUnbounded(v bool)`

SetUnbounded sets Unbounded field to given value.

### HasUnbounded

`func (o *ExpandedSegment) HasUnbounded() bool`

HasUnbounded returns a boolean if a field has been set.

### GetVersion

`func (o *ExpandedSegment) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExpandedSegment) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExpandedSegment) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExpandedSegment) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetGeneration

`func (o *ExpandedSegment) GetGeneration() int32`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *ExpandedSegment) GetGenerationOk() (*int32, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *ExpandedSegment) SetGeneration(v int32)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *ExpandedSegment) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetLinks

`func (o *ExpandedSegment) GetLinks() ParentAndSelfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExpandedSegment) GetLinksOk() (*ParentAndSelfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExpandedSegment) SetLinks(v ParentAndSelfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExpandedSegment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


