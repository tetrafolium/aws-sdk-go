package route53

import (
	"regexp"

	"github.com/tetrafolium/aws-sdk-go/aws/service"
)

func init() {
	initService = func(s *service.Service) {
		s.Handlers.Build.PushBack(sanitizeURL)
	}
}

var reSanitizeURL = regexp.MustCompile(`\/%2F\w+%2F`)

func sanitizeURL(r *service.Request) {
	r.HTTPRequest.URL.Opaque =
		reSanitizeURL.ReplaceAllString(r.HTTPRequest.URL.Opaque, "/")
}
