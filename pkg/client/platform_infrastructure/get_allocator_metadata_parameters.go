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

// NewGetAllocatorMetadataParams creates a new GetAllocatorMetadataParams object
// with the default values initialized.
func NewGetAllocatorMetadataParams() *GetAllocatorMetadataParams {
	var ()
	return &GetAllocatorMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllocatorMetadataParamsWithTimeout creates a new GetAllocatorMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllocatorMetadataParamsWithTimeout(timeout time.Duration) *GetAllocatorMetadataParams {
	var ()
	return &GetAllocatorMetadataParams{

		timeout: timeout,
	}
}

// NewGetAllocatorMetadataParamsWithContext creates a new GetAllocatorMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllocatorMetadataParamsWithContext(ctx context.Context) *GetAllocatorMetadataParams {
	var ()
	return &GetAllocatorMetadataParams{

		Context: ctx,
	}
}

// NewGetAllocatorMetadataParamsWithHTTPClient creates a new GetAllocatorMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllocatorMetadataParamsWithHTTPClient(client *http.Client) *GetAllocatorMetadataParams {
	var ()
	return &GetAllocatorMetadataParams{
		HTTPClient: client,
	}
}

/*GetAllocatorMetadataParams contains all the parameters to send to the API endpoint
for the get allocator metadata operation typically these are written to a http.Request
*/
type GetAllocatorMetadataParams struct {

	/*AllocatorID
	  The allocator identifier.

	*/
	AllocatorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get allocator metadata params
func (o *GetAllocatorMetadataParams) WithTimeout(timeout time.Duration) *GetAllocatorMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get allocator metadata params
func (o *GetAllocatorMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get allocator metadata params
func (o *GetAllocatorMetadataParams) WithContext(ctx context.Context) *GetAllocatorMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get allocator metadata params
func (o *GetAllocatorMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get allocator metadata params
func (o *GetAllocatorMetadataParams) WithHTTPClient(client *http.Client) *GetAllocatorMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get allocator metadata params
func (o *GetAllocatorMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocatorID adds the allocatorID to the get allocator metadata params
func (o *GetAllocatorMetadataParams) WithAllocatorID(allocatorID string) *GetAllocatorMetadataParams {
	o.SetAllocatorID(allocatorID)
	return o
}

// SetAllocatorID adds the allocatorId to the get allocator metadata params
func (o *GetAllocatorMetadataParams) SetAllocatorID(allocatorID string) {
	o.AllocatorID = allocatorID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllocatorMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param allocator_id
	if err := r.SetPathParam("allocator_id", o.AllocatorID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
