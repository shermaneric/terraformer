// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteStackSetInput struct {
	_ struct{} `type:"structure"`

	// The name or unique ID of the stack set that you're deleting. You can obtain
	// this value by running ListStackSets.
	//
	// StackSetName is a required field
	StackSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteStackSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteStackSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteStackSetInput"}

	if s.StackSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteStackSetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteStackSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteStackSet = "DeleteStackSet"

// DeleteStackSetRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Deletes a stack set. Before you can delete a stack set, all of its member
// stack instances must be deleted. For more information about how to do this,
// see DeleteStackInstances.
//
//    // Example sending a request using DeleteStackSetRequest.
//    req := client.DeleteStackSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DeleteStackSet
func (c *Client) DeleteStackSetRequest(input *DeleteStackSetInput) DeleteStackSetRequest {
	op := &aws.Operation{
		Name:       opDeleteStackSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteStackSetInput{}
	}

	req := c.newRequest(op, input, &DeleteStackSetOutput{})
	return DeleteStackSetRequest{Request: req, Input: input, Copy: c.DeleteStackSetRequest}
}

// DeleteStackSetRequest is the request type for the
// DeleteStackSet API operation.
type DeleteStackSetRequest struct {
	*aws.Request
	Input *DeleteStackSetInput
	Copy  func(*DeleteStackSetInput) DeleteStackSetRequest
}

// Send marshals and sends the DeleteStackSet API request.
func (r DeleteStackSetRequest) Send(ctx context.Context) (*DeleteStackSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteStackSetResponse{
		DeleteStackSetOutput: r.Request.Data.(*DeleteStackSetOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteStackSetResponse is the response type for the
// DeleteStackSet API operation.
type DeleteStackSetResponse struct {
	*DeleteStackSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteStackSet request.
func (r *DeleteStackSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
