package greeting

import (
	"regexp"
	"testing"
)

func TestName(t *testing.T) {
	name := "Amol"
	msg, err := Greet(name)
	want := regexp.MustCompile(`\b` + name + `\b`)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
