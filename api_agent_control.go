/*
LaunchDarkly REST API

This documentation describes LaunchDarkly's REST API. To access the complete OpenAPI spec directly, use [Get OpenAPI spec](https://launchdarkly.com/docs/api/other/get-openapi-spec).  To learn how to use LaunchDarkly using the user interface (UI) instead, read our [product documentation](https://launchdarkly.com/docs/home).  ## Authentication  LaunchDarkly's REST API uses the HTTPS protocol with a minimum TLS version of 1.2.  All REST API resources are authenticated with either [personal or service access tokens](https://launchdarkly.com/docs/home/account/api), or session cookies. Other authentication mechanisms are not supported. You can manage personal access tokens on your [**Authorization**](https://app.launchdarkly.com/settings/authorization) page in the LaunchDarkly UI.  LaunchDarkly also has SDK keys, mobile keys, and client-side IDs that are used by our server-side SDKs, mobile SDKs, and JavaScript-based SDKs, respectively. **These keys cannot be used to access our REST API**. These keys are environment-specific, and can only perform read-only operations such as fetching feature flag settings.  | Auth mechanism                                                                                  | Allowed resources                                                                                     | Use cases                                          | | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | -------------------------------------------------- | | [Personal or service access tokens](https://launchdarkly.com/docs/home/account/api) | Can be customized on a per-token basis                                                                | Building scripts, custom integrations, data export. | | SDK keys                                                                                        | Can only access read-only resources specific to server-side SDKs. Restricted to a single environment. | Server-side SDKs                     | | Mobile keys                                                                                     | Can only access read-only resources specific to mobile SDKs, and only for flags marked available to mobile keys. Restricted to a single environment.           | Mobile SDKs                                        | | Client-side ID                                                                                  | Can only access read-only resources specific to JavaScript-based client-side SDKs, and only for flags marked available to client-side. Restricted to a single environment.           | Client-side JavaScript                             |  > #### Keep your access tokens and SDK keys private > > Access tokens should _never_ be exposed in untrusted contexts. Never put an access token in client-side JavaScript, or embed it in a mobile application. LaunchDarkly has special mobile keys that you can embed in mobile apps. If you accidentally expose an access token or SDK key, you can reset it from your [**Authorization**](https://app.launchdarkly.com/settings/authorization) page. > > The client-side ID is safe to embed in untrusted contexts. It's designed for use in client-side JavaScript.  ### Authentication using request header  The preferred way to authenticate with the API is by adding an `Authorization` header containing your access token to your requests. The value of the `Authorization` header must be your access token.  Manage personal access tokens from the [**Authorization**](https://app.launchdarkly.com/settings/authorization) page.  ### Authentication using session cookie  For testing purposes, you can make API calls directly from your web browser. If you are logged in to the LaunchDarkly application, the API will use your existing session to authenticate calls.  Depending on the permissions granted as part of your [role](https://launchdarkly.com/docs/home/account/roles), you may not have permission to perform some API calls. You will receive a `401` response code in that case.  > ### Modifying the Origin header causes an error > > LaunchDarkly validates that the Origin header for any API request authenticated by a session cookie matches the expected Origin header. The expected Origin header is `https://app.launchdarkly.com`. > > If the Origin header does not match what's expected, LaunchDarkly returns an error. This error can prevent the LaunchDarkly app from working correctly. > > Any browser extension that intentionally changes the Origin header can cause this problem. For example, Cross-Origin Resource Sharing (CORS) extensions used during development can modify the Origin header and cause the app to fail. > > To prevent this error, do not modify your Origin header. > > LaunchDarkly does not require origin matching when authenticating with an access token, so this issue does not affect normal API usage.  ## Representations  All resources expect and return JSON response bodies. Error responses also send a JSON body. To learn more about the error format of the API, read [Errors](https://launchdarkly.com/docs/api#errors).  In practice this means that you always get a response with a `Content-Type` header set to `application/json`.  In addition, request bodies for `PATCH`, `POST`, and `PUT` requests must be encoded as JSON with a `Content-Type` header set to `application/json`.  ### Summary and detailed representations  When you fetch a list of resources, the response includes only the most important attributes of each resource. This is a _summary representation_ of the resource. When you fetch an individual resource, such as a single feature flag, you receive a _detailed representation_ of the resource.  The best way to find a detailed representation is to follow links. Every summary representation includes a link to its detailed representation.  ### Expanding responses  Sometimes the detailed representation of a resource does not include all of the attributes of the resource by default. If this is the case, the request method will clearly document this and describe which attributes you can include in an expanded response.  To include the additional attributes, append the `expand` request parameter to your request and add a comma-separated list of the attributes to include. For example, when you append `?expand=members,maintainers` to the [Get team](https://launchdarkly.com/docs/api/teams/get-team) endpoint, the expanded response includes both of these attributes.  ### Links and addressability  The best way to navigate the API is by following links. These are attributes in representations that link to other resources. The API always uses the same format for links:  - Links to other resources within the API are encapsulated in a `_links` object - If the resource has a corresponding link to HTML content on the site, it is stored in a special `_site` link  Each link has two attributes:  - An `href`, which contains the URL - A `type`, which describes the content type  For example, a feature resource might return the following:  ```json {   \"_links\": {     \"parent\": {       \"href\": \"/api/features\",       \"type\": \"application/json\"     },     \"self\": {       \"href\": \"/api/features/sort.order\",       \"type\": \"application/json\"     }   },   \"_site\": {     \"href\": \"/features/sort.order\",     \"type\": \"text/html\"   } } ```  From this, you can navigate to the parent collection of features by following the `parent` link, or navigate to the site page for the feature by following the `_site` link.  Collections are always represented as a JSON object with an `items` attribute containing an array of representations. Like all other representations, collections have `_links` defined at the top level.  Paginated collections include `first`, `last`, `next`, and `prev` links containing a URL with the respective set of elements in the collection.  ## Updates  Resources that accept partial updates use the `PATCH` verb. Most resources support the [JSON patch](https://launchdarkly.com/docs/api#updates-using-json-patch) format. Some resources also support the [JSON merge patch](https://launchdarkly.com/docs/api#updates-using-json-merge-patch) format, and some resources support the [semantic patch](https://launchdarkly.com/docs/api#updates-using-semantic-patch) format, which is a way to specify the modifications to perform as a set of executable instructions. Each resource supports optional [comments](https://launchdarkly.com/docs/api#updates-with-comments) that you can submit with updates. Comments appear in outgoing webhooks, the audit log, and other integrations.  When a resource supports both JSON patch and semantic patch, we document both in the request method. However, the specific request body fields and descriptions included in our documentation only match one type of patch or the other.  ### Updates using JSON patch  [JSON patch](https://datatracker.ietf.org/doc/html/rfc6902) is a way to specify the modifications to perform on a resource. JSON patch uses paths and a limited set of operations to describe how to transform the current state of the resource into a new state. JSON patch documents are always arrays, where each element contains an operation, a path to the field to update, and the new value.  For example, in this feature flag representation:  ```json {     \"name\": \"New recommendations engine\",     \"key\": \"engine.enable\",     \"description\": \"This is the description\",     ... } ``` You can change the feature flag's description with the following patch document:  ```json [{ \"op\": \"replace\", \"path\": \"/description\", \"value\": \"This is the new description\" }] ```  You can specify multiple modifications to perform in a single request. You can also test that certain preconditions are met before applying the patch:  ```json [   { \"op\": \"test\", \"path\": \"/version\", \"value\": 10 },   { \"op\": \"replace\", \"path\": \"/description\", \"value\": \"The new description\" } ] ```  The above patch request tests whether the feature flag's `version` is `10`, and if so, changes the feature flag's description.  Attributes that are not editable, such as a resource's `_links`, have names that start with an underscore.  ### Updates using JSON merge patch  [JSON merge patch](https://datatracker.ietf.org/doc/html/rfc7386) is another format for specifying the modifications to perform on a resource. JSON merge patch is less expressive than JSON patch. However, in many cases it is simpler to construct a merge patch document. For example, you can change a feature flag's description with the following merge patch document:  ```json {   \"description\": \"New flag description\" } ```  ### Updates using semantic patch  Some resources support the semantic patch format. A semantic patch is a way to specify the modifications to perform on a resource as a set of executable instructions.  Semantic patch allows you to be explicit about intent using precise, custom instructions. In many cases, you can define semantic patch instructions independently of the current state of the resource. This can be useful when defining a change that may be applied at a future date.  To make a semantic patch request, you must append `domain-model=launchdarkly.semanticpatch` to your `Content-Type` header.  Here's how:  ``` Content-Type: application/json; domain-model=launchdarkly.semanticpatch ```  If you call a semantic patch resource without this header, you will receive a `400` response because your semantic patch will be interpreted as a JSON patch.  The body of a semantic patch request takes the following properties:  * `comment` (string): (Optional) A description of the update. * `environmentKey` (string): (Required for some resources only) The environment key. * `instructions` (array): (Required) A list of actions the update should perform. Each action in the list must be an object with a `kind` property that indicates the instruction. If the instruction requires parameters, you must include those parameters as additional fields in the object. The documentation for each resource that supports semantic patch includes the available instructions and any additional parameters.  For example:  ```json {   \"comment\": \"optional comment\",   \"instructions\": [ {\"kind\": \"turnFlagOn\"} ] } ```  Semantic patches are not applied partially; either all of the instructions are applied or none of them are. If **any** instruction is invalid, the endpoint returns an error and will not change the resource. If all instructions are valid, the request succeeds and the resources are updated if necessary, or left unchanged if they are already in the state you request.  ### Updates with comments  You can submit optional comments with `PATCH` changes.  To submit a comment along with a JSON patch document, use the following format:  ```json {   \"comment\": \"This is a comment string\",   \"patch\": [{ \"op\": \"replace\", \"path\": \"/description\", \"value\": \"The new description\" }] } ```  To submit a comment along with a JSON merge patch document, use the following format:  ```json {   \"comment\": \"This is a comment string\",   \"merge\": { \"description\": \"New flag description\" } } ```  To submit a comment along with a semantic patch, use the following format:  ```json {   \"comment\": \"This is a comment string\",   \"instructions\": [ {\"kind\": \"turnFlagOn\"} ] } ```  ## Errors  The API always returns errors in a common format. Here's an example:  ```json {   \"code\": \"invalid_request\",   \"message\": \"A feature with that key already exists\",   \"id\": \"30ce6058-87da-11e4-b116-123b93f75cba\" } ```  The `code` indicates the general class of error. The `message` is a human-readable explanation of what went wrong. The `id` is a unique identifier. Use it when you're working with LaunchDarkly Support to debug a problem with a specific API call.  ### HTTP status error response codes  | Code | Definition        | Description                                                                                       | Possible Solution                                                | | ---- | ----------------- | ------------------------------------------------------------------------------------------- | ---------------------------------------------------------------- | | 400  | Invalid request       | The request cannot be understood.                                    | Ensure JSON syntax in request body is correct.                   | | 401  | Invalid access token      | Requestor is unauthorized or does not have permission for this API call.                                                | Ensure your API access token is valid and has the appropriate permissions.                                     | | 403  | Forbidden         | Requestor does not have access to this resource.                                                | Ensure that the account member or access token has proper permissions set. | | 404  | Invalid resource identifier | The requested resource is not valid. | Ensure that the resource is correctly identified by ID or key. | | 405  | Method not allowed | The request method is not allowed on this resource. | Ensure that the HTTP verb is correct. | | 409  | Conflict          | The API request can not be completed because it conflicts with a concurrent API request. | Retry your request.                                              | | 422  | Unprocessable entity | The API request can not be completed because the update description can not be understood. | Ensure that the request body is correct for the type of patch you are using, either JSON patch or semantic patch. | 429  | Too many requests | Read [Rate limiting](https://launchdarkly.com/docs/api#rate-limiting).                                               | Wait and try again later.                                        |  ## CORS  The LaunchDarkly API supports Cross Origin Resource Sharing (CORS) for AJAX requests from any origin. If an `Origin` header is given in a request, it will be echoed as an explicitly allowed origin. Otherwise the request returns a wildcard, `Access-Control-Allow-Origin: *`. For more information on CORS, read the [CORS W3C Recommendation](http://www.w3.org/TR/cors). Example CORS headers might look like:  ```http Access-Control-Allow-Headers: Accept, Content-Type, Content-Length, Accept-Encoding, Authorization Access-Control-Allow-Methods: OPTIONS, GET, DELETE, PATCH Access-Control-Allow-Origin: * Access-Control-Max-Age: 300 ```  You can make authenticated CORS calls just as you would make same-origin calls, using either [token or session-based authentication](https://launchdarkly.com/docs/api#authentication). If you are using session authentication, you should set the `withCredentials` property for your `xhr` request to `true`. You should never expose your access tokens to untrusted entities.  ## Rate limiting  We use several rate-limiting strategies to ensure the availability of our APIs. Rate-limited calls to our APIs return a `429` status code and include headers to indicate the current rate limit status. The specific headers returned depend on the API route that was called. Limits differ based on the route, authentication mechanism, and other factors.  Each set of headers below appears only when the corresponding limit is being enforced for your call. A given route may be subject to any combination of these limits, so a response can include one, several, or none of these headers. A missing header indicates that the limit was not applied to this specific call; it does not necessarily indicate that the limit does not exist. To reduce usage before hitting a `429` status, program against whichever rate limit headers are present rather than expecting a specific header.  We do not publicly document the specific number of calls permitted by any of these limits, and these limits may change. We encourage clients to program against the specification and rely on the headers described below, rather than hardcoding the current limits.  > ### Rate limiting and SDKs > > LaunchDarkly SDKs are never rate limited and do not use the API endpoints defined here. LaunchDarkly uses a different set of approaches, including streaming/server-sent events and a global CDN, to ensure availability to the routes used by LaunchDarkly SDKs.  ### Global rate limits  Authenticated requests are subject to a global limit. This is the maximum number of calls that your account can make to the API per ten seconds. All service and personal access tokens on the account share this limit, so exceeding the limit with one access token will impact other tokens. Calls that are subject to global rate limits may return the headers below:  | Header name                    | Description                                                                      | | ------------------------------ | -------------------------------------------------------------------------------- | | `X-Ratelimit-Global-Limit`     | The maximum number of requests the account is permitted to make per ten seconds. | | `X-Ratelimit-Global-Remaining` | The number of requests remaining in the current global rate limit window.        | | `X-Ratelimit-Reset`            | The time at which the current rate limit window resets in epoch milliseconds.    |  ### Route-level rate limits  Some authenticated routes have custom rate limits. These also reset every ten seconds. Any service or personal access tokens hitting the same route share this limit, so exceeding the limit with one access token may impact other tokens. Calls that are subject to route-level rate limits return the headers below:  | Header name                   | Description                                                                                           | | ----------------------------- | ----------------------------------------------------------------------------------------------------- | | `X-Ratelimit-Route-Limit`     | The maximum number of requests to the current route permitted per ten seconds.           | | `X-Ratelimit-Route-Remaining` | The number of requests remaining for the current route in the current rate limit window. | | `X-Ratelimit-Reset`           | The time at which the current rate limit window resets in epoch milliseconds.            |  A _route_ represents a specific URL pattern and verb. For example, the [Delete environment](https://launchdarkly.com/docs/api/environments/delete-environment) endpoint is considered a single route, and each call to delete an environment counts against your route-level rate limit for that route.  ### Access token rate limits  Some calls are rate limited per access token. Unlike the global and route-level limits, this limit applies to a single service or personal access token on its own. Exceeding a limit with one access token does not affect other tokens on the account. Calls that are subject to access token rate limits return these headers:  | Header name                        | Description                                                                             | | ---------------------------------- | --------------------------------------------------------------------------------------- | | `X-Ratelimit-Auth-Token-Limit`     | The maximum number of requests the access token can make per ten seconds.               | | `X-Ratelimit-Auth-Token-Remaining` | The number of requests remaining for the access token in the current rate limit window. | | `X-Ratelimit-Auth-Token-Reset`     | The time at which the current rate limit window resets in epoch milliseconds.           |  Unlike the other rate limits, access token rate limits report their own reset time in the `X-Ratelimit-Auth-Token-Reset` header instead of in `X-Ratelimit-Reset`.  ### IP-based rate limiting  We also employ IP-based rate limiting on some API routes. If you hit an IP-based rate limit, your API response will include a `Retry-After` header indicating how long to wait before re-trying the call. Clients must wait at least `Retry-After` seconds before making additional calls to our API, and should employ jitter and backoff strategies to avoid triggering rate limits again.  ## OpenAPI (Swagger) and client libraries  We have a [complete OpenAPI (Swagger) specification](https://app.launchdarkly.com/api/v2/openapi.json) for our API.  We auto-generate multiple client libraries based on our OpenAPI specification. To learn more, visit the [collection of client libraries on GitHub](https://github.com/search?q=topic%3Alaunchdarkly-api+org%3Alaunchdarkly&type=Repositories). Alternatively, you can use the specification to generate client libraries to interact with our REST API in your language of choice. Or, you can refer to our API endpoints' documentation for guidance on how to make requests with a common HTTP library in your language of choice.  Our OpenAPI specification is supported by several API-based tools such as Postman and Insomnia. In many cases, you can directly import our specification to explore our APIs.  ## Method overriding  Some firewalls and HTTP clients restrict the use of verbs other than `GET` and `POST`. In those environments, our API endpoints that use `DELETE`, `PATCH`, and `PUT` verbs are inaccessible.  To avoid this issue, our API supports the `X-HTTP-Method-Override` header, allowing clients to \"tunnel\" `DELETE`, `PATCH`, and `PUT` requests using a `POST` request.  For example, to call a `PATCH` endpoint using a `POST` request, you can include `X-HTTP-Method-Override:PATCH` as a header.  ## Beta resources  We sometimes release new API resources in **beta** status before we release them with general availability.  Resources that are in beta are still undergoing testing and development. They may change without notice, including becoming backwards incompatible.  We try to promote resources into general availability as quickly as possible. This happens after sufficient testing and when we're satisfied that we no longer need to make backwards-incompatible changes.  We mark beta resources with a \"Beta\" callout in our documentation, pictured below:  > ### This feature is in beta > > To use this feature, pass in a header including the `LD-API-Version` key with value set to `beta`. Use this header with each call. To learn more, read [Beta resources](https://launchdarkly.com/docs/api#beta-resources). > > Resources that are in beta are still undergoing testing and development. They may change without notice, including becoming backwards incompatible.  ### Using beta resources  To use a beta resource, you must include a header in the request. If you call a beta resource without this header, you receive a `403` response.  Use this header:  ``` LD-API-Version: beta ```  ## Federal and EU environments  In addition to the commercial versions, LaunchDarkly offers instances for federal agencies and those based in the European Union (EU).  ### Federal environments  The version of LaunchDarkly that is available on domains controlled by the United States government is different from the version of LaunchDarkly available to the general public. If you are an employee or contractor for a United States federal agency and use LaunchDarkly in your work, you likely use the federal instance of LaunchDarkly.  If you are working in the federal instance of LaunchDarkly, the base URI for each request is `https://app.launchdarkly.us`.  To learn more, read [LaunchDarkly in federal environments](https://launchdarkly.com/docs/home/infrastructure/federal).  ### EU environments  The version of LaunchDarkly that is available in the EU is different from the version of LaunchDarkly available to other regions. If you are based in the EU, you likely use the EU instance of LaunchDarkly. The LaunchDarkly EU instance complies with EU data residency principles, including the protection and confidentiality of EU customer information.  If you are working in the EU instance of LaunchDarkly, the base URI for each request is `https://app.eu.launchdarkly.com`.  To learn more, read [LaunchDarkly in the European Union (EU)](https://launchdarkly.com/docs/home/infrastructure/eu).  ## Versioning  We try hard to keep our REST API backwards compatible, but we occasionally have to make backwards-incompatible changes in the process of shipping new features. These breaking changes can cause unexpected behavior if you don't prepare for them accordingly.  Updates to our REST API include support for the latest features in LaunchDarkly. We also release a new version of our REST API every time we make a breaking change. We provide simultaneous support for multiple API versions so you can migrate from your current API version to a new version at your own pace.  ### Setting the API version per request  You can set the API version on a specific request by sending an `LD-API-Version` header, as shown in the example below:  ``` LD-API-Version: 20240415 ```  The header value is the version number of the API version you would like to request. The number for each version corresponds to the date the version was released in `yyyymmdd` format. In the example above the version `20240415` corresponds to April 15, 2024.  ### Setting the API version per access token  When you create an access token, you must specify a specific version of the API to use. This ensures that integrations using this token cannot be broken by version changes.  Tokens created before versioning was released have their version set to `20160426`, which is the version of the API that existed before the current versioning scheme, so that they continue working the same way they did before versioning.  If you would like to upgrade your integration to use a new API version, you can explicitly set the header described above.  > ### Best practice: Set the header for every client or integration > > We recommend that you set the API version header explicitly in any client or integration you build. > > Only rely on the access token API version during manual testing.  ### API version changelog  <table>   <tr>     <th>Version</th>     <th>Changes</th>     <th>End of life (EOL)</th>   </tr>   <tr>     <td>`20240415`</td>     <td>       <ul><li>Changed several endpoints from unpaginated to paginated. Use the `limit` and `offset` query parameters to page through the results.</li> <li>Changed the [list access tokens](https://launchdarkly.com/docs/api/access-tokens/get-tokens) endpoint: <ul><li>Response is now paginated with a default limit of `25`</li></ul></li> <li>Changed the [list account members](https://launchdarkly.com/docs/api/account-members/get-members) endpoint: <ul><li>The `accessCheck` filter is no longer available</li></ul></li> <li>Changed the [list custom roles](https://launchdarkly.com/docs/api/custom-roles/get-custom-roles) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li></ul></li> <li>Changed the [list feature flags](https://launchdarkly.com/docs/api/feature-flags/get-feature-flags) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li><li>The `environments` field is now only returned if the request is filtered by environment, using the `filterEnv` query parameter</li><li>The `followerId`, `hasDataExport`, `status`, `contextKindTargeted`, and `segmentTargeted` filters are no longer available</li><li>The `compare` query parameter is no longer available</li></ul></li> <li>Changed the [list segments](https://launchdarkly.com/docs/api/segments/get-segments) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li></ul></li> <li>Changed the [list teams](https://launchdarkly.com/docs/api/teams/get-teams) endpoint: <ul><li>The `expand` parameter no longer supports including `projects` or `roles`</li><li>In paginated results, the maximum page size is now 100</li></ul></li> <li>Changed the [get workflows](https://launchdarkly.com/docs/api/workflows/get-workflows) endpoint: <ul><li>Response is now paginated with a default limit of `20`</li><li>The `_conflicts` field in the response is no longer available</li></ul></li> </ul>     </td>     <td>Current</td>   </tr>   <tr>     <td>`20220603`</td>     <td>       <ul><li>Changed the [list projects](https://launchdarkly.com/docs/api/projects/get-projects) return value:<ul><li>Response is now paginated with a default limit of `20`.</li><li>Added support for filter and sort.</li><li>The project `environments` field is now expandable. This field is omitted by default.</li></ul></li><li>Changed the [get project](https://launchdarkly.com/docs/api/projects/get-project) return value:<ul><li>The `environments` field is now expandable. This field is omitted by default.</li></ul></li></ul>     </td>     <td>2025-04-15</td>   </tr>   <tr>     <td>`20210729`</td>     <td>       <ul><li>Changed the [create approval request](https://launchdarkly.com/docs/api/approvals/post-approval-request) return value. It now returns HTTP Status Code `201` instead of `200`.</li><li> Changed the [get user](https://launchdarkly.com/docs/api/users/get-user) return value. It now returns a user record, not a user. </li><li>Added additional optional fields to environment, segments, flags, members, and segments, including the ability to create big segments. </li><li> Added default values for flag variations when new environments are created. </li><li>Added filtering and pagination for getting flags and members, including `limit`, `number`, `filter`, and `sort` query parameters. </li><li>Added endpoints for expiring user targets for flags and segments, scheduled changes, access tokens, Relay Proxy configuration, integrations and subscriptions, and approvals. </li></ul>     </td>     <td>2023-06-03</td>   </tr>   <tr>     <td>`20191212`</td>     <td>       <ul><li>[List feature flags](https://launchdarkly.com/docs/api/feature-flags/get-feature-flags) now defaults to sending summaries of feature flag configurations, equivalent to setting the query parameter `summary=true`. Summaries omit flag targeting rules and individual user targets from the payload. </li><li> Added endpoints for flags, flag status, projects, environments, audit logs, members, users, custom roles, segments, usage, streams, events, and data export. </li></ul>     </td>     <td>2022-07-29</td>   </tr>   <tr>     <td>`20160426`</td>     <td>       <ul><li>Initial versioning of API. Tokens created before versioning have their version set to this.</li></ul>     </td>     <td>2020-12-12</td>   </tr> </table>  To learn more about how EOL is determined, read LaunchDarkly's [End of Life (EOL) Policy](https://launchdarkly.com/policies/end-of-life-policy/). 

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


type AgentControlApi interface {

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
	DeleteAgentOptimization Delete an agent optimization

	Delete an existing agent optimization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param optimizationKey
	@return ApiDeleteAgentOptimizationRequest
	*/
	DeleteAgentOptimization(ctx context.Context, projectKey string, optimizationKey string) ApiDeleteAgentOptimizationRequest

	// DeleteAgentOptimizationExecute executes the request
	DeleteAgentOptimizationExecute(r ApiDeleteAgentOptimizationRequest) (*http.Response, error)

	/*
	DeleteAgentOptimizationRun Delete an agent optimization run

	Delete all results for a specific run of an agent optimization. Returns a 404 if the optimization does not exist or if the run has no results. A run whose results have already been deleted is indistinguishable from a run that never existed, so repeating this request also returns a 404.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param optimizationKey
	@param runId
	@return ApiDeleteAgentOptimizationRunRequest
	*/
	DeleteAgentOptimizationRun(ctx context.Context, projectKey string, optimizationKey string, runId string) ApiDeleteAgentOptimizationRunRequest

	// DeleteAgentOptimizationRunExecute executes the request
	DeleteAgentOptimizationRunExecute(r ApiDeleteAgentOptimizationRunRequest) (*http.Response, error)

	/*
	DeleteAgentSkill Delete an agent skill

	Delete an agent skill, including all of its versions.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param skillKey
	@return ApiDeleteAgentSkillRequest
	*/
	DeleteAgentSkill(ctx context.Context, projectKey string, skillKey string) ApiDeleteAgentSkillRequest

	// DeleteAgentSkillExecute executes the request
	DeleteAgentSkillExecute(r ApiDeleteAgentSkillRequest) (*http.Response, error)

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
	DeletePromptSnippet Delete a prompt snippet

	Delete an existing prompt snippet.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param snippetKey
	@return ApiDeletePromptSnippetRequest
	*/
	DeletePromptSnippet(ctx context.Context, projectKey string, snippetKey string) ApiDeletePromptSnippetRequest

	// DeletePromptSnippetExecute executes the request
	DeletePromptSnippetExecute(r ApiDeletePromptSnippetRequest) (*http.Response, error)

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
	GetAIConfigQuickStats Get AI Config quick stats

	Retrieve aggregate quick stats for AI Configs in a project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiGetAIConfigQuickStatsRequest
	*/
	GetAIConfigQuickStats(ctx context.Context, projectKey string) ApiGetAIConfigQuickStatsRequest

	// GetAIConfigQuickStatsExecute executes the request
	//  @return QuickStats
	GetAIConfigQuickStatsExecute(r ApiGetAIConfigQuickStatsRequest) (*QuickStats, *http.Response, error)

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
	GetAgentOptimization Get an agent optimization

	Retrieve a specific agent optimization by its key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param optimizationKey
	@return ApiGetAgentOptimizationRequest
	*/
	GetAgentOptimization(ctx context.Context, projectKey string, optimizationKey string) ApiGetAgentOptimizationRequest

	// GetAgentOptimizationExecute executes the request
	//  @return AgentOptimization
	GetAgentOptimizationExecute(r ApiGetAgentOptimizationRequest) (*AgentOptimization, *http.Response, error)

	/*
	GetAgentSkill Get an agent skill

	Retrieve a specific agent skill by its key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param skillKey
	@return ApiGetAgentSkillRequest
	*/
	GetAgentSkill(ctx context.Context, projectKey string, skillKey string) ApiGetAgentSkillRequest

	// GetAgentSkillExecute executes the request
	//  @return AgentSkill
	GetAgentSkillExecute(r ApiGetAgentSkillRequest) (*AgentSkill, *http.Response, error)

	/*
	GetModelConfig Get AI model config

	Get an AI model config by key. Specify a version to retrieve a historical custom model config version. Global model configs accept version 0 or 1 and return their current representation.

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
	GetPromptSnippet Get a prompt snippet

	Retrieve a specific prompt snippet by its key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param snippetKey
	@return ApiGetPromptSnippetRequest
	*/
	GetPromptSnippet(ctx context.Context, projectKey string, snippetKey string) ApiGetPromptSnippetRequest

	// GetPromptSnippetExecute executes the request
	//  @return PromptSnippet
	GetPromptSnippetExecute(r ApiGetPromptSnippetRequest) (*PromptSnippet, *http.Response, error)

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
	ListAgentOptimizationResults List agent optimization runs

	Get the most recent result for each unique run of an agent optimization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param optimizationKey
	@return ApiListAgentOptimizationResultsRequest
	*/
	ListAgentOptimizationResults(ctx context.Context, projectKey string, optimizationKey string) ApiListAgentOptimizationResultsRequest

	// ListAgentOptimizationResultsExecute executes the request
	//  @return AgentOptimizationResults
	ListAgentOptimizationResultsExecute(r ApiListAgentOptimizationResultsRequest) (*AgentOptimizationResults, *http.Response, error)

	/*
	ListAgentOptimizationResultsByRunId List agent optimization results for a run

	Get all results for a specific run of an agent optimization, sorted by iteration descending.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param optimizationKey
	@param runId
	@return ApiListAgentOptimizationResultsByRunIdRequest
	*/
	ListAgentOptimizationResultsByRunId(ctx context.Context, projectKey string, optimizationKey string, runId string) ApiListAgentOptimizationResultsByRunIdRequest

	// ListAgentOptimizationResultsByRunIdExecute executes the request
	//  @return AgentOptimizationResults
	ListAgentOptimizationResultsByRunIdExecute(r ApiListAgentOptimizationResultsByRunIdRequest) (*AgentOptimizationResults, *http.Response, error)

	/*
	ListAgentOptimizations List agent optimizations

	Get a list of all agent optimizations in the given project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiListAgentOptimizationsRequest
	*/
	ListAgentOptimizations(ctx context.Context, projectKey string) ApiListAgentOptimizationsRequest

	// ListAgentOptimizationsExecute executes the request
	//  @return AgentOptimizations
	ListAgentOptimizationsExecute(r ApiListAgentOptimizationsRequest) (*AgentOptimizations, *http.Response, error)

	/*
	ListAgentSkillReferences List agent skill references

	Get all config variations that currently reference this agent skill.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param skillKey
	@return ApiListAgentSkillReferencesRequest
	*/
	ListAgentSkillReferences(ctx context.Context, projectKey string, skillKey string) ApiListAgentSkillReferencesRequest

	// ListAgentSkillReferencesExecute executes the request
	//  @return AgentSkillReferences
	ListAgentSkillReferencesExecute(r ApiListAgentSkillReferencesRequest) (*AgentSkillReferences, *http.Response, error)

	/*
	ListAgentSkillVersions List agent skill versions

	Get all versions of an agent skill in the given project, ordered by version descending.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param skillKey
	@return ApiListAgentSkillVersionsRequest
	*/
	ListAgentSkillVersions(ctx context.Context, projectKey string, skillKey string) ApiListAgentSkillVersionsRequest

	// ListAgentSkillVersionsExecute executes the request
	//  @return AgentSkills
	ListAgentSkillVersionsExecute(r ApiListAgentSkillVersionsRequest) (*AgentSkills, *http.Response, error)

	/*
	ListAgentSkills List agent skills

	Get a list of all agent skills in the given project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiListAgentSkillsRequest
	*/
	ListAgentSkills(ctx context.Context, projectKey string) ApiListAgentSkillsRequest

	// ListAgentSkillsExecute executes the request
	//  @return AgentSkills
	ListAgentSkillsExecute(r ApiListAgentSkillsRequest) (*AgentSkills, *http.Response, error)

	/*
	ListAllAgentOptimizationResults List all agent optimization results across versions

	Get all results across all versions of an agent optimization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param optimizationKey
	@return ApiListAllAgentOptimizationResultsRequest
	*/
	ListAllAgentOptimizationResults(ctx context.Context, projectKey string, optimizationKey string) ApiListAllAgentOptimizationResultsRequest

	// ListAllAgentOptimizationResultsExecute executes the request
	//  @return AgentOptimizationResults
	ListAllAgentOptimizationResultsExecute(r ApiListAllAgentOptimizationResultsRequest) (*AgentOptimizationResults, *http.Response, error)

	/*
	ListModelConfigVersions List AI model config versions

	Get a paginated list of all available versions of an AI model config, ordered from newest to oldest. Global model configs return their single current representation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param modelConfigKey
	@return ApiListModelConfigVersionsRequest
	*/
	ListModelConfigVersions(ctx context.Context, projectKey string, modelConfigKey string) ApiListModelConfigVersionsRequest

	// ListModelConfigVersionsExecute executes the request
	//  @return ModelConfigs
	ListModelConfigVersionsExecute(r ApiListModelConfigVersionsRequest) (*ModelConfigs, *http.Response, error)

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
	ListPromptSnippetReferences List prompt snippet references

	Get all config variations that currently reference this prompt snippet.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param snippetKey
	@return ApiListPromptSnippetReferencesRequest
	*/
	ListPromptSnippetReferences(ctx context.Context, projectKey string, snippetKey string) ApiListPromptSnippetReferencesRequest

	// ListPromptSnippetReferencesExecute executes the request
	//  @return SnippetReferences
	ListPromptSnippetReferencesExecute(r ApiListPromptSnippetReferencesRequest) (*SnippetReferences, *http.Response, error)

	/*
	ListPromptSnippetVersions List prompt snippet versions

	Get all versions of a prompt snippet in the given project, ordered by version descending.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param snippetKey
	@return ApiListPromptSnippetVersionsRequest
	*/
	ListPromptSnippetVersions(ctx context.Context, projectKey string, snippetKey string) ApiListPromptSnippetVersionsRequest

	// ListPromptSnippetVersionsExecute executes the request
	//  @return PromptSnippets
	ListPromptSnippetVersionsExecute(r ApiListPromptSnippetVersionsRequest) (*PromptSnippets, *http.Response, error)

	/*
	ListPromptSnippets List prompt snippets

	Get a list of all prompt snippets in the given project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiListPromptSnippetsRequest
	*/
	ListPromptSnippets(ctx context.Context, projectKey string) ApiListPromptSnippetsRequest

	// ListPromptSnippetsExecute executes the request
	//  @return PromptSnippets
	ListPromptSnippetsExecute(r ApiListPromptSnippetsRequest) (*PromptSnippets, *http.Response, error)

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
	PatchAgentOptimization Update an agent optimization

	Update an existing agent optimization. Creates a new version of the optimization.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param optimizationKey
	@return ApiPatchAgentOptimizationRequest
	*/
	PatchAgentOptimization(ctx context.Context, projectKey string, optimizationKey string) ApiPatchAgentOptimizationRequest

	// PatchAgentOptimizationExecute executes the request
	//  @return AgentOptimization
	PatchAgentOptimizationExecute(r ApiPatchAgentOptimizationRequest) (*AgentOptimization, *http.Response, error)

	/*
	PatchAgentOptimizationResult Update an agent optimization result

	Update status, activity, and AI request metadata fields on an agent optimization result.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param optimizationKey
	@param resultId
	@return ApiPatchAgentOptimizationResultRequest
	*/
	PatchAgentOptimizationResult(ctx context.Context, projectKey string, optimizationKey string, resultId string) ApiPatchAgentOptimizationResultRequest

	// PatchAgentOptimizationResultExecute executes the request
	//  @return AgentOptimizationResult
	PatchAgentOptimizationResultExecute(r ApiPatchAgentOptimizationResultRequest) (*AgentOptimizationResult, *http.Response, error)

	/*
	PatchAgentSkill Update an agent skill

	Update an existing agent skill. Creates a new version of the skill.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param skillKey
	@return ApiPatchAgentSkillRequest
	*/
	PatchAgentSkill(ctx context.Context, projectKey string, skillKey string) ApiPatchAgentSkillRequest

	// PatchAgentSkillExecute executes the request
	//  @return AgentSkill
	PatchAgentSkillExecute(r ApiPatchAgentSkillRequest) (*AgentSkill, *http.Response, error)

	/*
	PatchModelConfig Update an AI model config

	Update an AI model config.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param modelConfigKey
	@return ApiPatchModelConfigRequest
	*/
	PatchModelConfig(ctx context.Context, projectKey string, modelConfigKey string) ApiPatchModelConfigRequest

	// PatchModelConfigExecute executes the request
	//  @return ModelConfig
	PatchModelConfigExecute(r ApiPatchModelConfigRequest) (*ModelConfig, *http.Response, error)

	/*
	PatchPromptSnippet Update a prompt snippet

	Update an existing prompt snippet. Creates a new version of the snippet.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param snippetKey
	@return ApiPatchPromptSnippetRequest
	*/
	PatchPromptSnippet(ctx context.Context, projectKey string, snippetKey string) ApiPatchPromptSnippetRequest

	// PatchPromptSnippetExecute executes the request
	//  @return PromptSnippet
	PatchPromptSnippetExecute(r ApiPatchPromptSnippetRequest) (*PromptSnippet, *http.Response, error)

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
	PostAgentOptimization Create agent optimization

	Create a new agent optimization within the given project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiPostAgentOptimizationRequest
	*/
	PostAgentOptimization(ctx context.Context, projectKey string) ApiPostAgentOptimizationRequest

	// PostAgentOptimizationExecute executes the request
	//  @return AgentOptimization
	PostAgentOptimizationExecute(r ApiPostAgentOptimizationRequest) (*AgentOptimization, *http.Response, error)

	/*
	PostAgentOptimizationResult Create agent optimization result

	Create a new result for an agent optimization run.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@param optimizationKey
	@return ApiPostAgentOptimizationResultRequest
	*/
	PostAgentOptimizationResult(ctx context.Context, projectKey string, optimizationKey string) ApiPostAgentOptimizationResultRequest

	// PostAgentOptimizationResultExecute executes the request
	//  @return AgentOptimizationResult
	PostAgentOptimizationResultExecute(r ApiPostAgentOptimizationResultRequest) (*AgentOptimizationResult, *http.Response, error)

	/*
	PostAgentSkill Create an agent skill

	Create a new agent skill within the given project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiPostAgentSkillRequest
	*/
	PostAgentSkill(ctx context.Context, projectKey string) ApiPostAgentSkillRequest

	// PostAgentSkillExecute executes the request
	//  @return AgentSkill
	PostAgentSkillExecute(r ApiPostAgentSkillRequest) (*AgentSkill, *http.Response, error)

	/*
	PostModelConfig Create an AI model config

	Create an AI model config. You can use this in any variation for any config in your project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiPostModelConfigRequest
	*/
	PostModelConfig(ctx context.Context, projectKey string) ApiPostModelConfigRequest

	// PostModelConfigExecute executes the request
	//  @return ModelConfig
	PostModelConfigExecute(r ApiPostModelConfigRequest) (*ModelConfig, *http.Response, error)

	/*
	PostPromptSnippet Create a prompt snippet

	Create a new prompt snippet within the given project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiPostPromptSnippetRequest
	*/
	PostPromptSnippet(ctx context.Context, projectKey string) ApiPostPromptSnippetRequest

	// PostPromptSnippetExecute executes the request
	//  @return PromptSnippet
	PostPromptSnippetExecute(r ApiPostPromptSnippetRequest) (*PromptSnippet, *http.Response, error)

	/*
	PostRestrictedModels Add AI models to the restricted list

	Add AI models, by key, to the restricted list. Keys are included in the response from the [List AI model configs](https://launchdarkly.com/docs/api/agent-control/list-model-configs) endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectKey
	@return ApiPostRestrictedModelsRequest
	*/
	PostRestrictedModels(ctx context.Context, projectKey string) ApiPostRestrictedModelsRequest

	// PostRestrictedModelsExecute executes the request
	//  @return RestrictedModelsResponse
	PostRestrictedModelsExecute(r ApiPostRestrictedModelsRequest) (*RestrictedModelsResponse, *http.Response, error)
}

