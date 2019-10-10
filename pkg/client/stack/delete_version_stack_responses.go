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

package stack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// DeleteVersionStackReader is a Reader for the DeleteVersionStack structure.
type DeleteVersionStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVersionStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteVersionStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteVersionStackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 449:
		result := NewDeleteVersionStackRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteVersionStackOK creates a DeleteVersionStackOK with default headers values
func NewDeleteVersionStackOK() *DeleteVersionStackOK {
	return &DeleteVersionStackOK{}
}

/*DeleteVersionStackOK handles this case with default header values.

The `deleted` flag is applied to the specified Elastic Stack version.
*/
type DeleteVersionStackOK struct {
	Payload models.EmptyResponse
}

func (o *DeleteVersionStackOK) Error() string {
	return fmt.Sprintf("[DELETE /stack/versions/{version}][%d] deleteVersionStackOK  %+v", 200, o.Payload)
}

func (o *DeleteVersionStackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVersionStackNotFound creates a DeleteVersionStackNotFound with default headers values
func NewDeleteVersionStackNotFound() *DeleteVersionStackNotFound {
	return &DeleteVersionStackNotFound{}
}

/*DeleteVersionStackNotFound handles this case with default header values.

The Elastic Stack version can't be found. (code: `stackpack.version_not_found`)
*/
type DeleteVersionStackNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteVersionStackNotFound) Error() string {
	return fmt.Sprintf("[DELETE /stack/versions/{version}][%d] deleteVersionStackNotFound  %+v", 404, o.Payload)
}

func (o *DeleteVersionStackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVersionStackRetryWith creates a DeleteVersionStackRetryWith with default headers values
func NewDeleteVersionStackRetryWith() *DeleteVersionStackRetryWith {
	return &DeleteVersionStackRetryWith{}
}

/*DeleteVersionStackRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type DeleteVersionStackRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteVersionStackRetryWith) Error() string {
	return fmt.Sprintf("[DELETE /stack/versions/{version}][%d] deleteVersionStackRetryWith  %+v", 449, o.Payload)
}

func (o *DeleteVersionStackRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
