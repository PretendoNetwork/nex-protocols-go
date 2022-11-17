package nexproto

import (
	"fmt"
	"encoding/hex"

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
	M_uiRecipientType uint32
	M_principalID     uint32
	M_gatheringID     uint32
}

// ExtractFromStream extracts a MessageRecipient structure from a stream
func (messageRecipient *MessageRecipient) ExtractFromStream(stream *nex.StreamIn) error {
	messageRecipient.M_uiRecipientType = stream.ReadUInt32LE()
	messageRecipient.M_principalID = stream.ReadUInt32LE()
	messageRecipient.M_gatheringID = stream.ReadUInt32LE()

	return nil
}

// NewMessageRecipient returns a new MessageRecipient
func NewMessageRecipient() *MessageRecipient {
	return &MessageRecipient{}
}

type UserMessage struct {
	nex.Structure
	*nex.Data
	hierarchy          []nex.StructureInterface
	M_uiID             uint32
	M_uiParentID       uint32
	M_pidSender        uint32
	M_receptiontime    *nex.DateTime
	M_uiLifeTime       uint32
	M_uiFlags          uint32
	M_strSubject       string
	M_strSender        string
	M_messageRecipient *MessageRecipient
}

// ExtractFromStream extracts a UserMessage structure from a stream
func (userMessage *UserMessage) ExtractFromStream(stream *nex.StreamIn) error {
	fmt.Println("test")
	fmt.Println(stream.ReadUInt32LE())
	userMessage.M_uiID = stream.ReadUInt32LE()
	userMessage.M_uiParentID = stream.ReadUInt32LE()
	fmt.Println(stream.ReadUInt32LE())
	userMessage.M_pidSender = stream.ReadUInt32LE()
	userMessage.M_receptiontime = nex.NewDateTime(stream.ReadUInt64LE())
	//userMessage.M_uiLifeTime = stream.ReadUInt32LE()
	userMessage.M_uiFlags = stream.ReadUInt32LE()
	fmt.Println(stream.ReadUInt32LE())
	userMessage.M_strSubject, _ = stream.ReadString()
	userMessage.M_strSender, _ = stream.ReadString()

	return nil
}

func (userMessage *UserMessage) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(0)
	stream.WriteUInt32LE(userMessage.M_uiID)
	stream.WriteUInt32LE(userMessage.M_uiParentID)
	stream.WriteUInt32LE(0)
	stream.WriteUInt32LE(userMessage.M_pidSender)
	stream.WriteUInt64LE(userMessage.M_receptiontime.Value())
	//stream.WriteUInt32LE(userMessage.M_uiLifeTime)
	stream.WriteUInt32LE(userMessage.M_uiFlags)
	stream.WriteUInt32LE(1)
	stream.WriteString(userMessage.M_strSubject)
	stream.WriteString(userMessage.M_strSender)
	
	return stream.Bytes()
}

// NewUserMessage returns a new UserMessage
func NewUserMessage() *UserMessage {
	data := nex.NewData()

	userMessage := &UserMessage{}
	userMessage.Data = data
	userMessage.hierarchy = []nex.StructureInterface{data}

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

type TextMessage struct {
	nex.Structure
	*UserMessage
	hierarchy     []nex.StructureInterface
	M_StrTextBody string
}

func (textMessage *TextMessage) Bytes(stream *nex.StreamOut) []byte {
	lengthStream := nex.NewStreamOut(stream.Server)
	lengthStream.WriteStructure(textMessage.UserMessage)
	lengthStream.WriteString(textMessage.M_StrTextBody)
	length := len(lengthStream.Bytes())
	stream.WriteString("TextMessage")
	stream.WriteUInt32LE(uint32(length+4))
	stream.WriteUInt32LE(uint32(length))
	stream.WriteStructure(textMessage.UserMessage)
	stream.WriteString(textMessage.M_StrTextBody)
	
	return stream.Bytes()
}

// ExtractFromStream extracts a TextMessage structure from a stream
func (textMessage *TextMessage) ExtractFromStream(stream *nex.StreamIn) error {
	//messageRecipient, _ := stream.ReadStructure(NewMessageRecipient())
	textMessage.UserMessage.ExtractFromStream(stream)
	textMessage.M_StrTextBody, _ = stream.ReadString()

	return nil
}

// NewTextMessage returns a new TextMessage
func NewTextMessage() *TextMessage {
	textMessage := &TextMessage{}

	userMessage := NewUserMessage()

	textMessage.UserMessage = userMessage

	textMessage.hierarchy = []nex.StructureInterface{
		userMessage,
	}

	return textMessage
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
				go respondNotImplemented(packet, MessageDeliveryProtocolID)
				fmt.Printf("Unsupported MessageDelivery method ID: %#v\n", request.MethodID())
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
		logger.Warning("MessageDeliveryProtocol::DeliverMessage not implemented")
		go respondNotImplemented(packet, MessageDeliveryProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()
	fmt.Println(hex.EncodeToString(parameters))

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
	if dataHolderName == "TextMessage" {
		oUserMessage, _ = dataHolderContentStream.ReadStructure(NewTextMessage())
		fmt.Println(oUserMessage.(*TextMessage).M_StrTextBody)
	}

	go messageDeliveryProtocol.DeliverMessageHandler(nil, client, callID, oUserMessage)
}

// NewMessageDeliveryProtocol returns a new MessageDeliveryProtocol
func NewMessageDeliveryProtocol(server *nex.Server) *MessageDeliveryProtocol {
	messageDeliveryProtocol := &MessageDeliveryProtocol{server: server}

	messageDeliveryProtocol.Setup()

	return messageDeliveryProtocol
}
