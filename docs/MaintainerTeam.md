# MaintainerTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key of the maintainer team | 
**Name** | **string** | A human-friendly name for the maintainer team | 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 

## Methods

### NewMaintainerTeam

`func NewMaintainerTeam(key string, name string, ) *MaintainerTeam`

NewMaintainerTeam instantiates a new MaintainerTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintainerTeamWithDefaults

`func NewMaintainerTeamWithDefaults() *MaintainerTeam`

NewMaintainerTeamWithDefaults instantiates a new MaintainerTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MaintainerTeam) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MaintainerTeam) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MaintainerTeam) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *MaintainerTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MaintainerTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MaintainerTeam) SetName(v string)`

SetName sets Name field to given value.


### GetLinks

`func (o *MaintainerTeam) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MaintainerTeam) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MaintainerTeam) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MaintainerTeam) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


