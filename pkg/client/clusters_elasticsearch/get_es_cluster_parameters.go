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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetEsClusterParams creates a new GetEsClusterParams object
// with the default values initialized.
func NewGetEsClusterParams() *GetEsClusterParams {
	var (
		convertLegacyPlansDefault = bool(false)
		enrichWithTemplateDefault = bool(false)
		showMetadataDefault       = bool(false)
		showPlanDefaultsDefault   = bool(false)
		showPlanLogsDefault       = bool(false)
		showPlansDefault          = bool(true)
		showSecurityDefault       = bool(false)
		showSettingsDefault       = bool(false)
		showSystemAlertsDefault   = int64(0)
	)
	return &GetEsClusterParams{
		ConvertLegacyPlans: &convertLegacyPlansDefault,
		EnrichWithTemplate: &enrichWithTemplateDefault,
		ShowMetadata:       &showMetadataDefault,
		ShowPlanDefaults:   &showPlanDefaultsDefault,
		ShowPlanLogs:       &showPlanLogsDefault,
		ShowPlans:          &showPlansDefault,
		ShowSecurity:       &showSecurityDefault,
		ShowSettings:       &showSettingsDefault,
		ShowSystemAlerts:   &showSystemAlertsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEsClusterParamsWithTimeout creates a new GetEsClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEsClusterParamsWithTimeout(timeout time.Duration) *GetEsClusterParams {
	var (
		convertLegacyPlansDefault = bool(false)
		enrichWithTemplateDefault = bool(false)
		showMetadataDefault       = bool(false)
		showPlanDefaultsDefault   = bool(false)
		showPlanLogsDefault       = bool(false)
		showPlansDefault          = bool(true)
		showSecurityDefault       = bool(false)
		showSettingsDefault       = bool(false)
		showSystemAlertsDefault   = int64(0)
	)
	return &GetEsClusterParams{
		ConvertLegacyPlans: &convertLegacyPlansDefault,
		EnrichWithTemplate: &enrichWithTemplateDefault,
		ShowMetadata:       &showMetadataDefault,
		ShowPlanDefaults:   &showPlanDefaultsDefault,
		ShowPlanLogs:       &showPlanLogsDefault,
		ShowPlans:          &showPlansDefault,
		ShowSecurity:       &showSecurityDefault,
		ShowSettings:       &showSettingsDefault,
		ShowSystemAlerts:   &showSystemAlertsDefault,

		timeout: timeout,
	}
}

// NewGetEsClusterParamsWithContext creates a new GetEsClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEsClusterParamsWithContext(ctx context.Context) *GetEsClusterParams {
	var (
		convertLegacyPlansDefault = bool(false)
		enrichWithTemplateDefault = bool(false)
		showMetadataDefault       = bool(false)
		showPlanDefaultsDefault   = bool(false)
		showPlanLogsDefault       = bool(false)
		showPlansDefault          = bool(true)
		showSecurityDefault       = bool(false)
		showSettingsDefault       = bool(false)
		showSystemAlertsDefault   = int64(0)
	)
	return &GetEsClusterParams{
		ConvertLegacyPlans: &convertLegacyPlansDefault,
		EnrichWithTemplate: &enrichWithTemplateDefault,
		ShowMetadata:       &showMetadataDefault,
		ShowPlanDefaults:   &showPlanDefaultsDefault,
		ShowPlanLogs:       &showPlanLogsDefault,
		ShowPlans:          &showPlansDefault,
		ShowSecurity:       &showSecurityDefault,
		ShowSettings:       &showSettingsDefault,
		ShowSystemAlerts:   &showSystemAlertsDefault,

		Context: ctx,
	}
}

// NewGetEsClusterParamsWithHTTPClient creates a new GetEsClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEsClusterParamsWithHTTPClient(client *http.Client) *GetEsClusterParams {
	var (
		convertLegacyPlansDefault = bool(false)
		enrichWithTemplateDefault = bool(false)
		showMetadataDefault       = bool(false)
		showPlanDefaultsDefault   = bool(false)
		showPlanLogsDefault       = bool(false)
		showPlansDefault          = bool(true)
		showSecurityDefault       = bool(false)
		showSettingsDefault       = bool(false)
		showSystemAlertsDefault   = int64(0)
	)
	return &GetEsClusterParams{
		ConvertLegacyPlans: &convertLegacyPlansDefault,
		EnrichWithTemplate: &enrichWithTemplateDefault,
		ShowMetadata:       &showMetadataDefault,
		ShowPlanDefaults:   &showPlanDefaultsDefault,
		ShowPlanLogs:       &showPlanLogsDefault,
		ShowPlans:          &showPlansDefault,
		ShowSecurity:       &showSecurityDefault,
		ShowSettings:       &showSettingsDefault,
		ShowSystemAlerts:   &showSystemAlertsDefault,
		HTTPClient:         client,
	}
}

