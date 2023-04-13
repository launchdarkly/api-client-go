# CustomRoleSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The custom role&#39;s&#39; ID | 
**Key** | Pointer to **string** | The key of the custom role | [optional] 
**Name** | Pointer to **string** | The name of the custom role | [optional] 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

## Methods

### NewCustomRoleSummary

`func NewCustomRoleSummary(id string, links map[string]Link, ) *CustomRoleSummary`

NewCustomRoleSummary instantiates a new CustomRoleSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomRoleSummaryWithDefaults

`func NewCustomRoleSummaryWithDefaults() *CustomRoleSummary`

NewCustomRoleSummaryWithDefaults instantiates a new CustomRoleSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomRoleSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomRoleSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomRoleSummary) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *CustomRoleSummary) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CustomRoleSummary) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CustomRoleSummary) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CustomRoleSummary) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CustomRoleSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomRoleSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomRoleSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomRoleSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLinks

`func (o *CustomRoleSummary) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomRoleSummary) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomRoleSummary) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


