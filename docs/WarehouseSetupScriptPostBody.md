# WarehouseSetupScriptPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**SnowflakeHostAddress** | Pointer to **string** |  | [optional] 
**DatabaseName** | Pointer to **string** |  | [optional] 
**WarehouseName** | Pointer to **string** |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 
**SchemaName** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**IncludeNetworkPolicy** | Pointer to **bool** |  | [optional] 
**ClusterIdentifier** | Pointer to **string** |  | [optional] 
**ClusterRegion** | Pointer to **string** |  | [optional] 
**ClusterAwsAccountId** | Pointer to **string** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**ClickHouseDatabaseName** | Pointer to **string** |  | [optional] 
**ClickHouseUserName** | Pointer to **string** |  | [optional] 
**ClickHouseS3BucketName** | Pointer to **string** |  | [optional] 
**ClickHouseIncludeHostRestriction** | Pointer to **bool** |  | [optional] 
**ClickHouseServiceRoleArn** | Pointer to **string** |  | [optional] 
**ClickHousePassword** | Pointer to **string** |  | [optional] 

## Methods

### NewWarehouseSetupScriptPostBody

`func NewWarehouseSetupScriptPostBody() *WarehouseSetupScriptPostBody`

NewWarehouseSetupScriptPostBody instantiates a new WarehouseSetupScriptPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseSetupScriptPostBodyWithDefaults

`func NewWarehouseSetupScriptPostBodyWithDefaults() *WarehouseSetupScriptPostBody`

NewWarehouseSetupScriptPostBodyWithDefaults instantiates a new WarehouseSetupScriptPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WarehouseSetupScriptPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WarehouseSetupScriptPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WarehouseSetupScriptPostBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WarehouseSetupScriptPostBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSnowflakeHostAddress

`func (o *WarehouseSetupScriptPostBody) GetSnowflakeHostAddress() string`

GetSnowflakeHostAddress returns the SnowflakeHostAddress field if non-nil, zero value otherwise.

### GetSnowflakeHostAddressOk

`func (o *WarehouseSetupScriptPostBody) GetSnowflakeHostAddressOk() (*string, bool)`

GetSnowflakeHostAddressOk returns a tuple with the SnowflakeHostAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnowflakeHostAddress

`func (o *WarehouseSetupScriptPostBody) SetSnowflakeHostAddress(v string)`

SetSnowflakeHostAddress sets SnowflakeHostAddress field to given value.

### HasSnowflakeHostAddress

`func (o *WarehouseSetupScriptPostBody) HasSnowflakeHostAddress() bool`

HasSnowflakeHostAddress returns a boolean if a field has been set.

### GetDatabaseName

`func (o *WarehouseSetupScriptPostBody) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *WarehouseSetupScriptPostBody) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *WarehouseSetupScriptPostBody) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *WarehouseSetupScriptPostBody) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetWarehouseName

`func (o *WarehouseSetupScriptPostBody) GetWarehouseName() string`

GetWarehouseName returns the WarehouseName field if non-nil, zero value otherwise.

### GetWarehouseNameOk

`func (o *WarehouseSetupScriptPostBody) GetWarehouseNameOk() (*string, bool)`

GetWarehouseNameOk returns a tuple with the WarehouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseName

`func (o *WarehouseSetupScriptPostBody) SetWarehouseName(v string)`

SetWarehouseName sets WarehouseName field to given value.

### HasWarehouseName

`func (o *WarehouseSetupScriptPostBody) HasWarehouseName() bool`

HasWarehouseName returns a boolean if a field has been set.

### GetRoleName

`func (o *WarehouseSetupScriptPostBody) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *WarehouseSetupScriptPostBody) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *WarehouseSetupScriptPostBody) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *WarehouseSetupScriptPostBody) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetSchemaName

`func (o *WarehouseSetupScriptPostBody) GetSchemaName() string`

GetSchemaName returns the SchemaName field if non-nil, zero value otherwise.

### GetSchemaNameOk

`func (o *WarehouseSetupScriptPostBody) GetSchemaNameOk() (*string, bool)`

GetSchemaNameOk returns a tuple with the SchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaName

`func (o *WarehouseSetupScriptPostBody) SetSchemaName(v string)`

SetSchemaName sets SchemaName field to given value.

### HasSchemaName

`func (o *WarehouseSetupScriptPostBody) HasSchemaName() bool`

HasSchemaName returns a boolean if a field has been set.

### GetUserName

