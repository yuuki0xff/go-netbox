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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimInterfacesReadReader is a Reader for the DcimInterfacesRead structure.
type DcimInterfacesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfacesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInterfacesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInterfacesReadOK creates a DcimInterfacesReadOK with default headers values
func NewDcimInterfacesReadOK() *DcimInterfacesReadOK {
	return &DcimInterfacesReadOK{}
}

/*
DcimInterfacesReadOK describes a response with status code 200, with default header values.

DcimInterfacesReadOK dcim interfaces read o k
*/
type DcimInterfacesReadOK struct {
	Payload *models.Interface
}

// IsSuccess returns true when this dcim interfaces read o k response has a 2xx status code
func (o *DcimInterfacesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim interfaces read o k response has a 3xx status code
func (o *DcimInterfacesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim interfaces read o k response has a 4xx status code
func (o *DcimInterfacesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim interfaces read o k response has a 5xx status code
func (o *DcimInterfacesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim interfaces read o k response a status code equal to that given
func (o *DcimInterfacesReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimInterfacesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/interfaces/{id}/][%d] dcimInterfacesReadOK  %+v", 200, o.Payload)
}

func (o *DcimInterfacesReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/interfaces/{id}/][%d] dcimInterfacesReadOK  %+v", 200, o.Payload)
}

func (o *DcimInterfacesReadOK) GetPayload() *models.Interface {
	return o.Payload
}

func (o *DcimInterfacesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Interface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInterfacesReadDefault creates a DcimInterfacesReadDefault with default headers values
func NewDcimInterfacesReadDefault(code int) *DcimInterfacesReadDefault {
	return &DcimInterfacesReadDefault{
		_statusCode: code,
	}
}

/*
DcimInterfacesReadDefault describes a response with status code -1, with default header values.

DcimInterfacesReadDefault dcim interfaces read default
*/
type DcimInterfacesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim interfaces read default response
func (o *DcimInterfacesReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim interfaces read default response has a 2xx status code
func (o *DcimInterfacesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim interfaces read default response has a 3xx status code
func (o *DcimInterfacesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim interfaces read default response has a 4xx status code
func (o *DcimInterfacesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim interfaces read default response has a 5xx status code
func (o *DcimInterfacesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim interfaces read default response a status code equal to that given
func (o *DcimInterfacesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimInterfacesReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/interfaces/{id}/][%d] dcim_interfaces_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfacesReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/interfaces/{id}/][%d] dcim_interfaces_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfacesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInterfacesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
