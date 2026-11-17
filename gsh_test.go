package main

import (
	"strings"
	"testing"
)

func TestCd(t *testing.T) {
	got := execInput("cd")
	want := "path required"
	if strings.Compare(got.Error(), want) != 0 {
		t.Errorf("execInput(cd) = %s; want %s", got, want)
	}
}
