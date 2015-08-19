// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cognitosynciface provides an interface for the Amazon Cognito Sync.
package cognitosynciface

import (
	"github.com/tetrafolium/aws-sdk-go/aws/service"
	"github.com/tetrafolium/aws-sdk-go/service/cognitosync"
)

// CognitoSyncAPI is the interface type for cognitosync.CognitoSync.
type CognitoSyncAPI interface {
	BulkPublishRequest(*cognitosync.BulkPublishInput) (*service.Request, *cognitosync.BulkPublishOutput)

	BulkPublish(*cognitosync.BulkPublishInput) (*cognitosync.BulkPublishOutput, error)

	DeleteDatasetRequest(*cognitosync.DeleteDatasetInput) (*service.Request, *cognitosync.DeleteDatasetOutput)

	DeleteDataset(*cognitosync.DeleteDatasetInput) (*cognitosync.DeleteDatasetOutput, error)

	DescribeDatasetRequest(*cognitosync.DescribeDatasetInput) (*service.Request, *cognitosync.DescribeDatasetOutput)

	DescribeDataset(*cognitosync.DescribeDatasetInput) (*cognitosync.DescribeDatasetOutput, error)

	DescribeIdentityPoolUsageRequest(*cognitosync.DescribeIdentityPoolUsageInput) (*service.Request, *cognitosync.DescribeIdentityPoolUsageOutput)

	DescribeIdentityPoolUsage(*cognitosync.DescribeIdentityPoolUsageInput) (*cognitosync.DescribeIdentityPoolUsageOutput, error)

	DescribeIdentityUsageRequest(*cognitosync.DescribeIdentityUsageInput) (*service.Request, *cognitosync.DescribeIdentityUsageOutput)

	DescribeIdentityUsage(*cognitosync.DescribeIdentityUsageInput) (*cognitosync.DescribeIdentityUsageOutput, error)

	GetBulkPublishDetailsRequest(*cognitosync.GetBulkPublishDetailsInput) (*service.Request, *cognitosync.GetBulkPublishDetailsOutput)

	GetBulkPublishDetails(*cognitosync.GetBulkPublishDetailsInput) (*cognitosync.GetBulkPublishDetailsOutput, error)

	GetCognitoEventsRequest(*cognitosync.GetCognitoEventsInput) (*service.Request, *cognitosync.GetCognitoEventsOutput)

	GetCognitoEvents(*cognitosync.GetCognitoEventsInput) (*cognitosync.GetCognitoEventsOutput, error)

	GetIdentityPoolConfigurationRequest(*cognitosync.GetIdentityPoolConfigurationInput) (*service.Request, *cognitosync.GetIdentityPoolConfigurationOutput)

	GetIdentityPoolConfiguration(*cognitosync.GetIdentityPoolConfigurationInput) (*cognitosync.GetIdentityPoolConfigurationOutput, error)

	ListDatasetsRequest(*cognitosync.ListDatasetsInput) (*service.Request, *cognitosync.ListDatasetsOutput)

	ListDatasets(*cognitosync.ListDatasetsInput) (*cognitosync.ListDatasetsOutput, error)

	ListIdentityPoolUsageRequest(*cognitosync.ListIdentityPoolUsageInput) (*service.Request, *cognitosync.ListIdentityPoolUsageOutput)

	ListIdentityPoolUsage(*cognitosync.ListIdentityPoolUsageInput) (*cognitosync.ListIdentityPoolUsageOutput, error)

	ListRecordsRequest(*cognitosync.ListRecordsInput) (*service.Request, *cognitosync.ListRecordsOutput)

	ListRecords(*cognitosync.ListRecordsInput) (*cognitosync.ListRecordsOutput, error)

	RegisterDeviceRequest(*cognitosync.RegisterDeviceInput) (*service.Request, *cognitosync.RegisterDeviceOutput)

	RegisterDevice(*cognitosync.RegisterDeviceInput) (*cognitosync.RegisterDeviceOutput, error)

	SetCognitoEventsRequest(*cognitosync.SetCognitoEventsInput) (*service.Request, *cognitosync.SetCognitoEventsOutput)

	SetCognitoEvents(*cognitosync.SetCognitoEventsInput) (*cognitosync.SetCognitoEventsOutput, error)

	SetIdentityPoolConfigurationRequest(*cognitosync.SetIdentityPoolConfigurationInput) (*service.Request, *cognitosync.SetIdentityPoolConfigurationOutput)

	SetIdentityPoolConfiguration(*cognitosync.SetIdentityPoolConfigurationInput) (*cognitosync.SetIdentityPoolConfigurationOutput, error)

	SubscribeToDatasetRequest(*cognitosync.SubscribeToDatasetInput) (*service.Request, *cognitosync.SubscribeToDatasetOutput)

	SubscribeToDataset(*cognitosync.SubscribeToDatasetInput) (*cognitosync.SubscribeToDatasetOutput, error)

	UnsubscribeFromDatasetRequest(*cognitosync.UnsubscribeFromDatasetInput) (*service.Request, *cognitosync.UnsubscribeFromDatasetOutput)

	UnsubscribeFromDataset(*cognitosync.UnsubscribeFromDatasetInput) (*cognitosync.UnsubscribeFromDatasetOutput, error)

	UpdateRecordsRequest(*cognitosync.UpdateRecordsInput) (*service.Request, *cognitosync.UpdateRecordsOutput)

	UpdateRecords(*cognitosync.UpdateRecordsInput) (*cognitosync.UpdateRecordsOutput, error)
}
