package drivercore

// Network describes a virtual network, via which a cluster of Machines
// are connected. It provides machine-to-machine private connectivity,
// and external network connectivity via routing or NAT.
// It may also manage DHCP.
type Network interface {
	// Name is the name of the network.
	Name() string
	// CIDR is the network's IPv4 address range.
	CIDR() string

	// SetCIDR sets a new IPv4 address range, in CIDR format.
	SetCIDR(cidr string)
}
