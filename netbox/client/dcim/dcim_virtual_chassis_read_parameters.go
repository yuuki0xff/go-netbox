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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDcimVirtualChassisReadParams creates a new DcimVirtualChassisReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimVirtualChassisReadParams() *DcimVirtualChassisReadParams {
	return &DcimVirtualChassisReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimVirtualChassisReadParamsWithTimeout creates a new DcimVirtualChassisReadParams object
// with the ability to set a timeout on a request.
func NewDcimVirtualChassisReadParamsWithTimeout(timeout time.Duration) *DcimVirtualChassisReadParams {
	return &DcimVirtualChassisReadParams{
		timeout: timeout,
	}
}

// NewDcimVirtualChassisReadParamsWithContext creates a new DcimVirtualChassisReadParams object
// with the ability to set a context for a request.
func NewDcimVirtualChassisReadParamsWithContext(ctx context.Context) *DcimVirtualChassisReadParams {
	return &DcimVirtualChassisReadParams{
		Context: ctx,
	}
}

// NewDcimVirtualChassisReadParamsWithHTTPClient creates a new DcimVirtualChassisReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimVirtualChassisReadParamsWithHTTPClient(client *http.Client) *DcimVirtualChassisReadParams {
	return &DcimVirtualChassisReadParams{
		HTTPClient: client,
	}
}

/*
DcimVirtualChassisReadParams contains all the parameters to send to the API endpoint

	for the dcim virtual chassis read operation.

	Typically these are written to a http.Request.
*/
type DcimVirtualChassisReadParams struct {

	/* ID.

	   A unique integer value identifying this virtual chassis.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim virtual chassis read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualChassisReadParams) WithDefaults() *DcimVirtualChassisReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim virtual chassis read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualChassisReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim virtual chassis read params
func (o *DcimVirtualChassisReadParams) WithTimeout(timeout time.Duration) *DcimVirtualChassisReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim virtual chassis read params
func (o *DcimVirtualChassisReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim virtual chassis read params
func (o *DcimVirtualChassisReadParams) WithContext(ctx context.Context) *DcimVirtualChassisReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim virtual chassis read params
func (o *DcimVirtualChassisReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim virtual chassis read params
func (o *DcimVirtualChassisReadParams) WithHTTPClient(client *http.Client) *DcimVirtualChassisReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim virtual chassis read params
func (o *DcimVirtualChassisReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim virtual chassis read params
func (o *DcimVirtualChassisReadParams) WithID(id int64) *DcimVirtualChassisReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim virtual chassis read params
func (o *DcimVirtualChassisReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimVirtualChassisReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
