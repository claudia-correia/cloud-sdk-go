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

package platform_configuration_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetInstanceConfigurationParams creates a new GetInstanceConfigurationParams object
// with the default values initialized.
func NewGetInstanceConfigurationParams() *GetInstanceConfigurationParams {
	var ()
	return &GetInstanceConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstanceConfigurationParamsWithTimeout creates a new GetInstanceConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstanceConfigurationParamsWithTimeout(timeout time.Duration) *GetInstanceConfigurationParams {
	var ()
	return &GetInstanceConfigurationParams{

		timeout: timeout,
	}
}

// NewGetInstanceConfigurationParamsWithContext creates a new GetInstanceConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstanceConfigurationParamsWithContext(ctx context.Context) *GetInstanceConfigurationParams {
	var ()
	return &GetInstanceConfigurationParams{

		Context: ctx,
	}
}

// NewGetInstanceConfigurationParamsWithHTTPClient creates a new GetInstanceConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstanceConfigurationParamsWithHTTPClient(client *http.Client) *GetInstanceConfigurationParams {
	var ()
	return &GetInstanceConfigurationParams{
		HTTPClient: client,
	}
}

/*GetInstanceConfigurationParams contains all the parameters to send to the API endpoint
for the get instance configuration operation typically these are written to a http.Request
*/
type GetInstanceConfigurationParams struct {

	/*ID
	  ID of the instance configuration

	*/
	ID string
	/*ShowDeleted
	  If true, if the instance configuration has been marked for deletion it is still returned. Otherwise, instance configurations marked for deletion generate a 404

	*/
	ShowDeleted *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instance configuration params
func (o *GetInstanceConfigurationParams) WithTimeout(timeout time.Duration) *GetInstanceConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instance configuration params
func (o *GetInstanceConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instance configuration params
func (o *GetInstanceConfigurationParams) WithContext(ctx context.Context) *GetInstanceConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instance configuration params
func (o *GetInstanceConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instance configuration params
func (o *GetInstanceConfigurationParams) WithHTTPClient(client *http.Client) *GetInstanceConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instance configuration params
func (o *GetInstanceConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get instance configuration params
func (o *GetInstanceConfigurationParams) WithID(id string) *GetInstanceConfigurationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get instance configuration params
func (o *GetInstanceConfigurationParams) SetID(id string) {
	o.ID = id
}

// WithShowDeleted adds the showDeleted to the get instance configuration params
func (o *GetInstanceConfigurationParams) WithShowDeleted(showDeleted *bool) *GetInstanceConfigurationParams {
	o.SetShowDeleted(showDeleted)
	return o
}

// SetShowDeleted adds the showDeleted to the get instance configuration params
func (o *GetInstanceConfigurationParams) SetShowDeleted(showDeleted *bool) {
	o.ShowDeleted = showDeleted
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstanceConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.ShowDeleted != nil {

		// query param show_deleted
		var qrShowDeleted bool
		if o.ShowDeleted != nil {
			qrShowDeleted = *o.ShowDeleted
		}
		qShowDeleted := swag.FormatBool(qrShowDeleted)
		if qShowDeleted != "" {
			if err := r.SetQueryParam("show_deleted", qShowDeleted); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
