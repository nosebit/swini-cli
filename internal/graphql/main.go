package graphql

import (
	"net/http"

	"swini-cli/internal/auth"

	"github.com/Yamashou/gqlgenc/clientv2"
)

var (
	SharedClient *Client
)

type headerRoundTripper struct {
	headers http.Header
	rt      http.RoundTripper
}

func (h *headerRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range h.headers {
		for _, vv := range v {
			req.Header.Add(k, vv)
		}
	}
	return h.rt.RoundTrip(req)
}

func InitSharedClient(baseURL string, options *clientv2.Options) {
	headers, _ := auth.HttpHeaders()

	httpClient := &http.Client{
		Transport: &headerRoundTripper{
			headers: headers,
			rt:      http.DefaultTransport,
		},
	}

	SharedClient = NewClient(httpClient, baseURL, options)
}
