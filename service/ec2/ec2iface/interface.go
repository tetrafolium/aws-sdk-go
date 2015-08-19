// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package ec2iface provides an interface for the Amazon Elastic Compute Cloud.
package ec2iface

import (
	"github.com/tetrafolium/aws-sdk-go/aws/service"
	"github.com/tetrafolium/aws-sdk-go/service/ec2"
)

// EC2API is the interface type for ec2.EC2.
type EC2API interface {
	AcceptVpcPeeringConnectionRequest(*ec2.AcceptVpcPeeringConnectionInput) (*service.Request, *ec2.AcceptVpcPeeringConnectionOutput)

	AcceptVpcPeeringConnection(*ec2.AcceptVpcPeeringConnectionInput) (*ec2.AcceptVpcPeeringConnectionOutput, error)

	AllocateAddressRequest(*ec2.AllocateAddressInput) (*service.Request, *ec2.AllocateAddressOutput)

	AllocateAddress(*ec2.AllocateAddressInput) (*ec2.AllocateAddressOutput, error)

	AssignPrivateIpAddressesRequest(*ec2.AssignPrivateIpAddressesInput) (*service.Request, *ec2.AssignPrivateIpAddressesOutput)

	AssignPrivateIpAddresses(*ec2.AssignPrivateIpAddressesInput) (*ec2.AssignPrivateIpAddressesOutput, error)

	AssociateAddressRequest(*ec2.AssociateAddressInput) (*service.Request, *ec2.AssociateAddressOutput)

	AssociateAddress(*ec2.AssociateAddressInput) (*ec2.AssociateAddressOutput, error)

	AssociateDhcpOptionsRequest(*ec2.AssociateDhcpOptionsInput) (*service.Request, *ec2.AssociateDhcpOptionsOutput)

	AssociateDhcpOptions(*ec2.AssociateDhcpOptionsInput) (*ec2.AssociateDhcpOptionsOutput, error)

	AssociateRouteTableRequest(*ec2.AssociateRouteTableInput) (*service.Request, *ec2.AssociateRouteTableOutput)

	AssociateRouteTable(*ec2.AssociateRouteTableInput) (*ec2.AssociateRouteTableOutput, error)

	AttachClassicLinkVpcRequest(*ec2.AttachClassicLinkVpcInput) (*service.Request, *ec2.AttachClassicLinkVpcOutput)

	AttachClassicLinkVpc(*ec2.AttachClassicLinkVpcInput) (*ec2.AttachClassicLinkVpcOutput, error)

	AttachInternetGatewayRequest(*ec2.AttachInternetGatewayInput) (*service.Request, *ec2.AttachInternetGatewayOutput)

	AttachInternetGateway(*ec2.AttachInternetGatewayInput) (*ec2.AttachInternetGatewayOutput, error)

	AttachNetworkInterfaceRequest(*ec2.AttachNetworkInterfaceInput) (*service.Request, *ec2.AttachNetworkInterfaceOutput)

	AttachNetworkInterface(*ec2.AttachNetworkInterfaceInput) (*ec2.AttachNetworkInterfaceOutput, error)

	AttachVolumeRequest(*ec2.AttachVolumeInput) (*service.Request, *ec2.VolumeAttachment)

	AttachVolume(*ec2.AttachVolumeInput) (*ec2.VolumeAttachment, error)

	AttachVpnGatewayRequest(*ec2.AttachVpnGatewayInput) (*service.Request, *ec2.AttachVpnGatewayOutput)

	AttachVpnGateway(*ec2.AttachVpnGatewayInput) (*ec2.AttachVpnGatewayOutput, error)

	AuthorizeSecurityGroupEgressRequest(*ec2.AuthorizeSecurityGroupEgressInput) (*service.Request, *ec2.AuthorizeSecurityGroupEgressOutput)

	AuthorizeSecurityGroupEgress(*ec2.AuthorizeSecurityGroupEgressInput) (*ec2.AuthorizeSecurityGroupEgressOutput, error)

	AuthorizeSecurityGroupIngressRequest(*ec2.AuthorizeSecurityGroupIngressInput) (*service.Request, *ec2.AuthorizeSecurityGroupIngressOutput)

	AuthorizeSecurityGroupIngress(*ec2.AuthorizeSecurityGroupIngressInput) (*ec2.AuthorizeSecurityGroupIngressOutput, error)

	BundleInstanceRequest(*ec2.BundleInstanceInput) (*service.Request, *ec2.BundleInstanceOutput)

	BundleInstance(*ec2.BundleInstanceInput) (*ec2.BundleInstanceOutput, error)

	CancelBundleTaskRequest(*ec2.CancelBundleTaskInput) (*service.Request, *ec2.CancelBundleTaskOutput)

	CancelBundleTask(*ec2.CancelBundleTaskInput) (*ec2.CancelBundleTaskOutput, error)

	CancelConversionTaskRequest(*ec2.CancelConversionTaskInput) (*service.Request, *ec2.CancelConversionTaskOutput)

	CancelConversionTask(*ec2.CancelConversionTaskInput) (*ec2.CancelConversionTaskOutput, error)

	CancelExportTaskRequest(*ec2.CancelExportTaskInput) (*service.Request, *ec2.CancelExportTaskOutput)

	CancelExportTask(*ec2.CancelExportTaskInput) (*ec2.CancelExportTaskOutput, error)

	CancelImportTaskRequest(*ec2.CancelImportTaskInput) (*service.Request, *ec2.CancelImportTaskOutput)

	CancelImportTask(*ec2.CancelImportTaskInput) (*ec2.CancelImportTaskOutput, error)

	CancelReservedInstancesListingRequest(*ec2.CancelReservedInstancesListingInput) (*service.Request, *ec2.CancelReservedInstancesListingOutput)

	CancelReservedInstancesListing(*ec2.CancelReservedInstancesListingInput) (*ec2.CancelReservedInstancesListingOutput, error)

	CancelSpotFleetRequestsRequest(*ec2.CancelSpotFleetRequestsInput) (*service.Request, *ec2.CancelSpotFleetRequestsOutput)

	CancelSpotFleetRequests(*ec2.CancelSpotFleetRequestsInput) (*ec2.CancelSpotFleetRequestsOutput, error)

	CancelSpotInstanceRequestsRequest(*ec2.CancelSpotInstanceRequestsInput) (*service.Request, *ec2.CancelSpotInstanceRequestsOutput)

	CancelSpotInstanceRequests(*ec2.CancelSpotInstanceRequestsInput) (*ec2.CancelSpotInstanceRequestsOutput, error)

	ConfirmProductInstanceRequest(*ec2.ConfirmProductInstanceInput) (*service.Request, *ec2.ConfirmProductInstanceOutput)

	ConfirmProductInstance(*ec2.ConfirmProductInstanceInput) (*ec2.ConfirmProductInstanceOutput, error)

	CopyImageRequest(*ec2.CopyImageInput) (*service.Request, *ec2.CopyImageOutput)

	CopyImage(*ec2.CopyImageInput) (*ec2.CopyImageOutput, error)

	CopySnapshotRequest(*ec2.CopySnapshotInput) (*service.Request, *ec2.CopySnapshotOutput)

	CopySnapshot(*ec2.CopySnapshotInput) (*ec2.CopySnapshotOutput, error)

	CreateCustomerGatewayRequest(*ec2.CreateCustomerGatewayInput) (*service.Request, *ec2.CreateCustomerGatewayOutput)

	CreateCustomerGateway(*ec2.CreateCustomerGatewayInput) (*ec2.CreateCustomerGatewayOutput, error)

	CreateDhcpOptionsRequest(*ec2.CreateDhcpOptionsInput) (*service.Request, *ec2.CreateDhcpOptionsOutput)

	CreateDhcpOptions(*ec2.CreateDhcpOptionsInput) (*ec2.CreateDhcpOptionsOutput, error)

	CreateFlowLogsRequest(*ec2.CreateFlowLogsInput) (*service.Request, *ec2.CreateFlowLogsOutput)

	CreateFlowLogs(*ec2.CreateFlowLogsInput) (*ec2.CreateFlowLogsOutput, error)

	CreateImageRequest(*ec2.CreateImageInput) (*service.Request, *ec2.CreateImageOutput)

	CreateImage(*ec2.CreateImageInput) (*ec2.CreateImageOutput, error)

	CreateInstanceExportTaskRequest(*ec2.CreateInstanceExportTaskInput) (*service.Request, *ec2.CreateInstanceExportTaskOutput)

	CreateInstanceExportTask(*ec2.CreateInstanceExportTaskInput) (*ec2.CreateInstanceExportTaskOutput, error)

	CreateInternetGatewayRequest(*ec2.CreateInternetGatewayInput) (*service.Request, *ec2.CreateInternetGatewayOutput)

	CreateInternetGateway(*ec2.CreateInternetGatewayInput) (*ec2.CreateInternetGatewayOutput, error)

	CreateKeyPairRequest(*ec2.CreateKeyPairInput) (*service.Request, *ec2.CreateKeyPairOutput)

	CreateKeyPair(*ec2.CreateKeyPairInput) (*ec2.CreateKeyPairOutput, error)

	CreateNetworkAclRequest(*ec2.CreateNetworkAclInput) (*service.Request, *ec2.CreateNetworkAclOutput)

	CreateNetworkAcl(*ec2.CreateNetworkAclInput) (*ec2.CreateNetworkAclOutput, error)

	CreateNetworkAclEntryRequest(*ec2.CreateNetworkAclEntryInput) (*service.Request, *ec2.CreateNetworkAclEntryOutput)

	CreateNetworkAclEntry(*ec2.CreateNetworkAclEntryInput) (*ec2.CreateNetworkAclEntryOutput, error)

	CreateNetworkInterfaceRequest(*ec2.CreateNetworkInterfaceInput) (*service.Request, *ec2.CreateNetworkInterfaceOutput)

	CreateNetworkInterface(*ec2.CreateNetworkInterfaceInput) (*ec2.CreateNetworkInterfaceOutput, error)

	CreatePlacementGroupRequest(*ec2.CreatePlacementGroupInput) (*service.Request, *ec2.CreatePlacementGroupOutput)

	CreatePlacementGroup(*ec2.CreatePlacementGroupInput) (*ec2.CreatePlacementGroupOutput, error)

	CreateReservedInstancesListingRequest(*ec2.CreateReservedInstancesListingInput) (*service.Request, *ec2.CreateReservedInstancesListingOutput)

	CreateReservedInstancesListing(*ec2.CreateReservedInstancesListingInput) (*ec2.CreateReservedInstancesListingOutput, error)

	CreateRouteRequest(*ec2.CreateRouteInput) (*service.Request, *ec2.CreateRouteOutput)

	CreateRoute(*ec2.CreateRouteInput) (*ec2.CreateRouteOutput, error)

	CreateRouteTableRequest(*ec2.CreateRouteTableInput) (*service.Request, *ec2.CreateRouteTableOutput)

	CreateRouteTable(*ec2.CreateRouteTableInput) (*ec2.CreateRouteTableOutput, error)

	CreateSecurityGroupRequest(*ec2.CreateSecurityGroupInput) (*service.Request, *ec2.CreateSecurityGroupOutput)

	CreateSecurityGroup(*ec2.CreateSecurityGroupInput) (*ec2.CreateSecurityGroupOutput, error)

	CreateSnapshotRequest(*ec2.CreateSnapshotInput) (*service.Request, *ec2.Snapshot)

	CreateSnapshot(*ec2.CreateSnapshotInput) (*ec2.Snapshot, error)

	CreateSpotDatafeedSubscriptionRequest(*ec2.CreateSpotDatafeedSubscriptionInput) (*service.Request, *ec2.CreateSpotDatafeedSubscriptionOutput)

	CreateSpotDatafeedSubscription(*ec2.CreateSpotDatafeedSubscriptionInput) (*ec2.CreateSpotDatafeedSubscriptionOutput, error)

	CreateSubnetRequest(*ec2.CreateSubnetInput) (*service.Request, *ec2.CreateSubnetOutput)

	CreateSubnet(*ec2.CreateSubnetInput) (*ec2.CreateSubnetOutput, error)

	CreateTagsRequest(*ec2.CreateTagsInput) (*service.Request, *ec2.CreateTagsOutput)

	CreateTags(*ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error)

	CreateVolumeRequest(*ec2.CreateVolumeInput) (*service.Request, *ec2.Volume)

	CreateVolume(*ec2.CreateVolumeInput) (*ec2.Volume, error)

	CreateVpcRequest(*ec2.CreateVpcInput) (*service.Request, *ec2.CreateVpcOutput)

	CreateVpc(*ec2.CreateVpcInput) (*ec2.CreateVpcOutput, error)

	CreateVpcEndpointRequest(*ec2.CreateVpcEndpointInput) (*service.Request, *ec2.CreateVpcEndpointOutput)

	CreateVpcEndpoint(*ec2.CreateVpcEndpointInput) (*ec2.CreateVpcEndpointOutput, error)

	CreateVpcPeeringConnectionRequest(*ec2.CreateVpcPeeringConnectionInput) (*service.Request, *ec2.CreateVpcPeeringConnectionOutput)

	CreateVpcPeeringConnection(*ec2.CreateVpcPeeringConnectionInput) (*ec2.CreateVpcPeeringConnectionOutput, error)

	CreateVpnConnectionRequest(*ec2.CreateVpnConnectionInput) (*service.Request, *ec2.CreateVpnConnectionOutput)

	CreateVpnConnection(*ec2.CreateVpnConnectionInput) (*ec2.CreateVpnConnectionOutput, error)

	CreateVpnConnectionRouteRequest(*ec2.CreateVpnConnectionRouteInput) (*service.Request, *ec2.CreateVpnConnectionRouteOutput)

	CreateVpnConnectionRoute(*ec2.CreateVpnConnectionRouteInput) (*ec2.CreateVpnConnectionRouteOutput, error)

	CreateVpnGatewayRequest(*ec2.CreateVpnGatewayInput) (*service.Request, *ec2.CreateVpnGatewayOutput)

	CreateVpnGateway(*ec2.CreateVpnGatewayInput) (*ec2.CreateVpnGatewayOutput, error)

	DeleteCustomerGatewayRequest(*ec2.DeleteCustomerGatewayInput) (*service.Request, *ec2.DeleteCustomerGatewayOutput)

	DeleteCustomerGateway(*ec2.DeleteCustomerGatewayInput) (*ec2.DeleteCustomerGatewayOutput, error)

	DeleteDhcpOptionsRequest(*ec2.DeleteDhcpOptionsInput) (*service.Request, *ec2.DeleteDhcpOptionsOutput)

	DeleteDhcpOptions(*ec2.DeleteDhcpOptionsInput) (*ec2.DeleteDhcpOptionsOutput, error)

	DeleteFlowLogsRequest(*ec2.DeleteFlowLogsInput) (*service.Request, *ec2.DeleteFlowLogsOutput)

	DeleteFlowLogs(*ec2.DeleteFlowLogsInput) (*ec2.DeleteFlowLogsOutput, error)

	DeleteInternetGatewayRequest(*ec2.DeleteInternetGatewayInput) (*service.Request, *ec2.DeleteInternetGatewayOutput)

	DeleteInternetGateway(*ec2.DeleteInternetGatewayInput) (*ec2.DeleteInternetGatewayOutput, error)

	DeleteKeyPairRequest(*ec2.DeleteKeyPairInput) (*service.Request, *ec2.DeleteKeyPairOutput)

	DeleteKeyPair(*ec2.DeleteKeyPairInput) (*ec2.DeleteKeyPairOutput, error)

	DeleteNetworkAclRequest(*ec2.DeleteNetworkAclInput) (*service.Request, *ec2.DeleteNetworkAclOutput)

	DeleteNetworkAcl(*ec2.DeleteNetworkAclInput) (*ec2.DeleteNetworkAclOutput, error)

	DeleteNetworkAclEntryRequest(*ec2.DeleteNetworkAclEntryInput) (*service.Request, *ec2.DeleteNetworkAclEntryOutput)

	DeleteNetworkAclEntry(*ec2.DeleteNetworkAclEntryInput) (*ec2.DeleteNetworkAclEntryOutput, error)

	DeleteNetworkInterfaceRequest(*ec2.DeleteNetworkInterfaceInput) (*service.Request, *ec2.DeleteNetworkInterfaceOutput)

	DeleteNetworkInterface(*ec2.DeleteNetworkInterfaceInput) (*ec2.DeleteNetworkInterfaceOutput, error)

	DeletePlacementGroupRequest(*ec2.DeletePlacementGroupInput) (*service.Request, *ec2.DeletePlacementGroupOutput)

	DeletePlacementGroup(*ec2.DeletePlacementGroupInput) (*ec2.DeletePlacementGroupOutput, error)

	DeleteRouteRequest(*ec2.DeleteRouteInput) (*service.Request, *ec2.DeleteRouteOutput)

	DeleteRoute(*ec2.DeleteRouteInput) (*ec2.DeleteRouteOutput, error)

	DeleteRouteTableRequest(*ec2.DeleteRouteTableInput) (*service.Request, *ec2.DeleteRouteTableOutput)

	DeleteRouteTable(*ec2.DeleteRouteTableInput) (*ec2.DeleteRouteTableOutput, error)

	DeleteSecurityGroupRequest(*ec2.DeleteSecurityGroupInput) (*service.Request, *ec2.DeleteSecurityGroupOutput)

	DeleteSecurityGroup(*ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error)

	DeleteSnapshotRequest(*ec2.DeleteSnapshotInput) (*service.Request, *ec2.DeleteSnapshotOutput)

	DeleteSnapshot(*ec2.DeleteSnapshotInput) (*ec2.DeleteSnapshotOutput, error)

	DeleteSpotDatafeedSubscriptionRequest(*ec2.DeleteSpotDatafeedSubscriptionInput) (*service.Request, *ec2.DeleteSpotDatafeedSubscriptionOutput)

	DeleteSpotDatafeedSubscription(*ec2.DeleteSpotDatafeedSubscriptionInput) (*ec2.DeleteSpotDatafeedSubscriptionOutput, error)

	DeleteSubnetRequest(*ec2.DeleteSubnetInput) (*service.Request, *ec2.DeleteSubnetOutput)

	DeleteSubnet(*ec2.DeleteSubnetInput) (*ec2.DeleteSubnetOutput, error)

	DeleteTagsRequest(*ec2.DeleteTagsInput) (*service.Request, *ec2.DeleteTagsOutput)

	DeleteTags(*ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error)

	DeleteVolumeRequest(*ec2.DeleteVolumeInput) (*service.Request, *ec2.DeleteVolumeOutput)

	DeleteVolume(*ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error)

	DeleteVpcRequest(*ec2.DeleteVpcInput) (*service.Request, *ec2.DeleteVpcOutput)

	DeleteVpc(*ec2.DeleteVpcInput) (*ec2.DeleteVpcOutput, error)

	DeleteVpcEndpointsRequest(*ec2.DeleteVpcEndpointsInput) (*service.Request, *ec2.DeleteVpcEndpointsOutput)

	DeleteVpcEndpoints(*ec2.DeleteVpcEndpointsInput) (*ec2.DeleteVpcEndpointsOutput, error)

	DeleteVpcPeeringConnectionRequest(*ec2.DeleteVpcPeeringConnectionInput) (*service.Request, *ec2.DeleteVpcPeeringConnectionOutput)

	DeleteVpcPeeringConnection(*ec2.DeleteVpcPeeringConnectionInput) (*ec2.DeleteVpcPeeringConnectionOutput, error)

	DeleteVpnConnectionRequest(*ec2.DeleteVpnConnectionInput) (*service.Request, *ec2.DeleteVpnConnectionOutput)

	DeleteVpnConnection(*ec2.DeleteVpnConnectionInput) (*ec2.DeleteVpnConnectionOutput, error)

	DeleteVpnConnectionRouteRequest(*ec2.DeleteVpnConnectionRouteInput) (*service.Request, *ec2.DeleteVpnConnectionRouteOutput)

	DeleteVpnConnectionRoute(*ec2.DeleteVpnConnectionRouteInput) (*ec2.DeleteVpnConnectionRouteOutput, error)

	DeleteVpnGatewayRequest(*ec2.DeleteVpnGatewayInput) (*service.Request, *ec2.DeleteVpnGatewayOutput)

	DeleteVpnGateway(*ec2.DeleteVpnGatewayInput) (*ec2.DeleteVpnGatewayOutput, error)

	DeregisterImageRequest(*ec2.DeregisterImageInput) (*service.Request, *ec2.DeregisterImageOutput)

	DeregisterImage(*ec2.DeregisterImageInput) (*ec2.DeregisterImageOutput, error)

	DescribeAccountAttributesRequest(*ec2.DescribeAccountAttributesInput) (*service.Request, *ec2.DescribeAccountAttributesOutput)

	DescribeAccountAttributes(*ec2.DescribeAccountAttributesInput) (*ec2.DescribeAccountAttributesOutput, error)

	DescribeAddressesRequest(*ec2.DescribeAddressesInput) (*service.Request, *ec2.DescribeAddressesOutput)

	DescribeAddresses(*ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error)

	DescribeAvailabilityZonesRequest(*ec2.DescribeAvailabilityZonesInput) (*service.Request, *ec2.DescribeAvailabilityZonesOutput)

	DescribeAvailabilityZones(*ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error)

	DescribeBundleTasksRequest(*ec2.DescribeBundleTasksInput) (*service.Request, *ec2.DescribeBundleTasksOutput)

	DescribeBundleTasks(*ec2.DescribeBundleTasksInput) (*ec2.DescribeBundleTasksOutput, error)

	DescribeClassicLinkInstancesRequest(*ec2.DescribeClassicLinkInstancesInput) (*service.Request, *ec2.DescribeClassicLinkInstancesOutput)

	DescribeClassicLinkInstances(*ec2.DescribeClassicLinkInstancesInput) (*ec2.DescribeClassicLinkInstancesOutput, error)

	DescribeConversionTasksRequest(*ec2.DescribeConversionTasksInput) (*service.Request, *ec2.DescribeConversionTasksOutput)

	DescribeConversionTasks(*ec2.DescribeConversionTasksInput) (*ec2.DescribeConversionTasksOutput, error)

	DescribeCustomerGatewaysRequest(*ec2.DescribeCustomerGatewaysInput) (*service.Request, *ec2.DescribeCustomerGatewaysOutput)

	DescribeCustomerGateways(*ec2.DescribeCustomerGatewaysInput) (*ec2.DescribeCustomerGatewaysOutput, error)

	DescribeDhcpOptionsRequest(*ec2.DescribeDhcpOptionsInput) (*service.Request, *ec2.DescribeDhcpOptionsOutput)

	DescribeDhcpOptions(*ec2.DescribeDhcpOptionsInput) (*ec2.DescribeDhcpOptionsOutput, error)

	DescribeExportTasksRequest(*ec2.DescribeExportTasksInput) (*service.Request, *ec2.DescribeExportTasksOutput)

	DescribeExportTasks(*ec2.DescribeExportTasksInput) (*ec2.DescribeExportTasksOutput, error)

	DescribeFlowLogsRequest(*ec2.DescribeFlowLogsInput) (*service.Request, *ec2.DescribeFlowLogsOutput)

	DescribeFlowLogs(*ec2.DescribeFlowLogsInput) (*ec2.DescribeFlowLogsOutput, error)

	DescribeImageAttributeRequest(*ec2.DescribeImageAttributeInput) (*service.Request, *ec2.DescribeImageAttributeOutput)

	DescribeImageAttribute(*ec2.DescribeImageAttributeInput) (*ec2.DescribeImageAttributeOutput, error)

	DescribeImagesRequest(*ec2.DescribeImagesInput) (*service.Request, *ec2.DescribeImagesOutput)

	DescribeImages(*ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error)

	DescribeImportImageTasksRequest(*ec2.DescribeImportImageTasksInput) (*service.Request, *ec2.DescribeImportImageTasksOutput)

	DescribeImportImageTasks(*ec2.DescribeImportImageTasksInput) (*ec2.DescribeImportImageTasksOutput, error)

	DescribeImportSnapshotTasksRequest(*ec2.DescribeImportSnapshotTasksInput) (*service.Request, *ec2.DescribeImportSnapshotTasksOutput)

	DescribeImportSnapshotTasks(*ec2.DescribeImportSnapshotTasksInput) (*ec2.DescribeImportSnapshotTasksOutput, error)

	DescribeInstanceAttributeRequest(*ec2.DescribeInstanceAttributeInput) (*service.Request, *ec2.DescribeInstanceAttributeOutput)

	DescribeInstanceAttribute(*ec2.DescribeInstanceAttributeInput) (*ec2.DescribeInstanceAttributeOutput, error)

	DescribeInstanceStatusRequest(*ec2.DescribeInstanceStatusInput) (*service.Request, *ec2.DescribeInstanceStatusOutput)

	DescribeInstanceStatus(*ec2.DescribeInstanceStatusInput) (*ec2.DescribeInstanceStatusOutput, error)

	DescribeInstanceStatusPages(*ec2.DescribeInstanceStatusInput, func(*ec2.DescribeInstanceStatusOutput, bool) bool) error

	DescribeInstancesRequest(*ec2.DescribeInstancesInput) (*service.Request, *ec2.DescribeInstancesOutput)

	DescribeInstances(*ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error)

	DescribeInstancesPages(*ec2.DescribeInstancesInput, func(*ec2.DescribeInstancesOutput, bool) bool) error

	DescribeInternetGatewaysRequest(*ec2.DescribeInternetGatewaysInput) (*service.Request, *ec2.DescribeInternetGatewaysOutput)

	DescribeInternetGateways(*ec2.DescribeInternetGatewaysInput) (*ec2.DescribeInternetGatewaysOutput, error)

	DescribeKeyPairsRequest(*ec2.DescribeKeyPairsInput) (*service.Request, *ec2.DescribeKeyPairsOutput)

	DescribeKeyPairs(*ec2.DescribeKeyPairsInput) (*ec2.DescribeKeyPairsOutput, error)

	DescribeMovingAddressesRequest(*ec2.DescribeMovingAddressesInput) (*service.Request, *ec2.DescribeMovingAddressesOutput)

	DescribeMovingAddresses(*ec2.DescribeMovingAddressesInput) (*ec2.DescribeMovingAddressesOutput, error)

	DescribeNetworkAclsRequest(*ec2.DescribeNetworkAclsInput) (*service.Request, *ec2.DescribeNetworkAclsOutput)

	DescribeNetworkAcls(*ec2.DescribeNetworkAclsInput) (*ec2.DescribeNetworkAclsOutput, error)

	DescribeNetworkInterfaceAttributeRequest(*ec2.DescribeNetworkInterfaceAttributeInput) (*service.Request, *ec2.DescribeNetworkInterfaceAttributeOutput)

	DescribeNetworkInterfaceAttribute(*ec2.DescribeNetworkInterfaceAttributeInput) (*ec2.DescribeNetworkInterfaceAttributeOutput, error)

	DescribeNetworkInterfacesRequest(*ec2.DescribeNetworkInterfacesInput) (*service.Request, *ec2.DescribeNetworkInterfacesOutput)

	DescribeNetworkInterfaces(*ec2.DescribeNetworkInterfacesInput) (*ec2.DescribeNetworkInterfacesOutput, error)

	DescribePlacementGroupsRequest(*ec2.DescribePlacementGroupsInput) (*service.Request, *ec2.DescribePlacementGroupsOutput)

	DescribePlacementGroups(*ec2.DescribePlacementGroupsInput) (*ec2.DescribePlacementGroupsOutput, error)

	DescribePrefixListsRequest(*ec2.DescribePrefixListsInput) (*service.Request, *ec2.DescribePrefixListsOutput)

	DescribePrefixLists(*ec2.DescribePrefixListsInput) (*ec2.DescribePrefixListsOutput, error)

	DescribeRegionsRequest(*ec2.DescribeRegionsInput) (*service.Request, *ec2.DescribeRegionsOutput)

	DescribeRegions(*ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error)

	DescribeReservedInstancesRequest(*ec2.DescribeReservedInstancesInput) (*service.Request, *ec2.DescribeReservedInstancesOutput)

	DescribeReservedInstances(*ec2.DescribeReservedInstancesInput) (*ec2.DescribeReservedInstancesOutput, error)

	DescribeReservedInstancesListingsRequest(*ec2.DescribeReservedInstancesListingsInput) (*service.Request, *ec2.DescribeReservedInstancesListingsOutput)

	DescribeReservedInstancesListings(*ec2.DescribeReservedInstancesListingsInput) (*ec2.DescribeReservedInstancesListingsOutput, error)

	DescribeReservedInstancesModificationsRequest(*ec2.DescribeReservedInstancesModificationsInput) (*service.Request, *ec2.DescribeReservedInstancesModificationsOutput)

	DescribeReservedInstancesModifications(*ec2.DescribeReservedInstancesModificationsInput) (*ec2.DescribeReservedInstancesModificationsOutput, error)

	DescribeReservedInstancesModificationsPages(*ec2.DescribeReservedInstancesModificationsInput, func(*ec2.DescribeReservedInstancesModificationsOutput, bool) bool) error

	DescribeReservedInstancesOfferingsRequest(*ec2.DescribeReservedInstancesOfferingsInput) (*service.Request, *ec2.DescribeReservedInstancesOfferingsOutput)

	DescribeReservedInstancesOfferings(*ec2.DescribeReservedInstancesOfferingsInput) (*ec2.DescribeReservedInstancesOfferingsOutput, error)

	DescribeReservedInstancesOfferingsPages(*ec2.DescribeReservedInstancesOfferingsInput, func(*ec2.DescribeReservedInstancesOfferingsOutput, bool) bool) error

	DescribeRouteTablesRequest(*ec2.DescribeRouteTablesInput) (*service.Request, *ec2.DescribeRouteTablesOutput)

	DescribeRouteTables(*ec2.DescribeRouteTablesInput) (*ec2.DescribeRouteTablesOutput, error)

	DescribeSecurityGroupsRequest(*ec2.DescribeSecurityGroupsInput) (*service.Request, *ec2.DescribeSecurityGroupsOutput)

	DescribeSecurityGroups(*ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error)

	DescribeSnapshotAttributeRequest(*ec2.DescribeSnapshotAttributeInput) (*service.Request, *ec2.DescribeSnapshotAttributeOutput)

	DescribeSnapshotAttribute(*ec2.DescribeSnapshotAttributeInput) (*ec2.DescribeSnapshotAttributeOutput, error)

	DescribeSnapshotsRequest(*ec2.DescribeSnapshotsInput) (*service.Request, *ec2.DescribeSnapshotsOutput)

	DescribeSnapshots(*ec2.DescribeSnapshotsInput) (*ec2.DescribeSnapshotsOutput, error)

	DescribeSnapshotsPages(*ec2.DescribeSnapshotsInput, func(*ec2.DescribeSnapshotsOutput, bool) bool) error

	DescribeSpotDatafeedSubscriptionRequest(*ec2.DescribeSpotDatafeedSubscriptionInput) (*service.Request, *ec2.DescribeSpotDatafeedSubscriptionOutput)

	DescribeSpotDatafeedSubscription(*ec2.DescribeSpotDatafeedSubscriptionInput) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error)

	DescribeSpotFleetInstancesRequest(*ec2.DescribeSpotFleetInstancesInput) (*service.Request, *ec2.DescribeSpotFleetInstancesOutput)

	DescribeSpotFleetInstances(*ec2.DescribeSpotFleetInstancesInput) (*ec2.DescribeSpotFleetInstancesOutput, error)

	DescribeSpotFleetRequestHistoryRequest(*ec2.DescribeSpotFleetRequestHistoryInput) (*service.Request, *ec2.DescribeSpotFleetRequestHistoryOutput)

	DescribeSpotFleetRequestHistory(*ec2.DescribeSpotFleetRequestHistoryInput) (*ec2.DescribeSpotFleetRequestHistoryOutput, error)

	DescribeSpotFleetRequestsRequest(*ec2.DescribeSpotFleetRequestsInput) (*service.Request, *ec2.DescribeSpotFleetRequestsOutput)

	DescribeSpotFleetRequests(*ec2.DescribeSpotFleetRequestsInput) (*ec2.DescribeSpotFleetRequestsOutput, error)

	DescribeSpotInstanceRequestsRequest(*ec2.DescribeSpotInstanceRequestsInput) (*service.Request, *ec2.DescribeSpotInstanceRequestsOutput)

	DescribeSpotInstanceRequests(*ec2.DescribeSpotInstanceRequestsInput) (*ec2.DescribeSpotInstanceRequestsOutput, error)

	DescribeSpotPriceHistoryRequest(*ec2.DescribeSpotPriceHistoryInput) (*service.Request, *ec2.DescribeSpotPriceHistoryOutput)

	DescribeSpotPriceHistory(*ec2.DescribeSpotPriceHistoryInput) (*ec2.DescribeSpotPriceHistoryOutput, error)

	DescribeSpotPriceHistoryPages(*ec2.DescribeSpotPriceHistoryInput, func(*ec2.DescribeSpotPriceHistoryOutput, bool) bool) error

	DescribeSubnetsRequest(*ec2.DescribeSubnetsInput) (*service.Request, *ec2.DescribeSubnetsOutput)

	DescribeSubnets(*ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error)

	DescribeTagsRequest(*ec2.DescribeTagsInput) (*service.Request, *ec2.DescribeTagsOutput)

	DescribeTags(*ec2.DescribeTagsInput) (*ec2.DescribeTagsOutput, error)

	DescribeVolumeAttributeRequest(*ec2.DescribeVolumeAttributeInput) (*service.Request, *ec2.DescribeVolumeAttributeOutput)

	DescribeVolumeAttribute(*ec2.DescribeVolumeAttributeInput) (*ec2.DescribeVolumeAttributeOutput, error)

	DescribeVolumeStatusRequest(*ec2.DescribeVolumeStatusInput) (*service.Request, *ec2.DescribeVolumeStatusOutput)

	DescribeVolumeStatus(*ec2.DescribeVolumeStatusInput) (*ec2.DescribeVolumeStatusOutput, error)

	DescribeVolumeStatusPages(*ec2.DescribeVolumeStatusInput, func(*ec2.DescribeVolumeStatusOutput, bool) bool) error

	DescribeVolumesRequest(*ec2.DescribeVolumesInput) (*service.Request, *ec2.DescribeVolumesOutput)

	DescribeVolumes(*ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error)

	DescribeVolumesPages(*ec2.DescribeVolumesInput, func(*ec2.DescribeVolumesOutput, bool) bool) error

	DescribeVpcAttributeRequest(*ec2.DescribeVpcAttributeInput) (*service.Request, *ec2.DescribeVpcAttributeOutput)

	DescribeVpcAttribute(*ec2.DescribeVpcAttributeInput) (*ec2.DescribeVpcAttributeOutput, error)

	DescribeVpcClassicLinkRequest(*ec2.DescribeVpcClassicLinkInput) (*service.Request, *ec2.DescribeVpcClassicLinkOutput)

	DescribeVpcClassicLink(*ec2.DescribeVpcClassicLinkInput) (*ec2.DescribeVpcClassicLinkOutput, error)

	DescribeVpcEndpointServicesRequest(*ec2.DescribeVpcEndpointServicesInput) (*service.Request, *ec2.DescribeVpcEndpointServicesOutput)

	DescribeVpcEndpointServices(*ec2.DescribeVpcEndpointServicesInput) (*ec2.DescribeVpcEndpointServicesOutput, error)

	DescribeVpcEndpointsRequest(*ec2.DescribeVpcEndpointsInput) (*service.Request, *ec2.DescribeVpcEndpointsOutput)

	DescribeVpcEndpoints(*ec2.DescribeVpcEndpointsInput) (*ec2.DescribeVpcEndpointsOutput, error)

	DescribeVpcPeeringConnectionsRequest(*ec2.DescribeVpcPeeringConnectionsInput) (*service.Request, *ec2.DescribeVpcPeeringConnectionsOutput)

	DescribeVpcPeeringConnections(*ec2.DescribeVpcPeeringConnectionsInput) (*ec2.DescribeVpcPeeringConnectionsOutput, error)

	DescribeVpcsRequest(*ec2.DescribeVpcsInput) (*service.Request, *ec2.DescribeVpcsOutput)

	DescribeVpcs(*ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error)

	DescribeVpnConnectionsRequest(*ec2.DescribeVpnConnectionsInput) (*service.Request, *ec2.DescribeVpnConnectionsOutput)

	DescribeVpnConnections(*ec2.DescribeVpnConnectionsInput) (*ec2.DescribeVpnConnectionsOutput, error)

	DescribeVpnGatewaysRequest(*ec2.DescribeVpnGatewaysInput) (*service.Request, *ec2.DescribeVpnGatewaysOutput)

	DescribeVpnGateways(*ec2.DescribeVpnGatewaysInput) (*ec2.DescribeVpnGatewaysOutput, error)

	DetachClassicLinkVpcRequest(*ec2.DetachClassicLinkVpcInput) (*service.Request, *ec2.DetachClassicLinkVpcOutput)

	DetachClassicLinkVpc(*ec2.DetachClassicLinkVpcInput) (*ec2.DetachClassicLinkVpcOutput, error)

	DetachInternetGatewayRequest(*ec2.DetachInternetGatewayInput) (*service.Request, *ec2.DetachInternetGatewayOutput)

	DetachInternetGateway(*ec2.DetachInternetGatewayInput) (*ec2.DetachInternetGatewayOutput, error)

	DetachNetworkInterfaceRequest(*ec2.DetachNetworkInterfaceInput) (*service.Request, *ec2.DetachNetworkInterfaceOutput)

	DetachNetworkInterface(*ec2.DetachNetworkInterfaceInput) (*ec2.DetachNetworkInterfaceOutput, error)

	DetachVolumeRequest(*ec2.DetachVolumeInput) (*service.Request, *ec2.VolumeAttachment)

	DetachVolume(*ec2.DetachVolumeInput) (*ec2.VolumeAttachment, error)

	DetachVpnGatewayRequest(*ec2.DetachVpnGatewayInput) (*service.Request, *ec2.DetachVpnGatewayOutput)

	DetachVpnGateway(*ec2.DetachVpnGatewayInput) (*ec2.DetachVpnGatewayOutput, error)

	DisableVgwRoutePropagationRequest(*ec2.DisableVgwRoutePropagationInput) (*service.Request, *ec2.DisableVgwRoutePropagationOutput)

	DisableVgwRoutePropagation(*ec2.DisableVgwRoutePropagationInput) (*ec2.DisableVgwRoutePropagationOutput, error)

	DisableVpcClassicLinkRequest(*ec2.DisableVpcClassicLinkInput) (*service.Request, *ec2.DisableVpcClassicLinkOutput)

	DisableVpcClassicLink(*ec2.DisableVpcClassicLinkInput) (*ec2.DisableVpcClassicLinkOutput, error)

	DisassociateAddressRequest(*ec2.DisassociateAddressInput) (*service.Request, *ec2.DisassociateAddressOutput)

	DisassociateAddress(*ec2.DisassociateAddressInput) (*ec2.DisassociateAddressOutput, error)

	DisassociateRouteTableRequest(*ec2.DisassociateRouteTableInput) (*service.Request, *ec2.DisassociateRouteTableOutput)

	DisassociateRouteTable(*ec2.DisassociateRouteTableInput) (*ec2.DisassociateRouteTableOutput, error)

	EnableVgwRoutePropagationRequest(*ec2.EnableVgwRoutePropagationInput) (*service.Request, *ec2.EnableVgwRoutePropagationOutput)

	EnableVgwRoutePropagation(*ec2.EnableVgwRoutePropagationInput) (*ec2.EnableVgwRoutePropagationOutput, error)

	EnableVolumeIORequest(*ec2.EnableVolumeIOInput) (*service.Request, *ec2.EnableVolumeIOOutput)

	EnableVolumeIO(*ec2.EnableVolumeIOInput) (*ec2.EnableVolumeIOOutput, error)

	EnableVpcClassicLinkRequest(*ec2.EnableVpcClassicLinkInput) (*service.Request, *ec2.EnableVpcClassicLinkOutput)

	EnableVpcClassicLink(*ec2.EnableVpcClassicLinkInput) (*ec2.EnableVpcClassicLinkOutput, error)

	GetConsoleOutputRequest(*ec2.GetConsoleOutputInput) (*service.Request, *ec2.GetConsoleOutputOutput)

	GetConsoleOutput(*ec2.GetConsoleOutputInput) (*ec2.GetConsoleOutputOutput, error)

	GetPasswordDataRequest(*ec2.GetPasswordDataInput) (*service.Request, *ec2.GetPasswordDataOutput)

	GetPasswordData(*ec2.GetPasswordDataInput) (*ec2.GetPasswordDataOutput, error)

	ImportImageRequest(*ec2.ImportImageInput) (*service.Request, *ec2.ImportImageOutput)

	ImportImage(*ec2.ImportImageInput) (*ec2.ImportImageOutput, error)

	ImportInstanceRequest(*ec2.ImportInstanceInput) (*service.Request, *ec2.ImportInstanceOutput)

	ImportInstance(*ec2.ImportInstanceInput) (*ec2.ImportInstanceOutput, error)

	ImportKeyPairRequest(*ec2.ImportKeyPairInput) (*service.Request, *ec2.ImportKeyPairOutput)

	ImportKeyPair(*ec2.ImportKeyPairInput) (*ec2.ImportKeyPairOutput, error)

	ImportSnapshotRequest(*ec2.ImportSnapshotInput) (*service.Request, *ec2.ImportSnapshotOutput)

	ImportSnapshot(*ec2.ImportSnapshotInput) (*ec2.ImportSnapshotOutput, error)

	ImportVolumeRequest(*ec2.ImportVolumeInput) (*service.Request, *ec2.ImportVolumeOutput)

	ImportVolume(*ec2.ImportVolumeInput) (*ec2.ImportVolumeOutput, error)

	ModifyImageAttributeRequest(*ec2.ModifyImageAttributeInput) (*service.Request, *ec2.ModifyImageAttributeOutput)

	ModifyImageAttribute(*ec2.ModifyImageAttributeInput) (*ec2.ModifyImageAttributeOutput, error)

	ModifyInstanceAttributeRequest(*ec2.ModifyInstanceAttributeInput) (*service.Request, *ec2.ModifyInstanceAttributeOutput)

	ModifyInstanceAttribute(*ec2.ModifyInstanceAttributeInput) (*ec2.ModifyInstanceAttributeOutput, error)

	ModifyNetworkInterfaceAttributeRequest(*ec2.ModifyNetworkInterfaceAttributeInput) (*service.Request, *ec2.ModifyNetworkInterfaceAttributeOutput)

	ModifyNetworkInterfaceAttribute(*ec2.ModifyNetworkInterfaceAttributeInput) (*ec2.ModifyNetworkInterfaceAttributeOutput, error)

	ModifyReservedInstancesRequest(*ec2.ModifyReservedInstancesInput) (*service.Request, *ec2.ModifyReservedInstancesOutput)

	ModifyReservedInstances(*ec2.ModifyReservedInstancesInput) (*ec2.ModifyReservedInstancesOutput, error)

	ModifySnapshotAttributeRequest(*ec2.ModifySnapshotAttributeInput) (*service.Request, *ec2.ModifySnapshotAttributeOutput)

	ModifySnapshotAttribute(*ec2.ModifySnapshotAttributeInput) (*ec2.ModifySnapshotAttributeOutput, error)

	ModifySubnetAttributeRequest(*ec2.ModifySubnetAttributeInput) (*service.Request, *ec2.ModifySubnetAttributeOutput)

	ModifySubnetAttribute(*ec2.ModifySubnetAttributeInput) (*ec2.ModifySubnetAttributeOutput, error)

	ModifyVolumeAttributeRequest(*ec2.ModifyVolumeAttributeInput) (*service.Request, *ec2.ModifyVolumeAttributeOutput)

	ModifyVolumeAttribute(*ec2.ModifyVolumeAttributeInput) (*ec2.ModifyVolumeAttributeOutput, error)

	ModifyVpcAttributeRequest(*ec2.ModifyVpcAttributeInput) (*service.Request, *ec2.ModifyVpcAttributeOutput)

	ModifyVpcAttribute(*ec2.ModifyVpcAttributeInput) (*ec2.ModifyVpcAttributeOutput, error)

	ModifyVpcEndpointRequest(*ec2.ModifyVpcEndpointInput) (*service.Request, *ec2.ModifyVpcEndpointOutput)

	ModifyVpcEndpoint(*ec2.ModifyVpcEndpointInput) (*ec2.ModifyVpcEndpointOutput, error)

	MonitorInstancesRequest(*ec2.MonitorInstancesInput) (*service.Request, *ec2.MonitorInstancesOutput)

	MonitorInstances(*ec2.MonitorInstancesInput) (*ec2.MonitorInstancesOutput, error)

	MoveAddressToVpcRequest(*ec2.MoveAddressToVpcInput) (*service.Request, *ec2.MoveAddressToVpcOutput)

	MoveAddressToVpc(*ec2.MoveAddressToVpcInput) (*ec2.MoveAddressToVpcOutput, error)

	PurchaseReservedInstancesOfferingRequest(*ec2.PurchaseReservedInstancesOfferingInput) (*service.Request, *ec2.PurchaseReservedInstancesOfferingOutput)

	PurchaseReservedInstancesOffering(*ec2.PurchaseReservedInstancesOfferingInput) (*ec2.PurchaseReservedInstancesOfferingOutput, error)

	RebootInstancesRequest(*ec2.RebootInstancesInput) (*service.Request, *ec2.RebootInstancesOutput)

	RebootInstances(*ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error)

	RegisterImageRequest(*ec2.RegisterImageInput) (*service.Request, *ec2.RegisterImageOutput)

	RegisterImage(*ec2.RegisterImageInput) (*ec2.RegisterImageOutput, error)

	RejectVpcPeeringConnectionRequest(*ec2.RejectVpcPeeringConnectionInput) (*service.Request, *ec2.RejectVpcPeeringConnectionOutput)

	RejectVpcPeeringConnection(*ec2.RejectVpcPeeringConnectionInput) (*ec2.RejectVpcPeeringConnectionOutput, error)

	ReleaseAddressRequest(*ec2.ReleaseAddressInput) (*service.Request, *ec2.ReleaseAddressOutput)

	ReleaseAddress(*ec2.ReleaseAddressInput) (*ec2.ReleaseAddressOutput, error)

	ReplaceNetworkAclAssociationRequest(*ec2.ReplaceNetworkAclAssociationInput) (*service.Request, *ec2.ReplaceNetworkAclAssociationOutput)

	ReplaceNetworkAclAssociation(*ec2.ReplaceNetworkAclAssociationInput) (*ec2.ReplaceNetworkAclAssociationOutput, error)

	ReplaceNetworkAclEntryRequest(*ec2.ReplaceNetworkAclEntryInput) (*service.Request, *ec2.ReplaceNetworkAclEntryOutput)

	ReplaceNetworkAclEntry(*ec2.ReplaceNetworkAclEntryInput) (*ec2.ReplaceNetworkAclEntryOutput, error)

	ReplaceRouteRequest(*ec2.ReplaceRouteInput) (*service.Request, *ec2.ReplaceRouteOutput)

	ReplaceRoute(*ec2.ReplaceRouteInput) (*ec2.ReplaceRouteOutput, error)

	ReplaceRouteTableAssociationRequest(*ec2.ReplaceRouteTableAssociationInput) (*service.Request, *ec2.ReplaceRouteTableAssociationOutput)

	ReplaceRouteTableAssociation(*ec2.ReplaceRouteTableAssociationInput) (*ec2.ReplaceRouteTableAssociationOutput, error)

	ReportInstanceStatusRequest(*ec2.ReportInstanceStatusInput) (*service.Request, *ec2.ReportInstanceStatusOutput)

	ReportInstanceStatus(*ec2.ReportInstanceStatusInput) (*ec2.ReportInstanceStatusOutput, error)

	RequestSpotFleetRequest(*ec2.RequestSpotFleetInput) (*service.Request, *ec2.RequestSpotFleetOutput)

	RequestSpotFleet(*ec2.RequestSpotFleetInput) (*ec2.RequestSpotFleetOutput, error)

	RequestSpotInstancesRequest(*ec2.RequestSpotInstancesInput) (*service.Request, *ec2.RequestSpotInstancesOutput)

	RequestSpotInstances(*ec2.RequestSpotInstancesInput) (*ec2.RequestSpotInstancesOutput, error)

	ResetImageAttributeRequest(*ec2.ResetImageAttributeInput) (*service.Request, *ec2.ResetImageAttributeOutput)

	ResetImageAttribute(*ec2.ResetImageAttributeInput) (*ec2.ResetImageAttributeOutput, error)

	ResetInstanceAttributeRequest(*ec2.ResetInstanceAttributeInput) (*service.Request, *ec2.ResetInstanceAttributeOutput)

	ResetInstanceAttribute(*ec2.ResetInstanceAttributeInput) (*ec2.ResetInstanceAttributeOutput, error)

	ResetNetworkInterfaceAttributeRequest(*ec2.ResetNetworkInterfaceAttributeInput) (*service.Request, *ec2.ResetNetworkInterfaceAttributeOutput)

	ResetNetworkInterfaceAttribute(*ec2.ResetNetworkInterfaceAttributeInput) (*ec2.ResetNetworkInterfaceAttributeOutput, error)

	ResetSnapshotAttributeRequest(*ec2.ResetSnapshotAttributeInput) (*service.Request, *ec2.ResetSnapshotAttributeOutput)

	ResetSnapshotAttribute(*ec2.ResetSnapshotAttributeInput) (*ec2.ResetSnapshotAttributeOutput, error)

	RestoreAddressToClassicRequest(*ec2.RestoreAddressToClassicInput) (*service.Request, *ec2.RestoreAddressToClassicOutput)

	RestoreAddressToClassic(*ec2.RestoreAddressToClassicInput) (*ec2.RestoreAddressToClassicOutput, error)

	RevokeSecurityGroupEgressRequest(*ec2.RevokeSecurityGroupEgressInput) (*service.Request, *ec2.RevokeSecurityGroupEgressOutput)

	RevokeSecurityGroupEgress(*ec2.RevokeSecurityGroupEgressInput) (*ec2.RevokeSecurityGroupEgressOutput, error)

	RevokeSecurityGroupIngressRequest(*ec2.RevokeSecurityGroupIngressInput) (*service.Request, *ec2.RevokeSecurityGroupIngressOutput)

	RevokeSecurityGroupIngress(*ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error)

	RunInstancesRequest(*ec2.RunInstancesInput) (*service.Request, *ec2.Reservation)

	RunInstances(*ec2.RunInstancesInput) (*ec2.Reservation, error)

	StartInstancesRequest(*ec2.StartInstancesInput) (*service.Request, *ec2.StartInstancesOutput)

	StartInstances(*ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error)

	StopInstancesRequest(*ec2.StopInstancesInput) (*service.Request, *ec2.StopInstancesOutput)

	StopInstances(*ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error)

	TerminateInstancesRequest(*ec2.TerminateInstancesInput) (*service.Request, *ec2.TerminateInstancesOutput)

	TerminateInstances(*ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error)

	UnassignPrivateIpAddressesRequest(*ec2.UnassignPrivateIpAddressesInput) (*service.Request, *ec2.UnassignPrivateIpAddressesOutput)

	UnassignPrivateIpAddresses(*ec2.UnassignPrivateIpAddressesInput) (*ec2.UnassignPrivateIpAddressesOutput, error)

	UnmonitorInstancesRequest(*ec2.UnmonitorInstancesInput) (*service.Request, *ec2.UnmonitorInstancesOutput)

	UnmonitorInstances(*ec2.UnmonitorInstancesInput) (*ec2.UnmonitorInstancesOutput, error)
}
