package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateTypes(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "go-ts-generator-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test Go file
	goFilePath := filepath.Join(tempDir, "test_types.go")
	goFileContent := `package test

import "time"

// User represents a user in the system
type User struct {
	ID        int       ` + "`json:\"id\"`" + `
	Name      string    ` + "`json:\"name\"`" + `
	Email     string    ` + "`json:\"email\"`" + `
	CreatedAt time.Time ` + "`json:\"created_at\"`" + `
	UpdatedAt time.Time ` + "`json:\"updated_at\"`" + `
	Address   *Address  ` + "`json:\"address,omitempty\"`" + `
}

// Address represents a physical address
type Address struct {
	Street  string ` + "`json:\"street\"`" + `
	City    string ` + "`json:\"city\"`" + `
	State   string ` + "`json:\"state\"`" + `
	ZipCode string ` + "`json:\"zip_code\"`" + `
	Country string ` + "`json:\"country\"`" + `
}

// unexportedType is not exported
type unexportedType struct {
	field1 string ` + "`json:\"field1\"`" + `
	field2 int    ` + "`json:\"field2\"`" + `
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

	// Output the generated file content for debugging
	fmt.Printf("Generated TypeScript content:\n%s\n", tsContentStr)

	// Check for User interface
	if !strings.Contains(tsContentStr, "export interface User {") {
		t.Error("Generated TypeScript does not contain User interface")
	}

	// Check for Address interface
	if !strings.Contains(tsContentStr, "export interface Address {") {
		t.Error("Generated TypeScript does not contain Address interface")
	}

	// Check for unexportedType (should be converted to PascalCase)
	if !strings.Contains(tsContentStr, "export interface UnexportedType {") {
		t.Error("Generated TypeScript does not contain UnexportedType interface")
	}

	// Check for optional field (address in User)
	if !strings.Contains(tsContentStr, "address?: Address;") {
		t.Error("Generated TypeScript does not handle optional fields correctly")
	}

	// Check for time.Time conversion - field names are converted to camelCase
	if !strings.Contains(tsContentStr, "createdAt: string /* RFC3339 */;") {
		t.Error("Generated TypeScript does not handle time.Time correctly")
		// Output specific content when the test fails
		fmt.Printf("Looking for 'createdAt: string /* RFC3339 */;' in:\n%s\n", tsContentStr)
	}
}
