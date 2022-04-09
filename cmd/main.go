package main

import (
	"EmoSearch/core"
	"flag"
	"fmt"
	"os"
)

func main() {
	Banner()
	flag.Parse()
	if *core.URL != "" {
		fmt.Println("Searching")
		core.Run1(*core.URL)
	}

	if *core.Urllist != "" {
		fmt.Println("Searching")

		core.Runs()

	}
	core.Wg.Wait()

	if *core.URL == "" && *core.Urllist == "" {
		fmt.Println("Please input -url or -file to lock target")
		os.Exit(0)
	}

	fmt.Println("done")

}

func Banner() {
	banner := `
    _____                      __
    /    '                   /    )                                /
---/__-------_--_-----__-----\---------__-----__----)__-----__----/__-
  /         / /  )  /   )     \      /___)  /   )  /   )  /   '  /   )
_/____     / /  /  (___/  (____/    (___   (___(_ /      (___   /   / 
												by Ch1nfo ver:1.0
	`
	print(banner)
}
