// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package checkbook

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/checkbook-go/internal/apijson"
	"github.com/stainless-sdks/checkbook-go/internal/apiquery"
	"github.com/stainless-sdks/checkbook-go/internal/requestconfig"
	"github.com/stainless-sdks/checkbook-go/option"
	"github.com/stainless-sdks/checkbook-go/packages/param"
	"github.com/stainless-sdks/checkbook-go/packages/respjson"
)

// UserService contains methods and other services that help with interacting with
// the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
	APIKey  UserAPIKeyService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	r.APIKey = NewUserAPIKeyService(opts...)
	return
}

// Create a new marketplace user
func (r *UserService) New(ctx context.Context, body UserNewParams, opts ...option.RequestOption) (res *UserNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get user information
func (r *UserService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update existing user information
func (r *UserService) Update(ctx context.Context, body UserUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v3/user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Return the marketplace users
func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *UserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/user/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete the marketplace user
func (r *UserService) Delete(ctx context.Context, userID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("v3/user/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Add signature
func (r *UserService) AddSignature(ctx context.Context, body UserAddSignatureParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v3/user/signature"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Payment fields
type ColorParam struct {
	// Payment expiration
	Primary param.Opt[string] `json:"primary,omitzero"`
	// Payment expiration
	Secondary param.Opt[string] `json:"secondary,omitzero"`
	paramObj
}

func (r ColorParam) MarshalJSON() (data []byte, err error) {
	type shadow ColorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ColorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Address field
type MerchantResponseAddress struct {
	// City
	City string `json:"city,nullable"`
	// Country
	Country string `json:"country,nullable"`
	// Street line 1
	Line1 string `json:"line_1,nullable"`
	// Street line 2
	Line2 string `json:"line_2,nullable"`
	// State
	State string `json:"state,nullable"`
	// Zip/postal code
	Zip string `json:"zip,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Line1       respjson.Field
		Line2       respjson.Field
		State       respjson.Field
		Zip         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MerchantResponseAddress) RawJSON() string { return r.JSON.raw }
func (r *MerchantResponseAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Address fields
type UpdateMerchantAddressParam struct {
	// City
	City param.Opt[string] `json:"city,omitzero"`
	// Country
	Country param.Opt[string] `json:"country,omitzero"`
	// Street line 1
	Line1 param.Opt[string] `json:"line_1,omitzero"`
	// Street line 2
	Line2 param.Opt[string] `json:"line_2,omitzero"`
	// State
	State param.Opt[string] `json:"state,omitzero"`
	// Zip/postal code
	Zip param.Opt[string] `json:"zip,omitzero"`
	paramObj
}

func (r UpdateMerchantAddressParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateMerchantAddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateMerchantAddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response fields for user creation
type UserNewResponse struct {
	// Unique identifier for user
	ID string `json:"id,required"`
	// Publishable key of the user
	Key string `json:"key,required"`
	// Secret key of the user
	Secret string `json:"secret,required"`
	// Specific user_id provided by the marketplace owner
	UserID string `json:"user_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Key         respjson.Field
		Secret      respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserNewResponse) RawJSON() string { return r.JSON.raw }
func (r *UserNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response fields for user retrieval
type UserGetResponse struct {
	// Brand field
	Brand UserGetResponseBrand `json:"brand"`
	// Memrchant field
	Merchant UserGetResponseMerchant `json:"merchant"`
	// User field
	User UserGetResponseUser `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brand       respjson.Field
		Merchant    respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetResponse) RawJSON() string { return r.JSON.raw }
func (r *UserGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Brand field
type UserGetResponseBrand struct {
	// Email footer
	Footer string `json:"footer,nullable"`
	// Base64 encoded company logo
	Logo string `json:"logo"`
	// Email reply to address
	ReplyTo string `json:"reply_to,nullable"`
	// Email slogan
	Slogan string `json:"slogan,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Footer      respjson.Field
		Logo        respjson.Field
		ReplyTo     respjson.Field
		Slogan      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetResponseBrand) RawJSON() string { return r.JSON.raw }
func (r *UserGetResponseBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Memrchant field
type UserGetResponseMerchant struct {
	// Address field
	Address MerchantResponseAddress `json:"address"`
	// Business incorporation or formation date
	IncorporationDate string `json:"incorporation_date,nullable"`
	// Industry sector
	Industry string `json:"industry,nullable"`
	// Principal officer first name
	LegalFirstname string `json:"legal_firstname,nullable"`
	// Principal officer last name
	LegalLastname string `json:"legal_lastname,nullable"`
	// Address field
	PrincipalAddress MerchantResponseAddress `json:"principal_address"`
	// One of a standard set of values that indicate the citizenship status
	PrincipalCitizenshipStatus string `json:"principal_citizenship_status,nullable"`
	// One of a standard set of values that indicate customer occupation
	PrincipalOccupation string `json:"principal_occupation,nullable"`
	// Tax ID
	TaxID string `json:"tax_id,nullable"`
	// Company website
	Website string `json:"website,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address                    respjson.Field
		IncorporationDate          respjson.Field
		Industry                   respjson.Field
		LegalFirstname             respjson.Field
		LegalLastname              respjson.Field
		PrincipalAddress           respjson.Field
		PrincipalCitizenshipStatus respjson.Field
		PrincipalOccupation        respjson.Field
		TaxID                      respjson.Field
		Website                    respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetResponseMerchant) RawJSON() string { return r.JSON.raw }
func (r *UserGetResponseMerchant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User field
type UserGetResponseUser struct {
	// User id
	ID string `json:"id"`
	// Business name
	BusinessName string `json:"business_name,nullable"`
	// Number to be used for the next check
	CheckNumber int64 `json:"check_number,nullable"`
	// User's date of birth
	Dob string `json:"dob,nullable"`
	// User first name
	FirstName string `json:"first_name,nullable"`
	// Number to be used for the next invoice
	InvoiceNumber int64 `json:"invoice_number,nullable"`
	// User last name
	LastName string `json:"last_name,nullable"`
	// Phone number
	Phone string `json:"phone,nullable"`
	// Social Security Number
	Ssn string `json:"ssn,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		BusinessName  respjson.Field
		CheckNumber   respjson.Field
		Dob           respjson.Field
		FirstName     respjson.Field
		InvoiceNumber respjson.Field
		LastName      respjson.Field
		Phone         respjson.Field
		Ssn           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetResponseUser) RawJSON() string { return r.JSON.raw }
func (r *UserGetResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response fields for check query
type UserListResponse struct {
	// List of users
	Users []UserListResponseUser `json:"users,required"`
	// Current page
	Page int64 `json:"page"`
	// Total number of pages
	Pages int64 `json:"pages"`
	// Total number of users
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Users       respjson.Field
		Page        respjson.Field
		Pages       respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Check field
type UserListResponseUser struct {
	// Unique identifier for user
	ID string `json:"id,required"`
	// User creation timestamp
	Date string `json:"date,required"`
	// Publishable key of the user
	Key string `json:"key,required"`
	// Name of the user
	Name string `json:"name,required"`
	// Redacted secret key of the user
	Secret string `json:"secret,required"`
	// Specific user_id provided by the marketplace owner
	UserID string `json:"user_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Date        respjson.Field
		Key         respjson.Field
		Name        respjson.Field
		Secret      respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserListResponseUser) RawJSON() string { return r.JSON.raw }
func (r *UserListResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewParams struct {
	// Name of user
	Name string `json:"name,required"`
	// Unique identifier for new user
	UserID string `json:"user_id,required"`
	paramObj
}

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserUpdateParams struct {
	// Bank fields
	Bank UserUpdateParamsBank `json:"bank,omitzero"`
	// Brand fields
	Brand UserUpdateParamsBrand `json:"brand,omitzero"`
	// Developer fields
	Developer UserUpdateParamsDeveloper `json:"developer,omitzero"`
	// Merchant fields
	Merchant UserUpdateParamsMerchant `json:"merchant,omitzero"`
	// Payment fields
	Payment UserUpdateParamsPayment `json:"payment,omitzero"`
	// User field
	User UserUpdateParamsUser `json:"user,omitzero"`
	paramObj
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow UserUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bank fields
type UserUpdateParamsBank struct {
	// Billing bank account id (if user has multiple bank accounts)
	Billing param.Opt[string] `json:"billing,omitzero" format:"uuid"`
	// Default bank account id (if user has multiple bank accounts)
	Default param.Opt[string] `json:"default,omitzero" format:"uuid"`
	paramObj
}

func (r UserUpdateParamsBank) MarshalJSON() (data []byte, err error) {
	type shadow UserUpdateParamsBank
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUpdateParamsBank) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Brand fields
type UserUpdateParamsBrand struct {
	// Email footer
	Footer param.Opt[string] `json:"footer,omitzero"`
	// Email slogan
	Slogan param.Opt[string] `json:"slogan,omitzero"`
	// Base64 encoded company logo
	Logo param.Opt[string] `json:"logo,omitzero"`
	// Email reply to address
	ReplyTo param.Opt[string] `json:"reply_to,omitzero"`
	// Payment fields
	Button UserUpdateParamsBrandButton `json:"button,omitzero"`
	// Payment fields
	Check UserUpdateParamsBrandCheck `json:"check,omitzero"`
	paramObj
}

func (r UserUpdateParamsBrand) MarshalJSON() (data []byte, err error) {
	type shadow UserUpdateParamsBrand
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUpdateParamsBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payment fields
type UserUpdateParamsBrandButton struct {
	// Payment fields
	Color ColorParam `json:"color,omitzero"`
	paramObj
}

func (r UserUpdateParamsBrandButton) MarshalJSON() (data []byte, err error) {
	type shadow UserUpdateParamsBrandButton
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUpdateParamsBrandButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payment fields
type UserUpdateParamsBrandCheck struct {
	Logo param.Opt[string] `json:"logo,omitzero"`
	// Payment fields
	Color ColorParam `json:"color,omitzero"`
	paramObj
}

func (r UserUpdateParamsBrandCheck) MarshalJSON() (data []byte, err error) {
	type shadow UserUpdateParamsBrandCheck
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUpdateParamsBrandCheck) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Developer fields
type UserUpdateParamsDeveloper struct {
	// Webhook URI
	WebhookUri param.Opt[string] `json:"webhook_uri,omitzero"`
	paramObj
}

func (r UserUpdateParamsDeveloper) MarshalJSON() (data []byte, err error) {
	type shadow UserUpdateParamsDeveloper
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUpdateParamsDeveloper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Merchant fields
type UserUpdateParamsMerchant struct {
	// Business incorporation or formation date
	IncorporationDate param.Opt[time.Time] `json:"incorporation_date,omitzero" format:"date"`
	// Principal officer first name
	LegalFirstname param.Opt[string] `json:"legal_firstname,omitzero"`
	// Principal officer last name
	LegalLastname param.Opt[string] `json:"legal_lastname,omitzero"`
	// Occupation of principal
	PrincipalOccupation param.Opt[string] `json:"principal_occupation,omitzero"`
	// Tax ID
	TaxID param.Opt[string] `json:"tax_id,omitzero"`
	// Company website
	Website param.Opt[string] `json:"website,omitzero"`
	// Address fields
	Address UpdateMerchantAddressParam `json:"address,omitzero"`
	// Industry sector
	//
	// Any of 11, 21, 22, 23, 33, 42, 45, 49, 51, 52, 53, 54, 55, 56, 61, 62, 71, 72,
	// 81, 92, 441, 519, 524, 624, 4541, 5411, 5614, 54112, 541211, 541214, 541614.
	Industry int64 `json:"industry,omitzero"`
	// Address fields
	PrincipalAddress UpdateMerchantAddressParam `json:"principal_address,omitzero"`
	// One of a standard set of values that indicate the citizenship status
	//
	// Any of "us_citizen", "resident", "non_resident".
	PrincipalCitizenshipStatus string `json:"principal_citizenship_status,omitzero"`
	paramObj
}

func (r UserUpdateParamsMerchant) MarshalJSON() (data []byte, err error) {
	type shadow UserUpdateParamsMerchant
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUpdateParamsMerchant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UserUpdateParamsMerchant](
		"industry", 11, 21, 22, 23, 33, 42, 45, 49, 51, 52, 53, 54, 55, 56, 61, 62, 71, 72, 81, 92, 441, 519, 524, 624, 4541, 5411, 5614, 54112, 541211, 541214, 541614,
	)
	apijson.RegisterFieldValidator[UserUpdateParamsMerchant](
		"principal_citizenship_status", "us_citizen", "resident", "non_resident",
	)
}

// Payment fields
type UserUpdateParamsPayment struct {
	// Payment expiration
	Expiration param.Opt[int64] `json:"expiration,omitzero"`
	paramObj
}

func (r UserUpdateParamsPayment) MarshalJSON() (data []byte, err error) {
	type shadow UserUpdateParamsPayment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUpdateParamsPayment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User field
type UserUpdateParamsUser struct {
	// Business name
	BusinessName param.Opt[string] `json:"business_name,omitzero"`
	// Number to be used for the next payment
	CheckNumber param.Opt[int64] `json:"check_number,omitzero"`
	// User's date of birth
	Dob param.Opt[time.Time] `json:"dob,omitzero" format:"date"`
	// User first name
	FirstName param.Opt[string] `json:"first_name,omitzero"`
	// Number to be used for the next invoice
	InvoiceNumber param.Opt[int64] `json:"invoice_number,omitzero"`
	// User last name
	LastName param.Opt[string] `json:"last_name,omitzero"`
	// User password
	Password param.Opt[string] `json:"password,omitzero"`
	// Phone number
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Social security number
	Ssn param.Opt[string] `json:"ssn,omitzero"`
	paramObj
}

func (r UserUpdateParamsUser) MarshalJSON() (data []byte, err error) {
	type shadow UserUpdateParamsUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUpdateParamsUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserListParams struct {
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Query
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Items per page
	//
	// Any of 10, 25, 50, 100, 250.
	PerPage int64 `query:"per_page,omitzero" json:"-"`
	// Sort
	//
	// Any of "+DATE", "-DATE", "+USER_ID", "-USER_ID".
	Sort UserListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserListParams]'s query parameters as `url.Values`.
func (r UserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort
type UserListParamsSort string

const (
	UserListParamsSortPlusDate    UserListParamsSort = "+DATE"
	UserListParamsSortMinusDate   UserListParamsSort = "-DATE"
	UserListParamsSortPlusUserID  UserListParamsSort = "+USER_ID"
	UserListParamsSortMinusUserID UserListParamsSort = "-USER_ID"
)

type UserAddSignatureParams struct {
	// Base64 encoded image of user’s signature
	Signature string `json:"signature,required"`
	paramObj
}

func (r UserAddSignatureParams) MarshalJSON() (data []byte, err error) {
	type shadow UserAddSignatureParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserAddSignatureParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
