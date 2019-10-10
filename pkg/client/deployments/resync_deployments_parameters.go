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

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewResyncDeploymentsParams creates a new ResyncDeploymentsParams object
// with the default values initialized.
func NewResyncDeploymentsParams() *ResyncDeploymentsParams {

	return &ResyncDeploymentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResyncDeploymentsParamsWithTimeout creates a new ResyncDeploymentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResyncDeploymentsParamsWithTimeout(timeout time.Duration) *ResyncDeploymentsParams {

	return &ResyncDeploymentsParams{

		timeout: timeout,
	}
}

// NewResyncDeploymentsParamsWithContext creates a new ResyncDeploymentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewResyncDeploymentsParamsWithContext(ctx context.Context) *ResyncDeploymentsParams {

	return &ResyncDeploymentsParams{

		Context: ctx,
	}
}

// NewResyncDeploymentsParamsWithHTTPClient creates a new ResyncDeploymentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResyncDeploymentsParamsWithHTTPClient(client *http.Client) *ResyncDeploymentsParams {

	return &ResyncDeploymentsParams{
		HTTPClient: client,
	}
}

/*ResyncDeploymentsParams contains all the parameters to send to the API endpoint
for the resync deployments operation typically these are written to a http.Request
*/
type ResyncDeploymentsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the resync deployments params
func (o *ResyncDeploymentsParams) WithTimeout(timeout time.Duration) *ResyncDeploymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resync deployments params
func (o *ResyncDeploymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resync deployments params
func (o *ResyncDeploymentsParams) WithContext(ctx context.Context) *ResyncDeploymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resync deployments params
func (o *ResyncDeploymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resync deployments params
func (o *ResyncDeploymentsParams) WithHTTPClient(client *http.Client) *ResyncDeploymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resync deployments params
func (o *ResyncDeploymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ResyncDeploymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
