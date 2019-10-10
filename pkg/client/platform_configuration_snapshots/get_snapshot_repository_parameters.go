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

package platform_configuration_snapshots

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

// NewGetSnapshotRepositoryParams creates a new GetSnapshotRepositoryParams object
// with the default values initialized.
func NewGetSnapshotRepositoryParams() *GetSnapshotRepositoryParams {
	var ()
	return &GetSnapshotRepositoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSnapshotRepositoryParamsWithTimeout creates a new GetSnapshotRepositoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSnapshotRepositoryParamsWithTimeout(timeout time.Duration) *GetSnapshotRepositoryParams {
	var ()
	return &GetSnapshotRepositoryParams{

		timeout: timeout,
	}
}

// NewGetSnapshotRepositoryParamsWithContext creates a new GetSnapshotRepositoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSnapshotRepositoryParamsWithContext(ctx context.Context) *GetSnapshotRepositoryParams {
	var ()
	return &GetSnapshotRepositoryParams{

		Context: ctx,
	}
}

// NewGetSnapshotRepositoryParamsWithHTTPClient creates a new GetSnapshotRepositoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSnapshotRepositoryParamsWithHTTPClient(client *http.Client) *GetSnapshotRepositoryParams {
	var ()
	return &GetSnapshotRepositoryParams{
		HTTPClient: client,
	}
}

/*GetSnapshotRepositoryParams contains all the parameters to send to the API endpoint
for the get snapshot repository operation typically these are written to a http.Request
*/
type GetSnapshotRepositoryParams struct {

	/*RepositoryName
	  The name of the snapshot repository configuration.

	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get snapshot repository params
func (o *GetSnapshotRepositoryParams) WithTimeout(timeout time.Duration) *GetSnapshotRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get snapshot repository params
func (o *GetSnapshotRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get snapshot repository params
func (o *GetSnapshotRepositoryParams) WithContext(ctx context.Context) *GetSnapshotRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get snapshot repository params
func (o *GetSnapshotRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get snapshot repository params
func (o *GetSnapshotRepositoryParams) WithHTTPClient(client *http.Client) *GetSnapshotRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get snapshot repository params
func (o *GetSnapshotRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepositoryName adds the repositoryName to the get snapshot repository params
func (o *GetSnapshotRepositoryParams) WithRepositoryName(repositoryName string) *GetSnapshotRepositoryParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the get snapshot repository params
func (o *GetSnapshotRepositoryParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *GetSnapshotRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repository_name
	if err := r.SetPathParam("repository_name", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
