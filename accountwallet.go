// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package checkbook

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/checkbook-go/internal/apijson"
	shimjson "github.com/stainless-sdks/checkbook-go/internal/encoding/json"
	"github.com/stainless-sdks/checkbook-go/internal/requestconfig"
	"github.com/stainless-sdks/checkbook-go/option"
	"github.com/stainless-sdks/checkbook-go/packages/param"
	"github.com/stainless-sdks/checkbook-go/packages/respjson"
)

// AccountWalletService contains methods and other services that help with
// interacting with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWalletService] method instead.
type AccountWalletService struct {
	Options []option.RequestOption
}

// NewAccountWalletService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountWalletService(opts ...option.RequestOption) (r AccountWalletService) {
	r = AccountWalletService{}
	r.Options = opts
	return
}

// Update wallet
func (r *AccountWalletService) New(ctx context.Context, body AccountWalletNewParams, opts ...option.RequestOption) (res *AccountWalletNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/wallet"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get wallet accounts for user
func (r *AccountWalletService) List(ctx context.Context, opts ...option.RequestOption) (res *AccountWalletListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/wallet"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a wallet account
func (r *AccountWalletService) Delete(ctx context.Context, walletID string, opts ...option.RequestOption) (res *AccountWalletDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return
	}
	path := fmt.Sprintf("v3/account/wallet/%s", walletID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CreateWalletRequestParam struct {
	// Optional name
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r CreateWalletRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateWalletRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateWalletRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreateWalletResponse struct {
	// Unique identifier for account
	ID string `json:"id,required"`
	// Account creation timestamp
	Date string `json:"date,required"`
	// Name of the wallet
	Name    string                      `json:"name,nullable"`
	Numbers CreateWalletResponseNumbers `json:"numbers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Date        respjson.Field
		Name        respjson.Field
		Numbers     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateWalletResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateWalletResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreateWalletResponseNumbers struct {
	ACH  Number140416527981344 `json:"ACH"`
	Rtp  Number140416527981344 `json:"RTP"`
	Wire Number140416527981344 `json:"WIRE"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ACH         respjson.Field
		Rtp         respjson.Field
		Wire        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateWalletResponseNumbers) RawJSON() string { return r.JSON.raw }
func (r *CreateWalletResponseNumbers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number140416527981344 struct {
	// Account number
	Account string `json:"account"`
	// Routing number
	Routing string `json:"routing"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account     respjson.Field
		Routing     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number140416527981344) RawJSON() string { return r.JSON.raw }
func (r *Number140416527981344) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWalletNewResponse = any

type AccountWalletListResponse struct {
	Wallets []CreateWalletResponse `json:"wallets"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Wallets     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountWalletListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountWalletListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWalletDeleteResponse = any

type AccountWalletNewParams struct {
	CreateWalletRequest CreateWalletRequestParam
	paramObj
}

func (r AccountWalletNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateWalletRequest)
}
func (r *AccountWalletNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateWalletRequest)
}
