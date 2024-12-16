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

func TestUserNew(t *testing.T) {
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
	_, err := client.Users.New(context.TODO(), identety.UserNewParams{
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
		var apierr *identety.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserGet(t *testing.T) {
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
	_, err := client.Users.Get(context.TODO(), "id")
	if err != nil {
		var apierr *identety.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdate(t *testing.T) {
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
	_, err := client.Users.Update(
		context.TODO(),
		"id",
		identety.UserUpdateParams{
			Address: identety.F(identety.UserUpdateParamsAddress{
				Country:       identety.F("USA"),
				Locality:      identety.F("New York"),
				PostalCode:    identety.F("10001"),
				Region:        identety.F("NY"),
				StreetAddress: identety.F("123 Main St"),
			}),
			FamilyName: identety.F("Doe"),
			GivenName:  identety.F("John"),
			Locale:     identety.F("en-US"),
			Metadata: identety.F[any](map[string]interface{}{
				"customField": "value",
			}),
			Name:    identety.F("John Doe"),
			Picture: identety.F("https://example.com/photo.jpg"),
		},
	)
	if err != nil {
		var apierr *identety.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserListWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.List(context.TODO(), identety.UserListParams{
		Columns: identety.F(identety.UserListParamsColumnsID),
		Limit:   identety.F(0.000000),
		Page:    identety.F(0.000000),
		Sort:    identety.F(identety.UserListParamsSortAsc),
		SortBy:  identety.F("sortBy"),
	})
	if err != nil {
		var apierr *identety.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserDelete(t *testing.T) {
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
	_, err := client.Users.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *identety.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
