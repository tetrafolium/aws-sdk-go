// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package glacieriface provides an interface for the Amazon Glacier.
package glacieriface

import (
	"github.com/tetrafolium/aws-sdk-go/aws/service"
	"github.com/tetrafolium/aws-sdk-go/service/glacier"
)

// GlacierAPI is the interface type for glacier.Glacier.
type GlacierAPI interface {
	AbortMultipartUploadRequest(*glacier.AbortMultipartUploadInput) (*service.Request, *glacier.AbortMultipartUploadOutput)

	AbortMultipartUpload(*glacier.AbortMultipartUploadInput) (*glacier.AbortMultipartUploadOutput, error)

	AbortVaultLockRequest(*glacier.AbortVaultLockInput) (*service.Request, *glacier.AbortVaultLockOutput)

	AbortVaultLock(*glacier.AbortVaultLockInput) (*glacier.AbortVaultLockOutput, error)

	AddTagsToVaultRequest(*glacier.AddTagsToVaultInput) (*service.Request, *glacier.AddTagsToVaultOutput)

	AddTagsToVault(*glacier.AddTagsToVaultInput) (*glacier.AddTagsToVaultOutput, error)

	CompleteMultipartUploadRequest(*glacier.CompleteMultipartUploadInput) (*service.Request, *glacier.ArchiveCreationOutput)

	CompleteMultipartUpload(*glacier.CompleteMultipartUploadInput) (*glacier.ArchiveCreationOutput, error)

	CompleteVaultLockRequest(*glacier.CompleteVaultLockInput) (*service.Request, *glacier.CompleteVaultLockOutput)

	CompleteVaultLock(*glacier.CompleteVaultLockInput) (*glacier.CompleteVaultLockOutput, error)

	CreateVaultRequest(*glacier.CreateVaultInput) (*service.Request, *glacier.CreateVaultOutput)

	CreateVault(*glacier.CreateVaultInput) (*glacier.CreateVaultOutput, error)

	DeleteArchiveRequest(*glacier.DeleteArchiveInput) (*service.Request, *glacier.DeleteArchiveOutput)

	DeleteArchive(*glacier.DeleteArchiveInput) (*glacier.DeleteArchiveOutput, error)

	DeleteVaultRequest(*glacier.DeleteVaultInput) (*service.Request, *glacier.DeleteVaultOutput)

	DeleteVault(*glacier.DeleteVaultInput) (*glacier.DeleteVaultOutput, error)

	DeleteVaultAccessPolicyRequest(*glacier.DeleteVaultAccessPolicyInput) (*service.Request, *glacier.DeleteVaultAccessPolicyOutput)

	DeleteVaultAccessPolicy(*glacier.DeleteVaultAccessPolicyInput) (*glacier.DeleteVaultAccessPolicyOutput, error)

	DeleteVaultNotificationsRequest(*glacier.DeleteVaultNotificationsInput) (*service.Request, *glacier.DeleteVaultNotificationsOutput)

	DeleteVaultNotifications(*glacier.DeleteVaultNotificationsInput) (*glacier.DeleteVaultNotificationsOutput, error)

	DescribeJobRequest(*glacier.DescribeJobInput) (*service.Request, *glacier.JobDescription)

	DescribeJob(*glacier.DescribeJobInput) (*glacier.JobDescription, error)

	DescribeVaultRequest(*glacier.DescribeVaultInput) (*service.Request, *glacier.DescribeVaultOutput)

	DescribeVault(*glacier.DescribeVaultInput) (*glacier.DescribeVaultOutput, error)

	GetDataRetrievalPolicyRequest(*glacier.GetDataRetrievalPolicyInput) (*service.Request, *glacier.GetDataRetrievalPolicyOutput)

	GetDataRetrievalPolicy(*glacier.GetDataRetrievalPolicyInput) (*glacier.GetDataRetrievalPolicyOutput, error)

	GetJobOutputRequest(*glacier.GetJobOutputInput) (*service.Request, *glacier.GetJobOutputOutput)

	GetJobOutput(*glacier.GetJobOutputInput) (*glacier.GetJobOutputOutput, error)

	GetVaultAccessPolicyRequest(*glacier.GetVaultAccessPolicyInput) (*service.Request, *glacier.GetVaultAccessPolicyOutput)

	GetVaultAccessPolicy(*glacier.GetVaultAccessPolicyInput) (*glacier.GetVaultAccessPolicyOutput, error)

	GetVaultLockRequest(*glacier.GetVaultLockInput) (*service.Request, *glacier.GetVaultLockOutput)

	GetVaultLock(*glacier.GetVaultLockInput) (*glacier.GetVaultLockOutput, error)

	GetVaultNotificationsRequest(*glacier.GetVaultNotificationsInput) (*service.Request, *glacier.GetVaultNotificationsOutput)

	GetVaultNotifications(*glacier.GetVaultNotificationsInput) (*glacier.GetVaultNotificationsOutput, error)

	InitiateJobRequest(*glacier.InitiateJobInput) (*service.Request, *glacier.InitiateJobOutput)

	InitiateJob(*glacier.InitiateJobInput) (*glacier.InitiateJobOutput, error)

	InitiateMultipartUploadRequest(*glacier.InitiateMultipartUploadInput) (*service.Request, *glacier.InitiateMultipartUploadOutput)

	InitiateMultipartUpload(*glacier.InitiateMultipartUploadInput) (*glacier.InitiateMultipartUploadOutput, error)

	InitiateVaultLockRequest(*glacier.InitiateVaultLockInput) (*service.Request, *glacier.InitiateVaultLockOutput)

	InitiateVaultLock(*glacier.InitiateVaultLockInput) (*glacier.InitiateVaultLockOutput, error)

	ListJobsRequest(*glacier.ListJobsInput) (*service.Request, *glacier.ListJobsOutput)

	ListJobs(*glacier.ListJobsInput) (*glacier.ListJobsOutput, error)

	ListJobsPages(*glacier.ListJobsInput, func(*glacier.ListJobsOutput, bool) bool) error

	ListMultipartUploadsRequest(*glacier.ListMultipartUploadsInput) (*service.Request, *glacier.ListMultipartUploadsOutput)

	ListMultipartUploads(*glacier.ListMultipartUploadsInput) (*glacier.ListMultipartUploadsOutput, error)

	ListMultipartUploadsPages(*glacier.ListMultipartUploadsInput, func(*glacier.ListMultipartUploadsOutput, bool) bool) error

	ListPartsRequest(*glacier.ListPartsInput) (*service.Request, *glacier.ListPartsOutput)

	ListParts(*glacier.ListPartsInput) (*glacier.ListPartsOutput, error)

	ListPartsPages(*glacier.ListPartsInput, func(*glacier.ListPartsOutput, bool) bool) error

	ListTagsForVaultRequest(*glacier.ListTagsForVaultInput) (*service.Request, *glacier.ListTagsForVaultOutput)

	ListTagsForVault(*glacier.ListTagsForVaultInput) (*glacier.ListTagsForVaultOutput, error)

	ListVaultsRequest(*glacier.ListVaultsInput) (*service.Request, *glacier.ListVaultsOutput)

	ListVaults(*glacier.ListVaultsInput) (*glacier.ListVaultsOutput, error)

	ListVaultsPages(*glacier.ListVaultsInput, func(*glacier.ListVaultsOutput, bool) bool) error

	RemoveTagsFromVaultRequest(*glacier.RemoveTagsFromVaultInput) (*service.Request, *glacier.RemoveTagsFromVaultOutput)

	RemoveTagsFromVault(*glacier.RemoveTagsFromVaultInput) (*glacier.RemoveTagsFromVaultOutput, error)

	SetDataRetrievalPolicyRequest(*glacier.SetDataRetrievalPolicyInput) (*service.Request, *glacier.SetDataRetrievalPolicyOutput)

	SetDataRetrievalPolicy(*glacier.SetDataRetrievalPolicyInput) (*glacier.SetDataRetrievalPolicyOutput, error)

	SetVaultAccessPolicyRequest(*glacier.SetVaultAccessPolicyInput) (*service.Request, *glacier.SetVaultAccessPolicyOutput)

	SetVaultAccessPolicy(*glacier.SetVaultAccessPolicyInput) (*glacier.SetVaultAccessPolicyOutput, error)

	SetVaultNotificationsRequest(*glacier.SetVaultNotificationsInput) (*service.Request, *glacier.SetVaultNotificationsOutput)

	SetVaultNotifications(*glacier.SetVaultNotificationsInput) (*glacier.SetVaultNotificationsOutput, error)

	UploadArchiveRequest(*glacier.UploadArchiveInput) (*service.Request, *glacier.ArchiveCreationOutput)

	UploadArchive(*glacier.UploadArchiveInput) (*glacier.ArchiveCreationOutput, error)

	UploadMultipartPartRequest(*glacier.UploadMultipartPartInput) (*service.Request, *glacier.UploadMultipartPartOutput)

	UploadMultipartPart(*glacier.UploadMultipartPartInput) (*glacier.UploadMultipartPartOutput, error)
}
