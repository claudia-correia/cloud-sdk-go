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

// InstanceConfiguration The configuration template for Elasticsearch instances, Kibana instances, and APM Servers.
// swagger:model InstanceConfiguration
type InstanceConfiguration struct {

	// Optional filter to match allocators against
	AllocatorFilter *QueryContainer `json:"allocator_filter,omitempty"`

	// Date/time that this instance configuration was marked for deletion
	// Format: date-time
	DeletedOn *strfmt.DateTime `json:"deleted_on,omitempty"`

	// Optional description for the instance configuration
	Description string `json:"description,omitempty"`

	// Numerics representing possible instance sizes that the instance configuration supports.
	// Required: true
	DiscreteSizes *DiscreteSizes `json:"discrete_sizes"`

	// Unique identifier for the instance configuration
	ID string `json:"id,omitempty"`

	// The type of instance (elasticsearch, kibana)
	// Required: true
	InstanceType *string `json:"instance_type"`

	// Optional arbitrary metadata to associate with this template.
	Metadata interface{} `json:"metadata,omitempty"`

	// Display name for the instance configuration.
	// Required: true
	Name *string `json:"name"`

	// Node types (master, data) for the instance
	NodeTypes []string `json:"node_types,omitempty"`

	// Settings for the instance storage multiplier
	StorageMultiplier float64 `json:"storage_multiplier,omitempty"`

	// Indicates if a instance configuration is system owned (restricts the set of operations that can be performed on it)
	SystemOwned *bool `json:"system_owned,omitempty"`
}

// Validate validates this instance configuration
func (m *InstanceConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocatorFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscreteSizes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceConfiguration) validateAllocatorFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.AllocatorFilter) { // not required
		return nil
	}

	if m.AllocatorFilter != nil {
		if err := m.AllocatorFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocator_filter")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceConfiguration) validateDeletedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.DeletedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("deleted_on", "body", "date-time", m.DeletedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InstanceConfiguration) validateDiscreteSizes(formats strfmt.Registry) error {

	if err := validate.Required("discrete_sizes", "body", m.DiscreteSizes); err != nil {
		return err
	}

	if m.DiscreteSizes != nil {
		if err := m.DiscreteSizes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discrete_sizes")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceConfiguration) validateInstanceType(formats strfmt.Registry) error {

	if err := validate.Required("instance_type", "body", m.InstanceType); err != nil {
		return err
	}

	return nil
}

func (m *InstanceConfiguration) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceConfiguration) UnmarshalBinary(b []byte) error {
	var res InstanceConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
