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
// A machine can be started, stoppped normally and stopped forcibly.
// If the Driver and the Network use NAT, then a Machine can also have its
// ports forwarded to NAT router ports.
// A machine provides a way to execute predefined commands within its guest
// operating system.
type Machine interface {
	Name() string
	Status() MachineStatus
	Error() string

	IPAddress() string
	SSHAddress() string

	Start() error
	Stop() error
	ForceStop() error
	WaitForStateChange(timeoutinseconds int)

	ForwardPort(hostport int, machineport int) error
	UnforwardPort(machineport int) error
	ForwardSSHPort(hostport int) error

	ImplementsCommand(command PredefinedCommand) bool
	ExecuteCommand(command PredefinedCommand, params ...string) error
}
