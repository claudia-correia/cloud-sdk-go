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

// NewSetEsClusterMonitoringParams creates a new SetEsClusterMonitoringParams object
// with the default values initialized.
func NewSetEsClusterMonitoringParams() *SetEsClusterMonitoringParams {
	var ()
	return &SetEsClusterMonitoringParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetEsClusterMonitoringParamsWithTimeout creates a new SetEsClusterMonitoringParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetEsClusterMonitoringParamsWithTimeout(timeout time.Duration) *SetEsClusterMonitoringParams {
	var ()
	return &SetEsClusterMonitoringParams{

		timeout: timeout,
	}
}

// NewSetEsClusterMonitoringParamsWithContext creates a new SetEsClusterMonitoringParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetEsClusterMonitoringParamsWithContext(ctx context.Context) *SetEsClusterMonitoringParams {
	var ()
	return &SetEsClusterMonitoringParams{

		Context: ctx,
	}
}

// NewSetEsClusterMonitoringParamsWithHTTPClient creates a new SetEsClusterMonitoringParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetEsClusterMonitoringParamsWithHTTPClient(client *http.Client) *SetEsClusterMonitoringParams {
	var ()
	return &SetEsClusterMonitoringParams{
		HTTPClient: client,
	}
}

/*SetEsClusterMonitoringParams contains all the parameters to send to the API endpoint
for the set es cluster monitoring operation typically these are written to a http.Request
*/
type SetEsClusterMonitoringParams struct {

	/*ClusterID
	  Elasticsearch cluster identifier

	*/
	ClusterID string
	/*DestClusterID
	  The Elasticsearch cluster identifier for the monitoring destination.

	*/
	DestClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set es cluster monitoring params
func (o *SetEsClusterMonitoringParams) WithTimeout(timeout time.Duration) *SetEsClusterMonitoringParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set es cluster monitoring params
func (o *SetEsClusterMonitoringParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set es cluster monitoring params
func (o *SetEsClusterMonitoringParams) WithContext(ctx context.Context) *SetEsClusterMonitoringParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set es cluster monitoring params
func (o *SetEsClusterMonitoringParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set es cluster monitoring params
func (o *SetEsClusterMonitoringParams) WithHTTPClient(client *http.Client) *SetEsClusterMonitoringParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set es cluster monitoring params
func (o *SetEsClusterMonitoringParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the set es cluster monitoring params
func (o *SetEsClusterMonitoringParams) WithClusterID(clusterID string) *SetEsClusterMonitoringParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the set es cluster monitoring params
func (o *SetEsClusterMonitoringParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDestClusterID adds the destClusterID to the set es cluster monitoring params
func (o *SetEsClusterMonitoringParams) WithDestClusterID(destClusterID string) *SetEsClusterMonitoringParams {
	o.SetDestClusterID(destClusterID)
	return o
}

// SetDestClusterID adds the destClusterId to the set es cluster monitoring params
func (o *SetEsClusterMonitoringParams) SetDestClusterID(destClusterID string) {
	o.DestClusterID = destClusterID
}

// WriteToRequest writes these params to a swagger request
func (o *SetEsClusterMonitoringParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dest_cluster_id
	if err := r.SetPathParam("dest_cluster_id", o.DestClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
