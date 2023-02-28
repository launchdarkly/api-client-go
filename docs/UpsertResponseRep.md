# UpsertResponseRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status of the create or update operation | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 

## Methods

### NewUpsertResponseRep

`func NewUpsertResponseRep() *UpsertResponseRep`

NewUpsertResponseRep instantiates a new UpsertResponseRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertResponseRepWithDefaults

`func NewUpsertResponseRepWithDefaults() *UpsertResponseRep`

NewUpsertResponseRepWithDefaults instantiates a new UpsertResponseRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpsertResponseRep) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpsertResponseRep) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpsertResponseRep) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpsertResponseRep) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *UpsertResponseRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpsertResponseRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpsertResponseRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UpsertResponseRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


