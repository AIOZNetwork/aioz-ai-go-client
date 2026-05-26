package contract

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type SwaggerSpec struct {
	Paths       map[string]map[string]any `json:"paths"`
	Definitions map[string]any            `json:"definitions"`
}

func loadSpec(t *testing.T) *SwaggerSpec {
	t.Helper()

	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(filename), "..", "..", "..", "..")
	specPath := filepath.Join(projectRoot, "swagger", "sdk.json")

	data, err := os.ReadFile(specPath)
	require.NoError(t, err, "Failed to read swagger spec at %s", specPath)

	var spec SwaggerSpec
	require.NoError(t, json.Unmarshal(data, &spec))
	return &spec
}

func TestSpecDefinitionCount(t *testing.T) {
	spec := loadSpec(t)
	assert.GreaterOrEqual(t, len(spec.Definitions), 500,
		"Swagger spec should have at least 500 definitions, got %d", len(spec.Definitions))
}

func TestSpecEndpointCount(t *testing.T) {
	spec := loadSpec(t)
	endpointCount := 0
	for _, methods := range spec.Paths {
		for method := range methods {
			m := strings.ToUpper(method)
			if m == "GET" || m == "POST" || m == "PUT" || m == "DELETE" || m == "PATCH" {
				endpointCount++
			}
		}
	}
	assert.GreaterOrEqual(t, endpointCount, 200,
		"Swagger spec should have at least 200 endpoints, got %d", endpointCount)
}

func TestGeneratedModelsExist(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(filename), "..", "..")
	modelsDir := filepath.Join(projectRoot, "generated", "models")

	entries, err := os.ReadDir(modelsDir)
	require.NoError(t, err, "Generated models directory should exist")

	goFiles := 0
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".go") {
			goFiles++
		}
	}
	assert.GreaterOrEqual(t, goFiles, 500,
		"Should have at least 500 generated model files, got %d", goFiles)
}

func TestGeneratedClientExists(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(filename), "..", "..")
	clientDir := filepath.Join(projectRoot, "generated", "client")

	entries, err := os.ReadDir(clientDir)
	require.NoError(t, err, "Generated client directory should exist")

	// Count sub-directories (operation packages)
	dirs := 0
	for _, e := range entries {
		if e.IsDir() {
			dirs++
		}
	}
	assert.GreaterOrEqual(t, dirs, 30,
		"Should have at least 30 generated operation packages, got %d", dirs)
}
