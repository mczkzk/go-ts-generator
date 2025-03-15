package api

import "time"

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
