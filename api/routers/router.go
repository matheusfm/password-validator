package routers

import (
	"encoding/json"
	"github.com/matheusfm/password-validator/api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheusfm/password-validator/api/handlers"
)

func Routers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	router.HandleFunc("/validation", handlers.PasswordValidator).Methods(http.MethodPost)
	router.Use(middlewares.JsonContentType)
	return router
}
