package drivercore

// Network describes a virtual network, via which a cluster of Machines are connected.
// It provides machine-to-machine private connectivity, and external network connectivity
// via routing or NAT. It may also manage DHCP.
type Network interface {
	Name() string
	CIDR() string

	SetCIDR(cidr string)
}
