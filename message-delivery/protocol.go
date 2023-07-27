// Package protocol implements the Message Deliver protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Message Delivery protocol
	ProtocolID = 0x1B

	// MethodDeliverMessage is the method ID for the method DeliverMessage
	MethodDeliverMessage = 0x1
)

// Protocol stores all the RMC method handlers for the Message Delivery protocol and listens for requests
type Protocol struct {
	Server                *nex.Server
	deliverMessageHandler func(err error, client *nex.Client, callID uint32, oUserMessage *nex.DataHolder)
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
	case MethodDeliverMessage:
		go protocol.handleDeliverMessage(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported MessageDelivery method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new Message Delivery protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
