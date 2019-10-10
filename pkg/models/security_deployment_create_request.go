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
)

// SecurityDeploymentCreateRequest The creation request for the Elasticsearch security deployment.
// swagger:model SecurityDeploymentCreateRequest
type SecurityDeploymentCreateRequest struct {

	// The name of the security deployment cluster
	Name string `json:"name,omitempty"`

	// The cluster topology used to create the security deployment cluster
	Topology *SecurityDeploymentTopology `json:"topology,omitempty"`

	// The version of the Elasticsearch cluster
	Version string `json:"version,omitempty"`
}

// Validate validates this security deployment create request
func (m *SecurityDeploymentCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTopology(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityDeploymentCreateRequest) validateTopology(formats strfmt.Registry) error {

	if swag.IsZero(m.Topology) { // not required
		return nil
	}

	if m.Topology != nil {
		if err := m.Topology.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("topology")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityDeploymentCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityDeploymentCreateRequest) UnmarshalBinary(b []byte) error {
	var res SecurityDeploymentCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
