// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package identety

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/identety/identety-go-sdk/internal/apijson"
	"github.com/identety/identety-go-sdk/internal/apiquery"
	"github.com/identety/identety-go-sdk/internal/param"
	"github.com/identety/identety-go-sdk/internal/requestconfig"
	"github.com/identety/identety-go-sdk/option"
)

// ClientService contains methods and other services that help with interacting
// with the identety API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientService] method instead.
type ClientService struct {
	Options []option.RequestOption
}

// NewClientService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewClientService(opts ...option.RequestOption) (r *ClientService) {
	r = &ClientService{}
	r.Options = opts
	return
}

// Create client
func (r *ClientService) New(ctx context.Context, body ClientNewParams, opts ...option.RequestOption) (res *Client, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "clients"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get client details by id
func (r *ClientService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Client, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("clients/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update client
func (r *ClientService) Update(ctx context.Context, id string, body ClientUpdateParams, opts ...option.RequestOption) (res *Client, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("clients/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get all clients
func (r *ClientService) List(ctx context.Context, query ClientListParams, opts ...option.RequestOption) (res *ClientListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "clients"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete client
func (r *ClientService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *Client, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("clients/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Client struct {
	ID            string      `json:"id" api:"required"`
	AllowedGrants []string    `json:"allowedGrants" api:"required"`
	AllowedScopes []string    `json:"allowedScopes" api:"required"`
	ClientID      string      `json:"clientId" api:"required"`
	ClientSecret  string      `json:"clientSecret" api:"required"`
	IsActive      bool        `json:"isActive" api:"required"`
	Name          string      `json:"name" api:"required"`
	RedirectUris  []string    `json:"redirectUris" api:"required"`
	Settings      interface{} `json:"settings" api:"required"`
	Type          ClientType  `json:"type" api:"required"`
	JSON          clientJSON  `json:"-"`
}

// clientJSON contains the JSON metadata for the struct [Client]
type clientJSON struct {
	ID            apijson.Field
	AllowedGrants apijson.Field
	AllowedScopes apijson.Field
	ClientID      apijson.Field
	ClientSecret  apijson.Field
	IsActive      apijson.Field
	Name          apijson.Field
	RedirectUris  apijson.Field
	Settings      apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Client) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientJSON) RawJSON() string {
	return r.raw
}

type ClientType string

const (
	ClientTypePublic  ClientType = "public"
	ClientTypePrivate ClientType = "private"
	ClientTypeM2m     ClientType = "m2m"
)

func (r ClientType) IsKnown() bool {
	switch r {
	case ClientTypePublic, ClientTypePrivate, ClientTypeM2m:
		return true
	}
	return false
}

type ClientListResponse struct {
	Meta  interface{}            `json:"meta" api:"required"`
	Nodes []Client               `json:"nodes" api:"required"`
	JSON  clientListResponseJSON `json:"-"`
}

// clientListResponseJSON contains the JSON metadata for the struct
// [ClientListResponse]
type clientListResponseJSON struct {
	Meta        apijson.Field
	Nodes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientListResponseJSON) RawJSON() string {
	return r.raw
}

type ClientNewParams struct {
	// Client Name
	Name param.Field[string] `json:"name" api:"required"`
	// Client type
	Type param.Field[ClientNewParamsType] `json:"type" api:"required"`
	// Allowed Grants
	AllowedGrants param.Field[[]ClientNewParamsAllowedGrant] `json:"allowedGrants"`
	// Allowed Scopes
	AllowedScopes param.Field[[]string] `json:"allowedScopes"`
	// Redirect URIs
	RedirectUris param.Field[[]string] `json:"redirectUris"`
	// Client Settings
	Settings param.Field[interface{}] `json:"settings"`
}

func (r ClientNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Client type
type ClientNewParamsType string

const (
	ClientNewParamsTypePublic  ClientNewParamsType = "public"
	ClientNewParamsTypePrivate ClientNewParamsType = "private"
	ClientNewParamsTypeM2m     ClientNewParamsType = "m2m"
)

func (r ClientNewParamsType) IsKnown() bool {
	switch r {
	case ClientNewParamsTypePublic, ClientNewParamsTypePrivate, ClientNewParamsTypeM2m:
		return true
	}
	return false
}

type ClientNewParamsAllowedGrant string

const (
	ClientNewParamsAllowedGrantAuthorizationCode ClientNewParamsAllowedGrant = "authorization_code"
	ClientNewParamsAllowedGrantClientCredentials ClientNewParamsAllowedGrant = "client_credentials"
	ClientNewParamsAllowedGrantRefreshToken      ClientNewParamsAllowedGrant = "refresh_token"
)

func (r ClientNewParamsAllowedGrant) IsKnown() bool {
	switch r {
	case ClientNewParamsAllowedGrantAuthorizationCode, ClientNewParamsAllowedGrantClientCredentials, ClientNewParamsAllowedGrantRefreshToken:
		return true
	}
	return false
}

type ClientUpdateParams struct {
	// Client Name
	Name param.Field[string] `json:"name" api:"required"`
	// Allowed Grants
	AllowedGrants param.Field[[]ClientUpdateParamsAllowedGrant] `json:"allowedGrants"`
	// Allowed Scopes
	AllowedScopes param.Field[[]string] `json:"allowedScopes"`
	// Redirect URIs
	RedirectUris param.Field[[]string] `json:"redirectUris"`
	// Client Settings
	Settings param.Field[interface{}] `json:"settings"`
}

func (r ClientUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ClientUpdateParamsAllowedGrant string

const (
	ClientUpdateParamsAllowedGrantAuthorizationCode ClientUpdateParamsAllowedGrant = "authorization_code"
	ClientUpdateParamsAllowedGrantClientCredentials ClientUpdateParamsAllowedGrant = "client_credentials"
	ClientUpdateParamsAllowedGrantRefreshToken      ClientUpdateParamsAllowedGrant = "refresh_token"
)

func (r ClientUpdateParamsAllowedGrant) IsKnown() bool {
	switch r {
	case ClientUpdateParamsAllowedGrantAuthorizationCode, ClientUpdateParamsAllowedGrantClientCredentials, ClientUpdateParamsAllowedGrantRefreshToken:
		return true
	}
	return false
}

type ClientListParams struct {
	// Comma separated column names
	Columns param.Field[ClientListParamsColumns] `query:"columns" api:"required"`
	Limit   param.Field[float64]                 `query:"limit"`
	Page    param.Field[float64]                 `query:"page"`
	Sort    param.Field[ClientListParamsSort]    `query:"sort"`
	SortBy  param.Field[string]                  `query:"sortBy"`
}

// URLQuery serializes [ClientListParams]'s query parameters as `url.Values`.
func (r ClientListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Comma separated column names
type ClientListParamsColumns string

const (
	ClientListParamsColumnsID            ClientListParamsColumns = "id"
	ClientListParamsColumnsClientID      ClientListParamsColumns = "client_id"
	ClientListParamsColumnsClientSecret  ClientListParamsColumns = "client_secret"
	ClientListParamsColumnsName          ClientListParamsColumns = "name"
	ClientListParamsColumnsType          ClientListParamsColumns = "type"
	ClientListParamsColumnsRedirectUris  ClientListParamsColumns = "redirect_uris"
	ClientListParamsColumnsAllowedScopes ClientListParamsColumns = "allowed_scopes"
	ClientListParamsColumnsAllowedGrants ClientListParamsColumns = "allowed_grants"
	ClientListParamsColumnsIsActive      ClientListParamsColumns = "is_active"
	ClientListParamsColumnsRequirePkce   ClientListParamsColumns = "require_pkce"
	ClientListParamsColumnsSettings      ClientListParamsColumns = "settings"
	ClientListParamsColumnsTenantID      ClientListParamsColumns = "tenant_id"
	ClientListParamsColumnsCreatedAt     ClientListParamsColumns = "created_at"
	ClientListParamsColumnsUpdatedAt     ClientListParamsColumns = "updated_at"
)

func (r ClientListParamsColumns) IsKnown() bool {
	switch r {
	case ClientListParamsColumnsID, ClientListParamsColumnsClientID, ClientListParamsColumnsClientSecret, ClientListParamsColumnsName, ClientListParamsColumnsType, ClientListParamsColumnsRedirectUris, ClientListParamsColumnsAllowedScopes, ClientListParamsColumnsAllowedGrants, ClientListParamsColumnsIsActive, ClientListParamsColumnsRequirePkce, ClientListParamsColumnsSettings, ClientListParamsColumnsTenantID, ClientListParamsColumnsCreatedAt, ClientListParamsColumnsUpdatedAt:
		return true
	}
	return false
}

type ClientListParamsSort string

const (
	ClientListParamsSortAsc  ClientListParamsSort = "asc"
	ClientListParamsSortDesc ClientListParamsSort = "desc"
)

func (r ClientListParamsSort) IsKnown() bool {
	switch r {
	case ClientListParamsSortAsc, ClientListParamsSortDesc:
		return true
	}
	return false
}
