package drivermock

// New creates a new mock driver.
func New(drivername string, description string, usesNAT bool) *Driver {
	return &Driver{
		driverName:        drivername,
		driverDescription: description,
		usesNATNetworking: usesNAT,
		networkNameSuffix: drivername + "net",
		status:            "Ready",
	}
}
