/*
 * SwaggerHub Registry API
 *
 * The registry API for SwaggerHub
 *
 * API version: 1.0.97
 * Contact: info@swaggerhub.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GitHubExportSettings struct {

	Token string `json:"token"`

	Owner string `json:"owner"`

	Repository string `json:"repository"`

	Branch string `json:"branch"`

	Path string `json:"path,omitempty"`

	YamlPath string `json:"yamlPath,omitempty"`

	NotificationEmail string `json:"notificationEmail"`

	Service string `json:"service,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	LastResult *AbstractResult `json:"lastResult,omitempty"`
}