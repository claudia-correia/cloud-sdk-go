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
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new clusters kibana API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for clusters kibana API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CancelKibanaClusterPendingPlan cancels pending plan

Cancels the pending plan of the Kibana instance.
*/
func (a *Client) CancelKibanaClusterPendingPlan(params *CancelKibanaClusterPendingPlanParams, authInfo runtime.ClientAuthInfoWriter) (*CancelKibanaClusterPendingPlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelKibanaClusterPendingPlanParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cancel-kibana-cluster-pending-plan",
		Method:             "DELETE",
		PathPattern:        "/clusters/kibana/{cluster_id}/plan/pending",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CancelKibanaClusterPendingPlanReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CancelKibanaClusterPendingPlanOK), nil

}

/*
CreateKibanaCluster creates cluster

Creates a Kibana instance for the Elasticsearch cluster.
*/
func (a *Client) CreateKibanaCluster(params *CreateKibanaClusterParams, authInfo runtime.ClientAuthInfoWriter) (*CreateKibanaClusterOK, *CreateKibanaClusterCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateKibanaClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create-kibana-cluster",
		Method:             "POST",
		PathPattern:        "/clusters/kibana",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateKibanaClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateKibanaClusterOK:
		return value, nil, nil
	case *CreateKibanaClusterCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
DeleteKibProxyRequests proxies HTTP d e l e t e request

Proxies the HTTP DELETE request to the cluster. You must specify the `X-Management-Request` HTTP header. NOTE: Use this endpoint for management purposes. It does not provide high performance.
*/
func (a *Client) DeleteKibProxyRequests(params *DeleteKibProxyRequestsParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteKibProxyRequestsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteKibProxyRequestsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete-kib-proxy-requests",
		Method:             "DELETE",
		PathPattern:        "/clusters/kibana/{cluster_id}/proxy/{kibana_path}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteKibProxyRequestsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteKibProxyRequestsOK), nil

}

/*
DeleteKibanaCluster deletes cluster

Deletes the Kibana instance.
Before you delete the Kibana instance, you must first successfully issue a `_shutdown` command.
*/
func (a *Client) DeleteKibanaCluster(params *DeleteKibanaClusterParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteKibanaClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteKibanaClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete-kibana-cluster",
		Method:             "DELETE",
		PathPattern:        "/clusters/kibana/{cluster_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteKibanaClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteKibanaClusterOK), nil

}

/*
GetKibProxyRequests proxies HTTP g e t request

Proxies the HTTP GET request to the cluster. You must specify the `X-Management-Request` HTTP header. NOTE: Use this endpoint for management purposes. It does not provide high performance.
*/
func (a *Client) GetKibProxyRequests(params *GetKibProxyRequestsParams, authInfo runtime.ClientAuthInfoWriter) (*GetKibProxyRequestsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKibProxyRequestsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-kib-proxy-requests",
		Method:             "GET",
		PathPattern:        "/clusters/kibana/{cluster_id}/proxy/{kibana_path}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKibProxyRequestsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetKibProxyRequestsOK), nil

}

/*
GetKibanaCluster gets cluster

Retrieves the Kibana instance information.
*/
func (a *Client) GetKibanaCluster(params *GetKibanaClusterParams, authInfo runtime.ClientAuthInfoWriter) (*GetKibanaClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKibanaClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-kibana-cluster",
		Method:             "GET",
		PathPattern:        "/clusters/kibana/{cluster_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKibanaClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetKibanaClusterOK), nil

}

/*
GetKibanaClusterMetadataRaw gets cluster metadata

Advanced use only. Retrieves the internal metadata, in free-form JSON, for the Kibana instance.
*/
func (a *Client) GetKibanaClusterMetadataRaw(params *GetKibanaClusterMetadataRawParams, authInfo runtime.ClientAuthInfoWriter) (*GetKibanaClusterMetadataRawOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKibanaClusterMetadataRawParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-kibana-cluster-metadata-raw",
		Method:             "GET",
		PathPattern:        "/clusters/kibana/{cluster_id}/metadata/raw",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKibanaClusterMetadataRawReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetKibanaClusterMetadataRawOK), nil

}

