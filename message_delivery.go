package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// MessageDeliveryProtocolID is the protocol ID for the Message Delivery protocol
	MessageDeliveryProtocolID = 0x1B

	// MessageDeliveryMethodDeliverMessage is the method ID for the method DeliverMessage
	MessageDeliveryMethodDeliverMessage = 0x1
)

// AuthenticationProtocol handles the Authentication nex protocol
type MessageDeliveryProtocol struct {
	server                *nex.Server
	DeliverMessageHandler func(err error, client *nex.Client, callID uint32, oUserMessage nex.StructureInterface)
}

type MessageRecipient struct {
	nex.Structure
	m_uiRecipientType uint32
	m_principalID     uint32
	m_gatheringID     uint32
}

// ExtractFromStream extracts a MessageRecipient structure from a stream
func (messageRecipient *MessageRecipient) ExtractFromStream(stream *nex.StreamIn) error {
	messageRecipient.m_uiRecipientType = stream.ReadUInt32LE()
	messageRecipient.m_principalID = stream.ReadUInt32LE()
	messageRecipient.m_gatheringID = stream.ReadUInt32LE()

	return nil
}

// NewMessageRecipient returns a new MessageRecipient
func NewMessageRecipient() *MessageRecipient {
	return &MessageRecipient{}
}

type UserMessage struct {
	nex.Structure
	*nex.NullData
	hierarchy          []nex.StructureInterface
	m_uiID             uint32
	m_uiParentID       uint32
	m_pidSender        uint32
	m_receptiontime    *nex.DateTime
	m_uiLifeTime       uint32
	m_uiFlags          uint32
	m_strSubject       string
	m_strSender        string
	m_messageRecipient *MessageRecipient
}

// ExtractFromStream extracts a UserMessage structure from a stream
func (userMessage *UserMessage) ExtractFromStream(stream *nex.StreamIn) error {
	userMessage.m_uiID = stream.ReadUInt32LE()
	userMessage.m_uiParentID = stream.ReadUInt32LE()
	userMessage.m_pidSender = stream.ReadUInt32LE()
	userMessage.m_receptiontime = nex.NewDateTime(stream.ReadUInt64LE())
	userMessage.m_uiLifeTime = stream.ReadUInt32LE()
	userMessage.m_uiFlags = stream.ReadUInt32LE()
	userMessage.m_strSubject, _ = stream.ReadString()
	userMessage.m_strSender, _ = stream.ReadString()
	messageRecipient, _ := stream.ReadStructure(NewMessageRecipient())
	userMessage.m_messageRecipient, _ = messageRecipient.(*MessageRecipient)

	return nil
}

// NewUserMessage returns a new UserMessage
func NewUserMessage() *UserMessage {
	userMessage := &UserMessage{}

	nullData := nex.NewNullData()

	userMessage.NullData = nullData

	userMessage.hierarchy = []nex.StructureInterface{
		nullData,
	}

	return userMessage
}

type BinaryMessage struct {
	nex.Structure
	*UserMessage
	hierarchy    []nex.StructureInterface
	m_binaryBody []byte
}

func (binaryMessage *BinaryMessage) Bytes(stream *nex.StreamOut) []byte {
	return []byte{}
}

// ExtractFromStream extracts a BinaryMessage structure from a stream
func (binaryMessage *BinaryMessage) ExtractFromStream(stream *nex.StreamIn) error {
	binaryMessage.m_binaryBody, _ = stream.ReadQBuffer()

	return nil
}

// NewBinaryMessage returns a new BinaryMessage
func NewBinaryMessage() *BinaryMessage {
	binaryMessage := &BinaryMessage{}

	userMessage := NewUserMessage()

	binaryMessage.UserMessage = userMessage

	binaryMessage.hierarchy = []nex.StructureInterface{
		userMessage,
	}

	return binaryMessage
}

// Setup initializes the protocol
func (messageDeliveryProtocol *MessageDeliveryProtocol) Setup() {
	nexServer := messageDeliveryProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if MessageDeliveryProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case MessageDeliveryMethodDeliverMessage:
				go messageDeliveryProtocol.handleDeliverMessage(packet)
			default:
				fmt.Printf("Unsupported MessageDelivery method ID: %#v\n", request.MethodID())
				go respondNotImplemented(packet, MessageDeliveryProtocolID)
			}
		}
	})
}

// DeliverMessage sets the DeliverMessage handler function
func (messageDeliveryProtocol *MessageDeliveryProtocol) DeliverMessage(handler func(err error, client *nex.Client, callID uint32, oUserMessage nex.StructureInterface)) {
	messageDeliveryProtocol.DeliverMessageHandler = handler
}

func (messageDeliveryProtocol *MessageDeliveryProtocol) handleDeliverMessage(packet nex.PacketInterface) {
	if messageDeliveryProtocol.DeliverMessageHandler == nil {
		fmt.Println("[Warning] MessageDeliveryProtocol::DeliverMessage not implemented")
		go respondNotImplemented(packet, MessageDeliveryProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, messageDeliveryProtocol.server)

	dataHolderName, err := parametersStream.ReadString()

	if err != nil {
		go messageDeliveryProtocol.DeliverMessageHandler(err, client, callID, nil)
		return
	}

	_ = parametersStream.ReadUInt32LE() // length including this field

	dataHolderContent, err := parametersStream.ReadBuffer()

	if err != nil {
		go messageDeliveryProtocol.DeliverMessageHandler(err, client, callID, nil)
		return
	}

	dataHolderContentStream := nex.NewStreamIn(dataHolderContent, messageDeliveryProtocol.server)

	var oUserMessage nex.StructureInterface

	if dataHolderName == "BinaryMessage" {
		oUserMessage, _ = dataHolderContentStream.ReadStructure(NewBinaryMessage())
	}

	go messageDeliveryProtocol.DeliverMessageHandler(nil, client, callID, oUserMessage)
}

// NewMessageDeliveryProtocol returns a new MessageDeliveryProtocol
func NewMessageDeliveryProtocol(server *nex.Server) *MessageDeliveryProtocol {
	messageDeliveryProtocol := &MessageDeliveryProtocol{server: server}

	messageDeliveryProtocol.Setup()

	return messageDeliveryProtocol
}
