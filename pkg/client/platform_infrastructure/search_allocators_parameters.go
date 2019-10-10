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

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewSearchAllocatorsParams creates a new SearchAllocatorsParams object
// with the default values initialized.
func NewSearchAllocatorsParams() *SearchAllocatorsParams {
	var ()
	return &SearchAllocatorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchAllocatorsParamsWithTimeout creates a new SearchAllocatorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchAllocatorsParamsWithTimeout(timeout time.Duration) *SearchAllocatorsParams {
	var ()
	return &SearchAllocatorsParams{

		timeout: timeout,
	}
}

// NewSearchAllocatorsParamsWithContext creates a new SearchAllocatorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchAllocatorsParamsWithContext(ctx context.Context) *SearchAllocatorsParams {
	var ()
	return &SearchAllocatorsParams{

		Context: ctx,
	}
}

// NewSearchAllocatorsParamsWithHTTPClient creates a new SearchAllocatorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchAllocatorsParamsWithHTTPClient(client *http.Client) *SearchAllocatorsParams {
	var ()
	return &SearchAllocatorsParams{
		HTTPClient: client,
	}
}

/*SearchAllocatorsParams contains all the parameters to send to the API endpoint
for the search allocators operation typically these are written to a http.Request
*/
type SearchAllocatorsParams struct {

	/*Body
	  The optional search request to execute. If not supplied then all allocators are matched

	*/
	Body *models.SearchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search allocators params
func (o *SearchAllocatorsParams) WithTimeout(timeout time.Duration) *SearchAllocatorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search allocators params
func (o *SearchAllocatorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search allocators params
func (o *SearchAllocatorsParams) WithContext(ctx context.Context) *SearchAllocatorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search allocators params
func (o *SearchAllocatorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search allocators params
func (o *SearchAllocatorsParams) WithHTTPClient(client *http.Client) *SearchAllocatorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search allocators params
func (o *SearchAllocatorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search allocators params
func (o *SearchAllocatorsParams) WithBody(body *models.SearchRequest) *SearchAllocatorsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search allocators params
func (o *SearchAllocatorsParams) SetBody(body *models.SearchRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SearchAllocatorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
