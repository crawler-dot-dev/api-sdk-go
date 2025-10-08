// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crawlerdev

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/crawler.dev-go/internal/apijson"
	"github.com/stainless-sdks/crawler.dev-go/internal/requestconfig"
	"github.com/stainless-sdks/crawler.dev-go/option"
	"github.com/stainless-sdks/crawler.dev-go/packages/param"
	"github.com/stainless-sdks/crawler.dev-go/packages/respjson"
)

// URLService contains methods and other services that help with interacting with
// the crawler.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewURLService] method instead.
type URLService struct {
	Options []option.RequestOption
}

// NewURLService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewURLService(opts ...option.RequestOption) (r URLService) {
	r = URLService{}
	r.Options = opts
	return
}

// Extract text content from a webpage or document accessible via URL. Supports
// HTML, PDF, and other web-accessible content types.
func (r *URLService) ExtractText(ctx context.Context, body URLExtractTextParams, opts ...option.RequestOption) (res *URLExtractTextResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/urls/text"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type URLExtractTextResponse struct {
	ContentType   string `json:"contentType"`
	ExtractedText string `json:"extractedText"`
	FinalURL      string `json:"finalUrl"`
	SizeBytes     int64  `json:"sizeBytes"`
	StatusCode    int64  `json:"statusCode"`
	Success       bool   `json:"success"`
	TextLength    int64  `json:"textLength"`
	URL           string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType   respjson.Field
		ExtractedText respjson.Field
		FinalURL      respjson.Field
		SizeBytes     respjson.Field
		StatusCode    respjson.Field
		Success       respjson.Field
		TextLength    respjson.Field
		URL           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r URLExtractTextResponse) RawJSON() string { return r.JSON.raw }
func (r *URLExtractTextResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type URLExtractTextParams struct {
	// The URL to extract text from.
	URL string `json:"url,required"`
	// Whether to clean extracted text
	CleanText param.Opt[bool] `json:"clean_text,omitzero"`
	// Whether to render JavaScript for HTML content. This parameter is ignored for
	// binary content types (PDF, DOC, etc.) since they are not HTML.
	RenderJs param.Opt[bool] `json:"render_js,omitzero"`
	// Whether to remove boilerplate text
	StripBoilerplate param.Opt[bool] `json:"strip_boilerplate,omitzero"`
	paramObj
}

func (r URLExtractTextParams) MarshalJSON() (data []byte, err error) {
	type shadow URLExtractTextParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *URLExtractTextParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
