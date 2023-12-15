# Manifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** |  | [optional] 
**Capabilities** | Pointer to [**Capabilities**](Capabilities.md) |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**FormVariables** | Pointer to [**[]FormVariable**](FormVariable.md) |  | [optional] 
**HideOnIntegrationsPage** | Pointer to **bool** |  | [optional] 
**Icons** | Pointer to [**Icons**](Icons.md) |  | [optional] 
**Legacy** | Pointer to [**Legacy**](Legacy.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OauthIntegrationKey** | Pointer to **string** |  | [optional] 
**OtherCapabilities** | Pointer to **[]string** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**RequiresOAuth** | Pointer to **bool** |  | [optional] 
**SupportEmail** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewManifest

`func NewManifest() *Manifest`

NewManifest instantiates a new Manifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManifestWithDefaults

`func NewManifestWithDefaults() *Manifest`

NewManifestWithDefaults instantiates a new Manifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *Manifest) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Manifest) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Manifest) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Manifest) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCapabilities

`func (o *Manifest) GetCapabilities() Capabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Manifest) GetCapabilitiesOk() (*Capabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Manifest) SetCapabilities(v Capabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *Manifest) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCategories

`func (o *Manifest) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Manifest) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Manifest) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Manifest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetDescription

`func (o *Manifest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Manifest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Manifest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Manifest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetails

`func (o *Manifest) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Manifest) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Manifest) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Manifest) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetFormVariables

`func (o *Manifest) GetFormVariables() []FormVariable`

GetFormVariables returns the FormVariables field if non-nil, zero value otherwise.

### GetFormVariablesOk

`func (o *Manifest) GetFormVariablesOk() (*[]FormVariable, bool)`

GetFormVariablesOk returns a tuple with the FormVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormVariables

`func (o *Manifest) SetFormVariables(v []FormVariable)`

SetFormVariables sets FormVariables field to given value.

### HasFormVariables

`func (o *Manifest) HasFormVariables() bool`

HasFormVariables returns a boolean if a field has been set.

### GetHideOnIntegrationsPage

`func (o *Manifest) GetHideOnIntegrationsPage() bool`

GetHideOnIntegrationsPage returns the HideOnIntegrationsPage field if non-nil, zero value otherwise.

### GetHideOnIntegrationsPageOk

`func (o *Manifest) GetHideOnIntegrationsPageOk() (*bool, bool)`

GetHideOnIntegrationsPageOk returns a tuple with the HideOnIntegrationsPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideOnIntegrationsPage

`func (o *Manifest) SetHideOnIntegrationsPage(v bool)`

SetHideOnIntegrationsPage sets HideOnIntegrationsPage field to given value.

### HasHideOnIntegrationsPage

`func (o *Manifest) HasHideOnIntegrationsPage() bool`

HasHideOnIntegrationsPage returns a boolean if a field has been set.

### GetIcons

`func (o *Manifest) GetIcons() Icons`

GetIcons returns the Icons field if non-nil, zero value otherwise.

### GetIconsOk

`func (o *Manifest) GetIconsOk() (*Icons, bool)`

GetIconsOk returns a tuple with the Icons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcons

`func (o *Manifest) SetIcons(v Icons)`

SetIcons sets Icons field to given value.

### HasIcons

`func (o *Manifest) HasIcons() bool`

HasIcons returns a boolean if a field has been set.

### GetLegacy

`func (o *Manifest) GetLegacy() Legacy`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *Manifest) GetLegacyOk() (*Legacy, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *Manifest) SetLegacy(v Legacy)`

SetLegacy sets Legacy field to given value.

### HasLegacy

`func (o *Manifest) HasLegacy() bool`

HasLegacy returns a boolean if a field has been set.

### GetLinks

`func (o *Manifest) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Manifest) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Manifest) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Manifest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *Manifest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Manifest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Manifest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Manifest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOauthIntegrationKey

`func (o *Manifest) GetOauthIntegrationKey() string`

GetOauthIntegrationKey returns the OauthIntegrationKey field if non-nil, zero value otherwise.

### GetOauthIntegrationKeyOk

`func (o *Manifest) GetOauthIntegrationKeyOk() (*string, bool)`

GetOauthIntegrationKeyOk returns a tuple with the OauthIntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthIntegrationKey

`func (o *Manifest) SetOauthIntegrationKey(v string)`

SetOauthIntegrationKey sets OauthIntegrationKey field to given value.

### HasOauthIntegrationKey

`func (o *Manifest) HasOauthIntegrationKey() bool`

HasOauthIntegrationKey returns a boolean if a field has been set.

### GetOtherCapabilities

`func (o *Manifest) GetOtherCapabilities() []string`

GetOtherCapabilities returns the OtherCapabilities field if non-nil, zero value otherwise.

### GetOtherCapabilitiesOk

`func (o *Manifest) GetOtherCapabilitiesOk() (*[]string, bool)`

GetOtherCapabilitiesOk returns a tuple with the OtherCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCapabilities

`func (o *Manifest) SetOtherCapabilities(v []string)`

SetOtherCapabilities sets OtherCapabilities field to given value.

### HasOtherCapabilities

`func (o *Manifest) HasOtherCapabilities() bool`

HasOtherCapabilities returns a boolean if a field has been set.

### GetOverview

`func (o *Manifest) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *Manifest) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *Manifest) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *Manifest) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetRequiresOAuth

`func (o *Manifest) GetRequiresOAuth() bool`

GetRequiresOAuth returns the RequiresOAuth field if non-nil, zero value otherwise.

### GetRequiresOAuthOk

`func (o *Manifest) GetRequiresOAuthOk() (*bool, bool)`

GetRequiresOAuthOk returns a tuple with the RequiresOAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresOAuth

`func (o *Manifest) SetRequiresOAuth(v bool)`

SetRequiresOAuth sets RequiresOAuth field to given value.

### HasRequiresOAuth

`func (o *Manifest) HasRequiresOAuth() bool`

HasRequiresOAuth returns a boolean if a field has been set.

### GetSupportEmail

`func (o *Manifest) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *Manifest) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *Manifest) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.

### HasSupportEmail

`func (o *Manifest) HasSupportEmail() bool`

HasSupportEmail returns a boolean if a field has been set.

### GetVersion

`func (o *Manifest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Manifest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Manifest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Manifest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


