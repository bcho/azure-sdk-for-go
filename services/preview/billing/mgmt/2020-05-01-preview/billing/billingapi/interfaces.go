package billingapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/billing/mgmt/2020-05-01-preview/billing"
)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	Get(ctx context.Context, billingAccountName string, expand string) (result billing.Account, err error)
	List(ctx context.Context, expand string) (result billing.AccountListResultPage, err error)
	ListComplete(ctx context.Context, expand string) (result billing.AccountListResultIterator, err error)
	ListInvoiceSectionsByCreateSubscriptionPermission(ctx context.Context, billingAccountName string) (result billing.InvoiceSectionListWithCreateSubPermissionResultPage, err error)
	ListInvoiceSectionsByCreateSubscriptionPermissionComplete(ctx context.Context, billingAccountName string) (result billing.InvoiceSectionListWithCreateSubPermissionResultIterator, err error)
	Update(ctx context.Context, billingAccountName string, parameters billing.AccountUpdateRequest) (result billing.AccountsUpdateFuture, err error)
}

var _ AccountsClientAPI = (*billing.AccountsClient)(nil)

// AddressClientAPI contains the set of methods on the AddressClient type.
type AddressClientAPI interface {
	Validate(ctx context.Context, address billing.AddressDetails) (result billing.ValidateAddressResponse, err error)
}

var _ AddressClientAPI = (*billing.AddressClient)(nil)

// AvailableBalancesClientAPI contains the set of methods on the AvailableBalancesClient type.
type AvailableBalancesClientAPI interface {
	Get(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.AvailableBalance, err error)
}

var _ AvailableBalancesClientAPI = (*billing.AvailableBalancesClient)(nil)

// InstructionsClientAPI contains the set of methods on the InstructionsClient type.
type InstructionsClientAPI interface {
	Get(ctx context.Context, billingAccountName string, billingProfileName string, instructionName string) (result billing.Instruction, err error)
	ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.InstructionListResultPage, err error)
	ListByBillingProfileComplete(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.InstructionListResultIterator, err error)
	Put(ctx context.Context, billingAccountName string, billingProfileName string, instructionName string, parameters billing.Instruction) (result billing.Instruction, err error)
}

var _ InstructionsClientAPI = (*billing.InstructionsClient)(nil)

// ProfilesClientAPI contains the set of methods on the ProfilesClient type.
type ProfilesClientAPI interface {
	CreateOrUpdate(ctx context.Context, billingAccountName string, billingProfileName string, parameters billing.Profile) (result billing.ProfilesCreateOrUpdateFuture, err error)
	Get(ctx context.Context, billingAccountName string, billingProfileName string, expand string) (result billing.Profile, err error)
	ListByBillingAccount(ctx context.Context, billingAccountName string, expand string) (result billing.ProfileListResultPage, err error)
	ListByBillingAccountComplete(ctx context.Context, billingAccountName string, expand string) (result billing.ProfileListResultIterator, err error)
}

var _ ProfilesClientAPI = (*billing.ProfilesClient)(nil)

// CustomersClientAPI contains the set of methods on the CustomersClient type.
type CustomersClientAPI interface {
	Get(ctx context.Context, billingAccountName string, customerName string, expand string) (result billing.Customer, err error)
	ListByBillingAccount(ctx context.Context, billingAccountName string, search string, filter string) (result billing.CustomerListResultPage, err error)
	ListByBillingAccountComplete(ctx context.Context, billingAccountName string, search string, filter string) (result billing.CustomerListResultIterator, err error)
	ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, search string, filter string) (result billing.CustomerListResultPage, err error)
	ListByBillingProfileComplete(ctx context.Context, billingAccountName string, billingProfileName string, search string, filter string) (result billing.CustomerListResultIterator, err error)
}

var _ CustomersClientAPI = (*billing.CustomersClient)(nil)

// InvoiceSectionsClientAPI contains the set of methods on the InvoiceSectionsClient type.
type InvoiceSectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, parameters billing.InvoiceSection) (result billing.InvoiceSectionsCreateOrUpdateFuture, err error)
	Get(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result billing.InvoiceSection, err error)
	ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.InvoiceSectionListResultPage, err error)
	ListByBillingProfileComplete(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.InvoiceSectionListResultIterator, err error)
}

var _ InvoiceSectionsClientAPI = (*billing.InvoiceSectionsClient)(nil)

// PermissionsClientAPI contains the set of methods on the PermissionsClient type.
type PermissionsClientAPI interface {
	ListByBillingAccount(ctx context.Context, billingAccountName string) (result billing.PermissionsListResultPage, err error)
	ListByBillingAccountComplete(ctx context.Context, billingAccountName string) (result billing.PermissionsListResultIterator, err error)
	ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.PermissionsListResultPage, err error)
	ListByBillingProfileComplete(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.PermissionsListResultIterator, err error)
	ListByCustomer(ctx context.Context, billingAccountName string, customerName string) (result billing.PermissionsListResultPage, err error)
	ListByCustomerComplete(ctx context.Context, billingAccountName string, customerName string) (result billing.PermissionsListResultIterator, err error)
	ListByInvoiceSections(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result billing.PermissionsListResultPage, err error)
	ListByInvoiceSectionsComplete(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result billing.PermissionsListResultIterator, err error)
}

