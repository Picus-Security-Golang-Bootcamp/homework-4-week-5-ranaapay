package _type

import "PicusHomework4/src/storage/storageType"

type ResponseAuthorType struct {
	Id       uint               `json:"id"`
	Name     string             `json:"firstName"`
	LastName string             `json:"lastName"`
	Email    string             `json:"email"`
	Gender   string             `json:"gender"`
	Books    []storageType.Book `json:",omitempty"`
}

func NewResponseAuthorType(a *storageType.Author) *ResponseAuthorType {
	return &ResponseAuthorType{
		Id:       a.ID,
		Name:     a.Name,
		LastName: a.LastName,
		Email:    a.Email,
		Gender:   a.Gender,
		Books:    a.Books,
	}
}
