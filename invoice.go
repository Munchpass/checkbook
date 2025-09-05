// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package checkbook

import (
	"context"
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

// InvoiceService contains methods and other services that help with interacting
// with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvoiceService] method instead.
type InvoiceService struct {
	Options []option.RequestOption
}

// NewInvoiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInvoiceService(opts ...option.RequestOption) (r InvoiceService) {
	r = InvoiceService{}
	r.Options = opts
	return
}

// Create a new invoice
func (r *InvoiceService) New(ctx context.Context, body InvoiceNewParams, opts ...option.RequestOption) (res *InvoiceNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/invoice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the specified invoice
func (r *InvoiceService) Get(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *GetInvoice, err error) {
	opts = append(r.Options[:], opts...)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return
	}
	path := fmt.Sprintf("v3/invoice/%s", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get sent/received invoices
func (r *InvoiceService) List(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) (res *InvoiceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/invoice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get the attachment for an invoice
func (r *InvoiceService) GetAttachment(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *Error, err error) {
	opts = append(r.Options[:], opts...)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return
	}
	path := fmt.Sprintf("v3/invoice/%s/attachment", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Pay an outstanding invoice
func (r *InvoiceService) Pay(ctx context.Context, body InvoicePayParams, opts ...option.RequestOption) (res *InvoicePayResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/invoice/payment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Cancel the specified invoice
func (r *InvoiceService) Void(ctx context.Context, invoiceID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return
	}
	path := fmt.Sprintf("v3/invoice/%s", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type GetInvoice struct {
	// Unique identifier for invoice
	ID string `json:"id,required"`
	// Invoice amount
	Amount float64 `json:"amount,required"`
	// Invoice creation timestamp
	Date string `json:"date,required"`
	// Invoice description
	Description string `json:"description,required"`
	// Direction
	//
	// Any of "INCOMING", "OUTGOING".
	Direction GetInvoiceDirection `json:"direction,required"`
	// Invoice number
	Number string `json:"number,required"`
	// Current status of the invoice
	Status string `json:"status,required"`
	// URI for invoice attachment
	AttachmentUri string `json:"attachment_uri,nullable"`
	// Associated check number
	CheckID string `json:"check_id,nullable"`
	// Name of third party who sent/received the invoice
	Name string `json:"name,nullable"`
	// Email/id of the invoice recipient
	Recipient string `json:"recipient,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Amount        respjson.Field
		Date          respjson.Field
		Description   respjson.Field
		Direction     respjson.Field
		Number        respjson.Field
		Status        respjson.Field
		AttachmentUri respjson.Field
		CheckID       respjson.Field
		Name          respjson.Field
		Recipient     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetInvoice) RawJSON() string { return r.JSON.raw }
func (r *GetInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Direction
type GetInvoiceDirection string

const (
	GetInvoiceDirectionIncoming GetInvoiceDirection = "INCOMING"
	GetInvoiceDirectionOutgoing GetInvoiceDirection = "OUTGOING"
)

type InvoiceNewResponse struct {
	// Unique identifier for invoice
	ID string `json:"id,required"`
	// Invoice amount
	Amount float64 `json:"amount,required"`
	// Invoice creation timestamp
	Date string `json:"date,required"`
	// Invoice description
	Description string `json:"description,required"`
	// Name of third party who sent/received the invoice
	Name string `json:"name,required"`
	// Invoice number
	Number string `json:"number,required"`
	// Email/id of the invoice recipient
	Recipient string `json:"recipient,required"`
	// Current status of the invoice
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		Date        respjson.Field
		Description respjson.Field
		Name        respjson.Field
		Number      respjson.Field
		Recipient   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceNewResponse) RawJSON() string { return r.JSON.raw }
func (r *InvoiceNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceListResponse struct {
	// List of sent/received invoices
	Invoices []GetInvoice `json:"invoices,required"`
	Page     int64        `json:"page,required"`
	Pages    int64        `json:"pages,required"`
	Total    int64        `json:"total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Invoices    respjson.Field
		Page        respjson.Field
		Pages       respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceListResponse) RawJSON() string { return r.JSON.raw }
func (r *InvoiceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoicePayResponse struct {
	// Unique identifier for invoice
	ID string `json:"id,required"`
	// Invoice amount
	Amount float64 `json:"amount,required"`
	// Invoice creation timestamp
	Date string `json:"date,required"`
	// Name of third party who sent/received the invoice
	Name string `json:"name,required"`
	// Current status of the invoice
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		Date        respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoicePayResponse) RawJSON() string { return r.JSON.raw }
func (r *InvoicePayResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceNewParams struct {
	// Invoice amount
	Amount float64 `json:"amount,required"`
	// Description/memo for invoice
	Description string `json:"description,required"`
	// Name of recipient
	Name string `json:"name,required"`
	// Email/id of recipient
	Recipient string            `json:"recipient,required"`
	Number    param.Opt[string] `json:"number,omitzero"`
	// Optional credit account id for funds (if sender has multiple bank accounts)
	Account    param.Opt[string]               `json:"account,omitzero" format:"uuid"`
	Attachment InvoiceNewParamsAttachmentUnion `json:"attachment,omitzero"`
	paramObj
}

func (r InvoiceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InvoiceNewParamsAttachmentUnion struct {
	OfInvoiceAttachment *InvoiceNewParamsAttachmentInvoiceAttachment `json:",omitzero,inline"`
	OfString            param.Opt[string]                            `json:",omitzero,inline"`
	paramUnion
}

func (u InvoiceNewParamsAttachmentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInvoiceAttachment, u.OfString)
}
func (u *InvoiceNewParamsAttachmentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InvoiceNewParamsAttachmentUnion) asAny() any {
	if !param.IsOmitted(u.OfInvoiceAttachment) {
		return u.OfInvoiceAttachment
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// The properties Content, Filename are required.
type InvoiceNewParamsAttachmentInvoiceAttachment struct {
	// Attachment content
	Content string `json:"content,required"`
	// Attachment filename
	Filename string `json:"filename,required"`
	paramObj
}

func (r InvoiceNewParamsAttachmentInvoiceAttachment) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceNewParamsAttachmentInvoiceAttachment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceNewParamsAttachmentInvoiceAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceListParams struct {
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
	Direction InvoiceListParamsDirection `query:"direction,omitzero" json:"-"`
	// Items per page
	//
	// Any of 10, 25, 50, 100, 250.
	PerPage int64 `query:"per_page,omitzero" json:"-"`
	// Sort
	//
	// Any of "+NUMBER", "-NUMBER", "+TYPE", "-TYPE", "+AMOUNT", "-AMOUNT", "+STATUS",
	// "-STATUS", "+DATE", "-DATE", "+UPDATE", "-UPDATE", "+DESCRIPTION",
	// "-DESCRIPTION".
	Sort InvoiceListParamsSort `query:"sort,omitzero" json:"-"`
	// Status
	//
	// Any of "PAID", "IN_PROCESS", "UNPAID", "VOID", "EXPIRED", "PRINTED", "MAILED",
	// "FAILED", "REFUNDED".
	Status InvoiceListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InvoiceListParams]'s query parameters as `url.Values`.
func (r InvoiceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction
type InvoiceListParamsDirection string

const (
	InvoiceListParamsDirectionIncoming InvoiceListParamsDirection = "INCOMING"
	InvoiceListParamsDirectionOutgoing InvoiceListParamsDirection = "OUTGOING"
)

// Sort
type InvoiceListParamsSort string

const (
	InvoiceListParamsSortPlusNumber       InvoiceListParamsSort = "+NUMBER"
	InvoiceListParamsSortMinusNumber      InvoiceListParamsSort = "-NUMBER"
	InvoiceListParamsSortPlusType         InvoiceListParamsSort = "+TYPE"
	InvoiceListParamsSortMinusType        InvoiceListParamsSort = "-TYPE"
	InvoiceListParamsSortPlusAmount       InvoiceListParamsSort = "+AMOUNT"
	InvoiceListParamsSortMinusAmount      InvoiceListParamsSort = "-AMOUNT"
	InvoiceListParamsSortPlusStatus       InvoiceListParamsSort = "+STATUS"
	InvoiceListParamsSortMinusStatus      InvoiceListParamsSort = "-STATUS"
	InvoiceListParamsSortPlusDate         InvoiceListParamsSort = "+DATE"
	InvoiceListParamsSortMinusDate        InvoiceListParamsSort = "-DATE"
	InvoiceListParamsSortPlusUpdate       InvoiceListParamsSort = "+UPDATE"
	InvoiceListParamsSortMinusUpdate      InvoiceListParamsSort = "-UPDATE"
	InvoiceListParamsSortPlusDescription  InvoiceListParamsSort = "+DESCRIPTION"
	InvoiceListParamsSortMinusDescription InvoiceListParamsSort = "-DESCRIPTION"
)

// Status
type InvoiceListParamsStatus string

const (
	InvoiceListParamsStatusPaid      InvoiceListParamsStatus = "PAID"
	InvoiceListParamsStatusInProcess InvoiceListParamsStatus = "IN_PROCESS"
	InvoiceListParamsStatusUnpaid    InvoiceListParamsStatus = "UNPAID"
	InvoiceListParamsStatusVoid      InvoiceListParamsStatus = "VOID"
	InvoiceListParamsStatusExpired   InvoiceListParamsStatus = "EXPIRED"
	InvoiceListParamsStatusPrinted   InvoiceListParamsStatus = "PRINTED"
	InvoiceListParamsStatusMailed    InvoiceListParamsStatus = "MAILED"
	InvoiceListParamsStatusFailed    InvoiceListParamsStatus = "FAILED"
	InvoiceListParamsStatusRefunded  InvoiceListParamsStatus = "REFUNDED"
)

type InvoicePayParams struct {
	// Unique identifier for invoice
	ID string `json:"id,required" format:"uuid"`
	// Invoice amount
	Amount float64 `json:"amount,required"`
	// Optional credit account id for funds (if sender has multiple bank accounts)
	Account param.Opt[string] `json:"account,omitzero" format:"uuid"`
	paramObj
}

func (r InvoicePayParams) MarshalJSON() (data []byte, err error) {
	type shadow InvoicePayParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoicePayParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
