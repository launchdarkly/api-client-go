# DestinationRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**map[string]InlineResponse200**](InlineResponse200.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **float32** |  | [optional] 
**Config** | Pointer to **interface{}** |  | [optional] 
**On** | Pointer to **bool** |  | [optional] 
**Access** | Pointer to [**AccessRep**](AccessRep.md) |  | [optional] 

## Methods

### NewDestinationRep

`func NewDestinationRep() *DestinationRep`

NewDestinationRep instantiates a new DestinationRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationRepWithDefaults

`func NewDestinationRepWithDefaults() *DestinationRep`

NewDestinationRepWithDefaults instantiates a new DestinationRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DestinationRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DestinationRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DestinationRep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DestinationRep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *DestinationRep) GetLinks() map[string]InlineResponse200`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DestinationRep) GetLinksOk() (*map[string]InlineResponse200, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DestinationRep) SetLinks(v map[string]InlineResponse200)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DestinationRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *DestinationRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DestinationRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKind

`func (o *DestinationRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DestinationRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DestinationRep) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DestinationRep) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetVersion

`func (o *DestinationRep) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DestinationRep) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DestinationRep) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DestinationRep) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetConfig

`func (o *DestinationRep) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DestinationRep) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DestinationRep) SetConfig(v interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DestinationRep) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *DestinationRep) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *DestinationRep) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetOn

`func (o *DestinationRep) GetOn() bool`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *DestinationRep) GetOnOk() (*bool, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *DestinationRep) SetOn(v bool)`

SetOn sets On field to given value.

### HasOn

`func (o *DestinationRep) HasOn() bool`

HasOn returns a boolean if a field has been set.

### GetAccess

`func (o *DestinationRep) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *DestinationRep) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *DestinationRep) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *DestinationRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


