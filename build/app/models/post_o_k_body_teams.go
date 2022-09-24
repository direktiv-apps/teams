// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostOKBodyTeams post o k body teams
//
// swagger:model postOKBodyTeams
type PostOKBodyTeams struct {

	// result
	Result int64 `json:"result,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this post o k body teams
func (m *PostOKBodyTeams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post o k body teams based on context it is used
func (m *PostOKBodyTeams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostOKBodyTeams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostOKBodyTeams) UnmarshalBinary(b []byte) error {
	var res PostOKBodyTeams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
