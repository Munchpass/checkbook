// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package checkbook_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Munchpass/checkbook"
	"github.com/Munchpass/checkbook/internal/testutil"
	"github.com/Munchpass/checkbook/option"
)

func TestSubscriptionGet(t *testing.T) {
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
	_, err := client.Subscription.Get(context.TODO(), "subscription_id")
	if err != nil {
		var apierr *checkbook.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Subscription.Update(
		context.TODO(),
		"subscription_id",
		checkbook.SubscriptionUpdateParams{
			Autopay: checkbook.Bool(true),
			Skipped: []int64{0},
		},
	)
	if err != nil {
		var apierr *checkbook.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscription.List(context.TODO(), checkbook.SubscriptionListParams{
		Direction: checkbook.SubscriptionListParamsDirectionIncoming,
		EndDate:   checkbook.Time(time.Now()),
		Page:      checkbook.Int(1),
		PerPage:   10,
		Q:         checkbook.String("q"),
		Sort:      checkbook.SubscriptionListParamsSortPlusNumber,
		StartDate: checkbook.Time(time.Now()),
		Status:    checkbook.SubscriptionListParamsStatusPaid,
	})
	if err != nil {
		var apierr *checkbook.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionDelete(t *testing.T) {
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
	err := client.Subscription.Delete(context.TODO(), "subscription_id")
	if err != nil {
		var apierr *checkbook.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionNewInvoiceWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscription.NewInvoice(context.TODO(), checkbook.SubscriptionNewInvoiceParams{
		Amount:      0.01,
		Description: "description",
		Interval:    checkbook.SubscriptionNewInvoiceParamsIntervalWeekly,
		Name:        "xx",
		Recipient:   "recipient",
		Account:     checkbook.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Duration:    checkbook.Int(1),
		StartDate:   checkbook.Time(time.Now()),
	})
	if err != nil {
		var apierr *checkbook.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionNewPaymentWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscription.NewPayment(context.TODO(), checkbook.SubscriptionNewPaymentParams{
		Amount:   0.01,
		Interval: checkbook.SubscriptionNewPaymentParamsIntervalWeekly,
		Name:     "xx",
		Recipient: checkbook.SubscriptionNewPaymentParamsRecipientUnion{
			OfString: checkbook.String("string"),
		},
		Account:     checkbook.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Description: checkbook.String("description"),
		Duration:    checkbook.Int(1),
		RemittanceAdvice: []checkbook.RemittanceAdviceParam{{
			ID:          "id",
			Date:        "date",
			Amount:      checkbook.Float(0),
			Description: checkbook.String("description"),
		}},
		StartDate: checkbook.Time(time.Now()),
	})
	if err != nil {
		var apierr *checkbook.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
