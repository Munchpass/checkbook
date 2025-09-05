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

// UserAPIKeyService contains methods and other services that help with interacting
// with the checkbook API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserAPIKeyService] method instead.
type UserAPIKeyService struct {
	Options []option.RequestOption
}

// NewUserAPIKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserAPIKeyService(opts ...option.RequestOption) (r UserAPIKeyService) {
	r = UserAPIKeyService{}
	r.Options = opts
	return
}

// Generate new API keys for the user
func (r *UserAPIKeyService) New(ctx context.Context, body UserAPIKeyNewParams, opts ...option.RequestOption) (res *NewAPIKey, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/user/api_key"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Return the API keys for the user
func (r *UserAPIKeyService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserAPIKeyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v3/user/api_key"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete API key for user
func (r *UserAPIKeyService) Delete(ctx context.Context, keyID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return
	}
	path := fmt.Sprintf("v3/user/api_key/%s", keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type NewAPIKey struct {
	// Id of API key
	ID string `json:"id"`
	// When the API key expires
	ExpirationDate time.Time `json:"expiration_date,nullable" format:"date"`
	// Name of the API key
	Name string `json:"name,nullable"`
	// Secret key of the user
	Secret string `json:"secret"`
	// Webhook key of the user
	WebhookKey string `json:"webhook_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		ExpirationDate respjson.Field
		Name           respjson.Field
		Secret         respjson.Field
		WebhookKey     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewAPIKey) RawJSON() string { return r.JSON.raw }
func (r *NewAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response fields for api key list
type UserAPIKeyGetResponse struct {
	// List of API keys
	APIKeys []NewAPIKey `json:"api_keys,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKeys     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserAPIKeyGetResponse) RawJSON() string { return r.JSON.raw }
func (r *UserAPIKeyGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserAPIKeyNewParams struct {
	// Expiration date for API key
	ExpirationDate param.Opt[time.Time] `json:"expiration_date,omitzero" format:"date"`
	// Name of API Key
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r UserAPIKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UserAPIKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserAPIKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
