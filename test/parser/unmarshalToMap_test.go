package parser_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/zainkai/bumpit/pkg/parser"
)

// IsEqualJSONMaps Checks equality of two maps
func IsEqualJSONMaps(actual, expected parser.RawJsonMap) bool {
	for key, bArr := range expected {
		aArr, ok := actual[key]
		if !ok || string(*aArr) != string(*bArr) {
			// fmt.Println(string(*aArr), string(*bArr))
			return false
		}
	}
	return true
}

func strToRawMes(s string) *json.RawMessage {
	result := json.RawMessage(s) // doing this prevents homeless variables
	return &result
}

var tests = []struct {
	input    []byte
	expected parser.RawJsonMap
}{
	{
		[]byte(`{"foo":"bar"}`),
		parser.RawJsonMap{"foo": strToRawMes("\"bar\"")},
	},
}

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")
	os.Exit(exitVal)
}

func TestUnmarshalToMap(t *testing.T) {
	t.Log("GIVEN UnmarshalToMap")
	for _, test := range tests {
		t.Logf("WHEN parsing %s", test.input)
		actual, err := parser.UnmarshalToMap(test.input)

		t.Log("THEN the not throw error")
		if err != nil {
			t.Errorf("FAILED: received error from input \"%v\"", err)
		}

		t.Log("THEN the maps should be the same")
		if !IsEqualJSONMaps(actual, test.expected) {
			t.Errorf("FAILED: actual map doesn't equal expected")
		}
	}
}

// small benchmark tests
func BenchmarkUnmarshalToMapSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parser.UnmarshalToMap(tests[0].input)
	}
}
