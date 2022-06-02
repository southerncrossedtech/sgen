package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/southerncrossedtech/sgen/pkg/config"
	"github.com/southerncrossedtech/sgen/pkg/sdk"
)

const sgenVersion = "0.0.1"

func main() {
	// Handle simple flags immediately.
	for _, arg := range os.Args {
		if arg == "--version" {
			fmt.Println("sgen v" + sgenVersion)

			return
		}
	}

	// Setup zerolog logger to stdout
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Setup global log level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Set up the cobra root command
	rootCmd := &cobra.Command{
		Use:   "sgen [flags]",
		Short: "SDK from Swagger Generator.",
		Long: "sgen generates a golang client sdk based on the swagger file input for use as an importable package.\n" +
			`Complete documentation is available at http://github.com/southercrossedtech/sgen`,
		Example: `sgen -f docs/swagger.yaml`,
		RunE:    run,
	}

	// Setup command flags
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "debug mode prints stack traces on error")
	rootCmd.PersistentFlags().StringP("file", "f", "docs/swagger.yaml", "the name of the swagger input file, defaults to github.com/swaggo/swag default output")
	rootCmd.PersistentFlags().StringP("output", "o", "sdk/*.go", "output directory for all the generated files")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "print the version")

	viper.BindPFlags(rootCmd.PersistentFlags())
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n\n", err.Error())

		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	conf := config.Config{
		File:         viper.GetString("file"),
		Output:       viper.GetString("output"),
		DebugEnabled: viper.GetBool("debug"),
		Version:      viper.GetBool("version"),
	}

	if conf.Version {
		fmt.Println("sgen v" + sgenVersion)

		return nil
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

	viper.AddConfigPath(currentDir) // Set current directory
	viper.SetConfigName(conf.File)  // Set file path in current directory
	viper.SetConfigType("yaml")     // Set file type

	// Find and read the config file
	err = viper.ReadInConfig()
	if err != nil {
		log.Error().AnErr("error", err).Msg("failed to read in swagger file")

		return err
	}

	// Get the instance of a new sdk generator
	clientSDK := sdk.New(viper.GetViper(), currentDir, sgenVersion)

	// Render the api client sdk using the viper loaded swagger file as config
	err = clientSDK.Render()
	if err != nil {
		log.Error().AnErr("error", err).Msg("render client sdk error")

		return err
	}

	log.Info().Msg("sgen completed")

	return nil
}
