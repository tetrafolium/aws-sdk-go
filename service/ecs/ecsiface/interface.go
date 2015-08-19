// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package ecsiface provides an interface for the Amazon EC2 Container Service.
package ecsiface

import (
	"github.com/tetrafolium/aws-sdk-go/aws/service"
	"github.com/tetrafolium/aws-sdk-go/service/ecs"
)

// ECSAPI is the interface type for ecs.ECS.
type ECSAPI interface {
	CreateClusterRequest(*ecs.CreateClusterInput) (*service.Request, *ecs.CreateClusterOutput)

	CreateCluster(*ecs.CreateClusterInput) (*ecs.CreateClusterOutput, error)

	CreateServiceRequest(*ecs.CreateServiceInput) (*service.Request, *ecs.CreateServiceOutput)

	CreateService(*ecs.CreateServiceInput) (*ecs.CreateServiceOutput, error)

	DeleteClusterRequest(*ecs.DeleteClusterInput) (*service.Request, *ecs.DeleteClusterOutput)

	DeleteCluster(*ecs.DeleteClusterInput) (*ecs.DeleteClusterOutput, error)

	DeleteServiceRequest(*ecs.DeleteServiceInput) (*service.Request, *ecs.DeleteServiceOutput)

	DeleteService(*ecs.DeleteServiceInput) (*ecs.DeleteServiceOutput, error)

	DeregisterContainerInstanceRequest(*ecs.DeregisterContainerInstanceInput) (*service.Request, *ecs.DeregisterContainerInstanceOutput)

	DeregisterContainerInstance(*ecs.DeregisterContainerInstanceInput) (*ecs.DeregisterContainerInstanceOutput, error)

	DeregisterTaskDefinitionRequest(*ecs.DeregisterTaskDefinitionInput) (*service.Request, *ecs.DeregisterTaskDefinitionOutput)

	DeregisterTaskDefinition(*ecs.DeregisterTaskDefinitionInput) (*ecs.DeregisterTaskDefinitionOutput, error)

	DescribeClustersRequest(*ecs.DescribeClustersInput) (*service.Request, *ecs.DescribeClustersOutput)

	DescribeClusters(*ecs.DescribeClustersInput) (*ecs.DescribeClustersOutput, error)

	DescribeContainerInstancesRequest(*ecs.DescribeContainerInstancesInput) (*service.Request, *ecs.DescribeContainerInstancesOutput)

	DescribeContainerInstances(*ecs.DescribeContainerInstancesInput) (*ecs.DescribeContainerInstancesOutput, error)

	DescribeServicesRequest(*ecs.DescribeServicesInput) (*service.Request, *ecs.DescribeServicesOutput)

	DescribeServices(*ecs.DescribeServicesInput) (*ecs.DescribeServicesOutput, error)

	DescribeTaskDefinitionRequest(*ecs.DescribeTaskDefinitionInput) (*service.Request, *ecs.DescribeTaskDefinitionOutput)

	DescribeTaskDefinition(*ecs.DescribeTaskDefinitionInput) (*ecs.DescribeTaskDefinitionOutput, error)

	DescribeTasksRequest(*ecs.DescribeTasksInput) (*service.Request, *ecs.DescribeTasksOutput)

	DescribeTasks(*ecs.DescribeTasksInput) (*ecs.DescribeTasksOutput, error)

	DiscoverPollEndpointRequest(*ecs.DiscoverPollEndpointInput) (*service.Request, *ecs.DiscoverPollEndpointOutput)

	DiscoverPollEndpoint(*ecs.DiscoverPollEndpointInput) (*ecs.DiscoverPollEndpointOutput, error)

	ListClustersRequest(*ecs.ListClustersInput) (*service.Request, *ecs.ListClustersOutput)

	ListClusters(*ecs.ListClustersInput) (*ecs.ListClustersOutput, error)

	ListClustersPages(*ecs.ListClustersInput, func(*ecs.ListClustersOutput, bool) bool) error

	ListContainerInstancesRequest(*ecs.ListContainerInstancesInput) (*service.Request, *ecs.ListContainerInstancesOutput)

	ListContainerInstances(*ecs.ListContainerInstancesInput) (*ecs.ListContainerInstancesOutput, error)

	ListContainerInstancesPages(*ecs.ListContainerInstancesInput, func(*ecs.ListContainerInstancesOutput, bool) bool) error

	ListServicesRequest(*ecs.ListServicesInput) (*service.Request, *ecs.ListServicesOutput)

	ListServices(*ecs.ListServicesInput) (*ecs.ListServicesOutput, error)

	ListServicesPages(*ecs.ListServicesInput, func(*ecs.ListServicesOutput, bool) bool) error

	ListTaskDefinitionFamiliesRequest(*ecs.ListTaskDefinitionFamiliesInput) (*service.Request, *ecs.ListTaskDefinitionFamiliesOutput)

	ListTaskDefinitionFamilies(*ecs.ListTaskDefinitionFamiliesInput) (*ecs.ListTaskDefinitionFamiliesOutput, error)

	ListTaskDefinitionFamiliesPages(*ecs.ListTaskDefinitionFamiliesInput, func(*ecs.ListTaskDefinitionFamiliesOutput, bool) bool) error

	ListTaskDefinitionsRequest(*ecs.ListTaskDefinitionsInput) (*service.Request, *ecs.ListTaskDefinitionsOutput)

	ListTaskDefinitions(*ecs.ListTaskDefinitionsInput) (*ecs.ListTaskDefinitionsOutput, error)

	ListTaskDefinitionsPages(*ecs.ListTaskDefinitionsInput, func(*ecs.ListTaskDefinitionsOutput, bool) bool) error

	ListTasksRequest(*ecs.ListTasksInput) (*service.Request, *ecs.ListTasksOutput)

	ListTasks(*ecs.ListTasksInput) (*ecs.ListTasksOutput, error)

	ListTasksPages(*ecs.ListTasksInput, func(*ecs.ListTasksOutput, bool) bool) error

	RegisterContainerInstanceRequest(*ecs.RegisterContainerInstanceInput) (*service.Request, *ecs.RegisterContainerInstanceOutput)

	RegisterContainerInstance(*ecs.RegisterContainerInstanceInput) (*ecs.RegisterContainerInstanceOutput, error)

	RegisterTaskDefinitionRequest(*ecs.RegisterTaskDefinitionInput) (*service.Request, *ecs.RegisterTaskDefinitionOutput)

	RegisterTaskDefinition(*ecs.RegisterTaskDefinitionInput) (*ecs.RegisterTaskDefinitionOutput, error)

	RunTaskRequest(*ecs.RunTaskInput) (*service.Request, *ecs.RunTaskOutput)

	RunTask(*ecs.RunTaskInput) (*ecs.RunTaskOutput, error)

	StartTaskRequest(*ecs.StartTaskInput) (*service.Request, *ecs.StartTaskOutput)

	StartTask(*ecs.StartTaskInput) (*ecs.StartTaskOutput, error)

	StopTaskRequest(*ecs.StopTaskInput) (*service.Request, *ecs.StopTaskOutput)

	StopTask(*ecs.StopTaskInput) (*ecs.StopTaskOutput, error)

	SubmitContainerStateChangeRequest(*ecs.SubmitContainerStateChangeInput) (*service.Request, *ecs.SubmitContainerStateChangeOutput)

	SubmitContainerStateChange(*ecs.SubmitContainerStateChangeInput) (*ecs.SubmitContainerStateChangeOutput, error)

	SubmitTaskStateChangeRequest(*ecs.SubmitTaskStateChangeInput) (*service.Request, *ecs.SubmitTaskStateChangeOutput)

	SubmitTaskStateChange(*ecs.SubmitTaskStateChangeInput) (*ecs.SubmitTaskStateChangeOutput, error)

	UpdateContainerAgentRequest(*ecs.UpdateContainerAgentInput) (*service.Request, *ecs.UpdateContainerAgentOutput)

	UpdateContainerAgent(*ecs.UpdateContainerAgentInput) (*ecs.UpdateContainerAgentOutput, error)

	UpdateServiceRequest(*ecs.UpdateServiceInput) (*service.Request, *ecs.UpdateServiceOutput)

	UpdateService(*ecs.UpdateServiceInput) (*ecs.UpdateServiceOutput, error)
}
