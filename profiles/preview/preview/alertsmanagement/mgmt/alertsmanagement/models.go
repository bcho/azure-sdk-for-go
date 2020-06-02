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

package alertsmanagement

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/alertsmanagement/mgmt/2019-05-05/alertsmanagement"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ActionRuleStatus = original.ActionRuleStatus

const (
	Disabled ActionRuleStatus = original.Disabled
	Enabled  ActionRuleStatus = original.Enabled
)

type AlertModificationEvent = original.AlertModificationEvent

const (
	ActionRuleSuppressed   AlertModificationEvent = original.ActionRuleSuppressed
	ActionRuleTriggered    AlertModificationEvent = original.ActionRuleTriggered
	ActionsFailed          AlertModificationEvent = original.ActionsFailed
	ActionsSuppressed      AlertModificationEvent = original.ActionsSuppressed
	ActionsTriggered       AlertModificationEvent = original.ActionsTriggered
	AlertCreated           AlertModificationEvent = original.AlertCreated
	MonitorConditionChange AlertModificationEvent = original.MonitorConditionChange
	SeverityChange         AlertModificationEvent = original.SeverityChange
	StateChange            AlertModificationEvent = original.StateChange
)

type AlertState = original.AlertState

const (
	AlertStateAcknowledged AlertState = original.AlertStateAcknowledged
	AlertStateClosed       AlertState = original.AlertStateClosed
	AlertStateNew          AlertState = original.AlertStateNew
)

type AlertsSortByFields = original.AlertsSortByFields

const (
	AlertsSortByFieldsAlertState           AlertsSortByFields = original.AlertsSortByFieldsAlertState
	AlertsSortByFieldsLastModifiedDateTime AlertsSortByFields = original.AlertsSortByFieldsLastModifiedDateTime
	AlertsSortByFieldsMonitorCondition     AlertsSortByFields = original.AlertsSortByFieldsMonitorCondition
	AlertsSortByFieldsName                 AlertsSortByFields = original.AlertsSortByFieldsName
	AlertsSortByFieldsSeverity             AlertsSortByFields = original.AlertsSortByFieldsSeverity
	AlertsSortByFieldsStartDateTime        AlertsSortByFields = original.AlertsSortByFieldsStartDateTime
	AlertsSortByFieldsTargetResource       AlertsSortByFields = original.AlertsSortByFieldsTargetResource
	AlertsSortByFieldsTargetResourceGroup  AlertsSortByFields = original.AlertsSortByFieldsTargetResourceGroup
	AlertsSortByFieldsTargetResourceName   AlertsSortByFields = original.AlertsSortByFieldsTargetResourceName
	AlertsSortByFieldsTargetResourceType   AlertsSortByFields = original.AlertsSortByFieldsTargetResourceType
)

type AlertsSummaryGroupByFields = original.AlertsSummaryGroupByFields

const (
	AlertsSummaryGroupByFieldsAlertRule        AlertsSummaryGroupByFields = original.AlertsSummaryGroupByFieldsAlertRule
	AlertsSummaryGroupByFieldsAlertState       AlertsSummaryGroupByFields = original.AlertsSummaryGroupByFieldsAlertState
	AlertsSummaryGroupByFieldsMonitorCondition AlertsSummaryGroupByFields = original.AlertsSummaryGroupByFieldsMonitorCondition
	AlertsSummaryGroupByFieldsMonitorService   AlertsSummaryGroupByFields = original.AlertsSummaryGroupByFieldsMonitorService
	AlertsSummaryGroupByFieldsSeverity         AlertsSummaryGroupByFields = original.AlertsSummaryGroupByFieldsSeverity
	AlertsSummaryGroupByFieldsSignalType       AlertsSummaryGroupByFields = original.AlertsSummaryGroupByFieldsSignalType
)

type MetadataIdentifier = original.MetadataIdentifier

const (
	MetadataIdentifierAlertsMetaDataProperties MetadataIdentifier = original.MetadataIdentifierAlertsMetaDataProperties
	MetadataIdentifierMonitorServiceList       MetadataIdentifier = original.MetadataIdentifierMonitorServiceList
)

type MonitorCondition = original.MonitorCondition

const (
	Fired    MonitorCondition = original.Fired
	Resolved MonitorCondition = original.Resolved
)

type MonitorService = original.MonitorService

const (
	ActivityLogAdministrative MonitorService = original.ActivityLogAdministrative
	ActivityLogAutoscale      MonitorService = original.ActivityLogAutoscale
	ActivityLogPolicy         MonitorService = original.ActivityLogPolicy
	ActivityLogRecommendation MonitorService = original.ActivityLogRecommendation
	ActivityLogSecurity       MonitorService = original.ActivityLogSecurity
	ApplicationInsights       MonitorService = original.ApplicationInsights
	LogAnalytics              MonitorService = original.LogAnalytics
	Nagios                    MonitorService = original.Nagios
	Platform                  MonitorService = original.Platform
	SCOM                      MonitorService = original.SCOM
	ServiceHealth             MonitorService = original.ServiceHealth
	SmartDetector             MonitorService = original.SmartDetector
	VMInsights                MonitorService = original.VMInsights
	Zabbix                    MonitorService = original.Zabbix
)

