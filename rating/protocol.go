// Package protocol implements the Rating protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Rating protocol
	ProtocolID = 0x76

	// MethodUnk1 is the method ID for the method Unk1
	// TODO: find name if possible
	MethodUnk1 = 0x1

	// MethodUnk2 is the method ID for the method Unk2
	// TODO: find name if possible
	MethodUnk2 = 0x2
)

// Protocol handles the Rating protocol
type Protocol struct {
	endpoint                         nex.EndpointInterface
	Unk1                			 func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error) // TODO: find name if possible
	Unk2                			 func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error) // TODO: find name if possible
	Patches                          nex.ServiceProtocol
	PatchedMethods                   []uint32
}

// Interface implements the methods present on the Rating protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerUnk1(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) // TODO: find name if possible
	SetHandlerUnk2(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) // TODO: find name if possible
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerUnk1 sets the handler for the Unk1 method
// TODO: find name if possible
func (protocol *Protocol) SetHandlerUnk1(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.Unk1 = handler
}

// SetHandlerUnk2 sets the handler for the Unk2 method
// TODO: find name if possible
func (protocol *Protocol) SetHandlerUnk2(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.Unk2 = handler
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
	case MethodUnk1:
		protocol.handleUnk1(packet) // TODO: find name if possible
	case MethodUnk2:
		protocol.handleUnk2(packet) // TODO: find name if possible
	default:
		errMessage := fmt.Sprintf("Unsupported Rating method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Rating protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