/*
GetKibanaClusterMetadataSettings gets cluster metadata settings

Retrieves a structured version of the cluster metadata as a collection of top-level settings. If a particular setting isn't returned, then the free-form JSON endpoint (`/metadata/raw`) must be used.
*/
func (a *Client) GetKibanaClusterMetadataSettings(params *GetKibanaClusterMetadataSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetKibanaClusterMetadataSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKibanaClusterMetadataSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-kibana-cluster-metadata-settings",
		Method:             "GET",
		PathPattern:        "/clusters/kibana/{cluster_id}/metadata/settings",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKibanaClusterMetadataSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetKibanaClusterMetadataSettingsOK), nil

}

/*
GetKibanaClusterPendingPlan gets pending plan

Retrieves the pending plan of the Kibana instance.
*/
func (a *Client) GetKibanaClusterPendingPlan(params *GetKibanaClusterPendingPlanParams, authInfo runtime.ClientAuthInfoWriter) (*GetKibanaClusterPendingPlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKibanaClusterPendingPlanParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-kibana-cluster-pending-plan",
		Method:             "GET",
		PathPattern:        "/clusters/kibana/{cluster_id}/plan/pending",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKibanaClusterPendingPlanReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetKibanaClusterPendingPlanOK), nil

}

/*
GetKibanaClusterPlan gets plan

Retrieves the active plan of the Kibana instance. Transient settings are not show by this endpoint. To view the transient settings that have been applied with a specific plan, use the activity endpoint.
*/
func (a *Client) GetKibanaClusterPlan(params *GetKibanaClusterPlanParams, authInfo runtime.ClientAuthInfoWriter) (*GetKibanaClusterPlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKibanaClusterPlanParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-kibana-cluster-plan",
		Method:             "GET",
		PathPattern:        "/clusters/kibana/{cluster_id}/plan",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKibanaClusterPlanReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetKibanaClusterPlanOK), nil

}

/*
GetKibanaClusterPlanActivity gets plan activity

Retrieves the active and historical plan information for the Kibana instance.
*/
func (a *Client) GetKibanaClusterPlanActivity(params *GetKibanaClusterPlanActivityParams, authInfo runtime.ClientAuthInfoWriter) (*GetKibanaClusterPlanActivityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKibanaClusterPlanActivityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-kibana-cluster-plan-activity",
		Method:             "GET",
		PathPattern:        "/clusters/kibana/{cluster_id}/plan/activity",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKibanaClusterPlanActivityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetKibanaClusterPlanActivityOK), nil

}

/*
GetKibanaClusters gets clusters

Retrieves the information for all Kibana instances.
*/
func (a *Client) GetKibanaClusters(params *GetKibanaClustersParams, authInfo runtime.ClientAuthInfoWriter) (*GetKibanaClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKibanaClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-kibana-clusters",
		Method:             "GET",
		PathPattern:        "/clusters/kibana",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKibanaClustersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetKibanaClustersOK), nil

}

/*
MoveKibanaClusterInstances moves instances

Moves one or more Kibana instances.
*/
func (a *Client) MoveKibanaClusterInstances(params *MoveKibanaClusterInstancesParams, authInfo runtime.ClientAuthInfoWriter) (*MoveKibanaClusterInstancesAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMoveKibanaClusterInstancesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "move-kibana-cluster-instances",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/{cluster_id}/instances/{instance_ids}/_move",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MoveKibanaClusterInstancesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MoveKibanaClusterInstancesAccepted), nil

}

/*
MoveKibanaClusterInstancesAdvanced moves instances advanced

Moves one or more Kibana instances. The custom configuration settings are posted in the body.
*/
func (a *Client) MoveKibanaClusterInstancesAdvanced(params *MoveKibanaClusterInstancesAdvancedParams, authInfo runtime.ClientAuthInfoWriter) (*MoveKibanaClusterInstancesAdvancedAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMoveKibanaClusterInstancesAdvancedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "move-kibana-cluster-instances-advanced",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/{cluster_id}/instances/_move",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MoveKibanaClusterInstancesAdvancedReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MoveKibanaClusterInstancesAdvancedAccepted), nil

}