type Operator = original.Operator

const (
	Contains       Operator = original.Contains
	DoesNotContain Operator = original.DoesNotContain
	Equals         Operator = original.Equals
	NotEquals      Operator = original.NotEquals
)

type ScopeType = original.ScopeType

const (
	ScopeTypeResource      ScopeType = original.ScopeTypeResource
	ScopeTypeResourceGroup ScopeType = original.ScopeTypeResourceGroup
	ScopeTypeSubscription  ScopeType = original.ScopeTypeSubscription
)

type Severity = original.Severity

const (
	Sev0 Severity = original.Sev0
	Sev1 Severity = original.Sev1
	Sev2 Severity = original.Sev2
	Sev3 Severity = original.Sev3
	Sev4 Severity = original.Sev4
)

type SignalType = original.SignalType

const (
	Log     SignalType = original.Log
	Metric  SignalType = original.Metric
	Unknown SignalType = original.Unknown
)

type SmartGroupModificationEvent = original.SmartGroupModificationEvent

const (
	SmartGroupModificationEventAlertAdded        SmartGroupModificationEvent = original.SmartGroupModificationEventAlertAdded
	SmartGroupModificationEventAlertRemoved      SmartGroupModificationEvent = original.SmartGroupModificationEventAlertRemoved
	SmartGroupModificationEventSmartGroupCreated SmartGroupModificationEvent = original.SmartGroupModificationEventSmartGroupCreated
	SmartGroupModificationEventStateChange       SmartGroupModificationEvent = original.SmartGroupModificationEventStateChange
)

type SmartGroupsSortByFields = original.SmartGroupsSortByFields

const (
	SmartGroupsSortByFieldsAlertsCount          SmartGroupsSortByFields = original.SmartGroupsSortByFieldsAlertsCount
	SmartGroupsSortByFieldsLastModifiedDateTime SmartGroupsSortByFields = original.SmartGroupsSortByFieldsLastModifiedDateTime
	SmartGroupsSortByFieldsSeverity             SmartGroupsSortByFields = original.SmartGroupsSortByFieldsSeverity
	SmartGroupsSortByFieldsStartDateTime        SmartGroupsSortByFields = original.SmartGroupsSortByFieldsStartDateTime
	SmartGroupsSortByFieldsState                SmartGroupsSortByFields = original.SmartGroupsSortByFieldsState
)

type State = original.State

const (
	StateAcknowledged State = original.StateAcknowledged
	StateClosed       State = original.StateClosed
	StateNew          State = original.StateNew
)

type SuppressionType = original.SuppressionType

const (
	Always  SuppressionType = original.Always
	Daily   SuppressionType = original.Daily
	Monthly SuppressionType = original.Monthly
	Once    SuppressionType = original.Once
	Weekly  SuppressionType = original.Weekly
)

type TimeRange = original.TimeRange

const (
	Oned       TimeRange = original.Oned
	Oneh       TimeRange = original.Oneh
	Sevend     TimeRange = original.Sevend
	ThreeZerod TimeRange = original.ThreeZerod
)

type Type = original.Type

const (
	TypeActionGroup          Type = original.TypeActionGroup
	TypeActionRuleProperties Type = original.TypeActionRuleProperties
	TypeDiagnostics          Type = original.TypeDiagnostics
	TypeSuppression          Type = original.TypeSuppression
)

