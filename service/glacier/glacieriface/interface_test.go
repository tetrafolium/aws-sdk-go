// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package glacieriface_test

import (
	"testing"

	"github.com/tetrafolium/aws-sdk-go/service/glacier"
	"github.com/tetrafolium/aws-sdk-go/service/glacier/glacieriface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*glacieriface.GlacierAPI)(nil), glacier.New(nil))
}
