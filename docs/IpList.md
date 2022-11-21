# IpList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** | A list of the IP addresses LaunchDarkly&#39;s service uses | 
**OutboundAddresses** | **[]string** | A list of the IP addresses outgoing webhook notifications use | 

## Methods

### NewIpList

`func NewIpList(addresses []string, outboundAddresses []string, ) *IpList`

NewIpList instantiates a new IpList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpListWithDefaults

`func NewIpListWithDefaults() *IpList`

NewIpListWithDefaults instantiates a new IpList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *IpList) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *IpList) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *IpList) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetOutboundAddresses

`func (o *IpList) GetOutboundAddresses() []string`

GetOutboundAddresses returns the OutboundAddresses field if non-nil, zero value otherwise.

### GetOutboundAddressesOk

`func (o *IpList) GetOutboundAddressesOk() (*[]string, bool)`

GetOutboundAddressesOk returns a tuple with the OutboundAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundAddresses

`func (o *IpList) SetOutboundAddresses(v []string)`

SetOutboundAddresses sets OutboundAddresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


