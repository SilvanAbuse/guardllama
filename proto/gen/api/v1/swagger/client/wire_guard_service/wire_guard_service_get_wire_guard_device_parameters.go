// Code generated by go-swagger; DO NOT EDIT.

package wire_guard_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewWireGuardServiceGetWireGuardDeviceParams creates a new WireGuardServiceGetWireGuardDeviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWireGuardServiceGetWireGuardDeviceParams() *WireGuardServiceGetWireGuardDeviceParams {
	return &WireGuardServiceGetWireGuardDeviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWireGuardServiceGetWireGuardDeviceParamsWithTimeout creates a new WireGuardServiceGetWireGuardDeviceParams object
// with the ability to set a timeout on a request.
func NewWireGuardServiceGetWireGuardDeviceParamsWithTimeout(timeout time.Duration) *WireGuardServiceGetWireGuardDeviceParams {
	return &WireGuardServiceGetWireGuardDeviceParams{
		timeout: timeout,
	}
}

// NewWireGuardServiceGetWireGuardDeviceParamsWithContext creates a new WireGuardServiceGetWireGuardDeviceParams object
// with the ability to set a context for a request.
func NewWireGuardServiceGetWireGuardDeviceParamsWithContext(ctx context.Context) *WireGuardServiceGetWireGuardDeviceParams {
	return &WireGuardServiceGetWireGuardDeviceParams{
		Context: ctx,
	}
}

// NewWireGuardServiceGetWireGuardDeviceParamsWithHTTPClient creates a new WireGuardServiceGetWireGuardDeviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewWireGuardServiceGetWireGuardDeviceParamsWithHTTPClient(client *http.Client) *WireGuardServiceGetWireGuardDeviceParams {
	return &WireGuardServiceGetWireGuardDeviceParams{
		HTTPClient: client,
	}
}

/*
WireGuardServiceGetWireGuardDeviceParams contains all the parameters to send to the API endpoint

	for the wire guard service get wire guard device operation.

	Typically these are written to a http.Request.
*/
type WireGuardServiceGetWireGuardDeviceParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the wire guard service get wire guard device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WireGuardServiceGetWireGuardDeviceParams) WithDefaults() *WireGuardServiceGetWireGuardDeviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the wire guard service get wire guard device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WireGuardServiceGetWireGuardDeviceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the wire guard service get wire guard device params
func (o *WireGuardServiceGetWireGuardDeviceParams) WithTimeout(timeout time.Duration) *WireGuardServiceGetWireGuardDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the wire guard service get wire guard device params
func (o *WireGuardServiceGetWireGuardDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the wire guard service get wire guard device params
func (o *WireGuardServiceGetWireGuardDeviceParams) WithContext(ctx context.Context) *WireGuardServiceGetWireGuardDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the wire guard service get wire guard device params
func (o *WireGuardServiceGetWireGuardDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the wire guard service get wire guard device params
func (o *WireGuardServiceGetWireGuardDeviceParams) WithHTTPClient(client *http.Client) *WireGuardServiceGetWireGuardDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the wire guard service get wire guard device params
func (o *WireGuardServiceGetWireGuardDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *WireGuardServiceGetWireGuardDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
