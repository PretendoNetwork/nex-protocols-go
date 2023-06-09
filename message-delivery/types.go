package message_delivery

import (
	"bytes"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

type MessageRecipient struct {
	nex.Structure
	m_uiRecipientType uint32
	m_principalID     uint32
	m_gatheringID     uint32
}

// ExtractFromStream extracts a MessageRecipient structure from a stream
func (messageRecipient *MessageRecipient) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	messageRecipient.m_uiRecipientType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MessageRecipient.m_uiRecipientType from stream. %s", err.Error())
	}

	messageRecipient.m_principalID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MessageRecipient.m_principalID from stream. %s", err.Error())
	}

	messageRecipient.m_gatheringID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract MessageRecipient.m_gatheringID from stream. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MessageRecipient
func (messageRecipient *MessageRecipient) Copy() nex.StructureInterface {
	copied := NewMessageRecipient()

	copied.m_uiRecipientType = messageRecipient.m_uiRecipientType
	copied.m_principalID = messageRecipient.m_principalID
	copied.m_gatheringID = messageRecipient.m_gatheringID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (messageRecipient *MessageRecipient) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MessageRecipient)

	if messageRecipient.m_uiRecipientType != other.m_uiRecipientType {
		return false
	}

	if messageRecipient.m_principalID != other.m_principalID {
		return false
	}

	if messageRecipient.m_gatheringID != other.m_gatheringID {
		return false
	}

	return true
}

// NewMessageRecipient returns a new MessageRecipient
func NewMessageRecipient() *MessageRecipient {
	return &MessageRecipient{}
}

type UserMessage struct {
	nex.Structure
	*nex.Data
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
	var err error

	userMessage.m_uiID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.m_uiID from stream. %s", err.Error())
	}

	userMessage.m_uiParentID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.m_uiParentID from stream. %s", err.Error())
	}

	userMessage.m_pidSender, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.m_pidSender from stream. %s", err.Error())
	}

	userMessage.m_receptiontime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.m_receptiontime from stream. %s", err.Error())
	}

	userMessage.m_uiLifeTime, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.m_uiLifeTime from stream. %s", err.Error())
	}

	userMessage.m_uiFlags, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.m_uiFlags from stream. %s", err.Error())
	}

	userMessage.m_strSubject, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.m_strSubject from stream. %s", err.Error())
	}

	userMessage.m_strSender, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.m_strSender from stream. %s", err.Error())
	}

	messageRecipient, err := stream.ReadStructure(NewMessageRecipient())
	if err != nil {
		return fmt.Errorf("Failed to extract UserMessage.m_messageRecipient from stream. %s", err.Error())
	}

	userMessage.m_messageRecipient = messageRecipient.(*MessageRecipient)

	return nil
}

// Copy returns a new copied instance of UserMessage
func (userMessage *UserMessage) Copy() nex.StructureInterface {
	copied := NewUserMessage()

	copied.SetParentType(userMessage.ParentType().Copy())
	copied.m_uiID = userMessage.m_uiID
	copied.m_uiParentID = userMessage.m_uiParentID
	copied.m_pidSender = userMessage.m_pidSender
	copied.m_receptiontime = userMessage.m_receptiontime.Copy()
	copied.m_uiLifeTime = userMessage.m_uiLifeTime
	copied.m_uiFlags = userMessage.m_uiFlags
	copied.m_strSubject = userMessage.m_strSubject
	copied.m_strSender = userMessage.m_strSender
	copied.m_messageRecipient = userMessage.m_messageRecipient.Copy().(*MessageRecipient)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (userMessage *UserMessage) Equals(structure nex.StructureInterface) bool {
	other := structure.(*UserMessage)

	if !userMessage.ParentType().Equals(other.ParentType()) {
		return false
	}

	if userMessage.m_uiID != other.m_uiID {
		return false
	}

	if userMessage.m_uiParentID != other.m_uiParentID {
		return false
	}

	if userMessage.m_pidSender != other.m_pidSender {
		return false
	}

	if !userMessage.m_receptiontime.Equals(other.m_receptiontime) {
		return false
	}

	if userMessage.m_uiLifeTime != other.m_uiLifeTime {
		return false
	}

	if userMessage.m_uiFlags != other.m_uiFlags {
		return false
	}

	if userMessage.m_strSubject != other.m_strSubject {
		return false
	}

	if userMessage.m_strSender != other.m_strSender {
		return false
	}

	if !userMessage.m_messageRecipient.Equals(other.m_messageRecipient) {
		return false
	}

	return true
}

// NewUserMessage returns a new UserMessage
func NewUserMessage() *UserMessage {
	userMessage := &UserMessage{}
	userMessage.Data = nex.NewData()
	userMessage.SetParentType(userMessage.Data)

	return userMessage
}

type BinaryMessage struct {
	nex.Structure
	*UserMessage
	m_binaryBody []byte
}

// Bytes encodes the BinaryMessage and returns a byte array
func (binaryMessage *BinaryMessage) Bytes(stream *nex.StreamOut) []byte {
	return []byte{}
}

// ExtractFromStream extracts a BinaryMessage structure from a stream
func (binaryMessage *BinaryMessage) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	binaryMessage.m_binaryBody, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract BinaryMessage.m_binaryBody from stream. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of BinaryMessage
func (binaryMessage *BinaryMessage) Copy() nex.StructureInterface {
	copied := NewBinaryMessage()

	copied.SetParentType(binaryMessage.ParentType().Copy())
	copied.m_binaryBody = make([]byte, len(binaryMessage.m_binaryBody))

	copy(copied.m_binaryBody, binaryMessage.m_binaryBody)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (binaryMessage *BinaryMessage) Equals(structure nex.StructureInterface) bool {
	other := structure.(*BinaryMessage)

	if !binaryMessage.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !bytes.Equal(binaryMessage.m_binaryBody, other.m_binaryBody) {
		return false
	}

	return true
}

// NewBinaryMessage returns a new BinaryMessage
func NewBinaryMessage() *BinaryMessage {
	binaryMessage := &BinaryMessage{}
	binaryMessage.UserMessage = NewUserMessage()
	binaryMessage.SetParentType(binaryMessage.UserMessage)

	return binaryMessage
}
