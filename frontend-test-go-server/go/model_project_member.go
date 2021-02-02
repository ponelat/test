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

type ProjectMember struct {

	Name string `json:"name"`

	Type_ string `json:"type"`

	Roles []string `json:"roles,omitempty"`
}