var _ PermissionsClientAPI = (*billing.PermissionsClient)(nil)

// SubscriptionsClientAPI contains the set of methods on the SubscriptionsClient type.
type SubscriptionsClientAPI interface {
	Get(ctx context.Context, billingAccountName string) (result billing.Subscription, err error)
	ListByBillingAccount(ctx context.Context, billingAccountName string) (result billing.SubscriptionsListResultPage, err error)
	ListByBillingAccountComplete(ctx context.Context, billingAccountName string) (result billing.SubscriptionsListResultIterator, err error)
	ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.SubscriptionsListResultPage, err error)
	ListByBillingProfileComplete(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.SubscriptionsListResultIterator, err error)
	ListByCustomer(ctx context.Context, billingAccountName string, customerName string) (result billing.SubscriptionsListResultPage, err error)
	ListByCustomerComplete(ctx context.Context, billingAccountName string, customerName string) (result billing.SubscriptionsListResultIterator, err error)
	ListByInvoiceSection(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result billing.SubscriptionsListResultPage, err error)
	ListByInvoiceSectionComplete(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result billing.SubscriptionsListResultIterator, err error)
	Move(ctx context.Context, billingAccountName string, parameters billing.TransferBillingSubscriptionRequestProperties) (result billing.SubscriptionsMoveFuture, err error)
	Update(ctx context.Context, billingAccountName string, parameters billing.Subscription) (result billing.Subscription, err error)
	ValidateMove(ctx context.Context, billingAccountName string, parameters billing.TransferBillingSubscriptionRequestProperties) (result billing.ValidateSubscriptionTransferEligibilityResult, err error)
}

var _ SubscriptionsClientAPI = (*billing.SubscriptionsClient)(nil)

// ProductsClientAPI contains the set of methods on the ProductsClient type.
type ProductsClientAPI interface {
	Get(ctx context.Context, billingAccountName string, productName string) (result billing.Product, err error)
	ListByBillingAccount(ctx context.Context, billingAccountName string, filter string) (result billing.ProductsListResultPage, err error)
	ListByBillingAccountComplete(ctx context.Context, billingAccountName string, filter string) (result billing.ProductsListResultIterator, err error)
	ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, filter string) (result billing.ProductsListResultPage, err error)
	ListByBillingProfileComplete(ctx context.Context, billingAccountName string, billingProfileName string, filter string) (result billing.ProductsListResultIterator, err error)
	ListByCustomer(ctx context.Context, billingAccountName string, customerName string) (result billing.ProductsListResultPage, err error)
	ListByCustomerComplete(ctx context.Context, billingAccountName string, customerName string) (result billing.ProductsListResultIterator, err error)
	ListByInvoiceSection(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, filter string) (result billing.ProductsListResultPage, err error)
	ListByInvoiceSectionComplete(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, filter string) (result billing.ProductsListResultIterator, err error)
	Move(ctx context.Context, billingAccountName string, productName string, parameters billing.TransferProductRequestProperties) (result billing.Product, err error)
	Update(ctx context.Context, billingAccountName string, productName string, parameters billing.Product) (result billing.Product, err error)
	ValidateMove(ctx context.Context, billingAccountName string, productName string, parameters billing.TransferProductRequestProperties) (result billing.ValidateProductTransferEligibilityResult, err error)
}

var _ ProductsClientAPI = (*billing.ProductsClient)(nil)

// InvoicesClientAPI contains the set of methods on the InvoicesClient type.
type InvoicesClientAPI interface {
	DownloadBillingSubscriptionInvoice(ctx context.Context, invoiceName string, downloadToken string) (result billing.InvoicesDownloadBillingSubscriptionInvoiceFuture, err error)
	DownloadInvoice(ctx context.Context, billingAccountName string, invoiceName string) (result billing.InvoicesDownloadInvoiceFuture, err error)
	Get(ctx context.Context, billingAccountName string, invoiceName string) (result billing.Invoice, err error)
	GetByID(ctx context.Context, invoiceName string) (result billing.Invoice, err error)
	GetBySubscriptionAndInvoiceID(ctx context.Context, invoiceName string) (result billing.Invoice, err error)
	ListByBillingAccount(ctx context.Context, billingAccountName string, periodStartDate string, periodEndDate string) (result billing.InvoiceListResultPage, err error)
	ListByBillingAccountComplete(ctx context.Context, billingAccountName string, periodStartDate string, periodEndDate string) (result billing.InvoiceListResultIterator, err error)
	ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, periodStartDate string, periodEndDate string) (result billing.InvoiceListResultPage, err error)
	ListByBillingProfileComplete(ctx context.Context, billingAccountName string, billingProfileName string, periodStartDate string, periodEndDate string) (result billing.InvoiceListResultIterator, err error)
	ListByBillingSubscription(ctx context.Context, periodStartDate string, periodEndDate string) (result billing.InvoiceListResultPage, err error)
	ListByBillingSubscriptionComplete(ctx context.Context, periodStartDate string, periodEndDate string) (result billing.InvoiceListResultIterator, err error)
}

