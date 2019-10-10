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

// ClusterMetadataCPUResourcesSettings Specifies the CPU resource settings for the Elasticsearch cluster.
// swagger:model ClusterMetadataCpuResourcesSettings
type ClusterMetadataCPUResourcesSettings struct {

	// Indicates if the the CPU boost flag for a cluster is enabled or not
	// Required: true
	Boost *bool `json:"boost"`

	// Indicates if the the CPU hard limit flag for a cluster is enabled or not
	// Required: true
	HardLimit *bool `json:"hard_limit"`
}

// Validate validates this cluster metadata Cpu resources settings
func (m *ClusterMetadataCPUResourcesSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHardLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterMetadataCPUResourcesSettings) validateBoost(formats strfmt.Registry) error {

	if err := validate.Required("boost", "body", m.Boost); err != nil {
		return err
	}

	return nil
}

func (m *ClusterMetadataCPUResourcesSettings) validateHardLimit(formats strfmt.Registry) error {

	if err := validate.Required("hard_limit", "body", m.HardLimit); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterMetadataCPUResourcesSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterMetadataCPUResourcesSettings) UnmarshalBinary(b []byte) error {
	var res ClusterMetadataCPUResourcesSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
