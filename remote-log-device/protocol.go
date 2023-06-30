// Package remote_log_device implements the Remote Log Device NEX protocol
package remote_log_device

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// ProtocolID is the protocol ID for the RemoteLogDevice protocol
	ProtocolID = 0x1

	// MethodLog is the method ID for the method Log
	MethodLog = 0x1
)

// RemoteLogDeviceProtocol handles the RemoteLogDevice protocol
type RemoteLogDeviceProtocol struct {
	Server     *nex.Server
	LogHandler func(err error, client *nex.Client, callID uint32, strLine string)
}

// Setup initializes the protocol
func (protocol *RemoteLogDeviceProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *RemoteLogDeviceProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodLog:
		go protocol.handleLog(packet)
	default:
		fmt.Printf("Unsupported RemoteLogDevice method ID: %#v\n", request.MethodID())
	}
}

// NewRemoteLogDeviceProtocol returns a new RemoteLogDeviceProtocol
func NewRemoteLogDeviceProtocol(server *nex.Server) *RemoteLogDeviceProtocol {
	protocol := &RemoteLogDeviceProtocol{Server: server}

	protocol.Setup()

	return protocol
}
