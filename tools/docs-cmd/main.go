package main

import (
	"log"
	"os"

	"github.com/spf13/cobra/doc"

	"k8s.io/kubectl/pkg/cmd"
)

func main() {
	// set HOME env var so that default values involve user's home directory do not depend on the running user.
	os.Setenv("HOME", "/home/user")
	os.Setenv("XDG_CONFIG_HOME", "/home/user/.config")

	err := doc.GenMarkdownTree(cmd.NewDefaultKubectlCommand(), "./docs/commands")
	if err != nil {
		log.Fatal(err)
	}
}
