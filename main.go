package main

import (
	"os"

	"github.com/clarify/clarify-go/automation/automationcli"
)

func main() {
	os.Exit(automationcli.ParseAndRun(routines))
}
