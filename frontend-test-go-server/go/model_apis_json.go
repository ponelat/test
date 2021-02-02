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

type ApisJson struct {

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Url string `json:"url,omitempty"`

	Created time.Time `json:"created,omitempty"`

	Modified time.Time `json:"modified,omitempty"`

	SpecificationVersion string `json:"specificationVersion,omitempty"`

	Offset int32 `json:"offset,omitempty"`

	TotalCount int64 `json:"totalCount,omitempty"`

	Apis []ApisJsonApi `json:"apis,omitempty"`
}
