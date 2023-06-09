package message_delivery

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

// MessageDeliveryProtocol handles the Authentication nex protocol
type MessageDeliveryProtocol struct {
	Server                *nex.Server
	DeliverMessageHandler func(err error, client *nex.Client, callID uint32, oUserMessage *nex.DataHolder)
}

// Setup initializes the protocol
func (protocol *MessageDeliveryProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

func (protocol *MessageDeliveryProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodDeliverMessage:
		go protocol.HandleDeliverMessage(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported MessageDelivery method ID: %#v\n", request.MethodID())
	}
}

// NewMessageDeliveryProtocol returns a new MessageDeliveryProtocol
func NewMessageDeliveryProtocol(server *nex.Server) *MessageDeliveryProtocol {
	protocol := &MessageDeliveryProtocol{Server: server}

	protocol.Setup()

	return protocol
}
