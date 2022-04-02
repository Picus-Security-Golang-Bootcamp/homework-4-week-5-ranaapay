package cmd

import (
	"PicusHomework4/src/handler"
	postgres "PicusHomework4/src/pkg/db"
	"PicusHomework4/src/pkg/httpErrors"
	"PicusHomework4/src/service"
	"PicusHomework4/src/storage"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func Execute() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(httpErrors.ENVFileError.Error())
	}
	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Postgres connected")

	r := mux.NewRouter()
	InitializeBookRoute(r, db)
	InitializeAuthorRoute(r, db)

	srv := &http.Server{
		Addr:         "localhost:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
	log.Println(srv.ListenAndServe())
}

func InitializeBookRoute(r *mux.Router, db *gorm.DB) {
	bookRepository := storage.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	bookRouter := r.PathPrefix("/book").Subrouter()
	bookRouter.HandleFunc("/{id}", bookHandler.FindBookById).Methods(http.MethodGet)
	//bookRouter.HandleFunc("/{id}/author", bookHandler.FindBookByIdWithAuthor).Methods(http.MethodGet)
	bookRouter.HandleFunc("/", bookHandler.CreateBook).Methods(http.MethodPost)
	bookRouter.HandleFunc("/{id}", bookHandler.UpdateBook).Methods(http.MethodPut)
	bookRouter.HandleFunc("/{id}", bookHandler.DeleteBook).Methods(http.MethodDelete)
}

func InitializeAuthorRoute(r *mux.Router, db *gorm.DB) {
	authorRepository := storage.NewAuthorRepository(db)
	authorService := service.NewAuthorService(authorRepository)
	authorHandler := handler.NewAuthorHandler(authorService)

	authorRouter := r.PathPrefix("/author").Subrouter()
	authorRouter.HandleFunc("/{id}", authorHandler.FindAuthorById).Methods(http.MethodGet)
	authorRouter.HandleFunc("/{id}/books", authorHandler.FindAuthorByIdWithBooks).Methods(http.MethodGet)
	authorRouter.HandleFunc("/", authorHandler.CreateAuthor).Methods(http.MethodPost)
	authorRouter.HandleFunc("/{id}", authorHandler.UpdateAuthor).Methods(http.MethodPut)
	authorRouter.HandleFunc("/{id}", authorHandler.DeleteAuthor).Methods(http.MethodDelete)
}
