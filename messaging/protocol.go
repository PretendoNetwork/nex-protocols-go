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
	server                         nex.ServerInterface
	DeliverMessage                 func(err error, packet nex.PacketInterface, callID uint32, oUserMessage *nex.DataHolder) (*nex.RMCMessage, uint32)
	GetNumberOfMessages            func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient) (*nex.RMCMessage, uint32)
	GetMessagesHeaders             func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)
	RetrieveAllMessagesWithinRange func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)
	RetrieveMessages               func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, lstMsgIDs []uint32, bLeaveOnServer bool) (*nex.RMCMessage, uint32)
	DeleteMessages                 func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, lstMessagesToDelete []uint32) (*nex.RMCMessage, uint32)
	DeleteAllMessages              func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient) (*nex.RMCMessage, uint32)
	DeliverMessageMultiTarget      func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32) // TODO - Unknown request/response format
}

// Interface implements the methods present on the Messaging protocol struct
type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerDeliverMessage(handler func(err error, packet nex.PacketInterface, callID uint32, oUserMessage *nex.DataHolder) (*nex.RMCMessage, uint32))
	SetHandlerGetNumberOfMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient) (*nex.RMCMessage, uint32))
	SetHandlerGetMessagesHeaders(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32))
	SetHandlerRetrieveAllMessagesWithinRange(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32))
	SetHandlerRetrieveMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, lstMsgIDs []uint32, bLeaveOnServer bool) (*nex.RMCMessage, uint32))
	SetHandlerDeleteMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, lstMessagesToDelete []uint32) (*nex.RMCMessage, uint32))
	SetHandlerDeleteAllMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient) (*nex.RMCMessage, uint32))
	SetHandlerDeliverMessageMultiTarget(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerDeliverMessage sets the handler for the DeliverMessage method
func (protocol *Protocol) SetHandlerDeliverMessage(handler func(err error, packet nex.PacketInterface, callID uint32, oUserMessage *nex.DataHolder) (*nex.RMCMessage, uint32)) {
	protocol.DeliverMessage = handler
}

// SetHandlerGetNumberOfMessages sets the handler for the GetNumberOfMessages method
func (protocol *Protocol) SetHandlerGetNumberOfMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient) (*nex.RMCMessage, uint32)) {
	protocol.GetNumberOfMessages = handler
}

// SetHandlerGetMessagesHeaders sets the handler for the GetMessagesHeaders method
func (protocol *Protocol) SetHandlerGetMessagesHeaders(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)) {
	protocol.GetMessagesHeaders = handler
}

// SetHandlerRetrieveAllMessagesWithinRange sets the handler for the RetrieveAllMessagesWithinRange method
func (protocol *Protocol) SetHandlerRetrieveAllMessagesWithinRange(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)) {
	protocol.RetrieveAllMessagesWithinRange = handler
}

// SetHandlerRetrieveMessages sets the handler for the RetrieveMessages method
func (protocol *Protocol) SetHandlerRetrieveMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, lstMsgIDs []uint32, bLeaveOnServer bool) (*nex.RMCMessage, uint32)) {
	protocol.RetrieveMessages = handler
}

// SetHandlerDeleteMessages sets the handler for the DeleteMessages method
func (protocol *Protocol) SetHandlerDeleteMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, lstMessagesToDelete []uint32) (*nex.RMCMessage, uint32)) {
	protocol.DeleteMessages = handler
}

// SetHandlerDeleteAllMessages sets the handler for the DeleteAllMessages method
func (protocol *Protocol) SetHandlerDeleteAllMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient) (*nex.RMCMessage, uint32)) {
	protocol.DeleteAllMessages = handler
}

// SetHandlerDeliverMessageMultiTarget sets the handler for the DeliverMessageMultiTarget method
func (protocol *Protocol) SetHandlerDeliverMessageMultiTarget(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)) {
	protocol.DeliverMessageMultiTarget = handler
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.server.OnData(func(packet nex.PacketInterface) {
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
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Messaging method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Messaging protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{server: server}

	protocol.Setup()

	return protocol
}
