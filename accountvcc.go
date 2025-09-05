// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package checkbook

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Munchpass/checkbook/internal/apijson"
	"github.com/Munchpass/checkbook/internal/requestconfig"
	"github.com/Munchpass/checkbook/option"
	"github.com/Munchpass/checkbook/packages/param"
	"github.com/Munchpass/checkbook/packages/respjson"
)

// AccountVccService contains methods and other services that help with interacting
// with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountVccService] method instead.
type AccountVccService struct {
	Options     []option.RequestOption
	Transaction AccountVccTransactionService
}

// NewAccountVccService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountVccService(opts ...option.RequestOption) (r AccountVccService) {
	r = AccountVccService{}
	r.Options = opts
	r.Transaction = NewAccountVccTransactionService(opts...)
	return
}

// Add a new vcc
func (r *AccountVccService) New(ctx context.Context, body AccountVccNewParams, opts ...option.RequestOption) (res *AccountVccNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/vcc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update the specified vcc
func (r *AccountVccService) Update(ctx context.Context, vccID string, body AccountVccUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if vccID == "" {
		err = errors.New("missing required vcc_id parameter")
		return
	}
	path := fmt.Sprintf("v3/account/vcc/%s", vccID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Return the virtual cards
func (r *AccountVccService) List(ctx context.Context, opts ...option.RequestOption) (res *AccountVccListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/vcc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove the specified vcc
func (r *AccountVccService) Delete(ctx context.Context, vccID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if vccID == "" {
		err = errors.New("missing required vcc_id parameter")
		return
	}
	path := fmt.Sprintf("v3/account/vcc/%s", vccID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Address struct {
	// City
	City string `json:"city,nullable"`
	// Country
	Country string `json:"country,nullable"`
	// Street line 1
	Line1 string `json:"line_1,nullable"`
	// State
	State string `json:"state,nullable"`
	// Zip/postal code
	Zip string `json:"zip,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Line1       respjson.Field
		State       respjson.Field
		Zip         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Address) RawJSON() string { return r.JSON.raw }
func (r *Address) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Address to a AddressParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AddressParam.Overrides()
func (r Address) ToParam() AddressParam {
	return param.Override[AddressParam](json.RawMessage(r.RawJSON()))
}

type AddressParam struct {
	// City
	City param.Opt[string] `json:"city,omitzero"`
	// Country
	Country param.Opt[string] `json:"country,omitzero"`
	// Street line 1
	Line1 param.Opt[string] `json:"line_1,omitzero"`
	// State
	State param.Opt[string] `json:"state,omitzero"`
	// Zip/postal code
	Zip param.Opt[string] `json:"zip,omitzero"`
	paramObj
}

func (r AddressParam) MarshalJSON() (data []byte, err error) {
	type shadow AddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVccNewResponse struct {
	// Unique identifier for VCC
	ID string `json:"id,required"`
	// Last 4 of VCC number
	CardNumber string `json:"card_number,required"`
	// VCC expiration date
	ExpirationDate string `json:"expiration_date,required"`
	// CVV code
	Cvv string `json:"cvv"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CardNumber     respjson.Field
		ExpirationDate respjson.Field
		Cvv            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountVccNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountVccNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVccListResponse struct {
	// List of vcc accounts for user
	Vccs []AccountVccListResponseVcc `json:"vccs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Vccs        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountVccListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountVccListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVccListResponseVcc struct {
	// Unique identifier for VCC
	ID string `json:"id,required"`
	// VCC balance
	Balance string `json:"balance,required"`
	// Last 4 of VCC number
	CardNumber string `json:"card_number,required"`
	// VCC creation timestamp
	Date string `json:"date,required"`
	// Indicates the default VCC account for the user
	Default bool `json:"default,required"`
	// VCC expiration date
	ExpirationDate string  `json:"expiration_date,required"`
	Address        Address `json:"address"`
	// Name of the VCC
	Name string `json:"name,nullable"`
	// VCC ruleset prefix
	RulesetPrefix string `json:"ruleset_prefix,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Balance        respjson.Field
		CardNumber     respjson.Field
		Date           respjson.Field
		Default        respjson.Field
		ExpirationDate respjson.Field
		Address        respjson.Field
		Name           respjson.Field
		RulesetPrefix  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountVccListResponseVcc) RawJSON() string { return r.JSON.raw }
func (r *AccountVccListResponseVcc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVccNewParams struct {
	// Email address
	Email param.Opt[string] `json:"email,omitzero"`
	// Phone number
	Phone   param.Opt[string] `json:"phone,omitzero"`
	Address AddressParam      `json:"address,omitzero"`
	paramObj
}

func (r AccountVccNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountVccNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountVccNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountVccUpdateParams struct {
	// Indicates the default VCC account for the user
	Default param.Opt[bool] `json:"default,omitzero"`
	// Name of the VCC
	Name    param.Opt[string] `json:"name,omitzero"`
	Address AddressParam      `json:"address,omitzero"`
	paramObj
}

func (r AccountVccUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountVccUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountVccUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
