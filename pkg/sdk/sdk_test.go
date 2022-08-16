package sdk

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-openapi/loads"
	"github.com/stretchr/testify/require"
)

const testVersion = "0.0.1-test"

func Test_New(t *testing.T) {
	// Set path for config to current directory
	currentDir, err := os.Getwd()
	require.NoError(t, err)

	specDoc, err := loads.Spec(fmt.Sprintf("%s/%s", currentDir, "testdata/swagger.yaml"))
	require.NoError(t, err)

	client, err := New(specDoc, currentDir, testVersion)
	require.NoError(t, err)

	require.Equal(t, "httpbin", client.Spec.Info.Title)
	require.Equal(t, "v1.0", client.Spec.Info.Version)
	require.Equal(t, testVersion, client.SGen.Version)
}
