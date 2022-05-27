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
	Name() string
	Description() string
	UsesPerClusterNetworking() bool
	UsesNATNetworking() bool
	Status() string
	Error() string

	QualifiedNetworkName(clustername string) string
	ListNetworks() ([]Network, error)
	GetNetwork(clustername string) (Network, error)
	DeleteNetwork(clustername string) error
	NewNetwork(clustername string) (Network, error)

	QualifiedMachineName(machinename string, clustername string) string
	ListMachines() ([]Machine, error)
	GetMachine(machinename string, clustername string) (Machine, error)
	DeleteMachine(machinename string, clustername string) error
	NewMachine(machinename string, clustername string, k8sversion string) (Machine, error)

	UpdateImageList() error
	ValidK8sVersion(k8sversion string) bool
	K8sVersions() []string
	ListImages() ([]Image, error)
	GetImage(k8sversion string) (Image, error)
}
