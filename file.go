// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crawlerdev

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/stainless-sdks/crawler.dev-go/internal/apiform"
	"github.com/stainless-sdks/crawler.dev-go/internal/apijson"
	"github.com/stainless-sdks/crawler.dev-go/internal/requestconfig"
	"github.com/stainless-sdks/crawler.dev-go/option"
	"github.com/stainless-sdks/crawler.dev-go/packages/param"
	"github.com/stainless-sdks/crawler.dev-go/packages/respjson"
)

// FileService contains methods and other services that help with interacting with
// the crawler.dev API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	Options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r FileService) {
	r = FileService{}
	r.Options = opts
	return
}

// Upload a file and extract text content from it. Supports PDF, DOC, DOCX, TXT and
// other text-extractable document formats.
func (r *FileService) ExtractText(ctx context.Context, body FileExtractTextParams, opts ...option.RequestOption) (res *FileExtractTextResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/files/text"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type FileExtractTextResponse struct {
	ContentType   string `json:"contentType"`
	ExtractedText string `json:"extractedText"`
	Filename      string `json:"filename"`
	SizeBytes     int64  `json:"sizeBytes"`
	Success       bool   `json:"success"`
	TextLength    int64  `json:"textLength"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType   respjson.Field
		ExtractedText respjson.Field
		Filename      respjson.Field
		SizeBytes     respjson.Field
		Success       respjson.Field
		TextLength    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileExtractTextResponse) RawJSON() string { return r.JSON.raw }
func (r *FileExtractTextResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileExtractTextParams struct {
	// The file to upload.
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	// Whether to clean the extracted text
	CleanText param.Opt[bool] `json:"clean_text,omitzero"`
	// Whether to remove boilerplate text
	StripBoilerplate param.Opt[bool] `json:"strip_boilerplate,omitzero"`
	paramObj
}

func (r FileExtractTextParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
