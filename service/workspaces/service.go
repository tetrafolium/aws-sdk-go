// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package workspaces

import (
	"github.com/tetrafolium/aws-sdk-go/aws"
	"github.com/tetrafolium/aws-sdk-go/aws/defaults"
	"github.com/tetrafolium/aws-sdk-go/aws/service"
	"github.com/tetrafolium/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/tetrafolium/aws-sdk-go/internal/signer/v4"
)

// This is the Amazon WorkSpaces API Reference. This guide provides detailed
// information about Amazon WorkSpaces operations, data types, parameters, and
// errors.
type WorkSpaces struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*service.Request)

// New returns a new WorkSpaces client.
func New(config *aws.Config) *WorkSpaces {
	service := &service.Service{
		Config:       defaults.DefaultConfig.Merge(config),
		ServiceName:  "workspaces",
		APIVersion:   "2015-04-08",
		JSONVersion:  "1.1",
		TargetPrefix: "WorkspacesService",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &WorkSpaces{service}
}

// newRequest creates a new request for a WorkSpaces operation and runs any
// custom request initialization.
func (c *WorkSpaces) newRequest(op *service.Operation, params, data interface{}) *service.Request {
	req := service.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
