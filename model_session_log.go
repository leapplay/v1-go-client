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

type SessionLog struct {
	StationId string `json:"stationId,omitempty"`
	SessionId string `json:"sessionId,omitempty"`
	RequestedBy string `json:"requestedBy,omitempty"`
	LatestState string `json:"latestState,omitempty"`
	StartedUtc time.Time `json:"startedUtc,omitempty"`
	Reference string `json:"reference,omitempty"`
	MaxDurationLimit string `json:"maxDurationLimit,omitempty"`
	EndedUtc time.Time `json:"endedUtc,omitempty"`
	StoppedBy string `json:"stoppedBy,omitempty"`
	ChangeRequests []ChangeRequest `json:"changeRequests,omitempty"`
}
