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

type CollaborationMembership struct {

	Uuid string `json:"uuid,omitempty"`

	Name string `json:"name"`

	Title string `json:"title,omitempty"`

	StartTime time.Time `json:"startTime,omitempty"`

	Roles []string `json:"roles"`

	Donotdisturb bool `json:"donotdisturb,omitempty"`

	Blocked bool `json:"blocked,omitempty"`
}
