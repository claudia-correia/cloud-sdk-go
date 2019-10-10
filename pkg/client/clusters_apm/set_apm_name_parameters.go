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

package clusters_apm

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

// NewSetApmNameParams creates a new SetApmNameParams object
// with the default values initialized.
func NewSetApmNameParams() *SetApmNameParams {
	var ()
	return &SetApmNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetApmNameParamsWithTimeout creates a new SetApmNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetApmNameParamsWithTimeout(timeout time.Duration) *SetApmNameParams {
	var ()
	return &SetApmNameParams{

		timeout: timeout,
	}
}

// NewSetApmNameParamsWithContext creates a new SetApmNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetApmNameParamsWithContext(ctx context.Context) *SetApmNameParams {
	var ()
	return &SetApmNameParams{

		Context: ctx,
	}
}

// NewSetApmNameParamsWithHTTPClient creates a new SetApmNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetApmNameParamsWithHTTPClient(client *http.Client) *SetApmNameParams {
	var ()
	return &SetApmNameParams{
		HTTPClient: client,
	}
}

/*SetApmNameParams contains all the parameters to send to the API endpoint
for the set apm name operation typically these are written to a http.Request
*/
type SetApmNameParams struct {

	/*ClusterID
	  The APM deployment identifier.

	*/
	ClusterID string
	/*NewName
	  The new name for the cluster.

	*/
	NewName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set apm name params
func (o *SetApmNameParams) WithTimeout(timeout time.Duration) *SetApmNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set apm name params
func (o *SetApmNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set apm name params
func (o *SetApmNameParams) WithContext(ctx context.Context) *SetApmNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set apm name params
func (o *SetApmNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set apm name params
func (o *SetApmNameParams) WithHTTPClient(client *http.Client) *SetApmNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set apm name params
func (o *SetApmNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the set apm name params
func (o *SetApmNameParams) WithClusterID(clusterID string) *SetApmNameParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the set apm name params
func (o *SetApmNameParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithNewName adds the newName to the set apm name params
func (o *SetApmNameParams) WithNewName(newName string) *SetApmNameParams {
	o.SetNewName(newName)
	return o
}

// SetNewName adds the newName to the set apm name params
func (o *SetApmNameParams) SetNewName(newName string) {
	o.NewName = newName
}

// WriteToRequest writes these params to a swagger request
func (o *SetApmNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param new_name
	if err := r.SetPathParam("new_name", o.NewName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
