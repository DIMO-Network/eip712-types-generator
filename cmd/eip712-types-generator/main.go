package main

import (
	"context"
	"eip712-types-generator/internal/generator"
	_ "embed"
	"flag"
	"fmt"
	"os"

	_ "embed"

	"github.com/rs/zerolog"
)

//go:generate go run . -package=eip712_types -out=eip712_types.go -outDir=../../output -filepath=../../types/eip712_types.json generate
func main() {
	ctx := context.Background()
	var packageName, outFile, outDir, inFile string
	flag.StringVar(&packageName, "package", generator.DefaultPackageName, "Name of the package to generate")
	flag.StringVar(&outFile, "out", generator.DefaultOutFile, "Output file for the generated Go file")
	flag.StringVar(&outDir, "outDir", generator.DefaultOutDir, "Output directory for the generated Go file")
	flag.StringVar(&inFile, "filepath", generator.DefaultFilePath, "Path to eip-712 types json is empty")
	flag.Parse()

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", "eip712-types-generator").
		Logger()

	generator, err := generator.New(logger, packageName, inFile, outDir, outFile)
	if err != nil {
		logger.Err(err).Msg("failed to create generator")
		os.Exit(1)
	}

	if err := generator.Execute(ctx); err != nil {
		fmt.Println(err)
		logger.Fatal().Err(err)
		return
	}

}
