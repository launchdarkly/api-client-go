This repository contains a client library for LaunchDarkly's REST API. This client was automatically
generated from our [OpenAPI specification](https://github.com/launchdarkly/ld-openapi).

This REST API is for custom integrations, data export, or automating your feature flag workflows. *DO NOT* use this client library to include feature flags in your web or mobile application. To integrate feature flags with your application, please see the [SDK documentation](https://docs.launchdarkly.com/v2.0/docs)

# Go API client for ldapi

# Authentication

All REST API resources are authenticated with [personal access tokens](https://docs.launchdarkly.com/home/account-security/api-access-tokens) or session cookies. Other authentication mechanisms are not supported. You can manage personal access tokens on your [Account settings](https://app.launchdarkly.com/settings/tokens) page. 

LaunchDarkly also has SDK keys, mobile keys, and client-side IDs that are used by our server-side SDKs, mobile SDKs, and client-side JavaScript SDKs, respectively. **These keys cannot be used to access our REST API**. These keys are environment-specific, and can only perform read-only operations (fetching feature flag settings).

| Auth mechanism | Allowed resources | Use cases |
| -------------- | ----------------- | --------- |
| [Personal access tokens](https://docs.launchdarkly.com/home/account-security/api-access-tokens) | Can be customized on a per-token basis | Building scripts, custom integrations, data export |
| SDK keys | Can only access read-only SDK-specific resources and the firehose, restricted to a single environment | Server-side SDKs, Firehose API |
| Mobile keys | Can only access read-only mobile SDK-specific resources, restricted to a single environment | Mobile SDKs |
| Client-side ID | Single environment, only flags marked available to client-side | Client-side JavaScript |

<blockquote>
    <h3><span>❗️</span>Keep your access tokens and SDK keys private</h3>
    <p>Access tokens should *never* be exposed in untrusted contexts. Never put an access token in client-side JavaScript, or embed it in a mobile application. LaunchDarkly has special mobile keys that you can embed in mobile apps. If you accidentally expose an access token or SDK key, you can reset it from your [Account Settings](https://app.launchdarkly.com/settings#/tokens) page.</p>
    <p>The client-side ID is safe to embed in untrusted contexts. It's designed for use in client-side JavaScript.</p>
</blockquote>

## Via request header

The preferred way to authenticate with the API is by adding an `Authorization` header containing your access token to your requests. The value of the `Authorization` header must be your access token.

Manage personal access tokens from the [Account Settings](https://app.launchdarkly.com/settings/tokens) page.

## Via session cookie

For testing purposes, you can make API calls directly from your web browser. If you're logged in to the application, the API will use your existing session to authenticate calls.

If you have a [role](https://docs.launchdarkly.com/home/team/built-in-roles) other than Admin, or have a [custom role](https://docs.launchdarkly.com/home/team/custom-roles) defined, you may not have permission to perform some API calls. You will receive a `401` response code in that case.

<blockquote>
    <h3><span>❗️</span>Modifying the Origin header causes an error</h3>
    <p>We validate that the Origin header for any API request authenticated by a session cookie matches the expected Origin header. The expected Origin header is `https://app.launchdarkly.com`.</p>
    <p>If the Origin header does not match what's expected, LaunchDarkly returns an error. This error can prevent the LaunchDarkly app from working correctly. </p>
    <p>Any browser extension that intentionally changes the Origin header can cause this problem. For example, the `Allow-Control-Allow-Origin: *` Chrome extension changes the Origin header to http://evil.com and causes the app to fail.</p>
    <p>To prevent this error, do not modify your Origin header.</p>
    <p>LaunchDarkly does not require origin matching when authenticating with an Access Token, so this issue does not affect normal API usage.</p>
</blockquote>

# Representations

All resources expect and return JSON response bodies. Error responses will also send a JSON body. Read [Errors](#section/Errors) for a more detailed description of the error format used by the API. 

In practice this means that you always get a response with a `Content-Type` header set to `application/json`.

In addition, request bodies for `PUT`, `POST`, `REPORT` and `PATCH` requests must be encoded as JSON with a `Content-Type` header set to `application/json`.

## Summary and detailed representations

When you fetch a list of resources, the response includes only the most important attributes of each resource. This is a *summary representation* of the resource. When you fetch an individual resource (for example, a single feature flag), you receive a *detailed representation* containing all of the attributes of the resource.

The best way to find a detailed representation is to follow links. Every summary representation includes a link to its detailed representation.

## Links and addressability

The best way to navigate the API is by following links. These are attributes in representations that link to other resources. The API always uses the same format for links:

* Links to other resources within the API are encapsulated in a `_links` object.
* If the resource has a corresponding link to HTML content on the site, it is stored in a special `_site` link.

Each link has two attributes: an href (the URL) and a type (the content type). For example, a feature resource might return the following:

```json
{
    \"_links\": {
        \"parent\": {
            \"href\": \"/api/features\",
            \"type\": \"application/json\"
        },
        \"self\": {
            \"href\":\"/api/features/sort.order\",
            \"type\":\"application/json\"
        }
    },
    \"_site\":{
        \"href\":\"/features/sort.order\",
        \"type\":\"text/html\"
    }
}
```

From this, you can navigate to the parent collection of features by following the `parent` link, or navigate to the site page for the feature by following the `_site` link.

Collections are always represented as a JSON object with an `items` attribute containing an array of representations. Like all other representations, collections have `_links` defined at the top level.

Paginated collections include `first`, `last`, `next`, and `prev` links containing a URL with the respective set of elements in the collection.

# Updates

Resources that accept partial updates use the `PATCH` verb, and support the [JSON Patch](http://tools.ietf.org/html/rfc6902) format. Some resources also support the [JSON Merge Patch](https://tools.ietf.org/html/rfc7386) format. In addition, some resources support optional comments that can be submitted with updates. Comments appear in outgoing webhooks, the audit log, and other integrations.

## Updates via JSON Patch

[JSON Patch](http://tools.ietf.org/html/rfc6902) is a way to specify the modifications to perform on a resource. For example, in this feature flag representation:

```json
{
    \"name\": \"New recommendations engine\",
    \"key\": \"engine.enable\",
    \"description\": \"This is the description\",
    ...
}
```

You can change the feature flag's description with the following patch document:

```json
[
    { \"op\": \"replace\", \"path\": \"/description\", \"value\": \"This is the new description\"}
]
```

JSON Patch documents are always arrays. You can specify multiple modifications to perform in a single request. You can also test that certain preconditions are met before applying the patch:

```json
[
    { \"op\": \"test\", \"path\": \"/version\", \"value\": 10 },
    { \"op\": \"replace\", \"path\": \"/description\", \"value\": \"The new description\" }
]
```

The above patch request tests whether the feature flag's `version` is `10`, and if so, changes the feature flag's description.

Attributes that aren't editable, like a resource's `_links`, have names that start with an underscore.

## Updates via JSON Merge Patch

The API also supports the [JSON Merge Patch](https://tools.ietf.org/html/rfc7386) format, as well as the [Update feature flag](#operation/patchFeatureFlag) resource. 

JSON Merge Patch is less expressive than JSON Patch but in many cases, it is simpler to construct a merge patch document. For example, you can change a feature flag's description with the following merge patch document:

```json
{
    \"description\": \"New flag description\"
}
```

## Updates with comments

You can submit optional comments with `PATCH` changes. The [Update feature flag](#operation/patchFeatureFlag) resource supports comments.

To submit a comment along with a JSON Patch document, use the following format:

```json
{
    \"comment\": \"This is a comment string\",
    \"patch\": [ {\"op\": \"replace\", \"path\": \"/description\", \"value\": \"The new description\" } ]
}
```

To submit a comment along with a JSON Merge Patch document, use the following format:

```json
{
    \"comment\": \"This is a comment string\",
    \"merge\": { \"description\": \"New flag description\"}
}
```

## Updates via semantic patches

The API also supports the Semantic patch format. A semantic `PATCH` is a way to specify the modifications to perform on a resource as a set of executable instructions. 

JSON Patch uses paths and a limited set of operations to describe how to transform the current state of the resource into a new state. Semantic patch allows you to be explicit about intent using precise, custom instructions. In many cases, semantic patch instructions can also be defined independently of the current state of the resource. This can be useful when defining a change that may be applied at a future date.

For example, in this feature flag configuration in environment Production:

```json
{
    \"name\": \"Alternate sort order\",
    \"kind\": \"boolean\",
    \"key\": \"sort.order\",
   ...
    \"environments\": {
        \"production\": {
            \"on\": true,
            \"archived\": false,
            \"salt\": \"c29ydC5vcmRlcg==\",
            \"sel\": \"8de1085cb7354b0ab41c0e778376dfd3\",
            \"lastModified\": 1469131558260,
            \"version\": 81,
            \"targets\": [
                {
                    \"values\": [
                        \"Gerhard.Little@yahoo.com\"
                    ],
                    \"variation\": 0
                },
                {
                    \"values\": [
                        \"1461797806429-33-861961230\",
                        \"438580d8-02ee-418d-9eec-0085cab2bdf0\"
                    ],
                    \"variation\": 1
                }
            ],
            \"rules\": [],
            \"fallthrough\": {
                \"variation\": 0
            },
            \"offVariation\": 1,
            \"prerequisites\": [],
            \"_site\": {
                \"href\": \"/default/production/features/sort.order\",
                \"type\": \"text/html\"
            }
       }
    }
}
```

You can add a date you want a user to be removed from the feature flag's user targets. For example, “remove user 1461797806429-33-861961230 from the user target for variation 0 on the Alternate sort order flag in the production environment on Wed Jul 08 2020 at 15:27:41 pm”. This is done using the following:

```json
{
    \"comment\": \"update expiring user targets\",
    \"instructions\": [
        {
            \"kind\": \"removeExpireUserTargetDate\",
            \"userKey\": \"userKey\",
            \"variationId\": \"978d53f9-7fe3-4a63-992d-97bcb4535dc8\",
        },
        {
            \"kind\": \"updateExpireUserTargetDate\",
            \"userKey\": \"userKey2\",
            \"variationId\": \"978d53f9-7fe3-4a63-992d-97bcb4535dc8\",
            \"value\": 1587582000000
        },
        {
            \"kind\": \"addExpireUserTargetDate\",
            \"userKey\": \"userKey3\",
            \"variationId\": \"978d53f9-7fe3-4a63-992d-97bcb4535dc8\",
            \"value\": 1594247266386
        }
    ]
}
```

Here is another example. In this feature flag configuration:

```json
{
    \"name\": \"New recommendations engine\",
    \"key\": \"engine.enable\",
    \"environments\": {
        \"test\": {
            \"on\": true
        }
    }
}
```

You can change the feature flag's description with the following patch document as a set of executable instructions. For example, “add user X to targets for variation Y and remove user A from targets for variation B for test flag”:

```json
{
    \"comment\": \"\",
    \"instructions\": [
        {
            \"kind\": \"removeUserTargets\",
            \"values\": [\"438580d8-02ee-418d-9eec-0085cab2bdf0\"],
            \"variationId\": \"852cb784-54ff-46b9-8c35-5498d2e4f270\"
        },
        {
            \"kind\": \"addUserTargets\",
            \"values\": [\"438580d8-02ee-418d-9eec-0085cab2bdf0\"],
            \"variationId\": \"1bb18465-33b6-49aa-a3bd-eeb6650b33ad\"
        }
    ]
}
```

<blockquote>
    <h3><span>📘</span>Supported semantic patch API endpoints</h3>
    <p>TODO: update these links</p>
    <p><a href=\"#operation/patchFeatureFlag\">Update feature flag</a></p>
    <p><a href=\"\">Update expiring user targets on feature flag</a></p>
    <p><a href=\"\">Update expiring user target for flags</a></p>
</blockquote>

# Errors

The API always returns errors in a common format. Here's an example:

```json
{
    \"code\": \"invalid_request\",
    \"message\": \"A feature with that key already exists\",
    \"id\": \"30ce6058-87da-11e4-b116-123b93f75cba\"
}
```

The general class of error is indicated by the `code`. The `message` is a human-readable explanation of what went wrong. The `id` is a unique identifier. Use it when you're working with LaunchDarkly support to debug a problem with a specific API call.

## HTTP Status - Error Response Codes

| Code | Definition | Desc. | Possible Solution |
| ---- | ---------- | ----- | ----------------- |
| 400  | Bad Request | A request that fails may return this HTTP response code. | Ensure JSON syntax in request body is correct. |
| 401  | Unauthorized | User doesn't have permission to an API call. | Ensure your SDK key is good. |
| 403  | Forbidden | User does not have permission for operation. | Ensure that the user or access token has proper permissions set. |
| 409  | Conflict | The API request could not be completed because it conflicted with a concurrent API request. | Retry your request. |
| 429  | Too many requests | See [Rate limiting](ref:rate-limiting). | Wait and try again later. |

# CORS

The LaunchDarkly API supports Cross Origin Resource Sharing (CORS) for AJAX requests from any origin. If an `Origin` header is given in a request, it will be echoed as an explicitly allowed origin. Otherwise, a wildcard is returned: `Access-Control-Allow-Origin: *`. For more information on CORS, see the [CORS W3C Recommendation](http://www.w3.org/TR/cors). Example CORS headers might look like:

```http
Access-Control-Allow-Headers: Accept, Content-Type, Content-Length, Accept-Encoding, Authorization
Access-Control-Allow-Methods: OPTIONS, GET, DELETE, PATCH
Access-Control-Allow-Origin: *
Access-Control-Max-Age: 300
```

You can make authenticated CORS calls just as you would make same-origin calls, using either [token or session-based authentication](#section/Authentication). If you’re using session auth, you should set the `withCredentials` property for your `xhr` request to `true`. You should never expose your access tokens to untrusted users.

# Rate limiting

We use several rate limiting strategies to ensure the availability of our APIs. Rate-limited calls to our APIs will return a `429` status code. Calls to our APIs will include headers indicating the current rate limit status. The specific headers returned depend on the API route being called-- the limits differ based on the route, authentication mechanism, and other factors. Routes that are not rate limited may not contain any of the headers described below.

<blockquote>
    <h3><span>❗️</span>Rate limiting and SDKs</h3>
    <p>Our SDKs are never rate limited. Our SDKs do not use the API endpoints defined here. We use a different set of approaches, including streaming/server-sent events and a global CDN, to ensure availability to the routes used by our SDKs.</p>
    <p>The client-side ID is safe to embed in untrusted contexts. It's designed for use in client-side JavaScript.</p>
</blockquote>

## Global rate limits

Authenticated requests are subject to a global limit. This is the maximum number of calls that can be made to the API per ten seconds. All personal access tokens on the account share this limit, so exceeding the limit with one access token will impact other tokens. Calls that are subject to global rate limits will return the headers below:

| Header name | Description |
| ----------- | ----------- |
| `X-Ratelimit-Global-Remaining` | The maximum number of requests the account is permitted to make per ten seconds. |
| `X-Ratelimit-Reset` | The time at which the current rate limit window resets in epoch milliseconds. |

We do not publicly document the specific number of calls that can be made globally. This limit may change, and we encourage clients to program against the specification, relying on the two headers defined above, rather than hardcoding to the current limit.

## Route-level rate limits

Some authenticated routes have custom rate limits. These also reset every ten seconds. Any access tokens hitting the same route share this limit, so exceeding the limit with one access token may impact other tokens. Calls that are subject to route-level rate limits will return the headers below:

| Header name | Description |
| ----------- | ----------- |
| `X-Ratelimit-Route-Remaining` | The maximum number of requests to the current route the account is permitted to make per ten seconds. |
| `X-Ratelimit-Reset` | The time at which the current rate limit window resets in epoch milliseconds. |

A *route* represents a specific URL pattern and verb. For example, the [Delete environment](#operation/deleteEnvironment) endpoint is considered a single route, and each call to delete an environment counts against your route-level rate limit for that route. 

We do not publicly document the specific number of calls that can be made to each endpoint per ten seconds. These limits may change, and we encourage clients to program against the specification, relying on the two headers defined above, rather than hardcoding to the current limits.

## IP-based rate limiting

We also employ IP-based rate limiting on some API routes. If you hit an IP-based rate limit, your API response will include a `Retry-After` header indicating how long to wait before re-trying the call. Clients must wait at least `Retry-After` seconds before making additional calls to our API, and should employ jitter and backoff strategies to avoid triggering rate limits again.

# OpenAPI (Swagger)

We have a [complete OpenAPI (Swagger) specification](https://app.launchdarkly.com/api/v2/openapi.json) for our API.

You can use this specification to generate client libraries to interact with our REST API in your language of choice. 

# Client libraries

We auto-generate multiple client libraries based on our OpenAPI specification. To learn more, visit [GitHub](https://github.com/search?q=topic%3Alaunchdarkly-api+org%3Alaunchdarkly&type=Repositories).

# Method Overriding

Some firewalls and HTTP clients restrict the use of verbs other than `GET` and `POST`. In those environments, our API endpoints that use `PUT`, `PATCH`, and `DELETE` verbs will be inaccessible.

To avoid this issue, our API supports the `X-HTTP-Method-Override` header, allowing clients to \"tunnel\" `PUT`, `PATCH`, and `DELETE` requests via a `POST` request. 

For example, if you wish to call one of our `PATCH` resources via a `POST` request, you can include `X-HTTP-Method-Override:PATCH` as a header.

# Beta resources

We sometimes release new API resources in **beta** status before we release them with general availability. 

Resources that are in beta are still undergoing testing and development. They may change without notice, including becoming backwards incompatible. 

We try to promote resources into general availability as quickly as possible. This happens after sufficient testing and when we're satisfied that we no longer need to make backwards-incompatible changes.

We mark beta resources with a \"Beta\" callout in our documentation, pictured below:
<blockquote>
    <h3><span>📘</span>Beta</h3>
    <p>**This feature is in beta.** You must include a specific header to use it.\\n\\nTo learn more, read [Beta resources](ref:beta-resources).</p>
</blockquote>

## Using beta resources

To use a beta resource, you must include a header in the request. If you call a beta resource without this header, you'll receive a `403` response.

Use this header: 

```
LD-API-Version: beta
```

# Versioning

We try hard to keep our REST API backwards compatible, but we occasionally have to make backwards-incompatible changes in the process of shipping new features. These breaking changes can cause unexpected behavior if you don't prepare for them accordingly. 

Updates to our REST API include support for the latest features in LaunchDarkly. We also release a new version of our REST API every time we make a breaking change. We provide simultaneous support for multiple API versions so you can migrate from your current API version to a new version at your own pace. 

See new versions in the [Changelog](ref:changelog).

## Setting the API version per request

You can set the API version on a specific request by sending an `LD-API-Version` header, as shown in the example below:

```
LD-API-Version: 20191212
```

The header value is the version number of the API version you'd like to request. The number for each version corresponds to the date the version was released. In the example above the version `20191212` corresponds to December 12, 2019. 

## Setting the API version per access token

When creating an access token, you must specify a specific version of the API to use. This ensures that integrations using this token cannot be broken by version changes.

Tokens created before versioning was released have their version set to `20160426` (the version of the API that existed before versioning) so that they continue working the same way they did before versioning.

If you would like to upgrade your integration to use a new API version, you can explicitly set the header described above.

<blockquote>
    <h3><span>👍</span>Best practice: Set the header for every client or integration</h3>
    <p>We recommend that you set the API version header explicitly in any client or integration you build.</p>
    <p>Only rely on the access token API version during manual testing.</p>
</blockquote>

<blockquote>
    <h3><span>🚧</span>API Path Versioning</h3>
    <p>In the past, we've used path-based API versioning. For example, versioning resources by adding `v2` to endpoint URLs. We don't foresee the need to do this again, but may do so if we need to make major revisions to the API.</p>
</blockquote>


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://support.launchdarkly.com](https://support.launchdarkly.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./ldapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://app.launchdarkly.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccessTokensApi* | [**DeleteToken**](docs/AccessTokensApi.md#deletetoken) | **Delete** /api/v2/tokens/{id} | Delete access token
*AccessTokensApi* | [**GetToken**](docs/AccessTokensApi.md#gettoken) | **Get** /api/v2/tokens/{id} | Get access token
*AccessTokensApi* | [**GetTokens**](docs/AccessTokensApi.md#gettokens) | **Get** /api/v2/tokens | List access tokens
*AccessTokensApi* | [**PatchToken**](docs/AccessTokensApi.md#patchtoken) | **Patch** /api/v2/tokens/{id} | Patch access token
*AccessTokensApi* | [**PostToken**](docs/AccessTokensApi.md#posttoken) | **Post** /api/v2/tokens | Create access token
*AccessTokensApi* | [**ResetToken**](docs/AccessTokensApi.md#resettoken) | **Post** /api/v2/tokens/{id}/reset | Reset access token
*ApprovalsApi* | [**DeleteApprovalRequest**](docs/ApprovalsApi.md#deleteapprovalrequest) | **Delete** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{id} | Delete approval request
*ApprovalsApi* | [**GetApproval**](docs/ApprovalsApi.md#getapproval) | **Get** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{id} | Get approval request
*ApprovalsApi* | [**GetApprovals**](docs/ApprovalsApi.md#getapprovals) | **Get** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests | Get all approval requests
*ApprovalsApi* | [**PostApprovalRequest**](docs/ApprovalsApi.md#postapprovalrequest) | **Post** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests | Create approval request
*ApprovalsApi* | [**PostApprovalRequestApplyRequest**](docs/ApprovalsApi.md#postapprovalrequestapplyrequest) | **Post** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{id}/apply | Apply approval request
*ApprovalsApi* | [**PostApprovalRequestReview**](docs/ApprovalsApi.md#postapprovalrequestreview) | **Post** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{id}/reviews | Review approval request
*ApprovalsApi* | [**PostFlagCopyConfigApprovalRequest**](docs/ApprovalsApi.md#postflagcopyconfigapprovalrequest) | **Post** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/copy/environments/{environmentKey}/approval-requests-flag-copy | Create approval request to copy flag configurations across environments
*AuditLogApi* | [**GetAuditLogEntries**](docs/AuditLogApi.md#getauditlogentries) | **Get** /api/v2/auditlog | List audit log feature flag entries
*CodeReferencesApi* | [**DeleteRepository**](docs/CodeReferencesApi.md#deleterepository) | **Delete** /api/v2/code-refs/repositories/{repo} | Delete repository
*CodeReferencesApi* | [**GetBranch**](docs/CodeReferencesApi.md#getbranch) | **Get** /api/v2/code-refs/repositories/{repo}/branches/{branch} | Get branch
*CodeReferencesApi* | [**GetBranches**](docs/CodeReferencesApi.md#getbranches) | **Get** /api/v2/code-refs/repositories/{repo}/branches | List branches
*CodeReferencesApi* | [**GetRepositories**](docs/CodeReferencesApi.md#getrepositories) | **Get** /api/v2/code-refs/repositories | List repositories
*CodeReferencesApi* | [**GetRepository**](docs/CodeReferencesApi.md#getrepository) | **Get** /api/v2/code-refs/repositories/{repo} | Get repository
*CodeReferencesApi* | [**GetStatistics**](docs/CodeReferencesApi.md#getstatistics) | **Get** /api/v2/code-refs/statistics/{projKey} | Get number of code references for flags
*CodeReferencesApi* | [**PatchRepository**](docs/CodeReferencesApi.md#patchrepository) | **Patch** /api/v2/code-refs/repositories/{repo} | Update repository
*CodeReferencesApi* | [**PostRepository**](docs/CodeReferencesApi.md#postrepository) | **Post** /api/v2/code-refs/repositories | Create repository
*CodeReferencesApi* | [**PutBranch**](docs/CodeReferencesApi.md#putbranch) | **Put** /api/v2/code-refs/repositories/{repo}/branches/{branch} | Upsert branch
*CustomRolesApi* | [**DeleteCustomRole**](docs/CustomRolesApi.md#deletecustomrole) | **Delete** /api/v2/roles/{key} | Delete custom role
*CustomRolesApi* | [**GetCustomRole**](docs/CustomRolesApi.md#getcustomrole) | **Get** /api/v2/roles/{key} | Get custom role
*CustomRolesApi* | [**GetCustomRoles**](docs/CustomRolesApi.md#getcustomroles) | **Get** /api/v2/roles | List custom roles
*CustomRolesApi* | [**PatchCustomRole**](docs/CustomRolesApi.md#patchcustomrole) | **Patch** /api/v2/roles/{key} | Update custom role
*CustomRolesApi* | [**PostCustomRole**](docs/CustomRolesApi.md#postcustomrole) | **Post** /api/v2/roles | Create custom role
*DataExportDestinationsApi* | [**DeleteDestination**](docs/DataExportDestinationsApi.md#deletedestination) | **Delete** /api/v2/destinations/{projKey}/{envKey}/{id} | Delete Data Export destination
*DataExportDestinationsApi* | [**GetDestination**](docs/DataExportDestinationsApi.md#getdestination) | **Get** /api/v2/destinations/{projKey}/{envKey}/{id} | Get destination
*DataExportDestinationsApi* | [**GetDestinations**](docs/DataExportDestinationsApi.md#getdestinations) | **Get** /api/v2/destinations | List destinations
*DataExportDestinationsApi* | [**PatchDestination**](docs/DataExportDestinationsApi.md#patchdestination) | **Patch** /api/v2/destinations/{projKey}/{envKey}/{id} | Update Data Export destination
*DataExportDestinationsApi* | [**PostDestination**](docs/DataExportDestinationsApi.md#postdestination) | **Post** /api/v2/destinations/{projKey}/{envKey} | Create data export destination
*EnvironmentsApi* | [**DeleteEnvironment**](docs/EnvironmentsApi.md#deleteenvironment) | **Delete** /api/v2/projects/{projectKey}/environments/{environmentKey} | Delete environment by key
*EnvironmentsApi* | [**GetEnvironment**](docs/EnvironmentsApi.md#getenvironment) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey} | Get environment by key
*EnvironmentsApi* | [**PatchEnvironment**](docs/EnvironmentsApi.md#patchenvironment) | **Patch** /api/v2/projects/{projectKey}/environments/{environmentKey} | Patch environment by key
*EnvironmentsApi* | [**PostEnvironment**](docs/EnvironmentsApi.md#postenvironment) | **Post** /api/v2/projects/{projectKey}/environments | Create environment
*ExperimentsApi* | [**GetExperiment**](docs/ExperimentsApi.md#getexperiment) | **Get** /api/v2/projects/{projKey}/{flagKey}/experiments/{envKey}/{metricKey} | Get experiment results
*ExperimentsApi* | [**PostExperiment**](docs/ExperimentsApi.md#postexperiment) | **Post** /api/v2/projects/{projKey}/experiments | Create experiment
*ExperimentsApi* | [**ResetExperiment**](docs/ExperimentsApi.md#resetexperiment) | **Delete** /api/v2/flags/{projKey}/{flagKey}/experiments/{envKey}/{metricKey}/results | Reset experiment results
*FeatureFlagsApi* | [**CopyFeatureFlag**](docs/FeatureFlagsApi.md#copyfeatureflag) | **Post** /api/v2/flags/{projKey}/{featureFlagKey}/copy | Copy feature flag
*FeatureFlagsApi* | [**DeleteFeatureFlag**](docs/FeatureFlagsApi.md#deletefeatureflag) | **Delete** /api/v2/flags/{projKey}/{key} | Delete flag by key
*FeatureFlagsApi* | [**GetDependentFlags**](docs/FeatureFlagsApi.md#getdependentflags) | **Get** /api/v2/flags/{projKey}/{flagKey}/dependent-flags | List dependent feature flags
*FeatureFlagsApi* | [**GetDependentFlagsByEnv**](docs/FeatureFlagsApi.md#getdependentflagsbyenv) | **Get** /api/v2/flags/{projKey}/{envKey}/{flagKey}/dependent-flags | List dependent feature flags by environment
*FeatureFlagsApi* | [**GetFeatureFlag**](docs/FeatureFlagsApi.md#getfeatureflag) | **Get** /api/v2/flags/{projKey}/{key} | Get feature flag
*FeatureFlagsApi* | [**GetFeatureFlagStatus**](docs/FeatureFlagsApi.md#getfeatureflagstatus) | **Get** /api/v2/flag-statuses/{projKey}/{envKey}/{key} | List feature flag statuses
*FeatureFlagsApi* | [**GetFeatureFlagStatusAcrossEnvironments**](docs/FeatureFlagsApi.md#getfeatureflagstatusacrossenvironments) | **Get** /api/v2/flag-status/{projKey}/{key} | Get feature flag status
*FeatureFlagsApi* | [**GetFeatureFlagStatuses**](docs/FeatureFlagsApi.md#getfeatureflagstatuses) | **Get** /api/v2/flag-statuses/{projKey}/{envKey} | List feature flag statuses
*FeatureFlagsApi* | [**GetFeatureFlags**](docs/FeatureFlagsApi.md#getfeatureflags) | **Get** /api/v2/flags/{projKey} | List feature flags
*FeatureFlagsApi* | [**PatchFeatureFlag**](docs/FeatureFlagsApi.md#patchfeatureflag) | **Patch** /api/v2/flags/{projKey}/{key} | Update feature flag
*FeatureFlagsApi* | [**PostFeatureFlag**](docs/FeatureFlagsApi.md#postfeatureflag) | **Post** /api/v2/flags/{projKey} | Create a feature flag
*IntegrationsApi* | [**CreateSubscription**](docs/IntegrationsApi.md#createsubscription) | **Post** /api/v2/integrations/{integrationKey} | Create integration configuration
*IntegrationsApi* | [**DeleteSubscription**](docs/IntegrationsApi.md#deletesubscription) | **Delete** /api/v2/integrations/{integrationKey}/{id} | Delete integration configuration
*IntegrationsApi* | [**GetConfigurations**](docs/IntegrationsApi.md#getconfigurations) | **Get** /api/v2/integrations | List integration configurations
*IntegrationsApi* | [**GetIntegrationSubscriptionByID**](docs/IntegrationsApi.md#getintegrationsubscriptionbyid) | **Get** /api/v2/integrations/{integrationKey}/{id} | Get configured integration by ID
*IntegrationsApi* | [**UpdateSubscription**](docs/IntegrationsApi.md#updatesubscription) | **Patch** /api/v2/integrations/{integrationKey} | Update subscription
*IntegrationsBetaApi* | [**ValidateConnection**](docs/IntegrationsBetaApi.md#validateconnection) | **Post** /integrations/{integrationKey}/{id}/validate | Validate connection
*MetricsApi* | [**DeleteMetric**](docs/MetricsApi.md#deletemetric) | **Delete** /api/v2/metrics/{projectKey}/{key} | Delete metric
*MetricsApi* | [**GetMetric**](docs/MetricsApi.md#getmetric) | **Get** /api/v2/metrics/{projectKey}/{key} | Get metric
*MetricsApi* | [**GetMetrics**](docs/MetricsApi.md#getmetrics) | **Get** /api/v2/metrics/{projectKey} | List metrics
*MetricsApi* | [**PatchMetric**](docs/MetricsApi.md#patchmetric) | **Patch** /api/v2/metrics/{projectKey}/{key} | Update metric
*MetricsApi* | [**PostMetric**](docs/MetricsApi.md#postmetric) | **Post** /api/v2/metrics/{projectKey} | Create metric
*OtherApi* | [**GetIps**](docs/OtherApi.md#getips) | **Get** /api/v2/public-ip-list | Gets the public IP list
*OtherApi* | [**GetOpenapiSpec**](docs/OtherApi.md#getopenapispec) | **Get** /api/v2/openapi.json | Gets the OpenAPI spec in json
*OtherApi* | [**GetRoot**](docs/OtherApi.md#getroot) | **Get** /api/v2 | Root resource
*OtherApi* | [**GetVersions**](docs/OtherApi.md#getversions) | **Get** /api/v2/versions | Get version information
*ProjectsApi* | [**DeleteProject**](docs/ProjectsApi.md#deleteproject) | **Delete** /api/v2/projects/{projectKey} | Delete project by key
*ProjectsApi* | [**GetProject**](docs/ProjectsApi.md#getproject) | **Get** /api/v2/projects/{projectKey} | Get project by key
*ProjectsApi* | [**GetProjects**](docs/ProjectsApi.md#getprojects) | **Get** /api/v2/projects | Get all projects
*ProjectsApi* | [**PatchProject**](docs/ProjectsApi.md#patchproject) | **Patch** /api/v2/projects/{projectKey} | Patch project by key
*ProjectsApi* | [**PostProject**](docs/ProjectsApi.md#postproject) | **Post** /api/v2/projects | Create project
*ScheduledChangesApi* | [**DeleteFlagConfigScheduledChanges**](docs/ScheduledChangesApi.md#deleteflagconfigscheduledchanges) | **Delete** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes/{id} | Delete scheduled changes workflow
*ScheduledChangesApi* | [**GetFeatureFlagScheduledChange**](docs/ScheduledChangesApi.md#getfeatureflagscheduledchange) | **Get** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes/{id} | Get a scheduled change
*ScheduledChangesApi* | [**GetFlagConfigScheduledChanges**](docs/ScheduledChangesApi.md#getflagconfigscheduledchanges) | **Get** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes | List scheduled changes
*ScheduledChangesApi* | [**GetFlagConfigScheduledChangesConflicts**](docs/ScheduledChangesApi.md#getflagconfigscheduledchangesconflicts) | **Post** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes-conflicts | List conflicts for proposed instructions
*ScheduledChangesApi* | [**PatchFlagConfigScheduledChange**](docs/ScheduledChangesApi.md#patchflagconfigscheduledchange) | **Patch** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes/{id} | Update scheduled changes workflow
*ScheduledChangesApi* | [**PostFlagConfigScheduledChanges**](docs/ScheduledChangesApi.md#postflagconfigscheduledchanges) | **Post** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes | Create scheduled changes workflow
*SegmentsApi* | [**DeleteSegment**](docs/SegmentsApi.md#deletesegment) | **Delete** /api/v2/segments/{projKey}/{envKey}/{key} | Delete segment
*SegmentsApi* | [**GetBigSegmentTarget**](docs/SegmentsApi.md#getbigsegmenttarget) | **Get** /segments/{projKey}/{envKey}/{key}/users/{userKey} | Get user in big segment
*SegmentsApi* | [**GetExpiringUserTargetsOnSegment**](docs/SegmentsApi.md#getexpiringusertargetsonsegment) | **Get** /api/v2/segments/{projKey}/{segmentKey}/expiring-user-targets/{envKey} | Get expiring user targets for segment
*SegmentsApi* | [**GetSegment**](docs/SegmentsApi.md#getsegment) | **Get** /api/v2/segments/{projKey}/{envKey}/{key} | Get segment
*SegmentsApi* | [**GetSegments**](docs/SegmentsApi.md#getsegments) | **Get** /api/v2/segments/{projKey}/{envKey} | List segments
*SegmentsApi* | [**PatchExpiringUserTargetsOnSegment**](docs/SegmentsApi.md#patchexpiringusertargetsonsegment) | **Patch** /api/v2/segments/{projKey}/{segmentKey}/expiring-user-targets/{envKey} | Update expiring user targets on segment
*SegmentsApi* | [**PatchSegment**](docs/SegmentsApi.md#patchsegment) | **Patch** /api/v2/segments/{projKey}/{envKey}/{key} | Patch segment
*SegmentsApi* | [**PostSegment**](docs/SegmentsApi.md#postsegment) | **Post** /api/v2/segments/{projKey}/{envKey} | Create segment
*SegmentsApi* | [**UpdateBigSegmentTargets**](docs/SegmentsApi.md#updatebigsegmenttargets) | **Post** /segments/{projKey}/{envKey}/{key}/users | Update targets on a big segment
*TeamMembersApi* | [**DeleteMember**](docs/TeamMembersApi.md#deletemember) | **Delete** /api/v2/members/{id} | Delete team member by ID
*TeamMembersApi* | [**GetMember**](docs/TeamMembersApi.md#getmember) | **Get** /api/v2/members/{id} | Get team member by ID
*TeamMembersApi* | [**PatchMember**](docs/TeamMembersApi.md#patchmember) | **Patch** /api/v2/members/{id} | Modify a team member by ID
*TeamMembersApi* | [**PostMembers**](docs/TeamMembersApi.md#postmembers) | **Post** /api/v2/members | Invite new members
*UserSettingsApi* | [**GetUserFlagSetting**](docs/UserSettingsApi.md#getuserflagsetting) | **Get** /api/v2/users/{projKey}/{envKey}/{key}/flags/{featureKey} | Get flag setting for user
*UserSettingsApi* | [**GetUserFlagSettings**](docs/UserSettingsApi.md#getuserflagsettings) | **Get** /api/v2/users/{projKey}/{envKey}/{key}/flags | List flag settings for user
*UserSettingsApi* | [**PutFlagSetting**](docs/UserSettingsApi.md#putflagsetting) | **Put** /api/v2/users/{projKey}/{envKey}/{key}/flags/{featureKey} | Update flag settings for user
*UsersApi* | [**DeleteUser**](docs/UsersApi.md#deleteuser) | **Delete** /api/v2/users/{projKey}/{envKey}/{key} | Delete user
*UsersApi* | [**GetSearchUsers**](docs/UsersApi.md#getsearchusers) | **Get** /api/v2/user-search/{projKey}/{envKey} | Find users
*UsersApi* | [**GetUser**](docs/UsersApi.md#getuser) | **Get** /api/v2/users/{projKey}/{envKey}/{key} | Get user
*UsersApi* | [**GetUsers**](docs/UsersApi.md#getusers) | **Get** /api/v2/users/{projKey}/{envKey} | List users


## Documentation For Models

 - [AccessDeniedReasonRep](docs/AccessDeniedReasonRep.md)
 - [AccessDeniedRep](docs/AccessDeniedRep.md)
 - [AccessRep](docs/AccessRep.md)
 - [AccessRepDenied](docs/AccessRepDenied.md)
 - [AccessTokenCollectionRep](docs/AccessTokenCollectionRep.md)
 - [AccessTokenRep](docs/AccessTokenRep.md)
 - [AccessTokensAccessTokenPost](docs/AccessTokensAccessTokenPost.md)
 - [AccessTokensAccessTokenPostInlineRole](docs/AccessTokensAccessTokenPostInlineRole.md)
 - [AccountsClientSideAvailability](docs/AccountsClientSideAvailability.md)
 - [AccountsIntegrationSubscriptionMetadata](docs/AccountsIntegrationSubscriptionMetadata.md)
 - [Api2IpList](docs/Api2IpList.md)
 - [Api2ValuePut](docs/Api2ValuePut.md)
 - [ApiBranchCollectionRep](docs/ApiBranchCollectionRep.md)
 - [ApiBranchCollectionRepHunks](docs/ApiBranchCollectionRepHunks.md)
 - [ApiBranchCollectionRepItems](docs/ApiBranchCollectionRepItems.md)
 - [ApiBranchCollectionRepReferences](docs/ApiBranchCollectionRepReferences.md)
 - [ApiBranchRep](docs/ApiBranchRep.md)
 - [ApiHunkRep](docs/ApiHunkRep.md)
 - [ApiNewMemberForm](docs/ApiNewMemberForm.md)
 - [ApiReferenceRep](docs/ApiReferenceRep.md)
 - [ApiRepositoryCollectionRep](docs/ApiRepositoryCollectionRep.md)
 - [ApiRepositoryCollectionRepItems](docs/ApiRepositoryCollectionRepItems.md)
 - [ApiRepositoryRep](docs/ApiRepositoryRep.md)
 - [ApiStatisticCollectionRep](docs/ApiStatisticCollectionRep.md)
 - [ApiStatisticRep](docs/ApiStatisticRep.md)
 - [ApiV2IntegrationsItems](docs/ApiV2IntegrationsItems.md)
 - [ApprovalSettingsRep](docs/ApprovalSettingsRep.md)
 - [ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest](docs/ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest.md)
 - [ApprovalsEndpointsCreateFlagConfigApprovalRequestRequest](docs/ApprovalsEndpointsCreateFlagConfigApprovalRequestRequest.md)
 - [ApprovalsEndpointsPostApprovalRequestApplyRequest](docs/ApprovalsEndpointsPostApprovalRequestApplyRequest.md)
 - [ApprovalsEndpointsPostApprovalRequestReviewRequest](docs/ApprovalsEndpointsPostApprovalRequestReviewRequest.md)
 - [ApprovalsEndpointsSourceFlag](docs/ApprovalsEndpointsSourceFlag.md)
 - [AuditLogEntryListingRep](docs/AuditLogEntryListingRep.md)
 - [AuditLogEntryListingRepAccesses](docs/AuditLogEntryListingRepAccesses.md)
 - [AuditLogEntryListingRepCollection](docs/AuditLogEntryListingRepCollection.md)
 - [AuditLogEntryListingRepCollectionItems](docs/AuditLogEntryListingRepCollectionItems.md)
 - [AuthorizedAppDataRep](docs/AuthorizedAppDataRep.md)
 - [ClausesClause](docs/ClausesClause.md)
 - [ClientSideAvailabilityPost](docs/ClientSideAvailabilityPost.md)
 - [ConfidenceIntervalRep](docs/ConfidenceIntervalRep.md)
 - [CoreLink](docs/CoreLink.md)
 - [CustomProperty](docs/CustomProperty.md)
 - [CustomRoleCollectionRep](docs/CustomRoleCollectionRep.md)
 - [CustomRoleCollectionRepItems](docs/CustomRoleCollectionRepItems.md)
 - [CustomRoleRep](docs/CustomRoleRep.md)
 - [DefaultClientSideAvailabilityPost](docs/DefaultClientSideAvailabilityPost.md)
 - [DependentFlag](docs/DependentFlag.md)
 - [DependentFlagEnvironment](docs/DependentFlagEnvironment.md)
 - [DependentFlagWithEnvs](docs/DependentFlagWithEnvs.md)
 - [DependentFlagWithEnvsEnvironments](docs/DependentFlagWithEnvsEnvironments.md)
 - [DependentFlagsCollectionRep](docs/DependentFlagsCollectionRep.md)
 - [DestinationCollectionRep](docs/DestinationCollectionRep.md)
 - [DestinationPostRep](docs/DestinationPostRep.md)
 - [DestinationRep](docs/DestinationRep.md)
 - [EnvironmentPost](docs/EnvironmentPost.md)
 - [EnvironmentRep](docs/EnvironmentRep.md)
 - [ExperimentAllocationRep](docs/ExperimentAllocationRep.md)
 - [ExperimentEnvironmentSettingRep](docs/ExperimentEnvironmentSettingRep.md)
 - [ExperimentFlagRep](docs/ExperimentFlagRep.md)
 - [ExperimentFlagRepVariations](docs/ExperimentFlagRepVariations.md)
 - [ExperimentInfoRep](docs/ExperimentInfoRep.md)
 - [ExperimentInfoRepEnvironmentSettings](docs/ExperimentInfoRepEnvironmentSettings.md)
 - [ExperimentInfoRepItems](docs/ExperimentInfoRepItems.md)
 - [ExperimentMetadataRep](docs/ExperimentMetadataRep.md)
 - [ExperimentPost](docs/ExperimentPost.md)
 - [ExperimentRep](docs/ExperimentRep.md)
 - [ExperimentResultsRep](docs/ExperimentResultsRep.md)
 - [ExperimentResultsRepMetadata](docs/ExperimentResultsRepMetadata.md)
 - [ExperimentResultsRepSeries](docs/ExperimentResultsRepSeries.md)
 - [ExperimentResultsRepTotals](docs/ExperimentResultsRepTotals.md)
 - [ExperimentStatsRep](docs/ExperimentStatsRep.md)
 - [ExperimentSummaryRep](docs/ExperimentSummaryRep.md)
 - [ExperimentTimeSeriesSlice](docs/ExperimentTimeSeriesSlice.md)
 - [ExperimentTimeSeriesVariationSlice](docs/ExperimentTimeSeriesVariationSlice.md)
 - [ExperimentTotalsRep](docs/ExperimentTotalsRep.md)
 - [ExpiringUserTargetsEndpointsPatchSegmentInstruction](docs/ExpiringUserTargetsEndpointsPatchSegmentInstruction.md)
 - [ExpiringUserTargetsEndpointsPatchSegmentRequest](docs/ExpiringUserTargetsEndpointsPatchSegmentRequest.md)
 - [ExpiringUserTargetsEndpointsPatchSegmentRequestInstructions](docs/ExpiringUserTargetsEndpointsPatchSegmentRequestInstructions.md)
 - [FlagConfigurationRep](docs/FlagConfigurationRep.md)
 - [FlagConfigurationRepClauses](docs/FlagConfigurationRepClauses.md)
 - [FlagConfigurationRepPrerequisites](docs/FlagConfigurationRepPrerequisites.md)
 - [FlagConfigurationRepRules](docs/FlagConfigurationRepRules.md)
 - [FlagConfigurationRepTargets](docs/FlagConfigurationRepTargets.md)
 - [FlagDefaultsRep](docs/FlagDefaultsRep.md)
 - [FlagGlobalAttributesRep](docs/FlagGlobalAttributesRep.md)
 - [FlagListingRep](docs/FlagListingRep.md)
 - [FlagPost](docs/FlagPost.md)
 - [FlagPostVariations](docs/FlagPostVariations.md)
 - [FlagStatusCollectionRep](docs/FlagStatusCollectionRep.md)
 - [FlagStatusRep](docs/FlagStatusRep.md)
 - [FlagStatusRepFromEnvSummaries](docs/FlagStatusRepFromEnvSummaries.md)
 - [FlagStatusRepFromEnvSummariesEnvironments](docs/FlagStatusRepFromEnvSummariesEnvironments.md)
 - [FlagStatusesRep](docs/FlagStatusesRep.md)
 - [FlagSummary](docs/FlagSummary.md)
 - [FlagsFlagCopyConfigEnvironment](docs/FlagsFlagCopyConfigEnvironment.md)
 - [FlagsFlagCopyConfigPost](docs/FlagsFlagCopyConfigPost.md)
 - [FlagsPrerequisite](docs/FlagsPrerequisite.md)
 - [FlagsTarget](docs/FlagsTarget.md)
 - [FlagsVariate](docs/FlagsVariate.md)
 - [GlobalFlagCollectionRep](docs/GlobalFlagCollectionRep.md)
 - [GlobalFlagRep](docs/GlobalFlagRep.md)
 - [GlobalFlagRepEnvironments](docs/GlobalFlagRepEnvironments.md)
 - [GlobalSegmentCollectionRep](docs/GlobalSegmentCollectionRep.md)
 - [GoalsModification](docs/GoalsModification.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [IntegrationStatusRep](docs/IntegrationStatusRep.md)
 - [IntegrationSubscriptionRep](docs/IntegrationSubscriptionRep.md)
 - [IntegrationSubscriptionRepCollection](docs/IntegrationSubscriptionRepCollection.md)
 - [IntegrationSubscriptionStatusRep](docs/IntegrationSubscriptionStatusRep.md)
 - [IntegrationSubscriptionStatusRepErrors](docs/IntegrationSubscriptionStatusRepErrors.md)
 - [IntegrationSubscriptionTestEventRep](docs/IntegrationSubscriptionTestEventRep.md)
 - [IntegrationsSubscriptionPost](docs/IntegrationsSubscriptionPost.md)
 - [JSONPatchElt](docs/JSONPatchElt.md)
 - [LastSeenMetadata](docs/LastSeenMetadata.md)
 - [MemberCollectionRep](docs/MemberCollectionRep.md)
 - [MemberDataRep](docs/MemberDataRep.md)
 - [MemberRep](docs/MemberRep.md)
 - [MemberRepTeams](docs/MemberRepTeams.md)
 - [MemberSummaryRep](docs/MemberSummaryRep.md)
 - [MemberTeamSummaryRep](docs/MemberTeamSummaryRep.md)
 - [MetricCollectionRep](docs/MetricCollectionRep.md)
 - [MetricListingRep](docs/MetricListingRep.md)
 - [MetricRep](docs/MetricRep.md)
 - [MetricRepAttachedFeatures](docs/MetricRepAttachedFeatures.md)
 - [MetricsMetricPost](docs/MetricsMetricPost.md)
 - [MetricsMetricPostUrls](docs/MetricsMetricPostUrls.md)
 - [ModelsDerivedAttribute](docs/ModelsDerivedAttribute.md)
 - [ModelsUser](docs/ModelsUser.md)
 - [ModelsUserDerived](docs/ModelsUserDerived.md)
 - [MultiEnvDependentFlagsCollectionRep](docs/MultiEnvDependentFlagsCollectionRep.md)
 - [MultiEnvDependentFlagsCollectionRepItems](docs/MultiEnvDependentFlagsCollectionRepItems.md)
 - [ParentResourceRep](docs/ParentResourceRep.md)
 - [PatchWithComment](docs/PatchWithComment.md)
 - [ProjectCollectionRep](docs/ProjectCollectionRep.md)
 - [ProjectListingRep](docs/ProjectListingRep.md)
 - [ProjectPost](docs/ProjectPost.md)
 - [ProjectRep](docs/ProjectRep.md)
 - [PubNubDetailRep](docs/PubNubDetailRep.md)
 - [RolesActionList](docs/RolesActionList.md)
 - [RolesResourceAccess](docs/RolesResourceAccess.md)
 - [RolesResourceList](docs/RolesResourceList.md)
 - [RolesStatement](docs/RolesStatement.md)
 - [RolesStatementPost](docs/RolesStatementPost.md)
 - [RolesStatementPostData](docs/RolesStatementPostData.md)
 - [RolloutRep](docs/RolloutRep.md)
 - [RolloutRepVariations](docs/RolloutRepVariations.md)
 - [RuleRep](docs/RuleRep.md)
 - [ScheduledChangesCollectionRep](docs/ScheduledChangesCollectionRep.md)
 - [ScheduledChangesRep](docs/ScheduledChangesRep.md)
 - [SegmentRep](docs/SegmentRep.md)
 - [SegmentRepRules](docs/SegmentRepRules.md)
 - [SegmentRuleRep](docs/SegmentRuleRep.md)
 - [SegmentsSegmentPost](docs/SegmentsSegmentPost.md)
 - [SegmentsUnboundedTargetRep](docs/SegmentsUnboundedTargetRep.md)
 - [SharedUrlPost](docs/SharedUrlPost.md)
 - [StatementRep](docs/StatementRep.md)
 - [StatementStatementPost](docs/StatementStatementPost.md)
 - [StatementStatementPostData](docs/StatementStatementPostData.md)
 - [SubjectDataRep](docs/SubjectDataRep.md)
 - [TargetResourceRep](docs/TargetResourceRep.md)
 - [TitleRep](docs/TitleRep.md)
 - [TokenDataRep](docs/TokenDataRep.md)
 - [UnboundedSegmentMetadata](docs/UnboundedSegmentMetadata.md)
 - [UnboundedSegmentUserList](docs/UnboundedSegmentUserList.md)
 - [UnboundedSegmentUserState](docs/UnboundedSegmentUserState.md)
 - [UserListRep](docs/UserListRep.md)
 - [UserListRepItems](docs/UserListRepItems.md)
 - [UserRecord](docs/UserRecord.md)
 - [UserRep](docs/UserRep.md)
 - [UserSettingRep](docs/UserSettingRep.md)
 - [UserSettingsCollection](docs/UserSettingsCollection.md)
 - [UserSettingsCollectionItems](docs/UserSettingsCollectionItems.md)
 - [VariateRep](docs/VariateRep.md)
 - [VariationOrRolloutRep](docs/VariationOrRolloutRep.md)
 - [VariationSummary](docs/VariationSummary.md)
 - [VersionsRep](docs/VersionsRep.md)
 - [WebConflict](docs/WebConflict.md)
 - [WebConflictResponse](docs/WebConflictResponse.md)
 - [WebConflictResponseConflict](docs/WebConflictResponseConflict.md)
 - [WebConflictResponseConflicts](docs/WebConflictResponseConflicts.md)
 - [WebConflictResponseInstruction](docs/WebConflictResponseInstruction.md)
 - [WebConflictResponseItems](docs/WebConflictResponseItems.md)
 - [WebCopiedFromEnv](docs/WebCopiedFromEnv.md)
 - [WebExpiringUserTargetError](docs/WebExpiringUserTargetError.md)
 - [WebExpiringUserTargetItem](docs/WebExpiringUserTargetItem.md)
 - [WebExpiringUserTargetResponse](docs/WebExpiringUserTargetResponse.md)
 - [WebExpiringUserTargetResponseErrors](docs/WebExpiringUserTargetResponseErrors.md)
 - [WebExpiringUserTargetResponseItems](docs/WebExpiringUserTargetResponseItems.md)
 - [WebFlagConfigApprovalRequestResponse](docs/WebFlagConfigApprovalRequestResponse.md)
 - [WebFlagConfigApprovalRequestResponseAllReviews](docs/WebFlagConfigApprovalRequestResponseAllReviews.md)
 - [WebFlagConfigApprovalRequestResponseConflicts](docs/WebFlagConfigApprovalRequestResponseConflicts.md)
 - [WebFlagScheduledChangesInput](docs/WebFlagScheduledChangesInput.md)
 - [WebIntegrationMetadata](docs/WebIntegrationMetadata.md)
 - [WebIntegrationStatus](docs/WebIntegrationStatus.md)
 - [WebPostFlagScheduledChangesInput](docs/WebPostFlagScheduledChangesInput.md)
 - [WebReportFlagScheduledChangesInput](docs/WebReportFlagScheduledChangesInput.md)
 - [WebResourceIDResponse](docs/WebResourceIDResponse.md)
 - [WebReviewResponse](docs/WebReviewResponse.md)
 - [WeightedVariationRep](docs/WeightedVariationRep.md)


## Documentation For Authorization



### ApiKey

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@launchdarkly.com

## Sample Code

```go
package main

import (
	"context"
	"fmt"
	"os"

	ldapi "github.com/launchdarkly/api-client-go"
)

func main() {
	apiKey := os.Getenv("LD_API_KEY")
	if apiKey == "" {
		panic("LD_API_KEY env var was empty!")
	}
	client := ldapi.NewAPIClient(ldapi.NewConfiguration())

	auth := make(map[string]ldapi.APIKey)
	auth["ApiKey"] = ldapi.APIKey{
		Key: apiKey,
	}

	ctx := context.WithValue(context.Background(), ldapi.ContextAPIKeys, auth)

	flagName := "Test Flag Go"
	flagKey := "test-go"
	// Create a multi-variate feature flag
	valOneVal := []int{1, 2}
	valOne := map[string]interface{}{"one": valOneVal}
	valTwoVal := []int{4, 5}
	valTwo := map[string]interface{}{"two": valTwoVal}
	body := ldapi.Reps2GlobalFlagRep{
		Name: &flagName,
		Key:  &flagKey,
		Variations: &[]ldapi.Reps2VariateRep{
			{Value: &valOne},
			{Value: &valTwo},
		},
	}
	flag, resp, err := client.DefaultApi.PostFeatureFlag(ctx, "openapi").Reps2GlobalFlagRep(body).Execute()
	if err != nil {
		if resp.StatusCode != 409 {
			panic(fmt.Errorf("create failed: %s", err))
		} else {
			if _, err := client.DefaultApi.DeleteFeatureFlag(ctx, "openapi", *body.Key).Execute(); err != nil {
				panic(fmt.Errorf("delete failed: %s", err))
			}
			flag, resp, err = client.DefaultApi.PostFeatureFlag(ctx, "openapi").Reps2GlobalFlagRep(body).Execute()
			if err != nil {
				panic(fmt.Errorf("create failed: %s", err))
			}
		}
	}
	fmt.Printf("Created flag: %+v\n", flag)
	// Clean up new flag
	defer func() {
		if _, err := client.DefaultApi.DeleteFeatureFlag(ctx, "openapi", *body.Key).Execute(); err != nil {
			panic(fmt.Errorf("delete failed: %s", err))
		}
	}()
}
```
