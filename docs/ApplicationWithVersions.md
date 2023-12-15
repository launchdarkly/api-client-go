# ApplicationWithVersions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Name** | **string** |  | 
**Kind** | **string** |  | 
**Versions** | Pointer to [**[]Version**](Version.md) |  | [optional] 

## Methods

### NewApplicationWithVersions

`func NewApplicationWithVersions(key string, name string, kind string, ) *ApplicationWithVersions`

NewApplicationWithVersions instantiates a new ApplicationWithVersions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithVersionsWithDefaults

`func NewApplicationWithVersionsWithDefaults() *ApplicationWithVersions`

NewApplicationWithVersionsWithDefaults instantiates a new ApplicationWithVersions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ApplicationWithVersions) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApplicationWithVersions) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApplicationWithVersions) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ApplicationWithVersions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationWithVersions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationWithVersions) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *ApplicationWithVersions) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ApplicationWithVersions) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ApplicationWithVersions) SetKind(v string)`

SetKind sets Kind field to given value.


### GetVersions

`func (o *ApplicationWithVersions) GetVersions() []Version`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ApplicationWithVersions) GetVersionsOk() (*[]Version, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ApplicationWithVersions) SetVersions(v []Version)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ApplicationWithVersions) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


