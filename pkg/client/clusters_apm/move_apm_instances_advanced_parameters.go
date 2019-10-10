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

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewMoveApmInstancesAdvancedParams creates a new MoveApmInstancesAdvancedParams object
// with the default values initialized.
func NewMoveApmInstancesAdvancedParams() *MoveApmInstancesAdvancedParams {
	var (
		forceUpdateDefault  = bool(false)
		validateOnlyDefault = bool(false)
	)
	return &MoveApmInstancesAdvancedParams{
		ForceUpdate:  &forceUpdateDefault,
		ValidateOnly: &validateOnlyDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewMoveApmInstancesAdvancedParamsWithTimeout creates a new MoveApmInstancesAdvancedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMoveApmInstancesAdvancedParamsWithTimeout(timeout time.Duration) *MoveApmInstancesAdvancedParams {
	var (
		forceUpdateDefault  = bool(false)
		validateOnlyDefault = bool(false)
	)
	return &MoveApmInstancesAdvancedParams{
		ForceUpdate:  &forceUpdateDefault,
		ValidateOnly: &validateOnlyDefault,

		timeout: timeout,
	}
}

// NewMoveApmInstancesAdvancedParamsWithContext creates a new MoveApmInstancesAdvancedParams object
// with the default values initialized, and the ability to set a context for a request
func NewMoveApmInstancesAdvancedParamsWithContext(ctx context.Context) *MoveApmInstancesAdvancedParams {
	var (
		forceUpdateDefault  = bool(false)
		validateOnlyDefault = bool(false)
	)
	return &MoveApmInstancesAdvancedParams{
		ForceUpdate:  &forceUpdateDefault,
		ValidateOnly: &validateOnlyDefault,

		Context: ctx,
	}
}

// NewMoveApmInstancesAdvancedParamsWithHTTPClient creates a new MoveApmInstancesAdvancedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMoveApmInstancesAdvancedParamsWithHTTPClient(client *http.Client) *MoveApmInstancesAdvancedParams {
	var (
		forceUpdateDefault  = bool(false)
		validateOnlyDefault = bool(false)
	)
	return &MoveApmInstancesAdvancedParams{
		ForceUpdate:  &forceUpdateDefault,
		ValidateOnly: &validateOnlyDefault,
		HTTPClient:   client,
	}
}

/*MoveApmInstancesAdvancedParams contains all the parameters to send to the API endpoint
for the move apm instances advanced operation typically these are written to a http.Request
*/
type MoveApmInstancesAdvancedParams struct {

	/*Body
	  Overrides defaults for the move, including setting the configuration of instances specified in the path

	*/
	Body *models.TransientApmPlanConfiguration
	/*ClusterID
	  The APM deployment identifier.

	*/
	ClusterID string
	/*ForceUpdate
	  When `true`, cancels and overwrites the pending plans, or treats the instance as an error.

	*/
	ForceUpdate *bool
	/*ValidateOnly
	  When `true`, validates the move request, then returns the calculated plan without applying the plan.

	*/
	ValidateOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the move apm instances advanced params
func (o *MoveApmInstancesAdvancedParams) WithTimeout(timeout time.Duration) *MoveApmInstancesAdvancedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the move apm instances advanced params
func (o *MoveApmInstancesAdvancedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the move apm instances advanced params
func (o *MoveApmInstancesAdvancedParams) WithContext(ctx context.Context) *MoveApmInstancesAdvancedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the move apm instances advanced params
func (o *MoveApmInstancesAdvancedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the move apm instances advanced params
func (o *MoveApmInstancesAdvancedParams) WithHTTPClient(client *http.Client) *MoveApmInstancesAdvancedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the move apm instances advanced params
func (o *MoveApmInstancesAdvancedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the move apm instances advanced params
func (o *MoveApmInstancesAdvancedParams) WithBody(body *models.TransientApmPlanConfiguration) *MoveApmInstancesAdvancedParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the move apm instances advanced params
func (o *MoveApmInstancesAdvancedParams) SetBody(body *models.TransientApmPlanConfiguration) {
	o.Body = body
}

// WithClusterID adds the clusterID to the move apm instances advanced params
func (o *MoveApmInstancesAdvancedParams) WithClusterID(clusterID string) *MoveApmInstancesAdvancedParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the move apm instances advanced params
func (o *MoveApmInstancesAdvancedParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithForceUpdate adds the forceUpdate to the move apm instances advanced params
func (o *MoveApmInstancesAdvancedParams) WithForceUpdate(forceUpdate *bool) *MoveApmInstancesAdvancedParams {
	o.SetForceUpdate(forceUpdate)
	return o
}

// SetForceUpdate adds the forceUpdate to the move apm instances advanced params
func (o *MoveApmInstancesAdvancedParams) SetForceUpdate(forceUpdate *bool) {
	o.ForceUpdate = forceUpdate
}

// WithValidateOnly adds the validateOnly to the move apm instances advanced params
func (o *MoveApmInstancesAdvancedParams) WithValidateOnly(validateOnly *bool) *MoveApmInstancesAdvancedParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the move apm instances advanced params
func (o *MoveApmInstancesAdvancedParams) SetValidateOnly(validateOnly *bool) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *MoveApmInstancesAdvancedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.ForceUpdate != nil {

		// query param force_update
		var qrForceUpdate bool
		if o.ForceUpdate != nil {
			qrForceUpdate = *o.ForceUpdate
		}
		qForceUpdate := swag.FormatBool(qrForceUpdate)
		if qForceUpdate != "" {
			if err := r.SetQueryParam("force_update", qForceUpdate); err != nil {
				return err
			}
		}

	}

	if o.ValidateOnly != nil {

		// query param validate_only
		var qrValidateOnly bool
		if o.ValidateOnly != nil {
			qrValidateOnly = *o.ValidateOnly
		}
		qValidateOnly := swag.FormatBool(qrValidateOnly)
		if qValidateOnly != "" {
			if err := r.SetQueryParam("validate_only", qValidateOnly); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
