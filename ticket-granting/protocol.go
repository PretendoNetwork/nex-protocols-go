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
	Server                  *nex.Server
	loginHandler            func(err error, packet nex.PacketInterface, callID uint32, strUserName string) uint32
	loginExHandler          func(err error, packet nex.PacketInterface, callID uint32, strUserName string, oExtraData *nex.DataHolder) uint32
	requestTicketHandler    func(err error, packet nex.PacketInterface, callID uint32, idSource uint32, idTarget uint32) uint32
	getPIDHandler           func(err error, packet nex.PacketInterface, callID uint32, strUserName string) uint32
	getNameHandler          func(err error, packet nex.PacketInterface, callID uint32, id uint32) uint32
	loginWithContextHandler func(err error, packet nex.PacketInterface, callID uint32) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodLogin:
		go protocol.handleLogin(packet)
	case MethodLoginEx:
		go protocol.handleLoginEx(packet)
	case MethodRequestTicket:
		go protocol.handleRequestTicket(packet)
	case MethodGetPID:
		go protocol.handleGetPID(packet)
	case MethodGetName:
		go protocol.handleGetName(packet)
	case MethodLoginWithContext:
		go protocol.handleLoginWithContext(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Ticket Granting method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new Ticket Granting protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
