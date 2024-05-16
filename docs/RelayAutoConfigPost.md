# RelayAutoConfigPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-friendly name for the Relay Proxy configuration | 
**Policy** | [**[]Statement**](Statement.md) | A description of what environments and projects the Relay Proxy should include or exclude. To learn more, read [Write an inline policy](https://docs.launchdarkly.com/sdk/relay-proxy/automatic-configuration#write-an-inline-policy). | 

## Methods

### NewRelayAutoConfigPost

`func NewRelayAutoConfigPost(name string, policy []Statement, ) *RelayAutoConfigPost`

NewRelayAutoConfigPost instantiates a new RelayAutoConfigPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelayAutoConfigPostWithDefaults

`func NewRelayAutoConfigPostWithDefaults() *RelayAutoConfigPost`

NewRelayAutoConfigPostWithDefaults instantiates a new RelayAutoConfigPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RelayAutoConfigPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RelayAutoConfigPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RelayAutoConfigPost) SetName(v string)`

SetName sets Name field to given value.


### GetPolicy

`func (o *RelayAutoConfigPost) GetPolicy() []Statement`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *RelayAutoConfigPost) GetPolicyOk() (*[]Statement, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *RelayAutoConfigPost) SetPolicy(v []Statement)`

SetPolicy sets Policy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