/*
PostKibProxyRequests proxies HTTP p o s t request

Proxies the HTTP POST request to the cluster. You must specify the `X-Management-Request` HTTP header. NOTE: Use this endpoint for management purposes. It does not provide high performance.
*/
func (a *Client) PostKibProxyRequests(params *PostKibProxyRequestsParams, authInfo runtime.ClientAuthInfoWriter) (*PostKibProxyRequestsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostKibProxyRequestsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "post-kib-proxy-requests",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/{cluster_id}/proxy/{kibana_path}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostKibProxyRequestsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostKibProxyRequestsOK), nil

}

/*
PutKibProxyRequests proxies HTTP p u t request

Proxies the HTTP PUT request to the cluster. You must specify the `X-Management-Request` HTTP header. NOTE: Use this endpoint for management purposes. It does not provide high performance.
*/
func (a *Client) PutKibProxyRequests(params *PutKibProxyRequestsParams, authInfo runtime.ClientAuthInfoWriter) (*PutKibProxyRequestsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutKibProxyRequestsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "put-kib-proxy-requests",
		Method:             "PUT",
		PathPattern:        "/clusters/kibana/{cluster_id}/proxy/{kibana_path}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutKibProxyRequestsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutKibProxyRequestsOK), nil

}

/*
RestartKibanaCluster restarts cluster

Restarts the Kibana instance. When you restart an active instance, the existing plan is used and a `cluster_reboot` is applied. A `cluster_reboot` issues a Kibana restart command, then waits for the command to complete. When you restart an inactive instance, the most recent successful plan is applied.
*/
func (a *Client) RestartKibanaCluster(params *RestartKibanaClusterParams, authInfo runtime.ClientAuthInfoWriter) (*RestartKibanaClusterAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestartKibanaClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "restart-kibana-cluster",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/{cluster_id}/_restart",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RestartKibanaClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RestartKibanaClusterAccepted), nil

}

/*
ResyncKibanaCluster resynchronizes cluster

Immediately resynchronizes the search index and cache for the selected Kibana instance.
*/
func (a *Client) ResyncKibanaCluster(params *ResyncKibanaClusterParams, authInfo runtime.ClientAuthInfoWriter) (*ResyncKibanaClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResyncKibanaClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "resync-kibana-cluster",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/{cluster_id}/_resync",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResyncKibanaClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ResyncKibanaClusterOK), nil

}

/*
ResyncKibanaClusters resynchronizes clusters

Asynchronously resynchronizes the search index for all Kibana instances.
*/
func (a *Client) ResyncKibanaClusters(params *ResyncKibanaClustersParams, authInfo runtime.ClientAuthInfoWriter) (*ResyncKibanaClustersAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResyncKibanaClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "resync-kibana-clusters",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/_resync",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResyncKibanaClustersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ResyncKibanaClustersAccepted), nil

}

/*
SearchKibanaClusters searches clusters

Retrieves the information for all of the Kibana instances that match the specified query.
*/
func (a *Client) SearchKibanaClusters(params *SearchKibanaClustersParams, authInfo runtime.ClientAuthInfoWriter) (*SearchKibanaClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchKibanaClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "search-kibana-clusters",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/_search",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchKibanaClustersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SearchKibanaClustersOK), nil

}

/*
SetKibanaClusterMetadataRaw sets cluster metadata

Advanced use only. Sets the internal metadata, in free-form JSON, for the Kibana instance.
Only use the parameter to set the modified JSON that is returned from the get version of the metadata.
*/
func (a *Client) SetKibanaClusterMetadataRaw(params *SetKibanaClusterMetadataRawParams, authInfo runtime.ClientAuthInfoWriter) (*SetKibanaClusterMetadataRawOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetKibanaClusterMetadataRawParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "set-kibana-cluster-metadata-raw",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/{cluster_id}/metadata/raw",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetKibanaClusterMetadataRawReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetKibanaClusterMetadataRawOK), nil

}

