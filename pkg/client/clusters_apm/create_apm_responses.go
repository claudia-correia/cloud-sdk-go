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

// CreateApmReader is a Reader for the CreateApm structure.
type CreateApmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateApmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateApmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewCreateApmCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateApmBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 449:
		result := NewCreateApmRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateApmOK creates a CreateApmOK with default headers values
func NewCreateApmOK() *CreateApmOK {
	return &CreateApmOK{}
}

/*CreateApmOK handles this case with default header values.

The APM server plan is valid. The return object contains an internal representation of the plan that you can use for debugging.
*/
type CreateApmOK struct {
	Payload *models.ApmCrudResponse
}

func (o *CreateApmOK) Error() string {
	return fmt.Sprintf("[POST /clusters/apm][%d] createApmOK  %+v", 200, o.Payload)
}

func (o *CreateApmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApmCrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateApmCreated creates a CreateApmCreated with default headers values
func NewCreateApmCreated() *CreateApmCreated {
	return &CreateApmCreated{}
}

/*CreateApmCreated handles this case with default header values.

The APM server plan is valid and the creation process has started.
*/
type CreateApmCreated struct {
	Payload *models.ApmCrudResponse
}

func (o *CreateApmCreated) Error() string {
	return fmt.Sprintf("[POST /clusters/apm][%d] createApmCreated  %+v", 201, o.Payload)
}

func (o *CreateApmCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApmCrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateApmBadRequest creates a CreateApmBadRequest with default headers values
func NewCreateApmBadRequest() *CreateApmBadRequest {
	return &CreateApmBadRequest{}
}

/*CreateApmBadRequest handles this case with default header values.

The APM server plan contains errors. (code: 'clusters.cluster_invalid_plan' and 'clusters.plan_feature_not_implemented')
*/
type CreateApmBadRequest struct {
	Payload *models.BasicFailedReply
}

func (o *CreateApmBadRequest) Error() string {
	return fmt.Sprintf("[POST /clusters/apm][%d] createApmBadRequest  %+v", 400, o.Payload)
}

func (o *CreateApmBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateApmRetryWith creates a CreateApmRetryWith with default headers values
func NewCreateApmRetryWith() *CreateApmRetryWith {
	return &CreateApmRetryWith{}
}

/*CreateApmRetryWith handles this case with default header values.

Elevated permissions are required. (code: '"root.needs_elevated_permissions"')
*/
type CreateApmRetryWith struct {
	Payload *models.BasicFailedReply
}

func (o *CreateApmRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/apm][%d] createApmRetryWith  %+v", 449, o.Payload)
}

func (o *CreateApmRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
