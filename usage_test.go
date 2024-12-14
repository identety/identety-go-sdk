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
	user, err := client.Users.New(context.TODO(), identety.UserNewParams{
		Address: identety.F(identety.UserNewParamsAddress{
			Country:       identety.F("USA"),
			Locality:      identety.F("New York"),
			PostalCode:    identety.F("10001"),
			Region:        identety.F("NY"),
			StreetAddress: identety.F("123 Main St"),
		}),
		Email:      identety.F("john@example.com"),
		FamilyName: identety.F("Doe"),
		GivenName:  identety.F("John"),
		Locale:     identety.F("en-US"),
		Metadata: identety.F[any](map[string]interface{}{
			"customField": "value",
		}),
		Name:     identety.F("John Doe"),
		Password: identety.F("password123"),
		Picture:  identety.F("https://example.com/photo.jpg"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", user.ID)
}
