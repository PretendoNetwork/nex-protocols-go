// Package protocol implements the StorageManager protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the StorageManager protocol
	ProtocolID = 0x6E

	// MethodAcquireCardID is the method ID for the method AcquireCardID
	MethodAcquireCardID = 0x4

	// MethodActivateWithCardID is the method ID for the method ActivateWithCardID
	MethodActivateWithCardID = 0x5
)

// Protocol stores all the RMC method handlers for the StorageManager protocol and listens for requests
type Protocol struct {
	server             nex.ServerInterface
	AcquireCardID      func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	ActivateWithCardID func(err error, packet nex.PacketInterface, callID uint32, unknown *types.PrimitiveU8, cardID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)
}

// Interface implements the methods present on the StorageManager protocol struct
type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerAcquireCardID(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerActivateWithCardID(handler func(err error, packet nex.PacketInterface, callID uint32, unknown *types.PrimitiveU8, cardID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerAcquireCardID sets the handler for the AcquireCardID method
func (protocol *Protocol) SetHandlerAcquireCardID(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.AcquireCardID = handler
}

// SetHandlerActivateWithCardID sets the handler for the ActivateWithCardID method
func (protocol *Protocol) SetHandlerActivateWithCardID(handler func(err error, packet nex.PacketInterface, callID uint32, unknown *types.PrimitiveU8, cardID *types.PrimitiveU64) (*nex.RMCMessage, *nex.Error)) {
	protocol.ActivateWithCardID = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	switch message.MethodID {
	case MethodAcquireCardID:
		protocol.handleAcquireCardID(packet)
	case MethodActivateWithCardID:
		protocol.handleActivateWithCardID(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported StorageManager method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new StorageManager protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	return &Protocol{server: server}
}
