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

// AccountVccTransactionService contains methods and other services that help with
// interacting with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountVccTransactionService] method instead.
type AccountVccTransactionService struct {
	Options []option.RequestOption
}

// NewAccountVccTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountVccTransactionService(opts ...option.RequestOption) (r AccountVccTransactionService) {
	r = AccountVccTransactionService{}
	r.Options = opts
	return
}

// Get the requested transaction for the specified VCC
func (r *AccountVccTransactionService) Get(ctx context.Context, transactionID string, query AccountVccTransactionGetParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = append(r.Options[:], opts...)
	if query.VccID == "" {
		err = errors.New("missing required vcc_id parameter")
		return
	}
	if transactionID == "" {
		err = errors.New("missing required transaction_id parameter")
		return
	}
	path := fmt.Sprintf("v3/account/vcc/%s/transaction/%s", query.VccID, transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the transactions for the specified VCC
func (r *AccountVccTransactionService) List(ctx context.Context, vccID string, query AccountVccTransactionListParams, opts ...option.RequestOption) (res *AccountVccTransactionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if vccID == "" {
		err = errors.New("missing required vcc_id parameter")
		return
	}
	path := fmt.Sprintf("v3/account/vcc/%s/transaction", vccID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Transaction struct {
	// Transaction id
	ID string `json:"id,nullable"`
	// Amount
	Amount float64 `json:"amount"`
	// Balance
	Balance float64 `json:"balance"`
	// Creation timestamp
	CreatedTs string `json:"created_ts,nullable"`
	// Description
	Description string `json:"description,nullable"`
	// Action
	Status string `json:"status"`
	// Transaction type
	Type string `json:"type,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		Balance     respjson.Field
		CreatedTs   respjson.Field
		Description respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Transaction) RawJSON() string { return r.JSON.raw }
func (r *Transaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVccTransactionListResponse struct {
	// Transactions list
	Transactions []Transaction `json:"transactions,required"`
	// Current page
	Page int64 `json:"page"`
	// Total number of pages
	Pages int64 `json:"pages"`
	// Total number of transactions
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transactions respjson.Field
		Page         respjson.Field
		Pages        respjson.Field
		Total        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountVccTransactionListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountVccTransactionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVccTransactionGetParams struct {
	VccID string `path:"vcc_id,required" json:"-"`
	paramObj
}

type AccountVccTransactionListParams struct {
	Beta param.Opt[bool] `query:"beta,omitzero" json:"-"`
	// End date
	EndDate param.Opt[time.Time] `query:"end_date,omitzero" format:"date" json:"-"`
	// Start date
	StartDate param.Opt[time.Time] `query:"start_date,omitzero" format:"date" json:"-"`
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Items per page
	//
	// Any of 10, 25, 50.
	PerPage int64 `query:"per_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccountVccTransactionListParams]'s query parameters as
// `url.Values`.
func (r AccountVccTransactionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
