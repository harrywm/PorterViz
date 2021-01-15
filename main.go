package main

import (
	"context"
	"fmt"

	"porter/donut/gui"
	"porter/donut/networks"

	"github.com/kr/pretty"
)

func setup() {
	networkList := networks.AllNetworks()

	conts := networkList[0].GetContainers()

	for _, cont := range conts {

		contInsp, err := networks.RefreshClient().ContainerInspect(context.Background(), cont.ID)

		if err != nil {
			panic(err)
		}

		fmt.Println(pretty.Formatter(contInsp.Name))

	}
	//fmt.Printf("%# v", pretty.Formatter(networkList))
}

func main() {

	/* 	for {
		setup()
		time.Sleep(time.Second)
	} */

	gui.Main()

}
