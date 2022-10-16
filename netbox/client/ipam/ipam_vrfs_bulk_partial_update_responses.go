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

// IpamVrfsBulkPartialUpdateReader is a Reader for the IpamVrfsBulkPartialUpdate structure.
type IpamVrfsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVrfsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVrfsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVrfsBulkPartialUpdateOK creates a IpamVrfsBulkPartialUpdateOK with default headers values
func NewIpamVrfsBulkPartialUpdateOK() *IpamVrfsBulkPartialUpdateOK {
	return &IpamVrfsBulkPartialUpdateOK{}
}

/*
IpamVrfsBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamVrfsBulkPartialUpdateOK ipam vrfs bulk partial update o k
*/
type IpamVrfsBulkPartialUpdateOK struct {
	Payload *models.VRF
}

// IsSuccess returns true when this ipam vrfs bulk partial update o k response has a 2xx status code
func (o *IpamVrfsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vrfs bulk partial update o k response has a 3xx status code
func (o *IpamVrfsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vrfs bulk partial update o k response has a 4xx status code
func (o *IpamVrfsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vrfs bulk partial update o k response has a 5xx status code
func (o *IpamVrfsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vrfs bulk partial update o k response a status code equal to that given
func (o *IpamVrfsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamVrfsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vrfs/][%d] ipamVrfsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamVrfsBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/vrfs/][%d] ipamVrfsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamVrfsBulkPartialUpdateOK) GetPayload() *models.VRF {
	return o.Payload
}

func (o *IpamVrfsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVrfsBulkPartialUpdateDefault creates a IpamVrfsBulkPartialUpdateDefault with default headers values
func NewIpamVrfsBulkPartialUpdateDefault(code int) *IpamVrfsBulkPartialUpdateDefault {
	return &IpamVrfsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamVrfsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

IpamVrfsBulkPartialUpdateDefault ipam vrfs bulk partial update default
*/
type IpamVrfsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vrfs bulk partial update default response
func (o *IpamVrfsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam vrfs bulk partial update default response has a 2xx status code
func (o *IpamVrfsBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam vrfs bulk partial update default response has a 3xx status code
func (o *IpamVrfsBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam vrfs bulk partial update default response has a 4xx status code
func (o *IpamVrfsBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam vrfs bulk partial update default response has a 5xx status code
func (o *IpamVrfsBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam vrfs bulk partial update default response a status code equal to that given
func (o *IpamVrfsBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamVrfsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vrfs/][%d] ipam_vrfs_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVrfsBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/vrfs/][%d] ipam_vrfs_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVrfsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVrfsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
