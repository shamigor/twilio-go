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

// CallRecordingSource the model 'CallRecordingSource'
type CallRecordingSource string

// List of call_recording_source
const (
	CALLRECORDINGSOURCE_DIAL_VERB                      CallRecordingSource = "DialVerb"
	CALLRECORDINGSOURCE_CONFERENCE                     CallRecordingSource = "Conference"
	CALLRECORDINGSOURCE_OUTBOUND_API                   CallRecordingSource = "OutboundAPI"
	CALLRECORDINGSOURCE_TRUNKING                       CallRecordingSource = "Trunking"
	CALLRECORDINGSOURCE_RECORD_VERB                    CallRecordingSource = "RecordVerb"
	CALLRECORDINGSOURCE_START_CALL_RECORDING_API       CallRecordingSource = "StartCallRecordingAPI"
	CALLRECORDINGSOURCE_START_CONFERENCE_RECORDING_API CallRecordingSource = "StartConferenceRecordingAPI"
)
