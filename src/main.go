package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Biisairo/sugo/src/cmdbuild"
	"github.com/Biisairo/sugo/src/cmdstart"
)

var helpCommand = `SuGo - static web site generator

Usage:
	sugo <command> [arguments]

Commands:
	start	Create default template to build webpage
	build	Generate static webpage
	run	Build and run webpage

Flags:
	-root	string
		Root derectory for build contents (Default .)
	-config	string
		Config file path (Default config.toml)

Use "sugo --help" or "sugo -h" for more information about SuGo.
`

func main() {
	configFile := "config.toml"
	rootPath := ""
	showHelp := false

	flag.StringVar(&configFile, "config", configFile, "")
	flag.StringVar(&configFile, "c", configFile, "")
	flag.StringVar(&rootPath, "root", rootPath, "")
	flag.StringVar(&rootPath, "r", rootPath, "")
	flag.BoolVar(&showHelp, "help", showHelp, "")
	flag.BoolVar(&showHelp, "h", showHelp, "")
	flag.Parse()

	command := ""
	if 1 < len(os.Args) {
		command = os.Args[1]
	}

	if showHelp || (command != "start" && command != "build" && command != "run") {
		fmt.Print(helpCommand)
		return
	}

	if command == "start" {
		if err := cmdstart.Start(rootPath); err != nil {
			log.Fatal("Error creating site structure: ", err)
		}
		return
	}

	cmdbuild.Build(rootPath, configFile)

	if command == "run" {
		// 	// setup fsnotify watcher
		// 	go watchDirs([]string{
		// 		filepath.Join(rootPath, "content"),
		// 		filepath.Join(rootPath, "public"),
		// 		filepath.Join(rootPath, "templates"),
		// 	}, func() {
		// 		cmdbuild.Build(rootPath, configFile)
		// 	})

		// 	// serve site
		log.Default().Print("Listening on http://localhost:8080\n")
		_ = http.ListenAndServe("localhost:8080", http.FileServer(http.Dir("build")))
	}
}
