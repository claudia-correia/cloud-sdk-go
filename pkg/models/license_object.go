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

// LicenseObject The container for a license.
// swagger:model LicenseObject
type LicenseObject struct {

	// License data
	// Required: true
	License *LicenseInfo `json:"license"`

	// Information about current usage.
	// Read Only: true
	UsageStats *UsageStats `json:"usage_stats,omitempty"`
}

// Validate validates this license object
func (m *LicenseObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLicense(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsageStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LicenseObject) validateLicense(formats strfmt.Registry) error {

	if err := validate.Required("license", "body", m.License); err != nil {
		return err
	}

	if m.License != nil {
		if err := m.License.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("license")
			}
			return err
		}
	}

	return nil
}

func (m *LicenseObject) validateUsageStats(formats strfmt.Registry) error {

	if swag.IsZero(m.UsageStats) { // not required
		return nil
	}

	if m.UsageStats != nil {
		if err := m.UsageStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage_stats")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LicenseObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicenseObject) UnmarshalBinary(b []byte) error {
	var res LicenseObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
