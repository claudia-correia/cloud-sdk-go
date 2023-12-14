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

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetDeploymentEsResourceSnapshotDependencyReader is a Reader for the GetDeploymentEsResourceSnapshotDependency structure.
type GetDeploymentEsResourceSnapshotDependencyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentEsResourceSnapshotDependencyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentEsResourceSnapshotDependencyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDeploymentEsResourceSnapshotDependencyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDeploymentEsResourceSnapshotDependencyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentEsResourceSnapshotDependencyOK creates a GetDeploymentEsResourceSnapshotDependencyOK with default headers values
func NewGetDeploymentEsResourceSnapshotDependencyOK() *GetDeploymentEsResourceSnapshotDependencyOK {
	return &GetDeploymentEsResourceSnapshotDependencyOK{}
}

/*
GetDeploymentEsResourceSnapshotDependencyOK describes a response with status code 200, with default header values.

List of snapshot dependency
*/
type GetDeploymentEsResourceSnapshotDependencyOK struct {
	Payload *models.SnapshotDependencies
}

// IsSuccess returns true when this get deployment es resource snapshot dependency o k response has a 2xx status code
func (o *GetDeploymentEsResourceSnapshotDependencyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get deployment es resource snapshot dependency o k response has a 3xx status code
func (o *GetDeploymentEsResourceSnapshotDependencyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment es resource snapshot dependency o k response has a 4xx status code
func (o *GetDeploymentEsResourceSnapshotDependencyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get deployment es resource snapshot dependency o k response has a 5xx status code
func (o *GetDeploymentEsResourceSnapshotDependencyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment es resource snapshot dependency o k response a status code equal to that given
func (o *GetDeploymentEsResourceSnapshotDependencyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get deployment es resource snapshot dependency o k response
func (o *GetDeploymentEsResourceSnapshotDependencyOK) Code() int {
	return 200
}

func (o *GetDeploymentEsResourceSnapshotDependencyOK) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/elasticsearch/{ref_id}/snapshot/dependency][%d] getDeploymentEsResourceSnapshotDependencyOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentEsResourceSnapshotDependencyOK) String() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/elasticsearch/{ref_id}/snapshot/dependency][%d] getDeploymentEsResourceSnapshotDependencyOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentEsResourceSnapshotDependencyOK) GetPayload() *models.SnapshotDependencies {
	return o.Payload
}

func (o *GetDeploymentEsResourceSnapshotDependencyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotDependencies)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentEsResourceSnapshotDependencyNotFound creates a GetDeploymentEsResourceSnapshotDependencyNotFound with default headers values
func NewGetDeploymentEsResourceSnapshotDependencyNotFound() *GetDeploymentEsResourceSnapshotDependencyNotFound {
	return &GetDeploymentEsResourceSnapshotDependencyNotFound{}
}

/*
	GetDeploymentEsResourceSnapshotDependencyNotFound describes a response with status code 404, with default header values.

	* The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)

* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
*/
type GetDeploymentEsResourceSnapshotDependencyNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get deployment es resource snapshot dependency not found response has a 2xx status code
func (o *GetDeploymentEsResourceSnapshotDependencyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment es resource snapshot dependency not found response has a 3xx status code
func (o *GetDeploymentEsResourceSnapshotDependencyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment es resource snapshot dependency not found response has a 4xx status code
func (o *GetDeploymentEsResourceSnapshotDependencyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployment es resource snapshot dependency not found response has a 5xx status code
func (o *GetDeploymentEsResourceSnapshotDependencyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment es resource snapshot dependency not found response a status code equal to that given
func (o *GetDeploymentEsResourceSnapshotDependencyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get deployment es resource snapshot dependency not found response
func (o *GetDeploymentEsResourceSnapshotDependencyNotFound) Code() int {
	return 404
}

func (o *GetDeploymentEsResourceSnapshotDependencyNotFound) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/elasticsearch/{ref_id}/snapshot/dependency][%d] getDeploymentEsResourceSnapshotDependencyNotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentEsResourceSnapshotDependencyNotFound) String() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/elasticsearch/{ref_id}/snapshot/dependency][%d] getDeploymentEsResourceSnapshotDependencyNotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentEsResourceSnapshotDependencyNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentEsResourceSnapshotDependencyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentEsResourceSnapshotDependencyInternalServerError creates a GetDeploymentEsResourceSnapshotDependencyInternalServerError with default headers values
func NewGetDeploymentEsResourceSnapshotDependencyInternalServerError() *GetDeploymentEsResourceSnapshotDependencyInternalServerError {
	return &GetDeploymentEsResourceSnapshotDependencyInternalServerError{}
}

/*
GetDeploymentEsResourceSnapshotDependencyInternalServerError describes a response with status code 500, with default header values.

We have failed you. (code: `deployments.metadata_internal_error`)
*/
type GetDeploymentEsResourceSnapshotDependencyInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get deployment es resource snapshot dependency internal server error response has a 2xx status code
func (o *GetDeploymentEsResourceSnapshotDependencyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment es resource snapshot dependency internal server error response has a 3xx status code
func (o *GetDeploymentEsResourceSnapshotDependencyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment es resource snapshot dependency internal server error response has a 4xx status code
func (o *GetDeploymentEsResourceSnapshotDependencyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get deployment es resource snapshot dependency internal server error response has a 5xx status code
func (o *GetDeploymentEsResourceSnapshotDependencyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get deployment es resource snapshot dependency internal server error response a status code equal to that given
func (o *GetDeploymentEsResourceSnapshotDependencyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get deployment es resource snapshot dependency internal server error response
func (o *GetDeploymentEsResourceSnapshotDependencyInternalServerError) Code() int {
	return 500
}

func (o *GetDeploymentEsResourceSnapshotDependencyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/elasticsearch/{ref_id}/snapshot/dependency][%d] getDeploymentEsResourceSnapshotDependencyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDeploymentEsResourceSnapshotDependencyInternalServerError) String() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/elasticsearch/{ref_id}/snapshot/dependency][%d] getDeploymentEsResourceSnapshotDependencyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDeploymentEsResourceSnapshotDependencyInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentEsResourceSnapshotDependencyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
