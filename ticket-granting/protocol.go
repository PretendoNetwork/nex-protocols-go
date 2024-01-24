// Package protocol implements the Ticket Granting protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
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
	server           nex.ServerInterface
	Login            func(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String) (*nex.RMCMessage, uint32)
	LoginEx          func(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String, oExtraData *types.AnyDataHolder) (*nex.RMCMessage, uint32)
	RequestTicket    func(err error, packet nex.PacketInterface, callID uint32, idSource *types.PID, idTarget *types.PID) (*nex.RMCMessage, uint32)
	GetPID           func(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String) (*nex.RMCMessage, uint32)
	GetName          func(err error, packet nex.PacketInterface, callID uint32, id *types.PID) (*nex.RMCMessage, uint32)
	LoginWithContext func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
}

// Interface implements the methods present on the Ticket Granting protocol struct
type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerLogin(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String) (*nex.RMCMessage, uint32))
	SetHandlerLoginEx(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String, oExtraData *types.AnyDataHolder) (*nex.RMCMessage, uint32))
	SetHandlerRequestTicket(handler func(err error, packet nex.PacketInterface, callID uint32, idSource *types.PID, idTarget *types.PID) (*nex.RMCMessage, uint32))
	SetHandlerGetPID(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String) (*nex.RMCMessage, uint32))
	SetHandlerGetName(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PID) (*nex.RMCMessage, uint32))
	SetHandlerLoginWithContext(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerLogin sets the handler for the Login method
func (protocol *Protocol) SetHandlerLogin(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String) (*nex.RMCMessage, uint32)) {
	protocol.Login = handler
}

// SetHandlerLoginEx sets the handler for the LoginEx method
func (protocol *Protocol) SetHandlerLoginEx(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String, oExtraData *types.AnyDataHolder) (*nex.RMCMessage, uint32)) {
	protocol.LoginEx = handler
}

// SetHandlerRequestTicket sets the handler for the RequestTicket method
func (protocol *Protocol) SetHandlerRequestTicket(handler func(err error, packet nex.PacketInterface, callID uint32, idSource *types.PID, idTarget *types.PID) (*nex.RMCMessage, uint32)) {
	protocol.RequestTicket = handler
}

// SetHandlerGetPID sets the handler for the GetPID method
func (protocol *Protocol) SetHandlerGetPID(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName *types.String) (*nex.RMCMessage, uint32)) {
	protocol.GetPID = handler
}

// SetHandlerGetName sets the handler for the GetName method
func (protocol *Protocol) SetHandlerGetName(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PID) (*nex.RMCMessage, uint32)) {
	protocol.GetName = handler
}

// SetHandlerLoginWithContext sets the handler for the LoginWithContext method
func (protocol *Protocol) SetHandlerLoginWithContext(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.LoginWithContext = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	switch message.MethodID {
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
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		fmt.Printf("Unsupported Ticket Granting method ID: %#v\n", message.MethodID)
	}
}

// NewProtocol returns a new Ticket Granting protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	return &Protocol{server: server}
}
