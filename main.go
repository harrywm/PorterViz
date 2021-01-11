package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/kr/pretty"
)

// PorterContainer is an abstracted version of the Docker data type for a Container containing only the fields we are concerned with
type PorterContainer struct {
	Name    []string
	ID      string
	Image   string
	ImageID string
	Ports   []types.Port
	State   string
}

// PorterNetwork is an abstracted version of the Docker Network data type containing only the fields we are concerned with
type PorterNetwork struct {
	Name       string
	ID         string
	Driver     string
	Containers []PorterContainer
}

func (n *PorterNetwork) getInfo() string {
	return n.Name
}

func (n *PorterNetwork) getID() string {
	return n.ID
}

func (n *PorterNetwork) getDriver() string {
	return n.Driver
}

func (n *PorterNetwork) getContainers() []PorterContainer {
	return n.Containers
}

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
	return netContainers
}

func allNetworks() []PorterNetwork {

	var returnNets []PorterNetwork = []PorterNetwork{}
	var net PorterNetwork = PorterNetwork{}
	var cont []PorterContainer = []PorterContainer{}

	for _, network := range networks {

		currentContainer := getByNetwork(network)

		for _, container := range currentContainer {

			cont = append(cont,
				PorterContainer{
					Name:    container.Names,
					ID:      container.ID,
					Image:   container.Image,
					ImageID: container.ImageID,
					Ports:   container.Ports,
					State:   container.State,
				})
		}

		net = PorterNetwork{
			Name:       network.Name,
			ID:         network.ID,
			Driver:     network.Driver,
			Containers: cont,
		}

		returnNets = append(returnNets, net)
	}

	return returnNets
}

func main() {

	networks := allNetworks()

	conts := networks[0].getContainers()

	for _, cont := range conts {

		contInsp, err := cli.ContainerInspect(context.Background(), cont.ID)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%# v", pretty.Formatter(contInsp))

	}

	fmt.Printf("%# v", pretty.Formatter(networks))

	//js := mewn.String("./frontend/build/static/js/main.js")
	//css := mewn.String("./frontend/build/static/css/main.css")

	/* 	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "PORTER",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	}) */

	//app.Bind()
	//app.Run()
}
