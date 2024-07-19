package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/DIMO-Network/eip712-types-generator/internal/generator"
	"github.com/rs/zerolog"
)

func main() {
	var outDir, inFile, packageName, fileName string
	flag.StringVar(&outDir, "output", generator.DefaultOutDir, "Output directory for the generated Go file")
	flag.StringVar(&inFile, "filepath", generator.DefaultFilePath, "Path to eip-712 types json is empty")
	flag.StringVar(&packageName, "packageName", generator.DefaultPackageName, "Package name for the generated Go file")
	flag.StringVar(&fileName, "fileName", generator.DefaultFileName, "Resulting file name for the generated Go file")
	flag.Parse()

	logger := zerolog.New(os.Stdout).With().Timestamp().Str("app", "eip712-types-generator").Logger()

	if err := os.Mkdir(filepath.Clean(outDir), 0700); os.IsNotExist(err) {
		logger.Err(err).Msg("failed to create output directory")
		os.Exit(1)
	}

	typesJSON, err := os.ReadFile(filepath.Clean(inFile))
	if err != nil {
		logger.Err(err).Msg("failed to read types json")
		os.Exit(1)
	}

	generator, err := generator.New(logger, packageName)
	if err != nil {
		logger.Err(err).Msg("failed to create generator")
		os.Exit(1)
	}

	templB, err := generator.BuildTemplate(typesJSON)
	if err != nil {
		logger.Err(err).Msg("failed to build template")
		os.Exit(1)
	}

	out := filepath.Join(filepath.Clean(outDir), fileName)
	if err := generator.WriteToFile(templB, out); err != nil {
		logger.Err(err).Msg("failed to write to output file")
		os.Exit(1)
	}

	logger.Info().Msgf("successfully generated eip712 types at: %s", out)
}
