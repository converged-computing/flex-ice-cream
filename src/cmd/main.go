package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/converged-computing/flex-ice-cream/src/graph"
)

const (
// defaultspecFile = "icecream.yaml"
)

func main() {
	fmt.Println("This is the flex ice cream matcher")
	confFilePath := flag.String("conf", "conf/icecream.graphml", "icecream.graphml that describes the graph structure")
	specFilePath := flag.String("spec", "", "JobSpec (yaml file) that defines ice cream request")
	matchPolicy := flag.String("policy", "first", "Match policy (defaults to first)")
	flag.Parse()

	specFile := *specFilePath
	confFile := *confFilePath

	// The JobSpec file is required
	if specFile == "" {
		flag.Usage()
		os.Exit(0)
	}

	// The JobSpec file and graphml must exist
	if _, err := os.Stat(specFile); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("%s does not exist\n", specFile)
		os.Exit(0)
	}
	if _, err := os.Stat(confFile); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("%s does not exist\n", confFile)
		os.Exit(0)
	}

	// Create an ice cream graph, and match the spec to it.
	g := graph.FlexGraph{}
	g.Init(confFile, *matchPolicy, "")
	icecream, err := g.Order(specFile)
	if err != nil {
		fmt.Printf("Oh no! There was a problem with your ice cream request: %x", err)
		return
	}

	// Yay! Show the ice cream!
	icecream.Show()
}
