/*
 * Twilio - Serverless
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

// ServerlessV1Service struct for ServerlessV1Service
type ServerlessV1Service struct {
	// The SID of the Account that created the Service resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the Service resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the Service resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the Service resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// Whether to inject Account credentials into a function invocation context
	IncludeCredentials *bool `json:"include_credentials,omitempty"`
	// The URLs of the Service's nested resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The unique string that identifies the Service resource
	Sid *string `json:"sid,omitempty"`
	// Whether the Service resource's properties and subresources can be edited via the UI
	UiEditable *bool `json:"ui_editable,omitempty"`
	// A user-defined string that uniquely identifies the Service resource
	UniqueName *string `json:"unique_name,omitempty"`
	// The absolute URL of the Service resource
	Url *string `json:"url,omitempty"`
}
