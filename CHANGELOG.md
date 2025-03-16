# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.7.0] - 2025-03-17

### Changed
- Simplified TypeScript type generation for arrays of pointers (`[]*Type`) to use `Type[]` instead of `(Type | null | undefined)[]`
- Updated tests to reflect the new simplified array type format
- Updated documentation and examples to match the new type conversion

## [0.6.0] - 2025-03-16

### Added
- Improved Swagger/OpenAPI annotation parsing for better endpoint detection
- Support for different path parameter formats (`:param` and `{param}`)
- Support for alternative @Success annotation formats
- Enhanced handling of types with package prefixes (e.g., `responses.TypeName`)
- Better detection of function-level Swagger annotations

### Fixed
- Fixed issue where endpoint information wasn't extracted from certain Swagger annotation formats
- Fixed handling of path parameters in @Router annotations
- Improved type name extraction from package-qualified names

## [0.5.0] - 2025-03-15

### Added
- Support for processing multiple source directories in a single command
- Command-line now accepts comma-separated list of directories
- New `GenerateTypesFromMultipleDirs` function in the library API
- Updated documentation with examples for multiple directory usage

## [0.4.0] - 2025-03-15

### Added
- Support for parsing Swagger/OpenAPI annotations to extract API endpoint information
- New feature to add endpoint usage information to TypeScript type comments
- Automatic detection of types used as request bodies and responses
- Enhanced JSDoc comments showing which endpoints use each type

## [0.3.0] - 2025-03-15

### Added
- Support for form, param tags in addition to json and query tags
- Prioritization of tags in the order: json > form > param > query
- Improved JSDoc comments with validation rules from struct tags

### Changed
- Improved documentation with clearer examples
- Streamlined README.md by removing redundancies
- Enhanced handling of unexported types and fields

## [0.2.0] - 2025-03-15

### Added
- Support for type aliases in Go to TypeScript conversion
- Proper handling of array of pointers (`[]*Type`) as `(Type | undefined)[]` in TypeScript
- Added examples for type aliases in the basic examples
- Added comprehensive tests for both basic and API type conversions

### Fixed
- Fixed issue where `type Foo []*Bar` was incorrectly converted to `type Foo = any`
- Improved type detection for arrays with pointer elements

## [0.1.0] - 2025-03-15

### Added
- Initial release
- Support for converting Go structs to TypeScript interfaces
- Support for basic Go types (string, bool, int, etc.)
- Support for time.Time conversion
- Support for pointer types as optional fields
- Support for JSON and query tag parsing
- Support for API-related types with preserved field names
- Command-line tool for generating TypeScript types
- Basic examples and API examples 