// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package configserviceiface_test

import (
	"testing"

	"github.com/tetrafolium/aws-sdk-go/service/configservice"
	"github.com/tetrafolium/aws-sdk-go/service/configservice/configserviceiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*configserviceiface.ConfigServiceAPI)(nil), configservice.New(nil))
}
