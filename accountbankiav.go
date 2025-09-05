// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package checkbook

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/checkbook-go/internal/apijson"
	"github.com/stainless-sdks/checkbook-go/internal/requestconfig"
	"github.com/stainless-sdks/checkbook-go/option"
	"github.com/stainless-sdks/checkbook-go/packages/param"
	"github.com/stainless-sdks/checkbook-go/packages/respjson"
)

// AccountBankIavService contains methods and other services that help with
// interacting with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountBankIavService] method instead.
type AccountBankIavService struct {
	Options []option.RequestOption
}

// NewAccountBankIavService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountBankIavService(opts ...option.RequestOption) (r AccountBankIavService) {
	r = AccountBankIavService{}
	r.Options = opts
	return
}

// Add a new bank account with instant account verification
func (r *AccountBankIavService) New(ctx context.Context, body AccountBankIavNewParams, opts ...option.RequestOption) (res *AccountBankIavNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/bank/iav"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve the bank account(s) associated with the Plaid token
func (r *AccountBankIavService) Plaid(ctx context.Context, body AccountBankIavPlaidParams, opts ...option.RequestOption) (res *AccountBankIavPlaidResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/bank/iav/plaid"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountBankIavNewResponse struct {
	// List of accounts and information
	Accounts []AccountBankIavNewResponseAccount `json:"accounts"`
	// `institution_id` returned from the institutions endpoint
	InstitutionID string `json:"institution_id"`
	// MFA questions
	Mfa []AccountBankIavNewResponseMfa `json:"mfa"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accounts      respjson.Field
		InstitutionID respjson.Field
		Mfa           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBankIavNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountBankIavNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBankIavNewResponseAccount struct {
	// Last 4 of account number
	Account string `json:"account,required"`
	// Routing number
	Routing string `json:"routing,required"`
	// Name of the bank account
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account     respjson.Field
		Routing     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBankIavNewResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *AccountBankIavNewResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBankIavNewResponseMfa struct {
	// MFA description
	Description string `json:"description"`
	// MFA image
	Image string `json:"image"`
	// MFA name
	Name string `json:"name"`
	// MFA selections
	Selections []AccountBankIavNewResponseMfaSelection `json:"selections"`
	// MFA type
	//
	// Any of "SELECTION", "IMAGE", "TEXT".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Image       respjson.Field
		Name        respjson.Field
		Selections  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBankIavNewResponseMfa) RawJSON() string { return r.JSON.raw }
func (r *AccountBankIavNewResponseMfa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBankIavNewResponseMfaSelection struct {
	// MFA field description
	Description string `json:"description"`
	// MFA field name
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBankIavNewResponseMfaSelection) RawJSON() string { return r.JSON.raw }
func (r *AccountBankIavNewResponseMfaSelection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBankIavPlaidResponse struct {
	// List of valid bank accounts
	Accounts []AccountBankIavPlaidResponseAccount `json:"accounts"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accounts    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBankIavPlaidResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountBankIavPlaidResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBankIavPlaidResponseAccount struct {
	// Last 4 of account number
	Account string `json:"account,required"`
	// Routing number
	Routing string `json:"routing,required"`
	// Name of the bank account
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account     respjson.Field
		Routing     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBankIavPlaidResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *AccountBankIavPlaidResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBankIavNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfIavLoginSchema *AccountBankIavNewParamsBodyIavLoginSchema `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfIavmfaSchema *AccountBankIavNewParamsBodyIavmfaSchema `json:",inline"`

	paramObj
}

func (u AccountBankIavNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfIavLoginSchema, u.OfIavmfaSchema)
}
func (r *AccountBankIavNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property InstitutionID is required.
type AccountBankIavNewParamsBodyIavLoginSchema struct {
	// `institution_id` returned from the institutions endpoint
	InstitutionID string         `json:"institution_id,required"`
	ExtraFields   map[string]any `json:"-"`
	paramObj
}

func (r AccountBankIavNewParamsBodyIavLoginSchema) MarshalJSON() (data []byte, err error) {
	type shadow AccountBankIavNewParamsBodyIavLoginSchema
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *AccountBankIavNewParamsBodyIavLoginSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property InstitutionID is required.
type AccountBankIavNewParamsBodyIavmfaSchema struct {
	// `institution_id` returned from the institutions endpoint
	InstitutionID string                                       `json:"institution_id,required"`
	Mfa           []AccountBankIavNewParamsBodyIavmfaSchemaMfa `json:"mfa,omitzero"`
	paramObj
}

func (r AccountBankIavNewParamsBodyIavmfaSchema) MarshalJSON() (data []byte, err error) {
	type shadow AccountBankIavNewParamsBodyIavmfaSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountBankIavNewParamsBodyIavmfaSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBankIavNewParamsBodyIavmfaSchemaMfa struct {
	Name     param.Opt[string] `json:"name,omitzero"`
	Response param.Opt[string] `json:"response,omitzero"`
	// Any of "IMAGE", "SELECTION", "TEXT".
	Type        string         `json:"type,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r AccountBankIavNewParamsBodyIavmfaSchemaMfa) MarshalJSON() (data []byte, err error) {
	type shadow AccountBankIavNewParamsBodyIavmfaSchemaMfa
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *AccountBankIavNewParamsBodyIavmfaSchemaMfa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AccountBankIavNewParamsBodyIavmfaSchemaMfa](
		"type", "IMAGE", "SELECTION", "TEXT",
	)
}

type AccountBankIavPlaidParams struct {
	// Plaid processor token
	ProcessorToken string `json:"processor_token,required"`
	paramObj
}

func (r AccountBankIavPlaidParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountBankIavPlaidParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountBankIavPlaidParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
