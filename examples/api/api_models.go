package api

import (
	"mime/multipart"
	"time"
)

// UserRequest represents a request to create or update a user
type UserRequest struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

// UserResponse represents an API response with user data
type UserResponse struct {
	User_ID    int       `json:"user_id"`
	First_Name string    `json:"first_name"`
	Last_Name  string    `json:"last_name"`
	Email      string    `json:"email"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

// Address represents a physical address in API requests/responses
type Address struct {
	Street_Line1 string `json:"street_line1"`
	Street_Line2 string `json:"street_line2,omitempty"`
	City         string `json:"city"`
	State        string `json:"state"`
	Postal_Code  string `json:"postal_code"`
	Country      string `json:"country"`
}

// SearchParams represents query parameters for search endpoints
type SearchParams struct {
	Query      string   `query:"q"`
	Page       int      `query:"page"`
	Limit      int      `query:"limit"`
	Sort_By    string   `query:"sort_by"`
	Sort_Order string   `query:"sort_order"`
	Filters    []string `query:"filters"`
}

// LoginForm represents a login form submission
type LoginForm struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"pass"`
	Remember bool   `form:"remember_me" json:"remember"`
}

// RegisterForm represents a user registration form
type RegisterForm struct {
	Username        string `form:"username" json:"username" binding:"required" validate:"min=3,max=50"`
	Email           string `form:"email" json:"email" binding:"required" validate:"email"`
	Password        string `form:"password" json:"password" binding:"required" validate:"min=8,containsAny=!@#$%^&*"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required" validate:"eqfield=Password"`
	AcceptTerms     bool   `form:"accept_terms" json:"accept_terms" binding:"required" validate:"eq=true"`
}

// MixedTagsStruct demonstrates priority between JSON and form tags
type MixedTagsStruct struct {
	ID       int64  `json:"id" form:"user_id"`
	Name     string `json:"name" form:"user_name"`
	Email    string `json:"email" form:"user_email"`
	JSONOnly string `json:"json_only"`
	FormOnly string `form:"form_only"`
	NoTags   string
}

// FileUploadForm represents a form with file uploads
type FileUploadForm struct {
	UserID      int64                   `form:"user_id"`
	Title       string                  `form:"title"`
	Description string                  `form:"description"`
	File        *multipart.FileHeader   `form:"file"`
	Images      []*multipart.FileHeader `form:"images"`
}

// RouteParams represents URL parameters in a route
type RouteParams struct {
	UserID     int64  `param:"user_id" json:"id"`
	PostID     int64  `param:"post_id" json:"postId"`
	CommentID  string `param:"comment_id" json:"commentId"`
	CategoryID string `param:"category_id"`
}

// MixedParamStruct demonstrates priority between param and json tags
type MixedParamStruct struct {
	ID        int64  `json:"id" param:"user_id"`
	Name      string `json:"name" param:"user_name"`
	JSONOnly  string `json:"json_only"`
	ParamOnly string `param:"param_only"`
}

// SearchForm represents a search form with various filters
type SearchForm struct {
	Query      string   `form:"q" json:"query" validate:"omitempty,max=100"`
	Categories []string `form:"categories" json:"categories" validate:"omitempty,dive,max=50"`
	MinPrice   *float64 `form:"min_price" json:"minPrice" validate:"omitempty,min=0"`
	MaxPrice   *float64 `form:"max_price" json:"maxPrice" validate:"omitempty,gtfield=MinPrice"`
	SortBy     string   `form:"sort_by" json:"sortBy" validate:"omitempty,oneof=price date popularity rating"`
	SortOrder  string   `form:"sort_order" json:"sortOrder" validate:"omitempty,oneof=asc desc"`
	Page       int      `form:"page" json:"page" validate:"min=1"`
	Limit      int      `form:"limit" json:"limit" validate:"min=1,max=100"`
}

// NullableFieldsExample demonstrates different combinations of nullable and required fields
type NullableFieldsExample struct {
	// Basic required field
	RequiredField string `json:"required_field" validate:"required"`

	// Optional field with omitempty
	OptionalField string `json:"optional_field,omitempty"`

	// Nullable field (pointer type)
	NullableField *string `json:"nullable_field"`

	// Nullable and optional field
	NullableOptionalField *string `json:"nullable_optional_field,omitempty"`

	// Nullable but required field
	NullableRequiredField *string `json:"nullable_required_field" validate:"required"`

	// Field with binding:required
	BindingRequiredField string `json:"binding_required_field" binding:"required"`
}
