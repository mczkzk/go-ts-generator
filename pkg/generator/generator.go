// Package generator provides functionality to generate TypeScript type definitions from Go structs.
package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"
)

// TypeScriptType represents a TypeScript type definition
type TypeScriptType struct {
	Name        string
	Fields      []TypeScriptField
	IsInterface bool
	Comment     string
	IsExported  bool // Whether the type is exported
	IsAPIType   bool // Whether the type is API-related
}

// TypeScriptField represents a field in a TypeScript type
type TypeScriptField struct {
	Name       string
	Type       string
	Optional   bool
	Comment    string
	IsExported bool // Whether the field is exported
	Validation []string
}

// GenerateTypes parses Go files in the source directory and generates TypeScript type definitions
// in the target file.
func GenerateTypes(sourceDir, targetFile string) error {
	// Parse Go files in the source directory
	types, err := ParseGoFiles(sourceDir)
	if err != nil {
		return err
	}

	// Generate TypeScript type definitions
	if err := GenerateTypeScriptTypes(types, targetFile); err != nil {
		return err
	}

	return nil
}

// ParseGoFiles parses Go files in the source directory and collects TypeScript type information
func ParseGoFiles(sourceDir string) ([]TypeScriptType, error) {
	var types []TypeScriptType

	// Search for Go files in the source directory
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Process only Go files
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			// Parse the file
			fset := token.NewFileSet()
			node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if err != nil {
				fmt.Printf("Error parsing file %s: %v\n", path, err)
				return nil
			}

			// Determine if the file is API-related based on the file path
			isAPIFile := strings.Contains(path, "controller") || strings.Contains(path, "handler") || strings.Contains(path, "api")

			// Collect type definitions
			for _, decl := range node.Decls {
				if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
					for _, spec := range genDecl.Specs {
						if typeSpec, ok := spec.(*ast.TypeSpec); ok {
							// Check if the type is exported
							isExported := unicode.IsUpper(rune(typeSpec.Name.Name[0]))

							// For struct types
							if structType, ok := typeSpec.Type.(*ast.StructType); ok {
								tsType := TypeScriptType{
									Name:        typeSpec.Name.Name,
									IsInterface: true,
									IsExported:  isExported,
									IsAPIType: isAPIFile ||
										strings.Contains(typeSpec.Name.Name, "Request") ||
										strings.Contains(typeSpec.Name.Name, "Response") ||
										strings.Contains(typeSpec.Name.Name, "Params") ||
										strings.Contains(typeSpec.Name.Name, "Param") ||
										strings.Contains(typeSpec.Name.Name, "Form") ||
										strings.Contains(path, "form") ||
										strings.Contains(path, "api"),
								}

								// Get comments
								if genDecl.Doc != nil {
									tsType.Comment = genDecl.Doc.Text()
								}

								// Collect fields
								if structType.Fields != nil {
									for _, field := range structType.Fields.List {
										if len(field.Names) > 0 {
											fieldName := field.Names[0].Name
											fieldType, isPointer := getTypeString(field.Type)

											// Check if the field is exported
											isFieldExported := unicode.IsUpper(rune(fieldName[0]))

											// Field comment
											fieldComment := ""
											if field.Comment != nil {
												fieldComment = field.Comment.Text()
											}

											// Parse tags
											jsonName := fieldName
											optional := isPointer        // Pointer types are optional
											var validationRules []string // Store validation rules for JSDoc

											if field.Tag != nil {
												tag := strings.Trim(field.Tag.Value, "`")

												// Extract validation rules
												bindingTag := extractTag(tag, "binding")
												validateTag := extractTag(tag, "validate")

												// Add binding validation rules if present
												if bindingTag != "" {
													validationRules = append(validationRules, "binding: "+bindingTag)
												}

												// Add validate validation rules if present
												if validateTag != "" {
													validationRules = append(validationRules, "validate: "+validateTag)
												}

												// Parse tags in order of priority: json, form, param, query
												jsonTag := extractTag(tag, "json")
												formTag := extractTag(tag, "form")
												paramTag := extractTag(tag, "param")
												queryTag := extractTag(tag, "query")

												// First priority: json tag
												if jsonTag != "" {
													parts := strings.Split(jsonTag, ",")
													if parts[0] != "" && parts[0] != "-" {
														jsonName = parts[0]
													}
													for _, part := range parts[1:] {
														if part == "omitempty" {
															optional = true
														}
													}
												} else if formTag != "" {
													// Second priority: form tag
													parts := strings.Split(formTag, ",")
													if parts[0] != "" && parts[0] != "-" {
														jsonName = parts[0]
													}
													for _, part := range parts[1:] {
														if part == "omitempty" {
															optional = true
														}
													}
												} else if paramTag != "" {
													// Third priority: param tag
													parts := strings.Split(paramTag, ",")
													if parts[0] != "" && parts[0] != "-" {
														jsonName = parts[0]
													}
													for _, part := range parts[1:] {
														if part == "omitempty" {
															optional = true
														}
													}
												} else if queryTag != "" {
													// Fourth priority: query tag
													parts := strings.Split(queryTag, ",")
													if parts[0] != "" && parts[0] != "-" {
														jsonName = parts[0]
													}
													for _, part := range parts[1:] {
														if part == "omitempty" {
															optional = true
														}
													}
												}
											}

											// Convert to camelCase if not API-related
											finalFieldName := jsonName
											if !tsType.IsAPIType && isFieldExported {
												finalFieldName = toCamelCase(jsonName)
											}

											tsType.Fields = append(tsType.Fields, TypeScriptField{
												Name:       finalFieldName,
												Type:       fieldType,
												Optional:   optional,
												Comment:    fieldComment,
												IsExported: isFieldExported,
												Validation: validationRules, // Add validation rules
											})
										}
									}
								}

								types = append(types, tsType)
							} else {
								// For non-struct types (type aliases, etc.)
								tsTypeName := typeSpec.Name.Name
								tsTypeValue, _ := getTypeString(typeSpec.Type)

								tsType := TypeScriptType{
									Name:        tsTypeName,
									IsInterface: false,
									IsExported:  isExported,
									IsAPIType:   isAPIFile || strings.Contains(tsTypeName, "Request") || strings.Contains(tsTypeName, "Response") || strings.Contains(tsTypeName, "Params"),
								}

								// Get comments
								if genDecl.Doc != nil {
									tsType.Comment = genDecl.Doc.Text()
								}

								// Add a single field to represent the type alias
								tsType.Fields = append(tsType.Fields, TypeScriptField{
									Name:       "value",
									Type:       tsTypeValue,
									Optional:   false,
									Comment:    "",
									IsExported: true,
								})

								types = append(types, tsType)
							}
						}
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %v", err)
	}

	return types, nil
}

// getTypeString gets a TypeScript type string from a Go type expression
// Returns (type string, whether it's a pointer type)
func getTypeString(expr ast.Expr) (string, bool) {
	switch t := expr.(type) {
	case *ast.Ident:
		switch t.Name {
		case "string":
			return "string", false
		case "bool":
			return "boolean", false
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float32", "float64":
			return "number", false
		case "time.Time":
			return "string /* RFC3339 */", false
		default:
			// Convert non-exported type names to PascalCase
			typeName := t.Name
			if len(typeName) > 0 && unicode.IsLower(rune(typeName[0])) {
				typeName = strings.ToUpper(typeName[:1]) + typeName[1:]
			}
			return typeName, false
		}
	case *ast.ArrayType:
		// Check if the element type is a pointer
		if starExpr, isPointer := t.Elt.(*ast.StarExpr); isPointer {
			// For array of pointers, get the base type
			baseType, _ := getTypeString(starExpr.X)
			// Return as a union type array (Type | null | undefined)[]
			return "(" + baseType + " | null | undefined)[]", true
		}
		// Regular array type
		elemType, _ := getTypeString(t.Elt)
		return elemType + "[]", false
	case *ast.MapType:
		keyType, _ := getTypeString(t.Key)
		valueType, _ := getTypeString(t.Value)
		return "Record<" + keyType + ", " + valueType + ">", false
	case *ast.SelectorExpr:
		if ident, ok := t.X.(*ast.Ident); ok && ident.Name == "time" && t.Sel.Name == "Time" {
			return "string /* RFC3339 */", false
		}
		return t.Sel.Name, false
	case *ast.StarExpr:
		// For pointer types, get the base type and return a flag indicating it's a pointer
		baseType, _ := getTypeString(t.X)
		return baseType, true
	case *ast.InterfaceType:
		return "any", false
	default:
		return "any", false
	}
}

// extractTag extracts a specific tag value from a tag string
func extractTag(tag, key string) string {
	for _, t := range strings.Split(tag, " ") {
		if strings.HasPrefix(t, key+":") {
			value := strings.TrimPrefix(t, key+":")
			return strings.Trim(value, "\"")
		}
	}
	return ""
}

// GenerateTypeScriptTypes generates TypeScript type definitions
func GenerateTypeScriptTypes(types []TypeScriptType, targetFile string) error {
	file, err := os.Create(targetFile)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	// Collect undefined types
	undefinedTypes := make(map[string]bool)
	for _, t := range types {
		for _, field := range t.Fields {
			// Collect types that are not basic types and not defined
			if !isBasicType(field.Type) && !typeExists(field.Type, types) {
				undefinedTypes[field.Type] = true
			}
		}
	}

	// Write header
	header := `// This file is auto-generated. Do not edit directly.
// Generated at: %s
// Note: This file includes both exported and unexported types and fields.

/* eslint-disable */

`
	fmt.Fprintf(file, header, time.Now().Format("2006-01-02 15:04:05"))

	// Write placeholders for undefined types
	if len(undefinedTypes) > 0 {
		fmt.Fprintln(file, "// Placeholders for undefined types")
		// Avoid duplicates by extracting base types from array types
		processedTypes := make(map[string]bool)

		// Process non-array types first
		for typeName := range undefinedTypes {
			if !strings.HasSuffix(typeName, "[]") && !strings.HasPrefix(typeName, "(") {
				fmt.Fprintf(file, "type %s = any;\n", typeName)
				processedTypes[typeName] = true
			}
		}

		// Then process array types
		for typeName := range undefinedTypes {
			if strings.HasPrefix(typeName, "(") && strings.HasSuffix(typeName, ")[]") {
				// Extract base type from "(Type | null | undefined)[]" format
				baseType := strings.TrimSuffix(strings.TrimPrefix(typeName, "("), " | null | undefined)[]")
				if !processedTypes[baseType] {
					fmt.Fprintf(file, "type %s = any;\n", baseType)
					processedTypes[baseType] = true
				}
			} else if strings.HasSuffix(typeName, "[]") {
				// Extract base type from array type
				baseType := strings.TrimSuffix(typeName, "[]")
				if !processedTypes[baseType] {
					fmt.Fprintf(file, "type %s = any;\n", baseType)
					processedTypes[baseType] = true
				}
			}
		}
		fmt.Fprintln(file, "")
	}

	// Write type definitions
	for _, t := range types {
		// Write comments
		if t.Comment != "" {
			lines := strings.Split(strings.TrimSpace(t.Comment), "\n")
			fmt.Fprintln(file, "/**")
			for _, line := range lines {
				fmt.Fprintf(file, " * %s\n", strings.TrimSpace(line))
			}
			fmt.Fprintln(file, " */")
		}

		// Add a note for unexported types
		if !t.IsExported {
			fmt.Fprintln(file, "/**")
			fmt.Fprintln(file, " * Note: This is an unexported type. In Go code, it's defined with a lowercase identifier.")
			fmt.Fprintln(file, " * It cannot be accessed directly from outside the package.")
			fmt.Fprintln(file, " */")
		}

		// Type names are always in PascalCase (to match TypeScript naming conventions)
		typeName := t.Name
		if !t.IsExported {
			// For unexported types, convert to PascalCase
			if len(typeName) > 0 && unicode.IsLower(rune(typeName[0])) {
				typeName = strings.ToUpper(typeName[:1]) + typeName[1:]
			}
		}

		// Write interface definition or type alias
		if t.IsInterface {
			fmt.Fprintf(file, "export interface %s {\n", typeName)
			for _, field := range t.Fields {
				// Write field comments
				if field.Comment != "" || len(field.Validation) > 0 {
					lines := []string{}

					// Add field comment if present
					if field.Comment != "" {
						commentLines := strings.Split(strings.TrimSpace(field.Comment), "\n")
						for _, line := range commentLines {
							lines = append(lines, strings.TrimSpace(line))
						}
					}

					// Add validation rules if present
					if len(field.Validation) > 0 {
						if len(lines) > 0 {
							lines = append(lines, "") // Add empty line between comment and validation
						}
						lines = append(lines, "@validation")
						for _, rule := range field.Validation {
							lines = append(lines, "  - "+rule)
						}
					}

					fmt.Fprintln(file, "  /**")
					for _, line := range lines {
						fmt.Fprintf(file, "   * %s\n", line)
					}
					fmt.Fprintln(file, "   */")
				}

				// Add a note for unexported fields
				if !field.IsExported {
					fmt.Fprintln(file, "  /**")
					fmt.Fprintln(file, "   * Note: This is an unexported field. In Go code, it's defined with a lowercase identifier.")
					fmt.Fprintln(file, "   * It cannot be accessed directly from outside the package.")
					fmt.Fprintln(file, "   */")
				}

				// Write field definition
				optionalMark := ""
				if field.Optional {
					optionalMark = "?"
				}
				fmt.Fprintf(file, "  %s%s: %s;\n", field.Name, optionalMark, field.Type)
			}
			fmt.Fprintln(file, "}")
		} else {
			// For non-interface types (type aliases)
			if len(t.Fields) > 0 {
				// Use the type of the "value" field as the type alias
				fmt.Fprintf(file, "export type %s = %s;\n", typeName, t.Fields[0].Type)
			}
		}
		fmt.Fprintln(file)
	}

	return nil
}

// isBasicType determines if a type name is a basic type
func isBasicType(typeName string) bool {
	basicTypes := map[string]bool{
		"string":  true,
		"boolean": true,
		"number":  true,
		"any":     true,
	}

	// For array types, check the element type
	if strings.HasSuffix(typeName, "[]") {
		return isBasicType(strings.TrimSuffix(typeName, "[]"))
	}

	// For Record types
	if strings.HasPrefix(typeName, "Record<") {
		return true
	}

	// For string with RFC3339 annotation
	if strings.Contains(typeName, "/* RFC3339 */") {
		return true
	}

	return basicTypes[typeName]
}

// typeExists determines if a specified type is defined
func typeExists(typeName string, types []TypeScriptType) bool {
	// For array types with nullable elements, extract the base type
	if strings.HasPrefix(typeName, "(") && strings.HasSuffix(typeName, ")[]") {
		// Extract the base type from "(Type | null | undefined)[]"
		baseType := strings.TrimSuffix(strings.TrimPrefix(typeName, "("), " | null | undefined)[]")
		return typeExists(baseType, types)
	}

	// For array types, check the element type
	if strings.HasSuffix(typeName, "[]") {
		return typeExists(strings.TrimSuffix(typeName, "[]"), types)
	}

	// Normalize type name (convert to PascalCase if it starts with lowercase)
	normalizedTypeName := typeName
	if len(normalizedTypeName) > 0 && unicode.IsLower(rune(normalizedTypeName[0])) {
		normalizedTypeName = strings.ToUpper(normalizedTypeName[:1]) + normalizedTypeName[1:]
	}

	for _, t := range types {
		// Normalize the type name for comparison
		normalizedName := t.Name
		if len(normalizedName) > 0 && unicode.IsLower(rune(normalizedName[0])) {
			normalizedName = strings.ToUpper(normalizedName[:1]) + normalizedName[1:]
		}

		if normalizedName == normalizedTypeName {
			return true
		}
	}
	return false
}

// toCamelCase converts a snake_case or PascalCase string to camelCase
func toCamelCase(s string) string {
	// If already camelCase, return as is
	if !strings.Contains(s, "_") && !unicode.IsUpper(rune(s[0])) {
		return s
	}

	// For snake_case
	if strings.Contains(s, "_") {
		words := strings.Split(s, "_")
		result := strings.ToLower(words[0])
		for i := 1; i < len(words); i++ {
			if words[i] != "" {
				result += strings.ToUpper(words[i][:1]) + strings.ToLower(words[i][1:])
			}
		}
		return result
	}

	// For PascalCase
	return strings.ToLower(s[:1]) + s[1:]
}
