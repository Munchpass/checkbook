// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package checkbook

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/checkbook-go/internal/apijson"
	"github.com/stainless-sdks/checkbook-go/internal/apiquery"
	"github.com/stainless-sdks/checkbook-go/internal/requestconfig"
	"github.com/stainless-sdks/checkbook-go/option"
	"github.com/stainless-sdks/checkbook-go/packages/param"
	"github.com/stainless-sdks/checkbook-go/packages/respjson"
)

// DirectoryService contains methods and other services that help with interacting
// with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDirectoryService] method instead.
type DirectoryService struct {
	Options []option.RequestOption
	Account DirectoryAccountService
}

// NewDirectoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDirectoryService(opts ...option.RequestOption) (r DirectoryService) {
	r = DirectoryService{}
	r.Options = opts
	r.Account = NewDirectoryAccountService(opts...)
	return
}

// Create a new directory item
func (r *DirectoryService) New(ctx context.Context, body DirectoryNewParams, opts ...option.RequestOption) (res *CreateDirectoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/directory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Return the directory entry
func (r *DirectoryService) Get(ctx context.Context, query DirectoryGetParams, opts ...option.RequestOption) (res *DirectoryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/directory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update a directory item
func (r *DirectoryService) Update(ctx context.Context, directoryID string, body DirectoryUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if directoryID == "" {
		err = errors.New("missing required directory_id parameter")
		return
	}
	path := fmt.Sprintf("v3/directory/%s", directoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Remove the directory item
func (r *DirectoryService) Delete(ctx context.Context, directoryID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if directoryID == "" {
		err = errors.New("missing required directory_id parameter")
		return
	}
	path := fmt.Sprintf("v3/directory/%s", directoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type CreateDirectoryResponse struct {
	// Unique identifier for directory
	ID string `json:"id,required"`
	// Name
	Name     string                           `json:"name,required"`
	Accounts []CreateDirectoryResponseAccount `json:"accounts"`
	Address  CreateDirectoryResponseAddress   `json:"address"`
	// Email
	Email string `json:"email,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Accounts    respjson.Field
		Address     respjson.Field
		Email       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateDirectoryResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateDirectoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreateDirectoryResponseAccount struct {
	// Unique identifier for account
	ID string `json:"id,required"`
	// Account name
	Name string `json:"name,required"`
	// Last 4 of account number
	Number string `json:"number,required"`
	// Account type
	//
	// Any of "CARD", "BANK".
	Type string `json:"type,required"`
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
func (r CreateDirectoryResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *CreateDirectoryResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreateDirectoryResponseAddress struct {
	// City
	City string `json:"city,nullable"`
	// Country
	Country string `json:"country,nullable"`
	// Line 1
	Line1 string `json:"line_1,nullable"`
	// Line 2
	Line2 string `json:"line_2,nullable"`
	// State
	State string `json:"state,nullable"`
	// Zip code
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
func (r CreateDirectoryResponseAddress) RawJSON() string { return r.JSON.raw }
func (r *CreateDirectoryResponseAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryGetResponse struct {
	// List of directory entries
	Entries []CreateDirectoryResponse `json:"entries,required"`
	// Current page
	Page int64 `json:"page"`
	// Total number of pages
	Pages int64 `json:"pages"`
	// Total number of directory entries
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entries     respjson.Field
		Page        respjson.Field
		Pages       respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirectoryGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DirectoryGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryNewParams struct {
	// Name
	Name string `json:"name,required"`
	// Email
	Email   param.Opt[string]         `json:"email,omitzero"`
	Address DirectoryNewParamsAddress `json:"address,omitzero"`
	paramObj
}

func (r DirectoryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DirectoryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirectoryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties City, Line1, State, Zip are required.
type DirectoryNewParamsAddress struct {
	// City
	City string `json:"city,required"`
	// Line 1
	Line1 string `json:"line_1,required"`
	// State
	State string `json:"state,required"`
	// Zip/postal code
	Zip string `json:"zip,required"`
	// Country
	Country param.Opt[string] `json:"country,omitzero"`
	// Line 2
	Line2 param.Opt[string] `json:"line_2,omitzero"`
	paramObj
}

func (r DirectoryNewParamsAddress) MarshalJSON() (data []byte, err error) {
	type shadow DirectoryNewParamsAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirectoryNewParamsAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryGetParams struct {
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Query
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Items per page
	//
	// Any of 10, 25, 50.
	PerPage int64 `query:"per_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DirectoryGetParams]'s query parameters as `url.Values`.
func (r DirectoryGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DirectoryUpdateParams struct {
	// Email
	Email param.Opt[string] `json:"email,omitzero"`
	// Name
	Name    param.Opt[string]            `json:"name,omitzero"`
	Address DirectoryUpdateParamsAddress `json:"address,omitzero"`
	paramObj
}

func (r DirectoryUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DirectoryUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirectoryUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirectoryUpdateParamsAddress struct {
	// Country
	Country param.Opt[string] `json:"country,omitzero"`
	// Line 2
	Line2 param.Opt[string] `json:"line_2,omitzero"`
	// City
	City param.Opt[string] `json:"city,omitzero"`
	// Line 1
	Line1 param.Opt[string] `json:"line_1,omitzero"`
	// State
	State param.Opt[string] `json:"state,omitzero"`
	// Zip/postal code
	Zip param.Opt[string] `json:"zip,omitzero"`
	paramObj
}

func (r DirectoryUpdateParamsAddress) MarshalJSON() (data []byte, err error) {
	type shadow DirectoryUpdateParamsAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirectoryUpdateParamsAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
