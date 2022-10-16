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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// UsersGroupsBulkPartialUpdateReader is a Reader for the UsersGroupsBulkPartialUpdate structure.
type UsersGroupsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersGroupsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersGroupsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersGroupsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersGroupsBulkPartialUpdateOK creates a UsersGroupsBulkPartialUpdateOK with default headers values
func NewUsersGroupsBulkPartialUpdateOK() *UsersGroupsBulkPartialUpdateOK {
	return &UsersGroupsBulkPartialUpdateOK{}
}

/*
UsersGroupsBulkPartialUpdateOK describes a response with status code 200, with default header values.

UsersGroupsBulkPartialUpdateOK users groups bulk partial update o k
*/
type UsersGroupsBulkPartialUpdateOK struct {
	Payload *models.Group
}

// IsSuccess returns true when this users groups bulk partial update o k response has a 2xx status code
func (o *UsersGroupsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users groups bulk partial update o k response has a 3xx status code
func (o *UsersGroupsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users groups bulk partial update o k response has a 4xx status code
func (o *UsersGroupsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users groups bulk partial update o k response has a 5xx status code
func (o *UsersGroupsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users groups bulk partial update o k response a status code equal to that given
func (o *UsersGroupsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersGroupsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /users/groups/][%d] usersGroupsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersGroupsBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /users/groups/][%d] usersGroupsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersGroupsBulkPartialUpdateOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *UsersGroupsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersGroupsBulkPartialUpdateDefault creates a UsersGroupsBulkPartialUpdateDefault with default headers values
func NewUsersGroupsBulkPartialUpdateDefault(code int) *UsersGroupsBulkPartialUpdateDefault {
	return &UsersGroupsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
UsersGroupsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

UsersGroupsBulkPartialUpdateDefault users groups bulk partial update default
*/
type UsersGroupsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users groups bulk partial update default response
func (o *UsersGroupsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this users groups bulk partial update default response has a 2xx status code
func (o *UsersGroupsBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users groups bulk partial update default response has a 3xx status code
func (o *UsersGroupsBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users groups bulk partial update default response has a 4xx status code
func (o *UsersGroupsBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users groups bulk partial update default response has a 5xx status code
func (o *UsersGroupsBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users groups bulk partial update default response a status code equal to that given
func (o *UsersGroupsBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsersGroupsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /users/groups/][%d] users_groups_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersGroupsBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /users/groups/][%d] users_groups_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersGroupsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersGroupsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
