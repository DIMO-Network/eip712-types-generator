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
var structTemplate string

const (
	// DefaultPackageName is the default package name for the generated Go file
	DefaultPackageName = "eip712_types"
	// DefaultOutFile is the default output file for the generated Go file
	DefaultOutFile = "eip712_types.go"
	// DefaultOutDir is the default output directory for the generated Go file
	DefaultOutDir = "../../output"
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

	tmpl, err := template.New("").Parse(structTemplate)
	if err != nil {
		return nil, err
	}

	return &Generator{
		logger:   logger,
		pkg:      pkg,
		inPath:   path,
		outPath:  filepath.Join(filepath.Clean(outDir), filepath.Clean(outFile)),
		template: tmpl,
	}, nil
}

type TemplData struct {
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
		return err
	}

	var eip712Types map[string][]Arguments
	if err := json.Unmarshal(typesJSON, &eip712Types); err != nil {
		return err
	}

	tdata := TemplData{
		PackageName: g.pkg,
	}

	for k, v := range eip712Types {
		m := Method{
			Name:  k,
			Alias: strings.ToLower(k[:1]),
		}

		for _, arg := range v {
			arg.CapitalName = strings.ToUpper(arg.Name[:1]) + arg.Name[1:]
			arg.GoType = solidityToGolangTypesMap[arg.Type]
			arg.TypeDataStructMethod = fmt.Sprintf(typedDataStructMap[arg.Type], m.Alias, arg.CapitalName)
			m.Arguments = append(m.Arguments, arg)
		}

		tdata.Methods = append(tdata.Methods, m)

	}

	var outBuf bytes.Buffer
	if err := g.template.Execute(&outBuf, tdata); err != nil {
		return err
	}

	formattedData, err := imports.Process(g.outPath, outBuf.Bytes(), &imports.Options{
		AllErrors: true,
		Comments:  true,
	})

	goOutputFile, err := os.Create(g.outPath)
	if err != nil {
		panic(err)
	}

	_, err = goOutputFile.Write(formattedData)
	if err != nil {
		panic(err)
	}

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
