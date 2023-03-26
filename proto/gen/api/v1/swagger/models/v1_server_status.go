// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ServerStatus v1 server status
//
// swagger:model v1ServerStatus
type V1ServerStatus struct {

	// kubernetes distribution
	KubernetesDistribution string `json:"kubernetes_distribution,omitempty"`

	// kubernetes version
	KubernetesVersion string `json:"kubernetes_version,omitempty"`

	// machine ip
	MachineIP string `json:"machine_ip,omitempty"`

	// server version
	ServerVersion string `json:"server_version,omitempty"`
}

// Validate validates this v1 server status
func (m *V1ServerStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 server status based on context it is used
func (m *V1ServerStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ServerStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ServerStatus) UnmarshalBinary(b []byte) error {
	var res V1ServerStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
