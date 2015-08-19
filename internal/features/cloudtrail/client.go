//Package cloudtrail provides gucumber integration tests suppport.
package cloudtrail

import (
	"github.com/tetrafolium/aws-sdk-go/internal/features/shared"
	"github.com/tetrafolium/aws-sdk-go/service/cloudtrail"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@cloudtrail", func() {
		World["client"] = cloudtrail.New(nil)
	})
}
