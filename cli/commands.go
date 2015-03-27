package cli

import (
	"strings"
)

// split command line by whitespace -- return strings representing:
// action: something to do
// function: the source of the information
// args: parameters to be used upon the function
func ProcessCommand(line string) (string, string, []string) {
	line = line[:len(line)-1]
	strs := strings.Split(line, ` `)
	if len(strs) == 1 {
		return strs[0], "", nil
	} else if len(strs) == 2 {
		return strs[0], strs[1], nil
	} else {
		return strs[0], strs[1], strs[2:]
	}
}
