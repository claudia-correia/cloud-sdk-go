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

	"github.com/go-openapi/swag"
)

// RangeQuery The query that matches documents with fields that contain terms within a specified range.
// swagger:model RangeQuery
type RangeQuery struct {

	// An optional boost value to apply to the query.
	Boost float32 `json:"boost,omitempty"`

	// Formatted dates will be parsed using the format specified on the date field by default, but it can be overridden by passing the format parameter.
	Format string `json:"format,omitempty"`

	// Greater-than
	Gt interface{} `json:"gt,omitempty"`

	// Greater-than or equal to
	Gte interface{} `json:"gte,omitempty"`

	// Less-than
	Lt interface{} `json:"lt,omitempty"`

	// Less-than or equal to.
	Lte interface{} `json:"lte,omitempty"`

	// Dates can be converted from another timezone to UTC either by specifying the time zone in the date value itself (if the format accepts it), or it can be specified as the time_zone parameter.
	TimeZone string `json:"time_zone,omitempty"`
}

// Validate validates this range query
func (m *RangeQuery) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RangeQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RangeQuery) UnmarshalBinary(b []byte) error {
	var res RangeQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
