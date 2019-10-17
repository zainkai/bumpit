package internal_test

import (
	"testing"

	"github.com/zainkai/bumpit/internal/commands"
)

var testsIsVersionNo = []struct {
	TestName string
	Input    string
	Expected bool
}{
	{
		"simple",
		`"version": "1.1.1"`,
		true,
	}, {
		"simple with whitespace",
		`     "version": "1.1.1"`,
		true,
	}, {
		"simple with tabs",
		`\t\t"version": "3.2.1-4343ese"`,
		true,
	}, {
		"simple with tabs and spaces",
		`\t\   t"version":   "1.12.3"`,
		true,
	}, {
		"simple with lots of spaces",
		`                                      "version":                "3.34.3"`,
		true,
	}, {
		"bad input",
		`"version": `,
		false,
	}, {
		"no quotes around numbers",
		`"version": 1.1.1`,
		false,
	}, {
		"not versioning line",
		`              "express": "1.1.1"`,
		false,
	},
}

func TestIsVersionNo(t *testing.T) {
	for _, test := range testsIsVersionNo {
		actual := commands.IsVersionNo(test.Input)
		if actual != test.Expected {
			t.Errorf("Failed: %s", test.TestName)
		}
	}
}

func BenchmarkSimpleShort(b *testing.B) {
	input := `"version": "1.1.1"`
	for i := 0; i < b.N; i++ {
		commands.IsVersionNo(input)
	}
}
func BenchmarkSimpleLong(b *testing.B) {
	input := `            "version":            "1.1.1"           `
	for i := 0; i < b.N; i++ {
		commands.IsVersionNo(input)
	}
}

func BenchmarkBadInputShort(b *testing.B) {
	input := `"express": "1.1.1"`
	for i := 0; i < b.N; i++ {
		commands.IsVersionNo(input)
	}
}

func BenchmarkBadInputLong(b *testing.B) {
	input := `         \t        \t"express":            "1.1.1"                `
	for i := 0; i < b.N; i++ {
		commands.IsVersionNo(input)
	}
}
