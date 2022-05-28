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
		fmt.Printf("Error: %v\n\n", err)
		rootCmd.Help()

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

	if conf.DebugEnabled {
		log.Debug().Interface("cmdConfig", conf).Msg("using sgen with config")
	}

	return nil
}
