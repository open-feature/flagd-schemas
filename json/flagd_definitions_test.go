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

var compiledFlagDefinitionSchema *gojsonschema.Schema
var compiledTargetingSchema *gojsonschema.Schema

func init() {
	flagDefinitionSchemaLoader := gojsonschema.NewSchemaLoader()
	flagDefinitionSchemaLoader.AddSchemas(gojsonschema.NewStringLoader(flagd_definitions.TargetingSchema))
	targetingSchemaLoader := gojsonschema.NewSchemaLoader()
	var err error
	compiledFlagDefinitionSchema, err = flagDefinitionSchemaLoader.Compile(gojsonschema.NewStringLoader(flagd_definitions.FlagSchema))
	compiledTargetingSchema, err = targetingSchemaLoader.Compile(gojsonschema.NewStringLoader(flagd_definitions.TargetingSchema))
	if err != nil {
		message := fmt.Errorf("err: %v", err)
		log.Fatal(message)
	}
}

func TestPositiveFlagParsing(t *testing.T) {
	if err := walkPath(true, "./test/flags/positive", compiledFlagDefinitionSchema); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestNegativeFlagParsing(t *testing.T) {
	if err := walkPath(false, "./test/flags/negative", compiledFlagDefinitionSchema); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestPositiveTargetingParsing(t *testing.T) {
	if err := walkPath(true, "./test/targeting/positive", compiledTargetingSchema); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestNegativeTargetingParsing(t *testing.T) {
	if err := walkPath(false, "./test/targeting/negative", compiledTargetingSchema); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func walkPath(shouldPass bool, root string, schema *gojsonschema.Schema) error {
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

		p, err := schema.Validate(flagStringLoader)
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
