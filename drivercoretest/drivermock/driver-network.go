package drivermock

import (
	"errors"

	"github.com/kuttiproject/drivercore"
)

var mocknetworks = map[string]*Network{}

// QualifiedNetworkName adds a suffix to the specified cluster name.
func (vd *Driver) QualifiedNetworkName(clustername string) string {
	return clustername + vd.networkNameSuffix
}

// ListNetworks gets a list of networks.
func (vd *Driver) ListNetworks() ([]drivercore.Network, error) {
	panic("not implemented") // Will not implement
}

// GetNetwork gets a network.
func (vd *Driver) GetNetwork(clustername string) (drivercore.Network, error) {
	result, ok := mocknetworks[vd.QualifiedNetworkName(clustername)]
	if !ok {
		return nil, errors.New("network not found")
	}
	return result, nil
}

// DeleteNetwork deletes a network.
func (vd *Driver) DeleteNetwork(clustername string) error {
	netname := vd.QualifiedNetworkName(clustername)
	if _, ok := mocknetworks[netname]; !ok {
		return errors.New("network not found")
	}
	delete(mocknetworks, netname)
	return nil
}

// NewNetwork creates a network for a cluster.
func (vd *Driver) NewNetwork(clustername string) (drivercore.Network, error) {
	qnname := vd.QualifiedNetworkName(clustername)
	_, ok := mocknetworks[qnname]
	if ok {
		return nil, errors.New("network already exists")
	}
	newnetwork := &Network{
		driver:      vd,
		clustername: clustername,
		cidr:        DefaultNetCIDR,
	}
	mocknetworks[qnname] = newnetwork
	return newnetwork, nil
}
