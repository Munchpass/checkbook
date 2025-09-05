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

// DirectoryAccountService contains methods and other services that help with
// interacting with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDirectoryAccountService] method instead.
type DirectoryAccountService struct {
	Options []option.RequestOption
}

// NewDirectoryAccountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDirectoryAccountService(opts ...option.RequestOption) (r DirectoryAccountService) {
	r = DirectoryAccountService{}
	r.Options = opts
	return
}

// Remove a directory account
func (r *DirectoryAccountService) Delete(ctx context.Context, accountID string, body DirectoryAccountDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.DirectoryID == "" {
		err = errors.New("missing required directory_id parameter")
		return
	}
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("v3/directory/%s/account/%s", body.DirectoryID, accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Create a new directory bank account
func (r *DirectoryAccountService) NewBank(ctx context.Context, directoryID string, body DirectoryAccountNewBankParams, opts ...option.RequestOption) (res *DirectoryAccountNewBankResponse, err error) {
	opts = append(r.Options[:], opts...)
	if directoryID == "" {
		err = errors.New("missing required directory_id parameter")
		return
	}
	path := fmt.Sprintf("v3/directory/%s/account/bank", directoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new directory card account
func (r *DirectoryAccountService) NewCard(ctx context.Context, directoryID string, body DirectoryAccountNewCardParams, opts ...option.RequestOption) (res *DirectoryAccountNewCardResponse, err error) {
	opts = append(r.Options[:], opts...)
	if directoryID == "" {
		err = errors.New("missing required directory_id parameter")
		return
	}
	path := fmt.Sprintf("v3/directory/%s/account/card", directoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DirectoryAccountNewBankResponse struct {
	// Unique identifier for bank account
	ID string `json:"id,required"`
	// Bank account name
	Name string `json:"name,required"`
	// Last 4 of bank account
	Number string `json:"number,required"`
	// Account type
	//
	// Any of "BANK".
	Type DirectoryAccountNewBankResponseType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Number      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectoryAccountNewBankResponse) RawJSON() string { return r.JSON.raw }
func (r *DirectoryAccountNewBankResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account type
type DirectoryAccountNewBankResponseType string

const (
	DirectoryAccountNewBankResponseTypeBank DirectoryAccountNewBankResponseType = "BANK"
)

type DirectoryAccountNewCardResponse struct {
	// Unique identifier for card
	ID string `json:"id,required"`
	// Card name
	Name string `json:"name,required"`
	// Last 4 of card number
	Number string `json:"number,required"`
	// Account type
	//
	// Any of "CARD".
	Type DirectoryAccountNewCardResponseType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Number      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectoryAccountNewCardResponse) RawJSON() string { return r.JSON.raw }
func (r *DirectoryAccountNewCardResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account type
type DirectoryAccountNewCardResponseType string

const (
	DirectoryAccountNewCardResponseTypeCard DirectoryAccountNewCardResponseType = "CARD"
)

type DirectoryAccountDeleteParams struct {
	DirectoryID string `path:"directory_id,required" json:"-"`
	paramObj
}

type DirectoryAccountNewBankParams struct {
	// Account number
	Account string `json:"account,required"`
	// Routing number
	Routing string `json:"routing,required"`
	// Account type
	//
	// Any of "CHECKING", "SAVINGS", "BUSINESS".
	Type DirectoryAccountNewBankParamsType `json:"type,omitzero,required"`
	// Optional name
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r DirectoryAccountNewBankParams) MarshalJSON() (data []byte, err error) {
	type shadow DirectoryAccountNewBankParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirectoryAccountNewBankParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account type
type DirectoryAccountNewBankParamsType string

const (
	DirectoryAccountNewBankParamsTypeChecking DirectoryAccountNewBankParamsType = "CHECKING"
	DirectoryAccountNewBankParamsTypeSavings  DirectoryAccountNewBankParamsType = "SAVINGS"
	DirectoryAccountNewBankParamsTypeBusiness DirectoryAccountNewBankParamsType = "BUSINESS"
)

type DirectoryAccountNewCardParams struct {
	// Card number
	CardNumber string `json:"card_number,required"`
	// Card expiration date
	ExpirationDate string `json:"expiration_date,required"`
	// Card name
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r DirectoryAccountNewCardParams) MarshalJSON() (data []byte, err error) {
	type shadow DirectoryAccountNewCardParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirectoryAccountNewCardParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
