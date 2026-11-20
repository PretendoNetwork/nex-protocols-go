// Package protocol implements the IPToLocationAdmin protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the IPToLocationAdmin protocol
	ProtocolID = 0x2E

	// MethodGetLocationFromIP is the method ID for the method GetLocationFromIP
	MethodGetLocationFromIP = 0x1
)

// Protocol stores all the RMC method handlers for the IPToLocationAdmin protocol and listens for request
type Protocol struct {
	endpoint          nex.EndpointInterface
	GetLocationFromIP func(err error, packet nex.PacketInterface, callID uint32, ip types.String) (*nex.RMCMessage, *nex.Error)
	Patches           nex.ServiceProtocol
	PatchedMethods    []uint32
}

// Interface implements the methods present on IPToLOcationAdmin protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerGetLocationFromIP(handler func(err error, packet nex.PacketInterface, callID uint32, ip types.String) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerGetLocationFromIP sets the handler for the GetLocationFromIP
func (protocol *Protocol) SetHandlerGetLocationFromIP(handler func(err error, packet nex.PacketInterface, callID uint32, ip types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetLocationFromIP = handler
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
	case MethodGetLocationFromIP:
		protocol.handleGetLocationFromIP(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported IPToLocationAdmin method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new IPToLocationAdmin protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
