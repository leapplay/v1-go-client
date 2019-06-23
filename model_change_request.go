/*
 * Leap Play
 *
 * Learn more at https://www.leap-play.com
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package leapclient

import (
	"time"
)

type ChangeRequest struct {
	Id string `json:"id,omitempty"`
	CreatedOnUtc time.Time `json:"createdOnUtc,omitempty"`
	Reference string `json:"reference,omitempty"`
	TimeChange string `json:"timeChange,omitempty"`
}
