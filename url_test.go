// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crawlerdev_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/crawler-dot-dev/api-sdk-go"
	"github.com/crawler-dot-dev/api-sdk-go/internal/testutil"
	"github.com/crawler-dot-dev/api-sdk-go/option"
)

func TestURLExtractTextWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.URLs.ExtractText(context.TODO(), crawlerdev.URLExtractTextParams{
		URL:              "url",
		CleanText:        crawlerdev.Bool(true),
		RenderJs:         crawlerdev.Bool(true),
		StripBoilerplate: crawlerdev.Bool(true),
	})
	if err != nil {
		var apierr *crawlerdev.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
