package syslog

import (
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestDumb(t *testing.T) {
	line := "<195>"
	got, err := ParseReader("", strings.NewReader(line))
	if err != nil {
		t.Error(err)
	}

	spew.Dump(got)
}
