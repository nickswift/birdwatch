package main

import (
	"bufio"
	"fmt"
	"github.com/nickswift498/birdwatch/cli"
	config "github.com/nickswift498/birdwatch/config"
	"github.com/nickswift498/birdwatch/publ/cli/actions"
	"github.com/nickswift498/birdwatch/publ/cli/tasks"
	"os"
)

func title() {
	fmt.Println(`             _     _ _      _     _   `)
	fmt.Println(`            | |   | (_)    (_)   | |  `)
	fmt.Println(` _ __  _   _| |__ | |_  ___ _ ___| |_ `)
	fmt.Println(`| '_ \| | | | '_ \| | |/ __| / __| __|`)
	fmt.Println(`| |_) | |_| | |_) | | | (__| \__ \ |_ `)
	fmt.Println(`| .__/ \__,_|_.__/|_|_|\___|_|___/\__|`)
	fmt.Println(`| |                                   `)
	fmt.Println(`|_|                                   `)
	fmt.Println(``)
}
func version() string {
	return fmt.Sprintf("publicist v%d.%d", config.PUBL_VERSION_MAJOR, config.PUBL_VERSION_MINOR)
}

func main() {
	title()

	cliReader := bufio.NewReader(os.Stdin)

	for {
		// user prompt
		fmt.Printf("%s > ", version())
		cmd, _ := cliReader.ReadString('\n')

		paction, ptask, pargs := cli.ProcessCommand(cmd)
		if pargs == nil {
			pargs = []string{""}
		}

		// TODO: catch unrecognized commands

		// get action/task functions
		action := actions.Actions[paction]
		task := tasks.Tasks[ptask]

		action(task, pargs...)
	}
}
