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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ElasticsearchMonitoringInfo Information about the monitoring status for the Elasticsearch cluster.
// swagger:model ElasticsearchMonitoringInfo
type ElasticsearchMonitoringInfo struct {

	// The list of clusters Ids to which this cluster is currently sending monitoring data
	// Required: true
	DestinationClusterIds []string `json:"destination_cluster_ids"`

	// Whether the Monitoring configuration has been successfully applied
	// Required: true
	Healthy *bool `json:"healthy"`

	// The time the monitoring configuration was last changed
	// Required: true
	// Format: date-time
	LastModified *strfmt.DateTime `json:"last_modified"`

	// The status message from the last update (successful or not)
	// Required: true
	LastUpdateStatus *string `json:"last_update_status"`

	// The list of clusters Ids from which this cluster is currently receiving monitoring data
	// Required: true
	SourceClusterIds []string `json:"source_cluster_ids"`
}

// Validate validates this elasticsearch monitoring info
func (m *ElasticsearchMonitoringInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationClusterIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceClusterIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchMonitoringInfo) validateDestinationClusterIds(formats strfmt.Registry) error {

	if err := validate.Required("destination_cluster_ids", "body", m.DestinationClusterIds); err != nil {
		return err
	}

	return nil
}

func (m *ElasticsearchMonitoringInfo) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *ElasticsearchMonitoringInfo) validateLastModified(formats strfmt.Registry) error {

	if err := validate.Required("last_modified", "body", m.LastModified); err != nil {
		return err
	}

	if err := validate.FormatOf("last_modified", "body", "date-time", m.LastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ElasticsearchMonitoringInfo) validateLastUpdateStatus(formats strfmt.Registry) error {

	if err := validate.Required("last_update_status", "body", m.LastUpdateStatus); err != nil {
		return err
	}

	return nil
}

func (m *ElasticsearchMonitoringInfo) validateSourceClusterIds(formats strfmt.Registry) error {

	if err := validate.Required("source_cluster_ids", "body", m.SourceClusterIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchMonitoringInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchMonitoringInfo) UnmarshalBinary(b []byte) error {
	var res ElasticsearchMonitoringInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
