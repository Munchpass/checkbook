// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package checkbook_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/checkbook-go"
	"github.com/stainless-sdks/checkbook-go/internal/testutil"
	"github.com/stainless-sdks/checkbook-go/option"
)

func TestAccountWalletNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.Wallet.New(context.TODO(), checkbook.AccountWalletNewParams{
		CreateWalletRequest: checkbook.CreateWalletRequestParam{
			Name: checkbook.String("Wallet #1"),
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

func TestAccountWalletList(t *testing.T) {
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
	_, err := client.Account.Wallet.List(context.TODO())
	if err != nil {
		var apierr *checkbook.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
