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

type ModelFallthrough struct {
	Variation int32 `json:"variation,omitempty"`
	Rollout *Rollout `json:"rollout,omitempty"`
}
