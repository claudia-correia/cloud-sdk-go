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

// QueryStringQuery A query that uses the `SimpleQueryParser` for parsing.
// swagger:model QueryStringQuery
type QueryStringQuery struct {

	// When set, * or ? are allowed as the first character. Defaults to false.
	AllowLeadingWildcard *bool `json:"allow_leading_wildcard,omitempty"`

	// The analyzer used to analyze each term of the query when creating composite queries.
	Analyzer string `json:"analyzer,omitempty"`

	// The default field for query terms if no prefix field is specified.
	DefaultField string `json:"default_field,omitempty"`

	// The default operator used if no explicit operator is specified.
	DefaultOperator string `json:"default_operator,omitempty"`

	// The actual query to be parsed.
	// Required: true
	Query *string `json:"query"`
}

// Validate validates this query string query
func (m *QueryStringQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryStringQuery) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryStringQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryStringQuery) UnmarshalBinary(b []byte) error {
	var res QueryStringQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
