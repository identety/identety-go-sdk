// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package identety

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/identety-go/internal/requestconfig"
	"github.com/stainless-sdks/identety-go/option"
)

// OrgService contains methods and other services that help with interacting with
// the identety API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgService] method instead.
type OrgService struct {
	Options []option.RequestOption
}

// NewOrgService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrgService(opts ...option.RequestOption) (r *OrgService) {
	r = &OrgService{}
	r.Options = opts
	return
}

func (r *OrgService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("org/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}
