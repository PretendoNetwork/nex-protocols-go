// Package protocol implements the UbiNews protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the UbiNews protocol
	ProtocolID = 0x21

	// MethodGetChannel is the method ID for the method GetChannel
	MethodGetChannel = 0x1

	// MethodGetNewsHeaders is the method ID for the method GetNewsHeaders
	MethodGetNewsHeaders = 0x2

	// MethodGetNewsMessages is the method ID for the method GetNewsMessages
	MethodGetNewsMessages = 0x3
)

// Protocol stores all the RMC method handlers for the UbiNews protocol and listens for requests.
type Protocol struct {
	endpoint        nex.EndpointInterface
	GetChannel      func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	GetNewsHeaders  func(err error, packet nex.PacketInterface, callID uint32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error)
	GetNewsMessages func(err error, packet nex.PacketInterface, callID uint32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error)
	Patches         nex.ServiceProtocol
	PatchedMethods  []uint32
}

// Interface implements the methods present on the UbiNews protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerGetChannel(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetNewsHeaders(handler func(err error, packet nex.PacketInterface, callID uint32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetNewsMessages(handler func(err error, packet nex.PacketInterface, callID uint32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerGetChannel sets the handler for the GetChannel method
func (protocol *Protocol) SetHandlerGetChannel(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetChannel = handler
}

// SetHandlerGetNewsHeaders sets the handler for the GetNewsHeaders method
func (protocol *Protocol) SetHandlerGetNewsHeaders(handler func(err error, packet nex.PacketInterface, callID uint32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetNewsHeaders = handler
}

// SetHandlerGetNewsMessages set the handler for the GetNewsMessages method
func (protocol *Protocol) SetHandlerGetNewsMessages(handler func(err error, packet nex.PacketInterface, callID uint32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetNewsMessages = handler
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
	case MethodGetChannel:
		protocol.handleGetChannel(packet)
	case MethodGetNewsHeaders:
		protocol.handleGetNewsHeaders(packet)
	case MethodGetNewsMessages:
		protocol.handleGetNewsMessages(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported UbiNews method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new UbiNews protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
