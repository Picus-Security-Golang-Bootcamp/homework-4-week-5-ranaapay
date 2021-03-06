// Code generated by go-swagger; DO NOT EDIT.

package swagger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponseType response type
//
// swagger:model ResponseType
type ResponseType struct {

	// code
	Code int64 `json:"code,omitempty"`

	// message
	Message interface{} `json:"message,omitempty"`
}

// Validate validates this response type
func (m *ResponseType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this response type based on context it is used
func (m *ResponseType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponseType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseType) UnmarshalBinary(b []byte) error {
	var res ResponseType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
