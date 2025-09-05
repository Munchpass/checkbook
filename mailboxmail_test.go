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

func TestMailboxMailGet(t *testing.T) {
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
	_, err := client.Mailbox.Mail.Get(
		context.TODO(),
		"item_id",
		checkbook.MailboxMailGetParams{
			MailboxID: "mailbox_id",
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

func TestMailboxMailListWithOptionalParams(t *testing.T) {
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
	_, err := client.Mailbox.Mail.List(
		context.TODO(),
		"mailbox_id",
		checkbook.MailboxMailListParams{
			EndDate:   checkbook.Time(time.Now()),
			Page:      checkbook.Int(1),
			PerPage:   50,
			Q:         checkbook.String("payment"),
			StartDate: checkbook.Time(time.Now()),
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

func TestMailboxMailGetAttachment(t *testing.T) {
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
	_, err := client.Mailbox.Mail.GetAttachment(
		context.TODO(),
		"item_id",
		checkbook.MailboxMailGetAttachmentParams{
			MailboxID: "mailbox_id",
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
