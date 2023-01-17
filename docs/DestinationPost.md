# DestinationPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A human-readable name for your Data Export destination | [optional] 
**Kind** | Pointer to **string** | The type of Data Export destination | [optional] 
**Config** | Pointer to **interface{}** | An object with the configuration parameters required for the destination type | [optional] 
**On** | Pointer to **bool** | Whether the export is on. Displayed as the integration status in the LaunchDarkly UI. | [optional] 

## Methods

### NewDestinationPost

`func NewDestinationPost() *DestinationPost`

NewDestinationPost instantiates a new DestinationPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationPostWithDefaults

`func NewDestinationPostWithDefaults() *DestinationPost`

NewDestinationPostWithDefaults instantiates a new DestinationPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DestinationPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DestinationPost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKind

`func (o *DestinationPost) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DestinationPost) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DestinationPost) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DestinationPost) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetConfig

`func (o *DestinationPost) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DestinationPost) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DestinationPost) SetConfig(v interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DestinationPost) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *DestinationPost) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *DestinationPost) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetOn

`func (o *DestinationPost) GetOn() bool`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *DestinationPost) GetOnOk() (*bool, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *DestinationPost) SetOn(v bool)`

SetOn sets On field to given value.

### HasOn

`func (o *DestinationPost) HasOn() bool`

HasOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


