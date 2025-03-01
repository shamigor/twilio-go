/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// SyncV1ServiceSyncStreamStreamMessage struct for SyncV1ServiceSyncStreamStreamMessage
type SyncV1ServiceSyncStreamStreamMessage struct {
	// Stream Message body
	Data *map[string]interface{} `json:"data,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
}
