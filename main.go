package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/kr/pretty"
)

var cli = refreshClient()
var containers = containerList(cli)
var networks = networkList(cli)

func refreshClient() *client.Client {

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	return cli
}

func containerList(cli *client.Client) []types.Container {

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	return containers
}

func networkList(cli *client.Client) []types.NetworkResource {

	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		panic(err)
	}

	return networks
}

func getByNetwork(network types.NetworkResource) []types.Container {
	var netContainers []types.Container

	networkInspect, err := cli.NetworkInspect(context.Background(), network.ID, types.NetworkInspectOptions{})
	if err != nil {
		panic(err)
	}

	// If Network has Containers
	if len(networkInspect.Containers) != 0 {
		// For each container in the network
		for _, container := range networkInspect.Containers {
			// Retrieve full Inspect of Container by Name from Global Container List
			for i := range containers {
				if containers[i].Names[0][1:] == container.Name {
					netContainers = append(netContainers, containers[i])
				}
			}
		}
	}

	fmt.Printf("%# v", pretty.Formatter(netContainers))
	return netContainers

}

func main() {

	getByNetwork(networks[0])

}
