package model

/*
ConnectionTypeEnum represents a connection type
*/
type ConnectionTypeEnum int

const (
	// SSHConnType represents an SSH connection
	SSHConnType = iota
	// LOCALConType represents a local connection
	LOCALConType
)

/*
MCServer represents an MC server
*/
type MCServer struct {
	Name              string
	Host              string
	User              string
	PrivateKeyPath    string
	PrivateKeyContent string
	ConnectionType    ConnectionTypeEnum
	ConnectionPort    string
	Playbook          string
	ExtraVars         map[string]string
}
