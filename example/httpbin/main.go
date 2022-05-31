package main

import (
	"net/http"

	"github.com/southerncrossedtech/sgen/example/httpbin/api"
)

// @title httpbin
// @description A simple HTTP Request & Response Service.
// @version v1.0
//
// @host https://httpbin.org
// @BasePath /
// @query.collection.format multi
// @schemes https
//
// @license.name ISC License
// @license.url https://github.com/postmanlabs/httpbin/blob/master/LICENSE
func main() {
	http.HandleFunc("/get", api.Get)
	// http.HandleFunc("/put", api.Put)
	// http.HandleFunc("/post", api.Post)
	// http.HandleFunc("/delete", api.Delete)

	http.ListenAndServe(":8080", nil)
}
