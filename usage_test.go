// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crawlerdev_test

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/crawler-dot-dev/api-sdk-go"
	"github.com/crawler-dot-dev/api-sdk-go/internal/testutil"
	"github.com/crawler-dot-dev/api-sdk-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := crawlerdev.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	response, err := client.Files.ExtractText(context.TODO(), crawlerdev.FileExtractTextParams{
		File: io.Reader(bytes.NewBuffer([]byte("REPLACE_ME"))),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.ContentType)
}
