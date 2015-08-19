// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package elbiface provides an interface for the Elastic Load Balancing.
package elbiface

import (
	"github.com/tetrafolium/aws-sdk-go/aws/service"
	"github.com/tetrafolium/aws-sdk-go/service/elb"
)

// ELBAPI is the interface type for elb.ELB.
type ELBAPI interface {
	AddTagsRequest(*elb.AddTagsInput) (*service.Request, *elb.AddTagsOutput)

	AddTags(*elb.AddTagsInput) (*elb.AddTagsOutput, error)

	ApplySecurityGroupsToLoadBalancerRequest(*elb.ApplySecurityGroupsToLoadBalancerInput) (*service.Request, *elb.ApplySecurityGroupsToLoadBalancerOutput)

	ApplySecurityGroupsToLoadBalancer(*elb.ApplySecurityGroupsToLoadBalancerInput) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error)

	AttachLoadBalancerToSubnetsRequest(*elb.AttachLoadBalancerToSubnetsInput) (*service.Request, *elb.AttachLoadBalancerToSubnetsOutput)

	AttachLoadBalancerToSubnets(*elb.AttachLoadBalancerToSubnetsInput) (*elb.AttachLoadBalancerToSubnetsOutput, error)

	ConfigureHealthCheckRequest(*elb.ConfigureHealthCheckInput) (*service.Request, *elb.ConfigureHealthCheckOutput)

	ConfigureHealthCheck(*elb.ConfigureHealthCheckInput) (*elb.ConfigureHealthCheckOutput, error)

	CreateAppCookieStickinessPolicyRequest(*elb.CreateAppCookieStickinessPolicyInput) (*service.Request, *elb.CreateAppCookieStickinessPolicyOutput)

	CreateAppCookieStickinessPolicy(*elb.CreateAppCookieStickinessPolicyInput) (*elb.CreateAppCookieStickinessPolicyOutput, error)

	CreateLBCookieStickinessPolicyRequest(*elb.CreateLBCookieStickinessPolicyInput) (*service.Request, *elb.CreateLBCookieStickinessPolicyOutput)

	CreateLBCookieStickinessPolicy(*elb.CreateLBCookieStickinessPolicyInput) (*elb.CreateLBCookieStickinessPolicyOutput, error)

	CreateLoadBalancerRequest(*elb.CreateLoadBalancerInput) (*service.Request, *elb.CreateLoadBalancerOutput)

	CreateLoadBalancer(*elb.CreateLoadBalancerInput) (*elb.CreateLoadBalancerOutput, error)

	CreateLoadBalancerListenersRequest(*elb.CreateLoadBalancerListenersInput) (*service.Request, *elb.CreateLoadBalancerListenersOutput)

	CreateLoadBalancerListeners(*elb.CreateLoadBalancerListenersInput) (*elb.CreateLoadBalancerListenersOutput, error)

	CreateLoadBalancerPolicyRequest(*elb.CreateLoadBalancerPolicyInput) (*service.Request, *elb.CreateLoadBalancerPolicyOutput)

	CreateLoadBalancerPolicy(*elb.CreateLoadBalancerPolicyInput) (*elb.CreateLoadBalancerPolicyOutput, error)

	DeleteLoadBalancerRequest(*elb.DeleteLoadBalancerInput) (*service.Request, *elb.DeleteLoadBalancerOutput)

	DeleteLoadBalancer(*elb.DeleteLoadBalancerInput) (*elb.DeleteLoadBalancerOutput, error)

	DeleteLoadBalancerListenersRequest(*elb.DeleteLoadBalancerListenersInput) (*service.Request, *elb.DeleteLoadBalancerListenersOutput)

	DeleteLoadBalancerListeners(*elb.DeleteLoadBalancerListenersInput) (*elb.DeleteLoadBalancerListenersOutput, error)

	DeleteLoadBalancerPolicyRequest(*elb.DeleteLoadBalancerPolicyInput) (*service.Request, *elb.DeleteLoadBalancerPolicyOutput)

	DeleteLoadBalancerPolicy(*elb.DeleteLoadBalancerPolicyInput) (*elb.DeleteLoadBalancerPolicyOutput, error)

	DeregisterInstancesFromLoadBalancerRequest(*elb.DeregisterInstancesFromLoadBalancerInput) (*service.Request, *elb.DeregisterInstancesFromLoadBalancerOutput)

	DeregisterInstancesFromLoadBalancer(*elb.DeregisterInstancesFromLoadBalancerInput) (*elb.DeregisterInstancesFromLoadBalancerOutput, error)

	DescribeInstanceHealthRequest(*elb.DescribeInstanceHealthInput) (*service.Request, *elb.DescribeInstanceHealthOutput)

	DescribeInstanceHealth(*elb.DescribeInstanceHealthInput) (*elb.DescribeInstanceHealthOutput, error)

	DescribeLoadBalancerAttributesRequest(*elb.DescribeLoadBalancerAttributesInput) (*service.Request, *elb.DescribeLoadBalancerAttributesOutput)

	DescribeLoadBalancerAttributes(*elb.DescribeLoadBalancerAttributesInput) (*elb.DescribeLoadBalancerAttributesOutput, error)

	DescribeLoadBalancerPoliciesRequest(*elb.DescribeLoadBalancerPoliciesInput) (*service.Request, *elb.DescribeLoadBalancerPoliciesOutput)

	DescribeLoadBalancerPolicies(*elb.DescribeLoadBalancerPoliciesInput) (*elb.DescribeLoadBalancerPoliciesOutput, error)

	DescribeLoadBalancerPolicyTypesRequest(*elb.DescribeLoadBalancerPolicyTypesInput) (*service.Request, *elb.DescribeLoadBalancerPolicyTypesOutput)

	DescribeLoadBalancerPolicyTypes(*elb.DescribeLoadBalancerPolicyTypesInput) (*elb.DescribeLoadBalancerPolicyTypesOutput, error)

	DescribeLoadBalancersRequest(*elb.DescribeLoadBalancersInput) (*service.Request, *elb.DescribeLoadBalancersOutput)

	DescribeLoadBalancers(*elb.DescribeLoadBalancersInput) (*elb.DescribeLoadBalancersOutput, error)

	DescribeLoadBalancersPages(*elb.DescribeLoadBalancersInput, func(*elb.DescribeLoadBalancersOutput, bool) bool) error

	DescribeTagsRequest(*elb.DescribeTagsInput) (*service.Request, *elb.DescribeTagsOutput)

	DescribeTags(*elb.DescribeTagsInput) (*elb.DescribeTagsOutput, error)

	DetachLoadBalancerFromSubnetsRequest(*elb.DetachLoadBalancerFromSubnetsInput) (*service.Request, *elb.DetachLoadBalancerFromSubnetsOutput)

	DetachLoadBalancerFromSubnets(*elb.DetachLoadBalancerFromSubnetsInput) (*elb.DetachLoadBalancerFromSubnetsOutput, error)

	DisableAvailabilityZonesForLoadBalancerRequest(*elb.DisableAvailabilityZonesForLoadBalancerInput) (*service.Request, *elb.DisableAvailabilityZonesForLoadBalancerOutput)

	DisableAvailabilityZonesForLoadBalancer(*elb.DisableAvailabilityZonesForLoadBalancerInput) (*elb.DisableAvailabilityZonesForLoadBalancerOutput, error)

	EnableAvailabilityZonesForLoadBalancerRequest(*elb.EnableAvailabilityZonesForLoadBalancerInput) (*service.Request, *elb.EnableAvailabilityZonesForLoadBalancerOutput)

	EnableAvailabilityZonesForLoadBalancer(*elb.EnableAvailabilityZonesForLoadBalancerInput) (*elb.EnableAvailabilityZonesForLoadBalancerOutput, error)

	ModifyLoadBalancerAttributesRequest(*elb.ModifyLoadBalancerAttributesInput) (*service.Request, *elb.ModifyLoadBalancerAttributesOutput)

	ModifyLoadBalancerAttributes(*elb.ModifyLoadBalancerAttributesInput) (*elb.ModifyLoadBalancerAttributesOutput, error)

	RegisterInstancesWithLoadBalancerRequest(*elb.RegisterInstancesWithLoadBalancerInput) (*service.Request, *elb.RegisterInstancesWithLoadBalancerOutput)

	RegisterInstancesWithLoadBalancer(*elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error)

	RemoveTagsRequest(*elb.RemoveTagsInput) (*service.Request, *elb.RemoveTagsOutput)

	RemoveTags(*elb.RemoveTagsInput) (*elb.RemoveTagsOutput, error)

	SetLoadBalancerListenerSSLCertificateRequest(*elb.SetLoadBalancerListenerSSLCertificateInput) (*service.Request, *elb.SetLoadBalancerListenerSSLCertificateOutput)

	SetLoadBalancerListenerSSLCertificate(*elb.SetLoadBalancerListenerSSLCertificateInput) (*elb.SetLoadBalancerListenerSSLCertificateOutput, error)

	SetLoadBalancerPoliciesForBackendServerRequest(*elb.SetLoadBalancerPoliciesForBackendServerInput) (*service.Request, *elb.SetLoadBalancerPoliciesForBackendServerOutput)

	SetLoadBalancerPoliciesForBackendServer(*elb.SetLoadBalancerPoliciesForBackendServerInput) (*elb.SetLoadBalancerPoliciesForBackendServerOutput, error)

	SetLoadBalancerPoliciesOfListenerRequest(*elb.SetLoadBalancerPoliciesOfListenerInput) (*service.Request, *elb.SetLoadBalancerPoliciesOfListenerOutput)

	SetLoadBalancerPoliciesOfListener(*elb.SetLoadBalancerPoliciesOfListenerInput) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error)
}
