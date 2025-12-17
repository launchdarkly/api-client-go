/*
LaunchDarkly REST API

This documentation describes LaunchDarkly's REST API. To access the complete OpenAPI spec directly, use [Get OpenAPI spec](https://launchdarkly.com/docs/api/other/get-openapi-spec).  To learn how to use LaunchDarkly using the user interface (UI) instead, read our [product documentation](https://launchdarkly.com/docs/home).  ## Authentication  LaunchDarkly's REST API uses the HTTPS protocol with a minimum TLS version of 1.2.  All REST API resources are authenticated with either [personal or service access tokens](https://launchdarkly.com/docs/home/account/api), or session cookies. Other authentication mechanisms are not supported. You can manage personal access tokens on your [**Authorization**](https://app.launchdarkly.com/settings/authorization) page in the LaunchDarkly UI.  LaunchDarkly also has SDK keys, mobile keys, and client-side IDs that are used by our server-side SDKs, mobile SDKs, and JavaScript-based SDKs, respectively. **These keys cannot be used to access our REST API**. These keys are environment-specific, and can only perform read-only operations such as fetching feature flag settings.  | Auth mechanism                                                                                  | Allowed resources                                                                                     | Use cases                                          | | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | -------------------------------------------------- | | [Personal or service access tokens](https://launchdarkly.com/docs/home/account/api) | Can be customized on a per-token basis                                                                | Building scripts, custom integrations, data export. | | SDK keys                                                                                        | Can only access read-only resources specific to server-side SDKs. Restricted to a single environment. | Server-side SDKs                     | | Mobile keys                                                                                     | Can only access read-only resources specific to mobile SDKs, and only for flags marked available to mobile keys. Restricted to a single environment.           | Mobile SDKs                                        | | Client-side ID                                                                                  | Can only access read-only resources specific to JavaScript-based client-side SDKs, and only for flags marked available to client-side. Restricted to a single environment.           | Client-side JavaScript                             |  > #### Keep your access tokens and SDK keys private > > Access tokens should _never_ be exposed in untrusted contexts. Never put an access token in client-side JavaScript, or embed it in a mobile application. LaunchDarkly has special mobile keys that you can embed in mobile apps. If you accidentally expose an access token or SDK key, you can reset it from your [**Authorization**](https://app.launchdarkly.com/settings/authorization) page. > > The client-side ID is safe to embed in untrusted contexts. It's designed for use in client-side JavaScript.  ### Authentication using request header  The preferred way to authenticate with the API is by adding an `Authorization` header containing your access token to your requests. The value of the `Authorization` header must be your access token.  Manage personal access tokens from the [**Authorization**](https://app.launchdarkly.com/settings/authorization) page.  ### Authentication using session cookie  For testing purposes, you can make API calls directly from your web browser. If you are logged in to the LaunchDarkly application, the API will use your existing session to authenticate calls.  Depending on the permissions granted as part of your [role](https://launchdarkly.com/docs/home/account/roles), you may not have permission to perform some API calls. You will receive a `401` response code in that case.  > ### Modifying the Origin header causes an error > > LaunchDarkly validates that the Origin header for any API request authenticated by a session cookie matches the expected Origin header. The expected Origin header is `https://app.launchdarkly.com`. > > If the Origin header does not match what's expected, LaunchDarkly returns an error. This error can prevent the LaunchDarkly app from working correctly. > > Any browser extension that intentionally changes the Origin header can cause this problem. For example, the `Allow-Control-Allow-Origin: *` Chrome extension changes the Origin header to `http://evil.com` and causes the app to fail. > > To prevent this error, do not modify your Origin header. > > LaunchDarkly does not require origin matching when authenticating with an access token, so this issue does not affect normal API usage.  ## Representations  All resources expect and return JSON response bodies. Error responses also send a JSON body. To learn more about the error format of the API, read [Errors](https://launchdarkly.com/docs/api#errors).  In practice this means that you always get a response with a `Content-Type` header set to `application/json`.  In addition, request bodies for `PATCH`, `POST`, and `PUT` requests must be encoded as JSON with a `Content-Type` header set to `application/json`.  ### Summary and detailed representations  When you fetch a list of resources, the response includes only the most important attributes of each resource. This is a _summary representation_ of the resource. When you fetch an individual resource, such as a single feature flag, you receive a _detailed representation_ of the resource.  The best way to find a detailed representation is to follow links. Every summary representation includes a link to its detailed representation.  ### Expanding responses  Sometimes the detailed representation of a resource does not include all of the attributes of the resource by default. If this is the case, the request method will clearly document this and describe which attributes you can include in an expanded response.  To include the additional attributes, append the `expand` request parameter to your request and add a comma-separated list of the attributes to include. For example, when you append `?expand=members,maintainers` to the [Get team](https://launchdarkly.com/docs/api/teams/get-team) endpoint, the expanded response includes both of these attributes.  ### Links and addressability  The best way to navigate the API is by following links. These are attributes in representations that link to other resources. The API always uses the same format for links:  - Links to other resources within the API are encapsulated in a `_links` object - If the resource has a corresponding link to HTML content on the site, it is stored in a special `_site` link  Each link has two attributes:  - An `href`, which contains the URL - A `type`, which describes the content type  For example, a feature resource might return the following:  ```json {   \"_links\": {     \"parent\": {       \"href\": \"/api/features\",       \"type\": \"application/json\"     },     \"self\": {       \"href\": \"/api/features/sort.order\",       \"type\": \"application/json\"     }   },   \"_site\": {     \"href\": \"/features/sort.order\",     \"type\": \"text/html\"   } } ```  From this, you can navigate to the parent collection of features by following the `parent` link, or navigate to the site page for the feature by following the `_site` link.  Collections are always represented as a JSON object with an `items` attribute containing an array of representations. Like all other representations, collections have `_links` defined at the top level.  Paginated collections include `first`, `last`, `next`, and `prev` links containing a URL with the respective set of elements in the collection.  ## Updates  Resources that accept partial updates use the `PATCH` verb. Most resources support the [JSON patch](https://launchdarkly.com/docs/api#updates-using-json-patch) format. Some resources also support the [JSON merge patch](https://launchdarkly.com/docs/api#updates-using-json-merge-patch) format, and some resources support the [semantic patch](https://launchdarkly.com/docs/api#updates-using-semantic-patch) format, which is a way to specify the modifications to perform as a set of executable instructions. Each resource supports optional [comments](https://launchdarkly.com/docs/api#updates-with-comments) that you can submit with updates. Comments appear in outgoing webhooks, the audit log, and other integrations.  When a resource supports both JSON patch and semantic patch, we document both in the request method. However, the specific request body fields and descriptions included in our documentation only match one type of patch or the other.  ### Updates using JSON patch  [JSON patch](https://datatracker.ietf.org/doc/html/rfc6902) is a way to specify the modifications to perform on a resource. JSON patch uses paths and a limited set of operations to describe how to transform the current state of the resource into a new state. JSON patch documents are always arrays, where each element contains an operation, a path to the field to update, and the new value.  For example, in this feature flag representation:  ```json {     \"name\": \"New recommendations engine\",     \"key\": \"engine.enable\",     \"description\": \"This is the description\",     ... } ``` You can change the feature flag's description with the following patch document:  ```json [{ \"op\": \"replace\", \"path\": \"/description\", \"value\": \"This is the new description\" }] ```  You can specify multiple modifications to perform in a single request. You can also test that certain preconditions are met before applying the patch:  ```json [   { \"op\": \"test\", \"path\": \"/version\", \"value\": 10 },   { \"op\": \"replace\", \"path\": \"/description\", \"value\": \"The new description\" } ] ```  The above patch request tests whether the feature flag's `version` is `10`, and if so, changes the feature flag's description.  Attributes that are not editable, such as a resource's `_links`, have names that start with an underscore.  ### Updates using JSON merge patch  [JSON merge patch](https://datatracker.ietf.org/doc/html/rfc7386) is another format for specifying the modifications to perform on a resource. JSON merge patch is less expressive than JSON patch. However, in many cases it is simpler to construct a merge patch document. For example, you can change a feature flag's description with the following merge patch document:  ```json {   \"description\": \"New flag description\" } ```  ### Updates using semantic patch  Some resources support the semantic patch format. A semantic patch is a way to specify the modifications to perform on a resource as a set of executable instructions.  Semantic patch allows you to be explicit about intent using precise, custom instructions. In many cases, you can define semantic patch instructions independently of the current state of the resource. This can be useful when defining a change that may be applied at a future date.  To make a semantic patch request, you must append `domain-model=launchdarkly.semanticpatch` to your `Content-Type` header.  Here's how:  ``` Content-Type: application/json; domain-model=launchdarkly.semanticpatch ```  If you call a semantic patch resource without this header, you will receive a `400` response because your semantic patch will be interpreted as a JSON patch.  The body of a semantic patch request takes the following properties:  * `comment` (string): (Optional) A description of the update. * `environmentKey` (string): (Required for some resources only) The environment key. * `instructions` (array): (Required) A list of actions the update should perform. Each action in the list must be an object with a `kind` property that indicates the instruction. If the instruction requires parameters, you must include those parameters as additional fields in the object. The documentation for each resource that supports semantic patch includes the available instructions and any additional parameters.  For example:  ```json {   \"comment\": \"optional comment\",   \"instructions\": [ {\"kind\": \"turnFlagOn\"} ] } ```  Semantic patches are not applied partially; either all of the instructions are applied or none of them are. If **any** instruction is invalid, the endpoint returns an error and will not change the resource. If all instructions are valid, the request succeeds and the resources are updated if necessary, or left unchanged if they are already in the state you request.  ### Updates with comments  You can submit optional comments with `PATCH` changes.  To submit a comment along with a JSON patch document, use the following format:  ```json {   \"comment\": \"This is a comment string\",   \"patch\": [{ \"op\": \"replace\", \"path\": \"/description\", \"value\": \"The new description\" }] } ```  To submit a comment along with a JSON merge patch document, use the following format:  ```json {   \"comment\": \"This is a comment string\",   \"merge\": { \"description\": \"New flag description\" } } ```  To submit a comment along with a semantic patch, use the following format:  ```json {   \"comment\": \"This is a comment string\",   \"instructions\": [ {\"kind\": \"turnFlagOn\"} ] } ```  ## Errors  The API always returns errors in a common format. Here's an example:  ```json {   \"code\": \"invalid_request\",   \"message\": \"A feature with that key already exists\",   \"id\": \"30ce6058-87da-11e4-b116-123b93f75cba\" } ```  The `code` indicates the general class of error. The `message` is a human-readable explanation of what went wrong. The `id` is a unique identifier. Use it when you're working with LaunchDarkly Support to debug a problem with a specific API call.  ### HTTP status error response codes  | Code | Definition        | Description                                                                                       | Possible Solution                                                | | ---- | ----------------- | ------------------------------------------------------------------------------------------- | ---------------------------------------------------------------- | | 400  | Invalid request       | The request cannot be understood.                                    | Ensure JSON syntax in request body is correct.                   | | 401  | Invalid access token      | Requestor is unauthorized or does not have permission for this API call.                                                | Ensure your API access token is valid and has the appropriate permissions.                                     | | 403  | Forbidden         | Requestor does not have access to this resource.                                                | Ensure that the account member or access token has proper permissions set. | | 404  | Invalid resource identifier | The requested resource is not valid. | Ensure that the resource is correctly identified by ID or key. | | 405  | Method not allowed | The request method is not allowed on this resource. | Ensure that the HTTP verb is correct. | | 409  | Conflict          | The API request can not be completed because it conflicts with a concurrent API request. | Retry your request.                                              | | 422  | Unprocessable entity | The API request can not be completed because the update description can not be understood. | Ensure that the request body is correct for the type of patch you are using, either JSON patch or semantic patch. | 429  | Too many requests | Read [Rate limiting](https://launchdarkly.com/docs/api#rate-limiting).                                               | Wait and try again later.                                        |  ## CORS  The LaunchDarkly API supports Cross Origin Resource Sharing (CORS) for AJAX requests from any origin. If an `Origin` header is given in a request, it will be echoed as an explicitly allowed origin. Otherwise the request returns a wildcard, `Access-Control-Allow-Origin: *`. For more information on CORS, read the [CORS W3C Recommendation](http://www.w3.org/TR/cors). Example CORS headers might look like:  ```http Access-Control-Allow-Headers: Accept, Content-Type, Content-Length, Accept-Encoding, Authorization Access-Control-Allow-Methods: OPTIONS, GET, DELETE, PATCH Access-Control-Allow-Origin: * Access-Control-Max-Age: 300 ```  You can make authenticated CORS calls just as you would make same-origin calls, using either [token or session-based authentication](https://launchdarkly.com/docs/api#authentication). If you are using session authentication, you should set the `withCredentials` property for your `xhr` request to `true`. You should never expose your access tokens to untrusted entities.  ## Rate limiting  We use several rate limiting strategies to ensure the availability of our APIs. Rate-limited calls to our APIs return a `429` status code. Calls to our APIs include headers indicating the current rate limit status. The specific headers returned depend on the API route being called. The limits differ based on the route, authentication mechanism, and other factors. Routes that are not rate limited may not contain any of the headers described below.  > ### Rate limiting and SDKs > > LaunchDarkly SDKs are never rate limited and do not use the API endpoints defined here. LaunchDarkly uses a different set of approaches, including streaming/server-sent events and a global CDN, to ensure availability to the routes used by LaunchDarkly SDKs.  ### Global rate limits  Authenticated requests are subject to a global limit. This is the maximum number of calls that your account can make to the API per ten seconds. All service and personal access tokens on the account share this limit, so exceeding the limit with one access token will impact other tokens. Calls that are subject to global rate limits may return the headers below:  | Header name                    | Description                                                                      | | ------------------------------ | -------------------------------------------------------------------------------- | | `X-Ratelimit-Global-Remaining` | The maximum number of requests the account is permitted to make per ten seconds. | | `X-Ratelimit-Reset`            | The time at which the current rate limit window resets in epoch milliseconds.    |  We do not publicly document the specific number of calls that can be made globally. This limit may change, and we encourage clients to program against the specification, relying on the two headers defined above, rather than hardcoding to the current limit.  ### Route-level rate limits  Some authenticated routes have custom rate limits. These also reset every ten seconds. Any service or personal access tokens hitting the same route share this limit, so exceeding the limit with one access token may impact other tokens. Calls that are subject to route-level rate limits return the headers below:  | Header name                   | Description                                                                                           | | ----------------------------- | ----------------------------------------------------------------------------------------------------- | | `X-Ratelimit-Route-Remaining` | The maximum number of requests to the current route the account is permitted to make per ten seconds. | | `X-Ratelimit-Reset`           | The time at which the current rate limit window resets in epoch milliseconds.                         |  A _route_ represents a specific URL pattern and verb. For example, the [Delete environment](https://launchdarkly.com/docs/api/environments/delete-environment) endpoint is considered a single route, and each call to delete an environment counts against your route-level rate limit for that route.  We do not publicly document the specific number of calls that an account can make to each endpoint per ten seconds. These limits may change, and we encourage clients to program against the specification, relying on the two headers defined above, rather than hardcoding to the current limits.  ### IP-based rate limiting  We also employ IP-based rate limiting on some API routes. If you hit an IP-based rate limit, your API response will include a `Retry-After` header indicating how long to wait before re-trying the call. Clients must wait at least `Retry-After` seconds before making additional calls to our API, and should employ jitter and backoff strategies to avoid triggering rate limits again.  ## OpenAPI (Swagger) and client libraries  We have a [complete OpenAPI (Swagger) specification](https://app.launchdarkly.com/api/v2/openapi.json) for our API.  We auto-generate multiple client libraries based on our OpenAPI specification. To learn more, visit the [collection of client libraries on GitHub](https://github.com/search?q=topic%3Alaunchdarkly-api+org%3Alaunchdarkly&type=Repositories). You can also use this specification to generate client libraries to interact with our REST API in your language of choice.  Our OpenAPI specification is supported by several API-based tools such as Postman and Insomnia. In many cases, you can directly import our specification to explore our APIs.  ## Method overriding  Some firewalls and HTTP clients restrict the use of verbs other than `GET` and `POST`. In those environments, our API endpoints that use `DELETE`, `PATCH`, and `PUT` verbs are inaccessible.  To avoid this issue, our API supports the `X-HTTP-Method-Override` header, allowing clients to \"tunnel\" `DELETE`, `PATCH`, and `PUT` requests using a `POST` request.  For example, to call a `PATCH` endpoint using a `POST` request, you can include `X-HTTP-Method-Override:PATCH` as a header.  ## Beta resources  We sometimes release new API resources in **beta** status before we release them with general availability.  Resources that are in beta are still undergoing testing and development. They may change without notice, including becoming backwards incompatible.  We try to promote resources into general availability as quickly as possible. This happens after sufficient testing and when we're satisfied that we no longer need to make backwards-incompatible changes.  We mark beta resources with a \"Beta\" callout in our documentation, pictured below:  > ### This feature is in beta > > To use this feature, pass in a header including the `LD-API-Version` key with value set to `beta`. Use this header with each call. To learn more, read [Beta resources](https://launchdarkly.com/docs/api#beta-resources). > > Resources that are in beta are still undergoing testing and development. They may change without notice, including becoming backwards incompatible.  ### Using beta resources  To use a beta resource, you must include a header in the request. If you call a beta resource without this header, you receive a `403` response.  Use this header:  ``` LD-API-Version: beta ```  ## Federal and EU environments  In addition to the commercial versions, LaunchDarkly offers instances for federal agencies and those based in the European Union (EU).  ### Federal environments  The version of LaunchDarkly that is available on domains controlled by the United States government is different from the version of LaunchDarkly available to the general public. If you are an employee or contractor for a United States federal agency and use LaunchDarkly in your work, you likely use the federal instance of LaunchDarkly.  If you are working in the federal instance of LaunchDarkly, the base URI for each request is `https://app.launchdarkly.us`.  To learn more, read [LaunchDarkly in federal environments](https://launchdarkly.com/docs/home/infrastructure/federal).  ### EU environments  The version of LaunchDarkly that is available in the EU is different from the version of LaunchDarkly available to other regions. If you are based in the EU, you likely use the EU instance of LaunchDarkly. The LaunchDarkly EU instance complies with EU data residency principles, including the protection and confidentiality of EU customer information.  If you are working in the EU instance of LaunchDarkly, the base URI for each request is `https://app.eu.launchdarkly.com`.  To learn more, read [LaunchDarkly in the European Union (EU)](https://launchdarkly.com/docs/home/infrastructure/eu).  ## Versioning  We try hard to keep our REST API backwards compatible, but we occasionally have to make backwards-incompatible changes in the process of shipping new features. These breaking changes can cause unexpected behavior if you don't prepare for them accordingly.  Updates to our REST API include support for the latest features in LaunchDarkly. We also release a new version of our REST API every time we make a breaking change. We provide simultaneous support for multiple API versions so you can migrate from your current API version to a new version at your own pace.  ### Setting the API version per request  You can set the API version on a specific request by sending an `LD-API-Version` header, as shown in the example below:  ``` LD-API-Version: 20240415 ```  The header value is the version number of the API version you would like to request. The number for each version corresponds to the date the version was released in `yyyymmdd` format. In the example above the version `20240415` corresponds to April 15, 2024.  ### Setting the API version per access token  When you create an access token, you must specify a specific version of the API to use. This ensures that integrations using this token cannot be broken by version changes.  Tokens created before versioning was released have their version set to `20160426`, which is the version of the API that existed before the current versioning scheme, so that they continue working the same way they did before versioning.  If you would like to upgrade your integration to use a new API version, you can explicitly set the header described above.  > ### Best practice: Set the header for every client or integration > > We recommend that you set the API version header explicitly in any client or integration you build. > > Only rely on the access token API version during manual testing.  ### API version changelog  <table>   <tr>     <th>Version</th>     <th>Changes</th>     <th>End of life (EOL)</th>   </tr>   <tr>     <td>`20240415`</td>     <td>       <ul><li>Changed several endpoints from unpaginated to paginated. Use the `limit` and `offset` query parameters to page through the results.</li> <li>Changed the [list access tokens](https://launchdarkly.com/docs/api/access-tokens/get-tokens) endpoint: <ul><li>Response is now paginated with a default limit of `25`</li></ul></li> <li>Changed the [list account members](https://launchdarkly.com/docs/api/account-members/get-members) endpoint: <ul><li>The `accessCheck` filter is no longer available</li></ul></li> <li>Changed the [list custom roles](https://launchdarkly.com/docs/api/custom-roles/get-custom-roles) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li></ul></li> <li>Changed the [list feature flags](https://launchdarkly.com/docs/api/feature-flags/get-feature-flags) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li><li>The `environments` field is now only returned if the request is filtered by environment, using the `filterEnv` query parameter</li><li>The `followerId`, `hasDataExport`, `status`, `contextKindTargeted`, and `segmentTargeted` filters are no longer available</li><li>The `compare` query parameter is no longer available</li></ul></li> <li>Changed the [list segments](https://launchdarkly.com/docs/api/segments/get-segments) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li></ul></li> <li>Changed the [list teams](https://launchdarkly.com/docs/api/teams/get-teams) endpoint: <ul><li>The `expand` parameter no longer supports including `projects` or `roles`</li><li>In paginated results, the maximum page size is now 100</li></ul></li> <li>Changed the [get workflows](https://launchdarkly.com/docs/api/workflows/get-workflows) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li><li>The `_conflicts` field in the response is no longer available</li></ul></li> </ul>     </td>     <td>Current</td>   </tr>   <tr>     <td>`20220603`</td>     <td>       <ul><li>Changed the [list projects](https://launchdarkly.com/docs/api/projects/get-projects) return value:<ul><li>Response is now paginated with a default limit of `20`.</li><li>Added support for filter and sort.</li><li>The project `environments` field is now expandable. This field is omitted by default.</li></ul></li><li>Changed the [get project](https://launchdarkly.com/docs/api/projects/get-project) return value:<ul><li>The `environments` field is now expandable. This field is omitted by default.</li></ul></li></ul>     </td>     <td>2025-04-15</td>   </tr>   <tr>     <td>`20210729`</td>     <td>       <ul><li>Changed the [create approval request](https://launchdarkly.com/docs/api/approvals/post-approval-request) return value. It now returns HTTP Status Code `201` instead of `200`.</li><li> Changed the [get user](https://launchdarkly.com/docs/api/users/get-user) return value. It now returns a user record, not a user. </li><li>Added additional optional fields to environment, segments, flags, members, and segments, including the ability to create big segments. </li><li> Added default values for flag variations when new environments are created. </li><li>Added filtering and pagination for getting flags and members, including `limit`, `number`, `filter`, and `sort` query parameters. </li><li>Added endpoints for expiring user targets for flags and segments, scheduled changes, access tokens, Relay Proxy configuration, integrations and subscriptions, and approvals. </li></ul>     </td>     <td>2023-06-03</td>   </tr>   <tr>     <td>`20191212`</td>     <td>       <ul><li>[List feature flags](https://launchdarkly.com/docs/api/feature-flags/get-feature-flags) now defaults to sending summaries of feature flag configurations, equivalent to setting the query parameter `summary=true`. Summaries omit flag targeting rules and individual user targets from the payload. </li><li> Added endpoints for flags, flag status, projects, environments, audit logs, members, users, custom roles, segments, usage, streams, events, and data export. </li></ul>     </td>     <td>2022-07-29</td>   </tr>   <tr>     <td>`20160426`</td>     <td>       <ul><li>Initial versioning of API. Tokens created before versioning have their version set to this.</li></ul>     </td>     <td>2020-12-12</td>   </tr> </table>  To learn more about how EOL is determined, read LaunchDarkly's [End of Life (EOL) Policy](https://launchdarkly.com/policies/end-of-life-policy/). 

API version: 2.0
Contact: support@launchdarkly.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ldapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type AIConfigsBetaApi interface {

	/*
	DeleteAIConfig Delete AI Config

	Delete an existing AI Config.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param configKey
	@return ApiDeleteAIConfigRequest
	*/
	DeleteAIConfig(ctx context.Context, projectKey string, configKey string) ApiDeleteAIConfigRequest

	// DeleteAIConfigExecute executes the request
	DeleteAIConfigExecute(r ApiDeleteAIConfigRequest) (*http.Response, error)

	/*
	DeleteAIConfigVariation Delete AI Config variation

	Delete a specific variation of an AI Config by config key and variation key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param configKey
	@param variationKey
	@return ApiDeleteAIConfigVariationRequest
	*/
	DeleteAIConfigVariation(ctx context.Context, projectKey string, configKey string, variationKey string) ApiDeleteAIConfigVariationRequest

	// DeleteAIConfigVariationExecute executes the request
	DeleteAIConfigVariationExecute(r ApiDeleteAIConfigVariationRequest) (*http.Response, error)

	/*
	DeleteAITool Delete AI tool

	Delete an existing AI tool.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param toolKey
	@return ApiDeleteAIToolRequest
	*/
	DeleteAITool(ctx context.Context, projectKey string, toolKey string) ApiDeleteAIToolRequest

	// DeleteAIToolExecute executes the request
	DeleteAIToolExecute(r ApiDeleteAIToolRequest) (*http.Response, error)

	/*
	DeleteAgentGraph Delete agent graph

	Delete an existing agent graph and all of its edges.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param graphKey
	@return ApiDeleteAgentGraphRequest
	*/
	DeleteAgentGraph(ctx context.Context, projectKey string, graphKey string) ApiDeleteAgentGraphRequest

	// DeleteAgentGraphExecute executes the request
	DeleteAgentGraphExecute(r ApiDeleteAgentGraphRequest) (*http.Response, error)

	/*
	DeleteModelConfig Delete an AI model config

	Delete an AI model config.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param modelConfigKey
	@return ApiDeleteModelConfigRequest
	*/
	DeleteModelConfig(ctx context.Context, projectKey string, modelConfigKey string) ApiDeleteModelConfigRequest

	// DeleteModelConfigExecute executes the request
	DeleteModelConfigExecute(r ApiDeleteModelConfigRequest) (*http.Response, error)

	/*
	DeleteRestrictedModels Remove AI models from the restricted list

	Remove AI models, by key, from the restricted list.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiDeleteRestrictedModelsRequest
	*/
	DeleteRestrictedModels(ctx context.Context, projectKey string) ApiDeleteRestrictedModelsRequest

	// DeleteRestrictedModelsExecute executes the request
	DeleteRestrictedModelsExecute(r ApiDeleteRestrictedModelsRequest) (*http.Response, error)

	/*
	GetAIConfig Get AI Config

	Retrieve a specific AI Config by its key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param configKey
	@return ApiGetAIConfigRequest
	*/
	GetAIConfig(ctx context.Context, projectKey string, configKey string) ApiGetAIConfigRequest

	// GetAIConfigExecute executes the request
	//  @return AIConfig
	GetAIConfigExecute(r ApiGetAIConfigRequest) (*AIConfig, *http.Response, error)

	/*
	GetAIConfigMetrics Get AI Config metrics

	Retrieve usage metrics for an AI Config by config key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param configKey
	@return ApiGetAIConfigMetricsRequest
	*/
	GetAIConfigMetrics(ctx context.Context, projectKey string, configKey string) ApiGetAIConfigMetricsRequest

	// GetAIConfigMetricsExecute executes the request
	//  @return Metrics
	GetAIConfigMetricsExecute(r ApiGetAIConfigMetricsRequest) (*Metrics, *http.Response, error)

	/*
	GetAIConfigMetricsByVariation Get AI Config metrics by variation

	Retrieve usage metrics for an AI Config by config key, with results split by variation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param configKey
	@return ApiGetAIConfigMetricsByVariationRequest
	*/
	GetAIConfigMetricsByVariation(ctx context.Context, projectKey string, configKey string) ApiGetAIConfigMetricsByVariationRequest

	// GetAIConfigMetricsByVariationExecute executes the request
	//  @return []MetricByVariation
	GetAIConfigMetricsByVariationExecute(r ApiGetAIConfigMetricsByVariationRequest) ([]MetricByVariation, *http.Response, error)

	/*
	GetAIConfigTargeting Show an AI Config's targeting

	Retrieves a specific AI Config's targeting by its key

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param configKey
	@return ApiGetAIConfigTargetingRequest
	*/
	GetAIConfigTargeting(ctx context.Context, projectKey string, configKey string) ApiGetAIConfigTargetingRequest

	// GetAIConfigTargetingExecute executes the request
	//  @return AIConfigTargeting
	GetAIConfigTargetingExecute(r ApiGetAIConfigTargetingRequest) (*AIConfigTargeting, *http.Response, error)

	/*
	GetAIConfigVariation Get AI Config variation

	Get an AI Config variation by key. The response includes all variation versions for the given variation key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param configKey
	@param variationKey
	@return ApiGetAIConfigVariationRequest
	*/
	GetAIConfigVariation(ctx context.Context, projectKey string, configKey string, variationKey string) ApiGetAIConfigVariationRequest

	// GetAIConfigVariationExecute executes the request
	//  @return AIConfigVariationsResponse
	GetAIConfigVariationExecute(r ApiGetAIConfigVariationRequest) (*AIConfigVariationsResponse, *http.Response, error)

	/*
	GetAIConfigs List AI Configs

	Get a list of all AI Configs in the given project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiGetAIConfigsRequest
	*/
	GetAIConfigs(ctx context.Context, projectKey string) ApiGetAIConfigsRequest

	// GetAIConfigsExecute executes the request
	//  @return AIConfigs
	GetAIConfigsExecute(r ApiGetAIConfigsRequest) (*AIConfigs, *http.Response, error)

	/*
	GetAITool Get AI tool

	Retrieve a specific AI tool by its key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param toolKey
	@return ApiGetAIToolRequest
	*/
	GetAITool(ctx context.Context, projectKey string, toolKey string) ApiGetAIToolRequest

	// GetAIToolExecute executes the request
	//  @return AITool
	GetAIToolExecute(r ApiGetAIToolRequest) (*AITool, *http.Response, error)

	/*
	GetAgentGraph Get agent graph

	Retrieve a specific agent graph by its key, including its edges.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param graphKey
	@return ApiGetAgentGraphRequest
	*/
	GetAgentGraph(ctx context.Context, projectKey string, graphKey string) ApiGetAgentGraphRequest

	// GetAgentGraphExecute executes the request
	//  @return AgentGraph
	GetAgentGraphExecute(r ApiGetAgentGraphRequest) (*AgentGraph, *http.Response, error)

	/*
	GetModelConfig Get AI model config

	Get an AI model config by key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param modelConfigKey
	@return ApiGetModelConfigRequest
	*/
	GetModelConfig(ctx context.Context, projectKey string, modelConfigKey string) ApiGetModelConfigRequest

	// GetModelConfigExecute executes the request
	//  @return ModelConfig
	GetModelConfigExecute(r ApiGetModelConfigRequest) (*ModelConfig, *http.Response, error)

	/*
	ListAIToolVersions List AI tool versions

	Get a list of all versions of an AI tool in the given project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param toolKey
	@return ApiListAIToolVersionsRequest
	*/
	ListAIToolVersions(ctx context.Context, projectKey string, toolKey string) ApiListAIToolVersionsRequest

	// ListAIToolVersionsExecute executes the request
	//  @return AITools
	ListAIToolVersionsExecute(r ApiListAIToolVersionsRequest) (*AITools, *http.Response, error)

	/*
	ListAITools List AI tools

	Get a list of all AI tools in the given project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiListAIToolsRequest
	*/
	ListAITools(ctx context.Context, projectKey string) ApiListAIToolsRequest

	// ListAIToolsExecute executes the request
	//  @return AITools
	ListAIToolsExecute(r ApiListAIToolsRequest) (*AITools, *http.Response, error)

	/*
	ListAgentGraphs List agent graphs

	Get a list of all agent graphs in the given project. Returns metadata only, without edge data.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiListAgentGraphsRequest
	*/
	ListAgentGraphs(ctx context.Context, projectKey string) ApiListAgentGraphsRequest

	// ListAgentGraphsExecute executes the request
	//  @return AgentGraphs
	ListAgentGraphsExecute(r ApiListAgentGraphsRequest) (*AgentGraphs, *http.Response, error)

	/*
	ListModelConfigs List AI model configs

	Get all AI model configs for a project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiListModelConfigsRequest
	*/
	ListModelConfigs(ctx context.Context, projectKey string) ApiListModelConfigsRequest

	// ListModelConfigsExecute executes the request
	//  @return []ModelConfig
	ListModelConfigsExecute(r ApiListModelConfigsRequest) ([]ModelConfig, *http.Response, error)

	/*
	PatchAIConfig Update AI Config

	Edit an existing AI Config.

The request body must be a JSON object of the fields to update. The values you include replace the existing values for the fields.

Here's an example:
  ```
    {
      "description": "Example updated description",
      "tags": ["new-tag"]
    }
  ```


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param configKey
	@return ApiPatchAIConfigRequest
	*/
	PatchAIConfig(ctx context.Context, projectKey string, configKey string) ApiPatchAIConfigRequest

	// PatchAIConfigExecute executes the request
	//  @return AIConfig
	PatchAIConfigExecute(r ApiPatchAIConfigRequest) (*AIConfig, *http.Response, error)

	/*
	PatchAIConfigTargeting Update AI Config targeting

	Perform a partial update to an AI Config's targeting. The request body must be a valid semantic patch.

### Using semantic patches on an AI Config

To make a semantic patch request, you must append `domain-model=launchdarkly.semanticpatch` to your `Content-Type` header. To learn more, read [Updates using semantic patch](https://launchdarkly.com/docs/api#updates-using-semantic-patch).

The body of a semantic patch request for updating an AI Config's targeting takes the following properties:

* `comment` (string): (Optional) A description of the update.
* `environmentKey` (string): The key of the LaunchDarkly environment.
* `instructions` (array): (Required) A list of actions the update should perform. Each action in the list must be an object with a `kind` property that indicates the instruction. If the action requires parameters, you must include those parameters as additional fields in the object. The body of a single semantic patch can contain many different instructions.

### Instructions

Semantic patch requests support the following `kind` instructions for updating AI Configs.

<details>
<summary>Click to expand instructions for <strong>working with targeting and variations</strong> for AI Configs</summary>

#### addClauses

Adds the given clauses to the rule indicated by `ruleId`.

##### Parameters

- `ruleId`: ID of a rule in the AI Config.
- `clauses`: Array of clause objects, with `contextKind` (string), `attribute` (string), `op` (string), `negate` (boolean), and `values` (array of strings, numbers, or dates) properties. The `contextKind`, `attribute`, and `values` are case sensitive. The `op` must be lower-case.

Here's an example:

```json
{
  "environmentKey": "environment-key-123abc",
  "instructions": [{
    "kind": "addClauses",
    "ruleId": "a902ef4a-2faf-4eaf-88e1-ecc356708a29",
    "clauses": [{
      "contextKind": "user",
      "attribute": "country",
      "op": "in",
      "negate": false,
      "values": ["USA", "Canada"]
    }]
  }]
}
```

#### addRule

Adds a new targeting rule to the AI Config. The rule may contain `clauses` and serve the variation that `variationId` indicates, or serve a percentage rollout that `rolloutWeights`, `rolloutBucketBy`, and `rolloutContextKind` indicate.

If you set `beforeRuleId`, this adds the new rule before the indicated rule. Otherwise, adds the new rule to the end of the list.

##### Parameters

- `clauses`: Array of clause objects, with `contextKind` (string), `attribute` (string), `op` (string), `negate` (boolean), and `values` (array of strings, numbers, or dates) properties. The `contextKind`, `attribute`, and `values` are case sensitive. The `op` must be lower-case.
- `beforeRuleId`: (Optional) ID of a rule.
- Either
- `variationId`: ID of a variation.

or

- `rolloutWeights`: (Optional) Map of `variationId` to weight, in thousandths of a percent (0-100000).
- `rolloutBucketBy`: (Optional) Context attribute available in the specified `rolloutContextKind`.
- `rolloutContextKind`: (Optional) Context kind, defaults to `user`

Here's an example that uses a `variationId`:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "addRule",
  "variationId": "2f43f67c-3e4e-4945-a18a-26559378ca00",
  "clauses": [{
    "contextKind": "organization",
    "attribute": "located_in",
    "op": "in",
    "negate": false,
    "values": ["Sweden", "Norway"]
  }]
}]
}
```

Here's an example that uses a percentage rollout:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "addRule",
  "clauses": [{
    "contextKind": "organization",
    "attribute": "located_in",
    "op": "in",
    "negate": false,
    "values": ["Sweden", "Norway"]
  }],
  "rolloutContextKind": "organization",
  "rolloutWeights": {
    "2f43f67c-3e4e-4945-a18a-26559378ca00": 15000, // serve 15% this variation
    "e5830889-1ec5-4b0c-9cc9-c48790090c43": 85000  // serve 85% this variation
  }
}]
}
```

#### addTargets

Adds context keys to the individual context targets for the context kind that `contextKind` specifies and the variation that `variationId` specifies. Returns an error if this causes the AI Config to target the same context key in multiple variations.

##### Parameters

- `values`: List of context keys.
- `contextKind`: (Optional) Context kind to target, defaults to `user`
- `variationId`: ID of a variation.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "addTargets",
  "values": ["context-key-123abc", "context-key-456def"],
  "variationId": "2f43f67c-3e4e-4945-a18a-26559378ca00"
}]
}
```

#### addValuesToClause

Adds `values` to the values of the clause that `ruleId` and `clauseId` indicate. Does not update the context kind, attribute, or operator.

##### Parameters

- `ruleId`: ID of a rule in the AI Config.
- `clauseId`: ID of a clause in that rule.
- `values`: Array of strings, case sensitive.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "addValuesToClause",
  "ruleId": "a902ef4a-2faf-4eaf-88e1-ecc356708a29",
  "clauseId": "10a58772-3121-400f-846b-b8a04e8944ed",
  "values": ["beta_testers"]
}]
}
```

#### clearTargets

Removes all individual targets from the variation that `variationId` specifies. This includes both user and non-user targets.

##### Parameters

- `variationId`: ID of a variation.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [ { "kind": "clearTargets", "variationId": "2f43f67c-3e4e-4945-a18a-26559378ca00" } ]
}
```

#### removeClauses

Removes the clauses specified by `clauseIds` from the rule indicated by `ruleId`.

##### Parameters

- `ruleId`: ID of a rule.
- `clauseIds`: Array of IDs of clauses in the rule.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "removeClauses",
  "ruleId": "a902ef4a-2faf-4eaf-88e1-ecc356708a29",
  "clauseIds": ["10a58772-3121-400f-846b-b8a04e8944ed", "36a461dc-235e-4b08-97b9-73ce9365873e"]
}]
}
```

#### removeRule

Removes the targeting rule specified by `ruleId`. Does nothing if the rule does not exist.

##### Parameters

- `ruleId`: ID of a rule.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [ { "kind": "removeRule", "ruleId": "a902ef4a-2faf-4eaf-88e1-ecc356708a29" } ]
}
```

#### removeTargets

Removes context keys from the individual context targets for the context kind that `contextKind` specifies and the variation that `variationId` specifies. Does nothing if the flag does not target the context keys.

##### Parameters

- `values`: List of context keys.
- `contextKind`: (Optional) Context kind to target, defaults to `user`
- `variationId`: ID of a variation.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "removeTargets",
  "values": ["context-key-123abc", "context-key-456def"],
  "variationId": "2f43f67c-3e4e-4945-a18a-26559378ca00"
}]
}
```

#### removeValuesFromClause

Removes `values` from the values of the clause indicated by `ruleId` and `clauseId`. Does not update the context kind, attribute, or operator.

##### Parameters

- `ruleId`: ID of a rule.
- `clauseId`: ID of a clause in that rule.
- `values`: Array of strings, case sensitive.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "removeValuesFromClause",
  "ruleId": "a902ef4a-2faf-4eaf-88e1-ecc356708a29",
  "clauseId": "10a58772-3121-400f-846b-b8a04e8944ed",
  "values": ["beta_testers"]
}]
}
```

#### reorderRules

Rearranges the rules to match the order given in `ruleIds`. Returns an error if `ruleIds` does not match the current set of rules on the AI Config.

##### Parameters

- `ruleIds`: Array of IDs of all rules.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "reorderRules",
  "ruleIds": ["a902ef4a-2faf-4eaf-88e1-ecc356708a29", "63c238d1-835d-435e-8f21-c8d5e40b2a3d"]
}]
}
```

#### replaceRules

Removes all targeting rules for the AI Config and replaces them with the list you provide.

##### Parameters

- `rules`: A list of rules.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [
  {
    "kind": "replaceRules",
    "rules": [
      {
        "variationId": "2f43f67c-3e4e-4945-a18a-26559378ca00",
        "description": "My new rule",
        "clauses": [
          {
            "contextKind": "user",
            "attribute": "segmentMatch",
            "op": "segmentMatch",
            "values": ["test"]
          }
        ]
      }
    ]
  }
]
}
```

#### replaceTargets

Removes all existing targeting and replaces it with the list of targets you provide.

##### Parameters

- `targets`: A list of context targeting. Each item in the list includes an optional `contextKind` that defaults to `user`, a required `variationId`, and a required list of `values`.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [
  {
    "kind": "replaceTargets",
    "targets": [
      {
        "contextKind": "user",
        "variationId": "2f43f67c-3e4e-4945-a18a-26559378ca00",
        "values": ["user-key-123abc"]
      },
      {
        "contextKind": "device",
        "variationId": "e5830889-1ec5-4b0c-9cc9-c48790090c43",
        "values": ["device-key-456def"]
      }
    ]
  }
]
}
```

#### updateClause

Replaces the clause indicated by `ruleId` and `clauseId` with `clause`.

##### Parameters

- `ruleId`: ID of a rule.
- `clauseId`: ID of a clause in that rule.
- `clause`: New `clause` object, with `contextKind` (string), `attribute` (string), `op` (string), `negate` (boolean), and `values` (array of strings, numbers, or dates) properties. The `contextKind`, `attribute`, and `values` are case sensitive. The `op` must be lower-case.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "updateClause",
  "ruleId": "a902ef4a-2faf-4eaf-88e1-ecc356708a29",
  "clauseId": "10c7462a-2062-45ba-a8bb-dfb3de0f8af5",
  "clause": {
    "contextKind": "user",
    "attribute": "country",
    "op": "in",
    "negate": false,
    "values": ["Mexico", "Canada"]
  }
}]
}
```

#### updateDefaultVariation

Updates the default on or off variation of the AI Config.

##### Parameters

- `onVariationValue`: (Optional) The value of the variation of the new on variation.
- `offVariationValue`: (Optional) The value of the variation of the new off variation

Here's an example:

```json
{
"instructions": [ { "kind": "updateDefaultVariation", "OnVariationValue": true, "OffVariationValue": false } ]
}
```

#### updateFallthroughVariationOrRollout

Updates the default or "fallthrough" rule for the AI Config, which the AI Config serves when a context matches none of the targeting rules. The rule can serve either the variation that `variationId` indicates, or a percentage rollout that `rolloutWeights` and `rolloutBucketBy` indicate.

##### Parameters

- `variationId`: ID of a variation.

or

- `rolloutWeights`: Map of `variationId` to weight, in thousandths of a percent (0-100000).
- `rolloutBucketBy`: (Optional) Context attribute available in the specified `rolloutContextKind`.
- `rolloutContextKind`: (Optional) Context kind, defaults to `user`

Here's an example that uses a `variationId`:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "updateFallthroughVariationOrRollout",
  "variationId": "2f43f67c-3e4e-4945-a18a-26559378ca00"
}]
}
```

Here's an example that uses a percentage rollout:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "updateFallthroughVariationOrRollout",
  "rolloutContextKind": "user",
  "rolloutWeights": {
    "2f43f67c-3e4e-4945-a18a-26559378ca00": 15000, // serve 15% this variation
    "e5830889-1ec5-4b0c-9cc9-c48790090c43": 85000  // serve 85% this variation
  }
}]
}
```

#### updateOffVariation

Updates the default off variation to `variationId`. The AI Config serves the default off variation when the AI Config's targeting is **Off**.

##### Parameters

- `variationId`: ID of a variation.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [ { "kind": "updateOffVariation", "variationId": "2f43f67c-3e4e-4945-a18a-26559378ca00" } ]
}
```

#### updateRuleDescription

Updates the description of the targeting rule.

##### Parameters

- `description`: The new human-readable description for this rule.
- `ruleId`: The ID of the rule. You can retrieve this by making a GET request for the AI Config.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "updateRuleDescription",
  "description": "New rule description",
  "ruleId": "a902ef4a-2faf-4eaf-88e1-ecc356708a29"
}]
}
```

#### updateRuleTrackEvents

Updates whether or not LaunchDarkly tracks events for the AI Config associated with this rule.

##### Parameters

- `ruleId`: The ID of the rule. You can retrieve this by making a GET request for the AI Config.
- `trackEvents`: Whether or not events are tracked.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "updateRuleTrackEvents",
  "ruleId": "a902ef4a-2faf-4eaf-88e1-ecc356708a29",
  "trackEvents": true
}]
}
```

#### updateRuleVariationOrRollout

Updates what `ruleId` serves when its clauses evaluate to true. The rule can serve either the variation that `variationId` indicates, or a percent rollout that `rolloutWeights` and `rolloutBucketBy` indicate.

##### Parameters

- `ruleId`: ID of a rule.
- `variationId`: ID of a variation.

or

- `rolloutWeights`: Map of `variationId` to weight, in thousandths of a percent (0-100000).
- `rolloutBucketBy`: (Optional) Context attribute available in the specified `rolloutContextKind`.
- `rolloutContextKind`: (Optional) Context kind, defaults to `user`

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "updateRuleVariationOrRollout",
  "ruleId": "a902ef4a-2faf-4eaf-88e1-ecc356708a29",
  "variationId": "2f43f67c-3e4e-4945-a18a-26559378ca00"
}]
}
```

#### updateTrackEvents

Updates whether or not LaunchDarkly tracks events for the AI Config, for all rules.

##### Parameters

- `trackEvents`: Whether or not events are tracked.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [ { "kind": "updateTrackEvents", "trackEvents": true } ]
}
```

#### updateTrackEventsFallthrough

Updates whether or not LaunchDarkly tracks events for the AI Config, for the default rule.

##### Parameters

- `trackEvents`: Whether or not events are tracked.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [ { "kind": "updateTrackEventsFallthrough", "trackEvents": true } ]
}
```
</details>


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param configKey
	@return ApiPatchAIConfigTargetingRequest
	*/
	PatchAIConfigTargeting(ctx context.Context, projectKey string, configKey string) ApiPatchAIConfigTargetingRequest

	// PatchAIConfigTargetingExecute executes the request
	//  @return AIConfigTargeting
	PatchAIConfigTargetingExecute(r ApiPatchAIConfigTargetingRequest) (*AIConfigTargeting, *http.Response, error)

	/*
	PatchAIConfigVariation Update AI Config variation

	Edit an existing variation of an AI Config. This creates a new version of the variation.

The request body must be a JSON object of the fields to update. The values you include replace the existing values for the fields.

Here's an example:
```
  {
    "messages": [
      {
        "role": "system",
        "content": "The new message"
      }
    ]
  }
```


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param configKey
	@param variationKey
	@return ApiPatchAIConfigVariationRequest
	*/
	PatchAIConfigVariation(ctx context.Context, projectKey string, configKey string, variationKey string) ApiPatchAIConfigVariationRequest

	// PatchAIConfigVariationExecute executes the request
	//  @return AIConfigVariation
	PatchAIConfigVariationExecute(r ApiPatchAIConfigVariationRequest) (*AIConfigVariation, *http.Response, error)

	/*
	PatchAITool Update AI tool

	Edit an existing AI tool.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param toolKey
	@return ApiPatchAIToolRequest
	*/
	PatchAITool(ctx context.Context, projectKey string, toolKey string) ApiPatchAIToolRequest

	// PatchAIToolExecute executes the request
	//  @return AITool
	PatchAIToolExecute(r ApiPatchAIToolRequest) (*AITool, *http.Response, error)

	/*
	PatchAgentGraph Update agent graph

	Edit an existing agent graph.

The request body must be a JSON object of the fields to update. The values you include replace the existing values for the fields.

If the update includes `rootConfigKey` or `edges`, both must be present and will be treated as full replacements.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param graphKey
	@return ApiPatchAgentGraphRequest
	*/
	PatchAgentGraph(ctx context.Context, projectKey string, graphKey string) ApiPatchAgentGraphRequest

	// PatchAgentGraphExecute executes the request
	//  @return AgentGraph
	PatchAgentGraphExecute(r ApiPatchAgentGraphRequest) (*AgentGraph, *http.Response, error)

	/*
	PostAIConfig Create new AI Config

	Create a new AI Config within the given project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiPostAIConfigRequest
	*/
	PostAIConfig(ctx context.Context, projectKey string) ApiPostAIConfigRequest

	// PostAIConfigExecute executes the request
	//  @return AIConfig
	PostAIConfigExecute(r ApiPostAIConfigRequest) (*AIConfig, *http.Response, error)

	/*
	PostAIConfigVariation Create AI Config variation

	Create a new variation for a given AI Config.

The <code>model</code> in the request body requires a <code>modelName</code> and <code>parameters</code>, for example:

```
  "model": {
    "modelName": "claude-3-opus-20240229",
    "parameters": {
      "max_tokens": 1024
    }
  }
```


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param configKey
	@return ApiPostAIConfigVariationRequest
	*/
	PostAIConfigVariation(ctx context.Context, projectKey string, configKey string) ApiPostAIConfigVariationRequest

	// PostAIConfigVariationExecute executes the request
	//  @return AIConfigVariation
	PostAIConfigVariationExecute(r ApiPostAIConfigVariationRequest) (*AIConfigVariation, *http.Response, error)

	/*
	PostAITool Create an AI tool

	Create an AI tool

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiPostAIToolRequest
	*/
	PostAITool(ctx context.Context, projectKey string) ApiPostAIToolRequest

	// PostAIToolExecute executes the request
	//  @return AITool
	PostAIToolExecute(r ApiPostAIToolRequest) (*AITool, *http.Response, error)

	/*
	PostAgentGraph Create new agent graph

	Create a new agent graph within the given project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiPostAgentGraphRequest
	*/
	PostAgentGraph(ctx context.Context, projectKey string) ApiPostAgentGraphRequest

	// PostAgentGraphExecute executes the request
	//  @return AgentGraph
	PostAgentGraphExecute(r ApiPostAgentGraphRequest) (*AgentGraph, *http.Response, error)

	/*
	PostModelConfig Create an AI model config

	Create an AI model config. You can use this in any variation for any AI Config in your project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiPostModelConfigRequest
	*/
	PostModelConfig(ctx context.Context, projectKey string) ApiPostModelConfigRequest

	// PostModelConfigExecute executes the request
	//  @return ModelConfig
	PostModelConfigExecute(r ApiPostModelConfigRequest) (*ModelConfig, *http.Response, error)

	/*
	PostRestrictedModels Add AI models to the restricted list

	Add AI models, by key, to the restricted list. Keys are included in the response from the [List AI model configs](https://launchdarkly.com/docs/api/ai-configs-beta/list-model-configs) endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiPostRestrictedModelsRequest
	*/
	PostRestrictedModels(ctx context.Context, projectKey string) ApiPostRestrictedModelsRequest

	// PostRestrictedModelsExecute executes the request
	//  @return RestrictedModelsResponse
	PostRestrictedModelsExecute(r ApiPostRestrictedModelsRequest) (*RestrictedModelsResponse, *http.Response, error)
}

