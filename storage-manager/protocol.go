// Package protocol implements the StorageManager protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the StorageManager protocol
	ProtocolID = 0x6E

	// MethodAcquireNexUniqueID is the method ID for the method AcquireNexUniqueID
	MethodAcquireNexUniqueID = 0x1

	// MethodNexUniqueIDToPrincipalID is the method ID for the method NexUniqueIDToPrincipalID
	MethodNexUniqueIDToPrincipalID = 0x2

	// MethodUnk3 is the method ID for the method Unk3
	// TODO - Find name if possible
	MethodUnk3 = 0x3

	// MethodAcquireCardID is the method ID for the method AcquireCardID
	MethodAcquireCardID = 0x4

	// MethodActivateWithCardID is the method ID for the method ActivateWithCardID
	MethodActivateWithCardID = 0x5
)

// Protocol stores all the RMC method handlers for the StorageManager protocol and listens for requests
type Protocol struct {
	endpoint                 nex.EndpointInterface
	AcquireNexUniqueID       func(err error, packet nex.PacketInterface, callID uint32, slot types.UInt8) (*nex.RMCMessage, *nex.Error)
	NexUniqueIDToPrincipalID func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32) (*nex.RMCMessage, *nex.Error)
	Unk3                     func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error) // TODO - Find name if possible
	AcquireCardID            func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	ActivateWithCardID       func(err error, packet nex.PacketInterface, callID uint32, slot types.UInt8, cardID types.UInt64) (*nex.RMCMessage, *nex.Error)
	Patches                  nex.ServiceProtocol
	PatchedMethods           []uint32
}

// Interface implements the methods present on the StorageManager protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerAcquireNexUniqueID(handler func(err error, packet nex.PacketInterface, callID uint32, slot types.UInt8) (*nex.RMCMessage, *nex.Error))
	SetHandlerNexUniqueIDToPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUnk3(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) // TODO - Find name if possible
	SetHandlerAcquireCardID(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerActivateWithCardID(handler func(err error, packet nex.PacketInterface, callID uint32, slot types.UInt8, cardID types.UInt64) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerAcquireNexUniqueID sets the handler for the AcquireNexUniqueID method
func (protocol *Protocol) SetHandlerAcquireNexUniqueID(handler func(err error, packet nex.PacketInterface, callID uint32, slot types.UInt8) (*nex.RMCMessage, *nex.Error)) {
	protocol.AcquireNexUniqueID = handler
}

// SetHandlerNexUniqueIDToPrincipalID sets the handler for the NexUniqueIDToPrincipalID method
func (protocol *Protocol) SetHandlerNexUniqueIDToPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.NexUniqueIDToPrincipalID = handler
}

// SetHandlerUnk3 sets the handler for the Unk3 method
// TODO - Find name if possible
func (protocol *Protocol) SetHandlerUnk3(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.Unk3 = handler
}

// SetHandlerAcquireCardID sets the handler for the AcquireCardID method
func (protocol *Protocol) SetHandlerAcquireCardID(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.AcquireCardID = handler
}

// SetHandlerActivateWithCardID sets the handler for the ActivateWithCardID method
func (protocol *Protocol) SetHandlerActivateWithCardID(handler func(err error, packet nex.PacketInterface, callID uint32, unknown types.UInt8, cardID types.UInt64) (*nex.RMCMessage, *nex.Error)) {
	protocol.ActivateWithCardID = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if protocol.Patches != nil && slices.Contains(protocol.PatchedMethods, message.MethodID) {
		protocol.Patches.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodAcquireNexUniqueID:
		protocol.handleAcquireNexUniqueID(packet)
	case MethodNexUniqueIDToPrincipalID:
		protocol.handleNexUniqueIDToPrincipalID(packet)
	case MethodUnk3:
		protocol.handleUnk3(packet) // TODO - Find name if possible
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
func NewProtocol() *Protocol {
	return &Protocol{}
}
