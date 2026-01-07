// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package apicrawlerdevsdks

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/crawler-dot-dev/api-sdk-go/internal/apiform"
	"github.com/crawler-dot-dev/api-sdk-go/internal/apijson"
	"github.com/crawler-dot-dev/api-sdk-go/internal/requestconfig"
	"github.com/crawler-dot-dev/api-sdk-go/option"
	"github.com/crawler-dot-dev/api-sdk-go/packages/param"
	"github.com/crawler-dot-dev/api-sdk-go/packages/respjson"
)

// ExtractService contains methods and other services that help with interacting
// with the api.crawler.dev-sdks API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtractService] method instead.
type ExtractService struct {
	Options []option.RequestOption
}

// NewExtractService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewExtractService(opts ...option.RequestOption) (r ExtractService) {
	r = ExtractService{}
	r.Options = opts
	return
}

// Upload a file and extract text content from it. Supports PDF, DOC, DOCX, TXT and
// other text-extractable document formats.
func (r *ExtractService) FromFile(ctx context.Context, body ExtractFromFileParams, opts ...option.RequestOption) (res *ExtractFromFileResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/extract/file"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Extract text content from a webpage or document accessible via URL. Supports
// HTML, PDF, and other web-accessible content types.
func (r *ExtractService) FromURL(ctx context.Context, body ExtractFromURLParams, opts ...option.RequestOption) (res *ExtractFromURLResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/extract/url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ExtractFromFileResponse struct {
	ContentType string `json:"contentType"`
	Filename    string `json:"filename"`
	// Markdown representation (included when 'markdown' is in formats array, empty
	// string for non-HTML content)
	Markdown string `json:"markdown"`
	// The size of the entity in bytes
	Size int64 `json:"size"`
	// Extracted plain text (included when 'text' is in formats array)
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		Filename    respjson.Field
		Markdown    respjson.Field
		Size        respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractFromFileResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractFromFileResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractFromURLResponse struct {
	ContentType string `json:"contentType"`
	FinalURL    string `json:"finalUrl"`
	// Markdown representation (included when 'markdown' is in formats array, empty
	// string for non-HTML content)
	Markdown string `json:"markdown"`
	// The size of the entity in bytes
	Size       int64 `json:"size"`
	StatusCode int64 `json:"statusCode"`
	// Extracted plain text (included when 'text' is in formats array)
	Text string `json:"text"`
	URL  string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		FinalURL    respjson.Field
		Markdown    respjson.Field
		Size        respjson.Field
		StatusCode  respjson.Field
		Text        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractFromURLResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractFromURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractFromFileParams struct {
	// The file to upload.
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	// Whether to clean and normalize the extracted text. When enabled (true):
	//
	//   - For HTML content: Removes script, style, and other non-text elements before
	//     extraction
	//   - Normalizes whitespace (collapses multiple spaces/tabs, normalizes newlines)
	//   - Removes empty lines and trims leading/trailing whitespace
	//   - Normalizes Unicode characters (NFC)
	//   - For JSON content: Only minimal cleaning to preserve structure When disabled
	//     (false): Returns raw extracted text without any processing.
	CleanText param.Opt[bool] `json:"cleanText,omitzero"`
	// Array of output formats to include in the response. Options: 'text', 'markdown'.
	//
	//   - 'text': Extracted plain text (always available)
	//   - 'markdown': Markdown representation (only available for HTML content, empty
	//     string otherwise) Defaults to ['text'] if not specified.
	//
	// Any of "text", "markdown".
	Formats []string `json:"formats,omitzero"`
	// Maximum time before the file extraction gives up. Accepts either:
	//
	//   - Integer: milliseconds (e.g., 30000 for 30 seconds)
	//   - String: time format with unit (e.g., "1s", "5h", "3m", "4.4h") Supported
	//     units: s (seconds), m (minutes), h (hours), d (days), ms (milliseconds) Must
	//     be between 5 seconds and 2 minutes. Defaults to "30s" (30 seconds) if not
	//     specified. This controls the timeout for Tika extraction operations on
	//     uploaded files.
	MaxTimeout ExtractFromFileParamsMaxTimeoutUnion `json:"maxTimeout,omitzero"`
	paramObj
}

func (r ExtractFromFileParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractFromFileParamsMaxTimeoutUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractFromFileParamsMaxTimeoutUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ExtractFromFileParamsMaxTimeoutUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractFromFileParamsMaxTimeoutUnion) asAny() any {
	if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type ExtractFromURLParams struct {
	// The URL to extract text from.
	URL string `json:"url,required"`
	// Whether to clean extracted text
	CleanText param.Opt[bool] `json:"cleanText,omitzero"`
	// Maximum number of redirects to follow when fetching the URL. Must be between 0
	// (no redirects) and 20. Defaults to 5 if not specified.
	MaxRedirects param.Opt[int64] `json:"maxRedirects,omitzero"`
	// When enabled, we use a proxy for the request. If set to true, and the 'proxy'
	// option is set, it will be ignored. Defaults to false if not specified. Note:
	// Enabling stealthMode consumes an additional credit/quota point (2 credits total
	// instead of 1) for this request.
	StealthMode param.Opt[bool] `json:"stealthMode,omitzero"`
	// Maximum acceptable age of cached content. This parameter controls how fresh
	// cached data must be to be used.
	//
	//   - If a cached item exists and is younger than this value, it will be used (cache
	//     hit)
	//   - If a cached item exists but is older than this value, it will be ignored and
	//     fresh data will be fetched (cache miss)
	//   - If set to 0, caching is disabled for this request (always fetches fresh data)
	//   - When fresh data is fetched, it will be cached with this value as the TTL for
	//     future requests Accepts either:
	//   - Integer: milliseconds (e.g., 86400000 for 1 day)
	//   - String: time format with unit (e.g., "1s", "5h", "3m", "4.4h", "2d") Supported
	//     units: s (seconds), m (minutes), h (hours), d (days), ms (milliseconds) Must
	//     be between 0 (no caching) and 3 days. Defaults to "2d" (2 days) if not
	//     specified. Examples:
	//   - "1s": Only use cached items less than 1 second old; fetch fresh data if cache
	//     is older
	//   - "1h": Only use cached items less than 1 hour old; fetch fresh data if cache is
	//     older
	//   - 0: Disable caching entirely; always fetch fresh data
	CacheAge ExtractFromURLParamsCacheAgeUnion `json:"cacheAge,omitzero"`
	// Array of output formats to include in the response. Options: 'text', 'markdown'.
	//
	//   - 'text': Extracted plain text (always available)
	//   - 'markdown': Markdown representation (only available for HTML content, empty
	//     string otherwise) Defaults to ['text'] if not specified.
	//
	// Any of "text", "markdown".
	Formats []string `json:"formats,omitzero"`
	// Custom HTTP headers to send with the request (case-insensitive)
	Headers map[string]string `json:"headers,omitzero"`
	// Maximum content length for the URL response. Accepts either:
	//
	//   - Integer: bytes (e.g., 8388608 for 8MB)
	//   - String: size format with unit (e.g., "1kb", "55mb", "1.2gb") Supported units:
	//     b (bytes), kb (kilobytes), mb (megabytes), gb (gigabytes), tb (terabytes) Must
	//     be between 1KB and 8MB. Defaults to "8mb" (8MB) if not specified.
	MaxSize ExtractFromURLParamsMaxSizeUnion `json:"maxSize,omitzero"`
	// Maximum time before the crawler gives up on loading a URL. Accepts either:
	//
	//   - Integer: milliseconds (e.g., 15000 for 15 seconds)
	//   - String: time format with unit (e.g., "1s", "5h", "3m", "4.4h") Supported
	//     units: s (seconds), m (minutes), h (hours), d (days), ms (milliseconds) Must
	//     be between 1 second and 30 seconds. Defaults to "10s" (10 seconds) if not
	//     specified.
	MaxTimeout ExtractFromURLParamsMaxTimeoutUnion `json:"maxTimeout,omitzero"`
	// Proxy configuration for the request
	Proxy ExtractFromURLParamsProxy `json:"proxy,omitzero"`
	paramObj
}

func (r ExtractFromURLParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractFromURLParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractFromURLParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractFromURLParamsCacheAgeUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractFromURLParamsCacheAgeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ExtractFromURLParamsCacheAgeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractFromURLParamsCacheAgeUnion) asAny() any {
	if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractFromURLParamsMaxSizeUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractFromURLParamsMaxSizeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ExtractFromURLParamsMaxSizeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractFromURLParamsMaxSizeUnion) asAny() any {
	if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractFromURLParamsMaxTimeoutUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractFromURLParamsMaxTimeoutUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ExtractFromURLParamsMaxTimeoutUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractFromURLParamsMaxTimeoutUnion) asAny() any {
	if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Proxy configuration for the request
type ExtractFromURLParamsProxy struct {
	// Proxy password for authentication
	Password param.Opt[string] `json:"password,omitzero"`
	// Proxy server URL (e.g., http://proxy.example.com:8080 or
	// socks5://proxy.example.com:1080)
	Server param.Opt[string] `json:"server,omitzero"`
	// Proxy username for authentication
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r ExtractFromURLParamsProxy) MarshalJSON() (data []byte, err error) {
	type shadow ExtractFromURLParamsProxy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractFromURLParamsProxy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
