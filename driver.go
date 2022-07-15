package drivercore

// Driver describes operations to manage Networks, Machines and Images.
// The UsesPerClusterNetworking method returns true if the driver supports
// creation of Networks per cluster.
// The UsesNATNetworking() method returns true if the Networks of the driver
// use NAT. This means that ports of the Machines will need to be forwarded
// to physical host ports for host access.
// A driver is also responsible for maintaining a local cache of Images.
// It can update a list of published images from a driver-defined source,
// download local copies of images, and delete local copies (but not the
// image itself).
type Driver interface {
	// Name returns the unique name for a Driver. It should be short.
	Name() string
	// Description returns a short description of the Driver.
	Description() string
	// UsesPerClusterNetworking should return true if the Driver supports
	// creation of isolated Networks (NAT or routed) per cluster.
	UsesPerClusterNetworking() bool
	// UsesNATNetworking should return true if the networking used by the
	// Driver implements NAT, and therefore Machine ports need forwarding.
	UsesNATNetworking() bool
	// Status returns the current status of the Driver.
	Status() string
	// Error returns the last error reported by the Driver.
	Error() string

	// QualifiedNetworkName returns a unique Network name for a cluster.
	QualifiedNetworkName(clustername string) string
	ListNetworks() ([]Network, error)
	GetNetwork(clustername string) (Network, error)
	// DeleteNetwork deletes the Network for a cluster.
	DeleteNetwork(clustername string) error
	// NewNetwork creates a new Network for a cluster.
	NewNetwork(clustername string) (Network, error)

	// QualifiedMachineName returns a unique name for a Machine in a cluster.
	// This name is usually used internally by a Driver.
	QualifiedMachineName(machinename string, clustername string) string
	ListMachines() ([]Machine, error)
	// GetMachine returns a Machine in a cluster.
	GetMachine(machinename string, clustername string) (Machine, error)
	// DeleteMachine deletes a Machine in a cluster.
	DeleteMachine(machinename string, clustername string) error
	// NewMachine creates a new Machine in a cluster, usually using an Image
	// for the supplied Kubernetes version.
	NewMachine(machinename string, clustername string, k8sversion string) (Machine, error)

	// UpdateImageList fetches the latest list of Images from a driver-defined
	// source, and stores it locally.
	UpdateImageList() error
	// ValidK8sVersion returns true if the specified Kubernetes version is
	// available in the local list.
	ValidK8sVersion(k8sversion string) bool
	// K8sVersions returns all Kubernetes versions currently supported by a
	// Driver. This should come from the local list.
	K8sVersions() []string
	// ListImages returns the currently available Images from the local list.
	ListImages() ([]Image, error)
	// GetImage returns the Image for a Kubernetes version, or an error.
	GetImage(k8sversion string) (Image, error)
}
