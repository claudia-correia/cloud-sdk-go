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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewResyncAllocatorsParams creates a new ResyncAllocatorsParams object
// with the default values initialized.
func NewResyncAllocatorsParams() *ResyncAllocatorsParams {
	var (
		skipMatchingVersionDefault = bool(true)
	)
	return &ResyncAllocatorsParams{
		SkipMatchingVersion: &skipMatchingVersionDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewResyncAllocatorsParamsWithTimeout creates a new ResyncAllocatorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResyncAllocatorsParamsWithTimeout(timeout time.Duration) *ResyncAllocatorsParams {
	var (
		skipMatchingVersionDefault = bool(true)
	)
	return &ResyncAllocatorsParams{
		SkipMatchingVersion: &skipMatchingVersionDefault,

		timeout: timeout,
	}
}

// NewResyncAllocatorsParamsWithContext creates a new ResyncAllocatorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewResyncAllocatorsParamsWithContext(ctx context.Context) *ResyncAllocatorsParams {
	var (
		skipMatchingVersionDefault = bool(true)
	)
	return &ResyncAllocatorsParams{
		SkipMatchingVersion: &skipMatchingVersionDefault,

		Context: ctx,
	}
}

// NewResyncAllocatorsParamsWithHTTPClient creates a new ResyncAllocatorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResyncAllocatorsParamsWithHTTPClient(client *http.Client) *ResyncAllocatorsParams {
	var (
		skipMatchingVersionDefault = bool(true)
	)
	return &ResyncAllocatorsParams{
		SkipMatchingVersion: &skipMatchingVersionDefault,
		HTTPClient:          client,
	}
}

/*ResyncAllocatorsParams contains all the parameters to send to the API endpoint
for the resync allocators operation typically these are written to a http.Request
*/
type ResyncAllocatorsParams struct {

	/*SkipMatchingVersion
	  When true, skips the document indexing when the version matches the in-memory copy.

	*/
	SkipMatchingVersion *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the resync allocators params
func (o *ResyncAllocatorsParams) WithTimeout(timeout time.Duration) *ResyncAllocatorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resync allocators params
func (o *ResyncAllocatorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resync allocators params
func (o *ResyncAllocatorsParams) WithContext(ctx context.Context) *ResyncAllocatorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resync allocators params
func (o *ResyncAllocatorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resync allocators params
func (o *ResyncAllocatorsParams) WithHTTPClient(client *http.Client) *ResyncAllocatorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resync allocators params
func (o *ResyncAllocatorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSkipMatchingVersion adds the skipMatchingVersion to the resync allocators params
func (o *ResyncAllocatorsParams) WithSkipMatchingVersion(skipMatchingVersion *bool) *ResyncAllocatorsParams {
	o.SetSkipMatchingVersion(skipMatchingVersion)
	return o
}

// SetSkipMatchingVersion adds the skipMatchingVersion to the resync allocators params
func (o *ResyncAllocatorsParams) SetSkipMatchingVersion(skipMatchingVersion *bool) {
	o.SkipMatchingVersion = skipMatchingVersion
}

// WriteToRequest writes these params to a swagger request
func (o *ResyncAllocatorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SkipMatchingVersion != nil {

		// query param skip_matching_version
		var qrSkipMatchingVersion bool
		if o.SkipMatchingVersion != nil {
			qrSkipMatchingVersion = *o.SkipMatchingVersion
		}
		qSkipMatchingVersion := swag.FormatBool(qrSkipMatchingVersion)
		if qSkipMatchingVersion != "" {
			if err := r.SetQueryParam("skip_matching_version", qSkipMatchingVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
