/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.24
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type ProjectBody struct {
	Name string `json:"name"`
	Key string `json:"key"`
	IncludeInSnippetByDefault bool `json:"includeInSnippetByDefault,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Environments []EnvironmentPost `json:"environments,omitempty"`
}
