package drivermock

// New creates a new mock driver.
func New(drivername string, description string, usesNAT bool, usesNetworking bool) *Driver {
	return &Driver{
		driverName:               drivername,
		driverDescription:        description,
		usesPerClusterNetworking: usesNetworking,
		usesNATNetworking:        usesNAT,
		networkNameSuffix:        drivername + "net",
		status:                   "Ready",
	}
}
