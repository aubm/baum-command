package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var deep = flag.Int("d", 0, "Depth")
var commandArgs []string

func main() {
	flag.Parse()
	commandArgs = flag.Args()

	dirToScan := getDirToScan()
	filepath.Walk(dirToScan, func(p string, f os.FileInfo, err error) error {
		p = strings.Replace(p, dirToScan, ".", 1)
		dir, _ := path.Split(p)
		nbOfParts := len(strings.Split(dir, "/"))
		if nbOfParts < (*deep+2) || *deep == 0 {
			fmt.Println(p)
		}
		return nil
	})
}

func getDirToScan() string {
	var dir string
	if len(commandArgs) > 0 {
		dir = commandArgs[0]
		dir, _ = filepath.Abs(dir)
	} else {
		dir, _ = os.Getwd()
	}

	src, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Source does not exist")
			os.Exit(1)
		} else if os.IsPermission(err) {
			fmt.Println("Unauthorized source")
			os.Exit(1)
		} else {
			panic(err)
		}
	}

	if !src.IsDir() {
		fmt.Println("Source is not a directory")
		os.Exit(1)
	}

	return dir
}
