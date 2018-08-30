package schema

import (
	"encoding/json"
	"testing"

	"github.com/go-test/deep"
)

func testMarshalUnmarshal(t *testing.T, filename string, output interface{}, sourceKey string) {
	// Convert output object into JSON
	outputJSON, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		t.Error(err)
	}

	// Fetch original source with the given key
	sourceJSON := GetJSONKeyFromSource(t, filename, sourceKey)

	if sourceJSON == "" && string(outputJSON) == "null" {
		return
	}

	var sourceData map[string]interface{}
	var outputData map[string]interface{}

	err = json.Unmarshal([]byte(sourceJSON), &sourceData)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal([]byte(outputJSON), &outputData)
	if err != nil {
		t.Error(err)
	}

	if diff := deep.Equal(sourceData, outputData); diff != nil {
		t.Logf("Output data:\n%s", outputJSON)
		t.Fatal(diff)
	}
}

func TestWholeFile(t *testing.T) {
	filename := "example.json"
	source, _ := LoadSourceFile(t, filename)

	config := map[string]interface{}{
		"languages":   source.Languages,
		"classes":     source.Classes,
		"subclasses":  source.Subclasses,
		"monsters":    source.Monsters,
		"feats":       source.Feats,
		"backgrounds": source.Backgrounds,
		"invocations": source.Invocations,
		"subraces":    source.Subraces,
		"spells":      source.Spells,
		"encounters":  source.Encounters,
		"selections":  source.Selections,
		"races":       source.Races,
	}

	for key, obj := range config {
		testMarshalUnmarshal(t, filename, obj, key)
	}
}
