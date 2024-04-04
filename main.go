package main

import (
	"fmt"

	"github.com/vishvananda/netlink"
)

func main() {

	routes, err := netlink.RouteList(nil, netlink.FAMILY_V4)
	if err != nil {
		fmt.Print("something went wrong")
	}

	for _, v := range routes {
		fmt.Print(v)
	}
}