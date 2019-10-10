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

package platform_configuration_security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// UpdateLdapConfigurationReader is a Reader for the UpdateLdapConfiguration structure.
type UpdateLdapConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLdapConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateLdapConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateLdapConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateLdapConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewUpdateLdapConfigurationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 449:
		result := NewUpdateLdapConfigurationRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateLdapConfigurationOK creates a UpdateLdapConfigurationOK with default headers values
func NewUpdateLdapConfigurationOK() *UpdateLdapConfigurationOK {
	return &UpdateLdapConfigurationOK{}
}

/*UpdateLdapConfigurationOK handles this case with default header values.

The LDAP configuration was successfully updated
*/
type UpdateLdapConfigurationOK struct {
	/*The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string
	/*The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string
	/*The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload models.EmptyResponse
}

func (o *UpdateLdapConfigurationOK) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/realms/ldap/{realm_id}][%d] updateLdapConfigurationOK  %+v", 200, o.Payload)
}

func (o *UpdateLdapConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-resource-created
	o.XCloudResourceCreated = response.GetHeader("x-cloud-resource-created")

	// response header x-cloud-resource-last-modified
	o.XCloudResourceLastModified = response.GetHeader("x-cloud-resource-last-modified")

	// response header x-cloud-resource-version
	o.XCloudResourceVersion = response.GetHeader("x-cloud-resource-version")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLdapConfigurationBadRequest creates a UpdateLdapConfigurationBadRequest with default headers values
func NewUpdateLdapConfigurationBadRequest() *UpdateLdapConfigurationBadRequest {
	return &UpdateLdapConfigurationBadRequest{}
}

/*UpdateLdapConfigurationBadRequest handles this case with default header values.

* The realm id is already in use. (code: `security_realm.id_conflict`)
* The selected id is not valid. (code: `security_realm.invalid_id`)
* Order must be greater than zero. (code: `security_realm.invalid_order`)
* Invalid Elasticsearch Security realm type. (code: `security_realm.invalid_type`)
* The realm order is already in use. (code: `security_realm.order_conflict`)
* Advanced YAML format is invalid. (code: `security_realm.invalid_yaml`)
* The url format is invalid. (code: `security_realm.invalid_url`)
* Invalid LDAP URL. (code: `security_realm.ldap.invalid_url`)
* Invalid certificate bundle URL. (code: `security_realm.invalid_bundle_url`)
 */
type UpdateLdapConfigurationBadRequest struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateLdapConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/realms/ldap/{realm_id}][%d] updateLdapConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateLdapConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLdapConfigurationNotFound creates a UpdateLdapConfigurationNotFound with default headers values
func NewUpdateLdapConfigurationNotFound() *UpdateLdapConfigurationNotFound {
	return &UpdateLdapConfigurationNotFound{}
}

/*UpdateLdapConfigurationNotFound handles this case with default header values.

The realm specified by {realm_id} cannot be found. (code: `security_realm.not_found`)
*/
type UpdateLdapConfigurationNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateLdapConfigurationNotFound) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/realms/ldap/{realm_id}][%d] updateLdapConfigurationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateLdapConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLdapConfigurationConflict creates a UpdateLdapConfigurationConflict with default headers values
func NewUpdateLdapConfigurationConflict() *UpdateLdapConfigurationConflict {
	return &UpdateLdapConfigurationConflict{}
}

/*UpdateLdapConfigurationConflict handles this case with default header values.

There is a version conflict. (code: `security_realm.version_conflict`)
*/
type UpdateLdapConfigurationConflict struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateLdapConfigurationConflict) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/realms/ldap/{realm_id}][%d] updateLdapConfigurationConflict  %+v", 409, o.Payload)
}

func (o *UpdateLdapConfigurationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLdapConfigurationRetryWith creates a UpdateLdapConfigurationRetryWith with default headers values
func NewUpdateLdapConfigurationRetryWith() *UpdateLdapConfigurationRetryWith {
	return &UpdateLdapConfigurationRetryWith{}
}

/*UpdateLdapConfigurationRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type UpdateLdapConfigurationRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateLdapConfigurationRetryWith) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/realms/ldap/{realm_id}][%d] updateLdapConfigurationRetryWith  %+v", 449, o.Payload)
}

func (o *UpdateLdapConfigurationRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
