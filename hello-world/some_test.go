package helloworld

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "rick"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("rick")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("rick") = %q, %v, want match for %#q, nil`, msg, err, want)
	}

}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
