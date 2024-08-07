package generator

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"go/format"
	"strings"
	"text/template"
)

//go:embed types.tmpl
var packageTemplate string

var (
	solidityToGoType = map[string]string{
		"string[]": "[]string",
		"uint256":  "*big.Int",
		"address":  "common.Address",
		"string":   "string",
	}

	solidityToPresentationConverter = map[string]string{
		"uint256": "(*big.Int).String",
		"address": "common.Address.Hex",
	}
)

type Generator struct {
	template *template.Template
}

func New() (*Generator, error) {
	tmpl, err := template.New("").Parse(packageTemplate)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse template: %w", err)
	}

	return &Generator{
		template: tmpl,
	}, nil
}

type templateData struct {
	Package string
	Types   []typeDesc
}

type member struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	GoName      string `json:"-"`
	GoType      string `json:"-"`
	ConvertFunc string `json:"-"`
}

type typeDesc struct {
	Name    string
	Members []member
}

func (g *Generator) Execute(packageName string, data []byte) ([]byte, error) {
	var eip712Types map[string][]member
	if err := json.Unmarshal(data, &eip712Types); err != nil {
		return nil, fmt.Errorf("failed to unmarshal eip712 types json: %w", err)
	}

	templateData := templateData{
		Package: packageName,
	}

	for typeName, members := range eip712Types {
		t := typeDesc{
			Name: typeName,
		}

		for i, mem := range members {
			if mem.Name == "" {
				return nil, fmt.Errorf("type %s member at index %d has no name", typeName, i)
			}

			mem.GoName = strings.ToUpper(mem.Name[:1]) + mem.Name[1:]

			gt, ok := solidityToGoType[mem.Type]
			if !ok {
				return nil, fmt.Errorf("type %s member %s has unsupported type %s", typeName, mem.Name, mem.Type)
			}

			mem.GoType = gt
			mem.ConvertFunc = solidityToPresentationConverter[mem.Type]

			t.Members = append(t.Members, mem)
		}

		templateData.Types = append(templateData.Types, t)
	}

	var tmplOut bytes.Buffer
	if err := g.template.Execute(&tmplOut, templateData); err != nil {
		return nil, fmt.Errorf("failed to evaluate template: %w", err)
	}

	out, err := format.Source(tmplOut.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to format output file: %w", err)
	}

	return out, nil
}
