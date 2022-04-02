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

type AuthorHandler struct {
	service service.AuthorService
}

func NewAuthorHandler(authorService service.AuthorService) AuthorHandler {
	return AuthorHandler{service: authorService}
}

func (h *AuthorHandler) FindAuthorById(w http.ResponseWriter, r *http.Request){
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
	author, err := h.service.FindAuthorById(reqId)
	if err != nil {
		errResp := httpErrors.ParseError(err)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	authorResp := _type.NewResponseAuthorType(author)
	resp, _ := json.Marshal(_type.NewResponseType(http.StatusOK,authorResp))
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h *AuthorHandler) FindAuthorByIdWithBooks(w http.ResponseWriter, r *http.Request){
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
	author, err := h.service.FindAuthorByIdWithBooks(reqId)
	if err != nil {
		errResp := httpErrors.ParseError(err)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	authorResp := _type.NewResponseAuthorType(author)
	resp, _ := json.Marshal(_type.NewResponseType(http.StatusOK,authorResp))
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var authorReq _type.RequestAuthorType
	if r.Header.Get("Content-Type") != "application/json" {
		errResp := httpErrors.ParseError(httpErrors.ContentTypeError)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&authorReq); err != nil {
		errResp := httpErrors.ParseError(httpErrors.CannotDecodeError)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	if err := authorReq.ValidateAuthorRequest(); err != nil {
		errResp := httpErrors.ParseError(err)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	author := authorReq.NewRequestTypeToAuthor()
	resId, err := h.service.CreateAuthor(author)
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

func (h *AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
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
	var authorReq _type.RequestAuthorType
	if err = json.NewDecoder(r.Body).Decode(&authorReq); err != nil {
		errResp := httpErrors.ParseError(httpErrors.CannotDecodeError)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	if err = authorReq.ValidateAuthorRequest(); err != nil {
		errResp := httpErrors.ParseError(err)
		resp, _ := json.Marshal(_type.NewResponseType(errResp.ErrCode, errResp))
		w.WriteHeader(errResp.ErrCode)
		w.Write(resp)
		return
	}
	author := authorReq.NewRequestTypeToAuthor()
	author.ID = uint(reqId)
	if err = h.service.UpdateAuthorById(author); err != nil {
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

func (h *AuthorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
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
	if err = h.service.DeleteAuthorById(reqId); err != nil {
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