// AgentControlApiService AgentControlApi service
type AgentControlApiService service

type ApiDeleteAIConfigRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	configKey string
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
func (a *AgentControlApiService) DeleteAIConfig(ctx context.Context, projectKey string, configKey string) ApiDeleteAIConfigRequest {
	return ApiDeleteAIConfigRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
	}
}

// Execute executes the request
func (a *AgentControlApiService) DeleteAIConfigExecute(r ApiDeleteAIConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.DeleteAIConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	ApiService AgentControlApi
	projectKey string
	configKey string
	variationKey string
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
func (a *AgentControlApiService) DeleteAIConfigVariation(ctx context.Context, projectKey string, configKey string, variationKey string) ApiDeleteAIConfigVariationRequest {
	return ApiDeleteAIConfigVariationRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
		variationKey: variationKey,
	}
}

// Execute executes the request
func (a *AgentControlApiService) DeleteAIConfigVariationExecute(r ApiDeleteAIConfigVariationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.DeleteAIConfigVariation")
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
	ApiService AgentControlApi
	projectKey string
	toolKey string
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
func (a *AgentControlApiService) DeleteAITool(ctx context.Context, projectKey string, toolKey string) ApiDeleteAIToolRequest {
	return ApiDeleteAIToolRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		toolKey: toolKey,
	}
}

