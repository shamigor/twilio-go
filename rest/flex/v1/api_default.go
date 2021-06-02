/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	twilio "github.com/twilio/twilio-go/client"
)

type DefaultApiService struct {
	baseURL        string
	requestHandler *twilio.RequestHandler
}

func NewDefaultApiService(requestHandler *twilio.RequestHandler) *DefaultApiService {
	return &DefaultApiService{
		requestHandler: requestHandler,
		baseURL:        "https://flex-api.twilio.com",
	}
}

func NewDefaultApiServiceWithClient(client twilio.BaseClient) *DefaultApiService {
	return NewDefaultApiService(twilio.NewRequestHandler(client))
}

// Optional parameters for the method 'CreateChannel'
type CreateChannelParams struct {
	// The chat channel's friendly name.
	ChatFriendlyName *string `json:"ChatFriendlyName,omitempty"`
	// The chat channel's unique name.
	ChatUniqueName *string `json:"ChatUniqueName,omitempty"`
	// The chat participant's friendly name.
	ChatUserFriendlyName *string `json:"ChatUserFriendlyName,omitempty"`
	// The SID of the Flex Flow.
	FlexFlowSid *string `json:"FlexFlowSid,omitempty"`
	// The `identity` value that uniquely identifies the new resource's chat User.
	Identity *string `json:"Identity,omitempty"`
	// Whether to create the channel as long-lived.
	LongLived *bool `json:"LongLived,omitempty"`
	// The pre-engagement data.
	PreEngagementData *string `json:"PreEngagementData,omitempty"`
	// The Target Contact Identity, for example the phone number of an SMS.
	Target *string `json:"Target,omitempty"`
	// The Task attributes to be added for the TaskRouter Task.
	TaskAttributes *string `json:"TaskAttributes,omitempty"`
	// The SID of the TaskRouter Task. Only valid when integration type is `task`. `null` for integration types `studio` & `external`
	TaskSid *string `json:"TaskSid,omitempty"`
}

func (params *CreateChannelParams) SetChatFriendlyName(ChatFriendlyName string) *CreateChannelParams {
	params.ChatFriendlyName = &ChatFriendlyName
	return params
}
func (params *CreateChannelParams) SetChatUniqueName(ChatUniqueName string) *CreateChannelParams {
	params.ChatUniqueName = &ChatUniqueName
	return params
}
func (params *CreateChannelParams) SetChatUserFriendlyName(ChatUserFriendlyName string) *CreateChannelParams {
	params.ChatUserFriendlyName = &ChatUserFriendlyName
	return params
}
func (params *CreateChannelParams) SetFlexFlowSid(FlexFlowSid string) *CreateChannelParams {
	params.FlexFlowSid = &FlexFlowSid
	return params
}
func (params *CreateChannelParams) SetIdentity(Identity string) *CreateChannelParams {
	params.Identity = &Identity
	return params
}
func (params *CreateChannelParams) SetLongLived(LongLived bool) *CreateChannelParams {
	params.LongLived = &LongLived
	return params
}
func (params *CreateChannelParams) SetPreEngagementData(PreEngagementData string) *CreateChannelParams {
	params.PreEngagementData = &PreEngagementData
	return params
}
func (params *CreateChannelParams) SetTarget(Target string) *CreateChannelParams {
	params.Target = &Target
	return params
}
func (params *CreateChannelParams) SetTaskAttributes(TaskAttributes string) *CreateChannelParams {
	params.TaskAttributes = &TaskAttributes
	return params
}
func (params *CreateChannelParams) SetTaskSid(TaskSid string) *CreateChannelParams {
	params.TaskSid = &TaskSid
	return params
}

