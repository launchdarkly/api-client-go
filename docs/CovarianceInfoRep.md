# CovarianceInfoRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the covariance matrix | 
**FileName** | **string** | The file name of the uploaded covariance matrix | 
**CreatedAt** | **int64** |  | 

## Methods

### NewCovarianceInfoRep

`func NewCovarianceInfoRep(id string, fileName string, createdAt int64, ) *CovarianceInfoRep`

NewCovarianceInfoRep instantiates a new CovarianceInfoRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCovarianceInfoRepWithDefaults

`func NewCovarianceInfoRepWithDefaults() *CovarianceInfoRep`

NewCovarianceInfoRepWithDefaults instantiates a new CovarianceInfoRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CovarianceInfoRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CovarianceInfoRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CovarianceInfoRep) SetId(v string)`

SetId sets Id field to given value.


### GetFileName

`func (o *CovarianceInfoRep) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CovarianceInfoRep) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CovarianceInfoRep) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetCreatedAt

`func (o *CovarianceInfoRep) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CovarianceInfoRep) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CovarianceInfoRep) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


