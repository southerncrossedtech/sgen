package sdk

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

const testVersion = "0.0.1-test"

func Test_New(t *testing.T) {
	// Set path for config to current directory
	currentDir, err := os.Getwd()
	require.NoError(t, err)

	viper.AddConfigPath(currentDir)              // Set current directory
	viper.SetConfigName("testdata/swagger.yaml") // Set file path in current directory
	viper.SetConfigType("yaml")                  // Set file type

	// Find and read the config file
	err = viper.ReadInConfig()
	require.NoError(t, err)

	client := New(viper.GetViper(), currentDir, testVersion)

	require.Equal(t, "httpbin", client.Metadata.Title)
	require.Equal(t, "v1.0", client.Metadata.Version)
	require.Equal(t, testVersion, client.Metadata.SGenVersion)
}
