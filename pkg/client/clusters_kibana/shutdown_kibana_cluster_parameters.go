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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewShutdownKibanaClusterParams creates a new ShutdownKibanaClusterParams object
// with the default values initialized.
func NewShutdownKibanaClusterParams() *ShutdownKibanaClusterParams {
	var (
		hideDefault = bool(false)
	)
	return &ShutdownKibanaClusterParams{
		Hide: &hideDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewShutdownKibanaClusterParamsWithTimeout creates a new ShutdownKibanaClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShutdownKibanaClusterParamsWithTimeout(timeout time.Duration) *ShutdownKibanaClusterParams {
	var (
		hideDefault = bool(false)
	)
	return &ShutdownKibanaClusterParams{
		Hide: &hideDefault,

		timeout: timeout,
	}
}

// NewShutdownKibanaClusterParamsWithContext creates a new ShutdownKibanaClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewShutdownKibanaClusterParamsWithContext(ctx context.Context) *ShutdownKibanaClusterParams {
	var (
		hideDefault = bool(false)
	)
	return &ShutdownKibanaClusterParams{
		Hide: &hideDefault,

		Context: ctx,
	}
}

// NewShutdownKibanaClusterParamsWithHTTPClient creates a new ShutdownKibanaClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShutdownKibanaClusterParamsWithHTTPClient(client *http.Client) *ShutdownKibanaClusterParams {
	var (
		hideDefault = bool(false)
	)
	return &ShutdownKibanaClusterParams{
		Hide:       &hideDefault,
		HTTPClient: client,
	}
}

/*ShutdownKibanaClusterParams contains all the parameters to send to the API endpoint
for the shutdown kibana cluster operation typically these are written to a http.Request
*/
type ShutdownKibanaClusterParams struct {

	/*ClusterID
	  The Kibana deployment identifier.

	*/
	ClusterID string
	/*Hide
	  Hides the clusters during shutdown. NOTE: By default, hidden clusters are not listed.

	*/
	Hide *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the shutdown kibana cluster params
func (o *ShutdownKibanaClusterParams) WithTimeout(timeout time.Duration) *ShutdownKibanaClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the shutdown kibana cluster params
func (o *ShutdownKibanaClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the shutdown kibana cluster params
func (o *ShutdownKibanaClusterParams) WithContext(ctx context.Context) *ShutdownKibanaClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the shutdown kibana cluster params
func (o *ShutdownKibanaClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the shutdown kibana cluster params
func (o *ShutdownKibanaClusterParams) WithHTTPClient(client *http.Client) *ShutdownKibanaClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the shutdown kibana cluster params
func (o *ShutdownKibanaClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the shutdown kibana cluster params
func (o *ShutdownKibanaClusterParams) WithClusterID(clusterID string) *ShutdownKibanaClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the shutdown kibana cluster params
func (o *ShutdownKibanaClusterParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithHide adds the hide to the shutdown kibana cluster params
func (o *ShutdownKibanaClusterParams) WithHide(hide *bool) *ShutdownKibanaClusterParams {
	o.SetHide(hide)
	return o
}

// SetHide adds the hide to the shutdown kibana cluster params
func (o *ShutdownKibanaClusterParams) SetHide(hide *bool) {
	o.Hide = hide
}

// WriteToRequest writes these params to a swagger request
func (o *ShutdownKibanaClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.Hide != nil {

		// query param hide
		var qrHide bool
		if o.Hide != nil {
			qrHide = *o.Hide
		}
		qHide := swag.FormatBool(qrHide)
		if qHide != "" {
			if err := r.SetQueryParam("hide", qHide); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
