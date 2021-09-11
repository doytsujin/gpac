package main

import (
	"fmt"
	"strings"
)

func main() {
	gconf(":url[https://github.com]\n" +
		":build[cd /tmp]")
}

func gconf(gconfs string) {
	for _, line := range strings.Split(strings.TrimRight(gconfs, "\n"), "\n") {

		if string(line[0:4]) == ":url" {
			url := line[:len(line)-1]
			url = url[5:]
			fmt.Println(url)
		}
		if string(line[0:6]) == ":build" {
			build := line[:len(line)-1]
			build = build[7:]
			fmt.Println(build)
		}
	}
}