// Execute executes the request
func (a *AgentControlApiService) DeleteAIToolExecute(r ApiDeleteAIToolRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.DeleteAITool")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-tools/{toolKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toolKey"+"}", url.PathEscape(parameterValueToString(r.toolKey, "toolKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	ApiService AgentControlApi
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
func (a *AgentControlApiService) DeleteAgentGraph(ctx context.Context, projectKey string, graphKey string) ApiDeleteAgentGraphRequest {
	return ApiDeleteAgentGraphRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		graphKey: graphKey,
	}
}

// Execute executes the request
func (a *AgentControlApiService) DeleteAgentGraphExecute(r ApiDeleteAgentGraphRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.DeleteAgentGraph")
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

type ApiDeleteAgentOptimizationRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	optimizationKey string
}

func (r ApiDeleteAgentOptimizationRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAgentOptimizationExecute(r)
}

/*
DeleteAgentOptimization Delete an agent optimization

Delete an existing agent optimization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param optimizationKey
 @return ApiDeleteAgentOptimizationRequest
*/
func (a *AgentControlApiService) DeleteAgentOptimization(ctx context.Context, projectKey string, optimizationKey string) ApiDeleteAgentOptimizationRequest {
	return ApiDeleteAgentOptimizationRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		optimizationKey: optimizationKey,
	}
}

// Execute executes the request
func (a *AgentControlApiService) DeleteAgentOptimizationExecute(r ApiDeleteAgentOptimizationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.DeleteAgentOptimization")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"optimizationKey"+"}", url.PathEscape(parameterValueToString(r.optimizationKey, "optimizationKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiDeleteAgentOptimizationRunRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	lDAPIVersion *string
	projectKey string
	optimizationKey string
	runId string
}

