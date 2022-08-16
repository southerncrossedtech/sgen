package sdk

import (
	"fmt"
	"html/template"
	"os"

	"github.com/Masterminds/sprig/v3"
	"github.com/rs/zerolog/log"
	"github.com/southerncrossedtech/sgen/resources"
)

// Render renders the client api sdk from the given templates using the swagger configuration data loaded
// by viper and applying it to the templates to render a set of .go files.
func (c *ClientSDK) RenderClient() error {
	clientBytes, err := resources.Templates.ReadFile("templates/00_client.go.tpl")
	if err != nil {
		log.Error().AnErr("error", err).Msg("read template failed")

		return fmt.Errorf("read template failed: %w", err)
	}

	clientTpl, err := template.New("client").Funcs(sprig.FuncMap()).Parse(string(clientBytes))
	if err != nil {
		log.Error().AnErr("error", err).Msg("parse template failed")

		return fmt.Errorf("parse template failed: %w", err)
	}

	outputFile, err := os.Create(fmt.Sprintf("%s/%s/%s", c.CurrentDir, "output", "client.go"))
	if err != nil {
		log.Error().AnErr("error", err).Msg("failed to create output file")

		return fmt.Errorf("failed to create output file: %w", err)
	}

	err = clientTpl.Execute(outputFile, c.TemplateData)
	if err != nil {
		log.Error().AnErr("error", err).Msg("execute template failed")

		return fmt.Errorf("execute template failed: %w", err)
	}

	return nil
}