// AIConfigsBetaApiService AIConfigsBetaApi service
type AIConfigsBetaApiService service

type ApiDeleteAIConfigRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	configKey string
}

// Version of the endpoint.
func (r ApiDeleteAIConfigRequest) LDAPIVersion(lDAPIVersion string) ApiDeleteAIConfigRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

func (r ApiDeleteAIConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAIConfigExecute(r)
}

/*
DeleteAIConfig Delete AI Config

Delete an existing AI Config.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param configKey
 @return ApiDeleteAIConfigRequest
*/
func (a *AIConfigsBetaApiService) DeleteAIConfig(ctx context.Context, projectKey string, configKey string) ApiDeleteAIConfigRequest {
	return ApiDeleteAIConfigRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
	}
}

// Execute executes the request
func (a *AIConfigsBetaApiService) DeleteAIConfigExecute(r ApiDeleteAIConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.DeleteAIConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return nil, reportError("lDAPIVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteAIConfigVariationRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	configKey string
	variationKey string
}

// Version of the endpoint.
func (r ApiDeleteAIConfigVariationRequest) LDAPIVersion(lDAPIVersion string) ApiDeleteAIConfigVariationRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

func (r ApiDeleteAIConfigVariationRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAIConfigVariationExecute(r)
}

/*
DeleteAIConfigVariation Delete AI Config variation

Delete a specific variation of an AI Config by config key and variation key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param configKey
 @param variationKey
 @return ApiDeleteAIConfigVariationRequest
*/
func (a *AIConfigsBetaApiService) DeleteAIConfigVariation(ctx context.Context, projectKey string, configKey string, variationKey string) ApiDeleteAIConfigVariationRequest {
	return ApiDeleteAIConfigVariationRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
		variationKey: variationKey,
	}
}

