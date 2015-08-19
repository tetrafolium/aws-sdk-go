//Package cloudhsm provides gucumber integration tests suppport.
package cloudhsm

import (
	"github.com/tetrafolium/aws-sdk-go/internal/features/shared"
	"github.com/tetrafolium/aws-sdk-go/service/cloudhsm"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@cloudhsm", func() {
		World["client"] = cloudhsm.New(nil)
	})
}
