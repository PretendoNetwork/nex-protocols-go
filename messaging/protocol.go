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
	Server                                *nex.Server
	deliverMessageHandler                 func(err error, client *nex.Client, callID uint32, oUserMessage *nex.DataHolder) uint32
	getNumberOfMessagesHandler            func(err error, client *nex.Client, callID uint32, recipient *messaging_types.MessageRecipient) uint32
	getMessagesHeadersHandler             func(err error, client *nex.Client, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *nex.ResultRange) uint32
	retrieveAllMessagesWithinRangeHandler func(err error, client *nex.Client, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *nex.ResultRange) uint32
	retrieveMessagesHandler               func(err error, client *nex.Client, callID uint32, recipient *messaging_types.MessageRecipient, lstMsgIDs []uint32, bLeaveOnServer bool) uint32
	deleteMessagesHandler                 func(err error, client *nex.Client, callID uint32, recipient *messaging_types.MessageRecipient, lstMessagesToDelete []uint32) uint32
	deleteAllMessagesHandler              func(err error, client *nex.Client, callID uint32, recipient *messaging_types.MessageRecipient) uint32
	deliverMessageMultiTargetHandler      func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32 // TODO - Unknown request/response format
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
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Messaging method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new Messaging protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
