package drivermock

const (
	networkNameSuffix  = "kuttinet"
	networkNamePattern = "*" + networkNameSuffix
	dhcpaddress        = "192.168.125.3"
	dhcpnetmask        = "255.255.255.0"
	ipNetAddr          = "192.168.125"
	iphostbase         = 10
	forwardedPortBase  = 10000
)

// DefaultNetCIDR is the address range used by NAT networks.
var DefaultNetCIDR = "192.168.125.0/24"

// Driver implements the drivercore.Driver interface
type Driver struct {
	driverName               string
	driverDescription        string
	usesPerClusterNetworking bool
	usesNATNetworking        bool
	errormessage             string
	networkNameSuffix        string
	status                   string
}

// Name returns the driver name.
func (vd *Driver) Name() string {
	return vd.driverName
}

// Description returns the driver description.
func (vd *Driver) Description() string {
	return vd.driverDescription
}

// UsesPerClusterNetworking returns if this particular mock
// driver instance uses per cluster networking
func (vd *Driver) UsesPerClusterNetworking() bool {
	return vd.usesPerClusterNetworking
}

// UsesNATNetworking returns if this particular mock
// driver instance uses NAT networking
func (vd *Driver) UsesNATNetworking() bool {
	return vd.usesNATNetworking
}

// Status returns the driver status.
func (vd *Driver) Status() string {
	return vd.status
}

// Error returns the last error returned by the driver, or
// an empty string.
func (vd *Driver) Error() string {
	return vd.errormessage
}
