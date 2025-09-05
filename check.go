// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package checkbook

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/checkbook-go/internal/apijson"
	"github.com/stainless-sdks/checkbook-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/checkbook-go/internal/encoding/json"
	"github.com/stainless-sdks/checkbook-go/internal/requestconfig"
	"github.com/stainless-sdks/checkbook-go/option"
	"github.com/stainless-sdks/checkbook-go/packages/param"
	"github.com/stainless-sdks/checkbook-go/packages/respjson"
)

// CheckService contains methods and other services that help with interacting with
// the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCheckService] method instead.
type CheckService struct {
	Options []option.RequestOption
	Deposit CheckDepositService
}

// NewCheckService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCheckService(opts ...option.RequestOption) (r CheckService) {
	r = CheckService{}
	r.Options = opts
	r.Deposit = NewCheckDepositService(opts...)
	return
}

// Get the specified payment
func (r *CheckService) Get(ctx context.Context, checkID string, opts ...option.RequestOption) (res *GetCheck, err error) {
	opts = append(r.Options[:], opts...)
	if checkID == "" {
		err = errors.New("missing required check_id parameter")
		return
	}
	path := fmt.Sprintf("v3/check/%s", checkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Return the sent/received payments
func (r *CheckService) List(ctx context.Context, query CheckListParams, opts ...option.RequestOption) (res *CheckListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/check"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Create a digital payment
func (r *CheckService) NewDigital(ctx context.Context, body CheckNewDigitalParams, opts ...option.RequestOption) (res *GetCheck, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/check/digital"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new multi party payment
func (r *CheckService) NewMulti(ctx context.Context, body CheckNewMultiParams, opts ...option.RequestOption) (res *GetCheck, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/check/multi"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new paper check
func (r *CheckService) NewPhysical(ctx context.Context, body CheckNewPhysicalParams, opts ...option.RequestOption) (res *GetCheck, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/check/physical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Endorse a multi party payment
func (r *CheckService) Endorse(ctx context.Context, checkID string, body CheckEndorseParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if checkID == "" {
		err = errors.New("missing required check_id parameter")
		return
	}
	path := fmt.Sprintf("v3/check/endorse/%s", checkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get the attachment for a payment
func (r *CheckService) GetAttachment(ctx context.Context, checkID string, opts ...option.RequestOption) (res *Error, err error) {
	opts = append(r.Options[:], opts...)
	if checkID == "" {
		err = errors.New("missing required check_id parameter")
		return
	}
	path := fmt.Sprintf("v3/check/%s/attachment", checkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get details on a failed payment
func (r *CheckService) GetFailDetails(ctx context.Context, checkID string, opts ...option.RequestOption) (res *CheckGetFailDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if checkID == "" {
		err = errors.New("missing required check_id parameter")
		return
	}
	path := fmt.Sprintf("v3/check/%s/fail", checkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get tracking details on a mailed check
func (r *CheckService) GetTrackingDetails(ctx context.Context, checkID string, opts ...option.RequestOption) (res *CheckGetTrackingDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if checkID == "" {
		err = errors.New("missing required check_id parameter")
		return
	}
	path := fmt.Sprintf("v3/check/%s/tracking", checkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the verification code
func (r *CheckService) GetVerificationCode(ctx context.Context, checkID string, opts ...option.RequestOption) (res *Error, err error) {
	opts = append(r.Options[:], opts...)
	if checkID == "" {
		err = errors.New("missing required check_id parameter")
		return
	}
	path := fmt.Sprintf("v3/check/%s/verification", checkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Resend payment notification
func (r *CheckService) Notify(ctx context.Context, checkID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if checkID == "" {
		err = errors.New("missing required check_id parameter")
		return
	}
	path := fmt.Sprintf("v3/check/notify/%s", checkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Preview a new payment
func (r *CheckService) Preview(ctx context.Context, body CheckPreviewParams, opts ...option.RequestOption) (res *CheckPreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/check/preview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Print a check
func (r *CheckService) Print(ctx context.Context, checkID string, opts ...option.RequestOption) (res *Error, err error) {
	opts = append(r.Options[:], opts...)
	if checkID == "" {
		err = errors.New("missing required check_id parameter")
		return
	}
	path := fmt.Sprintf("v3/check/print/%s", checkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Trigger a webhook notification on sandbox
func (r *CheckService) TriggerWebhook(ctx context.Context, checkID string, body CheckTriggerWebhookParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if checkID == "" {
		err = errors.New("missing required check_id parameter")
		return
	}
	path := fmt.Sprintf("v3/check/webhook/%s", checkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Void the specified payment
func (r *CheckService) Void(ctx context.Context, checkID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if checkID == "" {
		err = errors.New("missing required check_id parameter")
		return
	}
	path := fmt.Sprintf("v3/check/%s", checkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type CheckListResponse struct {
	// List of sent/received payments
	Checks []CheckListResponseCheck `json:"checks,required"`
	Page   int64                    `json:"page,required"`
	Pages  int64                    `json:"pages,required"`
	Total  int64                    `json:"total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Checks      respjson.Field
		Page        respjson.Field
		Pages       respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckListResponse) RawJSON() string { return r.JSON.raw }
func (r *CheckListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckListResponseCheck struct {
	// Unique identifier for payment
	ID string `json:"id,required"`
	// Payment amount
	Amount float64 `json:"amount,required"`
	// Payment creation timestamp
	Date string `json:"date,required"`
	// Payment direction
	//
	// Any of "INCOMING", "OUTGOING".
	Direction string `json:"direction,required"`
	// Current status of the payment
	//
	// Any of "PAID", "IN_PROCESS", "UNPAID", "VOID", "EXPIRED", "PRINTED", "MAILED",
	// "FAILED", "REFUNDED".
	Status string `json:"status,required"`
	// Payment comment
	Comment string `json:"comment,nullable"`
	// Payment description/memo
	Description string `json:"description,nullable"`
	// URI where image of the payment can be accessed
	ImageUri string `json:"image_uri,nullable"`
	// Name of third party who received the payment
	Name      string                               `json:"name,nullable"`
	Number    CheckListResponseCheckNumberUnion    `json:"number"`
	Recipient CheckListResponseCheckRecipientUnion `json:"recipient"`
	// Email/id or physical address of the payment sender
	Sender string `json:"sender,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		Date        respjson.Field
		Direction   respjson.Field
		Status      respjson.Field
		Comment     respjson.Field
		Description respjson.Field
		ImageUri    respjson.Field
		Name        respjson.Field
		Number      respjson.Field
		Recipient   respjson.Field
		Sender      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckListResponseCheck) RawJSON() string { return r.JSON.raw }
func (r *CheckListResponseCheck) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CheckListResponseCheckNumberUnion contains all possible properties and values
// from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type CheckListResponseCheckNumberUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u CheckListResponseCheckNumberUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CheckListResponseCheckNumberUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CheckListResponseCheckNumberUnion) RawJSON() string { return u.JSON.raw }

func (r *CheckListResponseCheckNumberUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CheckListResponseCheckRecipientUnion contains all possible properties and values
// from [string], [[]string], [CheckAddress].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type CheckListResponseCheckRecipientUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field is from variant [CheckAddress].
	City string `json:"city"`
	// This field is from variant [CheckAddress].
	Line1 string `json:"line_1"`
	// This field is from variant [CheckAddress].
	State string `json:"state"`
	// This field is from variant [CheckAddress].
	Zip string `json:"zip"`
	// This field is from variant [CheckAddress].
	Country string `json:"country"`
	// This field is from variant [CheckAddress].
	Line2 string `json:"line_2"`
	JSON  struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		City          respjson.Field
		Line1         respjson.Field
		State         respjson.Field
		Zip           respjson.Field
		Country       respjson.Field
		Line2         respjson.Field
		raw           string
	} `json:"-"`
}

func (u CheckListResponseCheckRecipientUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CheckListResponseCheckRecipientUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CheckListResponseCheckRecipientUnion) AsCheckAddress() (v CheckAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CheckListResponseCheckRecipientUnion) RawJSON() string { return u.JSON.raw }

func (r *CheckListResponseCheckRecipientUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckGetFailDetailsResponse struct {
	// Check fail code
	Code string `json:"code,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckGetFailDetailsResponse) RawJSON() string { return r.JSON.raw }
func (r *CheckGetFailDetailsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckGetTrackingDetailsResponse struct {
	// List of tracking events
	Tracking []CheckGetTrackingDetailsResponseTracking `json:"tracking,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tracking    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckGetTrackingDetailsResponse) RawJSON() string { return r.JSON.raw }
func (r *CheckGetTrackingDetailsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckGetTrackingDetailsResponseTracking struct {
	// Location
	Location string `json:"location,required"`
	// Date
	ActionTs string `json:"action_ts,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Location    respjson.Field
		ActionTs    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckGetTrackingDetailsResponseTracking) RawJSON() string { return r.JSON.raw }
func (r *CheckGetTrackingDetailsResponseTracking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckPreviewResponse struct {
	// Base64 encoded check image
	Image string `json:"image,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckPreviewResponse) RawJSON() string { return r.JSON.raw }
func (r *CheckPreviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckListParams struct {
	// End date
	EndDate param.Opt[time.Time] `query:"end_date,omitzero" format:"date" json:"-"`
	// Start date
	StartDate param.Opt[time.Time] `query:"start_date,omitzero" format:"date" json:"-"`
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Query
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Direction
	//
	// Any of "INCOMING", "OUTGOING".
	Direction CheckListParamsDirection `query:"direction,omitzero" json:"-"`
	// Items per page
	//
	// Any of 10, 25, 50, 100, 250.
	PerPage int64 `query:"per_page,omitzero" json:"-"`
	// Sort
	//
	// Any of "+NUMBER", "-NUMBER", "+TYPE", "-TYPE", "+AMOUNT", "-AMOUNT", "+STATUS",
	// "-STATUS", "+DATE", "-DATE", "+UPDATE", "-UPDATE", "+DESCRIPTION",
	// "-DESCRIPTION".
	Sort CheckListParamsSort `query:"sort,omitzero" json:"-"`
	// Status
	//
	// Any of "PAID", "IN_PROCESS", "UNPAID", "VOID", "EXPIRED", "PRINTED", "MAILED",
	// "FAILED", "REFUNDED".
	Status CheckListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CheckListParams]'s query parameters as `url.Values`.
func (r CheckListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction
type CheckListParamsDirection string

const (
	CheckListParamsDirectionIncoming CheckListParamsDirection = "INCOMING"
	CheckListParamsDirectionOutgoing CheckListParamsDirection = "OUTGOING"
)

// Sort
type CheckListParamsSort string

const (
	CheckListParamsSortPlusNumber       CheckListParamsSort = "+NUMBER"
	CheckListParamsSortMinusNumber      CheckListParamsSort = "-NUMBER"
	CheckListParamsSortPlusType         CheckListParamsSort = "+TYPE"
	CheckListParamsSortMinusType        CheckListParamsSort = "-TYPE"
	CheckListParamsSortPlusAmount       CheckListParamsSort = "+AMOUNT"
	CheckListParamsSortMinusAmount      CheckListParamsSort = "-AMOUNT"
	CheckListParamsSortPlusStatus       CheckListParamsSort = "+STATUS"
	CheckListParamsSortMinusStatus      CheckListParamsSort = "-STATUS"
	CheckListParamsSortPlusDate         CheckListParamsSort = "+DATE"
	CheckListParamsSortMinusDate        CheckListParamsSort = "-DATE"
	CheckListParamsSortPlusUpdate       CheckListParamsSort = "+UPDATE"
	CheckListParamsSortMinusUpdate      CheckListParamsSort = "-UPDATE"
	CheckListParamsSortPlusDescription  CheckListParamsSort = "+DESCRIPTION"
	CheckListParamsSortMinusDescription CheckListParamsSort = "-DESCRIPTION"
)

// Status
type CheckListParamsStatus string

const (
	CheckListParamsStatusPaid      CheckListParamsStatus = "PAID"
	CheckListParamsStatusInProcess CheckListParamsStatus = "IN_PROCESS"
	CheckListParamsStatusUnpaid    CheckListParamsStatus = "UNPAID"
	CheckListParamsStatusVoid      CheckListParamsStatus = "VOID"
	CheckListParamsStatusExpired   CheckListParamsStatus = "EXPIRED"
	CheckListParamsStatusPrinted   CheckListParamsStatus = "PRINTED"
	CheckListParamsStatusMailed    CheckListParamsStatus = "MAILED"
	CheckListParamsStatusFailed    CheckListParamsStatus = "FAILED"
	CheckListParamsStatusRefunded  CheckListParamsStatus = "REFUNDED"
)

type CheckNewDigitalParams struct {
	CreateDigitalCheck CreateDigitalCheckParam
	paramObj
}

func (r CheckNewDigitalParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateDigitalCheck)
}
func (r *CheckNewDigitalParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateDigitalCheck)
}

type CheckNewMultiParams struct {
	CreateMultiCheck CreateMultiCheckParam
	paramObj
}

func (r CheckNewMultiParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateMultiCheck)
}
func (r *CheckNewMultiParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateMultiCheck)
}

type CheckNewPhysicalParams struct {
	CreatePhysicalCheck CreatePhysicalCheckParam
	paramObj
}

func (r CheckNewPhysicalParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreatePhysicalCheck)
}
func (r *CheckNewPhysicalParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreatePhysicalCheck)
}

type CheckEndorseParams struct {
	// Name of the endorser
	Name string `json:"name,required"`
	// Signature of the endorser
	Signature param.Opt[string] `json:"signature,omitzero"`
	paramObj
}

func (r CheckEndorseParams) MarshalJSON() (data []byte, err error) {
	type shadow CheckEndorseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CheckEndorseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckPreviewParams struct {
	// Check amount
	Amount float64 `json:"amount,required"`
	// Recipient name
	Name string `json:"name,required"`
	// Debit account id for funds (if sender has multiple bank accounts)
	Account param.Opt[string] `json:"account,omitzero" format:"uuid"`
	// Check description
	Description param.Opt[string]             `json:"description,omitzero"`
	Number      CheckPreviewParamsNumberUnion `json:"number,omitzero"`
	paramObj
}

func (r CheckPreviewParams) MarshalJSON() (data []byte, err error) {
	type shadow CheckPreviewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CheckPreviewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CheckPreviewParamsNumberUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u CheckPreviewParamsNumberUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *CheckPreviewParamsNumberUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CheckPreviewParamsNumberUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

type CheckTriggerWebhookParams struct {
	// Desired check status
	//
	// Any of "PAID", "IN_PROCESS", "UNPAID", "VOID", "EXPIRED", "PRINTED", "MAILED",
	// "FAILED", "REFUNDED".
	Status  CheckTriggerWebhookParamsStatus  `json:"status,omitzero,required"`
	Options CheckTriggerWebhookParamsOptions `json:"options,omitzero"`
	paramObj
}

func (r CheckTriggerWebhookParams) MarshalJSON() (data []byte, err error) {
	type shadow CheckTriggerWebhookParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CheckTriggerWebhookParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Desired check status
type CheckTriggerWebhookParamsStatus string

const (
	CheckTriggerWebhookParamsStatusPaid      CheckTriggerWebhookParamsStatus = "PAID"
	CheckTriggerWebhookParamsStatusInProcess CheckTriggerWebhookParamsStatus = "IN_PROCESS"
	CheckTriggerWebhookParamsStatusUnpaid    CheckTriggerWebhookParamsStatus = "UNPAID"
	CheckTriggerWebhookParamsStatusVoid      CheckTriggerWebhookParamsStatus = "VOID"
	CheckTriggerWebhookParamsStatusExpired   CheckTriggerWebhookParamsStatus = "EXPIRED"
	CheckTriggerWebhookParamsStatusPrinted   CheckTriggerWebhookParamsStatus = "PRINTED"
	CheckTriggerWebhookParamsStatusMailed    CheckTriggerWebhookParamsStatus = "MAILED"
	CheckTriggerWebhookParamsStatusFailed    CheckTriggerWebhookParamsStatus = "FAILED"
	CheckTriggerWebhookParamsStatusRefunded  CheckTriggerWebhookParamsStatus = "REFUNDED"
)

type CheckTriggerWebhookParamsOptions struct {
	// Desired return code
	//
	// Any of "R01", "R02", "R03", "A", "B", "C", "D", "04", "05", "06", "R901".
	ReturnCode string `json:"return_code,omitzero"`
	paramObj
}

func (r CheckTriggerWebhookParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow CheckTriggerWebhookParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CheckTriggerWebhookParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CheckTriggerWebhookParamsOptions](
		"return_code", "R01", "R02", "R03", "A", "B", "C", "D", "04", "05", "06", "R901",
	)
}
