package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestGenerateNullableTypes tests the generation of TypeScript types with nullable fields
func TestGenerateNullableTypes(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "go-ts-generator-nullable-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test Go file with nullable fields
	goFilePath := filepath.Join(tempDir, "nullable_models.go")
	goFileContent := `package nullable

// BasicField represents a struct with a basic required field
type BasicField struct {
	Name string ` + "`json:\"name\"`" + `
}

// OptionalField represents a struct with an optional field (omitempty)
type OptionalField struct {
	Name string ` + "`json:\"name,omitempty\"`" + `
}

// NullableField represents a struct with a nullable field (pointer)
type NullableField struct {
	Name *string ` + "`json:\"name\"`" + `
}

// NullableOptionalField represents a struct with a nullable and optional field
type NullableOptionalField struct {
	Name *string ` + "`json:\"name,omitempty\"`" + `
}

// RequiredField represents a struct with an explicitly required field
type RequiredField struct {
	Name string ` + "`json:\"name\" validate:\"required\"`" + `
}

// NullableRequiredField represents a struct with a nullable but required field
type NullableRequiredField struct {
	Name *string ` + "`json:\"name\" validate:\"required\"`" + `
}

// MixedFields represents a struct with various field types
type MixedFields struct {
	Required        string  ` + "`json:\"required\" validate:\"required\"`" + `
	Optional        string  ` + "`json:\"optional,omitempty\"`" + `
	Nullable        *string ` + "`json:\"nullable\"`" + `
	NullableOpt     *string ` + "`json:\"nullable_opt,omitempty\"`" + `
	NullableReq     *string ` + "`json:\"nullable_req\" validate:\"required\"`" + `
	BindingRequired string  ` + "`json:\"binding_required\" binding:\"required\"`" + `
}
`

	if err := os.WriteFile(goFilePath, []byte(goFileContent), 0644); err != nil {
		t.Fatalf("Failed to write test Go file: %v", err)
	}

	// Generate TypeScript types
	tsFilePath := filepath.Join(tempDir, "generated.ts")
	if err := GenerateTypes(tempDir, tsFilePath); err != nil {
		t.Fatalf("GenerateTypes failed: %v", err)
	}

	// Read the generated TypeScript file
	tsContent, err := os.ReadFile(tsFilePath)
	if err != nil {
		t.Fatalf("Failed to read generated TypeScript file: %v", err)
	}

	// Check for expected content
	tsContentStr := string(tsContent)

	// Basic required field
	if !strings.Contains(tsContentStr, "name: string;") {
		t.Error("Generated TypeScript does not handle basic required field correctly")
		t.Logf("Expected 'name: string;' but got something else in:\n%s", tsContentStr)
	}

	// Optional field (omitempty)
	if !strings.Contains(tsContentStr, "name?: string;") {
		t.Error("Generated TypeScript does not handle optional field correctly")
		t.Logf("Expected 'name?: string;' but got something else in:\n%s", tsContentStr)
	}

	// Nullable field (pointer)
	if !strings.Contains(tsContentStr, "name: string | null;") {
		t.Error("Generated TypeScript does not handle nullable field correctly")
		t.Logf("Expected 'name: string | null;' but got something else in:\n%s", tsContentStr)
	}

	// Nullable and optional field
	if !strings.Contains(tsContentStr, "name?: string | null;") {
		t.Error("Generated TypeScript does not handle nullable and optional field correctly")
		t.Logf("Expected 'name?: string | null;' but got something else in:\n%s", tsContentStr)
	}

	// Explicitly required field
	if !strings.Contains(tsContentStr, "required: string;") {
		t.Error("Generated TypeScript does not handle explicitly required field correctly")
		t.Logf("Expected 'required: string;' but got something else in:\n%s", tsContentStr)
	}

	// Nullable but required field
	if !strings.Contains(tsContentStr, "nullable_req: string | null;") {
		t.Error("Generated TypeScript does not handle nullable but required field correctly")
		t.Logf("Expected 'nullable_req: string | null;' but got something else in:\n%s", tsContentStr)
	}

	// Field with binding:required
	if !strings.Contains(tsContentStr, "binding_required: string;") {
		t.Error("Generated TypeScript does not handle binding:required field correctly")
		t.Logf("Expected 'binding_required: string;' but got something else in:\n%s", tsContentStr)
	}
}
