// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package support_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/tetrafolium/aws-sdk-go/aws"
	"github.com/tetrafolium/aws-sdk-go/aws/awserr"
	"github.com/tetrafolium/aws-sdk-go/aws/awsutil"
	"github.com/tetrafolium/aws-sdk-go/service/support"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSupport_AddAttachmentsToSet() {
	svc := support.New(nil)

	params := &support.AddAttachmentsToSetInput{
		Attachments: []*support.Attachment{ // Required
			{ // Required
				Data:     []byte("PAYLOAD"),
				FileName: aws.String("FileName"),
			},
			// More values...
		},
		AttachmentSetId: aws.String("AttachmentSetId"),
	}
	resp, err := svc.AddAttachmentsToSet(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleSupport_AddCommunicationToCase() {
	svc := support.New(nil)

	params := &support.AddCommunicationToCaseInput{
		CommunicationBody: aws.String("CommunicationBody"), // Required
		AttachmentSetId:   aws.String("AttachmentSetId"),
		CaseId:            aws.String("CaseId"),
		CcEmailAddresses: []*string{
			aws.String("CcEmailAddress"), // Required
			// More values...
		},
	}
	resp, err := svc.AddCommunicationToCase(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleSupport_CreateCase() {
	svc := support.New(nil)

	params := &support.CreateCaseInput{
		CommunicationBody: aws.String("CommunicationBody"), // Required
		Subject:           aws.String("Subject"),           // Required
		AttachmentSetId:   aws.String("AttachmentSetId"),
		CategoryCode:      aws.String("CategoryCode"),
		CcEmailAddresses: []*string{
			aws.String("CcEmailAddress"), // Required
			// More values...
		},
		IssueType:    aws.String("IssueType"),
		Language:     aws.String("Language"),
		ServiceCode:  aws.String("ServiceCode"),
		SeverityCode: aws.String("SeverityCode"),
	}
	resp, err := svc.CreateCase(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleSupport_DescribeAttachment() {
	svc := support.New(nil)

	params := &support.DescribeAttachmentInput{
		AttachmentId: aws.String("AttachmentId"), // Required
	}
	resp, err := svc.DescribeAttachment(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleSupport_DescribeCases() {
	svc := support.New(nil)

	params := &support.DescribeCasesInput{
		AfterTime:  aws.String("AfterTime"),
		BeforeTime: aws.String("BeforeTime"),
		CaseIdList: []*string{
			aws.String("CaseId"), // Required
			// More values...
		},
		DisplayId:             aws.String("DisplayId"),
		IncludeCommunications: aws.Bool(true),
		IncludeResolvedCases:  aws.Bool(true),
		Language:              aws.String("Language"),
		MaxResults:            aws.Int64(1),
		NextToken:             aws.String("NextToken"),
	}
	resp, err := svc.DescribeCases(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleSupport_DescribeCommunications() {
	svc := support.New(nil)

	params := &support.DescribeCommunicationsInput{
		CaseId:     aws.String("CaseId"), // Required
		AfterTime:  aws.String("AfterTime"),
		BeforeTime: aws.String("BeforeTime"),
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.DescribeCommunications(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleSupport_DescribeServices() {
	svc := support.New(nil)

	params := &support.DescribeServicesInput{
		Language: aws.String("Language"),
		ServiceCodeList: []*string{
			aws.String("ServiceCode"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeServices(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleSupport_DescribeSeverityLevels() {
	svc := support.New(nil)

	params := &support.DescribeSeverityLevelsInput{
		Language: aws.String("Language"),
	}
	resp, err := svc.DescribeSeverityLevels(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleSupport_DescribeTrustedAdvisorCheckRefreshStatuses() {
	svc := support.New(nil)

	params := &support.DescribeTrustedAdvisorCheckRefreshStatusesInput{
		CheckIds: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTrustedAdvisorCheckRefreshStatuses(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleSupport_DescribeTrustedAdvisorCheckResult() {
	svc := support.New(nil)

	params := &support.DescribeTrustedAdvisorCheckResultInput{
		CheckId:  aws.String("String"), // Required
		Language: aws.String("String"),
	}
	resp, err := svc.DescribeTrustedAdvisorCheckResult(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleSupport_DescribeTrustedAdvisorCheckSummaries() {
	svc := support.New(nil)

	params := &support.DescribeTrustedAdvisorCheckSummariesInput{
		CheckIds: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTrustedAdvisorCheckSummaries(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleSupport_DescribeTrustedAdvisorChecks() {
	svc := support.New(nil)

	params := &support.DescribeTrustedAdvisorChecksInput{
		Language: aws.String("String"), // Required
	}
	resp, err := svc.DescribeTrustedAdvisorChecks(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleSupport_RefreshTrustedAdvisorCheck() {
	svc := support.New(nil)

	params := &support.RefreshTrustedAdvisorCheckInput{
		CheckId: aws.String("String"), // Required
	}
	resp, err := svc.RefreshTrustedAdvisorCheck(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleSupport_ResolveCase() {
	svc := support.New(nil)

	params := &support.ResolveCaseInput{
		CaseId: aws.String("CaseId"),
	}
	resp, err := svc.ResolveCase(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}
