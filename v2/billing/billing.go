// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package billing implements the DocuSign SDK
// category Billing.
//
// Use the Billing category to manage the following billing related tasks:
//
// * Retrieve and update billing plan information.
// * Retrieve invoices.
// * Retrieve and update payment information.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/esign-rest-api/v2/reference/Billing
// Usage example:
//
//   import (
//       "github.com/jfcote87/esign"
//       "github.com/jfcote87/esign/v2/billing"
//       "github.com/jfcote87/esign/v2/model"
//   )
//   ...
//   billingService := billing.New(esignCredential)
package billing // import "github.com/jfcote87/esign/v2/billing"

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/jfcote87/esign"
	"github.com/jfcote87/esign/v2/model"
)

// Service implements DocuSign Billing Category API operations
type Service struct {
	credential esign.Credential
}

// New initializes a billing service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// PlansGet get the billing plan details.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/billing/billingplans/get
//
// SDK Method Billing::getBillingPlan
func (s *Service) PlansGet(billingPlanID string) *PlansGetOp {
	return &PlansGetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"", "v2", "billing_plans", billingPlanID}, "/"),
		QueryOpts:  make(url.Values),
	}
}

// PlansGetOp implements DocuSign API SDK Billing::getBillingPlan
type PlansGetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *PlansGetOp) Do(ctx context.Context) (*model.BillingPlanResponse, error) {
	var res *model.BillingPlanResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// PlansGetAccountPlan get Account Billing Plan
//
// https://developers.docusign.com/esign-rest-api/v2/reference/billing/billingplans/getaccountplan
//
// SDK Method Billing::getPlan
func (s *Service) PlansGetAccountPlan() *PlansGetAccountPlanOp {
	return &PlansGetAccountPlanOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "billing_plan",
		QueryOpts:  make(url.Values),
	}
}

// PlansGetAccountPlanOp implements DocuSign API SDK Billing::getPlan
type PlansGetAccountPlanOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *PlansGetAccountPlanOp) Do(ctx context.Context) (*model.AccountBillingPlanResponse, error) {
	var res *model.AccountBillingPlanResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// IncludeCreditCardInformation when set to **true**, payment information including credit card information will show in the return.
func (op *PlansGetAccountPlanOp) IncludeCreditCardInformation() *PlansGetAccountPlanOp {
	if op != nil {
		op.QueryOpts.Set("include_credit_card_information", "true")
	}
	return op
}

// IncludeMetadata when set to **true**, the `canUpgrade` and `renewalStatus` properities are included the response and an array of `supportedCountries` property is added to the `billingAddress` information.
func (op *PlansGetAccountPlanOp) IncludeMetadata() *PlansGetAccountPlanOp {
	if op != nil {
		op.QueryOpts.Set("include_metadata", "true")
	}
	return op
}

// IncludeSuccessorPlans when set to **true**, excludes successor information from the response.
func (op *PlansGetAccountPlanOp) IncludeSuccessorPlans() *PlansGetAccountPlanOp {
	if op != nil {
		op.QueryOpts.Set("include_successor_plans", "true")
	}
	return op
}

// PlansGetCreditCard get metadata for a given credit card.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/billing/billingplans/getcreditcard
//
// SDK Method Billing::getCreditCardInfo
func (s *Service) PlansGetCreditCard() *PlansGetCreditCardOp {
	return &PlansGetCreditCardOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "billing_plan/credit_card",
		QueryOpts:  make(url.Values),
	}
}

// PlansGetCreditCardOp implements DocuSign API SDK Billing::getCreditCardInfo
type PlansGetCreditCardOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *PlansGetCreditCardOp) Do(ctx context.Context) (*model.CreditCardInformation, error) {
	var res *model.CreditCardInformation
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// PlansList gets the list of available billing plans.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/billing/billingplans/list
//
// SDK Method Billing::listBillingPlans
func (s *Service) PlansList() *PlansListOp {
	return &PlansListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "/v2/billing_plans",
		QueryOpts:  make(url.Values),
	}
}

// PlansListOp implements DocuSign API SDK Billing::listBillingPlans
type PlansListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *PlansListOp) Do(ctx context.Context) (*model.BillingPlansResponse, error) {
	var res *model.BillingPlansResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// PlansPurchaseEnvelopes reserverd: Purchase additional envelopes.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/billing/billingplans/purchaseenvelopes
//
// SDK Method Billing::purchaseEnvelopes
func (s *Service) PlansPurchaseEnvelopes(purchasedEnvelopesInformation *model.PurchasedEnvelopesInformation) *PlansPurchaseEnvelopesOp {
	return &PlansPurchaseEnvelopesOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       "billing_plan/purchased_envelopes",
		Payload:    purchasedEnvelopesInformation,
		QueryOpts:  make(url.Values),
	}
}

// PlansPurchaseEnvelopesOp implements DocuSign API SDK Billing::purchaseEnvelopes
type PlansPurchaseEnvelopesOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *PlansPurchaseEnvelopesOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// PlansUpdate updates the account billing plan.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/billing/billingplans/update
//
// SDK Method Billing::updatePlan
func (s *Service) PlansUpdate(billingPlanInformation *model.BillingPlanInformation) *PlansUpdateOp {
	return &PlansUpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       "billing_plan",
		Payload:    billingPlanInformation,
		QueryOpts:  make(url.Values),
	}
}

