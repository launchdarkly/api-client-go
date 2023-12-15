# BillingContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**Address1** | Pointer to **string** |  | [optional] 
**Address2** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**PoNumber** | Pointer to **string** |  | [optional] 
**TaxAddress** | Pointer to [**TaxAddress**](TaxAddress.md) |  | [optional] 

## Methods

### NewBillingContact

`func NewBillingContact() *BillingContact`

NewBillingContact instantiates a new BillingContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingContactWithDefaults

`func NewBillingContactWithDefaults() *BillingContact`

NewBillingContactWithDefaults instantiates a new BillingContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BillingContact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingContact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingContact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingContact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *BillingContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BillingContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BillingContact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BillingContact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCompany

`func (o *BillingContact) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *BillingContact) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *BillingContact) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *BillingContact) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetAddress1

`func (o *BillingContact) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *BillingContact) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *BillingContact) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *BillingContact) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *BillingContact) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *BillingContact) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *BillingContact) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *BillingContact) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *BillingContact) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BillingContact) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BillingContact) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *BillingContact) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *BillingContact) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BillingContact) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BillingContact) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BillingContact) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *BillingContact) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *BillingContact) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *BillingContact) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *BillingContact) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *BillingContact) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BillingContact) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BillingContact) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *BillingContact) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPoNumber

`func (o *BillingContact) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *BillingContact) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *BillingContact) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *BillingContact) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetTaxAddress

`func (o *BillingContact) GetTaxAddress() TaxAddress`

GetTaxAddress returns the TaxAddress field if non-nil, zero value otherwise.

### GetTaxAddressOk

`func (o *BillingContact) GetTaxAddressOk() (*TaxAddress, bool)`

GetTaxAddressOk returns a tuple with the TaxAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAddress

`func (o *BillingContact) SetTaxAddress(v TaxAddress)`

SetTaxAddress sets TaxAddress field to given value.

### HasTaxAddress

`func (o *BillingContact) HasTaxAddress() bool`

HasTaxAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


