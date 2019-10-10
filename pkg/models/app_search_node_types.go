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

// AppSearchNodeTypes Node types to enable for an AppSearch instance
// swagger:model AppSearchNodeTypes
type AppSearchNodeTypes struct {

	// Defines whether this instance should run as Application/API server
	// Required: true
	Appserver *bool `json:"appserver"`

	// Defines whether this instance should run as background worker
	// Required: true
	Worker *bool `json:"worker"`
}

// Validate validates this app search node types
func (m *AppSearchNodeTypes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppserver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorker(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppSearchNodeTypes) validateAppserver(formats strfmt.Registry) error {

	if err := validate.Required("appserver", "body", m.Appserver); err != nil {
		return err
	}

	return nil
}

func (m *AppSearchNodeTypes) validateWorker(formats strfmt.Registry) error {

	if err := validate.Required("worker", "body", m.Worker); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppSearchNodeTypes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppSearchNodeTypes) UnmarshalBinary(b []byte) error {
	var res AppSearchNodeTypes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
