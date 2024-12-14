// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package identety

import (
	"context"
	"net/http"

	"github.com/identety/identety-go-sdk/internal/apijson"
	"github.com/identety/identety-go-sdk/internal/param"
	"github.com/identety/identety-go-sdk/internal/requestconfig"
	"github.com/identety/identety-go-sdk/option"
)

// UserService contains methods and other services that help with interacting with
// the identety API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	return
}

func (r *UserService) New(ctx context.Context, body UserNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type UserNewParams struct {
	Address    param.Field[UserNewParamsAddress] `json:"address,required"`
	Email      param.Field[string]               `json:"email,required"`
	FamilyName param.Field[string]               `json:"familyName,required"`
	GivenName  param.Field[string]               `json:"givenName,required"`
	Locale     param.Field[string]               `json:"locale,required"`
	Metadata   param.Field[interface{}]          `json:"metadata,required"`
	Name       param.Field[string]               `json:"name,required"`
	Password   param.Field[string]               `json:"password,required"`
	Picture    param.Field[string]               `json:"picture,required"`
}

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserNewParamsAddress struct {
	Country       param.Field[string] `json:"country,required"`
	Locality      param.Field[string] `json:"locality,required"`
	PostalCode    param.Field[string] `json:"postalCode,required"`
	Region        param.Field[string] `json:"region,required"`
	StreetAddress param.Field[string] `json:"streetAddress,required"`
}

func (r UserNewParamsAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
