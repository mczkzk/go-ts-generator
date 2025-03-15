// Command go-ts-generator generates TypeScript type definitions from Go structs.
package main

import (
	"fmt"
	"os"

	"github.com/mo49/go-ts-generator/pkg/generator"
)

// Version information
const (
	Version = "0.3.0"
)

func printHelp() {
	fmt.Println("go-ts-generator - Generate TypeScript type definitions from Go structs")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  go-ts-generator [options] <source_dir> <target_file>")
	fmt.Println("")
	fmt.Println("Arguments:")
	fmt.Println("  <source_dir>  - Directory containing Go files to parse")
	fmt.Println("  <target_file> - Target TypeScript file to generate")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  --help     - Show this help message")
	fmt.Println("  --version  - Show version information")
}

func main() {
	// Check for --version flag
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Printf("go-ts-generator version %s\n", Version)
		return
	}

	// Check for --help flag
	if len(os.Args) <= 1 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		printHelp()
		if len(os.Args) <= 1 {
			os.Exit(1)
		}
		return
	}

	// Get source directory and target file from command-line arguments
	if len(os.Args) < 3 {
		fmt.Println("Error: Missing required arguments")
		printHelp()
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