// PlansUpdateOp implements DocuSign API SDK Billing::updatePlan
type PlansUpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *PlansUpdateOp) Do(ctx context.Context) (*model.BillingPlanUpdateResponse, error) {
	var res *model.BillingPlanUpdateResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// PreviewBillingPlan when set to **true**, updates the account using a preview billing plan.
func (op *PlansUpdateOp) PreviewBillingPlan() *PlansUpdateOp {
	if op != nil {
		op.QueryOpts.Set("preview_billing_plan", "true")
	}
	return op
}

// InvoicesGet retrieves a billing invoice.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/billing/invoices/get
//
// SDK Method Billing::getInvoice
func (s *Service) InvoicesGet(invoiceID string) *InvoicesGetOp {
	return &InvoicesGetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"billing_invoices", invoiceID}, "/"),
		QueryOpts:  make(url.Values),
	}
}

// InvoicesGetOp implements DocuSign API SDK Billing::getInvoice
type InvoicesGetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *InvoicesGetOp) Do(ctx context.Context) (*model.BillingInvoice, error) {
	var res *model.BillingInvoice
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// PDF returns a pdf version of the invoice by setting
// the Accept header to application/pdf
//
// **not included in swagger definition
func (op *InvoicesGetOp) PDF(ctx context.Context) (*esign.Download, error) {
	var res *esign.Download
	if op == nil {
		return nil, esign.ErrNilOp
	}
	newOp := esign.Op(*op)
	newOp.Accept = "application/pdf"
	return res, (&newOp).Do(ctx, &res)
}

// InvoicesList get a List of Billing Invoices
//
// https://developers.docusign.com/esign-rest-api/v2/reference/billing/invoices/list
//
// SDK Method Billing::listInvoices
func (s *Service) InvoicesList() *InvoicesListOp {
	return &InvoicesListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "billing_invoices",
		QueryOpts:  make(url.Values),
	}
}

// InvoicesListOp implements DocuSign API SDK Billing::listInvoices
type InvoicesListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *InvoicesListOp) Do(ctx context.Context) (*model.BillingInvoicesResponse, error) {
	var res *model.BillingInvoicesResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// FromDate specifies the date/time of the earliest invoice in the account to retrieve.
func (op *InvoicesListOp) FromDate(val time.Time) *InvoicesListOp {
	if op != nil {
		op.QueryOpts.Set("from_date", val.Format(time.RFC3339))
	}
	return op
}

// ToDate specifies the date/time of the latest invoice in the account to retrieve.
func (op *InvoicesListOp) ToDate(val time.Time) *InvoicesListOp {
	if op != nil {
		op.QueryOpts.Set("to_date", val.Format(time.RFC3339))
	}
	return op
}

// InvoicesListPastDue get a list of past due invoices.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/billing/invoices/listpastdue
//
// SDK Method Billing::listInvoicesPastDue
func (s *Service) InvoicesListPastDue() *InvoicesListPastDueOp {
	return &InvoicesListPastDueOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "billing_invoices_past_due",
		QueryOpts:  make(url.Values),
	}
}

// InvoicesListPastDueOp implements DocuSign API SDK Billing::listInvoicesPastDue
type InvoicesListPastDueOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *InvoicesListPastDueOp) Do(ctx context.Context) (*model.BillingInvoicesSummary, error) {
	var res *model.BillingInvoicesSummary
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// PaymentsCreate posts a payment to a past due invoice.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/billing/payments/create
//
// SDK Method Billing::makePayment
func (s *Service) PaymentsCreate(billingPaymentRequest *model.BillingPaymentRequest) *PaymentsCreateOp {
	return &PaymentsCreateOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       "billing_payments",
		Payload:    billingPaymentRequest,
		QueryOpts:  make(url.Values),
	}
}

// PaymentsCreateOp implements DocuSign API SDK Billing::makePayment
type PaymentsCreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *PaymentsCreateOp) Do(ctx context.Context) (*model.BillingPaymentResponse, error) {
	var res *model.BillingPaymentResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// PaymentsGet gets billing payment information for a specific payment.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/billing/payments/get
//
// SDK Method Billing::getPayment
func (s *Service) PaymentsGet(paymentID string) *PaymentsGetOp {
	return &PaymentsGetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"billing_payments", paymentID}, "/"),
		QueryOpts:  make(url.Values),
	}
}

// PaymentsGetOp implements DocuSign API SDK Billing::getPayment
type PaymentsGetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *PaymentsGetOp) Do(ctx context.Context) (*model.BillingPaymentItem, error) {
	var res *model.BillingPaymentItem
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// PaymentsList gets payment information for one or more payments.
//
// https://developers.docusign.com/esign-rest-api/v2/reference/billing/payments/list
//
// SDK Method Billing::listPayments
func (s *Service) PaymentsList() *PaymentsListOp {
	return &PaymentsListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "billing_payments",
		QueryOpts:  make(url.Values),
	}
}

// PaymentsListOp implements DocuSign API SDK Billing::listPayments
type PaymentsListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *PaymentsListOp) Do(ctx context.Context) (*model.BillingPaymentsResponse, error) {
	var res *model.BillingPaymentsResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// FromDate specifies the date/time of the earliest payment in the account to retrieve.
func (op *PaymentsListOp) FromDate(val time.Time) *PaymentsListOp {
	if op != nil {
		op.QueryOpts.Set("from_date", val.Format(time.RFC3339))
	}
	return op
}

// ToDate specifies the date/time of the latest payment in the account to retrieve.
func (op *PaymentsListOp) ToDate(val time.Time) *PaymentsListOp {
	if op != nil {
		op.QueryOpts.Set("to_date", val.Format(time.RFC3339))
	}
	return op
}
