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

type ApisJsonApi struct {

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Tags []string `json:"tags,omitempty"`

	Properties []ApisJsonProperty `json:"properties,omitempty"`
}