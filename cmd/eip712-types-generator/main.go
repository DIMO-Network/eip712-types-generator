package main

import (
	"flag"
	"io"
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

	typesInfo, err := os.Stat(typesPath)
	if err != nil {
		logger.Fatal().Err(err).Msgf("Couldn't stat file %q.", typesPath)
	}

	typesFile, err := os.Open(typesPath)
	if err != nil {
		logger.Fatal().Err(err).Msgf("Couldn't open file %q.", typesPath)
	}

	in, err := io.ReadAll(typesFile)
	if err != nil {
		logger.Fatal().Err(err).Msgf("Couldn't read file %q.", typesPath)
	}

	g, err := generator.New()
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create generator.")
	}

	out, err := g.Execute(*packageName, in)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to execute template.")
	}

	if *outPath == "" {
		_, err := os.Stdin.Write(out)
		if err != nil {
			logger.Fatal().Err(err).Msg("Couldn't write to stdout.")
		}
	} else {
		err := os.WriteFile(*outPath, out, typesInfo.Mode())
		if err != nil {
			logger.Fatal().Err(err).Msgf("Couldn't write to file %q.", *outPath)
		}
	}
}
