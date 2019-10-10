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

// NewDeleteInstanceConfigurationParams creates a new DeleteInstanceConfigurationParams object
// with the default values initialized.
func NewDeleteInstanceConfigurationParams() *DeleteInstanceConfigurationParams {
	var ()
	return &DeleteInstanceConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstanceConfigurationParamsWithTimeout creates a new DeleteInstanceConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInstanceConfigurationParamsWithTimeout(timeout time.Duration) *DeleteInstanceConfigurationParams {
	var ()
	return &DeleteInstanceConfigurationParams{

		timeout: timeout,
	}
}

// NewDeleteInstanceConfigurationParamsWithContext creates a new DeleteInstanceConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInstanceConfigurationParamsWithContext(ctx context.Context) *DeleteInstanceConfigurationParams {
	var ()
	return &DeleteInstanceConfigurationParams{

		Context: ctx,
	}
}

// NewDeleteInstanceConfigurationParamsWithHTTPClient creates a new DeleteInstanceConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInstanceConfigurationParamsWithHTTPClient(client *http.Client) *DeleteInstanceConfigurationParams {
	var ()
	return &DeleteInstanceConfigurationParams{
		HTTPClient: client,
	}
}

/*DeleteInstanceConfigurationParams contains all the parameters to send to the API endpoint
for the delete instance configuration operation typically these are written to a http.Request
*/
type DeleteInstanceConfigurationParams struct {

	/*ID
	  ID of the instance configuration

	*/
	ID string
	/*Version
	  If specified, checks for conflicts against the version of the repository configuration (returned in 'x-cloud-resource-version' of the GET request)

	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) WithTimeout(timeout time.Duration) *DeleteInstanceConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) WithContext(ctx context.Context) *DeleteInstanceConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) WithHTTPClient(client *http.Client) *DeleteInstanceConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) WithID(id string) *DeleteInstanceConfigurationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) WithVersion(version *int64) *DeleteInstanceConfigurationParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstanceConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion int64
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
