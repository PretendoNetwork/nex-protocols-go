// Package types implements all the types used by the GameSession protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GameSessionUpdate is a type with the GameSession protocol
type GameSessionUpdate struct {
	types.Structure
	SessionKey GameSessionKey
	Attributes types.List[Property]
}

// WriteTo writes the GameSessionUpdate to the given writable
func (gsu GameSessionUpdate) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gsu.SessionKey.WriteTo(contentWritable)
	gsu.Attributes.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gsu.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameSessionUpdate from the given readable
func (gsu *GameSessionUpdate) ExtractFrom(readable types.Readable) error {
	if err := gsu.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionUpdate header. %s", err.Error())
	}

	if err := gsu.SessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionUpdate.m_sessionKey. %s", err.Error())
	}

	if err := gsu.Attributes.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionUpdate.m_attributes. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameSessionUpdate
func (gsu GameSessionUpdate) Copy() types.RVType {
	copied := NewGameSessionUpdate()

	copied.StructureVersion = gsu.StructureVersion
	copied.SessionKey = gsu.SessionKey.Copy().(GameSessionKey)
	copied.Attributes = gsu.Attributes.Copy().(types.List[Property])

	return copied
}

// Equals checks if the given GameSessionUpdate contains the same data as the current GameSessionUpdate
func (gsu GameSessionUpdate) Equals(o types.RVType) bool {
	if _, ok := o.(GameSessionUpdate); !ok {
		return false
	}

	other := o.(GameSessionUpdate)

	if gsu.StructureVersion != other.StructureVersion {
		return false
	}

	if !gsu.SessionKey.Equals(other.SessionKey) {
		return false
	}

	return gsu.Attributes.Equals(other.Attributes)
}

// CopyRef copies the current value of the GameSessionUpdate
// and returns a pointer to the new copy
func (gsu GameSessionUpdate) CopyRef() types.RVTypePtr {
	copied := gsu.Copy().(GameSessionUpdate)
	return &copied
}

// Deref takes a pointer of GameSessionUpdate
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gsu *GameSessionUpdate) Deref() types.RVType {
	return *gsu
}

// String returns the string representation of the GameSessionUpdate
func (gsu GameSessionUpdate) String() string {
	return gsu.FormatToString(0)
}

// FormatToString pretty-prints the GameSessionUpdate using the provided indentation level
func (gsu GameSessionUpdate) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameSessionUpdate{\n")
	fmt.Fprintf(&b, "%sm_sessionKey: %s, \n", indentationValues, gsu.SessionKey)
	fmt.Fprintf(&b, "%sm_attributes: %s,\n", indentationValues, gsu.Attributes)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewGameSessionUpdate returns a new GameSessionUpdate
func NewGameSessionUpdate() GameSessionUpdate {
	return GameSessionUpdate{
		SessionKey: NewGameSessionKey(),
		Attributes: types.NewList[Property](),
	}
}
