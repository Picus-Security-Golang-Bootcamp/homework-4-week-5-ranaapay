package service

import (
	"PicusHomework4/src/service/helper"
	"PicusHomework4/src/storage"
	"PicusHomework4/src/storage/storageType"
)

type AuthorService struct {
	repo *storage.AuthorRepository
}

func NewAuthorService(repo *storage.AuthorRepository) AuthorService {
	return AuthorService{repo: repo}
}

func (s *AuthorService) FindAuthorById(id int) (*storageType.Author, error) {
	res, err := s.repo.FindAuthorById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AuthorService) CreateAuthor(author storageType.Author) (uint, error) {
	resId, err := s.repo.CreateAuthor(author)
	if err != nil {
		return 0, err
	}
	return resId,nil
}

func (s *AuthorService) UpdateAuthorById(author storageType.Author) error {
	updateOptions := helper.SetAuthorUpdateOptions(author)
	if err := s.repo.UpdateAuthorById(author, updateOptions); err != nil {
		return err
	}
	return nil
}

func (s *AuthorService) DeleteAuthorById(id int) error {
	if err := s.repo.DeleteAuthorById(uint(id)); err != nil {
		return err
	}
	return nil
}

func (s *AuthorService) FindAuthorByIdWithBooks (id int) (*storageType.Author, error) {
	res, err := s.repo.GetAuthorByIdWithBooks(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
