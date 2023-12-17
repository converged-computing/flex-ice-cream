# Flex Ice Cream

> Need to figure our your ice cream supply? Just Flex it! 💪️™️🍦️

The Flux Framework "flux-sched" or fluxion project provides modular bindings in different languages for intelligent,
graph-based scheduling. When we extend fluxion to a tool or project that warrants logic of this type, we call this a flex!
Thus, the project here demonstrates flex-ice-cream, or using fluxion to:

1. Generate a [graph schema](conf/icecream.graphml) (graphml) that describes nodes and edges of an ice cream shop
2. Create an example [ice cream cone or cup](icecream.yaml) that a customer wants
3. Use the tool to determine if the shop can provide it!

This is a simple example intended for learning and fun. 

## Concepts

From the above, the following definitions might be useful.

 - **[Flux Framework](https://flux-framework.org)**: a modular framework for putting together a workload manager. It is traditionally for HPC, but components have been used in other places (e.g., here, Kubernetes, etc). It is analogous to Kubernetes in that it is modular and used for running batch workloads.
 - **[fluxion](fluxion)**: refers to [flux-framework/flux-sched](https://github.com/flux-framework/flux-sched) and is the scheduler component or module of Flux Framework. There are bindings in several langauges, and specifically the Go bindings (server at [flux-framework/flux-k8s](https://github.com/flux-framework/flux-k8s)) assemble into the project "fluence."
 - **flex** is an out of tree tool, plugin, or similar that uses fluxion to flexibly schedule or match some kind of graph-based resources. This project is an example of a flex!

## Usage

### Build

This demonstrates how to build the bindings. You will need to be in the VSCode developer container environment, or produce the same
on your host. Note that we currently are using [this commit](https://github.com/researchapps/flux-sched/commit/86f5bb331342f2883b057920cf58e2c042aef881) that
si a fork of [milroy's work](https://github.com/flux-framework/flux-sched/pull/1120) to ensure the module name matches what is added to go.mod (it won't work otherwise). When this is merged, we will update to flux-framework/flux-sched. Below shows the make command that builds our final binary!

```bash
make
```
```console
# This needs to match the flux-sched install and latest commit, for now we are using a fork of milroy's branch
# that has a go.mod updated to match the org name
# go get -u github.com/researchapps/flux-sched/resource/reapi/bindings/go/src/fluxcli@86f5bb331342f2883b057920cf58e2c042aef881
go mod tidy
mkdir -p ./bin
GOOS=linux CGO_CFLAGS="-I/opt/flux-sched/resource/reapi/bindings/c" CGO_LDFLAGS="-L/usr/lib -L/opt/flux-sched/resource -lfluxion-resource -L/opt/flux-sched/resource/libjobspec -ljobspec_conv -L//opt/flux-sched/resource/reapi/bindings -lreapi_cli -lflux-idset -lstdc++ -lczmq -ljansson -lhwloc -lboost_system -lflux-hostlist -lboost_graph -lyaml-cpp" go build -ldflags '-w' -o bin/icecream src/cmd/main.go
```

The output is generated in bin:

```bash
$ ls bin/
icecream
```

Also note that we are targeting header files and shared libraries that are still in the flux source code directory. This is necessary because, due to licenses,
we are not allowed to distribute them proper. This is why we provide the Developer environment here that has this build ready to go.

### Run

You can provide your request for ice cream (e.g., icecream.yaml) and the description of the graph (in graphml). Note that we need shared libs on the path:

```
export LD_LIBRARY_PATH=/usr/lib:/opt/flux-sched/resource:/opt/flux-sched/resource/reapi/bindings:/opt/flux-sched/resource/libjobspec
```

-L/usr/lib -L/opt/flux-sched/resource -lfluxion-resource -L/opt/flux-sched/resource/libjobspec -ljobspec_conv -L/opt/flux-sched/resource/reapi/bindings -lreapi_cli -lflux-idset -lstdc++ -lczmq -ljansson -lhwloc -lboost_system -lflux-hostlist -lboost_graph -lyaml-cpp" go build -ldflags '-w' -o bin/icecream src/cmd/main.go

```bash
./bin/icecream -spec icecream.yaml
```
```console
This is the flex ice cream matcher
Created flex resource graph &{%!s(*fluxcli.ReapiCtx=&{})}
 Match policy: first
 Load format: graphml (grug)
 Config file: conf/icecream.graphml

✨️ Init context complete!
   🍦️ Request: icecream.yaml

😍️ Your Ice Cream Order is Ready!
Order Number: 1
Recipe:
      ---------scoop0[4:x]
      ------cup0[1:s]
      ---ice-cream-shop0[1:s]
```

I'm not actually sure if that is upside down? Note that we need to adjust the above to be something like `Satisfies` instead of `MatchAllocate`, so that is TBA.
You can also customize the graphml input file:

```bash
./bin/icecream -spec icecream.yaml -conf ./config/icecream.graphml
```

Try an order that can't be satisfied (note this is a bug on my part, you should be able to ask for a cone, I'm just not sure how to do it yet):

```bash
./bin/icecream -spec examples/cone.yaml
```
```console
This is the flex ice cream matcher
Created flex resource graph &{%!s(*fluxcli.ReapiCtx=&{})}
 Match policy: first
 Load format: graphml (grug)
 Config file: conf/icecream.graphml

✨️ Init context complete!
   🍦️ Request: examples/cone.yaml

😭️ Oh no, we could not satisfy your order!
```


## License

HPCIC DevTools is distributed under the terms of the MIT license.
All new contributions must be made under this license.

See [LICENSE](https://github.com/converged-computing/cloud-select/blob/main/LICENSE),
[COPYRIGHT](https://github.com/converged-computing/cloud-select/blob/main/COPYRIGHT), and
[NOTICE](https://github.com/converged-computing/cloud-select/blob/main/NOTICE) for details.

SPDX-License-Identifier: (MIT)

LLNL-CODE- 842614
