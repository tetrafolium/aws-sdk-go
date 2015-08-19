// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package efs

import (
	"github.com/tetrafolium/aws-sdk-go/aws"
	"github.com/tetrafolium/aws-sdk-go/aws/defaults"
	"github.com/tetrafolium/aws-sdk-go/aws/service"
	"github.com/tetrafolium/aws-sdk-go/internal/protocol/restjson"
	"github.com/tetrafolium/aws-sdk-go/internal/signer/v4"
)

type EFS struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*service.Request)

// New returns a new EFS client.
func New(config *aws.Config) *EFS {
	service := &service.Service{
		Config:      defaults.DefaultConfig.Merge(config),
		ServiceName: "elasticfilesystem",
		APIVersion:  "2015-02-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &EFS{service}
}

// newRequest creates a new request for a EFS operation and runs any
// custom request initialization.
func (c *EFS) newRequest(op *service.Operation, params, data interface{}) *service.Request {
	req := service.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
