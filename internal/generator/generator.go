package generator

import (
	"bytes"
	"fmt"
	"os"
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
	// DefaultPackageName is the default package name for the generated Go file
	DefaultPackageName = "eip712_types"
	// DefaultFileName is the default file name for the generated Go file
	DefaultFileName = "eip712_types.go"
)

var (
	solidityToGolangTypesMap = map[string]string{
		"string[]": "[]string",
		"uint256":  "*big.Int",
		"address":  "common.Address",
		"string":   "string",
	}

	typedDataStructMap = map[string]string{
		"string[]": "anySlice(%s.%s)",
		"uint256":  "hexutil.EncodeBig(%s.%s)",
		"address":  "%s.%s.Hex()",
		"string":   "%s.%s",
	}
)

type Generator struct {
	logger   zerolog.Logger
	pkg      string
	template *template.Template
}

func New(logger zerolog.Logger, pkg string) (*Generator, error) {

	tmpl, err := template.New("").Parse(packageTemplate)
	if err != nil {
		return nil, err
	}

	return &Generator{
		logger:   logger,
		pkg:      pkg,
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

func (g *Generator) BuildTemplate(data []byte) ([]byte, error) {
	var eip712Types map[string][]Arguments
	if err := json.Unmarshal(data, &eip712Types); err != nil {
		return nil, fmt.Errorf("failed to unmarshal eip712 types json: %w", err)
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
		return nil, fmt.Errorf("failed to execute eip712 types template: %w", err)
	}
	return outBuf.Bytes(), nil
}

func (g *Generator) WriteToFile(data []byte, outPath string) error {
	formattedData, err := imports.Process(outPath, data, &imports.Options{
		AllErrors: true,
		Comments:  true,
	})
	if err != nil {
		return err
	}

	goOutputFile, err := os.Create(outPath)
	if err != nil {
		return err
	}

	_, err = goOutputFile.Write(formattedData)
	return err
}
