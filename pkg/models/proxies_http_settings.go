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

// ProxiesHTTPSettings The HTTP settings for all proxies.
// swagger:model ProxiesHttpSettings
type ProxiesHTTPSettings struct {

	// Secret string for the HTTP cookie
	// Required: true
	CookieSecret *string `json:"cookie_secret"`

	// Base URL for the dashboard
	// Required: true
	DashboardsBaseURL *string `json:"dashboards_base_url"`

	// Cutoff interval after disconnection in milliseconds
	// Required: true
	DisconnectedCutoff *int64 `json:"disconnected_cutoff"`

	// Minimum number of proxy instances
	// Required: true
	MinimumProxyServices *int32 `json:"minimum_proxy_services"`

	// Settings related to single-sign-on
	// Required: true
	SsoSettings *ProxiesSSOSettings `json:"sso_settings"`

	// User key for the HTTP cookie
	// Required: true
	UserCookieKey *string `json:"user_cookie_key"`
}

// Validate validates this proxies Http settings
func (m *ProxiesHTTPSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCookieSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDashboardsBaseURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisconnectedCutoff(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinimumProxyServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsoSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserCookieKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProxiesHTTPSettings) validateCookieSecret(formats strfmt.Registry) error {

	if err := validate.Required("cookie_secret", "body", m.CookieSecret); err != nil {
		return err
	}

	return nil
}

func (m *ProxiesHTTPSettings) validateDashboardsBaseURL(formats strfmt.Registry) error {

	if err := validate.Required("dashboards_base_url", "body", m.DashboardsBaseURL); err != nil {
		return err
	}

	return nil
}

func (m *ProxiesHTTPSettings) validateDisconnectedCutoff(formats strfmt.Registry) error {

	if err := validate.Required("disconnected_cutoff", "body", m.DisconnectedCutoff); err != nil {
		return err
	}

	return nil
}

func (m *ProxiesHTTPSettings) validateMinimumProxyServices(formats strfmt.Registry) error {

	if err := validate.Required("minimum_proxy_services", "body", m.MinimumProxyServices); err != nil {
		return err
	}

	return nil
}

func (m *ProxiesHTTPSettings) validateSsoSettings(formats strfmt.Registry) error {

	if err := validate.Required("sso_settings", "body", m.SsoSettings); err != nil {
		return err
	}

	if m.SsoSettings != nil {
		if err := m.SsoSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sso_settings")
			}
			return err
		}
	}

	return nil
}

func (m *ProxiesHTTPSettings) validateUserCookieKey(formats strfmt.Registry) error {

	if err := validate.Required("user_cookie_key", "body", m.UserCookieKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProxiesHTTPSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProxiesHTTPSettings) UnmarshalBinary(b []byte) error {
	var res ProxiesHTTPSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
