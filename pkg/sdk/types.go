package sdk

import (
	"reflect"
	"strings"

	"github.com/go-openapi/spec"
	"github.com/k0kubun/pp/v3"
)

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
	Spec       *spec.Swagger
	Operations map[string][]Operation
}

// Metadata contains the global details for a typical client sdk
type SGen struct {
	// SGenVersion is the version of the sdk generator library.
	Version string
}

type Operation struct {
	Path  string
	Props spec.OperationProps
}

func mapOperations(specDoc *spec.Swagger) map[string][]Operation {
	operations := make(map[string][]Operation)

	for path, operation := range specDoc.Paths.Paths {
		v := reflect.ValueOf(operation.PathItemProps)

		values := make([]interface{}, v.NumField())

		// We need to loop over the available fields for PathItemProps to
		// extract the relevant information.
		for i := 0; i < v.NumField(); i++ {
			// We exclude empty operations from being rendered in the sdk
			if v.Field(i).CanInterface() && !v.Field(i).IsNil() {
				values[i] = v.Field(i).Interface()

				// Cast the reflected value to it's operation struct counterpart
				o := values[i].(*spec.Operation)

				// We will group resources using it's first tag
				t := strings.ToLower(o.OperationProps.Tags[0])

				operations[t] = append(operations[t], Operation{
					Path:  path,
					Props: o.OperationProps,
				})
			}
		}
	}

	pp.Println(operations)

	return operations
}
