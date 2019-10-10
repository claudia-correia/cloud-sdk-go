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

package clusters_elasticsearch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// StopEsClusterInstancesAllMaintenanceModeReader is a Reader for the StopEsClusterInstancesAllMaintenanceMode structure.
type StopEsClusterInstancesAllMaintenanceModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopEsClusterInstancesAllMaintenanceModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewStopEsClusterInstancesAllMaintenanceModeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewStopEsClusterInstancesAllMaintenanceModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewStopEsClusterInstancesAllMaintenanceModeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 449:
		result := NewStopEsClusterInstancesAllMaintenanceModeRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStopEsClusterInstancesAllMaintenanceModeAccepted creates a StopEsClusterInstancesAllMaintenanceModeAccepted with default headers values
func NewStopEsClusterInstancesAllMaintenanceModeAccepted() *StopEsClusterInstancesAllMaintenanceModeAccepted {
	return &StopEsClusterInstancesAllMaintenanceModeAccepted{}
}

/*StopEsClusterInstancesAllMaintenanceModeAccepted handles this case with default header values.

The stop maintenance mode command was issued successfully, use the "GET" command on the /{cluster_id} resource to monitor progress
*/
type StopEsClusterInstancesAllMaintenanceModeAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *StopEsClusterInstancesAllMaintenanceModeAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/maintenance-mode/_stop][%d] stopEsClusterInstancesAllMaintenanceModeAccepted  %+v", 202, o.Payload)
}

func (o *StopEsClusterInstancesAllMaintenanceModeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopEsClusterInstancesAllMaintenanceModeForbidden creates a StopEsClusterInstancesAllMaintenanceModeForbidden with default headers values
func NewStopEsClusterInstancesAllMaintenanceModeForbidden() *StopEsClusterInstancesAllMaintenanceModeForbidden {
	return &StopEsClusterInstancesAllMaintenanceModeForbidden{}
}

/*StopEsClusterInstancesAllMaintenanceModeForbidden handles this case with default header values.

The stop maintenance mode command was prohibited for the given cluster. (code: `clusters.command_prohibited`)
*/
type StopEsClusterInstancesAllMaintenanceModeForbidden struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopEsClusterInstancesAllMaintenanceModeForbidden) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/maintenance-mode/_stop][%d] stopEsClusterInstancesAllMaintenanceModeForbidden  %+v", 403, o.Payload)
}

func (o *StopEsClusterInstancesAllMaintenanceModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopEsClusterInstancesAllMaintenanceModeNotFound creates a StopEsClusterInstancesAllMaintenanceModeNotFound with default headers values
func NewStopEsClusterInstancesAllMaintenanceModeNotFound() *StopEsClusterInstancesAllMaintenanceModeNotFound {
	return &StopEsClusterInstancesAllMaintenanceModeNotFound{}
}

/*StopEsClusterInstancesAllMaintenanceModeNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type StopEsClusterInstancesAllMaintenanceModeNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopEsClusterInstancesAllMaintenanceModeNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/maintenance-mode/_stop][%d] stopEsClusterInstancesAllMaintenanceModeNotFound  %+v", 404, o.Payload)
}

func (o *StopEsClusterInstancesAllMaintenanceModeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopEsClusterInstancesAllMaintenanceModeRetryWith creates a StopEsClusterInstancesAllMaintenanceModeRetryWith with default headers values
func NewStopEsClusterInstancesAllMaintenanceModeRetryWith() *StopEsClusterInstancesAllMaintenanceModeRetryWith {
	return &StopEsClusterInstancesAllMaintenanceModeRetryWith{}
}

/*StopEsClusterInstancesAllMaintenanceModeRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type StopEsClusterInstancesAllMaintenanceModeRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopEsClusterInstancesAllMaintenanceModeRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/maintenance-mode/_stop][%d] stopEsClusterInstancesAllMaintenanceModeRetryWith  %+v", 449, o.Payload)
}

func (o *StopEsClusterInstancesAllMaintenanceModeRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
