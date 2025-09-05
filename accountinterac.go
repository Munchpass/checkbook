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

// AccountInteracService contains methods and other services that help with
// interacting with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountInteracService] method instead.
type AccountInteracService struct {
	Options []option.RequestOption
}

// NewAccountInteracService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountInteracService(opts ...option.RequestOption) (r AccountInteracService) {
	r = AccountInteracService{}
	r.Options = opts
	return
}

// Add a new Interac account for a user
func (r *AccountInteracService) New(ctx context.Context, body AccountInteracNewParams, opts ...option.RequestOption) (res *InteracAccountResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/interac"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an existing Interac account
func (r *AccountInteracService) Update(ctx context.Context, interacID string, body AccountInteracUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if interacID == "" {
		err = errors.New("missing required interac_id parameter")
		return
	}
	path := fmt.Sprintf("v3/account/interac/%s", interacID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Return the Interac accounts of a user
func (r *AccountInteracService) List(ctx context.Context, opts ...option.RequestOption) (res *AccountInteracListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/interac"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove an existing Interac account
func (r *AccountInteracService) Delete(ctx context.Context, interacID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if interacID == "" {
		err = errors.New("missing required interac_id parameter")
		return
	}
	path := fmt.Sprintf("v3/account/interac/%s", interacID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type InteracAccountResponse struct {
	// Unique identifier for Interac account
	ID string `json:"id,required"`
	// Account creation timestamp
	Date string `json:"date,required"`
	// Name of the Interac account
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Date        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InteracAccountResponse) RawJSON() string { return r.JSON.raw }
func (r *InteracAccountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountInteracListResponse struct {
	// List of Interac accounts
	Accounts []InteracAccountResponse `json:"accounts,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accounts    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountInteracListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountInteracListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountInteracNewParams struct {
	// Interac username or phone number
	Username string `json:"username,required"`
	paramObj
}

func (r AccountInteracNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountInteracNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountInteracNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountInteracUpdateParams struct {
	// Name for the Interac account
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r AccountInteracUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountInteracUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountInteracUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