// Execute executes the request
func (a *AIConfigsBetaApiService) DeleteAIConfigVariationExecute(r ApiDeleteAIConfigVariationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.DeleteAIConfigVariation")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}/variations/{variationKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"variationKey"+"}", url.PathEscape(parameterValueToString(r.variationKey, "variationKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return nil, reportError("lDAPIVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteAIToolRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	toolKey string
}

// Version of the endpoint.
func (r ApiDeleteAIToolRequest) LDAPIVersion(lDAPIVersion string) ApiDeleteAIToolRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

func (r ApiDeleteAIToolRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAIToolExecute(r)
}

/*
DeleteAITool Delete AI tool

Delete an existing AI tool.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param toolKey
 @return ApiDeleteAIToolRequest
*/
func (a *AIConfigsBetaApiService) DeleteAITool(ctx context.Context, projectKey string, toolKey string) ApiDeleteAIToolRequest {
	return ApiDeleteAIToolRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		toolKey: toolKey,
	}
}

// Execute executes the request
func (a *AIConfigsBetaApiService) DeleteAIToolExecute(r ApiDeleteAIToolRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.DeleteAITool")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-tools/{toolKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toolKey"+"}", url.PathEscape(parameterValueToString(r.toolKey, "toolKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return nil, reportError("lDAPIVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteAgentGraphRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	graphKey string
}

// Version of the endpoint.
func (r ApiDeleteAgentGraphRequest) LDAPIVersion(lDAPIVersion string) ApiDeleteAgentGraphRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

func (r ApiDeleteAgentGraphRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAgentGraphExecute(r)
}

/*
DeleteAgentGraph Delete agent graph

Delete an existing agent graph and all of its edges.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param graphKey
 @return ApiDeleteAgentGraphRequest
*/
func (a *AIConfigsBetaApiService) DeleteAgentGraph(ctx context.Context, projectKey string, graphKey string) ApiDeleteAgentGraphRequest {
	return ApiDeleteAgentGraphRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		graphKey: graphKey,
	}
}

// Execute executes the request
func (a *AIConfigsBetaApiService) DeleteAgentGraphExecute(r ApiDeleteAgentGraphRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.DeleteAgentGraph")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/agent-graphs/{graphKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"graphKey"+"}", url.PathEscape(parameterValueToString(r.graphKey, "graphKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return nil, reportError("lDAPIVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteModelConfigRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	modelConfigKey string
}

// Version of the endpoint.
func (r ApiDeleteModelConfigRequest) LDAPIVersion(lDAPIVersion string) ApiDeleteModelConfigRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

func (r ApiDeleteModelConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteModelConfigExecute(r)
}

/*
DeleteModelConfig Delete an AI model config

Delete an AI model config.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param modelConfigKey
 @return ApiDeleteModelConfigRequest
*/
func (a *AIConfigsBetaApiService) DeleteModelConfig(ctx context.Context, projectKey string, modelConfigKey string) ApiDeleteModelConfigRequest {
	return ApiDeleteModelConfigRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		modelConfigKey: modelConfigKey,
	}
}