/*
SetKibanaClusterName sets cluster name

Assigns a name to the Kibana instance.
*/
func (a *Client) SetKibanaClusterName(params *SetKibanaClusterNameParams, authInfo runtime.ClientAuthInfoWriter) (*SetKibanaClusterNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetKibanaClusterNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "set-kibana-cluster-name",
		Method:             "PUT",
		PathPattern:        "/clusters/kibana/{cluster_id}/metadata/name/{new_name}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetKibanaClusterNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetKibanaClusterNameOK), nil

}

/*
ShutdownKibanaCluster shuts down cluster

Shuts down the active Kibana instance and removes all of the instance nodes. The instance definition is retained. WARNING: To avoid data loss, save the snapshot repository before you shut down the instance.
*/
func (a *Client) ShutdownKibanaCluster(params *ShutdownKibanaClusterParams, authInfo runtime.ClientAuthInfoWriter) (*ShutdownKibanaClusterAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShutdownKibanaClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "shutdown-kibana-cluster",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/{cluster_id}/_shutdown",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ShutdownKibanaClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ShutdownKibanaClusterAccepted), nil

}

/*
StartKibanaClusterInstances starts instances

Starts the specified Kibana instances.
*/
func (a *Client) StartKibanaClusterInstances(params *StartKibanaClusterInstancesParams, authInfo runtime.ClientAuthInfoWriter) (*StartKibanaClusterInstancesAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartKibanaClusterInstancesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "start-kibana-cluster-instances",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/{cluster_id}/instances/{instance_ids}/_start",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StartKibanaClusterInstancesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartKibanaClusterInstancesAccepted), nil

}

/*
StartKibanaClusterInstancesAll starts all instances

Starts all of the Kibana instances.
*/
func (a *Client) StartKibanaClusterInstancesAll(params *StartKibanaClusterInstancesAllParams, authInfo runtime.ClientAuthInfoWriter) (*StartKibanaClusterInstancesAllAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartKibanaClusterInstancesAllParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "start-kibana-cluster-instances-all",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/{cluster_id}/instances/_start",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StartKibanaClusterInstancesAllReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartKibanaClusterInstancesAllAccepted), nil

}

/*
StartKibanaClusterInstancesAllMaintenanceMode starts maintenance mode all instances

Starts maintenance mode on all of the Kibana instances.
*/
func (a *Client) StartKibanaClusterInstancesAllMaintenanceMode(params *StartKibanaClusterInstancesAllMaintenanceModeParams, authInfo runtime.ClientAuthInfoWriter) (*StartKibanaClusterInstancesAllMaintenanceModeAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartKibanaClusterInstancesAllMaintenanceModeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "start-kibana-cluster-instances-all-maintenance-mode",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/{cluster_id}/instances/maintenance-mode/_start",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StartKibanaClusterInstancesAllMaintenanceModeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartKibanaClusterInstancesAllMaintenanceModeAccepted), nil

}

/*
StartKibanaClusterMaintenanceMode starts maintenance mode

Starts maintenance mode on the specified Kibana instances.
*/
func (a *Client) StartKibanaClusterMaintenanceMode(params *StartKibanaClusterMaintenanceModeParams, authInfo runtime.ClientAuthInfoWriter) (*StartKibanaClusterMaintenanceModeAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartKibanaClusterMaintenanceModeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "start-kibana-cluster-maintenance-mode",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/{cluster_id}/instances/{instance_ids}/maintenance-mode/_start",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StartKibanaClusterMaintenanceModeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartKibanaClusterMaintenanceModeAccepted), nil

}