func (c *DefaultApiService) CreateChannel(params *CreateChannelParams) (*FlexV1Channel, error) {
	path := "/v1/Channels"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChatFriendlyName != nil {
		data.Set("ChatFriendlyName", *params.ChatFriendlyName)
	}
	if params != nil && params.ChatUniqueName != nil {
		data.Set("ChatUniqueName", *params.ChatUniqueName)
	}
	if params != nil && params.ChatUserFriendlyName != nil {
		data.Set("ChatUserFriendlyName", *params.ChatUserFriendlyName)
	}
	if params != nil && params.FlexFlowSid != nil {
		data.Set("FlexFlowSid", *params.FlexFlowSid)
	}
	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.LongLived != nil {
		data.Set("LongLived", fmt.Sprint(*params.LongLived))
	}
	if params != nil && params.PreEngagementData != nil {
		data.Set("PreEngagementData", *params.PreEngagementData)
	}
	if params != nil && params.Target != nil {
		data.Set("Target", *params.Target)
	}
	if params != nil && params.TaskAttributes != nil {
		data.Set("TaskAttributes", *params.TaskAttributes)
	}
	if params != nil && params.TaskSid != nil {
		data.Set("TaskSid", *params.TaskSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1Channel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'CreateFlexFlow'
type CreateFlexFlowParams struct {
	// The channel type. Can be: `web`, `facebook`, `sms`, `whatsapp`, `line` or `custom`.
	ChannelType *string `json:"ChannelType,omitempty"`
	// The SID of the chat service.
	ChatServiceSid *string `json:"ChatServiceSid,omitempty"`
	// The channel contact's Identity.
	ContactIdentity *string `json:"ContactIdentity,omitempty"`
	// Whether the new Flex Flow is enabled.
	Enabled *bool `json:"Enabled,omitempty"`
	// A descriptive string that you create to describe the Flex Flow resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The Task Channel SID (TCXXXX) or unique name (e.g., `sms`) to use for the Task that will be created. Applicable and required when `integrationType` is `task`. The default value is `default`.
	IntegrationChannel *string `json:"Integration.Channel,omitempty"`
	// In the context of outbound messaging, defines whether to create a Task immediately (and therefore reserve the conversation to current agent), or delay Task creation until the customer sends the first response. Set to false to create immediately, true to delay Task creation. This setting is only applicable for outbound messaging.
	IntegrationCreationOnMessage *bool `json:"Integration.CreationOnMessage,omitempty"`
	// The SID of the Studio Flow. Required when `integrationType` is `studio`.
	IntegrationFlowSid *string `json:"Integration.FlowSid,omitempty"`
	// The Task priority of a new Task. The default priority is 0. Optional when `integrationType` is `task`, not applicable otherwise.
	IntegrationPriority *int32 `json:"Integration.Priority,omitempty"`
	// The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3 (inclusive), default is 3. Optional when `integrationType` is `external`, not applicable otherwise.
	IntegrationRetryCount *int32 `json:"Integration.RetryCount,omitempty"`
	// The Task timeout in seconds for a new Task. Default is 86,400 seconds (24 hours). Optional when `integrationType` is `task`, not applicable otherwise.
	IntegrationTimeout *int32 `json:"Integration.Timeout,omitempty"`
	// The URL of the external webhook. Required when `integrationType` is `external`.
	IntegrationUrl *string `json:"Integration.Url,omitempty"`
	// The Workflow SID for a new Task. Required when `integrationType` is `task`.
	IntegrationWorkflowSid *string `json:"Integration.WorkflowSid,omitempty"`
	// The Workspace SID for a new Task. Required when `integrationType` is `task`.
	IntegrationWorkspaceSid *string `json:"Integration.WorkspaceSid,omitempty"`
	// The software that will handle inbound messages. [Integration Type](https://www.twilio.com/docs/flex/developer/messaging/manage-flows#integration-types) can be: `studio`, `external`, or `task`.
	IntegrationType *string `json:"IntegrationType,omitempty"`
	// When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to `false`.
	JanitorEnabled *bool `json:"JanitorEnabled,omitempty"`
	// When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to `false`.
	LongLived *bool `json:"LongLived,omitempty"`
}

func (params *CreateFlexFlowParams) SetChannelType(ChannelType string) *CreateFlexFlowParams {
	params.ChannelType = &ChannelType
	return params
}
func (params *CreateFlexFlowParams) SetChatServiceSid(ChatServiceSid string) *CreateFlexFlowParams {
	params.ChatServiceSid = &ChatServiceSid
	return params
}
func (params *CreateFlexFlowParams) SetContactIdentity(ContactIdentity string) *CreateFlexFlowParams {
	params.ContactIdentity = &ContactIdentity
	return params
}
func (params *CreateFlexFlowParams) SetEnabled(Enabled bool) *CreateFlexFlowParams {
	params.Enabled = &Enabled
	return params
}
func (params *CreateFlexFlowParams) SetFriendlyName(FriendlyName string) *CreateFlexFlowParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationChannel(IntegrationChannel string) *CreateFlexFlowParams {
	params.IntegrationChannel = &IntegrationChannel
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationCreationOnMessage(IntegrationCreationOnMessage bool) *CreateFlexFlowParams {
	params.IntegrationCreationOnMessage = &IntegrationCreationOnMessage
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationFlowSid(IntegrationFlowSid string) *CreateFlexFlowParams {
	params.IntegrationFlowSid = &IntegrationFlowSid
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationPriority(IntegrationPriority int32) *CreateFlexFlowParams {
	params.IntegrationPriority = &IntegrationPriority
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationRetryCount(IntegrationRetryCount int32) *CreateFlexFlowParams {
	params.IntegrationRetryCount = &IntegrationRetryCount
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationTimeout(IntegrationTimeout int32) *CreateFlexFlowParams {
	params.IntegrationTimeout = &IntegrationTimeout
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationUrl(IntegrationUrl string) *CreateFlexFlowParams {
	params.IntegrationUrl = &IntegrationUrl
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationWorkflowSid(IntegrationWorkflowSid string) *CreateFlexFlowParams {
	params.IntegrationWorkflowSid = &IntegrationWorkflowSid
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationWorkspaceSid(IntegrationWorkspaceSid string) *CreateFlexFlowParams {
	params.IntegrationWorkspaceSid = &IntegrationWorkspaceSid
	return params
}
func (params *CreateFlexFlowParams) SetIntegrationType(IntegrationType string) *CreateFlexFlowParams {
	params.IntegrationType = &IntegrationType
	return params
}
func (params *CreateFlexFlowParams) SetJanitorEnabled(JanitorEnabled bool) *CreateFlexFlowParams {
	params.JanitorEnabled = &JanitorEnabled
	return params
}
func (params *CreateFlexFlowParams) SetLongLived(LongLived bool) *CreateFlexFlowParams {
	params.LongLived = &LongLived
	return params
}

func (c *DefaultApiService) CreateFlexFlow(params *CreateFlexFlowParams) (*FlexV1FlexFlow, error) {
	path := "/v1/FlexFlows"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChannelType != nil {
		data.Set("ChannelType", *params.ChannelType)
	}
	if params != nil && params.ChatServiceSid != nil {
		data.Set("ChatServiceSid", *params.ChatServiceSid)
	}
	if params != nil && params.ContactIdentity != nil {
		data.Set("ContactIdentity", *params.ContactIdentity)
	}
	if params != nil && params.Enabled != nil {
		data.Set("Enabled", fmt.Sprint(*params.Enabled))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.IntegrationChannel != nil {
		data.Set("Integration.Channel", *params.IntegrationChannel)
	}
	if params != nil && params.IntegrationCreationOnMessage != nil {
		data.Set("Integration.CreationOnMessage", fmt.Sprint(*params.IntegrationCreationOnMessage))
	}
	if params != nil && params.IntegrationFlowSid != nil {
		data.Set("Integration.FlowSid", *params.IntegrationFlowSid)
	}
	if params != nil && params.IntegrationPriority != nil {
		data.Set("Integration.Priority", fmt.Sprint(*params.IntegrationPriority))
	}
	if params != nil && params.IntegrationRetryCount != nil {
		data.Set("Integration.RetryCount", fmt.Sprint(*params.IntegrationRetryCount))
	}
	if params != nil && params.IntegrationTimeout != nil {
		data.Set("Integration.Timeout", fmt.Sprint(*params.IntegrationTimeout))
	}
	if params != nil && params.IntegrationUrl != nil {
		data.Set("Integration.Url", *params.IntegrationUrl)
	}
	if params != nil && params.IntegrationWorkflowSid != nil {
		data.Set("Integration.WorkflowSid", *params.IntegrationWorkflowSid)
	}
	if params != nil && params.IntegrationWorkspaceSid != nil {
		data.Set("Integration.WorkspaceSid", *params.IntegrationWorkspaceSid)
	}
	if params != nil && params.IntegrationType != nil {
		data.Set("IntegrationType", *params.IntegrationType)
	}
	if params != nil && params.JanitorEnabled != nil {
		data.Set("JanitorEnabled", fmt.Sprint(*params.JanitorEnabled))
	}
	if params != nil && params.LongLived != nil {
		data.Set("LongLived", fmt.Sprint(*params.LongLived))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1FlexFlow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'CreateWebChannel'
type CreateWebChannelParams struct {
	// The chat channel's friendly name.
	ChatFriendlyName *string `json:"ChatFriendlyName,omitempty"`
	// The chat channel's unique name.
	ChatUniqueName *string `json:"ChatUniqueName,omitempty"`
	// The chat participant's friendly name.
	CustomerFriendlyName *string `json:"CustomerFriendlyName,omitempty"`
	// The SID of the Flex Flow.
	FlexFlowSid *string `json:"FlexFlowSid,omitempty"`
	// The chat identity.
	Identity *string `json:"Identity,omitempty"`
	// The pre-engagement data.
	PreEngagementData *string `json:"PreEngagementData,omitempty"`
}

func (params *CreateWebChannelParams) SetChatFriendlyName(ChatFriendlyName string) *CreateWebChannelParams {
	params.ChatFriendlyName = &ChatFriendlyName
	return params
}
func (params *CreateWebChannelParams) SetChatUniqueName(ChatUniqueName string) *CreateWebChannelParams {
	params.ChatUniqueName = &ChatUniqueName
	return params
}
func (params *CreateWebChannelParams) SetCustomerFriendlyName(CustomerFriendlyName string) *CreateWebChannelParams {
	params.CustomerFriendlyName = &CustomerFriendlyName
	return params
}
func (params *CreateWebChannelParams) SetFlexFlowSid(FlexFlowSid string) *CreateWebChannelParams {
	params.FlexFlowSid = &FlexFlowSid
	return params
}
func (params *CreateWebChannelParams) SetIdentity(Identity string) *CreateWebChannelParams {
	params.Identity = &Identity
	return params
}
func (params *CreateWebChannelParams) SetPreEngagementData(PreEngagementData string) *CreateWebChannelParams {
	params.PreEngagementData = &PreEngagementData
	return params
}

func (c *DefaultApiService) CreateWebChannel(params *CreateWebChannelParams) (*FlexV1WebChannel, error) {
	path := "/v1/WebChannels"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChatFriendlyName != nil {
		data.Set("ChatFriendlyName", *params.ChatFriendlyName)
	}
	if params != nil && params.ChatUniqueName != nil {
		data.Set("ChatUniqueName", *params.ChatUniqueName)
	}
	if params != nil && params.CustomerFriendlyName != nil {
		data.Set("CustomerFriendlyName", *params.CustomerFriendlyName)
	}
	if params != nil && params.FlexFlowSid != nil {
		data.Set("FlexFlowSid", *params.FlexFlowSid)
	}
	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.PreEngagementData != nil {
		data.Set("PreEngagementData", *params.PreEngagementData)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1WebChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *DefaultApiService) DeleteChannel(Sid string) error {
	path := "/v1/Channels/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *DefaultApiService) DeleteFlexFlow(Sid string) error {
	path := "/v1/FlexFlows/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *DefaultApiService) DeleteWebChannel(Sid string) error {
	path := "/v1/WebChannels/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *DefaultApiService) FetchChannel(Sid string) (*FlexV1Channel, error) {
	path := "/v1/Channels/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1Channel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'FetchConfiguration'
type FetchConfigurationParams struct {
	// The Pinned UI version of the Configuration resource to fetch.
	UiVersion *string `json:"UiVersion,omitempty"`
}

func (params *FetchConfigurationParams) SetUiVersion(UiVersion string) *FetchConfigurationParams {
	params.UiVersion = &UiVersion
	return params
}

func (c *DefaultApiService) FetchConfiguration(params *FetchConfigurationParams) (*FlexV1Configuration, error) {
	path := "/v1/Configuration"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.UiVersion != nil {
		data.Set("UiVersion", *params.UiVersion)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1Configuration{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *DefaultApiService) FetchFlexFlow(Sid string) (*FlexV1FlexFlow, error) {
	path := "/v1/FlexFlows/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1FlexFlow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *DefaultApiService) FetchWebChannel(Sid string) (*FlexV1WebChannel, error) {
	path := "/v1/WebChannels/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1WebChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListChannel'
type ListChannelParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListChannelParams) SetPageSize(PageSize int32) *ListChannelParams {
	params.PageSize = &PageSize
	return params
}

func (c *DefaultApiService) ListChannel(params *ListChannelParams) (*ListChannelResponse, error) {
	path := "/v1/Channels"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFlexFlow'
type ListFlexFlowParams struct {
	// The `friendly_name` of the Flex Flow resources to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListFlexFlowParams) SetFriendlyName(FriendlyName string) *ListFlexFlowParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListFlexFlowParams) SetPageSize(PageSize int32) *ListFlexFlowParams {
	params.PageSize = &PageSize
	return params
}

func (c *DefaultApiService) ListFlexFlow(params *ListFlexFlowParams) (*ListFlexFlowResponse, error) {
	path := "/v1/FlexFlows"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFlexFlowResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListWebChannel'
type ListWebChannelParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListWebChannelParams) SetPageSize(PageSize int32) *ListWebChannelParams {
	params.PageSize = &PageSize
	return params
}

func (c *DefaultApiService) ListWebChannel(params *ListWebChannelParams) (*ListWebChannelResponse, error) {
	path := "/v1/WebChannels"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListWebChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *DefaultApiService) UpdateConfiguration() (*FlexV1Configuration, error) {
	path := "/v1/Configuration"

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1Configuration{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateFlexFlow'
type UpdateFlexFlowParams struct {
	// The channel type. Can be: `web`, `facebook`, `sms`, `whatsapp`, `line` or `custom`.
	ChannelType *string `json:"ChannelType,omitempty"`
	// The SID of the chat service.
	ChatServiceSid *string `json:"ChatServiceSid,omitempty"`
	// The channel contact's Identity.
	ContactIdentity *string `json:"ContactIdentity,omitempty"`
	// Whether the new Flex Flow is enabled.
	Enabled *bool `json:"Enabled,omitempty"`
	// A descriptive string that you create to describe the Flex Flow resource.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The Task Channel SID (TCXXXX) or unique name (e.g., `sms`) to use for the Task that will be created. Applicable and required when `integrationType` is `task`. The default value is `default`.
	IntegrationChannel *string `json:"Integration.Channel,omitempty"`
	// In the context of outbound messaging, defines whether to create a Task immediately (and therefore reserve the conversation to current agent), or delay Task creation until the customer sends the first response. Set to false to create immediately, true to delay Task creation. This setting is only applicable for outbound messaging.
	IntegrationCreationOnMessage *bool `json:"Integration.CreationOnMessage,omitempty"`
	// The SID of the Studio Flow. Required when `integrationType` is `studio`.
	IntegrationFlowSid *string `json:"Integration.FlowSid,omitempty"`
	// The Task priority of a new Task. The default priority is 0. Optional when `integrationType` is `task`, not applicable otherwise.
	IntegrationPriority *int32 `json:"Integration.Priority,omitempty"`
	// The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3 (inclusive), default is 3. Optional when `integrationType` is `external`, not applicable otherwise.
	IntegrationRetryCount *int32 `json:"Integration.RetryCount,omitempty"`
	// The Task timeout in seconds for a new Task. Default is 86,400 seconds (24 hours). Optional when `integrationType` is `task`, not applicable otherwise.
	IntegrationTimeout *int32 `json:"Integration.Timeout,omitempty"`
	// The URL of the external webhook. Required when `integrationType` is `external`.
	IntegrationUrl *string `json:"Integration.Url,omitempty"`
	// The Workflow SID for a new Task. Required when `integrationType` is `task`.
	IntegrationWorkflowSid *string `json:"Integration.WorkflowSid,omitempty"`
	// The Workspace SID for a new Task. Required when `integrationType` is `task`.
	IntegrationWorkspaceSid *string `json:"Integration.WorkspaceSid,omitempty"`
	// The software that will handle inbound messages. [Integration Type](https://www.twilio.com/docs/flex/developer/messaging/manage-flows#integration-types) can be: `studio`, `external`, or `task`.
	IntegrationType *string `json:"IntegrationType,omitempty"`
	// When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to `false`.
	JanitorEnabled *bool `json:"JanitorEnabled,omitempty"`
	// When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to `false`.
	LongLived *bool `json:"LongLived,omitempty"`
}

func (params *UpdateFlexFlowParams) SetChannelType(ChannelType string) *UpdateFlexFlowParams {
	params.ChannelType = &ChannelType
	return params
}
func (params *UpdateFlexFlowParams) SetChatServiceSid(ChatServiceSid string) *UpdateFlexFlowParams {
	params.ChatServiceSid = &ChatServiceSid
	return params
}
func (params *UpdateFlexFlowParams) SetContactIdentity(ContactIdentity string) *UpdateFlexFlowParams {
	params.ContactIdentity = &ContactIdentity
	return params
}
func (params *UpdateFlexFlowParams) SetEnabled(Enabled bool) *UpdateFlexFlowParams {
	params.Enabled = &Enabled
	return params
}
func (params *UpdateFlexFlowParams) SetFriendlyName(FriendlyName string) *UpdateFlexFlowParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationChannel(IntegrationChannel string) *UpdateFlexFlowParams {
	params.IntegrationChannel = &IntegrationChannel
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationCreationOnMessage(IntegrationCreationOnMessage bool) *UpdateFlexFlowParams {
	params.IntegrationCreationOnMessage = &IntegrationCreationOnMessage
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationFlowSid(IntegrationFlowSid string) *UpdateFlexFlowParams {
	params.IntegrationFlowSid = &IntegrationFlowSid
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationPriority(IntegrationPriority int32) *UpdateFlexFlowParams {
	params.IntegrationPriority = &IntegrationPriority
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationRetryCount(IntegrationRetryCount int32) *UpdateFlexFlowParams {
	params.IntegrationRetryCount = &IntegrationRetryCount
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationTimeout(IntegrationTimeout int32) *UpdateFlexFlowParams {
	params.IntegrationTimeout = &IntegrationTimeout
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationUrl(IntegrationUrl string) *UpdateFlexFlowParams {
	params.IntegrationUrl = &IntegrationUrl
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationWorkflowSid(IntegrationWorkflowSid string) *UpdateFlexFlowParams {
	params.IntegrationWorkflowSid = &IntegrationWorkflowSid
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationWorkspaceSid(IntegrationWorkspaceSid string) *UpdateFlexFlowParams {
	params.IntegrationWorkspaceSid = &IntegrationWorkspaceSid
	return params
}
func (params *UpdateFlexFlowParams) SetIntegrationType(IntegrationType string) *UpdateFlexFlowParams {
	params.IntegrationType = &IntegrationType
	return params
}
func (params *UpdateFlexFlowParams) SetJanitorEnabled(JanitorEnabled bool) *UpdateFlexFlowParams {
	params.JanitorEnabled = &JanitorEnabled
	return params
}
func (params *UpdateFlexFlowParams) SetLongLived(LongLived bool) *UpdateFlexFlowParams {
	params.LongLived = &LongLived
	return params
}

func (c *DefaultApiService) UpdateFlexFlow(Sid string, params *UpdateFlexFlowParams) (*FlexV1FlexFlow, error) {
	path := "/v1/FlexFlows/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChannelType != nil {
		data.Set("ChannelType", *params.ChannelType)
	}
	if params != nil && params.ChatServiceSid != nil {
		data.Set("ChatServiceSid", *params.ChatServiceSid)
	}
	if params != nil && params.ContactIdentity != nil {
		data.Set("ContactIdentity", *params.ContactIdentity)
	}
	if params != nil && params.Enabled != nil {
		data.Set("Enabled", fmt.Sprint(*params.Enabled))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.IntegrationChannel != nil {
		data.Set("Integration.Channel", *params.IntegrationChannel)
	}
	if params != nil && params.IntegrationCreationOnMessage != nil {
		data.Set("Integration.CreationOnMessage", fmt.Sprint(*params.IntegrationCreationOnMessage))
	}
	if params != nil && params.IntegrationFlowSid != nil {
		data.Set("Integration.FlowSid", *params.IntegrationFlowSid)
	}
	if params != nil && params.IntegrationPriority != nil {
		data.Set("Integration.Priority", fmt.Sprint(*params.IntegrationPriority))
	}
	if params != nil && params.IntegrationRetryCount != nil {
		data.Set("Integration.RetryCount", fmt.Sprint(*params.IntegrationRetryCount))
	}
	if params != nil && params.IntegrationTimeout != nil {
		data.Set("Integration.Timeout", fmt.Sprint(*params.IntegrationTimeout))
	}
	if params != nil && params.IntegrationUrl != nil {
		data.Set("Integration.Url", *params.IntegrationUrl)
	}
	if params != nil && params.IntegrationWorkflowSid != nil {
		data.Set("Integration.WorkflowSid", *params.IntegrationWorkflowSid)
	}
	if params != nil && params.IntegrationWorkspaceSid != nil {
		data.Set("Integration.WorkspaceSid", *params.IntegrationWorkspaceSid)
	}
	if params != nil && params.IntegrationType != nil {
		data.Set("IntegrationType", *params.IntegrationType)
	}
	if params != nil && params.JanitorEnabled != nil {
		data.Set("JanitorEnabled", fmt.Sprint(*params.JanitorEnabled))
	}
	if params != nil && params.LongLived != nil {
		data.Set("LongLived", fmt.Sprint(*params.LongLived))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1FlexFlow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateWebChannel'
type UpdateWebChannelParams struct {
	// The chat status. Can only be `inactive`.
	ChatStatus *string `json:"ChatStatus,omitempty"`
	// The post-engagement data.
	PostEngagementData *string `json:"PostEngagementData,omitempty"`
}

func (params *UpdateWebChannelParams) SetChatStatus(ChatStatus string) *UpdateWebChannelParams {
	params.ChatStatus = &ChatStatus
	return params
}
func (params *UpdateWebChannelParams) SetPostEngagementData(PostEngagementData string) *UpdateWebChannelParams {
	params.PostEngagementData = &PostEngagementData
	return params
}

func (c *DefaultApiService) UpdateWebChannel(Sid string, params *UpdateWebChannelParams) (*FlexV1WebChannel, error) {
	path := "/v1/WebChannels/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChatStatus != nil {
		data.Set("ChatStatus", *params.ChatStatus)
	}
	if params != nil && params.PostEngagementData != nil {
		data.Set("PostEngagementData", *params.PostEngagementData)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1WebChannel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
