/*
 * Twilio - Messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// MessagingV1ServiceShortCode struct for MessagingV1ServiceShortCode
type MessagingV1ServiceShortCode struct {
	AccountSid string `json:"AccountSid,omitempty"`
	Capabilities []string `json:"Capabilities,omitempty"`
	CountryCode string `json:"CountryCode,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	ServiceSid string `json:"ServiceSid,omitempty"`
	ShortCode string `json:"ShortCode,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
}