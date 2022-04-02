// Code generated by go-swagger; DO NOT EDIT.

package swagger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BookRequest book request
//
// swagger:model BookRequest
type BookRequest struct {

	// author Id
	// Required: true
	AuthorID *uint64 `json:"authorId"`

	// isbn
	// Required: true
	Isbn *string `json:"isbn"`

	// name
	// Required: true
	Name *string `json:"name"`

	// page number
	// Required: true
	PageNumber *int64 `json:"pageNumber"`

	// price
	// Required: true
	Price *string `json:"price"`

	// publish time
	// Required: true
	PublishTime *string `json:"publishTime"`

	// stock code
	// Required: true
	StockCode *string `json:"stockCode"`

	// stock number
	// Required: true
	StockNumber *int64 `json:"stockNumber"`
}

// Validate validates this book request
func (m *BookRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsbn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStockCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStockNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BookRequest) validateAuthorID(formats strfmt.Registry) error {

	if err := validate.Required("authorId", "body", m.AuthorID); err != nil {
		return err
	}

	return nil
}

func (m *BookRequest) validateIsbn(formats strfmt.Registry) error {

	if err := validate.Required("isbn", "body", m.Isbn); err != nil {
		return err
	}

	return nil
}

func (m *BookRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *BookRequest) validatePageNumber(formats strfmt.Registry) error {

	if err := validate.Required("pageNumber", "body", m.PageNumber); err != nil {
		return err
	}

	return nil
}

func (m *BookRequest) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

func (m *BookRequest) validatePublishTime(formats strfmt.Registry) error {

	if err := validate.Required("publishTime", "body", m.PublishTime); err != nil {
		return err
	}

	return nil
}

func (m *BookRequest) validateStockCode(formats strfmt.Registry) error {

	if err := validate.Required("stockCode", "body", m.StockCode); err != nil {
		return err
	}

	return nil
}

func (m *BookRequest) validateStockNumber(formats strfmt.Registry) error {

	if err := validate.Required("stockNumber", "body", m.StockNumber); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this book request based on context it is used
func (m *BookRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BookRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BookRequest) UnmarshalBinary(b []byte) error {
	var res BookRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}