package main

import (
	"github.com/nickswift498/birdwatch/config"
	"github.com/nickswift498/birdwatch/client"
	"github.com/nickswift498/birdwatch/cli"
	"github.com/nickswift498/birdwatch/cli/actions"
	"github.com/nickswift498/birdwatch/cli/tasks"
	"fmt"
	"bufio"
	"os"
)

func title() {
	fmt.Println(``)
	fmt.Println(`+------------------------------------------------------------------------+`)
	fmt.Println(`| ██████╗ ██╗██████╗ ██████╗ ██╗    ██╗ █████╗ ████████╗ ██████╗██╗  ██╗ |`)
	fmt.Println(`| ██╔══██╗██║██╔══██╗██╔══██╗██║    ██║██╔══██╗╚══██╔══╝██╔════╝██║  ██║ |`)
	fmt.Println(`| ██████╔╝██║██████╔╝██║  ██║██║ █╗ ██║███████║   ██║   ██║     ███████║ |`)
	fmt.Println(`| ██╔══██╗██║██╔══██╗██║  ██║██║███╗██║██╔══██║   ██║   ██║     ██╔══██║ |`)
	fmt.Println(`| ██████╔╝██║██║  ██║██████╔╝╚███╔███╔╝██║  ██║   ██║   ╚██████╗██║  ██║ |`)
	fmt.Println(`| ╚═════╝ ╚═╝╚═╝  ╚═╝╚═════╝  ╚══╝╚══╝ ╚═╝  ╚═╝   ╚═╝    ╚═════╝╚═╝  ╚═╝ |`)
	fmt.Println(`+------------------------------------------------------------------------+`)
	fmt.Println(``)
}

func version() string {
	return fmt.Sprintf("birdwatch v%d.%d", config.VERSION_MAJOR, config.VERSION_MINOR)
}
func help() string {
	return ``
}

func main() {

	// show title
	title()

	// retrieve twitter client
	tc := client.TwitterClientFromEnv()

	// collect user input
	cliReader := bufio.NewReader(os.Stdin)
	for {
		// print prompt
		fmt.Printf("birdwatch > ")

		// catch input and do some throat clearing
		cmd, _ := cliReader.ReadString('\n')
		paction, _, pargs := cli.ProcessCommand(cmd)
		if pargs == nil {
			pargs = []string{""}
		}

		// catch special commands first
		var specialcmd, exitcmd bool

		switch paction {
			case "exit":
				exitcmd = true
				break
			case "version":
				fmt.Println(version())
				specialcmd = true
			case "help":
				fmt.Println(help())
				specialcmd = true
		}

		// special commands take precedence
		if specialcmd {
			continue
		}
		// exit command quits
		if exitcmd {
			break
		}

		// process command
		action, task, args := cli.ProcessCommand(cmd)

		// catch unrecognized commands
		cmdOK := true
		if _, ok := actions.Actions[action]; !ok {
			fmt.Println("unrecognized action")
			cmdOK = false
		}
		if _, ok := tasks.Tasks[task]; !ok {
			fmt.Println("unrecognized task")
			cmdOK = false
		}
		if !cmdOK {
			continue
		}

		actionfn := actions.Actions[action]
		taskfn := tasks.Tasks[task]

		// do action
		actionfn(tc, taskfn, args...)
	}

	fmt.Println("Excelsior!")
}
