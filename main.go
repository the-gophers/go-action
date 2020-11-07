package main

import (
	"fmt"
	"log"
	"os"

	flag "github.com/spf13/pflag"
)

var (
	sample string
)

func main() {
	parseAndValidateInput()
	fmt.Printf("sample was %q\n", sample)
	if os.Getenv("DRY_RUN") != "true" {
		// do something with your action
		PrintOutput("sampleOutput", "env var DRY_RUN was false or not specified")
		return
	}

	PrintOutput("sampleOutput", "dry run was true")
}

func parseAndValidateInput() {
	flag.StringVar(&sample, "sample", "", "some sample input")
	flag.Parse()

	if sample == "" {
		log.Fatal("--sample can't be empty")
	}
}

// PrintOutput prints formatted output to stdout so that GitHub Actions runtime will associate it with the output
// variable name.
func PrintOutput(key, message string) {
	fmt.Printf("::set-output name=%s::%s\n", key, message)
}
