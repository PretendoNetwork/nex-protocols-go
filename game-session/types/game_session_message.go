// Package types implements all the types used by the GameSession protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GameSessionMessage is a type with the GameSession protocol
type GameSessionMessage struct {
	types.Structure
	SessionKey GameSessionKey
	Message    types.String
}

// WriteTo writes the GameSessionMessage to the given writable
func (gsm GameSessionMessage) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gsm.SessionKey.WriteTo(contentWritable)
	gsm.Message.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gsm.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameSessionMessage from the given readable
func (gsm *GameSessionMessage) ExtractFrom(readable types.Readable) error {
	if err := gsm.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionMessage header. %s", err.Error())
	}

	if err := gsm.SessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionMessage.m_sessionKey. %s", err.Error())
	}

	if err := gsm.Message.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionMessage.m_message. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameSessionMessage
func (gsm GameSessionMessage) Copy() types.RVType {
	copied := NewGameSessionMessage()

	copied.StructureVersion = gsm.StructureVersion
	copied.SessionKey = gsm.SessionKey.Copy().(GameSessionKey)
	copied.Message = gsm.Message

	return copied
}

// Equals checks if the given GameSessionMessage contains the same data as the current GameSessionMessage
func (gsm GameSessionMessage) Equals(o types.RVType) bool {
	if _, ok := o.(GameSessionMessage); !ok {
		return false
	}

	other := o.(GameSessionMessage)

	if gsm.StructureVersion != other.StructureVersion {
		return false
	}

	if !gsm.SessionKey.Equals(other.SessionKey) {
		return false
	}

	return gsm.Message == other.Message
}

// CopyRef copies the current value of the GameSessionMessage
// and returns a pointer to the new copy
func (gsm GameSessionMessage) CopyRef() types.RVTypePtr {
	copied := gsm.Copy().(GameSessionMessage)
	return &copied
}

// Deref takes a pointer of GameSessionMessage
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gsm *GameSessionMessage) Deref() types.RVType {
	return *gsm
}

// String returns the string representation of the GameSessionMessage
func (gsm GameSessionMessage) String() string {
	return gsm.FormatToString(0)
}

// FormatToString pretty-prints the GameSessionMessage using the provided indentation level
func (gsm GameSessionMessage) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameSessionMessage{\n")
	fmt.Fprintf(&b, "%sm_sessionKey: %s, \n", indentationValues, gsm.SessionKey)
	fmt.Fprintf(&b, "%sm_message: %s,\n", indentationValues, gsm.Message)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewGameSessionMessage returns a new GameSessionMessage
func NewGameSessionMessage() GameSessionMessage {
	return GameSessionMessage{
		SessionKey: NewGameSessionKey(),
		Message:    types.NewString(""),
	}
}
