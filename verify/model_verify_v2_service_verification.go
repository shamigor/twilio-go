/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twilio
import (
	"time"
)
// VerifyV2ServiceVerification struct for VerifyV2ServiceVerification
type VerifyV2ServiceVerification struct {
	AccountSid string `json:"account_sid,omitempty"`
	Amount string `json:"amount,omitempty"`
	Channel string `json:"channel,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	DateUpdated time.Time `json:"date_updated,omitempty"`
	Lookup map[string]interface{} `json:"lookup,omitempty"`
	Payee string `json:"payee,omitempty"`
	SendCodeAttempts []map[string]interface{} `json:"send_code_attempts,omitempty"`
	ServiceSid string `json:"service_sid,omitempty"`
	Sid string `json:"sid,omitempty"`
	Status string `json:"status,omitempty"`
	To string `json:"to,omitempty"`
	Url string `json:"url,omitempty"`
	Valid bool `json:"valid,omitempty"`
}