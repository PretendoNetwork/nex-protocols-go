package message_delivery_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

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

	copied.Data = userMessage.Data.Copy().(*nex.Data)
	copied.SetParentType(copied.Data)
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
