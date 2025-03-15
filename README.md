# go-ts-generator

A Go module that generates TypeScript type definitions from Go struct definitions.

## Features

- Automatically generates TypeScript interfaces from Go structs
- Handles JSON and query tag parsing for field names
- Converts Go types to appropriate TypeScript types
- Preserves comments from Go code in the generated TypeScript
- Handles both exported and unexported types and fields
- Converts field names to camelCase for non-API types
- Special handling for API-related types (preserves exact field names)
- Handles pointer types as optional fields
- Supports `omitempty` tag for optional fields

## Installation

### As a command-line tool

```bash
go install github.com/mo49/go-ts-generator/cmd/go-ts-generator@latest
```

### As a library

```bash
go get github.com/mo49/go-ts-generator
```

## Usage

### Command-line

```bash
go-ts-generator <source_dir> <target_file>
```

Where:
- `<source_dir>` is the directory containing Go files to parse
- `<target_file>` is the target TypeScript file to generate

Example:
```bash
go-ts-generator ./models ./types/generated.ts
```

### As a library

```go
package main

import (
	"fmt"
	"github.com/mo49/go-ts-generator/pkg/generator"
)

func main() {
	err := generator.GenerateTypes("./models", "./types/generated.ts")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
```

## Type Conversion

| Go Type | TypeScript Type |
|---------|----------------|
| string | string |
| bool | boolean |
| int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64 | number |
| time.Time | string /* RFC3339 */ |
| []T | T[] |
| map[K]V | Record<K, V> |
| interface{} | any |
| *T | T? (optional) |

## Example

### Go Input

```go
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
```

### TypeScript Output

```typescript
// This file is auto-generated. Do not edit directly.
// Generated at: 2025-03-15 12:34:56
// Note: This file includes both exported and unexported types and fields.

/* eslint-disable */

/**
 * User represents a user in the system
 */
export interface User {
  id: number;
  name: string;
  email: string;
  createdAt: string /* RFC3339 */;
  updatedAt: string /* RFC3339 */;
  address?: Address;
}

/**
 * Address represents a physical address
 */
export interface Address {
  street: string;
  city: string;
  state: string;
  zipCode: string;
  country: string;
}
```

### API Type Example

For API-related types (file paths containing `controller`, `handler`, `api`, or type names containing `Request`, `Response`, `Params`), field names are preserved as-is:

```go
// UserResponse represents an API response
type UserResponse struct {
	User_ID    int       `json:"user_id"`
	First_Name string    `json:"first_name"`
	Last_Name  string    `json:"last_name"`
	Created_At time.Time `json:"created_at"`
}
```

```typescript
/**
 * UserResponse represents an API response
 */
export interface UserResponse {
  user_id: number;
  first_name: string;
  last_name: string;
  created_at: string /* RFC3339 */;
}
```

## License

MIT