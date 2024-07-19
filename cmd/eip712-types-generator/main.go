package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/DIMO-Network/eip712-types-generator/internal/generator"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("app", "eip712-types-generator").Logger()

	if len(os.Args) == 1 {
		logger.Fatal().Msg("Types file path required.")
	}

	packageName := flag.String("package", "main", "Package name for the generated Go file.")
	outPath := flag.String("outfile", "", "Output file for the generated code. If blank, then stdout will be used.")
	flag.Parse()

	typesPath := os.Args[1]

	rawTypes, mode, err := loadTypeFile(typesPath)
	if err != nil {
		logger.Fatal().Err(err).Msgf("Couldn't load types from file %q.", typesPath)
	}

	g, err := generator.New()
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create generator.")
	}

	out, err := g.Execute(*packageName, rawTypes)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to execute template.")
	}

	if *outPath == "" {
		_, err := os.Stdin.Write(out)
		if err != nil {
			logger.Fatal().Err(err).Msg("Couldn't write to stdout.")
		}
	} else {
		err := os.WriteFile(*outPath, out, mode)
		if err != nil {
			logger.Fatal().Err(err).Msgf("Couldn't write to file %q.", *outPath)
		}
	}
}

func loadTypeFile(path string) ([]byte, fs.FileMode, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, 0, fmt.Errorf("couldn't stat file: %w", err)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, 0, fmt.Errorf("couldn't open file: %w", err)
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, 0, fmt.Errorf("couldn't read file: %w", err)
	}

	return b, info.Mode(), nil
}
