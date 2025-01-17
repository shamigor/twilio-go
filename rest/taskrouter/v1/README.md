# Go API client for openapi

This is the public Twilio REST API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project from the OpenAPI specs located at [twilio/twilio-oai](https://github.com/twilio/twilio-oai/tree/main/spec).  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.16.1
- Package version: 1.0.0
- Build package: com.twilio.oai.TwilioGoGenerator
For more information, please visit [https://support.twilio.com](https://support.twilio.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://taskrouter.twilio.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**CreateActivity**](docs/DefaultApi.md#createactivity) | **Post** /v1/Workspaces/{WorkspaceSid}/Activities | 
*DefaultApi* | [**CreateTask**](docs/DefaultApi.md#createtask) | **Post** /v1/Workspaces/{WorkspaceSid}/Tasks | 
*DefaultApi* | [**CreateTaskChannel**](docs/DefaultApi.md#createtaskchannel) | **Post** /v1/Workspaces/{WorkspaceSid}/TaskChannels | 
*DefaultApi* | [**CreateTaskQueue**](docs/DefaultApi.md#createtaskqueue) | **Post** /v1/Workspaces/{WorkspaceSid}/TaskQueues | 
*DefaultApi* | [**CreateWorker**](docs/DefaultApi.md#createworker) | **Post** /v1/Workspaces/{WorkspaceSid}/Workers | 
*DefaultApi* | [**CreateWorkflow**](docs/DefaultApi.md#createworkflow) | **Post** /v1/Workspaces/{WorkspaceSid}/Workflows | 
*DefaultApi* | [**CreateWorkspace**](docs/DefaultApi.md#createworkspace) | **Post** /v1/Workspaces | 
*DefaultApi* | [**DeleteActivity**](docs/DefaultApi.md#deleteactivity) | **Delete** /v1/Workspaces/{WorkspaceSid}/Activities/{Sid} | 
*DefaultApi* | [**DeleteTask**](docs/DefaultApi.md#deletetask) | **Delete** /v1/Workspaces/{WorkspaceSid}/Tasks/{Sid} | 
*DefaultApi* | [**DeleteTaskChannel**](docs/DefaultApi.md#deletetaskchannel) | **Delete** /v1/Workspaces/{WorkspaceSid}/TaskChannels/{Sid} | 
*DefaultApi* | [**DeleteTaskQueue**](docs/DefaultApi.md#deletetaskqueue) | **Delete** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{Sid} | 
*DefaultApi* | [**DeleteWorker**](docs/DefaultApi.md#deleteworker) | **Delete** /v1/Workspaces/{WorkspaceSid}/Workers/{Sid} | 
*DefaultApi* | [**DeleteWorkflow**](docs/DefaultApi.md#deleteworkflow) | **Delete** /v1/Workspaces/{WorkspaceSid}/Workflows/{Sid} | 
*DefaultApi* | [**DeleteWorkspace**](docs/DefaultApi.md#deleteworkspace) | **Delete** /v1/Workspaces/{Sid} | 
*DefaultApi* | [**FetchActivity**](docs/DefaultApi.md#fetchactivity) | **Get** /v1/Workspaces/{WorkspaceSid}/Activities/{Sid} | 
*DefaultApi* | [**FetchEvent**](docs/DefaultApi.md#fetchevent) | **Get** /v1/Workspaces/{WorkspaceSid}/Events/{Sid} | 
*DefaultApi* | [**FetchTask**](docs/DefaultApi.md#fetchtask) | **Get** /v1/Workspaces/{WorkspaceSid}/Tasks/{Sid} | 
*DefaultApi* | [**FetchTaskChannel**](docs/DefaultApi.md#fetchtaskchannel) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskChannels/{Sid} | 
*DefaultApi* | [**FetchTaskQueue**](docs/DefaultApi.md#fetchtaskqueue) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{Sid} | 
*DefaultApi* | [**FetchTaskQueueCumulativeStatistics**](docs/DefaultApi.md#fetchtaskqueuecumulativestatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{TaskQueueSid}/CumulativeStatistics | 
*DefaultApi* | [**FetchTaskQueueRealTimeStatistics**](docs/DefaultApi.md#fetchtaskqueuerealtimestatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{TaskQueueSid}/RealTimeStatistics | 
*DefaultApi* | [**FetchTaskQueueStatistics**](docs/DefaultApi.md#fetchtaskqueuestatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{TaskQueueSid}/Statistics | 
*DefaultApi* | [**FetchTaskReservation**](docs/DefaultApi.md#fetchtaskreservation) | **Get** /v1/Workspaces/{WorkspaceSid}/Tasks/{TaskSid}/Reservations/{Sid} | 
*DefaultApi* | [**FetchWorker**](docs/DefaultApi.md#fetchworker) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/{Sid} | 
*DefaultApi* | [**FetchWorkerChannel**](docs/DefaultApi.md#fetchworkerchannel) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Channels/{Sid} | 
*DefaultApi* | [**FetchWorkerInstanceStatistics**](docs/DefaultApi.md#fetchworkerinstancestatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Statistics | 
*DefaultApi* | [**FetchWorkerReservation**](docs/DefaultApi.md#fetchworkerreservation) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Reservations/{Sid} | 
*DefaultApi* | [**FetchWorkerStatistics**](docs/DefaultApi.md#fetchworkerstatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/Statistics | 
*DefaultApi* | [**FetchWorkersCumulativeStatistics**](docs/DefaultApi.md#fetchworkerscumulativestatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/CumulativeStatistics | 
*DefaultApi* | [**FetchWorkersRealTimeStatistics**](docs/DefaultApi.md#fetchworkersrealtimestatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/RealTimeStatistics | 
*DefaultApi* | [**FetchWorkflow**](docs/DefaultApi.md#fetchworkflow) | **Get** /v1/Workspaces/{WorkspaceSid}/Workflows/{Sid} | 
*DefaultApi* | [**FetchWorkflowCumulativeStatistics**](docs/DefaultApi.md#fetchworkflowcumulativestatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Workflows/{WorkflowSid}/CumulativeStatistics | 
*DefaultApi* | [**FetchWorkflowRealTimeStatistics**](docs/DefaultApi.md#fetchworkflowrealtimestatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Workflows/{WorkflowSid}/RealTimeStatistics | 
*DefaultApi* | [**FetchWorkflowStatistics**](docs/DefaultApi.md#fetchworkflowstatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Workflows/{WorkflowSid}/Statistics | 
*DefaultApi* | [**FetchWorkspace**](docs/DefaultApi.md#fetchworkspace) | **Get** /v1/Workspaces/{Sid} | 
*DefaultApi* | [**FetchWorkspaceCumulativeStatistics**](docs/DefaultApi.md#fetchworkspacecumulativestatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/CumulativeStatistics | 
*DefaultApi* | [**FetchWorkspaceRealTimeStatistics**](docs/DefaultApi.md#fetchworkspacerealtimestatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/RealTimeStatistics | 
*DefaultApi* | [**FetchWorkspaceStatistics**](docs/DefaultApi.md#fetchworkspacestatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Statistics | 
*DefaultApi* | [**ListActivity**](docs/DefaultApi.md#listactivity) | **Get** /v1/Workspaces/{WorkspaceSid}/Activities | 
*DefaultApi* | [**ListEvent**](docs/DefaultApi.md#listevent) | **Get** /v1/Workspaces/{WorkspaceSid}/Events | 
*DefaultApi* | [**ListTask**](docs/DefaultApi.md#listtask) | **Get** /v1/Workspaces/{WorkspaceSid}/Tasks | 
*DefaultApi* | [**ListTaskChannel**](docs/DefaultApi.md#listtaskchannel) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskChannels | 
*DefaultApi* | [**ListTaskQueue**](docs/DefaultApi.md#listtaskqueue) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues | 
*DefaultApi* | [**ListTaskQueuesStatistics**](docs/DefaultApi.md#listtaskqueuesstatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues/Statistics | 
*DefaultApi* | [**ListTaskReservation**](docs/DefaultApi.md#listtaskreservation) | **Get** /v1/Workspaces/{WorkspaceSid}/Tasks/{TaskSid}/Reservations | 
*DefaultApi* | [**ListWorker**](docs/DefaultApi.md#listworker) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers | 
*DefaultApi* | [**ListWorkerChannel**](docs/DefaultApi.md#listworkerchannel) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Channels | 
*DefaultApi* | [**ListWorkerReservation**](docs/DefaultApi.md#listworkerreservation) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Reservations | 
*DefaultApi* | [**ListWorkflow**](docs/DefaultApi.md#listworkflow) | **Get** /v1/Workspaces/{WorkspaceSid}/Workflows | 
*DefaultApi* | [**ListWorkspace**](docs/DefaultApi.md#listworkspace) | **Get** /v1/Workspaces | 
*DefaultApi* | [**UpdateActivity**](docs/DefaultApi.md#updateactivity) | **Post** /v1/Workspaces/{WorkspaceSid}/Activities/{Sid} | 
*DefaultApi* | [**UpdateTask**](docs/DefaultApi.md#updatetask) | **Post** /v1/Workspaces/{WorkspaceSid}/Tasks/{Sid} | 
*DefaultApi* | [**UpdateTaskChannel**](docs/DefaultApi.md#updatetaskchannel) | **Post** /v1/Workspaces/{WorkspaceSid}/TaskChannels/{Sid} | 
*DefaultApi* | [**UpdateTaskQueue**](docs/DefaultApi.md#updatetaskqueue) | **Post** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{Sid} | 
*DefaultApi* | [**UpdateTaskReservation**](docs/DefaultApi.md#updatetaskreservation) | **Post** /v1/Workspaces/{WorkspaceSid}/Tasks/{TaskSid}/Reservations/{Sid} | 
*DefaultApi* | [**UpdateWorker**](docs/DefaultApi.md#updateworker) | **Post** /v1/Workspaces/{WorkspaceSid}/Workers/{Sid} | 
*DefaultApi* | [**UpdateWorkerChannel**](docs/DefaultApi.md#updateworkerchannel) | **Post** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Channels/{Sid} | 
*DefaultApi* | [**UpdateWorkerReservation**](docs/DefaultApi.md#updateworkerreservation) | **Post** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Reservations/{Sid} | 
*DefaultApi* | [**UpdateWorkflow**](docs/DefaultApi.md#updateworkflow) | **Post** /v1/Workspaces/{WorkspaceSid}/Workflows/{Sid} | 
*DefaultApi* | [**UpdateWorkspace**](docs/DefaultApi.md#updateworkspace) | **Post** /v1/Workspaces/{Sid} | 


## Documentation For Models

 - [ListActivityResponse](docs/ListActivityResponse.md)
 - [ListEventResponse](docs/ListEventResponse.md)
 - [ListTaskChannelResponse](docs/ListTaskChannelResponse.md)
 - [ListTaskQueueResponse](docs/ListTaskQueueResponse.md)
 - [ListTaskQueuesStatisticsResponse](docs/ListTaskQueuesStatisticsResponse.md)
 - [ListTaskReservationResponse](docs/ListTaskReservationResponse.md)
 - [ListTaskResponse](docs/ListTaskResponse.md)
 - [ListWorkerChannelResponse](docs/ListWorkerChannelResponse.md)
 - [ListWorkerReservationResponse](docs/ListWorkerReservationResponse.md)
 - [ListWorkerResponse](docs/ListWorkerResponse.md)
 - [ListWorkflowResponse](docs/ListWorkflowResponse.md)
 - [ListWorkspaceResponse](docs/ListWorkspaceResponse.md)
 - [ListWorkspaceResponseMeta](docs/ListWorkspaceResponseMeta.md)
 - [TaskrouterV1Workspace](docs/TaskrouterV1Workspace.md)
 - [TaskrouterV1WorkspaceActivity](docs/TaskrouterV1WorkspaceActivity.md)
 - [TaskrouterV1WorkspaceEvent](docs/TaskrouterV1WorkspaceEvent.md)
 - [TaskrouterV1WorkspaceTask](docs/TaskrouterV1WorkspaceTask.md)
 - [TaskrouterV1WorkspaceTaskChannel](docs/TaskrouterV1WorkspaceTaskChannel.md)
 - [TaskrouterV1WorkspaceTaskQueue](docs/TaskrouterV1WorkspaceTaskQueue.md)
 - [TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics](docs/TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics.md)
 - [TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics](docs/TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics.md)
 - [TaskrouterV1WorkspaceTaskQueueTaskQueueStatistics](docs/TaskrouterV1WorkspaceTaskQueueTaskQueueStatistics.md)
 - [TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics](docs/TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics.md)
 - [TaskrouterV1WorkspaceTaskTaskReservation](docs/TaskrouterV1WorkspaceTaskTaskReservation.md)
 - [TaskrouterV1WorkspaceWorker](docs/TaskrouterV1WorkspaceWorker.md)
 - [TaskrouterV1WorkspaceWorkerWorkerChannel](docs/TaskrouterV1WorkspaceWorkerWorkerChannel.md)
 - [TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics](docs/TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics.md)
 - [TaskrouterV1WorkspaceWorkerWorkerReservation](docs/TaskrouterV1WorkspaceWorkerWorkerReservation.md)
 - [TaskrouterV1WorkspaceWorkerWorkerStatistics](docs/TaskrouterV1WorkspaceWorkerWorkerStatistics.md)
 - [TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics](docs/TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics.md)
 - [TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics](docs/TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics.md)
 - [TaskrouterV1WorkspaceWorkflow](docs/TaskrouterV1WorkspaceWorkflow.md)
 - [TaskrouterV1WorkspaceWorkflowWorkflowCumulativeStatistics](docs/TaskrouterV1WorkspaceWorkflowWorkflowCumulativeStatistics.md)
 - [TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics](docs/TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics.md)
 - [TaskrouterV1WorkspaceWorkflowWorkflowStatistics](docs/TaskrouterV1WorkspaceWorkflowWorkflowStatistics.md)
 - [TaskrouterV1WorkspaceWorkspaceCumulativeStatistics](docs/TaskrouterV1WorkspaceWorkspaceCumulativeStatistics.md)
 - [TaskrouterV1WorkspaceWorkspaceRealTimeStatistics](docs/TaskrouterV1WorkspaceWorkspaceRealTimeStatistics.md)
 - [TaskrouterV1WorkspaceWorkspaceStatistics](docs/TaskrouterV1WorkspaceWorkspaceStatistics.md)


## Documentation For Authorization



## accountSid_authToken

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Author

support@twilio.com

