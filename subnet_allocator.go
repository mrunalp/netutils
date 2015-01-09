package netutils

import (
	"fmt"
	"net"
)

type SubnetAllocator struct {
	network  *net.IPNet
	capacity int
	allocMap map[string]bool
}

func NewSubnetAllocator(network string, capacity int, amap map[string]bool) (*SubnetAllocator, error) {
	_, netIP, err := net.ParseCIDR(network)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse network address: %q", network)
	}

	netMaskSize, _ := netIP.Mask.Size()
	fmt.Println(netMaskSize)
	if capacity > (32 - netMaskSize) {
		return nil, fmt.Errorf("hosts bits cannot be larger than available bits")
	}

	if amap == nil {
		amap = make(map[string]bool)
	}
	return &SubnetAllocator{network: netIP, allocMap: amap}, nil
}

func (sna *SubnetAllocator) GetNetwork() (*net.IPNet, error) {
	fmt.Println(sna.network)
	return sna.network, nil
}
