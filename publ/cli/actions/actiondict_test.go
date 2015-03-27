package actions

import (
	"testing"
)

func TestActiondictComposition(t *testing.T) {
	if Actions[`print`] == nil {
		t.Fatalf("bad key")
	}
}
