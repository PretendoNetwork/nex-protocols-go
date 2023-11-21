// Package protocol implements the Messaging protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/messaging/types"
)

const (
	// ProtocolID is the protocol ID for the Messaging protocol
	ProtocolID = 0x17

	// MethodDeliverMessage is the method ID for method DeliverMessage
	MethodDeliverMessage = 0x1

	// MethodGetNumberOfMessages is the method ID for method GetNumberOfMessages
	MethodGetNumberOfMessages = 0x2

	// MethodGetMessagesHeaders is the method ID for method GetMessagesHeaders
	MethodGetMessagesHeaders = 0x3

	// MethodRetrieveAllMessagesWithinRange is the method ID for method RetrieveAllMessagesWithinRange
	MethodRetrieveAllMessagesWithinRange = 0x4

	// MethodRetrieveMessages is the method ID for method RetrieveMessages
	MethodRetrieveMessages = 0x5

	// MethodDeleteMessages is the method ID for method DeleteMessages
	MethodDeleteMessages = 0x6

	// MethodDeleteAllMessages is the method ID for method DeleteAllMessages
	MethodDeleteAllMessages = 0x7

	// MethodDeliverMessageMultiTarget is the method ID for method DeliverMessageMultiTarget
	MethodDeliverMessageMultiTarget = 0x8
)

// Protocol stores all the RMC method handlers for the Messaging protocol and listens for requests
type Protocol struct {
	Server                         nex.ServerInterface
	DeliverMessage                 func(err error, packet nex.PacketInterface, callID uint32, oUserMessage *nex.DataHolder) (*nex.RMCMessage, uint32)
	GetNumberOfMessages            func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient) (*nex.RMCMessage, uint32)
	GetMessagesHeaders             func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)
	RetrieveAllMessagesWithinRange func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)
	RetrieveMessages               func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, lstMsgIDs []uint32, bLeaveOnServer bool) (*nex.RMCMessage, uint32)
	DeleteMessages                 func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, lstMessagesToDelete []uint32) (*nex.RMCMessage, uint32)
	DeleteAllMessages              func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient) (*nex.RMCMessage, uint32)
	DeliverMessageMultiTarget      func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32) // TODO - Unknown request/response format
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodDeliverMessage:
		protocol.handleDeliverMessage(packet)
	case MethodGetNumberOfMessages:
		protocol.handleGetNumberOfMessages(packet)
	case MethodGetMessagesHeaders:
		protocol.handleGetMessagesHeaders(packet)
	case MethodRetrieveAllMessagesWithinRange:
		protocol.handleRetrieveAllMessagesWithinRange(packet)
	case MethodRetrieveMessages:
		protocol.handleRetrieveMessages(packet)
	case MethodDeleteMessages:
		protocol.handleDeleteMessages(packet)
	case MethodDeleteAllMessages:
		protocol.handleDeleteAllMessages(packet)
	case MethodDeliverMessageMultiTarget:
		protocol.handleDeliverMessageMultiTarget(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Messaging method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Messaging protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
