package main

import (
	"context"
	"flag"
	"os"

	"eip-712-types-generator/internal/generator"

	"github.com/rs/zerolog"
)

//go:generate go run . -package=eip712_types -out=eip712_types.go -outDir=../../output -filepath=../../types/eip712_types.json generate
func main() {
	ctx := context.Background()
	var packageName, outFile, outDir, filepath string
	flag.StringVar(&packageName, "package", generator.DefaultPackageName, "Name of the package to generate")
	flag.StringVar(&outFile, "out", generator.DefaultOutFile, "Output file for the generated Go file")
	flag.StringVar(&outDir, "outDir", generator.DefaultOutDir, "Output directory for the generated Go file")
	flag.StringVar(&filepath, "filepath", generator.DefaultFilePath, "Path to eip-712 types json is empty")
	flag.Parse()

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "eip712-types-generator").
		Logger()

	generator, err := generator.New(logger, packageName, filepath, outDir, outFile)
	if err != nil {
		logger.Err(err).Msg("failed to create generator")
		os.Exit(1)
	}

	if err := generator.Execute(ctx); err != nil {
		logger.Err(err).Msg("failed to generate eip-712 types")
		os.Exit(1)
	}

	os.Exit(0)
}
