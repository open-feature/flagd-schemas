package flagd_definitions_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	flagd_definitions "github.com/open-feature/flagd-schemas/json"

	"github.com/xeipuuv/gojsonschema"
)

var compiledSchema *gojsonschema.Schema

func init() {
	schemaLoader := gojsonschema.NewSchemaLoader()
	schemaLoader.AddSchemas(gojsonschema.NewStringLoader(flagd_definitions.TargetingSchema))
	var err error
	compiledSchema, err = schemaLoader.Compile(gojsonschema.NewStringLoader(flagd_definitions.FlagSchema))
	if err != nil {
		message := fmt.Errorf("err: %v", err)
		log.Fatal(message)
	}
}

func TestPositiveParsing(t *testing.T) {
	if err := walkPath(true, "./test/positive"); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestNegativeParsing(t *testing.T) {
	if err := walkPath(false, "./test/negative"); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func walkPath(shouldPass bool, root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		ps := strings.Split(path, ".")
		if ps[len(ps)-1] != "json" {
			return nil
		}
		file, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		flagStringLoader := gojsonschema.NewStringLoader(string(file))

		p, err := compiledSchema.Validate(flagStringLoader)
		if err != nil {
			return err
		}

		if p.Valid() && shouldPass == false {
			return fmt.Errorf("file %s should have failed validation, but did not", path)
		}

		if !p.Valid() && shouldPass == true {
			return fmt.Errorf("file %s should not have failed validation, but did", path)
		}

		return nil
	})
}
