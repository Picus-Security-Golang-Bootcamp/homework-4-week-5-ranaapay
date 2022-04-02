package _type

import (
	"PicusHomework4/src/pkg/httpErrors"
	"PicusHomework4/src/storage/storageType"
)

type RequestAuthorType struct {
	Name     string `json:"firstName"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}

func (a RequestAuthorType) NewRequestTypeToAuthor() storageType.Author {
	return storageType.Author{
		Name:     a.Name,
		LastName: a.LastName,
		Email:    a.Email,
		Gender:   a.Gender,
	}
}

func (a RequestAuthorType) ValidateAuthorRequest() error {
	if a.Name == "" {
		return httpErrors.FirstNameValidationError
	}
	if a.LastName == "" {
		return httpErrors.LastNameValidationError
	}
	if a.Email == "" {
		return httpErrors.EmailValidationError
	}
	if a.Gender == "" {
		return httpErrors.GenderValidationError
	}
	return nil
}
