package url_helper

import (
	"net/url"
	"strings"
)

func CleanUrl(url *url.URL) *url.URL {
	url.Scheme = strings.TrimSuffix(url.Scheme, ":")
	url.Host = strings.TrimSuffix(url.Host, ":")
	url.Host = strings.TrimSuffix(url.Host, "/")

	// Remove any unnecessary slashes in the path
	url.Path = strings.ReplaceAll(url.Path, "//", "/")
	url.Path = strings.ReplaceAll(url.Path, "//", "/")
	return url
}
