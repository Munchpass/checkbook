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

// AccountBankService contains methods and other services that help with
// interacting with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountBankService] method instead.
type AccountBankService struct {
	Options []option.RequestOption
	Iav     AccountBankIavService
}

// NewAccountBankService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountBankService(opts ...option.RequestOption) (r AccountBankService) {
	r = AccountBankService{}
	r.Options = opts
	r.Iav = NewAccountBankIavService(opts...)
	return
}

// Add a new bank account
func (r *AccountBankService) New(ctx context.Context, body AccountBankNewParams, opts ...option.RequestOption) (res *AccountBankNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/bank"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an existing bank account
func (r *AccountBankService) Update(ctx context.Context, bankID string, body AccountBankUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if bankID == "" {
		err = errors.New("missing required bank_id parameter")
		return
	}
	path := fmt.Sprintf("v3/account/bank/%s", bankID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Get the bank accounts for a user
func (r *AccountBankService) List(ctx context.Context, opts ...option.RequestOption) (res *AccountBankListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/bank"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove the specified bank account
func (r *AccountBankService) Delete(ctx context.Context, bankID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if bankID == "" {
		err = errors.New("missing required bank_id parameter")
		return
	}
	path := fmt.Sprintf("v3/account/bank/%s", bankID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Release the micro-deposits for a bank account
func (r *AccountBankService) Release(ctx context.Context, body AccountBankReleaseParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v3/account/bank/release"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Return a list of our supported institutions for instant account verification
func (r *AccountBankService) GetInstitutions(ctx context.Context, opts ...option.RequestOption) (res *AccountBankGetInstitutionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/account/bank/institutions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Verify the micro-deposits for a bank account
func (r *AccountBankService) Verify(ctx context.Context, body AccountBankVerifyParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v3/account/bank/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type AccountBankNewResponse struct {
	// Unique identifier for account
	ID string `json:"id,required"`
	// Last 4 of account number
	Account string `json:"account,required"`
	// Indicates the billing account for the user
	Billing bool `json:"billing,required"`
	// Account creation timestamp
	Date string `json:"date,required"`
	// Indicates the default account for the user
	Default bool `json:"default,required"`
	// Routing number
	Routing string `json:"routing,required"`
	// Current status of account
	//
	// Any of "PENDING", "VERIFIED", "DEPOSIT_ONLY".
	Status AccountBankNewResponseStatus `json:"status,required"`
	// Bank account type
	//
	// Any of "CHECKING", "SAVINGS", "BUSINESS".
	Type AccountBankNewResponseType `json:"type,required"`
	// Indicates the current amount left in the account's balance (only for prefunded
	// accounts)
	Balance float64 `json:"balance,nullable"`
	// Name of the bank account
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Account     respjson.Field
		Billing     respjson.Field
		Date        respjson.Field
		Default     respjson.Field
		Routing     respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Balance     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBankNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountBankNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of account
type AccountBankNewResponseStatus string

const (
	AccountBankNewResponseStatusPending     AccountBankNewResponseStatus = "PENDING"
	AccountBankNewResponseStatusVerified    AccountBankNewResponseStatus = "VERIFIED"
	AccountBankNewResponseStatusDepositOnly AccountBankNewResponseStatus = "DEPOSIT_ONLY"
)

// Bank account type
type AccountBankNewResponseType string

const (
	AccountBankNewResponseTypeChecking AccountBankNewResponseType = "CHECKING"
	AccountBankNewResponseTypeSavings  AccountBankNewResponseType = "SAVINGS"
	AccountBankNewResponseTypeBusiness AccountBankNewResponseType = "BUSINESS"
)

type AccountBankListResponse struct {
	// List of bank accounts for user
	Banks []AccountBankListResponseBank `json:"banks,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Banks       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBankListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountBankListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBankListResponseBank struct {
	// Unique identifier for account
	ID string `json:"id,required"`
	// Last 4 of account number
	Account string `json:"account,required"`
	// Indicates the billing account for the user
	Billing bool `json:"billing,required"`
	// Account creation timestamp
	Date string `json:"date,required"`
	// Indicates the default account for the user
	Default bool `json:"default,required"`
	// Routing number
	Routing string `json:"routing,required"`
	// Current status of account
	//
	// Any of "PENDING", "VERIFIED", "DEPOSIT_ONLY".
	Status string `json:"status,required"`
	// Bank account type
	//
	// Any of "CHECKING", "SAVINGS", "BUSINESS".
	Type string `json:"type,required"`
	// Indicates the current amount left in the account's balance (only for prefunded
	// accounts)
	Balance float64 `json:"balance,nullable"`
	// Name of the bank account
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Account     respjson.Field
		Billing     respjson.Field
		Date        respjson.Field
		Default     respjson.Field
		Routing     respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Balance     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBankListResponseBank) RawJSON() string { return r.JSON.raw }
func (r *AccountBankListResponseBank) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBankGetInstitutionsResponse struct {
	// List of supported institutions for IAV
	Institutions []AccountBankGetInstitutionsResponseInstitution `json:"institutions,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Institutions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBankGetInstitutionsResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountBankGetInstitutionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBankGetInstitutionsResponseInstitution struct {
	// Login form for institution
	LoginForm []AccountBankGetInstitutionsResponseInstitutionLoginForm `json:"login_form,required"`
	// Name of institution
	Name string `json:"name,required"`
	// Unique identifier for institution
	ID string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LoginForm   respjson.Field
		Name        respjson.Field
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBankGetInstitutionsResponseInstitution) RawJSON() string { return r.JSON.raw }
func (r *AccountBankGetInstitutionsResponseInstitution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Login form field
type AccountBankGetInstitutionsResponseInstitutionLoginForm struct {
	// Description of the field
	Description string `json:"description"`
	// Name of the field
	Name string `json:"name"`
	// Field type
	//
	// Any of "TEXT", "PASSWORD", "OPTION".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBankGetInstitutionsResponseInstitutionLoginForm) RawJSON() string { return r.JSON.raw }
func (r *AccountBankGetInstitutionsResponseInstitutionLoginForm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBankNewParams struct {
	// Account number
	Account string `json:"account,required"`
	// Routing number
	Routing string `json:"routing,required"`
	// Account type
	//
	// Any of "CHECKING", "SAVINGS", "BUSINESS".
	Type AccountBankNewParamsType `json:"type,omitzero,required"`
	// Optional name
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r AccountBankNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountBankNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountBankNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account type
type AccountBankNewParamsType string

const (
	AccountBankNewParamsTypeChecking AccountBankNewParamsType = "CHECKING"
	AccountBankNewParamsTypeSavings  AccountBankNewParamsType = "SAVINGS"
	AccountBankNewParamsTypeBusiness AccountBankNewParamsType = "BUSINESS"
)

type AccountBankUpdateParams struct {
	// Specification of billing account
	Billing param.Opt[bool] `json:"billing,omitzero"`
	// Specification of default account
	Default param.Opt[bool] `json:"default,omitzero"`
	// Name for account
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r AccountBankUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountBankUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountBankUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBankReleaseParams struct {
	// ID of the bank account
	Account string `json:"account,required"`
	paramObj
}

func (r AccountBankReleaseParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountBankReleaseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountBankReleaseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBankVerifyParams struct {
	// Amount of first microdeposit
	Amount1 float64 `json:"amount_1,required"`
	// Amount of second microdeposit
	Amount2 float64 `json:"amount_2,required"`
	// ID of account to verify
	Account param.Opt[string] `json:"account,omitzero"`
	paramObj
}

func (r AccountBankVerifyParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountBankVerifyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountBankVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
