// Command go-ts-generator generates TypeScript type definitions from Go structs.
package main

import (
	"fmt"
	"os"

	"github.com/mo49/go-ts-generator/pkg/generator"
)

func main() {
	// Get source directory and target file from command-line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: go-ts-generator <source_dir> <target_file>")
		fmt.Println("  <source_dir>  - Directory containing Go files to parse")
		fmt.Println("  <target_file> - Target TypeScript file to generate")
		os.Exit(1)
	}

	sourceDir := os.Args[1]
	targetFile := os.Args[2]

	// Generate TypeScript types
	err := generator.GenerateTypes(sourceDir, targetFile)
	if err != nil {
		fmt.Printf("Error generating TypeScript types: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("TypeScript type definitions generated: %s\n", targetFile)
}
