/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.22
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type Clause struct {
	Attribute string `json:"attribute,omitempty"`
	Op string `json:"op,omitempty"`
	Values []interface{} `json:"values,omitempty"`
	Negate bool `json:"negate,omitempty"`
}
