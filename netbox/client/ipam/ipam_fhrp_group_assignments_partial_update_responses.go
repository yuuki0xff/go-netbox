// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// IpamFhrpGroupAssignmentsPartialUpdateReader is a Reader for the IpamFhrpGroupAssignmentsPartialUpdate structure.
type IpamFhrpGroupAssignmentsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupAssignmentsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamFhrpGroupAssignmentsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamFhrpGroupAssignmentsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamFhrpGroupAssignmentsPartialUpdateOK creates a IpamFhrpGroupAssignmentsPartialUpdateOK with default headers values
func NewIpamFhrpGroupAssignmentsPartialUpdateOK() *IpamFhrpGroupAssignmentsPartialUpdateOK {
	return &IpamFhrpGroupAssignmentsPartialUpdateOK{}
}

/*
IpamFhrpGroupAssignmentsPartialUpdateOK describes a response with status code 200, with default header values.

IpamFhrpGroupAssignmentsPartialUpdateOK ipam fhrp group assignments partial update o k
*/
type IpamFhrpGroupAssignmentsPartialUpdateOK struct {
	Payload *models.FHRPGroupAssignment
}

// IsSuccess returns true when this ipam fhrp group assignments partial update o k response has a 2xx status code
func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp group assignments partial update o k response has a 3xx status code
func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp group assignments partial update o k response has a 4xx status code
func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp group assignments partial update o k response has a 5xx status code
func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp group assignments partial update o k response a status code equal to that given
func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/fhrp-group-assignments/{id}/][%d] ipamFhrpGroupAssignmentsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/fhrp-group-assignments/{id}/][%d] ipamFhrpGroupAssignmentsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) GetPayload() *models.FHRPGroupAssignment {
	return o.Payload
}

func (o *IpamFhrpGroupAssignmentsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FHRPGroupAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamFhrpGroupAssignmentsPartialUpdateDefault creates a IpamFhrpGroupAssignmentsPartialUpdateDefault with default headers values
func NewIpamFhrpGroupAssignmentsPartialUpdateDefault(code int) *IpamFhrpGroupAssignmentsPartialUpdateDefault {
	return &IpamFhrpGroupAssignmentsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamFhrpGroupAssignmentsPartialUpdateDefault describes a response with status code -1, with default header values.

IpamFhrpGroupAssignmentsPartialUpdateDefault ipam fhrp group assignments partial update default
*/
type IpamFhrpGroupAssignmentsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam fhrp group assignments partial update default response
func (o *IpamFhrpGroupAssignmentsPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam fhrp group assignments partial update default response has a 2xx status code
func (o *IpamFhrpGroupAssignmentsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam fhrp group assignments partial update default response has a 3xx status code
func (o *IpamFhrpGroupAssignmentsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam fhrp group assignments partial update default response has a 4xx status code
func (o *IpamFhrpGroupAssignmentsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam fhrp group assignments partial update default response has a 5xx status code
func (o *IpamFhrpGroupAssignmentsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam fhrp group assignments partial update default response a status code equal to that given
func (o *IpamFhrpGroupAssignmentsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamFhrpGroupAssignmentsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/fhrp-group-assignments/{id}/][%d] ipam_fhrp-group-assignments_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/fhrp-group-assignments/{id}/][%d] ipam_fhrp-group-assignments_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamFhrpGroupAssignmentsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
