//Package opsworks provides gucumber integration tests suppport.
package opsworks

import (
	"github.com/tetrafolium/aws-sdk-go/internal/features/shared"
	"github.com/tetrafolium/aws-sdk-go/service/opsworks"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@opsworks", func() {
		World["client"] = opsworks.New(nil)
	})
}
