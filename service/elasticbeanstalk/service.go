// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package elasticbeanstalk

import (
	"github.com/tetrafolium/aws-sdk-go/aws"
	"github.com/tetrafolium/aws-sdk-go/aws/defaults"
	"github.com/tetrafolium/aws-sdk-go/aws/service"
	"github.com/tetrafolium/aws-sdk-go/internal/protocol/query"
	"github.com/tetrafolium/aws-sdk-go/internal/signer/v4"
)

// This is the AWS Elastic Beanstalk API Reference. This guide provides detailed
// information about AWS Elastic Beanstalk actions, data types, parameters,
// and errors.
//
// AWS Elastic Beanstalk is a tool that makes it easy for you to create, deploy,
// and manage scalable, fault-tolerant applications running on Amazon Web Services
// cloud resources.
//
//  For more information about this product, go to the AWS Elastic Beanstalk
// (http://aws.amazon.com/elasticbeanstalk/) details page. The location of the
// latest AWS Elastic Beanstalk WSDL is http://elasticbeanstalk.s3.amazonaws.com/doc/2010-12-01/AWSElasticBeanstalk.wsdl
// (http://elasticbeanstalk.s3.amazonaws.com/doc/2010-12-01/AWSElasticBeanstalk.wsdl).
// To install the Software Development Kits (SDKs), Integrated Development Environment
// (IDE) Toolkits, and command line tools that enable you to access the API,
// go to Tools for Amazon Web Services (https://aws.amazon.com/tools/).
//
// Endpoints
//
// For a list of region-specific endpoints that AWS Elastic Beanstalk supports,
// go to Regions and Endpoints (http://docs.aws.amazon.com/general/latest/gr/rande.html#elasticbeanstalk_region)
// in the Amazon Web Services Glossary.
type ElasticBeanstalk struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*service.Request)

// New returns a new ElasticBeanstalk client.
func New(config *aws.Config) *ElasticBeanstalk {
	service := &service.Service{
		Config:      defaults.DefaultConfig.Merge(config),
		ServiceName: "elasticbeanstalk",
		APIVersion:  "2010-12-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &ElasticBeanstalk{service}
}

// newRequest creates a new request for a ElasticBeanstalk operation and runs any
// custom request initialization.
func (c *ElasticBeanstalk) newRequest(op *service.Operation, params, data interface{}) *service.Request {
	req := service.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
