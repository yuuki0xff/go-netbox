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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// TenancyContactsUpdateReader is a Reader for the TenancyContactsUpdate structure.
type TenancyContactsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactsUpdateOK creates a TenancyContactsUpdateOK with default headers values
func NewTenancyContactsUpdateOK() *TenancyContactsUpdateOK {
	return &TenancyContactsUpdateOK{}
}

/*
TenancyContactsUpdateOK describes a response with status code 200, with default header values.

TenancyContactsUpdateOK tenancy contacts update o k
*/
type TenancyContactsUpdateOK struct {
	Payload *models.Contact
}

// IsSuccess returns true when this tenancy contacts update o k response has a 2xx status code
func (o *TenancyContactsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contacts update o k response has a 3xx status code
func (o *TenancyContactsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contacts update o k response has a 4xx status code
func (o *TenancyContactsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contacts update o k response has a 5xx status code
func (o *TenancyContactsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contacts update o k response a status code equal to that given
func (o *TenancyContactsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *TenancyContactsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /tenancy/contacts/{id}/][%d] tenancyContactsUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /tenancy/contacts/{id}/][%d] tenancyContactsUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactsUpdateOK) GetPayload() *models.Contact {
	return o.Payload
}

func (o *TenancyContactsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Contact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactsUpdateDefault creates a TenancyContactsUpdateDefault with default headers values
func NewTenancyContactsUpdateDefault(code int) *TenancyContactsUpdateDefault {
	return &TenancyContactsUpdateDefault{
		_statusCode: code,
	}
}

/*
TenancyContactsUpdateDefault describes a response with status code -1, with default header values.

TenancyContactsUpdateDefault tenancy contacts update default
*/
type TenancyContactsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy contacts update default response
func (o *TenancyContactsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this tenancy contacts update default response has a 2xx status code
func (o *TenancyContactsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contacts update default response has a 3xx status code
func (o *TenancyContactsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contacts update default response has a 4xx status code
func (o *TenancyContactsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contacts update default response has a 5xx status code
func (o *TenancyContactsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contacts update default response a status code equal to that given
func (o *TenancyContactsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *TenancyContactsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /tenancy/contacts/{id}/][%d] tenancy_contacts_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /tenancy/contacts/{id}/][%d] tenancy_contacts_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
