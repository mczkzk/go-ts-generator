package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateBasicTypes(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "go-ts-generator-basic-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test Go file
	goFilePath := filepath.Join(tempDir, "basic_models.go")
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

// Product represents a product in the catalog
type Product struct {
	ID          int       ` + "`json:\"id\"`" + `
	Name        string    ` + "`json:\"name\"`" + `
	Description string    ` + "`json:\"description\"`" + `
	Price       float64   ` + "`json:\"price\"`" + `
	Category    *Category ` + "`json:\"category\"`" + ` // Pointer type without omitempty
	CreatedAt   time.Time ` + "`json:\"created_at\"`" + `
}

// Category represents a product category
type Category struct {
	ID   int    ` + "`json:\"id\"`" + `
	Name string ` + "`json:\"name\"`" + `
}

// unexportedType is not exported
type unexportedType struct {
	field1 string ` + "`json:\"field1\"`" + `
	field2 int    ` + "`json:\"field2\"`" + `
}

// UserList represents a list of users
type UserList []*User

// StringArray is a simple string array
type StringArray []string

// CategoryMap is a map of category IDs to categories
type CategoryMap map[int]*Category
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

	// Check for basic struct conversion
	if !strings.Contains(tsContentStr, "export interface User {") {
		t.Error("Generated TypeScript does not contain User interface")
	}

	if !strings.Contains(tsContentStr, "export interface Address {") {
		t.Error("Generated TypeScript does not contain Address interface")
	}

	// Check for unexported type handling
	if !strings.Contains(tsContentStr, "export interface unexportedType {") {
		t.Error("Generated TypeScript does not keep unexported types in lowercase")
	}

	// Check for pointer field handling
	if !strings.Contains(tsContentStr, "address?: Address | null;") {
		t.Error("Generated TypeScript does not handle pointer fields with omitempty as optional and nullable")
		t.Logf("Expected 'address?: Address | null;' but got something else in:\n%s", tsContentStr)
	}

	// Check for time.Time conversion and original field names from JSON tags
	if !strings.Contains(tsContentStr, "created_at: string /* RFC3339 */;") {
		t.Error("Generated TypeScript does not handle time.Time correctly or preserve original field names from JSON tags")
		fmt.Printf("Looking for 'created_at: string /* RFC3339 */;' in:\n%s\n", tsContentStr)
	}

	// Check for array of pointers type alias
	if !strings.Contains(tsContentStr, "export type UserList = (User | null)[];") {
		t.Error("Generated TypeScript does not handle array of pointers type alias correctly")
		t.Logf("Expected 'export type UserList = (User | null)[];' but got something else in:\n%s", tsContentStr)
	}

	// Check for string array type alias
	if !strings.Contains(tsContentStr, "export type StringArray = string[];") {
		t.Error("Generated TypeScript does not handle string array type alias correctly")
		t.Logf("Expected 'export type StringArray = string[];' but got something else in:\n%s", tsContentStr)
	}

	// Check for map type alias
	if !strings.Contains(tsContentStr, "export type CategoryMap = Record<number, Category | null>;") {
		t.Error("Generated TypeScript does not handle map type alias correctly")
		t.Logf("Expected 'export type CategoryMap = Record<number, Category | null>;' but got something else in:\n%s", tsContentStr)
	}
}
