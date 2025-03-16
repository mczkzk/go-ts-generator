# go-ts-generator

A Go module that generates TypeScript type definitions from Go struct definitions.

## Features

- Automatically generates TypeScript interfaces from Go structs
- Handles JSON, form, param, and query tag parsing for field names (prioritized in that order)
- Converts Go types to appropriate TypeScript types
- Preserves comments from Go code in the generated TypeScript
- Handles both exported and unexported types and fields
- Converts field names to camelCase for non-API types
- Special handling for API-related types (preserves exact field names)
- Handles pointer types as optional fields
- Supports `omitempty` tag for optional fields
- Handles arrays of pointers as `(Type | null | undefined)[]` in TypeScript
- Includes validation rules from struct tags as JSDoc comments
- Parses Swagger/OpenAPI annotations to add API endpoint information to type comments
- Supports processing multiple source directories in a single command

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
go-ts-generator <source_dirs> <target_file>
```

Where:
- `<source_dirs>` is a comma-separated list of directories containing Go files to parse
- `<target_file>` is the target TypeScript file to generate

Examples:
```bash
# Single directory
go-ts-generator ./models ./types/generated.ts

# Multiple directories
go-ts-generator ./models,./api,./controllers ./types/generated.ts
```

### As a library

```go
package main

import (
	"fmt"
	"github.com/mo49/go-ts-generator/pkg/generator"
)

func main() {
	// Generate from a single directory
	err := generator.GenerateTypes("./models", "./types/generated.ts")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	
	// Or generate from multiple directories
	sourceDirs := []string{"./models", "./api", "./controllers"}
	err = generator.GenerateTypesFromMultipleDirs(sourceDirs, "./types/generated.ts")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
```

### Examples

Check out the [examples](./examples) directory for complete usage examples:
- [Basic examples](./examples/basic) - Standard Go structs and their TypeScript equivalents
- [API examples](./examples/api) - API-related types with preserved field names
- [Swagger examples](./examples/swagger) - Swagger/OpenAPI annotations for API documentation

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
export interface unexportedType {
  publicField: string;
  /**
   * Note: This is an unexported field. In Go code, it's defined with a lowercase identifier.
   * It cannot be accessed directly from outside the package.
   */
  privateField: number;
}
```

For more detailed examples, please see the [examples](./examples) directory.

## License

MIT