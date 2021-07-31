# MemberRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]CoreLink**](CoreLink.md) |  | 
**Id** | **string** | The member&#39;s ID | 
**FirstName** | Pointer to **string** | The member&#39;s first name | [optional] 
**LastName** | Pointer to **string** | The member&#39;s last name | [optional] 
**Role** | **string** | The member&#39;s built-in role. If the member has no custom roles, this role will be in effect. | 
**Email** | **string** | The member&#39;s e-mail address | 
**PendingInvite** | **bool** | Whether or not the member has a pending invitation | 
**Verified** | **bool** | Whether or not the member&#39;s e-mail address has been verified | 
**PendingEmail** | Pointer to **string** |  | [optional] 
**CustomRoles** | **[]string** | The set of custom roles (as keys) assigned to the member | 
**Mfa** | **string** | Whether or not multi-factor authentication is enabled for this member | 
**ExcludedDashboards** | **[]string** | Default dashboards that the team member has chosen to ignore | 
**LastSeen** | **int64** |  | 
**LastSeenMetadata** | Pointer to [**LastSeenMetadata**](LastSeenMetadata.md) |  | [optional] 
**IntegrationMetadata** | Pointer to [**map[string]AccountsIntegrationSubscriptionMetadata**](AccountsIntegrationSubscriptionMetadata.md) |  | [optional] 
**Teams** | Pointer to [**[]MemberTeamSummaryRep**](MemberTeamSummaryRep.md) |  | [optional] 

## Methods

### NewMemberRep

`func NewMemberRep(links map[string]CoreLink, id string, role string, email string, pendingInvite bool, verified bool, customRoles []string, mfa string, excludedDashboards []string, lastSeen int64, ) *MemberRep`

NewMemberRep instantiates a new MemberRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberRepWithDefaults

`func NewMemberRepWithDefaults() *MemberRep`

NewMemberRepWithDefaults instantiates a new MemberRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *MemberRep) GetLinks() map[string]CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MemberRep) GetLinksOk() (*map[string]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MemberRep) SetLinks(v map[string]CoreLink)`

SetLinks sets Links field to given value.


### GetId

`func (o *MemberRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberRep) SetId(v string)`

SetId sets Id field to given value.


### GetFirstName

`func (o *MemberRep) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MemberRep) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MemberRep) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *MemberRep) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *MemberRep) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MemberRep) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MemberRep) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *MemberRep) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetRole

`func (o *MemberRep) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MemberRep) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MemberRep) SetRole(v string)`

SetRole sets Role field to given value.


### GetEmail

`func (o *MemberRep) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MemberRep) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MemberRep) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPendingInvite

`func (o *MemberRep) GetPendingInvite() bool`

GetPendingInvite returns the PendingInvite field if non-nil, zero value otherwise.

### GetPendingInviteOk

`func (o *MemberRep) GetPendingInviteOk() (*bool, bool)`

GetPendingInviteOk returns a tuple with the PendingInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingInvite

`func (o *MemberRep) SetPendingInvite(v bool)`

SetPendingInvite sets PendingInvite field to given value.


### GetVerified

`func (o *MemberRep) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *MemberRep) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *MemberRep) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetPendingEmail

`func (o *MemberRep) GetPendingEmail() string`

GetPendingEmail returns the PendingEmail field if non-nil, zero value otherwise.

### GetPendingEmailOk

`func (o *MemberRep) GetPendingEmailOk() (*string, bool)`

GetPendingEmailOk returns a tuple with the PendingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingEmail

`func (o *MemberRep) SetPendingEmail(v string)`

SetPendingEmail sets PendingEmail field to given value.

### HasPendingEmail

`func (o *MemberRep) HasPendingEmail() bool`

HasPendingEmail returns a boolean if a field has been set.

### GetCustomRoles

`func (o *MemberRep) GetCustomRoles() []string`

GetCustomRoles returns the CustomRoles field if non-nil, zero value otherwise.

### GetCustomRolesOk

`func (o *MemberRep) GetCustomRolesOk() (*[]string, bool)`

GetCustomRolesOk returns a tuple with the CustomRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoles

`func (o *MemberRep) SetCustomRoles(v []string)`

SetCustomRoles sets CustomRoles field to given value.


### GetMfa

`func (o *MemberRep) GetMfa() string`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *MemberRep) GetMfaOk() (*string, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *MemberRep) SetMfa(v string)`

SetMfa sets Mfa field to given value.


### GetExcludedDashboards

`func (o *MemberRep) GetExcludedDashboards() []string`

GetExcludedDashboards returns the ExcludedDashboards field if non-nil, zero value otherwise.

### GetExcludedDashboardsOk

`func (o *MemberRep) GetExcludedDashboardsOk() (*[]string, bool)`

GetExcludedDashboardsOk returns a tuple with the ExcludedDashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedDashboards

`func (o *MemberRep) SetExcludedDashboards(v []string)`

SetExcludedDashboards sets ExcludedDashboards field to given value.


### GetLastSeen

`func (o *MemberRep) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *MemberRep) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *MemberRep) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.


### GetLastSeenMetadata

`func (o *MemberRep) GetLastSeenMetadata() LastSeenMetadata`

GetLastSeenMetadata returns the LastSeenMetadata field if non-nil, zero value otherwise.

### GetLastSeenMetadataOk

`func (o *MemberRep) GetLastSeenMetadataOk() (*LastSeenMetadata, bool)`

GetLastSeenMetadataOk returns a tuple with the LastSeenMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenMetadata

`func (o *MemberRep) SetLastSeenMetadata(v LastSeenMetadata)`

SetLastSeenMetadata sets LastSeenMetadata field to given value.

### HasLastSeenMetadata

`func (o *MemberRep) HasLastSeenMetadata() bool`

HasLastSeenMetadata returns a boolean if a field has been set.

### GetIntegrationMetadata

`func (o *MemberRep) GetIntegrationMetadata() map[string]AccountsIntegrationSubscriptionMetadata`

GetIntegrationMetadata returns the IntegrationMetadata field if non-nil, zero value otherwise.

### GetIntegrationMetadataOk

`func (o *MemberRep) GetIntegrationMetadataOk() (*map[string]AccountsIntegrationSubscriptionMetadata, bool)`

GetIntegrationMetadataOk returns a tuple with the IntegrationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationMetadata

`func (o *MemberRep) SetIntegrationMetadata(v map[string]AccountsIntegrationSubscriptionMetadata)`

SetIntegrationMetadata sets IntegrationMetadata field to given value.

### HasIntegrationMetadata

`func (o *MemberRep) HasIntegrationMetadata() bool`

HasIntegrationMetadata returns a boolean if a field has been set.

### GetTeams

`func (o *MemberRep) GetTeams() []MemberTeamSummaryRep`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *MemberRep) GetTeamsOk() (*[]MemberTeamSummaryRep, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *MemberRep) SetTeams(v []MemberTeamSummaryRep)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *MemberRep) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


