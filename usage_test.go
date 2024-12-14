// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package identety_test

import (
	"context"
	"os"
	"testing"

	"github.com/identety/identety-go-sdk"
	"github.com/identety/identety-go-sdk/internal/testutil"
	"github.com/identety/identety-go-sdk/option"
)

func TestUsage(t *testing.T) {
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
	client, err := client.Clients.New(context.TODO(), identety.ClientNewParams{
		Name: identety.F("name"),
		Type: identety.F(identety.ClientNewParamsTypePublic),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", client.ID)
}
