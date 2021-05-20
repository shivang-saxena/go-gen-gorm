package main

import (
	"context"
	"flag"
	"log"
	"path/filepath"
	"time"

	"github.com/shivang-saxena/go-gen-gorm/internal/parser"
	"github.com/shivang-saxena/go-gen-gorm/internal/queryset/generator"
)

func main() {
	const defaultOutPath = "generated_{in}"

	inFile := flag.String("in", "models.go", "path to input file")
	outFile := flag.String("out", defaultOutPath, "path to output file")
	timeout := flag.Duration("timeout", time.Minute, "timeout for generation")
	flag.Parse()

	if *outFile == defaultOutPath {
		*outFile = filepath.Join(filepath.Dir(*inFile), "generated_"+filepath.Base(*inFile))
	}

	g := generator.Generator{
		StructsParser: &parser.Structs{},
	}

	ctx, finish := context.WithTimeout(context.Background(), *timeout)
	defer finish()

	if err := g.Generate(ctx, *inFile, *outFile); err != nil {
		log.Fatalf("can't generate query sets: %s", err)
	}
}
