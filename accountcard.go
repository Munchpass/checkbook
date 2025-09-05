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

// AccountCardService contains methods and other services that help with
// interacting with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCardService] method instead.
type AccountCardService struct {
	Options []option.RequestOption
}

// NewAccountCardService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountCardService(opts ...option.RequestOption) (r AccountCardService) {
	r = AccountCardService{}
	r.Options = opts
	return
}

// Add a new card
func (r *AccountCardService) New(ctx context.Context, body AccountCardNewParams, opts ...option.RequestOption) (res *AccountCardNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/card"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update the specified card
func (r *AccountCardService) Update(ctx context.Context, cardID string, body AccountCardUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if cardID == "" {
		err = errors.New("missing required card_id parameter")
		return
	}
	path := fmt.Sprintf("v3/account/card/%s", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Return the cards
func (r *AccountCardService) List(ctx context.Context, opts ...option.RequestOption) (res *AccountCardListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/card"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove the specified card
func (r *AccountCardService) Delete(ctx context.Context, cardID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if cardID == "" {
		err = errors.New("missing required card_id parameter")
		return
	}
	path := fmt.Sprintf("v3/account/card/%s", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AccountCardNewResponse struct {
	// Unique identifier for card
	ID string `json:"id,required"`
	// Last 4 of card number
	CardNumber string `json:"card_number,required"`
	// Card creation timestamp
	Date string `json:"date,required"`
	// Indicates the default card
	Default bool `json:"default,required"`
	// Name of the card
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CardNumber  respjson.Field
		Date        respjson.Field
		Default     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountCardNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountCardNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCardListResponse struct {
	// List of cards
	Cards []AccountCardListResponseCard `json:"cards,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cards       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountCardListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountCardListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCardListResponseCard struct {
	// Unique identifier for card
	ID string `json:"id,required"`
	// Last 4 of card number
	CardNumber string `json:"card_number,required"`
	// Card creation timestamp
	Date string `json:"date,required"`
	// Indicates the default card
	Default bool `json:"default,required"`
	// Card expiration date
	ExpirationDate string `json:"expiration_date,required"`
	// Name of the card
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CardNumber     respjson.Field
		Date           respjson.Field
		Default        respjson.Field
		ExpirationDate respjson.Field
		Name           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountCardListResponseCard) RawJSON() string { return r.JSON.raw }
func (r *AccountCardListResponseCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCardNewParams struct {
	// Card number
	CardNumber string `json:"card_number,required"`
	// Expiration date (yyyy-mm)
	ExpirationDate string `json:"expiration_date,required"`
	// CVV code
	Cvv     param.Opt[string]           `json:"cvv,omitzero"`
	Address AccountCardNewParamsAddress `json:"address,omitzero"`
	paramObj
}

func (r AccountCardNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountCardNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountCardNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Line1 is required.
type AccountCardNewParamsAddress struct {
	// Street line 1
	Line1 string `json:"line_1,required"`
	// City
	City param.Opt[string] `json:"city,omitzero"`
	// Country
	Country param.Opt[string] `json:"country,omitzero"`
	// Street line 2
	Line2 param.Opt[string] `json:"line_2,omitzero"`
	// State
	State param.Opt[string] `json:"state,omitzero"`
	// Zip/postal code
	Zip param.Opt[string] `json:"zip,omitzero"`
	paramObj
}

func (r AccountCardNewParamsAddress) MarshalJSON() (data []byte, err error) {
	type shadow AccountCardNewParamsAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountCardNewParamsAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCardUpdateParams struct {
	// Indicates the default card
	Default param.Opt[bool] `json:"default,omitzero"`
	// Name of the card
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r AccountCardUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountCardUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountCardUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
