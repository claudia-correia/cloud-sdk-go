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
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AuthenticationInfo A user's authentication info
// swagger:model AuthenticationInfo
type AuthenticationInfo struct {

	// The UTC time when elevated permissions will expire, if the user has elevated permissions
	// Format: date-time
	ElevatedPermissionsExpireAt strfmt.DateTime `json:"elevated_permissions_expire_at,omitempty"`

	// True if the user has elevated permissions
	// Required: true
	HasElevatedPermissions *bool `json:"has_elevated_permissions"`

	// True if the user has an available TOTP device
	// Required: true
	HasTotpDevice *bool `json:"has_totp_device"`

	// The TOTP device source
	// Required: true
	// Enum: [native okta]
	TotpDeviceSource *string `json:"totp_device_source"`
}

// Validate validates this authentication info
func (m *AuthenticationInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElevatedPermissionsExpireAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasElevatedPermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasTotpDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotpDeviceSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthenticationInfo) validateElevatedPermissionsExpireAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ElevatedPermissionsExpireAt) { // not required
		return nil
	}

	if err := validate.FormatOf("elevated_permissions_expire_at", "body", "date-time", m.ElevatedPermissionsExpireAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuthenticationInfo) validateHasElevatedPermissions(formats strfmt.Registry) error {

	if err := validate.Required("has_elevated_permissions", "body", m.HasElevatedPermissions); err != nil {
		return err
	}

	return nil
}

func (m *AuthenticationInfo) validateHasTotpDevice(formats strfmt.Registry) error {

	if err := validate.Required("has_totp_device", "body", m.HasTotpDevice); err != nil {
		return err
	}

	return nil
}

var authenticationInfoTypeTotpDeviceSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["native","okta"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		authenticationInfoTypeTotpDeviceSourcePropEnum = append(authenticationInfoTypeTotpDeviceSourcePropEnum, v)
	}
}

const (

	// AuthenticationInfoTotpDeviceSourceNative captures enum value "native"
	AuthenticationInfoTotpDeviceSourceNative string = "native"

	// AuthenticationInfoTotpDeviceSourceOkta captures enum value "okta"
	AuthenticationInfoTotpDeviceSourceOkta string = "okta"
)

// prop value enum
func (m *AuthenticationInfo) validateTotpDeviceSourceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, authenticationInfoTypeTotpDeviceSourcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AuthenticationInfo) validateTotpDeviceSource(formats strfmt.Registry) error {

	if err := validate.Required("totp_device_source", "body", m.TotpDeviceSource); err != nil {
		return err
	}

	// value enum
	if err := m.validateTotpDeviceSourceEnum("totp_device_source", "body", *m.TotpDeviceSource); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthenticationInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthenticationInfo) UnmarshalBinary(b []byte) error {
	var res AuthenticationInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
