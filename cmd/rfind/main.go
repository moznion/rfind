package main

import (
	"flag"
	"fmt"
	"github.com/moznion/rfind"
	"os"
	"path/filepath"
)

func main() {
	maxUpperDepth := flag.Uint("max-upper-depth", 0, "Max number of path depth to search from the ORIGIN_PATH to ancestor direction (default: unlimited)")
	maxDepthFromRoot := flag.Uint("max-depth-from-root", 0, "Max number of path depth to search from root (default: unlimited)")
	limit := flag.Uint("limit", 0, "Limit the number of found items (default: unlimited)")
	fileOnly := flag.Bool("file-only", false, "Find only file")
	dirOnly := flag.Bool("dir-only", false, "Find only directory")

	flag.Usage = customUsage

	flag.Parse()

	if !*fileOnly && !*dirOnly {
		*fileOnly = true
		*dirOnly = true
	}

	rf := rfind.Rfind{
		MaxUpwardDepth:   *maxUpperDepth,
		MaxDepthFromRoot: *maxDepthFromRoot,
		Limit:            *limit,
		IsFile:           *fileOnly,
		IsDir:            *dirOnly,
	}

	args := flag.Args()
	if len(args) < 2 {
		_, _ = fmt.Fprintf(os.Stderr, "[error] invalid arguments, please refer to the usage\n")
		flag.Usage()
		os.Exit(1)
	}

	originPath := args[0]
	if !filepath.IsAbs(originPath) {
		cwd, err := os.Getwd()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "[error] failed to get current working directory: %s\n", err)
			os.Exit(1)
		}
		originPath = filepath.Join(cwd, originPath)
	}

	found, err := rf.Find(originPath, args[1:])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[error] failed to find: %s\n", err)
		os.Exit(1)
	}
	for _, f := range found {
		fmt.Printf("%s\n", f)
	}
}

func customUsage() {
	cmd := os.Args[0]

	_, _ = fmt.Fprintf(os.Stderr, `%s: A command-line tool that searches for files in reverse order (i.e. to ancestor direction).
The result of the found files' paths ensures that they are ordered by the length of the match, with longer matches appearing first.
Usage:
  %s [OPTIONS] ORIGIN_PATH TARGETS...
Options:
`, cmd, cmd)
	flag.PrintDefaults()
}
