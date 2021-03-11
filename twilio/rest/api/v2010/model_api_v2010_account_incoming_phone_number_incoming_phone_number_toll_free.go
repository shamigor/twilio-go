/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree struct for ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree
type ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree struct {
	AccountSid           *string                                        `json:"AccountSid,omitempty"`
	AddressRequirements  *IncomingPhoneNumberTollFreeAddressRequirement `json:"AddressRequirements,omitempty"`
	AddressSid           *string                                        `json:"AddressSid,omitempty"`
	ApiVersion           *string                                        `json:"ApiVersion,omitempty"`
	Beta                 *bool                                          `json:"Beta,omitempty"`
	BundleSid            *string                                        `json:"BundleSid,omitempty"`
	Capabilities         *PhoneNumberCapabilities                       `json:"Capabilities,omitempty"`
	DateCreated          *string                                        `json:"DateCreated,omitempty"`
	DateUpdated          *string                                        `json:"DateUpdated,omitempty"`
	EmergencyAddressSid  *string                                        `json:"EmergencyAddressSid,omitempty"`
	EmergencyStatus      *IncomingPhoneNumberTollFreeEmergencyStatus    `json:"EmergencyStatus,omitempty"`
	FriendlyName         *string                                        `json:"FriendlyName,omitempty"`
	IdentitySid          *string                                        `json:"IdentitySid,omitempty"`
	Origin               *string                                        `json:"Origin,omitempty"`
	PhoneNumber          *string                                        `json:"PhoneNumber,omitempty"`
	Sid                  *string                                        `json:"Sid,omitempty"`
	SmsApplicationSid    *string                                        `json:"SmsApplicationSid,omitempty"`
	SmsFallbackMethod    *HttpMethod                                    `json:"SmsFallbackMethod,omitempty"`
	SmsFallbackUrl       *string                                        `json:"SmsFallbackUrl,omitempty"`
	SmsMethod            *HttpMethod                                    `json:"SmsMethod,omitempty"`
	SmsUrl               *string                                        `json:"SmsUrl,omitempty"`
	Status               *string                                        `json:"Status,omitempty"`
	StatusCallback       *string                                        `json:"StatusCallback,omitempty"`
	StatusCallbackMethod *HttpMethod                                    `json:"StatusCallbackMethod,omitempty"`
	TrunkSid             *string                                        `json:"TrunkSid,omitempty"`
	Uri                  *string                                        `json:"Uri,omitempty"`
	VoiceApplicationSid  *string                                        `json:"VoiceApplicationSid,omitempty"`
	VoiceCallerIdLookup  *bool                                          `json:"VoiceCallerIdLookup,omitempty"`
	VoiceFallbackMethod  *HttpMethod                                    `json:"VoiceFallbackMethod,omitempty"`
	VoiceFallbackUrl     *string                                        `json:"VoiceFallbackUrl,omitempty"`
	VoiceMethod          *HttpMethod                                    `json:"VoiceMethod,omitempty"`
	VoiceReceiveMode     *IncomingPhoneNumberTollFreeVoiceReceiveMode   `json:"VoiceReceiveMode,omitempty"`
	VoiceUrl             *string                                        `json:"VoiceUrl,omitempty"`
}