// Version of the endpoint.
func (r ApiDeleteAgentOptimizationRunRequest) LDAPIVersion(lDAPIVersion string) ApiDeleteAgentOptimizationRunRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

func (r ApiDeleteAgentOptimizationRunRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAgentOptimizationRunExecute(r)
}

/*
DeleteAgentOptimizationRun Delete an agent optimization run

Delete all results for a specific run of an agent optimization. Returns a 404 if the optimization does not exist or if the run has no results. A run whose results have already been deleted is indistinguishable from a run that never existed, so repeating this request also returns a 404.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param optimizationKey
 @param runId
 @return ApiDeleteAgentOptimizationRunRequest
*/
func (a *AgentControlApiService) DeleteAgentOptimizationRun(ctx context.Context, projectKey string, optimizationKey string, runId string) ApiDeleteAgentOptimizationRunRequest {
	return ApiDeleteAgentOptimizationRunRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		optimizationKey: optimizationKey,
		runId: runId,
	}
}

// Execute executes the request
func (a *AgentControlApiService) DeleteAgentOptimizationRunExecute(r ApiDeleteAgentOptimizationRunRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.DeleteAgentOptimizationRun")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey}/runs/{runId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"optimizationKey"+"}", url.PathEscape(parameterValueToString(r.optimizationKey, "optimizationKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"runId"+"}", url.PathEscape(parameterValueToString(r.runId, "runId")), -1)

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

type ApiDeleteAgentSkillRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	skillKey string
}

func (r ApiDeleteAgentSkillRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAgentSkillExecute(r)
}

/*
DeleteAgentSkill Delete an agent skill

Delete an agent skill, including all of its versions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param skillKey
 @return ApiDeleteAgentSkillRequest
*/
func (a *AgentControlApiService) DeleteAgentSkill(ctx context.Context, projectKey string, skillKey string) ApiDeleteAgentSkillRequest {
	return ApiDeleteAgentSkillRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		skillKey: skillKey,
	}
}

// Execute executes the request
func (a *AgentControlApiService) DeleteAgentSkillExecute(r ApiDeleteAgentSkillRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.DeleteAgentSkill")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/skills/{skillKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"skillKey"+"}", url.PathEscape(parameterValueToString(r.skillKey, "skillKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		if localVarHTTPResponse.StatusCode == 409 {
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
	ApiService AgentControlApi
	projectKey string
	modelConfigKey string
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
func (a *AgentControlApiService) DeleteModelConfig(ctx context.Context, projectKey string, modelConfigKey string) ApiDeleteModelConfigRequest {
	return ApiDeleteModelConfigRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		modelConfigKey: modelConfigKey,
	}
}

// Execute executes the request
func (a *AgentControlApiService) DeleteModelConfigExecute(r ApiDeleteModelConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.DeleteModelConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/model-configs/{modelConfigKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"modelConfigKey"+"}", url.PathEscape(parameterValueToString(r.modelConfigKey, "modelConfigKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiDeletePromptSnippetRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	snippetKey string
}

func (r ApiDeletePromptSnippetRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePromptSnippetExecute(r)
}

/*
DeletePromptSnippet Delete a prompt snippet

Delete an existing prompt snippet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param snippetKey
 @return ApiDeletePromptSnippetRequest
*/
func (a *AgentControlApiService) DeletePromptSnippet(ctx context.Context, projectKey string, snippetKey string) ApiDeletePromptSnippetRequest {
	return ApiDeletePromptSnippetRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		snippetKey: snippetKey,
	}
}

// Execute executes the request
func (a *AgentControlApiService) DeletePromptSnippetExecute(r ApiDeletePromptSnippetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.DeletePromptSnippet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/prompt-snippets/{snippetKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snippetKey"+"}", url.PathEscape(parameterValueToString(r.snippetKey, "snippetKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	ApiService AgentControlApi
	projectKey string
	restrictedModelsRequest *RestrictedModelsRequest
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
func (a *AgentControlApiService) DeleteRestrictedModels(ctx context.Context, projectKey string) ApiDeleteRestrictedModelsRequest {
	return ApiDeleteRestrictedModelsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
func (a *AgentControlApiService) DeleteRestrictedModelsExecute(r ApiDeleteRestrictedModelsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.DeleteRestrictedModels")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/model-configs/restricted"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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
	ApiService AgentControlApi
	projectKey string
	configKey string
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
func (a *AgentControlApiService) GetAIConfig(ctx context.Context, projectKey string, configKey string) ApiGetAIConfigRequest {
	return ApiGetAIConfigRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
	}
}

// Execute executes the request
//  @return AIConfig
func (a *AgentControlApiService) GetAIConfigExecute(r ApiGetAIConfigRequest) (*AIConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AIConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.GetAIConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	ApiService AgentControlApi
	projectKey string
	configKey string
	from *int32
	to *int32
	env *string
	contextKind *string
	contextKey *string
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

// A context kind. Only metrics from events that include a context of this kind are included. Required if &#x60;contextKey&#x60; is provided.
func (r ApiGetAIConfigMetricsRequest) ContextKind(contextKind string) ApiGetAIConfigMetricsRequest {
	r.contextKind = &contextKind
	return r
}

// A context key. Only metrics from events whose context of the &#x60;contextKind&#x60; kind has this key are included. Requires &#x60;contextKind&#x60;.
func (r ApiGetAIConfigMetricsRequest) ContextKey(contextKey string) ApiGetAIConfigMetricsRequest {
	r.contextKey = &contextKey
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
func (a *AgentControlApiService) GetAIConfigMetrics(ctx context.Context, projectKey string, configKey string) ApiGetAIConfigMetricsRequest {
	return ApiGetAIConfigMetricsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
	}
}

// Execute executes the request
//  @return Metrics
func (a *AgentControlApiService) GetAIConfigMetricsExecute(r ApiGetAIConfigMetricsRequest) (*Metrics, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Metrics
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.GetAIConfigMetrics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}/metrics"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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
	if r.contextKind != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "contextKind", r.contextKind, "form", "")
	}
	if r.contextKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "contextKey", r.contextKey, "form", "")
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
	ApiService AgentControlApi
	projectKey string
	configKey string
	from *int32
	to *int32
	env *string
	contextKind *string
	contextKey *string
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

// A context kind. Only metrics from events that include a context of this kind are included. Required if &#x60;contextKey&#x60; is provided.
func (r ApiGetAIConfigMetricsByVariationRequest) ContextKind(contextKind string) ApiGetAIConfigMetricsByVariationRequest {
	r.contextKind = &contextKind
	return r
}

// A context key. Only metrics from events whose context of the &#x60;contextKind&#x60; kind has this key are included. Requires &#x60;contextKind&#x60;.
func (r ApiGetAIConfigMetricsByVariationRequest) ContextKey(contextKey string) ApiGetAIConfigMetricsByVariationRequest {
	r.contextKey = &contextKey
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
func (a *AgentControlApiService) GetAIConfigMetricsByVariation(ctx context.Context, projectKey string, configKey string) ApiGetAIConfigMetricsByVariationRequest {
	return ApiGetAIConfigMetricsByVariationRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
	}
}

// Execute executes the request
//  @return []MetricByVariation
func (a *AgentControlApiService) GetAIConfigMetricsByVariationExecute(r ApiGetAIConfigMetricsByVariationRequest) ([]MetricByVariation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []MetricByVariation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.GetAIConfigMetricsByVariation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}/metrics-by-variation"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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
	if r.contextKind != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "contextKind", r.contextKind, "form", "")
	}
	if r.contextKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "contextKey", r.contextKey, "form", "")
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

type ApiGetAIConfigQuickStatsRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	env *string
}

// An environment key. Only metrics from this environment will be included.
func (r ApiGetAIConfigQuickStatsRequest) Env(env string) ApiGetAIConfigQuickStatsRequest {
	r.env = &env
	return r
}

func (r ApiGetAIConfigQuickStatsRequest) Execute() (*QuickStats, *http.Response, error) {
	return r.ApiService.GetAIConfigQuickStatsExecute(r)
}

/*
GetAIConfigQuickStats Get AI Config quick stats

Retrieve aggregate quick stats for AI Configs in a project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiGetAIConfigQuickStatsRequest
*/
func (a *AgentControlApiService) GetAIConfigQuickStats(ctx context.Context, projectKey string) ApiGetAIConfigQuickStatsRequest {
	return ApiGetAIConfigQuickStatsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return QuickStats
func (a *AgentControlApiService) GetAIConfigQuickStatsExecute(r ApiGetAIConfigQuickStatsRequest) (*QuickStats, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QuickStats
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.GetAIConfigQuickStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/quick-stats"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.env == nil {
		return localVarReturnValue, nil, reportError("env is required and must be specified")
	}

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
	ApiService AgentControlApi
	projectKey string
	configKey string
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
func (a *AgentControlApiService) GetAIConfigTargeting(ctx context.Context, projectKey string, configKey string) ApiGetAIConfigTargetingRequest {
	return ApiGetAIConfigTargetingRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
	}
}

