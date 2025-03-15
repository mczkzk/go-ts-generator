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
	Username        string `form:"username" json:"username"`
	Email           string `form:"email" json:"email"`
	Password        string `form:"password" json:"password"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password"`
	AcceptTerms     bool   `form:"accept_terms" json:"accept_terms"`
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
