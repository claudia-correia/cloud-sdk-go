// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package trusted_environments

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
)

// NewGetTrustedEnvsParams creates a new GetTrustedEnvsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTrustedEnvsParams() *GetTrustedEnvsParams {
	return &GetTrustedEnvsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTrustedEnvsParamsWithTimeout creates a new GetTrustedEnvsParams object
// with the ability to set a timeout on a request.
func NewGetTrustedEnvsParamsWithTimeout(timeout time.Duration) *GetTrustedEnvsParams {
	return &GetTrustedEnvsParams{
		timeout: timeout,
	}
}

// NewGetTrustedEnvsParamsWithContext creates a new GetTrustedEnvsParams object
// with the ability to set a context for a request.
func NewGetTrustedEnvsParamsWithContext(ctx context.Context) *GetTrustedEnvsParams {
	return &GetTrustedEnvsParams{
		Context: ctx,
	}
}

// NewGetTrustedEnvsParamsWithHTTPClient creates a new GetTrustedEnvsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTrustedEnvsParamsWithHTTPClient(client *http.Client) *GetTrustedEnvsParams {
	return &GetTrustedEnvsParams{
		HTTPClient: client,
	}
}

/* GetTrustedEnvsParams contains all the parameters to send to the API endpoint
   for the get trusted envs operation.

   Typically these are written to a http.Request.
*/
type GetTrustedEnvsParams struct {

	/* OrganizationID.

	   (Optional) Organization Id for which to retrieve the trusted environments
	*/
	OrganizationID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get trusted envs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTrustedEnvsParams) WithDefaults() *GetTrustedEnvsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get trusted envs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTrustedEnvsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get trusted envs params
func (o *GetTrustedEnvsParams) WithTimeout(timeout time.Duration) *GetTrustedEnvsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get trusted envs params
func (o *GetTrustedEnvsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get trusted envs params
func (o *GetTrustedEnvsParams) WithContext(ctx context.Context) *GetTrustedEnvsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get trusted envs params
func (o *GetTrustedEnvsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get trusted envs params
func (o *GetTrustedEnvsParams) WithHTTPClient(client *http.Client) *GetTrustedEnvsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get trusted envs params
func (o *GetTrustedEnvsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get trusted envs params
func (o *GetTrustedEnvsParams) WithOrganizationID(organizationID *string) *GetTrustedEnvsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get trusted envs params
func (o *GetTrustedEnvsParams) SetOrganizationID(organizationID *string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTrustedEnvsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OrganizationID != nil {

		// query param organization_id
		var qrOrganizationID string

		if o.OrganizationID != nil {
			qrOrganizationID = *o.OrganizationID
		}
		qOrganizationID := qrOrganizationID
		if qOrganizationID != "" {

			if err := r.SetQueryParam("organization_id", qOrganizationID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
