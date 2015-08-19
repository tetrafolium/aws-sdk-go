// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package sqsiface provides an interface for the Amazon Simple Queue Service.
package sqsiface

import (
	"github.com/tetrafolium/aws-sdk-go/aws/service"
	"github.com/tetrafolium/aws-sdk-go/service/sqs"
)

// SQSAPI is the interface type for sqs.SQS.
type SQSAPI interface {
	AddPermissionRequest(*sqs.AddPermissionInput) (*service.Request, *sqs.AddPermissionOutput)

	AddPermission(*sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error)

	ChangeMessageVisibilityRequest(*sqs.ChangeMessageVisibilityInput) (*service.Request, *sqs.ChangeMessageVisibilityOutput)

	ChangeMessageVisibility(*sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, error)

	ChangeMessageVisibilityBatchRequest(*sqs.ChangeMessageVisibilityBatchInput) (*service.Request, *sqs.ChangeMessageVisibilityBatchOutput)

	ChangeMessageVisibilityBatch(*sqs.ChangeMessageVisibilityBatchInput) (*sqs.ChangeMessageVisibilityBatchOutput, error)

	CreateQueueRequest(*sqs.CreateQueueInput) (*service.Request, *sqs.CreateQueueOutput)

	CreateQueue(*sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error)

	DeleteMessageRequest(*sqs.DeleteMessageInput) (*service.Request, *sqs.DeleteMessageOutput)

	DeleteMessage(*sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error)

	DeleteMessageBatchRequest(*sqs.DeleteMessageBatchInput) (*service.Request, *sqs.DeleteMessageBatchOutput)

	DeleteMessageBatch(*sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, error)

	DeleteQueueRequest(*sqs.DeleteQueueInput) (*service.Request, *sqs.DeleteQueueOutput)

	DeleteQueue(*sqs.DeleteQueueInput) (*sqs.DeleteQueueOutput, error)

	GetQueueAttributesRequest(*sqs.GetQueueAttributesInput) (*service.Request, *sqs.GetQueueAttributesOutput)

	GetQueueAttributes(*sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error)

	GetQueueUrlRequest(*sqs.GetQueueUrlInput) (*service.Request, *sqs.GetQueueUrlOutput)

	GetQueueUrl(*sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error)

	ListDeadLetterSourceQueuesRequest(*sqs.ListDeadLetterSourceQueuesInput) (*service.Request, *sqs.ListDeadLetterSourceQueuesOutput)

	ListDeadLetterSourceQueues(*sqs.ListDeadLetterSourceQueuesInput) (*sqs.ListDeadLetterSourceQueuesOutput, error)

	ListQueuesRequest(*sqs.ListQueuesInput) (*service.Request, *sqs.ListQueuesOutput)

	ListQueues(*sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error)

	PurgeQueueRequest(*sqs.PurgeQueueInput) (*service.Request, *sqs.PurgeQueueOutput)

	PurgeQueue(*sqs.PurgeQueueInput) (*sqs.PurgeQueueOutput, error)

	ReceiveMessageRequest(*sqs.ReceiveMessageInput) (*service.Request, *sqs.ReceiveMessageOutput)

	ReceiveMessage(*sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error)

	RemovePermissionRequest(*sqs.RemovePermissionInput) (*service.Request, *sqs.RemovePermissionOutput)

	RemovePermission(*sqs.RemovePermissionInput) (*sqs.RemovePermissionOutput, error)

	SendMessageRequest(*sqs.SendMessageInput) (*service.Request, *sqs.SendMessageOutput)

	SendMessage(*sqs.SendMessageInput) (*sqs.SendMessageOutput, error)

	SendMessageBatchRequest(*sqs.SendMessageBatchInput) (*service.Request, *sqs.SendMessageBatchOutput)

	SendMessageBatch(*sqs.SendMessageBatchInput) (*sqs.SendMessageBatchOutput, error)

	SetQueueAttributesRequest(*sqs.SetQueueAttributesInput) (*service.Request, *sqs.SetQueueAttributesOutput)

	SetQueueAttributes(*sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, error)
}
