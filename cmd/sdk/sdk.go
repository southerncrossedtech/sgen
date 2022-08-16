package sdk

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-openapi/loads"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/southerncrossedtech/sgen/pkg/config"
	"github.com/southerncrossedtech/sgen/pkg/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var sgenVersion string

// Command returns the "mock" command.
func Command(v string) *cobra.Command {
	sgenVersion = v

	sdkCmd := &cobra.Command{
		Use:   "sdk",
		Short: "SDK from Swagger Generator.",
		Long: "sgen sdk generates a golang client sdk based on the swagger file input for use as an importable package.\n" +
			`Complete documentation is available at http://github.com/southercrossedtech/sgen`,
		Example: `sgen sdk -f docs/swagger.yaml`,
		RunE:    run,
	}

	// Setup command flags
	sdkCmd.PersistentFlags().BoolP("debug", "d", false, "debug mode prints stack traces on error")
	sdkCmd.PersistentFlags().StringP("file", "f", "docs/swagger.yaml", "the name of the swagger input file, defaults to github.com/swaggo/swag default output")
	sdkCmd.PersistentFlags().StringP("output", "o", "output", "output directory for all the generated files, defaults to ./output/*.go")

	viper.BindPFlags(sdkCmd.PersistentFlags())
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// cmd.AddCommand(ProbeCommand())

	return sdkCmd
}

func run(cmd *cobra.Command, args []string) error {
	conf := config.Config{
		File:         viper.GetString("file"),
		Output:       viper.GetString("output"),
		DebugEnabled: viper.GetBool("debug"),
		Version:      sgenVersion,
	}

	log.Info().Str("file", conf.File).Msg("running sgen")
	log.Debug().Interface("cmdConfig", conf).Msg("using sgen with config")

	if conf.DebugEnabled {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// Set path for config to current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Check if output directory exists, otherwise create it
	if _, err := os.Stat(conf.Output); os.IsNotExist(err) {
		err := os.Mkdir(conf.Output, os.ModePerm)
		if err != nil {
			log.Error().AnErr("error", err).Msg("failed to create output directory")

			return err
		}
	}

	viper.AddConfigPath(currentDir) // Set current directory
	viper.SetConfigName(conf.File)  // Set file path in current directory
	viper.SetConfigType("yaml")     // Set file type

	// Find and read the config file
	err = viper.ReadInConfig()
	if err != nil {
		log.Error().AnErr("error", err).Msg("failed to read in swagger file")

		return err
	}

	// Load the spec into the open api spec document model
	specDoc, err := loads.Spec(fmt.Sprintf("%s/%s", currentDir, conf.File))
	if err != nil {
		log.Error().AnErr("error", err).Msg("error loading spec doc")

		return err
	}

	// Get the instance of a new sdk generator
	clientSDK, err := sdk.New(specDoc, currentDir, conf)
	if err != nil {
		log.Error().AnErr("error", err).Msg("new sdk error")

		return err
	}

	// Render the api client sdk using the viper loaded swagger file as config
	err = clientSDK.Render()
	if err != nil {
		log.Error().AnErr("error", err).Msg("render client sdk error")

		return err
	}

	log.Info().Msg("sgen completed")

	return nil
}
