package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func isGitRepo() error {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return errors.New(stderr.String()) //nolint:err113
	}

	return nil
}

// getStagedDirs returns a list of Go packages that have staged changes.
func getStagedDirs() ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-only", "--cached")

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	files := strings.Split(string(output), "\n")
	dirSet := make(map[string]struct{})

	for _, file := range files {
		if file != "" {
			dir := filepath.Base(filepath.Dir(file))
			if dir == "." {
				dir = "chore"
			}

			dirSet[dir] = struct{}{}
		}
	}

	dirs := make([]string, 0, len(dirSet))

	for dir := range dirSet {
		dirs = append(dirs, dir)
	}

	return dirs, nil
}

func getStagedChanges() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get diff: %w", err)
	}

	return string(output), nil
}
