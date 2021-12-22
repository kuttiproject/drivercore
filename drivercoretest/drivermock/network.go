package drivermock

// Network implements the drivercore.Network interface
type Network struct {
	driver      *Driver
	clustername string
	cidr        string
}

// Name returns the driver-qualified name of the network.
func (vd *Network) Name() string {
	return vd.driver.QualifiedNetworkName(vd.clustername)
}

// CIDR returns the CIDR of the network.
func (vd *Network) CIDR() string {
	return vd.cidr
}

// SetCIDR has not been implemented for this mock.
func (vd *Network) SetCIDR(cidr string) {
	panic("not implemented") // TODO: Implement
}
