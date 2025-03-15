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
- Prioritizes tags in the following order: `json` > `form` > `param` > `query`
- Handles arrays of pointers as `(Type | null | undefined)[]` in TypeScript
- Includes validation rules from struct tags as JSDoc comments

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

### Examples

Check out the [examples](./examples) directory for complete usage examples:
- [Basic examples](./examples/basic) - Standard Go structs and their TypeScript equivalents
- [API examples](./examples/api) - API-related types with preserved field names

## Tag Priority

The generator prioritizes field tags in the following order:

1. `json` tag (highest priority) - Used for JSON field names in API communication
2. `form` tag - Used for HTML form field names
3. `param` tag - Used for URL parameter names
4. `query` tag (lowest priority) - Used for query parameter names

This means that if multiple tags are present on a field, the generator will use the name from the highest priority tag available.

Example:
```go
type MixedTagsStruct struct {
    ID    int64  `json:"id" form:"user_id"`
    Name  string `form:"user_name" param:"name"`
    Email string `param:"email" query:"user_email"`
}
```

Will generate:
```typescript
export interface MixedTagsStruct {
  id: number;         // from json tag
  user_name: string;  // from form tag
  email: string;      // from param tag
}
```

## Validation Rules

The generator extracts validation rules from struct tags and includes them as JSDoc comments in the generated TypeScript. This allows frontend developers to reference the same validation rules that are enforced on the backend.

Example:
```go
type RegisterForm struct {
    Username string `form:"username" binding:"required" validate:"min=3,max=50"`
    Email    string `form:"email" binding:"required" validate:"email"`
}
```

Will generate:
```typescript
export interface RegisterForm {
  /**
   * @validation
   *   - binding: required
   *   - validate: min=3,max=50
   */
  username: string;
  /**
   * @validation
   *   - binding: required
   *   - validate: email
   */
  email: string;
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
| []*T | (T \| null \| undefined)[] |
| map[K]V | Record<K, V> |
| interface{} | any |
| *T | T? (optional) |

## Quick Example

### Go Input

```go
// User represents a user in the system
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Address   *Address  `json:"address,omitempty"`
}

// Address represents a physical address
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// unexportedType is not exported
type unexportedType struct {
	publicField  string `json:"publicField"`
	privateField int    `json:"privateField"`
}
```

### TypeScript Output

```typescript
/**
 * User represents a user in the system
 */
export interface User {
  id: number;
  name: string;
  email: string;
  createdAt: string /* RFC3339 */;
  address?: Address;
}

/**
 * Address represents a physical address
 */
export interface Address {
  street: string;
  city: string;
  country: string;
}

/**
 * unexportedType is not exported
 */
/**
 * Note: This is an unexported type. In Go code, it's defined with a lowercase identifier.
 * It cannot be accessed directly from outside the package.
 */
export interface UnexportedType {
  publicField: string;
  /**
   * Note: This is an unexported field. In Go code, it's defined with a lowercase identifier.
   * It cannot be accessed directly from outside the package.
   */
  privateField: number;
}
```

For more detailed examples including:
- Pointer types without `omitempty`
- API-related types with preserved field names
- Query tag handling
- Unexported types and fields

Please see the [examples](./examples) directory.

## License

MIT