var _ InvoicesClientAPI = (*billing.InvoicesClient)(nil)

// TransactionsClientAPI contains the set of methods on the TransactionsClient type.
type TransactionsClientAPI interface {
	ListByInvoice(ctx context.Context, billingAccountName string, invoiceName string) (result billing.TransactionListResultPage, err error)
	ListByInvoiceComplete(ctx context.Context, billingAccountName string, invoiceName string) (result billing.TransactionListResultIterator, err error)
}

var _ TransactionsClientAPI = (*billing.TransactionsClient)(nil)

// PoliciesClientAPI contains the set of methods on the PoliciesClient type.
type PoliciesClientAPI interface {
	GetByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.Policy, err error)
	GetByCustomer(ctx context.Context, billingAccountName string, customerName string) (result billing.CustomerPolicy, err error)
	Update(ctx context.Context, billingAccountName string, billingProfileName string, parameters billing.Policy) (result billing.Policy, err error)
	UpdateCustomer(ctx context.Context, billingAccountName string, customerName string, parameters billing.CustomerPolicy) (result billing.CustomerPolicy, err error)
}

var _ PoliciesClientAPI = (*billing.PoliciesClient)(nil)

// PropertyClientAPI contains the set of methods on the PropertyClient type.
type PropertyClientAPI interface {
	Get(ctx context.Context) (result billing.Property, err error)
	Update(ctx context.Context, parameters billing.Property) (result billing.Property, err error)
}

var _ PropertyClientAPI = (*billing.PropertyClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result billing.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result billing.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*billing.OperationsClient)(nil)

// RoleDefinitionsClientAPI contains the set of methods on the RoleDefinitionsClient type.
type RoleDefinitionsClientAPI interface {
	GetByBillingAccount(ctx context.Context, billingAccountName string, billingRoleDefinitionName string) (result billing.RoleDefinition, err error)
	GetByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleDefinitionName string) (result billing.RoleDefinition, err error)
	GetByInvoiceSection(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleDefinitionName string) (result billing.RoleDefinition, err error)
	ListByBillingAccount(ctx context.Context, billingAccountName string) (result billing.RoleDefinitionListResultPage, err error)
	ListByBillingAccountComplete(ctx context.Context, billingAccountName string) (result billing.RoleDefinitionListResultIterator, err error)
	ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.RoleDefinitionListResultPage, err error)
	ListByBillingProfileComplete(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.RoleDefinitionListResultIterator, err error)
	ListByInvoiceSection(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result billing.RoleDefinitionListResultPage, err error)
	ListByInvoiceSectionComplete(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result billing.RoleDefinitionListResultIterator, err error)
}

var _ RoleDefinitionsClientAPI = (*billing.RoleDefinitionsClient)(nil)

// RoleAssignmentsClientAPI contains the set of methods on the RoleAssignmentsClient type.
type RoleAssignmentsClientAPI interface {
	DeleteByBillingAccount(ctx context.Context, billingAccountName string, billingRoleAssignmentName string) (result billing.RoleAssignment, err error)
	DeleteByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleAssignmentName string) (result billing.RoleAssignment, err error)
	DeleteByInvoiceSection(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleAssignmentName string) (result billing.RoleAssignment, err error)
	GetByBillingAccount(ctx context.Context, billingAccountName string, billingRoleAssignmentName string) (result billing.RoleAssignment, err error)
	GetByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleAssignmentName string) (result billing.RoleAssignment, err error)
	GetByInvoiceSection(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleAssignmentName string) (result billing.RoleAssignment, err error)
	ListByBillingAccount(ctx context.Context, billingAccountName string) (result billing.RoleAssignmentListResultPage, err error)
	ListByBillingAccountComplete(ctx context.Context, billingAccountName string) (result billing.RoleAssignmentListResultIterator, err error)
	ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.RoleAssignmentListResultPage, err error)
	ListByBillingProfileComplete(ctx context.Context, billingAccountName string, billingProfileName string) (result billing.RoleAssignmentListResultIterator, err error)
	ListByInvoiceSection(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result billing.RoleAssignmentListResultPage, err error)
	ListByInvoiceSectionComplete(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result billing.RoleAssignmentListResultIterator, err error)
}

var _ RoleAssignmentsClientAPI = (*billing.RoleAssignmentsClient)(nil)

// AgreementsClientAPI contains the set of methods on the AgreementsClient type.
type AgreementsClientAPI interface {
	Get(ctx context.Context, billingAccountName string, agreementName string, expand string) (result billing.Agreement, err error)
	ListByBillingAccount(ctx context.Context, billingAccountName string, expand string) (result billing.AgreementListResultPage, err error)
	ListByBillingAccountComplete(ctx context.Context, billingAccountName string, expand string) (result billing.AgreementListResultIterator, err error)
}

var _ AgreementsClientAPI = (*billing.AgreementsClient)(nil)
