# Destination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **float32** |  | [optional] 
**Config** | Pointer to **interface{}** |  | [optional] 
**On** | Pointer to **bool** |  | [optional] 
**Access** | Pointer to [**AccessRep**](AccessRep.md) |  | [optional] 

## Methods

### NewDestination

`func NewDestination() *Destination`

NewDestination instantiates a new Destination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationWithDefaults

`func NewDestinationWithDefaults() *Destination`

NewDestinationWithDefaults instantiates a new Destination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Destination) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Destination) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Destination) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Destination) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *Destination) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Destination) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Destination) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Destination) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *Destination) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Destination) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Destination) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Destination) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKind

`func (o *Destination) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Destination) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Destination) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Destination) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetVersion

`func (o *Destination) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Destination) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Destination) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Destination) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetConfig

`func (o *Destination) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Destination) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Destination) SetConfig(v interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Destination) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *Destination) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *Destination) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetOn

`func (o *Destination) GetOn() bool`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *Destination) GetOnOk() (*bool, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *Destination) SetOn(v bool)`

SetOn sets On field to given value.

### HasOn

`func (o *Destination) HasOn() bool`

HasOn returns a boolean if a field has been set.

### GetAccess

`func (o *Destination) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *Destination) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *Destination) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *Destination) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


