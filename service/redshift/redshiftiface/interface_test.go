// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package redshiftiface_test

import (
	"testing"

	"github.com/tetrafolium/aws-sdk-go/service/redshift"
	"github.com/tetrafolium/aws-sdk-go/service/redshift/redshiftiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*redshiftiface.RedshiftAPI)(nil), redshift.New(nil))
}
