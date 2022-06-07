package sdk

import (
	"fmt"

	"github.com/go-openapi/loads"
	"github.com/rs/zerolog/log"
)

// New sets up and returns an instance of the SDK generator
func New(specDoc *loads.Document, currentDir, sgenVersion string) (*ClientSDK, error) {
	// Setup the initial template data from the swagger file
	td := TemplateData{
		Spec: specDoc.OrigSpec(),
		SGen: SGen{
			Version: sgenVersion,
		},
	}

	return &ClientSDK{
		CurrentDir:   currentDir,
		TemplateData: td,
	}, nil
}

// Render renders the client api sdk from the given templates using the swagger configuration data loaded
// by viper and applying it to the templates to render a set of .go files.
func (c *ClientSDK) Render() error {
	log.Debug().Interface("template_data", c.TemplateData).Msg("rendering client sdk with")

	err := c.RenderClient()
	if err != nil {
		log.Error().AnErr("error", err).Msg("render client failed")

		return fmt.Errorf("render client: %w", err)
	}

	return nil
}
