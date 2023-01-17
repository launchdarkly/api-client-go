# ContextKindRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The context kind key | 
**Name** | **string** | The context kind name | 
**Description** | **string** | The context kind description | 
**Version** | **int32** | The context kind version | 
**CreationDate** | **int64** |  | 
**LastModified** | **int64** |  | 
**LastSeen** | Pointer to **int64** |  | [optional] 
**CreatedFrom** | **string** |  | 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 

## Methods

### NewContextKindRep

`func NewContextKindRep(key string, name string, description string, version int32, creationDate int64, lastModified int64, createdFrom string, ) *ContextKindRep`

NewContextKindRep instantiates a new ContextKindRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextKindRepWithDefaults

`func NewContextKindRepWithDefaults() *ContextKindRep`

NewContextKindRepWithDefaults instantiates a new ContextKindRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ContextKindRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ContextKindRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ContextKindRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ContextKindRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContextKindRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContextKindRep) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ContextKindRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContextKindRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContextKindRep) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetVersion

`func (o *ContextKindRep) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ContextKindRep) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ContextKindRep) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCreationDate

`func (o *ContextKindRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ContextKindRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ContextKindRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetLastModified

`func (o *ContextKindRep) GetLastModified() int64`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ContextKindRep) GetLastModifiedOk() (*int64, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ContextKindRep) SetLastModified(v int64)`

SetLastModified sets LastModified field to given value.


### GetLastSeen

`func (o *ContextKindRep) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ContextKindRep) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ContextKindRep) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ContextKindRep) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetCreatedFrom

`func (o *ContextKindRep) GetCreatedFrom() string`

GetCreatedFrom returns the CreatedFrom field if non-nil, zero value otherwise.

### GetCreatedFromOk

`func (o *ContextKindRep) GetCreatedFromOk() (*string, bool)`

GetCreatedFromOk returns a tuple with the CreatedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedFrom

`func (o *ContextKindRep) SetCreatedFrom(v string)`

SetCreatedFrom sets CreatedFrom field to given value.


### GetLinks

`func (o *ContextKindRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContextKindRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContextKindRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ContextKindRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


