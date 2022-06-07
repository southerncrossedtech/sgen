package sdk

import "github.com/go-openapi/spec"

type ClientSDK struct {
	// CurrentDir sets the path for sgen to run in. Defaults to the
	// current directory.
	CurrentDir string
	// TemplateData is the dynamic data parts fetched and assumed from
	// the swagger.yaml file to render the client api sdk.
	TemplateData
}

type TemplateData struct {
	SGen
	Spec *spec.Swagger
}

// Metadata contains the global details for a typical client sdk
type SGen struct {
	// SGenVersion is the version of the sdk generator library.
	Version string
}
