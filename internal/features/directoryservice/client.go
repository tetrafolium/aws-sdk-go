//Package directoryservice provides gucumber integration tests suppport.
package directoryservice

import (
	"github.com/tetrafolium/aws-sdk-go/internal/features/shared"
	"github.com/tetrafolium/aws-sdk-go/service/directoryservice"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@directoryservice", func() {
		// FIXME remove custom region
		World["client"] = directoryservice.New(nil)
	})
}
