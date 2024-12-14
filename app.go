// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package identety

import (
	"context"
	"net/http"

	"github.com/identety/identety-go-sdk/internal/requestconfig"
	"github.com/identety/identety-go-sdk/option"
)

// AppService contains methods and other services that help with interacting with
// the identety API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppService] method instead.
type AppService struct {
	Options []option.RequestOption
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r *AppService) {
	r = &AppService{}
	r.Options = opts
	return
}

func (r *AppService) Get(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := ""
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}