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
	Server                nex.ServerInterface
	RegisterApplication   func(err error, packet nex.PacketInterface, callID uint32, titleID uint64) (*nex.RMCMessage, uint32)
	UnregisterApplication func(err error, packet nex.PacketInterface, callID uint32, titleID uint64) (*nex.RMCMessage, uint32)
	SetApplicationInfo    func(err error, packet nex.PacketInterface, callID uint32, applicationInfo []*aauser_types.ApplicationInfo) (*nex.RMCMessage, uint32)
	GetApplicationInfo    func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
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
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