/*
StopKibanaClusterInstances stops instances

Stops the specified Kibana instances.
*/
func (a *Client) StopKibanaClusterInstances(params *StopKibanaClusterInstancesParams, authInfo runtime.ClientAuthInfoWriter) (*StopKibanaClusterInstancesAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopKibanaClusterInstancesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stop-kibana-cluster-instances",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/{cluster_id}/instances/{instance_ids}/_stop",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StopKibanaClusterInstancesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StopKibanaClusterInstancesAccepted), nil

}

/*
StopKibanaClusterInstancesAll stops all instances

Stops all of the Kibana instances.
*/
func (a *Client) StopKibanaClusterInstancesAll(params *StopKibanaClusterInstancesAllParams, authInfo runtime.ClientAuthInfoWriter) (*StopKibanaClusterInstancesAllAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopKibanaClusterInstancesAllParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stop-kibana-cluster-instances-all",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/{cluster_id}/instances/_stop",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StopKibanaClusterInstancesAllReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StopKibanaClusterInstancesAllAccepted), nil

}

/*
StopKibanaClusterInstancesAllMaintenanceMode stops maintenance mode all instances

Stops maintenance mode on all of the Kibana instances.
*/
func (a *Client) StopKibanaClusterInstancesAllMaintenanceMode(params *StopKibanaClusterInstancesAllMaintenanceModeParams, authInfo runtime.ClientAuthInfoWriter) (*StopKibanaClusterInstancesAllMaintenanceModeAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopKibanaClusterInstancesAllMaintenanceModeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stop-kibana-cluster-instances-all-maintenance-mode",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/{cluster_id}/instances/maintenance-mode/_stop",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StopKibanaClusterInstancesAllMaintenanceModeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StopKibanaClusterInstancesAllMaintenanceModeAccepted), nil

}

/*
StopKibanaClusterMaintenanceMode stops maintenance mode

Stops maintenance mode on the specified Kibana instances.
*/
func (a *Client) StopKibanaClusterMaintenanceMode(params *StopKibanaClusterMaintenanceModeParams, authInfo runtime.ClientAuthInfoWriter) (*StopKibanaClusterMaintenanceModeAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopKibanaClusterMaintenanceModeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stop-kibana-cluster-maintenance-mode",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/{cluster_id}/instances/{instance_ids}/maintenance-mode/_stop",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StopKibanaClusterMaintenanceModeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StopKibanaClusterMaintenanceModeAccepted), nil

}

/*
UpdateKibanaClusterMetadataSettings updates cluster metadata settings

All changes in the specified object are applied to the metadata object. Omitting existing fields causes the same values to be reapplied. Specifying a `null` value reverts the field to the default value, or removes the field when no default value exists.
*/
func (a *Client) UpdateKibanaClusterMetadataSettings(params *UpdateKibanaClusterMetadataSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateKibanaClusterMetadataSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateKibanaClusterMetadataSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update-kibana-cluster-metadata-settings",
		Method:             "PATCH",
		PathPattern:        "/clusters/kibana/{cluster_id}/metadata/settings",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateKibanaClusterMetadataSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateKibanaClusterMetadataSettingsOK), nil

}

/*
UpdateKibanaClusterPlan updates plan

Updates the configuration of the Kibana instance.
*/
func (a *Client) UpdateKibanaClusterPlan(params *UpdateKibanaClusterPlanParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateKibanaClusterPlanOK, *UpdateKibanaClusterPlanAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateKibanaClusterPlanParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update-kibana-cluster-plan",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/{cluster_id}/plan",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateKibanaClusterPlanReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateKibanaClusterPlanOK:
		return value, nil, nil
	case *UpdateKibanaClusterPlanAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
UpgradeKibanaCluster upgrades cluster

Upgrades an active Kibana instance.
*/
func (a *Client) UpgradeKibanaCluster(params *UpgradeKibanaClusterParams, authInfo runtime.ClientAuthInfoWriter) (*UpgradeKibanaClusterAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpgradeKibanaClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "upgrade-kibana-cluster",
		Method:             "POST",
		PathPattern:        "/clusters/kibana/{cluster_id}/_upgrade",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpgradeKibanaClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpgradeKibanaClusterAccepted), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
