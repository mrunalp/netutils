package netutils

import (
	"fmt"
	"net"
)

type SubnetAllocator struct {
	network  *net.IPNet
	capacity uint
	allocMap map[string]bool
}

func NewSubnetAllocator(network string, capacity uint, amap map[string]bool) (*SubnetAllocator, error) {
	_, netIP, err := net.ParseCIDR(network)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse network address: %q", network)
	}

	netMaskSize, _ := netIP.Mask.Size()
	fmt.Println(netMaskSize)
	if capacity > (32 - uint(netMaskSize)) {
		return nil, fmt.Errorf("Subnet capacity cannot be larger than number of networks available")
	}

	if amap == nil {
		amap = make(map[string]bool)
	}
	return &SubnetAllocator{network: netIP, capacity: capacity, allocMap: amap}, nil
}

func (sna *SubnetAllocator) GetNetwork() (*net.IPNet, error) {
	var (
		numSubnets    uint32
		numSubnetBits uint
	)
	baseipu := IPToUint32(sna.network.IP)
	netMaskSize, _ := sna.network.Mask.Size()
	numSubnetBits = 32 - uint(netMaskSize) - sna.capacity
	numSubnets = 1 << numSubnetBits

	var i uint32
	for i = 0; i < numSubnets; i++ {
		shifted := i << sna.capacity
		ipu := baseipu | shifted
		genIp := Uint32ToIP(ipu)
		genSubnet := &net.IPNet{IP: genIp, Mask: net.CIDRMask(int(numSubnetBits)+netMaskSize, 32)}
		if !sna.allocMap[genSubnet.String()] {
			sna.allocMap[genSubnet.String()] = true
			return genSubnet, nil
		}
	}

	return nil, fmt.Errorf("No subnets available.")
}
