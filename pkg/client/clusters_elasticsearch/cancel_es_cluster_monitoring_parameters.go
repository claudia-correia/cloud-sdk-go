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

package clusters_elasticsearch

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

// NewCancelEsClusterMonitoringParams creates a new CancelEsClusterMonitoringParams object
// with the default values initialized.
func NewCancelEsClusterMonitoringParams() *CancelEsClusterMonitoringParams {
	var ()
	return &CancelEsClusterMonitoringParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCancelEsClusterMonitoringParamsWithTimeout creates a new CancelEsClusterMonitoringParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCancelEsClusterMonitoringParamsWithTimeout(timeout time.Duration) *CancelEsClusterMonitoringParams {
	var ()
	return &CancelEsClusterMonitoringParams{

		timeout: timeout,
	}
}

// NewCancelEsClusterMonitoringParamsWithContext creates a new CancelEsClusterMonitoringParams object
// with the default values initialized, and the ability to set a context for a request
func NewCancelEsClusterMonitoringParamsWithContext(ctx context.Context) *CancelEsClusterMonitoringParams {
	var ()
	return &CancelEsClusterMonitoringParams{

		Context: ctx,
	}
}

// NewCancelEsClusterMonitoringParamsWithHTTPClient creates a new CancelEsClusterMonitoringParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCancelEsClusterMonitoringParamsWithHTTPClient(client *http.Client) *CancelEsClusterMonitoringParams {
	var ()
	return &CancelEsClusterMonitoringParams{
		HTTPClient: client,
	}
}

/*CancelEsClusterMonitoringParams contains all the parameters to send to the API endpoint
for the cancel es cluster monitoring operation typically these are written to a http.Request
*/
type CancelEsClusterMonitoringParams struct {

	/*ClusterID
	  Elasticsearch cluster identifier

	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cancel es cluster monitoring params
func (o *CancelEsClusterMonitoringParams) WithTimeout(timeout time.Duration) *CancelEsClusterMonitoringParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel es cluster monitoring params
func (o *CancelEsClusterMonitoringParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel es cluster monitoring params
func (o *CancelEsClusterMonitoringParams) WithContext(ctx context.Context) *CancelEsClusterMonitoringParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel es cluster monitoring params
func (o *CancelEsClusterMonitoringParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel es cluster monitoring params
func (o *CancelEsClusterMonitoringParams) WithHTTPClient(client *http.Client) *CancelEsClusterMonitoringParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel es cluster monitoring params
func (o *CancelEsClusterMonitoringParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the cancel es cluster monitoring params
func (o *CancelEsClusterMonitoringParams) WithClusterID(clusterID string) *CancelEsClusterMonitoringParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the cancel es cluster monitoring params
func (o *CancelEsClusterMonitoringParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *CancelEsClusterMonitoringParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
