// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package checkbook

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Munchpass/checkbook/internal/apijson"
	"github.com/Munchpass/checkbook/internal/apiquery"
	"github.com/Munchpass/checkbook/internal/requestconfig"
	"github.com/Munchpass/checkbook/option"
	"github.com/Munchpass/checkbook/packages/param"
	"github.com/Munchpass/checkbook/packages/respjson"
)

// MailboxService contains methods and other services that help with interacting
// with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMailboxService] method instead.
type MailboxService struct {
	Options []option.RequestOption
	Mail    MailboxMailService
}

// NewMailboxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMailboxService(opts ...option.RequestOption) (r MailboxService) {
	r = MailboxService{}
	r.Options = opts
	r.Mail = NewMailboxMailService(opts...)
	return
}

// Create a new mailbox
func (r *MailboxService) New(ctx context.Context, opts ...option.RequestOption) (res *CreateMailboxResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/mailbox"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get mailbox details
func (r *MailboxService) Get(ctx context.Context, mailboxID string, opts ...option.RequestOption) (res *CreateMailboxResponse, err error) {
	opts = append(r.Options[:], opts...)
	if mailboxID == "" {
		err = errors.New("missing required mailbox_id parameter")
		return
	}
	path := fmt.Sprintf("v3/mailbox/%s", mailboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Return the mailboxes for the current user
func (r *MailboxService) List(ctx context.Context, query MailboxListParams, opts ...option.RequestOption) (res *MailboxListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/mailbox"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CreateMailboxResponse struct {
	// Unique identifier for mailbox
	ID string `json:"id,required"`
	// Mailbox creation timestamp
	Date string `json:"date,required"`
	// Mailbox status
	Status  string         `json:"status,required"`
	Address MailboxAddress `json:"address"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Date        respjson.Field
		Status      respjson.Field
		Address     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateMailboxResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateMailboxResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxAddress struct {
	// City
	City string `json:"city,required"`
	// Line 1
	Line1 string `json:"line_1,required"`
	// State
	State string `json:"state,required"`
	// Zip/postal code
	Zip string `json:"zip,required"`
	// Country
	Country string `json:"country,nullable"`
	// Line 2
	Line2 string `json:"line_2,nullable"`
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
func (r MailboxAddress) RawJSON() string { return r.JSON.raw }
func (r *MailboxAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxListResponse struct {
	// List of sent/received payments
	Mailboxes []CreateMailboxResponse `json:"mailboxes,required"`
	Page      int64                   `json:"page,required"`
	Pages     int64                   `json:"pages,required"`
	Total     int64                   `json:"total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mailboxes   respjson.Field
		Page        respjson.Field
		Pages       respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxListResponse) RawJSON() string { return r.JSON.raw }
func (r *MailboxListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxListParams struct {
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Items per page
	//
	// Any of 10, 25, 50, 100, 250.
	PerPage int64 `query:"per_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MailboxListParams]'s query parameters as `url.Values`.
func (r MailboxListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