/*GetEsClusterParams contains all the parameters to send to the API endpoint
for the get es cluster operation typically these are written to a http.Request
*/
type GetEsClusterParams struct {

	/*ClusterID
	  The Elasticsearch cluster identifier.

	*/
	ClusterID string
	/*ConvertLegacyPlans
	  When `true`, converts the plans to the 2.0.x format. When `false`, uses the 1.x format. The default is `false`.

	*/
	ConvertLegacyPlans *bool
	/*EnrichWithTemplate
	  When plans are shown, includes the missing elements from the applicable deployment template.

	*/
	EnrichWithTemplate *bool
	/*ShowMetadata
	  Includes all of the cluster metadata in the response. NOTE: Responses can include a large amount of metadata, as well as credentials.

	*/
	ShowMetadata *bool
	/*ShowPlanDefaults
	  When plans are shown, includes the default values in the response. NOTE: This option results in large responses.

	*/
	ShowPlanDefaults *bool
	/*ShowPlanLogs
	  Includes the active, pending, and historical plan information in the attempt log. NOTE: This option can result in large responses.

	*/
	ShowPlanLogs *bool
	/*ShowPlans
	  Includes the active and pending plan information in the response. NOTE: This option can result in large responses.

	*/
	ShowPlans *bool
	/*ShowSecurity
	  Includes the Elasticsearch 2.x security information in the response. NOTE: Responses can include a large amount of metadata, as well as credentials.

	*/
	ShowSecurity *bool
	/*ShowSettings
	  Includes the cluster settings in the response.

	*/
	ShowSettings *bool
	/*ShowSystemAlerts
	  The number of system alerts to include in the response. For example, the number of forced restarts caused from a limited amount of memory. Only numbers greater than zero return a field. NOTE: Responses can include a large number of system alerts.

	*/
	ShowSystemAlerts *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get es cluster params
func (o *GetEsClusterParams) WithTimeout(timeout time.Duration) *GetEsClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get es cluster params
func (o *GetEsClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get es cluster params
func (o *GetEsClusterParams) WithContext(ctx context.Context) *GetEsClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get es cluster params
func (o *GetEsClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get es cluster params
func (o *GetEsClusterParams) WithHTTPClient(client *http.Client) *GetEsClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get es cluster params
func (o *GetEsClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get es cluster params
func (o *GetEsClusterParams) WithClusterID(clusterID string) *GetEsClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get es cluster params
func (o *GetEsClusterParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithConvertLegacyPlans adds the convertLegacyPlans to the get es cluster params
func (o *GetEsClusterParams) WithConvertLegacyPlans(convertLegacyPlans *bool) *GetEsClusterParams {
	o.SetConvertLegacyPlans(convertLegacyPlans)
	return o
}

// SetConvertLegacyPlans adds the convertLegacyPlans to the get es cluster params
func (o *GetEsClusterParams) SetConvertLegacyPlans(convertLegacyPlans *bool) {
	o.ConvertLegacyPlans = convertLegacyPlans
}

// WithEnrichWithTemplate adds the enrichWithTemplate to the get es cluster params
func (o *GetEsClusterParams) WithEnrichWithTemplate(enrichWithTemplate *bool) *GetEsClusterParams {
	o.SetEnrichWithTemplate(enrichWithTemplate)
	return o
}

// SetEnrichWithTemplate adds the enrichWithTemplate to the get es cluster params
func (o *GetEsClusterParams) SetEnrichWithTemplate(enrichWithTemplate *bool) {
	o.EnrichWithTemplate = enrichWithTemplate
}

// WithShowMetadata adds the showMetadata to the get es cluster params
func (o *GetEsClusterParams) WithShowMetadata(showMetadata *bool) *GetEsClusterParams {
	o.SetShowMetadata(showMetadata)
	return o
}

// SetShowMetadata adds the showMetadata to the get es cluster params
func (o *GetEsClusterParams) SetShowMetadata(showMetadata *bool) {
	o.ShowMetadata = showMetadata
}

// WithShowPlanDefaults adds the showPlanDefaults to the get es cluster params
func (o *GetEsClusterParams) WithShowPlanDefaults(showPlanDefaults *bool) *GetEsClusterParams {
	o.SetShowPlanDefaults(showPlanDefaults)
	return o
}

// SetShowPlanDefaults adds the showPlanDefaults to the get es cluster params
func (o *GetEsClusterParams) SetShowPlanDefaults(showPlanDefaults *bool) {
	o.ShowPlanDefaults = showPlanDefaults
}

// WithShowPlanLogs adds the showPlanLogs to the get es cluster params
func (o *GetEsClusterParams) WithShowPlanLogs(showPlanLogs *bool) *GetEsClusterParams {
	o.SetShowPlanLogs(showPlanLogs)
	return o
}

// SetShowPlanLogs adds the showPlanLogs to the get es cluster params
func (o *GetEsClusterParams) SetShowPlanLogs(showPlanLogs *bool) {
	o.ShowPlanLogs = showPlanLogs
}

// WithShowPlans adds the showPlans to the get es cluster params
func (o *GetEsClusterParams) WithShowPlans(showPlans *bool) *GetEsClusterParams {
	o.SetShowPlans(showPlans)
	return o
}

// SetShowPlans adds the showPlans to the get es cluster params
func (o *GetEsClusterParams) SetShowPlans(showPlans *bool) {
	o.ShowPlans = showPlans
}

// WithShowSecurity adds the showSecurity to the get es cluster params
func (o *GetEsClusterParams) WithShowSecurity(showSecurity *bool) *GetEsClusterParams {
	o.SetShowSecurity(showSecurity)
	return o
}

// SetShowSecurity adds the showSecurity to the get es cluster params
func (o *GetEsClusterParams) SetShowSecurity(showSecurity *bool) {
	o.ShowSecurity = showSecurity
}

// WithShowSettings adds the showSettings to the get es cluster params
func (o *GetEsClusterParams) WithShowSettings(showSettings *bool) *GetEsClusterParams {
	o.SetShowSettings(showSettings)
	return o
}

// SetShowSettings adds the showSettings to the get es cluster params
func (o *GetEsClusterParams) SetShowSettings(showSettings *bool) {
	o.ShowSettings = showSettings
}

// WithShowSystemAlerts adds the showSystemAlerts to the get es cluster params
func (o *GetEsClusterParams) WithShowSystemAlerts(showSystemAlerts *int64) *GetEsClusterParams {
	o.SetShowSystemAlerts(showSystemAlerts)
	return o
}

// SetShowSystemAlerts adds the showSystemAlerts to the get es cluster params
func (o *GetEsClusterParams) SetShowSystemAlerts(showSystemAlerts *int64) {
	o.ShowSystemAlerts = showSystemAlerts
}

// WriteToRequest writes these params to a swagger request
func (o *GetEsClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.ConvertLegacyPlans != nil {

		// query param convert_legacy_plans
		var qrConvertLegacyPlans bool
		if o.ConvertLegacyPlans != nil {
			qrConvertLegacyPlans = *o.ConvertLegacyPlans
		}
		qConvertLegacyPlans := swag.FormatBool(qrConvertLegacyPlans)
		if qConvertLegacyPlans != "" {
			if err := r.SetQueryParam("convert_legacy_plans", qConvertLegacyPlans); err != nil {
				return err
			}
		}

	}

	if o.EnrichWithTemplate != nil {

		// query param enrich_with_template
		var qrEnrichWithTemplate bool
		if o.EnrichWithTemplate != nil {
			qrEnrichWithTemplate = *o.EnrichWithTemplate
		}
		qEnrichWithTemplate := swag.FormatBool(qrEnrichWithTemplate)
		if qEnrichWithTemplate != "" {
			if err := r.SetQueryParam("enrich_with_template", qEnrichWithTemplate); err != nil {
				return err
			}
		}

	}

	if o.ShowMetadata != nil {

		// query param show_metadata
		var qrShowMetadata bool
		if o.ShowMetadata != nil {
			qrShowMetadata = *o.ShowMetadata
		}
		qShowMetadata := swag.FormatBool(qrShowMetadata)
		if qShowMetadata != "" {
			if err := r.SetQueryParam("show_metadata", qShowMetadata); err != nil {
				return err
			}
		}

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

	if o.ShowPlans != nil {

		// query param show_plans
		var qrShowPlans bool
		if o.ShowPlans != nil {
			qrShowPlans = *o.ShowPlans
		}
		qShowPlans := swag.FormatBool(qrShowPlans)
		if qShowPlans != "" {
			if err := r.SetQueryParam("show_plans", qShowPlans); err != nil {
				return err
			}
		}

	}

	if o.ShowSecurity != nil {

		// query param show_security
		var qrShowSecurity bool
		if o.ShowSecurity != nil {
			qrShowSecurity = *o.ShowSecurity
		}
		qShowSecurity := swag.FormatBool(qrShowSecurity)
		if qShowSecurity != "" {
			if err := r.SetQueryParam("show_security", qShowSecurity); err != nil {
				return err
			}
		}

	}

	if o.ShowSettings != nil {

		// query param show_settings
		var qrShowSettings bool
		if o.ShowSettings != nil {
			qrShowSettings = *o.ShowSettings
		}
		qShowSettings := swag.FormatBool(qrShowSettings)
		if qShowSettings != "" {
			if err := r.SetQueryParam("show_settings", qShowSettings); err != nil {
				return err
			}
		}

	}

	if o.ShowSystemAlerts != nil {

		// query param show_system_alerts
		var qrShowSystemAlerts int64
		if o.ShowSystemAlerts != nil {
			qrShowSystemAlerts = *o.ShowSystemAlerts
		}
		qShowSystemAlerts := swag.FormatInt64(qrShowSystemAlerts)
		if qShowSystemAlerts != "" {
			if err := r.SetQueryParam("show_system_alerts", qShowSystemAlerts); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
