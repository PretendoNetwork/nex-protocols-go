// Package types implements all the types used by the GameSession protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GameSessionInvitationReceived is a type with the GameSession protocol
type GameSessionInvitationReceived struct {
	types.Structure
	SessionKey   GameSessionKey
	SenderPID    types.PID
	Message      types.String
	CreationTime types.DateTime
}

// WriteTo writes the GameSessionInvitationReceived to the given writable
func (gsis GameSessionInvitationReceived) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gsis.SessionKey.WriteTo(contentWritable)
	gsis.SenderPID.WriteTo(contentWritable)
	gsis.Message.WriteTo(contentWritable)
	gsis.CreationTime.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gsis.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameSessionInvitationReceived from the given readable
func (gsis *GameSessionInvitationReceived) ExtractFrom(readable types.Readable) error {
	if err := gsis.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitationReceived header. %s", err.Error())
	}

	if err := gsis.SessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitationReceived.m_sessionKey. %s", err.Error())
	}

	if err := gsis.SenderPID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitationReceived.m_senderPID. %s", err.Error())
	}

	if err := gsis.Message.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitationReceived.m_message. %s", err.Error())
	}

	if err := gsis.CreationTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionInvitationReceived.m_creationTime. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameSessionInvitationReceived
func (gsis GameSessionInvitationReceived) Copy() types.RVType {
	copied := NewGameSessionInvitationReceived()

	copied.StructureVersion = gsis.StructureVersion
	copied.SessionKey = gsis.SessionKey.Copy().(GameSessionKey)
	copied.SenderPID = gsis.SenderPID
	copied.Message = gsis.Message
	copied.CreationTime = gsis.CreationTime

	return copied
}

// Equals checks if the given GameSessionInvitationReceived contains the same data as the current GameSessionInvitationReceived
func (gsis GameSessionInvitationReceived) Equals(o types.RVType) bool {
	if _, ok := o.(GameSessionInvitationReceived); !ok {
		return false
	}

	other := o.(GameSessionInvitationReceived)

	if gsis.StructureVersion != other.StructureVersion {
		return false
	}

	if !gsis.SessionKey.Equals(other.SessionKey) {
		return false
	}

	if !gsis.SenderPID.Equals(other.SenderPID) {
		return false
	}

	if gsis.Message != other.Message {
		return false
	}

	return gsis.CreationTime == other.CreationTime
}

// CopyRef copies the current value of the GameSessionInvitationReceived
// and returns a pointer to the new copy
func (gsis GameSessionInvitationReceived) CopyRef() types.RVTypePtr {
	copied := gsis.Copy().(GameSessionInvitationReceived)
	return &copied
}

// Deref takes a pointer of GameSessionInvitationReceived
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gsis *GameSessionInvitationReceived) Deref() types.RVType {
	return *gsis
}

// String returns the string representation of the GameSessionInvitationReceived
func (gsis GameSessionInvitationReceived) String() string {
	return gsis.FormatToString(0)
}

// FormatToString pretty-prints the GameSessionInvitationReceived using the provided indentation level
func (gsis GameSessionInvitationReceived) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameSessionInvitationReceived{\n")
	fmt.Fprintf(&b, "%sm_sessionKey: %s,\n", indentationValues, gsis.SessionKey)
	fmt.Fprintf(&b, "%sm_senderPID: %s,\n", indentationValues, gsis.SenderPID)
	fmt.Fprintf(&b, "%sm_message: %s,\n", indentationValues, gsis.Message)
	fmt.Fprintf(&b, "%sm_creationTime: %s,\n", indentationValues, gsis.CreationTime)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewGameSessionInvitationReceived returns a new GameSessionInvitationReceived
func NewGameSessionInvitationReceived() GameSessionInvitationReceived {
	return GameSessionInvitationReceived{
		SessionKey:   NewGameSessionKey(),
		SenderPID:    types.NewPID(0),
		Message:      types.NewString(""),
		CreationTime: types.NewDateTime(0),
	}
}
