# Examples

This directory contains examples of how to use go-ts-generator.

## Basic Example

The `basic` directory contains standard Go structs that demonstrate:
- Basic type conversion
- Handling of pointer types (both with and without `omitempty`)
- Handling of unexported types and fields
- Conversion of field names to camelCase

Files:
- `models.go` - Go struct definitions
- `generated.ts` - Generated TypeScript type definitions

To generate TypeScript types from these models:

```bash
go-ts-generator ./examples/basic ./examples/basic/generated.ts
```

## API Example

The `api` directory contains API-related structs that demonstrate:
- Preservation of field names in API types (no camelCase conversion)
- Handling of query tags
- Handling of request/response types

Files:
- `api_models.go` - Go struct definitions for API-related types
- `generated.ts` - Generated TypeScript type definitions

To generate TypeScript types from these API models:

```bash
go-ts-generator ./examples/api ./examples/api/generated.ts
```

## Running the Examples

You can run both examples with:

```bash
# Install the tool
go install github.com/mo49/go-ts-generator/cmd/go-ts-generator@latest

# Generate TypeScript types for basic models
go-ts-generator ./examples/basic ./examples/basic/generated.ts

# Generate TypeScript types for API models
go-ts-generator ./examples/api ./examples/api/generated.ts
```

The generated TypeScript files will show how different Go types are converted to TypeScript.

## Pre-generated Files

For convenience, this repository includes pre-generated TypeScript files (`generated.ts`) in both example directories. These files show the expected output of the tool when run on the example Go files. 