// Execute executes the request
func (a *AIConfigsBetaApiService) DeleteModelConfigExecute(r ApiDeleteModelConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.DeleteModelConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/model-configs/{modelConfigKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"modelConfigKey"+"}", url.PathEscape(parameterValueToString(r.modelConfigKey, "modelConfigKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return nil, reportError("lDAPIVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteRestrictedModelsRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	restrictedModelsRequest *RestrictedModelsRequest
}

// Version of the endpoint.
func (r ApiDeleteRestrictedModelsRequest) LDAPIVersion(lDAPIVersion string) ApiDeleteRestrictedModelsRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// List of AI model keys to remove from the restricted list
func (r ApiDeleteRestrictedModelsRequest) RestrictedModelsRequest(restrictedModelsRequest RestrictedModelsRequest) ApiDeleteRestrictedModelsRequest {
	r.restrictedModelsRequest = &restrictedModelsRequest
	return r
}

func (r ApiDeleteRestrictedModelsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteRestrictedModelsExecute(r)
}

/*
DeleteRestrictedModels Remove AI models from the restricted list

Remove AI models, by key, from the restricted list.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiDeleteRestrictedModelsRequest
*/
func (a *AIConfigsBetaApiService) DeleteRestrictedModels(ctx context.Context, projectKey string) ApiDeleteRestrictedModelsRequest {
	return ApiDeleteRestrictedModelsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
func (a *AIConfigsBetaApiService) DeleteRestrictedModelsExecute(r ApiDeleteRestrictedModelsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.DeleteRestrictedModels")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/model-configs/restricted"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return nil, reportError("lDAPIVersion is required and must be specified")
	}
	if r.restrictedModelsRequest == nil {
		return nil, reportError("restrictedModelsRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	// body params
	localVarPostBody = r.restrictedModelsRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetAIConfigRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	configKey string
}

// Version of the endpoint.
func (r ApiGetAIConfigRequest) LDAPIVersion(lDAPIVersion string) ApiGetAIConfigRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

func (r ApiGetAIConfigRequest) Execute() (*AIConfig, *http.Response, error) {
	return r.ApiService.GetAIConfigExecute(r)
}

/*
GetAIConfig Get AI Config

Retrieve a specific AI Config by its key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param configKey
 @return ApiGetAIConfigRequest
*/
func (a *AIConfigsBetaApiService) GetAIConfig(ctx context.Context, projectKey string, configKey string) ApiGetAIConfigRequest {
	return ApiGetAIConfigRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
	}
}

// Execute executes the request
//  @return AIConfig
func (a *AIConfigsBetaApiService) GetAIConfigExecute(r ApiGetAIConfigRequest) (*AIConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AIConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.GetAIConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAIConfigMetricsRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	configKey string
	from *int32
	to *int32
	env *string
}

// Version of the endpoint.
func (r ApiGetAIConfigMetricsRequest) LDAPIVersion(lDAPIVersion string) ApiGetAIConfigMetricsRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// The starting time, as milliseconds since epoch (inclusive).
func (r ApiGetAIConfigMetricsRequest) From(from int32) ApiGetAIConfigMetricsRequest {
	r.from = &from
	return r
}

// The ending time, as milliseconds since epoch (exclusive). May not be more than 100 days after &#x60;from&#x60;.
func (r ApiGetAIConfigMetricsRequest) To(to int32) ApiGetAIConfigMetricsRequest {
	r.to = &to
	return r
}

// An environment key. Only metrics from this environment will be included.
func (r ApiGetAIConfigMetricsRequest) Env(env string) ApiGetAIConfigMetricsRequest {
	r.env = &env
	return r
}

func (r ApiGetAIConfigMetricsRequest) Execute() (*Metrics, *http.Response, error) {
	return r.ApiService.GetAIConfigMetricsExecute(r)
}

/*
GetAIConfigMetrics Get AI Config metrics

Retrieve usage metrics for an AI Config by config key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param configKey
 @return ApiGetAIConfigMetricsRequest
*/
func (a *AIConfigsBetaApiService) GetAIConfigMetrics(ctx context.Context, projectKey string, configKey string) ApiGetAIConfigMetricsRequest {
	return ApiGetAIConfigMetricsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
	}
}

// Execute executes the request
//  @return Metrics
func (a *AIConfigsBetaApiService) GetAIConfigMetricsExecute(r ApiGetAIConfigMetricsRequest) (*Metrics, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Metrics
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.GetAIConfigMetrics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}/metrics"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}
	if r.from == nil {
		return localVarReturnValue, nil, reportError("from is required and must be specified")
	}
	if r.to == nil {
		return localVarReturnValue, nil, reportError("to is required and must be specified")
	}
	if r.env == nil {
		return localVarReturnValue, nil, reportError("env is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "to", r.to, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "env", r.env, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAIConfigMetricsByVariationRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	configKey string
	from *int32
	to *int32
	env *string
}

// Version of the endpoint.
func (r ApiGetAIConfigMetricsByVariationRequest) LDAPIVersion(lDAPIVersion string) ApiGetAIConfigMetricsByVariationRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// The starting time, as milliseconds since epoch (inclusive).
func (r ApiGetAIConfigMetricsByVariationRequest) From(from int32) ApiGetAIConfigMetricsByVariationRequest {
	r.from = &from
	return r
}

// The ending time, as milliseconds since epoch (exclusive). May not be more than 100 days after &#x60;from&#x60;.
func (r ApiGetAIConfigMetricsByVariationRequest) To(to int32) ApiGetAIConfigMetricsByVariationRequest {
	r.to = &to
	return r
}

// An environment key. Only metrics from this environment will be included.
func (r ApiGetAIConfigMetricsByVariationRequest) Env(env string) ApiGetAIConfigMetricsByVariationRequest {
	r.env = &env
	return r
}

func (r ApiGetAIConfigMetricsByVariationRequest) Execute() ([]MetricByVariation, *http.Response, error) {
	return r.ApiService.GetAIConfigMetricsByVariationExecute(r)
}

/*
GetAIConfigMetricsByVariation Get AI Config metrics by variation

Retrieve usage metrics for an AI Config by config key, with results split by variation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param configKey
 @return ApiGetAIConfigMetricsByVariationRequest
*/
func (a *AIConfigsBetaApiService) GetAIConfigMetricsByVariation(ctx context.Context, projectKey string, configKey string) ApiGetAIConfigMetricsByVariationRequest {
	return ApiGetAIConfigMetricsByVariationRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
	}
}

// Execute executes the request
//  @return []MetricByVariation
func (a *AIConfigsBetaApiService) GetAIConfigMetricsByVariationExecute(r ApiGetAIConfigMetricsByVariationRequest) ([]MetricByVariation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []MetricByVariation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.GetAIConfigMetricsByVariation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}/metrics-by-variation"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}
	if r.from == nil {
		return localVarReturnValue, nil, reportError("from is required and must be specified")
	}
	if r.to == nil {
		return localVarReturnValue, nil, reportError("to is required and must be specified")
	}
	if r.env == nil {
		return localVarReturnValue, nil, reportError("env is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "to", r.to, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "env", r.env, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAIConfigTargetingRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	configKey string
}

// Version of the endpoint.
func (r ApiGetAIConfigTargetingRequest) LDAPIVersion(lDAPIVersion string) ApiGetAIConfigTargetingRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

func (r ApiGetAIConfigTargetingRequest) Execute() (*AIConfigTargeting, *http.Response, error) {
	return r.ApiService.GetAIConfigTargetingExecute(r)
}

/*
GetAIConfigTargeting Show an AI Config's targeting

Retrieves a specific AI Config's targeting by its key

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param configKey
 @return ApiGetAIConfigTargetingRequest
*/
func (a *AIConfigsBetaApiService) GetAIConfigTargeting(ctx context.Context, projectKey string, configKey string) ApiGetAIConfigTargetingRequest {
	return ApiGetAIConfigTargetingRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
	}
}

// Execute executes the request
//  @return AIConfigTargeting
func (a *AIConfigsBetaApiService) GetAIConfigTargetingExecute(r ApiGetAIConfigTargetingRequest) (*AIConfigTargeting, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AIConfigTargeting
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.GetAIConfigTargeting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}/targeting"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAIConfigVariationRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	configKey string
	variationKey string
}

// Version of the endpoint.
func (r ApiGetAIConfigVariationRequest) LDAPIVersion(lDAPIVersion string) ApiGetAIConfigVariationRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

func (r ApiGetAIConfigVariationRequest) Execute() (*AIConfigVariationsResponse, *http.Response, error) {
	return r.ApiService.GetAIConfigVariationExecute(r)
}

/*
GetAIConfigVariation Get AI Config variation

Get an AI Config variation by key. The response includes all variation versions for the given variation key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param configKey
 @param variationKey
 @return ApiGetAIConfigVariationRequest
*/
func (a *AIConfigsBetaApiService) GetAIConfigVariation(ctx context.Context, projectKey string, configKey string, variationKey string) ApiGetAIConfigVariationRequest {
	return ApiGetAIConfigVariationRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
		variationKey: variationKey,
	}
}

// Execute executes the request
//  @return AIConfigVariationsResponse
func (a *AIConfigsBetaApiService) GetAIConfigVariationExecute(r ApiGetAIConfigVariationRequest) (*AIConfigVariationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AIConfigVariationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.GetAIConfigVariation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}/variations/{variationKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"variationKey"+"}", url.PathEscape(parameterValueToString(r.variationKey, "variationKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAIConfigsRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	sort *string
	limit *int32
	offset *int32
	filter *string
}

// Version of the endpoint.
func (r ApiGetAIConfigsRequest) LDAPIVersion(lDAPIVersion string) ApiGetAIConfigsRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// A sort to apply to the list of AI Configs.
func (r ApiGetAIConfigsRequest) Sort(sort string) ApiGetAIConfigsRequest {
	r.sort = &sort
	return r
}

// The number of AI Configs to return.
func (r ApiGetAIConfigsRequest) Limit(limit int32) ApiGetAIConfigsRequest {
	r.limit = &limit
	return r
}

// Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;.
func (r ApiGetAIConfigsRequest) Offset(offset int32) ApiGetAIConfigsRequest {
	r.offset = &offset
	return r
}

// A filter to apply to the list of AI Configs.
func (r ApiGetAIConfigsRequest) Filter(filter string) ApiGetAIConfigsRequest {
	r.filter = &filter
	return r
}

func (r ApiGetAIConfigsRequest) Execute() (*AIConfigs, *http.Response, error) {
	return r.ApiService.GetAIConfigsExecute(r)
}

/*
GetAIConfigs List AI Configs

Get a list of all AI Configs in the given project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiGetAIConfigsRequest
*/
func (a *AIConfigsBetaApiService) GetAIConfigs(ctx context.Context, projectKey string) ApiGetAIConfigsRequest {
	return ApiGetAIConfigsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return AIConfigs
func (a *AIConfigsBetaApiService) GetAIConfigsExecute(r ApiGetAIConfigsRequest) (*AIConfigs, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AIConfigs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.GetAIConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}

	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAIToolRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	toolKey string
}

// Version of the endpoint.
func (r ApiGetAIToolRequest) LDAPIVersion(lDAPIVersion string) ApiGetAIToolRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

func (r ApiGetAIToolRequest) Execute() (*AITool, *http.Response, error) {
	return r.ApiService.GetAIToolExecute(r)
}

/*
GetAITool Get AI tool

Retrieve a specific AI tool by its key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param toolKey
 @return ApiGetAIToolRequest
*/
func (a *AIConfigsBetaApiService) GetAITool(ctx context.Context, projectKey string, toolKey string) ApiGetAIToolRequest {
	return ApiGetAIToolRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		toolKey: toolKey,
	}
}

// Execute executes the request
//  @return AITool
func (a *AIConfigsBetaApiService) GetAIToolExecute(r ApiGetAIToolRequest) (*AITool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AITool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.GetAITool")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-tools/{toolKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toolKey"+"}", url.PathEscape(parameterValueToString(r.toolKey, "toolKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAgentGraphRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	graphKey string
}

// Version of the endpoint.
func (r ApiGetAgentGraphRequest) LDAPIVersion(lDAPIVersion string) ApiGetAgentGraphRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

func (r ApiGetAgentGraphRequest) Execute() (*AgentGraph, *http.Response, error) {
	return r.ApiService.GetAgentGraphExecute(r)
}

/*
GetAgentGraph Get agent graph

Retrieve a specific agent graph by its key, including its edges.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param graphKey
 @return ApiGetAgentGraphRequest
*/
func (a *AIConfigsBetaApiService) GetAgentGraph(ctx context.Context, projectKey string, graphKey string) ApiGetAgentGraphRequest {
	return ApiGetAgentGraphRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		graphKey: graphKey,
	}
}

// Execute executes the request
//  @return AgentGraph
func (a *AIConfigsBetaApiService) GetAgentGraphExecute(r ApiGetAgentGraphRequest) (*AgentGraph, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentGraph
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.GetAgentGraph")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/agent-graphs/{graphKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"graphKey"+"}", url.PathEscape(parameterValueToString(r.graphKey, "graphKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetModelConfigRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	modelConfigKey string
}

// Version of the endpoint.
func (r ApiGetModelConfigRequest) LDAPIVersion(lDAPIVersion string) ApiGetModelConfigRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

func (r ApiGetModelConfigRequest) Execute() (*ModelConfig, *http.Response, error) {
	return r.ApiService.GetModelConfigExecute(r)
}

/*
GetModelConfig Get AI model config

Get an AI model config by key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param modelConfigKey
 @return ApiGetModelConfigRequest
*/
func (a *AIConfigsBetaApiService) GetModelConfig(ctx context.Context, projectKey string, modelConfigKey string) ApiGetModelConfigRequest {
	return ApiGetModelConfigRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		modelConfigKey: modelConfigKey,
	}
}

// Execute executes the request
//  @return ModelConfig
func (a *AIConfigsBetaApiService) GetModelConfigExecute(r ApiGetModelConfigRequest) (*ModelConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.GetModelConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/model-configs/{modelConfigKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"modelConfigKey"+"}", url.PathEscape(parameterValueToString(r.modelConfigKey, "modelConfigKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListAIToolVersionsRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	toolKey string
	sort *string
	limit *int32
	offset *int32
}

// Version of the endpoint.
func (r ApiListAIToolVersionsRequest) LDAPIVersion(lDAPIVersion string) ApiListAIToolVersionsRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// A sort to apply to the list of AI Configs.
func (r ApiListAIToolVersionsRequest) Sort(sort string) ApiListAIToolVersionsRequest {
	r.sort = &sort
	return r
}

// The number of AI Configs to return.
func (r ApiListAIToolVersionsRequest) Limit(limit int32) ApiListAIToolVersionsRequest {
	r.limit = &limit
	return r
}

// Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;.
func (r ApiListAIToolVersionsRequest) Offset(offset int32) ApiListAIToolVersionsRequest {
	r.offset = &offset
	return r
}

func (r ApiListAIToolVersionsRequest) Execute() (*AITools, *http.Response, error) {
	return r.ApiService.ListAIToolVersionsExecute(r)
}

/*
ListAIToolVersions List AI tool versions

Get a list of all versions of an AI tool in the given project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param toolKey
 @return ApiListAIToolVersionsRequest
*/
func (a *AIConfigsBetaApiService) ListAIToolVersions(ctx context.Context, projectKey string, toolKey string) ApiListAIToolVersionsRequest {
	return ApiListAIToolVersionsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		toolKey: toolKey,
	}
}

// Execute executes the request
//  @return AITools
func (a *AIConfigsBetaApiService) ListAIToolVersionsExecute(r ApiListAIToolVersionsRequest) (*AITools, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AITools
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.ListAIToolVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-tools/{toolKey}/versions"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toolKey"+"}", url.PathEscape(parameterValueToString(r.toolKey, "toolKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}

	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListAIToolsRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	sort *string
	limit *int32
	offset *int32
	filter *string
}

// Version of the endpoint.
func (r ApiListAIToolsRequest) LDAPIVersion(lDAPIVersion string) ApiListAIToolsRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// A sort to apply to the list of AI Configs.
func (r ApiListAIToolsRequest) Sort(sort string) ApiListAIToolsRequest {
	r.sort = &sort
	return r
}

// The number of AI Configs to return.
func (r ApiListAIToolsRequest) Limit(limit int32) ApiListAIToolsRequest {
	r.limit = &limit
	return r
}

// Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;.
func (r ApiListAIToolsRequest) Offset(offset int32) ApiListAIToolsRequest {
	r.offset = &offset
	return r
}

// A filter to apply to the list of AI Configs.
func (r ApiListAIToolsRequest) Filter(filter string) ApiListAIToolsRequest {
	r.filter = &filter
	return r
}

func (r ApiListAIToolsRequest) Execute() (*AITools, *http.Response, error) {
	return r.ApiService.ListAIToolsExecute(r)
}

/*
ListAITools List AI tools

Get a list of all AI tools in the given project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiListAIToolsRequest
*/
func (a *AIConfigsBetaApiService) ListAITools(ctx context.Context, projectKey string) ApiListAIToolsRequest {
	return ApiListAIToolsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return AITools
func (a *AIConfigsBetaApiService) ListAIToolsExecute(r ApiListAIToolsRequest) (*AITools, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AITools
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.ListAITools")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-tools"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}

	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListAgentGraphsRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	limit *int32
	offset *int32
}

// Version of the endpoint.
func (r ApiListAgentGraphsRequest) LDAPIVersion(lDAPIVersion string) ApiListAgentGraphsRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// The number of AI Configs to return.
func (r ApiListAgentGraphsRequest) Limit(limit int32) ApiListAgentGraphsRequest {
	r.limit = &limit
	return r
}

// Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;.
func (r ApiListAgentGraphsRequest) Offset(offset int32) ApiListAgentGraphsRequest {
	r.offset = &offset
	return r
}

func (r ApiListAgentGraphsRequest) Execute() (*AgentGraphs, *http.Response, error) {
	return r.ApiService.ListAgentGraphsExecute(r)
}

/*
ListAgentGraphs List agent graphs

Get a list of all agent graphs in the given project. Returns metadata only, without edge data.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiListAgentGraphsRequest
*/
func (a *AIConfigsBetaApiService) ListAgentGraphs(ctx context.Context, projectKey string) ApiListAgentGraphsRequest {
	return ApiListAgentGraphsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return AgentGraphs
func (a *AIConfigsBetaApiService) ListAgentGraphsExecute(r ApiListAgentGraphsRequest) (*AgentGraphs, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentGraphs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.ListAgentGraphs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/agent-graphs"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListModelConfigsRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	restricted *bool
}

// Version of the endpoint.
func (r ApiListModelConfigsRequest) LDAPIVersion(lDAPIVersion string) ApiListModelConfigsRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// Whether to return only restricted models
func (r ApiListModelConfigsRequest) Restricted(restricted bool) ApiListModelConfigsRequest {
	r.restricted = &restricted
	return r
}

func (r ApiListModelConfigsRequest) Execute() ([]ModelConfig, *http.Response, error) {
	return r.ApiService.ListModelConfigsExecute(r)
}

/*
ListModelConfigs List AI model configs

Get all AI model configs for a project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiListModelConfigsRequest
*/
func (a *AIConfigsBetaApiService) ListModelConfigs(ctx context.Context, projectKey string) ApiListModelConfigsRequest {
	return ApiListModelConfigsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return []ModelConfig
func (a *AIConfigsBetaApiService) ListModelConfigsExecute(r ApiListModelConfigsRequest) ([]ModelConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ModelConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.ListModelConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/model-configs"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}

	if r.restricted != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "restricted", r.restricted, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPatchAIConfigRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	configKey string
	aIConfigPatch *AIConfigPatch
}

// Version of the endpoint.
func (r ApiPatchAIConfigRequest) LDAPIVersion(lDAPIVersion string) ApiPatchAIConfigRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// AI Config object to update
func (r ApiPatchAIConfigRequest) AIConfigPatch(aIConfigPatch AIConfigPatch) ApiPatchAIConfigRequest {
	r.aIConfigPatch = &aIConfigPatch
	return r
}

func (r ApiPatchAIConfigRequest) Execute() (*AIConfig, *http.Response, error) {
	return r.ApiService.PatchAIConfigExecute(r)
}

/*
PatchAIConfig Update AI Config

Edit an existing AI Config.

The request body must be a JSON object of the fields to update. The values you include replace the existing values for the fields.

Here's an example:
  ```
    {
      "description": "Example updated description",
      "tags": ["new-tag"]
    }
  ```


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param configKey
 @return ApiPatchAIConfigRequest
*/
func (a *AIConfigsBetaApiService) PatchAIConfig(ctx context.Context, projectKey string, configKey string) ApiPatchAIConfigRequest {
	return ApiPatchAIConfigRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
	}
}

// Execute executes the request
//  @return AIConfig
func (a *AIConfigsBetaApiService) PatchAIConfigExecute(r ApiPatchAIConfigRequest) (*AIConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AIConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.PatchAIConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	// body params
	localVarPostBody = r.aIConfigPatch
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPatchAIConfigTargetingRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	configKey string
	aIConfigTargetingPatch *AIConfigTargetingPatch
}

// Version of the endpoint.
func (r ApiPatchAIConfigTargetingRequest) LDAPIVersion(lDAPIVersion string) ApiPatchAIConfigTargetingRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// AI Config targeting semantic patch instructions
func (r ApiPatchAIConfigTargetingRequest) AIConfigTargetingPatch(aIConfigTargetingPatch AIConfigTargetingPatch) ApiPatchAIConfigTargetingRequest {
	r.aIConfigTargetingPatch = &aIConfigTargetingPatch
	return r
}

func (r ApiPatchAIConfigTargetingRequest) Execute() (*AIConfigTargeting, *http.Response, error) {
	return r.ApiService.PatchAIConfigTargetingExecute(r)
}

/*
PatchAIConfigTargeting Update AI Config targeting

Perform a partial update to an AI Config's targeting. The request body must be a valid semantic patch.

### Using semantic patches on an AI Config

To make a semantic patch request, you must append `domain-model=launchdarkly.semanticpatch` to your `Content-Type` header. To learn more, read [Updates using semantic patch](https://launchdarkly.com/docs/api#updates-using-semantic-patch).

The body of a semantic patch request for updating an AI Config's targeting takes the following properties:

* `comment` (string): (Optional) A description of the update.
* `environmentKey` (string): The key of the LaunchDarkly environment.
* `instructions` (array): (Required) A list of actions the update should perform. Each action in the list must be an object with a `kind` property that indicates the instruction. If the action requires parameters, you must include those parameters as additional fields in the object. The body of a single semantic patch can contain many different instructions.

### Instructions

Semantic patch requests support the following `kind` instructions for updating AI Configs.

<details>
<summary>Click to expand instructions for <strong>working with targeting and variations</strong> for AI Configs</summary>

#### addClauses

Adds the given clauses to the rule indicated by `ruleId`.

##### Parameters

- `ruleId`: ID of a rule in the AI Config.
- `clauses`: Array of clause objects, with `contextKind` (string), `attribute` (string), `op` (string), `negate` (boolean), and `values` (array of strings, numbers, or dates) properties. The `contextKind`, `attribute`, and `values` are case sensitive. The `op` must be lower-case.

Here's an example:

```json
{
  "environmentKey": "environment-key-123abc",
  "instructions": [{
    "kind": "addClauses",
    "ruleId": "a902ef4a-2faf-4eaf-88e1-ecc356708a29",
    "clauses": [{
      "contextKind": "user",
      "attribute": "country",
      "op": "in",
      "negate": false,
      "values": ["USA", "Canada"]
    }]
  }]
}
```

#### addRule

Adds a new targeting rule to the AI Config. The rule may contain `clauses` and serve the variation that `variationId` indicates, or serve a percentage rollout that `rolloutWeights`, `rolloutBucketBy`, and `rolloutContextKind` indicate.

If you set `beforeRuleId`, this adds the new rule before the indicated rule. Otherwise, adds the new rule to the end of the list.

##### Parameters

- `clauses`: Array of clause objects, with `contextKind` (string), `attribute` (string), `op` (string), `negate` (boolean), and `values` (array of strings, numbers, or dates) properties. The `contextKind`, `attribute`, and `values` are case sensitive. The `op` must be lower-case.
- `beforeRuleId`: (Optional) ID of a rule.
- Either
- `variationId`: ID of a variation.

or

- `rolloutWeights`: (Optional) Map of `variationId` to weight, in thousandths of a percent (0-100000).
- `rolloutBucketBy`: (Optional) Context attribute available in the specified `rolloutContextKind`.
- `rolloutContextKind`: (Optional) Context kind, defaults to `user`

Here's an example that uses a `variationId`:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "addRule",
  "variationId": "2f43f67c-3e4e-4945-a18a-26559378ca00",
  "clauses": [{
    "contextKind": "organization",
    "attribute": "located_in",
    "op": "in",
    "negate": false,
    "values": ["Sweden", "Norway"]
  }]
}]
}
```

Here's an example that uses a percentage rollout:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "addRule",
  "clauses": [{
    "contextKind": "organization",
    "attribute": "located_in",
    "op": "in",
    "negate": false,
    "values": ["Sweden", "Norway"]
  }],
  "rolloutContextKind": "organization",
  "rolloutWeights": {
    "2f43f67c-3e4e-4945-a18a-26559378ca00": 15000, // serve 15% this variation
    "e5830889-1ec5-4b0c-9cc9-c48790090c43": 85000  // serve 85% this variation
  }
}]
}
```

#### addTargets

Adds context keys to the individual context targets for the context kind that `contextKind` specifies and the variation that `variationId` specifies. Returns an error if this causes the AI Config to target the same context key in multiple variations.

##### Parameters

- `values`: List of context keys.
- `contextKind`: (Optional) Context kind to target, defaults to `user`
- `variationId`: ID of a variation.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "addTargets",
  "values": ["context-key-123abc", "context-key-456def"],
  "variationId": "2f43f67c-3e4e-4945-a18a-26559378ca00"
}]
}
```

#### addValuesToClause

Adds `values` to the values of the clause that `ruleId` and `clauseId` indicate. Does not update the context kind, attribute, or operator.

##### Parameters

- `ruleId`: ID of a rule in the AI Config.
- `clauseId`: ID of a clause in that rule.
- `values`: Array of strings, case sensitive.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "addValuesToClause",
  "ruleId": "a902ef4a-2faf-4eaf-88e1-ecc356708a29",
  "clauseId": "10a58772-3121-400f-846b-b8a04e8944ed",
  "values": ["beta_testers"]
}]
}
```

#### clearTargets

Removes all individual targets from the variation that `variationId` specifies. This includes both user and non-user targets.

##### Parameters

- `variationId`: ID of a variation.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [ { "kind": "clearTargets", "variationId": "2f43f67c-3e4e-4945-a18a-26559378ca00" } ]
}
```

#### removeClauses

Removes the clauses specified by `clauseIds` from the rule indicated by `ruleId`.

##### Parameters

- `ruleId`: ID of a rule.
- `clauseIds`: Array of IDs of clauses in the rule.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "removeClauses",
  "ruleId": "a902ef4a-2faf-4eaf-88e1-ecc356708a29",
  "clauseIds": ["10a58772-3121-400f-846b-b8a04e8944ed", "36a461dc-235e-4b08-97b9-73ce9365873e"]
}]
}
```

#### removeRule

Removes the targeting rule specified by `ruleId`. Does nothing if the rule does not exist.

##### Parameters

- `ruleId`: ID of a rule.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [ { "kind": "removeRule", "ruleId": "a902ef4a-2faf-4eaf-88e1-ecc356708a29" } ]
}
```

#### removeTargets

Removes context keys from the individual context targets for the context kind that `contextKind` specifies and the variation that `variationId` specifies. Does nothing if the flag does not target the context keys.

##### Parameters

- `values`: List of context keys.
- `contextKind`: (Optional) Context kind to target, defaults to `user`
- `variationId`: ID of a variation.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "removeTargets",
  "values": ["context-key-123abc", "context-key-456def"],
  "variationId": "2f43f67c-3e4e-4945-a18a-26559378ca00"
}]
}
```

#### removeValuesFromClause

Removes `values` from the values of the clause indicated by `ruleId` and `clauseId`. Does not update the context kind, attribute, or operator.

##### Parameters

- `ruleId`: ID of a rule.
- `clauseId`: ID of a clause in that rule.
- `values`: Array of strings, case sensitive.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "removeValuesFromClause",
  "ruleId": "a902ef4a-2faf-4eaf-88e1-ecc356708a29",
  "clauseId": "10a58772-3121-400f-846b-b8a04e8944ed",
  "values": ["beta_testers"]
}]
}
```

#### reorderRules

Rearranges the rules to match the order given in `ruleIds`. Returns an error if `ruleIds` does not match the current set of rules on the AI Config.

##### Parameters

- `ruleIds`: Array of IDs of all rules.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "reorderRules",
  "ruleIds": ["a902ef4a-2faf-4eaf-88e1-ecc356708a29", "63c238d1-835d-435e-8f21-c8d5e40b2a3d"]
}]
}
```

#### replaceRules

Removes all targeting rules for the AI Config and replaces them with the list you provide.

##### Parameters

- `rules`: A list of rules.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [
  {
    "kind": "replaceRules",
    "rules": [
      {
        "variationId": "2f43f67c-3e4e-4945-a18a-26559378ca00",
        "description": "My new rule",
        "clauses": [
          {
            "contextKind": "user",
            "attribute": "segmentMatch",
            "op": "segmentMatch",
            "values": ["test"]
          }
        ]
      }
    ]
  }
]
}
```

#### replaceTargets

Removes all existing targeting and replaces it with the list of targets you provide.

##### Parameters

- `targets`: A list of context targeting. Each item in the list includes an optional `contextKind` that defaults to `user`, a required `variationId`, and a required list of `values`.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [
  {
    "kind": "replaceTargets",
    "targets": [
      {
        "contextKind": "user",
        "variationId": "2f43f67c-3e4e-4945-a18a-26559378ca00",
        "values": ["user-key-123abc"]
      },
      {
        "contextKind": "device",
        "variationId": "e5830889-1ec5-4b0c-9cc9-c48790090c43",
        "values": ["device-key-456def"]
      }
    ]
  }
]
}
```

#### updateClause

Replaces the clause indicated by `ruleId` and `clauseId` with `clause`.

##### Parameters

- `ruleId`: ID of a rule.
- `clauseId`: ID of a clause in that rule.
- `clause`: New `clause` object, with `contextKind` (string), `attribute` (string), `op` (string), `negate` (boolean), and `values` (array of strings, numbers, or dates) properties. The `contextKind`, `attribute`, and `values` are case sensitive. The `op` must be lower-case.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "updateClause",
  "ruleId": "a902ef4a-2faf-4eaf-88e1-ecc356708a29",
  "clauseId": "10c7462a-2062-45ba-a8bb-dfb3de0f8af5",
  "clause": {
    "contextKind": "user",
    "attribute": "country",
    "op": "in",
    "negate": false,
    "values": ["Mexico", "Canada"]
  }
}]
}
```

#### updateDefaultVariation

Updates the default on or off variation of the AI Config.

##### Parameters

- `onVariationValue`: (Optional) The value of the variation of the new on variation.
- `offVariationValue`: (Optional) The value of the variation of the new off variation

Here's an example:

```json
{
"instructions": [ { "kind": "updateDefaultVariation", "OnVariationValue": true, "OffVariationValue": false } ]
}
```

#### updateFallthroughVariationOrRollout

Updates the default or "fallthrough" rule for the AI Config, which the AI Config serves when a context matches none of the targeting rules. The rule can serve either the variation that `variationId` indicates, or a percentage rollout that `rolloutWeights` and `rolloutBucketBy` indicate.

##### Parameters

- `variationId`: ID of a variation.

or

- `rolloutWeights`: Map of `variationId` to weight, in thousandths of a percent (0-100000).
- `rolloutBucketBy`: (Optional) Context attribute available in the specified `rolloutContextKind`.
- `rolloutContextKind`: (Optional) Context kind, defaults to `user`

Here's an example that uses a `variationId`:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "updateFallthroughVariationOrRollout",
  "variationId": "2f43f67c-3e4e-4945-a18a-26559378ca00"
}]
}
```

Here's an example that uses a percentage rollout:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "updateFallthroughVariationOrRollout",
  "rolloutContextKind": "user",
  "rolloutWeights": {
    "2f43f67c-3e4e-4945-a18a-26559378ca00": 15000, // serve 15% this variation
    "e5830889-1ec5-4b0c-9cc9-c48790090c43": 85000  // serve 85% this variation
  }
}]
}
```

#### updateOffVariation

Updates the default off variation to `variationId`. The AI Config serves the default off variation when the AI Config's targeting is **Off**.

##### Parameters

- `variationId`: ID of a variation.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [ { "kind": "updateOffVariation", "variationId": "2f43f67c-3e4e-4945-a18a-26559378ca00" } ]
}
```

#### updateRuleDescription

Updates the description of the targeting rule.

##### Parameters

- `description`: The new human-readable description for this rule.
- `ruleId`: The ID of the rule. You can retrieve this by making a GET request for the AI Config.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "updateRuleDescription",
  "description": "New rule description",
  "ruleId": "a902ef4a-2faf-4eaf-88e1-ecc356708a29"
}]
}
```

#### updateRuleTrackEvents

Updates whether or not LaunchDarkly tracks events for the AI Config associated with this rule.

##### Parameters

- `ruleId`: The ID of the rule. You can retrieve this by making a GET request for the AI Config.
- `trackEvents`: Whether or not events are tracked.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "updateRuleTrackEvents",
  "ruleId": "a902ef4a-2faf-4eaf-88e1-ecc356708a29",
  "trackEvents": true
}]
}
```

#### updateRuleVariationOrRollout

Updates what `ruleId` serves when its clauses evaluate to true. The rule can serve either the variation that `variationId` indicates, or a percent rollout that `rolloutWeights` and `rolloutBucketBy` indicate.

##### Parameters

- `ruleId`: ID of a rule.
- `variationId`: ID of a variation.

or

- `rolloutWeights`: Map of `variationId` to weight, in thousandths of a percent (0-100000).
- `rolloutBucketBy`: (Optional) Context attribute available in the specified `rolloutContextKind`.
- `rolloutContextKind`: (Optional) Context kind, defaults to `user`

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [{
  "kind": "updateRuleVariationOrRollout",
  "ruleId": "a902ef4a-2faf-4eaf-88e1-ecc356708a29",
  "variationId": "2f43f67c-3e4e-4945-a18a-26559378ca00"
}]
}
```

#### updateTrackEvents

Updates whether or not LaunchDarkly tracks events for the AI Config, for all rules.

##### Parameters

- `trackEvents`: Whether or not events are tracked.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [ { "kind": "updateTrackEvents", "trackEvents": true } ]
}
```

#### updateTrackEventsFallthrough

Updates whether or not LaunchDarkly tracks events for the AI Config, for the default rule.

##### Parameters

- `trackEvents`: Whether or not events are tracked.

Here's an example:

```json
{
"environmentKey": "environment-key-123abc",
"instructions": [ { "kind": "updateTrackEventsFallthrough", "trackEvents": true } ]
}
```
</details>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param configKey
 @return ApiPatchAIConfigTargetingRequest
*/
func (a *AIConfigsBetaApiService) PatchAIConfigTargeting(ctx context.Context, projectKey string, configKey string) ApiPatchAIConfigTargetingRequest {
	return ApiPatchAIConfigTargetingRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
	}
}

// Execute executes the request
//  @return AIConfigTargeting
func (a *AIConfigsBetaApiService) PatchAIConfigTargetingExecute(r ApiPatchAIConfigTargetingRequest) (*AIConfigTargeting, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AIConfigTargeting
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.PatchAIConfigTargeting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}/targeting"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	// body params
	localVarPostBody = r.aIConfigTargetingPatch
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPatchAIConfigVariationRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	configKey string
	variationKey string
	aIConfigVariationPatch *AIConfigVariationPatch
}

// Version of the endpoint.
func (r ApiPatchAIConfigVariationRequest) LDAPIVersion(lDAPIVersion string) ApiPatchAIConfigVariationRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// AI Config variation object to update
func (r ApiPatchAIConfigVariationRequest) AIConfigVariationPatch(aIConfigVariationPatch AIConfigVariationPatch) ApiPatchAIConfigVariationRequest {
	r.aIConfigVariationPatch = &aIConfigVariationPatch
	return r
}

func (r ApiPatchAIConfigVariationRequest) Execute() (*AIConfigVariation, *http.Response, error) {
	return r.ApiService.PatchAIConfigVariationExecute(r)
}

/*
PatchAIConfigVariation Update AI Config variation

Edit an existing variation of an AI Config. This creates a new version of the variation.

The request body must be a JSON object of the fields to update. The values you include replace the existing values for the fields.

Here's an example:
```
  {
    "messages": [
      {
        "role": "system",
        "content": "The new message"
      }
    ]
  }
```


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param configKey
 @param variationKey
 @return ApiPatchAIConfigVariationRequest
*/
func (a *AIConfigsBetaApiService) PatchAIConfigVariation(ctx context.Context, projectKey string, configKey string, variationKey string) ApiPatchAIConfigVariationRequest {
	return ApiPatchAIConfigVariationRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
		variationKey: variationKey,
	}
}

// Execute executes the request
//  @return AIConfigVariation
func (a *AIConfigsBetaApiService) PatchAIConfigVariationExecute(r ApiPatchAIConfigVariationRequest) (*AIConfigVariation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AIConfigVariation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.PatchAIConfigVariation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}/variations/{variationKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"variationKey"+"}", url.PathEscape(parameterValueToString(r.variationKey, "variationKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	// body params
	localVarPostBody = r.aIConfigVariationPatch
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPatchAIToolRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	toolKey string
	aIToolPatch *AIToolPatch
}

// Version of the endpoint.
func (r ApiPatchAIToolRequest) LDAPIVersion(lDAPIVersion string) ApiPatchAIToolRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// AI tool object to update
func (r ApiPatchAIToolRequest) AIToolPatch(aIToolPatch AIToolPatch) ApiPatchAIToolRequest {
	r.aIToolPatch = &aIToolPatch
	return r
}

func (r ApiPatchAIToolRequest) Execute() (*AITool, *http.Response, error) {
	return r.ApiService.PatchAIToolExecute(r)
}

/*
PatchAITool Update AI tool

Edit an existing AI tool.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param toolKey
 @return ApiPatchAIToolRequest
*/
func (a *AIConfigsBetaApiService) PatchAITool(ctx context.Context, projectKey string, toolKey string) ApiPatchAIToolRequest {
	return ApiPatchAIToolRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		toolKey: toolKey,
	}
}

// Execute executes the request
//  @return AITool
func (a *AIConfigsBetaApiService) PatchAIToolExecute(r ApiPatchAIToolRequest) (*AITool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AITool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.PatchAITool")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-tools/{toolKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toolKey"+"}", url.PathEscape(parameterValueToString(r.toolKey, "toolKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	// body params
	localVarPostBody = r.aIToolPatch
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPatchAgentGraphRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	graphKey string
	agentGraphPatch *AgentGraphPatch
}

// Version of the endpoint.
func (r ApiPatchAgentGraphRequest) LDAPIVersion(lDAPIVersion string) ApiPatchAgentGraphRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// Agent graph object to update
func (r ApiPatchAgentGraphRequest) AgentGraphPatch(agentGraphPatch AgentGraphPatch) ApiPatchAgentGraphRequest {
	r.agentGraphPatch = &agentGraphPatch
	return r
}

func (r ApiPatchAgentGraphRequest) Execute() (*AgentGraph, *http.Response, error) {
	return r.ApiService.PatchAgentGraphExecute(r)
}

/*
PatchAgentGraph Update agent graph

Edit an existing agent graph.

The request body must be a JSON object of the fields to update. The values you include replace the existing values for the fields.

If the update includes `rootConfigKey` or `edges`, both must be present and will be treated as full replacements.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param graphKey
 @return ApiPatchAgentGraphRequest
*/
func (a *AIConfigsBetaApiService) PatchAgentGraph(ctx context.Context, projectKey string, graphKey string) ApiPatchAgentGraphRequest {
	return ApiPatchAgentGraphRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		graphKey: graphKey,
	}
}

// Execute executes the request
//  @return AgentGraph
func (a *AIConfigsBetaApiService) PatchAgentGraphExecute(r ApiPatchAgentGraphRequest) (*AgentGraph, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentGraph
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.PatchAgentGraph")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/agent-graphs/{graphKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"graphKey"+"}", url.PathEscape(parameterValueToString(r.graphKey, "graphKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	// body params
	localVarPostBody = r.agentGraphPatch
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostAIConfigRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	aIConfigPost *AIConfigPost
}

// Version of the endpoint.
func (r ApiPostAIConfigRequest) LDAPIVersion(lDAPIVersion string) ApiPostAIConfigRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// AI Config object to create
func (r ApiPostAIConfigRequest) AIConfigPost(aIConfigPost AIConfigPost) ApiPostAIConfigRequest {
	r.aIConfigPost = &aIConfigPost
	return r
}

func (r ApiPostAIConfigRequest) Execute() (*AIConfig, *http.Response, error) {
	return r.ApiService.PostAIConfigExecute(r)
}

/*
PostAIConfig Create new AI Config

Create a new AI Config within the given project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiPostAIConfigRequest
*/
func (a *AIConfigsBetaApiService) PostAIConfig(ctx context.Context, projectKey string) ApiPostAIConfigRequest {
	return ApiPostAIConfigRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return AIConfig
func (a *AIConfigsBetaApiService) PostAIConfigExecute(r ApiPostAIConfigRequest) (*AIConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AIConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.PostAIConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}
	if r.aIConfigPost == nil {
		return localVarReturnValue, nil, reportError("aIConfigPost is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	// body params
	localVarPostBody = r.aIConfigPost
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostAIConfigVariationRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	configKey string
	aIConfigVariationPost *AIConfigVariationPost
}

// Version of the endpoint.
func (r ApiPostAIConfigVariationRequest) LDAPIVersion(lDAPIVersion string) ApiPostAIConfigVariationRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// AI Config variation object to create
func (r ApiPostAIConfigVariationRequest) AIConfigVariationPost(aIConfigVariationPost AIConfigVariationPost) ApiPostAIConfigVariationRequest {
	r.aIConfigVariationPost = &aIConfigVariationPost
	return r
}

func (r ApiPostAIConfigVariationRequest) Execute() (*AIConfigVariation, *http.Response, error) {
	return r.ApiService.PostAIConfigVariationExecute(r)
}

/*
PostAIConfigVariation Create AI Config variation

Create a new variation for a given AI Config.

The <code>model</code> in the request body requires a <code>modelName</code> and <code>parameters</code>, for example:

```
  "model": {
    "modelName": "claude-3-opus-20240229",
    "parameters": {
      "max_tokens": 1024
    }
  }
```


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param configKey
 @return ApiPostAIConfigVariationRequest
*/
func (a *AIConfigsBetaApiService) PostAIConfigVariation(ctx context.Context, projectKey string, configKey string) ApiPostAIConfigVariationRequest {
	return ApiPostAIConfigVariationRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
	}
}

// Execute executes the request
//  @return AIConfigVariation
func (a *AIConfigsBetaApiService) PostAIConfigVariationExecute(r ApiPostAIConfigVariationRequest) (*AIConfigVariation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AIConfigVariation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.PostAIConfigVariation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}/variations"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}
	if r.aIConfigVariationPost == nil {
		return localVarReturnValue, nil, reportError("aIConfigVariationPost is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	// body params
	localVarPostBody = r.aIConfigVariationPost
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostAIToolRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	aIToolPost *AIToolPost
}

// Version of the endpoint.
func (r ApiPostAIToolRequest) LDAPIVersion(lDAPIVersion string) ApiPostAIToolRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// AI tool object to create
func (r ApiPostAIToolRequest) AIToolPost(aIToolPost AIToolPost) ApiPostAIToolRequest {
	r.aIToolPost = &aIToolPost
	return r
}

func (r ApiPostAIToolRequest) Execute() (*AITool, *http.Response, error) {
	return r.ApiService.PostAIToolExecute(r)
}

/*
PostAITool Create an AI tool

Create an AI tool

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiPostAIToolRequest
*/
func (a *AIConfigsBetaApiService) PostAITool(ctx context.Context, projectKey string) ApiPostAIToolRequest {
	return ApiPostAIToolRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return AITool
func (a *AIConfigsBetaApiService) PostAIToolExecute(r ApiPostAIToolRequest) (*AITool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AITool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.PostAITool")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-tools"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}
	if r.aIToolPost == nil {
		return localVarReturnValue, nil, reportError("aIToolPost is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	// body params
	localVarPostBody = r.aIToolPost
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostAgentGraphRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	agentGraphPost *AgentGraphPost
}

// Version of the endpoint.
func (r ApiPostAgentGraphRequest) LDAPIVersion(lDAPIVersion string) ApiPostAgentGraphRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// Agent graph object to create
func (r ApiPostAgentGraphRequest) AgentGraphPost(agentGraphPost AgentGraphPost) ApiPostAgentGraphRequest {
	r.agentGraphPost = &agentGraphPost
	return r
}

func (r ApiPostAgentGraphRequest) Execute() (*AgentGraph, *http.Response, error) {
	return r.ApiService.PostAgentGraphExecute(r)
}

/*
PostAgentGraph Create new agent graph

Create a new agent graph within the given project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiPostAgentGraphRequest
*/
func (a *AIConfigsBetaApiService) PostAgentGraph(ctx context.Context, projectKey string) ApiPostAgentGraphRequest {
	return ApiPostAgentGraphRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return AgentGraph
func (a *AIConfigsBetaApiService) PostAgentGraphExecute(r ApiPostAgentGraphRequest) (*AgentGraph, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentGraph
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.PostAgentGraph")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/agent-graphs"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}
	if r.agentGraphPost == nil {
		return localVarReturnValue, nil, reportError("agentGraphPost is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	// body params
	localVarPostBody = r.agentGraphPost
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostModelConfigRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	modelConfigPost *ModelConfigPost
}

// Version of the endpoint.
func (r ApiPostModelConfigRequest) LDAPIVersion(lDAPIVersion string) ApiPostModelConfigRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// AI model config object to create
func (r ApiPostModelConfigRequest) ModelConfigPost(modelConfigPost ModelConfigPost) ApiPostModelConfigRequest {
	r.modelConfigPost = &modelConfigPost
	return r
}

func (r ApiPostModelConfigRequest) Execute() (*ModelConfig, *http.Response, error) {
	return r.ApiService.PostModelConfigExecute(r)
}

/*
PostModelConfig Create an AI model config

Create an AI model config. You can use this in any variation for any AI Config in your project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiPostModelConfigRequest
*/
func (a *AIConfigsBetaApiService) PostModelConfig(ctx context.Context, projectKey string) ApiPostModelConfigRequest {
	return ApiPostModelConfigRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return ModelConfig
func (a *AIConfigsBetaApiService) PostModelConfigExecute(r ApiPostModelConfigRequest) (*ModelConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.PostModelConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/model-configs"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}
	if r.modelConfigPost == nil {
		return localVarReturnValue, nil, reportError("modelConfigPost is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	// body params
	localVarPostBody = r.modelConfigPost
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostRestrictedModelsRequest struct {
	ctx context.Context
	ApiService AIConfigsBetaApi
	lDAPIVersion *string
	projectKey string
	restrictedModelsRequest *RestrictedModelsRequest
}

// Version of the endpoint.
func (r ApiPostRestrictedModelsRequest) LDAPIVersion(lDAPIVersion string) ApiPostRestrictedModelsRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// List of AI model keys to add to the restricted list.
func (r ApiPostRestrictedModelsRequest) RestrictedModelsRequest(restrictedModelsRequest RestrictedModelsRequest) ApiPostRestrictedModelsRequest {
	r.restrictedModelsRequest = &restrictedModelsRequest
	return r
}

func (r ApiPostRestrictedModelsRequest) Execute() (*RestrictedModelsResponse, *http.Response, error) {
	return r.ApiService.PostRestrictedModelsExecute(r)
}

/*
PostRestrictedModels Add AI models to the restricted list

Add AI models, by key, to the restricted list. Keys are included in the response from the [List AI model configs](https://launchdarkly.com/docs/api/ai-configs-beta/list-model-configs) endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiPostRestrictedModelsRequest
*/
func (a *AIConfigsBetaApiService) PostRestrictedModels(ctx context.Context, projectKey string) ApiPostRestrictedModelsRequest {
	return ApiPostRestrictedModelsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return RestrictedModelsResponse
func (a *AIConfigsBetaApiService) PostRestrictedModelsExecute(r ApiPostRestrictedModelsRequest) (*RestrictedModelsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RestrictedModelsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIConfigsBetaApiService.PostRestrictedModels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/model-configs/restricted"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lDAPIVersion == nil {
		return localVarReturnValue, nil, reportError("lDAPIVersion is required and must be specified")
	}
	if r.restrictedModelsRequest == nil {
		return localVarReturnValue, nil, reportError("restrictedModelsRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "LD-API-Version", r.lDAPIVersion, "simple", "")
	// body params
	localVarPostBody = r.restrictedModelsRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
