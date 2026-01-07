// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package apicrawlerdevsdks_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/api.crawler.dev-sdks-go"
	"github.com/stainless-sdks/api.crawler.dev-sdks-go/internal/testutil"
	"github.com/stainless-sdks/api.crawler.dev-sdks-go/option"
)

func TestExtractFromFileWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := apicrawlerdevsdks.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Extract.FromFile(context.TODO(), apicrawlerdevsdks.ExtractFromFileParams{
		File:      io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		CleanText: apicrawlerdevsdks.Bool(true),
		Formats:   []string{"text", "markdown"},
		MaxTimeout: apicrawlerdevsdks.ExtractFromFileParamsMaxTimeoutUnion{
			OfString: apicrawlerdevsdks.String("30s"),
		},
	})
	if err != nil {
		var apierr *apicrawlerdevsdks.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExtractFromURLWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := apicrawlerdevsdks.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Extract.FromURL(context.TODO(), apicrawlerdevsdks.ExtractFromURLParams{
		URL: "url",
		CacheAge: apicrawlerdevsdks.ExtractFromURLParamsCacheAgeUnion{
			OfString: apicrawlerdevsdks.String("1d"),
		},
		CleanText: apicrawlerdevsdks.Bool(true),
		Formats:   []string{"text", "markdown"},
		Headers: map[string]string{
			"User-Agent":      "Custom Bot/1.0",
			"X-API-Key":       "my-api-key",
			"Accept-Language": "en-US",
		},
		MaxRedirects: apicrawlerdevsdks.Int(5),
		MaxSize: apicrawlerdevsdks.ExtractFromURLParamsMaxSizeUnion{
			OfString: apicrawlerdevsdks.String("8mb"),
		},
		MaxTimeout: apicrawlerdevsdks.ExtractFromURLParamsMaxTimeoutUnion{
			OfString: apicrawlerdevsdks.String("15s"),
		},
		Proxy: apicrawlerdevsdks.ExtractFromURLParamsProxy{
			Password: apicrawlerdevsdks.String("password"),
			Server:   apicrawlerdevsdks.String("server"),
			Username: apicrawlerdevsdks.String("username"),
		},
		StealthMode: apicrawlerdevsdks.Bool(true),
	})
	if err != nil {
		var apierr *apicrawlerdevsdks.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
