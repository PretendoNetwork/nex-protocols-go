// Package messaging implements the Messaging protocol
package messaging

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

// MessagingProtocol handles the Messaging NEX protocol
type MessagingProtocol struct {
	Server                                *nex.Server
	deliverMessageHandler                 func(err error, client *nex.Client, callID uint32, oUserMessage *nex.DataHolder)
	getNumberOfMessagesHandler            func(err error, client *nex.Client, callID uint32, recipient *messaging_types.MessageRecipient)
	getMessagesHeadersHandler             func(err error, client *nex.Client, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *nex.ResultRange)
	retrieveAllMessagesWithinRangeHandler func(err error, client *nex.Client, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *nex.ResultRange)
	retrieveMessagesHandler               func(err error, client *nex.Client, callID uint32, recipient *messaging_types.MessageRecipient, lstMsgIDs []uint32, bLeaveOnServer bool)
	deleteMessagesHandler                 func(err error, client *nex.Client, callID uint32, recipient *messaging_types.MessageRecipient, lstMessagesToDelete []uint32)
	deleteAllMessagesHandler              func(err error, client *nex.Client, callID uint32, recipient *messaging_types.MessageRecipient)
	deliverMessageMultiTargetHandler      func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (protocol *MessagingProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *MessagingProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodDeliverMessage:
		go protocol.handleDeliverMessage(packet)
	case MethodGetNumberOfMessages:
		go protocol.handleGetNumberOfMessages(packet)
	case MethodGetMessagesHeaders:
		go protocol.handleGetMessagesHeaders(packet)
	case MethodRetrieveAllMessagesWithinRange:
		go protocol.handleRetrieveAllMessagesWithinRange(packet)
	case MethodRetrieveMessages:
		go protocol.handleRetrieveMessages(packet)
	case MethodDeleteMessages:
		go protocol.handleDeleteMessages(packet)
	case MethodDeleteAllMessages:
		go protocol.handleDeleteAllMessages(packet)
	case MethodDeliverMessageMultiTarget:
		go protocol.handleDeliverMessageMultiTarget(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported MessageDelivery method ID: %#v\n", request.MethodID())
	}
}

// NewMessagingProtocol returns a new MessagingProtocol
func NewMessagingProtocol(server *nex.Server) *MessagingProtocol {
	protocol := &MessagingProtocol{Server: server}

	protocol.Setup()

	return protocol
}
