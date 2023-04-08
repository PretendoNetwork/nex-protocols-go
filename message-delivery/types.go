package message_delivery

import nex "github.com/PretendoNetwork/nex-go"

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
	*nex.Data
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
