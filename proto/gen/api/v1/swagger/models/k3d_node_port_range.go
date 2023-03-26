// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// K3dNodePortRange k3d node port range
//
// swagger:model K3dNodePortRange
type K3dNodePortRange struct {

	// from port
	FromPort int32 `json:"from_port,omitempty"`

	// host
	Host string `json:"host,omitempty"`

	// protocol
	Protocol *NodePortRangeProtocol `json:"protocol,omitempty"`

	// to port
	ToPort int32 `json:"to_port,omitempty"`
}

// Validate validates this k3d node port range
func (m *K3dNodePortRange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K3dNodePortRange) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	if m.Protocol != nil {
		if err := m.Protocol.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protocol")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this k3d node port range based on the context it is used
func (m *K3dNodePortRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K3dNodePortRange) contextValidateProtocol(ctx context.Context, formats strfmt.Registry) error {

	if m.Protocol != nil {
		if err := m.Protocol.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protocol")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *K3dNodePortRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K3dNodePortRange) UnmarshalBinary(b []byte) error {
	var res K3dNodePortRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
