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

// TestGenerateFormTagTypes tests the generation of TypeScript types from Go structs with form tags
func TestGenerateFormTagTypes(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "go-ts-generator-form-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test Go file with form tags
	goFilePath := filepath.Join(tempDir, "form_models.go")
	goFileContent := `package form

import "mime/multipart"

// LoginForm represents a login form submission
type LoginForm struct {
	Username string ` + "`form:\"username\" json:\"user\"`" + `
	Password string ` + "`form:\"password\" json:\"pass\"`" + `
	Remember bool   ` + "`form:\"remember_me\" json:\"remember\"`" + `
}

// RegisterForm represents a user registration form
type RegisterForm struct {
	Username        string ` + "`form:\"username\" json:\"username\"`" + `
	Email           string ` + "`form:\"email\" json:\"email\"`" + `
	Password        string ` + "`form:\"password\" json:\"password\"`" + `
	ConfirmPassword string ` + "`form:\"confirm_password\" json:\"confirm_password\"`" + `
	AcceptTerms     bool   ` + "`form:\"accept_terms\" json:\"accept_terms\"`" + `
}

// MixedTagsStruct demonstrates priority between JSON and form tags
type MixedTagsStruct struct {
	ID       int64  ` + "`json:\"id\" form:\"user_id\"`" + `
	Name     string ` + "`json:\"name\" form:\"user_name\"`" + `
	Email    string ` + "`json:\"email\" form:\"user_email\"`" + `
	JSONOnly string ` + "`json:\"json_only\"`" + `
	FormOnly string ` + "`form:\"form_only\"`" + `
	NoTags   string
}

// FileUploadForm represents a form with file uploads
type FileUploadForm struct {
	UserID      int64                 ` + "`form:\"user_id\"`" + `
	Title       string                ` + "`form:\"title\"`" + `
	Description string                ` + "`form:\"description\"`" + `
	File        *multipart.FileHeader ` + "`form:\"file\"`" + `
	Images      []*multipart.FileHeader ` + "`form:\"images\"`" + `
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

	// Check for form tag usage in LoginForm
	if !strings.Contains(tsContentStr, "username: string;") {
		t.Error("Generated TypeScript does not use form tag for field name in LoginForm")
		t.Logf("Expected 'username: string;' but got something else in:\n%s", tsContentStr)
	}

	if !strings.Contains(tsContentStr, "password: string;") {
		t.Error("Generated TypeScript does not use form tag for field name in LoginForm")
		t.Logf("Expected 'password: string;' but got something else in:\n%s", tsContentStr)
	}

	if !strings.Contains(tsContentStr, "remember_me: boolean;") {
		t.Error("Generated TypeScript does not use form tag for field name in LoginForm")
		t.Logf("Expected 'remember_me: boolean;' but got something else in:\n%s", tsContentStr)
	}

	// Check for form tag usage in RegisterForm
	if !strings.Contains(tsContentStr, "confirm_password: string;") {
		t.Error("Generated TypeScript does not use form tag for field name in RegisterForm")
		t.Logf("Expected 'confirm_password: string;' but got something else in:\n%s", tsContentStr)
	}

	// Check for form tag priority in MixedTagsStruct
	if !strings.Contains(tsContentStr, "user_id: number;") {
		t.Error("Generated TypeScript does not prioritize form tag over json tag")
		t.Logf("Expected 'user_id: number;' but got something else in:\n%s", tsContentStr)
	}

	if !strings.Contains(tsContentStr, "user_name: string;") {
		t.Error("Generated TypeScript does not prioritize form tag over json tag")
		t.Logf("Expected 'user_name: string;' but got something else in:\n%s", tsContentStr)
	}

	// Check for json-only field
	if !strings.Contains(tsContentStr, "json_only: string;") {
		t.Error("Generated TypeScript does not handle json-only tag correctly")
		t.Logf("Expected 'json_only: string;' but got something else in:\n%s", tsContentStr)
	}

	// Check for form-only field
	if !strings.Contains(tsContentStr, "form_only: string;") {
		t.Error("Generated TypeScript does not handle form-only tag correctly")
		t.Logf("Expected 'form_only: string;' but got something else in:\n%s", tsContentStr)
	}

	// Check for FileUploadForm fields
	if !strings.Contains(tsContentStr, "file?: FileHeader;") {
		t.Error("Generated TypeScript does not handle file upload field correctly")
		t.Logf("Expected 'file?: FileHeader;' but got something else in:\n%s", tsContentStr)
	}

	if !strings.Contains(tsContentStr, "images?: (FileHeader | null | undefined)[];") {
		t.Error("Generated TypeScript does not handle multiple file upload field correctly")
		t.Logf("Expected 'images?: (FileHeader | null | undefined)[];' but got something else in:\n%s", tsContentStr)
	}
}

// TestGenerateParamTagTypes tests the generation of TypeScript types from Go structs with param tags
func TestGenerateParamTagTypes(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "go-ts-generator-param-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a test Go file with param tags
	goFilePath := filepath.Join(tempDir, "param_models.go")
	goFileContent := `package param

// RouteParams represents URL parameters in a route
type RouteParams struct {
	UserID     int64  ` + "`param:\"user_id\" json:\"id\"`" + `
	PostID     int64  ` + "`param:\"post_id\" json:\"postId\"`" + `
	CommentID  string ` + "`param:\"comment_id\" json:\"commentId\"`" + `
	CategoryID string ` + "`param:\"category_id\"`" + `
}

// MixedParamStruct demonstrates priority between param and json tags
type MixedParamStruct struct {
	ID        int64  ` + "`json:\"id\" param:\"user_id\"`" + `
	Name      string ` + "`json:\"name\" param:\"user_name\"`" + `
	JSONOnly  string ` + "`json:\"json_only\"`" + `
	ParamOnly string ` + "`param:\"param_only\"`" + `
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

	// Check for param tag usage in RouteParams
	if !strings.Contains(tsContentStr, "user_id: number;") {
		t.Error("Generated TypeScript does not use param tag for field name in RouteParams")
		t.Logf("Expected 'user_id: number;' but got something else in:\n%s", tsContentStr)
	}

	if !strings.Contains(tsContentStr, "post_id: number;") {
		t.Error("Generated TypeScript does not use param tag for field name in RouteParams")
		t.Logf("Expected 'post_id: number;' but got something else in:\n%s", tsContentStr)
	}

	if !strings.Contains(tsContentStr, "comment_id: string;") {
		t.Error("Generated TypeScript does not use param tag for field name in RouteParams")
		t.Logf("Expected 'comment_id: string;' but got something else in:\n%s", tsContentStr)
	}

	// Check for param tag priority in MixedParamStruct
	if !strings.Contains(tsContentStr, "user_id: number;") {
		t.Error("Generated TypeScript does not prioritize param tag over json tag")
		t.Logf("Expected 'user_id: number;' but got something else in:\n%s", tsContentStr)
	}

	if !strings.Contains(tsContentStr, "user_name: string;") {
		t.Error("Generated TypeScript does not prioritize param tag over json tag")
		t.Logf("Expected 'user_name: string;' but got something else in:\n%s", tsContentStr)
	}

	// Check for json-only field
	if !strings.Contains(tsContentStr, "json_only: string;") {
		t.Error("Generated TypeScript does not handle json-only tag correctly")
		t.Logf("Expected 'json_only: string;' but got something else in:\n%s", tsContentStr)
	}

	// Check for param-only field
	if !strings.Contains(tsContentStr, "param_only: string;") {
		t.Error("Generated TypeScript does not handle param-only tag correctly")
		t.Logf("Expected 'param_only: string;' but got something else in:\n%s", tsContentStr)
	}
}
