package drivercore

// Driver describes operations to manage Networks, Machines and Images.
// The RequiresPortForwarding() method is particularly important. It is expected to
// return true if the Networks of the driver use NAT. This means that ports of the
// Machines will need to be forwarded to physical ports for access.
// A driver is also responsible for maintaining a local cache of Images.
// It can update a list of published images from a driver-defined source,
// download local copies of images, and delete local copies (but not the
// image itself).
type Driver interface {
	Name() string
	Description() string
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
	NewMachine(machinename string, clustername string, k8sversion string, imagetype string) (Machine, error)

	UpateImageList() error
	ValidImageType(imagetype string) bool
	ValidImageTypes() []string
	ValidVersion(k8sversion string) bool
	ValidVersions() []string
	ListImages() ([]Image, error)
	GetImage(k8sversion string, imagetype string) (Image, error)
}
