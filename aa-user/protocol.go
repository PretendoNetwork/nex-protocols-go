// Package protocol implements the AAUser protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	aauser_types "github.com/PretendoNetwork/nex-protocols-go/aa-user/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the AAUser protocol
	ProtocolID = 0x7B

	// MethodRegisterApplication is the method ID for the method RegisterApplication
	MethodRegisterApplication = 0x1

	// MethodUnregisterApplication is the method ID for the method UnregisterApplication
	MethodUnregisterApplication = 0x2

	// MethodSetApplicationInfo is the method ID for the method RegisterApplication
	MethodSetApplicationInfo = 0x3

	// MethodGetApplicationInfo is the method ID for the method GetApplicationInfo
	MethodGetApplicationInfo = 0x4
)

// Protocol stores all the RMC method handlers for the AAUser protocol and listens for requests
type Protocol struct {
	server                nex.ServerInterface
	RegisterApplication   func(err error, packet nex.PacketInterface, callID uint32, titleID uint64) (*nex.RMCMessage, uint32)
	UnregisterApplication func(err error, packet nex.PacketInterface, callID uint32, titleID uint64) (*nex.RMCMessage, uint32)
	SetApplicationInfo    func(err error, packet nex.PacketInterface, callID uint32, applicationInfo []*aauser_types.ApplicationInfo) (*nex.RMCMessage, uint32)
	GetApplicationInfo    func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
}

// Interface implements the methods present on the AAUser Protocol struct
type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerRegisterApplication(handler func(err error, packet nex.PacketInterface, callID uint32, titleID uint64) (*nex.RMCMessage, uint32))
	SetHandlerUnregisterApplication(handler func(err error, packet nex.PacketInterface, callID uint32, titleID uint64) (*nex.RMCMessage, uint32))
	SetHandlerSetApplicationInfo(handler func(err error, packet nex.PacketInterface, callID uint32, applicationInfo []*aauser_types.ApplicationInfo) (*nex.RMCMessage, uint32))
	SetHandlerGetApplicationInfo(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerRegisterApplication sets the handler for the RegisterApplication method
func (protocol *Protocol) SetHandlerRegisterApplication(handler func(err error, packet nex.PacketInterface, callID uint32, titleID uint64) (*nex.RMCMessage, uint32)) {
	protocol.RegisterApplication = handler
}

// SetHandlerUnregisterApplication sets the handler for the UnregisterApplication method
func (protocol *Protocol) SetHandlerUnregisterApplication(handler func(err error, packet nex.PacketInterface, callID uint32, titleID uint64) (*nex.RMCMessage, uint32)) {
	protocol.UnregisterApplication = handler
}

// SetHandlerSetApplicationInfo sets the handler for the SetApplicationInfo method
func (protocol *Protocol) SetHandlerSetApplicationInfo(handler func(err error, packet nex.PacketInterface, callID uint32, applicationInfo []*aauser_types.ApplicationInfo) (*nex.RMCMessage, uint32)) {
	protocol.SetApplicationInfo = handler
}

// SetHandlerGetApplicationInfo sets the handler for the GetApplicationInfo method
func (protocol *Protocol) SetHandlerGetApplicationInfo(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.GetApplicationInfo = handler
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			switch message.MethodID {
			case MethodRegisterApplication:
				protocol.handleRegisterApplication(packet)
			case MethodUnregisterApplication:
				protocol.handleUnregisterApplication(packet)
			case MethodSetApplicationInfo:
				protocol.handleSetApplicationInfo(packet)
			case MethodGetApplicationInfo:
				protocol.handleGetApplicationInfo(packet)
			default:
				globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported AAUser method ID: %#v\n", message.MethodID)
			}
		}
	})
}

// NewProtocol returns a new AAUser protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{server: server}

	protocol.Setup()

	return protocol
}
