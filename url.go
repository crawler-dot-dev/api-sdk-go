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
	// Maximum cache time in milliseconds for the webpage. Must be between 0 (no
	// caching) and 259200000 (3 days). Defaults to 172800000 (2 days) if not
	// specified.
	CacheAge param.Opt[int64] `json:"cache_age,omitzero"`
	// Whether to clean extracted text
	CleanText param.Opt[bool] `json:"clean_text,omitzero"`
	// Maximum number of redirects to follow when fetching the URL. Must be between 0
	// (no redirects) and 20. Defaults to 5 if not specified.
	MaxRedirects param.Opt[int64] `json:"max_redirects,omitzero"`
	// Maximum content length in bytes for the URL response. Must be between 1024 (1KB)
	// and 52428800 (50MB). Defaults to 10485760 (10MB) if not specified.
	MaxSize param.Opt[int64] `json:"max_size,omitzero"`
	// Maximum time in milliseconds before the crawler gives up on loading a URL. Must
	// be between 1000 (1 second) and 30000 (30 seconds). Defaults to 10000 (10
	// seconds) if not specified.
	MaxTimeout param.Opt[int64] `json:"max_timeout,omitzero"`
	// When enabled, we use a proxy for the request. If set to true, and the 'proxy'
	// option is set, it will be ignored. Defaults to false if not specified. Note:
	// Enabling stealth_mode consumes an additional credit/quota point (2 credits total
	// instead of 1) for this request.
	StealthMode param.Opt[bool] `json:"stealth_mode,omitzero"`
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
