package main

import (
	"net/http"

	"github.com/southerncrossedtech/sgen/example/basic/api"
)

// @title example-api
// @description The example API contains a description that
// @description spans across multiple lines
//
// @host api.example.com
// @BasePath /
// @query.collection.format multi
// @schemes https
//
// @license.name The MIT License (MIT)
// @license.url https://choosealicense.com/licenses/mit/
//
// @securityDefinitions.apikey JWT Token
// @in header
// @name Authorization
func main() {
	http.HandleFunc("/example", api.Get)

	http.ListenAndServe(":8080", nil)
}
