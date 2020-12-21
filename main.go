package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

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

	/* 	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
		if err != nil {
			panic(err)
		}

	 	for _, container := range containers {

			networks := (container.NetworkSettings.Networks)

			for _, net := range networks {
				// CAN LISTEN ON THIS  --->
				//fmt.Println(net, "\n \n")
			}
		}  */
}
