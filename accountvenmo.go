// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package checkbook

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Munchpass/checkbook/internal/apijson"
	"github.com/Munchpass/checkbook/internal/requestconfig"
	"github.com/Munchpass/checkbook/option"
	"github.com/Munchpass/checkbook/packages/param"
	"github.com/Munchpass/checkbook/packages/respjson"
)

// AccountVenmoService contains methods and other services that help with
// interacting with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountVenmoService] method instead.
type AccountVenmoService struct {
	Options []option.RequestOption
}

// NewAccountVenmoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountVenmoService(opts ...option.RequestOption) (r AccountVenmoService) {
	r = AccountVenmoService{}
	r.Options = opts
	return
}

// Add a new Venmo account for a user
func (r *AccountVenmoService) New(ctx context.Context, body AccountVenmoNewParams, opts ...option.RequestOption) (res *VenmoAccountResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/venmo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an existing Venmo account
func (r *AccountVenmoService) Update(ctx context.Context, venmoID string, body AccountVenmoUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if venmoID == "" {
		err = errors.New("missing required venmo_id parameter")
		return
	}
	path := fmt.Sprintf("v3/account/venmo/%s", venmoID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Return the Venmo accounts of a user
func (r *AccountVenmoService) List(ctx context.Context, opts ...option.RequestOption) (res *AccountVenmoListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/venmo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove an existing Venmo account
func (r *AccountVenmoService) Delete(ctx context.Context, venmoID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if venmoID == "" {
		err = errors.New("missing required venmo_id parameter")
		return
	}
	path := fmt.Sprintf("v3/account/venmo/%s", venmoID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type VenmoAccountResponse struct {
	// Unique identifier for Venmo account
	ID string `json:"id,required"`
	// Account creation timestamp
	Date string `json:"date,required"`
	// Name of the Venmo account
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
func (r VenmoAccountResponse) RawJSON() string { return r.JSON.raw }
func (r *VenmoAccountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVenmoListResponse struct {
	// List of Venmo accounts
	Accounts []VenmoAccountResponse `json:"accounts,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accounts    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountVenmoListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountVenmoListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVenmoNewParams struct {
	// Venmo username or phone number
	Username string `json:"username,required"`
	paramObj
}

func (r AccountVenmoNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountVenmoNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountVenmoNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVenmoUpdateParams struct {
	// Name for the Venmo account
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r AccountVenmoUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountVenmoUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountVenmoUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
