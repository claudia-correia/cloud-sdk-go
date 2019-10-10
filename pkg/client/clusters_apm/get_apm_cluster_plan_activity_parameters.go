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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetApmClusterPlanActivityParams creates a new GetApmClusterPlanActivityParams object
// with the default values initialized.
func NewGetApmClusterPlanActivityParams() *GetApmClusterPlanActivityParams {
	var (
		showPlanDefaultsDefault = bool(false)
		showPlanLogsDefault     = bool(true)
	)
	return &GetApmClusterPlanActivityParams{
		ShowPlanDefaults: &showPlanDefaultsDefault,
		ShowPlanLogs:     &showPlanLogsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetApmClusterPlanActivityParamsWithTimeout creates a new GetApmClusterPlanActivityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetApmClusterPlanActivityParamsWithTimeout(timeout time.Duration) *GetApmClusterPlanActivityParams {
	var (
		showPlanDefaultsDefault = bool(false)
		showPlanLogsDefault     = bool(true)
	)
	return &GetApmClusterPlanActivityParams{
		ShowPlanDefaults: &showPlanDefaultsDefault,
		ShowPlanLogs:     &showPlanLogsDefault,

		timeout: timeout,
	}
}

// NewGetApmClusterPlanActivityParamsWithContext creates a new GetApmClusterPlanActivityParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetApmClusterPlanActivityParamsWithContext(ctx context.Context) *GetApmClusterPlanActivityParams {
	var (
		showPlanDefaultsDefault = bool(false)
		showPlanLogsDefault     = bool(true)
	)
	return &GetApmClusterPlanActivityParams{
		ShowPlanDefaults: &showPlanDefaultsDefault,
		ShowPlanLogs:     &showPlanLogsDefault,

		Context: ctx,
	}
}

// NewGetApmClusterPlanActivityParamsWithHTTPClient creates a new GetApmClusterPlanActivityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetApmClusterPlanActivityParamsWithHTTPClient(client *http.Client) *GetApmClusterPlanActivityParams {
	var (
		showPlanDefaultsDefault = bool(false)
		showPlanLogsDefault     = bool(true)
	)
	return &GetApmClusterPlanActivityParams{
		ShowPlanDefaults: &showPlanDefaultsDefault,
		ShowPlanLogs:     &showPlanLogsDefault,
		HTTPClient:       client,
	}
}

/*GetApmClusterPlanActivityParams contains all the parameters to send to the API endpoint
for the get apm cluster plan activity operation typically these are written to a http.Request
*/
type GetApmClusterPlanActivityParams struct {

	/*ClusterID
	  The APM deployment identifier.

	*/
	ClusterID string
	/*ShowPlanDefaults
	  When plans are shown, includes the default values in the response. NOTE: This option results in large responses.

	*/
	ShowPlanDefaults *bool
	/*ShowPlanLogs
	  Includes the active, pending, and historical plan information in the attempt log. NOTE: This option can result in large responses.

	*/
	ShowPlanLogs *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get apm cluster plan activity params
func (o *GetApmClusterPlanActivityParams) WithTimeout(timeout time.Duration) *GetApmClusterPlanActivityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get apm cluster plan activity params
func (o *GetApmClusterPlanActivityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get apm cluster plan activity params
func (o *GetApmClusterPlanActivityParams) WithContext(ctx context.Context) *GetApmClusterPlanActivityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get apm cluster plan activity params
func (o *GetApmClusterPlanActivityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get apm cluster plan activity params
func (o *GetApmClusterPlanActivityParams) WithHTTPClient(client *http.Client) *GetApmClusterPlanActivityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get apm cluster plan activity params
func (o *GetApmClusterPlanActivityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get apm cluster plan activity params
func (o *GetApmClusterPlanActivityParams) WithClusterID(clusterID string) *GetApmClusterPlanActivityParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get apm cluster plan activity params
func (o *GetApmClusterPlanActivityParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithShowPlanDefaults adds the showPlanDefaults to the get apm cluster plan activity params
func (o *GetApmClusterPlanActivityParams) WithShowPlanDefaults(showPlanDefaults *bool) *GetApmClusterPlanActivityParams {
	o.SetShowPlanDefaults(showPlanDefaults)
	return o
}

// SetShowPlanDefaults adds the showPlanDefaults to the get apm cluster plan activity params
func (o *GetApmClusterPlanActivityParams) SetShowPlanDefaults(showPlanDefaults *bool) {
	o.ShowPlanDefaults = showPlanDefaults
}

// WithShowPlanLogs adds the showPlanLogs to the get apm cluster plan activity params
func (o *GetApmClusterPlanActivityParams) WithShowPlanLogs(showPlanLogs *bool) *GetApmClusterPlanActivityParams {
	o.SetShowPlanLogs(showPlanLogs)
	return o
}

// SetShowPlanLogs adds the showPlanLogs to the get apm cluster plan activity params
func (o *GetApmClusterPlanActivityParams) SetShowPlanLogs(showPlanLogs *bool) {
	o.ShowPlanLogs = showPlanLogs
}

// WriteToRequest writes these params to a swagger request
func (o *GetApmClusterPlanActivityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.ShowPlanDefaults != nil {

		// query param show_plan_defaults
		var qrShowPlanDefaults bool
		if o.ShowPlanDefaults != nil {
			qrShowPlanDefaults = *o.ShowPlanDefaults
		}
		qShowPlanDefaults := swag.FormatBool(qrShowPlanDefaults)
		if qShowPlanDefaults != "" {
			if err := r.SetQueryParam("show_plan_defaults", qShowPlanDefaults); err != nil {
				return err
			}
		}

	}

	if o.ShowPlanLogs != nil {

		// query param show_plan_logs
		var qrShowPlanLogs bool
		if o.ShowPlanLogs != nil {
			qrShowPlanLogs = *o.ShowPlanLogs
		}
		qShowPlanLogs := swag.FormatBool(qrShowPlanLogs)
		if qShowPlanLogs != "" {
			if err := r.SetQueryParam("show_plan_logs", qShowPlanLogs); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
