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

// CredentialPushService the model 'CredentialPushService'
type CredentialPushService string

// List of credential_push_service
const (
	CREDENTIALPUSHSERVICE_GCM CredentialPushService = "gcm"
	CREDENTIALPUSHSERVICE_APN CredentialPushService = "apn"
	CREDENTIALPUSHSERVICE_FCM CredentialPushService = "fcm"
)
