// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package checkbook

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Munchpass/checkbook/internal/apijson"
	"github.com/Munchpass/checkbook/internal/requestconfig"
	"github.com/Munchpass/checkbook/option"
	"github.com/Munchpass/checkbook/packages/param"
	"github.com/Munchpass/checkbook/packages/respjson"
)

// CheckDepositService contains methods and other services that help with
// interacting with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCheckDepositService] method instead.
type CheckDepositService struct {
	Options []option.RequestOption
}

// NewCheckDepositService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCheckDepositService(opts ...option.RequestOption) (r CheckDepositService) {
	r = CheckDepositService{}
	r.Options = opts
	return
}

// Deposit a payment
func (r *CheckDepositService) New(ctx context.Context, checkID string, body CheckDepositNewParams, opts ...option.RequestOption) (res *GetCheck, err error) {
	opts = append(r.Options[:], opts...)
	if checkID == "" {
		err = errors.New("missing required check_id parameter")
		return
	}
	path := fmt.Sprintf("v3/check/deposit/%s", checkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a deposited payment
func (r *CheckDepositService) Get(ctx context.Context, checkID string, opts ...option.RequestOption) (res *CheckDepositGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if checkID == "" {
		err = errors.New("missing required check_id parameter")
		return
	}
	path := fmt.Sprintf("v3/check/%s/deposit", checkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CheckDepositGetResponse struct {
	// Check type
	CheckType string    `json:"check_type,required"`
	Eta       time.Time `json:"eta,required" format:"date"`
	Recipient string    `json:"recipient,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CheckType   respjson.Field
		Eta         respjson.Field
		Recipient   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckDepositGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CheckDepositGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckDepositNewParams struct {
	// ID of the account to deposit the payment
	Account string `json:"account,required" format:"uuid"`
	paramObj
}

func (r CheckDepositNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CheckDepositNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CheckDepositNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
