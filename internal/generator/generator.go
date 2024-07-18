package generator

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	_ "embed" // embed template file

	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
	"golang.org/x/tools/imports"
)

//go:embed types.tmpl
var packageTemplate string

const (
	// DefaultOutDir is the default output directory for the generated Go file
	DefaultOutDir = "../../types"
	// DefaultFilePath is the default path to eip-712 types json
	DefaultFilePath = "../../types/eip712_types.json"
)

type Generator struct {
	logger   zerolog.Logger
	pkg      string
	inPath   string
	outPath  string
	template *template.Template
}

func New(logger zerolog.Logger, pkg, path, outDir, outFile string) (*Generator, error) {
	if err := os.Mkdir(filepath.Clean(outDir), 0700); os.IsNotExist(err) {
		if err != nil {
			logger.Err(err).Msg("failed to create output directory")
			return nil, err
		}
	}

	tmpl, err := template.New("").Parse(packageTemplate)
	if err != nil {
		return nil, err
	}

	return &Generator{
		logger:   logger,
		pkg:      pkg,
		inPath:   path,
		outPath:  filepath.Join(filepath.Clean(outDir), outFile),
		template: tmpl,
	}, nil
}

type TemplateData struct {
	PackageName string
	Methods     []Method
}

type Arguments struct {
	CapitalName          string
	Name                 string
	Type                 string
	GoType               string
	TypeDataStructMethod string
}

type Method struct {
	Name      string
	Arguments []Arguments
	Alias     string
}

func (g *Generator) Execute(ctx context.Context) error {
	typesJSON, err := os.ReadFile(filepath.Clean(g.inPath))
	if err != nil {
		return fmt.Errorf("failed to read eip712 types json: %w", err)
	}

	var eip712Types map[string][]Arguments
	if err := json.Unmarshal(typesJSON, &eip712Types); err != nil {
		return fmt.Errorf("failed to unmarshal eip712 types json: %w", err)
	}

	templateData := TemplateData{
		PackageName: g.pkg,
	}

	for methodName, argArray := range eip712Types {
		m := Method{
			Name:  methodName,
			Alias: strings.ToLower(methodName[:1]),
		}

		for _, arg := range argArray {
			arg.CapitalName = strings.ToUpper(arg.Name[:1]) + arg.Name[1:]
			arg.GoType = solidityToGolangTypesMap[arg.Type]
			arg.TypeDataStructMethod = fmt.Sprintf(typedDataStructMap[arg.Type], m.Alias, arg.CapitalName)
			m.Arguments = append(m.Arguments, arg)
		}

		templateData.Methods = append(templateData.Methods, m)
	}

	var outBuf bytes.Buffer
	if err := g.template.Execute(&outBuf, templateData); err != nil {
		return fmt.Errorf("failed to execute eip712 types template: %w", err)
	}

	formattedData, err := imports.Process(g.outPath, outBuf.Bytes(), &imports.Options{
		AllErrors: true,
		Comments:  true,
	})

	goOutputFile, err := os.Create(g.outPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}

	if _, err = goOutputFile.Write(formattedData); err != nil {
		return fmt.Errorf("failed to write to output file: %w", err)
	}

	g.logger.Info().Msgf("successfully generated eip712 types at: %s", g.outPath)
	return nil
}

var solidityToGolangTypesMap = map[string]string{
	"string[]": "[]string",
	"uint256":  "*big.Int",
	"address":  "common.Address",
	"string":   "string",
}

var typedDataStructMap = map[string]string{
	"string[]": "anySlice(%s.%s)",
	"uint256":  "hexutil.EncodeBig(%s.%s)",
	"address":  "%s.%s.Hex()",
	"string":   "%s.%s",
}
