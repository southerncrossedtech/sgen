package api

import "net/http"

// Get example
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID get-example
// @Produce  json
// @Success 204 "No Content"
// @Router /example [get]
func Get(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNoContent)

	// _ = json.NewEncoder(w).Encode(payload)
}
