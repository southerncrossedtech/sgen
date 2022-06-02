package sdk

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/southerncrossedtech/sgen/resources"
	"github.com/spf13/viper"
)

// New sets up and returns an instance of the SDK generator
func New(viper *viper.Viper, currentDir, sgenVersion string) *ClientSDK {
	// Setup the initial template data from the swagger file
	td := TemplateData{
		Metadata: Metadata{
			Title:       strings.ToLower(viper.GetString("info.title")),
			Version:     viper.GetString("info.version"),
			SGenVersion: sgenVersion,
		},
	}

	return &ClientSDK{
		CurrentDir:   currentDir,
		TemplateData: td,
	}
}

// Render renders the client api sdk from the given templates using the swagger configuration data loaded
// by viper and applying it to the templates to render a set of .go files.
func (c *ClientSDK) Render() error {
	log.Debug().Interface("template_data", c.TemplateData).Msg("rendering client sdk with")

	clientBytes, err := resources.Templates.ReadFile("templates/00_client.go.tpl")
	if err != nil {
		log.Error().AnErr("error", err).Msg("read template failed")

		return err
	}

	clientTpl, err := template.New("client").Parse(string(clientBytes))
	if err != nil {
		log.Error().AnErr("error", err).Msg("parse template failed")

		return err
	}

	outputFile, err := os.Create(fmt.Sprintf("%s/%s/%s", c.CurrentDir, "output", "client.go"))
	if err != nil {
		log.Error().AnErr("error", err).Msg("failed to create output file")

		return err
	}

	err = clientTpl.Execute(outputFile, c.TemplateData)
	if err != nil {
		log.Error().AnErr("error", err).Msg("execute template failed")

		return err
	}

	return nil
}
