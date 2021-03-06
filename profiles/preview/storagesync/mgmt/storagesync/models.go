// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package storagesync

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/storagesync/mgmt/2019-10-01/storagesync"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ChangeDetectionMode = original.ChangeDetectionMode

const (
	Default   ChangeDetectionMode = original.Default
	Recursive ChangeDetectionMode = original.Recursive
)

type CloudTiering = original.CloudTiering

const (
	Off CloudTiering = original.Off
	On  CloudTiering = original.On
)

type CloudTiering1 = original.CloudTiering1

const (
	CloudTiering1Off CloudTiering1 = original.CloudTiering1Off
	CloudTiering1On  CloudTiering1 = original.CloudTiering1On
)

type CloudTiering2 = original.CloudTiering2

const (
	CloudTiering2Off CloudTiering2 = original.CloudTiering2Off
	CloudTiering2On  CloudTiering2 = original.CloudTiering2On
)

type CombinedHealth = original.CombinedHealth

const (
	CombinedHealthError                                    CombinedHealth = original.CombinedHealthError
	CombinedHealthHealthy                                  CombinedHealth = original.CombinedHealthHealthy
	CombinedHealthNoActivity                               CombinedHealth = original.CombinedHealthNoActivity
	CombinedHealthSyncBlockedForChangeDetectionPostRestore CombinedHealth = original.CombinedHealthSyncBlockedForChangeDetectionPostRestore
	CombinedHealthSyncBlockedForRestore                    CombinedHealth = original.CombinedHealthSyncBlockedForRestore
)

type DownloadHealth = original.DownloadHealth

const (
	DownloadHealthError                                    DownloadHealth = original.DownloadHealthError
	DownloadHealthHealthy                                  DownloadHealth = original.DownloadHealthHealthy
	DownloadHealthNoActivity                               DownloadHealth = original.DownloadHealthNoActivity
	DownloadHealthSyncBlockedForChangeDetectionPostRestore DownloadHealth = original.DownloadHealthSyncBlockedForChangeDetectionPostRestore
	DownloadHealthSyncBlockedForRestore                    DownloadHealth = original.DownloadHealthSyncBlockedForRestore
)

type Health = original.Health

const (
	HealthError   Health = original.HealthError
	HealthHealthy Health = original.HealthHealthy
)

type NameAvailabilityReason = original.NameAvailabilityReason

const (
	AlreadyExists NameAvailabilityReason = original.AlreadyExists
	Invalid       NameAvailabilityReason = original.Invalid
)

type OfflineDataTransfer = original.OfflineDataTransfer

const (
	OfflineDataTransferOff OfflineDataTransfer = original.OfflineDataTransferOff
	OfflineDataTransferOn  OfflineDataTransfer = original.OfflineDataTransferOn
)

type OfflineDataTransfer1 = original.OfflineDataTransfer1

const (
	OfflineDataTransfer1Off OfflineDataTransfer1 = original.OfflineDataTransfer1Off
	OfflineDataTransfer1On  OfflineDataTransfer1 = original.OfflineDataTransfer1On
)

type OfflineDataTransfer2 = original.OfflineDataTransfer2

const (
	OfflineDataTransfer2Off OfflineDataTransfer2 = original.OfflineDataTransfer2Off
	OfflineDataTransfer2On  OfflineDataTransfer2 = original.OfflineDataTransfer2On
)

type OfflineDataTransferStatus = original.OfflineDataTransferStatus

const (
	Complete   OfflineDataTransferStatus = original.Complete
	InProgress OfflineDataTransferStatus = original.InProgress
	NotRunning OfflineDataTransferStatus = original.NotRunning
	Stopping   OfflineDataTransferStatus = original.Stopping
)

type Operation = original.Operation

const (
	Cancel Operation = original.Cancel
	Do     Operation = original.Do
	Undo   Operation = original.Undo
)

type Reason = original.Reason

const (
	Deleted      Reason = original.Deleted
	Registered   Reason = original.Registered
	Suspended    Reason = original.Suspended
	Unregistered Reason = original.Unregistered
	Warned       Reason = original.Warned
)

type Status = original.Status

const (
	Aborted   Status = original.Aborted
	Active    Status = original.Active
	Expired   Status = original.Expired
	Failed    Status = original.Failed
	Succeeded Status = original.Succeeded
)

type SyncActivity = original.SyncActivity

const (
	Download          SyncActivity = original.Download
	Upload            SyncActivity = original.Upload
	UploadAndDownload SyncActivity = original.UploadAndDownload
)

type UploadHealth = original.UploadHealth

const (
	UploadHealthError                                    UploadHealth = original.UploadHealthError
	UploadHealthHealthy                                  UploadHealth = original.UploadHealthHealthy
	UploadHealthNoActivity                               UploadHealth = original.UploadHealthNoActivity
	UploadHealthSyncBlockedForChangeDetectionPostRestore UploadHealth = original.UploadHealthSyncBlockedForChangeDetectionPostRestore
	UploadHealthSyncBlockedForRestore                    UploadHealth = original.UploadHealthSyncBlockedForRestore
)

