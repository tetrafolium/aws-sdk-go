// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package kinesisiface provides an interface for the Amazon Kinesis.
package kinesisiface

import (
	"github.com/tetrafolium/aws-sdk-go/aws/service"
	"github.com/tetrafolium/aws-sdk-go/service/kinesis"
)

// KinesisAPI is the interface type for kinesis.Kinesis.
type KinesisAPI interface {
	AddTagsToStreamRequest(*kinesis.AddTagsToStreamInput) (*service.Request, *kinesis.AddTagsToStreamOutput)

	AddTagsToStream(*kinesis.AddTagsToStreamInput) (*kinesis.AddTagsToStreamOutput, error)

	CreateStreamRequest(*kinesis.CreateStreamInput) (*service.Request, *kinesis.CreateStreamOutput)

	CreateStream(*kinesis.CreateStreamInput) (*kinesis.CreateStreamOutput, error)

	DeleteStreamRequest(*kinesis.DeleteStreamInput) (*service.Request, *kinesis.DeleteStreamOutput)

	DeleteStream(*kinesis.DeleteStreamInput) (*kinesis.DeleteStreamOutput, error)

	DescribeStreamRequest(*kinesis.DescribeStreamInput) (*service.Request, *kinesis.DescribeStreamOutput)

	DescribeStream(*kinesis.DescribeStreamInput) (*kinesis.DescribeStreamOutput, error)

	DescribeStreamPages(*kinesis.DescribeStreamInput, func(*kinesis.DescribeStreamOutput, bool) bool) error

	GetRecordsRequest(*kinesis.GetRecordsInput) (*service.Request, *kinesis.GetRecordsOutput)

	GetRecords(*kinesis.GetRecordsInput) (*kinesis.GetRecordsOutput, error)

	GetShardIteratorRequest(*kinesis.GetShardIteratorInput) (*service.Request, *kinesis.GetShardIteratorOutput)

	GetShardIterator(*kinesis.GetShardIteratorInput) (*kinesis.GetShardIteratorOutput, error)

	ListStreamsRequest(*kinesis.ListStreamsInput) (*service.Request, *kinesis.ListStreamsOutput)

	ListStreams(*kinesis.ListStreamsInput) (*kinesis.ListStreamsOutput, error)

	ListStreamsPages(*kinesis.ListStreamsInput, func(*kinesis.ListStreamsOutput, bool) bool) error

	ListTagsForStreamRequest(*kinesis.ListTagsForStreamInput) (*service.Request, *kinesis.ListTagsForStreamOutput)

	ListTagsForStream(*kinesis.ListTagsForStreamInput) (*kinesis.ListTagsForStreamOutput, error)

	MergeShardsRequest(*kinesis.MergeShardsInput) (*service.Request, *kinesis.MergeShardsOutput)

	MergeShards(*kinesis.MergeShardsInput) (*kinesis.MergeShardsOutput, error)

	PutRecordRequest(*kinesis.PutRecordInput) (*service.Request, *kinesis.PutRecordOutput)

	PutRecord(*kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error)

	PutRecordsRequest(*kinesis.PutRecordsInput) (*service.Request, *kinesis.PutRecordsOutput)

	PutRecords(*kinesis.PutRecordsInput) (*kinesis.PutRecordsOutput, error)

	RemoveTagsFromStreamRequest(*kinesis.RemoveTagsFromStreamInput) (*service.Request, *kinesis.RemoveTagsFromStreamOutput)

	RemoveTagsFromStream(*kinesis.RemoveTagsFromStreamInput) (*kinesis.RemoveTagsFromStreamOutput, error)

	SplitShardRequest(*kinesis.SplitShardInput) (*service.Request, *kinesis.SplitShardOutput)

	SplitShard(*kinesis.SplitShardInput) (*kinesis.SplitShardOutput, error)
}
