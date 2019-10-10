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
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ElasticsearchShardsInfo Information about the shards and replicas that comprise the Elasticsearch indices.
// swagger:model ElasticsearchShardsInfo
type ElasticsearchShardsInfo struct {

	// available shards
	// Required: true
	AvailableShards []*ElasticsearchShardElement `json:"available_shards"`

	// Whether the shard situation is healthy (any unavailable shards is unhealthy)
	// Required: true
	Healthy *bool `json:"healthy"`

	// unavailable replicas
	// Required: true
	UnavailableReplicas []*ElasticsearchReplicaElement `json:"unavailable_replicas"`

	// unavailable shards
	// Required: true
	UnavailableShards []*ElasticsearchShardElement `json:"unavailable_shards"`
}

// Validate validates this elasticsearch shards info
func (m *ElasticsearchShardsInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableShards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnavailableReplicas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnavailableShards(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchShardsInfo) validateAvailableShards(formats strfmt.Registry) error {

	if err := validate.Required("available_shards", "body", m.AvailableShards); err != nil {
		return err
	}

	for i := 0; i < len(m.AvailableShards); i++ {
		if swag.IsZero(m.AvailableShards[i]) { // not required
			continue
		}

		if m.AvailableShards[i] != nil {
			if err := m.AvailableShards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("available_shards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElasticsearchShardsInfo) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *ElasticsearchShardsInfo) validateUnavailableReplicas(formats strfmt.Registry) error {

	if err := validate.Required("unavailable_replicas", "body", m.UnavailableReplicas); err != nil {
		return err
	}

	for i := 0; i < len(m.UnavailableReplicas); i++ {
		if swag.IsZero(m.UnavailableReplicas[i]) { // not required
			continue
		}

		if m.UnavailableReplicas[i] != nil {
			if err := m.UnavailableReplicas[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("unavailable_replicas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElasticsearchShardsInfo) validateUnavailableShards(formats strfmt.Registry) error {

	if err := validate.Required("unavailable_shards", "body", m.UnavailableShards); err != nil {
		return err
	}

	for i := 0; i < len(m.UnavailableShards); i++ {
		if swag.IsZero(m.UnavailableShards[i]) { // not required
			continue
		}

		if m.UnavailableShards[i] != nil {
			if err := m.UnavailableShards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("unavailable_shards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchShardsInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchShardsInfo) UnmarshalBinary(b []byte) error {
	var res ElasticsearchShardsInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
