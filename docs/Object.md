# Object

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**PkgPath** | Pointer to **string** |  | [optional] 
**PkgName** | Pointer to **string** |  | [optional] 
**ValidPos** | Pointer to **bool** |  | [optional] 
**Shadow** | Pointer to **map[string]bool** |  | [optional] 

## Methods

### NewObject

`func NewObject() *Object`

NewObject instantiates a new Object object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectWithDefaults

`func NewObjectWithDefaults() *Object`

NewObjectWithDefaults instantiates a new Object object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Object) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Object) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Object) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Object) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKind

`func (o *Object) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Object) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Object) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Object) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPkgPath

`func (o *Object) GetPkgPath() string`

GetPkgPath returns the PkgPath field if non-nil, zero value otherwise.

### GetPkgPathOk

`func (o *Object) GetPkgPathOk() (*string, bool)`

GetPkgPathOk returns a tuple with the PkgPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgPath

`func (o *Object) SetPkgPath(v string)`

SetPkgPath sets PkgPath field to given value.

### HasPkgPath

`func (o *Object) HasPkgPath() bool`

HasPkgPath returns a boolean if a field has been set.

### GetPkgName

`func (o *Object) GetPkgName() string`

GetPkgName returns the PkgName field if non-nil, zero value otherwise.

### GetPkgNameOk

`func (o *Object) GetPkgNameOk() (*string, bool)`

GetPkgNameOk returns a tuple with the PkgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgName

`func (o *Object) SetPkgName(v string)`

SetPkgName sets PkgName field to given value.

### HasPkgName

`func (o *Object) HasPkgName() bool`

HasPkgName returns a boolean if a field has been set.

### GetValidPos

`func (o *Object) GetValidPos() bool`

GetValidPos returns the ValidPos field if non-nil, zero value otherwise.

### GetValidPosOk

`func (o *Object) GetValidPosOk() (*bool, bool)`

GetValidPosOk returns a tuple with the ValidPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidPos

`func (o *Object) SetValidPos(v bool)`

SetValidPos sets ValidPos field to given value.

### HasValidPos

`func (o *Object) HasValidPos() bool`

HasValidPos returns a boolean if a field has been set.

### GetShadow

`func (o *Object) GetShadow() map[string]bool`

GetShadow returns the Shadow field if non-nil, zero value otherwise.

### GetShadowOk

`func (o *Object) GetShadowOk() (*map[string]bool, bool)`

GetShadowOk returns a tuple with the Shadow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadow

`func (o *Object) SetShadow(v map[string]bool)`

SetShadow sets Shadow field to given value.

### HasShadow

`func (o *Object) HasShadow() bool`

HasShadow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


