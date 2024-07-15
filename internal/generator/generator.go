package generator

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/goccy/go-json"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

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
	logger  zerolog.Logger
	pkg     string
	inPath  string
	outPath string
}

func New(logger zerolog.Logger, pkg string, path string, outDir string, outFile string) (*Generator, error) {
	if err := os.Mkdir(outDir, 0700); os.IsNotExist(err) {
		if err != nil {
			logger.Err(err).Msg("failed to create output directory")
			return nil, err
		}
	}

	return &Generator{
		logger:  logger,
		pkg:     pkg,
		inPath:  path,
		outPath: filepath.Join(outDir, outFile),
	}, nil
}

func (g *Generator) Execute(ctx context.Context) error {
	tf, err := os.Create(g.outPath)
	if err != nil {
		g.logger.Fatal().Err(err).Msg("failed to create output file")
	}

	jsonFile, err := os.Open(g.inPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			g.logger.Err(err).Msg(fmt.Sprintf("%s not found", g.inPath))
			return err
		}
		g.logger.Err(err).Msg(fmt.Sprintf("failed to open eip712 types file: %s", g.inPath))
		return err
	}
	defer jsonFile.Close()

	jsonB, err := io.ReadAll(jsonFile)
	if err != nil {
		g.logger.Err(err).Msg("failed to read eip712 types file")
		return err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(jsonB, &result); err != nil {
		g.logger.Err(err).Msg("failed to unmarshal eip712 types json")
		return err
	}

	fmt.Fprintf(tf, "package registry\n\n")
	for structName, detailsArray := range result {
		alias := strings.ToLower(structName[:1])

		solSig := fmt.Sprintf("//%s(", structName)
		structAttrs := fmt.Sprintf("type %s struct {\n", structName)
		typeFunc := fmt.Sprintf(`func (%s *%s) Type() []apitypes.Type {return []apitypes.Type{`, alias, structName)
		messageFunc := fmt.Sprintf(`func (%s *%s) Message() apitypes.TypedDataMessage {return apitypes.TypedDataMessage{`, alias, structName)
		for _, args := range detailsArray.([]interface{}) {
			argName, argNameOk := args.(map[string]interface{})["name"].(string)
			argType, argTypeOk := args.(map[string]interface{})["type"].(string)
			if !argNameOk || !argTypeOk {
				g.logger.Err(fmt.Errorf("failed to parse args for %s", structName)).Msg("failed to parse args for struct")
				return err
			}

			solSig += fmt.Sprintf("%s %s,", argType, argName)
			upperArgName := strings.ToUpper(argName[:1]) + argName[1:]
			typeFunc += fmt.Sprintf(`{Name: "%s", Type: "%s"},`, argName, argType)
			messageFunc += fmt.Sprintf(`"%s": %s,`, argName, fmt.Sprintf(typedDataStructMap[argType], alias, upperArgName))
			structAttrs += fmt.Sprintf("%s %s `json:\"%s\"`\n", upperArgName, solidityToGolangTypesMap[argType], argName)
		}

		solSig = strings.TrimRight(solSig, ",") + ")"
		typeFunc += "}}\n\n"
		messageFunc += "}}\n\n"

		fmt.Fprint(tf, solSig+"\n"+structAttrs+"\n")
		fmt.Fprint(tf, "}\n\n"+fmt.Sprintf(`func (%s *%s) Name() string {return "%s"}`, alias, structName, structName)+"\n\n")
		typedDataHash := fmt.Sprintf("func (%s *%s) TypedDataAndHash(domain apitypes.TypedDataDomain) ([]byte, error) {\ntd := &apitypes.TypedData {\nTypes: apitypes.Types{\n\t\"EIP712Domain\": []apitypes.Type{\n\t\t{Name: \"name\", Type: \"string\"},\n{Name: \"version\", Type: \"string\"},\n{Name: \"chainId\", Type: \"uint256\"},\n{Name: \"verifyingContract\", Type: \"address\"},\n}, %s.Name(): %s.Type(),\n},\nPrimaryType: %s.Name(),\nDomain: domain,\nMessage: %s.Message()}\nhash, _, err := apitypes.TypedDataAndHash(*td)\nreturn hash, err}\n\n", alias, structName, alias, alias, alias, alias)
		fmt.Fprint(tf, typeFunc)
		fmt.Fprint(tf, messageFunc)
		fmt.Fprint(tf, typedDataHash)

	}

	fmt.Fprint(tf, "func anySlice[A any](v []A) []any {\nn := len(v)\nout := make([]any, n)\nfor i := 0; i < n; i++ {\nout[i] = v[i]\n}\n\nreturn out\n}\n\n")

	if err := exec.Command("gofmt", "-w", g.outPath).Run(); err != nil {
		g.logger.Err(err).Msg("failed to run gofmt")
	}
	if err := exec.Command("goimports", "-w", g.outPath).Run(); err != nil {
		g.logger.Err(err).Msg("failed to run goimports")
	}

	if err := tf.Close(); err != nil {
		g.logger.Err(err).Msg("failed to close output file")
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
