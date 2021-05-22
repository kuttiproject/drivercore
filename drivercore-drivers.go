package drivercore

import "strings"

var drivers = map[string]Driver{}

// RegisteredDriver checks if a driver name is registered. The name
// check is case-insensitive.
func RegisteredDriver(name string) bool {
	_, ok := drivers[strings.ToLower(name)]
	return ok
}

// RegisteredDrivers returns all registered driver names.
func RegisteredDrivers() []string {
	result := make([]string, 0, len(drivers))
	for i := range drivers {
		result = append(result, i)
	}
	return result
}

// DriverCount returns the number of registered drivers.
func DriverCount() int {
	return len(drivers)
}

// ForEachDriver iterates over registered Drivers.
// The callback function can return false to stop the iteration.
func ForEachDriver(f func(Driver) bool) {
	for _, driver := range drivers {
		cancel := f(driver)
		if cancel {
			break
		}
	}
}

// GetDriver returns a Driver corresponding to the name.
// If there is no driver registered against the name, nil is returned.
// Names are case-insensitive.
func GetDriver(name string) (Driver, bool) {
	result, ok := drivers[strings.ToLower(name)]
	return result, ok
}

// RegisterDriver registers a Driver with a name to drivercore.
// If a driver with the specified name already exists, it is replaced.
// Driver names are converted to lower case.
func RegisterDriver(name string, d Driver) {
	if d != nil {
		drivers[strings.ToLower(name)] = d
	}
}
