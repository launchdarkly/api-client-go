# WarehouseDestinationSetupScriptRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | Pointer to **string** | The SQL setup script to run in your data warehouse | [optional] 
**PublicKey** | Pointer to **string** | The RSA public key (Snowflake only) to store as the destination public_key | [optional] 
**PublicKeyPkcs8** | Pointer to **string** | The PKCS8 RSA public key (Snowflake only) | [optional] 
**RedshiftIAMPermissionsPolicy** | Pointer to **string** | For Redshift, present only when clusterIdentifier, clusterRegion, and clusterAwsAccountId are supplied in the request body. | [optional] 
**RedshiftIAMTrustPolicy** | Pointer to **string** | For Redshift, present only when clusterIdentifier, clusterRegion, and clusterAwsAccountId are supplied in the request body. | [optional] 
**S3BucketName** | Pointer to **string** | The auto-generated S3 staging bucket name (ClickHouse and Redshift) | [optional] 

## Methods

### NewWarehouseDestinationSetupScriptRep

`func NewWarehouseDestinationSetupScriptRep() *WarehouseDestinationSetupScriptRep`

NewWarehouseDestinationSetupScriptRep instantiates a new WarehouseDestinationSetupScriptRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseDestinationSetupScriptRepWithDefaults

`func NewWarehouseDestinationSetupScriptRepWithDefaults() *WarehouseDestinationSetupScriptRep`

NewWarehouseDestinationSetupScriptRepWithDefaults instantiates a new WarehouseDestinationSetupScriptRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *WarehouseDestinationSetupScriptRep) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *WarehouseDestinationSetupScriptRep) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *WarehouseDestinationSetupScriptRep) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *WarehouseDestinationSetupScriptRep) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetPublicKey

`func (o *WarehouseDestinationSetupScriptRep) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *WarehouseDestinationSetupScriptRep) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *WarehouseDestinationSetupScriptRep) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *WarehouseDestinationSetupScriptRep) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetPublicKeyPkcs8

`func (o *WarehouseDestinationSetupScriptRep) GetPublicKeyPkcs8() string`

GetPublicKeyPkcs8 returns the PublicKeyPkcs8 field if non-nil, zero value otherwise.

### GetPublicKeyPkcs8Ok

`func (o *WarehouseDestinationSetupScriptRep) GetPublicKeyPkcs8Ok() (*string, bool)`

GetPublicKeyPkcs8Ok returns a tuple with the PublicKeyPkcs8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyPkcs8

`func (o *WarehouseDestinationSetupScriptRep) SetPublicKeyPkcs8(v string)`

SetPublicKeyPkcs8 sets PublicKeyPkcs8 field to given value.

### HasPublicKeyPkcs8

`func (o *WarehouseDestinationSetupScriptRep) HasPublicKeyPkcs8() bool`

HasPublicKeyPkcs8 returns a boolean if a field has been set.

### GetRedshiftIAMPermissionsPolicy

`func (o *WarehouseDestinationSetupScriptRep) GetRedshiftIAMPermissionsPolicy() string`

GetRedshiftIAMPermissionsPolicy returns the RedshiftIAMPermissionsPolicy field if non-nil, zero value otherwise.

### GetRedshiftIAMPermissionsPolicyOk

`func (o *WarehouseDestinationSetupScriptRep) GetRedshiftIAMPermissionsPolicyOk() (*string, bool)`

GetRedshiftIAMPermissionsPolicyOk returns a tuple with the RedshiftIAMPermissionsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftIAMPermissionsPolicy

`func (o *WarehouseDestinationSetupScriptRep) SetRedshiftIAMPermissionsPolicy(v string)`

SetRedshiftIAMPermissionsPolicy sets RedshiftIAMPermissionsPolicy field to given value.

### HasRedshiftIAMPermissionsPolicy

`func (o *WarehouseDestinationSetupScriptRep) HasRedshiftIAMPermissionsPolicy() bool`

HasRedshiftIAMPermissionsPolicy returns a boolean if a field has been set.

### GetRedshiftIAMTrustPolicy

`func (o *WarehouseDestinationSetupScriptRep) GetRedshiftIAMTrustPolicy() string`

GetRedshiftIAMTrustPolicy returns the RedshiftIAMTrustPolicy field if non-nil, zero value otherwise.

### GetRedshiftIAMTrustPolicyOk

`func (o *WarehouseDestinationSetupScriptRep) GetRedshiftIAMTrustPolicyOk() (*string, bool)`

GetRedshiftIAMTrustPolicyOk returns a tuple with the RedshiftIAMTrustPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedshiftIAMTrustPolicy

`func (o *WarehouseDestinationSetupScriptRep) SetRedshiftIAMTrustPolicy(v string)`

SetRedshiftIAMTrustPolicy sets RedshiftIAMTrustPolicy field to given value.

### HasRedshiftIAMTrustPolicy

`func (o *WarehouseDestinationSetupScriptRep) HasRedshiftIAMTrustPolicy() bool`

HasRedshiftIAMTrustPolicy returns a boolean if a field has been set.

### GetS3BucketName

`func (o *WarehouseDestinationSetupScriptRep) GetS3BucketName() string`

GetS3BucketName returns the S3BucketName field if non-nil, zero value otherwise.

### GetS3BucketNameOk

`func (o *WarehouseDestinationSetupScriptRep) GetS3BucketNameOk() (*string, bool)`

GetS3BucketNameOk returns a tuple with the S3BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketName

`func (o *WarehouseDestinationSetupScriptRep) SetS3BucketName(v string)`

SetS3BucketName sets S3BucketName field to given value.

### HasS3BucketName

`func (o *WarehouseDestinationSetupScriptRep) HasS3BucketName() bool`

HasS3BucketName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


