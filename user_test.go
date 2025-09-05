// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package checkbook_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/checkbook-go"
	"github.com/stainless-sdks/checkbook-go/internal/testutil"
	"github.com/stainless-sdks/checkbook-go/option"
)

func TestUserNew(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := checkbook.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.User.New(context.TODO(), checkbook.UserNewParams{
		Name:   "xx",
		UserID: "xx",
	})
	if err != nil {
		var apierr *checkbook.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := checkbook.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.User.Get(context.TODO())
	if err != nil {
		var apierr *checkbook.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := checkbook.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.User.Update(context.TODO(), checkbook.UserUpdateParams{
		Bank: checkbook.UserUpdateParamsBank{
			Billing: checkbook.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Default: checkbook.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
		Brand: checkbook.UserUpdateParamsBrand{
			Button: checkbook.UserUpdateParamsBrandButton{
				Color: checkbook.ColorParam{
					Primary:   checkbook.String("primary"),
					Secondary: checkbook.String("secondary"),
				},
			},
			Check: checkbook.UserUpdateParamsBrandCheck{
				Color: checkbook.ColorParam{
					Primary:   checkbook.String("primary"),
					Secondary: checkbook.String("secondary"),
				},
				Logo: checkbook.String("logo"),
			},
			Footer:  checkbook.String("footer"),
			Logo:    checkbook.String("logo"),
			ReplyTo: checkbook.String("reply_to"),
			Slogan:  checkbook.String("slogan"),
		},
		Developer: checkbook.UserUpdateParamsDeveloper{
			WebhookUri: checkbook.String("webhook_uri"),
		},
		Merchant: checkbook.UserUpdateParamsMerchant{
			Address: checkbook.UpdateMerchantAddressParam{
				City:    checkbook.String("city"),
				Country: checkbook.String("country"),
				Line1:   checkbook.String("line_1"),
				Line2:   checkbook.String("line_2"),
				State:   checkbook.String("state"),
				Zip:     checkbook.String("xxxxx"),
			},
			IncorporationDate: checkbook.Time(time.Now()),
			Industry:          11,
			LegalFirstname:    checkbook.String("legal_firstname"),
			LegalLastname:     checkbook.String("legal_lastname"),
			PrincipalAddress: checkbook.UpdateMerchantAddressParam{
				City:    checkbook.String("city"),
				Country: checkbook.String("country"),
				Line1:   checkbook.String("line_1"),
				Line2:   checkbook.String("line_2"),
				State:   checkbook.String("state"),
				Zip:     checkbook.String("xxxxx"),
			},
			PrincipalCitizenshipStatus: "us_citizen",
			PrincipalOccupation:        checkbook.String("principal_occupation"),
			TaxID:                      checkbook.String("xxxxxxxxx"),
			Website:                    checkbook.String("website"),
		},
		Payment: checkbook.UserUpdateParamsPayment{
			Expiration: checkbook.Int(10),
		},
		User: checkbook.UserUpdateParamsUser{
			BusinessName:  checkbook.String("xx"),
			CheckNumber:   checkbook.Int(1),
			Dob:           checkbook.Time(time.Now()),
			FirstName:     checkbook.String("xx"),
			InvoiceNumber: checkbook.Int(1),
			LastName:      checkbook.String("xx"),
			Password:      checkbook.String("xxxxxx"),
			Phone:         checkbook.String("phone"),
			Ssn:           checkbook.String("xxxx"),
		},
	})
	if err != nil {
		var apierr *checkbook.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := checkbook.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.User.List(context.TODO(), checkbook.UserListParams{
		Page:    checkbook.Int(1),
		PerPage: 10,
		Q:       checkbook.String("q"),
		Sort:    checkbook.UserListParamsSortPlusDate,
	})
	if err != nil {
		var apierr *checkbook.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := checkbook.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.User.Delete(context.TODO(), "user_id")
	if err != nil {
		var apierr *checkbook.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserAddSignature(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := checkbook.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.User.AddSignature(context.TODO(), checkbook.UserAddSignatureParams{
		Signature: "signature",
	})
	if err != nil {
		var apierr *checkbook.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
