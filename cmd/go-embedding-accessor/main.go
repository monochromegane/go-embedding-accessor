package main

import (
	"flag"
	"fmt"
	"os"

	accessor "github.com/monochromegane/go-embedding-accessor"
)

var pkg string
var outputDir string
var name string
var files string

func init() {
	flag.StringVar(&pkg, "pkg", "main", "Package name to use in the generated code.")
	flag.StringVar(&outputDir, "output-dir", ".", "Output directory")
	flag.StringVar(&name, "name", "asset", "Name of embeddings")
	flag.StringVar(&files, "files", "", "File path of embeddings for go:embed directive")
	flag.Parse()
}

func main() {
	err := accessor.Generate(pkg, outputDir, name, files)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
