package drivermock

import (
	"errors"
	"fmt"
	"time"

	"github.com/kuttiproject/drivercore"
)

// Machine implements the drivercore.Machine interface
type Machine struct {
	machinename  string
	status       drivercore.MachineStatus
	errormessage string
	ports        map[int]int
}

// Name returns the name of the machine.
func (vh *Machine) Name() string {
	return vh.machinename
}

// Status returns the status of the machine.
func (vh *Machine) Status() drivercore.MachineStatus {
	return vh.status
}

// Error returns the last error message returned by
// the machine, or an empty string.
func (vh *Machine) Error() string {
	return vh.errormessage
}

// IPAddress returns the IPv4 address of the machine.
func (vh *Machine) IPAddress() string {
	panic("not implemented") // TODO: Implement
}

// SSHAddress returns the SSH address of the machine.
func (vh *Machine) SSHAddress() string {
	vh.ensureports()
	hostport, ok := vh.ports[22]
	if !ok {
		return ""
	}

	return fmt.Sprintf("localhost:%v", hostport)
}

// Start starts the machine.
func (vh *Machine) Start() error {
	vh.status = drivercore.MachineStatusRunning
	return nil
}

// Stop stops the machine.
func (vh *Machine) Stop() error {
	vh.status = drivercore.MachineStatusStopped
	return nil
}

// ForceStop forcibly stops the machine.
func (vh *Machine) ForceStop() error {
	return vh.Stop()
}

// WaitForStateChange waits until the status of a machine has
// changed, or until timeout. In this mock, it just waits till
// 1/10th of timeout.
func (vh *Machine) WaitForStateChange(timeoutinseconds int) {
	time.Sleep(time.Second * time.Duration(timeoutinseconds/10))
}

// ForwardPort forwards a port.
func (vh *Machine) ForwardPort(hostport int, machineport int) error {
	vh.ensureports()

	for mp, hp := range vh.ports {
		if mp == machineport {
			return errors.New("port already forwarded")
		}
		if hp == hostport {
			return errors.New("host port is occupied")
		}
	}

	vh.ports[machineport] = hostport
	return nil
}

// UnforwardPort unforwards a port.
func (vh *Machine) UnforwardPort(machineport int) error {
	vh.ensureports()
	_, ok := vh.ports[machineport]
	if !ok {
		return errors.New("port not forwarded")
	}
	delete(vh.ports, machineport)
	return nil
}

// ForwardSSHPort forwards the SSH port.
func (vh *Machine) ForwardSSHPort(hostport int) error {
	return vh.ForwardPort(hostport, 22)
}

// ImplementsCommand returns true if a predefined command has been
// implemented, false otherwise.
func (vh *Machine) ImplementsCommand(command drivercore.PredefinedCommand) bool {
	panic("not implemented") // TODO: Implement
}

// ExecuteCommand executes a predefined command.
func (vh *Machine) ExecuteCommand(command drivercore.PredefinedCommand, params ...string) error {
	panic("not implemented") // TODO: Implement
}

func (vh *Machine) ensureports() {
	if vh.ports == nil {
		vh.ports = map[int]int{}
	}
}
