package main

import (
	"github.com/nickswift498/birdwatch/config"
	"github.com/nickswift498/birdwatch/cli"
	"fmt"
	"bufio"
	"os"
)

func title() {
	fmt.Println(`       __     __                          `)
	fmt.Println(`  ____/ /__  / /___  ________  ____ _____ `)
	fmt.Println(" / __  / _ \\/ / __ \\/ ___/ _ \\/ __ `/ __ \\")
	fmt.Println(`/ /_/ /  __/ / /_/ / /  /  __/ /_/ / / / /`)
	fmt.Println(`\__,_/\___/_/\____/_/   \___/\__,_/_/ /_/ `)
	fmt.Println(`                                          `)
	fmt.Println(`           It's 1.21 jiggawats!           `)
	fmt.Println(``)
}

func version() string {
	return fmt.Sprintf("delorean v%d.%d", config.DELOR_VERSION_MAJOR, config.DELOR_VERSION_MINOR)
}

func main() {
	title()

	// retrieve twitter client
	tc := client.TwitterClientFromEnv()

	// before we start, we need to build the user's history
	fmt.Println("No user history detected. Building...")
	
	// collect user input
	cliReader := bufio.NewReader(os.Stdin)
	for {
		// print prompt
		fmt.Printf("%s > ", version())

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
		}

		// special commands take precedence
		if specialcmd {
			continue
		}
		// exit command quits
		if exitcmd {
			break
		}
	}

	fmt.Println(`Great Scott!`)
}
