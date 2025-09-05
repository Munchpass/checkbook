// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package checkbook

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/checkbook-go/internal/apijson"
	"github.com/stainless-sdks/checkbook-go/internal/requestconfig"
	"github.com/stainless-sdks/checkbook-go/option"
	"github.com/stainless-sdks/checkbook-go/packages/param"
	"github.com/stainless-sdks/checkbook-go/packages/respjson"
)

// AccountWireService contains methods and other services that help with
// interacting with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWireService] method instead.
type AccountWireService struct {
	Options []option.RequestOption
}

// NewAccountWireService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountWireService(opts ...option.RequestOption) (r AccountWireService) {
	r = AccountWireService{}
	r.Options = opts
	return
}

// Create a new wire account
func (r *AccountWireService) New(ctx context.Context, body AccountWireNewParams, opts ...option.RequestOption) (res *WireAccountResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/wire"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an existing wire account
func (r *AccountWireService) Update(ctx context.Context, accountID string, body AccountWireUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("v3/account/wire/%s", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Return the wire accounts
func (r *AccountWireService) List(ctx context.Context, opts ...option.RequestOption) (res *AccountWireListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/wire"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove an existing wire account
func (r *AccountWireService) Delete(ctx context.Context, wireID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if wireID == "" {
		err = errors.New("missing required wire_id parameter")
		return
	}
	path := fmt.Sprintf("v3/account/wire/%s", wireID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type WireAccountResponse struct {
	// Unique identifier for account
	ID string `json:"id,required"`
	// Last 4 of account number
	Account string `json:"account,required"`
	// Account creation timestamp
	Date string `json:"date,required"`
	// Routing number
	Routing string `json:"routing,required"`
	// Name of the bank account
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Account     respjson.Field
		Date        respjson.Field
		Routing     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WireAccountResponse) RawJSON() string { return r.JSON.raw }
func (r *WireAccountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWireListResponse struct {
	// List of wire accounts
	Accounts []WireAccountResponse `json:"accounts,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accounts    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountWireListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountWireListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWireNewParams struct {
	// Account number
	Account string `json:"account,required"`
	// Routing number
	Routing string `json:"routing,required"`
	// Account type
	//
	// Any of "CHECKING", "SAVINGS", "BUSINESS".
	Type AccountWireNewParamsType `json:"type,omitzero,required"`
	// Optional name
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r AccountWireNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountWireNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountWireNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account type
type AccountWireNewParamsType string

const (
	AccountWireNewParamsTypeChecking AccountWireNewParamsType = "CHECKING"
	AccountWireNewParamsTypeSavings  AccountWireNewParamsType = "SAVINGS"
	AccountWireNewParamsTypeBusiness AccountWireNewParamsType = "BUSINESS"
)

type AccountWireUpdateParams struct {
	// Name for account
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r AccountWireUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountWireUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountWireUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
