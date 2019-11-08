/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.21
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type CustomRoles struct {
	Links *Links `json:"_links,omitempty"`
	Items []CustomRole `json:"items,omitempty"`
}
