// Package drivercore contains types, interfaces and functions that define kutti driver functionality.
// It also provides a central place for drivers to register themselves.
//
// The interfaces are:
//
// Driver
//
// This defines the interface for kutti "drivers". Each driver should be able to
// manage:
//
// - Machines, which represent Kubernetes nodes
//
// - Networks, which connect Machines and may manage DHCP and NAT
//
// - Images, which allow templated creation of Machines
//
// Implemented drivers should call the RegisterDriver function with a unique name on init.
//
// Network
//
// This defines a private network to which all nodes in a cluster will be connected.
// The network should allow connectivity between nodes, and public internet connectivity.
// For now, only IPv4 capability is assumed.
//
// Machine
//
// This defines a machine that will act as a Kubernetes node. The machine should allow start,
// stop, force stop, and wait operations, and provide a way to connect to it via SSH.
// It should also allow the execution of some predefined commands within its operating system,
// including:
//
// - RenameMachine
//
// - RestartMachine
//
// - CheckConnectivity
//
// - SetProxy
//
// - SetNoProxy
//
// - InitCluster
//
// - JoinCluster
//
// - InstallOverlayNetwork
//
// Image
//
// This defines a template from which a Machine can be created. An image has a property
// called K8sVersion , which specifies the version of kubernetes binaries present in
// Machines created from it. Drivers should include the functionality to download Images
// from driver-defined repositories, and check their integrity.
package drivercore