type APIError = original.APIError
type AzureEntityResource = original.AzureEntityResource
type BackupRequest = original.BackupRequest
type BaseClient = original.BaseClient
type CheckNameAvailabilityParameters = original.CheckNameAvailabilityParameters
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type CloudEndpoint = original.CloudEndpoint
type CloudEndpointArray = original.CloudEndpointArray
type CloudEndpointCreateParameters = original.CloudEndpointCreateParameters
type CloudEndpointCreateParametersProperties = original.CloudEndpointCreateParametersProperties
type CloudEndpointProperties = original.CloudEndpointProperties
type CloudEndpointsClient = original.CloudEndpointsClient
type CloudEndpointsCreateFuture = original.CloudEndpointsCreateFuture
type CloudEndpointsDeleteFuture = original.CloudEndpointsDeleteFuture
type CloudEndpointsPostBackupFuture = original.CloudEndpointsPostBackupFuture
type CloudEndpointsPostRestoreFuture = original.CloudEndpointsPostRestoreFuture
type CloudEndpointsPreBackupFuture = original.CloudEndpointsPreBackupFuture
type CloudEndpointsPreRestoreFuture = original.CloudEndpointsPreRestoreFuture
type CloudEndpointsTriggerChangeDetectionFuture = original.CloudEndpointsTriggerChangeDetectionFuture
type CloudTieringCachePerformance = original.CloudTieringCachePerformance
type CloudTieringDatePolicyStatus = original.CloudTieringDatePolicyStatus
type CloudTieringFilesNotTiering = original.CloudTieringFilesNotTiering
type CloudTieringSpaceSavings = original.CloudTieringSpaceSavings
type CloudTieringVolumeFreeSpacePolicyStatus = original.CloudTieringVolumeFreeSpacePolicyStatus
type Error = original.Error
type ErrorDetails = original.ErrorDetails
type FilesNotTieringError = original.FilesNotTieringError
type OperationDisplayInfo = original.OperationDisplayInfo
type OperationDisplayResource = original.OperationDisplayResource
type OperationEntity = original.OperationEntity
type OperationEntityListResult = original.OperationEntityListResult
type OperationEntityListResultIterator = original.OperationEntityListResultIterator
type OperationEntityListResultPage = original.OperationEntityListResultPage
type OperationStatus = original.OperationStatus
type OperationStatusClient = original.OperationStatusClient
type OperationsClient = original.OperationsClient
type PostBackupResponse = original.PostBackupResponse
type PostBackupResponseProperties = original.PostBackupResponseProperties
type PostRestoreRequest = original.PostRestoreRequest
type PreRestoreRequest = original.PreRestoreRequest
type ProxyResource = original.ProxyResource
type RecallActionParameters = original.RecallActionParameters
type RegisteredServer = original.RegisteredServer
type RegisteredServerArray = original.RegisteredServerArray
type RegisteredServerCreateParameters = original.RegisteredServerCreateParameters
type RegisteredServerCreateParametersProperties = original.RegisteredServerCreateParametersProperties
type RegisteredServerProperties = original.RegisteredServerProperties
type RegisteredServersClient = original.RegisteredServersClient
type RegisteredServersCreateFuture = original.RegisteredServersCreateFuture
type RegisteredServersDeleteFuture = original.RegisteredServersDeleteFuture
type RegisteredServersTriggerRolloverFuture = original.RegisteredServersTriggerRolloverFuture
type Resource = original.Resource
type ResourcesMoveInfo = original.ResourcesMoveInfo
type RestoreFileSpec = original.RestoreFileSpec
type ServerEndpoint = original.ServerEndpoint
type ServerEndpointArray = original.ServerEndpointArray
type ServerEndpointCloudTieringStatus = original.ServerEndpointCloudTieringStatus
type ServerEndpointCreateParameters = original.ServerEndpointCreateParameters
type ServerEndpointCreateParametersProperties = original.ServerEndpointCreateParametersProperties
type ServerEndpointFilesNotSyncingError = original.ServerEndpointFilesNotSyncingError
type ServerEndpointProperties = original.ServerEndpointProperties
type ServerEndpointRecallError = original.ServerEndpointRecallError
type ServerEndpointRecallStatus = original.ServerEndpointRecallStatus
type ServerEndpointSyncActivityStatus = original.ServerEndpointSyncActivityStatus
type ServerEndpointSyncSessionStatus = original.ServerEndpointSyncSessionStatus
type ServerEndpointSyncStatus = original.ServerEndpointSyncStatus
type ServerEndpointUpdateParameters = original.ServerEndpointUpdateParameters
type ServerEndpointUpdateProperties = original.ServerEndpointUpdateProperties
type ServerEndpointsClient = original.ServerEndpointsClient
type ServerEndpointsCreateFuture = original.ServerEndpointsCreateFuture
type ServerEndpointsDeleteFuture = original.ServerEndpointsDeleteFuture
type ServerEndpointsRecallActionFuture = original.ServerEndpointsRecallActionFuture
type ServerEndpointsUpdateFuture = original.ServerEndpointsUpdateFuture
type Service = original.Service
type ServiceArray = original.ServiceArray
type ServiceCreateParameters = original.ServiceCreateParameters
type ServiceProperties = original.ServiceProperties
type ServiceUpdateParameters = original.ServiceUpdateParameters
type ServicesClient = original.ServicesClient
type SubscriptionState = original.SubscriptionState
type SyncGroup = original.SyncGroup
type SyncGroupArray = original.SyncGroupArray
type SyncGroupCreateParameters = original.SyncGroupCreateParameters
type SyncGroupProperties = original.SyncGroupProperties
type SyncGroupsClient = original.SyncGroupsClient
type TrackedResource = original.TrackedResource
type TriggerChangeDetectionParameters = original.TriggerChangeDetectionParameters
type TriggerRolloverRequest = original.TriggerRolloverRequest
type Workflow = original.Workflow
type WorkflowArray = original.WorkflowArray
type WorkflowProperties = original.WorkflowProperties
type WorkflowsClient = original.WorkflowsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewCloudEndpointsClient(subscriptionID string) CloudEndpointsClient {
	return original.NewCloudEndpointsClient(subscriptionID)
}
func NewCloudEndpointsClientWithBaseURI(baseURI string, subscriptionID string) CloudEndpointsClient {
	return original.NewCloudEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationEntityListResultIterator(page OperationEntityListResultPage) OperationEntityListResultIterator {
	return original.NewOperationEntityListResultIterator(page)
}
func NewOperationEntityListResultPage(getNextPage func(context.Context, OperationEntityListResult) (OperationEntityListResult, error)) OperationEntityListResultPage {
	return original.NewOperationEntityListResultPage(getNextPage)
}
func NewOperationStatusClient(subscriptionID string) OperationStatusClient {
	return original.NewOperationStatusClient(subscriptionID)
}
func NewOperationStatusClientWithBaseURI(baseURI string, subscriptionID string) OperationStatusClient {
	return original.NewOperationStatusClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRegisteredServersClient(subscriptionID string) RegisteredServersClient {
	return original.NewRegisteredServersClient(subscriptionID)
}
func NewRegisteredServersClientWithBaseURI(baseURI string, subscriptionID string) RegisteredServersClient {
	return original.NewRegisteredServersClientWithBaseURI(baseURI, subscriptionID)
}
func NewServerEndpointsClient(subscriptionID string) ServerEndpointsClient {
	return original.NewServerEndpointsClient(subscriptionID)
}
func NewServerEndpointsClientWithBaseURI(baseURI string, subscriptionID string) ServerEndpointsClient {
	return original.NewServerEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServicesClient(subscriptionID string) ServicesClient {
	return original.NewServicesClient(subscriptionID)
}
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSyncGroupsClient(subscriptionID string) SyncGroupsClient {
	return original.NewSyncGroupsClient(subscriptionID)
}
func NewSyncGroupsClientWithBaseURI(baseURI string, subscriptionID string) SyncGroupsClient {
	return original.NewSyncGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewWorkflowsClient(subscriptionID string) WorkflowsClient {
	return original.NewWorkflowsClient(subscriptionID)
}
func NewWorkflowsClientWithBaseURI(baseURI string, subscriptionID string) WorkflowsClient {
	return original.NewWorkflowsClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleChangeDetectionModeValues() []ChangeDetectionMode {
	return original.PossibleChangeDetectionModeValues()
}
func PossibleCloudTiering1Values() []CloudTiering1 {
	return original.PossibleCloudTiering1Values()
}
func PossibleCloudTiering2Values() []CloudTiering2 {
	return original.PossibleCloudTiering2Values()
}
func PossibleCloudTieringValues() []CloudTiering {
	return original.PossibleCloudTieringValues()
}
func PossibleCombinedHealthValues() []CombinedHealth {
	return original.PossibleCombinedHealthValues()
}
func PossibleDownloadHealthValues() []DownloadHealth {
	return original.PossibleDownloadHealthValues()
}
func PossibleHealthValues() []Health {
	return original.PossibleHealthValues()
}
func PossibleNameAvailabilityReasonValues() []NameAvailabilityReason {
	return original.PossibleNameAvailabilityReasonValues()
}
func PossibleOfflineDataTransfer1Values() []OfflineDataTransfer1 {
	return original.PossibleOfflineDataTransfer1Values()
}
func PossibleOfflineDataTransfer2Values() []OfflineDataTransfer2 {
	return original.PossibleOfflineDataTransfer2Values()
}
func PossibleOfflineDataTransferStatusValues() []OfflineDataTransferStatus {
	return original.PossibleOfflineDataTransferStatusValues()
}
func PossibleOfflineDataTransferValues() []OfflineDataTransfer {
	return original.PossibleOfflineDataTransferValues()
}
func PossibleOperationValues() []Operation {
	return original.PossibleOperationValues()
}
func PossibleReasonValues() []Reason {
	return original.PossibleReasonValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleSyncActivityValues() []SyncActivity {
	return original.PossibleSyncActivityValues()
}
func PossibleUploadHealthValues() []UploadHealth {
	return original.PossibleUploadHealthValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
