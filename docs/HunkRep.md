# HunkRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartingLineNumber** | **int32** | Line number of beginning of code reference hunk | 
**Lines** | Pointer to **string** | Contextual lines of code that include the referenced feature flag | [optional] 
**ProjKey** | Pointer to **string** | The project key | [optional] 
**FlagKey** | Pointer to **string** | The feature flag key | [optional] 
**Aliases** | Pointer to **[]string** | An array of flag key aliases | [optional] 

## Methods

### NewHunkRep

`func NewHunkRep(startingLineNumber int32, ) *HunkRep`

NewHunkRep instantiates a new HunkRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHunkRepWithDefaults

`func NewHunkRepWithDefaults() *HunkRep`

NewHunkRepWithDefaults instantiates a new HunkRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartingLineNumber

`func (o *HunkRep) GetStartingLineNumber() int32`

GetStartingLineNumber returns the StartingLineNumber field if non-nil, zero value otherwise.

### GetStartingLineNumberOk

`func (o *HunkRep) GetStartingLineNumberOk() (*int32, bool)`

GetStartingLineNumberOk returns a tuple with the StartingLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingLineNumber

`func (o *HunkRep) SetStartingLineNumber(v int32)`

SetStartingLineNumber sets StartingLineNumber field to given value.


### GetLines

`func (o *HunkRep) GetLines() string`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *HunkRep) GetLinesOk() (*string, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *HunkRep) SetLines(v string)`

SetLines sets Lines field to given value.

### HasLines

`func (o *HunkRep) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetProjKey

`func (o *HunkRep) GetProjKey() string`

GetProjKey returns the ProjKey field if non-nil, zero value otherwise.

### GetProjKeyOk

`func (o *HunkRep) GetProjKeyOk() (*string, bool)`

GetProjKeyOk returns a tuple with the ProjKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjKey

`func (o *HunkRep) SetProjKey(v string)`

SetProjKey sets ProjKey field to given value.

### HasProjKey

`func (o *HunkRep) HasProjKey() bool`

HasProjKey returns a boolean if a field has been set.

### GetFlagKey

`func (o *HunkRep) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *HunkRep) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *HunkRep) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.

### HasFlagKey

`func (o *HunkRep) HasFlagKey() bool`

HasFlagKey returns a boolean if a field has been set.

### GetAliases

`func (o *HunkRep) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *HunkRep) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *HunkRep) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *HunkRep) HasAliases() bool`

HasAliases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


