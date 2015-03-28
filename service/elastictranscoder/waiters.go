// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package elastictranscoder

import (
	"github.com/awslabs/aws-sdk-go/internal/waiter"
)

var waiterJobComplete *waiter.Config

func (c *ElasticTranscoder) WaitUntilJobComplete(input *ReadJobInput) error {
	if waiterJobComplete == nil {
		waiterJobComplete = &waiter.Config{
			Operation:   "ReadJob",
			Delay:       30,
			MaxAttempts: 120,
			Acceptors: []waiter.WaitAcceptor{
				{
					State:    "success",
					Matcher:  "path",
					Argument: "Job.Status",
					Expected: "Complete",
				},
				{
					State:    "failure",
					Matcher:  "path",
					Argument: "Job.Status",
					Expected: "Canceled",
				},
				{
					State:    "failure",
					Matcher:  "path",
					Argument: "Job.Status",
					Expected: "Error",
				},
			},
		}
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterJobComplete,
	}
	return w.Wait()
}
