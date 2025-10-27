package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/zinrai/cuebridge"
)

const version = "0.1.0"

func main() {
	var (
		schemaPath  string
		configPath  string
		showVersion bool
	)

	setupFlags(&schemaPath, &configPath, &showVersion)
	flag.Parse()

	if showVersion {
		printVersion()
		os.Exit(0)
	}

	if err := validateArgs(schemaPath, configPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n\n", err)
		flag.Usage()
		os.Exit(1)
	}

	runValidation(schemaPath, configPath)
}

// setupFlags configures command-line flags
func setupFlags(schemaPath *string, configPath *string, showVersion *bool) {
	flag.StringVar(schemaPath, "schema", "", "Path to CUE schema file (required)")
	flag.StringVar(configPath, "config", "", "Path to JSON config file to validate (required)")
	flag.BoolVar(showVersion, "version", false, "Show version")
	flag.Usage = createUsageFunc()
}

// createUsageFunc creates the custom usage function
func createUsageFunc() func() {
	return func() {
		progName := filepath.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "integratify - JSON validation tool powered by CUE\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s -schema=<schema.cue> -config=<config.json>\n\n", progName)
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
}

// printVersion prints the version information
func printVersion() {
	fmt.Printf("integratify version %s\n", version)
}

// validateArgs validates command-line arguments
func validateArgs(schemaPath string, configPath string) error {
	if schemaPath == "" {
		return fmt.Errorf("-schema is required")
	}
	if configPath == "" {
		return fmt.Errorf("-config is required")
	}
	return nil
}

// runValidation runs the validation and handles the results
func runValidation(schemaPath string, configPath string) {
	// Create validator with #Config definition
	validator, err := cuebridge.NewValidator(schemaPath, "#Config")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading schema: %v\n", err)
		os.Exit(1)
	}

	// Validate the JSON file
	result, err := validator.Validate(cuebridge.ValidationInput{
		SourceType: cuebridge.SourceFile,
		FilePath:   configPath,
		Format:     cuebridge.FormatJSON,
		Name:       configPath,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error validating %s: %v\n", configPath, err)
		os.Exit(1)
	}

	// Format and display results
	output := cuebridge.FormatResults([]cuebridge.ValidationResult{result})
	fmt.Print(output)

	// Determine exit code
	if !result.Valid {
		os.Exit(1)
	}
}
