package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateAPITypes(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "go-ts-generator-api-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test Go file with API types
	goFilePath := filepath.Join(tempDir, "api_models.go")
	goFileContent := `package api

import "time"

// UserRequest represents a request to create or update a user
type UserRequest struct {
	Name    string  ` + "`json:\"name\"`" + `
	Email   string  ` + "`json:\"email\"`" + `
	Address Address ` + "`json:\"address\"`" + `
}

// UserResponse represents an API response with user data
type UserResponse struct {
	User_ID    int       ` + "`json:\"user_id\"`" + `
	First_Name string    ` + "`json:\"first_name\"`" + `
	Last_Name  string    ` + "`json:\"last_name\"`" + `
	Email      string    ` + "`json:\"email\"`" + `
	Created_At time.Time ` + "`json:\"created_at\"`" + `
	Updated_At time.Time ` + "`json:\"updated_at\"`" + `
}

// Address represents a physical address in API requests/responses
type Address struct {
	Street_Line1 string ` + "`json:\"street_line1\"`" + `
	Street_Line2 string ` + "`json:\"street_line2,omitempty\"`" + `
	City         string ` + "`json:\"city\"`" + `
	State        string ` + "`json:\"state\"`" + `
	Postal_Code  string ` + "`json:\"postal_code\"`" + `
	Country      string ` + "`json:\"country\"`" + `
}

// SearchParams represents query parameters for search endpoints
type SearchParams struct {
	Query      string   ` + "`query:\"q\"`" + `
	Page       int      ` + "`query:\"page\"`" + `
	Limit      int      ` + "`query:\"limit\"`" + `
	Sort_By    string   ` + "`query:\"sort_by\"`" + `
	Sort_Order string   ` + "`query:\"sort_order\"`" + `
	Filters    []string ` + "`query:\"filters\"`" + `
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

	// Check for API interfaces
	if !strings.Contains(tsContentStr, "export interface UserRequest {") {
		t.Error("Generated TypeScript does not contain UserRequest interface")
	}

	if !strings.Contains(tsContentStr, "export interface UserResponse {") {
		t.Error("Generated TypeScript does not contain UserResponse interface")
	}

	// Check for snake_case field names in API types (should be preserved, not converted to camelCase)
	if !strings.Contains(tsContentStr, "user_id: number;") {
		t.Error("Generated TypeScript does not preserve snake_case field names in API types")
		t.Logf("Expected 'user_id: number;' but got something else in:\n%s", tsContentStr)
	}

	if !strings.Contains(tsContentStr, "first_name: string;") {
		t.Error("Generated TypeScript does not preserve snake_case field names in API types")
		t.Logf("Expected 'first_name: string;' but got something else in:\n%s", tsContentStr)
	}

	// Check for query tag handling
	if !strings.Contains(tsContentStr, "q: string;") {
		t.Error("Generated TypeScript does not handle query tags correctly")
		t.Logf("Expected 'q: string;' but got something else in:\n%s", tsContentStr)
	}

	// Check for array type in API
	if !strings.Contains(tsContentStr, "filters: string[];") {
		t.Error("Generated TypeScript does not handle array types in API correctly")
		t.Logf("Expected 'filters: string[];' but got something else in:\n%s", tsContentStr)
	}

	// Check for optional field with omitempty
	if !strings.Contains(tsContentStr, "street_line2?: string;") {
		t.Error("Generated TypeScript does not handle omitempty tag correctly")
		t.Logf("Expected 'street_line2?: string;' but got something else in:\n%s", tsContentStr)
	}
}
