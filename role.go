/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.4
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type Role string

// List of Role
const (
	WRITER Role = "writer"
	READER Role = "reader"
	ADMIN Role = "admin"
	OWNER Role = "owner"
)
