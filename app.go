package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// App main entry for program
type App struct {
	Router *mux.Router
	DB     *sqlx.DB
}

// Initialize database and routes
func (a *App) Initialize(dataSourceName string) {
	var err error

	a.DB, err = sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

// Run serve up http listeners
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

// InitializeRoutes setup RESTful routing
func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/books", a.GetBooks).Methods("GET")
	a.Router.HandleFunc("/book", a.CreateBook).Methods("POST")

	a.Router.HandleFunc("/book/{id:[0-9]+}", a.GetBook).Methods("GET")
	a.Router.HandleFunc("/book/{id:[0-9]+}", a.UpdateBook).Methods("PUT")
	a.Router.HandleFunc("/book/{id:[0-9]+}", a.DeleteBook).Methods("DELETE")

	a.Router.HandleFunc("/authors", a.GetAuthors).Methods("GET")
	a.Router.HandleFunc("/author", a.CreateAuthor).Methods("POST")

	a.Router.HandleFunc("/author/{id:[0-9]+}", a.GetAuthor).Methods("GET")
	a.Router.HandleFunc("/author/{id:[0-9]+}", a.UpdateAuthor).Methods("PUT")
	a.Router.HandleFunc("/author/{id:[0-9]+}", a.DeleteAuthor).Methods("DELETE")

	a.Router.HandleFunc("/publishers", a.GetPublishers).Methods("GET")
	a.Router.HandleFunc("/publisher", a.CreatePublisher).Methods("POST")

	a.Router.HandleFunc("/publisher/{id:[0-9]+}", a.GetPublisher).Methods("GET")
	a.Router.HandleFunc("/publisher/{id:[0-9]+}", a.UpdatePublisher).Methods("PUT")
	a.Router.HandleFunc("/publisher/{id:[0-9]+}", a.DeletePublisher).Methods("DELETE")
}

// RespondWithError json error response
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON normal json response
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// GetBook a single book
func (a *App) GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	b := Book{ID: id}
	if err := b.GetBook(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Book not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	RespondWithJSON(w, http.StatusOK, b)
}

// GetBooks all books
func (a *App) GetBooks(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	books, err := GetBooks(a.DB, start, count)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, books)
}

// CreateBook new book
func (a *App) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&book); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := book.CreateBook(a.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, book)
}

// UpdateBook post request on book
func (a *App) UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	var book Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	book.ID = id

	if err := book.UpdateBook(a.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, book)
}

// DeleteBook remove book from db
func (a *App) DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}

	book := Book{ID: id}
	if err := book.DeleteBook(a.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// GetAuthor return a single author
func (a *App) GetAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid author ID")
		return
	}

	author := Author{ID: id}
	if err := author.GetAuthor(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Author not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	RespondWithJSON(w, http.StatusOK, author)
}

// GetAuthors all authors
func (a *App) GetAuthors(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	authors, err := GetAuthors(a.DB, start, count)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, authors)
}

// CreateAuthor a new author
func (a *App) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author Author
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&author); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := author.CreateAuthor(a.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, author)
}

// UpdateAuthor update author attributes
func (a *App) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid author ID")
		return
	}

	var author Author
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&author); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	author.ID = id

	if err := author.UpdateAuthor(a.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, author)
}

// DeleteAuthor remove author
func (a *App) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid author ID")
		return
	}

	author := Author{ID: id}
	if err := author.DeleteAuthor(a.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// GetPublisher a single publisher
func (a *App) GetPublisher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid publisher ID")
		return
	}

	p := Publisher{ID: id}
	if err := p.GetPublisher(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Publisher not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	RespondWithJSON(w, http.StatusOK, p)
}

// GetPublishers all publishers
func (a *App) GetPublishers(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	publishers, err := GetPublishers(a.DB, start, count)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, publishers)
}

// CreatePublisher a new publisher
func (a *App) CreatePublisher(w http.ResponseWriter, r *http.Request) {
	var publisher Publisher
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&publisher); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := publisher.CreatePublisher(a.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, publisher)
}

// UpdatePublisher update publisher attributes
func (a *App) UpdatePublisher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid publisher ID")
		return
	}

	var publisher Publisher
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&publisher); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	publisher.ID = id

	if err := publisher.UpdatePublisher(a.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, publisher)
}

// DeletePublisher remove publisher
func (a *App) DeletePublisher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid publisher ID")
		return
	}

	publisher := Publisher{ID: id}
	if err := publisher.DeletePublisher(a.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