`func (o *WarehouseSetupScriptPostBody) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *WarehouseSetupScriptPostBody) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *WarehouseSetupScriptPostBody) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *WarehouseSetupScriptPostBody) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetIncludeNetworkPolicy

`func (o *WarehouseSetupScriptPostBody) GetIncludeNetworkPolicy() bool`

GetIncludeNetworkPolicy returns the IncludeNetworkPolicy field if non-nil, zero value otherwise.

### GetIncludeNetworkPolicyOk

`func (o *WarehouseSetupScriptPostBody) GetIncludeNetworkPolicyOk() (*bool, bool)`

GetIncludeNetworkPolicyOk returns a tuple with the IncludeNetworkPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNetworkPolicy

`func (o *WarehouseSetupScriptPostBody) SetIncludeNetworkPolicy(v bool)`

SetIncludeNetworkPolicy sets IncludeNetworkPolicy field to given value.

### HasIncludeNetworkPolicy

`func (o *WarehouseSetupScriptPostBody) HasIncludeNetworkPolicy() bool`

HasIncludeNetworkPolicy returns a boolean if a field has been set.

### GetClusterIdentifier

`func (o *WarehouseSetupScriptPostBody) GetClusterIdentifier() string`

GetClusterIdentifier returns the ClusterIdentifier field if non-nil, zero value otherwise.

### GetClusterIdentifierOk

`func (o *WarehouseSetupScriptPostBody) GetClusterIdentifierOk() (*string, bool)`

GetClusterIdentifierOk returns a tuple with the ClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifier

`func (o *WarehouseSetupScriptPostBody) SetClusterIdentifier(v string)`

SetClusterIdentifier sets ClusterIdentifier field to given value.

### HasClusterIdentifier

`func (o *WarehouseSetupScriptPostBody) HasClusterIdentifier() bool`

HasClusterIdentifier returns a boolean if a field has been set.

### GetClusterRegion

`func (o *WarehouseSetupScriptPostBody) GetClusterRegion() string`

GetClusterRegion returns the ClusterRegion field if non-nil, zero value otherwise.

### GetClusterRegionOk

`func (o *WarehouseSetupScriptPostBody) GetClusterRegionOk() (*string, bool)`

GetClusterRegionOk returns a tuple with the ClusterRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRegion

`func (o *WarehouseSetupScriptPostBody) SetClusterRegion(v string)`

SetClusterRegion sets ClusterRegion field to given value.

### HasClusterRegion

`func (o *WarehouseSetupScriptPostBody) HasClusterRegion() bool`

HasClusterRegion returns a boolean if a field has been set.

### GetClusterAwsAccountId

`func (o *WarehouseSetupScriptPostBody) GetClusterAwsAccountId() string`

GetClusterAwsAccountId returns the ClusterAwsAccountId field if non-nil, zero value otherwise.

### GetClusterAwsAccountIdOk

`func (o *WarehouseSetupScriptPostBody) GetClusterAwsAccountIdOk() (*string, bool)`

GetClusterAwsAccountIdOk returns a tuple with the ClusterAwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterAwsAccountId

`func (o *WarehouseSetupScriptPostBody) SetClusterAwsAccountId(v string)`

SetClusterAwsAccountId sets ClusterAwsAccountId field to given value.

### HasClusterAwsAccountId

`func (o *WarehouseSetupScriptPostBody) HasClusterAwsAccountId() bool`

HasClusterAwsAccountId returns a boolean if a field has been set.

### GetEndpoint

`func (o *WarehouseSetupScriptPostBody) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *WarehouseSetupScriptPostBody) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *WarehouseSetupScriptPostBody) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *WarehouseSetupScriptPostBody) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetClickHouseDatabaseName

`func (o *WarehouseSetupScriptPostBody) GetClickHouseDatabaseName() string`

GetClickHouseDatabaseName returns the ClickHouseDatabaseName field if non-nil, zero value otherwise.

### GetClickHouseDatabaseNameOk

`func (o *WarehouseSetupScriptPostBody) GetClickHouseDatabaseNameOk() (*string, bool)`

GetClickHouseDatabaseNameOk returns a tuple with the ClickHouseDatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickHouseDatabaseName

`func (o *WarehouseSetupScriptPostBody) SetClickHouseDatabaseName(v string)`

SetClickHouseDatabaseName sets ClickHouseDatabaseName field to given value.

### HasClickHouseDatabaseName

`func (o *WarehouseSetupScriptPostBody) HasClickHouseDatabaseName() bool`

HasClickHouseDatabaseName returns a boolean if a field has been set.

