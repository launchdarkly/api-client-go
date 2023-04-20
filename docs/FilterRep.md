# FilterRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The filter key | 
**Name** | **string** | The filter name | 
**Description** | **string** | Textual description of the filter | 
**Rules** | **[]map[string]interface{}** |  | 
**Enabled** | **bool** | True if the filter&#39;s rules are active | 
**Archived** | **bool** | True if filter is archived | 
**CreatedAt** | **time.Time** | Creation time | 
**UpdatedAt** | **time.Time** | Last update time | 
**AppliedAt** | **time.Time** | Time the filter was last used to deliver a payload | 
**ArchivedAt** | **time.Time** | Time the filter was archived | 
**Etag** | **string** | An ETag value that may be sent in the If-Match header when updating the filter | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 

## Methods

### NewFilterRep

`func NewFilterRep(key string, name string, description string, rules []map[string]interface{}, enabled bool, archived bool, createdAt time.Time, updatedAt time.Time, appliedAt time.Time, archivedAt time.Time, etag string, links map[string]Link, ) *FilterRep`

NewFilterRep instantiates a new FilterRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterRepWithDefaults

`func NewFilterRepWithDefaults() *FilterRep`

NewFilterRepWithDefaults instantiates a new FilterRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FilterRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FilterRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FilterRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *FilterRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilterRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilterRep) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FilterRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FilterRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FilterRep) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRules

`func (o *FilterRep) GetRules() []map[string]interface{}`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FilterRep) GetRulesOk() (*[]map[string]interface{}, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FilterRep) SetRules(v []map[string]interface{})`

SetRules sets Rules field to given value.


### GetEnabled

`func (o *FilterRep) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FilterRep) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FilterRep) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetArchived

`func (o *FilterRep) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *FilterRep) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *FilterRep) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetCreatedAt

`func (o *FilterRep) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FilterRep) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FilterRep) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *FilterRep) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FilterRep) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FilterRep) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAppliedAt

`func (o *FilterRep) GetAppliedAt() time.Time`

GetAppliedAt returns the AppliedAt field if non-nil, zero value otherwise.

### GetAppliedAtOk

`func (o *FilterRep) GetAppliedAtOk() (*time.Time, bool)`

GetAppliedAtOk returns a tuple with the AppliedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAt

`func (o *FilterRep) SetAppliedAt(v time.Time)`

SetAppliedAt sets AppliedAt field to given value.


### GetArchivedAt

`func (o *FilterRep) GetArchivedAt() time.Time`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *FilterRep) GetArchivedAtOk() (*time.Time, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *FilterRep) SetArchivedAt(v time.Time)`

SetArchivedAt sets ArchivedAt field to given value.


### GetEtag

`func (o *FilterRep) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *FilterRep) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *FilterRep) SetEtag(v string)`

SetEtag sets Etag field to given value.


### GetLinks

`func (o *FilterRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FilterRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FilterRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetAccess

`func (o *FilterRep) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FilterRep) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FilterRep) SetAccess(v Access)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *FilterRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


