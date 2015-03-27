package cli

import (
	"fmt"
	"testing"
)

func TestProcessCommandOneItem(t *testing.T) {
	line := `action`
	a, b, c := ProcessCommand(line)

	fmt.Printf("%v,%v,%v\n", a, b, c)
}

func TestProcessCommandTwoItems(t *testing.T) {
	line := `action command`
	a, b, c := ProcessCommand(line)

	fmt.Printf("%v,%v,%v\n", a, b, c)
}

func TestProcessCommandThreeItems(t *testing.T) {
	line := `action command arg`
	a, b, c := ProcessCommand(line)

	fmt.Printf("%v,%v,%v\n", a, b, c)
}
