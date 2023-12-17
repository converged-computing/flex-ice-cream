package graph

import (
	"errors"
	"os"

	"github.com/converged-computing/flex-ice-cream/src/icecream"

	// TODO update back to flux-sched when merged
	"github.com/researchapps/flux-sched/resource/reapi/bindings/go/src/fluxcli"

	"fmt"
)

/*
Desired steps:

1. instantiate Fluxion
2. Create the context, pass in the graphml and specify instead of JFG we are using graphML
3. Then the defaults will work out of box
4. Then pass in a jobspec
5. Perform a satisfiability check rather than a match "can we represent this or not" There exist one or more matches for the schema in the resource graph.
6. Would we need/want a match? Is it possible to validate the schema without performing a match.

*/

type FlexGraph struct {
	cli *fluxcli.ReapiClient
}

// Init a new FlexGraph from a graphml filename
func (f *FlexGraph) Init(confFile string, matchPolicy string, label string) {

	// 1. instantiate fluxion
	f.cli = fluxcli.NewReapiClient()
	fmt.Printf("Created flex resource graph %s\n", f.cli)

	// 2. Load in the graphml GRUG file
	conf, err := os.ReadFile(confFile)
	if err != nil {
		fmt.Println("Error reading graphml file")
		return
	}

	// Set match policy to default (first) if not defined.
	// In practice this should not happen - the cmd/main.go sets a default.
	if matchPolicy == "" {
		matchPolicy = "first"
	}

	// Alert the user to all the chosen parameters
	// Note that "grug" == "graphml" but probably nobody knows what grug means
	fmt.Printf(" Match policy: %s\n", matchPolicy)
	fmt.Println(" Load format: graphml (grug)")
	fmt.Printf(" Config file: %s\n", confFile)

	// 2. Create the context, specify instead of JGF (default) we want graphml
	// 3. Remainder of defaults should work out of the box
	// Note that the options get passed as a json string to here:
	// https://github.com/flux-framework/flux-sched/blob/master/resource/reapi/bindings/c%2B%2B/reapi_cli_impl.hpp#L412
	opts := `{"matcher_policy": "%s", "load_file": "%s", "load_format": "grug", "match_format": "simple"}`
	p := fmt.Sprintf(opts, matchPolicy, confFile)

	// 4. Then pass in a jobspec... err, ice cream request :)
	err = f.cli.InitContext(string(conf), p)
	if err != nil {
		fmt.Printf("Error creating context: %x", err)
	}
	fmt.Printf("\n✨️ Init context complete!\n")
}

// Order is akin to doing a Satisfies, but right now it's a MatchAllocate
// The result of an order is a traversal of the graph that could satisfy the ice cream request
func (f *FlexGraph) Order(specFile string) (icecream.IceCreamRequest, error) {
	fmt.Printf("   🍦️ Request: %s\n", specFile)

	// Prepare the ice cream request!
	request := icecream.IceCreamRequest{}

	spec, err := os.ReadFile(specFile)
	if err != nil {
		return request, errors.New("Error reading jobspec")
	}

	// TODO this could be f.cli.Satisfies
	// Note that number originally was a jobid (it's now a number for the ice cream in the shop)
	// Note that recipe was originally "allocated"
	_, recipe, _, _, number, err := f.cli.MatchAllocate(false, string(spec))
	if err != nil {
		return request, err
	}

	// Populate the ice cream request for the customer
	request.Recipe = recipe
	request.Number = number
	return request, nil
}
