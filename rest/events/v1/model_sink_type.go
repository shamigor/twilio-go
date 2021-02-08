/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SinkType the model 'SinkType'
type SinkType string

// List of sink_type
const (
	SINKTYPE_KINESIS SinkType = "kinesis"
	SINKTYPE_WEBHOOK SinkType = "webhook"
)