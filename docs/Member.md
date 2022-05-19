# Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) |  | 
**Id** | **string** | The member&#39;s ID | 
**FirstName** | Pointer to **string** | The member&#39;s first name | [optional] 
**LastName** | Pointer to **string** | The member&#39;s last name | [optional] 
**Role** | **string** | The member&#39;s built-in role. If the member has no custom roles, this role will be in effect. | 
**Email** | **string** | The member&#39;s email address | 
**PendingInvite** | **bool** | Whether or not the member has a pending invitation | 
**Verified** | **bool** | Whether or not the member&#39;s email address has been verified | 
**PendingEmail** | Pointer to **string** |  | [optional] 
**CustomRoles** | **[]string** | The set of custom roles (as keys) assigned to the member | 
**Mfa** | **string** | Whether or not multi-factor authentication is enabled for this member | 
**ExcludedDashboards** | Pointer to **[]string** | Default dashboards that the member has chosen to ignore | [optional] 
**LastSeen** | **int64** |  | 
**LastSeenMetadata** | Pointer to [**LastSeenMetadata**](LastSeenMetadata.md) |  | [optional] 
**IntegrationMetadata** | Pointer to [**IntegrationMetadata**](IntegrationMetadata.md) |  | [optional] 
**Teams** | Pointer to [**[]MemberTeamSummaryRep**](MemberTeamSummaryRep.md) |  | [optional] 
**PermissionGrants** | Pointer to [**[]MemberPermissionGrantSummaryRep**](MemberPermissionGrantSummaryRep.md) |  | [optional] 
**CreationDate** | **int64** |  | 

## Methods

### NewMember

`func NewMember(links map[string]Link, id string, role string, email string, pendingInvite bool, verified bool, customRoles []string, mfa string, lastSeen int64, creationDate int64, ) *Member`

NewMember instantiates a new Member object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberWithDefaults

`func NewMemberWithDefaults() *Member`

NewMemberWithDefaults instantiates a new Member object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Member) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Member) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Member) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetId

`func (o *Member) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Member) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Member) SetId(v string)`

SetId sets Id field to given value.


### GetFirstName

`func (o *Member) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Member) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Member) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Member) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Member) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Member) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Member) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Member) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetRole

`func (o *Member) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Member) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Member) SetRole(v string)`

SetRole sets Role field to given value.


### GetEmail

`func (o *Member) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Member) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Member) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPendingInvite

`func (o *Member) GetPendingInvite() bool`

GetPendingInvite returns the PendingInvite field if non-nil, zero value otherwise.

### GetPendingInviteOk

`func (o *Member) GetPendingInviteOk() (*bool, bool)`

GetPendingInviteOk returns a tuple with the PendingInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingInvite

`func (o *Member) SetPendingInvite(v bool)`

SetPendingInvite sets PendingInvite field to given value.


### GetVerified

`func (o *Member) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *Member) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *Member) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetPendingEmail

`func (o *Member) GetPendingEmail() string`

GetPendingEmail returns the PendingEmail field if non-nil, zero value otherwise.

### GetPendingEmailOk

`func (o *Member) GetPendingEmailOk() (*string, bool)`

GetPendingEmailOk returns a tuple with the PendingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingEmail

`func (o *Member) SetPendingEmail(v string)`

SetPendingEmail sets PendingEmail field to given value.

### HasPendingEmail

`func (o *Member) HasPendingEmail() bool`

HasPendingEmail returns a boolean if a field has been set.

### GetCustomRoles

`func (o *Member) GetCustomRoles() []string`

GetCustomRoles returns the CustomRoles field if non-nil, zero value otherwise.

### GetCustomRolesOk

`func (o *Member) GetCustomRolesOk() (*[]string, bool)`

GetCustomRolesOk returns a tuple with the CustomRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoles

`func (o *Member) SetCustomRoles(v []string)`

SetCustomRoles sets CustomRoles field to given value.


### GetMfa

`func (o *Member) GetMfa() string`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *Member) GetMfaOk() (*string, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *Member) SetMfa(v string)`

SetMfa sets Mfa field to given value.


### GetExcludedDashboards

`func (o *Member) GetExcludedDashboards() []string`

GetExcludedDashboards returns the ExcludedDashboards field if non-nil, zero value otherwise.

### GetExcludedDashboardsOk

`func (o *Member) GetExcludedDashboardsOk() (*[]string, bool)`

GetExcludedDashboardsOk returns a tuple with the ExcludedDashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedDashboards

`func (o *Member) SetExcludedDashboards(v []string)`

SetExcludedDashboards sets ExcludedDashboards field to given value.

### HasExcludedDashboards

`func (o *Member) HasExcludedDashboards() bool`

HasExcludedDashboards returns a boolean if a field has been set.

### GetLastSeen

`func (o *Member) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *Member) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *Member) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.


### GetLastSeenMetadata

`func (o *Member) GetLastSeenMetadata() LastSeenMetadata`

GetLastSeenMetadata returns the LastSeenMetadata field if non-nil, zero value otherwise.

### GetLastSeenMetadataOk

`func (o *Member) GetLastSeenMetadataOk() (*LastSeenMetadata, bool)`

GetLastSeenMetadataOk returns a tuple with the LastSeenMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenMetadata

`func (o *Member) SetLastSeenMetadata(v LastSeenMetadata)`

SetLastSeenMetadata sets LastSeenMetadata field to given value.

### HasLastSeenMetadata

`func (o *Member) HasLastSeenMetadata() bool`

HasLastSeenMetadata returns a boolean if a field has been set.

### GetIntegrationMetadata

`func (o *Member) GetIntegrationMetadata() IntegrationMetadata`

GetIntegrationMetadata returns the IntegrationMetadata field if non-nil, zero value otherwise.

### GetIntegrationMetadataOk

`func (o *Member) GetIntegrationMetadataOk() (*IntegrationMetadata, bool)`

GetIntegrationMetadataOk returns a tuple with the IntegrationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationMetadata

`func (o *Member) SetIntegrationMetadata(v IntegrationMetadata)`

SetIntegrationMetadata sets IntegrationMetadata field to given value.

### HasIntegrationMetadata

`func (o *Member) HasIntegrationMetadata() bool`

HasIntegrationMetadata returns a boolean if a field has been set.

### GetTeams

`func (o *Member) GetTeams() []MemberTeamSummaryRep`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *Member) GetTeamsOk() (*[]MemberTeamSummaryRep, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *Member) SetTeams(v []MemberTeamSummaryRep)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *Member) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetPermissionGrants

`func (o *Member) GetPermissionGrants() []MemberPermissionGrantSummaryRep`

GetPermissionGrants returns the PermissionGrants field if non-nil, zero value otherwise.

### GetPermissionGrantsOk

`func (o *Member) GetPermissionGrantsOk() (*[]MemberPermissionGrantSummaryRep, bool)`

GetPermissionGrantsOk returns a tuple with the PermissionGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGrants

`func (o *Member) SetPermissionGrants(v []MemberPermissionGrantSummaryRep)`

SetPermissionGrants sets PermissionGrants field to given value.

### HasPermissionGrants

`func (o *Member) HasPermissionGrants() bool`

HasPermissionGrants returns a boolean if a field has been set.

### GetCreationDate

`func (o *Member) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Member) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Member) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


