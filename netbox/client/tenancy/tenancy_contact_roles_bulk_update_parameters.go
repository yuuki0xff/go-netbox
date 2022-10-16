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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// NewTenancyContactRolesBulkUpdateParams creates a new TenancyContactRolesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyContactRolesBulkUpdateParams() *TenancyContactRolesBulkUpdateParams {
	return &TenancyContactRolesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyContactRolesBulkUpdateParamsWithTimeout creates a new TenancyContactRolesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewTenancyContactRolesBulkUpdateParamsWithTimeout(timeout time.Duration) *TenancyContactRolesBulkUpdateParams {
	return &TenancyContactRolesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewTenancyContactRolesBulkUpdateParamsWithContext creates a new TenancyContactRolesBulkUpdateParams object
// with the ability to set a context for a request.
func NewTenancyContactRolesBulkUpdateParamsWithContext(ctx context.Context) *TenancyContactRolesBulkUpdateParams {
	return &TenancyContactRolesBulkUpdateParams{
		Context: ctx,
	}
}

// NewTenancyContactRolesBulkUpdateParamsWithHTTPClient creates a new TenancyContactRolesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyContactRolesBulkUpdateParamsWithHTTPClient(client *http.Client) *TenancyContactRolesBulkUpdateParams {
	return &TenancyContactRolesBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
TenancyContactRolesBulkUpdateParams contains all the parameters to send to the API endpoint

	for the tenancy contact roles bulk update operation.

	Typically these are written to a http.Request.
*/
type TenancyContactRolesBulkUpdateParams struct {

	// Data.
	Data *models.ContactRole

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy contact roles bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactRolesBulkUpdateParams) WithDefaults() *TenancyContactRolesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy contact roles bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactRolesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy contact roles bulk update params
func (o *TenancyContactRolesBulkUpdateParams) WithTimeout(timeout time.Duration) *TenancyContactRolesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy contact roles bulk update params
func (o *TenancyContactRolesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy contact roles bulk update params
func (o *TenancyContactRolesBulkUpdateParams) WithContext(ctx context.Context) *TenancyContactRolesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy contact roles bulk update params
func (o *TenancyContactRolesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy contact roles bulk update params
func (o *TenancyContactRolesBulkUpdateParams) WithHTTPClient(client *http.Client) *TenancyContactRolesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy contact roles bulk update params
func (o *TenancyContactRolesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tenancy contact roles bulk update params
func (o *TenancyContactRolesBulkUpdateParams) WithData(data *models.ContactRole) *TenancyContactRolesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tenancy contact roles bulk update params
func (o *TenancyContactRolesBulkUpdateParams) SetData(data *models.ContactRole) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyContactRolesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
