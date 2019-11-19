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

type User struct {
	Key string `json:"key,omitempty"`
	Secondary string `json:"secondary,omitempty"`
	Ip string `json:"ip,omitempty"`
	Country string `json:"country,omitempty"`
	Email string `json:"email,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	Name string `json:"name,omitempty"`
	Anonymous bool `json:"anonymous,omitempty"`
	Custom *interface{} `json:"custom,omitempty"`
}
