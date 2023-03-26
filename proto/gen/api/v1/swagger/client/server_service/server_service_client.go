// Code generated by go-swagger; DO NOT EDIT.

package server_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new server service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for server service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ServerServiceGetServer(params *ServerServiceGetServerParams, opts ...ClientOption) (*ServerServiceGetServerOK, error)

	ServerServiceGetServerConfig(params *ServerServiceGetServerConfigParams, opts ...ClientOption) (*ServerServiceGetServerConfigOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ServerServiceGetServer server service get server API
*/
func (a *Client) ServerServiceGetServer(params *ServerServiceGetServerParams, opts ...ClientOption) (*ServerServiceGetServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServerServiceGetServerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ServerService_GetServer",
		Method:             "GET",
		PathPattern:        "/server",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServerServiceGetServerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServerServiceGetServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServerServiceGetServerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ServerServiceGetServerConfig server service get server config API
*/
func (a *Client) ServerServiceGetServerConfig(params *ServerServiceGetServerConfigParams, opts ...ClientOption) (*ServerServiceGetServerConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServerServiceGetServerConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ServerService_GetServerConfig",
		Method:             "GET",
		PathPattern:        "/server/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServerServiceGetServerConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServerServiceGetServerConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServerServiceGetServerConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
