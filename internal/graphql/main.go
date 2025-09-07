package graphql

import (
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
)

var (
	SharedClient *Client
)

func InitSharedClient(baseURL string, options *clientv2.Options) {
	SharedClient = NewClient(&http.Client{}, baseURL, options)
}
