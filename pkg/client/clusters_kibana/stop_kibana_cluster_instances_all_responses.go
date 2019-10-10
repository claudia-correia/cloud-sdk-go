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

package clusters_kibana

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// StopKibanaClusterInstancesAllReader is a Reader for the StopKibanaClusterInstancesAll structure.
type StopKibanaClusterInstancesAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopKibanaClusterInstancesAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewStopKibanaClusterInstancesAllAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewStopKibanaClusterInstancesAllForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewStopKibanaClusterInstancesAllNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 449:
		result := NewStopKibanaClusterInstancesAllRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStopKibanaClusterInstancesAllAccepted creates a StopKibanaClusterInstancesAllAccepted with default headers values
func NewStopKibanaClusterInstancesAllAccepted() *StopKibanaClusterInstancesAllAccepted {
	return &StopKibanaClusterInstancesAllAccepted{}
}

/*StopKibanaClusterInstancesAllAccepted handles this case with default header values.

The stop command was issued successfully, use the "GET" command on the /{cluster_id} resource to monitor progress
*/
type StopKibanaClusterInstancesAllAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *StopKibanaClusterInstancesAllAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/_stop][%d] stopKibanaClusterInstancesAllAccepted  %+v", 202, o.Payload)
}

func (o *StopKibanaClusterInstancesAllAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopKibanaClusterInstancesAllForbidden creates a StopKibanaClusterInstancesAllForbidden with default headers values
func NewStopKibanaClusterInstancesAllForbidden() *StopKibanaClusterInstancesAllForbidden {
	return &StopKibanaClusterInstancesAllForbidden{}
}

/*StopKibanaClusterInstancesAllForbidden handles this case with default header values.

The stop command was prohibited for the given cluster. (code: `clusters.command_prohibited`)
*/
type StopKibanaClusterInstancesAllForbidden struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopKibanaClusterInstancesAllForbidden) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/_stop][%d] stopKibanaClusterInstancesAllForbidden  %+v", 403, o.Payload)
}

func (o *StopKibanaClusterInstancesAllForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopKibanaClusterInstancesAllNotFound creates a StopKibanaClusterInstancesAllNotFound with default headers values
func NewStopKibanaClusterInstancesAllNotFound() *StopKibanaClusterInstancesAllNotFound {
	return &StopKibanaClusterInstancesAllNotFound{}
}

/*StopKibanaClusterInstancesAllNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type StopKibanaClusterInstancesAllNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopKibanaClusterInstancesAllNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/_stop][%d] stopKibanaClusterInstancesAllNotFound  %+v", 404, o.Payload)
}

func (o *StopKibanaClusterInstancesAllNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopKibanaClusterInstancesAllRetryWith creates a StopKibanaClusterInstancesAllRetryWith with default headers values
func NewStopKibanaClusterInstancesAllRetryWith() *StopKibanaClusterInstancesAllRetryWith {
	return &StopKibanaClusterInstancesAllRetryWith{}
}

/*StopKibanaClusterInstancesAllRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type StopKibanaClusterInstancesAllRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopKibanaClusterInstancesAllRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/_stop][%d] stopKibanaClusterInstancesAllRetryWith  %+v", 449, o.Payload)
}

func (o *StopKibanaClusterInstancesAllRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
