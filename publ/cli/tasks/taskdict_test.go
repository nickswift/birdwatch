package tasks

import (
	"testing"
)

func TestTaskdictComposition(t *testing.T) {
	if Tasks[`simple`] == nil {
		t.Fatalf("bad key")
	}
}
