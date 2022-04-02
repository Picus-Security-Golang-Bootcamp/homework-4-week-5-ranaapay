package handler

import (
	_type "PicusHomework4/src/handler/type"
	"PicusHomework4/src/pkg/httpErrors"
	"PicusHomework4/src/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type BookHandler struct {
	service service.BookService
}

func NewBookHandler(bookService service.BookService) BookHandler {
	return BookHandler{service: bookService}
}

func (h *BookHandler) FindBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	reqId, err := strconv.Atoi(id)
	if err != nil {
		errResp := httpErrors.ParseError(httpErrors.ConvertIdError)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	book, err := h.service.FindBookById(reqId)
	if err != nil {
		errResp := httpErrors.ParseError(err)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	bookResp := _type.NewResponseBookType(book)
	resp, _ := json.Marshal(_type.NewResponseType(http.StatusOK,bookResp))
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var bookReq _type.RequestBookType
	if r.Header.Get("Content-Type") != "application/json" {
		errResp := httpErrors.ParseError(httpErrors.ContentTypeError)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&bookReq); err != nil {
		errResp := httpErrors.ParseError(httpErrors.CannotDecodeError)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	if err := bookReq.ValidateBookRequest(); err != nil {
		errResp := httpErrors.ParseError(err)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	book := bookReq.NewRequestTypeToBook()
	resId, err := h.service.CreateBook(book)
	if err != nil {
		errResp := httpErrors.ParseError(err)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	resp, _ := json.Marshal(_type.NewResponseType(http.StatusCreated,_type.IdResponseType{Id: int(resId)}))
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	reqId, err := strconv.Atoi(id)
	if err != nil {
		errResp := httpErrors.ParseError(httpErrors.ConvertIdError)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	var bookReq _type.RequestBookType
	if err = json.NewDecoder(r.Body).Decode(&bookReq); err != nil {
		errResp := httpErrors.ParseError(httpErrors.CannotDecodeError)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	if err = bookReq.ValidateBookRequest(); err != nil {
		errResp := httpErrors.ParseError(err)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	book := bookReq.NewRequestTypeToBook()
	book.ID = uint(reqId)
	if err = h.service.UpdateBookById(book); err != nil {
		errResp := httpErrors.ParseError(err)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	resp, _ := json.Marshal(_type.NewResponseType(http.StatusOK,_type.QuickResponseType{Result: true}))
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	reqId, err := strconv.Atoi(id)
	if err != nil {
		errResp := httpErrors.ParseError(httpErrors.ConvertIdError)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	if err = h.service.DeleteBookById(reqId); err != nil {
		errResp := httpErrors.ParseError(err)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	resp, _ := json.Marshal(_type.NewResponseType(http.StatusOK,_type.QuickResponseType{Result: true}))
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}


//___________________________________________________________________

func (h *BookHandler) FindBookByIdWithAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	reqId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book, err := h.service.FindBookByIdWithAuthor(reqId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	resp, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}