type ActionGroup = original.ActionGroup
type ActionRule = original.ActionRule
type ActionRuleProperties = original.ActionRuleProperties
type ActionRulesClient = original.ActionRulesClient
type ActionRulesList = original.ActionRulesList
type ActionRulesListIterator = original.ActionRulesListIterator
type ActionRulesListPage = original.ActionRulesListPage
type Alert = original.Alert
type AlertModification = original.AlertModification
type AlertModificationItem = original.AlertModificationItem
type AlertModificationProperties = original.AlertModificationProperties
type AlertProperties = original.AlertProperties
type AlertsClient = original.AlertsClient
type AlertsList = original.AlertsList
type AlertsListIterator = original.AlertsListIterator
type AlertsListPage = original.AlertsListPage
type AlertsMetaData = original.AlertsMetaData
type AlertsMetaDataProperties = original.AlertsMetaDataProperties
type AlertsSummary = original.AlertsSummary
type AlertsSummaryGroup = original.AlertsSummaryGroup
type AlertsSummaryGroupItem = original.AlertsSummaryGroupItem
type BaseClient = original.BaseClient
type BasicActionRuleProperties = original.BasicActionRuleProperties
type BasicAlertsMetaDataProperties = original.BasicAlertsMetaDataProperties
type Bool = original.Bool
type Condition = original.Condition
type Conditions = original.Conditions
type Diagnostics = original.Diagnostics
type ErrorResponse = original.ErrorResponse
type ErrorResponseBody = original.ErrorResponseBody
type Essentials = original.Essentials
type ManagedResource = original.ManagedResource
type MonitorServiceDetails = original.MonitorServiceDetails
type MonitorServiceList = original.MonitorServiceList
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationsClient = original.OperationsClient
type OperationsList = original.OperationsList
type OperationsListIterator = original.OperationsListIterator
type OperationsListPage = original.OperationsListPage
type PatchObject = original.PatchObject
type PatchProperties = original.PatchProperties
type Resource = original.Resource
type Scope = original.Scope
type SmartGroup = original.SmartGroup
type SmartGroupAggregatedProperty = original.SmartGroupAggregatedProperty
type SmartGroupModification = original.SmartGroupModification
type SmartGroupModificationItem = original.SmartGroupModificationItem
type SmartGroupModificationProperties = original.SmartGroupModificationProperties
type SmartGroupProperties = original.SmartGroupProperties
type SmartGroupsClient = original.SmartGroupsClient
type SmartGroupsList = original.SmartGroupsList
type SmartGroupsListIterator = original.SmartGroupsListIterator
type SmartGroupsListPage = original.SmartGroupsListPage
type Suppression = original.Suppression
type SuppressionConfig = original.SuppressionConfig
type SuppressionSchedule = original.SuppressionSchedule

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewActionRulesClient(subscriptionID string) ActionRulesClient {
	return original.NewActionRulesClient(subscriptionID)
}
func NewActionRulesClientWithBaseURI(baseURI string, subscriptionID string) ActionRulesClient {
	return original.NewActionRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewActionRulesListIterator(page ActionRulesListPage) ActionRulesListIterator {
	return original.NewActionRulesListIterator(page)
}
func NewActionRulesListPage(getNextPage func(context.Context, ActionRulesList) (ActionRulesList, error)) ActionRulesListPage {
	return original.NewActionRulesListPage(getNextPage)
}
func NewAlertsClient(subscriptionID string) AlertsClient {
	return original.NewAlertsClient(subscriptionID)
}
func NewAlertsClientWithBaseURI(baseURI string, subscriptionID string) AlertsClient {
	return original.NewAlertsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAlertsListIterator(page AlertsListPage) AlertsListIterator {
	return original.NewAlertsListIterator(page)
}
func NewAlertsListPage(getNextPage func(context.Context, AlertsList) (AlertsList, error)) AlertsListPage {
	return original.NewAlertsListPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsListIterator(page OperationsListPage) OperationsListIterator {
	return original.NewOperationsListIterator(page)
}
func NewOperationsListPage(getNextPage func(context.Context, OperationsList) (OperationsList, error)) OperationsListPage {
	return original.NewOperationsListPage(getNextPage)
}
func NewSmartGroupsClient(subscriptionID string) SmartGroupsClient {
	return original.NewSmartGroupsClient(subscriptionID)
}
func NewSmartGroupsClientWithBaseURI(baseURI string, subscriptionID string) SmartGroupsClient {
	return original.NewSmartGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSmartGroupsListIterator(page SmartGroupsListPage) SmartGroupsListIterator {
	return original.NewSmartGroupsListIterator(page)
}
func NewSmartGroupsListPage(getNextPage func(context.Context, SmartGroupsList) (SmartGroupsList, error)) SmartGroupsListPage {
	return original.NewSmartGroupsListPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleActionRuleStatusValues() []ActionRuleStatus {
	return original.PossibleActionRuleStatusValues()
}
func PossibleAlertModificationEventValues() []AlertModificationEvent {
	return original.PossibleAlertModificationEventValues()
}
func PossibleAlertStateValues() []AlertState {
	return original.PossibleAlertStateValues()
}
func PossibleAlertsSortByFieldsValues() []AlertsSortByFields {
	return original.PossibleAlertsSortByFieldsValues()
}
func PossibleAlertsSummaryGroupByFieldsValues() []AlertsSummaryGroupByFields {
	return original.PossibleAlertsSummaryGroupByFieldsValues()
}
func PossibleMetadataIdentifierValues() []MetadataIdentifier {
	return original.PossibleMetadataIdentifierValues()
}
func PossibleMonitorConditionValues() []MonitorCondition {
	return original.PossibleMonitorConditionValues()
}
func PossibleMonitorServiceValues() []MonitorService {
	return original.PossibleMonitorServiceValues()
}
func PossibleOperatorValues() []Operator {
	return original.PossibleOperatorValues()
}
func PossibleScopeTypeValues() []ScopeType {
	return original.PossibleScopeTypeValues()
}
func PossibleSeverityValues() []Severity {
	return original.PossibleSeverityValues()
}
func PossibleSignalTypeValues() []SignalType {
	return original.PossibleSignalTypeValues()
}
func PossibleSmartGroupModificationEventValues() []SmartGroupModificationEvent {
	return original.PossibleSmartGroupModificationEventValues()
}
func PossibleSmartGroupsSortByFieldsValues() []SmartGroupsSortByFields {
	return original.PossibleSmartGroupsSortByFieldsValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleSuppressionTypeValues() []SuppressionType {
	return original.PossibleSuppressionTypeValues()
}
func PossibleTimeRangeValues() []TimeRange {
	return original.PossibleTimeRangeValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
