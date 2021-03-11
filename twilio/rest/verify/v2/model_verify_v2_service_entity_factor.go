/*
 * Twilio - Verify
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

// VerifyV2ServiceEntityFactor struct for VerifyV2ServiceEntityFactor
type VerifyV2ServiceEntityFactor struct {
	AccountSid   *string                 `json:"AccountSid,omitempty"`
	Config       *map[string]interface{} `json:"Config,omitempty"`
	DateCreated  *time.Time              `json:"DateCreated,omitempty"`
	DateUpdated  *time.Time              `json:"DateUpdated,omitempty"`
	EntitySid    *string                 `json:"EntitySid,omitempty"`
	FactorType   *FactorFactorTypes      `json:"FactorType,omitempty"`
	FriendlyName *string                 `json:"FriendlyName,omitempty"`
	Identity     *string                 `json:"Identity,omitempty"`
	ServiceSid   *string                 `json:"ServiceSid,omitempty"`
	Sid          *string                 `json:"Sid,omitempty"`
	Status       *FactorFactorStatuses   `json:"Status,omitempty"`
	Url          *string                 `json:"Url,omitempty"`
}
