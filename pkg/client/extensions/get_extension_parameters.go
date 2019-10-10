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

package extensions

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

// NewGetExtensionParams creates a new GetExtensionParams object
// with the default values initialized.
func NewGetExtensionParams() *GetExtensionParams {
	var ()
	return &GetExtensionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetExtensionParamsWithTimeout creates a new GetExtensionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetExtensionParamsWithTimeout(timeout time.Duration) *GetExtensionParams {
	var ()
	return &GetExtensionParams{

		timeout: timeout,
	}
}

// NewGetExtensionParamsWithContext creates a new GetExtensionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetExtensionParamsWithContext(ctx context.Context) *GetExtensionParams {
	var ()
	return &GetExtensionParams{

		Context: ctx,
	}
}

// NewGetExtensionParamsWithHTTPClient creates a new GetExtensionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetExtensionParamsWithHTTPClient(client *http.Client) *GetExtensionParams {
	var ()
	return &GetExtensionParams{
		HTTPClient: client,
	}
}

/*GetExtensionParams contains all the parameters to send to the API endpoint
for the get extension operation typically these are written to a http.Request
*/
type GetExtensionParams struct {

	/*ExtensionID
	  Id of an Extension

	*/
	ExtensionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get extension params
func (o *GetExtensionParams) WithTimeout(timeout time.Duration) *GetExtensionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get extension params
func (o *GetExtensionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get extension params
func (o *GetExtensionParams) WithContext(ctx context.Context) *GetExtensionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get extension params
func (o *GetExtensionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get extension params
func (o *GetExtensionParams) WithHTTPClient(client *http.Client) *GetExtensionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get extension params
func (o *GetExtensionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtensionID adds the extensionID to the get extension params
func (o *GetExtensionParams) WithExtensionID(extensionID string) *GetExtensionParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the get extension params
func (o *GetExtensionParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetExtensionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param extension_id
	if err := r.SetPathParam("extension_id", o.ExtensionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
