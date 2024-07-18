package main

import (
	"context"
	"eip712-types-generator/internal/generator"
	_ "embed"
	"flag"
	"os"

	_ "embed"

	"github.com/rs/zerolog"
)

const (
	packageName = "eip712_types"
	fileName    = "eip712_types.go"
)

//go:generate go run . -output=../../types -filepath=../../types/eip712_types.json generate
func main() {
	ctx := context.Background()
	var outDir, inFile string
	flag.StringVar(&outDir, "output", generator.DefaultOutDir, "Output directory for the generated Go file")
	flag.StringVar(&inFile, "filepath", generator.DefaultFilePath, "Path to eip-712 types json is empty")
	flag.Parse()

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "eip712-types-generator").
		Logger()

	generator, err := generator.New(logger, packageName, inFile, outDir, fileName)
	if err != nil {
		logger.Err(err).Msg("failed to create generator")
		os.Exit(1)
	}

	if err := generator.Execute(ctx); err != nil {
		logger.Err(err).Msg("failed to execute generator")
		os.Exit(1)
	}

}
