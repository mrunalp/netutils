package netutils

import (
	"testing"
)

func TestAllocateSubnet(t *testing.T) {
	sna, err := NewSubnetAllocator("10.1.0.0/16", 8, nil)
	if err != nil {
		t.Fatal("Failed to initialize IP allocator: ", err)
	}

	sn, err := sna.GetNetwork()
	if err != nil {
		t.Fatal("Failed to get network: ", err)
	}
	if sn.String() != "10.1.0.0/24" {
		t.Fatal("Did not get expected subnet")
	}
	sn, err = sna.GetNetwork()
	if err != nil {
		t.Fatal("Failed to get network: ", err)
	}
	if sn.String() != "10.1.1.0/24" {
		t.Fatal("Did not get expected subnet")
	}
	sn, err = sna.GetNetwork()
	if err != nil {
		t.Fatal("Failed to get network: ", err)
	}
	if sn.String() != "10.1.2.0/24" {
		t.Fatal("Did not get expected subnet")
	}
}
