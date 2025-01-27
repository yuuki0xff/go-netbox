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
)

// DcimPowerOutletTemplatesBulkDeleteReader is a Reader for the DcimPowerOutletTemplatesBulkDelete structure.
type DcimPowerOutletTemplatesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerOutletTemplatesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerOutletTemplatesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerOutletTemplatesBulkDeleteNoContent creates a DcimPowerOutletTemplatesBulkDeleteNoContent with default headers values
func NewDcimPowerOutletTemplatesBulkDeleteNoContent() *DcimPowerOutletTemplatesBulkDeleteNoContent {
	return &DcimPowerOutletTemplatesBulkDeleteNoContent{}
}

/*
DcimPowerOutletTemplatesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerOutletTemplatesBulkDeleteNoContent dcim power outlet templates bulk delete no content
*/
type DcimPowerOutletTemplatesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim power outlet templates bulk delete no content response has a 2xx status code
func (o *DcimPowerOutletTemplatesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlet templates bulk delete no content response has a 3xx status code
func (o *DcimPowerOutletTemplatesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlet templates bulk delete no content response has a 4xx status code
func (o *DcimPowerOutletTemplatesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlet templates bulk delete no content response has a 5xx status code
func (o *DcimPowerOutletTemplatesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlet templates bulk delete no content response a status code equal to that given
func (o *DcimPowerOutletTemplatesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimPowerOutletTemplatesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-outlet-templates/][%d] dcimPowerOutletTemplatesBulkDeleteNoContent ", 204)
}

func (o *DcimPowerOutletTemplatesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/power-outlet-templates/][%d] dcimPowerOutletTemplatesBulkDeleteNoContent ", 204)
}

func (o *DcimPowerOutletTemplatesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimPowerOutletTemplatesBulkDeleteDefault creates a DcimPowerOutletTemplatesBulkDeleteDefault with default headers values
func NewDcimPowerOutletTemplatesBulkDeleteDefault(code int) *DcimPowerOutletTemplatesBulkDeleteDefault {
	return &DcimPowerOutletTemplatesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimPowerOutletTemplatesBulkDeleteDefault describes a response with status code -1, with default header values.

DcimPowerOutletTemplatesBulkDeleteDefault dcim power outlet templates bulk delete default
*/
type DcimPowerOutletTemplatesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power outlet templates bulk delete default response
func (o *DcimPowerOutletTemplatesBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power outlet templates bulk delete default response has a 2xx status code
func (o *DcimPowerOutletTemplatesBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power outlet templates bulk delete default response has a 3xx status code
func (o *DcimPowerOutletTemplatesBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power outlet templates bulk delete default response has a 4xx status code
func (o *DcimPowerOutletTemplatesBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power outlet templates bulk delete default response has a 5xx status code
func (o *DcimPowerOutletTemplatesBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power outlet templates bulk delete default response a status code equal to that given
func (o *DcimPowerOutletTemplatesBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerOutletTemplatesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-outlet-templates/][%d] dcim_power-outlet-templates_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletTemplatesBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/power-outlet-templates/][%d] dcim_power-outlet-templates_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletTemplatesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
