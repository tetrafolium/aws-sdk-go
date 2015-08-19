// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package codedeploy

import (
	"github.com/tetrafolium/aws-sdk-go/aws"
	"github.com/tetrafolium/aws-sdk-go/aws/defaults"
	"github.com/tetrafolium/aws-sdk-go/aws/service"
	"github.com/tetrafolium/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/tetrafolium/aws-sdk-go/internal/signer/v4"
)

// Overview This is the AWS CodeDeploy API Reference. This guide provides descriptions
// of the AWS CodeDeploy APIs. For additional information, see the AWS CodeDeploy
// User Guide (http://docs.aws.amazon.com/codedeploy/latest/userguide).
//
// Using the APIs You can use the AWS CodeDeploy APIs to work with the following
// items:
//
//   Applications are unique identifiers that AWS CodeDeploy uses to ensure
// that the correct combinations of revisions, deployment configurations, and
// deployment groups are being referenced during deployments.
//
// You can use the AWS CodeDeploy APIs to create, delete, get, list, and update
// applications.
//
//   Deployment configurations are sets of deployment rules and deployment
// success and failure conditions that AWS CodeDeploy uses during deployments.
//
// You can use the AWS CodeDeploy APIs to create, delete, get, and list deployment
// configurations.
//
//   Deployment groups are groups of instances to which application revisions
// can be deployed.
//
// You can use the AWS CodeDeploy APIs to create, delete, get, list, and update
// deployment groups.
//
//   Instances represent Amazon EC2 instances to which application revisions
// are deployed. Instances are identified by their Amazon EC2 tags or Auto Scaling
// group names. Instances belong to deployment groups.
//
// You can use the AWS CodeDeploy APIs to get and list instances.
//
//   Deployments represent the process of deploying revisions to instances.
//
// You can use the AWS CodeDeploy APIs to create, get, list, and stop deployments.
//
//   Application revisions are archive files that are stored in Amazon S3 buckets
// or GitHub repositories. These revisions contain source content (such as source
// code, web pages, executable files, any deployment scripts, and similar) along
// with an Application Specification file (AppSpec file). (The AppSpec file
// is unique to AWS CodeDeploy; it defines a series of deployment actions that
// you want AWS CodeDeploy to execute.) An application revision is uniquely
// identified by its Amazon S3 object key and its ETag, version, or both (for
// application revisions that are stored in Amazon S3 buckets) or by its repository
// name and commit ID (for applications revisions that are stored in GitHub
// repositories). Application revisions are deployed through deployment groups.
//
// You can use the AWS CodeDeploy APIs to get, list, and register application
// revisions.
type CodeDeploy struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*service.Request)

// New returns a new CodeDeploy client.
func New(config *aws.Config) *CodeDeploy {
	service := &service.Service{
		Config:       defaults.DefaultConfig.Merge(config),
		ServiceName:  "codedeploy",
		APIVersion:   "2014-10-06",
		JSONVersion:  "1.1",
		TargetPrefix: "CodeDeploy_20141006",
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

	return &CodeDeploy{service}
}

// newRequest creates a new request for a CodeDeploy operation and runs any
// custom request initialization.
func (c *CodeDeploy) newRequest(op *service.Operation, params, data interface{}) *service.Request {
	req := service.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
