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

package platform_configuration_security

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

// NewGetTLSCertificateParams creates a new GetTLSCertificateParams object
// with the default values initialized.
func NewGetTLSCertificateParams() *GetTLSCertificateParams {
	var ()
	return &GetTLSCertificateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTLSCertificateParamsWithTimeout creates a new GetTLSCertificateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTLSCertificateParamsWithTimeout(timeout time.Duration) *GetTLSCertificateParams {
	var ()
	return &GetTLSCertificateParams{

		timeout: timeout,
	}
}

// NewGetTLSCertificateParamsWithContext creates a new GetTLSCertificateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTLSCertificateParamsWithContext(ctx context.Context) *GetTLSCertificateParams {
	var ()
	return &GetTLSCertificateParams{

		Context: ctx,
	}
}

// NewGetTLSCertificateParamsWithHTTPClient creates a new GetTLSCertificateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTLSCertificateParamsWithHTTPClient(client *http.Client) *GetTLSCertificateParams {
	var ()
	return &GetTLSCertificateParams{
		HTTPClient: client,
	}
}

/*GetTLSCertificateParams contains all the parameters to send to the API endpoint
for the get tls certificate operation typically these are written to a http.Request
*/
type GetTLSCertificateParams struct {

	/*ServiceName
	  The service certificate chain to read.

	*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tls certificate params
func (o *GetTLSCertificateParams) WithTimeout(timeout time.Duration) *GetTLSCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tls certificate params
func (o *GetTLSCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tls certificate params
func (o *GetTLSCertificateParams) WithContext(ctx context.Context) *GetTLSCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tls certificate params
func (o *GetTLSCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tls certificate params
func (o *GetTLSCertificateParams) WithHTTPClient(client *http.Client) *GetTLSCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tls certificate params
func (o *GetTLSCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceName adds the serviceName to the get tls certificate params
func (o *GetTLSCertificateParams) WithServiceName(serviceName string) *GetTLSCertificateParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get tls certificate params
func (o *GetTLSCertificateParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetTLSCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param service_name
	if err := r.SetPathParam("service_name", o.ServiceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
