package service

import (
	"PicusHomework4/src/service/helper"
	"PicusHomework4/src/storage"
	"PicusHomework4/src/storage/storageType"
)

type BookService struct {
	repo storage.BookRepository
}

func NewBookService(repo storage.BookRepository) BookService {
	return BookService{repo: repo}
}

func (s *BookService) FindBookById(id int) (*storageType.Book, error) {
	res, err := s.repo.FindBookById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *BookService) CreateBook(book storageType.Book) (uint, error) {
	resId, err := s.repo.CreateBook(book)
	if err != nil {
		return 0, err
	}
	return resId,nil
}

func (s *BookService) UpdateBookById(book storageType.Book) error {
	updateOptions := helper.SetBookUpdateOptions(book)
	if err := s.repo.UpdateBookById(book, updateOptions); err != nil {
		return err
	}
	return nil
}

func (s *BookService) DeleteBookById(id int) error {
	if err := s.repo.DeleteBookById(uint(id)); err != nil {
		return err
	}
	return nil
}

//___________________________________________________________________

func (s *BookService) FindBookByIdWithAuthor (id int) (*storageType.Book, error) {
	res, err := s.repo.FindBookByIdWithAuthor(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}