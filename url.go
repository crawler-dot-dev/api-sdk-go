// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crawlerdev

import (
	"context"
	"net/http"
	"slices"

	"github.com/crawler-dot-dev/api-sdk-go/internal/apijson"
	"github.com/crawler-dot-dev/api-sdk-go/internal/requestconfig"
	"github.com/crawler-dot-dev/api-sdk-go/option"
	"github.com/crawler-dot-dev/api-sdk-go/packages/param"
	"github.com/crawler-dot-dev/api-sdk-go/packages/respjson"
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
	TextLength    int64  `json:"textLength"`
	URL           string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType   respjson.Field
		ExtractedText respjson.Field
		FinalURL      respjson.Field
		SizeBytes     respjson.Field
		StatusCode    respjson.Field
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
	// Custom HTTP headers to send with the request (case-insensitive)
	Headers map[string]string `json:"headers,omitzero"`
	// Proxy configuration for the request
	Proxy URLExtractTextParamsProxy `json:"proxy,omitzero"`
	paramObj
}

func (r URLExtractTextParams) MarshalJSON() (data []byte, err error) {
	type shadow URLExtractTextParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *URLExtractTextParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Proxy configuration for the request
type URLExtractTextParamsProxy struct {
	// Proxy password for authentication
	Password param.Opt[string] `json:"password,omitzero"`
	// Proxy server URL (e.g., http://proxy.example.com:8080 or
	// socks5://proxy.example.com:1080)
	Server param.Opt[string] `json:"server,omitzero"`
	// Proxy username for authentication
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r URLExtractTextParamsProxy) MarshalJSON() (data []byte, err error) {
	type shadow URLExtractTextParamsProxy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *URLExtractTextParamsProxy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
