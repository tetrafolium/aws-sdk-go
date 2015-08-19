//Package sqs provides gucumber integration tests suppport.
package sqs

import (
	"github.com/tetrafolium/aws-sdk-go/internal/features/shared"
	"github.com/tetrafolium/aws-sdk-go/service/sqs"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@sqs", func() {
		World["client"] = sqs.New(nil)
	})
}
