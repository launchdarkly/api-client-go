# FlagLinkPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The flag link key | [optional] 
**IntegrationKey** | Pointer to **string** | The integration key for an integration whose &lt;code&gt;manifest.json&lt;/code&gt; includes the &lt;code&gt;flagLink&lt;/code&gt; capability, if this is a flag link for an existing integration. Do not include for URL flag links. | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**DeepLink** | Pointer to **string** | The URL for the external resource you are linking the flag to | [optional] 
**Title** | Pointer to **string** | The title of the flag link | [optional] 
**Description** | Pointer to **string** | The description of the flag link | [optional] 
**Metadata** | Pointer to **map[string]string** | The metadata required by this integration in order to create a flag link, if this is a flag link for an existing integration. Defined in the integration&#39;s &lt;code&gt;manifest.json&lt;/code&gt; file under &lt;code&gt;flagLink&lt;/code&gt;. | [optional] 

## Methods

### NewFlagLinkPost

`func NewFlagLinkPost() *FlagLinkPost`

NewFlagLinkPost instantiates a new FlagLinkPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagLinkPostWithDefaults

`func NewFlagLinkPostWithDefaults() *FlagLinkPost`

NewFlagLinkPostWithDefaults instantiates a new FlagLinkPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FlagLinkPost) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlagLinkPost) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlagLinkPost) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FlagLinkPost) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetIntegrationKey

`func (o *FlagLinkPost) GetIntegrationKey() string`

GetIntegrationKey returns the IntegrationKey field if non-nil, zero value otherwise.

### GetIntegrationKeyOk

`func (o *FlagLinkPost) GetIntegrationKeyOk() (*string, bool)`

GetIntegrationKeyOk returns a tuple with the IntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKey

`func (o *FlagLinkPost) SetIntegrationKey(v string)`

SetIntegrationKey sets IntegrationKey field to given value.

### HasIntegrationKey

`func (o *FlagLinkPost) HasIntegrationKey() bool`

HasIntegrationKey returns a boolean if a field has been set.

### GetTimestamp

`func (o *FlagLinkPost) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FlagLinkPost) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FlagLinkPost) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *FlagLinkPost) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDeepLink

`func (o *FlagLinkPost) GetDeepLink() string`

GetDeepLink returns the DeepLink field if non-nil, zero value otherwise.

### GetDeepLinkOk

`func (o *FlagLinkPost) GetDeepLinkOk() (*string, bool)`

GetDeepLinkOk returns a tuple with the DeepLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepLink

`func (o *FlagLinkPost) SetDeepLink(v string)`

SetDeepLink sets DeepLink field to given value.

### HasDeepLink

`func (o *FlagLinkPost) HasDeepLink() bool`

HasDeepLink returns a boolean if a field has been set.

### GetTitle

`func (o *FlagLinkPost) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FlagLinkPost) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FlagLinkPost) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FlagLinkPost) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *FlagLinkPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlagLinkPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlagLinkPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlagLinkPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *FlagLinkPost) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FlagLinkPost) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FlagLinkPost) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FlagLinkPost) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


