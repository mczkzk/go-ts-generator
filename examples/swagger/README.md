# Swagger Example

This directory contains Go files with Swagger/OpenAPI annotations that demonstrate how go-ts-generator extracts API endpoint information.

## Files

- `api_swagger.go` - Go file with Swagger/OpenAPI annotations for API endpoints

## Features Demonstrated

- Extraction of API endpoint information from Swagger/OpenAPI annotations
- Association of request and response types with API endpoints
- Generation of JSDoc comments with API usage information

## Usage

This directory is intended to be used in combination with the API models directory:

```bash
go-ts-generator ./examples/api,./examples/swagger ./examples/api/generated.ts
```

This will generate TypeScript type definitions that include API endpoint information extracted from Swagger annotations.

## Notes

- The package name in this directory (`package swagger`) can be different from the package name in the API models directory (`package api`).
- go-ts-generator associates types by name, not by package, so types with the same name will be correctly associated across different packages. 