// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package checkbook

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/checkbook-go/internal/apijson"
	"github.com/stainless-sdks/checkbook-go/internal/apiquery"
	"github.com/stainless-sdks/checkbook-go/internal/requestconfig"
	"github.com/stainless-sdks/checkbook-go/option"
	"github.com/stainless-sdks/checkbook-go/packages/param"
	"github.com/stainless-sdks/checkbook-go/packages/respjson"
)

// MailboxMailService contains methods and other services that help with
// interacting with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMailboxMailService] method instead.
type MailboxMailService struct {
	Options []option.RequestOption
}

// NewMailboxMailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMailboxMailService(opts ...option.RequestOption) (r MailboxMailService) {
	r = MailboxMailService{}
	r.Options = opts
	return
}

// Get mailbox item
func (r *MailboxMailService) Get(ctx context.Context, itemID string, query MailboxMailGetParams, opts ...option.RequestOption) (res *MailResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.MailboxID == "" {
		err = errors.New("missing required mailbox_id parameter")
		return
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return
	}
	path := fmt.Sprintf("v3/mailbox/%s/mail/%s", query.MailboxID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get mailbox items
func (r *MailboxMailService) List(ctx context.Context, mailboxID string, query MailboxMailListParams, opts ...option.RequestOption) (res *MailboxMailListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if mailboxID == "" {
		err = errors.New("missing required mailbox_id parameter")
		return
	}
	path := fmt.Sprintf("v3/mailbox/%s/mail", mailboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get mailbox item
func (r *MailboxMailService) GetAttachment(ctx context.Context, itemID string, query MailboxMailGetAttachmentParams, opts ...option.RequestOption) (res *Error, err error) {
	opts = append(r.Options[:], opts...)
	if query.MailboxID == "" {
		err = errors.New("missing required mailbox_id parameter")
		return
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return
	}
	path := fmt.Sprintf("v3/mailbox/%s/mail/%s/attachment", query.MailboxID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Counterparty struct {
	Address MailboxAddress `json:"address,required"`
	// Name
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Counterparty) RawJSON() string { return r.JSON.raw }
func (r *Counterparty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailResponse struct {
	// Unique identifier for mail
	ID string `json:"id,required"`
	// Mail creation timestamp
	Date      string       `json:"date,required"`
	Recipient Counterparty `json:"recipient,required"`
	Sender    Counterparty `json:"sender,required"`
	// Mail status
	Status string `json:"status,required"`
	// List of checks present in mail
	Checks []MailResponseCheck `json:"checks"`
	Tags   map[string]string   `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Date        respjson.Field
		Recipient   respjson.Field
		Sender      respjson.Field
		Status      respjson.Field
		Checks      respjson.Field
		Tags        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailResponse) RawJSON() string { return r.JSON.raw }
func (r *MailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailResponseCheck struct {
	// Unique identifier for check
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailResponseCheck) RawJSON() string { return r.JSON.raw }
func (r *MailResponseCheck) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxMailListResponse struct {
	// List of sent/received payments
	Mail  []MailResponse `json:"mail,required"`
	Page  int64          `json:"page,required"`
	Pages int64          `json:"pages,required"`
	Total int64          `json:"total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mail        respjson.Field
		Page        respjson.Field
		Pages       respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailboxMailListResponse) RawJSON() string { return r.JSON.raw }
func (r *MailboxMailListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailboxMailGetParams struct {
	MailboxID string `path:"mailbox_id,required" json:"-"`
	paramObj
}

type MailboxMailListParams struct {
	// End date
	EndDate param.Opt[time.Time] `query:"end_date,omitzero" format:"date" json:"-"`
	// Start date
	StartDate param.Opt[time.Time] `query:"start_date,omitzero" format:"date" json:"-"`
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Query
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Items per page
	//
	// Any of 10, 25, 50, 100, 250.
	PerPage int64 `query:"per_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MailboxMailListParams]'s query parameters as `url.Values`.
func (r MailboxMailListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MailboxMailGetAttachmentParams struct {
	MailboxID string `path:"mailbox_id,required" json:"-"`
	paramObj
}
