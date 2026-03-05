# ViewsMemberTeamSummaryRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomRoleKeys** | **[]string** | A list of keys of the custom roles this team has access to | 
**Key** | **string** | The team key | 
**Links** | Pointer to [**map[string]ViewsLink**](ViewsLink.md) |  | [optional] 
**Name** | **string** | The team name | 

## Methods

### NewViewsMemberTeamSummaryRep

`func NewViewsMemberTeamSummaryRep(customRoleKeys []string, key string, name string, ) *ViewsMemberTeamSummaryRep`

NewViewsMemberTeamSummaryRep instantiates a new ViewsMemberTeamSummaryRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewsMemberTeamSummaryRepWithDefaults

`func NewViewsMemberTeamSummaryRepWithDefaults() *ViewsMemberTeamSummaryRep`

NewViewsMemberTeamSummaryRepWithDefaults instantiates a new ViewsMemberTeamSummaryRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomRoleKeys

`func (o *ViewsMemberTeamSummaryRep) GetCustomRoleKeys() []string`

GetCustomRoleKeys returns the CustomRoleKeys field if non-nil, zero value otherwise.

### GetCustomRoleKeysOk

`func (o *ViewsMemberTeamSummaryRep) GetCustomRoleKeysOk() (*[]string, bool)`

GetCustomRoleKeysOk returns a tuple with the CustomRoleKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoleKeys

`func (o *ViewsMemberTeamSummaryRep) SetCustomRoleKeys(v []string)`

SetCustomRoleKeys sets CustomRoleKeys field to given value.


### GetKey

`func (o *ViewsMemberTeamSummaryRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ViewsMemberTeamSummaryRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ViewsMemberTeamSummaryRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetLinks

`func (o *ViewsMemberTeamSummaryRep) GetLinks() map[string]ViewsLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ViewsMemberTeamSummaryRep) GetLinksOk() (*map[string]ViewsLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ViewsMemberTeamSummaryRep) SetLinks(v map[string]ViewsLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ViewsMemberTeamSummaryRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *ViewsMemberTeamSummaryRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewsMemberTeamSummaryRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewsMemberTeamSummaryRep) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


