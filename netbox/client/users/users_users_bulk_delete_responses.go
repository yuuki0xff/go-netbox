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
)

// UsersUsersBulkDeleteReader is a Reader for the UsersUsersBulkDelete structure.
type UsersUsersBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUsersUsersBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersUsersBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersUsersBulkDeleteNoContent creates a UsersUsersBulkDeleteNoContent with default headers values
func NewUsersUsersBulkDeleteNoContent() *UsersUsersBulkDeleteNoContent {
	return &UsersUsersBulkDeleteNoContent{}
}

/*
UsersUsersBulkDeleteNoContent describes a response with status code 204, with default header values.

UsersUsersBulkDeleteNoContent users users bulk delete no content
*/
type UsersUsersBulkDeleteNoContent struct {
}

// IsSuccess returns true when this users users bulk delete no content response has a 2xx status code
func (o *UsersUsersBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users users bulk delete no content response has a 3xx status code
func (o *UsersUsersBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users users bulk delete no content response has a 4xx status code
func (o *UsersUsersBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this users users bulk delete no content response has a 5xx status code
func (o *UsersUsersBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this users users bulk delete no content response a status code equal to that given
func (o *UsersUsersBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UsersUsersBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/users/][%d] usersUsersBulkDeleteNoContent ", 204)
}

func (o *UsersUsersBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /users/users/][%d] usersUsersBulkDeleteNoContent ", 204)
}

func (o *UsersUsersBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUsersUsersBulkDeleteDefault creates a UsersUsersBulkDeleteDefault with default headers values
func NewUsersUsersBulkDeleteDefault(code int) *UsersUsersBulkDeleteDefault {
	return &UsersUsersBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
UsersUsersBulkDeleteDefault describes a response with status code -1, with default header values.

UsersUsersBulkDeleteDefault users users bulk delete default
*/
type UsersUsersBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users users bulk delete default response
func (o *UsersUsersBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this users users bulk delete default response has a 2xx status code
func (o *UsersUsersBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users users bulk delete default response has a 3xx status code
func (o *UsersUsersBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users users bulk delete default response has a 4xx status code
func (o *UsersUsersBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users users bulk delete default response has a 5xx status code
func (o *UsersUsersBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users users bulk delete default response a status code equal to that given
func (o *UsersUsersBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsersUsersBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /users/users/][%d] users_users_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *UsersUsersBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /users/users/][%d] users_users_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *UsersUsersBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersUsersBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
