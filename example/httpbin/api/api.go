package api

import (
	"encoding/json"
	"net/http"
	"path"
)

// GetResponse from httpbin
// https://httpbin.org/#/HTTP_Methods/get_get
type GetResponse struct {
	Headers `json:"headers"`
	URL     string `json:"url"`
}

type Headers struct {
	XAPIVersion    string `json:"x-api-version,omitempty"`
	Authorization  string `json:"authorization,omitempty"`
	Accept         string `json:"accept,omitempty"`
	AcceptEncoding string `json:"accept-encoding,omitempty"`
	AcceptLanguage string `json:"accept-language,omitempty"`
	DNT            string `json:"dnt,omitempty"`
	Host           string `json:"host,omitempty"`
	Referer        string `json:"referer,omitempty"`
	UserAgent      string `json:"user-agent,omitempty"`
}

// Get
// @Summary httpbin get request
// @Description Runs a request against the httpbin.org api
// @ID http-methods-list
// @Tags http-methods
// @Accept json
// @Produce json
// @Success 200 {object} api.GetResponse "Ok"
// @Router /get [get]
func Get(w http.ResponseWriter, r *http.Request) {
	p := GetResponse{
		Headers: Headers{
			Authorization:  r.Header.Get("Authorization"),
			XAPIVersion:    r.Header.Get("X-Api-Version"),
			Accept:         r.Header.Get("Accept"),
			AcceptEncoding: r.Header.Get("Accept-Encoding"),
			AcceptLanguage: r.Header.Get("Accept-Language"),
			DNT:            r.Header.Get("Dnt"),
			Host:           r.Header.Get("Host"),
			Referer:        r.Header.Get("Referer"),
			UserAgent:      r.Header.Get("User-Agent"),
		},
		URL: path.Join(r.URL.Scheme, r.Host, r.RequestURI),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(p)
}

// PostResponse from httpbin
// https://httpbin.org/#/HTTP_Methods/post_post
type PostResponse struct {
	Headers `json:"headers"`
	Data    interface{} `json:"data"`
	URL     string      `json:"url"`
}

type PostRequest struct {
	Key string `json:"key"`
}

// Post
// @Summary httpbin post request
// @Description Runs a request against the httpbin.org api
// @ID http-methods-create
// @Tags http-methods
// @Accept json
// @Produce json
// @Param PostRequest body api.PostRequest true "request body"
// @Success 200 {object} api.PostResponse "Ok"
// @Router /post [post]
func Post(w http.ResponseWriter, r *http.Request) {
	var data interface{}

	json.NewDecoder(r.Body).Decode(&data)

	p := PostResponse{
		Headers: Headers{
			Accept:         r.Header.Get("Accept"),
			AcceptEncoding: r.Header.Get("Accept-Encoding"),
			AcceptLanguage: r.Header.Get("Accept-Language"),
			DNT:            r.Header.Get("Dnt"),
			Host:           r.Header.Get("Host"),
			Referer:        r.Header.Get("Referer"),
			UserAgent:      r.Header.Get("User-Agent"),
		},
		Data: data,
		URL:  path.Join(r.URL.Scheme, r.Host, r.RequestURI),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(p)
}

// PutResponse from httpbin
// https://httpbin.org/#/HTTP_Methods/put_put
type PutResponse struct {
	Headers `json:"headers"`
	Data    interface{} `json:"data"`
	URL     string      `json:"url"`
}

type PutRequest struct {
	Key string `json:"key"`
}

// Put
// @Summary httpbin put request
// @Description Runs a request against the httpbin.org api
// @ID http-methods-update
// @Tags http-methods
// @Accept json
// @Produce json
// @Param PutRequest body api.PutRequest true "request body"
// @Success 200 {object} api.PutResponse "Ok"
// @Router /put [put]
func Put(w http.ResponseWriter, r *http.Request) {
	var data interface{}

	json.NewDecoder(r.Body).Decode(&data)

	p := PostResponse{
		Headers: Headers{
			Accept:         r.Header.Get("Accept"),
			AcceptEncoding: r.Header.Get("Accept-Encoding"),
			AcceptLanguage: r.Header.Get("Accept-Language"),
			DNT:            r.Header.Get("Dnt"),
			Host:           r.Header.Get("Host"),
			Referer:        r.Header.Get("Referer"),
			UserAgent:      r.Header.Get("User-Agent"),
		},
		Data: data,
		URL:  path.Join(r.URL.Scheme, r.Host, r.RequestURI),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(p)
}

// DeleteResponse from httpbin
// https://httpbin.org/#/HTTP_Methods/delete_delete
type DeleteResponse struct {
	Headers `json:"headers"`
	Data    interface{} `json:"data"`
	URL     string      `json:"url"`
}

type DeleteRequest struct {
	Key string `json:"key"`
}

// Delete
// @Summary httpbin delete request
// @Description Runs a request against the httpbin.org api
// @ID http-methods-delete
// @Tags http-methods
// @Accept json
// @Produce json
// @Param DeleteRequest body api.DeleteRequest true "request body"
// @Success 200 {object} api.DeleteResponse "Ok"
// @Router /delete [delete]
func Delete(w http.ResponseWriter, r *http.Request) {
	var data interface{}

	json.NewDecoder(r.Body).Decode(&data)

	p := PostResponse{
		Headers: Headers{
			Accept:         r.Header.Get("Accept"),
			AcceptEncoding: r.Header.Get("Accept-Encoding"),
			AcceptLanguage: r.Header.Get("Accept-Language"),
			DNT:            r.Header.Get("Dnt"),
			Host:           r.Header.Get("Host"),
			Referer:        r.Header.Get("Referer"),
			UserAgent:      r.Header.Get("User-Agent"),
		},
		Data: data,
		URL:  path.Join(r.URL.Scheme, r.Host, r.RequestURI),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(p)
}
