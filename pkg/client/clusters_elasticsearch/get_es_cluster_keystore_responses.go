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

// GetEsClusterKeystoreReader is a Reader for the GetEsClusterKeystore structure.
type GetEsClusterKeystoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEsClusterKeystoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEsClusterKeystoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetEsClusterKeystoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 449:
		result := NewGetEsClusterKeystoreRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEsClusterKeystoreOK creates a GetEsClusterKeystoreOK with default headers values
func NewGetEsClusterKeystoreOK() *GetEsClusterKeystoreOK {
	return &GetEsClusterKeystoreOK{}
}

/*GetEsClusterKeystoreOK handles this case with default header values.

The keystore settings are returned
*/
type GetEsClusterKeystoreOK struct {
	Payload *models.KeystoreContents
}

func (o *GetEsClusterKeystoreOK) Error() string {
	return fmt.Sprintf("[GET /clusters/elasticsearch/{cluster_id}/keystore][%d] getEsClusterKeystoreOK  %+v", 200, o.Payload)
}

func (o *GetEsClusterKeystoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeystoreContents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEsClusterKeystoreNotFound creates a GetEsClusterKeystoreNotFound with default headers values
func NewGetEsClusterKeystoreNotFound() *GetEsClusterKeystoreNotFound {
	return &GetEsClusterKeystoreNotFound{}
}

/*GetEsClusterKeystoreNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found (code: 'clusters.cluster_not_found')
*/
type GetEsClusterKeystoreNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *GetEsClusterKeystoreNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/elasticsearch/{cluster_id}/keystore][%d] getEsClusterKeystoreNotFound  %+v", 404, o.Payload)
}

func (o *GetEsClusterKeystoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEsClusterKeystoreRetryWith creates a GetEsClusterKeystoreRetryWith with default headers values
func NewGetEsClusterKeystoreRetryWith() *GetEsClusterKeystoreRetryWith {
	return &GetEsClusterKeystoreRetryWith{}
}

/*GetEsClusterKeystoreRetryWith handles this case with default header values.

Elevated permissions are required. (code: '"root.unauthorized.rbac.elevated_permissions_required"')
*/
type GetEsClusterKeystoreRetryWith struct {
	Payload *models.BasicFailedReply
}

func (o *GetEsClusterKeystoreRetryWith) Error() string {
	return fmt.Sprintf("[GET /clusters/elasticsearch/{cluster_id}/keystore][%d] getEsClusterKeystoreRetryWith  %+v", 449, o.Payload)
}

func (o *GetEsClusterKeystoreRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
