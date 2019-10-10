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

package platform_configuration_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new platform configuration templates API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for platform configuration templates API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateDeploymentTemplate creates deployment template

Creates a deployment template.
*/
func (a *Client) CreateDeploymentTemplate(params *CreateDeploymentTemplateParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDeploymentTemplateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeploymentTemplateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create-deployment-template",
		Method:             "POST",
		PathPattern:        "/platform/configuration/templates/deployments",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDeploymentTemplateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDeploymentTemplateCreated), nil

}

/*
DeleteDeploymentTemplate deletes deployment template

Deletes a deployment template by id.
*/
func (a *Client) DeleteDeploymentTemplate(params *DeleteDeploymentTemplateParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDeploymentTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDeploymentTemplateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete-deployment-template",
		Method:             "DELETE",
		PathPattern:        "/platform/configuration/templates/deployments/{template_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDeploymentTemplateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDeploymentTemplateOK), nil

}

/*
GetDeploymentTemplate gets deployment template

Retrieves a deployment template by id.
*/
func (a *Client) GetDeploymentTemplate(params *GetDeploymentTemplateParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentTemplateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-deployment-template",
		Method:             "GET",
		PathPattern:        "/platform/configuration/templates/deployments/{template_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeploymentTemplateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDeploymentTemplateOK), nil

}

/*
GetDeploymentTemplates gets deployment templates

Retrieves all deployment templates.
*/
func (a *Client) GetDeploymentTemplates(params *GetDeploymentTemplatesParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeploymentTemplatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentTemplatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-deployment-templates",
		Method:             "GET",
		PathPattern:        "/platform/configuration/templates/deployments",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeploymentTemplatesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDeploymentTemplatesOK), nil

}

/*
SetDeploymentTemplate sets deployment template

Creates or updates a deployment template.
*/
func (a *Client) SetDeploymentTemplate(params *SetDeploymentTemplateParams, authInfo runtime.ClientAuthInfoWriter) (*SetDeploymentTemplateOK, *SetDeploymentTemplateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetDeploymentTemplateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "set-deployment-template",
		Method:             "PUT",
		PathPattern:        "/platform/configuration/templates/deployments/{template_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetDeploymentTemplateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *SetDeploymentTemplateOK:
		return value, nil, nil
	case *SetDeploymentTemplateCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
