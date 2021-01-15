package networks

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
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

// GetInfo ... Returns name of Docker Network object
func (n *PorterNetwork) GetInfo() string {
	return n.Name
}

// GetID ... Returns ID of Docker Network object
func (n *PorterNetwork) GetID() string {
	return n.ID
}

// GetDriver ... Returns Network Driver of Docker Network Object
func (n *PorterNetwork) GetDriver() string {
	return n.Driver
}

// GetContainers ... Returns a list of Containers connected to the Docker Network Object
func (n *PorterNetwork) GetContainers() []PorterContainer {
	return n.Containers
}

// ConnectContainer ... Connect Container to Network using ContainerID
func (n *PorterNetwork) ConnectContainer(containerID string) {

	err := RefreshClient().NetworkConnect(context.Background(), n.ID, containerID, &network.EndpointSettings{})

	if err != nil {
		panic(err)
	}
}

// DisconnectContainer ... Disconnect Container from Network using ContainerID
func (n *PorterNetwork) DisconnectContainer(containerID string, force bool) {

	err := RefreshClient().NetworkDisconnect(context.Background(), n.ID, containerID, force)

	if err != nil {
		panic(err)
	}
}

var cli = RefreshClient()
var containers = ContainerList(cli)
var networks = NetworkList(cli)

// RefreshClient ... Refreshes the Docker client instance. Returns refreshed Docker client object
func RefreshClient() *client.Client {

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	return cli
}

// ContainerList ... Returns full list of Containers from Docker Client instance
func ContainerList(cli *client.Client) []types.Container {

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	return containers
}

// NetworkList ... Returns full list of Networks from Docker Client instance
func NetworkList(cli *client.Client) []types.NetworkResource {

	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		panic(err)
	}
	return networks
}

// GetByNetwork ... Returns Containers connected to specified Network
func GetByNetwork(network types.NetworkResource) []types.Container {
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

// AllNetworks ... Returns all Networks from current Docker Client in Porter Object form
func AllNetworks() []PorterNetwork {

	var returnNets []PorterNetwork = []PorterNetwork{}
	var net PorterNetwork = PorterNetwork{}
	var cont []PorterContainer = []PorterContainer{}

	for _, network := range networks {

		currentContainer := GetByNetwork(network)

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
