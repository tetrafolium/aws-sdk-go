//Package autoscaling provides gucumber integration tests suppport.
package autoscaling

import (
	"github.com/tetrafolium/aws-sdk-go/internal/features/shared"
	"github.com/tetrafolium/aws-sdk-go/service/autoscaling"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@autoscaling", func() {
		World["client"] = autoscaling.New(nil)
	})
}
