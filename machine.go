package drivercore

// MachineStatus can be either stopped or running
type MachineStatus string

// Machine Statuses
const (
	MachineStatusStopped = MachineStatus("Stopped")
	MachineStatusRunning = MachineStatus("Running")
	MachineStatusUnknown = MachineStatus("Unknown")
	MachineStatusError   = MachineStatus("Error")
)

// PredefinedCommand provides commands that Machines can execute while running.
// Drivers implement these.
type PredefinedCommand string

// Prefined Commands
const (
	RenameMachine         = PredefinedCommand("RenameMachine")
	RestartMachine        = PredefinedCommand("RestartMachine")
	CheckConnectivity     = PredefinedCommand("CheckConnectivity")
	SetProxy              = PredefinedCommand("SetProxy")
	SetNoProxy            = PredefinedCommand("SetNoProxy")
	InitCluster           = PredefinedCommand("InitCluster")
	JoinCluster           = PredefinedCommand("JoinCluster")
	InstallOverlayNetwork = PredefinedCommand("InstallOverlayNetwork")
)

// Machine describes a node in a cluster.
// A machine, when created from an Image, is expected to have an operating
// system and a container runtime already installed.
// A machine can be started, stopped normally and stopped forcibly.
// If the Driver and the Network use NAT, then a Machine can also have its
// ports forwarded to NAT router ports.
// A machine provides a way to execute predefined commands within its guest
// operating system.
type Machine interface {
	// Name is the name of the machine.
	// The operating system hostname should match this.
	Name() string
	// Status can be drivercore.MachineStatusRunning, drivercore.MachineStatusStopped
	// drivercore.MachineStatusUnknown or drivercore.MachineStatusError.
	Status() MachineStatus
	// Error returns the last error caused when manipulating this machine.
	// A valid value can be expected only when Status() returns
	// drivercore.MachineStatusError.
	Error() string

	// IPAddress returns the current IP Address of this Machine.
	// A valid value can be expected only when Status() returns
	// drivercore.MachineStatusRunning.
	IPAddress() string
	// SSHAddress returns the host address and port number to SSH into this Machine.
	// For drivers that use NAT netwoking, the host address will be 'localhost'.
	SSHAddress() string

	// Start starts a Machine.
	// Note that a Machine may not be ready for further operations at the end of this,
	// and therefore its status may not change immediately.
	// See WaitForStateChange().
	Start() error
	// Stop stops a Machine.
	// Note that a Machine may not be ready for further operations at the end of this,
	// and therefore its status will not change immediately.
	// See WaitForStateChange().
	Stop() error
	// ForceStop stops a Machine forcibly.
	// This operation should set the status to drivercore.MachineStatusStopped.
	ForceStop() error
	// WaitForStateChange waits the specified number of seconds, or until the Machine
	// status changes.
	// WaitForStateChange should be called after calls to Start() or Stop(), before
	// any other operation. It should not be called _before_ Stop().
	WaitForStateChange(timeoutinseconds int)

	// ForwardPort creates a rule to forward the specified Machine port to the
	// specified physical host port.
	ForwardPort(hostport int, machineport int) error
	// UnforwardPort removes the rule which forwarded the specified Machine port.
	UnforwardPort(machineport int) error
	// ForwardSSHPort forwards the SSH port of this Machine to the specified
	// physical host port.
	ForwardSSHPort(hostport int) error

	// ImplementsCommand returns true if the driver implements the specified
	// predefined operation.
	ImplementsCommand(command PredefinedCommand) bool
	// ExecuteCommand executes the specified predefined operation.
	ExecuteCommand(command PredefinedCommand, params ...string) error
}
