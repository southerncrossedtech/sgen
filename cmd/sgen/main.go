package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/southerncrossedtech/sgen/cmd/sdk"
	"github.com/spf13/cobra"
)

const sgenVersion = "0.0.1"

func main() {
	// Handle simple flags immediately.
	for _, arg := range os.Args {
		if arg == "--version" || arg == "-v" {
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
		Use:   "sgen [command] [flags]",
		Short: "SDK from Swagger Generator.",
		Long: "sgen generates a golang client sdk based on the swagger file input for use as an importable package.\n\n" +
			"Complete documentation is available at http://github.com/southercrossedtech/sgen",
		Example: `sgen sdk -f docs/swagger.yaml`,
	}

	rootCmd.PersistentFlags().BoolP("version", "v", false, "print the version")

	// Setup available sub commands
	rootCmd.AddCommand(sdk.Command(sgenVersion))

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n\n", err.Error())

		os.Exit(1)
	}
}
