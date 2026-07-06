# RedshiftDataExportCompletedArtifactsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SqlSetupScript** | Pointer to **string** | The SQL setup script originally run against the Redshift cluster, rehydrated from the destination&#39;s persisted custom names. | [optional] 
**S3BucketName** | Pointer to **string** | The auto-generated S3 staging bucket name. | [optional] 

## Methods

### NewRedshiftDataExportCompletedArtifactsRep

`func NewRedshiftDataExportCompletedArtifactsRep() *RedshiftDataExportCompletedArtifactsRep`

NewRedshiftDataExportCompletedArtifactsRep instantiates a new RedshiftDataExportCompletedArtifactsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedshiftDataExportCompletedArtifactsRepWithDefaults

`func NewRedshiftDataExportCompletedArtifactsRepWithDefaults() *RedshiftDataExportCompletedArtifactsRep`

NewRedshiftDataExportCompletedArtifactsRepWithDefaults instantiates a new RedshiftDataExportCompletedArtifactsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSqlSetupScript

`func (o *RedshiftDataExportCompletedArtifactsRep) GetSqlSetupScript() string`

GetSqlSetupScript returns the SqlSetupScript field if non-nil, zero value otherwise.

### GetSqlSetupScriptOk

`func (o *RedshiftDataExportCompletedArtifactsRep) GetSqlSetupScriptOk() (*string, bool)`

GetSqlSetupScriptOk returns a tuple with the SqlSetupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlSetupScript

`func (o *RedshiftDataExportCompletedArtifactsRep) SetSqlSetupScript(v string)`

SetSqlSetupScript sets SqlSetupScript field to given value.

### HasSqlSetupScript

`func (o *RedshiftDataExportCompletedArtifactsRep) HasSqlSetupScript() bool`

HasSqlSetupScript returns a boolean if a field has been set.

### GetS3BucketName

`func (o *RedshiftDataExportCompletedArtifactsRep) GetS3BucketName() string`

GetS3BucketName returns the S3BucketName field if non-nil, zero value otherwise.

### GetS3BucketNameOk

`func (o *RedshiftDataExportCompletedArtifactsRep) GetS3BucketNameOk() (*string, bool)`

GetS3BucketNameOk returns a tuple with the S3BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketName

`func (o *RedshiftDataExportCompletedArtifactsRep) SetS3BucketName(v string)`

SetS3BucketName sets S3BucketName field to given value.

### HasS3BucketName

`func (o *RedshiftDataExportCompletedArtifactsRep) HasS3BucketName() bool`

HasS3BucketName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


