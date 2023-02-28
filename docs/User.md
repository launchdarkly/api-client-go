# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The user key. This is the only mandatory user attribute. | [optional] 
**Secondary** | Pointer to **string** | If provided, used with the user key to generate a variation in percentage rollouts | [optional] 
**Ip** | Pointer to **string** | The user&#39;s IP address | [optional] 
**Country** | Pointer to **string** | The user&#39;s country | [optional] 
**Email** | Pointer to **string** | The user&#39;s email | [optional] 
**FirstName** | Pointer to **string** | The user&#39;s first name | [optional] 
**LastName** | Pointer to **string** | The user&#39;s last name | [optional] 
**Avatar** | Pointer to **string** | An absolute URL to an avatar image. | [optional] 
**Name** | Pointer to **string** | The user&#39;s full name | [optional] 
**Anonymous** | Pointer to **bool** | Whether the user is anonymous. If true, this user does not appear on the Contexts list in the LaunchDarkly user interface. | [optional] 
**Custom** | Pointer to **map[string]interface{}** | Any other custom attributes for this user. Custom attributes contain any other user data that you would like to use to conditionally target your users. | [optional] 
**PrivateAttrs** | Pointer to **[]string** | A list of attribute names that are marked as private. You can use these attributes in targeting rules and segments. If you are using a server-side SDK, the SDK will not send the private attribute back to LaunchDarkly. If you are using a client-side SDK, the SDK will send the private attribute back to LaunchDarkly for evaluation. However, the SDK won&#39;t send the attribute to LaunchDarkly in events data, LaunchDarkly won&#39;t store the private attribute, and the private attribute will not appear on the Contexts list. | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *User) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *User) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *User) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *User) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSecondary

`func (o *User) GetSecondary() string`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *User) GetSecondaryOk() (*string, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *User) SetSecondary(v string)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *User) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.

### GetIp

`func (o *User) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *User) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *User) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *User) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetCountry

`func (o *User) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *User) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *User) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *User) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *User) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *User) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetAvatar

`func (o *User) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *User) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *User) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *User) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnonymous

`func (o *User) GetAnonymous() bool`

GetAnonymous returns the Anonymous field if non-nil, zero value otherwise.

### GetAnonymousOk

`func (o *User) GetAnonymousOk() (*bool, bool)`

GetAnonymousOk returns a tuple with the Anonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymous

`func (o *User) SetAnonymous(v bool)`

SetAnonymous sets Anonymous field to given value.

### HasAnonymous

`func (o *User) HasAnonymous() bool`

HasAnonymous returns a boolean if a field has been set.

### GetCustom

`func (o *User) GetCustom() map[string]interface{}`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *User) GetCustomOk() (*map[string]interface{}, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *User) SetCustom(v map[string]interface{})`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *User) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetPrivateAttrs

`func (o *User) GetPrivateAttrs() []string`

GetPrivateAttrs returns the PrivateAttrs field if non-nil, zero value otherwise.

### GetPrivateAttrsOk

`func (o *User) GetPrivateAttrsOk() (*[]string, bool)`

GetPrivateAttrsOk returns a tuple with the PrivateAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateAttrs

`func (o *User) SetPrivateAttrs(v []string)`

SetPrivateAttrs sets PrivateAttrs field to given value.

### HasPrivateAttrs

`func (o *User) HasPrivateAttrs() bool`

HasPrivateAttrs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


