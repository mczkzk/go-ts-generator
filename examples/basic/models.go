package models

import "time"

// User represents a user in the system
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Address   *Address  `json:"address,omitempty"`
}

// Address represents a physical address
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}

// Product represents a product in the catalog
type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Category    *Category `json:"category"` // Pointer type without omitempty
	CreatedAt   time.Time `json:"created_at"`
}

// Category represents a product category
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// unexportedType is not exported
type unexportedType struct {
	field1 string `json:"field1"`
	field2 int    `json:"field2"`
}

// UserList represents a list of users
type UserList []*User

// StringArray is a simple string array
type StringArray []string

// CategoryMap is a map of category IDs to categories
type CategoryMap map[int]*Category