### GetClickHouseUserName

`func (o *WarehouseSetupScriptPostBody) GetClickHouseUserName() string`

GetClickHouseUserName returns the ClickHouseUserName field if non-nil, zero value otherwise.

### GetClickHouseUserNameOk

`func (o *WarehouseSetupScriptPostBody) GetClickHouseUserNameOk() (*string, bool)`

GetClickHouseUserNameOk returns a tuple with the ClickHouseUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickHouseUserName

`func (o *WarehouseSetupScriptPostBody) SetClickHouseUserName(v string)`

SetClickHouseUserName sets ClickHouseUserName field to given value.

### HasClickHouseUserName

`func (o *WarehouseSetupScriptPostBody) HasClickHouseUserName() bool`

HasClickHouseUserName returns a boolean if a field has been set.

### GetClickHouseS3BucketName

`func (o *WarehouseSetupScriptPostBody) GetClickHouseS3BucketName() string`

GetClickHouseS3BucketName returns the ClickHouseS3BucketName field if non-nil, zero value otherwise.

### GetClickHouseS3BucketNameOk

`func (o *WarehouseSetupScriptPostBody) GetClickHouseS3BucketNameOk() (*string, bool)`

GetClickHouseS3BucketNameOk returns a tuple with the ClickHouseS3BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickHouseS3BucketName

`func (o *WarehouseSetupScriptPostBody) SetClickHouseS3BucketName(v string)`

SetClickHouseS3BucketName sets ClickHouseS3BucketName field to given value.

### HasClickHouseS3BucketName

`func (o *WarehouseSetupScriptPostBody) HasClickHouseS3BucketName() bool`

HasClickHouseS3BucketName returns a boolean if a field has been set.

### GetClickHouseIncludeHostRestriction

`func (o *WarehouseSetupScriptPostBody) GetClickHouseIncludeHostRestriction() bool`

GetClickHouseIncludeHostRestriction returns the ClickHouseIncludeHostRestriction field if non-nil, zero value otherwise.

### GetClickHouseIncludeHostRestrictionOk

`func (o *WarehouseSetupScriptPostBody) GetClickHouseIncludeHostRestrictionOk() (*bool, bool)`

GetClickHouseIncludeHostRestrictionOk returns a tuple with the ClickHouseIncludeHostRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickHouseIncludeHostRestriction

`func (o *WarehouseSetupScriptPostBody) SetClickHouseIncludeHostRestriction(v bool)`

SetClickHouseIncludeHostRestriction sets ClickHouseIncludeHostRestriction field to given value.

### HasClickHouseIncludeHostRestriction

`func (o *WarehouseSetupScriptPostBody) HasClickHouseIncludeHostRestriction() bool`

HasClickHouseIncludeHostRestriction returns a boolean if a field has been set.

### GetClickHouseServiceRoleArn

`func (o *WarehouseSetupScriptPostBody) GetClickHouseServiceRoleArn() string`

GetClickHouseServiceRoleArn returns the ClickHouseServiceRoleArn field if non-nil, zero value otherwise.

### GetClickHouseServiceRoleArnOk

`func (o *WarehouseSetupScriptPostBody) GetClickHouseServiceRoleArnOk() (*string, bool)`

GetClickHouseServiceRoleArnOk returns a tuple with the ClickHouseServiceRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickHouseServiceRoleArn

`func (o *WarehouseSetupScriptPostBody) SetClickHouseServiceRoleArn(v string)`

SetClickHouseServiceRoleArn sets ClickHouseServiceRoleArn field to given value.

### HasClickHouseServiceRoleArn

`func (o *WarehouseSetupScriptPostBody) HasClickHouseServiceRoleArn() bool`

HasClickHouseServiceRoleArn returns a boolean if a field has been set.

### GetClickHousePassword

`func (o *WarehouseSetupScriptPostBody) GetClickHousePassword() string`

GetClickHousePassword returns the ClickHousePassword field if non-nil, zero value otherwise.

### GetClickHousePasswordOk

`func (o *WarehouseSetupScriptPostBody) GetClickHousePasswordOk() (*string, bool)`

GetClickHousePasswordOk returns a tuple with the ClickHousePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickHousePassword

`func (o *WarehouseSetupScriptPostBody) SetClickHousePassword(v string)`

SetClickHousePassword sets ClickHousePassword field to given value.

### HasClickHousePassword

`func (o *WarehouseSetupScriptPostBody) HasClickHousePassword() bool`

HasClickHousePassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


