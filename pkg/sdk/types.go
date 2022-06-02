package sdk

// Metadata contains the global details for a typical client sdk
type Metadata struct {
	// SGenVersion is the version of the sdk generator library.
	SGenVersion string
	// Title contains the lowercase title for your sdk service. This is
	// typically what you would call your API.
	// Swagger destination: info.title
	Title string
	// Version contains the version number of your API.
	// Swagger destination: info.version
	Version string
}

type ClientSDK struct {
	// CurrentDir sets the path for sgen to run in. Defaults to the
	// current directory.
	CurrentDir string
	// TemplateData is the dynamic data parts fetched and assumed from
	// the swagger.yaml file to render the client api sdk.
	TemplateData
}

type TemplateData struct {
	Metadata
}
