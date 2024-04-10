// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// GetLrpReader is a Reader for the GetLrp structure.
type GetLrpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLrpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLrpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /lrp] GetLrp", response, response.Code())
	}
}

// NewGetLrpOK creates a GetLrpOK with default headers values
func NewGetLrpOK() *GetLrpOK {
	return &GetLrpOK{}
}

/*
GetLrpOK describes a response with status code 200, with default header values.

Success
*/
type GetLrpOK struct {
	Payload []*models.LRPSpec
}

// IsSuccess returns true when this get lrp o k response has a 2xx status code
func (o *GetLrpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get lrp o k response has a 3xx status code
func (o *GetLrpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get lrp o k response has a 4xx status code
func (o *GetLrpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get lrp o k response has a 5xx status code
func (o *GetLrpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get lrp o k response a status code equal to that given
func (o *GetLrpOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get lrp o k response
func (o *GetLrpOK) Code() int {
	return 200
}

func (o *GetLrpOK) Error() string {
	return fmt.Sprintf("[GET /lrp][%d] getLrpOK  %+v", 200, o.Payload)
}

func (o *GetLrpOK) String() string {
	return fmt.Sprintf("[GET /lrp][%d] getLrpOK  %+v", 200, o.Payload)
}

func (o *GetLrpOK) GetPayload() []*models.LRPSpec {
	return o.Payload
}

func (o *GetLrpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}