package message_delivery_types

import (
	"bytes"
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

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

	copied.UserMessage = binaryMessage.UserMessage.Copy().(*UserMessage)
	copied.SetParentType(copied.UserMessage)
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
