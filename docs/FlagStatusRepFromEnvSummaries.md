# FlagStatusRepFromEnvSummaries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environments** | Pointer to [**map[string]FlagStatusesRep**](FlagStatusesRep.md) |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**map[string]CoreLink**](CoreLink.md) |  | [optional] 

## Methods

### NewFlagStatusRepFromEnvSummaries

`func NewFlagStatusRepFromEnvSummaries() *FlagStatusRepFromEnvSummaries`

NewFlagStatusRepFromEnvSummaries instantiates a new FlagStatusRepFromEnvSummaries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagStatusRepFromEnvSummariesWithDefaults

`func NewFlagStatusRepFromEnvSummariesWithDefaults() *FlagStatusRepFromEnvSummaries`

NewFlagStatusRepFromEnvSummariesWithDefaults instantiates a new FlagStatusRepFromEnvSummaries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironments

`func (o *FlagStatusRepFromEnvSummaries) GetEnvironments() map[string]FlagStatusesRep`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *FlagStatusRepFromEnvSummaries) GetEnvironmentsOk() (*map[string]FlagStatusesRep, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *FlagStatusRepFromEnvSummaries) SetEnvironments(v map[string]FlagStatusesRep)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *FlagStatusRepFromEnvSummaries) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetKey

`func (o *FlagStatusRepFromEnvSummaries) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlagStatusRepFromEnvSummaries) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlagStatusRepFromEnvSummaries) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FlagStatusRepFromEnvSummaries) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLinks

`func (o *FlagStatusRepFromEnvSummaries) GetLinks() map[string]CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlagStatusRepFromEnvSummaries) GetLinksOk() (*map[string]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlagStatusRepFromEnvSummaries) SetLinks(v map[string]CoreLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FlagStatusRepFromEnvSummaries) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


