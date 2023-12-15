# ExperimentFlagRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**Variations** | Pointer to [**[]Variation**](Variation.md) |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**Site** | Pointer to [**Link**](Link.md) |  | [optional] 

## Methods

### NewExperimentFlagRep

`func NewExperimentFlagRep() *ExperimentFlagRep`

NewExperimentFlagRep instantiates a new ExperimentFlagRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentFlagRepWithDefaults

`func NewExperimentFlagRepWithDefaults() *ExperimentFlagRep`

NewExperimentFlagRepWithDefaults instantiates a new ExperimentFlagRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExperimentFlagRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExperimentFlagRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExperimentFlagRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExperimentFlagRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *ExperimentFlagRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ExperimentFlagRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ExperimentFlagRep) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ExperimentFlagRep) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetArchived

`func (o *ExperimentFlagRep) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ExperimentFlagRep) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ExperimentFlagRep) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ExperimentFlagRep) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetVariations

`func (o *ExperimentFlagRep) GetVariations() []Variation`

GetVariations returns the Variations field if non-nil, zero value otherwise.

### GetVariationsOk

`func (o *ExperimentFlagRep) GetVariationsOk() (*[]Variation, bool)`

GetVariationsOk returns a tuple with the Variations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariations

`func (o *ExperimentFlagRep) SetVariations(v []Variation)`

SetVariations sets Variations field to given value.

### HasVariations

`func (o *ExperimentFlagRep) HasVariations() bool`

HasVariations returns a boolean if a field has been set.

### GetLinks

`func (o *ExperimentFlagRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExperimentFlagRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExperimentFlagRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExperimentFlagRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetSite

`func (o *ExperimentFlagRep) GetSite() Link`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ExperimentFlagRep) GetSiteOk() (*Link, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ExperimentFlagRep) SetSite(v Link)`

SetSite sets Site field to given value.

### HasSite

`func (o *ExperimentFlagRep) HasSite() bool`

HasSite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


