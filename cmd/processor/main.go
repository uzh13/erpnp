package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/uzh13/erpnp/internal/core/model/v1_0"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	sub := os.Args[1]

	switch sub {
	case "validate":
		cmdValidate(os.Args[2:])
	case "transform":
		cmdTransform(os.Args[2:])
	case "add":
		cmdAdd(os.Args[2:])
	case "remove":
		cmdRemove(os.Args[2:])
	case "sum":
		cmdSum(os.Args[2:])
	case "wasm":
		fmt.Println("Use 'make serve-wasm' and open http://localhost:8080")
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage: erpnp <command> [options]")
	fmt.Println("Commands:")
	fmt.Println("  validate  - Validate ERPN file structure and schema")
	fmt.Println("  transform - Convert ERPN file between formats (JSON/YAML/TOML/JSON5)")
	fmt.Println("  add       - Add nodes (stub)")
	fmt.Println("  remove    - Remove nodes (stub)")
	fmt.Println("  sum       - Sum fields (stub)")
	fmt.Println("  wasm      - WebAssembly build info")
}

func cmdValidate(args []string) {
	fs := flag.NewFlagSet("validate", flag.ExitOnError)
	f := fs.String("file", "", "path to input file")
	summary := fs.Bool("summary", false, "show file summary")
	if err := fs.Parse(args); err != nil {
		log.Fatal(err)
	}
	if *f == "" {
		log.Fatal("--file required")
	}

	// Load and validate ERPN file
	erpn, err := v1_0.LoadFromFile(*f)
	if err != nil {
		log.Fatalf("Failed to load file %s: %v", *f, err)
	}

	// Validate structure
	errors := v1_0.ValidateERPN(erpn)
	if len(errors) > 0 {
		fmt.Printf("Validation failed for %s:\n", *f)
		for _, err := range errors {
			fmt.Printf("  - %v\n", err)
		}
		os.Exit(1)
	}

	fmt.Printf("✓ %s is valid\n", *f)

	if *summary {
		fmt.Println()
		erpn.PrintSummary(os.Stdout)
	}
}

func cmdTransform(args []string) {
	fs := flag.NewFlagSet("transform", flag.ExitOnError)
	in := fs.String("in", "", "input file")
	out := fs.String("out", "", "output file")
	format := fs.String("format", "", "output format (json, yaml, toml, json5) - auto-detected from output file if not specified")
	if err := fs.Parse(args); err != nil {
		log.Fatal(err)
	}

	if *in == "" {
		log.Fatal("--in required")
	}
	if *out == "" {
		log.Fatal("--out required")
	}

	// Detect output format if not specified
	outputFormat := *format
	if outputFormat == "" {
		outputFormat = v1_0.DetectFormat(*out)
		if outputFormat == "" {
			log.Fatal("Cannot detect output format from file extension. Use --format flag.")
		}
	}

	// Convert file
	err := v1_0.ConvertFile(*in, *out)
	if err != nil {
		log.Fatalf("Failed to convert %s to %s: %v", *in, *out, err)
	}

	fmt.Printf("✓ Converted %s to %s (%s format)\n", *in, *out, outputFormat)
}

func cmdAdd(args []string) {
	fmt.Println("(stub) add node")
}

func cmdRemove(args []string) {
	fmt.Println("(stub) remove node")
}

func cmdSum(args []string) {
	fmt.Println("(stub) sum fields")
}
