package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
		fmt.Println("Use 'make build-wasm' and open web/wasm/index.html")
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage: erpnp <command> [options]")
	fmt.Println("Commands: validate, transform, add, remove, sum, wasm")
}

func cmdValidate(args []string) {
	fs := flag.NewFlagSet("validate", flag.ExitOnError)
	f := fs.String("file", "", "path to input file")
	if err := fs.Parse(args); err != nil {
		log.Fatal(err)
	}
	if *f == "" {
		log.Fatal("--file required")
	}
	fmt.Printf("(stub) validate %s\n", *f)
	// TODO: call parser, run schema validation, print errors / OK
}

func cmdTransform(args []string) {
	fs := flag.NewFlagSet("transform", flag.ExitOnError)
	in := fs.String("in", "", "input file")
	out := fs.String("out", "", "output file (optional)")
	if err := fs.Parse(args); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("(stub) transform %s -> %s\n", *in, *out)
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
