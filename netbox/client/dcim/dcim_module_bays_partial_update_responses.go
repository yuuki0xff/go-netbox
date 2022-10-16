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

// DcimModuleBaysPartialUpdateReader is a Reader for the DcimModuleBaysPartialUpdate structure.
type DcimModuleBaysPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleBaysPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModuleBaysPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleBaysPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleBaysPartialUpdateOK creates a DcimModuleBaysPartialUpdateOK with default headers values
func NewDcimModuleBaysPartialUpdateOK() *DcimModuleBaysPartialUpdateOK {
	return &DcimModuleBaysPartialUpdateOK{}
}

/*
DcimModuleBaysPartialUpdateOK describes a response with status code 200, with default header values.

DcimModuleBaysPartialUpdateOK dcim module bays partial update o k
*/
type DcimModuleBaysPartialUpdateOK struct {
	Payload *models.ModuleBay
}

// IsSuccess returns true when this dcim module bays partial update o k response has a 2xx status code
func (o *DcimModuleBaysPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module bays partial update o k response has a 3xx status code
func (o *DcimModuleBaysPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bays partial update o k response has a 4xx status code
func (o *DcimModuleBaysPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module bays partial update o k response has a 5xx status code
func (o *DcimModuleBaysPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bays partial update o k response a status code equal to that given
func (o *DcimModuleBaysPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimModuleBaysPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/module-bays/{id}/][%d] dcimModuleBaysPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModuleBaysPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/module-bays/{id}/][%d] dcimModuleBaysPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModuleBaysPartialUpdateOK) GetPayload() *models.ModuleBay {
	return o.Payload
}

func (o *DcimModuleBaysPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleBaysPartialUpdateDefault creates a DcimModuleBaysPartialUpdateDefault with default headers values
func NewDcimModuleBaysPartialUpdateDefault(code int) *DcimModuleBaysPartialUpdateDefault {
	return &DcimModuleBaysPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimModuleBaysPartialUpdateDefault describes a response with status code -1, with default header values.

DcimModuleBaysPartialUpdateDefault dcim module bays partial update default
*/
type DcimModuleBaysPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module bays partial update default response
func (o *DcimModuleBaysPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim module bays partial update default response has a 2xx status code
func (o *DcimModuleBaysPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module bays partial update default response has a 3xx status code
func (o *DcimModuleBaysPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module bays partial update default response has a 4xx status code
func (o *DcimModuleBaysPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module bays partial update default response has a 5xx status code
func (o *DcimModuleBaysPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module bays partial update default response a status code equal to that given
func (o *DcimModuleBaysPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimModuleBaysPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/module-bays/{id}/][%d] dcim_module-bays_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBaysPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/module-bays/{id}/][%d] dcim_module-bays_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBaysPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBaysPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
