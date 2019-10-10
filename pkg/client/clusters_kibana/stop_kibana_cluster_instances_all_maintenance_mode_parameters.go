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

package clusters_kibana

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

// NewStopKibanaClusterInstancesAllMaintenanceModeParams creates a new StopKibanaClusterInstancesAllMaintenanceModeParams object
// with the default values initialized.
func NewStopKibanaClusterInstancesAllMaintenanceModeParams() *StopKibanaClusterInstancesAllMaintenanceModeParams {
	var ()
	return &StopKibanaClusterInstancesAllMaintenanceModeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopKibanaClusterInstancesAllMaintenanceModeParamsWithTimeout creates a new StopKibanaClusterInstancesAllMaintenanceModeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopKibanaClusterInstancesAllMaintenanceModeParamsWithTimeout(timeout time.Duration) *StopKibanaClusterInstancesAllMaintenanceModeParams {
	var ()
	return &StopKibanaClusterInstancesAllMaintenanceModeParams{

		timeout: timeout,
	}
}

// NewStopKibanaClusterInstancesAllMaintenanceModeParamsWithContext creates a new StopKibanaClusterInstancesAllMaintenanceModeParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopKibanaClusterInstancesAllMaintenanceModeParamsWithContext(ctx context.Context) *StopKibanaClusterInstancesAllMaintenanceModeParams {
	var ()
	return &StopKibanaClusterInstancesAllMaintenanceModeParams{

		Context: ctx,
	}
}

// NewStopKibanaClusterInstancesAllMaintenanceModeParamsWithHTTPClient creates a new StopKibanaClusterInstancesAllMaintenanceModeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopKibanaClusterInstancesAllMaintenanceModeParamsWithHTTPClient(client *http.Client) *StopKibanaClusterInstancesAllMaintenanceModeParams {
	var ()
	return &StopKibanaClusterInstancesAllMaintenanceModeParams{
		HTTPClient: client,
	}
}

/*StopKibanaClusterInstancesAllMaintenanceModeParams contains all the parameters to send to the API endpoint
for the stop kibana cluster instances all maintenance mode operation typically these are written to a http.Request
*/
type StopKibanaClusterInstancesAllMaintenanceModeParams struct {

	/*ClusterID
	  The Kibana deployment identifier.

	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop kibana cluster instances all maintenance mode params
func (o *StopKibanaClusterInstancesAllMaintenanceModeParams) WithTimeout(timeout time.Duration) *StopKibanaClusterInstancesAllMaintenanceModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop kibana cluster instances all maintenance mode params
func (o *StopKibanaClusterInstancesAllMaintenanceModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop kibana cluster instances all maintenance mode params
func (o *StopKibanaClusterInstancesAllMaintenanceModeParams) WithContext(ctx context.Context) *StopKibanaClusterInstancesAllMaintenanceModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop kibana cluster instances all maintenance mode params
func (o *StopKibanaClusterInstancesAllMaintenanceModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop kibana cluster instances all maintenance mode params
func (o *StopKibanaClusterInstancesAllMaintenanceModeParams) WithHTTPClient(client *http.Client) *StopKibanaClusterInstancesAllMaintenanceModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop kibana cluster instances all maintenance mode params
func (o *StopKibanaClusterInstancesAllMaintenanceModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the stop kibana cluster instances all maintenance mode params
func (o *StopKibanaClusterInstancesAllMaintenanceModeParams) WithClusterID(clusterID string) *StopKibanaClusterInstancesAllMaintenanceModeParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the stop kibana cluster instances all maintenance mode params
func (o *StopKibanaClusterInstancesAllMaintenanceModeParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *StopKibanaClusterInstancesAllMaintenanceModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
