/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// TaskrouterV1WorkspaceWorkspaceCumulativeStatistics struct for TaskrouterV1WorkspaceWorkspaceCumulativeStatistics
type TaskrouterV1WorkspaceWorkspaceCumulativeStatistics struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The average time in seconds between Task creation and acceptance
	AvgTaskAcceptanceTime *int32 `json:"avg_task_acceptance_time,omitempty"`
	// The end of the interval during which these statistics were calculated
	EndTime *time.Time `json:"end_time,omitempty"`
	// The total number of Reservations accepted by Workers
	ReservationsAccepted *int32 `json:"reservations_accepted,omitempty"`
	// The total number of Reservations that were canceled
	ReservationsCanceled *int32 `json:"reservations_canceled,omitempty"`
	// The total number of Reservations that were created for Workers
	ReservationsCreated *int32 `json:"reservations_created,omitempty"`
	// The total number of Reservations that were rejected
	ReservationsRejected *int32 `json:"reservations_rejected,omitempty"`
	// The total number of Reservations that were rescinded
	ReservationsRescinded *int32 `json:"reservations_rescinded,omitempty"`
	// The total number of Reservations that were timed out
	ReservationsTimedOut *int32 `json:"reservations_timed_out,omitempty"`
	// A list of objects that describe the Tasks canceled and reservations accepted above and below the specified thresholds
	SplitByWaitTime *map[string]interface{} `json:"split_by_wait_time,omitempty"`
	// The beginning of the interval during which these statistics were calculated
	StartTime *time.Time `json:"start_time,omitempty"`
	// The total number of Tasks that were canceled
	TasksCanceled *int32 `json:"tasks_canceled,omitempty"`
	// The total number of Tasks that were completed
	TasksCompleted *int32 `json:"tasks_completed,omitempty"`
	// The total number of Tasks created
	TasksCreated *int32 `json:"tasks_created,omitempty"`
	// The total number of Tasks that were deleted
	TasksDeleted *int32 `json:"tasks_deleted,omitempty"`
	// The total number of Tasks that were moved from one queue to another
	TasksMoved *int32 `json:"tasks_moved,omitempty"`
	// The total number of Tasks that were timed out of their Workflows
	TasksTimedOutInWorkflow *int32 `json:"tasks_timed_out_in_workflow,omitempty"`
	// The absolute URL of the Workspace statistics resource
	Url *string `json:"url,omitempty"`
	// The wait duration statistics for Tasks that were accepted
	WaitDurationUntilAccepted *map[string]interface{} `json:"wait_duration_until_accepted,omitempty"`
	// The wait duration statistics for Tasks that were canceled
	WaitDurationUntilCanceled *map[string]interface{} `json:"wait_duration_until_canceled,omitempty"`
	// The SID of the Workspace
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}
