package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	if os.Getenv("OPENAI_API_KEY") == "" {
		os.Stderr.WriteString("OPENAI_API_KEY environment variable is not set\n")
		os.Exit(1)
	}

	noPackagesFlag := flag.Bool("no-pkg", false, "exclude package name from commit message")
	noCmdFlag := flag.Bool("no-cmd", false, "keep only commit message")
	flag.Parse()

	if err := isGitRepo(); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	diff, err := getStagedChanges()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	if diff == "" {
		os.Stdout.WriteString("no staged changes\n")
		os.Exit(0)
	}

	msg, err := generateCommitMessage(diff, os.Getenv("OPENAI_API_KEY"))
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	dirs, err := getStagedDirs()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	str := fmt.Sprintf("git commit -m \"%s: %s\"\n", strings.Join(dirs, ", "), msg)

	switch {
	case *noPackagesFlag && !*noCmdFlag:
		str = fmt.Sprintf("git commit -m \"%s\"\n", msg)

	case !*noPackagesFlag && *noCmdFlag:
		str = fmt.Sprintf("%s: %s\n", strings.Join(dirs, ", "), msg)

	case *noPackagesFlag && *noCmdFlag:
		str = msg + "\n"
	}

	os.Stderr.WriteString(str)
}
