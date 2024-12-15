// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package identety

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/identety-go/internal/apijson"
	"github.com/stainless-sdks/identety-go/internal/apiquery"
	"github.com/stainless-sdks/identety-go/internal/param"
	"github.com/stainless-sdks/identety-go/internal/requestconfig"
	"github.com/stainless-sdks/identety-go/option"
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

// Create client
func (r *UserService) New(ctx context.Context, body UserNewParams, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	path := "users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get client details by id
func (r *UserService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("users/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update client
func (r *UserService) Update(ctx context.Context, id string, body UserUpdateParams, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("users/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get all users
func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *UserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete client
func (r *UserService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("users/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type User struct {
	ID         string      `json:"id,required"`
	Address    UserAddress `json:"address,required"`
	FamilyName string      `json:"familyName,required"`
	GivenName  string      `json:"givenName,required"`
	Locale     string      `json:"locale,required"`
	Metadata   interface{} `json:"metadata,required"`
	Name       string      `json:"name,required"`
	Picture    string      `json:"picture,required"`
	JSON       userJSON    `json:"-"`
}

// userJSON contains the JSON metadata for the struct [User]
type userJSON struct {
	ID          apijson.Field
	Address     apijson.Field
	FamilyName  apijson.Field
	GivenName   apijson.Field
	Locale      apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Picture     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *User) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userJSON) RawJSON() string {
	return r.raw
}

type UserAddress struct {
	Country       string          `json:"country,required"`
	Locality      string          `json:"locality,required"`
	PostalCode    string          `json:"postalCode,required"`
	Region        string          `json:"region,required"`
	StreetAddress string          `json:"streetAddress,required"`
	JSON          userAddressJSON `json:"-"`
}

// userAddressJSON contains the JSON metadata for the struct [UserAddress]
type userAddressJSON struct {
	Country       apijson.Field
	Locality      apijson.Field
	PostalCode    apijson.Field
	Region        apijson.Field
	StreetAddress apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userAddressJSON) RawJSON() string {
	return r.raw
}

type UserListResponse struct {
	Meta  interface{}          `json:"meta,required"`
	Nodes []User               `json:"nodes,required"`
	JSON  userListResponseJSON `json:"-"`
}

// userListResponseJSON contains the JSON metadata for the struct
// [UserListResponse]
type userListResponseJSON struct {
	Meta        apijson.Field
	Nodes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListResponseJSON) RawJSON() string {
	return r.raw
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

type UserUpdateParams struct {
	Address    param.Field[UserUpdateParamsAddress] `json:"address,required"`
	FamilyName param.Field[string]                  `json:"familyName,required"`
	GivenName  param.Field[string]                  `json:"givenName,required"`
	Locale     param.Field[string]                  `json:"locale,required"`
	Metadata   param.Field[interface{}]             `json:"metadata,required"`
	Name       param.Field[string]                  `json:"name,required"`
	Picture    param.Field[string]                  `json:"picture,required"`
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateParamsAddress struct {
	Country       param.Field[string] `json:"country,required"`
	Locality      param.Field[string] `json:"locality,required"`
	PostalCode    param.Field[string] `json:"postalCode,required"`
	Region        param.Field[string] `json:"region,required"`
	StreetAddress param.Field[string] `json:"streetAddress,required"`
}

func (r UserUpdateParamsAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserListParams struct {
	// Comma separated column names
	Columns param.Field[UserListParamsColumns] `query:"columns,required"`
	Limit   param.Field[float64]               `query:"limit"`
	Page    param.Field[float64]               `query:"page"`
	Sort    param.Field[UserListParamsSort]    `query:"sort"`
	SortBy  param.Field[string]                `query:"sortBy"`
}

// URLQuery serializes [UserListParams]'s query parameters as `url.Values`.
func (r UserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Comma separated column names
type UserListParamsColumns string

const (
	UserListParamsColumnsID                  UserListParamsColumns = "id"
	UserListParamsColumnsEmail               UserListParamsColumns = "email"
	UserListParamsColumnsName                UserListParamsColumns = "name"
	UserListParamsColumnsPhoneNumber         UserListParamsColumns = "phone_number"
	UserListParamsColumnsAddress             UserListParamsColumns = "address"
	UserListParamsColumnsLocale              UserListParamsColumns = "locale"
	UserListParamsColumnsEmailVerified       UserListParamsColumns = "email_verified"
	UserListParamsColumnsPhoneNumberVerified UserListParamsColumns = "phone_number_verified"
	UserListParamsColumnsCreatedAt           UserListParamsColumns = "created_at"
	UserListParamsColumnsUpdatedAt           UserListParamsColumns = "updated_at"
)

func (r UserListParamsColumns) IsKnown() bool {
	switch r {
	case UserListParamsColumnsID, UserListParamsColumnsEmail, UserListParamsColumnsName, UserListParamsColumnsPhoneNumber, UserListParamsColumnsAddress, UserListParamsColumnsLocale, UserListParamsColumnsEmailVerified, UserListParamsColumnsPhoneNumberVerified, UserListParamsColumnsCreatedAt, UserListParamsColumnsUpdatedAt:
		return true
	}
	return false
}

type UserListParamsSort string

const (
	UserListParamsSortAsc  UserListParamsSort = "asc"
	UserListParamsSortDesc UserListParamsSort = "desc"
)

func (r UserListParamsSort) IsKnown() bool {
	switch r {
	case UserListParamsSortAsc, UserListParamsSortDesc:
		return true
	}
	return false
}
