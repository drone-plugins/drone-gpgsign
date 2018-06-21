package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/mattn/go-zglob"
	"github.com/pkg/errors"
)

type (
	Config struct {
		Key        string
		Passphrase string
		Detach     bool
		Clear      bool
		Files      []string
		Excludes   []string
	}

	Plugin struct {
		Config Config
	}
)

func (p *Plugin) Exec() error {
	excludes := make([]string, 0)
	files := make([]string, 0)

	for _, name := range p.Config.Excludes {
		matches, err := zglob.Glob(name)

		if err != nil {
			return errors.Wrap(err, "failed to match excludes")
		}

		for _, match := range matches {
			if _, err := os.Stat(match); err == nil {
				excludes = append(excludes, match)
			}
		}
	}

	for _, name := range p.Config.Files {
		matches, err := zglob.Glob(name)

		if err != nil {
			return errors.Wrap(err, "failed to match files")
		}

		for _, match := range matches {
			excluded := false

			for _, exclude := range excludes {
				if match == exclude {
					excluded = true
					break
				}
			}

			if _, err := os.Stat(match); err == nil && !excluded {
				files = append(files, match)
			}
		}
	}

	if len(files) == 0 {
		log.Printf("no files found")
		return nil
	}

	key, err := base64.StdEncoding.DecodeString(p.Config.Key)

	if err != nil {
		key = []byte(p.Config.Key)
	}

	cmd := p.importKey()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = strings.NewReader(string(key))

	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "failed to import key")
	}

	for _, name := range files {
		cmd := p.signFile(name)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if p.Config.Passphrase != "" {
			cmd.Stdin = strings.NewReader(p.Config.Passphrase)
		}

		if err := cmd.Run(); err != nil {
			return errors.Wrap(err, "failed to sign file")
		}
	}

	return nil
}

func (p *Plugin) importKey() *exec.Cmd {
	args := []string{
		"--batch",
		"--import",
		"-",
	}

	fmt.Println("$ gpg", strings.Join(args, " "))

	return exec.Command(
		"gpg",
		args...,
	)
}

func (p *Plugin) signFile(name string) *exec.Cmd {
	args := []string{
		"--batch",
		"--yes",
		"--armor",
	}

	if p.Config.Passphrase != "" {
		args = append(args, "--pinentry-mode", "loopback", "--passphrase-fd", "0")
	}

	if p.Config.Detach {
		args = append(args, "--detach-sign")
	} else if p.Config.Clear {
		args = append(args, "--clear-sign")
	} else {
		args = append(args, "--sign")
	}

	args = append(args, name)

	fmt.Println("$ gpg", strings.Join(args, " "))

	return exec.Command(
		"gpg",
		args...,
	)
}
