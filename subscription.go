// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package checkbook

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Munchpass/checkbook/internal/apijson"
	"github.com/Munchpass/checkbook/internal/apiquery"
	"github.com/Munchpass/checkbook/internal/requestconfig"
	"github.com/Munchpass/checkbook/option"
	"github.com/Munchpass/checkbook/packages/param"
	"github.com/Munchpass/checkbook/packages/respjson"
)

// SubscriptionService contains methods and other services that help with
// interacting with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionService] method instead.
type SubscriptionService struct {
	Options []option.RequestOption
}

// NewSubscriptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscriptionService(opts ...option.RequestOption) (r SubscriptionService) {
	r = SubscriptionService{}
	r.Options = opts
	return
}

// Get the specified subscription
func (r *SubscriptionService) Get(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (res *GetSubscriptionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("v3/subscription/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the specified subscription
func (r *SubscriptionService) Update(ctx context.Context, subscriptionID string, body SubscriptionUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("v3/subscription/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Return the subscriptions
func (r *SubscriptionService) List(ctx context.Context, query SubscriptionListParams, opts ...option.RequestOption) (res *SubscriptionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/subscription"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Remove the specified subscription
func (r *SubscriptionService) Delete(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("v3/subscription/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Create a new invoice subscription
func (r *SubscriptionService) NewInvoice(ctx context.Context, body SubscriptionNewInvoiceParams, opts ...option.RequestOption) (res *CreateSubscriptionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/subscription/invoice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new payment subscription
func (r *SubscriptionService) NewPayment(ctx context.Context, body SubscriptionNewPaymentParams, opts ...option.RequestOption) (res *CreateSubscriptionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/subscription/check"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CreateSubscriptionResponse struct {
	// Unique identifier for subscription
	ID string `json:"id,required"`
	// Subscription amount
	Amount float64 `json:"amount,required"`
	// Subscription creation timestamp
	Date string `json:"date,required"`
	// Subscription description
	Description string `json:"description,required"`
	// How often the subscription will recur
	//
	// Any of "WEEKLY", "MONTHLY".
	Interval CreateSubscriptionResponseInterval `json:"interval,required"`
	// Subscription start date
	StartDate string `json:"start_date,required"`
	// Type of the subscription
	//
	// Any of "INVOICE", "CHECK".
	Type CreateSubscriptionResponseType `json:"type,required"`
	// Debit/credit account id for funds
	Account string `json:"account,nullable"`
	// Number of times the subscription will recur (null indicates indefinite)
	Duration int64 `json:"duration,nullable"`
	// Name of third party who is receiving the check/invoice
	Name      string                                   `json:"name,nullable"`
	Recipient CreateSubscriptionResponseRecipientUnion `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		Date        respjson.Field
		Description respjson.Field
		Interval    respjson.Field
		StartDate   respjson.Field
		Type        respjson.Field
		Account     respjson.Field
		Duration    respjson.Field
		Name        respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateSubscriptionResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateSubscriptionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How often the subscription will recur
type CreateSubscriptionResponseInterval string

const (
	CreateSubscriptionResponseIntervalWeekly  CreateSubscriptionResponseInterval = "WEEKLY"
	CreateSubscriptionResponseIntervalMonthly CreateSubscriptionResponseInterval = "MONTHLY"
)

// Type of the subscription
type CreateSubscriptionResponseType string

const (
	CreateSubscriptionResponseTypeInvoice CreateSubscriptionResponseType = "INVOICE"
	CreateSubscriptionResponseTypeCheck   CreateSubscriptionResponseType = "CHECK"
)

// CreateSubscriptionResponseRecipientUnion contains all possible properties and
// values from [string], [SubscriptionAddress].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type CreateSubscriptionResponseRecipientUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [SubscriptionAddress].
	City string `json:"city"`
	// This field is from variant [SubscriptionAddress].
	Line1 string `json:"line_1"`
	// This field is from variant [SubscriptionAddress].
	State string `json:"state"`
	// This field is from variant [SubscriptionAddress].
	Zip string `json:"zip"`
	// This field is from variant [SubscriptionAddress].
	Country string `json:"country"`
	// This field is from variant [SubscriptionAddress].
	Line2 string `json:"line_2"`
	JSON  struct {
		OfString respjson.Field
		City     respjson.Field
		Line1    respjson.Field
		State    respjson.Field
		Zip      respjson.Field
		Country  respjson.Field
		Line2    respjson.Field
		raw      string
	} `json:"-"`
}

func (u CreateSubscriptionResponseRecipientUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CreateSubscriptionResponseRecipientUnion) AsSubscriptionAddress() (v SubscriptionAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CreateSubscriptionResponseRecipientUnion) RawJSON() string { return u.JSON.raw }

func (r *CreateSubscriptionResponseRecipientUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetSubscriptionResponse struct {
	// Unique identifier for subscription
	ID string `json:"id,required"`
	// Subscription amount
	Amount float64 `json:"amount,required"`
	// Subscription creation timestamp
	Date string `json:"date,required"`
	// Subscription description
	Description string `json:"description,required"`
	// How often the subscription will recur
	//
	// Any of "WEEKLY", "MONTHLY".
	Interval GetSubscriptionResponseInterval `json:"interval,required"`
	// List of skipped subscription indexes (e.g. [1] indicates the first subscription
	// will be skipped)
	Skipped []int64 `json:"skipped,required"`
	// Subscription start date
	StartDate string `json:"start_date,required"`
	// Current status of the check or invoice
	//
	// Any of "PAID", "IN_PROCESS", "UNPAID", "VOID", "EXPIRED", "PRINTED", "MAILED",
	// "FAILED", "REFUNDED", "CANCELED", "OVERDUE".
	Status GetSubscriptionResponseStatus `json:"status,required"`
	// Type of the subscription
	//
	// Any of "INVOICE", "CHECK".
	Type GetSubscriptionResponseType `json:"type,required"`
	// Debit/credit account id for funds
	Account string `json:"account,nullable"`
	// Autopay invoice
	Autopay bool `json:"autopay"`
	// Number of times the subscription will recur (null indicates indefinite)
	Duration int64 `json:"duration,nullable"`
	// Name of third party who is receiving the check/invoice
	Name      string                                `json:"name,nullable"`
	Recipient GetSubscriptionResponseRecipientUnion `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		Date        respjson.Field
		Description respjson.Field
		Interval    respjson.Field
		Skipped     respjson.Field
		StartDate   respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Account     respjson.Field
		Autopay     respjson.Field
		Duration    respjson.Field
		Name        respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetSubscriptionResponse) RawJSON() string { return r.JSON.raw }
func (r *GetSubscriptionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How often the subscription will recur
type GetSubscriptionResponseInterval string

const (
	GetSubscriptionResponseIntervalWeekly  GetSubscriptionResponseInterval = "WEEKLY"
	GetSubscriptionResponseIntervalMonthly GetSubscriptionResponseInterval = "MONTHLY"
)

// Current status of the check or invoice
type GetSubscriptionResponseStatus string

const (
	GetSubscriptionResponseStatusPaid      GetSubscriptionResponseStatus = "PAID"
	GetSubscriptionResponseStatusInProcess GetSubscriptionResponseStatus = "IN_PROCESS"
	GetSubscriptionResponseStatusUnpaid    GetSubscriptionResponseStatus = "UNPAID"
	GetSubscriptionResponseStatusVoid      GetSubscriptionResponseStatus = "VOID"
	GetSubscriptionResponseStatusExpired   GetSubscriptionResponseStatus = "EXPIRED"
	GetSubscriptionResponseStatusPrinted   GetSubscriptionResponseStatus = "PRINTED"
	GetSubscriptionResponseStatusMailed    GetSubscriptionResponseStatus = "MAILED"
	GetSubscriptionResponseStatusFailed    GetSubscriptionResponseStatus = "FAILED"
	GetSubscriptionResponseStatusRefunded  GetSubscriptionResponseStatus = "REFUNDED"
	GetSubscriptionResponseStatusCanceled  GetSubscriptionResponseStatus = "CANCELED"
	GetSubscriptionResponseStatusOverdue   GetSubscriptionResponseStatus = "OVERDUE"
)

// Type of the subscription
type GetSubscriptionResponseType string

const (
	GetSubscriptionResponseTypeInvoice GetSubscriptionResponseType = "INVOICE"
	GetSubscriptionResponseTypeCheck   GetSubscriptionResponseType = "CHECK"
)

// GetSubscriptionResponseRecipientUnion contains all possible properties and
// values from [string], [SubscriptionAddress].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type GetSubscriptionResponseRecipientUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [SubscriptionAddress].
	City string `json:"city"`
	// This field is from variant [SubscriptionAddress].
	Line1 string `json:"line_1"`
	// This field is from variant [SubscriptionAddress].
	State string `json:"state"`
	// This field is from variant [SubscriptionAddress].
	Zip string `json:"zip"`
	// This field is from variant [SubscriptionAddress].
	Country string `json:"country"`
	// This field is from variant [SubscriptionAddress].
	Line2 string `json:"line_2"`
	JSON  struct {
		OfString respjson.Field
		City     respjson.Field
		Line1    respjson.Field
		State    respjson.Field
		Zip      respjson.Field
		Country  respjson.Field
		Line2    respjson.Field
		raw      string
	} `json:"-"`
}

func (u GetSubscriptionResponseRecipientUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GetSubscriptionResponseRecipientUnion) AsSubscriptionAddress() (v SubscriptionAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GetSubscriptionResponseRecipientUnion) RawJSON() string { return u.JSON.raw }

func (r *GetSubscriptionResponseRecipientUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionAddress struct {
	// City
	City string `json:"city,required"`
	// Line 1
	Line1 string `json:"line_1,required"`
	// State
	State string `json:"state,required"`
	// Zip/postal code
	Zip string `json:"zip,required"`
	// Country
	Country string `json:"country"`
	// Line 2
	Line2 string `json:"line_2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Line1       respjson.Field
		State       respjson.Field
		Zip         respjson.Field
		Country     respjson.Field
		Line2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionAddress) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SubscriptionAddress to a SubscriptionAddressParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SubscriptionAddressParam.Overrides()
func (r SubscriptionAddress) ToParam() SubscriptionAddressParam {
	return param.Override[SubscriptionAddressParam](json.RawMessage(r.RawJSON()))
}

// The properties City, Line1, State, Zip are required.
type SubscriptionAddressParam struct {
	// City
	City string `json:"city,required"`
	// Line 1
	Line1 string `json:"line_1,required"`
	// State
	State string `json:"state,required"`
	// Zip/postal code
	Zip string `json:"zip,required"`
	// Country
	Country param.Opt[string] `json:"country,omitzero"`
	// Line 2
	Line2 param.Opt[string] `json:"line_2,omitzero"`
	paramObj
}

func (r SubscriptionAddressParam) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionAddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionAddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionListResponse struct {
	Page  int64 `json:"page,required"`
	Pages int64 `json:"pages,required"`
	// List of subscriptions
	Subscriptions []GetSubscriptionResponse `json:"subscriptions,required"`
	Total         int64                     `json:"total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page          respjson.Field
		Pages         respjson.Field
		Subscriptions respjson.Field
		Total         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionListResponse) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionUpdateParams struct {
	// Autopay invoice
	Autopay param.Opt[bool] `json:"autopay,omitzero"`
	// List of skipped subscription indexes (e.g. [1] indicates the first subscription
	// will be skipped)
	Skipped []int64 `json:"skipped,omitzero"`
	paramObj
}

func (r SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionListParams struct {
	// End date
	EndDate param.Opt[time.Time] `query:"end_date,omitzero" format:"date" json:"-"`
	// Start date
	StartDate param.Opt[time.Time] `query:"start_date,omitzero" format:"date" json:"-"`
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Query
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Direction
	//
	// Any of "INCOMING", "OUTGOING".
	Direction SubscriptionListParamsDirection `query:"direction,omitzero" json:"-"`
	// Items per page
	//
	// Any of 10, 25, 50, 100, 250.
	PerPage int64 `query:"per_page,omitzero" json:"-"`
	// Sort
	//
	// Any of "+NUMBER", "-NUMBER", "+TYPE", "-TYPE", "+AMOUNT", "-AMOUNT", "+STATUS",
	// "-STATUS", "+DATE", "-DATE", "+UPDATE", "-UPDATE", "+DESCRIPTION",
	// "-DESCRIPTION".
	Sort SubscriptionListParamsSort `query:"sort,omitzero" json:"-"`
	// Status
	//
	// Any of "PAID", "IN_PROCESS", "UNPAID", "VOID", "EXPIRED", "PRINTED", "MAILED",
	// "FAILED", "REFUNDED".
	Status SubscriptionListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SubscriptionListParams]'s query parameters as `url.Values`.
func (r SubscriptionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction
type SubscriptionListParamsDirection string

const (
	SubscriptionListParamsDirectionIncoming SubscriptionListParamsDirection = "INCOMING"
	SubscriptionListParamsDirectionOutgoing SubscriptionListParamsDirection = "OUTGOING"
)

// Sort
type SubscriptionListParamsSort string

const (
	SubscriptionListParamsSortPlusNumber       SubscriptionListParamsSort = "+NUMBER"
	SubscriptionListParamsSortMinusNumber      SubscriptionListParamsSort = "-NUMBER"
	SubscriptionListParamsSortPlusType         SubscriptionListParamsSort = "+TYPE"
	SubscriptionListParamsSortMinusType        SubscriptionListParamsSort = "-TYPE"
	SubscriptionListParamsSortPlusAmount       SubscriptionListParamsSort = "+AMOUNT"
	SubscriptionListParamsSortMinusAmount      SubscriptionListParamsSort = "-AMOUNT"
	SubscriptionListParamsSortPlusStatus       SubscriptionListParamsSort = "+STATUS"
	SubscriptionListParamsSortMinusStatus      SubscriptionListParamsSort = "-STATUS"
	SubscriptionListParamsSortPlusDate         SubscriptionListParamsSort = "+DATE"
	SubscriptionListParamsSortMinusDate        SubscriptionListParamsSort = "-DATE"
	SubscriptionListParamsSortPlusUpdate       SubscriptionListParamsSort = "+UPDATE"
	SubscriptionListParamsSortMinusUpdate      SubscriptionListParamsSort = "-UPDATE"
	SubscriptionListParamsSortPlusDescription  SubscriptionListParamsSort = "+DESCRIPTION"
	SubscriptionListParamsSortMinusDescription SubscriptionListParamsSort = "-DESCRIPTION"
)

// Status
type SubscriptionListParamsStatus string

const (
	SubscriptionListParamsStatusPaid      SubscriptionListParamsStatus = "PAID"
	SubscriptionListParamsStatusInProcess SubscriptionListParamsStatus = "IN_PROCESS"
	SubscriptionListParamsStatusUnpaid    SubscriptionListParamsStatus = "UNPAID"
	SubscriptionListParamsStatusVoid      SubscriptionListParamsStatus = "VOID"
	SubscriptionListParamsStatusExpired   SubscriptionListParamsStatus = "EXPIRED"
	SubscriptionListParamsStatusPrinted   SubscriptionListParamsStatus = "PRINTED"
	SubscriptionListParamsStatusMailed    SubscriptionListParamsStatus = "MAILED"
	SubscriptionListParamsStatusFailed    SubscriptionListParamsStatus = "FAILED"
	SubscriptionListParamsStatusRefunded  SubscriptionListParamsStatus = "REFUNDED"
)

type SubscriptionNewInvoiceParams struct {
	// Invoice amount
	Amount float64 `json:"amount,required"`
	// Description for invoice
	Description string `json:"description,required"`
	// How often the subscription will recur
	//
	// Any of "WEEKLY", "MONTHLY".
	Interval SubscriptionNewInvoiceParamsInterval `json:"interval,omitzero,required"`
	// Name of recipient
	Name string `json:"name,required"`
	// Email/id of recipient
	Recipient string `json:"recipient,required"`
	// Optional number of times the subscription should recur (defaults to indefinite)
	Duration param.Opt[int64] `json:"duration,omitzero"`
	// Start date for subscription (this is the first date the subsription will be sent
	// out and defaults to the current timestamp)
	StartDate param.Opt[time.Time] `json:"start_date,omitzero" format:"date"`
	// Debit account id for funds (if sender has multiple bank accounts)
	Account param.Opt[string] `json:"account,omitzero" format:"uuid"`
	paramObj
}

func (r SubscriptionNewInvoiceParams) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionNewInvoiceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionNewInvoiceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How often the subscription will recur
type SubscriptionNewInvoiceParamsInterval string

const (
	SubscriptionNewInvoiceParamsIntervalWeekly  SubscriptionNewInvoiceParamsInterval = "WEEKLY"
	SubscriptionNewInvoiceParamsIntervalMonthly SubscriptionNewInvoiceParamsInterval = "MONTHLY"
)

type SubscriptionNewPaymentParams struct {
	// Invoice amount
	Amount float64 `json:"amount,required"`
	// How often the subscription will recur
	//
	// Any of "WEEKLY", "MONTHLY".
	Interval SubscriptionNewPaymentParamsInterval `json:"interval,omitzero,required"`
	// Name of recipient
	Name string `json:"name,required"`
	// Email/id of recipient
	Recipient SubscriptionNewPaymentParamsRecipientUnion `json:"recipient,omitzero,required"`
	// Description for check
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional number of times the subscription should recur (defaults to indefinite)
	Duration param.Opt[int64] `json:"duration,omitzero"`
	// Start date for subscription (this is the first date the subsription will be sent
	// out and defaults to the current timestamp)
	StartDate param.Opt[time.Time] `json:"start_date,omitzero" format:"date"`
	// Debit account id for funds (if sender has multiple bank accounts)
	Account param.Opt[string] `json:"account,omitzero" format:"uuid"`
	// List of the remittance records
	RemittanceAdvice []RemittanceAdviceParam `json:"remittance_advice,omitzero"`
	paramObj
}

func (r SubscriptionNewPaymentParams) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionNewPaymentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionNewPaymentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How often the subscription will recur
type SubscriptionNewPaymentParamsInterval string

const (
	SubscriptionNewPaymentParamsIntervalWeekly  SubscriptionNewPaymentParamsInterval = "WEEKLY"
	SubscriptionNewPaymentParamsIntervalMonthly SubscriptionNewPaymentParamsInterval = "MONTHLY"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SubscriptionNewPaymentParamsRecipientUnion struct {
	OfString              param.Opt[string]         `json:",omitzero,inline"`
	OfSubscriptionAddress *SubscriptionAddressParam `json:",omitzero,inline"`
	paramUnion
}

func (u SubscriptionNewPaymentParamsRecipientUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfSubscriptionAddress)
}
func (u *SubscriptionNewPaymentParamsRecipientUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SubscriptionNewPaymentParamsRecipientUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfSubscriptionAddress) {
		return u.OfSubscriptionAddress
	}
	return nil
}
