package internal_test

import (
	"testing"

	"github.com/zainkai/bumpit/internal/commands"
)

type benchmarkProcessVersionLine struct {
	s string
	t commands.UPDATE_SEMVER
}

var testsProcessVersionLine = []struct {
	TestName   string
	InputStr   string
	UpdateType commands.UPDATE_SEMVER
	Expected   string
}{
	{
		"Short-Patch",
		`"version": "1.1.1"`,
		commands.PATCH,
		`"version": "1.1.2"`,
	}, {
		"Short-Minor",
		`"version": "1.1.1"`,
		commands.MINOR,
		`"version": "1.2.0"`,
	}, {
		"Short-Major",
		`"version": "1.1.1"`,
		commands.MAJOR,
		`"version": "2.0.0"`,
	}, {
		"Long-Patch",
		`        "version":    "1.1.1"`,
		commands.PATCH,
		`        "version":    "1.1.2"`,
	}, {
		"Long-Minor",
		`       "version": "1.1.1"`,
		commands.MINOR,
		`       "version": "1.2.0"`,
	}, {
		"Long-Major",
		`                         "version":"1.1.1"`,
		commands.MAJOR,
		`                         "version":"2.0.0"`,
	},
}

func TestProcessVersionLine(t *testing.T) {
	for _, test := range testsProcessVersionLine {
		actual := commands.ProcessVersionLine(test.InputStr, test.UpdateType)
		if actual != test.Expected {
			t.Errorf("Failed: %s actual: %s", test.TestName, actual)
		}
	}
}

func BenchmarkShort(b *testing.B) {
	test := benchmarkProcessVersionLine{
		s: `"version": "1.1.1"`,
		t: commands.PATCH,
	}
	for i := 0; i < b.N; i++ {
		commands.ProcessVersionLine(test.s, test.t)
	}
}

func BenchmarkLong(b *testing.B) {
	test := benchmarkProcessVersionLine{
		s: `                               "version":          "1.1.1"    `,
		t: commands.PATCH,
	}
	for i := 0; i < b.N; i++ {
		commands.ProcessVersionLine(test.s, test.t)
	}
}
