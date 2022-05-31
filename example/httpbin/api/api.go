package api

import "net/http"

// GetResponse from httpbin
// https://httpbin.org/#/HTTP_Methods/get_get
type GetResponse struct {
	Headers struct {
		Accept         string
		AcceptEncoding string
		AcceptLanguage string
		DNT            int
		Host           string
		Referer        string
		UserAgent      string
	}
	URL string
}

// Get
// @Summary httpbin get request
// @Description Runs a request against the httpbin.org api
// @ID httpbin-get
// @Accept json
// @Produce json
// @Success 200 {object} api.GetResponse "Ok"
// @Router /get [get]
func Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
