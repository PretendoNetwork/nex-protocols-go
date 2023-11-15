// Package protocol implements the Ticket Granting protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Ticket Granting protocol
	ProtocolID = 0xA

	// MethodLogin is the method ID for the method Login
	MethodLogin = 0x1

	// MethodLoginEx is the method ID for the method LoginEx
	MethodLoginEx = 0x2

	// MethodRequestTicket is the method ID for the method RequestTicket
	MethodRequestTicket = 0x3

	// MethodGetPID is the method ID for the method GetPID
	MethodGetPID = 0x4

	// MethodGetName is the method ID for the method GetName
	MethodGetName = 0x5

	// MethodLoginWithContext is the method ID for the method LoginWithContext
	MethodLoginWithContext = 0x6
)

// Protocol stores all the RMC method handlers for the Ticket Granting protocol and listens for requests
type Protocol struct {
	Server           nex.ServerInterface
	Login            func(err error, packet nex.PacketInterface, callID uint32, strUserName string) (*nex.RMCMessage, uint32)
	LoginEx          func(err error, packet nex.PacketInterface, callID uint32, strUserName string, oExtraData *nex.DataHolder) (*nex.RMCMessage, uint32)
	RequestTicket    func(err error, packet nex.PacketInterface, callID uint32, idSource *nex.PID, idTarget *nex.PID) (*nex.RMCMessage, uint32)
	GetPID           func(err error, packet nex.PacketInterface, callID uint32, strUserName string) (*nex.RMCMessage, uint32)
	GetName          func(err error, packet nex.PacketInterface, callID uint32, id *nex.PID) (*nex.RMCMessage, uint32)
	LoginWithContext func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		if request.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodLogin:
		protocol.handleLogin(packet)
	case MethodLoginEx:
		protocol.handleLoginEx(packet)
	case MethodRequestTicket:
		protocol.handleRequestTicket(packet)
	case MethodGetPID:
		protocol.handleGetPID(packet)
	case MethodGetName:
		protocol.handleGetName(packet)
	case MethodLoginWithContext:
		protocol.handleLoginWithContext(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Ticket Granting method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Ticket Granting protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
