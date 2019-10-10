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

package authentication

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

// NewSamlCallbackParams creates a new SamlCallbackParams object
// with the default values initialized.
func NewSamlCallbackParams() *SamlCallbackParams {
	var ()
	return &SamlCallbackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSamlCallbackParamsWithTimeout creates a new SamlCallbackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSamlCallbackParamsWithTimeout(timeout time.Duration) *SamlCallbackParams {
	var ()
	return &SamlCallbackParams{

		timeout: timeout,
	}
}

// NewSamlCallbackParamsWithContext creates a new SamlCallbackParams object
// with the default values initialized, and the ability to set a context for a request
func NewSamlCallbackParamsWithContext(ctx context.Context) *SamlCallbackParams {
	var ()
	return &SamlCallbackParams{

		Context: ctx,
	}
}

// NewSamlCallbackParamsWithHTTPClient creates a new SamlCallbackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSamlCallbackParamsWithHTTPClient(client *http.Client) *SamlCallbackParams {
	var ()
	return &SamlCallbackParams{
		HTTPClient: client,
	}
}

/*SamlCallbackParams contains all the parameters to send to the API endpoint
for the saml callback operation typically these are written to a http.Request
*/
type SamlCallbackParams struct {

	/*RelayState
	  The optional relay state that the API (service provider) sent to the identity provider.

	*/
	RelayState *string
	/*SAMLResponse
	  A message issued by the identity provider to the service provider

	*/
	SAMLResponse string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the saml callback params
func (o *SamlCallbackParams) WithTimeout(timeout time.Duration) *SamlCallbackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the saml callback params
func (o *SamlCallbackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the saml callback params
func (o *SamlCallbackParams) WithContext(ctx context.Context) *SamlCallbackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the saml callback params
func (o *SamlCallbackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the saml callback params
func (o *SamlCallbackParams) WithHTTPClient(client *http.Client) *SamlCallbackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the saml callback params
func (o *SamlCallbackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRelayState adds the relayState to the saml callback params
func (o *SamlCallbackParams) WithRelayState(relayState *string) *SamlCallbackParams {
	o.SetRelayState(relayState)
	return o
}

// SetRelayState adds the relayState to the saml callback params
func (o *SamlCallbackParams) SetRelayState(relayState *string) {
	o.RelayState = relayState
}

// WithSAMLResponse adds the sAMLResponse to the saml callback params
func (o *SamlCallbackParams) WithSAMLResponse(sAMLResponse string) *SamlCallbackParams {
	o.SetSAMLResponse(sAMLResponse)
	return o
}

// SetSAMLResponse adds the sAMLResponse to the saml callback params
func (o *SamlCallbackParams) SetSAMLResponse(sAMLResponse string) {
	o.SAMLResponse = sAMLResponse
}

// WriteToRequest writes these params to a swagger request
func (o *SamlCallbackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RelayState != nil {

		// form param RelayState
		var frRelayState string
		if o.RelayState != nil {
			frRelayState = *o.RelayState
		}
		fRelayState := frRelayState
		if fRelayState != "" {
			if err := r.SetFormParam("RelayState", fRelayState); err != nil {
				return err
			}
		}

	}

	// form param SAMLResponse
	frSAMLResponse := o.SAMLResponse
	fSAMLResponse := frSAMLResponse
	if fSAMLResponse != "" {
		if err := r.SetFormParam("SAMLResponse", fSAMLResponse); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
