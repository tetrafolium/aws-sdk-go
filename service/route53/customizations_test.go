package route53_test

import (
	"testing"

	"github.com/tetrafolium/aws-sdk-go/aws"
	"github.com/tetrafolium/aws-sdk-go/internal/util/utilassert"
	"github.com/tetrafolium/aws-sdk-go/service/route53"
)

func TestBuildCorrectURI(t *testing.T) {
	svc := route53.New(nil)
	svc.Handlers.Validate.Clear()
	req, _ := svc.GetHostedZoneRequest(&route53.GetHostedZoneInput{
		Id: aws.String("/hostedzone/ABCDEFG"),
	})

	req.Build()

	utilassert.Match(t, `\/hostedzone\/ABCDEFG$`, req.HTTPRequest.URL.String())
}
