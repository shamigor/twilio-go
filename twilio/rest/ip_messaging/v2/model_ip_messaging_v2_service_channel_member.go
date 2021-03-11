/*
 * Twilio - Ip_messaging
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

// IpMessagingV2ServiceChannelMember struct for IpMessagingV2ServiceChannelMember
type IpMessagingV2ServiceChannelMember struct {
	AccountSid               *string    `json:"AccountSid,omitempty"`
	Attributes               *string    `json:"Attributes,omitempty"`
	ChannelSid               *string    `json:"ChannelSid,omitempty"`
	DateCreated              *time.Time `json:"DateCreated,omitempty"`
	DateUpdated              *time.Time `json:"DateUpdated,omitempty"`
	Identity                 *string    `json:"Identity,omitempty"`
	LastConsumedMessageIndex *int32     `json:"LastConsumedMessageIndex,omitempty"`
	LastConsumptionTimestamp *time.Time `json:"LastConsumptionTimestamp,omitempty"`
	RoleSid                  *string    `json:"RoleSid,omitempty"`
	ServiceSid               *string    `json:"ServiceSid,omitempty"`
	Sid                      *string    `json:"Sid,omitempty"`
	Url                      *string    `json:"Url,omitempty"`
}
