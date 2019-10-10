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

// SetEsClusterCcsSettingsReader is a Reader for the SetEsClusterCcsSettings structure.
type SetEsClusterCcsSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetEsClusterCcsSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewSetEsClusterCcsSettingsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewSetEsClusterCcsSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetEsClusterCcsSettingsAccepted creates a SetEsClusterCcsSettingsAccepted with default headers values
func NewSetEsClusterCcsSettingsAccepted() *SetEsClusterCcsSettingsAccepted {
	return &SetEsClusterCcsSettingsAccepted{}
}

/*SetEsClusterCcsSettingsAccepted handles this case with default header values.

The configuration for remote clusters was updated
*/
type SetEsClusterCcsSettingsAccepted struct {
	Payload models.EmptyResponse
}

func (o *SetEsClusterCcsSettingsAccepted) Error() string {
	return fmt.Sprintf("[PUT /clusters/elasticsearch/{cluster_id}/ccs/settings][%d] setEsClusterCcsSettingsAccepted  %+v", 202, o.Payload)
}

func (o *SetEsClusterCcsSettingsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetEsClusterCcsSettingsNotFound creates a SetEsClusterCcsSettingsNotFound with default headers values
func NewSetEsClusterCcsSettingsNotFound() *SetEsClusterCcsSettingsNotFound {
	return &SetEsClusterCcsSettingsNotFound{}
}

/*SetEsClusterCcsSettingsNotFound handles this case with default header values.

The cluster specified by {cluster_id} or a remote cluster cannot be found (code: 'clusters.cluster_not_found')
*/
type SetEsClusterCcsSettingsNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *SetEsClusterCcsSettingsNotFound) Error() string {
	return fmt.Sprintf("[PUT /clusters/elasticsearch/{cluster_id}/ccs/settings][%d] setEsClusterCcsSettingsNotFound  %+v", 404, o.Payload)
}

func (o *SetEsClusterCcsSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
