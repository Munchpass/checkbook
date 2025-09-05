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

// AccountPaypalService contains methods and other services that help with
// interacting with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountPaypalService] method instead.
type AccountPaypalService struct {
	Options []option.RequestOption
}

// NewAccountPaypalService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountPaypalService(opts ...option.RequestOption) (r AccountPaypalService) {
	r = AccountPaypalService{}
	r.Options = opts
	return
}

// Add a new Paypal account for a user
func (r *AccountPaypalService) New(ctx context.Context, body AccountPaypalNewParams, opts ...option.RequestOption) (res *PaypalAccountResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/paypal"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an existing Paypal account
func (r *AccountPaypalService) Update(ctx context.Context, paypalID string, body AccountPaypalUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if paypalID == "" {
		err = errors.New("missing required paypal_id parameter")
		return
	}
	path := fmt.Sprintf("v3/account/paypal/%s", paypalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Return the Paypal accounts of a user
func (r *AccountPaypalService) List(ctx context.Context, opts ...option.RequestOption) (res *AccountPaypalListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/paypal"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove an existing PayPal account
func (r *AccountPaypalService) Delete(ctx context.Context, paypalID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if paypalID == "" {
		err = errors.New("missing required paypal_id parameter")
		return
	}
	path := fmt.Sprintf("v3/account/paypal/%s", paypalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type PaypalAccountResponse struct {
	// Unique identifier for Paypal account
	ID string `json:"id,required"`
	// Account creation timestamp
	Date string `json:"date,required"`
	// Name of the PayPal account
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
func (r PaypalAccountResponse) RawJSON() string { return r.JSON.raw }
func (r *PaypalAccountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPaypalListResponse struct {
	// List of Paypal accounts
	Accounts []PaypalAccountResponse `json:"accounts,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accounts    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountPaypalListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountPaypalListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPaypalNewParams struct {
	// PayPal email or phone number
	Username string `json:"username,required"`
	paramObj
}

func (r AccountPaypalNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountPaypalNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountPaypalNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPaypalUpdateParams struct {
	// Name for the PayPal account
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r AccountPaypalUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountPaypalUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountPaypalUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
