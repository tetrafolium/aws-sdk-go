// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package ec2

import (
	"github.com/tetrafolium/aws-sdk-go/aws"
	"github.com/tetrafolium/aws-sdk-go/aws/defaults"
	"github.com/tetrafolium/aws-sdk-go/aws/service"
	"github.com/tetrafolium/aws-sdk-go/internal/protocol/ec2query"
	"github.com/tetrafolium/aws-sdk-go/internal/signer/v4"
)

// Amazon Elastic Compute Cloud (Amazon EC2) provides resizable computing capacity
// in the Amazon Web Services (AWS) cloud. Using Amazon EC2 eliminates your
// need to invest in hardware up front, so you can develop and deploy applications
// faster.
type EC2 struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*service.Request)

// New returns a new EC2 client.
func New(config *aws.Config) *EC2 {
	service := &service.Service{
		Config:      defaults.DefaultConfig.Merge(config),
		ServiceName: "ec2",
		APIVersion:  "2015-04-15",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(ec2query.Build)
	service.Handlers.Unmarshal.PushBack(ec2query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(ec2query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(ec2query.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &EC2{service}
}

// newRequest creates a new request for a EC2 operation and runs any
// custom request initialization.
func (c *EC2) newRequest(op *service.Operation, params, data interface{}) *service.Request {
	req := service.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
