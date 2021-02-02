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

import (
	"time"
)

type PluginDefinition struct {

	Id string `json:"id,omitempty"`

	Name string `json:"name"`

	ImplementingClass string `json:"implementingClass"`

	Version string `json:"version"`

	CreatedBy string `json:"createdBy"`

	CreatedOn time.Time `json:"createdOn"`

	Enabled bool `json:"enabled"`

	ConfigurationSchema *interface{} `json:"configurationSchema"`

	// count of allowed configurations
	Limit int64 `json:"limit,omitempty"`
}
