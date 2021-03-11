/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// WirelessV1RatePlan struct for WirelessV1RatePlan
type WirelessV1RatePlan struct {
	AccountSid                    *string    `json:"AccountSid,omitempty"`
	DataEnabled                   *bool      `json:"DataEnabled,omitempty"`
	DataLimit                     *int32     `json:"DataLimit,omitempty"`
	DataMetering                  *string    `json:"DataMetering,omitempty"`
	DateCreated                   *time.Time `json:"DateCreated,omitempty"`
	DateUpdated                   *time.Time `json:"DateUpdated,omitempty"`
	FriendlyName                  *string    `json:"FriendlyName,omitempty"`
	InternationalRoaming          *[]string  `json:"InternationalRoaming,omitempty"`
	InternationalRoamingDataLimit *int32     `json:"InternationalRoamingDataLimit,omitempty"`
	MessagingEnabled              *bool      `json:"MessagingEnabled,omitempty"`
	NationalRoamingDataLimit      *int32     `json:"NationalRoamingDataLimit,omitempty"`
	NationalRoamingEnabled        *bool      `json:"NationalRoamingEnabled,omitempty"`
	Sid                           *string    `json:"Sid,omitempty"`
	UniqueName                    *string    `json:"UniqueName,omitempty"`
	Url                           *string    `json:"Url,omitempty"`
	VoiceEnabled                  *bool      `json:"VoiceEnabled,omitempty"`
}
