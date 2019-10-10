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

package clusters_apm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// DeleteApmReader is a Reader for the DeleteApm structure.
type DeleteApmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteApmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteApmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteApmNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeleteApmPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 449:
		result := NewDeleteApmRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApmOK creates a DeleteApmOK with default headers values
func NewDeleteApmOK() *DeleteApmOK {
	return &DeleteApmOK{}
}

/*DeleteApmOK handles this case with default header values.

The APM server is deleted.
*/
type DeleteApmOK struct {
	Payload models.EmptyResponse
}

func (o *DeleteApmOK) Error() string {
	return fmt.Sprintf("[DELETE /clusters/apm/{cluster_id}][%d] deleteApmOK  %+v", 200, o.Payload)
}

func (o *DeleteApmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteApmNotFound creates a DeleteApmNotFound with default headers values
func NewDeleteApmNotFound() *DeleteApmNotFound {
	return &DeleteApmNotFound{}
}

/*DeleteApmNotFound handles this case with default header values.

The {cluster_id} can't be found. (code: 'clusters.cluster_not_found')
*/
type DeleteApmNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *DeleteApmNotFound) Error() string {
	return fmt.Sprintf("[DELETE /clusters/apm/{cluster_id}][%d] deleteApmNotFound  %+v", 404, o.Payload)
}

func (o *DeleteApmNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteApmPreconditionFailed creates a DeleteApmPreconditionFailed with default headers values
func NewDeleteApmPreconditionFailed() *DeleteApmPreconditionFailed {
	return &DeleteApmPreconditionFailed{}
}

/*DeleteApmPreconditionFailed handles this case with default header values.

To delete the APM server, you must first shut it down. (code: 'clusters.cluster_plan_state_error')
*/
type DeleteApmPreconditionFailed struct {
	Payload *models.BasicFailedReply
}

func (o *DeleteApmPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /clusters/apm/{cluster_id}][%d] deleteApmPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteApmPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteApmRetryWith creates a DeleteApmRetryWith with default headers values
func NewDeleteApmRetryWith() *DeleteApmRetryWith {
	return &DeleteApmRetryWith{}
}

/*DeleteApmRetryWith handles this case with default header values.

Elevated permissions are required. (code: '"root.needs_elevated_permissions"')
*/
type DeleteApmRetryWith struct {
	Payload *models.BasicFailedReply
}

func (o *DeleteApmRetryWith) Error() string {
	return fmt.Sprintf("[DELETE /clusters/apm/{cluster_id}][%d] deleteApmRetryWith  %+v", 449, o.Payload)
}

func (o *DeleteApmRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
