// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package identety_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/identety/identety-go-sdk"
	"github.com/identety/identety-go-sdk/internal/testutil"
	"github.com/identety/identety-go-sdk/option"
)

func TestOrgGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := identety.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Orgs.Get(context.TODO(), "id")
	if err != nil {
		var apierr *identety.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
