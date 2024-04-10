package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var Version string = "UNKNOWN"

func main() {
	var wantVersion bool
	flag.BoolVar(&wantVersion, "version", false, "show version")

	var autoDetect bool
	flag.BoolVar(&autoDetect, "auto", false, "enable only when running in GitHub Actions")

	var title string
	flag.StringVar(&title, "title", "", "the group title, defaults to the command")

	flag.Parse()

	if wantVersion {
		fmt.Fprintf(os.Stdout, "actgroup %s\n", Version)
		os.Exit(0)
	}

	run := NewRun(autoDetect, title, flag.Args())

	if run.Command == "" {
		flag.Usage()
		os.Exit(1)
	}

	cmd := exec.Command(run.Command, run.Args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if run.Enabled() {
		// https://docs.github.com/en/actions/using-workflows/workflow-commands-for-github-actions?tool=bash#grouping-log-lines
		fmt.Fprintf(os.Stdout, "::group::%s\n", run.Title)
	}

	exitCode := 0

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\n%s\n", err)
		exitCode := cmd.ProcessState.ExitCode()
		fmt.Fprintf(os.Stderr, "\nexit code %d\n", exitCode)
	}

	if run.Enabled() {
		fmt.Fprintf(os.Stdout, "::endgroup::\n")
	}
	os.Exit(exitCode)
}

type Run struct {
	Title      string
	Command    string
	Args       []string
	AutoDetect bool
}

func (r *Run) Enabled() bool {
	if !r.AutoDetect {
		return true
	}

	return InGitHubActions()
}

func NewRun(autoDetect bool, title string, rest []string) *Run {
	run := &Run{
		AutoDetect: autoDetect,
		Title:      title,
	}

	if len(rest) > 0 {
		run.Command = rest[0]
	}
	if len(rest) > 1 {
		run.Args = rest[1:]
	}

	if run.Title == "" {
		run.Title = run.Command
	}

	return run
}

func InGitHubActions() bool {
	// https://docs.github.com/en/actions/learn-github-actions/variables#default-environment-variables
	return os.Getenv("GITHUB_ACTIONS") == "true"
}
