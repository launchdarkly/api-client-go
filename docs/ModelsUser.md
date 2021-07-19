# ModelsUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Secondary** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Anonymous** | Pointer to **bool** |  | [optional] 
**Custom** | Pointer to **map[string]interface{}** |  | [optional] 
**Derived** | Pointer to [**map[string]ModelsUserDerived**](ModelsUserDerived.md) |  | [optional] 
**PrivateAttrs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewModelsUser

`func NewModelsUser() *ModelsUser`

NewModelsUser instantiates a new ModelsUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsUserWithDefaults

`func NewModelsUserWithDefaults() *ModelsUser`

NewModelsUserWithDefaults instantiates a new ModelsUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ModelsUser) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ModelsUser) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ModelsUser) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ModelsUser) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSecondary

`func (o *ModelsUser) GetSecondary() string`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *ModelsUser) GetSecondaryOk() (*string, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *ModelsUser) SetSecondary(v string)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *ModelsUser) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.

### GetIp

`func (o *ModelsUser) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ModelsUser) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ModelsUser) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ModelsUser) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetCountry

`func (o *ModelsUser) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ModelsUser) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ModelsUser) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ModelsUser) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *ModelsUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ModelsUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ModelsUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ModelsUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *ModelsUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ModelsUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ModelsUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ModelsUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ModelsUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ModelsUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ModelsUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ModelsUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetAvatar

`func (o *ModelsUser) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ModelsUser) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ModelsUser) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *ModelsUser) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetName

`func (o *ModelsUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnonymous

`func (o *ModelsUser) GetAnonymous() bool`

GetAnonymous returns the Anonymous field if non-nil, zero value otherwise.

### GetAnonymousOk

`func (o *ModelsUser) GetAnonymousOk() (*bool, bool)`

GetAnonymousOk returns a tuple with the Anonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymous

`func (o *ModelsUser) SetAnonymous(v bool)`

SetAnonymous sets Anonymous field to given value.

### HasAnonymous

`func (o *ModelsUser) HasAnonymous() bool`

HasAnonymous returns a boolean if a field has been set.

### GetCustom

`func (o *ModelsUser) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *ModelsUser) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *ModelsUser) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *ModelsUser) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetDerived

`func (o *ModelsUser) GetDerived() map[string]ModelsUserDerived`

GetDerived returns the Derived field if non-nil, zero value otherwise.

### GetDerivedOk

`func (o *ModelsUser) GetDerivedOk() (*map[string]ModelsUserDerived, bool)`

GetDerivedOk returns a tuple with the Derived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerived

`func (o *ModelsUser) SetDerived(v map[string]ModelsUserDerived)`

SetDerived sets Derived field to given value.

### HasDerived

`func (o *ModelsUser) HasDerived() bool`

HasDerived returns a boolean if a field has been set.

### GetPrivateAttrs

`func (o *ModelsUser) GetPrivateAttrs() []string`

GetPrivateAttrs returns the PrivateAttrs field if non-nil, zero value otherwise.

### GetPrivateAttrsOk

`func (o *ModelsUser) GetPrivateAttrsOk() (*[]string, bool)`

GetPrivateAttrsOk returns a tuple with the PrivateAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateAttrs

`func (o *ModelsUser) SetPrivateAttrs(v []string)`

SetPrivateAttrs sets PrivateAttrs field to given value.

### HasPrivateAttrs

`func (o *ModelsUser) HasPrivateAttrs() bool`

HasPrivateAttrs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


