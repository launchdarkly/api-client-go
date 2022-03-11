# FlagLinkRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) |  | 
**Key** | Pointer to **string** |  | [optional] 
**IntegrationKey** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**DeepLink** | **string** |  | 
**Timestamp** | [**TimestampRep**](TimestampRep.md) |  | 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**CreatedAt** | **int64** |  | 
**Member** | Pointer to [**FlagLinkMember**](FlagLinkMember.md) |  | [optional] 

## Methods

### NewFlagLinkRep

`func NewFlagLinkRep(links map[string]Link, id string, deepLink string, timestamp TimestampRep, createdAt int64, ) *FlagLinkRep`

NewFlagLinkRep instantiates a new FlagLinkRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagLinkRepWithDefaults

`func NewFlagLinkRepWithDefaults() *FlagLinkRep`

NewFlagLinkRepWithDefaults instantiates a new FlagLinkRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FlagLinkRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlagLinkRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlagLinkRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetKey

`func (o *FlagLinkRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlagLinkRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlagLinkRep) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FlagLinkRep) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetIntegrationKey

`func (o *FlagLinkRep) GetIntegrationKey() string`

GetIntegrationKey returns the IntegrationKey field if non-nil, zero value otherwise.

### GetIntegrationKeyOk

`func (o *FlagLinkRep) GetIntegrationKeyOk() (*string, bool)`

GetIntegrationKeyOk returns a tuple with the IntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKey

`func (o *FlagLinkRep) SetIntegrationKey(v string)`

SetIntegrationKey sets IntegrationKey field to given value.

### HasIntegrationKey

`func (o *FlagLinkRep) HasIntegrationKey() bool`

HasIntegrationKey returns a boolean if a field has been set.

### GetId

`func (o *FlagLinkRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlagLinkRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlagLinkRep) SetId(v string)`

SetId sets Id field to given value.


### GetDeepLink

`func (o *FlagLinkRep) GetDeepLink() string`

GetDeepLink returns the DeepLink field if non-nil, zero value otherwise.

### GetDeepLinkOk

`func (o *FlagLinkRep) GetDeepLinkOk() (*string, bool)`

GetDeepLinkOk returns a tuple with the DeepLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepLink

`func (o *FlagLinkRep) SetDeepLink(v string)`

SetDeepLink sets DeepLink field to given value.


### GetTimestamp

`func (o *FlagLinkRep) GetTimestamp() TimestampRep`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FlagLinkRep) GetTimestampOk() (*TimestampRep, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FlagLinkRep) SetTimestamp(v TimestampRep)`

SetTimestamp sets Timestamp field to given value.


### GetTitle

`func (o *FlagLinkRep) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FlagLinkRep) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FlagLinkRep) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FlagLinkRep) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *FlagLinkRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlagLinkRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlagLinkRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlagLinkRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *FlagLinkRep) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FlagLinkRep) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FlagLinkRep) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FlagLinkRep) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FlagLinkRep) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FlagLinkRep) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FlagLinkRep) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetMember

`func (o *FlagLinkRep) GetMember() FlagLinkMember`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *FlagLinkRep) GetMemberOk() (*FlagLinkMember, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *FlagLinkRep) SetMember(v FlagLinkMember)`

SetMember sets Member field to given value.

### HasMember

`func (o *FlagLinkRep) HasMember() bool`

HasMember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


