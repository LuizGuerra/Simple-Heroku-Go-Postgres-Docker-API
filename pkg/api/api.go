package api

import (
	"Ictus-Backend/pkg/db"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg/v10"
	"log"
	"net/http"
	"strconv"
)

func NewApi(pgDb *pg.DB) *chi.Mux { //
	// set up router
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.WithValue("DB", pgDb))

	r.Route("/homes", func(r chi.Router) {
		r.Get("/", getHomes)
		r.Get("/{homeID}", getHomeById)
		r.Post("/", createHome)
		r.Put("/{homeID}", updateHome)
		r.Delete("/{homeID}", deleteHome)
	})
	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Welcome!"))
	//})

	return r
}

func CanWriteError(err error, w http.ResponseWriter, status int, message string) bool {
	if err != nil {
		w.WriteHeader(status)

		err2 := json.NewEncoder(w).Encode(message)
		if err2 != nil {
			log.Printf("Error sending error message %s\n", err)
		}

		log.Println(message)
		return true
	}
	return false
}

// Json Encode Message
func jem(w http.ResponseWriter, v any) {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Printf("Error sending error message %s\n", err)
	}
}

func getHomes(w http.ResponseWriter, r *http.Request) {
	// get database from context
	pgDb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		logMessage := "Critical: server failed to retrieve database."
		jem(w, logMessage)
		log.Println(logMessage)
		return
	}

	// query the houses
	homes, err := db.GetHomes(pgDb)
	if CanWriteError(err, w, http.StatusInternalServerError,
		"Failed getting home array") {
		return
	}

	// return the response
	err = json.NewEncoder(w).Encode(homes)
	if err != nil {
		log.Printf("Error sending error message %s\n", err)
	}
	w.WriteHeader(http.StatusOK)
}

func getHomeById(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")

	// get database from context
	pgDb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		logMessage := "Critical: server failed to retrieve database."
		jem(w, logMessage)
		log.Println(logMessage)
		return
	}

	// query the house
	home, err := db.GetHome(pgDb, homeID)
	if CanWriteError(err, w, http.StatusInternalServerError,
		"Failed getting home") {
		return
	}

	// return the response
	jem(w, home)
	w.WriteHeader(http.StatusOK)
}

func createHome(w http.ResponseWriter, r *http.Request) {
	// parse request body
	req := &CreateHomeRequest{}
	if json.NewDecoder(r.Body).Decode(req) != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// get the database from context
	pgDb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		logMessage := "Critical: server failed to retrieve database."
		jem(w, logMessage)
		log.Println(logMessage)
		return
	}
	// insert our home
	newHome, err := db.CreateHome(pgDb, &db.Home{
		Price:   req.Price,
		AgentId: req.AgentID,
	})
	if CanWriteError(err, w, http.StatusInternalServerError,
		"Server failed to write new data.") {
		return
	}
	// return response
	jem(w, newHome)
	w.WriteHeader(http.StatusCreated)
}

type CreateHomeRequest struct {
	Price   int64 `json:"price"`
	AgentID int64 `json:"agent_id"`
}

type UpdateHomeRequest struct {
	Id      int64 `json:"id"`
	Price   int64 `json:"price"`
	AgentID int64 `json:"agent_id"`
}

func updateHome(w http.ResponseWriter, r *http.Request) {
	// parse request body
	req := &UpdateHomeRequest{}
	if json.NewDecoder(r.Body).Decode(req) != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// get the database from context
	pgDb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		logMessage := "Critical: server failed to retrieve database."
		jem(w, logMessage)
		log.Println(logMessage)
		return
	}
	// update the home
	updatedHome, err := db.UpdateHome(pgDb, &db.Home{
		ID:      req.Id,
		Price:   req.Price,
		AgentId: req.AgentID,
	})
	if CanWriteError(err, w, http.StatusInternalServerError,
		"Server failed to write new data.") {
		return
	}
	// return response
	jem(w, updatedHome)
	w.WriteHeader(http.StatusCreated)
}

func deleteHome(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")

	// get database from context
	pgDb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		logMessage := "Critical: server failed to retrieve database."
		jem(w, logMessage)
		log.Println(logMessage)
		return
	}

	// query the house
	intHomeID, err := strconv.ParseInt(homeID, 10, 64)
	if CanWriteError(err, w, http.StatusBadRequest,
		"Passed ID don't match int64") {
		return
	}
	err = db.DeleteHomeWEINE(pgDb, intHomeID)
	if CanWriteError(err, w, http.StatusNotFound,
		"Failed getting home") {
		return
	}

	// return the response
	w.WriteHeader(http.StatusOK)
}
