# Examples

This directory contains examples of how to use go-ts-generator.

## Basic Example

The `basic` directory contains standard Go structs that demonstrate:
- Basic type conversion
- Handling of pointer types (both with and without `omitempty`)
- Handling of unexported types and fields
- Preservation of original field names from struct tags

Files:
- `models.go` - Go struct definitions
- `generated.ts` - Generated TypeScript type definitions

## API Example

The `api` directory contains API-related structs that demonstrate:
- Handling of query tags
- Handling of request/response types

Files:
- `api_models.go` - Go struct definitions for API-related types
- `generated.ts` - Generated TypeScript type definitions

## Swagger Example

The `swagger` directory contains Swagger/OpenAPI annotations that demonstrate:
- Extraction of API endpoint information from Swagger comments
- Association of types with their usage in API endpoints

Files:
- `api_swagger.go` - Go file with Swagger/OpenAPI annotations

## Usage

You can generate TypeScript types from one or multiple directories:

```bash
# Install the tool
go install github.com/mo49/go-ts-generator/cmd/go-ts-generator@latest

# Generate from a single directory
go-ts-generator ./examples/basic ./examples/basic/generated.ts
go-ts-generator ./examples/api ./examples/api/generated.ts

# Generate from multiple directories
go-ts-generator ./examples/api,./examples/swagger ./examples/combined_output.ts
```

When using multiple directories, the tool will:
- Collect all type definitions from all specified directories
- Combine related information (like Swagger annotations with their types)
- Generate a single TypeScript file with complete type definitions

## Pre-generated Files

For convenience, this repository includes pre-generated TypeScript files (`generated.ts`) in the example directories. These files show the expected output of the tool when run on the example Go files. 