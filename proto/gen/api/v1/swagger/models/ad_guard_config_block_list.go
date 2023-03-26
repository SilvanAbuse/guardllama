// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdGuardConfigBlockList ad guard config block list
//
// swagger:model AdGuardConfigBlockList
type AdGuardConfigBlockList struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this ad guard config block list
func (m *AdGuardConfigBlockList) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ad guard config block list based on context it is used
func (m *AdGuardConfigBlockList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AdGuardConfigBlockList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdGuardConfigBlockList) UnmarshalBinary(b []byte) error {
	var res AdGuardConfigBlockList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
