package netutils

import (
	"testing"
)

func TestAllocateSubnet(t *testing.T) {
	sna, err := NewSubnetAllocator("10.1.2.0/24", 8, nil)
	if err != nil {
		t.Fatal("Failed to initialize IP allocator: %v", err)
	}
	t.Log(sna.GetNetwork())
}