// Execute executes the request
//  @return AIConfigTargeting
func (a *AgentControlApiService) GetAIConfigTargetingExecute(r ApiGetAIConfigTargetingRequest) (*AIConfigTargeting, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AIConfigTargeting
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.GetAIConfigTargeting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}/targeting"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	ApiService AgentControlApi
	projectKey string
	configKey string
	variationKey string
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
func (a *AgentControlApiService) GetAIConfigVariation(ctx context.Context, projectKey string, configKey string, variationKey string) ApiGetAIConfigVariationRequest {
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
func (a *AgentControlApiService) GetAIConfigVariationExecute(r ApiGetAIConfigVariationRequest) (*AIConfigVariationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AIConfigVariationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.GetAIConfigVariation")
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
	ApiService AgentControlApi
	projectKey string
	sort *string
	limit *int32
	offset *int32
	filter *string
}

// A sort to apply to the list of AgentControl configs.
func (r ApiGetAIConfigsRequest) Sort(sort string) ApiGetAIConfigsRequest {
	r.sort = &sort
	return r
}

// The number of resources to return.
func (r ApiGetAIConfigsRequest) Limit(limit int32) ApiGetAIConfigsRequest {
	r.limit = &limit
	return r
}

// Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;.
func (r ApiGetAIConfigsRequest) Offset(offset int32) ApiGetAIConfigsRequest {
	r.offset = &offset
	return r
}

// A filter to apply to the list.
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
func (a *AgentControlApiService) GetAIConfigs(ctx context.Context, projectKey string) ApiGetAIConfigsRequest {
	return ApiGetAIConfigsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return AIConfigs
func (a *AgentControlApiService) GetAIConfigsExecute(r ApiGetAIConfigsRequest) (*AIConfigs, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AIConfigs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.GetAIConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	ApiService AgentControlApi
	projectKey string
	toolKey string
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
func (a *AgentControlApiService) GetAITool(ctx context.Context, projectKey string, toolKey string) ApiGetAIToolRequest {
	return ApiGetAIToolRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		toolKey: toolKey,
	}
}

// Execute executes the request
//  @return AITool
func (a *AgentControlApiService) GetAIToolExecute(r ApiGetAIToolRequest) (*AITool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AITool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.GetAITool")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-tools/{toolKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toolKey"+"}", url.PathEscape(parameterValueToString(r.toolKey, "toolKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	ApiService AgentControlApi
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
func (a *AgentControlApiService) GetAgentGraph(ctx context.Context, projectKey string, graphKey string) ApiGetAgentGraphRequest {
	return ApiGetAgentGraphRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		graphKey: graphKey,
	}
}

// Execute executes the request
//  @return AgentGraph
func (a *AgentControlApiService) GetAgentGraphExecute(r ApiGetAgentGraphRequest) (*AgentGraph, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentGraph
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.GetAgentGraph")
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

type ApiGetAgentOptimizationRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	optimizationKey string
}

func (r ApiGetAgentOptimizationRequest) Execute() (*AgentOptimization, *http.Response, error) {
	return r.ApiService.GetAgentOptimizationExecute(r)
}

/*
GetAgentOptimization Get an agent optimization

Retrieve a specific agent optimization by its key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param optimizationKey
 @return ApiGetAgentOptimizationRequest
*/
func (a *AgentControlApiService) GetAgentOptimization(ctx context.Context, projectKey string, optimizationKey string) ApiGetAgentOptimizationRequest {
	return ApiGetAgentOptimizationRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		optimizationKey: optimizationKey,
	}
}

// Execute executes the request
//  @return AgentOptimization
func (a *AgentControlApiService) GetAgentOptimizationExecute(r ApiGetAgentOptimizationRequest) (*AgentOptimization, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentOptimization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.GetAgentOptimization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"optimizationKey"+"}", url.PathEscape(parameterValueToString(r.optimizationKey, "optimizationKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetAgentSkillRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	skillKey string
}

func (r ApiGetAgentSkillRequest) Execute() (*AgentSkill, *http.Response, error) {
	return r.ApiService.GetAgentSkillExecute(r)
}

/*
GetAgentSkill Get an agent skill

Retrieve a specific agent skill by its key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param skillKey
 @return ApiGetAgentSkillRequest
*/
func (a *AgentControlApiService) GetAgentSkill(ctx context.Context, projectKey string, skillKey string) ApiGetAgentSkillRequest {
	return ApiGetAgentSkillRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		skillKey: skillKey,
	}
}

// Execute executes the request
//  @return AgentSkill
func (a *AgentControlApiService) GetAgentSkillExecute(r ApiGetAgentSkillRequest) (*AgentSkill, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentSkill
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.GetAgentSkill")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/skills/{skillKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"skillKey"+"}", url.PathEscape(parameterValueToString(r.skillKey, "skillKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetModelConfigRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	modelConfigKey string
	version *int32
}

// Specific model config version to return. Omit to return the latest version.
func (r ApiGetModelConfigRequest) Version(version int32) ApiGetModelConfigRequest {
	r.version = &version
	return r
}

func (r ApiGetModelConfigRequest) Execute() (*ModelConfig, *http.Response, error) {
	return r.ApiService.GetModelConfigExecute(r)
}

/*
GetModelConfig Get AI model config

Get an AI model config by key. Specify a version to retrieve a historical custom model config version. Global model configs accept version 0 or 1 and return their current representation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param modelConfigKey
 @return ApiGetModelConfigRequest
*/
func (a *AgentControlApiService) GetModelConfig(ctx context.Context, projectKey string, modelConfigKey string) ApiGetModelConfigRequest {
	return ApiGetModelConfigRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		modelConfigKey: modelConfigKey,
	}
}

// Execute executes the request
//  @return ModelConfig
func (a *AgentControlApiService) GetModelConfigExecute(r ApiGetModelConfigRequest) (*ModelConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.GetModelConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/model-configs/{modelConfigKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"modelConfigKey"+"}", url.PathEscape(parameterValueToString(r.modelConfigKey, "modelConfigKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "form", "")
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

type ApiGetPromptSnippetRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	snippetKey string
}

func (r ApiGetPromptSnippetRequest) Execute() (*PromptSnippet, *http.Response, error) {
	return r.ApiService.GetPromptSnippetExecute(r)
}

/*
GetPromptSnippet Get a prompt snippet

Retrieve a specific prompt snippet by its key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param snippetKey
 @return ApiGetPromptSnippetRequest
*/
func (a *AgentControlApiService) GetPromptSnippet(ctx context.Context, projectKey string, snippetKey string) ApiGetPromptSnippetRequest {
	return ApiGetPromptSnippetRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		snippetKey: snippetKey,
	}
}

// Execute executes the request
//  @return PromptSnippet
func (a *AgentControlApiService) GetPromptSnippetExecute(r ApiGetPromptSnippetRequest) (*PromptSnippet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PromptSnippet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.GetPromptSnippet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/prompt-snippets/{snippetKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snippetKey"+"}", url.PathEscape(parameterValueToString(r.snippetKey, "snippetKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListAIToolVersionsRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	toolKey string
	sort *string
	limit *int32
	offset *int32
}

// A sort to apply to the list of AgentControl configs.
func (r ApiListAIToolVersionsRequest) Sort(sort string) ApiListAIToolVersionsRequest {
	r.sort = &sort
	return r
}

// The number of resources to return.
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
func (a *AgentControlApiService) ListAIToolVersions(ctx context.Context, projectKey string, toolKey string) ApiListAIToolVersionsRequest {
	return ApiListAIToolVersionsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		toolKey: toolKey,
	}
}

// Execute executes the request
//  @return AITools
func (a *AgentControlApiService) ListAIToolVersionsExecute(r ApiListAIToolVersionsRequest) (*AITools, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AITools
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.ListAIToolVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-tools/{toolKey}/versions"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toolKey"+"}", url.PathEscape(parameterValueToString(r.toolKey, "toolKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	ApiService AgentControlApi
	projectKey string
	sort *string
	limit *int32
	offset *int32
	filter *string
}

// A sort to apply to the list of AgentControl configs.
func (r ApiListAIToolsRequest) Sort(sort string) ApiListAIToolsRequest {
	r.sort = &sort
	return r
}

// The number of resources to return.
func (r ApiListAIToolsRequest) Limit(limit int32) ApiListAIToolsRequest {
	r.limit = &limit
	return r
}

// Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;.
func (r ApiListAIToolsRequest) Offset(offset int32) ApiListAIToolsRequest {
	r.offset = &offset
	return r
}

// A filter to apply to the list.
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
func (a *AgentControlApiService) ListAITools(ctx context.Context, projectKey string) ApiListAIToolsRequest {
	return ApiListAIToolsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return AITools
func (a *AgentControlApiService) ListAIToolsExecute(r ApiListAIToolsRequest) (*AITools, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AITools
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.ListAITools")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-tools"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	ApiService AgentControlApi
	lDAPIVersion *string
	projectKey string
	limit *int32
	offset *int32
	filter *string
}

// Version of the endpoint.
func (r ApiListAgentGraphsRequest) LDAPIVersion(lDAPIVersion string) ApiListAgentGraphsRequest {
	r.lDAPIVersion = &lDAPIVersion
	return r
}

// The number of resources to return.
func (r ApiListAgentGraphsRequest) Limit(limit int32) ApiListAgentGraphsRequest {
	r.limit = &limit
	return r
}

// Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;.
func (r ApiListAgentGraphsRequest) Offset(offset int32) ApiListAgentGraphsRequest {
	r.offset = &offset
	return r
}

// A filter to apply to the list.
func (r ApiListAgentGraphsRequest) Filter(filter string) ApiListAgentGraphsRequest {
	r.filter = &filter
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
func (a *AgentControlApiService) ListAgentGraphs(ctx context.Context, projectKey string) ApiListAgentGraphsRequest {
	return ApiListAgentGraphsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return AgentGraphs
func (a *AgentControlApiService) ListAgentGraphsExecute(r ApiListAgentGraphsRequest) (*AgentGraphs, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentGraphs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.ListAgentGraphs")
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

type ApiListAgentOptimizationResultsRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	optimizationKey string
	limit *int32
	offset *int32
}

// The number of resources to return.
func (r ApiListAgentOptimizationResultsRequest) Limit(limit int32) ApiListAgentOptimizationResultsRequest {
	r.limit = &limit
	return r
}

// Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;.
func (r ApiListAgentOptimizationResultsRequest) Offset(offset int32) ApiListAgentOptimizationResultsRequest {
	r.offset = &offset
	return r
}

func (r ApiListAgentOptimizationResultsRequest) Execute() (*AgentOptimizationResults, *http.Response, error) {
	return r.ApiService.ListAgentOptimizationResultsExecute(r)
}

/*
ListAgentOptimizationResults List agent optimization runs

Get the most recent result for each unique run of an agent optimization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param optimizationKey
 @return ApiListAgentOptimizationResultsRequest
*/
func (a *AgentControlApiService) ListAgentOptimizationResults(ctx context.Context, projectKey string, optimizationKey string) ApiListAgentOptimizationResultsRequest {
	return ApiListAgentOptimizationResultsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		optimizationKey: optimizationKey,
	}
}

// Execute executes the request
//  @return AgentOptimizationResults
func (a *AgentControlApiService) ListAgentOptimizationResultsExecute(r ApiListAgentOptimizationResultsRequest) (*AgentOptimizationResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentOptimizationResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.ListAgentOptimizationResults")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey}/results"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"optimizationKey"+"}", url.PathEscape(parameterValueToString(r.optimizationKey, "optimizationKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListAgentOptimizationResultsByRunIdRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	optimizationKey string
	runId string
}

func (r ApiListAgentOptimizationResultsByRunIdRequest) Execute() (*AgentOptimizationResults, *http.Response, error) {
	return r.ApiService.ListAgentOptimizationResultsByRunIdExecute(r)
}

/*
ListAgentOptimizationResultsByRunId List agent optimization results for a run

Get all results for a specific run of an agent optimization, sorted by iteration descending.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param optimizationKey
 @param runId
 @return ApiListAgentOptimizationResultsByRunIdRequest
*/
func (a *AgentControlApiService) ListAgentOptimizationResultsByRunId(ctx context.Context, projectKey string, optimizationKey string, runId string) ApiListAgentOptimizationResultsByRunIdRequest {
	return ApiListAgentOptimizationResultsByRunIdRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		optimizationKey: optimizationKey,
		runId: runId,
	}
}

// Execute executes the request
//  @return AgentOptimizationResults
func (a *AgentControlApiService) ListAgentOptimizationResultsByRunIdExecute(r ApiListAgentOptimizationResultsByRunIdRequest) (*AgentOptimizationResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentOptimizationResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.ListAgentOptimizationResultsByRunId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey}/runs/{runId}/results"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"optimizationKey"+"}", url.PathEscape(parameterValueToString(r.optimizationKey, "optimizationKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"runId"+"}", url.PathEscape(parameterValueToString(r.runId, "runId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListAgentOptimizationsRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	limit *int32
	offset *int32
	filter *string
}

// The number of resources to return.
func (r ApiListAgentOptimizationsRequest) Limit(limit int32) ApiListAgentOptimizationsRequest {
	r.limit = &limit
	return r
}

// Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;.
func (r ApiListAgentOptimizationsRequest) Offset(offset int32) ApiListAgentOptimizationsRequest {
	r.offset = &offset
	return r
}

// A filter to apply to the list.
func (r ApiListAgentOptimizationsRequest) Filter(filter string) ApiListAgentOptimizationsRequest {
	r.filter = &filter
	return r
}

func (r ApiListAgentOptimizationsRequest) Execute() (*AgentOptimizations, *http.Response, error) {
	return r.ApiService.ListAgentOptimizationsExecute(r)
}

/*
ListAgentOptimizations List agent optimizations

Get a list of all agent optimizations in the given project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiListAgentOptimizationsRequest
*/
func (a *AgentControlApiService) ListAgentOptimizations(ctx context.Context, projectKey string) ApiListAgentOptimizationsRequest {
	return ApiListAgentOptimizationsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return AgentOptimizations
func (a *AgentControlApiService) ListAgentOptimizationsExecute(r ApiListAgentOptimizationsRequest) (*AgentOptimizations, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentOptimizations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.ListAgentOptimizations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/agent-optimizations"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListAgentSkillReferencesRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	skillKey string
	limit *int32
	offset *int32
}

// The number of resources to return.
func (r ApiListAgentSkillReferencesRequest) Limit(limit int32) ApiListAgentSkillReferencesRequest {
	r.limit = &limit
	return r
}

// Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;.
func (r ApiListAgentSkillReferencesRequest) Offset(offset int32) ApiListAgentSkillReferencesRequest {
	r.offset = &offset
	return r
}

func (r ApiListAgentSkillReferencesRequest) Execute() (*AgentSkillReferences, *http.Response, error) {
	return r.ApiService.ListAgentSkillReferencesExecute(r)
}

/*
ListAgentSkillReferences List agent skill references

Get all config variations that currently reference this agent skill.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param skillKey
 @return ApiListAgentSkillReferencesRequest
*/
func (a *AgentControlApiService) ListAgentSkillReferences(ctx context.Context, projectKey string, skillKey string) ApiListAgentSkillReferencesRequest {
	return ApiListAgentSkillReferencesRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		skillKey: skillKey,
	}
}

// Execute executes the request
//  @return AgentSkillReferences
func (a *AgentControlApiService) ListAgentSkillReferencesExecute(r ApiListAgentSkillReferencesRequest) (*AgentSkillReferences, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentSkillReferences
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.ListAgentSkillReferences")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/skills/{skillKey}/references"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"skillKey"+"}", url.PathEscape(parameterValueToString(r.skillKey, "skillKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListAgentSkillVersionsRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	skillKey string
	limit *int32
	offset *int32
}

// The number of resources to return.
func (r ApiListAgentSkillVersionsRequest) Limit(limit int32) ApiListAgentSkillVersionsRequest {
	r.limit = &limit
	return r
}

// Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;.
func (r ApiListAgentSkillVersionsRequest) Offset(offset int32) ApiListAgentSkillVersionsRequest {
	r.offset = &offset
	return r
}

func (r ApiListAgentSkillVersionsRequest) Execute() (*AgentSkills, *http.Response, error) {
	return r.ApiService.ListAgentSkillVersionsExecute(r)
}

/*
ListAgentSkillVersions List agent skill versions

Get all versions of an agent skill in the given project, ordered by version descending.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param skillKey
 @return ApiListAgentSkillVersionsRequest
*/
func (a *AgentControlApiService) ListAgentSkillVersions(ctx context.Context, projectKey string, skillKey string) ApiListAgentSkillVersionsRequest {
	return ApiListAgentSkillVersionsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		skillKey: skillKey,
	}
}

// Execute executes the request
//  @return AgentSkills
func (a *AgentControlApiService) ListAgentSkillVersionsExecute(r ApiListAgentSkillVersionsRequest) (*AgentSkills, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentSkills
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.ListAgentSkillVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/skills/{skillKey}/versions"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"skillKey"+"}", url.PathEscape(parameterValueToString(r.skillKey, "skillKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListAgentSkillsRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	limit *int32
	offset *int32
	filter *string
}

// The number of resources to return.
func (r ApiListAgentSkillsRequest) Limit(limit int32) ApiListAgentSkillsRequest {
	r.limit = &limit
	return r
}

// Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;.
func (r ApiListAgentSkillsRequest) Offset(offset int32) ApiListAgentSkillsRequest {
	r.offset = &offset
	return r
}

// A filter to apply to the list.
func (r ApiListAgentSkillsRequest) Filter(filter string) ApiListAgentSkillsRequest {
	r.filter = &filter
	return r
}

func (r ApiListAgentSkillsRequest) Execute() (*AgentSkills, *http.Response, error) {
	return r.ApiService.ListAgentSkillsExecute(r)
}

/*
ListAgentSkills List agent skills

Get a list of all agent skills in the given project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiListAgentSkillsRequest
*/
func (a *AgentControlApiService) ListAgentSkills(ctx context.Context, projectKey string) ApiListAgentSkillsRequest {
	return ApiListAgentSkillsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return AgentSkills
func (a *AgentControlApiService) ListAgentSkillsExecute(r ApiListAgentSkillsRequest) (*AgentSkills, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentSkills
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.ListAgentSkills")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/skills"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListAllAgentOptimizationResultsRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	optimizationKey string
	limit *int32
	offset *int32
}

// The number of resources to return.
func (r ApiListAllAgentOptimizationResultsRequest) Limit(limit int32) ApiListAllAgentOptimizationResultsRequest {
	r.limit = &limit
	return r
}

// Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;.
func (r ApiListAllAgentOptimizationResultsRequest) Offset(offset int32) ApiListAllAgentOptimizationResultsRequest {
	r.offset = &offset
	return r
}

func (r ApiListAllAgentOptimizationResultsRequest) Execute() (*AgentOptimizationResults, *http.Response, error) {
	return r.ApiService.ListAllAgentOptimizationResultsExecute(r)
}

/*
ListAllAgentOptimizationResults List all agent optimization results across versions

Get all results across all versions of an agent optimization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param optimizationKey
 @return ApiListAllAgentOptimizationResultsRequest
*/
func (a *AgentControlApiService) ListAllAgentOptimizationResults(ctx context.Context, projectKey string, optimizationKey string) ApiListAllAgentOptimizationResultsRequest {
	return ApiListAllAgentOptimizationResultsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		optimizationKey: optimizationKey,
	}
}

// Execute executes the request
//  @return AgentOptimizationResults
func (a *AgentControlApiService) ListAllAgentOptimizationResultsExecute(r ApiListAllAgentOptimizationResultsRequest) (*AgentOptimizationResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentOptimizationResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.ListAllAgentOptimizationResults")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey}/all-results"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"optimizationKey"+"}", url.PathEscape(parameterValueToString(r.optimizationKey, "optimizationKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListModelConfigVersionsRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	modelConfigKey string
	limit *int32
	offset *int32
}

// The number of resources to return.
func (r ApiListModelConfigVersionsRequest) Limit(limit int32) ApiListModelConfigVersionsRequest {
	r.limit = &limit
	return r
}

// Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;.
func (r ApiListModelConfigVersionsRequest) Offset(offset int32) ApiListModelConfigVersionsRequest {
	r.offset = &offset
	return r
}

func (r ApiListModelConfigVersionsRequest) Execute() (*ModelConfigs, *http.Response, error) {
	return r.ApiService.ListModelConfigVersionsExecute(r)
}

/*
ListModelConfigVersions List AI model config versions

Get a paginated list of all available versions of an AI model config, ordered from newest to oldest. Global model configs return their single current representation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param modelConfigKey
 @return ApiListModelConfigVersionsRequest
*/
func (a *AgentControlApiService) ListModelConfigVersions(ctx context.Context, projectKey string, modelConfigKey string) ApiListModelConfigVersionsRequest {
	return ApiListModelConfigVersionsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		modelConfigKey: modelConfigKey,
	}
}

// Execute executes the request
//  @return ModelConfigs
func (a *AgentControlApiService) ListModelConfigVersionsExecute(r ApiListModelConfigVersionsRequest) (*ModelConfigs, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelConfigs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.ListModelConfigVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/model-configs/{modelConfigKey}/versions"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"modelConfigKey"+"}", url.PathEscape(parameterValueToString(r.modelConfigKey, "modelConfigKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListModelConfigsRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	restricted *bool
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
func (a *AgentControlApiService) ListModelConfigs(ctx context.Context, projectKey string) ApiListModelConfigsRequest {
	return ApiListModelConfigsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return []ModelConfig
func (a *AgentControlApiService) ListModelConfigsExecute(r ApiListModelConfigsRequest) ([]ModelConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ModelConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.ListModelConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/model-configs"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListPromptSnippetReferencesRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	snippetKey string
	limit *int32
	offset *int32
}

// The number of resources to return.
func (r ApiListPromptSnippetReferencesRequest) Limit(limit int32) ApiListPromptSnippetReferencesRequest {
	r.limit = &limit
	return r
}

// Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;.
func (r ApiListPromptSnippetReferencesRequest) Offset(offset int32) ApiListPromptSnippetReferencesRequest {
	r.offset = &offset
	return r
}

func (r ApiListPromptSnippetReferencesRequest) Execute() (*SnippetReferences, *http.Response, error) {
	return r.ApiService.ListPromptSnippetReferencesExecute(r)
}

/*
ListPromptSnippetReferences List prompt snippet references

Get all config variations that currently reference this prompt snippet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param snippetKey
 @return ApiListPromptSnippetReferencesRequest
*/
func (a *AgentControlApiService) ListPromptSnippetReferences(ctx context.Context, projectKey string, snippetKey string) ApiListPromptSnippetReferencesRequest {
	return ApiListPromptSnippetReferencesRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		snippetKey: snippetKey,
	}
}

// Execute executes the request
//  @return SnippetReferences
func (a *AgentControlApiService) ListPromptSnippetReferencesExecute(r ApiListPromptSnippetReferencesRequest) (*SnippetReferences, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SnippetReferences
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.ListPromptSnippetReferences")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/prompt-snippets/{snippetKey}/references"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snippetKey"+"}", url.PathEscape(parameterValueToString(r.snippetKey, "snippetKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListPromptSnippetVersionsRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	snippetKey string
	limit *int32
	offset *int32
}

// The number of resources to return.
func (r ApiListPromptSnippetVersionsRequest) Limit(limit int32) ApiListPromptSnippetVersionsRequest {
	r.limit = &limit
	return r
}

// Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;.
func (r ApiListPromptSnippetVersionsRequest) Offset(offset int32) ApiListPromptSnippetVersionsRequest {
	r.offset = &offset
	return r
}

func (r ApiListPromptSnippetVersionsRequest) Execute() (*PromptSnippets, *http.Response, error) {
	return r.ApiService.ListPromptSnippetVersionsExecute(r)
}

/*
ListPromptSnippetVersions List prompt snippet versions

Get all versions of a prompt snippet in the given project, ordered by version descending.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param snippetKey
 @return ApiListPromptSnippetVersionsRequest
*/
func (a *AgentControlApiService) ListPromptSnippetVersions(ctx context.Context, projectKey string, snippetKey string) ApiListPromptSnippetVersionsRequest {
	return ApiListPromptSnippetVersionsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		snippetKey: snippetKey,
	}
}

// Execute executes the request
//  @return PromptSnippets
func (a *AgentControlApiService) ListPromptSnippetVersionsExecute(r ApiListPromptSnippetVersionsRequest) (*PromptSnippets, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PromptSnippets
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.ListPromptSnippetVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/prompt-snippets/{snippetKey}/versions"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snippetKey"+"}", url.PathEscape(parameterValueToString(r.snippetKey, "snippetKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiListPromptSnippetsRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	limit *int32
	offset *int32
	filter *string
}

// The number of resources to return.
func (r ApiListPromptSnippetsRequest) Limit(limit int32) ApiListPromptSnippetsRequest {
	r.limit = &limit
	return r
}

// Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;.
func (r ApiListPromptSnippetsRequest) Offset(offset int32) ApiListPromptSnippetsRequest {
	r.offset = &offset
	return r
}

// A filter to apply to the list.
func (r ApiListPromptSnippetsRequest) Filter(filter string) ApiListPromptSnippetsRequest {
	r.filter = &filter
	return r
}

func (r ApiListPromptSnippetsRequest) Execute() (*PromptSnippets, *http.Response, error) {
	return r.ApiService.ListPromptSnippetsExecute(r)
}

/*
ListPromptSnippets List prompt snippets

Get a list of all prompt snippets in the given project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiListPromptSnippetsRequest
*/
func (a *AgentControlApiService) ListPromptSnippets(ctx context.Context, projectKey string) ApiListPromptSnippetsRequest {
	return ApiListPromptSnippetsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return PromptSnippets
func (a *AgentControlApiService) ListPromptSnippetsExecute(r ApiListPromptSnippetsRequest) (*PromptSnippets, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PromptSnippets
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.ListPromptSnippets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/prompt-snippets"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	ApiService AgentControlApi
	projectKey string
	configKey string
	aIConfigPatch *AIConfigPatch
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
func (a *AgentControlApiService) PatchAIConfig(ctx context.Context, projectKey string, configKey string) ApiPatchAIConfigRequest {
	return ApiPatchAIConfigRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
	}
}

// Execute executes the request
//  @return AIConfig
func (a *AgentControlApiService) PatchAIConfigExecute(r ApiPatchAIConfigRequest) (*AIConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AIConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PatchAIConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	ApiService AgentControlApi
	projectKey string
	configKey string
	aIConfigTargetingPatch *AIConfigTargetingPatch
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
func (a *AgentControlApiService) PatchAIConfigTargeting(ctx context.Context, projectKey string, configKey string) ApiPatchAIConfigTargetingRequest {
	return ApiPatchAIConfigTargetingRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
	}
}

// Execute executes the request
//  @return AIConfigTargeting
func (a *AgentControlApiService) PatchAIConfigTargetingExecute(r ApiPatchAIConfigTargetingRequest) (*AIConfigTargeting, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AIConfigTargeting
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PatchAIConfigTargeting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}/targeting"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	ApiService AgentControlApi
	projectKey string
	configKey string
	variationKey string
	aIConfigVariationPatch *AIConfigVariationPatch
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
func (a *AgentControlApiService) PatchAIConfigVariation(ctx context.Context, projectKey string, configKey string, variationKey string) ApiPatchAIConfigVariationRequest {
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
func (a *AgentControlApiService) PatchAIConfigVariationExecute(r ApiPatchAIConfigVariationRequest) (*AIConfigVariation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AIConfigVariation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PatchAIConfigVariation")
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
	ApiService AgentControlApi
	projectKey string
	toolKey string
	aIToolPatch *AIToolPatch
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
func (a *AgentControlApiService) PatchAITool(ctx context.Context, projectKey string, toolKey string) ApiPatchAIToolRequest {
	return ApiPatchAIToolRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		toolKey: toolKey,
	}
}

// Execute executes the request
//  @return AITool
func (a *AgentControlApiService) PatchAIToolExecute(r ApiPatchAIToolRequest) (*AITool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AITool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PatchAITool")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-tools/{toolKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toolKey"+"}", url.PathEscape(parameterValueToString(r.toolKey, "toolKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	ApiService AgentControlApi
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
func (a *AgentControlApiService) PatchAgentGraph(ctx context.Context, projectKey string, graphKey string) ApiPatchAgentGraphRequest {
	return ApiPatchAgentGraphRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		graphKey: graphKey,
	}
}

// Execute executes the request
//  @return AgentGraph
func (a *AgentControlApiService) PatchAgentGraphExecute(r ApiPatchAgentGraphRequest) (*AgentGraph, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentGraph
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PatchAgentGraph")
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

type ApiPatchAgentOptimizationRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	optimizationKey string
	agentOptimizationPatch *AgentOptimizationPatch
}

// Agent optimization fields to update
func (r ApiPatchAgentOptimizationRequest) AgentOptimizationPatch(agentOptimizationPatch AgentOptimizationPatch) ApiPatchAgentOptimizationRequest {
	r.agentOptimizationPatch = &agentOptimizationPatch
	return r
}

func (r ApiPatchAgentOptimizationRequest) Execute() (*AgentOptimization, *http.Response, error) {
	return r.ApiService.PatchAgentOptimizationExecute(r)
}

/*
PatchAgentOptimization Update an agent optimization

Update an existing agent optimization. Creates a new version of the optimization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param optimizationKey
 @return ApiPatchAgentOptimizationRequest
*/
func (a *AgentControlApiService) PatchAgentOptimization(ctx context.Context, projectKey string, optimizationKey string) ApiPatchAgentOptimizationRequest {
	return ApiPatchAgentOptimizationRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		optimizationKey: optimizationKey,
	}
}

// Execute executes the request
//  @return AgentOptimization
func (a *AgentControlApiService) PatchAgentOptimizationExecute(r ApiPatchAgentOptimizationRequest) (*AgentOptimization, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentOptimization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PatchAgentOptimization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"optimizationKey"+"}", url.PathEscape(parameterValueToString(r.optimizationKey, "optimizationKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentOptimizationPatch == nil {
		return localVarReturnValue, nil, reportError("agentOptimizationPatch is required and must be specified")
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
	// body params
	localVarPostBody = r.agentOptimizationPatch
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
		if localVarHTTPResponse.StatusCode == 409 {
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

type ApiPatchAgentOptimizationResultRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	optimizationKey string
	resultId string
	agentOptimizationResultPatch *AgentOptimizationResultPatch
}

// Agent optimization result fields to update
func (r ApiPatchAgentOptimizationResultRequest) AgentOptimizationResultPatch(agentOptimizationResultPatch AgentOptimizationResultPatch) ApiPatchAgentOptimizationResultRequest {
	r.agentOptimizationResultPatch = &agentOptimizationResultPatch
	return r
}

func (r ApiPatchAgentOptimizationResultRequest) Execute() (*AgentOptimizationResult, *http.Response, error) {
	return r.ApiService.PatchAgentOptimizationResultExecute(r)
}

/*
PatchAgentOptimizationResult Update an agent optimization result

Update status, activity, and AI request metadata fields on an agent optimization result.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param optimizationKey
 @param resultId
 @return ApiPatchAgentOptimizationResultRequest
*/
func (a *AgentControlApiService) PatchAgentOptimizationResult(ctx context.Context, projectKey string, optimizationKey string, resultId string) ApiPatchAgentOptimizationResultRequest {
	return ApiPatchAgentOptimizationResultRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		optimizationKey: optimizationKey,
		resultId: resultId,
	}
}

// Execute executes the request
//  @return AgentOptimizationResult
func (a *AgentControlApiService) PatchAgentOptimizationResultExecute(r ApiPatchAgentOptimizationResultRequest) (*AgentOptimizationResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentOptimizationResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PatchAgentOptimizationResult")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey}/results/{resultId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"optimizationKey"+"}", url.PathEscape(parameterValueToString(r.optimizationKey, "optimizationKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resultId"+"}", url.PathEscape(parameterValueToString(r.resultId, "resultId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentOptimizationResultPatch == nil {
		return localVarReturnValue, nil, reportError("agentOptimizationResultPatch is required and must be specified")
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
	// body params
	localVarPostBody = r.agentOptimizationResultPatch
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

type ApiPatchAgentSkillRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	skillKey string
	agentSkillPatch *AgentSkillPatch
}

// Agent skill fields to update
func (r ApiPatchAgentSkillRequest) AgentSkillPatch(agentSkillPatch AgentSkillPatch) ApiPatchAgentSkillRequest {
	r.agentSkillPatch = &agentSkillPatch
	return r
}

func (r ApiPatchAgentSkillRequest) Execute() (*AgentSkill, *http.Response, error) {
	return r.ApiService.PatchAgentSkillExecute(r)
}

/*
PatchAgentSkill Update an agent skill

Update an existing agent skill. Creates a new version of the skill.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param skillKey
 @return ApiPatchAgentSkillRequest
*/
func (a *AgentControlApiService) PatchAgentSkill(ctx context.Context, projectKey string, skillKey string) ApiPatchAgentSkillRequest {
	return ApiPatchAgentSkillRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		skillKey: skillKey,
	}
}

// Execute executes the request
//  @return AgentSkill
func (a *AgentControlApiService) PatchAgentSkillExecute(r ApiPatchAgentSkillRequest) (*AgentSkill, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentSkill
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PatchAgentSkill")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/skills/{skillKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"skillKey"+"}", url.PathEscape(parameterValueToString(r.skillKey, "skillKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentSkillPatch == nil {
		return localVarReturnValue, nil, reportError("agentSkillPatch is required and must be specified")
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
	// body params
	localVarPostBody = r.agentSkillPatch
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
		if localVarHTTPResponse.StatusCode == 409 {
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

type ApiPatchModelConfigRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	modelConfigKey string
	modelConfigPatch *ModelConfigPatch
}

// AI model config object to update
func (r ApiPatchModelConfigRequest) ModelConfigPatch(modelConfigPatch ModelConfigPatch) ApiPatchModelConfigRequest {
	r.modelConfigPatch = &modelConfigPatch
	return r
}

func (r ApiPatchModelConfigRequest) Execute() (*ModelConfig, *http.Response, error) {
	return r.ApiService.PatchModelConfigExecute(r)
}

/*
PatchModelConfig Update an AI model config

Update an AI model config.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param modelConfigKey
 @return ApiPatchModelConfigRequest
*/
func (a *AgentControlApiService) PatchModelConfig(ctx context.Context, projectKey string, modelConfigKey string) ApiPatchModelConfigRequest {
	return ApiPatchModelConfigRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		modelConfigKey: modelConfigKey,
	}
}

// Execute executes the request
//  @return ModelConfig
func (a *AgentControlApiService) PatchModelConfigExecute(r ApiPatchModelConfigRequest) (*ModelConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PatchModelConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/model-configs/{modelConfigKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"modelConfigKey"+"}", url.PathEscape(parameterValueToString(r.modelConfigKey, "modelConfigKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.modelConfigPatch == nil {
		return localVarReturnValue, nil, reportError("modelConfigPatch is required and must be specified")
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
	// body params
	localVarPostBody = r.modelConfigPatch
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

type ApiPatchPromptSnippetRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	snippetKey string
	promptSnippetPatch *PromptSnippetPatch
}

// Prompt snippet fields to update
func (r ApiPatchPromptSnippetRequest) PromptSnippetPatch(promptSnippetPatch PromptSnippetPatch) ApiPatchPromptSnippetRequest {
	r.promptSnippetPatch = &promptSnippetPatch
	return r
}

func (r ApiPatchPromptSnippetRequest) Execute() (*PromptSnippet, *http.Response, error) {
	return r.ApiService.PatchPromptSnippetExecute(r)
}

/*
PatchPromptSnippet Update a prompt snippet

Update an existing prompt snippet. Creates a new version of the snippet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param snippetKey
 @return ApiPatchPromptSnippetRequest
*/
func (a *AgentControlApiService) PatchPromptSnippet(ctx context.Context, projectKey string, snippetKey string) ApiPatchPromptSnippetRequest {
	return ApiPatchPromptSnippetRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		snippetKey: snippetKey,
	}
}

// Execute executes the request
//  @return PromptSnippet
func (a *AgentControlApiService) PatchPromptSnippetExecute(r ApiPatchPromptSnippetRequest) (*PromptSnippet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PromptSnippet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PatchPromptSnippet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/prompt-snippets/{snippetKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snippetKey"+"}", url.PathEscape(parameterValueToString(r.snippetKey, "snippetKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.promptSnippetPatch == nil {
		return localVarReturnValue, nil, reportError("promptSnippetPatch is required and must be specified")
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
	// body params
	localVarPostBody = r.promptSnippetPatch
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
		if localVarHTTPResponse.StatusCode == 409 {
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
	ApiService AgentControlApi
	projectKey string
	aIConfigPost *AIConfigPost
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
func (a *AgentControlApiService) PostAIConfig(ctx context.Context, projectKey string) ApiPostAIConfigRequest {
	return ApiPostAIConfigRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return AIConfig
func (a *AgentControlApiService) PostAIConfigExecute(r ApiPostAIConfigRequest) (*AIConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AIConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PostAIConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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
		if localVarHTTPResponse.StatusCode == 429 {
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
	ApiService AgentControlApi
	projectKey string
	configKey string
	aIConfigVariationPost *AIConfigVariationPost
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
func (a *AgentControlApiService) PostAIConfigVariation(ctx context.Context, projectKey string, configKey string) ApiPostAIConfigVariationRequest {
	return ApiPostAIConfigVariationRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		configKey: configKey,
	}
}

// Execute executes the request
//  @return AIConfigVariation
func (a *AgentControlApiService) PostAIConfigVariationExecute(r ApiPostAIConfigVariationRequest) (*AIConfigVariation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AIConfigVariation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PostAIConfigVariation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/{configKey}/variations"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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
		if localVarHTTPResponse.StatusCode == 429 {
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
	ApiService AgentControlApi
	projectKey string
	aIToolPost *AIToolPost
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
func (a *AgentControlApiService) PostAITool(ctx context.Context, projectKey string) ApiPostAIToolRequest {
	return ApiPostAIToolRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return AITool
func (a *AgentControlApiService) PostAIToolExecute(r ApiPostAIToolRequest) (*AITool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AITool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PostAITool")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-tools"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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

type ApiPostAgentGraphRequest struct {
	ctx context.Context
	ApiService AgentControlApi
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
func (a *AgentControlApiService) PostAgentGraph(ctx context.Context, projectKey string) ApiPostAgentGraphRequest {
	return ApiPostAgentGraphRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return AgentGraph
func (a *AgentControlApiService) PostAgentGraphExecute(r ApiPostAgentGraphRequest) (*AgentGraph, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentGraph
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PostAgentGraph")
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

type ApiPostAgentOptimizationRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	agentOptimizationPost *AgentOptimizationPost
}

// Agent optimization object to create
func (r ApiPostAgentOptimizationRequest) AgentOptimizationPost(agentOptimizationPost AgentOptimizationPost) ApiPostAgentOptimizationRequest {
	r.agentOptimizationPost = &agentOptimizationPost
	return r
}

func (r ApiPostAgentOptimizationRequest) Execute() (*AgentOptimization, *http.Response, error) {
	return r.ApiService.PostAgentOptimizationExecute(r)
}

/*
PostAgentOptimization Create agent optimization

Create a new agent optimization within the given project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiPostAgentOptimizationRequest
*/
func (a *AgentControlApiService) PostAgentOptimization(ctx context.Context, projectKey string) ApiPostAgentOptimizationRequest {
	return ApiPostAgentOptimizationRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return AgentOptimization
func (a *AgentControlApiService) PostAgentOptimizationExecute(r ApiPostAgentOptimizationRequest) (*AgentOptimization, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentOptimization
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PostAgentOptimization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/agent-optimizations"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentOptimizationPost == nil {
		return localVarReturnValue, nil, reportError("agentOptimizationPost is required and must be specified")
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
	// body params
	localVarPostBody = r.agentOptimizationPost
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
		if localVarHTTPResponse.StatusCode == 409 {
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

type ApiPostAgentOptimizationResultRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	optimizationKey string
	agentOptimizationResultPost *AgentOptimizationResultPost
}

// Agent optimization result to create
func (r ApiPostAgentOptimizationResultRequest) AgentOptimizationResultPost(agentOptimizationResultPost AgentOptimizationResultPost) ApiPostAgentOptimizationResultRequest {
	r.agentOptimizationResultPost = &agentOptimizationResultPost
	return r
}

func (r ApiPostAgentOptimizationResultRequest) Execute() (*AgentOptimizationResult, *http.Response, error) {
	return r.ApiService.PostAgentOptimizationResultExecute(r)
}

/*
PostAgentOptimizationResult Create agent optimization result

Create a new result for an agent optimization run.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @param optimizationKey
 @return ApiPostAgentOptimizationResultRequest
*/
func (a *AgentControlApiService) PostAgentOptimizationResult(ctx context.Context, projectKey string, optimizationKey string) ApiPostAgentOptimizationResultRequest {
	return ApiPostAgentOptimizationResultRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		optimizationKey: optimizationKey,
	}
}

// Execute executes the request
//  @return AgentOptimizationResult
func (a *AgentControlApiService) PostAgentOptimizationResultExecute(r ApiPostAgentOptimizationResultRequest) (*AgentOptimizationResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentOptimizationResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PostAgentOptimizationResult")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey}/results"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"optimizationKey"+"}", url.PathEscape(parameterValueToString(r.optimizationKey, "optimizationKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentOptimizationResultPost == nil {
		return localVarReturnValue, nil, reportError("agentOptimizationResultPost is required and must be specified")
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
	// body params
	localVarPostBody = r.agentOptimizationResultPost
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
		if localVarHTTPResponse.StatusCode == 409 {
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
		if localVarHTTPResponse.StatusCode == 429 {
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

type ApiPostAgentSkillRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	agentSkillPost *AgentSkillPost
}

// Agent skill object to create
func (r ApiPostAgentSkillRequest) AgentSkillPost(agentSkillPost AgentSkillPost) ApiPostAgentSkillRequest {
	r.agentSkillPost = &agentSkillPost
	return r
}

func (r ApiPostAgentSkillRequest) Execute() (*AgentSkill, *http.Response, error) {
	return r.ApiService.PostAgentSkillExecute(r)
}

/*
PostAgentSkill Create an agent skill

Create a new agent skill within the given project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiPostAgentSkillRequest
*/
func (a *AgentControlApiService) PostAgentSkill(ctx context.Context, projectKey string) ApiPostAgentSkillRequest {
	return ApiPostAgentSkillRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return AgentSkill
func (a *AgentControlApiService) PostAgentSkillExecute(r ApiPostAgentSkillRequest) (*AgentSkill, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AgentSkill
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PostAgentSkill")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/skills"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentSkillPost == nil {
		return localVarReturnValue, nil, reportError("agentSkillPost is required and must be specified")
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
	// body params
	localVarPostBody = r.agentSkillPost
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
		if localVarHTTPResponse.StatusCode == 409 {
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
	ApiService AgentControlApi
	projectKey string
	modelConfigPost *ModelConfigPost
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

Create an AI model config. You can use this in any variation for any config in your project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiPostModelConfigRequest
*/
func (a *AgentControlApiService) PostModelConfig(ctx context.Context, projectKey string) ApiPostModelConfigRequest {
	return ApiPostModelConfigRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return ModelConfig
func (a *AgentControlApiService) PostModelConfigExecute(r ApiPostModelConfigRequest) (*ModelConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PostModelConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/model-configs"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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

type ApiPostPromptSnippetRequest struct {
	ctx context.Context
	ApiService AgentControlApi
	projectKey string
	promptSnippetPost *PromptSnippetPost
}

// Prompt snippet object to create
func (r ApiPostPromptSnippetRequest) PromptSnippetPost(promptSnippetPost PromptSnippetPost) ApiPostPromptSnippetRequest {
	r.promptSnippetPost = &promptSnippetPost
	return r
}

func (r ApiPostPromptSnippetRequest) Execute() (*PromptSnippet, *http.Response, error) {
	return r.ApiService.PostPromptSnippetExecute(r)
}

/*
PostPromptSnippet Create a prompt snippet

Create a new prompt snippet within the given project.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiPostPromptSnippetRequest
*/
func (a *AgentControlApiService) PostPromptSnippet(ctx context.Context, projectKey string) ApiPostPromptSnippetRequest {
	return ApiPostPromptSnippetRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return PromptSnippet
func (a *AgentControlApiService) PostPromptSnippetExecute(r ApiPostPromptSnippetRequest) (*PromptSnippet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PromptSnippet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PostPromptSnippet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/prompt-snippets"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.promptSnippetPost == nil {
		return localVarReturnValue, nil, reportError("promptSnippetPost is required and must be specified")
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
	// body params
	localVarPostBody = r.promptSnippetPost
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
		if localVarHTTPResponse.StatusCode == 409 {
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
	ApiService AgentControlApi
	projectKey string
	restrictedModelsRequest *RestrictedModelsRequest
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

Add AI models, by key, to the restricted list. Keys are included in the response from the [List AI model configs](https://launchdarkly.com/docs/api/agent-control/list-model-configs) endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey
 @return ApiPostRestrictedModelsRequest
*/
func (a *AgentControlApiService) PostRestrictedModels(ctx context.Context, projectKey string) ApiPostRestrictedModelsRequest {
	return ApiPostRestrictedModelsRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return RestrictedModelsResponse
func (a *AgentControlApiService) PostRestrictedModelsExecute(r ApiPostRestrictedModelsRequest) (*RestrictedModelsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RestrictedModelsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentControlApiService.PostRestrictedModels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/ai-configs/model-configs/restricted"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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
