package route

import (
	"../app/controller"
	"encoding/json"
	"net/http"
)

// Handler route
func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	res := controller.SetupEndpoint()
	json.NewEncoder(w).Encode(res)
}
