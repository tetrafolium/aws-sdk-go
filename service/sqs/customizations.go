package sqs

import "github.com/tetrafolium/aws-sdk-go/aws/service"

func init() {
	initRequest = func(r *service.Request) {
		setupChecksumValidation(r)
	}
}
