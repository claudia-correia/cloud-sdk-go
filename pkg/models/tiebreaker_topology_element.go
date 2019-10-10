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

// TiebreakerTopologyElement Defines the topology, such as capacity and location, of the special tiebreaker node. TIP: When there are master nodes present, the master node topology is the default.
// swagger:model TiebreakerTopologyElement
type TiebreakerTopologyElement struct {

	// The memory capacity in MB for the tiebreaker node.
	// Required: true
	MemoryPerNode *int32 `json:"memory_per_node"`
}

// Validate validates this tiebreaker topology element
func (m *TiebreakerTopologyElement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMemoryPerNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TiebreakerTopologyElement) validateMemoryPerNode(formats strfmt.Registry) error {

	if err := validate.Required("memory_per_node", "body", m.MemoryPerNode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TiebreakerTopologyElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TiebreakerTopologyElement) UnmarshalBinary(b []byte) error {
	var res TiebreakerTopologyElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
