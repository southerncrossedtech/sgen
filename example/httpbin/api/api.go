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
// @ID httpbin-list
// @Tags httpbin
// @Accept json
// @Produce json
// @Success 200 {object} api.GetResponse "Ok"
// @Router /get [get]
func Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// PostResponse from httpbin
// https://httpbin.org/#/HTTP_Methods/post_post
type PostResponse struct {
	Headers struct {
		Accept         string
		AcceptEncoding string
		ContentLength  uint
		Host           string
		Referer        string
		UserAgent      string
	}
	URL string
}

// Post
// @Summary httpbin post request
// @Description Runs a request against the httpbin.org api
// @ID httpbin-create
// @Tags httpbin
// @Accept json
// @Produce json
// @Success 200 {object} api.PostResponse "Ok"
// @Router /post [post]
func Post(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// PutResponse from httpbin
// https://httpbin.org/#/HTTP_Methods/put_put
type PutResponse struct {
	Headers struct {
		Accept         string
		AcceptEncoding string
		ContentLength  uint
		Host           string
		Referer        string
		UserAgent      string
	}
	URL string
}

// Put
// @Summary httpbin put request
// @Description Runs a request against the httpbin.org api
// @ID httpbin-update
// @Tags httpbin
// @Accept json
// @Produce json
// @Success 200 {object} api.PutResponse "Ok"
// @Router /put [put]
func Put(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// DeleteResponse from httpbin
// https://httpbin.org/#/HTTP_Methods/delete_delete
type DeleteResponse struct {
	Headers struct {
		Accept         string
		AcceptEncoding string
		Host           string
		UserAgent      string
	}
	URL string
}

// Delete
// @Summary httpbin delete request
// @Description Runs a request against the httpbin.org api
// @ID httpbin-delete
// @Tags httpbin
// @Accept json
// @Produce json
// @Success 200 {object} api.DeleteResponse "Ok"
// @Router /delete [delete]
func Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
