package main

import (
	"bufio"
	"fmt"
	"github.com/nickswift498/birdwatch/cli"
	"github.com/nickswift498/birdwatch/config"
	"github.com/nickswift498/birdwatch/publ/cli/actions"
	"github.com/nickswift498/birdwatch/publ/cli/tasks"
	"github.com/nickswift498/birdwatch/client"
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
func help() string {
	return "help string here"
}

func main() {
	title()

	// get Twitter API
	tc := client.TwitterClientFromEnv()

	cliReader := bufio.NewReader(os.Stdin)
	for {
		// user prompt
		fmt.Printf("%s > ", version())
		cmd, _ := cliReader.ReadString('\n')

		paction, ptask, pargs := cli.ProcessCommand(cmd)
		if pargs == nil {
			pargs = []string{""}
		}

		// catch special commands first
		var specialcmd, exitcmd bool

		switch paction {
			case "exit":
				exitcmd = true
				break
			case "help":
				specialcmd = true
				fmt.Println(help())
				break
		}
		// special commands take precedence
		if specialcmd {
			continue
		}
		// exit command quits
		if exitcmd {
			break
		}

		// catch unrecognized commands
		cmdOK := true
		if _, ok := actions.Actions[paction]; !ok {
			fmt.Println("unrecognized action")
			cmdOK = false
		}
		if _, ok := tasks.Tasks[ptask]; !ok {
			fmt.Println("unrecognized task")
			cmdOK = false
		}
		if !cmdOK {
			continue
		}

		// get action/task functions
		action := actions.Actions[paction]
		task := tasks.Tasks[ptask]

		// do action
		action(tc, task, pargs...)
	}

	// print witty parting-message
	fmt.Println("excelsior!")
}
