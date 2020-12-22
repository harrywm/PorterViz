package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func refreshClient() *client.Client {

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	return cli
}

func networkList(cli *client.Client) {

	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		panic(err)
	}

	// THIS WORKS FOR LISTING FULL INSPECT OF ALL NETWORKS
	for _, network := range networks {
		//fmt.Println(network.ID)
		networkInspect, err := cli.NetworkInspect(context.Background(), network.ID, types.NetworkInspectOptions{})
		if err != nil {
			panic(err)
		}
		if len(networkInspect.Containers) != 0 {
			fmt.Println(networkInspect.Containers)
		}
	}
}

func main() {

	networkList(refreshClient())

}
