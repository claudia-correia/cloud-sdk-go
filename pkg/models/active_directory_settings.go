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

// ActiveDirectorySettings Elasticsearch Security Active Directory realm configuration
// swagger:model ActiveDirectorySettings
type ActiveDirectorySettings struct {

	// When true, bindDb credentials are ignored
	// Required: true
	BindAnonymously *bool `json:"bind_anonymously"`

	// The distinguished name of the user that is used to bind to the Active Directory and perform searches.
	BindDn string `json:"bind_dn,omitempty"`

	// The user password that is used to bind to the Active Directory server.
	BindPassword string `json:"bind_password,omitempty"`

	// The SSL trusted CA certificate bundle URL. The bundle should be a zip file containing a single keystore file 'keystore.ks' in the directory '/active_directory/:id/truststore', where :id is the value of the [id] field.
	CertificateURL string `json:"certificate_url,omitempty"`

	// The password to the certificate bundle URL truststore
	CertificateURLTruststorePassword string `json:"certificate_url_truststore_password,omitempty"`

	// The format of the keystore file. Should be jks to use the Java Keystore format or PKCS12 to use PKCS#12 files.  The default is jks.
	// Enum: [jks PKCS12]
	CertificateURLTruststoreType string `json:"certificate_url_truststore_type,omitempty"`

	// Specifies the domain name of the Active Directory (the forest root domain name).
	// Required: true
	DomainName *string `json:"domain_name"`

	// When true, enables the security realm
	Enabled *bool `json:"enabled,omitempty"`

	// The Active Directory group search configuration
	GroupSearch *ActiveDirectoryGroupSearch `json:"group_search,omitempty"`

	// The identifier for the security realm
	// Required: true
	ID *string `json:"id"`

	// The Active Directory load balancing behavior
	LoadBalance *ActiveDirectorySecurityRealmLoadBalance `json:"load_balance,omitempty"`

	// The friendly name of the security realm
	// Required: true
	Name *string `json:"name"`

	// The order that the security realm is evaluated
	Order int32 `json:"order,omitempty"`

	// Advanced configuration options in YAML format. Any settings defined here will override any configuration set via the API.
	OverrideYaml string `json:"override_yaml,omitempty"`

	// The role mapping rules associated with the security realm
	RoleMappings *ActiveDirectorySecurityRealmRoleMappingRules `json:"role_mappings,omitempty"`

	// The Active Directory URLs used to authenticate against, in the format ldap[s]://server:port. Note that ldap and ldaps protocols cannot be mixed together.
	// Required: true
	Urls []string `json:"urls"`

	// The Active Directory user search configuration.
	UserSearch *ActiveDirectoryUserSearch `json:"user_search,omitempty"`
}

// Validate validates this active directory settings
func (m *ActiveDirectorySettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBindAnonymously(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificateURLTruststoreType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomainName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupSearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoadBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleMappings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUrls(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserSearch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveDirectorySettings) validateBindAnonymously(formats strfmt.Registry) error {

	if err := validate.Required("bind_anonymously", "body", m.BindAnonymously); err != nil {
		return err
	}

	return nil
}

var activeDirectorySettingsTypeCertificateURLTruststoreTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["jks","PKCS12"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		activeDirectorySettingsTypeCertificateURLTruststoreTypePropEnum = append(activeDirectorySettingsTypeCertificateURLTruststoreTypePropEnum, v)
	}
}

const (

	// ActiveDirectorySettingsCertificateURLTruststoreTypeJks captures enum value "jks"
	ActiveDirectorySettingsCertificateURLTruststoreTypeJks string = "jks"

	// ActiveDirectorySettingsCertificateURLTruststoreTypePKCS12 captures enum value "PKCS12"
	ActiveDirectorySettingsCertificateURLTruststoreTypePKCS12 string = "PKCS12"
)

// prop value enum
func (m *ActiveDirectorySettings) validateCertificateURLTruststoreTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, activeDirectorySettingsTypeCertificateURLTruststoreTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ActiveDirectorySettings) validateCertificateURLTruststoreType(formats strfmt.Registry) error {

	if swag.IsZero(m.CertificateURLTruststoreType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCertificateURLTruststoreTypeEnum("certificate_url_truststore_type", "body", m.CertificateURLTruststoreType); err != nil {
		return err
	}

	return nil
}

func (m *ActiveDirectorySettings) validateDomainName(formats strfmt.Registry) error {

	if err := validate.Required("domain_name", "body", m.DomainName); err != nil {
		return err
	}

	return nil
}

func (m *ActiveDirectorySettings) validateGroupSearch(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupSearch) { // not required
		return nil
	}

	if m.GroupSearch != nil {
		if err := m.GroupSearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group_search")
			}
			return err
		}
	}

	return nil
}

func (m *ActiveDirectorySettings) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ActiveDirectorySettings) validateLoadBalance(formats strfmt.Registry) error {

	if swag.IsZero(m.LoadBalance) { // not required
		return nil
	}

	if m.LoadBalance != nil {
		if err := m.LoadBalance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("load_balance")
			}
			return err
		}
	}

	return nil
}

func (m *ActiveDirectorySettings) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ActiveDirectorySettings) validateRoleMappings(formats strfmt.Registry) error {

	if swag.IsZero(m.RoleMappings) { // not required
		return nil
	}

	if m.RoleMappings != nil {
		if err := m.RoleMappings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role_mappings")
			}
			return err
		}
	}

	return nil
}

func (m *ActiveDirectorySettings) validateUrls(formats strfmt.Registry) error {

	if err := validate.Required("urls", "body", m.Urls); err != nil {
		return err
	}

	return nil
}

func (m *ActiveDirectorySettings) validateUserSearch(formats strfmt.Registry) error {

	if swag.IsZero(m.UserSearch) { // not required
		return nil
	}

	if m.UserSearch != nil {
		if err := m.UserSearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user_search")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActiveDirectorySettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveDirectorySettings) UnmarshalBinary(b []byte) error {
	var res ActiveDirectorySettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
