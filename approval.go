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

// ApprovalService contains methods and other services that help with interacting
// with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApprovalService] method instead.
type ApprovalService struct {
	Options []option.RequestOption
}

// NewApprovalService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewApprovalService(opts ...option.RequestOption) (r ApprovalService) {
	r = ApprovalService{}
	r.Options = opts
	return
}

// Get the specified payment approval
func (r *ApprovalService) Get(ctx context.Context, approvalID string, opts ...option.RequestOption) (res *GetApproval, err error) {
	opts = append(r.Options[:], opts...)
	if approvalID == "" {
		err = errors.New("missing required approval_id parameter")
		return
	}
	path := fmt.Sprintf("v3/approval/%s", approvalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the specified paynent approval
func (r *ApprovalService) Update(ctx context.Context, approvalID string, body ApprovalUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if approvalID == "" {
		err = errors.New("missing required approval_id parameter")
		return
	}
	path := fmt.Sprintf("v3/approval/%s", approvalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Return approvals
func (r *ApprovalService) List(ctx context.Context, query ApprovalListParams, opts ...option.RequestOption) (res *ApprovalListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/approval"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Cancel the specified check approval
func (r *ApprovalService) Delete(ctx context.Context, approvalID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if approvalID == "" {
		err = errors.New("missing required approval_id parameter")
		return
	}
	path := fmt.Sprintf("v3/approval/%s", approvalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Create a new approval digital payment
func (r *ApprovalService) NewDigital(ctx context.Context, body ApprovalNewDigitalParams, opts ...option.RequestOption) (res *GetApproval, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/approval/digital"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new multi-party payment approval
func (r *ApprovalService) NewMulti(ctx context.Context, body ApprovalNewMultiParams, opts ...option.RequestOption) (res *GetApproval, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/approval/multi"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a new physical check approval
func (r *ApprovalService) NewPhysical(ctx context.Context, body ApprovalNewPhysicalParams, opts ...option.RequestOption) (res *GetApproval, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/approval/physical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a live payment from an approval
func (r *ApprovalService) Release(ctx context.Context, body ApprovalReleaseParams, opts ...option.RequestOption) (res *GetCheck, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/approval/release"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the attachment for a payment approval
func (r *ApprovalService) GetAttachment(ctx context.Context, approvalID string, opts ...option.RequestOption) (res *Error, err error) {
	opts = append(r.Options[:], opts...)
	if approvalID == "" {
		err = errors.New("missing required approval_id parameter")
		return
	}
	path := fmt.Sprintf("v3/approval/%s/attachment", approvalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// The properties Content, Filename are required.
type AttachmentParam struct {
	// Base64 encoded attachment content
	Content string `json:"content,required"`
	// Attachment filename
	Filename string `json:"filename,required"`
	paramObj
}

func (r AttachmentParam) MarshalJSON() (data []byte, err error) {
	type shadow AttachmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttachmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckAddress struct {
	// City
	City string `json:"city,required"`
	// Line 1
	Line1 string `json:"line_1,required"`
	// State
	State string `json:"state,required"`
	// Zip/postal code
	Zip string `json:"zip,required"`
	// Country
	Country string `json:"country"`
	// Line 2
	Line2 string `json:"line_2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Line1       respjson.Field
		State       respjson.Field
		Zip         respjson.Field
		Country     respjson.Field
		Line2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckAddress) RawJSON() string { return r.JSON.raw }
func (r *CheckAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, Name, Recipient are required.
type CreateDigitalCheckParam struct {
	// Payment amount
	Amount float64 `json:"amount,required"`
	// Name of recipient
	Name      string                                `json:"name,required"`
	Recipient CreateDigitalCheckRecipientUnionParam `json:"recipient,omitzero,required"`
	// Optional description/memo for payment
	Description param.Opt[string] `json:"description,omitzero"`
	// Debit account id for funds (if sender has multiple bank accounts)
	Account param.Opt[string] `json:"account,omitzero" format:"uuid"`
	// Comment field for payment
	Comment    param.Opt[string]                      `json:"comment,omitzero"`
	Attachment CreateDigitalCheckAttachmentUnionParam `json:"attachment,omitzero"`
	// Any of "PRINT", "MAIL", "BANK", "CARD", "VCC", "RTP", "PAYPAL", "WALLET",
	// "VENMO", "INTERAC", "WIRE".
	DepositOptions []string                           `json:"deposit_options,omitzero"`
	Number         CreateDigitalCheckNumberUnionParam `json:"number,omitzero"`
	Pin            PinParam                           `json:"pin,omitzero"`
	// List of the remittance records
	RemittanceAdvice []RemittanceAdviceParam `json:"remittance_advice,omitzero"`
	paramObj
}

func (r CreateDigitalCheckParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateDigitalCheckParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateDigitalCheckParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CreateDigitalCheckRecipientUnionParam struct {
	OfString           param.Opt[string]                                 `json:",omitzero,inline"`
	OfDigitalRecipient *CreateDigitalCheckRecipientDigitalRecipientParam `json:",omitzero,inline"`
	paramUnion
}

func (u CreateDigitalCheckRecipientUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfDigitalRecipient)
}
func (u *CreateDigitalCheckRecipientUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CreateDigitalCheckRecipientUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfDigitalRecipient) {
		return u.OfDigitalRecipient
	}
	return nil
}

// The property ID is required.
type CreateDigitalCheckRecipientDigitalRecipientParam struct {
	// Recipient ID
	ID string `json:"id,required"`
	// Recipient account
	Account param.Opt[string] `json:"account,omitzero" format:"uuid"`
	paramObj
}

func (r CreateDigitalCheckRecipientDigitalRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateDigitalCheckRecipientDigitalRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateDigitalCheckRecipientDigitalRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CreateDigitalCheckAttachmentUnionParam struct {
	OfAttachment *AttachmentParam  `json:",omitzero,inline"`
	OfString     param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u CreateDigitalCheckAttachmentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAttachment, u.OfString)
}
func (u *CreateDigitalCheckAttachmentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CreateDigitalCheckAttachmentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAttachment) {
		return u.OfAttachment
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CreateDigitalCheckNumberUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u CreateDigitalCheckNumberUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *CreateDigitalCheckNumberUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CreateDigitalCheckNumberUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// The properties Amount, Recipients are required.
type CreateMultiCheckParam struct {
	// Payment amount
	Amount float64 `json:"amount,required"`
	// Recipients
	Recipients []CreateMultiCheckRecipientParam `json:"recipients,omitzero,required"`
	// Optional description/memo for payment
	Description param.Opt[string] `json:"description,omitzero"`
	// Debit account id for funds (if sender has multiple bank accounts)
	Account param.Opt[string] `json:"account,omitzero" format:"uuid"`
	// Comment field for payment
	Comment    param.Opt[string]                    `json:"comment,omitzero"`
	Attachment CreateMultiCheckAttachmentUnionParam `json:"attachment,omitzero"`
	// Any of "PRINT", "MAIL", "BANK", "CARD", "VCC", "RTP", "PAYPAL", "WALLET",
	// "VENMO", "INTERAC", "WIRE".
	DepositOptions []string                         `json:"deposit_options,omitzero"`
	Number         CreateMultiCheckNumberUnionParam `json:"number,omitzero"`
	// List of the remittance records
	RemittanceAdvice []RemittanceAdviceParam `json:"remittance_advice,omitzero"`
	paramObj
}

func (r CreateMultiCheckParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateMultiCheckParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateMultiCheckParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Recipient are required.
type CreateMultiCheckRecipientParam struct {
	// Name of recipient
	Name string `json:"name,required"`
	// Email or text enabled phone number/id of recipient
	Recipient   string          `json:"recipient,required"`
	EndorseOnly param.Opt[bool] `json:"endorse_only,omitzero"`
	Pin         PinParam        `json:"pin,omitzero"`
	paramObj
}

func (r CreateMultiCheckRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateMultiCheckRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateMultiCheckRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CreateMultiCheckAttachmentUnionParam struct {
	OfAttachment *AttachmentParam  `json:",omitzero,inline"`
	OfString     param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u CreateMultiCheckAttachmentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAttachment, u.OfString)
}
func (u *CreateMultiCheckAttachmentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CreateMultiCheckAttachmentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAttachment) {
		return u.OfAttachment
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CreateMultiCheckNumberUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u CreateMultiCheckNumberUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *CreateMultiCheckNumberUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CreateMultiCheckNumberUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// The properties Amount, Name, Recipient are required.
type CreatePhysicalCheckParam struct {
	// Payment amount
	Amount float64 `json:"amount,required"`
	// Name of recipient
	Name      string                            `json:"name,required"`
	Recipient CreatePhysicalCheckRecipientParam `json:"recipient,omitzero,required"`
	// Optional description/memo for payment
	Description param.Opt[string] `json:"description,omitzero"`
	// Debit account id for funds (if sender has multiple bank accounts)
	Account param.Opt[string] `json:"account,omitzero" format:"uuid"`
	// Comment field for payment
	Comment    param.Opt[string]                       `json:"comment,omitzero"`
	Attachment CreatePhysicalCheckAttachmentUnionParam `json:"attachment,omitzero"`
	// Delivery options
	//
	// Any of "USPS_FIRST_CLASS", "OVERNIGHT", "TWO_DAY", "USPS_CERTIFIED".
	MailType         CreatePhysicalCheckMailType                   `json:"mail_type,omitzero"`
	Number           CreatePhysicalCheckNumberUnionParam           `json:"number,omitzero"`
	RemittanceAdvice CreatePhysicalCheckRemittanceAdviceUnionParam `json:"remittance_advice,omitzero"`
	paramObj
}

func (r CreatePhysicalCheckParam) MarshalJSON() (data []byte, err error) {
	type shadow CreatePhysicalCheckParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatePhysicalCheckParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties City, Line1, State, Zip are required.
type CreatePhysicalCheckRecipientParam struct {
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
	// Name on envelope
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r CreatePhysicalCheckRecipientParam) MarshalJSON() (data []byte, err error) {
	type shadow CreatePhysicalCheckRecipientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatePhysicalCheckRecipientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CreatePhysicalCheckAttachmentUnionParam struct {
	OfAttachment *AttachmentParam  `json:",omitzero,inline"`
	OfString     param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u CreatePhysicalCheckAttachmentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAttachment, u.OfString)
}
func (u *CreatePhysicalCheckAttachmentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CreatePhysicalCheckAttachmentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAttachment) {
		return u.OfAttachment
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Delivery options
type CreatePhysicalCheckMailType string

const (
	CreatePhysicalCheckMailTypeUspsFirstClass CreatePhysicalCheckMailType = "USPS_FIRST_CLASS"
	CreatePhysicalCheckMailTypeOvernight      CreatePhysicalCheckMailType = "OVERNIGHT"
	CreatePhysicalCheckMailTypeTwoDay         CreatePhysicalCheckMailType = "TWO_DAY"
	CreatePhysicalCheckMailTypeUspsCertified  CreatePhysicalCheckMailType = "USPS_CERTIFIED"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CreatePhysicalCheckNumberUnionParam struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u CreatePhysicalCheckNumberUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *CreatePhysicalCheckNumberUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CreatePhysicalCheckNumberUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CreatePhysicalCheckRemittanceAdviceUnionParam struct {
	OfString           param.Opt[string]       `json:",omitzero,inline"`
	OfRemittanceAdvice []RemittanceAdviceParam `json:",omitzero,inline"`
	paramUnion
}

func (u CreatePhysicalCheckRemittanceAdviceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfRemittanceAdvice)
}
func (u *CreatePhysicalCheckRemittanceAdviceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CreatePhysicalCheckRemittanceAdviceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfRemittanceAdvice) {
		return &u.OfRemittanceAdvice
	}
	return nil
}

type Error struct {
	Message string `json:"message,required"`
	Errors  any    `json:"errors"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Errors      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Error) RawJSON() string { return r.JSON.raw }
func (r *Error) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetApproval struct {
	// Unique identifier for payment
	ID string `json:"id,required"`
	// Payment amount
	Amount float64 `json:"amount,required"`
	// Payment creation timestamp
	Date string `json:"date,required"`
	// Remittance advice
	RemittanceAdvice [][]string `json:"remittance_advice,required"`
	// Current status of the check
	//
	// Any of "UNPAID", "APPROVED", "VOID".
	Status GetApprovalStatus `json:"status,required"`
	// Payment comment
	Comment string `json:"comment,nullable"`
	// Payment description/memo
	Description string `json:"description,nullable"`
	// URI where image of the payment can be accessed
	ImageUri string `json:"image_uri,nullable"`
	// Name of third party who received the payment
	Name      string                    `json:"name,nullable"`
	Number    GetApprovalNumberUnion    `json:"number"`
	Recipient GetApprovalRecipientUnion `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Amount           respjson.Field
		Date             respjson.Field
		RemittanceAdvice respjson.Field
		Status           respjson.Field
		Comment          respjson.Field
		Description      respjson.Field
		ImageUri         respjson.Field
		Name             respjson.Field
		Number           respjson.Field
		Recipient        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetApproval) RawJSON() string { return r.JSON.raw }
func (r *GetApproval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the check
type GetApprovalStatus string

const (
	GetApprovalStatusUnpaid   GetApprovalStatus = "UNPAID"
	GetApprovalStatusApproved GetApprovalStatus = "APPROVED"
	GetApprovalStatusVoid     GetApprovalStatus = "VOID"
)

// GetApprovalNumberUnion contains all possible properties and values from [int64],
// [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type GetApprovalNumberUnion struct {
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

func (u GetApprovalNumberUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GetApprovalNumberUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GetApprovalNumberUnion) RawJSON() string { return u.JSON.raw }

func (r *GetApprovalNumberUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetApprovalRecipientUnion contains all possible properties and values from
// [string], [[]string], [CheckAddress].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type GetApprovalRecipientUnion struct {
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

func (u GetApprovalRecipientUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GetApprovalRecipientUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GetApprovalRecipientUnion) AsCheckAddress() (v CheckAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GetApprovalRecipientUnion) RawJSON() string { return u.JSON.raw }

func (r *GetApprovalRecipientUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetCheck struct {
	// Unique identifier for payment
	ID string `json:"id,required"`
	// Payment amount
	Amount float64 `json:"amount,required"`
	// Payment creation timestamp
	Date string `json:"date,required"`
	// Remittance advice
	RemittanceAdvice [][]string `json:"remittance_advice,required"`
	// Current status of the payment
	//
	// Any of "PAID", "IN_PROCESS", "UNPAID", "VOID", "EXPIRED", "PRINTED", "MAILED",
	// "FAILED", "REFUNDED".
	Status GetCheckStatus `json:"status,required"`
	// Payment comment
	Comment string `json:"comment,nullable"`
	// Payment description/memo
	Description string `json:"description,nullable"`
	// URI where image of the payment can be accessed
	ImageUri string `json:"image_uri,nullable"`
	// Name of third party who received the payment
	Name      string                 `json:"name,nullable"`
	Number    GetCheckNumberUnion    `json:"number"`
	Recipient GetCheckRecipientUnion `json:"recipient"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Amount           respjson.Field
		Date             respjson.Field
		RemittanceAdvice respjson.Field
		Status           respjson.Field
		Comment          respjson.Field
		Description      respjson.Field
		ImageUri         respjson.Field
		Name             respjson.Field
		Number           respjson.Field
		Recipient        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetCheck) RawJSON() string { return r.JSON.raw }
func (r *GetCheck) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the payment
type GetCheckStatus string

const (
	GetCheckStatusPaid      GetCheckStatus = "PAID"
	GetCheckStatusInProcess GetCheckStatus = "IN_PROCESS"
	GetCheckStatusUnpaid    GetCheckStatus = "UNPAID"
	GetCheckStatusVoid      GetCheckStatus = "VOID"
	GetCheckStatusExpired   GetCheckStatus = "EXPIRED"
	GetCheckStatusPrinted   GetCheckStatus = "PRINTED"
	GetCheckStatusMailed    GetCheckStatus = "MAILED"
	GetCheckStatusFailed    GetCheckStatus = "FAILED"
	GetCheckStatusRefunded  GetCheckStatus = "REFUNDED"
)

// GetCheckNumberUnion contains all possible properties and values from [int64],
// [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type GetCheckNumberUnion struct {
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

func (u GetCheckNumberUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GetCheckNumberUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GetCheckNumberUnion) RawJSON() string { return u.JSON.raw }

func (r *GetCheckNumberUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetCheckRecipientUnion contains all possible properties and values from
// [string], [[]string], [CheckAddress].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type GetCheckRecipientUnion struct {
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

func (u GetCheckRecipientUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GetCheckRecipientUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GetCheckRecipientUnion) AsCheckAddress() (v CheckAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GetCheckRecipientUnion) RawJSON() string { return u.JSON.raw }

func (r *GetCheckRecipientUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Value are required.
type PinParam struct {
	// PIN description
	Description string `json:"description,required"`
	// Expected PIN value
	Value string `json:"value,required"`
	paramObj
}

func (r PinParam) MarshalJSON() (data []byte, err error) {
	type shadow PinParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PinParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Date are required.
type RemittanceAdviceParam struct {
	// Remittance id
	ID string `json:"id,required"`
	// Remittance description
	Date string `json:"date,required"`
	// Remittance amount
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Description for invoice
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r RemittanceAdviceParam) MarshalJSON() (data []byte, err error) {
	type shadow RemittanceAdviceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RemittanceAdviceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApprovalListResponse struct {
	// List of sent/received payment
	Checks []ApprovalListResponseCheck `json:"checks,required"`
	Page   int64                       `json:"page,required"`
	Pages  int64                       `json:"pages,required"`
	Total  int64                       `json:"total,required"`
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
func (r ApprovalListResponse) RawJSON() string { return r.JSON.raw }
func (r *ApprovalListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApprovalListResponseCheck struct {
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
	// Current status of the entry
	//
	// Any of "UNPAID", "APPROVED", "VOID".
	Status string `json:"status,required"`
	// Payment comment
	Comment string `json:"comment,nullable"`
	// Payment description/memo
	Description string `json:"description,nullable"`
	// URI where image of the payment can be accessed
	ImageUri string `json:"image_uri,nullable"`
	// Name of third party who received the payment
	Name      string                                  `json:"name,nullable"`
	Number    ApprovalListResponseCheckNumberUnion    `json:"number"`
	Recipient ApprovalListResponseCheckRecipientUnion `json:"recipient"`
	// Email/id or physical address of the check sender
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
func (r ApprovalListResponseCheck) RawJSON() string { return r.JSON.raw }
func (r *ApprovalListResponseCheck) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ApprovalListResponseCheckNumberUnion contains all possible properties and values
// from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ApprovalListResponseCheckNumberUnion struct {
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

func (u ApprovalListResponseCheckNumberUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ApprovalListResponseCheckNumberUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ApprovalListResponseCheckNumberUnion) RawJSON() string { return u.JSON.raw }

func (r *ApprovalListResponseCheckNumberUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ApprovalListResponseCheckRecipientUnion contains all possible properties and
// values from [string], [[]string], [CheckAddress].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type ApprovalListResponseCheckRecipientUnion struct {
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

func (u ApprovalListResponseCheckRecipientUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ApprovalListResponseCheckRecipientUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ApprovalListResponseCheckRecipientUnion) AsCheckAddress() (v CheckAddress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ApprovalListResponseCheckRecipientUnion) RawJSON() string { return u.JSON.raw }

func (r *ApprovalListResponseCheckRecipientUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApprovalUpdateParams struct {
	// Optional description/memo for check
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional debit account id for funds (if sender has multiple bank accounts)
	Account param.Opt[string] `json:"account,omitzero" format:"uuid"`
	// Check amount
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Name of recipient
	Name param.Opt[string] `json:"name,omitzero"`
	// Check number
	Number param.Opt[string] `json:"number,omitzero"`
	// Email or text enabled phone number/id of recipient
	Recipient param.Opt[string] `json:"recipient,omitzero"`
	paramObj
}

func (r ApprovalUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ApprovalUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApprovalUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApprovalListParams struct {
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
	Direction ApprovalListParamsDirection `query:"direction,omitzero" json:"-"`
	// Items per page
	//
	// Any of 10, 25, 50, 100, 250.
	PerPage int64 `query:"per_page,omitzero" json:"-"`
	// Sort
	//
	// Any of "+NUMBER", "-NUMBER", "+TYPE", "-TYPE", "+AMOUNT", "-AMOUNT", "+STATUS",
	// "-STATUS", "+DATE", "-DATE", "+UPDATE", "-UPDATE", "+DESCRIPTION",
	// "-DESCRIPTION".
	Sort ApprovalListParamsSort `query:"sort,omitzero" json:"-"`
	// Status
	//
	// Any of "PAID", "IN_PROCESS", "UNPAID", "VOID", "EXPIRED", "PRINTED", "MAILED",
	// "FAILED", "REFUNDED".
	Status ApprovalListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ApprovalListParams]'s query parameters as `url.Values`.
func (r ApprovalListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction
type ApprovalListParamsDirection string

const (
	ApprovalListParamsDirectionIncoming ApprovalListParamsDirection = "INCOMING"
	ApprovalListParamsDirectionOutgoing ApprovalListParamsDirection = "OUTGOING"
)

// Sort
type ApprovalListParamsSort string

const (
	ApprovalListParamsSortPlusNumber       ApprovalListParamsSort = "+NUMBER"
	ApprovalListParamsSortMinusNumber      ApprovalListParamsSort = "-NUMBER"
	ApprovalListParamsSortPlusType         ApprovalListParamsSort = "+TYPE"
	ApprovalListParamsSortMinusType        ApprovalListParamsSort = "-TYPE"
	ApprovalListParamsSortPlusAmount       ApprovalListParamsSort = "+AMOUNT"
	ApprovalListParamsSortMinusAmount      ApprovalListParamsSort = "-AMOUNT"
	ApprovalListParamsSortPlusStatus       ApprovalListParamsSort = "+STATUS"
	ApprovalListParamsSortMinusStatus      ApprovalListParamsSort = "-STATUS"
	ApprovalListParamsSortPlusDate         ApprovalListParamsSort = "+DATE"
	ApprovalListParamsSortMinusDate        ApprovalListParamsSort = "-DATE"
	ApprovalListParamsSortPlusUpdate       ApprovalListParamsSort = "+UPDATE"
	ApprovalListParamsSortMinusUpdate      ApprovalListParamsSort = "-UPDATE"
	ApprovalListParamsSortPlusDescription  ApprovalListParamsSort = "+DESCRIPTION"
	ApprovalListParamsSortMinusDescription ApprovalListParamsSort = "-DESCRIPTION"
)

// Status
type ApprovalListParamsStatus string

const (
	ApprovalListParamsStatusPaid      ApprovalListParamsStatus = "PAID"
	ApprovalListParamsStatusInProcess ApprovalListParamsStatus = "IN_PROCESS"
	ApprovalListParamsStatusUnpaid    ApprovalListParamsStatus = "UNPAID"
	ApprovalListParamsStatusVoid      ApprovalListParamsStatus = "VOID"
	ApprovalListParamsStatusExpired   ApprovalListParamsStatus = "EXPIRED"
	ApprovalListParamsStatusPrinted   ApprovalListParamsStatus = "PRINTED"
	ApprovalListParamsStatusMailed    ApprovalListParamsStatus = "MAILED"
	ApprovalListParamsStatusFailed    ApprovalListParamsStatus = "FAILED"
	ApprovalListParamsStatusRefunded  ApprovalListParamsStatus = "REFUNDED"
)

type ApprovalNewDigitalParams struct {
	CreateDigitalCheck CreateDigitalCheckParam
	paramObj
}

func (r ApprovalNewDigitalParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateDigitalCheck)
}
func (r *ApprovalNewDigitalParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateDigitalCheck)
}

type ApprovalNewMultiParams struct {
	CreateMultiCheck CreateMultiCheckParam
	paramObj
}

func (r ApprovalNewMultiParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateMultiCheck)
}
func (r *ApprovalNewMultiParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateMultiCheck)
}

type ApprovalNewPhysicalParams struct {
	CreatePhysicalCheck CreatePhysicalCheckParam
	paramObj
}

func (r ApprovalNewPhysicalParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreatePhysicalCheck)
}
func (r *ApprovalNewPhysicalParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreatePhysicalCheck)
}

type ApprovalReleaseParams struct {
	// Unique identifier for check
	ID string `json:"id,required" format:"uuid"`
	paramObj
}

func (r ApprovalReleaseParams) MarshalJSON() (data []byte, err error) {
	type shadow ApprovalReleaseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApprovalReleaseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
