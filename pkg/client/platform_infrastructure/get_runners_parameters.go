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

package platform_infrastructure

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

// NewGetRunnersParams creates a new GetRunnersParams object
// with the default values initialized.
func NewGetRunnersParams() *GetRunnersParams {
	var ()
	return &GetRunnersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunnersParamsWithTimeout creates a new GetRunnersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRunnersParamsWithTimeout(timeout time.Duration) *GetRunnersParams {
	var ()
	return &GetRunnersParams{

		timeout: timeout,
	}
}

// NewGetRunnersParamsWithContext creates a new GetRunnersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRunnersParamsWithContext(ctx context.Context) *GetRunnersParams {
	var ()
	return &GetRunnersParams{

		Context: ctx,
	}
}

// NewGetRunnersParamsWithHTTPClient creates a new GetRunnersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRunnersParamsWithHTTPClient(client *http.Client) *GetRunnersParams {
	var ()
	return &GetRunnersParams{
		HTTPClient: client,
	}
}

/*GetRunnersParams contains all the parameters to send to the API endpoint
for the get runners operation typically these are written to a http.Request
*/
type GetRunnersParams struct {

	/*Q
	  An optional query to filter runners by. Maps to an Elasticsearch query_string query.

	*/
	Q *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get runners params
func (o *GetRunnersParams) WithTimeout(timeout time.Duration) *GetRunnersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runners params
func (o *GetRunnersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runners params
func (o *GetRunnersParams) WithContext(ctx context.Context) *GetRunnersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runners params
func (o *GetRunnersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runners params
func (o *GetRunnersParams) WithHTTPClient(client *http.Client) *GetRunnersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runners params
func (o *GetRunnersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQ adds the q to the get runners params
func (o *GetRunnersParams) WithQ(q *string) *GetRunnersParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the get runners params
func (o *GetRunnersParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunnersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
