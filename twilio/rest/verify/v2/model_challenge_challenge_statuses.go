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

// ChallengeChallengeStatuses the model 'ChallengeChallengeStatuses'
type ChallengeChallengeStatuses string

// List of challenge_challenge_statuses
const (
	CHALLENGECHALLENGESTATUSES_PENDING  ChallengeChallengeStatuses = "pending"
	CHALLENGECHALLENGESTATUSES_EXPIRED  ChallengeChallengeStatuses = "expired"
	CHALLENGECHALLENGESTATUSES_APPROVED ChallengeChallengeStatuses = "approved"
	CHALLENGECHALLENGESTATUSES_DENIED   ChallengeChallengeStatuses = "denied"
)
