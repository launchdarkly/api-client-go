# RelayAutoConfigRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Creator** | Pointer to [**MemberSummary**](MemberSummary.md) |  | [optional] 
**Access** | Pointer to [**Access**](Access.md) |  | [optional] 
**Name** | **string** | A human-friendly name for the Relay Proxy configuration | 
**Policy** | [**[]Statement**](Statement.md) | A description of what environments and projects the Relay Proxy should include or exclude | 
**FullKey** | Pointer to **string** | The Relay Proxy configuration key | [optional] 
**DisplayKey** | **string** | The last few characters of the Relay Proxy configuration key, displayed in the LaunchDarkly UI | 
**CreationDate** | **int64** |  | 
**LastModified** | **int64** |  | 

## Methods

### NewRelayAutoConfigRep

`func NewRelayAutoConfigRep(id string, name string, policy []Statement, displayKey string, creationDate int64, lastModified int64, ) *RelayAutoConfigRep`

NewRelayAutoConfigRep instantiates a new RelayAutoConfigRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelayAutoConfigRepWithDefaults

`func NewRelayAutoConfigRepWithDefaults() *RelayAutoConfigRep`

NewRelayAutoConfigRepWithDefaults instantiates a new RelayAutoConfigRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RelayAutoConfigRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelayAutoConfigRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelayAutoConfigRep) SetId(v string)`

SetId sets Id field to given value.


### GetCreator

`func (o *RelayAutoConfigRep) GetCreator() MemberSummary`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RelayAutoConfigRep) GetCreatorOk() (*MemberSummary, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RelayAutoConfigRep) SetCreator(v MemberSummary)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RelayAutoConfigRep) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetAccess

`func (o *RelayAutoConfigRep) GetAccess() Access`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *RelayAutoConfigRep) GetAccessOk() (*Access, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *RelayAutoConfigRep) SetAccess(v Access)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *RelayAutoConfigRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetName

`func (o *RelayAutoConfigRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RelayAutoConfigRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RelayAutoConfigRep) SetName(v string)`

SetName sets Name field to given value.


### GetPolicy

`func (o *RelayAutoConfigRep) GetPolicy() []Statement`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *RelayAutoConfigRep) GetPolicyOk() (*[]Statement, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *RelayAutoConfigRep) SetPolicy(v []Statement)`

SetPolicy sets Policy field to given value.


### GetFullKey

`func (o *RelayAutoConfigRep) GetFullKey() string`

GetFullKey returns the FullKey field if non-nil, zero value otherwise.

### GetFullKeyOk

`func (o *RelayAutoConfigRep) GetFullKeyOk() (*string, bool)`

GetFullKeyOk returns a tuple with the FullKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullKey

`func (o *RelayAutoConfigRep) SetFullKey(v string)`

SetFullKey sets FullKey field to given value.

### HasFullKey

`func (o *RelayAutoConfigRep) HasFullKey() bool`

HasFullKey returns a boolean if a field has been set.

### GetDisplayKey

`func (o *RelayAutoConfigRep) GetDisplayKey() string`

GetDisplayKey returns the DisplayKey field if non-nil, zero value otherwise.

### GetDisplayKeyOk

`func (o *RelayAutoConfigRep) GetDisplayKeyOk() (*string, bool)`

GetDisplayKeyOk returns a tuple with the DisplayKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayKey

`func (o *RelayAutoConfigRep) SetDisplayKey(v string)`

SetDisplayKey sets DisplayKey field to given value.


### GetCreationDate

`func (o *RelayAutoConfigRep) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *RelayAutoConfigRep) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *RelayAutoConfigRep) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetLastModified

`func (o *RelayAutoConfigRep) GetLastModified() int64`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *RelayAutoConfigRep) GetLastModifiedOk() (*int64, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *RelayAutoConfigRep) SetLastModified(v int64)`

SetLastModified sets LastModified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


