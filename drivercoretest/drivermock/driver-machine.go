package drivermock

import (
	"errors"

	"github.com/kuttiproject/drivercore"
)

var mockmachines = map[string]*Machine{}

// QualifiedMachineName returns the driver-qualified machine name.
func (vd *Driver) QualifiedMachineName(machinename string, clustername string) string {
	return clustername + "-" + machinename
}

// ListMachines returns the list of machines.
func (vd *Driver) ListMachines() ([]drivercore.Machine, error) {
	result := make([]drivercore.Machine, len(mockmachines))
	index := 0
	for _, value := range mockmachines {
		result[index] = value
		index++
	}

	return result, nil
}

// GetMachine returns the specified machine, or an error.
func (vd *Driver) GetMachine(machinename string, clustername string) (drivercore.Machine, error) {
	result, ok := mockmachines[vd.QualifiedMachineName(machinename, clustername)]
	if !ok {
		return nil, errors.New("machine not found")
	}
	return result, nil
}

// DeleteMachine deletes the specified machine.
func (vd *Driver) DeleteMachine(machinename string, clustername string) error {
	mname := vd.QualifiedMachineName(machinename, clustername)
	if _, ok := mockmachines[mname]; !ok {
		return errors.New("machine not found")
	}
	delete(mockmachines, mname)
	return nil
}

// NewMachine creates a new machine.
func (vd *Driver) NewMachine(machinename string, clustername string, k8sversion string) (drivercore.Machine, error) {
	qname := vd.QualifiedMachineName(machinename, clustername)
	_, ok := mockmachines[qname]
	if ok {
		return nil, errors.New("machine already exists")
	}
	newmachine := &Machine{
		machinename: machinename,
		status:      drivercore.MachineStatusStopped,
	}
	mockmachines[qname] = newmachine
	return newmachine, nil
}
