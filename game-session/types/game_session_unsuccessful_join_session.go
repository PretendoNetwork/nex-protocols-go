// Package types implements all the types used by the GameSession protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GameSessionUnsuccessfulJoinSession is a type with the GameSession protocol
type GameSessionUnsuccessfulJoinSession struct {
	types.Structure
	SessionKey    GameSessionKey
	ErrorCategory types.UInt32
	ErrorCode     types.Int32
}

// WriteTo writes the GameSessionUnsuccessfulJoinSession to the given writable
func (gsujs GameSessionUnsuccessfulJoinSession) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gsujs.SessionKey.WriteTo(contentWritable)
	gsujs.ErrorCategory.WriteTo(contentWritable)
	gsujs.ErrorCode.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gsujs.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameSessionUnsuccessfulJoinSession from the given readable
func (gsujs *GameSessionUnsuccessfulJoinSession) ExtractFrom(readable types.Readable) error {
	if err := gsujs.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionUnsuccessfulJoinSession header. %s", err.Error())
	}

	if err := gsujs.SessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionUnsuccessfulJoinSession.m_sessionKey. %s", err.Error())
	}

	if err := gsujs.ErrorCategory.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionUnsuccessfulJoinSession.m_errorCategory. %s", err.Error())
	}

	if err := gsujs.ErrorCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionUnsuccessfulJoinSession.m_errorCode. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameSessionUnsuccessfulJoinSession
func (gsujs GameSessionUnsuccessfulJoinSession) Copy() types.RVType {
	copied := NewGameSessionUnsuccessfulJoinSession()

	copied.StructureVersion = gsujs.StructureVersion
	copied.SessionKey = gsujs.SessionKey.Copy().(GameSessionKey)
	copied.ErrorCategory = gsujs.ErrorCategory
	copied.ErrorCode = gsujs.ErrorCode

	return copied
}

// Equals checks if the given GameSessionUnsuccessfulJoinSession contains the same data as the current GameSessionUnsuccessfulJoinSession
func (gsujs GameSessionUnsuccessfulJoinSession) Equals(o types.RVType) bool {
	if _, ok := o.(GameSessionUnsuccessfulJoinSession); !ok {
		return false
	}

	other := o.(GameSessionUnsuccessfulJoinSession)

	if gsujs.StructureVersion != other.StructureVersion {
		return false
	}

	if !gsujs.SessionKey.Equals(other.SessionKey) {
		return false
	}

	if gsujs.ErrorCategory != other.ErrorCategory {
		return false
	}

	return gsujs.ErrorCode == other.ErrorCode
}

// CopyRef copies the current value of the GameSessionUnsuccessfulJoinSession
// and returns a pointer to the new copy
func (gsujs GameSessionUnsuccessfulJoinSession) CopyRef() types.RVTypePtr {
	copied := gsujs.Copy().(GameSessionUnsuccessfulJoinSession)
	return &copied
}

// Deref takes a pointer of GameSessionUnsuccessfulJoinSession
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gsujs *GameSessionUnsuccessfulJoinSession) Deref() types.RVType {
	return *gsujs
}

// String returns the string representation of the GameSessionUnsuccessfulJoinSession
func (gsujs GameSessionUnsuccessfulJoinSession) String() string {
	return gsujs.FormatToString(0)
}

// FormatToString pretty-prints the GameSessionUnsuccessfulJoinSession using the provided indentation level
func (gsujs GameSessionUnsuccessfulJoinSession) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameSessionUnsuccessfulJoinSession{\n")
	fmt.Fprintf(&b, "%sm_sessionKey: %s, \n", indentationValues, gsujs.SessionKey)
	fmt.Fprintf(&b, "%sm_errorCategory: %s,\n", indentationValues, gsujs.ErrorCategory)
	fmt.Fprintf(&b, "%sm_errorCode: %s,\n", indentationValues, gsujs.ErrorCode)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewGameSessionUnsuccessfulJoinSession returns a new GameSessionUnsuccessfulJoinSession
func NewGameSessionUnsuccessfulJoinSession() GameSessionUnsuccessfulJoinSession {
	return GameSessionUnsuccessfulJoinSession{
		SessionKey:    NewGameSessionKey(),
		ErrorCategory: types.NewUInt32(0),
		ErrorCode:     types.NewInt32(0),
	}
}
