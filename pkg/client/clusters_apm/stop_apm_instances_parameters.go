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

// NewStopApmInstancesParams creates a new StopApmInstancesParams object
// with the default values initialized.
func NewStopApmInstancesParams() *StopApmInstancesParams {
	var (
		ignoreMissingDefault = bool(false)
	)
	return &StopApmInstancesParams{
		IgnoreMissing: &ignoreMissingDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewStopApmInstancesParamsWithTimeout creates a new StopApmInstancesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopApmInstancesParamsWithTimeout(timeout time.Duration) *StopApmInstancesParams {
	var (
		ignoreMissingDefault = bool(false)
	)
	return &StopApmInstancesParams{
		IgnoreMissing: &ignoreMissingDefault,

		timeout: timeout,
	}
}

// NewStopApmInstancesParamsWithContext creates a new StopApmInstancesParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopApmInstancesParamsWithContext(ctx context.Context) *StopApmInstancesParams {
	var (
		ignoreMissingDefault = bool(false)
	)
	return &StopApmInstancesParams{
		IgnoreMissing: &ignoreMissingDefault,

		Context: ctx,
	}
}

// NewStopApmInstancesParamsWithHTTPClient creates a new StopApmInstancesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopApmInstancesParamsWithHTTPClient(client *http.Client) *StopApmInstancesParams {
	var (
		ignoreMissingDefault = bool(false)
	)
	return &StopApmInstancesParams{
		IgnoreMissing: &ignoreMissingDefault,
		HTTPClient:    client,
	}
}

/*StopApmInstancesParams contains all the parameters to send to the API endpoint
for the stop apm instances operation typically these are written to a http.Request
*/
type StopApmInstancesParams struct {

	/*ClusterID
	  The APM deployment identifier.

	*/
	ClusterID string
	/*IgnoreMissing
	  When `true` and the instance does not exist, proceeds to the next instance, or treats the instance as an error.

	*/
	IgnoreMissing *bool
	/*InstanceIds
	  A comma-separated list of instance identifiers.

	*/
	InstanceIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop apm instances params
func (o *StopApmInstancesParams) WithTimeout(timeout time.Duration) *StopApmInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop apm instances params
func (o *StopApmInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop apm instances params
func (o *StopApmInstancesParams) WithContext(ctx context.Context) *StopApmInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop apm instances params
func (o *StopApmInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop apm instances params
func (o *StopApmInstancesParams) WithHTTPClient(client *http.Client) *StopApmInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop apm instances params
func (o *StopApmInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the stop apm instances params
func (o *StopApmInstancesParams) WithClusterID(clusterID string) *StopApmInstancesParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the stop apm instances params
func (o *StopApmInstancesParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithIgnoreMissing adds the ignoreMissing to the stop apm instances params
func (o *StopApmInstancesParams) WithIgnoreMissing(ignoreMissing *bool) *StopApmInstancesParams {
	o.SetIgnoreMissing(ignoreMissing)
	return o
}

// SetIgnoreMissing adds the ignoreMissing to the stop apm instances params
func (o *StopApmInstancesParams) SetIgnoreMissing(ignoreMissing *bool) {
	o.IgnoreMissing = ignoreMissing
}

// WithInstanceIds adds the instanceIds to the stop apm instances params
func (o *StopApmInstancesParams) WithInstanceIds(instanceIds []string) *StopApmInstancesParams {
	o.SetInstanceIds(instanceIds)
	return o
}

// SetInstanceIds adds the instanceIds to the stop apm instances params
func (o *StopApmInstancesParams) SetInstanceIds(instanceIds []string) {
	o.InstanceIds = instanceIds
}

// WriteToRequest writes these params to a swagger request
func (o *StopApmInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.IgnoreMissing != nil {

		// query param ignore_missing
		var qrIgnoreMissing bool
		if o.IgnoreMissing != nil {
			qrIgnoreMissing = *o.IgnoreMissing
		}
		qIgnoreMissing := swag.FormatBool(qrIgnoreMissing)
		if qIgnoreMissing != "" {
			if err := r.SetQueryParam("ignore_missing", qIgnoreMissing); err != nil {
				return err
			}
		}

	}

	valuesInstanceIds := o.InstanceIds

	joinedInstanceIds := swag.JoinByFormat(valuesInstanceIds, "csv")
	// path array param instance_ids
	// SetPathParam does not support variadric arguments, since we used JoinByFormat
	// we can send the first item in the array as it's all the items of the previous
	// array joined together
	if len(joinedInstanceIds) > 0 {
		if err := r.SetPathParam("instance_ids", joinedInstanceIds[0]); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
